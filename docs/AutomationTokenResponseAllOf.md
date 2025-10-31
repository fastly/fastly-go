# AutomationTokenResponseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to [**ReadOnlyId**](ReadOnlyId.md) |  | [optional] 
**CustomerId** | Pointer to [**ReadOnlyCustomerId**](ReadOnlyCustomerId.md) |  | [optional] 
**Role** | Pointer to **string** |  | [optional] 
**Ip** | Pointer to **string** | The IP address of the client that last used the token. | [optional] 
**UserAgent** | Pointer to **string** | The User-Agent header of the client that last used the token. | [optional] 
**TlsAccess** | Pointer to **bool** | Indicates whether TLS access is enabled for the token. | [optional] 
**LastUsedAt** | Pointer to **time.Time** | A UTC timestamp of when the token was last used. | [optional] [readonly] 
**CreatedAt** | Pointer to **string** | A UTC timestamp of when the token was created. | [optional] 
**ExpiresAt** | Pointer to **string** | (optional) A UTC timestamp of when the token will expire. | [optional] 

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

### GetId

`func (o *AutomationTokenResponseAllOf) GetId() ReadOnlyId`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AutomationTokenResponseAllOf) GetIdOk() (*ReadOnlyId, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AutomationTokenResponseAllOf) SetId(v ReadOnlyId)`

SetId sets Id field to given value.

### HasId

`func (o *AutomationTokenResponseAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCustomerId

`func (o *AutomationTokenResponseAllOf) GetCustomerId() ReadOnlyCustomerId`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *AutomationTokenResponseAllOf) GetCustomerIdOk() (*ReadOnlyCustomerId, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *AutomationTokenResponseAllOf) SetCustomerId(v ReadOnlyCustomerId)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *AutomationTokenResponseAllOf) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

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

### GetIp

`func (o *AutomationTokenResponseAllOf) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *AutomationTokenResponseAllOf) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *AutomationTokenResponseAllOf) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *AutomationTokenResponseAllOf) HasIp() bool`

HasIp returns a boolean if a field has been set.

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

### GetTlsAccess

`func (o *AutomationTokenResponseAllOf) GetTlsAccess() bool`

GetTlsAccess returns the TlsAccess field if non-nil, zero value otherwise.

### GetTlsAccessOk

`func (o *AutomationTokenResponseAllOf) GetTlsAccessOk() (*bool, bool)`

GetTlsAccessOk returns a tuple with the TlsAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsAccess

`func (o *AutomationTokenResponseAllOf) SetTlsAccess(v bool)`

SetTlsAccess sets TlsAccess field to given value.

### HasTlsAccess

`func (o *AutomationTokenResponseAllOf) HasTlsAccess() bool`

HasTlsAccess returns a boolean if a field has been set.

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


