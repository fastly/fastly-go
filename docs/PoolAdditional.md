# PoolAdditional

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name for the Pool. | [optional] 
**Shield** | Pointer to **NullableString** | Selected POP to serve as a shield for the servers. Defaults to `null` meaning no origin shielding if not set. Refer to the [POPs API endpoint](https://www.fastly.com/documentation/reference/api/utils/pops/) to get a list of available POPs used for shielding. | [optional] [default to "null"]
**RequestCondition** | Pointer to **NullableString** | Condition which, if met, will select this configuration during a request. Optional. | [optional] 
**TlsCiphers** | Pointer to **NullableString** | List of OpenSSL ciphers (see the [openssl.org manpages](https://www.openssl.org/docs/man1.1.1/man1/ciphers.html) for details). Optional. | [optional] 
**TlsSniHostname** | Pointer to **NullableString** | SNI hostname. Optional. | [optional] 
**MinTlsVersion** | Pointer to **NullableInt32** | Minimum allowed TLS version on connections to this server. Optional. | [optional] 
**MaxTlsVersion** | Pointer to **NullableInt32** | Maximum allowed TLS version on connections to this server. Optional. | [optional] 
**Healthcheck** | Pointer to **NullableString** | Name of the healthcheck to use with this pool. Can be empty and could be reused across multiple backend and pools. | [optional] 
**Comment** | Pointer to **NullableString** | A freeform descriptive note. | [optional] 
**Type** | Pointer to **string** | What type of load balance group to use. | [optional] 
**OverrideHost** | Pointer to **NullableString** | The hostname to [override the Host header](https://www.fastly.com/documentation/guides/full-site-delivery/domains-and-origins/specifying-an-override-host/). Defaults to `null` meaning no override of the Host header will occur. This setting can also be added to a Server definition. If the field is set on a Server definition it will override the Pool setting. | [optional] [default to "null"]

## Methods

### NewPoolAdditional

`func NewPoolAdditional() *PoolAdditional`

NewPoolAdditional instantiates a new PoolAdditional object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolAdditionalWithDefaults

`func NewPoolAdditionalWithDefaults() *PoolAdditional`

NewPoolAdditionalWithDefaults instantiates a new PoolAdditional object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PoolAdditional) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PoolAdditional) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PoolAdditional) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PoolAdditional) HasName() bool`

HasName returns a boolean if a field has been set.

### GetShield

`func (o *PoolAdditional) GetShield() string`

GetShield returns the Shield field if non-nil, zero value otherwise.

### GetShieldOk

`func (o *PoolAdditional) GetShieldOk() (*string, bool)`

GetShieldOk returns a tuple with the Shield field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShield

`func (o *PoolAdditional) SetShield(v string)`

SetShield sets Shield field to given value.

### HasShield

`func (o *PoolAdditional) HasShield() bool`

HasShield returns a boolean if a field has been set.

### SetShieldNil

`func (o *PoolAdditional) SetShieldNil(b bool)`

 SetShieldNil sets the value for Shield to be an explicit nil

### UnsetShield
`func (o *PoolAdditional) UnsetShield()`

UnsetShield ensures that no value is present for Shield, not even an explicit nil
### GetRequestCondition

`func (o *PoolAdditional) GetRequestCondition() string`

GetRequestCondition returns the RequestCondition field if non-nil, zero value otherwise.

### GetRequestConditionOk

`func (o *PoolAdditional) GetRequestConditionOk() (*string, bool)`

GetRequestConditionOk returns a tuple with the RequestCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCondition

`func (o *PoolAdditional) SetRequestCondition(v string)`

SetRequestCondition sets RequestCondition field to given value.

### HasRequestCondition

`func (o *PoolAdditional) HasRequestCondition() bool`

HasRequestCondition returns a boolean if a field has been set.

### SetRequestConditionNil

`func (o *PoolAdditional) SetRequestConditionNil(b bool)`

 SetRequestConditionNil sets the value for RequestCondition to be an explicit nil

### UnsetRequestCondition
`func (o *PoolAdditional) UnsetRequestCondition()`

UnsetRequestCondition ensures that no value is present for RequestCondition, not even an explicit nil
### GetTlsCiphers

`func (o *PoolAdditional) GetTlsCiphers() string`

GetTlsCiphers returns the TlsCiphers field if non-nil, zero value otherwise.

### GetTlsCiphersOk

`func (o *PoolAdditional) GetTlsCiphersOk() (*string, bool)`

GetTlsCiphersOk returns a tuple with the TlsCiphers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsCiphers

`func (o *PoolAdditional) SetTlsCiphers(v string)`

SetTlsCiphers sets TlsCiphers field to given value.

### HasTlsCiphers

`func (o *PoolAdditional) HasTlsCiphers() bool`

HasTlsCiphers returns a boolean if a field has been set.

### SetTlsCiphersNil

`func (o *PoolAdditional) SetTlsCiphersNil(b bool)`

 SetTlsCiphersNil sets the value for TlsCiphers to be an explicit nil

### UnsetTlsCiphers
`func (o *PoolAdditional) UnsetTlsCiphers()`

UnsetTlsCiphers ensures that no value is present for TlsCiphers, not even an explicit nil
### GetTlsSniHostname

`func (o *PoolAdditional) GetTlsSniHostname() string`

GetTlsSniHostname returns the TlsSniHostname field if non-nil, zero value otherwise.

### GetTlsSniHostnameOk

`func (o *PoolAdditional) GetTlsSniHostnameOk() (*string, bool)`

GetTlsSniHostnameOk returns a tuple with the TlsSniHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsSniHostname

`func (o *PoolAdditional) SetTlsSniHostname(v string)`

SetTlsSniHostname sets TlsSniHostname field to given value.

### HasTlsSniHostname

`func (o *PoolAdditional) HasTlsSniHostname() bool`

HasTlsSniHostname returns a boolean if a field has been set.

### SetTlsSniHostnameNil

`func (o *PoolAdditional) SetTlsSniHostnameNil(b bool)`

 SetTlsSniHostnameNil sets the value for TlsSniHostname to be an explicit nil

### UnsetTlsSniHostname
`func (o *PoolAdditional) UnsetTlsSniHostname()`

UnsetTlsSniHostname ensures that no value is present for TlsSniHostname, not even an explicit nil
### GetMinTlsVersion

`func (o *PoolAdditional) GetMinTlsVersion() int32`

GetMinTlsVersion returns the MinTlsVersion field if non-nil, zero value otherwise.

### GetMinTlsVersionOk

`func (o *PoolAdditional) GetMinTlsVersionOk() (*int32, bool)`

GetMinTlsVersionOk returns a tuple with the MinTlsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinTlsVersion

`func (o *PoolAdditional) SetMinTlsVersion(v int32)`

SetMinTlsVersion sets MinTlsVersion field to given value.

### HasMinTlsVersion

`func (o *PoolAdditional) HasMinTlsVersion() bool`

HasMinTlsVersion returns a boolean if a field has been set.

### SetMinTlsVersionNil

`func (o *PoolAdditional) SetMinTlsVersionNil(b bool)`

 SetMinTlsVersionNil sets the value for MinTlsVersion to be an explicit nil

### UnsetMinTlsVersion
`func (o *PoolAdditional) UnsetMinTlsVersion()`

UnsetMinTlsVersion ensures that no value is present for MinTlsVersion, not even an explicit nil
### GetMaxTlsVersion

`func (o *PoolAdditional) GetMaxTlsVersion() int32`

GetMaxTlsVersion returns the MaxTlsVersion field if non-nil, zero value otherwise.

### GetMaxTlsVersionOk

`func (o *PoolAdditional) GetMaxTlsVersionOk() (*int32, bool)`

GetMaxTlsVersionOk returns a tuple with the MaxTlsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTlsVersion

`func (o *PoolAdditional) SetMaxTlsVersion(v int32)`

SetMaxTlsVersion sets MaxTlsVersion field to given value.

### HasMaxTlsVersion

`func (o *PoolAdditional) HasMaxTlsVersion() bool`

HasMaxTlsVersion returns a boolean if a field has been set.

### SetMaxTlsVersionNil

`func (o *PoolAdditional) SetMaxTlsVersionNil(b bool)`

 SetMaxTlsVersionNil sets the value for MaxTlsVersion to be an explicit nil

### UnsetMaxTlsVersion
`func (o *PoolAdditional) UnsetMaxTlsVersion()`

UnsetMaxTlsVersion ensures that no value is present for MaxTlsVersion, not even an explicit nil
### GetHealthcheck

`func (o *PoolAdditional) GetHealthcheck() string`

GetHealthcheck returns the Healthcheck field if non-nil, zero value otherwise.

### GetHealthcheckOk

`func (o *PoolAdditional) GetHealthcheckOk() (*string, bool)`

GetHealthcheckOk returns a tuple with the Healthcheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthcheck

`func (o *PoolAdditional) SetHealthcheck(v string)`

SetHealthcheck sets Healthcheck field to given value.

### HasHealthcheck

`func (o *PoolAdditional) HasHealthcheck() bool`

HasHealthcheck returns a boolean if a field has been set.

### SetHealthcheckNil

`func (o *PoolAdditional) SetHealthcheckNil(b bool)`

 SetHealthcheckNil sets the value for Healthcheck to be an explicit nil

### UnsetHealthcheck
`func (o *PoolAdditional) UnsetHealthcheck()`

UnsetHealthcheck ensures that no value is present for Healthcheck, not even an explicit nil
### GetComment

`func (o *PoolAdditional) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *PoolAdditional) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *PoolAdditional) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *PoolAdditional) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *PoolAdditional) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *PoolAdditional) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetType

`func (o *PoolAdditional) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PoolAdditional) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PoolAdditional) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PoolAdditional) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOverrideHost

`func (o *PoolAdditional) GetOverrideHost() string`

GetOverrideHost returns the OverrideHost field if non-nil, zero value otherwise.

### GetOverrideHostOk

`func (o *PoolAdditional) GetOverrideHostOk() (*string, bool)`

GetOverrideHostOk returns a tuple with the OverrideHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideHost

`func (o *PoolAdditional) SetOverrideHost(v string)`

SetOverrideHost sets OverrideHost field to given value.

### HasOverrideHost

`func (o *PoolAdditional) HasOverrideHost() bool`

HasOverrideHost returns a boolean if a field has been set.

### SetOverrideHostNil

`func (o *PoolAdditional) SetOverrideHostNil(b bool)`

 SetOverrideHostNil sets the value for OverrideHost to be an explicit nil

### UnsetOverrideHost
`func (o *PoolAdditional) UnsetOverrideHost()`

UnsetOverrideHost ensures that no value is present for OverrideHost, not even an explicit nil

[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


