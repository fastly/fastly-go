# ArcProviderConnectionAwsIam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The provider name. Must be `bedrock` for `aws-iam` connections. | 
**Models** | **[]string** | The Bedrock model identifiers this connection is allowed to use. | 
**AuthType** | **string** | The authentication type for this connection. | 
**CustomerRoleArn** | **string** | The ARN of the IAM role in your AWS account that AI Runtime Control assumes to access Bedrock. See [get Bedrock setup information](https://www.fastly.com/documentation/reference/api/ai-runtime-control/providers#arc-get-bedrock-setup) for the role ARN and external ID your trust policy must reference. | 
**Region** | **string** | The AWS region for Bedrock API requests. | 

## Methods

### NewArcProviderConnectionAwsIam

`func NewArcProviderConnectionAwsIam(name string, models []string, authType string, customerRoleArn string, region string, ) *ArcProviderConnectionAwsIam`

NewArcProviderConnectionAwsIam instantiates a new ArcProviderConnectionAwsIam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArcProviderConnectionAwsIamWithDefaults

`func NewArcProviderConnectionAwsIamWithDefaults() *ArcProviderConnectionAwsIam`

NewArcProviderConnectionAwsIamWithDefaults instantiates a new ArcProviderConnectionAwsIam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ArcProviderConnectionAwsIam) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ArcProviderConnectionAwsIam) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ArcProviderConnectionAwsIam) SetName(v string)`

SetName sets Name field to given value.


### GetModels

`func (o *ArcProviderConnectionAwsIam) GetModels() []string`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *ArcProviderConnectionAwsIam) GetModelsOk() (*[]string, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *ArcProviderConnectionAwsIam) SetModels(v []string)`

SetModels sets Models field to given value.


### GetAuthType

`func (o *ArcProviderConnectionAwsIam) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *ArcProviderConnectionAwsIam) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *ArcProviderConnectionAwsIam) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.


### GetCustomerRoleArn

`func (o *ArcProviderConnectionAwsIam) GetCustomerRoleArn() string`

GetCustomerRoleArn returns the CustomerRoleArn field if non-nil, zero value otherwise.

### GetCustomerRoleArnOk

`func (o *ArcProviderConnectionAwsIam) GetCustomerRoleArnOk() (*string, bool)`

GetCustomerRoleArnOk returns a tuple with the CustomerRoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerRoleArn

`func (o *ArcProviderConnectionAwsIam) SetCustomerRoleArn(v string)`

SetCustomerRoleArn sets CustomerRoleArn field to given value.


### GetRegion

`func (o *ArcProviderConnectionAwsIam) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ArcProviderConnectionAwsIam) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ArcProviderConnectionAwsIam) SetRegion(v string)`

SetRegion sets Region field to given value.



[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
