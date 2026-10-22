# ArcRotateVirtualKeyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier of the virtual key. | [readonly] 
**CustomerId** | **string** | The ID of the customer account that owns the key. | [readonly] 
**Name** | **string** | The display name of the virtual key. | 
**UserId** | **string** | The ID of the user who created the key. | [readonly] 
**ExpiresAt** | Pointer to **NullableTime** | The date and time the key expires, in RFC 3339 format. `null` if the key doesn&#39;t expire. | [optional] 
**CreatedAt** | **time.Time** | The date and time the key was created, in RFC 3339 format. | [readonly] 
**UpdatedAt** | **time.Time** | The date and time the key was last updated, in RFC 3339 format. | [readonly] 
**AccessToken** | **string** | The new access token for this virtual key. This is only returned when the key is created or rotated, and cannot be retrieved again afterward. | [readonly] 

## Methods

### NewArcRotateVirtualKeyResponse

`func NewArcRotateVirtualKeyResponse(id string, customerId string, name string, userId string, createdAt time.Time, updatedAt time.Time, accessToken string, ) *ArcRotateVirtualKeyResponse`

NewArcRotateVirtualKeyResponse instantiates a new ArcRotateVirtualKeyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArcRotateVirtualKeyResponseWithDefaults

`func NewArcRotateVirtualKeyResponseWithDefaults() *ArcRotateVirtualKeyResponse`

NewArcRotateVirtualKeyResponseWithDefaults instantiates a new ArcRotateVirtualKeyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ArcRotateVirtualKeyResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArcRotateVirtualKeyResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArcRotateVirtualKeyResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCustomerId

`func (o *ArcRotateVirtualKeyResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *ArcRotateVirtualKeyResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *ArcRotateVirtualKeyResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetName

`func (o *ArcRotateVirtualKeyResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ArcRotateVirtualKeyResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ArcRotateVirtualKeyResponse) SetName(v string)`

SetName sets Name field to given value.


### GetUserId

`func (o *ArcRotateVirtualKeyResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ArcRotateVirtualKeyResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ArcRotateVirtualKeyResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetExpiresAt

`func (o *ArcRotateVirtualKeyResponse) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ArcRotateVirtualKeyResponse) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ArcRotateVirtualKeyResponse) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *ArcRotateVirtualKeyResponse) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *ArcRotateVirtualKeyResponse) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *ArcRotateVirtualKeyResponse) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetCreatedAt

`func (o *ArcRotateVirtualKeyResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ArcRotateVirtualKeyResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ArcRotateVirtualKeyResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ArcRotateVirtualKeyResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ArcRotateVirtualKeyResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ArcRotateVirtualKeyResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetAccessToken

`func (o *ArcRotateVirtualKeyResponse) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *ArcRotateVirtualKeyResponse) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *ArcRotateVirtualKeyResponse) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.



[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


