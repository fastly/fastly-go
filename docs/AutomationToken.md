# AutomationToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the token. | [optional] 
**Role** | Pointer to **string** | The primary role on the token. For tokens with multiple roles, this returns the highest-priority role for backwards compatibility. | [optional] 
**Roles** | Pointer to [**[]AutomationTokenRole**](AutomationTokenRole.md) | All roles assigned to the token. Each role includes an `id` and `name`. | [optional] 
**Services** | Pointer to **[]string** | The service IDs of the services the token has access to. Empty when `limit_services` is `false` (full access) or when the token has no service access. | [optional] 
**LimitServices** | Pointer to **bool** | When `true`, the token is restricted to only the services in the `services` array. When `false`, the token has access to all services. | [optional] 
**LimitWorkspaces** | Pointer to **bool** | When `true`, the token is restricted to only the workspaces in the `workspaces` array. When `false`, the token has access to all workspaces. Only applies to roles that support workspace access. | [optional] 
**Workspaces** | Pointer to **[]string** | The workspace IDs of the workspaces the token has access to. Returns an empty array for tokens without workspace access. | [optional] 
**Scope** | Pointer to **string** | A space-delimited list of authorization scope. | [optional] [default to "global"]
**ExpiresAt** | Pointer to **string** | A UTC timestamp of when the token expires. | [optional] 

## Methods

### NewAutomationToken

`func NewAutomationToken() *AutomationToken`

NewAutomationToken instantiates a new AutomationToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutomationTokenWithDefaults

`func NewAutomationTokenWithDefaults() *AutomationToken`

NewAutomationTokenWithDefaults instantiates a new AutomationToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AutomationToken) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AutomationToken) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AutomationToken) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AutomationToken) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRole

`func (o *AutomationToken) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *AutomationToken) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *AutomationToken) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *AutomationToken) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetRoles

`func (o *AutomationToken) GetRoles() []AutomationTokenRole`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *AutomationToken) GetRolesOk() (*[]AutomationTokenRole, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *AutomationToken) SetRoles(v []AutomationTokenRole)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *AutomationToken) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetServices

`func (o *AutomationToken) GetServices() []string`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *AutomationToken) GetServicesOk() (*[]string, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *AutomationToken) SetServices(v []string)`

SetServices sets Services field to given value.

### HasServices

`func (o *AutomationToken) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetLimitServices

`func (o *AutomationToken) GetLimitServices() bool`

GetLimitServices returns the LimitServices field if non-nil, zero value otherwise.

### GetLimitServicesOk

`func (o *AutomationToken) GetLimitServicesOk() (*bool, bool)`

GetLimitServicesOk returns a tuple with the LimitServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitServices

`func (o *AutomationToken) SetLimitServices(v bool)`

SetLimitServices sets LimitServices field to given value.

### HasLimitServices

`func (o *AutomationToken) HasLimitServices() bool`

HasLimitServices returns a boolean if a field has been set.

### GetLimitWorkspaces

`func (o *AutomationToken) GetLimitWorkspaces() bool`

GetLimitWorkspaces returns the LimitWorkspaces field if non-nil, zero value otherwise.

### GetLimitWorkspacesOk

`func (o *AutomationToken) GetLimitWorkspacesOk() (*bool, bool)`

GetLimitWorkspacesOk returns a tuple with the LimitWorkspaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitWorkspaces

`func (o *AutomationToken) SetLimitWorkspaces(v bool)`

SetLimitWorkspaces sets LimitWorkspaces field to given value.

### HasLimitWorkspaces

`func (o *AutomationToken) HasLimitWorkspaces() bool`

HasLimitWorkspaces returns a boolean if a field has been set.

### GetWorkspaces

`func (o *AutomationToken) GetWorkspaces() []string`

GetWorkspaces returns the Workspaces field if non-nil, zero value otherwise.

### GetWorkspacesOk

`func (o *AutomationToken) GetWorkspacesOk() (*[]string, bool)`

GetWorkspacesOk returns a tuple with the Workspaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaces

`func (o *AutomationToken) SetWorkspaces(v []string)`

SetWorkspaces sets Workspaces field to given value.

### HasWorkspaces

`func (o *AutomationToken) HasWorkspaces() bool`

HasWorkspaces returns a boolean if a field has been set.

### GetScope

`func (o *AutomationToken) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *AutomationToken) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *AutomationToken) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *AutomationToken) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetExpiresAt

`func (o *AutomationToken) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *AutomationToken) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *AutomationToken) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *AutomationToken) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


