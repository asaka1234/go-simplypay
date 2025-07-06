package go_simplypay

type SimplyPayInitParams struct {
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

type SimplyPayDepositReq struct {
	Amount     string                      `json:"amount" mapstructure:"amount"`           // 金额
	Attach     *string                     `json:"attach,omitempty" mapstructure:"attach"` // 附加信息, 商户附加信息，原样返回
	MerOrderNo string                      `json:"merOrderNo" mapstructure:"merOrderNo"`   // 商户订单号, 商户订单号
	Extra      SimplyPayINRDepositReqExtra `json:"extra" mapstructure:"extra"`
	// 以下sdk帮搞
	// AppID string `json:"appId"` // 应用号, appID
	// NotifyURL  string                  `json:"notifyUrl"`           // 异步通知地址, 异步通知地址
	// Currency   string                      `json:"currency" mapstructure:"currency"`       // 金额币种, INR
	// Sign string `json:"sign"` // 签名
}

type SimplyPayINRDepositReqExtra struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	Mobile string `json:"mobile"`
}

// 返回数据
type SimplyPayDepositResponse struct {
	CommonResp `json:",inline" mapstructure:",squash"`
	Data       struct {
		// 订单状态, 支付中 0 1 -4
		// 支付成功 2 3
		// 支付失败 -1 -2 -3
		OrderStatus int    `json:"orderStatus"`
		OrderNo     string `json:"orderNo"`    //交易订单号
		MerOrderNo  string `json:"merOrderNo"` //商户订单号
		Amount      int    `json:"amount"`
		Currency    string `json:"currency"`
		Attach      string `json:"attach"`
		Extra       struct {
			Name   string `json:"name"`
			Email  string `json:"email"`
			Mobile string `json:"mobile"`
		} `json:"extra"`
		CreateTime int64  `json:"createTime"`
		UpdateTime int64  `json:"updateTime"`
		Sign       string `json:"sign"`
		Params     struct {
			PaymentLink string `json:"paymentLink"` //收银台链接
		} `json:"params"`
	} `json:"data"`
}

// ------------------------------------------------------------

type SimplyPayDepositBackReq struct {
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

type SimplyPayWithdrawReq struct {
	Amount     string                       `json:"amount" mapstructure:"amount"`         // 金额
	Attach     *string                      `json:"attach" mapstructure:"attach"`         // 附加信息, 商户附加信息，原样返回
	Extra      SimplyPayINRWithdrawReqExtra `json:"extra" mapstructure:"extra"`           // 扩展信息 (代付给谁)
	MerOrderNo string                       `json:"merOrderNo" mapstructure:"merOrderNo"` // 商户订单号, 商户订单号
	//以下sdk设置
	//AppID string `json:"appId"` // 应用号, appID
	//NotifyURL string `json:"notifyUrl"` // 异步通知地址, 异步通知地址
	//Currency string `json:"currency" mapstructure:"currency"` // 金额币种, INR
	//Sign  string `json:"sign"`  // 签名
}

type SimplyPayINRWithdrawReqExtra struct {
	PayoutType string `json:"payoutType"` //fixed, IFSC
	Ifsc       string `json:"ifsc"`       //ifsc号码
	Account    string `json:"account"`    //银行收款账号
	Name       string `json:"name"`       //客户名称
	Email      string `json:"email"`      //客户邮箱
	Mobile     string `json:"mobile"`     //客户手机
}

// 返回数据
type SimplyPayWithdrawResponse struct {
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

type SimplyPayWithdrawBackReq struct {
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
