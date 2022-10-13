package fec

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/shitutech/fec-sdk-go/v2/models"
)

// SupportBankList 查询通道支持的银行列表
func (s *Client) SupportBankList(request *models.SupportBankListRequest) (*models.SupportBankListResponse, error) {
	_result := &models.SupportBankListResponse{}

	encodeData, err := json.Marshal(request)
	if err != nil {
		return _result, errors.New("业务数据 JSON 编码失败")
	}

	respData, err := s.doRequest(string(encodeData), "/api/fec/v2/bank/list")
	if err != nil {
		return _result, err
	}

	type respStruct struct {
		commonResp
		Result models.SupportBankListResponse `json:"result,omitempty"`
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
