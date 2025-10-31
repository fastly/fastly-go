# TlsCertificateResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**TypeTlsCertificate**](TypeTlsCertificate.md) |  | [optional] [default to TYPETLSCERTIFICATE_TLS_CERTIFICATE]
**Attributes** | Pointer to [**TlsCertificateResponseAttributes**](TlsCertificateResponseAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**RelationshipTlsDomains**](RelationshipTlsDomains.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewTlsCertificateResponseData

`func NewTlsCertificateResponseData() *TlsCertificateResponseData`

NewTlsCertificateResponseData instantiates a new TlsCertificateResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTlsCertificateResponseDataWithDefaults

`func NewTlsCertificateResponseDataWithDefaults() *TlsCertificateResponseData`

NewTlsCertificateResponseDataWithDefaults instantiates a new TlsCertificateResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TlsCertificateResponseData) GetType() TypeTlsCertificate`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TlsCertificateResponseData) GetTypeOk() (*TypeTlsCertificate, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TlsCertificateResponseData) SetType(v TypeTlsCertificate)`

SetType sets Type field to given value.

### HasType

`func (o *TlsCertificateResponseData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAttributes

`func (o *TlsCertificateResponseData) GetAttributes() TlsCertificateResponseAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TlsCertificateResponseData) GetAttributesOk() (*TlsCertificateResponseAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TlsCertificateResponseData) SetAttributes(v TlsCertificateResponseAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *TlsCertificateResponseData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *TlsCertificateResponseData) GetRelationships() RelationshipTlsDomains`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *TlsCertificateResponseData) GetRelationshipsOk() (*RelationshipTlsDomains, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *TlsCertificateResponseData) SetRelationships(v RelationshipTlsDomains)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *TlsCertificateResponseData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetId

`func (o *TlsCertificateResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TlsCertificateResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TlsCertificateResponseData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TlsCertificateResponseData) HasId() bool`

HasId returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


