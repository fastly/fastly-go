# LoggingBigquery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the BigQuery logging object. Used as a primary key for API access. | [optional] 
**Placement** | Pointer to **NullableString** | Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.  | [optional] 
**FormatVersion** | Pointer to **int32** | The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.  | [optional] [default to 2]
**ResponseCondition** | Pointer to **NullableString** | The name of an existing condition in the configured endpoint, or leave blank to always execute. | [optional] 
**Format** | Pointer to **string** | A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats). Must produce JSON that matches the schema of your BigQuery table. | [optional] 
**User** | Pointer to **string** | Your Google Cloud Platform service account email address. The `client_email` field in your service account authentication JSON. Not required if `account_name` is specified. | [optional] 
**SecretKey** | Pointer to **string** | Your Google Cloud Platform account secret key. The `private_key` field in your service account authentication JSON. Not required if `account_name` is specified. | [optional] 
**AccountName** | Pointer to **string** | The name of the Google Cloud Platform service account associated with the target log collection service. Not required if `user` and `secret_key` are provided. | [optional] 
**Dataset** | Pointer to **string** | Your BigQuery dataset. | [optional] 
**Table** | Pointer to **string** | Your BigQuery table. | [optional] 
**TemplateSuffix** | Pointer to **NullableString** | BigQuery table name suffix template. Optional. | [optional] 
**ProjectID** | Pointer to **string** | Your Google Cloud Platform project ID. Required | [optional] 

## Methods

### NewLoggingBigquery

`func NewLoggingBigquery() *LoggingBigquery`

NewLoggingBigquery instantiates a new LoggingBigquery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoggingBigqueryWithDefaults

`func NewLoggingBigqueryWithDefaults() *LoggingBigquery`

NewLoggingBigqueryWithDefaults instantiates a new LoggingBigquery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *LoggingBigquery) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LoggingBigquery) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LoggingBigquery) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LoggingBigquery) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlacement

`func (o *LoggingBigquery) GetPlacement() string`

GetPlacement returns the Placement field if non-nil, zero value otherwise.

### GetPlacementOk

`func (o *LoggingBigquery) GetPlacementOk() (*string, bool)`

GetPlacementOk returns a tuple with the Placement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacement

`func (o *LoggingBigquery) SetPlacement(v string)`

SetPlacement sets Placement field to given value.

### HasPlacement

`func (o *LoggingBigquery) HasPlacement() bool`

HasPlacement returns a boolean if a field has been set.

### SetPlacementNil

`func (o *LoggingBigquery) SetPlacementNil(b bool)`

 SetPlacementNil sets the value for Placement to be an explicit nil

### UnsetPlacement
`func (o *LoggingBigquery) UnsetPlacement()`

UnsetPlacement ensures that no value is present for Placement, not even an explicit nil
### GetFormatVersion

`func (o *LoggingBigquery) GetFormatVersion() int32`

GetFormatVersion returns the FormatVersion field if non-nil, zero value otherwise.

### GetFormatVersionOk

`func (o *LoggingBigquery) GetFormatVersionOk() (*int32, bool)`

GetFormatVersionOk returns a tuple with the FormatVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatVersion

`func (o *LoggingBigquery) SetFormatVersion(v int32)`

SetFormatVersion sets FormatVersion field to given value.

### HasFormatVersion

`func (o *LoggingBigquery) HasFormatVersion() bool`

HasFormatVersion returns a boolean if a field has been set.

### GetResponseCondition

`func (o *LoggingBigquery) GetResponseCondition() string`

GetResponseCondition returns the ResponseCondition field if non-nil, zero value otherwise.

### GetResponseConditionOk

`func (o *LoggingBigquery) GetResponseConditionOk() (*string, bool)`

GetResponseConditionOk returns a tuple with the ResponseCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseCondition

`func (o *LoggingBigquery) SetResponseCondition(v string)`

SetResponseCondition sets ResponseCondition field to given value.

### HasResponseCondition

`func (o *LoggingBigquery) HasResponseCondition() bool`

HasResponseCondition returns a boolean if a field has been set.

### SetResponseConditionNil

`func (o *LoggingBigquery) SetResponseConditionNil(b bool)`

 SetResponseConditionNil sets the value for ResponseCondition to be an explicit nil

### UnsetResponseCondition
`func (o *LoggingBigquery) UnsetResponseCondition()`

UnsetResponseCondition ensures that no value is present for ResponseCondition, not even an explicit nil
### GetFormat

`func (o *LoggingBigquery) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *LoggingBigquery) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *LoggingBigquery) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *LoggingBigquery) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetUser

`func (o *LoggingBigquery) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *LoggingBigquery) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *LoggingBigquery) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *LoggingBigquery) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetSecretKey

`func (o *LoggingBigquery) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *LoggingBigquery) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *LoggingBigquery) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *LoggingBigquery) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### GetAccountName

`func (o *LoggingBigquery) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *LoggingBigquery) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *LoggingBigquery) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *LoggingBigquery) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetDataset

`func (o *LoggingBigquery) GetDataset() string`

GetDataset returns the Dataset field if non-nil, zero value otherwise.

### GetDatasetOk

`func (o *LoggingBigquery) GetDatasetOk() (*string, bool)`

GetDatasetOk returns a tuple with the Dataset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataset

`func (o *LoggingBigquery) SetDataset(v string)`

SetDataset sets Dataset field to given value.

### HasDataset

`func (o *LoggingBigquery) HasDataset() bool`

HasDataset returns a boolean if a field has been set.

### GetTable

`func (o *LoggingBigquery) GetTable() string`

GetTable returns the Table field if non-nil, zero value otherwise.

### GetTableOk

`func (o *LoggingBigquery) GetTableOk() (*string, bool)`

GetTableOk returns a tuple with the Table field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTable

`func (o *LoggingBigquery) SetTable(v string)`

SetTable sets Table field to given value.

### HasTable

`func (o *LoggingBigquery) HasTable() bool`

HasTable returns a boolean if a field has been set.

### GetTemplateSuffix

`func (o *LoggingBigquery) GetTemplateSuffix() string`

GetTemplateSuffix returns the TemplateSuffix field if non-nil, zero value otherwise.

### GetTemplateSuffixOk

`func (o *LoggingBigquery) GetTemplateSuffixOk() (*string, bool)`

GetTemplateSuffixOk returns a tuple with the TemplateSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateSuffix

`func (o *LoggingBigquery) SetTemplateSuffix(v string)`

SetTemplateSuffix sets TemplateSuffix field to given value.

### HasTemplateSuffix

`func (o *LoggingBigquery) HasTemplateSuffix() bool`

HasTemplateSuffix returns a boolean if a field has been set.

### SetTemplateSuffixNil

`func (o *LoggingBigquery) SetTemplateSuffixNil(b bool)`

 SetTemplateSuffixNil sets the value for TemplateSuffix to be an explicit nil

### UnsetTemplateSuffix
`func (o *LoggingBigquery) UnsetTemplateSuffix()`

UnsetTemplateSuffix ensures that no value is present for TemplateSuffix, not even an explicit nil
### GetProjectID

`func (o *LoggingBigquery) GetProjectID() string`

GetProjectID returns the ProjectID field if non-nil, zero value otherwise.

### GetProjectIDOk

`func (o *LoggingBigquery) GetProjectIDOk() (*string, bool)`

GetProjectIDOk returns a tuple with the ProjectID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectID

`func (o *LoggingBigquery) SetProjectID(v string)`

SetProjectID sets ProjectID field to given value.

### HasProjectID

`func (o *LoggingBigquery) HasProjectID() bool`

HasProjectID returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
