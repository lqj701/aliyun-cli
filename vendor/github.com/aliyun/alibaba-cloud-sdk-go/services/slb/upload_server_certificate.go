package slb

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

// UploadServerCertificate invokes the slb.UploadServerCertificate API synchronously
// api document: https://help.aliyun.com/api/slb/uploadservercertificate.html
func (client *Client) UploadServerCertificate(request *UploadServerCertificateRequest) (response *UploadServerCertificateResponse, err error) {
	response = CreateUploadServerCertificateResponse()
	err = client.DoAction(request, response)
	return
}

// UploadServerCertificateWithChan invokes the slb.UploadServerCertificate API asynchronously
// api document: https://help.aliyun.com/api/slb/uploadservercertificate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UploadServerCertificateWithChan(request *UploadServerCertificateRequest) (<-chan *UploadServerCertificateResponse, <-chan error) {
	responseChan := make(chan *UploadServerCertificateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UploadServerCertificate(request)
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

// UploadServerCertificateWithCallback invokes the slb.UploadServerCertificate API asynchronously
// api document: https://help.aliyun.com/api/slb/uploadservercertificate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UploadServerCertificateWithCallback(request *UploadServerCertificateRequest, callback func(response *UploadServerCertificateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UploadServerCertificateResponse
		var err error
		defer close(result)
		response, err = client.UploadServerCertificate(request)
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

// UploadServerCertificateRequest is the request struct for api UploadServerCertificate
type UploadServerCertificateRequest struct {
	*requests.RpcRequest
	OwnerId                 requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount    string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId         requests.Integer `position:"Query" name:"ResourceOwnerId"`
	AliCloudCertificateId   string           `position:"Query" name:"AliCloudCertificateId"`
	AliCloudCertificateName string           `position:"Query" name:"AliCloudCertificateName"`
	ServerCertificate       string           `position:"Query" name:"ServerCertificate"`
	PrivateKey              string           `position:"Query" name:"PrivateKey"`
	ServerCertificateName   string           `position:"Query" name:"ServerCertificateName"`
	OwnerAccount            string           `position:"Query" name:"OwnerAccount"`
	AccessKeyId             string           `position:"Query" name:"access_key_id"`
	Tags                    string           `position:"Query" name:"Tags"`
	ResourceGroupId         string           `position:"Query" name:"ResourceGroupId"`
}

// UploadServerCertificateResponse is the response struct for api UploadServerCertificate
type UploadServerCertificateResponse struct {
	*responses.BaseResponse
	RequestId               string `json:"RequestId" xml:"RequestId"`
	ServerCertificateId     string `json:"ServerCertificateId" xml:"ServerCertificateId"`
	Fingerprint             string `json:"Fingerprint" xml:"Fingerprint"`
	ServerCertificateName   string `json:"ServerCertificateName" xml:"ServerCertificateName"`
	RegionId                string `json:"RegionId" xml:"RegionId"`
	RegionIdAlias           string `json:"RegionIdAlias" xml:"RegionIdAlias"`
	AliCloudCertificateId   string `json:"AliCloudCertificateId" xml:"AliCloudCertificateId"`
	AliCloudCertificateName string `json:"AliCloudCertificateName" xml:"AliCloudCertificateName"`
	IsAliCloudCertificate   int    `json:"IsAliCloudCertificate" xml:"IsAliCloudCertificate"`
	ResourceGroupId         string `json:"ResourceGroupId" xml:"ResourceGroupId"`
	CreateTime              string `json:"CreateTime" xml:"CreateTime"`
	CreateTimeStamp         int    `json:"CreateTimeStamp" xml:"CreateTimeStamp"`
}

// CreateUploadServerCertificateRequest creates a request to invoke UploadServerCertificate API
func CreateUploadServerCertificateRequest() (request *UploadServerCertificateRequest) {
	request = &UploadServerCertificateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "UploadServerCertificate", "slb", "openAPI")
	return
}

// CreateUploadServerCertificateResponse creates a response to parse from UploadServerCertificate response
func CreateUploadServerCertificateResponse() (response *UploadServerCertificateResponse) {
	response = &UploadServerCertificateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
