// Copyright 2018-2025 JDCLOUD.COM
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
    . "github.com/jdcloud-api/jdcloud-sdk-go/core"
    "reflect"
    monitor "github.com/jdcloud-api/jdcloud-sdk-go/services/monitor/models"
)

type DescribeMetricDataRequest struct {

    JDCloudRequest

    /* 地域 Id  */
    RegionId string `json:"regionId"`

    /* 监控项英文标识(id)  */
    Metric string `json:"metric"`

    /* 资源的类型，取值vm, lb, ip, database 等  */
    ServiceCode string `json:"serviceCode"`

    /* 资源的uuid  */
    ResourceId string `json:"resourceId"`

    /* 查询时间范围的开始时间， UTC时间，格式：yyyy-MM-dd'T'HH:mm:ssZ（默认为当前时间，早于30d时，将被重置为30d） (Optional) */
    StartTime *string `json:"startTime"`

    /* 查询时间范围的结束时间， UTC时间，格式：2016-12- yyyy-MM-dd'T'HH:mm:ssZ（为空时，将由startTime与timeInterval计算得出） (Optional) */
    EndTime *string `json:"endTime"`

    /* 时间间隔：1h，6h，12h，1d，3d，7d，14d，固定时间间隔，timeInterval 与 endTime 至少填一项 (Optional) */
    TimeInterval *string `json:"timeInterval"`
}

/*
 * param regionId: 地域 Id 
 * param metric: 监控项英文标识(id) 
 * param serviceCode: 资源的类型，取值vm, lb, ip, database 等 
 * param resourceId: 资源的uuid 
 * param startTime: 查询时间范围的开始时间， UTC时间，格式：yyyy-MM-dd'T'HH:mm:ssZ（默认为当前时间，早于30d时，将被重置为30d） (Optional)
 * param endTime: 查询时间范围的结束时间， UTC时间，格式：2016-12- yyyy-MM-dd'T'HH:mm:ssZ（为空时，将由startTime与timeInterval计算得出） (Optional)
 * param timeInterval: 时间间隔：1h，6h，12h，1d，3d，7d，14d，固定时间间隔，timeInterval 与 endTime 至少填一项 (Optional)
 */
func NewDescribeMetricDataRequest(
    regionId string,
    metric string,
    serviceCode string,
    resourceId string,
) *DescribeMetricDataRequest {

	return &DescribeMetricDataRequest{
        JDCloudRequest: JDCloudRequest{
			URL:     "/regions/{regionId}/metrics/{metric}/metricData",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        Metric: metric,
        ServiceCode: serviceCode,
        ResourceId: resourceId,
	}
}

func (r *DescribeMetricDataRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

func (r *DescribeMetricDataRequest) SetMetric(metric string) {
    r.Metric = metric
}

func (r *DescribeMetricDataRequest) SetServiceCode(serviceCode string) {
    r.ServiceCode = serviceCode
}

func (r *DescribeMetricDataRequest) SetResourceId(resourceId string) {
    r.ResourceId = resourceId
}

func (r *DescribeMetricDataRequest) SetStartTime(startTime string) {
    r.StartTime = &startTime
}

func (r *DescribeMetricDataRequest) SetEndTime(endTime string) {
    r.EndTime = &endTime
}

func (r *DescribeMetricDataRequest) SetTimeInterval(timeInterval string) {
    r.TimeInterval = &timeInterval
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeMetricDataRequest) GetRegionId() string {
    fieldName := "RegionId"
    reqType := reflect.TypeOf(r)
    value := reflect.ValueOf(r)
    _, ok := reqType.FieldByName(fieldName)
    if ok {
        return value.FieldByName(fieldName).String()
    }

    return ""
}

type DescribeMetricDataResponse struct {
    RequestID string `json:"requestId"`
    Error ErrorResponse `json:"error"`
    Result DescribeMetricDataResult `json:"result"`
}

type DescribeMetricDataResult struct {
    MetricDatas []monitor.MetricData `json:"metricDatas"`
}