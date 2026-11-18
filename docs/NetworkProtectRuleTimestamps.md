# NetworkProtectRuleTimestamps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** | An ISO 8601 timestamp in UTC. | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** | An ISO 8601 timestamp in UTC. | [optional] [readonly] 

## Methods

### NewNetworkProtectRuleTimestamps

`func NewNetworkProtectRuleTimestamps() *NetworkProtectRuleTimestamps`

NewNetworkProtectRuleTimestamps instantiates a new NetworkProtectRuleTimestamps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkProtectRuleTimestampsWithDefaults

`func NewNetworkProtectRuleTimestampsWithDefaults() *NetworkProtectRuleTimestamps`

NewNetworkProtectRuleTimestampsWithDefaults instantiates a new NetworkProtectRuleTimestamps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *NetworkProtectRuleTimestamps) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NetworkProtectRuleTimestamps) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NetworkProtectRuleTimestamps) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *NetworkProtectRuleTimestamps) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *NetworkProtectRuleTimestamps) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *NetworkProtectRuleTimestamps) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *NetworkProtectRuleTimestamps) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *NetworkProtectRuleTimestamps) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
