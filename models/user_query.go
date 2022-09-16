package models

type UserQueryRequest struct {
	// 查询类型 1:查询基本信息 2:查询开户信息 3:查询会员ID
	Type string `json:"type"`
	// 会员ID （查询类型 1/2必传）
	MemberId string `json:"memberId,omitempty"`
	// 支付通道 1:众邦 （查询类型 2 必传）
	PayPass string `json:"payPass,omitempty"`
	// 身份证号 （查询类型 3 必传）
	IdCard string `json:"idCard,omitempty"`
}

type UserQueryResponse struct {
	BizCommonResponse
	// 姓名 （查询类型 1 返回）
	Name string `json:"name,omitempty"`
	// 身份证号 （查询类型 1 返回）
	IdCard string `json:"idCard,omitempty"`
	// 业务类型 1委托代征2.个体户注册(分包)3.自然人代开4.临时税务登记 （查询类型 1 返回）
	BusTypeList interface{} `json:"busTypeList,omitempty"`
	// 开户行编号 （查询类型 1 返回）
	BankNo string `json:"bankNo,omitempty"`
	// 开户行 （查询类型1 返回）
	BankName string `json:"bankName,omitempty"`
	// 卡号 （查询类型 1 返回）
	CardNo string `json:"cardNo,omitempty"`
	// 银行预留手机号 （查询类型 1 返回）
	BankPhone string `json:"bankPhone,omitempty"`
	// 账户状态 0:未开通 1:开通中 2:开通成功 3:开通失败 （查询类型 2 返回）
	AcctStatus string `json:"acctStatus,omitempty"`
	// 会员ID （查询类型 3 返回）
	MemberId string `json:"memberId,omitempty"`
}
