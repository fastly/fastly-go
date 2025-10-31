# TlsCertificateDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertBlob** | Pointer to **string** | The PEM-formatted certificate blob. Required. | [optional] 
**Name** | Pointer to **string** | A customizable name for your certificate. Defaults to the certificate&#39;s Common Name or first Subject Alternative Names (SAN) entry. Optional. | [optional] 
**AllowUntrustedRoot** | Pointer to **bool** | Indicates that the supplied certificate was not signed by a trusted CA. | [optional] 

## Methods

### NewTlsCertificateDataAttributes

`func NewTlsCertificateDataAttributes() *TlsCertificateDataAttributes`

NewTlsCertificateDataAttributes instantiates a new TlsCertificateDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTlsCertificateDataAttributesWithDefaults

`func NewTlsCertificateDataAttributesWithDefaults() *TlsCertificateDataAttributes`

NewTlsCertificateDataAttributesWithDefaults instantiates a new TlsCertificateDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertBlob

`func (o *TlsCertificateDataAttributes) GetCertBlob() string`

GetCertBlob returns the CertBlob field if non-nil, zero value otherwise.

### GetCertBlobOk

`func (o *TlsCertificateDataAttributes) GetCertBlobOk() (*string, bool)`

GetCertBlobOk returns a tuple with the CertBlob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertBlob

`func (o *TlsCertificateDataAttributes) SetCertBlob(v string)`

SetCertBlob sets CertBlob field to given value.

### HasCertBlob

`func (o *TlsCertificateDataAttributes) HasCertBlob() bool`

HasCertBlob returns a boolean if a field has been set.

### GetName

`func (o *TlsCertificateDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TlsCertificateDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TlsCertificateDataAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TlsCertificateDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAllowUntrustedRoot

`func (o *TlsCertificateDataAttributes) GetAllowUntrustedRoot() bool`

GetAllowUntrustedRoot returns the AllowUntrustedRoot field if non-nil, zero value otherwise.

### GetAllowUntrustedRootOk

`func (o *TlsCertificateDataAttributes) GetAllowUntrustedRootOk() (*bool, bool)`

GetAllowUntrustedRootOk returns a tuple with the AllowUntrustedRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUntrustedRoot

`func (o *TlsCertificateDataAttributes) SetAllowUntrustedRoot(v bool)`

SetAllowUntrustedRoot sets AllowUntrustedRoot field to given value.

### HasAllowUntrustedRoot

`func (o *TlsCertificateDataAttributes) HasAllowUntrustedRoot() bool`

HasAllowUntrustedRoot returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


