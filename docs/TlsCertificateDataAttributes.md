# TLSCertificateDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertBlob** | Pointer to **string** | The PEM-formatted certificate blob. Required. | [optional] 
**Name** | Pointer to **string** | A customizable name for your certificate. Defaults to the certificate&#39;s Common Name or first Subject Alternative Names (SAN) entry. Optional. | [optional] 
**AllowUntrustedRoot** | Pointer to **bool** | Indicates that the supplied certificate was not signed by a trusted CA. | [optional] 

## Methods

### NewTLSCertificateDataAttributes

`func NewTLSCertificateDataAttributes() *TLSCertificateDataAttributes`

NewTLSCertificateDataAttributes instantiates a new TLSCertificateDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTLSCertificateDataAttributesWithDefaults

`func NewTLSCertificateDataAttributesWithDefaults() *TLSCertificateDataAttributes`

NewTLSCertificateDataAttributesWithDefaults instantiates a new TLSCertificateDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertBlob

`func (o *TLSCertificateDataAttributes) GetCertBlob() string`

GetCertBlob returns the CertBlob field if non-nil, zero value otherwise.

### GetCertBlobOk

`func (o *TLSCertificateDataAttributes) GetCertBlobOk() (*string, bool)`

GetCertBlobOk returns a tuple with the CertBlob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertBlob

`func (o *TLSCertificateDataAttributes) SetCertBlob(v string)`

SetCertBlob sets CertBlob field to given value.

### HasCertBlob

`func (o *TLSCertificateDataAttributes) HasCertBlob() bool`

HasCertBlob returns a boolean if a field has been set.

### GetName

`func (o *TLSCertificateDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TLSCertificateDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TLSCertificateDataAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TLSCertificateDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAllowUntrustedRoot

`func (o *TLSCertificateDataAttributes) GetAllowUntrustedRoot() bool`

GetAllowUntrustedRoot returns the AllowUntrustedRoot field if non-nil, zero value otherwise.

### GetAllowUntrustedRootOk

`func (o *TLSCertificateDataAttributes) GetAllowUntrustedRootOk() (*bool, bool)`

GetAllowUntrustedRootOk returns a tuple with the AllowUntrustedRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUntrustedRoot

`func (o *TLSCertificateDataAttributes) SetAllowUntrustedRoot(v bool)`

SetAllowUntrustedRoot sets AllowUntrustedRoot field to given value.

### HasAllowUntrustedRoot

`func (o *TLSCertificateDataAttributes) HasAllowUntrustedRoot() bool`

HasAllowUntrustedRoot returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
