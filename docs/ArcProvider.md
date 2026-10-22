# ArcProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The provider identifier. | 
**DisplayName** | **string** | The human-readable name of the provider. | 
**DefaultBaseUrl** | **string** | The default API base URL for the provider. | 
**Models** | [**[]ArcModel**](ArcModel.md) | The models available from this provider. | 

## Methods

### NewArcProvider

`func NewArcProvider(id string, displayName string, defaultBaseUrl string, models []ArcModel, ) *ArcProvider`

NewArcProvider instantiates a new ArcProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArcProviderWithDefaults

`func NewArcProviderWithDefaults() *ArcProvider`

NewArcProviderWithDefaults instantiates a new ArcProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ArcProvider) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArcProvider) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArcProvider) SetId(v string)`

SetId sets Id field to given value.


### GetDisplayName

`func (o *ArcProvider) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ArcProvider) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ArcProvider) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetDefaultBaseUrl

`func (o *ArcProvider) GetDefaultBaseUrl() string`

GetDefaultBaseUrl returns the DefaultBaseUrl field if non-nil, zero value otherwise.

### GetDefaultBaseUrlOk

`func (o *ArcProvider) GetDefaultBaseUrlOk() (*string, bool)`

GetDefaultBaseUrlOk returns a tuple with the DefaultBaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBaseUrl

`func (o *ArcProvider) SetDefaultBaseUrl(v string)`

SetDefaultBaseUrl sets DefaultBaseUrl field to given value.


### GetModels

`func (o *ArcProvider) GetModels() []ArcModel`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *ArcProvider) GetModelsOk() (*[]ArcModel, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *ArcProvider) SetModels(v []ArcModel)`

SetModels sets Models field to given value.



[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


