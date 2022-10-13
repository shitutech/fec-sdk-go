package fec

const (
	/*
		账户类型1商户开户2服务商开户
	*/
	AccountTypeMch     = "1"
	AccountTypeService = "2"

	/*
	 支付通道（1:众邦 2：支付宝 4.招商银行）
	*/
	PayPassZb       = "1"
	PayPassAlipay   = "2"
	PayPassZhaoHang = "4"

	/*
		业务类型 1委托代征2.个体户注册(分包)3.自然人代开4.临时税务登记
	*/
	BusTypeEntrusted      = "1"
	BusTypeSelfEmployed   = "2"
	BusTypeNatural        = "3"
	BusTypeProvisionalTax = "4"

	/*
		查询类型 1:查询基本信息 2:查询开户信息 3:查询会员ID
	*/
	UserQueryTypeBase     = "1"
	UserQueryTypeAccount  = "2"
	UserQueryTypeMemberId = "3"

	/*
		3. 根据身份证号4.根据商户会员Id查询系统会员Id
	*/
	UserSystemQueryTypeId  = "3"
	UserSystemQueryTypeMch = "4"

	/*
		账户状态 0:未开通 1:开通中 2:开通成功 3:开通失败 （查询类型 2 返回）
	*/
	AcctStatusNotActivated          = "1"
	AcctStatusActivated             = "2"
	AcctStatusActivatedSuccessfully = "3"
	AcctStatusActivatedFailed       = "4"

	/*
		变更类型 1:变更注册手机号 2:变更影像件 3:拓展业务类型 4：变更结算卡信息
	*/
	ChangeTypeMobile         = "1"
	ChangeTypeIdCardImage    = "2"
	ChangeTypeBizType        = "3"
	ChangeTypeSettlementCard = "4"

	/*
		手续费承担方 1:企业承担 2:客户承担
	*/
	CostUndertakerEnterprise = "1"
	CostUndertakerCustomer   = "2"

	/*
		批次状态[0:已提交1:已接单2:已拒单3:发放中4:发放成功5:已撤销6:发放失败 7:发放部分成功]
	*/
	BatchStatusSubmitted                         = "0"
	BatchStatusOrderReceived                     = "1"
	BatchStatusOrderRejected                     = "2"
	BatchStatusDistributing                      = "3"
	BatchStatusDistributingSuccessful            = "4"
	BatchStatusCancelled                         = "5"
	BatchStatusDistributingFailed                = "6"
	BatchStatusPartiallyDistributingSuccessfully = "7"

	/*
		支付状态[0:待支付 1:支付中 2:支付成功 3:支付失败4:已撤回]
	*/
	PayStatusPendingPayment    = "0"
	PayStatusPaymentInProgress = "1"
	PayStatusPaymentSuccessful = "2"
	PayStatusPaymentFailed     = "3"
	PayStatusWithdrawn         = "4"

	/*
		是否自定义流水号 【0.不启用1.启用】
	*/
	OrderIsCustomN = "0"
	OrderIsCustomY = "1"

	/*
		是否使用默认绑定卡 （0 否1 是 ）【支付宝到户是传1】
	*/
	OrderIsDefaultCardN = "0"
	OrderIsDefaultCardY = "1"
)
