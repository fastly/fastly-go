# TLSPrivateKeyResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**TypeTLSPrivateKey**](TypeTLSPrivateKey.md) |  | [optional] [default to TYPETLSPRIVATEKEY_TLS_PRIVATE_KEY]
**ID** | Pointer to **string** |  | [optional] [readonly] 
**Attributes** | Pointer to [**TLSPrivateKeyResponseAttributes**](TlsPrivateKeyResponseAttributes.md) |  | [optional] 

## Methods

### NewTLSPrivateKeyResponseData

`func NewTLSPrivateKeyResponseData() *TLSPrivateKeyResponseData`

NewTLSPrivateKeyResponseData instantiates a new TLSPrivateKeyResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTLSPrivateKeyResponseDataWithDefaults

`func NewTLSPrivateKeyResponseDataWithDefaults() *TLSPrivateKeyResponseData`

NewTLSPrivateKeyResponseDataWithDefaults instantiates a new TLSPrivateKeyResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TLSPrivateKeyResponseData) GetType() TypeTLSPrivateKey`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TLSPrivateKeyResponseData) GetTypeOk() (*TypeTLSPrivateKey, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TLSPrivateKeyResponseData) SetType(v TypeTLSPrivateKey)`

SetType sets Type field to given value.

### HasType

`func (o *TLSPrivateKeyResponseData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetID

`func (o *TLSPrivateKeyResponseData) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *TLSPrivateKeyResponseData) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *TLSPrivateKeyResponseData) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *TLSPrivateKeyResponseData) HasID() bool`

HasID returns a boolean if a field has been set.

### GetAttributes

`func (o *TLSPrivateKeyResponseData) GetAttributes() TLSPrivateKeyResponseAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TLSPrivateKeyResponseData) GetAttributesOk() (*TLSPrivateKeyResponseAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TLSPrivateKeyResponseData) SetAttributes(v TLSPrivateKeyResponseAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *TLSPrivateKeyResponseData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
