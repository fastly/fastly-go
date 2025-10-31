# TlsPrivateKeyResponseAttributesAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A customizable name for your private key. | [optional] 
**KeyLength** | Pointer to **int32** | The key length used to generate the private key. | [optional] [readonly] 
**KeyType** | Pointer to **string** | The algorithm used to generate the private key. Must be `RSA`. | [optional] [readonly] 
**Replace** | Pointer to **bool** | A recommendation from Fastly to replace this private key and all associated certificates. | [optional] [readonly] 
**PublicKeySha1** | Pointer to **string** | Useful for safely identifying the key. | [optional] [readonly] 

## Methods

### NewTlsPrivateKeyResponseAttributesAllOf

`func NewTlsPrivateKeyResponseAttributesAllOf() *TlsPrivateKeyResponseAttributesAllOf`

NewTlsPrivateKeyResponseAttributesAllOf instantiates a new TlsPrivateKeyResponseAttributesAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTlsPrivateKeyResponseAttributesAllOfWithDefaults

`func NewTlsPrivateKeyResponseAttributesAllOfWithDefaults() *TlsPrivateKeyResponseAttributesAllOf`

NewTlsPrivateKeyResponseAttributesAllOfWithDefaults instantiates a new TlsPrivateKeyResponseAttributesAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TlsPrivateKeyResponseAttributesAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TlsPrivateKeyResponseAttributesAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TlsPrivateKeyResponseAttributesAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TlsPrivateKeyResponseAttributesAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetKeyLength

`func (o *TlsPrivateKeyResponseAttributesAllOf) GetKeyLength() int32`

GetKeyLength returns the KeyLength field if non-nil, zero value otherwise.

### GetKeyLengthOk

`func (o *TlsPrivateKeyResponseAttributesAllOf) GetKeyLengthOk() (*int32, bool)`

GetKeyLengthOk returns a tuple with the KeyLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyLength

`func (o *TlsPrivateKeyResponseAttributesAllOf) SetKeyLength(v int32)`

SetKeyLength sets KeyLength field to given value.

### HasKeyLength

`func (o *TlsPrivateKeyResponseAttributesAllOf) HasKeyLength() bool`

HasKeyLength returns a boolean if a field has been set.

### GetKeyType

`func (o *TlsPrivateKeyResponseAttributesAllOf) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *TlsPrivateKeyResponseAttributesAllOf) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *TlsPrivateKeyResponseAttributesAllOf) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.

### HasKeyType

`func (o *TlsPrivateKeyResponseAttributesAllOf) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.

### GetReplace

`func (o *TlsPrivateKeyResponseAttributesAllOf) GetReplace() bool`

GetReplace returns the Replace field if non-nil, zero value otherwise.

### GetReplaceOk

`func (o *TlsPrivateKeyResponseAttributesAllOf) GetReplaceOk() (*bool, bool)`

GetReplaceOk returns a tuple with the Replace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplace

`func (o *TlsPrivateKeyResponseAttributesAllOf) SetReplace(v bool)`

SetReplace sets Replace field to given value.

### HasReplace

`func (o *TlsPrivateKeyResponseAttributesAllOf) HasReplace() bool`

HasReplace returns a boolean if a field has been set.

### GetPublicKeySha1

`func (o *TlsPrivateKeyResponseAttributesAllOf) GetPublicKeySha1() string`

GetPublicKeySha1 returns the PublicKeySha1 field if non-nil, zero value otherwise.

### GetPublicKeySha1Ok

`func (o *TlsPrivateKeyResponseAttributesAllOf) GetPublicKeySha1Ok() (*string, bool)`

GetPublicKeySha1Ok returns a tuple with the PublicKeySha1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKeySha1

`func (o *TlsPrivateKeyResponseAttributesAllOf) SetPublicKeySha1(v string)`

SetPublicKeySha1 sets PublicKeySha1 field to given value.

### HasPublicKeySha1

`func (o *TlsPrivateKeyResponseAttributesAllOf) HasPublicKeySha1() bool`

HasPublicKeySha1 returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


