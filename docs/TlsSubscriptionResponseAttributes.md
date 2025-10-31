# TlsSubscriptionResponseAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**DeletedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**UpdatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**State** | Pointer to **string** | The current state of your subscription. | [optional] 
**HasActiveOrder** | Pointer to **bool** | Subscription has an active order | [optional] 

## Methods

### NewTlsSubscriptionResponseAttributes

`func NewTlsSubscriptionResponseAttributes() *TlsSubscriptionResponseAttributes`

NewTlsSubscriptionResponseAttributes instantiates a new TlsSubscriptionResponseAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTlsSubscriptionResponseAttributesWithDefaults

`func NewTlsSubscriptionResponseAttributesWithDefaults() *TlsSubscriptionResponseAttributes`

NewTlsSubscriptionResponseAttributesWithDefaults instantiates a new TlsSubscriptionResponseAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *TlsSubscriptionResponseAttributes) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TlsSubscriptionResponseAttributes) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TlsSubscriptionResponseAttributes) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TlsSubscriptionResponseAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *TlsSubscriptionResponseAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *TlsSubscriptionResponseAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetDeletedAt

`func (o *TlsSubscriptionResponseAttributes) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *TlsSubscriptionResponseAttributes) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *TlsSubscriptionResponseAttributes) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *TlsSubscriptionResponseAttributes) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *TlsSubscriptionResponseAttributes) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *TlsSubscriptionResponseAttributes) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetUpdatedAt

`func (o *TlsSubscriptionResponseAttributes) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TlsSubscriptionResponseAttributes) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TlsSubscriptionResponseAttributes) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TlsSubscriptionResponseAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *TlsSubscriptionResponseAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *TlsSubscriptionResponseAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetState

`func (o *TlsSubscriptionResponseAttributes) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *TlsSubscriptionResponseAttributes) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *TlsSubscriptionResponseAttributes) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *TlsSubscriptionResponseAttributes) HasState() bool`

HasState returns a boolean if a field has been set.

### GetHasActiveOrder

`func (o *TlsSubscriptionResponseAttributes) GetHasActiveOrder() bool`

GetHasActiveOrder returns the HasActiveOrder field if non-nil, zero value otherwise.

### GetHasActiveOrderOk

`func (o *TlsSubscriptionResponseAttributes) GetHasActiveOrderOk() (*bool, bool)`

GetHasActiveOrderOk returns a tuple with the HasActiveOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasActiveOrder

`func (o *TlsSubscriptionResponseAttributes) SetHasActiveOrder(v bool)`

SetHasActiveOrder sets HasActiveOrder field to given value.

### HasHasActiveOrder

`func (o *TlsSubscriptionResponseAttributes) HasHasActiveOrder() bool`

HasHasActiveOrder returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


