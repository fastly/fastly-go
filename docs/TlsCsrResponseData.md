# TLSCsrResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ID** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**TypeTLSCsr**](TypeTLSCsr.md) |  | [optional] [default to TYPETLSCSR_CSR]
**Attributes** | Pointer to [**TLSCsrResponseAttributes**](TlsCsrResponseAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**RelationshipsForTLSCsr**](RelationshipsForTLSCsr.md) |  | [optional] 

## Methods

### NewTLSCsrResponseData

`func NewTLSCsrResponseData() *TLSCsrResponseData`

NewTLSCsrResponseData instantiates a new TLSCsrResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTLSCsrResponseDataWithDefaults

`func NewTLSCsrResponseDataWithDefaults() *TLSCsrResponseData`

NewTLSCsrResponseDataWithDefaults instantiates a new TLSCsrResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetID

`func (o *TLSCsrResponseData) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *TLSCsrResponseData) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *TLSCsrResponseData) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *TLSCsrResponseData) HasID() bool`

HasID returns a boolean if a field has been set.

### GetType

`func (o *TLSCsrResponseData) GetType() TypeTLSCsr`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TLSCsrResponseData) GetTypeOk() (*TypeTLSCsr, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TLSCsrResponseData) SetType(v TypeTLSCsr)`

SetType sets Type field to given value.

### HasType

`func (o *TLSCsrResponseData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAttributes

`func (o *TLSCsrResponseData) GetAttributes() TLSCsrResponseAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TLSCsrResponseData) GetAttributesOk() (*TLSCsrResponseAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TLSCsrResponseData) SetAttributes(v TLSCsrResponseAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *TLSCsrResponseData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *TLSCsrResponseData) GetRelationships() RelationshipsForTLSCsr`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *TLSCsrResponseData) GetRelationshipsOk() (*RelationshipsForTLSCsr, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *TLSCsrResponseData) SetRelationships(v RelationshipsForTLSCsr)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *TLSCsrResponseData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
