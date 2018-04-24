package mts

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

// SearchWaterMarkTemplate invokes the mts.SearchWaterMarkTemplate API synchronously
// api document: https://help.aliyun.com/api/mts/searchwatermarktemplate.html
func (client *Client) SearchWaterMarkTemplate(request *SearchWaterMarkTemplateRequest) (response *SearchWaterMarkTemplateResponse, err error) {
	response = CreateSearchWaterMarkTemplateResponse()
	err = client.DoAction(request, response)
	return
}

// SearchWaterMarkTemplateWithChan invokes the mts.SearchWaterMarkTemplate API asynchronously
// api document: https://help.aliyun.com/api/mts/searchwatermarktemplate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SearchWaterMarkTemplateWithChan(request *SearchWaterMarkTemplateRequest) (<-chan *SearchWaterMarkTemplateResponse, <-chan error) {
	responseChan := make(chan *SearchWaterMarkTemplateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SearchWaterMarkTemplate(request)
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

// SearchWaterMarkTemplateWithCallback invokes the mts.SearchWaterMarkTemplate API asynchronously
// api document: https://help.aliyun.com/api/mts/searchwatermarktemplate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SearchWaterMarkTemplateWithCallback(request *SearchWaterMarkTemplateRequest, callback func(response *SearchWaterMarkTemplateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SearchWaterMarkTemplateResponse
		var err error
		defer close(result)
		response, err = client.SearchWaterMarkTemplate(request)
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

// SearchWaterMarkTemplateRequest is the request struct for api SearchWaterMarkTemplate
type SearchWaterMarkTemplateRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	State                string           `position:"Query" name:"State"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
}

// SearchWaterMarkTemplateResponse is the response struct for api SearchWaterMarkTemplate
type SearchWaterMarkTemplateResponse struct {
	*responses.BaseResponse
	RequestId             string                                         `json:"RequestId" xml:"RequestId"`
	TotalCount            int                                            `json:"TotalCount" xml:"TotalCount"`
	PageNumber            int                                            `json:"PageNumber" xml:"PageNumber"`
	PageSize              int                                            `json:"PageSize" xml:"PageSize"`
	WaterMarkTemplateList WaterMarkTemplateListInSearchWaterMarkTemplate `json:"WaterMarkTemplateList" xml:"WaterMarkTemplateList"`
}

// CreateSearchWaterMarkTemplateRequest creates a request to invoke SearchWaterMarkTemplate API
func CreateSearchWaterMarkTemplateRequest() (request *SearchWaterMarkTemplateRequest) {
	request = &SearchWaterMarkTemplateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "SearchWaterMarkTemplate", "mts", "openAPI")
	return
}

// CreateSearchWaterMarkTemplateResponse creates a response to parse from SearchWaterMarkTemplate response
func CreateSearchWaterMarkTemplateResponse() (response *SearchWaterMarkTemplateResponse) {
	response = &SearchWaterMarkTemplateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
