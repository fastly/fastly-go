# LoggingBigqueryAllOf

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

### NewLoggingBigqueryAllOf

`func NewLoggingBigqueryAllOf() *LoggingBigqueryAllOf`

NewLoggingBigqueryAllOf instantiates a new LoggingBigqueryAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoggingBigqueryAllOfWithDefaults

`func NewLoggingBigqueryAllOfWithDefaults() *LoggingBigqueryAllOf`

NewLoggingBigqueryAllOfWithDefaults instantiates a new LoggingBigqueryAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *LoggingBigqueryAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LoggingBigqueryAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LoggingBigqueryAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LoggingBigqueryAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFormat

`func (o *LoggingBigqueryAllOf) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *LoggingBigqueryAllOf) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *LoggingBigqueryAllOf) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *LoggingBigqueryAllOf) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetDataset

`func (o *LoggingBigqueryAllOf) GetDataset() string`

GetDataset returns the Dataset field if non-nil, zero value otherwise.

### GetDatasetOk

`func (o *LoggingBigqueryAllOf) GetDatasetOk() (*string, bool)`

GetDatasetOk returns a tuple with the Dataset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataset

`func (o *LoggingBigqueryAllOf) SetDataset(v string)`

SetDataset sets Dataset field to given value.

### HasDataset

`func (o *LoggingBigqueryAllOf) HasDataset() bool`

HasDataset returns a boolean if a field has been set.

### GetTable

`func (o *LoggingBigqueryAllOf) GetTable() string`

GetTable returns the Table field if non-nil, zero value otherwise.

### GetTableOk

`func (o *LoggingBigqueryAllOf) GetTableOk() (*string, bool)`

GetTableOk returns a tuple with the Table field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTable

`func (o *LoggingBigqueryAllOf) SetTable(v string)`

SetTable sets Table field to given value.

### HasTable

`func (o *LoggingBigqueryAllOf) HasTable() bool`

HasTable returns a boolean if a field has been set.

### GetTemplateSuffix

`func (o *LoggingBigqueryAllOf) GetTemplateSuffix() string`

GetTemplateSuffix returns the TemplateSuffix field if non-nil, zero value otherwise.

### GetTemplateSuffixOk

`func (o *LoggingBigqueryAllOf) GetTemplateSuffixOk() (*string, bool)`

GetTemplateSuffixOk returns a tuple with the TemplateSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateSuffix

`func (o *LoggingBigqueryAllOf) SetTemplateSuffix(v string)`

SetTemplateSuffix sets TemplateSuffix field to given value.

### HasTemplateSuffix

`func (o *LoggingBigqueryAllOf) HasTemplateSuffix() bool`

HasTemplateSuffix returns a boolean if a field has been set.

### SetTemplateSuffixNil

`func (o *LoggingBigqueryAllOf) SetTemplateSuffixNil(b bool)`

 SetTemplateSuffixNil sets the value for TemplateSuffix to be an explicit nil

### UnsetTemplateSuffix
`func (o *LoggingBigqueryAllOf) UnsetTemplateSuffix()`

UnsetTemplateSuffix ensures that no value is present for TemplateSuffix, not even an explicit nil
### GetProjectID

`func (o *LoggingBigqueryAllOf) GetProjectID() string`

GetProjectID returns the ProjectID field if non-nil, zero value otherwise.

### GetProjectIDOk

`func (o *LoggingBigqueryAllOf) GetProjectIDOk() (*string, bool)`

GetProjectIDOk returns a tuple with the ProjectID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectID

`func (o *LoggingBigqueryAllOf) SetProjectID(v string)`

SetProjectID sets ProjectID field to given value.

### HasProjectID

`func (o *LoggingBigqueryAllOf) HasProjectID() bool`

HasProjectID returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
