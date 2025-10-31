# TlsCertificateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**TypeTlsCertificate**](TypeTlsCertificate.md) |  | [optional] [default to TYPETLSCERTIFICATE_TLS_CERTIFICATE]
**Attributes** | Pointer to [**TlsCertificateDataAttributes**](TlsCertificateDataAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**RelationshipTlsDomains**](RelationshipTlsDomains.md) |  | [optional] 

## Methods

### NewTlsCertificateData

`func NewTlsCertificateData() *TlsCertificateData`

NewTlsCertificateData instantiates a new TlsCertificateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTlsCertificateDataWithDefaults

`func NewTlsCertificateDataWithDefaults() *TlsCertificateData`

NewTlsCertificateDataWithDefaults instantiates a new TlsCertificateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TlsCertificateData) GetType() TypeTlsCertificate`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TlsCertificateData) GetTypeOk() (*TypeTlsCertificate, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TlsCertificateData) SetType(v TypeTlsCertificate)`

SetType sets Type field to given value.

### HasType

`func (o *TlsCertificateData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAttributes

`func (o *TlsCertificateData) GetAttributes() TlsCertificateDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TlsCertificateData) GetAttributesOk() (*TlsCertificateDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TlsCertificateData) SetAttributes(v TlsCertificateDataAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *TlsCertificateData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *TlsCertificateData) GetRelationships() RelationshipTlsDomains`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *TlsCertificateData) GetRelationshipsOk() (*RelationshipTlsDomains, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *TlsCertificateData) SetRelationships(v RelationshipTlsDomains)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *TlsCertificateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


