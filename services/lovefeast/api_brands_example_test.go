// Copyright 2021 Tencent Inc. All rights reserved.
//
// 爱心餐对外API
//
// 微信支付爱心餐公益计划旨在面向深圳市的市政一线环卫工人提供每周一餐的1分钱用餐公益服务。在受助端，微信支付联动上千家餐饮门店关爱特殊人群，通过微信支付数字化能力将人群身份认证与公益福利领用全流程线上化，实现公益福利精准到人。在捐赠端，微信支付发挥连接优势与平台能力，结合用户就餐场景通过爱心餐一块捐插件让用户可在点餐时顺手捐1元，带动更多社会力量致谢城市美容师。
//
// API version: 0.0.4

// Code generated by WechatPay APIv3 Generator based on [OpenAPI Generator](https://openapi-generator.tech); DO NOT EDIT.

package lovefeast_test

import (
	"context"
	"log"

	"github.com/Marek-maxs/wechatpay-go/core"
	"github.com/Marek-maxs/wechatpay-go/core/option"
	"github.com/Marek-maxs/wechatpay-go/services/lovefeast"
	"github.com/Marek-maxs/wechatpay-go/utils"
)

func ExampleBrandsApiService_GetBrand() {
	var (
		mchID                      string = "190000****"                               // 商户号
		mchCertificateSerialNumber string = "3775************************************" // 商户证书序列号
		mchAPIv3Key                string = "2ab9****************************"         // 商户APIv3密钥
	)

	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath("/path/to/merchant/apiclient_key.pem")
	if err != nil {
		log.Printf("load merchant private key error:%s", err)
		return
	}

	ctx := context.Background()
	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(mchID, mchCertificateSerialNumber, mchPrivateKey, mchAPIv3Key),
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
		return
	}

	svc := lovefeast.BrandsApiService{Client: client}
	resp, result, err := svc.GetBrand(ctx,
		lovefeast.GetBrandRequest{
			BrandId: core.Int64(1442),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call GetBrand err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
	}
}
