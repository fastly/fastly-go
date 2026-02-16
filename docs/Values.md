# Values

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EdgeRequests** | Pointer to **int64** | Number of requests sent by end users to Fastly. | [optional] 
**EdgeRespHeaderBytes** | Pointer to **int64** | Total header bytes delivered from Fastly to the end user. | [optional] 
**EdgeRespBodyBytes** | Pointer to **int64** | Total body bytes delivered from Fastly to the end user. | [optional] 
**Status1xx** | Pointer to **int64** | Number of 1xx \&quot;Informational\&quot; category status codes delivered. | [optional] 
**Status2xx** | Pointer to **int64** | Number of 2xx \&quot;Success\&quot; status codes delivered. | [optional] 
**Status3xx** | Pointer to **int64** | Number of 3xx \&quot;Redirection\&quot; codes delivered. | [optional] 
**Status4xx** | Pointer to **int64** | Number of 4xx \&quot;Client Error\&quot; codes delivered. | [optional] 
**Status5xx** | Pointer to **int64** | Number of 5xx \&quot;Server Error\&quot; codes delivered. | [optional] 
**Status200** | Pointer to **int64** | Number of responses delivered with status code 200 (Success). | [optional] 
**Status204** | Pointer to **int64** | Number of responses delivered with status code 204 (No Content). | [optional] 
**Status206** | Pointer to **int64** | Number of responses delivered with status code 206 (Partial Content). | [optional] 
**Status301** | Pointer to **int64** | Number of responses delivered with status code 301 (Moved Permanently). | [optional] 
**Status302** | Pointer to **int64** | Number of responses delivered with status code 302 (Found). | [optional] 
**Status304** | Pointer to **int64** | Number of responses delivered with status code 304 (Not Modified). | [optional] 
**Status400** | Pointer to **int64** | Number of responses delivered with status code 400 (Bad Request). | [optional] 
**Status401** | Pointer to **int64** | Number of responses delivered with status code 401 (Unauthorized). | [optional] 
**Status403** | Pointer to **int64** | Number of responses delivered with status code 403 (Forbidden). | [optional] 
**Status404** | Pointer to **int64** | Number of responses delivered with status code 404 (Not Found). | [optional] 
**Status416** | Pointer to **int64** | Number of responses delivered with status code 416 (Range Not Satisfiable). | [optional] 
**Status429** | Pointer to **int64** | Number of responses delivered with status code 429 (Too Many Requests). | [optional] 
**Status500** | Pointer to **int64** | Number of responses delivered with status code 500 (Internal Server Error). | [optional] 
**Status501** | Pointer to **int64** | Number of responses delivered with status code 501 (Not Implemented). | [optional] 
**Status502** | Pointer to **int64** | Number of responses delivered with status code 502 (Bad Gateway). | [optional] 
**Status503** | Pointer to **int64** | Number of responses delivered with status code 503 (Service Unavailable). | [optional] 
**Status504** | Pointer to **int64** | Number of responses delivered with status code 504 (Gateway Timeout). | [optional] 
**Status505** | Pointer to **int64** | Number of responses delivered with status code 505 (HTTP Version Not Supported). | [optional] 
**Status530** | Pointer to **int64** | Number of responses delivered with status code 530. | [optional] 
**Requests** | Pointer to **int64** | Number of requests processed. | [optional] 
**RespHeaderBytes** | Pointer to **int64** | Total header bytes delivered. | [optional] 
**RespBodyBytes** | Pointer to **int64** | Total body bytes delivered. | [optional] 
**BereqHeaderBytes** | Pointer to **int64** | Total header bytes sent to origin. | [optional] 
**BereqBodyBytes** | Pointer to **int64** | Total body bytes sent to origin. | [optional] 
**EdgeHitRequests** | Pointer to **int64** | Number of requests sent by end users to Fastly that resulted in a hit at the edge. | [optional] 
**EdgeMissRequests** | Pointer to **int64** | Number of requests sent by end users to Fastly that resulted in a miss at the edge. | [optional] 
**OriginFetches** | Pointer to **int64** | Number of requests sent to origin. | [optional] 
**OriginFetchRespHeaderBytes** | Pointer to **int64** | Total header bytes received from origin. | [optional] 
**OriginFetchRespBodyBytes** | Pointer to **int64** | Total body bytes received from origin. | [optional] 
**Bandwidth** | Pointer to **int64** | Total bytes delivered (`resp_header_bytes` + `resp_body_bytes` + `bereq_header_bytes` + `bereq_body_bytes`). | [optional] 
**EdgeHitRatio** | Pointer to **float32** | Ratio of cache hits to cache misses at the edge, between 0 and 1 (`edge_hit_requests` / (`edge_hit_requests` + `edge_miss_requests`)). | [optional] 
**OriginOffload** | Pointer to **float32** | Origin Offload measures the ratio of bytes served to end users that were cached by Fastly, over the bytes served to end users, between 0 and 1. ((`edge_resp_body_bytes` + `edge_resp_header_bytes`) - (`origin_fetch_resp_body_bytes` + `origin_fetch_resp_header_bytes`)) / (`edge_resp_body_bytes` + `edge_resp_header_bytes`). Previously, Origin Offload used a different formula. [Learn more](https://www.fastly.com/documentation/reference/changes/2024/06/add-origin_offload-metric). | [optional] 
**OriginStatus200** | Pointer to **int64** | Number of responses received from origin with status code 200 (Success). | [optional] 
**OriginStatus204** | Pointer to **int64** | Number of responses received from origin with status code 204 (No Content). | [optional] 
**OriginStatus206** | Pointer to **int64** | Number of responses received from origin with status code 206 (Partial Content). | [optional] 
**OriginStatus301** | Pointer to **int64** | Number of responses received from origin with status code 301 (Moved Permanently). | [optional] 
**OriginStatus302** | Pointer to **int64** | Number of responses received from origin with status code 302 (Found). | [optional] 
**OriginStatus304** | Pointer to **int64** | Number of responses received from origin with status code 304 (Not Modified). | [optional] 
**OriginStatus400** | Pointer to **int64** | Number of responses received from origin with status code 400 (Bad Request). | [optional] 
**OriginStatus401** | Pointer to **int64** | Number of responses received from origin with status code 401 (Unauthorized). | [optional] 
**OriginStatus403** | Pointer to **int64** | Number of responses received from origin with status code 403 (Forbidden). | [optional] 
**OriginStatus404** | Pointer to **int64** | Number of responses received from origin with status code 404 (Not Found). | [optional] 
**OriginStatus416** | Pointer to **int64** | Number of responses received from origin with status code 416 (Range Not Satisfiable). | [optional] 
**OriginStatus429** | Pointer to **int64** | Number of responses received from origin with status code 429 (Too Many Requests). | [optional] 
**OriginStatus500** | Pointer to **int64** | Number of responses received from origin with status code 500 (Internal Server Error). | [optional] 
**OriginStatus501** | Pointer to **int64** | Number of responses received from origin with status code 501 (Not Implemented). | [optional] 
**OriginStatus502** | Pointer to **int64** | Number of responses received from origin with status code 502 (Bad Gateway). | [optional] 
**OriginStatus503** | Pointer to **int64** | Number of responses received from origin with status code 503 (Service Unavailable). | [optional] 
**OriginStatus504** | Pointer to **int64** | Number of responses received from origin with status code 504 (Gateway Timeout). | [optional] 
**OriginStatus505** | Pointer to **int64** | Number of responses received from origin with status code 505 (HTTP Version Not Supported). | [optional] 
**OriginStatus530** | Pointer to **int64** | Number of responses received from origin with status code 530. | [optional] 
**OriginStatus1xx** | Pointer to **int64** | Number of \&quot;Informational\&quot; category status codes received from origin. | [optional] 
**OriginStatus2xx** | Pointer to **int64** | Number of \&quot;Success\&quot; status codes received from origin. | [optional] 
**OriginStatus3xx** | Pointer to **int64** | Number of \&quot;Redirection\&quot; codes received from origin. | [optional] 
**OriginStatus4xx** | Pointer to **int64** | Number of \&quot;Client Error\&quot; codes received from origin. | [optional] 
**OriginStatus5xx** | Pointer to **int64** | Number of \&quot;Server Error\&quot; codes received from origin. | [optional] 
**ComputeBereqBodyBytes** | Pointer to **int64** | Total body bytes sent to backends (origins) by the Compute platform. | [optional] 
**ComputeBereqErrors** | Pointer to **int64** | Number of backend request errors, including timeouts, by the Compute platform. | [optional] 
**ComputeBereqHeaderBytes** | Pointer to **int64** | Total header bytes sent to backends (origins) by the Compute platform. | [optional] 
**ComputeBereqs** | Pointer to **int64** | Number of backend requests started by the Compute platform. | [optional] 
**ComputeBerespBodyBytes** | Pointer to **int64** | Total body bytes received from backends (origins) by the Compute platform. | [optional] 
**ComputeBerespHeaderBytes** | Pointer to **int64** | Total header bytes received from backends (origins) by the Compute platform. | [optional] 
**ComputeExecutionTimeMs** | Pointer to **int64** | The amount of active CPU time used to process your requests (in milliseconds). | [optional] 
**ComputeOriginStatus1xx** | Pointer to **int64** | Number of \&quot;Informational\&quot; category status codes received from origin by the Compute platform. | [optional] 
**ComputeOriginStatus200** | Pointer to **int64** | Number of responses received from origin with status code 200 (Success) by the Compute platform. | [optional] 
**ComputeOriginStatus204** | Pointer to **int64** | Number of responses received from origin with status code 204 (No Content) by the Compute platform. | [optional] 
**ComputeOriginStatus206** | Pointer to **int64** | Number of responses received from origin with status code 206 (Partial Content) by the Compute platform. | [optional] 
**ComputeOriginStatus2xx** | Pointer to **int64** | Number of \&quot;Success\&quot; status codes received from origin by the Compute platform. | [optional] 
**ComputeOriginStatus301** | Pointer to **int64** | Number of responses received from origin with status code 301 (Moved Permanently) by the Compute platform. | [optional] 
**ComputeOriginStatus302** | Pointer to **int64** | Number of responses received from origin with status code 302 (Found) by the Compute platform. | [optional] 
**ComputeOriginStatus304** | Pointer to **int64** | Number of responses received from origin with status code 304 (Not Modified) by the Compute platform. | [optional] 
**ComputeOriginStatus3xx** | Pointer to **int64** | Number of \&quot;Redirection\&quot; codes received from origin by the Compute platform. | [optional] 
**ComputeOriginStatus400** | Pointer to **int64** | Number of responses received from origin with status code 400 (Bad Request) by the Compute platform. | [optional] 
**ComputeOriginStatus401** | Pointer to **int64** | Number of responses received from origin with status code 401 (Unauthorized) by the Compute platform. | [optional] 
**ComputeOriginStatus403** | Pointer to **int64** | Number of responses received from origin with status code 403 (Forbidden) by the Compute platform. | [optional] 
**ComputeOriginStatus404** | Pointer to **int64** | Number of responses received from origin with status code 404 (Not Found) by the Compute platform. | [optional] 
**ComputeOriginStatus416** | Pointer to **int64** | Number of responses received from origin with status code 416 (Range Not Satisfiable) by the Compute platform. | [optional] 
**ComputeOriginStatus429** | Pointer to **int64** | Number of responses received from origin with status code 429 (Too Many Requests) by the Compute platform. | [optional] 
**ComputeOriginStatus4xx** | Pointer to **int64** | Number of \&quot;Client Error\&quot; codes received from origin by the Compute platform. | [optional] 
**ComputeOriginStatus500** | Pointer to **int64** | Number of responses received from origin with status code 500 (Internal Server Error) by the Compute platform. | [optional] 
**ComputeOriginStatus501** | Pointer to **int64** | Number of responses received from origin with status code 501 (Not Implemented) by the Compute platform. | [optional] 
**ComputeOriginStatus502** | Pointer to **int64** | Number of responses received from origin with status code 502 (Bad Gateway) by the Compute platform. | [optional] 
**ComputeOriginStatus503** | Pointer to **int64** | Number of responses received from origin with status code 503 (Service Unavailable) by the Compute platform. | [optional] 
**ComputeOriginStatus504** | Pointer to **int64** | Number of responses received from origin with status code 504 (Gateway Timeout) by the Compute platform. | [optional] 
**ComputeOriginStatus505** | Pointer to **int64** | Number of responses received from origin with status code 505 (HTTP Version Not Supported) by the Compute platform. | [optional] 
**ComputeOriginStatus530** | Pointer to **int64** | Number of responses received from origin with status code 530 by the Compute platform. | [optional] 
**ComputeOriginStatus5xx** | Pointer to **int64** | Number of \&quot;Server Error\&quot; codes received from origin by the Compute platform. | [optional] 
**ComputeReqBodyBytes** | Pointer to **int64** | Total body bytes received by the Compute platform. | [optional] 
**ComputeReqHeaderBytes** | Pointer to **int64** | Total header bytes received by the Compute platform. | [optional] 
**ComputeRequestTimeBilledMs** | Pointer to **int64** | The total amount of request processing time you will be billed for, measured in 50 millisecond increments. | [optional] 
**ComputeRequestTimeMs** | Pointer to **int64** | The total amount of time used to process your requests, including active CPU time (in milliseconds). | [optional] 
**ComputeRequest** | Pointer to **int64** | The total number of requests that were received by the Compute platform. | [optional] 
**ComputeRespBodyBytes** | Pointer to **int64** | Total body bytes sent from Compute to the end user. | [optional] 
**ComputeRespHeaderBytes** | Pointer to **int64** | Total header bytes sent from Compute to the end user. | [optional] 
**ComputeRespStatus103** | Pointer to **int64** | Number of responses delivered with status code 103 (Early Hints) by the Compute platform. | [optional] 
**ComputeRespStatus1xx** | Pointer to **int64** | Number of 1xx \&quot;Informational\&quot; category status codes delivered by the Compute platform. | [optional] 
**ComputeRespStatus200** | Pointer to **int64** | Number of responses delivered with status code 200 (Success) by the Compute platform. | [optional] 
**ComputeRespStatus204** | Pointer to **int64** | Number of responses delivered with status code 204 (No Content) by the Compute platform. | [optional] 
**ComputeRespStatus206** | Pointer to **int64** | Number of responses delivered with status code 206 (Partial Content) by the Compute platform. | [optional] 
**ComputeRespStatus2xx** | Pointer to **int64** | Number of 2xx \&quot;Success\&quot; status codes delivered by the Compute platform. | [optional] 
**ComputeRespStatus301** | Pointer to **int64** | Number of responses delivered with status code 301 (Moved Permanently) by the Compute platform. | [optional] 
**ComputeRespStatus302** | Pointer to **int64** | Number of responses delivered with status code 302 (Found) by the Compute platform. | [optional] 
**ComputeRespStatus304** | Pointer to **int64** | Number of responses delivered with status code 304 (Not Modified) by the Compute platform. | [optional] 
**ComputeRespStatus3xx** | Pointer to **int64** | Number of 3xx \&quot;Redirection\&quot; codes delivered by the Compute platform. | [optional] 
**ComputeRespStatus400** | Pointer to **int64** | Number of responses delivered with status code 400 (Bad Request) by the Compute platform. | [optional] 
**ComputeRespStatus401** | Pointer to **int64** | Number of responses delivered with status code 401 (Unauthorized) by the Compute platform. | [optional] 
**ComputeRespStatus403** | Pointer to **int64** | Number of responses delivered with status code 403 (Forbidden) by the Compute platform. | [optional] 
**ComputeRespStatus404** | Pointer to **int64** | Number of responses delivered with status code 404 (Not Found) by the Compute platform. | [optional] 
**ComputeRespStatus416** | Pointer to **int64** | Number of responses delivered with status code 416 (Range Not Satisfiable) by the Compute platform. | [optional] 
**ComputeRespStatus429** | Pointer to **int64** | Number of responses delivered with status code 429 (Too Many Requests) by the Compute platform. | [optional] 
**ComputeRespStatus4xx** | Pointer to **int64** | Number of 4xx \&quot;Client Error\&quot; codes delivered by the Compute platform. | [optional] 
**ComputeRespStatus500** | Pointer to **int64** | Number of responses delivered with status code 500 (Internal Server Error) by the Compute platform. | [optional] 
**ComputeRespStatus501** | Pointer to **int64** | Number of responses delivered with status code 501 (Not Implemented) by the Compute platform. | [optional] 
**ComputeRespStatus502** | Pointer to **int64** | Number of responses delivered with status code 502 (Bad Gateway) by the Compute platform. | [optional] 
**ComputeRespStatus503** | Pointer to **int64** | Number of responses delivered with status code 503 (Service Unavailable) by the Compute platform. | [optional] 
**ComputeRespStatus504** | Pointer to **int64** | Number of responses delivered with status code 504 (Gateway Timeout) by the Compute platform. | [optional] 
**ComputeRespStatus505** | Pointer to **int64** | Number of responses delivered with status code 505 (HTTP Version Not Supported) by the Compute platform. | [optional] 
**ComputeRespStatus530** | Pointer to **int64** | Number of responses delivered with status code 530 by the Compute platform. | [optional] 
**ComputeRespStatus5xx** | Pointer to **int64** | Number of \&quot;Server Error\&quot; category status codes delivered by the Compute platform. | [optional] 

## Methods

### NewValues

`func NewValues() *Values`

NewValues instantiates a new Values object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValuesWithDefaults

`func NewValuesWithDefaults() *Values`

NewValuesWithDefaults instantiates a new Values object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEdgeRequests

`func (o *Values) GetEdgeRequests() int64`

GetEdgeRequests returns the EdgeRequests field if non-nil, zero value otherwise.

### GetEdgeRequestsOk

`func (o *Values) GetEdgeRequestsOk() (*int64, bool)`

GetEdgeRequestsOk returns a tuple with the EdgeRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeRequests

`func (o *Values) SetEdgeRequests(v int64)`

SetEdgeRequests sets EdgeRequests field to given value.

### HasEdgeRequests

`func (o *Values) HasEdgeRequests() bool`

HasEdgeRequests returns a boolean if a field has been set.

### GetEdgeRespHeaderBytes

`func (o *Values) GetEdgeRespHeaderBytes() int64`

GetEdgeRespHeaderBytes returns the EdgeRespHeaderBytes field if non-nil, zero value otherwise.

### GetEdgeRespHeaderBytesOk

`func (o *Values) GetEdgeRespHeaderBytesOk() (*int64, bool)`

GetEdgeRespHeaderBytesOk returns a tuple with the EdgeRespHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeRespHeaderBytes

`func (o *Values) SetEdgeRespHeaderBytes(v int64)`

SetEdgeRespHeaderBytes sets EdgeRespHeaderBytes field to given value.

### HasEdgeRespHeaderBytes

`func (o *Values) HasEdgeRespHeaderBytes() bool`

HasEdgeRespHeaderBytes returns a boolean if a field has been set.

### GetEdgeRespBodyBytes

`func (o *Values) GetEdgeRespBodyBytes() int64`

GetEdgeRespBodyBytes returns the EdgeRespBodyBytes field if non-nil, zero value otherwise.

### GetEdgeRespBodyBytesOk

`func (o *Values) GetEdgeRespBodyBytesOk() (*int64, bool)`

GetEdgeRespBodyBytesOk returns a tuple with the EdgeRespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeRespBodyBytes

`func (o *Values) SetEdgeRespBodyBytes(v int64)`

SetEdgeRespBodyBytes sets EdgeRespBodyBytes field to given value.

### HasEdgeRespBodyBytes

`func (o *Values) HasEdgeRespBodyBytes() bool`

HasEdgeRespBodyBytes returns a boolean if a field has been set.

### GetStatus1xx

`func (o *Values) GetStatus1xx() int64`

GetStatus1xx returns the Status1xx field if non-nil, zero value otherwise.

### GetStatus1xxOk

`func (o *Values) GetStatus1xxOk() (*int64, bool)`

GetStatus1xxOk returns a tuple with the Status1xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus1xx

`func (o *Values) SetStatus1xx(v int64)`

SetStatus1xx sets Status1xx field to given value.

### HasStatus1xx

`func (o *Values) HasStatus1xx() bool`

HasStatus1xx returns a boolean if a field has been set.

### GetStatus2xx

`func (o *Values) GetStatus2xx() int64`

GetStatus2xx returns the Status2xx field if non-nil, zero value otherwise.

### GetStatus2xxOk

`func (o *Values) GetStatus2xxOk() (*int64, bool)`

GetStatus2xxOk returns a tuple with the Status2xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus2xx

`func (o *Values) SetStatus2xx(v int64)`

SetStatus2xx sets Status2xx field to given value.

### HasStatus2xx

`func (o *Values) HasStatus2xx() bool`

HasStatus2xx returns a boolean if a field has been set.

### GetStatus3xx

`func (o *Values) GetStatus3xx() int64`

GetStatus3xx returns the Status3xx field if non-nil, zero value otherwise.

### GetStatus3xxOk

`func (o *Values) GetStatus3xxOk() (*int64, bool)`

GetStatus3xxOk returns a tuple with the Status3xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus3xx

`func (o *Values) SetStatus3xx(v int64)`

SetStatus3xx sets Status3xx field to given value.

### HasStatus3xx

`func (o *Values) HasStatus3xx() bool`

HasStatus3xx returns a boolean if a field has been set.

### GetStatus4xx

`func (o *Values) GetStatus4xx() int64`

GetStatus4xx returns the Status4xx field if non-nil, zero value otherwise.

### GetStatus4xxOk

`func (o *Values) GetStatus4xxOk() (*int64, bool)`

GetStatus4xxOk returns a tuple with the Status4xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus4xx

`func (o *Values) SetStatus4xx(v int64)`

SetStatus4xx sets Status4xx field to given value.

### HasStatus4xx

`func (o *Values) HasStatus4xx() bool`

HasStatus4xx returns a boolean if a field has been set.

### GetStatus5xx

`func (o *Values) GetStatus5xx() int64`

GetStatus5xx returns the Status5xx field if non-nil, zero value otherwise.

### GetStatus5xxOk

`func (o *Values) GetStatus5xxOk() (*int64, bool)`

GetStatus5xxOk returns a tuple with the Status5xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus5xx

`func (o *Values) SetStatus5xx(v int64)`

SetStatus5xx sets Status5xx field to given value.

### HasStatus5xx

`func (o *Values) HasStatus5xx() bool`

HasStatus5xx returns a boolean if a field has been set.

### GetStatus200

`func (o *Values) GetStatus200() int64`

GetStatus200 returns the Status200 field if non-nil, zero value otherwise.

### GetStatus200Ok

`func (o *Values) GetStatus200Ok() (*int64, bool)`

GetStatus200Ok returns a tuple with the Status200 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus200

`func (o *Values) SetStatus200(v int64)`

SetStatus200 sets Status200 field to given value.

### HasStatus200

`func (o *Values) HasStatus200() bool`

HasStatus200 returns a boolean if a field has been set.

### GetStatus204

`func (o *Values) GetStatus204() int64`

GetStatus204 returns the Status204 field if non-nil, zero value otherwise.

### GetStatus204Ok

`func (o *Values) GetStatus204Ok() (*int64, bool)`

GetStatus204Ok returns a tuple with the Status204 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus204

`func (o *Values) SetStatus204(v int64)`

SetStatus204 sets Status204 field to given value.

### HasStatus204

`func (o *Values) HasStatus204() bool`

HasStatus204 returns a boolean if a field has been set.

### GetStatus206

`func (o *Values) GetStatus206() int64`

GetStatus206 returns the Status206 field if non-nil, zero value otherwise.

### GetStatus206Ok

`func (o *Values) GetStatus206Ok() (*int64, bool)`

GetStatus206Ok returns a tuple with the Status206 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus206

`func (o *Values) SetStatus206(v int64)`

SetStatus206 sets Status206 field to given value.

### HasStatus206

`func (o *Values) HasStatus206() bool`

HasStatus206 returns a boolean if a field has been set.

### GetStatus301

`func (o *Values) GetStatus301() int64`

GetStatus301 returns the Status301 field if non-nil, zero value otherwise.

### GetStatus301Ok

`func (o *Values) GetStatus301Ok() (*int64, bool)`

GetStatus301Ok returns a tuple with the Status301 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus301

`func (o *Values) SetStatus301(v int64)`

SetStatus301 sets Status301 field to given value.

### HasStatus301

`func (o *Values) HasStatus301() bool`

HasStatus301 returns a boolean if a field has been set.

### GetStatus302

`func (o *Values) GetStatus302() int64`

GetStatus302 returns the Status302 field if non-nil, zero value otherwise.

### GetStatus302Ok

`func (o *Values) GetStatus302Ok() (*int64, bool)`

GetStatus302Ok returns a tuple with the Status302 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus302

`func (o *Values) SetStatus302(v int64)`

SetStatus302 sets Status302 field to given value.

### HasStatus302

`func (o *Values) HasStatus302() bool`

HasStatus302 returns a boolean if a field has been set.

### GetStatus304

`func (o *Values) GetStatus304() int64`

GetStatus304 returns the Status304 field if non-nil, zero value otherwise.

### GetStatus304Ok

`func (o *Values) GetStatus304Ok() (*int64, bool)`

GetStatus304Ok returns a tuple with the Status304 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus304

`func (o *Values) SetStatus304(v int64)`

SetStatus304 sets Status304 field to given value.

### HasStatus304

`func (o *Values) HasStatus304() bool`

HasStatus304 returns a boolean if a field has been set.

### GetStatus400

`func (o *Values) GetStatus400() int64`

GetStatus400 returns the Status400 field if non-nil, zero value otherwise.

### GetStatus400Ok

`func (o *Values) GetStatus400Ok() (*int64, bool)`

GetStatus400Ok returns a tuple with the Status400 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus400

`func (o *Values) SetStatus400(v int64)`

SetStatus400 sets Status400 field to given value.

### HasStatus400

`func (o *Values) HasStatus400() bool`

HasStatus400 returns a boolean if a field has been set.

### GetStatus401

`func (o *Values) GetStatus401() int64`

GetStatus401 returns the Status401 field if non-nil, zero value otherwise.

### GetStatus401Ok

`func (o *Values) GetStatus401Ok() (*int64, bool)`

GetStatus401Ok returns a tuple with the Status401 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus401

`func (o *Values) SetStatus401(v int64)`

SetStatus401 sets Status401 field to given value.

### HasStatus401

`func (o *Values) HasStatus401() bool`

HasStatus401 returns a boolean if a field has been set.

### GetStatus403

`func (o *Values) GetStatus403() int64`

GetStatus403 returns the Status403 field if non-nil, zero value otherwise.

### GetStatus403Ok

`func (o *Values) GetStatus403Ok() (*int64, bool)`

GetStatus403Ok returns a tuple with the Status403 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus403

`func (o *Values) SetStatus403(v int64)`

SetStatus403 sets Status403 field to given value.

### HasStatus403

`func (o *Values) HasStatus403() bool`

HasStatus403 returns a boolean if a field has been set.

### GetStatus404

`func (o *Values) GetStatus404() int64`

GetStatus404 returns the Status404 field if non-nil, zero value otherwise.

### GetStatus404Ok

`func (o *Values) GetStatus404Ok() (*int64, bool)`

GetStatus404Ok returns a tuple with the Status404 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus404

`func (o *Values) SetStatus404(v int64)`

SetStatus404 sets Status404 field to given value.

### HasStatus404

`func (o *Values) HasStatus404() bool`

HasStatus404 returns a boolean if a field has been set.

### GetStatus416

`func (o *Values) GetStatus416() int64`

GetStatus416 returns the Status416 field if non-nil, zero value otherwise.

### GetStatus416Ok

`func (o *Values) GetStatus416Ok() (*int64, bool)`

GetStatus416Ok returns a tuple with the Status416 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus416

`func (o *Values) SetStatus416(v int64)`

SetStatus416 sets Status416 field to given value.

### HasStatus416

`func (o *Values) HasStatus416() bool`

HasStatus416 returns a boolean if a field has been set.

### GetStatus429

`func (o *Values) GetStatus429() int64`

GetStatus429 returns the Status429 field if non-nil, zero value otherwise.

### GetStatus429Ok

`func (o *Values) GetStatus429Ok() (*int64, bool)`

GetStatus429Ok returns a tuple with the Status429 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus429

`func (o *Values) SetStatus429(v int64)`

SetStatus429 sets Status429 field to given value.

### HasStatus429

`func (o *Values) HasStatus429() bool`

HasStatus429 returns a boolean if a field has been set.

### GetStatus500

`func (o *Values) GetStatus500() int64`

GetStatus500 returns the Status500 field if non-nil, zero value otherwise.

### GetStatus500Ok

`func (o *Values) GetStatus500Ok() (*int64, bool)`

GetStatus500Ok returns a tuple with the Status500 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus500

`func (o *Values) SetStatus500(v int64)`

SetStatus500 sets Status500 field to given value.

### HasStatus500

`func (o *Values) HasStatus500() bool`

HasStatus500 returns a boolean if a field has been set.

### GetStatus501

`func (o *Values) GetStatus501() int64`

GetStatus501 returns the Status501 field if non-nil, zero value otherwise.

### GetStatus501Ok

`func (o *Values) GetStatus501Ok() (*int64, bool)`

GetStatus501Ok returns a tuple with the Status501 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus501

`func (o *Values) SetStatus501(v int64)`

SetStatus501 sets Status501 field to given value.

### HasStatus501

`func (o *Values) HasStatus501() bool`

HasStatus501 returns a boolean if a field has been set.

### GetStatus502

`func (o *Values) GetStatus502() int64`

GetStatus502 returns the Status502 field if non-nil, zero value otherwise.

### GetStatus502Ok

`func (o *Values) GetStatus502Ok() (*int64, bool)`

GetStatus502Ok returns a tuple with the Status502 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus502

`func (o *Values) SetStatus502(v int64)`

SetStatus502 sets Status502 field to given value.

### HasStatus502

`func (o *Values) HasStatus502() bool`

HasStatus502 returns a boolean if a field has been set.

### GetStatus503

`func (o *Values) GetStatus503() int64`

GetStatus503 returns the Status503 field if non-nil, zero value otherwise.

### GetStatus503Ok

`func (o *Values) GetStatus503Ok() (*int64, bool)`

GetStatus503Ok returns a tuple with the Status503 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus503

`func (o *Values) SetStatus503(v int64)`

SetStatus503 sets Status503 field to given value.

### HasStatus503

`func (o *Values) HasStatus503() bool`

HasStatus503 returns a boolean if a field has been set.

### GetStatus504

`func (o *Values) GetStatus504() int64`

GetStatus504 returns the Status504 field if non-nil, zero value otherwise.

### GetStatus504Ok

`func (o *Values) GetStatus504Ok() (*int64, bool)`

GetStatus504Ok returns a tuple with the Status504 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus504

`func (o *Values) SetStatus504(v int64)`

SetStatus504 sets Status504 field to given value.

### HasStatus504

`func (o *Values) HasStatus504() bool`

HasStatus504 returns a boolean if a field has been set.

### GetStatus505

`func (o *Values) GetStatus505() int64`

GetStatus505 returns the Status505 field if non-nil, zero value otherwise.

### GetStatus505Ok

`func (o *Values) GetStatus505Ok() (*int64, bool)`

GetStatus505Ok returns a tuple with the Status505 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus505

`func (o *Values) SetStatus505(v int64)`

SetStatus505 sets Status505 field to given value.

### HasStatus505

`func (o *Values) HasStatus505() bool`

HasStatus505 returns a boolean if a field has been set.

### GetStatus530

`func (o *Values) GetStatus530() int64`

GetStatus530 returns the Status530 field if non-nil, zero value otherwise.

### GetStatus530Ok

`func (o *Values) GetStatus530Ok() (*int64, bool)`

GetStatus530Ok returns a tuple with the Status530 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus530

`func (o *Values) SetStatus530(v int64)`

SetStatus530 sets Status530 field to given value.

### HasStatus530

`func (o *Values) HasStatus530() bool`

HasStatus530 returns a boolean if a field has been set.

### GetRequests

`func (o *Values) GetRequests() int64`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *Values) GetRequestsOk() (*int64, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *Values) SetRequests(v int64)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *Values) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### GetRespHeaderBytes

`func (o *Values) GetRespHeaderBytes() int64`

GetRespHeaderBytes returns the RespHeaderBytes field if non-nil, zero value otherwise.

### GetRespHeaderBytesOk

`func (o *Values) GetRespHeaderBytesOk() (*int64, bool)`

GetRespHeaderBytesOk returns a tuple with the RespHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRespHeaderBytes

`func (o *Values) SetRespHeaderBytes(v int64)`

SetRespHeaderBytes sets RespHeaderBytes field to given value.

### HasRespHeaderBytes

`func (o *Values) HasRespHeaderBytes() bool`

HasRespHeaderBytes returns a boolean if a field has been set.

### GetRespBodyBytes

`func (o *Values) GetRespBodyBytes() int64`

GetRespBodyBytes returns the RespBodyBytes field if non-nil, zero value otherwise.

### GetRespBodyBytesOk

`func (o *Values) GetRespBodyBytesOk() (*int64, bool)`

GetRespBodyBytesOk returns a tuple with the RespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRespBodyBytes

`func (o *Values) SetRespBodyBytes(v int64)`

SetRespBodyBytes sets RespBodyBytes field to given value.

### HasRespBodyBytes

`func (o *Values) HasRespBodyBytes() bool`

HasRespBodyBytes returns a boolean if a field has been set.

### GetBereqHeaderBytes

`func (o *Values) GetBereqHeaderBytes() int64`

GetBereqHeaderBytes returns the BereqHeaderBytes field if non-nil, zero value otherwise.

### GetBereqHeaderBytesOk

`func (o *Values) GetBereqHeaderBytesOk() (*int64, bool)`

GetBereqHeaderBytesOk returns a tuple with the BereqHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBereqHeaderBytes

`func (o *Values) SetBereqHeaderBytes(v int64)`

SetBereqHeaderBytes sets BereqHeaderBytes field to given value.

### HasBereqHeaderBytes

`func (o *Values) HasBereqHeaderBytes() bool`

HasBereqHeaderBytes returns a boolean if a field has been set.

### GetBereqBodyBytes

`func (o *Values) GetBereqBodyBytes() int64`

GetBereqBodyBytes returns the BereqBodyBytes field if non-nil, zero value otherwise.

### GetBereqBodyBytesOk

`func (o *Values) GetBereqBodyBytesOk() (*int64, bool)`

GetBereqBodyBytesOk returns a tuple with the BereqBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBereqBodyBytes

`func (o *Values) SetBereqBodyBytes(v int64)`

SetBereqBodyBytes sets BereqBodyBytes field to given value.

### HasBereqBodyBytes

`func (o *Values) HasBereqBodyBytes() bool`

HasBereqBodyBytes returns a boolean if a field has been set.

### GetEdgeHitRequests

`func (o *Values) GetEdgeHitRequests() int64`

GetEdgeHitRequests returns the EdgeHitRequests field if non-nil, zero value otherwise.

### GetEdgeHitRequestsOk

`func (o *Values) GetEdgeHitRequestsOk() (*int64, bool)`

GetEdgeHitRequestsOk returns a tuple with the EdgeHitRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeHitRequests

`func (o *Values) SetEdgeHitRequests(v int64)`

SetEdgeHitRequests sets EdgeHitRequests field to given value.

### HasEdgeHitRequests

`func (o *Values) HasEdgeHitRequests() bool`

HasEdgeHitRequests returns a boolean if a field has been set.

### GetEdgeMissRequests

`func (o *Values) GetEdgeMissRequests() int64`

GetEdgeMissRequests returns the EdgeMissRequests field if non-nil, zero value otherwise.

### GetEdgeMissRequestsOk

`func (o *Values) GetEdgeMissRequestsOk() (*int64, bool)`

GetEdgeMissRequestsOk returns a tuple with the EdgeMissRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeMissRequests

`func (o *Values) SetEdgeMissRequests(v int64)`

SetEdgeMissRequests sets EdgeMissRequests field to given value.

### HasEdgeMissRequests

`func (o *Values) HasEdgeMissRequests() bool`

HasEdgeMissRequests returns a boolean if a field has been set.

### GetOriginFetches

`func (o *Values) GetOriginFetches() int64`

GetOriginFetches returns the OriginFetches field if non-nil, zero value otherwise.

### GetOriginFetchesOk

`func (o *Values) GetOriginFetchesOk() (*int64, bool)`

GetOriginFetchesOk returns a tuple with the OriginFetches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginFetches

`func (o *Values) SetOriginFetches(v int64)`

SetOriginFetches sets OriginFetches field to given value.

### HasOriginFetches

`func (o *Values) HasOriginFetches() bool`

HasOriginFetches returns a boolean if a field has been set.

### GetOriginFetchRespHeaderBytes

`func (o *Values) GetOriginFetchRespHeaderBytes() int64`

GetOriginFetchRespHeaderBytes returns the OriginFetchRespHeaderBytes field if non-nil, zero value otherwise.

### GetOriginFetchRespHeaderBytesOk

`func (o *Values) GetOriginFetchRespHeaderBytesOk() (*int64, bool)`

GetOriginFetchRespHeaderBytesOk returns a tuple with the OriginFetchRespHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginFetchRespHeaderBytes

`func (o *Values) SetOriginFetchRespHeaderBytes(v int64)`

SetOriginFetchRespHeaderBytes sets OriginFetchRespHeaderBytes field to given value.

### HasOriginFetchRespHeaderBytes

`func (o *Values) HasOriginFetchRespHeaderBytes() bool`

HasOriginFetchRespHeaderBytes returns a boolean if a field has been set.

### GetOriginFetchRespBodyBytes

`func (o *Values) GetOriginFetchRespBodyBytes() int64`

GetOriginFetchRespBodyBytes returns the OriginFetchRespBodyBytes field if non-nil, zero value otherwise.

### GetOriginFetchRespBodyBytesOk

`func (o *Values) GetOriginFetchRespBodyBytesOk() (*int64, bool)`

GetOriginFetchRespBodyBytesOk returns a tuple with the OriginFetchRespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginFetchRespBodyBytes

`func (o *Values) SetOriginFetchRespBodyBytes(v int64)`

SetOriginFetchRespBodyBytes sets OriginFetchRespBodyBytes field to given value.

### HasOriginFetchRespBodyBytes

`func (o *Values) HasOriginFetchRespBodyBytes() bool`

HasOriginFetchRespBodyBytes returns a boolean if a field has been set.

### GetBandwidth

`func (o *Values) GetBandwidth() int64`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *Values) GetBandwidthOk() (*int64, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *Values) SetBandwidth(v int64)`

SetBandwidth sets Bandwidth field to given value.

### HasBandwidth

`func (o *Values) HasBandwidth() bool`

HasBandwidth returns a boolean if a field has been set.

### GetEdgeHitRatio

`func (o *Values) GetEdgeHitRatio() float32`

GetEdgeHitRatio returns the EdgeHitRatio field if non-nil, zero value otherwise.

### GetEdgeHitRatioOk

`func (o *Values) GetEdgeHitRatioOk() (*float32, bool)`

GetEdgeHitRatioOk returns a tuple with the EdgeHitRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeHitRatio

`func (o *Values) SetEdgeHitRatio(v float32)`

SetEdgeHitRatio sets EdgeHitRatio field to given value.

### HasEdgeHitRatio

`func (o *Values) HasEdgeHitRatio() bool`

HasEdgeHitRatio returns a boolean if a field has been set.

### GetOriginOffload

`func (o *Values) GetOriginOffload() float32`

GetOriginOffload returns the OriginOffload field if non-nil, zero value otherwise.

### GetOriginOffloadOk

`func (o *Values) GetOriginOffloadOk() (*float32, bool)`

GetOriginOffloadOk returns a tuple with the OriginOffload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginOffload

`func (o *Values) SetOriginOffload(v float32)`

SetOriginOffload sets OriginOffload field to given value.

### HasOriginOffload

`func (o *Values) HasOriginOffload() bool`

HasOriginOffload returns a boolean if a field has been set.

### GetOriginStatus200

`func (o *Values) GetOriginStatus200() int64`

GetOriginStatus200 returns the OriginStatus200 field if non-nil, zero value otherwise.

### GetOriginStatus200Ok

`func (o *Values) GetOriginStatus200Ok() (*int64, bool)`

GetOriginStatus200Ok returns a tuple with the OriginStatus200 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginStatus200

`func (o *Values) SetOriginStatus200(v int64)`

SetOriginStatus200 sets OriginStatus200 field to given value.

### HasOriginStatus200

`func (o *Values) HasOriginStatus200() bool`

HasOriginStatus200 returns a boolean if a field has been set.

### GetOriginStatus204

`func (o *Values) GetOriginStatus204() int64`

GetOriginStatus204 returns the OriginStatus204 field if non-nil, zero value otherwise.

### GetOriginStatus204Ok

`func (o *Values) GetOriginStatus204Ok() (*int64, bool)`

GetOriginStatus204Ok returns a tuple with the OriginStatus204 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginStatus204

`func (o *Values) SetOriginStatus204(v int64)`

SetOriginStatus204 sets OriginStatus204 field to given value.

### HasOriginStatus204

`func (o *Values) HasOriginStatus204() bool`

HasOriginStatus204 returns a boolean if a field has been set.

### GetOriginStatus206

`func (o *Values) GetOriginStatus206() int64`

GetOriginStatus206 returns the OriginStatus206 field if non-nil, zero value otherwise.

### GetOriginStatus206Ok

`func (o *Values) GetOriginStatus206Ok() (*int64, bool)`

GetOriginStatus206Ok returns a tuple with the OriginStatus206 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginStatus206

`func (o *Values) SetOriginStatus206(v int64)`

SetOriginStatus206 sets OriginStatus206 field to given value.

### HasOriginStatus206

`func (o *Values) HasOriginStatus206() bool`

HasOriginStatus206 returns a boolean if a field has been set.

### GetOriginStatus301

`func (o *Values) GetOriginStatus301() int64`

GetOriginStatus301 returns the OriginStatus301 field if non-nil, zero value otherwise.

### GetOriginStatus301Ok

`func (o *Values) GetOriginStatus301Ok() (*int64, bool)`

GetOriginStatus301Ok returns a tuple with the OriginStatus301 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginStatus301

`func (o *Values) SetOriginStatus301(v int64)`

SetOriginStatus301 sets OriginStatus301 field to given value.

### HasOriginStatus301

`func (o *Values) HasOriginStatus301() bool`

HasOriginStatus301 returns a boolean if a field has been set.

### GetOriginStatus302

`func (o *Values) GetOriginStatus302() int64`

GetOriginStatus302 returns the OriginStatus302 field if non-nil, zero value otherwise.

### GetOriginStatus302Ok

`func (o *Values) GetOriginStatus302Ok() (*int64, bool)`

GetOriginStatus302Ok returns a tuple with the OriginStatus302 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginStatus302

`func (o *Values) SetOriginStatus302(v int64)`

SetOriginStatus302 sets OriginStatus302 field to given value.

### HasOriginStatus302

`func (o *Values) HasOriginStatus302() bool`

HasOriginStatus302 returns a boolean if a field has been set.

### GetOriginStatus304

`func (o *Values) GetOriginStatus304() int64`

GetOriginStatus304 returns the OriginStatus304 field if non-nil, zero value otherwise.

### GetOriginStatus304Ok

`func (o *Values) GetOriginStatus304Ok() (*int64, bool)`

GetOriginStatus304Ok returns a tuple with the OriginStatus304 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginStatus304

`func (o *Values) SetOriginStatus304(v int64)`

SetOriginStatus304 sets OriginStatus304 field to given value.

### HasOriginStatus304

`func (o *Values) HasOriginStatus304() bool`

HasOriginStatus304 returns a boolean if a field has been set.

### GetOriginStatus400

`func (o *Values) GetOriginStatus400() int64`

GetOriginStatus400 returns the OriginStatus400 field if non-nil, zero value otherwise.

### GetOriginStatus400Ok

`func (o *Values) GetOriginStatus400Ok() (*int64, bool)`

GetOriginStatus400Ok returns a tuple with the OriginStatus400 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginStatus400

`func (o *Values) SetOriginStatus400(v int64)`

SetOriginStatus400 sets OriginStatus400 field to given value.

### HasOriginStatus400

`func (o *Values) HasOriginStatus400() bool`

HasOriginStatus400 returns a boolean if a field has been set.

### GetOriginStatus401

`func (o *Values) GetOriginStatus401() int64`

GetOriginStatus401 returns the OriginStatus401 field if non-nil, zero value otherwise.

### GetOriginStatus401Ok

`func (o *Values) GetOriginStatus401Ok() (*int64, bool)`

GetOriginStatus401Ok returns a tuple with the OriginStatus401 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginStatus401

`func (o *Values) SetOriginStatus401(v int64)`

SetOriginStatus401 sets OriginStatus401 field to given value.

### HasOriginStatus401

`func (o *Values) HasOriginStatus401() bool`

HasOriginStatus401 returns a boolean if a field has been set.

### GetOriginStatus403

`func (o *Values) GetOriginStatus403() int64`

GetOriginStatus403 returns the OriginStatus403 field if non-nil, zero value otherwise.

### GetOriginStatus403Ok

`func (o *Values) GetOriginStatus403Ok() (*int64, bool)`

GetOriginStatus403Ok returns a tuple with the OriginStatus403 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginStatus403

`func (o *Values) SetOriginStatus403(v int64)`

SetOriginStatus403 sets OriginStatus403 field to given value.

### HasOriginStatus403

`func (o *Values) HasOriginStatus403() bool`

HasOriginStatus403 returns a boolean if a field has been set.

### GetOriginStatus404

`func (o *Values) GetOriginStatus404() int64`

GetOriginStatus404 returns the OriginStatus404 field if non-nil, zero value otherwise.

### GetOriginStatus404Ok

`func (o *Values) GetOriginStatus404Ok() (*int64, bool)`

GetOriginStatus404Ok returns a tuple with the OriginStatus404 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginStatus404

`func (o *Values) SetOriginStatus404(v int64)`

SetOriginStatus404 sets OriginStatus404 field to given value.

### HasOriginStatus404

`func (o *Values) HasOriginStatus404() bool`

HasOriginStatus404 returns a boolean if a field has been set.

### GetOriginStatus416

`func (o *Values) GetOriginStatus416() int64`

GetOriginStatus416 returns the OriginStatus416 field if non-nil, zero value otherwise.

### GetOriginStatus416Ok

`func (o *Values) GetOriginStatus416Ok() (*int64, bool)`

GetOriginStatus416Ok returns a tuple with the OriginStatus416 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginStatus416

`func (o *Values) SetOriginStatus416(v int64)`

SetOriginStatus416 sets OriginStatus416 field to given value.

### HasOriginStatus416

`func (o *Values) HasOriginStatus416() bool`

HasOriginStatus416 returns a boolean if a field has been set.

### GetOriginStatus429

`func (o *Values) GetOriginStatus429() int64`

GetOriginStatus429 returns the OriginStatus429 field if non-nil, zero value otherwise.

### GetOriginStatus429Ok

`func (o *Values) GetOriginStatus429Ok() (*int64, bool)`

GetOriginStatus429Ok returns a tuple with the OriginStatus429 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginStatus429

`func (o *Values) SetOriginStatus429(v int64)`

SetOriginStatus429 sets OriginStatus429 field to given value.

### HasOriginStatus429

`func (o *Values) HasOriginStatus429() bool`

HasOriginStatus429 returns a boolean if a field has been set.

### GetOriginStatus500

`func (o *Values) GetOriginStatus500() int64`

GetOriginStatus500 returns the OriginStatus500 field if non-nil, zero value otherwise.

### GetOriginStatus500Ok

`func (o *Values) GetOriginStatus500Ok() (*int64, bool)`

GetOriginStatus500Ok returns a tuple with the OriginStatus500 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginStatus500

`func (o *Values) SetOriginStatus500(v int64)`

SetOriginStatus500 sets OriginStatus500 field to given value.

### HasOriginStatus500

`func (o *Values) HasOriginStatus500() bool`

HasOriginStatus500 returns a boolean if a field has been set.

### GetOriginStatus501

`func (o *Values) GetOriginStatus501() int64`

GetOriginStatus501 returns the OriginStatus501 field if non-nil, zero value otherwise.

### GetOriginStatus501Ok

`func (o *Values) GetOriginStatus501Ok() (*int64, bool)`

GetOriginStatus501Ok returns a tuple with the OriginStatus501 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginStatus501

`func (o *Values) SetOriginStatus501(v int64)`

SetOriginStatus501 sets OriginStatus501 field to given value.

### HasOriginStatus501

`func (o *Values) HasOriginStatus501() bool`

HasOriginStatus501 returns a boolean if a field has been set.

### GetOriginStatus502

`func (o *Values) GetOriginStatus502() int64`

GetOriginStatus502 returns the OriginStatus502 field if non-nil, zero value otherwise.

### GetOriginStatus502Ok

`func (o *Values) GetOriginStatus502Ok() (*int64, bool)`

GetOriginStatus502Ok returns a tuple with the OriginStatus502 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginStatus502

`func (o *Values) SetOriginStatus502(v int64)`

SetOriginStatus502 sets OriginStatus502 field to given value.

### HasOriginStatus502

`func (o *Values) HasOriginStatus502() bool`

HasOriginStatus502 returns a boolean if a field has been set.

### GetOriginStatus503

`func (o *Values) GetOriginStatus503() int64`

GetOriginStatus503 returns the OriginStatus503 field if non-nil, zero value otherwise.

### GetOriginStatus503Ok

`func (o *Values) GetOriginStatus503Ok() (*int64, bool)`

GetOriginStatus503Ok returns a tuple with the OriginStatus503 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginStatus503

`func (o *Values) SetOriginStatus503(v int64)`

SetOriginStatus503 sets OriginStatus503 field to given value.

### HasOriginStatus503

`func (o *Values) HasOriginStatus503() bool`

HasOriginStatus503 returns a boolean if a field has been set.

### GetOriginStatus504

`func (o *Values) GetOriginStatus504() int64`

GetOriginStatus504 returns the OriginStatus504 field if non-nil, zero value otherwise.

### GetOriginStatus504Ok

`func (o *Values) GetOriginStatus504Ok() (*int64, bool)`

GetOriginStatus504Ok returns a tuple with the OriginStatus504 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginStatus504

`func (o *Values) SetOriginStatus504(v int64)`

SetOriginStatus504 sets OriginStatus504 field to given value.

### HasOriginStatus504

`func (o *Values) HasOriginStatus504() bool`

HasOriginStatus504 returns a boolean if a field has been set.

### GetOriginStatus505

`func (o *Values) GetOriginStatus505() int64`

GetOriginStatus505 returns the OriginStatus505 field if non-nil, zero value otherwise.

### GetOriginStatus505Ok

`func (o *Values) GetOriginStatus505Ok() (*int64, bool)`

GetOriginStatus505Ok returns a tuple with the OriginStatus505 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginStatus505

`func (o *Values) SetOriginStatus505(v int64)`

SetOriginStatus505 sets OriginStatus505 field to given value.

### HasOriginStatus505

`func (o *Values) HasOriginStatus505() bool`

HasOriginStatus505 returns a boolean if a field has been set.

### GetOriginStatus530

`func (o *Values) GetOriginStatus530() int64`

GetOriginStatus530 returns the OriginStatus530 field if non-nil, zero value otherwise.

### GetOriginStatus530Ok

`func (o *Values) GetOriginStatus530Ok() (*int64, bool)`

GetOriginStatus530Ok returns a tuple with the OriginStatus530 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginStatus530

`func (o *Values) SetOriginStatus530(v int64)`

SetOriginStatus530 sets OriginStatus530 field to given value.

### HasOriginStatus530

`func (o *Values) HasOriginStatus530() bool`

HasOriginStatus530 returns a boolean if a field has been set.

### GetOriginStatus1xx

`func (o *Values) GetOriginStatus1xx() int64`

GetOriginStatus1xx returns the OriginStatus1xx field if non-nil, zero value otherwise.

### GetOriginStatus1xxOk

`func (o *Values) GetOriginStatus1xxOk() (*int64, bool)`

GetOriginStatus1xxOk returns a tuple with the OriginStatus1xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginStatus1xx

`func (o *Values) SetOriginStatus1xx(v int64)`

SetOriginStatus1xx sets OriginStatus1xx field to given value.

### HasOriginStatus1xx

`func (o *Values) HasOriginStatus1xx() bool`

HasOriginStatus1xx returns a boolean if a field has been set.

### GetOriginStatus2xx

`func (o *Values) GetOriginStatus2xx() int64`

GetOriginStatus2xx returns the OriginStatus2xx field if non-nil, zero value otherwise.

### GetOriginStatus2xxOk

`func (o *Values) GetOriginStatus2xxOk() (*int64, bool)`

GetOriginStatus2xxOk returns a tuple with the OriginStatus2xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginStatus2xx

`func (o *Values) SetOriginStatus2xx(v int64)`

SetOriginStatus2xx sets OriginStatus2xx field to given value.

### HasOriginStatus2xx

`func (o *Values) HasOriginStatus2xx() bool`

HasOriginStatus2xx returns a boolean if a field has been set.

### GetOriginStatus3xx

`func (o *Values) GetOriginStatus3xx() int64`

GetOriginStatus3xx returns the OriginStatus3xx field if non-nil, zero value otherwise.

### GetOriginStatus3xxOk

`func (o *Values) GetOriginStatus3xxOk() (*int64, bool)`

GetOriginStatus3xxOk returns a tuple with the OriginStatus3xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginStatus3xx

`func (o *Values) SetOriginStatus3xx(v int64)`

SetOriginStatus3xx sets OriginStatus3xx field to given value.

### HasOriginStatus3xx

`func (o *Values) HasOriginStatus3xx() bool`

HasOriginStatus3xx returns a boolean if a field has been set.

### GetOriginStatus4xx

`func (o *Values) GetOriginStatus4xx() int64`

GetOriginStatus4xx returns the OriginStatus4xx field if non-nil, zero value otherwise.

### GetOriginStatus4xxOk

`func (o *Values) GetOriginStatus4xxOk() (*int64, bool)`

GetOriginStatus4xxOk returns a tuple with the OriginStatus4xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginStatus4xx

`func (o *Values) SetOriginStatus4xx(v int64)`

SetOriginStatus4xx sets OriginStatus4xx field to given value.

### HasOriginStatus4xx

`func (o *Values) HasOriginStatus4xx() bool`

HasOriginStatus4xx returns a boolean if a field has been set.

### GetOriginStatus5xx

`func (o *Values) GetOriginStatus5xx() int64`

GetOriginStatus5xx returns the OriginStatus5xx field if non-nil, zero value otherwise.

### GetOriginStatus5xxOk

`func (o *Values) GetOriginStatus5xxOk() (*int64, bool)`

GetOriginStatus5xxOk returns a tuple with the OriginStatus5xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginStatus5xx

`func (o *Values) SetOriginStatus5xx(v int64)`

SetOriginStatus5xx sets OriginStatus5xx field to given value.

### HasOriginStatus5xx

`func (o *Values) HasOriginStatus5xx() bool`

HasOriginStatus5xx returns a boolean if a field has been set.

### GetComputeBereqBodyBytes

`func (o *Values) GetComputeBereqBodyBytes() int64`

GetComputeBereqBodyBytes returns the ComputeBereqBodyBytes field if non-nil, zero value otherwise.

### GetComputeBereqBodyBytesOk

`func (o *Values) GetComputeBereqBodyBytesOk() (*int64, bool)`

GetComputeBereqBodyBytesOk returns a tuple with the ComputeBereqBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBereqBodyBytes

`func (o *Values) SetComputeBereqBodyBytes(v int64)`

SetComputeBereqBodyBytes sets ComputeBereqBodyBytes field to given value.

### HasComputeBereqBodyBytes

`func (o *Values) HasComputeBereqBodyBytes() bool`

HasComputeBereqBodyBytes returns a boolean if a field has been set.

### GetComputeBereqErrors

`func (o *Values) GetComputeBereqErrors() int64`

GetComputeBereqErrors returns the ComputeBereqErrors field if non-nil, zero value otherwise.

### GetComputeBereqErrorsOk

`func (o *Values) GetComputeBereqErrorsOk() (*int64, bool)`

GetComputeBereqErrorsOk returns a tuple with the ComputeBereqErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBereqErrors

`func (o *Values) SetComputeBereqErrors(v int64)`

SetComputeBereqErrors sets ComputeBereqErrors field to given value.

### HasComputeBereqErrors

`func (o *Values) HasComputeBereqErrors() bool`

HasComputeBereqErrors returns a boolean if a field has been set.

### GetComputeBereqHeaderBytes

`func (o *Values) GetComputeBereqHeaderBytes() int64`

GetComputeBereqHeaderBytes returns the ComputeBereqHeaderBytes field if non-nil, zero value otherwise.

### GetComputeBereqHeaderBytesOk

`func (o *Values) GetComputeBereqHeaderBytesOk() (*int64, bool)`

GetComputeBereqHeaderBytesOk returns a tuple with the ComputeBereqHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBereqHeaderBytes

`func (o *Values) SetComputeBereqHeaderBytes(v int64)`

SetComputeBereqHeaderBytes sets ComputeBereqHeaderBytes field to given value.

### HasComputeBereqHeaderBytes

`func (o *Values) HasComputeBereqHeaderBytes() bool`

HasComputeBereqHeaderBytes returns a boolean if a field has been set.

### GetComputeBereqs

`func (o *Values) GetComputeBereqs() int64`

GetComputeBereqs returns the ComputeBereqs field if non-nil, zero value otherwise.

### GetComputeBereqsOk

`func (o *Values) GetComputeBereqsOk() (*int64, bool)`

GetComputeBereqsOk returns a tuple with the ComputeBereqs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBereqs

`func (o *Values) SetComputeBereqs(v int64)`

SetComputeBereqs sets ComputeBereqs field to given value.

### HasComputeBereqs

`func (o *Values) HasComputeBereqs() bool`

HasComputeBereqs returns a boolean if a field has been set.

### GetComputeBerespBodyBytes

`func (o *Values) GetComputeBerespBodyBytes() int64`

GetComputeBerespBodyBytes returns the ComputeBerespBodyBytes field if non-nil, zero value otherwise.

### GetComputeBerespBodyBytesOk

`func (o *Values) GetComputeBerespBodyBytesOk() (*int64, bool)`

GetComputeBerespBodyBytesOk returns a tuple with the ComputeBerespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBerespBodyBytes

`func (o *Values) SetComputeBerespBodyBytes(v int64)`

SetComputeBerespBodyBytes sets ComputeBerespBodyBytes field to given value.

### HasComputeBerespBodyBytes

`func (o *Values) HasComputeBerespBodyBytes() bool`

HasComputeBerespBodyBytes returns a boolean if a field has been set.

### GetComputeBerespHeaderBytes

`func (o *Values) GetComputeBerespHeaderBytes() int64`

GetComputeBerespHeaderBytes returns the ComputeBerespHeaderBytes field if non-nil, zero value otherwise.

### GetComputeBerespHeaderBytesOk

`func (o *Values) GetComputeBerespHeaderBytesOk() (*int64, bool)`

GetComputeBerespHeaderBytesOk returns a tuple with the ComputeBerespHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBerespHeaderBytes

`func (o *Values) SetComputeBerespHeaderBytes(v int64)`

SetComputeBerespHeaderBytes sets ComputeBerespHeaderBytes field to given value.

### HasComputeBerespHeaderBytes

`func (o *Values) HasComputeBerespHeaderBytes() bool`

HasComputeBerespHeaderBytes returns a boolean if a field has been set.

### GetComputeExecutionTimeMs

`func (o *Values) GetComputeExecutionTimeMs() int64`

GetComputeExecutionTimeMs returns the ComputeExecutionTimeMs field if non-nil, zero value otherwise.

### GetComputeExecutionTimeMsOk

`func (o *Values) GetComputeExecutionTimeMsOk() (*int64, bool)`

GetComputeExecutionTimeMsOk returns a tuple with the ComputeExecutionTimeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeExecutionTimeMs

`func (o *Values) SetComputeExecutionTimeMs(v int64)`

SetComputeExecutionTimeMs sets ComputeExecutionTimeMs field to given value.

### HasComputeExecutionTimeMs

`func (o *Values) HasComputeExecutionTimeMs() bool`

HasComputeExecutionTimeMs returns a boolean if a field has been set.

### GetComputeOriginStatus1xx

`func (o *Values) GetComputeOriginStatus1xx() int64`

GetComputeOriginStatus1xx returns the ComputeOriginStatus1xx field if non-nil, zero value otherwise.

### GetComputeOriginStatus1xxOk

`func (o *Values) GetComputeOriginStatus1xxOk() (*int64, bool)`

GetComputeOriginStatus1xxOk returns a tuple with the ComputeOriginStatus1xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeOriginStatus1xx

`func (o *Values) SetComputeOriginStatus1xx(v int64)`

SetComputeOriginStatus1xx sets ComputeOriginStatus1xx field to given value.

### HasComputeOriginStatus1xx

`func (o *Values) HasComputeOriginStatus1xx() bool`

HasComputeOriginStatus1xx returns a boolean if a field has been set.

### GetComputeOriginStatus200

`func (o *Values) GetComputeOriginStatus200() int64`

GetComputeOriginStatus200 returns the ComputeOriginStatus200 field if non-nil, zero value otherwise.

### GetComputeOriginStatus200Ok

`func (o *Values) GetComputeOriginStatus200Ok() (*int64, bool)`

GetComputeOriginStatus200Ok returns a tuple with the ComputeOriginStatus200 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeOriginStatus200

`func (o *Values) SetComputeOriginStatus200(v int64)`

SetComputeOriginStatus200 sets ComputeOriginStatus200 field to given value.

### HasComputeOriginStatus200

`func (o *Values) HasComputeOriginStatus200() bool`

HasComputeOriginStatus200 returns a boolean if a field has been set.

### GetComputeOriginStatus204

`func (o *Values) GetComputeOriginStatus204() int64`

GetComputeOriginStatus204 returns the ComputeOriginStatus204 field if non-nil, zero value otherwise.

### GetComputeOriginStatus204Ok

`func (o *Values) GetComputeOriginStatus204Ok() (*int64, bool)`

GetComputeOriginStatus204Ok returns a tuple with the ComputeOriginStatus204 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeOriginStatus204

`func (o *Values) SetComputeOriginStatus204(v int64)`

SetComputeOriginStatus204 sets ComputeOriginStatus204 field to given value.

### HasComputeOriginStatus204

`func (o *Values) HasComputeOriginStatus204() bool`

HasComputeOriginStatus204 returns a boolean if a field has been set.

### GetComputeOriginStatus206

`func (o *Values) GetComputeOriginStatus206() int64`

GetComputeOriginStatus206 returns the ComputeOriginStatus206 field if non-nil, zero value otherwise.

### GetComputeOriginStatus206Ok

`func (o *Values) GetComputeOriginStatus206Ok() (*int64, bool)`

GetComputeOriginStatus206Ok returns a tuple with the ComputeOriginStatus206 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeOriginStatus206

`func (o *Values) SetComputeOriginStatus206(v int64)`

SetComputeOriginStatus206 sets ComputeOriginStatus206 field to given value.

### HasComputeOriginStatus206

`func (o *Values) HasComputeOriginStatus206() bool`

HasComputeOriginStatus206 returns a boolean if a field has been set.

### GetComputeOriginStatus2xx

`func (o *Values) GetComputeOriginStatus2xx() int64`

GetComputeOriginStatus2xx returns the ComputeOriginStatus2xx field if non-nil, zero value otherwise.

### GetComputeOriginStatus2xxOk

`func (o *Values) GetComputeOriginStatus2xxOk() (*int64, bool)`

GetComputeOriginStatus2xxOk returns a tuple with the ComputeOriginStatus2xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeOriginStatus2xx

`func (o *Values) SetComputeOriginStatus2xx(v int64)`

SetComputeOriginStatus2xx sets ComputeOriginStatus2xx field to given value.

### HasComputeOriginStatus2xx

`func (o *Values) HasComputeOriginStatus2xx() bool`

HasComputeOriginStatus2xx returns a boolean if a field has been set.

### GetComputeOriginStatus301

`func (o *Values) GetComputeOriginStatus301() int64`

GetComputeOriginStatus301 returns the ComputeOriginStatus301 field if non-nil, zero value otherwise.

### GetComputeOriginStatus301Ok

`func (o *Values) GetComputeOriginStatus301Ok() (*int64, bool)`

GetComputeOriginStatus301Ok returns a tuple with the ComputeOriginStatus301 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeOriginStatus301

`func (o *Values) SetComputeOriginStatus301(v int64)`

SetComputeOriginStatus301 sets ComputeOriginStatus301 field to given value.

### HasComputeOriginStatus301

`func (o *Values) HasComputeOriginStatus301() bool`

HasComputeOriginStatus301 returns a boolean if a field has been set.

### GetComputeOriginStatus302

`func (o *Values) GetComputeOriginStatus302() int64`

GetComputeOriginStatus302 returns the ComputeOriginStatus302 field if non-nil, zero value otherwise.

### GetComputeOriginStatus302Ok

`func (o *Values) GetComputeOriginStatus302Ok() (*int64, bool)`

GetComputeOriginStatus302Ok returns a tuple with the ComputeOriginStatus302 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeOriginStatus302

`func (o *Values) SetComputeOriginStatus302(v int64)`

SetComputeOriginStatus302 sets ComputeOriginStatus302 field to given value.

### HasComputeOriginStatus302

`func (o *Values) HasComputeOriginStatus302() bool`

HasComputeOriginStatus302 returns a boolean if a field has been set.

### GetComputeOriginStatus304

`func (o *Values) GetComputeOriginStatus304() int64`

GetComputeOriginStatus304 returns the ComputeOriginStatus304 field if non-nil, zero value otherwise.

### GetComputeOriginStatus304Ok

`func (o *Values) GetComputeOriginStatus304Ok() (*int64, bool)`

GetComputeOriginStatus304Ok returns a tuple with the ComputeOriginStatus304 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeOriginStatus304

`func (o *Values) SetComputeOriginStatus304(v int64)`

SetComputeOriginStatus304 sets ComputeOriginStatus304 field to given value.

### HasComputeOriginStatus304

`func (o *Values) HasComputeOriginStatus304() bool`

HasComputeOriginStatus304 returns a boolean if a field has been set.

### GetComputeOriginStatus3xx

`func (o *Values) GetComputeOriginStatus3xx() int64`

GetComputeOriginStatus3xx returns the ComputeOriginStatus3xx field if non-nil, zero value otherwise.

### GetComputeOriginStatus3xxOk

`func (o *Values) GetComputeOriginStatus3xxOk() (*int64, bool)`

GetComputeOriginStatus3xxOk returns a tuple with the ComputeOriginStatus3xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeOriginStatus3xx

`func (o *Values) SetComputeOriginStatus3xx(v int64)`

SetComputeOriginStatus3xx sets ComputeOriginStatus3xx field to given value.

### HasComputeOriginStatus3xx

`func (o *Values) HasComputeOriginStatus3xx() bool`

HasComputeOriginStatus3xx returns a boolean if a field has been set.

### GetComputeOriginStatus400

`func (o *Values) GetComputeOriginStatus400() int64`

GetComputeOriginStatus400 returns the ComputeOriginStatus400 field if non-nil, zero value otherwise.

### GetComputeOriginStatus400Ok

`func (o *Values) GetComputeOriginStatus400Ok() (*int64, bool)`

GetComputeOriginStatus400Ok returns a tuple with the ComputeOriginStatus400 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeOriginStatus400

`func (o *Values) SetComputeOriginStatus400(v int64)`

SetComputeOriginStatus400 sets ComputeOriginStatus400 field to given value.

### HasComputeOriginStatus400

`func (o *Values) HasComputeOriginStatus400() bool`

HasComputeOriginStatus400 returns a boolean if a field has been set.

### GetComputeOriginStatus401

`func (o *Values) GetComputeOriginStatus401() int64`

GetComputeOriginStatus401 returns the ComputeOriginStatus401 field if non-nil, zero value otherwise.

### GetComputeOriginStatus401Ok

`func (o *Values) GetComputeOriginStatus401Ok() (*int64, bool)`

GetComputeOriginStatus401Ok returns a tuple with the ComputeOriginStatus401 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeOriginStatus401

`func (o *Values) SetComputeOriginStatus401(v int64)`

SetComputeOriginStatus401 sets ComputeOriginStatus401 field to given value.

### HasComputeOriginStatus401

`func (o *Values) HasComputeOriginStatus401() bool`

HasComputeOriginStatus401 returns a boolean if a field has been set.

### GetComputeOriginStatus403

`func (o *Values) GetComputeOriginStatus403() int64`

GetComputeOriginStatus403 returns the ComputeOriginStatus403 field if non-nil, zero value otherwise.

### GetComputeOriginStatus403Ok

`func (o *Values) GetComputeOriginStatus403Ok() (*int64, bool)`

GetComputeOriginStatus403Ok returns a tuple with the ComputeOriginStatus403 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeOriginStatus403

`func (o *Values) SetComputeOriginStatus403(v int64)`

SetComputeOriginStatus403 sets ComputeOriginStatus403 field to given value.

### HasComputeOriginStatus403

`func (o *Values) HasComputeOriginStatus403() bool`

HasComputeOriginStatus403 returns a boolean if a field has been set.

### GetComputeOriginStatus404

`func (o *Values) GetComputeOriginStatus404() int64`

GetComputeOriginStatus404 returns the ComputeOriginStatus404 field if non-nil, zero value otherwise.

### GetComputeOriginStatus404Ok

`func (o *Values) GetComputeOriginStatus404Ok() (*int64, bool)`

GetComputeOriginStatus404Ok returns a tuple with the ComputeOriginStatus404 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeOriginStatus404

`func (o *Values) SetComputeOriginStatus404(v int64)`

SetComputeOriginStatus404 sets ComputeOriginStatus404 field to given value.

### HasComputeOriginStatus404

`func (o *Values) HasComputeOriginStatus404() bool`

HasComputeOriginStatus404 returns a boolean if a field has been set.

### GetComputeOriginStatus416

`func (o *Values) GetComputeOriginStatus416() int64`

GetComputeOriginStatus416 returns the ComputeOriginStatus416 field if non-nil, zero value otherwise.

### GetComputeOriginStatus416Ok

`func (o *Values) GetComputeOriginStatus416Ok() (*int64, bool)`

GetComputeOriginStatus416Ok returns a tuple with the ComputeOriginStatus416 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeOriginStatus416

`func (o *Values) SetComputeOriginStatus416(v int64)`

SetComputeOriginStatus416 sets ComputeOriginStatus416 field to given value.

### HasComputeOriginStatus416

`func (o *Values) HasComputeOriginStatus416() bool`

HasComputeOriginStatus416 returns a boolean if a field has been set.

### GetComputeOriginStatus429

`func (o *Values) GetComputeOriginStatus429() int64`

GetComputeOriginStatus429 returns the ComputeOriginStatus429 field if non-nil, zero value otherwise.

### GetComputeOriginStatus429Ok

`func (o *Values) GetComputeOriginStatus429Ok() (*int64, bool)`

GetComputeOriginStatus429Ok returns a tuple with the ComputeOriginStatus429 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeOriginStatus429

`func (o *Values) SetComputeOriginStatus429(v int64)`

SetComputeOriginStatus429 sets ComputeOriginStatus429 field to given value.

### HasComputeOriginStatus429

`func (o *Values) HasComputeOriginStatus429() bool`

HasComputeOriginStatus429 returns a boolean if a field has been set.

### GetComputeOriginStatus4xx

`func (o *Values) GetComputeOriginStatus4xx() int64`

GetComputeOriginStatus4xx returns the ComputeOriginStatus4xx field if non-nil, zero value otherwise.

### GetComputeOriginStatus4xxOk

`func (o *Values) GetComputeOriginStatus4xxOk() (*int64, bool)`

GetComputeOriginStatus4xxOk returns a tuple with the ComputeOriginStatus4xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeOriginStatus4xx

`func (o *Values) SetComputeOriginStatus4xx(v int64)`

SetComputeOriginStatus4xx sets ComputeOriginStatus4xx field to given value.

### HasComputeOriginStatus4xx

`func (o *Values) HasComputeOriginStatus4xx() bool`

HasComputeOriginStatus4xx returns a boolean if a field has been set.

### GetComputeOriginStatus500

`func (o *Values) GetComputeOriginStatus500() int64`

GetComputeOriginStatus500 returns the ComputeOriginStatus500 field if non-nil, zero value otherwise.

### GetComputeOriginStatus500Ok

`func (o *Values) GetComputeOriginStatus500Ok() (*int64, bool)`

GetComputeOriginStatus500Ok returns a tuple with the ComputeOriginStatus500 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeOriginStatus500

`func (o *Values) SetComputeOriginStatus500(v int64)`

SetComputeOriginStatus500 sets ComputeOriginStatus500 field to given value.

### HasComputeOriginStatus500

`func (o *Values) HasComputeOriginStatus500() bool`

HasComputeOriginStatus500 returns a boolean if a field has been set.

### GetComputeOriginStatus501

`func (o *Values) GetComputeOriginStatus501() int64`

GetComputeOriginStatus501 returns the ComputeOriginStatus501 field if non-nil, zero value otherwise.

### GetComputeOriginStatus501Ok

`func (o *Values) GetComputeOriginStatus501Ok() (*int64, bool)`

GetComputeOriginStatus501Ok returns a tuple with the ComputeOriginStatus501 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeOriginStatus501

`func (o *Values) SetComputeOriginStatus501(v int64)`

SetComputeOriginStatus501 sets ComputeOriginStatus501 field to given value.

### HasComputeOriginStatus501

`func (o *Values) HasComputeOriginStatus501() bool`

HasComputeOriginStatus501 returns a boolean if a field has been set.

### GetComputeOriginStatus502

`func (o *Values) GetComputeOriginStatus502() int64`

GetComputeOriginStatus502 returns the ComputeOriginStatus502 field if non-nil, zero value otherwise.

### GetComputeOriginStatus502Ok

`func (o *Values) GetComputeOriginStatus502Ok() (*int64, bool)`

GetComputeOriginStatus502Ok returns a tuple with the ComputeOriginStatus502 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeOriginStatus502

`func (o *Values) SetComputeOriginStatus502(v int64)`

SetComputeOriginStatus502 sets ComputeOriginStatus502 field to given value.

### HasComputeOriginStatus502

`func (o *Values) HasComputeOriginStatus502() bool`

HasComputeOriginStatus502 returns a boolean if a field has been set.

### GetComputeOriginStatus503

`func (o *Values) GetComputeOriginStatus503() int64`

GetComputeOriginStatus503 returns the ComputeOriginStatus503 field if non-nil, zero value otherwise.

### GetComputeOriginStatus503Ok

`func (o *Values) GetComputeOriginStatus503Ok() (*int64, bool)`

GetComputeOriginStatus503Ok returns a tuple with the ComputeOriginStatus503 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeOriginStatus503

`func (o *Values) SetComputeOriginStatus503(v int64)`

SetComputeOriginStatus503 sets ComputeOriginStatus503 field to given value.

### HasComputeOriginStatus503

`func (o *Values) HasComputeOriginStatus503() bool`

HasComputeOriginStatus503 returns a boolean if a field has been set.

### GetComputeOriginStatus504

`func (o *Values) GetComputeOriginStatus504() int64`

GetComputeOriginStatus504 returns the ComputeOriginStatus504 field if non-nil, zero value otherwise.

### GetComputeOriginStatus504Ok

`func (o *Values) GetComputeOriginStatus504Ok() (*int64, bool)`

GetComputeOriginStatus504Ok returns a tuple with the ComputeOriginStatus504 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeOriginStatus504

`func (o *Values) SetComputeOriginStatus504(v int64)`

SetComputeOriginStatus504 sets ComputeOriginStatus504 field to given value.

### HasComputeOriginStatus504

`func (o *Values) HasComputeOriginStatus504() bool`

HasComputeOriginStatus504 returns a boolean if a field has been set.

### GetComputeOriginStatus505

`func (o *Values) GetComputeOriginStatus505() int64`

GetComputeOriginStatus505 returns the ComputeOriginStatus505 field if non-nil, zero value otherwise.

### GetComputeOriginStatus505Ok

`func (o *Values) GetComputeOriginStatus505Ok() (*int64, bool)`

GetComputeOriginStatus505Ok returns a tuple with the ComputeOriginStatus505 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeOriginStatus505

`func (o *Values) SetComputeOriginStatus505(v int64)`

SetComputeOriginStatus505 sets ComputeOriginStatus505 field to given value.

### HasComputeOriginStatus505

`func (o *Values) HasComputeOriginStatus505() bool`

HasComputeOriginStatus505 returns a boolean if a field has been set.

### GetComputeOriginStatus530

`func (o *Values) GetComputeOriginStatus530() int64`

GetComputeOriginStatus530 returns the ComputeOriginStatus530 field if non-nil, zero value otherwise.

### GetComputeOriginStatus530Ok

`func (o *Values) GetComputeOriginStatus530Ok() (*int64, bool)`

GetComputeOriginStatus530Ok returns a tuple with the ComputeOriginStatus530 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeOriginStatus530

`func (o *Values) SetComputeOriginStatus530(v int64)`

SetComputeOriginStatus530 sets ComputeOriginStatus530 field to given value.

### HasComputeOriginStatus530

`func (o *Values) HasComputeOriginStatus530() bool`

HasComputeOriginStatus530 returns a boolean if a field has been set.

### GetComputeOriginStatus5xx

`func (o *Values) GetComputeOriginStatus5xx() int64`

GetComputeOriginStatus5xx returns the ComputeOriginStatus5xx field if non-nil, zero value otherwise.

### GetComputeOriginStatus5xxOk

`func (o *Values) GetComputeOriginStatus5xxOk() (*int64, bool)`

GetComputeOriginStatus5xxOk returns a tuple with the ComputeOriginStatus5xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeOriginStatus5xx

`func (o *Values) SetComputeOriginStatus5xx(v int64)`

SetComputeOriginStatus5xx sets ComputeOriginStatus5xx field to given value.

### HasComputeOriginStatus5xx

`func (o *Values) HasComputeOriginStatus5xx() bool`

HasComputeOriginStatus5xx returns a boolean if a field has been set.

### GetComputeReqBodyBytes

`func (o *Values) GetComputeReqBodyBytes() int64`

GetComputeReqBodyBytes returns the ComputeReqBodyBytes field if non-nil, zero value otherwise.

### GetComputeReqBodyBytesOk

`func (o *Values) GetComputeReqBodyBytesOk() (*int64, bool)`

GetComputeReqBodyBytesOk returns a tuple with the ComputeReqBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeReqBodyBytes

`func (o *Values) SetComputeReqBodyBytes(v int64)`

SetComputeReqBodyBytes sets ComputeReqBodyBytes field to given value.

### HasComputeReqBodyBytes

`func (o *Values) HasComputeReqBodyBytes() bool`

HasComputeReqBodyBytes returns a boolean if a field has been set.

### GetComputeReqHeaderBytes

`func (o *Values) GetComputeReqHeaderBytes() int64`

GetComputeReqHeaderBytes returns the ComputeReqHeaderBytes field if non-nil, zero value otherwise.

### GetComputeReqHeaderBytesOk

`func (o *Values) GetComputeReqHeaderBytesOk() (*int64, bool)`

GetComputeReqHeaderBytesOk returns a tuple with the ComputeReqHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeReqHeaderBytes

`func (o *Values) SetComputeReqHeaderBytes(v int64)`

SetComputeReqHeaderBytes sets ComputeReqHeaderBytes field to given value.

### HasComputeReqHeaderBytes

`func (o *Values) HasComputeReqHeaderBytes() bool`

HasComputeReqHeaderBytes returns a boolean if a field has been set.

### GetComputeRequestTimeBilledMs

`func (o *Values) GetComputeRequestTimeBilledMs() int64`

GetComputeRequestTimeBilledMs returns the ComputeRequestTimeBilledMs field if non-nil, zero value otherwise.

### GetComputeRequestTimeBilledMsOk

`func (o *Values) GetComputeRequestTimeBilledMsOk() (*int64, bool)`

GetComputeRequestTimeBilledMsOk returns a tuple with the ComputeRequestTimeBilledMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRequestTimeBilledMs

`func (o *Values) SetComputeRequestTimeBilledMs(v int64)`

SetComputeRequestTimeBilledMs sets ComputeRequestTimeBilledMs field to given value.

### HasComputeRequestTimeBilledMs

`func (o *Values) HasComputeRequestTimeBilledMs() bool`

HasComputeRequestTimeBilledMs returns a boolean if a field has been set.

### GetComputeRequestTimeMs

`func (o *Values) GetComputeRequestTimeMs() int64`

GetComputeRequestTimeMs returns the ComputeRequestTimeMs field if non-nil, zero value otherwise.

### GetComputeRequestTimeMsOk

`func (o *Values) GetComputeRequestTimeMsOk() (*int64, bool)`

GetComputeRequestTimeMsOk returns a tuple with the ComputeRequestTimeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRequestTimeMs

`func (o *Values) SetComputeRequestTimeMs(v int64)`

SetComputeRequestTimeMs sets ComputeRequestTimeMs field to given value.

### HasComputeRequestTimeMs

`func (o *Values) HasComputeRequestTimeMs() bool`

HasComputeRequestTimeMs returns a boolean if a field has been set.

### GetComputeRequest

`func (o *Values) GetComputeRequest() int64`

GetComputeRequest returns the ComputeRequest field if non-nil, zero value otherwise.

### GetComputeRequestOk

`func (o *Values) GetComputeRequestOk() (*int64, bool)`

GetComputeRequestOk returns a tuple with the ComputeRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRequest

`func (o *Values) SetComputeRequest(v int64)`

SetComputeRequest sets ComputeRequest field to given value.

### HasComputeRequest

`func (o *Values) HasComputeRequest() bool`

HasComputeRequest returns a boolean if a field has been set.

### GetComputeRespBodyBytes

`func (o *Values) GetComputeRespBodyBytes() int64`

GetComputeRespBodyBytes returns the ComputeRespBodyBytes field if non-nil, zero value otherwise.

### GetComputeRespBodyBytesOk

`func (o *Values) GetComputeRespBodyBytesOk() (*int64, bool)`

GetComputeRespBodyBytesOk returns a tuple with the ComputeRespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespBodyBytes

`func (o *Values) SetComputeRespBodyBytes(v int64)`

SetComputeRespBodyBytes sets ComputeRespBodyBytes field to given value.

### HasComputeRespBodyBytes

`func (o *Values) HasComputeRespBodyBytes() bool`

HasComputeRespBodyBytes returns a boolean if a field has been set.

### GetComputeRespHeaderBytes

`func (o *Values) GetComputeRespHeaderBytes() int64`

GetComputeRespHeaderBytes returns the ComputeRespHeaderBytes field if non-nil, zero value otherwise.

### GetComputeRespHeaderBytesOk

`func (o *Values) GetComputeRespHeaderBytesOk() (*int64, bool)`

GetComputeRespHeaderBytesOk returns a tuple with the ComputeRespHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespHeaderBytes

`func (o *Values) SetComputeRespHeaderBytes(v int64)`

SetComputeRespHeaderBytes sets ComputeRespHeaderBytes field to given value.

### HasComputeRespHeaderBytes

`func (o *Values) HasComputeRespHeaderBytes() bool`

HasComputeRespHeaderBytes returns a boolean if a field has been set.

### GetComputeRespStatus103

`func (o *Values) GetComputeRespStatus103() int64`

GetComputeRespStatus103 returns the ComputeRespStatus103 field if non-nil, zero value otherwise.

### GetComputeRespStatus103Ok

`func (o *Values) GetComputeRespStatus103Ok() (*int64, bool)`

GetComputeRespStatus103Ok returns a tuple with the ComputeRespStatus103 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespStatus103

`func (o *Values) SetComputeRespStatus103(v int64)`

SetComputeRespStatus103 sets ComputeRespStatus103 field to given value.

### HasComputeRespStatus103

`func (o *Values) HasComputeRespStatus103() bool`

HasComputeRespStatus103 returns a boolean if a field has been set.

### GetComputeRespStatus1xx

`func (o *Values) GetComputeRespStatus1xx() int64`

GetComputeRespStatus1xx returns the ComputeRespStatus1xx field if non-nil, zero value otherwise.

### GetComputeRespStatus1xxOk

`func (o *Values) GetComputeRespStatus1xxOk() (*int64, bool)`

GetComputeRespStatus1xxOk returns a tuple with the ComputeRespStatus1xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespStatus1xx

`func (o *Values) SetComputeRespStatus1xx(v int64)`

SetComputeRespStatus1xx sets ComputeRespStatus1xx field to given value.

### HasComputeRespStatus1xx

`func (o *Values) HasComputeRespStatus1xx() bool`

HasComputeRespStatus1xx returns a boolean if a field has been set.

### GetComputeRespStatus200

`func (o *Values) GetComputeRespStatus200() int64`

GetComputeRespStatus200 returns the ComputeRespStatus200 field if non-nil, zero value otherwise.

### GetComputeRespStatus200Ok

`func (o *Values) GetComputeRespStatus200Ok() (*int64, bool)`

GetComputeRespStatus200Ok returns a tuple with the ComputeRespStatus200 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespStatus200

`func (o *Values) SetComputeRespStatus200(v int64)`

SetComputeRespStatus200 sets ComputeRespStatus200 field to given value.

### HasComputeRespStatus200

`func (o *Values) HasComputeRespStatus200() bool`

HasComputeRespStatus200 returns a boolean if a field has been set.

### GetComputeRespStatus204

`func (o *Values) GetComputeRespStatus204() int64`

GetComputeRespStatus204 returns the ComputeRespStatus204 field if non-nil, zero value otherwise.

### GetComputeRespStatus204Ok

`func (o *Values) GetComputeRespStatus204Ok() (*int64, bool)`

GetComputeRespStatus204Ok returns a tuple with the ComputeRespStatus204 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespStatus204

`func (o *Values) SetComputeRespStatus204(v int64)`

SetComputeRespStatus204 sets ComputeRespStatus204 field to given value.

### HasComputeRespStatus204

`func (o *Values) HasComputeRespStatus204() bool`

HasComputeRespStatus204 returns a boolean if a field has been set.

### GetComputeRespStatus206

`func (o *Values) GetComputeRespStatus206() int64`

GetComputeRespStatus206 returns the ComputeRespStatus206 field if non-nil, zero value otherwise.

### GetComputeRespStatus206Ok

`func (o *Values) GetComputeRespStatus206Ok() (*int64, bool)`

GetComputeRespStatus206Ok returns a tuple with the ComputeRespStatus206 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespStatus206

`func (o *Values) SetComputeRespStatus206(v int64)`

SetComputeRespStatus206 sets ComputeRespStatus206 field to given value.

### HasComputeRespStatus206

`func (o *Values) HasComputeRespStatus206() bool`

HasComputeRespStatus206 returns a boolean if a field has been set.

### GetComputeRespStatus2xx

`func (o *Values) GetComputeRespStatus2xx() int64`

GetComputeRespStatus2xx returns the ComputeRespStatus2xx field if non-nil, zero value otherwise.

### GetComputeRespStatus2xxOk

`func (o *Values) GetComputeRespStatus2xxOk() (*int64, bool)`

GetComputeRespStatus2xxOk returns a tuple with the ComputeRespStatus2xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespStatus2xx

`func (o *Values) SetComputeRespStatus2xx(v int64)`

SetComputeRespStatus2xx sets ComputeRespStatus2xx field to given value.

### HasComputeRespStatus2xx

`func (o *Values) HasComputeRespStatus2xx() bool`

HasComputeRespStatus2xx returns a boolean if a field has been set.

### GetComputeRespStatus301

`func (o *Values) GetComputeRespStatus301() int64`

GetComputeRespStatus301 returns the ComputeRespStatus301 field if non-nil, zero value otherwise.

### GetComputeRespStatus301Ok

`func (o *Values) GetComputeRespStatus301Ok() (*int64, bool)`

GetComputeRespStatus301Ok returns a tuple with the ComputeRespStatus301 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespStatus301

`func (o *Values) SetComputeRespStatus301(v int64)`

SetComputeRespStatus301 sets ComputeRespStatus301 field to given value.

### HasComputeRespStatus301

`func (o *Values) HasComputeRespStatus301() bool`

HasComputeRespStatus301 returns a boolean if a field has been set.

### GetComputeRespStatus302

`func (o *Values) GetComputeRespStatus302() int64`

GetComputeRespStatus302 returns the ComputeRespStatus302 field if non-nil, zero value otherwise.

### GetComputeRespStatus302Ok

`func (o *Values) GetComputeRespStatus302Ok() (*int64, bool)`

GetComputeRespStatus302Ok returns a tuple with the ComputeRespStatus302 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespStatus302

`func (o *Values) SetComputeRespStatus302(v int64)`

SetComputeRespStatus302 sets ComputeRespStatus302 field to given value.

### HasComputeRespStatus302

`func (o *Values) HasComputeRespStatus302() bool`

HasComputeRespStatus302 returns a boolean if a field has been set.

### GetComputeRespStatus304

`func (o *Values) GetComputeRespStatus304() int64`

GetComputeRespStatus304 returns the ComputeRespStatus304 field if non-nil, zero value otherwise.

### GetComputeRespStatus304Ok

`func (o *Values) GetComputeRespStatus304Ok() (*int64, bool)`

GetComputeRespStatus304Ok returns a tuple with the ComputeRespStatus304 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespStatus304

`func (o *Values) SetComputeRespStatus304(v int64)`

SetComputeRespStatus304 sets ComputeRespStatus304 field to given value.

### HasComputeRespStatus304

`func (o *Values) HasComputeRespStatus304() bool`

HasComputeRespStatus304 returns a boolean if a field has been set.

### GetComputeRespStatus3xx

`func (o *Values) GetComputeRespStatus3xx() int64`

GetComputeRespStatus3xx returns the ComputeRespStatus3xx field if non-nil, zero value otherwise.

### GetComputeRespStatus3xxOk

`func (o *Values) GetComputeRespStatus3xxOk() (*int64, bool)`

GetComputeRespStatus3xxOk returns a tuple with the ComputeRespStatus3xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespStatus3xx

`func (o *Values) SetComputeRespStatus3xx(v int64)`

SetComputeRespStatus3xx sets ComputeRespStatus3xx field to given value.

### HasComputeRespStatus3xx

`func (o *Values) HasComputeRespStatus3xx() bool`

HasComputeRespStatus3xx returns a boolean if a field has been set.

### GetComputeRespStatus400

`func (o *Values) GetComputeRespStatus400() int64`

GetComputeRespStatus400 returns the ComputeRespStatus400 field if non-nil, zero value otherwise.

### GetComputeRespStatus400Ok

`func (o *Values) GetComputeRespStatus400Ok() (*int64, bool)`

GetComputeRespStatus400Ok returns a tuple with the ComputeRespStatus400 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespStatus400

`func (o *Values) SetComputeRespStatus400(v int64)`

SetComputeRespStatus400 sets ComputeRespStatus400 field to given value.

### HasComputeRespStatus400

`func (o *Values) HasComputeRespStatus400() bool`

HasComputeRespStatus400 returns a boolean if a field has been set.

### GetComputeRespStatus401

`func (o *Values) GetComputeRespStatus401() int64`

GetComputeRespStatus401 returns the ComputeRespStatus401 field if non-nil, zero value otherwise.

### GetComputeRespStatus401Ok

`func (o *Values) GetComputeRespStatus401Ok() (*int64, bool)`

GetComputeRespStatus401Ok returns a tuple with the ComputeRespStatus401 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespStatus401

`func (o *Values) SetComputeRespStatus401(v int64)`

SetComputeRespStatus401 sets ComputeRespStatus401 field to given value.

### HasComputeRespStatus401

`func (o *Values) HasComputeRespStatus401() bool`

HasComputeRespStatus401 returns a boolean if a field has been set.

### GetComputeRespStatus403

`func (o *Values) GetComputeRespStatus403() int64`

GetComputeRespStatus403 returns the ComputeRespStatus403 field if non-nil, zero value otherwise.

### GetComputeRespStatus403Ok

`func (o *Values) GetComputeRespStatus403Ok() (*int64, bool)`

GetComputeRespStatus403Ok returns a tuple with the ComputeRespStatus403 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespStatus403

`func (o *Values) SetComputeRespStatus403(v int64)`

SetComputeRespStatus403 sets ComputeRespStatus403 field to given value.

### HasComputeRespStatus403

`func (o *Values) HasComputeRespStatus403() bool`

HasComputeRespStatus403 returns a boolean if a field has been set.

### GetComputeRespStatus404

`func (o *Values) GetComputeRespStatus404() int64`

GetComputeRespStatus404 returns the ComputeRespStatus404 field if non-nil, zero value otherwise.

### GetComputeRespStatus404Ok

`func (o *Values) GetComputeRespStatus404Ok() (*int64, bool)`

GetComputeRespStatus404Ok returns a tuple with the ComputeRespStatus404 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespStatus404

`func (o *Values) SetComputeRespStatus404(v int64)`

SetComputeRespStatus404 sets ComputeRespStatus404 field to given value.

### HasComputeRespStatus404

`func (o *Values) HasComputeRespStatus404() bool`

HasComputeRespStatus404 returns a boolean if a field has been set.

### GetComputeRespStatus416

`func (o *Values) GetComputeRespStatus416() int64`

GetComputeRespStatus416 returns the ComputeRespStatus416 field if non-nil, zero value otherwise.

### GetComputeRespStatus416Ok

`func (o *Values) GetComputeRespStatus416Ok() (*int64, bool)`

GetComputeRespStatus416Ok returns a tuple with the ComputeRespStatus416 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespStatus416

`func (o *Values) SetComputeRespStatus416(v int64)`

SetComputeRespStatus416 sets ComputeRespStatus416 field to given value.

### HasComputeRespStatus416

`func (o *Values) HasComputeRespStatus416() bool`

HasComputeRespStatus416 returns a boolean if a field has been set.

### GetComputeRespStatus429

`func (o *Values) GetComputeRespStatus429() int64`

GetComputeRespStatus429 returns the ComputeRespStatus429 field if non-nil, zero value otherwise.

### GetComputeRespStatus429Ok

`func (o *Values) GetComputeRespStatus429Ok() (*int64, bool)`

GetComputeRespStatus429Ok returns a tuple with the ComputeRespStatus429 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespStatus429

`func (o *Values) SetComputeRespStatus429(v int64)`

SetComputeRespStatus429 sets ComputeRespStatus429 field to given value.

### HasComputeRespStatus429

`func (o *Values) HasComputeRespStatus429() bool`

HasComputeRespStatus429 returns a boolean if a field has been set.

### GetComputeRespStatus4xx

`func (o *Values) GetComputeRespStatus4xx() int64`

GetComputeRespStatus4xx returns the ComputeRespStatus4xx field if non-nil, zero value otherwise.

### GetComputeRespStatus4xxOk

`func (o *Values) GetComputeRespStatus4xxOk() (*int64, bool)`

GetComputeRespStatus4xxOk returns a tuple with the ComputeRespStatus4xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespStatus4xx

`func (o *Values) SetComputeRespStatus4xx(v int64)`

SetComputeRespStatus4xx sets ComputeRespStatus4xx field to given value.

### HasComputeRespStatus4xx

`func (o *Values) HasComputeRespStatus4xx() bool`

HasComputeRespStatus4xx returns a boolean if a field has been set.

### GetComputeRespStatus500

`func (o *Values) GetComputeRespStatus500() int64`

GetComputeRespStatus500 returns the ComputeRespStatus500 field if non-nil, zero value otherwise.

### GetComputeRespStatus500Ok

`func (o *Values) GetComputeRespStatus500Ok() (*int64, bool)`

GetComputeRespStatus500Ok returns a tuple with the ComputeRespStatus500 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespStatus500

`func (o *Values) SetComputeRespStatus500(v int64)`

SetComputeRespStatus500 sets ComputeRespStatus500 field to given value.

### HasComputeRespStatus500

`func (o *Values) HasComputeRespStatus500() bool`

HasComputeRespStatus500 returns a boolean if a field has been set.

### GetComputeRespStatus501

`func (o *Values) GetComputeRespStatus501() int64`

GetComputeRespStatus501 returns the ComputeRespStatus501 field if non-nil, zero value otherwise.

### GetComputeRespStatus501Ok

`func (o *Values) GetComputeRespStatus501Ok() (*int64, bool)`

GetComputeRespStatus501Ok returns a tuple with the ComputeRespStatus501 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespStatus501

`func (o *Values) SetComputeRespStatus501(v int64)`

SetComputeRespStatus501 sets ComputeRespStatus501 field to given value.

### HasComputeRespStatus501

`func (o *Values) HasComputeRespStatus501() bool`

HasComputeRespStatus501 returns a boolean if a field has been set.

### GetComputeRespStatus502

`func (o *Values) GetComputeRespStatus502() int64`

GetComputeRespStatus502 returns the ComputeRespStatus502 field if non-nil, zero value otherwise.

### GetComputeRespStatus502Ok

`func (o *Values) GetComputeRespStatus502Ok() (*int64, bool)`

GetComputeRespStatus502Ok returns a tuple with the ComputeRespStatus502 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespStatus502

`func (o *Values) SetComputeRespStatus502(v int64)`

SetComputeRespStatus502 sets ComputeRespStatus502 field to given value.

### HasComputeRespStatus502

`func (o *Values) HasComputeRespStatus502() bool`

HasComputeRespStatus502 returns a boolean if a field has been set.

### GetComputeRespStatus503

`func (o *Values) GetComputeRespStatus503() int64`

GetComputeRespStatus503 returns the ComputeRespStatus503 field if non-nil, zero value otherwise.

### GetComputeRespStatus503Ok

`func (o *Values) GetComputeRespStatus503Ok() (*int64, bool)`

GetComputeRespStatus503Ok returns a tuple with the ComputeRespStatus503 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespStatus503

`func (o *Values) SetComputeRespStatus503(v int64)`

SetComputeRespStatus503 sets ComputeRespStatus503 field to given value.

### HasComputeRespStatus503

`func (o *Values) HasComputeRespStatus503() bool`

HasComputeRespStatus503 returns a boolean if a field has been set.

### GetComputeRespStatus504

`func (o *Values) GetComputeRespStatus504() int64`

GetComputeRespStatus504 returns the ComputeRespStatus504 field if non-nil, zero value otherwise.

### GetComputeRespStatus504Ok

`func (o *Values) GetComputeRespStatus504Ok() (*int64, bool)`

GetComputeRespStatus504Ok returns a tuple with the ComputeRespStatus504 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespStatus504

`func (o *Values) SetComputeRespStatus504(v int64)`

SetComputeRespStatus504 sets ComputeRespStatus504 field to given value.

### HasComputeRespStatus504

`func (o *Values) HasComputeRespStatus504() bool`

HasComputeRespStatus504 returns a boolean if a field has been set.

### GetComputeRespStatus505

`func (o *Values) GetComputeRespStatus505() int64`

GetComputeRespStatus505 returns the ComputeRespStatus505 field if non-nil, zero value otherwise.

### GetComputeRespStatus505Ok

`func (o *Values) GetComputeRespStatus505Ok() (*int64, bool)`

GetComputeRespStatus505Ok returns a tuple with the ComputeRespStatus505 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespStatus505

`func (o *Values) SetComputeRespStatus505(v int64)`

SetComputeRespStatus505 sets ComputeRespStatus505 field to given value.

### HasComputeRespStatus505

`func (o *Values) HasComputeRespStatus505() bool`

HasComputeRespStatus505 returns a boolean if a field has been set.

### GetComputeRespStatus530

`func (o *Values) GetComputeRespStatus530() int64`

GetComputeRespStatus530 returns the ComputeRespStatus530 field if non-nil, zero value otherwise.

### GetComputeRespStatus530Ok

`func (o *Values) GetComputeRespStatus530Ok() (*int64, bool)`

GetComputeRespStatus530Ok returns a tuple with the ComputeRespStatus530 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespStatus530

`func (o *Values) SetComputeRespStatus530(v int64)`

SetComputeRespStatus530 sets ComputeRespStatus530 field to given value.

### HasComputeRespStatus530

`func (o *Values) HasComputeRespStatus530() bool`

HasComputeRespStatus530 returns a boolean if a field has been set.

### GetComputeRespStatus5xx

`func (o *Values) GetComputeRespStatus5xx() int64`

GetComputeRespStatus5xx returns the ComputeRespStatus5xx field if non-nil, zero value otherwise.

### GetComputeRespStatus5xxOk

`func (o *Values) GetComputeRespStatus5xxOk() (*int64, bool)`

GetComputeRespStatus5xxOk returns a tuple with the ComputeRespStatus5xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespStatus5xx

`func (o *Values) SetComputeRespStatus5xx(v int64)`

SetComputeRespStatus5xx sets ComputeRespStatus5xx field to given value.

### HasComputeRespStatus5xx

`func (o *Values) HasComputeRespStatus5xx() bool`

HasComputeRespStatus5xx returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


