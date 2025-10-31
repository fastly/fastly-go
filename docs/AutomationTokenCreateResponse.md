# AutomationTokenCreateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the token. | [optional] 
**Role** | Pointer to **string** | The role on the token. | [optional] 
**Services** | Pointer to **[]string** | (Optional) The service IDs of the services the token will have access to. Separate service IDs with a space. If no services are specified, the token will have access to all services on the account.  | [optional] 
**Scope** | Pointer to **string** | A space-delimited list of authorization scope. | [optional] [default to "global"]
**ExpiresAt** | Pointer to **string** | A UTC timestamp of when the token expires. | [optional] 
**Id** | Pointer to [**ReadOnlyId**](ReadOnlyId.md) |  | [optional] 
**UserId** | Pointer to [**ReadOnlyUserId**](ReadOnlyUserId.md) |  | [optional] 
**CustomerId** | Pointer to [**ReadOnlyCustomerId**](ReadOnlyCustomerId.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** | A UTC timestamp of when the token was created.  | [optional] [readonly] 
**AccessToken** | Pointer to **string** |  | [optional] [readonly] 
**TlsAccess** | Pointer to **bool** | Indicates whether TLS access is enabled for the token. | [optional] 
**LastUsedAt** | Pointer to **time.Time** | A UTC timestamp of when the token was last used. | [optional] [readonly] 
**UserAgent** | Pointer to **string** | The User-Agent header of the client that last used the token. | [optional] 

## Methods

### NewAutomationTokenCreateResponse

`func NewAutomationTokenCreateResponse() *AutomationTokenCreateResponse`

NewAutomationTokenCreateResponse instantiates a new AutomationTokenCreateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutomationTokenCreateResponseWithDefaults

`func NewAutomationTokenCreateResponseWithDefaults() *AutomationTokenCreateResponse`

NewAutomationTokenCreateResponseWithDefaults instantiates a new AutomationTokenCreateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AutomationTokenCreateResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AutomationTokenCreateResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AutomationTokenCreateResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AutomationTokenCreateResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRole

`func (o *AutomationTokenCreateResponse) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *AutomationTokenCreateResponse) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *AutomationTokenCreateResponse) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *AutomationTokenCreateResponse) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetServices

`func (o *AutomationTokenCreateResponse) GetServices() []string`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *AutomationTokenCreateResponse) GetServicesOk() (*[]string, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *AutomationTokenCreateResponse) SetServices(v []string)`

SetServices sets Services field to given value.

### HasServices

`func (o *AutomationTokenCreateResponse) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetScope

`func (o *AutomationTokenCreateResponse) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *AutomationTokenCreateResponse) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *AutomationTokenCreateResponse) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *AutomationTokenCreateResponse) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetExpiresAt

`func (o *AutomationTokenCreateResponse) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *AutomationTokenCreateResponse) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *AutomationTokenCreateResponse) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *AutomationTokenCreateResponse) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetId

`func (o *AutomationTokenCreateResponse) GetId() ReadOnlyId`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AutomationTokenCreateResponse) GetIdOk() (*ReadOnlyId, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AutomationTokenCreateResponse) SetId(v ReadOnlyId)`

SetId sets Id field to given value.

### HasId

`func (o *AutomationTokenCreateResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUserId

`func (o *AutomationTokenCreateResponse) GetUserId() ReadOnlyUserId`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *AutomationTokenCreateResponse) GetUserIdOk() (*ReadOnlyUserId, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *AutomationTokenCreateResponse) SetUserId(v ReadOnlyUserId)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *AutomationTokenCreateResponse) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetCustomerId

`func (o *AutomationTokenCreateResponse) GetCustomerId() ReadOnlyCustomerId`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *AutomationTokenCreateResponse) GetCustomerIdOk() (*ReadOnlyCustomerId, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *AutomationTokenCreateResponse) SetCustomerId(v ReadOnlyCustomerId)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *AutomationTokenCreateResponse) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AutomationTokenCreateResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AutomationTokenCreateResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AutomationTokenCreateResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AutomationTokenCreateResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetAccessToken

`func (o *AutomationTokenCreateResponse) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *AutomationTokenCreateResponse) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *AutomationTokenCreateResponse) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *AutomationTokenCreateResponse) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetTlsAccess

`func (o *AutomationTokenCreateResponse) GetTlsAccess() bool`

GetTlsAccess returns the TlsAccess field if non-nil, zero value otherwise.

### GetTlsAccessOk

`func (o *AutomationTokenCreateResponse) GetTlsAccessOk() (*bool, bool)`

GetTlsAccessOk returns a tuple with the TlsAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsAccess

`func (o *AutomationTokenCreateResponse) SetTlsAccess(v bool)`

SetTlsAccess sets TlsAccess field to given value.

### HasTlsAccess

`func (o *AutomationTokenCreateResponse) HasTlsAccess() bool`

HasTlsAccess returns a boolean if a field has been set.

### GetLastUsedAt

`func (o *AutomationTokenCreateResponse) GetLastUsedAt() time.Time`

GetLastUsedAt returns the LastUsedAt field if non-nil, zero value otherwise.

### GetLastUsedAtOk

`func (o *AutomationTokenCreateResponse) GetLastUsedAtOk() (*time.Time, bool)`

GetLastUsedAtOk returns a tuple with the LastUsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedAt

`func (o *AutomationTokenCreateResponse) SetLastUsedAt(v time.Time)`

SetLastUsedAt sets LastUsedAt field to given value.

### HasLastUsedAt

`func (o *AutomationTokenCreateResponse) HasLastUsedAt() bool`

HasLastUsedAt returns a boolean if a field has been set.

### GetUserAgent

`func (o *AutomationTokenCreateResponse) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *AutomationTokenCreateResponse) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *AutomationTokenCreateResponse) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *AutomationTokenCreateResponse) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


