// Package fastly is an API client library for interacting with most facets of the Fastly API.
package fastly

/*
Fastly API

Via the Fastly API you can perform any of the operations that are possible within the management console,  including creating services, domains, and backends, configuring rules or uploading your own application code, as well as account operations such as user administration and billing reports. The API is organized into collections of endpoints that allow manipulation of objects related to Fastly services and accounts. For the most accurate and up-to-date API reference content, visit our [Developer Hub](https://www.fastly.com/documentation/reference/api/)

API version: 1.0.0
Contact: oss@fastly.com
*/

// This code is auto-generated; DO NOT EDIT.

import (
	"encoding/json"
)

// DomainInspectorMeasurements A measurements object provides a count of the total number of `requests` received by Fastly for your domain in the reported time period and for the relevant POP, as specified in the [entry](#entry-data-model). It also includes the number of responses for specific HTTP response status codes and for status code ranges. This dataset is sparse: only the keys with non-zero values will be included in the record. Where a specific status code does not have a field in this model (e.g., `429 Too Many Requests`), any responses with that code will be counted as part of the range count (`4xx` in this case) but will not be separately identified in the data.
type DomainInspectorMeasurements struct {
	// Number of requests sent by end users to Fastly.
	EdgeRequests *int64 `json:"edge_requests,omitempty"`
	// Total header bytes delivered from Fastly to the end user.
	EdgeRespHeaderBytes *int64 `json:"edge_resp_header_bytes,omitempty"`
	// Total body bytes delivered from Fastly to the end user.
	EdgeRespBodyBytes *int64 `json:"edge_resp_body_bytes,omitempty"`
	// Number of 1xx \"Informational\" category status codes delivered.
	Status1xx *int64 `json:"status_1xx,omitempty"`
	// Number of 2xx \"Success\" status codes delivered.
	Status2xx *int64 `json:"status_2xx,omitempty"`
	// Number of 3xx \"Redirection\" codes delivered.
	Status3xx *int64 `json:"status_3xx,omitempty"`
	// Number of 4xx \"Client Error\" codes delivered.
	Status4xx *int64 `json:"status_4xx,omitempty"`
	// Number of 5xx \"Server Error\" codes delivered.
	Status5xx *int64 `json:"status_5xx,omitempty"`
	// Number of responses delivered with status code 200 (Success).
	Status200 *int64 `json:"status_200,omitempty"`
	// Number of responses delivered with status code 204 (No Content).
	Status204 *int64 `json:"status_204,omitempty"`
	// Number of responses delivered with status code 206 (Partial Content).
	Status206 *int64 `json:"status_206,omitempty"`
	// Number of responses delivered with status code 301 (Moved Permanently).
	Status301 *int64 `json:"status_301,omitempty"`
	// Number of responses delivered with status code 302 (Found).
	Status302 *int64 `json:"status_302,omitempty"`
	// Number of responses delivered with status code 304 (Not Modified).
	Status304 *int64 `json:"status_304,omitempty"`
	// Number of responses delivered with status code 400 (Bad Request).
	Status400 *int64 `json:"status_400,omitempty"`
	// Number of responses delivered with status code 401 (Unauthorized).
	Status401 *int64 `json:"status_401,omitempty"`
	// Number of responses delivered with status code 403 (Forbidden).
	Status403 *int64 `json:"status_403,omitempty"`
	// Number of responses delivered with status code 404 (Not Found).
	Status404 *int64 `json:"status_404,omitempty"`
	// Number of responses delivered with status code 416 (Range Not Satisfiable).
	Status416 *int64 `json:"status_416,omitempty"`
	// Number of responses delivered with status code 429 (Too Many Requests).
	Status429 *int64 `json:"status_429,omitempty"`
	// Number of responses delivered with status code 500 (Internal Server Error).
	Status500 *int64 `json:"status_500,omitempty"`
	// Number of responses delivered with status code 501 (Not Implemented).
	Status501 *int64 `json:"status_501,omitempty"`
	// Number of responses delivered with status code 502 (Bad Gateway).
	Status502 *int64 `json:"status_502,omitempty"`
	// Number of responses delivered with status code 503 (Service Unavailable).
	Status503 *int64 `json:"status_503,omitempty"`
	// Number of responses delivered with status code 504 (Gateway Timeout).
	Status504 *int64 `json:"status_504,omitempty"`
	// Number of responses delivered with status code 505 (HTTP Version Not Supported).
	Status505 *int64 `json:"status_505,omitempty"`
	// Number of responses delivered with status code 530.
	Status530 *int64 `json:"status_530,omitempty"`
	// Number of requests processed.
	Requests *int64 `json:"requests,omitempty"`
	// Total header bytes delivered.
	RespHeaderBytes *int64 `json:"resp_header_bytes,omitempty"`
	// Total body bytes delivered.
	RespBodyBytes *int64 `json:"resp_body_bytes,omitempty"`
	// Total header bytes sent to origin.
	BereqHeaderBytes *int64 `json:"bereq_header_bytes,omitempty"`
	// Total body bytes sent to origin.
	BereqBodyBytes *int64 `json:"bereq_body_bytes,omitempty"`
	// Number of requests sent by end users to Fastly that resulted in a hit at the edge.
	EdgeHitRequests *int64 `json:"edge_hit_requests,omitempty"`
	// Number of requests sent by end users to Fastly that resulted in a miss at the edge.
	EdgeMissRequests *int64 `json:"edge_miss_requests,omitempty"`
	// Number of requests sent to origin.
	OriginFetches *int64 `json:"origin_fetches,omitempty"`
	// Total header bytes received from origin.
	OriginFetchRespHeaderBytes *int64 `json:"origin_fetch_resp_header_bytes,omitempty"`
	// Total body bytes received from origin.
	OriginFetchRespBodyBytes *int64 `json:"origin_fetch_resp_body_bytes,omitempty"`
	// Total bytes delivered (`resp_header_bytes` + `resp_body_bytes` + `bereq_header_bytes` + `bereq_body_bytes`).
	Bandwidth *int64 `json:"bandwidth,omitempty"`
	// Ratio of cache hits to cache misses at the edge, between 0 and 1 (`edge_hit_requests` / (`edge_hit_requests` + `edge_miss_requests`)).
	EdgeHitRatio *float32 `json:"edge_hit_ratio,omitempty"`
	// Origin Offload measures the ratio of bytes served to end users that were cached by Fastly, over the bytes served to end users, between 0 and 1. ((`edge_resp_body_bytes` + `edge_resp_header_bytes`) - (`origin_fetch_resp_body_bytes` + `origin_fetch_resp_header_bytes`)) / (`edge_resp_body_bytes` + `edge_resp_header_bytes`). Previously, Origin Offload used a different formula. [Learn more](https://www.fastly.com/documentation/reference/changes/2024/06/add-origin_offload-metric).
	OriginOffload *float32 `json:"origin_offload,omitempty"`
	// Number of responses received from origin with status code 200 (Success).
	OriginStatus200 *int64 `json:"origin_status_200,omitempty"`
	// Number of responses received from origin with status code 204 (No Content).
	OriginStatus204 *int64 `json:"origin_status_204,omitempty"`
	// Number of responses received from origin with status code 206 (Partial Content).
	OriginStatus206 *int64 `json:"origin_status_206,omitempty"`
	// Number of responses received from origin with status code 301 (Moved Permanently).
	OriginStatus301 *int64 `json:"origin_status_301,omitempty"`
	// Number of responses received from origin with status code 302 (Found).
	OriginStatus302 *int64 `json:"origin_status_302,omitempty"`
	// Number of responses received from origin with status code 304 (Not Modified).
	OriginStatus304 *int64 `json:"origin_status_304,omitempty"`
	// Number of responses received from origin with status code 400 (Bad Request).
	OriginStatus400 *int64 `json:"origin_status_400,omitempty"`
	// Number of responses received from origin with status code 401 (Unauthorized).
	OriginStatus401 *int64 `json:"origin_status_401,omitempty"`
	// Number of responses received from origin with status code 403 (Forbidden).
	OriginStatus403 *int64 `json:"origin_status_403,omitempty"`
	// Number of responses received from origin with status code 404 (Not Found).
	OriginStatus404 *int64 `json:"origin_status_404,omitempty"`
	// Number of responses received from origin with status code 416 (Range Not Satisfiable).
	OriginStatus416 *int64 `json:"origin_status_416,omitempty"`
	// Number of responses received from origin with status code 429 (Too Many Requests).
	OriginStatus429 *int64 `json:"origin_status_429,omitempty"`
	// Number of responses received from origin with status code 500 (Internal Server Error).
	OriginStatus500 *int64 `json:"origin_status_500,omitempty"`
	// Number of responses received from origin with status code 501 (Not Implemented).
	OriginStatus501 *int64 `json:"origin_status_501,omitempty"`
	// Number of responses received from origin with status code 502 (Bad Gateway).
	OriginStatus502 *int64 `json:"origin_status_502,omitempty"`
	// Number of responses received from origin with status code 503 (Service Unavailable).
	OriginStatus503 *int64 `json:"origin_status_503,omitempty"`
	// Number of responses received from origin with status code 504 (Gateway Timeout).
	OriginStatus504 *int64 `json:"origin_status_504,omitempty"`
	// Number of responses received from origin with status code 505 (HTTP Version Not Supported).
	OriginStatus505 *int64 `json:"origin_status_505,omitempty"`
	// Number of responses received from origin with status code 530.
	OriginStatus530 *int64 `json:"origin_status_530,omitempty"`
	// Number of \"Informational\" category status codes received from origin.
	OriginStatus1xx *int64 `json:"origin_status_1xx,omitempty"`
	// Number of \"Success\" status codes received from origin.
	OriginStatus2xx *int64 `json:"origin_status_2xx,omitempty"`
	// Number of \"Redirection\" codes received from origin.
	OriginStatus3xx *int64 `json:"origin_status_3xx,omitempty"`
	// Number of \"Client Error\" codes received from origin.
	OriginStatus4xx *int64 `json:"origin_status_4xx,omitempty"`
	// Number of \"Server Error\" codes received from origin.
	OriginStatus5xx      *int64 `json:"origin_status_5xx,omitempty"`
	AdditionalProperties map[string]any
}

type _DomainInspectorMeasurements DomainInspectorMeasurements

// NewDomainInspectorMeasurements instantiates a new DomainInspectorMeasurements object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomainInspectorMeasurements() *DomainInspectorMeasurements {
	this := DomainInspectorMeasurements{}
	return &this
}

// NewDomainInspectorMeasurementsWithDefaults instantiates a new DomainInspectorMeasurements object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainInspectorMeasurementsWithDefaults() *DomainInspectorMeasurements {
	this := DomainInspectorMeasurements{}
	return &this
}

// GetEdgeRequests returns the EdgeRequests field value if set, zero value otherwise.
func (o *DomainInspectorMeasurements) GetEdgeRequests() int64 {
	if o == nil || o.EdgeRequests == nil {
		var ret int64
		return ret
	}
	return *o.EdgeRequests
}

// GetEdgeRequestsOk returns a tuple with the EdgeRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainInspectorMeasurements) GetEdgeRequestsOk() (*int64, bool) {
	if o == nil || o.EdgeRequests == nil {
		return nil, false
	}
	return o.EdgeRequests, true
}

// HasEdgeRequests returns a boolean if a field has been set.
func (o *DomainInspectorMeasurements) HasEdgeRequests() bool {
	if o != nil && o.EdgeRequests != nil {
		return true
	}

	return false
}

// SetEdgeRequests gets a reference to the given int64 and assigns it to the EdgeRequests field.
func (o *DomainInspectorMeasurements) SetEdgeRequests(v int64) {
	o.EdgeRequests = &v
}

// GetEdgeRespHeaderBytes returns the EdgeRespHeaderBytes field value if set, zero value otherwise.
func (o *DomainInspectorMeasurements) GetEdgeRespHeaderBytes() int64 {
	if o == nil || o.EdgeRespHeaderBytes == nil {
		var ret int64
		return ret
	}
	return *o.EdgeRespHeaderBytes
}

// GetEdgeRespHeaderBytesOk returns a tuple with the EdgeRespHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainInspectorMeasurements) GetEdgeRespHeaderBytesOk() (*int64, bool) {
	if o == nil || o.EdgeRespHeaderBytes == nil {
		return nil, false
	}
	return o.EdgeRespHeaderBytes, true
}

// HasEdgeRespHeaderBytes returns a boolean if a field has been set.
func (o *DomainInspectorMeasurements) HasEdgeRespHeaderBytes() bool {
	if o != nil && o.EdgeRespHeaderBytes != nil {
		return true
	}

	return false
}

// SetEdgeRespHeaderBytes gets a reference to the given int64 and assigns it to the EdgeRespHeaderBytes field.
func (o *DomainInspectorMeasurements) SetEdgeRespHeaderBytes(v int64) {
	o.EdgeRespHeaderBytes = &v
}

// GetEdgeRespBodyBytes returns the EdgeRespBodyBytes field value if set, zero value otherwise.
func (o *DomainInspectorMeasurements) GetEdgeRespBodyBytes() int64 {
	if o == nil || o.EdgeRespBodyBytes == nil {
		var ret int64
		return ret
	}
	return *o.EdgeRespBodyBytes
}

// GetEdgeRespBodyBytesOk returns a tuple with the EdgeRespBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainInspectorMeasurements) GetEdgeRespBodyBytesOk() (*int64, bool) {
	if o == nil || o.EdgeRespBodyBytes == nil {
		return nil, false
	}
	return o.EdgeRespBodyBytes, true
}

// HasEdgeRespBodyBytes returns a boolean if a field has been set.
func (o *DomainInspectorMeasurements) HasEdgeRespBodyBytes() bool {
	if o != nil && o.EdgeRespBodyBytes != nil {
		return true
	}

	return false
}

// SetEdgeRespBodyBytes gets a reference to the given int64 and assigns it to the EdgeRespBodyBytes field.
func (o *DomainInspectorMeasurements) SetEdgeRespBodyBytes(v int64) {
	o.EdgeRespBodyBytes = &v
}

// GetStatus1xx returns the Status1xx field value if set, zero value otherwise.
func (o *DomainInspectorMeasurements) GetStatus1xx() int64 {
	if o == nil || o.Status1xx == nil {
		var ret int64
		return ret
	}
	return *o.Status1xx
}

// GetStatus1xxOk returns a tuple with the Status1xx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainInspectorMeasurements) GetStatus1xxOk() (*int64, bool) {
	if o == nil || o.Status1xx == nil {
		return nil, false
	}
	return o.Status1xx, true
}

// HasStatus1xx returns a boolean if a field has been set.
func (o *DomainInspectorMeasurements) HasStatus1xx() bool {
	if o != nil && o.Status1xx != nil {
		return true
	}

	return false
}

// SetStatus1xx gets a reference to the given int64 and assigns it to the Status1xx field.
func (o *DomainInspectorMeasurements) SetStatus1xx(v int64) {
	o.Status1xx = &v
}

// GetStatus2xx returns the Status2xx field value if set, zero value otherwise.
func (o *DomainInspectorMeasurements) GetStatus2xx() int64 {
	if o == nil || o.Status2xx == nil {
		var ret int64
		return ret
	}
	return *o.Status2xx
}

// GetStatus2xxOk returns a tuple with the Status2xx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainInspectorMeasurements) GetStatus2xxOk() (*int64, bool) {
	if o == nil || o.Status2xx == nil {
		return nil, false
	}
	return o.Status2xx, true
}

// HasStatus2xx returns a boolean if a field has been set.
func (o *DomainInspectorMeasurements) HasStatus2xx() bool {
	if o != nil && o.Status2xx != nil {
		return true
	}

	return false
}

// SetStatus2xx gets a reference to the given int64 and assigns it to the Status2xx field.
func (o *DomainInspectorMeasurements) SetStatus2xx(v int64) {
	o.Status2xx = &v
}

// GetStatus3xx returns the Status3xx field value if set, zero value otherwise.
func (o *DomainInspectorMeasurements) GetStatus3xx() int64 {
	if o == nil || o.Status3xx == nil {
		var ret int64
		return ret
	}
	return *o.Status3xx
}

// GetStatus3xxOk returns a tuple with the Status3xx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainInspectorMeasurements) GetStatus3xxOk() (*int64, bool) {
	if o == nil || o.Status3xx == nil {
		return nil, false
	}
	return o.Status3xx, true
}

// HasStatus3xx returns a boolean if a field has been set.
func (o *DomainInspectorMeasurements) HasStatus3xx() bool {
	if o != nil && o.Status3xx != nil {
		return true
	}

	return false
}

// SetStatus3xx gets a reference to the given int64 and assigns it to the Status3xx field.
func (o *DomainInspectorMeasurements) SetStatus3xx(v int64) {
	o.Status3xx = &v
}

// GetStatus4xx returns the Status4xx field value if set, zero value otherwise.
func (o *DomainInspectorMeasurements) GetStatus4xx() int64 {
	if o == nil || o.Status4xx == nil {
		var ret int64
		return ret
	}
	return *o.Status4xx
}

// GetStatus4xxOk returns a tuple with the Status4xx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainInspectorMeasurements) GetStatus4xxOk() (*int64, bool) {
	if o == nil || o.Status4xx == nil {
		return nil, false
	}
	return o.Status4xx, true
}

// HasStatus4xx returns a boolean if a field has been set.
func (o *DomainInspectorMeasurements) HasStatus4xx() bool {
	if o != nil && o.Status4xx != nil {
		return true
	}

	return false
}

// SetStatus4xx gets a reference to the given int64 and assigns it to the Status4xx field.
func (o *DomainInspectorMeasurements) SetStatus4xx(v int64) {
	o.Status4xx = &v
}

// GetStatus5xx returns the Status5xx field value if set, zero value otherwise.
func (o *DomainInspectorMeasurements) GetStatus5xx() int64 {
	if o == nil || o.Status5xx == nil {
		var ret int64
		return ret
	}
	return *o.Status5xx
}

// GetStatus5xxOk returns a tuple with the Status5xx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainInspectorMeasurements) GetStatus5xxOk() (*int64, bool) {
	if o == nil || o.Status5xx == nil {
		return nil, false
	}
	return o.Status5xx, true
}

// HasStatus5xx returns a boolean if a field has been set.
func (o *DomainInspectorMeasurements) HasStatus5xx() bool {
	if o != nil && o.Status5xx != nil {
		return true
	}

	return false
}

// SetStatus5xx gets a reference to the given int64 and assigns it to the Status5xx field.
func (o *DomainInspectorMeasurements) SetStatus5xx(v int64) {
	o.Status5xx = &v
}

// GetStatus200 returns the Status200 field value if set, zero value otherwise.
func (o *DomainInspectorMeasurements) GetStatus200() int64 {
	if o == nil || o.Status200 == nil {
		var ret int64
		return ret
	}
	return *o.Status200
}

// GetStatus200Ok returns a tuple with the Status200 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainInspectorMeasurements) GetStatus200Ok() (*int64, bool) {
	if o == nil || o.Status200 == nil {
		return nil, false
	}
	return o.Status200, true
}

// HasStatus200 returns a boolean if a field has been set.
func (o *DomainInspectorMeasurements) HasStatus200() bool {
	if o != nil && o.Status200 != nil {
		return true
	}

	return false
}

// SetStatus200 gets a reference to the given int64 and assigns it to the Status200 field.
func (o *DomainInspectorMeasurements) SetStatus200(v int64) {
	o.Status200 = &v
}

// GetStatus204 returns the Status204 field value if set, zero value otherwise.
func (o *DomainInspectorMeasurements) GetStatus204() int64 {
	if o == nil || o.Status204 == nil {
		var ret int64
		return ret
	}
	return *o.Status204
}

// GetStatus204Ok returns a tuple with the Status204 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainInspectorMeasurements) GetStatus204Ok() (*int64, bool) {
	if o == nil || o.Status204 == nil {
		return nil, false
	}
	return o.Status204, true
}

// HasStatus204 returns a boolean if a field has been set.
func (o *DomainInspectorMeasurements) HasStatus204() bool {
	if o != nil && o.Status204 != nil {
		return true
	}

	return false
}

// SetStatus204 gets a reference to the given int64 and assigns it to the Status204 field.
func (o *DomainInspectorMeasurements) SetStatus204(v int64) {
	o.Status204 = &v
}

// GetStatus206 returns the Status206 field value if set, zero value otherwise.
func (o *DomainInspectorMeasurements) GetStatus206() int64 {
	if o == nil || o.Status206 == nil {
		var ret int64
		return ret
	}
	return *o.Status206
}

// GetStatus206Ok returns a tuple with the Status206 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainInspectorMeasurements) GetStatus206Ok() (*int64, bool) {
	if o == nil || o.Status206 == nil {
		return nil, false
	}
	return o.Status206, true
}

// HasStatus206 returns a boolean if a field has been set.
func (o *DomainInspectorMeasurements) HasStatus206() bool {
	if o != nil && o.Status206 != nil {
		return true
	}

	return false
}

// SetStatus206 gets a reference to the given int64 and assigns it to the Status206 field.
func (o *DomainInspectorMeasurements) SetStatus206(v int64) {
	o.Status206 = &v
}

// GetStatus301 returns the Status301 field value if set, zero value otherwise.
func (o *DomainInspectorMeasurements) GetStatus301() int64 {
	if o == nil || o.Status301 == nil {
		var ret int64
		return ret
	}
	return *o.Status301
}

// GetStatus301Ok returns a tuple with the Status301 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainInspectorMeasurements) GetStatus301Ok() (*int64, bool) {
	if o == nil || o.Status301 == nil {
		return nil, false
	}
	return o.Status301, true
}

// HasStatus301 returns a boolean if a field has been set.
func (o *DomainInspectorMeasurements) HasStatus301() bool {
	if o != nil && o.Status301 != nil {
		return true
	}

	return false
}

// SetStatus301 gets a reference to the given int64 and assigns it to the Status301 field.
func (o *DomainInspectorMeasurements) SetStatus301(v int64) {
	o.Status301 = &v
}

// GetStatus302 returns the Status302 field value if set, zero value otherwise.
func (o *DomainInspectorMeasurements) GetStatus302() int64 {
	if o == nil || o.Status302 == nil {
		var ret int64
		return ret
	}
	return *o.Status302
}

// GetStatus302Ok returns a tuple with the Status302 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainInspectorMeasurements) GetStatus302Ok() (*int64, bool) {
	if o == nil || o.Status302 == nil {
		return nil, false
	}
	return o.Status302, true
}

// HasStatus302 returns a boolean if a field has been set.
func (o *DomainInspectorMeasurements) HasStatus302() bool {
	if o != nil && o.Status302 != nil {
		return true
	}

	return false
}

// SetStatus302 gets a reference to the given int64 and assigns it to the Status302 field.
func (o *DomainInspectorMeasurements) SetStatus302(v int64) {
	o.Status302 = &v
}

// GetStatus304 returns the Status304 field value if set, zero value otherwise.
func (o *DomainInspectorMeasurements) GetStatus304() int64 {
	if o == nil || o.Status304 == nil {
		var ret int64
		return ret
	}
	return *o.Status304
}

// GetStatus304Ok returns a tuple with the Status304 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainInspectorMeasurements) GetStatus304Ok() (*int64, bool) {
	if o == nil || o.Status304 == nil {
		return nil, false
	}
	return o.Status304, true
}

// HasStatus304 returns a boolean if a field has been set.
func (o *DomainInspectorMeasurements) HasStatus304() bool {
	if o != nil && o.Status304 != nil {
		return true
	}

	return false
}

// SetStatus304 gets a reference to the given int64 and assigns it to the Status304 field.
func (o *DomainInspectorMeasurements) SetStatus304(v int64) {
	o.Status304 = &v
}

// GetStatus400 returns the Status400 field value if set, zero value otherwise.
func (o *DomainInspectorMeasurements) GetStatus400() int64 {
	if o == nil || o.Status400 == nil {
		var ret int64
		return ret
	}
	return *o.Status400
}

// GetStatus400Ok returns a tuple with the Status400 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainInspectorMeasurements) GetStatus400Ok() (*int64, bool) {
	if o == nil || o.Status400 == nil {
		return nil, false
	}
	return o.Status400, true
}

// HasStatus400 returns a boolean if a field has been set.
func (o *DomainInspectorMeasurements) HasStatus400() bool {
	if o != nil && o.Status400 != nil {
		return true
	}

	return false
}

// SetStatus400 gets a reference to the given int64 and assigns it to the Status400 field.
func (o *DomainInspectorMeasurements) SetStatus400(v int64) {
	o.Status400 = &v
}

// GetStatus401 returns the Status401 field value if set, zero value otherwise.
func (o *DomainInspectorMeasurements) GetStatus401() int64 {
	if o == nil || o.Status401 == nil {
		var ret int64
		return ret
	}
	return *o.Status401
}

// GetStatus401Ok returns a tuple with the Status401 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainInspectorMeasurements) GetStatus401Ok() (*int64, bool) {
	if o == nil || o.Status401 == nil {
		return nil, false
	}
	return o.Status401, true
}

// HasStatus401 returns a boolean if a field has been set.
func (o *DomainInspectorMeasurements) HasStatus401() bool {
	if o != nil && o.Status401 != nil {
		return true
	}

	return false
}

// SetStatus401 gets a reference to the given int64 and assigns it to the Status401 field.
func (o *DomainInspectorMeasurements) SetStatus401(v int64) {
	o.Status401 = &v
}

// GetStatus403 returns the Status403 field value if set, zero value otherwise.
func (o *DomainInspectorMeasurements) GetStatus403() int64 {
	if o == nil || o.Status403 == nil {
		var ret int64
		return ret
	}
	return *o.Status403
}

// GetStatus403Ok returns a tuple with the Status403 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainInspectorMeasurements) GetStatus403Ok() (*int64, bool) {
	if o == nil || o.Status403 == nil {
		return nil, false
	}
	return o.Status403, true
}

// HasStatus403 returns a boolean if a field has been set.
func (o *DomainInspectorMeasurements) HasStatus403() bool {
	if o != nil && o.Status403 != nil {
		return true
	}

	return false
}

// SetStatus403 gets a reference to the given int64 and assigns it to the Status403 field.
func (o *DomainInspectorMeasurements) SetStatus403(v int64) {
	o.Status403 = &v
}

// GetStatus404 returns the Status404 field value if set, zero value otherwise.
func (o *DomainInspectorMeasurements) GetStatus404() int64 {
	if o == nil || o.Status404 == nil {
		var ret int64
		return ret
	}
	return *o.Status404
}

// GetStatus404Ok returns a tuple with the Status404 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainInspectorMeasurements) GetStatus404Ok() (*int64, bool) {
	if o == nil || o.Status404 == nil {
		return nil, false
	}
	return o.Status404, true
}

// HasStatus404 returns a boolean if a field has been set.
func (o *DomainInspectorMeasurements) HasStatus404() bool {
	if o != nil && o.Status404 != nil {
		return true
	}

	return false
}

// SetStatus404 gets a reference to the given int64 and assigns it to the Status404 field.
func (o *DomainInspectorMeasurements) SetStatus404(v int64) {
	o.Status404 = &v
}

// GetStatus416 returns the Status416 field value if set, zero value otherwise.
func (o *DomainInspectorMeasurements) GetStatus416() int64 {
	if o == nil || o.Status416 == nil {
		var ret int64
		return ret
	}
	return *o.Status416
}

// GetStatus416Ok returns a tuple with the Status416 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainInspectorMeasurements) GetStatus416Ok() (*int64, bool) {
	if o == nil || o.Status416 == nil {
		return nil, false
	}
	return o.Status416, true
}

// HasStatus416 returns a boolean if a field has been set.
func (o *DomainInspectorMeasurements) HasStatus416() bool {
	if o != nil && o.Status416 != nil {
		return true
	}

	return false
}

// SetStatus416 gets a reference to the given int64 and assigns it to the Status416 field.
func (o *DomainInspectorMeasurements) SetStatus416(v int64) {
	o.Status416 = &v
}

// GetStatus429 returns the Status429 field value if set, zero value otherwise.
func (o *DomainInspectorMeasurements) GetStatus429() int64 {
	if o == nil || o.Status429 == nil {
		var ret int64
		return ret
	}
	return *o.Status429
}

// GetStatus429Ok returns a tuple with the Status429 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainInspectorMeasurements) GetStatus429Ok() (*int64, bool) {
	if o == nil || o.Status429 == nil {
		return nil, false
	}
	return o.Status429, true
}

// HasStatus429 returns a boolean if a field has been set.
func (o *DomainInspectorMeasurements) HasStatus429() bool {
	if o != nil && o.Status429 != nil {
		return true
	}

	return false
}

// SetStatus429 gets a reference to the given int64 and assigns it to the Status429 field.
func (o *DomainInspectorMeasurements) SetStatus429(v int64) {
	o.Status429 = &v
}

// GetStatus500 returns the Status500 field value if set, zero value otherwise.
func (o *DomainInspectorMeasurements) GetStatus500() int64 {
	if o == nil || o.Status500 == nil {
		var ret int64
		return ret
	}
	return *o.Status500
}

// GetStatus500Ok returns a tuple with the Status500 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainInspectorMeasurements) GetStatus500Ok() (*int64, bool) {
	if o == nil || o.Status500 == nil {
		return nil, false
	}
	return o.Status500, true
}

// HasStatus500 returns a boolean if a field has been set.
func (o *DomainInspectorMeasurements) HasStatus500() bool {
	if o != nil && o.Status500 != nil {
		return true
	}

	return false
}

// SetStatus500 gets a reference to the given int64 and assigns it to the Status500 field.
func (o *DomainInspectorMeasurements) SetStatus500(v int64) {
	o.Status500 = &v
}

// GetStatus501 returns the Status501 field value if set, zero value otherwise.
func (o *DomainInspectorMeasurements) GetStatus501() int64 {
	if o == nil || o.Status501 == nil {
		var ret int64
		return ret
	}
	return *o.Status501
}

// GetStatus501Ok returns a tuple with the Status501 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainInspectorMeasurements) GetStatus501Ok() (*int64, bool) {
	if o == nil || o.Status501 == nil {
		return nil, false
	}
	return o.Status501, true
}

// HasStatus501 returns a boolean if a field has been set.
func (o *DomainInspectorMeasurements) HasStatus501() bool {
	if o != nil && o.Status501 != nil {
		return true
	}

	return false
}

// SetStatus501 gets a reference to the given int64 and assigns it to the Status501 field.
func (o *DomainInspectorMeasurements) SetStatus501(v int64) {
	o.Status501 = &v
}

// GetStatus502 returns the Status502 field value if set, zero value otherwise.
func (o *DomainInspectorMeasurements) GetStatus502() int64 {
	if o == nil || o.Status502 == nil {
		var ret int64
		return ret
	}
	return *o.Status502
}

// GetStatus502Ok returns a tuple with the Status502 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainInspectorMeasurements) GetStatus502Ok() (*int64, bool) {
	if o == nil || o.Status502 == nil {
		return nil, false
	}
	return o.Status502, true
}

// HasStatus502 returns a boolean if a field has been set.
func (o *DomainInspectorMeasurements) HasStatus502() bool {
	if o != nil && o.Status502 != nil {
		return true
	}

	return false
}

// SetStatus502 gets a reference to the given int64 and assigns it to the Status502 field.
func (o *DomainInspectorMeasurements) SetStatus502(v int64) {
	o.Status502 = &v
}

// GetStatus503 returns the Status503 field value if set, zero value otherwise.
func (o *DomainInspectorMeasurements) GetStatus503() int64 {
	if o == nil || o.Status503 == nil {
		var ret int64
		return ret
	}
	return *o.Status503
}

// GetStatus503Ok returns a tuple with the Status503 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainInspectorMeasurements) GetStatus503Ok() (*int64, bool) {
	if o == nil || o.Status503 == nil {
		return nil, false
	}
	return o.Status503, true
}

// HasStatus503 returns a boolean if a field has been set.
func (o *DomainInspectorMeasurements) HasStatus503() bool {
	if o != nil && o.Status503 != nil {
		return true
	}

	return false
}

// SetStatus503 gets a reference to the given int64 and assigns it to the Status503 field.
func (o *DomainInspectorMeasurements) SetStatus503(v int64) {
	o.Status503 = &v
}

// GetStatus504 returns the Status504 field value if set, zero value otherwise.
func (o *DomainInspectorMeasurements) GetStatus504() int64 {
	if o == nil || o.Status504 == nil {
		var ret int64
		return ret
	}
	return *o.Status504
}

// GetStatus504Ok returns a tuple with the Status504 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainInspectorMeasurements) GetStatus504Ok() (*int64, bool) {
	if o == nil || o.Status504 == nil {
		return nil, false
	}
	return o.Status504, true
}

// HasStatus504 returns a boolean if a field has been set.
func (o *DomainInspectorMeasurements) HasStatus504() bool {
	if o != nil && o.Status504 != nil {
		return true
	}

	return false
}

// SetStatus504 gets a reference to the given int64 and assigns it to the Status504 field.
func (o *DomainInspectorMeasurements) SetStatus504(v int64) {
	o.Status504 = &v
}

// GetStatus505 returns the Status505 field value if set, zero value otherwise.
func (o *DomainInspectorMeasurements) GetStatus505() int64 {
	if o == nil || o.Status505 == nil {
		var ret int64
		return ret
	}
	return *o.Status505
}

// GetStatus505Ok returns a tuple with the Status505 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainInspectorMeasurements) GetStatus505Ok() (*int64, bool) {
	if o == nil || o.Status505 == nil {
		return nil, false
	}
	return o.Status505, true
}

// HasStatus505 returns a boolean if a field has been set.
func (o *DomainInspectorMeasurements) HasStatus505() bool {
	if o != nil && o.Status505 != nil {
		return true
	}

	return false
}

// SetStatus505 gets a reference to the given int64 and assigns it to the Status505 field.
func (o *DomainInspectorMeasurements) SetStatus505(v int64) {
	o.Status505 = &v
}

// GetStatus530 returns the Status530 field value if set, zero value otherwise.
func (o *DomainInspectorMeasurements) GetStatus530() int64 {
	if o == nil || o.Status530 == nil {
		var ret int64
		return ret
	}
	return *o.Status530
}

// GetStatus530Ok returns a tuple with the Status530 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainInspectorMeasurements) GetStatus530Ok() (*int64, bool) {
	if o == nil || o.Status530 == nil {
		return nil, false
	}
	return o.Status530, true
}

// HasStatus530 returns a boolean if a field has been set.
func (o *DomainInspectorMeasurements) HasStatus530() bool {
	if o != nil && o.Status530 != nil {
		return true
	}

	return false
}

// SetStatus530 gets a reference to the given int64 and assigns it to the Status530 field.
func (o *DomainInspectorMeasurements) SetStatus530(v int64) {
	o.Status530 = &v
}

// GetRequests returns the Requests field value if set, zero value otherwise.
func (o *DomainInspectorMeasurements) GetRequests() int64 {
	if o == nil || o.Requests == nil {
		var ret int64
		return ret
	}
	return *o.Requests
}

// GetRequestsOk returns a tuple with the Requests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainInspectorMeasurements) GetRequestsOk() (*int64, bool) {
	if o == nil || o.Requests == nil {
		return nil, false
	}
	return o.Requests, true
}

// HasRequests returns a boolean if a field has been set.
func (o *DomainInspectorMeasurements) HasRequests() bool {
	if o != nil && o.Requests != nil {
		return true
	}

	return false
}

// SetRequests gets a reference to the given int64 and assigns it to the Requests field.
func (o *DomainInspectorMeasurements) SetRequests(v int64) {
	o.Requests = &v
}

// GetRespHeaderBytes returns the RespHeaderBytes field value if set, zero value otherwise.
func (o *DomainInspectorMeasurements) GetRespHeaderBytes() int64 {
	if o == nil || o.RespHeaderBytes == nil {
		var ret int64
		return ret
	}
	return *o.RespHeaderBytes
}

// GetRespHeaderBytesOk returns a tuple with the RespHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainInspectorMeasurements) GetRespHeaderBytesOk() (*int64, bool) {
	if o == nil || o.RespHeaderBytes == nil {
		return nil, false
	}
	return o.RespHeaderBytes, true
}

// HasRespHeaderBytes returns a boolean if a field has been set.
func (o *DomainInspectorMeasurements) HasRespHeaderBytes() bool {
	if o != nil && o.RespHeaderBytes != nil {
		return true
	}

	return false
}

// SetRespHeaderBytes gets a reference to the given int64 and assigns it to the RespHeaderBytes field.
func (o *DomainInspectorMeasurements) SetRespHeaderBytes(v int64) {
	o.RespHeaderBytes = &v
}

// GetRespBodyBytes returns the RespBodyBytes field value if set, zero value otherwise.
func (o *DomainInspectorMeasurements) GetRespBodyBytes() int64 {
	if o == nil || o.RespBodyBytes == nil {
		var ret int64
		return ret
	}
	return *o.RespBodyBytes
}

// GetRespBodyBytesOk returns a tuple with the RespBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainInspectorMeasurements) GetRespBodyBytesOk() (*int64, bool) {
	if o == nil || o.RespBodyBytes == nil {
		return nil, false
	}
	return o.RespBodyBytes, true
}

// HasRespBodyBytes returns a boolean if a field has been set.
func (o *DomainInspectorMeasurements) HasRespBodyBytes() bool {
	if o != nil && o.RespBodyBytes != nil {
		return true
	}

	return false
}

// SetRespBodyBytes gets a reference to the given int64 and assigns it to the RespBodyBytes field.
func (o *DomainInspectorMeasurements) SetRespBodyBytes(v int64) {
	o.RespBodyBytes = &v
}

// GetBereqHeaderBytes returns the BereqHeaderBytes field value if set, zero value otherwise.
func (o *DomainInspectorMeasurements) GetBereqHeaderBytes() int64 {
	if o == nil || o.BereqHeaderBytes == nil {
		var ret int64
		return ret
	}
	return *o.BereqHeaderBytes
}

// GetBereqHeaderBytesOk returns a tuple with the BereqHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainInspectorMeasurements) GetBereqHeaderBytesOk() (*int64, bool) {
	if o == nil || o.BereqHeaderBytes == nil {
		return nil, false
	}
	return o.BereqHeaderBytes, true
}

// HasBereqHeaderBytes returns a boolean if a field has been set.
func (o *DomainInspectorMeasurements) HasBereqHeaderBytes() bool {
	if o != nil && o.BereqHeaderBytes != nil {
		return true
	}

	return false
}

// SetBereqHeaderBytes gets a reference to the given int64 and assigns it to the BereqHeaderBytes field.
func (o *DomainInspectorMeasurements) SetBereqHeaderBytes(v int64) {
	o.BereqHeaderBytes = &v
}

// GetBereqBodyBytes returns the BereqBodyBytes field value if set, zero value otherwise.
func (o *DomainInspectorMeasurements) GetBereqBodyBytes() int64 {
	if o == nil || o.BereqBodyBytes == nil {
		var ret int64
		return ret
	}
	return *o.BereqBodyBytes
}

// GetBereqBodyBytesOk returns a tuple with the BereqBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainInspectorMeasurements) GetBereqBodyBytesOk() (*int64, bool) {
	if o == nil || o.BereqBodyBytes == nil {
		return nil, false
	}
	return o.BereqBodyBytes, true
}

// HasBereqBodyBytes returns a boolean if a field has been set.
func (o *DomainInspectorMeasurements) HasBereqBodyBytes() bool {
	if o != nil && o.BereqBodyBytes != nil {
		return true
	}

	return false
}

// SetBereqBodyBytes gets a reference to the given int64 and assigns it to the BereqBodyBytes field.
func (o *DomainInspectorMeasurements) SetBereqBodyBytes(v int64) {
	o.BereqBodyBytes = &v
}

// GetEdgeHitRequests returns the EdgeHitRequests field value if set, zero value otherwise.
func (o *DomainInspectorMeasurements) GetEdgeHitRequests() int64 {
	if o == nil || o.EdgeHitRequests == nil {
		var ret int64
		return ret
	}
	return *o.EdgeHitRequests
}

// GetEdgeHitRequestsOk returns a tuple with the EdgeHitRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainInspectorMeasurements) GetEdgeHitRequestsOk() (*int64, bool) {
	if o == nil || o.EdgeHitRequests == nil {
		return nil, false
	}
	return o.EdgeHitRequests, true
}

// HasEdgeHitRequests returns a boolean if a field has been set.
func (o *DomainInspectorMeasurements) HasEdgeHitRequests() bool {
	if o != nil && o.EdgeHitRequests != nil {
		return true
	}

	return false
}

// SetEdgeHitRequests gets a reference to the given int64 and assigns it to the EdgeHitRequests field.
func (o *DomainInspectorMeasurements) SetEdgeHitRequests(v int64) {
	o.EdgeHitRequests = &v
}

// GetEdgeMissRequests returns the EdgeMissRequests field value if set, zero value otherwise.
func (o *DomainInspectorMeasurements) GetEdgeMissRequests() int64 {
	if o == nil || o.EdgeMissRequests == nil {
		var ret int64
		return ret
	}
	return *o.EdgeMissRequests
}

// GetEdgeMissRequestsOk returns a tuple with the EdgeMissRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainInspectorMeasurements) GetEdgeMissRequestsOk() (*int64, bool) {
	if o == nil || o.EdgeMissRequests == nil {
		return nil, false
	}
	return o.EdgeMissRequests, true
}

// HasEdgeMissRequests returns a boolean if a field has been set.
func (o *DomainInspectorMeasurements) HasEdgeMissRequests() bool {
	if o != nil && o.EdgeMissRequests != nil {
		return true
	}

	return false
}

// SetEdgeMissRequests gets a reference to the given int64 and assigns it to the EdgeMissRequests field.
func (o *DomainInspectorMeasurements) SetEdgeMissRequests(v int64) {
	o.EdgeMissRequests = &v
}

// GetOriginFetches returns the OriginFetches field value if set, zero value otherwise.
func (o *DomainInspectorMeasurements) GetOriginFetches() int64 {
	if o == nil || o.OriginFetches == nil {
		var ret int64
		return ret
	}
	return *o.OriginFetches
}

// GetOriginFetchesOk returns a tuple with the OriginFetches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainInspectorMeasurements) GetOriginFetchesOk() (*int64, bool) {
	if o == nil || o.OriginFetches == nil {
		return nil, false
	}
	return o.OriginFetches, true
}

// HasOriginFetches returns a boolean if a field has been set.
func (o *DomainInspectorMeasurements) HasOriginFetches() bool {
	if o != nil && o.OriginFetches != nil {
		return true
	}

	return false
}

// SetOriginFetches gets a reference to the given int64 and assigns it to the OriginFetches field.
func (o *DomainInspectorMeasurements) SetOriginFetches(v int64) {
	o.OriginFetches = &v
}

// GetOriginFetchRespHeaderBytes returns the OriginFetchRespHeaderBytes field value if set, zero value otherwise.
func (o *DomainInspectorMeasurements) GetOriginFetchRespHeaderBytes() int64 {
	if o == nil || o.OriginFetchRespHeaderBytes == nil {
		var ret int64
		return ret
	}
	return *o.OriginFetchRespHeaderBytes
}

// GetOriginFetchRespHeaderBytesOk returns a tuple with the OriginFetchRespHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainInspectorMeasurements) GetOriginFetchRespHeaderBytesOk() (*int64, bool) {
	if o == nil || o.OriginFetchRespHeaderBytes == nil {
		return nil, false
	}
	return o.OriginFetchRespHeaderBytes, true
}

// HasOriginFetchRespHeaderBytes returns a boolean if a field has been set.
func (o *DomainInspectorMeasurements) HasOriginFetchRespHeaderBytes() bool {
	if o != nil && o.OriginFetchRespHeaderBytes != nil {
		return true
	}

	return false
}

// SetOriginFetchRespHeaderBytes gets a reference to the given int64 and assigns it to the OriginFetchRespHeaderBytes field.
func (o *DomainInspectorMeasurements) SetOriginFetchRespHeaderBytes(v int64) {
	o.OriginFetchRespHeaderBytes = &v
}

// GetOriginFetchRespBodyBytes returns the OriginFetchRespBodyBytes field value if set, zero value otherwise.
func (o *DomainInspectorMeasurements) GetOriginFetchRespBodyBytes() int64 {
	if o == nil || o.OriginFetchRespBodyBytes == nil {
		var ret int64
		return ret
	}
	return *o.OriginFetchRespBodyBytes
}

// GetOriginFetchRespBodyBytesOk returns a tuple with the OriginFetchRespBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainInspectorMeasurements) GetOriginFetchRespBodyBytesOk() (*int64, bool) {
	if o == nil || o.OriginFetchRespBodyBytes == nil {
		return nil, false
	}
	return o.OriginFetchRespBodyBytes, true
}

// HasOriginFetchRespBodyBytes returns a boolean if a field has been set.
func (o *DomainInspectorMeasurements) HasOriginFetchRespBodyBytes() bool {
	if o != nil && o.OriginFetchRespBodyBytes != nil {
		return true
	}

	return false
}

// SetOriginFetchRespBodyBytes gets a reference to the given int64 and assigns it to the OriginFetchRespBodyBytes field.
func (o *DomainInspectorMeasurements) SetOriginFetchRespBodyBytes(v int64) {
	o.OriginFetchRespBodyBytes = &v
}

// GetBandwidth returns the Bandwidth field value if set, zero value otherwise.
func (o *DomainInspectorMeasurements) GetBandwidth() int64 {
	if o == nil || o.Bandwidth == nil {
		var ret int64
		return ret
	}
	return *o.Bandwidth
}

// GetBandwidthOk returns a tuple with the Bandwidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainInspectorMeasurements) GetBandwidthOk() (*int64, bool) {
	if o == nil || o.Bandwidth == nil {
		return nil, false
	}
	return o.Bandwidth, true
}

// HasBandwidth returns a boolean if a field has been set.
func (o *DomainInspectorMeasurements) HasBandwidth() bool {
	if o != nil && o.Bandwidth != nil {
		return true
	}

	return false
}

// SetBandwidth gets a reference to the given int64 and assigns it to the Bandwidth field.
func (o *DomainInspectorMeasurements) SetBandwidth(v int64) {
	o.Bandwidth = &v
}

// GetEdgeHitRatio returns the EdgeHitRatio field value if set, zero value otherwise.
func (o *DomainInspectorMeasurements) GetEdgeHitRatio() float32 {
	if o == nil || o.EdgeHitRatio == nil {
		var ret float32
		return ret
	}
	return *o.EdgeHitRatio
}

// GetEdgeHitRatioOk returns a tuple with the EdgeHitRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainInspectorMeasurements) GetEdgeHitRatioOk() (*float32, bool) {
	if o == nil || o.EdgeHitRatio == nil {
		return nil, false
	}
	return o.EdgeHitRatio, true
}

// HasEdgeHitRatio returns a boolean if a field has been set.
func (o *DomainInspectorMeasurements) HasEdgeHitRatio() bool {
	if o != nil && o.EdgeHitRatio != nil {
		return true
	}

	return false
}

// SetEdgeHitRatio gets a reference to the given float32 and assigns it to the EdgeHitRatio field.
func (o *DomainInspectorMeasurements) SetEdgeHitRatio(v float32) {
	o.EdgeHitRatio = &v
}

// GetOriginOffload returns the OriginOffload field value if set, zero value otherwise.
func (o *DomainInspectorMeasurements) GetOriginOffload() float32 {
	if o == nil || o.OriginOffload == nil {
		var ret float32
		return ret
	}
	return *o.OriginOffload
}

// GetOriginOffloadOk returns a tuple with the OriginOffload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainInspectorMeasurements) GetOriginOffloadOk() (*float32, bool) {
	if o == nil || o.OriginOffload == nil {
		return nil, false
	}
	return o.OriginOffload, true
}

// HasOriginOffload returns a boolean if a field has been set.
func (o *DomainInspectorMeasurements) HasOriginOffload() bool {
	if o != nil && o.OriginOffload != nil {
		return true
	}

	return false
}

// SetOriginOffload gets a reference to the given float32 and assigns it to the OriginOffload field.
func (o *DomainInspectorMeasurements) SetOriginOffload(v float32) {
	o.OriginOffload = &v
}

// GetOriginStatus200 returns the OriginStatus200 field value if set, zero value otherwise.
func (o *DomainInspectorMeasurements) GetOriginStatus200() int64 {
	if o == nil || o.OriginStatus200 == nil {
		var ret int64
		return ret
	}
	return *o.OriginStatus200
}

// GetOriginStatus200Ok returns a tuple with the OriginStatus200 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainInspectorMeasurements) GetOriginStatus200Ok() (*int64, bool) {
	if o == nil || o.OriginStatus200 == nil {
		return nil, false
	}
	return o.OriginStatus200, true
}

// HasOriginStatus200 returns a boolean if a field has been set.
func (o *DomainInspectorMeasurements) HasOriginStatus200() bool {
	if o != nil && o.OriginStatus200 != nil {
		return true
	}

	return false
}

// SetOriginStatus200 gets a reference to the given int64 and assigns it to the OriginStatus200 field.
func (o *DomainInspectorMeasurements) SetOriginStatus200(v int64) {
	o.OriginStatus200 = &v
}

// GetOriginStatus204 returns the OriginStatus204 field value if set, zero value otherwise.
func (o *DomainInspectorMeasurements) GetOriginStatus204() int64 {
	if o == nil || o.OriginStatus204 == nil {
		var ret int64
		return ret
	}
	return *o.OriginStatus204
}

// GetOriginStatus204Ok returns a tuple with the OriginStatus204 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainInspectorMeasurements) GetOriginStatus204Ok() (*int64, bool) {
	if o == nil || o.OriginStatus204 == nil {
		return nil, false
	}
	return o.OriginStatus204, true
}

// HasOriginStatus204 returns a boolean if a field has been set.
func (o *DomainInspectorMeasurements) HasOriginStatus204() bool {
	if o != nil && o.OriginStatus204 != nil {
		return true
	}

	return false
}

// SetOriginStatus204 gets a reference to the given int64 and assigns it to the OriginStatus204 field.
func (o *DomainInspectorMeasurements) SetOriginStatus204(v int64) {
	o.OriginStatus204 = &v
}

// GetOriginStatus206 returns the OriginStatus206 field value if set, zero value otherwise.
func (o *DomainInspectorMeasurements) GetOriginStatus206() int64 {
	if o == nil || o.OriginStatus206 == nil {
		var ret int64
		return ret
	}
	return *o.OriginStatus206
}

// GetOriginStatus206Ok returns a tuple with the OriginStatus206 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainInspectorMeasurements) GetOriginStatus206Ok() (*int64, bool) {
	if o == nil || o.OriginStatus206 == nil {
		return nil, false
	}
	return o.OriginStatus206, true
}

// HasOriginStatus206 returns a boolean if a field has been set.
func (o *DomainInspectorMeasurements) HasOriginStatus206() bool {
	if o != nil && o.OriginStatus206 != nil {
		return true
	}

	return false
}

// SetOriginStatus206 gets a reference to the given int64 and assigns it to the OriginStatus206 field.
func (o *DomainInspectorMeasurements) SetOriginStatus206(v int64) {
	o.OriginStatus206 = &v
}

// GetOriginStatus301 returns the OriginStatus301 field value if set, zero value otherwise.
func (o *DomainInspectorMeasurements) GetOriginStatus301() int64 {
	if o == nil || o.OriginStatus301 == nil {
		var ret int64
		return ret
	}
	return *o.OriginStatus301
}

// GetOriginStatus301Ok returns a tuple with the OriginStatus301 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainInspectorMeasurements) GetOriginStatus301Ok() (*int64, bool) {
	if o == nil || o.OriginStatus301 == nil {
		return nil, false
	}
	return o.OriginStatus301, true
}

// HasOriginStatus301 returns a boolean if a field has been set.
func (o *DomainInspectorMeasurements) HasOriginStatus301() bool {
	if o != nil && o.OriginStatus301 != nil {
		return true
	}

	return false
}

// SetOriginStatus301 gets a reference to the given int64 and assigns it to the OriginStatus301 field.
func (o *DomainInspectorMeasurements) SetOriginStatus301(v int64) {
	o.OriginStatus301 = &v
}

// GetOriginStatus302 returns the OriginStatus302 field value if set, zero value otherwise.
func (o *DomainInspectorMeasurements) GetOriginStatus302() int64 {
	if o == nil || o.OriginStatus302 == nil {
		var ret int64
		return ret
	}
	return *o.OriginStatus302
}

// GetOriginStatus302Ok returns a tuple with the OriginStatus302 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainInspectorMeasurements) GetOriginStatus302Ok() (*int64, bool) {
	if o == nil || o.OriginStatus302 == nil {
		return nil, false
	}
	return o.OriginStatus302, true
}

// HasOriginStatus302 returns a boolean if a field has been set.
func (o *DomainInspectorMeasurements) HasOriginStatus302() bool {
	if o != nil && o.OriginStatus302 != nil {
		return true
	}

	return false
}

// SetOriginStatus302 gets a reference to the given int64 and assigns it to the OriginStatus302 field.
func (o *DomainInspectorMeasurements) SetOriginStatus302(v int64) {
	o.OriginStatus302 = &v
}

// GetOriginStatus304 returns the OriginStatus304 field value if set, zero value otherwise.
func (o *DomainInspectorMeasurements) GetOriginStatus304() int64 {
	if o == nil || o.OriginStatus304 == nil {
		var ret int64
		return ret
	}
	return *o.OriginStatus304
}

// GetOriginStatus304Ok returns a tuple with the OriginStatus304 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainInspectorMeasurements) GetOriginStatus304Ok() (*int64, bool) {
	if o == nil || o.OriginStatus304 == nil {
		return nil, false
	}
	return o.OriginStatus304, true
}

// HasOriginStatus304 returns a boolean if a field has been set.
func (o *DomainInspectorMeasurements) HasOriginStatus304() bool {
	if o != nil && o.OriginStatus304 != nil {
		return true
	}

	return false
}

// SetOriginStatus304 gets a reference to the given int64 and assigns it to the OriginStatus304 field.
func (o *DomainInspectorMeasurements) SetOriginStatus304(v int64) {
	o.OriginStatus304 = &v
}

// GetOriginStatus400 returns the OriginStatus400 field value if set, zero value otherwise.
func (o *DomainInspectorMeasurements) GetOriginStatus400() int64 {
	if o == nil || o.OriginStatus400 == nil {
		var ret int64
		return ret
	}
	return *o.OriginStatus400
}

// GetOriginStatus400Ok returns a tuple with the OriginStatus400 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainInspectorMeasurements) GetOriginStatus400Ok() (*int64, bool) {
	if o == nil || o.OriginStatus400 == nil {
		return nil, false
	}
	return o.OriginStatus400, true
}

// HasOriginStatus400 returns a boolean if a field has been set.
func (o *DomainInspectorMeasurements) HasOriginStatus400() bool {
	if o != nil && o.OriginStatus400 != nil {
		return true
	}

	return false
}

// SetOriginStatus400 gets a reference to the given int64 and assigns it to the OriginStatus400 field.
func (o *DomainInspectorMeasurements) SetOriginStatus400(v int64) {
	o.OriginStatus400 = &v
}

// GetOriginStatus401 returns the OriginStatus401 field value if set, zero value otherwise.
func (o *DomainInspectorMeasurements) GetOriginStatus401() int64 {
	if o == nil || o.OriginStatus401 == nil {
		var ret int64
		return ret
	}
	return *o.OriginStatus401
}

// GetOriginStatus401Ok returns a tuple with the OriginStatus401 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainInspectorMeasurements) GetOriginStatus401Ok() (*int64, bool) {
	if o == nil || o.OriginStatus401 == nil {
		return nil, false
	}
	return o.OriginStatus401, true
}

// HasOriginStatus401 returns a boolean if a field has been set.
func (o *DomainInspectorMeasurements) HasOriginStatus401() bool {
	if o != nil && o.OriginStatus401 != nil {
		return true
	}

	return false
}

// SetOriginStatus401 gets a reference to the given int64 and assigns it to the OriginStatus401 field.
func (o *DomainInspectorMeasurements) SetOriginStatus401(v int64) {
	o.OriginStatus401 = &v
}

// GetOriginStatus403 returns the OriginStatus403 field value if set, zero value otherwise.
func (o *DomainInspectorMeasurements) GetOriginStatus403() int64 {
	if o == nil || o.OriginStatus403 == nil {
		var ret int64
		return ret
	}
	return *o.OriginStatus403
}

// GetOriginStatus403Ok returns a tuple with the OriginStatus403 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainInspectorMeasurements) GetOriginStatus403Ok() (*int64, bool) {
	if o == nil || o.OriginStatus403 == nil {
		return nil, false
	}
	return o.OriginStatus403, true
}

// HasOriginStatus403 returns a boolean if a field has been set.
func (o *DomainInspectorMeasurements) HasOriginStatus403() bool {
	if o != nil && o.OriginStatus403 != nil {
		return true
	}

	return false
}

// SetOriginStatus403 gets a reference to the given int64 and assigns it to the OriginStatus403 field.
func (o *DomainInspectorMeasurements) SetOriginStatus403(v int64) {
	o.OriginStatus403 = &v
}

// GetOriginStatus404 returns the OriginStatus404 field value if set, zero value otherwise.
func (o *DomainInspectorMeasurements) GetOriginStatus404() int64 {
	if o == nil || o.OriginStatus404 == nil {
		var ret int64
		return ret
	}
	return *o.OriginStatus404
}

// GetOriginStatus404Ok returns a tuple with the OriginStatus404 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainInspectorMeasurements) GetOriginStatus404Ok() (*int64, bool) {
	if o == nil || o.OriginStatus404 == nil {
		return nil, false
	}
	return o.OriginStatus404, true
}

// HasOriginStatus404 returns a boolean if a field has been set.
func (o *DomainInspectorMeasurements) HasOriginStatus404() bool {
	if o != nil && o.OriginStatus404 != nil {
		return true
	}

	return false
}

// SetOriginStatus404 gets a reference to the given int64 and assigns it to the OriginStatus404 field.
func (o *DomainInspectorMeasurements) SetOriginStatus404(v int64) {
	o.OriginStatus404 = &v
}

// GetOriginStatus416 returns the OriginStatus416 field value if set, zero value otherwise.
func (o *DomainInspectorMeasurements) GetOriginStatus416() int64 {
	if o == nil || o.OriginStatus416 == nil {
		var ret int64
		return ret
	}
	return *o.OriginStatus416
}

// GetOriginStatus416Ok returns a tuple with the OriginStatus416 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainInspectorMeasurements) GetOriginStatus416Ok() (*int64, bool) {
	if o == nil || o.OriginStatus416 == nil {
		return nil, false
	}
	return o.OriginStatus416, true
}

// HasOriginStatus416 returns a boolean if a field has been set.
func (o *DomainInspectorMeasurements) HasOriginStatus416() bool {
	if o != nil && o.OriginStatus416 != nil {
		return true
	}

	return false
}

// SetOriginStatus416 gets a reference to the given int64 and assigns it to the OriginStatus416 field.
func (o *DomainInspectorMeasurements) SetOriginStatus416(v int64) {
	o.OriginStatus416 = &v
}

// GetOriginStatus429 returns the OriginStatus429 field value if set, zero value otherwise.
func (o *DomainInspectorMeasurements) GetOriginStatus429() int64 {
	if o == nil || o.OriginStatus429 == nil {
		var ret int64
		return ret
	}
	return *o.OriginStatus429
}

// GetOriginStatus429Ok returns a tuple with the OriginStatus429 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainInspectorMeasurements) GetOriginStatus429Ok() (*int64, bool) {
	if o == nil || o.OriginStatus429 == nil {
		return nil, false
	}
	return o.OriginStatus429, true
}

// HasOriginStatus429 returns a boolean if a field has been set.
func (o *DomainInspectorMeasurements) HasOriginStatus429() bool {
	if o != nil && o.OriginStatus429 != nil {
		return true
	}

	return false
}

// SetOriginStatus429 gets a reference to the given int64 and assigns it to the OriginStatus429 field.
func (o *DomainInspectorMeasurements) SetOriginStatus429(v int64) {
	o.OriginStatus429 = &v
}

// GetOriginStatus500 returns the OriginStatus500 field value if set, zero value otherwise.
func (o *DomainInspectorMeasurements) GetOriginStatus500() int64 {
	if o == nil || o.OriginStatus500 == nil {
		var ret int64
		return ret
	}
	return *o.OriginStatus500
}

// GetOriginStatus500Ok returns a tuple with the OriginStatus500 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainInspectorMeasurements) GetOriginStatus500Ok() (*int64, bool) {
	if o == nil || o.OriginStatus500 == nil {
		return nil, false
	}
	return o.OriginStatus500, true
}

// HasOriginStatus500 returns a boolean if a field has been set.
func (o *DomainInspectorMeasurements) HasOriginStatus500() bool {
	if o != nil && o.OriginStatus500 != nil {
		return true
	}

	return false
}

// SetOriginStatus500 gets a reference to the given int64 and assigns it to the OriginStatus500 field.
func (o *DomainInspectorMeasurements) SetOriginStatus500(v int64) {
	o.OriginStatus500 = &v
}

// GetOriginStatus501 returns the OriginStatus501 field value if set, zero value otherwise.
func (o *DomainInspectorMeasurements) GetOriginStatus501() int64 {
	if o == nil || o.OriginStatus501 == nil {
		var ret int64
		return ret
	}
	return *o.OriginStatus501
}

// GetOriginStatus501Ok returns a tuple with the OriginStatus501 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainInspectorMeasurements) GetOriginStatus501Ok() (*int64, bool) {
	if o == nil || o.OriginStatus501 == nil {
		return nil, false
	}
	return o.OriginStatus501, true
}

// HasOriginStatus501 returns a boolean if a field has been set.
func (o *DomainInspectorMeasurements) HasOriginStatus501() bool {
	if o != nil && o.OriginStatus501 != nil {
		return true
	}

	return false
}

// SetOriginStatus501 gets a reference to the given int64 and assigns it to the OriginStatus501 field.
func (o *DomainInspectorMeasurements) SetOriginStatus501(v int64) {
	o.OriginStatus501 = &v
}

// GetOriginStatus502 returns the OriginStatus502 field value if set, zero value otherwise.
func (o *DomainInspectorMeasurements) GetOriginStatus502() int64 {
	if o == nil || o.OriginStatus502 == nil {
		var ret int64
		return ret
	}
	return *o.OriginStatus502
}

// GetOriginStatus502Ok returns a tuple with the OriginStatus502 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainInspectorMeasurements) GetOriginStatus502Ok() (*int64, bool) {
	if o == nil || o.OriginStatus502 == nil {
		return nil, false
	}
	return o.OriginStatus502, true
}

// HasOriginStatus502 returns a boolean if a field has been set.
func (o *DomainInspectorMeasurements) HasOriginStatus502() bool {
	if o != nil && o.OriginStatus502 != nil {
		return true
	}

	return false
}

// SetOriginStatus502 gets a reference to the given int64 and assigns it to the OriginStatus502 field.
func (o *DomainInspectorMeasurements) SetOriginStatus502(v int64) {
	o.OriginStatus502 = &v
}

// GetOriginStatus503 returns the OriginStatus503 field value if set, zero value otherwise.
func (o *DomainInspectorMeasurements) GetOriginStatus503() int64 {
	if o == nil || o.OriginStatus503 == nil {
		var ret int64
		return ret
	}
	return *o.OriginStatus503
}

// GetOriginStatus503Ok returns a tuple with the OriginStatus503 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainInspectorMeasurements) GetOriginStatus503Ok() (*int64, bool) {
	if o == nil || o.OriginStatus503 == nil {
		return nil, false
	}
	return o.OriginStatus503, true
}

// HasOriginStatus503 returns a boolean if a field has been set.
func (o *DomainInspectorMeasurements) HasOriginStatus503() bool {
	if o != nil && o.OriginStatus503 != nil {
		return true
	}

	return false
}

// SetOriginStatus503 gets a reference to the given int64 and assigns it to the OriginStatus503 field.
func (o *DomainInspectorMeasurements) SetOriginStatus503(v int64) {
	o.OriginStatus503 = &v
}

// GetOriginStatus504 returns the OriginStatus504 field value if set, zero value otherwise.
func (o *DomainInspectorMeasurements) GetOriginStatus504() int64 {
	if o == nil || o.OriginStatus504 == nil {
		var ret int64
		return ret
	}
	return *o.OriginStatus504
}

// GetOriginStatus504Ok returns a tuple with the OriginStatus504 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainInspectorMeasurements) GetOriginStatus504Ok() (*int64, bool) {
	if o == nil || o.OriginStatus504 == nil {
		return nil, false
	}
	return o.OriginStatus504, true
}

// HasOriginStatus504 returns a boolean if a field has been set.
func (o *DomainInspectorMeasurements) HasOriginStatus504() bool {
	if o != nil && o.OriginStatus504 != nil {
		return true
	}

	return false
}

// SetOriginStatus504 gets a reference to the given int64 and assigns it to the OriginStatus504 field.
func (o *DomainInspectorMeasurements) SetOriginStatus504(v int64) {
	o.OriginStatus504 = &v
}

// GetOriginStatus505 returns the OriginStatus505 field value if set, zero value otherwise.
func (o *DomainInspectorMeasurements) GetOriginStatus505() int64 {
	if o == nil || o.OriginStatus505 == nil {
		var ret int64
		return ret
	}
	return *o.OriginStatus505
}

// GetOriginStatus505Ok returns a tuple with the OriginStatus505 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainInspectorMeasurements) GetOriginStatus505Ok() (*int64, bool) {
	if o == nil || o.OriginStatus505 == nil {
		return nil, false
	}
	return o.OriginStatus505, true
}

// HasOriginStatus505 returns a boolean if a field has been set.
func (o *DomainInspectorMeasurements) HasOriginStatus505() bool {
	if o != nil && o.OriginStatus505 != nil {
		return true
	}

	return false
}

// SetOriginStatus505 gets a reference to the given int64 and assigns it to the OriginStatus505 field.
func (o *DomainInspectorMeasurements) SetOriginStatus505(v int64) {
	o.OriginStatus505 = &v
}

// GetOriginStatus530 returns the OriginStatus530 field value if set, zero value otherwise.
func (o *DomainInspectorMeasurements) GetOriginStatus530() int64 {
	if o == nil || o.OriginStatus530 == nil {
		var ret int64
		return ret
	}
	return *o.OriginStatus530
}

// GetOriginStatus530Ok returns a tuple with the OriginStatus530 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainInspectorMeasurements) GetOriginStatus530Ok() (*int64, bool) {
	if o == nil || o.OriginStatus530 == nil {
		return nil, false
	}
	return o.OriginStatus530, true
}

// HasOriginStatus530 returns a boolean if a field has been set.
func (o *DomainInspectorMeasurements) HasOriginStatus530() bool {
	if o != nil && o.OriginStatus530 != nil {
		return true
	}

	return false
}

// SetOriginStatus530 gets a reference to the given int64 and assigns it to the OriginStatus530 field.
func (o *DomainInspectorMeasurements) SetOriginStatus530(v int64) {
	o.OriginStatus530 = &v
}

// GetOriginStatus1xx returns the OriginStatus1xx field value if set, zero value otherwise.
func (o *DomainInspectorMeasurements) GetOriginStatus1xx() int64 {
	if o == nil || o.OriginStatus1xx == nil {
		var ret int64
		return ret
	}
	return *o.OriginStatus1xx
}

// GetOriginStatus1xxOk returns a tuple with the OriginStatus1xx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainInspectorMeasurements) GetOriginStatus1xxOk() (*int64, bool) {
	if o == nil || o.OriginStatus1xx == nil {
		return nil, false
	}
	return o.OriginStatus1xx, true
}

// HasOriginStatus1xx returns a boolean if a field has been set.
func (o *DomainInspectorMeasurements) HasOriginStatus1xx() bool {
	if o != nil && o.OriginStatus1xx != nil {
		return true
	}

	return false
}

// SetOriginStatus1xx gets a reference to the given int64 and assigns it to the OriginStatus1xx field.
func (o *DomainInspectorMeasurements) SetOriginStatus1xx(v int64) {
	o.OriginStatus1xx = &v
}

// GetOriginStatus2xx returns the OriginStatus2xx field value if set, zero value otherwise.
func (o *DomainInspectorMeasurements) GetOriginStatus2xx() int64 {
	if o == nil || o.OriginStatus2xx == nil {
		var ret int64
		return ret
	}
	return *o.OriginStatus2xx
}

// GetOriginStatus2xxOk returns a tuple with the OriginStatus2xx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainInspectorMeasurements) GetOriginStatus2xxOk() (*int64, bool) {
	if o == nil || o.OriginStatus2xx == nil {
		return nil, false
	}
	return o.OriginStatus2xx, true
}

// HasOriginStatus2xx returns a boolean if a field has been set.
func (o *DomainInspectorMeasurements) HasOriginStatus2xx() bool {
	if o != nil && o.OriginStatus2xx != nil {
		return true
	}

	return false
}

// SetOriginStatus2xx gets a reference to the given int64 and assigns it to the OriginStatus2xx field.
func (o *DomainInspectorMeasurements) SetOriginStatus2xx(v int64) {
	o.OriginStatus2xx = &v
}

// GetOriginStatus3xx returns the OriginStatus3xx field value if set, zero value otherwise.
func (o *DomainInspectorMeasurements) GetOriginStatus3xx() int64 {
	if o == nil || o.OriginStatus3xx == nil {
		var ret int64
		return ret
	}
	return *o.OriginStatus3xx
}

// GetOriginStatus3xxOk returns a tuple with the OriginStatus3xx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainInspectorMeasurements) GetOriginStatus3xxOk() (*int64, bool) {
	if o == nil || o.OriginStatus3xx == nil {
		return nil, false
	}
	return o.OriginStatus3xx, true
}

// HasOriginStatus3xx returns a boolean if a field has been set.
func (o *DomainInspectorMeasurements) HasOriginStatus3xx() bool {
	if o != nil && o.OriginStatus3xx != nil {
		return true
	}

	return false
}

// SetOriginStatus3xx gets a reference to the given int64 and assigns it to the OriginStatus3xx field.
func (o *DomainInspectorMeasurements) SetOriginStatus3xx(v int64) {
	o.OriginStatus3xx = &v
}

// GetOriginStatus4xx returns the OriginStatus4xx field value if set, zero value otherwise.
func (o *DomainInspectorMeasurements) GetOriginStatus4xx() int64 {
	if o == nil || o.OriginStatus4xx == nil {
		var ret int64
		return ret
	}
	return *o.OriginStatus4xx
}

// GetOriginStatus4xxOk returns a tuple with the OriginStatus4xx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainInspectorMeasurements) GetOriginStatus4xxOk() (*int64, bool) {
	if o == nil || o.OriginStatus4xx == nil {
		return nil, false
	}
	return o.OriginStatus4xx, true
}

// HasOriginStatus4xx returns a boolean if a field has been set.
func (o *DomainInspectorMeasurements) HasOriginStatus4xx() bool {
	if o != nil && o.OriginStatus4xx != nil {
		return true
	}

	return false
}

// SetOriginStatus4xx gets a reference to the given int64 and assigns it to the OriginStatus4xx field.
func (o *DomainInspectorMeasurements) SetOriginStatus4xx(v int64) {
	o.OriginStatus4xx = &v
}

// GetOriginStatus5xx returns the OriginStatus5xx field value if set, zero value otherwise.
func (o *DomainInspectorMeasurements) GetOriginStatus5xx() int64 {
	if o == nil || o.OriginStatus5xx == nil {
		var ret int64
		return ret
	}
	return *o.OriginStatus5xx
}

// GetOriginStatus5xxOk returns a tuple with the OriginStatus5xx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainInspectorMeasurements) GetOriginStatus5xxOk() (*int64, bool) {
	if o == nil || o.OriginStatus5xx == nil {
		return nil, false
	}
	return o.OriginStatus5xx, true
}

// HasOriginStatus5xx returns a boolean if a field has been set.
func (o *DomainInspectorMeasurements) HasOriginStatus5xx() bool {
	if o != nil && o.OriginStatus5xx != nil {
		return true
	}

	return false
}

// SetOriginStatus5xx gets a reference to the given int64 and assigns it to the OriginStatus5xx field.
func (o *DomainInspectorMeasurements) SetOriginStatus5xx(v int64) {
	o.OriginStatus5xx = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o DomainInspectorMeasurements) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.EdgeRequests != nil {
		toSerialize["edge_requests"] = o.EdgeRequests
	}
	if o.EdgeRespHeaderBytes != nil {
		toSerialize["edge_resp_header_bytes"] = o.EdgeRespHeaderBytes
	}
	if o.EdgeRespBodyBytes != nil {
		toSerialize["edge_resp_body_bytes"] = o.EdgeRespBodyBytes
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
	if o.Status530 != nil {
		toSerialize["status_530"] = o.Status530
	}
	if o.Requests != nil {
		toSerialize["requests"] = o.Requests
	}
	if o.RespHeaderBytes != nil {
		toSerialize["resp_header_bytes"] = o.RespHeaderBytes
	}
	if o.RespBodyBytes != nil {
		toSerialize["resp_body_bytes"] = o.RespBodyBytes
	}
	if o.BereqHeaderBytes != nil {
		toSerialize["bereq_header_bytes"] = o.BereqHeaderBytes
	}
	if o.BereqBodyBytes != nil {
		toSerialize["bereq_body_bytes"] = o.BereqBodyBytes
	}
	if o.EdgeHitRequests != nil {
		toSerialize["edge_hit_requests"] = o.EdgeHitRequests
	}
	if o.EdgeMissRequests != nil {
		toSerialize["edge_miss_requests"] = o.EdgeMissRequests
	}
	if o.OriginFetches != nil {
		toSerialize["origin_fetches"] = o.OriginFetches
	}
	if o.OriginFetchRespHeaderBytes != nil {
		toSerialize["origin_fetch_resp_header_bytes"] = o.OriginFetchRespHeaderBytes
	}
	if o.OriginFetchRespBodyBytes != nil {
		toSerialize["origin_fetch_resp_body_bytes"] = o.OriginFetchRespBodyBytes
	}
	if o.Bandwidth != nil {
		toSerialize["bandwidth"] = o.Bandwidth
	}
	if o.EdgeHitRatio != nil {
		toSerialize["edge_hit_ratio"] = o.EdgeHitRatio
	}
	if o.OriginOffload != nil {
		toSerialize["origin_offload"] = o.OriginOffload
	}
	if o.OriginStatus200 != nil {
		toSerialize["origin_status_200"] = o.OriginStatus200
	}
	if o.OriginStatus204 != nil {
		toSerialize["origin_status_204"] = o.OriginStatus204
	}
	if o.OriginStatus206 != nil {
		toSerialize["origin_status_206"] = o.OriginStatus206
	}
	if o.OriginStatus301 != nil {
		toSerialize["origin_status_301"] = o.OriginStatus301
	}
	if o.OriginStatus302 != nil {
		toSerialize["origin_status_302"] = o.OriginStatus302
	}
	if o.OriginStatus304 != nil {
		toSerialize["origin_status_304"] = o.OriginStatus304
	}
	if o.OriginStatus400 != nil {
		toSerialize["origin_status_400"] = o.OriginStatus400
	}
	if o.OriginStatus401 != nil {
		toSerialize["origin_status_401"] = o.OriginStatus401
	}
	if o.OriginStatus403 != nil {
		toSerialize["origin_status_403"] = o.OriginStatus403
	}
	if o.OriginStatus404 != nil {
		toSerialize["origin_status_404"] = o.OriginStatus404
	}
	if o.OriginStatus416 != nil {
		toSerialize["origin_status_416"] = o.OriginStatus416
	}
	if o.OriginStatus429 != nil {
		toSerialize["origin_status_429"] = o.OriginStatus429
	}
	if o.OriginStatus500 != nil {
		toSerialize["origin_status_500"] = o.OriginStatus500
	}
	if o.OriginStatus501 != nil {
		toSerialize["origin_status_501"] = o.OriginStatus501
	}
	if o.OriginStatus502 != nil {
		toSerialize["origin_status_502"] = o.OriginStatus502
	}
	if o.OriginStatus503 != nil {
		toSerialize["origin_status_503"] = o.OriginStatus503
	}
	if o.OriginStatus504 != nil {
		toSerialize["origin_status_504"] = o.OriginStatus504
	}
	if o.OriginStatus505 != nil {
		toSerialize["origin_status_505"] = o.OriginStatus505
	}
	if o.OriginStatus530 != nil {
		toSerialize["origin_status_530"] = o.OriginStatus530
	}
	if o.OriginStatus1xx != nil {
		toSerialize["origin_status_1xx"] = o.OriginStatus1xx
	}
	if o.OriginStatus2xx != nil {
		toSerialize["origin_status_2xx"] = o.OriginStatus2xx
	}
	if o.OriginStatus3xx != nil {
		toSerialize["origin_status_3xx"] = o.OriginStatus3xx
	}
	if o.OriginStatus4xx != nil {
		toSerialize["origin_status_4xx"] = o.OriginStatus4xx
	}
	if o.OriginStatus5xx != nil {
		toSerialize["origin_status_5xx"] = o.OriginStatus5xx
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *DomainInspectorMeasurements) UnmarshalJSON(bytes []byte) (err error) {
	varDomainInspectorMeasurements := _DomainInspectorMeasurements{}

	if err = json.Unmarshal(bytes, &varDomainInspectorMeasurements); err == nil {
		*o = DomainInspectorMeasurements(varDomainInspectorMeasurements)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "edge_requests")
		delete(additionalProperties, "edge_resp_header_bytes")
		delete(additionalProperties, "edge_resp_body_bytes")
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
		delete(additionalProperties, "status_530")
		delete(additionalProperties, "requests")
		delete(additionalProperties, "resp_header_bytes")
		delete(additionalProperties, "resp_body_bytes")
		delete(additionalProperties, "bereq_header_bytes")
		delete(additionalProperties, "bereq_body_bytes")
		delete(additionalProperties, "edge_hit_requests")
		delete(additionalProperties, "edge_miss_requests")
		delete(additionalProperties, "origin_fetches")
		delete(additionalProperties, "origin_fetch_resp_header_bytes")
		delete(additionalProperties, "origin_fetch_resp_body_bytes")
		delete(additionalProperties, "bandwidth")
		delete(additionalProperties, "edge_hit_ratio")
		delete(additionalProperties, "origin_offload")
		delete(additionalProperties, "origin_status_200")
		delete(additionalProperties, "origin_status_204")
		delete(additionalProperties, "origin_status_206")
		delete(additionalProperties, "origin_status_301")
		delete(additionalProperties, "origin_status_302")
		delete(additionalProperties, "origin_status_304")
		delete(additionalProperties, "origin_status_400")
		delete(additionalProperties, "origin_status_401")
		delete(additionalProperties, "origin_status_403")
		delete(additionalProperties, "origin_status_404")
		delete(additionalProperties, "origin_status_416")
		delete(additionalProperties, "origin_status_429")
		delete(additionalProperties, "origin_status_500")
		delete(additionalProperties, "origin_status_501")
		delete(additionalProperties, "origin_status_502")
		delete(additionalProperties, "origin_status_503")
		delete(additionalProperties, "origin_status_504")
		delete(additionalProperties, "origin_status_505")
		delete(additionalProperties, "origin_status_530")
		delete(additionalProperties, "origin_status_1xx")
		delete(additionalProperties, "origin_status_2xx")
		delete(additionalProperties, "origin_status_3xx")
		delete(additionalProperties, "origin_status_4xx")
		delete(additionalProperties, "origin_status_5xx")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableDomainInspectorMeasurements is a helper abstraction for handling nullable domaininspectormeasurements types.
type NullableDomainInspectorMeasurements struct {
	value *DomainInspectorMeasurements
	isSet bool
}

// Get returns the value.
func (v NullableDomainInspectorMeasurements) Get() *DomainInspectorMeasurements {
	return v.value
}

// Set modifies the value.
func (v *NullableDomainInspectorMeasurements) Set(val *DomainInspectorMeasurements) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableDomainInspectorMeasurements) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableDomainInspectorMeasurements) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDomainInspectorMeasurements returns a pointer to a new instance of NullableDomainInspectorMeasurements.
func NewNullableDomainInspectorMeasurements(val *DomainInspectorMeasurements) *NullableDomainInspectorMeasurements {
	return &NullableDomainInspectorMeasurements{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableDomainInspectorMeasurements) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableDomainInspectorMeasurements) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
