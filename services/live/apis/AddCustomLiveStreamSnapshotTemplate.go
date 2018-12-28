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

type AddCustomLiveStreamSnapshotTemplateRequest struct {

    core.JDCloudRequest

    /* 图片格式  */
    Format string `json:"format"`

    /* 图片宽度  */
    Width int `json:"width"`

    /* 范围  */
    Height int `json:"height"`

    /* 截图与设定的宽高不匹配时的处理规则  */
    FillType int `json:"fillType"`

    /* 截图周期  */
    SnapshotInterval int `json:"snapshotInterval"`

    /* 存储模式  */
    SaveMode int `json:"saveMode"`

    /* 保存bucket  */
    SaveBucket string `json:"saveBucket"`

    /* 保存endPoint  */
    SaveEndpoint string `json:"saveEndpoint"`

    /* 录制模板自定义名称  */
    Template string `json:"template"`
}

/*
 * param format: 图片格式 (Required)
 * param width: 图片宽度 (Required)
 * param height: 范围 (Required)
 * param fillType: 截图与设定的宽高不匹配时的处理规则 (Required)
 * param snapshotInterval: 截图周期 (Required)
 * param saveMode: 存储模式 (Required)
 * param saveBucket: 保存bucket (Required)
 * param saveEndpoint: 保存endPoint (Required)
 * param template: 录制模板自定义名称 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewAddCustomLiveStreamSnapshotTemplateRequest(
    format string,
    width int,
    height int,
    fillType int,
    snapshotInterval int,
    saveMode int,
    saveBucket string,
    saveEndpoint string,
    template string,
) *AddCustomLiveStreamSnapshotTemplateRequest {

	return &AddCustomLiveStreamSnapshotTemplateRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/snapshotCustoms:template",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        Format: format,
        Width: width,
        Height: height,
        FillType: fillType,
        SnapshotInterval: snapshotInterval,
        SaveMode: saveMode,
        SaveBucket: saveBucket,
        SaveEndpoint: saveEndpoint,
        Template: template,
	}
}

/*
 * param format: 图片格式 (Required)
 * param width: 图片宽度 (Required)
 * param height: 范围 (Required)
 * param fillType: 截图与设定的宽高不匹配时的处理规则 (Required)
 * param snapshotInterval: 截图周期 (Required)
 * param saveMode: 存储模式 (Required)
 * param saveBucket: 保存bucket (Required)
 * param saveEndpoint: 保存endPoint (Required)
 * param template: 录制模板自定义名称 (Required)
 */
func NewAddCustomLiveStreamSnapshotTemplateRequestWithAllParams(
    format string,
    width int,
    height int,
    fillType int,
    snapshotInterval int,
    saveMode int,
    saveBucket string,
    saveEndpoint string,
    template string,
) *AddCustomLiveStreamSnapshotTemplateRequest {

    return &AddCustomLiveStreamSnapshotTemplateRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/snapshotCustoms:template",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        Format: format,
        Width: width,
        Height: height,
        FillType: fillType,
        SnapshotInterval: snapshotInterval,
        SaveMode: saveMode,
        SaveBucket: saveBucket,
        SaveEndpoint: saveEndpoint,
        Template: template,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewAddCustomLiveStreamSnapshotTemplateRequestWithoutParam() *AddCustomLiveStreamSnapshotTemplateRequest {

    return &AddCustomLiveStreamSnapshotTemplateRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/snapshotCustoms:template",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param format: 图片格式(Required) */
func (r *AddCustomLiveStreamSnapshotTemplateRequest) SetFormat(format string) {
    r.Format = format
}

/* param width: 图片宽度(Required) */
func (r *AddCustomLiveStreamSnapshotTemplateRequest) SetWidth(width int) {
    r.Width = width
}

/* param height: 范围(Required) */
func (r *AddCustomLiveStreamSnapshotTemplateRequest) SetHeight(height int) {
    r.Height = height
}

/* param fillType: 截图与设定的宽高不匹配时的处理规则(Required) */
func (r *AddCustomLiveStreamSnapshotTemplateRequest) SetFillType(fillType int) {
    r.FillType = fillType
}

/* param snapshotInterval: 截图周期(Required) */
func (r *AddCustomLiveStreamSnapshotTemplateRequest) SetSnapshotInterval(snapshotInterval int) {
    r.SnapshotInterval = snapshotInterval
}

/* param saveMode: 存储模式(Required) */
func (r *AddCustomLiveStreamSnapshotTemplateRequest) SetSaveMode(saveMode int) {
    r.SaveMode = saveMode
}

/* param saveBucket: 保存bucket(Required) */
func (r *AddCustomLiveStreamSnapshotTemplateRequest) SetSaveBucket(saveBucket string) {
    r.SaveBucket = saveBucket
}

/* param saveEndpoint: 保存endPoint(Required) */
func (r *AddCustomLiveStreamSnapshotTemplateRequest) SetSaveEndpoint(saveEndpoint string) {
    r.SaveEndpoint = saveEndpoint
}

/* param template: 录制模板自定义名称(Required) */
func (r *AddCustomLiveStreamSnapshotTemplateRequest) SetTemplate(template string) {
    r.Template = template
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r AddCustomLiveStreamSnapshotTemplateRequest) GetRegionId() string {
    return ""
}

type AddCustomLiveStreamSnapshotTemplateResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result AddCustomLiveStreamSnapshotTemplateResult `json:"result"`
}

type AddCustomLiveStreamSnapshotTemplateResult struct {
    Feedback bool `json:"feedback"`
}