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
    jmr "github.com/jdcloud-api/jdcloud-sdk-go/services/jmr/models"
)

type QueryVpcSubnetsRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /*   */
    VpcId string `json:"vpcId"`
}

/*
 * param regionId: 地域ID (Required)
 * param vpcId:  (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryVpcSubnetsRequest(
    regionId string,
    vpcId string,
) *QueryVpcSubnetsRequest {

	return &QueryVpcSubnetsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/vpcSubnets/{vpcId}:query",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        VpcId: vpcId,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param vpcId:  (Required)
 */
func NewQueryVpcSubnetsRequestWithAllParams(
    regionId string,
    vpcId string,
) *QueryVpcSubnetsRequest {

    return &QueryVpcSubnetsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/vpcSubnets/{vpcId}:query",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        VpcId: vpcId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryVpcSubnetsRequestWithoutParam() *QueryVpcSubnetsRequest {

    return &QueryVpcSubnetsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/vpcSubnets/{vpcId}:query",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *QueryVpcSubnetsRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param vpcId: (Required) */
func (r *QueryVpcSubnetsRequest) SetVpcId(vpcId string) {
    r.VpcId = vpcId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryVpcSubnetsRequest) GetRegionId() string {
    return r.RegionId
}

type QueryVpcSubnetsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryVpcSubnetsResult `json:"result"`
}

type QueryVpcSubnetsResult struct {
    Status string `json:"status"`
    Message string `json:"message"`
    Data jmr.AvailableNumData `json:"data"`
}