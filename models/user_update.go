package models

type UserUpdateRequest struct {
	// 会员ID
	MemberId string `json:"memberId"`
	// 变更类型 1:变更注册手机号 2:变更影像件 3:拓展业务类型 4：变更结算卡信息
	// 该值在请求时固定设置为 1
	ChangeType string `json:"changeType"`
	// 商户编号
	// 无需设置，其值将被配置覆盖
	MerchantNo string `json:"merchantNo"`
}

type UserUpdateMobileRequest struct {
	UserUpdateRequest
	// 注册手机号
	Phone string `json:"phone"`
}

type UserUpdateIdCardImageRequest struct {
	UserUpdateRequest
	// 身份证正面 base64格式
	ImgIdCardFront string `json:"imgIdCardFront"`
	// 身份证反面 base64格式
	ImgIdCardBack string `json:"imgIdCardBack"`
}

type UserUpdateBizTypeRequest struct {
	UserUpdateRequest
	// 服务商号 变更类型为3必传
	ProviderNo string `json:"providerNo"`
	// 业务类型 委托代征2.个体户注册(分包)3.自然人代开4.临时税务登记 变更类型为3必传
	BusType string `json:"busType"`
}

type UserUpdateSettlementCardRequest struct {
	UserUpdateRequest
	// 支付通道（1:众邦, 2: 支付宝） 变更类型为4必传
	PayPass string `json:"payPass"`
	// 开户行编号 变更类型为4必传
	BankNo string `json:"bankNo"`
	// 结算卡号 变更类型为4必传
	CardNo string `json:"cardNo"`
	// 银行预留手机号 变更类型为4必传
	BankPhone string `json:"bankPhone"`
	// 银行卡照片 变更类型为4必传(暂时不传, 后期上线后会有变动)
	ImgBank string `json:"imgBank,omitempty"`
}

type UserUpdateResponse struct {
	BizCommonResponse
}
