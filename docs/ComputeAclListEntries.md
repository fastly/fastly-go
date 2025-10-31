# ComputeAclListEntries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**ComputeAclListEntriesMeta**](ComputeAclListEntriesMeta.md) |  | [optional] 
**Entries** | Pointer to [**[]ComputeAclListEntriesItem**](ComputeAclListEntriesItem.md) |  | [optional] 

## Methods

### NewComputeAclListEntries

`func NewComputeAclListEntries() *ComputeAclListEntries`

NewComputeAclListEntries instantiates a new ComputeAclListEntries object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputeAclListEntriesWithDefaults

`func NewComputeAclListEntriesWithDefaults() *ComputeAclListEntries`

NewComputeAclListEntriesWithDefaults instantiates a new ComputeAclListEntries object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ComputeAclListEntries) GetMeta() ComputeAclListEntriesMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ComputeAclListEntries) GetMetaOk() (*ComputeAclListEntriesMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ComputeAclListEntries) SetMeta(v ComputeAclListEntriesMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ComputeAclListEntries) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetEntries

`func (o *ComputeAclListEntries) GetEntries() []ComputeAclListEntriesItem`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *ComputeAclListEntries) GetEntriesOk() (*[]ComputeAclListEntriesItem, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *ComputeAclListEntries) SetEntries(v []ComputeAclListEntriesItem)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *ComputeAclListEntries) HasEntries() bool`

HasEntries returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


