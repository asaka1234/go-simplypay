package go_simplypay

import (
	"fmt"
	"testing"
)

// 印度 代付
func TestPKWithdraw(t *testing.T) {
	vLog := VLog{}
	//构造client
	cli := NewClient(vLog, &SimplyPayInitParams{MERCHANT_ID, ACCESS_SECRET, IP, DEPOSIT_URL, WITHDRAW_URL, DEPOSIT_BACK_URL, WITHDRAW_BACK_URL})

	//发请求
	resp, err := cli.Withdraw(GenWithdrawRequestDemo())
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
		return
	}
	fmt.Printf("resp:%+v\n", resp)
}

func GenWithdrawRequestDemo() SimplyPayWithdrawReq {
	return SimplyPayWithdrawReq{
		MerOrderNo: "111",
		Amount:     "200", //100-50000
		Extra: SimplyPayINRWithdrawReqExtra{
			PayoutType: "IFSC",
			Ifsc:       "129171971",
			Account:    "1891917917912",
			Name:       "john",
			Email:      "demo@gmail.com",
			Mobile:     "0312345678",
		},
	}
}
