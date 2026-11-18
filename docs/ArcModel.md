# ArcModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ID** | **string** | The model identifier. | 
**DisplayName** | **string** | The human-readable name of the model. | 
**ProviderID** | **string** | The ID of the provider this model belongs to. | 

## Methods

### NewArcModel

`func NewArcModel(id string, displayName string, providerID string, ) *ArcModel`

NewArcModel instantiates a new ArcModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArcModelWithDefaults

`func NewArcModelWithDefaults() *ArcModel`

NewArcModelWithDefaults instantiates a new ArcModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetID

`func (o *ArcModel) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *ArcModel) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *ArcModel) SetID(v string)`

SetID sets ID field to given value.


### GetDisplayName

`func (o *ArcModel) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ArcModel) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ArcModel) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetProviderID

`func (o *ArcModel) GetProviderID() string`

GetProviderID returns the ProviderID field if non-nil, zero value otherwise.

### GetProviderIDOk

`func (o *ArcModel) GetProviderIDOk() (*string, bool)`

GetProviderIDOk returns a tuple with the ProviderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderID

`func (o *ArcModel) SetProviderID(v string)`

SetProviderID sets ProviderID field to given value.



[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
