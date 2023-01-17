# TLSBulkCertificateResponseDataAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ID** | Pointer to **string** |  | [optional] [readonly] 
**Attributes** | Pointer to [**TLSBulkCertificateResponseAttributes**](TlsBulkCertificateResponseAttributes.md) |  | [optional] 

## Methods

### NewTLSBulkCertificateResponseDataAllOf

`func NewTLSBulkCertificateResponseDataAllOf() *TLSBulkCertificateResponseDataAllOf`

NewTLSBulkCertificateResponseDataAllOf instantiates a new TLSBulkCertificateResponseDataAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTLSBulkCertificateResponseDataAllOfWithDefaults

`func NewTLSBulkCertificateResponseDataAllOfWithDefaults() *TLSBulkCertificateResponseDataAllOf`

NewTLSBulkCertificateResponseDataAllOfWithDefaults instantiates a new TLSBulkCertificateResponseDataAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetID

`func (o *TLSBulkCertificateResponseDataAllOf) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *TLSBulkCertificateResponseDataAllOf) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *TLSBulkCertificateResponseDataAllOf) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *TLSBulkCertificateResponseDataAllOf) HasID() bool`

HasID returns a boolean if a field has been set.

### GetAttributes

`func (o *TLSBulkCertificateResponseDataAllOf) GetAttributes() TLSBulkCertificateResponseAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TLSBulkCertificateResponseDataAllOf) GetAttributesOk() (*TLSBulkCertificateResponseAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TLSBulkCertificateResponseDataAllOf) SetAttributes(v TLSBulkCertificateResponseAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *TLSBulkCertificateResponseDataAllOf) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
