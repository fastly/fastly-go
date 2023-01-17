# WafExclusionResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**DeletedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**UpdatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**Condition** | Pointer to **string** | A conditional expression in VCL used to determine if the condition is met. | [optional] 
**ExclusionType** | Pointer to **string** | The type of exclusion. | [optional] 
**Logging** | Pointer to **bool** | Whether to generate a log upon matching. | [optional] [default to true]
**Name** | Pointer to **string** | Name of the exclusion. | [optional] 
**Number** | Pointer to **int32** | A numeric ID identifying a WAF exclusion. | [optional] 
**Variable** | Pointer to **NullableString** | The variable to exclude. An optional selector can be specified after the variable separated by a colon (`:`) to restrict the variable to a particular parameter. Required for `exclusion_type&#x3D;variable`. | [optional] 

## Methods

### NewWafExclusionResponseDataAttributes

`func NewWafExclusionResponseDataAttributes() *WafExclusionResponseDataAttributes`

NewWafExclusionResponseDataAttributes instantiates a new WafExclusionResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafExclusionResponseDataAttributesWithDefaults

`func NewWafExclusionResponseDataAttributesWithDefaults() *WafExclusionResponseDataAttributes`

NewWafExclusionResponseDataAttributesWithDefaults instantiates a new WafExclusionResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *WafExclusionResponseDataAttributes) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WafExclusionResponseDataAttributes) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WafExclusionResponseDataAttributes) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *WafExclusionResponseDataAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *WafExclusionResponseDataAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *WafExclusionResponseDataAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetDeletedAt

`func (o *WafExclusionResponseDataAttributes) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *WafExclusionResponseDataAttributes) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *WafExclusionResponseDataAttributes) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *WafExclusionResponseDataAttributes) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *WafExclusionResponseDataAttributes) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *WafExclusionResponseDataAttributes) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetUpdatedAt

`func (o *WafExclusionResponseDataAttributes) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *WafExclusionResponseDataAttributes) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *WafExclusionResponseDataAttributes) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *WafExclusionResponseDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *WafExclusionResponseDataAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *WafExclusionResponseDataAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetCondition

`func (o *WafExclusionResponseDataAttributes) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *WafExclusionResponseDataAttributes) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *WafExclusionResponseDataAttributes) SetCondition(v string)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *WafExclusionResponseDataAttributes) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetExclusionType

`func (o *WafExclusionResponseDataAttributes) GetExclusionType() string`

GetExclusionType returns the ExclusionType field if non-nil, zero value otherwise.

### GetExclusionTypeOk

`func (o *WafExclusionResponseDataAttributes) GetExclusionTypeOk() (*string, bool)`

GetExclusionTypeOk returns a tuple with the ExclusionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclusionType

`func (o *WafExclusionResponseDataAttributes) SetExclusionType(v string)`

SetExclusionType sets ExclusionType field to given value.

### HasExclusionType

`func (o *WafExclusionResponseDataAttributes) HasExclusionType() bool`

HasExclusionType returns a boolean if a field has been set.

### GetLogging

`func (o *WafExclusionResponseDataAttributes) GetLogging() bool`

GetLogging returns the Logging field if non-nil, zero value otherwise.

### GetLoggingOk

`func (o *WafExclusionResponseDataAttributes) GetLoggingOk() (*bool, bool)`

GetLoggingOk returns a tuple with the Logging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogging

`func (o *WafExclusionResponseDataAttributes) SetLogging(v bool)`

SetLogging sets Logging field to given value.

### HasLogging

`func (o *WafExclusionResponseDataAttributes) HasLogging() bool`

HasLogging returns a boolean if a field has been set.

### GetName

`func (o *WafExclusionResponseDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WafExclusionResponseDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WafExclusionResponseDataAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WafExclusionResponseDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumber

`func (o *WafExclusionResponseDataAttributes) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *WafExclusionResponseDataAttributes) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *WafExclusionResponseDataAttributes) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *WafExclusionResponseDataAttributes) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetVariable

`func (o *WafExclusionResponseDataAttributes) GetVariable() string`

GetVariable returns the Variable field if non-nil, zero value otherwise.

### GetVariableOk

`func (o *WafExclusionResponseDataAttributes) GetVariableOk() (*string, bool)`

GetVariableOk returns a tuple with the Variable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariable

`func (o *WafExclusionResponseDataAttributes) SetVariable(v string)`

SetVariable sets Variable field to given value.

### HasVariable

`func (o *WafExclusionResponseDataAttributes) HasVariable() bool`

HasVariable returns a boolean if a field has been set.

### SetVariableNil

`func (o *WafExclusionResponseDataAttributes) SetVariableNil(b bool)`

 SetVariableNil sets the value for Variable to be an explicit nil

### UnsetVariable
`func (o *WafExclusionResponseDataAttributes) UnsetVariable()`

UnsetVariable ensures that no value is present for Variable, not even an explicit nil

[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
