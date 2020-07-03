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

package eci

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// UpdateContainerGroup invokes the eci.UpdateContainerGroup API synchronously
// api document: https://help.aliyun.com/api/eci/updatecontainergroup.html
func (client *Client) UpdateContainerGroup(request *UpdateContainerGroupRequest) (response *UpdateContainerGroupResponse, err error) {
	response = CreateUpdateContainerGroupResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateContainerGroupWithChan invokes the eci.UpdateContainerGroup API asynchronously
// api document: https://help.aliyun.com/api/eci/updatecontainergroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateContainerGroupWithChan(request *UpdateContainerGroupRequest) (<-chan *UpdateContainerGroupResponse, <-chan error) {
	responseChan := make(chan *UpdateContainerGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateContainerGroup(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// UpdateContainerGroupWithCallback invokes the eci.UpdateContainerGroup API asynchronously
// api document: https://help.aliyun.com/api/eci/updatecontainergroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateContainerGroupWithCallback(request *UpdateContainerGroupRequest, callback func(response *UpdateContainerGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateContainerGroupResponse
		var err error
		defer close(result)
		response, err = client.UpdateContainerGroup(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// UpdateContainerGroupRequest is the request struct for api UpdateContainerGroup
type UpdateContainerGroupRequest struct {
	*requests.RpcRequest
	OwnerId                 requests.Integer                               `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount    string                                         `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId         requests.Integer                               `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount            string                                         `position:"Query" name:"OwnerAccount"`
	RegionId                string                                         `position:"Query" name:"RegionId"`
	ContainerGroupId        string                                         `position:"Query" name:"ContainerGroupId"`
	RestartPolicy           string                                         `position:"Query" name:"RestartPolicy"`
	Tag                     *[]UpdateContainerGroupTag                     `position:"Query" name:"Tag" type:"Repeated"`
	Volume                  *[]UpdateContainerGroupVolume                  `position:"Query" name:"Volume" type:"Repeated"`
	Container               *[]UpdateContainerGroupContainer               `position:"Query" name:"Container" type:"Repeated"`
	InitContainer           *[]UpdateContainerGroupInitContainer           `position:"Query" name:"InitContainer" type:"Repeated"`
	ImageRegistryCredential *[]UpdateContainerGroupImageRegistryCredential `position:"Query" name:"ImageRegistryCredential" type:"Repeated"`
	ClientToken             string                                         `position:"Query" name:"ClientToken"`
	Cpu                     requests.Float                                 `position:"Query" name:"Cpu"`
	Memory                  requests.Float                                 `position:"Query" name:"Memory"`
	DnsConfig               UpdateContainerGroupDnsConfig                  `position:"Query" name:"DnsConfig" type:"Struct"`
}

type UpdateContainerGroupTag struct {
	Key   string `name:"Key"`
	Value string `name:"Value"`
}

type UpdateContainerGroupVolume struct {
	Name             string                               `name:"Name"`
	Type             string                               `name:"Type"`
	NFSVolume        UpdateContainerGroupNFSVolume        `name:"NFSVolume" type:"Struct"`
	ConfigFileVolume UpdateContainerGroupConfigFileVolume `name:"ConfigFileVolume" type:"Struct"`
	EmptyDirVolume   UpdateContainerGroupEmptyDirVolume   `name:"EmptyDirVolume" type:"Struct"`
}

type UpdateContainerGroupContainer struct {
	Name            string                              `name:"Name"`
	Image           string                              `name:"Image"`
	Cpu             requests.Float                      `name:"Cpu"`
	Memory          requests.Float                      `name:"Memory"`
	WorkingDir      string                              `name:"WorkingDir"`
	ImagePullPolicy string                              `name:"ImagePullPolicy"`
	Stdin           requests.Boolean                    `name:"Stdin"`
	StdinOnce       requests.Boolean                    `name:"StdinOnce"`
	Tty             requests.Boolean                    `name:"Tty"`
	Command         []string                            `name:"Command" type:"Repeated"`
	Arg             []string                            `name:"Arg" type:"Repeated"`
	EnvironmentVar  []string                            `name:"EnvironmentVar" type:"Repeated"`
	Port            []string                            `name:"Port" type:"Repeated"`
	VolumeMount     []string                            `name:"VolumeMount" type:"Repeated"`
	Gpu             requests.Integer                    `name:"Gpu"`
	ReadinessProbe  UpdateContainerGroupReadinessProbe  `name:"ReadinessProbe" type:"Struct"`
	LivenessProbe   UpdateContainerGroupLivenessProbe   `name:"LivenessProbe" type:"Struct"`
	SecurityContext UpdateContainerGroupSecurityContext `name:"SecurityContext" type:"Struct"`
}

type UpdateContainerGroupInitContainer struct {
	Name            string                                `name:"Name"`
	Image           string                                `name:"Image"`
	Cpu             requests.Float                        `name:"Cpu"`
	Memory          requests.Float                        `name:"Memory"`
	WorkingDir      string                                `name:"WorkingDir"`
	ImagePullPolicy string                                `name:"ImagePullPolicy"`
	Stdin           requests.Boolean                      `name:"Stdin"`
	StdinOnce       requests.Boolean                      `name:"StdinOnce"`
	Tty             requests.Boolean                      `name:"Tty"`
	Command         []string                              `name:"Command" type:"Repeated"`
	Arg             []string                              `name:"Arg" type:"Repeated"`
	EnvironmentVar  *[]UpdateContainerGroupEnvironmentVar `name:"EnvironmentVar" type:"Repeated"`
	Port            *[]UpdateContainerGroupPort           `name:"Port" type:"Repeated"`
	VolumeMount     *[]UpdateContainerGroupVolumeMount    `name:"VolumeMount" type:"Repeated"`
	Gpu             requests.Integer                      `name:"Gpu"`
	SecurityContext UpdateContainerGroupSecurityContext   `name:"SecurityContext" type:"Struct"`
}

type UpdateContainerGroupImageRegistryCredential struct {
	Server   string `name:"Server"`
	UserName string `name:"UserName"`
	Password string `name:"Password"`
}

type UpdateContainerGroupDnsConfig struct {
	NameServer []string                      `name:"NameServer"`
	Search     []string                      `name:"Search"`
	Option     *[]UpdateContainerGroupOption `name:"Option"`
}

type UpdateContainerGroupNFSVolume struct {
	Server   string           `name:"Server"`
	Path     string           `name:"Path"`
	ReadOnly requests.Boolean `name:"ReadOnly"`
}

type UpdateContainerGroupConfigFileVolume struct {
	ConfigFileToPath *[]UpdateContainerGroupConfigFileToPath `name:"ConfigFileToPath"`
}

type UpdateContainerGroupConfigFileToPath struct {
	Content string `name:"Content"`
	Path    string `name:"Path"`
}

type UpdateContainerGroupEmptyDirVolume struct {
	Medium string `name:"Medium"`
}

type UpdateContainerGroupReadinessProbe struct {
	InitialDelaySeconds requests.Integer              `name:"InitialDelaySeconds"`
	PeriodSeconds       requests.Integer              `name:"PeriodSeconds"`
	SuccessThreshold    requests.Integer              `name:"SuccessThreshold"`
	FailureThreshold    requests.Integer              `name:"FailureThreshold"`
	TimeoutSeconds      requests.Integer              `name:"TimeoutSeconds"`
	TcpSocket           UpdateContainerGroupTcpSocket `name:"TcpSocket"`
	Exec                UpdateContainerGroupExec      `name:"Exec"`
	HttpGet             UpdateContainerGroupHttpGet   `name:"HttpGet"`
}

type UpdateContainerGroupTcpSocket struct {
	Port requests.Integer `name:"Port"`
}

type UpdateContainerGroupExec struct {
	Command []string `name:"Command"`
}

type UpdateContainerGroupHttpGet struct {
	Path   string           `name:"Path"`
	Port   requests.Integer `name:"Port"`
	Scheme string           `name:"Scheme"`
}

type UpdateContainerGroupLivenessProbe struct {
	InitialDelaySeconds requests.Integer              `name:"InitialDelaySeconds"`
	PeriodSeconds       requests.Integer              `name:"PeriodSeconds"`
	SuccessThreshold    requests.Integer              `name:"SuccessThreshold"`
	FailureThreshold    requests.Integer              `name:"FailureThreshold"`
	TimeoutSeconds      requests.Integer              `name:"TimeoutSeconds"`
	TcpSocket           UpdateContainerGroupTcpSocket `name:"TcpSocket"`
	Exec                UpdateContainerGroupExec      `name:"Exec"`
	HttpGet             UpdateContainerGroupHttpGet   `name:"HttpGet"`
}

type UpdateContainerGroupSecurityContext struct {
	ReadOnlyRootFilesystem requests.Boolean               `name:"ReadOnlyRootFilesystem"`
	RunAsUser              requests.Integer               `name:"RunAsUser"`
	Capability             UpdateContainerGroupCapability `name:"Capability"`
}

type UpdateContainerGroupCapability struct {
	Add []string `name:"Add"`
}

type UpdateContainerGroupEnvironmentVar struct {
	Key   string `name:"Key"`
	Value string `name:"Value"`
}

type UpdateContainerGroupPort struct {
	Port     requests.Integer `name:"Port"`
	Protocol string           `name:"Protocol"`
}

type UpdateContainerGroupVolumeMount struct {
	Name      string           `name:"Name"`
	MountPath string           `name:"MountPath"`
	SubPath   string           `name:"SubPath"`
	ReadOnly  requests.Boolean `name:"ReadOnly"`
}

type UpdateContainerGroupOption struct {
	Name  string `name:"Name"`
	Value string `name:"Value"`
}

// UpdateContainerGroupResponse is the response struct for api UpdateContainerGroup
type UpdateContainerGroupResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateContainerGroupRequest creates a request to invoke UpdateContainerGroup API
func CreateUpdateContainerGroupRequest() (request *UpdateContainerGroupRequest) {
	request = &UpdateContainerGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Eci", "2018-08-08", "UpdateContainerGroup", "eci", "openAPI")
	return
}

// CreateUpdateContainerGroupResponse creates a response to parse from UpdateContainerGroup response
func CreateUpdateContainerGroupResponse() (response *UpdateContainerGroupResponse) {
	response = &UpdateContainerGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}