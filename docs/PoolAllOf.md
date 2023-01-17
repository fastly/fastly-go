# PoolAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewPoolAllOf

`func NewPoolAllOf() *PoolAllOf`

NewPoolAllOf instantiates a new PoolAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolAllOfWithDefaults

`func NewPoolAllOfWithDefaults() *PoolAllOf`

NewPoolAllOfWithDefaults instantiates a new PoolAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PoolAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PoolAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PoolAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PoolAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetShield

`func (o *PoolAllOf) GetShield() string`

GetShield returns the Shield field if non-nil, zero value otherwise.

### GetShieldOk

`func (o *PoolAllOf) GetShieldOk() (*string, bool)`

GetShieldOk returns a tuple with the Shield field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShield

`func (o *PoolAllOf) SetShield(v string)`

SetShield sets Shield field to given value.

### HasShield

`func (o *PoolAllOf) HasShield() bool`

HasShield returns a boolean if a field has been set.

### SetShieldNil

`func (o *PoolAllOf) SetShieldNil(b bool)`

 SetShieldNil sets the value for Shield to be an explicit nil

### UnsetShield
`func (o *PoolAllOf) UnsetShield()`

UnsetShield ensures that no value is present for Shield, not even an explicit nil
### GetRequestCondition

`func (o *PoolAllOf) GetRequestCondition() string`

GetRequestCondition returns the RequestCondition field if non-nil, zero value otherwise.

### GetRequestConditionOk

`func (o *PoolAllOf) GetRequestConditionOk() (*string, bool)`

GetRequestConditionOk returns a tuple with the RequestCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCondition

`func (o *PoolAllOf) SetRequestCondition(v string)`

SetRequestCondition sets RequestCondition field to given value.

### HasRequestCondition

`func (o *PoolAllOf) HasRequestCondition() bool`

HasRequestCondition returns a boolean if a field has been set.

### SetRequestConditionNil

`func (o *PoolAllOf) SetRequestConditionNil(b bool)`

 SetRequestConditionNil sets the value for RequestCondition to be an explicit nil

### UnsetRequestCondition
`func (o *PoolAllOf) UnsetRequestCondition()`

UnsetRequestCondition ensures that no value is present for RequestCondition, not even an explicit nil
### GetMaxConnDefault

`func (o *PoolAllOf) GetMaxConnDefault() int32`

GetMaxConnDefault returns the MaxConnDefault field if non-nil, zero value otherwise.

### GetMaxConnDefaultOk

`func (o *PoolAllOf) GetMaxConnDefaultOk() (*int32, bool)`

GetMaxConnDefaultOk returns a tuple with the MaxConnDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConnDefault

`func (o *PoolAllOf) SetMaxConnDefault(v int32)`

SetMaxConnDefault sets MaxConnDefault field to given value.

### HasMaxConnDefault

`func (o *PoolAllOf) HasMaxConnDefault() bool`

HasMaxConnDefault returns a boolean if a field has been set.

### GetConnectTimeout

`func (o *PoolAllOf) GetConnectTimeout() int32`

GetConnectTimeout returns the ConnectTimeout field if non-nil, zero value otherwise.

### GetConnectTimeoutOk

`func (o *PoolAllOf) GetConnectTimeoutOk() (*int32, bool)`

GetConnectTimeoutOk returns a tuple with the ConnectTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectTimeout

`func (o *PoolAllOf) SetConnectTimeout(v int32)`

SetConnectTimeout sets ConnectTimeout field to given value.

### HasConnectTimeout

`func (o *PoolAllOf) HasConnectTimeout() bool`

HasConnectTimeout returns a boolean if a field has been set.

### GetFirstByteTimeout

`func (o *PoolAllOf) GetFirstByteTimeout() int32`

GetFirstByteTimeout returns the FirstByteTimeout field if non-nil, zero value otherwise.

### GetFirstByteTimeoutOk

`func (o *PoolAllOf) GetFirstByteTimeoutOk() (*int32, bool)`

GetFirstByteTimeoutOk returns a tuple with the FirstByteTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstByteTimeout

`func (o *PoolAllOf) SetFirstByteTimeout(v int32)`

SetFirstByteTimeout sets FirstByteTimeout field to given value.

### HasFirstByteTimeout

`func (o *PoolAllOf) HasFirstByteTimeout() bool`

HasFirstByteTimeout returns a boolean if a field has been set.

### GetQuorum

`func (o *PoolAllOf) GetQuorum() int32`

GetQuorum returns the Quorum field if non-nil, zero value otherwise.

### GetQuorumOk

`func (o *PoolAllOf) GetQuorumOk() (*int32, bool)`

GetQuorumOk returns a tuple with the Quorum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuorum

`func (o *PoolAllOf) SetQuorum(v int32)`

SetQuorum sets Quorum field to given value.

### HasQuorum

`func (o *PoolAllOf) HasQuorum() bool`

HasQuorum returns a boolean if a field has been set.

### GetTLSCiphers

`func (o *PoolAllOf) GetTLSCiphers() string`

GetTLSCiphers returns the TLSCiphers field if non-nil, zero value otherwise.

### GetTLSCiphersOk

`func (o *PoolAllOf) GetTLSCiphersOk() (*string, bool)`

GetTLSCiphersOk returns a tuple with the TLSCiphers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTLSCiphers

`func (o *PoolAllOf) SetTLSCiphers(v string)`

SetTLSCiphers sets TLSCiphers field to given value.

### HasTLSCiphers

`func (o *PoolAllOf) HasTLSCiphers() bool`

HasTLSCiphers returns a boolean if a field has been set.

### SetTLSCiphersNil

`func (o *PoolAllOf) SetTLSCiphersNil(b bool)`

 SetTLSCiphersNil sets the value for TLSCiphers to be an explicit nil

### UnsetTLSCiphers
`func (o *PoolAllOf) UnsetTLSCiphers()`

UnsetTLSCiphers ensures that no value is present for TLSCiphers, not even an explicit nil
### GetTLSSniHostname

`func (o *PoolAllOf) GetTLSSniHostname() string`

GetTLSSniHostname returns the TLSSniHostname field if non-nil, zero value otherwise.

### GetTLSSniHostnameOk

`func (o *PoolAllOf) GetTLSSniHostnameOk() (*string, bool)`

GetTLSSniHostnameOk returns a tuple with the TLSSniHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTLSSniHostname

`func (o *PoolAllOf) SetTLSSniHostname(v string)`

SetTLSSniHostname sets TLSSniHostname field to given value.

### HasTLSSniHostname

`func (o *PoolAllOf) HasTLSSniHostname() bool`

HasTLSSniHostname returns a boolean if a field has been set.

### SetTLSSniHostnameNil

`func (o *PoolAllOf) SetTLSSniHostnameNil(b bool)`

 SetTLSSniHostnameNil sets the value for TLSSniHostname to be an explicit nil

### UnsetTLSSniHostname
`func (o *PoolAllOf) UnsetTLSSniHostname()`

UnsetTLSSniHostname ensures that no value is present for TLSSniHostname, not even an explicit nil
### GetTLSCheckCert

`func (o *PoolAllOf) GetTLSCheckCert() int32`

GetTLSCheckCert returns the TLSCheckCert field if non-nil, zero value otherwise.

### GetTLSCheckCertOk

`func (o *PoolAllOf) GetTLSCheckCertOk() (*int32, bool)`

GetTLSCheckCertOk returns a tuple with the TLSCheckCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTLSCheckCert

`func (o *PoolAllOf) SetTLSCheckCert(v int32)`

SetTLSCheckCert sets TLSCheckCert field to given value.

### HasTLSCheckCert

`func (o *PoolAllOf) HasTLSCheckCert() bool`

HasTLSCheckCert returns a boolean if a field has been set.

### SetTLSCheckCertNil

`func (o *PoolAllOf) SetTLSCheckCertNil(b bool)`

 SetTLSCheckCertNil sets the value for TLSCheckCert to be an explicit nil

### UnsetTLSCheckCert
`func (o *PoolAllOf) UnsetTLSCheckCert()`

UnsetTLSCheckCert ensures that no value is present for TLSCheckCert, not even an explicit nil
### GetMinTLSVersion

`func (o *PoolAllOf) GetMinTLSVersion() int32`

GetMinTLSVersion returns the MinTLSVersion field if non-nil, zero value otherwise.

### GetMinTLSVersionOk

`func (o *PoolAllOf) GetMinTLSVersionOk() (*int32, bool)`

GetMinTLSVersionOk returns a tuple with the MinTLSVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinTLSVersion

`func (o *PoolAllOf) SetMinTLSVersion(v int32)`

SetMinTLSVersion sets MinTLSVersion field to given value.

### HasMinTLSVersion

`func (o *PoolAllOf) HasMinTLSVersion() bool`

HasMinTLSVersion returns a boolean if a field has been set.

### SetMinTLSVersionNil

`func (o *PoolAllOf) SetMinTLSVersionNil(b bool)`

 SetMinTLSVersionNil sets the value for MinTLSVersion to be an explicit nil

### UnsetMinTLSVersion
`func (o *PoolAllOf) UnsetMinTLSVersion()`

UnsetMinTLSVersion ensures that no value is present for MinTLSVersion, not even an explicit nil
### GetMaxTLSVersion

`func (o *PoolAllOf) GetMaxTLSVersion() int32`

GetMaxTLSVersion returns the MaxTLSVersion field if non-nil, zero value otherwise.

### GetMaxTLSVersionOk

`func (o *PoolAllOf) GetMaxTLSVersionOk() (*int32, bool)`

GetMaxTLSVersionOk returns a tuple with the MaxTLSVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTLSVersion

`func (o *PoolAllOf) SetMaxTLSVersion(v int32)`

SetMaxTLSVersion sets MaxTLSVersion field to given value.

### HasMaxTLSVersion

`func (o *PoolAllOf) HasMaxTLSVersion() bool`

HasMaxTLSVersion returns a boolean if a field has been set.

### SetMaxTLSVersionNil

`func (o *PoolAllOf) SetMaxTLSVersionNil(b bool)`

 SetMaxTLSVersionNil sets the value for MaxTLSVersion to be an explicit nil

### UnsetMaxTLSVersion
`func (o *PoolAllOf) UnsetMaxTLSVersion()`

UnsetMaxTLSVersion ensures that no value is present for MaxTLSVersion, not even an explicit nil
### GetHealthcheck

`func (o *PoolAllOf) GetHealthcheck() string`

GetHealthcheck returns the Healthcheck field if non-nil, zero value otherwise.

### GetHealthcheckOk

`func (o *PoolAllOf) GetHealthcheckOk() (*string, bool)`

GetHealthcheckOk returns a tuple with the Healthcheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthcheck

`func (o *PoolAllOf) SetHealthcheck(v string)`

SetHealthcheck sets Healthcheck field to given value.

### HasHealthcheck

`func (o *PoolAllOf) HasHealthcheck() bool`

HasHealthcheck returns a boolean if a field has been set.

### SetHealthcheckNil

`func (o *PoolAllOf) SetHealthcheckNil(b bool)`

 SetHealthcheckNil sets the value for Healthcheck to be an explicit nil

### UnsetHealthcheck
`func (o *PoolAllOf) UnsetHealthcheck()`

UnsetHealthcheck ensures that no value is present for Healthcheck, not even an explicit nil
### GetComment

`func (o *PoolAllOf) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *PoolAllOf) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *PoolAllOf) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *PoolAllOf) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *PoolAllOf) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *PoolAllOf) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetType

`func (o *PoolAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PoolAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PoolAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PoolAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOverrideHost

`func (o *PoolAllOf) GetOverrideHost() string`

GetOverrideHost returns the OverrideHost field if non-nil, zero value otherwise.

### GetOverrideHostOk

`func (o *PoolAllOf) GetOverrideHostOk() (*string, bool)`

GetOverrideHostOk returns a tuple with the OverrideHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideHost

`func (o *PoolAllOf) SetOverrideHost(v string)`

SetOverrideHost sets OverrideHost field to given value.

### HasOverrideHost

`func (o *PoolAllOf) HasOverrideHost() bool`

HasOverrideHost returns a boolean if a field has been set.

### SetOverrideHostNil

`func (o *PoolAllOf) SetOverrideHostNil(b bool)`

 SetOverrideHostNil sets the value for OverrideHost to be an explicit nil

### UnsetOverrideHost
`func (o *PoolAllOf) UnsetOverrideHost()`

UnsetOverrideHost ensures that no value is present for OverrideHost, not even an explicit nil

[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
