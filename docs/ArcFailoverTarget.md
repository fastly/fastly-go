# ArcFailoverTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderConnectionID** | **string** | The ID of the provider connection to retry against. | 

## Methods

### NewArcFailoverTarget

`func NewArcFailoverTarget(providerConnectionID string, ) *ArcFailoverTarget`

NewArcFailoverTarget instantiates a new ArcFailoverTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArcFailoverTargetWithDefaults

`func NewArcFailoverTargetWithDefaults() *ArcFailoverTarget`

NewArcFailoverTargetWithDefaults instantiates a new ArcFailoverTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderConnectionID

`func (o *ArcFailoverTarget) GetProviderConnectionID() string`

GetProviderConnectionID returns the ProviderConnectionID field if non-nil, zero value otherwise.

### GetProviderConnectionIDOk

`func (o *ArcFailoverTarget) GetProviderConnectionIDOk() (*string, bool)`

GetProviderConnectionIDOk returns a tuple with the ProviderConnectionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderConnectionID

`func (o *ArcFailoverTarget) SetProviderConnectionID(v string)`

SetProviderConnectionID sets ProviderConnectionID field to given value.



[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
