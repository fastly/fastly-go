# Directive

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | CSP directive name (e.g., script-src, style-src) | [optional] 
**Values** | Pointer to **[]string** | Directive values | [optional] 

## Methods

### NewDirective

`func NewDirective() *Directive`

NewDirective instantiates a new Directive object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectiveWithDefaults

`func NewDirectiveWithDefaults() *Directive`

NewDirectiveWithDefaults instantiates a new Directive object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Directive) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Directive) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Directive) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Directive) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValues

`func (o *Directive) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *Directive) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *Directive) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *Directive) HasValues() bool`

HasValues returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


