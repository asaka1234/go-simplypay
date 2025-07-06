package go_simplypay

import (
	"crypto/tls"
	"github.com/asaka1234/go-simplypay/utils"
	"github.com/mitchellh/mapstructure"
)

// 代收
func (cli *Client) Deposit(req SimplyPayDepositReq) (*SimplyPayDepositResponse, error) {

	rawURL := cli.Params.DepositUrl

	var params map[string]interface{}
	mapstructure.Decode(req, &params)
	params["appId"] = cli.Params.MerchantId
	params["notifyUrl"] = cli.Params.DepositBackUrl
	params["currency"] = "INR"

	//签名
	signStr := utils.Sign(params, cli.Params.AccessKey)
	params["sign"] = signStr

	//返回值会放到这里
	var result SimplyPayDepositResponse

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

	return &result, nil
}
