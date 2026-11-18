# ArcProvidersMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** | The count of providers returned. | [optional] 

## Methods

### NewArcProvidersMeta

`func NewArcProvidersMeta() *ArcProvidersMeta`

NewArcProvidersMeta instantiates a new ArcProvidersMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArcProvidersMetaWithDefaults

`func NewArcProvidersMetaWithDefaults() *ArcProvidersMeta`

NewArcProvidersMetaWithDefaults instantiates a new ArcProvidersMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *ArcProvidersMeta) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ArcProvidersMeta) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ArcProvidersMeta) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ArcProvidersMeta) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
