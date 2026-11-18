# AutomationTokenCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the token. | 
**Role** | Pointer to **string** | The role for the token. The `roles` field is the preferred way to assign roles. This field is supported for backwards compatibility with single-role tokens. | [optional] 
**Roles** | Pointer to **[]string** | Array of role IDs to assign to the token. Use this field for tokens with multiple roles. When provided, this takes precedence over the `role` field. To get available role IDs, use [GET /roles](https://www.fastly.com/documentation/reference/api/account/roles/). | [optional] 
**Services** | Pointer to **[]string** | List of service IDs the token can access. Only used when `limit_services` is `true`. | [optional] 
**LimitServices** | Pointer to **bool** | When `true`, restricts the token to only the services specified in the `services` array. When `false` (default), the token has access to all services. | [optional] [default to false]
**LimitWorkspaces** | Pointer to **bool** | When `true` (default), restricts the token to only the workspaces specified in the `workspaces` array. When `false`, the token has access to all workspaces. Only applies to roles that support workspace access (such as Next-Gen WAF roles). For roles without workspace support, this field has no effect. | [optional] [default to true]
**Workspaces** | Pointer to **[]string** | List of workspace IDs the token can access. Only used when `limit_workspaces` is `true` and the token has roles that support workspace access. | [optional] 
**Scope** | Pointer to **string** |  | [optional] [default to "global"]
**ExpiresAt** | **NullableTime** | A UTC timestamp of when the token will expire. | 
**TLSAccess** | Pointer to **bool** | Indicates whether TLS access is enabled for the token. Only applicable when using the single-role `role` field. For tokens with multiple roles, assign TLS roles (`tls_admin` or `tls_viewer`) via the `roles` field instead. | [optional] 

## Methods

### NewAutomationTokenCreateRequest

`func NewAutomationTokenCreateRequest(name string, expiresAt NullableTime, ) *AutomationTokenCreateRequest`

NewAutomationTokenCreateRequest instantiates a new AutomationTokenCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutomationTokenCreateRequestWithDefaults

`func NewAutomationTokenCreateRequestWithDefaults() *AutomationTokenCreateRequest`

NewAutomationTokenCreateRequestWithDefaults instantiates a new AutomationTokenCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AutomationTokenCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AutomationTokenCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AutomationTokenCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetRole

`func (o *AutomationTokenCreateRequest) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *AutomationTokenCreateRequest) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *AutomationTokenCreateRequest) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *AutomationTokenCreateRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetRoles

`func (o *AutomationTokenCreateRequest) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *AutomationTokenCreateRequest) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *AutomationTokenCreateRequest) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *AutomationTokenCreateRequest) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetServices

`func (o *AutomationTokenCreateRequest) GetServices() []string`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *AutomationTokenCreateRequest) GetServicesOk() (*[]string, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *AutomationTokenCreateRequest) SetServices(v []string)`

SetServices sets Services field to given value.

### HasServices

`func (o *AutomationTokenCreateRequest) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetLimitServices

`func (o *AutomationTokenCreateRequest) GetLimitServices() bool`

GetLimitServices returns the LimitServices field if non-nil, zero value otherwise.

### GetLimitServicesOk

`func (o *AutomationTokenCreateRequest) GetLimitServicesOk() (*bool, bool)`

GetLimitServicesOk returns a tuple with the LimitServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitServices

`func (o *AutomationTokenCreateRequest) SetLimitServices(v bool)`

SetLimitServices sets LimitServices field to given value.

### HasLimitServices

`func (o *AutomationTokenCreateRequest) HasLimitServices() bool`

HasLimitServices returns a boolean if a field has been set.

### GetLimitWorkspaces

`func (o *AutomationTokenCreateRequest) GetLimitWorkspaces() bool`

GetLimitWorkspaces returns the LimitWorkspaces field if non-nil, zero value otherwise.

### GetLimitWorkspacesOk

`func (o *AutomationTokenCreateRequest) GetLimitWorkspacesOk() (*bool, bool)`

GetLimitWorkspacesOk returns a tuple with the LimitWorkspaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitWorkspaces

`func (o *AutomationTokenCreateRequest) SetLimitWorkspaces(v bool)`

SetLimitWorkspaces sets LimitWorkspaces field to given value.

### HasLimitWorkspaces

`func (o *AutomationTokenCreateRequest) HasLimitWorkspaces() bool`

HasLimitWorkspaces returns a boolean if a field has been set.

### GetWorkspaces

`func (o *AutomationTokenCreateRequest) GetWorkspaces() []string`

GetWorkspaces returns the Workspaces field if non-nil, zero value otherwise.

### GetWorkspacesOk

`func (o *AutomationTokenCreateRequest) GetWorkspacesOk() (*[]string, bool)`

GetWorkspacesOk returns a tuple with the Workspaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaces

`func (o *AutomationTokenCreateRequest) SetWorkspaces(v []string)`

SetWorkspaces sets Workspaces field to given value.

### HasWorkspaces

`func (o *AutomationTokenCreateRequest) HasWorkspaces() bool`

HasWorkspaces returns a boolean if a field has been set.

### GetScope

`func (o *AutomationTokenCreateRequest) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *AutomationTokenCreateRequest) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *AutomationTokenCreateRequest) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *AutomationTokenCreateRequest) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetExpiresAt

`func (o *AutomationTokenCreateRequest) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *AutomationTokenCreateRequest) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *AutomationTokenCreateRequest) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.


### SetExpiresAtNil

`func (o *AutomationTokenCreateRequest) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *AutomationTokenCreateRequest) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetTLSAccess

`func (o *AutomationTokenCreateRequest) GetTLSAccess() bool`

GetTLSAccess returns the TLSAccess field if non-nil, zero value otherwise.

### GetTLSAccessOk

`func (o *AutomationTokenCreateRequest) GetTLSAccessOk() (*bool, bool)`

GetTLSAccessOk returns a tuple with the TLSAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTLSAccess

`func (o *AutomationTokenCreateRequest) SetTLSAccess(v bool)`

SetTLSAccess sets TLSAccess field to given value.

### HasTLSAccess

`func (o *AutomationTokenCreateRequest) HasTLSAccess() bool`

HasTLSAccess returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
