# ArcFailoverTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderConnectionId** | **string** | The ID of the provider connection to retry against. | 

## Methods

### NewArcFailoverTarget

`func NewArcFailoverTarget(providerConnectionId string, ) *ArcFailoverTarget`

NewArcFailoverTarget instantiates a new ArcFailoverTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArcFailoverTargetWithDefaults

`func NewArcFailoverTargetWithDefaults() *ArcFailoverTarget`

NewArcFailoverTargetWithDefaults instantiates a new ArcFailoverTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderConnectionId

`func (o *ArcFailoverTarget) GetProviderConnectionId() string`

GetProviderConnectionId returns the ProviderConnectionId field if non-nil, zero value otherwise.

### GetProviderConnectionIdOk

`func (o *ArcFailoverTarget) GetProviderConnectionIdOk() (*string, bool)`

GetProviderConnectionIdOk returns a tuple with the ProviderConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderConnectionId

`func (o *ArcFailoverTarget) SetProviderConnectionId(v string)`

SetProviderConnectionId sets ProviderConnectionId field to given value.



[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


