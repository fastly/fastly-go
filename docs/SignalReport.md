# SignalReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the attack type. | [optional] 
**DisplayName** | Pointer to **string** | Display name of the attack type. | [optional] 
**Count** | Pointer to **int32** | Total count of attacks of this type. | [optional] 
**TopWorkspaces** | Pointer to [**[]TopWorkspace**](TopWorkspace.md) | Top workspaces affected by this attack type. | [optional] 

## Methods

### NewSignalReport

`func NewSignalReport() *SignalReport`

NewSignalReport instantiates a new SignalReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignalReportWithDefaults

`func NewSignalReportWithDefaults() *SignalReport`

NewSignalReportWithDefaults instantiates a new SignalReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SignalReport) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SignalReport) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SignalReport) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SignalReport) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDisplayName

`func (o *SignalReport) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *SignalReport) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *SignalReport) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *SignalReport) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetCount

`func (o *SignalReport) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *SignalReport) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *SignalReport) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *SignalReport) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetTopWorkspaces

`func (o *SignalReport) GetTopWorkspaces() []TopWorkspace`

GetTopWorkspaces returns the TopWorkspaces field if non-nil, zero value otherwise.

### GetTopWorkspacesOk

`func (o *SignalReport) GetTopWorkspacesOk() (*[]TopWorkspace, bool)`

GetTopWorkspacesOk returns a tuple with the TopWorkspaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopWorkspaces

`func (o *SignalReport) SetTopWorkspaces(v []TopWorkspace)`

SetTopWorkspaces sets TopWorkspaces field to given value.

### HasTopWorkspaces

`func (o *SignalReport) HasTopWorkspaces() bool`

HasTopWorkspaces returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
