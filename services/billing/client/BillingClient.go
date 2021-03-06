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
    billing "github.com/jdcloud-api/jdcloud-sdk-go/services/billing/apis"
    "encoding/json"
    "errors"
)

type BillingClient struct {
    core.JDCloudClient
}

func NewBillingClient(credential *core.Credential) *BillingClient {
    if credential == nil {
        return nil
    }

    config := core.NewConfig()
    config.SetEndpoint("billing.jdcloud-api.com")

    return &BillingClient{
        core.JDCloudClient{
            Credential:  *credential,
            Config:      *config,
            ServiceName: "billing",
            Revision:    "0.4.0",
            Logger:      core.NewDefaultLogger(core.LogInfo),
        }}
}

func (c *BillingClient) SetConfig(config *core.Config) {
    c.Config = *config
}

func (c *BillingClient) SetLogger(logger core.Logger) {
    c.Logger = logger
}

/* 查询消费总览 */
func (c *BillingClient) QueryConsumptionOverView(request *billing.QueryConsumptionOverViewRequest) (*billing.QueryConsumptionOverViewResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &billing.QueryConsumptionOverViewResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 查询资源账单列表 */
func (c *BillingClient) QueryResourceBills(request *billing.QueryResourceBillsRequest) (*billing.QueryResourceBillsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &billing.QueryResourceBillsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 查询资源单列表接口，不含已删除资源 */
func (c *BillingClient) GetExpiringOrders(request *billing.GetExpiringOrdersRequest) (*billing.GetExpiringOrdersResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &billing.GetExpiringOrdersResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 查询资源账单详情 */
func (c *BillingClient) QueryConsumeRecords(request *billing.QueryConsumeRecordsRequest) (*billing.QueryConsumeRecordsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &billing.QueryConsumeRecordsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 计算赔偿金额并发放代金券 */
func (c *BillingClient) CalculateCompensateFeeAndSendCoupons(request *billing.CalculateCompensateFeeAndSendCouponsRequest) (*billing.CalculateCompensateFeeAndSendCouponsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &billing.CalculateCompensateFeeAndSendCouponsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 查询消费记录列表 */
func (c *BillingClient) QueryConsumeBills(request *billing.QueryConsumeBillsRequest) (*billing.QueryConsumeBillsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &billing.QueryConsumeBillsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 查询资源账单，消费记录中的费用信息 */
func (c *BillingClient) QueryBillStatisticsInfo(request *billing.QueryBillStatisticsInfoRequest) (*billing.QueryBillStatisticsInfoResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &billing.QueryBillStatisticsInfoResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 查询消费统计信息 */
func (c *BillingClient) AdminQueryBillStatisticsInfo(request *billing.AdminQueryBillStatisticsInfoRequest) (*billing.AdminQueryBillStatisticsInfoResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &billing.AdminQueryBillStatisticsInfoResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 查询计费价格信息 */
func (c *BillingClient) CalculateTotalPrice(request *billing.CalculateTotalPriceRequest) (*billing.CalculateTotalPriceResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &billing.CalculateTotalPriceResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 查询消费记录详情 */
func (c *BillingClient) GetBillDetail(request *billing.GetBillDetailRequest) (*billing.GetBillDetailResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &billing.GetBillDetailResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 查询资源单列表 */
func (c *BillingClient) QueryPageByCondition(request *billing.QueryPageByConditionRequest) (*billing.QueryPageByConditionResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &billing.QueryPageByConditionResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 查询用户是否欠费 */
func (c *BillingClient) IsArrear(request *billing.IsArrearRequest) (*billing.IsArrearResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &billing.IsArrearResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 查询用于在账单展示的资源名称信息 */
func (c *BillingClient) GetResourceName(request *billing.GetResourceNameRequest) (*billing.GetResourceNameResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &billing.GetResourceNameResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 删除资源 */
func (c *BillingClient) SendResourceOrderStatusMessage(request *billing.SendResourceOrderStatusMessageRequest) (*billing.SendResourceOrderStatusMessageResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &billing.SendResourceOrderStatusMessageResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

