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
    monitor "github.com/jdcloud-api/jdcloud-sdk-go/services/monitor/models"
)

type ApplyTemplateRequest struct {

    core.JDCloudRequest

    /* 模板 id  */
    TemplateId string `json:"templateId"`

    /* 幂等性校验参数,最长36位  */
    ClientToken string `json:"clientToken"`

    /* 联系人列表 (Optional) */
    Contacts []monitor.BaseContact `json:"contacts"`

    /* dataCenters，该资源所处地域，与resourceIds一一对应  */
    DataCenters []string `json:"dataCenters"`

    /* 资源ID列表，resourceIds数量为1--40，每一项不能为空或包含*  */
    ResourceIds []string `json:"resourceIds"`

    /* 资源类型  */
    ServiceCode string `json:"serviceCode"`

    /* 标签组，模板下面规则对应的tags，与resourceIds一一对应 (Optional) */
    TagsArray []interface{} `json:"tagsArray"`

    /* 模板类型，1表示默认模板，2表示用户自定义模板  */
    TemplateType int64 `json:"templateType"`

    /* 回调content 注:仅webHookUrl和webHookContent均不为空时,才会创建webHook (Optional) */
    WebHookContent *string `json:"webHookContent"`

    /* webHook协议,https或http 注:此处需和webHookUrl相匹配 (Optional) */
    WebHookProtocol *string `json:"webHookProtocol"`

    /* 回调secret,用户请求签名,防伪造 (Optional) */
    WebHookSecret *string `json:"webHookSecret"`

    /* 回调url (Optional) */
    WebHookUrl *string `json:"webHookUrl"`
}

/*
 * param templateId: 模板 id (Required)
 * param clientToken: 幂等性校验参数,最长36位 (Required)
 * param dataCenters: dataCenters，该资源所处地域，与resourceIds一一对应 (Required)
 * param resourceIds: 资源ID列表，resourceIds数量为1--40，每一项不能为空或包含* (Required)
 * param serviceCode: 资源类型 (Required)
 * param templateType: 模板类型，1表示默认模板，2表示用户自定义模板 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewApplyTemplateRequest(
    templateId string,
    clientToken string,
    dataCenters []string,
    resourceIds []string,
    serviceCode string,
    templateType int64,
) *ApplyTemplateRequest {

	return &ApplyTemplateRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/template/{templateId}/resources",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        TemplateId: templateId,
        ClientToken: clientToken,
        DataCenters: dataCenters,
        ResourceIds: resourceIds,
        ServiceCode: serviceCode,
        TemplateType: templateType,
	}
}

/*
 * param templateId: 模板 id (Required)
 * param clientToken: 幂等性校验参数,最长36位 (Required)
 * param contacts: 联系人列表 (Optional)
 * param dataCenters: dataCenters，该资源所处地域，与resourceIds一一对应 (Required)
 * param resourceIds: 资源ID列表，resourceIds数量为1--40，每一项不能为空或包含* (Required)
 * param serviceCode: 资源类型 (Required)
 * param tagsArray: 标签组，模板下面规则对应的tags，与resourceIds一一对应 (Optional)
 * param templateType: 模板类型，1表示默认模板，2表示用户自定义模板 (Required)
 * param webHookContent: 回调content 注:仅webHookUrl和webHookContent均不为空时,才会创建webHook (Optional)
 * param webHookProtocol: webHook协议,https或http 注:此处需和webHookUrl相匹配 (Optional)
 * param webHookSecret: 回调secret,用户请求签名,防伪造 (Optional)
 * param webHookUrl: 回调url (Optional)
 */
func NewApplyTemplateRequestWithAllParams(
    templateId string,
    clientToken string,
    contacts []monitor.BaseContact,
    dataCenters []string,
    resourceIds []string,
    serviceCode string,
    tagsArray []interface{},
    templateType int64,
    webHookContent *string,
    webHookProtocol *string,
    webHookSecret *string,
    webHookUrl *string,
) *ApplyTemplateRequest {

    return &ApplyTemplateRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/template/{templateId}/resources",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        TemplateId: templateId,
        ClientToken: clientToken,
        Contacts: contacts,
        DataCenters: dataCenters,
        ResourceIds: resourceIds,
        ServiceCode: serviceCode,
        TagsArray: tagsArray,
        TemplateType: templateType,
        WebHookContent: webHookContent,
        WebHookProtocol: webHookProtocol,
        WebHookSecret: webHookSecret,
        WebHookUrl: webHookUrl,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewApplyTemplateRequestWithoutParam() *ApplyTemplateRequest {

    return &ApplyTemplateRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/template/{templateId}/resources",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param templateId: 模板 id(Required) */
func (r *ApplyTemplateRequest) SetTemplateId(templateId string) {
    r.TemplateId = templateId
}

/* param clientToken: 幂等性校验参数,最长36位(Required) */
func (r *ApplyTemplateRequest) SetClientToken(clientToken string) {
    r.ClientToken = clientToken
}

/* param contacts: 联系人列表(Optional) */
func (r *ApplyTemplateRequest) SetContacts(contacts []monitor.BaseContact) {
    r.Contacts = contacts
}

/* param dataCenters: dataCenters，该资源所处地域，与resourceIds一一对应(Required) */
func (r *ApplyTemplateRequest) SetDataCenters(dataCenters []string) {
    r.DataCenters = dataCenters
}

/* param resourceIds: 资源ID列表，resourceIds数量为1--40，每一项不能为空或包含*(Required) */
func (r *ApplyTemplateRequest) SetResourceIds(resourceIds []string) {
    r.ResourceIds = resourceIds
}

/* param serviceCode: 资源类型(Required) */
func (r *ApplyTemplateRequest) SetServiceCode(serviceCode string) {
    r.ServiceCode = serviceCode
}

/* param tagsArray: 标签组，模板下面规则对应的tags，与resourceIds一一对应(Optional) */
func (r *ApplyTemplateRequest) SetTagsArray(tagsArray []interface{}) {
    r.TagsArray = tagsArray
}

/* param templateType: 模板类型，1表示默认模板，2表示用户自定义模板(Required) */
func (r *ApplyTemplateRequest) SetTemplateType(templateType int64) {
    r.TemplateType = templateType
}

/* param webHookContent: 回调content 注:仅webHookUrl和webHookContent均不为空时,才会创建webHook(Optional) */
func (r *ApplyTemplateRequest) SetWebHookContent(webHookContent string) {
    r.WebHookContent = &webHookContent
}

/* param webHookProtocol: webHook协议,https或http 注:此处需和webHookUrl相匹配(Optional) */
func (r *ApplyTemplateRequest) SetWebHookProtocol(webHookProtocol string) {
    r.WebHookProtocol = &webHookProtocol
}

/* param webHookSecret: 回调secret,用户请求签名,防伪造(Optional) */
func (r *ApplyTemplateRequest) SetWebHookSecret(webHookSecret string) {
    r.WebHookSecret = &webHookSecret
}

/* param webHookUrl: 回调url(Optional) */
func (r *ApplyTemplateRequest) SetWebHookUrl(webHookUrl string) {
    r.WebHookUrl = &webHookUrl
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ApplyTemplateRequest) GetRegionId() string {
    return ""
}

type ApplyTemplateResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ApplyTemplateResult `json:"result"`
}

type ApplyTemplateResult struct {
    RuleIds []int64 `json:"ruleIds"`
    Success bool `json:"success"`
}