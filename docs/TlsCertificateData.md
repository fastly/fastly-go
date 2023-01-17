# TLSCertificateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**TypeTLSCertificate**](TypeTLSCertificate.md) |  | [optional] [default to TYPETLSCERTIFICATE_TLS_CERTIFICATE]
**Attributes** | Pointer to [**TLSCertificateDataAttributes**](TlsCertificateDataAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**RelationshipTLSDomains**](RelationshipTLSDomains.md) |  | [optional] 

## Methods

### NewTLSCertificateData

`func NewTLSCertificateData() *TLSCertificateData`

NewTLSCertificateData instantiates a new TLSCertificateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTLSCertificateDataWithDefaults

`func NewTLSCertificateDataWithDefaults() *TLSCertificateData`

NewTLSCertificateDataWithDefaults instantiates a new TLSCertificateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TLSCertificateData) GetType() TypeTLSCertificate`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TLSCertificateData) GetTypeOk() (*TypeTLSCertificate, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TLSCertificateData) SetType(v TypeTLSCertificate)`

SetType sets Type field to given value.

### HasType

`func (o *TLSCertificateData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAttributes

`func (o *TLSCertificateData) GetAttributes() TLSCertificateDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TLSCertificateData) GetAttributesOk() (*TLSCertificateDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TLSCertificateData) SetAttributes(v TLSCertificateDataAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *TLSCertificateData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *TLSCertificateData) GetRelationships() RelationshipTLSDomains`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *TLSCertificateData) GetRelationshipsOk() (*RelationshipTLSDomains, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *TLSCertificateData) SetRelationships(v RelationshipTLSDomains)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *TLSCertificateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
