# AutomationTokenCreateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the token. | [optional] 
**Role** | Pointer to **string** | The role on the token. | [optional] 
**Services** | Pointer to **[]string** | (Optional) The service IDs of the services the token will have access to. Separate service IDs with a space. If no services are specified, the token will have access to all services on the account.  | [optional] 
**Scope** | Pointer to **string** | A space-delimited list of authorization scope. | [optional] [default to "global"]
**ExpiresAt** | Pointer to **string** | A UTC timestamp of when the token expires. | [optional] 
**ID** | Pointer to [**ReadOnlyID**](ReadOnlyID.md) |  | [optional] 
**UserID** | Pointer to [**ReadOnlyUserID**](ReadOnlyUserID.md) |  | [optional] 
**CustomerID** | Pointer to [**ReadOnlyCustomerID**](ReadOnlyCustomerID.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** | A UTC timestamp of when the token was created.  | [optional] [readonly] 
**AccessToken** | Pointer to **string** |  | [optional] [readonly] 
**TLSAccess** | Pointer to **bool** | Indicates whether TLS access is enabled for the token. | [optional] 
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

### GetID

`func (o *AutomationTokenCreateResponse) GetID() ReadOnlyID`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *AutomationTokenCreateResponse) GetIDOk() (*ReadOnlyID, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *AutomationTokenCreateResponse) SetID(v ReadOnlyID)`

SetID sets ID field to given value.

### HasID

`func (o *AutomationTokenCreateResponse) HasID() bool`

HasID returns a boolean if a field has been set.

### GetUserID

`func (o *AutomationTokenCreateResponse) GetUserID() ReadOnlyUserID`

GetUserID returns the UserID field if non-nil, zero value otherwise.

### GetUserIDOk

`func (o *AutomationTokenCreateResponse) GetUserIDOk() (*ReadOnlyUserID, bool)`

GetUserIDOk returns a tuple with the UserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserID

`func (o *AutomationTokenCreateResponse) SetUserID(v ReadOnlyUserID)`

SetUserID sets UserID field to given value.

### HasUserID

`func (o *AutomationTokenCreateResponse) HasUserID() bool`

HasUserID returns a boolean if a field has been set.

### GetCustomerID

`func (o *AutomationTokenCreateResponse) GetCustomerID() ReadOnlyCustomerID`

GetCustomerID returns the CustomerID field if non-nil, zero value otherwise.

### GetCustomerIDOk

`func (o *AutomationTokenCreateResponse) GetCustomerIDOk() (*ReadOnlyCustomerID, bool)`

GetCustomerIDOk returns a tuple with the CustomerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerID

`func (o *AutomationTokenCreateResponse) SetCustomerID(v ReadOnlyCustomerID)`

SetCustomerID sets CustomerID field to given value.

### HasCustomerID

`func (o *AutomationTokenCreateResponse) HasCustomerID() bool`

HasCustomerID returns a boolean if a field has been set.

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

### GetTLSAccess

`func (o *AutomationTokenCreateResponse) GetTLSAccess() bool`

GetTLSAccess returns the TLSAccess field if non-nil, zero value otherwise.

### GetTLSAccessOk

`func (o *AutomationTokenCreateResponse) GetTLSAccessOk() (*bool, bool)`

GetTLSAccessOk returns a tuple with the TLSAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTLSAccess

`func (o *AutomationTokenCreateResponse) SetTLSAccess(v bool)`

SetTLSAccess sets TLSAccess field to given value.

### HasTLSAccess

`func (o *AutomationTokenCreateResponse) HasTLSAccess() bool`

HasTLSAccess returns a boolean if a field has been set.

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
