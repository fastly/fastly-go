# AutomationTokenCreateResponseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to [**ReadOnlyId**](ReadOnlyId.md) |  | [optional] 
**UserId** | Pointer to [**ReadOnlyUserId**](ReadOnlyUserId.md) |  | [optional] 
**CustomerId** | Pointer to [**ReadOnlyCustomerId**](ReadOnlyCustomerId.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** | A UTC timestamp of when the token was created.  | [optional] [readonly] 
**AccessToken** | Pointer to **string** |  | [optional] [readonly] 
**TlsAccess** | Pointer to **bool** | Indicates whether TLS access is enabled for the token. | [optional] 
**LastUsedAt** | Pointer to **time.Time** | A UTC timestamp of when the token was last used. | [optional] [readonly] 
**UserAgent** | Pointer to **string** | The User-Agent header of the client that last used the token. | [optional] 

## Methods

### NewAutomationTokenCreateResponseAllOf

`func NewAutomationTokenCreateResponseAllOf() *AutomationTokenCreateResponseAllOf`

NewAutomationTokenCreateResponseAllOf instantiates a new AutomationTokenCreateResponseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutomationTokenCreateResponseAllOfWithDefaults

`func NewAutomationTokenCreateResponseAllOfWithDefaults() *AutomationTokenCreateResponseAllOf`

NewAutomationTokenCreateResponseAllOfWithDefaults instantiates a new AutomationTokenCreateResponseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AutomationTokenCreateResponseAllOf) GetId() ReadOnlyId`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AutomationTokenCreateResponseAllOf) GetIdOk() (*ReadOnlyId, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AutomationTokenCreateResponseAllOf) SetId(v ReadOnlyId)`

SetId sets Id field to given value.

### HasId

`func (o *AutomationTokenCreateResponseAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUserId

`func (o *AutomationTokenCreateResponseAllOf) GetUserId() ReadOnlyUserId`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *AutomationTokenCreateResponseAllOf) GetUserIdOk() (*ReadOnlyUserId, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *AutomationTokenCreateResponseAllOf) SetUserId(v ReadOnlyUserId)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *AutomationTokenCreateResponseAllOf) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetCustomerId

`func (o *AutomationTokenCreateResponseAllOf) GetCustomerId() ReadOnlyCustomerId`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *AutomationTokenCreateResponseAllOf) GetCustomerIdOk() (*ReadOnlyCustomerId, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *AutomationTokenCreateResponseAllOf) SetCustomerId(v ReadOnlyCustomerId)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *AutomationTokenCreateResponseAllOf) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AutomationTokenCreateResponseAllOf) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AutomationTokenCreateResponseAllOf) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AutomationTokenCreateResponseAllOf) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AutomationTokenCreateResponseAllOf) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetAccessToken

`func (o *AutomationTokenCreateResponseAllOf) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *AutomationTokenCreateResponseAllOf) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *AutomationTokenCreateResponseAllOf) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *AutomationTokenCreateResponseAllOf) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetTlsAccess

`func (o *AutomationTokenCreateResponseAllOf) GetTlsAccess() bool`

GetTlsAccess returns the TlsAccess field if non-nil, zero value otherwise.

### GetTlsAccessOk

`func (o *AutomationTokenCreateResponseAllOf) GetTlsAccessOk() (*bool, bool)`

GetTlsAccessOk returns a tuple with the TlsAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsAccess

`func (o *AutomationTokenCreateResponseAllOf) SetTlsAccess(v bool)`

SetTlsAccess sets TlsAccess field to given value.

### HasTlsAccess

`func (o *AutomationTokenCreateResponseAllOf) HasTlsAccess() bool`

HasTlsAccess returns a boolean if a field has been set.

### GetLastUsedAt

`func (o *AutomationTokenCreateResponseAllOf) GetLastUsedAt() time.Time`

GetLastUsedAt returns the LastUsedAt field if non-nil, zero value otherwise.

### GetLastUsedAtOk

`func (o *AutomationTokenCreateResponseAllOf) GetLastUsedAtOk() (*time.Time, bool)`

GetLastUsedAtOk returns a tuple with the LastUsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedAt

`func (o *AutomationTokenCreateResponseAllOf) SetLastUsedAt(v time.Time)`

SetLastUsedAt sets LastUsedAt field to given value.

### HasLastUsedAt

`func (o *AutomationTokenCreateResponseAllOf) HasLastUsedAt() bool`

HasLastUsedAt returns a boolean if a field has been set.

### GetUserAgent

`func (o *AutomationTokenCreateResponseAllOf) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *AutomationTokenCreateResponseAllOf) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *AutomationTokenCreateResponseAllOf) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *AutomationTokenCreateResponseAllOf) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


