# ArcProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ID** | **string** | The provider identifier. | 
**DisplayName** | **string** | The human-readable name of the provider. | 
**DefaultBaseURL** | **string** | The default API base URL for the provider. | 
**Models** | [**[]ArcModel**](ArcModel.md) | The models available from this provider. | 

## Methods

### NewArcProvider

`func NewArcProvider(id string, displayName string, defaultBaseURL string, models []ArcModel, ) *ArcProvider`

NewArcProvider instantiates a new ArcProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArcProviderWithDefaults

`func NewArcProviderWithDefaults() *ArcProvider`

NewArcProviderWithDefaults instantiates a new ArcProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetID

`func (o *ArcProvider) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *ArcProvider) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *ArcProvider) SetID(v string)`

SetID sets ID field to given value.


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


### GetDefaultBaseURL

`func (o *ArcProvider) GetDefaultBaseURL() string`

GetDefaultBaseURL returns the DefaultBaseURL field if non-nil, zero value otherwise.

### GetDefaultBaseURLOk

`func (o *ArcProvider) GetDefaultBaseURLOk() (*string, bool)`

GetDefaultBaseURLOk returns a tuple with the DefaultBaseURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBaseURL

`func (o *ArcProvider) SetDefaultBaseURL(v string)`

SetDefaultBaseURL sets DefaultBaseURL field to given value.


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
