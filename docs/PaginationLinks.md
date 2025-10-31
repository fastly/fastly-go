# PaginationLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**First** | Pointer to **NullableString** | The first page of data. | [optional] 
**Last** | Pointer to **NullableString** | The last page of data. | [optional] 
**Prev** | Pointer to **NullableString** | The previous page of data. | [optional] 
**Next** | Pointer to **NullableString** | The next page of data. | [optional] 

## Methods

### NewPaginationLinks

`func NewPaginationLinks() *PaginationLinks`

NewPaginationLinks instantiates a new PaginationLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginationLinksWithDefaults

`func NewPaginationLinksWithDefaults() *PaginationLinks`

NewPaginationLinksWithDefaults instantiates a new PaginationLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirst

`func (o *PaginationLinks) GetFirst() string`

GetFirst returns the First field if non-nil, zero value otherwise.

### GetFirstOk

`func (o *PaginationLinks) GetFirstOk() (*string, bool)`

GetFirstOk returns a tuple with the First field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirst

`func (o *PaginationLinks) SetFirst(v string)`

SetFirst sets First field to given value.

### HasFirst

`func (o *PaginationLinks) HasFirst() bool`

HasFirst returns a boolean if a field has been set.

### SetFirstNil

`func (o *PaginationLinks) SetFirstNil(b bool)`

 SetFirstNil sets the value for First to be an explicit nil

### UnsetFirst
`func (o *PaginationLinks) UnsetFirst()`

UnsetFirst ensures that no value is present for First, not even an explicit nil
### GetLast

`func (o *PaginationLinks) GetLast() string`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *PaginationLinks) GetLastOk() (*string, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast

`func (o *PaginationLinks) SetLast(v string)`

SetLast sets Last field to given value.

### HasLast

`func (o *PaginationLinks) HasLast() bool`

HasLast returns a boolean if a field has been set.

### SetLastNil

`func (o *PaginationLinks) SetLastNil(b bool)`

 SetLastNil sets the value for Last to be an explicit nil

### UnsetLast
`func (o *PaginationLinks) UnsetLast()`

UnsetLast ensures that no value is present for Last, not even an explicit nil
### GetPrev

`func (o *PaginationLinks) GetPrev() string`

GetPrev returns the Prev field if non-nil, zero value otherwise.

### GetPrevOk

`func (o *PaginationLinks) GetPrevOk() (*string, bool)`

GetPrevOk returns a tuple with the Prev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrev

`func (o *PaginationLinks) SetPrev(v string)`

SetPrev sets Prev field to given value.

### HasPrev

`func (o *PaginationLinks) HasPrev() bool`

HasPrev returns a boolean if a field has been set.

### SetPrevNil

`func (o *PaginationLinks) SetPrevNil(b bool)`

 SetPrevNil sets the value for Prev to be an explicit nil

### UnsetPrev
`func (o *PaginationLinks) UnsetPrev()`

UnsetPrev ensures that no value is present for Prev, not even an explicit nil
### GetNext

`func (o *PaginationLinks) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginationLinks) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginationLinks) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginationLinks) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *PaginationLinks) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *PaginationLinks) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil

[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


