# LoggingDatadogResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name for the real-time logging configuration. | [optional] 
**Placement** | Pointer to **NullableString** | Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.  | [optional] 
**ResponseCondition** | Pointer to **NullableString** | The name of an existing condition in the configured endpoint, or leave blank to always execute. | [optional] 
**Format** | Pointer to **string** | A Fastly [log format string](https://www.fastly.com/documentation/guides/integrations/streaming-logs/custom-log-formats/). Must produce valid JSON that Datadog can ingest.  | [optional] [default to "{\"ddsource\":\"fastly\",\"service\":\"%{req.service_id}V\",\"date\":\"%{begin:%Y-%m-%dT%H:%M:%S%Z}t\",\"time_start\":\"%{begin:%Y-%m-%dT%H:%M:%S%Z}t\",\"time_end\":\"%{end:%Y-%m-%dT%H:%M:%S%Z}t\",\"http\":{\"request_time_ms\":\"%D\",\"method\":\"%m\",\"url\":\"%{json.escape(req.url)}V\",\"useragent\":\"%{User-Agent}i\",\"referer\":\"%{Referer}i\",\"protocol\":\"%H\",\"request_x_forwarded_for\":\"%{X-Forwarded-For}i\",\"status_code\":\"%s\"},\"network\":{\"client\":{\"ip\":\"%h\",\"name\":\"%{client.as.name}V\",\"number\":\"%{client.as.number}V\",\"connection_speed\":\"%{client.geo.conn_speed}V\"},\"destination\":{\"ip\":\"%A\"},\"geoip\":{\"geo_city\":\"%{client.geo.city.utf8}V\",\"geo_country_code\":\"%{client.geo.country_code}V\",\"geo_continent_code\":\"%{client.geo.continent_code}V\",\"geo_region\":\"%{client.geo.region}V\"},\"bytes_written\":\"%B\",\"bytes_read\":\"%{req.body_bytes_read}V\"},\"host\":\"%{Fastly-Orig-Host}i\",\"origin_host\":\"%v\",\"is_ipv6\":\"%{if(req.is_ipv6, \\\"true\\\", \\\"false\\\")}V\",\"is_tls\":\"%{if(req.is_ssl, \\\"true\\\", \\\"false\\\")}V\",\"tls_client_protocol\":\"%{json.escape(tls.client.protocol)}V\",\"tls_client_servername\":\"%{json.escape(tls.client.servername)}V\",\"tls_client_cipher\":\"%{json.escape(tls.client.cipher)}V\",\"tls_client_cipher_sha\":\"%{json.escape(tls.client.ciphers_sha)}V\",\"tls_client_tlsexts_sha\":\"%{json.escape(tls.client.tlsexts_sha)}V\",\"is_h2\":\"%{if(fastly_info.is_h2, \\\"true\\\", \\\"false\\\")}V\",\"is_h2_push\":\"%{if(fastly_info.h2.is_push, \\\"true\\\", \\\"false\\\")}V\",\"h2_stream_id\":\"%{fastly_info.h2.stream_id}V\",\"request_accept_content\":\"%{Accept}i\",\"request_accept_language\":\"%{Accept-Language}i\",\"request_accept_encoding\":\"%{Accept-Encoding}i\",\"request_accept_charset\":\"%{Accept-Charset}i\",\"request_connection\":\"%{Connection}i\",\"request_dnt\":\"%{DNT}i\",\"request_forwarded\":\"%{Forwarded}i\",\"request_via\":\"%{Via}i\",\"request_cache_control\":\"%{Cache-Control}i\",\"request_x_requested_with\":\"%{X-Requested-With}i\",\"request_x_att_device_id\":\"%{X-ATT-Device-ID}i\",\"content_type\":\"%{Content-Type}o\",\"is_cacheable\":\"%{if(fastly_info.state~\\\"^(HIT|MISS)$\\\", \\\"true\\\", \\\"false\\\")}V\",\"response_age\":\"%{Age}o\",\"response_cache_control\":\"%{Cache-Control}o\",\"response_expires\":\"%{Expires}o\",\"response_last_modified\":\"%{Last-Modified}o\",\"response_tsv\":\"%{TSV}o\",\"server_datacenter\":\"%{server.datacenter}V\",\"req_header_size\":\"%{req.header_bytes_read}V\",\"resp_header_size\":\"%{resp.header_bytes_written}V\",\"socket_cwnd\":\"%{client.socket.cwnd}V\",\"socket_nexthop\":\"%{client.socket.nexthop}V\",\"socket_tcpi_rcv_mss\":\"%{client.socket.tcpi_rcv_mss}V\",\"socket_tcpi_snd_mss\":\"%{client.socket.tcpi_snd_mss}V\",\"socket_tcpi_rtt\":\"%{client.socket.tcpi_rtt}V\",\"socket_tcpi_rttvar\":\"%{client.socket.tcpi_rttvar}V\",\"socket_tcpi_rcv_rtt\":\"%{client.socket.tcpi_rcv_rtt}V\",\"socket_tcpi_rcv_space\":\"%{client.socket.tcpi_rcv_space}V\",\"socket_tcpi_last_data_sent\":\"%{client.socket.tcpi_last_data_sent}V\",\"socket_tcpi_total_retrans\":\"%{client.socket.tcpi_total_retrans}V\",\"socket_tcpi_delta_retrans\":\"%{client.socket.tcpi_delta_retrans}V\",\"socket_ploss\":\"%{client.socket.ploss}V\"}"]
**LogProcessingRegion** | Pointer to **string** | The geographic region where the logs will be processed before streaming. Valid values are `us`, `eu`, and `none` for global. | [optional] [default to "none"]
**FormatVersion** | Pointer to **string** | The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.  | [optional] [default to "2"]
**Region** | Pointer to **string** | The region that log data will be sent to. | [optional] [default to "US"]
**Token** | Pointer to **string** | The API key from your Datadog account. Required. | [optional] 
**CreatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**DeletedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**UpdatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**ServiceID** | Pointer to **string** |  | [optional] [readonly] 
**Version** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewLoggingDatadogResponse

`func NewLoggingDatadogResponse() *LoggingDatadogResponse`

NewLoggingDatadogResponse instantiates a new LoggingDatadogResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoggingDatadogResponseWithDefaults

`func NewLoggingDatadogResponseWithDefaults() *LoggingDatadogResponse`

NewLoggingDatadogResponseWithDefaults instantiates a new LoggingDatadogResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *LoggingDatadogResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LoggingDatadogResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LoggingDatadogResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LoggingDatadogResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlacement

`func (o *LoggingDatadogResponse) GetPlacement() string`

GetPlacement returns the Placement field if non-nil, zero value otherwise.

### GetPlacementOk

`func (o *LoggingDatadogResponse) GetPlacementOk() (*string, bool)`

GetPlacementOk returns a tuple with the Placement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacement

`func (o *LoggingDatadogResponse) SetPlacement(v string)`

SetPlacement sets Placement field to given value.

### HasPlacement

`func (o *LoggingDatadogResponse) HasPlacement() bool`

HasPlacement returns a boolean if a field has been set.

### SetPlacementNil

`func (o *LoggingDatadogResponse) SetPlacementNil(b bool)`

 SetPlacementNil sets the value for Placement to be an explicit nil

### UnsetPlacement
`func (o *LoggingDatadogResponse) UnsetPlacement()`

UnsetPlacement ensures that no value is present for Placement, not even an explicit nil
### GetResponseCondition

`func (o *LoggingDatadogResponse) GetResponseCondition() string`

GetResponseCondition returns the ResponseCondition field if non-nil, zero value otherwise.

### GetResponseConditionOk

`func (o *LoggingDatadogResponse) GetResponseConditionOk() (*string, bool)`

GetResponseConditionOk returns a tuple with the ResponseCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseCondition

`func (o *LoggingDatadogResponse) SetResponseCondition(v string)`

SetResponseCondition sets ResponseCondition field to given value.

### HasResponseCondition

`func (o *LoggingDatadogResponse) HasResponseCondition() bool`

HasResponseCondition returns a boolean if a field has been set.

### SetResponseConditionNil

`func (o *LoggingDatadogResponse) SetResponseConditionNil(b bool)`

 SetResponseConditionNil sets the value for ResponseCondition to be an explicit nil

### UnsetResponseCondition
`func (o *LoggingDatadogResponse) UnsetResponseCondition()`

UnsetResponseCondition ensures that no value is present for ResponseCondition, not even an explicit nil
### GetFormat

`func (o *LoggingDatadogResponse) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *LoggingDatadogResponse) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *LoggingDatadogResponse) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *LoggingDatadogResponse) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetLogProcessingRegion

`func (o *LoggingDatadogResponse) GetLogProcessingRegion() string`

GetLogProcessingRegion returns the LogProcessingRegion field if non-nil, zero value otherwise.

### GetLogProcessingRegionOk

`func (o *LoggingDatadogResponse) GetLogProcessingRegionOk() (*string, bool)`

GetLogProcessingRegionOk returns a tuple with the LogProcessingRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogProcessingRegion

`func (o *LoggingDatadogResponse) SetLogProcessingRegion(v string)`

SetLogProcessingRegion sets LogProcessingRegion field to given value.

### HasLogProcessingRegion

`func (o *LoggingDatadogResponse) HasLogProcessingRegion() bool`

HasLogProcessingRegion returns a boolean if a field has been set.

### GetFormatVersion

`func (o *LoggingDatadogResponse) GetFormatVersion() string`

GetFormatVersion returns the FormatVersion field if non-nil, zero value otherwise.

### GetFormatVersionOk

`func (o *LoggingDatadogResponse) GetFormatVersionOk() (*string, bool)`

GetFormatVersionOk returns a tuple with the FormatVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatVersion

`func (o *LoggingDatadogResponse) SetFormatVersion(v string)`

SetFormatVersion sets FormatVersion field to given value.

### HasFormatVersion

`func (o *LoggingDatadogResponse) HasFormatVersion() bool`

HasFormatVersion returns a boolean if a field has been set.

### GetRegion

`func (o *LoggingDatadogResponse) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *LoggingDatadogResponse) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *LoggingDatadogResponse) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *LoggingDatadogResponse) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetToken

`func (o *LoggingDatadogResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *LoggingDatadogResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *LoggingDatadogResponse) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *LoggingDatadogResponse) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetCreatedAt

`func (o *LoggingDatadogResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LoggingDatadogResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LoggingDatadogResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *LoggingDatadogResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *LoggingDatadogResponse) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *LoggingDatadogResponse) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetDeletedAt

`func (o *LoggingDatadogResponse) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *LoggingDatadogResponse) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *LoggingDatadogResponse) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *LoggingDatadogResponse) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *LoggingDatadogResponse) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *LoggingDatadogResponse) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetUpdatedAt

`func (o *LoggingDatadogResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *LoggingDatadogResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *LoggingDatadogResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *LoggingDatadogResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *LoggingDatadogResponse) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *LoggingDatadogResponse) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetServiceID

`func (o *LoggingDatadogResponse) GetServiceID() string`

GetServiceID returns the ServiceID field if non-nil, zero value otherwise.

### GetServiceIDOk

`func (o *LoggingDatadogResponse) GetServiceIDOk() (*string, bool)`

GetServiceIDOk returns a tuple with the ServiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceID

`func (o *LoggingDatadogResponse) SetServiceID(v string)`

SetServiceID sets ServiceID field to given value.

### HasServiceID

`func (o *LoggingDatadogResponse) HasServiceID() bool`

HasServiceID returns a boolean if a field has been set.

### GetVersion

`func (o *LoggingDatadogResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *LoggingDatadogResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *LoggingDatadogResponse) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *LoggingDatadogResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
