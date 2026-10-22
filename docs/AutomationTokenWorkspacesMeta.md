# AutomationTokenWorkspacesMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | **int32** | Number of results requested. | 
**NextCursor** | Pointer to **string** | Cursor for the next page. Empty string when there are no more results. | [optional] 
**PreviousCursor** | Pointer to **string** | Cursor for the previous page. Empty string when on the first page. | [optional] 

## Methods

### NewAutomationTokenWorkspacesMeta

`func NewAutomationTokenWorkspacesMeta(limit int32, ) *AutomationTokenWorkspacesMeta`

NewAutomationTokenWorkspacesMeta instantiates a new AutomationTokenWorkspacesMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutomationTokenWorkspacesMetaWithDefaults

`func NewAutomationTokenWorkspacesMetaWithDefaults() *AutomationTokenWorkspacesMeta`

NewAutomationTokenWorkspacesMetaWithDefaults instantiates a new AutomationTokenWorkspacesMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *AutomationTokenWorkspacesMeta) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *AutomationTokenWorkspacesMeta) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *AutomationTokenWorkspacesMeta) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetNextCursor

`func (o *AutomationTokenWorkspacesMeta) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *AutomationTokenWorkspacesMeta) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *AutomationTokenWorkspacesMeta) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *AutomationTokenWorkspacesMeta) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.

### GetPreviousCursor

`func (o *AutomationTokenWorkspacesMeta) GetPreviousCursor() string`

GetPreviousCursor returns the PreviousCursor field if non-nil, zero value otherwise.

### GetPreviousCursorOk

`func (o *AutomationTokenWorkspacesMeta) GetPreviousCursorOk() (*string, bool)`

GetPreviousCursorOk returns a tuple with the PreviousCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousCursor

`func (o *AutomationTokenWorkspacesMeta) SetPreviousCursor(v string)`

SetPreviousCursor sets PreviousCursor field to given value.

### HasPreviousCursor

`func (o *AutomationTokenWorkspacesMeta) HasPreviousCursor() bool`

HasPreviousCursor returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


