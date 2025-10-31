# Vcl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **string** | The VCL code to be included. | [optional] 
**Main** | Pointer to **bool** | Set to `true` when this is the main VCL, otherwise `false`. | [optional] 
**Name** | Pointer to **string** | The name of this VCL. | [optional] 

## Methods

### NewVcl

`func NewVcl() *Vcl`

NewVcl instantiates a new Vcl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVclWithDefaults

`func NewVclWithDefaults() *Vcl`

NewVclWithDefaults instantiates a new Vcl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *Vcl) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *Vcl) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *Vcl) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *Vcl) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetMain

`func (o *Vcl) GetMain() bool`

GetMain returns the Main field if non-nil, zero value otherwise.

### GetMainOk

`func (o *Vcl) GetMainOk() (*bool, bool)`

GetMainOk returns a tuple with the Main field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMain

`func (o *Vcl) SetMain(v bool)`

SetMain sets Main field to given value.

### HasMain

`func (o *Vcl) HasMain() bool`

HasMain returns a boolean if a field has been set.

### GetName

`func (o *Vcl) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Vcl) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Vcl) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Vcl) HasName() bool`

HasName returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


