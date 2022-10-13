package models

type AcctInfoRequest struct {
	// 账户类型1商户开户2服务商开户
	AccountType string `json:"accountType"`
	// 支付通道（1:众邦 2：支付宝 4.招商银行）
	PayPass string `json:"payPass"`
	// 服务商编号
	ProviderNo string `json:"providerNo"`
}

func (s *AcctInfoRequest) SetAccountType(v string) *AcctInfoRequest {
	s.AccountType = v
	return s
}

func (s *AcctInfoRequest) SetPayPass(v string) *AcctInfoRequest {
	s.PayPass = v
	return s
}

func (s *AcctInfoRequest) SetProviderNo(v string) *AcctInfoRequest {
	s.ProviderNo = v
	return s
}

type AcctInfoResponse struct {
	BizCommonResponse
	// 账户余额 单位元，精确到两位小数点 示例： 10.05元
	Balance string `json:"balance"`
	// 可用余额 单位元，精确到两位小数点 示例： 10.05元
	AvailableFee string `json:"availableFee"`
	// 冻结金额 单位元，精确到两位小数点 示例： 10.05元
	FrozenBalance string `json:"frozenBalance"`
}

func (s *AcctInfoResponse) SetBalance(v string) *AcctInfoResponse {
	s.Balance = v
	return s
}

func (s *AcctInfoResponse) SetAvailableFee(v string) *AcctInfoResponse {
	s.AvailableFee = v
	return s
}

func (s *AcctInfoResponse) SetFrozenBalance(v string) *AcctInfoResponse {
	s.FrozenBalance = v
	return s
}
