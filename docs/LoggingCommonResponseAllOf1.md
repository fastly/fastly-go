# LoggingCommonResponseAllOf1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FormatVersion** | Pointer to **string** | The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.  | [optional] [default to "2"]

## Methods

### NewLoggingCommonResponseAllOf1

`func NewLoggingCommonResponseAllOf1() *LoggingCommonResponseAllOf1`

NewLoggingCommonResponseAllOf1 instantiates a new LoggingCommonResponseAllOf1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoggingCommonResponseAllOf1WithDefaults

`func NewLoggingCommonResponseAllOf1WithDefaults() *LoggingCommonResponseAllOf1`

NewLoggingCommonResponseAllOf1WithDefaults instantiates a new LoggingCommonResponseAllOf1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormatVersion

`func (o *LoggingCommonResponseAllOf1) GetFormatVersion() string`

GetFormatVersion returns the FormatVersion field if non-nil, zero value otherwise.

### GetFormatVersionOk

`func (o *LoggingCommonResponseAllOf1) GetFormatVersionOk() (*string, bool)`

GetFormatVersionOk returns a tuple with the FormatVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatVersion

`func (o *LoggingCommonResponseAllOf1) SetFormatVersion(v string)`

SetFormatVersion sets FormatVersion field to given value.

### HasFormatVersion

`func (o *LoggingCommonResponseAllOf1) HasFormatVersion() bool`

HasFormatVersion returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


