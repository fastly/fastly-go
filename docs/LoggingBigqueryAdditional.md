# LoggingBigqueryAdditional

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the BigQuery logging object. Used as a primary key for API access. | [optional] 
**Format** | Pointer to **string** | A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats). Must produce JSON that matches the schema of your BigQuery table. | [optional] 
**Dataset** | Pointer to **string** | Your BigQuery dataset. | [optional] 
**Table** | Pointer to **string** | Your BigQuery table. | [optional] 
**TemplateSuffix** | Pointer to **NullableString** | BigQuery table name suffix template. Optional. | [optional] 
**ProjectID** | Pointer to **string** | Your Google Cloud Platform project ID. Required | [optional] 

## Methods

### NewLoggingBigqueryAdditional

`func NewLoggingBigqueryAdditional() *LoggingBigqueryAdditional`

NewLoggingBigqueryAdditional instantiates a new LoggingBigqueryAdditional object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoggingBigqueryAdditionalWithDefaults

`func NewLoggingBigqueryAdditionalWithDefaults() *LoggingBigqueryAdditional`

NewLoggingBigqueryAdditionalWithDefaults instantiates a new LoggingBigqueryAdditional object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *LoggingBigqueryAdditional) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LoggingBigqueryAdditional) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LoggingBigqueryAdditional) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LoggingBigqueryAdditional) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFormat

`func (o *LoggingBigqueryAdditional) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *LoggingBigqueryAdditional) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *LoggingBigqueryAdditional) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *LoggingBigqueryAdditional) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetDataset

`func (o *LoggingBigqueryAdditional) GetDataset() string`

GetDataset returns the Dataset field if non-nil, zero value otherwise.

### GetDatasetOk

`func (o *LoggingBigqueryAdditional) GetDatasetOk() (*string, bool)`

GetDatasetOk returns a tuple with the Dataset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataset

`func (o *LoggingBigqueryAdditional) SetDataset(v string)`

SetDataset sets Dataset field to given value.

### HasDataset

`func (o *LoggingBigqueryAdditional) HasDataset() bool`

HasDataset returns a boolean if a field has been set.

### GetTable

`func (o *LoggingBigqueryAdditional) GetTable() string`

GetTable returns the Table field if non-nil, zero value otherwise.

### GetTableOk

`func (o *LoggingBigqueryAdditional) GetTableOk() (*string, bool)`

GetTableOk returns a tuple with the Table field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTable

`func (o *LoggingBigqueryAdditional) SetTable(v string)`

SetTable sets Table field to given value.

### HasTable

`func (o *LoggingBigqueryAdditional) HasTable() bool`

HasTable returns a boolean if a field has been set.

### GetTemplateSuffix

`func (o *LoggingBigqueryAdditional) GetTemplateSuffix() string`

GetTemplateSuffix returns the TemplateSuffix field if non-nil, zero value otherwise.

### GetTemplateSuffixOk

`func (o *LoggingBigqueryAdditional) GetTemplateSuffixOk() (*string, bool)`

GetTemplateSuffixOk returns a tuple with the TemplateSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateSuffix

`func (o *LoggingBigqueryAdditional) SetTemplateSuffix(v string)`

SetTemplateSuffix sets TemplateSuffix field to given value.

### HasTemplateSuffix

`func (o *LoggingBigqueryAdditional) HasTemplateSuffix() bool`

HasTemplateSuffix returns a boolean if a field has been set.

### SetTemplateSuffixNil

`func (o *LoggingBigqueryAdditional) SetTemplateSuffixNil(b bool)`

 SetTemplateSuffixNil sets the value for TemplateSuffix to be an explicit nil

### UnsetTemplateSuffix
`func (o *LoggingBigqueryAdditional) UnsetTemplateSuffix()`

UnsetTemplateSuffix ensures that no value is present for TemplateSuffix, not even an explicit nil
### GetProjectID

`func (o *LoggingBigqueryAdditional) GetProjectID() string`

GetProjectID returns the ProjectID field if non-nil, zero value otherwise.

### GetProjectIDOk

`func (o *LoggingBigqueryAdditional) GetProjectIDOk() (*string, bool)`

GetProjectIDOk returns a tuple with the ProjectID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectID

`func (o *LoggingBigqueryAdditional) SetProjectID(v string)`

SetProjectID sets ProjectID field to given value.

### HasProjectID

`func (o *LoggingBigqueryAdditional) HasProjectID() bool`

HasProjectID returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
