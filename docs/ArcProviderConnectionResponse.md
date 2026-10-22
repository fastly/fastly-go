# ArcProviderConnectionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier of the provider connection. | [readonly] 
**SecretId** | **string** | An opaque identifier for the secret material backing this connection. Not the secret itself. | [readonly] 
**Name** | **string** | The display name of the provider connection. | 
**Models** | **[]string** | The AI model identifiers this connection is allowed to use. | 
**BaseUrl** | Pointer to **string** | The base URL of the provider&#39;s API. Present only for `api-key` connections. | [optional] 
**AuthType** | **string** | The authentication type for this connection. | 
**Region** | Pointer to **string** | The AWS region for Bedrock API requests. Present only for `aws-iam` connections. | [optional] 
**CreatedAt** | **time.Time** | The date and time the connection was created, in RFC 3339 format. | [readonly] 
**UpdatedAt** | **time.Time** | The date and time the connection was last updated, in RFC 3339 format. | [readonly] 

## Methods

### NewArcProviderConnectionResponse

`func NewArcProviderConnectionResponse(id string, secretId string, name string, models []string, authType string, createdAt time.Time, updatedAt time.Time, ) *ArcProviderConnectionResponse`

NewArcProviderConnectionResponse instantiates a new ArcProviderConnectionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArcProviderConnectionResponseWithDefaults

`func NewArcProviderConnectionResponseWithDefaults() *ArcProviderConnectionResponse`

NewArcProviderConnectionResponseWithDefaults instantiates a new ArcProviderConnectionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ArcProviderConnectionResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArcProviderConnectionResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArcProviderConnectionResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSecretId

`func (o *ArcProviderConnectionResponse) GetSecretId() string`

GetSecretId returns the SecretId field if non-nil, zero value otherwise.

### GetSecretIdOk

`func (o *ArcProviderConnectionResponse) GetSecretIdOk() (*string, bool)`

GetSecretIdOk returns a tuple with the SecretId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretId

`func (o *ArcProviderConnectionResponse) SetSecretId(v string)`

SetSecretId sets SecretId field to given value.


### GetName

`func (o *ArcProviderConnectionResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ArcProviderConnectionResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ArcProviderConnectionResponse) SetName(v string)`

SetName sets Name field to given value.


### GetModels

`func (o *ArcProviderConnectionResponse) GetModels() []string`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *ArcProviderConnectionResponse) GetModelsOk() (*[]string, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *ArcProviderConnectionResponse) SetModels(v []string)`

SetModels sets Models field to given value.


### GetBaseUrl

`func (o *ArcProviderConnectionResponse) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *ArcProviderConnectionResponse) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *ArcProviderConnectionResponse) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.

### HasBaseUrl

`func (o *ArcProviderConnectionResponse) HasBaseUrl() bool`

HasBaseUrl returns a boolean if a field has been set.

### GetAuthType

`func (o *ArcProviderConnectionResponse) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *ArcProviderConnectionResponse) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *ArcProviderConnectionResponse) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.


### GetRegion

`func (o *ArcProviderConnectionResponse) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ArcProviderConnectionResponse) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ArcProviderConnectionResponse) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *ArcProviderConnectionResponse) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ArcProviderConnectionResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ArcProviderConnectionResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ArcProviderConnectionResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ArcProviderConnectionResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ArcProviderConnectionResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ArcProviderConnectionResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


