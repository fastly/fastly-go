// Package fastly is an API client library for interacting with most facets of the Fastly API.
package fastly

/*
Fastly API

Via the Fastly API you can perform any of the operations that are possible within the management console,  including creating services, domains, and backends, configuring rules or uploading your own application code, as well as account operations such as user administration and billing reports. The API is organized into collections of endpoints that allow manipulation of objects related to Fastly services and accounts. For the most accurate and up-to-date API reference content, visit our [Developer Hub](https://developer.fastly.com/reference/api/) 

API version: 1.0.0
Contact: oss@fastly.com
*/

// This code is auto-generated; DO NOT EDIT.


import (
	"encoding/json"
)

// OriginInspectorMeasurements A measurements object provides a count of the total number of `responses` received by Fastly from your origin servers in the reported time period, for the relevant POP and backend name as specified in the [entry](#entry-data-model). It also includes the number of responses for specific HTTP response status codes and for status code ranges. This dataset is sparse: only the keys with non-zero values will be included in the record. Where a specific status code does not have a field in this model (e.g., `429 Too Many Requests`), any responses with that code will be counted as part of the range count (`4xx` in this case) but will not be separately identified in the data. 
type OriginInspectorMeasurements struct {
	// Number of responses from origin.
	Responses *int32 `json:"responses,omitempty"`
	// Number of header bytes from origin.
	RespHeaderBytes *int32 `json:"resp_header_bytes,omitempty"`
	// Number of body bytes from origin.
	RespBodyBytes *int32 `json:"resp_body_bytes,omitempty"`
	// Number of 1xx \"Informational\" status codes delivered from origin.
	Status1xx *int32 `json:"status_1xx,omitempty"`
	// Number of 2xx \"Success\" status codes delivered from origin.
	Status2xx *int32 `json:"status_2xx,omitempty"`
	// Number of 3xx \"Redirection\" codes delivered from origin.
	Status3xx *int32 `json:"status_3xx,omitempty"`
	// Number of 4xx \"Client Error\" codes delivered from origin.
	Status4xx *int32 `json:"status_4xx,omitempty"`
	// Number of 5xx \"Server Error\" codes delivered from origin.
	Status5xx *int32 `json:"status_5xx,omitempty"`
	// Number of responses received with status code 200 (Success) from origin.
	Status200 *int32 `json:"status_200,omitempty"`
	// Number of responses received with status code 204 (No Content) from origin.
	Status204 *int32 `json:"status_204,omitempty"`
	// Number of responses received with status code 206 (Partial Content) from origin.
	Status206 *int32 `json:"status_206,omitempty"`
	// Number of responses received with status code 301 (Moved Permanently) from origin.
	Status301 *int32 `json:"status_301,omitempty"`
	// Number of responses received with status code 302 (Found) from origin.
	Status302 *int32 `json:"status_302,omitempty"`
	// Number of responses received with status code 304 (Not Modified) from origin.
	Status304 *int32 `json:"status_304,omitempty"`
	// Number of responses received with status code 400 (Bad Request) from origin.
	Status400 *int32 `json:"status_400,omitempty"`
	// Number of responses received with status code 401 (Unauthorized) from origin.
	Status401 *int32 `json:"status_401,omitempty"`
	// Number of responses received with status code 403 (Forbidden) from origin.
	Status403 *int32 `json:"status_403,omitempty"`
	// Number of responses received with status code 404 (Not Found) from origin.
	Status404 *int32 `json:"status_404,omitempty"`
	// Number of responses received with status code 416 (Range Not Satisfiable) from origin.
	Status416 *int32 `json:"status_416,omitempty"`
	// Number of responses received with status code 429 (Too Many Requests) from origin.
	Status429 *int32 `json:"status_429,omitempty"`
	// Number of responses received with status code 500 (Internal Server Error) from origin.
	Status500 *int32 `json:"status_500,omitempty"`
	// Number of responses received with status code 501 (Not Implemented) from origin.
	Status501 *int32 `json:"status_501,omitempty"`
	// Number of responses received with status code 502 (Bad Gateway) from origin.
	Status502 *int32 `json:"status_502,omitempty"`
	// Number of responses received with status code 503 (Service Unavailable) from origin.
	Status503 *int32 `json:"status_503,omitempty"`
	// Number of responses received with status code 504 (Gateway Timeout) from origin.
	Status504 *int32 `json:"status_504,omitempty"`
	// Number of responses received with status code 505 (HTTP Version Not Supported) from origin.
	Status505 *int32 `json:"status_505,omitempty"`
	// Number of responses from origin with latency between 0 and 1 millisecond.
	Latency0To1ms *int32 `json:"latency_0_to_1ms,omitempty"`
	// Number of responses from origin with latency between 1 and 5 milliseconds.
	Latency1To5ms *int32 `json:"latency_1_to_5ms,omitempty"`
	// Number of responses from origin with latency between 5 and 10 milliseconds.
	Latency5To10ms *int32 `json:"latency_5_to_10ms,omitempty"`
	// Number of responses from origin with latency between 10 and 50 milliseconds.
	Latency10To50ms *int32 `json:"latency_10_to_50ms,omitempty"`
	// Number of responses from origin with latency between 50 and 100 milliseconds.
	Latency50To100ms *int32 `json:"latency_50_to_100ms,omitempty"`
	// Number of responses from origin with latency between 100 and 250 milliseconds.
	Latency100To250ms *int32 `json:"latency_100_to_250ms,omitempty"`
	// Number of responses from origin with latency between 250 and 500 milliseconds.
	Latency250To500ms *int32 `json:"latency_250_to_500ms,omitempty"`
	// Number of responses from origin with latency between 500 and 1,000 milliseconds.
	Latency500To1000ms *int32 `json:"latency_500_to_1000ms,omitempty"`
	// Number of responses from origin with latency between 1,000 and 5,000 milliseconds.
	Latency1000To5000ms *int32 `json:"latency_1000_to_5000ms,omitempty"`
	// Number of responses from origin with latency between 5,000 and 10,000 milliseconds.
	Latency5000To10000ms *int32 `json:"latency_5000_to_10000ms,omitempty"`
	// Number of responses from origin with latency between 10,000 and 60,000 milliseconds.
	Latency10000To60000ms *int32 `json:"latency_10000_to_60000ms,omitempty"`
	// Number of responses from origin with latency of 60,000 milliseconds and above.
	Latency60000ms *int32 `json:"latency_60000ms,omitempty"`
	// Number of responses received for origin requests made by the Fastly WAF.
	WafResponses *int32 `json:"waf_responses,omitempty"`
	// Number of header bytes received for origin requests made by the Fastly WAF.
	WafRespHeaderBytes *int32 `json:"waf_resp_header_bytes,omitempty"`
	// Number of body bytes received for origin requests made by the Fastly WAF.
	WafRespBodyBytes *int32 `json:"waf_resp_body_bytes,omitempty"`
	// Number of 1xx \"Informational\" status codes received for origin requests made by the Fastly WAF.
	WafStatus1xx *int32 `json:"waf_status_1xx,omitempty"`
	// Number of 2xx \"Success\" status codes received for origin requests made by the Fastly WAF.
	WafStatus2xx *int32 `json:"waf_status_2xx,omitempty"`
	// Number of 3xx \"Redirection\" codes received for origin requests made by the Fastly WAF.
	WafStatus3xx *int32 `json:"waf_status_3xx,omitempty"`
	// Number of 4xx \"Client Error\" codes received for origin requests made by the Fastly WAF.
	WafStatus4xx *int32 `json:"waf_status_4xx,omitempty"`
	// Number of 5xx \"Server Error\" codes received for origin requests made by the Fastly WAF.
	WafStatus5xx *int32 `json:"waf_status_5xx,omitempty"`
	// Number of responses received with status code 200 (Success) received for origin requests made by the Fastly WAF.
	WafStatus200 *int32 `json:"waf_status_200,omitempty"`
	// Number of responses received with status code 204 (No Content) received for origin requests made by the Fastly WAF.
	WafStatus204 *int32 `json:"waf_status_204,omitempty"`
	// Number of responses received with status code 206 (Partial Content) received for origin requests made by the Fastly WAF.
	WafStatus206 *int32 `json:"waf_status_206,omitempty"`
	// Number of responses received with status code 301 (Moved Permanently) received for origin requests made by the Fastly WAF.
	WafStatus301 *int32 `json:"waf_status_301,omitempty"`
	// Number of responses received with status code 302 (Found) received for origin requests made by the Fastly WAF.
	WafStatus302 *int32 `json:"waf_status_302,omitempty"`
	// Number of responses received with status code 304 (Not Modified) received for origin requests made by the Fastly WAF.
	WafStatus304 *int32 `json:"waf_status_304,omitempty"`
	// Number of responses received with status code 400 (Bad Request) received for origin requests made by the Fastly WAF.
	WafStatus400 *int32 `json:"waf_status_400,omitempty"`
	// Number of responses received with status code 401 (Unauthorized) received for origin requests made by the Fastly WAF.
	WafStatus401 *int32 `json:"waf_status_401,omitempty"`
	// Number of responses received with status code 403 (Forbidden) received for origin requests made by the Fastly WAF.
	WafStatus403 *int32 `json:"waf_status_403,omitempty"`
	// Number of responses received with status code 404 (Not Found) received for origin requests made by the Fastly WAF.
	WafStatus404 *int32 `json:"waf_status_404,omitempty"`
	// Number of responses received with status code 416 (Range Not Satisfiable) received for origin requests made by the Fastly WAF.
	WafStatus416 *int32 `json:"waf_status_416,omitempty"`
	// Number of responses received with status code 429 (Too Many Requests) received for origin requests made by the Fastly WAF.
	WafStatus429 *int32 `json:"waf_status_429,omitempty"`
	// Number of responses received with status code 500 (Internal Server Error) received for origin requests made by the Fastly WAF.
	WafStatus500 *int32 `json:"waf_status_500,omitempty"`
	// Number of responses received with status code 501 (Not Implemented) received for origin requests made by the Fastly WAF.
	WafStatus501 *int32 `json:"waf_status_501,omitempty"`
	// Number of responses received with status code 502 (Bad Gateway) received for origin requests made by the Fastly WAF.
	WafStatus502 *int32 `json:"waf_status_502,omitempty"`
	// Number of responses received with status code 503 (Service Unavailable) received for origin requests made by the Fastly WAF.
	WafStatus503 *int32 `json:"waf_status_503,omitempty"`
	// Number of responses received with status code 504 (Gateway Timeout) received for origin requests made by the Fastly WAF.
	WafStatus504 *int32 `json:"waf_status_504,omitempty"`
	// Number of responses received with status code 505 (HTTP Version Not Supported) received for origin requests made by the Fastly WAF.
	WafStatus505 *int32 `json:"waf_status_505,omitempty"`
	// Number of responses with latency between 0 and 1 millisecond received for origin requests made by the Fastly WAF.
	WafLatency0To1ms *int32 `json:"waf_latency_0_to_1ms,omitempty"`
	// Number of responses with latency between 1 and 5 milliseconds received for origin requests made by the Fastly WAF.
	WafLatency1To5ms *int32 `json:"waf_latency_1_to_5ms,omitempty"`
	// Number of responses with latency between 5 and 10 milliseconds received for origin requests made by the Fastly WAF.
	WafLatency5To10ms *int32 `json:"waf_latency_5_to_10ms,omitempty"`
	// Number of responses with latency between 10 and 50 milliseconds received for origin requests made by the Fastly WAF.
	WafLatency10To50ms *int32 `json:"waf_latency_10_to_50ms,omitempty"`
	// Number of responses with latency between 50 and 100 milliseconds received for origin requests made by the Fastly WAF.
	WafLatency50To100ms *int32 `json:"waf_latency_50_to_100ms,omitempty"`
	// Number of responses with latency between 100 and 250 milliseconds received for origin requests made by the Fastly WAF.
	WafLatency100To250ms *int32 `json:"waf_latency_100_to_250ms,omitempty"`
	// Number of responses with latency between 250 and 500 milliseconds received for origin requests made by the Fastly WAF.
	WafLatency250To500ms *int32 `json:"waf_latency_250_to_500ms,omitempty"`
	// Number of responses with latency between 500 and 1,000 milliseconds received for origin requests made by the Fastly WAF.
	WafLatency500To1000ms *int32 `json:"waf_latency_500_to_1000ms,omitempty"`
	// Number of responses with latency between 1,000 and 5,000 milliseconds received for origin requests made by the Fastly WAF.
	WafLatency1000To5000ms *int32 `json:"waf_latency_1000_to_5000ms,omitempty"`
	// Number of responses with latency between 5,000 and 10,000 milliseconds received for origin requests made by the Fastly WAF.
	WafLatency5000To10000ms *int32 `json:"waf_latency_5000_to_10000ms,omitempty"`
	// Number of responses with latency between 10,000 and 60,000 milliseconds received for origin requests made by the Fastly WAF.
	WafLatency10000To60000ms *int32 `json:"waf_latency_10000_to_60000ms,omitempty"`
	// Number of responses with latency of 60,000 milliseconds and above received for origin requests made by the Fastly WAF.
	WafLatency60000ms *int32 `json:"waf_latency_60000ms,omitempty"`
	// Number of responses for origin received by Compute@Edge.
	ComputeResponses *int32 `json:"compute_responses,omitempty"`
	// Number of header bytes for origin received by Compute@Edge.
	ComputeRespHeaderBytes *int32 `json:"compute_resp_header_bytes,omitempty"`
	// Number of body bytes for origin received by Compute@Edge.
	ComputeRespBodyBytes *int32 `json:"compute_resp_body_bytes,omitempty"`
	// Number of 1xx \"Informational\" status codes for origin received by Compute@Edge.
	ComputeStatus1xx *int32 `json:"compute_status_1xx,omitempty"`
	// Number of 2xx \"Success\" status codes for origin received by Compute@Edge.
	ComputeStatus2xx *int32 `json:"compute_status_2xx,omitempty"`
	// Number of 3xx \"Redirection\" codes for origin received by Compute@Edge.
	ComputeStatus3xx *int32 `json:"compute_status_3xx,omitempty"`
	// Number of 4xx \"Client Error\" codes for origin received by Compute@Edge.
	ComputeStatus4xx *int32 `json:"compute_status_4xx,omitempty"`
	// Number of 5xx \"Server Error\" codes for origin received by Compute@Edge.
	ComputeStatus5xx *int32 `json:"compute_status_5xx,omitempty"`
	// Number of responses received with status code 200 (Success) for origin received by Compute@Edge.
	ComputeStatus200 *int32 `json:"compute_status_200,omitempty"`
	// Number of responses received with status code 204 (No Content) for origin received by Compute@Edge.
	ComputeStatus204 *int32 `json:"compute_status_204,omitempty"`
	// Number of responses received with status code 206 (Partial Content) for origin received by Compute@Edge.
	ComputeStatus206 *int32 `json:"compute_status_206,omitempty"`
	// Number of responses received with status code 301 (Moved Permanently) for origin received by Compute@Edge.
	ComputeStatus301 *int32 `json:"compute_status_301,omitempty"`
	// Number of responses received with status code 302 (Found) for origin received by Compute@Edge.
	ComputeStatus302 *int32 `json:"compute_status_302,omitempty"`
	// Number of responses received with status code 304 (Not Modified) for origin received by Compute@Edge.
	ComputeStatus304 *int32 `json:"compute_status_304,omitempty"`
	// Number of responses received with status code 400 (Bad Request) for origin received by Compute@Edge.
	ComputeStatus400 *int32 `json:"compute_status_400,omitempty"`
	// Number of responses received with status code 401 (Unauthorized) for origin received by Compute@Edge.
	ComputeStatus401 *int32 `json:"compute_status_401,omitempty"`
	// Number of responses received with status code 403 (Forbidden) for origin received by Compute@Edge.
	ComputeStatus403 *int32 `json:"compute_status_403,omitempty"`
	// Number of responses received with status code 404 (Not Found) for origin received by Compute@Edge.
	ComputeStatus404 *int32 `json:"compute_status_404,omitempty"`
	// Number of responses received with status code 416 (Range Not Satisfiable) for origin received by Compute@Edge.
	ComputeStatus416 *int32 `json:"compute_status_416,omitempty"`
	// Number of responses received with status code 429 (Too Many Requests) for origin received by Compute@Edge.
	ComputeStatus429 *int32 `json:"compute_status_429,omitempty"`
	// Number of responses received with status code 500 (Internal Server Error) for origin received by Compute@Edge.
	ComputeStatus500 *int32 `json:"compute_status_500,omitempty"`
	// Number of responses received with status code 501 (Not Implemented) for origin received by Compute@Edge.
	ComputeStatus501 *int32 `json:"compute_status_501,omitempty"`
	// Number of responses received with status code 502 (Bad Gateway) for origin received by Compute@Edge.
	ComputeStatus502 *int32 `json:"compute_status_502,omitempty"`
	// Number of responses received with status code 503 (Service Unavailable) for origin received by Compute@Edge.
	ComputeStatus503 *int32 `json:"compute_status_503,omitempty"`
	// Number of responses received with status code 504 (Gateway Timeout) for origin received by Compute@Edge.
	ComputeStatus504 *int32 `json:"compute_status_504,omitempty"`
	// Number of responses received with status code 505 (HTTP Version Not Supported) for origin received by Compute@Edge.
	ComputeStatus505 *int32 `json:"compute_status_505,omitempty"`
	// Number of responses with latency between 0 and 1 millisecond for origin received by Compute@Edge.
	ComputeLatency0To1ms *int32 `json:"compute_latency_0_to_1ms,omitempty"`
	// Number of responses with latency between 1 and 5 milliseconds for origin received by Compute@Edge.
	ComputeLatency1To5ms *int32 `json:"compute_latency_1_to_5ms,omitempty"`
	// Number of responses with latency between 5 and 10 milliseconds for origin received by Compute@Edge.
	ComputeLatency5To10ms *int32 `json:"compute_latency_5_to_10ms,omitempty"`
	// Number of responses with latency between 10 and 50 milliseconds for origin received by Compute@Edge.
	ComputeLatency10To50ms *int32 `json:"compute_latency_10_to_50ms,omitempty"`
	// Number of responses with latency between 50 and 100 milliseconds for origin received by Compute@Edge.
	ComputeLatency50To100ms *int32 `json:"compute_latency_50_to_100ms,omitempty"`
	// Number of responses with latency between 100 and 250 milliseconds for origin received by Compute@Edge.
	ComputeLatency100To250ms *int32 `json:"compute_latency_100_to_250ms,omitempty"`
	// Number of responses with latency between 250 and 500 milliseconds for origin received by Compute@Edge.
	ComputeLatency250To500ms *int32 `json:"compute_latency_250_to_500ms,omitempty"`
	// Number of responses with latency between 500 and 1,000 milliseconds for origin received by Compute@Edge.
	ComputeLatency500To1000ms *int32 `json:"compute_latency_500_to_1000ms,omitempty"`
	// Number of responses with latency between 1,000 and 5,000 milliseconds for origin received by Compute@Edge.
	ComputeLatency1000To5000ms *int32 `json:"compute_latency_1000_to_5000ms,omitempty"`
	// Number of responses with latency between 5,000 and 10,000 milliseconds for origin received by Compute@Edge.
	ComputeLatency5000To10000ms *int32 `json:"compute_latency_5000_to_10000ms,omitempty"`
	// Number of responses with latency between 10,000 and 60,000 milliseconds for origin received by Compute@Edge.
	ComputeLatency10000To60000ms *int32 `json:"compute_latency_10000_to_60000ms,omitempty"`
	// Number of responses with latency of 60,000 milliseconds and above for origin received by Compute@Edge.
	ComputeLatency60000ms *int32 `json:"compute_latency_60000ms,omitempty"`
	// Number of responses received for origin requests made by all sources.
	AllResponses *int32 `json:"all_responses,omitempty"`
	// Number of header bytes received for origin requests made by all sources.
	AllRespHeaderBytes *int32 `json:"all_resp_header_bytes,omitempty"`
	// Number of body bytes received for origin requests made by all sources.
	AllRespBodyBytes *int32 `json:"all_resp_body_bytes,omitempty"`
	// Number of 1xx \"Informational\" category status codes delivered received for origin requests made by all sources.
	AllStatus1xx *int32 `json:"all_status_1xx,omitempty"`
	// Number of 2xx \"Success\" status codes received for origin requests made by all sources.
	AllStatus2xx *int32 `json:"all_status_2xx,omitempty"`
	// Number of 3xx \"Redirection\" codes received for origin requests made by all sources.
	AllStatus3xx *int32 `json:"all_status_3xx,omitempty"`
	// Number of 4xx \"Client Error\" codes received for origin requests made by all sources.
	AllStatus4xx *int32 `json:"all_status_4xx,omitempty"`
	// Number of 5xx \"Server Error\" codes received for origin requests made by all sources.
	AllStatus5xx *int32 `json:"all_status_5xx,omitempty"`
	// Number of responses received with status code 200 (Success) received for origin requests made by all sources.
	AllStatus200 *int32 `json:"all_status_200,omitempty"`
	// Number of responses received with status code 204 (No Content) received for origin requests made by all sources.
	AllStatus204 *int32 `json:"all_status_204,omitempty"`
	// Number of responses received with status code 206 (Partial Content) received for origin requests made by all sources.
	AllStatus206 *int32 `json:"all_status_206,omitempty"`
	// Number of responses received with status code 301 (Moved Permanently) received for origin requests made by all sources.
	AllStatus301 *int32 `json:"all_status_301,omitempty"`
	// Number of responses received with status code 302 (Found) received for origin requests made by all sources.
	AllStatus302 *int32 `json:"all_status_302,omitempty"`
	// Number of responses received with status code 304 (Not Modified) received for origin requests made by all sources.
	AllStatus304 *int32 `json:"all_status_304,omitempty"`
	// Number of responses received with status code 400 (Bad Request) received for origin requests made by all sources.
	AllStatus400 *int32 `json:"all_status_400,omitempty"`
	// Number of responses received with status code 401 (Unauthorized) received for origin requests made by all sources.
	AllStatus401 *int32 `json:"all_status_401,omitempty"`
	// Number of responses received with status code 403 (Forbidden) received for origin requests made by all sources.
	AllStatus403 *int32 `json:"all_status_403,omitempty"`
	// Number of responses received with status code 404 (Not Found) received for origin requests made by all sources.
	AllStatus404 *int32 `json:"all_status_404,omitempty"`
	// Number of responses received with status code 416 (Range Not Satisfiable) received for origin requests made by all sources.
	AllStatus416 *int32 `json:"all_status_416,omitempty"`
	// Number of responses received with status code 429 (Too Many Requests) received for origin requests made by all sources.
	AllStatus429 *int32 `json:"all_status_429,omitempty"`
	// Number of responses received with status code 500 (Internal Server Error) received for origin requests made by all sources.
	AllStatus500 *int32 `json:"all_status_500,omitempty"`
	// Number of responses received with status code 501 (Not Implemented) received for origin requests made by all sources.
	AllStatus501 *int32 `json:"all_status_501,omitempty"`
	// Number of responses received with status code 502 (Bad Gateway) received for origin requests made by all sources.
	AllStatus502 *int32 `json:"all_status_502,omitempty"`
	// Number of responses received with status code 503 (Service Unavailable) received for origin requests made by all sources.
	AllStatus503 *int32 `json:"all_status_503,omitempty"`
	// Number of responses received with status code 504 (Gateway Timeout) received for origin requests made by all sources.
	AllStatus504 *int32 `json:"all_status_504,omitempty"`
	// Number of responses received with status code 505 (HTTP Version Not Supported) received for origin requests made by all sources.
	AllStatus505 *int32 `json:"all_status_505,omitempty"`
	// Number of responses with latency between 0 and 1 millisecond received for origin requests made by all sources.
	AllLatency0To1ms *int32 `json:"all_latency_0_to_1ms,omitempty"`
	// Number of responses with latency between 1 and 5 milliseconds received for origin requests made by all sources.
	AllLatency1To5ms *int32 `json:"all_latency_1_to_5ms,omitempty"`
	// Number of responses with latency between 5 and 10 milliseconds received for origin requests made by all sources.
	AllLatency5To10ms *int32 `json:"all_latency_5_to_10ms,omitempty"`
	// Number of responses with latency between 10 and 50 milliseconds received for origin requests made by all sources.
	AllLatency10To50ms *int32 `json:"all_latency_10_to_50ms,omitempty"`
	// Number of responses with latency between 50 and 100 milliseconds received for origin requests made by all sources.
	AllLatency50To100ms *int32 `json:"all_latency_50_to_100ms,omitempty"`
	// Number of responses with latency between 100 and 250 milliseconds received for origin requests made by all sources.
	AllLatency100To250ms *int32 `json:"all_latency_100_to_250ms,omitempty"`
	// Number of responses with latency between 250 and 500 milliseconds received for origin requests made by all sources.
	AllLatency250To500ms *int32 `json:"all_latency_250_to_500ms,omitempty"`
	// Number of responses with latency between 500 and 1,000 milliseconds received for origin requests made by all sources.
	AllLatency500To1000ms *int32 `json:"all_latency_500_to_1000ms,omitempty"`
	// Number of responses with latency between 1,000 and 5,000 milliseconds received for origin requests made by all sources.
	AllLatency1000To5000ms *int32 `json:"all_latency_1000_to_5000ms,omitempty"`
	// Number of responses with latency between 5,000 and 10,000 milliseconds received for origin requests made by all sources.
	AllLatency5000To10000ms *int32 `json:"all_latency_5000_to_10000ms,omitempty"`
	// Number of responses with latency between 10,000 and 60,000 milliseconds received for origin requests made by all sources.
	AllLatency10000To60000ms *int32 `json:"all_latency_10000_to_60000ms,omitempty"`
	// Number of responses with latency of 60,000 milliseconds and above received for origin requests made by all sources.
	AllLatency60000ms *int32 `json:"all_latency_60000ms,omitempty"`
	AdditionalProperties map[string]any
}

type _OriginInspectorMeasurements OriginInspectorMeasurements

// NewOriginInspectorMeasurements instantiates a new OriginInspectorMeasurements object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOriginInspectorMeasurements() *OriginInspectorMeasurements {
	this := OriginInspectorMeasurements{}
	return &this
}

// NewOriginInspectorMeasurementsWithDefaults instantiates a new OriginInspectorMeasurements object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOriginInspectorMeasurementsWithDefaults() *OriginInspectorMeasurements {
	this := OriginInspectorMeasurements{}
	return &this
}

// GetResponses returns the Responses field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetResponses() int32 {
	if o == nil || o.Responses == nil {
		var ret int32
		return ret
	}
	return *o.Responses
}

// GetResponsesOk returns a tuple with the Responses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetResponsesOk() (*int32, bool) {
	if o == nil || o.Responses == nil {
		return nil, false
	}
	return o.Responses, true
}

// HasResponses returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasResponses() bool {
	if o != nil && o.Responses != nil {
		return true
	}

	return false
}

// SetResponses gets a reference to the given int32 and assigns it to the Responses field.
func (o *OriginInspectorMeasurements) SetResponses(v int32) {
	o.Responses = &v
}

// GetRespHeaderBytes returns the RespHeaderBytes field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetRespHeaderBytes() int32 {
	if o == nil || o.RespHeaderBytes == nil {
		var ret int32
		return ret
	}
	return *o.RespHeaderBytes
}

// GetRespHeaderBytesOk returns a tuple with the RespHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetRespHeaderBytesOk() (*int32, bool) {
	if o == nil || o.RespHeaderBytes == nil {
		return nil, false
	}
	return o.RespHeaderBytes, true
}

// HasRespHeaderBytes returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasRespHeaderBytes() bool {
	if o != nil && o.RespHeaderBytes != nil {
		return true
	}

	return false
}

// SetRespHeaderBytes gets a reference to the given int32 and assigns it to the RespHeaderBytes field.
func (o *OriginInspectorMeasurements) SetRespHeaderBytes(v int32) {
	o.RespHeaderBytes = &v
}

// GetRespBodyBytes returns the RespBodyBytes field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetRespBodyBytes() int32 {
	if o == nil || o.RespBodyBytes == nil {
		var ret int32
		return ret
	}
	return *o.RespBodyBytes
}

// GetRespBodyBytesOk returns a tuple with the RespBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetRespBodyBytesOk() (*int32, bool) {
	if o == nil || o.RespBodyBytes == nil {
		return nil, false
	}
	return o.RespBodyBytes, true
}

// HasRespBodyBytes returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasRespBodyBytes() bool {
	if o != nil && o.RespBodyBytes != nil {
		return true
	}

	return false
}

// SetRespBodyBytes gets a reference to the given int32 and assigns it to the RespBodyBytes field.
func (o *OriginInspectorMeasurements) SetRespBodyBytes(v int32) {
	o.RespBodyBytes = &v
}

// GetStatus1xx returns the Status1xx field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetStatus1xx() int32 {
	if o == nil || o.Status1xx == nil {
		var ret int32
		return ret
	}
	return *o.Status1xx
}

// GetStatus1xxOk returns a tuple with the Status1xx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetStatus1xxOk() (*int32, bool) {
	if o == nil || o.Status1xx == nil {
		return nil, false
	}
	return o.Status1xx, true
}

// HasStatus1xx returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasStatus1xx() bool {
	if o != nil && o.Status1xx != nil {
		return true
	}

	return false
}

// SetStatus1xx gets a reference to the given int32 and assigns it to the Status1xx field.
func (o *OriginInspectorMeasurements) SetStatus1xx(v int32) {
	o.Status1xx = &v
}

// GetStatus2xx returns the Status2xx field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetStatus2xx() int32 {
	if o == nil || o.Status2xx == nil {
		var ret int32
		return ret
	}
	return *o.Status2xx
}

// GetStatus2xxOk returns a tuple with the Status2xx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetStatus2xxOk() (*int32, bool) {
	if o == nil || o.Status2xx == nil {
		return nil, false
	}
	return o.Status2xx, true
}

// HasStatus2xx returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasStatus2xx() bool {
	if o != nil && o.Status2xx != nil {
		return true
	}

	return false
}

// SetStatus2xx gets a reference to the given int32 and assigns it to the Status2xx field.
func (o *OriginInspectorMeasurements) SetStatus2xx(v int32) {
	o.Status2xx = &v
}

// GetStatus3xx returns the Status3xx field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetStatus3xx() int32 {
	if o == nil || o.Status3xx == nil {
		var ret int32
		return ret
	}
	return *o.Status3xx
}

// GetStatus3xxOk returns a tuple with the Status3xx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetStatus3xxOk() (*int32, bool) {
	if o == nil || o.Status3xx == nil {
		return nil, false
	}
	return o.Status3xx, true
}

// HasStatus3xx returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasStatus3xx() bool {
	if o != nil && o.Status3xx != nil {
		return true
	}

	return false
}

// SetStatus3xx gets a reference to the given int32 and assigns it to the Status3xx field.
func (o *OriginInspectorMeasurements) SetStatus3xx(v int32) {
	o.Status3xx = &v
}

// GetStatus4xx returns the Status4xx field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetStatus4xx() int32 {
	if o == nil || o.Status4xx == nil {
		var ret int32
		return ret
	}
	return *o.Status4xx
}

// GetStatus4xxOk returns a tuple with the Status4xx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetStatus4xxOk() (*int32, bool) {
	if o == nil || o.Status4xx == nil {
		return nil, false
	}
	return o.Status4xx, true
}

// HasStatus4xx returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasStatus4xx() bool {
	if o != nil && o.Status4xx != nil {
		return true
	}

	return false
}

// SetStatus4xx gets a reference to the given int32 and assigns it to the Status4xx field.
func (o *OriginInspectorMeasurements) SetStatus4xx(v int32) {
	o.Status4xx = &v
}

// GetStatus5xx returns the Status5xx field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetStatus5xx() int32 {
	if o == nil || o.Status5xx == nil {
		var ret int32
		return ret
	}
	return *o.Status5xx
}

// GetStatus5xxOk returns a tuple with the Status5xx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetStatus5xxOk() (*int32, bool) {
	if o == nil || o.Status5xx == nil {
		return nil, false
	}
	return o.Status5xx, true
}

// HasStatus5xx returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasStatus5xx() bool {
	if o != nil && o.Status5xx != nil {
		return true
	}

	return false
}

// SetStatus5xx gets a reference to the given int32 and assigns it to the Status5xx field.
func (o *OriginInspectorMeasurements) SetStatus5xx(v int32) {
	o.Status5xx = &v
}

// GetStatus200 returns the Status200 field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetStatus200() int32 {
	if o == nil || o.Status200 == nil {
		var ret int32
		return ret
	}
	return *o.Status200
}

// GetStatus200Ok returns a tuple with the Status200 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetStatus200Ok() (*int32, bool) {
	if o == nil || o.Status200 == nil {
		return nil, false
	}
	return o.Status200, true
}

// HasStatus200 returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasStatus200() bool {
	if o != nil && o.Status200 != nil {
		return true
	}

	return false
}

// SetStatus200 gets a reference to the given int32 and assigns it to the Status200 field.
func (o *OriginInspectorMeasurements) SetStatus200(v int32) {
	o.Status200 = &v
}

// GetStatus204 returns the Status204 field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetStatus204() int32 {
	if o == nil || o.Status204 == nil {
		var ret int32
		return ret
	}
	return *o.Status204
}

// GetStatus204Ok returns a tuple with the Status204 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetStatus204Ok() (*int32, bool) {
	if o == nil || o.Status204 == nil {
		return nil, false
	}
	return o.Status204, true
}

// HasStatus204 returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasStatus204() bool {
	if o != nil && o.Status204 != nil {
		return true
	}

	return false
}

// SetStatus204 gets a reference to the given int32 and assigns it to the Status204 field.
func (o *OriginInspectorMeasurements) SetStatus204(v int32) {
	o.Status204 = &v
}

// GetStatus206 returns the Status206 field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetStatus206() int32 {
	if o == nil || o.Status206 == nil {
		var ret int32
		return ret
	}
	return *o.Status206
}

// GetStatus206Ok returns a tuple with the Status206 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetStatus206Ok() (*int32, bool) {
	if o == nil || o.Status206 == nil {
		return nil, false
	}
	return o.Status206, true
}

// HasStatus206 returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasStatus206() bool {
	if o != nil && o.Status206 != nil {
		return true
	}

	return false
}

// SetStatus206 gets a reference to the given int32 and assigns it to the Status206 field.
func (o *OriginInspectorMeasurements) SetStatus206(v int32) {
	o.Status206 = &v
}

// GetStatus301 returns the Status301 field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetStatus301() int32 {
	if o == nil || o.Status301 == nil {
		var ret int32
		return ret
	}
	return *o.Status301
}

// GetStatus301Ok returns a tuple with the Status301 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetStatus301Ok() (*int32, bool) {
	if o == nil || o.Status301 == nil {
		return nil, false
	}
	return o.Status301, true
}

// HasStatus301 returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasStatus301() bool {
	if o != nil && o.Status301 != nil {
		return true
	}

	return false
}

// SetStatus301 gets a reference to the given int32 and assigns it to the Status301 field.
func (o *OriginInspectorMeasurements) SetStatus301(v int32) {
	o.Status301 = &v
}

// GetStatus302 returns the Status302 field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetStatus302() int32 {
	if o == nil || o.Status302 == nil {
		var ret int32
		return ret
	}
	return *o.Status302
}

// GetStatus302Ok returns a tuple with the Status302 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetStatus302Ok() (*int32, bool) {
	if o == nil || o.Status302 == nil {
		return nil, false
	}
	return o.Status302, true
}

// HasStatus302 returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasStatus302() bool {
	if o != nil && o.Status302 != nil {
		return true
	}

	return false
}

// SetStatus302 gets a reference to the given int32 and assigns it to the Status302 field.
func (o *OriginInspectorMeasurements) SetStatus302(v int32) {
	o.Status302 = &v
}

// GetStatus304 returns the Status304 field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetStatus304() int32 {
	if o == nil || o.Status304 == nil {
		var ret int32
		return ret
	}
	return *o.Status304
}

// GetStatus304Ok returns a tuple with the Status304 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetStatus304Ok() (*int32, bool) {
	if o == nil || o.Status304 == nil {
		return nil, false
	}
	return o.Status304, true
}

// HasStatus304 returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasStatus304() bool {
	if o != nil && o.Status304 != nil {
		return true
	}

	return false
}

// SetStatus304 gets a reference to the given int32 and assigns it to the Status304 field.
func (o *OriginInspectorMeasurements) SetStatus304(v int32) {
	o.Status304 = &v
}

// GetStatus400 returns the Status400 field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetStatus400() int32 {
	if o == nil || o.Status400 == nil {
		var ret int32
		return ret
	}
	return *o.Status400
}

// GetStatus400Ok returns a tuple with the Status400 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetStatus400Ok() (*int32, bool) {
	if o == nil || o.Status400 == nil {
		return nil, false
	}
	return o.Status400, true
}

// HasStatus400 returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasStatus400() bool {
	if o != nil && o.Status400 != nil {
		return true
	}

	return false
}

// SetStatus400 gets a reference to the given int32 and assigns it to the Status400 field.
func (o *OriginInspectorMeasurements) SetStatus400(v int32) {
	o.Status400 = &v
}

// GetStatus401 returns the Status401 field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetStatus401() int32 {
	if o == nil || o.Status401 == nil {
		var ret int32
		return ret
	}
	return *o.Status401
}

// GetStatus401Ok returns a tuple with the Status401 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetStatus401Ok() (*int32, bool) {
	if o == nil || o.Status401 == nil {
		return nil, false
	}
	return o.Status401, true
}

// HasStatus401 returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasStatus401() bool {
	if o != nil && o.Status401 != nil {
		return true
	}

	return false
}

// SetStatus401 gets a reference to the given int32 and assigns it to the Status401 field.
func (o *OriginInspectorMeasurements) SetStatus401(v int32) {
	o.Status401 = &v
}

// GetStatus403 returns the Status403 field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetStatus403() int32 {
	if o == nil || o.Status403 == nil {
		var ret int32
		return ret
	}
	return *o.Status403
}

// GetStatus403Ok returns a tuple with the Status403 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetStatus403Ok() (*int32, bool) {
	if o == nil || o.Status403 == nil {
		return nil, false
	}
	return o.Status403, true
}

// HasStatus403 returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasStatus403() bool {
	if o != nil && o.Status403 != nil {
		return true
	}

	return false
}

// SetStatus403 gets a reference to the given int32 and assigns it to the Status403 field.
func (o *OriginInspectorMeasurements) SetStatus403(v int32) {
	o.Status403 = &v
}

// GetStatus404 returns the Status404 field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetStatus404() int32 {
	if o == nil || o.Status404 == nil {
		var ret int32
		return ret
	}
	return *o.Status404
}

// GetStatus404Ok returns a tuple with the Status404 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetStatus404Ok() (*int32, bool) {
	if o == nil || o.Status404 == nil {
		return nil, false
	}
	return o.Status404, true
}

// HasStatus404 returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasStatus404() bool {
	if o != nil && o.Status404 != nil {
		return true
	}

	return false
}

// SetStatus404 gets a reference to the given int32 and assigns it to the Status404 field.
func (o *OriginInspectorMeasurements) SetStatus404(v int32) {
	o.Status404 = &v
}

// GetStatus416 returns the Status416 field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetStatus416() int32 {
	if o == nil || o.Status416 == nil {
		var ret int32
		return ret
	}
	return *o.Status416
}

// GetStatus416Ok returns a tuple with the Status416 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetStatus416Ok() (*int32, bool) {
	if o == nil || o.Status416 == nil {
		return nil, false
	}
	return o.Status416, true
}

// HasStatus416 returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasStatus416() bool {
	if o != nil && o.Status416 != nil {
		return true
	}

	return false
}

// SetStatus416 gets a reference to the given int32 and assigns it to the Status416 field.
func (o *OriginInspectorMeasurements) SetStatus416(v int32) {
	o.Status416 = &v
}

// GetStatus429 returns the Status429 field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetStatus429() int32 {
	if o == nil || o.Status429 == nil {
		var ret int32
		return ret
	}
	return *o.Status429
}

// GetStatus429Ok returns a tuple with the Status429 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetStatus429Ok() (*int32, bool) {
	if o == nil || o.Status429 == nil {
		return nil, false
	}
	return o.Status429, true
}

// HasStatus429 returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasStatus429() bool {
	if o != nil && o.Status429 != nil {
		return true
	}

	return false
}

// SetStatus429 gets a reference to the given int32 and assigns it to the Status429 field.
func (o *OriginInspectorMeasurements) SetStatus429(v int32) {
	o.Status429 = &v
}

// GetStatus500 returns the Status500 field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetStatus500() int32 {
	if o == nil || o.Status500 == nil {
		var ret int32
		return ret
	}
	return *o.Status500
}

// GetStatus500Ok returns a tuple with the Status500 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetStatus500Ok() (*int32, bool) {
	if o == nil || o.Status500 == nil {
		return nil, false
	}
	return o.Status500, true
}

// HasStatus500 returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasStatus500() bool {
	if o != nil && o.Status500 != nil {
		return true
	}

	return false
}

// SetStatus500 gets a reference to the given int32 and assigns it to the Status500 field.
func (o *OriginInspectorMeasurements) SetStatus500(v int32) {
	o.Status500 = &v
}

// GetStatus501 returns the Status501 field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetStatus501() int32 {
	if o == nil || o.Status501 == nil {
		var ret int32
		return ret
	}
	return *o.Status501
}

// GetStatus501Ok returns a tuple with the Status501 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetStatus501Ok() (*int32, bool) {
	if o == nil || o.Status501 == nil {
		return nil, false
	}
	return o.Status501, true
}

// HasStatus501 returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasStatus501() bool {
	if o != nil && o.Status501 != nil {
		return true
	}

	return false
}

// SetStatus501 gets a reference to the given int32 and assigns it to the Status501 field.
func (o *OriginInspectorMeasurements) SetStatus501(v int32) {
	o.Status501 = &v
}

// GetStatus502 returns the Status502 field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetStatus502() int32 {
	if o == nil || o.Status502 == nil {
		var ret int32
		return ret
	}
	return *o.Status502
}

// GetStatus502Ok returns a tuple with the Status502 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetStatus502Ok() (*int32, bool) {
	if o == nil || o.Status502 == nil {
		return nil, false
	}
	return o.Status502, true
}

// HasStatus502 returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasStatus502() bool {
	if o != nil && o.Status502 != nil {
		return true
	}

	return false
}

// SetStatus502 gets a reference to the given int32 and assigns it to the Status502 field.
func (o *OriginInspectorMeasurements) SetStatus502(v int32) {
	o.Status502 = &v
}

// GetStatus503 returns the Status503 field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetStatus503() int32 {
	if o == nil || o.Status503 == nil {
		var ret int32
		return ret
	}
	return *o.Status503
}

// GetStatus503Ok returns a tuple with the Status503 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetStatus503Ok() (*int32, bool) {
	if o == nil || o.Status503 == nil {
		return nil, false
	}
	return o.Status503, true
}

// HasStatus503 returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasStatus503() bool {
	if o != nil && o.Status503 != nil {
		return true
	}

	return false
}

// SetStatus503 gets a reference to the given int32 and assigns it to the Status503 field.
func (o *OriginInspectorMeasurements) SetStatus503(v int32) {
	o.Status503 = &v
}

// GetStatus504 returns the Status504 field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetStatus504() int32 {
	if o == nil || o.Status504 == nil {
		var ret int32
		return ret
	}
	return *o.Status504
}

// GetStatus504Ok returns a tuple with the Status504 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetStatus504Ok() (*int32, bool) {
	if o == nil || o.Status504 == nil {
		return nil, false
	}
	return o.Status504, true
}

// HasStatus504 returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasStatus504() bool {
	if o != nil && o.Status504 != nil {
		return true
	}

	return false
}

// SetStatus504 gets a reference to the given int32 and assigns it to the Status504 field.
func (o *OriginInspectorMeasurements) SetStatus504(v int32) {
	o.Status504 = &v
}

// GetStatus505 returns the Status505 field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetStatus505() int32 {
	if o == nil || o.Status505 == nil {
		var ret int32
		return ret
	}
	return *o.Status505
}

// GetStatus505Ok returns a tuple with the Status505 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetStatus505Ok() (*int32, bool) {
	if o == nil || o.Status505 == nil {
		return nil, false
	}
	return o.Status505, true
}

// HasStatus505 returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasStatus505() bool {
	if o != nil && o.Status505 != nil {
		return true
	}

	return false
}

// SetStatus505 gets a reference to the given int32 and assigns it to the Status505 field.
func (o *OriginInspectorMeasurements) SetStatus505(v int32) {
	o.Status505 = &v
}

// GetLatency0To1ms returns the Latency0To1ms field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetLatency0To1ms() int32 {
	if o == nil || o.Latency0To1ms == nil {
		var ret int32
		return ret
	}
	return *o.Latency0To1ms
}

// GetLatency0To1msOk returns a tuple with the Latency0To1ms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetLatency0To1msOk() (*int32, bool) {
	if o == nil || o.Latency0To1ms == nil {
		return nil, false
	}
	return o.Latency0To1ms, true
}

// HasLatency0To1ms returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasLatency0To1ms() bool {
	if o != nil && o.Latency0To1ms != nil {
		return true
	}

	return false
}

// SetLatency0To1ms gets a reference to the given int32 and assigns it to the Latency0To1ms field.
func (o *OriginInspectorMeasurements) SetLatency0To1ms(v int32) {
	o.Latency0To1ms = &v
}

// GetLatency1To5ms returns the Latency1To5ms field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetLatency1To5ms() int32 {
	if o == nil || o.Latency1To5ms == nil {
		var ret int32
		return ret
	}
	return *o.Latency1To5ms
}

// GetLatency1To5msOk returns a tuple with the Latency1To5ms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetLatency1To5msOk() (*int32, bool) {
	if o == nil || o.Latency1To5ms == nil {
		return nil, false
	}
	return o.Latency1To5ms, true
}

// HasLatency1To5ms returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasLatency1To5ms() bool {
	if o != nil && o.Latency1To5ms != nil {
		return true
	}

	return false
}

// SetLatency1To5ms gets a reference to the given int32 and assigns it to the Latency1To5ms field.
func (o *OriginInspectorMeasurements) SetLatency1To5ms(v int32) {
	o.Latency1To5ms = &v
}

// GetLatency5To10ms returns the Latency5To10ms field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetLatency5To10ms() int32 {
	if o == nil || o.Latency5To10ms == nil {
		var ret int32
		return ret
	}
	return *o.Latency5To10ms
}

// GetLatency5To10msOk returns a tuple with the Latency5To10ms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetLatency5To10msOk() (*int32, bool) {
	if o == nil || o.Latency5To10ms == nil {
		return nil, false
	}
	return o.Latency5To10ms, true
}

// HasLatency5To10ms returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasLatency5To10ms() bool {
	if o != nil && o.Latency5To10ms != nil {
		return true
	}

	return false
}

// SetLatency5To10ms gets a reference to the given int32 and assigns it to the Latency5To10ms field.
func (o *OriginInspectorMeasurements) SetLatency5To10ms(v int32) {
	o.Latency5To10ms = &v
}

// GetLatency10To50ms returns the Latency10To50ms field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetLatency10To50ms() int32 {
	if o == nil || o.Latency10To50ms == nil {
		var ret int32
		return ret
	}
	return *o.Latency10To50ms
}

// GetLatency10To50msOk returns a tuple with the Latency10To50ms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetLatency10To50msOk() (*int32, bool) {
	if o == nil || o.Latency10To50ms == nil {
		return nil, false
	}
	return o.Latency10To50ms, true
}

// HasLatency10To50ms returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasLatency10To50ms() bool {
	if o != nil && o.Latency10To50ms != nil {
		return true
	}

	return false
}

// SetLatency10To50ms gets a reference to the given int32 and assigns it to the Latency10To50ms field.
func (o *OriginInspectorMeasurements) SetLatency10To50ms(v int32) {
	o.Latency10To50ms = &v
}

// GetLatency50To100ms returns the Latency50To100ms field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetLatency50To100ms() int32 {
	if o == nil || o.Latency50To100ms == nil {
		var ret int32
		return ret
	}
	return *o.Latency50To100ms
}

// GetLatency50To100msOk returns a tuple with the Latency50To100ms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetLatency50To100msOk() (*int32, bool) {
	if o == nil || o.Latency50To100ms == nil {
		return nil, false
	}
	return o.Latency50To100ms, true
}

// HasLatency50To100ms returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasLatency50To100ms() bool {
	if o != nil && o.Latency50To100ms != nil {
		return true
	}

	return false
}

// SetLatency50To100ms gets a reference to the given int32 and assigns it to the Latency50To100ms field.
func (o *OriginInspectorMeasurements) SetLatency50To100ms(v int32) {
	o.Latency50To100ms = &v
}

// GetLatency100To250ms returns the Latency100To250ms field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetLatency100To250ms() int32 {
	if o == nil || o.Latency100To250ms == nil {
		var ret int32
		return ret
	}
	return *o.Latency100To250ms
}

// GetLatency100To250msOk returns a tuple with the Latency100To250ms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetLatency100To250msOk() (*int32, bool) {
	if o == nil || o.Latency100To250ms == nil {
		return nil, false
	}
	return o.Latency100To250ms, true
}

// HasLatency100To250ms returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasLatency100To250ms() bool {
	if o != nil && o.Latency100To250ms != nil {
		return true
	}

	return false
}

// SetLatency100To250ms gets a reference to the given int32 and assigns it to the Latency100To250ms field.
func (o *OriginInspectorMeasurements) SetLatency100To250ms(v int32) {
	o.Latency100To250ms = &v
}

// GetLatency250To500ms returns the Latency250To500ms field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetLatency250To500ms() int32 {
	if o == nil || o.Latency250To500ms == nil {
		var ret int32
		return ret
	}
	return *o.Latency250To500ms
}

// GetLatency250To500msOk returns a tuple with the Latency250To500ms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetLatency250To500msOk() (*int32, bool) {
	if o == nil || o.Latency250To500ms == nil {
		return nil, false
	}
	return o.Latency250To500ms, true
}

// HasLatency250To500ms returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasLatency250To500ms() bool {
	if o != nil && o.Latency250To500ms != nil {
		return true
	}

	return false
}

// SetLatency250To500ms gets a reference to the given int32 and assigns it to the Latency250To500ms field.
func (o *OriginInspectorMeasurements) SetLatency250To500ms(v int32) {
	o.Latency250To500ms = &v
}

// GetLatency500To1000ms returns the Latency500To1000ms field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetLatency500To1000ms() int32 {
	if o == nil || o.Latency500To1000ms == nil {
		var ret int32
		return ret
	}
	return *o.Latency500To1000ms
}

// GetLatency500To1000msOk returns a tuple with the Latency500To1000ms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetLatency500To1000msOk() (*int32, bool) {
	if o == nil || o.Latency500To1000ms == nil {
		return nil, false
	}
	return o.Latency500To1000ms, true
}

// HasLatency500To1000ms returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasLatency500To1000ms() bool {
	if o != nil && o.Latency500To1000ms != nil {
		return true
	}

	return false
}

// SetLatency500To1000ms gets a reference to the given int32 and assigns it to the Latency500To1000ms field.
func (o *OriginInspectorMeasurements) SetLatency500To1000ms(v int32) {
	o.Latency500To1000ms = &v
}

// GetLatency1000To5000ms returns the Latency1000To5000ms field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetLatency1000To5000ms() int32 {
	if o == nil || o.Latency1000To5000ms == nil {
		var ret int32
		return ret
	}
	return *o.Latency1000To5000ms
}

// GetLatency1000To5000msOk returns a tuple with the Latency1000To5000ms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetLatency1000To5000msOk() (*int32, bool) {
	if o == nil || o.Latency1000To5000ms == nil {
		return nil, false
	}
	return o.Latency1000To5000ms, true
}

// HasLatency1000To5000ms returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasLatency1000To5000ms() bool {
	if o != nil && o.Latency1000To5000ms != nil {
		return true
	}

	return false
}

// SetLatency1000To5000ms gets a reference to the given int32 and assigns it to the Latency1000To5000ms field.
func (o *OriginInspectorMeasurements) SetLatency1000To5000ms(v int32) {
	o.Latency1000To5000ms = &v
}

// GetLatency5000To10000ms returns the Latency5000To10000ms field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetLatency5000To10000ms() int32 {
	if o == nil || o.Latency5000To10000ms == nil {
		var ret int32
		return ret
	}
	return *o.Latency5000To10000ms
}

// GetLatency5000To10000msOk returns a tuple with the Latency5000To10000ms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetLatency5000To10000msOk() (*int32, bool) {
	if o == nil || o.Latency5000To10000ms == nil {
		return nil, false
	}
	return o.Latency5000To10000ms, true
}

// HasLatency5000To10000ms returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasLatency5000To10000ms() bool {
	if o != nil && o.Latency5000To10000ms != nil {
		return true
	}

	return false
}

// SetLatency5000To10000ms gets a reference to the given int32 and assigns it to the Latency5000To10000ms field.
func (o *OriginInspectorMeasurements) SetLatency5000To10000ms(v int32) {
	o.Latency5000To10000ms = &v
}

// GetLatency10000To60000ms returns the Latency10000To60000ms field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetLatency10000To60000ms() int32 {
	if o == nil || o.Latency10000To60000ms == nil {
		var ret int32
		return ret
	}
	return *o.Latency10000To60000ms
}

// GetLatency10000To60000msOk returns a tuple with the Latency10000To60000ms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetLatency10000To60000msOk() (*int32, bool) {
	if o == nil || o.Latency10000To60000ms == nil {
		return nil, false
	}
	return o.Latency10000To60000ms, true
}

// HasLatency10000To60000ms returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasLatency10000To60000ms() bool {
	if o != nil && o.Latency10000To60000ms != nil {
		return true
	}

	return false
}

// SetLatency10000To60000ms gets a reference to the given int32 and assigns it to the Latency10000To60000ms field.
func (o *OriginInspectorMeasurements) SetLatency10000To60000ms(v int32) {
	o.Latency10000To60000ms = &v
}

// GetLatency60000ms returns the Latency60000ms field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetLatency60000ms() int32 {
	if o == nil || o.Latency60000ms == nil {
		var ret int32
		return ret
	}
	return *o.Latency60000ms
}

// GetLatency60000msOk returns a tuple with the Latency60000ms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetLatency60000msOk() (*int32, bool) {
	if o == nil || o.Latency60000ms == nil {
		return nil, false
	}
	return o.Latency60000ms, true
}

// HasLatency60000ms returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasLatency60000ms() bool {
	if o != nil && o.Latency60000ms != nil {
		return true
	}

	return false
}

// SetLatency60000ms gets a reference to the given int32 and assigns it to the Latency60000ms field.
func (o *OriginInspectorMeasurements) SetLatency60000ms(v int32) {
	o.Latency60000ms = &v
}

// GetWafResponses returns the WafResponses field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetWafResponses() int32 {
	if o == nil || o.WafResponses == nil {
		var ret int32
		return ret
	}
	return *o.WafResponses
}

// GetWafResponsesOk returns a tuple with the WafResponses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetWafResponsesOk() (*int32, bool) {
	if o == nil || o.WafResponses == nil {
		return nil, false
	}
	return o.WafResponses, true
}

// HasWafResponses returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasWafResponses() bool {
	if o != nil && o.WafResponses != nil {
		return true
	}

	return false
}

// SetWafResponses gets a reference to the given int32 and assigns it to the WafResponses field.
func (o *OriginInspectorMeasurements) SetWafResponses(v int32) {
	o.WafResponses = &v
}

// GetWafRespHeaderBytes returns the WafRespHeaderBytes field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetWafRespHeaderBytes() int32 {
	if o == nil || o.WafRespHeaderBytes == nil {
		var ret int32
		return ret
	}
	return *o.WafRespHeaderBytes
}

// GetWafRespHeaderBytesOk returns a tuple with the WafRespHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetWafRespHeaderBytesOk() (*int32, bool) {
	if o == nil || o.WafRespHeaderBytes == nil {
		return nil, false
	}
	return o.WafRespHeaderBytes, true
}

// HasWafRespHeaderBytes returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasWafRespHeaderBytes() bool {
	if o != nil && o.WafRespHeaderBytes != nil {
		return true
	}

	return false
}

// SetWafRespHeaderBytes gets a reference to the given int32 and assigns it to the WafRespHeaderBytes field.
func (o *OriginInspectorMeasurements) SetWafRespHeaderBytes(v int32) {
	o.WafRespHeaderBytes = &v
}

// GetWafRespBodyBytes returns the WafRespBodyBytes field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetWafRespBodyBytes() int32 {
	if o == nil || o.WafRespBodyBytes == nil {
		var ret int32
		return ret
	}
	return *o.WafRespBodyBytes
}

// GetWafRespBodyBytesOk returns a tuple with the WafRespBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetWafRespBodyBytesOk() (*int32, bool) {
	if o == nil || o.WafRespBodyBytes == nil {
		return nil, false
	}
	return o.WafRespBodyBytes, true
}

// HasWafRespBodyBytes returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasWafRespBodyBytes() bool {
	if o != nil && o.WafRespBodyBytes != nil {
		return true
	}

	return false
}

// SetWafRespBodyBytes gets a reference to the given int32 and assigns it to the WafRespBodyBytes field.
func (o *OriginInspectorMeasurements) SetWafRespBodyBytes(v int32) {
	o.WafRespBodyBytes = &v
}

// GetWafStatus1xx returns the WafStatus1xx field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetWafStatus1xx() int32 {
	if o == nil || o.WafStatus1xx == nil {
		var ret int32
		return ret
	}
	return *o.WafStatus1xx
}

// GetWafStatus1xxOk returns a tuple with the WafStatus1xx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetWafStatus1xxOk() (*int32, bool) {
	if o == nil || o.WafStatus1xx == nil {
		return nil, false
	}
	return o.WafStatus1xx, true
}

// HasWafStatus1xx returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasWafStatus1xx() bool {
	if o != nil && o.WafStatus1xx != nil {
		return true
	}

	return false
}

// SetWafStatus1xx gets a reference to the given int32 and assigns it to the WafStatus1xx field.
func (o *OriginInspectorMeasurements) SetWafStatus1xx(v int32) {
	o.WafStatus1xx = &v
}

// GetWafStatus2xx returns the WafStatus2xx field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetWafStatus2xx() int32 {
	if o == nil || o.WafStatus2xx == nil {
		var ret int32
		return ret
	}
	return *o.WafStatus2xx
}

// GetWafStatus2xxOk returns a tuple with the WafStatus2xx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetWafStatus2xxOk() (*int32, bool) {
	if o == nil || o.WafStatus2xx == nil {
		return nil, false
	}
	return o.WafStatus2xx, true
}

// HasWafStatus2xx returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasWafStatus2xx() bool {
	if o != nil && o.WafStatus2xx != nil {
		return true
	}

	return false
}

// SetWafStatus2xx gets a reference to the given int32 and assigns it to the WafStatus2xx field.
func (o *OriginInspectorMeasurements) SetWafStatus2xx(v int32) {
	o.WafStatus2xx = &v
}

// GetWafStatus3xx returns the WafStatus3xx field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetWafStatus3xx() int32 {
	if o == nil || o.WafStatus3xx == nil {
		var ret int32
		return ret
	}
	return *o.WafStatus3xx
}

// GetWafStatus3xxOk returns a tuple with the WafStatus3xx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetWafStatus3xxOk() (*int32, bool) {
	if o == nil || o.WafStatus3xx == nil {
		return nil, false
	}
	return o.WafStatus3xx, true
}

// HasWafStatus3xx returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasWafStatus3xx() bool {
	if o != nil && o.WafStatus3xx != nil {
		return true
	}

	return false
}

// SetWafStatus3xx gets a reference to the given int32 and assigns it to the WafStatus3xx field.
func (o *OriginInspectorMeasurements) SetWafStatus3xx(v int32) {
	o.WafStatus3xx = &v
}

// GetWafStatus4xx returns the WafStatus4xx field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetWafStatus4xx() int32 {
	if o == nil || o.WafStatus4xx == nil {
		var ret int32
		return ret
	}
	return *o.WafStatus4xx
}

// GetWafStatus4xxOk returns a tuple with the WafStatus4xx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetWafStatus4xxOk() (*int32, bool) {
	if o == nil || o.WafStatus4xx == nil {
		return nil, false
	}
	return o.WafStatus4xx, true
}

// HasWafStatus4xx returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasWafStatus4xx() bool {
	if o != nil && o.WafStatus4xx != nil {
		return true
	}

	return false
}

// SetWafStatus4xx gets a reference to the given int32 and assigns it to the WafStatus4xx field.
func (o *OriginInspectorMeasurements) SetWafStatus4xx(v int32) {
	o.WafStatus4xx = &v
}

// GetWafStatus5xx returns the WafStatus5xx field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetWafStatus5xx() int32 {
	if o == nil || o.WafStatus5xx == nil {
		var ret int32
		return ret
	}
	return *o.WafStatus5xx
}

// GetWafStatus5xxOk returns a tuple with the WafStatus5xx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetWafStatus5xxOk() (*int32, bool) {
	if o == nil || o.WafStatus5xx == nil {
		return nil, false
	}
	return o.WafStatus5xx, true
}

// HasWafStatus5xx returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasWafStatus5xx() bool {
	if o != nil && o.WafStatus5xx != nil {
		return true
	}

	return false
}

// SetWafStatus5xx gets a reference to the given int32 and assigns it to the WafStatus5xx field.
func (o *OriginInspectorMeasurements) SetWafStatus5xx(v int32) {
	o.WafStatus5xx = &v
}

// GetWafStatus200 returns the WafStatus200 field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetWafStatus200() int32 {
	if o == nil || o.WafStatus200 == nil {
		var ret int32
		return ret
	}
	return *o.WafStatus200
}

// GetWafStatus200Ok returns a tuple with the WafStatus200 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetWafStatus200Ok() (*int32, bool) {
	if o == nil || o.WafStatus200 == nil {
		return nil, false
	}
	return o.WafStatus200, true
}

// HasWafStatus200 returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasWafStatus200() bool {
	if o != nil && o.WafStatus200 != nil {
		return true
	}

	return false
}

// SetWafStatus200 gets a reference to the given int32 and assigns it to the WafStatus200 field.
func (o *OriginInspectorMeasurements) SetWafStatus200(v int32) {
	o.WafStatus200 = &v
}

// GetWafStatus204 returns the WafStatus204 field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetWafStatus204() int32 {
	if o == nil || o.WafStatus204 == nil {
		var ret int32
		return ret
	}
	return *o.WafStatus204
}

// GetWafStatus204Ok returns a tuple with the WafStatus204 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetWafStatus204Ok() (*int32, bool) {
	if o == nil || o.WafStatus204 == nil {
		return nil, false
	}
	return o.WafStatus204, true
}

// HasWafStatus204 returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasWafStatus204() bool {
	if o != nil && o.WafStatus204 != nil {
		return true
	}

	return false
}

// SetWafStatus204 gets a reference to the given int32 and assigns it to the WafStatus204 field.
func (o *OriginInspectorMeasurements) SetWafStatus204(v int32) {
	o.WafStatus204 = &v
}

// GetWafStatus206 returns the WafStatus206 field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetWafStatus206() int32 {
	if o == nil || o.WafStatus206 == nil {
		var ret int32
		return ret
	}
	return *o.WafStatus206
}

// GetWafStatus206Ok returns a tuple with the WafStatus206 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetWafStatus206Ok() (*int32, bool) {
	if o == nil || o.WafStatus206 == nil {
		return nil, false
	}
	return o.WafStatus206, true
}

// HasWafStatus206 returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasWafStatus206() bool {
	if o != nil && o.WafStatus206 != nil {
		return true
	}

	return false
}

// SetWafStatus206 gets a reference to the given int32 and assigns it to the WafStatus206 field.
func (o *OriginInspectorMeasurements) SetWafStatus206(v int32) {
	o.WafStatus206 = &v
}

// GetWafStatus301 returns the WafStatus301 field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetWafStatus301() int32 {
	if o == nil || o.WafStatus301 == nil {
		var ret int32
		return ret
	}
	return *o.WafStatus301
}

// GetWafStatus301Ok returns a tuple with the WafStatus301 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetWafStatus301Ok() (*int32, bool) {
	if o == nil || o.WafStatus301 == nil {
		return nil, false
	}
	return o.WafStatus301, true
}

// HasWafStatus301 returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasWafStatus301() bool {
	if o != nil && o.WafStatus301 != nil {
		return true
	}

	return false
}

// SetWafStatus301 gets a reference to the given int32 and assigns it to the WafStatus301 field.
func (o *OriginInspectorMeasurements) SetWafStatus301(v int32) {
	o.WafStatus301 = &v
}

// GetWafStatus302 returns the WafStatus302 field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetWafStatus302() int32 {
	if o == nil || o.WafStatus302 == nil {
		var ret int32
		return ret
	}
	return *o.WafStatus302
}

// GetWafStatus302Ok returns a tuple with the WafStatus302 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetWafStatus302Ok() (*int32, bool) {
	if o == nil || o.WafStatus302 == nil {
		return nil, false
	}
	return o.WafStatus302, true
}

// HasWafStatus302 returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasWafStatus302() bool {
	if o != nil && o.WafStatus302 != nil {
		return true
	}

	return false
}

// SetWafStatus302 gets a reference to the given int32 and assigns it to the WafStatus302 field.
func (o *OriginInspectorMeasurements) SetWafStatus302(v int32) {
	o.WafStatus302 = &v
}

// GetWafStatus304 returns the WafStatus304 field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetWafStatus304() int32 {
	if o == nil || o.WafStatus304 == nil {
		var ret int32
		return ret
	}
	return *o.WafStatus304
}

// GetWafStatus304Ok returns a tuple with the WafStatus304 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetWafStatus304Ok() (*int32, bool) {
	if o == nil || o.WafStatus304 == nil {
		return nil, false
	}
	return o.WafStatus304, true
}

// HasWafStatus304 returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasWafStatus304() bool {
	if o != nil && o.WafStatus304 != nil {
		return true
	}

	return false
}

// SetWafStatus304 gets a reference to the given int32 and assigns it to the WafStatus304 field.
func (o *OriginInspectorMeasurements) SetWafStatus304(v int32) {
	o.WafStatus304 = &v
}

// GetWafStatus400 returns the WafStatus400 field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetWafStatus400() int32 {
	if o == nil || o.WafStatus400 == nil {
		var ret int32
		return ret
	}
	return *o.WafStatus400
}

// GetWafStatus400Ok returns a tuple with the WafStatus400 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetWafStatus400Ok() (*int32, bool) {
	if o == nil || o.WafStatus400 == nil {
		return nil, false
	}
	return o.WafStatus400, true
}

// HasWafStatus400 returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasWafStatus400() bool {
	if o != nil && o.WafStatus400 != nil {
		return true
	}

	return false
}

// SetWafStatus400 gets a reference to the given int32 and assigns it to the WafStatus400 field.
func (o *OriginInspectorMeasurements) SetWafStatus400(v int32) {
	o.WafStatus400 = &v
}

// GetWafStatus401 returns the WafStatus401 field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetWafStatus401() int32 {
	if o == nil || o.WafStatus401 == nil {
		var ret int32
		return ret
	}
	return *o.WafStatus401
}

// GetWafStatus401Ok returns a tuple with the WafStatus401 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetWafStatus401Ok() (*int32, bool) {
	if o == nil || o.WafStatus401 == nil {
		return nil, false
	}
	return o.WafStatus401, true
}

// HasWafStatus401 returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasWafStatus401() bool {
	if o != nil && o.WafStatus401 != nil {
		return true
	}

	return false
}

// SetWafStatus401 gets a reference to the given int32 and assigns it to the WafStatus401 field.
func (o *OriginInspectorMeasurements) SetWafStatus401(v int32) {
	o.WafStatus401 = &v
}

// GetWafStatus403 returns the WafStatus403 field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetWafStatus403() int32 {
	if o == nil || o.WafStatus403 == nil {
		var ret int32
		return ret
	}
	return *o.WafStatus403
}

// GetWafStatus403Ok returns a tuple with the WafStatus403 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetWafStatus403Ok() (*int32, bool) {
	if o == nil || o.WafStatus403 == nil {
		return nil, false
	}
	return o.WafStatus403, true
}

// HasWafStatus403 returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasWafStatus403() bool {
	if o != nil && o.WafStatus403 != nil {
		return true
	}

	return false
}

// SetWafStatus403 gets a reference to the given int32 and assigns it to the WafStatus403 field.
func (o *OriginInspectorMeasurements) SetWafStatus403(v int32) {
	o.WafStatus403 = &v
}

// GetWafStatus404 returns the WafStatus404 field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetWafStatus404() int32 {
	if o == nil || o.WafStatus404 == nil {
		var ret int32
		return ret
	}
	return *o.WafStatus404
}

// GetWafStatus404Ok returns a tuple with the WafStatus404 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetWafStatus404Ok() (*int32, bool) {
	if o == nil || o.WafStatus404 == nil {
		return nil, false
	}
	return o.WafStatus404, true
}

// HasWafStatus404 returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasWafStatus404() bool {
	if o != nil && o.WafStatus404 != nil {
		return true
	}

	return false
}

// SetWafStatus404 gets a reference to the given int32 and assigns it to the WafStatus404 field.
func (o *OriginInspectorMeasurements) SetWafStatus404(v int32) {
	o.WafStatus404 = &v
}

// GetWafStatus416 returns the WafStatus416 field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetWafStatus416() int32 {
	if o == nil || o.WafStatus416 == nil {
		var ret int32
		return ret
	}
	return *o.WafStatus416
}

// GetWafStatus416Ok returns a tuple with the WafStatus416 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetWafStatus416Ok() (*int32, bool) {
	if o == nil || o.WafStatus416 == nil {
		return nil, false
	}
	return o.WafStatus416, true
}

// HasWafStatus416 returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasWafStatus416() bool {
	if o != nil && o.WafStatus416 != nil {
		return true
	}

	return false
}

// SetWafStatus416 gets a reference to the given int32 and assigns it to the WafStatus416 field.
func (o *OriginInspectorMeasurements) SetWafStatus416(v int32) {
	o.WafStatus416 = &v
}

// GetWafStatus429 returns the WafStatus429 field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetWafStatus429() int32 {
	if o == nil || o.WafStatus429 == nil {
		var ret int32
		return ret
	}
	return *o.WafStatus429
}

// GetWafStatus429Ok returns a tuple with the WafStatus429 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetWafStatus429Ok() (*int32, bool) {
	if o == nil || o.WafStatus429 == nil {
		return nil, false
	}
	return o.WafStatus429, true
}

// HasWafStatus429 returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasWafStatus429() bool {
	if o != nil && o.WafStatus429 != nil {
		return true
	}

	return false
}

// SetWafStatus429 gets a reference to the given int32 and assigns it to the WafStatus429 field.
func (o *OriginInspectorMeasurements) SetWafStatus429(v int32) {
	o.WafStatus429 = &v
}

// GetWafStatus500 returns the WafStatus500 field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetWafStatus500() int32 {
	if o == nil || o.WafStatus500 == nil {
		var ret int32
		return ret
	}
	return *o.WafStatus500
}

// GetWafStatus500Ok returns a tuple with the WafStatus500 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetWafStatus500Ok() (*int32, bool) {
	if o == nil || o.WafStatus500 == nil {
		return nil, false
	}
	return o.WafStatus500, true
}

// HasWafStatus500 returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasWafStatus500() bool {
	if o != nil && o.WafStatus500 != nil {
		return true
	}

	return false
}

// SetWafStatus500 gets a reference to the given int32 and assigns it to the WafStatus500 field.
func (o *OriginInspectorMeasurements) SetWafStatus500(v int32) {
	o.WafStatus500 = &v
}

// GetWafStatus501 returns the WafStatus501 field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetWafStatus501() int32 {
	if o == nil || o.WafStatus501 == nil {
		var ret int32
		return ret
	}
	return *o.WafStatus501
}

// GetWafStatus501Ok returns a tuple with the WafStatus501 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetWafStatus501Ok() (*int32, bool) {
	if o == nil || o.WafStatus501 == nil {
		return nil, false
	}
	return o.WafStatus501, true
}

// HasWafStatus501 returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasWafStatus501() bool {
	if o != nil && o.WafStatus501 != nil {
		return true
	}

	return false
}

// SetWafStatus501 gets a reference to the given int32 and assigns it to the WafStatus501 field.
func (o *OriginInspectorMeasurements) SetWafStatus501(v int32) {
	o.WafStatus501 = &v
}

// GetWafStatus502 returns the WafStatus502 field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetWafStatus502() int32 {
	if o == nil || o.WafStatus502 == nil {
		var ret int32
		return ret
	}
	return *o.WafStatus502
}

// GetWafStatus502Ok returns a tuple with the WafStatus502 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetWafStatus502Ok() (*int32, bool) {
	if o == nil || o.WafStatus502 == nil {
		return nil, false
	}
	return o.WafStatus502, true
}

// HasWafStatus502 returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasWafStatus502() bool {
	if o != nil && o.WafStatus502 != nil {
		return true
	}

	return false
}

// SetWafStatus502 gets a reference to the given int32 and assigns it to the WafStatus502 field.
func (o *OriginInspectorMeasurements) SetWafStatus502(v int32) {
	o.WafStatus502 = &v
}

// GetWafStatus503 returns the WafStatus503 field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetWafStatus503() int32 {
	if o == nil || o.WafStatus503 == nil {
		var ret int32
		return ret
	}
	return *o.WafStatus503
}

// GetWafStatus503Ok returns a tuple with the WafStatus503 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetWafStatus503Ok() (*int32, bool) {
	if o == nil || o.WafStatus503 == nil {
		return nil, false
	}
	return o.WafStatus503, true
}

// HasWafStatus503 returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasWafStatus503() bool {
	if o != nil && o.WafStatus503 != nil {
		return true
	}

	return false
}

// SetWafStatus503 gets a reference to the given int32 and assigns it to the WafStatus503 field.
func (o *OriginInspectorMeasurements) SetWafStatus503(v int32) {
	o.WafStatus503 = &v
}

// GetWafStatus504 returns the WafStatus504 field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetWafStatus504() int32 {
	if o == nil || o.WafStatus504 == nil {
		var ret int32
		return ret
	}
	return *o.WafStatus504
}

// GetWafStatus504Ok returns a tuple with the WafStatus504 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetWafStatus504Ok() (*int32, bool) {
	if o == nil || o.WafStatus504 == nil {
		return nil, false
	}
	return o.WafStatus504, true
}

// HasWafStatus504 returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasWafStatus504() bool {
	if o != nil && o.WafStatus504 != nil {
		return true
	}

	return false
}

// SetWafStatus504 gets a reference to the given int32 and assigns it to the WafStatus504 field.
func (o *OriginInspectorMeasurements) SetWafStatus504(v int32) {
	o.WafStatus504 = &v
}

// GetWafStatus505 returns the WafStatus505 field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetWafStatus505() int32 {
	if o == nil || o.WafStatus505 == nil {
		var ret int32
		return ret
	}
	return *o.WafStatus505
}

// GetWafStatus505Ok returns a tuple with the WafStatus505 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetWafStatus505Ok() (*int32, bool) {
	if o == nil || o.WafStatus505 == nil {
		return nil, false
	}
	return o.WafStatus505, true
}

// HasWafStatus505 returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasWafStatus505() bool {
	if o != nil && o.WafStatus505 != nil {
		return true
	}

	return false
}

// SetWafStatus505 gets a reference to the given int32 and assigns it to the WafStatus505 field.
func (o *OriginInspectorMeasurements) SetWafStatus505(v int32) {
	o.WafStatus505 = &v
}

// GetWafLatency0To1ms returns the WafLatency0To1ms field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetWafLatency0To1ms() int32 {
	if o == nil || o.WafLatency0To1ms == nil {
		var ret int32
		return ret
	}
	return *o.WafLatency0To1ms
}

// GetWafLatency0To1msOk returns a tuple with the WafLatency0To1ms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetWafLatency0To1msOk() (*int32, bool) {
	if o == nil || o.WafLatency0To1ms == nil {
		return nil, false
	}
	return o.WafLatency0To1ms, true
}

// HasWafLatency0To1ms returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasWafLatency0To1ms() bool {
	if o != nil && o.WafLatency0To1ms != nil {
		return true
	}

	return false
}

// SetWafLatency0To1ms gets a reference to the given int32 and assigns it to the WafLatency0To1ms field.
func (o *OriginInspectorMeasurements) SetWafLatency0To1ms(v int32) {
	o.WafLatency0To1ms = &v
}

// GetWafLatency1To5ms returns the WafLatency1To5ms field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetWafLatency1To5ms() int32 {
	if o == nil || o.WafLatency1To5ms == nil {
		var ret int32
		return ret
	}
	return *o.WafLatency1To5ms
}

// GetWafLatency1To5msOk returns a tuple with the WafLatency1To5ms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetWafLatency1To5msOk() (*int32, bool) {
	if o == nil || o.WafLatency1To5ms == nil {
		return nil, false
	}
	return o.WafLatency1To5ms, true
}

// HasWafLatency1To5ms returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasWafLatency1To5ms() bool {
	if o != nil && o.WafLatency1To5ms != nil {
		return true
	}

	return false
}

// SetWafLatency1To5ms gets a reference to the given int32 and assigns it to the WafLatency1To5ms field.
func (o *OriginInspectorMeasurements) SetWafLatency1To5ms(v int32) {
	o.WafLatency1To5ms = &v
}

// GetWafLatency5To10ms returns the WafLatency5To10ms field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetWafLatency5To10ms() int32 {
	if o == nil || o.WafLatency5To10ms == nil {
		var ret int32
		return ret
	}
	return *o.WafLatency5To10ms
}

// GetWafLatency5To10msOk returns a tuple with the WafLatency5To10ms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetWafLatency5To10msOk() (*int32, bool) {
	if o == nil || o.WafLatency5To10ms == nil {
		return nil, false
	}
	return o.WafLatency5To10ms, true
}

// HasWafLatency5To10ms returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasWafLatency5To10ms() bool {
	if o != nil && o.WafLatency5To10ms != nil {
		return true
	}

	return false
}

// SetWafLatency5To10ms gets a reference to the given int32 and assigns it to the WafLatency5To10ms field.
func (o *OriginInspectorMeasurements) SetWafLatency5To10ms(v int32) {
	o.WafLatency5To10ms = &v
}

// GetWafLatency10To50ms returns the WafLatency10To50ms field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetWafLatency10To50ms() int32 {
	if o == nil || o.WafLatency10To50ms == nil {
		var ret int32
		return ret
	}
	return *o.WafLatency10To50ms
}

// GetWafLatency10To50msOk returns a tuple with the WafLatency10To50ms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetWafLatency10To50msOk() (*int32, bool) {
	if o == nil || o.WafLatency10To50ms == nil {
		return nil, false
	}
	return o.WafLatency10To50ms, true
}

// HasWafLatency10To50ms returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasWafLatency10To50ms() bool {
	if o != nil && o.WafLatency10To50ms != nil {
		return true
	}

	return false
}

// SetWafLatency10To50ms gets a reference to the given int32 and assigns it to the WafLatency10To50ms field.
func (o *OriginInspectorMeasurements) SetWafLatency10To50ms(v int32) {
	o.WafLatency10To50ms = &v
}

// GetWafLatency50To100ms returns the WafLatency50To100ms field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetWafLatency50To100ms() int32 {
	if o == nil || o.WafLatency50To100ms == nil {
		var ret int32
		return ret
	}
	return *o.WafLatency50To100ms
}

// GetWafLatency50To100msOk returns a tuple with the WafLatency50To100ms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetWafLatency50To100msOk() (*int32, bool) {
	if o == nil || o.WafLatency50To100ms == nil {
		return nil, false
	}
	return o.WafLatency50To100ms, true
}

// HasWafLatency50To100ms returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasWafLatency50To100ms() bool {
	if o != nil && o.WafLatency50To100ms != nil {
		return true
	}

	return false
}

// SetWafLatency50To100ms gets a reference to the given int32 and assigns it to the WafLatency50To100ms field.
func (o *OriginInspectorMeasurements) SetWafLatency50To100ms(v int32) {
	o.WafLatency50To100ms = &v
}

// GetWafLatency100To250ms returns the WafLatency100To250ms field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetWafLatency100To250ms() int32 {
	if o == nil || o.WafLatency100To250ms == nil {
		var ret int32
		return ret
	}
	return *o.WafLatency100To250ms
}

// GetWafLatency100To250msOk returns a tuple with the WafLatency100To250ms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetWafLatency100To250msOk() (*int32, bool) {
	if o == nil || o.WafLatency100To250ms == nil {
		return nil, false
	}
	return o.WafLatency100To250ms, true
}

// HasWafLatency100To250ms returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasWafLatency100To250ms() bool {
	if o != nil && o.WafLatency100To250ms != nil {
		return true
	}

	return false
}

// SetWafLatency100To250ms gets a reference to the given int32 and assigns it to the WafLatency100To250ms field.
func (o *OriginInspectorMeasurements) SetWafLatency100To250ms(v int32) {
	o.WafLatency100To250ms = &v
}

// GetWafLatency250To500ms returns the WafLatency250To500ms field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetWafLatency250To500ms() int32 {
	if o == nil || o.WafLatency250To500ms == nil {
		var ret int32
		return ret
	}
	return *o.WafLatency250To500ms
}

// GetWafLatency250To500msOk returns a tuple with the WafLatency250To500ms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetWafLatency250To500msOk() (*int32, bool) {
	if o == nil || o.WafLatency250To500ms == nil {
		return nil, false
	}
	return o.WafLatency250To500ms, true
}

// HasWafLatency250To500ms returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasWafLatency250To500ms() bool {
	if o != nil && o.WafLatency250To500ms != nil {
		return true
	}

	return false
}

// SetWafLatency250To500ms gets a reference to the given int32 and assigns it to the WafLatency250To500ms field.
func (o *OriginInspectorMeasurements) SetWafLatency250To500ms(v int32) {
	o.WafLatency250To500ms = &v
}

// GetWafLatency500To1000ms returns the WafLatency500To1000ms field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetWafLatency500To1000ms() int32 {
	if o == nil || o.WafLatency500To1000ms == nil {
		var ret int32
		return ret
	}
	return *o.WafLatency500To1000ms
}

// GetWafLatency500To1000msOk returns a tuple with the WafLatency500To1000ms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetWafLatency500To1000msOk() (*int32, bool) {
	if o == nil || o.WafLatency500To1000ms == nil {
		return nil, false
	}
	return o.WafLatency500To1000ms, true
}

// HasWafLatency500To1000ms returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasWafLatency500To1000ms() bool {
	if o != nil && o.WafLatency500To1000ms != nil {
		return true
	}

	return false
}

// SetWafLatency500To1000ms gets a reference to the given int32 and assigns it to the WafLatency500To1000ms field.
func (o *OriginInspectorMeasurements) SetWafLatency500To1000ms(v int32) {
	o.WafLatency500To1000ms = &v
}

// GetWafLatency1000To5000ms returns the WafLatency1000To5000ms field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetWafLatency1000To5000ms() int32 {
	if o == nil || o.WafLatency1000To5000ms == nil {
		var ret int32
		return ret
	}
	return *o.WafLatency1000To5000ms
}

// GetWafLatency1000To5000msOk returns a tuple with the WafLatency1000To5000ms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetWafLatency1000To5000msOk() (*int32, bool) {
	if o == nil || o.WafLatency1000To5000ms == nil {
		return nil, false
	}
	return o.WafLatency1000To5000ms, true
}

// HasWafLatency1000To5000ms returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasWafLatency1000To5000ms() bool {
	if o != nil && o.WafLatency1000To5000ms != nil {
		return true
	}

	return false
}

// SetWafLatency1000To5000ms gets a reference to the given int32 and assigns it to the WafLatency1000To5000ms field.
func (o *OriginInspectorMeasurements) SetWafLatency1000To5000ms(v int32) {
	o.WafLatency1000To5000ms = &v
}

// GetWafLatency5000To10000ms returns the WafLatency5000To10000ms field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetWafLatency5000To10000ms() int32 {
	if o == nil || o.WafLatency5000To10000ms == nil {
		var ret int32
		return ret
	}
	return *o.WafLatency5000To10000ms
}

// GetWafLatency5000To10000msOk returns a tuple with the WafLatency5000To10000ms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetWafLatency5000To10000msOk() (*int32, bool) {
	if o == nil || o.WafLatency5000To10000ms == nil {
		return nil, false
	}
	return o.WafLatency5000To10000ms, true
}

// HasWafLatency5000To10000ms returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasWafLatency5000To10000ms() bool {
	if o != nil && o.WafLatency5000To10000ms != nil {
		return true
	}

	return false
}

// SetWafLatency5000To10000ms gets a reference to the given int32 and assigns it to the WafLatency5000To10000ms field.
func (o *OriginInspectorMeasurements) SetWafLatency5000To10000ms(v int32) {
	o.WafLatency5000To10000ms = &v
}

// GetWafLatency10000To60000ms returns the WafLatency10000To60000ms field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetWafLatency10000To60000ms() int32 {
	if o == nil || o.WafLatency10000To60000ms == nil {
		var ret int32
		return ret
	}
	return *o.WafLatency10000To60000ms
}

// GetWafLatency10000To60000msOk returns a tuple with the WafLatency10000To60000ms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetWafLatency10000To60000msOk() (*int32, bool) {
	if o == nil || o.WafLatency10000To60000ms == nil {
		return nil, false
	}
	return o.WafLatency10000To60000ms, true
}

// HasWafLatency10000To60000ms returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasWafLatency10000To60000ms() bool {
	if o != nil && o.WafLatency10000To60000ms != nil {
		return true
	}

	return false
}

// SetWafLatency10000To60000ms gets a reference to the given int32 and assigns it to the WafLatency10000To60000ms field.
func (o *OriginInspectorMeasurements) SetWafLatency10000To60000ms(v int32) {
	o.WafLatency10000To60000ms = &v
}

// GetWafLatency60000ms returns the WafLatency60000ms field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetWafLatency60000ms() int32 {
	if o == nil || o.WafLatency60000ms == nil {
		var ret int32
		return ret
	}
	return *o.WafLatency60000ms
}

// GetWafLatency60000msOk returns a tuple with the WafLatency60000ms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetWafLatency60000msOk() (*int32, bool) {
	if o == nil || o.WafLatency60000ms == nil {
		return nil, false
	}
	return o.WafLatency60000ms, true
}

// HasWafLatency60000ms returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasWafLatency60000ms() bool {
	if o != nil && o.WafLatency60000ms != nil {
		return true
	}

	return false
}

// SetWafLatency60000ms gets a reference to the given int32 and assigns it to the WafLatency60000ms field.
func (o *OriginInspectorMeasurements) SetWafLatency60000ms(v int32) {
	o.WafLatency60000ms = &v
}

// GetComputeResponses returns the ComputeResponses field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetComputeResponses() int32 {
	if o == nil || o.ComputeResponses == nil {
		var ret int32
		return ret
	}
	return *o.ComputeResponses
}

// GetComputeResponsesOk returns a tuple with the ComputeResponses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetComputeResponsesOk() (*int32, bool) {
	if o == nil || o.ComputeResponses == nil {
		return nil, false
	}
	return o.ComputeResponses, true
}

// HasComputeResponses returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasComputeResponses() bool {
	if o != nil && o.ComputeResponses != nil {
		return true
	}

	return false
}

// SetComputeResponses gets a reference to the given int32 and assigns it to the ComputeResponses field.
func (o *OriginInspectorMeasurements) SetComputeResponses(v int32) {
	o.ComputeResponses = &v
}

// GetComputeRespHeaderBytes returns the ComputeRespHeaderBytes field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetComputeRespHeaderBytes() int32 {
	if o == nil || o.ComputeRespHeaderBytes == nil {
		var ret int32
		return ret
	}
	return *o.ComputeRespHeaderBytes
}

// GetComputeRespHeaderBytesOk returns a tuple with the ComputeRespHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetComputeRespHeaderBytesOk() (*int32, bool) {
	if o == nil || o.ComputeRespHeaderBytes == nil {
		return nil, false
	}
	return o.ComputeRespHeaderBytes, true
}

// HasComputeRespHeaderBytes returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasComputeRespHeaderBytes() bool {
	if o != nil && o.ComputeRespHeaderBytes != nil {
		return true
	}

	return false
}

// SetComputeRespHeaderBytes gets a reference to the given int32 and assigns it to the ComputeRespHeaderBytes field.
func (o *OriginInspectorMeasurements) SetComputeRespHeaderBytes(v int32) {
	o.ComputeRespHeaderBytes = &v
}

// GetComputeRespBodyBytes returns the ComputeRespBodyBytes field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetComputeRespBodyBytes() int32 {
	if o == nil || o.ComputeRespBodyBytes == nil {
		var ret int32
		return ret
	}
	return *o.ComputeRespBodyBytes
}

// GetComputeRespBodyBytesOk returns a tuple with the ComputeRespBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetComputeRespBodyBytesOk() (*int32, bool) {
	if o == nil || o.ComputeRespBodyBytes == nil {
		return nil, false
	}
	return o.ComputeRespBodyBytes, true
}

// HasComputeRespBodyBytes returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasComputeRespBodyBytes() bool {
	if o != nil && o.ComputeRespBodyBytes != nil {
		return true
	}

	return false
}

// SetComputeRespBodyBytes gets a reference to the given int32 and assigns it to the ComputeRespBodyBytes field.
func (o *OriginInspectorMeasurements) SetComputeRespBodyBytes(v int32) {
	o.ComputeRespBodyBytes = &v
}

// GetComputeStatus1xx returns the ComputeStatus1xx field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetComputeStatus1xx() int32 {
	if o == nil || o.ComputeStatus1xx == nil {
		var ret int32
		return ret
	}
	return *o.ComputeStatus1xx
}

// GetComputeStatus1xxOk returns a tuple with the ComputeStatus1xx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetComputeStatus1xxOk() (*int32, bool) {
	if o == nil || o.ComputeStatus1xx == nil {
		return nil, false
	}
	return o.ComputeStatus1xx, true
}

// HasComputeStatus1xx returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasComputeStatus1xx() bool {
	if o != nil && o.ComputeStatus1xx != nil {
		return true
	}

	return false
}

// SetComputeStatus1xx gets a reference to the given int32 and assigns it to the ComputeStatus1xx field.
func (o *OriginInspectorMeasurements) SetComputeStatus1xx(v int32) {
	o.ComputeStatus1xx = &v
}

// GetComputeStatus2xx returns the ComputeStatus2xx field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetComputeStatus2xx() int32 {
	if o == nil || o.ComputeStatus2xx == nil {
		var ret int32
		return ret
	}
	return *o.ComputeStatus2xx
}

// GetComputeStatus2xxOk returns a tuple with the ComputeStatus2xx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetComputeStatus2xxOk() (*int32, bool) {
	if o == nil || o.ComputeStatus2xx == nil {
		return nil, false
	}
	return o.ComputeStatus2xx, true
}

// HasComputeStatus2xx returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasComputeStatus2xx() bool {
	if o != nil && o.ComputeStatus2xx != nil {
		return true
	}

	return false
}

// SetComputeStatus2xx gets a reference to the given int32 and assigns it to the ComputeStatus2xx field.
func (o *OriginInspectorMeasurements) SetComputeStatus2xx(v int32) {
	o.ComputeStatus2xx = &v
}

// GetComputeStatus3xx returns the ComputeStatus3xx field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetComputeStatus3xx() int32 {
	if o == nil || o.ComputeStatus3xx == nil {
		var ret int32
		return ret
	}
	return *o.ComputeStatus3xx
}

// GetComputeStatus3xxOk returns a tuple with the ComputeStatus3xx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetComputeStatus3xxOk() (*int32, bool) {
	if o == nil || o.ComputeStatus3xx == nil {
		return nil, false
	}
	return o.ComputeStatus3xx, true
}

// HasComputeStatus3xx returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasComputeStatus3xx() bool {
	if o != nil && o.ComputeStatus3xx != nil {
		return true
	}

	return false
}

// SetComputeStatus3xx gets a reference to the given int32 and assigns it to the ComputeStatus3xx field.
func (o *OriginInspectorMeasurements) SetComputeStatus3xx(v int32) {
	o.ComputeStatus3xx = &v
}

// GetComputeStatus4xx returns the ComputeStatus4xx field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetComputeStatus4xx() int32 {
	if o == nil || o.ComputeStatus4xx == nil {
		var ret int32
		return ret
	}
	return *o.ComputeStatus4xx
}

// GetComputeStatus4xxOk returns a tuple with the ComputeStatus4xx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetComputeStatus4xxOk() (*int32, bool) {
	if o == nil || o.ComputeStatus4xx == nil {
		return nil, false
	}
	return o.ComputeStatus4xx, true
}

// HasComputeStatus4xx returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasComputeStatus4xx() bool {
	if o != nil && o.ComputeStatus4xx != nil {
		return true
	}

	return false
}

// SetComputeStatus4xx gets a reference to the given int32 and assigns it to the ComputeStatus4xx field.
func (o *OriginInspectorMeasurements) SetComputeStatus4xx(v int32) {
	o.ComputeStatus4xx = &v
}

// GetComputeStatus5xx returns the ComputeStatus5xx field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetComputeStatus5xx() int32 {
	if o == nil || o.ComputeStatus5xx == nil {
		var ret int32
		return ret
	}
	return *o.ComputeStatus5xx
}

// GetComputeStatus5xxOk returns a tuple with the ComputeStatus5xx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetComputeStatus5xxOk() (*int32, bool) {
	if o == nil || o.ComputeStatus5xx == nil {
		return nil, false
	}
	return o.ComputeStatus5xx, true
}

// HasComputeStatus5xx returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasComputeStatus5xx() bool {
	if o != nil && o.ComputeStatus5xx != nil {
		return true
	}

	return false
}

// SetComputeStatus5xx gets a reference to the given int32 and assigns it to the ComputeStatus5xx field.
func (o *OriginInspectorMeasurements) SetComputeStatus5xx(v int32) {
	o.ComputeStatus5xx = &v
}

// GetComputeStatus200 returns the ComputeStatus200 field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetComputeStatus200() int32 {
	if o == nil || o.ComputeStatus200 == nil {
		var ret int32
		return ret
	}
	return *o.ComputeStatus200
}

// GetComputeStatus200Ok returns a tuple with the ComputeStatus200 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetComputeStatus200Ok() (*int32, bool) {
	if o == nil || o.ComputeStatus200 == nil {
		return nil, false
	}
	return o.ComputeStatus200, true
}

// HasComputeStatus200 returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasComputeStatus200() bool {
	if o != nil && o.ComputeStatus200 != nil {
		return true
	}

	return false
}

// SetComputeStatus200 gets a reference to the given int32 and assigns it to the ComputeStatus200 field.
func (o *OriginInspectorMeasurements) SetComputeStatus200(v int32) {
	o.ComputeStatus200 = &v
}

// GetComputeStatus204 returns the ComputeStatus204 field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetComputeStatus204() int32 {
	if o == nil || o.ComputeStatus204 == nil {
		var ret int32
		return ret
	}
	return *o.ComputeStatus204
}

// GetComputeStatus204Ok returns a tuple with the ComputeStatus204 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetComputeStatus204Ok() (*int32, bool) {
	if o == nil || o.ComputeStatus204 == nil {
		return nil, false
	}
	return o.ComputeStatus204, true
}

// HasComputeStatus204 returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasComputeStatus204() bool {
	if o != nil && o.ComputeStatus204 != nil {
		return true
	}

	return false
}

// SetComputeStatus204 gets a reference to the given int32 and assigns it to the ComputeStatus204 field.
func (o *OriginInspectorMeasurements) SetComputeStatus204(v int32) {
	o.ComputeStatus204 = &v
}

// GetComputeStatus206 returns the ComputeStatus206 field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetComputeStatus206() int32 {
	if o == nil || o.ComputeStatus206 == nil {
		var ret int32
		return ret
	}
	return *o.ComputeStatus206
}

// GetComputeStatus206Ok returns a tuple with the ComputeStatus206 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetComputeStatus206Ok() (*int32, bool) {
	if o == nil || o.ComputeStatus206 == nil {
		return nil, false
	}
	return o.ComputeStatus206, true
}

// HasComputeStatus206 returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasComputeStatus206() bool {
	if o != nil && o.ComputeStatus206 != nil {
		return true
	}

	return false
}

// SetComputeStatus206 gets a reference to the given int32 and assigns it to the ComputeStatus206 field.
func (o *OriginInspectorMeasurements) SetComputeStatus206(v int32) {
	o.ComputeStatus206 = &v
}

// GetComputeStatus301 returns the ComputeStatus301 field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetComputeStatus301() int32 {
	if o == nil || o.ComputeStatus301 == nil {
		var ret int32
		return ret
	}
	return *o.ComputeStatus301
}

// GetComputeStatus301Ok returns a tuple with the ComputeStatus301 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetComputeStatus301Ok() (*int32, bool) {
	if o == nil || o.ComputeStatus301 == nil {
		return nil, false
	}
	return o.ComputeStatus301, true
}

// HasComputeStatus301 returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasComputeStatus301() bool {
	if o != nil && o.ComputeStatus301 != nil {
		return true
	}

	return false
}

// SetComputeStatus301 gets a reference to the given int32 and assigns it to the ComputeStatus301 field.
func (o *OriginInspectorMeasurements) SetComputeStatus301(v int32) {
	o.ComputeStatus301 = &v
}

// GetComputeStatus302 returns the ComputeStatus302 field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetComputeStatus302() int32 {
	if o == nil || o.ComputeStatus302 == nil {
		var ret int32
		return ret
	}
	return *o.ComputeStatus302
}

// GetComputeStatus302Ok returns a tuple with the ComputeStatus302 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetComputeStatus302Ok() (*int32, bool) {
	if o == nil || o.ComputeStatus302 == nil {
		return nil, false
	}
	return o.ComputeStatus302, true
}

// HasComputeStatus302 returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasComputeStatus302() bool {
	if o != nil && o.ComputeStatus302 != nil {
		return true
	}

	return false
}

// SetComputeStatus302 gets a reference to the given int32 and assigns it to the ComputeStatus302 field.
func (o *OriginInspectorMeasurements) SetComputeStatus302(v int32) {
	o.ComputeStatus302 = &v
}

// GetComputeStatus304 returns the ComputeStatus304 field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetComputeStatus304() int32 {
	if o == nil || o.ComputeStatus304 == nil {
		var ret int32
		return ret
	}
	return *o.ComputeStatus304
}

// GetComputeStatus304Ok returns a tuple with the ComputeStatus304 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetComputeStatus304Ok() (*int32, bool) {
	if o == nil || o.ComputeStatus304 == nil {
		return nil, false
	}
	return o.ComputeStatus304, true
}

// HasComputeStatus304 returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasComputeStatus304() bool {
	if o != nil && o.ComputeStatus304 != nil {
		return true
	}

	return false
}

// SetComputeStatus304 gets a reference to the given int32 and assigns it to the ComputeStatus304 field.
func (o *OriginInspectorMeasurements) SetComputeStatus304(v int32) {
	o.ComputeStatus304 = &v
}

// GetComputeStatus400 returns the ComputeStatus400 field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetComputeStatus400() int32 {
	if o == nil || o.ComputeStatus400 == nil {
		var ret int32
		return ret
	}
	return *o.ComputeStatus400
}

// GetComputeStatus400Ok returns a tuple with the ComputeStatus400 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetComputeStatus400Ok() (*int32, bool) {
	if o == nil || o.ComputeStatus400 == nil {
		return nil, false
	}
	return o.ComputeStatus400, true
}

// HasComputeStatus400 returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasComputeStatus400() bool {
	if o != nil && o.ComputeStatus400 != nil {
		return true
	}

	return false
}

// SetComputeStatus400 gets a reference to the given int32 and assigns it to the ComputeStatus400 field.
func (o *OriginInspectorMeasurements) SetComputeStatus400(v int32) {
	o.ComputeStatus400 = &v
}

// GetComputeStatus401 returns the ComputeStatus401 field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetComputeStatus401() int32 {
	if o == nil || o.ComputeStatus401 == nil {
		var ret int32
		return ret
	}
	return *o.ComputeStatus401
}

// GetComputeStatus401Ok returns a tuple with the ComputeStatus401 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetComputeStatus401Ok() (*int32, bool) {
	if o == nil || o.ComputeStatus401 == nil {
		return nil, false
	}
	return o.ComputeStatus401, true
}

// HasComputeStatus401 returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasComputeStatus401() bool {
	if o != nil && o.ComputeStatus401 != nil {
		return true
	}

	return false
}

// SetComputeStatus401 gets a reference to the given int32 and assigns it to the ComputeStatus401 field.
func (o *OriginInspectorMeasurements) SetComputeStatus401(v int32) {
	o.ComputeStatus401 = &v
}

// GetComputeStatus403 returns the ComputeStatus403 field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetComputeStatus403() int32 {
	if o == nil || o.ComputeStatus403 == nil {
		var ret int32
		return ret
	}
	return *o.ComputeStatus403
}

// GetComputeStatus403Ok returns a tuple with the ComputeStatus403 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetComputeStatus403Ok() (*int32, bool) {
	if o == nil || o.ComputeStatus403 == nil {
		return nil, false
	}
	return o.ComputeStatus403, true
}

// HasComputeStatus403 returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasComputeStatus403() bool {
	if o != nil && o.ComputeStatus403 != nil {
		return true
	}

	return false
}

// SetComputeStatus403 gets a reference to the given int32 and assigns it to the ComputeStatus403 field.
func (o *OriginInspectorMeasurements) SetComputeStatus403(v int32) {
	o.ComputeStatus403 = &v
}

// GetComputeStatus404 returns the ComputeStatus404 field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetComputeStatus404() int32 {
	if o == nil || o.ComputeStatus404 == nil {
		var ret int32
		return ret
	}
	return *o.ComputeStatus404
}

// GetComputeStatus404Ok returns a tuple with the ComputeStatus404 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetComputeStatus404Ok() (*int32, bool) {
	if o == nil || o.ComputeStatus404 == nil {
		return nil, false
	}
	return o.ComputeStatus404, true
}

// HasComputeStatus404 returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasComputeStatus404() bool {
	if o != nil && o.ComputeStatus404 != nil {
		return true
	}

	return false
}

// SetComputeStatus404 gets a reference to the given int32 and assigns it to the ComputeStatus404 field.
func (o *OriginInspectorMeasurements) SetComputeStatus404(v int32) {
	o.ComputeStatus404 = &v
}

// GetComputeStatus416 returns the ComputeStatus416 field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetComputeStatus416() int32 {
	if o == nil || o.ComputeStatus416 == nil {
		var ret int32
		return ret
	}
	return *o.ComputeStatus416
}

// GetComputeStatus416Ok returns a tuple with the ComputeStatus416 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetComputeStatus416Ok() (*int32, bool) {
	if o == nil || o.ComputeStatus416 == nil {
		return nil, false
	}
	return o.ComputeStatus416, true
}

// HasComputeStatus416 returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasComputeStatus416() bool {
	if o != nil && o.ComputeStatus416 != nil {
		return true
	}

	return false
}

// SetComputeStatus416 gets a reference to the given int32 and assigns it to the ComputeStatus416 field.
func (o *OriginInspectorMeasurements) SetComputeStatus416(v int32) {
	o.ComputeStatus416 = &v
}

// GetComputeStatus429 returns the ComputeStatus429 field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetComputeStatus429() int32 {
	if o == nil || o.ComputeStatus429 == nil {
		var ret int32
		return ret
	}
	return *o.ComputeStatus429
}

// GetComputeStatus429Ok returns a tuple with the ComputeStatus429 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetComputeStatus429Ok() (*int32, bool) {
	if o == nil || o.ComputeStatus429 == nil {
		return nil, false
	}
	return o.ComputeStatus429, true
}

// HasComputeStatus429 returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasComputeStatus429() bool {
	if o != nil && o.ComputeStatus429 != nil {
		return true
	}

	return false
}

// SetComputeStatus429 gets a reference to the given int32 and assigns it to the ComputeStatus429 field.
func (o *OriginInspectorMeasurements) SetComputeStatus429(v int32) {
	o.ComputeStatus429 = &v
}

// GetComputeStatus500 returns the ComputeStatus500 field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetComputeStatus500() int32 {
	if o == nil || o.ComputeStatus500 == nil {
		var ret int32
		return ret
	}
	return *o.ComputeStatus500
}

// GetComputeStatus500Ok returns a tuple with the ComputeStatus500 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetComputeStatus500Ok() (*int32, bool) {
	if o == nil || o.ComputeStatus500 == nil {
		return nil, false
	}
	return o.ComputeStatus500, true
}

// HasComputeStatus500 returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasComputeStatus500() bool {
	if o != nil && o.ComputeStatus500 != nil {
		return true
	}

	return false
}

// SetComputeStatus500 gets a reference to the given int32 and assigns it to the ComputeStatus500 field.
func (o *OriginInspectorMeasurements) SetComputeStatus500(v int32) {
	o.ComputeStatus500 = &v
}

// GetComputeStatus501 returns the ComputeStatus501 field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetComputeStatus501() int32 {
	if o == nil || o.ComputeStatus501 == nil {
		var ret int32
		return ret
	}
	return *o.ComputeStatus501
}

// GetComputeStatus501Ok returns a tuple with the ComputeStatus501 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetComputeStatus501Ok() (*int32, bool) {
	if o == nil || o.ComputeStatus501 == nil {
		return nil, false
	}
	return o.ComputeStatus501, true
}

// HasComputeStatus501 returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasComputeStatus501() bool {
	if o != nil && o.ComputeStatus501 != nil {
		return true
	}

	return false
}

// SetComputeStatus501 gets a reference to the given int32 and assigns it to the ComputeStatus501 field.
func (o *OriginInspectorMeasurements) SetComputeStatus501(v int32) {
	o.ComputeStatus501 = &v
}

// GetComputeStatus502 returns the ComputeStatus502 field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetComputeStatus502() int32 {
	if o == nil || o.ComputeStatus502 == nil {
		var ret int32
		return ret
	}
	return *o.ComputeStatus502
}

// GetComputeStatus502Ok returns a tuple with the ComputeStatus502 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetComputeStatus502Ok() (*int32, bool) {
	if o == nil || o.ComputeStatus502 == nil {
		return nil, false
	}
	return o.ComputeStatus502, true
}

// HasComputeStatus502 returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasComputeStatus502() bool {
	if o != nil && o.ComputeStatus502 != nil {
		return true
	}

	return false
}

// SetComputeStatus502 gets a reference to the given int32 and assigns it to the ComputeStatus502 field.
func (o *OriginInspectorMeasurements) SetComputeStatus502(v int32) {
	o.ComputeStatus502 = &v
}

// GetComputeStatus503 returns the ComputeStatus503 field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetComputeStatus503() int32 {
	if o == nil || o.ComputeStatus503 == nil {
		var ret int32
		return ret
	}
	return *o.ComputeStatus503
}

// GetComputeStatus503Ok returns a tuple with the ComputeStatus503 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetComputeStatus503Ok() (*int32, bool) {
	if o == nil || o.ComputeStatus503 == nil {
		return nil, false
	}
	return o.ComputeStatus503, true
}

// HasComputeStatus503 returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasComputeStatus503() bool {
	if o != nil && o.ComputeStatus503 != nil {
		return true
	}

	return false
}

// SetComputeStatus503 gets a reference to the given int32 and assigns it to the ComputeStatus503 field.
func (o *OriginInspectorMeasurements) SetComputeStatus503(v int32) {
	o.ComputeStatus503 = &v
}

// GetComputeStatus504 returns the ComputeStatus504 field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetComputeStatus504() int32 {
	if o == nil || o.ComputeStatus504 == nil {
		var ret int32
		return ret
	}
	return *o.ComputeStatus504
}

// GetComputeStatus504Ok returns a tuple with the ComputeStatus504 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetComputeStatus504Ok() (*int32, bool) {
	if o == nil || o.ComputeStatus504 == nil {
		return nil, false
	}
	return o.ComputeStatus504, true
}

// HasComputeStatus504 returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasComputeStatus504() bool {
	if o != nil && o.ComputeStatus504 != nil {
		return true
	}

	return false
}

// SetComputeStatus504 gets a reference to the given int32 and assigns it to the ComputeStatus504 field.
func (o *OriginInspectorMeasurements) SetComputeStatus504(v int32) {
	o.ComputeStatus504 = &v
}

// GetComputeStatus505 returns the ComputeStatus505 field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetComputeStatus505() int32 {
	if o == nil || o.ComputeStatus505 == nil {
		var ret int32
		return ret
	}
	return *o.ComputeStatus505
}

// GetComputeStatus505Ok returns a tuple with the ComputeStatus505 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetComputeStatus505Ok() (*int32, bool) {
	if o == nil || o.ComputeStatus505 == nil {
		return nil, false
	}
	return o.ComputeStatus505, true
}

// HasComputeStatus505 returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasComputeStatus505() bool {
	if o != nil && o.ComputeStatus505 != nil {
		return true
	}

	return false
}

// SetComputeStatus505 gets a reference to the given int32 and assigns it to the ComputeStatus505 field.
func (o *OriginInspectorMeasurements) SetComputeStatus505(v int32) {
	o.ComputeStatus505 = &v
}

// GetComputeLatency0To1ms returns the ComputeLatency0To1ms field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetComputeLatency0To1ms() int32 {
	if o == nil || o.ComputeLatency0To1ms == nil {
		var ret int32
		return ret
	}
	return *o.ComputeLatency0To1ms
}

// GetComputeLatency0To1msOk returns a tuple with the ComputeLatency0To1ms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetComputeLatency0To1msOk() (*int32, bool) {
	if o == nil || o.ComputeLatency0To1ms == nil {
		return nil, false
	}
	return o.ComputeLatency0To1ms, true
}

// HasComputeLatency0To1ms returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasComputeLatency0To1ms() bool {
	if o != nil && o.ComputeLatency0To1ms != nil {
		return true
	}

	return false
}

// SetComputeLatency0To1ms gets a reference to the given int32 and assigns it to the ComputeLatency0To1ms field.
func (o *OriginInspectorMeasurements) SetComputeLatency0To1ms(v int32) {
	o.ComputeLatency0To1ms = &v
}

// GetComputeLatency1To5ms returns the ComputeLatency1To5ms field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetComputeLatency1To5ms() int32 {
	if o == nil || o.ComputeLatency1To5ms == nil {
		var ret int32
		return ret
	}
	return *o.ComputeLatency1To5ms
}

// GetComputeLatency1To5msOk returns a tuple with the ComputeLatency1To5ms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetComputeLatency1To5msOk() (*int32, bool) {
	if o == nil || o.ComputeLatency1To5ms == nil {
		return nil, false
	}
	return o.ComputeLatency1To5ms, true
}

// HasComputeLatency1To5ms returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasComputeLatency1To5ms() bool {
	if o != nil && o.ComputeLatency1To5ms != nil {
		return true
	}

	return false
}

// SetComputeLatency1To5ms gets a reference to the given int32 and assigns it to the ComputeLatency1To5ms field.
func (o *OriginInspectorMeasurements) SetComputeLatency1To5ms(v int32) {
	o.ComputeLatency1To5ms = &v
}

// GetComputeLatency5To10ms returns the ComputeLatency5To10ms field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetComputeLatency5To10ms() int32 {
	if o == nil || o.ComputeLatency5To10ms == nil {
		var ret int32
		return ret
	}
	return *o.ComputeLatency5To10ms
}

// GetComputeLatency5To10msOk returns a tuple with the ComputeLatency5To10ms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetComputeLatency5To10msOk() (*int32, bool) {
	if o == nil || o.ComputeLatency5To10ms == nil {
		return nil, false
	}
	return o.ComputeLatency5To10ms, true
}

// HasComputeLatency5To10ms returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasComputeLatency5To10ms() bool {
	if o != nil && o.ComputeLatency5To10ms != nil {
		return true
	}

	return false
}

// SetComputeLatency5To10ms gets a reference to the given int32 and assigns it to the ComputeLatency5To10ms field.
func (o *OriginInspectorMeasurements) SetComputeLatency5To10ms(v int32) {
	o.ComputeLatency5To10ms = &v
}

// GetComputeLatency10To50ms returns the ComputeLatency10To50ms field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetComputeLatency10To50ms() int32 {
	if o == nil || o.ComputeLatency10To50ms == nil {
		var ret int32
		return ret
	}
	return *o.ComputeLatency10To50ms
}

// GetComputeLatency10To50msOk returns a tuple with the ComputeLatency10To50ms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetComputeLatency10To50msOk() (*int32, bool) {
	if o == nil || o.ComputeLatency10To50ms == nil {
		return nil, false
	}
	return o.ComputeLatency10To50ms, true
}

// HasComputeLatency10To50ms returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasComputeLatency10To50ms() bool {
	if o != nil && o.ComputeLatency10To50ms != nil {
		return true
	}

	return false
}

// SetComputeLatency10To50ms gets a reference to the given int32 and assigns it to the ComputeLatency10To50ms field.
func (o *OriginInspectorMeasurements) SetComputeLatency10To50ms(v int32) {
	o.ComputeLatency10To50ms = &v
}

// GetComputeLatency50To100ms returns the ComputeLatency50To100ms field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetComputeLatency50To100ms() int32 {
	if o == nil || o.ComputeLatency50To100ms == nil {
		var ret int32
		return ret
	}
	return *o.ComputeLatency50To100ms
}

// GetComputeLatency50To100msOk returns a tuple with the ComputeLatency50To100ms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetComputeLatency50To100msOk() (*int32, bool) {
	if o == nil || o.ComputeLatency50To100ms == nil {
		return nil, false
	}
	return o.ComputeLatency50To100ms, true
}

// HasComputeLatency50To100ms returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasComputeLatency50To100ms() bool {
	if o != nil && o.ComputeLatency50To100ms != nil {
		return true
	}

	return false
}

// SetComputeLatency50To100ms gets a reference to the given int32 and assigns it to the ComputeLatency50To100ms field.
func (o *OriginInspectorMeasurements) SetComputeLatency50To100ms(v int32) {
	o.ComputeLatency50To100ms = &v
}

// GetComputeLatency100To250ms returns the ComputeLatency100To250ms field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetComputeLatency100To250ms() int32 {
	if o == nil || o.ComputeLatency100To250ms == nil {
		var ret int32
		return ret
	}
	return *o.ComputeLatency100To250ms
}

// GetComputeLatency100To250msOk returns a tuple with the ComputeLatency100To250ms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetComputeLatency100To250msOk() (*int32, bool) {
	if o == nil || o.ComputeLatency100To250ms == nil {
		return nil, false
	}
	return o.ComputeLatency100To250ms, true
}

// HasComputeLatency100To250ms returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasComputeLatency100To250ms() bool {
	if o != nil && o.ComputeLatency100To250ms != nil {
		return true
	}

	return false
}

// SetComputeLatency100To250ms gets a reference to the given int32 and assigns it to the ComputeLatency100To250ms field.
func (o *OriginInspectorMeasurements) SetComputeLatency100To250ms(v int32) {
	o.ComputeLatency100To250ms = &v
}

// GetComputeLatency250To500ms returns the ComputeLatency250To500ms field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetComputeLatency250To500ms() int32 {
	if o == nil || o.ComputeLatency250To500ms == nil {
		var ret int32
		return ret
	}
	return *o.ComputeLatency250To500ms
}

// GetComputeLatency250To500msOk returns a tuple with the ComputeLatency250To500ms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetComputeLatency250To500msOk() (*int32, bool) {
	if o == nil || o.ComputeLatency250To500ms == nil {
		return nil, false
	}
	return o.ComputeLatency250To500ms, true
}

// HasComputeLatency250To500ms returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasComputeLatency250To500ms() bool {
	if o != nil && o.ComputeLatency250To500ms != nil {
		return true
	}

	return false
}

// SetComputeLatency250To500ms gets a reference to the given int32 and assigns it to the ComputeLatency250To500ms field.
func (o *OriginInspectorMeasurements) SetComputeLatency250To500ms(v int32) {
	o.ComputeLatency250To500ms = &v
}

// GetComputeLatency500To1000ms returns the ComputeLatency500To1000ms field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetComputeLatency500To1000ms() int32 {
	if o == nil || o.ComputeLatency500To1000ms == nil {
		var ret int32
		return ret
	}
	return *o.ComputeLatency500To1000ms
}

// GetComputeLatency500To1000msOk returns a tuple with the ComputeLatency500To1000ms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetComputeLatency500To1000msOk() (*int32, bool) {
	if o == nil || o.ComputeLatency500To1000ms == nil {
		return nil, false
	}
	return o.ComputeLatency500To1000ms, true
}

// HasComputeLatency500To1000ms returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasComputeLatency500To1000ms() bool {
	if o != nil && o.ComputeLatency500To1000ms != nil {
		return true
	}

	return false
}

// SetComputeLatency500To1000ms gets a reference to the given int32 and assigns it to the ComputeLatency500To1000ms field.
func (o *OriginInspectorMeasurements) SetComputeLatency500To1000ms(v int32) {
	o.ComputeLatency500To1000ms = &v
}

// GetComputeLatency1000To5000ms returns the ComputeLatency1000To5000ms field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetComputeLatency1000To5000ms() int32 {
	if o == nil || o.ComputeLatency1000To5000ms == nil {
		var ret int32
		return ret
	}
	return *o.ComputeLatency1000To5000ms
}

// GetComputeLatency1000To5000msOk returns a tuple with the ComputeLatency1000To5000ms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetComputeLatency1000To5000msOk() (*int32, bool) {
	if o == nil || o.ComputeLatency1000To5000ms == nil {
		return nil, false
	}
	return o.ComputeLatency1000To5000ms, true
}

// HasComputeLatency1000To5000ms returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasComputeLatency1000To5000ms() bool {
	if o != nil && o.ComputeLatency1000To5000ms != nil {
		return true
	}

	return false
}

// SetComputeLatency1000To5000ms gets a reference to the given int32 and assigns it to the ComputeLatency1000To5000ms field.
func (o *OriginInspectorMeasurements) SetComputeLatency1000To5000ms(v int32) {
	o.ComputeLatency1000To5000ms = &v
}

// GetComputeLatency5000To10000ms returns the ComputeLatency5000To10000ms field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetComputeLatency5000To10000ms() int32 {
	if o == nil || o.ComputeLatency5000To10000ms == nil {
		var ret int32
		return ret
	}
	return *o.ComputeLatency5000To10000ms
}

// GetComputeLatency5000To10000msOk returns a tuple with the ComputeLatency5000To10000ms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetComputeLatency5000To10000msOk() (*int32, bool) {
	if o == nil || o.ComputeLatency5000To10000ms == nil {
		return nil, false
	}
	return o.ComputeLatency5000To10000ms, true
}

// HasComputeLatency5000To10000ms returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasComputeLatency5000To10000ms() bool {
	if o != nil && o.ComputeLatency5000To10000ms != nil {
		return true
	}

	return false
}

// SetComputeLatency5000To10000ms gets a reference to the given int32 and assigns it to the ComputeLatency5000To10000ms field.
func (o *OriginInspectorMeasurements) SetComputeLatency5000To10000ms(v int32) {
	o.ComputeLatency5000To10000ms = &v
}

// GetComputeLatency10000To60000ms returns the ComputeLatency10000To60000ms field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetComputeLatency10000To60000ms() int32 {
	if o == nil || o.ComputeLatency10000To60000ms == nil {
		var ret int32
		return ret
	}
	return *o.ComputeLatency10000To60000ms
}

// GetComputeLatency10000To60000msOk returns a tuple with the ComputeLatency10000To60000ms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetComputeLatency10000To60000msOk() (*int32, bool) {
	if o == nil || o.ComputeLatency10000To60000ms == nil {
		return nil, false
	}
	return o.ComputeLatency10000To60000ms, true
}

// HasComputeLatency10000To60000ms returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasComputeLatency10000To60000ms() bool {
	if o != nil && o.ComputeLatency10000To60000ms != nil {
		return true
	}

	return false
}

// SetComputeLatency10000To60000ms gets a reference to the given int32 and assigns it to the ComputeLatency10000To60000ms field.
func (o *OriginInspectorMeasurements) SetComputeLatency10000To60000ms(v int32) {
	o.ComputeLatency10000To60000ms = &v
}

// GetComputeLatency60000ms returns the ComputeLatency60000ms field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetComputeLatency60000ms() int32 {
	if o == nil || o.ComputeLatency60000ms == nil {
		var ret int32
		return ret
	}
	return *o.ComputeLatency60000ms
}

// GetComputeLatency60000msOk returns a tuple with the ComputeLatency60000ms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetComputeLatency60000msOk() (*int32, bool) {
	if o == nil || o.ComputeLatency60000ms == nil {
		return nil, false
	}
	return o.ComputeLatency60000ms, true
}

// HasComputeLatency60000ms returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasComputeLatency60000ms() bool {
	if o != nil && o.ComputeLatency60000ms != nil {
		return true
	}

	return false
}

// SetComputeLatency60000ms gets a reference to the given int32 and assigns it to the ComputeLatency60000ms field.
func (o *OriginInspectorMeasurements) SetComputeLatency60000ms(v int32) {
	o.ComputeLatency60000ms = &v
}

// GetAllResponses returns the AllResponses field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetAllResponses() int32 {
	if o == nil || o.AllResponses == nil {
		var ret int32
		return ret
	}
	return *o.AllResponses
}

// GetAllResponsesOk returns a tuple with the AllResponses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetAllResponsesOk() (*int32, bool) {
	if o == nil || o.AllResponses == nil {
		return nil, false
	}
	return o.AllResponses, true
}

// HasAllResponses returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasAllResponses() bool {
	if o != nil && o.AllResponses != nil {
		return true
	}

	return false
}

// SetAllResponses gets a reference to the given int32 and assigns it to the AllResponses field.
func (o *OriginInspectorMeasurements) SetAllResponses(v int32) {
	o.AllResponses = &v
}

// GetAllRespHeaderBytes returns the AllRespHeaderBytes field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetAllRespHeaderBytes() int32 {
	if o == nil || o.AllRespHeaderBytes == nil {
		var ret int32
		return ret
	}
	return *o.AllRespHeaderBytes
}

// GetAllRespHeaderBytesOk returns a tuple with the AllRespHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetAllRespHeaderBytesOk() (*int32, bool) {
	if o == nil || o.AllRespHeaderBytes == nil {
		return nil, false
	}
	return o.AllRespHeaderBytes, true
}

// HasAllRespHeaderBytes returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasAllRespHeaderBytes() bool {
	if o != nil && o.AllRespHeaderBytes != nil {
		return true
	}

	return false
}

// SetAllRespHeaderBytes gets a reference to the given int32 and assigns it to the AllRespHeaderBytes field.
func (o *OriginInspectorMeasurements) SetAllRespHeaderBytes(v int32) {
	o.AllRespHeaderBytes = &v
}

// GetAllRespBodyBytes returns the AllRespBodyBytes field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetAllRespBodyBytes() int32 {
	if o == nil || o.AllRespBodyBytes == nil {
		var ret int32
		return ret
	}
	return *o.AllRespBodyBytes
}

// GetAllRespBodyBytesOk returns a tuple with the AllRespBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetAllRespBodyBytesOk() (*int32, bool) {
	if o == nil || o.AllRespBodyBytes == nil {
		return nil, false
	}
	return o.AllRespBodyBytes, true
}

// HasAllRespBodyBytes returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasAllRespBodyBytes() bool {
	if o != nil && o.AllRespBodyBytes != nil {
		return true
	}

	return false
}

// SetAllRespBodyBytes gets a reference to the given int32 and assigns it to the AllRespBodyBytes field.
func (o *OriginInspectorMeasurements) SetAllRespBodyBytes(v int32) {
	o.AllRespBodyBytes = &v
}

// GetAllStatus1xx returns the AllStatus1xx field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetAllStatus1xx() int32 {
	if o == nil || o.AllStatus1xx == nil {
		var ret int32
		return ret
	}
	return *o.AllStatus1xx
}

// GetAllStatus1xxOk returns a tuple with the AllStatus1xx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetAllStatus1xxOk() (*int32, bool) {
	if o == nil || o.AllStatus1xx == nil {
		return nil, false
	}
	return o.AllStatus1xx, true
}

// HasAllStatus1xx returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasAllStatus1xx() bool {
	if o != nil && o.AllStatus1xx != nil {
		return true
	}

	return false
}

// SetAllStatus1xx gets a reference to the given int32 and assigns it to the AllStatus1xx field.
func (o *OriginInspectorMeasurements) SetAllStatus1xx(v int32) {
	o.AllStatus1xx = &v
}

// GetAllStatus2xx returns the AllStatus2xx field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetAllStatus2xx() int32 {
	if o == nil || o.AllStatus2xx == nil {
		var ret int32
		return ret
	}
	return *o.AllStatus2xx
}

// GetAllStatus2xxOk returns a tuple with the AllStatus2xx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetAllStatus2xxOk() (*int32, bool) {
	if o == nil || o.AllStatus2xx == nil {
		return nil, false
	}
	return o.AllStatus2xx, true
}

// HasAllStatus2xx returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasAllStatus2xx() bool {
	if o != nil && o.AllStatus2xx != nil {
		return true
	}

	return false
}

// SetAllStatus2xx gets a reference to the given int32 and assigns it to the AllStatus2xx field.
func (o *OriginInspectorMeasurements) SetAllStatus2xx(v int32) {
	o.AllStatus2xx = &v
}

// GetAllStatus3xx returns the AllStatus3xx field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetAllStatus3xx() int32 {
	if o == nil || o.AllStatus3xx == nil {
		var ret int32
		return ret
	}
	return *o.AllStatus3xx
}

// GetAllStatus3xxOk returns a tuple with the AllStatus3xx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetAllStatus3xxOk() (*int32, bool) {
	if o == nil || o.AllStatus3xx == nil {
		return nil, false
	}
	return o.AllStatus3xx, true
}

// HasAllStatus3xx returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasAllStatus3xx() bool {
	if o != nil && o.AllStatus3xx != nil {
		return true
	}

	return false
}

// SetAllStatus3xx gets a reference to the given int32 and assigns it to the AllStatus3xx field.
func (o *OriginInspectorMeasurements) SetAllStatus3xx(v int32) {
	o.AllStatus3xx = &v
}

// GetAllStatus4xx returns the AllStatus4xx field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetAllStatus4xx() int32 {
	if o == nil || o.AllStatus4xx == nil {
		var ret int32
		return ret
	}
	return *o.AllStatus4xx
}

// GetAllStatus4xxOk returns a tuple with the AllStatus4xx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetAllStatus4xxOk() (*int32, bool) {
	if o == nil || o.AllStatus4xx == nil {
		return nil, false
	}
	return o.AllStatus4xx, true
}

// HasAllStatus4xx returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasAllStatus4xx() bool {
	if o != nil && o.AllStatus4xx != nil {
		return true
	}

	return false
}

// SetAllStatus4xx gets a reference to the given int32 and assigns it to the AllStatus4xx field.
func (o *OriginInspectorMeasurements) SetAllStatus4xx(v int32) {
	o.AllStatus4xx = &v
}

// GetAllStatus5xx returns the AllStatus5xx field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetAllStatus5xx() int32 {
	if o == nil || o.AllStatus5xx == nil {
		var ret int32
		return ret
	}
	return *o.AllStatus5xx
}

// GetAllStatus5xxOk returns a tuple with the AllStatus5xx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetAllStatus5xxOk() (*int32, bool) {
	if o == nil || o.AllStatus5xx == nil {
		return nil, false
	}
	return o.AllStatus5xx, true
}

// HasAllStatus5xx returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasAllStatus5xx() bool {
	if o != nil && o.AllStatus5xx != nil {
		return true
	}

	return false
}

// SetAllStatus5xx gets a reference to the given int32 and assigns it to the AllStatus5xx field.
func (o *OriginInspectorMeasurements) SetAllStatus5xx(v int32) {
	o.AllStatus5xx = &v
}

// GetAllStatus200 returns the AllStatus200 field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetAllStatus200() int32 {
	if o == nil || o.AllStatus200 == nil {
		var ret int32
		return ret
	}
	return *o.AllStatus200
}

// GetAllStatus200Ok returns a tuple with the AllStatus200 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetAllStatus200Ok() (*int32, bool) {
	if o == nil || o.AllStatus200 == nil {
		return nil, false
	}
	return o.AllStatus200, true
}

// HasAllStatus200 returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasAllStatus200() bool {
	if o != nil && o.AllStatus200 != nil {
		return true
	}

	return false
}

// SetAllStatus200 gets a reference to the given int32 and assigns it to the AllStatus200 field.
func (o *OriginInspectorMeasurements) SetAllStatus200(v int32) {
	o.AllStatus200 = &v
}

// GetAllStatus204 returns the AllStatus204 field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetAllStatus204() int32 {
	if o == nil || o.AllStatus204 == nil {
		var ret int32
		return ret
	}
	return *o.AllStatus204
}

// GetAllStatus204Ok returns a tuple with the AllStatus204 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetAllStatus204Ok() (*int32, bool) {
	if o == nil || o.AllStatus204 == nil {
		return nil, false
	}
	return o.AllStatus204, true
}

// HasAllStatus204 returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasAllStatus204() bool {
	if o != nil && o.AllStatus204 != nil {
		return true
	}

	return false
}

// SetAllStatus204 gets a reference to the given int32 and assigns it to the AllStatus204 field.
func (o *OriginInspectorMeasurements) SetAllStatus204(v int32) {
	o.AllStatus204 = &v
}

// GetAllStatus206 returns the AllStatus206 field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetAllStatus206() int32 {
	if o == nil || o.AllStatus206 == nil {
		var ret int32
		return ret
	}
	return *o.AllStatus206
}

// GetAllStatus206Ok returns a tuple with the AllStatus206 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetAllStatus206Ok() (*int32, bool) {
	if o == nil || o.AllStatus206 == nil {
		return nil, false
	}
	return o.AllStatus206, true
}

// HasAllStatus206 returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasAllStatus206() bool {
	if o != nil && o.AllStatus206 != nil {
		return true
	}

	return false
}

// SetAllStatus206 gets a reference to the given int32 and assigns it to the AllStatus206 field.
func (o *OriginInspectorMeasurements) SetAllStatus206(v int32) {
	o.AllStatus206 = &v
}

// GetAllStatus301 returns the AllStatus301 field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetAllStatus301() int32 {
	if o == nil || o.AllStatus301 == nil {
		var ret int32
		return ret
	}
	return *o.AllStatus301
}

// GetAllStatus301Ok returns a tuple with the AllStatus301 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetAllStatus301Ok() (*int32, bool) {
	if o == nil || o.AllStatus301 == nil {
		return nil, false
	}
	return o.AllStatus301, true
}

// HasAllStatus301 returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasAllStatus301() bool {
	if o != nil && o.AllStatus301 != nil {
		return true
	}

	return false
}

// SetAllStatus301 gets a reference to the given int32 and assigns it to the AllStatus301 field.
func (o *OriginInspectorMeasurements) SetAllStatus301(v int32) {
	o.AllStatus301 = &v
}

// GetAllStatus302 returns the AllStatus302 field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetAllStatus302() int32 {
	if o == nil || o.AllStatus302 == nil {
		var ret int32
		return ret
	}
	return *o.AllStatus302
}

// GetAllStatus302Ok returns a tuple with the AllStatus302 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetAllStatus302Ok() (*int32, bool) {
	if o == nil || o.AllStatus302 == nil {
		return nil, false
	}
	return o.AllStatus302, true
}

// HasAllStatus302 returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasAllStatus302() bool {
	if o != nil && o.AllStatus302 != nil {
		return true
	}

	return false
}

// SetAllStatus302 gets a reference to the given int32 and assigns it to the AllStatus302 field.
func (o *OriginInspectorMeasurements) SetAllStatus302(v int32) {
	o.AllStatus302 = &v
}

// GetAllStatus304 returns the AllStatus304 field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetAllStatus304() int32 {
	if o == nil || o.AllStatus304 == nil {
		var ret int32
		return ret
	}
	return *o.AllStatus304
}

// GetAllStatus304Ok returns a tuple with the AllStatus304 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetAllStatus304Ok() (*int32, bool) {
	if o == nil || o.AllStatus304 == nil {
		return nil, false
	}
	return o.AllStatus304, true
}

// HasAllStatus304 returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasAllStatus304() bool {
	if o != nil && o.AllStatus304 != nil {
		return true
	}

	return false
}

// SetAllStatus304 gets a reference to the given int32 and assigns it to the AllStatus304 field.
func (o *OriginInspectorMeasurements) SetAllStatus304(v int32) {
	o.AllStatus304 = &v
}

// GetAllStatus400 returns the AllStatus400 field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetAllStatus400() int32 {
	if o == nil || o.AllStatus400 == nil {
		var ret int32
		return ret
	}
	return *o.AllStatus400
}

// GetAllStatus400Ok returns a tuple with the AllStatus400 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetAllStatus400Ok() (*int32, bool) {
	if o == nil || o.AllStatus400 == nil {
		return nil, false
	}
	return o.AllStatus400, true
}

// HasAllStatus400 returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasAllStatus400() bool {
	if o != nil && o.AllStatus400 != nil {
		return true
	}

	return false
}

// SetAllStatus400 gets a reference to the given int32 and assigns it to the AllStatus400 field.
func (o *OriginInspectorMeasurements) SetAllStatus400(v int32) {
	o.AllStatus400 = &v
}

// GetAllStatus401 returns the AllStatus401 field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetAllStatus401() int32 {
	if o == nil || o.AllStatus401 == nil {
		var ret int32
		return ret
	}
	return *o.AllStatus401
}

// GetAllStatus401Ok returns a tuple with the AllStatus401 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetAllStatus401Ok() (*int32, bool) {
	if o == nil || o.AllStatus401 == nil {
		return nil, false
	}
	return o.AllStatus401, true
}

// HasAllStatus401 returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasAllStatus401() bool {
	if o != nil && o.AllStatus401 != nil {
		return true
	}

	return false
}

// SetAllStatus401 gets a reference to the given int32 and assigns it to the AllStatus401 field.
func (o *OriginInspectorMeasurements) SetAllStatus401(v int32) {
	o.AllStatus401 = &v
}

// GetAllStatus403 returns the AllStatus403 field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetAllStatus403() int32 {
	if o == nil || o.AllStatus403 == nil {
		var ret int32
		return ret
	}
	return *o.AllStatus403
}

// GetAllStatus403Ok returns a tuple with the AllStatus403 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetAllStatus403Ok() (*int32, bool) {
	if o == nil || o.AllStatus403 == nil {
		return nil, false
	}
	return o.AllStatus403, true
}

// HasAllStatus403 returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasAllStatus403() bool {
	if o != nil && o.AllStatus403 != nil {
		return true
	}

	return false
}

// SetAllStatus403 gets a reference to the given int32 and assigns it to the AllStatus403 field.
func (o *OriginInspectorMeasurements) SetAllStatus403(v int32) {
	o.AllStatus403 = &v
}

// GetAllStatus404 returns the AllStatus404 field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetAllStatus404() int32 {
	if o == nil || o.AllStatus404 == nil {
		var ret int32
		return ret
	}
	return *o.AllStatus404
}

// GetAllStatus404Ok returns a tuple with the AllStatus404 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetAllStatus404Ok() (*int32, bool) {
	if o == nil || o.AllStatus404 == nil {
		return nil, false
	}
	return o.AllStatus404, true
}

// HasAllStatus404 returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasAllStatus404() bool {
	if o != nil && o.AllStatus404 != nil {
		return true
	}

	return false
}

// SetAllStatus404 gets a reference to the given int32 and assigns it to the AllStatus404 field.
func (o *OriginInspectorMeasurements) SetAllStatus404(v int32) {
	o.AllStatus404 = &v
}

// GetAllStatus416 returns the AllStatus416 field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetAllStatus416() int32 {
	if o == nil || o.AllStatus416 == nil {
		var ret int32
		return ret
	}
	return *o.AllStatus416
}

// GetAllStatus416Ok returns a tuple with the AllStatus416 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetAllStatus416Ok() (*int32, bool) {
	if o == nil || o.AllStatus416 == nil {
		return nil, false
	}
	return o.AllStatus416, true
}

// HasAllStatus416 returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasAllStatus416() bool {
	if o != nil && o.AllStatus416 != nil {
		return true
	}

	return false
}

// SetAllStatus416 gets a reference to the given int32 and assigns it to the AllStatus416 field.
func (o *OriginInspectorMeasurements) SetAllStatus416(v int32) {
	o.AllStatus416 = &v
}

// GetAllStatus429 returns the AllStatus429 field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetAllStatus429() int32 {
	if o == nil || o.AllStatus429 == nil {
		var ret int32
		return ret
	}
	return *o.AllStatus429
}

// GetAllStatus429Ok returns a tuple with the AllStatus429 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetAllStatus429Ok() (*int32, bool) {
	if o == nil || o.AllStatus429 == nil {
		return nil, false
	}
	return o.AllStatus429, true
}

// HasAllStatus429 returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasAllStatus429() bool {
	if o != nil && o.AllStatus429 != nil {
		return true
	}

	return false
}

// SetAllStatus429 gets a reference to the given int32 and assigns it to the AllStatus429 field.
func (o *OriginInspectorMeasurements) SetAllStatus429(v int32) {
	o.AllStatus429 = &v
}

// GetAllStatus500 returns the AllStatus500 field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetAllStatus500() int32 {
	if o == nil || o.AllStatus500 == nil {
		var ret int32
		return ret
	}
	return *o.AllStatus500
}

// GetAllStatus500Ok returns a tuple with the AllStatus500 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetAllStatus500Ok() (*int32, bool) {
	if o == nil || o.AllStatus500 == nil {
		return nil, false
	}
	return o.AllStatus500, true
}

// HasAllStatus500 returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasAllStatus500() bool {
	if o != nil && o.AllStatus500 != nil {
		return true
	}

	return false
}

// SetAllStatus500 gets a reference to the given int32 and assigns it to the AllStatus500 field.
func (o *OriginInspectorMeasurements) SetAllStatus500(v int32) {
	o.AllStatus500 = &v
}

// GetAllStatus501 returns the AllStatus501 field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetAllStatus501() int32 {
	if o == nil || o.AllStatus501 == nil {
		var ret int32
		return ret
	}
	return *o.AllStatus501
}

// GetAllStatus501Ok returns a tuple with the AllStatus501 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetAllStatus501Ok() (*int32, bool) {
	if o == nil || o.AllStatus501 == nil {
		return nil, false
	}
	return o.AllStatus501, true
}

// HasAllStatus501 returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasAllStatus501() bool {
	if o != nil && o.AllStatus501 != nil {
		return true
	}

	return false
}

// SetAllStatus501 gets a reference to the given int32 and assigns it to the AllStatus501 field.
func (o *OriginInspectorMeasurements) SetAllStatus501(v int32) {
	o.AllStatus501 = &v
}

// GetAllStatus502 returns the AllStatus502 field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetAllStatus502() int32 {
	if o == nil || o.AllStatus502 == nil {
		var ret int32
		return ret
	}
	return *o.AllStatus502
}

// GetAllStatus502Ok returns a tuple with the AllStatus502 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetAllStatus502Ok() (*int32, bool) {
	if o == nil || o.AllStatus502 == nil {
		return nil, false
	}
	return o.AllStatus502, true
}

// HasAllStatus502 returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasAllStatus502() bool {
	if o != nil && o.AllStatus502 != nil {
		return true
	}

	return false
}

// SetAllStatus502 gets a reference to the given int32 and assigns it to the AllStatus502 field.
func (o *OriginInspectorMeasurements) SetAllStatus502(v int32) {
	o.AllStatus502 = &v
}

// GetAllStatus503 returns the AllStatus503 field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetAllStatus503() int32 {
	if o == nil || o.AllStatus503 == nil {
		var ret int32
		return ret
	}
	return *o.AllStatus503
}

// GetAllStatus503Ok returns a tuple with the AllStatus503 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetAllStatus503Ok() (*int32, bool) {
	if o == nil || o.AllStatus503 == nil {
		return nil, false
	}
	return o.AllStatus503, true
}

// HasAllStatus503 returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasAllStatus503() bool {
	if o != nil && o.AllStatus503 != nil {
		return true
	}

	return false
}

// SetAllStatus503 gets a reference to the given int32 and assigns it to the AllStatus503 field.
func (o *OriginInspectorMeasurements) SetAllStatus503(v int32) {
	o.AllStatus503 = &v
}

// GetAllStatus504 returns the AllStatus504 field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetAllStatus504() int32 {
	if o == nil || o.AllStatus504 == nil {
		var ret int32
		return ret
	}
	return *o.AllStatus504
}

// GetAllStatus504Ok returns a tuple with the AllStatus504 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetAllStatus504Ok() (*int32, bool) {
	if o == nil || o.AllStatus504 == nil {
		return nil, false
	}
	return o.AllStatus504, true
}

// HasAllStatus504 returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasAllStatus504() bool {
	if o != nil && o.AllStatus504 != nil {
		return true
	}

	return false
}

// SetAllStatus504 gets a reference to the given int32 and assigns it to the AllStatus504 field.
func (o *OriginInspectorMeasurements) SetAllStatus504(v int32) {
	o.AllStatus504 = &v
}

// GetAllStatus505 returns the AllStatus505 field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetAllStatus505() int32 {
	if o == nil || o.AllStatus505 == nil {
		var ret int32
		return ret
	}
	return *o.AllStatus505
}

// GetAllStatus505Ok returns a tuple with the AllStatus505 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetAllStatus505Ok() (*int32, bool) {
	if o == nil || o.AllStatus505 == nil {
		return nil, false
	}
	return o.AllStatus505, true
}

// HasAllStatus505 returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasAllStatus505() bool {
	if o != nil && o.AllStatus505 != nil {
		return true
	}

	return false
}

// SetAllStatus505 gets a reference to the given int32 and assigns it to the AllStatus505 field.
func (o *OriginInspectorMeasurements) SetAllStatus505(v int32) {
	o.AllStatus505 = &v
}

// GetAllLatency0To1ms returns the AllLatency0To1ms field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetAllLatency0To1ms() int32 {
	if o == nil || o.AllLatency0To1ms == nil {
		var ret int32
		return ret
	}
	return *o.AllLatency0To1ms
}

// GetAllLatency0To1msOk returns a tuple with the AllLatency0To1ms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetAllLatency0To1msOk() (*int32, bool) {
	if o == nil || o.AllLatency0To1ms == nil {
		return nil, false
	}
	return o.AllLatency0To1ms, true
}

// HasAllLatency0To1ms returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasAllLatency0To1ms() bool {
	if o != nil && o.AllLatency0To1ms != nil {
		return true
	}

	return false
}

// SetAllLatency0To1ms gets a reference to the given int32 and assigns it to the AllLatency0To1ms field.
func (o *OriginInspectorMeasurements) SetAllLatency0To1ms(v int32) {
	o.AllLatency0To1ms = &v
}

// GetAllLatency1To5ms returns the AllLatency1To5ms field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetAllLatency1To5ms() int32 {
	if o == nil || o.AllLatency1To5ms == nil {
		var ret int32
		return ret
	}
	return *o.AllLatency1To5ms
}

// GetAllLatency1To5msOk returns a tuple with the AllLatency1To5ms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetAllLatency1To5msOk() (*int32, bool) {
	if o == nil || o.AllLatency1To5ms == nil {
		return nil, false
	}
	return o.AllLatency1To5ms, true
}

// HasAllLatency1To5ms returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasAllLatency1To5ms() bool {
	if o != nil && o.AllLatency1To5ms != nil {
		return true
	}

	return false
}

// SetAllLatency1To5ms gets a reference to the given int32 and assigns it to the AllLatency1To5ms field.
func (o *OriginInspectorMeasurements) SetAllLatency1To5ms(v int32) {
	o.AllLatency1To5ms = &v
}

// GetAllLatency5To10ms returns the AllLatency5To10ms field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetAllLatency5To10ms() int32 {
	if o == nil || o.AllLatency5To10ms == nil {
		var ret int32
		return ret
	}
	return *o.AllLatency5To10ms
}

// GetAllLatency5To10msOk returns a tuple with the AllLatency5To10ms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetAllLatency5To10msOk() (*int32, bool) {
	if o == nil || o.AllLatency5To10ms == nil {
		return nil, false
	}
	return o.AllLatency5To10ms, true
}

// HasAllLatency5To10ms returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasAllLatency5To10ms() bool {
	if o != nil && o.AllLatency5To10ms != nil {
		return true
	}

	return false
}

// SetAllLatency5To10ms gets a reference to the given int32 and assigns it to the AllLatency5To10ms field.
func (o *OriginInspectorMeasurements) SetAllLatency5To10ms(v int32) {
	o.AllLatency5To10ms = &v
}

// GetAllLatency10To50ms returns the AllLatency10To50ms field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetAllLatency10To50ms() int32 {
	if o == nil || o.AllLatency10To50ms == nil {
		var ret int32
		return ret
	}
	return *o.AllLatency10To50ms
}

// GetAllLatency10To50msOk returns a tuple with the AllLatency10To50ms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetAllLatency10To50msOk() (*int32, bool) {
	if o == nil || o.AllLatency10To50ms == nil {
		return nil, false
	}
	return o.AllLatency10To50ms, true
}

// HasAllLatency10To50ms returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasAllLatency10To50ms() bool {
	if o != nil && o.AllLatency10To50ms != nil {
		return true
	}

	return false
}

// SetAllLatency10To50ms gets a reference to the given int32 and assigns it to the AllLatency10To50ms field.
func (o *OriginInspectorMeasurements) SetAllLatency10To50ms(v int32) {
	o.AllLatency10To50ms = &v
}

// GetAllLatency50To100ms returns the AllLatency50To100ms field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetAllLatency50To100ms() int32 {
	if o == nil || o.AllLatency50To100ms == nil {
		var ret int32
		return ret
	}
	return *o.AllLatency50To100ms
}

// GetAllLatency50To100msOk returns a tuple with the AllLatency50To100ms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetAllLatency50To100msOk() (*int32, bool) {
	if o == nil || o.AllLatency50To100ms == nil {
		return nil, false
	}
	return o.AllLatency50To100ms, true
}

// HasAllLatency50To100ms returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasAllLatency50To100ms() bool {
	if o != nil && o.AllLatency50To100ms != nil {
		return true
	}

	return false
}

// SetAllLatency50To100ms gets a reference to the given int32 and assigns it to the AllLatency50To100ms field.
func (o *OriginInspectorMeasurements) SetAllLatency50To100ms(v int32) {
	o.AllLatency50To100ms = &v
}

// GetAllLatency100To250ms returns the AllLatency100To250ms field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetAllLatency100To250ms() int32 {
	if o == nil || o.AllLatency100To250ms == nil {
		var ret int32
		return ret
	}
	return *o.AllLatency100To250ms
}

// GetAllLatency100To250msOk returns a tuple with the AllLatency100To250ms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetAllLatency100To250msOk() (*int32, bool) {
	if o == nil || o.AllLatency100To250ms == nil {
		return nil, false
	}
	return o.AllLatency100To250ms, true
}

// HasAllLatency100To250ms returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasAllLatency100To250ms() bool {
	if o != nil && o.AllLatency100To250ms != nil {
		return true
	}

	return false
}

// SetAllLatency100To250ms gets a reference to the given int32 and assigns it to the AllLatency100To250ms field.
func (o *OriginInspectorMeasurements) SetAllLatency100To250ms(v int32) {
	o.AllLatency100To250ms = &v
}

// GetAllLatency250To500ms returns the AllLatency250To500ms field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetAllLatency250To500ms() int32 {
	if o == nil || o.AllLatency250To500ms == nil {
		var ret int32
		return ret
	}
	return *o.AllLatency250To500ms
}

// GetAllLatency250To500msOk returns a tuple with the AllLatency250To500ms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetAllLatency250To500msOk() (*int32, bool) {
	if o == nil || o.AllLatency250To500ms == nil {
		return nil, false
	}
	return o.AllLatency250To500ms, true
}

// HasAllLatency250To500ms returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasAllLatency250To500ms() bool {
	if o != nil && o.AllLatency250To500ms != nil {
		return true
	}

	return false
}

// SetAllLatency250To500ms gets a reference to the given int32 and assigns it to the AllLatency250To500ms field.
func (o *OriginInspectorMeasurements) SetAllLatency250To500ms(v int32) {
	o.AllLatency250To500ms = &v
}

// GetAllLatency500To1000ms returns the AllLatency500To1000ms field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetAllLatency500To1000ms() int32 {
	if o == nil || o.AllLatency500To1000ms == nil {
		var ret int32
		return ret
	}
	return *o.AllLatency500To1000ms
}

// GetAllLatency500To1000msOk returns a tuple with the AllLatency500To1000ms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetAllLatency500To1000msOk() (*int32, bool) {
	if o == nil || o.AllLatency500To1000ms == nil {
		return nil, false
	}
	return o.AllLatency500To1000ms, true
}

// HasAllLatency500To1000ms returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasAllLatency500To1000ms() bool {
	if o != nil && o.AllLatency500To1000ms != nil {
		return true
	}

	return false
}

// SetAllLatency500To1000ms gets a reference to the given int32 and assigns it to the AllLatency500To1000ms field.
func (o *OriginInspectorMeasurements) SetAllLatency500To1000ms(v int32) {
	o.AllLatency500To1000ms = &v
}

// GetAllLatency1000To5000ms returns the AllLatency1000To5000ms field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetAllLatency1000To5000ms() int32 {
	if o == nil || o.AllLatency1000To5000ms == nil {
		var ret int32
		return ret
	}
	return *o.AllLatency1000To5000ms
}

// GetAllLatency1000To5000msOk returns a tuple with the AllLatency1000To5000ms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetAllLatency1000To5000msOk() (*int32, bool) {
	if o == nil || o.AllLatency1000To5000ms == nil {
		return nil, false
	}
	return o.AllLatency1000To5000ms, true
}

// HasAllLatency1000To5000ms returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasAllLatency1000To5000ms() bool {
	if o != nil && o.AllLatency1000To5000ms != nil {
		return true
	}

	return false
}

// SetAllLatency1000To5000ms gets a reference to the given int32 and assigns it to the AllLatency1000To5000ms field.
func (o *OriginInspectorMeasurements) SetAllLatency1000To5000ms(v int32) {
	o.AllLatency1000To5000ms = &v
}

// GetAllLatency5000To10000ms returns the AllLatency5000To10000ms field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetAllLatency5000To10000ms() int32 {
	if o == nil || o.AllLatency5000To10000ms == nil {
		var ret int32
		return ret
	}
	return *o.AllLatency5000To10000ms
}

// GetAllLatency5000To10000msOk returns a tuple with the AllLatency5000To10000ms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetAllLatency5000To10000msOk() (*int32, bool) {
	if o == nil || o.AllLatency5000To10000ms == nil {
		return nil, false
	}
	return o.AllLatency5000To10000ms, true
}

// HasAllLatency5000To10000ms returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasAllLatency5000To10000ms() bool {
	if o != nil && o.AllLatency5000To10000ms != nil {
		return true
	}

	return false
}

// SetAllLatency5000To10000ms gets a reference to the given int32 and assigns it to the AllLatency5000To10000ms field.
func (o *OriginInspectorMeasurements) SetAllLatency5000To10000ms(v int32) {
	o.AllLatency5000To10000ms = &v
}

// GetAllLatency10000To60000ms returns the AllLatency10000To60000ms field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetAllLatency10000To60000ms() int32 {
	if o == nil || o.AllLatency10000To60000ms == nil {
		var ret int32
		return ret
	}
	return *o.AllLatency10000To60000ms
}

// GetAllLatency10000To60000msOk returns a tuple with the AllLatency10000To60000ms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetAllLatency10000To60000msOk() (*int32, bool) {
	if o == nil || o.AllLatency10000To60000ms == nil {
		return nil, false
	}
	return o.AllLatency10000To60000ms, true
}

// HasAllLatency10000To60000ms returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasAllLatency10000To60000ms() bool {
	if o != nil && o.AllLatency10000To60000ms != nil {
		return true
	}

	return false
}

// SetAllLatency10000To60000ms gets a reference to the given int32 and assigns it to the AllLatency10000To60000ms field.
func (o *OriginInspectorMeasurements) SetAllLatency10000To60000ms(v int32) {
	o.AllLatency10000To60000ms = &v
}

// GetAllLatency60000ms returns the AllLatency60000ms field value if set, zero value otherwise.
func (o *OriginInspectorMeasurements) GetAllLatency60000ms() int32 {
	if o == nil || o.AllLatency60000ms == nil {
		var ret int32
		return ret
	}
	return *o.AllLatency60000ms
}

// GetAllLatency60000msOk returns a tuple with the AllLatency60000ms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorMeasurements) GetAllLatency60000msOk() (*int32, bool) {
	if o == nil || o.AllLatency60000ms == nil {
		return nil, false
	}
	return o.AllLatency60000ms, true
}

// HasAllLatency60000ms returns a boolean if a field has been set.
func (o *OriginInspectorMeasurements) HasAllLatency60000ms() bool {
	if o != nil && o.AllLatency60000ms != nil {
		return true
	}

	return false
}

// SetAllLatency60000ms gets a reference to the given int32 and assigns it to the AllLatency60000ms field.
func (o *OriginInspectorMeasurements) SetAllLatency60000ms(v int32) {
	o.AllLatency60000ms = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o OriginInspectorMeasurements) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Responses != nil {
		toSerialize["responses"] = o.Responses
	}
	if o.RespHeaderBytes != nil {
		toSerialize["resp_header_bytes"] = o.RespHeaderBytes
	}
	if o.RespBodyBytes != nil {
		toSerialize["resp_body_bytes"] = o.RespBodyBytes
	}
	if o.Status1xx != nil {
		toSerialize["status_1xx"] = o.Status1xx
	}
	if o.Status2xx != nil {
		toSerialize["status_2xx"] = o.Status2xx
	}
	if o.Status3xx != nil {
		toSerialize["status_3xx"] = o.Status3xx
	}
	if o.Status4xx != nil {
		toSerialize["status_4xx"] = o.Status4xx
	}
	if o.Status5xx != nil {
		toSerialize["status_5xx"] = o.Status5xx
	}
	if o.Status200 != nil {
		toSerialize["status_200"] = o.Status200
	}
	if o.Status204 != nil {
		toSerialize["status_204"] = o.Status204
	}
	if o.Status206 != nil {
		toSerialize["status_206"] = o.Status206
	}
	if o.Status301 != nil {
		toSerialize["status_301"] = o.Status301
	}
	if o.Status302 != nil {
		toSerialize["status_302"] = o.Status302
	}
	if o.Status304 != nil {
		toSerialize["status_304"] = o.Status304
	}
	if o.Status400 != nil {
		toSerialize["status_400"] = o.Status400
	}
	if o.Status401 != nil {
		toSerialize["status_401"] = o.Status401
	}
	if o.Status403 != nil {
		toSerialize["status_403"] = o.Status403
	}
	if o.Status404 != nil {
		toSerialize["status_404"] = o.Status404
	}
	if o.Status416 != nil {
		toSerialize["status_416"] = o.Status416
	}
	if o.Status429 != nil {
		toSerialize["status_429"] = o.Status429
	}
	if o.Status500 != nil {
		toSerialize["status_500"] = o.Status500
	}
	if o.Status501 != nil {
		toSerialize["status_501"] = o.Status501
	}
	if o.Status502 != nil {
		toSerialize["status_502"] = o.Status502
	}
	if o.Status503 != nil {
		toSerialize["status_503"] = o.Status503
	}
	if o.Status504 != nil {
		toSerialize["status_504"] = o.Status504
	}
	if o.Status505 != nil {
		toSerialize["status_505"] = o.Status505
	}
	if o.Latency0To1ms != nil {
		toSerialize["latency_0_to_1ms"] = o.Latency0To1ms
	}
	if o.Latency1To5ms != nil {
		toSerialize["latency_1_to_5ms"] = o.Latency1To5ms
	}
	if o.Latency5To10ms != nil {
		toSerialize["latency_5_to_10ms"] = o.Latency5To10ms
	}
	if o.Latency10To50ms != nil {
		toSerialize["latency_10_to_50ms"] = o.Latency10To50ms
	}
	if o.Latency50To100ms != nil {
		toSerialize["latency_50_to_100ms"] = o.Latency50To100ms
	}
	if o.Latency100To250ms != nil {
		toSerialize["latency_100_to_250ms"] = o.Latency100To250ms
	}
	if o.Latency250To500ms != nil {
		toSerialize["latency_250_to_500ms"] = o.Latency250To500ms
	}
	if o.Latency500To1000ms != nil {
		toSerialize["latency_500_to_1000ms"] = o.Latency500To1000ms
	}
	if o.Latency1000To5000ms != nil {
		toSerialize["latency_1000_to_5000ms"] = o.Latency1000To5000ms
	}
	if o.Latency5000To10000ms != nil {
		toSerialize["latency_5000_to_10000ms"] = o.Latency5000To10000ms
	}
	if o.Latency10000To60000ms != nil {
		toSerialize["latency_10000_to_60000ms"] = o.Latency10000To60000ms
	}
	if o.Latency60000ms != nil {
		toSerialize["latency_60000ms"] = o.Latency60000ms
	}
	if o.WafResponses != nil {
		toSerialize["waf_responses"] = o.WafResponses
	}
	if o.WafRespHeaderBytes != nil {
		toSerialize["waf_resp_header_bytes"] = o.WafRespHeaderBytes
	}
	if o.WafRespBodyBytes != nil {
		toSerialize["waf_resp_body_bytes"] = o.WafRespBodyBytes
	}
	if o.WafStatus1xx != nil {
		toSerialize["waf_status_1xx"] = o.WafStatus1xx
	}
	if o.WafStatus2xx != nil {
		toSerialize["waf_status_2xx"] = o.WafStatus2xx
	}
	if o.WafStatus3xx != nil {
		toSerialize["waf_status_3xx"] = o.WafStatus3xx
	}
	if o.WafStatus4xx != nil {
		toSerialize["waf_status_4xx"] = o.WafStatus4xx
	}
	if o.WafStatus5xx != nil {
		toSerialize["waf_status_5xx"] = o.WafStatus5xx
	}
	if o.WafStatus200 != nil {
		toSerialize["waf_status_200"] = o.WafStatus200
	}
	if o.WafStatus204 != nil {
		toSerialize["waf_status_204"] = o.WafStatus204
	}
	if o.WafStatus206 != nil {
		toSerialize["waf_status_206"] = o.WafStatus206
	}
	if o.WafStatus301 != nil {
		toSerialize["waf_status_301"] = o.WafStatus301
	}
	if o.WafStatus302 != nil {
		toSerialize["waf_status_302"] = o.WafStatus302
	}
	if o.WafStatus304 != nil {
		toSerialize["waf_status_304"] = o.WafStatus304
	}
	if o.WafStatus400 != nil {
		toSerialize["waf_status_400"] = o.WafStatus400
	}
	if o.WafStatus401 != nil {
		toSerialize["waf_status_401"] = o.WafStatus401
	}
	if o.WafStatus403 != nil {
		toSerialize["waf_status_403"] = o.WafStatus403
	}
	if o.WafStatus404 != nil {
		toSerialize["waf_status_404"] = o.WafStatus404
	}
	if o.WafStatus416 != nil {
		toSerialize["waf_status_416"] = o.WafStatus416
	}
	if o.WafStatus429 != nil {
		toSerialize["waf_status_429"] = o.WafStatus429
	}
	if o.WafStatus500 != nil {
		toSerialize["waf_status_500"] = o.WafStatus500
	}
	if o.WafStatus501 != nil {
		toSerialize["waf_status_501"] = o.WafStatus501
	}
	if o.WafStatus502 != nil {
		toSerialize["waf_status_502"] = o.WafStatus502
	}
	if o.WafStatus503 != nil {
		toSerialize["waf_status_503"] = o.WafStatus503
	}
	if o.WafStatus504 != nil {
		toSerialize["waf_status_504"] = o.WafStatus504
	}
	if o.WafStatus505 != nil {
		toSerialize["waf_status_505"] = o.WafStatus505
	}
	if o.WafLatency0To1ms != nil {
		toSerialize["waf_latency_0_to_1ms"] = o.WafLatency0To1ms
	}
	if o.WafLatency1To5ms != nil {
		toSerialize["waf_latency_1_to_5ms"] = o.WafLatency1To5ms
	}
	if o.WafLatency5To10ms != nil {
		toSerialize["waf_latency_5_to_10ms"] = o.WafLatency5To10ms
	}
	if o.WafLatency10To50ms != nil {
		toSerialize["waf_latency_10_to_50ms"] = o.WafLatency10To50ms
	}
	if o.WafLatency50To100ms != nil {
		toSerialize["waf_latency_50_to_100ms"] = o.WafLatency50To100ms
	}
	if o.WafLatency100To250ms != nil {
		toSerialize["waf_latency_100_to_250ms"] = o.WafLatency100To250ms
	}
	if o.WafLatency250To500ms != nil {
		toSerialize["waf_latency_250_to_500ms"] = o.WafLatency250To500ms
	}
	if o.WafLatency500To1000ms != nil {
		toSerialize["waf_latency_500_to_1000ms"] = o.WafLatency500To1000ms
	}
	if o.WafLatency1000To5000ms != nil {
		toSerialize["waf_latency_1000_to_5000ms"] = o.WafLatency1000To5000ms
	}
	if o.WafLatency5000To10000ms != nil {
		toSerialize["waf_latency_5000_to_10000ms"] = o.WafLatency5000To10000ms
	}
	if o.WafLatency10000To60000ms != nil {
		toSerialize["waf_latency_10000_to_60000ms"] = o.WafLatency10000To60000ms
	}
	if o.WafLatency60000ms != nil {
		toSerialize["waf_latency_60000ms"] = o.WafLatency60000ms
	}
	if o.ComputeResponses != nil {
		toSerialize["compute_responses"] = o.ComputeResponses
	}
	if o.ComputeRespHeaderBytes != nil {
		toSerialize["compute_resp_header_bytes"] = o.ComputeRespHeaderBytes
	}
	if o.ComputeRespBodyBytes != nil {
		toSerialize["compute_resp_body_bytes"] = o.ComputeRespBodyBytes
	}
	if o.ComputeStatus1xx != nil {
		toSerialize["compute_status_1xx"] = o.ComputeStatus1xx
	}
	if o.ComputeStatus2xx != nil {
		toSerialize["compute_status_2xx"] = o.ComputeStatus2xx
	}
	if o.ComputeStatus3xx != nil {
		toSerialize["compute_status_3xx"] = o.ComputeStatus3xx
	}
	if o.ComputeStatus4xx != nil {
		toSerialize["compute_status_4xx"] = o.ComputeStatus4xx
	}
	if o.ComputeStatus5xx != nil {
		toSerialize["compute_status_5xx"] = o.ComputeStatus5xx
	}
	if o.ComputeStatus200 != nil {
		toSerialize["compute_status_200"] = o.ComputeStatus200
	}
	if o.ComputeStatus204 != nil {
		toSerialize["compute_status_204"] = o.ComputeStatus204
	}
	if o.ComputeStatus206 != nil {
		toSerialize["compute_status_206"] = o.ComputeStatus206
	}
	if o.ComputeStatus301 != nil {
		toSerialize["compute_status_301"] = o.ComputeStatus301
	}
	if o.ComputeStatus302 != nil {
		toSerialize["compute_status_302"] = o.ComputeStatus302
	}
	if o.ComputeStatus304 != nil {
		toSerialize["compute_status_304"] = o.ComputeStatus304
	}
	if o.ComputeStatus400 != nil {
		toSerialize["compute_status_400"] = o.ComputeStatus400
	}
	if o.ComputeStatus401 != nil {
		toSerialize["compute_status_401"] = o.ComputeStatus401
	}
	if o.ComputeStatus403 != nil {
		toSerialize["compute_status_403"] = o.ComputeStatus403
	}
	if o.ComputeStatus404 != nil {
		toSerialize["compute_status_404"] = o.ComputeStatus404
	}
	if o.ComputeStatus416 != nil {
		toSerialize["compute_status_416"] = o.ComputeStatus416
	}
	if o.ComputeStatus429 != nil {
		toSerialize["compute_status_429"] = o.ComputeStatus429
	}
	if o.ComputeStatus500 != nil {
		toSerialize["compute_status_500"] = o.ComputeStatus500
	}
	if o.ComputeStatus501 != nil {
		toSerialize["compute_status_501"] = o.ComputeStatus501
	}
	if o.ComputeStatus502 != nil {
		toSerialize["compute_status_502"] = o.ComputeStatus502
	}
	if o.ComputeStatus503 != nil {
		toSerialize["compute_status_503"] = o.ComputeStatus503
	}
	if o.ComputeStatus504 != nil {
		toSerialize["compute_status_504"] = o.ComputeStatus504
	}
	if o.ComputeStatus505 != nil {
		toSerialize["compute_status_505"] = o.ComputeStatus505
	}
	if o.ComputeLatency0To1ms != nil {
		toSerialize["compute_latency_0_to_1ms"] = o.ComputeLatency0To1ms
	}
	if o.ComputeLatency1To5ms != nil {
		toSerialize["compute_latency_1_to_5ms"] = o.ComputeLatency1To5ms
	}
	if o.ComputeLatency5To10ms != nil {
		toSerialize["compute_latency_5_to_10ms"] = o.ComputeLatency5To10ms
	}
	if o.ComputeLatency10To50ms != nil {
		toSerialize["compute_latency_10_to_50ms"] = o.ComputeLatency10To50ms
	}
	if o.ComputeLatency50To100ms != nil {
		toSerialize["compute_latency_50_to_100ms"] = o.ComputeLatency50To100ms
	}
	if o.ComputeLatency100To250ms != nil {
		toSerialize["compute_latency_100_to_250ms"] = o.ComputeLatency100To250ms
	}
	if o.ComputeLatency250To500ms != nil {
		toSerialize["compute_latency_250_to_500ms"] = o.ComputeLatency250To500ms
	}
	if o.ComputeLatency500To1000ms != nil {
		toSerialize["compute_latency_500_to_1000ms"] = o.ComputeLatency500To1000ms
	}
	if o.ComputeLatency1000To5000ms != nil {
		toSerialize["compute_latency_1000_to_5000ms"] = o.ComputeLatency1000To5000ms
	}
	if o.ComputeLatency5000To10000ms != nil {
		toSerialize["compute_latency_5000_to_10000ms"] = o.ComputeLatency5000To10000ms
	}
	if o.ComputeLatency10000To60000ms != nil {
		toSerialize["compute_latency_10000_to_60000ms"] = o.ComputeLatency10000To60000ms
	}
	if o.ComputeLatency60000ms != nil {
		toSerialize["compute_latency_60000ms"] = o.ComputeLatency60000ms
	}
	if o.AllResponses != nil {
		toSerialize["all_responses"] = o.AllResponses
	}
	if o.AllRespHeaderBytes != nil {
		toSerialize["all_resp_header_bytes"] = o.AllRespHeaderBytes
	}
	if o.AllRespBodyBytes != nil {
		toSerialize["all_resp_body_bytes"] = o.AllRespBodyBytes
	}
	if o.AllStatus1xx != nil {
		toSerialize["all_status_1xx"] = o.AllStatus1xx
	}
	if o.AllStatus2xx != nil {
		toSerialize["all_status_2xx"] = o.AllStatus2xx
	}
	if o.AllStatus3xx != nil {
		toSerialize["all_status_3xx"] = o.AllStatus3xx
	}
	if o.AllStatus4xx != nil {
		toSerialize["all_status_4xx"] = o.AllStatus4xx
	}
	if o.AllStatus5xx != nil {
		toSerialize["all_status_5xx"] = o.AllStatus5xx
	}
	if o.AllStatus200 != nil {
		toSerialize["all_status_200"] = o.AllStatus200
	}
	if o.AllStatus204 != nil {
		toSerialize["all_status_204"] = o.AllStatus204
	}
	if o.AllStatus206 != nil {
		toSerialize["all_status_206"] = o.AllStatus206
	}
	if o.AllStatus301 != nil {
		toSerialize["all_status_301"] = o.AllStatus301
	}
	if o.AllStatus302 != nil {
		toSerialize["all_status_302"] = o.AllStatus302
	}
	if o.AllStatus304 != nil {
		toSerialize["all_status_304"] = o.AllStatus304
	}
	if o.AllStatus400 != nil {
		toSerialize["all_status_400"] = o.AllStatus400
	}
	if o.AllStatus401 != nil {
		toSerialize["all_status_401"] = o.AllStatus401
	}
	if o.AllStatus403 != nil {
		toSerialize["all_status_403"] = o.AllStatus403
	}
	if o.AllStatus404 != nil {
		toSerialize["all_status_404"] = o.AllStatus404
	}
	if o.AllStatus416 != nil {
		toSerialize["all_status_416"] = o.AllStatus416
	}
	if o.AllStatus429 != nil {
		toSerialize["all_status_429"] = o.AllStatus429
	}
	if o.AllStatus500 != nil {
		toSerialize["all_status_500"] = o.AllStatus500
	}
	if o.AllStatus501 != nil {
		toSerialize["all_status_501"] = o.AllStatus501
	}
	if o.AllStatus502 != nil {
		toSerialize["all_status_502"] = o.AllStatus502
	}
	if o.AllStatus503 != nil {
		toSerialize["all_status_503"] = o.AllStatus503
	}
	if o.AllStatus504 != nil {
		toSerialize["all_status_504"] = o.AllStatus504
	}
	if o.AllStatus505 != nil {
		toSerialize["all_status_505"] = o.AllStatus505
	}
	if o.AllLatency0To1ms != nil {
		toSerialize["all_latency_0_to_1ms"] = o.AllLatency0To1ms
	}
	if o.AllLatency1To5ms != nil {
		toSerialize["all_latency_1_to_5ms"] = o.AllLatency1To5ms
	}
	if o.AllLatency5To10ms != nil {
		toSerialize["all_latency_5_to_10ms"] = o.AllLatency5To10ms
	}
	if o.AllLatency10To50ms != nil {
		toSerialize["all_latency_10_to_50ms"] = o.AllLatency10To50ms
	}
	if o.AllLatency50To100ms != nil {
		toSerialize["all_latency_50_to_100ms"] = o.AllLatency50To100ms
	}
	if o.AllLatency100To250ms != nil {
		toSerialize["all_latency_100_to_250ms"] = o.AllLatency100To250ms
	}
	if o.AllLatency250To500ms != nil {
		toSerialize["all_latency_250_to_500ms"] = o.AllLatency250To500ms
	}
	if o.AllLatency500To1000ms != nil {
		toSerialize["all_latency_500_to_1000ms"] = o.AllLatency500To1000ms
	}
	if o.AllLatency1000To5000ms != nil {
		toSerialize["all_latency_1000_to_5000ms"] = o.AllLatency1000To5000ms
	}
	if o.AllLatency5000To10000ms != nil {
		toSerialize["all_latency_5000_to_10000ms"] = o.AllLatency5000To10000ms
	}
	if o.AllLatency10000To60000ms != nil {
		toSerialize["all_latency_10000_to_60000ms"] = o.AllLatency10000To60000ms
	}
	if o.AllLatency60000ms != nil {
		toSerialize["all_latency_60000ms"] = o.AllLatency60000ms
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *OriginInspectorMeasurements) UnmarshalJSON(bytes []byte) (err error) {
	varOriginInspectorMeasurements := _OriginInspectorMeasurements{}

	if err = json.Unmarshal(bytes, &varOriginInspectorMeasurements); err == nil {
		*o = OriginInspectorMeasurements(varOriginInspectorMeasurements)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "responses")
		delete(additionalProperties, "resp_header_bytes")
		delete(additionalProperties, "resp_body_bytes")
		delete(additionalProperties, "status_1xx")
		delete(additionalProperties, "status_2xx")
		delete(additionalProperties, "status_3xx")
		delete(additionalProperties, "status_4xx")
		delete(additionalProperties, "status_5xx")
		delete(additionalProperties, "status_200")
		delete(additionalProperties, "status_204")
		delete(additionalProperties, "status_206")
		delete(additionalProperties, "status_301")
		delete(additionalProperties, "status_302")
		delete(additionalProperties, "status_304")
		delete(additionalProperties, "status_400")
		delete(additionalProperties, "status_401")
		delete(additionalProperties, "status_403")
		delete(additionalProperties, "status_404")
		delete(additionalProperties, "status_416")
		delete(additionalProperties, "status_429")
		delete(additionalProperties, "status_500")
		delete(additionalProperties, "status_501")
		delete(additionalProperties, "status_502")
		delete(additionalProperties, "status_503")
		delete(additionalProperties, "status_504")
		delete(additionalProperties, "status_505")
		delete(additionalProperties, "latency_0_to_1ms")
		delete(additionalProperties, "latency_1_to_5ms")
		delete(additionalProperties, "latency_5_to_10ms")
		delete(additionalProperties, "latency_10_to_50ms")
		delete(additionalProperties, "latency_50_to_100ms")
		delete(additionalProperties, "latency_100_to_250ms")
		delete(additionalProperties, "latency_250_to_500ms")
		delete(additionalProperties, "latency_500_to_1000ms")
		delete(additionalProperties, "latency_1000_to_5000ms")
		delete(additionalProperties, "latency_5000_to_10000ms")
		delete(additionalProperties, "latency_10000_to_60000ms")
		delete(additionalProperties, "latency_60000ms")
		delete(additionalProperties, "waf_responses")
		delete(additionalProperties, "waf_resp_header_bytes")
		delete(additionalProperties, "waf_resp_body_bytes")
		delete(additionalProperties, "waf_status_1xx")
		delete(additionalProperties, "waf_status_2xx")
		delete(additionalProperties, "waf_status_3xx")
		delete(additionalProperties, "waf_status_4xx")
		delete(additionalProperties, "waf_status_5xx")
		delete(additionalProperties, "waf_status_200")
		delete(additionalProperties, "waf_status_204")
		delete(additionalProperties, "waf_status_206")
		delete(additionalProperties, "waf_status_301")
		delete(additionalProperties, "waf_status_302")
		delete(additionalProperties, "waf_status_304")
		delete(additionalProperties, "waf_status_400")
		delete(additionalProperties, "waf_status_401")
		delete(additionalProperties, "waf_status_403")
		delete(additionalProperties, "waf_status_404")
		delete(additionalProperties, "waf_status_416")
		delete(additionalProperties, "waf_status_429")
		delete(additionalProperties, "waf_status_500")
		delete(additionalProperties, "waf_status_501")
		delete(additionalProperties, "waf_status_502")
		delete(additionalProperties, "waf_status_503")
		delete(additionalProperties, "waf_status_504")
		delete(additionalProperties, "waf_status_505")
		delete(additionalProperties, "waf_latency_0_to_1ms")
		delete(additionalProperties, "waf_latency_1_to_5ms")
		delete(additionalProperties, "waf_latency_5_to_10ms")
		delete(additionalProperties, "waf_latency_10_to_50ms")
		delete(additionalProperties, "waf_latency_50_to_100ms")
		delete(additionalProperties, "waf_latency_100_to_250ms")
		delete(additionalProperties, "waf_latency_250_to_500ms")
		delete(additionalProperties, "waf_latency_500_to_1000ms")
		delete(additionalProperties, "waf_latency_1000_to_5000ms")
		delete(additionalProperties, "waf_latency_5000_to_10000ms")
		delete(additionalProperties, "waf_latency_10000_to_60000ms")
		delete(additionalProperties, "waf_latency_60000ms")
		delete(additionalProperties, "compute_responses")
		delete(additionalProperties, "compute_resp_header_bytes")
		delete(additionalProperties, "compute_resp_body_bytes")
		delete(additionalProperties, "compute_status_1xx")
		delete(additionalProperties, "compute_status_2xx")
		delete(additionalProperties, "compute_status_3xx")
		delete(additionalProperties, "compute_status_4xx")
		delete(additionalProperties, "compute_status_5xx")
		delete(additionalProperties, "compute_status_200")
		delete(additionalProperties, "compute_status_204")
		delete(additionalProperties, "compute_status_206")
		delete(additionalProperties, "compute_status_301")
		delete(additionalProperties, "compute_status_302")
		delete(additionalProperties, "compute_status_304")
		delete(additionalProperties, "compute_status_400")
		delete(additionalProperties, "compute_status_401")
		delete(additionalProperties, "compute_status_403")
		delete(additionalProperties, "compute_status_404")
		delete(additionalProperties, "compute_status_416")
		delete(additionalProperties, "compute_status_429")
		delete(additionalProperties, "compute_status_500")
		delete(additionalProperties, "compute_status_501")
		delete(additionalProperties, "compute_status_502")
		delete(additionalProperties, "compute_status_503")
		delete(additionalProperties, "compute_status_504")
		delete(additionalProperties, "compute_status_505")
		delete(additionalProperties, "compute_latency_0_to_1ms")
		delete(additionalProperties, "compute_latency_1_to_5ms")
		delete(additionalProperties, "compute_latency_5_to_10ms")
		delete(additionalProperties, "compute_latency_10_to_50ms")
		delete(additionalProperties, "compute_latency_50_to_100ms")
		delete(additionalProperties, "compute_latency_100_to_250ms")
		delete(additionalProperties, "compute_latency_250_to_500ms")
		delete(additionalProperties, "compute_latency_500_to_1000ms")
		delete(additionalProperties, "compute_latency_1000_to_5000ms")
		delete(additionalProperties, "compute_latency_5000_to_10000ms")
		delete(additionalProperties, "compute_latency_10000_to_60000ms")
		delete(additionalProperties, "compute_latency_60000ms")
		delete(additionalProperties, "all_responses")
		delete(additionalProperties, "all_resp_header_bytes")
		delete(additionalProperties, "all_resp_body_bytes")
		delete(additionalProperties, "all_status_1xx")
		delete(additionalProperties, "all_status_2xx")
		delete(additionalProperties, "all_status_3xx")
		delete(additionalProperties, "all_status_4xx")
		delete(additionalProperties, "all_status_5xx")
		delete(additionalProperties, "all_status_200")
		delete(additionalProperties, "all_status_204")
		delete(additionalProperties, "all_status_206")
		delete(additionalProperties, "all_status_301")
		delete(additionalProperties, "all_status_302")
		delete(additionalProperties, "all_status_304")
		delete(additionalProperties, "all_status_400")
		delete(additionalProperties, "all_status_401")
		delete(additionalProperties, "all_status_403")
		delete(additionalProperties, "all_status_404")
		delete(additionalProperties, "all_status_416")
		delete(additionalProperties, "all_status_429")
		delete(additionalProperties, "all_status_500")
		delete(additionalProperties, "all_status_501")
		delete(additionalProperties, "all_status_502")
		delete(additionalProperties, "all_status_503")
		delete(additionalProperties, "all_status_504")
		delete(additionalProperties, "all_status_505")
		delete(additionalProperties, "all_latency_0_to_1ms")
		delete(additionalProperties, "all_latency_1_to_5ms")
		delete(additionalProperties, "all_latency_5_to_10ms")
		delete(additionalProperties, "all_latency_10_to_50ms")
		delete(additionalProperties, "all_latency_50_to_100ms")
		delete(additionalProperties, "all_latency_100_to_250ms")
		delete(additionalProperties, "all_latency_250_to_500ms")
		delete(additionalProperties, "all_latency_500_to_1000ms")
		delete(additionalProperties, "all_latency_1000_to_5000ms")
		delete(additionalProperties, "all_latency_5000_to_10000ms")
		delete(additionalProperties, "all_latency_10000_to_60000ms")
		delete(additionalProperties, "all_latency_60000ms")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableOriginInspectorMeasurements is a helper abstraction for handling nullable origininspectormeasurements types. 
type NullableOriginInspectorMeasurements struct {
	value *OriginInspectorMeasurements
	isSet bool
}

// Get returns the value.
func (v NullableOriginInspectorMeasurements) Get() *OriginInspectorMeasurements {
	return v.value
}

// Set modifies the value.
func (v *NullableOriginInspectorMeasurements) Set(val *OriginInspectorMeasurements) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableOriginInspectorMeasurements) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableOriginInspectorMeasurements) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableOriginInspectorMeasurements returns a pointer to a new instance of NullableOriginInspectorMeasurements.
func NewNullableOriginInspectorMeasurements(val *OriginInspectorMeasurements) *NullableOriginInspectorMeasurements {
	return &NullableOriginInspectorMeasurements{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableOriginInspectorMeasurements) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableOriginInspectorMeasurements) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
