# TLSBulkCertificateResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**TypeTLSBulkCertificate**](TypeTLSBulkCertificate.md) |  | [optional] [default to TYPETLSBULKCERTIFICATE_TLS_BULK_CERTIFICATE]
**Attributes** | Pointer to [**TLSBulkCertificateResponseAttributes**](TlsBulkCertificateResponseAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**RelationshipsForTLSBulkCertificate**](RelationshipsForTLSBulkCertificate.md) |  | [optional] 
**ID** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewTLSBulkCertificateResponseData

`func NewTLSBulkCertificateResponseData() *TLSBulkCertificateResponseData`

NewTLSBulkCertificateResponseData instantiates a new TLSBulkCertificateResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTLSBulkCertificateResponseDataWithDefaults

`func NewTLSBulkCertificateResponseDataWithDefaults() *TLSBulkCertificateResponseData`

NewTLSBulkCertificateResponseDataWithDefaults instantiates a new TLSBulkCertificateResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TLSBulkCertificateResponseData) GetType() TypeTLSBulkCertificate`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TLSBulkCertificateResponseData) GetTypeOk() (*TypeTLSBulkCertificate, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TLSBulkCertificateResponseData) SetType(v TypeTLSBulkCertificate)`

SetType sets Type field to given value.

### HasType

`func (o *TLSBulkCertificateResponseData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAttributes

`func (o *TLSBulkCertificateResponseData) GetAttributes() TLSBulkCertificateResponseAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TLSBulkCertificateResponseData) GetAttributesOk() (*TLSBulkCertificateResponseAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TLSBulkCertificateResponseData) SetAttributes(v TLSBulkCertificateResponseAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *TLSBulkCertificateResponseData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *TLSBulkCertificateResponseData) GetRelationships() RelationshipsForTLSBulkCertificate`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *TLSBulkCertificateResponseData) GetRelationshipsOk() (*RelationshipsForTLSBulkCertificate, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *TLSBulkCertificateResponseData) SetRelationships(v RelationshipsForTLSBulkCertificate)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *TLSBulkCertificateResponseData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetID

`func (o *TLSBulkCertificateResponseData) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *TLSBulkCertificateResponseData) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *TLSBulkCertificateResponseData) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *TLSBulkCertificateResponseData) HasID() bool`

HasID returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
