# PoolAdditional

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name for the Pool. | [optional] 
**Shield** | Pointer to **NullableString** | Selected POP to serve as a shield for the servers. Defaults to `null` meaning no origin shielding if not set. Refer to the [POPs API endpoint](https://www.fastly.com/documentation/reference/api/utils/pops/) to get a list of available POPs used for shielding. | [optional] [default to "null"]
**RequestCondition** | Pointer to **NullableString** | Condition which, if met, will select this configuration during a request. Optional. | [optional] 
**TLSCiphers** | Pointer to **NullableString** | List of OpenSSL ciphers (see the [openssl.org manpages](https://www.openssl.org/docs/man1.1.1/man1/ciphers.html) for details). Optional. | [optional] 
**TLSSniHostname** | Pointer to **NullableString** | SNI hostname. Optional. | [optional] 
**MinTLSVersion** | Pointer to **NullableInt32** | Minimum allowed TLS version on connections to this server. Optional. | [optional] 
**MaxTLSVersion** | Pointer to **NullableInt32** | Maximum allowed TLS version on connections to this server. Optional. | [optional] 
**Healthcheck** | Pointer to **NullableString** | Name of the healthcheck to use with this pool. Can be empty and could be reused across multiple backend and pools. | [optional] 
**Comment** | Pointer to **NullableString** | A freeform descriptive note. | [optional] 
**Type** | Pointer to **string** | What type of load balance group to use. | [optional] 
**OverrideHost** | Pointer to **NullableString** | The hostname to [override the Host header](https://docs.fastly.com/en/guides/specifying-an-override-host). Defaults to `null` meaning no override of the Host header will occur. This setting can also be added to a Server definition. If the field is set on a Server definition it will override the Pool setting. | [optional] [default to "null"]

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
### GetTLSCiphers

`func (o *PoolAdditional) GetTLSCiphers() string`

GetTLSCiphers returns the TLSCiphers field if non-nil, zero value otherwise.

### GetTLSCiphersOk

`func (o *PoolAdditional) GetTLSCiphersOk() (*string, bool)`

GetTLSCiphersOk returns a tuple with the TLSCiphers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTLSCiphers

`func (o *PoolAdditional) SetTLSCiphers(v string)`

SetTLSCiphers sets TLSCiphers field to given value.

### HasTLSCiphers

`func (o *PoolAdditional) HasTLSCiphers() bool`

HasTLSCiphers returns a boolean if a field has been set.

### SetTLSCiphersNil

`func (o *PoolAdditional) SetTLSCiphersNil(b bool)`

 SetTLSCiphersNil sets the value for TLSCiphers to be an explicit nil

### UnsetTLSCiphers
`func (o *PoolAdditional) UnsetTLSCiphers()`

UnsetTLSCiphers ensures that no value is present for TLSCiphers, not even an explicit nil
### GetTLSSniHostname

`func (o *PoolAdditional) GetTLSSniHostname() string`

GetTLSSniHostname returns the TLSSniHostname field if non-nil, zero value otherwise.

### GetTLSSniHostnameOk

`func (o *PoolAdditional) GetTLSSniHostnameOk() (*string, bool)`

GetTLSSniHostnameOk returns a tuple with the TLSSniHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTLSSniHostname

`func (o *PoolAdditional) SetTLSSniHostname(v string)`

SetTLSSniHostname sets TLSSniHostname field to given value.

### HasTLSSniHostname

`func (o *PoolAdditional) HasTLSSniHostname() bool`

HasTLSSniHostname returns a boolean if a field has been set.

### SetTLSSniHostnameNil

`func (o *PoolAdditional) SetTLSSniHostnameNil(b bool)`

 SetTLSSniHostnameNil sets the value for TLSSniHostname to be an explicit nil

### UnsetTLSSniHostname
`func (o *PoolAdditional) UnsetTLSSniHostname()`

UnsetTLSSniHostname ensures that no value is present for TLSSniHostname, not even an explicit nil
### GetMinTLSVersion

`func (o *PoolAdditional) GetMinTLSVersion() int32`

GetMinTLSVersion returns the MinTLSVersion field if non-nil, zero value otherwise.

### GetMinTLSVersionOk

`func (o *PoolAdditional) GetMinTLSVersionOk() (*int32, bool)`

GetMinTLSVersionOk returns a tuple with the MinTLSVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinTLSVersion

`func (o *PoolAdditional) SetMinTLSVersion(v int32)`

SetMinTLSVersion sets MinTLSVersion field to given value.

### HasMinTLSVersion

`func (o *PoolAdditional) HasMinTLSVersion() bool`

HasMinTLSVersion returns a boolean if a field has been set.

### SetMinTLSVersionNil

`func (o *PoolAdditional) SetMinTLSVersionNil(b bool)`

 SetMinTLSVersionNil sets the value for MinTLSVersion to be an explicit nil

### UnsetMinTLSVersion
`func (o *PoolAdditional) UnsetMinTLSVersion()`

UnsetMinTLSVersion ensures that no value is present for MinTLSVersion, not even an explicit nil
### GetMaxTLSVersion

`func (o *PoolAdditional) GetMaxTLSVersion() int32`

GetMaxTLSVersion returns the MaxTLSVersion field if non-nil, zero value otherwise.

### GetMaxTLSVersionOk

`func (o *PoolAdditional) GetMaxTLSVersionOk() (*int32, bool)`

GetMaxTLSVersionOk returns a tuple with the MaxTLSVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTLSVersion

`func (o *PoolAdditional) SetMaxTLSVersion(v int32)`

SetMaxTLSVersion sets MaxTLSVersion field to given value.

### HasMaxTLSVersion

`func (o *PoolAdditional) HasMaxTLSVersion() bool`

HasMaxTLSVersion returns a boolean if a field has been set.

### SetMaxTLSVersionNil

`func (o *PoolAdditional) SetMaxTLSVersionNil(b bool)`

 SetMaxTLSVersionNil sets the value for MaxTLSVersion to be an explicit nil

### UnsetMaxTLSVersion
`func (o *PoolAdditional) UnsetMaxTLSVersion()`

UnsetMaxTLSVersion ensures that no value is present for MaxTLSVersion, not even an explicit nil
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
