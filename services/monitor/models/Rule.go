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

package models


type Rule struct {

    /* AutoScalingPolicyID (Optional) */
    AutoScalingPolicyId string `json:"autoScalingPolicyId"`

    /* calculateUnit (Optional) */
    CalculateUnit string `json:"calculateUnit"`

    /* calculation (Optional) */
    Calculation string `json:"calculation"`

    /* createTime (Optional) */
    CreateTime string `json:"createTime"`

    /* dataMeaning (Optional) */
    DataMeaning int64 `json:"dataMeaning"`

    /* 地域 (Optional) */
    Datacenter string `json:"datacenter"`

    /* deleted (Optional) */
    Deleted int64 `json:"deleted"`

    /* downSample (Optional) */
    DownSample string `json:"downSample"`

    /* enableTime (Optional) */
    EnableTime string `json:"enableTime"`

    /* enabled (Optional) */
    Enabled int64 `json:"enabled"`

    /* uuid (Optional) */
    Id string `json:"id"`

    /*  (Optional) */
    Idpk int64 `json:"idpk"`

    /* isLatest (Optional) */
    IsLatest int64 `json:"isLatest"`

    /* metric (Optional) */
    Metric string `json:"metric"`

    /* metricId (Optional) */
    MetricId int64 `json:"metricId"`

    /* metricName (Optional) */
    MetricName string `json:"metricName"`

    /*  (Optional) */
    NoticeLevel NoticeLevel `json:"noticeLevel"`

    /* notice_period (Optional) */
    NoticePeriod int64 `json:"noticePeriod"`

    /* operation (Optional) */
    Operation string `json:"operation"`

    /* period (Optional) */
    Period int64 `json:"period"`

    /* pin (Optional) */
    Pin string `json:"pin"`

    /* preVersionId (Optional) */
    PreVersionId int64 `json:"preVersionId"`

    /* region (Optional) */
    Region string `json:"region"`

    /* 资源 id (Optional) */
    ResourceId string `json:"resourceId"`

    /* rootRuleID (Optional) */
    RootRuleId int64 `json:"rootRuleId"`

    /* rutye 1:经典监控  5：自定义监控   6：站点监控 7：可用性监控 (Optional) */
    RuleType int64 `json:"ruleType"`

    /* serviceCode (Optional) */
    ServiceCode string `json:"serviceCode"`

    /* status (Optional) */
    Status int64 `json:"status"`

    /* statusTime (Optional) */
    StatusTime string `json:"statusTime"`

    /* tags (Optional) */
    Tags interface{} `json:"tags"`

    /* tagsNonGrouping (Optional) */
    TagsNonGrouping string `json:"tagsNonGrouping"`

    /* threshold (Optional) */
    Threshold float64 `json:"threshold"`

    /* times (Optional) */
    Times int64 `json:"times"`

    /* updateTime (Optional) */
    UpdateTime string `json:"updateTime"`

    /* version (Optional) */
    Version int64 `json:"version"`
}