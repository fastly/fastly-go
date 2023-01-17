# LoggingDatadog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name for the real-time logging configuration. | [optional] 
**Placement** | Pointer to **NullableString** | Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.  | [optional] 
**FormatVersion** | Pointer to **int32** | The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.  | [optional] [default to 2]
**ResponseCondition** | Pointer to **NullableString** | The name of an existing condition in the configured endpoint, or leave blank to always execute. | [optional] 
**Format** | Pointer to **string** | A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats). Must produce valid JSON that Datadog can ingest.  | [optional] [default to "{\"ddsource\":\"fastly\",\"service\":\"%{req.service_id}V\",\"date\":\"%{begin:%Y-%m-%dT%H:%M:%S%Z}t\",\"time_start\":\"%{begin:%Y-%m-%dT%H:%M:%S%Z}t\",\"time_end\":\"%{end:%Y-%m-%dT%H:%M:%S%Z}t\",\"http\":{\"request_time_ms\":\"%D\",\"method\":\"%m\",\"url\":\"%{json.escape(req.url)}V\",\"useragent\":\"%{User-Agent}i\",\"referer\":\"%{Referer}i\",\"protocol\":\"%H\",\"request_x_forwarded_for\":\"%{X-Forwarded-For}i\",\"status_code\":\"%s\"},\"network\":{\"client\":{\"ip\":\"%h\",\"name\":\"%{client.as.name}V\",\"number\":\"%{client.as.number}V\",\"connection_speed\":\"%{client.geo.conn_speed}V\"},\"destination\":{\"ip\":\"%A\"},\"geoip\":{\"geo_city\":\"%{client.geo.city.utf8}V\",\"geo_country_code\":\"%{client.geo.country_code}V\",\"geo_continent_code\":\"%{client.geo.continent_code}V\",\"geo_region\":\"%{client.geo.region}V\"},\"bytes_written\":\"%B\",\"bytes_read\":\"%{req.body_bytes_read}V\"},\"host\":\"%{Fastly-Orig-Host}i\",\"origin_host\":\"%v\",\"is_ipv6\":\"%{if(req.is_ipv6, \\\"true\\\", \\\"false\\\")}V\",\"is_tls\":\"%{if(req.is_ssl, \\\"true\\\", \\\"false\\\")}V\",\"tls_client_protocol\":\"%{json.escape(tls.client.protocol)}V\",\"tls_client_servername\":\"%{json.escape(tls.client.servername)}V\",\"tls_client_cipher\":\"%{json.escape(tls.client.cipher)}V\",\"tls_client_cipher_sha\":\"%{json.escape(tls.client.ciphers_sha)}V\",\"tls_client_tlsexts_sha\":\"%{json.escape(tls.client.tlsexts_sha)}V\",\"is_h2\":\"%{if(fastly_info.is_h2, \\\"true\\\", \\\"false\\\")}V\",\"is_h2_push\":\"%{if(fastly_info.h2.is_push, \\\"true\\\", \\\"false\\\")}V\",\"h2_stream_id\":\"%{fastly_info.h2.stream_id}V\",\"request_accept_content\":\"%{Accept}i\",\"request_accept_language\":\"%{Accept-Language}i\",\"request_accept_encoding\":\"%{Accept-Encoding}i\",\"request_accept_charset\":\"%{Accept-Charset}i\",\"request_connection\":\"%{Connection}i\",\"request_dnt\":\"%{DNT}i\",\"request_forwarded\":\"%{Forwarded}i\",\"request_via\":\"%{Via}i\",\"request_cache_control\":\"%{Cache-Control}i\",\"request_x_requested_with\":\"%{X-Requested-With}i\",\"request_x_att_device_id\":\"%{X-ATT-Device-ID}i\",\"content_type\":\"%{Content-Type}o\",\"is_cacheable\":\"%{if(fastly_info.state~\\\"^(HIT|MISS)$\\\", \\\"true\\\", \\\"false\\\")}V\",\"response_age\":\"%{Age}o\",\"response_cache_control\":\"%{Cache-Control}o\",\"response_expires\":\"%{Expires}o\",\"response_last_modified\":\"%{Last-Modified}o\",\"response_tsv\":\"%{TSV}o\",\"server_datacenter\":\"%{server.datacenter}V\",\"req_header_size\":\"%{req.header_bytes_read}V\",\"resp_header_size\":\"%{resp.header_bytes_written}V\",\"socket_cwnd\":\"%{client.socket.cwnd}V\",\"socket_nexthop\":\"%{client.socket.nexthop}V\",\"socket_tcpi_rcv_mss\":\"%{client.socket.tcpi_rcv_mss}V\",\"socket_tcpi_snd_mss\":\"%{client.socket.tcpi_snd_mss}V\",\"socket_tcpi_rtt\":\"%{client.socket.tcpi_rtt}V\",\"socket_tcpi_rttvar\":\"%{client.socket.tcpi_rttvar}V\",\"socket_tcpi_rcv_rtt\":\"%{client.socket.tcpi_rcv_rtt}V\",\"socket_tcpi_rcv_space\":\"%{client.socket.tcpi_rcv_space}V\",\"socket_tcpi_last_data_sent\":\"%{client.socket.tcpi_last_data_sent}V\",\"socket_tcpi_total_retrans\":\"%{client.socket.tcpi_total_retrans}V\",\"socket_tcpi_delta_retrans\":\"%{client.socket.tcpi_delta_retrans}V\",\"socket_ploss\":\"%{client.socket.ploss}V\"}"]
**Region** | Pointer to **string** | The region that log data will be sent to. | [optional] [default to "US"]
**Token** | Pointer to **string** | The API key from your Datadog account. Required. | [optional] 

## Methods

### NewLoggingDatadog

`func NewLoggingDatadog() *LoggingDatadog`

NewLoggingDatadog instantiates a new LoggingDatadog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoggingDatadogWithDefaults

`func NewLoggingDatadogWithDefaults() *LoggingDatadog`

NewLoggingDatadogWithDefaults instantiates a new LoggingDatadog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *LoggingDatadog) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LoggingDatadog) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LoggingDatadog) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LoggingDatadog) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlacement

`func (o *LoggingDatadog) GetPlacement() string`

GetPlacement returns the Placement field if non-nil, zero value otherwise.

### GetPlacementOk

`func (o *LoggingDatadog) GetPlacementOk() (*string, bool)`

GetPlacementOk returns a tuple with the Placement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacement

`func (o *LoggingDatadog) SetPlacement(v string)`

SetPlacement sets Placement field to given value.

### HasPlacement

`func (o *LoggingDatadog) HasPlacement() bool`

HasPlacement returns a boolean if a field has been set.

### SetPlacementNil

`func (o *LoggingDatadog) SetPlacementNil(b bool)`

 SetPlacementNil sets the value for Placement to be an explicit nil

### UnsetPlacement
`func (o *LoggingDatadog) UnsetPlacement()`

UnsetPlacement ensures that no value is present for Placement, not even an explicit nil
### GetFormatVersion

`func (o *LoggingDatadog) GetFormatVersion() int32`

GetFormatVersion returns the FormatVersion field if non-nil, zero value otherwise.

### GetFormatVersionOk

`func (o *LoggingDatadog) GetFormatVersionOk() (*int32, bool)`

GetFormatVersionOk returns a tuple with the FormatVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatVersion

`func (o *LoggingDatadog) SetFormatVersion(v int32)`

SetFormatVersion sets FormatVersion field to given value.

### HasFormatVersion

`func (o *LoggingDatadog) HasFormatVersion() bool`

HasFormatVersion returns a boolean if a field has been set.

### GetResponseCondition

`func (o *LoggingDatadog) GetResponseCondition() string`

GetResponseCondition returns the ResponseCondition field if non-nil, zero value otherwise.

### GetResponseConditionOk

`func (o *LoggingDatadog) GetResponseConditionOk() (*string, bool)`

GetResponseConditionOk returns a tuple with the ResponseCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseCondition

`func (o *LoggingDatadog) SetResponseCondition(v string)`

SetResponseCondition sets ResponseCondition field to given value.

### HasResponseCondition

`func (o *LoggingDatadog) HasResponseCondition() bool`

HasResponseCondition returns a boolean if a field has been set.

### SetResponseConditionNil

`func (o *LoggingDatadog) SetResponseConditionNil(b bool)`

 SetResponseConditionNil sets the value for ResponseCondition to be an explicit nil

### UnsetResponseCondition
`func (o *LoggingDatadog) UnsetResponseCondition()`

UnsetResponseCondition ensures that no value is present for ResponseCondition, not even an explicit nil
### GetFormat

`func (o *LoggingDatadog) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *LoggingDatadog) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *LoggingDatadog) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *LoggingDatadog) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetRegion

`func (o *LoggingDatadog) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *LoggingDatadog) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *LoggingDatadog) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *LoggingDatadog) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetToken

`func (o *LoggingDatadog) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *LoggingDatadog) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *LoggingDatadog) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *LoggingDatadog) HasToken() bool`

HasToken returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
