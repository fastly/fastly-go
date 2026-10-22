# ArcProviderConnectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The provider name. Must be `bedrock` for `aws-iam` connections. | 
**Models** | **[]string** | The Bedrock model identifiers this connection is allowed to use. | 
**BaseUrl** | **string** | The base URL of the provider&#39;s API. Must be an absolute `http` or `https` URL and must not end in `/v1`. | 
**ApiKey** | **string** | The provider&#39;s secret API key. | 
**AuthType** | **string** | The authentication type for this connection. | 
**CustomerRoleArn** | **string** | The ARN of the IAM role in your AWS account that AI Runtime Control assumes to access Bedrock. See [get Bedrock setup information](https://www.fastly.com/documentation/reference/api/ai-runtime-control/providers#arc-get-bedrock-setup) for the role ARN and external ID your trust policy must reference. | 
**Region** | **string** | The AWS region for Bedrock API requests. | 

## Methods

### NewArcProviderConnectionRequest

`func NewArcProviderConnectionRequest(name string, models []string, baseUrl string, apiKey string, authType string, customerRoleArn string, region string, ) *ArcProviderConnectionRequest`

NewArcProviderConnectionRequest instantiates a new ArcProviderConnectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArcProviderConnectionRequestWithDefaults

`func NewArcProviderConnectionRequestWithDefaults() *ArcProviderConnectionRequest`

NewArcProviderConnectionRequestWithDefaults instantiates a new ArcProviderConnectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ArcProviderConnectionRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ArcProviderConnectionRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ArcProviderConnectionRequest) SetName(v string)`

SetName sets Name field to given value.


### GetModels

`func (o *ArcProviderConnectionRequest) GetModels() []string`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *ArcProviderConnectionRequest) GetModelsOk() (*[]string, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *ArcProviderConnectionRequest) SetModels(v []string)`

SetModels sets Models field to given value.


### GetBaseUrl

`func (o *ArcProviderConnectionRequest) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *ArcProviderConnectionRequest) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *ArcProviderConnectionRequest) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.


### GetApiKey

`func (o *ArcProviderConnectionRequest) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *ArcProviderConnectionRequest) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *ArcProviderConnectionRequest) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.


### GetAuthType

`func (o *ArcProviderConnectionRequest) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *ArcProviderConnectionRequest) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *ArcProviderConnectionRequest) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.


### GetCustomerRoleArn

`func (o *ArcProviderConnectionRequest) GetCustomerRoleArn() string`

GetCustomerRoleArn returns the CustomerRoleArn field if non-nil, zero value otherwise.

### GetCustomerRoleArnOk

`func (o *ArcProviderConnectionRequest) GetCustomerRoleArnOk() (*string, bool)`

GetCustomerRoleArnOk returns a tuple with the CustomerRoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerRoleArn

`func (o *ArcProviderConnectionRequest) SetCustomerRoleArn(v string)`

SetCustomerRoleArn sets CustomerRoleArn field to given value.


### GetRegion

`func (o *ArcProviderConnectionRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ArcProviderConnectionRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ArcProviderConnectionRequest) SetRegion(v string)`

SetRegion sets Region field to given value.



[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


