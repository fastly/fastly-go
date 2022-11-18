# TLSBulkCertificateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**TypeTLSBulkCertificate**](TypeTLSBulkCertificate.md) |  | [optional] [default to TYPETLSBULKCERTIFICATE_TLS_BULK_CERTIFICATE]
**Attributes** | Pointer to [**TLSBulkCertificateDataAttributes**](TlsBulkCertificateDataAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**RelationshipsForTLSBulkCertificate**](RelationshipsForTLSBulkCertificate.md) |  | [optional] 

## Methods

### NewTLSBulkCertificateData

`func NewTLSBulkCertificateData() *TLSBulkCertificateData`

NewTLSBulkCertificateData instantiates a new TLSBulkCertificateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTLSBulkCertificateDataWithDefaults

`func NewTLSBulkCertificateDataWithDefaults() *TLSBulkCertificateData`

NewTLSBulkCertificateDataWithDefaults instantiates a new TLSBulkCertificateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TLSBulkCertificateData) GetType() TypeTLSBulkCertificate`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TLSBulkCertificateData) GetTypeOk() (*TypeTLSBulkCertificate, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TLSBulkCertificateData) SetType(v TypeTLSBulkCertificate)`

SetType sets Type field to given value.

### HasType

`func (o *TLSBulkCertificateData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAttributes

`func (o *TLSBulkCertificateData) GetAttributes() TLSBulkCertificateDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TLSBulkCertificateData) GetAttributesOk() (*TLSBulkCertificateDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TLSBulkCertificateData) SetAttributes(v TLSBulkCertificateDataAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *TLSBulkCertificateData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *TLSBulkCertificateData) GetRelationships() RelationshipsForTLSBulkCertificate`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *TLSBulkCertificateData) GetRelationshipsOk() (*RelationshipsForTLSBulkCertificate, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *TLSBulkCertificateData) SetRelationships(v RelationshipsForTLSBulkCertificate)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *TLSBulkCertificateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
