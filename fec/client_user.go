package fec

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/shitutech/fec-sdk-go/v2/models"
)

// UserRegister 用户注册
func (s *Client) UserRegister(request *models.UserRegisterRequest) (*models.UserRegisterResponse, error) {
	_result := &models.UserRegisterResponse{}

	// 设置服务商号
	request.ProviderNo = s.config.ProviderNo()

	encodeData, err := json.Marshal(request)
	if err != nil {
		return _result, errors.New("业务数据 JSON 编码失败")
	}

	respData, err := s.doRequest(string(encodeData), "/api/fec/v2/acct/register")
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

// UserQuery 会员基本信息查询
func (s *Client) UserQuery(request *models.UserQueryRequest) (*models.UserQueryResponse, error) {
	_result := &models.UserQueryResponse{}

	encodeData, err := json.Marshal(request)
	if err != nil {
		return _result, errors.New("业务数据 JSON 编码失败")
	}

	respData, err := s.doRequest(string(encodeData), "/api/fec/v2/acct/query")
	if err != nil {
		return _result, err
	}

	type respStruct struct {
		commonResp
		Result models.UserQueryResponse `json:"result,omitempty"`
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

// UserSystemId 查询会员系统id
func (s *Client) UserSystemId(request *models.UserSystemIdRequest) (*models.UserSystemIdResponse, error) {
	_result := &models.UserSystemIdResponse{}

	encodeData, err := json.Marshal(request)
	if err != nil {
		return _result, errors.New("业务数据 JSON 编码失败")
	}

	respData, err := s.doRequest(string(encodeData), "/api/fec/v2/system/id")
	if err != nil {
		return _result, err
	}

	type respStruct struct {
		commonResp
		Result models.UserSystemIdResponse `json:"result,omitempty"`
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

// UserIdSettlementCard 绑定/变更结算卡
func (s *Client) UserIdSettlementCard(request *models.UserUpdateSettlementCardRequest) (*models.UserUpdateResponse, error) {
	_result := &models.UserUpdateResponse{}

	encodeData, err := json.Marshal(request)
	if err != nil {
		return _result, errors.New("业务数据 JSON 编码失败")
	}

	return s.userUpdateRespDeal(&encodeData)
}

func (s *Client) userUpdateRespDeal(bizReqData *[]byte) (*models.UserUpdateResponse, error) {
	_result := &models.UserUpdateResponse{}

	respData, err := s.doRequest(string(*bizReqData), "/api/fec/v2/acct/update")
	if err != nil {
		return _result, err
	}

	type respStruct struct {
		commonResp
		Result models.UserUpdateResponse `json:"result,omitempty"`
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

// UserBindCards 查询会员绑定的结算卡列表
func (s *Client) UserBindCards(request *models.UserBindCardsRequest) (*models.UserBindCardsResponse, error) {
	_result := &models.UserBindCardsResponse{}

	encodeData, err := json.Marshal(request)
	if err != nil {
		return _result, errors.New("业务数据 JSON 编码失败")
	}

	respData, err := s.doRequest(string(encodeData), "/api/fec/v2/bank/cards")
	if err != nil {
		return _result, err
	}

	type respStruct struct {
		commonResp
		Result models.UserBindCardsResponse `json:"result,omitempty"`
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
