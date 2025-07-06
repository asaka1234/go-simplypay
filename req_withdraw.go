package go_simplypay

import (
	"crypto/tls"
	"github.com/asaka1234/go-simplypay/utils"
	"github.com/mitchellh/mapstructure"
)

// withdraw
func (cli *Client) Withdraw(req SimplyPayWithdrawReq) (*SimplyPayWithdrawResponse, error) {

	rawURL := cli.Params.WithdrawUrl

	var params map[string]interface{}
	mapstructure.Decode(req, &params)
	params["appId"] = cli.Params.MerchantId
	params["notifyUrl"] = cli.Params.WithdrawBackUrl
	params["currency"] = "INR" //印度

	//签名
	signStr := utils.Sign(params, cli.Params.AccessKey)
	params["sign"] = signStr

	//返回值会放到这里
	var result SimplyPayWithdrawResponse

	_, err := cli.ryClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetCloseConnection(true).
		R().
		SetHeaders(getHeaders()).
		SetBody(params).
		SetDebug(cli.debugMode).
		SetResult(&result).
		Post(rawURL)

	if err != nil {
		return nil, err
	}

	return &result, err
}
