# TLSConfigurationResponseAttributesAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Default** | Pointer to **bool** | Signifies whether or not Fastly will use this configuration as a default when creating a new [TLS Activation](https://www.fastly.com/documentation/reference/api/tls/custom-certs/activations/). | [optional] [readonly] 
**HTTPProtocols** | Pointer to **[]string** | HTTP protocols available on your configuration. | [optional] [readonly] 
**TLSProtocols** | Pointer to **[]string** | TLS protocols available on your configuration. | [optional] [readonly] 
**Bulk** | Pointer to **bool** | Signifies whether the configuration is used for Platform TLS or not. | [optional] [readonly] 

## Methods

### NewTLSConfigurationResponseAttributesAllOf

`func NewTLSConfigurationResponseAttributesAllOf() *TLSConfigurationResponseAttributesAllOf`

NewTLSConfigurationResponseAttributesAllOf instantiates a new TLSConfigurationResponseAttributesAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTLSConfigurationResponseAttributesAllOfWithDefaults

`func NewTLSConfigurationResponseAttributesAllOfWithDefaults() *TLSConfigurationResponseAttributesAllOf`

NewTLSConfigurationResponseAttributesAllOfWithDefaults instantiates a new TLSConfigurationResponseAttributesAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefault

`func (o *TLSConfigurationResponseAttributesAllOf) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *TLSConfigurationResponseAttributesAllOf) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *TLSConfigurationResponseAttributesAllOf) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *TLSConfigurationResponseAttributesAllOf) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetHTTPProtocols

`func (o *TLSConfigurationResponseAttributesAllOf) GetHTTPProtocols() []string`

GetHTTPProtocols returns the HTTPProtocols field if non-nil, zero value otherwise.

### GetHTTPProtocolsOk

`func (o *TLSConfigurationResponseAttributesAllOf) GetHTTPProtocolsOk() (*[]string, bool)`

GetHTTPProtocolsOk returns a tuple with the HTTPProtocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHTTPProtocols

`func (o *TLSConfigurationResponseAttributesAllOf) SetHTTPProtocols(v []string)`

SetHTTPProtocols sets HTTPProtocols field to given value.

### HasHTTPProtocols

`func (o *TLSConfigurationResponseAttributesAllOf) HasHTTPProtocols() bool`

HasHTTPProtocols returns a boolean if a field has been set.

### GetTLSProtocols

`func (o *TLSConfigurationResponseAttributesAllOf) GetTLSProtocols() []string`

GetTLSProtocols returns the TLSProtocols field if non-nil, zero value otherwise.

### GetTLSProtocolsOk

`func (o *TLSConfigurationResponseAttributesAllOf) GetTLSProtocolsOk() (*[]string, bool)`

GetTLSProtocolsOk returns a tuple with the TLSProtocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTLSProtocols

`func (o *TLSConfigurationResponseAttributesAllOf) SetTLSProtocols(v []string)`

SetTLSProtocols sets TLSProtocols field to given value.

### HasTLSProtocols

`func (o *TLSConfigurationResponseAttributesAllOf) HasTLSProtocols() bool`

HasTLSProtocols returns a boolean if a field has been set.

### GetBulk

`func (o *TLSConfigurationResponseAttributesAllOf) GetBulk() bool`

GetBulk returns the Bulk field if non-nil, zero value otherwise.

### GetBulkOk

`func (o *TLSConfigurationResponseAttributesAllOf) GetBulkOk() (*bool, bool)`

GetBulkOk returns a tuple with the Bulk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulk

`func (o *TLSConfigurationResponseAttributesAllOf) SetBulk(v bool)`

SetBulk sets Bulk field to given value.

### HasBulk

`func (o *TLSConfigurationResponseAttributesAllOf) HasBulk() bool`

HasBulk returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
