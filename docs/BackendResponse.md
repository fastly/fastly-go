# BackendResponse

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
**KeepaliveTime** | Pointer to **NullableInt32** | How long to keep a persistent connection to the backend between requests. | [optional] 
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
**SslCheckCert** | Pointer to **bool** | Be strict on checking SSL certs. | [optional] [default to true]
**SslCiphers** | Pointer to **NullableString** | List of [OpenSSL ciphers](https://www.openssl.org/docs/manmaster/man1/ciphers.html) to support for connections to this origin. If your backend server is not able to negotiate a connection meeting this constraint, a synthetic `503` error response will be generated. | [optional] 
**SslClientCert** | Pointer to **NullableString** | Client certificate attached to origin. | [optional] 
**SslClientKey** | Pointer to **NullableString** | Client key attached to origin. | [optional] 
**SslHostname** | Pointer to **NullableString** | Use `ssl_cert_hostname` and `ssl_sni_hostname` to configure certificate validation. | [optional] 
**SslSniHostname** | Pointer to **NullableString** | Overrides `ssl_hostname`, but only for SNI in the handshake. Does not affect cert validation at all. | [optional] 
**UseSsl** | Pointer to **bool** | Whether or not to require TLS for connections to this backend. | [optional] 
**Weight** | Pointer to **int32** | Weight used to load balance this backend against others. May be any positive integer. If `auto_loadbalance` is true, the chance of this backend being selected is equal to its own weight over the sum of all weights for backends that have `auto_loadbalance` set to true. | [optional] 
**CreatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**DeletedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**UpdatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**ServiceID** | Pointer to **string** |  | [optional] [readonly] 
**Version** | Pointer to **int32** |  | [optional] [readonly] 
**Locked** | Pointer to **bool** | Indicates whether the version of the service this backend is attached to accepts edits. | [optional] [readonly] 

## Methods

### NewBackendResponse

`func NewBackendResponse() *BackendResponse`

NewBackendResponse instantiates a new BackendResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackendResponseWithDefaults

`func NewBackendResponseWithDefaults() *BackendResponse`

NewBackendResponseWithDefaults instantiates a new BackendResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *BackendResponse) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *BackendResponse) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *BackendResponse) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *BackendResponse) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAutoLoadbalance

`func (o *BackendResponse) GetAutoLoadbalance() bool`

GetAutoLoadbalance returns the AutoLoadbalance field if non-nil, zero value otherwise.

### GetAutoLoadbalanceOk

`func (o *BackendResponse) GetAutoLoadbalanceOk() (*bool, bool)`

GetAutoLoadbalanceOk returns a tuple with the AutoLoadbalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoLoadbalance

`func (o *BackendResponse) SetAutoLoadbalance(v bool)`

SetAutoLoadbalance sets AutoLoadbalance field to given value.

### HasAutoLoadbalance

`func (o *BackendResponse) HasAutoLoadbalance() bool`

HasAutoLoadbalance returns a boolean if a field has been set.

### GetBetweenBytesTimeout

`func (o *BackendResponse) GetBetweenBytesTimeout() int32`

GetBetweenBytesTimeout returns the BetweenBytesTimeout field if non-nil, zero value otherwise.

### GetBetweenBytesTimeoutOk

`func (o *BackendResponse) GetBetweenBytesTimeoutOk() (*int32, bool)`

GetBetweenBytesTimeoutOk returns a tuple with the BetweenBytesTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBetweenBytesTimeout

`func (o *BackendResponse) SetBetweenBytesTimeout(v int32)`

SetBetweenBytesTimeout sets BetweenBytesTimeout field to given value.

### HasBetweenBytesTimeout

`func (o *BackendResponse) HasBetweenBytesTimeout() bool`

HasBetweenBytesTimeout returns a boolean if a field has been set.

### GetClientCert

`func (o *BackendResponse) GetClientCert() string`

GetClientCert returns the ClientCert field if non-nil, zero value otherwise.

### GetClientCertOk

`func (o *BackendResponse) GetClientCertOk() (*string, bool)`

GetClientCertOk returns a tuple with the ClientCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCert

`func (o *BackendResponse) SetClientCert(v string)`

SetClientCert sets ClientCert field to given value.

### HasClientCert

`func (o *BackendResponse) HasClientCert() bool`

HasClientCert returns a boolean if a field has been set.

### SetClientCertNil

`func (o *BackendResponse) SetClientCertNil(b bool)`

 SetClientCertNil sets the value for ClientCert to be an explicit nil

### UnsetClientCert
`func (o *BackendResponse) UnsetClientCert()`

UnsetClientCert ensures that no value is present for ClientCert, not even an explicit nil
### GetComment

`func (o *BackendResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *BackendResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *BackendResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *BackendResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *BackendResponse) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *BackendResponse) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetConnectTimeout

`func (o *BackendResponse) GetConnectTimeout() int32`

GetConnectTimeout returns the ConnectTimeout field if non-nil, zero value otherwise.

### GetConnectTimeoutOk

`func (o *BackendResponse) GetConnectTimeoutOk() (*int32, bool)`

GetConnectTimeoutOk returns a tuple with the ConnectTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectTimeout

`func (o *BackendResponse) SetConnectTimeout(v int32)`

SetConnectTimeout sets ConnectTimeout field to given value.

### HasConnectTimeout

`func (o *BackendResponse) HasConnectTimeout() bool`

HasConnectTimeout returns a boolean if a field has been set.

### GetFirstByteTimeout

`func (o *BackendResponse) GetFirstByteTimeout() int32`

GetFirstByteTimeout returns the FirstByteTimeout field if non-nil, zero value otherwise.

### GetFirstByteTimeoutOk

`func (o *BackendResponse) GetFirstByteTimeoutOk() (*int32, bool)`

GetFirstByteTimeoutOk returns a tuple with the FirstByteTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstByteTimeout

`func (o *BackendResponse) SetFirstByteTimeout(v int32)`

SetFirstByteTimeout sets FirstByteTimeout field to given value.

### HasFirstByteTimeout

`func (o *BackendResponse) HasFirstByteTimeout() bool`

HasFirstByteTimeout returns a boolean if a field has been set.

### GetHealthcheck

`func (o *BackendResponse) GetHealthcheck() string`

GetHealthcheck returns the Healthcheck field if non-nil, zero value otherwise.

### GetHealthcheckOk

`func (o *BackendResponse) GetHealthcheckOk() (*string, bool)`

GetHealthcheckOk returns a tuple with the Healthcheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthcheck

`func (o *BackendResponse) SetHealthcheck(v string)`

SetHealthcheck sets Healthcheck field to given value.

### HasHealthcheck

`func (o *BackendResponse) HasHealthcheck() bool`

HasHealthcheck returns a boolean if a field has been set.

### SetHealthcheckNil

`func (o *BackendResponse) SetHealthcheckNil(b bool)`

 SetHealthcheckNil sets the value for Healthcheck to be an explicit nil

### UnsetHealthcheck
`func (o *BackendResponse) UnsetHealthcheck()`

UnsetHealthcheck ensures that no value is present for Healthcheck, not even an explicit nil
### GetHostname

`func (o *BackendResponse) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *BackendResponse) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *BackendResponse) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *BackendResponse) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### SetHostnameNil

`func (o *BackendResponse) SetHostnameNil(b bool)`

 SetHostnameNil sets the value for Hostname to be an explicit nil

### UnsetHostname
`func (o *BackendResponse) UnsetHostname()`

UnsetHostname ensures that no value is present for Hostname, not even an explicit nil
### GetIpv4

`func (o *BackendResponse) GetIpv4() string`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *BackendResponse) GetIpv4Ok() (*string, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *BackendResponse) SetIpv4(v string)`

SetIpv4 sets Ipv4 field to given value.

### HasIpv4

`func (o *BackendResponse) HasIpv4() bool`

HasIpv4 returns a boolean if a field has been set.

### SetIpv4Nil

`func (o *BackendResponse) SetIpv4Nil(b bool)`

 SetIpv4Nil sets the value for Ipv4 to be an explicit nil

### UnsetIpv4
`func (o *BackendResponse) UnsetIpv4()`

UnsetIpv4 ensures that no value is present for Ipv4, not even an explicit nil
### GetIpv6

`func (o *BackendResponse) GetIpv6() string`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *BackendResponse) GetIpv6Ok() (*string, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *BackendResponse) SetIpv6(v string)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *BackendResponse) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.

### SetIpv6Nil

`func (o *BackendResponse) SetIpv6Nil(b bool)`

 SetIpv6Nil sets the value for Ipv6 to be an explicit nil

### UnsetIpv6
`func (o *BackendResponse) UnsetIpv6()`

UnsetIpv6 ensures that no value is present for Ipv6, not even an explicit nil
### GetKeepaliveTime

`func (o *BackendResponse) GetKeepaliveTime() int32`

GetKeepaliveTime returns the KeepaliveTime field if non-nil, zero value otherwise.

### GetKeepaliveTimeOk

`func (o *BackendResponse) GetKeepaliveTimeOk() (*int32, bool)`

GetKeepaliveTimeOk returns a tuple with the KeepaliveTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepaliveTime

`func (o *BackendResponse) SetKeepaliveTime(v int32)`

SetKeepaliveTime sets KeepaliveTime field to given value.

### HasKeepaliveTime

`func (o *BackendResponse) HasKeepaliveTime() bool`

HasKeepaliveTime returns a boolean if a field has been set.

### SetKeepaliveTimeNil

`func (o *BackendResponse) SetKeepaliveTimeNil(b bool)`

 SetKeepaliveTimeNil sets the value for KeepaliveTime to be an explicit nil

### UnsetKeepaliveTime
`func (o *BackendResponse) UnsetKeepaliveTime()`

UnsetKeepaliveTime ensures that no value is present for KeepaliveTime, not even an explicit nil
### GetMaxConn

`func (o *BackendResponse) GetMaxConn() int32`

GetMaxConn returns the MaxConn field if non-nil, zero value otherwise.

### GetMaxConnOk

`func (o *BackendResponse) GetMaxConnOk() (*int32, bool)`

GetMaxConnOk returns a tuple with the MaxConn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConn

`func (o *BackendResponse) SetMaxConn(v int32)`

SetMaxConn sets MaxConn field to given value.

### HasMaxConn

`func (o *BackendResponse) HasMaxConn() bool`

HasMaxConn returns a boolean if a field has been set.

### GetMaxTLSVersion

`func (o *BackendResponse) GetMaxTLSVersion() string`

GetMaxTLSVersion returns the MaxTLSVersion field if non-nil, zero value otherwise.

### GetMaxTLSVersionOk

`func (o *BackendResponse) GetMaxTLSVersionOk() (*string, bool)`

GetMaxTLSVersionOk returns a tuple with the MaxTLSVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTLSVersion

`func (o *BackendResponse) SetMaxTLSVersion(v string)`

SetMaxTLSVersion sets MaxTLSVersion field to given value.

### HasMaxTLSVersion

`func (o *BackendResponse) HasMaxTLSVersion() bool`

HasMaxTLSVersion returns a boolean if a field has been set.

### SetMaxTLSVersionNil

`func (o *BackendResponse) SetMaxTLSVersionNil(b bool)`

 SetMaxTLSVersionNil sets the value for MaxTLSVersion to be an explicit nil

### UnsetMaxTLSVersion
`func (o *BackendResponse) UnsetMaxTLSVersion()`

UnsetMaxTLSVersion ensures that no value is present for MaxTLSVersion, not even an explicit nil
### GetMinTLSVersion

`func (o *BackendResponse) GetMinTLSVersion() string`

GetMinTLSVersion returns the MinTLSVersion field if non-nil, zero value otherwise.

### GetMinTLSVersionOk

`func (o *BackendResponse) GetMinTLSVersionOk() (*string, bool)`

GetMinTLSVersionOk returns a tuple with the MinTLSVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinTLSVersion

`func (o *BackendResponse) SetMinTLSVersion(v string)`

SetMinTLSVersion sets MinTLSVersion field to given value.

### HasMinTLSVersion

`func (o *BackendResponse) HasMinTLSVersion() bool`

HasMinTLSVersion returns a boolean if a field has been set.

### SetMinTLSVersionNil

`func (o *BackendResponse) SetMinTLSVersionNil(b bool)`

 SetMinTLSVersionNil sets the value for MinTLSVersion to be an explicit nil

### UnsetMinTLSVersion
`func (o *BackendResponse) UnsetMinTLSVersion()`

UnsetMinTLSVersion ensures that no value is present for MinTLSVersion, not even an explicit nil
### GetName

`func (o *BackendResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BackendResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BackendResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BackendResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOverrideHost

`func (o *BackendResponse) GetOverrideHost() string`

GetOverrideHost returns the OverrideHost field if non-nil, zero value otherwise.

### GetOverrideHostOk

`func (o *BackendResponse) GetOverrideHostOk() (*string, bool)`

GetOverrideHostOk returns a tuple with the OverrideHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideHost

`func (o *BackendResponse) SetOverrideHost(v string)`

SetOverrideHost sets OverrideHost field to given value.

### HasOverrideHost

`func (o *BackendResponse) HasOverrideHost() bool`

HasOverrideHost returns a boolean if a field has been set.

### SetOverrideHostNil

`func (o *BackendResponse) SetOverrideHostNil(b bool)`

 SetOverrideHostNil sets the value for OverrideHost to be an explicit nil

### UnsetOverrideHost
`func (o *BackendResponse) UnsetOverrideHost()`

UnsetOverrideHost ensures that no value is present for OverrideHost, not even an explicit nil
### GetPort

`func (o *BackendResponse) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *BackendResponse) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *BackendResponse) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *BackendResponse) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetRequestCondition

`func (o *BackendResponse) GetRequestCondition() string`

GetRequestCondition returns the RequestCondition field if non-nil, zero value otherwise.

### GetRequestConditionOk

`func (o *BackendResponse) GetRequestConditionOk() (*string, bool)`

GetRequestConditionOk returns a tuple with the RequestCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCondition

`func (o *BackendResponse) SetRequestCondition(v string)`

SetRequestCondition sets RequestCondition field to given value.

### HasRequestCondition

`func (o *BackendResponse) HasRequestCondition() bool`

HasRequestCondition returns a boolean if a field has been set.

### GetShield

`func (o *BackendResponse) GetShield() string`

GetShield returns the Shield field if non-nil, zero value otherwise.

### GetShieldOk

`func (o *BackendResponse) GetShieldOk() (*string, bool)`

GetShieldOk returns a tuple with the Shield field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShield

`func (o *BackendResponse) SetShield(v string)`

SetShield sets Shield field to given value.

### HasShield

`func (o *BackendResponse) HasShield() bool`

HasShield returns a boolean if a field has been set.

### SetShieldNil

`func (o *BackendResponse) SetShieldNil(b bool)`

 SetShieldNil sets the value for Shield to be an explicit nil

### UnsetShield
`func (o *BackendResponse) UnsetShield()`

UnsetShield ensures that no value is present for Shield, not even an explicit nil
### GetSslCaCert

`func (o *BackendResponse) GetSslCaCert() string`

GetSslCaCert returns the SslCaCert field if non-nil, zero value otherwise.

### GetSslCaCertOk

`func (o *BackendResponse) GetSslCaCertOk() (*string, bool)`

GetSslCaCertOk returns a tuple with the SslCaCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCaCert

`func (o *BackendResponse) SetSslCaCert(v string)`

SetSslCaCert sets SslCaCert field to given value.

### HasSslCaCert

`func (o *BackendResponse) HasSslCaCert() bool`

HasSslCaCert returns a boolean if a field has been set.

### SetSslCaCertNil

`func (o *BackendResponse) SetSslCaCertNil(b bool)`

 SetSslCaCertNil sets the value for SslCaCert to be an explicit nil

### UnsetSslCaCert
`func (o *BackendResponse) UnsetSslCaCert()`

UnsetSslCaCert ensures that no value is present for SslCaCert, not even an explicit nil
### GetSslCertHostname

`func (o *BackendResponse) GetSslCertHostname() string`

GetSslCertHostname returns the SslCertHostname field if non-nil, zero value otherwise.

### GetSslCertHostnameOk

`func (o *BackendResponse) GetSslCertHostnameOk() (*string, bool)`

GetSslCertHostnameOk returns a tuple with the SslCertHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCertHostname

`func (o *BackendResponse) SetSslCertHostname(v string)`

SetSslCertHostname sets SslCertHostname field to given value.

### HasSslCertHostname

`func (o *BackendResponse) HasSslCertHostname() bool`

HasSslCertHostname returns a boolean if a field has been set.

### SetSslCertHostnameNil

`func (o *BackendResponse) SetSslCertHostnameNil(b bool)`

 SetSslCertHostnameNil sets the value for SslCertHostname to be an explicit nil

### UnsetSslCertHostname
`func (o *BackendResponse) UnsetSslCertHostname()`

UnsetSslCertHostname ensures that no value is present for SslCertHostname, not even an explicit nil
### GetSslCheckCert

`func (o *BackendResponse) GetSslCheckCert() bool`

GetSslCheckCert returns the SslCheckCert field if non-nil, zero value otherwise.

### GetSslCheckCertOk

`func (o *BackendResponse) GetSslCheckCertOk() (*bool, bool)`

GetSslCheckCertOk returns a tuple with the SslCheckCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCheckCert

`func (o *BackendResponse) SetSslCheckCert(v bool)`

SetSslCheckCert sets SslCheckCert field to given value.

### HasSslCheckCert

`func (o *BackendResponse) HasSslCheckCert() bool`

HasSslCheckCert returns a boolean if a field has been set.

### GetSslCiphers

`func (o *BackendResponse) GetSslCiphers() string`

GetSslCiphers returns the SslCiphers field if non-nil, zero value otherwise.

### GetSslCiphersOk

`func (o *BackendResponse) GetSslCiphersOk() (*string, bool)`

GetSslCiphersOk returns a tuple with the SslCiphers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCiphers

`func (o *BackendResponse) SetSslCiphers(v string)`

SetSslCiphers sets SslCiphers field to given value.

### HasSslCiphers

`func (o *BackendResponse) HasSslCiphers() bool`

HasSslCiphers returns a boolean if a field has been set.

### SetSslCiphersNil

`func (o *BackendResponse) SetSslCiphersNil(b bool)`

 SetSslCiphersNil sets the value for SslCiphers to be an explicit nil

### UnsetSslCiphers
`func (o *BackendResponse) UnsetSslCiphers()`

UnsetSslCiphers ensures that no value is present for SslCiphers, not even an explicit nil
### GetSslClientCert

`func (o *BackendResponse) GetSslClientCert() string`

GetSslClientCert returns the SslClientCert field if non-nil, zero value otherwise.

### GetSslClientCertOk

`func (o *BackendResponse) GetSslClientCertOk() (*string, bool)`

GetSslClientCertOk returns a tuple with the SslClientCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslClientCert

`func (o *BackendResponse) SetSslClientCert(v string)`

SetSslClientCert sets SslClientCert field to given value.

### HasSslClientCert

`func (o *BackendResponse) HasSslClientCert() bool`

HasSslClientCert returns a boolean if a field has been set.

### SetSslClientCertNil

`func (o *BackendResponse) SetSslClientCertNil(b bool)`

 SetSslClientCertNil sets the value for SslClientCert to be an explicit nil

### UnsetSslClientCert
`func (o *BackendResponse) UnsetSslClientCert()`

UnsetSslClientCert ensures that no value is present for SslClientCert, not even an explicit nil
### GetSslClientKey

`func (o *BackendResponse) GetSslClientKey() string`

GetSslClientKey returns the SslClientKey field if non-nil, zero value otherwise.

### GetSslClientKeyOk

`func (o *BackendResponse) GetSslClientKeyOk() (*string, bool)`

GetSslClientKeyOk returns a tuple with the SslClientKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslClientKey

`func (o *BackendResponse) SetSslClientKey(v string)`

SetSslClientKey sets SslClientKey field to given value.

### HasSslClientKey

`func (o *BackendResponse) HasSslClientKey() bool`

HasSslClientKey returns a boolean if a field has been set.

### SetSslClientKeyNil

`func (o *BackendResponse) SetSslClientKeyNil(b bool)`

 SetSslClientKeyNil sets the value for SslClientKey to be an explicit nil

### UnsetSslClientKey
`func (o *BackendResponse) UnsetSslClientKey()`

UnsetSslClientKey ensures that no value is present for SslClientKey, not even an explicit nil
### GetSslHostname

`func (o *BackendResponse) GetSslHostname() string`

GetSslHostname returns the SslHostname field if non-nil, zero value otherwise.

### GetSslHostnameOk

`func (o *BackendResponse) GetSslHostnameOk() (*string, bool)`

GetSslHostnameOk returns a tuple with the SslHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslHostname

`func (o *BackendResponse) SetSslHostname(v string)`

SetSslHostname sets SslHostname field to given value.

### HasSslHostname

`func (o *BackendResponse) HasSslHostname() bool`

HasSslHostname returns a boolean if a field has been set.

### SetSslHostnameNil

`func (o *BackendResponse) SetSslHostnameNil(b bool)`

 SetSslHostnameNil sets the value for SslHostname to be an explicit nil

### UnsetSslHostname
`func (o *BackendResponse) UnsetSslHostname()`

UnsetSslHostname ensures that no value is present for SslHostname, not even an explicit nil
### GetSslSniHostname

`func (o *BackendResponse) GetSslSniHostname() string`

GetSslSniHostname returns the SslSniHostname field if non-nil, zero value otherwise.

### GetSslSniHostnameOk

`func (o *BackendResponse) GetSslSniHostnameOk() (*string, bool)`

GetSslSniHostnameOk returns a tuple with the SslSniHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslSniHostname

`func (o *BackendResponse) SetSslSniHostname(v string)`

SetSslSniHostname sets SslSniHostname field to given value.

### HasSslSniHostname

`func (o *BackendResponse) HasSslSniHostname() bool`

HasSslSniHostname returns a boolean if a field has been set.

### SetSslSniHostnameNil

`func (o *BackendResponse) SetSslSniHostnameNil(b bool)`

 SetSslSniHostnameNil sets the value for SslSniHostname to be an explicit nil

### UnsetSslSniHostname
`func (o *BackendResponse) UnsetSslSniHostname()`

UnsetSslSniHostname ensures that no value is present for SslSniHostname, not even an explicit nil
### GetUseSsl

`func (o *BackendResponse) GetUseSsl() bool`

GetUseSsl returns the UseSsl field if non-nil, zero value otherwise.

### GetUseSslOk

`func (o *BackendResponse) GetUseSslOk() (*bool, bool)`

GetUseSslOk returns a tuple with the UseSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSsl

`func (o *BackendResponse) SetUseSsl(v bool)`

SetUseSsl sets UseSsl field to given value.

### HasUseSsl

`func (o *BackendResponse) HasUseSsl() bool`

HasUseSsl returns a boolean if a field has been set.

### GetWeight

`func (o *BackendResponse) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *BackendResponse) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *BackendResponse) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *BackendResponse) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetCreatedAt

`func (o *BackendResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BackendResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BackendResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BackendResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *BackendResponse) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *BackendResponse) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetDeletedAt

`func (o *BackendResponse) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *BackendResponse) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *BackendResponse) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *BackendResponse) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *BackendResponse) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *BackendResponse) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetUpdatedAt

`func (o *BackendResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BackendResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BackendResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *BackendResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *BackendResponse) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *BackendResponse) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetServiceID

`func (o *BackendResponse) GetServiceID() string`

GetServiceID returns the ServiceID field if non-nil, zero value otherwise.

### GetServiceIDOk

`func (o *BackendResponse) GetServiceIDOk() (*string, bool)`

GetServiceIDOk returns a tuple with the ServiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceID

`func (o *BackendResponse) SetServiceID(v string)`

SetServiceID sets ServiceID field to given value.

### HasServiceID

`func (o *BackendResponse) HasServiceID() bool`

HasServiceID returns a boolean if a field has been set.

### GetVersion

`func (o *BackendResponse) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *BackendResponse) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *BackendResponse) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *BackendResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetLocked

`func (o *BackendResponse) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *BackendResponse) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *BackendResponse) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *BackendResponse) HasLocked() bool`

HasLocked returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
