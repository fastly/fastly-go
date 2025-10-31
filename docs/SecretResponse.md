# SecretResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the secret. | [optional] 
**Digest** | Pointer to **string** | An opaque identifier of the plaintext secret value. This can be used to determine if a secret value has changed. | [optional] [readonly] 
**CreatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**Recreated** | Pointer to **NullableBool** | True if the secret replaced a secret with the same name. | [optional] 

## Methods

### NewSecretResponse

`func NewSecretResponse() *SecretResponse`

NewSecretResponse instantiates a new SecretResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretResponseWithDefaults

`func NewSecretResponseWithDefaults() *SecretResponse`

NewSecretResponseWithDefaults instantiates a new SecretResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SecretResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecretResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecretResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SecretResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDigest

`func (o *SecretResponse) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *SecretResponse) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *SecretResponse) SetDigest(v string)`

SetDigest sets Digest field to given value.

### HasDigest

`func (o *SecretResponse) HasDigest() bool`

HasDigest returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SecretResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SecretResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SecretResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SecretResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *SecretResponse) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *SecretResponse) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetRecreated

`func (o *SecretResponse) GetRecreated() bool`

GetRecreated returns the Recreated field if non-nil, zero value otherwise.

### GetRecreatedOk

`func (o *SecretResponse) GetRecreatedOk() (*bool, bool)`

GetRecreatedOk returns a tuple with the Recreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecreated

`func (o *SecretResponse) SetRecreated(v bool)`

SetRecreated sets Recreated field to given value.

### HasRecreated

`func (o *SecretResponse) HasRecreated() bool`

HasRecreated returns a boolean if a field has been set.

### SetRecreatedNil

`func (o *SecretResponse) SetRecreatedNil(b bool)`

 SetRecreatedNil sets the value for Recreated to be an explicit nil

### UnsetRecreated
`func (o *SecretResponse) UnsetRecreated()`

UnsetRecreated ensures that no value is present for Recreated, not even an explicit nil

[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


