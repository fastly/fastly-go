# TagCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the operation tag. | [optional] 
**Description** | Pointer to **string** | A description of the operation tag. | [optional] 

## Methods

### NewTagCreate

`func NewTagCreate() *TagCreate`

NewTagCreate instantiates a new TagCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagCreateWithDefaults

`func NewTagCreateWithDefaults() *TagCreate`

NewTagCreateWithDefaults instantiates a new TagCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TagCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TagCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TagCreate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TagCreate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *TagCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TagCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TagCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TagCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


