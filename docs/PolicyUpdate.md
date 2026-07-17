# PolicyUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Mode** | Pointer to **string** |  | [optional] 
**Directives** | Pointer to [**[]Directive**](Directive.md) |  | [optional] 

## Methods

### NewPolicyUpdate

`func NewPolicyUpdate() *PolicyUpdate`

NewPolicyUpdate instantiates a new PolicyUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyUpdateWithDefaults

`func NewPolicyUpdateWithDefaults() *PolicyUpdate`

NewPolicyUpdateWithDefaults instantiates a new PolicyUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PolicyUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PolicyUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PolicyUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PolicyUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PolicyUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PolicyUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PolicyUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PolicyUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMode

`func (o *PolicyUpdate) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *PolicyUpdate) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *PolicyUpdate) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *PolicyUpdate) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetDirectives

`func (o *PolicyUpdate) GetDirectives() []Directive`

GetDirectives returns the Directives field if non-nil, zero value otherwise.

### GetDirectivesOk

`func (o *PolicyUpdate) GetDirectivesOk() (*[]Directive, bool)`

GetDirectivesOk returns a tuple with the Directives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectives

`func (o *PolicyUpdate) SetDirectives(v []Directive)`

SetDirectives sets Directives field to given value.

### HasDirectives

`func (o *PolicyUpdate) HasDirectives() bool`

HasDirectives returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


