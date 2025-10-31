# TlsPrivateKeyResponseAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**DeletedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**UpdatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**Name** | Pointer to **string** | A customizable name for your private key. | [optional] 
**KeyLength** | Pointer to **int32** | The key length used to generate the private key. | [optional] [readonly] 
**KeyType** | Pointer to **string** | The algorithm used to generate the private key. Must be `RSA`. | [optional] [readonly] 
**Replace** | Pointer to **bool** | A recommendation from Fastly to replace this private key and all associated certificates. | [optional] [readonly] 
**PublicKeySha1** | Pointer to **string** | Useful for safely identifying the key. | [optional] [readonly] 

## Methods

### NewTlsPrivateKeyResponseAttributes

`func NewTlsPrivateKeyResponseAttributes() *TlsPrivateKeyResponseAttributes`

NewTlsPrivateKeyResponseAttributes instantiates a new TlsPrivateKeyResponseAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTlsPrivateKeyResponseAttributesWithDefaults

`func NewTlsPrivateKeyResponseAttributesWithDefaults() *TlsPrivateKeyResponseAttributes`

NewTlsPrivateKeyResponseAttributesWithDefaults instantiates a new TlsPrivateKeyResponseAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *TlsPrivateKeyResponseAttributes) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TlsPrivateKeyResponseAttributes) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TlsPrivateKeyResponseAttributes) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TlsPrivateKeyResponseAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *TlsPrivateKeyResponseAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *TlsPrivateKeyResponseAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetDeletedAt

`func (o *TlsPrivateKeyResponseAttributes) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *TlsPrivateKeyResponseAttributes) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *TlsPrivateKeyResponseAttributes) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *TlsPrivateKeyResponseAttributes) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *TlsPrivateKeyResponseAttributes) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *TlsPrivateKeyResponseAttributes) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetUpdatedAt

`func (o *TlsPrivateKeyResponseAttributes) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TlsPrivateKeyResponseAttributes) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TlsPrivateKeyResponseAttributes) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TlsPrivateKeyResponseAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *TlsPrivateKeyResponseAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *TlsPrivateKeyResponseAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetName

`func (o *TlsPrivateKeyResponseAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TlsPrivateKeyResponseAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TlsPrivateKeyResponseAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TlsPrivateKeyResponseAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetKeyLength

`func (o *TlsPrivateKeyResponseAttributes) GetKeyLength() int32`

GetKeyLength returns the KeyLength field if non-nil, zero value otherwise.

### GetKeyLengthOk

`func (o *TlsPrivateKeyResponseAttributes) GetKeyLengthOk() (*int32, bool)`

GetKeyLengthOk returns a tuple with the KeyLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyLength

`func (o *TlsPrivateKeyResponseAttributes) SetKeyLength(v int32)`

SetKeyLength sets KeyLength field to given value.

### HasKeyLength

`func (o *TlsPrivateKeyResponseAttributes) HasKeyLength() bool`

HasKeyLength returns a boolean if a field has been set.

### GetKeyType

`func (o *TlsPrivateKeyResponseAttributes) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *TlsPrivateKeyResponseAttributes) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *TlsPrivateKeyResponseAttributes) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.

### HasKeyType

`func (o *TlsPrivateKeyResponseAttributes) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.

### GetReplace

`func (o *TlsPrivateKeyResponseAttributes) GetReplace() bool`

GetReplace returns the Replace field if non-nil, zero value otherwise.

### GetReplaceOk

`func (o *TlsPrivateKeyResponseAttributes) GetReplaceOk() (*bool, bool)`

GetReplaceOk returns a tuple with the Replace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplace

`func (o *TlsPrivateKeyResponseAttributes) SetReplace(v bool)`

SetReplace sets Replace field to given value.

### HasReplace

`func (o *TlsPrivateKeyResponseAttributes) HasReplace() bool`

HasReplace returns a boolean if a field has been set.

### GetPublicKeySha1

`func (o *TlsPrivateKeyResponseAttributes) GetPublicKeySha1() string`

GetPublicKeySha1 returns the PublicKeySha1 field if non-nil, zero value otherwise.

### GetPublicKeySha1Ok

`func (o *TlsPrivateKeyResponseAttributes) GetPublicKeySha1Ok() (*string, bool)`

GetPublicKeySha1Ok returns a tuple with the PublicKeySha1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKeySha1

`func (o *TlsPrivateKeyResponseAttributes) SetPublicKeySha1(v string)`

SetPublicKeySha1 sets PublicKeySha1 field to given value.

### HasPublicKeySha1

`func (o *TlsPrivateKeyResponseAttributes) HasPublicKeySha1() bool`

HasPublicKeySha1 returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


