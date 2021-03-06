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

type ResumeLiveStreamRequest struct {

    core.JDCloudRequest

    /* 应用名称  */
    AppName string `json:"appName"`

    /* 您的加速域名  */
    PublishDomain string `json:"publishDomain"`

    /* 流名称  */
    StreamName string `json:"streamName"`
}

/*
 * param appName: 应用名称 (Required)
 * param publishDomain: 您的加速域名 (Required)
 * param streamName: 流名称 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewResumeLiveStreamRequest(
    appName string,
    publishDomain string,
    streamName string,
) *ResumeLiveStreamRequest {

	return &ResumeLiveStreamRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/streams:resume",
			Method:  "PUT",
			Header:  nil,
			Version: "v1",
		},
        AppName: appName,
        PublishDomain: publishDomain,
        StreamName: streamName,
	}
}

/*
 * param appName: 应用名称 (Required)
 * param publishDomain: 您的加速域名 (Required)
 * param streamName: 流名称 (Required)
 */
func NewResumeLiveStreamRequestWithAllParams(
    appName string,
    publishDomain string,
    streamName string,
) *ResumeLiveStreamRequest {

    return &ResumeLiveStreamRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/streams:resume",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
        AppName: appName,
        PublishDomain: publishDomain,
        StreamName: streamName,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewResumeLiveStreamRequestWithoutParam() *ResumeLiveStreamRequest {

    return &ResumeLiveStreamRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/streams:resume",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param appName: 应用名称(Required) */
func (r *ResumeLiveStreamRequest) SetAppName(appName string) {
    r.AppName = appName
}

/* param publishDomain: 您的加速域名(Required) */
func (r *ResumeLiveStreamRequest) SetPublishDomain(publishDomain string) {
    r.PublishDomain = publishDomain
}

/* param streamName: 流名称(Required) */
func (r *ResumeLiveStreamRequest) SetStreamName(streamName string) {
    r.StreamName = streamName
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ResumeLiveStreamRequest) GetRegionId() string {
    return ""
}

type ResumeLiveStreamResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ResumeLiveStreamResult `json:"result"`
}

type ResumeLiveStreamResult struct {
    PublishDomain string `json:"publishDomain"`
}