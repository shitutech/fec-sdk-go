package models

type AcctInfoRequest struct {
	PayPass string `json:"payPass"`
}

func (s *AcctInfoRequest) SetPayPass(v string) *AcctInfoRequest {
	s.PayPass = v
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
