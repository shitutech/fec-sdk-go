package models

type UserUpdateSettlementCardRequest struct {
	// 系统会员 ID
	SystemId string `json:"systemId"`
	// 支付通道（1:众邦 2:支付宝 4.招商银行）
	PayPass string `json:"payPass"`
	// 开户行编号
	BankNo string `json:"bankNo"`
	// 结算卡号
	CardNo string `json:"cardNo"`
	// 银行预留手机号
	BankPhone string `json:"bankPhone"`
	// 银行卡照片
	ImgBank string `json:"imgBank,omitempty"`
}

type UserUpdateResponse struct {
	BizCommonResponse
}
