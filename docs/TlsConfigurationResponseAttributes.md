# TLSConfigurationResponseAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**DeletedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**UpdatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**Default** | Pointer to **bool** | Signifies whether or not Fastly will use this configuration as a default when creating a new [TLS Activation](/reference/api/tls/custom-certs/activations/). | [optional] [readonly] 
**HTTPProtocols** | Pointer to **[]string** | HTTP protocols available on your configuration. | [optional] [readonly] 
**TLSProtocols** | Pointer to **[]string** | TLS protocols available on your configuration. | [optional] [readonly] 
**Bulk** | Pointer to **bool** | Signifies whether the configuration is used for Platform TLS or not. | [optional] [readonly] 

## Methods

### NewTLSConfigurationResponseAttributes

`func NewTLSConfigurationResponseAttributes() *TLSConfigurationResponseAttributes`

NewTLSConfigurationResponseAttributes instantiates a new TLSConfigurationResponseAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTLSConfigurationResponseAttributesWithDefaults

`func NewTLSConfigurationResponseAttributesWithDefaults() *TLSConfigurationResponseAttributes`

NewTLSConfigurationResponseAttributesWithDefaults instantiates a new TLSConfigurationResponseAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *TLSConfigurationResponseAttributes) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TLSConfigurationResponseAttributes) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TLSConfigurationResponseAttributes) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TLSConfigurationResponseAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *TLSConfigurationResponseAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *TLSConfigurationResponseAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetDeletedAt

`func (o *TLSConfigurationResponseAttributes) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *TLSConfigurationResponseAttributes) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *TLSConfigurationResponseAttributes) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *TLSConfigurationResponseAttributes) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *TLSConfigurationResponseAttributes) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *TLSConfigurationResponseAttributes) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetUpdatedAt

`func (o *TLSConfigurationResponseAttributes) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TLSConfigurationResponseAttributes) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TLSConfigurationResponseAttributes) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TLSConfigurationResponseAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *TLSConfigurationResponseAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *TLSConfigurationResponseAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetDefault

`func (o *TLSConfigurationResponseAttributes) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *TLSConfigurationResponseAttributes) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *TLSConfigurationResponseAttributes) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *TLSConfigurationResponseAttributes) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetHTTPProtocols

`func (o *TLSConfigurationResponseAttributes) GetHTTPProtocols() []string`

GetHTTPProtocols returns the HTTPProtocols field if non-nil, zero value otherwise.

### GetHTTPProtocolsOk

`func (o *TLSConfigurationResponseAttributes) GetHTTPProtocolsOk() (*[]string, bool)`

GetHTTPProtocolsOk returns a tuple with the HTTPProtocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHTTPProtocols

`func (o *TLSConfigurationResponseAttributes) SetHTTPProtocols(v []string)`

SetHTTPProtocols sets HTTPProtocols field to given value.

### HasHTTPProtocols

`func (o *TLSConfigurationResponseAttributes) HasHTTPProtocols() bool`

HasHTTPProtocols returns a boolean if a field has been set.

### GetTLSProtocols

`func (o *TLSConfigurationResponseAttributes) GetTLSProtocols() []string`

GetTLSProtocols returns the TLSProtocols field if non-nil, zero value otherwise.

### GetTLSProtocolsOk

`func (o *TLSConfigurationResponseAttributes) GetTLSProtocolsOk() (*[]string, bool)`

GetTLSProtocolsOk returns a tuple with the TLSProtocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTLSProtocols

`func (o *TLSConfigurationResponseAttributes) SetTLSProtocols(v []string)`

SetTLSProtocols sets TLSProtocols field to given value.

### HasTLSProtocols

`func (o *TLSConfigurationResponseAttributes) HasTLSProtocols() bool`

HasTLSProtocols returns a boolean if a field has been set.

### GetBulk

`func (o *TLSConfigurationResponseAttributes) GetBulk() bool`

GetBulk returns the Bulk field if non-nil, zero value otherwise.

### GetBulkOk

`func (o *TLSConfigurationResponseAttributes) GetBulkOk() (*bool, bool)`

GetBulkOk returns a tuple with the Bulk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulk

`func (o *TLSConfigurationResponseAttributes) SetBulk(v bool)`

SetBulk sets Bulk field to given value.

### HasBulk

`func (o *TLSConfigurationResponseAttributes) HasBulk() bool`

HasBulk returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
