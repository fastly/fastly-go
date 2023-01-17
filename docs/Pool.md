# Pool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TLSCaCert** | Pointer to **NullableString** | A secure certificate to authenticate a server with. Must be in PEM format. | [optional] [default to "null"]
**TLSClientCert** | Pointer to **NullableString** | The client certificate used to make authenticated requests. Must be in PEM format. | [optional] [default to "null"]
**TLSClientKey** | Pointer to **NullableString** | The client private key used to make authenticated requests. Must be in PEM format. | [optional] [default to "null"]
**TLSCertHostname** | Pointer to **NullableString** | The hostname used to verify a server&#39;s certificate. It can either be the Common Name (CN) or a Subject Alternative Name (SAN). | [optional] [default to "null"]
**UseTLS** | Pointer to **int32** | Whether to use TLS. | [optional] [default to 0]
**Name** | Pointer to **string** | Name for the Pool. | [optional] 
**Shield** | Pointer to **NullableString** | Selected POP to serve as a shield for the servers. Defaults to `null` meaning no origin shielding if not set. Refer to the [POPs API endpoint](/reference/api/utils/pops/) to get a list of available POPs used for shielding. | [optional] [default to "null"]
**RequestCondition** | Pointer to **NullableString** | Condition which, if met, will select this configuration during a request. Optional. | [optional] 
**MaxConnDefault** | Pointer to **int32** | Maximum number of connections. Optional. | [optional] [default to 200]
**ConnectTimeout** | Pointer to **int32** | How long to wait for a timeout in milliseconds. Optional. | [optional] 
**FirstByteTimeout** | Pointer to **int32** | How long to wait for the first byte in milliseconds. Optional. | [optional] 
**Quorum** | Pointer to **int32** | Percentage of capacity (`0-100`) that needs to be operationally available for a pool to be considered up. | [optional] [default to 75]
**TLSCiphers** | Pointer to **NullableString** | List of OpenSSL ciphers (see the [openssl.org manpages](https://www.openssl.org/docs/man1.1.1/man1/ciphers.html) for details). Optional. | [optional] 
**TLSSniHostname** | Pointer to **NullableString** | SNI hostname. Optional. | [optional] 
**TLSCheckCert** | Pointer to **NullableInt32** | Be strict on checking TLS certs. Optional. | [optional] 
**MinTLSVersion** | Pointer to **NullableInt32** | Minimum allowed TLS version on connections to this server. Optional. | [optional] 
**MaxTLSVersion** | Pointer to **NullableInt32** | Maximum allowed TLS version on connections to this server. Optional. | [optional] 
**Healthcheck** | Pointer to **NullableString** | Name of the healthcheck to use with this pool. Can be empty and could be reused across multiple backend and pools. | [optional] 
**Comment** | Pointer to **NullableString** | A freeform descriptive note. | [optional] 
**Type** | Pointer to **string** | What type of load balance group to use. | [optional] 
**OverrideHost** | Pointer to **NullableString** | The hostname to [override the Host header](https://docs.fastly.com/en/guides/specifying-an-override-host). Defaults to `null` meaning no override of the Host header will occur. This setting can also be added to a Server definition. If the field is set on a Server definition it will override the Pool setting. | [optional] [default to "null"]

## Methods

### NewPool

`func NewPool() *Pool`

NewPool instantiates a new Pool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolWithDefaults

`func NewPoolWithDefaults() *Pool`

NewPoolWithDefaults instantiates a new Pool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTLSCaCert

`func (o *Pool) GetTLSCaCert() string`

GetTLSCaCert returns the TLSCaCert field if non-nil, zero value otherwise.

### GetTLSCaCertOk

`func (o *Pool) GetTLSCaCertOk() (*string, bool)`

GetTLSCaCertOk returns a tuple with the TLSCaCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTLSCaCert

`func (o *Pool) SetTLSCaCert(v string)`

SetTLSCaCert sets TLSCaCert field to given value.

### HasTLSCaCert

`func (o *Pool) HasTLSCaCert() bool`

HasTLSCaCert returns a boolean if a field has been set.

### SetTLSCaCertNil

`func (o *Pool) SetTLSCaCertNil(b bool)`

 SetTLSCaCertNil sets the value for TLSCaCert to be an explicit nil

### UnsetTLSCaCert
`func (o *Pool) UnsetTLSCaCert()`

UnsetTLSCaCert ensures that no value is present for TLSCaCert, not even an explicit nil
### GetTLSClientCert

`func (o *Pool) GetTLSClientCert() string`

GetTLSClientCert returns the TLSClientCert field if non-nil, zero value otherwise.

### GetTLSClientCertOk

`func (o *Pool) GetTLSClientCertOk() (*string, bool)`

GetTLSClientCertOk returns a tuple with the TLSClientCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTLSClientCert

`func (o *Pool) SetTLSClientCert(v string)`

SetTLSClientCert sets TLSClientCert field to given value.

### HasTLSClientCert

`func (o *Pool) HasTLSClientCert() bool`

HasTLSClientCert returns a boolean if a field has been set.

### SetTLSClientCertNil

`func (o *Pool) SetTLSClientCertNil(b bool)`

 SetTLSClientCertNil sets the value for TLSClientCert to be an explicit nil

### UnsetTLSClientCert
`func (o *Pool) UnsetTLSClientCert()`

UnsetTLSClientCert ensures that no value is present for TLSClientCert, not even an explicit nil
### GetTLSClientKey

`func (o *Pool) GetTLSClientKey() string`

GetTLSClientKey returns the TLSClientKey field if non-nil, zero value otherwise.

### GetTLSClientKeyOk

`func (o *Pool) GetTLSClientKeyOk() (*string, bool)`

GetTLSClientKeyOk returns a tuple with the TLSClientKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTLSClientKey

`func (o *Pool) SetTLSClientKey(v string)`

SetTLSClientKey sets TLSClientKey field to given value.

### HasTLSClientKey

`func (o *Pool) HasTLSClientKey() bool`

HasTLSClientKey returns a boolean if a field has been set.

### SetTLSClientKeyNil

`func (o *Pool) SetTLSClientKeyNil(b bool)`

 SetTLSClientKeyNil sets the value for TLSClientKey to be an explicit nil

### UnsetTLSClientKey
`func (o *Pool) UnsetTLSClientKey()`

UnsetTLSClientKey ensures that no value is present for TLSClientKey, not even an explicit nil
### GetTLSCertHostname

`func (o *Pool) GetTLSCertHostname() string`

GetTLSCertHostname returns the TLSCertHostname field if non-nil, zero value otherwise.

### GetTLSCertHostnameOk

`func (o *Pool) GetTLSCertHostnameOk() (*string, bool)`

GetTLSCertHostnameOk returns a tuple with the TLSCertHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTLSCertHostname

`func (o *Pool) SetTLSCertHostname(v string)`

SetTLSCertHostname sets TLSCertHostname field to given value.

### HasTLSCertHostname

`func (o *Pool) HasTLSCertHostname() bool`

HasTLSCertHostname returns a boolean if a field has been set.

### SetTLSCertHostnameNil

`func (o *Pool) SetTLSCertHostnameNil(b bool)`

 SetTLSCertHostnameNil sets the value for TLSCertHostname to be an explicit nil

### UnsetTLSCertHostname
`func (o *Pool) UnsetTLSCertHostname()`

UnsetTLSCertHostname ensures that no value is present for TLSCertHostname, not even an explicit nil
### GetUseTLS

`func (o *Pool) GetUseTLS() int32`

GetUseTLS returns the UseTLS field if non-nil, zero value otherwise.

### GetUseTLSOk

`func (o *Pool) GetUseTLSOk() (*int32, bool)`

GetUseTLSOk returns a tuple with the UseTLS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTLS

`func (o *Pool) SetUseTLS(v int32)`

SetUseTLS sets UseTLS field to given value.

### HasUseTLS

`func (o *Pool) HasUseTLS() bool`

HasUseTLS returns a boolean if a field has been set.

### GetName

`func (o *Pool) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Pool) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Pool) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Pool) HasName() bool`

HasName returns a boolean if a field has been set.

### GetShield

`func (o *Pool) GetShield() string`

GetShield returns the Shield field if non-nil, zero value otherwise.

### GetShieldOk

`func (o *Pool) GetShieldOk() (*string, bool)`

GetShieldOk returns a tuple with the Shield field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShield

`func (o *Pool) SetShield(v string)`

SetShield sets Shield field to given value.

### HasShield

`func (o *Pool) HasShield() bool`

HasShield returns a boolean if a field has been set.

### SetShieldNil

`func (o *Pool) SetShieldNil(b bool)`

 SetShieldNil sets the value for Shield to be an explicit nil

### UnsetShield
`func (o *Pool) UnsetShield()`

UnsetShield ensures that no value is present for Shield, not even an explicit nil
### GetRequestCondition

`func (o *Pool) GetRequestCondition() string`

GetRequestCondition returns the RequestCondition field if non-nil, zero value otherwise.

### GetRequestConditionOk

`func (o *Pool) GetRequestConditionOk() (*string, bool)`

GetRequestConditionOk returns a tuple with the RequestCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCondition

`func (o *Pool) SetRequestCondition(v string)`

SetRequestCondition sets RequestCondition field to given value.

### HasRequestCondition

`func (o *Pool) HasRequestCondition() bool`

HasRequestCondition returns a boolean if a field has been set.

### SetRequestConditionNil

`func (o *Pool) SetRequestConditionNil(b bool)`

 SetRequestConditionNil sets the value for RequestCondition to be an explicit nil

### UnsetRequestCondition
`func (o *Pool) UnsetRequestCondition()`

UnsetRequestCondition ensures that no value is present for RequestCondition, not even an explicit nil
### GetMaxConnDefault

`func (o *Pool) GetMaxConnDefault() int32`

GetMaxConnDefault returns the MaxConnDefault field if non-nil, zero value otherwise.

### GetMaxConnDefaultOk

`func (o *Pool) GetMaxConnDefaultOk() (*int32, bool)`

GetMaxConnDefaultOk returns a tuple with the MaxConnDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConnDefault

`func (o *Pool) SetMaxConnDefault(v int32)`

SetMaxConnDefault sets MaxConnDefault field to given value.

### HasMaxConnDefault

`func (o *Pool) HasMaxConnDefault() bool`

HasMaxConnDefault returns a boolean if a field has been set.

### GetConnectTimeout

`func (o *Pool) GetConnectTimeout() int32`

GetConnectTimeout returns the ConnectTimeout field if non-nil, zero value otherwise.

### GetConnectTimeoutOk

`func (o *Pool) GetConnectTimeoutOk() (*int32, bool)`

GetConnectTimeoutOk returns a tuple with the ConnectTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectTimeout

`func (o *Pool) SetConnectTimeout(v int32)`

SetConnectTimeout sets ConnectTimeout field to given value.

### HasConnectTimeout

`func (o *Pool) HasConnectTimeout() bool`

HasConnectTimeout returns a boolean if a field has been set.

### GetFirstByteTimeout

`func (o *Pool) GetFirstByteTimeout() int32`

GetFirstByteTimeout returns the FirstByteTimeout field if non-nil, zero value otherwise.

### GetFirstByteTimeoutOk

`func (o *Pool) GetFirstByteTimeoutOk() (*int32, bool)`

GetFirstByteTimeoutOk returns a tuple with the FirstByteTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstByteTimeout

`func (o *Pool) SetFirstByteTimeout(v int32)`

SetFirstByteTimeout sets FirstByteTimeout field to given value.

### HasFirstByteTimeout

`func (o *Pool) HasFirstByteTimeout() bool`

HasFirstByteTimeout returns a boolean if a field has been set.

### GetQuorum

`func (o *Pool) GetQuorum() int32`

GetQuorum returns the Quorum field if non-nil, zero value otherwise.

### GetQuorumOk

`func (o *Pool) GetQuorumOk() (*int32, bool)`

GetQuorumOk returns a tuple with the Quorum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuorum

`func (o *Pool) SetQuorum(v int32)`

SetQuorum sets Quorum field to given value.

### HasQuorum

`func (o *Pool) HasQuorum() bool`

HasQuorum returns a boolean if a field has been set.

### GetTLSCiphers

`func (o *Pool) GetTLSCiphers() string`

GetTLSCiphers returns the TLSCiphers field if non-nil, zero value otherwise.

### GetTLSCiphersOk

`func (o *Pool) GetTLSCiphersOk() (*string, bool)`

GetTLSCiphersOk returns a tuple with the TLSCiphers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTLSCiphers

`func (o *Pool) SetTLSCiphers(v string)`

SetTLSCiphers sets TLSCiphers field to given value.

### HasTLSCiphers

`func (o *Pool) HasTLSCiphers() bool`

HasTLSCiphers returns a boolean if a field has been set.

### SetTLSCiphersNil

`func (o *Pool) SetTLSCiphersNil(b bool)`

 SetTLSCiphersNil sets the value for TLSCiphers to be an explicit nil

### UnsetTLSCiphers
`func (o *Pool) UnsetTLSCiphers()`

UnsetTLSCiphers ensures that no value is present for TLSCiphers, not even an explicit nil
### GetTLSSniHostname

`func (o *Pool) GetTLSSniHostname() string`

GetTLSSniHostname returns the TLSSniHostname field if non-nil, zero value otherwise.

### GetTLSSniHostnameOk

`func (o *Pool) GetTLSSniHostnameOk() (*string, bool)`

GetTLSSniHostnameOk returns a tuple with the TLSSniHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTLSSniHostname

`func (o *Pool) SetTLSSniHostname(v string)`

SetTLSSniHostname sets TLSSniHostname field to given value.

### HasTLSSniHostname

`func (o *Pool) HasTLSSniHostname() bool`

HasTLSSniHostname returns a boolean if a field has been set.

### SetTLSSniHostnameNil

`func (o *Pool) SetTLSSniHostnameNil(b bool)`

 SetTLSSniHostnameNil sets the value for TLSSniHostname to be an explicit nil

### UnsetTLSSniHostname
`func (o *Pool) UnsetTLSSniHostname()`

UnsetTLSSniHostname ensures that no value is present for TLSSniHostname, not even an explicit nil
### GetTLSCheckCert

`func (o *Pool) GetTLSCheckCert() int32`

GetTLSCheckCert returns the TLSCheckCert field if non-nil, zero value otherwise.

### GetTLSCheckCertOk

`func (o *Pool) GetTLSCheckCertOk() (*int32, bool)`

GetTLSCheckCertOk returns a tuple with the TLSCheckCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTLSCheckCert

`func (o *Pool) SetTLSCheckCert(v int32)`

SetTLSCheckCert sets TLSCheckCert field to given value.

### HasTLSCheckCert

`func (o *Pool) HasTLSCheckCert() bool`

HasTLSCheckCert returns a boolean if a field has been set.

### SetTLSCheckCertNil

`func (o *Pool) SetTLSCheckCertNil(b bool)`

 SetTLSCheckCertNil sets the value for TLSCheckCert to be an explicit nil

### UnsetTLSCheckCert
`func (o *Pool) UnsetTLSCheckCert()`

UnsetTLSCheckCert ensures that no value is present for TLSCheckCert, not even an explicit nil
### GetMinTLSVersion

`func (o *Pool) GetMinTLSVersion() int32`

GetMinTLSVersion returns the MinTLSVersion field if non-nil, zero value otherwise.

### GetMinTLSVersionOk

`func (o *Pool) GetMinTLSVersionOk() (*int32, bool)`

GetMinTLSVersionOk returns a tuple with the MinTLSVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinTLSVersion

`func (o *Pool) SetMinTLSVersion(v int32)`

SetMinTLSVersion sets MinTLSVersion field to given value.

### HasMinTLSVersion

`func (o *Pool) HasMinTLSVersion() bool`

HasMinTLSVersion returns a boolean if a field has been set.

### SetMinTLSVersionNil

`func (o *Pool) SetMinTLSVersionNil(b bool)`

 SetMinTLSVersionNil sets the value for MinTLSVersion to be an explicit nil

### UnsetMinTLSVersion
`func (o *Pool) UnsetMinTLSVersion()`

UnsetMinTLSVersion ensures that no value is present for MinTLSVersion, not even an explicit nil
### GetMaxTLSVersion

`func (o *Pool) GetMaxTLSVersion() int32`

GetMaxTLSVersion returns the MaxTLSVersion field if non-nil, zero value otherwise.

### GetMaxTLSVersionOk

`func (o *Pool) GetMaxTLSVersionOk() (*int32, bool)`

GetMaxTLSVersionOk returns a tuple with the MaxTLSVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTLSVersion

`func (o *Pool) SetMaxTLSVersion(v int32)`

SetMaxTLSVersion sets MaxTLSVersion field to given value.

### HasMaxTLSVersion

`func (o *Pool) HasMaxTLSVersion() bool`

HasMaxTLSVersion returns a boolean if a field has been set.

### SetMaxTLSVersionNil

`func (o *Pool) SetMaxTLSVersionNil(b bool)`

 SetMaxTLSVersionNil sets the value for MaxTLSVersion to be an explicit nil

### UnsetMaxTLSVersion
`func (o *Pool) UnsetMaxTLSVersion()`

UnsetMaxTLSVersion ensures that no value is present for MaxTLSVersion, not even an explicit nil
### GetHealthcheck

`func (o *Pool) GetHealthcheck() string`

GetHealthcheck returns the Healthcheck field if non-nil, zero value otherwise.

### GetHealthcheckOk

`func (o *Pool) GetHealthcheckOk() (*string, bool)`

GetHealthcheckOk returns a tuple with the Healthcheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthcheck

`func (o *Pool) SetHealthcheck(v string)`

SetHealthcheck sets Healthcheck field to given value.

### HasHealthcheck

`func (o *Pool) HasHealthcheck() bool`

HasHealthcheck returns a boolean if a field has been set.

### SetHealthcheckNil

`func (o *Pool) SetHealthcheckNil(b bool)`

 SetHealthcheckNil sets the value for Healthcheck to be an explicit nil

### UnsetHealthcheck
`func (o *Pool) UnsetHealthcheck()`

UnsetHealthcheck ensures that no value is present for Healthcheck, not even an explicit nil
### GetComment

`func (o *Pool) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Pool) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Pool) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Pool) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *Pool) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *Pool) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetType

`func (o *Pool) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Pool) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Pool) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Pool) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOverrideHost

`func (o *Pool) GetOverrideHost() string`

GetOverrideHost returns the OverrideHost field if non-nil, zero value otherwise.

### GetOverrideHostOk

`func (o *Pool) GetOverrideHostOk() (*string, bool)`

GetOverrideHostOk returns a tuple with the OverrideHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideHost

`func (o *Pool) SetOverrideHost(v string)`

SetOverrideHost sets OverrideHost field to given value.

### HasOverrideHost

`func (o *Pool) HasOverrideHost() bool`

HasOverrideHost returns a boolean if a field has been set.

### SetOverrideHostNil

`func (o *Pool) SetOverrideHostNil(b bool)`

 SetOverrideHostNil sets the value for OverrideHost to be an explicit nil

### UnsetOverrideHost
`func (o *Pool) UnsetOverrideHost()`

UnsetOverrideHost ensures that no value is present for OverrideHost, not even an explicit nil

[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
