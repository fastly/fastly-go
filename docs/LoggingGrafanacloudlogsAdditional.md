# LoggingGrafanacloudlogsAdditional

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Format** | Pointer to **string** | A Fastly [log format string](https://www.fastly.com/documentation/guides/integrations/streaming-logs/custom-log-formats/). | [optional] 
**User** | Pointer to **string** | The Grafana Cloud Logs Dataset you want to log to. | [optional] 
**Url** | Pointer to **string** | The URL of the Loki instance in your Grafana stack. | [optional] 
**Token** | Pointer to **string** | The Grafana Access Policy token with `logs:write` access scoped to your Loki instance. | [optional] 
**Index** | Pointer to **string** | The Stream Labels, a JSON string used to identify the stream. | [optional] 

## Methods

### NewLoggingGrafanacloudlogsAdditional

`func NewLoggingGrafanacloudlogsAdditional() *LoggingGrafanacloudlogsAdditional`

NewLoggingGrafanacloudlogsAdditional instantiates a new LoggingGrafanacloudlogsAdditional object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoggingGrafanacloudlogsAdditionalWithDefaults

`func NewLoggingGrafanacloudlogsAdditionalWithDefaults() *LoggingGrafanacloudlogsAdditional`

NewLoggingGrafanacloudlogsAdditionalWithDefaults instantiates a new LoggingGrafanacloudlogsAdditional object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormat

`func (o *LoggingGrafanacloudlogsAdditional) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *LoggingGrafanacloudlogsAdditional) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *LoggingGrafanacloudlogsAdditional) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *LoggingGrafanacloudlogsAdditional) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetUser

`func (o *LoggingGrafanacloudlogsAdditional) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *LoggingGrafanacloudlogsAdditional) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *LoggingGrafanacloudlogsAdditional) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *LoggingGrafanacloudlogsAdditional) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetUrl

`func (o *LoggingGrafanacloudlogsAdditional) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *LoggingGrafanacloudlogsAdditional) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *LoggingGrafanacloudlogsAdditional) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *LoggingGrafanacloudlogsAdditional) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetToken

`func (o *LoggingGrafanacloudlogsAdditional) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *LoggingGrafanacloudlogsAdditional) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *LoggingGrafanacloudlogsAdditional) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *LoggingGrafanacloudlogsAdditional) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetIndex

`func (o *LoggingGrafanacloudlogsAdditional) GetIndex() string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *LoggingGrafanacloudlogsAdditional) GetIndexOk() (*string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *LoggingGrafanacloudlogsAdditional) SetIndex(v string)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *LoggingGrafanacloudlogsAdditional) HasIndex() bool`

HasIndex returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


