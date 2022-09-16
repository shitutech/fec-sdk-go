package models

type UserAccountRequest struct {
	// 会员ID
	MemberId string `json:"memberId,omitempty"`
	// 支付通道（1:众邦, 2: 支付宝）
	PayPass string `json:"payPass"`
	// 开户行编号 （众邦必传）
	BankNo string `json:"bankNo,omitempty"`
	// 卡号 （众邦必传）
	CardNo string `json:"cardNo,omitempty"`
	// 银行预留手机号（众邦必传）
	BankPhone string `json:"bankPhone,omitempty"`
	// 面部高清照，格式要求：Base64值；Base64编码后的大小不超2M
	ImgFace string `json:"imgFace"`
	// (暂未上线)用于活体检测的视频，视频的Base64值；Base64编码后的大小不超5M，支持mp4、avi、flv格式。
	// 请使用标准的Base64编码方式(带=补位)，编码规范参考RFC4648
	VideoBase64 string `json:"videoBase64,omitempty"`
	// 证件有效期开始日期 （众邦必传） 格式：yyyyMMdd 20000101
	ValidDateStart string `json:"validDateStart,omitempty"`
	// 证件有效期截止日期 （众邦必传） 格式：yyyyMMdd 20180101 长期上送”长期”中文字符
	ValidDateEnd string `json:"validDateEnd,omitempty"`
	// 住址 （众邦必传）
	Address string `json:"address,omitempty"`
	// IP 地址 （众邦必传）
	IpAdr string `json:"ipAdr,omitempty"`
}

type UserAccountLiveRequest struct {
	// 会员ID
	MemberId string `json:"memberId,omitempty"`
	// 支付通道（1:众邦）
	PayPass string `json:"payPass"`
	// 开户行编号 （众邦必传）
	BankNo string `json:"bankNo,omitempty"`
	// 卡号 （众邦必传）
	CardNo string `json:"cardNo,omitempty"`
	// 商户编号
	// 该值无需赋值，将由配置覆盖
	MerchantNo string `json:"merchantNo"`
	// 银行预留手机号（众邦必传）
	BankPhone string `json:"bankPhone,omitempty"`
	// 银行卡照片 格式要求：Base64值；Base64编码后的大小不超2M（众邦必传）
	ImgBank string `json:"imgBank,omitempty"`
	// 面部高清照，格式要求：Base64值；Base64编码后的大小不超2M
	ImgFace string `json:"imgFace"`
	// 用于活体检测的视频,要求规范:需居中露出完整的人脸,3~5秒即可,眨眼张嘴均可;视频的Base64值；
	// Base64编码后的大小不超5M，支持mp4、avi、flv格式。请使用标准的Base64编码方式(带=补位)，
	// 编码规范参考RFC4648
	VideoBase64 string `json:"videoBase64,omitempty"`
	// 证件有效期开始日期 （众邦必传） 格式：yyyyMMdd 20000101
	ValidDateStart string `json:"validDateStart,omitempty"`
	// 证件有效期截止日期 （众邦必传） 格式：yyyyMMdd 20180101 长期上送”长期”中文字符
	ValidDateEnd string `json:"validDateEnd,omitempty"`
	// 住址 （众邦必传）
	Address string `json:"address,omitempty"`
	// IP 地址 （众邦必传）
	IpAdr string `json:"ipAdr,omitempty"`
}

type UserAccountResponse struct {
	BizCommonResponse
}
