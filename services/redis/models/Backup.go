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


type Backup struct {

    /* 备份操作ID  */
    BaseId string `json:"baseId"`

    /* 备份文件的名称  */
    BackupFileName string `json:"backupFileName"`

    /* 备份实例ID  */
    SpaceId string `json:"spaceId"`

    /* 备份开始时间  */
    BackupStartTime string `json:"backupStartTime"`

    /* 备份结束时间  */
    BackupEndTime string `json:"backupEndTime"`

    /* 备份类型，1表示手动备份，0表示自动备份  */
    BackupType int `json:"backupType"`

    /* 备份文件大小，如果实例是集群版，则表示每个分片备份文件大小的总和  */
    BackupSize int `json:"backupSize"`

    /* 备份任务状态状态，1表示失败，2表示成功  */
    BackupStatus int `json:"backupStatus"`

    /* 备份文件下载的URL地址，集群版有多个URL地址  */
    BackupDownloadURL string `json:"backupDownloadURL"`
}
