package models

type SupportBankListRequest struct {
	// 支付通道 1:众邦 2:支付宝 4.招商银行 （ 1/2必传）
	PayPass string `json:"payPass"`
}

type SupportBankListItem struct {
	// 银行名称
	Title string `json:"title,omitempty"`
	// 银行名称
	Text string `json:"text,omitempty"`
	// 银行编号
	Value string `json:"value,omitempty"`
}

type SupportBankListResponse struct {
	BizCommonResponse
	// 通道支持的银行列表, 详见以下list
	BankList []SupportBankListItem `json:"bankList,omitempty"`
}
