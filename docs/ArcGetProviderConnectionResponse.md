# ArcGetProviderConnectionResponse

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
**VirtualKeyIds** | **[]string** | The IDs of the active virtual keys whose failover policy references this connection. Empty if no virtual key references it. | [readonly] 

## Methods

### NewArcGetProviderConnectionResponse

`func NewArcGetProviderConnectionResponse(id string, secretId string, name string, models []string, authType string, createdAt time.Time, updatedAt time.Time, virtualKeyIds []string, ) *ArcGetProviderConnectionResponse`

NewArcGetProviderConnectionResponse instantiates a new ArcGetProviderConnectionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArcGetProviderConnectionResponseWithDefaults

`func NewArcGetProviderConnectionResponseWithDefaults() *ArcGetProviderConnectionResponse`

NewArcGetProviderConnectionResponseWithDefaults instantiates a new ArcGetProviderConnectionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ArcGetProviderConnectionResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArcGetProviderConnectionResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArcGetProviderConnectionResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSecretId

`func (o *ArcGetProviderConnectionResponse) GetSecretId() string`

GetSecretId returns the SecretId field if non-nil, zero value otherwise.

### GetSecretIdOk

`func (o *ArcGetProviderConnectionResponse) GetSecretIdOk() (*string, bool)`

GetSecretIdOk returns a tuple with the SecretId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretId

`func (o *ArcGetProviderConnectionResponse) SetSecretId(v string)`

SetSecretId sets SecretId field to given value.


### GetName

`func (o *ArcGetProviderConnectionResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ArcGetProviderConnectionResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ArcGetProviderConnectionResponse) SetName(v string)`

SetName sets Name field to given value.


### GetModels

`func (o *ArcGetProviderConnectionResponse) GetModels() []string`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *ArcGetProviderConnectionResponse) GetModelsOk() (*[]string, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *ArcGetProviderConnectionResponse) SetModels(v []string)`

SetModels sets Models field to given value.


### GetBaseUrl

`func (o *ArcGetProviderConnectionResponse) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *ArcGetProviderConnectionResponse) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *ArcGetProviderConnectionResponse) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.

### HasBaseUrl

`func (o *ArcGetProviderConnectionResponse) HasBaseUrl() bool`

HasBaseUrl returns a boolean if a field has been set.

### GetAuthType

`func (o *ArcGetProviderConnectionResponse) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *ArcGetProviderConnectionResponse) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *ArcGetProviderConnectionResponse) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.


### GetRegion

`func (o *ArcGetProviderConnectionResponse) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ArcGetProviderConnectionResponse) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ArcGetProviderConnectionResponse) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *ArcGetProviderConnectionResponse) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ArcGetProviderConnectionResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ArcGetProviderConnectionResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ArcGetProviderConnectionResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ArcGetProviderConnectionResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ArcGetProviderConnectionResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ArcGetProviderConnectionResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetVirtualKeyIds

`func (o *ArcGetProviderConnectionResponse) GetVirtualKeyIds() []string`

GetVirtualKeyIds returns the VirtualKeyIds field if non-nil, zero value otherwise.

### GetVirtualKeyIdsOk

`func (o *ArcGetProviderConnectionResponse) GetVirtualKeyIdsOk() (*[]string, bool)`

GetVirtualKeyIdsOk returns a tuple with the VirtualKeyIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualKeyIds

`func (o *ArcGetProviderConnectionResponse) SetVirtualKeyIds(v []string)`

SetVirtualKeyIds sets VirtualKeyIds field to given value.



[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


