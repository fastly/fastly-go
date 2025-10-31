# TlsConfigurationResponseAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**DeletedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**UpdatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**Default** | Pointer to **bool** | Signifies whether or not Fastly will use this configuration as a default when creating a new [TLS Activation](https://www.fastly.com/documentation/reference/api/tls/custom-certs/activations/). | [optional] [readonly] 
**HttpProtocols** | Pointer to **[]string** | HTTP protocols available on your configuration. | [optional] [readonly] 
**TlsProtocols** | Pointer to **[]string** | TLS protocols available on your configuration. | [optional] [readonly] 
**Bulk** | Pointer to **bool** | Signifies whether the configuration is used for Platform TLS or not. | [optional] [readonly] 

## Methods

### NewTlsConfigurationResponseAttributes

`func NewTlsConfigurationResponseAttributes() *TlsConfigurationResponseAttributes`

NewTlsConfigurationResponseAttributes instantiates a new TlsConfigurationResponseAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTlsConfigurationResponseAttributesWithDefaults

`func NewTlsConfigurationResponseAttributesWithDefaults() *TlsConfigurationResponseAttributes`

NewTlsConfigurationResponseAttributesWithDefaults instantiates a new TlsConfigurationResponseAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *TlsConfigurationResponseAttributes) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TlsConfigurationResponseAttributes) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TlsConfigurationResponseAttributes) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TlsConfigurationResponseAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *TlsConfigurationResponseAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *TlsConfigurationResponseAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetDeletedAt

`func (o *TlsConfigurationResponseAttributes) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *TlsConfigurationResponseAttributes) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *TlsConfigurationResponseAttributes) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *TlsConfigurationResponseAttributes) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *TlsConfigurationResponseAttributes) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *TlsConfigurationResponseAttributes) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetUpdatedAt

`func (o *TlsConfigurationResponseAttributes) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TlsConfigurationResponseAttributes) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TlsConfigurationResponseAttributes) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TlsConfigurationResponseAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *TlsConfigurationResponseAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *TlsConfigurationResponseAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetDefault

`func (o *TlsConfigurationResponseAttributes) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *TlsConfigurationResponseAttributes) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *TlsConfigurationResponseAttributes) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *TlsConfigurationResponseAttributes) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetHttpProtocols

`func (o *TlsConfigurationResponseAttributes) GetHttpProtocols() []string`

GetHttpProtocols returns the HttpProtocols field if non-nil, zero value otherwise.

### GetHttpProtocolsOk

`func (o *TlsConfigurationResponseAttributes) GetHttpProtocolsOk() (*[]string, bool)`

GetHttpProtocolsOk returns a tuple with the HttpProtocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpProtocols

`func (o *TlsConfigurationResponseAttributes) SetHttpProtocols(v []string)`

SetHttpProtocols sets HttpProtocols field to given value.

### HasHttpProtocols

`func (o *TlsConfigurationResponseAttributes) HasHttpProtocols() bool`

HasHttpProtocols returns a boolean if a field has been set.

### GetTlsProtocols

`func (o *TlsConfigurationResponseAttributes) GetTlsProtocols() []string`

GetTlsProtocols returns the TlsProtocols field if non-nil, zero value otherwise.

### GetTlsProtocolsOk

`func (o *TlsConfigurationResponseAttributes) GetTlsProtocolsOk() (*[]string, bool)`

GetTlsProtocolsOk returns a tuple with the TlsProtocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsProtocols

`func (o *TlsConfigurationResponseAttributes) SetTlsProtocols(v []string)`

SetTlsProtocols sets TlsProtocols field to given value.

### HasTlsProtocols

`func (o *TlsConfigurationResponseAttributes) HasTlsProtocols() bool`

HasTlsProtocols returns a boolean if a field has been set.

### GetBulk

`func (o *TlsConfigurationResponseAttributes) GetBulk() bool`

GetBulk returns the Bulk field if non-nil, zero value otherwise.

### GetBulkOk

`func (o *TlsConfigurationResponseAttributes) GetBulkOk() (*bool, bool)`

GetBulkOk returns a tuple with the Bulk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulk

`func (o *TlsConfigurationResponseAttributes) SetBulk(v bool)`

SetBulk sets Bulk field to given value.

### HasBulk

`func (o *TlsConfigurationResponseAttributes) HasBulk() bool`

HasBulk returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


