# TLSCertificateResponseDataAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ID** | Pointer to **string** |  | [optional] [readonly] 
**Attributes** | Pointer to [**TLSCertificateResponseAttributes**](TlsCertificateResponseAttributes.md) |  | [optional] 

## Methods

### NewTLSCertificateResponseDataAllOf

`func NewTLSCertificateResponseDataAllOf() *TLSCertificateResponseDataAllOf`

NewTLSCertificateResponseDataAllOf instantiates a new TLSCertificateResponseDataAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTLSCertificateResponseDataAllOfWithDefaults

`func NewTLSCertificateResponseDataAllOfWithDefaults() *TLSCertificateResponseDataAllOf`

NewTLSCertificateResponseDataAllOfWithDefaults instantiates a new TLSCertificateResponseDataAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetID

`func (o *TLSCertificateResponseDataAllOf) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *TLSCertificateResponseDataAllOf) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *TLSCertificateResponseDataAllOf) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *TLSCertificateResponseDataAllOf) HasID() bool`

HasID returns a boolean if a field has been set.

### GetAttributes

`func (o *TLSCertificateResponseDataAllOf) GetAttributes() TLSCertificateResponseAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TLSCertificateResponseDataAllOf) GetAttributesOk() (*TLSCertificateResponseAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TLSCertificateResponseDataAllOf) SetAttributes(v TLSCertificateResponseAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *TLSCertificateResponseDataAllOf) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
