package ens

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// EPNInstance is a nested struct in ens response
type EPNInstance struct {
	PrivateIpAddress        string `json:"PrivateIpAddress" xml:"PrivateIpAddress"`
	EPNInstanceName         string `json:"EPNInstanceName" xml:"EPNInstanceName"`
	ModifyTime              string `json:"ModifyTime" xml:"ModifyTime"`
	VSwitchId               string `json:"VSwitchId" xml:"VSwitchId"`
	InstanceName            string `json:"InstanceName" xml:"InstanceName"`
	VSwitchName             string `json:"VSwitchName" xml:"VSwitchName"`
	EnsRegionId             string `json:"EnsRegionId" xml:"EnsRegionId"`
	CreationTime            string `json:"CreationTime" xml:"CreationTime"`
	EPNInstanceType         string `json:"EPNInstanceType" xml:"EPNInstanceType"`
	PublicIpAddress         string `json:"PublicIpAddress" xml:"PublicIpAddress"`
	NetworkingModel         string `json:"NetworkingModel" xml:"NetworkingModel"`
	InternetMaxBandwidthOut int    `json:"InternetMaxBandwidthOut" xml:"InternetMaxBandwidthOut"`
	StartTime               string `json:"StartTime" xml:"StartTime"`
	EndTime                 string `json:"EndTime" xml:"EndTime"`
	InstanceId              string `json:"InstanceId" xml:"InstanceId"`
	Status                  string `json:"Status" xml:"Status"`
	Isp                     string `json:"Isp" xml:"Isp"`
	CidrBlock               string `json:"CidrBlock" xml:"CidrBlock"`
	EPNInstanceId           string `json:"EPNInstanceId" xml:"EPNInstanceId"`
}