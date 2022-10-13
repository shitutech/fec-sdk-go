package models

type AcctListRequest struct {
	// 账户类型1商户开户2服务商开户
	AccountType string `json:"accountType"`
	// 支付通道（1:众邦 2：支付宝 4.招商银行）
	PayPass string `json:"payPass"`
	// 服务商编号
	ProviderNo string `json:"providerNo"`
}

type payAccountList struct {
	// 商户编号
	MerchantNo string `json:"merchantNo"`
	// 账户名称
	AccountName string `json:"accountName"`
	// 账户号
	AccountNo string `json:"accountNo"`
	// 支付通道1.众邦2.支付宝4.招商银行
	PayPass string `json:"payPass"`
	// 账户类型【1.商户开户1服务商开户】
	AccountType string `json:"accountType"`
	// 服务商编号
	ProviderNo string `json:"providerNo"`
	// 服务商名称
	ProviderName string `json:"providerName"`
}

type AcctListResponse struct {
	BizCommonResponse
	PayAccountList []payAccountList `json:"payAccountList"`
}
