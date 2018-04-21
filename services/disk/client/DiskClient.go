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

package client

import (
    . "github.com/jdcloud-api/jdcloud-sdk-go/core"
    . "github.com/jdcloud-api/jdcloud-sdk-go/services/disk/apis"
    "encoding/json"
    "errors"
)

type DiskClient struct {
    JDCloudClient
}

func NewDiskClient(credential *Credential) *DiskClient {
    if credential == nil {
        return nil
    }

    config := NewConfig()
    config.SetEndpoint("disk.jdcloud-api.com")

    return &DiskClient{
        JDCloudClient{
            Credential:  *credential,
            Config:      *config,
            ServiceName: "disk",
            Revision:    "0.3.0",
            Logger:      NewDefaultLogger(LOG_INFO),
        }}
}

func (c *DiskClient) SetConfig(config *Config) {
    c.Config = *config
}

func (c *DiskClient) SetLogger(logger Logger) {
    c.Logger = logger
}

/* 为指定云硬盘创建快照调用成功返回后，新生成的快照的状态为 creating */
func (c *DiskClient) CreateSnapshot(request *CreateSnapshotRequest) (*CreateSnapshotResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &CreateSnapshotResponse{}
    err = json.Unmarshal(resp, jdResp)
    return jdResp, err
}

/* 从已有快照恢复一块云硬盘 */
func (c *DiskClient) RestoreDisk(request *RestoreDiskRequest) (*RestoreDiskResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &RestoreDiskResponse{}
    err = json.Unmarshal(resp, jdResp)
    return jdResp, err
}

/* 查询云硬盘列表 */
func (c *DiskClient) DescribeDisks(request *DescribeDisksRequest) (*DescribeDisksResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &DescribeDisksResponse{}
    err = json.Unmarshal(resp, jdResp)
    return jdResp, err
}

/* 查询云硬盘快照列表 */
func (c *DiskClient) DescribeSnapshots(request *DescribeSnapshotsRequest) (*DescribeSnapshotsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &DescribeSnapshotsResponse{}
    err = json.Unmarshal(resp, jdResp)
    return jdResp, err
}

/* 云硬盘信息详情 */
func (c *DiskClient) DescribeDisk(request *DescribeDiskRequest) (*DescribeDiskResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &DescribeDiskResponse{}
    err = json.Unmarshal(resp, jdResp)
    return jdResp, err
}

/* 云硬盘快照信息详情 */
func (c *DiskClient) DescribeSnapshot(request *DescribeSnapshotRequest) (*DescribeSnapshotResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &DescribeSnapshotResponse{}
    err = json.Unmarshal(resp, jdResp)
    return jdResp, err
}
