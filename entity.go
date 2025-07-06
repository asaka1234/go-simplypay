package go_starpago

type StarPagoInitParams struct {
	MerchantId int64  `json:"merchantId" mapstructure:"merchantId" config:"merchantId"  yaml:"merchantId"` // merchantId
	AccessKey  string `json:"accessKey" mapstructure:"accessKey" config:"accessKey"  yaml:"accessKey"`     //接入秘钥
	Ip         string `json:"ip" mapstructure:"ip" config:"ip"  yaml:"ip"`                                 //回调时,对方的ip白名单

	DepositUrl      string `json:"depositUrl" mapstructure:"depositUrl" config:"depositUrl"  yaml:"depositUrl"`
	WithdrawUrl     string `json:"withdrawUrl" mapstructure:"withdrawUrl" config:"withdrawUrl"  yaml:"withdrawUrl"`
	DepositBackUrl  string `json:"depositBackUrl" mapstructure:"depositBackUrl" config:"depositBackUrl"  yaml:"depositBackUrl"`
	WithdrawBackUrl string `json:"WithdrawBackUrl" mapstructure:"WithdrawBackUrl" config:"WithdrawBackUrl"  yaml:"WithdrawBackUrl"`
}

type CommonResp struct {
	Code  int64  `json:"code" mapstructure:"code"`   // 响应状态码, 0为成功
	Error string `json:"error" mapstructure:"error"` // 错误信息
	Msg   string `json:"msg" mapstructure:"msg"`     // 返回文字描述
}

// ----------pre order-------------------------

type StarPagoDepositReq struct {
	Amount     string      `json:"amount" mapstructure:"amount"`           // 金额
	Attach     *string     `json:"attach,omitempty" mapstructure:"attach"` // 附加信息, 商户附加信息，原样返回
	Currency   string      `json:"currency" mapstructure:"currency"`       // 金额币种, PKR
	MerOrderNo string      `json:"merOrderNo" mapstructure:"merOrderNo"`   // 商户订单号, 商户订单号
	Extra      interface{} `json:"extra" mapstructure:"extra"`
	PayMethod  string      `json:"payMethod"` //代收方式
	// 以下sdk帮搞
	// AppID string `json:"appId"` // 应用号, appID
	// NotifyURL  string                  `json:"notifyUrl"`           // 异步通知地址, 异步通知地址
	// Sign string `json:"sign"` // 签名
}

// 巴基斯坦 extra
type StarPagoPKDepositReqExtra struct {
	PayType  string  `json:"payType" mapstructure:"payType"`             //  代收钱包选择: PK_WALLET_JAZZCASH / PK_WALLET_EASYPAISA 其中一个
	AutoFill *string `json:"autoFill,omitempty" mapstructure:"autoFill"` // 收银台页面是否自动填充提交上来的用户信息(1: 是, 0: 否) 默认值为 0
	Direct   *string `json:"direct,omitempty" mapstructure:"direct"`     //// 支持用户直连钱包，无需跳转到其他三方收银台(1: 直连, 0: 不直连) 默认值为0

	AccountNo *string `json:"accountNo,omitempty" mapstructure:"accountNo"` // 巴基斯坦的用户cnic: payType 为 PK_WALLET_JAZZCASH 时该项必填
	Mobile    *string `json:"mobile" mapstructure:"mobile"`                 // 03开头的11位数字（当autoFill="1"或者direct="1"的时候需要传真实的电话号码，否则用户钱包可能收不到支付信息）
}

// 菲律宾 extra
type StarPagoPHDepositReqExtra struct {
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	AccountName string `json:"accountName"`
	ChannelCode string `json:"channelCode"`
}

// 返回数据
type StarPagoDepositResponse struct {
	CommonResp `json:",inline" mapstructure:",squash"`
	Data       struct {
		Amount     int64  `json:"amount"` // 金额
		Attach     string `json:"attach"`
		CreateTime int64  `json:"createTime"`
		Currency   string `json:"currency"`   // 金额币种, 币种
		MerOrderNo string `json:"merOrderNo"` // 商户订单号, 商户订单号
		Message    string `json:"message"`    // 描述信息, 订单异常时会返回此字段，用作异常描述
		OrderNo    string `json:"orderNo"`    // 交易订单号, 订单号
		// 订单状态, 支付中 0 1 -4
		// 支付成功 2 3
		// 支付失败 -1 -2 -3
		OrderStatus int64                     `json:"orderStatus"`
		Params      StarPagoDepositRespParams `json:"params"` // 支付数据
		Sign        string                    `json:"sign"`   // 签名
		UpdateTime  int64                     `json:"updateTime"`
	} `json:"data"` // 金额
}

// 支付数据
type StarPagoDepositRespParams struct {
	BankCode    string `json:"bankCode"`
	Content     string `json:"content"`     // 支付内容 VA 为数字编号，QRIS 为二维码内容
	PaymentLink string `json:"paymentLink"` // 收银台地址
	PayMethod   string `json:"payMethod"`
	URL         string `json:"url"`
}

// ------------------------------------------------------------

type StarPagoDepositBackReq struct {
	Amount      int    `json:"amount" mapstructure:"amount"`           //订单金额
	Attach      string `json:"attach" mapstructure:"attach"`           //附加信息（商户附加信息，原样返回）
	OrderStatus int    `json:"orderStatus" mapstructure:"orderStatus"` //订单状态支付中 0 1 -4, 支付成功 2 3, 支付失败 -1 -2 -3
	OrderNo     string `json:"orderNo" mapstructure:"orderNo"`         //交易订单号
	MerOrderNo  string `json:"merOrderNo" mapstructure:"merOrderNo"`   //商户订单号
	Currency    string `json:"currency" mapstructure:"currency"`       //金额币种
	Message     string `json:"message" mapstructure:"message"`         //描述信息（订单异常时会返回此字段，用作异常描述）
	CreateTime  int64  `json:"createTime" mapstructure:"createTime"`   //毫秒, 1663236430613
	UpdateTime  int64  `json:"updateTime" mapstructure:"updateTime"`   //毫秒, 1663236430613
	Sign        string `json:"sign" mapstructure:"sign"`
}

// response 是 succeess / ok

//===========withdraw===================================

type StarPagoWithdrawReq struct {
	Amount     string      `json:"amount" mapstructure:"amount"`         // 金额
	Attach     *string     `json:"attach" mapstructure:"attach"`         // 附加信息, 商户附加信息，原样返回
	Currency   string      `json:"currency" mapstructure:"currency"`     // 金额币种, PKR
	Extra      interface{} `json:"extra" mapstructure:"extra"`           // 扩展信息 (代付给谁)
	MerOrderNo string      `json:"merOrderNo" mapstructure:"merOrderNo"` // 商户订单号, 商户订单号
	PayMethod  string      `json:"payMethod"`                            // 代付方式
	//以下sdk设置
	//AppID string `json:"appId"` // 应用号, appID
	//NotifyURL string `json:"notifyUrl"` // 异步通知地址, 异步通知地址
	//Sign  string `json:"sign"`  // 签名
}

// 巴基斯坦 extra
type StarPagoPKWithdrawReqExtra struct {
	BankCode string `json:"bankCode" mapstructure:"bankCode"` // 参考代付银行对照表

	// 持卡人姓名, 收款人姓名(尽量真实)
	//AccountName *string `json:"accountName,omitempty" mapstructure:"accountName"`
	// 收款银行账号(银行卡代付校验真实性，否则可能出款到错误的用户身上)
	AccountNo string `json:"accountNo" mapstructure:"accountNo"`

	// 巴基斯坦必填，身份证号13位纯数字（EP不交易真实,JZ校验） 尽量真实
	Cnic string `json:"cnic" mapstructure:"cnic"`

	// 收款人邮箱（格式正确就行）
	//Email *string `json:"email,omitempty" mapstructure:"email"`
	// 收款人电话 (校验真实性，03开头11位或者3开头10位数)钱包代付时必须真实，否则可能出现出款到错误的用户钱包去
	Mobile string `json:"mobile" mapstructure:"mobile"`
}

// 菲律宾 extra
type StarPagoPHWithdrawReqExtra struct {
	Email string `json:"email"`
	Phone string `json:"phone"` //或者是 mobile
	//BankCode    string `json:"bankCode"`
	AccountNo   string `json:"accountNo"`
	AccountName string `json:"accountName"`
}

// 返回数据
type StarPagoWithdrawResponse struct {
	CommonResp `json:",inline" mapstructure:",squash"`

	Data struct {
		Amount     int64  `json:"amount"`     // 订单金额
		Currency   string `json:"currency"`   // 金额币种
		MerOrderNo string `json:"merOrderNo"` // 商户订单号
		Message    string `json:"message"`    //// 描述信息, 订单异常时会返回此字段，用作异常描述
		OrderNo    string `json:"orderNo"`    // 交易订单号
		// 订单状态, 支付中 0 1 -4
		// 支付成功 2 3
		// 支付失败 -1 -2 -3
		OrderStatus int64  `json:"orderStatus"`
		Sign        string `json:"sign"` // 签名
		CreateTime  int64  `json:"createTime"`
		UpdateTime  int64  `json:"updateTime"`
	} `json:"data"` // 订单金额
}

type StarPagoWithdrawBackReq struct {
	OrderStatus int    `json:"orderStatus" mapstructure:"orderStatus"` //订单状态（订单状态支付中 0 1 -4, 支付成功 2 3, 支付失败 -1 -2 -3
	Attach      string `json:"attach" mapstructure:"attach"`           //附加信息（商户附加信息，原样返回）
	OrderNo     string `json:"orderNo" mapstructure:"orderNo"`         //交易订单号
	MerOrderNo  string `json:"merOrderNo" mapstructure:"merOrderNo"`   //商户订单号
	Amount      int    `json:"amount" mapstructure:"amount"`           //订单金额
	Currency    string `json:"currency" mapstructure:"currency"`       //金额币种
	Message     string `json:"message" mapstructure:"message"`         //描述信息（订单异常时会返回此字段，用作异常描述）
	CreateTime  int64  `json:"createTime" mapstructure:"createTime"`   //毫秒, 1663236430613
	UpdateTime  int64  `json:"updateTime" mapstructure:"updateTime"`   //毫秒, 1663236430613
	Sign        string `json:"sign" mapstructure:"sign"`
}
