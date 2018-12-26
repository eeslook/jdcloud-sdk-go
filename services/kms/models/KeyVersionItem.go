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

package models


type KeyVersionItem struct {

    /* 版本号  */
    KeyVersion string `json:"keyVersion"`

    /* 创建时间，采用ISO8601标准，格式为: YYYY-MM-DDTHH:mm:ssZ  */
    CreateTime string `json:"createTime"`

    /* Key当前状态: 0:已启用、1:已禁用、2:计划删除  */
    KeyStatus int `json:"keyStatus"`

    /* 计划删除的时间，采用ISO8601标准，格式为: YYYY-MM-DDTHH:mm:ssZ  */
    DeleteTime string `json:"deleteTime"`
}