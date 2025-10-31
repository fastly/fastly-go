# TlsPrivateKeyData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**TypeTlsPrivateKey**](TypeTlsPrivateKey.md) |  | [optional] [default to TYPETLSPRIVATEKEY_TLS_PRIVATE_KEY]
**Attributes** | Pointer to [**TlsPrivateKeyDataAttributes**](TlsPrivateKeyDataAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**RelationshipsForTlsPrivateKey**](RelationshipsForTlsPrivateKey.md) |  | [optional] 

## Methods

### NewTlsPrivateKeyData

`func NewTlsPrivateKeyData() *TlsPrivateKeyData`

NewTlsPrivateKeyData instantiates a new TlsPrivateKeyData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTlsPrivateKeyDataWithDefaults

`func NewTlsPrivateKeyDataWithDefaults() *TlsPrivateKeyData`

NewTlsPrivateKeyDataWithDefaults instantiates a new TlsPrivateKeyData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TlsPrivateKeyData) GetType() TypeTlsPrivateKey`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TlsPrivateKeyData) GetTypeOk() (*TypeTlsPrivateKey, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TlsPrivateKeyData) SetType(v TypeTlsPrivateKey)`

SetType sets Type field to given value.

### HasType

`func (o *TlsPrivateKeyData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAttributes

`func (o *TlsPrivateKeyData) GetAttributes() TlsPrivateKeyDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TlsPrivateKeyData) GetAttributesOk() (*TlsPrivateKeyDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TlsPrivateKeyData) SetAttributes(v TlsPrivateKeyDataAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *TlsPrivateKeyData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *TlsPrivateKeyData) GetRelationships() RelationshipsForTlsPrivateKey`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *TlsPrivateKeyData) GetRelationshipsOk() (*RelationshipsForTlsPrivateKey, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *TlsPrivateKeyData) SetRelationships(v RelationshipsForTlsPrivateKey)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *TlsPrivateKeyData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


