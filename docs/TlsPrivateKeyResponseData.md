# TlsPrivateKeyResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**TypeTlsPrivateKey**](TypeTlsPrivateKey.md) |  | [optional] [default to TYPETLSPRIVATEKEY_TLS_PRIVATE_KEY]
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Attributes** | Pointer to [**TlsPrivateKeyResponseAttributes**](TlsPrivateKeyResponseAttributes.md) |  | [optional] 

## Methods

### NewTlsPrivateKeyResponseData

`func NewTlsPrivateKeyResponseData() *TlsPrivateKeyResponseData`

NewTlsPrivateKeyResponseData instantiates a new TlsPrivateKeyResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTlsPrivateKeyResponseDataWithDefaults

`func NewTlsPrivateKeyResponseDataWithDefaults() *TlsPrivateKeyResponseData`

NewTlsPrivateKeyResponseDataWithDefaults instantiates a new TlsPrivateKeyResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TlsPrivateKeyResponseData) GetType() TypeTlsPrivateKey`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TlsPrivateKeyResponseData) GetTypeOk() (*TypeTlsPrivateKey, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TlsPrivateKeyResponseData) SetType(v TypeTlsPrivateKey)`

SetType sets Type field to given value.

### HasType

`func (o *TlsPrivateKeyResponseData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *TlsPrivateKeyResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TlsPrivateKeyResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TlsPrivateKeyResponseData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TlsPrivateKeyResponseData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAttributes

`func (o *TlsPrivateKeyResponseData) GetAttributes() TlsPrivateKeyResponseAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TlsPrivateKeyResponseData) GetAttributesOk() (*TlsPrivateKeyResponseAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TlsPrivateKeyResponseData) SetAttributes(v TlsPrivateKeyResponseAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *TlsPrivateKeyResponseData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


