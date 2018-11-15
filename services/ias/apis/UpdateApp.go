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

type UpdateAppRequest struct {

    core.JDCloudRequest

    /*   */
    RegionId string `json:"regionId"`

    /*   */
    ClientId string `json:"clientId"`

    /* 应用名 (Optional) */
    ClientName *string `json:"clientName"`

    /* tokenEndpointAuthMethod (Optional) */
    TokenEndpointAuthMethod *string `json:"tokenEndpointAuthMethod"`

    /* grantTypes (Optional) */
    GrantTypes *string `json:"grantTypes"`

    /* redirectUris (Optional) */
    RedirectUris *string `json:"redirectUris"`

    /* clientUri (Optional) */
    ClientUri *string `json:"clientUri"`

    /* logoUri (Optional) */
    LogoUri *string `json:"logoUri"`

    /* tosUri (Optional) */
    TosUri *string `json:"tosUri"`

    /* policyUri (Optional) */
    PolicyUri *string `json:"policyUri"`

    /* scope (Optional) */
    Scope *string `json:"scope"`

    /* jwksUri (Optional) */
    JwksUri *string `json:"jwksUri"`

    /* jwks (Optional) */
    Jwks *string `json:"jwks"`

    /* contacts (Optional) */
    Contacts *string `json:"contacts"`

    /* extension (Optional) */
    Extension *string `json:"extension"`

    /* accessTokenValiditySeconds (Optional) */
    AccessTokenValiditySeconds *int `json:"accessTokenValiditySeconds"`

    /* refreshTokenValiditySeconds (Optional) */
    RefreshTokenValiditySeconds *int `json:"refreshTokenValiditySeconds"`

    /* multiTenant (Optional) */
    MultiTenant *bool `json:"multiTenant"`

    /* secret (Optional) */
    Secret *string `json:"secret"`

    /* userType (Optional) */
    UserType *string `json:"userType"`
}

/*
 * param regionId:  (Required)
 * param clientId:  (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewUpdateAppRequest(
    regionId string,
    clientId string,
) *UpdateAppRequest {

	return &UpdateAppRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/app/{clientId}",
			Method:  "PUT",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        ClientId: clientId,
	}
}

/*
 * param regionId:  (Required)
 * param clientId:  (Required)
 * param clientName: 应用名 (Optional)
 * param tokenEndpointAuthMethod: tokenEndpointAuthMethod (Optional)
 * param grantTypes: grantTypes (Optional)
 * param redirectUris: redirectUris (Optional)
 * param clientUri: clientUri (Optional)
 * param logoUri: logoUri (Optional)
 * param tosUri: tosUri (Optional)
 * param policyUri: policyUri (Optional)
 * param scope: scope (Optional)
 * param jwksUri: jwksUri (Optional)
 * param jwks: jwks (Optional)
 * param contacts: contacts (Optional)
 * param extension: extension (Optional)
 * param accessTokenValiditySeconds: accessTokenValiditySeconds (Optional)
 * param refreshTokenValiditySeconds: refreshTokenValiditySeconds (Optional)
 * param multiTenant: multiTenant (Optional)
 * param secret: secret (Optional)
 * param userType: userType (Optional)
 */
func NewUpdateAppRequestWithAllParams(
    regionId string,
    clientId string,
    clientName *string,
    tokenEndpointAuthMethod *string,
    grantTypes *string,
    redirectUris *string,
    clientUri *string,
    logoUri *string,
    tosUri *string,
    policyUri *string,
    scope *string,
    jwksUri *string,
    jwks *string,
    contacts *string,
    extension *string,
    accessTokenValiditySeconds *int,
    refreshTokenValiditySeconds *int,
    multiTenant *bool,
    secret *string,
    userType *string,
) *UpdateAppRequest {

    return &UpdateAppRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/app/{clientId}",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        ClientId: clientId,
        ClientName: clientName,
        TokenEndpointAuthMethod: tokenEndpointAuthMethod,
        GrantTypes: grantTypes,
        RedirectUris: redirectUris,
        ClientUri: clientUri,
        LogoUri: logoUri,
        TosUri: tosUri,
        PolicyUri: policyUri,
        Scope: scope,
        JwksUri: jwksUri,
        Jwks: jwks,
        Contacts: contacts,
        Extension: extension,
        AccessTokenValiditySeconds: accessTokenValiditySeconds,
        RefreshTokenValiditySeconds: refreshTokenValiditySeconds,
        MultiTenant: multiTenant,
        Secret: secret,
        UserType: userType,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewUpdateAppRequestWithoutParam() *UpdateAppRequest {

    return &UpdateAppRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/app/{clientId}",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: (Required) */
func (r *UpdateAppRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param clientId: (Required) */
func (r *UpdateAppRequest) SetClientId(clientId string) {
    r.ClientId = clientId
}

/* param clientName: 应用名(Optional) */
func (r *UpdateAppRequest) SetClientName(clientName string) {
    r.ClientName = &clientName
}

/* param tokenEndpointAuthMethod: tokenEndpointAuthMethod(Optional) */
func (r *UpdateAppRequest) SetTokenEndpointAuthMethod(tokenEndpointAuthMethod string) {
    r.TokenEndpointAuthMethod = &tokenEndpointAuthMethod
}

/* param grantTypes: grantTypes(Optional) */
func (r *UpdateAppRequest) SetGrantTypes(grantTypes string) {
    r.GrantTypes = &grantTypes
}

/* param redirectUris: redirectUris(Optional) */
func (r *UpdateAppRequest) SetRedirectUris(redirectUris string) {
    r.RedirectUris = &redirectUris
}

/* param clientUri: clientUri(Optional) */
func (r *UpdateAppRequest) SetClientUri(clientUri string) {
    r.ClientUri = &clientUri
}

/* param logoUri: logoUri(Optional) */
func (r *UpdateAppRequest) SetLogoUri(logoUri string) {
    r.LogoUri = &logoUri
}

/* param tosUri: tosUri(Optional) */
func (r *UpdateAppRequest) SetTosUri(tosUri string) {
    r.TosUri = &tosUri
}

/* param policyUri: policyUri(Optional) */
func (r *UpdateAppRequest) SetPolicyUri(policyUri string) {
    r.PolicyUri = &policyUri
}

/* param scope: scope(Optional) */
func (r *UpdateAppRequest) SetScope(scope string) {
    r.Scope = &scope
}

/* param jwksUri: jwksUri(Optional) */
func (r *UpdateAppRequest) SetJwksUri(jwksUri string) {
    r.JwksUri = &jwksUri
}

/* param jwks: jwks(Optional) */
func (r *UpdateAppRequest) SetJwks(jwks string) {
    r.Jwks = &jwks
}

/* param contacts: contacts(Optional) */
func (r *UpdateAppRequest) SetContacts(contacts string) {
    r.Contacts = &contacts
}

/* param extension: extension(Optional) */
func (r *UpdateAppRequest) SetExtension(extension string) {
    r.Extension = &extension
}

/* param accessTokenValiditySeconds: accessTokenValiditySeconds(Optional) */
func (r *UpdateAppRequest) SetAccessTokenValiditySeconds(accessTokenValiditySeconds int) {
    r.AccessTokenValiditySeconds = &accessTokenValiditySeconds
}

/* param refreshTokenValiditySeconds: refreshTokenValiditySeconds(Optional) */
func (r *UpdateAppRequest) SetRefreshTokenValiditySeconds(refreshTokenValiditySeconds int) {
    r.RefreshTokenValiditySeconds = &refreshTokenValiditySeconds
}

/* param multiTenant: multiTenant(Optional) */
func (r *UpdateAppRequest) SetMultiTenant(multiTenant bool) {
    r.MultiTenant = &multiTenant
}

/* param secret: secret(Optional) */
func (r *UpdateAppRequest) SetSecret(secret string) {
    r.Secret = &secret
}

/* param userType: userType(Optional) */
func (r *UpdateAppRequest) SetUserType(userType string) {
    r.UserType = &userType
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r UpdateAppRequest) GetRegionId() string {
    return r.RegionId
}

type UpdateAppResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result UpdateAppResult `json:"result"`
}

type UpdateAppResult struct {
    ClientId string `json:"clientId"`
    ClientName string `json:"clientName"`
    TokenEndpointAuthMethod string `json:"tokenEndpointAuthMethod"`
    GrantTypes string `json:"grantTypes"`
    ResponseTypes string `json:"responseTypes"`
    RedirectUris string `json:"redirectUris"`
    ClientUri string `json:"clientUri"`
    LogoUri string `json:"logoUri"`
    TosUri string `json:"tosUri"`
    PolicyUri string `json:"policyUri"`
    Scope string `json:"scope"`
    JwksUri string `json:"jwksUri"`
    Jwks string `json:"jwks"`
    Contacts string `json:"contacts"`
    Extension string `json:"extension"`
    AccessTokenValiditySeconds int `json:"accessTokenValiditySeconds"`
    RefreshTokenValiditySeconds int `json:"refreshTokenValiditySeconds"`
    MultiTenant bool `json:"multiTenant"`
    SecretUpdateTime int64 `json:"secretUpdateTime"`
    UpdateTime int64 `json:"updateTime"`
    CreateTime int64 `json:"createTime"`
    Account string `json:"account"`
    UserType string `json:"userType"`
}