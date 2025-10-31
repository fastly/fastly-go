# TlsCsrResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**TypeTlsCsr**](TypeTlsCsr.md) |  | [optional] [default to TYPETLSCSR_CSR]
**Attributes** | Pointer to [**TlsCsrResponseAttributes**](TlsCsrResponseAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**RelationshipTlsPrivateKey**](RelationshipTlsPrivateKey.md) |  | [optional] 

## Methods

### NewTlsCsrResponseData

`func NewTlsCsrResponseData() *TlsCsrResponseData`

NewTlsCsrResponseData instantiates a new TlsCsrResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTlsCsrResponseDataWithDefaults

`func NewTlsCsrResponseDataWithDefaults() *TlsCsrResponseData`

NewTlsCsrResponseDataWithDefaults instantiates a new TlsCsrResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TlsCsrResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TlsCsrResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TlsCsrResponseData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TlsCsrResponseData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *TlsCsrResponseData) GetType() TypeTlsCsr`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TlsCsrResponseData) GetTypeOk() (*TypeTlsCsr, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TlsCsrResponseData) SetType(v TypeTlsCsr)`

SetType sets Type field to given value.

### HasType

`func (o *TlsCsrResponseData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAttributes

`func (o *TlsCsrResponseData) GetAttributes() TlsCsrResponseAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TlsCsrResponseData) GetAttributesOk() (*TlsCsrResponseAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TlsCsrResponseData) SetAttributes(v TlsCsrResponseAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *TlsCsrResponseData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *TlsCsrResponseData) GetRelationships() RelationshipTlsPrivateKey`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *TlsCsrResponseData) GetRelationshipsOk() (*RelationshipTlsPrivateKey, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *TlsCsrResponseData) SetRelationships(v RelationshipTlsPrivateKey)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *TlsCsrResponseData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


