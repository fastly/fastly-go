# AutomationTokenResponseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ID** | Pointer to [**ReadOnlyID**](ReadOnlyID.md) |  | [optional] 
**CustomerID** | Pointer to [**ReadOnlyCustomerID**](ReadOnlyCustomerID.md) |  | [optional] 
**Role** | Pointer to **string** |  | [optional] 
**IP** | Pointer to **string** | The IP address of the client that last used the token. | [optional] 
**UserAgent** | Pointer to **string** | The User-Agent header of the client that last used the token. | [optional] 
**SudoExpiresAt** | Pointer to **string** |  | [optional] [readonly] 
**LastUsedAt** | Pointer to **time.Time** | A UTC time-stamp of when the token was last used. | [optional] [readonly] 
**CreatedAt** | Pointer to **string** | A UTC time-stamp of when the token was created. | [optional] 
**ExpiresAt** | Pointer to **string** | (optional) A UTC time-stamp of when the token will expire. | [optional] 

## Methods

### NewAutomationTokenResponseAllOf

`func NewAutomationTokenResponseAllOf() *AutomationTokenResponseAllOf`

NewAutomationTokenResponseAllOf instantiates a new AutomationTokenResponseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutomationTokenResponseAllOfWithDefaults

`func NewAutomationTokenResponseAllOfWithDefaults() *AutomationTokenResponseAllOf`

NewAutomationTokenResponseAllOfWithDefaults instantiates a new AutomationTokenResponseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetID

`func (o *AutomationTokenResponseAllOf) GetID() ReadOnlyID`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *AutomationTokenResponseAllOf) GetIDOk() (*ReadOnlyID, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *AutomationTokenResponseAllOf) SetID(v ReadOnlyID)`

SetID sets ID field to given value.

### HasID

`func (o *AutomationTokenResponseAllOf) HasID() bool`

HasID returns a boolean if a field has been set.

### GetCustomerID

`func (o *AutomationTokenResponseAllOf) GetCustomerID() ReadOnlyCustomerID`

GetCustomerID returns the CustomerID field if non-nil, zero value otherwise.

### GetCustomerIDOk

`func (o *AutomationTokenResponseAllOf) GetCustomerIDOk() (*ReadOnlyCustomerID, bool)`

GetCustomerIDOk returns a tuple with the CustomerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerID

`func (o *AutomationTokenResponseAllOf) SetCustomerID(v ReadOnlyCustomerID)`

SetCustomerID sets CustomerID field to given value.

### HasCustomerID

`func (o *AutomationTokenResponseAllOf) HasCustomerID() bool`

HasCustomerID returns a boolean if a field has been set.

### GetRole

`func (o *AutomationTokenResponseAllOf) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *AutomationTokenResponseAllOf) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *AutomationTokenResponseAllOf) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *AutomationTokenResponseAllOf) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetIP

`func (o *AutomationTokenResponseAllOf) GetIP() string`

GetIP returns the IP field if non-nil, zero value otherwise.

### GetIPOk

`func (o *AutomationTokenResponseAllOf) GetIPOk() (*string, bool)`

GetIPOk returns a tuple with the IP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIP

`func (o *AutomationTokenResponseAllOf) SetIP(v string)`

SetIP sets IP field to given value.

### HasIP

`func (o *AutomationTokenResponseAllOf) HasIP() bool`

HasIP returns a boolean if a field has been set.

### GetUserAgent

`func (o *AutomationTokenResponseAllOf) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *AutomationTokenResponseAllOf) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *AutomationTokenResponseAllOf) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *AutomationTokenResponseAllOf) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.

### GetSudoExpiresAt

`func (o *AutomationTokenResponseAllOf) GetSudoExpiresAt() string`

GetSudoExpiresAt returns the SudoExpiresAt field if non-nil, zero value otherwise.

### GetSudoExpiresAtOk

`func (o *AutomationTokenResponseAllOf) GetSudoExpiresAtOk() (*string, bool)`

GetSudoExpiresAtOk returns a tuple with the SudoExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSudoExpiresAt

`func (o *AutomationTokenResponseAllOf) SetSudoExpiresAt(v string)`

SetSudoExpiresAt sets SudoExpiresAt field to given value.

### HasSudoExpiresAt

`func (o *AutomationTokenResponseAllOf) HasSudoExpiresAt() bool`

HasSudoExpiresAt returns a boolean if a field has been set.

### GetLastUsedAt

`func (o *AutomationTokenResponseAllOf) GetLastUsedAt() time.Time`

GetLastUsedAt returns the LastUsedAt field if non-nil, zero value otherwise.

### GetLastUsedAtOk

`func (o *AutomationTokenResponseAllOf) GetLastUsedAtOk() (*time.Time, bool)`

GetLastUsedAtOk returns a tuple with the LastUsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedAt

`func (o *AutomationTokenResponseAllOf) SetLastUsedAt(v time.Time)`

SetLastUsedAt sets LastUsedAt field to given value.

### HasLastUsedAt

`func (o *AutomationTokenResponseAllOf) HasLastUsedAt() bool`

HasLastUsedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AutomationTokenResponseAllOf) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AutomationTokenResponseAllOf) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AutomationTokenResponseAllOf) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AutomationTokenResponseAllOf) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetExpiresAt

`func (o *AutomationTokenResponseAllOf) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *AutomationTokenResponseAllOf) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *AutomationTokenResponseAllOf) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *AutomationTokenResponseAllOf) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
