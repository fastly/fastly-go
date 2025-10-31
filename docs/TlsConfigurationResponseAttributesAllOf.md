# TlsConfigurationResponseAttributesAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Default** | Pointer to **bool** | Signifies whether or not Fastly will use this configuration as a default when creating a new [TLS Activation](https://www.fastly.com/documentation/reference/api/tls/custom-certs/activations/). | [optional] [readonly] 
**HttpProtocols** | Pointer to **[]string** | HTTP protocols available on your configuration. | [optional] [readonly] 
**TlsProtocols** | Pointer to **[]string** | TLS protocols available on your configuration. | [optional] [readonly] 
**Bulk** | Pointer to **bool** | Signifies whether the configuration is used for Platform TLS or not. | [optional] [readonly] 

## Methods

### NewTlsConfigurationResponseAttributesAllOf

`func NewTlsConfigurationResponseAttributesAllOf() *TlsConfigurationResponseAttributesAllOf`

NewTlsConfigurationResponseAttributesAllOf instantiates a new TlsConfigurationResponseAttributesAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTlsConfigurationResponseAttributesAllOfWithDefaults

`func NewTlsConfigurationResponseAttributesAllOfWithDefaults() *TlsConfigurationResponseAttributesAllOf`

NewTlsConfigurationResponseAttributesAllOfWithDefaults instantiates a new TlsConfigurationResponseAttributesAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefault

`func (o *TlsConfigurationResponseAttributesAllOf) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *TlsConfigurationResponseAttributesAllOf) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *TlsConfigurationResponseAttributesAllOf) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *TlsConfigurationResponseAttributesAllOf) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetHttpProtocols

`func (o *TlsConfigurationResponseAttributesAllOf) GetHttpProtocols() []string`

GetHttpProtocols returns the HttpProtocols field if non-nil, zero value otherwise.

### GetHttpProtocolsOk

`func (o *TlsConfigurationResponseAttributesAllOf) GetHttpProtocolsOk() (*[]string, bool)`

GetHttpProtocolsOk returns a tuple with the HttpProtocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpProtocols

`func (o *TlsConfigurationResponseAttributesAllOf) SetHttpProtocols(v []string)`

SetHttpProtocols sets HttpProtocols field to given value.

### HasHttpProtocols

`func (o *TlsConfigurationResponseAttributesAllOf) HasHttpProtocols() bool`

HasHttpProtocols returns a boolean if a field has been set.

### GetTlsProtocols

`func (o *TlsConfigurationResponseAttributesAllOf) GetTlsProtocols() []string`

GetTlsProtocols returns the TlsProtocols field if non-nil, zero value otherwise.

### GetTlsProtocolsOk

`func (o *TlsConfigurationResponseAttributesAllOf) GetTlsProtocolsOk() (*[]string, bool)`

GetTlsProtocolsOk returns a tuple with the TlsProtocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsProtocols

`func (o *TlsConfigurationResponseAttributesAllOf) SetTlsProtocols(v []string)`

SetTlsProtocols sets TlsProtocols field to given value.

### HasTlsProtocols

`func (o *TlsConfigurationResponseAttributesAllOf) HasTlsProtocols() bool`

HasTlsProtocols returns a boolean if a field has been set.

### GetBulk

`func (o *TlsConfigurationResponseAttributesAllOf) GetBulk() bool`

GetBulk returns the Bulk field if non-nil, zero value otherwise.

### GetBulkOk

`func (o *TlsConfigurationResponseAttributesAllOf) GetBulkOk() (*bool, bool)`

GetBulkOk returns a tuple with the Bulk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulk

`func (o *TlsConfigurationResponseAttributesAllOf) SetBulk(v bool)`

SetBulk sets Bulk field to given value.

### HasBulk

`func (o *TlsConfigurationResponseAttributesAllOf) HasBulk() bool`

HasBulk returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


