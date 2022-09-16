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

	// 设置服务商号
	request.ProviderNo = s.config.ProviderNo()

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

// UserQuery 用户信息查询
func (s *Client) UserQuery(request *models.UserQueryRequest) (*models.UserQueryResponse, error) {
	_result := &models.UserQueryResponse{}

	encodeData, err := json.Marshal(request)
	if err != nil {
		return _result, errors.New("业务数据 JSON 编码失败")
	}

	respData, err := s.doRequest(string(encodeData), "/api/fec/acct/query")
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

// UserUpdateMobile 用户变更注册手机号
func (s *Client) UserUpdateMobile(request *models.UserUpdateMobileRequest) (*models.UserUpdateResponse, error) {
	_result := &models.UserUpdateResponse{}

	request.ChangeType = ChangeTypeMobile
	request.MerchantNo = s.config.MerchantNo()

	encodeData, err := json.Marshal(request)
	if err != nil {
		return _result, errors.New("业务数据 JSON 编码失败")
	}

	return s.userUpdateRespDeal(&encodeData)
}

// UserIdCardImage 用户影像件
func (s *Client) UserIdCardImage(request *models.UserUpdateIdCardImageRequest) (*models.UserUpdateResponse, error) {
	_result := &models.UserUpdateResponse{}

	request.ChangeType = ChangeTypeIdCardImage
	request.MerchantNo = s.config.MerchantNo()

	encodeData, err := json.Marshal(request)
	if err != nil {
		return _result, errors.New("业务数据 JSON 编码失败")
	}

	return s.userUpdateRespDeal(&encodeData)
}

// UserIdBizType 拓展业务类型
func (s *Client) UserIdBizType(request *models.UserUpdateBizTypeRequest) (*models.UserUpdateResponse, error) {
	_result := &models.UserUpdateResponse{}

	request.ChangeType = ChangeTypeBizType
	request.MerchantNo = s.config.MerchantNo()
	request.ProviderNo = s.config.ProviderNo()

	encodeData, err := json.Marshal(request)
	if err != nil {
		return _result, errors.New("业务数据 JSON 编码失败")
	}

	return s.userUpdateRespDeal(&encodeData)
}

// UserIdSettlementCard 结算卡信息
func (s *Client) UserIdSettlementCard(request *models.UserUpdateSettlementCardRequest) (*models.UserUpdateResponse, error) {
	_result := &models.UserUpdateResponse{}

	request.ChangeType = ChangeTypeSettlementCard
	request.MerchantNo = s.config.MerchantNo()

	encodeData, err := json.Marshal(request)
	if err != nil {
		return _result, errors.New("业务数据 JSON 编码失败")
	}

	return s.userUpdateRespDeal(&encodeData)
}

func (s *Client) userUpdateRespDeal(bizReqData *[]byte) (*models.UserUpdateResponse, error) {
	_result := &models.UserUpdateResponse{}

	respData, err := s.doRequest(string(*bizReqData), "/api/fec/acct/update")
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
