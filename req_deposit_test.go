package go_simplypay

import (
	"fmt"
	"testing"
)

type VLog struct {
}

func (l VLog) Debugf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
func (l VLog) Infof(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
func (l VLog) Warnf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
func (l VLog) Errorf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}

// 印度 代收
func TestDeposit(t *testing.T) {
	vLog := VLog{}
	//构造client
	cli := NewClient(vLog, &SimplyPayInitParams{MERCHANT_ID, ACCESS_SECRET, IP, DEPOSIT_URL, WITHDRAW_URL, DEPOSIT_BACK_URL, WITHDRAW_BACK_URL})

	//发请求
	resp, err := cli.Deposit(GenDepositRequestDemo())
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
		return
	}
	fmt.Printf("resp:%+v\n", resp)
}

// 印度
func GenDepositRequestDemo() SimplyPayDepositReq {
	return SimplyPayDepositReq{
		MerOrderNo: "323231224", //商户id
		Amount:     "200",       //100->50000
		Extra: SimplyPayINRDepositReqExtra{
			Name:   "john", //这个不需要其他参数
			Email:  "demo@gmail.com",
			Mobile: "+4478781234",
		},
	}
}
