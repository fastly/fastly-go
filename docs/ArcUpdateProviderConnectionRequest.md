# ArcUpdateProviderConnectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Models** | Pointer to **[]string** | The AI model identifiers this connection is allowed to use. Must be non-empty if included. | [optional] 
**BaseUrl** | Pointer to **string** | The base URL of the provider&#39;s API (`api-key` connections only). | [optional] 
**ApiKey** | Pointer to **string** | The provider&#39;s secret API key (`api-key` connections only). | [optional] 
**CustomerRoleArn** | Pointer to **string** | The ARN of the IAM role to assume for Bedrock requests (`aws-iam` connections only). | [optional] 
**Region** | Pointer to **string** | The AWS region for Bedrock API requests (`aws-iam` connections only). | [optional] 

## Methods

### NewArcUpdateProviderConnectionRequest

`func NewArcUpdateProviderConnectionRequest() *ArcUpdateProviderConnectionRequest`

NewArcUpdateProviderConnectionRequest instantiates a new ArcUpdateProviderConnectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArcUpdateProviderConnectionRequestWithDefaults

`func NewArcUpdateProviderConnectionRequestWithDefaults() *ArcUpdateProviderConnectionRequest`

NewArcUpdateProviderConnectionRequestWithDefaults instantiates a new ArcUpdateProviderConnectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModels

`func (o *ArcUpdateProviderConnectionRequest) GetModels() []string`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *ArcUpdateProviderConnectionRequest) GetModelsOk() (*[]string, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *ArcUpdateProviderConnectionRequest) SetModels(v []string)`

SetModels sets Models field to given value.

### HasModels

`func (o *ArcUpdateProviderConnectionRequest) HasModels() bool`

HasModels returns a boolean if a field has been set.

### GetBaseUrl

`func (o *ArcUpdateProviderConnectionRequest) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *ArcUpdateProviderConnectionRequest) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *ArcUpdateProviderConnectionRequest) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.

### HasBaseUrl

`func (o *ArcUpdateProviderConnectionRequest) HasBaseUrl() bool`

HasBaseUrl returns a boolean if a field has been set.

### GetApiKey

`func (o *ArcUpdateProviderConnectionRequest) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *ArcUpdateProviderConnectionRequest) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *ArcUpdateProviderConnectionRequest) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *ArcUpdateProviderConnectionRequest) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetCustomerRoleArn

`func (o *ArcUpdateProviderConnectionRequest) GetCustomerRoleArn() string`

GetCustomerRoleArn returns the CustomerRoleArn field if non-nil, zero value otherwise.

### GetCustomerRoleArnOk

`func (o *ArcUpdateProviderConnectionRequest) GetCustomerRoleArnOk() (*string, bool)`

GetCustomerRoleArnOk returns a tuple with the CustomerRoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerRoleArn

`func (o *ArcUpdateProviderConnectionRequest) SetCustomerRoleArn(v string)`

SetCustomerRoleArn sets CustomerRoleArn field to given value.

### HasCustomerRoleArn

`func (o *ArcUpdateProviderConnectionRequest) HasCustomerRoleArn() bool`

HasCustomerRoleArn returns a boolean if a field has been set.

### GetRegion

`func (o *ArcUpdateProviderConnectionRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ArcUpdateProviderConnectionRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ArcUpdateProviderConnectionRequest) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *ArcUpdateProviderConnectionRequest) HasRegion() bool`

HasRegion returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


