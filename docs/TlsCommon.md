# TLSCommon

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TLSCaCert** | Pointer to **NullableString** | A secure certificate to authenticate a server with. Must be in PEM format. | [optional] [default to "null"]
**TLSClientCert** | Pointer to **NullableString** | The client certificate used to make authenticated requests. Must be in PEM format. | [optional] [default to "null"]
**TLSClientKey** | Pointer to **NullableString** | The client private key used to make authenticated requests. Must be in PEM format. | [optional] [default to "null"]
**TLSCertHostname** | Pointer to **NullableString** | The hostname used to verify a server&#39;s certificate. It can either be the Common Name (CN) or a Subject Alternative Name (SAN). | [optional] [default to "null"]
**UseTLS** | Pointer to **int32** | Whether to use TLS. | [optional] [default to 0]

## Methods

### NewTLSCommon

`func NewTLSCommon() *TLSCommon`

NewTLSCommon instantiates a new TLSCommon object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTLSCommonWithDefaults

`func NewTLSCommonWithDefaults() *TLSCommon`

NewTLSCommonWithDefaults instantiates a new TLSCommon object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTLSCaCert

`func (o *TLSCommon) GetTLSCaCert() string`

GetTLSCaCert returns the TLSCaCert field if non-nil, zero value otherwise.

### GetTLSCaCertOk

`func (o *TLSCommon) GetTLSCaCertOk() (*string, bool)`

GetTLSCaCertOk returns a tuple with the TLSCaCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTLSCaCert

`func (o *TLSCommon) SetTLSCaCert(v string)`

SetTLSCaCert sets TLSCaCert field to given value.

### HasTLSCaCert

`func (o *TLSCommon) HasTLSCaCert() bool`

HasTLSCaCert returns a boolean if a field has been set.

### SetTLSCaCertNil

`func (o *TLSCommon) SetTLSCaCertNil(b bool)`

 SetTLSCaCertNil sets the value for TLSCaCert to be an explicit nil

### UnsetTLSCaCert
`func (o *TLSCommon) UnsetTLSCaCert()`

UnsetTLSCaCert ensures that no value is present for TLSCaCert, not even an explicit nil
### GetTLSClientCert

`func (o *TLSCommon) GetTLSClientCert() string`

GetTLSClientCert returns the TLSClientCert field if non-nil, zero value otherwise.

### GetTLSClientCertOk

`func (o *TLSCommon) GetTLSClientCertOk() (*string, bool)`

GetTLSClientCertOk returns a tuple with the TLSClientCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTLSClientCert

`func (o *TLSCommon) SetTLSClientCert(v string)`

SetTLSClientCert sets TLSClientCert field to given value.

### HasTLSClientCert

`func (o *TLSCommon) HasTLSClientCert() bool`

HasTLSClientCert returns a boolean if a field has been set.

### SetTLSClientCertNil

`func (o *TLSCommon) SetTLSClientCertNil(b bool)`

 SetTLSClientCertNil sets the value for TLSClientCert to be an explicit nil

### UnsetTLSClientCert
`func (o *TLSCommon) UnsetTLSClientCert()`

UnsetTLSClientCert ensures that no value is present for TLSClientCert, not even an explicit nil
### GetTLSClientKey

`func (o *TLSCommon) GetTLSClientKey() string`

GetTLSClientKey returns the TLSClientKey field if non-nil, zero value otherwise.

### GetTLSClientKeyOk

`func (o *TLSCommon) GetTLSClientKeyOk() (*string, bool)`

GetTLSClientKeyOk returns a tuple with the TLSClientKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTLSClientKey

`func (o *TLSCommon) SetTLSClientKey(v string)`

SetTLSClientKey sets TLSClientKey field to given value.

### HasTLSClientKey

`func (o *TLSCommon) HasTLSClientKey() bool`

HasTLSClientKey returns a boolean if a field has been set.

### SetTLSClientKeyNil

`func (o *TLSCommon) SetTLSClientKeyNil(b bool)`

 SetTLSClientKeyNil sets the value for TLSClientKey to be an explicit nil

### UnsetTLSClientKey
`func (o *TLSCommon) UnsetTLSClientKey()`

UnsetTLSClientKey ensures that no value is present for TLSClientKey, not even an explicit nil
### GetTLSCertHostname

`func (o *TLSCommon) GetTLSCertHostname() string`

GetTLSCertHostname returns the TLSCertHostname field if non-nil, zero value otherwise.

### GetTLSCertHostnameOk

`func (o *TLSCommon) GetTLSCertHostnameOk() (*string, bool)`

GetTLSCertHostnameOk returns a tuple with the TLSCertHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTLSCertHostname

`func (o *TLSCommon) SetTLSCertHostname(v string)`

SetTLSCertHostname sets TLSCertHostname field to given value.

### HasTLSCertHostname

`func (o *TLSCommon) HasTLSCertHostname() bool`

HasTLSCertHostname returns a boolean if a field has been set.

### SetTLSCertHostnameNil

`func (o *TLSCommon) SetTLSCertHostnameNil(b bool)`

 SetTLSCertHostnameNil sets the value for TLSCertHostname to be an explicit nil

### UnsetTLSCertHostname
`func (o *TLSCommon) UnsetTLSCertHostname()`

UnsetTLSCertHostname ensures that no value is present for TLSCertHostname, not even an explicit nil
### GetUseTLS

`func (o *TLSCommon) GetUseTLS() int32`

GetUseTLS returns the UseTLS field if non-nil, zero value otherwise.

### GetUseTLSOk

`func (o *TLSCommon) GetUseTLSOk() (*int32, bool)`

GetUseTLSOk returns a tuple with the UseTLS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTLS

`func (o *TLSCommon) SetUseTLS(v int32)`

SetUseTLS sets UseTLS field to given value.

### HasUseTLS

`func (o *TLSCommon) HasUseTLS() bool`

HasUseTLS returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
