# PoolResponseCommon

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BetweenBytesTimeout** | Pointer to **string** | Maximum duration in milliseconds that Fastly will wait while receiving no data on a download from a backend. If exceeded, for Delivery services, the response received so far will be considered complete and the fetch will end. For Compute services, timeout expiration is treated as a failure of the backend connection, and an error is generated. May be set at runtime using `bereq.between_bytes_timeout`. | [optional] 
**ConnectTimeout** | Pointer to **string** | How long to wait for a timeout in milliseconds. | [optional] 
**FirstByteTimeout** | Pointer to **string** | How long to wait for the first byte in milliseconds. | [optional] 
**MaxConnDefault** | Pointer to **string** | Maximum number of connections. | [optional] [default to "200"]
**TlsCheckCert** | Pointer to **NullableString** | Be strict on checking TLS certs. | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewPoolResponseCommon

`func NewPoolResponseCommon() *PoolResponseCommon`

NewPoolResponseCommon instantiates a new PoolResponseCommon object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolResponseCommonWithDefaults

`func NewPoolResponseCommonWithDefaults() *PoolResponseCommon`

NewPoolResponseCommonWithDefaults instantiates a new PoolResponseCommon object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBetweenBytesTimeout

`func (o *PoolResponseCommon) GetBetweenBytesTimeout() string`

GetBetweenBytesTimeout returns the BetweenBytesTimeout field if non-nil, zero value otherwise.

### GetBetweenBytesTimeoutOk

`func (o *PoolResponseCommon) GetBetweenBytesTimeoutOk() (*string, bool)`

GetBetweenBytesTimeoutOk returns a tuple with the BetweenBytesTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBetweenBytesTimeout

`func (o *PoolResponseCommon) SetBetweenBytesTimeout(v string)`

SetBetweenBytesTimeout sets BetweenBytesTimeout field to given value.

### HasBetweenBytesTimeout

`func (o *PoolResponseCommon) HasBetweenBytesTimeout() bool`

HasBetweenBytesTimeout returns a boolean if a field has been set.

### GetConnectTimeout

`func (o *PoolResponseCommon) GetConnectTimeout() string`

GetConnectTimeout returns the ConnectTimeout field if non-nil, zero value otherwise.

### GetConnectTimeoutOk

`func (o *PoolResponseCommon) GetConnectTimeoutOk() (*string, bool)`

GetConnectTimeoutOk returns a tuple with the ConnectTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectTimeout

`func (o *PoolResponseCommon) SetConnectTimeout(v string)`

SetConnectTimeout sets ConnectTimeout field to given value.

### HasConnectTimeout

`func (o *PoolResponseCommon) HasConnectTimeout() bool`

HasConnectTimeout returns a boolean if a field has been set.

### GetFirstByteTimeout

`func (o *PoolResponseCommon) GetFirstByteTimeout() string`

GetFirstByteTimeout returns the FirstByteTimeout field if non-nil, zero value otherwise.

### GetFirstByteTimeoutOk

`func (o *PoolResponseCommon) GetFirstByteTimeoutOk() (*string, bool)`

GetFirstByteTimeoutOk returns a tuple with the FirstByteTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstByteTimeout

`func (o *PoolResponseCommon) SetFirstByteTimeout(v string)`

SetFirstByteTimeout sets FirstByteTimeout field to given value.

### HasFirstByteTimeout

`func (o *PoolResponseCommon) HasFirstByteTimeout() bool`

HasFirstByteTimeout returns a boolean if a field has been set.

### GetMaxConnDefault

`func (o *PoolResponseCommon) GetMaxConnDefault() string`

GetMaxConnDefault returns the MaxConnDefault field if non-nil, zero value otherwise.

### GetMaxConnDefaultOk

`func (o *PoolResponseCommon) GetMaxConnDefaultOk() (*string, bool)`

GetMaxConnDefaultOk returns a tuple with the MaxConnDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConnDefault

`func (o *PoolResponseCommon) SetMaxConnDefault(v string)`

SetMaxConnDefault sets MaxConnDefault field to given value.

### HasMaxConnDefault

`func (o *PoolResponseCommon) HasMaxConnDefault() bool`

HasMaxConnDefault returns a boolean if a field has been set.

### GetTlsCheckCert

`func (o *PoolResponseCommon) GetTlsCheckCert() string`

GetTlsCheckCert returns the TlsCheckCert field if non-nil, zero value otherwise.

### GetTlsCheckCertOk

`func (o *PoolResponseCommon) GetTlsCheckCertOk() (*string, bool)`

GetTlsCheckCertOk returns a tuple with the TlsCheckCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsCheckCert

`func (o *PoolResponseCommon) SetTlsCheckCert(v string)`

SetTlsCheckCert sets TlsCheckCert field to given value.

### HasTlsCheckCert

`func (o *PoolResponseCommon) HasTlsCheckCert() bool`

HasTlsCheckCert returns a boolean if a field has been set.

### SetTlsCheckCertNil

`func (o *PoolResponseCommon) SetTlsCheckCertNil(b bool)`

 SetTlsCheckCertNil sets the value for TlsCheckCert to be an explicit nil

### UnsetTlsCheckCert
`func (o *PoolResponseCommon) UnsetTlsCheckCert()`

UnsetTlsCheckCert ensures that no value is present for TlsCheckCert, not even an explicit nil
### GetId

`func (o *PoolResponseCommon) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PoolResponseCommon) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PoolResponseCommon) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PoolResponseCommon) HasId() bool`

HasId returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


