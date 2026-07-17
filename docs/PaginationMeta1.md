# PaginationMeta1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentPage** | Pointer to **int32** | Current page. | [optional] 
**PerPage** | Pointer to **int32** | Number of records per page. | [optional] [default to 20]
**RecordCount** | Pointer to **int32** | Total records in result set. | [optional] 
**TotalPages** | Pointer to **int32** | Total pages in result set. | [optional] 

## Methods

### NewPaginationMeta1

`func NewPaginationMeta1() *PaginationMeta1`

NewPaginationMeta1 instantiates a new PaginationMeta1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginationMeta1WithDefaults

`func NewPaginationMeta1WithDefaults() *PaginationMeta1`

NewPaginationMeta1WithDefaults instantiates a new PaginationMeta1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentPage

`func (o *PaginationMeta1) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *PaginationMeta1) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *PaginationMeta1) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.

### HasCurrentPage

`func (o *PaginationMeta1) HasCurrentPage() bool`

HasCurrentPage returns a boolean if a field has been set.

### GetPerPage

`func (o *PaginationMeta1) GetPerPage() int32`

GetPerPage returns the PerPage field if non-nil, zero value otherwise.

### GetPerPageOk

`func (o *PaginationMeta1) GetPerPageOk() (*int32, bool)`

GetPerPageOk returns a tuple with the PerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerPage

`func (o *PaginationMeta1) SetPerPage(v int32)`

SetPerPage sets PerPage field to given value.

### HasPerPage

`func (o *PaginationMeta1) HasPerPage() bool`

HasPerPage returns a boolean if a field has been set.

### GetRecordCount

`func (o *PaginationMeta1) GetRecordCount() int32`

GetRecordCount returns the RecordCount field if non-nil, zero value otherwise.

### GetRecordCountOk

`func (o *PaginationMeta1) GetRecordCountOk() (*int32, bool)`

GetRecordCountOk returns a tuple with the RecordCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordCount

`func (o *PaginationMeta1) SetRecordCount(v int32)`

SetRecordCount sets RecordCount field to given value.

### HasRecordCount

`func (o *PaginationMeta1) HasRecordCount() bool`

HasRecordCount returns a boolean if a field has been set.

### GetTotalPages

`func (o *PaginationMeta1) GetTotalPages() int32`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *PaginationMeta1) GetTotalPagesOk() (*int32, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *PaginationMeta1) SetTotalPages(v int32)`

SetTotalPages sets TotalPages field to given value.

### HasTotalPages

`func (o *PaginationMeta1) HasTotalPages() bool`

HasTotalPages returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


