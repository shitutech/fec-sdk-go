package fec

import (
	"bytes"
	"crypto"
	"crypto/aes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/andreburgaud/crypt2go/ecb"
	"github.com/shitutech/fec-sdk-go/helpers"
	"github.com/shitutech/fec-sdk-go/models"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type commonResp struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Success   bool   `json:"success"`
	Timestamp int64  `json:"timestamp"`
}

type Client struct {
	// 配置
	config *Config
}

func NewClient(config *Config) *Client {
	s := new(Client)

	s.config = config

	return s
}

// AcctInfo 商户信息查询
func (s *Client) AcctInfo(request *models.AcctInfoRequest) (*models.AcctInfoResponse, error) {
	_result := &models.AcctInfoResponse{}

	encodeData, err := json.Marshal(request)
	if err != nil {
		return _result, errors.New("业务数据 JSON 编码失败")
	}

	respData, err := s.doRequest(string(encodeData), "/api/fec/acct/info")
	if err != nil {
		return _result, err
	}

	// {"code":400,"message":"signature error","success":false,"timestamp":1663324104}
	// {"code":200,"message":"操作成功！","success":true,"timestamp":1663323110463,"result":{"statusCode":"1000","msg":null,"balance"0","availableFee":"0.00","frozenBalance":"0"}}
	type respStruct struct {
		commonResp
		Result models.AcctInfoResponse `json:"result,omitempty"`
	}
	var respObj respStruct

	err = json.Unmarshal([]byte(respData), &respObj)
	if err != nil {
		return _result, errors.New("响应数据 JSON 解码失败")
	}

	if respObj.Code != 200 {
		return _result, errors.New(fmt.Sprintf("上游服务响应报告异常。Err: %d::%s", respObj.Code, respObj.Message))
	}

	if respObj.Result.StatusCode != "1000" {
		return _result, errors.New(fmt.Sprintf("三方服务响应报告异常。Err: %s:::%s",
			respObj.Result.StatusCode, respObj.Result.Msg))
	}

	return &respObj.Result, nil
}

func (s *Client) doRequest(bizData string, apiPath string) (string, error) {
	reqParams := make(map[string]string)
	reqParams["requestTime"] = strings.Replace(time.Now().Format("20060102150405.000"), ".", "", 1)
	reqParams["nonce"] = helpers.GetRandomString(32)
	reqParams["merchantNo"] = s.config.MerchantNo()

	aesKey := helpers.GetRandomString(24)

	// 业务数据加密
	encryptedBizData, err := s.encryptBizData(&bizData, &aesKey)
	if err != nil {
		return "", err
	}

	// aes密钥加密
	encryptedAesKey, err := s.encryptAesKey(&aesKey)
	if err != nil {
		return "", err
	}

	// 签名计算
	signStr := fmt.Sprintf("%s%s%s%s%s", reqParams["requestTime"], reqParams["nonce"],
		reqParams["merchantNo"], encryptedBizData, encryptedAesKey)

	sign, err := s.signRsa(&signStr)
	if err != nil {
		return "", err
	}

	reqParams["requestData"] = encryptedBizData
	reqParams["encryptKey"] = encryptedAesKey
	reqParams["sign"] = sign

	reqStr, err := json.Marshal(reqParams)
	if err != nil {
		return "", err
	}

	httpClient := &http.Client{Timeout: 30 * time.Second}
	defer httpClient.CloseIdleConnections()

	url := "https://fec.51wanquan.com" + apiPath
	req, err := http.NewRequest("POST", url, strings.NewReader(string(reqStr)))
	if err != nil {
		return "", err
	}

	req.Header.Add("User-Agent", fmt.Sprintf("fecSDKGo/v%s", VERSION))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		return "", errors.New("HTTP " + strconv.Itoa(resp.StatusCode))
	}

	respBodyStr := strings.TrimSpace(string(respBody))
	if len(respBodyStr) < 1 {
		return "", errors.New("响应为空")
	}

	return respBodyStr, nil
}

func (s *Client) encryptBizData(bizData *string, aesKey *string) (string, error) {
	cipher, err := aes.NewCipher([]byte(*aesKey))
	if err != nil {
		return "", err
	}

	mode := ecb.NewECBEncrypter(cipher)
	//padder := padding.NewPkcs5Padding()
	//pt, err := padder.Pad([]byte(*bizData))
	//if err != nil {
	//	return "", err
	//}

	pt := s.PKCS5Padding([]byte(*bizData), mode.BlockSize())

	ct := make([]byte, len(pt))
	mode.CryptBlocks(ct, pt)

	hashData := base64.StdEncoding.EncodeToString(ct)

	return hashData, nil
}

func (s *Client) PKCS5Padding(cipherText []byte, blockSize int) []byte {
	paddingLen := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(paddingLen)}, paddingLen)
	return append(cipherText, padText...)
}

func (s *Client) encryptAesKey(aesKey *string) (string, error) {
	pubKey := s.config.SystemPublicKey()
	if strings.Contains(pubKey, "\n") || strings.Contains(pubKey, "-") {
		return "", errors.New("所配置的公钥只能是单一的字符串，不可包含换行符和公钥头尾标识")
	}

	// 解密pem格式的公钥
	block, _ := pem.Decode([]byte(s.wordWrapForRsa(pubKey, false)))
	if block == nil {
		return "", errors.New("所配置的公钥错误")
	}

	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return "", err
	}

	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	//加密
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, pub, []byte(*aesKey))

	hashData := base64.StdEncoding.EncodeToString(ciphertext)

	return hashData, nil
}

func (s *Client) signRsa(signStr *string) (string, error) {
	privKey := s.config.PrivateKey()
	if strings.Contains(privKey, "\n") || strings.Contains(privKey, "-") {
		return "", errors.New("所配置的私钥只能是单一的字符串，不可包含换行符和私钥头尾标识")
	}

	// 解密pem格式的私钥
	block, _ := pem.Decode([]byte(s.wordWrapForRsa(privKey, true)))
	if block == nil {
		return "", errors.New("所配置的私钥错误")
	}

	// 解析私钥
	var privInterface interface{}
	privInterface, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		privInterface, err = x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			return "", err
		}
	}

	priv := privInterface.(*rsa.PrivateKey)

	hash := sha1.New()
	_, err = hash.Write([]byte(*signStr))
	if err != nil {
		return "", err
	}

	hashSum := hash.Sum(nil)

	signature, err := rsa.SignPKCS1v15(rand.Reader, priv, crypto.SHA1, hashSum)
	if err != nil {
		return "", err
	}

	signatureEncode := base64.StdEncoding.EncodeToString(signature)

	return signatureEncode, nil
}

func (s *Client) wordWrapForRsa(rsaStr string, isPrivate bool) string {
	var lines []string
	if !isPrivate {
		lines = append(lines, "-----BEGIN PUBLIC KEY-----")
	} else {
		lines = append(lines, "-----BEGIN PRIVATE KEY-----") // pkcs8
	}

	for {
		if len(rsaStr) > 64 {
			lines = append(lines, rsaStr[:64])
			rsaStr = rsaStr[64:]
		} else {
			lines = append(lines, rsaStr[:])
			break
		}
	}

	if !isPrivate {
		lines = append(lines, "-----END PUBLIC KEY-----")
	} else {
		lines = append(lines, "-----END PRIVATE KEY-----")
	}

	return strings.Join(lines, "\n")
}
