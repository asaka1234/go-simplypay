package go_simplypay

import (
	"errors"
	"github.com/asaka1234/go-simplypay/utils"
	"github.com/mitchellh/mapstructure"
)

// 充值的回调处理(传入一个处理函数)
func (cli *Client) DepositCallback(req SimplyPayDepositBackReq, processor func(SimplyPayDepositBackReq) error) error {
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
func (cli *Client) WithdrawCallBack(req SimplyPayWithdrawBackReq, processor func(SimplyPayWithdrawBackReq) error) error {
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
