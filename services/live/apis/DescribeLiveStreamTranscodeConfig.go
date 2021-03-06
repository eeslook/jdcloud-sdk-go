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
    live "github.com/jdcloud-api/jdcloud-sdk-go/services/live/models"
    common "github.com/jdcloud-api/jdcloud-sdk-go/services/common/models"
)

type DescribeLiveStreamTranscodeConfigRequest struct {

    core.JDCloudRequest

    /* 页码, 默认为1, 取值范围：[1,∞) (Optional) */
    PageNum *int `json:"pageNum"`

    /* 分页大小，默认为20，取值范围：[10,100] (Optional) */
    PageSize *int `json:"pageSize"`

    /* 转码模板查询过滤条件, 不传递分页参数时默认返回10条 (Optional) */
    Filters []common.Filter `json:"filters"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeLiveStreamTranscodeConfigRequest(
) *DescribeLiveStreamTranscodeConfigRequest {

	return &DescribeLiveStreamTranscodeConfigRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/transcodes:config",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param pageNum: 页码, 默认为1, 取值范围：[1,∞) (Optional)
 * param pageSize: 分页大小，默认为20，取值范围：[10,100] (Optional)
 * param filters: 转码模板查询过滤条件, 不传递分页参数时默认返回10条 (Optional)
 */
func NewDescribeLiveStreamTranscodeConfigRequestWithAllParams(
    pageNum *int,
    pageSize *int,
    filters []common.Filter,
) *DescribeLiveStreamTranscodeConfigRequest {

    return &DescribeLiveStreamTranscodeConfigRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/transcodes:config",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        PageNum: pageNum,
        PageSize: pageSize,
        Filters: filters,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeLiveStreamTranscodeConfigRequestWithoutParam() *DescribeLiveStreamTranscodeConfigRequest {

    return &DescribeLiveStreamTranscodeConfigRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/transcodes:config",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param pageNum: 页码, 默认为1, 取值范围：[1,∞)(Optional) */
func (r *DescribeLiveStreamTranscodeConfigRequest) SetPageNum(pageNum int) {
    r.PageNum = &pageNum
}

/* param pageSize: 分页大小，默认为20，取值范围：[10,100](Optional) */
func (r *DescribeLiveStreamTranscodeConfigRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param filters: 转码模板查询过滤条件, 不传递分页参数时默认返回10条(Optional) */
func (r *DescribeLiveStreamTranscodeConfigRequest) SetFilters(filters []common.Filter) {
    r.Filters = filters
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeLiveStreamTranscodeConfigRequest) GetRegionId() string {
    return ""
}

type DescribeLiveStreamTranscodeConfigResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeLiveStreamTranscodeConfigResult `json:"result"`
}

type DescribeLiveStreamTranscodeConfigResult struct {
    PageNumber int `json:"pageNumber"`
    PageSize int `json:"pageSize"`
    TotalCount int `json:"totalCount"`
    TranscodeConfigs []live.TemplateConfig `json:"transcodeConfigs"`
}