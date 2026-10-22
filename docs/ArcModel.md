# ArcModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The model identifier. | 
**DisplayName** | **string** | The human-readable name of the model. | 
**ProviderId** | **string** | The ID of the provider this model belongs to. | 

## Methods

### NewArcModel

`func NewArcModel(id string, displayName string, providerId string, ) *ArcModel`

NewArcModel instantiates a new ArcModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArcModelWithDefaults

`func NewArcModelWithDefaults() *ArcModel`

NewArcModelWithDefaults instantiates a new ArcModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ArcModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArcModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArcModel) SetId(v string)`

SetId sets Id field to given value.


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


### GetProviderId

`func (o *ArcModel) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *ArcModel) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *ArcModel) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.



[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


