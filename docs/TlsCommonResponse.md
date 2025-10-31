# TlsCommonResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TlsCaCert** | Pointer to **NullableString** | A secure certificate to authenticate a server with. Must be in PEM format. | [optional] [default to "null"]
**TlsClientCert** | Pointer to **NullableString** | The client certificate used to make authenticated requests. Must be in PEM format. | [optional] [default to "null"]
**TlsClientKey** | Pointer to **NullableString** | The client private key used to make authenticated requests. Must be in PEM format. | [optional] [default to "null"]
**TlsCertHostname** | Pointer to **NullableString** | The hostname used to verify a server&#39;s certificate. It can either be the Common Name (CN) or a Subject Alternative Name (SAN). | [optional] [default to "null"]
**UseTls** | Pointer to **string** | Whether to use TLS. | [optional] [default to "0"]

## Methods

### NewTlsCommonResponse

`func NewTlsCommonResponse() *TlsCommonResponse`

NewTlsCommonResponse instantiates a new TlsCommonResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTlsCommonResponseWithDefaults

`func NewTlsCommonResponseWithDefaults() *TlsCommonResponse`

NewTlsCommonResponseWithDefaults instantiates a new TlsCommonResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTlsCaCert

`func (o *TlsCommonResponse) GetTlsCaCert() string`

GetTlsCaCert returns the TlsCaCert field if non-nil, zero value otherwise.

### GetTlsCaCertOk

`func (o *TlsCommonResponse) GetTlsCaCertOk() (*string, bool)`

GetTlsCaCertOk returns a tuple with the TlsCaCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsCaCert

`func (o *TlsCommonResponse) SetTlsCaCert(v string)`

SetTlsCaCert sets TlsCaCert field to given value.

### HasTlsCaCert

`func (o *TlsCommonResponse) HasTlsCaCert() bool`

HasTlsCaCert returns a boolean if a field has been set.

### SetTlsCaCertNil

`func (o *TlsCommonResponse) SetTlsCaCertNil(b bool)`

 SetTlsCaCertNil sets the value for TlsCaCert to be an explicit nil

### UnsetTlsCaCert
`func (o *TlsCommonResponse) UnsetTlsCaCert()`

UnsetTlsCaCert ensures that no value is present for TlsCaCert, not even an explicit nil
### GetTlsClientCert

`func (o *TlsCommonResponse) GetTlsClientCert() string`

GetTlsClientCert returns the TlsClientCert field if non-nil, zero value otherwise.

### GetTlsClientCertOk

`func (o *TlsCommonResponse) GetTlsClientCertOk() (*string, bool)`

GetTlsClientCertOk returns a tuple with the TlsClientCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsClientCert

`func (o *TlsCommonResponse) SetTlsClientCert(v string)`

SetTlsClientCert sets TlsClientCert field to given value.

### HasTlsClientCert

`func (o *TlsCommonResponse) HasTlsClientCert() bool`

HasTlsClientCert returns a boolean if a field has been set.

### SetTlsClientCertNil

`func (o *TlsCommonResponse) SetTlsClientCertNil(b bool)`

 SetTlsClientCertNil sets the value for TlsClientCert to be an explicit nil

### UnsetTlsClientCert
`func (o *TlsCommonResponse) UnsetTlsClientCert()`

UnsetTlsClientCert ensures that no value is present for TlsClientCert, not even an explicit nil
### GetTlsClientKey

`func (o *TlsCommonResponse) GetTlsClientKey() string`

GetTlsClientKey returns the TlsClientKey field if non-nil, zero value otherwise.

### GetTlsClientKeyOk

`func (o *TlsCommonResponse) GetTlsClientKeyOk() (*string, bool)`

GetTlsClientKeyOk returns a tuple with the TlsClientKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsClientKey

`func (o *TlsCommonResponse) SetTlsClientKey(v string)`

SetTlsClientKey sets TlsClientKey field to given value.

### HasTlsClientKey

`func (o *TlsCommonResponse) HasTlsClientKey() bool`

HasTlsClientKey returns a boolean if a field has been set.

### SetTlsClientKeyNil

`func (o *TlsCommonResponse) SetTlsClientKeyNil(b bool)`

 SetTlsClientKeyNil sets the value for TlsClientKey to be an explicit nil

### UnsetTlsClientKey
`func (o *TlsCommonResponse) UnsetTlsClientKey()`

UnsetTlsClientKey ensures that no value is present for TlsClientKey, not even an explicit nil
### GetTlsCertHostname

`func (o *TlsCommonResponse) GetTlsCertHostname() string`

GetTlsCertHostname returns the TlsCertHostname field if non-nil, zero value otherwise.

### GetTlsCertHostnameOk

`func (o *TlsCommonResponse) GetTlsCertHostnameOk() (*string, bool)`

GetTlsCertHostnameOk returns a tuple with the TlsCertHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsCertHostname

`func (o *TlsCommonResponse) SetTlsCertHostname(v string)`

SetTlsCertHostname sets TlsCertHostname field to given value.

### HasTlsCertHostname

`func (o *TlsCommonResponse) HasTlsCertHostname() bool`

HasTlsCertHostname returns a boolean if a field has been set.

### SetTlsCertHostnameNil

`func (o *TlsCommonResponse) SetTlsCertHostnameNil(b bool)`

 SetTlsCertHostnameNil sets the value for TlsCertHostname to be an explicit nil

### UnsetTlsCertHostname
`func (o *TlsCommonResponse) UnsetTlsCertHostname()`

UnsetTlsCertHostname ensures that no value is present for TlsCertHostname, not even an explicit nil
### GetUseTls

`func (o *TlsCommonResponse) GetUseTls() string`

GetUseTls returns the UseTls field if non-nil, zero value otherwise.

### GetUseTlsOk

`func (o *TlsCommonResponse) GetUseTlsOk() (*string, bool)`

GetUseTlsOk returns a tuple with the UseTls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTls

`func (o *TlsCommonResponse) SetUseTls(v string)`

SetUseTls sets UseTls field to given value.

### HasUseTls

`func (o *TlsCommonResponse) HasUseTls() bool`

HasUseTls returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


