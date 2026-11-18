# ArcBedrockSetupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArcRoleArn** | **string** | The ARN of the AI Runtime Control IAM role your AWS account&#39;s trust policy must trust. | 
**ExternalID** | **string** | The external ID to require in your trust policy. This is your Fastly customer ID. | 

## Methods

### NewArcBedrockSetupResponse

`func NewArcBedrockSetupResponse(arcRoleArn string, externalID string, ) *ArcBedrockSetupResponse`

NewArcBedrockSetupResponse instantiates a new ArcBedrockSetupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArcBedrockSetupResponseWithDefaults

`func NewArcBedrockSetupResponseWithDefaults() *ArcBedrockSetupResponse`

NewArcBedrockSetupResponseWithDefaults instantiates a new ArcBedrockSetupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArcRoleArn

`func (o *ArcBedrockSetupResponse) GetArcRoleArn() string`

GetArcRoleArn returns the ArcRoleArn field if non-nil, zero value otherwise.

### GetArcRoleArnOk

`func (o *ArcBedrockSetupResponse) GetArcRoleArnOk() (*string, bool)`

GetArcRoleArnOk returns a tuple with the ArcRoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArcRoleArn

`func (o *ArcBedrockSetupResponse) SetArcRoleArn(v string)`

SetArcRoleArn sets ArcRoleArn field to given value.


### GetExternalID

`func (o *ArcBedrockSetupResponse) GetExternalID() string`

GetExternalID returns the ExternalID field if non-nil, zero value otherwise.

### GetExternalIDOk

`func (o *ArcBedrockSetupResponse) GetExternalIDOk() (*string, bool)`

GetExternalIDOk returns a tuple with the ExternalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalID

`func (o *ArcBedrockSetupResponse) SetExternalID(v string)`

SetExternalID sets ExternalID field to given value.



[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
