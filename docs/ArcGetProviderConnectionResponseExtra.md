# ArcGetProviderConnectionResponseExtra

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VirtualKeyIds** | **[]string** | The IDs of the active virtual keys whose failover policy references this connection. Empty if no virtual key references it. | [readonly] 

## Methods

### NewArcGetProviderConnectionResponseExtra

`func NewArcGetProviderConnectionResponseExtra(virtualKeyIds []string, ) *ArcGetProviderConnectionResponseExtra`

NewArcGetProviderConnectionResponseExtra instantiates a new ArcGetProviderConnectionResponseExtra object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArcGetProviderConnectionResponseExtraWithDefaults

`func NewArcGetProviderConnectionResponseExtraWithDefaults() *ArcGetProviderConnectionResponseExtra`

NewArcGetProviderConnectionResponseExtraWithDefaults instantiates a new ArcGetProviderConnectionResponseExtra object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVirtualKeyIds

`func (o *ArcGetProviderConnectionResponseExtra) GetVirtualKeyIds() []string`

GetVirtualKeyIds returns the VirtualKeyIds field if non-nil, zero value otherwise.

### GetVirtualKeyIdsOk

`func (o *ArcGetProviderConnectionResponseExtra) GetVirtualKeyIdsOk() (*[]string, bool)`

GetVirtualKeyIdsOk returns a tuple with the VirtualKeyIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualKeyIds

`func (o *ArcGetProviderConnectionResponseExtra) SetVirtualKeyIds(v []string)`

SetVirtualKeyIds sets VirtualKeyIds field to given value.



[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


