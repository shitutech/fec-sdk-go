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

| API             | 请求类                            | 响应类                             |
|-----------------|--------------------------------|---------------------------------|
| 用户注册            | UserRegisterRequest::class     | UserRegisterResponse::class     |
| 用户信息查询          | UserQueryRequest::class        | UserQueryResponse::class        |
| 用户信息变更 - 注册手机号  | UserUpdatePhoneRequest::class  | UserUpdateResponse::class       |
| 用户信息变更 - 影像件    | UserUpdateImageRequest::class  | UserUpdateResponse::class       |
| 用户信息变更 - 拓展业务类型 | UserUpdateExpandRequest::class | UserUpdateResponse::class       |
| 用户信息变更 - 结算卡信息  | UserUpdateCardRequest::class   | UserUpdateResponse::class       |
| 用户账户开户          | UserOpenRequest::class         | UserOpenResponse::class         |
| 用户账户开户（活体认证）    | UserOpenVideoRequest::class    | UserOpenResponse::class         |
| 订单支付            | OrderPayRequest::class         | OrderPayResponse::class         |
| 批次订单号查询订单       | OrderQueryBatchRequest::class  | OrderQueryBatchResponse::class  |
| 子订单详情查询         | OrderQueryDetailRequest::class | OrderQueryDetailResponse::class |
| 商户信息查询          | AcctInfoRequest                | AcctInfoResponse                |