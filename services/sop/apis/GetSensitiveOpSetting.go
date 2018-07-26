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

type GetSensitiveOpSettingRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 操作action serviceName:actionName  */
    Action string `json:"action"`
}

/*
 * param regionId: Region ID (Required)
 * param action: 操作action serviceName:actionName (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGetSensitiveOpSettingRequest(
    regionId string,
    action string,
) *GetSensitiveOpSettingRequest {

	return &GetSensitiveOpSettingRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/sensitiveOpSetting",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        Action: action,
	}
}

/*
 * param regionId: Region ID (Required)
 * param action: 操作action serviceName:actionName (Required)
 */
func NewGetSensitiveOpSettingRequestWithAllParams(
    regionId string,
    action string,
) *GetSensitiveOpSettingRequest {

    return &GetSensitiveOpSettingRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/sensitiveOpSetting",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Action: action,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGetSensitiveOpSettingRequestWithoutParam() *GetSensitiveOpSettingRequest {

    return &GetSensitiveOpSettingRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/sensitiveOpSetting",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *GetSensitiveOpSettingRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param action: 操作action serviceName:actionName(Required) */
func (r *GetSensitiveOpSettingRequest) SetAction(action string) {
    r.Action = action
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GetSensitiveOpSettingRequest) GetRegionId() string {
    return r.RegionId
}

type GetSensitiveOpSettingResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GetSensitiveOpSettingResult `json:"result"`
}

type GetSensitiveOpSettingResult struct {
    Status int `json:"status"`
    Type int `json:"type"`
    ExtInfo string `json:"extInfo"`
}