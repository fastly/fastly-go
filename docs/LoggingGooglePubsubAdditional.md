# LoggingGooglePubsubAdditional

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Topic** | Pointer to **string** | The Google Cloud Pub/Sub topic to which logs will be published. Required. | [optional] 
**Format** | Pointer to **string** | A Fastly [log format string](https://www.fastly.com/documentation/guides/integrations/streaming-logs/custom-log-formats/). | [optional] [default to "%h %l %u %t \"%r\" %&gt;s %b"]
**ProjectId** | Pointer to **string** | Your Google Cloud Platform project ID. Required | [optional] 

## Methods

### NewLoggingGooglePubsubAdditional

`func NewLoggingGooglePubsubAdditional() *LoggingGooglePubsubAdditional`

NewLoggingGooglePubsubAdditional instantiates a new LoggingGooglePubsubAdditional object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoggingGooglePubsubAdditionalWithDefaults

`func NewLoggingGooglePubsubAdditionalWithDefaults() *LoggingGooglePubsubAdditional`

NewLoggingGooglePubsubAdditionalWithDefaults instantiates a new LoggingGooglePubsubAdditional object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTopic

`func (o *LoggingGooglePubsubAdditional) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *LoggingGooglePubsubAdditional) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *LoggingGooglePubsubAdditional) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *LoggingGooglePubsubAdditional) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### GetFormat

`func (o *LoggingGooglePubsubAdditional) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *LoggingGooglePubsubAdditional) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *LoggingGooglePubsubAdditional) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *LoggingGooglePubsubAdditional) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetProjectId

`func (o *LoggingGooglePubsubAdditional) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *LoggingGooglePubsubAdditional) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *LoggingGooglePubsubAdditional) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *LoggingGooglePubsubAdditional) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


