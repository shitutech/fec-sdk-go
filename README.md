# 介绍

灵工 API Golang SDK

# 安装

```shell
go get github.com/shitutech/fec-sdk-go
```

```shell
## 更新
go get -u github.com/shitutech/fec-sdk-go
```

# 调用

```go
package main

import (
	"fmt"
	"github.com/shitutech/fec-sdk-go/fec"
	"github.com/shitutech/fec-sdk-go/models"
)

func main() {
	c := fec.NewConfig()
	c.SetMerchantNo("A*****-******-******").
		SetProviderNo("S*****").
		SetProductNo("P**********").
		SetTaskCode("STC**********").
		SetPrivateKey("privateKey-single-line")

	_client := fec.NewClient(c)
	acctInfoReq := &models.AcctInfoRequest{PayPass: "2"}

	accInfoResp, err := _client.AcctInfo(acctInfoReq)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(accInfoResp.Balance)
}
```

# 接口

| API             | Client 请求方法                | 请求                                  | 响应                           |
|-----------------|----------------------------|-------------------------------------|------------------------------|
| 用户注册            | ``UserRegister()``         | ``UserRegisterRequest``             | ``UserRegisterResponse``     |
| 用户信息查询          | ``UserQuery()``            | ``UserQueryRequest``                | ``UserQueryResponse``        |
| 用户信息变更 - 注册手机号  | ``UserUpdateMobile()``     | ``UserUpdateMobileRequest``         | ``UserUpdateResponse``       |
| 用户信息变更 - 影像件    | ``UserIdCardImage()``      | ``UserUpdateIdCardImageRequest``    | ``UserUpdateResponse``       |
| 用户信息变更 - 拓展业务类型 | ``UserIdBizType()``        | ``UserUpdateBizTypeRequest``        | ``UserUpdateResponse``       |
| 用户信息变更 - 结算卡信息  | ``UserIdSettlementCard()`` | ``UserUpdateSettlementCardRequest`` | ``UserUpdateResponse``       |
| 用户账户开户          |                            | ``UserOpenRequest``                 | ``UserOpenResponse``         |
| 用户账户开户（活体认证）    |                            | ``UserOpenVideoRequest``            | ``UserOpenResponse``         |
| 订单支付            |                            | ``OrderPayRequest``                 | ``OrderPayResponse``         |
| 批次订单号查询订单       |                            | ``OrderQueryBatchRequest``          | ``OrderQueryBatchResponse``  |
| 子订单详情查询         |                            | ``OrderQueryDetailRequest``         | ``OrderQueryDetailResponse`` |
| 商户信息查询          | ``AcctInfo()``             | ``AcctInfoRequest``                 | ``AcctInfoResponse``         |