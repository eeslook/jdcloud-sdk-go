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

package client

import (
    "github.com/jdcloud-api/jdcloud-sdk-go/core"
    jdfusion "github.com/jdcloud-api/jdcloud-sdk-go/services/jdfusion/apis"
    "encoding/json"
    "errors"
)

type JdfusionClient struct {
    core.JDCloudClient
}

func NewJdfusionClient(credential *core.Credential) *JdfusionClient {
    if credential == nil {
        return nil
    }

    config := core.NewConfig()
    config.SetEndpoint("jdfusion.jdcloud-api.com")

    return &JdfusionClient{
        core.JDCloudClient{
            Credential:  *credential,
            Config:      *config,
            ServiceName: "jdfusion",
            Revision:    "0.1.0",
            Logger:      core.NewDefaultLogger(core.LogInfo),
        }}
}

func (c *JdfusionClient) SetConfig(config *core.Config) {
    c.Config = *config
}

func (c *JdfusionClient) SetLogger(logger core.Logger) {
    c.Logger = logger
}

/* 虚拟机规格列表 */
func (c *JdfusionClient) GetVmInstanceTypes(request *jdfusion.GetVmInstanceTypesRequest) (*jdfusion.GetVmInstanceTypesResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdfusion.GetVmInstanceTypesResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 读取指定ID的运行结果和运行状态 */
func (c *JdfusionClient) GetTaskInfoHistoryById(request *jdfusion.GetTaskInfoHistoryByIdRequest) (*jdfusion.GetTaskInfoHistoryByIdResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdfusion.GetTaskInfoHistoryByIdResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 根据云提供商查询对应的网卡资源信息 */
func (c *JdfusionClient) GetVpcNetworkInterfaces(request *jdfusion.GetVpcNetworkInterfacesRequest) (*jdfusion.GetVpcNetworkInterfacesResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdfusion.GetVpcNetworkInterfacesResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 根据云提供商查询对应的负载均衡资源信息 */
func (c *JdfusionClient) GetVpcSlbs(request *jdfusion.GetVpcSlbsRequest) (*jdfusion.GetVpcSlbsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdfusion.GetVpcSlbsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 网卡挂载虚拟机 */
func (c *JdfusionClient) AttachVpcNetworkInterfaceById(request *jdfusion.AttachVpcNetworkInterfaceByIdRequest) (*jdfusion.AttachVpcNetworkInterfaceByIdResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdfusion.AttachVpcNetworkInterfaceByIdResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 根据云提供商创建云硬盘 */
func (c *JdfusionClient) CreateDisk(request *jdfusion.CreateDiskRequest) (*jdfusion.CreateDiskResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdfusion.CreateDiskResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 根据云提供商创建网卡 */
func (c *JdfusionClient) CreateVpcNetworkInterface(request *jdfusion.CreateVpcNetworkInterfaceRequest) (*jdfusion.CreateVpcNetworkInterfaceResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdfusion.CreateVpcNetworkInterfaceResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 只能重启状态为 运行中（Running）的实例。 */
func (c *JdfusionClient) RebootVmInstanceById(request *jdfusion.RebootVmInstanceByIdRequest) (*jdfusion.RebootVmInstanceByIdResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdfusion.RebootVmInstanceByIdResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 根据云提供商查询对应的云硬盘资源信息 */
func (c *JdfusionClient) GetDiskById(request *jdfusion.GetDiskByIdRequest) (*jdfusion.GetDiskByIdResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdfusion.GetDiskByIdResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 根据云提供商查询对应的VM资源信息 */
func (c *JdfusionClient) GetVmInstancesById(request *jdfusion.GetVmInstancesByIdRequest) (*jdfusion.GetVmInstancesByIdResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdfusion.GetVmInstancesByIdResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 根据云提供商创建私有网络 */
func (c *JdfusionClient) CreateVpc(request *jdfusion.CreateVpcRequest) (*jdfusion.CreateVpcResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdfusion.CreateVpcResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 根据云提供商查询对应的网卡资源信息 */
func (c *JdfusionClient) GetVpcNetworkInterfaceById(request *jdfusion.GetVpcNetworkInterfaceByIdRequest) (*jdfusion.GetVpcNetworkInterfaceByIdResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdfusion.GetVpcNetworkInterfaceByIdResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 根据云提供商查询对应的SLB资源信息 */
func (c *JdfusionClient) GetVpcSlbById(request *jdfusion.GetVpcSlbByIdRequest) (*jdfusion.GetVpcSlbByIdResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdfusion.GetVpcSlbByIdResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 根据云提供商创建子网 */
func (c *JdfusionClient) CreateVpcSubnet(request *jdfusion.CreateVpcSubnetRequest) (*jdfusion.CreateVpcSubnetResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdfusion.CreateVpcSubnetResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 删除弹性网卡 */
func (c *JdfusionClient) DeleteVpcNetworkInterfaceById(request *jdfusion.DeleteVpcNetworkInterfaceByIdRequest) (*jdfusion.DeleteVpcNetworkInterfaceByIdResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdfusion.DeleteVpcNetworkInterfaceByIdResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 卸载网卡 */
func (c *JdfusionClient) DetachVpcNetworkInterfaceById(request *jdfusion.DetachVpcNetworkInterfaceByIdRequest) (*jdfusion.DetachVpcNetworkInterfaceByIdResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdfusion.DetachVpcNetworkInterfaceByIdResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 为指定用户关联云 */
func (c *JdfusionClient) RegistCloudInfo(request *jdfusion.RegistCloudInfoRequest) (*jdfusion.RegistCloudInfoResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdfusion.RegistCloudInfoResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 根据云提供商查询对应的公网IP资源信息 */
func (c *JdfusionClient) GetVpcEipById(request *jdfusion.GetVpcEipByIdRequest) (*jdfusion.GetVpcEipByIdResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdfusion.GetVpcEipByIdResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 删除安全组 */
func (c *JdfusionClient) DeleteVpcSecurityGroupById(request *jdfusion.DeleteVpcSecurityGroupByIdRequest) (*jdfusion.DeleteVpcSecurityGroupByIdResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdfusion.DeleteVpcSecurityGroupByIdResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 删除云硬盘 */
func (c *JdfusionClient) RemoveDiskById(request *jdfusion.RemoveDiskByIdRequest) (*jdfusion.RemoveDiskByIdResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdfusion.RemoveDiskByIdResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 根据云提供商分配公网IP */
func (c *JdfusionClient) CreateVpcEip(request *jdfusion.CreateVpcEipRequest) (*jdfusion.CreateVpcEipResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdfusion.CreateVpcEipResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 弹性公网IP绑定到虚拟机 */
func (c *JdfusionClient) AssociateVpcEipById(request *jdfusion.AssociateVpcEipByIdRequest) (*jdfusion.AssociateVpcEipByIdResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdfusion.AssociateVpcEipByIdResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 根据云提供商创建虚拟机 */
func (c *JdfusionClient) CreateVmInstance(request *jdfusion.CreateVmInstanceRequest) (*jdfusion.CreateVmInstanceResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdfusion.CreateVmInstanceResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 获取指定云信息 */
func (c *JdfusionClient) UnregistCloudInfo(request *jdfusion.UnregistCloudInfoRequest) (*jdfusion.UnregistCloudInfoResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdfusion.UnregistCloudInfoResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 解绑公网IP */
func (c *JdfusionClient) DisassociateVpcEipById(request *jdfusion.DisassociateVpcEipByIdRequest) (*jdfusion.DisassociateVpcEipByIdResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdfusion.DisassociateVpcEipByIdResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 根据云提供商查询对应的安全组资源信息 */
func (c *JdfusionClient) GetVpcSecurityGroupById(request *jdfusion.GetVpcSecurityGroupByIdRequest) (*jdfusion.GetVpcSecurityGroupByIdResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdfusion.GetVpcSecurityGroupByIdResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 删除VPC */
func (c *JdfusionClient) DeleteVpcById(request *jdfusion.DeleteVpcByIdRequest) (*jdfusion.DeleteVpcByIdResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdfusion.DeleteVpcByIdResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 启动一台实例。实例状态必须为 已停止（Stopped），才可以调用该接口。 */
func (c *JdfusionClient) StartVmInstanceById(request *jdfusion.StartVmInstanceByIdRequest) (*jdfusion.StartVmInstanceByIdResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdfusion.StartVmInstanceByIdResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 获取指定云信息 */
func (c *JdfusionClient) GetCloudInfoById(request *jdfusion.GetCloudInfoByIdRequest) (*jdfusion.GetCloudInfoByIdResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdfusion.GetCloudInfoByIdResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 停止运行一台实例。只有状态为 运行中（Running）的实例才可以进行此操作。 */
func (c *JdfusionClient) StopVmInstanceById(request *jdfusion.StopVmInstanceByIdRequest) (*jdfusion.StopVmInstanceByIdResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdfusion.StopVmInstanceByIdResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 通过虚拟机id删除虚拟机 */
func (c *JdfusionClient) DeleteVmInstanceById(request *jdfusion.DeleteVmInstanceByIdRequest) (*jdfusion.DeleteVmInstanceByIdResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdfusion.DeleteVmInstanceByIdResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 删除密钥对 */
func (c *JdfusionClient) DeleteVmKeypairByName(request *jdfusion.DeleteVmKeypairByNameRequest) (*jdfusion.DeleteVmKeypairByNameResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdfusion.DeleteVmKeypairByNameResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 删除子网 */
func (c *JdfusionClient) DeleteVpcSubnetById(request *jdfusion.DeleteVpcSubnetByIdRequest) (*jdfusion.DeleteVpcSubnetByIdResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdfusion.DeleteVpcSubnetByIdResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 根据云提供商查询对应的VPC资源信息 */
func (c *JdfusionClient) GetVpcById(request *jdfusion.GetVpcByIdRequest) (*jdfusion.GetVpcByIdResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdfusion.GetVpcByIdResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 根据云提供商查询对应的虚拟机资源信息 */
func (c *JdfusionClient) GetVmInstances(request *jdfusion.GetVmInstancesRequest) (*jdfusion.GetVmInstancesResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdfusion.GetVmInstancesResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 根据云提供商查询对应的subnet资源信息 */
func (c *JdfusionClient) GetVpcSubnetById(request *jdfusion.GetVpcSubnetByIdRequest) (*jdfusion.GetVpcSubnetByIdResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdfusion.GetVpcSubnetByIdResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 根据云提供商创建密钥对 */
func (c *JdfusionClient) CreateVmKeypair(request *jdfusion.CreateVmKeypairRequest) (*jdfusion.CreateVmKeypairResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdfusion.CreateVmKeypairResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 根据云提供商查询对应的公网IP资源信息 */
func (c *JdfusionClient) GetVpcEips(request *jdfusion.GetVpcEipsRequest) (*jdfusion.GetVpcEipsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdfusion.GetVpcEipsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 根据云提供商创建负载均衡 */
func (c *JdfusionClient) CreateVpcSlb(request *jdfusion.CreateVpcSlbRequest) (*jdfusion.CreateVpcSlbResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdfusion.CreateVpcSlbResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 根据云提供商查询对应的安全组资源信息 */
func (c *JdfusionClient) GetVpcSecurityGroups(request *jdfusion.GetVpcSecurityGroupsRequest) (*jdfusion.GetVpcSecurityGroupsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdfusion.GetVpcSecurityGroupsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 根据云提供商查询对应的密钥对资源信息 */
func (c *JdfusionClient) GetVmKeypairsByName(request *jdfusion.GetVmKeypairsByNameRequest) (*jdfusion.GetVmKeypairsByNameResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdfusion.GetVmKeypairsByNameResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 云硬盘挂载至虚拟机 */
func (c *JdfusionClient) AttachDiskToVmInstanceByDiskId(request *jdfusion.AttachDiskToVmInstanceByDiskIdRequest) (*jdfusion.AttachDiskToVmInstanceByDiskIdResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdfusion.AttachDiskToVmInstanceByDiskIdResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 释放公网IP */
func (c *JdfusionClient) DeleteVpcEipById(request *jdfusion.DeleteVpcEipByIdRequest) (*jdfusion.DeleteVpcEipByIdResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdfusion.DeleteVpcEipByIdResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 根据云提供商创建安全组 */
func (c *JdfusionClient) CreateVpcSecurityGroup(request *jdfusion.CreateVpcSecurityGroupRequest) (*jdfusion.CreateVpcSecurityGroupResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdfusion.CreateVpcSecurityGroupResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 根据云提供商查询对应的subnet资源信息 */
func (c *JdfusionClient) GetVpcSubnets(request *jdfusion.GetVpcSubnetsRequest) (*jdfusion.GetVpcSubnetsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdfusion.GetVpcSubnetsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 根据过滤条件，取得服务器组的信息 */
func (c *JdfusionClient) GetVpcVServerGroups(request *jdfusion.GetVpcVServerGroupsRequest) (*jdfusion.GetVpcVServerGroupsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdfusion.GetVpcVServerGroupsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询私有网络资源列表 */
func (c *JdfusionClient) GetVpcs(request *jdfusion.GetVpcsRequest) (*jdfusion.GetVpcsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdfusion.GetVpcsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 读取指定ID的运行结果和运行状态 */
func (c *JdfusionClient) GetTaskInfoById(request *jdfusion.GetTaskInfoByIdRequest) (*jdfusion.GetTaskInfoByIdResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdfusion.GetTaskInfoByIdResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 获取当前用户的云注册信息 */
func (c *JdfusionClient) GetCloudInfos(request *jdfusion.GetCloudInfosRequest) (*jdfusion.GetCloudInfosResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdfusion.GetCloudInfosResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 根据云提供商查询对应的密钥对资源信息 */
func (c *JdfusionClient) GetVmKeypairs(request *jdfusion.GetVmKeypairsRequest) (*jdfusion.GetVmKeypairsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdfusion.GetVmKeypairsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 根据过滤条件，取得镜像资源的信息 */
func (c *JdfusionClient) GetVmImages(request *jdfusion.GetVmImagesRequest) (*jdfusion.GetVmImagesResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdfusion.GetVmImagesResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 删除SLB */
func (c *JdfusionClient) DeleteVpcSlbById(request *jdfusion.DeleteVpcSlbByIdRequest) (*jdfusion.DeleteVpcSlbByIdResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdfusion.DeleteVpcSlbByIdResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 创建虚拟服务器组，并添加后端服务器 */
func (c *JdfusionClient) CreateVpcVServerGroup(request *jdfusion.CreateVpcVServerGroupRequest) (*jdfusion.CreateVpcVServerGroupResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdfusion.CreateVpcVServerGroupResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 根据云提供商查询对应的云硬盘资源信息 */
func (c *JdfusionClient) GetDisks(request *jdfusion.GetDisksRequest) (*jdfusion.GetDisksResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdfusion.GetDisksResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 从虚拟机卸载云硬盘 */
func (c *JdfusionClient) DetachDiskToVmInstanceByDiskId(request *jdfusion.DetachDiskToVmInstanceByDiskIdRequest) (*jdfusion.DetachDiskToVmInstanceByDiskIdResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdfusion.DetachDiskToVmInstanceByDiskIdResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 创建HTTP监听器 */
func (c *JdfusionClient) CreateVpcLBHttpListener(request *jdfusion.CreateVpcLBHttpListenerRequest) (*jdfusion.CreateVpcLBHttpListenerResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jdfusion.CreateVpcLBHttpListenerResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}
