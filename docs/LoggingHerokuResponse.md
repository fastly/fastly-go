# LoggingHerokuResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name for the real-time logging configuration. | [optional] 
**Placement** | Pointer to **NullableString** | Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.  | [optional] 
**ResponseCondition** | Pointer to **NullableString** | The name of an existing condition in the configured endpoint, or leave blank to always execute. | [optional] 
**Format** | Pointer to **string** | A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats). | [optional] [default to "%h %l %u %t \"%r\" %&gt;s %b"]
**FormatVersion** | Pointer to **string** | The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.  | [optional] [default to "2"]
**Token** | Pointer to **string** | The token to use for authentication ([https://devcenter.heroku.com/articles/add-on-partner-log-integration](https://devcenter.heroku.com/articles/add-on-partner-log-integration)). | [optional] 
**URL** | Pointer to **string** | The URL to stream logs to. | [optional] 
**CreatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**DeletedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**UpdatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**ServiceID** | Pointer to **string** |  | [optional] [readonly] 
**Version** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewLoggingHerokuResponse

`func NewLoggingHerokuResponse() *LoggingHerokuResponse`

NewLoggingHerokuResponse instantiates a new LoggingHerokuResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoggingHerokuResponseWithDefaults

`func NewLoggingHerokuResponseWithDefaults() *LoggingHerokuResponse`

NewLoggingHerokuResponseWithDefaults instantiates a new LoggingHerokuResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *LoggingHerokuResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LoggingHerokuResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LoggingHerokuResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LoggingHerokuResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlacement

`func (o *LoggingHerokuResponse) GetPlacement() string`

GetPlacement returns the Placement field if non-nil, zero value otherwise.

### GetPlacementOk

`func (o *LoggingHerokuResponse) GetPlacementOk() (*string, bool)`

GetPlacementOk returns a tuple with the Placement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacement

`func (o *LoggingHerokuResponse) SetPlacement(v string)`

SetPlacement sets Placement field to given value.

### HasPlacement

`func (o *LoggingHerokuResponse) HasPlacement() bool`

HasPlacement returns a boolean if a field has been set.

### SetPlacementNil

`func (o *LoggingHerokuResponse) SetPlacementNil(b bool)`

 SetPlacementNil sets the value for Placement to be an explicit nil

### UnsetPlacement
`func (o *LoggingHerokuResponse) UnsetPlacement()`

UnsetPlacement ensures that no value is present for Placement, not even an explicit nil
### GetResponseCondition

`func (o *LoggingHerokuResponse) GetResponseCondition() string`

GetResponseCondition returns the ResponseCondition field if non-nil, zero value otherwise.

### GetResponseConditionOk

`func (o *LoggingHerokuResponse) GetResponseConditionOk() (*string, bool)`

GetResponseConditionOk returns a tuple with the ResponseCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseCondition

`func (o *LoggingHerokuResponse) SetResponseCondition(v string)`

SetResponseCondition sets ResponseCondition field to given value.

### HasResponseCondition

`func (o *LoggingHerokuResponse) HasResponseCondition() bool`

HasResponseCondition returns a boolean if a field has been set.

### SetResponseConditionNil

`func (o *LoggingHerokuResponse) SetResponseConditionNil(b bool)`

 SetResponseConditionNil sets the value for ResponseCondition to be an explicit nil

### UnsetResponseCondition
`func (o *LoggingHerokuResponse) UnsetResponseCondition()`

UnsetResponseCondition ensures that no value is present for ResponseCondition, not even an explicit nil
### GetFormat

`func (o *LoggingHerokuResponse) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *LoggingHerokuResponse) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *LoggingHerokuResponse) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *LoggingHerokuResponse) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetFormatVersion

`func (o *LoggingHerokuResponse) GetFormatVersion() string`

GetFormatVersion returns the FormatVersion field if non-nil, zero value otherwise.

### GetFormatVersionOk

`func (o *LoggingHerokuResponse) GetFormatVersionOk() (*string, bool)`

GetFormatVersionOk returns a tuple with the FormatVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatVersion

`func (o *LoggingHerokuResponse) SetFormatVersion(v string)`

SetFormatVersion sets FormatVersion field to given value.

### HasFormatVersion

`func (o *LoggingHerokuResponse) HasFormatVersion() bool`

HasFormatVersion returns a boolean if a field has been set.

### GetToken

`func (o *LoggingHerokuResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *LoggingHerokuResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *LoggingHerokuResponse) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *LoggingHerokuResponse) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetURL

`func (o *LoggingHerokuResponse) GetURL() string`

GetURL returns the URL field if non-nil, zero value otherwise.

### GetURLOk

`func (o *LoggingHerokuResponse) GetURLOk() (*string, bool)`

GetURLOk returns a tuple with the URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetURL

`func (o *LoggingHerokuResponse) SetURL(v string)`

SetURL sets URL field to given value.

### HasURL

`func (o *LoggingHerokuResponse) HasURL() bool`

HasURL returns a boolean if a field has been set.

### GetCreatedAt

`func (o *LoggingHerokuResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LoggingHerokuResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LoggingHerokuResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *LoggingHerokuResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *LoggingHerokuResponse) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *LoggingHerokuResponse) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetDeletedAt

`func (o *LoggingHerokuResponse) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *LoggingHerokuResponse) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *LoggingHerokuResponse) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *LoggingHerokuResponse) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *LoggingHerokuResponse) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *LoggingHerokuResponse) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetUpdatedAt

`func (o *LoggingHerokuResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *LoggingHerokuResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *LoggingHerokuResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *LoggingHerokuResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *LoggingHerokuResponse) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *LoggingHerokuResponse) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetServiceID

`func (o *LoggingHerokuResponse) GetServiceID() string`

GetServiceID returns the ServiceID field if non-nil, zero value otherwise.

### GetServiceIDOk

`func (o *LoggingHerokuResponse) GetServiceIDOk() (*string, bool)`

GetServiceIDOk returns a tuple with the ServiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceID

`func (o *LoggingHerokuResponse) SetServiceID(v string)`

SetServiceID sets ServiceID field to given value.

### HasServiceID

`func (o *LoggingHerokuResponse) HasServiceID() bool`

HasServiceID returns a boolean if a field has been set.

### GetVersion

`func (o *LoggingHerokuResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *LoggingHerokuResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *LoggingHerokuResponse) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *LoggingHerokuResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
