# TLSCertificateResponseAttributesAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IssuedTo** | Pointer to **string** | The hostname for which a certificate was issued. | [optional] [readonly] 
**Issuer** | Pointer to **string** | The certificate authority that issued the certificate. | [optional] [readonly] 
**SerialNumber** | Pointer to **string** | A value assigned by the issuer that is unique to a certificate. | [optional] [readonly] 
**SignatureAlgorithm** | Pointer to **string** | The algorithm used to sign the certificate. | [optional] [readonly] 
**NotAfter** | Pointer to **time.Time** | Time-stamp (GMT) when the certificate will expire. Must be in the future to be used to terminate TLS traffic. | [optional] [readonly] 
**NotBefore** | Pointer to **time.Time** | Time-stamp (GMT) when the certificate will become valid. Must be in the past to be used to terminate TLS traffic. | [optional] [readonly] 
**Replace** | Pointer to **bool** | A recommendation from Fastly indicating the key associated with this certificate is in need of rotation. | [optional] [readonly] 

## Methods

### NewTLSCertificateResponseAttributesAllOf

`func NewTLSCertificateResponseAttributesAllOf() *TLSCertificateResponseAttributesAllOf`

NewTLSCertificateResponseAttributesAllOf instantiates a new TLSCertificateResponseAttributesAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTLSCertificateResponseAttributesAllOfWithDefaults

`func NewTLSCertificateResponseAttributesAllOfWithDefaults() *TLSCertificateResponseAttributesAllOf`

NewTLSCertificateResponseAttributesAllOfWithDefaults instantiates a new TLSCertificateResponseAttributesAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssuedTo

`func (o *TLSCertificateResponseAttributesAllOf) GetIssuedTo() string`

GetIssuedTo returns the IssuedTo field if non-nil, zero value otherwise.

### GetIssuedToOk

`func (o *TLSCertificateResponseAttributesAllOf) GetIssuedToOk() (*string, bool)`

GetIssuedToOk returns a tuple with the IssuedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedTo

`func (o *TLSCertificateResponseAttributesAllOf) SetIssuedTo(v string)`

SetIssuedTo sets IssuedTo field to given value.

### HasIssuedTo

`func (o *TLSCertificateResponseAttributesAllOf) HasIssuedTo() bool`

HasIssuedTo returns a boolean if a field has been set.

### GetIssuer

`func (o *TLSCertificateResponseAttributesAllOf) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *TLSCertificateResponseAttributesAllOf) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *TLSCertificateResponseAttributesAllOf) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *TLSCertificateResponseAttributesAllOf) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetSerialNumber

`func (o *TLSCertificateResponseAttributesAllOf) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *TLSCertificateResponseAttributesAllOf) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *TLSCertificateResponseAttributesAllOf) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *TLSCertificateResponseAttributesAllOf) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetSignatureAlgorithm

`func (o *TLSCertificateResponseAttributesAllOf) GetSignatureAlgorithm() string`

GetSignatureAlgorithm returns the SignatureAlgorithm field if non-nil, zero value otherwise.

### GetSignatureAlgorithmOk

`func (o *TLSCertificateResponseAttributesAllOf) GetSignatureAlgorithmOk() (*string, bool)`

GetSignatureAlgorithmOk returns a tuple with the SignatureAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureAlgorithm

`func (o *TLSCertificateResponseAttributesAllOf) SetSignatureAlgorithm(v string)`

SetSignatureAlgorithm sets SignatureAlgorithm field to given value.

### HasSignatureAlgorithm

`func (o *TLSCertificateResponseAttributesAllOf) HasSignatureAlgorithm() bool`

HasSignatureAlgorithm returns a boolean if a field has been set.

### GetNotAfter

`func (o *TLSCertificateResponseAttributesAllOf) GetNotAfter() time.Time`

GetNotAfter returns the NotAfter field if non-nil, zero value otherwise.

### GetNotAfterOk

`func (o *TLSCertificateResponseAttributesAllOf) GetNotAfterOk() (*time.Time, bool)`

GetNotAfterOk returns a tuple with the NotAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAfter

`func (o *TLSCertificateResponseAttributesAllOf) SetNotAfter(v time.Time)`

SetNotAfter sets NotAfter field to given value.

### HasNotAfter

`func (o *TLSCertificateResponseAttributesAllOf) HasNotAfter() bool`

HasNotAfter returns a boolean if a field has been set.

### GetNotBefore

`func (o *TLSCertificateResponseAttributesAllOf) GetNotBefore() time.Time`

GetNotBefore returns the NotBefore field if non-nil, zero value otherwise.

### GetNotBeforeOk

`func (o *TLSCertificateResponseAttributesAllOf) GetNotBeforeOk() (*time.Time, bool)`

GetNotBeforeOk returns a tuple with the NotBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotBefore

`func (o *TLSCertificateResponseAttributesAllOf) SetNotBefore(v time.Time)`

SetNotBefore sets NotBefore field to given value.

### HasNotBefore

`func (o *TLSCertificateResponseAttributesAllOf) HasNotBefore() bool`

HasNotBefore returns a boolean if a field has been set.

### GetReplace

`func (o *TLSCertificateResponseAttributesAllOf) GetReplace() bool`

GetReplace returns the Replace field if non-nil, zero value otherwise.

### GetReplaceOk

`func (o *TLSCertificateResponseAttributesAllOf) GetReplaceOk() (*bool, bool)`

GetReplaceOk returns a tuple with the Replace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplace

`func (o *TLSCertificateResponseAttributesAllOf) SetReplace(v bool)`

SetReplace sets Replace field to given value.

### HasReplace

`func (o *TLSCertificateResponseAttributesAllOf) HasReplace() bool`

HasReplace returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
