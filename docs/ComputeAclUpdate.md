# ComputeACLUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entries** | Pointer to [**[]ComputeACLUpdateEntry**](ComputeACLUpdateEntry.md) |  | [optional] 

## Methods

### NewComputeACLUpdate

`func NewComputeACLUpdate() *ComputeACLUpdate`

NewComputeACLUpdate instantiates a new ComputeACLUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputeACLUpdateWithDefaults

`func NewComputeACLUpdateWithDefaults() *ComputeACLUpdate`

NewComputeACLUpdateWithDefaults instantiates a new ComputeACLUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntries

`func (o *ComputeACLUpdate) GetEntries() []ComputeACLUpdateEntry`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *ComputeACLUpdate) GetEntriesOk() (*[]ComputeACLUpdateEntry, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *ComputeACLUpdate) SetEntries(v []ComputeACLUpdateEntry)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *ComputeACLUpdate) HasEntries() bool`

HasEntries returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
