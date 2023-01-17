# WafExclusionResponseDataAttributesAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Condition** | Pointer to **string** | A conditional expression in VCL used to determine if the condition is met. | [optional] 
**ExclusionType** | Pointer to **string** | The type of exclusion. | [optional] 
**Logging** | Pointer to **bool** | Whether to generate a log upon matching. | [optional] [default to true]
**Name** | Pointer to **string** | Name of the exclusion. | [optional] 
**Number** | Pointer to **int32** | A numeric ID identifying a WAF exclusion. | [optional] 
**Variable** | Pointer to **NullableString** | The variable to exclude. An optional selector can be specified after the variable separated by a colon (`:`) to restrict the variable to a particular parameter. Required for `exclusion_type&#x3D;variable`. | [optional] 

## Methods

### NewWafExclusionResponseDataAttributesAllOf

`func NewWafExclusionResponseDataAttributesAllOf() *WafExclusionResponseDataAttributesAllOf`

NewWafExclusionResponseDataAttributesAllOf instantiates a new WafExclusionResponseDataAttributesAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafExclusionResponseDataAttributesAllOfWithDefaults

`func NewWafExclusionResponseDataAttributesAllOfWithDefaults() *WafExclusionResponseDataAttributesAllOf`

NewWafExclusionResponseDataAttributesAllOfWithDefaults instantiates a new WafExclusionResponseDataAttributesAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCondition

`func (o *WafExclusionResponseDataAttributesAllOf) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *WafExclusionResponseDataAttributesAllOf) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *WafExclusionResponseDataAttributesAllOf) SetCondition(v string)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *WafExclusionResponseDataAttributesAllOf) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetExclusionType

`func (o *WafExclusionResponseDataAttributesAllOf) GetExclusionType() string`

GetExclusionType returns the ExclusionType field if non-nil, zero value otherwise.

### GetExclusionTypeOk

`func (o *WafExclusionResponseDataAttributesAllOf) GetExclusionTypeOk() (*string, bool)`

GetExclusionTypeOk returns a tuple with the ExclusionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclusionType

`func (o *WafExclusionResponseDataAttributesAllOf) SetExclusionType(v string)`

SetExclusionType sets ExclusionType field to given value.

### HasExclusionType

`func (o *WafExclusionResponseDataAttributesAllOf) HasExclusionType() bool`

HasExclusionType returns a boolean if a field has been set.

### GetLogging

`func (o *WafExclusionResponseDataAttributesAllOf) GetLogging() bool`

GetLogging returns the Logging field if non-nil, zero value otherwise.

### GetLoggingOk

`func (o *WafExclusionResponseDataAttributesAllOf) GetLoggingOk() (*bool, bool)`

GetLoggingOk returns a tuple with the Logging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogging

`func (o *WafExclusionResponseDataAttributesAllOf) SetLogging(v bool)`

SetLogging sets Logging field to given value.

### HasLogging

`func (o *WafExclusionResponseDataAttributesAllOf) HasLogging() bool`

HasLogging returns a boolean if a field has been set.

### GetName

`func (o *WafExclusionResponseDataAttributesAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WafExclusionResponseDataAttributesAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WafExclusionResponseDataAttributesAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WafExclusionResponseDataAttributesAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumber

`func (o *WafExclusionResponseDataAttributesAllOf) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *WafExclusionResponseDataAttributesAllOf) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *WafExclusionResponseDataAttributesAllOf) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *WafExclusionResponseDataAttributesAllOf) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetVariable

`func (o *WafExclusionResponseDataAttributesAllOf) GetVariable() string`

GetVariable returns the Variable field if non-nil, zero value otherwise.

### GetVariableOk

`func (o *WafExclusionResponseDataAttributesAllOf) GetVariableOk() (*string, bool)`

GetVariableOk returns a tuple with the Variable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariable

`func (o *WafExclusionResponseDataAttributesAllOf) SetVariable(v string)`

SetVariable sets Variable field to given value.

### HasVariable

`func (o *WafExclusionResponseDataAttributesAllOf) HasVariable() bool`

HasVariable returns a boolean if a field has been set.

### SetVariableNil

`func (o *WafExclusionResponseDataAttributesAllOf) SetVariableNil(b bool)`

 SetVariableNil sets the value for Variable to be an explicit nil

### UnsetVariable
`func (o *WafExclusionResponseDataAttributesAllOf) UnsetVariable()`

UnsetVariable ensures that no value is present for Variable, not even an explicit nil

[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
