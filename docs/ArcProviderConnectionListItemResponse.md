# ArcProviderConnectionListItemResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ID** | **string** | The unique identifier of the provider connection. | [readonly] 
**SecretID** | **string** | An opaque identifier for the secret material backing this connection. Not the secret itself. | [readonly] 
**Name** | **string** | The display name of the provider connection. | 
**Models** | **[]string** | The AI model identifiers this connection is allowed to use. | 
**BaseURL** | Pointer to **string** | The base URL of the provider&#39;s API. Present only for `api-key` connections. | [optional] 
**AuthType** | **string** | The authentication type for this connection. | 
**Region** | Pointer to **string** | The AWS region for Bedrock API requests. Present only for `aws-iam` connections. | [optional] 
**CreatedAt** | **time.Time** | The date and time the connection was created, in RFC 3339 format. | [readonly] 
**UpdatedAt** | **time.Time** | The date and time the connection was last updated, in RFC 3339 format. | [readonly] 
**DeletedAt** | Pointer to **time.Time** | The date and time the connection was deleted, in RFC 3339 format. Present only for deleted connections, which are returned only when `include_deleted` is `true`. | [optional] [readonly] 

## Methods

### NewArcProviderConnectionListItemResponse

`func NewArcProviderConnectionListItemResponse(id string, secretID string, name string, models []string, authType string, createdAt time.Time, updatedAt time.Time, ) *ArcProviderConnectionListItemResponse`

NewArcProviderConnectionListItemResponse instantiates a new ArcProviderConnectionListItemResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArcProviderConnectionListItemResponseWithDefaults

`func NewArcProviderConnectionListItemResponseWithDefaults() *ArcProviderConnectionListItemResponse`

NewArcProviderConnectionListItemResponseWithDefaults instantiates a new ArcProviderConnectionListItemResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetID

`func (o *ArcProviderConnectionListItemResponse) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *ArcProviderConnectionListItemResponse) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *ArcProviderConnectionListItemResponse) SetID(v string)`

SetID sets ID field to given value.


### GetSecretID

`func (o *ArcProviderConnectionListItemResponse) GetSecretID() string`

GetSecretID returns the SecretID field if non-nil, zero value otherwise.

### GetSecretIDOk

`func (o *ArcProviderConnectionListItemResponse) GetSecretIDOk() (*string, bool)`

GetSecretIDOk returns a tuple with the SecretID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretID

`func (o *ArcProviderConnectionListItemResponse) SetSecretID(v string)`

SetSecretID sets SecretID field to given value.


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


### GetBaseURL

`func (o *ArcProviderConnectionListItemResponse) GetBaseURL() string`

GetBaseURL returns the BaseURL field if non-nil, zero value otherwise.

### GetBaseURLOk

`func (o *ArcProviderConnectionListItemResponse) GetBaseURLOk() (*string, bool)`

GetBaseURLOk returns a tuple with the BaseURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseURL

`func (o *ArcProviderConnectionListItemResponse) SetBaseURL(v string)`

SetBaseURL sets BaseURL field to given value.

### HasBaseURL

`func (o *ArcProviderConnectionListItemResponse) HasBaseURL() bool`

HasBaseURL returns a boolean if a field has been set.

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
