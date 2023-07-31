# PoolResponsePostAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Quorum** | Pointer to **int32** | Percentage of capacity (`0-100`) that needs to be operationally available for a pool to be considered up. | [optional] [default to 75]

## Methods

### NewPoolResponsePostAllOf

`func NewPoolResponsePostAllOf() *PoolResponsePostAllOf`

NewPoolResponsePostAllOf instantiates a new PoolResponsePostAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolResponsePostAllOfWithDefaults

`func NewPoolResponsePostAllOfWithDefaults() *PoolResponsePostAllOf`

NewPoolResponsePostAllOfWithDefaults instantiates a new PoolResponsePostAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuorum

`func (o *PoolResponsePostAllOf) GetQuorum() int32`

GetQuorum returns the Quorum field if non-nil, zero value otherwise.

### GetQuorumOk

`func (o *PoolResponsePostAllOf) GetQuorumOk() (*int32, bool)`

GetQuorumOk returns a tuple with the Quorum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuorum

`func (o *PoolResponsePostAllOf) SetQuorum(v int32)`

SetQuorum sets Quorum field to given value.

### HasQuorum

`func (o *PoolResponsePostAllOf) HasQuorum() bool`

HasQuorum returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
