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

type ModifyCacheInstanceClassRequest struct {

    core.JDCloudRequest

    /* 缓存Redis实例所在区域的Region ID。目前缓存Redis有华北、华南、华东区域，对应Region ID为cn-north-1、cn-south-1、cn-east-2  */
    RegionId string `json:"regionId"`

    /* 缓存Redis实例ID，是访问实例的唯一标识。  */
    CacheInstanceId string `json:"cacheInstanceId"`

    /* 变更后的缓存Redis规格，详细参见：<a href="https://www.jdcloud.com/help/detail/411/isCatalog/1">实例规格代码</a>  */
    CacheInstanceClass string `json:"cacheInstanceClass"`
}

/*
 * param regionId: 缓存Redis实例所在区域的Region ID。目前缓存Redis有华北、华南、华东区域，对应Region ID为cn-north-1、cn-south-1、cn-east-2 (Required)
 * param cacheInstanceId: 缓存Redis实例ID，是访问实例的唯一标识。 (Required)
 * param cacheInstanceClass: 变更后的缓存Redis规格，详细参见：<a href="https://www.jdcloud.com/help/detail/411/isCatalog/1">实例规格代码</a> (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewModifyCacheInstanceClassRequest(
    regionId string,
    cacheInstanceId string,
    cacheInstanceClass string,
) *ModifyCacheInstanceClassRequest {

	return &ModifyCacheInstanceClassRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/cacheInstance/{cacheInstanceId}:modifyCacheInstanceClass",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        CacheInstanceId: cacheInstanceId,
        CacheInstanceClass: cacheInstanceClass,
	}
}

/*
 * param regionId: 缓存Redis实例所在区域的Region ID。目前缓存Redis有华北、华南、华东区域，对应Region ID为cn-north-1、cn-south-1、cn-east-2 (Required)
 * param cacheInstanceId: 缓存Redis实例ID，是访问实例的唯一标识。 (Required)
 * param cacheInstanceClass: 变更后的缓存Redis规格，详细参见：<a href="https://www.jdcloud.com/help/detail/411/isCatalog/1">实例规格代码</a> (Required)
 */
func NewModifyCacheInstanceClassRequestWithAllParams(
    regionId string,
    cacheInstanceId string,
    cacheInstanceClass string,
) *ModifyCacheInstanceClassRequest {

    return &ModifyCacheInstanceClassRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/cacheInstance/{cacheInstanceId}:modifyCacheInstanceClass",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        CacheInstanceId: cacheInstanceId,
        CacheInstanceClass: cacheInstanceClass,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewModifyCacheInstanceClassRequestWithoutParam() *ModifyCacheInstanceClassRequest {

    return &ModifyCacheInstanceClassRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/cacheInstance/{cacheInstanceId}:modifyCacheInstanceClass",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 缓存Redis实例所在区域的Region ID。目前缓存Redis有华北、华南、华东区域，对应Region ID为cn-north-1、cn-south-1、cn-east-2(Required) */
func (r *ModifyCacheInstanceClassRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param cacheInstanceId: 缓存Redis实例ID，是访问实例的唯一标识。(Required) */
func (r *ModifyCacheInstanceClassRequest) SetCacheInstanceId(cacheInstanceId string) {
    r.CacheInstanceId = cacheInstanceId
}

/* param cacheInstanceClass: 变更后的缓存Redis规格，详细参见：<a href="https://www.jdcloud.com/help/detail/411/isCatalog/1">实例规格代码</a>(Required) */
func (r *ModifyCacheInstanceClassRequest) SetCacheInstanceClass(cacheInstanceClass string) {
    r.CacheInstanceClass = cacheInstanceClass
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ModifyCacheInstanceClassRequest) GetRegionId() string {
    return r.RegionId
}

type ModifyCacheInstanceClassResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ModifyCacheInstanceClassResult `json:"result"`
}

type ModifyCacheInstanceClassResult struct {
    OrderNum string `json:"orderNum"`
}