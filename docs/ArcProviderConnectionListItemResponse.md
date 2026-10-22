# ArcProviderConnectionListItemResponse

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
**DeletedAt** | Pointer to **time.Time** | The date and time the connection was deleted, in RFC 3339 format. Present only for deleted connections, which are returned only when `include_deleted` is `true`. | [optional] [readonly] 

## Methods

### NewArcProviderConnectionListItemResponse

`func NewArcProviderConnectionListItemResponse(id string, secretId string, name string, models []string, authType string, createdAt time.Time, updatedAt time.Time, ) *ArcProviderConnectionListItemResponse`

NewArcProviderConnectionListItemResponse instantiates a new ArcProviderConnectionListItemResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArcProviderConnectionListItemResponseWithDefaults

`func NewArcProviderConnectionListItemResponseWithDefaults() *ArcProviderConnectionListItemResponse`

NewArcProviderConnectionListItemResponseWithDefaults instantiates a new ArcProviderConnectionListItemResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ArcProviderConnectionListItemResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArcProviderConnectionListItemResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArcProviderConnectionListItemResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSecretId

`func (o *ArcProviderConnectionListItemResponse) GetSecretId() string`

GetSecretId returns the SecretId field if non-nil, zero value otherwise.

### GetSecretIdOk

`func (o *ArcProviderConnectionListItemResponse) GetSecretIdOk() (*string, bool)`

GetSecretIdOk returns a tuple with the SecretId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretId

`func (o *ArcProviderConnectionListItemResponse) SetSecretId(v string)`

SetSecretId sets SecretId field to given value.


### GetName

`func (o *ArcProviderConnectionListItemResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ArcProviderConnectionListItemResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ArcProviderConnectionListItemResponse) SetName(v string)`

SetName sets Name field to given value.


### GetModels

`func (o *ArcProviderConnectionListItemResponse) GetModels() []string`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *ArcProviderConnectionListItemResponse) GetModelsOk() (*[]string, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *ArcProviderConnectionListItemResponse) SetModels(v []string)`

SetModels sets Models field to given value.


### GetBaseUrl

`func (o *ArcProviderConnectionListItemResponse) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *ArcProviderConnectionListItemResponse) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *ArcProviderConnectionListItemResponse) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.

### HasBaseUrl

`func (o *ArcProviderConnectionListItemResponse) HasBaseUrl() bool`

HasBaseUrl returns a boolean if a field has been set.

### GetAuthType

`func (o *ArcProviderConnectionListItemResponse) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *ArcProviderConnectionListItemResponse) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *ArcProviderConnectionListItemResponse) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.


### GetRegion

`func (o *ArcProviderConnectionListItemResponse) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ArcProviderConnectionListItemResponse) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ArcProviderConnectionListItemResponse) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *ArcProviderConnectionListItemResponse) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ArcProviderConnectionListItemResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ArcProviderConnectionListItemResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ArcProviderConnectionListItemResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ArcProviderConnectionListItemResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ArcProviderConnectionListItemResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ArcProviderConnectionListItemResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *ArcProviderConnectionListItemResponse) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *ArcProviderConnectionListItemResponse) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *ArcProviderConnectionListItemResponse) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *ArcProviderConnectionListItemResponse) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


