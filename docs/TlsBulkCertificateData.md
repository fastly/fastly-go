# TlsBulkCertificateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**TypeTlsBulkCertificate**](TypeTlsBulkCertificate.md) |  | [optional] [default to TYPETLSBULKCERTIFICATE_TLS_BULK_CERTIFICATE]
**Attributes** | Pointer to [**TlsBulkCertificateDataAttributes**](TlsBulkCertificateDataAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**RelationshipsForTlsBulkCertificate**](RelationshipsForTlsBulkCertificate.md) |  | [optional] 

## Methods

### NewTlsBulkCertificateData

`func NewTlsBulkCertificateData() *TlsBulkCertificateData`

NewTlsBulkCertificateData instantiates a new TlsBulkCertificateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTlsBulkCertificateDataWithDefaults

`func NewTlsBulkCertificateDataWithDefaults() *TlsBulkCertificateData`

NewTlsBulkCertificateDataWithDefaults instantiates a new TlsBulkCertificateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TlsBulkCertificateData) GetType() TypeTlsBulkCertificate`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TlsBulkCertificateData) GetTypeOk() (*TypeTlsBulkCertificate, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TlsBulkCertificateData) SetType(v TypeTlsBulkCertificate)`

SetType sets Type field to given value.

### HasType

`func (o *TlsBulkCertificateData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAttributes

`func (o *TlsBulkCertificateData) GetAttributes() TlsBulkCertificateDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TlsBulkCertificateData) GetAttributesOk() (*TlsBulkCertificateDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TlsBulkCertificateData) SetAttributes(v TlsBulkCertificateDataAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *TlsBulkCertificateData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *TlsBulkCertificateData) GetRelationships() RelationshipsForTlsBulkCertificate`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *TlsBulkCertificateData) GetRelationshipsOk() (*RelationshipsForTlsBulkCertificate, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *TlsBulkCertificateData) SetRelationships(v RelationshipsForTlsBulkCertificate)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *TlsBulkCertificateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


