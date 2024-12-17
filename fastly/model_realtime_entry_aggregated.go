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

// RealtimeEntryAggregated Aggregates [measurements](#measurements-data-model) across all Fastly POPs.
type RealtimeEntryAggregated struct {
	// Number of requests processed.
	Requests *int64 `json:"requests,omitempty"`
	// Number of log lines sent (alias for `log`).
	Logging *int64 `json:"logging,omitempty"`
	// Number of log lines sent.
	Log *int64 `json:"log,omitempty"`
	// Total header bytes delivered (edge_resp_header_bytes + shield_resp_header_bytes).
	RespHeaderBytes *int64 `json:"resp_header_bytes,omitempty"`
	// Total header bytes delivered (alias for resp_header_bytes).
	HeaderSize *int64 `json:"header_size,omitempty"`
	// Total body bytes delivered (edge_resp_body_bytes + shield_resp_body_bytes).
	RespBodyBytes *int64 `json:"resp_body_bytes,omitempty"`
	// Total body bytes delivered (alias for resp_body_bytes).
	BodySize *int64 `json:"body_size,omitempty"`
	// Number of cache hits.
	Hits *int64 `json:"hits,omitempty"`
	// Number of cache misses.
	Miss *int64 `json:"miss,omitempty"`
	// Number of requests that passed through the CDN without being cached.
	Pass *int64 `json:"pass,omitempty"`
	// Number of requests that returned a synthetic response (i.e., response objects created with the `synthetic` VCL statement).
	Synth *int64 `json:"synth,omitempty"`
	// Number of cache errors.
	Errors *int64 `json:"errors,omitempty"`
	// Total amount of time spent processing cache hits (in seconds).
	HitsTime *float32 `json:"hits_time,omitempty"`
	// Total amount of time spent processing cache misses (in seconds).
	MissTime *float32 `json:"miss_time,omitempty"`
	// A histogram. The value in each bucket is the number of requests to the origin whose responses arrived during the time period represented by the bucket. The key of each bucket represents the upper bound (in response time) of that bucket. The buckets vary in width and cover the time periods 0-10ms (in 1ms increments), 10-250ms (in 10ms increments), 250-1,000ms (in 50ms increments), 1,000-3,000ms (in 100ms increments), 3,000-10,000ms (in 500 ms increments), 10,000-20,000ms (in 1,000ms increments), 20,000-60,000ms (in 5,000ms increments), and 60,000ms through infinity (in a single bucket).
	MissHistogram map[string]map[string]any `json:"miss_histogram,omitempty"`
	// The total number of requests that were received for your service by Fastly.
	ComputeRequests *int64 `json:"compute_requests,omitempty"`
	// The amount of active CPU time used to process your requests (in milliseconds).
	ComputeExecutionTimeMs *float32 `json:"compute_execution_time_ms,omitempty"`
	// The amount of RAM used for your service by Fastly (in bytes).
	ComputeRAMUsed *int64 `json:"compute_ram_used,omitempty"`
	// The total, actual amount of time used to process your requests, including active CPU time (in milliseconds).
	ComputeRequestTimeMs *float32 `json:"compute_request_time_ms,omitempty"`
	// The total amount of request processing time you will be billed for, measured in 50 millisecond increments.
	ComputeRequestTimeBilledMs *float32 `json:"compute_request_time_billed_ms,omitempty"`
	// Number of requests from edge to the shield POP.
	Shield *int64 `json:"shield,omitempty"`
	// Number of requests that were received over IPv6.
	Ipv6 *int64 `json:"ipv6,omitempty"`
	// Number of responses that came from the Fastly Image Optimizer service. If the service receives 10 requests for an image, this stat will be 10 regardless of how many times the image was transformed.
	Imgopto *int64 `json:"imgopto,omitempty"`
	// Number of responses that came from the Fastly Image Optimizer service via a shield.
	ImgoptoShield *int64 `json:"imgopto_shield,omitempty"`
	// Number of transforms performed by the Fastly Image Optimizer service.
	ImgoptoTransforms *int64 `json:"imgopto_transforms,omitempty"`
	// Number of responses that came from the Fastly On-the-Fly Packaging service for video-on-demand.
	Otfp *int64 `json:"otfp,omitempty"`
	// Number of responses that came from the Fastly On-the-Fly Packaging service for video-on-demand via a shield.
	OtfpShield *int64 `json:"otfp_shield,omitempty"`
	// Number of responses that were manifest files from the Fastly On-the-Fly Packaging service for video-on-demand.
	OtfpManifests *int64 `json:"otfp_manifests,omitempty"`
	// Number of responses with the video segment or video manifest MIME type (i.e., application/x-mpegurl, application/vnd.apple.mpegurl, application/f4m, application/dash+xml, application/vnd.ms-sstr+xml, ideo/mp2t, audio/aac, video/f4f, video/x-flv, video/mp4, audio/mp4).
	Video *int64 `json:"video,omitempty"`
	// Number of responses with the PCI flag turned on.
	Pci *int64 `json:"pci,omitempty"`
	// Number of requests received over HTTP/2.
	HTTP2 *int64 `json:"http2,omitempty"`
	// Number of requests received over HTTP/3.
	HTTP3 *int64 `json:"http3,omitempty"`
	// Number of restarts performed.
	Restarts *int64 `json:"restarts,omitempty"`
	// Total header bytes received.
	ReqHeaderBytes *int64 `json:"req_header_bytes,omitempty"`
	// Total body bytes received.
	ReqBodyBytes *int64 `json:"req_body_bytes,omitempty"`
	// Total header bytes sent to origin.
	BereqHeaderBytes *int64 `json:"bereq_header_bytes,omitempty"`
	// Total body bytes sent to origin.
	BereqBodyBytes *int64 `json:"bereq_body_bytes,omitempty"`
	// Number of requests that triggered a WAF rule and were blocked.
	WafBlocked *int64 `json:"waf_blocked,omitempty"`
	// Number of requests that triggered a WAF rule and were logged.
	WafLogged *int64 `json:"waf_logged,omitempty"`
	// Number of requests that triggered a WAF rule and were passed.
	WafPassed *int64 `json:"waf_passed,omitempty"`
	// Total header bytes received from requests that triggered a WAF rule.
	AttackReqHeaderBytes *int64 `json:"attack_req_header_bytes,omitempty"`
	// Total body bytes received from requests that triggered a WAF rule.
	AttackReqBodyBytes *int64 `json:"attack_req_body_bytes,omitempty"`
	// Total bytes delivered for requests that triggered a WAF rule and returned a synthetic response.
	AttackRespSynthBytes *int64 `json:"attack_resp_synth_bytes,omitempty"`
	// Total header bytes received from requests that triggered a WAF rule that was logged.
	AttackLoggedReqHeaderBytes *int64 `json:"attack_logged_req_header_bytes,omitempty"`
	// Total body bytes received from requests that triggered a WAF rule that was logged.
	AttackLoggedReqBodyBytes *int64 `json:"attack_logged_req_body_bytes,omitempty"`
	// Total header bytes received from requests that triggered a WAF rule that was blocked.
	AttackBlockedReqHeaderBytes *int64 `json:"attack_blocked_req_header_bytes,omitempty"`
	// Total body bytes received from requests that triggered a WAF rule that was blocked.
	AttackBlockedReqBodyBytes *int64 `json:"attack_blocked_req_body_bytes,omitempty"`
	// Total header bytes received from requests that triggered a WAF rule that was passed.
	AttackPassedReqHeaderBytes *int64 `json:"attack_passed_req_header_bytes,omitempty"`
	// Total body bytes received from requests that triggered a WAF rule that was passed.
	AttackPassedReqBodyBytes *int64 `json:"attack_passed_req_body_bytes,omitempty"`
	// Total header bytes delivered via a shield.
	ShieldRespHeaderBytes *int64 `json:"shield_resp_header_bytes,omitempty"`
	// Total body bytes delivered via a shield.
	ShieldRespBodyBytes *int64 `json:"shield_resp_body_bytes,omitempty"`
	// Total header bytes delivered from the Fastly On-the-Fly Packaging service for video-on-demand.
	OtfpRespHeaderBytes *int64 `json:"otfp_resp_header_bytes,omitempty"`
	// Total body bytes delivered from the Fastly On-the-Fly Packaging service for video-on-demand.
	OtfpRespBodyBytes *int64 `json:"otfp_resp_body_bytes,omitempty"`
	// Total header bytes delivered via a shield for the Fastly On-the-Fly Packaging service for video-on-demand.
	OtfpShieldRespHeaderBytes *int64 `json:"otfp_shield_resp_header_bytes,omitempty"`
	// Total body bytes delivered via a shield for the Fastly On-the-Fly Packaging service for video-on-demand.
	OtfpShieldRespBodyBytes *int64 `json:"otfp_shield_resp_body_bytes,omitempty"`
	// Total amount of time spent delivering a response via a shield from the Fastly On-the-Fly Packaging service for video-on-demand (in seconds).
	OtfpShieldTime *float32 `json:"otfp_shield_time,omitempty"`
	// Total amount of time spent delivering a response from the Fastly On-the-Fly Packaging service for video-on-demand (in seconds).
	OtfpDeliverTime *float32 `json:"otfp_deliver_time,omitempty"`
	// Total header bytes delivered from the Fastly Image Optimizer service, including shield traffic.
	ImgoptoRespHeaderBytes *int64 `json:"imgopto_resp_header_bytes,omitempty"`
	// Total body bytes delivered from the Fastly Image Optimizer service, including shield traffic.
	ImgoptoRespBodyBytes *int64 `json:"imgopto_resp_body_bytes,omitempty"`
	// Total header bytes delivered via a shield from the Fastly Image Optimizer service.
	ImgoptoShieldRespHeaderBytes *int64 `json:"imgopto_shield_resp_header_bytes,omitempty"`
	// Total body bytes delivered via a shield from the Fastly Image Optimizer service.
	ImgoptoShieldRespBodyBytes *int64 `json:"imgopto_shield_resp_body_bytes,omitempty"`
	// Number of \"Informational\" category status codes delivered.
	Status1xx *int64 `json:"status_1xx,omitempty"`
	// Number of \"Success\" status codes delivered.
	Status2xx *int64 `json:"status_2xx,omitempty"`
	// Number of \"Redirection\" codes delivered.
	Status3xx *int64 `json:"status_3xx,omitempty"`
	// Number of \"Client Error\" codes delivered.
	Status4xx *int64 `json:"status_4xx,omitempty"`
	// Number of \"Server Error\" codes delivered.
	Status5xx *int64 `json:"status_5xx,omitempty"`
	// Number of responses sent with status code 200 (Success).
	Status200 *int64 `json:"status_200,omitempty"`
	// Number of responses sent with status code 204 (No Content).
	Status204 *int64 `json:"status_204,omitempty"`
	// Number of responses sent with status code 206 (Partial Content).
	Status206 *int64 `json:"status_206,omitempty"`
	// Number of responses sent with status code 301 (Moved Permanently).
	Status301 *int64 `json:"status_301,omitempty"`
	// Number of responses sent with status code 302 (Found).
	Status302 *int64 `json:"status_302,omitempty"`
	// Number of responses sent with status code 304 (Not Modified).
	Status304 *int64 `json:"status_304,omitempty"`
	// Number of responses sent with status code 400 (Bad Request).
	Status400 *int64 `json:"status_400,omitempty"`
	// Number of responses sent with status code 401 (Unauthorized).
	Status401 *int64 `json:"status_401,omitempty"`
	// Number of responses sent with status code 403 (Forbidden).
	Status403 *int64 `json:"status_403,omitempty"`
	// Number of responses sent with status code 404 (Not Found).
	Status404 *int64 `json:"status_404,omitempty"`
	// Number of responses sent with status code 406 (Not Acceptable).
	Status406 *int64 `json:"status_406,omitempty"`
	// Number of responses sent with status code 416 (Range Not Satisfiable).
	Status416 *int64 `json:"status_416,omitempty"`
	// Number of responses sent with status code 429 (Too Many Requests).
	Status429 *int64 `json:"status_429,omitempty"`
	// Number of responses sent with status code 500 (Internal Server Error).
	Status500 *int64 `json:"status_500,omitempty"`
	// Number of responses sent with status code 501 (Not Implemented).
	Status501 *int64 `json:"status_501,omitempty"`
	// Number of responses sent with status code 502 (Bad Gateway).
	Status502 *int64 `json:"status_502,omitempty"`
	// Number of responses sent with status code 503 (Service Unavailable).
	Status503 *int64 `json:"status_503,omitempty"`
	// Number of responses sent with status code 504 (Gateway Timeout).
	Status504 *int64 `json:"status_504,omitempty"`
	// Number of responses sent with status code 505 (HTTP Version Not Supported).
	Status505 *int64 `json:"status_505,omitempty"`
	// Number of requests that were designated uncachable.
	Uncacheable *int64 `json:"uncacheable,omitempty"`
	// Total amount of time spent processing cache passes (in seconds).
	PassTime *float32 `json:"pass_time,omitempty"`
	// Number of requests that were received over TLS.
	TLS *int64 `json:"tls,omitempty"`
	// Number of requests received over TLS 1.0.
	TLSV10 *int64 `json:"tls_v10,omitempty"`
	// Number of requests received over TLS 1.1.
	TLSV11 *int64 `json:"tls_v11,omitempty"`
	// Number of requests received over TLS 1.2.
	TLSV12 *int64 `json:"tls_v12,omitempty"`
	// Number of requests received over TLS 1.3.
	TLSV13 *int64 `json:"tls_v13,omitempty"`
	// Number of objects served that were under 1KB in size.
	ObjectSize1k *int64 `json:"object_size_1k,omitempty"`
	// Number of objects served that were between 1KB and 10KB in size.
	ObjectSize10k *int64 `json:"object_size_10k,omitempty"`
	// Number of objects served that were between 10KB and 100KB in size.
	ObjectSize100k *int64 `json:"object_size_100k,omitempty"`
	// Number of objects served that were between 100KB and 1MB in size.
	ObjectSize1m *int64 `json:"object_size_1m,omitempty"`
	// Number of objects served that were between 1MB and 10MB in size.
	ObjectSize10m *int64 `json:"object_size_10m,omitempty"`
	// Number of objects served that were between 10MB and 100MB in size.
	ObjectSize100m *int64 `json:"object_size_100m,omitempty"`
	// Number of objects served that were between 100MB and 1GB in size.
	ObjectSize1g *int64 `json:"object_size_1g,omitempty"`
	// Number of objects served that were larger than 1GB in size.
	ObjectSizeOther *int64 `json:"object_size_other,omitempty"`
	// Time spent inside the `vcl_recv` Varnish subroutine (in nanoseconds).
	RecvSubTime *float32 `json:"recv_sub_time,omitempty"`
	// Number of executions of the `vcl_recv` Varnish subroutine.
	RecvSubCount *int64 `json:"recv_sub_count,omitempty"`
	// Time spent inside the `vcl_hash` Varnish subroutine (in nanoseconds).
	HashSubTime *float32 `json:"hash_sub_time,omitempty"`
	// Number of executions of the `vcl_hash` Varnish subroutine.
	HashSubCount *int64 `json:"hash_sub_count,omitempty"`
	// Time spent inside the `vcl_miss` Varnish subroutine (in nanoseconds).
	MissSubTime *float32 `json:"miss_sub_time,omitempty"`
	// Number of executions of the `vcl_miss` Varnish subroutine.
	MissSubCount *int64 `json:"miss_sub_count,omitempty"`
	// Time spent inside the `vcl_fetch` Varnish subroutine (in nanoseconds).
	FetchSubTime *float32 `json:"fetch_sub_time,omitempty"`
	// Number of executions of the `vcl_fetch` Varnish subroutine.
	FetchSubCount *int64 `json:"fetch_sub_count,omitempty"`
	// Time spent inside the `vcl_pass` Varnish subroutine (in nanoseconds).
	PassSubTime *float32 `json:"pass_sub_time,omitempty"`
	// Number of executions of the `vcl_pass` Varnish subroutine.
	PassSubCount *int64 `json:"pass_sub_count,omitempty"`
	// Time spent inside the `vcl_pipe` Varnish subroutine (in nanoseconds).
	PipeSubTime *float32 `json:"pipe_sub_time,omitempty"`
	// Number of executions of the `vcl_pipe` Varnish subroutine.
	PipeSubCount *int64 `json:"pipe_sub_count,omitempty"`
	// Time spent inside the `vcl_deliver` Varnish subroutine (in nanoseconds).
	DeliverSubTime *float32 `json:"deliver_sub_time,omitempty"`
	// Number of executions of the `vcl_deliver` Varnish subroutine.
	DeliverSubCount *int64 `json:"deliver_sub_count,omitempty"`
	// Time spent inside the `vcl_error` Varnish subroutine (in nanoseconds).
	ErrorSubTime *float32 `json:"error_sub_time,omitempty"`
	// Number of executions of the `vcl_error` Varnish subroutine.
	ErrorSubCount *int64 `json:"error_sub_count,omitempty"`
	// Time spent inside the `vcl_hit` Varnish subroutine (in nanoseconds).
	HitSubTime *float32 `json:"hit_sub_time,omitempty"`
	// Number of executions of the `vcl_hit` Varnish subroutine.
	HitSubCount *int64 `json:"hit_sub_count,omitempty"`
	// Time spent inside the `vcl_prehash` Varnish subroutine (in nanoseconds).
	PrehashSubTime *float32 `json:"prehash_sub_time,omitempty"`
	// Number of executions of the `vcl_prehash` Varnish subroutine.
	PrehashSubCount *int64 `json:"prehash_sub_count,omitempty"`
	// Time spent inside the `vcl_predeliver` Varnish subroutine (in nanoseconds).
	PredeliverSubTime *float32 `json:"predeliver_sub_time,omitempty"`
	// Number of executions of the `vcl_predeliver` Varnish subroutine.
	PredeliverSubCount *int64 `json:"predeliver_sub_count,omitempty"`
	// Total body bytes delivered for cache hits.
	HitRespBodyBytes *int64 `json:"hit_resp_body_bytes,omitempty"`
	// Total body bytes delivered for cache misses.
	MissRespBodyBytes *int64 `json:"miss_resp_body_bytes,omitempty"`
	// Total body bytes delivered for cache passes.
	PassRespBodyBytes *int64 `json:"pass_resp_body_bytes,omitempty"`
	// Total header bytes received by the Compute platform.
	ComputeReqHeaderBytes *int64 `json:"compute_req_header_bytes,omitempty"`
	// Total body bytes received by the Compute platform.
	ComputeReqBodyBytes *int64 `json:"compute_req_body_bytes,omitempty"`
	// Total header bytes sent from Compute to end user.
	ComputeRespHeaderBytes *int64 `json:"compute_resp_header_bytes,omitempty"`
	// Total body bytes sent from Compute to end user.
	ComputeRespBodyBytes *int64 `json:"compute_resp_body_bytes,omitempty"`
	// Number of video responses that came from the Fastly Image Optimizer service.
	Imgvideo *int64 `json:"imgvideo,omitempty"`
	// Number of video frames that came from the Fastly Image Optimizer service. A video frame is an individual image within a sequence of video.
	ImgvideoFrames *int64 `json:"imgvideo_frames,omitempty"`
	// Total header bytes of video delivered from the Fastly Image Optimizer service.
	ImgvideoRespHeaderBytes *int64 `json:"imgvideo_resp_header_bytes,omitempty"`
	// Total body bytes of video delivered from the Fastly Image Optimizer service.
	ImgvideoRespBodyBytes *int64 `json:"imgvideo_resp_body_bytes,omitempty"`
	// Number of video responses delivered via a shield that came from the Fastly Image Optimizer service.
	ImgvideoShield *int64 `json:"imgvideo_shield,omitempty"`
	// Number of video frames delivered via a shield that came from the Fastly Image Optimizer service. A video frame is an individual image within a sequence of video.
	ImgvideoShieldFrames *int64 `json:"imgvideo_shield_frames,omitempty"`
	// Total header bytes of video delivered via a shield from the Fastly Image Optimizer service.
	ImgvideoShieldRespHeaderBytes *int64 `json:"imgvideo_shield_resp_header_bytes,omitempty"`
	// Total body bytes of video delivered via a shield from the Fastly Image Optimizer service.
	ImgvideoShieldRespBodyBytes *int64 `json:"imgvideo_shield_resp_body_bytes,omitempty"`
	// Total log bytes sent.
	LogBytes *int64 `json:"log_bytes,omitempty"`
	// Number of requests sent by end users to Fastly.
	EdgeRequests *int64 `json:"edge_requests,omitempty"`
	// Total header bytes delivered from Fastly to the end user.
	EdgeRespHeaderBytes *int64 `json:"edge_resp_header_bytes,omitempty"`
	// Total body bytes delivered from Fastly to the end user.
	EdgeRespBodyBytes *int64 `json:"edge_resp_body_bytes,omitempty"`
	// Number of responses received from origin with a `304` status code in response to an `If-Modified-Since` or `If-None-Match` request. Under regular scenarios, a revalidation will imply a cache hit. However, if using Fastly Image Optimizer or segmented caching this may result in a cache miss.
	OriginRevalidations *int64 `json:"origin_revalidations,omitempty"`
	// Number of requests sent to origin.
	OriginFetches *int64 `json:"origin_fetches,omitempty"`
	// Total request header bytes sent to origin.
	OriginFetchHeaderBytes *int64 `json:"origin_fetch_header_bytes,omitempty"`
	// Total request body bytes sent to origin.
	OriginFetchBodyBytes *int64 `json:"origin_fetch_body_bytes,omitempty"`
	// Total header bytes received from origin.
	OriginFetchRespHeaderBytes *int64 `json:"origin_fetch_resp_header_bytes,omitempty"`
	// Total body bytes received from origin.
	OriginFetchRespBodyBytes *int64 `json:"origin_fetch_resp_body_bytes,omitempty"`
	// Number of responses received from origin with a `304` status code, in response to an `If-Modified-Since` or `If-None-Match` request to a shield. Under regular scenarios, a revalidation will imply a cache hit. However, if using segmented caching this may result in a cache miss.
	ShieldRevalidations *int64 `json:"shield_revalidations,omitempty"`
	// Number of requests made from one Fastly POP to another, as part of shielding.
	ShieldFetches *int64 `json:"shield_fetches,omitempty"`
	// Total request header bytes sent to a shield.
	ShieldFetchHeaderBytes *int64 `json:"shield_fetch_header_bytes,omitempty"`
	// Total request body bytes sent to a shield.
	ShieldFetchBodyBytes *int64 `json:"shield_fetch_body_bytes,omitempty"`
	// Total response header bytes sent from a shield to the edge.
	ShieldFetchRespHeaderBytes *int64 `json:"shield_fetch_resp_header_bytes,omitempty"`
	// Total response body bytes sent from a shield to the edge.
	ShieldFetchRespBodyBytes *int64 `json:"shield_fetch_resp_body_bytes,omitempty"`
	// Number of `Range` requests to origin for segments of resources when using segmented caching.
	SegblockOriginFetches *int64 `json:"segblock_origin_fetches,omitempty"`
	// Number of `Range` requests to a shield for segments of resources when using segmented caching.
	SegblockShieldFetches *int64 `json:"segblock_shield_fetches,omitempty"`
	// Number of \"Informational\" category status codes delivered by the Compute platform.
	ComputeRespStatus1xx *int64 `json:"compute_resp_status_1xx,omitempty"`
	// Number of \"Success\" category status codes delivered by the Compute platform.
	ComputeRespStatus2xx *int64 `json:"compute_resp_status_2xx,omitempty"`
	// Number of \"Redirection\" category status codes delivered by the Compute platform.
	ComputeRespStatus3xx *int64 `json:"compute_resp_status_3xx,omitempty"`
	// Number of \"Client Error\" category status codes delivered by the Compute platform.
	ComputeRespStatus4xx *int64 `json:"compute_resp_status_4xx,omitempty"`
	// Number of \"Server Error\" category status codes delivered by the Compute platform.
	ComputeRespStatus5xx *int64 `json:"compute_resp_status_5xx,omitempty"`
	// Number of requests sent by end users to Fastly that resulted in a hit at the edge.
	EdgeHitRequests *int64 `json:"edge_hit_requests,omitempty"`
	// Number of requests sent by end users to Fastly that resulted in a miss at the edge.
	EdgeMissRequests *int64 `json:"edge_miss_requests,omitempty"`
	// Total header bytes sent to backends (origins) by the Compute platform.
	ComputeBereqHeaderBytes *int64 `json:"compute_bereq_header_bytes,omitempty"`
	// Total body bytes sent to backends (origins) by the Compute platform.
	ComputeBereqBodyBytes *int64 `json:"compute_bereq_body_bytes,omitempty"`
	// Total header bytes received from backends (origins) by the Compute platform.
	ComputeBerespHeaderBytes *int64 `json:"compute_beresp_header_bytes,omitempty"`
	// Total body bytes received from backends (origins) by the Compute platform.
	ComputeBerespBodyBytes *int64 `json:"compute_beresp_body_bytes,omitempty"`
	// The total number of completed requests made to backends (origins) that returned cacheable content.
	OriginCacheFetches *int64 `json:"origin_cache_fetches,omitempty"`
	// The total number of completed requests made to shields that returned cacheable content.
	ShieldCacheFetches *int64 `json:"shield_cache_fetches,omitempty"`
	// Number of backend requests started.
	ComputeBereqs *int64 `json:"compute_bereqs,omitempty"`
	// Number of backend request errors, including timeouts.
	ComputeBereqErrors *int64 `json:"compute_bereq_errors,omitempty"`
	// Number of times a guest exceeded its resource limit, includes heap, stack, globals, and code execution timeout.
	ComputeResourceLimitExceeded *int64 `json:"compute_resource_limit_exceeded,omitempty"`
	// Number of times a guest exceeded its heap limit.
	ComputeHeapLimitExceeded *int64 `json:"compute_heap_limit_exceeded,omitempty"`
	// Number of times a guest exceeded its stack limit.
	ComputeStackLimitExceeded *int64 `json:"compute_stack_limit_exceeded,omitempty"`
	// Number of times a guest exceeded its globals limit.
	ComputeGlobalsLimitExceeded *int64 `json:"compute_globals_limit_exceeded,omitempty"`
	// Number of times a service experienced a guest code error.
	ComputeGuestErrors *int64 `json:"compute_guest_errors,omitempty"`
	// Number of times a service experienced a guest runtime error.
	ComputeRuntimeErrors *int64 `json:"compute_runtime_errors,omitempty"`
	// Body bytes delivered for edge hits.
	EdgeHitRespBodyBytes *int64 `json:"edge_hit_resp_body_bytes,omitempty"`
	// Header bytes delivered for edge hits.
	EdgeHitRespHeaderBytes *int64 `json:"edge_hit_resp_header_bytes,omitempty"`
	// Body bytes delivered for edge misses.
	EdgeMissRespBodyBytes *int64 `json:"edge_miss_resp_body_bytes,omitempty"`
	// Header bytes delivered for edge misses.
	EdgeMissRespHeaderBytes *int64 `json:"edge_miss_resp_header_bytes,omitempty"`
	// Body bytes received from origin for cacheable content.
	OriginCacheFetchRespBodyBytes *int64 `json:"origin_cache_fetch_resp_body_bytes,omitempty"`
	// Header bytes received from an origin for cacheable content.
	OriginCacheFetchRespHeaderBytes *int64 `json:"origin_cache_fetch_resp_header_bytes,omitempty"`
	// Number of requests that resulted in a hit at a shield.
	ShieldHitRequests *int64 `json:"shield_hit_requests,omitempty"`
	// Number of requests that resulted in a miss at a shield.
	ShieldMissRequests *int64 `json:"shield_miss_requests,omitempty"`
	// Header bytes delivered for shield hits.
	ShieldHitRespHeaderBytes *int64 `json:"shield_hit_resp_header_bytes,omitempty"`
	// Body bytes delivered for shield hits.
	ShieldHitRespBodyBytes *int64 `json:"shield_hit_resp_body_bytes,omitempty"`
	// Header bytes delivered for shield misses.
	ShieldMissRespHeaderBytes *int64 `json:"shield_miss_resp_header_bytes,omitempty"`
	// Body bytes delivered for shield misses.
	ShieldMissRespBodyBytes *int64 `json:"shield_miss_resp_body_bytes,omitempty"`
	// Total header bytes received from end users over passthrough WebSocket connections.
	WebsocketReqHeaderBytes *int64 `json:"websocket_req_header_bytes,omitempty"`
	// Total message content bytes received from end users over passthrough WebSocket connections.
	WebsocketReqBodyBytes *int64 `json:"websocket_req_body_bytes,omitempty"`
	// Total header bytes sent to end users over passthrough WebSocket connections.
	WebsocketRespHeaderBytes *int64 `json:"websocket_resp_header_bytes,omitempty"`
	// Total header bytes sent to backends over passthrough WebSocket connections.
	WebsocketBereqHeaderBytes *int64 `json:"websocket_bereq_header_bytes,omitempty"`
	// Total message content bytes sent to backends over passthrough WebSocket connections.
	WebsocketBereqBodyBytes *int64 `json:"websocket_bereq_body_bytes,omitempty"`
	// Total header bytes received from backends over passthrough WebSocket connections.
	WebsocketBerespHeaderBytes *int64 `json:"websocket_beresp_header_bytes,omitempty"`
	// Total message content bytes received from backends over passthrough WebSocket connections.
	WebsocketBerespBodyBytes *int64 `json:"websocket_beresp_body_bytes,omitempty"`
	// Total duration of passthrough WebSocket connections with end users.
	WebsocketConnTimeMs *int64 `json:"websocket_conn_time_ms,omitempty"`
	// Total message content bytes sent to end users over passthrough WebSocket connections.
	WebsocketRespBodyBytes *int64 `json:"websocket_resp_body_bytes,omitempty"`
	// Total published messages received from the publish API endpoint.
	FanoutRecvPublishes *int64 `json:"fanout_recv_publishes,omitempty"`
	// Total published messages sent to end users.
	FanoutSendPublishes *int64 `json:"fanout_send_publishes,omitempty"`
	// The total number of class a operations for the KV store.
	KvStoreClassAOperations *int64 `json:"kv_store_class_a_operations,omitempty"`
	// The total number of class b operations for the KV store.
	KvStoreClassBOperations *int64 `json:"kv_store_class_b_operations,omitempty"`
	// Use kv_store_class_a_operations.
	// Deprecated
	ObjectStoreClassAOperations *int64 `json:"object_store_class_a_operations,omitempty"`
	// Use kv_store_class_b_operations.
	// Deprecated
	ObjectStoreClassBOperations *int64 `json:"object_store_class_b_operations,omitempty"`
	// Total header bytes received from end users over Fanout connections.
	FanoutReqHeaderBytes *int64 `json:"fanout_req_header_bytes,omitempty"`
	// Total body or message content bytes received from end users over Fanout connections.
	FanoutReqBodyBytes *int64 `json:"fanout_req_body_bytes,omitempty"`
	// Total header bytes sent to end users over Fanout connections.
	FanoutRespHeaderBytes *int64 `json:"fanout_resp_header_bytes,omitempty"`
	// Total body or message content bytes sent to end users over Fanout connections, excluding published message content.
	FanoutRespBodyBytes *int64 `json:"fanout_resp_body_bytes,omitempty"`
	// Total header bytes sent to backends over Fanout connections.
	FanoutBereqHeaderBytes *int64 `json:"fanout_bereq_header_bytes,omitempty"`
	// Total body or message content bytes sent to backends over Fanout connections.
	FanoutBereqBodyBytes *int64 `json:"fanout_bereq_body_bytes,omitempty"`
	// Total header bytes received from backends over Fanout connections.
	FanoutBerespHeaderBytes *int64 `json:"fanout_beresp_header_bytes,omitempty"`
	// Total body or message content bytes received from backends over Fanout connections.
	FanoutBerespBodyBytes *int64 `json:"fanout_beresp_body_bytes,omitempty"`
	// Total duration of Fanout connections with end users.
	FanoutConnTimeMs *int64 `json:"fanout_conn_time_ms,omitempty"`
	// For HTTP/2, the number of connections the limit-streams action was applied to. The limit-streams action caps the allowed number of concurrent streams in a connection.
	DdosActionLimitStreamsConnections *int64 `json:"ddos_action_limit_streams_connections,omitempty"`
	// For HTTP/2, the number of requests made on a connection for which the limit-streams action was taken. The limit-streams action caps the allowed number of concurrent streams in a connection.
	DdosActionLimitStreamsRequests *int64 `json:"ddos_action_limit_streams_requests,omitempty"`
	// The number of times the tarpit-accept action was taken. The tarpit-accept action adds a delay when accepting future connections.
	DdosActionTarpitAccept *int64 `json:"ddos_action_tarpit_accept,omitempty"`
	// The number of times the tarpit action was taken. The tarpit action delays writing the response to the client.
	DdosActionTarpit *int64 `json:"ddos_action_tarpit,omitempty"`
	// The number of times the close action was taken. The close action aborts the connection as soon as possible. The close action takes effect either right after accept, right after the client hello, or right after the response was sent.
	DdosActionClose *int64 `json:"ddos_action_close,omitempty"`
	// The number of times the blackhole action was taken. The blackhole action quietly closes a TCP connection without sending a reset. The blackhole action quietly closes a TCP connection without notifying its peer (all TCP state is dropped).
	DdosActionBlackhole *int64 `json:"ddos_action_blackhole,omitempty"`
	// The number of challenge-start tokens created.
	BotChallengeStarts *int64 `json:"bot_challenge_starts,omitempty"`
	// The number of challenge-complete tokens that passed validation.
	BotChallengeCompleteTokensPassed *int64 `json:"bot_challenge_complete_tokens_passed,omitempty"`
	// The number of challenge-complete tokens that failed validation.
	BotChallengeCompleteTokensFailed *int64 `json:"bot_challenge_complete_tokens_failed,omitempty"`
	// The number of challenge-complete tokens checked.
	BotChallengeCompleteTokensChecked *int64 `json:"bot_challenge_complete_tokens_checked,omitempty"`
	// The number of challenge-complete tokens not checked because the feature was disabled.
	BotChallengeCompleteTokensDisabled *int64 `json:"bot_challenge_complete_tokens_disabled,omitempty"`
	// The number of challenges issued. For example, the issuance of a CAPTCHA challenge.
	BotChallengesIssued *int64 `json:"bot_challenges_issued,omitempty"`
	// The number of successful challenge solutions processed. For example, a correct CAPTCHA solution.
	BotChallengesSucceeded *int64 `json:"bot_challenges_succeeded,omitempty"`
	// The number of failed challenge solutions processed. For example, an incorrect CAPTCHA solution.
	BotChallengesFailed *int64 `json:"bot_challenges_failed,omitempty"`
	// The number of challenge-complete tokens issued. For example, issuing a challenge-complete token after a series of CAPTCHA challenges ending in success.
	BotChallengeCompleteTokensIssued *int64 `json:"bot_challenge_complete_tokens_issued,omitempty"`
	// The number of times the downgrade action was taken. The downgrade action restricts the client to http1.
	DdosActionDowngrade *int64 `json:"ddos_action_downgrade,omitempty"`
	// The number of connections the downgrade action was applied to. The downgrade action restricts the connection to http1.
	DdosActionDowngradedConnections *int64 `json:"ddos_action_downgraded_connections,omitempty"`
	// Number of cache hits for a VCL service.
	AllHitRequests *int64 `json:"all_hit_requests,omitempty"`
	// Number of cache misses for a VCL service.
	AllMissRequests *int64 `json:"all_miss_requests,omitempty"`
	// Number of requests that passed through the CDN without being cached for a VCL service.
	AllPassRequests *int64 `json:"all_pass_requests,omitempty"`
	// Number of cache errors for a VCL service.
	AllErrorRequests *int64 `json:"all_error_requests,omitempty"`
	// Number of requests that returned a synthetic response (i.e., response objects created with the `synthetic` VCL statement) for a VCL service.
	AllSynthRequests *int64 `json:"all_synth_requests,omitempty"`
	// Number of requests sent by end users to Fastly that resulted in a hit at the edge for a VCL service.
	AllEdgeHitRequests *int64 `json:"all_edge_hit_requests,omitempty"`
	// Number of requests sent by end users to Fastly that resulted in a miss at the edge for a VCL service.
	AllEdgeMissRequests *int64 `json:"all_edge_miss_requests,omitempty"`
	// Number of \"Informational\" category status codes delivered for all sources.
	AllStatus1xx *int64 `json:"all_status_1xx,omitempty"`
	// Number of \"Success\" status codes delivered for all sources.
	AllStatus2xx *int64 `json:"all_status_2xx,omitempty"`
	// Number of \"Redirection\" codes delivered for all sources.
	AllStatus3xx *int64 `json:"all_status_3xx,omitempty"`
	// Number of \"Client Error\" codes delivered for all sources.
	AllStatus4xx *int64 `json:"all_status_4xx,omitempty"`
	// Number of \"Server Error\" codes delivered for all sources.
	AllStatus5xx *int64 `json:"all_status_5xx,omitempty"`
	// Origin Offload measures the ratio of bytes served to end users that were cached by Fastly, over the bytes served to end users, between 0 and 1. ((`edge_resp_body_bytes` + `edge_resp_header_bytes`) - (`origin_fetch_resp_body_bytes` + `origin_fetch_resp_header_bytes`)) / (`edge_resp_body_bytes` + `edge_resp_header_bytes`).
	OriginOffload *float32 `json:"origin_offload,omitempty"`
	// Number of requests where Fastly responded with 400 due to the request being a GET or HEAD request containing a body.
	RequestDeniedGetHeadBody *int64 `json:"request_denied_get_head_body,omitempty"`
	// Number of requests classified as a DDoS attack against a customer origin or service.
	ServiceDdosRequestsDetected *int64 `json:"service_ddos_requests_detected,omitempty"`
	// Number of requests classified as a DDoS attack against a customer origin or service that were mitigated by the Fastly platform.
	ServiceDdosRequestsMitigated *int64 `json:"service_ddos_requests_mitigated,omitempty"`
	// Number of requests analyzed for DDoS attacks against a customer origin or service, but with no DDoS detected.
	ServiceDdosRequestsAllowed *int64 `json:"service_ddos_requests_allowed,omitempty"`
	AdditionalProperties       map[string]any
}

type _RealtimeEntryAggregated RealtimeEntryAggregated

// NewRealtimeEntryAggregated instantiates a new RealtimeEntryAggregated object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRealtimeEntryAggregated() *RealtimeEntryAggregated {
	this := RealtimeEntryAggregated{}
	return &this
}

// NewRealtimeEntryAggregatedWithDefaults instantiates a new RealtimeEntryAggregated object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRealtimeEntryAggregatedWithDefaults() *RealtimeEntryAggregated {
	this := RealtimeEntryAggregated{}
	return &this
}

// GetRequests returns the Requests field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetRequests() int64 {
	if o == nil || o.Requests == nil {
		var ret int64
		return ret
	}
	return *o.Requests
}

// GetRequestsOk returns a tuple with the Requests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetRequestsOk() (*int64, bool) {
	if o == nil || o.Requests == nil {
		return nil, false
	}
	return o.Requests, true
}

// HasRequests returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasRequests() bool {
	if o != nil && o.Requests != nil {
		return true
	}

	return false
}

// SetRequests gets a reference to the given int64 and assigns it to the Requests field.
func (o *RealtimeEntryAggregated) SetRequests(v int64) {
	o.Requests = &v
}

// GetLogging returns the Logging field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetLogging() int64 {
	if o == nil || o.Logging == nil {
		var ret int64
		return ret
	}
	return *o.Logging
}

// GetLoggingOk returns a tuple with the Logging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetLoggingOk() (*int64, bool) {
	if o == nil || o.Logging == nil {
		return nil, false
	}
	return o.Logging, true
}

// HasLogging returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasLogging() bool {
	if o != nil && o.Logging != nil {
		return true
	}

	return false
}

// SetLogging gets a reference to the given int64 and assigns it to the Logging field.
func (o *RealtimeEntryAggregated) SetLogging(v int64) {
	o.Logging = &v
}

// GetLog returns the Log field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetLog() int64 {
	if o == nil || o.Log == nil {
		var ret int64
		return ret
	}
	return *o.Log
}

// GetLogOk returns a tuple with the Log field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetLogOk() (*int64, bool) {
	if o == nil || o.Log == nil {
		return nil, false
	}
	return o.Log, true
}

// HasLog returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasLog() bool {
	if o != nil && o.Log != nil {
		return true
	}

	return false
}

// SetLog gets a reference to the given int64 and assigns it to the Log field.
func (o *RealtimeEntryAggregated) SetLog(v int64) {
	o.Log = &v
}

// GetRespHeaderBytes returns the RespHeaderBytes field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetRespHeaderBytes() int64 {
	if o == nil || o.RespHeaderBytes == nil {
		var ret int64
		return ret
	}
	return *o.RespHeaderBytes
}

// GetRespHeaderBytesOk returns a tuple with the RespHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetRespHeaderBytesOk() (*int64, bool) {
	if o == nil || o.RespHeaderBytes == nil {
		return nil, false
	}
	return o.RespHeaderBytes, true
}

// HasRespHeaderBytes returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasRespHeaderBytes() bool {
	if o != nil && o.RespHeaderBytes != nil {
		return true
	}

	return false
}

// SetRespHeaderBytes gets a reference to the given int64 and assigns it to the RespHeaderBytes field.
func (o *RealtimeEntryAggregated) SetRespHeaderBytes(v int64) {
	o.RespHeaderBytes = &v
}

// GetHeaderSize returns the HeaderSize field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetHeaderSize() int64 {
	if o == nil || o.HeaderSize == nil {
		var ret int64
		return ret
	}
	return *o.HeaderSize
}

// GetHeaderSizeOk returns a tuple with the HeaderSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetHeaderSizeOk() (*int64, bool) {
	if o == nil || o.HeaderSize == nil {
		return nil, false
	}
	return o.HeaderSize, true
}

// HasHeaderSize returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasHeaderSize() bool {
	if o != nil && o.HeaderSize != nil {
		return true
	}

	return false
}

// SetHeaderSize gets a reference to the given int64 and assigns it to the HeaderSize field.
func (o *RealtimeEntryAggregated) SetHeaderSize(v int64) {
	o.HeaderSize = &v
}

// GetRespBodyBytes returns the RespBodyBytes field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetRespBodyBytes() int64 {
	if o == nil || o.RespBodyBytes == nil {
		var ret int64
		return ret
	}
	return *o.RespBodyBytes
}

// GetRespBodyBytesOk returns a tuple with the RespBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetRespBodyBytesOk() (*int64, bool) {
	if o == nil || o.RespBodyBytes == nil {
		return nil, false
	}
	return o.RespBodyBytes, true
}

// HasRespBodyBytes returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasRespBodyBytes() bool {
	if o != nil && o.RespBodyBytes != nil {
		return true
	}

	return false
}

// SetRespBodyBytes gets a reference to the given int64 and assigns it to the RespBodyBytes field.
func (o *RealtimeEntryAggregated) SetRespBodyBytes(v int64) {
	o.RespBodyBytes = &v
}

// GetBodySize returns the BodySize field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetBodySize() int64 {
	if o == nil || o.BodySize == nil {
		var ret int64
		return ret
	}
	return *o.BodySize
}

// GetBodySizeOk returns a tuple with the BodySize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetBodySizeOk() (*int64, bool) {
	if o == nil || o.BodySize == nil {
		return nil, false
	}
	return o.BodySize, true
}

// HasBodySize returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasBodySize() bool {
	if o != nil && o.BodySize != nil {
		return true
	}

	return false
}

// SetBodySize gets a reference to the given int64 and assigns it to the BodySize field.
func (o *RealtimeEntryAggregated) SetBodySize(v int64) {
	o.BodySize = &v
}

// GetHits returns the Hits field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetHits() int64 {
	if o == nil || o.Hits == nil {
		var ret int64
		return ret
	}
	return *o.Hits
}

// GetHitsOk returns a tuple with the Hits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetHitsOk() (*int64, bool) {
	if o == nil || o.Hits == nil {
		return nil, false
	}
	return o.Hits, true
}

// HasHits returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasHits() bool {
	if o != nil && o.Hits != nil {
		return true
	}

	return false
}

// SetHits gets a reference to the given int64 and assigns it to the Hits field.
func (o *RealtimeEntryAggregated) SetHits(v int64) {
	o.Hits = &v
}

// GetMiss returns the Miss field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetMiss() int64 {
	if o == nil || o.Miss == nil {
		var ret int64
		return ret
	}
	return *o.Miss
}

// GetMissOk returns a tuple with the Miss field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetMissOk() (*int64, bool) {
	if o == nil || o.Miss == nil {
		return nil, false
	}
	return o.Miss, true
}

// HasMiss returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasMiss() bool {
	if o != nil && o.Miss != nil {
		return true
	}

	return false
}

// SetMiss gets a reference to the given int64 and assigns it to the Miss field.
func (o *RealtimeEntryAggregated) SetMiss(v int64) {
	o.Miss = &v
}

// GetPass returns the Pass field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetPass() int64 {
	if o == nil || o.Pass == nil {
		var ret int64
		return ret
	}
	return *o.Pass
}

// GetPassOk returns a tuple with the Pass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetPassOk() (*int64, bool) {
	if o == nil || o.Pass == nil {
		return nil, false
	}
	return o.Pass, true
}

// HasPass returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasPass() bool {
	if o != nil && o.Pass != nil {
		return true
	}

	return false
}

// SetPass gets a reference to the given int64 and assigns it to the Pass field.
func (o *RealtimeEntryAggregated) SetPass(v int64) {
	o.Pass = &v
}

// GetSynth returns the Synth field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetSynth() int64 {
	if o == nil || o.Synth == nil {
		var ret int64
		return ret
	}
	return *o.Synth
}

// GetSynthOk returns a tuple with the Synth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetSynthOk() (*int64, bool) {
	if o == nil || o.Synth == nil {
		return nil, false
	}
	return o.Synth, true
}

// HasSynth returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasSynth() bool {
	if o != nil && o.Synth != nil {
		return true
	}

	return false
}

// SetSynth gets a reference to the given int64 and assigns it to the Synth field.
func (o *RealtimeEntryAggregated) SetSynth(v int64) {
	o.Synth = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetErrors() int64 {
	if o == nil || o.Errors == nil {
		var ret int64
		return ret
	}
	return *o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetErrorsOk() (*int64, bool) {
	if o == nil || o.Errors == nil {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasErrors() bool {
	if o != nil && o.Errors != nil {
		return true
	}

	return false
}

// SetErrors gets a reference to the given int64 and assigns it to the Errors field.
func (o *RealtimeEntryAggregated) SetErrors(v int64) {
	o.Errors = &v
}

// GetHitsTime returns the HitsTime field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetHitsTime() float32 {
	if o == nil || o.HitsTime == nil {
		var ret float32
		return ret
	}
	return *o.HitsTime
}

// GetHitsTimeOk returns a tuple with the HitsTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetHitsTimeOk() (*float32, bool) {
	if o == nil || o.HitsTime == nil {
		return nil, false
	}
	return o.HitsTime, true
}

// HasHitsTime returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasHitsTime() bool {
	if o != nil && o.HitsTime != nil {
		return true
	}

	return false
}

// SetHitsTime gets a reference to the given float32 and assigns it to the HitsTime field.
func (o *RealtimeEntryAggregated) SetHitsTime(v float32) {
	o.HitsTime = &v
}

// GetMissTime returns the MissTime field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetMissTime() float32 {
	if o == nil || o.MissTime == nil {
		var ret float32
		return ret
	}
	return *o.MissTime
}

// GetMissTimeOk returns a tuple with the MissTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetMissTimeOk() (*float32, bool) {
	if o == nil || o.MissTime == nil {
		return nil, false
	}
	return o.MissTime, true
}

// HasMissTime returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasMissTime() bool {
	if o != nil && o.MissTime != nil {
		return true
	}

	return false
}

// SetMissTime gets a reference to the given float32 and assigns it to the MissTime field.
func (o *RealtimeEntryAggregated) SetMissTime(v float32) {
	o.MissTime = &v
}

// GetMissHistogram returns the MissHistogram field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetMissHistogram() map[string]map[string]any {
	if o == nil || o.MissHistogram == nil {
		var ret map[string]map[string]any
		return ret
	}
	return o.MissHistogram
}

// GetMissHistogramOk returns a tuple with the MissHistogram field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetMissHistogramOk() (map[string]map[string]any, bool) {
	if o == nil || o.MissHistogram == nil {
		return nil, false
	}
	return o.MissHistogram, true
}

// HasMissHistogram returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasMissHistogram() bool {
	if o != nil && o.MissHistogram != nil {
		return true
	}

	return false
}

// SetMissHistogram gets a reference to the given map[string]map[string]any and assigns it to the MissHistogram field.
func (o *RealtimeEntryAggregated) SetMissHistogram(v map[string]map[string]any) {
	o.MissHistogram = v
}

// GetComputeRequests returns the ComputeRequests field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetComputeRequests() int64 {
	if o == nil || o.ComputeRequests == nil {
		var ret int64
		return ret
	}
	return *o.ComputeRequests
}

// GetComputeRequestsOk returns a tuple with the ComputeRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetComputeRequestsOk() (*int64, bool) {
	if o == nil || o.ComputeRequests == nil {
		return nil, false
	}
	return o.ComputeRequests, true
}

// HasComputeRequests returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasComputeRequests() bool {
	if o != nil && o.ComputeRequests != nil {
		return true
	}

	return false
}

// SetComputeRequests gets a reference to the given int64 and assigns it to the ComputeRequests field.
func (o *RealtimeEntryAggregated) SetComputeRequests(v int64) {
	o.ComputeRequests = &v
}

// GetComputeExecutionTimeMs returns the ComputeExecutionTimeMs field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetComputeExecutionTimeMs() float32 {
	if o == nil || o.ComputeExecutionTimeMs == nil {
		var ret float32
		return ret
	}
	return *o.ComputeExecutionTimeMs
}

// GetComputeExecutionTimeMsOk returns a tuple with the ComputeExecutionTimeMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetComputeExecutionTimeMsOk() (*float32, bool) {
	if o == nil || o.ComputeExecutionTimeMs == nil {
		return nil, false
	}
	return o.ComputeExecutionTimeMs, true
}

// HasComputeExecutionTimeMs returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasComputeExecutionTimeMs() bool {
	if o != nil && o.ComputeExecutionTimeMs != nil {
		return true
	}

	return false
}

// SetComputeExecutionTimeMs gets a reference to the given float32 and assigns it to the ComputeExecutionTimeMs field.
func (o *RealtimeEntryAggregated) SetComputeExecutionTimeMs(v float32) {
	o.ComputeExecutionTimeMs = &v
}

// GetComputeRAMUsed returns the ComputeRAMUsed field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetComputeRAMUsed() int64 {
	if o == nil || o.ComputeRAMUsed == nil {
		var ret int64
		return ret
	}
	return *o.ComputeRAMUsed
}

// GetComputeRAMUsedOk returns a tuple with the ComputeRAMUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetComputeRAMUsedOk() (*int64, bool) {
	if o == nil || o.ComputeRAMUsed == nil {
		return nil, false
	}
	return o.ComputeRAMUsed, true
}

// HasComputeRAMUsed returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasComputeRAMUsed() bool {
	if o != nil && o.ComputeRAMUsed != nil {
		return true
	}

	return false
}

// SetComputeRAMUsed gets a reference to the given int64 and assigns it to the ComputeRAMUsed field.
func (o *RealtimeEntryAggregated) SetComputeRAMUsed(v int64) {
	o.ComputeRAMUsed = &v
}

// GetComputeRequestTimeMs returns the ComputeRequestTimeMs field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetComputeRequestTimeMs() float32 {
	if o == nil || o.ComputeRequestTimeMs == nil {
		var ret float32
		return ret
	}
	return *o.ComputeRequestTimeMs
}

// GetComputeRequestTimeMsOk returns a tuple with the ComputeRequestTimeMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetComputeRequestTimeMsOk() (*float32, bool) {
	if o == nil || o.ComputeRequestTimeMs == nil {
		return nil, false
	}
	return o.ComputeRequestTimeMs, true
}

// HasComputeRequestTimeMs returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasComputeRequestTimeMs() bool {
	if o != nil && o.ComputeRequestTimeMs != nil {
		return true
	}

	return false
}

// SetComputeRequestTimeMs gets a reference to the given float32 and assigns it to the ComputeRequestTimeMs field.
func (o *RealtimeEntryAggregated) SetComputeRequestTimeMs(v float32) {
	o.ComputeRequestTimeMs = &v
}

// GetComputeRequestTimeBilledMs returns the ComputeRequestTimeBilledMs field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetComputeRequestTimeBilledMs() float32 {
	if o == nil || o.ComputeRequestTimeBilledMs == nil {
		var ret float32
		return ret
	}
	return *o.ComputeRequestTimeBilledMs
}

// GetComputeRequestTimeBilledMsOk returns a tuple with the ComputeRequestTimeBilledMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetComputeRequestTimeBilledMsOk() (*float32, bool) {
	if o == nil || o.ComputeRequestTimeBilledMs == nil {
		return nil, false
	}
	return o.ComputeRequestTimeBilledMs, true
}

// HasComputeRequestTimeBilledMs returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasComputeRequestTimeBilledMs() bool {
	if o != nil && o.ComputeRequestTimeBilledMs != nil {
		return true
	}

	return false
}

// SetComputeRequestTimeBilledMs gets a reference to the given float32 and assigns it to the ComputeRequestTimeBilledMs field.
func (o *RealtimeEntryAggregated) SetComputeRequestTimeBilledMs(v float32) {
	o.ComputeRequestTimeBilledMs = &v
}

// GetShield returns the Shield field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetShield() int64 {
	if o == nil || o.Shield == nil {
		var ret int64
		return ret
	}
	return *o.Shield
}

// GetShieldOk returns a tuple with the Shield field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetShieldOk() (*int64, bool) {
	if o == nil || o.Shield == nil {
		return nil, false
	}
	return o.Shield, true
}

// HasShield returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasShield() bool {
	if o != nil && o.Shield != nil {
		return true
	}

	return false
}

// SetShield gets a reference to the given int64 and assigns it to the Shield field.
func (o *RealtimeEntryAggregated) SetShield(v int64) {
	o.Shield = &v
}

// GetIpv6 returns the Ipv6 field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetIpv6() int64 {
	if o == nil || o.Ipv6 == nil {
		var ret int64
		return ret
	}
	return *o.Ipv6
}

// GetIpv6Ok returns a tuple with the Ipv6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetIpv6Ok() (*int64, bool) {
	if o == nil || o.Ipv6 == nil {
		return nil, false
	}
	return o.Ipv6, true
}

// HasIpv6 returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasIpv6() bool {
	if o != nil && o.Ipv6 != nil {
		return true
	}

	return false
}

// SetIpv6 gets a reference to the given int64 and assigns it to the Ipv6 field.
func (o *RealtimeEntryAggregated) SetIpv6(v int64) {
	o.Ipv6 = &v
}

// GetImgopto returns the Imgopto field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetImgopto() int64 {
	if o == nil || o.Imgopto == nil {
		var ret int64
		return ret
	}
	return *o.Imgopto
}

// GetImgoptoOk returns a tuple with the Imgopto field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetImgoptoOk() (*int64, bool) {
	if o == nil || o.Imgopto == nil {
		return nil, false
	}
	return o.Imgopto, true
}

// HasImgopto returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasImgopto() bool {
	if o != nil && o.Imgopto != nil {
		return true
	}

	return false
}

// SetImgopto gets a reference to the given int64 and assigns it to the Imgopto field.
func (o *RealtimeEntryAggregated) SetImgopto(v int64) {
	o.Imgopto = &v
}

// GetImgoptoShield returns the ImgoptoShield field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetImgoptoShield() int64 {
	if o == nil || o.ImgoptoShield == nil {
		var ret int64
		return ret
	}
	return *o.ImgoptoShield
}

// GetImgoptoShieldOk returns a tuple with the ImgoptoShield field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetImgoptoShieldOk() (*int64, bool) {
	if o == nil || o.ImgoptoShield == nil {
		return nil, false
	}
	return o.ImgoptoShield, true
}

// HasImgoptoShield returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasImgoptoShield() bool {
	if o != nil && o.ImgoptoShield != nil {
		return true
	}

	return false
}

// SetImgoptoShield gets a reference to the given int64 and assigns it to the ImgoptoShield field.
func (o *RealtimeEntryAggregated) SetImgoptoShield(v int64) {
	o.ImgoptoShield = &v
}

// GetImgoptoTransforms returns the ImgoptoTransforms field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetImgoptoTransforms() int64 {
	if o == nil || o.ImgoptoTransforms == nil {
		var ret int64
		return ret
	}
	return *o.ImgoptoTransforms
}

// GetImgoptoTransformsOk returns a tuple with the ImgoptoTransforms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetImgoptoTransformsOk() (*int64, bool) {
	if o == nil || o.ImgoptoTransforms == nil {
		return nil, false
	}
	return o.ImgoptoTransforms, true
}

// HasImgoptoTransforms returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasImgoptoTransforms() bool {
	if o != nil && o.ImgoptoTransforms != nil {
		return true
	}

	return false
}

// SetImgoptoTransforms gets a reference to the given int64 and assigns it to the ImgoptoTransforms field.
func (o *RealtimeEntryAggregated) SetImgoptoTransforms(v int64) {
	o.ImgoptoTransforms = &v
}

// GetOtfp returns the Otfp field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetOtfp() int64 {
	if o == nil || o.Otfp == nil {
		var ret int64
		return ret
	}
	return *o.Otfp
}

// GetOtfpOk returns a tuple with the Otfp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetOtfpOk() (*int64, bool) {
	if o == nil || o.Otfp == nil {
		return nil, false
	}
	return o.Otfp, true
}

// HasOtfp returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasOtfp() bool {
	if o != nil && o.Otfp != nil {
		return true
	}

	return false
}

// SetOtfp gets a reference to the given int64 and assigns it to the Otfp field.
func (o *RealtimeEntryAggregated) SetOtfp(v int64) {
	o.Otfp = &v
}

// GetOtfpShield returns the OtfpShield field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetOtfpShield() int64 {
	if o == nil || o.OtfpShield == nil {
		var ret int64
		return ret
	}
	return *o.OtfpShield
}

// GetOtfpShieldOk returns a tuple with the OtfpShield field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetOtfpShieldOk() (*int64, bool) {
	if o == nil || o.OtfpShield == nil {
		return nil, false
	}
	return o.OtfpShield, true
}

// HasOtfpShield returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasOtfpShield() bool {
	if o != nil && o.OtfpShield != nil {
		return true
	}

	return false
}

// SetOtfpShield gets a reference to the given int64 and assigns it to the OtfpShield field.
func (o *RealtimeEntryAggregated) SetOtfpShield(v int64) {
	o.OtfpShield = &v
}

// GetOtfpManifests returns the OtfpManifests field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetOtfpManifests() int64 {
	if o == nil || o.OtfpManifests == nil {
		var ret int64
		return ret
	}
	return *o.OtfpManifests
}

// GetOtfpManifestsOk returns a tuple with the OtfpManifests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetOtfpManifestsOk() (*int64, bool) {
	if o == nil || o.OtfpManifests == nil {
		return nil, false
	}
	return o.OtfpManifests, true
}

// HasOtfpManifests returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasOtfpManifests() bool {
	if o != nil && o.OtfpManifests != nil {
		return true
	}

	return false
}

// SetOtfpManifests gets a reference to the given int64 and assigns it to the OtfpManifests field.
func (o *RealtimeEntryAggregated) SetOtfpManifests(v int64) {
	o.OtfpManifests = &v
}

// GetVideo returns the Video field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetVideo() int64 {
	if o == nil || o.Video == nil {
		var ret int64
		return ret
	}
	return *o.Video
}

// GetVideoOk returns a tuple with the Video field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetVideoOk() (*int64, bool) {
	if o == nil || o.Video == nil {
		return nil, false
	}
	return o.Video, true
}

// HasVideo returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasVideo() bool {
	if o != nil && o.Video != nil {
		return true
	}

	return false
}

// SetVideo gets a reference to the given int64 and assigns it to the Video field.
func (o *RealtimeEntryAggregated) SetVideo(v int64) {
	o.Video = &v
}

// GetPci returns the Pci field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetPci() int64 {
	if o == nil || o.Pci == nil {
		var ret int64
		return ret
	}
	return *o.Pci
}

// GetPciOk returns a tuple with the Pci field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetPciOk() (*int64, bool) {
	if o == nil || o.Pci == nil {
		return nil, false
	}
	return o.Pci, true
}

// HasPci returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasPci() bool {
	if o != nil && o.Pci != nil {
		return true
	}

	return false
}

// SetPci gets a reference to the given int64 and assigns it to the Pci field.
func (o *RealtimeEntryAggregated) SetPci(v int64) {
	o.Pci = &v
}

// GetHTTP2 returns the HTTP2 field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetHTTP2() int64 {
	if o == nil || o.HTTP2 == nil {
		var ret int64
		return ret
	}
	return *o.HTTP2
}

// GetHTTP2Ok returns a tuple with the HTTP2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetHTTP2Ok() (*int64, bool) {
	if o == nil || o.HTTP2 == nil {
		return nil, false
	}
	return o.HTTP2, true
}

// HasHTTP2 returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasHTTP2() bool {
	if o != nil && o.HTTP2 != nil {
		return true
	}

	return false
}

// SetHTTP2 gets a reference to the given int64 and assigns it to the HTTP2 field.
func (o *RealtimeEntryAggregated) SetHTTP2(v int64) {
	o.HTTP2 = &v
}

// GetHTTP3 returns the HTTP3 field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetHTTP3() int64 {
	if o == nil || o.HTTP3 == nil {
		var ret int64
		return ret
	}
	return *o.HTTP3
}

// GetHTTP3Ok returns a tuple with the HTTP3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetHTTP3Ok() (*int64, bool) {
	if o == nil || o.HTTP3 == nil {
		return nil, false
	}
	return o.HTTP3, true
}

// HasHTTP3 returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasHTTP3() bool {
	if o != nil && o.HTTP3 != nil {
		return true
	}

	return false
}

// SetHTTP3 gets a reference to the given int64 and assigns it to the HTTP3 field.
func (o *RealtimeEntryAggregated) SetHTTP3(v int64) {
	o.HTTP3 = &v
}

// GetRestarts returns the Restarts field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetRestarts() int64 {
	if o == nil || o.Restarts == nil {
		var ret int64
		return ret
	}
	return *o.Restarts
}

// GetRestartsOk returns a tuple with the Restarts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetRestartsOk() (*int64, bool) {
	if o == nil || o.Restarts == nil {
		return nil, false
	}
	return o.Restarts, true
}

// HasRestarts returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasRestarts() bool {
	if o != nil && o.Restarts != nil {
		return true
	}

	return false
}

// SetRestarts gets a reference to the given int64 and assigns it to the Restarts field.
func (o *RealtimeEntryAggregated) SetRestarts(v int64) {
	o.Restarts = &v
}

// GetReqHeaderBytes returns the ReqHeaderBytes field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetReqHeaderBytes() int64 {
	if o == nil || o.ReqHeaderBytes == nil {
		var ret int64
		return ret
	}
	return *o.ReqHeaderBytes
}

// GetReqHeaderBytesOk returns a tuple with the ReqHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetReqHeaderBytesOk() (*int64, bool) {
	if o == nil || o.ReqHeaderBytes == nil {
		return nil, false
	}
	return o.ReqHeaderBytes, true
}

// HasReqHeaderBytes returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasReqHeaderBytes() bool {
	if o != nil && o.ReqHeaderBytes != nil {
		return true
	}

	return false
}

// SetReqHeaderBytes gets a reference to the given int64 and assigns it to the ReqHeaderBytes field.
func (o *RealtimeEntryAggregated) SetReqHeaderBytes(v int64) {
	o.ReqHeaderBytes = &v
}

// GetReqBodyBytes returns the ReqBodyBytes field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetReqBodyBytes() int64 {
	if o == nil || o.ReqBodyBytes == nil {
		var ret int64
		return ret
	}
	return *o.ReqBodyBytes
}

// GetReqBodyBytesOk returns a tuple with the ReqBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetReqBodyBytesOk() (*int64, bool) {
	if o == nil || o.ReqBodyBytes == nil {
		return nil, false
	}
	return o.ReqBodyBytes, true
}

// HasReqBodyBytes returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasReqBodyBytes() bool {
	if o != nil && o.ReqBodyBytes != nil {
		return true
	}

	return false
}

// SetReqBodyBytes gets a reference to the given int64 and assigns it to the ReqBodyBytes field.
func (o *RealtimeEntryAggregated) SetReqBodyBytes(v int64) {
	o.ReqBodyBytes = &v
}

// GetBereqHeaderBytes returns the BereqHeaderBytes field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetBereqHeaderBytes() int64 {
	if o == nil || o.BereqHeaderBytes == nil {
		var ret int64
		return ret
	}
	return *o.BereqHeaderBytes
}

// GetBereqHeaderBytesOk returns a tuple with the BereqHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetBereqHeaderBytesOk() (*int64, bool) {
	if o == nil || o.BereqHeaderBytes == nil {
		return nil, false
	}
	return o.BereqHeaderBytes, true
}

// HasBereqHeaderBytes returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasBereqHeaderBytes() bool {
	if o != nil && o.BereqHeaderBytes != nil {
		return true
	}

	return false
}

// SetBereqHeaderBytes gets a reference to the given int64 and assigns it to the BereqHeaderBytes field.
func (o *RealtimeEntryAggregated) SetBereqHeaderBytes(v int64) {
	o.BereqHeaderBytes = &v
}

// GetBereqBodyBytes returns the BereqBodyBytes field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetBereqBodyBytes() int64 {
	if o == nil || o.BereqBodyBytes == nil {
		var ret int64
		return ret
	}
	return *o.BereqBodyBytes
}

// GetBereqBodyBytesOk returns a tuple with the BereqBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetBereqBodyBytesOk() (*int64, bool) {
	if o == nil || o.BereqBodyBytes == nil {
		return nil, false
	}
	return o.BereqBodyBytes, true
}

// HasBereqBodyBytes returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasBereqBodyBytes() bool {
	if o != nil && o.BereqBodyBytes != nil {
		return true
	}

	return false
}

// SetBereqBodyBytes gets a reference to the given int64 and assigns it to the BereqBodyBytes field.
func (o *RealtimeEntryAggregated) SetBereqBodyBytes(v int64) {
	o.BereqBodyBytes = &v
}

// GetWafBlocked returns the WafBlocked field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetWafBlocked() int64 {
	if o == nil || o.WafBlocked == nil {
		var ret int64
		return ret
	}
	return *o.WafBlocked
}

// GetWafBlockedOk returns a tuple with the WafBlocked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetWafBlockedOk() (*int64, bool) {
	if o == nil || o.WafBlocked == nil {
		return nil, false
	}
	return o.WafBlocked, true
}

// HasWafBlocked returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasWafBlocked() bool {
	if o != nil && o.WafBlocked != nil {
		return true
	}

	return false
}

// SetWafBlocked gets a reference to the given int64 and assigns it to the WafBlocked field.
func (o *RealtimeEntryAggregated) SetWafBlocked(v int64) {
	o.WafBlocked = &v
}

// GetWafLogged returns the WafLogged field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetWafLogged() int64 {
	if o == nil || o.WafLogged == nil {
		var ret int64
		return ret
	}
	return *o.WafLogged
}

// GetWafLoggedOk returns a tuple with the WafLogged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetWafLoggedOk() (*int64, bool) {
	if o == nil || o.WafLogged == nil {
		return nil, false
	}
	return o.WafLogged, true
}

// HasWafLogged returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasWafLogged() bool {
	if o != nil && o.WafLogged != nil {
		return true
	}

	return false
}

// SetWafLogged gets a reference to the given int64 and assigns it to the WafLogged field.
func (o *RealtimeEntryAggregated) SetWafLogged(v int64) {
	o.WafLogged = &v
}

// GetWafPassed returns the WafPassed field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetWafPassed() int64 {
	if o == nil || o.WafPassed == nil {
		var ret int64
		return ret
	}
	return *o.WafPassed
}

// GetWafPassedOk returns a tuple with the WafPassed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetWafPassedOk() (*int64, bool) {
	if o == nil || o.WafPassed == nil {
		return nil, false
	}
	return o.WafPassed, true
}

// HasWafPassed returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasWafPassed() bool {
	if o != nil && o.WafPassed != nil {
		return true
	}

	return false
}

// SetWafPassed gets a reference to the given int64 and assigns it to the WafPassed field.
func (o *RealtimeEntryAggregated) SetWafPassed(v int64) {
	o.WafPassed = &v
}

// GetAttackReqHeaderBytes returns the AttackReqHeaderBytes field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetAttackReqHeaderBytes() int64 {
	if o == nil || o.AttackReqHeaderBytes == nil {
		var ret int64
		return ret
	}
	return *o.AttackReqHeaderBytes
}

// GetAttackReqHeaderBytesOk returns a tuple with the AttackReqHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetAttackReqHeaderBytesOk() (*int64, bool) {
	if o == nil || o.AttackReqHeaderBytes == nil {
		return nil, false
	}
	return o.AttackReqHeaderBytes, true
}

// HasAttackReqHeaderBytes returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasAttackReqHeaderBytes() bool {
	if o != nil && o.AttackReqHeaderBytes != nil {
		return true
	}

	return false
}

// SetAttackReqHeaderBytes gets a reference to the given int64 and assigns it to the AttackReqHeaderBytes field.
func (o *RealtimeEntryAggregated) SetAttackReqHeaderBytes(v int64) {
	o.AttackReqHeaderBytes = &v
}

// GetAttackReqBodyBytes returns the AttackReqBodyBytes field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetAttackReqBodyBytes() int64 {
	if o == nil || o.AttackReqBodyBytes == nil {
		var ret int64
		return ret
	}
	return *o.AttackReqBodyBytes
}

// GetAttackReqBodyBytesOk returns a tuple with the AttackReqBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetAttackReqBodyBytesOk() (*int64, bool) {
	if o == nil || o.AttackReqBodyBytes == nil {
		return nil, false
	}
	return o.AttackReqBodyBytes, true
}

// HasAttackReqBodyBytes returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasAttackReqBodyBytes() bool {
	if o != nil && o.AttackReqBodyBytes != nil {
		return true
	}

	return false
}

// SetAttackReqBodyBytes gets a reference to the given int64 and assigns it to the AttackReqBodyBytes field.
func (o *RealtimeEntryAggregated) SetAttackReqBodyBytes(v int64) {
	o.AttackReqBodyBytes = &v
}

// GetAttackRespSynthBytes returns the AttackRespSynthBytes field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetAttackRespSynthBytes() int64 {
	if o == nil || o.AttackRespSynthBytes == nil {
		var ret int64
		return ret
	}
	return *o.AttackRespSynthBytes
}

// GetAttackRespSynthBytesOk returns a tuple with the AttackRespSynthBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetAttackRespSynthBytesOk() (*int64, bool) {
	if o == nil || o.AttackRespSynthBytes == nil {
		return nil, false
	}
	return o.AttackRespSynthBytes, true
}

// HasAttackRespSynthBytes returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasAttackRespSynthBytes() bool {
	if o != nil && o.AttackRespSynthBytes != nil {
		return true
	}

	return false
}

// SetAttackRespSynthBytes gets a reference to the given int64 and assigns it to the AttackRespSynthBytes field.
func (o *RealtimeEntryAggregated) SetAttackRespSynthBytes(v int64) {
	o.AttackRespSynthBytes = &v
}

// GetAttackLoggedReqHeaderBytes returns the AttackLoggedReqHeaderBytes field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetAttackLoggedReqHeaderBytes() int64 {
	if o == nil || o.AttackLoggedReqHeaderBytes == nil {
		var ret int64
		return ret
	}
	return *o.AttackLoggedReqHeaderBytes
}

// GetAttackLoggedReqHeaderBytesOk returns a tuple with the AttackLoggedReqHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetAttackLoggedReqHeaderBytesOk() (*int64, bool) {
	if o == nil || o.AttackLoggedReqHeaderBytes == nil {
		return nil, false
	}
	return o.AttackLoggedReqHeaderBytes, true
}

// HasAttackLoggedReqHeaderBytes returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasAttackLoggedReqHeaderBytes() bool {
	if o != nil && o.AttackLoggedReqHeaderBytes != nil {
		return true
	}

	return false
}

// SetAttackLoggedReqHeaderBytes gets a reference to the given int64 and assigns it to the AttackLoggedReqHeaderBytes field.
func (o *RealtimeEntryAggregated) SetAttackLoggedReqHeaderBytes(v int64) {
	o.AttackLoggedReqHeaderBytes = &v
}

// GetAttackLoggedReqBodyBytes returns the AttackLoggedReqBodyBytes field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetAttackLoggedReqBodyBytes() int64 {
	if o == nil || o.AttackLoggedReqBodyBytes == nil {
		var ret int64
		return ret
	}
	return *o.AttackLoggedReqBodyBytes
}

// GetAttackLoggedReqBodyBytesOk returns a tuple with the AttackLoggedReqBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetAttackLoggedReqBodyBytesOk() (*int64, bool) {
	if o == nil || o.AttackLoggedReqBodyBytes == nil {
		return nil, false
	}
	return o.AttackLoggedReqBodyBytes, true
}

// HasAttackLoggedReqBodyBytes returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasAttackLoggedReqBodyBytes() bool {
	if o != nil && o.AttackLoggedReqBodyBytes != nil {
		return true
	}

	return false
}

// SetAttackLoggedReqBodyBytes gets a reference to the given int64 and assigns it to the AttackLoggedReqBodyBytes field.
func (o *RealtimeEntryAggregated) SetAttackLoggedReqBodyBytes(v int64) {
	o.AttackLoggedReqBodyBytes = &v
}

// GetAttackBlockedReqHeaderBytes returns the AttackBlockedReqHeaderBytes field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetAttackBlockedReqHeaderBytes() int64 {
	if o == nil || o.AttackBlockedReqHeaderBytes == nil {
		var ret int64
		return ret
	}
	return *o.AttackBlockedReqHeaderBytes
}

// GetAttackBlockedReqHeaderBytesOk returns a tuple with the AttackBlockedReqHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetAttackBlockedReqHeaderBytesOk() (*int64, bool) {
	if o == nil || o.AttackBlockedReqHeaderBytes == nil {
		return nil, false
	}
	return o.AttackBlockedReqHeaderBytes, true
}

// HasAttackBlockedReqHeaderBytes returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasAttackBlockedReqHeaderBytes() bool {
	if o != nil && o.AttackBlockedReqHeaderBytes != nil {
		return true
	}

	return false
}

// SetAttackBlockedReqHeaderBytes gets a reference to the given int64 and assigns it to the AttackBlockedReqHeaderBytes field.
func (o *RealtimeEntryAggregated) SetAttackBlockedReqHeaderBytes(v int64) {
	o.AttackBlockedReqHeaderBytes = &v
}

// GetAttackBlockedReqBodyBytes returns the AttackBlockedReqBodyBytes field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetAttackBlockedReqBodyBytes() int64 {
	if o == nil || o.AttackBlockedReqBodyBytes == nil {
		var ret int64
		return ret
	}
	return *o.AttackBlockedReqBodyBytes
}

// GetAttackBlockedReqBodyBytesOk returns a tuple with the AttackBlockedReqBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetAttackBlockedReqBodyBytesOk() (*int64, bool) {
	if o == nil || o.AttackBlockedReqBodyBytes == nil {
		return nil, false
	}
	return o.AttackBlockedReqBodyBytes, true
}

// HasAttackBlockedReqBodyBytes returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasAttackBlockedReqBodyBytes() bool {
	if o != nil && o.AttackBlockedReqBodyBytes != nil {
		return true
	}

	return false
}

// SetAttackBlockedReqBodyBytes gets a reference to the given int64 and assigns it to the AttackBlockedReqBodyBytes field.
func (o *RealtimeEntryAggregated) SetAttackBlockedReqBodyBytes(v int64) {
	o.AttackBlockedReqBodyBytes = &v
}

// GetAttackPassedReqHeaderBytes returns the AttackPassedReqHeaderBytes field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetAttackPassedReqHeaderBytes() int64 {
	if o == nil || o.AttackPassedReqHeaderBytes == nil {
		var ret int64
		return ret
	}
	return *o.AttackPassedReqHeaderBytes
}

// GetAttackPassedReqHeaderBytesOk returns a tuple with the AttackPassedReqHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetAttackPassedReqHeaderBytesOk() (*int64, bool) {
	if o == nil || o.AttackPassedReqHeaderBytes == nil {
		return nil, false
	}
	return o.AttackPassedReqHeaderBytes, true
}

// HasAttackPassedReqHeaderBytes returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasAttackPassedReqHeaderBytes() bool {
	if o != nil && o.AttackPassedReqHeaderBytes != nil {
		return true
	}

	return false
}

// SetAttackPassedReqHeaderBytes gets a reference to the given int64 and assigns it to the AttackPassedReqHeaderBytes field.
func (o *RealtimeEntryAggregated) SetAttackPassedReqHeaderBytes(v int64) {
	o.AttackPassedReqHeaderBytes = &v
}

// GetAttackPassedReqBodyBytes returns the AttackPassedReqBodyBytes field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetAttackPassedReqBodyBytes() int64 {
	if o == nil || o.AttackPassedReqBodyBytes == nil {
		var ret int64
		return ret
	}
	return *o.AttackPassedReqBodyBytes
}

// GetAttackPassedReqBodyBytesOk returns a tuple with the AttackPassedReqBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetAttackPassedReqBodyBytesOk() (*int64, bool) {
	if o == nil || o.AttackPassedReqBodyBytes == nil {
		return nil, false
	}
	return o.AttackPassedReqBodyBytes, true
}

// HasAttackPassedReqBodyBytes returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasAttackPassedReqBodyBytes() bool {
	if o != nil && o.AttackPassedReqBodyBytes != nil {
		return true
	}

	return false
}

// SetAttackPassedReqBodyBytes gets a reference to the given int64 and assigns it to the AttackPassedReqBodyBytes field.
func (o *RealtimeEntryAggregated) SetAttackPassedReqBodyBytes(v int64) {
	o.AttackPassedReqBodyBytes = &v
}

// GetShieldRespHeaderBytes returns the ShieldRespHeaderBytes field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetShieldRespHeaderBytes() int64 {
	if o == nil || o.ShieldRespHeaderBytes == nil {
		var ret int64
		return ret
	}
	return *o.ShieldRespHeaderBytes
}

// GetShieldRespHeaderBytesOk returns a tuple with the ShieldRespHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetShieldRespHeaderBytesOk() (*int64, bool) {
	if o == nil || o.ShieldRespHeaderBytes == nil {
		return nil, false
	}
	return o.ShieldRespHeaderBytes, true
}

// HasShieldRespHeaderBytes returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasShieldRespHeaderBytes() bool {
	if o != nil && o.ShieldRespHeaderBytes != nil {
		return true
	}

	return false
}

// SetShieldRespHeaderBytes gets a reference to the given int64 and assigns it to the ShieldRespHeaderBytes field.
func (o *RealtimeEntryAggregated) SetShieldRespHeaderBytes(v int64) {
	o.ShieldRespHeaderBytes = &v
}

// GetShieldRespBodyBytes returns the ShieldRespBodyBytes field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetShieldRespBodyBytes() int64 {
	if o == nil || o.ShieldRespBodyBytes == nil {
		var ret int64
		return ret
	}
	return *o.ShieldRespBodyBytes
}

// GetShieldRespBodyBytesOk returns a tuple with the ShieldRespBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetShieldRespBodyBytesOk() (*int64, bool) {
	if o == nil || o.ShieldRespBodyBytes == nil {
		return nil, false
	}
	return o.ShieldRespBodyBytes, true
}

// HasShieldRespBodyBytes returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasShieldRespBodyBytes() bool {
	if o != nil && o.ShieldRespBodyBytes != nil {
		return true
	}

	return false
}

// SetShieldRespBodyBytes gets a reference to the given int64 and assigns it to the ShieldRespBodyBytes field.
func (o *RealtimeEntryAggregated) SetShieldRespBodyBytes(v int64) {
	o.ShieldRespBodyBytes = &v
}

// GetOtfpRespHeaderBytes returns the OtfpRespHeaderBytes field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetOtfpRespHeaderBytes() int64 {
	if o == nil || o.OtfpRespHeaderBytes == nil {
		var ret int64
		return ret
	}
	return *o.OtfpRespHeaderBytes
}

// GetOtfpRespHeaderBytesOk returns a tuple with the OtfpRespHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetOtfpRespHeaderBytesOk() (*int64, bool) {
	if o == nil || o.OtfpRespHeaderBytes == nil {
		return nil, false
	}
	return o.OtfpRespHeaderBytes, true
}

// HasOtfpRespHeaderBytes returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasOtfpRespHeaderBytes() bool {
	if o != nil && o.OtfpRespHeaderBytes != nil {
		return true
	}

	return false
}

// SetOtfpRespHeaderBytes gets a reference to the given int64 and assigns it to the OtfpRespHeaderBytes field.
func (o *RealtimeEntryAggregated) SetOtfpRespHeaderBytes(v int64) {
	o.OtfpRespHeaderBytes = &v
}

// GetOtfpRespBodyBytes returns the OtfpRespBodyBytes field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetOtfpRespBodyBytes() int64 {
	if o == nil || o.OtfpRespBodyBytes == nil {
		var ret int64
		return ret
	}
	return *o.OtfpRespBodyBytes
}

// GetOtfpRespBodyBytesOk returns a tuple with the OtfpRespBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetOtfpRespBodyBytesOk() (*int64, bool) {
	if o == nil || o.OtfpRespBodyBytes == nil {
		return nil, false
	}
	return o.OtfpRespBodyBytes, true
}

// HasOtfpRespBodyBytes returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasOtfpRespBodyBytes() bool {
	if o != nil && o.OtfpRespBodyBytes != nil {
		return true
	}

	return false
}

// SetOtfpRespBodyBytes gets a reference to the given int64 and assigns it to the OtfpRespBodyBytes field.
func (o *RealtimeEntryAggregated) SetOtfpRespBodyBytes(v int64) {
	o.OtfpRespBodyBytes = &v
}

// GetOtfpShieldRespHeaderBytes returns the OtfpShieldRespHeaderBytes field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetOtfpShieldRespHeaderBytes() int64 {
	if o == nil || o.OtfpShieldRespHeaderBytes == nil {
		var ret int64
		return ret
	}
	return *o.OtfpShieldRespHeaderBytes
}

// GetOtfpShieldRespHeaderBytesOk returns a tuple with the OtfpShieldRespHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetOtfpShieldRespHeaderBytesOk() (*int64, bool) {
	if o == nil || o.OtfpShieldRespHeaderBytes == nil {
		return nil, false
	}
	return o.OtfpShieldRespHeaderBytes, true
}

// HasOtfpShieldRespHeaderBytes returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasOtfpShieldRespHeaderBytes() bool {
	if o != nil && o.OtfpShieldRespHeaderBytes != nil {
		return true
	}

	return false
}

// SetOtfpShieldRespHeaderBytes gets a reference to the given int64 and assigns it to the OtfpShieldRespHeaderBytes field.
func (o *RealtimeEntryAggregated) SetOtfpShieldRespHeaderBytes(v int64) {
	o.OtfpShieldRespHeaderBytes = &v
}

// GetOtfpShieldRespBodyBytes returns the OtfpShieldRespBodyBytes field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetOtfpShieldRespBodyBytes() int64 {
	if o == nil || o.OtfpShieldRespBodyBytes == nil {
		var ret int64
		return ret
	}
	return *o.OtfpShieldRespBodyBytes
}

// GetOtfpShieldRespBodyBytesOk returns a tuple with the OtfpShieldRespBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetOtfpShieldRespBodyBytesOk() (*int64, bool) {
	if o == nil || o.OtfpShieldRespBodyBytes == nil {
		return nil, false
	}
	return o.OtfpShieldRespBodyBytes, true
}

// HasOtfpShieldRespBodyBytes returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasOtfpShieldRespBodyBytes() bool {
	if o != nil && o.OtfpShieldRespBodyBytes != nil {
		return true
	}

	return false
}

// SetOtfpShieldRespBodyBytes gets a reference to the given int64 and assigns it to the OtfpShieldRespBodyBytes field.
func (o *RealtimeEntryAggregated) SetOtfpShieldRespBodyBytes(v int64) {
	o.OtfpShieldRespBodyBytes = &v
}

// GetOtfpShieldTime returns the OtfpShieldTime field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetOtfpShieldTime() float32 {
	if o == nil || o.OtfpShieldTime == nil {
		var ret float32
		return ret
	}
	return *o.OtfpShieldTime
}

// GetOtfpShieldTimeOk returns a tuple with the OtfpShieldTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetOtfpShieldTimeOk() (*float32, bool) {
	if o == nil || o.OtfpShieldTime == nil {
		return nil, false
	}
	return o.OtfpShieldTime, true
}

// HasOtfpShieldTime returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasOtfpShieldTime() bool {
	if o != nil && o.OtfpShieldTime != nil {
		return true
	}

	return false
}

// SetOtfpShieldTime gets a reference to the given float32 and assigns it to the OtfpShieldTime field.
func (o *RealtimeEntryAggregated) SetOtfpShieldTime(v float32) {
	o.OtfpShieldTime = &v
}

// GetOtfpDeliverTime returns the OtfpDeliverTime field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetOtfpDeliverTime() float32 {
	if o == nil || o.OtfpDeliverTime == nil {
		var ret float32
		return ret
	}
	return *o.OtfpDeliverTime
}

// GetOtfpDeliverTimeOk returns a tuple with the OtfpDeliverTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetOtfpDeliverTimeOk() (*float32, bool) {
	if o == nil || o.OtfpDeliverTime == nil {
		return nil, false
	}
	return o.OtfpDeliverTime, true
}

// HasOtfpDeliverTime returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasOtfpDeliverTime() bool {
	if o != nil && o.OtfpDeliverTime != nil {
		return true
	}

	return false
}

// SetOtfpDeliverTime gets a reference to the given float32 and assigns it to the OtfpDeliverTime field.
func (o *RealtimeEntryAggregated) SetOtfpDeliverTime(v float32) {
	o.OtfpDeliverTime = &v
}

// GetImgoptoRespHeaderBytes returns the ImgoptoRespHeaderBytes field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetImgoptoRespHeaderBytes() int64 {
	if o == nil || o.ImgoptoRespHeaderBytes == nil {
		var ret int64
		return ret
	}
	return *o.ImgoptoRespHeaderBytes
}

// GetImgoptoRespHeaderBytesOk returns a tuple with the ImgoptoRespHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetImgoptoRespHeaderBytesOk() (*int64, bool) {
	if o == nil || o.ImgoptoRespHeaderBytes == nil {
		return nil, false
	}
	return o.ImgoptoRespHeaderBytes, true
}

// HasImgoptoRespHeaderBytes returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasImgoptoRespHeaderBytes() bool {
	if o != nil && o.ImgoptoRespHeaderBytes != nil {
		return true
	}

	return false
}

// SetImgoptoRespHeaderBytes gets a reference to the given int64 and assigns it to the ImgoptoRespHeaderBytes field.
func (o *RealtimeEntryAggregated) SetImgoptoRespHeaderBytes(v int64) {
	o.ImgoptoRespHeaderBytes = &v
}

// GetImgoptoRespBodyBytes returns the ImgoptoRespBodyBytes field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetImgoptoRespBodyBytes() int64 {
	if o == nil || o.ImgoptoRespBodyBytes == nil {
		var ret int64
		return ret
	}
	return *o.ImgoptoRespBodyBytes
}

// GetImgoptoRespBodyBytesOk returns a tuple with the ImgoptoRespBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetImgoptoRespBodyBytesOk() (*int64, bool) {
	if o == nil || o.ImgoptoRespBodyBytes == nil {
		return nil, false
	}
	return o.ImgoptoRespBodyBytes, true
}

// HasImgoptoRespBodyBytes returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasImgoptoRespBodyBytes() bool {
	if o != nil && o.ImgoptoRespBodyBytes != nil {
		return true
	}

	return false
}

// SetImgoptoRespBodyBytes gets a reference to the given int64 and assigns it to the ImgoptoRespBodyBytes field.
func (o *RealtimeEntryAggregated) SetImgoptoRespBodyBytes(v int64) {
	o.ImgoptoRespBodyBytes = &v
}

// GetImgoptoShieldRespHeaderBytes returns the ImgoptoShieldRespHeaderBytes field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetImgoptoShieldRespHeaderBytes() int64 {
	if o == nil || o.ImgoptoShieldRespHeaderBytes == nil {
		var ret int64
		return ret
	}
	return *o.ImgoptoShieldRespHeaderBytes
}

// GetImgoptoShieldRespHeaderBytesOk returns a tuple with the ImgoptoShieldRespHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetImgoptoShieldRespHeaderBytesOk() (*int64, bool) {
	if o == nil || o.ImgoptoShieldRespHeaderBytes == nil {
		return nil, false
	}
	return o.ImgoptoShieldRespHeaderBytes, true
}

// HasImgoptoShieldRespHeaderBytes returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasImgoptoShieldRespHeaderBytes() bool {
	if o != nil && o.ImgoptoShieldRespHeaderBytes != nil {
		return true
	}

	return false
}

// SetImgoptoShieldRespHeaderBytes gets a reference to the given int64 and assigns it to the ImgoptoShieldRespHeaderBytes field.
func (o *RealtimeEntryAggregated) SetImgoptoShieldRespHeaderBytes(v int64) {
	o.ImgoptoShieldRespHeaderBytes = &v
}

// GetImgoptoShieldRespBodyBytes returns the ImgoptoShieldRespBodyBytes field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetImgoptoShieldRespBodyBytes() int64 {
	if o == nil || o.ImgoptoShieldRespBodyBytes == nil {
		var ret int64
		return ret
	}
	return *o.ImgoptoShieldRespBodyBytes
}

// GetImgoptoShieldRespBodyBytesOk returns a tuple with the ImgoptoShieldRespBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetImgoptoShieldRespBodyBytesOk() (*int64, bool) {
	if o == nil || o.ImgoptoShieldRespBodyBytes == nil {
		return nil, false
	}
	return o.ImgoptoShieldRespBodyBytes, true
}

// HasImgoptoShieldRespBodyBytes returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasImgoptoShieldRespBodyBytes() bool {
	if o != nil && o.ImgoptoShieldRespBodyBytes != nil {
		return true
	}

	return false
}

// SetImgoptoShieldRespBodyBytes gets a reference to the given int64 and assigns it to the ImgoptoShieldRespBodyBytes field.
func (o *RealtimeEntryAggregated) SetImgoptoShieldRespBodyBytes(v int64) {
	o.ImgoptoShieldRespBodyBytes = &v
}

// GetStatus1xx returns the Status1xx field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetStatus1xx() int64 {
	if o == nil || o.Status1xx == nil {
		var ret int64
		return ret
	}
	return *o.Status1xx
}

// GetStatus1xxOk returns a tuple with the Status1xx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetStatus1xxOk() (*int64, bool) {
	if o == nil || o.Status1xx == nil {
		return nil, false
	}
	return o.Status1xx, true
}

// HasStatus1xx returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasStatus1xx() bool {
	if o != nil && o.Status1xx != nil {
		return true
	}

	return false
}

// SetStatus1xx gets a reference to the given int64 and assigns it to the Status1xx field.
func (o *RealtimeEntryAggregated) SetStatus1xx(v int64) {
	o.Status1xx = &v
}

// GetStatus2xx returns the Status2xx field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetStatus2xx() int64 {
	if o == nil || o.Status2xx == nil {
		var ret int64
		return ret
	}
	return *o.Status2xx
}

// GetStatus2xxOk returns a tuple with the Status2xx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetStatus2xxOk() (*int64, bool) {
	if o == nil || o.Status2xx == nil {
		return nil, false
	}
	return o.Status2xx, true
}

// HasStatus2xx returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasStatus2xx() bool {
	if o != nil && o.Status2xx != nil {
		return true
	}

	return false
}

// SetStatus2xx gets a reference to the given int64 and assigns it to the Status2xx field.
func (o *RealtimeEntryAggregated) SetStatus2xx(v int64) {
	o.Status2xx = &v
}

// GetStatus3xx returns the Status3xx field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetStatus3xx() int64 {
	if o == nil || o.Status3xx == nil {
		var ret int64
		return ret
	}
	return *o.Status3xx
}

// GetStatus3xxOk returns a tuple with the Status3xx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetStatus3xxOk() (*int64, bool) {
	if o == nil || o.Status3xx == nil {
		return nil, false
	}
	return o.Status3xx, true
}

// HasStatus3xx returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasStatus3xx() bool {
	if o != nil && o.Status3xx != nil {
		return true
	}

	return false
}

// SetStatus3xx gets a reference to the given int64 and assigns it to the Status3xx field.
func (o *RealtimeEntryAggregated) SetStatus3xx(v int64) {
	o.Status3xx = &v
}

// GetStatus4xx returns the Status4xx field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetStatus4xx() int64 {
	if o == nil || o.Status4xx == nil {
		var ret int64
		return ret
	}
	return *o.Status4xx
}

// GetStatus4xxOk returns a tuple with the Status4xx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetStatus4xxOk() (*int64, bool) {
	if o == nil || o.Status4xx == nil {
		return nil, false
	}
	return o.Status4xx, true
}

// HasStatus4xx returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasStatus4xx() bool {
	if o != nil && o.Status4xx != nil {
		return true
	}

	return false
}

// SetStatus4xx gets a reference to the given int64 and assigns it to the Status4xx field.
func (o *RealtimeEntryAggregated) SetStatus4xx(v int64) {
	o.Status4xx = &v
}

// GetStatus5xx returns the Status5xx field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetStatus5xx() int64 {
	if o == nil || o.Status5xx == nil {
		var ret int64
		return ret
	}
	return *o.Status5xx
}

// GetStatus5xxOk returns a tuple with the Status5xx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetStatus5xxOk() (*int64, bool) {
	if o == nil || o.Status5xx == nil {
		return nil, false
	}
	return o.Status5xx, true
}

// HasStatus5xx returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasStatus5xx() bool {
	if o != nil && o.Status5xx != nil {
		return true
	}

	return false
}

// SetStatus5xx gets a reference to the given int64 and assigns it to the Status5xx field.
func (o *RealtimeEntryAggregated) SetStatus5xx(v int64) {
	o.Status5xx = &v
}

// GetStatus200 returns the Status200 field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetStatus200() int64 {
	if o == nil || o.Status200 == nil {
		var ret int64
		return ret
	}
	return *o.Status200
}

// GetStatus200Ok returns a tuple with the Status200 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetStatus200Ok() (*int64, bool) {
	if o == nil || o.Status200 == nil {
		return nil, false
	}
	return o.Status200, true
}

// HasStatus200 returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasStatus200() bool {
	if o != nil && o.Status200 != nil {
		return true
	}

	return false
}

// SetStatus200 gets a reference to the given int64 and assigns it to the Status200 field.
func (o *RealtimeEntryAggregated) SetStatus200(v int64) {
	o.Status200 = &v
}

// GetStatus204 returns the Status204 field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetStatus204() int64 {
	if o == nil || o.Status204 == nil {
		var ret int64
		return ret
	}
	return *o.Status204
}

// GetStatus204Ok returns a tuple with the Status204 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetStatus204Ok() (*int64, bool) {
	if o == nil || o.Status204 == nil {
		return nil, false
	}
	return o.Status204, true
}

// HasStatus204 returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasStatus204() bool {
	if o != nil && o.Status204 != nil {
		return true
	}

	return false
}

// SetStatus204 gets a reference to the given int64 and assigns it to the Status204 field.
func (o *RealtimeEntryAggregated) SetStatus204(v int64) {
	o.Status204 = &v
}

// GetStatus206 returns the Status206 field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetStatus206() int64 {
	if o == nil || o.Status206 == nil {
		var ret int64
		return ret
	}
	return *o.Status206
}

// GetStatus206Ok returns a tuple with the Status206 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetStatus206Ok() (*int64, bool) {
	if o == nil || o.Status206 == nil {
		return nil, false
	}
	return o.Status206, true
}

// HasStatus206 returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasStatus206() bool {
	if o != nil && o.Status206 != nil {
		return true
	}

	return false
}

// SetStatus206 gets a reference to the given int64 and assigns it to the Status206 field.
func (o *RealtimeEntryAggregated) SetStatus206(v int64) {
	o.Status206 = &v
}

// GetStatus301 returns the Status301 field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetStatus301() int64 {
	if o == nil || o.Status301 == nil {
		var ret int64
		return ret
	}
	return *o.Status301
}

// GetStatus301Ok returns a tuple with the Status301 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetStatus301Ok() (*int64, bool) {
	if o == nil || o.Status301 == nil {
		return nil, false
	}
	return o.Status301, true
}

// HasStatus301 returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasStatus301() bool {
	if o != nil && o.Status301 != nil {
		return true
	}

	return false
}

// SetStatus301 gets a reference to the given int64 and assigns it to the Status301 field.
func (o *RealtimeEntryAggregated) SetStatus301(v int64) {
	o.Status301 = &v
}

// GetStatus302 returns the Status302 field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetStatus302() int64 {
	if o == nil || o.Status302 == nil {
		var ret int64
		return ret
	}
	return *o.Status302
}

// GetStatus302Ok returns a tuple with the Status302 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetStatus302Ok() (*int64, bool) {
	if o == nil || o.Status302 == nil {
		return nil, false
	}
	return o.Status302, true
}

// HasStatus302 returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasStatus302() bool {
	if o != nil && o.Status302 != nil {
		return true
	}

	return false
}

// SetStatus302 gets a reference to the given int64 and assigns it to the Status302 field.
func (o *RealtimeEntryAggregated) SetStatus302(v int64) {
	o.Status302 = &v
}

// GetStatus304 returns the Status304 field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetStatus304() int64 {
	if o == nil || o.Status304 == nil {
		var ret int64
		return ret
	}
	return *o.Status304
}

// GetStatus304Ok returns a tuple with the Status304 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetStatus304Ok() (*int64, bool) {
	if o == nil || o.Status304 == nil {
		return nil, false
	}
	return o.Status304, true
}

// HasStatus304 returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasStatus304() bool {
	if o != nil && o.Status304 != nil {
		return true
	}

	return false
}

// SetStatus304 gets a reference to the given int64 and assigns it to the Status304 field.
func (o *RealtimeEntryAggregated) SetStatus304(v int64) {
	o.Status304 = &v
}

// GetStatus400 returns the Status400 field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetStatus400() int64 {
	if o == nil || o.Status400 == nil {
		var ret int64
		return ret
	}
	return *o.Status400
}

// GetStatus400Ok returns a tuple with the Status400 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetStatus400Ok() (*int64, bool) {
	if o == nil || o.Status400 == nil {
		return nil, false
	}
	return o.Status400, true
}

// HasStatus400 returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasStatus400() bool {
	if o != nil && o.Status400 != nil {
		return true
	}

	return false
}

// SetStatus400 gets a reference to the given int64 and assigns it to the Status400 field.
func (o *RealtimeEntryAggregated) SetStatus400(v int64) {
	o.Status400 = &v
}

// GetStatus401 returns the Status401 field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetStatus401() int64 {
	if o == nil || o.Status401 == nil {
		var ret int64
		return ret
	}
	return *o.Status401
}

// GetStatus401Ok returns a tuple with the Status401 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetStatus401Ok() (*int64, bool) {
	if o == nil || o.Status401 == nil {
		return nil, false
	}
	return o.Status401, true
}

// HasStatus401 returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasStatus401() bool {
	if o != nil && o.Status401 != nil {
		return true
	}

	return false
}

// SetStatus401 gets a reference to the given int64 and assigns it to the Status401 field.
func (o *RealtimeEntryAggregated) SetStatus401(v int64) {
	o.Status401 = &v
}

// GetStatus403 returns the Status403 field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetStatus403() int64 {
	if o == nil || o.Status403 == nil {
		var ret int64
		return ret
	}
	return *o.Status403
}

// GetStatus403Ok returns a tuple with the Status403 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetStatus403Ok() (*int64, bool) {
	if o == nil || o.Status403 == nil {
		return nil, false
	}
	return o.Status403, true
}

// HasStatus403 returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasStatus403() bool {
	if o != nil && o.Status403 != nil {
		return true
	}

	return false
}

// SetStatus403 gets a reference to the given int64 and assigns it to the Status403 field.
func (o *RealtimeEntryAggregated) SetStatus403(v int64) {
	o.Status403 = &v
}

// GetStatus404 returns the Status404 field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetStatus404() int64 {
	if o == nil || o.Status404 == nil {
		var ret int64
		return ret
	}
	return *o.Status404
}

// GetStatus404Ok returns a tuple with the Status404 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetStatus404Ok() (*int64, bool) {
	if o == nil || o.Status404 == nil {
		return nil, false
	}
	return o.Status404, true
}

// HasStatus404 returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasStatus404() bool {
	if o != nil && o.Status404 != nil {
		return true
	}

	return false
}

// SetStatus404 gets a reference to the given int64 and assigns it to the Status404 field.
func (o *RealtimeEntryAggregated) SetStatus404(v int64) {
	o.Status404 = &v
}

// GetStatus406 returns the Status406 field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetStatus406() int64 {
	if o == nil || o.Status406 == nil {
		var ret int64
		return ret
	}
	return *o.Status406
}

// GetStatus406Ok returns a tuple with the Status406 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetStatus406Ok() (*int64, bool) {
	if o == nil || o.Status406 == nil {
		return nil, false
	}
	return o.Status406, true
}

// HasStatus406 returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasStatus406() bool {
	if o != nil && o.Status406 != nil {
		return true
	}

	return false
}

// SetStatus406 gets a reference to the given int64 and assigns it to the Status406 field.
func (o *RealtimeEntryAggregated) SetStatus406(v int64) {
	o.Status406 = &v
}

// GetStatus416 returns the Status416 field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetStatus416() int64 {
	if o == nil || o.Status416 == nil {
		var ret int64
		return ret
	}
	return *o.Status416
}

// GetStatus416Ok returns a tuple with the Status416 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetStatus416Ok() (*int64, bool) {
	if o == nil || o.Status416 == nil {
		return nil, false
	}
	return o.Status416, true
}

// HasStatus416 returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasStatus416() bool {
	if o != nil && o.Status416 != nil {
		return true
	}

	return false
}

// SetStatus416 gets a reference to the given int64 and assigns it to the Status416 field.
func (o *RealtimeEntryAggregated) SetStatus416(v int64) {
	o.Status416 = &v
}

// GetStatus429 returns the Status429 field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetStatus429() int64 {
	if o == nil || o.Status429 == nil {
		var ret int64
		return ret
	}
	return *o.Status429
}

// GetStatus429Ok returns a tuple with the Status429 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetStatus429Ok() (*int64, bool) {
	if o == nil || o.Status429 == nil {
		return nil, false
	}
	return o.Status429, true
}

// HasStatus429 returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasStatus429() bool {
	if o != nil && o.Status429 != nil {
		return true
	}

	return false
}

// SetStatus429 gets a reference to the given int64 and assigns it to the Status429 field.
func (o *RealtimeEntryAggregated) SetStatus429(v int64) {
	o.Status429 = &v
}

// GetStatus500 returns the Status500 field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetStatus500() int64 {
	if o == nil || o.Status500 == nil {
		var ret int64
		return ret
	}
	return *o.Status500
}

// GetStatus500Ok returns a tuple with the Status500 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetStatus500Ok() (*int64, bool) {
	if o == nil || o.Status500 == nil {
		return nil, false
	}
	return o.Status500, true
}

// HasStatus500 returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasStatus500() bool {
	if o != nil && o.Status500 != nil {
		return true
	}

	return false
}

// SetStatus500 gets a reference to the given int64 and assigns it to the Status500 field.
func (o *RealtimeEntryAggregated) SetStatus500(v int64) {
	o.Status500 = &v
}

// GetStatus501 returns the Status501 field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetStatus501() int64 {
	if o == nil || o.Status501 == nil {
		var ret int64
		return ret
	}
	return *o.Status501
}

// GetStatus501Ok returns a tuple with the Status501 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetStatus501Ok() (*int64, bool) {
	if o == nil || o.Status501 == nil {
		return nil, false
	}
	return o.Status501, true
}

// HasStatus501 returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasStatus501() bool {
	if o != nil && o.Status501 != nil {
		return true
	}

	return false
}

// SetStatus501 gets a reference to the given int64 and assigns it to the Status501 field.
func (o *RealtimeEntryAggregated) SetStatus501(v int64) {
	o.Status501 = &v
}

// GetStatus502 returns the Status502 field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetStatus502() int64 {
	if o == nil || o.Status502 == nil {
		var ret int64
		return ret
	}
	return *o.Status502
}

// GetStatus502Ok returns a tuple with the Status502 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetStatus502Ok() (*int64, bool) {
	if o == nil || o.Status502 == nil {
		return nil, false
	}
	return o.Status502, true
}

// HasStatus502 returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasStatus502() bool {
	if o != nil && o.Status502 != nil {
		return true
	}

	return false
}

// SetStatus502 gets a reference to the given int64 and assigns it to the Status502 field.
func (o *RealtimeEntryAggregated) SetStatus502(v int64) {
	o.Status502 = &v
}

// GetStatus503 returns the Status503 field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetStatus503() int64 {
	if o == nil || o.Status503 == nil {
		var ret int64
		return ret
	}
	return *o.Status503
}

// GetStatus503Ok returns a tuple with the Status503 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetStatus503Ok() (*int64, bool) {
	if o == nil || o.Status503 == nil {
		return nil, false
	}
	return o.Status503, true
}

// HasStatus503 returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasStatus503() bool {
	if o != nil && o.Status503 != nil {
		return true
	}

	return false
}

// SetStatus503 gets a reference to the given int64 and assigns it to the Status503 field.
func (o *RealtimeEntryAggregated) SetStatus503(v int64) {
	o.Status503 = &v
}

// GetStatus504 returns the Status504 field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetStatus504() int64 {
	if o == nil || o.Status504 == nil {
		var ret int64
		return ret
	}
	return *o.Status504
}

// GetStatus504Ok returns a tuple with the Status504 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetStatus504Ok() (*int64, bool) {
	if o == nil || o.Status504 == nil {
		return nil, false
	}
	return o.Status504, true
}

// HasStatus504 returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasStatus504() bool {
	if o != nil && o.Status504 != nil {
		return true
	}

	return false
}

// SetStatus504 gets a reference to the given int64 and assigns it to the Status504 field.
func (o *RealtimeEntryAggregated) SetStatus504(v int64) {
	o.Status504 = &v
}

// GetStatus505 returns the Status505 field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetStatus505() int64 {
	if o == nil || o.Status505 == nil {
		var ret int64
		return ret
	}
	return *o.Status505
}

// GetStatus505Ok returns a tuple with the Status505 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetStatus505Ok() (*int64, bool) {
	if o == nil || o.Status505 == nil {
		return nil, false
	}
	return o.Status505, true
}

// HasStatus505 returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasStatus505() bool {
	if o != nil && o.Status505 != nil {
		return true
	}

	return false
}

// SetStatus505 gets a reference to the given int64 and assigns it to the Status505 field.
func (o *RealtimeEntryAggregated) SetStatus505(v int64) {
	o.Status505 = &v
}

// GetUncacheable returns the Uncacheable field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetUncacheable() int64 {
	if o == nil || o.Uncacheable == nil {
		var ret int64
		return ret
	}
	return *o.Uncacheable
}

// GetUncacheableOk returns a tuple with the Uncacheable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetUncacheableOk() (*int64, bool) {
	if o == nil || o.Uncacheable == nil {
		return nil, false
	}
	return o.Uncacheable, true
}

// HasUncacheable returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasUncacheable() bool {
	if o != nil && o.Uncacheable != nil {
		return true
	}

	return false
}

// SetUncacheable gets a reference to the given int64 and assigns it to the Uncacheable field.
func (o *RealtimeEntryAggregated) SetUncacheable(v int64) {
	o.Uncacheable = &v
}

// GetPassTime returns the PassTime field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetPassTime() float32 {
	if o == nil || o.PassTime == nil {
		var ret float32
		return ret
	}
	return *o.PassTime
}

// GetPassTimeOk returns a tuple with the PassTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetPassTimeOk() (*float32, bool) {
	if o == nil || o.PassTime == nil {
		return nil, false
	}
	return o.PassTime, true
}

// HasPassTime returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasPassTime() bool {
	if o != nil && o.PassTime != nil {
		return true
	}

	return false
}

// SetPassTime gets a reference to the given float32 and assigns it to the PassTime field.
func (o *RealtimeEntryAggregated) SetPassTime(v float32) {
	o.PassTime = &v
}

// GetTLS returns the TLS field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetTLS() int64 {
	if o == nil || o.TLS == nil {
		var ret int64
		return ret
	}
	return *o.TLS
}

// GetTLSOk returns a tuple with the TLS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetTLSOk() (*int64, bool) {
	if o == nil || o.TLS == nil {
		return nil, false
	}
	return o.TLS, true
}

// HasTLS returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasTLS() bool {
	if o != nil && o.TLS != nil {
		return true
	}

	return false
}

// SetTLS gets a reference to the given int64 and assigns it to the TLS field.
func (o *RealtimeEntryAggregated) SetTLS(v int64) {
	o.TLS = &v
}

// GetTLSV10 returns the TLSV10 field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetTLSV10() int64 {
	if o == nil || o.TLSV10 == nil {
		var ret int64
		return ret
	}
	return *o.TLSV10
}

// GetTLSV10Ok returns a tuple with the TLSV10 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetTLSV10Ok() (*int64, bool) {
	if o == nil || o.TLSV10 == nil {
		return nil, false
	}
	return o.TLSV10, true
}

// HasTLSV10 returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasTLSV10() bool {
	if o != nil && o.TLSV10 != nil {
		return true
	}

	return false
}

// SetTLSV10 gets a reference to the given int64 and assigns it to the TLSV10 field.
func (o *RealtimeEntryAggregated) SetTLSV10(v int64) {
	o.TLSV10 = &v
}

// GetTLSV11 returns the TLSV11 field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetTLSV11() int64 {
	if o == nil || o.TLSV11 == nil {
		var ret int64
		return ret
	}
	return *o.TLSV11
}

// GetTLSV11Ok returns a tuple with the TLSV11 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetTLSV11Ok() (*int64, bool) {
	if o == nil || o.TLSV11 == nil {
		return nil, false
	}
	return o.TLSV11, true
}

// HasTLSV11 returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasTLSV11() bool {
	if o != nil && o.TLSV11 != nil {
		return true
	}

	return false
}

// SetTLSV11 gets a reference to the given int64 and assigns it to the TLSV11 field.
func (o *RealtimeEntryAggregated) SetTLSV11(v int64) {
	o.TLSV11 = &v
}

// GetTLSV12 returns the TLSV12 field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetTLSV12() int64 {
	if o == nil || o.TLSV12 == nil {
		var ret int64
		return ret
	}
	return *o.TLSV12
}

// GetTLSV12Ok returns a tuple with the TLSV12 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetTLSV12Ok() (*int64, bool) {
	if o == nil || o.TLSV12 == nil {
		return nil, false
	}
	return o.TLSV12, true
}

// HasTLSV12 returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasTLSV12() bool {
	if o != nil && o.TLSV12 != nil {
		return true
	}

	return false
}

// SetTLSV12 gets a reference to the given int64 and assigns it to the TLSV12 field.
func (o *RealtimeEntryAggregated) SetTLSV12(v int64) {
	o.TLSV12 = &v
}

// GetTLSV13 returns the TLSV13 field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetTLSV13() int64 {
	if o == nil || o.TLSV13 == nil {
		var ret int64
		return ret
	}
	return *o.TLSV13
}

// GetTLSV13Ok returns a tuple with the TLSV13 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetTLSV13Ok() (*int64, bool) {
	if o == nil || o.TLSV13 == nil {
		return nil, false
	}
	return o.TLSV13, true
}

// HasTLSV13 returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasTLSV13() bool {
	if o != nil && o.TLSV13 != nil {
		return true
	}

	return false
}

// SetTLSV13 gets a reference to the given int64 and assigns it to the TLSV13 field.
func (o *RealtimeEntryAggregated) SetTLSV13(v int64) {
	o.TLSV13 = &v
}

// GetObjectSize1k returns the ObjectSize1k field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetObjectSize1k() int64 {
	if o == nil || o.ObjectSize1k == nil {
		var ret int64
		return ret
	}
	return *o.ObjectSize1k
}

// GetObjectSize1kOk returns a tuple with the ObjectSize1k field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetObjectSize1kOk() (*int64, bool) {
	if o == nil || o.ObjectSize1k == nil {
		return nil, false
	}
	return o.ObjectSize1k, true
}

// HasObjectSize1k returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasObjectSize1k() bool {
	if o != nil && o.ObjectSize1k != nil {
		return true
	}

	return false
}

// SetObjectSize1k gets a reference to the given int64 and assigns it to the ObjectSize1k field.
func (o *RealtimeEntryAggregated) SetObjectSize1k(v int64) {
	o.ObjectSize1k = &v
}

// GetObjectSize10k returns the ObjectSize10k field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetObjectSize10k() int64 {
	if o == nil || o.ObjectSize10k == nil {
		var ret int64
		return ret
	}
	return *o.ObjectSize10k
}

// GetObjectSize10kOk returns a tuple with the ObjectSize10k field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetObjectSize10kOk() (*int64, bool) {
	if o == nil || o.ObjectSize10k == nil {
		return nil, false
	}
	return o.ObjectSize10k, true
}

// HasObjectSize10k returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasObjectSize10k() bool {
	if o != nil && o.ObjectSize10k != nil {
		return true
	}

	return false
}

// SetObjectSize10k gets a reference to the given int64 and assigns it to the ObjectSize10k field.
func (o *RealtimeEntryAggregated) SetObjectSize10k(v int64) {
	o.ObjectSize10k = &v
}

// GetObjectSize100k returns the ObjectSize100k field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetObjectSize100k() int64 {
	if o == nil || o.ObjectSize100k == nil {
		var ret int64
		return ret
	}
	return *o.ObjectSize100k
}

// GetObjectSize100kOk returns a tuple with the ObjectSize100k field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetObjectSize100kOk() (*int64, bool) {
	if o == nil || o.ObjectSize100k == nil {
		return nil, false
	}
	return o.ObjectSize100k, true
}

// HasObjectSize100k returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasObjectSize100k() bool {
	if o != nil && o.ObjectSize100k != nil {
		return true
	}

	return false
}

// SetObjectSize100k gets a reference to the given int64 and assigns it to the ObjectSize100k field.
func (o *RealtimeEntryAggregated) SetObjectSize100k(v int64) {
	o.ObjectSize100k = &v
}

// GetObjectSize1m returns the ObjectSize1m field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetObjectSize1m() int64 {
	if o == nil || o.ObjectSize1m == nil {
		var ret int64
		return ret
	}
	return *o.ObjectSize1m
}

// GetObjectSize1mOk returns a tuple with the ObjectSize1m field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetObjectSize1mOk() (*int64, bool) {
	if o == nil || o.ObjectSize1m == nil {
		return nil, false
	}
	return o.ObjectSize1m, true
}

// HasObjectSize1m returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasObjectSize1m() bool {
	if o != nil && o.ObjectSize1m != nil {
		return true
	}

	return false
}

// SetObjectSize1m gets a reference to the given int64 and assigns it to the ObjectSize1m field.
func (o *RealtimeEntryAggregated) SetObjectSize1m(v int64) {
	o.ObjectSize1m = &v
}

// GetObjectSize10m returns the ObjectSize10m field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetObjectSize10m() int64 {
	if o == nil || o.ObjectSize10m == nil {
		var ret int64
		return ret
	}
	return *o.ObjectSize10m
}

// GetObjectSize10mOk returns a tuple with the ObjectSize10m field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetObjectSize10mOk() (*int64, bool) {
	if o == nil || o.ObjectSize10m == nil {
		return nil, false
	}
	return o.ObjectSize10m, true
}

// HasObjectSize10m returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasObjectSize10m() bool {
	if o != nil && o.ObjectSize10m != nil {
		return true
	}

	return false
}

// SetObjectSize10m gets a reference to the given int64 and assigns it to the ObjectSize10m field.
func (o *RealtimeEntryAggregated) SetObjectSize10m(v int64) {
	o.ObjectSize10m = &v
}

// GetObjectSize100m returns the ObjectSize100m field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetObjectSize100m() int64 {
	if o == nil || o.ObjectSize100m == nil {
		var ret int64
		return ret
	}
	return *o.ObjectSize100m
}

// GetObjectSize100mOk returns a tuple with the ObjectSize100m field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetObjectSize100mOk() (*int64, bool) {
	if o == nil || o.ObjectSize100m == nil {
		return nil, false
	}
	return o.ObjectSize100m, true
}

// HasObjectSize100m returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasObjectSize100m() bool {
	if o != nil && o.ObjectSize100m != nil {
		return true
	}

	return false
}

// SetObjectSize100m gets a reference to the given int64 and assigns it to the ObjectSize100m field.
func (o *RealtimeEntryAggregated) SetObjectSize100m(v int64) {
	o.ObjectSize100m = &v
}

// GetObjectSize1g returns the ObjectSize1g field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetObjectSize1g() int64 {
	if o == nil || o.ObjectSize1g == nil {
		var ret int64
		return ret
	}
	return *o.ObjectSize1g
}

// GetObjectSize1gOk returns a tuple with the ObjectSize1g field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetObjectSize1gOk() (*int64, bool) {
	if o == nil || o.ObjectSize1g == nil {
		return nil, false
	}
	return o.ObjectSize1g, true
}

// HasObjectSize1g returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasObjectSize1g() bool {
	if o != nil && o.ObjectSize1g != nil {
		return true
	}

	return false
}

// SetObjectSize1g gets a reference to the given int64 and assigns it to the ObjectSize1g field.
func (o *RealtimeEntryAggregated) SetObjectSize1g(v int64) {
	o.ObjectSize1g = &v
}

// GetObjectSizeOther returns the ObjectSizeOther field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetObjectSizeOther() int64 {
	if o == nil || o.ObjectSizeOther == nil {
		var ret int64
		return ret
	}
	return *o.ObjectSizeOther
}

// GetObjectSizeOtherOk returns a tuple with the ObjectSizeOther field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetObjectSizeOtherOk() (*int64, bool) {
	if o == nil || o.ObjectSizeOther == nil {
		return nil, false
	}
	return o.ObjectSizeOther, true
}

// HasObjectSizeOther returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasObjectSizeOther() bool {
	if o != nil && o.ObjectSizeOther != nil {
		return true
	}

	return false
}

// SetObjectSizeOther gets a reference to the given int64 and assigns it to the ObjectSizeOther field.
func (o *RealtimeEntryAggregated) SetObjectSizeOther(v int64) {
	o.ObjectSizeOther = &v
}

// GetRecvSubTime returns the RecvSubTime field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetRecvSubTime() float32 {
	if o == nil || o.RecvSubTime == nil {
		var ret float32
		return ret
	}
	return *o.RecvSubTime
}

// GetRecvSubTimeOk returns a tuple with the RecvSubTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetRecvSubTimeOk() (*float32, bool) {
	if o == nil || o.RecvSubTime == nil {
		return nil, false
	}
	return o.RecvSubTime, true
}

// HasRecvSubTime returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasRecvSubTime() bool {
	if o != nil && o.RecvSubTime != nil {
		return true
	}

	return false
}

// SetRecvSubTime gets a reference to the given float32 and assigns it to the RecvSubTime field.
func (o *RealtimeEntryAggregated) SetRecvSubTime(v float32) {
	o.RecvSubTime = &v
}

// GetRecvSubCount returns the RecvSubCount field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetRecvSubCount() int64 {
	if o == nil || o.RecvSubCount == nil {
		var ret int64
		return ret
	}
	return *o.RecvSubCount
}

// GetRecvSubCountOk returns a tuple with the RecvSubCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetRecvSubCountOk() (*int64, bool) {
	if o == nil || o.RecvSubCount == nil {
		return nil, false
	}
	return o.RecvSubCount, true
}

// HasRecvSubCount returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasRecvSubCount() bool {
	if o != nil && o.RecvSubCount != nil {
		return true
	}

	return false
}

// SetRecvSubCount gets a reference to the given int64 and assigns it to the RecvSubCount field.
func (o *RealtimeEntryAggregated) SetRecvSubCount(v int64) {
	o.RecvSubCount = &v
}

// GetHashSubTime returns the HashSubTime field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetHashSubTime() float32 {
	if o == nil || o.HashSubTime == nil {
		var ret float32
		return ret
	}
	return *o.HashSubTime
}

// GetHashSubTimeOk returns a tuple with the HashSubTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetHashSubTimeOk() (*float32, bool) {
	if o == nil || o.HashSubTime == nil {
		return nil, false
	}
	return o.HashSubTime, true
}

// HasHashSubTime returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasHashSubTime() bool {
	if o != nil && o.HashSubTime != nil {
		return true
	}

	return false
}

// SetHashSubTime gets a reference to the given float32 and assigns it to the HashSubTime field.
func (o *RealtimeEntryAggregated) SetHashSubTime(v float32) {
	o.HashSubTime = &v
}

// GetHashSubCount returns the HashSubCount field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetHashSubCount() int64 {
	if o == nil || o.HashSubCount == nil {
		var ret int64
		return ret
	}
	return *o.HashSubCount
}

// GetHashSubCountOk returns a tuple with the HashSubCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetHashSubCountOk() (*int64, bool) {
	if o == nil || o.HashSubCount == nil {
		return nil, false
	}
	return o.HashSubCount, true
}

// HasHashSubCount returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasHashSubCount() bool {
	if o != nil && o.HashSubCount != nil {
		return true
	}

	return false
}

// SetHashSubCount gets a reference to the given int64 and assigns it to the HashSubCount field.
func (o *RealtimeEntryAggregated) SetHashSubCount(v int64) {
	o.HashSubCount = &v
}

// GetMissSubTime returns the MissSubTime field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetMissSubTime() float32 {
	if o == nil || o.MissSubTime == nil {
		var ret float32
		return ret
	}
	return *o.MissSubTime
}

// GetMissSubTimeOk returns a tuple with the MissSubTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetMissSubTimeOk() (*float32, bool) {
	if o == nil || o.MissSubTime == nil {
		return nil, false
	}
	return o.MissSubTime, true
}

// HasMissSubTime returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasMissSubTime() bool {
	if o != nil && o.MissSubTime != nil {
		return true
	}

	return false
}

// SetMissSubTime gets a reference to the given float32 and assigns it to the MissSubTime field.
func (o *RealtimeEntryAggregated) SetMissSubTime(v float32) {
	o.MissSubTime = &v
}

// GetMissSubCount returns the MissSubCount field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetMissSubCount() int64 {
	if o == nil || o.MissSubCount == nil {
		var ret int64
		return ret
	}
	return *o.MissSubCount
}

// GetMissSubCountOk returns a tuple with the MissSubCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetMissSubCountOk() (*int64, bool) {
	if o == nil || o.MissSubCount == nil {
		return nil, false
	}
	return o.MissSubCount, true
}

// HasMissSubCount returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasMissSubCount() bool {
	if o != nil && o.MissSubCount != nil {
		return true
	}

	return false
}

// SetMissSubCount gets a reference to the given int64 and assigns it to the MissSubCount field.
func (o *RealtimeEntryAggregated) SetMissSubCount(v int64) {
	o.MissSubCount = &v
}

// GetFetchSubTime returns the FetchSubTime field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetFetchSubTime() float32 {
	if o == nil || o.FetchSubTime == nil {
		var ret float32
		return ret
	}
	return *o.FetchSubTime
}

// GetFetchSubTimeOk returns a tuple with the FetchSubTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetFetchSubTimeOk() (*float32, bool) {
	if o == nil || o.FetchSubTime == nil {
		return nil, false
	}
	return o.FetchSubTime, true
}

// HasFetchSubTime returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasFetchSubTime() bool {
	if o != nil && o.FetchSubTime != nil {
		return true
	}

	return false
}

// SetFetchSubTime gets a reference to the given float32 and assigns it to the FetchSubTime field.
func (o *RealtimeEntryAggregated) SetFetchSubTime(v float32) {
	o.FetchSubTime = &v
}

// GetFetchSubCount returns the FetchSubCount field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetFetchSubCount() int64 {
	if o == nil || o.FetchSubCount == nil {
		var ret int64
		return ret
	}
	return *o.FetchSubCount
}

// GetFetchSubCountOk returns a tuple with the FetchSubCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetFetchSubCountOk() (*int64, bool) {
	if o == nil || o.FetchSubCount == nil {
		return nil, false
	}
	return o.FetchSubCount, true
}

// HasFetchSubCount returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasFetchSubCount() bool {
	if o != nil && o.FetchSubCount != nil {
		return true
	}

	return false
}

// SetFetchSubCount gets a reference to the given int64 and assigns it to the FetchSubCount field.
func (o *RealtimeEntryAggregated) SetFetchSubCount(v int64) {
	o.FetchSubCount = &v
}

// GetPassSubTime returns the PassSubTime field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetPassSubTime() float32 {
	if o == nil || o.PassSubTime == nil {
		var ret float32
		return ret
	}
	return *o.PassSubTime
}

// GetPassSubTimeOk returns a tuple with the PassSubTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetPassSubTimeOk() (*float32, bool) {
	if o == nil || o.PassSubTime == nil {
		return nil, false
	}
	return o.PassSubTime, true
}

// HasPassSubTime returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasPassSubTime() bool {
	if o != nil && o.PassSubTime != nil {
		return true
	}

	return false
}

// SetPassSubTime gets a reference to the given float32 and assigns it to the PassSubTime field.
func (o *RealtimeEntryAggregated) SetPassSubTime(v float32) {
	o.PassSubTime = &v
}

// GetPassSubCount returns the PassSubCount field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetPassSubCount() int64 {
	if o == nil || o.PassSubCount == nil {
		var ret int64
		return ret
	}
	return *o.PassSubCount
}

// GetPassSubCountOk returns a tuple with the PassSubCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetPassSubCountOk() (*int64, bool) {
	if o == nil || o.PassSubCount == nil {
		return nil, false
	}
	return o.PassSubCount, true
}

// HasPassSubCount returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasPassSubCount() bool {
	if o != nil && o.PassSubCount != nil {
		return true
	}

	return false
}

// SetPassSubCount gets a reference to the given int64 and assigns it to the PassSubCount field.
func (o *RealtimeEntryAggregated) SetPassSubCount(v int64) {
	o.PassSubCount = &v
}

// GetPipeSubTime returns the PipeSubTime field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetPipeSubTime() float32 {
	if o == nil || o.PipeSubTime == nil {
		var ret float32
		return ret
	}
	return *o.PipeSubTime
}

// GetPipeSubTimeOk returns a tuple with the PipeSubTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetPipeSubTimeOk() (*float32, bool) {
	if o == nil || o.PipeSubTime == nil {
		return nil, false
	}
	return o.PipeSubTime, true
}

// HasPipeSubTime returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasPipeSubTime() bool {
	if o != nil && o.PipeSubTime != nil {
		return true
	}

	return false
}

// SetPipeSubTime gets a reference to the given float32 and assigns it to the PipeSubTime field.
func (o *RealtimeEntryAggregated) SetPipeSubTime(v float32) {
	o.PipeSubTime = &v
}

// GetPipeSubCount returns the PipeSubCount field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetPipeSubCount() int64 {
	if o == nil || o.PipeSubCount == nil {
		var ret int64
		return ret
	}
	return *o.PipeSubCount
}

// GetPipeSubCountOk returns a tuple with the PipeSubCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetPipeSubCountOk() (*int64, bool) {
	if o == nil || o.PipeSubCount == nil {
		return nil, false
	}
	return o.PipeSubCount, true
}

// HasPipeSubCount returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasPipeSubCount() bool {
	if o != nil && o.PipeSubCount != nil {
		return true
	}

	return false
}

// SetPipeSubCount gets a reference to the given int64 and assigns it to the PipeSubCount field.
func (o *RealtimeEntryAggregated) SetPipeSubCount(v int64) {
	o.PipeSubCount = &v
}

// GetDeliverSubTime returns the DeliverSubTime field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetDeliverSubTime() float32 {
	if o == nil || o.DeliverSubTime == nil {
		var ret float32
		return ret
	}
	return *o.DeliverSubTime
}

// GetDeliverSubTimeOk returns a tuple with the DeliverSubTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetDeliverSubTimeOk() (*float32, bool) {
	if o == nil || o.DeliverSubTime == nil {
		return nil, false
	}
	return o.DeliverSubTime, true
}

// HasDeliverSubTime returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasDeliverSubTime() bool {
	if o != nil && o.DeliverSubTime != nil {
		return true
	}

	return false
}

// SetDeliverSubTime gets a reference to the given float32 and assigns it to the DeliverSubTime field.
func (o *RealtimeEntryAggregated) SetDeliverSubTime(v float32) {
	o.DeliverSubTime = &v
}

// GetDeliverSubCount returns the DeliverSubCount field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetDeliverSubCount() int64 {
	if o == nil || o.DeliverSubCount == nil {
		var ret int64
		return ret
	}
	return *o.DeliverSubCount
}

// GetDeliverSubCountOk returns a tuple with the DeliverSubCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetDeliverSubCountOk() (*int64, bool) {
	if o == nil || o.DeliverSubCount == nil {
		return nil, false
	}
	return o.DeliverSubCount, true
}

// HasDeliverSubCount returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasDeliverSubCount() bool {
	if o != nil && o.DeliverSubCount != nil {
		return true
	}

	return false
}

// SetDeliverSubCount gets a reference to the given int64 and assigns it to the DeliverSubCount field.
func (o *RealtimeEntryAggregated) SetDeliverSubCount(v int64) {
	o.DeliverSubCount = &v
}

// GetErrorSubTime returns the ErrorSubTime field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetErrorSubTime() float32 {
	if o == nil || o.ErrorSubTime == nil {
		var ret float32
		return ret
	}
	return *o.ErrorSubTime
}

// GetErrorSubTimeOk returns a tuple with the ErrorSubTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetErrorSubTimeOk() (*float32, bool) {
	if o == nil || o.ErrorSubTime == nil {
		return nil, false
	}
	return o.ErrorSubTime, true
}

// HasErrorSubTime returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasErrorSubTime() bool {
	if o != nil && o.ErrorSubTime != nil {
		return true
	}

	return false
}

// SetErrorSubTime gets a reference to the given float32 and assigns it to the ErrorSubTime field.
func (o *RealtimeEntryAggregated) SetErrorSubTime(v float32) {
	o.ErrorSubTime = &v
}

// GetErrorSubCount returns the ErrorSubCount field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetErrorSubCount() int64 {
	if o == nil || o.ErrorSubCount == nil {
		var ret int64
		return ret
	}
	return *o.ErrorSubCount
}

// GetErrorSubCountOk returns a tuple with the ErrorSubCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetErrorSubCountOk() (*int64, bool) {
	if o == nil || o.ErrorSubCount == nil {
		return nil, false
	}
	return o.ErrorSubCount, true
}

// HasErrorSubCount returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasErrorSubCount() bool {
	if o != nil && o.ErrorSubCount != nil {
		return true
	}

	return false
}

// SetErrorSubCount gets a reference to the given int64 and assigns it to the ErrorSubCount field.
func (o *RealtimeEntryAggregated) SetErrorSubCount(v int64) {
	o.ErrorSubCount = &v
}

// GetHitSubTime returns the HitSubTime field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetHitSubTime() float32 {
	if o == nil || o.HitSubTime == nil {
		var ret float32
		return ret
	}
	return *o.HitSubTime
}

// GetHitSubTimeOk returns a tuple with the HitSubTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetHitSubTimeOk() (*float32, bool) {
	if o == nil || o.HitSubTime == nil {
		return nil, false
	}
	return o.HitSubTime, true
}

// HasHitSubTime returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasHitSubTime() bool {
	if o != nil && o.HitSubTime != nil {
		return true
	}

	return false
}

// SetHitSubTime gets a reference to the given float32 and assigns it to the HitSubTime field.
func (o *RealtimeEntryAggregated) SetHitSubTime(v float32) {
	o.HitSubTime = &v
}

// GetHitSubCount returns the HitSubCount field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetHitSubCount() int64 {
	if o == nil || o.HitSubCount == nil {
		var ret int64
		return ret
	}
	return *o.HitSubCount
}

// GetHitSubCountOk returns a tuple with the HitSubCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetHitSubCountOk() (*int64, bool) {
	if o == nil || o.HitSubCount == nil {
		return nil, false
	}
	return o.HitSubCount, true
}

// HasHitSubCount returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasHitSubCount() bool {
	if o != nil && o.HitSubCount != nil {
		return true
	}

	return false
}

// SetHitSubCount gets a reference to the given int64 and assigns it to the HitSubCount field.
func (o *RealtimeEntryAggregated) SetHitSubCount(v int64) {
	o.HitSubCount = &v
}

// GetPrehashSubTime returns the PrehashSubTime field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetPrehashSubTime() float32 {
	if o == nil || o.PrehashSubTime == nil {
		var ret float32
		return ret
	}
	return *o.PrehashSubTime
}

// GetPrehashSubTimeOk returns a tuple with the PrehashSubTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetPrehashSubTimeOk() (*float32, bool) {
	if o == nil || o.PrehashSubTime == nil {
		return nil, false
	}
	return o.PrehashSubTime, true
}

// HasPrehashSubTime returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasPrehashSubTime() bool {
	if o != nil && o.PrehashSubTime != nil {
		return true
	}

	return false
}

// SetPrehashSubTime gets a reference to the given float32 and assigns it to the PrehashSubTime field.
func (o *RealtimeEntryAggregated) SetPrehashSubTime(v float32) {
	o.PrehashSubTime = &v
}

// GetPrehashSubCount returns the PrehashSubCount field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetPrehashSubCount() int64 {
	if o == nil || o.PrehashSubCount == nil {
		var ret int64
		return ret
	}
	return *o.PrehashSubCount
}

// GetPrehashSubCountOk returns a tuple with the PrehashSubCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetPrehashSubCountOk() (*int64, bool) {
	if o == nil || o.PrehashSubCount == nil {
		return nil, false
	}
	return o.PrehashSubCount, true
}

// HasPrehashSubCount returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasPrehashSubCount() bool {
	if o != nil && o.PrehashSubCount != nil {
		return true
	}

	return false
}

// SetPrehashSubCount gets a reference to the given int64 and assigns it to the PrehashSubCount field.
func (o *RealtimeEntryAggregated) SetPrehashSubCount(v int64) {
	o.PrehashSubCount = &v
}

// GetPredeliverSubTime returns the PredeliverSubTime field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetPredeliverSubTime() float32 {
	if o == nil || o.PredeliverSubTime == nil {
		var ret float32
		return ret
	}
	return *o.PredeliverSubTime
}

// GetPredeliverSubTimeOk returns a tuple with the PredeliverSubTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetPredeliverSubTimeOk() (*float32, bool) {
	if o == nil || o.PredeliverSubTime == nil {
		return nil, false
	}
	return o.PredeliverSubTime, true
}

// HasPredeliverSubTime returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasPredeliverSubTime() bool {
	if o != nil && o.PredeliverSubTime != nil {
		return true
	}

	return false
}

// SetPredeliverSubTime gets a reference to the given float32 and assigns it to the PredeliverSubTime field.
func (o *RealtimeEntryAggregated) SetPredeliverSubTime(v float32) {
	o.PredeliverSubTime = &v
}

// GetPredeliverSubCount returns the PredeliverSubCount field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetPredeliverSubCount() int64 {
	if o == nil || o.PredeliverSubCount == nil {
		var ret int64
		return ret
	}
	return *o.PredeliverSubCount
}

// GetPredeliverSubCountOk returns a tuple with the PredeliverSubCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetPredeliverSubCountOk() (*int64, bool) {
	if o == nil || o.PredeliverSubCount == nil {
		return nil, false
	}
	return o.PredeliverSubCount, true
}

// HasPredeliverSubCount returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasPredeliverSubCount() bool {
	if o != nil && o.PredeliverSubCount != nil {
		return true
	}

	return false
}

// SetPredeliverSubCount gets a reference to the given int64 and assigns it to the PredeliverSubCount field.
func (o *RealtimeEntryAggregated) SetPredeliverSubCount(v int64) {
	o.PredeliverSubCount = &v
}

// GetHitRespBodyBytes returns the HitRespBodyBytes field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetHitRespBodyBytes() int64 {
	if o == nil || o.HitRespBodyBytes == nil {
		var ret int64
		return ret
	}
	return *o.HitRespBodyBytes
}

// GetHitRespBodyBytesOk returns a tuple with the HitRespBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetHitRespBodyBytesOk() (*int64, bool) {
	if o == nil || o.HitRespBodyBytes == nil {
		return nil, false
	}
	return o.HitRespBodyBytes, true
}

// HasHitRespBodyBytes returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasHitRespBodyBytes() bool {
	if o != nil && o.HitRespBodyBytes != nil {
		return true
	}

	return false
}

// SetHitRespBodyBytes gets a reference to the given int64 and assigns it to the HitRespBodyBytes field.
func (o *RealtimeEntryAggregated) SetHitRespBodyBytes(v int64) {
	o.HitRespBodyBytes = &v
}

// GetMissRespBodyBytes returns the MissRespBodyBytes field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetMissRespBodyBytes() int64 {
	if o == nil || o.MissRespBodyBytes == nil {
		var ret int64
		return ret
	}
	return *o.MissRespBodyBytes
}

// GetMissRespBodyBytesOk returns a tuple with the MissRespBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetMissRespBodyBytesOk() (*int64, bool) {
	if o == nil || o.MissRespBodyBytes == nil {
		return nil, false
	}
	return o.MissRespBodyBytes, true
}

// HasMissRespBodyBytes returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasMissRespBodyBytes() bool {
	if o != nil && o.MissRespBodyBytes != nil {
		return true
	}

	return false
}

// SetMissRespBodyBytes gets a reference to the given int64 and assigns it to the MissRespBodyBytes field.
func (o *RealtimeEntryAggregated) SetMissRespBodyBytes(v int64) {
	o.MissRespBodyBytes = &v
}

// GetPassRespBodyBytes returns the PassRespBodyBytes field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetPassRespBodyBytes() int64 {
	if o == nil || o.PassRespBodyBytes == nil {
		var ret int64
		return ret
	}
	return *o.PassRespBodyBytes
}

// GetPassRespBodyBytesOk returns a tuple with the PassRespBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetPassRespBodyBytesOk() (*int64, bool) {
	if o == nil || o.PassRespBodyBytes == nil {
		return nil, false
	}
	return o.PassRespBodyBytes, true
}

// HasPassRespBodyBytes returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasPassRespBodyBytes() bool {
	if o != nil && o.PassRespBodyBytes != nil {
		return true
	}

	return false
}

// SetPassRespBodyBytes gets a reference to the given int64 and assigns it to the PassRespBodyBytes field.
func (o *RealtimeEntryAggregated) SetPassRespBodyBytes(v int64) {
	o.PassRespBodyBytes = &v
}

// GetComputeReqHeaderBytes returns the ComputeReqHeaderBytes field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetComputeReqHeaderBytes() int64 {
	if o == nil || o.ComputeReqHeaderBytes == nil {
		var ret int64
		return ret
	}
	return *o.ComputeReqHeaderBytes
}

// GetComputeReqHeaderBytesOk returns a tuple with the ComputeReqHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetComputeReqHeaderBytesOk() (*int64, bool) {
	if o == nil || o.ComputeReqHeaderBytes == nil {
		return nil, false
	}
	return o.ComputeReqHeaderBytes, true
}

// HasComputeReqHeaderBytes returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasComputeReqHeaderBytes() bool {
	if o != nil && o.ComputeReqHeaderBytes != nil {
		return true
	}

	return false
}

// SetComputeReqHeaderBytes gets a reference to the given int64 and assigns it to the ComputeReqHeaderBytes field.
func (o *RealtimeEntryAggregated) SetComputeReqHeaderBytes(v int64) {
	o.ComputeReqHeaderBytes = &v
}

// GetComputeReqBodyBytes returns the ComputeReqBodyBytes field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetComputeReqBodyBytes() int64 {
	if o == nil || o.ComputeReqBodyBytes == nil {
		var ret int64
		return ret
	}
	return *o.ComputeReqBodyBytes
}

// GetComputeReqBodyBytesOk returns a tuple with the ComputeReqBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetComputeReqBodyBytesOk() (*int64, bool) {
	if o == nil || o.ComputeReqBodyBytes == nil {
		return nil, false
	}
	return o.ComputeReqBodyBytes, true
}

// HasComputeReqBodyBytes returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasComputeReqBodyBytes() bool {
	if o != nil && o.ComputeReqBodyBytes != nil {
		return true
	}

	return false
}

// SetComputeReqBodyBytes gets a reference to the given int64 and assigns it to the ComputeReqBodyBytes field.
func (o *RealtimeEntryAggregated) SetComputeReqBodyBytes(v int64) {
	o.ComputeReqBodyBytes = &v
}

// GetComputeRespHeaderBytes returns the ComputeRespHeaderBytes field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetComputeRespHeaderBytes() int64 {
	if o == nil || o.ComputeRespHeaderBytes == nil {
		var ret int64
		return ret
	}
	return *o.ComputeRespHeaderBytes
}

// GetComputeRespHeaderBytesOk returns a tuple with the ComputeRespHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetComputeRespHeaderBytesOk() (*int64, bool) {
	if o == nil || o.ComputeRespHeaderBytes == nil {
		return nil, false
	}
	return o.ComputeRespHeaderBytes, true
}

// HasComputeRespHeaderBytes returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasComputeRespHeaderBytes() bool {
	if o != nil && o.ComputeRespHeaderBytes != nil {
		return true
	}

	return false
}

// SetComputeRespHeaderBytes gets a reference to the given int64 and assigns it to the ComputeRespHeaderBytes field.
func (o *RealtimeEntryAggregated) SetComputeRespHeaderBytes(v int64) {
	o.ComputeRespHeaderBytes = &v
}

// GetComputeRespBodyBytes returns the ComputeRespBodyBytes field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetComputeRespBodyBytes() int64 {
	if o == nil || o.ComputeRespBodyBytes == nil {
		var ret int64
		return ret
	}
	return *o.ComputeRespBodyBytes
}

// GetComputeRespBodyBytesOk returns a tuple with the ComputeRespBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetComputeRespBodyBytesOk() (*int64, bool) {
	if o == nil || o.ComputeRespBodyBytes == nil {
		return nil, false
	}
	return o.ComputeRespBodyBytes, true
}

// HasComputeRespBodyBytes returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasComputeRespBodyBytes() bool {
	if o != nil && o.ComputeRespBodyBytes != nil {
		return true
	}

	return false
}

// SetComputeRespBodyBytes gets a reference to the given int64 and assigns it to the ComputeRespBodyBytes field.
func (o *RealtimeEntryAggregated) SetComputeRespBodyBytes(v int64) {
	o.ComputeRespBodyBytes = &v
}

// GetImgvideo returns the Imgvideo field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetImgvideo() int64 {
	if o == nil || o.Imgvideo == nil {
		var ret int64
		return ret
	}
	return *o.Imgvideo
}

// GetImgvideoOk returns a tuple with the Imgvideo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetImgvideoOk() (*int64, bool) {
	if o == nil || o.Imgvideo == nil {
		return nil, false
	}
	return o.Imgvideo, true
}

// HasImgvideo returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasImgvideo() bool {
	if o != nil && o.Imgvideo != nil {
		return true
	}

	return false
}

// SetImgvideo gets a reference to the given int64 and assigns it to the Imgvideo field.
func (o *RealtimeEntryAggregated) SetImgvideo(v int64) {
	o.Imgvideo = &v
}

// GetImgvideoFrames returns the ImgvideoFrames field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetImgvideoFrames() int64 {
	if o == nil || o.ImgvideoFrames == nil {
		var ret int64
		return ret
	}
	return *o.ImgvideoFrames
}

// GetImgvideoFramesOk returns a tuple with the ImgvideoFrames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetImgvideoFramesOk() (*int64, bool) {
	if o == nil || o.ImgvideoFrames == nil {
		return nil, false
	}
	return o.ImgvideoFrames, true
}

// HasImgvideoFrames returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasImgvideoFrames() bool {
	if o != nil && o.ImgvideoFrames != nil {
		return true
	}

	return false
}

// SetImgvideoFrames gets a reference to the given int64 and assigns it to the ImgvideoFrames field.
func (o *RealtimeEntryAggregated) SetImgvideoFrames(v int64) {
	o.ImgvideoFrames = &v
}

// GetImgvideoRespHeaderBytes returns the ImgvideoRespHeaderBytes field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetImgvideoRespHeaderBytes() int64 {
	if o == nil || o.ImgvideoRespHeaderBytes == nil {
		var ret int64
		return ret
	}
	return *o.ImgvideoRespHeaderBytes
}

// GetImgvideoRespHeaderBytesOk returns a tuple with the ImgvideoRespHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetImgvideoRespHeaderBytesOk() (*int64, bool) {
	if o == nil || o.ImgvideoRespHeaderBytes == nil {
		return nil, false
	}
	return o.ImgvideoRespHeaderBytes, true
}

// HasImgvideoRespHeaderBytes returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasImgvideoRespHeaderBytes() bool {
	if o != nil && o.ImgvideoRespHeaderBytes != nil {
		return true
	}

	return false
}

// SetImgvideoRespHeaderBytes gets a reference to the given int64 and assigns it to the ImgvideoRespHeaderBytes field.
func (o *RealtimeEntryAggregated) SetImgvideoRespHeaderBytes(v int64) {
	o.ImgvideoRespHeaderBytes = &v
}

// GetImgvideoRespBodyBytes returns the ImgvideoRespBodyBytes field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetImgvideoRespBodyBytes() int64 {
	if o == nil || o.ImgvideoRespBodyBytes == nil {
		var ret int64
		return ret
	}
	return *o.ImgvideoRespBodyBytes
}

// GetImgvideoRespBodyBytesOk returns a tuple with the ImgvideoRespBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetImgvideoRespBodyBytesOk() (*int64, bool) {
	if o == nil || o.ImgvideoRespBodyBytes == nil {
		return nil, false
	}
	return o.ImgvideoRespBodyBytes, true
}

// HasImgvideoRespBodyBytes returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasImgvideoRespBodyBytes() bool {
	if o != nil && o.ImgvideoRespBodyBytes != nil {
		return true
	}

	return false
}

// SetImgvideoRespBodyBytes gets a reference to the given int64 and assigns it to the ImgvideoRespBodyBytes field.
func (o *RealtimeEntryAggregated) SetImgvideoRespBodyBytes(v int64) {
	o.ImgvideoRespBodyBytes = &v
}

// GetImgvideoShield returns the ImgvideoShield field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetImgvideoShield() int64 {
	if o == nil || o.ImgvideoShield == nil {
		var ret int64
		return ret
	}
	return *o.ImgvideoShield
}

// GetImgvideoShieldOk returns a tuple with the ImgvideoShield field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetImgvideoShieldOk() (*int64, bool) {
	if o == nil || o.ImgvideoShield == nil {
		return nil, false
	}
	return o.ImgvideoShield, true
}

// HasImgvideoShield returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasImgvideoShield() bool {
	if o != nil && o.ImgvideoShield != nil {
		return true
	}

	return false
}

// SetImgvideoShield gets a reference to the given int64 and assigns it to the ImgvideoShield field.
func (o *RealtimeEntryAggregated) SetImgvideoShield(v int64) {
	o.ImgvideoShield = &v
}

// GetImgvideoShieldFrames returns the ImgvideoShieldFrames field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetImgvideoShieldFrames() int64 {
	if o == nil || o.ImgvideoShieldFrames == nil {
		var ret int64
		return ret
	}
	return *o.ImgvideoShieldFrames
}

// GetImgvideoShieldFramesOk returns a tuple with the ImgvideoShieldFrames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetImgvideoShieldFramesOk() (*int64, bool) {
	if o == nil || o.ImgvideoShieldFrames == nil {
		return nil, false
	}
	return o.ImgvideoShieldFrames, true
}

// HasImgvideoShieldFrames returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasImgvideoShieldFrames() bool {
	if o != nil && o.ImgvideoShieldFrames != nil {
		return true
	}

	return false
}

// SetImgvideoShieldFrames gets a reference to the given int64 and assigns it to the ImgvideoShieldFrames field.
func (o *RealtimeEntryAggregated) SetImgvideoShieldFrames(v int64) {
	o.ImgvideoShieldFrames = &v
}

// GetImgvideoShieldRespHeaderBytes returns the ImgvideoShieldRespHeaderBytes field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetImgvideoShieldRespHeaderBytes() int64 {
	if o == nil || o.ImgvideoShieldRespHeaderBytes == nil {
		var ret int64
		return ret
	}
	return *o.ImgvideoShieldRespHeaderBytes
}

// GetImgvideoShieldRespHeaderBytesOk returns a tuple with the ImgvideoShieldRespHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetImgvideoShieldRespHeaderBytesOk() (*int64, bool) {
	if o == nil || o.ImgvideoShieldRespHeaderBytes == nil {
		return nil, false
	}
	return o.ImgvideoShieldRespHeaderBytes, true
}

// HasImgvideoShieldRespHeaderBytes returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasImgvideoShieldRespHeaderBytes() bool {
	if o != nil && o.ImgvideoShieldRespHeaderBytes != nil {
		return true
	}

	return false
}

// SetImgvideoShieldRespHeaderBytes gets a reference to the given int64 and assigns it to the ImgvideoShieldRespHeaderBytes field.
func (o *RealtimeEntryAggregated) SetImgvideoShieldRespHeaderBytes(v int64) {
	o.ImgvideoShieldRespHeaderBytes = &v
}

// GetImgvideoShieldRespBodyBytes returns the ImgvideoShieldRespBodyBytes field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetImgvideoShieldRespBodyBytes() int64 {
	if o == nil || o.ImgvideoShieldRespBodyBytes == nil {
		var ret int64
		return ret
	}
	return *o.ImgvideoShieldRespBodyBytes
}

// GetImgvideoShieldRespBodyBytesOk returns a tuple with the ImgvideoShieldRespBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetImgvideoShieldRespBodyBytesOk() (*int64, bool) {
	if o == nil || o.ImgvideoShieldRespBodyBytes == nil {
		return nil, false
	}
	return o.ImgvideoShieldRespBodyBytes, true
}

// HasImgvideoShieldRespBodyBytes returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasImgvideoShieldRespBodyBytes() bool {
	if o != nil && o.ImgvideoShieldRespBodyBytes != nil {
		return true
	}

	return false
}

// SetImgvideoShieldRespBodyBytes gets a reference to the given int64 and assigns it to the ImgvideoShieldRespBodyBytes field.
func (o *RealtimeEntryAggregated) SetImgvideoShieldRespBodyBytes(v int64) {
	o.ImgvideoShieldRespBodyBytes = &v
}

// GetLogBytes returns the LogBytes field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetLogBytes() int64 {
	if o == nil || o.LogBytes == nil {
		var ret int64
		return ret
	}
	return *o.LogBytes
}

// GetLogBytesOk returns a tuple with the LogBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetLogBytesOk() (*int64, bool) {
	if o == nil || o.LogBytes == nil {
		return nil, false
	}
	return o.LogBytes, true
}

// HasLogBytes returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasLogBytes() bool {
	if o != nil && o.LogBytes != nil {
		return true
	}

	return false
}

// SetLogBytes gets a reference to the given int64 and assigns it to the LogBytes field.
func (o *RealtimeEntryAggregated) SetLogBytes(v int64) {
	o.LogBytes = &v
}

// GetEdgeRequests returns the EdgeRequests field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetEdgeRequests() int64 {
	if o == nil || o.EdgeRequests == nil {
		var ret int64
		return ret
	}
	return *o.EdgeRequests
}

// GetEdgeRequestsOk returns a tuple with the EdgeRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetEdgeRequestsOk() (*int64, bool) {
	if o == nil || o.EdgeRequests == nil {
		return nil, false
	}
	return o.EdgeRequests, true
}

// HasEdgeRequests returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasEdgeRequests() bool {
	if o != nil && o.EdgeRequests != nil {
		return true
	}

	return false
}

// SetEdgeRequests gets a reference to the given int64 and assigns it to the EdgeRequests field.
func (o *RealtimeEntryAggregated) SetEdgeRequests(v int64) {
	o.EdgeRequests = &v
}

// GetEdgeRespHeaderBytes returns the EdgeRespHeaderBytes field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetEdgeRespHeaderBytes() int64 {
	if o == nil || o.EdgeRespHeaderBytes == nil {
		var ret int64
		return ret
	}
	return *o.EdgeRespHeaderBytes
}

// GetEdgeRespHeaderBytesOk returns a tuple with the EdgeRespHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetEdgeRespHeaderBytesOk() (*int64, bool) {
	if o == nil || o.EdgeRespHeaderBytes == nil {
		return nil, false
	}
	return o.EdgeRespHeaderBytes, true
}

// HasEdgeRespHeaderBytes returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasEdgeRespHeaderBytes() bool {
	if o != nil && o.EdgeRespHeaderBytes != nil {
		return true
	}

	return false
}

// SetEdgeRespHeaderBytes gets a reference to the given int64 and assigns it to the EdgeRespHeaderBytes field.
func (o *RealtimeEntryAggregated) SetEdgeRespHeaderBytes(v int64) {
	o.EdgeRespHeaderBytes = &v
}

// GetEdgeRespBodyBytes returns the EdgeRespBodyBytes field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetEdgeRespBodyBytes() int64 {
	if o == nil || o.EdgeRespBodyBytes == nil {
		var ret int64
		return ret
	}
	return *o.EdgeRespBodyBytes
}

// GetEdgeRespBodyBytesOk returns a tuple with the EdgeRespBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetEdgeRespBodyBytesOk() (*int64, bool) {
	if o == nil || o.EdgeRespBodyBytes == nil {
		return nil, false
	}
	return o.EdgeRespBodyBytes, true
}

// HasEdgeRespBodyBytes returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasEdgeRespBodyBytes() bool {
	if o != nil && o.EdgeRespBodyBytes != nil {
		return true
	}

	return false
}

// SetEdgeRespBodyBytes gets a reference to the given int64 and assigns it to the EdgeRespBodyBytes field.
func (o *RealtimeEntryAggregated) SetEdgeRespBodyBytes(v int64) {
	o.EdgeRespBodyBytes = &v
}

// GetOriginRevalidations returns the OriginRevalidations field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetOriginRevalidations() int64 {
	if o == nil || o.OriginRevalidations == nil {
		var ret int64
		return ret
	}
	return *o.OriginRevalidations
}

// GetOriginRevalidationsOk returns a tuple with the OriginRevalidations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetOriginRevalidationsOk() (*int64, bool) {
	if o == nil || o.OriginRevalidations == nil {
		return nil, false
	}
	return o.OriginRevalidations, true
}

// HasOriginRevalidations returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasOriginRevalidations() bool {
	if o != nil && o.OriginRevalidations != nil {
		return true
	}

	return false
}

// SetOriginRevalidations gets a reference to the given int64 and assigns it to the OriginRevalidations field.
func (o *RealtimeEntryAggregated) SetOriginRevalidations(v int64) {
	o.OriginRevalidations = &v
}

// GetOriginFetches returns the OriginFetches field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetOriginFetches() int64 {
	if o == nil || o.OriginFetches == nil {
		var ret int64
		return ret
	}
	return *o.OriginFetches
}

// GetOriginFetchesOk returns a tuple with the OriginFetches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetOriginFetchesOk() (*int64, bool) {
	if o == nil || o.OriginFetches == nil {
		return nil, false
	}
	return o.OriginFetches, true
}

// HasOriginFetches returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasOriginFetches() bool {
	if o != nil && o.OriginFetches != nil {
		return true
	}

	return false
}

// SetOriginFetches gets a reference to the given int64 and assigns it to the OriginFetches field.
func (o *RealtimeEntryAggregated) SetOriginFetches(v int64) {
	o.OriginFetches = &v
}

// GetOriginFetchHeaderBytes returns the OriginFetchHeaderBytes field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetOriginFetchHeaderBytes() int64 {
	if o == nil || o.OriginFetchHeaderBytes == nil {
		var ret int64
		return ret
	}
	return *o.OriginFetchHeaderBytes
}

// GetOriginFetchHeaderBytesOk returns a tuple with the OriginFetchHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetOriginFetchHeaderBytesOk() (*int64, bool) {
	if o == nil || o.OriginFetchHeaderBytes == nil {
		return nil, false
	}
	return o.OriginFetchHeaderBytes, true
}

// HasOriginFetchHeaderBytes returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasOriginFetchHeaderBytes() bool {
	if o != nil && o.OriginFetchHeaderBytes != nil {
		return true
	}

	return false
}

// SetOriginFetchHeaderBytes gets a reference to the given int64 and assigns it to the OriginFetchHeaderBytes field.
func (o *RealtimeEntryAggregated) SetOriginFetchHeaderBytes(v int64) {
	o.OriginFetchHeaderBytes = &v
}

// GetOriginFetchBodyBytes returns the OriginFetchBodyBytes field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetOriginFetchBodyBytes() int64 {
	if o == nil || o.OriginFetchBodyBytes == nil {
		var ret int64
		return ret
	}
	return *o.OriginFetchBodyBytes
}

// GetOriginFetchBodyBytesOk returns a tuple with the OriginFetchBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetOriginFetchBodyBytesOk() (*int64, bool) {
	if o == nil || o.OriginFetchBodyBytes == nil {
		return nil, false
	}
	return o.OriginFetchBodyBytes, true
}

// HasOriginFetchBodyBytes returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasOriginFetchBodyBytes() bool {
	if o != nil && o.OriginFetchBodyBytes != nil {
		return true
	}

	return false
}

// SetOriginFetchBodyBytes gets a reference to the given int64 and assigns it to the OriginFetchBodyBytes field.
func (o *RealtimeEntryAggregated) SetOriginFetchBodyBytes(v int64) {
	o.OriginFetchBodyBytes = &v
}

// GetOriginFetchRespHeaderBytes returns the OriginFetchRespHeaderBytes field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetOriginFetchRespHeaderBytes() int64 {
	if o == nil || o.OriginFetchRespHeaderBytes == nil {
		var ret int64
		return ret
	}
	return *o.OriginFetchRespHeaderBytes
}

// GetOriginFetchRespHeaderBytesOk returns a tuple with the OriginFetchRespHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetOriginFetchRespHeaderBytesOk() (*int64, bool) {
	if o == nil || o.OriginFetchRespHeaderBytes == nil {
		return nil, false
	}
	return o.OriginFetchRespHeaderBytes, true
}

// HasOriginFetchRespHeaderBytes returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasOriginFetchRespHeaderBytes() bool {
	if o != nil && o.OriginFetchRespHeaderBytes != nil {
		return true
	}

	return false
}

// SetOriginFetchRespHeaderBytes gets a reference to the given int64 and assigns it to the OriginFetchRespHeaderBytes field.
func (o *RealtimeEntryAggregated) SetOriginFetchRespHeaderBytes(v int64) {
	o.OriginFetchRespHeaderBytes = &v
}

// GetOriginFetchRespBodyBytes returns the OriginFetchRespBodyBytes field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetOriginFetchRespBodyBytes() int64 {
	if o == nil || o.OriginFetchRespBodyBytes == nil {
		var ret int64
		return ret
	}
	return *o.OriginFetchRespBodyBytes
}

// GetOriginFetchRespBodyBytesOk returns a tuple with the OriginFetchRespBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetOriginFetchRespBodyBytesOk() (*int64, bool) {
	if o == nil || o.OriginFetchRespBodyBytes == nil {
		return nil, false
	}
	return o.OriginFetchRespBodyBytes, true
}

// HasOriginFetchRespBodyBytes returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasOriginFetchRespBodyBytes() bool {
	if o != nil && o.OriginFetchRespBodyBytes != nil {
		return true
	}

	return false
}

// SetOriginFetchRespBodyBytes gets a reference to the given int64 and assigns it to the OriginFetchRespBodyBytes field.
func (o *RealtimeEntryAggregated) SetOriginFetchRespBodyBytes(v int64) {
	o.OriginFetchRespBodyBytes = &v
}

// GetShieldRevalidations returns the ShieldRevalidations field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetShieldRevalidations() int64 {
	if o == nil || o.ShieldRevalidations == nil {
		var ret int64
		return ret
	}
	return *o.ShieldRevalidations
}

// GetShieldRevalidationsOk returns a tuple with the ShieldRevalidations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetShieldRevalidationsOk() (*int64, bool) {
	if o == nil || o.ShieldRevalidations == nil {
		return nil, false
	}
	return o.ShieldRevalidations, true
}

// HasShieldRevalidations returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasShieldRevalidations() bool {
	if o != nil && o.ShieldRevalidations != nil {
		return true
	}

	return false
}

// SetShieldRevalidations gets a reference to the given int64 and assigns it to the ShieldRevalidations field.
func (o *RealtimeEntryAggregated) SetShieldRevalidations(v int64) {
	o.ShieldRevalidations = &v
}

// GetShieldFetches returns the ShieldFetches field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetShieldFetches() int64 {
	if o == nil || o.ShieldFetches == nil {
		var ret int64
		return ret
	}
	return *o.ShieldFetches
}

// GetShieldFetchesOk returns a tuple with the ShieldFetches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetShieldFetchesOk() (*int64, bool) {
	if o == nil || o.ShieldFetches == nil {
		return nil, false
	}
	return o.ShieldFetches, true
}

// HasShieldFetches returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasShieldFetches() bool {
	if o != nil && o.ShieldFetches != nil {
		return true
	}

	return false
}

// SetShieldFetches gets a reference to the given int64 and assigns it to the ShieldFetches field.
func (o *RealtimeEntryAggregated) SetShieldFetches(v int64) {
	o.ShieldFetches = &v
}

// GetShieldFetchHeaderBytes returns the ShieldFetchHeaderBytes field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetShieldFetchHeaderBytes() int64 {
	if o == nil || o.ShieldFetchHeaderBytes == nil {
		var ret int64
		return ret
	}
	return *o.ShieldFetchHeaderBytes
}

// GetShieldFetchHeaderBytesOk returns a tuple with the ShieldFetchHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetShieldFetchHeaderBytesOk() (*int64, bool) {
	if o == nil || o.ShieldFetchHeaderBytes == nil {
		return nil, false
	}
	return o.ShieldFetchHeaderBytes, true
}

// HasShieldFetchHeaderBytes returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasShieldFetchHeaderBytes() bool {
	if o != nil && o.ShieldFetchHeaderBytes != nil {
		return true
	}

	return false
}

// SetShieldFetchHeaderBytes gets a reference to the given int64 and assigns it to the ShieldFetchHeaderBytes field.
func (o *RealtimeEntryAggregated) SetShieldFetchHeaderBytes(v int64) {
	o.ShieldFetchHeaderBytes = &v
}

// GetShieldFetchBodyBytes returns the ShieldFetchBodyBytes field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetShieldFetchBodyBytes() int64 {
	if o == nil || o.ShieldFetchBodyBytes == nil {
		var ret int64
		return ret
	}
	return *o.ShieldFetchBodyBytes
}

// GetShieldFetchBodyBytesOk returns a tuple with the ShieldFetchBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetShieldFetchBodyBytesOk() (*int64, bool) {
	if o == nil || o.ShieldFetchBodyBytes == nil {
		return nil, false
	}
	return o.ShieldFetchBodyBytes, true
}

// HasShieldFetchBodyBytes returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasShieldFetchBodyBytes() bool {
	if o != nil && o.ShieldFetchBodyBytes != nil {
		return true
	}

	return false
}

// SetShieldFetchBodyBytes gets a reference to the given int64 and assigns it to the ShieldFetchBodyBytes field.
func (o *RealtimeEntryAggregated) SetShieldFetchBodyBytes(v int64) {
	o.ShieldFetchBodyBytes = &v
}

// GetShieldFetchRespHeaderBytes returns the ShieldFetchRespHeaderBytes field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetShieldFetchRespHeaderBytes() int64 {
	if o == nil || o.ShieldFetchRespHeaderBytes == nil {
		var ret int64
		return ret
	}
	return *o.ShieldFetchRespHeaderBytes
}

// GetShieldFetchRespHeaderBytesOk returns a tuple with the ShieldFetchRespHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetShieldFetchRespHeaderBytesOk() (*int64, bool) {
	if o == nil || o.ShieldFetchRespHeaderBytes == nil {
		return nil, false
	}
	return o.ShieldFetchRespHeaderBytes, true
}

// HasShieldFetchRespHeaderBytes returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasShieldFetchRespHeaderBytes() bool {
	if o != nil && o.ShieldFetchRespHeaderBytes != nil {
		return true
	}

	return false
}

// SetShieldFetchRespHeaderBytes gets a reference to the given int64 and assigns it to the ShieldFetchRespHeaderBytes field.
func (o *RealtimeEntryAggregated) SetShieldFetchRespHeaderBytes(v int64) {
	o.ShieldFetchRespHeaderBytes = &v
}

// GetShieldFetchRespBodyBytes returns the ShieldFetchRespBodyBytes field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetShieldFetchRespBodyBytes() int64 {
	if o == nil || o.ShieldFetchRespBodyBytes == nil {
		var ret int64
		return ret
	}
	return *o.ShieldFetchRespBodyBytes
}

// GetShieldFetchRespBodyBytesOk returns a tuple with the ShieldFetchRespBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetShieldFetchRespBodyBytesOk() (*int64, bool) {
	if o == nil || o.ShieldFetchRespBodyBytes == nil {
		return nil, false
	}
	return o.ShieldFetchRespBodyBytes, true
}

// HasShieldFetchRespBodyBytes returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasShieldFetchRespBodyBytes() bool {
	if o != nil && o.ShieldFetchRespBodyBytes != nil {
		return true
	}

	return false
}

// SetShieldFetchRespBodyBytes gets a reference to the given int64 and assigns it to the ShieldFetchRespBodyBytes field.
func (o *RealtimeEntryAggregated) SetShieldFetchRespBodyBytes(v int64) {
	o.ShieldFetchRespBodyBytes = &v
}

// GetSegblockOriginFetches returns the SegblockOriginFetches field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetSegblockOriginFetches() int64 {
	if o == nil || o.SegblockOriginFetches == nil {
		var ret int64
		return ret
	}
	return *o.SegblockOriginFetches
}

// GetSegblockOriginFetchesOk returns a tuple with the SegblockOriginFetches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetSegblockOriginFetchesOk() (*int64, bool) {
	if o == nil || o.SegblockOriginFetches == nil {
		return nil, false
	}
	return o.SegblockOriginFetches, true
}

// HasSegblockOriginFetches returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasSegblockOriginFetches() bool {
	if o != nil && o.SegblockOriginFetches != nil {
		return true
	}

	return false
}

// SetSegblockOriginFetches gets a reference to the given int64 and assigns it to the SegblockOriginFetches field.
func (o *RealtimeEntryAggregated) SetSegblockOriginFetches(v int64) {
	o.SegblockOriginFetches = &v
}

// GetSegblockShieldFetches returns the SegblockShieldFetches field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetSegblockShieldFetches() int64 {
	if o == nil || o.SegblockShieldFetches == nil {
		var ret int64
		return ret
	}
	return *o.SegblockShieldFetches
}

// GetSegblockShieldFetchesOk returns a tuple with the SegblockShieldFetches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetSegblockShieldFetchesOk() (*int64, bool) {
	if o == nil || o.SegblockShieldFetches == nil {
		return nil, false
	}
	return o.SegblockShieldFetches, true
}

// HasSegblockShieldFetches returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasSegblockShieldFetches() bool {
	if o != nil && o.SegblockShieldFetches != nil {
		return true
	}

	return false
}

// SetSegblockShieldFetches gets a reference to the given int64 and assigns it to the SegblockShieldFetches field.
func (o *RealtimeEntryAggregated) SetSegblockShieldFetches(v int64) {
	o.SegblockShieldFetches = &v
}

// GetComputeRespStatus1xx returns the ComputeRespStatus1xx field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetComputeRespStatus1xx() int64 {
	if o == nil || o.ComputeRespStatus1xx == nil {
		var ret int64
		return ret
	}
	return *o.ComputeRespStatus1xx
}

// GetComputeRespStatus1xxOk returns a tuple with the ComputeRespStatus1xx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetComputeRespStatus1xxOk() (*int64, bool) {
	if o == nil || o.ComputeRespStatus1xx == nil {
		return nil, false
	}
	return o.ComputeRespStatus1xx, true
}

// HasComputeRespStatus1xx returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasComputeRespStatus1xx() bool {
	if o != nil && o.ComputeRespStatus1xx != nil {
		return true
	}

	return false
}

// SetComputeRespStatus1xx gets a reference to the given int64 and assigns it to the ComputeRespStatus1xx field.
func (o *RealtimeEntryAggregated) SetComputeRespStatus1xx(v int64) {
	o.ComputeRespStatus1xx = &v
}

// GetComputeRespStatus2xx returns the ComputeRespStatus2xx field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetComputeRespStatus2xx() int64 {
	if o == nil || o.ComputeRespStatus2xx == nil {
		var ret int64
		return ret
	}
	return *o.ComputeRespStatus2xx
}

// GetComputeRespStatus2xxOk returns a tuple with the ComputeRespStatus2xx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetComputeRespStatus2xxOk() (*int64, bool) {
	if o == nil || o.ComputeRespStatus2xx == nil {
		return nil, false
	}
	return o.ComputeRespStatus2xx, true
}

// HasComputeRespStatus2xx returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasComputeRespStatus2xx() bool {
	if o != nil && o.ComputeRespStatus2xx != nil {
		return true
	}

	return false
}

// SetComputeRespStatus2xx gets a reference to the given int64 and assigns it to the ComputeRespStatus2xx field.
func (o *RealtimeEntryAggregated) SetComputeRespStatus2xx(v int64) {
	o.ComputeRespStatus2xx = &v
}

// GetComputeRespStatus3xx returns the ComputeRespStatus3xx field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetComputeRespStatus3xx() int64 {
	if o == nil || o.ComputeRespStatus3xx == nil {
		var ret int64
		return ret
	}
	return *o.ComputeRespStatus3xx
}

// GetComputeRespStatus3xxOk returns a tuple with the ComputeRespStatus3xx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetComputeRespStatus3xxOk() (*int64, bool) {
	if o == nil || o.ComputeRespStatus3xx == nil {
		return nil, false
	}
	return o.ComputeRespStatus3xx, true
}

// HasComputeRespStatus3xx returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasComputeRespStatus3xx() bool {
	if o != nil && o.ComputeRespStatus3xx != nil {
		return true
	}

	return false
}

// SetComputeRespStatus3xx gets a reference to the given int64 and assigns it to the ComputeRespStatus3xx field.
func (o *RealtimeEntryAggregated) SetComputeRespStatus3xx(v int64) {
	o.ComputeRespStatus3xx = &v
}

// GetComputeRespStatus4xx returns the ComputeRespStatus4xx field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetComputeRespStatus4xx() int64 {
	if o == nil || o.ComputeRespStatus4xx == nil {
		var ret int64
		return ret
	}
	return *o.ComputeRespStatus4xx
}

// GetComputeRespStatus4xxOk returns a tuple with the ComputeRespStatus4xx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetComputeRespStatus4xxOk() (*int64, bool) {
	if o == nil || o.ComputeRespStatus4xx == nil {
		return nil, false
	}
	return o.ComputeRespStatus4xx, true
}

// HasComputeRespStatus4xx returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasComputeRespStatus4xx() bool {
	if o != nil && o.ComputeRespStatus4xx != nil {
		return true
	}

	return false
}

// SetComputeRespStatus4xx gets a reference to the given int64 and assigns it to the ComputeRespStatus4xx field.
func (o *RealtimeEntryAggregated) SetComputeRespStatus4xx(v int64) {
	o.ComputeRespStatus4xx = &v
}

// GetComputeRespStatus5xx returns the ComputeRespStatus5xx field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetComputeRespStatus5xx() int64 {
	if o == nil || o.ComputeRespStatus5xx == nil {
		var ret int64
		return ret
	}
	return *o.ComputeRespStatus5xx
}

// GetComputeRespStatus5xxOk returns a tuple with the ComputeRespStatus5xx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetComputeRespStatus5xxOk() (*int64, bool) {
	if o == nil || o.ComputeRespStatus5xx == nil {
		return nil, false
	}
	return o.ComputeRespStatus5xx, true
}

// HasComputeRespStatus5xx returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasComputeRespStatus5xx() bool {
	if o != nil && o.ComputeRespStatus5xx != nil {
		return true
	}

	return false
}

// SetComputeRespStatus5xx gets a reference to the given int64 and assigns it to the ComputeRespStatus5xx field.
func (o *RealtimeEntryAggregated) SetComputeRespStatus5xx(v int64) {
	o.ComputeRespStatus5xx = &v
}

// GetEdgeHitRequests returns the EdgeHitRequests field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetEdgeHitRequests() int64 {
	if o == nil || o.EdgeHitRequests == nil {
		var ret int64
		return ret
	}
	return *o.EdgeHitRequests
}

// GetEdgeHitRequestsOk returns a tuple with the EdgeHitRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetEdgeHitRequestsOk() (*int64, bool) {
	if o == nil || o.EdgeHitRequests == nil {
		return nil, false
	}
	return o.EdgeHitRequests, true
}

// HasEdgeHitRequests returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasEdgeHitRequests() bool {
	if o != nil && o.EdgeHitRequests != nil {
		return true
	}

	return false
}

// SetEdgeHitRequests gets a reference to the given int64 and assigns it to the EdgeHitRequests field.
func (o *RealtimeEntryAggregated) SetEdgeHitRequests(v int64) {
	o.EdgeHitRequests = &v
}

// GetEdgeMissRequests returns the EdgeMissRequests field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetEdgeMissRequests() int64 {
	if o == nil || o.EdgeMissRequests == nil {
		var ret int64
		return ret
	}
	return *o.EdgeMissRequests
}

// GetEdgeMissRequestsOk returns a tuple with the EdgeMissRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetEdgeMissRequestsOk() (*int64, bool) {
	if o == nil || o.EdgeMissRequests == nil {
		return nil, false
	}
	return o.EdgeMissRequests, true
}

// HasEdgeMissRequests returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasEdgeMissRequests() bool {
	if o != nil && o.EdgeMissRequests != nil {
		return true
	}

	return false
}

// SetEdgeMissRequests gets a reference to the given int64 and assigns it to the EdgeMissRequests field.
func (o *RealtimeEntryAggregated) SetEdgeMissRequests(v int64) {
	o.EdgeMissRequests = &v
}

// GetComputeBereqHeaderBytes returns the ComputeBereqHeaderBytes field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetComputeBereqHeaderBytes() int64 {
	if o == nil || o.ComputeBereqHeaderBytes == nil {
		var ret int64
		return ret
	}
	return *o.ComputeBereqHeaderBytes
}

// GetComputeBereqHeaderBytesOk returns a tuple with the ComputeBereqHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetComputeBereqHeaderBytesOk() (*int64, bool) {
	if o == nil || o.ComputeBereqHeaderBytes == nil {
		return nil, false
	}
	return o.ComputeBereqHeaderBytes, true
}

// HasComputeBereqHeaderBytes returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasComputeBereqHeaderBytes() bool {
	if o != nil && o.ComputeBereqHeaderBytes != nil {
		return true
	}

	return false
}

// SetComputeBereqHeaderBytes gets a reference to the given int64 and assigns it to the ComputeBereqHeaderBytes field.
func (o *RealtimeEntryAggregated) SetComputeBereqHeaderBytes(v int64) {
	o.ComputeBereqHeaderBytes = &v
}

// GetComputeBereqBodyBytes returns the ComputeBereqBodyBytes field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetComputeBereqBodyBytes() int64 {
	if o == nil || o.ComputeBereqBodyBytes == nil {
		var ret int64
		return ret
	}
	return *o.ComputeBereqBodyBytes
}

// GetComputeBereqBodyBytesOk returns a tuple with the ComputeBereqBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetComputeBereqBodyBytesOk() (*int64, bool) {
	if o == nil || o.ComputeBereqBodyBytes == nil {
		return nil, false
	}
	return o.ComputeBereqBodyBytes, true
}

// HasComputeBereqBodyBytes returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasComputeBereqBodyBytes() bool {
	if o != nil && o.ComputeBereqBodyBytes != nil {
		return true
	}

	return false
}

// SetComputeBereqBodyBytes gets a reference to the given int64 and assigns it to the ComputeBereqBodyBytes field.
func (o *RealtimeEntryAggregated) SetComputeBereqBodyBytes(v int64) {
	o.ComputeBereqBodyBytes = &v
}

// GetComputeBerespHeaderBytes returns the ComputeBerespHeaderBytes field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetComputeBerespHeaderBytes() int64 {
	if o == nil || o.ComputeBerespHeaderBytes == nil {
		var ret int64
		return ret
	}
	return *o.ComputeBerespHeaderBytes
}

// GetComputeBerespHeaderBytesOk returns a tuple with the ComputeBerespHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetComputeBerespHeaderBytesOk() (*int64, bool) {
	if o == nil || o.ComputeBerespHeaderBytes == nil {
		return nil, false
	}
	return o.ComputeBerespHeaderBytes, true
}

// HasComputeBerespHeaderBytes returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasComputeBerespHeaderBytes() bool {
	if o != nil && o.ComputeBerespHeaderBytes != nil {
		return true
	}

	return false
}

// SetComputeBerespHeaderBytes gets a reference to the given int64 and assigns it to the ComputeBerespHeaderBytes field.
func (o *RealtimeEntryAggregated) SetComputeBerespHeaderBytes(v int64) {
	o.ComputeBerespHeaderBytes = &v
}

// GetComputeBerespBodyBytes returns the ComputeBerespBodyBytes field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetComputeBerespBodyBytes() int64 {
	if o == nil || o.ComputeBerespBodyBytes == nil {
		var ret int64
		return ret
	}
	return *o.ComputeBerespBodyBytes
}

// GetComputeBerespBodyBytesOk returns a tuple with the ComputeBerespBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetComputeBerespBodyBytesOk() (*int64, bool) {
	if o == nil || o.ComputeBerespBodyBytes == nil {
		return nil, false
	}
	return o.ComputeBerespBodyBytes, true
}

// HasComputeBerespBodyBytes returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasComputeBerespBodyBytes() bool {
	if o != nil && o.ComputeBerespBodyBytes != nil {
		return true
	}

	return false
}

// SetComputeBerespBodyBytes gets a reference to the given int64 and assigns it to the ComputeBerespBodyBytes field.
func (o *RealtimeEntryAggregated) SetComputeBerespBodyBytes(v int64) {
	o.ComputeBerespBodyBytes = &v
}

// GetOriginCacheFetches returns the OriginCacheFetches field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetOriginCacheFetches() int64 {
	if o == nil || o.OriginCacheFetches == nil {
		var ret int64
		return ret
	}
	return *o.OriginCacheFetches
}

// GetOriginCacheFetchesOk returns a tuple with the OriginCacheFetches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetOriginCacheFetchesOk() (*int64, bool) {
	if o == nil || o.OriginCacheFetches == nil {
		return nil, false
	}
	return o.OriginCacheFetches, true
}

// HasOriginCacheFetches returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasOriginCacheFetches() bool {
	if o != nil && o.OriginCacheFetches != nil {
		return true
	}

	return false
}

// SetOriginCacheFetches gets a reference to the given int64 and assigns it to the OriginCacheFetches field.
func (o *RealtimeEntryAggregated) SetOriginCacheFetches(v int64) {
	o.OriginCacheFetches = &v
}

// GetShieldCacheFetches returns the ShieldCacheFetches field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetShieldCacheFetches() int64 {
	if o == nil || o.ShieldCacheFetches == nil {
		var ret int64
		return ret
	}
	return *o.ShieldCacheFetches
}

// GetShieldCacheFetchesOk returns a tuple with the ShieldCacheFetches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetShieldCacheFetchesOk() (*int64, bool) {
	if o == nil || o.ShieldCacheFetches == nil {
		return nil, false
	}
	return o.ShieldCacheFetches, true
}

// HasShieldCacheFetches returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasShieldCacheFetches() bool {
	if o != nil && o.ShieldCacheFetches != nil {
		return true
	}

	return false
}

// SetShieldCacheFetches gets a reference to the given int64 and assigns it to the ShieldCacheFetches field.
func (o *RealtimeEntryAggregated) SetShieldCacheFetches(v int64) {
	o.ShieldCacheFetches = &v
}

// GetComputeBereqs returns the ComputeBereqs field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetComputeBereqs() int64 {
	if o == nil || o.ComputeBereqs == nil {
		var ret int64
		return ret
	}
	return *o.ComputeBereqs
}

// GetComputeBereqsOk returns a tuple with the ComputeBereqs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetComputeBereqsOk() (*int64, bool) {
	if o == nil || o.ComputeBereqs == nil {
		return nil, false
	}
	return o.ComputeBereqs, true
}

// HasComputeBereqs returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasComputeBereqs() bool {
	if o != nil && o.ComputeBereqs != nil {
		return true
	}

	return false
}

// SetComputeBereqs gets a reference to the given int64 and assigns it to the ComputeBereqs field.
func (o *RealtimeEntryAggregated) SetComputeBereqs(v int64) {
	o.ComputeBereqs = &v
}

// GetComputeBereqErrors returns the ComputeBereqErrors field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetComputeBereqErrors() int64 {
	if o == nil || o.ComputeBereqErrors == nil {
		var ret int64
		return ret
	}
	return *o.ComputeBereqErrors
}

// GetComputeBereqErrorsOk returns a tuple with the ComputeBereqErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetComputeBereqErrorsOk() (*int64, bool) {
	if o == nil || o.ComputeBereqErrors == nil {
		return nil, false
	}
	return o.ComputeBereqErrors, true
}

// HasComputeBereqErrors returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasComputeBereqErrors() bool {
	if o != nil && o.ComputeBereqErrors != nil {
		return true
	}

	return false
}

// SetComputeBereqErrors gets a reference to the given int64 and assigns it to the ComputeBereqErrors field.
func (o *RealtimeEntryAggregated) SetComputeBereqErrors(v int64) {
	o.ComputeBereqErrors = &v
}

// GetComputeResourceLimitExceeded returns the ComputeResourceLimitExceeded field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetComputeResourceLimitExceeded() int64 {
	if o == nil || o.ComputeResourceLimitExceeded == nil {
		var ret int64
		return ret
	}
	return *o.ComputeResourceLimitExceeded
}

// GetComputeResourceLimitExceededOk returns a tuple with the ComputeResourceLimitExceeded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetComputeResourceLimitExceededOk() (*int64, bool) {
	if o == nil || o.ComputeResourceLimitExceeded == nil {
		return nil, false
	}
	return o.ComputeResourceLimitExceeded, true
}

// HasComputeResourceLimitExceeded returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasComputeResourceLimitExceeded() bool {
	if o != nil && o.ComputeResourceLimitExceeded != nil {
		return true
	}

	return false
}

// SetComputeResourceLimitExceeded gets a reference to the given int64 and assigns it to the ComputeResourceLimitExceeded field.
func (o *RealtimeEntryAggregated) SetComputeResourceLimitExceeded(v int64) {
	o.ComputeResourceLimitExceeded = &v
}

// GetComputeHeapLimitExceeded returns the ComputeHeapLimitExceeded field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetComputeHeapLimitExceeded() int64 {
	if o == nil || o.ComputeHeapLimitExceeded == nil {
		var ret int64
		return ret
	}
	return *o.ComputeHeapLimitExceeded
}

// GetComputeHeapLimitExceededOk returns a tuple with the ComputeHeapLimitExceeded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetComputeHeapLimitExceededOk() (*int64, bool) {
	if o == nil || o.ComputeHeapLimitExceeded == nil {
		return nil, false
	}
	return o.ComputeHeapLimitExceeded, true
}

// HasComputeHeapLimitExceeded returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasComputeHeapLimitExceeded() bool {
	if o != nil && o.ComputeHeapLimitExceeded != nil {
		return true
	}

	return false
}

// SetComputeHeapLimitExceeded gets a reference to the given int64 and assigns it to the ComputeHeapLimitExceeded field.
func (o *RealtimeEntryAggregated) SetComputeHeapLimitExceeded(v int64) {
	o.ComputeHeapLimitExceeded = &v
}

// GetComputeStackLimitExceeded returns the ComputeStackLimitExceeded field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetComputeStackLimitExceeded() int64 {
	if o == nil || o.ComputeStackLimitExceeded == nil {
		var ret int64
		return ret
	}
	return *o.ComputeStackLimitExceeded
}

// GetComputeStackLimitExceededOk returns a tuple with the ComputeStackLimitExceeded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetComputeStackLimitExceededOk() (*int64, bool) {
	if o == nil || o.ComputeStackLimitExceeded == nil {
		return nil, false
	}
	return o.ComputeStackLimitExceeded, true
}

// HasComputeStackLimitExceeded returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasComputeStackLimitExceeded() bool {
	if o != nil && o.ComputeStackLimitExceeded != nil {
		return true
	}

	return false
}

// SetComputeStackLimitExceeded gets a reference to the given int64 and assigns it to the ComputeStackLimitExceeded field.
func (o *RealtimeEntryAggregated) SetComputeStackLimitExceeded(v int64) {
	o.ComputeStackLimitExceeded = &v
}

// GetComputeGlobalsLimitExceeded returns the ComputeGlobalsLimitExceeded field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetComputeGlobalsLimitExceeded() int64 {
	if o == nil || o.ComputeGlobalsLimitExceeded == nil {
		var ret int64
		return ret
	}
	return *o.ComputeGlobalsLimitExceeded
}

// GetComputeGlobalsLimitExceededOk returns a tuple with the ComputeGlobalsLimitExceeded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetComputeGlobalsLimitExceededOk() (*int64, bool) {
	if o == nil || o.ComputeGlobalsLimitExceeded == nil {
		return nil, false
	}
	return o.ComputeGlobalsLimitExceeded, true
}

// HasComputeGlobalsLimitExceeded returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasComputeGlobalsLimitExceeded() bool {
	if o != nil && o.ComputeGlobalsLimitExceeded != nil {
		return true
	}

	return false
}

// SetComputeGlobalsLimitExceeded gets a reference to the given int64 and assigns it to the ComputeGlobalsLimitExceeded field.
func (o *RealtimeEntryAggregated) SetComputeGlobalsLimitExceeded(v int64) {
	o.ComputeGlobalsLimitExceeded = &v
}

// GetComputeGuestErrors returns the ComputeGuestErrors field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetComputeGuestErrors() int64 {
	if o == nil || o.ComputeGuestErrors == nil {
		var ret int64
		return ret
	}
	return *o.ComputeGuestErrors
}

// GetComputeGuestErrorsOk returns a tuple with the ComputeGuestErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetComputeGuestErrorsOk() (*int64, bool) {
	if o == nil || o.ComputeGuestErrors == nil {
		return nil, false
	}
	return o.ComputeGuestErrors, true
}

// HasComputeGuestErrors returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasComputeGuestErrors() bool {
	if o != nil && o.ComputeGuestErrors != nil {
		return true
	}

	return false
}

// SetComputeGuestErrors gets a reference to the given int64 and assigns it to the ComputeGuestErrors field.
func (o *RealtimeEntryAggregated) SetComputeGuestErrors(v int64) {
	o.ComputeGuestErrors = &v
}

// GetComputeRuntimeErrors returns the ComputeRuntimeErrors field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetComputeRuntimeErrors() int64 {
	if o == nil || o.ComputeRuntimeErrors == nil {
		var ret int64
		return ret
	}
	return *o.ComputeRuntimeErrors
}

// GetComputeRuntimeErrorsOk returns a tuple with the ComputeRuntimeErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetComputeRuntimeErrorsOk() (*int64, bool) {
	if o == nil || o.ComputeRuntimeErrors == nil {
		return nil, false
	}
	return o.ComputeRuntimeErrors, true
}

// HasComputeRuntimeErrors returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasComputeRuntimeErrors() bool {
	if o != nil && o.ComputeRuntimeErrors != nil {
		return true
	}

	return false
}

// SetComputeRuntimeErrors gets a reference to the given int64 and assigns it to the ComputeRuntimeErrors field.
func (o *RealtimeEntryAggregated) SetComputeRuntimeErrors(v int64) {
	o.ComputeRuntimeErrors = &v
}

// GetEdgeHitRespBodyBytes returns the EdgeHitRespBodyBytes field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetEdgeHitRespBodyBytes() int64 {
	if o == nil || o.EdgeHitRespBodyBytes == nil {
		var ret int64
		return ret
	}
	return *o.EdgeHitRespBodyBytes
}

// GetEdgeHitRespBodyBytesOk returns a tuple with the EdgeHitRespBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetEdgeHitRespBodyBytesOk() (*int64, bool) {
	if o == nil || o.EdgeHitRespBodyBytes == nil {
		return nil, false
	}
	return o.EdgeHitRespBodyBytes, true
}

// HasEdgeHitRespBodyBytes returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasEdgeHitRespBodyBytes() bool {
	if o != nil && o.EdgeHitRespBodyBytes != nil {
		return true
	}

	return false
}

// SetEdgeHitRespBodyBytes gets a reference to the given int64 and assigns it to the EdgeHitRespBodyBytes field.
func (o *RealtimeEntryAggregated) SetEdgeHitRespBodyBytes(v int64) {
	o.EdgeHitRespBodyBytes = &v
}

// GetEdgeHitRespHeaderBytes returns the EdgeHitRespHeaderBytes field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetEdgeHitRespHeaderBytes() int64 {
	if o == nil || o.EdgeHitRespHeaderBytes == nil {
		var ret int64
		return ret
	}
	return *o.EdgeHitRespHeaderBytes
}

// GetEdgeHitRespHeaderBytesOk returns a tuple with the EdgeHitRespHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetEdgeHitRespHeaderBytesOk() (*int64, bool) {
	if o == nil || o.EdgeHitRespHeaderBytes == nil {
		return nil, false
	}
	return o.EdgeHitRespHeaderBytes, true
}

// HasEdgeHitRespHeaderBytes returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasEdgeHitRespHeaderBytes() bool {
	if o != nil && o.EdgeHitRespHeaderBytes != nil {
		return true
	}

	return false
}

// SetEdgeHitRespHeaderBytes gets a reference to the given int64 and assigns it to the EdgeHitRespHeaderBytes field.
func (o *RealtimeEntryAggregated) SetEdgeHitRespHeaderBytes(v int64) {
	o.EdgeHitRespHeaderBytes = &v
}

// GetEdgeMissRespBodyBytes returns the EdgeMissRespBodyBytes field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetEdgeMissRespBodyBytes() int64 {
	if o == nil || o.EdgeMissRespBodyBytes == nil {
		var ret int64
		return ret
	}
	return *o.EdgeMissRespBodyBytes
}

// GetEdgeMissRespBodyBytesOk returns a tuple with the EdgeMissRespBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetEdgeMissRespBodyBytesOk() (*int64, bool) {
	if o == nil || o.EdgeMissRespBodyBytes == nil {
		return nil, false
	}
	return o.EdgeMissRespBodyBytes, true
}

// HasEdgeMissRespBodyBytes returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasEdgeMissRespBodyBytes() bool {
	if o != nil && o.EdgeMissRespBodyBytes != nil {
		return true
	}

	return false
}

// SetEdgeMissRespBodyBytes gets a reference to the given int64 and assigns it to the EdgeMissRespBodyBytes field.
func (o *RealtimeEntryAggregated) SetEdgeMissRespBodyBytes(v int64) {
	o.EdgeMissRespBodyBytes = &v
}

// GetEdgeMissRespHeaderBytes returns the EdgeMissRespHeaderBytes field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetEdgeMissRespHeaderBytes() int64 {
	if o == nil || o.EdgeMissRespHeaderBytes == nil {
		var ret int64
		return ret
	}
	return *o.EdgeMissRespHeaderBytes
}

// GetEdgeMissRespHeaderBytesOk returns a tuple with the EdgeMissRespHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetEdgeMissRespHeaderBytesOk() (*int64, bool) {
	if o == nil || o.EdgeMissRespHeaderBytes == nil {
		return nil, false
	}
	return o.EdgeMissRespHeaderBytes, true
}

// HasEdgeMissRespHeaderBytes returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasEdgeMissRespHeaderBytes() bool {
	if o != nil && o.EdgeMissRespHeaderBytes != nil {
		return true
	}

	return false
}

// SetEdgeMissRespHeaderBytes gets a reference to the given int64 and assigns it to the EdgeMissRespHeaderBytes field.
func (o *RealtimeEntryAggregated) SetEdgeMissRespHeaderBytes(v int64) {
	o.EdgeMissRespHeaderBytes = &v
}

// GetOriginCacheFetchRespBodyBytes returns the OriginCacheFetchRespBodyBytes field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetOriginCacheFetchRespBodyBytes() int64 {
	if o == nil || o.OriginCacheFetchRespBodyBytes == nil {
		var ret int64
		return ret
	}
	return *o.OriginCacheFetchRespBodyBytes
}

// GetOriginCacheFetchRespBodyBytesOk returns a tuple with the OriginCacheFetchRespBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetOriginCacheFetchRespBodyBytesOk() (*int64, bool) {
	if o == nil || o.OriginCacheFetchRespBodyBytes == nil {
		return nil, false
	}
	return o.OriginCacheFetchRespBodyBytes, true
}

// HasOriginCacheFetchRespBodyBytes returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasOriginCacheFetchRespBodyBytes() bool {
	if o != nil && o.OriginCacheFetchRespBodyBytes != nil {
		return true
	}

	return false
}

// SetOriginCacheFetchRespBodyBytes gets a reference to the given int64 and assigns it to the OriginCacheFetchRespBodyBytes field.
func (o *RealtimeEntryAggregated) SetOriginCacheFetchRespBodyBytes(v int64) {
	o.OriginCacheFetchRespBodyBytes = &v
}

// GetOriginCacheFetchRespHeaderBytes returns the OriginCacheFetchRespHeaderBytes field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetOriginCacheFetchRespHeaderBytes() int64 {
	if o == nil || o.OriginCacheFetchRespHeaderBytes == nil {
		var ret int64
		return ret
	}
	return *o.OriginCacheFetchRespHeaderBytes
}

// GetOriginCacheFetchRespHeaderBytesOk returns a tuple with the OriginCacheFetchRespHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetOriginCacheFetchRespHeaderBytesOk() (*int64, bool) {
	if o == nil || o.OriginCacheFetchRespHeaderBytes == nil {
		return nil, false
	}
	return o.OriginCacheFetchRespHeaderBytes, true
}

// HasOriginCacheFetchRespHeaderBytes returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasOriginCacheFetchRespHeaderBytes() bool {
	if o != nil && o.OriginCacheFetchRespHeaderBytes != nil {
		return true
	}

	return false
}

// SetOriginCacheFetchRespHeaderBytes gets a reference to the given int64 and assigns it to the OriginCacheFetchRespHeaderBytes field.
func (o *RealtimeEntryAggregated) SetOriginCacheFetchRespHeaderBytes(v int64) {
	o.OriginCacheFetchRespHeaderBytes = &v
}

// GetShieldHitRequests returns the ShieldHitRequests field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetShieldHitRequests() int64 {
	if o == nil || o.ShieldHitRequests == nil {
		var ret int64
		return ret
	}
	return *o.ShieldHitRequests
}

// GetShieldHitRequestsOk returns a tuple with the ShieldHitRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetShieldHitRequestsOk() (*int64, bool) {
	if o == nil || o.ShieldHitRequests == nil {
		return nil, false
	}
	return o.ShieldHitRequests, true
}

// HasShieldHitRequests returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasShieldHitRequests() bool {
	if o != nil && o.ShieldHitRequests != nil {
		return true
	}

	return false
}

// SetShieldHitRequests gets a reference to the given int64 and assigns it to the ShieldHitRequests field.
func (o *RealtimeEntryAggregated) SetShieldHitRequests(v int64) {
	o.ShieldHitRequests = &v
}

// GetShieldMissRequests returns the ShieldMissRequests field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetShieldMissRequests() int64 {
	if o == nil || o.ShieldMissRequests == nil {
		var ret int64
		return ret
	}
	return *o.ShieldMissRequests
}

// GetShieldMissRequestsOk returns a tuple with the ShieldMissRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetShieldMissRequestsOk() (*int64, bool) {
	if o == nil || o.ShieldMissRequests == nil {
		return nil, false
	}
	return o.ShieldMissRequests, true
}

// HasShieldMissRequests returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasShieldMissRequests() bool {
	if o != nil && o.ShieldMissRequests != nil {
		return true
	}

	return false
}

// SetShieldMissRequests gets a reference to the given int64 and assigns it to the ShieldMissRequests field.
func (o *RealtimeEntryAggregated) SetShieldMissRequests(v int64) {
	o.ShieldMissRequests = &v
}

// GetShieldHitRespHeaderBytes returns the ShieldHitRespHeaderBytes field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetShieldHitRespHeaderBytes() int64 {
	if o == nil || o.ShieldHitRespHeaderBytes == nil {
		var ret int64
		return ret
	}
	return *o.ShieldHitRespHeaderBytes
}

// GetShieldHitRespHeaderBytesOk returns a tuple with the ShieldHitRespHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetShieldHitRespHeaderBytesOk() (*int64, bool) {
	if o == nil || o.ShieldHitRespHeaderBytes == nil {
		return nil, false
	}
	return o.ShieldHitRespHeaderBytes, true
}

// HasShieldHitRespHeaderBytes returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasShieldHitRespHeaderBytes() bool {
	if o != nil && o.ShieldHitRespHeaderBytes != nil {
		return true
	}

	return false
}

// SetShieldHitRespHeaderBytes gets a reference to the given int64 and assigns it to the ShieldHitRespHeaderBytes field.
func (o *RealtimeEntryAggregated) SetShieldHitRespHeaderBytes(v int64) {
	o.ShieldHitRespHeaderBytes = &v
}

// GetShieldHitRespBodyBytes returns the ShieldHitRespBodyBytes field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetShieldHitRespBodyBytes() int64 {
	if o == nil || o.ShieldHitRespBodyBytes == nil {
		var ret int64
		return ret
	}
	return *o.ShieldHitRespBodyBytes
}

// GetShieldHitRespBodyBytesOk returns a tuple with the ShieldHitRespBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetShieldHitRespBodyBytesOk() (*int64, bool) {
	if o == nil || o.ShieldHitRespBodyBytes == nil {
		return nil, false
	}
	return o.ShieldHitRespBodyBytes, true
}

// HasShieldHitRespBodyBytes returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasShieldHitRespBodyBytes() bool {
	if o != nil && o.ShieldHitRespBodyBytes != nil {
		return true
	}

	return false
}

// SetShieldHitRespBodyBytes gets a reference to the given int64 and assigns it to the ShieldHitRespBodyBytes field.
func (o *RealtimeEntryAggregated) SetShieldHitRespBodyBytes(v int64) {
	o.ShieldHitRespBodyBytes = &v
}

// GetShieldMissRespHeaderBytes returns the ShieldMissRespHeaderBytes field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetShieldMissRespHeaderBytes() int64 {
	if o == nil || o.ShieldMissRespHeaderBytes == nil {
		var ret int64
		return ret
	}
	return *o.ShieldMissRespHeaderBytes
}

// GetShieldMissRespHeaderBytesOk returns a tuple with the ShieldMissRespHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetShieldMissRespHeaderBytesOk() (*int64, bool) {
	if o == nil || o.ShieldMissRespHeaderBytes == nil {
		return nil, false
	}
	return o.ShieldMissRespHeaderBytes, true
}

// HasShieldMissRespHeaderBytes returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasShieldMissRespHeaderBytes() bool {
	if o != nil && o.ShieldMissRespHeaderBytes != nil {
		return true
	}

	return false
}

// SetShieldMissRespHeaderBytes gets a reference to the given int64 and assigns it to the ShieldMissRespHeaderBytes field.
func (o *RealtimeEntryAggregated) SetShieldMissRespHeaderBytes(v int64) {
	o.ShieldMissRespHeaderBytes = &v
}

// GetShieldMissRespBodyBytes returns the ShieldMissRespBodyBytes field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetShieldMissRespBodyBytes() int64 {
	if o == nil || o.ShieldMissRespBodyBytes == nil {
		var ret int64
		return ret
	}
	return *o.ShieldMissRespBodyBytes
}

// GetShieldMissRespBodyBytesOk returns a tuple with the ShieldMissRespBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetShieldMissRespBodyBytesOk() (*int64, bool) {
	if o == nil || o.ShieldMissRespBodyBytes == nil {
		return nil, false
	}
	return o.ShieldMissRespBodyBytes, true
}

// HasShieldMissRespBodyBytes returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasShieldMissRespBodyBytes() bool {
	if o != nil && o.ShieldMissRespBodyBytes != nil {
		return true
	}

	return false
}

// SetShieldMissRespBodyBytes gets a reference to the given int64 and assigns it to the ShieldMissRespBodyBytes field.
func (o *RealtimeEntryAggregated) SetShieldMissRespBodyBytes(v int64) {
	o.ShieldMissRespBodyBytes = &v
}

// GetWebsocketReqHeaderBytes returns the WebsocketReqHeaderBytes field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetWebsocketReqHeaderBytes() int64 {
	if o == nil || o.WebsocketReqHeaderBytes == nil {
		var ret int64
		return ret
	}
	return *o.WebsocketReqHeaderBytes
}

// GetWebsocketReqHeaderBytesOk returns a tuple with the WebsocketReqHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetWebsocketReqHeaderBytesOk() (*int64, bool) {
	if o == nil || o.WebsocketReqHeaderBytes == nil {
		return nil, false
	}
	return o.WebsocketReqHeaderBytes, true
}

// HasWebsocketReqHeaderBytes returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasWebsocketReqHeaderBytes() bool {
	if o != nil && o.WebsocketReqHeaderBytes != nil {
		return true
	}

	return false
}

// SetWebsocketReqHeaderBytes gets a reference to the given int64 and assigns it to the WebsocketReqHeaderBytes field.
func (o *RealtimeEntryAggregated) SetWebsocketReqHeaderBytes(v int64) {
	o.WebsocketReqHeaderBytes = &v
}

// GetWebsocketReqBodyBytes returns the WebsocketReqBodyBytes field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetWebsocketReqBodyBytes() int64 {
	if o == nil || o.WebsocketReqBodyBytes == nil {
		var ret int64
		return ret
	}
	return *o.WebsocketReqBodyBytes
}

// GetWebsocketReqBodyBytesOk returns a tuple with the WebsocketReqBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetWebsocketReqBodyBytesOk() (*int64, bool) {
	if o == nil || o.WebsocketReqBodyBytes == nil {
		return nil, false
	}
	return o.WebsocketReqBodyBytes, true
}

// HasWebsocketReqBodyBytes returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasWebsocketReqBodyBytes() bool {
	if o != nil && o.WebsocketReqBodyBytes != nil {
		return true
	}

	return false
}

// SetWebsocketReqBodyBytes gets a reference to the given int64 and assigns it to the WebsocketReqBodyBytes field.
func (o *RealtimeEntryAggregated) SetWebsocketReqBodyBytes(v int64) {
	o.WebsocketReqBodyBytes = &v
}

// GetWebsocketRespHeaderBytes returns the WebsocketRespHeaderBytes field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetWebsocketRespHeaderBytes() int64 {
	if o == nil || o.WebsocketRespHeaderBytes == nil {
		var ret int64
		return ret
	}
	return *o.WebsocketRespHeaderBytes
}

// GetWebsocketRespHeaderBytesOk returns a tuple with the WebsocketRespHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetWebsocketRespHeaderBytesOk() (*int64, bool) {
	if o == nil || o.WebsocketRespHeaderBytes == nil {
		return nil, false
	}
	return o.WebsocketRespHeaderBytes, true
}

// HasWebsocketRespHeaderBytes returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasWebsocketRespHeaderBytes() bool {
	if o != nil && o.WebsocketRespHeaderBytes != nil {
		return true
	}

	return false
}

// SetWebsocketRespHeaderBytes gets a reference to the given int64 and assigns it to the WebsocketRespHeaderBytes field.
func (o *RealtimeEntryAggregated) SetWebsocketRespHeaderBytes(v int64) {
	o.WebsocketRespHeaderBytes = &v
}

// GetWebsocketBereqHeaderBytes returns the WebsocketBereqHeaderBytes field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetWebsocketBereqHeaderBytes() int64 {
	if o == nil || o.WebsocketBereqHeaderBytes == nil {
		var ret int64
		return ret
	}
	return *o.WebsocketBereqHeaderBytes
}

// GetWebsocketBereqHeaderBytesOk returns a tuple with the WebsocketBereqHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetWebsocketBereqHeaderBytesOk() (*int64, bool) {
	if o == nil || o.WebsocketBereqHeaderBytes == nil {
		return nil, false
	}
	return o.WebsocketBereqHeaderBytes, true
}

// HasWebsocketBereqHeaderBytes returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasWebsocketBereqHeaderBytes() bool {
	if o != nil && o.WebsocketBereqHeaderBytes != nil {
		return true
	}

	return false
}

// SetWebsocketBereqHeaderBytes gets a reference to the given int64 and assigns it to the WebsocketBereqHeaderBytes field.
func (o *RealtimeEntryAggregated) SetWebsocketBereqHeaderBytes(v int64) {
	o.WebsocketBereqHeaderBytes = &v
}

// GetWebsocketBereqBodyBytes returns the WebsocketBereqBodyBytes field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetWebsocketBereqBodyBytes() int64 {
	if o == nil || o.WebsocketBereqBodyBytes == nil {
		var ret int64
		return ret
	}
	return *o.WebsocketBereqBodyBytes
}

// GetWebsocketBereqBodyBytesOk returns a tuple with the WebsocketBereqBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetWebsocketBereqBodyBytesOk() (*int64, bool) {
	if o == nil || o.WebsocketBereqBodyBytes == nil {
		return nil, false
	}
	return o.WebsocketBereqBodyBytes, true
}

// HasWebsocketBereqBodyBytes returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasWebsocketBereqBodyBytes() bool {
	if o != nil && o.WebsocketBereqBodyBytes != nil {
		return true
	}

	return false
}

// SetWebsocketBereqBodyBytes gets a reference to the given int64 and assigns it to the WebsocketBereqBodyBytes field.
func (o *RealtimeEntryAggregated) SetWebsocketBereqBodyBytes(v int64) {
	o.WebsocketBereqBodyBytes = &v
}

// GetWebsocketBerespHeaderBytes returns the WebsocketBerespHeaderBytes field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetWebsocketBerespHeaderBytes() int64 {
	if o == nil || o.WebsocketBerespHeaderBytes == nil {
		var ret int64
		return ret
	}
	return *o.WebsocketBerespHeaderBytes
}

// GetWebsocketBerespHeaderBytesOk returns a tuple with the WebsocketBerespHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetWebsocketBerespHeaderBytesOk() (*int64, bool) {
	if o == nil || o.WebsocketBerespHeaderBytes == nil {
		return nil, false
	}
	return o.WebsocketBerespHeaderBytes, true
}

// HasWebsocketBerespHeaderBytes returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasWebsocketBerespHeaderBytes() bool {
	if o != nil && o.WebsocketBerespHeaderBytes != nil {
		return true
	}

	return false
}

// SetWebsocketBerespHeaderBytes gets a reference to the given int64 and assigns it to the WebsocketBerespHeaderBytes field.
func (o *RealtimeEntryAggregated) SetWebsocketBerespHeaderBytes(v int64) {
	o.WebsocketBerespHeaderBytes = &v
}

// GetWebsocketBerespBodyBytes returns the WebsocketBerespBodyBytes field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetWebsocketBerespBodyBytes() int64 {
	if o == nil || o.WebsocketBerespBodyBytes == nil {
		var ret int64
		return ret
	}
	return *o.WebsocketBerespBodyBytes
}

// GetWebsocketBerespBodyBytesOk returns a tuple with the WebsocketBerespBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetWebsocketBerespBodyBytesOk() (*int64, bool) {
	if o == nil || o.WebsocketBerespBodyBytes == nil {
		return nil, false
	}
	return o.WebsocketBerespBodyBytes, true
}

// HasWebsocketBerespBodyBytes returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasWebsocketBerespBodyBytes() bool {
	if o != nil && o.WebsocketBerespBodyBytes != nil {
		return true
	}

	return false
}

// SetWebsocketBerespBodyBytes gets a reference to the given int64 and assigns it to the WebsocketBerespBodyBytes field.
func (o *RealtimeEntryAggregated) SetWebsocketBerespBodyBytes(v int64) {
	o.WebsocketBerespBodyBytes = &v
}

// GetWebsocketConnTimeMs returns the WebsocketConnTimeMs field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetWebsocketConnTimeMs() int64 {
	if o == nil || o.WebsocketConnTimeMs == nil {
		var ret int64
		return ret
	}
	return *o.WebsocketConnTimeMs
}

// GetWebsocketConnTimeMsOk returns a tuple with the WebsocketConnTimeMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetWebsocketConnTimeMsOk() (*int64, bool) {
	if o == nil || o.WebsocketConnTimeMs == nil {
		return nil, false
	}
	return o.WebsocketConnTimeMs, true
}

// HasWebsocketConnTimeMs returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasWebsocketConnTimeMs() bool {
	if o != nil && o.WebsocketConnTimeMs != nil {
		return true
	}

	return false
}

// SetWebsocketConnTimeMs gets a reference to the given int64 and assigns it to the WebsocketConnTimeMs field.
func (o *RealtimeEntryAggregated) SetWebsocketConnTimeMs(v int64) {
	o.WebsocketConnTimeMs = &v
}

// GetWebsocketRespBodyBytes returns the WebsocketRespBodyBytes field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetWebsocketRespBodyBytes() int64 {
	if o == nil || o.WebsocketRespBodyBytes == nil {
		var ret int64
		return ret
	}
	return *o.WebsocketRespBodyBytes
}

// GetWebsocketRespBodyBytesOk returns a tuple with the WebsocketRespBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetWebsocketRespBodyBytesOk() (*int64, bool) {
	if o == nil || o.WebsocketRespBodyBytes == nil {
		return nil, false
	}
	return o.WebsocketRespBodyBytes, true
}

// HasWebsocketRespBodyBytes returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasWebsocketRespBodyBytes() bool {
	if o != nil && o.WebsocketRespBodyBytes != nil {
		return true
	}

	return false
}

// SetWebsocketRespBodyBytes gets a reference to the given int64 and assigns it to the WebsocketRespBodyBytes field.
func (o *RealtimeEntryAggregated) SetWebsocketRespBodyBytes(v int64) {
	o.WebsocketRespBodyBytes = &v
}

// GetFanoutRecvPublishes returns the FanoutRecvPublishes field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetFanoutRecvPublishes() int64 {
	if o == nil || o.FanoutRecvPublishes == nil {
		var ret int64
		return ret
	}
	return *o.FanoutRecvPublishes
}

// GetFanoutRecvPublishesOk returns a tuple with the FanoutRecvPublishes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetFanoutRecvPublishesOk() (*int64, bool) {
	if o == nil || o.FanoutRecvPublishes == nil {
		return nil, false
	}
	return o.FanoutRecvPublishes, true
}

// HasFanoutRecvPublishes returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasFanoutRecvPublishes() bool {
	if o != nil && o.FanoutRecvPublishes != nil {
		return true
	}

	return false
}

// SetFanoutRecvPublishes gets a reference to the given int64 and assigns it to the FanoutRecvPublishes field.
func (o *RealtimeEntryAggregated) SetFanoutRecvPublishes(v int64) {
	o.FanoutRecvPublishes = &v
}

// GetFanoutSendPublishes returns the FanoutSendPublishes field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetFanoutSendPublishes() int64 {
	if o == nil || o.FanoutSendPublishes == nil {
		var ret int64
		return ret
	}
	return *o.FanoutSendPublishes
}

// GetFanoutSendPublishesOk returns a tuple with the FanoutSendPublishes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetFanoutSendPublishesOk() (*int64, bool) {
	if o == nil || o.FanoutSendPublishes == nil {
		return nil, false
	}
	return o.FanoutSendPublishes, true
}

// HasFanoutSendPublishes returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasFanoutSendPublishes() bool {
	if o != nil && o.FanoutSendPublishes != nil {
		return true
	}

	return false
}

// SetFanoutSendPublishes gets a reference to the given int64 and assigns it to the FanoutSendPublishes field.
func (o *RealtimeEntryAggregated) SetFanoutSendPublishes(v int64) {
	o.FanoutSendPublishes = &v
}

// GetKvStoreClassAOperations returns the KvStoreClassAOperations field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetKvStoreClassAOperations() int64 {
	if o == nil || o.KvStoreClassAOperations == nil {
		var ret int64
		return ret
	}
	return *o.KvStoreClassAOperations
}

// GetKvStoreClassAOperationsOk returns a tuple with the KvStoreClassAOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetKvStoreClassAOperationsOk() (*int64, bool) {
	if o == nil || o.KvStoreClassAOperations == nil {
		return nil, false
	}
	return o.KvStoreClassAOperations, true
}

// HasKvStoreClassAOperations returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasKvStoreClassAOperations() bool {
	if o != nil && o.KvStoreClassAOperations != nil {
		return true
	}

	return false
}

// SetKvStoreClassAOperations gets a reference to the given int64 and assigns it to the KvStoreClassAOperations field.
func (o *RealtimeEntryAggregated) SetKvStoreClassAOperations(v int64) {
	o.KvStoreClassAOperations = &v
}

// GetKvStoreClassBOperations returns the KvStoreClassBOperations field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetKvStoreClassBOperations() int64 {
	if o == nil || o.KvStoreClassBOperations == nil {
		var ret int64
		return ret
	}
	return *o.KvStoreClassBOperations
}

// GetKvStoreClassBOperationsOk returns a tuple with the KvStoreClassBOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetKvStoreClassBOperationsOk() (*int64, bool) {
	if o == nil || o.KvStoreClassBOperations == nil {
		return nil, false
	}
	return o.KvStoreClassBOperations, true
}

// HasKvStoreClassBOperations returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasKvStoreClassBOperations() bool {
	if o != nil && o.KvStoreClassBOperations != nil {
		return true
	}

	return false
}

// SetKvStoreClassBOperations gets a reference to the given int64 and assigns it to the KvStoreClassBOperations field.
func (o *RealtimeEntryAggregated) SetKvStoreClassBOperations(v int64) {
	o.KvStoreClassBOperations = &v
}

// GetObjectStoreClassAOperations returns the ObjectStoreClassAOperations field value if set, zero value otherwise.
// Deprecated
func (o *RealtimeEntryAggregated) GetObjectStoreClassAOperations() int64 {
	if o == nil || o.ObjectStoreClassAOperations == nil {
		var ret int64
		return ret
	}
	return *o.ObjectStoreClassAOperations
}

// GetObjectStoreClassAOperationsOk returns a tuple with the ObjectStoreClassAOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *RealtimeEntryAggregated) GetObjectStoreClassAOperationsOk() (*int64, bool) {
	if o == nil || o.ObjectStoreClassAOperations == nil {
		return nil, false
	}
	return o.ObjectStoreClassAOperations, true
}

// HasObjectStoreClassAOperations returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasObjectStoreClassAOperations() bool {
	if o != nil && o.ObjectStoreClassAOperations != nil {
		return true
	}

	return false
}

// SetObjectStoreClassAOperations gets a reference to the given int64 and assigns it to the ObjectStoreClassAOperations field.
// Deprecated
func (o *RealtimeEntryAggregated) SetObjectStoreClassAOperations(v int64) {
	o.ObjectStoreClassAOperations = &v
}

// GetObjectStoreClassBOperations returns the ObjectStoreClassBOperations field value if set, zero value otherwise.
// Deprecated
func (o *RealtimeEntryAggregated) GetObjectStoreClassBOperations() int64 {
	if o == nil || o.ObjectStoreClassBOperations == nil {
		var ret int64
		return ret
	}
	return *o.ObjectStoreClassBOperations
}

// GetObjectStoreClassBOperationsOk returns a tuple with the ObjectStoreClassBOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *RealtimeEntryAggregated) GetObjectStoreClassBOperationsOk() (*int64, bool) {
	if o == nil || o.ObjectStoreClassBOperations == nil {
		return nil, false
	}
	return o.ObjectStoreClassBOperations, true
}

// HasObjectStoreClassBOperations returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasObjectStoreClassBOperations() bool {
	if o != nil && o.ObjectStoreClassBOperations != nil {
		return true
	}

	return false
}

// SetObjectStoreClassBOperations gets a reference to the given int64 and assigns it to the ObjectStoreClassBOperations field.
// Deprecated
func (o *RealtimeEntryAggregated) SetObjectStoreClassBOperations(v int64) {
	o.ObjectStoreClassBOperations = &v
}

// GetFanoutReqHeaderBytes returns the FanoutReqHeaderBytes field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetFanoutReqHeaderBytes() int64 {
	if o == nil || o.FanoutReqHeaderBytes == nil {
		var ret int64
		return ret
	}
	return *o.FanoutReqHeaderBytes
}

// GetFanoutReqHeaderBytesOk returns a tuple with the FanoutReqHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetFanoutReqHeaderBytesOk() (*int64, bool) {
	if o == nil || o.FanoutReqHeaderBytes == nil {
		return nil, false
	}
	return o.FanoutReqHeaderBytes, true
}

// HasFanoutReqHeaderBytes returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasFanoutReqHeaderBytes() bool {
	if o != nil && o.FanoutReqHeaderBytes != nil {
		return true
	}

	return false
}

// SetFanoutReqHeaderBytes gets a reference to the given int64 and assigns it to the FanoutReqHeaderBytes field.
func (o *RealtimeEntryAggregated) SetFanoutReqHeaderBytes(v int64) {
	o.FanoutReqHeaderBytes = &v
}

// GetFanoutReqBodyBytes returns the FanoutReqBodyBytes field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetFanoutReqBodyBytes() int64 {
	if o == nil || o.FanoutReqBodyBytes == nil {
		var ret int64
		return ret
	}
	return *o.FanoutReqBodyBytes
}

// GetFanoutReqBodyBytesOk returns a tuple with the FanoutReqBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetFanoutReqBodyBytesOk() (*int64, bool) {
	if o == nil || o.FanoutReqBodyBytes == nil {
		return nil, false
	}
	return o.FanoutReqBodyBytes, true
}

// HasFanoutReqBodyBytes returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasFanoutReqBodyBytes() bool {
	if o != nil && o.FanoutReqBodyBytes != nil {
		return true
	}

	return false
}

// SetFanoutReqBodyBytes gets a reference to the given int64 and assigns it to the FanoutReqBodyBytes field.
func (o *RealtimeEntryAggregated) SetFanoutReqBodyBytes(v int64) {
	o.FanoutReqBodyBytes = &v
}

// GetFanoutRespHeaderBytes returns the FanoutRespHeaderBytes field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetFanoutRespHeaderBytes() int64 {
	if o == nil || o.FanoutRespHeaderBytes == nil {
		var ret int64
		return ret
	}
	return *o.FanoutRespHeaderBytes
}

// GetFanoutRespHeaderBytesOk returns a tuple with the FanoutRespHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetFanoutRespHeaderBytesOk() (*int64, bool) {
	if o == nil || o.FanoutRespHeaderBytes == nil {
		return nil, false
	}
	return o.FanoutRespHeaderBytes, true
}

// HasFanoutRespHeaderBytes returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasFanoutRespHeaderBytes() bool {
	if o != nil && o.FanoutRespHeaderBytes != nil {
		return true
	}

	return false
}

// SetFanoutRespHeaderBytes gets a reference to the given int64 and assigns it to the FanoutRespHeaderBytes field.
func (o *RealtimeEntryAggregated) SetFanoutRespHeaderBytes(v int64) {
	o.FanoutRespHeaderBytes = &v
}

// GetFanoutRespBodyBytes returns the FanoutRespBodyBytes field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetFanoutRespBodyBytes() int64 {
	if o == nil || o.FanoutRespBodyBytes == nil {
		var ret int64
		return ret
	}
	return *o.FanoutRespBodyBytes
}

// GetFanoutRespBodyBytesOk returns a tuple with the FanoutRespBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetFanoutRespBodyBytesOk() (*int64, bool) {
	if o == nil || o.FanoutRespBodyBytes == nil {
		return nil, false
	}
	return o.FanoutRespBodyBytes, true
}

// HasFanoutRespBodyBytes returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasFanoutRespBodyBytes() bool {
	if o != nil && o.FanoutRespBodyBytes != nil {
		return true
	}

	return false
}

// SetFanoutRespBodyBytes gets a reference to the given int64 and assigns it to the FanoutRespBodyBytes field.
func (o *RealtimeEntryAggregated) SetFanoutRespBodyBytes(v int64) {
	o.FanoutRespBodyBytes = &v
}

// GetFanoutBereqHeaderBytes returns the FanoutBereqHeaderBytes field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetFanoutBereqHeaderBytes() int64 {
	if o == nil || o.FanoutBereqHeaderBytes == nil {
		var ret int64
		return ret
	}
	return *o.FanoutBereqHeaderBytes
}

// GetFanoutBereqHeaderBytesOk returns a tuple with the FanoutBereqHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetFanoutBereqHeaderBytesOk() (*int64, bool) {
	if o == nil || o.FanoutBereqHeaderBytes == nil {
		return nil, false
	}
	return o.FanoutBereqHeaderBytes, true
}

// HasFanoutBereqHeaderBytes returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasFanoutBereqHeaderBytes() bool {
	if o != nil && o.FanoutBereqHeaderBytes != nil {
		return true
	}

	return false
}

// SetFanoutBereqHeaderBytes gets a reference to the given int64 and assigns it to the FanoutBereqHeaderBytes field.
func (o *RealtimeEntryAggregated) SetFanoutBereqHeaderBytes(v int64) {
	o.FanoutBereqHeaderBytes = &v
}

// GetFanoutBereqBodyBytes returns the FanoutBereqBodyBytes field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetFanoutBereqBodyBytes() int64 {
	if o == nil || o.FanoutBereqBodyBytes == nil {
		var ret int64
		return ret
	}
	return *o.FanoutBereqBodyBytes
}

// GetFanoutBereqBodyBytesOk returns a tuple with the FanoutBereqBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetFanoutBereqBodyBytesOk() (*int64, bool) {
	if o == nil || o.FanoutBereqBodyBytes == nil {
		return nil, false
	}
	return o.FanoutBereqBodyBytes, true
}

// HasFanoutBereqBodyBytes returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasFanoutBereqBodyBytes() bool {
	if o != nil && o.FanoutBereqBodyBytes != nil {
		return true
	}

	return false
}

// SetFanoutBereqBodyBytes gets a reference to the given int64 and assigns it to the FanoutBereqBodyBytes field.
func (o *RealtimeEntryAggregated) SetFanoutBereqBodyBytes(v int64) {
	o.FanoutBereqBodyBytes = &v
}

// GetFanoutBerespHeaderBytes returns the FanoutBerespHeaderBytes field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetFanoutBerespHeaderBytes() int64 {
	if o == nil || o.FanoutBerespHeaderBytes == nil {
		var ret int64
		return ret
	}
	return *o.FanoutBerespHeaderBytes
}

// GetFanoutBerespHeaderBytesOk returns a tuple with the FanoutBerespHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetFanoutBerespHeaderBytesOk() (*int64, bool) {
	if o == nil || o.FanoutBerespHeaderBytes == nil {
		return nil, false
	}
	return o.FanoutBerespHeaderBytes, true
}

// HasFanoutBerespHeaderBytes returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasFanoutBerespHeaderBytes() bool {
	if o != nil && o.FanoutBerespHeaderBytes != nil {
		return true
	}

	return false
}

// SetFanoutBerespHeaderBytes gets a reference to the given int64 and assigns it to the FanoutBerespHeaderBytes field.
func (o *RealtimeEntryAggregated) SetFanoutBerespHeaderBytes(v int64) {
	o.FanoutBerespHeaderBytes = &v
}

// GetFanoutBerespBodyBytes returns the FanoutBerespBodyBytes field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetFanoutBerespBodyBytes() int64 {
	if o == nil || o.FanoutBerespBodyBytes == nil {
		var ret int64
		return ret
	}
	return *o.FanoutBerespBodyBytes
}

// GetFanoutBerespBodyBytesOk returns a tuple with the FanoutBerespBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetFanoutBerespBodyBytesOk() (*int64, bool) {
	if o == nil || o.FanoutBerespBodyBytes == nil {
		return nil, false
	}
	return o.FanoutBerespBodyBytes, true
}

// HasFanoutBerespBodyBytes returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasFanoutBerespBodyBytes() bool {
	if o != nil && o.FanoutBerespBodyBytes != nil {
		return true
	}

	return false
}

// SetFanoutBerespBodyBytes gets a reference to the given int64 and assigns it to the FanoutBerespBodyBytes field.
func (o *RealtimeEntryAggregated) SetFanoutBerespBodyBytes(v int64) {
	o.FanoutBerespBodyBytes = &v
}

// GetFanoutConnTimeMs returns the FanoutConnTimeMs field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetFanoutConnTimeMs() int64 {
	if o == nil || o.FanoutConnTimeMs == nil {
		var ret int64
		return ret
	}
	return *o.FanoutConnTimeMs
}

// GetFanoutConnTimeMsOk returns a tuple with the FanoutConnTimeMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetFanoutConnTimeMsOk() (*int64, bool) {
	if o == nil || o.FanoutConnTimeMs == nil {
		return nil, false
	}
	return o.FanoutConnTimeMs, true
}

// HasFanoutConnTimeMs returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasFanoutConnTimeMs() bool {
	if o != nil && o.FanoutConnTimeMs != nil {
		return true
	}

	return false
}

// SetFanoutConnTimeMs gets a reference to the given int64 and assigns it to the FanoutConnTimeMs field.
func (o *RealtimeEntryAggregated) SetFanoutConnTimeMs(v int64) {
	o.FanoutConnTimeMs = &v
}

// GetDdosActionLimitStreamsConnections returns the DdosActionLimitStreamsConnections field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetDdosActionLimitStreamsConnections() int64 {
	if o == nil || o.DdosActionLimitStreamsConnections == nil {
		var ret int64
		return ret
	}
	return *o.DdosActionLimitStreamsConnections
}

// GetDdosActionLimitStreamsConnectionsOk returns a tuple with the DdosActionLimitStreamsConnections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetDdosActionLimitStreamsConnectionsOk() (*int64, bool) {
	if o == nil || o.DdosActionLimitStreamsConnections == nil {
		return nil, false
	}
	return o.DdosActionLimitStreamsConnections, true
}

// HasDdosActionLimitStreamsConnections returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasDdosActionLimitStreamsConnections() bool {
	if o != nil && o.DdosActionLimitStreamsConnections != nil {
		return true
	}

	return false
}

// SetDdosActionLimitStreamsConnections gets a reference to the given int64 and assigns it to the DdosActionLimitStreamsConnections field.
func (o *RealtimeEntryAggregated) SetDdosActionLimitStreamsConnections(v int64) {
	o.DdosActionLimitStreamsConnections = &v
}

// GetDdosActionLimitStreamsRequests returns the DdosActionLimitStreamsRequests field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetDdosActionLimitStreamsRequests() int64 {
	if o == nil || o.DdosActionLimitStreamsRequests == nil {
		var ret int64
		return ret
	}
	return *o.DdosActionLimitStreamsRequests
}

// GetDdosActionLimitStreamsRequestsOk returns a tuple with the DdosActionLimitStreamsRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetDdosActionLimitStreamsRequestsOk() (*int64, bool) {
	if o == nil || o.DdosActionLimitStreamsRequests == nil {
		return nil, false
	}
	return o.DdosActionLimitStreamsRequests, true
}

// HasDdosActionLimitStreamsRequests returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasDdosActionLimitStreamsRequests() bool {
	if o != nil && o.DdosActionLimitStreamsRequests != nil {
		return true
	}

	return false
}

// SetDdosActionLimitStreamsRequests gets a reference to the given int64 and assigns it to the DdosActionLimitStreamsRequests field.
func (o *RealtimeEntryAggregated) SetDdosActionLimitStreamsRequests(v int64) {
	o.DdosActionLimitStreamsRequests = &v
}

// GetDdosActionTarpitAccept returns the DdosActionTarpitAccept field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetDdosActionTarpitAccept() int64 {
	if o == nil || o.DdosActionTarpitAccept == nil {
		var ret int64
		return ret
	}
	return *o.DdosActionTarpitAccept
}

// GetDdosActionTarpitAcceptOk returns a tuple with the DdosActionTarpitAccept field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetDdosActionTarpitAcceptOk() (*int64, bool) {
	if o == nil || o.DdosActionTarpitAccept == nil {
		return nil, false
	}
	return o.DdosActionTarpitAccept, true
}

// HasDdosActionTarpitAccept returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasDdosActionTarpitAccept() bool {
	if o != nil && o.DdosActionTarpitAccept != nil {
		return true
	}

	return false
}

// SetDdosActionTarpitAccept gets a reference to the given int64 and assigns it to the DdosActionTarpitAccept field.
func (o *RealtimeEntryAggregated) SetDdosActionTarpitAccept(v int64) {
	o.DdosActionTarpitAccept = &v
}

// GetDdosActionTarpit returns the DdosActionTarpit field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetDdosActionTarpit() int64 {
	if o == nil || o.DdosActionTarpit == nil {
		var ret int64
		return ret
	}
	return *o.DdosActionTarpit
}

// GetDdosActionTarpitOk returns a tuple with the DdosActionTarpit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetDdosActionTarpitOk() (*int64, bool) {
	if o == nil || o.DdosActionTarpit == nil {
		return nil, false
	}
	return o.DdosActionTarpit, true
}

// HasDdosActionTarpit returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasDdosActionTarpit() bool {
	if o != nil && o.DdosActionTarpit != nil {
		return true
	}

	return false
}

// SetDdosActionTarpit gets a reference to the given int64 and assigns it to the DdosActionTarpit field.
func (o *RealtimeEntryAggregated) SetDdosActionTarpit(v int64) {
	o.DdosActionTarpit = &v
}

// GetDdosActionClose returns the DdosActionClose field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetDdosActionClose() int64 {
	if o == nil || o.DdosActionClose == nil {
		var ret int64
		return ret
	}
	return *o.DdosActionClose
}

// GetDdosActionCloseOk returns a tuple with the DdosActionClose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetDdosActionCloseOk() (*int64, bool) {
	if o == nil || o.DdosActionClose == nil {
		return nil, false
	}
	return o.DdosActionClose, true
}

// HasDdosActionClose returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasDdosActionClose() bool {
	if o != nil && o.DdosActionClose != nil {
		return true
	}

	return false
}

// SetDdosActionClose gets a reference to the given int64 and assigns it to the DdosActionClose field.
func (o *RealtimeEntryAggregated) SetDdosActionClose(v int64) {
	o.DdosActionClose = &v
}

// GetDdosActionBlackhole returns the DdosActionBlackhole field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetDdosActionBlackhole() int64 {
	if o == nil || o.DdosActionBlackhole == nil {
		var ret int64
		return ret
	}
	return *o.DdosActionBlackhole
}

// GetDdosActionBlackholeOk returns a tuple with the DdosActionBlackhole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetDdosActionBlackholeOk() (*int64, bool) {
	if o == nil || o.DdosActionBlackhole == nil {
		return nil, false
	}
	return o.DdosActionBlackhole, true
}

// HasDdosActionBlackhole returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasDdosActionBlackhole() bool {
	if o != nil && o.DdosActionBlackhole != nil {
		return true
	}

	return false
}

// SetDdosActionBlackhole gets a reference to the given int64 and assigns it to the DdosActionBlackhole field.
func (o *RealtimeEntryAggregated) SetDdosActionBlackhole(v int64) {
	o.DdosActionBlackhole = &v
}

// GetBotChallengeStarts returns the BotChallengeStarts field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetBotChallengeStarts() int64 {
	if o == nil || o.BotChallengeStarts == nil {
		var ret int64
		return ret
	}
	return *o.BotChallengeStarts
}

// GetBotChallengeStartsOk returns a tuple with the BotChallengeStarts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetBotChallengeStartsOk() (*int64, bool) {
	if o == nil || o.BotChallengeStarts == nil {
		return nil, false
	}
	return o.BotChallengeStarts, true
}

// HasBotChallengeStarts returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasBotChallengeStarts() bool {
	if o != nil && o.BotChallengeStarts != nil {
		return true
	}

	return false
}

// SetBotChallengeStarts gets a reference to the given int64 and assigns it to the BotChallengeStarts field.
func (o *RealtimeEntryAggregated) SetBotChallengeStarts(v int64) {
	o.BotChallengeStarts = &v
}

// GetBotChallengeCompleteTokensPassed returns the BotChallengeCompleteTokensPassed field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetBotChallengeCompleteTokensPassed() int64 {
	if o == nil || o.BotChallengeCompleteTokensPassed == nil {
		var ret int64
		return ret
	}
	return *o.BotChallengeCompleteTokensPassed
}

// GetBotChallengeCompleteTokensPassedOk returns a tuple with the BotChallengeCompleteTokensPassed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetBotChallengeCompleteTokensPassedOk() (*int64, bool) {
	if o == nil || o.BotChallengeCompleteTokensPassed == nil {
		return nil, false
	}
	return o.BotChallengeCompleteTokensPassed, true
}

// HasBotChallengeCompleteTokensPassed returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasBotChallengeCompleteTokensPassed() bool {
	if o != nil && o.BotChallengeCompleteTokensPassed != nil {
		return true
	}

	return false
}

// SetBotChallengeCompleteTokensPassed gets a reference to the given int64 and assigns it to the BotChallengeCompleteTokensPassed field.
func (o *RealtimeEntryAggregated) SetBotChallengeCompleteTokensPassed(v int64) {
	o.BotChallengeCompleteTokensPassed = &v
}

// GetBotChallengeCompleteTokensFailed returns the BotChallengeCompleteTokensFailed field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetBotChallengeCompleteTokensFailed() int64 {
	if o == nil || o.BotChallengeCompleteTokensFailed == nil {
		var ret int64
		return ret
	}
	return *o.BotChallengeCompleteTokensFailed
}

// GetBotChallengeCompleteTokensFailedOk returns a tuple with the BotChallengeCompleteTokensFailed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetBotChallengeCompleteTokensFailedOk() (*int64, bool) {
	if o == nil || o.BotChallengeCompleteTokensFailed == nil {
		return nil, false
	}
	return o.BotChallengeCompleteTokensFailed, true
}

// HasBotChallengeCompleteTokensFailed returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasBotChallengeCompleteTokensFailed() bool {
	if o != nil && o.BotChallengeCompleteTokensFailed != nil {
		return true
	}

	return false
}

// SetBotChallengeCompleteTokensFailed gets a reference to the given int64 and assigns it to the BotChallengeCompleteTokensFailed field.
func (o *RealtimeEntryAggregated) SetBotChallengeCompleteTokensFailed(v int64) {
	o.BotChallengeCompleteTokensFailed = &v
}

// GetBotChallengeCompleteTokensChecked returns the BotChallengeCompleteTokensChecked field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetBotChallengeCompleteTokensChecked() int64 {
	if o == nil || o.BotChallengeCompleteTokensChecked == nil {
		var ret int64
		return ret
	}
	return *o.BotChallengeCompleteTokensChecked
}

// GetBotChallengeCompleteTokensCheckedOk returns a tuple with the BotChallengeCompleteTokensChecked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetBotChallengeCompleteTokensCheckedOk() (*int64, bool) {
	if o == nil || o.BotChallengeCompleteTokensChecked == nil {
		return nil, false
	}
	return o.BotChallengeCompleteTokensChecked, true
}

// HasBotChallengeCompleteTokensChecked returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasBotChallengeCompleteTokensChecked() bool {
	if o != nil && o.BotChallengeCompleteTokensChecked != nil {
		return true
	}

	return false
}

// SetBotChallengeCompleteTokensChecked gets a reference to the given int64 and assigns it to the BotChallengeCompleteTokensChecked field.
func (o *RealtimeEntryAggregated) SetBotChallengeCompleteTokensChecked(v int64) {
	o.BotChallengeCompleteTokensChecked = &v
}

// GetBotChallengeCompleteTokensDisabled returns the BotChallengeCompleteTokensDisabled field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetBotChallengeCompleteTokensDisabled() int64 {
	if o == nil || o.BotChallengeCompleteTokensDisabled == nil {
		var ret int64
		return ret
	}
	return *o.BotChallengeCompleteTokensDisabled
}

// GetBotChallengeCompleteTokensDisabledOk returns a tuple with the BotChallengeCompleteTokensDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetBotChallengeCompleteTokensDisabledOk() (*int64, bool) {
	if o == nil || o.BotChallengeCompleteTokensDisabled == nil {
		return nil, false
	}
	return o.BotChallengeCompleteTokensDisabled, true
}

// HasBotChallengeCompleteTokensDisabled returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasBotChallengeCompleteTokensDisabled() bool {
	if o != nil && o.BotChallengeCompleteTokensDisabled != nil {
		return true
	}

	return false
}

// SetBotChallengeCompleteTokensDisabled gets a reference to the given int64 and assigns it to the BotChallengeCompleteTokensDisabled field.
func (o *RealtimeEntryAggregated) SetBotChallengeCompleteTokensDisabled(v int64) {
	o.BotChallengeCompleteTokensDisabled = &v
}

// GetBotChallengesIssued returns the BotChallengesIssued field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetBotChallengesIssued() int64 {
	if o == nil || o.BotChallengesIssued == nil {
		var ret int64
		return ret
	}
	return *o.BotChallengesIssued
}

// GetBotChallengesIssuedOk returns a tuple with the BotChallengesIssued field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetBotChallengesIssuedOk() (*int64, bool) {
	if o == nil || o.BotChallengesIssued == nil {
		return nil, false
	}
	return o.BotChallengesIssued, true
}

// HasBotChallengesIssued returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasBotChallengesIssued() bool {
	if o != nil && o.BotChallengesIssued != nil {
		return true
	}

	return false
}

// SetBotChallengesIssued gets a reference to the given int64 and assigns it to the BotChallengesIssued field.
func (o *RealtimeEntryAggregated) SetBotChallengesIssued(v int64) {
	o.BotChallengesIssued = &v
}

// GetBotChallengesSucceeded returns the BotChallengesSucceeded field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetBotChallengesSucceeded() int64 {
	if o == nil || o.BotChallengesSucceeded == nil {
		var ret int64
		return ret
	}
	return *o.BotChallengesSucceeded
}

// GetBotChallengesSucceededOk returns a tuple with the BotChallengesSucceeded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetBotChallengesSucceededOk() (*int64, bool) {
	if o == nil || o.BotChallengesSucceeded == nil {
		return nil, false
	}
	return o.BotChallengesSucceeded, true
}

// HasBotChallengesSucceeded returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasBotChallengesSucceeded() bool {
	if o != nil && o.BotChallengesSucceeded != nil {
		return true
	}

	return false
}

// SetBotChallengesSucceeded gets a reference to the given int64 and assigns it to the BotChallengesSucceeded field.
func (o *RealtimeEntryAggregated) SetBotChallengesSucceeded(v int64) {
	o.BotChallengesSucceeded = &v
}

// GetBotChallengesFailed returns the BotChallengesFailed field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetBotChallengesFailed() int64 {
	if o == nil || o.BotChallengesFailed == nil {
		var ret int64
		return ret
	}
	return *o.BotChallengesFailed
}

// GetBotChallengesFailedOk returns a tuple with the BotChallengesFailed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetBotChallengesFailedOk() (*int64, bool) {
	if o == nil || o.BotChallengesFailed == nil {
		return nil, false
	}
	return o.BotChallengesFailed, true
}

// HasBotChallengesFailed returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasBotChallengesFailed() bool {
	if o != nil && o.BotChallengesFailed != nil {
		return true
	}

	return false
}

// SetBotChallengesFailed gets a reference to the given int64 and assigns it to the BotChallengesFailed field.
func (o *RealtimeEntryAggregated) SetBotChallengesFailed(v int64) {
	o.BotChallengesFailed = &v
}

// GetBotChallengeCompleteTokensIssued returns the BotChallengeCompleteTokensIssued field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetBotChallengeCompleteTokensIssued() int64 {
	if o == nil || o.BotChallengeCompleteTokensIssued == nil {
		var ret int64
		return ret
	}
	return *o.BotChallengeCompleteTokensIssued
}

// GetBotChallengeCompleteTokensIssuedOk returns a tuple with the BotChallengeCompleteTokensIssued field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetBotChallengeCompleteTokensIssuedOk() (*int64, bool) {
	if o == nil || o.BotChallengeCompleteTokensIssued == nil {
		return nil, false
	}
	return o.BotChallengeCompleteTokensIssued, true
}

// HasBotChallengeCompleteTokensIssued returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasBotChallengeCompleteTokensIssued() bool {
	if o != nil && o.BotChallengeCompleteTokensIssued != nil {
		return true
	}

	return false
}

// SetBotChallengeCompleteTokensIssued gets a reference to the given int64 and assigns it to the BotChallengeCompleteTokensIssued field.
func (o *RealtimeEntryAggregated) SetBotChallengeCompleteTokensIssued(v int64) {
	o.BotChallengeCompleteTokensIssued = &v
}

// GetDdosActionDowngrade returns the DdosActionDowngrade field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetDdosActionDowngrade() int64 {
	if o == nil || o.DdosActionDowngrade == nil {
		var ret int64
		return ret
	}
	return *o.DdosActionDowngrade
}

// GetDdosActionDowngradeOk returns a tuple with the DdosActionDowngrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetDdosActionDowngradeOk() (*int64, bool) {
	if o == nil || o.DdosActionDowngrade == nil {
		return nil, false
	}
	return o.DdosActionDowngrade, true
}

// HasDdosActionDowngrade returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasDdosActionDowngrade() bool {
	if o != nil && o.DdosActionDowngrade != nil {
		return true
	}

	return false
}

// SetDdosActionDowngrade gets a reference to the given int64 and assigns it to the DdosActionDowngrade field.
func (o *RealtimeEntryAggregated) SetDdosActionDowngrade(v int64) {
	o.DdosActionDowngrade = &v
}

// GetDdosActionDowngradedConnections returns the DdosActionDowngradedConnections field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetDdosActionDowngradedConnections() int64 {
	if o == nil || o.DdosActionDowngradedConnections == nil {
		var ret int64
		return ret
	}
	return *o.DdosActionDowngradedConnections
}

// GetDdosActionDowngradedConnectionsOk returns a tuple with the DdosActionDowngradedConnections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetDdosActionDowngradedConnectionsOk() (*int64, bool) {
	if o == nil || o.DdosActionDowngradedConnections == nil {
		return nil, false
	}
	return o.DdosActionDowngradedConnections, true
}

// HasDdosActionDowngradedConnections returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasDdosActionDowngradedConnections() bool {
	if o != nil && o.DdosActionDowngradedConnections != nil {
		return true
	}

	return false
}

// SetDdosActionDowngradedConnections gets a reference to the given int64 and assigns it to the DdosActionDowngradedConnections field.
func (o *RealtimeEntryAggregated) SetDdosActionDowngradedConnections(v int64) {
	o.DdosActionDowngradedConnections = &v
}

// GetAllHitRequests returns the AllHitRequests field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetAllHitRequests() int64 {
	if o == nil || o.AllHitRequests == nil {
		var ret int64
		return ret
	}
	return *o.AllHitRequests
}

// GetAllHitRequestsOk returns a tuple with the AllHitRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetAllHitRequestsOk() (*int64, bool) {
	if o == nil || o.AllHitRequests == nil {
		return nil, false
	}
	return o.AllHitRequests, true
}

// HasAllHitRequests returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasAllHitRequests() bool {
	if o != nil && o.AllHitRequests != nil {
		return true
	}

	return false
}

// SetAllHitRequests gets a reference to the given int64 and assigns it to the AllHitRequests field.
func (o *RealtimeEntryAggregated) SetAllHitRequests(v int64) {
	o.AllHitRequests = &v
}

// GetAllMissRequests returns the AllMissRequests field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetAllMissRequests() int64 {
	if o == nil || o.AllMissRequests == nil {
		var ret int64
		return ret
	}
	return *o.AllMissRequests
}

// GetAllMissRequestsOk returns a tuple with the AllMissRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetAllMissRequestsOk() (*int64, bool) {
	if o == nil || o.AllMissRequests == nil {
		return nil, false
	}
	return o.AllMissRequests, true
}

// HasAllMissRequests returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasAllMissRequests() bool {
	if o != nil && o.AllMissRequests != nil {
		return true
	}

	return false
}

// SetAllMissRequests gets a reference to the given int64 and assigns it to the AllMissRequests field.
func (o *RealtimeEntryAggregated) SetAllMissRequests(v int64) {
	o.AllMissRequests = &v
}

// GetAllPassRequests returns the AllPassRequests field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetAllPassRequests() int64 {
	if o == nil || o.AllPassRequests == nil {
		var ret int64
		return ret
	}
	return *o.AllPassRequests
}

// GetAllPassRequestsOk returns a tuple with the AllPassRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetAllPassRequestsOk() (*int64, bool) {
	if o == nil || o.AllPassRequests == nil {
		return nil, false
	}
	return o.AllPassRequests, true
}

// HasAllPassRequests returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasAllPassRequests() bool {
	if o != nil && o.AllPassRequests != nil {
		return true
	}

	return false
}

// SetAllPassRequests gets a reference to the given int64 and assigns it to the AllPassRequests field.
func (o *RealtimeEntryAggregated) SetAllPassRequests(v int64) {
	o.AllPassRequests = &v
}

// GetAllErrorRequests returns the AllErrorRequests field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetAllErrorRequests() int64 {
	if o == nil || o.AllErrorRequests == nil {
		var ret int64
		return ret
	}
	return *o.AllErrorRequests
}

// GetAllErrorRequestsOk returns a tuple with the AllErrorRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetAllErrorRequestsOk() (*int64, bool) {
	if o == nil || o.AllErrorRequests == nil {
		return nil, false
	}
	return o.AllErrorRequests, true
}

// HasAllErrorRequests returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasAllErrorRequests() bool {
	if o != nil && o.AllErrorRequests != nil {
		return true
	}

	return false
}

// SetAllErrorRequests gets a reference to the given int64 and assigns it to the AllErrorRequests field.
func (o *RealtimeEntryAggregated) SetAllErrorRequests(v int64) {
	o.AllErrorRequests = &v
}

// GetAllSynthRequests returns the AllSynthRequests field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetAllSynthRequests() int64 {
	if o == nil || o.AllSynthRequests == nil {
		var ret int64
		return ret
	}
	return *o.AllSynthRequests
}

// GetAllSynthRequestsOk returns a tuple with the AllSynthRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetAllSynthRequestsOk() (*int64, bool) {
	if o == nil || o.AllSynthRequests == nil {
		return nil, false
	}
	return o.AllSynthRequests, true
}

// HasAllSynthRequests returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasAllSynthRequests() bool {
	if o != nil && o.AllSynthRequests != nil {
		return true
	}

	return false
}

// SetAllSynthRequests gets a reference to the given int64 and assigns it to the AllSynthRequests field.
func (o *RealtimeEntryAggregated) SetAllSynthRequests(v int64) {
	o.AllSynthRequests = &v
}

// GetAllEdgeHitRequests returns the AllEdgeHitRequests field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetAllEdgeHitRequests() int64 {
	if o == nil || o.AllEdgeHitRequests == nil {
		var ret int64
		return ret
	}
	return *o.AllEdgeHitRequests
}

// GetAllEdgeHitRequestsOk returns a tuple with the AllEdgeHitRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetAllEdgeHitRequestsOk() (*int64, bool) {
	if o == nil || o.AllEdgeHitRequests == nil {
		return nil, false
	}
	return o.AllEdgeHitRequests, true
}

// HasAllEdgeHitRequests returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasAllEdgeHitRequests() bool {
	if o != nil && o.AllEdgeHitRequests != nil {
		return true
	}

	return false
}

// SetAllEdgeHitRequests gets a reference to the given int64 and assigns it to the AllEdgeHitRequests field.
func (o *RealtimeEntryAggregated) SetAllEdgeHitRequests(v int64) {
	o.AllEdgeHitRequests = &v
}

// GetAllEdgeMissRequests returns the AllEdgeMissRequests field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetAllEdgeMissRequests() int64 {
	if o == nil || o.AllEdgeMissRequests == nil {
		var ret int64
		return ret
	}
	return *o.AllEdgeMissRequests
}

// GetAllEdgeMissRequestsOk returns a tuple with the AllEdgeMissRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetAllEdgeMissRequestsOk() (*int64, bool) {
	if o == nil || o.AllEdgeMissRequests == nil {
		return nil, false
	}
	return o.AllEdgeMissRequests, true
}

// HasAllEdgeMissRequests returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasAllEdgeMissRequests() bool {
	if o != nil && o.AllEdgeMissRequests != nil {
		return true
	}

	return false
}

// SetAllEdgeMissRequests gets a reference to the given int64 and assigns it to the AllEdgeMissRequests field.
func (o *RealtimeEntryAggregated) SetAllEdgeMissRequests(v int64) {
	o.AllEdgeMissRequests = &v
}

// GetAllStatus1xx returns the AllStatus1xx field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetAllStatus1xx() int64 {
	if o == nil || o.AllStatus1xx == nil {
		var ret int64
		return ret
	}
	return *o.AllStatus1xx
}

// GetAllStatus1xxOk returns a tuple with the AllStatus1xx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetAllStatus1xxOk() (*int64, bool) {
	if o == nil || o.AllStatus1xx == nil {
		return nil, false
	}
	return o.AllStatus1xx, true
}

// HasAllStatus1xx returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasAllStatus1xx() bool {
	if o != nil && o.AllStatus1xx != nil {
		return true
	}

	return false
}

// SetAllStatus1xx gets a reference to the given int64 and assigns it to the AllStatus1xx field.
func (o *RealtimeEntryAggregated) SetAllStatus1xx(v int64) {
	o.AllStatus1xx = &v
}

// GetAllStatus2xx returns the AllStatus2xx field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetAllStatus2xx() int64 {
	if o == nil || o.AllStatus2xx == nil {
		var ret int64
		return ret
	}
	return *o.AllStatus2xx
}

// GetAllStatus2xxOk returns a tuple with the AllStatus2xx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetAllStatus2xxOk() (*int64, bool) {
	if o == nil || o.AllStatus2xx == nil {
		return nil, false
	}
	return o.AllStatus2xx, true
}

// HasAllStatus2xx returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasAllStatus2xx() bool {
	if o != nil && o.AllStatus2xx != nil {
		return true
	}

	return false
}

// SetAllStatus2xx gets a reference to the given int64 and assigns it to the AllStatus2xx field.
func (o *RealtimeEntryAggregated) SetAllStatus2xx(v int64) {
	o.AllStatus2xx = &v
}

// GetAllStatus3xx returns the AllStatus3xx field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetAllStatus3xx() int64 {
	if o == nil || o.AllStatus3xx == nil {
		var ret int64
		return ret
	}
	return *o.AllStatus3xx
}

// GetAllStatus3xxOk returns a tuple with the AllStatus3xx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetAllStatus3xxOk() (*int64, bool) {
	if o == nil || o.AllStatus3xx == nil {
		return nil, false
	}
	return o.AllStatus3xx, true
}

// HasAllStatus3xx returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasAllStatus3xx() bool {
	if o != nil && o.AllStatus3xx != nil {
		return true
	}

	return false
}

// SetAllStatus3xx gets a reference to the given int64 and assigns it to the AllStatus3xx field.
func (o *RealtimeEntryAggregated) SetAllStatus3xx(v int64) {
	o.AllStatus3xx = &v
}

// GetAllStatus4xx returns the AllStatus4xx field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetAllStatus4xx() int64 {
	if o == nil || o.AllStatus4xx == nil {
		var ret int64
		return ret
	}
	return *o.AllStatus4xx
}

// GetAllStatus4xxOk returns a tuple with the AllStatus4xx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetAllStatus4xxOk() (*int64, bool) {
	if o == nil || o.AllStatus4xx == nil {
		return nil, false
	}
	return o.AllStatus4xx, true
}

// HasAllStatus4xx returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasAllStatus4xx() bool {
	if o != nil && o.AllStatus4xx != nil {
		return true
	}

	return false
}

// SetAllStatus4xx gets a reference to the given int64 and assigns it to the AllStatus4xx field.
func (o *RealtimeEntryAggregated) SetAllStatus4xx(v int64) {
	o.AllStatus4xx = &v
}

// GetAllStatus5xx returns the AllStatus5xx field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetAllStatus5xx() int64 {
	if o == nil || o.AllStatus5xx == nil {
		var ret int64
		return ret
	}
	return *o.AllStatus5xx
}

// GetAllStatus5xxOk returns a tuple with the AllStatus5xx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetAllStatus5xxOk() (*int64, bool) {
	if o == nil || o.AllStatus5xx == nil {
		return nil, false
	}
	return o.AllStatus5xx, true
}

// HasAllStatus5xx returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasAllStatus5xx() bool {
	if o != nil && o.AllStatus5xx != nil {
		return true
	}

	return false
}

// SetAllStatus5xx gets a reference to the given int64 and assigns it to the AllStatus5xx field.
func (o *RealtimeEntryAggregated) SetAllStatus5xx(v int64) {
	o.AllStatus5xx = &v
}

// GetOriginOffload returns the OriginOffload field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetOriginOffload() float32 {
	if o == nil || o.OriginOffload == nil {
		var ret float32
		return ret
	}
	return *o.OriginOffload
}

// GetOriginOffloadOk returns a tuple with the OriginOffload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetOriginOffloadOk() (*float32, bool) {
	if o == nil || o.OriginOffload == nil {
		return nil, false
	}
	return o.OriginOffload, true
}

// HasOriginOffload returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasOriginOffload() bool {
	if o != nil && o.OriginOffload != nil {
		return true
	}

	return false
}

// SetOriginOffload gets a reference to the given float32 and assigns it to the OriginOffload field.
func (o *RealtimeEntryAggregated) SetOriginOffload(v float32) {
	o.OriginOffload = &v
}

// GetRequestDeniedGetHeadBody returns the RequestDeniedGetHeadBody field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetRequestDeniedGetHeadBody() int64 {
	if o == nil || o.RequestDeniedGetHeadBody == nil {
		var ret int64
		return ret
	}
	return *o.RequestDeniedGetHeadBody
}

// GetRequestDeniedGetHeadBodyOk returns a tuple with the RequestDeniedGetHeadBody field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetRequestDeniedGetHeadBodyOk() (*int64, bool) {
	if o == nil || o.RequestDeniedGetHeadBody == nil {
		return nil, false
	}
	return o.RequestDeniedGetHeadBody, true
}

// HasRequestDeniedGetHeadBody returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasRequestDeniedGetHeadBody() bool {
	if o != nil && o.RequestDeniedGetHeadBody != nil {
		return true
	}

	return false
}

// SetRequestDeniedGetHeadBody gets a reference to the given int64 and assigns it to the RequestDeniedGetHeadBody field.
func (o *RealtimeEntryAggregated) SetRequestDeniedGetHeadBody(v int64) {
	o.RequestDeniedGetHeadBody = &v
}

// GetServiceDdosRequestsDetected returns the ServiceDdosRequestsDetected field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetServiceDdosRequestsDetected() int64 {
	if o == nil || o.ServiceDdosRequestsDetected == nil {
		var ret int64
		return ret
	}
	return *o.ServiceDdosRequestsDetected
}

// GetServiceDdosRequestsDetectedOk returns a tuple with the ServiceDdosRequestsDetected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetServiceDdosRequestsDetectedOk() (*int64, bool) {
	if o == nil || o.ServiceDdosRequestsDetected == nil {
		return nil, false
	}
	return o.ServiceDdosRequestsDetected, true
}

// HasServiceDdosRequestsDetected returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasServiceDdosRequestsDetected() bool {
	if o != nil && o.ServiceDdosRequestsDetected != nil {
		return true
	}

	return false
}

// SetServiceDdosRequestsDetected gets a reference to the given int64 and assigns it to the ServiceDdosRequestsDetected field.
func (o *RealtimeEntryAggregated) SetServiceDdosRequestsDetected(v int64) {
	o.ServiceDdosRequestsDetected = &v
}

// GetServiceDdosRequestsMitigated returns the ServiceDdosRequestsMitigated field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetServiceDdosRequestsMitigated() int64 {
	if o == nil || o.ServiceDdosRequestsMitigated == nil {
		var ret int64
		return ret
	}
	return *o.ServiceDdosRequestsMitigated
}

// GetServiceDdosRequestsMitigatedOk returns a tuple with the ServiceDdosRequestsMitigated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetServiceDdosRequestsMitigatedOk() (*int64, bool) {
	if o == nil || o.ServiceDdosRequestsMitigated == nil {
		return nil, false
	}
	return o.ServiceDdosRequestsMitigated, true
}

// HasServiceDdosRequestsMitigated returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasServiceDdosRequestsMitigated() bool {
	if o != nil && o.ServiceDdosRequestsMitigated != nil {
		return true
	}

	return false
}

// SetServiceDdosRequestsMitigated gets a reference to the given int64 and assigns it to the ServiceDdosRequestsMitigated field.
func (o *RealtimeEntryAggregated) SetServiceDdosRequestsMitigated(v int64) {
	o.ServiceDdosRequestsMitigated = &v
}

// GetServiceDdosRequestsAllowed returns the ServiceDdosRequestsAllowed field value if set, zero value otherwise.
func (o *RealtimeEntryAggregated) GetServiceDdosRequestsAllowed() int64 {
	if o == nil || o.ServiceDdosRequestsAllowed == nil {
		var ret int64
		return ret
	}
	return *o.ServiceDdosRequestsAllowed
}

// GetServiceDdosRequestsAllowedOk returns a tuple with the ServiceDdosRequestsAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealtimeEntryAggregated) GetServiceDdosRequestsAllowedOk() (*int64, bool) {
	if o == nil || o.ServiceDdosRequestsAllowed == nil {
		return nil, false
	}
	return o.ServiceDdosRequestsAllowed, true
}

// HasServiceDdosRequestsAllowed returns a boolean if a field has been set.
func (o *RealtimeEntryAggregated) HasServiceDdosRequestsAllowed() bool {
	if o != nil && o.ServiceDdosRequestsAllowed != nil {
		return true
	}

	return false
}

// SetServiceDdosRequestsAllowed gets a reference to the given int64 and assigns it to the ServiceDdosRequestsAllowed field.
func (o *RealtimeEntryAggregated) SetServiceDdosRequestsAllowed(v int64) {
	o.ServiceDdosRequestsAllowed = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RealtimeEntryAggregated) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Requests != nil {
		toSerialize["requests"] = o.Requests
	}
	if o.Logging != nil {
		toSerialize["logging"] = o.Logging
	}
	if o.Log != nil {
		toSerialize["log"] = o.Log
	}
	if o.RespHeaderBytes != nil {
		toSerialize["resp_header_bytes"] = o.RespHeaderBytes
	}
	if o.HeaderSize != nil {
		toSerialize["header_size"] = o.HeaderSize
	}
	if o.RespBodyBytes != nil {
		toSerialize["resp_body_bytes"] = o.RespBodyBytes
	}
	if o.BodySize != nil {
		toSerialize["body_size"] = o.BodySize
	}
	if o.Hits != nil {
		toSerialize["hits"] = o.Hits
	}
	if o.Miss != nil {
		toSerialize["miss"] = o.Miss
	}
	if o.Pass != nil {
		toSerialize["pass"] = o.Pass
	}
	if o.Synth != nil {
		toSerialize["synth"] = o.Synth
	}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}
	if o.HitsTime != nil {
		toSerialize["hits_time"] = o.HitsTime
	}
	if o.MissTime != nil {
		toSerialize["miss_time"] = o.MissTime
	}
	if o.MissHistogram != nil {
		toSerialize["miss_histogram"] = o.MissHistogram
	}
	if o.ComputeRequests != nil {
		toSerialize["compute_requests"] = o.ComputeRequests
	}
	if o.ComputeExecutionTimeMs != nil {
		toSerialize["compute_execution_time_ms"] = o.ComputeExecutionTimeMs
	}
	if o.ComputeRAMUsed != nil {
		toSerialize["compute_ram_used"] = o.ComputeRAMUsed
	}
	if o.ComputeRequestTimeMs != nil {
		toSerialize["compute_request_time_ms"] = o.ComputeRequestTimeMs
	}
	if o.ComputeRequestTimeBilledMs != nil {
		toSerialize["compute_request_time_billed_ms"] = o.ComputeRequestTimeBilledMs
	}
	if o.Shield != nil {
		toSerialize["shield"] = o.Shield
	}
	if o.Ipv6 != nil {
		toSerialize["ipv6"] = o.Ipv6
	}
	if o.Imgopto != nil {
		toSerialize["imgopto"] = o.Imgopto
	}
	if o.ImgoptoShield != nil {
		toSerialize["imgopto_shield"] = o.ImgoptoShield
	}
	if o.ImgoptoTransforms != nil {
		toSerialize["imgopto_transforms"] = o.ImgoptoTransforms
	}
	if o.Otfp != nil {
		toSerialize["otfp"] = o.Otfp
	}
	if o.OtfpShield != nil {
		toSerialize["otfp_shield"] = o.OtfpShield
	}
	if o.OtfpManifests != nil {
		toSerialize["otfp_manifests"] = o.OtfpManifests
	}
	if o.Video != nil {
		toSerialize["video"] = o.Video
	}
	if o.Pci != nil {
		toSerialize["pci"] = o.Pci
	}
	if o.HTTP2 != nil {
		toSerialize["http2"] = o.HTTP2
	}
	if o.HTTP3 != nil {
		toSerialize["http3"] = o.HTTP3
	}
	if o.Restarts != nil {
		toSerialize["restarts"] = o.Restarts
	}
	if o.ReqHeaderBytes != nil {
		toSerialize["req_header_bytes"] = o.ReqHeaderBytes
	}
	if o.ReqBodyBytes != nil {
		toSerialize["req_body_bytes"] = o.ReqBodyBytes
	}
	if o.BereqHeaderBytes != nil {
		toSerialize["bereq_header_bytes"] = o.BereqHeaderBytes
	}
	if o.BereqBodyBytes != nil {
		toSerialize["bereq_body_bytes"] = o.BereqBodyBytes
	}
	if o.WafBlocked != nil {
		toSerialize["waf_blocked"] = o.WafBlocked
	}
	if o.WafLogged != nil {
		toSerialize["waf_logged"] = o.WafLogged
	}
	if o.WafPassed != nil {
		toSerialize["waf_passed"] = o.WafPassed
	}
	if o.AttackReqHeaderBytes != nil {
		toSerialize["attack_req_header_bytes"] = o.AttackReqHeaderBytes
	}
	if o.AttackReqBodyBytes != nil {
		toSerialize["attack_req_body_bytes"] = o.AttackReqBodyBytes
	}
	if o.AttackRespSynthBytes != nil {
		toSerialize["attack_resp_synth_bytes"] = o.AttackRespSynthBytes
	}
	if o.AttackLoggedReqHeaderBytes != nil {
		toSerialize["attack_logged_req_header_bytes"] = o.AttackLoggedReqHeaderBytes
	}
	if o.AttackLoggedReqBodyBytes != nil {
		toSerialize["attack_logged_req_body_bytes"] = o.AttackLoggedReqBodyBytes
	}
	if o.AttackBlockedReqHeaderBytes != nil {
		toSerialize["attack_blocked_req_header_bytes"] = o.AttackBlockedReqHeaderBytes
	}
	if o.AttackBlockedReqBodyBytes != nil {
		toSerialize["attack_blocked_req_body_bytes"] = o.AttackBlockedReqBodyBytes
	}
	if o.AttackPassedReqHeaderBytes != nil {
		toSerialize["attack_passed_req_header_bytes"] = o.AttackPassedReqHeaderBytes
	}
	if o.AttackPassedReqBodyBytes != nil {
		toSerialize["attack_passed_req_body_bytes"] = o.AttackPassedReqBodyBytes
	}
	if o.ShieldRespHeaderBytes != nil {
		toSerialize["shield_resp_header_bytes"] = o.ShieldRespHeaderBytes
	}
	if o.ShieldRespBodyBytes != nil {
		toSerialize["shield_resp_body_bytes"] = o.ShieldRespBodyBytes
	}
	if o.OtfpRespHeaderBytes != nil {
		toSerialize["otfp_resp_header_bytes"] = o.OtfpRespHeaderBytes
	}
	if o.OtfpRespBodyBytes != nil {
		toSerialize["otfp_resp_body_bytes"] = o.OtfpRespBodyBytes
	}
	if o.OtfpShieldRespHeaderBytes != nil {
		toSerialize["otfp_shield_resp_header_bytes"] = o.OtfpShieldRespHeaderBytes
	}
	if o.OtfpShieldRespBodyBytes != nil {
		toSerialize["otfp_shield_resp_body_bytes"] = o.OtfpShieldRespBodyBytes
	}
	if o.OtfpShieldTime != nil {
		toSerialize["otfp_shield_time"] = o.OtfpShieldTime
	}
	if o.OtfpDeliverTime != nil {
		toSerialize["otfp_deliver_time"] = o.OtfpDeliverTime
	}
	if o.ImgoptoRespHeaderBytes != nil {
		toSerialize["imgopto_resp_header_bytes"] = o.ImgoptoRespHeaderBytes
	}
	if o.ImgoptoRespBodyBytes != nil {
		toSerialize["imgopto_resp_body_bytes"] = o.ImgoptoRespBodyBytes
	}
	if o.ImgoptoShieldRespHeaderBytes != nil {
		toSerialize["imgopto_shield_resp_header_bytes"] = o.ImgoptoShieldRespHeaderBytes
	}
	if o.ImgoptoShieldRespBodyBytes != nil {
		toSerialize["imgopto_shield_resp_body_bytes"] = o.ImgoptoShieldRespBodyBytes
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
	if o.Status406 != nil {
		toSerialize["status_406"] = o.Status406
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
	if o.Uncacheable != nil {
		toSerialize["uncacheable"] = o.Uncacheable
	}
	if o.PassTime != nil {
		toSerialize["pass_time"] = o.PassTime
	}
	if o.TLS != nil {
		toSerialize["tls"] = o.TLS
	}
	if o.TLSV10 != nil {
		toSerialize["tls_v10"] = o.TLSV10
	}
	if o.TLSV11 != nil {
		toSerialize["tls_v11"] = o.TLSV11
	}
	if o.TLSV12 != nil {
		toSerialize["tls_v12"] = o.TLSV12
	}
	if o.TLSV13 != nil {
		toSerialize["tls_v13"] = o.TLSV13
	}
	if o.ObjectSize1k != nil {
		toSerialize["object_size_1k"] = o.ObjectSize1k
	}
	if o.ObjectSize10k != nil {
		toSerialize["object_size_10k"] = o.ObjectSize10k
	}
	if o.ObjectSize100k != nil {
		toSerialize["object_size_100k"] = o.ObjectSize100k
	}
	if o.ObjectSize1m != nil {
		toSerialize["object_size_1m"] = o.ObjectSize1m
	}
	if o.ObjectSize10m != nil {
		toSerialize["object_size_10m"] = o.ObjectSize10m
	}
	if o.ObjectSize100m != nil {
		toSerialize["object_size_100m"] = o.ObjectSize100m
	}
	if o.ObjectSize1g != nil {
		toSerialize["object_size_1g"] = o.ObjectSize1g
	}
	if o.ObjectSizeOther != nil {
		toSerialize["object_size_other"] = o.ObjectSizeOther
	}
	if o.RecvSubTime != nil {
		toSerialize["recv_sub_time"] = o.RecvSubTime
	}
	if o.RecvSubCount != nil {
		toSerialize["recv_sub_count"] = o.RecvSubCount
	}
	if o.HashSubTime != nil {
		toSerialize["hash_sub_time"] = o.HashSubTime
	}
	if o.HashSubCount != nil {
		toSerialize["hash_sub_count"] = o.HashSubCount
	}
	if o.MissSubTime != nil {
		toSerialize["miss_sub_time"] = o.MissSubTime
	}
	if o.MissSubCount != nil {
		toSerialize["miss_sub_count"] = o.MissSubCount
	}
	if o.FetchSubTime != nil {
		toSerialize["fetch_sub_time"] = o.FetchSubTime
	}
	if o.FetchSubCount != nil {
		toSerialize["fetch_sub_count"] = o.FetchSubCount
	}
	if o.PassSubTime != nil {
		toSerialize["pass_sub_time"] = o.PassSubTime
	}
	if o.PassSubCount != nil {
		toSerialize["pass_sub_count"] = o.PassSubCount
	}
	if o.PipeSubTime != nil {
		toSerialize["pipe_sub_time"] = o.PipeSubTime
	}
	if o.PipeSubCount != nil {
		toSerialize["pipe_sub_count"] = o.PipeSubCount
	}
	if o.DeliverSubTime != nil {
		toSerialize["deliver_sub_time"] = o.DeliverSubTime
	}
	if o.DeliverSubCount != nil {
		toSerialize["deliver_sub_count"] = o.DeliverSubCount
	}
	if o.ErrorSubTime != nil {
		toSerialize["error_sub_time"] = o.ErrorSubTime
	}
	if o.ErrorSubCount != nil {
		toSerialize["error_sub_count"] = o.ErrorSubCount
	}
	if o.HitSubTime != nil {
		toSerialize["hit_sub_time"] = o.HitSubTime
	}
	if o.HitSubCount != nil {
		toSerialize["hit_sub_count"] = o.HitSubCount
	}
	if o.PrehashSubTime != nil {
		toSerialize["prehash_sub_time"] = o.PrehashSubTime
	}
	if o.PrehashSubCount != nil {
		toSerialize["prehash_sub_count"] = o.PrehashSubCount
	}
	if o.PredeliverSubTime != nil {
		toSerialize["predeliver_sub_time"] = o.PredeliverSubTime
	}
	if o.PredeliverSubCount != nil {
		toSerialize["predeliver_sub_count"] = o.PredeliverSubCount
	}
	if o.HitRespBodyBytes != nil {
		toSerialize["hit_resp_body_bytes"] = o.HitRespBodyBytes
	}
	if o.MissRespBodyBytes != nil {
		toSerialize["miss_resp_body_bytes"] = o.MissRespBodyBytes
	}
	if o.PassRespBodyBytes != nil {
		toSerialize["pass_resp_body_bytes"] = o.PassRespBodyBytes
	}
	if o.ComputeReqHeaderBytes != nil {
		toSerialize["compute_req_header_bytes"] = o.ComputeReqHeaderBytes
	}
	if o.ComputeReqBodyBytes != nil {
		toSerialize["compute_req_body_bytes"] = o.ComputeReqBodyBytes
	}
	if o.ComputeRespHeaderBytes != nil {
		toSerialize["compute_resp_header_bytes"] = o.ComputeRespHeaderBytes
	}
	if o.ComputeRespBodyBytes != nil {
		toSerialize["compute_resp_body_bytes"] = o.ComputeRespBodyBytes
	}
	if o.Imgvideo != nil {
		toSerialize["imgvideo"] = o.Imgvideo
	}
	if o.ImgvideoFrames != nil {
		toSerialize["imgvideo_frames"] = o.ImgvideoFrames
	}
	if o.ImgvideoRespHeaderBytes != nil {
		toSerialize["imgvideo_resp_header_bytes"] = o.ImgvideoRespHeaderBytes
	}
	if o.ImgvideoRespBodyBytes != nil {
		toSerialize["imgvideo_resp_body_bytes"] = o.ImgvideoRespBodyBytes
	}
	if o.ImgvideoShield != nil {
		toSerialize["imgvideo_shield"] = o.ImgvideoShield
	}
	if o.ImgvideoShieldFrames != nil {
		toSerialize["imgvideo_shield_frames"] = o.ImgvideoShieldFrames
	}
	if o.ImgvideoShieldRespHeaderBytes != nil {
		toSerialize["imgvideo_shield_resp_header_bytes"] = o.ImgvideoShieldRespHeaderBytes
	}
	if o.ImgvideoShieldRespBodyBytes != nil {
		toSerialize["imgvideo_shield_resp_body_bytes"] = o.ImgvideoShieldRespBodyBytes
	}
	if o.LogBytes != nil {
		toSerialize["log_bytes"] = o.LogBytes
	}
	if o.EdgeRequests != nil {
		toSerialize["edge_requests"] = o.EdgeRequests
	}
	if o.EdgeRespHeaderBytes != nil {
		toSerialize["edge_resp_header_bytes"] = o.EdgeRespHeaderBytes
	}
	if o.EdgeRespBodyBytes != nil {
		toSerialize["edge_resp_body_bytes"] = o.EdgeRespBodyBytes
	}
	if o.OriginRevalidations != nil {
		toSerialize["origin_revalidations"] = o.OriginRevalidations
	}
	if o.OriginFetches != nil {
		toSerialize["origin_fetches"] = o.OriginFetches
	}
	if o.OriginFetchHeaderBytes != nil {
		toSerialize["origin_fetch_header_bytes"] = o.OriginFetchHeaderBytes
	}
	if o.OriginFetchBodyBytes != nil {
		toSerialize["origin_fetch_body_bytes"] = o.OriginFetchBodyBytes
	}
	if o.OriginFetchRespHeaderBytes != nil {
		toSerialize["origin_fetch_resp_header_bytes"] = o.OriginFetchRespHeaderBytes
	}
	if o.OriginFetchRespBodyBytes != nil {
		toSerialize["origin_fetch_resp_body_bytes"] = o.OriginFetchRespBodyBytes
	}
	if o.ShieldRevalidations != nil {
		toSerialize["shield_revalidations"] = o.ShieldRevalidations
	}
	if o.ShieldFetches != nil {
		toSerialize["shield_fetches"] = o.ShieldFetches
	}
	if o.ShieldFetchHeaderBytes != nil {
		toSerialize["shield_fetch_header_bytes"] = o.ShieldFetchHeaderBytes
	}
	if o.ShieldFetchBodyBytes != nil {
		toSerialize["shield_fetch_body_bytes"] = o.ShieldFetchBodyBytes
	}
	if o.ShieldFetchRespHeaderBytes != nil {
		toSerialize["shield_fetch_resp_header_bytes"] = o.ShieldFetchRespHeaderBytes
	}
	if o.ShieldFetchRespBodyBytes != nil {
		toSerialize["shield_fetch_resp_body_bytes"] = o.ShieldFetchRespBodyBytes
	}
	if o.SegblockOriginFetches != nil {
		toSerialize["segblock_origin_fetches"] = o.SegblockOriginFetches
	}
	if o.SegblockShieldFetches != nil {
		toSerialize["segblock_shield_fetches"] = o.SegblockShieldFetches
	}
	if o.ComputeRespStatus1xx != nil {
		toSerialize["compute_resp_status_1xx"] = o.ComputeRespStatus1xx
	}
	if o.ComputeRespStatus2xx != nil {
		toSerialize["compute_resp_status_2xx"] = o.ComputeRespStatus2xx
	}
	if o.ComputeRespStatus3xx != nil {
		toSerialize["compute_resp_status_3xx"] = o.ComputeRespStatus3xx
	}
	if o.ComputeRespStatus4xx != nil {
		toSerialize["compute_resp_status_4xx"] = o.ComputeRespStatus4xx
	}
	if o.ComputeRespStatus5xx != nil {
		toSerialize["compute_resp_status_5xx"] = o.ComputeRespStatus5xx
	}
	if o.EdgeHitRequests != nil {
		toSerialize["edge_hit_requests"] = o.EdgeHitRequests
	}
	if o.EdgeMissRequests != nil {
		toSerialize["edge_miss_requests"] = o.EdgeMissRequests
	}
	if o.ComputeBereqHeaderBytes != nil {
		toSerialize["compute_bereq_header_bytes"] = o.ComputeBereqHeaderBytes
	}
	if o.ComputeBereqBodyBytes != nil {
		toSerialize["compute_bereq_body_bytes"] = o.ComputeBereqBodyBytes
	}
	if o.ComputeBerespHeaderBytes != nil {
		toSerialize["compute_beresp_header_bytes"] = o.ComputeBerespHeaderBytes
	}
	if o.ComputeBerespBodyBytes != nil {
		toSerialize["compute_beresp_body_bytes"] = o.ComputeBerespBodyBytes
	}
	if o.OriginCacheFetches != nil {
		toSerialize["origin_cache_fetches"] = o.OriginCacheFetches
	}
	if o.ShieldCacheFetches != nil {
		toSerialize["shield_cache_fetches"] = o.ShieldCacheFetches
	}
	if o.ComputeBereqs != nil {
		toSerialize["compute_bereqs"] = o.ComputeBereqs
	}
	if o.ComputeBereqErrors != nil {
		toSerialize["compute_bereq_errors"] = o.ComputeBereqErrors
	}
	if o.ComputeResourceLimitExceeded != nil {
		toSerialize["compute_resource_limit_exceeded"] = o.ComputeResourceLimitExceeded
	}
	if o.ComputeHeapLimitExceeded != nil {
		toSerialize["compute_heap_limit_exceeded"] = o.ComputeHeapLimitExceeded
	}
	if o.ComputeStackLimitExceeded != nil {
		toSerialize["compute_stack_limit_exceeded"] = o.ComputeStackLimitExceeded
	}
	if o.ComputeGlobalsLimitExceeded != nil {
		toSerialize["compute_globals_limit_exceeded"] = o.ComputeGlobalsLimitExceeded
	}
	if o.ComputeGuestErrors != nil {
		toSerialize["compute_guest_errors"] = o.ComputeGuestErrors
	}
	if o.ComputeRuntimeErrors != nil {
		toSerialize["compute_runtime_errors"] = o.ComputeRuntimeErrors
	}
	if o.EdgeHitRespBodyBytes != nil {
		toSerialize["edge_hit_resp_body_bytes"] = o.EdgeHitRespBodyBytes
	}
	if o.EdgeHitRespHeaderBytes != nil {
		toSerialize["edge_hit_resp_header_bytes"] = o.EdgeHitRespHeaderBytes
	}
	if o.EdgeMissRespBodyBytes != nil {
		toSerialize["edge_miss_resp_body_bytes"] = o.EdgeMissRespBodyBytes
	}
	if o.EdgeMissRespHeaderBytes != nil {
		toSerialize["edge_miss_resp_header_bytes"] = o.EdgeMissRespHeaderBytes
	}
	if o.OriginCacheFetchRespBodyBytes != nil {
		toSerialize["origin_cache_fetch_resp_body_bytes"] = o.OriginCacheFetchRespBodyBytes
	}
	if o.OriginCacheFetchRespHeaderBytes != nil {
		toSerialize["origin_cache_fetch_resp_header_bytes"] = o.OriginCacheFetchRespHeaderBytes
	}
	if o.ShieldHitRequests != nil {
		toSerialize["shield_hit_requests"] = o.ShieldHitRequests
	}
	if o.ShieldMissRequests != nil {
		toSerialize["shield_miss_requests"] = o.ShieldMissRequests
	}
	if o.ShieldHitRespHeaderBytes != nil {
		toSerialize["shield_hit_resp_header_bytes"] = o.ShieldHitRespHeaderBytes
	}
	if o.ShieldHitRespBodyBytes != nil {
		toSerialize["shield_hit_resp_body_bytes"] = o.ShieldHitRespBodyBytes
	}
	if o.ShieldMissRespHeaderBytes != nil {
		toSerialize["shield_miss_resp_header_bytes"] = o.ShieldMissRespHeaderBytes
	}
	if o.ShieldMissRespBodyBytes != nil {
		toSerialize["shield_miss_resp_body_bytes"] = o.ShieldMissRespBodyBytes
	}
	if o.WebsocketReqHeaderBytes != nil {
		toSerialize["websocket_req_header_bytes"] = o.WebsocketReqHeaderBytes
	}
	if o.WebsocketReqBodyBytes != nil {
		toSerialize["websocket_req_body_bytes"] = o.WebsocketReqBodyBytes
	}
	if o.WebsocketRespHeaderBytes != nil {
		toSerialize["websocket_resp_header_bytes"] = o.WebsocketRespHeaderBytes
	}
	if o.WebsocketBereqHeaderBytes != nil {
		toSerialize["websocket_bereq_header_bytes"] = o.WebsocketBereqHeaderBytes
	}
	if o.WebsocketBereqBodyBytes != nil {
		toSerialize["websocket_bereq_body_bytes"] = o.WebsocketBereqBodyBytes
	}
	if o.WebsocketBerespHeaderBytes != nil {
		toSerialize["websocket_beresp_header_bytes"] = o.WebsocketBerespHeaderBytes
	}
	if o.WebsocketBerespBodyBytes != nil {
		toSerialize["websocket_beresp_body_bytes"] = o.WebsocketBerespBodyBytes
	}
	if o.WebsocketConnTimeMs != nil {
		toSerialize["websocket_conn_time_ms"] = o.WebsocketConnTimeMs
	}
	if o.WebsocketRespBodyBytes != nil {
		toSerialize["websocket_resp_body_bytes"] = o.WebsocketRespBodyBytes
	}
	if o.FanoutRecvPublishes != nil {
		toSerialize["fanout_recv_publishes"] = o.FanoutRecvPublishes
	}
	if o.FanoutSendPublishes != nil {
		toSerialize["fanout_send_publishes"] = o.FanoutSendPublishes
	}
	if o.KvStoreClassAOperations != nil {
		toSerialize["kv_store_class_a_operations"] = o.KvStoreClassAOperations
	}
	if o.KvStoreClassBOperations != nil {
		toSerialize["kv_store_class_b_operations"] = o.KvStoreClassBOperations
	}
	if o.ObjectStoreClassAOperations != nil {
		toSerialize["object_store_class_a_operations"] = o.ObjectStoreClassAOperations
	}
	if o.ObjectStoreClassBOperations != nil {
		toSerialize["object_store_class_b_operations"] = o.ObjectStoreClassBOperations
	}
	if o.FanoutReqHeaderBytes != nil {
		toSerialize["fanout_req_header_bytes"] = o.FanoutReqHeaderBytes
	}
	if o.FanoutReqBodyBytes != nil {
		toSerialize["fanout_req_body_bytes"] = o.FanoutReqBodyBytes
	}
	if o.FanoutRespHeaderBytes != nil {
		toSerialize["fanout_resp_header_bytes"] = o.FanoutRespHeaderBytes
	}
	if o.FanoutRespBodyBytes != nil {
		toSerialize["fanout_resp_body_bytes"] = o.FanoutRespBodyBytes
	}
	if o.FanoutBereqHeaderBytes != nil {
		toSerialize["fanout_bereq_header_bytes"] = o.FanoutBereqHeaderBytes
	}
	if o.FanoutBereqBodyBytes != nil {
		toSerialize["fanout_bereq_body_bytes"] = o.FanoutBereqBodyBytes
	}
	if o.FanoutBerespHeaderBytes != nil {
		toSerialize["fanout_beresp_header_bytes"] = o.FanoutBerespHeaderBytes
	}
	if o.FanoutBerespBodyBytes != nil {
		toSerialize["fanout_beresp_body_bytes"] = o.FanoutBerespBodyBytes
	}
	if o.FanoutConnTimeMs != nil {
		toSerialize["fanout_conn_time_ms"] = o.FanoutConnTimeMs
	}
	if o.DdosActionLimitStreamsConnections != nil {
		toSerialize["ddos_action_limit_streams_connections"] = o.DdosActionLimitStreamsConnections
	}
	if o.DdosActionLimitStreamsRequests != nil {
		toSerialize["ddos_action_limit_streams_requests"] = o.DdosActionLimitStreamsRequests
	}
	if o.DdosActionTarpitAccept != nil {
		toSerialize["ddos_action_tarpit_accept"] = o.DdosActionTarpitAccept
	}
	if o.DdosActionTarpit != nil {
		toSerialize["ddos_action_tarpit"] = o.DdosActionTarpit
	}
	if o.DdosActionClose != nil {
		toSerialize["ddos_action_close"] = o.DdosActionClose
	}
	if o.DdosActionBlackhole != nil {
		toSerialize["ddos_action_blackhole"] = o.DdosActionBlackhole
	}
	if o.BotChallengeStarts != nil {
		toSerialize["bot_challenge_starts"] = o.BotChallengeStarts
	}
	if o.BotChallengeCompleteTokensPassed != nil {
		toSerialize["bot_challenge_complete_tokens_passed"] = o.BotChallengeCompleteTokensPassed
	}
	if o.BotChallengeCompleteTokensFailed != nil {
		toSerialize["bot_challenge_complete_tokens_failed"] = o.BotChallengeCompleteTokensFailed
	}
	if o.BotChallengeCompleteTokensChecked != nil {
		toSerialize["bot_challenge_complete_tokens_checked"] = o.BotChallengeCompleteTokensChecked
	}
	if o.BotChallengeCompleteTokensDisabled != nil {
		toSerialize["bot_challenge_complete_tokens_disabled"] = o.BotChallengeCompleteTokensDisabled
	}
	if o.BotChallengesIssued != nil {
		toSerialize["bot_challenges_issued"] = o.BotChallengesIssued
	}
	if o.BotChallengesSucceeded != nil {
		toSerialize["bot_challenges_succeeded"] = o.BotChallengesSucceeded
	}
	if o.BotChallengesFailed != nil {
		toSerialize["bot_challenges_failed"] = o.BotChallengesFailed
	}
	if o.BotChallengeCompleteTokensIssued != nil {
		toSerialize["bot_challenge_complete_tokens_issued"] = o.BotChallengeCompleteTokensIssued
	}
	if o.DdosActionDowngrade != nil {
		toSerialize["ddos_action_downgrade"] = o.DdosActionDowngrade
	}
	if o.DdosActionDowngradedConnections != nil {
		toSerialize["ddos_action_downgraded_connections"] = o.DdosActionDowngradedConnections
	}
	if o.AllHitRequests != nil {
		toSerialize["all_hit_requests"] = o.AllHitRequests
	}
	if o.AllMissRequests != nil {
		toSerialize["all_miss_requests"] = o.AllMissRequests
	}
	if o.AllPassRequests != nil {
		toSerialize["all_pass_requests"] = o.AllPassRequests
	}
	if o.AllErrorRequests != nil {
		toSerialize["all_error_requests"] = o.AllErrorRequests
	}
	if o.AllSynthRequests != nil {
		toSerialize["all_synth_requests"] = o.AllSynthRequests
	}
	if o.AllEdgeHitRequests != nil {
		toSerialize["all_edge_hit_requests"] = o.AllEdgeHitRequests
	}
	if o.AllEdgeMissRequests != nil {
		toSerialize["all_edge_miss_requests"] = o.AllEdgeMissRequests
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
	if o.OriginOffload != nil {
		toSerialize["origin_offload"] = o.OriginOffload
	}
	if o.RequestDeniedGetHeadBody != nil {
		toSerialize["request_denied_get_head_body"] = o.RequestDeniedGetHeadBody
	}
	if o.ServiceDdosRequestsDetected != nil {
		toSerialize["service_ddos_requests_detected"] = o.ServiceDdosRequestsDetected
	}
	if o.ServiceDdosRequestsMitigated != nil {
		toSerialize["service_ddos_requests_mitigated"] = o.ServiceDdosRequestsMitigated
	}
	if o.ServiceDdosRequestsAllowed != nil {
		toSerialize["service_ddos_requests_allowed"] = o.ServiceDdosRequestsAllowed
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *RealtimeEntryAggregated) UnmarshalJSON(bytes []byte) (err error) {
	varRealtimeEntryAggregated := _RealtimeEntryAggregated{}

	if err = json.Unmarshal(bytes, &varRealtimeEntryAggregated); err == nil {
		*o = RealtimeEntryAggregated(varRealtimeEntryAggregated)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "requests")
		delete(additionalProperties, "logging")
		delete(additionalProperties, "log")
		delete(additionalProperties, "resp_header_bytes")
		delete(additionalProperties, "header_size")
		delete(additionalProperties, "resp_body_bytes")
		delete(additionalProperties, "body_size")
		delete(additionalProperties, "hits")
		delete(additionalProperties, "miss")
		delete(additionalProperties, "pass")
		delete(additionalProperties, "synth")
		delete(additionalProperties, "errors")
		delete(additionalProperties, "hits_time")
		delete(additionalProperties, "miss_time")
		delete(additionalProperties, "miss_histogram")
		delete(additionalProperties, "compute_requests")
		delete(additionalProperties, "compute_execution_time_ms")
		delete(additionalProperties, "compute_ram_used")
		delete(additionalProperties, "compute_request_time_ms")
		delete(additionalProperties, "compute_request_time_billed_ms")
		delete(additionalProperties, "shield")
		delete(additionalProperties, "ipv6")
		delete(additionalProperties, "imgopto")
		delete(additionalProperties, "imgopto_shield")
		delete(additionalProperties, "imgopto_transforms")
		delete(additionalProperties, "otfp")
		delete(additionalProperties, "otfp_shield")
		delete(additionalProperties, "otfp_manifests")
		delete(additionalProperties, "video")
		delete(additionalProperties, "pci")
		delete(additionalProperties, "http2")
		delete(additionalProperties, "http3")
		delete(additionalProperties, "restarts")
		delete(additionalProperties, "req_header_bytes")
		delete(additionalProperties, "req_body_bytes")
		delete(additionalProperties, "bereq_header_bytes")
		delete(additionalProperties, "bereq_body_bytes")
		delete(additionalProperties, "waf_blocked")
		delete(additionalProperties, "waf_logged")
		delete(additionalProperties, "waf_passed")
		delete(additionalProperties, "attack_req_header_bytes")
		delete(additionalProperties, "attack_req_body_bytes")
		delete(additionalProperties, "attack_resp_synth_bytes")
		delete(additionalProperties, "attack_logged_req_header_bytes")
		delete(additionalProperties, "attack_logged_req_body_bytes")
		delete(additionalProperties, "attack_blocked_req_header_bytes")
		delete(additionalProperties, "attack_blocked_req_body_bytes")
		delete(additionalProperties, "attack_passed_req_header_bytes")
		delete(additionalProperties, "attack_passed_req_body_bytes")
		delete(additionalProperties, "shield_resp_header_bytes")
		delete(additionalProperties, "shield_resp_body_bytes")
		delete(additionalProperties, "otfp_resp_header_bytes")
		delete(additionalProperties, "otfp_resp_body_bytes")
		delete(additionalProperties, "otfp_shield_resp_header_bytes")
		delete(additionalProperties, "otfp_shield_resp_body_bytes")
		delete(additionalProperties, "otfp_shield_time")
		delete(additionalProperties, "otfp_deliver_time")
		delete(additionalProperties, "imgopto_resp_header_bytes")
		delete(additionalProperties, "imgopto_resp_body_bytes")
		delete(additionalProperties, "imgopto_shield_resp_header_bytes")
		delete(additionalProperties, "imgopto_shield_resp_body_bytes")
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
		delete(additionalProperties, "status_406")
		delete(additionalProperties, "status_416")
		delete(additionalProperties, "status_429")
		delete(additionalProperties, "status_500")
		delete(additionalProperties, "status_501")
		delete(additionalProperties, "status_502")
		delete(additionalProperties, "status_503")
		delete(additionalProperties, "status_504")
		delete(additionalProperties, "status_505")
		delete(additionalProperties, "uncacheable")
		delete(additionalProperties, "pass_time")
		delete(additionalProperties, "tls")
		delete(additionalProperties, "tls_v10")
		delete(additionalProperties, "tls_v11")
		delete(additionalProperties, "tls_v12")
		delete(additionalProperties, "tls_v13")
		delete(additionalProperties, "object_size_1k")
		delete(additionalProperties, "object_size_10k")
		delete(additionalProperties, "object_size_100k")
		delete(additionalProperties, "object_size_1m")
		delete(additionalProperties, "object_size_10m")
		delete(additionalProperties, "object_size_100m")
		delete(additionalProperties, "object_size_1g")
		delete(additionalProperties, "object_size_other")
		delete(additionalProperties, "recv_sub_time")
		delete(additionalProperties, "recv_sub_count")
		delete(additionalProperties, "hash_sub_time")
		delete(additionalProperties, "hash_sub_count")
		delete(additionalProperties, "miss_sub_time")
		delete(additionalProperties, "miss_sub_count")
		delete(additionalProperties, "fetch_sub_time")
		delete(additionalProperties, "fetch_sub_count")
		delete(additionalProperties, "pass_sub_time")
		delete(additionalProperties, "pass_sub_count")
		delete(additionalProperties, "pipe_sub_time")
		delete(additionalProperties, "pipe_sub_count")
		delete(additionalProperties, "deliver_sub_time")
		delete(additionalProperties, "deliver_sub_count")
		delete(additionalProperties, "error_sub_time")
		delete(additionalProperties, "error_sub_count")
		delete(additionalProperties, "hit_sub_time")
		delete(additionalProperties, "hit_sub_count")
		delete(additionalProperties, "prehash_sub_time")
		delete(additionalProperties, "prehash_sub_count")
		delete(additionalProperties, "predeliver_sub_time")
		delete(additionalProperties, "predeliver_sub_count")
		delete(additionalProperties, "hit_resp_body_bytes")
		delete(additionalProperties, "miss_resp_body_bytes")
		delete(additionalProperties, "pass_resp_body_bytes")
		delete(additionalProperties, "compute_req_header_bytes")
		delete(additionalProperties, "compute_req_body_bytes")
		delete(additionalProperties, "compute_resp_header_bytes")
		delete(additionalProperties, "compute_resp_body_bytes")
		delete(additionalProperties, "imgvideo")
		delete(additionalProperties, "imgvideo_frames")
		delete(additionalProperties, "imgvideo_resp_header_bytes")
		delete(additionalProperties, "imgvideo_resp_body_bytes")
		delete(additionalProperties, "imgvideo_shield")
		delete(additionalProperties, "imgvideo_shield_frames")
		delete(additionalProperties, "imgvideo_shield_resp_header_bytes")
		delete(additionalProperties, "imgvideo_shield_resp_body_bytes")
		delete(additionalProperties, "log_bytes")
		delete(additionalProperties, "edge_requests")
		delete(additionalProperties, "edge_resp_header_bytes")
		delete(additionalProperties, "edge_resp_body_bytes")
		delete(additionalProperties, "origin_revalidations")
		delete(additionalProperties, "origin_fetches")
		delete(additionalProperties, "origin_fetch_header_bytes")
		delete(additionalProperties, "origin_fetch_body_bytes")
		delete(additionalProperties, "origin_fetch_resp_header_bytes")
		delete(additionalProperties, "origin_fetch_resp_body_bytes")
		delete(additionalProperties, "shield_revalidations")
		delete(additionalProperties, "shield_fetches")
		delete(additionalProperties, "shield_fetch_header_bytes")
		delete(additionalProperties, "shield_fetch_body_bytes")
		delete(additionalProperties, "shield_fetch_resp_header_bytes")
		delete(additionalProperties, "shield_fetch_resp_body_bytes")
		delete(additionalProperties, "segblock_origin_fetches")
		delete(additionalProperties, "segblock_shield_fetches")
		delete(additionalProperties, "compute_resp_status_1xx")
		delete(additionalProperties, "compute_resp_status_2xx")
		delete(additionalProperties, "compute_resp_status_3xx")
		delete(additionalProperties, "compute_resp_status_4xx")
		delete(additionalProperties, "compute_resp_status_5xx")
		delete(additionalProperties, "edge_hit_requests")
		delete(additionalProperties, "edge_miss_requests")
		delete(additionalProperties, "compute_bereq_header_bytes")
		delete(additionalProperties, "compute_bereq_body_bytes")
		delete(additionalProperties, "compute_beresp_header_bytes")
		delete(additionalProperties, "compute_beresp_body_bytes")
		delete(additionalProperties, "origin_cache_fetches")
		delete(additionalProperties, "shield_cache_fetches")
		delete(additionalProperties, "compute_bereqs")
		delete(additionalProperties, "compute_bereq_errors")
		delete(additionalProperties, "compute_resource_limit_exceeded")
		delete(additionalProperties, "compute_heap_limit_exceeded")
		delete(additionalProperties, "compute_stack_limit_exceeded")
		delete(additionalProperties, "compute_globals_limit_exceeded")
		delete(additionalProperties, "compute_guest_errors")
		delete(additionalProperties, "compute_runtime_errors")
		delete(additionalProperties, "edge_hit_resp_body_bytes")
		delete(additionalProperties, "edge_hit_resp_header_bytes")
		delete(additionalProperties, "edge_miss_resp_body_bytes")
		delete(additionalProperties, "edge_miss_resp_header_bytes")
		delete(additionalProperties, "origin_cache_fetch_resp_body_bytes")
		delete(additionalProperties, "origin_cache_fetch_resp_header_bytes")
		delete(additionalProperties, "shield_hit_requests")
		delete(additionalProperties, "shield_miss_requests")
		delete(additionalProperties, "shield_hit_resp_header_bytes")
		delete(additionalProperties, "shield_hit_resp_body_bytes")
		delete(additionalProperties, "shield_miss_resp_header_bytes")
		delete(additionalProperties, "shield_miss_resp_body_bytes")
		delete(additionalProperties, "websocket_req_header_bytes")
		delete(additionalProperties, "websocket_req_body_bytes")
		delete(additionalProperties, "websocket_resp_header_bytes")
		delete(additionalProperties, "websocket_bereq_header_bytes")
		delete(additionalProperties, "websocket_bereq_body_bytes")
		delete(additionalProperties, "websocket_beresp_header_bytes")
		delete(additionalProperties, "websocket_beresp_body_bytes")
		delete(additionalProperties, "websocket_conn_time_ms")
		delete(additionalProperties, "websocket_resp_body_bytes")
		delete(additionalProperties, "fanout_recv_publishes")
		delete(additionalProperties, "fanout_send_publishes")
		delete(additionalProperties, "kv_store_class_a_operations")
		delete(additionalProperties, "kv_store_class_b_operations")
		delete(additionalProperties, "object_store_class_a_operations")
		delete(additionalProperties, "object_store_class_b_operations")
		delete(additionalProperties, "fanout_req_header_bytes")
		delete(additionalProperties, "fanout_req_body_bytes")
		delete(additionalProperties, "fanout_resp_header_bytes")
		delete(additionalProperties, "fanout_resp_body_bytes")
		delete(additionalProperties, "fanout_bereq_header_bytes")
		delete(additionalProperties, "fanout_bereq_body_bytes")
		delete(additionalProperties, "fanout_beresp_header_bytes")
		delete(additionalProperties, "fanout_beresp_body_bytes")
		delete(additionalProperties, "fanout_conn_time_ms")
		delete(additionalProperties, "ddos_action_limit_streams_connections")
		delete(additionalProperties, "ddos_action_limit_streams_requests")
		delete(additionalProperties, "ddos_action_tarpit_accept")
		delete(additionalProperties, "ddos_action_tarpit")
		delete(additionalProperties, "ddos_action_close")
		delete(additionalProperties, "ddos_action_blackhole")
		delete(additionalProperties, "bot_challenge_starts")
		delete(additionalProperties, "bot_challenge_complete_tokens_passed")
		delete(additionalProperties, "bot_challenge_complete_tokens_failed")
		delete(additionalProperties, "bot_challenge_complete_tokens_checked")
		delete(additionalProperties, "bot_challenge_complete_tokens_disabled")
		delete(additionalProperties, "bot_challenges_issued")
		delete(additionalProperties, "bot_challenges_succeeded")
		delete(additionalProperties, "bot_challenges_failed")
		delete(additionalProperties, "bot_challenge_complete_tokens_issued")
		delete(additionalProperties, "ddos_action_downgrade")
		delete(additionalProperties, "ddos_action_downgraded_connections")
		delete(additionalProperties, "all_hit_requests")
		delete(additionalProperties, "all_miss_requests")
		delete(additionalProperties, "all_pass_requests")
		delete(additionalProperties, "all_error_requests")
		delete(additionalProperties, "all_synth_requests")
		delete(additionalProperties, "all_edge_hit_requests")
		delete(additionalProperties, "all_edge_miss_requests")
		delete(additionalProperties, "all_status_1xx")
		delete(additionalProperties, "all_status_2xx")
		delete(additionalProperties, "all_status_3xx")
		delete(additionalProperties, "all_status_4xx")
		delete(additionalProperties, "all_status_5xx")
		delete(additionalProperties, "origin_offload")
		delete(additionalProperties, "request_denied_get_head_body")
		delete(additionalProperties, "service_ddos_requests_detected")
		delete(additionalProperties, "service_ddos_requests_mitigated")
		delete(additionalProperties, "service_ddos_requests_allowed")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRealtimeEntryAggregated is a helper abstraction for handling nullable realtimeentryaggregated types.
type NullableRealtimeEntryAggregated struct {
	value *RealtimeEntryAggregated
	isSet bool
}

// Get returns the value.
func (v NullableRealtimeEntryAggregated) Get() *RealtimeEntryAggregated {
	return v.value
}

// Set modifies the value.
func (v *NullableRealtimeEntryAggregated) Set(val *RealtimeEntryAggregated) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRealtimeEntryAggregated) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRealtimeEntryAggregated) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRealtimeEntryAggregated returns a pointer to a new instance of NullableRealtimeEntryAggregated.
func NewNullableRealtimeEntryAggregated(val *RealtimeEntryAggregated) *NullableRealtimeEntryAggregated {
	return &NullableRealtimeEntryAggregated{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRealtimeEntryAggregated) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableRealtimeEntryAggregated) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
