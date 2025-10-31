# TlsBulkCertificateResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**TypeTlsBulkCertificate**](TypeTlsBulkCertificate.md) |  | [optional] [default to TYPETLSBULKCERTIFICATE_TLS_BULK_CERTIFICATE]
**Attributes** | Pointer to [**TlsBulkCertificateResponseAttributes**](TlsBulkCertificateResponseAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**RelationshipsForTlsBulkCertificate**](RelationshipsForTlsBulkCertificate.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewTlsBulkCertificateResponseData

`func NewTlsBulkCertificateResponseData() *TlsBulkCertificateResponseData`

NewTlsBulkCertificateResponseData instantiates a new TlsBulkCertificateResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTlsBulkCertificateResponseDataWithDefaults

`func NewTlsBulkCertificateResponseDataWithDefaults() *TlsBulkCertificateResponseData`

NewTlsBulkCertificateResponseDataWithDefaults instantiates a new TlsBulkCertificateResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TlsBulkCertificateResponseData) GetType() TypeTlsBulkCertificate`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TlsBulkCertificateResponseData) GetTypeOk() (*TypeTlsBulkCertificate, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TlsBulkCertificateResponseData) SetType(v TypeTlsBulkCertificate)`

SetType sets Type field to given value.

### HasType

`func (o *TlsBulkCertificateResponseData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAttributes

`func (o *TlsBulkCertificateResponseData) GetAttributes() TlsBulkCertificateResponseAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TlsBulkCertificateResponseData) GetAttributesOk() (*TlsBulkCertificateResponseAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TlsBulkCertificateResponseData) SetAttributes(v TlsBulkCertificateResponseAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *TlsBulkCertificateResponseData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *TlsBulkCertificateResponseData) GetRelationships() RelationshipsForTlsBulkCertificate`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *TlsBulkCertificateResponseData) GetRelationshipsOk() (*RelationshipsForTlsBulkCertificate, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *TlsBulkCertificateResponseData) SetRelationships(v RelationshipsForTlsBulkCertificate)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *TlsBulkCertificateResponseData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetId

`func (o *TlsBulkCertificateResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TlsBulkCertificateResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TlsBulkCertificateResponseData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TlsBulkCertificateResponseData) HasId() bool`

HasId returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


