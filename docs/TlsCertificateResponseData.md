# TLSCertificateResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**TypeTLSCertificate**](TypeTLSCertificate.md) |  | [optional] [default to TYPETLSCERTIFICATE_TLS_CERTIFICATE]
**Attributes** | Pointer to [**TLSCertificateResponseAttributes**](TlsCertificateResponseAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**RelationshipTLSDomains**](RelationshipTLSDomains.md) |  | [optional] 
**ID** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewTLSCertificateResponseData

`func NewTLSCertificateResponseData() *TLSCertificateResponseData`

NewTLSCertificateResponseData instantiates a new TLSCertificateResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTLSCertificateResponseDataWithDefaults

`func NewTLSCertificateResponseDataWithDefaults() *TLSCertificateResponseData`

NewTLSCertificateResponseDataWithDefaults instantiates a new TLSCertificateResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TLSCertificateResponseData) GetType() TypeTLSCertificate`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TLSCertificateResponseData) GetTypeOk() (*TypeTLSCertificate, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TLSCertificateResponseData) SetType(v TypeTLSCertificate)`

SetType sets Type field to given value.

### HasType

`func (o *TLSCertificateResponseData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAttributes

`func (o *TLSCertificateResponseData) GetAttributes() TLSCertificateResponseAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TLSCertificateResponseData) GetAttributesOk() (*TLSCertificateResponseAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TLSCertificateResponseData) SetAttributes(v TLSCertificateResponseAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *TLSCertificateResponseData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *TLSCertificateResponseData) GetRelationships() RelationshipTLSDomains`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *TLSCertificateResponseData) GetRelationshipsOk() (*RelationshipTLSDomains, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *TLSCertificateResponseData) SetRelationships(v RelationshipTLSDomains)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *TLSCertificateResponseData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetID

`func (o *TLSCertificateResponseData) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *TLSCertificateResponseData) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *TLSCertificateResponseData) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *TLSCertificateResponseData) HasID() bool`

HasID returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
