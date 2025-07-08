# LoggingElasticsearchAdditional

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | Pointer to **string** | The name of the Elasticsearch index to send documents (logs) to. The index must follow the Elasticsearch [index format rules](https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-create-index.html). We support [strftime](https://www.man7.org/linux/man-pages/man3/strftime.3.html) interpolated variables inside braces prefixed with a pound symbol. For example, `#{%F}` will interpolate as `YYYY-MM-DD` with today&#39;s date. | [optional] 
**URL** | Pointer to **string** | The URL to stream logs to. Must use HTTPS. | [optional] 
**Pipeline** | Pointer to **NullableString** | The ID of the Elasticsearch ingest pipeline to apply pre-process transformations to before indexing. Learn more about creating a pipeline in the [Elasticsearch docs](https://www.elastic.co/guide/en/elasticsearch/reference/current/ingest.html). | [optional] 
**User** | Pointer to **NullableString** | Basic Auth username. | [optional] 
**Password** | Pointer to **NullableString** | Basic Auth password. | [optional] 
**Format** | Pointer to **string** | A Fastly [log format string](https://www.fastly.com/documentation/guides/integrations/streaming-logs/custom-log-formats/). Must produce valid JSON that Elasticsearch can ingest. | [optional] 

## Methods

### NewLoggingElasticsearchAdditional

`func NewLoggingElasticsearchAdditional() *LoggingElasticsearchAdditional`

NewLoggingElasticsearchAdditional instantiates a new LoggingElasticsearchAdditional object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoggingElasticsearchAdditionalWithDefaults

`func NewLoggingElasticsearchAdditionalWithDefaults() *LoggingElasticsearchAdditional`

NewLoggingElasticsearchAdditionalWithDefaults instantiates a new LoggingElasticsearchAdditional object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *LoggingElasticsearchAdditional) GetIndex() string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *LoggingElasticsearchAdditional) GetIndexOk() (*string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *LoggingElasticsearchAdditional) SetIndex(v string)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *LoggingElasticsearchAdditional) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetURL

`func (o *LoggingElasticsearchAdditional) GetURL() string`

GetURL returns the URL field if non-nil, zero value otherwise.

### GetURLOk

`func (o *LoggingElasticsearchAdditional) GetURLOk() (*string, bool)`

GetURLOk returns a tuple with the URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetURL

`func (o *LoggingElasticsearchAdditional) SetURL(v string)`

SetURL sets URL field to given value.

### HasURL

`func (o *LoggingElasticsearchAdditional) HasURL() bool`

HasURL returns a boolean if a field has been set.

### GetPipeline

`func (o *LoggingElasticsearchAdditional) GetPipeline() string`

GetPipeline returns the Pipeline field if non-nil, zero value otherwise.

### GetPipelineOk

`func (o *LoggingElasticsearchAdditional) GetPipelineOk() (*string, bool)`

GetPipelineOk returns a tuple with the Pipeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipeline

`func (o *LoggingElasticsearchAdditional) SetPipeline(v string)`

SetPipeline sets Pipeline field to given value.

### HasPipeline

`func (o *LoggingElasticsearchAdditional) HasPipeline() bool`

HasPipeline returns a boolean if a field has been set.

### SetPipelineNil

`func (o *LoggingElasticsearchAdditional) SetPipelineNil(b bool)`

 SetPipelineNil sets the value for Pipeline to be an explicit nil

### UnsetPipeline
`func (o *LoggingElasticsearchAdditional) UnsetPipeline()`

UnsetPipeline ensures that no value is present for Pipeline, not even an explicit nil
### GetUser

`func (o *LoggingElasticsearchAdditional) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *LoggingElasticsearchAdditional) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *LoggingElasticsearchAdditional) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *LoggingElasticsearchAdditional) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *LoggingElasticsearchAdditional) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *LoggingElasticsearchAdditional) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetPassword

`func (o *LoggingElasticsearchAdditional) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *LoggingElasticsearchAdditional) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *LoggingElasticsearchAdditional) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *LoggingElasticsearchAdditional) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *LoggingElasticsearchAdditional) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *LoggingElasticsearchAdditional) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetFormat

`func (o *LoggingElasticsearchAdditional) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *LoggingElasticsearchAdditional) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *LoggingElasticsearchAdditional) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *LoggingElasticsearchAdditional) HasFormat() bool`

HasFormat returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
