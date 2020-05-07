package resourcemanager

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

// ListPolicyAttachments invokes the resourcemanager.ListPolicyAttachments API synchronously
// api document: https://help.aliyun.com/api/resourcemanager/listpolicyattachments.html
func (client *Client) ListPolicyAttachments(request *ListPolicyAttachmentsRequest) (response *ListPolicyAttachmentsResponse, err error) {
	response = CreateListPolicyAttachmentsResponse()
	err = client.DoAction(request, response)
	return
}

// ListPolicyAttachmentsWithChan invokes the resourcemanager.ListPolicyAttachments API asynchronously
// api document: https://help.aliyun.com/api/resourcemanager/listpolicyattachments.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListPolicyAttachmentsWithChan(request *ListPolicyAttachmentsRequest) (<-chan *ListPolicyAttachmentsResponse, <-chan error) {
	responseChan := make(chan *ListPolicyAttachmentsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListPolicyAttachments(request)
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

// ListPolicyAttachmentsWithCallback invokes the resourcemanager.ListPolicyAttachments API asynchronously
// api document: https://help.aliyun.com/api/resourcemanager/listpolicyattachments.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListPolicyAttachmentsWithCallback(request *ListPolicyAttachmentsRequest, callback func(response *ListPolicyAttachmentsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListPolicyAttachmentsResponse
		var err error
		defer close(result)
		response, err = client.ListPolicyAttachments(request)
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

// ListPolicyAttachmentsRequest is the request struct for api ListPolicyAttachments
type ListPolicyAttachmentsRequest struct {
	*requests.RpcRequest
	Language        string           `position:"Query" name:"Language"`
	PageNumber      requests.Integer `position:"Query" name:"PageNumber"`
	ResourceGroupId string           `position:"Query" name:"ResourceGroupId"`
	PageSize        requests.Integer `position:"Query" name:"PageSize"`
	PolicyType      string           `position:"Query" name:"PolicyType"`
	PrincipalType   string           `position:"Query" name:"PrincipalType"`
	PolicyName      string           `position:"Query" name:"PolicyName"`
	PrincipalName   string           `position:"Query" name:"PrincipalName"`
}

// ListPolicyAttachmentsResponse is the response struct for api ListPolicyAttachments
type ListPolicyAttachmentsResponse struct {
	*responses.BaseResponse
	RequestId         string            `json:"RequestId" xml:"RequestId"`
	PageNumber        int               `json:"PageNumber" xml:"PageNumber"`
	PageSize          int               `json:"PageSize" xml:"PageSize"`
	TotalCount        int               `json:"TotalCount" xml:"TotalCount"`
	PolicyAttachments PolicyAttachments `json:"PolicyAttachments" xml:"PolicyAttachments"`
}

// CreateListPolicyAttachmentsRequest creates a request to invoke ListPolicyAttachments API
func CreateListPolicyAttachmentsRequest() (request *ListPolicyAttachmentsRequest) {
	request = &ListPolicyAttachmentsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ResourceManager", "2020-03-31", "ListPolicyAttachments", "resourcemanager", "openAPI")
	return
}

// CreateListPolicyAttachmentsResponse creates a response to parse from ListPolicyAttachments response
func CreateListPolicyAttachmentsResponse() (response *ListPolicyAttachmentsResponse) {
	response = &ListPolicyAttachmentsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
