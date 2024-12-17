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
**OriginStatus1xx** | Pointer to **int64** | Number of \&quot;Informational\&quot; category status codes received from origin. | [optional] 
**OriginStatus2xx** | Pointer to **int64** | Number of \&quot;Success\&quot; status codes received from origin. | [optional] 
**OriginStatus3xx** | Pointer to **int64** | Number of \&quot;Redirection\&quot; codes received from origin. | [optional] 
**OriginStatus4xx** | Pointer to **int64** | Number of \&quot;Client Error\&quot; codes received from origin. | [optional] 
**OriginStatus5xx** | Pointer to **int64** | Number of \&quot;Server Error\&quot; codes received from origin. | [optional] 

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


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
