# ArcRotateVirtualKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiresAt** | Pointer to **time.Time** | The date and time the new access token expires, in RFC 3339 format. | [optional] 

## Methods

### NewArcRotateVirtualKeyRequest

`func NewArcRotateVirtualKeyRequest() *ArcRotateVirtualKeyRequest`

NewArcRotateVirtualKeyRequest instantiates a new ArcRotateVirtualKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArcRotateVirtualKeyRequestWithDefaults

`func NewArcRotateVirtualKeyRequestWithDefaults() *ArcRotateVirtualKeyRequest`

NewArcRotateVirtualKeyRequestWithDefaults instantiates a new ArcRotateVirtualKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiresAt

`func (o *ArcRotateVirtualKeyRequest) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ArcRotateVirtualKeyRequest) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ArcRotateVirtualKeyRequest) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *ArcRotateVirtualKeyRequest) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


