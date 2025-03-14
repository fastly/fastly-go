# OriginInspectorValues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Responses** | Pointer to **int64** | Number of responses from origin. | [optional] 
**RespHeaderBytes** | Pointer to **int64** | Number of header bytes from origin. | [optional] 
**RespBodyBytes** | Pointer to **int64** | Number of body bytes from origin. | [optional] 
**Status1xx** | Pointer to **int64** | Number of 1xx \&quot;Informational\&quot; status codes delivered from origin. | [optional] 
**Status2xx** | Pointer to **int64** | Number of 2xx \&quot;Success\&quot; status codes delivered from origin. | [optional] 
**Status3xx** | Pointer to **int64** | Number of 3xx \&quot;Redirection\&quot; codes delivered from origin. | [optional] 
**Status4xx** | Pointer to **int64** | Number of 4xx \&quot;Client Error\&quot; codes delivered from origin. | [optional] 
**Status5xx** | Pointer to **int64** | Number of 5xx \&quot;Server Error\&quot; codes delivered from origin. | [optional] 
**Status200** | Pointer to **int64** | Number of responses received with status code 200 (Success) from origin. | [optional] 
**Status204** | Pointer to **int64** | Number of responses received with status code 204 (No Content) from origin. | [optional] 
**Status206** | Pointer to **int64** | Number of responses received with status code 206 (Partial Content) from origin. | [optional] 
**Status301** | Pointer to **int64** | Number of responses received with status code 301 (Moved Permanently) from origin. | [optional] 
**Status302** | Pointer to **int64** | Number of responses received with status code 302 (Found) from origin. | [optional] 
**Status304** | Pointer to **int64** | Number of responses received with status code 304 (Not Modified) from origin. | [optional] 
**Status400** | Pointer to **int64** | Number of responses received with status code 400 (Bad Request) from origin. | [optional] 
**Status401** | Pointer to **int64** | Number of responses received with status code 401 (Unauthorized) from origin. | [optional] 
**Status403** | Pointer to **int64** | Number of responses received with status code 403 (Forbidden) from origin. | [optional] 
**Status404** | Pointer to **int64** | Number of responses received with status code 404 (Not Found) from origin. | [optional] 
**Status416** | Pointer to **int64** | Number of responses received with status code 416 (Range Not Satisfiable) from origin. | [optional] 
**Status429** | Pointer to **int64** | Number of responses received with status code 429 (Too Many Requests) from origin. | [optional] 
**Status500** | Pointer to **int64** | Number of responses received with status code 500 (Internal Server Error) from origin. | [optional] 
**Status501** | Pointer to **int64** | Number of responses received with status code 501 (Not Implemented) from origin. | [optional] 
**Status502** | Pointer to **int64** | Number of responses received with status code 502 (Bad Gateway) from origin. | [optional] 
**Status503** | Pointer to **int64** | Number of responses received with status code 503 (Service Unavailable) from origin. | [optional] 
**Status504** | Pointer to **int64** | Number of responses received with status code 504 (Gateway Timeout) from origin. | [optional] 
**Status505** | Pointer to **int64** | Number of responses received with status code 505 (HTTP Version Not Supported) from origin. | [optional] 
**Status530** | Pointer to **int64** | Number of responses received from origin with status code 530. | [optional] 
**Latency0To1ms** | Pointer to **int64** | Number of responses from origin with latency between 0 and 1 millisecond. | [optional] 
**Latency1To5ms** | Pointer to **int64** | Number of responses from origin with latency between 1 and 5 milliseconds. | [optional] 
**Latency5To10ms** | Pointer to **int64** | Number of responses from origin with latency between 5 and 10 milliseconds. | [optional] 
**Latency10To50ms** | Pointer to **int64** | Number of responses from origin with latency between 10 and 50 milliseconds. | [optional] 
**Latency50To100ms** | Pointer to **int64** | Number of responses from origin with latency between 50 and 100 milliseconds. | [optional] 
**Latency100To250ms** | Pointer to **int64** | Number of responses from origin with latency between 100 and 250 milliseconds. | [optional] 
**Latency250To500ms** | Pointer to **int64** | Number of responses from origin with latency between 250 and 500 milliseconds. | [optional] 
**Latency500To1000ms** | Pointer to **int64** | Number of responses from origin with latency between 500 and 1,000 milliseconds. | [optional] 
**Latency1000To5000ms** | Pointer to **int64** | Number of responses from origin with latency between 1,000 and 5,000 milliseconds. | [optional] 
**Latency5000To10000ms** | Pointer to **int64** | Number of responses from origin with latency between 5,000 and 10,000 milliseconds. | [optional] 
**Latency10000To60000ms** | Pointer to **int64** | Number of responses from origin with latency between 10,000 and 60,000 milliseconds. | [optional] 
**Latency60000ms** | Pointer to **int64** | Number of responses from origin with latency of 60,000 milliseconds and above. | [optional] 
**WafResponses** | Pointer to **int64** | Number of responses received for origin requests made by the Fastly WAF. | [optional] 
**WafRespHeaderBytes** | Pointer to **int64** | Number of header bytes received for origin requests made by the Fastly WAF. | [optional] 
**WafRespBodyBytes** | Pointer to **int64** | Number of body bytes received for origin requests made by the Fastly WAF. | [optional] 
**WafStatus1xx** | Pointer to **int64** | Number of 1xx \&quot;Informational\&quot; status codes received for origin requests made by the Fastly WAF. | [optional] 
**WafStatus2xx** | Pointer to **int64** | Number of 2xx \&quot;Success\&quot; status codes received for origin requests made by the Fastly WAF. | [optional] 
**WafStatus3xx** | Pointer to **int64** | Number of 3xx \&quot;Redirection\&quot; codes received for origin requests made by the Fastly WAF. | [optional] 
**WafStatus4xx** | Pointer to **int64** | Number of 4xx \&quot;Client Error\&quot; codes received for origin requests made by the Fastly WAF. | [optional] 
**WafStatus5xx** | Pointer to **int64** | Number of 5xx \&quot;Server Error\&quot; codes received for origin requests made by the Fastly WAF. | [optional] 
**WafStatus200** | Pointer to **int64** | Number of responses received with status code 200 (Success) received for origin requests made by the Fastly WAF. | [optional] 
**WafStatus204** | Pointer to **int64** | Number of responses received with status code 204 (No Content) received for origin requests made by the Fastly WAF. | [optional] 
**WafStatus206** | Pointer to **int64** | Number of responses received with status code 206 (Partial Content) received for origin requests made by the Fastly WAF. | [optional] 
**WafStatus301** | Pointer to **int64** | Number of responses received with status code 301 (Moved Permanently) received for origin requests made by the Fastly WAF. | [optional] 
**WafStatus302** | Pointer to **int64** | Number of responses received with status code 302 (Found) received for origin requests made by the Fastly WAF. | [optional] 
**WafStatus304** | Pointer to **int64** | Number of responses received with status code 304 (Not Modified) received for origin requests made by the Fastly WAF. | [optional] 
**WafStatus400** | Pointer to **int64** | Number of responses received with status code 400 (Bad Request) received for origin requests made by the Fastly WAF. | [optional] 
**WafStatus401** | Pointer to **int64** | Number of responses received with status code 401 (Unauthorized) received for origin requests made by the Fastly WAF. | [optional] 
**WafStatus403** | Pointer to **int64** | Number of responses received with status code 403 (Forbidden) received for origin requests made by the Fastly WAF. | [optional] 
**WafStatus404** | Pointer to **int64** | Number of responses received with status code 404 (Not Found) received for origin requests made by the Fastly WAF. | [optional] 
**WafStatus416** | Pointer to **int64** | Number of responses received with status code 416 (Range Not Satisfiable) received for origin requests made by the Fastly WAF. | [optional] 
**WafStatus429** | Pointer to **int64** | Number of responses received with status code 429 (Too Many Requests) received for origin requests made by the Fastly WAF. | [optional] 
**WafStatus500** | Pointer to **int64** | Number of responses received with status code 500 (Internal Server Error) received for origin requests made by the Fastly WAF. | [optional] 
**WafStatus501** | Pointer to **int64** | Number of responses received with status code 501 (Not Implemented) received for origin requests made by the Fastly WAF. | [optional] 
**WafStatus502** | Pointer to **int64** | Number of responses received with status code 502 (Bad Gateway) received for origin requests made by the Fastly WAF. | [optional] 
**WafStatus503** | Pointer to **int64** | Number of responses received with status code 503 (Service Unavailable) received for origin requests made by the Fastly WAF. | [optional] 
**WafStatus504** | Pointer to **int64** | Number of responses received with status code 504 (Gateway Timeout) received for origin requests made by the Fastly WAF. | [optional] 
**WafStatus505** | Pointer to **int64** | Number of responses received with status code 505 (HTTP Version Not Supported) received for origin requests made by the Fastly WAF. | [optional] 
**WafStatus530** | Pointer to **int64** | Number of responses received with status code 530 received for origin requests made by the Fastly WAF. | [optional] 
**WafLatency0To1ms** | Pointer to **int64** | Number of responses with latency between 0 and 1 millisecond received for origin requests made by the Fastly WAF. | [optional] 
**WafLatency1To5ms** | Pointer to **int64** | Number of responses with latency between 1 and 5 milliseconds received for origin requests made by the Fastly WAF. | [optional] 
**WafLatency5To10ms** | Pointer to **int64** | Number of responses with latency between 5 and 10 milliseconds received for origin requests made by the Fastly WAF. | [optional] 
**WafLatency10To50ms** | Pointer to **int64** | Number of responses with latency between 10 and 50 milliseconds received for origin requests made by the Fastly WAF. | [optional] 
**WafLatency50To100ms** | Pointer to **int64** | Number of responses with latency between 50 and 100 milliseconds received for origin requests made by the Fastly WAF. | [optional] 
**WafLatency100To250ms** | Pointer to **int64** | Number of responses with latency between 100 and 250 milliseconds received for origin requests made by the Fastly WAF. | [optional] 
**WafLatency250To500ms** | Pointer to **int64** | Number of responses with latency between 250 and 500 milliseconds received for origin requests made by the Fastly WAF. | [optional] 
**WafLatency500To1000ms** | Pointer to **int64** | Number of responses with latency between 500 and 1,000 milliseconds received for origin requests made by the Fastly WAF. | [optional] 
**WafLatency1000To5000ms** | Pointer to **int64** | Number of responses with latency between 1,000 and 5,000 milliseconds received for origin requests made by the Fastly WAF. | [optional] 
**WafLatency5000To10000ms** | Pointer to **int64** | Number of responses with latency between 5,000 and 10,000 milliseconds received for origin requests made by the Fastly WAF. | [optional] 
**WafLatency10000To60000ms** | Pointer to **int64** | Number of responses with latency between 10,000 and 60,000 milliseconds received for origin requests made by the Fastly WAF. | [optional] 
**WafLatency60000ms** | Pointer to **int64** | Number of responses with latency of 60,000 milliseconds and above received for origin requests made by the Fastly WAF. | [optional] 
**ComputeResponses** | Pointer to **int64** | Number of responses for origin received by the Compute platform. | [optional] 
**ComputeRespHeaderBytes** | Pointer to **int64** | Number of header bytes for origin received by the Compute platform. | [optional] 
**ComputeRespBodyBytes** | Pointer to **int64** | Number of body bytes for origin received by the Compute platform. | [optional] 
**ComputeStatus1xx** | Pointer to **int64** | Number of 1xx \&quot;Informational\&quot; status codes for origin received by the Compute platform. | [optional] 
**ComputeStatus2xx** | Pointer to **int64** | Number of 2xx \&quot;Success\&quot; status codes for origin received by the Compute platform. | [optional] 
**ComputeStatus3xx** | Pointer to **int64** | Number of 3xx \&quot;Redirection\&quot; codes for origin received by the Compute platform. | [optional] 
**ComputeStatus4xx** | Pointer to **int64** | Number of 4xx \&quot;Client Error\&quot; codes for origin received by the Compute platform. | [optional] 
**ComputeStatus5xx** | Pointer to **int64** | Number of 5xx \&quot;Server Error\&quot; codes for origin received by the Compute platform. | [optional] 
**ComputeStatus200** | Pointer to **int64** | Number of responses received with status code 200 (Success) for origin received by the Compute platform. | [optional] 
**ComputeStatus204** | Pointer to **int64** | Number of responses received with status code 204 (No Content) for origin received by the Compute platform. | [optional] 
**ComputeStatus206** | Pointer to **int64** | Number of responses received with status code 206 (Partial Content) for origin received by the Compute platform. | [optional] 
**ComputeStatus301** | Pointer to **int64** | Number of responses received with status code 301 (Moved Permanently) for origin received by the Compute platform. | [optional] 
**ComputeStatus302** | Pointer to **int64** | Number of responses received with status code 302 (Found) for origin received by the Compute platform. | [optional] 
**ComputeStatus304** | Pointer to **int64** | Number of responses received with status code 304 (Not Modified) for origin received by the Compute platform. | [optional] 
**ComputeStatus400** | Pointer to **int64** | Number of responses received with status code 400 (Bad Request) for origin received by the Compute platform. | [optional] 
**ComputeStatus401** | Pointer to **int64** | Number of responses received with status code 401 (Unauthorized) for origin received by the Compute platform. | [optional] 
**ComputeStatus403** | Pointer to **int64** | Number of responses received with status code 403 (Forbidden) for origin received by the Compute platform. | [optional] 
**ComputeStatus404** | Pointer to **int64** | Number of responses received with status code 404 (Not Found) for origin received by the Compute platform. | [optional] 
**ComputeStatus416** | Pointer to **int64** | Number of responses received with status code 416 (Range Not Satisfiable) for origin received by the Compute platform. | [optional] 
**ComputeStatus429** | Pointer to **int64** | Number of responses received with status code 429 (Too Many Requests) for origin received by the Compute platform. | [optional] 
**ComputeStatus500** | Pointer to **int64** | Number of responses received with status code 500 (Internal Server Error) for origin received by the Compute platform. | [optional] 
**ComputeStatus501** | Pointer to **int64** | Number of responses received with status code 501 (Not Implemented) for origin received by the Compute platform. | [optional] 
**ComputeStatus502** | Pointer to **int64** | Number of responses received with status code 502 (Bad Gateway) for origin received by the Compute platform. | [optional] 
**ComputeStatus503** | Pointer to **int64** | Number of responses received with status code 503 (Service Unavailable) for origin received by the Compute platform. | [optional] 
**ComputeStatus504** | Pointer to **int64** | Number of responses received with status code 504 (Gateway Timeout) for origin received by the Compute platform. | [optional] 
**ComputeStatus505** | Pointer to **int64** | Number of responses received with status code 505 (HTTP Version Not Supported) for origin received by the Compute platform. | [optional] 
**ComputeStatus530** | Pointer to **int64** | Number of responses received with status code 530 for origin received by the Compute platform. | [optional] 
**ComputeLatency0To1ms** | Pointer to **int64** | Number of responses with latency between 0 and 1 millisecond for origin received by the Compute platform. | [optional] 
**ComputeLatency1To5ms** | Pointer to **int64** | Number of responses with latency between 1 and 5 milliseconds for origin received by the Compute platform. | [optional] 
**ComputeLatency5To10ms** | Pointer to **int64** | Number of responses with latency between 5 and 10 milliseconds for origin received by the Compute platform. | [optional] 
**ComputeLatency10To50ms** | Pointer to **int64** | Number of responses with latency between 10 and 50 milliseconds for origin received by the Compute platform. | [optional] 
**ComputeLatency50To100ms** | Pointer to **int64** | Number of responses with latency between 50 and 100 milliseconds for origin received by the Compute platform. | [optional] 
**ComputeLatency100To250ms** | Pointer to **int64** | Number of responses with latency between 100 and 250 milliseconds for origin received by the Compute platform. | [optional] 
**ComputeLatency250To500ms** | Pointer to **int64** | Number of responses with latency between 250 and 500 milliseconds for origin received by the Compute platform. | [optional] 
**ComputeLatency500To1000ms** | Pointer to **int64** | Number of responses with latency between 500 and 1,000 milliseconds for origin received by the Compute platform. | [optional] 
**ComputeLatency1000To5000ms** | Pointer to **int64** | Number of responses with latency between 1,000 and 5,000 milliseconds for origin received by the Compute platform. | [optional] 
**ComputeLatency5000To10000ms** | Pointer to **int64** | Number of responses with latency between 5,000 and 10,000 milliseconds for origin received by the Compute platform. | [optional] 
**ComputeLatency10000To60000ms** | Pointer to **int64** | Number of responses with latency between 10,000 and 60,000 milliseconds for origin received by the Compute platform. | [optional] 
**ComputeLatency60000ms** | Pointer to **int64** | Number of responses with latency of 60,000 milliseconds and above for origin received by the Compute platform. | [optional] 
**AllResponses** | Pointer to **int64** | Number of responses received for origin requests made by all sources. | [optional] 
**AllRespHeaderBytes** | Pointer to **int64** | Number of header bytes received for origin requests made by all sources. | [optional] 
**AllRespBodyBytes** | Pointer to **int64** | Number of body bytes received for origin requests made by all sources. | [optional] 
**AllStatus1xx** | Pointer to **int64** | Number of 1xx \&quot;Informational\&quot; category status codes delivered received for origin requests made by all sources. | [optional] 
**AllStatus2xx** | Pointer to **int64** | Number of 2xx \&quot;Success\&quot; status codes received for origin requests made by all sources. | [optional] 
**AllStatus3xx** | Pointer to **int64** | Number of 3xx \&quot;Redirection\&quot; codes received for origin requests made by all sources. | [optional] 
**AllStatus4xx** | Pointer to **int64** | Number of 4xx \&quot;Client Error\&quot; codes received for origin requests made by all sources. | [optional] 
**AllStatus5xx** | Pointer to **int64** | Number of 5xx \&quot;Server Error\&quot; codes received for origin requests made by all sources. | [optional] 
**AllStatus200** | Pointer to **int64** | Number of responses received with status code 200 (Success) received for origin requests made by all sources. | [optional] 
**AllStatus204** | Pointer to **int64** | Number of responses received with status code 204 (No Content) received for origin requests made by all sources. | [optional] 
**AllStatus206** | Pointer to **int64** | Number of responses received with status code 206 (Partial Content) received for origin requests made by all sources. | [optional] 
**AllStatus301** | Pointer to **int64** | Number of responses received with status code 301 (Moved Permanently) received for origin requests made by all sources. | [optional] 
**AllStatus302** | Pointer to **int64** | Number of responses received with status code 302 (Found) received for origin requests made by all sources. | [optional] 
**AllStatus304** | Pointer to **int64** | Number of responses received with status code 304 (Not Modified) received for origin requests made by all sources. | [optional] 
**AllStatus400** | Pointer to **int64** | Number of responses received with status code 400 (Bad Request) received for origin requests made by all sources. | [optional] 
**AllStatus401** | Pointer to **int64** | Number of responses received with status code 401 (Unauthorized) received for origin requests made by all sources. | [optional] 
**AllStatus403** | Pointer to **int64** | Number of responses received with status code 403 (Forbidden) received for origin requests made by all sources. | [optional] 
**AllStatus404** | Pointer to **int64** | Number of responses received with status code 404 (Not Found) received for origin requests made by all sources. | [optional] 
**AllStatus416** | Pointer to **int64** | Number of responses received with status code 416 (Range Not Satisfiable) received for origin requests made by all sources. | [optional] 
**AllStatus429** | Pointer to **int64** | Number of responses received with status code 429 (Too Many Requests) received for origin requests made by all sources. | [optional] 
**AllStatus500** | Pointer to **int64** | Number of responses received with status code 500 (Internal Server Error) received for origin requests made by all sources. | [optional] 
**AllStatus501** | Pointer to **int64** | Number of responses received with status code 501 (Not Implemented) received for origin requests made by all sources. | [optional] 
**AllStatus502** | Pointer to **int64** | Number of responses received with status code 502 (Bad Gateway) received for origin requests made by all sources. | [optional] 
**AllStatus503** | Pointer to **int64** | Number of responses received with status code 503 (Service Unavailable) received for origin requests made by all sources. | [optional] 
**AllStatus504** | Pointer to **int64** | Number of responses received with status code 504 (Gateway Timeout) received for origin requests made by all sources. | [optional] 
**AllStatus505** | Pointer to **int64** | Number of responses received with status code 505 (HTTP Version Not Supported) received for origin requests made by all sources. | [optional] 
**AllStatus530** | Pointer to **int64** | Number of responses received with status code 530 received for origin requests made by all sources. | [optional] 
**AllLatency0To1ms** | Pointer to **int64** | Number of responses with latency between 0 and 1 millisecond received for origin requests made by all sources. | [optional] 
**AllLatency1To5ms** | Pointer to **int64** | Number of responses with latency between 1 and 5 milliseconds received for origin requests made by all sources. | [optional] 
**AllLatency5To10ms** | Pointer to **int64** | Number of responses with latency between 5 and 10 milliseconds received for origin requests made by all sources. | [optional] 
**AllLatency10To50ms** | Pointer to **int64** | Number of responses with latency between 10 and 50 milliseconds received for origin requests made by all sources. | [optional] 
**AllLatency50To100ms** | Pointer to **int64** | Number of responses with latency between 50 and 100 milliseconds received for origin requests made by all sources. | [optional] 
**AllLatency100To250ms** | Pointer to **int64** | Number of responses with latency between 100 and 250 milliseconds received for origin requests made by all sources. | [optional] 
**AllLatency250To500ms** | Pointer to **int64** | Number of responses with latency between 250 and 500 milliseconds received for origin requests made by all sources. | [optional] 
**AllLatency500To1000ms** | Pointer to **int64** | Number of responses with latency between 500 and 1,000 milliseconds received for origin requests made by all sources. | [optional] 
**AllLatency1000To5000ms** | Pointer to **int64** | Number of responses with latency between 1,000 and 5,000 milliseconds received for origin requests made by all sources. | [optional] 
**AllLatency5000To10000ms** | Pointer to **int64** | Number of responses with latency between 5,000 and 10,000 milliseconds received for origin requests made by all sources. | [optional] 
**AllLatency10000To60000ms** | Pointer to **int64** | Number of responses with latency between 10,000 and 60,000 milliseconds received for origin requests made by all sources. | [optional] 
**AllLatency60000ms** | Pointer to **int64** | Number of responses with latency of 60,000 milliseconds and above received for origin requests made by all sources. | [optional] 

## Methods

### NewOriginInspectorValues

`func NewOriginInspectorValues() *OriginInspectorValues`

NewOriginInspectorValues instantiates a new OriginInspectorValues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOriginInspectorValuesWithDefaults

`func NewOriginInspectorValuesWithDefaults() *OriginInspectorValues`

NewOriginInspectorValuesWithDefaults instantiates a new OriginInspectorValues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponses

`func (o *OriginInspectorValues) GetResponses() int64`

GetResponses returns the Responses field if non-nil, zero value otherwise.

### GetResponsesOk

`func (o *OriginInspectorValues) GetResponsesOk() (*int64, bool)`

GetResponsesOk returns a tuple with the Responses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponses

`func (o *OriginInspectorValues) SetResponses(v int64)`

SetResponses sets Responses field to given value.

### HasResponses

`func (o *OriginInspectorValues) HasResponses() bool`

HasResponses returns a boolean if a field has been set.

### GetRespHeaderBytes

`func (o *OriginInspectorValues) GetRespHeaderBytes() int64`

GetRespHeaderBytes returns the RespHeaderBytes field if non-nil, zero value otherwise.

### GetRespHeaderBytesOk

`func (o *OriginInspectorValues) GetRespHeaderBytesOk() (*int64, bool)`

GetRespHeaderBytesOk returns a tuple with the RespHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRespHeaderBytes

`func (o *OriginInspectorValues) SetRespHeaderBytes(v int64)`

SetRespHeaderBytes sets RespHeaderBytes field to given value.

### HasRespHeaderBytes

`func (o *OriginInspectorValues) HasRespHeaderBytes() bool`

HasRespHeaderBytes returns a boolean if a field has been set.

### GetRespBodyBytes

`func (o *OriginInspectorValues) GetRespBodyBytes() int64`

GetRespBodyBytes returns the RespBodyBytes field if non-nil, zero value otherwise.

### GetRespBodyBytesOk

`func (o *OriginInspectorValues) GetRespBodyBytesOk() (*int64, bool)`

GetRespBodyBytesOk returns a tuple with the RespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRespBodyBytes

`func (o *OriginInspectorValues) SetRespBodyBytes(v int64)`

SetRespBodyBytes sets RespBodyBytes field to given value.

### HasRespBodyBytes

`func (o *OriginInspectorValues) HasRespBodyBytes() bool`

HasRespBodyBytes returns a boolean if a field has been set.

### GetStatus1xx

`func (o *OriginInspectorValues) GetStatus1xx() int64`

GetStatus1xx returns the Status1xx field if non-nil, zero value otherwise.

### GetStatus1xxOk

`func (o *OriginInspectorValues) GetStatus1xxOk() (*int64, bool)`

GetStatus1xxOk returns a tuple with the Status1xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus1xx

`func (o *OriginInspectorValues) SetStatus1xx(v int64)`

SetStatus1xx sets Status1xx field to given value.

### HasStatus1xx

`func (o *OriginInspectorValues) HasStatus1xx() bool`

HasStatus1xx returns a boolean if a field has been set.

### GetStatus2xx

`func (o *OriginInspectorValues) GetStatus2xx() int64`

GetStatus2xx returns the Status2xx field if non-nil, zero value otherwise.

### GetStatus2xxOk

`func (o *OriginInspectorValues) GetStatus2xxOk() (*int64, bool)`

GetStatus2xxOk returns a tuple with the Status2xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus2xx

`func (o *OriginInspectorValues) SetStatus2xx(v int64)`

SetStatus2xx sets Status2xx field to given value.

### HasStatus2xx

`func (o *OriginInspectorValues) HasStatus2xx() bool`

HasStatus2xx returns a boolean if a field has been set.

### GetStatus3xx

`func (o *OriginInspectorValues) GetStatus3xx() int64`

GetStatus3xx returns the Status3xx field if non-nil, zero value otherwise.

### GetStatus3xxOk

`func (o *OriginInspectorValues) GetStatus3xxOk() (*int64, bool)`

GetStatus3xxOk returns a tuple with the Status3xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus3xx

`func (o *OriginInspectorValues) SetStatus3xx(v int64)`

SetStatus3xx sets Status3xx field to given value.

### HasStatus3xx

`func (o *OriginInspectorValues) HasStatus3xx() bool`

HasStatus3xx returns a boolean if a field has been set.

### GetStatus4xx

`func (o *OriginInspectorValues) GetStatus4xx() int64`

GetStatus4xx returns the Status4xx field if non-nil, zero value otherwise.

### GetStatus4xxOk

`func (o *OriginInspectorValues) GetStatus4xxOk() (*int64, bool)`

GetStatus4xxOk returns a tuple with the Status4xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus4xx

`func (o *OriginInspectorValues) SetStatus4xx(v int64)`

SetStatus4xx sets Status4xx field to given value.

### HasStatus4xx

`func (o *OriginInspectorValues) HasStatus4xx() bool`

HasStatus4xx returns a boolean if a field has been set.

### GetStatus5xx

`func (o *OriginInspectorValues) GetStatus5xx() int64`

GetStatus5xx returns the Status5xx field if non-nil, zero value otherwise.

### GetStatus5xxOk

`func (o *OriginInspectorValues) GetStatus5xxOk() (*int64, bool)`

GetStatus5xxOk returns a tuple with the Status5xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus5xx

`func (o *OriginInspectorValues) SetStatus5xx(v int64)`

SetStatus5xx sets Status5xx field to given value.

### HasStatus5xx

`func (o *OriginInspectorValues) HasStatus5xx() bool`

HasStatus5xx returns a boolean if a field has been set.

### GetStatus200

`func (o *OriginInspectorValues) GetStatus200() int64`

GetStatus200 returns the Status200 field if non-nil, zero value otherwise.

### GetStatus200Ok

`func (o *OriginInspectorValues) GetStatus200Ok() (*int64, bool)`

GetStatus200Ok returns a tuple with the Status200 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus200

`func (o *OriginInspectorValues) SetStatus200(v int64)`

SetStatus200 sets Status200 field to given value.

### HasStatus200

`func (o *OriginInspectorValues) HasStatus200() bool`

HasStatus200 returns a boolean if a field has been set.

### GetStatus204

`func (o *OriginInspectorValues) GetStatus204() int64`

GetStatus204 returns the Status204 field if non-nil, zero value otherwise.

### GetStatus204Ok

`func (o *OriginInspectorValues) GetStatus204Ok() (*int64, bool)`

GetStatus204Ok returns a tuple with the Status204 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus204

`func (o *OriginInspectorValues) SetStatus204(v int64)`

SetStatus204 sets Status204 field to given value.

### HasStatus204

`func (o *OriginInspectorValues) HasStatus204() bool`

HasStatus204 returns a boolean if a field has been set.

### GetStatus206

`func (o *OriginInspectorValues) GetStatus206() int64`

GetStatus206 returns the Status206 field if non-nil, zero value otherwise.

### GetStatus206Ok

`func (o *OriginInspectorValues) GetStatus206Ok() (*int64, bool)`

GetStatus206Ok returns a tuple with the Status206 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus206

`func (o *OriginInspectorValues) SetStatus206(v int64)`

SetStatus206 sets Status206 field to given value.

### HasStatus206

`func (o *OriginInspectorValues) HasStatus206() bool`

HasStatus206 returns a boolean if a field has been set.

### GetStatus301

`func (o *OriginInspectorValues) GetStatus301() int64`

GetStatus301 returns the Status301 field if non-nil, zero value otherwise.

### GetStatus301Ok

`func (o *OriginInspectorValues) GetStatus301Ok() (*int64, bool)`

GetStatus301Ok returns a tuple with the Status301 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus301

`func (o *OriginInspectorValues) SetStatus301(v int64)`

SetStatus301 sets Status301 field to given value.

### HasStatus301

`func (o *OriginInspectorValues) HasStatus301() bool`

HasStatus301 returns a boolean if a field has been set.

### GetStatus302

`func (o *OriginInspectorValues) GetStatus302() int64`

GetStatus302 returns the Status302 field if non-nil, zero value otherwise.

### GetStatus302Ok

`func (o *OriginInspectorValues) GetStatus302Ok() (*int64, bool)`

GetStatus302Ok returns a tuple with the Status302 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus302

`func (o *OriginInspectorValues) SetStatus302(v int64)`

SetStatus302 sets Status302 field to given value.

### HasStatus302

`func (o *OriginInspectorValues) HasStatus302() bool`

HasStatus302 returns a boolean if a field has been set.

### GetStatus304

`func (o *OriginInspectorValues) GetStatus304() int64`

GetStatus304 returns the Status304 field if non-nil, zero value otherwise.

### GetStatus304Ok

`func (o *OriginInspectorValues) GetStatus304Ok() (*int64, bool)`

GetStatus304Ok returns a tuple with the Status304 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus304

`func (o *OriginInspectorValues) SetStatus304(v int64)`

SetStatus304 sets Status304 field to given value.

### HasStatus304

`func (o *OriginInspectorValues) HasStatus304() bool`

HasStatus304 returns a boolean if a field has been set.

### GetStatus400

`func (o *OriginInspectorValues) GetStatus400() int64`

GetStatus400 returns the Status400 field if non-nil, zero value otherwise.

### GetStatus400Ok

`func (o *OriginInspectorValues) GetStatus400Ok() (*int64, bool)`

GetStatus400Ok returns a tuple with the Status400 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus400

`func (o *OriginInspectorValues) SetStatus400(v int64)`

SetStatus400 sets Status400 field to given value.

### HasStatus400

`func (o *OriginInspectorValues) HasStatus400() bool`

HasStatus400 returns a boolean if a field has been set.

### GetStatus401

`func (o *OriginInspectorValues) GetStatus401() int64`

GetStatus401 returns the Status401 field if non-nil, zero value otherwise.

### GetStatus401Ok

`func (o *OriginInspectorValues) GetStatus401Ok() (*int64, bool)`

GetStatus401Ok returns a tuple with the Status401 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus401

`func (o *OriginInspectorValues) SetStatus401(v int64)`

SetStatus401 sets Status401 field to given value.

### HasStatus401

`func (o *OriginInspectorValues) HasStatus401() bool`

HasStatus401 returns a boolean if a field has been set.

### GetStatus403

`func (o *OriginInspectorValues) GetStatus403() int64`

GetStatus403 returns the Status403 field if non-nil, zero value otherwise.

### GetStatus403Ok

`func (o *OriginInspectorValues) GetStatus403Ok() (*int64, bool)`

GetStatus403Ok returns a tuple with the Status403 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus403

`func (o *OriginInspectorValues) SetStatus403(v int64)`

SetStatus403 sets Status403 field to given value.

### HasStatus403

`func (o *OriginInspectorValues) HasStatus403() bool`

HasStatus403 returns a boolean if a field has been set.

### GetStatus404

`func (o *OriginInspectorValues) GetStatus404() int64`

GetStatus404 returns the Status404 field if non-nil, zero value otherwise.

### GetStatus404Ok

`func (o *OriginInspectorValues) GetStatus404Ok() (*int64, bool)`

GetStatus404Ok returns a tuple with the Status404 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus404

`func (o *OriginInspectorValues) SetStatus404(v int64)`

SetStatus404 sets Status404 field to given value.

### HasStatus404

`func (o *OriginInspectorValues) HasStatus404() bool`

HasStatus404 returns a boolean if a field has been set.

### GetStatus416

`func (o *OriginInspectorValues) GetStatus416() int64`

GetStatus416 returns the Status416 field if non-nil, zero value otherwise.

### GetStatus416Ok

`func (o *OriginInspectorValues) GetStatus416Ok() (*int64, bool)`

GetStatus416Ok returns a tuple with the Status416 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus416

`func (o *OriginInspectorValues) SetStatus416(v int64)`

SetStatus416 sets Status416 field to given value.

### HasStatus416

`func (o *OriginInspectorValues) HasStatus416() bool`

HasStatus416 returns a boolean if a field has been set.

### GetStatus429

`func (o *OriginInspectorValues) GetStatus429() int64`

GetStatus429 returns the Status429 field if non-nil, zero value otherwise.

### GetStatus429Ok

`func (o *OriginInspectorValues) GetStatus429Ok() (*int64, bool)`

GetStatus429Ok returns a tuple with the Status429 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus429

`func (o *OriginInspectorValues) SetStatus429(v int64)`

SetStatus429 sets Status429 field to given value.

### HasStatus429

`func (o *OriginInspectorValues) HasStatus429() bool`

HasStatus429 returns a boolean if a field has been set.

### GetStatus500

`func (o *OriginInspectorValues) GetStatus500() int64`

GetStatus500 returns the Status500 field if non-nil, zero value otherwise.

### GetStatus500Ok

`func (o *OriginInspectorValues) GetStatus500Ok() (*int64, bool)`

GetStatus500Ok returns a tuple with the Status500 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus500

`func (o *OriginInspectorValues) SetStatus500(v int64)`

SetStatus500 sets Status500 field to given value.

### HasStatus500

`func (o *OriginInspectorValues) HasStatus500() bool`

HasStatus500 returns a boolean if a field has been set.

### GetStatus501

`func (o *OriginInspectorValues) GetStatus501() int64`

GetStatus501 returns the Status501 field if non-nil, zero value otherwise.

### GetStatus501Ok

`func (o *OriginInspectorValues) GetStatus501Ok() (*int64, bool)`

GetStatus501Ok returns a tuple with the Status501 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus501

`func (o *OriginInspectorValues) SetStatus501(v int64)`

SetStatus501 sets Status501 field to given value.

### HasStatus501

`func (o *OriginInspectorValues) HasStatus501() bool`

HasStatus501 returns a boolean if a field has been set.

### GetStatus502

`func (o *OriginInspectorValues) GetStatus502() int64`

GetStatus502 returns the Status502 field if non-nil, zero value otherwise.

### GetStatus502Ok

`func (o *OriginInspectorValues) GetStatus502Ok() (*int64, bool)`

GetStatus502Ok returns a tuple with the Status502 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus502

`func (o *OriginInspectorValues) SetStatus502(v int64)`

SetStatus502 sets Status502 field to given value.

### HasStatus502

`func (o *OriginInspectorValues) HasStatus502() bool`

HasStatus502 returns a boolean if a field has been set.

### GetStatus503

`func (o *OriginInspectorValues) GetStatus503() int64`

GetStatus503 returns the Status503 field if non-nil, zero value otherwise.

### GetStatus503Ok

`func (o *OriginInspectorValues) GetStatus503Ok() (*int64, bool)`

GetStatus503Ok returns a tuple with the Status503 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus503

`func (o *OriginInspectorValues) SetStatus503(v int64)`

SetStatus503 sets Status503 field to given value.

### HasStatus503

`func (o *OriginInspectorValues) HasStatus503() bool`

HasStatus503 returns a boolean if a field has been set.

### GetStatus504

`func (o *OriginInspectorValues) GetStatus504() int64`

GetStatus504 returns the Status504 field if non-nil, zero value otherwise.

### GetStatus504Ok

`func (o *OriginInspectorValues) GetStatus504Ok() (*int64, bool)`

GetStatus504Ok returns a tuple with the Status504 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus504

`func (o *OriginInspectorValues) SetStatus504(v int64)`

SetStatus504 sets Status504 field to given value.

### HasStatus504

`func (o *OriginInspectorValues) HasStatus504() bool`

HasStatus504 returns a boolean if a field has been set.

### GetStatus505

`func (o *OriginInspectorValues) GetStatus505() int64`

GetStatus505 returns the Status505 field if non-nil, zero value otherwise.

### GetStatus505Ok

`func (o *OriginInspectorValues) GetStatus505Ok() (*int64, bool)`

GetStatus505Ok returns a tuple with the Status505 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus505

`func (o *OriginInspectorValues) SetStatus505(v int64)`

SetStatus505 sets Status505 field to given value.

### HasStatus505

`func (o *OriginInspectorValues) HasStatus505() bool`

HasStatus505 returns a boolean if a field has been set.

### GetStatus530

`func (o *OriginInspectorValues) GetStatus530() int64`

GetStatus530 returns the Status530 field if non-nil, zero value otherwise.

### GetStatus530Ok

`func (o *OriginInspectorValues) GetStatus530Ok() (*int64, bool)`

GetStatus530Ok returns a tuple with the Status530 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus530

`func (o *OriginInspectorValues) SetStatus530(v int64)`

SetStatus530 sets Status530 field to given value.

### HasStatus530

`func (o *OriginInspectorValues) HasStatus530() bool`

HasStatus530 returns a boolean if a field has been set.

### GetLatency0To1ms

`func (o *OriginInspectorValues) GetLatency0To1ms() int64`

GetLatency0To1ms returns the Latency0To1ms field if non-nil, zero value otherwise.

### GetLatency0To1msOk

`func (o *OriginInspectorValues) GetLatency0To1msOk() (*int64, bool)`

GetLatency0To1msOk returns a tuple with the Latency0To1ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatency0To1ms

`func (o *OriginInspectorValues) SetLatency0To1ms(v int64)`

SetLatency0To1ms sets Latency0To1ms field to given value.

### HasLatency0To1ms

`func (o *OriginInspectorValues) HasLatency0To1ms() bool`

HasLatency0To1ms returns a boolean if a field has been set.

### GetLatency1To5ms

`func (o *OriginInspectorValues) GetLatency1To5ms() int64`

GetLatency1To5ms returns the Latency1To5ms field if non-nil, zero value otherwise.

### GetLatency1To5msOk

`func (o *OriginInspectorValues) GetLatency1To5msOk() (*int64, bool)`

GetLatency1To5msOk returns a tuple with the Latency1To5ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatency1To5ms

`func (o *OriginInspectorValues) SetLatency1To5ms(v int64)`

SetLatency1To5ms sets Latency1To5ms field to given value.

### HasLatency1To5ms

`func (o *OriginInspectorValues) HasLatency1To5ms() bool`

HasLatency1To5ms returns a boolean if a field has been set.

### GetLatency5To10ms

`func (o *OriginInspectorValues) GetLatency5To10ms() int64`

GetLatency5To10ms returns the Latency5To10ms field if non-nil, zero value otherwise.

### GetLatency5To10msOk

`func (o *OriginInspectorValues) GetLatency5To10msOk() (*int64, bool)`

GetLatency5To10msOk returns a tuple with the Latency5To10ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatency5To10ms

`func (o *OriginInspectorValues) SetLatency5To10ms(v int64)`

SetLatency5To10ms sets Latency5To10ms field to given value.

### HasLatency5To10ms

`func (o *OriginInspectorValues) HasLatency5To10ms() bool`

HasLatency5To10ms returns a boolean if a field has been set.

### GetLatency10To50ms

`func (o *OriginInspectorValues) GetLatency10To50ms() int64`

GetLatency10To50ms returns the Latency10To50ms field if non-nil, zero value otherwise.

### GetLatency10To50msOk

`func (o *OriginInspectorValues) GetLatency10To50msOk() (*int64, bool)`

GetLatency10To50msOk returns a tuple with the Latency10To50ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatency10To50ms

`func (o *OriginInspectorValues) SetLatency10To50ms(v int64)`

SetLatency10To50ms sets Latency10To50ms field to given value.

### HasLatency10To50ms

`func (o *OriginInspectorValues) HasLatency10To50ms() bool`

HasLatency10To50ms returns a boolean if a field has been set.

### GetLatency50To100ms

`func (o *OriginInspectorValues) GetLatency50To100ms() int64`

GetLatency50To100ms returns the Latency50To100ms field if non-nil, zero value otherwise.

### GetLatency50To100msOk

`func (o *OriginInspectorValues) GetLatency50To100msOk() (*int64, bool)`

GetLatency50To100msOk returns a tuple with the Latency50To100ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatency50To100ms

`func (o *OriginInspectorValues) SetLatency50To100ms(v int64)`

SetLatency50To100ms sets Latency50To100ms field to given value.

### HasLatency50To100ms

`func (o *OriginInspectorValues) HasLatency50To100ms() bool`

HasLatency50To100ms returns a boolean if a field has been set.

### GetLatency100To250ms

`func (o *OriginInspectorValues) GetLatency100To250ms() int64`

GetLatency100To250ms returns the Latency100To250ms field if non-nil, zero value otherwise.

### GetLatency100To250msOk

`func (o *OriginInspectorValues) GetLatency100To250msOk() (*int64, bool)`

GetLatency100To250msOk returns a tuple with the Latency100To250ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatency100To250ms

`func (o *OriginInspectorValues) SetLatency100To250ms(v int64)`

SetLatency100To250ms sets Latency100To250ms field to given value.

### HasLatency100To250ms

`func (o *OriginInspectorValues) HasLatency100To250ms() bool`

HasLatency100To250ms returns a boolean if a field has been set.

### GetLatency250To500ms

`func (o *OriginInspectorValues) GetLatency250To500ms() int64`

GetLatency250To500ms returns the Latency250To500ms field if non-nil, zero value otherwise.

### GetLatency250To500msOk

`func (o *OriginInspectorValues) GetLatency250To500msOk() (*int64, bool)`

GetLatency250To500msOk returns a tuple with the Latency250To500ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatency250To500ms

`func (o *OriginInspectorValues) SetLatency250To500ms(v int64)`

SetLatency250To500ms sets Latency250To500ms field to given value.

### HasLatency250To500ms

`func (o *OriginInspectorValues) HasLatency250To500ms() bool`

HasLatency250To500ms returns a boolean if a field has been set.

### GetLatency500To1000ms

`func (o *OriginInspectorValues) GetLatency500To1000ms() int64`

GetLatency500To1000ms returns the Latency500To1000ms field if non-nil, zero value otherwise.

### GetLatency500To1000msOk

`func (o *OriginInspectorValues) GetLatency500To1000msOk() (*int64, bool)`

GetLatency500To1000msOk returns a tuple with the Latency500To1000ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatency500To1000ms

`func (o *OriginInspectorValues) SetLatency500To1000ms(v int64)`

SetLatency500To1000ms sets Latency500To1000ms field to given value.

### HasLatency500To1000ms

`func (o *OriginInspectorValues) HasLatency500To1000ms() bool`

HasLatency500To1000ms returns a boolean if a field has been set.

### GetLatency1000To5000ms

`func (o *OriginInspectorValues) GetLatency1000To5000ms() int64`

GetLatency1000To5000ms returns the Latency1000To5000ms field if non-nil, zero value otherwise.

### GetLatency1000To5000msOk

`func (o *OriginInspectorValues) GetLatency1000To5000msOk() (*int64, bool)`

GetLatency1000To5000msOk returns a tuple with the Latency1000To5000ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatency1000To5000ms

`func (o *OriginInspectorValues) SetLatency1000To5000ms(v int64)`

SetLatency1000To5000ms sets Latency1000To5000ms field to given value.

### HasLatency1000To5000ms

`func (o *OriginInspectorValues) HasLatency1000To5000ms() bool`

HasLatency1000To5000ms returns a boolean if a field has been set.

### GetLatency5000To10000ms

`func (o *OriginInspectorValues) GetLatency5000To10000ms() int64`

GetLatency5000To10000ms returns the Latency5000To10000ms field if non-nil, zero value otherwise.

### GetLatency5000To10000msOk

`func (o *OriginInspectorValues) GetLatency5000To10000msOk() (*int64, bool)`

GetLatency5000To10000msOk returns a tuple with the Latency5000To10000ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatency5000To10000ms

`func (o *OriginInspectorValues) SetLatency5000To10000ms(v int64)`

SetLatency5000To10000ms sets Latency5000To10000ms field to given value.

### HasLatency5000To10000ms

`func (o *OriginInspectorValues) HasLatency5000To10000ms() bool`

HasLatency5000To10000ms returns a boolean if a field has been set.

### GetLatency10000To60000ms

`func (o *OriginInspectorValues) GetLatency10000To60000ms() int64`

GetLatency10000To60000ms returns the Latency10000To60000ms field if non-nil, zero value otherwise.

### GetLatency10000To60000msOk

`func (o *OriginInspectorValues) GetLatency10000To60000msOk() (*int64, bool)`

GetLatency10000To60000msOk returns a tuple with the Latency10000To60000ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatency10000To60000ms

`func (o *OriginInspectorValues) SetLatency10000To60000ms(v int64)`

SetLatency10000To60000ms sets Latency10000To60000ms field to given value.

### HasLatency10000To60000ms

`func (o *OriginInspectorValues) HasLatency10000To60000ms() bool`

HasLatency10000To60000ms returns a boolean if a field has been set.

### GetLatency60000ms

`func (o *OriginInspectorValues) GetLatency60000ms() int64`

GetLatency60000ms returns the Latency60000ms field if non-nil, zero value otherwise.

### GetLatency60000msOk

`func (o *OriginInspectorValues) GetLatency60000msOk() (*int64, bool)`

GetLatency60000msOk returns a tuple with the Latency60000ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatency60000ms

`func (o *OriginInspectorValues) SetLatency60000ms(v int64)`

SetLatency60000ms sets Latency60000ms field to given value.

### HasLatency60000ms

`func (o *OriginInspectorValues) HasLatency60000ms() bool`

HasLatency60000ms returns a boolean if a field has been set.

### GetWafResponses

`func (o *OriginInspectorValues) GetWafResponses() int64`

GetWafResponses returns the WafResponses field if non-nil, zero value otherwise.

### GetWafResponsesOk

`func (o *OriginInspectorValues) GetWafResponsesOk() (*int64, bool)`

GetWafResponsesOk returns a tuple with the WafResponses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafResponses

`func (o *OriginInspectorValues) SetWafResponses(v int64)`

SetWafResponses sets WafResponses field to given value.

### HasWafResponses

`func (o *OriginInspectorValues) HasWafResponses() bool`

HasWafResponses returns a boolean if a field has been set.

### GetWafRespHeaderBytes

`func (o *OriginInspectorValues) GetWafRespHeaderBytes() int64`

GetWafRespHeaderBytes returns the WafRespHeaderBytes field if non-nil, zero value otherwise.

### GetWafRespHeaderBytesOk

`func (o *OriginInspectorValues) GetWafRespHeaderBytesOk() (*int64, bool)`

GetWafRespHeaderBytesOk returns a tuple with the WafRespHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafRespHeaderBytes

`func (o *OriginInspectorValues) SetWafRespHeaderBytes(v int64)`

SetWafRespHeaderBytes sets WafRespHeaderBytes field to given value.

### HasWafRespHeaderBytes

`func (o *OriginInspectorValues) HasWafRespHeaderBytes() bool`

HasWafRespHeaderBytes returns a boolean if a field has been set.

### GetWafRespBodyBytes

`func (o *OriginInspectorValues) GetWafRespBodyBytes() int64`

GetWafRespBodyBytes returns the WafRespBodyBytes field if non-nil, zero value otherwise.

### GetWafRespBodyBytesOk

`func (o *OriginInspectorValues) GetWafRespBodyBytesOk() (*int64, bool)`

GetWafRespBodyBytesOk returns a tuple with the WafRespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafRespBodyBytes

`func (o *OriginInspectorValues) SetWafRespBodyBytes(v int64)`

SetWafRespBodyBytes sets WafRespBodyBytes field to given value.

### HasWafRespBodyBytes

`func (o *OriginInspectorValues) HasWafRespBodyBytes() bool`

HasWafRespBodyBytes returns a boolean if a field has been set.

### GetWafStatus1xx

`func (o *OriginInspectorValues) GetWafStatus1xx() int64`

GetWafStatus1xx returns the WafStatus1xx field if non-nil, zero value otherwise.

### GetWafStatus1xxOk

`func (o *OriginInspectorValues) GetWafStatus1xxOk() (*int64, bool)`

GetWafStatus1xxOk returns a tuple with the WafStatus1xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafStatus1xx

`func (o *OriginInspectorValues) SetWafStatus1xx(v int64)`

SetWafStatus1xx sets WafStatus1xx field to given value.

### HasWafStatus1xx

`func (o *OriginInspectorValues) HasWafStatus1xx() bool`

HasWafStatus1xx returns a boolean if a field has been set.

### GetWafStatus2xx

`func (o *OriginInspectorValues) GetWafStatus2xx() int64`

GetWafStatus2xx returns the WafStatus2xx field if non-nil, zero value otherwise.

### GetWafStatus2xxOk

`func (o *OriginInspectorValues) GetWafStatus2xxOk() (*int64, bool)`

GetWafStatus2xxOk returns a tuple with the WafStatus2xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafStatus2xx

`func (o *OriginInspectorValues) SetWafStatus2xx(v int64)`

SetWafStatus2xx sets WafStatus2xx field to given value.

### HasWafStatus2xx

`func (o *OriginInspectorValues) HasWafStatus2xx() bool`

HasWafStatus2xx returns a boolean if a field has been set.

### GetWafStatus3xx

`func (o *OriginInspectorValues) GetWafStatus3xx() int64`

GetWafStatus3xx returns the WafStatus3xx field if non-nil, zero value otherwise.

### GetWafStatus3xxOk

`func (o *OriginInspectorValues) GetWafStatus3xxOk() (*int64, bool)`

GetWafStatus3xxOk returns a tuple with the WafStatus3xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafStatus3xx

`func (o *OriginInspectorValues) SetWafStatus3xx(v int64)`

SetWafStatus3xx sets WafStatus3xx field to given value.

### HasWafStatus3xx

`func (o *OriginInspectorValues) HasWafStatus3xx() bool`

HasWafStatus3xx returns a boolean if a field has been set.

### GetWafStatus4xx

`func (o *OriginInspectorValues) GetWafStatus4xx() int64`

GetWafStatus4xx returns the WafStatus4xx field if non-nil, zero value otherwise.

### GetWafStatus4xxOk

`func (o *OriginInspectorValues) GetWafStatus4xxOk() (*int64, bool)`

GetWafStatus4xxOk returns a tuple with the WafStatus4xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafStatus4xx

`func (o *OriginInspectorValues) SetWafStatus4xx(v int64)`

SetWafStatus4xx sets WafStatus4xx field to given value.

### HasWafStatus4xx

`func (o *OriginInspectorValues) HasWafStatus4xx() bool`

HasWafStatus4xx returns a boolean if a field has been set.

### GetWafStatus5xx

`func (o *OriginInspectorValues) GetWafStatus5xx() int64`

GetWafStatus5xx returns the WafStatus5xx field if non-nil, zero value otherwise.

### GetWafStatus5xxOk

`func (o *OriginInspectorValues) GetWafStatus5xxOk() (*int64, bool)`

GetWafStatus5xxOk returns a tuple with the WafStatus5xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafStatus5xx

`func (o *OriginInspectorValues) SetWafStatus5xx(v int64)`

SetWafStatus5xx sets WafStatus5xx field to given value.

### HasWafStatus5xx

`func (o *OriginInspectorValues) HasWafStatus5xx() bool`

HasWafStatus5xx returns a boolean if a field has been set.

### GetWafStatus200

`func (o *OriginInspectorValues) GetWafStatus200() int64`

GetWafStatus200 returns the WafStatus200 field if non-nil, zero value otherwise.

### GetWafStatus200Ok

`func (o *OriginInspectorValues) GetWafStatus200Ok() (*int64, bool)`

GetWafStatus200Ok returns a tuple with the WafStatus200 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafStatus200

`func (o *OriginInspectorValues) SetWafStatus200(v int64)`

SetWafStatus200 sets WafStatus200 field to given value.

### HasWafStatus200

`func (o *OriginInspectorValues) HasWafStatus200() bool`

HasWafStatus200 returns a boolean if a field has been set.

### GetWafStatus204

`func (o *OriginInspectorValues) GetWafStatus204() int64`

GetWafStatus204 returns the WafStatus204 field if non-nil, zero value otherwise.

### GetWafStatus204Ok

`func (o *OriginInspectorValues) GetWafStatus204Ok() (*int64, bool)`

GetWafStatus204Ok returns a tuple with the WafStatus204 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafStatus204

`func (o *OriginInspectorValues) SetWafStatus204(v int64)`

SetWafStatus204 sets WafStatus204 field to given value.

### HasWafStatus204

`func (o *OriginInspectorValues) HasWafStatus204() bool`

HasWafStatus204 returns a boolean if a field has been set.

### GetWafStatus206

`func (o *OriginInspectorValues) GetWafStatus206() int64`

GetWafStatus206 returns the WafStatus206 field if non-nil, zero value otherwise.

### GetWafStatus206Ok

`func (o *OriginInspectorValues) GetWafStatus206Ok() (*int64, bool)`

GetWafStatus206Ok returns a tuple with the WafStatus206 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafStatus206

`func (o *OriginInspectorValues) SetWafStatus206(v int64)`

SetWafStatus206 sets WafStatus206 field to given value.

### HasWafStatus206

`func (o *OriginInspectorValues) HasWafStatus206() bool`

HasWafStatus206 returns a boolean if a field has been set.

### GetWafStatus301

`func (o *OriginInspectorValues) GetWafStatus301() int64`

GetWafStatus301 returns the WafStatus301 field if non-nil, zero value otherwise.

### GetWafStatus301Ok

`func (o *OriginInspectorValues) GetWafStatus301Ok() (*int64, bool)`

GetWafStatus301Ok returns a tuple with the WafStatus301 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafStatus301

`func (o *OriginInspectorValues) SetWafStatus301(v int64)`

SetWafStatus301 sets WafStatus301 field to given value.

### HasWafStatus301

`func (o *OriginInspectorValues) HasWafStatus301() bool`

HasWafStatus301 returns a boolean if a field has been set.

### GetWafStatus302

`func (o *OriginInspectorValues) GetWafStatus302() int64`

GetWafStatus302 returns the WafStatus302 field if non-nil, zero value otherwise.

### GetWafStatus302Ok

`func (o *OriginInspectorValues) GetWafStatus302Ok() (*int64, bool)`

GetWafStatus302Ok returns a tuple with the WafStatus302 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafStatus302

`func (o *OriginInspectorValues) SetWafStatus302(v int64)`

SetWafStatus302 sets WafStatus302 field to given value.

### HasWafStatus302

`func (o *OriginInspectorValues) HasWafStatus302() bool`

HasWafStatus302 returns a boolean if a field has been set.

### GetWafStatus304

`func (o *OriginInspectorValues) GetWafStatus304() int64`

GetWafStatus304 returns the WafStatus304 field if non-nil, zero value otherwise.

### GetWafStatus304Ok

`func (o *OriginInspectorValues) GetWafStatus304Ok() (*int64, bool)`

GetWafStatus304Ok returns a tuple with the WafStatus304 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafStatus304

`func (o *OriginInspectorValues) SetWafStatus304(v int64)`

SetWafStatus304 sets WafStatus304 field to given value.

### HasWafStatus304

`func (o *OriginInspectorValues) HasWafStatus304() bool`

HasWafStatus304 returns a boolean if a field has been set.

### GetWafStatus400

`func (o *OriginInspectorValues) GetWafStatus400() int64`

GetWafStatus400 returns the WafStatus400 field if non-nil, zero value otherwise.

### GetWafStatus400Ok

`func (o *OriginInspectorValues) GetWafStatus400Ok() (*int64, bool)`

GetWafStatus400Ok returns a tuple with the WafStatus400 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafStatus400

`func (o *OriginInspectorValues) SetWafStatus400(v int64)`

SetWafStatus400 sets WafStatus400 field to given value.

### HasWafStatus400

`func (o *OriginInspectorValues) HasWafStatus400() bool`

HasWafStatus400 returns a boolean if a field has been set.

### GetWafStatus401

`func (o *OriginInspectorValues) GetWafStatus401() int64`

GetWafStatus401 returns the WafStatus401 field if non-nil, zero value otherwise.

### GetWafStatus401Ok

`func (o *OriginInspectorValues) GetWafStatus401Ok() (*int64, bool)`

GetWafStatus401Ok returns a tuple with the WafStatus401 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafStatus401

`func (o *OriginInspectorValues) SetWafStatus401(v int64)`

SetWafStatus401 sets WafStatus401 field to given value.

### HasWafStatus401

`func (o *OriginInspectorValues) HasWafStatus401() bool`

HasWafStatus401 returns a boolean if a field has been set.

### GetWafStatus403

`func (o *OriginInspectorValues) GetWafStatus403() int64`

GetWafStatus403 returns the WafStatus403 field if non-nil, zero value otherwise.

### GetWafStatus403Ok

`func (o *OriginInspectorValues) GetWafStatus403Ok() (*int64, bool)`

GetWafStatus403Ok returns a tuple with the WafStatus403 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafStatus403

`func (o *OriginInspectorValues) SetWafStatus403(v int64)`

SetWafStatus403 sets WafStatus403 field to given value.

### HasWafStatus403

`func (o *OriginInspectorValues) HasWafStatus403() bool`

HasWafStatus403 returns a boolean if a field has been set.

### GetWafStatus404

`func (o *OriginInspectorValues) GetWafStatus404() int64`

GetWafStatus404 returns the WafStatus404 field if non-nil, zero value otherwise.

### GetWafStatus404Ok

`func (o *OriginInspectorValues) GetWafStatus404Ok() (*int64, bool)`

GetWafStatus404Ok returns a tuple with the WafStatus404 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafStatus404

`func (o *OriginInspectorValues) SetWafStatus404(v int64)`

SetWafStatus404 sets WafStatus404 field to given value.

### HasWafStatus404

`func (o *OriginInspectorValues) HasWafStatus404() bool`

HasWafStatus404 returns a boolean if a field has been set.

### GetWafStatus416

`func (o *OriginInspectorValues) GetWafStatus416() int64`

GetWafStatus416 returns the WafStatus416 field if non-nil, zero value otherwise.

### GetWafStatus416Ok

`func (o *OriginInspectorValues) GetWafStatus416Ok() (*int64, bool)`

GetWafStatus416Ok returns a tuple with the WafStatus416 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafStatus416

`func (o *OriginInspectorValues) SetWafStatus416(v int64)`

SetWafStatus416 sets WafStatus416 field to given value.

### HasWafStatus416

`func (o *OriginInspectorValues) HasWafStatus416() bool`

HasWafStatus416 returns a boolean if a field has been set.

### GetWafStatus429

`func (o *OriginInspectorValues) GetWafStatus429() int64`

GetWafStatus429 returns the WafStatus429 field if non-nil, zero value otherwise.

### GetWafStatus429Ok

`func (o *OriginInspectorValues) GetWafStatus429Ok() (*int64, bool)`

GetWafStatus429Ok returns a tuple with the WafStatus429 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafStatus429

`func (o *OriginInspectorValues) SetWafStatus429(v int64)`

SetWafStatus429 sets WafStatus429 field to given value.

### HasWafStatus429

`func (o *OriginInspectorValues) HasWafStatus429() bool`

HasWafStatus429 returns a boolean if a field has been set.

### GetWafStatus500

`func (o *OriginInspectorValues) GetWafStatus500() int64`

GetWafStatus500 returns the WafStatus500 field if non-nil, zero value otherwise.

### GetWafStatus500Ok

`func (o *OriginInspectorValues) GetWafStatus500Ok() (*int64, bool)`

GetWafStatus500Ok returns a tuple with the WafStatus500 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafStatus500

`func (o *OriginInspectorValues) SetWafStatus500(v int64)`

SetWafStatus500 sets WafStatus500 field to given value.

### HasWafStatus500

`func (o *OriginInspectorValues) HasWafStatus500() bool`

HasWafStatus500 returns a boolean if a field has been set.

### GetWafStatus501

`func (o *OriginInspectorValues) GetWafStatus501() int64`

GetWafStatus501 returns the WafStatus501 field if non-nil, zero value otherwise.

### GetWafStatus501Ok

`func (o *OriginInspectorValues) GetWafStatus501Ok() (*int64, bool)`

GetWafStatus501Ok returns a tuple with the WafStatus501 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafStatus501

`func (o *OriginInspectorValues) SetWafStatus501(v int64)`

SetWafStatus501 sets WafStatus501 field to given value.

### HasWafStatus501

`func (o *OriginInspectorValues) HasWafStatus501() bool`

HasWafStatus501 returns a boolean if a field has been set.

### GetWafStatus502

`func (o *OriginInspectorValues) GetWafStatus502() int64`

GetWafStatus502 returns the WafStatus502 field if non-nil, zero value otherwise.

### GetWafStatus502Ok

`func (o *OriginInspectorValues) GetWafStatus502Ok() (*int64, bool)`

GetWafStatus502Ok returns a tuple with the WafStatus502 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafStatus502

`func (o *OriginInspectorValues) SetWafStatus502(v int64)`

SetWafStatus502 sets WafStatus502 field to given value.

### HasWafStatus502

`func (o *OriginInspectorValues) HasWafStatus502() bool`

HasWafStatus502 returns a boolean if a field has been set.

### GetWafStatus503

`func (o *OriginInspectorValues) GetWafStatus503() int64`

GetWafStatus503 returns the WafStatus503 field if non-nil, zero value otherwise.

### GetWafStatus503Ok

`func (o *OriginInspectorValues) GetWafStatus503Ok() (*int64, bool)`

GetWafStatus503Ok returns a tuple with the WafStatus503 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafStatus503

`func (o *OriginInspectorValues) SetWafStatus503(v int64)`

SetWafStatus503 sets WafStatus503 field to given value.

### HasWafStatus503

`func (o *OriginInspectorValues) HasWafStatus503() bool`

HasWafStatus503 returns a boolean if a field has been set.

### GetWafStatus504

`func (o *OriginInspectorValues) GetWafStatus504() int64`

GetWafStatus504 returns the WafStatus504 field if non-nil, zero value otherwise.

### GetWafStatus504Ok

`func (o *OriginInspectorValues) GetWafStatus504Ok() (*int64, bool)`

GetWafStatus504Ok returns a tuple with the WafStatus504 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafStatus504

`func (o *OriginInspectorValues) SetWafStatus504(v int64)`

SetWafStatus504 sets WafStatus504 field to given value.

### HasWafStatus504

`func (o *OriginInspectorValues) HasWafStatus504() bool`

HasWafStatus504 returns a boolean if a field has been set.

### GetWafStatus505

`func (o *OriginInspectorValues) GetWafStatus505() int64`

GetWafStatus505 returns the WafStatus505 field if non-nil, zero value otherwise.

### GetWafStatus505Ok

`func (o *OriginInspectorValues) GetWafStatus505Ok() (*int64, bool)`

GetWafStatus505Ok returns a tuple with the WafStatus505 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafStatus505

`func (o *OriginInspectorValues) SetWafStatus505(v int64)`

SetWafStatus505 sets WafStatus505 field to given value.

### HasWafStatus505

`func (o *OriginInspectorValues) HasWafStatus505() bool`

HasWafStatus505 returns a boolean if a field has been set.

### GetWafStatus530

`func (o *OriginInspectorValues) GetWafStatus530() int64`

GetWafStatus530 returns the WafStatus530 field if non-nil, zero value otherwise.

### GetWafStatus530Ok

`func (o *OriginInspectorValues) GetWafStatus530Ok() (*int64, bool)`

GetWafStatus530Ok returns a tuple with the WafStatus530 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafStatus530

`func (o *OriginInspectorValues) SetWafStatus530(v int64)`

SetWafStatus530 sets WafStatus530 field to given value.

### HasWafStatus530

`func (o *OriginInspectorValues) HasWafStatus530() bool`

HasWafStatus530 returns a boolean if a field has been set.

### GetWafLatency0To1ms

`func (o *OriginInspectorValues) GetWafLatency0To1ms() int64`

GetWafLatency0To1ms returns the WafLatency0To1ms field if non-nil, zero value otherwise.

### GetWafLatency0To1msOk

`func (o *OriginInspectorValues) GetWafLatency0To1msOk() (*int64, bool)`

GetWafLatency0To1msOk returns a tuple with the WafLatency0To1ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafLatency0To1ms

`func (o *OriginInspectorValues) SetWafLatency0To1ms(v int64)`

SetWafLatency0To1ms sets WafLatency0To1ms field to given value.

### HasWafLatency0To1ms

`func (o *OriginInspectorValues) HasWafLatency0To1ms() bool`

HasWafLatency0To1ms returns a boolean if a field has been set.

### GetWafLatency1To5ms

`func (o *OriginInspectorValues) GetWafLatency1To5ms() int64`

GetWafLatency1To5ms returns the WafLatency1To5ms field if non-nil, zero value otherwise.

### GetWafLatency1To5msOk

`func (o *OriginInspectorValues) GetWafLatency1To5msOk() (*int64, bool)`

GetWafLatency1To5msOk returns a tuple with the WafLatency1To5ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafLatency1To5ms

`func (o *OriginInspectorValues) SetWafLatency1To5ms(v int64)`

SetWafLatency1To5ms sets WafLatency1To5ms field to given value.

### HasWafLatency1To5ms

`func (o *OriginInspectorValues) HasWafLatency1To5ms() bool`

HasWafLatency1To5ms returns a boolean if a field has been set.

### GetWafLatency5To10ms

`func (o *OriginInspectorValues) GetWafLatency5To10ms() int64`

GetWafLatency5To10ms returns the WafLatency5To10ms field if non-nil, zero value otherwise.

### GetWafLatency5To10msOk

`func (o *OriginInspectorValues) GetWafLatency5To10msOk() (*int64, bool)`

GetWafLatency5To10msOk returns a tuple with the WafLatency5To10ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafLatency5To10ms

`func (o *OriginInspectorValues) SetWafLatency5To10ms(v int64)`

SetWafLatency5To10ms sets WafLatency5To10ms field to given value.

### HasWafLatency5To10ms

`func (o *OriginInspectorValues) HasWafLatency5To10ms() bool`

HasWafLatency5To10ms returns a boolean if a field has been set.

### GetWafLatency10To50ms

`func (o *OriginInspectorValues) GetWafLatency10To50ms() int64`

GetWafLatency10To50ms returns the WafLatency10To50ms field if non-nil, zero value otherwise.

### GetWafLatency10To50msOk

`func (o *OriginInspectorValues) GetWafLatency10To50msOk() (*int64, bool)`

GetWafLatency10To50msOk returns a tuple with the WafLatency10To50ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafLatency10To50ms

`func (o *OriginInspectorValues) SetWafLatency10To50ms(v int64)`

SetWafLatency10To50ms sets WafLatency10To50ms field to given value.

### HasWafLatency10To50ms

`func (o *OriginInspectorValues) HasWafLatency10To50ms() bool`

HasWafLatency10To50ms returns a boolean if a field has been set.

### GetWafLatency50To100ms

`func (o *OriginInspectorValues) GetWafLatency50To100ms() int64`

GetWafLatency50To100ms returns the WafLatency50To100ms field if non-nil, zero value otherwise.

### GetWafLatency50To100msOk

`func (o *OriginInspectorValues) GetWafLatency50To100msOk() (*int64, bool)`

GetWafLatency50To100msOk returns a tuple with the WafLatency50To100ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafLatency50To100ms

`func (o *OriginInspectorValues) SetWafLatency50To100ms(v int64)`

SetWafLatency50To100ms sets WafLatency50To100ms field to given value.

### HasWafLatency50To100ms

`func (o *OriginInspectorValues) HasWafLatency50To100ms() bool`

HasWafLatency50To100ms returns a boolean if a field has been set.

### GetWafLatency100To250ms

`func (o *OriginInspectorValues) GetWafLatency100To250ms() int64`

GetWafLatency100To250ms returns the WafLatency100To250ms field if non-nil, zero value otherwise.

### GetWafLatency100To250msOk

`func (o *OriginInspectorValues) GetWafLatency100To250msOk() (*int64, bool)`

GetWafLatency100To250msOk returns a tuple with the WafLatency100To250ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafLatency100To250ms

`func (o *OriginInspectorValues) SetWafLatency100To250ms(v int64)`

SetWafLatency100To250ms sets WafLatency100To250ms field to given value.

### HasWafLatency100To250ms

`func (o *OriginInspectorValues) HasWafLatency100To250ms() bool`

HasWafLatency100To250ms returns a boolean if a field has been set.

### GetWafLatency250To500ms

`func (o *OriginInspectorValues) GetWafLatency250To500ms() int64`

GetWafLatency250To500ms returns the WafLatency250To500ms field if non-nil, zero value otherwise.

### GetWafLatency250To500msOk

`func (o *OriginInspectorValues) GetWafLatency250To500msOk() (*int64, bool)`

GetWafLatency250To500msOk returns a tuple with the WafLatency250To500ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafLatency250To500ms

`func (o *OriginInspectorValues) SetWafLatency250To500ms(v int64)`

SetWafLatency250To500ms sets WafLatency250To500ms field to given value.

### HasWafLatency250To500ms

`func (o *OriginInspectorValues) HasWafLatency250To500ms() bool`

HasWafLatency250To500ms returns a boolean if a field has been set.

### GetWafLatency500To1000ms

`func (o *OriginInspectorValues) GetWafLatency500To1000ms() int64`

GetWafLatency500To1000ms returns the WafLatency500To1000ms field if non-nil, zero value otherwise.

### GetWafLatency500To1000msOk

`func (o *OriginInspectorValues) GetWafLatency500To1000msOk() (*int64, bool)`

GetWafLatency500To1000msOk returns a tuple with the WafLatency500To1000ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafLatency500To1000ms

`func (o *OriginInspectorValues) SetWafLatency500To1000ms(v int64)`

SetWafLatency500To1000ms sets WafLatency500To1000ms field to given value.

### HasWafLatency500To1000ms

`func (o *OriginInspectorValues) HasWafLatency500To1000ms() bool`

HasWafLatency500To1000ms returns a boolean if a field has been set.

### GetWafLatency1000To5000ms

`func (o *OriginInspectorValues) GetWafLatency1000To5000ms() int64`

GetWafLatency1000To5000ms returns the WafLatency1000To5000ms field if non-nil, zero value otherwise.

### GetWafLatency1000To5000msOk

`func (o *OriginInspectorValues) GetWafLatency1000To5000msOk() (*int64, bool)`

GetWafLatency1000To5000msOk returns a tuple with the WafLatency1000To5000ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafLatency1000To5000ms

`func (o *OriginInspectorValues) SetWafLatency1000To5000ms(v int64)`

SetWafLatency1000To5000ms sets WafLatency1000To5000ms field to given value.

### HasWafLatency1000To5000ms

`func (o *OriginInspectorValues) HasWafLatency1000To5000ms() bool`

HasWafLatency1000To5000ms returns a boolean if a field has been set.

### GetWafLatency5000To10000ms

`func (o *OriginInspectorValues) GetWafLatency5000To10000ms() int64`

GetWafLatency5000To10000ms returns the WafLatency5000To10000ms field if non-nil, zero value otherwise.

### GetWafLatency5000To10000msOk

`func (o *OriginInspectorValues) GetWafLatency5000To10000msOk() (*int64, bool)`

GetWafLatency5000To10000msOk returns a tuple with the WafLatency5000To10000ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafLatency5000To10000ms

`func (o *OriginInspectorValues) SetWafLatency5000To10000ms(v int64)`

SetWafLatency5000To10000ms sets WafLatency5000To10000ms field to given value.

### HasWafLatency5000To10000ms

`func (o *OriginInspectorValues) HasWafLatency5000To10000ms() bool`

HasWafLatency5000To10000ms returns a boolean if a field has been set.

### GetWafLatency10000To60000ms

`func (o *OriginInspectorValues) GetWafLatency10000To60000ms() int64`

GetWafLatency10000To60000ms returns the WafLatency10000To60000ms field if non-nil, zero value otherwise.

### GetWafLatency10000To60000msOk

`func (o *OriginInspectorValues) GetWafLatency10000To60000msOk() (*int64, bool)`

GetWafLatency10000To60000msOk returns a tuple with the WafLatency10000To60000ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafLatency10000To60000ms

`func (o *OriginInspectorValues) SetWafLatency10000To60000ms(v int64)`

SetWafLatency10000To60000ms sets WafLatency10000To60000ms field to given value.

### HasWafLatency10000To60000ms

`func (o *OriginInspectorValues) HasWafLatency10000To60000ms() bool`

HasWafLatency10000To60000ms returns a boolean if a field has been set.

### GetWafLatency60000ms

`func (o *OriginInspectorValues) GetWafLatency60000ms() int64`

GetWafLatency60000ms returns the WafLatency60000ms field if non-nil, zero value otherwise.

### GetWafLatency60000msOk

`func (o *OriginInspectorValues) GetWafLatency60000msOk() (*int64, bool)`

GetWafLatency60000msOk returns a tuple with the WafLatency60000ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafLatency60000ms

`func (o *OriginInspectorValues) SetWafLatency60000ms(v int64)`

SetWafLatency60000ms sets WafLatency60000ms field to given value.

### HasWafLatency60000ms

`func (o *OriginInspectorValues) HasWafLatency60000ms() bool`

HasWafLatency60000ms returns a boolean if a field has been set.

### GetComputeResponses

`func (o *OriginInspectorValues) GetComputeResponses() int64`

GetComputeResponses returns the ComputeResponses field if non-nil, zero value otherwise.

### GetComputeResponsesOk

`func (o *OriginInspectorValues) GetComputeResponsesOk() (*int64, bool)`

GetComputeResponsesOk returns a tuple with the ComputeResponses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeResponses

`func (o *OriginInspectorValues) SetComputeResponses(v int64)`

SetComputeResponses sets ComputeResponses field to given value.

### HasComputeResponses

`func (o *OriginInspectorValues) HasComputeResponses() bool`

HasComputeResponses returns a boolean if a field has been set.

### GetComputeRespHeaderBytes

`func (o *OriginInspectorValues) GetComputeRespHeaderBytes() int64`

GetComputeRespHeaderBytes returns the ComputeRespHeaderBytes field if non-nil, zero value otherwise.

### GetComputeRespHeaderBytesOk

`func (o *OriginInspectorValues) GetComputeRespHeaderBytesOk() (*int64, bool)`

GetComputeRespHeaderBytesOk returns a tuple with the ComputeRespHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespHeaderBytes

`func (o *OriginInspectorValues) SetComputeRespHeaderBytes(v int64)`

SetComputeRespHeaderBytes sets ComputeRespHeaderBytes field to given value.

### HasComputeRespHeaderBytes

`func (o *OriginInspectorValues) HasComputeRespHeaderBytes() bool`

HasComputeRespHeaderBytes returns a boolean if a field has been set.

### GetComputeRespBodyBytes

`func (o *OriginInspectorValues) GetComputeRespBodyBytes() int64`

GetComputeRespBodyBytes returns the ComputeRespBodyBytes field if non-nil, zero value otherwise.

### GetComputeRespBodyBytesOk

`func (o *OriginInspectorValues) GetComputeRespBodyBytesOk() (*int64, bool)`

GetComputeRespBodyBytesOk returns a tuple with the ComputeRespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespBodyBytes

`func (o *OriginInspectorValues) SetComputeRespBodyBytes(v int64)`

SetComputeRespBodyBytes sets ComputeRespBodyBytes field to given value.

### HasComputeRespBodyBytes

`func (o *OriginInspectorValues) HasComputeRespBodyBytes() bool`

HasComputeRespBodyBytes returns a boolean if a field has been set.

### GetComputeStatus1xx

`func (o *OriginInspectorValues) GetComputeStatus1xx() int64`

GetComputeStatus1xx returns the ComputeStatus1xx field if non-nil, zero value otherwise.

### GetComputeStatus1xxOk

`func (o *OriginInspectorValues) GetComputeStatus1xxOk() (*int64, bool)`

GetComputeStatus1xxOk returns a tuple with the ComputeStatus1xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeStatus1xx

`func (o *OriginInspectorValues) SetComputeStatus1xx(v int64)`

SetComputeStatus1xx sets ComputeStatus1xx field to given value.

### HasComputeStatus1xx

`func (o *OriginInspectorValues) HasComputeStatus1xx() bool`

HasComputeStatus1xx returns a boolean if a field has been set.

### GetComputeStatus2xx

`func (o *OriginInspectorValues) GetComputeStatus2xx() int64`

GetComputeStatus2xx returns the ComputeStatus2xx field if non-nil, zero value otherwise.

### GetComputeStatus2xxOk

`func (o *OriginInspectorValues) GetComputeStatus2xxOk() (*int64, bool)`

GetComputeStatus2xxOk returns a tuple with the ComputeStatus2xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeStatus2xx

`func (o *OriginInspectorValues) SetComputeStatus2xx(v int64)`

SetComputeStatus2xx sets ComputeStatus2xx field to given value.

### HasComputeStatus2xx

`func (o *OriginInspectorValues) HasComputeStatus2xx() bool`

HasComputeStatus2xx returns a boolean if a field has been set.

### GetComputeStatus3xx

`func (o *OriginInspectorValues) GetComputeStatus3xx() int64`

GetComputeStatus3xx returns the ComputeStatus3xx field if non-nil, zero value otherwise.

### GetComputeStatus3xxOk

`func (o *OriginInspectorValues) GetComputeStatus3xxOk() (*int64, bool)`

GetComputeStatus3xxOk returns a tuple with the ComputeStatus3xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeStatus3xx

`func (o *OriginInspectorValues) SetComputeStatus3xx(v int64)`

SetComputeStatus3xx sets ComputeStatus3xx field to given value.

### HasComputeStatus3xx

`func (o *OriginInspectorValues) HasComputeStatus3xx() bool`

HasComputeStatus3xx returns a boolean if a field has been set.

### GetComputeStatus4xx

`func (o *OriginInspectorValues) GetComputeStatus4xx() int64`

GetComputeStatus4xx returns the ComputeStatus4xx field if non-nil, zero value otherwise.

### GetComputeStatus4xxOk

`func (o *OriginInspectorValues) GetComputeStatus4xxOk() (*int64, bool)`

GetComputeStatus4xxOk returns a tuple with the ComputeStatus4xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeStatus4xx

`func (o *OriginInspectorValues) SetComputeStatus4xx(v int64)`

SetComputeStatus4xx sets ComputeStatus4xx field to given value.

### HasComputeStatus4xx

`func (o *OriginInspectorValues) HasComputeStatus4xx() bool`

HasComputeStatus4xx returns a boolean if a field has been set.

### GetComputeStatus5xx

`func (o *OriginInspectorValues) GetComputeStatus5xx() int64`

GetComputeStatus5xx returns the ComputeStatus5xx field if non-nil, zero value otherwise.

### GetComputeStatus5xxOk

`func (o *OriginInspectorValues) GetComputeStatus5xxOk() (*int64, bool)`

GetComputeStatus5xxOk returns a tuple with the ComputeStatus5xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeStatus5xx

`func (o *OriginInspectorValues) SetComputeStatus5xx(v int64)`

SetComputeStatus5xx sets ComputeStatus5xx field to given value.

### HasComputeStatus5xx

`func (o *OriginInspectorValues) HasComputeStatus5xx() bool`

HasComputeStatus5xx returns a boolean if a field has been set.

### GetComputeStatus200

`func (o *OriginInspectorValues) GetComputeStatus200() int64`

GetComputeStatus200 returns the ComputeStatus200 field if non-nil, zero value otherwise.

### GetComputeStatus200Ok

`func (o *OriginInspectorValues) GetComputeStatus200Ok() (*int64, bool)`

GetComputeStatus200Ok returns a tuple with the ComputeStatus200 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeStatus200

`func (o *OriginInspectorValues) SetComputeStatus200(v int64)`

SetComputeStatus200 sets ComputeStatus200 field to given value.

### HasComputeStatus200

`func (o *OriginInspectorValues) HasComputeStatus200() bool`

HasComputeStatus200 returns a boolean if a field has been set.

### GetComputeStatus204

`func (o *OriginInspectorValues) GetComputeStatus204() int64`

GetComputeStatus204 returns the ComputeStatus204 field if non-nil, zero value otherwise.

### GetComputeStatus204Ok

`func (o *OriginInspectorValues) GetComputeStatus204Ok() (*int64, bool)`

GetComputeStatus204Ok returns a tuple with the ComputeStatus204 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeStatus204

`func (o *OriginInspectorValues) SetComputeStatus204(v int64)`

SetComputeStatus204 sets ComputeStatus204 field to given value.

### HasComputeStatus204

`func (o *OriginInspectorValues) HasComputeStatus204() bool`

HasComputeStatus204 returns a boolean if a field has been set.

### GetComputeStatus206

`func (o *OriginInspectorValues) GetComputeStatus206() int64`

GetComputeStatus206 returns the ComputeStatus206 field if non-nil, zero value otherwise.

### GetComputeStatus206Ok

`func (o *OriginInspectorValues) GetComputeStatus206Ok() (*int64, bool)`

GetComputeStatus206Ok returns a tuple with the ComputeStatus206 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeStatus206

`func (o *OriginInspectorValues) SetComputeStatus206(v int64)`

SetComputeStatus206 sets ComputeStatus206 field to given value.

### HasComputeStatus206

`func (o *OriginInspectorValues) HasComputeStatus206() bool`

HasComputeStatus206 returns a boolean if a field has been set.

### GetComputeStatus301

`func (o *OriginInspectorValues) GetComputeStatus301() int64`

GetComputeStatus301 returns the ComputeStatus301 field if non-nil, zero value otherwise.

### GetComputeStatus301Ok

`func (o *OriginInspectorValues) GetComputeStatus301Ok() (*int64, bool)`

GetComputeStatus301Ok returns a tuple with the ComputeStatus301 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeStatus301

`func (o *OriginInspectorValues) SetComputeStatus301(v int64)`

SetComputeStatus301 sets ComputeStatus301 field to given value.

### HasComputeStatus301

`func (o *OriginInspectorValues) HasComputeStatus301() bool`

HasComputeStatus301 returns a boolean if a field has been set.

### GetComputeStatus302

`func (o *OriginInspectorValues) GetComputeStatus302() int64`

GetComputeStatus302 returns the ComputeStatus302 field if non-nil, zero value otherwise.

### GetComputeStatus302Ok

`func (o *OriginInspectorValues) GetComputeStatus302Ok() (*int64, bool)`

GetComputeStatus302Ok returns a tuple with the ComputeStatus302 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeStatus302

`func (o *OriginInspectorValues) SetComputeStatus302(v int64)`

SetComputeStatus302 sets ComputeStatus302 field to given value.

### HasComputeStatus302

`func (o *OriginInspectorValues) HasComputeStatus302() bool`

HasComputeStatus302 returns a boolean if a field has been set.

### GetComputeStatus304

`func (o *OriginInspectorValues) GetComputeStatus304() int64`

GetComputeStatus304 returns the ComputeStatus304 field if non-nil, zero value otherwise.

### GetComputeStatus304Ok

`func (o *OriginInspectorValues) GetComputeStatus304Ok() (*int64, bool)`

GetComputeStatus304Ok returns a tuple with the ComputeStatus304 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeStatus304

`func (o *OriginInspectorValues) SetComputeStatus304(v int64)`

SetComputeStatus304 sets ComputeStatus304 field to given value.

### HasComputeStatus304

`func (o *OriginInspectorValues) HasComputeStatus304() bool`

HasComputeStatus304 returns a boolean if a field has been set.

### GetComputeStatus400

`func (o *OriginInspectorValues) GetComputeStatus400() int64`

GetComputeStatus400 returns the ComputeStatus400 field if non-nil, zero value otherwise.

### GetComputeStatus400Ok

`func (o *OriginInspectorValues) GetComputeStatus400Ok() (*int64, bool)`

GetComputeStatus400Ok returns a tuple with the ComputeStatus400 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeStatus400

`func (o *OriginInspectorValues) SetComputeStatus400(v int64)`

SetComputeStatus400 sets ComputeStatus400 field to given value.

### HasComputeStatus400

`func (o *OriginInspectorValues) HasComputeStatus400() bool`

HasComputeStatus400 returns a boolean if a field has been set.

### GetComputeStatus401

`func (o *OriginInspectorValues) GetComputeStatus401() int64`

GetComputeStatus401 returns the ComputeStatus401 field if non-nil, zero value otherwise.

### GetComputeStatus401Ok

`func (o *OriginInspectorValues) GetComputeStatus401Ok() (*int64, bool)`

GetComputeStatus401Ok returns a tuple with the ComputeStatus401 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeStatus401

`func (o *OriginInspectorValues) SetComputeStatus401(v int64)`

SetComputeStatus401 sets ComputeStatus401 field to given value.

### HasComputeStatus401

`func (o *OriginInspectorValues) HasComputeStatus401() bool`

HasComputeStatus401 returns a boolean if a field has been set.

### GetComputeStatus403

`func (o *OriginInspectorValues) GetComputeStatus403() int64`

GetComputeStatus403 returns the ComputeStatus403 field if non-nil, zero value otherwise.

### GetComputeStatus403Ok

`func (o *OriginInspectorValues) GetComputeStatus403Ok() (*int64, bool)`

GetComputeStatus403Ok returns a tuple with the ComputeStatus403 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeStatus403

`func (o *OriginInspectorValues) SetComputeStatus403(v int64)`

SetComputeStatus403 sets ComputeStatus403 field to given value.

### HasComputeStatus403

`func (o *OriginInspectorValues) HasComputeStatus403() bool`

HasComputeStatus403 returns a boolean if a field has been set.

### GetComputeStatus404

`func (o *OriginInspectorValues) GetComputeStatus404() int64`

GetComputeStatus404 returns the ComputeStatus404 field if non-nil, zero value otherwise.

### GetComputeStatus404Ok

`func (o *OriginInspectorValues) GetComputeStatus404Ok() (*int64, bool)`

GetComputeStatus404Ok returns a tuple with the ComputeStatus404 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeStatus404

`func (o *OriginInspectorValues) SetComputeStatus404(v int64)`

SetComputeStatus404 sets ComputeStatus404 field to given value.

### HasComputeStatus404

`func (o *OriginInspectorValues) HasComputeStatus404() bool`

HasComputeStatus404 returns a boolean if a field has been set.

### GetComputeStatus416

`func (o *OriginInspectorValues) GetComputeStatus416() int64`

GetComputeStatus416 returns the ComputeStatus416 field if non-nil, zero value otherwise.

### GetComputeStatus416Ok

`func (o *OriginInspectorValues) GetComputeStatus416Ok() (*int64, bool)`

GetComputeStatus416Ok returns a tuple with the ComputeStatus416 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeStatus416

`func (o *OriginInspectorValues) SetComputeStatus416(v int64)`

SetComputeStatus416 sets ComputeStatus416 field to given value.

### HasComputeStatus416

`func (o *OriginInspectorValues) HasComputeStatus416() bool`

HasComputeStatus416 returns a boolean if a field has been set.

### GetComputeStatus429

`func (o *OriginInspectorValues) GetComputeStatus429() int64`

GetComputeStatus429 returns the ComputeStatus429 field if non-nil, zero value otherwise.

### GetComputeStatus429Ok

`func (o *OriginInspectorValues) GetComputeStatus429Ok() (*int64, bool)`

GetComputeStatus429Ok returns a tuple with the ComputeStatus429 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeStatus429

`func (o *OriginInspectorValues) SetComputeStatus429(v int64)`

SetComputeStatus429 sets ComputeStatus429 field to given value.

### HasComputeStatus429

`func (o *OriginInspectorValues) HasComputeStatus429() bool`

HasComputeStatus429 returns a boolean if a field has been set.

### GetComputeStatus500

`func (o *OriginInspectorValues) GetComputeStatus500() int64`

GetComputeStatus500 returns the ComputeStatus500 field if non-nil, zero value otherwise.

### GetComputeStatus500Ok

`func (o *OriginInspectorValues) GetComputeStatus500Ok() (*int64, bool)`

GetComputeStatus500Ok returns a tuple with the ComputeStatus500 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeStatus500

`func (o *OriginInspectorValues) SetComputeStatus500(v int64)`

SetComputeStatus500 sets ComputeStatus500 field to given value.

### HasComputeStatus500

`func (o *OriginInspectorValues) HasComputeStatus500() bool`

HasComputeStatus500 returns a boolean if a field has been set.

### GetComputeStatus501

`func (o *OriginInspectorValues) GetComputeStatus501() int64`

GetComputeStatus501 returns the ComputeStatus501 field if non-nil, zero value otherwise.

### GetComputeStatus501Ok

`func (o *OriginInspectorValues) GetComputeStatus501Ok() (*int64, bool)`

GetComputeStatus501Ok returns a tuple with the ComputeStatus501 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeStatus501

`func (o *OriginInspectorValues) SetComputeStatus501(v int64)`

SetComputeStatus501 sets ComputeStatus501 field to given value.

### HasComputeStatus501

`func (o *OriginInspectorValues) HasComputeStatus501() bool`

HasComputeStatus501 returns a boolean if a field has been set.

### GetComputeStatus502

`func (o *OriginInspectorValues) GetComputeStatus502() int64`

GetComputeStatus502 returns the ComputeStatus502 field if non-nil, zero value otherwise.

### GetComputeStatus502Ok

`func (o *OriginInspectorValues) GetComputeStatus502Ok() (*int64, bool)`

GetComputeStatus502Ok returns a tuple with the ComputeStatus502 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeStatus502

`func (o *OriginInspectorValues) SetComputeStatus502(v int64)`

SetComputeStatus502 sets ComputeStatus502 field to given value.

### HasComputeStatus502

`func (o *OriginInspectorValues) HasComputeStatus502() bool`

HasComputeStatus502 returns a boolean if a field has been set.

### GetComputeStatus503

`func (o *OriginInspectorValues) GetComputeStatus503() int64`

GetComputeStatus503 returns the ComputeStatus503 field if non-nil, zero value otherwise.

### GetComputeStatus503Ok

`func (o *OriginInspectorValues) GetComputeStatus503Ok() (*int64, bool)`

GetComputeStatus503Ok returns a tuple with the ComputeStatus503 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeStatus503

`func (o *OriginInspectorValues) SetComputeStatus503(v int64)`

SetComputeStatus503 sets ComputeStatus503 field to given value.

### HasComputeStatus503

`func (o *OriginInspectorValues) HasComputeStatus503() bool`

HasComputeStatus503 returns a boolean if a field has been set.

### GetComputeStatus504

`func (o *OriginInspectorValues) GetComputeStatus504() int64`

GetComputeStatus504 returns the ComputeStatus504 field if non-nil, zero value otherwise.

### GetComputeStatus504Ok

`func (o *OriginInspectorValues) GetComputeStatus504Ok() (*int64, bool)`

GetComputeStatus504Ok returns a tuple with the ComputeStatus504 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeStatus504

`func (o *OriginInspectorValues) SetComputeStatus504(v int64)`

SetComputeStatus504 sets ComputeStatus504 field to given value.

### HasComputeStatus504

`func (o *OriginInspectorValues) HasComputeStatus504() bool`

HasComputeStatus504 returns a boolean if a field has been set.

### GetComputeStatus505

`func (o *OriginInspectorValues) GetComputeStatus505() int64`

GetComputeStatus505 returns the ComputeStatus505 field if non-nil, zero value otherwise.

### GetComputeStatus505Ok

`func (o *OriginInspectorValues) GetComputeStatus505Ok() (*int64, bool)`

GetComputeStatus505Ok returns a tuple with the ComputeStatus505 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeStatus505

`func (o *OriginInspectorValues) SetComputeStatus505(v int64)`

SetComputeStatus505 sets ComputeStatus505 field to given value.

### HasComputeStatus505

`func (o *OriginInspectorValues) HasComputeStatus505() bool`

HasComputeStatus505 returns a boolean if a field has been set.

### GetComputeStatus530

`func (o *OriginInspectorValues) GetComputeStatus530() int64`

GetComputeStatus530 returns the ComputeStatus530 field if non-nil, zero value otherwise.

### GetComputeStatus530Ok

`func (o *OriginInspectorValues) GetComputeStatus530Ok() (*int64, bool)`

GetComputeStatus530Ok returns a tuple with the ComputeStatus530 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeStatus530

`func (o *OriginInspectorValues) SetComputeStatus530(v int64)`

SetComputeStatus530 sets ComputeStatus530 field to given value.

### HasComputeStatus530

`func (o *OriginInspectorValues) HasComputeStatus530() bool`

HasComputeStatus530 returns a boolean if a field has been set.

### GetComputeLatency0To1ms

`func (o *OriginInspectorValues) GetComputeLatency0To1ms() int64`

GetComputeLatency0To1ms returns the ComputeLatency0To1ms field if non-nil, zero value otherwise.

### GetComputeLatency0To1msOk

`func (o *OriginInspectorValues) GetComputeLatency0To1msOk() (*int64, bool)`

GetComputeLatency0To1msOk returns a tuple with the ComputeLatency0To1ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeLatency0To1ms

`func (o *OriginInspectorValues) SetComputeLatency0To1ms(v int64)`

SetComputeLatency0To1ms sets ComputeLatency0To1ms field to given value.

### HasComputeLatency0To1ms

`func (o *OriginInspectorValues) HasComputeLatency0To1ms() bool`

HasComputeLatency0To1ms returns a boolean if a field has been set.

### GetComputeLatency1To5ms

`func (o *OriginInspectorValues) GetComputeLatency1To5ms() int64`

GetComputeLatency1To5ms returns the ComputeLatency1To5ms field if non-nil, zero value otherwise.

### GetComputeLatency1To5msOk

`func (o *OriginInspectorValues) GetComputeLatency1To5msOk() (*int64, bool)`

GetComputeLatency1To5msOk returns a tuple with the ComputeLatency1To5ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeLatency1To5ms

`func (o *OriginInspectorValues) SetComputeLatency1To5ms(v int64)`

SetComputeLatency1To5ms sets ComputeLatency1To5ms field to given value.

### HasComputeLatency1To5ms

`func (o *OriginInspectorValues) HasComputeLatency1To5ms() bool`

HasComputeLatency1To5ms returns a boolean if a field has been set.

### GetComputeLatency5To10ms

`func (o *OriginInspectorValues) GetComputeLatency5To10ms() int64`

GetComputeLatency5To10ms returns the ComputeLatency5To10ms field if non-nil, zero value otherwise.

### GetComputeLatency5To10msOk

`func (o *OriginInspectorValues) GetComputeLatency5To10msOk() (*int64, bool)`

GetComputeLatency5To10msOk returns a tuple with the ComputeLatency5To10ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeLatency5To10ms

`func (o *OriginInspectorValues) SetComputeLatency5To10ms(v int64)`

SetComputeLatency5To10ms sets ComputeLatency5To10ms field to given value.

### HasComputeLatency5To10ms

`func (o *OriginInspectorValues) HasComputeLatency5To10ms() bool`

HasComputeLatency5To10ms returns a boolean if a field has been set.

### GetComputeLatency10To50ms

`func (o *OriginInspectorValues) GetComputeLatency10To50ms() int64`

GetComputeLatency10To50ms returns the ComputeLatency10To50ms field if non-nil, zero value otherwise.

### GetComputeLatency10To50msOk

`func (o *OriginInspectorValues) GetComputeLatency10To50msOk() (*int64, bool)`

GetComputeLatency10To50msOk returns a tuple with the ComputeLatency10To50ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeLatency10To50ms

`func (o *OriginInspectorValues) SetComputeLatency10To50ms(v int64)`

SetComputeLatency10To50ms sets ComputeLatency10To50ms field to given value.

### HasComputeLatency10To50ms

`func (o *OriginInspectorValues) HasComputeLatency10To50ms() bool`

HasComputeLatency10To50ms returns a boolean if a field has been set.

### GetComputeLatency50To100ms

`func (o *OriginInspectorValues) GetComputeLatency50To100ms() int64`

GetComputeLatency50To100ms returns the ComputeLatency50To100ms field if non-nil, zero value otherwise.

### GetComputeLatency50To100msOk

`func (o *OriginInspectorValues) GetComputeLatency50To100msOk() (*int64, bool)`

GetComputeLatency50To100msOk returns a tuple with the ComputeLatency50To100ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeLatency50To100ms

`func (o *OriginInspectorValues) SetComputeLatency50To100ms(v int64)`

SetComputeLatency50To100ms sets ComputeLatency50To100ms field to given value.

### HasComputeLatency50To100ms

`func (o *OriginInspectorValues) HasComputeLatency50To100ms() bool`

HasComputeLatency50To100ms returns a boolean if a field has been set.

### GetComputeLatency100To250ms

`func (o *OriginInspectorValues) GetComputeLatency100To250ms() int64`

GetComputeLatency100To250ms returns the ComputeLatency100To250ms field if non-nil, zero value otherwise.

### GetComputeLatency100To250msOk

`func (o *OriginInspectorValues) GetComputeLatency100To250msOk() (*int64, bool)`

GetComputeLatency100To250msOk returns a tuple with the ComputeLatency100To250ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeLatency100To250ms

`func (o *OriginInspectorValues) SetComputeLatency100To250ms(v int64)`

SetComputeLatency100To250ms sets ComputeLatency100To250ms field to given value.

### HasComputeLatency100To250ms

`func (o *OriginInspectorValues) HasComputeLatency100To250ms() bool`

HasComputeLatency100To250ms returns a boolean if a field has been set.

### GetComputeLatency250To500ms

`func (o *OriginInspectorValues) GetComputeLatency250To500ms() int64`

GetComputeLatency250To500ms returns the ComputeLatency250To500ms field if non-nil, zero value otherwise.

### GetComputeLatency250To500msOk

`func (o *OriginInspectorValues) GetComputeLatency250To500msOk() (*int64, bool)`

GetComputeLatency250To500msOk returns a tuple with the ComputeLatency250To500ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeLatency250To500ms

`func (o *OriginInspectorValues) SetComputeLatency250To500ms(v int64)`

SetComputeLatency250To500ms sets ComputeLatency250To500ms field to given value.

### HasComputeLatency250To500ms

`func (o *OriginInspectorValues) HasComputeLatency250To500ms() bool`

HasComputeLatency250To500ms returns a boolean if a field has been set.

### GetComputeLatency500To1000ms

`func (o *OriginInspectorValues) GetComputeLatency500To1000ms() int64`

GetComputeLatency500To1000ms returns the ComputeLatency500To1000ms field if non-nil, zero value otherwise.

### GetComputeLatency500To1000msOk

`func (o *OriginInspectorValues) GetComputeLatency500To1000msOk() (*int64, bool)`

GetComputeLatency500To1000msOk returns a tuple with the ComputeLatency500To1000ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeLatency500To1000ms

`func (o *OriginInspectorValues) SetComputeLatency500To1000ms(v int64)`

SetComputeLatency500To1000ms sets ComputeLatency500To1000ms field to given value.

### HasComputeLatency500To1000ms

`func (o *OriginInspectorValues) HasComputeLatency500To1000ms() bool`

HasComputeLatency500To1000ms returns a boolean if a field has been set.

### GetComputeLatency1000To5000ms

`func (o *OriginInspectorValues) GetComputeLatency1000To5000ms() int64`

GetComputeLatency1000To5000ms returns the ComputeLatency1000To5000ms field if non-nil, zero value otherwise.

### GetComputeLatency1000To5000msOk

`func (o *OriginInspectorValues) GetComputeLatency1000To5000msOk() (*int64, bool)`

GetComputeLatency1000To5000msOk returns a tuple with the ComputeLatency1000To5000ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeLatency1000To5000ms

`func (o *OriginInspectorValues) SetComputeLatency1000To5000ms(v int64)`

SetComputeLatency1000To5000ms sets ComputeLatency1000To5000ms field to given value.

### HasComputeLatency1000To5000ms

`func (o *OriginInspectorValues) HasComputeLatency1000To5000ms() bool`

HasComputeLatency1000To5000ms returns a boolean if a field has been set.

### GetComputeLatency5000To10000ms

`func (o *OriginInspectorValues) GetComputeLatency5000To10000ms() int64`

GetComputeLatency5000To10000ms returns the ComputeLatency5000To10000ms field if non-nil, zero value otherwise.

### GetComputeLatency5000To10000msOk

`func (o *OriginInspectorValues) GetComputeLatency5000To10000msOk() (*int64, bool)`

GetComputeLatency5000To10000msOk returns a tuple with the ComputeLatency5000To10000ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeLatency5000To10000ms

`func (o *OriginInspectorValues) SetComputeLatency5000To10000ms(v int64)`

SetComputeLatency5000To10000ms sets ComputeLatency5000To10000ms field to given value.

### HasComputeLatency5000To10000ms

`func (o *OriginInspectorValues) HasComputeLatency5000To10000ms() bool`

HasComputeLatency5000To10000ms returns a boolean if a field has been set.

### GetComputeLatency10000To60000ms

`func (o *OriginInspectorValues) GetComputeLatency10000To60000ms() int64`

GetComputeLatency10000To60000ms returns the ComputeLatency10000To60000ms field if non-nil, zero value otherwise.

### GetComputeLatency10000To60000msOk

`func (o *OriginInspectorValues) GetComputeLatency10000To60000msOk() (*int64, bool)`

GetComputeLatency10000To60000msOk returns a tuple with the ComputeLatency10000To60000ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeLatency10000To60000ms

`func (o *OriginInspectorValues) SetComputeLatency10000To60000ms(v int64)`

SetComputeLatency10000To60000ms sets ComputeLatency10000To60000ms field to given value.

### HasComputeLatency10000To60000ms

`func (o *OriginInspectorValues) HasComputeLatency10000To60000ms() bool`

HasComputeLatency10000To60000ms returns a boolean if a field has been set.

### GetComputeLatency60000ms

`func (o *OriginInspectorValues) GetComputeLatency60000ms() int64`

GetComputeLatency60000ms returns the ComputeLatency60000ms field if non-nil, zero value otherwise.

### GetComputeLatency60000msOk

`func (o *OriginInspectorValues) GetComputeLatency60000msOk() (*int64, bool)`

GetComputeLatency60000msOk returns a tuple with the ComputeLatency60000ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeLatency60000ms

`func (o *OriginInspectorValues) SetComputeLatency60000ms(v int64)`

SetComputeLatency60000ms sets ComputeLatency60000ms field to given value.

### HasComputeLatency60000ms

`func (o *OriginInspectorValues) HasComputeLatency60000ms() bool`

HasComputeLatency60000ms returns a boolean if a field has been set.

### GetAllResponses

`func (o *OriginInspectorValues) GetAllResponses() int64`

GetAllResponses returns the AllResponses field if non-nil, zero value otherwise.

### GetAllResponsesOk

`func (o *OriginInspectorValues) GetAllResponsesOk() (*int64, bool)`

GetAllResponsesOk returns a tuple with the AllResponses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllResponses

`func (o *OriginInspectorValues) SetAllResponses(v int64)`

SetAllResponses sets AllResponses field to given value.

### HasAllResponses

`func (o *OriginInspectorValues) HasAllResponses() bool`

HasAllResponses returns a boolean if a field has been set.

### GetAllRespHeaderBytes

`func (o *OriginInspectorValues) GetAllRespHeaderBytes() int64`

GetAllRespHeaderBytes returns the AllRespHeaderBytes field if non-nil, zero value otherwise.

### GetAllRespHeaderBytesOk

`func (o *OriginInspectorValues) GetAllRespHeaderBytesOk() (*int64, bool)`

GetAllRespHeaderBytesOk returns a tuple with the AllRespHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllRespHeaderBytes

`func (o *OriginInspectorValues) SetAllRespHeaderBytes(v int64)`

SetAllRespHeaderBytes sets AllRespHeaderBytes field to given value.

### HasAllRespHeaderBytes

`func (o *OriginInspectorValues) HasAllRespHeaderBytes() bool`

HasAllRespHeaderBytes returns a boolean if a field has been set.

### GetAllRespBodyBytes

`func (o *OriginInspectorValues) GetAllRespBodyBytes() int64`

GetAllRespBodyBytes returns the AllRespBodyBytes field if non-nil, zero value otherwise.

### GetAllRespBodyBytesOk

`func (o *OriginInspectorValues) GetAllRespBodyBytesOk() (*int64, bool)`

GetAllRespBodyBytesOk returns a tuple with the AllRespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllRespBodyBytes

`func (o *OriginInspectorValues) SetAllRespBodyBytes(v int64)`

SetAllRespBodyBytes sets AllRespBodyBytes field to given value.

### HasAllRespBodyBytes

`func (o *OriginInspectorValues) HasAllRespBodyBytes() bool`

HasAllRespBodyBytes returns a boolean if a field has been set.

### GetAllStatus1xx

`func (o *OriginInspectorValues) GetAllStatus1xx() int64`

GetAllStatus1xx returns the AllStatus1xx field if non-nil, zero value otherwise.

### GetAllStatus1xxOk

`func (o *OriginInspectorValues) GetAllStatus1xxOk() (*int64, bool)`

GetAllStatus1xxOk returns a tuple with the AllStatus1xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllStatus1xx

`func (o *OriginInspectorValues) SetAllStatus1xx(v int64)`

SetAllStatus1xx sets AllStatus1xx field to given value.

### HasAllStatus1xx

`func (o *OriginInspectorValues) HasAllStatus1xx() bool`

HasAllStatus1xx returns a boolean if a field has been set.

### GetAllStatus2xx

`func (o *OriginInspectorValues) GetAllStatus2xx() int64`

GetAllStatus2xx returns the AllStatus2xx field if non-nil, zero value otherwise.

### GetAllStatus2xxOk

`func (o *OriginInspectorValues) GetAllStatus2xxOk() (*int64, bool)`

GetAllStatus2xxOk returns a tuple with the AllStatus2xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllStatus2xx

`func (o *OriginInspectorValues) SetAllStatus2xx(v int64)`

SetAllStatus2xx sets AllStatus2xx field to given value.

### HasAllStatus2xx

`func (o *OriginInspectorValues) HasAllStatus2xx() bool`

HasAllStatus2xx returns a boolean if a field has been set.

### GetAllStatus3xx

`func (o *OriginInspectorValues) GetAllStatus3xx() int64`

GetAllStatus3xx returns the AllStatus3xx field if non-nil, zero value otherwise.

### GetAllStatus3xxOk

`func (o *OriginInspectorValues) GetAllStatus3xxOk() (*int64, bool)`

GetAllStatus3xxOk returns a tuple with the AllStatus3xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllStatus3xx

`func (o *OriginInspectorValues) SetAllStatus3xx(v int64)`

SetAllStatus3xx sets AllStatus3xx field to given value.

### HasAllStatus3xx

`func (o *OriginInspectorValues) HasAllStatus3xx() bool`

HasAllStatus3xx returns a boolean if a field has been set.

### GetAllStatus4xx

`func (o *OriginInspectorValues) GetAllStatus4xx() int64`

GetAllStatus4xx returns the AllStatus4xx field if non-nil, zero value otherwise.

### GetAllStatus4xxOk

`func (o *OriginInspectorValues) GetAllStatus4xxOk() (*int64, bool)`

GetAllStatus4xxOk returns a tuple with the AllStatus4xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllStatus4xx

`func (o *OriginInspectorValues) SetAllStatus4xx(v int64)`

SetAllStatus4xx sets AllStatus4xx field to given value.

### HasAllStatus4xx

`func (o *OriginInspectorValues) HasAllStatus4xx() bool`

HasAllStatus4xx returns a boolean if a field has been set.

### GetAllStatus5xx

`func (o *OriginInspectorValues) GetAllStatus5xx() int64`

GetAllStatus5xx returns the AllStatus5xx field if non-nil, zero value otherwise.

### GetAllStatus5xxOk

`func (o *OriginInspectorValues) GetAllStatus5xxOk() (*int64, bool)`

GetAllStatus5xxOk returns a tuple with the AllStatus5xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllStatus5xx

`func (o *OriginInspectorValues) SetAllStatus5xx(v int64)`

SetAllStatus5xx sets AllStatus5xx field to given value.

### HasAllStatus5xx

`func (o *OriginInspectorValues) HasAllStatus5xx() bool`

HasAllStatus5xx returns a boolean if a field has been set.

### GetAllStatus200

`func (o *OriginInspectorValues) GetAllStatus200() int64`

GetAllStatus200 returns the AllStatus200 field if non-nil, zero value otherwise.

### GetAllStatus200Ok

`func (o *OriginInspectorValues) GetAllStatus200Ok() (*int64, bool)`

GetAllStatus200Ok returns a tuple with the AllStatus200 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllStatus200

`func (o *OriginInspectorValues) SetAllStatus200(v int64)`

SetAllStatus200 sets AllStatus200 field to given value.

### HasAllStatus200

`func (o *OriginInspectorValues) HasAllStatus200() bool`

HasAllStatus200 returns a boolean if a field has been set.

### GetAllStatus204

`func (o *OriginInspectorValues) GetAllStatus204() int64`

GetAllStatus204 returns the AllStatus204 field if non-nil, zero value otherwise.

### GetAllStatus204Ok

`func (o *OriginInspectorValues) GetAllStatus204Ok() (*int64, bool)`

GetAllStatus204Ok returns a tuple with the AllStatus204 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllStatus204

`func (o *OriginInspectorValues) SetAllStatus204(v int64)`

SetAllStatus204 sets AllStatus204 field to given value.

### HasAllStatus204

`func (o *OriginInspectorValues) HasAllStatus204() bool`

HasAllStatus204 returns a boolean if a field has been set.

### GetAllStatus206

`func (o *OriginInspectorValues) GetAllStatus206() int64`

GetAllStatus206 returns the AllStatus206 field if non-nil, zero value otherwise.

### GetAllStatus206Ok

`func (o *OriginInspectorValues) GetAllStatus206Ok() (*int64, bool)`

GetAllStatus206Ok returns a tuple with the AllStatus206 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllStatus206

`func (o *OriginInspectorValues) SetAllStatus206(v int64)`

SetAllStatus206 sets AllStatus206 field to given value.

### HasAllStatus206

`func (o *OriginInspectorValues) HasAllStatus206() bool`

HasAllStatus206 returns a boolean if a field has been set.

### GetAllStatus301

`func (o *OriginInspectorValues) GetAllStatus301() int64`

GetAllStatus301 returns the AllStatus301 field if non-nil, zero value otherwise.

### GetAllStatus301Ok

`func (o *OriginInspectorValues) GetAllStatus301Ok() (*int64, bool)`

GetAllStatus301Ok returns a tuple with the AllStatus301 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllStatus301

`func (o *OriginInspectorValues) SetAllStatus301(v int64)`

SetAllStatus301 sets AllStatus301 field to given value.

### HasAllStatus301

`func (o *OriginInspectorValues) HasAllStatus301() bool`

HasAllStatus301 returns a boolean if a field has been set.

### GetAllStatus302

`func (o *OriginInspectorValues) GetAllStatus302() int64`

GetAllStatus302 returns the AllStatus302 field if non-nil, zero value otherwise.

### GetAllStatus302Ok

`func (o *OriginInspectorValues) GetAllStatus302Ok() (*int64, bool)`

GetAllStatus302Ok returns a tuple with the AllStatus302 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllStatus302

`func (o *OriginInspectorValues) SetAllStatus302(v int64)`

SetAllStatus302 sets AllStatus302 field to given value.

### HasAllStatus302

`func (o *OriginInspectorValues) HasAllStatus302() bool`

HasAllStatus302 returns a boolean if a field has been set.

### GetAllStatus304

`func (o *OriginInspectorValues) GetAllStatus304() int64`

GetAllStatus304 returns the AllStatus304 field if non-nil, zero value otherwise.

### GetAllStatus304Ok

`func (o *OriginInspectorValues) GetAllStatus304Ok() (*int64, bool)`

GetAllStatus304Ok returns a tuple with the AllStatus304 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllStatus304

`func (o *OriginInspectorValues) SetAllStatus304(v int64)`

SetAllStatus304 sets AllStatus304 field to given value.

### HasAllStatus304

`func (o *OriginInspectorValues) HasAllStatus304() bool`

HasAllStatus304 returns a boolean if a field has been set.

### GetAllStatus400

`func (o *OriginInspectorValues) GetAllStatus400() int64`

GetAllStatus400 returns the AllStatus400 field if non-nil, zero value otherwise.

### GetAllStatus400Ok

`func (o *OriginInspectorValues) GetAllStatus400Ok() (*int64, bool)`

GetAllStatus400Ok returns a tuple with the AllStatus400 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllStatus400

`func (o *OriginInspectorValues) SetAllStatus400(v int64)`

SetAllStatus400 sets AllStatus400 field to given value.

### HasAllStatus400

`func (o *OriginInspectorValues) HasAllStatus400() bool`

HasAllStatus400 returns a boolean if a field has been set.

### GetAllStatus401

`func (o *OriginInspectorValues) GetAllStatus401() int64`

GetAllStatus401 returns the AllStatus401 field if non-nil, zero value otherwise.

### GetAllStatus401Ok

`func (o *OriginInspectorValues) GetAllStatus401Ok() (*int64, bool)`

GetAllStatus401Ok returns a tuple with the AllStatus401 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllStatus401

`func (o *OriginInspectorValues) SetAllStatus401(v int64)`

SetAllStatus401 sets AllStatus401 field to given value.

### HasAllStatus401

`func (o *OriginInspectorValues) HasAllStatus401() bool`

HasAllStatus401 returns a boolean if a field has been set.

### GetAllStatus403

`func (o *OriginInspectorValues) GetAllStatus403() int64`

GetAllStatus403 returns the AllStatus403 field if non-nil, zero value otherwise.

### GetAllStatus403Ok

`func (o *OriginInspectorValues) GetAllStatus403Ok() (*int64, bool)`

GetAllStatus403Ok returns a tuple with the AllStatus403 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllStatus403

`func (o *OriginInspectorValues) SetAllStatus403(v int64)`

SetAllStatus403 sets AllStatus403 field to given value.

### HasAllStatus403

`func (o *OriginInspectorValues) HasAllStatus403() bool`

HasAllStatus403 returns a boolean if a field has been set.

### GetAllStatus404

`func (o *OriginInspectorValues) GetAllStatus404() int64`

GetAllStatus404 returns the AllStatus404 field if non-nil, zero value otherwise.

### GetAllStatus404Ok

`func (o *OriginInspectorValues) GetAllStatus404Ok() (*int64, bool)`

GetAllStatus404Ok returns a tuple with the AllStatus404 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllStatus404

`func (o *OriginInspectorValues) SetAllStatus404(v int64)`

SetAllStatus404 sets AllStatus404 field to given value.

### HasAllStatus404

`func (o *OriginInspectorValues) HasAllStatus404() bool`

HasAllStatus404 returns a boolean if a field has been set.

### GetAllStatus416

`func (o *OriginInspectorValues) GetAllStatus416() int64`

GetAllStatus416 returns the AllStatus416 field if non-nil, zero value otherwise.

### GetAllStatus416Ok

`func (o *OriginInspectorValues) GetAllStatus416Ok() (*int64, bool)`

GetAllStatus416Ok returns a tuple with the AllStatus416 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllStatus416

`func (o *OriginInspectorValues) SetAllStatus416(v int64)`

SetAllStatus416 sets AllStatus416 field to given value.

### HasAllStatus416

`func (o *OriginInspectorValues) HasAllStatus416() bool`

HasAllStatus416 returns a boolean if a field has been set.

### GetAllStatus429

`func (o *OriginInspectorValues) GetAllStatus429() int64`

GetAllStatus429 returns the AllStatus429 field if non-nil, zero value otherwise.

### GetAllStatus429Ok

`func (o *OriginInspectorValues) GetAllStatus429Ok() (*int64, bool)`

GetAllStatus429Ok returns a tuple with the AllStatus429 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllStatus429

`func (o *OriginInspectorValues) SetAllStatus429(v int64)`

SetAllStatus429 sets AllStatus429 field to given value.

### HasAllStatus429

`func (o *OriginInspectorValues) HasAllStatus429() bool`

HasAllStatus429 returns a boolean if a field has been set.

### GetAllStatus500

`func (o *OriginInspectorValues) GetAllStatus500() int64`

GetAllStatus500 returns the AllStatus500 field if non-nil, zero value otherwise.

### GetAllStatus500Ok

`func (o *OriginInspectorValues) GetAllStatus500Ok() (*int64, bool)`

GetAllStatus500Ok returns a tuple with the AllStatus500 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllStatus500

`func (o *OriginInspectorValues) SetAllStatus500(v int64)`

SetAllStatus500 sets AllStatus500 field to given value.

### HasAllStatus500

`func (o *OriginInspectorValues) HasAllStatus500() bool`

HasAllStatus500 returns a boolean if a field has been set.

### GetAllStatus501

`func (o *OriginInspectorValues) GetAllStatus501() int64`

GetAllStatus501 returns the AllStatus501 field if non-nil, zero value otherwise.

### GetAllStatus501Ok

`func (o *OriginInspectorValues) GetAllStatus501Ok() (*int64, bool)`

GetAllStatus501Ok returns a tuple with the AllStatus501 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllStatus501

`func (o *OriginInspectorValues) SetAllStatus501(v int64)`

SetAllStatus501 sets AllStatus501 field to given value.

### HasAllStatus501

`func (o *OriginInspectorValues) HasAllStatus501() bool`

HasAllStatus501 returns a boolean if a field has been set.

### GetAllStatus502

`func (o *OriginInspectorValues) GetAllStatus502() int64`

GetAllStatus502 returns the AllStatus502 field if non-nil, zero value otherwise.

### GetAllStatus502Ok

`func (o *OriginInspectorValues) GetAllStatus502Ok() (*int64, bool)`

GetAllStatus502Ok returns a tuple with the AllStatus502 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllStatus502

`func (o *OriginInspectorValues) SetAllStatus502(v int64)`

SetAllStatus502 sets AllStatus502 field to given value.

### HasAllStatus502

`func (o *OriginInspectorValues) HasAllStatus502() bool`

HasAllStatus502 returns a boolean if a field has been set.

### GetAllStatus503

`func (o *OriginInspectorValues) GetAllStatus503() int64`

GetAllStatus503 returns the AllStatus503 field if non-nil, zero value otherwise.

### GetAllStatus503Ok

`func (o *OriginInspectorValues) GetAllStatus503Ok() (*int64, bool)`

GetAllStatus503Ok returns a tuple with the AllStatus503 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllStatus503

`func (o *OriginInspectorValues) SetAllStatus503(v int64)`

SetAllStatus503 sets AllStatus503 field to given value.

### HasAllStatus503

`func (o *OriginInspectorValues) HasAllStatus503() bool`

HasAllStatus503 returns a boolean if a field has been set.

### GetAllStatus504

`func (o *OriginInspectorValues) GetAllStatus504() int64`

GetAllStatus504 returns the AllStatus504 field if non-nil, zero value otherwise.

### GetAllStatus504Ok

`func (o *OriginInspectorValues) GetAllStatus504Ok() (*int64, bool)`

GetAllStatus504Ok returns a tuple with the AllStatus504 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllStatus504

`func (o *OriginInspectorValues) SetAllStatus504(v int64)`

SetAllStatus504 sets AllStatus504 field to given value.

### HasAllStatus504

`func (o *OriginInspectorValues) HasAllStatus504() bool`

HasAllStatus504 returns a boolean if a field has been set.

### GetAllStatus505

`func (o *OriginInspectorValues) GetAllStatus505() int64`

GetAllStatus505 returns the AllStatus505 field if non-nil, zero value otherwise.

### GetAllStatus505Ok

`func (o *OriginInspectorValues) GetAllStatus505Ok() (*int64, bool)`

GetAllStatus505Ok returns a tuple with the AllStatus505 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllStatus505

`func (o *OriginInspectorValues) SetAllStatus505(v int64)`

SetAllStatus505 sets AllStatus505 field to given value.

### HasAllStatus505

`func (o *OriginInspectorValues) HasAllStatus505() bool`

HasAllStatus505 returns a boolean if a field has been set.

### GetAllStatus530

`func (o *OriginInspectorValues) GetAllStatus530() int64`

GetAllStatus530 returns the AllStatus530 field if non-nil, zero value otherwise.

### GetAllStatus530Ok

`func (o *OriginInspectorValues) GetAllStatus530Ok() (*int64, bool)`

GetAllStatus530Ok returns a tuple with the AllStatus530 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllStatus530

`func (o *OriginInspectorValues) SetAllStatus530(v int64)`

SetAllStatus530 sets AllStatus530 field to given value.

### HasAllStatus530

`func (o *OriginInspectorValues) HasAllStatus530() bool`

HasAllStatus530 returns a boolean if a field has been set.

### GetAllLatency0To1ms

`func (o *OriginInspectorValues) GetAllLatency0To1ms() int64`

GetAllLatency0To1ms returns the AllLatency0To1ms field if non-nil, zero value otherwise.

### GetAllLatency0To1msOk

`func (o *OriginInspectorValues) GetAllLatency0To1msOk() (*int64, bool)`

GetAllLatency0To1msOk returns a tuple with the AllLatency0To1ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllLatency0To1ms

`func (o *OriginInspectorValues) SetAllLatency0To1ms(v int64)`

SetAllLatency0To1ms sets AllLatency0To1ms field to given value.

### HasAllLatency0To1ms

`func (o *OriginInspectorValues) HasAllLatency0To1ms() bool`

HasAllLatency0To1ms returns a boolean if a field has been set.

### GetAllLatency1To5ms

`func (o *OriginInspectorValues) GetAllLatency1To5ms() int64`

GetAllLatency1To5ms returns the AllLatency1To5ms field if non-nil, zero value otherwise.

### GetAllLatency1To5msOk

`func (o *OriginInspectorValues) GetAllLatency1To5msOk() (*int64, bool)`

GetAllLatency1To5msOk returns a tuple with the AllLatency1To5ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllLatency1To5ms

`func (o *OriginInspectorValues) SetAllLatency1To5ms(v int64)`

SetAllLatency1To5ms sets AllLatency1To5ms field to given value.

### HasAllLatency1To5ms

`func (o *OriginInspectorValues) HasAllLatency1To5ms() bool`

HasAllLatency1To5ms returns a boolean if a field has been set.

### GetAllLatency5To10ms

`func (o *OriginInspectorValues) GetAllLatency5To10ms() int64`

GetAllLatency5To10ms returns the AllLatency5To10ms field if non-nil, zero value otherwise.

### GetAllLatency5To10msOk

`func (o *OriginInspectorValues) GetAllLatency5To10msOk() (*int64, bool)`

GetAllLatency5To10msOk returns a tuple with the AllLatency5To10ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllLatency5To10ms

`func (o *OriginInspectorValues) SetAllLatency5To10ms(v int64)`

SetAllLatency5To10ms sets AllLatency5To10ms field to given value.

### HasAllLatency5To10ms

`func (o *OriginInspectorValues) HasAllLatency5To10ms() bool`

HasAllLatency5To10ms returns a boolean if a field has been set.

### GetAllLatency10To50ms

`func (o *OriginInspectorValues) GetAllLatency10To50ms() int64`

GetAllLatency10To50ms returns the AllLatency10To50ms field if non-nil, zero value otherwise.

### GetAllLatency10To50msOk

`func (o *OriginInspectorValues) GetAllLatency10To50msOk() (*int64, bool)`

GetAllLatency10To50msOk returns a tuple with the AllLatency10To50ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllLatency10To50ms

`func (o *OriginInspectorValues) SetAllLatency10To50ms(v int64)`

SetAllLatency10To50ms sets AllLatency10To50ms field to given value.

### HasAllLatency10To50ms

`func (o *OriginInspectorValues) HasAllLatency10To50ms() bool`

HasAllLatency10To50ms returns a boolean if a field has been set.

### GetAllLatency50To100ms

`func (o *OriginInspectorValues) GetAllLatency50To100ms() int64`

GetAllLatency50To100ms returns the AllLatency50To100ms field if non-nil, zero value otherwise.

### GetAllLatency50To100msOk

`func (o *OriginInspectorValues) GetAllLatency50To100msOk() (*int64, bool)`

GetAllLatency50To100msOk returns a tuple with the AllLatency50To100ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllLatency50To100ms

`func (o *OriginInspectorValues) SetAllLatency50To100ms(v int64)`

SetAllLatency50To100ms sets AllLatency50To100ms field to given value.

### HasAllLatency50To100ms

`func (o *OriginInspectorValues) HasAllLatency50To100ms() bool`

HasAllLatency50To100ms returns a boolean if a field has been set.

### GetAllLatency100To250ms

`func (o *OriginInspectorValues) GetAllLatency100To250ms() int64`

GetAllLatency100To250ms returns the AllLatency100To250ms field if non-nil, zero value otherwise.

### GetAllLatency100To250msOk

`func (o *OriginInspectorValues) GetAllLatency100To250msOk() (*int64, bool)`

GetAllLatency100To250msOk returns a tuple with the AllLatency100To250ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllLatency100To250ms

`func (o *OriginInspectorValues) SetAllLatency100To250ms(v int64)`

SetAllLatency100To250ms sets AllLatency100To250ms field to given value.

### HasAllLatency100To250ms

`func (o *OriginInspectorValues) HasAllLatency100To250ms() bool`

HasAllLatency100To250ms returns a boolean if a field has been set.

### GetAllLatency250To500ms

`func (o *OriginInspectorValues) GetAllLatency250To500ms() int64`

GetAllLatency250To500ms returns the AllLatency250To500ms field if non-nil, zero value otherwise.

### GetAllLatency250To500msOk

`func (o *OriginInspectorValues) GetAllLatency250To500msOk() (*int64, bool)`

GetAllLatency250To500msOk returns a tuple with the AllLatency250To500ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllLatency250To500ms

`func (o *OriginInspectorValues) SetAllLatency250To500ms(v int64)`

SetAllLatency250To500ms sets AllLatency250To500ms field to given value.

### HasAllLatency250To500ms

`func (o *OriginInspectorValues) HasAllLatency250To500ms() bool`

HasAllLatency250To500ms returns a boolean if a field has been set.

### GetAllLatency500To1000ms

`func (o *OriginInspectorValues) GetAllLatency500To1000ms() int64`

GetAllLatency500To1000ms returns the AllLatency500To1000ms field if non-nil, zero value otherwise.

### GetAllLatency500To1000msOk

`func (o *OriginInspectorValues) GetAllLatency500To1000msOk() (*int64, bool)`

GetAllLatency500To1000msOk returns a tuple with the AllLatency500To1000ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllLatency500To1000ms

`func (o *OriginInspectorValues) SetAllLatency500To1000ms(v int64)`

SetAllLatency500To1000ms sets AllLatency500To1000ms field to given value.

### HasAllLatency500To1000ms

`func (o *OriginInspectorValues) HasAllLatency500To1000ms() bool`

HasAllLatency500To1000ms returns a boolean if a field has been set.

### GetAllLatency1000To5000ms

`func (o *OriginInspectorValues) GetAllLatency1000To5000ms() int64`

GetAllLatency1000To5000ms returns the AllLatency1000To5000ms field if non-nil, zero value otherwise.

### GetAllLatency1000To5000msOk

`func (o *OriginInspectorValues) GetAllLatency1000To5000msOk() (*int64, bool)`

GetAllLatency1000To5000msOk returns a tuple with the AllLatency1000To5000ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllLatency1000To5000ms

`func (o *OriginInspectorValues) SetAllLatency1000To5000ms(v int64)`

SetAllLatency1000To5000ms sets AllLatency1000To5000ms field to given value.

### HasAllLatency1000To5000ms

`func (o *OriginInspectorValues) HasAllLatency1000To5000ms() bool`

HasAllLatency1000To5000ms returns a boolean if a field has been set.

### GetAllLatency5000To10000ms

`func (o *OriginInspectorValues) GetAllLatency5000To10000ms() int64`

GetAllLatency5000To10000ms returns the AllLatency5000To10000ms field if non-nil, zero value otherwise.

### GetAllLatency5000To10000msOk

`func (o *OriginInspectorValues) GetAllLatency5000To10000msOk() (*int64, bool)`

GetAllLatency5000To10000msOk returns a tuple with the AllLatency5000To10000ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllLatency5000To10000ms

`func (o *OriginInspectorValues) SetAllLatency5000To10000ms(v int64)`

SetAllLatency5000To10000ms sets AllLatency5000To10000ms field to given value.

### HasAllLatency5000To10000ms

`func (o *OriginInspectorValues) HasAllLatency5000To10000ms() bool`

HasAllLatency5000To10000ms returns a boolean if a field has been set.

### GetAllLatency10000To60000ms

`func (o *OriginInspectorValues) GetAllLatency10000To60000ms() int64`

GetAllLatency10000To60000ms returns the AllLatency10000To60000ms field if non-nil, zero value otherwise.

### GetAllLatency10000To60000msOk

`func (o *OriginInspectorValues) GetAllLatency10000To60000msOk() (*int64, bool)`

GetAllLatency10000To60000msOk returns a tuple with the AllLatency10000To60000ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllLatency10000To60000ms

`func (o *OriginInspectorValues) SetAllLatency10000To60000ms(v int64)`

SetAllLatency10000To60000ms sets AllLatency10000To60000ms field to given value.

### HasAllLatency10000To60000ms

`func (o *OriginInspectorValues) HasAllLatency10000To60000ms() bool`

HasAllLatency10000To60000ms returns a boolean if a field has been set.

### GetAllLatency60000ms

`func (o *OriginInspectorValues) GetAllLatency60000ms() int64`

GetAllLatency60000ms returns the AllLatency60000ms field if non-nil, zero value otherwise.

### GetAllLatency60000msOk

`func (o *OriginInspectorValues) GetAllLatency60000msOk() (*int64, bool)`

GetAllLatency60000msOk returns a tuple with the AllLatency60000ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllLatency60000ms

`func (o *OriginInspectorValues) SetAllLatency60000ms(v int64)`

SetAllLatency60000ms sets AllLatency60000ms field to given value.

### HasAllLatency60000ms

`func (o *OriginInspectorValues) HasAllLatency60000ms() bool`

HasAllLatency60000ms returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
