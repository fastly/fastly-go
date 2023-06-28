# AutomationTokenResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the token. | [optional] 
**Role** | Pointer to **string** |  | [optional] 
**Services** | Pointer to **[]string** | (Optional) The service IDs of the services the token will have access to. Separate service IDs with a space. If no services are specified, the token will have access to all services on the account.  | [optional] 
**Scope** | Pointer to **string** | A space-delimited list of authorization scope. | [optional] [default to "global"]
**ExpiresAt** | Pointer to **string** | (optional) A UTC time-stamp of when the token will expire. | [optional] 
**CreatedAt** | Pointer to **string** | A UTC time-stamp of when the token was created. | [optional] 
**DeletedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**UpdatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**ID** | Pointer to [**ReadOnlyID**](ReadOnlyID.md) |  | [optional] 
**CustomerID** | Pointer to [**ReadOnlyCustomerID**](ReadOnlyCustomerID.md) |  | [optional] 
**IP** | Pointer to **string** | The IP address of the client that last used the token. | [optional] 
**UserAgent** | Pointer to **string** | The User-Agent header of the client that last used the token. | [optional] 
**SudoExpiresAt** | Pointer to **string** |  | [optional] [readonly] 
**LastUsedAt** | Pointer to **time.Time** | A UTC time-stamp of when the token was last used. | [optional] [readonly] 

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

### GetDeletedAt

`func (o *AutomationTokenResponse) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *AutomationTokenResponse) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *AutomationTokenResponse) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *AutomationTokenResponse) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *AutomationTokenResponse) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *AutomationTokenResponse) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetUpdatedAt

`func (o *AutomationTokenResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AutomationTokenResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AutomationTokenResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AutomationTokenResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *AutomationTokenResponse) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *AutomationTokenResponse) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetID

`func (o *AutomationTokenResponse) GetID() ReadOnlyID`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *AutomationTokenResponse) GetIDOk() (*ReadOnlyID, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *AutomationTokenResponse) SetID(v ReadOnlyID)`

SetID sets ID field to given value.

### HasID

`func (o *AutomationTokenResponse) HasID() bool`

HasID returns a boolean if a field has been set.

### GetCustomerID

`func (o *AutomationTokenResponse) GetCustomerID() ReadOnlyCustomerID`

GetCustomerID returns the CustomerID field if non-nil, zero value otherwise.

### GetCustomerIDOk

`func (o *AutomationTokenResponse) GetCustomerIDOk() (*ReadOnlyCustomerID, bool)`

GetCustomerIDOk returns a tuple with the CustomerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerID

`func (o *AutomationTokenResponse) SetCustomerID(v ReadOnlyCustomerID)`

SetCustomerID sets CustomerID field to given value.

### HasCustomerID

`func (o *AutomationTokenResponse) HasCustomerID() bool`

HasCustomerID returns a boolean if a field has been set.

### GetIP

`func (o *AutomationTokenResponse) GetIP() string`

GetIP returns the IP field if non-nil, zero value otherwise.

### GetIPOk

`func (o *AutomationTokenResponse) GetIPOk() (*string, bool)`

GetIPOk returns a tuple with the IP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIP

`func (o *AutomationTokenResponse) SetIP(v string)`

SetIP sets IP field to given value.

### HasIP

`func (o *AutomationTokenResponse) HasIP() bool`

HasIP returns a boolean if a field has been set.

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

### GetSudoExpiresAt

`func (o *AutomationTokenResponse) GetSudoExpiresAt() string`

GetSudoExpiresAt returns the SudoExpiresAt field if non-nil, zero value otherwise.

### GetSudoExpiresAtOk

`func (o *AutomationTokenResponse) GetSudoExpiresAtOk() (*string, bool)`

GetSudoExpiresAtOk returns a tuple with the SudoExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSudoExpiresAt

`func (o *AutomationTokenResponse) SetSudoExpiresAt(v string)`

SetSudoExpiresAt sets SudoExpiresAt field to given value.

### HasSudoExpiresAt

`func (o *AutomationTokenResponse) HasSudoExpiresAt() bool`

HasSudoExpiresAt returns a boolean if a field has been set.

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


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
