package models

type OrderSubQueryRequest struct {
	// 是否自定义流水号 【0.否1.是】
	IsCustom string `json:"isCustom"`
	// 批次流水号
	OrderNo string `json:"orderNo"`
	// 自定义批次号 isCustom=1时必传
	BatchOrderNo string `json:"batchOrderNo,omitempty"`
}

type OrderSubQueryResponse struct {
	BizCommonResponse
	// 交易流水号
	OrderNo string `json:"orderNo,omitempty"`
	// 会员名称
	Name string `json:"name,omitempty"`
	// 身份证号
	IdCard string `json:"idCard,omitempty"`
	// 金额
	Fee string `json:"fee,omitempty"`
	// 支付状态[0:待支付 1:支付中 2:支付成功 3:支付失败4:已撤回]
	PayStatus string `json:"payStatus,omitempty"`
	// 支付时间 格式：yyyyMMddHHmmss
	PayTime string `json:"payTime,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty"`
}
