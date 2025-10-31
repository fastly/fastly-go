# AutomationTokenResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the token. | [optional] 
**Role** | Pointer to **string** |  | [optional] 
**Services** | Pointer to **[]string** | (Optional) The service IDs of the services the token will have access to. Separate service IDs with a space. If no services are specified, the token will have access to all services on the account.  | [optional] 
**Scope** | Pointer to **string** | A space-delimited list of authorization scope. | [optional] [default to "global"]
**ExpiresAt** | Pointer to **string** | (optional) A UTC timestamp of when the token will expire. | [optional] 
**Id** | Pointer to [**ReadOnlyId**](ReadOnlyId.md) |  | [optional] 
**CustomerId** | Pointer to [**ReadOnlyCustomerId**](ReadOnlyCustomerId.md) |  | [optional] 
**Ip** | Pointer to **string** | The IP address of the client that last used the token. | [optional] 
**UserAgent** | Pointer to **string** | The User-Agent header of the client that last used the token. | [optional] 
**TlsAccess** | Pointer to **bool** | Indicates whether TLS access is enabled for the token. | [optional] 
**LastUsedAt** | Pointer to **time.Time** | A UTC timestamp of when the token was last used. | [optional] [readonly] 
**CreatedAt** | Pointer to **string** | A UTC timestamp of when the token was created. | [optional] 

## Methods

### NewAutomationTokenResponse

`func NewAutomationTokenResponse() *AutomationTokenResponse`

NewAutomationTokenResponse instantiates a new AutomationTokenResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutomationTokenResponseWithDefaults

`func NewAutomationTokenResponseWithDefaults() *AutomationTokenResponse`

NewAutomationTokenResponseWithDefaults instantiates a new AutomationTokenResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AutomationTokenResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AutomationTokenResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AutomationTokenResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AutomationTokenResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRole

`func (o *AutomationTokenResponse) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *AutomationTokenResponse) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *AutomationTokenResponse) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *AutomationTokenResponse) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetServices

`func (o *AutomationTokenResponse) GetServices() []string`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *AutomationTokenResponse) GetServicesOk() (*[]string, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *AutomationTokenResponse) SetServices(v []string)`

SetServices sets Services field to given value.

### HasServices

`func (o *AutomationTokenResponse) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetScope

`func (o *AutomationTokenResponse) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *AutomationTokenResponse) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *AutomationTokenResponse) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *AutomationTokenResponse) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetExpiresAt

`func (o *AutomationTokenResponse) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *AutomationTokenResponse) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *AutomationTokenResponse) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *AutomationTokenResponse) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetId

`func (o *AutomationTokenResponse) GetId() ReadOnlyId`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AutomationTokenResponse) GetIdOk() (*ReadOnlyId, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AutomationTokenResponse) SetId(v ReadOnlyId)`

SetId sets Id field to given value.

### HasId

`func (o *AutomationTokenResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCustomerId

`func (o *AutomationTokenResponse) GetCustomerId() ReadOnlyCustomerId`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *AutomationTokenResponse) GetCustomerIdOk() (*ReadOnlyCustomerId, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *AutomationTokenResponse) SetCustomerId(v ReadOnlyCustomerId)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *AutomationTokenResponse) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetIp

`func (o *AutomationTokenResponse) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *AutomationTokenResponse) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *AutomationTokenResponse) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *AutomationTokenResponse) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetUserAgent

`func (o *AutomationTokenResponse) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *AutomationTokenResponse) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *AutomationTokenResponse) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *AutomationTokenResponse) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.

### GetTlsAccess

`func (o *AutomationTokenResponse) GetTlsAccess() bool`

GetTlsAccess returns the TlsAccess field if non-nil, zero value otherwise.

### GetTlsAccessOk

`func (o *AutomationTokenResponse) GetTlsAccessOk() (*bool, bool)`

GetTlsAccessOk returns a tuple with the TlsAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsAccess

`func (o *AutomationTokenResponse) SetTlsAccess(v bool)`

SetTlsAccess sets TlsAccess field to given value.

### HasTlsAccess

`func (o *AutomationTokenResponse) HasTlsAccess() bool`

HasTlsAccess returns a boolean if a field has been set.

### GetLastUsedAt

`func (o *AutomationTokenResponse) GetLastUsedAt() time.Time`

GetLastUsedAt returns the LastUsedAt field if non-nil, zero value otherwise.

### GetLastUsedAtOk

`func (o *AutomationTokenResponse) GetLastUsedAtOk() (*time.Time, bool)`

GetLastUsedAtOk returns a tuple with the LastUsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedAt

`func (o *AutomationTokenResponse) SetLastUsedAt(v time.Time)`

SetLastUsedAt sets LastUsedAt field to given value.

### HasLastUsedAt

`func (o *AutomationTokenResponse) HasLastUsedAt() bool`

HasLastUsedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AutomationTokenResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AutomationTokenResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AutomationTokenResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AutomationTokenResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


