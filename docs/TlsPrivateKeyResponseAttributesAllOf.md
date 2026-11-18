# TLSPrivateKeyResponseAttributesAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A customizable name for your private key. | [optional] 
**KeyLength** | Pointer to **int32** | The key length used to generate the private key. | [optional] [readonly] 
**KeyType** | Pointer to **string** | The algorithm used to generate the private key. Must be `RSA`. | [optional] [readonly] 
**Replace** | Pointer to **bool** | A recommendation from Fastly to replace this private key and all associated certificates. | [optional] [readonly] 
**PublicKeySha1** | Pointer to **string** | Useful for safely identifying the key. | [optional] [readonly] 

## Methods

### NewTLSPrivateKeyResponseAttributesAllOf

`func NewTLSPrivateKeyResponseAttributesAllOf() *TLSPrivateKeyResponseAttributesAllOf`

NewTLSPrivateKeyResponseAttributesAllOf instantiates a new TLSPrivateKeyResponseAttributesAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTLSPrivateKeyResponseAttributesAllOfWithDefaults

`func NewTLSPrivateKeyResponseAttributesAllOfWithDefaults() *TLSPrivateKeyResponseAttributesAllOf`

NewTLSPrivateKeyResponseAttributesAllOfWithDefaults instantiates a new TLSPrivateKeyResponseAttributesAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TLSPrivateKeyResponseAttributesAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TLSPrivateKeyResponseAttributesAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TLSPrivateKeyResponseAttributesAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TLSPrivateKeyResponseAttributesAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetKeyLength

`func (o *TLSPrivateKeyResponseAttributesAllOf) GetKeyLength() int32`

GetKeyLength returns the KeyLength field if non-nil, zero value otherwise.

### GetKeyLengthOk

`func (o *TLSPrivateKeyResponseAttributesAllOf) GetKeyLengthOk() (*int32, bool)`

GetKeyLengthOk returns a tuple with the KeyLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyLength

`func (o *TLSPrivateKeyResponseAttributesAllOf) SetKeyLength(v int32)`

SetKeyLength sets KeyLength field to given value.

### HasKeyLength

`func (o *TLSPrivateKeyResponseAttributesAllOf) HasKeyLength() bool`

HasKeyLength returns a boolean if a field has been set.

### GetKeyType

`func (o *TLSPrivateKeyResponseAttributesAllOf) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *TLSPrivateKeyResponseAttributesAllOf) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *TLSPrivateKeyResponseAttributesAllOf) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.

### HasKeyType

`func (o *TLSPrivateKeyResponseAttributesAllOf) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.

### GetReplace

`func (o *TLSPrivateKeyResponseAttributesAllOf) GetReplace() bool`

GetReplace returns the Replace field if non-nil, zero value otherwise.

### GetReplaceOk

`func (o *TLSPrivateKeyResponseAttributesAllOf) GetReplaceOk() (*bool, bool)`

GetReplaceOk returns a tuple with the Replace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplace

`func (o *TLSPrivateKeyResponseAttributesAllOf) SetReplace(v bool)`

SetReplace sets Replace field to given value.

### HasReplace

`func (o *TLSPrivateKeyResponseAttributesAllOf) HasReplace() bool`

HasReplace returns a boolean if a field has been set.

### GetPublicKeySha1

`func (o *TLSPrivateKeyResponseAttributesAllOf) GetPublicKeySha1() string`

GetPublicKeySha1 returns the PublicKeySha1 field if non-nil, zero value otherwise.

### GetPublicKeySha1Ok

`func (o *TLSPrivateKeyResponseAttributesAllOf) GetPublicKeySha1Ok() (*string, bool)`

GetPublicKeySha1Ok returns a tuple with the PublicKeySha1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKeySha1

`func (o *TLSPrivateKeyResponseAttributesAllOf) SetPublicKeySha1(v string)`

SetPublicKeySha1 sets PublicKeySha1 field to given value.

### HasPublicKeySha1

`func (o *TLSPrivateKeyResponseAttributesAllOf) HasPublicKeySha1() bool`

HasPublicKeySha1 returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
