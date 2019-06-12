package gpdb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// RestartDBInstance invokes the gpdb.RestartDBInstance API synchronously
// api document: https://help.aliyun.com/api/gpdb/restartdbinstance.html
func (client *Client) RestartDBInstance(request *RestartDBInstanceRequest) (response *RestartDBInstanceResponse, err error) {
	response = CreateRestartDBInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// RestartDBInstanceWithChan invokes the gpdb.RestartDBInstance API asynchronously
// api document: https://help.aliyun.com/api/gpdb/restartdbinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RestartDBInstanceWithChan(request *RestartDBInstanceRequest) (<-chan *RestartDBInstanceResponse, <-chan error) {
	responseChan := make(chan *RestartDBInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RestartDBInstance(request)
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

// RestartDBInstanceWithCallback invokes the gpdb.RestartDBInstance API asynchronously
// api document: https://help.aliyun.com/api/gpdb/restartdbinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RestartDBInstanceWithCallback(request *RestartDBInstanceRequest, callback func(response *RestartDBInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RestartDBInstanceResponse
		var err error
		defer close(result)
		response, err = client.RestartDBInstance(request)
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

// RestartDBInstanceRequest is the request struct for api RestartDBInstance
type RestartDBInstanceRequest struct {
	*requests.RpcRequest
	ClientToken  string `position:"Query" name:"ClientToken"`
	DBInstanceId string `position:"Query" name:"DBInstanceId"`
}

// RestartDBInstanceResponse is the response struct for api RestartDBInstance
type RestartDBInstanceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateRestartDBInstanceRequest creates a request to invoke RestartDBInstance API
func CreateRestartDBInstanceRequest() (request *RestartDBInstanceRequest) {
	request = &RestartDBInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("gpdb", "2016-05-03", "RestartDBInstance", "gpdb", "openAPI")
	return
}

// CreateRestartDBInstanceResponse creates a response to parse from RestartDBInstance response
func CreateRestartDBInstanceResponse() (response *RestartDBInstanceResponse) {
	response = &RestartDBInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}