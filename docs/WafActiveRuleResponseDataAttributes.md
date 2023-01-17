# WafActiveRuleResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**DeletedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**UpdatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**LatestRevision** | Pointer to **int32** | The latest rule revision number that is available for the associated rule revision. | [optional] [readonly] 
**Outdated** | Pointer to **bool** | Indicates if the associated rule revision is up to date or not. | [optional] [readonly] 

## Methods

### NewWafActiveRuleResponseDataAttributes

`func NewWafActiveRuleResponseDataAttributes() *WafActiveRuleResponseDataAttributes`

NewWafActiveRuleResponseDataAttributes instantiates a new WafActiveRuleResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafActiveRuleResponseDataAttributesWithDefaults

`func NewWafActiveRuleResponseDataAttributesWithDefaults() *WafActiveRuleResponseDataAttributes`

NewWafActiveRuleResponseDataAttributesWithDefaults instantiates a new WafActiveRuleResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *WafActiveRuleResponseDataAttributes) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WafActiveRuleResponseDataAttributes) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WafActiveRuleResponseDataAttributes) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *WafActiveRuleResponseDataAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *WafActiveRuleResponseDataAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *WafActiveRuleResponseDataAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetDeletedAt

`func (o *WafActiveRuleResponseDataAttributes) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *WafActiveRuleResponseDataAttributes) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *WafActiveRuleResponseDataAttributes) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *WafActiveRuleResponseDataAttributes) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *WafActiveRuleResponseDataAttributes) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *WafActiveRuleResponseDataAttributes) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetUpdatedAt

`func (o *WafActiveRuleResponseDataAttributes) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *WafActiveRuleResponseDataAttributes) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *WafActiveRuleResponseDataAttributes) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *WafActiveRuleResponseDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *WafActiveRuleResponseDataAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *WafActiveRuleResponseDataAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetLatestRevision

`func (o *WafActiveRuleResponseDataAttributes) GetLatestRevision() int32`

GetLatestRevision returns the LatestRevision field if non-nil, zero value otherwise.

### GetLatestRevisionOk

`func (o *WafActiveRuleResponseDataAttributes) GetLatestRevisionOk() (*int32, bool)`

GetLatestRevisionOk returns a tuple with the LatestRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestRevision

`func (o *WafActiveRuleResponseDataAttributes) SetLatestRevision(v int32)`

SetLatestRevision sets LatestRevision field to given value.

### HasLatestRevision

`func (o *WafActiveRuleResponseDataAttributes) HasLatestRevision() bool`

HasLatestRevision returns a boolean if a field has been set.

### GetOutdated

`func (o *WafActiveRuleResponseDataAttributes) GetOutdated() bool`

GetOutdated returns the Outdated field if non-nil, zero value otherwise.

### GetOutdatedOk

`func (o *WafActiveRuleResponseDataAttributes) GetOutdatedOk() (*bool, bool)`

GetOutdatedOk returns a tuple with the Outdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutdated

`func (o *WafActiveRuleResponseDataAttributes) SetOutdated(v bool)`

SetOutdated sets Outdated field to given value.

### HasOutdated

`func (o *WafActiveRuleResponseDataAttributes) HasOutdated() bool`

HasOutdated returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
