# TLSPrivateKeyData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**TypeTLSPrivateKey**](TypeTLSPrivateKey.md) |  | [optional] [default to TYPETLSPRIVATEKEY_TLS_PRIVATE_KEY]
**Attributes** | Pointer to [**TLSPrivateKeyDataAttributes**](TlsPrivateKeyDataAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**RelationshipsForTLSPrivateKey**](RelationshipsForTLSPrivateKey.md) |  | [optional] 

## Methods

### NewTLSPrivateKeyData

`func NewTLSPrivateKeyData() *TLSPrivateKeyData`

NewTLSPrivateKeyData instantiates a new TLSPrivateKeyData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTLSPrivateKeyDataWithDefaults

`func NewTLSPrivateKeyDataWithDefaults() *TLSPrivateKeyData`

NewTLSPrivateKeyDataWithDefaults instantiates a new TLSPrivateKeyData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TLSPrivateKeyData) GetType() TypeTLSPrivateKey`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TLSPrivateKeyData) GetTypeOk() (*TypeTLSPrivateKey, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TLSPrivateKeyData) SetType(v TypeTLSPrivateKey)`

SetType sets Type field to given value.

### HasType

`func (o *TLSPrivateKeyData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAttributes

`func (o *TLSPrivateKeyData) GetAttributes() TLSPrivateKeyDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TLSPrivateKeyData) GetAttributesOk() (*TLSPrivateKeyDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TLSPrivateKeyData) SetAttributes(v TLSPrivateKeyDataAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *TLSPrivateKeyData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *TLSPrivateKeyData) GetRelationships() RelationshipsForTLSPrivateKey`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *TLSPrivateKeyData) GetRelationshipsOk() (*RelationshipsForTLSPrivateKey, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *TLSPrivateKeyData) SetRelationships(v RelationshipsForTLSPrivateKey)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *TLSPrivateKeyData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
