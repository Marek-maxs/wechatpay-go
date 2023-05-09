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

type UserServicesApiService services.Service

// QueryUserService 查询用户服务可用信息
//
// 在商户生成乘车码前，商户请求查询用户服务可用信息接口，查询 用户服务可用信息，通过用户服务可用信息中的服务可用状态，来判断是否可以正常使用公共出行代扣服务
func (a *UserServicesApiService) QueryUserService(ctx context.Context, req QueryUserServiceRequest) (resp *UserServiceEntity, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodGet
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	// Make sure Path Params are properly set
	if req.ContractId == nil {
		return nil, nil, fmt.Errorf("field `ContractId` is required and must be specified in QueryUserServiceRequest")
	}

	localVarPath := consts.WechatPayAPIServer + "/v3/qrcode/user-services/contract-id/{contract_id}"
	// Build Path with Path Params
	localVarPath = strings.Replace(localVarPath, "{"+"contract_id"+"}", neturl.PathEscape(core.ParameterToString(*req.ContractId, "")), -1)

	// Make sure All Required Params are properly set
	if req.Appid == nil {
		return nil, nil, fmt.Errorf("field `Appid` is required and must be specified in QueryUserServiceRequest")
	}

	// Setup Query Params
	localVarQueryParams = neturl.Values{}
	localVarQueryParams.Add("appid", core.ParameterToString(*req.Appid, ""))
	if req.SubAppid != nil {
		localVarQueryParams.Add("sub_appid", core.ParameterToString(*req.SubAppid, ""))
	}
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

	// Extract UserServiceEntity from Http Response
	resp = new(UserServiceEntity)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}
