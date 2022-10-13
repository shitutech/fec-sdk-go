package models

type OrderPayListItem struct {
	// 系统会员ID
	MemberId string `json:"memberId"`
	// 会员名称
	Name string `json:"name"`
	// 身份证号
	IdCard string `json:"idCard"`
	// 开户行编号 （isDefaultCard=0必传）
	BankNo string `json:"bankNo,omitempty"`
	// 开户行名称 （isDefaultCard=0必传）
	BankName string `json:"bankName,omitempty"`
	// 结算卡号 （isDefaultCard=0必传）
	CardNo string `json:"cardNo,omitempty"`
	// 结算卡号银行预留手机号 （isDefaultCard=0必传）
	BankPhone string `json:"bankPhone,omitempty"`
	// 支付宝号 支付通道2 必传
	AliPayNo string `json:"aliPayNo,omitempty"`
	// 金额 单位元，精确到两位小数点 示例： 10.05元
	Fee string `json:"fee"`
}

type OrderPayRequest struct {
	// 服务商号
	// 无需设置，将由配置覆盖
	ProviderNo string `json:"providerNo"`
	// 账户类型【1.商户开户2.服务商开户】 对应账户和服务商关系
	AccountType string `json:"accountType"`
	// 任务编号
	// 无需设置，将由配置覆盖
	TaskCode string `json:"taskCode"`
	// 产品编号
	// 无需设置，将由配置覆盖
	ProductCode string `json:"productCode"`
	// 手续费承担方 1:企业承担 2:客户承担
	CostUndertaker string `json:"costUndertaker"`
	// 支付通道（1:众邦 2:支付宝 4.招商银行）
	PayPass string `json:"payPass"`
	// 是否自定义流水号 【0.不启用1.启用】
	IsCustom string `json:"isCustom"`
	// 批次号 格式：20200101-00001 当天日期+五位序列号 （isCustom=0时不为空）
	BatchNumber string `json:"batchNumber,omitempty"`
	// 不超过20位 自定义流水号 （isCustom=1时不为空）
	BatchOrderNo string `json:"batchOrderNo,omitempty"`
	// 是否使用默认绑定卡 （0 否1 是 ）【支付宝到户是传1】
	IsDefaultCard string `json:"isDefaultCard"`
	// 支付信息 数据格式[{“memberId”:”134814782XXX”,”name”:”李XX”,”idCard”:”2305231XXX”,”fee”:”10.00”}]
	PayList []OrderPayListItem `json:"payList"`
}

type OrderRespPayListElement struct {
	// 系统会员ID
	MemberId string `json:"memberId"`
	// 会员名称
	Name string `json:"name"`
	// 身份证号
	IdCard string `json:"idCard"`
	// 开户行编号
	BankNo string `json:"bankNo,omitempty"`
	// 开户行名称
	BankName string `json:"bankName,omitempty"`
	// 结算卡号
	CardNo string `json:"cardNo,omitempty"`
	// 结算卡号银行预留手机号
	BankPhone string `json:"bankPhone,omitempty"`
	// 金额
	Fee string `json:"fee"`
	// 交易流水号
	OrderNo string `json:"orderNo"`
	// 支付状态[0:待支付 1:支付中 2:支付成功 3:支付失败4:已撤回]
	PayStatus string `json:"payStatus"`
}

type OrderRespFailListElement struct {
	// 系统会员ID
	MemberId string `json:"memberId"`
	// 会员名称
	Name string `json:"name"`
	// 身份证号
	IdCard string `json:"idCard"`
	// 开户行编号
	BankNo string `json:"bankNo,omitempty"`
	// 开户行名称
	BankName string `json:"bankName,omitempty"`
	// 结算卡号
	CardNo string `json:"cardNo,omitempty"`
	// 结算卡号银行预留手机号
	BankPhone string `json:"bankPhone,omitempty"`
	// 支付宝号 支付通道2返回
	AliPayNo string `json:"aliPayNo,omitempty"`
	// 金额
	Fee string `json:"fee"`
	// 备注
	Remark string `json:"remark"`
}

type OrderPayResponse struct {
	BizCommonResponse
	// 系统批次流水号
	BatchOrderId string `json:"batchOrderId,omitempty"`
	// 自定义流水号
	BatchOrderNo string `json:"batchOrderNo,omitempty"`
	// 支付信息 （code=1001返回以下信息）
	PayList []OrderRespPayListElement `json:"payList,omitempty"`
	// 失败信息 （code=1002返回以下信息）
	FailList []OrderRespFailListElement `json:"failList,omitempty"`
}
