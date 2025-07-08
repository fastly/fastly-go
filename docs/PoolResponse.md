# PoolResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TLSCaCert** | Pointer to **NullableString** | A secure certificate to authenticate a server with. Must be in PEM format. | [optional] [default to "null"]
**TLSClientCert** | Pointer to **NullableString** | The client certificate used to make authenticated requests. Must be in PEM format. | [optional] [default to "null"]
**TLSClientKey** | Pointer to **NullableString** | The client private key used to make authenticated requests. Must be in PEM format. | [optional] [default to "null"]
**TLSCertHostname** | Pointer to **NullableString** | The hostname used to verify a server&#39;s certificate. It can either be the Common Name (CN) or a Subject Alternative Name (SAN). | [optional] [default to "null"]
**UseTLS** | Pointer to **string** | Whether to use TLS. | [optional] [default to "0"]
**CreatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**DeletedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**UpdatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**ServiceID** | Pointer to **string** |  | [optional] [readonly] 
**Version** | Pointer to **string** |  | [optional] [readonly] 
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
**OverrideHost** | Pointer to **NullableString** | The hostname to [override the Host header](https://www.fastly.com/documentation/guides/full-site-delivery/domains-and-origins/specifying-an-override-host/). Defaults to `null` meaning no override of the Host header will occur. This setting can also be added to a Server definition. If the field is set on a Server definition it will override the Pool setting. | [optional] [default to "null"]
**BetweenBytesTimeout** | Pointer to **string** | Maximum duration in milliseconds that Fastly will wait while receiving no data on a download from a backend. If exceeded, the response received so far will be considered complete and the fetch will end. May be set at runtime using `bereq.between_bytes_timeout`. | [optional] 
**ConnectTimeout** | Pointer to **string** | How long to wait for a timeout in milliseconds. | [optional] 
**FirstByteTimeout** | Pointer to **string** | How long to wait for the first byte in milliseconds. | [optional] 
**MaxConnDefault** | Pointer to **string** | Maximum number of connections. | [optional] [default to "200"]
**TLSCheckCert** | Pointer to **NullableString** | Be strict on checking TLS certs. | [optional] 
**ID** | Pointer to **string** |  | [optional] [readonly] 
**Quorum** | Pointer to **string** | Percentage of capacity (`0-100`) that needs to be operationally available for a pool to be considered up. | [optional] [default to "75"]

## Methods

### NewPoolResponse

`func NewPoolResponse() *PoolResponse`

NewPoolResponse instantiates a new PoolResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolResponseWithDefaults

`func NewPoolResponseWithDefaults() *PoolResponse`

NewPoolResponseWithDefaults instantiates a new PoolResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTLSCaCert

`func (o *PoolResponse) GetTLSCaCert() string`

GetTLSCaCert returns the TLSCaCert field if non-nil, zero value otherwise.

### GetTLSCaCertOk

`func (o *PoolResponse) GetTLSCaCertOk() (*string, bool)`

GetTLSCaCertOk returns a tuple with the TLSCaCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTLSCaCert

`func (o *PoolResponse) SetTLSCaCert(v string)`

SetTLSCaCert sets TLSCaCert field to given value.

### HasTLSCaCert

`func (o *PoolResponse) HasTLSCaCert() bool`

HasTLSCaCert returns a boolean if a field has been set.

### SetTLSCaCertNil

`func (o *PoolResponse) SetTLSCaCertNil(b bool)`

 SetTLSCaCertNil sets the value for TLSCaCert to be an explicit nil

### UnsetTLSCaCert
`func (o *PoolResponse) UnsetTLSCaCert()`

UnsetTLSCaCert ensures that no value is present for TLSCaCert, not even an explicit nil
### GetTLSClientCert

`func (o *PoolResponse) GetTLSClientCert() string`

GetTLSClientCert returns the TLSClientCert field if non-nil, zero value otherwise.

### GetTLSClientCertOk

`func (o *PoolResponse) GetTLSClientCertOk() (*string, bool)`

GetTLSClientCertOk returns a tuple with the TLSClientCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTLSClientCert

`func (o *PoolResponse) SetTLSClientCert(v string)`

SetTLSClientCert sets TLSClientCert field to given value.

### HasTLSClientCert

`func (o *PoolResponse) HasTLSClientCert() bool`

HasTLSClientCert returns a boolean if a field has been set.

### SetTLSClientCertNil

`func (o *PoolResponse) SetTLSClientCertNil(b bool)`

 SetTLSClientCertNil sets the value for TLSClientCert to be an explicit nil

### UnsetTLSClientCert
`func (o *PoolResponse) UnsetTLSClientCert()`

UnsetTLSClientCert ensures that no value is present for TLSClientCert, not even an explicit nil
### GetTLSClientKey

`func (o *PoolResponse) GetTLSClientKey() string`

GetTLSClientKey returns the TLSClientKey field if non-nil, zero value otherwise.

### GetTLSClientKeyOk

`func (o *PoolResponse) GetTLSClientKeyOk() (*string, bool)`

GetTLSClientKeyOk returns a tuple with the TLSClientKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTLSClientKey

`func (o *PoolResponse) SetTLSClientKey(v string)`

SetTLSClientKey sets TLSClientKey field to given value.

### HasTLSClientKey

`func (o *PoolResponse) HasTLSClientKey() bool`

HasTLSClientKey returns a boolean if a field has been set.

### SetTLSClientKeyNil

`func (o *PoolResponse) SetTLSClientKeyNil(b bool)`

 SetTLSClientKeyNil sets the value for TLSClientKey to be an explicit nil

### UnsetTLSClientKey
`func (o *PoolResponse) UnsetTLSClientKey()`

UnsetTLSClientKey ensures that no value is present for TLSClientKey, not even an explicit nil
### GetTLSCertHostname

`func (o *PoolResponse) GetTLSCertHostname() string`

GetTLSCertHostname returns the TLSCertHostname field if non-nil, zero value otherwise.

### GetTLSCertHostnameOk

`func (o *PoolResponse) GetTLSCertHostnameOk() (*string, bool)`

GetTLSCertHostnameOk returns a tuple with the TLSCertHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTLSCertHostname

`func (o *PoolResponse) SetTLSCertHostname(v string)`

SetTLSCertHostname sets TLSCertHostname field to given value.

### HasTLSCertHostname

`func (o *PoolResponse) HasTLSCertHostname() bool`

HasTLSCertHostname returns a boolean if a field has been set.

### SetTLSCertHostnameNil

`func (o *PoolResponse) SetTLSCertHostnameNil(b bool)`

 SetTLSCertHostnameNil sets the value for TLSCertHostname to be an explicit nil

### UnsetTLSCertHostname
`func (o *PoolResponse) UnsetTLSCertHostname()`

UnsetTLSCertHostname ensures that no value is present for TLSCertHostname, not even an explicit nil
### GetUseTLS

`func (o *PoolResponse) GetUseTLS() string`

GetUseTLS returns the UseTLS field if non-nil, zero value otherwise.

### GetUseTLSOk

`func (o *PoolResponse) GetUseTLSOk() (*string, bool)`

GetUseTLSOk returns a tuple with the UseTLS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTLS

`func (o *PoolResponse) SetUseTLS(v string)`

SetUseTLS sets UseTLS field to given value.

### HasUseTLS

`func (o *PoolResponse) HasUseTLS() bool`

HasUseTLS returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PoolResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PoolResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PoolResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PoolResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *PoolResponse) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *PoolResponse) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetDeletedAt

`func (o *PoolResponse) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *PoolResponse) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *PoolResponse) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *PoolResponse) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *PoolResponse) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *PoolResponse) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetUpdatedAt

`func (o *PoolResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PoolResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PoolResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PoolResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *PoolResponse) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *PoolResponse) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetServiceID

`func (o *PoolResponse) GetServiceID() string`

GetServiceID returns the ServiceID field if non-nil, zero value otherwise.

### GetServiceIDOk

`func (o *PoolResponse) GetServiceIDOk() (*string, bool)`

GetServiceIDOk returns a tuple with the ServiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceID

`func (o *PoolResponse) SetServiceID(v string)`

SetServiceID sets ServiceID field to given value.

### HasServiceID

`func (o *PoolResponse) HasServiceID() bool`

HasServiceID returns a boolean if a field has been set.

### GetVersion

`func (o *PoolResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PoolResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PoolResponse) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PoolResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetName

`func (o *PoolResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PoolResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PoolResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PoolResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetShield

`func (o *PoolResponse) GetShield() string`

GetShield returns the Shield field if non-nil, zero value otherwise.

### GetShieldOk

`func (o *PoolResponse) GetShieldOk() (*string, bool)`

GetShieldOk returns a tuple with the Shield field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShield

`func (o *PoolResponse) SetShield(v string)`

SetShield sets Shield field to given value.

### HasShield

`func (o *PoolResponse) HasShield() bool`

HasShield returns a boolean if a field has been set.

### SetShieldNil

`func (o *PoolResponse) SetShieldNil(b bool)`

 SetShieldNil sets the value for Shield to be an explicit nil

### UnsetShield
`func (o *PoolResponse) UnsetShield()`

UnsetShield ensures that no value is present for Shield, not even an explicit nil
### GetRequestCondition

`func (o *PoolResponse) GetRequestCondition() string`

GetRequestCondition returns the RequestCondition field if non-nil, zero value otherwise.

### GetRequestConditionOk

`func (o *PoolResponse) GetRequestConditionOk() (*string, bool)`

GetRequestConditionOk returns a tuple with the RequestCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCondition

`func (o *PoolResponse) SetRequestCondition(v string)`

SetRequestCondition sets RequestCondition field to given value.

### HasRequestCondition

`func (o *PoolResponse) HasRequestCondition() bool`

HasRequestCondition returns a boolean if a field has been set.

### SetRequestConditionNil

`func (o *PoolResponse) SetRequestConditionNil(b bool)`

 SetRequestConditionNil sets the value for RequestCondition to be an explicit nil

### UnsetRequestCondition
`func (o *PoolResponse) UnsetRequestCondition()`

UnsetRequestCondition ensures that no value is present for RequestCondition, not even an explicit nil
### GetTLSCiphers

`func (o *PoolResponse) GetTLSCiphers() string`

GetTLSCiphers returns the TLSCiphers field if non-nil, zero value otherwise.

### GetTLSCiphersOk

`func (o *PoolResponse) GetTLSCiphersOk() (*string, bool)`

GetTLSCiphersOk returns a tuple with the TLSCiphers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTLSCiphers

`func (o *PoolResponse) SetTLSCiphers(v string)`

SetTLSCiphers sets TLSCiphers field to given value.

### HasTLSCiphers

`func (o *PoolResponse) HasTLSCiphers() bool`

HasTLSCiphers returns a boolean if a field has been set.

### SetTLSCiphersNil

`func (o *PoolResponse) SetTLSCiphersNil(b bool)`

 SetTLSCiphersNil sets the value for TLSCiphers to be an explicit nil

### UnsetTLSCiphers
`func (o *PoolResponse) UnsetTLSCiphers()`

UnsetTLSCiphers ensures that no value is present for TLSCiphers, not even an explicit nil
### GetTLSSniHostname

`func (o *PoolResponse) GetTLSSniHostname() string`

GetTLSSniHostname returns the TLSSniHostname field if non-nil, zero value otherwise.

### GetTLSSniHostnameOk

`func (o *PoolResponse) GetTLSSniHostnameOk() (*string, bool)`

GetTLSSniHostnameOk returns a tuple with the TLSSniHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTLSSniHostname

`func (o *PoolResponse) SetTLSSniHostname(v string)`

SetTLSSniHostname sets TLSSniHostname field to given value.

### HasTLSSniHostname

`func (o *PoolResponse) HasTLSSniHostname() bool`

HasTLSSniHostname returns a boolean if a field has been set.

### SetTLSSniHostnameNil

`func (o *PoolResponse) SetTLSSniHostnameNil(b bool)`

 SetTLSSniHostnameNil sets the value for TLSSniHostname to be an explicit nil

### UnsetTLSSniHostname
`func (o *PoolResponse) UnsetTLSSniHostname()`

UnsetTLSSniHostname ensures that no value is present for TLSSniHostname, not even an explicit nil
### GetMinTLSVersion

`func (o *PoolResponse) GetMinTLSVersion() int32`

GetMinTLSVersion returns the MinTLSVersion field if non-nil, zero value otherwise.

### GetMinTLSVersionOk

`func (o *PoolResponse) GetMinTLSVersionOk() (*int32, bool)`

GetMinTLSVersionOk returns a tuple with the MinTLSVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinTLSVersion

`func (o *PoolResponse) SetMinTLSVersion(v int32)`

SetMinTLSVersion sets MinTLSVersion field to given value.

### HasMinTLSVersion

`func (o *PoolResponse) HasMinTLSVersion() bool`

HasMinTLSVersion returns a boolean if a field has been set.

### SetMinTLSVersionNil

`func (o *PoolResponse) SetMinTLSVersionNil(b bool)`

 SetMinTLSVersionNil sets the value for MinTLSVersion to be an explicit nil

### UnsetMinTLSVersion
`func (o *PoolResponse) UnsetMinTLSVersion()`

UnsetMinTLSVersion ensures that no value is present for MinTLSVersion, not even an explicit nil
### GetMaxTLSVersion

`func (o *PoolResponse) GetMaxTLSVersion() int32`

GetMaxTLSVersion returns the MaxTLSVersion field if non-nil, zero value otherwise.

### GetMaxTLSVersionOk

`func (o *PoolResponse) GetMaxTLSVersionOk() (*int32, bool)`

GetMaxTLSVersionOk returns a tuple with the MaxTLSVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTLSVersion

`func (o *PoolResponse) SetMaxTLSVersion(v int32)`

SetMaxTLSVersion sets MaxTLSVersion field to given value.

### HasMaxTLSVersion

`func (o *PoolResponse) HasMaxTLSVersion() bool`

HasMaxTLSVersion returns a boolean if a field has been set.

### SetMaxTLSVersionNil

`func (o *PoolResponse) SetMaxTLSVersionNil(b bool)`

 SetMaxTLSVersionNil sets the value for MaxTLSVersion to be an explicit nil

### UnsetMaxTLSVersion
`func (o *PoolResponse) UnsetMaxTLSVersion()`

UnsetMaxTLSVersion ensures that no value is present for MaxTLSVersion, not even an explicit nil
### GetHealthcheck

`func (o *PoolResponse) GetHealthcheck() string`

GetHealthcheck returns the Healthcheck field if non-nil, zero value otherwise.

### GetHealthcheckOk

`func (o *PoolResponse) GetHealthcheckOk() (*string, bool)`

GetHealthcheckOk returns a tuple with the Healthcheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthcheck

`func (o *PoolResponse) SetHealthcheck(v string)`

SetHealthcheck sets Healthcheck field to given value.

### HasHealthcheck

`func (o *PoolResponse) HasHealthcheck() bool`

HasHealthcheck returns a boolean if a field has been set.

### SetHealthcheckNil

`func (o *PoolResponse) SetHealthcheckNil(b bool)`

 SetHealthcheckNil sets the value for Healthcheck to be an explicit nil

### UnsetHealthcheck
`func (o *PoolResponse) UnsetHealthcheck()`

UnsetHealthcheck ensures that no value is present for Healthcheck, not even an explicit nil
### GetComment

`func (o *PoolResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *PoolResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *PoolResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *PoolResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *PoolResponse) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *PoolResponse) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetType

`func (o *PoolResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PoolResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PoolResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PoolResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOverrideHost

`func (o *PoolResponse) GetOverrideHost() string`

GetOverrideHost returns the OverrideHost field if non-nil, zero value otherwise.

### GetOverrideHostOk

`func (o *PoolResponse) GetOverrideHostOk() (*string, bool)`

GetOverrideHostOk returns a tuple with the OverrideHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideHost

`func (o *PoolResponse) SetOverrideHost(v string)`

SetOverrideHost sets OverrideHost field to given value.

### HasOverrideHost

`func (o *PoolResponse) HasOverrideHost() bool`

HasOverrideHost returns a boolean if a field has been set.

### SetOverrideHostNil

`func (o *PoolResponse) SetOverrideHostNil(b bool)`

 SetOverrideHostNil sets the value for OverrideHost to be an explicit nil

### UnsetOverrideHost
`func (o *PoolResponse) UnsetOverrideHost()`

UnsetOverrideHost ensures that no value is present for OverrideHost, not even an explicit nil
### GetBetweenBytesTimeout

`func (o *PoolResponse) GetBetweenBytesTimeout() string`

GetBetweenBytesTimeout returns the BetweenBytesTimeout field if non-nil, zero value otherwise.

### GetBetweenBytesTimeoutOk

`func (o *PoolResponse) GetBetweenBytesTimeoutOk() (*string, bool)`

GetBetweenBytesTimeoutOk returns a tuple with the BetweenBytesTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBetweenBytesTimeout

`func (o *PoolResponse) SetBetweenBytesTimeout(v string)`

SetBetweenBytesTimeout sets BetweenBytesTimeout field to given value.

### HasBetweenBytesTimeout

`func (o *PoolResponse) HasBetweenBytesTimeout() bool`

HasBetweenBytesTimeout returns a boolean if a field has been set.

### GetConnectTimeout

`func (o *PoolResponse) GetConnectTimeout() string`

GetConnectTimeout returns the ConnectTimeout field if non-nil, zero value otherwise.

### GetConnectTimeoutOk

`func (o *PoolResponse) GetConnectTimeoutOk() (*string, bool)`

GetConnectTimeoutOk returns a tuple with the ConnectTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectTimeout

`func (o *PoolResponse) SetConnectTimeout(v string)`

SetConnectTimeout sets ConnectTimeout field to given value.

### HasConnectTimeout

`func (o *PoolResponse) HasConnectTimeout() bool`

HasConnectTimeout returns a boolean if a field has been set.

### GetFirstByteTimeout

`func (o *PoolResponse) GetFirstByteTimeout() string`

GetFirstByteTimeout returns the FirstByteTimeout field if non-nil, zero value otherwise.

### GetFirstByteTimeoutOk

`func (o *PoolResponse) GetFirstByteTimeoutOk() (*string, bool)`

GetFirstByteTimeoutOk returns a tuple with the FirstByteTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstByteTimeout

`func (o *PoolResponse) SetFirstByteTimeout(v string)`

SetFirstByteTimeout sets FirstByteTimeout field to given value.

### HasFirstByteTimeout

`func (o *PoolResponse) HasFirstByteTimeout() bool`

HasFirstByteTimeout returns a boolean if a field has been set.

### GetMaxConnDefault

`func (o *PoolResponse) GetMaxConnDefault() string`

GetMaxConnDefault returns the MaxConnDefault field if non-nil, zero value otherwise.

### GetMaxConnDefaultOk

`func (o *PoolResponse) GetMaxConnDefaultOk() (*string, bool)`

GetMaxConnDefaultOk returns a tuple with the MaxConnDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConnDefault

`func (o *PoolResponse) SetMaxConnDefault(v string)`

SetMaxConnDefault sets MaxConnDefault field to given value.

### HasMaxConnDefault

`func (o *PoolResponse) HasMaxConnDefault() bool`

HasMaxConnDefault returns a boolean if a field has been set.

### GetTLSCheckCert

`func (o *PoolResponse) GetTLSCheckCert() string`

GetTLSCheckCert returns the TLSCheckCert field if non-nil, zero value otherwise.

### GetTLSCheckCertOk

`func (o *PoolResponse) GetTLSCheckCertOk() (*string, bool)`

GetTLSCheckCertOk returns a tuple with the TLSCheckCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTLSCheckCert

`func (o *PoolResponse) SetTLSCheckCert(v string)`

SetTLSCheckCert sets TLSCheckCert field to given value.

### HasTLSCheckCert

`func (o *PoolResponse) HasTLSCheckCert() bool`

HasTLSCheckCert returns a boolean if a field has been set.

### SetTLSCheckCertNil

`func (o *PoolResponse) SetTLSCheckCertNil(b bool)`

 SetTLSCheckCertNil sets the value for TLSCheckCert to be an explicit nil

### UnsetTLSCheckCert
`func (o *PoolResponse) UnsetTLSCheckCert()`

UnsetTLSCheckCert ensures that no value is present for TLSCheckCert, not even an explicit nil
### GetID

`func (o *PoolResponse) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *PoolResponse) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *PoolResponse) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *PoolResponse) HasID() bool`

HasID returns a boolean if a field has been set.

### GetQuorum

`func (o *PoolResponse) GetQuorum() string`

GetQuorum returns the Quorum field if non-nil, zero value otherwise.

### GetQuorumOk

`func (o *PoolResponse) GetQuorumOk() (*string, bool)`

GetQuorumOk returns a tuple with the Quorum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuorum

`func (o *PoolResponse) SetQuorum(v string)`

SetQuorum sets Quorum field to given value.

### HasQuorum

`func (o *PoolResponse) HasQuorum() bool`

HasQuorum returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
