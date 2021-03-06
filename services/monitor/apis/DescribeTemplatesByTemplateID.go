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
    monitor "github.com/jdcloud-api/jdcloud-sdk-go/services/monitor/models"
)

type DescribeTemplatesByTemplateIDRequest struct {

    core.JDCloudRequest

    /* 模板 id  */
    TemplateId string `json:"templateId"`

    /* 模板类型，区分默认模板和用户自定义模板：1表示默认模板，2表示用户自定义模板
in: query (Optional) */
    TemplateType *int64 `json:"templateType"`
}

/*
 * param templateId: 模板 id (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeTemplatesByTemplateIDRequest(
    templateId string,
) *DescribeTemplatesByTemplateIDRequest {

	return &DescribeTemplatesByTemplateIDRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/template/{templateId}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        TemplateId: templateId,
	}
}

/*
 * param templateId: 模板 id (Required)
 * param templateType: 模板类型，区分默认模板和用户自定义模板：1表示默认模板，2表示用户自定义模板
in: query (Optional)
 */
func NewDescribeTemplatesByTemplateIDRequestWithAllParams(
    templateId string,
    templateType *int64,
) *DescribeTemplatesByTemplateIDRequest {

    return &DescribeTemplatesByTemplateIDRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/template/{templateId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        TemplateId: templateId,
        TemplateType: templateType,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeTemplatesByTemplateIDRequestWithoutParam() *DescribeTemplatesByTemplateIDRequest {

    return &DescribeTemplatesByTemplateIDRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/template/{templateId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param templateId: 模板 id(Required) */
func (r *DescribeTemplatesByTemplateIDRequest) SetTemplateId(templateId string) {
    r.TemplateId = templateId
}

/* param templateType: 模板类型，区分默认模板和用户自定义模板：1表示默认模板，2表示用户自定义模板
in: query(Optional) */
func (r *DescribeTemplatesByTemplateIDRequest) SetTemplateType(templateType int64) {
    r.TemplateType = &templateType
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeTemplatesByTemplateIDRequest) GetRegionId() string {
    return ""
}

type DescribeTemplatesByTemplateIDResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeTemplatesByTemplateIDResult `json:"result"`
}

type DescribeTemplatesByTemplateIDResult struct {
    Template monitor.TemplateVo `json:"template"`
}