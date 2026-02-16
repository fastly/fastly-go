# TagGet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the operation tag. | 
**Description** | Pointer to **string** | A description of the operation tag. | [optional] 
**Id** | **string** | The unique identifier of the operation tag. | [readonly] 
**Count** | Pointer to **int32** | The number of operations associated with this operation tag. | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** | The date and time the operation tag was created. | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** | The date and time the operation tag was last updated. | [optional] [readonly] 

## Methods

### NewTagGet

`func NewTagGet(name string, id string, ) *TagGet`

NewTagGet instantiates a new TagGet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagGetWithDefaults

`func NewTagGetWithDefaults() *TagGet`

NewTagGetWithDefaults instantiates a new TagGet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TagGet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TagGet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TagGet) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *TagGet) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TagGet) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TagGet) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TagGet) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *TagGet) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TagGet) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TagGet) SetId(v string)`

SetId sets Id field to given value.


### GetCount

`func (o *TagGet) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *TagGet) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *TagGet) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *TagGet) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TagGet) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TagGet) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TagGet) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TagGet) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *TagGet) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TagGet) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TagGet) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TagGet) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


