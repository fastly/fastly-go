# DraftDiff

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Added** | Pointer to [**[]PathWithRules**](PathWithRules.md) | Paths that exist in the draft but not in the active version. | [optional] 
**Deleted** | Pointer to [**[]PathWithRules**](PathWithRules.md) | Paths that exist in the active version but not in the draft. | [optional] 
**Modified** | Pointer to [**[]PathChange**](PathChange.md) | Paths that exist in both versions but have changed. | [optional] 

## Methods

### NewDraftDiff

`func NewDraftDiff() *DraftDiff`

NewDraftDiff instantiates a new DraftDiff object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDraftDiffWithDefaults

`func NewDraftDiffWithDefaults() *DraftDiff`

NewDraftDiffWithDefaults instantiates a new DraftDiff object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdded

`func (o *DraftDiff) GetAdded() []PathWithRules`

GetAdded returns the Added field if non-nil, zero value otherwise.

### GetAddedOk

`func (o *DraftDiff) GetAddedOk() (*[]PathWithRules, bool)`

GetAddedOk returns a tuple with the Added field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdded

`func (o *DraftDiff) SetAdded(v []PathWithRules)`

SetAdded sets Added field to given value.

### HasAdded

`func (o *DraftDiff) HasAdded() bool`

HasAdded returns a boolean if a field has been set.

### GetDeleted

`func (o *DraftDiff) GetDeleted() []PathWithRules`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *DraftDiff) GetDeletedOk() (*[]PathWithRules, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *DraftDiff) SetDeleted(v []PathWithRules)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *DraftDiff) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetModified

`func (o *DraftDiff) GetModified() []PathChange`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *DraftDiff) GetModifiedOk() (*[]PathChange, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *DraftDiff) SetModified(v []PathChange)`

SetModified sets Modified field to given value.

### HasModified

`func (o *DraftDiff) HasModified() bool`

HasModified returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


