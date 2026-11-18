# TokenResponseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ID** | Pointer to **string** |  | [optional] [readonly] 
**UserID** | Pointer to **string** |  | [optional] [readonly] 
**CreatedAt** | Pointer to **string** | Time-stamp (UTC) of when the token was created. | [optional] 
**LastUsedAt** | Pointer to **string** | Time-stamp (UTC) of when the token was last used. | [optional] [readonly] 
**ExpiresAt** | Pointer to **string** | Time-stamp (UTC) of when the token will expire (optional). | [optional] 
**IP** | Pointer to **string** | IP Address of the client that last used the token. | [optional] 
**UserAgent** | Pointer to **string** | User-Agent header of the client that last used the token. | [optional] 

## Methods

### NewTokenResponseAllOf

`func NewTokenResponseAllOf() *TokenResponseAllOf`

NewTokenResponseAllOf instantiates a new TokenResponseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenResponseAllOfWithDefaults

`func NewTokenResponseAllOfWithDefaults() *TokenResponseAllOf`

NewTokenResponseAllOfWithDefaults instantiates a new TokenResponseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetID

`func (o *TokenResponseAllOf) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *TokenResponseAllOf) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *TokenResponseAllOf) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *TokenResponseAllOf) HasID() bool`

HasID returns a boolean if a field has been set.

### GetUserID

`func (o *TokenResponseAllOf) GetUserID() string`

GetUserID returns the UserID field if non-nil, zero value otherwise.

### GetUserIDOk

`func (o *TokenResponseAllOf) GetUserIDOk() (*string, bool)`

GetUserIDOk returns a tuple with the UserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserID

`func (o *TokenResponseAllOf) SetUserID(v string)`

SetUserID sets UserID field to given value.

### HasUserID

`func (o *TokenResponseAllOf) HasUserID() bool`

HasUserID returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TokenResponseAllOf) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TokenResponseAllOf) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TokenResponseAllOf) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TokenResponseAllOf) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastUsedAt

`func (o *TokenResponseAllOf) GetLastUsedAt() string`

GetLastUsedAt returns the LastUsedAt field if non-nil, zero value otherwise.

### GetLastUsedAtOk

`func (o *TokenResponseAllOf) GetLastUsedAtOk() (*string, bool)`

GetLastUsedAtOk returns a tuple with the LastUsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedAt

`func (o *TokenResponseAllOf) SetLastUsedAt(v string)`

SetLastUsedAt sets LastUsedAt field to given value.

### HasLastUsedAt

`func (o *TokenResponseAllOf) HasLastUsedAt() bool`

HasLastUsedAt returns a boolean if a field has been set.

### GetExpiresAt

`func (o *TokenResponseAllOf) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *TokenResponseAllOf) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *TokenResponseAllOf) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *TokenResponseAllOf) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetIP

`func (o *TokenResponseAllOf) GetIP() string`

GetIP returns the IP field if non-nil, zero value otherwise.

### GetIPOk

`func (o *TokenResponseAllOf) GetIPOk() (*string, bool)`

GetIPOk returns a tuple with the IP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIP

`func (o *TokenResponseAllOf) SetIP(v string)`

SetIP sets IP field to given value.

### HasIP

`func (o *TokenResponseAllOf) HasIP() bool`

HasIP returns a boolean if a field has been set.

### GetUserAgent

`func (o *TokenResponseAllOf) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *TokenResponseAllOf) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *TokenResponseAllOf) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *TokenResponseAllOf) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
