# ArcCreateVirtualKeyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier of the virtual key. | [readonly] 
**CustomerId** | **string** | The ID of the customer account that owns the key. | [readonly] 
**Name** | **string** | The display name of the virtual key. | 
**UserId** | **string** | The ID of the user who created the key. | [readonly] 
**ExpiresAt** | Pointer to **NullableTime** | The date and time the key expires, in RFC 3339 format. `null` if the key doesn&#39;t expire. | [optional] 
**RpmLimit** | Pointer to **NullableInt32** | The requests-per-minute limit for this key. `null` if there is no limit. | [optional] 
**TpmLimit** | Pointer to **NullableInt32** | The tokens-per-minute limit for this key. `null` if there is no limit. | [optional] 
**CreatedAt** | **time.Time** | The date and time the key was created, in RFC 3339 format. | [readonly] 
**UpdatedAt** | **time.Time** | The date and time the key was last updated, in RFC 3339 format. | [readonly] 
**AccessToken** | **string** | The access token for this virtual key. This is only returned when the key is created or rotated, and cannot be retrieved again afterward. | [readonly] 

## Methods

### NewArcCreateVirtualKeyResponse

`func NewArcCreateVirtualKeyResponse(id string, customerId string, name string, userId string, createdAt time.Time, updatedAt time.Time, accessToken string, ) *ArcCreateVirtualKeyResponse`

NewArcCreateVirtualKeyResponse instantiates a new ArcCreateVirtualKeyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArcCreateVirtualKeyResponseWithDefaults

`func NewArcCreateVirtualKeyResponseWithDefaults() *ArcCreateVirtualKeyResponse`

NewArcCreateVirtualKeyResponseWithDefaults instantiates a new ArcCreateVirtualKeyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ArcCreateVirtualKeyResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArcCreateVirtualKeyResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArcCreateVirtualKeyResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCustomerId

`func (o *ArcCreateVirtualKeyResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *ArcCreateVirtualKeyResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *ArcCreateVirtualKeyResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetName

`func (o *ArcCreateVirtualKeyResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ArcCreateVirtualKeyResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ArcCreateVirtualKeyResponse) SetName(v string)`

SetName sets Name field to given value.


### GetUserId

`func (o *ArcCreateVirtualKeyResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ArcCreateVirtualKeyResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ArcCreateVirtualKeyResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetExpiresAt

`func (o *ArcCreateVirtualKeyResponse) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ArcCreateVirtualKeyResponse) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ArcCreateVirtualKeyResponse) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *ArcCreateVirtualKeyResponse) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *ArcCreateVirtualKeyResponse) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *ArcCreateVirtualKeyResponse) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetRpmLimit

`func (o *ArcCreateVirtualKeyResponse) GetRpmLimit() int32`

GetRpmLimit returns the RpmLimit field if non-nil, zero value otherwise.

### GetRpmLimitOk

`func (o *ArcCreateVirtualKeyResponse) GetRpmLimitOk() (*int32, bool)`

GetRpmLimitOk returns a tuple with the RpmLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpmLimit

`func (o *ArcCreateVirtualKeyResponse) SetRpmLimit(v int32)`

SetRpmLimit sets RpmLimit field to given value.

### HasRpmLimit

`func (o *ArcCreateVirtualKeyResponse) HasRpmLimit() bool`

HasRpmLimit returns a boolean if a field has been set.

### SetRpmLimitNil

`func (o *ArcCreateVirtualKeyResponse) SetRpmLimitNil(b bool)`

 SetRpmLimitNil sets the value for RpmLimit to be an explicit nil

### UnsetRpmLimit
`func (o *ArcCreateVirtualKeyResponse) UnsetRpmLimit()`

UnsetRpmLimit ensures that no value is present for RpmLimit, not even an explicit nil
### GetTpmLimit

`func (o *ArcCreateVirtualKeyResponse) GetTpmLimit() int32`

GetTpmLimit returns the TpmLimit field if non-nil, zero value otherwise.

### GetTpmLimitOk

`func (o *ArcCreateVirtualKeyResponse) GetTpmLimitOk() (*int32, bool)`

GetTpmLimitOk returns a tuple with the TpmLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpmLimit

`func (o *ArcCreateVirtualKeyResponse) SetTpmLimit(v int32)`

SetTpmLimit sets TpmLimit field to given value.

### HasTpmLimit

`func (o *ArcCreateVirtualKeyResponse) HasTpmLimit() bool`

HasTpmLimit returns a boolean if a field has been set.

### SetTpmLimitNil

`func (o *ArcCreateVirtualKeyResponse) SetTpmLimitNil(b bool)`

 SetTpmLimitNil sets the value for TpmLimit to be an explicit nil

### UnsetTpmLimit
`func (o *ArcCreateVirtualKeyResponse) UnsetTpmLimit()`

UnsetTpmLimit ensures that no value is present for TpmLimit, not even an explicit nil
### GetCreatedAt

`func (o *ArcCreateVirtualKeyResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ArcCreateVirtualKeyResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ArcCreateVirtualKeyResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ArcCreateVirtualKeyResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ArcCreateVirtualKeyResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ArcCreateVirtualKeyResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetAccessToken

`func (o *ArcCreateVirtualKeyResponse) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *ArcCreateVirtualKeyResponse) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *ArcCreateVirtualKeyResponse) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.



[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


