# LoggingGooglePubsubAdditional

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Topic** | Pointer to **string** | The Google Cloud Pub/Sub topic to which logs will be published. Required. | [optional] 
**Format** | Pointer to **string** | A Fastly [log format string](https://www.fastly.com/documentation/guides/integrations/streaming-logs/custom-log-formats/). | [optional] [default to "%h %l %u %t \"%r\" %&gt;s %b"]
**ProjectID** | Pointer to **string** | Your Google Cloud Platform project ID. Required | [optional] 

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

### GetProjectID

`func (o *LoggingGooglePubsubAdditional) GetProjectID() string`

GetProjectID returns the ProjectID field if non-nil, zero value otherwise.

### GetProjectIDOk

`func (o *LoggingGooglePubsubAdditional) GetProjectIDOk() (*string, bool)`

GetProjectIDOk returns a tuple with the ProjectID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectID

`func (o *LoggingGooglePubsubAdditional) SetProjectID(v string)`

SetProjectID sets ProjectID field to given value.

### HasProjectID

`func (o *LoggingGooglePubsubAdditional) HasProjectID() bool`

HasProjectID returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
