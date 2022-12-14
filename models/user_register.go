package models

type UserRegisterRequest struct {
	// 服务商号。设置会被覆盖
	ProviderNo string `json:"providerNo"`
	// 姓名
	Name string `json:"name"`
	// 身份证号
	IdCard string `json:"idCard"`
	// 手机号
	Phone string `json:"phone"`
	// 身份证正面 base64格式 Base64编码后的大小不超2M
	ImgIdCardFront string `json:"imgIdCardFront"`
	// 身份证反面 base64格式 Base64编码后的大小不超2M
	ImgIdCardBack string `json:"imgIdCardBack"`
	// 业务类型 1委托代征2.个体户注册(分包)3.自然人代开4.临时税务登记
	BusType string `json:"busType"`
	// 面部高清照，格式要求：Base64值；Base64编码后的大小不超2M
	ImgFace string `json:"imgFace,omitempty"`
}

type UserRegisterResponse struct {
	BizCommonResponse
	// 商户会员ID
	MemberId string `json:"memberId,omitempty"`
	// 系统会员ID
	SystemId string `json:"systemId,omitempty"`
}
