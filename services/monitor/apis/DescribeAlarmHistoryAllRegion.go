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

type DescribeAlarmHistoryAllRegionRequest struct {

    core.JDCloudRequest

    /* 当前所在页，默认为1 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 页面大小，默认为20；取值范围[1, 100] (Optional) */
    PageSize *int `json:"pageSize"`

    /* 产品线 (Optional) */
    ServiceCode *string `json:"serviceCode"`

    /* 资源Id (Optional) */
    ResourceId *string `json:"resourceId"`

    /* 规则Id (Optional) */
    AlarmId *string `json:"alarmId"`

    /* 正在报警, 取值为1 (Optional) */
    Alarming *int `json:"alarming"`

    /* 开始时间 (Optional) */
    StartTime *string `json:"startTime"`

    /* 结束时间 (Optional) */
    EndTime *string `json:"endTime"`

    /* 服务码或资源Id列表
filter name 为serviceCodes表示查询多个产品线的规则
filter name 为resourceIds表示查询多个资源的规则 (Optional) */
    Filters []monitor.Filter `json:"filters"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeAlarmHistoryAllRegionRequest(
) *DescribeAlarmHistoryAllRegionRequest {

	return &DescribeAlarmHistoryAllRegionRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/rule/queryNotice",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param pageNumber: 当前所在页，默认为1 (Optional)
 * param pageSize: 页面大小，默认为20；取值范围[1, 100] (Optional)
 * param serviceCode: 产品线 (Optional)
 * param resourceId: 资源Id (Optional)
 * param alarmId: 规则Id (Optional)
 * param alarming: 正在报警, 取值为1 (Optional)
 * param startTime: 开始时间 (Optional)
 * param endTime: 结束时间 (Optional)
 * param filters: 服务码或资源Id列表
filter name 为serviceCodes表示查询多个产品线的规则
filter name 为resourceIds表示查询多个资源的规则 (Optional)
 */
func NewDescribeAlarmHistoryAllRegionRequestWithAllParams(
    pageNumber *int,
    pageSize *int,
    serviceCode *string,
    resourceId *string,
    alarmId *string,
    alarming *int,
    startTime *string,
    endTime *string,
    filters []monitor.Filter,
) *DescribeAlarmHistoryAllRegionRequest {

    return &DescribeAlarmHistoryAllRegionRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/rule/queryNotice",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        PageNumber: pageNumber,
        PageSize: pageSize,
        ServiceCode: serviceCode,
        ResourceId: resourceId,
        AlarmId: alarmId,
        Alarming: alarming,
        StartTime: startTime,
        EndTime: endTime,
        Filters: filters,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeAlarmHistoryAllRegionRequestWithoutParam() *DescribeAlarmHistoryAllRegionRequest {

    return &DescribeAlarmHistoryAllRegionRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/rule/queryNotice",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param pageNumber: 当前所在页，默认为1(Optional) */
func (r *DescribeAlarmHistoryAllRegionRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 页面大小，默认为20；取值范围[1, 100](Optional) */
func (r *DescribeAlarmHistoryAllRegionRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param serviceCode: 产品线(Optional) */
func (r *DescribeAlarmHistoryAllRegionRequest) SetServiceCode(serviceCode string) {
    r.ServiceCode = &serviceCode
}

/* param resourceId: 资源Id(Optional) */
func (r *DescribeAlarmHistoryAllRegionRequest) SetResourceId(resourceId string) {
    r.ResourceId = &resourceId
}

/* param alarmId: 规则Id(Optional) */
func (r *DescribeAlarmHistoryAllRegionRequest) SetAlarmId(alarmId string) {
    r.AlarmId = &alarmId
}

/* param alarming: 正在报警, 取值为1(Optional) */
func (r *DescribeAlarmHistoryAllRegionRequest) SetAlarming(alarming int) {
    r.Alarming = &alarming
}

/* param startTime: 开始时间(Optional) */
func (r *DescribeAlarmHistoryAllRegionRequest) SetStartTime(startTime string) {
    r.StartTime = &startTime
}

/* param endTime: 结束时间(Optional) */
func (r *DescribeAlarmHistoryAllRegionRequest) SetEndTime(endTime string) {
    r.EndTime = &endTime
}

/* param filters: 服务码或资源Id列表
filter name 为serviceCodes表示查询多个产品线的规则
filter name 为resourceIds表示查询多个资源的规则(Optional) */
func (r *DescribeAlarmHistoryAllRegionRequest) SetFilters(filters []monitor.Filter) {
    r.Filters = filters
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeAlarmHistoryAllRegionRequest) GetRegionId() string {
    return ""
}

type DescribeAlarmHistoryAllRegionResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeAlarmHistoryAllRegionResult `json:"result"`
}

type DescribeAlarmHistoryAllRegionResult struct {
    List []monitor.AlarmHistoryWithDetail `json:"list"`
    Total int64 `json:"total"`
}