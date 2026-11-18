# TLSBulkCertificateResponseAttributesAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotAfter** | Pointer to **time.Time** | Time-stamp (GMT) when the certificate will expire. Must be in the future to be used to terminate TLS traffic. | [optional] [readonly] 
**NotBefore** | Pointer to **time.Time** | Time-stamp (GMT) when the certificate will become valid. Must be in the past to be used to terminate TLS traffic. | [optional] [readonly] 
**Replace** | Pointer to **bool** | A recommendation from Fastly indicating the key associated with this certificate is in need of rotation. | [optional] [readonly] 

## Methods

### NewTLSBulkCertificateResponseAttributesAllOf

`func NewTLSBulkCertificateResponseAttributesAllOf() *TLSBulkCertificateResponseAttributesAllOf`

NewTLSBulkCertificateResponseAttributesAllOf instantiates a new TLSBulkCertificateResponseAttributesAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTLSBulkCertificateResponseAttributesAllOfWithDefaults

`func NewTLSBulkCertificateResponseAttributesAllOfWithDefaults() *TLSBulkCertificateResponseAttributesAllOf`

NewTLSBulkCertificateResponseAttributesAllOfWithDefaults instantiates a new TLSBulkCertificateResponseAttributesAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotAfter

`func (o *TLSBulkCertificateResponseAttributesAllOf) GetNotAfter() time.Time`

GetNotAfter returns the NotAfter field if non-nil, zero value otherwise.

### GetNotAfterOk

`func (o *TLSBulkCertificateResponseAttributesAllOf) GetNotAfterOk() (*time.Time, bool)`

GetNotAfterOk returns a tuple with the NotAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAfter

`func (o *TLSBulkCertificateResponseAttributesAllOf) SetNotAfter(v time.Time)`

SetNotAfter sets NotAfter field to given value.

### HasNotAfter

`func (o *TLSBulkCertificateResponseAttributesAllOf) HasNotAfter() bool`

HasNotAfter returns a boolean if a field has been set.

### GetNotBefore

`func (o *TLSBulkCertificateResponseAttributesAllOf) GetNotBefore() time.Time`

GetNotBefore returns the NotBefore field if non-nil, zero value otherwise.

### GetNotBeforeOk

`func (o *TLSBulkCertificateResponseAttributesAllOf) GetNotBeforeOk() (*time.Time, bool)`

GetNotBeforeOk returns a tuple with the NotBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotBefore

`func (o *TLSBulkCertificateResponseAttributesAllOf) SetNotBefore(v time.Time)`

SetNotBefore sets NotBefore field to given value.

### HasNotBefore

`func (o *TLSBulkCertificateResponseAttributesAllOf) HasNotBefore() bool`

HasNotBefore returns a boolean if a field has been set.

### GetReplace

`func (o *TLSBulkCertificateResponseAttributesAllOf) GetReplace() bool`

GetReplace returns the Replace field if non-nil, zero value otherwise.

### GetReplaceOk

`func (o *TLSBulkCertificateResponseAttributesAllOf) GetReplaceOk() (*bool, bool)`

GetReplaceOk returns a tuple with the Replace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplace

`func (o *TLSBulkCertificateResponseAttributesAllOf) SetReplace(v bool)`

SetReplace sets Replace field to given value.

### HasReplace

`func (o *TLSBulkCertificateResponseAttributesAllOf) HasReplace() bool`

HasReplace returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
