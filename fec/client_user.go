package fec

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/shitutech/fec-sdk-go/models"
)

// UserRegister 用户注册
func (s *Client) UserRegister(request *models.UserRegisterRequest) (*models.UserRegisterResponse, error) {
	_result := &models.UserRegisterResponse{}

	encodeData, err := json.Marshal(request)
	if err != nil {
		return _result, errors.New("业务数据 JSON 编码失败")
	}

	respData, err := s.doRequest(string(encodeData), "/api/fec/acct/register")
	if err != nil {
		return _result, err
	}

	type respStruct struct {
		commonResp
		Result models.UserRegisterResponse `json:"result,omitempty"`
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
