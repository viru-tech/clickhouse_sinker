package kms

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

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// ExportDataKey invokes the kms.ExportDataKey API synchronously
func (client *Client) ExportDataKey(request *ExportDataKeyRequest) (response *ExportDataKeyResponse, err error) {
	response = CreateExportDataKeyResponse()
	err = client.DoAction(request, response)
	return
}

// ExportDataKeyWithChan invokes the kms.ExportDataKey API asynchronously
func (client *Client) ExportDataKeyWithChan(request *ExportDataKeyRequest) (<-chan *ExportDataKeyResponse, <-chan error) {
	responseChan := make(chan *ExportDataKeyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ExportDataKey(request)
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

// ExportDataKeyWithCallback invokes the kms.ExportDataKey API asynchronously
func (client *Client) ExportDataKeyWithCallback(request *ExportDataKeyRequest, callback func(response *ExportDataKeyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ExportDataKeyResponse
		var err error
		defer close(result)
		response, err = client.ExportDataKey(request)
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

// ExportDataKeyRequest is the request struct for api ExportDataKey
type ExportDataKeyRequest struct {
	*requests.RpcRequest
	PublicKeyBlob     string `position:"Query" name:"PublicKeyBlob"`
	EncryptionContext string `position:"Query" name:"EncryptionContext"`
	WrappingAlgorithm string `position:"Query" name:"WrappingAlgorithm"`
	CiphertextBlob    string `position:"Query" name:"CiphertextBlob"`
	WrappingKeySpec   string `position:"Query" name:"WrappingKeySpec"`
}

// ExportDataKeyResponse is the response struct for api ExportDataKey
type ExportDataKeyResponse struct {
	*responses.BaseResponse
	KeyVersionId    string `json:"KeyVersionId" xml:"KeyVersionId"`
	KeyId           string `json:"KeyId" xml:"KeyId"`
	RequestId       string `json:"RequestId" xml:"RequestId"`
	ExportedDataKey string `json:"ExportedDataKey" xml:"ExportedDataKey"`
}

// CreateExportDataKeyRequest creates a request to invoke ExportDataKey API
func CreateExportDataKeyRequest() (request *ExportDataKeyRequest) {
	request = &ExportDataKeyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Kms", "2016-01-20", "ExportDataKey", "kms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateExportDataKeyResponse creates a response to parse from ExportDataKey response
func CreateExportDataKeyResponse() (response *ExportDataKeyResponse) {
	response = &ExportDataKeyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
