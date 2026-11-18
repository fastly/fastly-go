# ComputeACLListEntries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**ComputeACLListEntriesMeta**](ComputeACLListEntriesMeta.md) |  | [optional] 
**Entries** | Pointer to [**[]ComputeACLListEntriesItem**](ComputeACLListEntriesItem.md) |  | [optional] 

## Methods

### NewComputeACLListEntries

`func NewComputeACLListEntries() *ComputeACLListEntries`

NewComputeACLListEntries instantiates a new ComputeACLListEntries object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputeACLListEntriesWithDefaults

`func NewComputeACLListEntriesWithDefaults() *ComputeACLListEntries`

NewComputeACLListEntriesWithDefaults instantiates a new ComputeACLListEntries object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ComputeACLListEntries) GetMeta() ComputeACLListEntriesMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ComputeACLListEntries) GetMetaOk() (*ComputeACLListEntriesMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ComputeACLListEntries) SetMeta(v ComputeACLListEntriesMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ComputeACLListEntries) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetEntries

`func (o *ComputeACLListEntries) GetEntries() []ComputeACLListEntriesItem`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *ComputeACLListEntries) GetEntriesOk() (*[]ComputeACLListEntriesItem, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *ComputeACLListEntries) SetEntries(v []ComputeACLListEntriesItem)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *ComputeACLListEntries) HasEntries() bool`

HasEntries returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
