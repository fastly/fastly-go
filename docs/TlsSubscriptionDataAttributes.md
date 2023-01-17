# TLSSubscriptionDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertificateAuthority** | Pointer to **string** | The entity that issues and certifies the TLS certificates for your subscription. | [optional] 

## Methods

### NewTLSSubscriptionDataAttributes

`func NewTLSSubscriptionDataAttributes() *TLSSubscriptionDataAttributes`

NewTLSSubscriptionDataAttributes instantiates a new TLSSubscriptionDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTLSSubscriptionDataAttributesWithDefaults

`func NewTLSSubscriptionDataAttributesWithDefaults() *TLSSubscriptionDataAttributes`

NewTLSSubscriptionDataAttributesWithDefaults instantiates a new TLSSubscriptionDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificateAuthority

`func (o *TLSSubscriptionDataAttributes) GetCertificateAuthority() string`

GetCertificateAuthority returns the CertificateAuthority field if non-nil, zero value otherwise.

### GetCertificateAuthorityOk

`func (o *TLSSubscriptionDataAttributes) GetCertificateAuthorityOk() (*string, bool)`

GetCertificateAuthorityOk returns a tuple with the CertificateAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateAuthority

`func (o *TLSSubscriptionDataAttributes) SetCertificateAuthority(v string)`

SetCertificateAuthority sets CertificateAuthority field to given value.

### HasCertificateAuthority

`func (o *TLSSubscriptionDataAttributes) HasCertificateAuthority() bool`

HasCertificateAuthority returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
