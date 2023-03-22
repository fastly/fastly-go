# BulkUpdateACLEntriesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entries** | Pointer to [**[]BulkUpdateACLEntry**](BulkUpdateACLEntry.md) |  | [optional] 

## Methods

### NewBulkUpdateACLEntriesRequest

`func NewBulkUpdateACLEntriesRequest() *BulkUpdateACLEntriesRequest`

NewBulkUpdateACLEntriesRequest instantiates a new BulkUpdateACLEntriesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkUpdateACLEntriesRequestWithDefaults

`func NewBulkUpdateACLEntriesRequestWithDefaults() *BulkUpdateACLEntriesRequest`

NewBulkUpdateACLEntriesRequestWithDefaults instantiates a new BulkUpdateACLEntriesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntries

`func (o *BulkUpdateACLEntriesRequest) GetEntries() []BulkUpdateACLEntry`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *BulkUpdateACLEntriesRequest) GetEntriesOk() (*[]BulkUpdateACLEntry, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *BulkUpdateACLEntriesRequest) SetEntries(v []BulkUpdateACLEntry)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *BulkUpdateACLEntriesRequest) HasEntries() bool`

HasEntries returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
