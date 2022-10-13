package fec

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/shitutech/fec-sdk-go/v2/models"
)

// AcctInfo 商户账户信息查询
func (s *Client) AcctInfo(request *models.AcctInfoRequest) (*models.AcctInfoResponse, error) {
	_result := &models.AcctInfoResponse{}

	encodeData, err := json.Marshal(request)
	if err != nil {
		return _result, errors.New("业务数据 JSON 编码失败")
	}

	respData, err := s.doRequest(string(encodeData), "/api/fec/v2/acct/info")
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

// AcctList 商户账户列表查询
func (s *Client) AcctList(request *models.AcctListRequest) (*models.AcctListResponse, error) {
	_result := &models.AcctListResponse{}

	encodeData, err := json.Marshal(request)
	if err != nil {
		return _result, errors.New("业务数据 JSON 编码失败")
	}

	respData, err := s.doRequest(string(encodeData), "/api/fec/v2/acct/query/list")
	if err != nil {
		return _result, err
	}

	fmt.Println(respData)

	type respStruct struct {
		commonResp
		Result models.AcctListResponse `json:"result,omitempty"`
	}
	var respObj respStruct

	err = json.Unmarshal([]byte(respData), &respObj)
	if err != nil {
		return _result, errors.New("响应数据 JSON 解码失败")
	}

	if respObj.Code != 200 {
		return _result, errors.New(fmt.Sprintf("上游服务响应报告异常。Err: %d::%s", respObj.Code, respObj.Message))
	}

	if respObj.Result.StatusCode != "" {
		return _result, errors.New(fmt.Sprintf("三方服务响应报告异常。Err: %s:::%s",
			respObj.Result.StatusCode, respObj.Result.Msg))
	}

	return &respObj.Result, nil
}
