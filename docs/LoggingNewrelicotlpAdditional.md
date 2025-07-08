# LoggingNewrelicotlpAdditional

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Format** | Pointer to **string** | A Fastly [log format string](https://www.fastly.com/documentation/guides/integrations/streaming-logs/custom-log-formats/). | [optional] [default to "{\"timestamp\":\"%{begin:%Y-%m-%dT%H:%M:%S}t\",\"time_elapsed\":\"%{time.elapsed.usec}V\",\"is_tls\":\"%{if(req.is_ssl, \\\"true\\\", \\\"false\\\")}V\",\"client_ip\":\"%{req.http.Fastly-Client-IP}V\",\"geo_city\":\"%{client.geo.city}V\",\"geo_country_code\":\"%{client.geo.country_code}V\",\"request\":\"%{req.request}V\",\"host\":\"%{req.http.Fastly-Orig-Host}V\",\"url\":\"%{json.escape(req.url)}V\",\"request_referer\":\"%{json.escape(req.http.Referer)}V\",\"request_user_agent\":\"%{json.escape(req.http.User-Agent)}V\",\"request_accept_language\":\"%{json.escape(req.http.Accept-Language)}V\",\"request_accept_charset\":\"%{json.escape(req.http.Accept-Charset)}V\",\"cache_status\":\"%{regsub(fastly_info.state, \\\"^(HIT-(SYNTH)|(HITPASS|HIT|MISS|PASS|ERROR|PIPE)).*\\\", \\\"\\\\2\\\\3\\\") }V\"}"]
**Token** | Pointer to **string** | The Insert API key from the Account page of your New Relic account. Required. | [optional] 
**Region** | Pointer to **string** | The region to which to stream logs. | [optional] [default to "US"]
**URL** | Pointer to **NullableString** | (Optional) URL of the New Relic Trace Observer, if you are using New Relic Infinite Tracing. | [optional] [default to "null"]

## Methods

### NewLoggingNewrelicotlpAdditional

`func NewLoggingNewrelicotlpAdditional() *LoggingNewrelicotlpAdditional`

NewLoggingNewrelicotlpAdditional instantiates a new LoggingNewrelicotlpAdditional object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoggingNewrelicotlpAdditionalWithDefaults

`func NewLoggingNewrelicotlpAdditionalWithDefaults() *LoggingNewrelicotlpAdditional`

NewLoggingNewrelicotlpAdditionalWithDefaults instantiates a new LoggingNewrelicotlpAdditional object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormat

`func (o *LoggingNewrelicotlpAdditional) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *LoggingNewrelicotlpAdditional) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *LoggingNewrelicotlpAdditional) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *LoggingNewrelicotlpAdditional) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetToken

`func (o *LoggingNewrelicotlpAdditional) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *LoggingNewrelicotlpAdditional) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *LoggingNewrelicotlpAdditional) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *LoggingNewrelicotlpAdditional) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetRegion

`func (o *LoggingNewrelicotlpAdditional) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *LoggingNewrelicotlpAdditional) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *LoggingNewrelicotlpAdditional) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *LoggingNewrelicotlpAdditional) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetURL

`func (o *LoggingNewrelicotlpAdditional) GetURL() string`

GetURL returns the URL field if non-nil, zero value otherwise.

### GetURLOk

`func (o *LoggingNewrelicotlpAdditional) GetURLOk() (*string, bool)`

GetURLOk returns a tuple with the URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetURL

`func (o *LoggingNewrelicotlpAdditional) SetURL(v string)`

SetURL sets URL field to given value.

### HasURL

`func (o *LoggingNewrelicotlpAdditional) HasURL() bool`

HasURL returns a boolean if a field has been set.

### SetURLNil

`func (o *LoggingNewrelicotlpAdditional) SetURLNil(b bool)`

 SetURLNil sets the value for URL to be an explicit nil

### UnsetURL
`func (o *LoggingNewrelicotlpAdditional) UnsetURL()`

UnsetURL ensures that no value is present for URL, not even an explicit nil

[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
