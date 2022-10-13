# 介绍

灵工 API Golang SDK

# 安装

```shell
go get github.com/shitutech/fec-sdk-go/v2
```

```shell
## 更新
go get -u github.com/shitutech/fec-sdk-go/v2
```

# 调用

```go
package main

import (
	"fmt"
	"github.com/shitutech/fec-sdk-go/v2/fec"
	"github.com/shitutech/fec-sdk-go/v2/models"
)

func main() {
	c := fec.NewConfig()
	c.SetMerchantNo("A*****-******-******").
		SetProviderNo("S*****").
		SetProductNo("P**********").
		SetTaskCode("STC**********").
		SetPrivateKey("privateKey-single-line")

	_client := fec.NewClient(c)
	acctInfoReq := &models.AcctInfoRequest{
		AccountType: fec.AccountTypeService,
		PayPass: fec.PayPassZhaoHang,
		ProviderNo: c.ProviderNo(),
	}

	accInfoResp, err := _client.AcctInfo(acctInfoReq)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(accInfoResp)
}
```

# 接口

| API            | Client 请求方法                | 请求                                  | 响应                          |
|----------------|----------------------------|-------------------------------------|-----------------------------|
| 用户注册           | ``UserRegister()``         | ``UserRegisterRequest``             | ``UserRegisterResponse``    |
| 用户信息查询         | ``UserQuery()``            | ``UserQueryRequest``                | ``UserQueryResponse``       |
| 查询会员系统         | ``UserSystemId()``         | ``UserSystemIdRequest``             | ``UserSystemIdResponse``    |
| 用户信息变更 - 结算卡信息 | ``UserIdSettlementCard()`` | ``UserUpdateSettlementCardRequest`` | ``UserUpdateResponse``      |
| 查询会员绑定的结算卡列表   | ``UserBindCards``          | ``UserBindCardsRequest``            | ``UserBindCardsResponse``   |
| 查询通道支持的银行列表    | ``SupportBankList``        | ``SupportBankListRequest``          | ``SupportBankListResponse`` |
| 订单支付           | ``OrderPay()``             | ``OrderPayRequest``                 | ``OrderPayResponse``        |
| 批次订单号查询订单      | ``OrderBatchQuery()``      | ``OrderBatchQueryRequest``          | ``OrderBatchQueryResponse`` |
| 子订单详情查询        | ``OrderSubQuery()``        | ``OrderSubQueryRequest``            | ``OrderSubQueryResponse``   |
| 商户账户信息查询       | ``AcctInfo()``             | ``AcctInfoRequest``                 | ``AcctInfoResponse``        |
| 商户账户列表查询       | ``AcctList()``             | ``AcctListRequest``                 | ``AcctListResponse``        |