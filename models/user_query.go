package models

type UserQueryRequest struct {
	// 系统会员 ID
	SystemId string `json:"systemId"`
	// 支付通道（1:众邦 2:支付宝 4.招商银行）
	PayPass string `json:"payPass"`
}

type UserQueryResponse struct {
	BizCommonResponse
	// 姓名
	Name string `json:"name"`
	// 身份证号
	IdCard string `json:"idCard"`
	// 开户行编号 （ payPass不为空且绑定过银行卡返回）
	BankNo string `json:"bankNo,omitempty"`
	// 开户行 （payPass不为空且绑定过银行卡返回）
	BankName string `json:"bankName,omitempty"`
	// 卡号 （payPass不为空且绑定过银行卡返回）
	CardNo string `json:"cardNo,omitempty"`
	// 银行预留手机号 （payPass不为空且绑定过银行卡返回）
	BankPhone string `json:"bankPhone,omitempty"`
	// 系统会员 ID
	SystemId string `json:"systemId"`
}

type UserSystemIdRequest struct {
	// 3. 根据身份证号4.根据商户会员Id查询系统会员Id
	Type string `json:"type"`
	// 身份证号（查询类型3必传;会员身份证号）
	IdCard string `json:"idCard,omitempty"`
	// 商户会员ID（查询类型4必传;v1.0用户注册接口返回）
	MemberId string `json:"memberId,omitempty"`
}

type UserSystemIdResponse struct {
	BizCommonResponse
	// 系统会员ID
	SystemId string `json:"systemId"`
}

type UserBindCardsRequest struct {
	// 系统会员 ID
	SystemId string `json:"systemId"`
}

type UserBindCardItem struct {
	// 支付通道 （1.众邦银行2.支付宝4.招商银行）
	PayPass string `json:"payPass,omitempty"`
	// 是否通道默认提现卡【0否1是】
	IzDefault string `json:"izDefault,omitempty"`
	// 开户行编号
	BankNo string `json:"bankNo,omitempty"`
	// 开户行
	BankName string `json:"bankName,omitempty"`
	// 银行卡号
	CardNo string `json:"cardNo,omitempty"`
	// 银行预留手机号
	BankPhone string `json:"bankPhone,omitempty"`
}

type UserBindCardsResponse struct {
	BizCommonResponse
	// 绑卡列表, 详见以下list
	MemberCardList []UserBindCardItem `json:"memberCardList,omitempty"`
}
