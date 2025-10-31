# BulkUpdateAclEntriesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entries** | Pointer to [**[]BulkUpdateAclEntry**](BulkUpdateAclEntry.md) |  | [optional] 

## Methods

### NewBulkUpdateAclEntriesRequest

`func NewBulkUpdateAclEntriesRequest() *BulkUpdateAclEntriesRequest`

NewBulkUpdateAclEntriesRequest instantiates a new BulkUpdateAclEntriesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkUpdateAclEntriesRequestWithDefaults

`func NewBulkUpdateAclEntriesRequestWithDefaults() *BulkUpdateAclEntriesRequest`

NewBulkUpdateAclEntriesRequestWithDefaults instantiates a new BulkUpdateAclEntriesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntries

`func (o *BulkUpdateAclEntriesRequest) GetEntries() []BulkUpdateAclEntry`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *BulkUpdateAclEntriesRequest) GetEntriesOk() (*[]BulkUpdateAclEntry, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *BulkUpdateAclEntriesRequest) SetEntries(v []BulkUpdateAclEntry)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *BulkUpdateAclEntriesRequest) HasEntries() bool`

HasEntries returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


