# DomainInspectorMeasurements

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

### NewDomainInspectorMeasurements

`func NewDomainInspectorMeasurements() *DomainInspectorMeasurements`

NewDomainInspectorMeasurements instantiates a new DomainInspectorMeasurements object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainInspectorMeasurementsWithDefaults

`func NewDomainInspectorMeasurementsWithDefaults() *DomainInspectorMeasurements`

NewDomainInspectorMeasurementsWithDefaults instantiates a new DomainInspectorMeasurements object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEdgeRequests

`func (o *DomainInspectorMeasurements) GetEdgeRequests() int64`

GetEdgeRequests returns the EdgeRequests field if non-nil, zero value otherwise.

### GetEdgeRequestsOk

`func (o *DomainInspectorMeasurements) GetEdgeRequestsOk() (*int64, bool)`

GetEdgeRequestsOk returns a tuple with the EdgeRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeRequests

`func (o *DomainInspectorMeasurements) SetEdgeRequests(v int64)`

SetEdgeRequests sets EdgeRequests field to given value.

### HasEdgeRequests

`func (o *DomainInspectorMeasurements) HasEdgeRequests() bool`

HasEdgeRequests returns a boolean if a field has been set.

### GetEdgeRespHeaderBytes

`func (o *DomainInspectorMeasurements) GetEdgeRespHeaderBytes() int64`

GetEdgeRespHeaderBytes returns the EdgeRespHeaderBytes field if non-nil, zero value otherwise.

### GetEdgeRespHeaderBytesOk

`func (o *DomainInspectorMeasurements) GetEdgeRespHeaderBytesOk() (*int64, bool)`

GetEdgeRespHeaderBytesOk returns a tuple with the EdgeRespHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeRespHeaderBytes

`func (o *DomainInspectorMeasurements) SetEdgeRespHeaderBytes(v int64)`

SetEdgeRespHeaderBytes sets EdgeRespHeaderBytes field to given value.

### HasEdgeRespHeaderBytes

`func (o *DomainInspectorMeasurements) HasEdgeRespHeaderBytes() bool`

HasEdgeRespHeaderBytes returns a boolean if a field has been set.

### GetEdgeRespBodyBytes

`func (o *DomainInspectorMeasurements) GetEdgeRespBodyBytes() int64`

GetEdgeRespBodyBytes returns the EdgeRespBodyBytes field if non-nil, zero value otherwise.

### GetEdgeRespBodyBytesOk

`func (o *DomainInspectorMeasurements) GetEdgeRespBodyBytesOk() (*int64, bool)`

GetEdgeRespBodyBytesOk returns a tuple with the EdgeRespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeRespBodyBytes

`func (o *DomainInspectorMeasurements) SetEdgeRespBodyBytes(v int64)`

SetEdgeRespBodyBytes sets EdgeRespBodyBytes field to given value.

### HasEdgeRespBodyBytes

`func (o *DomainInspectorMeasurements) HasEdgeRespBodyBytes() bool`

HasEdgeRespBodyBytes returns a boolean if a field has been set.

### GetStatus1xx

`func (o *DomainInspectorMeasurements) GetStatus1xx() int64`

GetStatus1xx returns the Status1xx field if non-nil, zero value otherwise.

### GetStatus1xxOk

`func (o *DomainInspectorMeasurements) GetStatus1xxOk() (*int64, bool)`

GetStatus1xxOk returns a tuple with the Status1xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus1xx

`func (o *DomainInspectorMeasurements) SetStatus1xx(v int64)`

SetStatus1xx sets Status1xx field to given value.

### HasStatus1xx

`func (o *DomainInspectorMeasurements) HasStatus1xx() bool`

HasStatus1xx returns a boolean if a field has been set.

### GetStatus2xx

`func (o *DomainInspectorMeasurements) GetStatus2xx() int64`

GetStatus2xx returns the Status2xx field if non-nil, zero value otherwise.

### GetStatus2xxOk

`func (o *DomainInspectorMeasurements) GetStatus2xxOk() (*int64, bool)`

GetStatus2xxOk returns a tuple with the Status2xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus2xx

`func (o *DomainInspectorMeasurements) SetStatus2xx(v int64)`

SetStatus2xx sets Status2xx field to given value.

### HasStatus2xx

`func (o *DomainInspectorMeasurements) HasStatus2xx() bool`

HasStatus2xx returns a boolean if a field has been set.

### GetStatus3xx

`func (o *DomainInspectorMeasurements) GetStatus3xx() int64`

GetStatus3xx returns the Status3xx field if non-nil, zero value otherwise.

### GetStatus3xxOk

`func (o *DomainInspectorMeasurements) GetStatus3xxOk() (*int64, bool)`

GetStatus3xxOk returns a tuple with the Status3xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus3xx

`func (o *DomainInspectorMeasurements) SetStatus3xx(v int64)`

SetStatus3xx sets Status3xx field to given value.

### HasStatus3xx

`func (o *DomainInspectorMeasurements) HasStatus3xx() bool`

HasStatus3xx returns a boolean if a field has been set.

### GetStatus4xx

`func (o *DomainInspectorMeasurements) GetStatus4xx() int64`

GetStatus4xx returns the Status4xx field if non-nil, zero value otherwise.

### GetStatus4xxOk

`func (o *DomainInspectorMeasurements) GetStatus4xxOk() (*int64, bool)`

GetStatus4xxOk returns a tuple with the Status4xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus4xx

`func (o *DomainInspectorMeasurements) SetStatus4xx(v int64)`

SetStatus4xx sets Status4xx field to given value.

### HasStatus4xx

`func (o *DomainInspectorMeasurements) HasStatus4xx() bool`

HasStatus4xx returns a boolean if a field has been set.

### GetStatus5xx

`func (o *DomainInspectorMeasurements) GetStatus5xx() int64`

GetStatus5xx returns the Status5xx field if non-nil, zero value otherwise.

### GetStatus5xxOk

`func (o *DomainInspectorMeasurements) GetStatus5xxOk() (*int64, bool)`

GetStatus5xxOk returns a tuple with the Status5xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus5xx

`func (o *DomainInspectorMeasurements) SetStatus5xx(v int64)`

SetStatus5xx sets Status5xx field to given value.

### HasStatus5xx

`func (o *DomainInspectorMeasurements) HasStatus5xx() bool`

HasStatus5xx returns a boolean if a field has been set.

### GetStatus200

`func (o *DomainInspectorMeasurements) GetStatus200() int64`

GetStatus200 returns the Status200 field if non-nil, zero value otherwise.

### GetStatus200Ok

`func (o *DomainInspectorMeasurements) GetStatus200Ok() (*int64, bool)`

GetStatus200Ok returns a tuple with the Status200 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus200

`func (o *DomainInspectorMeasurements) SetStatus200(v int64)`

SetStatus200 sets Status200 field to given value.

### HasStatus200

`func (o *DomainInspectorMeasurements) HasStatus200() bool`

HasStatus200 returns a boolean if a field has been set.

### GetStatus204

`func (o *DomainInspectorMeasurements) GetStatus204() int64`

GetStatus204 returns the Status204 field if non-nil, zero value otherwise.

### GetStatus204Ok

`func (o *DomainInspectorMeasurements) GetStatus204Ok() (*int64, bool)`

GetStatus204Ok returns a tuple with the Status204 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus204

`func (o *DomainInspectorMeasurements) SetStatus204(v int64)`

SetStatus204 sets Status204 field to given value.

### HasStatus204

`func (o *DomainInspectorMeasurements) HasStatus204() bool`

HasStatus204 returns a boolean if a field has been set.

### GetStatus206

`func (o *DomainInspectorMeasurements) GetStatus206() int64`

GetStatus206 returns the Status206 field if non-nil, zero value otherwise.

### GetStatus206Ok

`func (o *DomainInspectorMeasurements) GetStatus206Ok() (*int64, bool)`

GetStatus206Ok returns a tuple with the Status206 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus206

`func (o *DomainInspectorMeasurements) SetStatus206(v int64)`

SetStatus206 sets Status206 field to given value.

### HasStatus206

`func (o *DomainInspectorMeasurements) HasStatus206() bool`

HasStatus206 returns a boolean if a field has been set.

### GetStatus301

`func (o *DomainInspectorMeasurements) GetStatus301() int64`

GetStatus301 returns the Status301 field if non-nil, zero value otherwise.

### GetStatus301Ok

`func (o *DomainInspectorMeasurements) GetStatus301Ok() (*int64, bool)`

GetStatus301Ok returns a tuple with the Status301 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus301

`func (o *DomainInspectorMeasurements) SetStatus301(v int64)`

SetStatus301 sets Status301 field to given value.

### HasStatus301

`func (o *DomainInspectorMeasurements) HasStatus301() bool`

HasStatus301 returns a boolean if a field has been set.

### GetStatus302

`func (o *DomainInspectorMeasurements) GetStatus302() int64`

GetStatus302 returns the Status302 field if non-nil, zero value otherwise.

### GetStatus302Ok

`func (o *DomainInspectorMeasurements) GetStatus302Ok() (*int64, bool)`

GetStatus302Ok returns a tuple with the Status302 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus302

`func (o *DomainInspectorMeasurements) SetStatus302(v int64)`

SetStatus302 sets Status302 field to given value.

### HasStatus302

`func (o *DomainInspectorMeasurements) HasStatus302() bool`

HasStatus302 returns a boolean if a field has been set.

### GetStatus304

`func (o *DomainInspectorMeasurements) GetStatus304() int64`

GetStatus304 returns the Status304 field if non-nil, zero value otherwise.

### GetStatus304Ok

`func (o *DomainInspectorMeasurements) GetStatus304Ok() (*int64, bool)`

GetStatus304Ok returns a tuple with the Status304 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus304

`func (o *DomainInspectorMeasurements) SetStatus304(v int64)`

SetStatus304 sets Status304 field to given value.

### HasStatus304

`func (o *DomainInspectorMeasurements) HasStatus304() bool`

HasStatus304 returns a boolean if a field has been set.

### GetStatus400

`func (o *DomainInspectorMeasurements) GetStatus400() int64`

GetStatus400 returns the Status400 field if non-nil, zero value otherwise.

### GetStatus400Ok

`func (o *DomainInspectorMeasurements) GetStatus400Ok() (*int64, bool)`

GetStatus400Ok returns a tuple with the Status400 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus400

`func (o *DomainInspectorMeasurements) SetStatus400(v int64)`

SetStatus400 sets Status400 field to given value.

### HasStatus400

`func (o *DomainInspectorMeasurements) HasStatus400() bool`

HasStatus400 returns a boolean if a field has been set.

### GetStatus401

`func (o *DomainInspectorMeasurements) GetStatus401() int64`

GetStatus401 returns the Status401 field if non-nil, zero value otherwise.

### GetStatus401Ok

`func (o *DomainInspectorMeasurements) GetStatus401Ok() (*int64, bool)`

GetStatus401Ok returns a tuple with the Status401 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus401

`func (o *DomainInspectorMeasurements) SetStatus401(v int64)`

SetStatus401 sets Status401 field to given value.

### HasStatus401

`func (o *DomainInspectorMeasurements) HasStatus401() bool`

HasStatus401 returns a boolean if a field has been set.

### GetStatus403

`func (o *DomainInspectorMeasurements) GetStatus403() int64`

GetStatus403 returns the Status403 field if non-nil, zero value otherwise.

### GetStatus403Ok

`func (o *DomainInspectorMeasurements) GetStatus403Ok() (*int64, bool)`

GetStatus403Ok returns a tuple with the Status403 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus403

`func (o *DomainInspectorMeasurements) SetStatus403(v int64)`

SetStatus403 sets Status403 field to given value.

### HasStatus403

`func (o *DomainInspectorMeasurements) HasStatus403() bool`

HasStatus403 returns a boolean if a field has been set.

### GetStatus404

`func (o *DomainInspectorMeasurements) GetStatus404() int64`

GetStatus404 returns the Status404 field if non-nil, zero value otherwise.

### GetStatus404Ok

`func (o *DomainInspectorMeasurements) GetStatus404Ok() (*int64, bool)`

GetStatus404Ok returns a tuple with the Status404 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus404

`func (o *DomainInspectorMeasurements) SetStatus404(v int64)`

SetStatus404 sets Status404 field to given value.

### HasStatus404

`func (o *DomainInspectorMeasurements) HasStatus404() bool`

HasStatus404 returns a boolean if a field has been set.

### GetStatus416

`func (o *DomainInspectorMeasurements) GetStatus416() int64`

GetStatus416 returns the Status416 field if non-nil, zero value otherwise.

### GetStatus416Ok

`func (o *DomainInspectorMeasurements) GetStatus416Ok() (*int64, bool)`

GetStatus416Ok returns a tuple with the Status416 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus416

`func (o *DomainInspectorMeasurements) SetStatus416(v int64)`

SetStatus416 sets Status416 field to given value.

### HasStatus416

`func (o *DomainInspectorMeasurements) HasStatus416() bool`

HasStatus416 returns a boolean if a field has been set.

### GetStatus429

`func (o *DomainInspectorMeasurements) GetStatus429() int64`

GetStatus429 returns the Status429 field if non-nil, zero value otherwise.

### GetStatus429Ok

`func (o *DomainInspectorMeasurements) GetStatus429Ok() (*int64, bool)`

GetStatus429Ok returns a tuple with the Status429 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus429

`func (o *DomainInspectorMeasurements) SetStatus429(v int64)`

SetStatus429 sets Status429 field to given value.

### HasStatus429

`func (o *DomainInspectorMeasurements) HasStatus429() bool`

HasStatus429 returns a boolean if a field has been set.

### GetStatus500

`func (o *DomainInspectorMeasurements) GetStatus500() int64`

GetStatus500 returns the Status500 field if non-nil, zero value otherwise.

### GetStatus500Ok

`func (o *DomainInspectorMeasurements) GetStatus500Ok() (*int64, bool)`

GetStatus500Ok returns a tuple with the Status500 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus500

`func (o *DomainInspectorMeasurements) SetStatus500(v int64)`

SetStatus500 sets Status500 field to given value.

### HasStatus500

`func (o *DomainInspectorMeasurements) HasStatus500() bool`

HasStatus500 returns a boolean if a field has been set.

### GetStatus501

`func (o *DomainInspectorMeasurements) GetStatus501() int64`

GetStatus501 returns the Status501 field if non-nil, zero value otherwise.

### GetStatus501Ok

`func (o *DomainInspectorMeasurements) GetStatus501Ok() (*int64, bool)`

GetStatus501Ok returns a tuple with the Status501 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus501

`func (o *DomainInspectorMeasurements) SetStatus501(v int64)`

SetStatus501 sets Status501 field to given value.

### HasStatus501

`func (o *DomainInspectorMeasurements) HasStatus501() bool`

HasStatus501 returns a boolean if a field has been set.

### GetStatus502

`func (o *DomainInspectorMeasurements) GetStatus502() int64`

GetStatus502 returns the Status502 field if non-nil, zero value otherwise.

### GetStatus502Ok

`func (o *DomainInspectorMeasurements) GetStatus502Ok() (*int64, bool)`

GetStatus502Ok returns a tuple with the Status502 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus502

`func (o *DomainInspectorMeasurements) SetStatus502(v int64)`

SetStatus502 sets Status502 field to given value.

### HasStatus502

`func (o *DomainInspectorMeasurements) HasStatus502() bool`

HasStatus502 returns a boolean if a field has been set.

### GetStatus503

`func (o *DomainInspectorMeasurements) GetStatus503() int64`

GetStatus503 returns the Status503 field if non-nil, zero value otherwise.

### GetStatus503Ok

`func (o *DomainInspectorMeasurements) GetStatus503Ok() (*int64, bool)`

GetStatus503Ok returns a tuple with the Status503 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus503

`func (o *DomainInspectorMeasurements) SetStatus503(v int64)`

SetStatus503 sets Status503 field to given value.

### HasStatus503

`func (o *DomainInspectorMeasurements) HasStatus503() bool`

HasStatus503 returns a boolean if a field has been set.

### GetStatus504

`func (o *DomainInspectorMeasurements) GetStatus504() int64`

GetStatus504 returns the Status504 field if non-nil, zero value otherwise.

### GetStatus504Ok

`func (o *DomainInspectorMeasurements) GetStatus504Ok() (*int64, bool)`

GetStatus504Ok returns a tuple with the Status504 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus504

`func (o *DomainInspectorMeasurements) SetStatus504(v int64)`

SetStatus504 sets Status504 field to given value.

### HasStatus504

`func (o *DomainInspectorMeasurements) HasStatus504() bool`

HasStatus504 returns a boolean if a field has been set.

### GetStatus505

`func (o *DomainInspectorMeasurements) GetStatus505() int64`

GetStatus505 returns the Status505 field if non-nil, zero value otherwise.

### GetStatus505Ok

`func (o *DomainInspectorMeasurements) GetStatus505Ok() (*int64, bool)`

GetStatus505Ok returns a tuple with the Status505 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus505

`func (o *DomainInspectorMeasurements) SetStatus505(v int64)`

SetStatus505 sets Status505 field to given value.

### HasStatus505

`func (o *DomainInspectorMeasurements) HasStatus505() bool`

HasStatus505 returns a boolean if a field has been set.

### GetStatus530

`func (o *DomainInspectorMeasurements) GetStatus530() int64`

GetStatus530 returns the Status530 field if non-nil, zero value otherwise.

### GetStatus530Ok

`func (o *DomainInspectorMeasurements) GetStatus530Ok() (*int64, bool)`

GetStatus530Ok returns a tuple with the Status530 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus530

`func (o *DomainInspectorMeasurements) SetStatus530(v int64)`

SetStatus530 sets Status530 field to given value.

### HasStatus530

`func (o *DomainInspectorMeasurements) HasStatus530() bool`

HasStatus530 returns a boolean if a field has been set.

### GetRequests

`func (o *DomainInspectorMeasurements) GetRequests() int64`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *DomainInspectorMeasurements) GetRequestsOk() (*int64, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *DomainInspectorMeasurements) SetRequests(v int64)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *DomainInspectorMeasurements) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### GetRespHeaderBytes

`func (o *DomainInspectorMeasurements) GetRespHeaderBytes() int64`

GetRespHeaderBytes returns the RespHeaderBytes field if non-nil, zero value otherwise.

### GetRespHeaderBytesOk

`func (o *DomainInspectorMeasurements) GetRespHeaderBytesOk() (*int64, bool)`

GetRespHeaderBytesOk returns a tuple with the RespHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRespHeaderBytes

`func (o *DomainInspectorMeasurements) SetRespHeaderBytes(v int64)`

SetRespHeaderBytes sets RespHeaderBytes field to given value.

### HasRespHeaderBytes

`func (o *DomainInspectorMeasurements) HasRespHeaderBytes() bool`

HasRespHeaderBytes returns a boolean if a field has been set.

### GetRespBodyBytes

`func (o *DomainInspectorMeasurements) GetRespBodyBytes() int64`

GetRespBodyBytes returns the RespBodyBytes field if non-nil, zero value otherwise.

### GetRespBodyBytesOk

`func (o *DomainInspectorMeasurements) GetRespBodyBytesOk() (*int64, bool)`

GetRespBodyBytesOk returns a tuple with the RespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRespBodyBytes

`func (o *DomainInspectorMeasurements) SetRespBodyBytes(v int64)`

SetRespBodyBytes sets RespBodyBytes field to given value.

### HasRespBodyBytes

`func (o *DomainInspectorMeasurements) HasRespBodyBytes() bool`

HasRespBodyBytes returns a boolean if a field has been set.

### GetBereqHeaderBytes

`func (o *DomainInspectorMeasurements) GetBereqHeaderBytes() int64`

GetBereqHeaderBytes returns the BereqHeaderBytes field if non-nil, zero value otherwise.

### GetBereqHeaderBytesOk

`func (o *DomainInspectorMeasurements) GetBereqHeaderBytesOk() (*int64, bool)`

GetBereqHeaderBytesOk returns a tuple with the BereqHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBereqHeaderBytes

`func (o *DomainInspectorMeasurements) SetBereqHeaderBytes(v int64)`

SetBereqHeaderBytes sets BereqHeaderBytes field to given value.

### HasBereqHeaderBytes

`func (o *DomainInspectorMeasurements) HasBereqHeaderBytes() bool`

HasBereqHeaderBytes returns a boolean if a field has been set.

### GetBereqBodyBytes

`func (o *DomainInspectorMeasurements) GetBereqBodyBytes() int64`

GetBereqBodyBytes returns the BereqBodyBytes field if non-nil, zero value otherwise.

### GetBereqBodyBytesOk

`func (o *DomainInspectorMeasurements) GetBereqBodyBytesOk() (*int64, bool)`

GetBereqBodyBytesOk returns a tuple with the BereqBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBereqBodyBytes

`func (o *DomainInspectorMeasurements) SetBereqBodyBytes(v int64)`

SetBereqBodyBytes sets BereqBodyBytes field to given value.

### HasBereqBodyBytes

`func (o *DomainInspectorMeasurements) HasBereqBodyBytes() bool`

HasBereqBodyBytes returns a boolean if a field has been set.

### GetEdgeHitRequests

`func (o *DomainInspectorMeasurements) GetEdgeHitRequests() int64`

GetEdgeHitRequests returns the EdgeHitRequests field if non-nil, zero value otherwise.

### GetEdgeHitRequestsOk

`func (o *DomainInspectorMeasurements) GetEdgeHitRequestsOk() (*int64, bool)`

GetEdgeHitRequestsOk returns a tuple with the EdgeHitRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeHitRequests

`func (o *DomainInspectorMeasurements) SetEdgeHitRequests(v int64)`

SetEdgeHitRequests sets EdgeHitRequests field to given value.

### HasEdgeHitRequests

`func (o *DomainInspectorMeasurements) HasEdgeHitRequests() bool`

HasEdgeHitRequests returns a boolean if a field has been set.

### GetEdgeMissRequests

`func (o *DomainInspectorMeasurements) GetEdgeMissRequests() int64`

GetEdgeMissRequests returns the EdgeMissRequests field if non-nil, zero value otherwise.

### GetEdgeMissRequestsOk

`func (o *DomainInspectorMeasurements) GetEdgeMissRequestsOk() (*int64, bool)`

GetEdgeMissRequestsOk returns a tuple with the EdgeMissRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeMissRequests

`func (o *DomainInspectorMeasurements) SetEdgeMissRequests(v int64)`

SetEdgeMissRequests sets EdgeMissRequests field to given value.

### HasEdgeMissRequests

`func (o *DomainInspectorMeasurements) HasEdgeMissRequests() bool`

HasEdgeMissRequests returns a boolean if a field has been set.

### GetOriginFetches

`func (o *DomainInspectorMeasurements) GetOriginFetches() int64`

GetOriginFetches returns the OriginFetches field if non-nil, zero value otherwise.

### GetOriginFetchesOk

`func (o *DomainInspectorMeasurements) GetOriginFetchesOk() (*int64, bool)`

GetOriginFetchesOk returns a tuple with the OriginFetches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginFetches

`func (o *DomainInspectorMeasurements) SetOriginFetches(v int64)`

SetOriginFetches sets OriginFetches field to given value.

### HasOriginFetches

`func (o *DomainInspectorMeasurements) HasOriginFetches() bool`

HasOriginFetches returns a boolean if a field has been set.

### GetOriginFetchRespHeaderBytes

`func (o *DomainInspectorMeasurements) GetOriginFetchRespHeaderBytes() int64`

GetOriginFetchRespHeaderBytes returns the OriginFetchRespHeaderBytes field if non-nil, zero value otherwise.

### GetOriginFetchRespHeaderBytesOk

`func (o *DomainInspectorMeasurements) GetOriginFetchRespHeaderBytesOk() (*int64, bool)`

GetOriginFetchRespHeaderBytesOk returns a tuple with the OriginFetchRespHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginFetchRespHeaderBytes

`func (o *DomainInspectorMeasurements) SetOriginFetchRespHeaderBytes(v int64)`

SetOriginFetchRespHeaderBytes sets OriginFetchRespHeaderBytes field to given value.

### HasOriginFetchRespHeaderBytes

`func (o *DomainInspectorMeasurements) HasOriginFetchRespHeaderBytes() bool`

HasOriginFetchRespHeaderBytes returns a boolean if a field has been set.

### GetOriginFetchRespBodyBytes

`func (o *DomainInspectorMeasurements) GetOriginFetchRespBodyBytes() int64`

GetOriginFetchRespBodyBytes returns the OriginFetchRespBodyBytes field if non-nil, zero value otherwise.

### GetOriginFetchRespBodyBytesOk

`func (o *DomainInspectorMeasurements) GetOriginFetchRespBodyBytesOk() (*int64, bool)`

GetOriginFetchRespBodyBytesOk returns a tuple with the OriginFetchRespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginFetchRespBodyBytes

`func (o *DomainInspectorMeasurements) SetOriginFetchRespBodyBytes(v int64)`

SetOriginFetchRespBodyBytes sets OriginFetchRespBodyBytes field to given value.

### HasOriginFetchRespBodyBytes

`func (o *DomainInspectorMeasurements) HasOriginFetchRespBodyBytes() bool`

HasOriginFetchRespBodyBytes returns a boolean if a field has been set.

### GetBandwidth

`func (o *DomainInspectorMeasurements) GetBandwidth() int64`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *DomainInspectorMeasurements) GetBandwidthOk() (*int64, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *DomainInspectorMeasurements) SetBandwidth(v int64)`

SetBandwidth sets Bandwidth field to given value.

### HasBandwidth

`func (o *DomainInspectorMeasurements) HasBandwidth() bool`

HasBandwidth returns a boolean if a field has been set.

### GetEdgeHitRatio

`func (o *DomainInspectorMeasurements) GetEdgeHitRatio() float32`

GetEdgeHitRatio returns the EdgeHitRatio field if non-nil, zero value otherwise.

### GetEdgeHitRatioOk

`func (o *DomainInspectorMeasurements) GetEdgeHitRatioOk() (*float32, bool)`

GetEdgeHitRatioOk returns a tuple with the EdgeHitRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeHitRatio

`func (o *DomainInspectorMeasurements) SetEdgeHitRatio(v float32)`

SetEdgeHitRatio sets EdgeHitRatio field to given value.

### HasEdgeHitRatio

`func (o *DomainInspectorMeasurements) HasEdgeHitRatio() bool`

HasEdgeHitRatio returns a boolean if a field has been set.

### GetOriginOffload

`func (o *DomainInspectorMeasurements) GetOriginOffload() float32`

GetOriginOffload returns the OriginOffload field if non-nil, zero value otherwise.

### GetOriginOffloadOk

`func (o *DomainInspectorMeasurements) GetOriginOffloadOk() (*float32, bool)`

GetOriginOffloadOk returns a tuple with the OriginOffload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginOffload

`func (o *DomainInspectorMeasurements) SetOriginOffload(v float32)`

SetOriginOffload sets OriginOffload field to given value.

### HasOriginOffload

`func (o *DomainInspectorMeasurements) HasOriginOffload() bool`

HasOriginOffload returns a boolean if a field has been set.

### GetOriginStatus200

`func (o *DomainInspectorMeasurements) GetOriginStatus200() int64`

GetOriginStatus200 returns the OriginStatus200 field if non-nil, zero value otherwise.

### GetOriginStatus200Ok

`func (o *DomainInspectorMeasurements) GetOriginStatus200Ok() (*int64, bool)`

GetOriginStatus200Ok returns a tuple with the OriginStatus200 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginStatus200

`func (o *DomainInspectorMeasurements) SetOriginStatus200(v int64)`

SetOriginStatus200 sets OriginStatus200 field to given value.

### HasOriginStatus200

`func (o *DomainInspectorMeasurements) HasOriginStatus200() bool`

HasOriginStatus200 returns a boolean if a field has been set.

### GetOriginStatus204

`func (o *DomainInspectorMeasurements) GetOriginStatus204() int64`

GetOriginStatus204 returns the OriginStatus204 field if non-nil, zero value otherwise.

### GetOriginStatus204Ok

`func (o *DomainInspectorMeasurements) GetOriginStatus204Ok() (*int64, bool)`

GetOriginStatus204Ok returns a tuple with the OriginStatus204 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginStatus204

`func (o *DomainInspectorMeasurements) SetOriginStatus204(v int64)`

SetOriginStatus204 sets OriginStatus204 field to given value.

### HasOriginStatus204

`func (o *DomainInspectorMeasurements) HasOriginStatus204() bool`

HasOriginStatus204 returns a boolean if a field has been set.

### GetOriginStatus206

`func (o *DomainInspectorMeasurements) GetOriginStatus206() int64`

GetOriginStatus206 returns the OriginStatus206 field if non-nil, zero value otherwise.

### GetOriginStatus206Ok

`func (o *DomainInspectorMeasurements) GetOriginStatus206Ok() (*int64, bool)`

GetOriginStatus206Ok returns a tuple with the OriginStatus206 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginStatus206

`func (o *DomainInspectorMeasurements) SetOriginStatus206(v int64)`

SetOriginStatus206 sets OriginStatus206 field to given value.

### HasOriginStatus206

`func (o *DomainInspectorMeasurements) HasOriginStatus206() bool`

HasOriginStatus206 returns a boolean if a field has been set.

### GetOriginStatus301

`func (o *DomainInspectorMeasurements) GetOriginStatus301() int64`

GetOriginStatus301 returns the OriginStatus301 field if non-nil, zero value otherwise.

### GetOriginStatus301Ok

`func (o *DomainInspectorMeasurements) GetOriginStatus301Ok() (*int64, bool)`

GetOriginStatus301Ok returns a tuple with the OriginStatus301 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginStatus301

`func (o *DomainInspectorMeasurements) SetOriginStatus301(v int64)`

SetOriginStatus301 sets OriginStatus301 field to given value.

### HasOriginStatus301

`func (o *DomainInspectorMeasurements) HasOriginStatus301() bool`

HasOriginStatus301 returns a boolean if a field has been set.

### GetOriginStatus302

`func (o *DomainInspectorMeasurements) GetOriginStatus302() int64`

GetOriginStatus302 returns the OriginStatus302 field if non-nil, zero value otherwise.

### GetOriginStatus302Ok

`func (o *DomainInspectorMeasurements) GetOriginStatus302Ok() (*int64, bool)`

GetOriginStatus302Ok returns a tuple with the OriginStatus302 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginStatus302

`func (o *DomainInspectorMeasurements) SetOriginStatus302(v int64)`

SetOriginStatus302 sets OriginStatus302 field to given value.

### HasOriginStatus302

`func (o *DomainInspectorMeasurements) HasOriginStatus302() bool`

HasOriginStatus302 returns a boolean if a field has been set.

### GetOriginStatus304

`func (o *DomainInspectorMeasurements) GetOriginStatus304() int64`

GetOriginStatus304 returns the OriginStatus304 field if non-nil, zero value otherwise.

### GetOriginStatus304Ok

`func (o *DomainInspectorMeasurements) GetOriginStatus304Ok() (*int64, bool)`

GetOriginStatus304Ok returns a tuple with the OriginStatus304 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginStatus304

`func (o *DomainInspectorMeasurements) SetOriginStatus304(v int64)`

SetOriginStatus304 sets OriginStatus304 field to given value.

### HasOriginStatus304

`func (o *DomainInspectorMeasurements) HasOriginStatus304() bool`

HasOriginStatus304 returns a boolean if a field has been set.

### GetOriginStatus400

`func (o *DomainInspectorMeasurements) GetOriginStatus400() int64`

GetOriginStatus400 returns the OriginStatus400 field if non-nil, zero value otherwise.

### GetOriginStatus400Ok

`func (o *DomainInspectorMeasurements) GetOriginStatus400Ok() (*int64, bool)`

GetOriginStatus400Ok returns a tuple with the OriginStatus400 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginStatus400

`func (o *DomainInspectorMeasurements) SetOriginStatus400(v int64)`

SetOriginStatus400 sets OriginStatus400 field to given value.

### HasOriginStatus400

`func (o *DomainInspectorMeasurements) HasOriginStatus400() bool`

HasOriginStatus400 returns a boolean if a field has been set.

### GetOriginStatus401

`func (o *DomainInspectorMeasurements) GetOriginStatus401() int64`

GetOriginStatus401 returns the OriginStatus401 field if non-nil, zero value otherwise.

### GetOriginStatus401Ok

`func (o *DomainInspectorMeasurements) GetOriginStatus401Ok() (*int64, bool)`

GetOriginStatus401Ok returns a tuple with the OriginStatus401 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginStatus401

`func (o *DomainInspectorMeasurements) SetOriginStatus401(v int64)`

SetOriginStatus401 sets OriginStatus401 field to given value.

### HasOriginStatus401

`func (o *DomainInspectorMeasurements) HasOriginStatus401() bool`

HasOriginStatus401 returns a boolean if a field has been set.

### GetOriginStatus403

`func (o *DomainInspectorMeasurements) GetOriginStatus403() int64`

GetOriginStatus403 returns the OriginStatus403 field if non-nil, zero value otherwise.

### GetOriginStatus403Ok

`func (o *DomainInspectorMeasurements) GetOriginStatus403Ok() (*int64, bool)`

GetOriginStatus403Ok returns a tuple with the OriginStatus403 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginStatus403

`func (o *DomainInspectorMeasurements) SetOriginStatus403(v int64)`

SetOriginStatus403 sets OriginStatus403 field to given value.

### HasOriginStatus403

`func (o *DomainInspectorMeasurements) HasOriginStatus403() bool`

HasOriginStatus403 returns a boolean if a field has been set.

### GetOriginStatus404

`func (o *DomainInspectorMeasurements) GetOriginStatus404() int64`

GetOriginStatus404 returns the OriginStatus404 field if non-nil, zero value otherwise.

### GetOriginStatus404Ok

`func (o *DomainInspectorMeasurements) GetOriginStatus404Ok() (*int64, bool)`

GetOriginStatus404Ok returns a tuple with the OriginStatus404 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginStatus404

`func (o *DomainInspectorMeasurements) SetOriginStatus404(v int64)`

SetOriginStatus404 sets OriginStatus404 field to given value.

### HasOriginStatus404

`func (o *DomainInspectorMeasurements) HasOriginStatus404() bool`

HasOriginStatus404 returns a boolean if a field has been set.

### GetOriginStatus416

`func (o *DomainInspectorMeasurements) GetOriginStatus416() int64`

GetOriginStatus416 returns the OriginStatus416 field if non-nil, zero value otherwise.

### GetOriginStatus416Ok

`func (o *DomainInspectorMeasurements) GetOriginStatus416Ok() (*int64, bool)`

GetOriginStatus416Ok returns a tuple with the OriginStatus416 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginStatus416

`func (o *DomainInspectorMeasurements) SetOriginStatus416(v int64)`

SetOriginStatus416 sets OriginStatus416 field to given value.

### HasOriginStatus416

`func (o *DomainInspectorMeasurements) HasOriginStatus416() bool`

HasOriginStatus416 returns a boolean if a field has been set.

### GetOriginStatus429

`func (o *DomainInspectorMeasurements) GetOriginStatus429() int64`

GetOriginStatus429 returns the OriginStatus429 field if non-nil, zero value otherwise.

### GetOriginStatus429Ok

`func (o *DomainInspectorMeasurements) GetOriginStatus429Ok() (*int64, bool)`

GetOriginStatus429Ok returns a tuple with the OriginStatus429 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginStatus429

`func (o *DomainInspectorMeasurements) SetOriginStatus429(v int64)`

SetOriginStatus429 sets OriginStatus429 field to given value.

### HasOriginStatus429

`func (o *DomainInspectorMeasurements) HasOriginStatus429() bool`

HasOriginStatus429 returns a boolean if a field has been set.

### GetOriginStatus500

`func (o *DomainInspectorMeasurements) GetOriginStatus500() int64`

GetOriginStatus500 returns the OriginStatus500 field if non-nil, zero value otherwise.

### GetOriginStatus500Ok

`func (o *DomainInspectorMeasurements) GetOriginStatus500Ok() (*int64, bool)`

GetOriginStatus500Ok returns a tuple with the OriginStatus500 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginStatus500

`func (o *DomainInspectorMeasurements) SetOriginStatus500(v int64)`

SetOriginStatus500 sets OriginStatus500 field to given value.

### HasOriginStatus500

`func (o *DomainInspectorMeasurements) HasOriginStatus500() bool`

HasOriginStatus500 returns a boolean if a field has been set.

### GetOriginStatus501

`func (o *DomainInspectorMeasurements) GetOriginStatus501() int64`

GetOriginStatus501 returns the OriginStatus501 field if non-nil, zero value otherwise.

### GetOriginStatus501Ok

`func (o *DomainInspectorMeasurements) GetOriginStatus501Ok() (*int64, bool)`

GetOriginStatus501Ok returns a tuple with the OriginStatus501 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginStatus501

`func (o *DomainInspectorMeasurements) SetOriginStatus501(v int64)`

SetOriginStatus501 sets OriginStatus501 field to given value.

### HasOriginStatus501

`func (o *DomainInspectorMeasurements) HasOriginStatus501() bool`

HasOriginStatus501 returns a boolean if a field has been set.

### GetOriginStatus502

`func (o *DomainInspectorMeasurements) GetOriginStatus502() int64`

GetOriginStatus502 returns the OriginStatus502 field if non-nil, zero value otherwise.

### GetOriginStatus502Ok

`func (o *DomainInspectorMeasurements) GetOriginStatus502Ok() (*int64, bool)`

GetOriginStatus502Ok returns a tuple with the OriginStatus502 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginStatus502

`func (o *DomainInspectorMeasurements) SetOriginStatus502(v int64)`

SetOriginStatus502 sets OriginStatus502 field to given value.

### HasOriginStatus502

`func (o *DomainInspectorMeasurements) HasOriginStatus502() bool`

HasOriginStatus502 returns a boolean if a field has been set.

### GetOriginStatus503

`func (o *DomainInspectorMeasurements) GetOriginStatus503() int64`

GetOriginStatus503 returns the OriginStatus503 field if non-nil, zero value otherwise.

### GetOriginStatus503Ok

`func (o *DomainInspectorMeasurements) GetOriginStatus503Ok() (*int64, bool)`

GetOriginStatus503Ok returns a tuple with the OriginStatus503 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginStatus503

`func (o *DomainInspectorMeasurements) SetOriginStatus503(v int64)`

SetOriginStatus503 sets OriginStatus503 field to given value.

### HasOriginStatus503

`func (o *DomainInspectorMeasurements) HasOriginStatus503() bool`

HasOriginStatus503 returns a boolean if a field has been set.

### GetOriginStatus504

`func (o *DomainInspectorMeasurements) GetOriginStatus504() int64`

GetOriginStatus504 returns the OriginStatus504 field if non-nil, zero value otherwise.

### GetOriginStatus504Ok

`func (o *DomainInspectorMeasurements) GetOriginStatus504Ok() (*int64, bool)`

GetOriginStatus504Ok returns a tuple with the OriginStatus504 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginStatus504

`func (o *DomainInspectorMeasurements) SetOriginStatus504(v int64)`

SetOriginStatus504 sets OriginStatus504 field to given value.

### HasOriginStatus504

`func (o *DomainInspectorMeasurements) HasOriginStatus504() bool`

HasOriginStatus504 returns a boolean if a field has been set.

### GetOriginStatus505

`func (o *DomainInspectorMeasurements) GetOriginStatus505() int64`

GetOriginStatus505 returns the OriginStatus505 field if non-nil, zero value otherwise.

### GetOriginStatus505Ok

`func (o *DomainInspectorMeasurements) GetOriginStatus505Ok() (*int64, bool)`

GetOriginStatus505Ok returns a tuple with the OriginStatus505 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginStatus505

`func (o *DomainInspectorMeasurements) SetOriginStatus505(v int64)`

SetOriginStatus505 sets OriginStatus505 field to given value.

### HasOriginStatus505

`func (o *DomainInspectorMeasurements) HasOriginStatus505() bool`

HasOriginStatus505 returns a boolean if a field has been set.

### GetOriginStatus530

`func (o *DomainInspectorMeasurements) GetOriginStatus530() int64`

GetOriginStatus530 returns the OriginStatus530 field if non-nil, zero value otherwise.

### GetOriginStatus530Ok

`func (o *DomainInspectorMeasurements) GetOriginStatus530Ok() (*int64, bool)`

GetOriginStatus530Ok returns a tuple with the OriginStatus530 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginStatus530

`func (o *DomainInspectorMeasurements) SetOriginStatus530(v int64)`

SetOriginStatus530 sets OriginStatus530 field to given value.

### HasOriginStatus530

`func (o *DomainInspectorMeasurements) HasOriginStatus530() bool`

HasOriginStatus530 returns a boolean if a field has been set.

### GetOriginStatus1xx

`func (o *DomainInspectorMeasurements) GetOriginStatus1xx() int64`

GetOriginStatus1xx returns the OriginStatus1xx field if non-nil, zero value otherwise.

### GetOriginStatus1xxOk

`func (o *DomainInspectorMeasurements) GetOriginStatus1xxOk() (*int64, bool)`

GetOriginStatus1xxOk returns a tuple with the OriginStatus1xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginStatus1xx

`func (o *DomainInspectorMeasurements) SetOriginStatus1xx(v int64)`

SetOriginStatus1xx sets OriginStatus1xx field to given value.

### HasOriginStatus1xx

`func (o *DomainInspectorMeasurements) HasOriginStatus1xx() bool`

HasOriginStatus1xx returns a boolean if a field has been set.

### GetOriginStatus2xx

`func (o *DomainInspectorMeasurements) GetOriginStatus2xx() int64`

GetOriginStatus2xx returns the OriginStatus2xx field if non-nil, zero value otherwise.

### GetOriginStatus2xxOk

`func (o *DomainInspectorMeasurements) GetOriginStatus2xxOk() (*int64, bool)`

GetOriginStatus2xxOk returns a tuple with the OriginStatus2xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginStatus2xx

`func (o *DomainInspectorMeasurements) SetOriginStatus2xx(v int64)`

SetOriginStatus2xx sets OriginStatus2xx field to given value.

### HasOriginStatus2xx

`func (o *DomainInspectorMeasurements) HasOriginStatus2xx() bool`

HasOriginStatus2xx returns a boolean if a field has been set.

### GetOriginStatus3xx

`func (o *DomainInspectorMeasurements) GetOriginStatus3xx() int64`

GetOriginStatus3xx returns the OriginStatus3xx field if non-nil, zero value otherwise.

### GetOriginStatus3xxOk

`func (o *DomainInspectorMeasurements) GetOriginStatus3xxOk() (*int64, bool)`

GetOriginStatus3xxOk returns a tuple with the OriginStatus3xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginStatus3xx

`func (o *DomainInspectorMeasurements) SetOriginStatus3xx(v int64)`

SetOriginStatus3xx sets OriginStatus3xx field to given value.

### HasOriginStatus3xx

`func (o *DomainInspectorMeasurements) HasOriginStatus3xx() bool`

HasOriginStatus3xx returns a boolean if a field has been set.

### GetOriginStatus4xx

`func (o *DomainInspectorMeasurements) GetOriginStatus4xx() int64`

GetOriginStatus4xx returns the OriginStatus4xx field if non-nil, zero value otherwise.

### GetOriginStatus4xxOk

`func (o *DomainInspectorMeasurements) GetOriginStatus4xxOk() (*int64, bool)`

GetOriginStatus4xxOk returns a tuple with the OriginStatus4xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginStatus4xx

`func (o *DomainInspectorMeasurements) SetOriginStatus4xx(v int64)`

SetOriginStatus4xx sets OriginStatus4xx field to given value.

### HasOriginStatus4xx

`func (o *DomainInspectorMeasurements) HasOriginStatus4xx() bool`

HasOriginStatus4xx returns a boolean if a field has been set.

### GetOriginStatus5xx

`func (o *DomainInspectorMeasurements) GetOriginStatus5xx() int64`

GetOriginStatus5xx returns the OriginStatus5xx field if non-nil, zero value otherwise.

### GetOriginStatus5xxOk

`func (o *DomainInspectorMeasurements) GetOriginStatus5xxOk() (*int64, bool)`

GetOriginStatus5xxOk returns a tuple with the OriginStatus5xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginStatus5xx

`func (o *DomainInspectorMeasurements) SetOriginStatus5xx(v int64)`

SetOriginStatus5xx sets OriginStatus5xx field to given value.

### HasOriginStatus5xx

`func (o *DomainInspectorMeasurements) HasOriginStatus5xx() bool`

HasOriginStatus5xx returns a boolean if a field has been set.

### GetComputeBereqBodyBytes

`func (o *DomainInspectorMeasurements) GetComputeBereqBodyBytes() int64`

GetComputeBereqBodyBytes returns the ComputeBereqBodyBytes field if non-nil, zero value otherwise.

### GetComputeBereqBodyBytesOk

`func (o *DomainInspectorMeasurements) GetComputeBereqBodyBytesOk() (*int64, bool)`

GetComputeBereqBodyBytesOk returns a tuple with the ComputeBereqBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBereqBodyBytes

`func (o *DomainInspectorMeasurements) SetComputeBereqBodyBytes(v int64)`

SetComputeBereqBodyBytes sets ComputeBereqBodyBytes field to given value.

### HasComputeBereqBodyBytes

`func (o *DomainInspectorMeasurements) HasComputeBereqBodyBytes() bool`

HasComputeBereqBodyBytes returns a boolean if a field has been set.

### GetComputeBereqErrors

`func (o *DomainInspectorMeasurements) GetComputeBereqErrors() int64`

GetComputeBereqErrors returns the ComputeBereqErrors field if non-nil, zero value otherwise.

### GetComputeBereqErrorsOk

`func (o *DomainInspectorMeasurements) GetComputeBereqErrorsOk() (*int64, bool)`

GetComputeBereqErrorsOk returns a tuple with the ComputeBereqErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBereqErrors

`func (o *DomainInspectorMeasurements) SetComputeBereqErrors(v int64)`

SetComputeBereqErrors sets ComputeBereqErrors field to given value.

### HasComputeBereqErrors

`func (o *DomainInspectorMeasurements) HasComputeBereqErrors() bool`

HasComputeBereqErrors returns a boolean if a field has been set.

### GetComputeBereqHeaderBytes

`func (o *DomainInspectorMeasurements) GetComputeBereqHeaderBytes() int64`

GetComputeBereqHeaderBytes returns the ComputeBereqHeaderBytes field if non-nil, zero value otherwise.

### GetComputeBereqHeaderBytesOk

`func (o *DomainInspectorMeasurements) GetComputeBereqHeaderBytesOk() (*int64, bool)`

GetComputeBereqHeaderBytesOk returns a tuple with the ComputeBereqHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBereqHeaderBytes

`func (o *DomainInspectorMeasurements) SetComputeBereqHeaderBytes(v int64)`

SetComputeBereqHeaderBytes sets ComputeBereqHeaderBytes field to given value.

### HasComputeBereqHeaderBytes

`func (o *DomainInspectorMeasurements) HasComputeBereqHeaderBytes() bool`

HasComputeBereqHeaderBytes returns a boolean if a field has been set.

### GetComputeBereqs

`func (o *DomainInspectorMeasurements) GetComputeBereqs() int64`

GetComputeBereqs returns the ComputeBereqs field if non-nil, zero value otherwise.

### GetComputeBereqsOk

`func (o *DomainInspectorMeasurements) GetComputeBereqsOk() (*int64, bool)`

GetComputeBereqsOk returns a tuple with the ComputeBereqs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBereqs

`func (o *DomainInspectorMeasurements) SetComputeBereqs(v int64)`

SetComputeBereqs sets ComputeBereqs field to given value.

### HasComputeBereqs

`func (o *DomainInspectorMeasurements) HasComputeBereqs() bool`

HasComputeBereqs returns a boolean if a field has been set.

### GetComputeBerespBodyBytes

`func (o *DomainInspectorMeasurements) GetComputeBerespBodyBytes() int64`

GetComputeBerespBodyBytes returns the ComputeBerespBodyBytes field if non-nil, zero value otherwise.

### GetComputeBerespBodyBytesOk

`func (o *DomainInspectorMeasurements) GetComputeBerespBodyBytesOk() (*int64, bool)`

GetComputeBerespBodyBytesOk returns a tuple with the ComputeBerespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBerespBodyBytes

`func (o *DomainInspectorMeasurements) SetComputeBerespBodyBytes(v int64)`

SetComputeBerespBodyBytes sets ComputeBerespBodyBytes field to given value.

### HasComputeBerespBodyBytes

`func (o *DomainInspectorMeasurements) HasComputeBerespBodyBytes() bool`

HasComputeBerespBodyBytes returns a boolean if a field has been set.

### GetComputeBerespHeaderBytes

`func (o *DomainInspectorMeasurements) GetComputeBerespHeaderBytes() int64`

GetComputeBerespHeaderBytes returns the ComputeBerespHeaderBytes field if non-nil, zero value otherwise.

### GetComputeBerespHeaderBytesOk

`func (o *DomainInspectorMeasurements) GetComputeBerespHeaderBytesOk() (*int64, bool)`

GetComputeBerespHeaderBytesOk returns a tuple with the ComputeBerespHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBerespHeaderBytes

`func (o *DomainInspectorMeasurements) SetComputeBerespHeaderBytes(v int64)`

SetComputeBerespHeaderBytes sets ComputeBerespHeaderBytes field to given value.

### HasComputeBerespHeaderBytes

`func (o *DomainInspectorMeasurements) HasComputeBerespHeaderBytes() bool`

HasComputeBerespHeaderBytes returns a boolean if a field has been set.

### GetComputeExecutionTimeMs

`func (o *DomainInspectorMeasurements) GetComputeExecutionTimeMs() int64`

GetComputeExecutionTimeMs returns the ComputeExecutionTimeMs field if non-nil, zero value otherwise.

### GetComputeExecutionTimeMsOk

`func (o *DomainInspectorMeasurements) GetComputeExecutionTimeMsOk() (*int64, bool)`

GetComputeExecutionTimeMsOk returns a tuple with the ComputeExecutionTimeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeExecutionTimeMs

`func (o *DomainInspectorMeasurements) SetComputeExecutionTimeMs(v int64)`

SetComputeExecutionTimeMs sets ComputeExecutionTimeMs field to given value.

### HasComputeExecutionTimeMs

`func (o *DomainInspectorMeasurements) HasComputeExecutionTimeMs() bool`

HasComputeExecutionTimeMs returns a boolean if a field has been set.

### GetComputeOriginStatus1xx

`func (o *DomainInspectorMeasurements) GetComputeOriginStatus1xx() int64`

GetComputeOriginStatus1xx returns the ComputeOriginStatus1xx field if non-nil, zero value otherwise.

### GetComputeOriginStatus1xxOk

`func (o *DomainInspectorMeasurements) GetComputeOriginStatus1xxOk() (*int64, bool)`

GetComputeOriginStatus1xxOk returns a tuple with the ComputeOriginStatus1xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeOriginStatus1xx

`func (o *DomainInspectorMeasurements) SetComputeOriginStatus1xx(v int64)`

SetComputeOriginStatus1xx sets ComputeOriginStatus1xx field to given value.

### HasComputeOriginStatus1xx

`func (o *DomainInspectorMeasurements) HasComputeOriginStatus1xx() bool`

HasComputeOriginStatus1xx returns a boolean if a field has been set.

### GetComputeOriginStatus200

`func (o *DomainInspectorMeasurements) GetComputeOriginStatus200() int64`

GetComputeOriginStatus200 returns the ComputeOriginStatus200 field if non-nil, zero value otherwise.

### GetComputeOriginStatus200Ok

`func (o *DomainInspectorMeasurements) GetComputeOriginStatus200Ok() (*int64, bool)`

GetComputeOriginStatus200Ok returns a tuple with the ComputeOriginStatus200 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeOriginStatus200

`func (o *DomainInspectorMeasurements) SetComputeOriginStatus200(v int64)`

SetComputeOriginStatus200 sets ComputeOriginStatus200 field to given value.

### HasComputeOriginStatus200

`func (o *DomainInspectorMeasurements) HasComputeOriginStatus200() bool`

HasComputeOriginStatus200 returns a boolean if a field has been set.

### GetComputeOriginStatus204

`func (o *DomainInspectorMeasurements) GetComputeOriginStatus204() int64`

GetComputeOriginStatus204 returns the ComputeOriginStatus204 field if non-nil, zero value otherwise.

### GetComputeOriginStatus204Ok

`func (o *DomainInspectorMeasurements) GetComputeOriginStatus204Ok() (*int64, bool)`

GetComputeOriginStatus204Ok returns a tuple with the ComputeOriginStatus204 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeOriginStatus204

`func (o *DomainInspectorMeasurements) SetComputeOriginStatus204(v int64)`

SetComputeOriginStatus204 sets ComputeOriginStatus204 field to given value.

### HasComputeOriginStatus204

`func (o *DomainInspectorMeasurements) HasComputeOriginStatus204() bool`

HasComputeOriginStatus204 returns a boolean if a field has been set.

### GetComputeOriginStatus206

`func (o *DomainInspectorMeasurements) GetComputeOriginStatus206() int64`

GetComputeOriginStatus206 returns the ComputeOriginStatus206 field if non-nil, zero value otherwise.

### GetComputeOriginStatus206Ok

`func (o *DomainInspectorMeasurements) GetComputeOriginStatus206Ok() (*int64, bool)`

GetComputeOriginStatus206Ok returns a tuple with the ComputeOriginStatus206 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeOriginStatus206

`func (o *DomainInspectorMeasurements) SetComputeOriginStatus206(v int64)`

SetComputeOriginStatus206 sets ComputeOriginStatus206 field to given value.

### HasComputeOriginStatus206

`func (o *DomainInspectorMeasurements) HasComputeOriginStatus206() bool`

HasComputeOriginStatus206 returns a boolean if a field has been set.

### GetComputeOriginStatus2xx

`func (o *DomainInspectorMeasurements) GetComputeOriginStatus2xx() int64`

GetComputeOriginStatus2xx returns the ComputeOriginStatus2xx field if non-nil, zero value otherwise.

### GetComputeOriginStatus2xxOk

`func (o *DomainInspectorMeasurements) GetComputeOriginStatus2xxOk() (*int64, bool)`

GetComputeOriginStatus2xxOk returns a tuple with the ComputeOriginStatus2xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeOriginStatus2xx

`func (o *DomainInspectorMeasurements) SetComputeOriginStatus2xx(v int64)`

SetComputeOriginStatus2xx sets ComputeOriginStatus2xx field to given value.

### HasComputeOriginStatus2xx

`func (o *DomainInspectorMeasurements) HasComputeOriginStatus2xx() bool`

HasComputeOriginStatus2xx returns a boolean if a field has been set.

### GetComputeOriginStatus301

`func (o *DomainInspectorMeasurements) GetComputeOriginStatus301() int64`

GetComputeOriginStatus301 returns the ComputeOriginStatus301 field if non-nil, zero value otherwise.

### GetComputeOriginStatus301Ok

`func (o *DomainInspectorMeasurements) GetComputeOriginStatus301Ok() (*int64, bool)`

GetComputeOriginStatus301Ok returns a tuple with the ComputeOriginStatus301 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeOriginStatus301

`func (o *DomainInspectorMeasurements) SetComputeOriginStatus301(v int64)`

SetComputeOriginStatus301 sets ComputeOriginStatus301 field to given value.

### HasComputeOriginStatus301

`func (o *DomainInspectorMeasurements) HasComputeOriginStatus301() bool`

HasComputeOriginStatus301 returns a boolean if a field has been set.

### GetComputeOriginStatus302

`func (o *DomainInspectorMeasurements) GetComputeOriginStatus302() int64`

GetComputeOriginStatus302 returns the ComputeOriginStatus302 field if non-nil, zero value otherwise.

### GetComputeOriginStatus302Ok

`func (o *DomainInspectorMeasurements) GetComputeOriginStatus302Ok() (*int64, bool)`

GetComputeOriginStatus302Ok returns a tuple with the ComputeOriginStatus302 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeOriginStatus302

`func (o *DomainInspectorMeasurements) SetComputeOriginStatus302(v int64)`

SetComputeOriginStatus302 sets ComputeOriginStatus302 field to given value.

### HasComputeOriginStatus302

`func (o *DomainInspectorMeasurements) HasComputeOriginStatus302() bool`

HasComputeOriginStatus302 returns a boolean if a field has been set.

### GetComputeOriginStatus304

`func (o *DomainInspectorMeasurements) GetComputeOriginStatus304() int64`

GetComputeOriginStatus304 returns the ComputeOriginStatus304 field if non-nil, zero value otherwise.

### GetComputeOriginStatus304Ok

`func (o *DomainInspectorMeasurements) GetComputeOriginStatus304Ok() (*int64, bool)`

GetComputeOriginStatus304Ok returns a tuple with the ComputeOriginStatus304 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeOriginStatus304

`func (o *DomainInspectorMeasurements) SetComputeOriginStatus304(v int64)`

SetComputeOriginStatus304 sets ComputeOriginStatus304 field to given value.

### HasComputeOriginStatus304

`func (o *DomainInspectorMeasurements) HasComputeOriginStatus304() bool`

HasComputeOriginStatus304 returns a boolean if a field has been set.

### GetComputeOriginStatus3xx

`func (o *DomainInspectorMeasurements) GetComputeOriginStatus3xx() int64`

GetComputeOriginStatus3xx returns the ComputeOriginStatus3xx field if non-nil, zero value otherwise.

### GetComputeOriginStatus3xxOk

`func (o *DomainInspectorMeasurements) GetComputeOriginStatus3xxOk() (*int64, bool)`

GetComputeOriginStatus3xxOk returns a tuple with the ComputeOriginStatus3xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeOriginStatus3xx

`func (o *DomainInspectorMeasurements) SetComputeOriginStatus3xx(v int64)`

SetComputeOriginStatus3xx sets ComputeOriginStatus3xx field to given value.

### HasComputeOriginStatus3xx

`func (o *DomainInspectorMeasurements) HasComputeOriginStatus3xx() bool`

HasComputeOriginStatus3xx returns a boolean if a field has been set.

### GetComputeOriginStatus400

`func (o *DomainInspectorMeasurements) GetComputeOriginStatus400() int64`

GetComputeOriginStatus400 returns the ComputeOriginStatus400 field if non-nil, zero value otherwise.

### GetComputeOriginStatus400Ok

`func (o *DomainInspectorMeasurements) GetComputeOriginStatus400Ok() (*int64, bool)`

GetComputeOriginStatus400Ok returns a tuple with the ComputeOriginStatus400 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeOriginStatus400

`func (o *DomainInspectorMeasurements) SetComputeOriginStatus400(v int64)`

SetComputeOriginStatus400 sets ComputeOriginStatus400 field to given value.

### HasComputeOriginStatus400

`func (o *DomainInspectorMeasurements) HasComputeOriginStatus400() bool`

HasComputeOriginStatus400 returns a boolean if a field has been set.

### GetComputeOriginStatus401

`func (o *DomainInspectorMeasurements) GetComputeOriginStatus401() int64`

GetComputeOriginStatus401 returns the ComputeOriginStatus401 field if non-nil, zero value otherwise.

### GetComputeOriginStatus401Ok

`func (o *DomainInspectorMeasurements) GetComputeOriginStatus401Ok() (*int64, bool)`

GetComputeOriginStatus401Ok returns a tuple with the ComputeOriginStatus401 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeOriginStatus401

`func (o *DomainInspectorMeasurements) SetComputeOriginStatus401(v int64)`

SetComputeOriginStatus401 sets ComputeOriginStatus401 field to given value.

### HasComputeOriginStatus401

`func (o *DomainInspectorMeasurements) HasComputeOriginStatus401() bool`

HasComputeOriginStatus401 returns a boolean if a field has been set.

### GetComputeOriginStatus403

`func (o *DomainInspectorMeasurements) GetComputeOriginStatus403() int64`

GetComputeOriginStatus403 returns the ComputeOriginStatus403 field if non-nil, zero value otherwise.

### GetComputeOriginStatus403Ok

`func (o *DomainInspectorMeasurements) GetComputeOriginStatus403Ok() (*int64, bool)`

GetComputeOriginStatus403Ok returns a tuple with the ComputeOriginStatus403 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeOriginStatus403

`func (o *DomainInspectorMeasurements) SetComputeOriginStatus403(v int64)`

SetComputeOriginStatus403 sets ComputeOriginStatus403 field to given value.

### HasComputeOriginStatus403

`func (o *DomainInspectorMeasurements) HasComputeOriginStatus403() bool`

HasComputeOriginStatus403 returns a boolean if a field has been set.

### GetComputeOriginStatus404

`func (o *DomainInspectorMeasurements) GetComputeOriginStatus404() int64`

GetComputeOriginStatus404 returns the ComputeOriginStatus404 field if non-nil, zero value otherwise.

### GetComputeOriginStatus404Ok

`func (o *DomainInspectorMeasurements) GetComputeOriginStatus404Ok() (*int64, bool)`

GetComputeOriginStatus404Ok returns a tuple with the ComputeOriginStatus404 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeOriginStatus404

`func (o *DomainInspectorMeasurements) SetComputeOriginStatus404(v int64)`

SetComputeOriginStatus404 sets ComputeOriginStatus404 field to given value.

### HasComputeOriginStatus404

`func (o *DomainInspectorMeasurements) HasComputeOriginStatus404() bool`

HasComputeOriginStatus404 returns a boolean if a field has been set.

### GetComputeOriginStatus416

`func (o *DomainInspectorMeasurements) GetComputeOriginStatus416() int64`

GetComputeOriginStatus416 returns the ComputeOriginStatus416 field if non-nil, zero value otherwise.

### GetComputeOriginStatus416Ok

`func (o *DomainInspectorMeasurements) GetComputeOriginStatus416Ok() (*int64, bool)`

GetComputeOriginStatus416Ok returns a tuple with the ComputeOriginStatus416 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeOriginStatus416

`func (o *DomainInspectorMeasurements) SetComputeOriginStatus416(v int64)`

SetComputeOriginStatus416 sets ComputeOriginStatus416 field to given value.

### HasComputeOriginStatus416

`func (o *DomainInspectorMeasurements) HasComputeOriginStatus416() bool`

HasComputeOriginStatus416 returns a boolean if a field has been set.

### GetComputeOriginStatus429

`func (o *DomainInspectorMeasurements) GetComputeOriginStatus429() int64`

GetComputeOriginStatus429 returns the ComputeOriginStatus429 field if non-nil, zero value otherwise.

### GetComputeOriginStatus429Ok

`func (o *DomainInspectorMeasurements) GetComputeOriginStatus429Ok() (*int64, bool)`

GetComputeOriginStatus429Ok returns a tuple with the ComputeOriginStatus429 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeOriginStatus429

`func (o *DomainInspectorMeasurements) SetComputeOriginStatus429(v int64)`

SetComputeOriginStatus429 sets ComputeOriginStatus429 field to given value.

### HasComputeOriginStatus429

`func (o *DomainInspectorMeasurements) HasComputeOriginStatus429() bool`

HasComputeOriginStatus429 returns a boolean if a field has been set.

### GetComputeOriginStatus4xx

`func (o *DomainInspectorMeasurements) GetComputeOriginStatus4xx() int64`

GetComputeOriginStatus4xx returns the ComputeOriginStatus4xx field if non-nil, zero value otherwise.

### GetComputeOriginStatus4xxOk

`func (o *DomainInspectorMeasurements) GetComputeOriginStatus4xxOk() (*int64, bool)`

GetComputeOriginStatus4xxOk returns a tuple with the ComputeOriginStatus4xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeOriginStatus4xx

`func (o *DomainInspectorMeasurements) SetComputeOriginStatus4xx(v int64)`

SetComputeOriginStatus4xx sets ComputeOriginStatus4xx field to given value.

### HasComputeOriginStatus4xx

`func (o *DomainInspectorMeasurements) HasComputeOriginStatus4xx() bool`

HasComputeOriginStatus4xx returns a boolean if a field has been set.

### GetComputeOriginStatus500

`func (o *DomainInspectorMeasurements) GetComputeOriginStatus500() int64`

GetComputeOriginStatus500 returns the ComputeOriginStatus500 field if non-nil, zero value otherwise.

### GetComputeOriginStatus500Ok

`func (o *DomainInspectorMeasurements) GetComputeOriginStatus500Ok() (*int64, bool)`

GetComputeOriginStatus500Ok returns a tuple with the ComputeOriginStatus500 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeOriginStatus500

`func (o *DomainInspectorMeasurements) SetComputeOriginStatus500(v int64)`

SetComputeOriginStatus500 sets ComputeOriginStatus500 field to given value.

### HasComputeOriginStatus500

`func (o *DomainInspectorMeasurements) HasComputeOriginStatus500() bool`

HasComputeOriginStatus500 returns a boolean if a field has been set.

### GetComputeOriginStatus501

`func (o *DomainInspectorMeasurements) GetComputeOriginStatus501() int64`

GetComputeOriginStatus501 returns the ComputeOriginStatus501 field if non-nil, zero value otherwise.

### GetComputeOriginStatus501Ok

`func (o *DomainInspectorMeasurements) GetComputeOriginStatus501Ok() (*int64, bool)`

GetComputeOriginStatus501Ok returns a tuple with the ComputeOriginStatus501 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeOriginStatus501

`func (o *DomainInspectorMeasurements) SetComputeOriginStatus501(v int64)`

SetComputeOriginStatus501 sets ComputeOriginStatus501 field to given value.

### HasComputeOriginStatus501

`func (o *DomainInspectorMeasurements) HasComputeOriginStatus501() bool`

HasComputeOriginStatus501 returns a boolean if a field has been set.

### GetComputeOriginStatus502

`func (o *DomainInspectorMeasurements) GetComputeOriginStatus502() int64`

GetComputeOriginStatus502 returns the ComputeOriginStatus502 field if non-nil, zero value otherwise.

### GetComputeOriginStatus502Ok

`func (o *DomainInspectorMeasurements) GetComputeOriginStatus502Ok() (*int64, bool)`

GetComputeOriginStatus502Ok returns a tuple with the ComputeOriginStatus502 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeOriginStatus502

`func (o *DomainInspectorMeasurements) SetComputeOriginStatus502(v int64)`

SetComputeOriginStatus502 sets ComputeOriginStatus502 field to given value.

### HasComputeOriginStatus502

`func (o *DomainInspectorMeasurements) HasComputeOriginStatus502() bool`

HasComputeOriginStatus502 returns a boolean if a field has been set.

### GetComputeOriginStatus503

`func (o *DomainInspectorMeasurements) GetComputeOriginStatus503() int64`

GetComputeOriginStatus503 returns the ComputeOriginStatus503 field if non-nil, zero value otherwise.

### GetComputeOriginStatus503Ok

`func (o *DomainInspectorMeasurements) GetComputeOriginStatus503Ok() (*int64, bool)`

GetComputeOriginStatus503Ok returns a tuple with the ComputeOriginStatus503 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeOriginStatus503

`func (o *DomainInspectorMeasurements) SetComputeOriginStatus503(v int64)`

SetComputeOriginStatus503 sets ComputeOriginStatus503 field to given value.

### HasComputeOriginStatus503

`func (o *DomainInspectorMeasurements) HasComputeOriginStatus503() bool`

HasComputeOriginStatus503 returns a boolean if a field has been set.

### GetComputeOriginStatus504

`func (o *DomainInspectorMeasurements) GetComputeOriginStatus504() int64`

GetComputeOriginStatus504 returns the ComputeOriginStatus504 field if non-nil, zero value otherwise.

### GetComputeOriginStatus504Ok

`func (o *DomainInspectorMeasurements) GetComputeOriginStatus504Ok() (*int64, bool)`

GetComputeOriginStatus504Ok returns a tuple with the ComputeOriginStatus504 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeOriginStatus504

`func (o *DomainInspectorMeasurements) SetComputeOriginStatus504(v int64)`

SetComputeOriginStatus504 sets ComputeOriginStatus504 field to given value.

### HasComputeOriginStatus504

`func (o *DomainInspectorMeasurements) HasComputeOriginStatus504() bool`

HasComputeOriginStatus504 returns a boolean if a field has been set.

### GetComputeOriginStatus505

`func (o *DomainInspectorMeasurements) GetComputeOriginStatus505() int64`

GetComputeOriginStatus505 returns the ComputeOriginStatus505 field if non-nil, zero value otherwise.

### GetComputeOriginStatus505Ok

`func (o *DomainInspectorMeasurements) GetComputeOriginStatus505Ok() (*int64, bool)`

GetComputeOriginStatus505Ok returns a tuple with the ComputeOriginStatus505 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeOriginStatus505

`func (o *DomainInspectorMeasurements) SetComputeOriginStatus505(v int64)`

SetComputeOriginStatus505 sets ComputeOriginStatus505 field to given value.

### HasComputeOriginStatus505

`func (o *DomainInspectorMeasurements) HasComputeOriginStatus505() bool`

HasComputeOriginStatus505 returns a boolean if a field has been set.

### GetComputeOriginStatus530

`func (o *DomainInspectorMeasurements) GetComputeOriginStatus530() int64`

GetComputeOriginStatus530 returns the ComputeOriginStatus530 field if non-nil, zero value otherwise.

### GetComputeOriginStatus530Ok

`func (o *DomainInspectorMeasurements) GetComputeOriginStatus530Ok() (*int64, bool)`

GetComputeOriginStatus530Ok returns a tuple with the ComputeOriginStatus530 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeOriginStatus530

`func (o *DomainInspectorMeasurements) SetComputeOriginStatus530(v int64)`

SetComputeOriginStatus530 sets ComputeOriginStatus530 field to given value.

### HasComputeOriginStatus530

`func (o *DomainInspectorMeasurements) HasComputeOriginStatus530() bool`

HasComputeOriginStatus530 returns a boolean if a field has been set.

### GetComputeOriginStatus5xx

`func (o *DomainInspectorMeasurements) GetComputeOriginStatus5xx() int64`

GetComputeOriginStatus5xx returns the ComputeOriginStatus5xx field if non-nil, zero value otherwise.

### GetComputeOriginStatus5xxOk

`func (o *DomainInspectorMeasurements) GetComputeOriginStatus5xxOk() (*int64, bool)`

GetComputeOriginStatus5xxOk returns a tuple with the ComputeOriginStatus5xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeOriginStatus5xx

`func (o *DomainInspectorMeasurements) SetComputeOriginStatus5xx(v int64)`

SetComputeOriginStatus5xx sets ComputeOriginStatus5xx field to given value.

### HasComputeOriginStatus5xx

`func (o *DomainInspectorMeasurements) HasComputeOriginStatus5xx() bool`

HasComputeOriginStatus5xx returns a boolean if a field has been set.

### GetComputeReqBodyBytes

`func (o *DomainInspectorMeasurements) GetComputeReqBodyBytes() int64`

GetComputeReqBodyBytes returns the ComputeReqBodyBytes field if non-nil, zero value otherwise.

### GetComputeReqBodyBytesOk

`func (o *DomainInspectorMeasurements) GetComputeReqBodyBytesOk() (*int64, bool)`

GetComputeReqBodyBytesOk returns a tuple with the ComputeReqBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeReqBodyBytes

`func (o *DomainInspectorMeasurements) SetComputeReqBodyBytes(v int64)`

SetComputeReqBodyBytes sets ComputeReqBodyBytes field to given value.

### HasComputeReqBodyBytes

`func (o *DomainInspectorMeasurements) HasComputeReqBodyBytes() bool`

HasComputeReqBodyBytes returns a boolean if a field has been set.

### GetComputeReqHeaderBytes

`func (o *DomainInspectorMeasurements) GetComputeReqHeaderBytes() int64`

GetComputeReqHeaderBytes returns the ComputeReqHeaderBytes field if non-nil, zero value otherwise.

### GetComputeReqHeaderBytesOk

`func (o *DomainInspectorMeasurements) GetComputeReqHeaderBytesOk() (*int64, bool)`

GetComputeReqHeaderBytesOk returns a tuple with the ComputeReqHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeReqHeaderBytes

`func (o *DomainInspectorMeasurements) SetComputeReqHeaderBytes(v int64)`

SetComputeReqHeaderBytes sets ComputeReqHeaderBytes field to given value.

### HasComputeReqHeaderBytes

`func (o *DomainInspectorMeasurements) HasComputeReqHeaderBytes() bool`

HasComputeReqHeaderBytes returns a boolean if a field has been set.

### GetComputeRequestTimeBilledMs

`func (o *DomainInspectorMeasurements) GetComputeRequestTimeBilledMs() int64`

GetComputeRequestTimeBilledMs returns the ComputeRequestTimeBilledMs field if non-nil, zero value otherwise.

### GetComputeRequestTimeBilledMsOk

`func (o *DomainInspectorMeasurements) GetComputeRequestTimeBilledMsOk() (*int64, bool)`

GetComputeRequestTimeBilledMsOk returns a tuple with the ComputeRequestTimeBilledMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRequestTimeBilledMs

`func (o *DomainInspectorMeasurements) SetComputeRequestTimeBilledMs(v int64)`

SetComputeRequestTimeBilledMs sets ComputeRequestTimeBilledMs field to given value.

### HasComputeRequestTimeBilledMs

`func (o *DomainInspectorMeasurements) HasComputeRequestTimeBilledMs() bool`

HasComputeRequestTimeBilledMs returns a boolean if a field has been set.

### GetComputeRequestTimeMs

`func (o *DomainInspectorMeasurements) GetComputeRequestTimeMs() int64`

GetComputeRequestTimeMs returns the ComputeRequestTimeMs field if non-nil, zero value otherwise.

### GetComputeRequestTimeMsOk

`func (o *DomainInspectorMeasurements) GetComputeRequestTimeMsOk() (*int64, bool)`

GetComputeRequestTimeMsOk returns a tuple with the ComputeRequestTimeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRequestTimeMs

`func (o *DomainInspectorMeasurements) SetComputeRequestTimeMs(v int64)`

SetComputeRequestTimeMs sets ComputeRequestTimeMs field to given value.

### HasComputeRequestTimeMs

`func (o *DomainInspectorMeasurements) HasComputeRequestTimeMs() bool`

HasComputeRequestTimeMs returns a boolean if a field has been set.

### GetComputeRequest

`func (o *DomainInspectorMeasurements) GetComputeRequest() int64`

GetComputeRequest returns the ComputeRequest field if non-nil, zero value otherwise.

### GetComputeRequestOk

`func (o *DomainInspectorMeasurements) GetComputeRequestOk() (*int64, bool)`

GetComputeRequestOk returns a tuple with the ComputeRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRequest

`func (o *DomainInspectorMeasurements) SetComputeRequest(v int64)`

SetComputeRequest sets ComputeRequest field to given value.

### HasComputeRequest

`func (o *DomainInspectorMeasurements) HasComputeRequest() bool`

HasComputeRequest returns a boolean if a field has been set.

### GetComputeRespBodyBytes

`func (o *DomainInspectorMeasurements) GetComputeRespBodyBytes() int64`

GetComputeRespBodyBytes returns the ComputeRespBodyBytes field if non-nil, zero value otherwise.

### GetComputeRespBodyBytesOk

`func (o *DomainInspectorMeasurements) GetComputeRespBodyBytesOk() (*int64, bool)`

GetComputeRespBodyBytesOk returns a tuple with the ComputeRespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespBodyBytes

`func (o *DomainInspectorMeasurements) SetComputeRespBodyBytes(v int64)`

SetComputeRespBodyBytes sets ComputeRespBodyBytes field to given value.

### HasComputeRespBodyBytes

`func (o *DomainInspectorMeasurements) HasComputeRespBodyBytes() bool`

HasComputeRespBodyBytes returns a boolean if a field has been set.

### GetComputeRespHeaderBytes

`func (o *DomainInspectorMeasurements) GetComputeRespHeaderBytes() int64`

GetComputeRespHeaderBytes returns the ComputeRespHeaderBytes field if non-nil, zero value otherwise.

### GetComputeRespHeaderBytesOk

`func (o *DomainInspectorMeasurements) GetComputeRespHeaderBytesOk() (*int64, bool)`

GetComputeRespHeaderBytesOk returns a tuple with the ComputeRespHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespHeaderBytes

`func (o *DomainInspectorMeasurements) SetComputeRespHeaderBytes(v int64)`

SetComputeRespHeaderBytes sets ComputeRespHeaderBytes field to given value.

### HasComputeRespHeaderBytes

`func (o *DomainInspectorMeasurements) HasComputeRespHeaderBytes() bool`

HasComputeRespHeaderBytes returns a boolean if a field has been set.

### GetComputeRespStatus103

`func (o *DomainInspectorMeasurements) GetComputeRespStatus103() int64`

GetComputeRespStatus103 returns the ComputeRespStatus103 field if non-nil, zero value otherwise.

### GetComputeRespStatus103Ok

`func (o *DomainInspectorMeasurements) GetComputeRespStatus103Ok() (*int64, bool)`

GetComputeRespStatus103Ok returns a tuple with the ComputeRespStatus103 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespStatus103

`func (o *DomainInspectorMeasurements) SetComputeRespStatus103(v int64)`

SetComputeRespStatus103 sets ComputeRespStatus103 field to given value.

### HasComputeRespStatus103

`func (o *DomainInspectorMeasurements) HasComputeRespStatus103() bool`

HasComputeRespStatus103 returns a boolean if a field has been set.

### GetComputeRespStatus1xx

`func (o *DomainInspectorMeasurements) GetComputeRespStatus1xx() int64`

GetComputeRespStatus1xx returns the ComputeRespStatus1xx field if non-nil, zero value otherwise.

### GetComputeRespStatus1xxOk

`func (o *DomainInspectorMeasurements) GetComputeRespStatus1xxOk() (*int64, bool)`

GetComputeRespStatus1xxOk returns a tuple with the ComputeRespStatus1xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespStatus1xx

`func (o *DomainInspectorMeasurements) SetComputeRespStatus1xx(v int64)`

SetComputeRespStatus1xx sets ComputeRespStatus1xx field to given value.

### HasComputeRespStatus1xx

`func (o *DomainInspectorMeasurements) HasComputeRespStatus1xx() bool`

HasComputeRespStatus1xx returns a boolean if a field has been set.

### GetComputeRespStatus200

`func (o *DomainInspectorMeasurements) GetComputeRespStatus200() int64`

GetComputeRespStatus200 returns the ComputeRespStatus200 field if non-nil, zero value otherwise.

### GetComputeRespStatus200Ok

`func (o *DomainInspectorMeasurements) GetComputeRespStatus200Ok() (*int64, bool)`

GetComputeRespStatus200Ok returns a tuple with the ComputeRespStatus200 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespStatus200

`func (o *DomainInspectorMeasurements) SetComputeRespStatus200(v int64)`

SetComputeRespStatus200 sets ComputeRespStatus200 field to given value.

### HasComputeRespStatus200

`func (o *DomainInspectorMeasurements) HasComputeRespStatus200() bool`

HasComputeRespStatus200 returns a boolean if a field has been set.

### GetComputeRespStatus204

`func (o *DomainInspectorMeasurements) GetComputeRespStatus204() int64`

GetComputeRespStatus204 returns the ComputeRespStatus204 field if non-nil, zero value otherwise.

### GetComputeRespStatus204Ok

`func (o *DomainInspectorMeasurements) GetComputeRespStatus204Ok() (*int64, bool)`

GetComputeRespStatus204Ok returns a tuple with the ComputeRespStatus204 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespStatus204

`func (o *DomainInspectorMeasurements) SetComputeRespStatus204(v int64)`

SetComputeRespStatus204 sets ComputeRespStatus204 field to given value.

### HasComputeRespStatus204

`func (o *DomainInspectorMeasurements) HasComputeRespStatus204() bool`

HasComputeRespStatus204 returns a boolean if a field has been set.

### GetComputeRespStatus206

`func (o *DomainInspectorMeasurements) GetComputeRespStatus206() int64`

GetComputeRespStatus206 returns the ComputeRespStatus206 field if non-nil, zero value otherwise.

### GetComputeRespStatus206Ok

`func (o *DomainInspectorMeasurements) GetComputeRespStatus206Ok() (*int64, bool)`

GetComputeRespStatus206Ok returns a tuple with the ComputeRespStatus206 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespStatus206

`func (o *DomainInspectorMeasurements) SetComputeRespStatus206(v int64)`

SetComputeRespStatus206 sets ComputeRespStatus206 field to given value.

### HasComputeRespStatus206

`func (o *DomainInspectorMeasurements) HasComputeRespStatus206() bool`

HasComputeRespStatus206 returns a boolean if a field has been set.

### GetComputeRespStatus2xx

`func (o *DomainInspectorMeasurements) GetComputeRespStatus2xx() int64`

GetComputeRespStatus2xx returns the ComputeRespStatus2xx field if non-nil, zero value otherwise.

### GetComputeRespStatus2xxOk

`func (o *DomainInspectorMeasurements) GetComputeRespStatus2xxOk() (*int64, bool)`

GetComputeRespStatus2xxOk returns a tuple with the ComputeRespStatus2xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespStatus2xx

`func (o *DomainInspectorMeasurements) SetComputeRespStatus2xx(v int64)`

SetComputeRespStatus2xx sets ComputeRespStatus2xx field to given value.

### HasComputeRespStatus2xx

`func (o *DomainInspectorMeasurements) HasComputeRespStatus2xx() bool`

HasComputeRespStatus2xx returns a boolean if a field has been set.

### GetComputeRespStatus301

`func (o *DomainInspectorMeasurements) GetComputeRespStatus301() int64`

GetComputeRespStatus301 returns the ComputeRespStatus301 field if non-nil, zero value otherwise.

### GetComputeRespStatus301Ok

`func (o *DomainInspectorMeasurements) GetComputeRespStatus301Ok() (*int64, bool)`

GetComputeRespStatus301Ok returns a tuple with the ComputeRespStatus301 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespStatus301

`func (o *DomainInspectorMeasurements) SetComputeRespStatus301(v int64)`

SetComputeRespStatus301 sets ComputeRespStatus301 field to given value.

### HasComputeRespStatus301

`func (o *DomainInspectorMeasurements) HasComputeRespStatus301() bool`

HasComputeRespStatus301 returns a boolean if a field has been set.

### GetComputeRespStatus302

`func (o *DomainInspectorMeasurements) GetComputeRespStatus302() int64`

GetComputeRespStatus302 returns the ComputeRespStatus302 field if non-nil, zero value otherwise.

### GetComputeRespStatus302Ok

`func (o *DomainInspectorMeasurements) GetComputeRespStatus302Ok() (*int64, bool)`

GetComputeRespStatus302Ok returns a tuple with the ComputeRespStatus302 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespStatus302

`func (o *DomainInspectorMeasurements) SetComputeRespStatus302(v int64)`

SetComputeRespStatus302 sets ComputeRespStatus302 field to given value.

### HasComputeRespStatus302

`func (o *DomainInspectorMeasurements) HasComputeRespStatus302() bool`

HasComputeRespStatus302 returns a boolean if a field has been set.

### GetComputeRespStatus304

`func (o *DomainInspectorMeasurements) GetComputeRespStatus304() int64`

GetComputeRespStatus304 returns the ComputeRespStatus304 field if non-nil, zero value otherwise.

### GetComputeRespStatus304Ok

`func (o *DomainInspectorMeasurements) GetComputeRespStatus304Ok() (*int64, bool)`

GetComputeRespStatus304Ok returns a tuple with the ComputeRespStatus304 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespStatus304

`func (o *DomainInspectorMeasurements) SetComputeRespStatus304(v int64)`

SetComputeRespStatus304 sets ComputeRespStatus304 field to given value.

### HasComputeRespStatus304

`func (o *DomainInspectorMeasurements) HasComputeRespStatus304() bool`

HasComputeRespStatus304 returns a boolean if a field has been set.

### GetComputeRespStatus3xx

`func (o *DomainInspectorMeasurements) GetComputeRespStatus3xx() int64`

GetComputeRespStatus3xx returns the ComputeRespStatus3xx field if non-nil, zero value otherwise.

### GetComputeRespStatus3xxOk

`func (o *DomainInspectorMeasurements) GetComputeRespStatus3xxOk() (*int64, bool)`

GetComputeRespStatus3xxOk returns a tuple with the ComputeRespStatus3xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespStatus3xx

`func (o *DomainInspectorMeasurements) SetComputeRespStatus3xx(v int64)`

SetComputeRespStatus3xx sets ComputeRespStatus3xx field to given value.

### HasComputeRespStatus3xx

`func (o *DomainInspectorMeasurements) HasComputeRespStatus3xx() bool`

HasComputeRespStatus3xx returns a boolean if a field has been set.

### GetComputeRespStatus400

`func (o *DomainInspectorMeasurements) GetComputeRespStatus400() int64`

GetComputeRespStatus400 returns the ComputeRespStatus400 field if non-nil, zero value otherwise.

### GetComputeRespStatus400Ok

`func (o *DomainInspectorMeasurements) GetComputeRespStatus400Ok() (*int64, bool)`

GetComputeRespStatus400Ok returns a tuple with the ComputeRespStatus400 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespStatus400

`func (o *DomainInspectorMeasurements) SetComputeRespStatus400(v int64)`

SetComputeRespStatus400 sets ComputeRespStatus400 field to given value.

### HasComputeRespStatus400

`func (o *DomainInspectorMeasurements) HasComputeRespStatus400() bool`

HasComputeRespStatus400 returns a boolean if a field has been set.

### GetComputeRespStatus401

`func (o *DomainInspectorMeasurements) GetComputeRespStatus401() int64`

GetComputeRespStatus401 returns the ComputeRespStatus401 field if non-nil, zero value otherwise.

### GetComputeRespStatus401Ok

`func (o *DomainInspectorMeasurements) GetComputeRespStatus401Ok() (*int64, bool)`

GetComputeRespStatus401Ok returns a tuple with the ComputeRespStatus401 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespStatus401

`func (o *DomainInspectorMeasurements) SetComputeRespStatus401(v int64)`

SetComputeRespStatus401 sets ComputeRespStatus401 field to given value.

### HasComputeRespStatus401

`func (o *DomainInspectorMeasurements) HasComputeRespStatus401() bool`

HasComputeRespStatus401 returns a boolean if a field has been set.

### GetComputeRespStatus403

`func (o *DomainInspectorMeasurements) GetComputeRespStatus403() int64`

GetComputeRespStatus403 returns the ComputeRespStatus403 field if non-nil, zero value otherwise.

### GetComputeRespStatus403Ok

`func (o *DomainInspectorMeasurements) GetComputeRespStatus403Ok() (*int64, bool)`

GetComputeRespStatus403Ok returns a tuple with the ComputeRespStatus403 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespStatus403

`func (o *DomainInspectorMeasurements) SetComputeRespStatus403(v int64)`

SetComputeRespStatus403 sets ComputeRespStatus403 field to given value.

### HasComputeRespStatus403

`func (o *DomainInspectorMeasurements) HasComputeRespStatus403() bool`

HasComputeRespStatus403 returns a boolean if a field has been set.

### GetComputeRespStatus404

`func (o *DomainInspectorMeasurements) GetComputeRespStatus404() int64`

GetComputeRespStatus404 returns the ComputeRespStatus404 field if non-nil, zero value otherwise.

### GetComputeRespStatus404Ok

`func (o *DomainInspectorMeasurements) GetComputeRespStatus404Ok() (*int64, bool)`

GetComputeRespStatus404Ok returns a tuple with the ComputeRespStatus404 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespStatus404

`func (o *DomainInspectorMeasurements) SetComputeRespStatus404(v int64)`

SetComputeRespStatus404 sets ComputeRespStatus404 field to given value.

### HasComputeRespStatus404

`func (o *DomainInspectorMeasurements) HasComputeRespStatus404() bool`

HasComputeRespStatus404 returns a boolean if a field has been set.

### GetComputeRespStatus416

`func (o *DomainInspectorMeasurements) GetComputeRespStatus416() int64`

GetComputeRespStatus416 returns the ComputeRespStatus416 field if non-nil, zero value otherwise.

### GetComputeRespStatus416Ok

`func (o *DomainInspectorMeasurements) GetComputeRespStatus416Ok() (*int64, bool)`

GetComputeRespStatus416Ok returns a tuple with the ComputeRespStatus416 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespStatus416

`func (o *DomainInspectorMeasurements) SetComputeRespStatus416(v int64)`

SetComputeRespStatus416 sets ComputeRespStatus416 field to given value.

### HasComputeRespStatus416

`func (o *DomainInspectorMeasurements) HasComputeRespStatus416() bool`

HasComputeRespStatus416 returns a boolean if a field has been set.

### GetComputeRespStatus429

`func (o *DomainInspectorMeasurements) GetComputeRespStatus429() int64`

GetComputeRespStatus429 returns the ComputeRespStatus429 field if non-nil, zero value otherwise.

### GetComputeRespStatus429Ok

`func (o *DomainInspectorMeasurements) GetComputeRespStatus429Ok() (*int64, bool)`

GetComputeRespStatus429Ok returns a tuple with the ComputeRespStatus429 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespStatus429

`func (o *DomainInspectorMeasurements) SetComputeRespStatus429(v int64)`

SetComputeRespStatus429 sets ComputeRespStatus429 field to given value.

### HasComputeRespStatus429

`func (o *DomainInspectorMeasurements) HasComputeRespStatus429() bool`

HasComputeRespStatus429 returns a boolean if a field has been set.

### GetComputeRespStatus4xx

`func (o *DomainInspectorMeasurements) GetComputeRespStatus4xx() int64`

GetComputeRespStatus4xx returns the ComputeRespStatus4xx field if non-nil, zero value otherwise.

### GetComputeRespStatus4xxOk

`func (o *DomainInspectorMeasurements) GetComputeRespStatus4xxOk() (*int64, bool)`

GetComputeRespStatus4xxOk returns a tuple with the ComputeRespStatus4xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespStatus4xx

`func (o *DomainInspectorMeasurements) SetComputeRespStatus4xx(v int64)`

SetComputeRespStatus4xx sets ComputeRespStatus4xx field to given value.

### HasComputeRespStatus4xx

`func (o *DomainInspectorMeasurements) HasComputeRespStatus4xx() bool`

HasComputeRespStatus4xx returns a boolean if a field has been set.

### GetComputeRespStatus500

`func (o *DomainInspectorMeasurements) GetComputeRespStatus500() int64`

GetComputeRespStatus500 returns the ComputeRespStatus500 field if non-nil, zero value otherwise.

### GetComputeRespStatus500Ok

`func (o *DomainInspectorMeasurements) GetComputeRespStatus500Ok() (*int64, bool)`

GetComputeRespStatus500Ok returns a tuple with the ComputeRespStatus500 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespStatus500

`func (o *DomainInspectorMeasurements) SetComputeRespStatus500(v int64)`

SetComputeRespStatus500 sets ComputeRespStatus500 field to given value.

### HasComputeRespStatus500

`func (o *DomainInspectorMeasurements) HasComputeRespStatus500() bool`

HasComputeRespStatus500 returns a boolean if a field has been set.

### GetComputeRespStatus501

`func (o *DomainInspectorMeasurements) GetComputeRespStatus501() int64`

GetComputeRespStatus501 returns the ComputeRespStatus501 field if non-nil, zero value otherwise.

### GetComputeRespStatus501Ok

`func (o *DomainInspectorMeasurements) GetComputeRespStatus501Ok() (*int64, bool)`

GetComputeRespStatus501Ok returns a tuple with the ComputeRespStatus501 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespStatus501

`func (o *DomainInspectorMeasurements) SetComputeRespStatus501(v int64)`

SetComputeRespStatus501 sets ComputeRespStatus501 field to given value.

### HasComputeRespStatus501

`func (o *DomainInspectorMeasurements) HasComputeRespStatus501() bool`

HasComputeRespStatus501 returns a boolean if a field has been set.

### GetComputeRespStatus502

`func (o *DomainInspectorMeasurements) GetComputeRespStatus502() int64`

GetComputeRespStatus502 returns the ComputeRespStatus502 field if non-nil, zero value otherwise.

### GetComputeRespStatus502Ok

`func (o *DomainInspectorMeasurements) GetComputeRespStatus502Ok() (*int64, bool)`

GetComputeRespStatus502Ok returns a tuple with the ComputeRespStatus502 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespStatus502

`func (o *DomainInspectorMeasurements) SetComputeRespStatus502(v int64)`

SetComputeRespStatus502 sets ComputeRespStatus502 field to given value.

### HasComputeRespStatus502

`func (o *DomainInspectorMeasurements) HasComputeRespStatus502() bool`

HasComputeRespStatus502 returns a boolean if a field has been set.

### GetComputeRespStatus503

`func (o *DomainInspectorMeasurements) GetComputeRespStatus503() int64`

GetComputeRespStatus503 returns the ComputeRespStatus503 field if non-nil, zero value otherwise.

### GetComputeRespStatus503Ok

`func (o *DomainInspectorMeasurements) GetComputeRespStatus503Ok() (*int64, bool)`

GetComputeRespStatus503Ok returns a tuple with the ComputeRespStatus503 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespStatus503

`func (o *DomainInspectorMeasurements) SetComputeRespStatus503(v int64)`

SetComputeRespStatus503 sets ComputeRespStatus503 field to given value.

### HasComputeRespStatus503

`func (o *DomainInspectorMeasurements) HasComputeRespStatus503() bool`

HasComputeRespStatus503 returns a boolean if a field has been set.

### GetComputeRespStatus504

`func (o *DomainInspectorMeasurements) GetComputeRespStatus504() int64`

GetComputeRespStatus504 returns the ComputeRespStatus504 field if non-nil, zero value otherwise.

### GetComputeRespStatus504Ok

`func (o *DomainInspectorMeasurements) GetComputeRespStatus504Ok() (*int64, bool)`

GetComputeRespStatus504Ok returns a tuple with the ComputeRespStatus504 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespStatus504

`func (o *DomainInspectorMeasurements) SetComputeRespStatus504(v int64)`

SetComputeRespStatus504 sets ComputeRespStatus504 field to given value.

### HasComputeRespStatus504

`func (o *DomainInspectorMeasurements) HasComputeRespStatus504() bool`

HasComputeRespStatus504 returns a boolean if a field has been set.

### GetComputeRespStatus505

`func (o *DomainInspectorMeasurements) GetComputeRespStatus505() int64`

GetComputeRespStatus505 returns the ComputeRespStatus505 field if non-nil, zero value otherwise.

### GetComputeRespStatus505Ok

`func (o *DomainInspectorMeasurements) GetComputeRespStatus505Ok() (*int64, bool)`

GetComputeRespStatus505Ok returns a tuple with the ComputeRespStatus505 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespStatus505

`func (o *DomainInspectorMeasurements) SetComputeRespStatus505(v int64)`

SetComputeRespStatus505 sets ComputeRespStatus505 field to given value.

### HasComputeRespStatus505

`func (o *DomainInspectorMeasurements) HasComputeRespStatus505() bool`

HasComputeRespStatus505 returns a boolean if a field has been set.

### GetComputeRespStatus530

`func (o *DomainInspectorMeasurements) GetComputeRespStatus530() int64`

GetComputeRespStatus530 returns the ComputeRespStatus530 field if non-nil, zero value otherwise.

### GetComputeRespStatus530Ok

`func (o *DomainInspectorMeasurements) GetComputeRespStatus530Ok() (*int64, bool)`

GetComputeRespStatus530Ok returns a tuple with the ComputeRespStatus530 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespStatus530

`func (o *DomainInspectorMeasurements) SetComputeRespStatus530(v int64)`

SetComputeRespStatus530 sets ComputeRespStatus530 field to given value.

### HasComputeRespStatus530

`func (o *DomainInspectorMeasurements) HasComputeRespStatus530() bool`

HasComputeRespStatus530 returns a boolean if a field has been set.

### GetComputeRespStatus5xx

`func (o *DomainInspectorMeasurements) GetComputeRespStatus5xx() int64`

GetComputeRespStatus5xx returns the ComputeRespStatus5xx field if non-nil, zero value otherwise.

### GetComputeRespStatus5xxOk

`func (o *DomainInspectorMeasurements) GetComputeRespStatus5xxOk() (*int64, bool)`

GetComputeRespStatus5xxOk returns a tuple with the ComputeRespStatus5xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespStatus5xx

`func (o *DomainInspectorMeasurements) SetComputeRespStatus5xx(v int64)`

SetComputeRespStatus5xx sets ComputeRespStatus5xx field to given value.

### HasComputeRespStatus5xx

`func (o *DomainInspectorMeasurements) HasComputeRespStatus5xx() bool`

HasComputeRespStatus5xx returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


