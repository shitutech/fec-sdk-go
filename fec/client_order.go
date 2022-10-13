package fec

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/shitutech/fec-sdk-go/models"
)

// OrderPay 订单支付
func (s *Client) OrderPay(request *models.OrderPayRequest) (*models.OrderPayResponse, error) {
	_result := &models.OrderPayResponse{}

	// 设置服务商号
	request.ProviderNo = s.config.ProviderNo()
	request.TaskCode = s.config.TaskCode()
	request.ProductCode = s.config.ProductNo()

	encodeData, err := json.Marshal(request)
	if err != nil {
		return _result, errors.New("业务数据 JSON 编码失败")
	}

	respData, err := s.doRequest(string(encodeData), "/api/fec/order/pay")
	if err != nil {
		return _result, err
	}

	type respStruct struct {
		commonResp
		Result models.OrderPayResponse `json:"result,omitempty"`
	}
	var respObj respStruct

	err = json.Unmarshal([]byte(respData), &respObj)
	if err != nil {
		return _result, errors.New("响应数据 JSON 解码失败")
	}

	if respObj.Code != 200 {
		return _result, errors.New(fmt.Sprintf("上游服务响应报告异常。Err: %d::%s", respObj.Code, respObj.Message))
	}

	if respObj.Result.StatusCode != "1001" && respObj.Result.StatusCode != "1002" {
		return _result, errors.New(fmt.Sprintf("三方服务响应报告异常。Err: %s:::%s",
			respObj.Result.StatusCode, respObj.Result.Msg))
	}

	return &respObj.Result, nil
}

// OrderBatchQuery 批次订单号查询
func (s *Client) OrderBatchQuery(request *models.OrderBatchQueryRequest) (*models.OrderBatchQueryResponse, error) {
	_result := &models.OrderBatchQueryResponse{}

	encodeData, err := json.Marshal(request)
	if err != nil {
		return _result, errors.New("业务数据 JSON 编码失败")
	}

	respData, err := s.doRequest(string(encodeData), "/api/fec/order/query/batch")
	if err != nil {
		return _result, err
	}

	type respStruct struct {
		commonResp
		Result models.OrderBatchQueryResponse `json:"result,omitempty"`
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

// OrderSubQuery 子订单详情查询
func (s *Client) OrderSubQuery(request *models.OrderSubQueryRequest) (*models.OrderSubQueryResponse, error) {
	_result := &models.OrderSubQueryResponse{}

	encodeData, err := json.Marshal(request)
	if err != nil {
		return _result, errors.New("业务数据 JSON 编码失败")
	}

	respData, err := s.doRequest(string(encodeData), "/api/fec/order/query/detail")
	if err != nil {
		return _result, err
	}

	type respStruct struct {
		commonResp
		Result models.OrderSubQueryResponse `json:"result,omitempty"`
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
