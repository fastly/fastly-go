# ComputeAclUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entries** | Pointer to [**[]ComputeAclUpdateEntry**](ComputeAclUpdateEntry.md) |  | [optional] 

## Methods

### NewComputeAclUpdate

`func NewComputeAclUpdate() *ComputeAclUpdate`

NewComputeAclUpdate instantiates a new ComputeAclUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputeAclUpdateWithDefaults

`func NewComputeAclUpdateWithDefaults() *ComputeAclUpdate`

NewComputeAclUpdateWithDefaults instantiates a new ComputeAclUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntries

`func (o *ComputeAclUpdate) GetEntries() []ComputeAclUpdateEntry`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *ComputeAclUpdate) GetEntriesOk() (*[]ComputeAclUpdateEntry, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *ComputeAclUpdate) SetEntries(v []ComputeAclUpdateEntry)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *ComputeAclUpdate) HasEntries() bool`

HasEntries returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


