# OriginInspectorMeasurements

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Responses** | Pointer to **int32** | Number of responses from origin. | [optional] 
**RespHeaderBytes** | Pointer to **int32** | Number of header bytes from origin. | [optional] 
**RespBodyBytes** | Pointer to **int32** | Number of body bytes from origin. | [optional] 
**Status1xx** | Pointer to **int32** | Number of 1xx \&quot;Informational\&quot; status codes delivered from origin. | [optional] 
**Status2xx** | Pointer to **int32** | Number of 2xx \&quot;Success\&quot; status codes delivered from origin. | [optional] 
**Status3xx** | Pointer to **int32** | Number of 3xx \&quot;Redirection\&quot; codes delivered from origin. | [optional] 
**Status4xx** | Pointer to **int32** | Number of 4xx \&quot;Client Error\&quot; codes delivered from origin. | [optional] 
**Status5xx** | Pointer to **int32** | Number of 5xx \&quot;Server Error\&quot; codes delivered from origin. | [optional] 
**Status200** | Pointer to **int32** | Number of responses received with status code 200 (Success) from origin. | [optional] 
**Status204** | Pointer to **int32** | Number of responses received with status code 204 (No Content) from origin. | [optional] 
**Status206** | Pointer to **int32** | Number of responses received with status code 206 (Partial Content) from origin. | [optional] 
**Status301** | Pointer to **int32** | Number of responses received with status code 301 (Moved Permanently) from origin. | [optional] 
**Status302** | Pointer to **int32** | Number of responses received with status code 302 (Found) from origin. | [optional] 
**Status304** | Pointer to **int32** | Number of responses received with status code 304 (Not Modified) from origin. | [optional] 
**Status400** | Pointer to **int32** | Number of responses received with status code 400 (Bad Request) from origin. | [optional] 
**Status401** | Pointer to **int32** | Number of responses received with status code 401 (Unauthorized) from origin. | [optional] 
**Status403** | Pointer to **int32** | Number of responses received with status code 403 (Forbidden) from origin. | [optional] 
**Status404** | Pointer to **int32** | Number of responses received with status code 404 (Not Found) from origin. | [optional] 
**Status416** | Pointer to **int32** | Number of responses received with status code 416 (Range Not Satisfiable) from origin. | [optional] 
**Status429** | Pointer to **int32** | Number of responses received with status code 429 (Too Many Requests) from origin. | [optional] 
**Status500** | Pointer to **int32** | Number of responses received with status code 500 (Internal Server Error) from origin. | [optional] 
**Status501** | Pointer to **int32** | Number of responses received with status code 501 (Not Implemented) from origin. | [optional] 
**Status502** | Pointer to **int32** | Number of responses received with status code 502 (Bad Gateway) from origin. | [optional] 
**Status503** | Pointer to **int32** | Number of responses received with status code 503 (Service Unavailable) from origin. | [optional] 
**Status504** | Pointer to **int32** | Number of responses received with status code 504 (Gateway Timeout) from origin. | [optional] 
**Status505** | Pointer to **int32** | Number of responses received with status code 505 (HTTP Version Not Supported) from origin. | [optional] 
**Latency0To1ms** | Pointer to **int32** | Number of responses from origin with latency between 0 and 1 millisecond. | [optional] 
**Latency1To5ms** | Pointer to **int32** | Number of responses from origin with latency between 1 and 5 milliseconds. | [optional] 
**Latency5To10ms** | Pointer to **int32** | Number of responses from origin with latency between 5 and 10 milliseconds. | [optional] 
**Latency10To50ms** | Pointer to **int32** | Number of responses from origin with latency between 10 and 50 milliseconds. | [optional] 
**Latency50To100ms** | Pointer to **int32** | Number of responses from origin with latency between 50 and 100 milliseconds. | [optional] 
**Latency100To250ms** | Pointer to **int32** | Number of responses from origin with latency between 100 and 250 milliseconds. | [optional] 
**Latency250To500ms** | Pointer to **int32** | Number of responses from origin with latency between 250 and 500 milliseconds. | [optional] 
**Latency500To1000ms** | Pointer to **int32** | Number of responses from origin with latency between 500 and 1,000 milliseconds. | [optional] 
**Latency1000To5000ms** | Pointer to **int32** | Number of responses from origin with latency between 1,000 and 5,000 milliseconds. | [optional] 
**Latency5000To10000ms** | Pointer to **int32** | Number of responses from origin with latency between 5,000 and 10,000 milliseconds. | [optional] 
**Latency10000To60000ms** | Pointer to **int32** | Number of responses from origin with latency between 10,000 and 60,000 milliseconds. | [optional] 
**Latency60000ms** | Pointer to **int32** | Number of responses from origin with latency of 60,000 milliseconds and above. | [optional] 
**WafResponses** | Pointer to **int32** | Number of responses received for origin requests made by the Fastly WAF. | [optional] 
**WafRespHeaderBytes** | Pointer to **int32** | Number of header bytes received for origin requests made by the Fastly WAF. | [optional] 
**WafRespBodyBytes** | Pointer to **int32** | Number of body bytes received for origin requests made by the Fastly WAF. | [optional] 
**WafStatus1xx** | Pointer to **int32** | Number of 1xx \&quot;Informational\&quot; status codes received for origin requests made by the Fastly WAF. | [optional] 
**WafStatus2xx** | Pointer to **int32** | Number of 2xx \&quot;Success\&quot; status codes received for origin requests made by the Fastly WAF. | [optional] 
**WafStatus3xx** | Pointer to **int32** | Number of 3xx \&quot;Redirection\&quot; codes received for origin requests made by the Fastly WAF. | [optional] 
**WafStatus4xx** | Pointer to **int32** | Number of 4xx \&quot;Client Error\&quot; codes received for origin requests made by the Fastly WAF. | [optional] 
**WafStatus5xx** | Pointer to **int32** | Number of 5xx \&quot;Server Error\&quot; codes received for origin requests made by the Fastly WAF. | [optional] 
**WafStatus200** | Pointer to **int32** | Number of responses received with status code 200 (Success) received for origin requests made by the Fastly WAF. | [optional] 
**WafStatus204** | Pointer to **int32** | Number of responses received with status code 204 (No Content) received for origin requests made by the Fastly WAF. | [optional] 
**WafStatus206** | Pointer to **int32** | Number of responses received with status code 206 (Partial Content) received for origin requests made by the Fastly WAF. | [optional] 
**WafStatus301** | Pointer to **int32** | Number of responses received with status code 301 (Moved Permanently) received for origin requests made by the Fastly WAF. | [optional] 
**WafStatus302** | Pointer to **int32** | Number of responses received with status code 302 (Found) received for origin requests made by the Fastly WAF. | [optional] 
**WafStatus304** | Pointer to **int32** | Number of responses received with status code 304 (Not Modified) received for origin requests made by the Fastly WAF. | [optional] 
**WafStatus400** | Pointer to **int32** | Number of responses received with status code 400 (Bad Request) received for origin requests made by the Fastly WAF. | [optional] 
**WafStatus401** | Pointer to **int32** | Number of responses received with status code 401 (Unauthorized) received for origin requests made by the Fastly WAF. | [optional] 
**WafStatus403** | Pointer to **int32** | Number of responses received with status code 403 (Forbidden) received for origin requests made by the Fastly WAF. | [optional] 
**WafStatus404** | Pointer to **int32** | Number of responses received with status code 404 (Not Found) received for origin requests made by the Fastly WAF. | [optional] 
**WafStatus416** | Pointer to **int32** | Number of responses received with status code 416 (Range Not Satisfiable) received for origin requests made by the Fastly WAF. | [optional] 
**WafStatus429** | Pointer to **int32** | Number of responses received with status code 429 (Too Many Requests) received for origin requests made by the Fastly WAF. | [optional] 
**WafStatus500** | Pointer to **int32** | Number of responses received with status code 500 (Internal Server Error) received for origin requests made by the Fastly WAF. | [optional] 
**WafStatus501** | Pointer to **int32** | Number of responses received with status code 501 (Not Implemented) received for origin requests made by the Fastly WAF. | [optional] 
**WafStatus502** | Pointer to **int32** | Number of responses received with status code 502 (Bad Gateway) received for origin requests made by the Fastly WAF. | [optional] 
**WafStatus503** | Pointer to **int32** | Number of responses received with status code 503 (Service Unavailable) received for origin requests made by the Fastly WAF. | [optional] 
**WafStatus504** | Pointer to **int32** | Number of responses received with status code 504 (Gateway Timeout) received for origin requests made by the Fastly WAF. | [optional] 
**WafStatus505** | Pointer to **int32** | Number of responses received with status code 505 (HTTP Version Not Supported) received for origin requests made by the Fastly WAF. | [optional] 
**WafLatency0To1ms** | Pointer to **int32** | Number of responses with latency between 0 and 1 millisecond received for origin requests made by the Fastly WAF. | [optional] 
**WafLatency1To5ms** | Pointer to **int32** | Number of responses with latency between 1 and 5 milliseconds received for origin requests made by the Fastly WAF. | [optional] 
**WafLatency5To10ms** | Pointer to **int32** | Number of responses with latency between 5 and 10 milliseconds received for origin requests made by the Fastly WAF. | [optional] 
**WafLatency10To50ms** | Pointer to **int32** | Number of responses with latency between 10 and 50 milliseconds received for origin requests made by the Fastly WAF. | [optional] 
**WafLatency50To100ms** | Pointer to **int32** | Number of responses with latency between 50 and 100 milliseconds received for origin requests made by the Fastly WAF. | [optional] 
**WafLatency100To250ms** | Pointer to **int32** | Number of responses with latency between 100 and 250 milliseconds received for origin requests made by the Fastly WAF. | [optional] 
**WafLatency250To500ms** | Pointer to **int32** | Number of responses with latency between 250 and 500 milliseconds received for origin requests made by the Fastly WAF. | [optional] 
**WafLatency500To1000ms** | Pointer to **int32** | Number of responses with latency between 500 and 1,000 milliseconds received for origin requests made by the Fastly WAF. | [optional] 
**WafLatency1000To5000ms** | Pointer to **int32** | Number of responses with latency between 1,000 and 5,000 milliseconds received for origin requests made by the Fastly WAF. | [optional] 
**WafLatency5000To10000ms** | Pointer to **int32** | Number of responses with latency between 5,000 and 10,000 milliseconds received for origin requests made by the Fastly WAF. | [optional] 
**WafLatency10000To60000ms** | Pointer to **int32** | Number of responses with latency between 10,000 and 60,000 milliseconds received for origin requests made by the Fastly WAF. | [optional] 
**WafLatency60000ms** | Pointer to **int32** | Number of responses with latency of 60,000 milliseconds and above received for origin requests made by the Fastly WAF. | [optional] 
**ComputeResponses** | Pointer to **int32** | Number of responses for origin received by the Compute platform. | [optional] 
**ComputeRespHeaderBytes** | Pointer to **int32** | Number of header bytes for origin received by the Compute platform. | [optional] 
**ComputeRespBodyBytes** | Pointer to **int32** | Number of body bytes for origin received by the Compute platform. | [optional] 
**ComputeStatus1xx** | Pointer to **int32** | Number of 1xx \&quot;Informational\&quot; status codes for origin received by the Compute platform. | [optional] 
**ComputeStatus2xx** | Pointer to **int32** | Number of 2xx \&quot;Success\&quot; status codes for origin received by the Compute platform. | [optional] 
**ComputeStatus3xx** | Pointer to **int32** | Number of 3xx \&quot;Redirection\&quot; codes for origin received by the Compute platform. | [optional] 
**ComputeStatus4xx** | Pointer to **int32** | Number of 4xx \&quot;Client Error\&quot; codes for origin received by the Compute platform. | [optional] 
**ComputeStatus5xx** | Pointer to **int32** | Number of 5xx \&quot;Server Error\&quot; codes for origin received by the Compute platform. | [optional] 
**ComputeStatus200** | Pointer to **int32** | Number of responses received with status code 200 (Success) for origin received by the Compute platform. | [optional] 
**ComputeStatus204** | Pointer to **int32** | Number of responses received with status code 204 (No Content) for origin received by the Compute platform. | [optional] 
**ComputeStatus206** | Pointer to **int32** | Number of responses received with status code 206 (Partial Content) for origin received by the Compute platform. | [optional] 
**ComputeStatus301** | Pointer to **int32** | Number of responses received with status code 301 (Moved Permanently) for origin received by the Compute platform. | [optional] 
**ComputeStatus302** | Pointer to **int32** | Number of responses received with status code 302 (Found) for origin received by the Compute platform. | [optional] 
**ComputeStatus304** | Pointer to **int32** | Number of responses received with status code 304 (Not Modified) for origin received by the Compute platform. | [optional] 
**ComputeStatus400** | Pointer to **int32** | Number of responses received with status code 400 (Bad Request) for origin received by the Compute platform. | [optional] 
**ComputeStatus401** | Pointer to **int32** | Number of responses received with status code 401 (Unauthorized) for origin received by the Compute platform. | [optional] 
**ComputeStatus403** | Pointer to **int32** | Number of responses received with status code 403 (Forbidden) for origin received by the Compute platform. | [optional] 
**ComputeStatus404** | Pointer to **int32** | Number of responses received with status code 404 (Not Found) for origin received by the Compute platform. | [optional] 
**ComputeStatus416** | Pointer to **int32** | Number of responses received with status code 416 (Range Not Satisfiable) for origin received by the Compute platform. | [optional] 
**ComputeStatus429** | Pointer to **int32** | Number of responses received with status code 429 (Too Many Requests) for origin received by the Compute platform. | [optional] 
**ComputeStatus500** | Pointer to **int32** | Number of responses received with status code 500 (Internal Server Error) for origin received by the Compute platform. | [optional] 
**ComputeStatus501** | Pointer to **int32** | Number of responses received with status code 501 (Not Implemented) for origin received by the Compute platform. | [optional] 
**ComputeStatus502** | Pointer to **int32** | Number of responses received with status code 502 (Bad Gateway) for origin received by the Compute platform. | [optional] 
**ComputeStatus503** | Pointer to **int32** | Number of responses received with status code 503 (Service Unavailable) for origin received by the Compute platform. | [optional] 
**ComputeStatus504** | Pointer to **int32** | Number of responses received with status code 504 (Gateway Timeout) for origin received by the Compute platform. | [optional] 
**ComputeStatus505** | Pointer to **int32** | Number of responses received with status code 505 (HTTP Version Not Supported) for origin received by the Compute platform. | [optional] 
**ComputeLatency0To1ms** | Pointer to **int32** | Number of responses with latency between 0 and 1 millisecond for origin received by the Compute platform. | [optional] 
**ComputeLatency1To5ms** | Pointer to **int32** | Number of responses with latency between 1 and 5 milliseconds for origin received by the Compute platform. | [optional] 
**ComputeLatency5To10ms** | Pointer to **int32** | Number of responses with latency between 5 and 10 milliseconds for origin received by the Compute platform. | [optional] 
**ComputeLatency10To50ms** | Pointer to **int32** | Number of responses with latency between 10 and 50 milliseconds for origin received by the Compute platform. | [optional] 
**ComputeLatency50To100ms** | Pointer to **int32** | Number of responses with latency between 50 and 100 milliseconds for origin received by the Compute platform. | [optional] 
**ComputeLatency100To250ms** | Pointer to **int32** | Number of responses with latency between 100 and 250 milliseconds for origin received by the Compute platform. | [optional] 
**ComputeLatency250To500ms** | Pointer to **int32** | Number of responses with latency between 250 and 500 milliseconds for origin received by the Compute platform. | [optional] 
**ComputeLatency500To1000ms** | Pointer to **int32** | Number of responses with latency between 500 and 1,000 milliseconds for origin received by the Compute platform. | [optional] 
**ComputeLatency1000To5000ms** | Pointer to **int32** | Number of responses with latency between 1,000 and 5,000 milliseconds for origin received by the Compute platform. | [optional] 
**ComputeLatency5000To10000ms** | Pointer to **int32** | Number of responses with latency between 5,000 and 10,000 milliseconds for origin received by the Compute platform. | [optional] 
**ComputeLatency10000To60000ms** | Pointer to **int32** | Number of responses with latency between 10,000 and 60,000 milliseconds for origin received by the Compute platform. | [optional] 
**ComputeLatency60000ms** | Pointer to **int32** | Number of responses with latency of 60,000 milliseconds and above for origin received by the Compute platform. | [optional] 
**AllResponses** | Pointer to **int32** | Number of responses received for origin requests made by all sources. | [optional] 
**AllRespHeaderBytes** | Pointer to **int32** | Number of header bytes received for origin requests made by all sources. | [optional] 
**AllRespBodyBytes** | Pointer to **int32** | Number of body bytes received for origin requests made by all sources. | [optional] 
**AllStatus1xx** | Pointer to **int32** | Number of 1xx \&quot;Informational\&quot; category status codes delivered received for origin requests made by all sources. | [optional] 
**AllStatus2xx** | Pointer to **int32** | Number of 2xx \&quot;Success\&quot; status codes received for origin requests made by all sources. | [optional] 
**AllStatus3xx** | Pointer to **int32** | Number of 3xx \&quot;Redirection\&quot; codes received for origin requests made by all sources. | [optional] 
**AllStatus4xx** | Pointer to **int32** | Number of 4xx \&quot;Client Error\&quot; codes received for origin requests made by all sources. | [optional] 
**AllStatus5xx** | Pointer to **int32** | Number of 5xx \&quot;Server Error\&quot; codes received for origin requests made by all sources. | [optional] 
**AllStatus200** | Pointer to **int32** | Number of responses received with status code 200 (Success) received for origin requests made by all sources. | [optional] 
**AllStatus204** | Pointer to **int32** | Number of responses received with status code 204 (No Content) received for origin requests made by all sources. | [optional] 
**AllStatus206** | Pointer to **int32** | Number of responses received with status code 206 (Partial Content) received for origin requests made by all sources. | [optional] 
**AllStatus301** | Pointer to **int32** | Number of responses received with status code 301 (Moved Permanently) received for origin requests made by all sources. | [optional] 
**AllStatus302** | Pointer to **int32** | Number of responses received with status code 302 (Found) received for origin requests made by all sources. | [optional] 
**AllStatus304** | Pointer to **int32** | Number of responses received with status code 304 (Not Modified) received for origin requests made by all sources. | [optional] 
**AllStatus400** | Pointer to **int32** | Number of responses received with status code 400 (Bad Request) received for origin requests made by all sources. | [optional] 
**AllStatus401** | Pointer to **int32** | Number of responses received with status code 401 (Unauthorized) received for origin requests made by all sources. | [optional] 
**AllStatus403** | Pointer to **int32** | Number of responses received with status code 403 (Forbidden) received for origin requests made by all sources. | [optional] 
**AllStatus404** | Pointer to **int32** | Number of responses received with status code 404 (Not Found) received for origin requests made by all sources. | [optional] 
**AllStatus416** | Pointer to **int32** | Number of responses received with status code 416 (Range Not Satisfiable) received for origin requests made by all sources. | [optional] 
**AllStatus429** | Pointer to **int32** | Number of responses received with status code 429 (Too Many Requests) received for origin requests made by all sources. | [optional] 
**AllStatus500** | Pointer to **int32** | Number of responses received with status code 500 (Internal Server Error) received for origin requests made by all sources. | [optional] 
**AllStatus501** | Pointer to **int32** | Number of responses received with status code 501 (Not Implemented) received for origin requests made by all sources. | [optional] 
**AllStatus502** | Pointer to **int32** | Number of responses received with status code 502 (Bad Gateway) received for origin requests made by all sources. | [optional] 
**AllStatus503** | Pointer to **int32** | Number of responses received with status code 503 (Service Unavailable) received for origin requests made by all sources. | [optional] 
**AllStatus504** | Pointer to **int32** | Number of responses received with status code 504 (Gateway Timeout) received for origin requests made by all sources. | [optional] 
**AllStatus505** | Pointer to **int32** | Number of responses received with status code 505 (HTTP Version Not Supported) received for origin requests made by all sources. | [optional] 
**AllLatency0To1ms** | Pointer to **int32** | Number of responses with latency between 0 and 1 millisecond received for origin requests made by all sources. | [optional] 
**AllLatency1To5ms** | Pointer to **int32** | Number of responses with latency between 1 and 5 milliseconds received for origin requests made by all sources. | [optional] 
**AllLatency5To10ms** | Pointer to **int32** | Number of responses with latency between 5 and 10 milliseconds received for origin requests made by all sources. | [optional] 
**AllLatency10To50ms** | Pointer to **int32** | Number of responses with latency between 10 and 50 milliseconds received for origin requests made by all sources. | [optional] 
**AllLatency50To100ms** | Pointer to **int32** | Number of responses with latency between 50 and 100 milliseconds received for origin requests made by all sources. | [optional] 
**AllLatency100To250ms** | Pointer to **int32** | Number of responses with latency between 100 and 250 milliseconds received for origin requests made by all sources. | [optional] 
**AllLatency250To500ms** | Pointer to **int32** | Number of responses with latency between 250 and 500 milliseconds received for origin requests made by all sources. | [optional] 
**AllLatency500To1000ms** | Pointer to **int32** | Number of responses with latency between 500 and 1,000 milliseconds received for origin requests made by all sources. | [optional] 
**AllLatency1000To5000ms** | Pointer to **int32** | Number of responses with latency between 1,000 and 5,000 milliseconds received for origin requests made by all sources. | [optional] 
**AllLatency5000To10000ms** | Pointer to **int32** | Number of responses with latency between 5,000 and 10,000 milliseconds received for origin requests made by all sources. | [optional] 
**AllLatency10000To60000ms** | Pointer to **int32** | Number of responses with latency between 10,000 and 60,000 milliseconds received for origin requests made by all sources. | [optional] 
**AllLatency60000ms** | Pointer to **int32** | Number of responses with latency of 60,000 milliseconds and above received for origin requests made by all sources. | [optional] 

## Methods

### NewOriginInspectorMeasurements

`func NewOriginInspectorMeasurements() *OriginInspectorMeasurements`

NewOriginInspectorMeasurements instantiates a new OriginInspectorMeasurements object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOriginInspectorMeasurementsWithDefaults

`func NewOriginInspectorMeasurementsWithDefaults() *OriginInspectorMeasurements`

NewOriginInspectorMeasurementsWithDefaults instantiates a new OriginInspectorMeasurements object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponses

`func (o *OriginInspectorMeasurements) GetResponses() int32`

GetResponses returns the Responses field if non-nil, zero value otherwise.

### GetResponsesOk

`func (o *OriginInspectorMeasurements) GetResponsesOk() (*int32, bool)`

GetResponsesOk returns a tuple with the Responses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponses

`func (o *OriginInspectorMeasurements) SetResponses(v int32)`

SetResponses sets Responses field to given value.

### HasResponses

`func (o *OriginInspectorMeasurements) HasResponses() bool`

HasResponses returns a boolean if a field has been set.

### GetRespHeaderBytes

`func (o *OriginInspectorMeasurements) GetRespHeaderBytes() int32`

GetRespHeaderBytes returns the RespHeaderBytes field if non-nil, zero value otherwise.

### GetRespHeaderBytesOk

`func (o *OriginInspectorMeasurements) GetRespHeaderBytesOk() (*int32, bool)`

GetRespHeaderBytesOk returns a tuple with the RespHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRespHeaderBytes

`func (o *OriginInspectorMeasurements) SetRespHeaderBytes(v int32)`

SetRespHeaderBytes sets RespHeaderBytes field to given value.

### HasRespHeaderBytes

`func (o *OriginInspectorMeasurements) HasRespHeaderBytes() bool`

HasRespHeaderBytes returns a boolean if a field has been set.

### GetRespBodyBytes

`func (o *OriginInspectorMeasurements) GetRespBodyBytes() int32`

GetRespBodyBytes returns the RespBodyBytes field if non-nil, zero value otherwise.

### GetRespBodyBytesOk

`func (o *OriginInspectorMeasurements) GetRespBodyBytesOk() (*int32, bool)`

GetRespBodyBytesOk returns a tuple with the RespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRespBodyBytes

`func (o *OriginInspectorMeasurements) SetRespBodyBytes(v int32)`

SetRespBodyBytes sets RespBodyBytes field to given value.

### HasRespBodyBytes

`func (o *OriginInspectorMeasurements) HasRespBodyBytes() bool`

HasRespBodyBytes returns a boolean if a field has been set.

### GetStatus1xx

`func (o *OriginInspectorMeasurements) GetStatus1xx() int32`

GetStatus1xx returns the Status1xx field if non-nil, zero value otherwise.

### GetStatus1xxOk

`func (o *OriginInspectorMeasurements) GetStatus1xxOk() (*int32, bool)`

GetStatus1xxOk returns a tuple with the Status1xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus1xx

`func (o *OriginInspectorMeasurements) SetStatus1xx(v int32)`

SetStatus1xx sets Status1xx field to given value.

### HasStatus1xx

`func (o *OriginInspectorMeasurements) HasStatus1xx() bool`

HasStatus1xx returns a boolean if a field has been set.

### GetStatus2xx

`func (o *OriginInspectorMeasurements) GetStatus2xx() int32`

GetStatus2xx returns the Status2xx field if non-nil, zero value otherwise.

### GetStatus2xxOk

`func (o *OriginInspectorMeasurements) GetStatus2xxOk() (*int32, bool)`

GetStatus2xxOk returns a tuple with the Status2xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus2xx

`func (o *OriginInspectorMeasurements) SetStatus2xx(v int32)`

SetStatus2xx sets Status2xx field to given value.

### HasStatus2xx

`func (o *OriginInspectorMeasurements) HasStatus2xx() bool`

HasStatus2xx returns a boolean if a field has been set.

### GetStatus3xx

`func (o *OriginInspectorMeasurements) GetStatus3xx() int32`

GetStatus3xx returns the Status3xx field if non-nil, zero value otherwise.

### GetStatus3xxOk

`func (o *OriginInspectorMeasurements) GetStatus3xxOk() (*int32, bool)`

GetStatus3xxOk returns a tuple with the Status3xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus3xx

`func (o *OriginInspectorMeasurements) SetStatus3xx(v int32)`

SetStatus3xx sets Status3xx field to given value.

### HasStatus3xx

`func (o *OriginInspectorMeasurements) HasStatus3xx() bool`

HasStatus3xx returns a boolean if a field has been set.

### GetStatus4xx

`func (o *OriginInspectorMeasurements) GetStatus4xx() int32`

GetStatus4xx returns the Status4xx field if non-nil, zero value otherwise.

### GetStatus4xxOk

`func (o *OriginInspectorMeasurements) GetStatus4xxOk() (*int32, bool)`

GetStatus4xxOk returns a tuple with the Status4xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus4xx

`func (o *OriginInspectorMeasurements) SetStatus4xx(v int32)`

SetStatus4xx sets Status4xx field to given value.

### HasStatus4xx

`func (o *OriginInspectorMeasurements) HasStatus4xx() bool`

HasStatus4xx returns a boolean if a field has been set.

### GetStatus5xx

`func (o *OriginInspectorMeasurements) GetStatus5xx() int32`

GetStatus5xx returns the Status5xx field if non-nil, zero value otherwise.

### GetStatus5xxOk

`func (o *OriginInspectorMeasurements) GetStatus5xxOk() (*int32, bool)`

GetStatus5xxOk returns a tuple with the Status5xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus5xx

`func (o *OriginInspectorMeasurements) SetStatus5xx(v int32)`

SetStatus5xx sets Status5xx field to given value.

### HasStatus5xx

`func (o *OriginInspectorMeasurements) HasStatus5xx() bool`

HasStatus5xx returns a boolean if a field has been set.

### GetStatus200

`func (o *OriginInspectorMeasurements) GetStatus200() int32`

GetStatus200 returns the Status200 field if non-nil, zero value otherwise.

### GetStatus200Ok

`func (o *OriginInspectorMeasurements) GetStatus200Ok() (*int32, bool)`

GetStatus200Ok returns a tuple with the Status200 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus200

`func (o *OriginInspectorMeasurements) SetStatus200(v int32)`

SetStatus200 sets Status200 field to given value.

### HasStatus200

`func (o *OriginInspectorMeasurements) HasStatus200() bool`

HasStatus200 returns a boolean if a field has been set.

### GetStatus204

`func (o *OriginInspectorMeasurements) GetStatus204() int32`

GetStatus204 returns the Status204 field if non-nil, zero value otherwise.

### GetStatus204Ok

`func (o *OriginInspectorMeasurements) GetStatus204Ok() (*int32, bool)`

GetStatus204Ok returns a tuple with the Status204 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus204

`func (o *OriginInspectorMeasurements) SetStatus204(v int32)`

SetStatus204 sets Status204 field to given value.

### HasStatus204

`func (o *OriginInspectorMeasurements) HasStatus204() bool`

HasStatus204 returns a boolean if a field has been set.

### GetStatus206

`func (o *OriginInspectorMeasurements) GetStatus206() int32`

GetStatus206 returns the Status206 field if non-nil, zero value otherwise.

### GetStatus206Ok

`func (o *OriginInspectorMeasurements) GetStatus206Ok() (*int32, bool)`

GetStatus206Ok returns a tuple with the Status206 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus206

`func (o *OriginInspectorMeasurements) SetStatus206(v int32)`

SetStatus206 sets Status206 field to given value.

### HasStatus206

`func (o *OriginInspectorMeasurements) HasStatus206() bool`

HasStatus206 returns a boolean if a field has been set.

### GetStatus301

`func (o *OriginInspectorMeasurements) GetStatus301() int32`

GetStatus301 returns the Status301 field if non-nil, zero value otherwise.

### GetStatus301Ok

`func (o *OriginInspectorMeasurements) GetStatus301Ok() (*int32, bool)`

GetStatus301Ok returns a tuple with the Status301 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus301

`func (o *OriginInspectorMeasurements) SetStatus301(v int32)`

SetStatus301 sets Status301 field to given value.

### HasStatus301

`func (o *OriginInspectorMeasurements) HasStatus301() bool`

HasStatus301 returns a boolean if a field has been set.

### GetStatus302

`func (o *OriginInspectorMeasurements) GetStatus302() int32`

GetStatus302 returns the Status302 field if non-nil, zero value otherwise.

### GetStatus302Ok

`func (o *OriginInspectorMeasurements) GetStatus302Ok() (*int32, bool)`

GetStatus302Ok returns a tuple with the Status302 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus302

`func (o *OriginInspectorMeasurements) SetStatus302(v int32)`

SetStatus302 sets Status302 field to given value.

### HasStatus302

`func (o *OriginInspectorMeasurements) HasStatus302() bool`

HasStatus302 returns a boolean if a field has been set.

### GetStatus304

`func (o *OriginInspectorMeasurements) GetStatus304() int32`

GetStatus304 returns the Status304 field if non-nil, zero value otherwise.

### GetStatus304Ok

`func (o *OriginInspectorMeasurements) GetStatus304Ok() (*int32, bool)`

GetStatus304Ok returns a tuple with the Status304 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus304

`func (o *OriginInspectorMeasurements) SetStatus304(v int32)`

SetStatus304 sets Status304 field to given value.

### HasStatus304

`func (o *OriginInspectorMeasurements) HasStatus304() bool`

HasStatus304 returns a boolean if a field has been set.

### GetStatus400

`func (o *OriginInspectorMeasurements) GetStatus400() int32`

GetStatus400 returns the Status400 field if non-nil, zero value otherwise.

### GetStatus400Ok

`func (o *OriginInspectorMeasurements) GetStatus400Ok() (*int32, bool)`

GetStatus400Ok returns a tuple with the Status400 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus400

`func (o *OriginInspectorMeasurements) SetStatus400(v int32)`

SetStatus400 sets Status400 field to given value.

### HasStatus400

`func (o *OriginInspectorMeasurements) HasStatus400() bool`

HasStatus400 returns a boolean if a field has been set.

### GetStatus401

`func (o *OriginInspectorMeasurements) GetStatus401() int32`

GetStatus401 returns the Status401 field if non-nil, zero value otherwise.

### GetStatus401Ok

`func (o *OriginInspectorMeasurements) GetStatus401Ok() (*int32, bool)`

GetStatus401Ok returns a tuple with the Status401 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus401

`func (o *OriginInspectorMeasurements) SetStatus401(v int32)`

SetStatus401 sets Status401 field to given value.

### HasStatus401

`func (o *OriginInspectorMeasurements) HasStatus401() bool`

HasStatus401 returns a boolean if a field has been set.

### GetStatus403

`func (o *OriginInspectorMeasurements) GetStatus403() int32`

GetStatus403 returns the Status403 field if non-nil, zero value otherwise.

### GetStatus403Ok

`func (o *OriginInspectorMeasurements) GetStatus403Ok() (*int32, bool)`

GetStatus403Ok returns a tuple with the Status403 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus403

`func (o *OriginInspectorMeasurements) SetStatus403(v int32)`

SetStatus403 sets Status403 field to given value.

### HasStatus403

`func (o *OriginInspectorMeasurements) HasStatus403() bool`

HasStatus403 returns a boolean if a field has been set.

### GetStatus404

`func (o *OriginInspectorMeasurements) GetStatus404() int32`

GetStatus404 returns the Status404 field if non-nil, zero value otherwise.

### GetStatus404Ok

`func (o *OriginInspectorMeasurements) GetStatus404Ok() (*int32, bool)`

GetStatus404Ok returns a tuple with the Status404 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus404

`func (o *OriginInspectorMeasurements) SetStatus404(v int32)`

SetStatus404 sets Status404 field to given value.

### HasStatus404

`func (o *OriginInspectorMeasurements) HasStatus404() bool`

HasStatus404 returns a boolean if a field has been set.

### GetStatus416

`func (o *OriginInspectorMeasurements) GetStatus416() int32`

GetStatus416 returns the Status416 field if non-nil, zero value otherwise.

### GetStatus416Ok

`func (o *OriginInspectorMeasurements) GetStatus416Ok() (*int32, bool)`

GetStatus416Ok returns a tuple with the Status416 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus416

`func (o *OriginInspectorMeasurements) SetStatus416(v int32)`

SetStatus416 sets Status416 field to given value.

### HasStatus416

`func (o *OriginInspectorMeasurements) HasStatus416() bool`

HasStatus416 returns a boolean if a field has been set.

### GetStatus429

`func (o *OriginInspectorMeasurements) GetStatus429() int32`

GetStatus429 returns the Status429 field if non-nil, zero value otherwise.

### GetStatus429Ok

`func (o *OriginInspectorMeasurements) GetStatus429Ok() (*int32, bool)`

GetStatus429Ok returns a tuple with the Status429 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus429

`func (o *OriginInspectorMeasurements) SetStatus429(v int32)`

SetStatus429 sets Status429 field to given value.

### HasStatus429

`func (o *OriginInspectorMeasurements) HasStatus429() bool`

HasStatus429 returns a boolean if a field has been set.

### GetStatus500

`func (o *OriginInspectorMeasurements) GetStatus500() int32`

GetStatus500 returns the Status500 field if non-nil, zero value otherwise.

### GetStatus500Ok

`func (o *OriginInspectorMeasurements) GetStatus500Ok() (*int32, bool)`

GetStatus500Ok returns a tuple with the Status500 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus500

`func (o *OriginInspectorMeasurements) SetStatus500(v int32)`

SetStatus500 sets Status500 field to given value.

### HasStatus500

`func (o *OriginInspectorMeasurements) HasStatus500() bool`

HasStatus500 returns a boolean if a field has been set.

### GetStatus501

`func (o *OriginInspectorMeasurements) GetStatus501() int32`

GetStatus501 returns the Status501 field if non-nil, zero value otherwise.

### GetStatus501Ok

`func (o *OriginInspectorMeasurements) GetStatus501Ok() (*int32, bool)`

GetStatus501Ok returns a tuple with the Status501 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus501

`func (o *OriginInspectorMeasurements) SetStatus501(v int32)`

SetStatus501 sets Status501 field to given value.

### HasStatus501

`func (o *OriginInspectorMeasurements) HasStatus501() bool`

HasStatus501 returns a boolean if a field has been set.

### GetStatus502

`func (o *OriginInspectorMeasurements) GetStatus502() int32`

GetStatus502 returns the Status502 field if non-nil, zero value otherwise.

### GetStatus502Ok

`func (o *OriginInspectorMeasurements) GetStatus502Ok() (*int32, bool)`

GetStatus502Ok returns a tuple with the Status502 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus502

`func (o *OriginInspectorMeasurements) SetStatus502(v int32)`

SetStatus502 sets Status502 field to given value.

### HasStatus502

`func (o *OriginInspectorMeasurements) HasStatus502() bool`

HasStatus502 returns a boolean if a field has been set.

### GetStatus503

`func (o *OriginInspectorMeasurements) GetStatus503() int32`

GetStatus503 returns the Status503 field if non-nil, zero value otherwise.

### GetStatus503Ok

`func (o *OriginInspectorMeasurements) GetStatus503Ok() (*int32, bool)`

GetStatus503Ok returns a tuple with the Status503 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus503

`func (o *OriginInspectorMeasurements) SetStatus503(v int32)`

SetStatus503 sets Status503 field to given value.

### HasStatus503

`func (o *OriginInspectorMeasurements) HasStatus503() bool`

HasStatus503 returns a boolean if a field has been set.

### GetStatus504

`func (o *OriginInspectorMeasurements) GetStatus504() int32`

GetStatus504 returns the Status504 field if non-nil, zero value otherwise.

### GetStatus504Ok

`func (o *OriginInspectorMeasurements) GetStatus504Ok() (*int32, bool)`

GetStatus504Ok returns a tuple with the Status504 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus504

`func (o *OriginInspectorMeasurements) SetStatus504(v int32)`

SetStatus504 sets Status504 field to given value.

### HasStatus504

`func (o *OriginInspectorMeasurements) HasStatus504() bool`

HasStatus504 returns a boolean if a field has been set.

### GetStatus505

`func (o *OriginInspectorMeasurements) GetStatus505() int32`

GetStatus505 returns the Status505 field if non-nil, zero value otherwise.

### GetStatus505Ok

`func (o *OriginInspectorMeasurements) GetStatus505Ok() (*int32, bool)`

GetStatus505Ok returns a tuple with the Status505 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus505

`func (o *OriginInspectorMeasurements) SetStatus505(v int32)`

SetStatus505 sets Status505 field to given value.

### HasStatus505

`func (o *OriginInspectorMeasurements) HasStatus505() bool`

HasStatus505 returns a boolean if a field has been set.

### GetLatency0To1ms

`func (o *OriginInspectorMeasurements) GetLatency0To1ms() int32`

GetLatency0To1ms returns the Latency0To1ms field if non-nil, zero value otherwise.

### GetLatency0To1msOk

`func (o *OriginInspectorMeasurements) GetLatency0To1msOk() (*int32, bool)`

GetLatency0To1msOk returns a tuple with the Latency0To1ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatency0To1ms

`func (o *OriginInspectorMeasurements) SetLatency0To1ms(v int32)`

SetLatency0To1ms sets Latency0To1ms field to given value.

### HasLatency0To1ms

`func (o *OriginInspectorMeasurements) HasLatency0To1ms() bool`

HasLatency0To1ms returns a boolean if a field has been set.

### GetLatency1To5ms

`func (o *OriginInspectorMeasurements) GetLatency1To5ms() int32`

GetLatency1To5ms returns the Latency1To5ms field if non-nil, zero value otherwise.

### GetLatency1To5msOk

`func (o *OriginInspectorMeasurements) GetLatency1To5msOk() (*int32, bool)`

GetLatency1To5msOk returns a tuple with the Latency1To5ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatency1To5ms

`func (o *OriginInspectorMeasurements) SetLatency1To5ms(v int32)`

SetLatency1To5ms sets Latency1To5ms field to given value.

### HasLatency1To5ms

`func (o *OriginInspectorMeasurements) HasLatency1To5ms() bool`

HasLatency1To5ms returns a boolean if a field has been set.

### GetLatency5To10ms

`func (o *OriginInspectorMeasurements) GetLatency5To10ms() int32`

GetLatency5To10ms returns the Latency5To10ms field if non-nil, zero value otherwise.

### GetLatency5To10msOk

`func (o *OriginInspectorMeasurements) GetLatency5To10msOk() (*int32, bool)`

GetLatency5To10msOk returns a tuple with the Latency5To10ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatency5To10ms

`func (o *OriginInspectorMeasurements) SetLatency5To10ms(v int32)`

SetLatency5To10ms sets Latency5To10ms field to given value.

### HasLatency5To10ms

`func (o *OriginInspectorMeasurements) HasLatency5To10ms() bool`

HasLatency5To10ms returns a boolean if a field has been set.

### GetLatency10To50ms

`func (o *OriginInspectorMeasurements) GetLatency10To50ms() int32`

GetLatency10To50ms returns the Latency10To50ms field if non-nil, zero value otherwise.

### GetLatency10To50msOk

`func (o *OriginInspectorMeasurements) GetLatency10To50msOk() (*int32, bool)`

GetLatency10To50msOk returns a tuple with the Latency10To50ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatency10To50ms

`func (o *OriginInspectorMeasurements) SetLatency10To50ms(v int32)`

SetLatency10To50ms sets Latency10To50ms field to given value.

### HasLatency10To50ms

`func (o *OriginInspectorMeasurements) HasLatency10To50ms() bool`

HasLatency10To50ms returns a boolean if a field has been set.

### GetLatency50To100ms

`func (o *OriginInspectorMeasurements) GetLatency50To100ms() int32`

GetLatency50To100ms returns the Latency50To100ms field if non-nil, zero value otherwise.

### GetLatency50To100msOk

`func (o *OriginInspectorMeasurements) GetLatency50To100msOk() (*int32, bool)`

GetLatency50To100msOk returns a tuple with the Latency50To100ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatency50To100ms

`func (o *OriginInspectorMeasurements) SetLatency50To100ms(v int32)`

SetLatency50To100ms sets Latency50To100ms field to given value.

### HasLatency50To100ms

`func (o *OriginInspectorMeasurements) HasLatency50To100ms() bool`

HasLatency50To100ms returns a boolean if a field has been set.

### GetLatency100To250ms

`func (o *OriginInspectorMeasurements) GetLatency100To250ms() int32`

GetLatency100To250ms returns the Latency100To250ms field if non-nil, zero value otherwise.

### GetLatency100To250msOk

`func (o *OriginInspectorMeasurements) GetLatency100To250msOk() (*int32, bool)`

GetLatency100To250msOk returns a tuple with the Latency100To250ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatency100To250ms

`func (o *OriginInspectorMeasurements) SetLatency100To250ms(v int32)`

SetLatency100To250ms sets Latency100To250ms field to given value.

### HasLatency100To250ms

`func (o *OriginInspectorMeasurements) HasLatency100To250ms() bool`

HasLatency100To250ms returns a boolean if a field has been set.

### GetLatency250To500ms

`func (o *OriginInspectorMeasurements) GetLatency250To500ms() int32`

GetLatency250To500ms returns the Latency250To500ms field if non-nil, zero value otherwise.

### GetLatency250To500msOk

`func (o *OriginInspectorMeasurements) GetLatency250To500msOk() (*int32, bool)`

GetLatency250To500msOk returns a tuple with the Latency250To500ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatency250To500ms

`func (o *OriginInspectorMeasurements) SetLatency250To500ms(v int32)`

SetLatency250To500ms sets Latency250To500ms field to given value.

### HasLatency250To500ms

`func (o *OriginInspectorMeasurements) HasLatency250To500ms() bool`

HasLatency250To500ms returns a boolean if a field has been set.

### GetLatency500To1000ms

`func (o *OriginInspectorMeasurements) GetLatency500To1000ms() int32`

GetLatency500To1000ms returns the Latency500To1000ms field if non-nil, zero value otherwise.

### GetLatency500To1000msOk

`func (o *OriginInspectorMeasurements) GetLatency500To1000msOk() (*int32, bool)`

GetLatency500To1000msOk returns a tuple with the Latency500To1000ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatency500To1000ms

`func (o *OriginInspectorMeasurements) SetLatency500To1000ms(v int32)`

SetLatency500To1000ms sets Latency500To1000ms field to given value.

### HasLatency500To1000ms

`func (o *OriginInspectorMeasurements) HasLatency500To1000ms() bool`

HasLatency500To1000ms returns a boolean if a field has been set.

### GetLatency1000To5000ms

`func (o *OriginInspectorMeasurements) GetLatency1000To5000ms() int32`

GetLatency1000To5000ms returns the Latency1000To5000ms field if non-nil, zero value otherwise.

### GetLatency1000To5000msOk

`func (o *OriginInspectorMeasurements) GetLatency1000To5000msOk() (*int32, bool)`

GetLatency1000To5000msOk returns a tuple with the Latency1000To5000ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatency1000To5000ms

`func (o *OriginInspectorMeasurements) SetLatency1000To5000ms(v int32)`

SetLatency1000To5000ms sets Latency1000To5000ms field to given value.

### HasLatency1000To5000ms

`func (o *OriginInspectorMeasurements) HasLatency1000To5000ms() bool`

HasLatency1000To5000ms returns a boolean if a field has been set.

### GetLatency5000To10000ms

`func (o *OriginInspectorMeasurements) GetLatency5000To10000ms() int32`

GetLatency5000To10000ms returns the Latency5000To10000ms field if non-nil, zero value otherwise.

### GetLatency5000To10000msOk

`func (o *OriginInspectorMeasurements) GetLatency5000To10000msOk() (*int32, bool)`

GetLatency5000To10000msOk returns a tuple with the Latency5000To10000ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatency5000To10000ms

`func (o *OriginInspectorMeasurements) SetLatency5000To10000ms(v int32)`

SetLatency5000To10000ms sets Latency5000To10000ms field to given value.

### HasLatency5000To10000ms

`func (o *OriginInspectorMeasurements) HasLatency5000To10000ms() bool`

HasLatency5000To10000ms returns a boolean if a field has been set.

### GetLatency10000To60000ms

`func (o *OriginInspectorMeasurements) GetLatency10000To60000ms() int32`

GetLatency10000To60000ms returns the Latency10000To60000ms field if non-nil, zero value otherwise.

### GetLatency10000To60000msOk

`func (o *OriginInspectorMeasurements) GetLatency10000To60000msOk() (*int32, bool)`

GetLatency10000To60000msOk returns a tuple with the Latency10000To60000ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatency10000To60000ms

`func (o *OriginInspectorMeasurements) SetLatency10000To60000ms(v int32)`

SetLatency10000To60000ms sets Latency10000To60000ms field to given value.

### HasLatency10000To60000ms

`func (o *OriginInspectorMeasurements) HasLatency10000To60000ms() bool`

HasLatency10000To60000ms returns a boolean if a field has been set.

### GetLatency60000ms

`func (o *OriginInspectorMeasurements) GetLatency60000ms() int32`

GetLatency60000ms returns the Latency60000ms field if non-nil, zero value otherwise.

### GetLatency60000msOk

`func (o *OriginInspectorMeasurements) GetLatency60000msOk() (*int32, bool)`

GetLatency60000msOk returns a tuple with the Latency60000ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatency60000ms

`func (o *OriginInspectorMeasurements) SetLatency60000ms(v int32)`

SetLatency60000ms sets Latency60000ms field to given value.

### HasLatency60000ms

`func (o *OriginInspectorMeasurements) HasLatency60000ms() bool`

HasLatency60000ms returns a boolean if a field has been set.

### GetWafResponses

`func (o *OriginInspectorMeasurements) GetWafResponses() int32`

GetWafResponses returns the WafResponses field if non-nil, zero value otherwise.

### GetWafResponsesOk

`func (o *OriginInspectorMeasurements) GetWafResponsesOk() (*int32, bool)`

GetWafResponsesOk returns a tuple with the WafResponses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafResponses

`func (o *OriginInspectorMeasurements) SetWafResponses(v int32)`

SetWafResponses sets WafResponses field to given value.

### HasWafResponses

`func (o *OriginInspectorMeasurements) HasWafResponses() bool`

HasWafResponses returns a boolean if a field has been set.

### GetWafRespHeaderBytes

`func (o *OriginInspectorMeasurements) GetWafRespHeaderBytes() int32`

GetWafRespHeaderBytes returns the WafRespHeaderBytes field if non-nil, zero value otherwise.

### GetWafRespHeaderBytesOk

`func (o *OriginInspectorMeasurements) GetWafRespHeaderBytesOk() (*int32, bool)`

GetWafRespHeaderBytesOk returns a tuple with the WafRespHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafRespHeaderBytes

`func (o *OriginInspectorMeasurements) SetWafRespHeaderBytes(v int32)`

SetWafRespHeaderBytes sets WafRespHeaderBytes field to given value.

### HasWafRespHeaderBytes

`func (o *OriginInspectorMeasurements) HasWafRespHeaderBytes() bool`

HasWafRespHeaderBytes returns a boolean if a field has been set.

### GetWafRespBodyBytes

`func (o *OriginInspectorMeasurements) GetWafRespBodyBytes() int32`

GetWafRespBodyBytes returns the WafRespBodyBytes field if non-nil, zero value otherwise.

### GetWafRespBodyBytesOk

`func (o *OriginInspectorMeasurements) GetWafRespBodyBytesOk() (*int32, bool)`

GetWafRespBodyBytesOk returns a tuple with the WafRespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafRespBodyBytes

`func (o *OriginInspectorMeasurements) SetWafRespBodyBytes(v int32)`

SetWafRespBodyBytes sets WafRespBodyBytes field to given value.

### HasWafRespBodyBytes

`func (o *OriginInspectorMeasurements) HasWafRespBodyBytes() bool`

HasWafRespBodyBytes returns a boolean if a field has been set.

### GetWafStatus1xx

`func (o *OriginInspectorMeasurements) GetWafStatus1xx() int32`

GetWafStatus1xx returns the WafStatus1xx field if non-nil, zero value otherwise.

### GetWafStatus1xxOk

`func (o *OriginInspectorMeasurements) GetWafStatus1xxOk() (*int32, bool)`

GetWafStatus1xxOk returns a tuple with the WafStatus1xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafStatus1xx

`func (o *OriginInspectorMeasurements) SetWafStatus1xx(v int32)`

SetWafStatus1xx sets WafStatus1xx field to given value.

### HasWafStatus1xx

`func (o *OriginInspectorMeasurements) HasWafStatus1xx() bool`

HasWafStatus1xx returns a boolean if a field has been set.

### GetWafStatus2xx

`func (o *OriginInspectorMeasurements) GetWafStatus2xx() int32`

GetWafStatus2xx returns the WafStatus2xx field if non-nil, zero value otherwise.

### GetWafStatus2xxOk

`func (o *OriginInspectorMeasurements) GetWafStatus2xxOk() (*int32, bool)`

GetWafStatus2xxOk returns a tuple with the WafStatus2xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafStatus2xx

`func (o *OriginInspectorMeasurements) SetWafStatus2xx(v int32)`

SetWafStatus2xx sets WafStatus2xx field to given value.

### HasWafStatus2xx

`func (o *OriginInspectorMeasurements) HasWafStatus2xx() bool`

HasWafStatus2xx returns a boolean if a field has been set.

### GetWafStatus3xx

`func (o *OriginInspectorMeasurements) GetWafStatus3xx() int32`

GetWafStatus3xx returns the WafStatus3xx field if non-nil, zero value otherwise.

### GetWafStatus3xxOk

`func (o *OriginInspectorMeasurements) GetWafStatus3xxOk() (*int32, bool)`

GetWafStatus3xxOk returns a tuple with the WafStatus3xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafStatus3xx

`func (o *OriginInspectorMeasurements) SetWafStatus3xx(v int32)`

SetWafStatus3xx sets WafStatus3xx field to given value.

### HasWafStatus3xx

`func (o *OriginInspectorMeasurements) HasWafStatus3xx() bool`

HasWafStatus3xx returns a boolean if a field has been set.

### GetWafStatus4xx

`func (o *OriginInspectorMeasurements) GetWafStatus4xx() int32`

GetWafStatus4xx returns the WafStatus4xx field if non-nil, zero value otherwise.

### GetWafStatus4xxOk

`func (o *OriginInspectorMeasurements) GetWafStatus4xxOk() (*int32, bool)`

GetWafStatus4xxOk returns a tuple with the WafStatus4xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafStatus4xx

`func (o *OriginInspectorMeasurements) SetWafStatus4xx(v int32)`

SetWafStatus4xx sets WafStatus4xx field to given value.

### HasWafStatus4xx

`func (o *OriginInspectorMeasurements) HasWafStatus4xx() bool`

HasWafStatus4xx returns a boolean if a field has been set.

### GetWafStatus5xx

`func (o *OriginInspectorMeasurements) GetWafStatus5xx() int32`

GetWafStatus5xx returns the WafStatus5xx field if non-nil, zero value otherwise.

### GetWafStatus5xxOk

`func (o *OriginInspectorMeasurements) GetWafStatus5xxOk() (*int32, bool)`

GetWafStatus5xxOk returns a tuple with the WafStatus5xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafStatus5xx

`func (o *OriginInspectorMeasurements) SetWafStatus5xx(v int32)`

SetWafStatus5xx sets WafStatus5xx field to given value.

### HasWafStatus5xx

`func (o *OriginInspectorMeasurements) HasWafStatus5xx() bool`

HasWafStatus5xx returns a boolean if a field has been set.

### GetWafStatus200

`func (o *OriginInspectorMeasurements) GetWafStatus200() int32`

GetWafStatus200 returns the WafStatus200 field if non-nil, zero value otherwise.

### GetWafStatus200Ok

`func (o *OriginInspectorMeasurements) GetWafStatus200Ok() (*int32, bool)`

GetWafStatus200Ok returns a tuple with the WafStatus200 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafStatus200

`func (o *OriginInspectorMeasurements) SetWafStatus200(v int32)`

SetWafStatus200 sets WafStatus200 field to given value.

### HasWafStatus200

`func (o *OriginInspectorMeasurements) HasWafStatus200() bool`

HasWafStatus200 returns a boolean if a field has been set.

### GetWafStatus204

`func (o *OriginInspectorMeasurements) GetWafStatus204() int32`

GetWafStatus204 returns the WafStatus204 field if non-nil, zero value otherwise.

### GetWafStatus204Ok

`func (o *OriginInspectorMeasurements) GetWafStatus204Ok() (*int32, bool)`

GetWafStatus204Ok returns a tuple with the WafStatus204 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafStatus204

`func (o *OriginInspectorMeasurements) SetWafStatus204(v int32)`

SetWafStatus204 sets WafStatus204 field to given value.

### HasWafStatus204

`func (o *OriginInspectorMeasurements) HasWafStatus204() bool`

HasWafStatus204 returns a boolean if a field has been set.

### GetWafStatus206

`func (o *OriginInspectorMeasurements) GetWafStatus206() int32`

GetWafStatus206 returns the WafStatus206 field if non-nil, zero value otherwise.

### GetWafStatus206Ok

`func (o *OriginInspectorMeasurements) GetWafStatus206Ok() (*int32, bool)`

GetWafStatus206Ok returns a tuple with the WafStatus206 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafStatus206

`func (o *OriginInspectorMeasurements) SetWafStatus206(v int32)`

SetWafStatus206 sets WafStatus206 field to given value.

### HasWafStatus206

`func (o *OriginInspectorMeasurements) HasWafStatus206() bool`

HasWafStatus206 returns a boolean if a field has been set.

### GetWafStatus301

`func (o *OriginInspectorMeasurements) GetWafStatus301() int32`

GetWafStatus301 returns the WafStatus301 field if non-nil, zero value otherwise.

### GetWafStatus301Ok

`func (o *OriginInspectorMeasurements) GetWafStatus301Ok() (*int32, bool)`

GetWafStatus301Ok returns a tuple with the WafStatus301 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafStatus301

`func (o *OriginInspectorMeasurements) SetWafStatus301(v int32)`

SetWafStatus301 sets WafStatus301 field to given value.

### HasWafStatus301

`func (o *OriginInspectorMeasurements) HasWafStatus301() bool`

HasWafStatus301 returns a boolean if a field has been set.

### GetWafStatus302

`func (o *OriginInspectorMeasurements) GetWafStatus302() int32`

GetWafStatus302 returns the WafStatus302 field if non-nil, zero value otherwise.

### GetWafStatus302Ok

`func (o *OriginInspectorMeasurements) GetWafStatus302Ok() (*int32, bool)`

GetWafStatus302Ok returns a tuple with the WafStatus302 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafStatus302

`func (o *OriginInspectorMeasurements) SetWafStatus302(v int32)`

SetWafStatus302 sets WafStatus302 field to given value.

### HasWafStatus302

`func (o *OriginInspectorMeasurements) HasWafStatus302() bool`

HasWafStatus302 returns a boolean if a field has been set.

### GetWafStatus304

`func (o *OriginInspectorMeasurements) GetWafStatus304() int32`

GetWafStatus304 returns the WafStatus304 field if non-nil, zero value otherwise.

### GetWafStatus304Ok

`func (o *OriginInspectorMeasurements) GetWafStatus304Ok() (*int32, bool)`

GetWafStatus304Ok returns a tuple with the WafStatus304 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafStatus304

`func (o *OriginInspectorMeasurements) SetWafStatus304(v int32)`

SetWafStatus304 sets WafStatus304 field to given value.

### HasWafStatus304

`func (o *OriginInspectorMeasurements) HasWafStatus304() bool`

HasWafStatus304 returns a boolean if a field has been set.

### GetWafStatus400

`func (o *OriginInspectorMeasurements) GetWafStatus400() int32`

GetWafStatus400 returns the WafStatus400 field if non-nil, zero value otherwise.

### GetWafStatus400Ok

`func (o *OriginInspectorMeasurements) GetWafStatus400Ok() (*int32, bool)`

GetWafStatus400Ok returns a tuple with the WafStatus400 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafStatus400

`func (o *OriginInspectorMeasurements) SetWafStatus400(v int32)`

SetWafStatus400 sets WafStatus400 field to given value.

### HasWafStatus400

`func (o *OriginInspectorMeasurements) HasWafStatus400() bool`

HasWafStatus400 returns a boolean if a field has been set.

### GetWafStatus401

`func (o *OriginInspectorMeasurements) GetWafStatus401() int32`

GetWafStatus401 returns the WafStatus401 field if non-nil, zero value otherwise.

### GetWafStatus401Ok

`func (o *OriginInspectorMeasurements) GetWafStatus401Ok() (*int32, bool)`

GetWafStatus401Ok returns a tuple with the WafStatus401 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafStatus401

`func (o *OriginInspectorMeasurements) SetWafStatus401(v int32)`

SetWafStatus401 sets WafStatus401 field to given value.

### HasWafStatus401

`func (o *OriginInspectorMeasurements) HasWafStatus401() bool`

HasWafStatus401 returns a boolean if a field has been set.

### GetWafStatus403

`func (o *OriginInspectorMeasurements) GetWafStatus403() int32`

GetWafStatus403 returns the WafStatus403 field if non-nil, zero value otherwise.

### GetWafStatus403Ok

`func (o *OriginInspectorMeasurements) GetWafStatus403Ok() (*int32, bool)`

GetWafStatus403Ok returns a tuple with the WafStatus403 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafStatus403

`func (o *OriginInspectorMeasurements) SetWafStatus403(v int32)`

SetWafStatus403 sets WafStatus403 field to given value.

### HasWafStatus403

`func (o *OriginInspectorMeasurements) HasWafStatus403() bool`

HasWafStatus403 returns a boolean if a field has been set.

### GetWafStatus404

`func (o *OriginInspectorMeasurements) GetWafStatus404() int32`

GetWafStatus404 returns the WafStatus404 field if non-nil, zero value otherwise.

### GetWafStatus404Ok

`func (o *OriginInspectorMeasurements) GetWafStatus404Ok() (*int32, bool)`

GetWafStatus404Ok returns a tuple with the WafStatus404 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafStatus404

`func (o *OriginInspectorMeasurements) SetWafStatus404(v int32)`

SetWafStatus404 sets WafStatus404 field to given value.

### HasWafStatus404

`func (o *OriginInspectorMeasurements) HasWafStatus404() bool`

HasWafStatus404 returns a boolean if a field has been set.

### GetWafStatus416

`func (o *OriginInspectorMeasurements) GetWafStatus416() int32`

GetWafStatus416 returns the WafStatus416 field if non-nil, zero value otherwise.

### GetWafStatus416Ok

`func (o *OriginInspectorMeasurements) GetWafStatus416Ok() (*int32, bool)`

GetWafStatus416Ok returns a tuple with the WafStatus416 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafStatus416

`func (o *OriginInspectorMeasurements) SetWafStatus416(v int32)`

SetWafStatus416 sets WafStatus416 field to given value.

### HasWafStatus416

`func (o *OriginInspectorMeasurements) HasWafStatus416() bool`

HasWafStatus416 returns a boolean if a field has been set.

### GetWafStatus429

`func (o *OriginInspectorMeasurements) GetWafStatus429() int32`

GetWafStatus429 returns the WafStatus429 field if non-nil, zero value otherwise.

### GetWafStatus429Ok

`func (o *OriginInspectorMeasurements) GetWafStatus429Ok() (*int32, bool)`

GetWafStatus429Ok returns a tuple with the WafStatus429 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafStatus429

`func (o *OriginInspectorMeasurements) SetWafStatus429(v int32)`

SetWafStatus429 sets WafStatus429 field to given value.

### HasWafStatus429

`func (o *OriginInspectorMeasurements) HasWafStatus429() bool`

HasWafStatus429 returns a boolean if a field has been set.

### GetWafStatus500

`func (o *OriginInspectorMeasurements) GetWafStatus500() int32`

GetWafStatus500 returns the WafStatus500 field if non-nil, zero value otherwise.

### GetWafStatus500Ok

`func (o *OriginInspectorMeasurements) GetWafStatus500Ok() (*int32, bool)`

GetWafStatus500Ok returns a tuple with the WafStatus500 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafStatus500

`func (o *OriginInspectorMeasurements) SetWafStatus500(v int32)`

SetWafStatus500 sets WafStatus500 field to given value.

### HasWafStatus500

`func (o *OriginInspectorMeasurements) HasWafStatus500() bool`

HasWafStatus500 returns a boolean if a field has been set.

### GetWafStatus501

`func (o *OriginInspectorMeasurements) GetWafStatus501() int32`

GetWafStatus501 returns the WafStatus501 field if non-nil, zero value otherwise.

### GetWafStatus501Ok

`func (o *OriginInspectorMeasurements) GetWafStatus501Ok() (*int32, bool)`

GetWafStatus501Ok returns a tuple with the WafStatus501 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafStatus501

`func (o *OriginInspectorMeasurements) SetWafStatus501(v int32)`

SetWafStatus501 sets WafStatus501 field to given value.

### HasWafStatus501

`func (o *OriginInspectorMeasurements) HasWafStatus501() bool`

HasWafStatus501 returns a boolean if a field has been set.

### GetWafStatus502

`func (o *OriginInspectorMeasurements) GetWafStatus502() int32`

GetWafStatus502 returns the WafStatus502 field if non-nil, zero value otherwise.

### GetWafStatus502Ok

`func (o *OriginInspectorMeasurements) GetWafStatus502Ok() (*int32, bool)`

GetWafStatus502Ok returns a tuple with the WafStatus502 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafStatus502

`func (o *OriginInspectorMeasurements) SetWafStatus502(v int32)`

SetWafStatus502 sets WafStatus502 field to given value.

### HasWafStatus502

`func (o *OriginInspectorMeasurements) HasWafStatus502() bool`

HasWafStatus502 returns a boolean if a field has been set.

### GetWafStatus503

`func (o *OriginInspectorMeasurements) GetWafStatus503() int32`

GetWafStatus503 returns the WafStatus503 field if non-nil, zero value otherwise.

### GetWafStatus503Ok

`func (o *OriginInspectorMeasurements) GetWafStatus503Ok() (*int32, bool)`

GetWafStatus503Ok returns a tuple with the WafStatus503 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafStatus503

`func (o *OriginInspectorMeasurements) SetWafStatus503(v int32)`

SetWafStatus503 sets WafStatus503 field to given value.

### HasWafStatus503

`func (o *OriginInspectorMeasurements) HasWafStatus503() bool`

HasWafStatus503 returns a boolean if a field has been set.

### GetWafStatus504

`func (o *OriginInspectorMeasurements) GetWafStatus504() int32`

GetWafStatus504 returns the WafStatus504 field if non-nil, zero value otherwise.

### GetWafStatus504Ok

`func (o *OriginInspectorMeasurements) GetWafStatus504Ok() (*int32, bool)`

GetWafStatus504Ok returns a tuple with the WafStatus504 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafStatus504

`func (o *OriginInspectorMeasurements) SetWafStatus504(v int32)`

SetWafStatus504 sets WafStatus504 field to given value.

### HasWafStatus504

`func (o *OriginInspectorMeasurements) HasWafStatus504() bool`

HasWafStatus504 returns a boolean if a field has been set.

### GetWafStatus505

`func (o *OriginInspectorMeasurements) GetWafStatus505() int32`

GetWafStatus505 returns the WafStatus505 field if non-nil, zero value otherwise.

### GetWafStatus505Ok

`func (o *OriginInspectorMeasurements) GetWafStatus505Ok() (*int32, bool)`

GetWafStatus505Ok returns a tuple with the WafStatus505 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafStatus505

`func (o *OriginInspectorMeasurements) SetWafStatus505(v int32)`

SetWafStatus505 sets WafStatus505 field to given value.

### HasWafStatus505

`func (o *OriginInspectorMeasurements) HasWafStatus505() bool`

HasWafStatus505 returns a boolean if a field has been set.

### GetWafLatency0To1ms

`func (o *OriginInspectorMeasurements) GetWafLatency0To1ms() int32`

GetWafLatency0To1ms returns the WafLatency0To1ms field if non-nil, zero value otherwise.

### GetWafLatency0To1msOk

`func (o *OriginInspectorMeasurements) GetWafLatency0To1msOk() (*int32, bool)`

GetWafLatency0To1msOk returns a tuple with the WafLatency0To1ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafLatency0To1ms

`func (o *OriginInspectorMeasurements) SetWafLatency0To1ms(v int32)`

SetWafLatency0To1ms sets WafLatency0To1ms field to given value.

### HasWafLatency0To1ms

`func (o *OriginInspectorMeasurements) HasWafLatency0To1ms() bool`

HasWafLatency0To1ms returns a boolean if a field has been set.

### GetWafLatency1To5ms

`func (o *OriginInspectorMeasurements) GetWafLatency1To5ms() int32`

GetWafLatency1To5ms returns the WafLatency1To5ms field if non-nil, zero value otherwise.

### GetWafLatency1To5msOk

`func (o *OriginInspectorMeasurements) GetWafLatency1To5msOk() (*int32, bool)`

GetWafLatency1To5msOk returns a tuple with the WafLatency1To5ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafLatency1To5ms

`func (o *OriginInspectorMeasurements) SetWafLatency1To5ms(v int32)`

SetWafLatency1To5ms sets WafLatency1To5ms field to given value.

### HasWafLatency1To5ms

`func (o *OriginInspectorMeasurements) HasWafLatency1To5ms() bool`

HasWafLatency1To5ms returns a boolean if a field has been set.

### GetWafLatency5To10ms

`func (o *OriginInspectorMeasurements) GetWafLatency5To10ms() int32`

GetWafLatency5To10ms returns the WafLatency5To10ms field if non-nil, zero value otherwise.

### GetWafLatency5To10msOk

`func (o *OriginInspectorMeasurements) GetWafLatency5To10msOk() (*int32, bool)`

GetWafLatency5To10msOk returns a tuple with the WafLatency5To10ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafLatency5To10ms

`func (o *OriginInspectorMeasurements) SetWafLatency5To10ms(v int32)`

SetWafLatency5To10ms sets WafLatency5To10ms field to given value.

### HasWafLatency5To10ms

`func (o *OriginInspectorMeasurements) HasWafLatency5To10ms() bool`

HasWafLatency5To10ms returns a boolean if a field has been set.

### GetWafLatency10To50ms

`func (o *OriginInspectorMeasurements) GetWafLatency10To50ms() int32`

GetWafLatency10To50ms returns the WafLatency10To50ms field if non-nil, zero value otherwise.

### GetWafLatency10To50msOk

`func (o *OriginInspectorMeasurements) GetWafLatency10To50msOk() (*int32, bool)`

GetWafLatency10To50msOk returns a tuple with the WafLatency10To50ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafLatency10To50ms

`func (o *OriginInspectorMeasurements) SetWafLatency10To50ms(v int32)`

SetWafLatency10To50ms sets WafLatency10To50ms field to given value.

### HasWafLatency10To50ms

`func (o *OriginInspectorMeasurements) HasWafLatency10To50ms() bool`

HasWafLatency10To50ms returns a boolean if a field has been set.

### GetWafLatency50To100ms

`func (o *OriginInspectorMeasurements) GetWafLatency50To100ms() int32`

GetWafLatency50To100ms returns the WafLatency50To100ms field if non-nil, zero value otherwise.

### GetWafLatency50To100msOk

`func (o *OriginInspectorMeasurements) GetWafLatency50To100msOk() (*int32, bool)`

GetWafLatency50To100msOk returns a tuple with the WafLatency50To100ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafLatency50To100ms

`func (o *OriginInspectorMeasurements) SetWafLatency50To100ms(v int32)`

SetWafLatency50To100ms sets WafLatency50To100ms field to given value.

### HasWafLatency50To100ms

`func (o *OriginInspectorMeasurements) HasWafLatency50To100ms() bool`

HasWafLatency50To100ms returns a boolean if a field has been set.

### GetWafLatency100To250ms

`func (o *OriginInspectorMeasurements) GetWafLatency100To250ms() int32`

GetWafLatency100To250ms returns the WafLatency100To250ms field if non-nil, zero value otherwise.

### GetWafLatency100To250msOk

`func (o *OriginInspectorMeasurements) GetWafLatency100To250msOk() (*int32, bool)`

GetWafLatency100To250msOk returns a tuple with the WafLatency100To250ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafLatency100To250ms

`func (o *OriginInspectorMeasurements) SetWafLatency100To250ms(v int32)`

SetWafLatency100To250ms sets WafLatency100To250ms field to given value.

### HasWafLatency100To250ms

`func (o *OriginInspectorMeasurements) HasWafLatency100To250ms() bool`

HasWafLatency100To250ms returns a boolean if a field has been set.

### GetWafLatency250To500ms

`func (o *OriginInspectorMeasurements) GetWafLatency250To500ms() int32`

GetWafLatency250To500ms returns the WafLatency250To500ms field if non-nil, zero value otherwise.

### GetWafLatency250To500msOk

`func (o *OriginInspectorMeasurements) GetWafLatency250To500msOk() (*int32, bool)`

GetWafLatency250To500msOk returns a tuple with the WafLatency250To500ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafLatency250To500ms

`func (o *OriginInspectorMeasurements) SetWafLatency250To500ms(v int32)`

SetWafLatency250To500ms sets WafLatency250To500ms field to given value.

### HasWafLatency250To500ms

`func (o *OriginInspectorMeasurements) HasWafLatency250To500ms() bool`

HasWafLatency250To500ms returns a boolean if a field has been set.

### GetWafLatency500To1000ms

`func (o *OriginInspectorMeasurements) GetWafLatency500To1000ms() int32`

GetWafLatency500To1000ms returns the WafLatency500To1000ms field if non-nil, zero value otherwise.

### GetWafLatency500To1000msOk

`func (o *OriginInspectorMeasurements) GetWafLatency500To1000msOk() (*int32, bool)`

GetWafLatency500To1000msOk returns a tuple with the WafLatency500To1000ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafLatency500To1000ms

`func (o *OriginInspectorMeasurements) SetWafLatency500To1000ms(v int32)`

SetWafLatency500To1000ms sets WafLatency500To1000ms field to given value.

### HasWafLatency500To1000ms

`func (o *OriginInspectorMeasurements) HasWafLatency500To1000ms() bool`

HasWafLatency500To1000ms returns a boolean if a field has been set.

### GetWafLatency1000To5000ms

`func (o *OriginInspectorMeasurements) GetWafLatency1000To5000ms() int32`

GetWafLatency1000To5000ms returns the WafLatency1000To5000ms field if non-nil, zero value otherwise.

### GetWafLatency1000To5000msOk

`func (o *OriginInspectorMeasurements) GetWafLatency1000To5000msOk() (*int32, bool)`

GetWafLatency1000To5000msOk returns a tuple with the WafLatency1000To5000ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafLatency1000To5000ms

`func (o *OriginInspectorMeasurements) SetWafLatency1000To5000ms(v int32)`

SetWafLatency1000To5000ms sets WafLatency1000To5000ms field to given value.

### HasWafLatency1000To5000ms

`func (o *OriginInspectorMeasurements) HasWafLatency1000To5000ms() bool`

HasWafLatency1000To5000ms returns a boolean if a field has been set.

### GetWafLatency5000To10000ms

`func (o *OriginInspectorMeasurements) GetWafLatency5000To10000ms() int32`

GetWafLatency5000To10000ms returns the WafLatency5000To10000ms field if non-nil, zero value otherwise.

### GetWafLatency5000To10000msOk

`func (o *OriginInspectorMeasurements) GetWafLatency5000To10000msOk() (*int32, bool)`

GetWafLatency5000To10000msOk returns a tuple with the WafLatency5000To10000ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafLatency5000To10000ms

`func (o *OriginInspectorMeasurements) SetWafLatency5000To10000ms(v int32)`

SetWafLatency5000To10000ms sets WafLatency5000To10000ms field to given value.

### HasWafLatency5000To10000ms

`func (o *OriginInspectorMeasurements) HasWafLatency5000To10000ms() bool`

HasWafLatency5000To10000ms returns a boolean if a field has been set.

### GetWafLatency10000To60000ms

`func (o *OriginInspectorMeasurements) GetWafLatency10000To60000ms() int32`

GetWafLatency10000To60000ms returns the WafLatency10000To60000ms field if non-nil, zero value otherwise.

### GetWafLatency10000To60000msOk

`func (o *OriginInspectorMeasurements) GetWafLatency10000To60000msOk() (*int32, bool)`

GetWafLatency10000To60000msOk returns a tuple with the WafLatency10000To60000ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafLatency10000To60000ms

`func (o *OriginInspectorMeasurements) SetWafLatency10000To60000ms(v int32)`

SetWafLatency10000To60000ms sets WafLatency10000To60000ms field to given value.

### HasWafLatency10000To60000ms

`func (o *OriginInspectorMeasurements) HasWafLatency10000To60000ms() bool`

HasWafLatency10000To60000ms returns a boolean if a field has been set.

### GetWafLatency60000ms

`func (o *OriginInspectorMeasurements) GetWafLatency60000ms() int32`

GetWafLatency60000ms returns the WafLatency60000ms field if non-nil, zero value otherwise.

### GetWafLatency60000msOk

`func (o *OriginInspectorMeasurements) GetWafLatency60000msOk() (*int32, bool)`

GetWafLatency60000msOk returns a tuple with the WafLatency60000ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafLatency60000ms

`func (o *OriginInspectorMeasurements) SetWafLatency60000ms(v int32)`

SetWafLatency60000ms sets WafLatency60000ms field to given value.

### HasWafLatency60000ms

`func (o *OriginInspectorMeasurements) HasWafLatency60000ms() bool`

HasWafLatency60000ms returns a boolean if a field has been set.

### GetComputeResponses

`func (o *OriginInspectorMeasurements) GetComputeResponses() int32`

GetComputeResponses returns the ComputeResponses field if non-nil, zero value otherwise.

### GetComputeResponsesOk

`func (o *OriginInspectorMeasurements) GetComputeResponsesOk() (*int32, bool)`

GetComputeResponsesOk returns a tuple with the ComputeResponses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeResponses

`func (o *OriginInspectorMeasurements) SetComputeResponses(v int32)`

SetComputeResponses sets ComputeResponses field to given value.

### HasComputeResponses

`func (o *OriginInspectorMeasurements) HasComputeResponses() bool`

HasComputeResponses returns a boolean if a field has been set.

### GetComputeRespHeaderBytes

`func (o *OriginInspectorMeasurements) GetComputeRespHeaderBytes() int32`

GetComputeRespHeaderBytes returns the ComputeRespHeaderBytes field if non-nil, zero value otherwise.

### GetComputeRespHeaderBytesOk

`func (o *OriginInspectorMeasurements) GetComputeRespHeaderBytesOk() (*int32, bool)`

GetComputeRespHeaderBytesOk returns a tuple with the ComputeRespHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespHeaderBytes

`func (o *OriginInspectorMeasurements) SetComputeRespHeaderBytes(v int32)`

SetComputeRespHeaderBytes sets ComputeRespHeaderBytes field to given value.

### HasComputeRespHeaderBytes

`func (o *OriginInspectorMeasurements) HasComputeRespHeaderBytes() bool`

HasComputeRespHeaderBytes returns a boolean if a field has been set.

### GetComputeRespBodyBytes

`func (o *OriginInspectorMeasurements) GetComputeRespBodyBytes() int32`

GetComputeRespBodyBytes returns the ComputeRespBodyBytes field if non-nil, zero value otherwise.

### GetComputeRespBodyBytesOk

`func (o *OriginInspectorMeasurements) GetComputeRespBodyBytesOk() (*int32, bool)`

GetComputeRespBodyBytesOk returns a tuple with the ComputeRespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespBodyBytes

`func (o *OriginInspectorMeasurements) SetComputeRespBodyBytes(v int32)`

SetComputeRespBodyBytes sets ComputeRespBodyBytes field to given value.

### HasComputeRespBodyBytes

`func (o *OriginInspectorMeasurements) HasComputeRespBodyBytes() bool`

HasComputeRespBodyBytes returns a boolean if a field has been set.

### GetComputeStatus1xx

`func (o *OriginInspectorMeasurements) GetComputeStatus1xx() int32`

GetComputeStatus1xx returns the ComputeStatus1xx field if non-nil, zero value otherwise.

### GetComputeStatus1xxOk

`func (o *OriginInspectorMeasurements) GetComputeStatus1xxOk() (*int32, bool)`

GetComputeStatus1xxOk returns a tuple with the ComputeStatus1xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeStatus1xx

`func (o *OriginInspectorMeasurements) SetComputeStatus1xx(v int32)`

SetComputeStatus1xx sets ComputeStatus1xx field to given value.

### HasComputeStatus1xx

`func (o *OriginInspectorMeasurements) HasComputeStatus1xx() bool`

HasComputeStatus1xx returns a boolean if a field has been set.

### GetComputeStatus2xx

`func (o *OriginInspectorMeasurements) GetComputeStatus2xx() int32`

GetComputeStatus2xx returns the ComputeStatus2xx field if non-nil, zero value otherwise.

### GetComputeStatus2xxOk

`func (o *OriginInspectorMeasurements) GetComputeStatus2xxOk() (*int32, bool)`

GetComputeStatus2xxOk returns a tuple with the ComputeStatus2xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeStatus2xx

`func (o *OriginInspectorMeasurements) SetComputeStatus2xx(v int32)`

SetComputeStatus2xx sets ComputeStatus2xx field to given value.

### HasComputeStatus2xx

`func (o *OriginInspectorMeasurements) HasComputeStatus2xx() bool`

HasComputeStatus2xx returns a boolean if a field has been set.

### GetComputeStatus3xx

`func (o *OriginInspectorMeasurements) GetComputeStatus3xx() int32`

GetComputeStatus3xx returns the ComputeStatus3xx field if non-nil, zero value otherwise.

### GetComputeStatus3xxOk

`func (o *OriginInspectorMeasurements) GetComputeStatus3xxOk() (*int32, bool)`

GetComputeStatus3xxOk returns a tuple with the ComputeStatus3xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeStatus3xx

`func (o *OriginInspectorMeasurements) SetComputeStatus3xx(v int32)`

SetComputeStatus3xx sets ComputeStatus3xx field to given value.

### HasComputeStatus3xx

`func (o *OriginInspectorMeasurements) HasComputeStatus3xx() bool`

HasComputeStatus3xx returns a boolean if a field has been set.

### GetComputeStatus4xx

`func (o *OriginInspectorMeasurements) GetComputeStatus4xx() int32`

GetComputeStatus4xx returns the ComputeStatus4xx field if non-nil, zero value otherwise.

### GetComputeStatus4xxOk

`func (o *OriginInspectorMeasurements) GetComputeStatus4xxOk() (*int32, bool)`

GetComputeStatus4xxOk returns a tuple with the ComputeStatus4xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeStatus4xx

`func (o *OriginInspectorMeasurements) SetComputeStatus4xx(v int32)`

SetComputeStatus4xx sets ComputeStatus4xx field to given value.

### HasComputeStatus4xx

`func (o *OriginInspectorMeasurements) HasComputeStatus4xx() bool`

HasComputeStatus4xx returns a boolean if a field has been set.

### GetComputeStatus5xx

`func (o *OriginInspectorMeasurements) GetComputeStatus5xx() int32`

GetComputeStatus5xx returns the ComputeStatus5xx field if non-nil, zero value otherwise.

### GetComputeStatus5xxOk

`func (o *OriginInspectorMeasurements) GetComputeStatus5xxOk() (*int32, bool)`

GetComputeStatus5xxOk returns a tuple with the ComputeStatus5xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeStatus5xx

`func (o *OriginInspectorMeasurements) SetComputeStatus5xx(v int32)`

SetComputeStatus5xx sets ComputeStatus5xx field to given value.

### HasComputeStatus5xx

`func (o *OriginInspectorMeasurements) HasComputeStatus5xx() bool`

HasComputeStatus5xx returns a boolean if a field has been set.

### GetComputeStatus200

`func (o *OriginInspectorMeasurements) GetComputeStatus200() int32`

GetComputeStatus200 returns the ComputeStatus200 field if non-nil, zero value otherwise.

### GetComputeStatus200Ok

`func (o *OriginInspectorMeasurements) GetComputeStatus200Ok() (*int32, bool)`

GetComputeStatus200Ok returns a tuple with the ComputeStatus200 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeStatus200

`func (o *OriginInspectorMeasurements) SetComputeStatus200(v int32)`

SetComputeStatus200 sets ComputeStatus200 field to given value.

### HasComputeStatus200

`func (o *OriginInspectorMeasurements) HasComputeStatus200() bool`

HasComputeStatus200 returns a boolean if a field has been set.

### GetComputeStatus204

`func (o *OriginInspectorMeasurements) GetComputeStatus204() int32`

GetComputeStatus204 returns the ComputeStatus204 field if non-nil, zero value otherwise.

### GetComputeStatus204Ok

`func (o *OriginInspectorMeasurements) GetComputeStatus204Ok() (*int32, bool)`

GetComputeStatus204Ok returns a tuple with the ComputeStatus204 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeStatus204

`func (o *OriginInspectorMeasurements) SetComputeStatus204(v int32)`

SetComputeStatus204 sets ComputeStatus204 field to given value.

### HasComputeStatus204

`func (o *OriginInspectorMeasurements) HasComputeStatus204() bool`

HasComputeStatus204 returns a boolean if a field has been set.

### GetComputeStatus206

`func (o *OriginInspectorMeasurements) GetComputeStatus206() int32`

GetComputeStatus206 returns the ComputeStatus206 field if non-nil, zero value otherwise.

### GetComputeStatus206Ok

`func (o *OriginInspectorMeasurements) GetComputeStatus206Ok() (*int32, bool)`

GetComputeStatus206Ok returns a tuple with the ComputeStatus206 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeStatus206

`func (o *OriginInspectorMeasurements) SetComputeStatus206(v int32)`

SetComputeStatus206 sets ComputeStatus206 field to given value.

### HasComputeStatus206

`func (o *OriginInspectorMeasurements) HasComputeStatus206() bool`

HasComputeStatus206 returns a boolean if a field has been set.

### GetComputeStatus301

`func (o *OriginInspectorMeasurements) GetComputeStatus301() int32`

GetComputeStatus301 returns the ComputeStatus301 field if non-nil, zero value otherwise.

### GetComputeStatus301Ok

`func (o *OriginInspectorMeasurements) GetComputeStatus301Ok() (*int32, bool)`

GetComputeStatus301Ok returns a tuple with the ComputeStatus301 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeStatus301

`func (o *OriginInspectorMeasurements) SetComputeStatus301(v int32)`

SetComputeStatus301 sets ComputeStatus301 field to given value.

### HasComputeStatus301

`func (o *OriginInspectorMeasurements) HasComputeStatus301() bool`

HasComputeStatus301 returns a boolean if a field has been set.

### GetComputeStatus302

`func (o *OriginInspectorMeasurements) GetComputeStatus302() int32`

GetComputeStatus302 returns the ComputeStatus302 field if non-nil, zero value otherwise.

### GetComputeStatus302Ok

`func (o *OriginInspectorMeasurements) GetComputeStatus302Ok() (*int32, bool)`

GetComputeStatus302Ok returns a tuple with the ComputeStatus302 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeStatus302

`func (o *OriginInspectorMeasurements) SetComputeStatus302(v int32)`

SetComputeStatus302 sets ComputeStatus302 field to given value.

### HasComputeStatus302

`func (o *OriginInspectorMeasurements) HasComputeStatus302() bool`

HasComputeStatus302 returns a boolean if a field has been set.

### GetComputeStatus304

`func (o *OriginInspectorMeasurements) GetComputeStatus304() int32`

GetComputeStatus304 returns the ComputeStatus304 field if non-nil, zero value otherwise.

### GetComputeStatus304Ok

`func (o *OriginInspectorMeasurements) GetComputeStatus304Ok() (*int32, bool)`

GetComputeStatus304Ok returns a tuple with the ComputeStatus304 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeStatus304

`func (o *OriginInspectorMeasurements) SetComputeStatus304(v int32)`

SetComputeStatus304 sets ComputeStatus304 field to given value.

### HasComputeStatus304

`func (o *OriginInspectorMeasurements) HasComputeStatus304() bool`

HasComputeStatus304 returns a boolean if a field has been set.

### GetComputeStatus400

`func (o *OriginInspectorMeasurements) GetComputeStatus400() int32`

GetComputeStatus400 returns the ComputeStatus400 field if non-nil, zero value otherwise.

### GetComputeStatus400Ok

`func (o *OriginInspectorMeasurements) GetComputeStatus400Ok() (*int32, bool)`

GetComputeStatus400Ok returns a tuple with the ComputeStatus400 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeStatus400

`func (o *OriginInspectorMeasurements) SetComputeStatus400(v int32)`

SetComputeStatus400 sets ComputeStatus400 field to given value.

### HasComputeStatus400

`func (o *OriginInspectorMeasurements) HasComputeStatus400() bool`

HasComputeStatus400 returns a boolean if a field has been set.

### GetComputeStatus401

`func (o *OriginInspectorMeasurements) GetComputeStatus401() int32`

GetComputeStatus401 returns the ComputeStatus401 field if non-nil, zero value otherwise.

### GetComputeStatus401Ok

`func (o *OriginInspectorMeasurements) GetComputeStatus401Ok() (*int32, bool)`

GetComputeStatus401Ok returns a tuple with the ComputeStatus401 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeStatus401

`func (o *OriginInspectorMeasurements) SetComputeStatus401(v int32)`

SetComputeStatus401 sets ComputeStatus401 field to given value.

### HasComputeStatus401

`func (o *OriginInspectorMeasurements) HasComputeStatus401() bool`

HasComputeStatus401 returns a boolean if a field has been set.

### GetComputeStatus403

`func (o *OriginInspectorMeasurements) GetComputeStatus403() int32`

GetComputeStatus403 returns the ComputeStatus403 field if non-nil, zero value otherwise.

### GetComputeStatus403Ok

`func (o *OriginInspectorMeasurements) GetComputeStatus403Ok() (*int32, bool)`

GetComputeStatus403Ok returns a tuple with the ComputeStatus403 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeStatus403

`func (o *OriginInspectorMeasurements) SetComputeStatus403(v int32)`

SetComputeStatus403 sets ComputeStatus403 field to given value.

### HasComputeStatus403

`func (o *OriginInspectorMeasurements) HasComputeStatus403() bool`

HasComputeStatus403 returns a boolean if a field has been set.

### GetComputeStatus404

`func (o *OriginInspectorMeasurements) GetComputeStatus404() int32`

GetComputeStatus404 returns the ComputeStatus404 field if non-nil, zero value otherwise.

### GetComputeStatus404Ok

`func (o *OriginInspectorMeasurements) GetComputeStatus404Ok() (*int32, bool)`

GetComputeStatus404Ok returns a tuple with the ComputeStatus404 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeStatus404

`func (o *OriginInspectorMeasurements) SetComputeStatus404(v int32)`

SetComputeStatus404 sets ComputeStatus404 field to given value.

### HasComputeStatus404

`func (o *OriginInspectorMeasurements) HasComputeStatus404() bool`

HasComputeStatus404 returns a boolean if a field has been set.

### GetComputeStatus416

`func (o *OriginInspectorMeasurements) GetComputeStatus416() int32`

GetComputeStatus416 returns the ComputeStatus416 field if non-nil, zero value otherwise.

### GetComputeStatus416Ok

`func (o *OriginInspectorMeasurements) GetComputeStatus416Ok() (*int32, bool)`

GetComputeStatus416Ok returns a tuple with the ComputeStatus416 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeStatus416

`func (o *OriginInspectorMeasurements) SetComputeStatus416(v int32)`

SetComputeStatus416 sets ComputeStatus416 field to given value.

### HasComputeStatus416

`func (o *OriginInspectorMeasurements) HasComputeStatus416() bool`

HasComputeStatus416 returns a boolean if a field has been set.

### GetComputeStatus429

`func (o *OriginInspectorMeasurements) GetComputeStatus429() int32`

GetComputeStatus429 returns the ComputeStatus429 field if non-nil, zero value otherwise.

### GetComputeStatus429Ok

`func (o *OriginInspectorMeasurements) GetComputeStatus429Ok() (*int32, bool)`

GetComputeStatus429Ok returns a tuple with the ComputeStatus429 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeStatus429

`func (o *OriginInspectorMeasurements) SetComputeStatus429(v int32)`

SetComputeStatus429 sets ComputeStatus429 field to given value.

### HasComputeStatus429

`func (o *OriginInspectorMeasurements) HasComputeStatus429() bool`

HasComputeStatus429 returns a boolean if a field has been set.

### GetComputeStatus500

`func (o *OriginInspectorMeasurements) GetComputeStatus500() int32`

GetComputeStatus500 returns the ComputeStatus500 field if non-nil, zero value otherwise.

### GetComputeStatus500Ok

`func (o *OriginInspectorMeasurements) GetComputeStatus500Ok() (*int32, bool)`

GetComputeStatus500Ok returns a tuple with the ComputeStatus500 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeStatus500

`func (o *OriginInspectorMeasurements) SetComputeStatus500(v int32)`

SetComputeStatus500 sets ComputeStatus500 field to given value.

### HasComputeStatus500

`func (o *OriginInspectorMeasurements) HasComputeStatus500() bool`

HasComputeStatus500 returns a boolean if a field has been set.

### GetComputeStatus501

`func (o *OriginInspectorMeasurements) GetComputeStatus501() int32`

GetComputeStatus501 returns the ComputeStatus501 field if non-nil, zero value otherwise.

### GetComputeStatus501Ok

`func (o *OriginInspectorMeasurements) GetComputeStatus501Ok() (*int32, bool)`

GetComputeStatus501Ok returns a tuple with the ComputeStatus501 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeStatus501

`func (o *OriginInspectorMeasurements) SetComputeStatus501(v int32)`

SetComputeStatus501 sets ComputeStatus501 field to given value.

### HasComputeStatus501

`func (o *OriginInspectorMeasurements) HasComputeStatus501() bool`

HasComputeStatus501 returns a boolean if a field has been set.

### GetComputeStatus502

`func (o *OriginInspectorMeasurements) GetComputeStatus502() int32`

GetComputeStatus502 returns the ComputeStatus502 field if non-nil, zero value otherwise.

### GetComputeStatus502Ok

`func (o *OriginInspectorMeasurements) GetComputeStatus502Ok() (*int32, bool)`

GetComputeStatus502Ok returns a tuple with the ComputeStatus502 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeStatus502

`func (o *OriginInspectorMeasurements) SetComputeStatus502(v int32)`

SetComputeStatus502 sets ComputeStatus502 field to given value.

### HasComputeStatus502

`func (o *OriginInspectorMeasurements) HasComputeStatus502() bool`

HasComputeStatus502 returns a boolean if a field has been set.

### GetComputeStatus503

`func (o *OriginInspectorMeasurements) GetComputeStatus503() int32`

GetComputeStatus503 returns the ComputeStatus503 field if non-nil, zero value otherwise.

### GetComputeStatus503Ok

`func (o *OriginInspectorMeasurements) GetComputeStatus503Ok() (*int32, bool)`

GetComputeStatus503Ok returns a tuple with the ComputeStatus503 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeStatus503

`func (o *OriginInspectorMeasurements) SetComputeStatus503(v int32)`

SetComputeStatus503 sets ComputeStatus503 field to given value.

### HasComputeStatus503

`func (o *OriginInspectorMeasurements) HasComputeStatus503() bool`

HasComputeStatus503 returns a boolean if a field has been set.

### GetComputeStatus504

`func (o *OriginInspectorMeasurements) GetComputeStatus504() int32`

GetComputeStatus504 returns the ComputeStatus504 field if non-nil, zero value otherwise.

### GetComputeStatus504Ok

`func (o *OriginInspectorMeasurements) GetComputeStatus504Ok() (*int32, bool)`

GetComputeStatus504Ok returns a tuple with the ComputeStatus504 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeStatus504

`func (o *OriginInspectorMeasurements) SetComputeStatus504(v int32)`

SetComputeStatus504 sets ComputeStatus504 field to given value.

### HasComputeStatus504

`func (o *OriginInspectorMeasurements) HasComputeStatus504() bool`

HasComputeStatus504 returns a boolean if a field has been set.

### GetComputeStatus505

`func (o *OriginInspectorMeasurements) GetComputeStatus505() int32`

GetComputeStatus505 returns the ComputeStatus505 field if non-nil, zero value otherwise.

### GetComputeStatus505Ok

`func (o *OriginInspectorMeasurements) GetComputeStatus505Ok() (*int32, bool)`

GetComputeStatus505Ok returns a tuple with the ComputeStatus505 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeStatus505

`func (o *OriginInspectorMeasurements) SetComputeStatus505(v int32)`

SetComputeStatus505 sets ComputeStatus505 field to given value.

### HasComputeStatus505

`func (o *OriginInspectorMeasurements) HasComputeStatus505() bool`

HasComputeStatus505 returns a boolean if a field has been set.

### GetComputeLatency0To1ms

`func (o *OriginInspectorMeasurements) GetComputeLatency0To1ms() int32`

GetComputeLatency0To1ms returns the ComputeLatency0To1ms field if non-nil, zero value otherwise.

### GetComputeLatency0To1msOk

`func (o *OriginInspectorMeasurements) GetComputeLatency0To1msOk() (*int32, bool)`

GetComputeLatency0To1msOk returns a tuple with the ComputeLatency0To1ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeLatency0To1ms

`func (o *OriginInspectorMeasurements) SetComputeLatency0To1ms(v int32)`

SetComputeLatency0To1ms sets ComputeLatency0To1ms field to given value.

### HasComputeLatency0To1ms

`func (o *OriginInspectorMeasurements) HasComputeLatency0To1ms() bool`

HasComputeLatency0To1ms returns a boolean if a field has been set.

### GetComputeLatency1To5ms

`func (o *OriginInspectorMeasurements) GetComputeLatency1To5ms() int32`

GetComputeLatency1To5ms returns the ComputeLatency1To5ms field if non-nil, zero value otherwise.

### GetComputeLatency1To5msOk

`func (o *OriginInspectorMeasurements) GetComputeLatency1To5msOk() (*int32, bool)`

GetComputeLatency1To5msOk returns a tuple with the ComputeLatency1To5ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeLatency1To5ms

`func (o *OriginInspectorMeasurements) SetComputeLatency1To5ms(v int32)`

SetComputeLatency1To5ms sets ComputeLatency1To5ms field to given value.

### HasComputeLatency1To5ms

`func (o *OriginInspectorMeasurements) HasComputeLatency1To5ms() bool`

HasComputeLatency1To5ms returns a boolean if a field has been set.

### GetComputeLatency5To10ms

`func (o *OriginInspectorMeasurements) GetComputeLatency5To10ms() int32`

GetComputeLatency5To10ms returns the ComputeLatency5To10ms field if non-nil, zero value otherwise.

### GetComputeLatency5To10msOk

`func (o *OriginInspectorMeasurements) GetComputeLatency5To10msOk() (*int32, bool)`

GetComputeLatency5To10msOk returns a tuple with the ComputeLatency5To10ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeLatency5To10ms

`func (o *OriginInspectorMeasurements) SetComputeLatency5To10ms(v int32)`

SetComputeLatency5To10ms sets ComputeLatency5To10ms field to given value.

### HasComputeLatency5To10ms

`func (o *OriginInspectorMeasurements) HasComputeLatency5To10ms() bool`

HasComputeLatency5To10ms returns a boolean if a field has been set.

### GetComputeLatency10To50ms

`func (o *OriginInspectorMeasurements) GetComputeLatency10To50ms() int32`

GetComputeLatency10To50ms returns the ComputeLatency10To50ms field if non-nil, zero value otherwise.

### GetComputeLatency10To50msOk

`func (o *OriginInspectorMeasurements) GetComputeLatency10To50msOk() (*int32, bool)`

GetComputeLatency10To50msOk returns a tuple with the ComputeLatency10To50ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeLatency10To50ms

`func (o *OriginInspectorMeasurements) SetComputeLatency10To50ms(v int32)`

SetComputeLatency10To50ms sets ComputeLatency10To50ms field to given value.

### HasComputeLatency10To50ms

`func (o *OriginInspectorMeasurements) HasComputeLatency10To50ms() bool`

HasComputeLatency10To50ms returns a boolean if a field has been set.

### GetComputeLatency50To100ms

`func (o *OriginInspectorMeasurements) GetComputeLatency50To100ms() int32`

GetComputeLatency50To100ms returns the ComputeLatency50To100ms field if non-nil, zero value otherwise.

### GetComputeLatency50To100msOk

`func (o *OriginInspectorMeasurements) GetComputeLatency50To100msOk() (*int32, bool)`

GetComputeLatency50To100msOk returns a tuple with the ComputeLatency50To100ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeLatency50To100ms

`func (o *OriginInspectorMeasurements) SetComputeLatency50To100ms(v int32)`

SetComputeLatency50To100ms sets ComputeLatency50To100ms field to given value.

### HasComputeLatency50To100ms

`func (o *OriginInspectorMeasurements) HasComputeLatency50To100ms() bool`

HasComputeLatency50To100ms returns a boolean if a field has been set.

### GetComputeLatency100To250ms

`func (o *OriginInspectorMeasurements) GetComputeLatency100To250ms() int32`

GetComputeLatency100To250ms returns the ComputeLatency100To250ms field if non-nil, zero value otherwise.

### GetComputeLatency100To250msOk

`func (o *OriginInspectorMeasurements) GetComputeLatency100To250msOk() (*int32, bool)`

GetComputeLatency100To250msOk returns a tuple with the ComputeLatency100To250ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeLatency100To250ms

`func (o *OriginInspectorMeasurements) SetComputeLatency100To250ms(v int32)`

SetComputeLatency100To250ms sets ComputeLatency100To250ms field to given value.

### HasComputeLatency100To250ms

`func (o *OriginInspectorMeasurements) HasComputeLatency100To250ms() bool`

HasComputeLatency100To250ms returns a boolean if a field has been set.

### GetComputeLatency250To500ms

`func (o *OriginInspectorMeasurements) GetComputeLatency250To500ms() int32`

GetComputeLatency250To500ms returns the ComputeLatency250To500ms field if non-nil, zero value otherwise.

### GetComputeLatency250To500msOk

`func (o *OriginInspectorMeasurements) GetComputeLatency250To500msOk() (*int32, bool)`

GetComputeLatency250To500msOk returns a tuple with the ComputeLatency250To500ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeLatency250To500ms

`func (o *OriginInspectorMeasurements) SetComputeLatency250To500ms(v int32)`

SetComputeLatency250To500ms sets ComputeLatency250To500ms field to given value.

### HasComputeLatency250To500ms

`func (o *OriginInspectorMeasurements) HasComputeLatency250To500ms() bool`

HasComputeLatency250To500ms returns a boolean if a field has been set.

### GetComputeLatency500To1000ms

`func (o *OriginInspectorMeasurements) GetComputeLatency500To1000ms() int32`

GetComputeLatency500To1000ms returns the ComputeLatency500To1000ms field if non-nil, zero value otherwise.

### GetComputeLatency500To1000msOk

`func (o *OriginInspectorMeasurements) GetComputeLatency500To1000msOk() (*int32, bool)`

GetComputeLatency500To1000msOk returns a tuple with the ComputeLatency500To1000ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeLatency500To1000ms

`func (o *OriginInspectorMeasurements) SetComputeLatency500To1000ms(v int32)`

SetComputeLatency500To1000ms sets ComputeLatency500To1000ms field to given value.

### HasComputeLatency500To1000ms

`func (o *OriginInspectorMeasurements) HasComputeLatency500To1000ms() bool`

HasComputeLatency500To1000ms returns a boolean if a field has been set.

### GetComputeLatency1000To5000ms

`func (o *OriginInspectorMeasurements) GetComputeLatency1000To5000ms() int32`

GetComputeLatency1000To5000ms returns the ComputeLatency1000To5000ms field if non-nil, zero value otherwise.

### GetComputeLatency1000To5000msOk

`func (o *OriginInspectorMeasurements) GetComputeLatency1000To5000msOk() (*int32, bool)`

GetComputeLatency1000To5000msOk returns a tuple with the ComputeLatency1000To5000ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeLatency1000To5000ms

`func (o *OriginInspectorMeasurements) SetComputeLatency1000To5000ms(v int32)`

SetComputeLatency1000To5000ms sets ComputeLatency1000To5000ms field to given value.

### HasComputeLatency1000To5000ms

`func (o *OriginInspectorMeasurements) HasComputeLatency1000To5000ms() bool`

HasComputeLatency1000To5000ms returns a boolean if a field has been set.

### GetComputeLatency5000To10000ms

`func (o *OriginInspectorMeasurements) GetComputeLatency5000To10000ms() int32`

GetComputeLatency5000To10000ms returns the ComputeLatency5000To10000ms field if non-nil, zero value otherwise.

### GetComputeLatency5000To10000msOk

`func (o *OriginInspectorMeasurements) GetComputeLatency5000To10000msOk() (*int32, bool)`

GetComputeLatency5000To10000msOk returns a tuple with the ComputeLatency5000To10000ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeLatency5000To10000ms

`func (o *OriginInspectorMeasurements) SetComputeLatency5000To10000ms(v int32)`

SetComputeLatency5000To10000ms sets ComputeLatency5000To10000ms field to given value.

### HasComputeLatency5000To10000ms

`func (o *OriginInspectorMeasurements) HasComputeLatency5000To10000ms() bool`

HasComputeLatency5000To10000ms returns a boolean if a field has been set.

### GetComputeLatency10000To60000ms

`func (o *OriginInspectorMeasurements) GetComputeLatency10000To60000ms() int32`

GetComputeLatency10000To60000ms returns the ComputeLatency10000To60000ms field if non-nil, zero value otherwise.

### GetComputeLatency10000To60000msOk

`func (o *OriginInspectorMeasurements) GetComputeLatency10000To60000msOk() (*int32, bool)`

GetComputeLatency10000To60000msOk returns a tuple with the ComputeLatency10000To60000ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeLatency10000To60000ms

`func (o *OriginInspectorMeasurements) SetComputeLatency10000To60000ms(v int32)`

SetComputeLatency10000To60000ms sets ComputeLatency10000To60000ms field to given value.

### HasComputeLatency10000To60000ms

`func (o *OriginInspectorMeasurements) HasComputeLatency10000To60000ms() bool`

HasComputeLatency10000To60000ms returns a boolean if a field has been set.

### GetComputeLatency60000ms

`func (o *OriginInspectorMeasurements) GetComputeLatency60000ms() int32`

GetComputeLatency60000ms returns the ComputeLatency60000ms field if non-nil, zero value otherwise.

### GetComputeLatency60000msOk

`func (o *OriginInspectorMeasurements) GetComputeLatency60000msOk() (*int32, bool)`

GetComputeLatency60000msOk returns a tuple with the ComputeLatency60000ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeLatency60000ms

`func (o *OriginInspectorMeasurements) SetComputeLatency60000ms(v int32)`

SetComputeLatency60000ms sets ComputeLatency60000ms field to given value.

### HasComputeLatency60000ms

`func (o *OriginInspectorMeasurements) HasComputeLatency60000ms() bool`

HasComputeLatency60000ms returns a boolean if a field has been set.

### GetAllResponses

`func (o *OriginInspectorMeasurements) GetAllResponses() int32`

GetAllResponses returns the AllResponses field if non-nil, zero value otherwise.

### GetAllResponsesOk

`func (o *OriginInspectorMeasurements) GetAllResponsesOk() (*int32, bool)`

GetAllResponsesOk returns a tuple with the AllResponses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllResponses

`func (o *OriginInspectorMeasurements) SetAllResponses(v int32)`

SetAllResponses sets AllResponses field to given value.

### HasAllResponses

`func (o *OriginInspectorMeasurements) HasAllResponses() bool`

HasAllResponses returns a boolean if a field has been set.

### GetAllRespHeaderBytes

`func (o *OriginInspectorMeasurements) GetAllRespHeaderBytes() int32`

GetAllRespHeaderBytes returns the AllRespHeaderBytes field if non-nil, zero value otherwise.

### GetAllRespHeaderBytesOk

`func (o *OriginInspectorMeasurements) GetAllRespHeaderBytesOk() (*int32, bool)`

GetAllRespHeaderBytesOk returns a tuple with the AllRespHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllRespHeaderBytes

`func (o *OriginInspectorMeasurements) SetAllRespHeaderBytes(v int32)`

SetAllRespHeaderBytes sets AllRespHeaderBytes field to given value.

### HasAllRespHeaderBytes

`func (o *OriginInspectorMeasurements) HasAllRespHeaderBytes() bool`

HasAllRespHeaderBytes returns a boolean if a field has been set.

### GetAllRespBodyBytes

`func (o *OriginInspectorMeasurements) GetAllRespBodyBytes() int32`

GetAllRespBodyBytes returns the AllRespBodyBytes field if non-nil, zero value otherwise.

### GetAllRespBodyBytesOk

`func (o *OriginInspectorMeasurements) GetAllRespBodyBytesOk() (*int32, bool)`

GetAllRespBodyBytesOk returns a tuple with the AllRespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllRespBodyBytes

`func (o *OriginInspectorMeasurements) SetAllRespBodyBytes(v int32)`

SetAllRespBodyBytes sets AllRespBodyBytes field to given value.

### HasAllRespBodyBytes

`func (o *OriginInspectorMeasurements) HasAllRespBodyBytes() bool`

HasAllRespBodyBytes returns a boolean if a field has been set.

### GetAllStatus1xx

`func (o *OriginInspectorMeasurements) GetAllStatus1xx() int32`

GetAllStatus1xx returns the AllStatus1xx field if non-nil, zero value otherwise.

### GetAllStatus1xxOk

`func (o *OriginInspectorMeasurements) GetAllStatus1xxOk() (*int32, bool)`

GetAllStatus1xxOk returns a tuple with the AllStatus1xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllStatus1xx

`func (o *OriginInspectorMeasurements) SetAllStatus1xx(v int32)`

SetAllStatus1xx sets AllStatus1xx field to given value.

### HasAllStatus1xx

`func (o *OriginInspectorMeasurements) HasAllStatus1xx() bool`

HasAllStatus1xx returns a boolean if a field has been set.

### GetAllStatus2xx

`func (o *OriginInspectorMeasurements) GetAllStatus2xx() int32`

GetAllStatus2xx returns the AllStatus2xx field if non-nil, zero value otherwise.

### GetAllStatus2xxOk

`func (o *OriginInspectorMeasurements) GetAllStatus2xxOk() (*int32, bool)`

GetAllStatus2xxOk returns a tuple with the AllStatus2xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllStatus2xx

`func (o *OriginInspectorMeasurements) SetAllStatus2xx(v int32)`

SetAllStatus2xx sets AllStatus2xx field to given value.

### HasAllStatus2xx

`func (o *OriginInspectorMeasurements) HasAllStatus2xx() bool`

HasAllStatus2xx returns a boolean if a field has been set.

### GetAllStatus3xx

`func (o *OriginInspectorMeasurements) GetAllStatus3xx() int32`

GetAllStatus3xx returns the AllStatus3xx field if non-nil, zero value otherwise.

### GetAllStatus3xxOk

`func (o *OriginInspectorMeasurements) GetAllStatus3xxOk() (*int32, bool)`

GetAllStatus3xxOk returns a tuple with the AllStatus3xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllStatus3xx

`func (o *OriginInspectorMeasurements) SetAllStatus3xx(v int32)`

SetAllStatus3xx sets AllStatus3xx field to given value.

### HasAllStatus3xx

`func (o *OriginInspectorMeasurements) HasAllStatus3xx() bool`

HasAllStatus3xx returns a boolean if a field has been set.

### GetAllStatus4xx

`func (o *OriginInspectorMeasurements) GetAllStatus4xx() int32`

GetAllStatus4xx returns the AllStatus4xx field if non-nil, zero value otherwise.

### GetAllStatus4xxOk

`func (o *OriginInspectorMeasurements) GetAllStatus4xxOk() (*int32, bool)`

GetAllStatus4xxOk returns a tuple with the AllStatus4xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllStatus4xx

`func (o *OriginInspectorMeasurements) SetAllStatus4xx(v int32)`

SetAllStatus4xx sets AllStatus4xx field to given value.

### HasAllStatus4xx

`func (o *OriginInspectorMeasurements) HasAllStatus4xx() bool`

HasAllStatus4xx returns a boolean if a field has been set.

### GetAllStatus5xx

`func (o *OriginInspectorMeasurements) GetAllStatus5xx() int32`

GetAllStatus5xx returns the AllStatus5xx field if non-nil, zero value otherwise.

### GetAllStatus5xxOk

`func (o *OriginInspectorMeasurements) GetAllStatus5xxOk() (*int32, bool)`

GetAllStatus5xxOk returns a tuple with the AllStatus5xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllStatus5xx

`func (o *OriginInspectorMeasurements) SetAllStatus5xx(v int32)`

SetAllStatus5xx sets AllStatus5xx field to given value.

### HasAllStatus5xx

`func (o *OriginInspectorMeasurements) HasAllStatus5xx() bool`

HasAllStatus5xx returns a boolean if a field has been set.

### GetAllStatus200

`func (o *OriginInspectorMeasurements) GetAllStatus200() int32`

GetAllStatus200 returns the AllStatus200 field if non-nil, zero value otherwise.

### GetAllStatus200Ok

`func (o *OriginInspectorMeasurements) GetAllStatus200Ok() (*int32, bool)`

GetAllStatus200Ok returns a tuple with the AllStatus200 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllStatus200

`func (o *OriginInspectorMeasurements) SetAllStatus200(v int32)`

SetAllStatus200 sets AllStatus200 field to given value.

### HasAllStatus200

`func (o *OriginInspectorMeasurements) HasAllStatus200() bool`

HasAllStatus200 returns a boolean if a field has been set.

### GetAllStatus204

`func (o *OriginInspectorMeasurements) GetAllStatus204() int32`

GetAllStatus204 returns the AllStatus204 field if non-nil, zero value otherwise.

### GetAllStatus204Ok

`func (o *OriginInspectorMeasurements) GetAllStatus204Ok() (*int32, bool)`

GetAllStatus204Ok returns a tuple with the AllStatus204 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllStatus204

`func (o *OriginInspectorMeasurements) SetAllStatus204(v int32)`

SetAllStatus204 sets AllStatus204 field to given value.

### HasAllStatus204

`func (o *OriginInspectorMeasurements) HasAllStatus204() bool`

HasAllStatus204 returns a boolean if a field has been set.

### GetAllStatus206

`func (o *OriginInspectorMeasurements) GetAllStatus206() int32`

GetAllStatus206 returns the AllStatus206 field if non-nil, zero value otherwise.

### GetAllStatus206Ok

`func (o *OriginInspectorMeasurements) GetAllStatus206Ok() (*int32, bool)`

GetAllStatus206Ok returns a tuple with the AllStatus206 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllStatus206

`func (o *OriginInspectorMeasurements) SetAllStatus206(v int32)`

SetAllStatus206 sets AllStatus206 field to given value.

### HasAllStatus206

`func (o *OriginInspectorMeasurements) HasAllStatus206() bool`

HasAllStatus206 returns a boolean if a field has been set.

### GetAllStatus301

`func (o *OriginInspectorMeasurements) GetAllStatus301() int32`

GetAllStatus301 returns the AllStatus301 field if non-nil, zero value otherwise.

### GetAllStatus301Ok

`func (o *OriginInspectorMeasurements) GetAllStatus301Ok() (*int32, bool)`

GetAllStatus301Ok returns a tuple with the AllStatus301 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllStatus301

`func (o *OriginInspectorMeasurements) SetAllStatus301(v int32)`

SetAllStatus301 sets AllStatus301 field to given value.

### HasAllStatus301

`func (o *OriginInspectorMeasurements) HasAllStatus301() bool`

HasAllStatus301 returns a boolean if a field has been set.

### GetAllStatus302

`func (o *OriginInspectorMeasurements) GetAllStatus302() int32`

GetAllStatus302 returns the AllStatus302 field if non-nil, zero value otherwise.

### GetAllStatus302Ok

`func (o *OriginInspectorMeasurements) GetAllStatus302Ok() (*int32, bool)`

GetAllStatus302Ok returns a tuple with the AllStatus302 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllStatus302

`func (o *OriginInspectorMeasurements) SetAllStatus302(v int32)`

SetAllStatus302 sets AllStatus302 field to given value.

### HasAllStatus302

`func (o *OriginInspectorMeasurements) HasAllStatus302() bool`

HasAllStatus302 returns a boolean if a field has been set.

### GetAllStatus304

`func (o *OriginInspectorMeasurements) GetAllStatus304() int32`

GetAllStatus304 returns the AllStatus304 field if non-nil, zero value otherwise.

### GetAllStatus304Ok

`func (o *OriginInspectorMeasurements) GetAllStatus304Ok() (*int32, bool)`

GetAllStatus304Ok returns a tuple with the AllStatus304 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllStatus304

`func (o *OriginInspectorMeasurements) SetAllStatus304(v int32)`

SetAllStatus304 sets AllStatus304 field to given value.

### HasAllStatus304

`func (o *OriginInspectorMeasurements) HasAllStatus304() bool`

HasAllStatus304 returns a boolean if a field has been set.

### GetAllStatus400

`func (o *OriginInspectorMeasurements) GetAllStatus400() int32`

GetAllStatus400 returns the AllStatus400 field if non-nil, zero value otherwise.

### GetAllStatus400Ok

`func (o *OriginInspectorMeasurements) GetAllStatus400Ok() (*int32, bool)`

GetAllStatus400Ok returns a tuple with the AllStatus400 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllStatus400

`func (o *OriginInspectorMeasurements) SetAllStatus400(v int32)`

SetAllStatus400 sets AllStatus400 field to given value.

### HasAllStatus400

`func (o *OriginInspectorMeasurements) HasAllStatus400() bool`

HasAllStatus400 returns a boolean if a field has been set.

### GetAllStatus401

`func (o *OriginInspectorMeasurements) GetAllStatus401() int32`

GetAllStatus401 returns the AllStatus401 field if non-nil, zero value otherwise.

### GetAllStatus401Ok

`func (o *OriginInspectorMeasurements) GetAllStatus401Ok() (*int32, bool)`

GetAllStatus401Ok returns a tuple with the AllStatus401 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllStatus401

`func (o *OriginInspectorMeasurements) SetAllStatus401(v int32)`

SetAllStatus401 sets AllStatus401 field to given value.

### HasAllStatus401

`func (o *OriginInspectorMeasurements) HasAllStatus401() bool`

HasAllStatus401 returns a boolean if a field has been set.

### GetAllStatus403

`func (o *OriginInspectorMeasurements) GetAllStatus403() int32`

GetAllStatus403 returns the AllStatus403 field if non-nil, zero value otherwise.

### GetAllStatus403Ok

`func (o *OriginInspectorMeasurements) GetAllStatus403Ok() (*int32, bool)`

GetAllStatus403Ok returns a tuple with the AllStatus403 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllStatus403

`func (o *OriginInspectorMeasurements) SetAllStatus403(v int32)`

SetAllStatus403 sets AllStatus403 field to given value.

### HasAllStatus403

`func (o *OriginInspectorMeasurements) HasAllStatus403() bool`

HasAllStatus403 returns a boolean if a field has been set.

### GetAllStatus404

`func (o *OriginInspectorMeasurements) GetAllStatus404() int32`

GetAllStatus404 returns the AllStatus404 field if non-nil, zero value otherwise.

### GetAllStatus404Ok

`func (o *OriginInspectorMeasurements) GetAllStatus404Ok() (*int32, bool)`

GetAllStatus404Ok returns a tuple with the AllStatus404 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllStatus404

`func (o *OriginInspectorMeasurements) SetAllStatus404(v int32)`

SetAllStatus404 sets AllStatus404 field to given value.

### HasAllStatus404

`func (o *OriginInspectorMeasurements) HasAllStatus404() bool`

HasAllStatus404 returns a boolean if a field has been set.

### GetAllStatus416

`func (o *OriginInspectorMeasurements) GetAllStatus416() int32`

GetAllStatus416 returns the AllStatus416 field if non-nil, zero value otherwise.

### GetAllStatus416Ok

`func (o *OriginInspectorMeasurements) GetAllStatus416Ok() (*int32, bool)`

GetAllStatus416Ok returns a tuple with the AllStatus416 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllStatus416

`func (o *OriginInspectorMeasurements) SetAllStatus416(v int32)`

SetAllStatus416 sets AllStatus416 field to given value.

### HasAllStatus416

`func (o *OriginInspectorMeasurements) HasAllStatus416() bool`

HasAllStatus416 returns a boolean if a field has been set.

### GetAllStatus429

`func (o *OriginInspectorMeasurements) GetAllStatus429() int32`

GetAllStatus429 returns the AllStatus429 field if non-nil, zero value otherwise.

### GetAllStatus429Ok

`func (o *OriginInspectorMeasurements) GetAllStatus429Ok() (*int32, bool)`

GetAllStatus429Ok returns a tuple with the AllStatus429 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllStatus429

`func (o *OriginInspectorMeasurements) SetAllStatus429(v int32)`

SetAllStatus429 sets AllStatus429 field to given value.

### HasAllStatus429

`func (o *OriginInspectorMeasurements) HasAllStatus429() bool`

HasAllStatus429 returns a boolean if a field has been set.

### GetAllStatus500

`func (o *OriginInspectorMeasurements) GetAllStatus500() int32`

GetAllStatus500 returns the AllStatus500 field if non-nil, zero value otherwise.

### GetAllStatus500Ok

`func (o *OriginInspectorMeasurements) GetAllStatus500Ok() (*int32, bool)`

GetAllStatus500Ok returns a tuple with the AllStatus500 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllStatus500

`func (o *OriginInspectorMeasurements) SetAllStatus500(v int32)`

SetAllStatus500 sets AllStatus500 field to given value.

### HasAllStatus500

`func (o *OriginInspectorMeasurements) HasAllStatus500() bool`

HasAllStatus500 returns a boolean if a field has been set.

### GetAllStatus501

`func (o *OriginInspectorMeasurements) GetAllStatus501() int32`

GetAllStatus501 returns the AllStatus501 field if non-nil, zero value otherwise.

### GetAllStatus501Ok

`func (o *OriginInspectorMeasurements) GetAllStatus501Ok() (*int32, bool)`

GetAllStatus501Ok returns a tuple with the AllStatus501 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllStatus501

`func (o *OriginInspectorMeasurements) SetAllStatus501(v int32)`

SetAllStatus501 sets AllStatus501 field to given value.

### HasAllStatus501

`func (o *OriginInspectorMeasurements) HasAllStatus501() bool`

HasAllStatus501 returns a boolean if a field has been set.

### GetAllStatus502

`func (o *OriginInspectorMeasurements) GetAllStatus502() int32`

GetAllStatus502 returns the AllStatus502 field if non-nil, zero value otherwise.

### GetAllStatus502Ok

`func (o *OriginInspectorMeasurements) GetAllStatus502Ok() (*int32, bool)`

GetAllStatus502Ok returns a tuple with the AllStatus502 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllStatus502

`func (o *OriginInspectorMeasurements) SetAllStatus502(v int32)`

SetAllStatus502 sets AllStatus502 field to given value.

### HasAllStatus502

`func (o *OriginInspectorMeasurements) HasAllStatus502() bool`

HasAllStatus502 returns a boolean if a field has been set.

### GetAllStatus503

`func (o *OriginInspectorMeasurements) GetAllStatus503() int32`

GetAllStatus503 returns the AllStatus503 field if non-nil, zero value otherwise.

### GetAllStatus503Ok

`func (o *OriginInspectorMeasurements) GetAllStatus503Ok() (*int32, bool)`

GetAllStatus503Ok returns a tuple with the AllStatus503 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllStatus503

`func (o *OriginInspectorMeasurements) SetAllStatus503(v int32)`

SetAllStatus503 sets AllStatus503 field to given value.

### HasAllStatus503

`func (o *OriginInspectorMeasurements) HasAllStatus503() bool`

HasAllStatus503 returns a boolean if a field has been set.

### GetAllStatus504

`func (o *OriginInspectorMeasurements) GetAllStatus504() int32`

GetAllStatus504 returns the AllStatus504 field if non-nil, zero value otherwise.

### GetAllStatus504Ok

`func (o *OriginInspectorMeasurements) GetAllStatus504Ok() (*int32, bool)`

GetAllStatus504Ok returns a tuple with the AllStatus504 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllStatus504

`func (o *OriginInspectorMeasurements) SetAllStatus504(v int32)`

SetAllStatus504 sets AllStatus504 field to given value.

### HasAllStatus504

`func (o *OriginInspectorMeasurements) HasAllStatus504() bool`

HasAllStatus504 returns a boolean if a field has been set.

### GetAllStatus505

`func (o *OriginInspectorMeasurements) GetAllStatus505() int32`

GetAllStatus505 returns the AllStatus505 field if non-nil, zero value otherwise.

### GetAllStatus505Ok

`func (o *OriginInspectorMeasurements) GetAllStatus505Ok() (*int32, bool)`

GetAllStatus505Ok returns a tuple with the AllStatus505 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllStatus505

`func (o *OriginInspectorMeasurements) SetAllStatus505(v int32)`

SetAllStatus505 sets AllStatus505 field to given value.

### HasAllStatus505

`func (o *OriginInspectorMeasurements) HasAllStatus505() bool`

HasAllStatus505 returns a boolean if a field has been set.

### GetAllLatency0To1ms

`func (o *OriginInspectorMeasurements) GetAllLatency0To1ms() int32`

GetAllLatency0To1ms returns the AllLatency0To1ms field if non-nil, zero value otherwise.

### GetAllLatency0To1msOk

`func (o *OriginInspectorMeasurements) GetAllLatency0To1msOk() (*int32, bool)`

GetAllLatency0To1msOk returns a tuple with the AllLatency0To1ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllLatency0To1ms

`func (o *OriginInspectorMeasurements) SetAllLatency0To1ms(v int32)`

SetAllLatency0To1ms sets AllLatency0To1ms field to given value.

### HasAllLatency0To1ms

`func (o *OriginInspectorMeasurements) HasAllLatency0To1ms() bool`

HasAllLatency0To1ms returns a boolean if a field has been set.

### GetAllLatency1To5ms

`func (o *OriginInspectorMeasurements) GetAllLatency1To5ms() int32`

GetAllLatency1To5ms returns the AllLatency1To5ms field if non-nil, zero value otherwise.

### GetAllLatency1To5msOk

`func (o *OriginInspectorMeasurements) GetAllLatency1To5msOk() (*int32, bool)`

GetAllLatency1To5msOk returns a tuple with the AllLatency1To5ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllLatency1To5ms

`func (o *OriginInspectorMeasurements) SetAllLatency1To5ms(v int32)`

SetAllLatency1To5ms sets AllLatency1To5ms field to given value.

### HasAllLatency1To5ms

`func (o *OriginInspectorMeasurements) HasAllLatency1To5ms() bool`

HasAllLatency1To5ms returns a boolean if a field has been set.

### GetAllLatency5To10ms

`func (o *OriginInspectorMeasurements) GetAllLatency5To10ms() int32`

GetAllLatency5To10ms returns the AllLatency5To10ms field if non-nil, zero value otherwise.

### GetAllLatency5To10msOk

`func (o *OriginInspectorMeasurements) GetAllLatency5To10msOk() (*int32, bool)`

GetAllLatency5To10msOk returns a tuple with the AllLatency5To10ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllLatency5To10ms

`func (o *OriginInspectorMeasurements) SetAllLatency5To10ms(v int32)`

SetAllLatency5To10ms sets AllLatency5To10ms field to given value.

### HasAllLatency5To10ms

`func (o *OriginInspectorMeasurements) HasAllLatency5To10ms() bool`

HasAllLatency5To10ms returns a boolean if a field has been set.

### GetAllLatency10To50ms

`func (o *OriginInspectorMeasurements) GetAllLatency10To50ms() int32`

GetAllLatency10To50ms returns the AllLatency10To50ms field if non-nil, zero value otherwise.

### GetAllLatency10To50msOk

`func (o *OriginInspectorMeasurements) GetAllLatency10To50msOk() (*int32, bool)`

GetAllLatency10To50msOk returns a tuple with the AllLatency10To50ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllLatency10To50ms

`func (o *OriginInspectorMeasurements) SetAllLatency10To50ms(v int32)`

SetAllLatency10To50ms sets AllLatency10To50ms field to given value.

### HasAllLatency10To50ms

`func (o *OriginInspectorMeasurements) HasAllLatency10To50ms() bool`

HasAllLatency10To50ms returns a boolean if a field has been set.

### GetAllLatency50To100ms

`func (o *OriginInspectorMeasurements) GetAllLatency50To100ms() int32`

GetAllLatency50To100ms returns the AllLatency50To100ms field if non-nil, zero value otherwise.

### GetAllLatency50To100msOk

`func (o *OriginInspectorMeasurements) GetAllLatency50To100msOk() (*int32, bool)`

GetAllLatency50To100msOk returns a tuple with the AllLatency50To100ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllLatency50To100ms

`func (o *OriginInspectorMeasurements) SetAllLatency50To100ms(v int32)`

SetAllLatency50To100ms sets AllLatency50To100ms field to given value.

### HasAllLatency50To100ms

`func (o *OriginInspectorMeasurements) HasAllLatency50To100ms() bool`

HasAllLatency50To100ms returns a boolean if a field has been set.

### GetAllLatency100To250ms

`func (o *OriginInspectorMeasurements) GetAllLatency100To250ms() int32`

GetAllLatency100To250ms returns the AllLatency100To250ms field if non-nil, zero value otherwise.

### GetAllLatency100To250msOk

`func (o *OriginInspectorMeasurements) GetAllLatency100To250msOk() (*int32, bool)`

GetAllLatency100To250msOk returns a tuple with the AllLatency100To250ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllLatency100To250ms

`func (o *OriginInspectorMeasurements) SetAllLatency100To250ms(v int32)`

SetAllLatency100To250ms sets AllLatency100To250ms field to given value.

### HasAllLatency100To250ms

`func (o *OriginInspectorMeasurements) HasAllLatency100To250ms() bool`

HasAllLatency100To250ms returns a boolean if a field has been set.

### GetAllLatency250To500ms

`func (o *OriginInspectorMeasurements) GetAllLatency250To500ms() int32`

GetAllLatency250To500ms returns the AllLatency250To500ms field if non-nil, zero value otherwise.

### GetAllLatency250To500msOk

`func (o *OriginInspectorMeasurements) GetAllLatency250To500msOk() (*int32, bool)`

GetAllLatency250To500msOk returns a tuple with the AllLatency250To500ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllLatency250To500ms

`func (o *OriginInspectorMeasurements) SetAllLatency250To500ms(v int32)`

SetAllLatency250To500ms sets AllLatency250To500ms field to given value.

### HasAllLatency250To500ms

`func (o *OriginInspectorMeasurements) HasAllLatency250To500ms() bool`

HasAllLatency250To500ms returns a boolean if a field has been set.

### GetAllLatency500To1000ms

`func (o *OriginInspectorMeasurements) GetAllLatency500To1000ms() int32`

GetAllLatency500To1000ms returns the AllLatency500To1000ms field if non-nil, zero value otherwise.

### GetAllLatency500To1000msOk

`func (o *OriginInspectorMeasurements) GetAllLatency500To1000msOk() (*int32, bool)`

GetAllLatency500To1000msOk returns a tuple with the AllLatency500To1000ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllLatency500To1000ms

`func (o *OriginInspectorMeasurements) SetAllLatency500To1000ms(v int32)`

SetAllLatency500To1000ms sets AllLatency500To1000ms field to given value.

### HasAllLatency500To1000ms

`func (o *OriginInspectorMeasurements) HasAllLatency500To1000ms() bool`

HasAllLatency500To1000ms returns a boolean if a field has been set.

### GetAllLatency1000To5000ms

`func (o *OriginInspectorMeasurements) GetAllLatency1000To5000ms() int32`

GetAllLatency1000To5000ms returns the AllLatency1000To5000ms field if non-nil, zero value otherwise.

### GetAllLatency1000To5000msOk

`func (o *OriginInspectorMeasurements) GetAllLatency1000To5000msOk() (*int32, bool)`

GetAllLatency1000To5000msOk returns a tuple with the AllLatency1000To5000ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllLatency1000To5000ms

`func (o *OriginInspectorMeasurements) SetAllLatency1000To5000ms(v int32)`

SetAllLatency1000To5000ms sets AllLatency1000To5000ms field to given value.

### HasAllLatency1000To5000ms

`func (o *OriginInspectorMeasurements) HasAllLatency1000To5000ms() bool`

HasAllLatency1000To5000ms returns a boolean if a field has been set.

### GetAllLatency5000To10000ms

`func (o *OriginInspectorMeasurements) GetAllLatency5000To10000ms() int32`

GetAllLatency5000To10000ms returns the AllLatency5000To10000ms field if non-nil, zero value otherwise.

### GetAllLatency5000To10000msOk

`func (o *OriginInspectorMeasurements) GetAllLatency5000To10000msOk() (*int32, bool)`

GetAllLatency5000To10000msOk returns a tuple with the AllLatency5000To10000ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllLatency5000To10000ms

`func (o *OriginInspectorMeasurements) SetAllLatency5000To10000ms(v int32)`

SetAllLatency5000To10000ms sets AllLatency5000To10000ms field to given value.

### HasAllLatency5000To10000ms

`func (o *OriginInspectorMeasurements) HasAllLatency5000To10000ms() bool`

HasAllLatency5000To10000ms returns a boolean if a field has been set.

### GetAllLatency10000To60000ms

`func (o *OriginInspectorMeasurements) GetAllLatency10000To60000ms() int32`

GetAllLatency10000To60000ms returns the AllLatency10000To60000ms field if non-nil, zero value otherwise.

### GetAllLatency10000To60000msOk

`func (o *OriginInspectorMeasurements) GetAllLatency10000To60000msOk() (*int32, bool)`

GetAllLatency10000To60000msOk returns a tuple with the AllLatency10000To60000ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllLatency10000To60000ms

`func (o *OriginInspectorMeasurements) SetAllLatency10000To60000ms(v int32)`

SetAllLatency10000To60000ms sets AllLatency10000To60000ms field to given value.

### HasAllLatency10000To60000ms

`func (o *OriginInspectorMeasurements) HasAllLatency10000To60000ms() bool`

HasAllLatency10000To60000ms returns a boolean if a field has been set.

### GetAllLatency60000ms

`func (o *OriginInspectorMeasurements) GetAllLatency60000ms() int32`

GetAllLatency60000ms returns the AllLatency60000ms field if non-nil, zero value otherwise.

### GetAllLatency60000msOk

`func (o *OriginInspectorMeasurements) GetAllLatency60000msOk() (*int32, bool)`

GetAllLatency60000msOk returns a tuple with the AllLatency60000ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllLatency60000ms

`func (o *OriginInspectorMeasurements) SetAllLatency60000ms(v int32)`

SetAllLatency60000ms sets AllLatency60000ms field to given value.

### HasAllLatency60000ms

`func (o *OriginInspectorMeasurements) HasAllLatency60000ms() bool`

HasAllLatency60000ms returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
