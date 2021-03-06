package rds

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

// ModifyBackupPolicy invokes the rds.ModifyBackupPolicy API synchronously
// api document: https://help.aliyun.com/api/rds/modifybackuppolicy.html
func (client *Client) ModifyBackupPolicy(request *ModifyBackupPolicyRequest) (response *ModifyBackupPolicyResponse, err error) {
	response = CreateModifyBackupPolicyResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyBackupPolicyWithChan invokes the rds.ModifyBackupPolicy API asynchronously
// api document: https://help.aliyun.com/api/rds/modifybackuppolicy.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyBackupPolicyWithChan(request *ModifyBackupPolicyRequest) (<-chan *ModifyBackupPolicyResponse, <-chan error) {
	responseChan := make(chan *ModifyBackupPolicyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyBackupPolicy(request)
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

// ModifyBackupPolicyWithCallback invokes the rds.ModifyBackupPolicy API asynchronously
// api document: https://help.aliyun.com/api/rds/modifybackuppolicy.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyBackupPolicyWithCallback(request *ModifyBackupPolicyRequest, callback func(response *ModifyBackupPolicyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyBackupPolicyResponse
		var err error
		defer close(result)
		response, err = client.ModifyBackupPolicy(request)
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

// ModifyBackupPolicyRequest is the request struct for api ModifyBackupPolicy
type ModifyBackupPolicyRequest struct {
	*requests.RpcRequest
	PreferredBackupPeriod    string           `position:"Query" name:"PreferredBackupPeriod"`
	ResourceOwnerId          requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount     string           `position:"Query" name:"ResourceOwnerAccount"`
	LocalLogRetentionHours   string           `position:"Query" name:"LocalLogRetentionHours"`
	OwnerAccount             string           `position:"Query" name:"OwnerAccount"`
	LogBackupFrequency       string           `position:"Query" name:"LogBackupFrequency"`
	CompressType             string           `position:"Query" name:"CompressType"`
	BackupLog                string           `position:"Query" name:"BackupLog"`
	LocalLogRetentionSpace   string           `position:"Query" name:"LocalLogRetentionSpace"`
	OwnerId                  requests.Integer `position:"Query" name:"OwnerId"`
	Duplication              string           `position:"Query" name:"Duplication"`
	PreferredBackupTime      string           `position:"Query" name:"PreferredBackupTime"`
	BackupRetentionPeriod    string           `position:"Query" name:"BackupRetentionPeriod"`
	DuplicationContent       string           `position:"Query" name:"DuplicationContent"`
	HighSpaceUsageProtection string           `position:"Query" name:"HighSpaceUsageProtection"`
	DBInstanceId             string           `position:"Query" name:"DBInstanceId"`
	DuplicationLocation      string           `position:"Query" name:"DuplicationLocation"`
	LogBackupRetentionPeriod string           `position:"Query" name:"LogBackupRetentionPeriod"`
	EnableBackupLog          string           `position:"Query" name:"EnableBackupLog"`
	BackupPolicyMode         string           `position:"Query" name:"BackupPolicyMode"`
}

// ModifyBackupPolicyResponse is the response struct for api ModifyBackupPolicy
type ModifyBackupPolicyResponse struct {
	*responses.BaseResponse
	RequestId                string `json:"RequestId" xml:"RequestId"`
	DBInstanceID             string `json:"DBInstanceID" xml:"DBInstanceID"`
	EnableBackupLog          string `json:"EnableBackupLog" xml:"EnableBackupLog"`
	LocalLogRetentionHours   int    `json:"LocalLogRetentionHours" xml:"LocalLogRetentionHours"`
	LocalLogRetentionSpace   string `json:"LocalLogRetentionSpace" xml:"LocalLogRetentionSpace"`
	HighSpaceUsageProtection string `json:"HighSpaceUsageProtection" xml:"HighSpaceUsageProtection"`
	CompressType             string `json:"CompressType" xml:"CompressType"`
}

// CreateModifyBackupPolicyRequest creates a request to invoke ModifyBackupPolicy API
func CreateModifyBackupPolicyRequest() (request *ModifyBackupPolicyRequest) {
	request = &ModifyBackupPolicyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "ModifyBackupPolicy", "rds", "openAPI")
	return
}

// CreateModifyBackupPolicyResponse creates a response to parse from ModifyBackupPolicy response
func CreateModifyBackupPolicyResponse() (response *ModifyBackupPolicyResponse) {
	response = &ModifyBackupPolicyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
