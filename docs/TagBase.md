# TagBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the operation tag. | [optional] 
**Description** | Pointer to **string** | A description of the operation tag. | [optional] 

## Methods

### NewTagBase

`func NewTagBase() *TagBase`

NewTagBase instantiates a new TagBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagBaseWithDefaults

`func NewTagBaseWithDefaults() *TagBase`

NewTagBaseWithDefaults instantiates a new TagBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TagBase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TagBase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TagBase) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TagBase) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *TagBase) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TagBase) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TagBase) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TagBase) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


