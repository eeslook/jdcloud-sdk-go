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


type DwDatabase struct {

    /* 数据库id (Optional) */
    Id int `json:"id"`

    /* 用户名 (Optional) */
    UserName string `json:"userName"`

    /* 数据库名称  */
    DatabaseName string `json:"databaseName"`

    /* 总表数量 (Optional) */
    TotalTableQuantity int `json:"totalTableQuantity"`

    /* 创建时间 (Optional) */
    CreateTime string `json:"createTime"`

    /* 最新更新时间 (Optional) */
    LastUpdateTime string `json:"lastUpdateTime"`

    /* 最新更新时间 (Optional) */
    PhysicalStorageCapacity string `json:"physicalStorageCapacity"`

    /* 类别 (Optional) */
    Category string `json:"category"`

    /* 来源 (Optional) */
    Source string `json:"source"`

    /* 所有者 (Optional) */
    Owner string `json:"owner"`

    /* 位置 (Optional) */
    Location string `json:"location"`

    /* 描述信息 (Optional) */
    Comments string `json:"comments"`
}
