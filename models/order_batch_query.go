package models

type OrderBatchQueryRequest struct {
	// 是否自定义流水号 【0.否1.是】
	IsCustom string `json:"isCustom"`
	// 批次流水号 isCustom=0时必传
	BatchOrderId string `json:"batchOrderId"`
	// 自定义批次号 isCustom=1时必传
	BatchOrderNo string `json:"batchOrderNo"`
}

type OrderBatchPayListItem struct {
	// 会员ID
	MemberId string `json:"memberId,omitempty"`
	// 系统会员ID
	SystemId string `json:"systemId,omitempty"`
	// 会员名称
	Name string `json:"name,omitempty"`
	// 身份证号
	IdCard string `json:"idCard,omitempty"`
	// 金额
	Fee string `json:"fee,omitempty"`
	// 交易流水号
	OrderNo string `json:"orderNo,omitempty"`
	// 支付状态[0:待支付 1:支付中 2:支付成功 3:支付失败4:已撤回]
	PayStatus string `json:"payStatus,omitempty"`
	// 支付时间 格式：yyyyMMddHHmmss
	PayTime string `json:"payTime,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty"`
}

type OrderBatchQueryResponse struct {
	BizCommonResponse
	// 批次流水号
	BatchOrderId string `json:"batchOrderId,omitempty"`
	// 自定义流水号, 自定义时有值
	BatchOrderNo string `json:"batchOrderNo,omitempty"`
	// 商户号
	MerchantNo string `json:"merchantNo,omitempty"`
	// 服务商号
	ProviderNo string `json:"providerNo,omitempty"`
	// 任务编号
	TaskCode string `json:"taskCode,omitempty"`
	// 产品编号
	ProductCode string `json:"productCode,omitempty"`
	// 手续费承担方 1:企业承担 2:客户承担 查询类型1返回
	CostUndertaker string `json:"costUndertaker,omitempty"`
	// 业务类型
	BusType string `json:"busType,omitempty"`
	// 支付通道
	PayPass string `json:"payPass,omitempty"`
	// 批次状态[0:已提交1:已接单2:已拒单3:发放中4:发放成功5:已撤销6:发放失败 7:发放部分成功]
	BatchStatus string `json:"batchStatus,omitempty"`
	// 支付信息
	PayList []OrderBatchPayListItem `json:"payList,omitempty"`
}
