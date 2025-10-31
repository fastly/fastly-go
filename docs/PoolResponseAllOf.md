# PoolResponseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Quorum** | Pointer to **string** | Percentage of capacity (`0-100`) that needs to be operationally available for a pool to be considered up. | [optional] [default to "75"]

## Methods

### NewPoolResponseAllOf

`func NewPoolResponseAllOf() *PoolResponseAllOf`

NewPoolResponseAllOf instantiates a new PoolResponseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolResponseAllOfWithDefaults

`func NewPoolResponseAllOfWithDefaults() *PoolResponseAllOf`

NewPoolResponseAllOfWithDefaults instantiates a new PoolResponseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuorum

`func (o *PoolResponseAllOf) GetQuorum() string`

GetQuorum returns the Quorum field if non-nil, zero value otherwise.

### GetQuorumOk

`func (o *PoolResponseAllOf) GetQuorumOk() (*string, bool)`

GetQuorumOk returns a tuple with the Quorum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuorum

`func (o *PoolResponseAllOf) SetQuorum(v string)`

SetQuorum sets Quorum field to given value.

### HasQuorum

`func (o *PoolResponseAllOf) HasQuorum() bool`

HasQuorum returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


