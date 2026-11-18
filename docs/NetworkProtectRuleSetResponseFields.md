# NetworkProtectRuleSetResponseFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ID** | **string** | A read-only unique identifier for a Network Protect resource. | [readonly] 
**CreatedAt** | **time.Time** | An ISO 8601 timestamp in UTC. | [readonly] 
**UpdatedAt** | **time.Time** | An ISO 8601 timestamp in UTC. | [readonly] 

## Methods

### NewNetworkProtectRuleSetResponseFields

`func NewNetworkProtectRuleSetResponseFields(id string, createdAt time.Time, updatedAt time.Time, ) *NetworkProtectRuleSetResponseFields`

NewNetworkProtectRuleSetResponseFields instantiates a new NetworkProtectRuleSetResponseFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkProtectRuleSetResponseFieldsWithDefaults

`func NewNetworkProtectRuleSetResponseFieldsWithDefaults() *NetworkProtectRuleSetResponseFields`

NewNetworkProtectRuleSetResponseFieldsWithDefaults instantiates a new NetworkProtectRuleSetResponseFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetID

`func (o *NetworkProtectRuleSetResponseFields) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *NetworkProtectRuleSetResponseFields) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *NetworkProtectRuleSetResponseFields) SetID(v string)`

SetID sets ID field to given value.


### GetCreatedAt

`func (o *NetworkProtectRuleSetResponseFields) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NetworkProtectRuleSetResponseFields) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NetworkProtectRuleSetResponseFields) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *NetworkProtectRuleSetResponseFields) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *NetworkProtectRuleSetResponseFields) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *NetworkProtectRuleSetResponseFields) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
