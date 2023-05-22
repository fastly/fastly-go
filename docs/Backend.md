# Backend

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | A hostname, IPv4, or IPv6 address for the backend. This is the preferred way to specify the location of your backend. | [optional] 
**AutoLoadbalance** | Pointer to **bool** | Whether or not this backend should be automatically load balanced. If true, all backends with this setting that don&#39;t have a `request_condition` will be selected based on their `weight`. | [optional] 
**BetweenBytesTimeout** | Pointer to **int32** | Maximum duration in milliseconds that Fastly will wait while receiving no data on a download from a backend. If exceeded, the response received so far will be considered complete and the fetch will end. May be set at runtime using `bereq.between_bytes_timeout`. | [optional] 
**ClientCert** | Pointer to **NullableString** | Unused. | [optional] 
**Comment** | Pointer to **NullableString** | A freeform descriptive note. | [optional] 
**ConnectTimeout** | Pointer to **int32** | Maximum duration in milliseconds to wait for a connection to this backend to be established. If exceeded, the connection is aborted and a synthethic `503` response will be presented instead. May be set at runtime using `bereq.connect_timeout`. | [optional] 
**FirstByteTimeout** | Pointer to **int32** | Maximum duration in milliseconds to wait for the server response to begin after a TCP connection is established and the request has been sent. If exceeded, the connection is aborted and a synthethic `503` response will be presented instead. May be set at runtime using `bereq.first_byte_timeout`. | [optional] 
**Healthcheck** | Pointer to **NullableString** | The name of the healthcheck to use with this backend. | [optional] 
**Hostname** | Pointer to **NullableString** | The hostname of the backend. May be used as an alternative to `address` to set the backend location. | [optional] 
**Ipv4** | Pointer to **NullableString** | IPv4 address of the backend. May be used as an alternative to `address` to set the backend location. | [optional] 
**Ipv6** | Pointer to **NullableString** | IPv6 address of the backend. May be used as an alternative to `address` to set the backend location. | [optional] 
**KeepaliveTime** | Pointer to **NullableInt32** | How long in seconds to keep a persistent connection to the backend between requests. | [optional] 
**MaxConn** | Pointer to **int32** | Maximum number of concurrent connections this backend will accept. | [optional] 
**MaxTLSVersion** | Pointer to **NullableString** | Maximum allowed TLS version on SSL connections to this backend. If your backend server is not able to negotiate a connection meeting this constraint, a synthetic `503` error response will be generated. | [optional] 
**MinTLSVersion** | Pointer to **NullableString** | Minimum allowed TLS version on SSL connections to this backend. If your backend server is not able to negotiate a connection meeting this constraint, a synthetic `503` error response will be generated. | [optional] 
**Name** | Pointer to **string** | The name of the backend. | [optional] 
**OverrideHost** | Pointer to **NullableString** | If set, will replace the client-supplied HTTP `Host` header on connections to this backend. Applied after VCL has been processed, so this setting will take precedence over changing `bereq.http.Host` in VCL. | [optional] 
**Port** | Pointer to **int32** | Port on which the backend server is listening for connections from Fastly. Setting `port` to 80 or 443 will also set `use_ssl` automatically (to false and true respectively), unless explicitly overridden by setting `use_ssl` in the same request. | [optional] 
**RequestCondition** | Pointer to **string** | Name of a Condition, which if satisfied, will select this backend during a request. If set, will override any `auto_loadbalance` setting. By default, the first backend added to a service is selected for all requests. | [optional] 
**Shield** | Pointer to **NullableString** | Identifier of the POP to use as a [shield](https://docs.fastly.com/en/guides/shielding). | [optional] 
**SslCaCert** | Pointer to **NullableString** | CA certificate attached to origin. | [optional] 
**SslCertHostname** | Pointer to **NullableString** | Overrides `ssl_hostname`, but only for cert verification. Does not affect SNI at all. | [optional] 
**SslCheckCert** | Pointer to **NullableBool** | Be strict on checking SSL certs. | [optional] [default to true]
**SslCiphers** | Pointer to **NullableString** | List of [OpenSSL ciphers](https://www.openssl.org/docs/manmaster/man1/ciphers.html) to support for connections to this origin. If your backend server is not able to negotiate a connection meeting this constraint, a synthetic `503` error response will be generated. | [optional] 
**SslClientCert** | Pointer to **NullableString** | Client certificate attached to origin. | [optional] 
**SslClientKey** | Pointer to **NullableString** | Client key attached to origin. | [optional] 
**SslHostname** | Pointer to **NullableString** | Use `ssl_cert_hostname` and `ssl_sni_hostname` to configure certificate validation. | [optional] 
**SslSniHostname** | Pointer to **NullableString** | Overrides `ssl_hostname`, but only for SNI in the handshake. Does not affect cert validation at all. | [optional] 
**UseSsl** | Pointer to **bool** | Whether or not to require TLS for connections to this backend. | [optional] 
**Weight** | Pointer to **int32** | Weight used to load balance this backend against others. May be any positive integer. If `auto_loadbalance` is true, the chance of this backend being selected is equal to its own weight over the sum of all weights for backends that have `auto_loadbalance` set to true. | [optional] 

## Methods

### NewBackend

`func NewBackend() *Backend`

NewBackend instantiates a new Backend object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackendWithDefaults

`func NewBackendWithDefaults() *Backend`

NewBackendWithDefaults instantiates a new Backend object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *Backend) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Backend) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Backend) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *Backend) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAutoLoadbalance

`func (o *Backend) GetAutoLoadbalance() bool`

GetAutoLoadbalance returns the AutoLoadbalance field if non-nil, zero value otherwise.

### GetAutoLoadbalanceOk

`func (o *Backend) GetAutoLoadbalanceOk() (*bool, bool)`

GetAutoLoadbalanceOk returns a tuple with the AutoLoadbalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoLoadbalance

`func (o *Backend) SetAutoLoadbalance(v bool)`

SetAutoLoadbalance sets AutoLoadbalance field to given value.

### HasAutoLoadbalance

`func (o *Backend) HasAutoLoadbalance() bool`

HasAutoLoadbalance returns a boolean if a field has been set.

### GetBetweenBytesTimeout

`func (o *Backend) GetBetweenBytesTimeout() int32`

GetBetweenBytesTimeout returns the BetweenBytesTimeout field if non-nil, zero value otherwise.

### GetBetweenBytesTimeoutOk

`func (o *Backend) GetBetweenBytesTimeoutOk() (*int32, bool)`

GetBetweenBytesTimeoutOk returns a tuple with the BetweenBytesTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBetweenBytesTimeout

`func (o *Backend) SetBetweenBytesTimeout(v int32)`

SetBetweenBytesTimeout sets BetweenBytesTimeout field to given value.

### HasBetweenBytesTimeout

`func (o *Backend) HasBetweenBytesTimeout() bool`

HasBetweenBytesTimeout returns a boolean if a field has been set.

### GetClientCert

`func (o *Backend) GetClientCert() string`

GetClientCert returns the ClientCert field if non-nil, zero value otherwise.

### GetClientCertOk

`func (o *Backend) GetClientCertOk() (*string, bool)`

GetClientCertOk returns a tuple with the ClientCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCert

`func (o *Backend) SetClientCert(v string)`

SetClientCert sets ClientCert field to given value.

### HasClientCert

`func (o *Backend) HasClientCert() bool`

HasClientCert returns a boolean if a field has been set.

### SetClientCertNil

`func (o *Backend) SetClientCertNil(b bool)`

 SetClientCertNil sets the value for ClientCert to be an explicit nil

### UnsetClientCert
`func (o *Backend) UnsetClientCert()`

UnsetClientCert ensures that no value is present for ClientCert, not even an explicit nil
### GetComment

`func (o *Backend) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Backend) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Backend) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Backend) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *Backend) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *Backend) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetConnectTimeout

`func (o *Backend) GetConnectTimeout() int32`

GetConnectTimeout returns the ConnectTimeout field if non-nil, zero value otherwise.

### GetConnectTimeoutOk

`func (o *Backend) GetConnectTimeoutOk() (*int32, bool)`

GetConnectTimeoutOk returns a tuple with the ConnectTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectTimeout

`func (o *Backend) SetConnectTimeout(v int32)`

SetConnectTimeout sets ConnectTimeout field to given value.

### HasConnectTimeout

`func (o *Backend) HasConnectTimeout() bool`

HasConnectTimeout returns a boolean if a field has been set.

### GetFirstByteTimeout

`func (o *Backend) GetFirstByteTimeout() int32`

GetFirstByteTimeout returns the FirstByteTimeout field if non-nil, zero value otherwise.

### GetFirstByteTimeoutOk

`func (o *Backend) GetFirstByteTimeoutOk() (*int32, bool)`

GetFirstByteTimeoutOk returns a tuple with the FirstByteTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstByteTimeout

`func (o *Backend) SetFirstByteTimeout(v int32)`

SetFirstByteTimeout sets FirstByteTimeout field to given value.

### HasFirstByteTimeout

`func (o *Backend) HasFirstByteTimeout() bool`

HasFirstByteTimeout returns a boolean if a field has been set.

### GetHealthcheck

`func (o *Backend) GetHealthcheck() string`

GetHealthcheck returns the Healthcheck field if non-nil, zero value otherwise.

### GetHealthcheckOk

`func (o *Backend) GetHealthcheckOk() (*string, bool)`

GetHealthcheckOk returns a tuple with the Healthcheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthcheck

`func (o *Backend) SetHealthcheck(v string)`

SetHealthcheck sets Healthcheck field to given value.

### HasHealthcheck

`func (o *Backend) HasHealthcheck() bool`

HasHealthcheck returns a boolean if a field has been set.

### SetHealthcheckNil

`func (o *Backend) SetHealthcheckNil(b bool)`

 SetHealthcheckNil sets the value for Healthcheck to be an explicit nil

### UnsetHealthcheck
`func (o *Backend) UnsetHealthcheck()`

UnsetHealthcheck ensures that no value is present for Healthcheck, not even an explicit nil
### GetHostname

`func (o *Backend) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *Backend) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *Backend) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *Backend) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### SetHostnameNil

`func (o *Backend) SetHostnameNil(b bool)`

 SetHostnameNil sets the value for Hostname to be an explicit nil

### UnsetHostname
`func (o *Backend) UnsetHostname()`

UnsetHostname ensures that no value is present for Hostname, not even an explicit nil
### GetIpv4

`func (o *Backend) GetIpv4() string`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *Backend) GetIpv4Ok() (*string, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *Backend) SetIpv4(v string)`

SetIpv4 sets Ipv4 field to given value.

### HasIpv4

`func (o *Backend) HasIpv4() bool`

HasIpv4 returns a boolean if a field has been set.

### SetIpv4Nil

`func (o *Backend) SetIpv4Nil(b bool)`

 SetIpv4Nil sets the value for Ipv4 to be an explicit nil

### UnsetIpv4
`func (o *Backend) UnsetIpv4()`

UnsetIpv4 ensures that no value is present for Ipv4, not even an explicit nil
### GetIpv6

`func (o *Backend) GetIpv6() string`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *Backend) GetIpv6Ok() (*string, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *Backend) SetIpv6(v string)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *Backend) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.

### SetIpv6Nil

`func (o *Backend) SetIpv6Nil(b bool)`

 SetIpv6Nil sets the value for Ipv6 to be an explicit nil

### UnsetIpv6
`func (o *Backend) UnsetIpv6()`

UnsetIpv6 ensures that no value is present for Ipv6, not even an explicit nil
### GetKeepaliveTime

`func (o *Backend) GetKeepaliveTime() int32`

GetKeepaliveTime returns the KeepaliveTime field if non-nil, zero value otherwise.

### GetKeepaliveTimeOk

`func (o *Backend) GetKeepaliveTimeOk() (*int32, bool)`

GetKeepaliveTimeOk returns a tuple with the KeepaliveTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepaliveTime

`func (o *Backend) SetKeepaliveTime(v int32)`

SetKeepaliveTime sets KeepaliveTime field to given value.

### HasKeepaliveTime

`func (o *Backend) HasKeepaliveTime() bool`

HasKeepaliveTime returns a boolean if a field has been set.

### SetKeepaliveTimeNil

`func (o *Backend) SetKeepaliveTimeNil(b bool)`

 SetKeepaliveTimeNil sets the value for KeepaliveTime to be an explicit nil

### UnsetKeepaliveTime
`func (o *Backend) UnsetKeepaliveTime()`

UnsetKeepaliveTime ensures that no value is present for KeepaliveTime, not even an explicit nil
### GetMaxConn

`func (o *Backend) GetMaxConn() int32`

GetMaxConn returns the MaxConn field if non-nil, zero value otherwise.

### GetMaxConnOk

`func (o *Backend) GetMaxConnOk() (*int32, bool)`

GetMaxConnOk returns a tuple with the MaxConn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConn

`func (o *Backend) SetMaxConn(v int32)`

SetMaxConn sets MaxConn field to given value.

### HasMaxConn

`func (o *Backend) HasMaxConn() bool`

HasMaxConn returns a boolean if a field has been set.

### GetMaxTLSVersion

`func (o *Backend) GetMaxTLSVersion() string`

GetMaxTLSVersion returns the MaxTLSVersion field if non-nil, zero value otherwise.

### GetMaxTLSVersionOk

`func (o *Backend) GetMaxTLSVersionOk() (*string, bool)`

GetMaxTLSVersionOk returns a tuple with the MaxTLSVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTLSVersion

`func (o *Backend) SetMaxTLSVersion(v string)`

SetMaxTLSVersion sets MaxTLSVersion field to given value.

### HasMaxTLSVersion

`func (o *Backend) HasMaxTLSVersion() bool`

HasMaxTLSVersion returns a boolean if a field has been set.

### SetMaxTLSVersionNil

`func (o *Backend) SetMaxTLSVersionNil(b bool)`

 SetMaxTLSVersionNil sets the value for MaxTLSVersion to be an explicit nil

### UnsetMaxTLSVersion
`func (o *Backend) UnsetMaxTLSVersion()`

UnsetMaxTLSVersion ensures that no value is present for MaxTLSVersion, not even an explicit nil
### GetMinTLSVersion

`func (o *Backend) GetMinTLSVersion() string`

GetMinTLSVersion returns the MinTLSVersion field if non-nil, zero value otherwise.

### GetMinTLSVersionOk

`func (o *Backend) GetMinTLSVersionOk() (*string, bool)`

GetMinTLSVersionOk returns a tuple with the MinTLSVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinTLSVersion

`func (o *Backend) SetMinTLSVersion(v string)`

SetMinTLSVersion sets MinTLSVersion field to given value.

### HasMinTLSVersion

`func (o *Backend) HasMinTLSVersion() bool`

HasMinTLSVersion returns a boolean if a field has been set.

### SetMinTLSVersionNil

`func (o *Backend) SetMinTLSVersionNil(b bool)`

 SetMinTLSVersionNil sets the value for MinTLSVersion to be an explicit nil

### UnsetMinTLSVersion
`func (o *Backend) UnsetMinTLSVersion()`

UnsetMinTLSVersion ensures that no value is present for MinTLSVersion, not even an explicit nil
### GetName

`func (o *Backend) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Backend) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Backend) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Backend) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOverrideHost

`func (o *Backend) GetOverrideHost() string`

GetOverrideHost returns the OverrideHost field if non-nil, zero value otherwise.

### GetOverrideHostOk

`func (o *Backend) GetOverrideHostOk() (*string, bool)`

GetOverrideHostOk returns a tuple with the OverrideHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideHost

`func (o *Backend) SetOverrideHost(v string)`

SetOverrideHost sets OverrideHost field to given value.

### HasOverrideHost

`func (o *Backend) HasOverrideHost() bool`

HasOverrideHost returns a boolean if a field has been set.

### SetOverrideHostNil

`func (o *Backend) SetOverrideHostNil(b bool)`

 SetOverrideHostNil sets the value for OverrideHost to be an explicit nil

### UnsetOverrideHost
`func (o *Backend) UnsetOverrideHost()`

UnsetOverrideHost ensures that no value is present for OverrideHost, not even an explicit nil
### GetPort

`func (o *Backend) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *Backend) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *Backend) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *Backend) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetRequestCondition

`func (o *Backend) GetRequestCondition() string`

GetRequestCondition returns the RequestCondition field if non-nil, zero value otherwise.

### GetRequestConditionOk

`func (o *Backend) GetRequestConditionOk() (*string, bool)`

GetRequestConditionOk returns a tuple with the RequestCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCondition

`func (o *Backend) SetRequestCondition(v string)`

SetRequestCondition sets RequestCondition field to given value.

### HasRequestCondition

`func (o *Backend) HasRequestCondition() bool`

HasRequestCondition returns a boolean if a field has been set.

### GetShield

`func (o *Backend) GetShield() string`

GetShield returns the Shield field if non-nil, zero value otherwise.

### GetShieldOk

`func (o *Backend) GetShieldOk() (*string, bool)`

GetShieldOk returns a tuple with the Shield field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShield

`func (o *Backend) SetShield(v string)`

SetShield sets Shield field to given value.

### HasShield

`func (o *Backend) HasShield() bool`

HasShield returns a boolean if a field has been set.

### SetShieldNil

`func (o *Backend) SetShieldNil(b bool)`

 SetShieldNil sets the value for Shield to be an explicit nil

### UnsetShield
`func (o *Backend) UnsetShield()`

UnsetShield ensures that no value is present for Shield, not even an explicit nil
### GetSslCaCert

`func (o *Backend) GetSslCaCert() string`

GetSslCaCert returns the SslCaCert field if non-nil, zero value otherwise.

### GetSslCaCertOk

`func (o *Backend) GetSslCaCertOk() (*string, bool)`

GetSslCaCertOk returns a tuple with the SslCaCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCaCert

`func (o *Backend) SetSslCaCert(v string)`

SetSslCaCert sets SslCaCert field to given value.

### HasSslCaCert

`func (o *Backend) HasSslCaCert() bool`

HasSslCaCert returns a boolean if a field has been set.

### SetSslCaCertNil

`func (o *Backend) SetSslCaCertNil(b bool)`

 SetSslCaCertNil sets the value for SslCaCert to be an explicit nil

### UnsetSslCaCert
`func (o *Backend) UnsetSslCaCert()`

UnsetSslCaCert ensures that no value is present for SslCaCert, not even an explicit nil
### GetSslCertHostname

`func (o *Backend) GetSslCertHostname() string`

GetSslCertHostname returns the SslCertHostname field if non-nil, zero value otherwise.

### GetSslCertHostnameOk

`func (o *Backend) GetSslCertHostnameOk() (*string, bool)`

GetSslCertHostnameOk returns a tuple with the SslCertHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCertHostname

`func (o *Backend) SetSslCertHostname(v string)`

SetSslCertHostname sets SslCertHostname field to given value.

### HasSslCertHostname

`func (o *Backend) HasSslCertHostname() bool`

HasSslCertHostname returns a boolean if a field has been set.

### SetSslCertHostnameNil

`func (o *Backend) SetSslCertHostnameNil(b bool)`

 SetSslCertHostnameNil sets the value for SslCertHostname to be an explicit nil

### UnsetSslCertHostname
`func (o *Backend) UnsetSslCertHostname()`

UnsetSslCertHostname ensures that no value is present for SslCertHostname, not even an explicit nil
### GetSslCheckCert

`func (o *Backend) GetSslCheckCert() bool`

GetSslCheckCert returns the SslCheckCert field if non-nil, zero value otherwise.

### GetSslCheckCertOk

`func (o *Backend) GetSslCheckCertOk() (*bool, bool)`

GetSslCheckCertOk returns a tuple with the SslCheckCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCheckCert

`func (o *Backend) SetSslCheckCert(v bool)`

SetSslCheckCert sets SslCheckCert field to given value.

### HasSslCheckCert

`func (o *Backend) HasSslCheckCert() bool`

HasSslCheckCert returns a boolean if a field has been set.

### SetSslCheckCertNil

`func (o *Backend) SetSslCheckCertNil(b bool)`

 SetSslCheckCertNil sets the value for SslCheckCert to be an explicit nil

### UnsetSslCheckCert
`func (o *Backend) UnsetSslCheckCert()`

UnsetSslCheckCert ensures that no value is present for SslCheckCert, not even an explicit nil
### GetSslCiphers

`func (o *Backend) GetSslCiphers() string`

GetSslCiphers returns the SslCiphers field if non-nil, zero value otherwise.

### GetSslCiphersOk

`func (o *Backend) GetSslCiphersOk() (*string, bool)`

GetSslCiphersOk returns a tuple with the SslCiphers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCiphers

`func (o *Backend) SetSslCiphers(v string)`

SetSslCiphers sets SslCiphers field to given value.

### HasSslCiphers

`func (o *Backend) HasSslCiphers() bool`

HasSslCiphers returns a boolean if a field has been set.

### SetSslCiphersNil

`func (o *Backend) SetSslCiphersNil(b bool)`

 SetSslCiphersNil sets the value for SslCiphers to be an explicit nil

### UnsetSslCiphers
`func (o *Backend) UnsetSslCiphers()`

UnsetSslCiphers ensures that no value is present for SslCiphers, not even an explicit nil
### GetSslClientCert

`func (o *Backend) GetSslClientCert() string`

GetSslClientCert returns the SslClientCert field if non-nil, zero value otherwise.

### GetSslClientCertOk

`func (o *Backend) GetSslClientCertOk() (*string, bool)`

GetSslClientCertOk returns a tuple with the SslClientCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslClientCert

`func (o *Backend) SetSslClientCert(v string)`

SetSslClientCert sets SslClientCert field to given value.

### HasSslClientCert

`func (o *Backend) HasSslClientCert() bool`

HasSslClientCert returns a boolean if a field has been set.

### SetSslClientCertNil

`func (o *Backend) SetSslClientCertNil(b bool)`

 SetSslClientCertNil sets the value for SslClientCert to be an explicit nil

### UnsetSslClientCert
`func (o *Backend) UnsetSslClientCert()`

UnsetSslClientCert ensures that no value is present for SslClientCert, not even an explicit nil
### GetSslClientKey

`func (o *Backend) GetSslClientKey() string`

GetSslClientKey returns the SslClientKey field if non-nil, zero value otherwise.

### GetSslClientKeyOk

`func (o *Backend) GetSslClientKeyOk() (*string, bool)`

GetSslClientKeyOk returns a tuple with the SslClientKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslClientKey

`func (o *Backend) SetSslClientKey(v string)`

SetSslClientKey sets SslClientKey field to given value.

### HasSslClientKey

`func (o *Backend) HasSslClientKey() bool`

HasSslClientKey returns a boolean if a field has been set.

### SetSslClientKeyNil

`func (o *Backend) SetSslClientKeyNil(b bool)`

 SetSslClientKeyNil sets the value for SslClientKey to be an explicit nil

### UnsetSslClientKey
`func (o *Backend) UnsetSslClientKey()`

UnsetSslClientKey ensures that no value is present for SslClientKey, not even an explicit nil
### GetSslHostname

`func (o *Backend) GetSslHostname() string`

GetSslHostname returns the SslHostname field if non-nil, zero value otherwise.

### GetSslHostnameOk

`func (o *Backend) GetSslHostnameOk() (*string, bool)`

GetSslHostnameOk returns a tuple with the SslHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslHostname

`func (o *Backend) SetSslHostname(v string)`

SetSslHostname sets SslHostname field to given value.

### HasSslHostname

`func (o *Backend) HasSslHostname() bool`

HasSslHostname returns a boolean if a field has been set.

### SetSslHostnameNil

`func (o *Backend) SetSslHostnameNil(b bool)`

 SetSslHostnameNil sets the value for SslHostname to be an explicit nil

### UnsetSslHostname
`func (o *Backend) UnsetSslHostname()`

UnsetSslHostname ensures that no value is present for SslHostname, not even an explicit nil
### GetSslSniHostname

`func (o *Backend) GetSslSniHostname() string`

GetSslSniHostname returns the SslSniHostname field if non-nil, zero value otherwise.

### GetSslSniHostnameOk

`func (o *Backend) GetSslSniHostnameOk() (*string, bool)`

GetSslSniHostnameOk returns a tuple with the SslSniHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslSniHostname

`func (o *Backend) SetSslSniHostname(v string)`

SetSslSniHostname sets SslSniHostname field to given value.

### HasSslSniHostname

`func (o *Backend) HasSslSniHostname() bool`

HasSslSniHostname returns a boolean if a field has been set.

### SetSslSniHostnameNil

`func (o *Backend) SetSslSniHostnameNil(b bool)`

 SetSslSniHostnameNil sets the value for SslSniHostname to be an explicit nil

### UnsetSslSniHostname
`func (o *Backend) UnsetSslSniHostname()`

UnsetSslSniHostname ensures that no value is present for SslSniHostname, not even an explicit nil
### GetUseSsl

`func (o *Backend) GetUseSsl() bool`

GetUseSsl returns the UseSsl field if non-nil, zero value otherwise.

### GetUseSslOk

`func (o *Backend) GetUseSslOk() (*bool, bool)`

GetUseSslOk returns a tuple with the UseSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSsl

`func (o *Backend) SetUseSsl(v bool)`

SetUseSsl sets UseSsl field to given value.

### HasUseSsl

`func (o *Backend) HasUseSsl() bool`

HasUseSsl returns a boolean if a field has been set.

### GetWeight

`func (o *Backend) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *Backend) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *Backend) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *Backend) HasWeight() bool`

HasWeight returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
