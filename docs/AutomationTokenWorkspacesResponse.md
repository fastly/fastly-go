# AutomationTokenWorkspacesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]AutomationTokenWorkspace**](AutomationTokenWorkspace.md) | Workspaces the token can access. Empty if the token has no workspace-specific permissions or if `limit_workspaces` is `false` (unrestricted access). | 
**Meta** | [**AutomationTokenWorkspacesMeta**](AutomationTokenWorkspacesMeta.md) |  | 
**LimitWorkspaces** | **bool** | Indicates whether workspace access is restricted: - `true`: Token can only access workspaces listed in `data` - `false`: Token has unrestricted access to all workspaces  | 

## Methods

### NewAutomationTokenWorkspacesResponse

`func NewAutomationTokenWorkspacesResponse(data []AutomationTokenWorkspace, meta AutomationTokenWorkspacesMeta, limitWorkspaces bool, ) *AutomationTokenWorkspacesResponse`

NewAutomationTokenWorkspacesResponse instantiates a new AutomationTokenWorkspacesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutomationTokenWorkspacesResponseWithDefaults

`func NewAutomationTokenWorkspacesResponseWithDefaults() *AutomationTokenWorkspacesResponse`

NewAutomationTokenWorkspacesResponseWithDefaults instantiates a new AutomationTokenWorkspacesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AutomationTokenWorkspacesResponse) GetData() []AutomationTokenWorkspace`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AutomationTokenWorkspacesResponse) GetDataOk() (*[]AutomationTokenWorkspace, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AutomationTokenWorkspacesResponse) SetData(v []AutomationTokenWorkspace)`

SetData sets Data field to given value.


### GetMeta

`func (o *AutomationTokenWorkspacesResponse) GetMeta() AutomationTokenWorkspacesMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AutomationTokenWorkspacesResponse) GetMetaOk() (*AutomationTokenWorkspacesMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AutomationTokenWorkspacesResponse) SetMeta(v AutomationTokenWorkspacesMeta)`

SetMeta sets Meta field to given value.


### GetLimitWorkspaces

`func (o *AutomationTokenWorkspacesResponse) GetLimitWorkspaces() bool`

GetLimitWorkspaces returns the LimitWorkspaces field if non-nil, zero value otherwise.

### GetLimitWorkspacesOk

`func (o *AutomationTokenWorkspacesResponse) GetLimitWorkspacesOk() (*bool, bool)`

GetLimitWorkspacesOk returns a tuple with the LimitWorkspaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitWorkspaces

`func (o *AutomationTokenWorkspacesResponse) SetLimitWorkspaces(v bool)`

SetLimitWorkspaces sets LimitWorkspaces field to given value.



[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
