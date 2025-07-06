package go_starpago

import (
	"errors"
	"github.com/asaka1234/go-starpago/utils"
	"github.com/mitchellh/mapstructure"
)

// 充值的回调处理(传入一个处理函数)
func (cli *Client) DepositCallback(req StarPagoDepositBackReq, processor func(StarPagoDepositBackReq) error) error {
	//验证签名
	var params map[string]interface{}
	mapstructure.Decode(req, &params)

	verifyResult := utils.VerifySign(params, cli.Params.AccessKey)
	if !verifyResult {
		//验签失败
		return errors.New("verify sign error!")
	}

	//开始处理
	return processor(req)
}

//==========================================

// 充值的回调处理(传入一个处理函数)
func (cli *Client) WithdrawCallBack(req StarPagoWithdrawBackReq, processor func(StarPagoWithdrawBackReq) error) error {
	//验证签名
	var params map[string]interface{}
	mapstructure.Decode(req, &params)

	verifyResult := utils.VerifySign(params, cli.Params.AccessKey)
	if !verifyResult {
		//验签失败
		return errors.New("verify sign error!")
	}

	//开始处理
	return processor(req)
}
