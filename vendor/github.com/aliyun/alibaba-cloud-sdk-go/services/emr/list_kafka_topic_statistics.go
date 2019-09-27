package emr

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

// ListKafkaTopicStatistics invokes the emr.ListKafkaTopicStatistics API synchronously
// api document: https://help.aliyun.com/api/emr/listkafkatopicstatistics.html
func (client *Client) ListKafkaTopicStatistics(request *ListKafkaTopicStatisticsRequest) (response *ListKafkaTopicStatisticsResponse, err error) {
	response = CreateListKafkaTopicStatisticsResponse()
	err = client.DoAction(request, response)
	return
}

// ListKafkaTopicStatisticsWithChan invokes the emr.ListKafkaTopicStatistics API asynchronously
// api document: https://help.aliyun.com/api/emr/listkafkatopicstatistics.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListKafkaTopicStatisticsWithChan(request *ListKafkaTopicStatisticsRequest) (<-chan *ListKafkaTopicStatisticsResponse, <-chan error) {
	responseChan := make(chan *ListKafkaTopicStatisticsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListKafkaTopicStatistics(request)
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

// ListKafkaTopicStatisticsWithCallback invokes the emr.ListKafkaTopicStatistics API asynchronously
// api document: https://help.aliyun.com/api/emr/listkafkatopicstatistics.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListKafkaTopicStatisticsWithCallback(request *ListKafkaTopicStatisticsRequest, callback func(response *ListKafkaTopicStatisticsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListKafkaTopicStatisticsResponse
		var err error
		defer close(result)
		response, err = client.ListKafkaTopicStatistics(request)
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

// ListKafkaTopicStatisticsRequest is the request struct for api ListKafkaTopicStatistics
type ListKafkaTopicStatisticsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ActiveOnly      requests.Boolean `position:"Query" name:"ActiveOnly"`
	PageSize        requests.Integer `position:"Query" name:"PageSize"`
	DataSourceId    string           `position:"Query" name:"DataSourceId"`
	TopicName       string           `position:"Query" name:"TopicName"`
	ClusterId       string           `position:"Query" name:"ClusterId"`
	PageNumber      requests.Integer `position:"Query" name:"PageNumber"`
	FuzzyTopicName  string           `position:"Query" name:"FuzzyTopicName"`
}

// ListKafkaTopicStatisticsResponse is the response struct for api ListKafkaTopicStatistics
type ListKafkaTopicStatisticsResponse struct {
	*responses.BaseResponse
	RequestId  string                              `json:"RequestId" xml:"RequestId"`
	Total      int                                 `json:"Total" xml:"Total"`
	PageSize   int                                 `json:"PageSize" xml:"PageSize"`
	PageNumber int                                 `json:"PageNumber" xml:"PageNumber"`
	TopicList  TopicListInListKafkaTopicStatistics `json:"TopicList" xml:"TopicList"`
}

// CreateListKafkaTopicStatisticsRequest creates a request to invoke ListKafkaTopicStatistics API
func CreateListKafkaTopicStatisticsRequest() (request *ListKafkaTopicStatisticsRequest) {
	request = &ListKafkaTopicStatisticsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "ListKafkaTopicStatistics", "emr", "openAPI")
	return
}

// CreateListKafkaTopicStatisticsResponse creates a response to parse from ListKafkaTopicStatistics response
func CreateListKafkaTopicStatisticsResponse() (response *ListKafkaTopicStatisticsResponse) {
	response = &ListKafkaTopicStatisticsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}