# PolicyCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Mode** | **string** |  | 
**Directives** | Pointer to [**[]Directive**](Directive.md) |  | [optional] 

## Methods

### NewPolicyCreate

`func NewPolicyCreate(name string, mode string, ) *PolicyCreate`

NewPolicyCreate instantiates a new PolicyCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyCreateWithDefaults

`func NewPolicyCreateWithDefaults() *PolicyCreate`

NewPolicyCreateWithDefaults instantiates a new PolicyCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PolicyCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PolicyCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PolicyCreate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *PolicyCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PolicyCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PolicyCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PolicyCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMode

`func (o *PolicyCreate) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *PolicyCreate) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *PolicyCreate) SetMode(v string)`

SetMode sets Mode field to given value.


### GetDirectives

`func (o *PolicyCreate) GetDirectives() []Directive`

GetDirectives returns the Directives field if non-nil, zero value otherwise.

### GetDirectivesOk

`func (o *PolicyCreate) GetDirectivesOk() (*[]Directive, bool)`

GetDirectivesOk returns a tuple with the Directives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectives

`func (o *PolicyCreate) SetDirectives(v []Directive)`

SetDirectives sets Directives field to given value.

### HasDirectives

`func (o *PolicyCreate) HasDirectives() bool`

HasDirectives returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


