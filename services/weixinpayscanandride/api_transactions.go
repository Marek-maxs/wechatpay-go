// Copyright 2021 Tencent Inc. All rights reserved.
//
// 公共出行平台代扣服务对外API
//
// 公共出行平台代扣服务对外API
//
// API version: 1.0.0

// Code generated by WechatPay APIv3 Generator based on [OpenAPI Generator](https://openapi-generator.tech); DO NOT EDIT.

package weixinpayscanandride

import (
	"context"
	"fmt"
	nethttp "net/http"
	neturl "net/url"
	"strings"

	"github.com/Marek-maxs/wechatpay-go/core"
	"github.com/Marek-maxs/wechatpay-go/core/consts"
	"github.com/Marek-maxs/wechatpay-go/services"
)

type TransactionsApiService services.Service

// CreateTransaction 扣费受理
//
// 商户请求扣费受理接口，会完成订单受理。微信支付进行异步扣款，支付完成后，会将订单支付结果发送给商户。
func (a *TransactionsApiService) CreateTransaction(ctx context.Context, req CreateTransactionRequest) (resp *TransactionsEntity, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPost
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	localVarPath := consts.WechatPayAPIServer + "/v3/qrcode/transactions"
	// Make sure All Required Params are properly set

	// Setup Body Params
	localVarPostBody = req

	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{"application/json"}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)

	// Perform Http Request
	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	if err != nil {
		return nil, result, err
	}

	// Extract TransactionsEntity from Http Response
	resp = new(TransactionsEntity)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// QueryTransaction 查询订单
//
// 商户通过商户订单号，来查询订单信息
func (a *TransactionsApiService) QueryTransaction(ctx context.Context, req QueryTransactionRequest) (resp *TransactionsEntity, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodGet
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	// Make sure Path Params are properly set
	if req.OutTradeNo == nil {
		return nil, nil, fmt.Errorf("field `OutTradeNo` is required and must be specified in QueryTransactionRequest")
	}

	localVarPath := consts.WechatPayAPIServer + "/v3/qrcode/transactions/out-trade-no/{out_trade_no}"
	// Build Path with Path Params
	localVarPath = strings.Replace(localVarPath, "{"+"out_trade_no"+"}", neturl.PathEscape(core.ParameterToString(*req.OutTradeNo, "")), -1)

	// Make sure All Required Params are properly set

	// Setup Query Params
	localVarQueryParams = neturl.Values{}
	if req.SubMchid != nil {
		localVarQueryParams.Add("sub_mchid", core.ParameterToString(*req.SubMchid, ""))
	}

	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)

	// Perform Http Request
	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	if err != nil {
		return nil, result, err
	}

	// Extract TransactionsEntity from Http Response
	resp = new(TransactionsEntity)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}
