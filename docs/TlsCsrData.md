# TLSCsrData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**TypeTLSCsr**](TypeTLSCsr.md) |  | [optional] [default to TYPETLSCSR_CSR]
**Attributes** | Pointer to [**TLSCsrDataAttributes**](TlsCsrDataAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**RelationshipTLSPrivateKey**](RelationshipTLSPrivateKey.md) |  | [optional] 

## Methods

### NewTLSCsrData

`func NewTLSCsrData() *TLSCsrData`

NewTLSCsrData instantiates a new TLSCsrData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTLSCsrDataWithDefaults

`func NewTLSCsrDataWithDefaults() *TLSCsrData`

NewTLSCsrDataWithDefaults instantiates a new TLSCsrData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TLSCsrData) GetType() TypeTLSCsr`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TLSCsrData) GetTypeOk() (*TypeTLSCsr, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TLSCsrData) SetType(v TypeTLSCsr)`

SetType sets Type field to given value.

### HasType

`func (o *TLSCsrData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAttributes

`func (o *TLSCsrData) GetAttributes() TLSCsrDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TLSCsrData) GetAttributesOk() (*TLSCsrDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TLSCsrData) SetAttributes(v TLSCsrDataAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *TLSCsrData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *TLSCsrData) GetRelationships() RelationshipTLSPrivateKey`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *TLSCsrData) GetRelationshipsOk() (*RelationshipTLSPrivateKey, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *TLSCsrData) SetRelationships(v RelationshipTLSPrivateKey)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *TLSCsrData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
