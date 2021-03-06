// Copyright 2018 JDCLOUD.COM
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// NOTE: This class is auto generated by the jdcloud code generator program.

package apis

import (
    "github.com/jdcloud-api/jdcloud-sdk-go/core"
)

type SetHttpHeaderRequest struct {

    core.JDCloudRequest

    /* 用户域名  */
    Domain string `json:"domain"`

    /* header类型[resp,req] (Optional) */
    HeaderType *string `json:"headerType"`

    /* header名 (Optional) */
    HeaderName *string `json:"headerName"`

    /* header值 (Optional) */
    HeaderValue *string `json:"headerValue"`
}

/*
 * param domain: 用户域名 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewSetHttpHeaderRequest(
    domain string,
) *SetHttpHeaderRequest {

	return &SetHttpHeaderRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/domain/{domain}/httpHeader",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        Domain: domain,
	}
}

/*
 * param domain: 用户域名 (Required)
 * param headerType: header类型[resp,req] (Optional)
 * param headerName: header名 (Optional)
 * param headerValue: header值 (Optional)
 */
func NewSetHttpHeaderRequestWithAllParams(
    domain string,
    headerType *string,
    headerName *string,
    headerValue *string,
) *SetHttpHeaderRequest {

    return &SetHttpHeaderRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/domain/{domain}/httpHeader",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        Domain: domain,
        HeaderType: headerType,
        HeaderName: headerName,
        HeaderValue: headerValue,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewSetHttpHeaderRequestWithoutParam() *SetHttpHeaderRequest {

    return &SetHttpHeaderRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/domain/{domain}/httpHeader",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param domain: 用户域名(Required) */
func (r *SetHttpHeaderRequest) SetDomain(domain string) {
    r.Domain = domain
}

/* param headerType: header类型[resp,req](Optional) */
func (r *SetHttpHeaderRequest) SetHeaderType(headerType string) {
    r.HeaderType = &headerType
}

/* param headerName: header名(Optional) */
func (r *SetHttpHeaderRequest) SetHeaderName(headerName string) {
    r.HeaderName = &headerName
}

/* param headerValue: header值(Optional) */
func (r *SetHttpHeaderRequest) SetHeaderValue(headerValue string) {
    r.HeaderValue = &headerValue
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r SetHttpHeaderRequest) GetRegionId() string {
    return ""
}

type SetHttpHeaderResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result SetHttpHeaderResult `json:"result"`
}

type SetHttpHeaderResult struct {
}