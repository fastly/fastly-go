# AutomationTokenCreateResponseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ID** | Pointer to [**ReadOnlyID**](ReadOnlyID.md) |  | [optional] 
**UserID** | Pointer to [**ReadOnlyUserID**](ReadOnlyUserID.md) |  | [optional] 
**CustomerID** | Pointer to [**ReadOnlyCustomerID**](ReadOnlyCustomerID.md) |  | [optional] 
**SudoExpiresAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** | A UTC time-stamp of when the token was created.  | [optional] [readonly] 
**AccessToken** | Pointer to **string** |  | [optional] [readonly] 
**LastUsedAt** | Pointer to **time.Time** | A UTC time-stamp of when the token was last used. | [optional] [readonly] 
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

### GetID

`func (o *AutomationTokenCreateResponseAllOf) GetID() ReadOnlyID`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *AutomationTokenCreateResponseAllOf) GetIDOk() (*ReadOnlyID, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *AutomationTokenCreateResponseAllOf) SetID(v ReadOnlyID)`

SetID sets ID field to given value.

### HasID

`func (o *AutomationTokenCreateResponseAllOf) HasID() bool`

HasID returns a boolean if a field has been set.

### GetUserID

`func (o *AutomationTokenCreateResponseAllOf) GetUserID() ReadOnlyUserID`

GetUserID returns the UserID field if non-nil, zero value otherwise.

### GetUserIDOk

`func (o *AutomationTokenCreateResponseAllOf) GetUserIDOk() (*ReadOnlyUserID, bool)`

GetUserIDOk returns a tuple with the UserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserID

`func (o *AutomationTokenCreateResponseAllOf) SetUserID(v ReadOnlyUserID)`

SetUserID sets UserID field to given value.

### HasUserID

`func (o *AutomationTokenCreateResponseAllOf) HasUserID() bool`

HasUserID returns a boolean if a field has been set.

### GetCustomerID

`func (o *AutomationTokenCreateResponseAllOf) GetCustomerID() ReadOnlyCustomerID`

GetCustomerID returns the CustomerID field if non-nil, zero value otherwise.

### GetCustomerIDOk

`func (o *AutomationTokenCreateResponseAllOf) GetCustomerIDOk() (*ReadOnlyCustomerID, bool)`

GetCustomerIDOk returns a tuple with the CustomerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerID

`func (o *AutomationTokenCreateResponseAllOf) SetCustomerID(v ReadOnlyCustomerID)`

SetCustomerID sets CustomerID field to given value.

### HasCustomerID

`func (o *AutomationTokenCreateResponseAllOf) HasCustomerID() bool`

HasCustomerID returns a boolean if a field has been set.

### GetSudoExpiresAt

`func (o *AutomationTokenCreateResponseAllOf) GetSudoExpiresAt() time.Time`

GetSudoExpiresAt returns the SudoExpiresAt field if non-nil, zero value otherwise.

### GetSudoExpiresAtOk

`func (o *AutomationTokenCreateResponseAllOf) GetSudoExpiresAtOk() (*time.Time, bool)`

GetSudoExpiresAtOk returns a tuple with the SudoExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSudoExpiresAt

`func (o *AutomationTokenCreateResponseAllOf) SetSudoExpiresAt(v time.Time)`

SetSudoExpiresAt sets SudoExpiresAt field to given value.

### HasSudoExpiresAt

`func (o *AutomationTokenCreateResponseAllOf) HasSudoExpiresAt() bool`

HasSudoExpiresAt returns a boolean if a field has been set.

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
