# TlsBulkCertificateDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowUntrustedRoot** | Pointer to **bool** | Allow certificates that chain to untrusted roots. | [optional] [default to false]
**CertBlob** | Pointer to **string** | The PEM-formatted certificate blob. Required. | [optional] 
**IntermediatesBlob** | Pointer to **string** | The PEM-formatted chain of intermediate blobs. Required. | [optional] 

## Methods

### NewTlsBulkCertificateDataAttributes

`func NewTlsBulkCertificateDataAttributes() *TlsBulkCertificateDataAttributes`

NewTlsBulkCertificateDataAttributes instantiates a new TlsBulkCertificateDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTlsBulkCertificateDataAttributesWithDefaults

`func NewTlsBulkCertificateDataAttributesWithDefaults() *TlsBulkCertificateDataAttributes`

NewTlsBulkCertificateDataAttributesWithDefaults instantiates a new TlsBulkCertificateDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowUntrustedRoot

`func (o *TlsBulkCertificateDataAttributes) GetAllowUntrustedRoot() bool`

GetAllowUntrustedRoot returns the AllowUntrustedRoot field if non-nil, zero value otherwise.

### GetAllowUntrustedRootOk

`func (o *TlsBulkCertificateDataAttributes) GetAllowUntrustedRootOk() (*bool, bool)`

GetAllowUntrustedRootOk returns a tuple with the AllowUntrustedRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUntrustedRoot

`func (o *TlsBulkCertificateDataAttributes) SetAllowUntrustedRoot(v bool)`

SetAllowUntrustedRoot sets AllowUntrustedRoot field to given value.

### HasAllowUntrustedRoot

`func (o *TlsBulkCertificateDataAttributes) HasAllowUntrustedRoot() bool`

HasAllowUntrustedRoot returns a boolean if a field has been set.

### GetCertBlob

`func (o *TlsBulkCertificateDataAttributes) GetCertBlob() string`

GetCertBlob returns the CertBlob field if non-nil, zero value otherwise.

### GetCertBlobOk

`func (o *TlsBulkCertificateDataAttributes) GetCertBlobOk() (*string, bool)`

GetCertBlobOk returns a tuple with the CertBlob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertBlob

`func (o *TlsBulkCertificateDataAttributes) SetCertBlob(v string)`

SetCertBlob sets CertBlob field to given value.

### HasCertBlob

`func (o *TlsBulkCertificateDataAttributes) HasCertBlob() bool`

HasCertBlob returns a boolean if a field has been set.

### GetIntermediatesBlob

`func (o *TlsBulkCertificateDataAttributes) GetIntermediatesBlob() string`

GetIntermediatesBlob returns the IntermediatesBlob field if non-nil, zero value otherwise.

### GetIntermediatesBlobOk

`func (o *TlsBulkCertificateDataAttributes) GetIntermediatesBlobOk() (*string, bool)`

GetIntermediatesBlobOk returns a tuple with the IntermediatesBlob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntermediatesBlob

`func (o *TlsBulkCertificateDataAttributes) SetIntermediatesBlob(v string)`

SetIntermediatesBlob sets IntermediatesBlob field to given value.

### HasIntermediatesBlob

`func (o *TlsBulkCertificateDataAttributes) HasIntermediatesBlob() bool`

HasIntermediatesBlob returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


