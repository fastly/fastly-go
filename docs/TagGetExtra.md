# TagGetExtra

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier of the operation tag. | [readonly] 
**Count** | Pointer to **int32** | The number of operations associated with this operation tag. | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** | The date and time the operation tag was created. | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** | The date and time the operation tag was last updated. | [optional] [readonly] 

## Methods

### NewTagGetExtra

`func NewTagGetExtra(id string, ) *TagGetExtra`

NewTagGetExtra instantiates a new TagGetExtra object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagGetExtraWithDefaults

`func NewTagGetExtraWithDefaults() *TagGetExtra`

NewTagGetExtraWithDefaults instantiates a new TagGetExtra object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TagGetExtra) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TagGetExtra) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TagGetExtra) SetId(v string)`

SetId sets Id field to given value.


### GetCount

`func (o *TagGetExtra) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *TagGetExtra) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *TagGetExtra) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *TagGetExtra) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TagGetExtra) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TagGetExtra) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TagGetExtra) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TagGetExtra) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *TagGetExtra) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TagGetExtra) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TagGetExtra) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TagGetExtra) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


