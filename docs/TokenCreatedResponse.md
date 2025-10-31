# TokenCreatedResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Services** | Pointer to **[]string** | List of alphanumeric strings identifying services (optional). If no services are specified, the token will have access to all services on the account.  | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the token. | [optional] 
**Scope** | Pointer to **string** | Space-delimited list of authorization scope. | [optional] [default to "global"]
**CreatedAt** | Pointer to **string** | Time-stamp (UTC) of when the token was created. | [optional] 
**DeletedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**UpdatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**UserId** | Pointer to **string** |  | [optional] [readonly] 
**LastUsedAt** | Pointer to **string** | Time-stamp (UTC) of when the token was last used. | [optional] [readonly] 
**ExpiresAt** | Pointer to **string** | Time-stamp (UTC) of when the token will expire (optional). | [optional] 
**Ip** | Pointer to **string** | IP Address of the client that last used the token. | [optional] 
**UserAgent** | Pointer to **string** | User-Agent header of the client that last used the token. | [optional] 
**AccessToken** | Pointer to **string** | The alphanumeric string for accessing the API (only available on token creation). | [optional] 

## Methods

### NewTokenCreatedResponse

`func NewTokenCreatedResponse() *TokenCreatedResponse`

NewTokenCreatedResponse instantiates a new TokenCreatedResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenCreatedResponseWithDefaults

`func NewTokenCreatedResponseWithDefaults() *TokenCreatedResponse`

NewTokenCreatedResponseWithDefaults instantiates a new TokenCreatedResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServices

`func (o *TokenCreatedResponse) GetServices() []string`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *TokenCreatedResponse) GetServicesOk() (*[]string, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *TokenCreatedResponse) SetServices(v []string)`

SetServices sets Services field to given value.

### HasServices

`func (o *TokenCreatedResponse) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetName

`func (o *TokenCreatedResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TokenCreatedResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TokenCreatedResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TokenCreatedResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetScope

`func (o *TokenCreatedResponse) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *TokenCreatedResponse) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *TokenCreatedResponse) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *TokenCreatedResponse) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TokenCreatedResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TokenCreatedResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TokenCreatedResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TokenCreatedResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *TokenCreatedResponse) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *TokenCreatedResponse) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *TokenCreatedResponse) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *TokenCreatedResponse) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *TokenCreatedResponse) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *TokenCreatedResponse) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetUpdatedAt

`func (o *TokenCreatedResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TokenCreatedResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TokenCreatedResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TokenCreatedResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *TokenCreatedResponse) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *TokenCreatedResponse) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetId

`func (o *TokenCreatedResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TokenCreatedResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TokenCreatedResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TokenCreatedResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUserId

`func (o *TokenCreatedResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *TokenCreatedResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *TokenCreatedResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *TokenCreatedResponse) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetLastUsedAt

`func (o *TokenCreatedResponse) GetLastUsedAt() string`

GetLastUsedAt returns the LastUsedAt field if non-nil, zero value otherwise.

### GetLastUsedAtOk

`func (o *TokenCreatedResponse) GetLastUsedAtOk() (*string, bool)`

GetLastUsedAtOk returns a tuple with the LastUsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedAt

`func (o *TokenCreatedResponse) SetLastUsedAt(v string)`

SetLastUsedAt sets LastUsedAt field to given value.

### HasLastUsedAt

`func (o *TokenCreatedResponse) HasLastUsedAt() bool`

HasLastUsedAt returns a boolean if a field has been set.

### GetExpiresAt

`func (o *TokenCreatedResponse) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *TokenCreatedResponse) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *TokenCreatedResponse) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *TokenCreatedResponse) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetIp

`func (o *TokenCreatedResponse) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *TokenCreatedResponse) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *TokenCreatedResponse) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *TokenCreatedResponse) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetUserAgent

`func (o *TokenCreatedResponse) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *TokenCreatedResponse) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *TokenCreatedResponse) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *TokenCreatedResponse) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.

### GetAccessToken

`func (o *TokenCreatedResponse) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *TokenCreatedResponse) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *TokenCreatedResponse) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *TokenCreatedResponse) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


