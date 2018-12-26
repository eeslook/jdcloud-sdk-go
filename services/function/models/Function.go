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


type Function struct {

    /*  (Optional) */
    FunctionId string `json:"functionId"`

    /*  (Optional) */
    Name string `json:"name"`

    /*  (Optional) */
    Description string `json:"description"`

    /*  (Optional) */
    Entrance string `json:"entrance"`

    /*  (Optional) */
    Memory int `json:"memory"`

    /*  (Optional) */
    RunTime string `json:"runTime"`

    /*  (Optional) */
    OverTime int `json:"overTime"`

    /*  (Optional) */
    Version string `json:"version"`

    /*  (Optional) */
    Code Code `json:"code"`

    /*  (Optional) */
    Environment Env `json:"environment"`

    /*  (Optional) */
    LogSetId string `json:"logSetId"`

    /*  (Optional) */
    LogTopicId string `json:"logTopicId"`

    /*  (Optional) */
    CodeCheckSum string `json:"codeCheckSum"`

    /*  (Optional) */
    CodeSize int `json:"codeSize"`

    /*  (Optional) */
    DownloadUrl string `json:"downloadUrl"`
}