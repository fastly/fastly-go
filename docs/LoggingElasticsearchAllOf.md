# LoggingElasticsearchAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | Pointer to **string** | The name of the Elasticsearch index to send documents (logs) to. The index must follow the Elasticsearch [index format rules](https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-create-index.html). We support [strftime](https://www.man7.org/linux/man-pages/man3/strftime.3.html) interpolated variables inside braces prefixed with a pound symbol. For example, `#{%F}` will interpolate as `YYYY-MM-DD` with today&#39;s date. | [optional] 
**URL** | Pointer to **string** | The URL to stream logs to. Must use HTTPS. | [optional] 
**Pipeline** | Pointer to **NullableString** | The ID of the Elasticsearch ingest pipeline to apply pre-process transformations to before indexing. Learn more about creating a pipeline in the [Elasticsearch docs](https://www.elastic.co/guide/en/elasticsearch/reference/current/ingest.html). | [optional] 
**User** | Pointer to **NullableString** | Basic Auth username. | [optional] 
**Password** | Pointer to **NullableString** | Basic Auth password. | [optional] 
**Format** | Pointer to **string** | A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats). Must produce valid JSON that Elasticsearch can ingest. | [optional] 

## Methods

### NewLoggingElasticsearchAllOf

`func NewLoggingElasticsearchAllOf() *LoggingElasticsearchAllOf`

NewLoggingElasticsearchAllOf instantiates a new LoggingElasticsearchAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoggingElasticsearchAllOfWithDefaults

`func NewLoggingElasticsearchAllOfWithDefaults() *LoggingElasticsearchAllOf`

NewLoggingElasticsearchAllOfWithDefaults instantiates a new LoggingElasticsearchAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *LoggingElasticsearchAllOf) GetIndex() string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *LoggingElasticsearchAllOf) GetIndexOk() (*string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *LoggingElasticsearchAllOf) SetIndex(v string)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *LoggingElasticsearchAllOf) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetURL

`func (o *LoggingElasticsearchAllOf) GetURL() string`

GetURL returns the URL field if non-nil, zero value otherwise.

### GetURLOk

`func (o *LoggingElasticsearchAllOf) GetURLOk() (*string, bool)`

GetURLOk returns a tuple with the URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetURL

`func (o *LoggingElasticsearchAllOf) SetURL(v string)`

SetURL sets URL field to given value.

### HasURL

`func (o *LoggingElasticsearchAllOf) HasURL() bool`

HasURL returns a boolean if a field has been set.

### GetPipeline

`func (o *LoggingElasticsearchAllOf) GetPipeline() string`

GetPipeline returns the Pipeline field if non-nil, zero value otherwise.

### GetPipelineOk

`func (o *LoggingElasticsearchAllOf) GetPipelineOk() (*string, bool)`

GetPipelineOk returns a tuple with the Pipeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipeline

`func (o *LoggingElasticsearchAllOf) SetPipeline(v string)`

SetPipeline sets Pipeline field to given value.

### HasPipeline

`func (o *LoggingElasticsearchAllOf) HasPipeline() bool`

HasPipeline returns a boolean if a field has been set.

### SetPipelineNil

`func (o *LoggingElasticsearchAllOf) SetPipelineNil(b bool)`

 SetPipelineNil sets the value for Pipeline to be an explicit nil

### UnsetPipeline
`func (o *LoggingElasticsearchAllOf) UnsetPipeline()`

UnsetPipeline ensures that no value is present for Pipeline, not even an explicit nil
### GetUser

`func (o *LoggingElasticsearchAllOf) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *LoggingElasticsearchAllOf) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *LoggingElasticsearchAllOf) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *LoggingElasticsearchAllOf) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *LoggingElasticsearchAllOf) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *LoggingElasticsearchAllOf) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetPassword

`func (o *LoggingElasticsearchAllOf) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *LoggingElasticsearchAllOf) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *LoggingElasticsearchAllOf) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *LoggingElasticsearchAllOf) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *LoggingElasticsearchAllOf) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *LoggingElasticsearchAllOf) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetFormat

`func (o *LoggingElasticsearchAllOf) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *LoggingElasticsearchAllOf) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *LoggingElasticsearchAllOf) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *LoggingElasticsearchAllOf) HasFormat() bool`

HasFormat returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
