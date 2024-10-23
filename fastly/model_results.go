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

// Results The results of stats queries, may be grouped by service depending on endpoint, and aggregated over the appropriate time span.
type Results struct {
	// Number of requests processed.
	Requests *int32 `json:"requests,omitempty"`
	// Number of cache hits.
	Hits *int32 `json:"hits,omitempty"`
	// Total amount of time spent processing cache hits (in seconds).
	HitsTime *float32 `json:"hits_time,omitempty"`
	// Number of cache misses.
	Miss *int32 `json:"miss,omitempty"`
	// Total amount of time spent processing cache misses (in seconds).
	MissTime *float32 `json:"miss_time,omitempty"`
	// Number of requests that passed through the CDN without being cached.
	Pass *int32 `json:"pass,omitempty"`
	// Total amount of time spent processing cache passes (in seconds).
	PassTime *float32 `json:"pass_time,omitempty"`
	// Number of cache errors.
	Errors *int32 `json:"errors,omitempty"`
	// Number of restarts performed.
	Restarts *int32 `json:"restarts,omitempty"`
	// Ratio of cache hits to cache misses (between 0 and 1).
	HitRatio NullableFloat32 `json:"hit_ratio,omitempty"`
	// Total bytes delivered (`resp_header_bytes` + `resp_body_bytes` + `bereq_header_bytes` + `bereq_body_bytes` + `compute_resp_header_bytes` + `compute_resp_body_bytes` + `compute_bereq_header_bytes` + `compute_bereq_body_bytes` + `websocket_resp_header_bytes` + `websocket_resp_body_bytes` + `websocket_bereq_header_bytes` + `websocket_bereq_body_bytes` + `fanout_resp_header_bytes` + `fanout_resp_body_bytes` + `fanout_bereq_header_bytes` + `fanout_bereq_body_bytes`).
	Bandwidth *int32 `json:"bandwidth,omitempty"`
	// Total body bytes delivered (alias for resp_body_bytes).
	BodySize *int32 `json:"body_size,omitempty"`
	// Total header bytes delivered (alias for resp_header_bytes).
	HeaderSize *int32 `json:"header_size,omitempty"`
	// Total body bytes received.
	ReqBodyBytes *int32 `json:"req_body_bytes,omitempty"`
	// Total header bytes received.
	ReqHeaderBytes *int32 `json:"req_header_bytes,omitempty"`
	// Total body bytes delivered (edge_resp_body_bytes + shield_resp_body_bytes).
	RespBodyBytes *int32 `json:"resp_body_bytes,omitempty"`
	// Total header bytes delivered (edge_resp_header_bytes + shield_resp_header_bytes).
	RespHeaderBytes *int32 `json:"resp_header_bytes,omitempty"`
	// Total body bytes sent to origin.
	BereqBodyBytes *int32 `json:"bereq_body_bytes,omitempty"`
	// Total header bytes sent to origin.
	BereqHeaderBytes *int32 `json:"bereq_header_bytes,omitempty"`
	// Number of requests that were designated uncachable.
	Uncacheable *int32 `json:"uncacheable,omitempty"`
	// Optional. Pipe operations performed (legacy feature).
	Pipe *int32 `json:"pipe,omitempty"`
	// Number of requests that returned a synthetic response (i.e., response objects created with the `synthetic` VCL statement).
	Synth *int32 `json:"synth,omitempty"`
	// Number of requests that were received over TLS.
	TLS *int32 `json:"tls,omitempty"`
	// Number of requests received over TLS 1.0.
	TLSV10 *int32 `json:"tls_v10,omitempty"`
	// Number of requests received over TLS 1.1.
	TLSV11 *int32 `json:"tls_v11,omitempty"`
	// Number of requests received over TLS 1.2.
	TLSV12 *int32 `json:"tls_v12,omitempty"`
	// Number of requests received over TLS 1.3.
	TLSV13 *int32 `json:"tls_v13,omitempty"`
	// Number of requests sent by end users to Fastly.
	EdgeRequests *int32 `json:"edge_requests,omitempty"`
	// Total header bytes delivered from Fastly to the end user.
	EdgeRespHeaderBytes *int32 `json:"edge_resp_header_bytes,omitempty"`
	// Total body bytes delivered from Fastly to the end user.
	EdgeRespBodyBytes *int32 `json:"edge_resp_body_bytes,omitempty"`
	// Number of requests sent by end users to Fastly that resulted in a hit at the edge.
	EdgeHitRequests *int32 `json:"edge_hit_requests,omitempty"`
	// Number of requests sent by end users to Fastly that resulted in a miss at the edge.
	EdgeMissRequests *int32 `json:"edge_miss_requests,omitempty"`
	// Number of requests sent to origin.
	OriginFetches *int32 `json:"origin_fetches,omitempty"`
	// Total request header bytes sent to origin.
	OriginFetchHeaderBytes *int32 `json:"origin_fetch_header_bytes,omitempty"`
	// Total request body bytes sent to origin.
	OriginFetchBodyBytes *int32 `json:"origin_fetch_body_bytes,omitempty"`
	// Total header bytes received from origin.
	OriginFetchRespHeaderBytes *int32 `json:"origin_fetch_resp_header_bytes,omitempty"`
	// Total body bytes received from origin.
	OriginFetchRespBodyBytes *int32 `json:"origin_fetch_resp_body_bytes,omitempty"`
	// Number of responses received from origin with a `304` status code in response to an `If-Modified-Since` or `If-None-Match` request. Under regular scenarios, a revalidation will imply a cache hit. However, if using Fastly Image Optimizer or segmented caching this may result in a cache miss.
	OriginRevalidations *int32 `json:"origin_revalidations,omitempty"`
	// The total number of completed requests made to backends (origins) that returned cacheable content.
	OriginCacheFetches *int32 `json:"origin_cache_fetches,omitempty"`
	// Number of requests from edge to the shield POP.
	Shield *int32 `json:"shield,omitempty"`
	// Total body bytes delivered via a shield.
	ShieldRespBodyBytes *int32 `json:"shield_resp_body_bytes,omitempty"`
	// Total header bytes delivered via a shield.
	ShieldRespHeaderBytes *int32 `json:"shield_resp_header_bytes,omitempty"`
	// Number of requests made from one Fastly POP to another, as part of shielding.
	ShieldFetches *int32 `json:"shield_fetches,omitempty"`
	// Total request header bytes sent to a shield.
	ShieldFetchHeaderBytes *int32 `json:"shield_fetch_header_bytes,omitempty"`
	// Total request body bytes sent to a shield.
	ShieldFetchBodyBytes *int32 `json:"shield_fetch_body_bytes,omitempty"`
	// Total response header bytes sent from a shield to the edge.
	ShieldFetchRespHeaderBytes *int32 `json:"shield_fetch_resp_header_bytes,omitempty"`
	// Total response body bytes sent from a shield to the edge.
	ShieldFetchRespBodyBytes *int32 `json:"shield_fetch_resp_body_bytes,omitempty"`
	// Number of responses received from origin with a `304` status code, in response to an `If-Modified-Since` or `If-None-Match` request to a shield. Under regular scenarios, a revalidation will imply a cache hit. However, if using segmented caching this may result in a cache miss.
	ShieldRevalidations *int32 `json:"shield_revalidations,omitempty"`
	// The total number of completed requests made to shields that returned cacheable content.
	ShieldCacheFetches *int32 `json:"shield_cache_fetches,omitempty"`
	// Number of requests that were received over IPv6.
	Ipv6 *int32 `json:"ipv6,omitempty"`
	// Number of responses that came from the Fastly On-the-Fly Packaging service for video-on-demand.
	Otfp *int32 `json:"otfp,omitempty"`
	// Total body bytes delivered from the Fastly On-the-Fly Packaging service for video-on-demand.
	OtfpRespBodyBytes *int32 `json:"otfp_resp_body_bytes,omitempty"`
	// Total header bytes delivered from the Fastly On-the-Fly Packaging service for video-on-demand.
	OtfpRespHeaderBytes *int32 `json:"otfp_resp_header_bytes,omitempty"`
	// Total body bytes delivered via a shield for the Fastly On-the-Fly Packaging service for video-on-demand.
	OtfpShieldRespBodyBytes *int32 `json:"otfp_shield_resp_body_bytes,omitempty"`
	// Total header bytes delivered via a shield for the Fastly On-the-Fly Packaging service for video-on-demand.
	OtfpShieldRespHeaderBytes *int32 `json:"otfp_shield_resp_header_bytes,omitempty"`
	// Number of responses that were manifest files from the Fastly On-the-Fly Packaging service for video-on-demand.
	OtfpManifests *int32 `json:"otfp_manifests,omitempty"`
	// Total amount of time spent delivering a response from the Fastly On-the-Fly Packaging service for video-on-demand (in seconds).
	OtfpDeliverTime *float32 `json:"otfp_deliver_time,omitempty"`
	// Total amount of time spent delivering a response via a shield from the Fastly On-the-Fly Packaging service for video-on-demand (in seconds).
	OtfpShieldTime *float32 `json:"otfp_shield_time,omitempty"`
	// Number of responses with the video segment or video manifest MIME type (i.e., application/x-mpegurl, application/vnd.apple.mpegurl, application/f4m, application/dash+xml, application/vnd.ms-sstr+xml, ideo/mp2t, audio/aac, video/f4f, video/x-flv, video/mp4, audio/mp4).
	Video *int32 `json:"video,omitempty"`
	// Number of responses with the PCI flag turned on.
	Pci *int32 `json:"pci,omitempty"`
	// Number of log lines sent.
	Log *int32 `json:"log,omitempty"`
	// Total log bytes sent.
	LogBytes *int32 `json:"log_bytes,omitempty"`
	// Number of requests received over HTTP/2.
	HTTP2 *int32 `json:"http2,omitempty"`
	// Number of requests received over HTTP/3.
	HTTP3 *int32 `json:"http3,omitempty"`
	// Number of requests that triggered a WAF rule and were logged.
	WafLogged *int32 `json:"waf_logged,omitempty"`
	// Number of requests that triggered a WAF rule and were blocked.
	WafBlocked *int32 `json:"waf_blocked,omitempty"`
	// Number of requests that triggered a WAF rule and were passed.
	WafPassed *int32 `json:"waf_passed,omitempty"`
	// Total body bytes received from requests that triggered a WAF rule.
	AttackReqBodyBytes *int32 `json:"attack_req_body_bytes,omitempty"`
	// Total header bytes received from requests that triggered a WAF rule.
	AttackReqHeaderBytes *int32 `json:"attack_req_header_bytes,omitempty"`
	// Total body bytes received from requests that triggered a WAF rule that was logged.
	AttackLoggedReqBodyBytes *int32 `json:"attack_logged_req_body_bytes,omitempty"`
	// Total header bytes received from requests that triggered a WAF rule that was logged.
	AttackLoggedReqHeaderBytes *int32 `json:"attack_logged_req_header_bytes,omitempty"`
	// Total body bytes received from requests that triggered a WAF rule that was blocked.
	AttackBlockedReqBodyBytes *int32 `json:"attack_blocked_req_body_bytes,omitempty"`
	// Total header bytes received from requests that triggered a WAF rule that was blocked.
	AttackBlockedReqHeaderBytes *int32 `json:"attack_blocked_req_header_bytes,omitempty"`
	// Total body bytes received from requests that triggered a WAF rule that was passed.
	AttackPassedReqBodyBytes *int32 `json:"attack_passed_req_body_bytes,omitempty"`
	// Total header bytes received from requests that triggered a WAF rule that was passed.
	AttackPassedReqHeaderBytes *int32 `json:"attack_passed_req_header_bytes,omitempty"`
	// Total bytes delivered for requests that triggered a WAF rule and returned a synthetic response.
	AttackRespSynthBytes *int32 `json:"attack_resp_synth_bytes,omitempty"`
	// Number of responses that came from the Fastly Image Optimizer service. If the service receives 10 requests for an image, this stat will be 10 regardless of how many times the image was transformed.
	Imgopto *int32 `json:"imgopto,omitempty"`
	// Total body bytes delivered from the Fastly Image Optimizer service, including shield traffic.
	ImgoptoRespBodyBytes *int32 `json:"imgopto_resp_body_bytes,omitempty"`
	// Total header bytes delivered from the Fastly Image Optimizer service, including shield traffic.
	ImgoptoRespHeaderBytes *int32 `json:"imgopto_resp_header_bytes,omitempty"`
	// Number of responses that came from the Fastly Image Optimizer service via a shield.
	ImgoptoShield *int32 `json:"imgopto_shield,omitempty"`
	// Total body bytes delivered via a shield from the Fastly Image Optimizer service.
	ImgoptoShieldRespBodyBytes *int32 `json:"imgopto_shield_resp_body_bytes,omitempty"`
	// Total header bytes delivered via a shield from the Fastly Image Optimizer service.
	ImgoptoShieldRespHeaderBytes *int32 `json:"imgopto_shield_resp_header_bytes,omitempty"`
	// Number of transforms performed by the Fastly Image Optimizer service.
	ImgoptoTransforms *int32 `json:"imgopto_transforms,omitempty"`
	// Number of video responses that came from the Fastly Image Optimizer service.
	Imgvideo *int32 `json:"imgvideo,omitempty"`
	// Number of video frames that came from the Fastly Image Optimizer service. A video frame is an individual image within a sequence of video.
	ImgvideoFrames *int32 `json:"imgvideo_frames,omitempty"`
	// Total header bytes of video delivered from the Fastly Image Optimizer service.
	ImgvideoRespHeaderBytes *int32 `json:"imgvideo_resp_header_bytes,omitempty"`
	// Total body bytes of video delivered from the Fastly Image Optimizer service.
	ImgvideoRespBodyBytes *int32 `json:"imgvideo_resp_body_bytes,omitempty"`
	// Total header bytes of video delivered via a shield from the Fastly Image Optimizer service.
	ImgvideoShieldRespHeaderBytes *int32 `json:"imgvideo_shield_resp_header_bytes,omitempty"`
	// Total body bytes of video delivered via a shield from the Fastly Image Optimizer service.
	ImgvideoShieldRespBodyBytes *int32 `json:"imgvideo_shield_resp_body_bytes,omitempty"`
	// Number of video responses delivered via a shield that came from the Fastly Image Optimizer service.
	ImgvideoShield *int32 `json:"imgvideo_shield,omitempty"`
	// Number of video frames delivered via a shield that came from the Fastly Image Optimizer service. A video frame is an individual image within a sequence of video.
	ImgvideoShieldFrames *int32 `json:"imgvideo_shield_frames,omitempty"`
	// Number of responses sent with status code 200 (Success).
	Status200 *int32 `json:"status_200,omitempty"`
	// Number of responses sent with status code 204 (No Content).
	Status204 *int32 `json:"status_204,omitempty"`
	// Number of responses sent with status code 206 (Partial Content).
	Status206 *int32 `json:"status_206,omitempty"`
	// Number of responses sent with status code 301 (Moved Permanently).
	Status301 *int32 `json:"status_301,omitempty"`
	// Number of responses sent with status code 302 (Found).
	Status302 *int32 `json:"status_302,omitempty"`
	// Number of responses sent with status code 304 (Not Modified).
	Status304 *int32 `json:"status_304,omitempty"`
	// Number of responses sent with status code 400 (Bad Request).
	Status400 *int32 `json:"status_400,omitempty"`
	// Number of responses sent with status code 401 (Unauthorized).
	Status401 *int32 `json:"status_401,omitempty"`
	// Number of responses sent with status code 403 (Forbidden).
	Status403 *int32 `json:"status_403,omitempty"`
	// Number of responses sent with status code 404 (Not Found).
	Status404 *int32 `json:"status_404,omitempty"`
	// Number of responses sent with status code 406 (Not Acceptable).
	Status406 *int32 `json:"status_406,omitempty"`
	// Number of responses sent with status code 416 (Range Not Satisfiable).
	Status416 *int32 `json:"status_416,omitempty"`
	// Number of responses sent with status code 429 (Too Many Requests).
	Status429 *int32 `json:"status_429,omitempty"`
	// Number of responses sent with status code 500 (Internal Server Error).
	Status500 *int32 `json:"status_500,omitempty"`
	// Number of responses sent with status code 501 (Not Implemented).
	Status501 *int32 `json:"status_501,omitempty"`
	// Number of responses sent with status code 502 (Bad Gateway).
	Status502 *int32 `json:"status_502,omitempty"`
	// Number of responses sent with status code 503 (Service Unavailable).
	Status503 *int32 `json:"status_503,omitempty"`
	// Number of responses sent with status code 504 (Gateway Timeout).
	Status504 *int32 `json:"status_504,omitempty"`
	// Number of responses sent with status code 505 (HTTP Version Not Supported).
	Status505 *int32 `json:"status_505,omitempty"`
	// Number of \"Informational\" category status codes delivered.
	Status1xx *int32 `json:"status_1xx,omitempty"`
	// Number of \"Success\" status codes delivered.
	Status2xx *int32 `json:"status_2xx,omitempty"`
	// Number of \"Redirection\" codes delivered.
	Status3xx *int32 `json:"status_3xx,omitempty"`
	// Number of \"Client Error\" codes delivered.
	Status4xx *int32 `json:"status_4xx,omitempty"`
	// Number of \"Server Error\" codes delivered.
	Status5xx *int32 `json:"status_5xx,omitempty"`
	// Number of objects served that were under 1KB in size.
	ObjectSize1k *int32 `json:"object_size_1k,omitempty"`
	// Number of objects served that were between 1KB and 10KB in size.
	ObjectSize10k *int32 `json:"object_size_10k,omitempty"`
	// Number of objects served that were between 10KB and 100KB in size.
	ObjectSize100k *int32 `json:"object_size_100k,omitempty"`
	// Number of objects served that were between 100KB and 1MB in size.
	ObjectSize1m *int32 `json:"object_size_1m,omitempty"`
	// Number of objects served that were between 1MB and 10MB in size.
	ObjectSize10m *int32 `json:"object_size_10m,omitempty"`
	// Number of objects served that were between 10MB and 100MB in size.
	ObjectSize100m *int32 `json:"object_size_100m,omitempty"`
	// Number of objects served that were between 100MB and 1GB in size.
	ObjectSize1g *int32 `json:"object_size_1g,omitempty"`
	// Time spent inside the `vcl_recv` Varnish subroutine (in seconds).
	RecvSubTime *float32 `json:"recv_sub_time,omitempty"`
	// Number of executions of the `vcl_recv` Varnish subroutine.
	RecvSubCount *int32 `json:"recv_sub_count,omitempty"`
	// Time spent inside the `vcl_hash` Varnish subroutine (in seconds).
	HashSubTime *float32 `json:"hash_sub_time,omitempty"`
	// Number of executions of the `vcl_hash` Varnish subroutine.
	HashSubCount *int32 `json:"hash_sub_count,omitempty"`
	// Time spent inside the `vcl_miss` Varnish subroutine (in seconds).
	MissSubTime *float32 `json:"miss_sub_time,omitempty"`
	// Number of executions of the `vcl_miss` Varnish subroutine.
	MissSubCount *int32 `json:"miss_sub_count,omitempty"`
	// Time spent inside the `vcl_fetch` Varnish subroutine (in seconds).
	FetchSubTime *float32 `json:"fetch_sub_time,omitempty"`
	// Number of executions of the `vcl_fetch` Varnish subroutine.
	FetchSubCount *int32 `json:"fetch_sub_count,omitempty"`
	// Time spent inside the `vcl_pass` Varnish subroutine (in seconds).
	PassSubTime *float32 `json:"pass_sub_time,omitempty"`
	// Number of executions of the `vcl_pass` Varnish subroutine.
	PassSubCount *int32 `json:"pass_sub_count,omitempty"`
	// Time spent inside the `vcl_pipe` Varnish subroutine (in seconds).
	PipeSubTime *float32 `json:"pipe_sub_time,omitempty"`
	// Number of executions of the `vcl_pipe` Varnish subroutine.
	PipeSubCount *int32 `json:"pipe_sub_count,omitempty"`
	// Time spent inside the `vcl_deliver` Varnish subroutine (in seconds).
	DeliverSubTime *float32 `json:"deliver_sub_time,omitempty"`
	// Number of executions of the `vcl_deliver` Varnish subroutine.
	DeliverSubCount *int32 `json:"deliver_sub_count,omitempty"`
	// Time spent inside the `vcl_error` Varnish subroutine (in seconds).
	ErrorSubTime *float32 `json:"error_sub_time,omitempty"`
	// Number of executions of the `vcl_error` Varnish subroutine.
	ErrorSubCount *int32 `json:"error_sub_count,omitempty"`
	// Time spent inside the `vcl_hit` Varnish subroutine (in seconds).
	HitSubTime *float32 `json:"hit_sub_time,omitempty"`
	// Number of executions of the `vcl_hit` Varnish subroutine.
	HitSubCount *int32 `json:"hit_sub_count,omitempty"`
	// Time spent inside the `vcl_prehash` Varnish subroutine (in seconds).
	PrehashSubTime *float32 `json:"prehash_sub_time,omitempty"`
	// Number of executions of the `vcl_prehash` Varnish subroutine.
	PrehashSubCount *int32 `json:"prehash_sub_count,omitempty"`
	// Time spent inside the `vcl_predeliver` Varnish subroutine (in seconds).
	PredeliverSubTime *float32 `json:"predeliver_sub_time,omitempty"`
	// Number of executions of the `vcl_predeliver` Varnish subroutine.
	PredeliverSubCount *int32 `json:"predeliver_sub_count,omitempty"`
	// Number of bytes transferred during TLS handshake.
	TLSHandshakeSentBytes *int32 `json:"tls_handshake_sent_bytes,omitempty"`
	// Total body bytes delivered for cache hits.
	HitRespBodyBytes *int32 `json:"hit_resp_body_bytes,omitempty"`
	// Total body bytes delivered for cache misses.
	MissRespBodyBytes *int32 `json:"miss_resp_body_bytes,omitempty"`
	// Total body bytes delivered for cache passes.
	PassRespBodyBytes *int32 `json:"pass_resp_body_bytes,omitempty"`
	// Number of `Range` requests to origin for segments of resources when using segmented caching.
	SegblockOriginFetches *int32 `json:"segblock_origin_fetches,omitempty"`
	// Number of `Range` requests to a shield for segments of resources when using segmented caching.
	SegblockShieldFetches *int32 `json:"segblock_shield_fetches,omitempty"`
	// The total number of requests that were received for your service by Fastly.
	ComputeRequests *int32 `json:"compute_requests,omitempty"`
	// The total, actual amount of time used to process your requests, including active CPU time (in milliseconds).
	ComputeRequestTimeMs *float32 `json:"compute_request_time_ms,omitempty"`
	// The total amount of request processing time you will be billed for, measured in 50 millisecond increments.
	ComputeRequestTimeBilledMs *float32 `json:"compute_request_time_billed_ms,omitempty"`
	// The amount of RAM used for your service by Fastly (in bytes).
	ComputeRAMUsed *int32 `json:"compute_ram_used,omitempty"`
	// The amount of active CPU time used to process your requests (in milliseconds).
	ComputeExecutionTimeMs *float32 `json:"compute_execution_time_ms,omitempty"`
	// Total header bytes received by the Compute platform.
	ComputeReqHeaderBytes *int32 `json:"compute_req_header_bytes,omitempty"`
	// Total body bytes received by the Compute platform.
	ComputeReqBodyBytes *int32 `json:"compute_req_body_bytes,omitempty"`
	// Total header bytes sent from Compute to end user.
	ComputeRespHeaderBytes *int32 `json:"compute_resp_header_bytes,omitempty"`
	// Total body bytes sent from Compute to end user.
	ComputeRespBodyBytes *int32 `json:"compute_resp_body_bytes,omitempty"`
	// Number of \"Informational\" category status codes delivered by the Compute platform.
	ComputeRespStatus1xx *int32 `json:"compute_resp_status_1xx,omitempty"`
	// Number of \"Success\" category status codes delivered by the Compute platform.
	ComputeRespStatus2xx *int32 `json:"compute_resp_status_2xx,omitempty"`
	// Number of \"Redirection\" category status codes delivered by the Compute platform.
	ComputeRespStatus3xx *int32 `json:"compute_resp_status_3xx,omitempty"`
	// Number of \"Client Error\" category status codes delivered by the Compute platform.
	ComputeRespStatus4xx *int32 `json:"compute_resp_status_4xx,omitempty"`
	// Number of \"Server Error\" category status codes delivered by the Compute platform.
	ComputeRespStatus5xx *int32 `json:"compute_resp_status_5xx,omitempty"`
	// Total header bytes sent to backends (origins) by the Compute platform.
	ComputeBereqHeaderBytes *int32 `json:"compute_bereq_header_bytes,omitempty"`
	// Total body bytes sent to backends (origins) by the Compute platform.
	ComputeBereqBodyBytes *int32 `json:"compute_bereq_body_bytes,omitempty"`
	// Total header bytes received from backends (origins) by the Compute platform.
	ComputeBerespHeaderBytes *int32 `json:"compute_beresp_header_bytes,omitempty"`
	// Total body bytes received from backends (origins) by the Compute platform.
	ComputeBerespBodyBytes *int32 `json:"compute_beresp_body_bytes,omitempty"`
	// Number of backend requests started.
	ComputeBereqs *int32 `json:"compute_bereqs,omitempty"`
	// Number of backend request errors, including timeouts.
	ComputeBereqErrors *int32 `json:"compute_bereq_errors,omitempty"`
	// Number of times a guest exceeded its resource limit, includes heap, stack, globals, and code execution timeout.
	ComputeResourceLimitExceeded *int32 `json:"compute_resource_limit_exceeded,omitempty"`
	// Number of times a guest exceeded its heap limit.
	ComputeHeapLimitExceeded *int32 `json:"compute_heap_limit_exceeded,omitempty"`
	// Number of times a guest exceeded its stack limit.
	ComputeStackLimitExceeded *int32 `json:"compute_stack_limit_exceeded,omitempty"`
	// Number of times a guest exceeded its globals limit.
	ComputeGlobalsLimitExceeded *int32 `json:"compute_globals_limit_exceeded,omitempty"`
	// Number of times a service experienced a guest code error.
	ComputeGuestErrors *int32 `json:"compute_guest_errors,omitempty"`
	// Number of times a service experienced a guest runtime error.
	ComputeRuntimeErrors *int32 `json:"compute_runtime_errors,omitempty"`
	// Body bytes delivered for edge hits.
	EdgeHitRespBodyBytes *int32 `json:"edge_hit_resp_body_bytes,omitempty"`
	// Header bytes delivered for edge hits.
	EdgeHitRespHeaderBytes *int32 `json:"edge_hit_resp_header_bytes,omitempty"`
	// Body bytes delivered for edge misses.
	EdgeMissRespBodyBytes *int32 `json:"edge_miss_resp_body_bytes,omitempty"`
	// Header bytes delivered for edge misses.
	EdgeMissRespHeaderBytes *int32 `json:"edge_miss_resp_header_bytes,omitempty"`
	// Body bytes received from origin for cacheable content.
	OriginCacheFetchRespBodyBytes *int32 `json:"origin_cache_fetch_resp_body_bytes,omitempty"`
	// Header bytes received from an origin for cacheable content.
	OriginCacheFetchRespHeaderBytes *int32 `json:"origin_cache_fetch_resp_header_bytes,omitempty"`
	// Number of requests that resulted in a hit at a shield.
	ShieldHitRequests *int32 `json:"shield_hit_requests,omitempty"`
	// Number of requests that resulted in a miss at a shield.
	ShieldMissRequests *int32 `json:"shield_miss_requests,omitempty"`
	// Header bytes delivered for shield hits.
	ShieldHitRespHeaderBytes *int32 `json:"shield_hit_resp_header_bytes,omitempty"`
	// Body bytes delivered for shield hits.
	ShieldHitRespBodyBytes *int32 `json:"shield_hit_resp_body_bytes,omitempty"`
	// Header bytes delivered for shield misses.
	ShieldMissRespHeaderBytes *int32 `json:"shield_miss_resp_header_bytes,omitempty"`
	// Body bytes delivered for shield misses.
	ShieldMissRespBodyBytes *int32 `json:"shield_miss_resp_body_bytes,omitempty"`
	// Total header bytes received from end users over passthrough WebSocket connections.
	WebsocketReqHeaderBytes *int32 `json:"websocket_req_header_bytes,omitempty"`
	// Total message content bytes received from end users over passthrough WebSocket connections.
	WebsocketReqBodyBytes *int32 `json:"websocket_req_body_bytes,omitempty"`
	// Total header bytes sent to end users over passthrough WebSocket connections.
	WebsocketRespHeaderBytes *int32 `json:"websocket_resp_header_bytes,omitempty"`
	// Total message content bytes sent to end users over passthrough WebSocket connections.
	WebsocketRespBodyBytes *int32 `json:"websocket_resp_body_bytes,omitempty"`
	// Total header bytes sent to backends over passthrough WebSocket connections.
	WebsocketBereqHeaderBytes *int32 `json:"websocket_bereq_header_bytes,omitempty"`
	// Total message content bytes sent to backends over passthrough WebSocket connections.
	WebsocketBereqBodyBytes *int32 `json:"websocket_bereq_body_bytes,omitempty"`
	// Total header bytes received from backends over passthrough WebSocket connections.
	WebsocketBerespHeaderBytes *int32 `json:"websocket_beresp_header_bytes,omitempty"`
	// Total message content bytes received from backends over passthrough WebSocket connections.
	WebsocketBerespBodyBytes *int32 `json:"websocket_beresp_body_bytes,omitempty"`
	// Total duration of passthrough WebSocket connections with end users.
	WebsocketConnTimeMs *int32 `json:"websocket_conn_time_ms,omitempty"`
	// Total published messages received from the publish API endpoint.
	FanoutRecvPublishes *int32 `json:"fanout_recv_publishes,omitempty"`
	// Total published messages sent to end users.
	FanoutSendPublishes *int32 `json:"fanout_send_publishes,omitempty"`
	// The total number of class a operations for the KV store.
	KvStoreClassAOperations *int32 `json:"kv_store_class_a_operations,omitempty"`
	// The total number of class b operations for the KV store.
	KvStoreClassBOperations *int32 `json:"kv_store_class_b_operations,omitempty"`
	// Use kv_store_class_a_operations.
	// Deprecated
	ObjectStoreClassAOperations *int32 `json:"object_store_class_a_operations,omitempty"`
	// Use kv_store_class_b_operations.
	// Deprecated
	ObjectStoreClassBOperations *int32 `json:"object_store_class_b_operations,omitempty"`
	// Total header bytes received from end users over Fanout connections.
	FanoutReqHeaderBytes *int32 `json:"fanout_req_header_bytes,omitempty"`
	// Total body or message content bytes received from end users over Fanout connections.
	FanoutReqBodyBytes *int32 `json:"fanout_req_body_bytes,omitempty"`
	// Total header bytes sent to end users over Fanout connections.
	FanoutRespHeaderBytes *int32 `json:"fanout_resp_header_bytes,omitempty"`
	// Total body or message content bytes sent to end users over Fanout connections, excluding published message content.
	FanoutRespBodyBytes *int32 `json:"fanout_resp_body_bytes,omitempty"`
	// Total header bytes sent to backends over Fanout connections.
	FanoutBereqHeaderBytes *int32 `json:"fanout_bereq_header_bytes,omitempty"`
	// Total body or message content bytes sent to backends over Fanout connections.
	FanoutBereqBodyBytes *int32 `json:"fanout_bereq_body_bytes,omitempty"`
	// Total header bytes received from backends over Fanout connections.
	FanoutBerespHeaderBytes *int32 `json:"fanout_beresp_header_bytes,omitempty"`
	// Total body or message content bytes received from backends over Fanout connections.
	FanoutBerespBodyBytes *int32 `json:"fanout_beresp_body_bytes,omitempty"`
	// Total duration of Fanout connections with end users.
	FanoutConnTimeMs *int32 `json:"fanout_conn_time_ms,omitempty"`
	// For HTTP/2, the number of connections the limit-streams action was applied to. The limit-streams action caps the allowed number of concurrent streams in a connection.
	DdosActionLimitStreamsConnections *int32 `json:"ddos_action_limit_streams_connections,omitempty"`
	// For HTTP/2, the number of requests made on a connection for which the limit-streams action was taken. The limit-streams action caps the allowed number of concurrent streams in a connection.
	DdosActionLimitStreamsRequests *int32 `json:"ddos_action_limit_streams_requests,omitempty"`
	// The number of times the tarpit-accept action was taken. The tarpit-accept action adds a delay when accepting future connections.
	DdosActionTarpitAccept *int32 `json:"ddos_action_tarpit_accept,omitempty"`
	// The number of times the tarpit action was taken. The tarpit action delays writing the response to the client.
	DdosActionTarpit *int32 `json:"ddos_action_tarpit,omitempty"`
	// The number of times the close action was taken. The close action aborts the connection as soon as possible. The close action takes effect either right after accept, right after the client hello, or right after the response was sent.
	DdosActionClose *int32 `json:"ddos_action_close,omitempty"`
	// The number of times the blackhole action was taken. The blackhole action quietly closes a TCP connection without sending a reset. The blackhole action quietly closes a TCP connection without notifying its peer (all TCP state is dropped).
	DdosActionBlackhole *int32 `json:"ddos_action_blackhole,omitempty"`
	// The number of challenge-start tokens created.
	BotChallengeStarts *int32 `json:"bot_challenge_starts,omitempty"`
	// The number of challenge-complete tokens that passed validation.
	BotChallengeCompleteTokensPassed *int32 `json:"bot_challenge_complete_tokens_passed,omitempty"`
	// The number of challenge-complete tokens that failed validation.
	BotChallengeCompleteTokensFailed *int32 `json:"bot_challenge_complete_tokens_failed,omitempty"`
	// The number of challenge-complete tokens checked.
	BotChallengeCompleteTokensChecked *int32 `json:"bot_challenge_complete_tokens_checked,omitempty"`
	// The number of challenge-complete tokens not checked because the feature was disabled.
	BotChallengeCompleteTokensDisabled *int32 `json:"bot_challenge_complete_tokens_disabled,omitempty"`
	// The number of challenge-complete tokens issued. For example, issuing a challenge-complete token after a series of CAPTCHA challenges ending in success.
	BotChallengeCompleteTokensIssued *int32 `json:"bot_challenge_complete_tokens_issued,omitempty"`
	// The number of challenges issued. For example, the issuance of a CAPTCHA challenge.
	BotChallengesIssued *int32 `json:"bot_challenges_issued,omitempty"`
	// The number of successful challenge solutions processed. For example, a correct CAPTCHA solution.
	BotChallengesSucceeded *int32 `json:"bot_challenges_succeeded,omitempty"`
	// The number of failed challenge solutions processed. For example, an incorrect CAPTCHA solution.
	BotChallengesFailed *int32 `json:"bot_challenges_failed,omitempty"`
	// The number of times the downgrade action was taken. The downgrade action restricts the client to http1.
	DdosActionDowngrade *int32 `json:"ddos_action_downgrade,omitempty"`
	// The number of connections the downgrade action was applied to. The downgrade action restricts the connection to http1.
	DdosActionDowngradedConnections *int32 `json:"ddos_action_downgraded_connections,omitempty"`
	// Number of cache hits for a VCL service running on Compute.
	VclOnComputeHitRequests *int32 `json:"vcl_on_compute_hit_requests,omitempty"`
	// Number of cache misses for a VCL service running on Compute.
	VclOnComputeMissRequests *int32 `json:"vcl_on_compute_miss_requests,omitempty"`
	// Number of requests that passed through the CDN without being cached for a VCL service running on Compute.
	VclOnComputePassRequests *int32 `json:"vcl_on_compute_pass_requests,omitempty"`
	// Number of cache errors for a VCL service running on Compute.
	VclOnComputeErrorRequests *int32 `json:"vcl_on_compute_error_requests,omitempty"`
	// Number of requests that returned a synthetic response (i.e., response objects created with the `synthetic` VCL statement) for a VCL service running on Compute.
	VclOnComputeSynthRequests *int32 `json:"vcl_on_compute_synth_requests,omitempty"`
	// Number of requests sent by end users to Fastly that resulted in a hit at the edge for a VCL service running on Compute.
	VclOnComputeEdgeHitRequests *int32 `json:"vcl_on_compute_edge_hit_requests,omitempty"`
	// Number of requests sent by end users to Fastly that resulted in a miss at the edge for a VCL service running on Compute.
	VclOnComputeEdgeMissRequests *int32 `json:"vcl_on_compute_edge_miss_requests,omitempty"`
	// Number of cache hits for a VCL service.
	AllHitRequests *int32 `json:"all_hit_requests,omitempty"`
	// Number of cache misses for a VCL service.
	AllMissRequests *int32 `json:"all_miss_requests,omitempty"`
	// Number of requests that passed through the CDN without being cached for a VCL service.
	AllPassRequests *int32 `json:"all_pass_requests,omitempty"`
	// Number of cache errors for a VCL service.
	AllErrorRequests *int32 `json:"all_error_requests,omitempty"`
	// Number of requests that returned a synthetic response (i.e., response objects created with the `synthetic` VCL statement) for a VCL service.
	AllSynthRequests *int32 `json:"all_synth_requests,omitempty"`
	// Number of requests sent by end users to Fastly that resulted in a hit at the edge for a VCL service.
	AllEdgeHitRequests *int32 `json:"all_edge_hit_requests,omitempty"`
	// Number of requests sent by end users to Fastly that resulted in a miss at the edge for a VCL service.
	AllEdgeMissRequests *int32 `json:"all_edge_miss_requests,omitempty"`
	// Number of \"Informational\" category status codes delivered for all sources.
	AllStatus1xx *int32 `json:"all_status_1xx,omitempty"`
	// Number of \"Success\" status codes delivered for all sources.
	AllStatus2xx *int32 `json:"all_status_2xx,omitempty"`
	// Number of \"Redirection\" codes delivered for all sources.
	AllStatus3xx *int32 `json:"all_status_3xx,omitempty"`
	// Number of \"Client Error\" codes delivered for all sources.
	AllStatus4xx *int32 `json:"all_status_4xx,omitempty"`
	// Number of \"Server Error\" codes delivered for all sources.
	AllStatus5xx *int32 `json:"all_status_5xx,omitempty"`
	// Origin Offload measures the ratio of bytes served to end users that were cached by Fastly, over the bytes served to end users, between 0 and 1. ((`edge_resp_body_bytes` + `edge_resp_header_bytes`) - (`origin_fetch_resp_body_bytes` + `origin_fetch_resp_header_bytes`)) / (`edge_resp_body_bytes` + `edge_resp_header_bytes`).
	OriginOffload *float32 `json:"origin_offload,omitempty"`
	// Number of requests where Fastly responded with 400 due to the request being a GET or HEAD request containing a body.
	RequestDeniedGetHeadBody *int32  `json:"request_denied_get_head_body,omitempty"`
	ServiceID                *string `json:"service_id,omitempty"`
	// Timestamp for the start of the time period being reported
	StartTime            *int32 `json:"start_time,omitempty"`
	AdditionalProperties map[string]any
}

type _Results Results

// NewResults instantiates a new Results object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResults() *Results {
	this := Results{}
	return &this
}

// NewResultsWithDefaults instantiates a new Results object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResultsWithDefaults() *Results {
	this := Results{}
	return &this
}

// GetRequests returns the Requests field value if set, zero value otherwise.
func (o *Results) GetRequests() int32 {
	if o == nil || o.Requests == nil {
		var ret int32
		return ret
	}
	return *o.Requests
}

// GetRequestsOk returns a tuple with the Requests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetRequestsOk() (*int32, bool) {
	if o == nil || o.Requests == nil {
		return nil, false
	}
	return o.Requests, true
}

// HasRequests returns a boolean if a field has been set.
func (o *Results) HasRequests() bool {
	if o != nil && o.Requests != nil {
		return true
	}

	return false
}

// SetRequests gets a reference to the given int32 and assigns it to the Requests field.
func (o *Results) SetRequests(v int32) {
	o.Requests = &v
}

// GetHits returns the Hits field value if set, zero value otherwise.
func (o *Results) GetHits() int32 {
	if o == nil || o.Hits == nil {
		var ret int32
		return ret
	}
	return *o.Hits
}

// GetHitsOk returns a tuple with the Hits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetHitsOk() (*int32, bool) {
	if o == nil || o.Hits == nil {
		return nil, false
	}
	return o.Hits, true
}

// HasHits returns a boolean if a field has been set.
func (o *Results) HasHits() bool {
	if o != nil && o.Hits != nil {
		return true
	}

	return false
}

// SetHits gets a reference to the given int32 and assigns it to the Hits field.
func (o *Results) SetHits(v int32) {
	o.Hits = &v
}

// GetHitsTime returns the HitsTime field value if set, zero value otherwise.
func (o *Results) GetHitsTime() float32 {
	if o == nil || o.HitsTime == nil {
		var ret float32
		return ret
	}
	return *o.HitsTime
}

// GetHitsTimeOk returns a tuple with the HitsTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetHitsTimeOk() (*float32, bool) {
	if o == nil || o.HitsTime == nil {
		return nil, false
	}
	return o.HitsTime, true
}

// HasHitsTime returns a boolean if a field has been set.
func (o *Results) HasHitsTime() bool {
	if o != nil && o.HitsTime != nil {
		return true
	}

	return false
}

// SetHitsTime gets a reference to the given float32 and assigns it to the HitsTime field.
func (o *Results) SetHitsTime(v float32) {
	o.HitsTime = &v
}

// GetMiss returns the Miss field value if set, zero value otherwise.
func (o *Results) GetMiss() int32 {
	if o == nil || o.Miss == nil {
		var ret int32
		return ret
	}
	return *o.Miss
}

// GetMissOk returns a tuple with the Miss field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetMissOk() (*int32, bool) {
	if o == nil || o.Miss == nil {
		return nil, false
	}
	return o.Miss, true
}

// HasMiss returns a boolean if a field has been set.
func (o *Results) HasMiss() bool {
	if o != nil && o.Miss != nil {
		return true
	}

	return false
}

// SetMiss gets a reference to the given int32 and assigns it to the Miss field.
func (o *Results) SetMiss(v int32) {
	o.Miss = &v
}

// GetMissTime returns the MissTime field value if set, zero value otherwise.
func (o *Results) GetMissTime() float32 {
	if o == nil || o.MissTime == nil {
		var ret float32
		return ret
	}
	return *o.MissTime
}

// GetMissTimeOk returns a tuple with the MissTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetMissTimeOk() (*float32, bool) {
	if o == nil || o.MissTime == nil {
		return nil, false
	}
	return o.MissTime, true
}

// HasMissTime returns a boolean if a field has been set.
func (o *Results) HasMissTime() bool {
	if o != nil && o.MissTime != nil {
		return true
	}

	return false
}

// SetMissTime gets a reference to the given float32 and assigns it to the MissTime field.
func (o *Results) SetMissTime(v float32) {
	o.MissTime = &v
}

// GetPass returns the Pass field value if set, zero value otherwise.
func (o *Results) GetPass() int32 {
	if o == nil || o.Pass == nil {
		var ret int32
		return ret
	}
	return *o.Pass
}

// GetPassOk returns a tuple with the Pass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetPassOk() (*int32, bool) {
	if o == nil || o.Pass == nil {
		return nil, false
	}
	return o.Pass, true
}

// HasPass returns a boolean if a field has been set.
func (o *Results) HasPass() bool {
	if o != nil && o.Pass != nil {
		return true
	}

	return false
}

// SetPass gets a reference to the given int32 and assigns it to the Pass field.
func (o *Results) SetPass(v int32) {
	o.Pass = &v
}

// GetPassTime returns the PassTime field value if set, zero value otherwise.
func (o *Results) GetPassTime() float32 {
	if o == nil || o.PassTime == nil {
		var ret float32
		return ret
	}
	return *o.PassTime
}

// GetPassTimeOk returns a tuple with the PassTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetPassTimeOk() (*float32, bool) {
	if o == nil || o.PassTime == nil {
		return nil, false
	}
	return o.PassTime, true
}

// HasPassTime returns a boolean if a field has been set.
func (o *Results) HasPassTime() bool {
	if o != nil && o.PassTime != nil {
		return true
	}

	return false
}

// SetPassTime gets a reference to the given float32 and assigns it to the PassTime field.
func (o *Results) SetPassTime(v float32) {
	o.PassTime = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *Results) GetErrors() int32 {
	if o == nil || o.Errors == nil {
		var ret int32
		return ret
	}
	return *o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetErrorsOk() (*int32, bool) {
	if o == nil || o.Errors == nil {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *Results) HasErrors() bool {
	if o != nil && o.Errors != nil {
		return true
	}

	return false
}

// SetErrors gets a reference to the given int32 and assigns it to the Errors field.
func (o *Results) SetErrors(v int32) {
	o.Errors = &v
}

// GetRestarts returns the Restarts field value if set, zero value otherwise.
func (o *Results) GetRestarts() int32 {
	if o == nil || o.Restarts == nil {
		var ret int32
		return ret
	}
	return *o.Restarts
}

// GetRestartsOk returns a tuple with the Restarts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetRestartsOk() (*int32, bool) {
	if o == nil || o.Restarts == nil {
		return nil, false
	}
	return o.Restarts, true
}

// HasRestarts returns a boolean if a field has been set.
func (o *Results) HasRestarts() bool {
	if o != nil && o.Restarts != nil {
		return true
	}

	return false
}

// SetRestarts gets a reference to the given int32 and assigns it to the Restarts field.
func (o *Results) SetRestarts(v int32) {
	o.Restarts = &v
}

// GetHitRatio returns the HitRatio field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Results) GetHitRatio() float32 {
	if o == nil || o.HitRatio.Get() == nil {
		var ret float32
		return ret
	}
	return *o.HitRatio.Get()
}

// GetHitRatioOk returns a tuple with the HitRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Results) GetHitRatioOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.HitRatio.Get(), o.HitRatio.IsSet()
}

// HasHitRatio returns a boolean if a field has been set.
func (o *Results) HasHitRatio() bool {
	if o != nil && o.HitRatio.IsSet() {
		return true
	}

	return false
}

// SetHitRatio gets a reference to the given NullableFloat32 and assigns it to the HitRatio field.
func (o *Results) SetHitRatio(v float32) {
	o.HitRatio.Set(&v)
}

// SetHitRatioNil sets the value for HitRatio to be an explicit nil
func (o *Results) SetHitRatioNil() {
	o.HitRatio.Set(nil)
}

// UnsetHitRatio ensures that no value is present for HitRatio, not even an explicit nil
func (o *Results) UnsetHitRatio() {
	o.HitRatio.Unset()
}

// GetBandwidth returns the Bandwidth field value if set, zero value otherwise.
func (o *Results) GetBandwidth() int32 {
	if o == nil || o.Bandwidth == nil {
		var ret int32
		return ret
	}
	return *o.Bandwidth
}

// GetBandwidthOk returns a tuple with the Bandwidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetBandwidthOk() (*int32, bool) {
	if o == nil || o.Bandwidth == nil {
		return nil, false
	}
	return o.Bandwidth, true
}

// HasBandwidth returns a boolean if a field has been set.
func (o *Results) HasBandwidth() bool {
	if o != nil && o.Bandwidth != nil {
		return true
	}

	return false
}

// SetBandwidth gets a reference to the given int32 and assigns it to the Bandwidth field.
func (o *Results) SetBandwidth(v int32) {
	o.Bandwidth = &v
}

// GetBodySize returns the BodySize field value if set, zero value otherwise.
func (o *Results) GetBodySize() int32 {
	if o == nil || o.BodySize == nil {
		var ret int32
		return ret
	}
	return *o.BodySize
}

// GetBodySizeOk returns a tuple with the BodySize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetBodySizeOk() (*int32, bool) {
	if o == nil || o.BodySize == nil {
		return nil, false
	}
	return o.BodySize, true
}

// HasBodySize returns a boolean if a field has been set.
func (o *Results) HasBodySize() bool {
	if o != nil && o.BodySize != nil {
		return true
	}

	return false
}

// SetBodySize gets a reference to the given int32 and assigns it to the BodySize field.
func (o *Results) SetBodySize(v int32) {
	o.BodySize = &v
}

// GetHeaderSize returns the HeaderSize field value if set, zero value otherwise.
func (o *Results) GetHeaderSize() int32 {
	if o == nil || o.HeaderSize == nil {
		var ret int32
		return ret
	}
	return *o.HeaderSize
}

// GetHeaderSizeOk returns a tuple with the HeaderSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetHeaderSizeOk() (*int32, bool) {
	if o == nil || o.HeaderSize == nil {
		return nil, false
	}
	return o.HeaderSize, true
}

// HasHeaderSize returns a boolean if a field has been set.
func (o *Results) HasHeaderSize() bool {
	if o != nil && o.HeaderSize != nil {
		return true
	}

	return false
}

// SetHeaderSize gets a reference to the given int32 and assigns it to the HeaderSize field.
func (o *Results) SetHeaderSize(v int32) {
	o.HeaderSize = &v
}

// GetReqBodyBytes returns the ReqBodyBytes field value if set, zero value otherwise.
func (o *Results) GetReqBodyBytes() int32 {
	if o == nil || o.ReqBodyBytes == nil {
		var ret int32
		return ret
	}
	return *o.ReqBodyBytes
}

// GetReqBodyBytesOk returns a tuple with the ReqBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetReqBodyBytesOk() (*int32, bool) {
	if o == nil || o.ReqBodyBytes == nil {
		return nil, false
	}
	return o.ReqBodyBytes, true
}

// HasReqBodyBytes returns a boolean if a field has been set.
func (o *Results) HasReqBodyBytes() bool {
	if o != nil && o.ReqBodyBytes != nil {
		return true
	}

	return false
}

// SetReqBodyBytes gets a reference to the given int32 and assigns it to the ReqBodyBytes field.
func (o *Results) SetReqBodyBytes(v int32) {
	o.ReqBodyBytes = &v
}

// GetReqHeaderBytes returns the ReqHeaderBytes field value if set, zero value otherwise.
func (o *Results) GetReqHeaderBytes() int32 {
	if o == nil || o.ReqHeaderBytes == nil {
		var ret int32
		return ret
	}
	return *o.ReqHeaderBytes
}

// GetReqHeaderBytesOk returns a tuple with the ReqHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetReqHeaderBytesOk() (*int32, bool) {
	if o == nil || o.ReqHeaderBytes == nil {
		return nil, false
	}
	return o.ReqHeaderBytes, true
}

// HasReqHeaderBytes returns a boolean if a field has been set.
func (o *Results) HasReqHeaderBytes() bool {
	if o != nil && o.ReqHeaderBytes != nil {
		return true
	}

	return false
}

// SetReqHeaderBytes gets a reference to the given int32 and assigns it to the ReqHeaderBytes field.
func (o *Results) SetReqHeaderBytes(v int32) {
	o.ReqHeaderBytes = &v
}

// GetRespBodyBytes returns the RespBodyBytes field value if set, zero value otherwise.
func (o *Results) GetRespBodyBytes() int32 {
	if o == nil || o.RespBodyBytes == nil {
		var ret int32
		return ret
	}
	return *o.RespBodyBytes
}

// GetRespBodyBytesOk returns a tuple with the RespBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetRespBodyBytesOk() (*int32, bool) {
	if o == nil || o.RespBodyBytes == nil {
		return nil, false
	}
	return o.RespBodyBytes, true
}

// HasRespBodyBytes returns a boolean if a field has been set.
func (o *Results) HasRespBodyBytes() bool {
	if o != nil && o.RespBodyBytes != nil {
		return true
	}

	return false
}

// SetRespBodyBytes gets a reference to the given int32 and assigns it to the RespBodyBytes field.
func (o *Results) SetRespBodyBytes(v int32) {
	o.RespBodyBytes = &v
}

// GetRespHeaderBytes returns the RespHeaderBytes field value if set, zero value otherwise.
func (o *Results) GetRespHeaderBytes() int32 {
	if o == nil || o.RespHeaderBytes == nil {
		var ret int32
		return ret
	}
	return *o.RespHeaderBytes
}

// GetRespHeaderBytesOk returns a tuple with the RespHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetRespHeaderBytesOk() (*int32, bool) {
	if o == nil || o.RespHeaderBytes == nil {
		return nil, false
	}
	return o.RespHeaderBytes, true
}

// HasRespHeaderBytes returns a boolean if a field has been set.
func (o *Results) HasRespHeaderBytes() bool {
	if o != nil && o.RespHeaderBytes != nil {
		return true
	}

	return false
}

// SetRespHeaderBytes gets a reference to the given int32 and assigns it to the RespHeaderBytes field.
func (o *Results) SetRespHeaderBytes(v int32) {
	o.RespHeaderBytes = &v
}

// GetBereqBodyBytes returns the BereqBodyBytes field value if set, zero value otherwise.
func (o *Results) GetBereqBodyBytes() int32 {
	if o == nil || o.BereqBodyBytes == nil {
		var ret int32
		return ret
	}
	return *o.BereqBodyBytes
}

// GetBereqBodyBytesOk returns a tuple with the BereqBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetBereqBodyBytesOk() (*int32, bool) {
	if o == nil || o.BereqBodyBytes == nil {
		return nil, false
	}
	return o.BereqBodyBytes, true
}

// HasBereqBodyBytes returns a boolean if a field has been set.
func (o *Results) HasBereqBodyBytes() bool {
	if o != nil && o.BereqBodyBytes != nil {
		return true
	}

	return false
}

// SetBereqBodyBytes gets a reference to the given int32 and assigns it to the BereqBodyBytes field.
func (o *Results) SetBereqBodyBytes(v int32) {
	o.BereqBodyBytes = &v
}

// GetBereqHeaderBytes returns the BereqHeaderBytes field value if set, zero value otherwise.
func (o *Results) GetBereqHeaderBytes() int32 {
	if o == nil || o.BereqHeaderBytes == nil {
		var ret int32
		return ret
	}
	return *o.BereqHeaderBytes
}

// GetBereqHeaderBytesOk returns a tuple with the BereqHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetBereqHeaderBytesOk() (*int32, bool) {
	if o == nil || o.BereqHeaderBytes == nil {
		return nil, false
	}
	return o.BereqHeaderBytes, true
}

// HasBereqHeaderBytes returns a boolean if a field has been set.
func (o *Results) HasBereqHeaderBytes() bool {
	if o != nil && o.BereqHeaderBytes != nil {
		return true
	}

	return false
}

// SetBereqHeaderBytes gets a reference to the given int32 and assigns it to the BereqHeaderBytes field.
func (o *Results) SetBereqHeaderBytes(v int32) {
	o.BereqHeaderBytes = &v
}

// GetUncacheable returns the Uncacheable field value if set, zero value otherwise.
func (o *Results) GetUncacheable() int32 {
	if o == nil || o.Uncacheable == nil {
		var ret int32
		return ret
	}
	return *o.Uncacheable
}

// GetUncacheableOk returns a tuple with the Uncacheable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetUncacheableOk() (*int32, bool) {
	if o == nil || o.Uncacheable == nil {
		return nil, false
	}
	return o.Uncacheable, true
}

// HasUncacheable returns a boolean if a field has been set.
func (o *Results) HasUncacheable() bool {
	if o != nil && o.Uncacheable != nil {
		return true
	}

	return false
}

// SetUncacheable gets a reference to the given int32 and assigns it to the Uncacheable field.
func (o *Results) SetUncacheable(v int32) {
	o.Uncacheable = &v
}

// GetPipe returns the Pipe field value if set, zero value otherwise.
func (o *Results) GetPipe() int32 {
	if o == nil || o.Pipe == nil {
		var ret int32
		return ret
	}
	return *o.Pipe
}

// GetPipeOk returns a tuple with the Pipe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetPipeOk() (*int32, bool) {
	if o == nil || o.Pipe == nil {
		return nil, false
	}
	return o.Pipe, true
}

// HasPipe returns a boolean if a field has been set.
func (o *Results) HasPipe() bool {
	if o != nil && o.Pipe != nil {
		return true
	}

	return false
}

// SetPipe gets a reference to the given int32 and assigns it to the Pipe field.
func (o *Results) SetPipe(v int32) {
	o.Pipe = &v
}

// GetSynth returns the Synth field value if set, zero value otherwise.
func (o *Results) GetSynth() int32 {
	if o == nil || o.Synth == nil {
		var ret int32
		return ret
	}
	return *o.Synth
}

// GetSynthOk returns a tuple with the Synth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetSynthOk() (*int32, bool) {
	if o == nil || o.Synth == nil {
		return nil, false
	}
	return o.Synth, true
}

// HasSynth returns a boolean if a field has been set.
func (o *Results) HasSynth() bool {
	if o != nil && o.Synth != nil {
		return true
	}

	return false
}

// SetSynth gets a reference to the given int32 and assigns it to the Synth field.
func (o *Results) SetSynth(v int32) {
	o.Synth = &v
}

// GetTLS returns the TLS field value if set, zero value otherwise.
func (o *Results) GetTLS() int32 {
	if o == nil || o.TLS == nil {
		var ret int32
		return ret
	}
	return *o.TLS
}

// GetTLSOk returns a tuple with the TLS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetTLSOk() (*int32, bool) {
	if o == nil || o.TLS == nil {
		return nil, false
	}
	return o.TLS, true
}

// HasTLS returns a boolean if a field has been set.
func (o *Results) HasTLS() bool {
	if o != nil && o.TLS != nil {
		return true
	}

	return false
}

// SetTLS gets a reference to the given int32 and assigns it to the TLS field.
func (o *Results) SetTLS(v int32) {
	o.TLS = &v
}

// GetTLSV10 returns the TLSV10 field value if set, zero value otherwise.
func (o *Results) GetTLSV10() int32 {
	if o == nil || o.TLSV10 == nil {
		var ret int32
		return ret
	}
	return *o.TLSV10
}

// GetTLSV10Ok returns a tuple with the TLSV10 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetTLSV10Ok() (*int32, bool) {
	if o == nil || o.TLSV10 == nil {
		return nil, false
	}
	return o.TLSV10, true
}

// HasTLSV10 returns a boolean if a field has been set.
func (o *Results) HasTLSV10() bool {
	if o != nil && o.TLSV10 != nil {
		return true
	}

	return false
}

// SetTLSV10 gets a reference to the given int32 and assigns it to the TLSV10 field.
func (o *Results) SetTLSV10(v int32) {
	o.TLSV10 = &v
}

// GetTLSV11 returns the TLSV11 field value if set, zero value otherwise.
func (o *Results) GetTLSV11() int32 {
	if o == nil || o.TLSV11 == nil {
		var ret int32
		return ret
	}
	return *o.TLSV11
}

// GetTLSV11Ok returns a tuple with the TLSV11 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetTLSV11Ok() (*int32, bool) {
	if o == nil || o.TLSV11 == nil {
		return nil, false
	}
	return o.TLSV11, true
}

// HasTLSV11 returns a boolean if a field has been set.
func (o *Results) HasTLSV11() bool {
	if o != nil && o.TLSV11 != nil {
		return true
	}

	return false
}

// SetTLSV11 gets a reference to the given int32 and assigns it to the TLSV11 field.
func (o *Results) SetTLSV11(v int32) {
	o.TLSV11 = &v
}

// GetTLSV12 returns the TLSV12 field value if set, zero value otherwise.
func (o *Results) GetTLSV12() int32 {
	if o == nil || o.TLSV12 == nil {
		var ret int32
		return ret
	}
	return *o.TLSV12
}

// GetTLSV12Ok returns a tuple with the TLSV12 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetTLSV12Ok() (*int32, bool) {
	if o == nil || o.TLSV12 == nil {
		return nil, false
	}
	return o.TLSV12, true
}

// HasTLSV12 returns a boolean if a field has been set.
func (o *Results) HasTLSV12() bool {
	if o != nil && o.TLSV12 != nil {
		return true
	}

	return false
}

// SetTLSV12 gets a reference to the given int32 and assigns it to the TLSV12 field.
func (o *Results) SetTLSV12(v int32) {
	o.TLSV12 = &v
}

// GetTLSV13 returns the TLSV13 field value if set, zero value otherwise.
func (o *Results) GetTLSV13() int32 {
	if o == nil || o.TLSV13 == nil {
		var ret int32
		return ret
	}
	return *o.TLSV13
}

// GetTLSV13Ok returns a tuple with the TLSV13 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetTLSV13Ok() (*int32, bool) {
	if o == nil || o.TLSV13 == nil {
		return nil, false
	}
	return o.TLSV13, true
}

// HasTLSV13 returns a boolean if a field has been set.
func (o *Results) HasTLSV13() bool {
	if o != nil && o.TLSV13 != nil {
		return true
	}

	return false
}

// SetTLSV13 gets a reference to the given int32 and assigns it to the TLSV13 field.
func (o *Results) SetTLSV13(v int32) {
	o.TLSV13 = &v
}

// GetEdgeRequests returns the EdgeRequests field value if set, zero value otherwise.
func (o *Results) GetEdgeRequests() int32 {
	if o == nil || o.EdgeRequests == nil {
		var ret int32
		return ret
	}
	return *o.EdgeRequests
}

// GetEdgeRequestsOk returns a tuple with the EdgeRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetEdgeRequestsOk() (*int32, bool) {
	if o == nil || o.EdgeRequests == nil {
		return nil, false
	}
	return o.EdgeRequests, true
}

// HasEdgeRequests returns a boolean if a field has been set.
func (o *Results) HasEdgeRequests() bool {
	if o != nil && o.EdgeRequests != nil {
		return true
	}

	return false
}

// SetEdgeRequests gets a reference to the given int32 and assigns it to the EdgeRequests field.
func (o *Results) SetEdgeRequests(v int32) {
	o.EdgeRequests = &v
}

// GetEdgeRespHeaderBytes returns the EdgeRespHeaderBytes field value if set, zero value otherwise.
func (o *Results) GetEdgeRespHeaderBytes() int32 {
	if o == nil || o.EdgeRespHeaderBytes == nil {
		var ret int32
		return ret
	}
	return *o.EdgeRespHeaderBytes
}

// GetEdgeRespHeaderBytesOk returns a tuple with the EdgeRespHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetEdgeRespHeaderBytesOk() (*int32, bool) {
	if o == nil || o.EdgeRespHeaderBytes == nil {
		return nil, false
	}
	return o.EdgeRespHeaderBytes, true
}

// HasEdgeRespHeaderBytes returns a boolean if a field has been set.
func (o *Results) HasEdgeRespHeaderBytes() bool {
	if o != nil && o.EdgeRespHeaderBytes != nil {
		return true
	}

	return false
}

// SetEdgeRespHeaderBytes gets a reference to the given int32 and assigns it to the EdgeRespHeaderBytes field.
func (o *Results) SetEdgeRespHeaderBytes(v int32) {
	o.EdgeRespHeaderBytes = &v
}

// GetEdgeRespBodyBytes returns the EdgeRespBodyBytes field value if set, zero value otherwise.
func (o *Results) GetEdgeRespBodyBytes() int32 {
	if o == nil || o.EdgeRespBodyBytes == nil {
		var ret int32
		return ret
	}
	return *o.EdgeRespBodyBytes
}

// GetEdgeRespBodyBytesOk returns a tuple with the EdgeRespBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetEdgeRespBodyBytesOk() (*int32, bool) {
	if o == nil || o.EdgeRespBodyBytes == nil {
		return nil, false
	}
	return o.EdgeRespBodyBytes, true
}

// HasEdgeRespBodyBytes returns a boolean if a field has been set.
func (o *Results) HasEdgeRespBodyBytes() bool {
	if o != nil && o.EdgeRespBodyBytes != nil {
		return true
	}

	return false
}

// SetEdgeRespBodyBytes gets a reference to the given int32 and assigns it to the EdgeRespBodyBytes field.
func (o *Results) SetEdgeRespBodyBytes(v int32) {
	o.EdgeRespBodyBytes = &v
}

// GetEdgeHitRequests returns the EdgeHitRequests field value if set, zero value otherwise.
func (o *Results) GetEdgeHitRequests() int32 {
	if o == nil || o.EdgeHitRequests == nil {
		var ret int32
		return ret
	}
	return *o.EdgeHitRequests
}

// GetEdgeHitRequestsOk returns a tuple with the EdgeHitRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetEdgeHitRequestsOk() (*int32, bool) {
	if o == nil || o.EdgeHitRequests == nil {
		return nil, false
	}
	return o.EdgeHitRequests, true
}

// HasEdgeHitRequests returns a boolean if a field has been set.
func (o *Results) HasEdgeHitRequests() bool {
	if o != nil && o.EdgeHitRequests != nil {
		return true
	}

	return false
}

// SetEdgeHitRequests gets a reference to the given int32 and assigns it to the EdgeHitRequests field.
func (o *Results) SetEdgeHitRequests(v int32) {
	o.EdgeHitRequests = &v
}

// GetEdgeMissRequests returns the EdgeMissRequests field value if set, zero value otherwise.
func (o *Results) GetEdgeMissRequests() int32 {
	if o == nil || o.EdgeMissRequests == nil {
		var ret int32
		return ret
	}
	return *o.EdgeMissRequests
}

// GetEdgeMissRequestsOk returns a tuple with the EdgeMissRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetEdgeMissRequestsOk() (*int32, bool) {
	if o == nil || o.EdgeMissRequests == nil {
		return nil, false
	}
	return o.EdgeMissRequests, true
}

// HasEdgeMissRequests returns a boolean if a field has been set.
func (o *Results) HasEdgeMissRequests() bool {
	if o != nil && o.EdgeMissRequests != nil {
		return true
	}

	return false
}

// SetEdgeMissRequests gets a reference to the given int32 and assigns it to the EdgeMissRequests field.
func (o *Results) SetEdgeMissRequests(v int32) {
	o.EdgeMissRequests = &v
}

// GetOriginFetches returns the OriginFetches field value if set, zero value otherwise.
func (o *Results) GetOriginFetches() int32 {
	if o == nil || o.OriginFetches == nil {
		var ret int32
		return ret
	}
	return *o.OriginFetches
}

// GetOriginFetchesOk returns a tuple with the OriginFetches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetOriginFetchesOk() (*int32, bool) {
	if o == nil || o.OriginFetches == nil {
		return nil, false
	}
	return o.OriginFetches, true
}

// HasOriginFetches returns a boolean if a field has been set.
func (o *Results) HasOriginFetches() bool {
	if o != nil && o.OriginFetches != nil {
		return true
	}

	return false
}

// SetOriginFetches gets a reference to the given int32 and assigns it to the OriginFetches field.
func (o *Results) SetOriginFetches(v int32) {
	o.OriginFetches = &v
}

// GetOriginFetchHeaderBytes returns the OriginFetchHeaderBytes field value if set, zero value otherwise.
func (o *Results) GetOriginFetchHeaderBytes() int32 {
	if o == nil || o.OriginFetchHeaderBytes == nil {
		var ret int32
		return ret
	}
	return *o.OriginFetchHeaderBytes
}

// GetOriginFetchHeaderBytesOk returns a tuple with the OriginFetchHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetOriginFetchHeaderBytesOk() (*int32, bool) {
	if o == nil || o.OriginFetchHeaderBytes == nil {
		return nil, false
	}
	return o.OriginFetchHeaderBytes, true
}

// HasOriginFetchHeaderBytes returns a boolean if a field has been set.
func (o *Results) HasOriginFetchHeaderBytes() bool {
	if o != nil && o.OriginFetchHeaderBytes != nil {
		return true
	}

	return false
}

// SetOriginFetchHeaderBytes gets a reference to the given int32 and assigns it to the OriginFetchHeaderBytes field.
func (o *Results) SetOriginFetchHeaderBytes(v int32) {
	o.OriginFetchHeaderBytes = &v
}

// GetOriginFetchBodyBytes returns the OriginFetchBodyBytes field value if set, zero value otherwise.
func (o *Results) GetOriginFetchBodyBytes() int32 {
	if o == nil || o.OriginFetchBodyBytes == nil {
		var ret int32
		return ret
	}
	return *o.OriginFetchBodyBytes
}

// GetOriginFetchBodyBytesOk returns a tuple with the OriginFetchBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetOriginFetchBodyBytesOk() (*int32, bool) {
	if o == nil || o.OriginFetchBodyBytes == nil {
		return nil, false
	}
	return o.OriginFetchBodyBytes, true
}

// HasOriginFetchBodyBytes returns a boolean if a field has been set.
func (o *Results) HasOriginFetchBodyBytes() bool {
	if o != nil && o.OriginFetchBodyBytes != nil {
		return true
	}

	return false
}

// SetOriginFetchBodyBytes gets a reference to the given int32 and assigns it to the OriginFetchBodyBytes field.
func (o *Results) SetOriginFetchBodyBytes(v int32) {
	o.OriginFetchBodyBytes = &v
}

// GetOriginFetchRespHeaderBytes returns the OriginFetchRespHeaderBytes field value if set, zero value otherwise.
func (o *Results) GetOriginFetchRespHeaderBytes() int32 {
	if o == nil || o.OriginFetchRespHeaderBytes == nil {
		var ret int32
		return ret
	}
	return *o.OriginFetchRespHeaderBytes
}

// GetOriginFetchRespHeaderBytesOk returns a tuple with the OriginFetchRespHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetOriginFetchRespHeaderBytesOk() (*int32, bool) {
	if o == nil || o.OriginFetchRespHeaderBytes == nil {
		return nil, false
	}
	return o.OriginFetchRespHeaderBytes, true
}

// HasOriginFetchRespHeaderBytes returns a boolean if a field has been set.
func (o *Results) HasOriginFetchRespHeaderBytes() bool {
	if o != nil && o.OriginFetchRespHeaderBytes != nil {
		return true
	}

	return false
}

// SetOriginFetchRespHeaderBytes gets a reference to the given int32 and assigns it to the OriginFetchRespHeaderBytes field.
func (o *Results) SetOriginFetchRespHeaderBytes(v int32) {
	o.OriginFetchRespHeaderBytes = &v
}

// GetOriginFetchRespBodyBytes returns the OriginFetchRespBodyBytes field value if set, zero value otherwise.
func (o *Results) GetOriginFetchRespBodyBytes() int32 {
	if o == nil || o.OriginFetchRespBodyBytes == nil {
		var ret int32
		return ret
	}
	return *o.OriginFetchRespBodyBytes
}

// GetOriginFetchRespBodyBytesOk returns a tuple with the OriginFetchRespBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetOriginFetchRespBodyBytesOk() (*int32, bool) {
	if o == nil || o.OriginFetchRespBodyBytes == nil {
		return nil, false
	}
	return o.OriginFetchRespBodyBytes, true
}

// HasOriginFetchRespBodyBytes returns a boolean if a field has been set.
func (o *Results) HasOriginFetchRespBodyBytes() bool {
	if o != nil && o.OriginFetchRespBodyBytes != nil {
		return true
	}

	return false
}

// SetOriginFetchRespBodyBytes gets a reference to the given int32 and assigns it to the OriginFetchRespBodyBytes field.
func (o *Results) SetOriginFetchRespBodyBytes(v int32) {
	o.OriginFetchRespBodyBytes = &v
}

// GetOriginRevalidations returns the OriginRevalidations field value if set, zero value otherwise.
func (o *Results) GetOriginRevalidations() int32 {
	if o == nil || o.OriginRevalidations == nil {
		var ret int32
		return ret
	}
	return *o.OriginRevalidations
}

// GetOriginRevalidationsOk returns a tuple with the OriginRevalidations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetOriginRevalidationsOk() (*int32, bool) {
	if o == nil || o.OriginRevalidations == nil {
		return nil, false
	}
	return o.OriginRevalidations, true
}

// HasOriginRevalidations returns a boolean if a field has been set.
func (o *Results) HasOriginRevalidations() bool {
	if o != nil && o.OriginRevalidations != nil {
		return true
	}

	return false
}

// SetOriginRevalidations gets a reference to the given int32 and assigns it to the OriginRevalidations field.
func (o *Results) SetOriginRevalidations(v int32) {
	o.OriginRevalidations = &v
}

// GetOriginCacheFetches returns the OriginCacheFetches field value if set, zero value otherwise.
func (o *Results) GetOriginCacheFetches() int32 {
	if o == nil || o.OriginCacheFetches == nil {
		var ret int32
		return ret
	}
	return *o.OriginCacheFetches
}

// GetOriginCacheFetchesOk returns a tuple with the OriginCacheFetches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetOriginCacheFetchesOk() (*int32, bool) {
	if o == nil || o.OriginCacheFetches == nil {
		return nil, false
	}
	return o.OriginCacheFetches, true
}

// HasOriginCacheFetches returns a boolean if a field has been set.
func (o *Results) HasOriginCacheFetches() bool {
	if o != nil && o.OriginCacheFetches != nil {
		return true
	}

	return false
}

// SetOriginCacheFetches gets a reference to the given int32 and assigns it to the OriginCacheFetches field.
func (o *Results) SetOriginCacheFetches(v int32) {
	o.OriginCacheFetches = &v
}

// GetShield returns the Shield field value if set, zero value otherwise.
func (o *Results) GetShield() int32 {
	if o == nil || o.Shield == nil {
		var ret int32
		return ret
	}
	return *o.Shield
}

// GetShieldOk returns a tuple with the Shield field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetShieldOk() (*int32, bool) {
	if o == nil || o.Shield == nil {
		return nil, false
	}
	return o.Shield, true
}

// HasShield returns a boolean if a field has been set.
func (o *Results) HasShield() bool {
	if o != nil && o.Shield != nil {
		return true
	}

	return false
}

// SetShield gets a reference to the given int32 and assigns it to the Shield field.
func (o *Results) SetShield(v int32) {
	o.Shield = &v
}

// GetShieldRespBodyBytes returns the ShieldRespBodyBytes field value if set, zero value otherwise.
func (o *Results) GetShieldRespBodyBytes() int32 {
	if o == nil || o.ShieldRespBodyBytes == nil {
		var ret int32
		return ret
	}
	return *o.ShieldRespBodyBytes
}

// GetShieldRespBodyBytesOk returns a tuple with the ShieldRespBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetShieldRespBodyBytesOk() (*int32, bool) {
	if o == nil || o.ShieldRespBodyBytes == nil {
		return nil, false
	}
	return o.ShieldRespBodyBytes, true
}

// HasShieldRespBodyBytes returns a boolean if a field has been set.
func (o *Results) HasShieldRespBodyBytes() bool {
	if o != nil && o.ShieldRespBodyBytes != nil {
		return true
	}

	return false
}

// SetShieldRespBodyBytes gets a reference to the given int32 and assigns it to the ShieldRespBodyBytes field.
func (o *Results) SetShieldRespBodyBytes(v int32) {
	o.ShieldRespBodyBytes = &v
}

// GetShieldRespHeaderBytes returns the ShieldRespHeaderBytes field value if set, zero value otherwise.
func (o *Results) GetShieldRespHeaderBytes() int32 {
	if o == nil || o.ShieldRespHeaderBytes == nil {
		var ret int32
		return ret
	}
	return *o.ShieldRespHeaderBytes
}

// GetShieldRespHeaderBytesOk returns a tuple with the ShieldRespHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetShieldRespHeaderBytesOk() (*int32, bool) {
	if o == nil || o.ShieldRespHeaderBytes == nil {
		return nil, false
	}
	return o.ShieldRespHeaderBytes, true
}

// HasShieldRespHeaderBytes returns a boolean if a field has been set.
func (o *Results) HasShieldRespHeaderBytes() bool {
	if o != nil && o.ShieldRespHeaderBytes != nil {
		return true
	}

	return false
}

// SetShieldRespHeaderBytes gets a reference to the given int32 and assigns it to the ShieldRespHeaderBytes field.
func (o *Results) SetShieldRespHeaderBytes(v int32) {
	o.ShieldRespHeaderBytes = &v
}

// GetShieldFetches returns the ShieldFetches field value if set, zero value otherwise.
func (o *Results) GetShieldFetches() int32 {
	if o == nil || o.ShieldFetches == nil {
		var ret int32
		return ret
	}
	return *o.ShieldFetches
}

// GetShieldFetchesOk returns a tuple with the ShieldFetches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetShieldFetchesOk() (*int32, bool) {
	if o == nil || o.ShieldFetches == nil {
		return nil, false
	}
	return o.ShieldFetches, true
}

// HasShieldFetches returns a boolean if a field has been set.
func (o *Results) HasShieldFetches() bool {
	if o != nil && o.ShieldFetches != nil {
		return true
	}

	return false
}

// SetShieldFetches gets a reference to the given int32 and assigns it to the ShieldFetches field.
func (o *Results) SetShieldFetches(v int32) {
	o.ShieldFetches = &v
}

// GetShieldFetchHeaderBytes returns the ShieldFetchHeaderBytes field value if set, zero value otherwise.
func (o *Results) GetShieldFetchHeaderBytes() int32 {
	if o == nil || o.ShieldFetchHeaderBytes == nil {
		var ret int32
		return ret
	}
	return *o.ShieldFetchHeaderBytes
}

// GetShieldFetchHeaderBytesOk returns a tuple with the ShieldFetchHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetShieldFetchHeaderBytesOk() (*int32, bool) {
	if o == nil || o.ShieldFetchHeaderBytes == nil {
		return nil, false
	}
	return o.ShieldFetchHeaderBytes, true
}

// HasShieldFetchHeaderBytes returns a boolean if a field has been set.
func (o *Results) HasShieldFetchHeaderBytes() bool {
	if o != nil && o.ShieldFetchHeaderBytes != nil {
		return true
	}

	return false
}

// SetShieldFetchHeaderBytes gets a reference to the given int32 and assigns it to the ShieldFetchHeaderBytes field.
func (o *Results) SetShieldFetchHeaderBytes(v int32) {
	o.ShieldFetchHeaderBytes = &v
}

// GetShieldFetchBodyBytes returns the ShieldFetchBodyBytes field value if set, zero value otherwise.
func (o *Results) GetShieldFetchBodyBytes() int32 {
	if o == nil || o.ShieldFetchBodyBytes == nil {
		var ret int32
		return ret
	}
	return *o.ShieldFetchBodyBytes
}

// GetShieldFetchBodyBytesOk returns a tuple with the ShieldFetchBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetShieldFetchBodyBytesOk() (*int32, bool) {
	if o == nil || o.ShieldFetchBodyBytes == nil {
		return nil, false
	}
	return o.ShieldFetchBodyBytes, true
}

// HasShieldFetchBodyBytes returns a boolean if a field has been set.
func (o *Results) HasShieldFetchBodyBytes() bool {
	if o != nil && o.ShieldFetchBodyBytes != nil {
		return true
	}

	return false
}

// SetShieldFetchBodyBytes gets a reference to the given int32 and assigns it to the ShieldFetchBodyBytes field.
func (o *Results) SetShieldFetchBodyBytes(v int32) {
	o.ShieldFetchBodyBytes = &v
}

// GetShieldFetchRespHeaderBytes returns the ShieldFetchRespHeaderBytes field value if set, zero value otherwise.
func (o *Results) GetShieldFetchRespHeaderBytes() int32 {
	if o == nil || o.ShieldFetchRespHeaderBytes == nil {
		var ret int32
		return ret
	}
	return *o.ShieldFetchRespHeaderBytes
}

// GetShieldFetchRespHeaderBytesOk returns a tuple with the ShieldFetchRespHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetShieldFetchRespHeaderBytesOk() (*int32, bool) {
	if o == nil || o.ShieldFetchRespHeaderBytes == nil {
		return nil, false
	}
	return o.ShieldFetchRespHeaderBytes, true
}

// HasShieldFetchRespHeaderBytes returns a boolean if a field has been set.
func (o *Results) HasShieldFetchRespHeaderBytes() bool {
	if o != nil && o.ShieldFetchRespHeaderBytes != nil {
		return true
	}

	return false
}

// SetShieldFetchRespHeaderBytes gets a reference to the given int32 and assigns it to the ShieldFetchRespHeaderBytes field.
func (o *Results) SetShieldFetchRespHeaderBytes(v int32) {
	o.ShieldFetchRespHeaderBytes = &v
}

// GetShieldFetchRespBodyBytes returns the ShieldFetchRespBodyBytes field value if set, zero value otherwise.
func (o *Results) GetShieldFetchRespBodyBytes() int32 {
	if o == nil || o.ShieldFetchRespBodyBytes == nil {
		var ret int32
		return ret
	}
	return *o.ShieldFetchRespBodyBytes
}

// GetShieldFetchRespBodyBytesOk returns a tuple with the ShieldFetchRespBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetShieldFetchRespBodyBytesOk() (*int32, bool) {
	if o == nil || o.ShieldFetchRespBodyBytes == nil {
		return nil, false
	}
	return o.ShieldFetchRespBodyBytes, true
}

// HasShieldFetchRespBodyBytes returns a boolean if a field has been set.
func (o *Results) HasShieldFetchRespBodyBytes() bool {
	if o != nil && o.ShieldFetchRespBodyBytes != nil {
		return true
	}

	return false
}

// SetShieldFetchRespBodyBytes gets a reference to the given int32 and assigns it to the ShieldFetchRespBodyBytes field.
func (o *Results) SetShieldFetchRespBodyBytes(v int32) {
	o.ShieldFetchRespBodyBytes = &v
}

// GetShieldRevalidations returns the ShieldRevalidations field value if set, zero value otherwise.
func (o *Results) GetShieldRevalidations() int32 {
	if o == nil || o.ShieldRevalidations == nil {
		var ret int32
		return ret
	}
	return *o.ShieldRevalidations
}

// GetShieldRevalidationsOk returns a tuple with the ShieldRevalidations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetShieldRevalidationsOk() (*int32, bool) {
	if o == nil || o.ShieldRevalidations == nil {
		return nil, false
	}
	return o.ShieldRevalidations, true
}

// HasShieldRevalidations returns a boolean if a field has been set.
func (o *Results) HasShieldRevalidations() bool {
	if o != nil && o.ShieldRevalidations != nil {
		return true
	}

	return false
}

// SetShieldRevalidations gets a reference to the given int32 and assigns it to the ShieldRevalidations field.
func (o *Results) SetShieldRevalidations(v int32) {
	o.ShieldRevalidations = &v
}

// GetShieldCacheFetches returns the ShieldCacheFetches field value if set, zero value otherwise.
func (o *Results) GetShieldCacheFetches() int32 {
	if o == nil || o.ShieldCacheFetches == nil {
		var ret int32
		return ret
	}
	return *o.ShieldCacheFetches
}

// GetShieldCacheFetchesOk returns a tuple with the ShieldCacheFetches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetShieldCacheFetchesOk() (*int32, bool) {
	if o == nil || o.ShieldCacheFetches == nil {
		return nil, false
	}
	return o.ShieldCacheFetches, true
}

// HasShieldCacheFetches returns a boolean if a field has been set.
func (o *Results) HasShieldCacheFetches() bool {
	if o != nil && o.ShieldCacheFetches != nil {
		return true
	}

	return false
}

// SetShieldCacheFetches gets a reference to the given int32 and assigns it to the ShieldCacheFetches field.
func (o *Results) SetShieldCacheFetches(v int32) {
	o.ShieldCacheFetches = &v
}

// GetIpv6 returns the Ipv6 field value if set, zero value otherwise.
func (o *Results) GetIpv6() int32 {
	if o == nil || o.Ipv6 == nil {
		var ret int32
		return ret
	}
	return *o.Ipv6
}

// GetIpv6Ok returns a tuple with the Ipv6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetIpv6Ok() (*int32, bool) {
	if o == nil || o.Ipv6 == nil {
		return nil, false
	}
	return o.Ipv6, true
}

// HasIpv6 returns a boolean if a field has been set.
func (o *Results) HasIpv6() bool {
	if o != nil && o.Ipv6 != nil {
		return true
	}

	return false
}

// SetIpv6 gets a reference to the given int32 and assigns it to the Ipv6 field.
func (o *Results) SetIpv6(v int32) {
	o.Ipv6 = &v
}

// GetOtfp returns the Otfp field value if set, zero value otherwise.
func (o *Results) GetOtfp() int32 {
	if o == nil || o.Otfp == nil {
		var ret int32
		return ret
	}
	return *o.Otfp
}

// GetOtfpOk returns a tuple with the Otfp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetOtfpOk() (*int32, bool) {
	if o == nil || o.Otfp == nil {
		return nil, false
	}
	return o.Otfp, true
}

// HasOtfp returns a boolean if a field has been set.
func (o *Results) HasOtfp() bool {
	if o != nil && o.Otfp != nil {
		return true
	}

	return false
}

// SetOtfp gets a reference to the given int32 and assigns it to the Otfp field.
func (o *Results) SetOtfp(v int32) {
	o.Otfp = &v
}

// GetOtfpRespBodyBytes returns the OtfpRespBodyBytes field value if set, zero value otherwise.
func (o *Results) GetOtfpRespBodyBytes() int32 {
	if o == nil || o.OtfpRespBodyBytes == nil {
		var ret int32
		return ret
	}
	return *o.OtfpRespBodyBytes
}

// GetOtfpRespBodyBytesOk returns a tuple with the OtfpRespBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetOtfpRespBodyBytesOk() (*int32, bool) {
	if o == nil || o.OtfpRespBodyBytes == nil {
		return nil, false
	}
	return o.OtfpRespBodyBytes, true
}

// HasOtfpRespBodyBytes returns a boolean if a field has been set.
func (o *Results) HasOtfpRespBodyBytes() bool {
	if o != nil && o.OtfpRespBodyBytes != nil {
		return true
	}

	return false
}

// SetOtfpRespBodyBytes gets a reference to the given int32 and assigns it to the OtfpRespBodyBytes field.
func (o *Results) SetOtfpRespBodyBytes(v int32) {
	o.OtfpRespBodyBytes = &v
}

// GetOtfpRespHeaderBytes returns the OtfpRespHeaderBytes field value if set, zero value otherwise.
func (o *Results) GetOtfpRespHeaderBytes() int32 {
	if o == nil || o.OtfpRespHeaderBytes == nil {
		var ret int32
		return ret
	}
	return *o.OtfpRespHeaderBytes
}

// GetOtfpRespHeaderBytesOk returns a tuple with the OtfpRespHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetOtfpRespHeaderBytesOk() (*int32, bool) {
	if o == nil || o.OtfpRespHeaderBytes == nil {
		return nil, false
	}
	return o.OtfpRespHeaderBytes, true
}

// HasOtfpRespHeaderBytes returns a boolean if a field has been set.
func (o *Results) HasOtfpRespHeaderBytes() bool {
	if o != nil && o.OtfpRespHeaderBytes != nil {
		return true
	}

	return false
}

// SetOtfpRespHeaderBytes gets a reference to the given int32 and assigns it to the OtfpRespHeaderBytes field.
func (o *Results) SetOtfpRespHeaderBytes(v int32) {
	o.OtfpRespHeaderBytes = &v
}

// GetOtfpShieldRespBodyBytes returns the OtfpShieldRespBodyBytes field value if set, zero value otherwise.
func (o *Results) GetOtfpShieldRespBodyBytes() int32 {
	if o == nil || o.OtfpShieldRespBodyBytes == nil {
		var ret int32
		return ret
	}
	return *o.OtfpShieldRespBodyBytes
}

// GetOtfpShieldRespBodyBytesOk returns a tuple with the OtfpShieldRespBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetOtfpShieldRespBodyBytesOk() (*int32, bool) {
	if o == nil || o.OtfpShieldRespBodyBytes == nil {
		return nil, false
	}
	return o.OtfpShieldRespBodyBytes, true
}

// HasOtfpShieldRespBodyBytes returns a boolean if a field has been set.
func (o *Results) HasOtfpShieldRespBodyBytes() bool {
	if o != nil && o.OtfpShieldRespBodyBytes != nil {
		return true
	}

	return false
}

// SetOtfpShieldRespBodyBytes gets a reference to the given int32 and assigns it to the OtfpShieldRespBodyBytes field.
func (o *Results) SetOtfpShieldRespBodyBytes(v int32) {
	o.OtfpShieldRespBodyBytes = &v
}

// GetOtfpShieldRespHeaderBytes returns the OtfpShieldRespHeaderBytes field value if set, zero value otherwise.
func (o *Results) GetOtfpShieldRespHeaderBytes() int32 {
	if o == nil || o.OtfpShieldRespHeaderBytes == nil {
		var ret int32
		return ret
	}
	return *o.OtfpShieldRespHeaderBytes
}

// GetOtfpShieldRespHeaderBytesOk returns a tuple with the OtfpShieldRespHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetOtfpShieldRespHeaderBytesOk() (*int32, bool) {
	if o == nil || o.OtfpShieldRespHeaderBytes == nil {
		return nil, false
	}
	return o.OtfpShieldRespHeaderBytes, true
}

// HasOtfpShieldRespHeaderBytes returns a boolean if a field has been set.
func (o *Results) HasOtfpShieldRespHeaderBytes() bool {
	if o != nil && o.OtfpShieldRespHeaderBytes != nil {
		return true
	}

	return false
}

// SetOtfpShieldRespHeaderBytes gets a reference to the given int32 and assigns it to the OtfpShieldRespHeaderBytes field.
func (o *Results) SetOtfpShieldRespHeaderBytes(v int32) {
	o.OtfpShieldRespHeaderBytes = &v
}

// GetOtfpManifests returns the OtfpManifests field value if set, zero value otherwise.
func (o *Results) GetOtfpManifests() int32 {
	if o == nil || o.OtfpManifests == nil {
		var ret int32
		return ret
	}
	return *o.OtfpManifests
}

// GetOtfpManifestsOk returns a tuple with the OtfpManifests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetOtfpManifestsOk() (*int32, bool) {
	if o == nil || o.OtfpManifests == nil {
		return nil, false
	}
	return o.OtfpManifests, true
}

// HasOtfpManifests returns a boolean if a field has been set.
func (o *Results) HasOtfpManifests() bool {
	if o != nil && o.OtfpManifests != nil {
		return true
	}

	return false
}

// SetOtfpManifests gets a reference to the given int32 and assigns it to the OtfpManifests field.
func (o *Results) SetOtfpManifests(v int32) {
	o.OtfpManifests = &v
}

// GetOtfpDeliverTime returns the OtfpDeliverTime field value if set, zero value otherwise.
func (o *Results) GetOtfpDeliverTime() float32 {
	if o == nil || o.OtfpDeliverTime == nil {
		var ret float32
		return ret
	}
	return *o.OtfpDeliverTime
}

// GetOtfpDeliverTimeOk returns a tuple with the OtfpDeliverTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetOtfpDeliverTimeOk() (*float32, bool) {
	if o == nil || o.OtfpDeliverTime == nil {
		return nil, false
	}
	return o.OtfpDeliverTime, true
}

// HasOtfpDeliverTime returns a boolean if a field has been set.
func (o *Results) HasOtfpDeliverTime() bool {
	if o != nil && o.OtfpDeliverTime != nil {
		return true
	}

	return false
}

// SetOtfpDeliverTime gets a reference to the given float32 and assigns it to the OtfpDeliverTime field.
func (o *Results) SetOtfpDeliverTime(v float32) {
	o.OtfpDeliverTime = &v
}

// GetOtfpShieldTime returns the OtfpShieldTime field value if set, zero value otherwise.
func (o *Results) GetOtfpShieldTime() float32 {
	if o == nil || o.OtfpShieldTime == nil {
		var ret float32
		return ret
	}
	return *o.OtfpShieldTime
}

// GetOtfpShieldTimeOk returns a tuple with the OtfpShieldTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetOtfpShieldTimeOk() (*float32, bool) {
	if o == nil || o.OtfpShieldTime == nil {
		return nil, false
	}
	return o.OtfpShieldTime, true
}

// HasOtfpShieldTime returns a boolean if a field has been set.
func (o *Results) HasOtfpShieldTime() bool {
	if o != nil && o.OtfpShieldTime != nil {
		return true
	}

	return false
}

// SetOtfpShieldTime gets a reference to the given float32 and assigns it to the OtfpShieldTime field.
func (o *Results) SetOtfpShieldTime(v float32) {
	o.OtfpShieldTime = &v
}

// GetVideo returns the Video field value if set, zero value otherwise.
func (o *Results) GetVideo() int32 {
	if o == nil || o.Video == nil {
		var ret int32
		return ret
	}
	return *o.Video
}

// GetVideoOk returns a tuple with the Video field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetVideoOk() (*int32, bool) {
	if o == nil || o.Video == nil {
		return nil, false
	}
	return o.Video, true
}

// HasVideo returns a boolean if a field has been set.
func (o *Results) HasVideo() bool {
	if o != nil && o.Video != nil {
		return true
	}

	return false
}

// SetVideo gets a reference to the given int32 and assigns it to the Video field.
func (o *Results) SetVideo(v int32) {
	o.Video = &v
}

// GetPci returns the Pci field value if set, zero value otherwise.
func (o *Results) GetPci() int32 {
	if o == nil || o.Pci == nil {
		var ret int32
		return ret
	}
	return *o.Pci
}

// GetPciOk returns a tuple with the Pci field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetPciOk() (*int32, bool) {
	if o == nil || o.Pci == nil {
		return nil, false
	}
	return o.Pci, true
}

// HasPci returns a boolean if a field has been set.
func (o *Results) HasPci() bool {
	if o != nil && o.Pci != nil {
		return true
	}

	return false
}

// SetPci gets a reference to the given int32 and assigns it to the Pci field.
func (o *Results) SetPci(v int32) {
	o.Pci = &v
}

// GetLog returns the Log field value if set, zero value otherwise.
func (o *Results) GetLog() int32 {
	if o == nil || o.Log == nil {
		var ret int32
		return ret
	}
	return *o.Log
}

// GetLogOk returns a tuple with the Log field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetLogOk() (*int32, bool) {
	if o == nil || o.Log == nil {
		return nil, false
	}
	return o.Log, true
}

// HasLog returns a boolean if a field has been set.
func (o *Results) HasLog() bool {
	if o != nil && o.Log != nil {
		return true
	}

	return false
}

// SetLog gets a reference to the given int32 and assigns it to the Log field.
func (o *Results) SetLog(v int32) {
	o.Log = &v
}

// GetLogBytes returns the LogBytes field value if set, zero value otherwise.
func (o *Results) GetLogBytes() int32 {
	if o == nil || o.LogBytes == nil {
		var ret int32
		return ret
	}
	return *o.LogBytes
}

// GetLogBytesOk returns a tuple with the LogBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetLogBytesOk() (*int32, bool) {
	if o == nil || o.LogBytes == nil {
		return nil, false
	}
	return o.LogBytes, true
}

// HasLogBytes returns a boolean if a field has been set.
func (o *Results) HasLogBytes() bool {
	if o != nil && o.LogBytes != nil {
		return true
	}

	return false
}

// SetLogBytes gets a reference to the given int32 and assigns it to the LogBytes field.
func (o *Results) SetLogBytes(v int32) {
	o.LogBytes = &v
}

// GetHTTP2 returns the HTTP2 field value if set, zero value otherwise.
func (o *Results) GetHTTP2() int32 {
	if o == nil || o.HTTP2 == nil {
		var ret int32
		return ret
	}
	return *o.HTTP2
}

// GetHTTP2Ok returns a tuple with the HTTP2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetHTTP2Ok() (*int32, bool) {
	if o == nil || o.HTTP2 == nil {
		return nil, false
	}
	return o.HTTP2, true
}

// HasHTTP2 returns a boolean if a field has been set.
func (o *Results) HasHTTP2() bool {
	if o != nil && o.HTTP2 != nil {
		return true
	}

	return false
}

// SetHTTP2 gets a reference to the given int32 and assigns it to the HTTP2 field.
func (o *Results) SetHTTP2(v int32) {
	o.HTTP2 = &v
}

// GetHTTP3 returns the HTTP3 field value if set, zero value otherwise.
func (o *Results) GetHTTP3() int32 {
	if o == nil || o.HTTP3 == nil {
		var ret int32
		return ret
	}
	return *o.HTTP3
}

// GetHTTP3Ok returns a tuple with the HTTP3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetHTTP3Ok() (*int32, bool) {
	if o == nil || o.HTTP3 == nil {
		return nil, false
	}
	return o.HTTP3, true
}

// HasHTTP3 returns a boolean if a field has been set.
func (o *Results) HasHTTP3() bool {
	if o != nil && o.HTTP3 != nil {
		return true
	}

	return false
}

// SetHTTP3 gets a reference to the given int32 and assigns it to the HTTP3 field.
func (o *Results) SetHTTP3(v int32) {
	o.HTTP3 = &v
}

// GetWafLogged returns the WafLogged field value if set, zero value otherwise.
func (o *Results) GetWafLogged() int32 {
	if o == nil || o.WafLogged == nil {
		var ret int32
		return ret
	}
	return *o.WafLogged
}

// GetWafLoggedOk returns a tuple with the WafLogged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetWafLoggedOk() (*int32, bool) {
	if o == nil || o.WafLogged == nil {
		return nil, false
	}
	return o.WafLogged, true
}

// HasWafLogged returns a boolean if a field has been set.
func (o *Results) HasWafLogged() bool {
	if o != nil && o.WafLogged != nil {
		return true
	}

	return false
}

// SetWafLogged gets a reference to the given int32 and assigns it to the WafLogged field.
func (o *Results) SetWafLogged(v int32) {
	o.WafLogged = &v
}

// GetWafBlocked returns the WafBlocked field value if set, zero value otherwise.
func (o *Results) GetWafBlocked() int32 {
	if o == nil || o.WafBlocked == nil {
		var ret int32
		return ret
	}
	return *o.WafBlocked
}

// GetWafBlockedOk returns a tuple with the WafBlocked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetWafBlockedOk() (*int32, bool) {
	if o == nil || o.WafBlocked == nil {
		return nil, false
	}
	return o.WafBlocked, true
}

// HasWafBlocked returns a boolean if a field has been set.
func (o *Results) HasWafBlocked() bool {
	if o != nil && o.WafBlocked != nil {
		return true
	}

	return false
}

// SetWafBlocked gets a reference to the given int32 and assigns it to the WafBlocked field.
func (o *Results) SetWafBlocked(v int32) {
	o.WafBlocked = &v
}

// GetWafPassed returns the WafPassed field value if set, zero value otherwise.
func (o *Results) GetWafPassed() int32 {
	if o == nil || o.WafPassed == nil {
		var ret int32
		return ret
	}
	return *o.WafPassed
}

// GetWafPassedOk returns a tuple with the WafPassed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetWafPassedOk() (*int32, bool) {
	if o == nil || o.WafPassed == nil {
		return nil, false
	}
	return o.WafPassed, true
}

// HasWafPassed returns a boolean if a field has been set.
func (o *Results) HasWafPassed() bool {
	if o != nil && o.WafPassed != nil {
		return true
	}

	return false
}

// SetWafPassed gets a reference to the given int32 and assigns it to the WafPassed field.
func (o *Results) SetWafPassed(v int32) {
	o.WafPassed = &v
}

// GetAttackReqBodyBytes returns the AttackReqBodyBytes field value if set, zero value otherwise.
func (o *Results) GetAttackReqBodyBytes() int32 {
	if o == nil || o.AttackReqBodyBytes == nil {
		var ret int32
		return ret
	}
	return *o.AttackReqBodyBytes
}

// GetAttackReqBodyBytesOk returns a tuple with the AttackReqBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetAttackReqBodyBytesOk() (*int32, bool) {
	if o == nil || o.AttackReqBodyBytes == nil {
		return nil, false
	}
	return o.AttackReqBodyBytes, true
}

// HasAttackReqBodyBytes returns a boolean if a field has been set.
func (o *Results) HasAttackReqBodyBytes() bool {
	if o != nil && o.AttackReqBodyBytes != nil {
		return true
	}

	return false
}

// SetAttackReqBodyBytes gets a reference to the given int32 and assigns it to the AttackReqBodyBytes field.
func (o *Results) SetAttackReqBodyBytes(v int32) {
	o.AttackReqBodyBytes = &v
}

// GetAttackReqHeaderBytes returns the AttackReqHeaderBytes field value if set, zero value otherwise.
func (o *Results) GetAttackReqHeaderBytes() int32 {
	if o == nil || o.AttackReqHeaderBytes == nil {
		var ret int32
		return ret
	}
	return *o.AttackReqHeaderBytes
}

// GetAttackReqHeaderBytesOk returns a tuple with the AttackReqHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetAttackReqHeaderBytesOk() (*int32, bool) {
	if o == nil || o.AttackReqHeaderBytes == nil {
		return nil, false
	}
	return o.AttackReqHeaderBytes, true
}

// HasAttackReqHeaderBytes returns a boolean if a field has been set.
func (o *Results) HasAttackReqHeaderBytes() bool {
	if o != nil && o.AttackReqHeaderBytes != nil {
		return true
	}

	return false
}

// SetAttackReqHeaderBytes gets a reference to the given int32 and assigns it to the AttackReqHeaderBytes field.
func (o *Results) SetAttackReqHeaderBytes(v int32) {
	o.AttackReqHeaderBytes = &v
}

// GetAttackLoggedReqBodyBytes returns the AttackLoggedReqBodyBytes field value if set, zero value otherwise.
func (o *Results) GetAttackLoggedReqBodyBytes() int32 {
	if o == nil || o.AttackLoggedReqBodyBytes == nil {
		var ret int32
		return ret
	}
	return *o.AttackLoggedReqBodyBytes
}

// GetAttackLoggedReqBodyBytesOk returns a tuple with the AttackLoggedReqBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetAttackLoggedReqBodyBytesOk() (*int32, bool) {
	if o == nil || o.AttackLoggedReqBodyBytes == nil {
		return nil, false
	}
	return o.AttackLoggedReqBodyBytes, true
}

// HasAttackLoggedReqBodyBytes returns a boolean if a field has been set.
func (o *Results) HasAttackLoggedReqBodyBytes() bool {
	if o != nil && o.AttackLoggedReqBodyBytes != nil {
		return true
	}

	return false
}

// SetAttackLoggedReqBodyBytes gets a reference to the given int32 and assigns it to the AttackLoggedReqBodyBytes field.
func (o *Results) SetAttackLoggedReqBodyBytes(v int32) {
	o.AttackLoggedReqBodyBytes = &v
}

// GetAttackLoggedReqHeaderBytes returns the AttackLoggedReqHeaderBytes field value if set, zero value otherwise.
func (o *Results) GetAttackLoggedReqHeaderBytes() int32 {
	if o == nil || o.AttackLoggedReqHeaderBytes == nil {
		var ret int32
		return ret
	}
	return *o.AttackLoggedReqHeaderBytes
}

// GetAttackLoggedReqHeaderBytesOk returns a tuple with the AttackLoggedReqHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetAttackLoggedReqHeaderBytesOk() (*int32, bool) {
	if o == nil || o.AttackLoggedReqHeaderBytes == nil {
		return nil, false
	}
	return o.AttackLoggedReqHeaderBytes, true
}

// HasAttackLoggedReqHeaderBytes returns a boolean if a field has been set.
func (o *Results) HasAttackLoggedReqHeaderBytes() bool {
	if o != nil && o.AttackLoggedReqHeaderBytes != nil {
		return true
	}

	return false
}

// SetAttackLoggedReqHeaderBytes gets a reference to the given int32 and assigns it to the AttackLoggedReqHeaderBytes field.
func (o *Results) SetAttackLoggedReqHeaderBytes(v int32) {
	o.AttackLoggedReqHeaderBytes = &v
}

// GetAttackBlockedReqBodyBytes returns the AttackBlockedReqBodyBytes field value if set, zero value otherwise.
func (o *Results) GetAttackBlockedReqBodyBytes() int32 {
	if o == nil || o.AttackBlockedReqBodyBytes == nil {
		var ret int32
		return ret
	}
	return *o.AttackBlockedReqBodyBytes
}

// GetAttackBlockedReqBodyBytesOk returns a tuple with the AttackBlockedReqBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetAttackBlockedReqBodyBytesOk() (*int32, bool) {
	if o == nil || o.AttackBlockedReqBodyBytes == nil {
		return nil, false
	}
	return o.AttackBlockedReqBodyBytes, true
}

// HasAttackBlockedReqBodyBytes returns a boolean if a field has been set.
func (o *Results) HasAttackBlockedReqBodyBytes() bool {
	if o != nil && o.AttackBlockedReqBodyBytes != nil {
		return true
	}

	return false
}

// SetAttackBlockedReqBodyBytes gets a reference to the given int32 and assigns it to the AttackBlockedReqBodyBytes field.
func (o *Results) SetAttackBlockedReqBodyBytes(v int32) {
	o.AttackBlockedReqBodyBytes = &v
}

// GetAttackBlockedReqHeaderBytes returns the AttackBlockedReqHeaderBytes field value if set, zero value otherwise.
func (o *Results) GetAttackBlockedReqHeaderBytes() int32 {
	if o == nil || o.AttackBlockedReqHeaderBytes == nil {
		var ret int32
		return ret
	}
	return *o.AttackBlockedReqHeaderBytes
}

// GetAttackBlockedReqHeaderBytesOk returns a tuple with the AttackBlockedReqHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetAttackBlockedReqHeaderBytesOk() (*int32, bool) {
	if o == nil || o.AttackBlockedReqHeaderBytes == nil {
		return nil, false
	}
	return o.AttackBlockedReqHeaderBytes, true
}

// HasAttackBlockedReqHeaderBytes returns a boolean if a field has been set.
func (o *Results) HasAttackBlockedReqHeaderBytes() bool {
	if o != nil && o.AttackBlockedReqHeaderBytes != nil {
		return true
	}

	return false
}

// SetAttackBlockedReqHeaderBytes gets a reference to the given int32 and assigns it to the AttackBlockedReqHeaderBytes field.
func (o *Results) SetAttackBlockedReqHeaderBytes(v int32) {
	o.AttackBlockedReqHeaderBytes = &v
}

// GetAttackPassedReqBodyBytes returns the AttackPassedReqBodyBytes field value if set, zero value otherwise.
func (o *Results) GetAttackPassedReqBodyBytes() int32 {
	if o == nil || o.AttackPassedReqBodyBytes == nil {
		var ret int32
		return ret
	}
	return *o.AttackPassedReqBodyBytes
}

// GetAttackPassedReqBodyBytesOk returns a tuple with the AttackPassedReqBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetAttackPassedReqBodyBytesOk() (*int32, bool) {
	if o == nil || o.AttackPassedReqBodyBytes == nil {
		return nil, false
	}
	return o.AttackPassedReqBodyBytes, true
}

// HasAttackPassedReqBodyBytes returns a boolean if a field has been set.
func (o *Results) HasAttackPassedReqBodyBytes() bool {
	if o != nil && o.AttackPassedReqBodyBytes != nil {
		return true
	}

	return false
}

// SetAttackPassedReqBodyBytes gets a reference to the given int32 and assigns it to the AttackPassedReqBodyBytes field.
func (o *Results) SetAttackPassedReqBodyBytes(v int32) {
	o.AttackPassedReqBodyBytes = &v
}

// GetAttackPassedReqHeaderBytes returns the AttackPassedReqHeaderBytes field value if set, zero value otherwise.
func (o *Results) GetAttackPassedReqHeaderBytes() int32 {
	if o == nil || o.AttackPassedReqHeaderBytes == nil {
		var ret int32
		return ret
	}
	return *o.AttackPassedReqHeaderBytes
}

// GetAttackPassedReqHeaderBytesOk returns a tuple with the AttackPassedReqHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetAttackPassedReqHeaderBytesOk() (*int32, bool) {
	if o == nil || o.AttackPassedReqHeaderBytes == nil {
		return nil, false
	}
	return o.AttackPassedReqHeaderBytes, true
}

// HasAttackPassedReqHeaderBytes returns a boolean if a field has been set.
func (o *Results) HasAttackPassedReqHeaderBytes() bool {
	if o != nil && o.AttackPassedReqHeaderBytes != nil {
		return true
	}

	return false
}

// SetAttackPassedReqHeaderBytes gets a reference to the given int32 and assigns it to the AttackPassedReqHeaderBytes field.
func (o *Results) SetAttackPassedReqHeaderBytes(v int32) {
	o.AttackPassedReqHeaderBytes = &v
}

// GetAttackRespSynthBytes returns the AttackRespSynthBytes field value if set, zero value otherwise.
func (o *Results) GetAttackRespSynthBytes() int32 {
	if o == nil || o.AttackRespSynthBytes == nil {
		var ret int32
		return ret
	}
	return *o.AttackRespSynthBytes
}

// GetAttackRespSynthBytesOk returns a tuple with the AttackRespSynthBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetAttackRespSynthBytesOk() (*int32, bool) {
	if o == nil || o.AttackRespSynthBytes == nil {
		return nil, false
	}
	return o.AttackRespSynthBytes, true
}

// HasAttackRespSynthBytes returns a boolean if a field has been set.
func (o *Results) HasAttackRespSynthBytes() bool {
	if o != nil && o.AttackRespSynthBytes != nil {
		return true
	}

	return false
}

// SetAttackRespSynthBytes gets a reference to the given int32 and assigns it to the AttackRespSynthBytes field.
func (o *Results) SetAttackRespSynthBytes(v int32) {
	o.AttackRespSynthBytes = &v
}

// GetImgopto returns the Imgopto field value if set, zero value otherwise.
func (o *Results) GetImgopto() int32 {
	if o == nil || o.Imgopto == nil {
		var ret int32
		return ret
	}
	return *o.Imgopto
}

// GetImgoptoOk returns a tuple with the Imgopto field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetImgoptoOk() (*int32, bool) {
	if o == nil || o.Imgopto == nil {
		return nil, false
	}
	return o.Imgopto, true
}

// HasImgopto returns a boolean if a field has been set.
func (o *Results) HasImgopto() bool {
	if o != nil && o.Imgopto != nil {
		return true
	}

	return false
}

// SetImgopto gets a reference to the given int32 and assigns it to the Imgopto field.
func (o *Results) SetImgopto(v int32) {
	o.Imgopto = &v
}

// GetImgoptoRespBodyBytes returns the ImgoptoRespBodyBytes field value if set, zero value otherwise.
func (o *Results) GetImgoptoRespBodyBytes() int32 {
	if o == nil || o.ImgoptoRespBodyBytes == nil {
		var ret int32
		return ret
	}
	return *o.ImgoptoRespBodyBytes
}

// GetImgoptoRespBodyBytesOk returns a tuple with the ImgoptoRespBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetImgoptoRespBodyBytesOk() (*int32, bool) {
	if o == nil || o.ImgoptoRespBodyBytes == nil {
		return nil, false
	}
	return o.ImgoptoRespBodyBytes, true
}

// HasImgoptoRespBodyBytes returns a boolean if a field has been set.
func (o *Results) HasImgoptoRespBodyBytes() bool {
	if o != nil && o.ImgoptoRespBodyBytes != nil {
		return true
	}

	return false
}

// SetImgoptoRespBodyBytes gets a reference to the given int32 and assigns it to the ImgoptoRespBodyBytes field.
func (o *Results) SetImgoptoRespBodyBytes(v int32) {
	o.ImgoptoRespBodyBytes = &v
}

// GetImgoptoRespHeaderBytes returns the ImgoptoRespHeaderBytes field value if set, zero value otherwise.
func (o *Results) GetImgoptoRespHeaderBytes() int32 {
	if o == nil || o.ImgoptoRespHeaderBytes == nil {
		var ret int32
		return ret
	}
	return *o.ImgoptoRespHeaderBytes
}

// GetImgoptoRespHeaderBytesOk returns a tuple with the ImgoptoRespHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetImgoptoRespHeaderBytesOk() (*int32, bool) {
	if o == nil || o.ImgoptoRespHeaderBytes == nil {
		return nil, false
	}
	return o.ImgoptoRespHeaderBytes, true
}

// HasImgoptoRespHeaderBytes returns a boolean if a field has been set.
func (o *Results) HasImgoptoRespHeaderBytes() bool {
	if o != nil && o.ImgoptoRespHeaderBytes != nil {
		return true
	}

	return false
}

// SetImgoptoRespHeaderBytes gets a reference to the given int32 and assigns it to the ImgoptoRespHeaderBytes field.
func (o *Results) SetImgoptoRespHeaderBytes(v int32) {
	o.ImgoptoRespHeaderBytes = &v
}

// GetImgoptoShield returns the ImgoptoShield field value if set, zero value otherwise.
func (o *Results) GetImgoptoShield() int32 {
	if o == nil || o.ImgoptoShield == nil {
		var ret int32
		return ret
	}
	return *o.ImgoptoShield
}

// GetImgoptoShieldOk returns a tuple with the ImgoptoShield field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetImgoptoShieldOk() (*int32, bool) {
	if o == nil || o.ImgoptoShield == nil {
		return nil, false
	}
	return o.ImgoptoShield, true
}

// HasImgoptoShield returns a boolean if a field has been set.
func (o *Results) HasImgoptoShield() bool {
	if o != nil && o.ImgoptoShield != nil {
		return true
	}

	return false
}

// SetImgoptoShield gets a reference to the given int32 and assigns it to the ImgoptoShield field.
func (o *Results) SetImgoptoShield(v int32) {
	o.ImgoptoShield = &v
}

// GetImgoptoShieldRespBodyBytes returns the ImgoptoShieldRespBodyBytes field value if set, zero value otherwise.
func (o *Results) GetImgoptoShieldRespBodyBytes() int32 {
	if o == nil || o.ImgoptoShieldRespBodyBytes == nil {
		var ret int32
		return ret
	}
	return *o.ImgoptoShieldRespBodyBytes
}

// GetImgoptoShieldRespBodyBytesOk returns a tuple with the ImgoptoShieldRespBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetImgoptoShieldRespBodyBytesOk() (*int32, bool) {
	if o == nil || o.ImgoptoShieldRespBodyBytes == nil {
		return nil, false
	}
	return o.ImgoptoShieldRespBodyBytes, true
}

// HasImgoptoShieldRespBodyBytes returns a boolean if a field has been set.
func (o *Results) HasImgoptoShieldRespBodyBytes() bool {
	if o != nil && o.ImgoptoShieldRespBodyBytes != nil {
		return true
	}

	return false
}

// SetImgoptoShieldRespBodyBytes gets a reference to the given int32 and assigns it to the ImgoptoShieldRespBodyBytes field.
func (o *Results) SetImgoptoShieldRespBodyBytes(v int32) {
	o.ImgoptoShieldRespBodyBytes = &v
}

// GetImgoptoShieldRespHeaderBytes returns the ImgoptoShieldRespHeaderBytes field value if set, zero value otherwise.
func (o *Results) GetImgoptoShieldRespHeaderBytes() int32 {
	if o == nil || o.ImgoptoShieldRespHeaderBytes == nil {
		var ret int32
		return ret
	}
	return *o.ImgoptoShieldRespHeaderBytes
}

// GetImgoptoShieldRespHeaderBytesOk returns a tuple with the ImgoptoShieldRespHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetImgoptoShieldRespHeaderBytesOk() (*int32, bool) {
	if o == nil || o.ImgoptoShieldRespHeaderBytes == nil {
		return nil, false
	}
	return o.ImgoptoShieldRespHeaderBytes, true
}

// HasImgoptoShieldRespHeaderBytes returns a boolean if a field has been set.
func (o *Results) HasImgoptoShieldRespHeaderBytes() bool {
	if o != nil && o.ImgoptoShieldRespHeaderBytes != nil {
		return true
	}

	return false
}

// SetImgoptoShieldRespHeaderBytes gets a reference to the given int32 and assigns it to the ImgoptoShieldRespHeaderBytes field.
func (o *Results) SetImgoptoShieldRespHeaderBytes(v int32) {
	o.ImgoptoShieldRespHeaderBytes = &v
}

// GetImgoptoTransforms returns the ImgoptoTransforms field value if set, zero value otherwise.
func (o *Results) GetImgoptoTransforms() int32 {
	if o == nil || o.ImgoptoTransforms == nil {
		var ret int32
		return ret
	}
	return *o.ImgoptoTransforms
}

// GetImgoptoTransformsOk returns a tuple with the ImgoptoTransforms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetImgoptoTransformsOk() (*int32, bool) {
	if o == nil || o.ImgoptoTransforms == nil {
		return nil, false
	}
	return o.ImgoptoTransforms, true
}

// HasImgoptoTransforms returns a boolean if a field has been set.
func (o *Results) HasImgoptoTransforms() bool {
	if o != nil && o.ImgoptoTransforms != nil {
		return true
	}

	return false
}

// SetImgoptoTransforms gets a reference to the given int32 and assigns it to the ImgoptoTransforms field.
func (o *Results) SetImgoptoTransforms(v int32) {
	o.ImgoptoTransforms = &v
}

// GetImgvideo returns the Imgvideo field value if set, zero value otherwise.
func (o *Results) GetImgvideo() int32 {
	if o == nil || o.Imgvideo == nil {
		var ret int32
		return ret
	}
	return *o.Imgvideo
}

// GetImgvideoOk returns a tuple with the Imgvideo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetImgvideoOk() (*int32, bool) {
	if o == nil || o.Imgvideo == nil {
		return nil, false
	}
	return o.Imgvideo, true
}

// HasImgvideo returns a boolean if a field has been set.
func (o *Results) HasImgvideo() bool {
	if o != nil && o.Imgvideo != nil {
		return true
	}

	return false
}

// SetImgvideo gets a reference to the given int32 and assigns it to the Imgvideo field.
func (o *Results) SetImgvideo(v int32) {
	o.Imgvideo = &v
}

// GetImgvideoFrames returns the ImgvideoFrames field value if set, zero value otherwise.
func (o *Results) GetImgvideoFrames() int32 {
	if o == nil || o.ImgvideoFrames == nil {
		var ret int32
		return ret
	}
	return *o.ImgvideoFrames
}

// GetImgvideoFramesOk returns a tuple with the ImgvideoFrames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetImgvideoFramesOk() (*int32, bool) {
	if o == nil || o.ImgvideoFrames == nil {
		return nil, false
	}
	return o.ImgvideoFrames, true
}

// HasImgvideoFrames returns a boolean if a field has been set.
func (o *Results) HasImgvideoFrames() bool {
	if o != nil && o.ImgvideoFrames != nil {
		return true
	}

	return false
}

// SetImgvideoFrames gets a reference to the given int32 and assigns it to the ImgvideoFrames field.
func (o *Results) SetImgvideoFrames(v int32) {
	o.ImgvideoFrames = &v
}

// GetImgvideoRespHeaderBytes returns the ImgvideoRespHeaderBytes field value if set, zero value otherwise.
func (o *Results) GetImgvideoRespHeaderBytes() int32 {
	if o == nil || o.ImgvideoRespHeaderBytes == nil {
		var ret int32
		return ret
	}
	return *o.ImgvideoRespHeaderBytes
}

// GetImgvideoRespHeaderBytesOk returns a tuple with the ImgvideoRespHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetImgvideoRespHeaderBytesOk() (*int32, bool) {
	if o == nil || o.ImgvideoRespHeaderBytes == nil {
		return nil, false
	}
	return o.ImgvideoRespHeaderBytes, true
}

// HasImgvideoRespHeaderBytes returns a boolean if a field has been set.
func (o *Results) HasImgvideoRespHeaderBytes() bool {
	if o != nil && o.ImgvideoRespHeaderBytes != nil {
		return true
	}

	return false
}

// SetImgvideoRespHeaderBytes gets a reference to the given int32 and assigns it to the ImgvideoRespHeaderBytes field.
func (o *Results) SetImgvideoRespHeaderBytes(v int32) {
	o.ImgvideoRespHeaderBytes = &v
}

// GetImgvideoRespBodyBytes returns the ImgvideoRespBodyBytes field value if set, zero value otherwise.
func (o *Results) GetImgvideoRespBodyBytes() int32 {
	if o == nil || o.ImgvideoRespBodyBytes == nil {
		var ret int32
		return ret
	}
	return *o.ImgvideoRespBodyBytes
}

// GetImgvideoRespBodyBytesOk returns a tuple with the ImgvideoRespBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetImgvideoRespBodyBytesOk() (*int32, bool) {
	if o == nil || o.ImgvideoRespBodyBytes == nil {
		return nil, false
	}
	return o.ImgvideoRespBodyBytes, true
}

// HasImgvideoRespBodyBytes returns a boolean if a field has been set.
func (o *Results) HasImgvideoRespBodyBytes() bool {
	if o != nil && o.ImgvideoRespBodyBytes != nil {
		return true
	}

	return false
}

// SetImgvideoRespBodyBytes gets a reference to the given int32 and assigns it to the ImgvideoRespBodyBytes field.
func (o *Results) SetImgvideoRespBodyBytes(v int32) {
	o.ImgvideoRespBodyBytes = &v
}

// GetImgvideoShieldRespHeaderBytes returns the ImgvideoShieldRespHeaderBytes field value if set, zero value otherwise.
func (o *Results) GetImgvideoShieldRespHeaderBytes() int32 {
	if o == nil || o.ImgvideoShieldRespHeaderBytes == nil {
		var ret int32
		return ret
	}
	return *o.ImgvideoShieldRespHeaderBytes
}

// GetImgvideoShieldRespHeaderBytesOk returns a tuple with the ImgvideoShieldRespHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetImgvideoShieldRespHeaderBytesOk() (*int32, bool) {
	if o == nil || o.ImgvideoShieldRespHeaderBytes == nil {
		return nil, false
	}
	return o.ImgvideoShieldRespHeaderBytes, true
}

// HasImgvideoShieldRespHeaderBytes returns a boolean if a field has been set.
func (o *Results) HasImgvideoShieldRespHeaderBytes() bool {
	if o != nil && o.ImgvideoShieldRespHeaderBytes != nil {
		return true
	}

	return false
}

// SetImgvideoShieldRespHeaderBytes gets a reference to the given int32 and assigns it to the ImgvideoShieldRespHeaderBytes field.
func (o *Results) SetImgvideoShieldRespHeaderBytes(v int32) {
	o.ImgvideoShieldRespHeaderBytes = &v
}

// GetImgvideoShieldRespBodyBytes returns the ImgvideoShieldRespBodyBytes field value if set, zero value otherwise.
func (o *Results) GetImgvideoShieldRespBodyBytes() int32 {
	if o == nil || o.ImgvideoShieldRespBodyBytes == nil {
		var ret int32
		return ret
	}
	return *o.ImgvideoShieldRespBodyBytes
}

// GetImgvideoShieldRespBodyBytesOk returns a tuple with the ImgvideoShieldRespBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetImgvideoShieldRespBodyBytesOk() (*int32, bool) {
	if o == nil || o.ImgvideoShieldRespBodyBytes == nil {
		return nil, false
	}
	return o.ImgvideoShieldRespBodyBytes, true
}

// HasImgvideoShieldRespBodyBytes returns a boolean if a field has been set.
func (o *Results) HasImgvideoShieldRespBodyBytes() bool {
	if o != nil && o.ImgvideoShieldRespBodyBytes != nil {
		return true
	}

	return false
}

// SetImgvideoShieldRespBodyBytes gets a reference to the given int32 and assigns it to the ImgvideoShieldRespBodyBytes field.
func (o *Results) SetImgvideoShieldRespBodyBytes(v int32) {
	o.ImgvideoShieldRespBodyBytes = &v
}

// GetImgvideoShield returns the ImgvideoShield field value if set, zero value otherwise.
func (o *Results) GetImgvideoShield() int32 {
	if o == nil || o.ImgvideoShield == nil {
		var ret int32
		return ret
	}
	return *o.ImgvideoShield
}

// GetImgvideoShieldOk returns a tuple with the ImgvideoShield field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetImgvideoShieldOk() (*int32, bool) {
	if o == nil || o.ImgvideoShield == nil {
		return nil, false
	}
	return o.ImgvideoShield, true
}

// HasImgvideoShield returns a boolean if a field has been set.
func (o *Results) HasImgvideoShield() bool {
	if o != nil && o.ImgvideoShield != nil {
		return true
	}

	return false
}

// SetImgvideoShield gets a reference to the given int32 and assigns it to the ImgvideoShield field.
func (o *Results) SetImgvideoShield(v int32) {
	o.ImgvideoShield = &v
}

// GetImgvideoShieldFrames returns the ImgvideoShieldFrames field value if set, zero value otherwise.
func (o *Results) GetImgvideoShieldFrames() int32 {
	if o == nil || o.ImgvideoShieldFrames == nil {
		var ret int32
		return ret
	}
	return *o.ImgvideoShieldFrames
}

// GetImgvideoShieldFramesOk returns a tuple with the ImgvideoShieldFrames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetImgvideoShieldFramesOk() (*int32, bool) {
	if o == nil || o.ImgvideoShieldFrames == nil {
		return nil, false
	}
	return o.ImgvideoShieldFrames, true
}

// HasImgvideoShieldFrames returns a boolean if a field has been set.
func (o *Results) HasImgvideoShieldFrames() bool {
	if o != nil && o.ImgvideoShieldFrames != nil {
		return true
	}

	return false
}

// SetImgvideoShieldFrames gets a reference to the given int32 and assigns it to the ImgvideoShieldFrames field.
func (o *Results) SetImgvideoShieldFrames(v int32) {
	o.ImgvideoShieldFrames = &v
}

// GetStatus200 returns the Status200 field value if set, zero value otherwise.
func (o *Results) GetStatus200() int32 {
	if o == nil || o.Status200 == nil {
		var ret int32
		return ret
	}
	return *o.Status200
}

// GetStatus200Ok returns a tuple with the Status200 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetStatus200Ok() (*int32, bool) {
	if o == nil || o.Status200 == nil {
		return nil, false
	}
	return o.Status200, true
}

// HasStatus200 returns a boolean if a field has been set.
func (o *Results) HasStatus200() bool {
	if o != nil && o.Status200 != nil {
		return true
	}

	return false
}

// SetStatus200 gets a reference to the given int32 and assigns it to the Status200 field.
func (o *Results) SetStatus200(v int32) {
	o.Status200 = &v
}

// GetStatus204 returns the Status204 field value if set, zero value otherwise.
func (o *Results) GetStatus204() int32 {
	if o == nil || o.Status204 == nil {
		var ret int32
		return ret
	}
	return *o.Status204
}

// GetStatus204Ok returns a tuple with the Status204 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetStatus204Ok() (*int32, bool) {
	if o == nil || o.Status204 == nil {
		return nil, false
	}
	return o.Status204, true
}

// HasStatus204 returns a boolean if a field has been set.
func (o *Results) HasStatus204() bool {
	if o != nil && o.Status204 != nil {
		return true
	}

	return false
}

// SetStatus204 gets a reference to the given int32 and assigns it to the Status204 field.
func (o *Results) SetStatus204(v int32) {
	o.Status204 = &v
}

// GetStatus206 returns the Status206 field value if set, zero value otherwise.
func (o *Results) GetStatus206() int32 {
	if o == nil || o.Status206 == nil {
		var ret int32
		return ret
	}
	return *o.Status206
}

// GetStatus206Ok returns a tuple with the Status206 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetStatus206Ok() (*int32, bool) {
	if o == nil || o.Status206 == nil {
		return nil, false
	}
	return o.Status206, true
}

// HasStatus206 returns a boolean if a field has been set.
func (o *Results) HasStatus206() bool {
	if o != nil && o.Status206 != nil {
		return true
	}

	return false
}

// SetStatus206 gets a reference to the given int32 and assigns it to the Status206 field.
func (o *Results) SetStatus206(v int32) {
	o.Status206 = &v
}

// GetStatus301 returns the Status301 field value if set, zero value otherwise.
func (o *Results) GetStatus301() int32 {
	if o == nil || o.Status301 == nil {
		var ret int32
		return ret
	}
	return *o.Status301
}

// GetStatus301Ok returns a tuple with the Status301 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetStatus301Ok() (*int32, bool) {
	if o == nil || o.Status301 == nil {
		return nil, false
	}
	return o.Status301, true
}

// HasStatus301 returns a boolean if a field has been set.
func (o *Results) HasStatus301() bool {
	if o != nil && o.Status301 != nil {
		return true
	}

	return false
}

// SetStatus301 gets a reference to the given int32 and assigns it to the Status301 field.
func (o *Results) SetStatus301(v int32) {
	o.Status301 = &v
}

// GetStatus302 returns the Status302 field value if set, zero value otherwise.
func (o *Results) GetStatus302() int32 {
	if o == nil || o.Status302 == nil {
		var ret int32
		return ret
	}
	return *o.Status302
}

// GetStatus302Ok returns a tuple with the Status302 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetStatus302Ok() (*int32, bool) {
	if o == nil || o.Status302 == nil {
		return nil, false
	}
	return o.Status302, true
}

// HasStatus302 returns a boolean if a field has been set.
func (o *Results) HasStatus302() bool {
	if o != nil && o.Status302 != nil {
		return true
	}

	return false
}

// SetStatus302 gets a reference to the given int32 and assigns it to the Status302 field.
func (o *Results) SetStatus302(v int32) {
	o.Status302 = &v
}

// GetStatus304 returns the Status304 field value if set, zero value otherwise.
func (o *Results) GetStatus304() int32 {
	if o == nil || o.Status304 == nil {
		var ret int32
		return ret
	}
	return *o.Status304
}

// GetStatus304Ok returns a tuple with the Status304 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetStatus304Ok() (*int32, bool) {
	if o == nil || o.Status304 == nil {
		return nil, false
	}
	return o.Status304, true
}

// HasStatus304 returns a boolean if a field has been set.
func (o *Results) HasStatus304() bool {
	if o != nil && o.Status304 != nil {
		return true
	}

	return false
}

// SetStatus304 gets a reference to the given int32 and assigns it to the Status304 field.
func (o *Results) SetStatus304(v int32) {
	o.Status304 = &v
}

// GetStatus400 returns the Status400 field value if set, zero value otherwise.
func (o *Results) GetStatus400() int32 {
	if o == nil || o.Status400 == nil {
		var ret int32
		return ret
	}
	return *o.Status400
}

// GetStatus400Ok returns a tuple with the Status400 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetStatus400Ok() (*int32, bool) {
	if o == nil || o.Status400 == nil {
		return nil, false
	}
	return o.Status400, true
}

// HasStatus400 returns a boolean if a field has been set.
func (o *Results) HasStatus400() bool {
	if o != nil && o.Status400 != nil {
		return true
	}

	return false
}

// SetStatus400 gets a reference to the given int32 and assigns it to the Status400 field.
func (o *Results) SetStatus400(v int32) {
	o.Status400 = &v
}

// GetStatus401 returns the Status401 field value if set, zero value otherwise.
func (o *Results) GetStatus401() int32 {
	if o == nil || o.Status401 == nil {
		var ret int32
		return ret
	}
	return *o.Status401
}

// GetStatus401Ok returns a tuple with the Status401 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetStatus401Ok() (*int32, bool) {
	if o == nil || o.Status401 == nil {
		return nil, false
	}
	return o.Status401, true
}

// HasStatus401 returns a boolean if a field has been set.
func (o *Results) HasStatus401() bool {
	if o != nil && o.Status401 != nil {
		return true
	}

	return false
}

// SetStatus401 gets a reference to the given int32 and assigns it to the Status401 field.
func (o *Results) SetStatus401(v int32) {
	o.Status401 = &v
}

// GetStatus403 returns the Status403 field value if set, zero value otherwise.
func (o *Results) GetStatus403() int32 {
	if o == nil || o.Status403 == nil {
		var ret int32
		return ret
	}
	return *o.Status403
}

// GetStatus403Ok returns a tuple with the Status403 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetStatus403Ok() (*int32, bool) {
	if o == nil || o.Status403 == nil {
		return nil, false
	}
	return o.Status403, true
}

// HasStatus403 returns a boolean if a field has been set.
func (o *Results) HasStatus403() bool {
	if o != nil && o.Status403 != nil {
		return true
	}

	return false
}

// SetStatus403 gets a reference to the given int32 and assigns it to the Status403 field.
func (o *Results) SetStatus403(v int32) {
	o.Status403 = &v
}

// GetStatus404 returns the Status404 field value if set, zero value otherwise.
func (o *Results) GetStatus404() int32 {
	if o == nil || o.Status404 == nil {
		var ret int32
		return ret
	}
	return *o.Status404
}

// GetStatus404Ok returns a tuple with the Status404 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetStatus404Ok() (*int32, bool) {
	if o == nil || o.Status404 == nil {
		return nil, false
	}
	return o.Status404, true
}

// HasStatus404 returns a boolean if a field has been set.
func (o *Results) HasStatus404() bool {
	if o != nil && o.Status404 != nil {
		return true
	}

	return false
}

// SetStatus404 gets a reference to the given int32 and assigns it to the Status404 field.
func (o *Results) SetStatus404(v int32) {
	o.Status404 = &v
}

// GetStatus406 returns the Status406 field value if set, zero value otherwise.
func (o *Results) GetStatus406() int32 {
	if o == nil || o.Status406 == nil {
		var ret int32
		return ret
	}
	return *o.Status406
}

// GetStatus406Ok returns a tuple with the Status406 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetStatus406Ok() (*int32, bool) {
	if o == nil || o.Status406 == nil {
		return nil, false
	}
	return o.Status406, true
}

// HasStatus406 returns a boolean if a field has been set.
func (o *Results) HasStatus406() bool {
	if o != nil && o.Status406 != nil {
		return true
	}

	return false
}

// SetStatus406 gets a reference to the given int32 and assigns it to the Status406 field.
func (o *Results) SetStatus406(v int32) {
	o.Status406 = &v
}

// GetStatus416 returns the Status416 field value if set, zero value otherwise.
func (o *Results) GetStatus416() int32 {
	if o == nil || o.Status416 == nil {
		var ret int32
		return ret
	}
	return *o.Status416
}

// GetStatus416Ok returns a tuple with the Status416 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetStatus416Ok() (*int32, bool) {
	if o == nil || o.Status416 == nil {
		return nil, false
	}
	return o.Status416, true
}

// HasStatus416 returns a boolean if a field has been set.
func (o *Results) HasStatus416() bool {
	if o != nil && o.Status416 != nil {
		return true
	}

	return false
}

// SetStatus416 gets a reference to the given int32 and assigns it to the Status416 field.
func (o *Results) SetStatus416(v int32) {
	o.Status416 = &v
}

// GetStatus429 returns the Status429 field value if set, zero value otherwise.
func (o *Results) GetStatus429() int32 {
	if o == nil || o.Status429 == nil {
		var ret int32
		return ret
	}
	return *o.Status429
}

// GetStatus429Ok returns a tuple with the Status429 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetStatus429Ok() (*int32, bool) {
	if o == nil || o.Status429 == nil {
		return nil, false
	}
	return o.Status429, true
}

// HasStatus429 returns a boolean if a field has been set.
func (o *Results) HasStatus429() bool {
	if o != nil && o.Status429 != nil {
		return true
	}

	return false
}

// SetStatus429 gets a reference to the given int32 and assigns it to the Status429 field.
func (o *Results) SetStatus429(v int32) {
	o.Status429 = &v
}

// GetStatus500 returns the Status500 field value if set, zero value otherwise.
func (o *Results) GetStatus500() int32 {
	if o == nil || o.Status500 == nil {
		var ret int32
		return ret
	}
	return *o.Status500
}

// GetStatus500Ok returns a tuple with the Status500 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetStatus500Ok() (*int32, bool) {
	if o == nil || o.Status500 == nil {
		return nil, false
	}
	return o.Status500, true
}

// HasStatus500 returns a boolean if a field has been set.
func (o *Results) HasStatus500() bool {
	if o != nil && o.Status500 != nil {
		return true
	}

	return false
}

// SetStatus500 gets a reference to the given int32 and assigns it to the Status500 field.
func (o *Results) SetStatus500(v int32) {
	o.Status500 = &v
}

// GetStatus501 returns the Status501 field value if set, zero value otherwise.
func (o *Results) GetStatus501() int32 {
	if o == nil || o.Status501 == nil {
		var ret int32
		return ret
	}
	return *o.Status501
}

// GetStatus501Ok returns a tuple with the Status501 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetStatus501Ok() (*int32, bool) {
	if o == nil || o.Status501 == nil {
		return nil, false
	}
	return o.Status501, true
}

// HasStatus501 returns a boolean if a field has been set.
func (o *Results) HasStatus501() bool {
	if o != nil && o.Status501 != nil {
		return true
	}

	return false
}

// SetStatus501 gets a reference to the given int32 and assigns it to the Status501 field.
func (o *Results) SetStatus501(v int32) {
	o.Status501 = &v
}

// GetStatus502 returns the Status502 field value if set, zero value otherwise.
func (o *Results) GetStatus502() int32 {
	if o == nil || o.Status502 == nil {
		var ret int32
		return ret
	}
	return *o.Status502
}

// GetStatus502Ok returns a tuple with the Status502 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetStatus502Ok() (*int32, bool) {
	if o == nil || o.Status502 == nil {
		return nil, false
	}
	return o.Status502, true
}

// HasStatus502 returns a boolean if a field has been set.
func (o *Results) HasStatus502() bool {
	if o != nil && o.Status502 != nil {
		return true
	}

	return false
}

// SetStatus502 gets a reference to the given int32 and assigns it to the Status502 field.
func (o *Results) SetStatus502(v int32) {
	o.Status502 = &v
}

// GetStatus503 returns the Status503 field value if set, zero value otherwise.
func (o *Results) GetStatus503() int32 {
	if o == nil || o.Status503 == nil {
		var ret int32
		return ret
	}
	return *o.Status503
}

// GetStatus503Ok returns a tuple with the Status503 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetStatus503Ok() (*int32, bool) {
	if o == nil || o.Status503 == nil {
		return nil, false
	}
	return o.Status503, true
}

// HasStatus503 returns a boolean if a field has been set.
func (o *Results) HasStatus503() bool {
	if o != nil && o.Status503 != nil {
		return true
	}

	return false
}

// SetStatus503 gets a reference to the given int32 and assigns it to the Status503 field.
func (o *Results) SetStatus503(v int32) {
	o.Status503 = &v
}

// GetStatus504 returns the Status504 field value if set, zero value otherwise.
func (o *Results) GetStatus504() int32 {
	if o == nil || o.Status504 == nil {
		var ret int32
		return ret
	}
	return *o.Status504
}

// GetStatus504Ok returns a tuple with the Status504 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetStatus504Ok() (*int32, bool) {
	if o == nil || o.Status504 == nil {
		return nil, false
	}
	return o.Status504, true
}

// HasStatus504 returns a boolean if a field has been set.
func (o *Results) HasStatus504() bool {
	if o != nil && o.Status504 != nil {
		return true
	}

	return false
}

// SetStatus504 gets a reference to the given int32 and assigns it to the Status504 field.
func (o *Results) SetStatus504(v int32) {
	o.Status504 = &v
}

// GetStatus505 returns the Status505 field value if set, zero value otherwise.
func (o *Results) GetStatus505() int32 {
	if o == nil || o.Status505 == nil {
		var ret int32
		return ret
	}
	return *o.Status505
}

// GetStatus505Ok returns a tuple with the Status505 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetStatus505Ok() (*int32, bool) {
	if o == nil || o.Status505 == nil {
		return nil, false
	}
	return o.Status505, true
}

// HasStatus505 returns a boolean if a field has been set.
func (o *Results) HasStatus505() bool {
	if o != nil && o.Status505 != nil {
		return true
	}

	return false
}

// SetStatus505 gets a reference to the given int32 and assigns it to the Status505 field.
func (o *Results) SetStatus505(v int32) {
	o.Status505 = &v
}

// GetStatus1xx returns the Status1xx field value if set, zero value otherwise.
func (o *Results) GetStatus1xx() int32 {
	if o == nil || o.Status1xx == nil {
		var ret int32
		return ret
	}
	return *o.Status1xx
}

// GetStatus1xxOk returns a tuple with the Status1xx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetStatus1xxOk() (*int32, bool) {
	if o == nil || o.Status1xx == nil {
		return nil, false
	}
	return o.Status1xx, true
}

// HasStatus1xx returns a boolean if a field has been set.
func (o *Results) HasStatus1xx() bool {
	if o != nil && o.Status1xx != nil {
		return true
	}

	return false
}

// SetStatus1xx gets a reference to the given int32 and assigns it to the Status1xx field.
func (o *Results) SetStatus1xx(v int32) {
	o.Status1xx = &v
}

// GetStatus2xx returns the Status2xx field value if set, zero value otherwise.
func (o *Results) GetStatus2xx() int32 {
	if o == nil || o.Status2xx == nil {
		var ret int32
		return ret
	}
	return *o.Status2xx
}

// GetStatus2xxOk returns a tuple with the Status2xx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetStatus2xxOk() (*int32, bool) {
	if o == nil || o.Status2xx == nil {
		return nil, false
	}
	return o.Status2xx, true
}

// HasStatus2xx returns a boolean if a field has been set.
func (o *Results) HasStatus2xx() bool {
	if o != nil && o.Status2xx != nil {
		return true
	}

	return false
}

// SetStatus2xx gets a reference to the given int32 and assigns it to the Status2xx field.
func (o *Results) SetStatus2xx(v int32) {
	o.Status2xx = &v
}

// GetStatus3xx returns the Status3xx field value if set, zero value otherwise.
func (o *Results) GetStatus3xx() int32 {
	if o == nil || o.Status3xx == nil {
		var ret int32
		return ret
	}
	return *o.Status3xx
}

// GetStatus3xxOk returns a tuple with the Status3xx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetStatus3xxOk() (*int32, bool) {
	if o == nil || o.Status3xx == nil {
		return nil, false
	}
	return o.Status3xx, true
}

// HasStatus3xx returns a boolean if a field has been set.
func (o *Results) HasStatus3xx() bool {
	if o != nil && o.Status3xx != nil {
		return true
	}

	return false
}

// SetStatus3xx gets a reference to the given int32 and assigns it to the Status3xx field.
func (o *Results) SetStatus3xx(v int32) {
	o.Status3xx = &v
}

// GetStatus4xx returns the Status4xx field value if set, zero value otherwise.
func (o *Results) GetStatus4xx() int32 {
	if o == nil || o.Status4xx == nil {
		var ret int32
		return ret
	}
	return *o.Status4xx
}

// GetStatus4xxOk returns a tuple with the Status4xx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetStatus4xxOk() (*int32, bool) {
	if o == nil || o.Status4xx == nil {
		return nil, false
	}
	return o.Status4xx, true
}

// HasStatus4xx returns a boolean if a field has been set.
func (o *Results) HasStatus4xx() bool {
	if o != nil && o.Status4xx != nil {
		return true
	}

	return false
}

// SetStatus4xx gets a reference to the given int32 and assigns it to the Status4xx field.
func (o *Results) SetStatus4xx(v int32) {
	o.Status4xx = &v
}

// GetStatus5xx returns the Status5xx field value if set, zero value otherwise.
func (o *Results) GetStatus5xx() int32 {
	if o == nil || o.Status5xx == nil {
		var ret int32
		return ret
	}
	return *o.Status5xx
}

// GetStatus5xxOk returns a tuple with the Status5xx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetStatus5xxOk() (*int32, bool) {
	if o == nil || o.Status5xx == nil {
		return nil, false
	}
	return o.Status5xx, true
}

// HasStatus5xx returns a boolean if a field has been set.
func (o *Results) HasStatus5xx() bool {
	if o != nil && o.Status5xx != nil {
		return true
	}

	return false
}

// SetStatus5xx gets a reference to the given int32 and assigns it to the Status5xx field.
func (o *Results) SetStatus5xx(v int32) {
	o.Status5xx = &v
}

// GetObjectSize1k returns the ObjectSize1k field value if set, zero value otherwise.
func (o *Results) GetObjectSize1k() int32 {
	if o == nil || o.ObjectSize1k == nil {
		var ret int32
		return ret
	}
	return *o.ObjectSize1k
}

// GetObjectSize1kOk returns a tuple with the ObjectSize1k field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetObjectSize1kOk() (*int32, bool) {
	if o == nil || o.ObjectSize1k == nil {
		return nil, false
	}
	return o.ObjectSize1k, true
}

// HasObjectSize1k returns a boolean if a field has been set.
func (o *Results) HasObjectSize1k() bool {
	if o != nil && o.ObjectSize1k != nil {
		return true
	}

	return false
}

// SetObjectSize1k gets a reference to the given int32 and assigns it to the ObjectSize1k field.
func (o *Results) SetObjectSize1k(v int32) {
	o.ObjectSize1k = &v
}

// GetObjectSize10k returns the ObjectSize10k field value if set, zero value otherwise.
func (o *Results) GetObjectSize10k() int32 {
	if o == nil || o.ObjectSize10k == nil {
		var ret int32
		return ret
	}
	return *o.ObjectSize10k
}

// GetObjectSize10kOk returns a tuple with the ObjectSize10k field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetObjectSize10kOk() (*int32, bool) {
	if o == nil || o.ObjectSize10k == nil {
		return nil, false
	}
	return o.ObjectSize10k, true
}

// HasObjectSize10k returns a boolean if a field has been set.
func (o *Results) HasObjectSize10k() bool {
	if o != nil && o.ObjectSize10k != nil {
		return true
	}

	return false
}

// SetObjectSize10k gets a reference to the given int32 and assigns it to the ObjectSize10k field.
func (o *Results) SetObjectSize10k(v int32) {
	o.ObjectSize10k = &v
}

// GetObjectSize100k returns the ObjectSize100k field value if set, zero value otherwise.
func (o *Results) GetObjectSize100k() int32 {
	if o == nil || o.ObjectSize100k == nil {
		var ret int32
		return ret
	}
	return *o.ObjectSize100k
}

// GetObjectSize100kOk returns a tuple with the ObjectSize100k field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetObjectSize100kOk() (*int32, bool) {
	if o == nil || o.ObjectSize100k == nil {
		return nil, false
	}
	return o.ObjectSize100k, true
}

// HasObjectSize100k returns a boolean if a field has been set.
func (o *Results) HasObjectSize100k() bool {
	if o != nil && o.ObjectSize100k != nil {
		return true
	}

	return false
}

// SetObjectSize100k gets a reference to the given int32 and assigns it to the ObjectSize100k field.
func (o *Results) SetObjectSize100k(v int32) {
	o.ObjectSize100k = &v
}

// GetObjectSize1m returns the ObjectSize1m field value if set, zero value otherwise.
func (o *Results) GetObjectSize1m() int32 {
	if o == nil || o.ObjectSize1m == nil {
		var ret int32
		return ret
	}
	return *o.ObjectSize1m
}

// GetObjectSize1mOk returns a tuple with the ObjectSize1m field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetObjectSize1mOk() (*int32, bool) {
	if o == nil || o.ObjectSize1m == nil {
		return nil, false
	}
	return o.ObjectSize1m, true
}

// HasObjectSize1m returns a boolean if a field has been set.
func (o *Results) HasObjectSize1m() bool {
	if o != nil && o.ObjectSize1m != nil {
		return true
	}

	return false
}

// SetObjectSize1m gets a reference to the given int32 and assigns it to the ObjectSize1m field.
func (o *Results) SetObjectSize1m(v int32) {
	o.ObjectSize1m = &v
}

// GetObjectSize10m returns the ObjectSize10m field value if set, zero value otherwise.
func (o *Results) GetObjectSize10m() int32 {
	if o == nil || o.ObjectSize10m == nil {
		var ret int32
		return ret
	}
	return *o.ObjectSize10m
}

// GetObjectSize10mOk returns a tuple with the ObjectSize10m field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetObjectSize10mOk() (*int32, bool) {
	if o == nil || o.ObjectSize10m == nil {
		return nil, false
	}
	return o.ObjectSize10m, true
}

// HasObjectSize10m returns a boolean if a field has been set.
func (o *Results) HasObjectSize10m() bool {
	if o != nil && o.ObjectSize10m != nil {
		return true
	}

	return false
}

// SetObjectSize10m gets a reference to the given int32 and assigns it to the ObjectSize10m field.
func (o *Results) SetObjectSize10m(v int32) {
	o.ObjectSize10m = &v
}

// GetObjectSize100m returns the ObjectSize100m field value if set, zero value otherwise.
func (o *Results) GetObjectSize100m() int32 {
	if o == nil || o.ObjectSize100m == nil {
		var ret int32
		return ret
	}
	return *o.ObjectSize100m
}

// GetObjectSize100mOk returns a tuple with the ObjectSize100m field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetObjectSize100mOk() (*int32, bool) {
	if o == nil || o.ObjectSize100m == nil {
		return nil, false
	}
	return o.ObjectSize100m, true
}

// HasObjectSize100m returns a boolean if a field has been set.
func (o *Results) HasObjectSize100m() bool {
	if o != nil && o.ObjectSize100m != nil {
		return true
	}

	return false
}

// SetObjectSize100m gets a reference to the given int32 and assigns it to the ObjectSize100m field.
func (o *Results) SetObjectSize100m(v int32) {
	o.ObjectSize100m = &v
}

// GetObjectSize1g returns the ObjectSize1g field value if set, zero value otherwise.
func (o *Results) GetObjectSize1g() int32 {
	if o == nil || o.ObjectSize1g == nil {
		var ret int32
		return ret
	}
	return *o.ObjectSize1g
}

// GetObjectSize1gOk returns a tuple with the ObjectSize1g field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetObjectSize1gOk() (*int32, bool) {
	if o == nil || o.ObjectSize1g == nil {
		return nil, false
	}
	return o.ObjectSize1g, true
}

// HasObjectSize1g returns a boolean if a field has been set.
func (o *Results) HasObjectSize1g() bool {
	if o != nil && o.ObjectSize1g != nil {
		return true
	}

	return false
}

// SetObjectSize1g gets a reference to the given int32 and assigns it to the ObjectSize1g field.
func (o *Results) SetObjectSize1g(v int32) {
	o.ObjectSize1g = &v
}

// GetRecvSubTime returns the RecvSubTime field value if set, zero value otherwise.
func (o *Results) GetRecvSubTime() float32 {
	if o == nil || o.RecvSubTime == nil {
		var ret float32
		return ret
	}
	return *o.RecvSubTime
}

// GetRecvSubTimeOk returns a tuple with the RecvSubTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetRecvSubTimeOk() (*float32, bool) {
	if o == nil || o.RecvSubTime == nil {
		return nil, false
	}
	return o.RecvSubTime, true
}

// HasRecvSubTime returns a boolean if a field has been set.
func (o *Results) HasRecvSubTime() bool {
	if o != nil && o.RecvSubTime != nil {
		return true
	}

	return false
}

// SetRecvSubTime gets a reference to the given float32 and assigns it to the RecvSubTime field.
func (o *Results) SetRecvSubTime(v float32) {
	o.RecvSubTime = &v
}

// GetRecvSubCount returns the RecvSubCount field value if set, zero value otherwise.
func (o *Results) GetRecvSubCount() int32 {
	if o == nil || o.RecvSubCount == nil {
		var ret int32
		return ret
	}
	return *o.RecvSubCount
}

// GetRecvSubCountOk returns a tuple with the RecvSubCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetRecvSubCountOk() (*int32, bool) {
	if o == nil || o.RecvSubCount == nil {
		return nil, false
	}
	return o.RecvSubCount, true
}

// HasRecvSubCount returns a boolean if a field has been set.
func (o *Results) HasRecvSubCount() bool {
	if o != nil && o.RecvSubCount != nil {
		return true
	}

	return false
}

// SetRecvSubCount gets a reference to the given int32 and assigns it to the RecvSubCount field.
func (o *Results) SetRecvSubCount(v int32) {
	o.RecvSubCount = &v
}

// GetHashSubTime returns the HashSubTime field value if set, zero value otherwise.
func (o *Results) GetHashSubTime() float32 {
	if o == nil || o.HashSubTime == nil {
		var ret float32
		return ret
	}
	return *o.HashSubTime
}

// GetHashSubTimeOk returns a tuple with the HashSubTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetHashSubTimeOk() (*float32, bool) {
	if o == nil || o.HashSubTime == nil {
		return nil, false
	}
	return o.HashSubTime, true
}

// HasHashSubTime returns a boolean if a field has been set.
func (o *Results) HasHashSubTime() bool {
	if o != nil && o.HashSubTime != nil {
		return true
	}

	return false
}

// SetHashSubTime gets a reference to the given float32 and assigns it to the HashSubTime field.
func (o *Results) SetHashSubTime(v float32) {
	o.HashSubTime = &v
}

// GetHashSubCount returns the HashSubCount field value if set, zero value otherwise.
func (o *Results) GetHashSubCount() int32 {
	if o == nil || o.HashSubCount == nil {
		var ret int32
		return ret
	}
	return *o.HashSubCount
}

// GetHashSubCountOk returns a tuple with the HashSubCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetHashSubCountOk() (*int32, bool) {
	if o == nil || o.HashSubCount == nil {
		return nil, false
	}
	return o.HashSubCount, true
}

// HasHashSubCount returns a boolean if a field has been set.
func (o *Results) HasHashSubCount() bool {
	if o != nil && o.HashSubCount != nil {
		return true
	}

	return false
}

// SetHashSubCount gets a reference to the given int32 and assigns it to the HashSubCount field.
func (o *Results) SetHashSubCount(v int32) {
	o.HashSubCount = &v
}

// GetMissSubTime returns the MissSubTime field value if set, zero value otherwise.
func (o *Results) GetMissSubTime() float32 {
	if o == nil || o.MissSubTime == nil {
		var ret float32
		return ret
	}
	return *o.MissSubTime
}

// GetMissSubTimeOk returns a tuple with the MissSubTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetMissSubTimeOk() (*float32, bool) {
	if o == nil || o.MissSubTime == nil {
		return nil, false
	}
	return o.MissSubTime, true
}

// HasMissSubTime returns a boolean if a field has been set.
func (o *Results) HasMissSubTime() bool {
	if o != nil && o.MissSubTime != nil {
		return true
	}

	return false
}

// SetMissSubTime gets a reference to the given float32 and assigns it to the MissSubTime field.
func (o *Results) SetMissSubTime(v float32) {
	o.MissSubTime = &v
}

// GetMissSubCount returns the MissSubCount field value if set, zero value otherwise.
func (o *Results) GetMissSubCount() int32 {
	if o == nil || o.MissSubCount == nil {
		var ret int32
		return ret
	}
	return *o.MissSubCount
}

// GetMissSubCountOk returns a tuple with the MissSubCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetMissSubCountOk() (*int32, bool) {
	if o == nil || o.MissSubCount == nil {
		return nil, false
	}
	return o.MissSubCount, true
}

// HasMissSubCount returns a boolean if a field has been set.
func (o *Results) HasMissSubCount() bool {
	if o != nil && o.MissSubCount != nil {
		return true
	}

	return false
}

// SetMissSubCount gets a reference to the given int32 and assigns it to the MissSubCount field.
func (o *Results) SetMissSubCount(v int32) {
	o.MissSubCount = &v
}

// GetFetchSubTime returns the FetchSubTime field value if set, zero value otherwise.
func (o *Results) GetFetchSubTime() float32 {
	if o == nil || o.FetchSubTime == nil {
		var ret float32
		return ret
	}
	return *o.FetchSubTime
}

// GetFetchSubTimeOk returns a tuple with the FetchSubTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetFetchSubTimeOk() (*float32, bool) {
	if o == nil || o.FetchSubTime == nil {
		return nil, false
	}
	return o.FetchSubTime, true
}

// HasFetchSubTime returns a boolean if a field has been set.
func (o *Results) HasFetchSubTime() bool {
	if o != nil && o.FetchSubTime != nil {
		return true
	}

	return false
}

// SetFetchSubTime gets a reference to the given float32 and assigns it to the FetchSubTime field.
func (o *Results) SetFetchSubTime(v float32) {
	o.FetchSubTime = &v
}

// GetFetchSubCount returns the FetchSubCount field value if set, zero value otherwise.
func (o *Results) GetFetchSubCount() int32 {
	if o == nil || o.FetchSubCount == nil {
		var ret int32
		return ret
	}
	return *o.FetchSubCount
}

// GetFetchSubCountOk returns a tuple with the FetchSubCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetFetchSubCountOk() (*int32, bool) {
	if o == nil || o.FetchSubCount == nil {
		return nil, false
	}
	return o.FetchSubCount, true
}

// HasFetchSubCount returns a boolean if a field has been set.
func (o *Results) HasFetchSubCount() bool {
	if o != nil && o.FetchSubCount != nil {
		return true
	}

	return false
}

// SetFetchSubCount gets a reference to the given int32 and assigns it to the FetchSubCount field.
func (o *Results) SetFetchSubCount(v int32) {
	o.FetchSubCount = &v
}

// GetPassSubTime returns the PassSubTime field value if set, zero value otherwise.
func (o *Results) GetPassSubTime() float32 {
	if o == nil || o.PassSubTime == nil {
		var ret float32
		return ret
	}
	return *o.PassSubTime
}

// GetPassSubTimeOk returns a tuple with the PassSubTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetPassSubTimeOk() (*float32, bool) {
	if o == nil || o.PassSubTime == nil {
		return nil, false
	}
	return o.PassSubTime, true
}

// HasPassSubTime returns a boolean if a field has been set.
func (o *Results) HasPassSubTime() bool {
	if o != nil && o.PassSubTime != nil {
		return true
	}

	return false
}

// SetPassSubTime gets a reference to the given float32 and assigns it to the PassSubTime field.
func (o *Results) SetPassSubTime(v float32) {
	o.PassSubTime = &v
}

// GetPassSubCount returns the PassSubCount field value if set, zero value otherwise.
func (o *Results) GetPassSubCount() int32 {
	if o == nil || o.PassSubCount == nil {
		var ret int32
		return ret
	}
	return *o.PassSubCount
}

// GetPassSubCountOk returns a tuple with the PassSubCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetPassSubCountOk() (*int32, bool) {
	if o == nil || o.PassSubCount == nil {
		return nil, false
	}
	return o.PassSubCount, true
}

// HasPassSubCount returns a boolean if a field has been set.
func (o *Results) HasPassSubCount() bool {
	if o != nil && o.PassSubCount != nil {
		return true
	}

	return false
}

// SetPassSubCount gets a reference to the given int32 and assigns it to the PassSubCount field.
func (o *Results) SetPassSubCount(v int32) {
	o.PassSubCount = &v
}

// GetPipeSubTime returns the PipeSubTime field value if set, zero value otherwise.
func (o *Results) GetPipeSubTime() float32 {
	if o == nil || o.PipeSubTime == nil {
		var ret float32
		return ret
	}
	return *o.PipeSubTime
}

// GetPipeSubTimeOk returns a tuple with the PipeSubTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetPipeSubTimeOk() (*float32, bool) {
	if o == nil || o.PipeSubTime == nil {
		return nil, false
	}
	return o.PipeSubTime, true
}

// HasPipeSubTime returns a boolean if a field has been set.
func (o *Results) HasPipeSubTime() bool {
	if o != nil && o.PipeSubTime != nil {
		return true
	}

	return false
}

// SetPipeSubTime gets a reference to the given float32 and assigns it to the PipeSubTime field.
func (o *Results) SetPipeSubTime(v float32) {
	o.PipeSubTime = &v
}

// GetPipeSubCount returns the PipeSubCount field value if set, zero value otherwise.
func (o *Results) GetPipeSubCount() int32 {
	if o == nil || o.PipeSubCount == nil {
		var ret int32
		return ret
	}
	return *o.PipeSubCount
}

// GetPipeSubCountOk returns a tuple with the PipeSubCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetPipeSubCountOk() (*int32, bool) {
	if o == nil || o.PipeSubCount == nil {
		return nil, false
	}
	return o.PipeSubCount, true
}

// HasPipeSubCount returns a boolean if a field has been set.
func (o *Results) HasPipeSubCount() bool {
	if o != nil && o.PipeSubCount != nil {
		return true
	}

	return false
}

// SetPipeSubCount gets a reference to the given int32 and assigns it to the PipeSubCount field.
func (o *Results) SetPipeSubCount(v int32) {
	o.PipeSubCount = &v
}

// GetDeliverSubTime returns the DeliverSubTime field value if set, zero value otherwise.
func (o *Results) GetDeliverSubTime() float32 {
	if o == nil || o.DeliverSubTime == nil {
		var ret float32
		return ret
	}
	return *o.DeliverSubTime
}

// GetDeliverSubTimeOk returns a tuple with the DeliverSubTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetDeliverSubTimeOk() (*float32, bool) {
	if o == nil || o.DeliverSubTime == nil {
		return nil, false
	}
	return o.DeliverSubTime, true
}

// HasDeliverSubTime returns a boolean if a field has been set.
func (o *Results) HasDeliverSubTime() bool {
	if o != nil && o.DeliverSubTime != nil {
		return true
	}

	return false
}

// SetDeliverSubTime gets a reference to the given float32 and assigns it to the DeliverSubTime field.
func (o *Results) SetDeliverSubTime(v float32) {
	o.DeliverSubTime = &v
}

// GetDeliverSubCount returns the DeliverSubCount field value if set, zero value otherwise.
func (o *Results) GetDeliverSubCount() int32 {
	if o == nil || o.DeliverSubCount == nil {
		var ret int32
		return ret
	}
	return *o.DeliverSubCount
}

// GetDeliverSubCountOk returns a tuple with the DeliverSubCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetDeliverSubCountOk() (*int32, bool) {
	if o == nil || o.DeliverSubCount == nil {
		return nil, false
	}
	return o.DeliverSubCount, true
}

// HasDeliverSubCount returns a boolean if a field has been set.
func (o *Results) HasDeliverSubCount() bool {
	if o != nil && o.DeliverSubCount != nil {
		return true
	}

	return false
}

// SetDeliverSubCount gets a reference to the given int32 and assigns it to the DeliverSubCount field.
func (o *Results) SetDeliverSubCount(v int32) {
	o.DeliverSubCount = &v
}

// GetErrorSubTime returns the ErrorSubTime field value if set, zero value otherwise.
func (o *Results) GetErrorSubTime() float32 {
	if o == nil || o.ErrorSubTime == nil {
		var ret float32
		return ret
	}
	return *o.ErrorSubTime
}

// GetErrorSubTimeOk returns a tuple with the ErrorSubTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetErrorSubTimeOk() (*float32, bool) {
	if o == nil || o.ErrorSubTime == nil {
		return nil, false
	}
	return o.ErrorSubTime, true
}

// HasErrorSubTime returns a boolean if a field has been set.
func (o *Results) HasErrorSubTime() bool {
	if o != nil && o.ErrorSubTime != nil {
		return true
	}

	return false
}

// SetErrorSubTime gets a reference to the given float32 and assigns it to the ErrorSubTime field.
func (o *Results) SetErrorSubTime(v float32) {
	o.ErrorSubTime = &v
}

// GetErrorSubCount returns the ErrorSubCount field value if set, zero value otherwise.
func (o *Results) GetErrorSubCount() int32 {
	if o == nil || o.ErrorSubCount == nil {
		var ret int32
		return ret
	}
	return *o.ErrorSubCount
}

// GetErrorSubCountOk returns a tuple with the ErrorSubCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetErrorSubCountOk() (*int32, bool) {
	if o == nil || o.ErrorSubCount == nil {
		return nil, false
	}
	return o.ErrorSubCount, true
}

// HasErrorSubCount returns a boolean if a field has been set.
func (o *Results) HasErrorSubCount() bool {
	if o != nil && o.ErrorSubCount != nil {
		return true
	}

	return false
}

// SetErrorSubCount gets a reference to the given int32 and assigns it to the ErrorSubCount field.
func (o *Results) SetErrorSubCount(v int32) {
	o.ErrorSubCount = &v
}

// GetHitSubTime returns the HitSubTime field value if set, zero value otherwise.
func (o *Results) GetHitSubTime() float32 {
	if o == nil || o.HitSubTime == nil {
		var ret float32
		return ret
	}
	return *o.HitSubTime
}

// GetHitSubTimeOk returns a tuple with the HitSubTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetHitSubTimeOk() (*float32, bool) {
	if o == nil || o.HitSubTime == nil {
		return nil, false
	}
	return o.HitSubTime, true
}

// HasHitSubTime returns a boolean if a field has been set.
func (o *Results) HasHitSubTime() bool {
	if o != nil && o.HitSubTime != nil {
		return true
	}

	return false
}

// SetHitSubTime gets a reference to the given float32 and assigns it to the HitSubTime field.
func (o *Results) SetHitSubTime(v float32) {
	o.HitSubTime = &v
}

// GetHitSubCount returns the HitSubCount field value if set, zero value otherwise.
func (o *Results) GetHitSubCount() int32 {
	if o == nil || o.HitSubCount == nil {
		var ret int32
		return ret
	}
	return *o.HitSubCount
}

// GetHitSubCountOk returns a tuple with the HitSubCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetHitSubCountOk() (*int32, bool) {
	if o == nil || o.HitSubCount == nil {
		return nil, false
	}
	return o.HitSubCount, true
}

// HasHitSubCount returns a boolean if a field has been set.
func (o *Results) HasHitSubCount() bool {
	if o != nil && o.HitSubCount != nil {
		return true
	}

	return false
}

// SetHitSubCount gets a reference to the given int32 and assigns it to the HitSubCount field.
func (o *Results) SetHitSubCount(v int32) {
	o.HitSubCount = &v
}

// GetPrehashSubTime returns the PrehashSubTime field value if set, zero value otherwise.
func (o *Results) GetPrehashSubTime() float32 {
	if o == nil || o.PrehashSubTime == nil {
		var ret float32
		return ret
	}
	return *o.PrehashSubTime
}

// GetPrehashSubTimeOk returns a tuple with the PrehashSubTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetPrehashSubTimeOk() (*float32, bool) {
	if o == nil || o.PrehashSubTime == nil {
		return nil, false
	}
	return o.PrehashSubTime, true
}

// HasPrehashSubTime returns a boolean if a field has been set.
func (o *Results) HasPrehashSubTime() bool {
	if o != nil && o.PrehashSubTime != nil {
		return true
	}

	return false
}

// SetPrehashSubTime gets a reference to the given float32 and assigns it to the PrehashSubTime field.
func (o *Results) SetPrehashSubTime(v float32) {
	o.PrehashSubTime = &v
}

// GetPrehashSubCount returns the PrehashSubCount field value if set, zero value otherwise.
func (o *Results) GetPrehashSubCount() int32 {
	if o == nil || o.PrehashSubCount == nil {
		var ret int32
		return ret
	}
	return *o.PrehashSubCount
}

// GetPrehashSubCountOk returns a tuple with the PrehashSubCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetPrehashSubCountOk() (*int32, bool) {
	if o == nil || o.PrehashSubCount == nil {
		return nil, false
	}
	return o.PrehashSubCount, true
}

// HasPrehashSubCount returns a boolean if a field has been set.
func (o *Results) HasPrehashSubCount() bool {
	if o != nil && o.PrehashSubCount != nil {
		return true
	}

	return false
}

// SetPrehashSubCount gets a reference to the given int32 and assigns it to the PrehashSubCount field.
func (o *Results) SetPrehashSubCount(v int32) {
	o.PrehashSubCount = &v
}

// GetPredeliverSubTime returns the PredeliverSubTime field value if set, zero value otherwise.
func (o *Results) GetPredeliverSubTime() float32 {
	if o == nil || o.PredeliverSubTime == nil {
		var ret float32
		return ret
	}
	return *o.PredeliverSubTime
}

// GetPredeliverSubTimeOk returns a tuple with the PredeliverSubTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetPredeliverSubTimeOk() (*float32, bool) {
	if o == nil || o.PredeliverSubTime == nil {
		return nil, false
	}
	return o.PredeliverSubTime, true
}

// HasPredeliverSubTime returns a boolean if a field has been set.
func (o *Results) HasPredeliverSubTime() bool {
	if o != nil && o.PredeliverSubTime != nil {
		return true
	}

	return false
}

// SetPredeliverSubTime gets a reference to the given float32 and assigns it to the PredeliverSubTime field.
func (o *Results) SetPredeliverSubTime(v float32) {
	o.PredeliverSubTime = &v
}

// GetPredeliverSubCount returns the PredeliverSubCount field value if set, zero value otherwise.
func (o *Results) GetPredeliverSubCount() int32 {
	if o == nil || o.PredeliverSubCount == nil {
		var ret int32
		return ret
	}
	return *o.PredeliverSubCount
}

// GetPredeliverSubCountOk returns a tuple with the PredeliverSubCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetPredeliverSubCountOk() (*int32, bool) {
	if o == nil || o.PredeliverSubCount == nil {
		return nil, false
	}
	return o.PredeliverSubCount, true
}

// HasPredeliverSubCount returns a boolean if a field has been set.
func (o *Results) HasPredeliverSubCount() bool {
	if o != nil && o.PredeliverSubCount != nil {
		return true
	}

	return false
}

// SetPredeliverSubCount gets a reference to the given int32 and assigns it to the PredeliverSubCount field.
func (o *Results) SetPredeliverSubCount(v int32) {
	o.PredeliverSubCount = &v
}

// GetTLSHandshakeSentBytes returns the TLSHandshakeSentBytes field value if set, zero value otherwise.
func (o *Results) GetTLSHandshakeSentBytes() int32 {
	if o == nil || o.TLSHandshakeSentBytes == nil {
		var ret int32
		return ret
	}
	return *o.TLSHandshakeSentBytes
}

// GetTLSHandshakeSentBytesOk returns a tuple with the TLSHandshakeSentBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetTLSHandshakeSentBytesOk() (*int32, bool) {
	if o == nil || o.TLSHandshakeSentBytes == nil {
		return nil, false
	}
	return o.TLSHandshakeSentBytes, true
}

// HasTLSHandshakeSentBytes returns a boolean if a field has been set.
func (o *Results) HasTLSHandshakeSentBytes() bool {
	if o != nil && o.TLSHandshakeSentBytes != nil {
		return true
	}

	return false
}

// SetTLSHandshakeSentBytes gets a reference to the given int32 and assigns it to the TLSHandshakeSentBytes field.
func (o *Results) SetTLSHandshakeSentBytes(v int32) {
	o.TLSHandshakeSentBytes = &v
}

// GetHitRespBodyBytes returns the HitRespBodyBytes field value if set, zero value otherwise.
func (o *Results) GetHitRespBodyBytes() int32 {
	if o == nil || o.HitRespBodyBytes == nil {
		var ret int32
		return ret
	}
	return *o.HitRespBodyBytes
}

// GetHitRespBodyBytesOk returns a tuple with the HitRespBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetHitRespBodyBytesOk() (*int32, bool) {
	if o == nil || o.HitRespBodyBytes == nil {
		return nil, false
	}
	return o.HitRespBodyBytes, true
}

// HasHitRespBodyBytes returns a boolean if a field has been set.
func (o *Results) HasHitRespBodyBytes() bool {
	if o != nil && o.HitRespBodyBytes != nil {
		return true
	}

	return false
}

// SetHitRespBodyBytes gets a reference to the given int32 and assigns it to the HitRespBodyBytes field.
func (o *Results) SetHitRespBodyBytes(v int32) {
	o.HitRespBodyBytes = &v
}

// GetMissRespBodyBytes returns the MissRespBodyBytes field value if set, zero value otherwise.
func (o *Results) GetMissRespBodyBytes() int32 {
	if o == nil || o.MissRespBodyBytes == nil {
		var ret int32
		return ret
	}
	return *o.MissRespBodyBytes
}

// GetMissRespBodyBytesOk returns a tuple with the MissRespBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetMissRespBodyBytesOk() (*int32, bool) {
	if o == nil || o.MissRespBodyBytes == nil {
		return nil, false
	}
	return o.MissRespBodyBytes, true
}

// HasMissRespBodyBytes returns a boolean if a field has been set.
func (o *Results) HasMissRespBodyBytes() bool {
	if o != nil && o.MissRespBodyBytes != nil {
		return true
	}

	return false
}

// SetMissRespBodyBytes gets a reference to the given int32 and assigns it to the MissRespBodyBytes field.
func (o *Results) SetMissRespBodyBytes(v int32) {
	o.MissRespBodyBytes = &v
}

// GetPassRespBodyBytes returns the PassRespBodyBytes field value if set, zero value otherwise.
func (o *Results) GetPassRespBodyBytes() int32 {
	if o == nil || o.PassRespBodyBytes == nil {
		var ret int32
		return ret
	}
	return *o.PassRespBodyBytes
}

// GetPassRespBodyBytesOk returns a tuple with the PassRespBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetPassRespBodyBytesOk() (*int32, bool) {
	if o == nil || o.PassRespBodyBytes == nil {
		return nil, false
	}
	return o.PassRespBodyBytes, true
}

// HasPassRespBodyBytes returns a boolean if a field has been set.
func (o *Results) HasPassRespBodyBytes() bool {
	if o != nil && o.PassRespBodyBytes != nil {
		return true
	}

	return false
}

// SetPassRespBodyBytes gets a reference to the given int32 and assigns it to the PassRespBodyBytes field.
func (o *Results) SetPassRespBodyBytes(v int32) {
	o.PassRespBodyBytes = &v
}

// GetSegblockOriginFetches returns the SegblockOriginFetches field value if set, zero value otherwise.
func (o *Results) GetSegblockOriginFetches() int32 {
	if o == nil || o.SegblockOriginFetches == nil {
		var ret int32
		return ret
	}
	return *o.SegblockOriginFetches
}

// GetSegblockOriginFetchesOk returns a tuple with the SegblockOriginFetches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetSegblockOriginFetchesOk() (*int32, bool) {
	if o == nil || o.SegblockOriginFetches == nil {
		return nil, false
	}
	return o.SegblockOriginFetches, true
}

// HasSegblockOriginFetches returns a boolean if a field has been set.
func (o *Results) HasSegblockOriginFetches() bool {
	if o != nil && o.SegblockOriginFetches != nil {
		return true
	}

	return false
}

// SetSegblockOriginFetches gets a reference to the given int32 and assigns it to the SegblockOriginFetches field.
func (o *Results) SetSegblockOriginFetches(v int32) {
	o.SegblockOriginFetches = &v
}

// GetSegblockShieldFetches returns the SegblockShieldFetches field value if set, zero value otherwise.
func (o *Results) GetSegblockShieldFetches() int32 {
	if o == nil || o.SegblockShieldFetches == nil {
		var ret int32
		return ret
	}
	return *o.SegblockShieldFetches
}

// GetSegblockShieldFetchesOk returns a tuple with the SegblockShieldFetches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetSegblockShieldFetchesOk() (*int32, bool) {
	if o == nil || o.SegblockShieldFetches == nil {
		return nil, false
	}
	return o.SegblockShieldFetches, true
}

// HasSegblockShieldFetches returns a boolean if a field has been set.
func (o *Results) HasSegblockShieldFetches() bool {
	if o != nil && o.SegblockShieldFetches != nil {
		return true
	}

	return false
}

// SetSegblockShieldFetches gets a reference to the given int32 and assigns it to the SegblockShieldFetches field.
func (o *Results) SetSegblockShieldFetches(v int32) {
	o.SegblockShieldFetches = &v
}

// GetComputeRequests returns the ComputeRequests field value if set, zero value otherwise.
func (o *Results) GetComputeRequests() int32 {
	if o == nil || o.ComputeRequests == nil {
		var ret int32
		return ret
	}
	return *o.ComputeRequests
}

// GetComputeRequestsOk returns a tuple with the ComputeRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetComputeRequestsOk() (*int32, bool) {
	if o == nil || o.ComputeRequests == nil {
		return nil, false
	}
	return o.ComputeRequests, true
}

// HasComputeRequests returns a boolean if a field has been set.
func (o *Results) HasComputeRequests() bool {
	if o != nil && o.ComputeRequests != nil {
		return true
	}

	return false
}

// SetComputeRequests gets a reference to the given int32 and assigns it to the ComputeRequests field.
func (o *Results) SetComputeRequests(v int32) {
	o.ComputeRequests = &v
}

// GetComputeRequestTimeMs returns the ComputeRequestTimeMs field value if set, zero value otherwise.
func (o *Results) GetComputeRequestTimeMs() float32 {
	if o == nil || o.ComputeRequestTimeMs == nil {
		var ret float32
		return ret
	}
	return *o.ComputeRequestTimeMs
}

// GetComputeRequestTimeMsOk returns a tuple with the ComputeRequestTimeMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetComputeRequestTimeMsOk() (*float32, bool) {
	if o == nil || o.ComputeRequestTimeMs == nil {
		return nil, false
	}
	return o.ComputeRequestTimeMs, true
}

// HasComputeRequestTimeMs returns a boolean if a field has been set.
func (o *Results) HasComputeRequestTimeMs() bool {
	if o != nil && o.ComputeRequestTimeMs != nil {
		return true
	}

	return false
}

// SetComputeRequestTimeMs gets a reference to the given float32 and assigns it to the ComputeRequestTimeMs field.
func (o *Results) SetComputeRequestTimeMs(v float32) {
	o.ComputeRequestTimeMs = &v
}

// GetComputeRequestTimeBilledMs returns the ComputeRequestTimeBilledMs field value if set, zero value otherwise.
func (o *Results) GetComputeRequestTimeBilledMs() float32 {
	if o == nil || o.ComputeRequestTimeBilledMs == nil {
		var ret float32
		return ret
	}
	return *o.ComputeRequestTimeBilledMs
}

// GetComputeRequestTimeBilledMsOk returns a tuple with the ComputeRequestTimeBilledMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetComputeRequestTimeBilledMsOk() (*float32, bool) {
	if o == nil || o.ComputeRequestTimeBilledMs == nil {
		return nil, false
	}
	return o.ComputeRequestTimeBilledMs, true
}

// HasComputeRequestTimeBilledMs returns a boolean if a field has been set.
func (o *Results) HasComputeRequestTimeBilledMs() bool {
	if o != nil && o.ComputeRequestTimeBilledMs != nil {
		return true
	}

	return false
}

// SetComputeRequestTimeBilledMs gets a reference to the given float32 and assigns it to the ComputeRequestTimeBilledMs field.
func (o *Results) SetComputeRequestTimeBilledMs(v float32) {
	o.ComputeRequestTimeBilledMs = &v
}

// GetComputeRAMUsed returns the ComputeRAMUsed field value if set, zero value otherwise.
func (o *Results) GetComputeRAMUsed() int32 {
	if o == nil || o.ComputeRAMUsed == nil {
		var ret int32
		return ret
	}
	return *o.ComputeRAMUsed
}

// GetComputeRAMUsedOk returns a tuple with the ComputeRAMUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetComputeRAMUsedOk() (*int32, bool) {
	if o == nil || o.ComputeRAMUsed == nil {
		return nil, false
	}
	return o.ComputeRAMUsed, true
}

// HasComputeRAMUsed returns a boolean if a field has been set.
func (o *Results) HasComputeRAMUsed() bool {
	if o != nil && o.ComputeRAMUsed != nil {
		return true
	}

	return false
}

// SetComputeRAMUsed gets a reference to the given int32 and assigns it to the ComputeRAMUsed field.
func (o *Results) SetComputeRAMUsed(v int32) {
	o.ComputeRAMUsed = &v
}

// GetComputeExecutionTimeMs returns the ComputeExecutionTimeMs field value if set, zero value otherwise.
func (o *Results) GetComputeExecutionTimeMs() float32 {
	if o == nil || o.ComputeExecutionTimeMs == nil {
		var ret float32
		return ret
	}
	return *o.ComputeExecutionTimeMs
}

// GetComputeExecutionTimeMsOk returns a tuple with the ComputeExecutionTimeMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetComputeExecutionTimeMsOk() (*float32, bool) {
	if o == nil || o.ComputeExecutionTimeMs == nil {
		return nil, false
	}
	return o.ComputeExecutionTimeMs, true
}

// HasComputeExecutionTimeMs returns a boolean if a field has been set.
func (o *Results) HasComputeExecutionTimeMs() bool {
	if o != nil && o.ComputeExecutionTimeMs != nil {
		return true
	}

	return false
}

// SetComputeExecutionTimeMs gets a reference to the given float32 and assigns it to the ComputeExecutionTimeMs field.
func (o *Results) SetComputeExecutionTimeMs(v float32) {
	o.ComputeExecutionTimeMs = &v
}

// GetComputeReqHeaderBytes returns the ComputeReqHeaderBytes field value if set, zero value otherwise.
func (o *Results) GetComputeReqHeaderBytes() int32 {
	if o == nil || o.ComputeReqHeaderBytes == nil {
		var ret int32
		return ret
	}
	return *o.ComputeReqHeaderBytes
}

// GetComputeReqHeaderBytesOk returns a tuple with the ComputeReqHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetComputeReqHeaderBytesOk() (*int32, bool) {
	if o == nil || o.ComputeReqHeaderBytes == nil {
		return nil, false
	}
	return o.ComputeReqHeaderBytes, true
}

// HasComputeReqHeaderBytes returns a boolean if a field has been set.
func (o *Results) HasComputeReqHeaderBytes() bool {
	if o != nil && o.ComputeReqHeaderBytes != nil {
		return true
	}

	return false
}

// SetComputeReqHeaderBytes gets a reference to the given int32 and assigns it to the ComputeReqHeaderBytes field.
func (o *Results) SetComputeReqHeaderBytes(v int32) {
	o.ComputeReqHeaderBytes = &v
}

// GetComputeReqBodyBytes returns the ComputeReqBodyBytes field value if set, zero value otherwise.
func (o *Results) GetComputeReqBodyBytes() int32 {
	if o == nil || o.ComputeReqBodyBytes == nil {
		var ret int32
		return ret
	}
	return *o.ComputeReqBodyBytes
}

// GetComputeReqBodyBytesOk returns a tuple with the ComputeReqBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetComputeReqBodyBytesOk() (*int32, bool) {
	if o == nil || o.ComputeReqBodyBytes == nil {
		return nil, false
	}
	return o.ComputeReqBodyBytes, true
}

// HasComputeReqBodyBytes returns a boolean if a field has been set.
func (o *Results) HasComputeReqBodyBytes() bool {
	if o != nil && o.ComputeReqBodyBytes != nil {
		return true
	}

	return false
}

// SetComputeReqBodyBytes gets a reference to the given int32 and assigns it to the ComputeReqBodyBytes field.
func (o *Results) SetComputeReqBodyBytes(v int32) {
	o.ComputeReqBodyBytes = &v
}

// GetComputeRespHeaderBytes returns the ComputeRespHeaderBytes field value if set, zero value otherwise.
func (o *Results) GetComputeRespHeaderBytes() int32 {
	if o == nil || o.ComputeRespHeaderBytes == nil {
		var ret int32
		return ret
	}
	return *o.ComputeRespHeaderBytes
}

// GetComputeRespHeaderBytesOk returns a tuple with the ComputeRespHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetComputeRespHeaderBytesOk() (*int32, bool) {
	if o == nil || o.ComputeRespHeaderBytes == nil {
		return nil, false
	}
	return o.ComputeRespHeaderBytes, true
}

// HasComputeRespHeaderBytes returns a boolean if a field has been set.
func (o *Results) HasComputeRespHeaderBytes() bool {
	if o != nil && o.ComputeRespHeaderBytes != nil {
		return true
	}

	return false
}

// SetComputeRespHeaderBytes gets a reference to the given int32 and assigns it to the ComputeRespHeaderBytes field.
func (o *Results) SetComputeRespHeaderBytes(v int32) {
	o.ComputeRespHeaderBytes = &v
}

// GetComputeRespBodyBytes returns the ComputeRespBodyBytes field value if set, zero value otherwise.
func (o *Results) GetComputeRespBodyBytes() int32 {
	if o == nil || o.ComputeRespBodyBytes == nil {
		var ret int32
		return ret
	}
	return *o.ComputeRespBodyBytes
}

// GetComputeRespBodyBytesOk returns a tuple with the ComputeRespBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetComputeRespBodyBytesOk() (*int32, bool) {
	if o == nil || o.ComputeRespBodyBytes == nil {
		return nil, false
	}
	return o.ComputeRespBodyBytes, true
}

// HasComputeRespBodyBytes returns a boolean if a field has been set.
func (o *Results) HasComputeRespBodyBytes() bool {
	if o != nil && o.ComputeRespBodyBytes != nil {
		return true
	}

	return false
}

// SetComputeRespBodyBytes gets a reference to the given int32 and assigns it to the ComputeRespBodyBytes field.
func (o *Results) SetComputeRespBodyBytes(v int32) {
	o.ComputeRespBodyBytes = &v
}

// GetComputeRespStatus1xx returns the ComputeRespStatus1xx field value if set, zero value otherwise.
func (o *Results) GetComputeRespStatus1xx() int32 {
	if o == nil || o.ComputeRespStatus1xx == nil {
		var ret int32
		return ret
	}
	return *o.ComputeRespStatus1xx
}

// GetComputeRespStatus1xxOk returns a tuple with the ComputeRespStatus1xx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetComputeRespStatus1xxOk() (*int32, bool) {
	if o == nil || o.ComputeRespStatus1xx == nil {
		return nil, false
	}
	return o.ComputeRespStatus1xx, true
}

// HasComputeRespStatus1xx returns a boolean if a field has been set.
func (o *Results) HasComputeRespStatus1xx() bool {
	if o != nil && o.ComputeRespStatus1xx != nil {
		return true
	}

	return false
}

// SetComputeRespStatus1xx gets a reference to the given int32 and assigns it to the ComputeRespStatus1xx field.
func (o *Results) SetComputeRespStatus1xx(v int32) {
	o.ComputeRespStatus1xx = &v
}

// GetComputeRespStatus2xx returns the ComputeRespStatus2xx field value if set, zero value otherwise.
func (o *Results) GetComputeRespStatus2xx() int32 {
	if o == nil || o.ComputeRespStatus2xx == nil {
		var ret int32
		return ret
	}
	return *o.ComputeRespStatus2xx
}

// GetComputeRespStatus2xxOk returns a tuple with the ComputeRespStatus2xx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetComputeRespStatus2xxOk() (*int32, bool) {
	if o == nil || o.ComputeRespStatus2xx == nil {
		return nil, false
	}
	return o.ComputeRespStatus2xx, true
}

// HasComputeRespStatus2xx returns a boolean if a field has been set.
func (o *Results) HasComputeRespStatus2xx() bool {
	if o != nil && o.ComputeRespStatus2xx != nil {
		return true
	}

	return false
}

// SetComputeRespStatus2xx gets a reference to the given int32 and assigns it to the ComputeRespStatus2xx field.
func (o *Results) SetComputeRespStatus2xx(v int32) {
	o.ComputeRespStatus2xx = &v
}

// GetComputeRespStatus3xx returns the ComputeRespStatus3xx field value if set, zero value otherwise.
func (o *Results) GetComputeRespStatus3xx() int32 {
	if o == nil || o.ComputeRespStatus3xx == nil {
		var ret int32
		return ret
	}
	return *o.ComputeRespStatus3xx
}

// GetComputeRespStatus3xxOk returns a tuple with the ComputeRespStatus3xx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetComputeRespStatus3xxOk() (*int32, bool) {
	if o == nil || o.ComputeRespStatus3xx == nil {
		return nil, false
	}
	return o.ComputeRespStatus3xx, true
}

// HasComputeRespStatus3xx returns a boolean if a field has been set.
func (o *Results) HasComputeRespStatus3xx() bool {
	if o != nil && o.ComputeRespStatus3xx != nil {
		return true
	}

	return false
}

// SetComputeRespStatus3xx gets a reference to the given int32 and assigns it to the ComputeRespStatus3xx field.
func (o *Results) SetComputeRespStatus3xx(v int32) {
	o.ComputeRespStatus3xx = &v
}

// GetComputeRespStatus4xx returns the ComputeRespStatus4xx field value if set, zero value otherwise.
func (o *Results) GetComputeRespStatus4xx() int32 {
	if o == nil || o.ComputeRespStatus4xx == nil {
		var ret int32
		return ret
	}
	return *o.ComputeRespStatus4xx
}

// GetComputeRespStatus4xxOk returns a tuple with the ComputeRespStatus4xx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetComputeRespStatus4xxOk() (*int32, bool) {
	if o == nil || o.ComputeRespStatus4xx == nil {
		return nil, false
	}
	return o.ComputeRespStatus4xx, true
}

// HasComputeRespStatus4xx returns a boolean if a field has been set.
func (o *Results) HasComputeRespStatus4xx() bool {
	if o != nil && o.ComputeRespStatus4xx != nil {
		return true
	}

	return false
}

// SetComputeRespStatus4xx gets a reference to the given int32 and assigns it to the ComputeRespStatus4xx field.
func (o *Results) SetComputeRespStatus4xx(v int32) {
	o.ComputeRespStatus4xx = &v
}

// GetComputeRespStatus5xx returns the ComputeRespStatus5xx field value if set, zero value otherwise.
func (o *Results) GetComputeRespStatus5xx() int32 {
	if o == nil || o.ComputeRespStatus5xx == nil {
		var ret int32
		return ret
	}
	return *o.ComputeRespStatus5xx
}

// GetComputeRespStatus5xxOk returns a tuple with the ComputeRespStatus5xx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetComputeRespStatus5xxOk() (*int32, bool) {
	if o == nil || o.ComputeRespStatus5xx == nil {
		return nil, false
	}
	return o.ComputeRespStatus5xx, true
}

// HasComputeRespStatus5xx returns a boolean if a field has been set.
func (o *Results) HasComputeRespStatus5xx() bool {
	if o != nil && o.ComputeRespStatus5xx != nil {
		return true
	}

	return false
}

// SetComputeRespStatus5xx gets a reference to the given int32 and assigns it to the ComputeRespStatus5xx field.
func (o *Results) SetComputeRespStatus5xx(v int32) {
	o.ComputeRespStatus5xx = &v
}

// GetComputeBereqHeaderBytes returns the ComputeBereqHeaderBytes field value if set, zero value otherwise.
func (o *Results) GetComputeBereqHeaderBytes() int32 {
	if o == nil || o.ComputeBereqHeaderBytes == nil {
		var ret int32
		return ret
	}
	return *o.ComputeBereqHeaderBytes
}

// GetComputeBereqHeaderBytesOk returns a tuple with the ComputeBereqHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetComputeBereqHeaderBytesOk() (*int32, bool) {
	if o == nil || o.ComputeBereqHeaderBytes == nil {
		return nil, false
	}
	return o.ComputeBereqHeaderBytes, true
}

// HasComputeBereqHeaderBytes returns a boolean if a field has been set.
func (o *Results) HasComputeBereqHeaderBytes() bool {
	if o != nil && o.ComputeBereqHeaderBytes != nil {
		return true
	}

	return false
}

// SetComputeBereqHeaderBytes gets a reference to the given int32 and assigns it to the ComputeBereqHeaderBytes field.
func (o *Results) SetComputeBereqHeaderBytes(v int32) {
	o.ComputeBereqHeaderBytes = &v
}

// GetComputeBereqBodyBytes returns the ComputeBereqBodyBytes field value if set, zero value otherwise.
func (o *Results) GetComputeBereqBodyBytes() int32 {
	if o == nil || o.ComputeBereqBodyBytes == nil {
		var ret int32
		return ret
	}
	return *o.ComputeBereqBodyBytes
}

// GetComputeBereqBodyBytesOk returns a tuple with the ComputeBereqBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetComputeBereqBodyBytesOk() (*int32, bool) {
	if o == nil || o.ComputeBereqBodyBytes == nil {
		return nil, false
	}
	return o.ComputeBereqBodyBytes, true
}

// HasComputeBereqBodyBytes returns a boolean if a field has been set.
func (o *Results) HasComputeBereqBodyBytes() bool {
	if o != nil && o.ComputeBereqBodyBytes != nil {
		return true
	}

	return false
}

// SetComputeBereqBodyBytes gets a reference to the given int32 and assigns it to the ComputeBereqBodyBytes field.
func (o *Results) SetComputeBereqBodyBytes(v int32) {
	o.ComputeBereqBodyBytes = &v
}

// GetComputeBerespHeaderBytes returns the ComputeBerespHeaderBytes field value if set, zero value otherwise.
func (o *Results) GetComputeBerespHeaderBytes() int32 {
	if o == nil || o.ComputeBerespHeaderBytes == nil {
		var ret int32
		return ret
	}
	return *o.ComputeBerespHeaderBytes
}

// GetComputeBerespHeaderBytesOk returns a tuple with the ComputeBerespHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetComputeBerespHeaderBytesOk() (*int32, bool) {
	if o == nil || o.ComputeBerespHeaderBytes == nil {
		return nil, false
	}
	return o.ComputeBerespHeaderBytes, true
}

// HasComputeBerespHeaderBytes returns a boolean if a field has been set.
func (o *Results) HasComputeBerespHeaderBytes() bool {
	if o != nil && o.ComputeBerespHeaderBytes != nil {
		return true
	}

	return false
}

// SetComputeBerespHeaderBytes gets a reference to the given int32 and assigns it to the ComputeBerespHeaderBytes field.
func (o *Results) SetComputeBerespHeaderBytes(v int32) {
	o.ComputeBerespHeaderBytes = &v
}

// GetComputeBerespBodyBytes returns the ComputeBerespBodyBytes field value if set, zero value otherwise.
func (o *Results) GetComputeBerespBodyBytes() int32 {
	if o == nil || o.ComputeBerespBodyBytes == nil {
		var ret int32
		return ret
	}
	return *o.ComputeBerespBodyBytes
}

// GetComputeBerespBodyBytesOk returns a tuple with the ComputeBerespBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetComputeBerespBodyBytesOk() (*int32, bool) {
	if o == nil || o.ComputeBerespBodyBytes == nil {
		return nil, false
	}
	return o.ComputeBerespBodyBytes, true
}

// HasComputeBerespBodyBytes returns a boolean if a field has been set.
func (o *Results) HasComputeBerespBodyBytes() bool {
	if o != nil && o.ComputeBerespBodyBytes != nil {
		return true
	}

	return false
}

// SetComputeBerespBodyBytes gets a reference to the given int32 and assigns it to the ComputeBerespBodyBytes field.
func (o *Results) SetComputeBerespBodyBytes(v int32) {
	o.ComputeBerespBodyBytes = &v
}

// GetComputeBereqs returns the ComputeBereqs field value if set, zero value otherwise.
func (o *Results) GetComputeBereqs() int32 {
	if o == nil || o.ComputeBereqs == nil {
		var ret int32
		return ret
	}
	return *o.ComputeBereqs
}

// GetComputeBereqsOk returns a tuple with the ComputeBereqs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetComputeBereqsOk() (*int32, bool) {
	if o == nil || o.ComputeBereqs == nil {
		return nil, false
	}
	return o.ComputeBereqs, true
}

// HasComputeBereqs returns a boolean if a field has been set.
func (o *Results) HasComputeBereqs() bool {
	if o != nil && o.ComputeBereqs != nil {
		return true
	}

	return false
}

// SetComputeBereqs gets a reference to the given int32 and assigns it to the ComputeBereqs field.
func (o *Results) SetComputeBereqs(v int32) {
	o.ComputeBereqs = &v
}

// GetComputeBereqErrors returns the ComputeBereqErrors field value if set, zero value otherwise.
func (o *Results) GetComputeBereqErrors() int32 {
	if o == nil || o.ComputeBereqErrors == nil {
		var ret int32
		return ret
	}
	return *o.ComputeBereqErrors
}

// GetComputeBereqErrorsOk returns a tuple with the ComputeBereqErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetComputeBereqErrorsOk() (*int32, bool) {
	if o == nil || o.ComputeBereqErrors == nil {
		return nil, false
	}
	return o.ComputeBereqErrors, true
}

// HasComputeBereqErrors returns a boolean if a field has been set.
func (o *Results) HasComputeBereqErrors() bool {
	if o != nil && o.ComputeBereqErrors != nil {
		return true
	}

	return false
}

// SetComputeBereqErrors gets a reference to the given int32 and assigns it to the ComputeBereqErrors field.
func (o *Results) SetComputeBereqErrors(v int32) {
	o.ComputeBereqErrors = &v
}

// GetComputeResourceLimitExceeded returns the ComputeResourceLimitExceeded field value if set, zero value otherwise.
func (o *Results) GetComputeResourceLimitExceeded() int32 {
	if o == nil || o.ComputeResourceLimitExceeded == nil {
		var ret int32
		return ret
	}
	return *o.ComputeResourceLimitExceeded
}

// GetComputeResourceLimitExceededOk returns a tuple with the ComputeResourceLimitExceeded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetComputeResourceLimitExceededOk() (*int32, bool) {
	if o == nil || o.ComputeResourceLimitExceeded == nil {
		return nil, false
	}
	return o.ComputeResourceLimitExceeded, true
}

// HasComputeResourceLimitExceeded returns a boolean if a field has been set.
func (o *Results) HasComputeResourceLimitExceeded() bool {
	if o != nil && o.ComputeResourceLimitExceeded != nil {
		return true
	}

	return false
}

// SetComputeResourceLimitExceeded gets a reference to the given int32 and assigns it to the ComputeResourceLimitExceeded field.
func (o *Results) SetComputeResourceLimitExceeded(v int32) {
	o.ComputeResourceLimitExceeded = &v
}

// GetComputeHeapLimitExceeded returns the ComputeHeapLimitExceeded field value if set, zero value otherwise.
func (o *Results) GetComputeHeapLimitExceeded() int32 {
	if o == nil || o.ComputeHeapLimitExceeded == nil {
		var ret int32
		return ret
	}
	return *o.ComputeHeapLimitExceeded
}

// GetComputeHeapLimitExceededOk returns a tuple with the ComputeHeapLimitExceeded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetComputeHeapLimitExceededOk() (*int32, bool) {
	if o == nil || o.ComputeHeapLimitExceeded == nil {
		return nil, false
	}
	return o.ComputeHeapLimitExceeded, true
}

// HasComputeHeapLimitExceeded returns a boolean if a field has been set.
func (o *Results) HasComputeHeapLimitExceeded() bool {
	if o != nil && o.ComputeHeapLimitExceeded != nil {
		return true
	}

	return false
}

// SetComputeHeapLimitExceeded gets a reference to the given int32 and assigns it to the ComputeHeapLimitExceeded field.
func (o *Results) SetComputeHeapLimitExceeded(v int32) {
	o.ComputeHeapLimitExceeded = &v
}

// GetComputeStackLimitExceeded returns the ComputeStackLimitExceeded field value if set, zero value otherwise.
func (o *Results) GetComputeStackLimitExceeded() int32 {
	if o == nil || o.ComputeStackLimitExceeded == nil {
		var ret int32
		return ret
	}
	return *o.ComputeStackLimitExceeded
}

// GetComputeStackLimitExceededOk returns a tuple with the ComputeStackLimitExceeded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetComputeStackLimitExceededOk() (*int32, bool) {
	if o == nil || o.ComputeStackLimitExceeded == nil {
		return nil, false
	}
	return o.ComputeStackLimitExceeded, true
}

// HasComputeStackLimitExceeded returns a boolean if a field has been set.
func (o *Results) HasComputeStackLimitExceeded() bool {
	if o != nil && o.ComputeStackLimitExceeded != nil {
		return true
	}

	return false
}

// SetComputeStackLimitExceeded gets a reference to the given int32 and assigns it to the ComputeStackLimitExceeded field.
func (o *Results) SetComputeStackLimitExceeded(v int32) {
	o.ComputeStackLimitExceeded = &v
}

// GetComputeGlobalsLimitExceeded returns the ComputeGlobalsLimitExceeded field value if set, zero value otherwise.
func (o *Results) GetComputeGlobalsLimitExceeded() int32 {
	if o == nil || o.ComputeGlobalsLimitExceeded == nil {
		var ret int32
		return ret
	}
	return *o.ComputeGlobalsLimitExceeded
}

// GetComputeGlobalsLimitExceededOk returns a tuple with the ComputeGlobalsLimitExceeded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetComputeGlobalsLimitExceededOk() (*int32, bool) {
	if o == nil || o.ComputeGlobalsLimitExceeded == nil {
		return nil, false
	}
	return o.ComputeGlobalsLimitExceeded, true
}

// HasComputeGlobalsLimitExceeded returns a boolean if a field has been set.
func (o *Results) HasComputeGlobalsLimitExceeded() bool {
	if o != nil && o.ComputeGlobalsLimitExceeded != nil {
		return true
	}

	return false
}

// SetComputeGlobalsLimitExceeded gets a reference to the given int32 and assigns it to the ComputeGlobalsLimitExceeded field.
func (o *Results) SetComputeGlobalsLimitExceeded(v int32) {
	o.ComputeGlobalsLimitExceeded = &v
}

// GetComputeGuestErrors returns the ComputeGuestErrors field value if set, zero value otherwise.
func (o *Results) GetComputeGuestErrors() int32 {
	if o == nil || o.ComputeGuestErrors == nil {
		var ret int32
		return ret
	}
	return *o.ComputeGuestErrors
}

// GetComputeGuestErrorsOk returns a tuple with the ComputeGuestErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetComputeGuestErrorsOk() (*int32, bool) {
	if o == nil || o.ComputeGuestErrors == nil {
		return nil, false
	}
	return o.ComputeGuestErrors, true
}

// HasComputeGuestErrors returns a boolean if a field has been set.
func (o *Results) HasComputeGuestErrors() bool {
	if o != nil && o.ComputeGuestErrors != nil {
		return true
	}

	return false
}

// SetComputeGuestErrors gets a reference to the given int32 and assigns it to the ComputeGuestErrors field.
func (o *Results) SetComputeGuestErrors(v int32) {
	o.ComputeGuestErrors = &v
}

// GetComputeRuntimeErrors returns the ComputeRuntimeErrors field value if set, zero value otherwise.
func (o *Results) GetComputeRuntimeErrors() int32 {
	if o == nil || o.ComputeRuntimeErrors == nil {
		var ret int32
		return ret
	}
	return *o.ComputeRuntimeErrors
}

// GetComputeRuntimeErrorsOk returns a tuple with the ComputeRuntimeErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetComputeRuntimeErrorsOk() (*int32, bool) {
	if o == nil || o.ComputeRuntimeErrors == nil {
		return nil, false
	}
	return o.ComputeRuntimeErrors, true
}

// HasComputeRuntimeErrors returns a boolean if a field has been set.
func (o *Results) HasComputeRuntimeErrors() bool {
	if o != nil && o.ComputeRuntimeErrors != nil {
		return true
	}

	return false
}

// SetComputeRuntimeErrors gets a reference to the given int32 and assigns it to the ComputeRuntimeErrors field.
func (o *Results) SetComputeRuntimeErrors(v int32) {
	o.ComputeRuntimeErrors = &v
}

// GetEdgeHitRespBodyBytes returns the EdgeHitRespBodyBytes field value if set, zero value otherwise.
func (o *Results) GetEdgeHitRespBodyBytes() int32 {
	if o == nil || o.EdgeHitRespBodyBytes == nil {
		var ret int32
		return ret
	}
	return *o.EdgeHitRespBodyBytes
}

// GetEdgeHitRespBodyBytesOk returns a tuple with the EdgeHitRespBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetEdgeHitRespBodyBytesOk() (*int32, bool) {
	if o == nil || o.EdgeHitRespBodyBytes == nil {
		return nil, false
	}
	return o.EdgeHitRespBodyBytes, true
}

// HasEdgeHitRespBodyBytes returns a boolean if a field has been set.
func (o *Results) HasEdgeHitRespBodyBytes() bool {
	if o != nil && o.EdgeHitRespBodyBytes != nil {
		return true
	}

	return false
}

// SetEdgeHitRespBodyBytes gets a reference to the given int32 and assigns it to the EdgeHitRespBodyBytes field.
func (o *Results) SetEdgeHitRespBodyBytes(v int32) {
	o.EdgeHitRespBodyBytes = &v
}

// GetEdgeHitRespHeaderBytes returns the EdgeHitRespHeaderBytes field value if set, zero value otherwise.
func (o *Results) GetEdgeHitRespHeaderBytes() int32 {
	if o == nil || o.EdgeHitRespHeaderBytes == nil {
		var ret int32
		return ret
	}
	return *o.EdgeHitRespHeaderBytes
}

// GetEdgeHitRespHeaderBytesOk returns a tuple with the EdgeHitRespHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetEdgeHitRespHeaderBytesOk() (*int32, bool) {
	if o == nil || o.EdgeHitRespHeaderBytes == nil {
		return nil, false
	}
	return o.EdgeHitRespHeaderBytes, true
}

// HasEdgeHitRespHeaderBytes returns a boolean if a field has been set.
func (o *Results) HasEdgeHitRespHeaderBytes() bool {
	if o != nil && o.EdgeHitRespHeaderBytes != nil {
		return true
	}

	return false
}

// SetEdgeHitRespHeaderBytes gets a reference to the given int32 and assigns it to the EdgeHitRespHeaderBytes field.
func (o *Results) SetEdgeHitRespHeaderBytes(v int32) {
	o.EdgeHitRespHeaderBytes = &v
}

// GetEdgeMissRespBodyBytes returns the EdgeMissRespBodyBytes field value if set, zero value otherwise.
func (o *Results) GetEdgeMissRespBodyBytes() int32 {
	if o == nil || o.EdgeMissRespBodyBytes == nil {
		var ret int32
		return ret
	}
	return *o.EdgeMissRespBodyBytes
}

// GetEdgeMissRespBodyBytesOk returns a tuple with the EdgeMissRespBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetEdgeMissRespBodyBytesOk() (*int32, bool) {
	if o == nil || o.EdgeMissRespBodyBytes == nil {
		return nil, false
	}
	return o.EdgeMissRespBodyBytes, true
}

// HasEdgeMissRespBodyBytes returns a boolean if a field has been set.
func (o *Results) HasEdgeMissRespBodyBytes() bool {
	if o != nil && o.EdgeMissRespBodyBytes != nil {
		return true
	}

	return false
}

// SetEdgeMissRespBodyBytes gets a reference to the given int32 and assigns it to the EdgeMissRespBodyBytes field.
func (o *Results) SetEdgeMissRespBodyBytes(v int32) {
	o.EdgeMissRespBodyBytes = &v
}

// GetEdgeMissRespHeaderBytes returns the EdgeMissRespHeaderBytes field value if set, zero value otherwise.
func (o *Results) GetEdgeMissRespHeaderBytes() int32 {
	if o == nil || o.EdgeMissRespHeaderBytes == nil {
		var ret int32
		return ret
	}
	return *o.EdgeMissRespHeaderBytes
}

// GetEdgeMissRespHeaderBytesOk returns a tuple with the EdgeMissRespHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetEdgeMissRespHeaderBytesOk() (*int32, bool) {
	if o == nil || o.EdgeMissRespHeaderBytes == nil {
		return nil, false
	}
	return o.EdgeMissRespHeaderBytes, true
}

// HasEdgeMissRespHeaderBytes returns a boolean if a field has been set.
func (o *Results) HasEdgeMissRespHeaderBytes() bool {
	if o != nil && o.EdgeMissRespHeaderBytes != nil {
		return true
	}

	return false
}

// SetEdgeMissRespHeaderBytes gets a reference to the given int32 and assigns it to the EdgeMissRespHeaderBytes field.
func (o *Results) SetEdgeMissRespHeaderBytes(v int32) {
	o.EdgeMissRespHeaderBytes = &v
}

// GetOriginCacheFetchRespBodyBytes returns the OriginCacheFetchRespBodyBytes field value if set, zero value otherwise.
func (o *Results) GetOriginCacheFetchRespBodyBytes() int32 {
	if o == nil || o.OriginCacheFetchRespBodyBytes == nil {
		var ret int32
		return ret
	}
	return *o.OriginCacheFetchRespBodyBytes
}

// GetOriginCacheFetchRespBodyBytesOk returns a tuple with the OriginCacheFetchRespBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetOriginCacheFetchRespBodyBytesOk() (*int32, bool) {
	if o == nil || o.OriginCacheFetchRespBodyBytes == nil {
		return nil, false
	}
	return o.OriginCacheFetchRespBodyBytes, true
}

// HasOriginCacheFetchRespBodyBytes returns a boolean if a field has been set.
func (o *Results) HasOriginCacheFetchRespBodyBytes() bool {
	if o != nil && o.OriginCacheFetchRespBodyBytes != nil {
		return true
	}

	return false
}

// SetOriginCacheFetchRespBodyBytes gets a reference to the given int32 and assigns it to the OriginCacheFetchRespBodyBytes field.
func (o *Results) SetOriginCacheFetchRespBodyBytes(v int32) {
	o.OriginCacheFetchRespBodyBytes = &v
}

// GetOriginCacheFetchRespHeaderBytes returns the OriginCacheFetchRespHeaderBytes field value if set, zero value otherwise.
func (o *Results) GetOriginCacheFetchRespHeaderBytes() int32 {
	if o == nil || o.OriginCacheFetchRespHeaderBytes == nil {
		var ret int32
		return ret
	}
	return *o.OriginCacheFetchRespHeaderBytes
}

// GetOriginCacheFetchRespHeaderBytesOk returns a tuple with the OriginCacheFetchRespHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetOriginCacheFetchRespHeaderBytesOk() (*int32, bool) {
	if o == nil || o.OriginCacheFetchRespHeaderBytes == nil {
		return nil, false
	}
	return o.OriginCacheFetchRespHeaderBytes, true
}

// HasOriginCacheFetchRespHeaderBytes returns a boolean if a field has been set.
func (o *Results) HasOriginCacheFetchRespHeaderBytes() bool {
	if o != nil && o.OriginCacheFetchRespHeaderBytes != nil {
		return true
	}

	return false
}

// SetOriginCacheFetchRespHeaderBytes gets a reference to the given int32 and assigns it to the OriginCacheFetchRespHeaderBytes field.
func (o *Results) SetOriginCacheFetchRespHeaderBytes(v int32) {
	o.OriginCacheFetchRespHeaderBytes = &v
}

// GetShieldHitRequests returns the ShieldHitRequests field value if set, zero value otherwise.
func (o *Results) GetShieldHitRequests() int32 {
	if o == nil || o.ShieldHitRequests == nil {
		var ret int32
		return ret
	}
	return *o.ShieldHitRequests
}

// GetShieldHitRequestsOk returns a tuple with the ShieldHitRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetShieldHitRequestsOk() (*int32, bool) {
	if o == nil || o.ShieldHitRequests == nil {
		return nil, false
	}
	return o.ShieldHitRequests, true
}

// HasShieldHitRequests returns a boolean if a field has been set.
func (o *Results) HasShieldHitRequests() bool {
	if o != nil && o.ShieldHitRequests != nil {
		return true
	}

	return false
}

// SetShieldHitRequests gets a reference to the given int32 and assigns it to the ShieldHitRequests field.
func (o *Results) SetShieldHitRequests(v int32) {
	o.ShieldHitRequests = &v
}

// GetShieldMissRequests returns the ShieldMissRequests field value if set, zero value otherwise.
func (o *Results) GetShieldMissRequests() int32 {
	if o == nil || o.ShieldMissRequests == nil {
		var ret int32
		return ret
	}
	return *o.ShieldMissRequests
}

// GetShieldMissRequestsOk returns a tuple with the ShieldMissRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetShieldMissRequestsOk() (*int32, bool) {
	if o == nil || o.ShieldMissRequests == nil {
		return nil, false
	}
	return o.ShieldMissRequests, true
}

// HasShieldMissRequests returns a boolean if a field has been set.
func (o *Results) HasShieldMissRequests() bool {
	if o != nil && o.ShieldMissRequests != nil {
		return true
	}

	return false
}

// SetShieldMissRequests gets a reference to the given int32 and assigns it to the ShieldMissRequests field.
func (o *Results) SetShieldMissRequests(v int32) {
	o.ShieldMissRequests = &v
}

// GetShieldHitRespHeaderBytes returns the ShieldHitRespHeaderBytes field value if set, zero value otherwise.
func (o *Results) GetShieldHitRespHeaderBytes() int32 {
	if o == nil || o.ShieldHitRespHeaderBytes == nil {
		var ret int32
		return ret
	}
	return *o.ShieldHitRespHeaderBytes
}

// GetShieldHitRespHeaderBytesOk returns a tuple with the ShieldHitRespHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetShieldHitRespHeaderBytesOk() (*int32, bool) {
	if o == nil || o.ShieldHitRespHeaderBytes == nil {
		return nil, false
	}
	return o.ShieldHitRespHeaderBytes, true
}

// HasShieldHitRespHeaderBytes returns a boolean if a field has been set.
func (o *Results) HasShieldHitRespHeaderBytes() bool {
	if o != nil && o.ShieldHitRespHeaderBytes != nil {
		return true
	}

	return false
}

// SetShieldHitRespHeaderBytes gets a reference to the given int32 and assigns it to the ShieldHitRespHeaderBytes field.
func (o *Results) SetShieldHitRespHeaderBytes(v int32) {
	o.ShieldHitRespHeaderBytes = &v
}

// GetShieldHitRespBodyBytes returns the ShieldHitRespBodyBytes field value if set, zero value otherwise.
func (o *Results) GetShieldHitRespBodyBytes() int32 {
	if o == nil || o.ShieldHitRespBodyBytes == nil {
		var ret int32
		return ret
	}
	return *o.ShieldHitRespBodyBytes
}

// GetShieldHitRespBodyBytesOk returns a tuple with the ShieldHitRespBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetShieldHitRespBodyBytesOk() (*int32, bool) {
	if o == nil || o.ShieldHitRespBodyBytes == nil {
		return nil, false
	}
	return o.ShieldHitRespBodyBytes, true
}

// HasShieldHitRespBodyBytes returns a boolean if a field has been set.
func (o *Results) HasShieldHitRespBodyBytes() bool {
	if o != nil && o.ShieldHitRespBodyBytes != nil {
		return true
	}

	return false
}

// SetShieldHitRespBodyBytes gets a reference to the given int32 and assigns it to the ShieldHitRespBodyBytes field.
func (o *Results) SetShieldHitRespBodyBytes(v int32) {
	o.ShieldHitRespBodyBytes = &v
}

// GetShieldMissRespHeaderBytes returns the ShieldMissRespHeaderBytes field value if set, zero value otherwise.
func (o *Results) GetShieldMissRespHeaderBytes() int32 {
	if o == nil || o.ShieldMissRespHeaderBytes == nil {
		var ret int32
		return ret
	}
	return *o.ShieldMissRespHeaderBytes
}

// GetShieldMissRespHeaderBytesOk returns a tuple with the ShieldMissRespHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetShieldMissRespHeaderBytesOk() (*int32, bool) {
	if o == nil || o.ShieldMissRespHeaderBytes == nil {
		return nil, false
	}
	return o.ShieldMissRespHeaderBytes, true
}

// HasShieldMissRespHeaderBytes returns a boolean if a field has been set.
func (o *Results) HasShieldMissRespHeaderBytes() bool {
	if o != nil && o.ShieldMissRespHeaderBytes != nil {
		return true
	}

	return false
}

// SetShieldMissRespHeaderBytes gets a reference to the given int32 and assigns it to the ShieldMissRespHeaderBytes field.
func (o *Results) SetShieldMissRespHeaderBytes(v int32) {
	o.ShieldMissRespHeaderBytes = &v
}

// GetShieldMissRespBodyBytes returns the ShieldMissRespBodyBytes field value if set, zero value otherwise.
func (o *Results) GetShieldMissRespBodyBytes() int32 {
	if o == nil || o.ShieldMissRespBodyBytes == nil {
		var ret int32
		return ret
	}
	return *o.ShieldMissRespBodyBytes
}

// GetShieldMissRespBodyBytesOk returns a tuple with the ShieldMissRespBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetShieldMissRespBodyBytesOk() (*int32, bool) {
	if o == nil || o.ShieldMissRespBodyBytes == nil {
		return nil, false
	}
	return o.ShieldMissRespBodyBytes, true
}

// HasShieldMissRespBodyBytes returns a boolean if a field has been set.
func (o *Results) HasShieldMissRespBodyBytes() bool {
	if o != nil && o.ShieldMissRespBodyBytes != nil {
		return true
	}

	return false
}

// SetShieldMissRespBodyBytes gets a reference to the given int32 and assigns it to the ShieldMissRespBodyBytes field.
func (o *Results) SetShieldMissRespBodyBytes(v int32) {
	o.ShieldMissRespBodyBytes = &v
}

// GetWebsocketReqHeaderBytes returns the WebsocketReqHeaderBytes field value if set, zero value otherwise.
func (o *Results) GetWebsocketReqHeaderBytes() int32 {
	if o == nil || o.WebsocketReqHeaderBytes == nil {
		var ret int32
		return ret
	}
	return *o.WebsocketReqHeaderBytes
}

// GetWebsocketReqHeaderBytesOk returns a tuple with the WebsocketReqHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetWebsocketReqHeaderBytesOk() (*int32, bool) {
	if o == nil || o.WebsocketReqHeaderBytes == nil {
		return nil, false
	}
	return o.WebsocketReqHeaderBytes, true
}

// HasWebsocketReqHeaderBytes returns a boolean if a field has been set.
func (o *Results) HasWebsocketReqHeaderBytes() bool {
	if o != nil && o.WebsocketReqHeaderBytes != nil {
		return true
	}

	return false
}

// SetWebsocketReqHeaderBytes gets a reference to the given int32 and assigns it to the WebsocketReqHeaderBytes field.
func (o *Results) SetWebsocketReqHeaderBytes(v int32) {
	o.WebsocketReqHeaderBytes = &v
}

// GetWebsocketReqBodyBytes returns the WebsocketReqBodyBytes field value if set, zero value otherwise.
func (o *Results) GetWebsocketReqBodyBytes() int32 {
	if o == nil || o.WebsocketReqBodyBytes == nil {
		var ret int32
		return ret
	}
	return *o.WebsocketReqBodyBytes
}

// GetWebsocketReqBodyBytesOk returns a tuple with the WebsocketReqBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetWebsocketReqBodyBytesOk() (*int32, bool) {
	if o == nil || o.WebsocketReqBodyBytes == nil {
		return nil, false
	}
	return o.WebsocketReqBodyBytes, true
}

// HasWebsocketReqBodyBytes returns a boolean if a field has been set.
func (o *Results) HasWebsocketReqBodyBytes() bool {
	if o != nil && o.WebsocketReqBodyBytes != nil {
		return true
	}

	return false
}

// SetWebsocketReqBodyBytes gets a reference to the given int32 and assigns it to the WebsocketReqBodyBytes field.
func (o *Results) SetWebsocketReqBodyBytes(v int32) {
	o.WebsocketReqBodyBytes = &v
}

// GetWebsocketRespHeaderBytes returns the WebsocketRespHeaderBytes field value if set, zero value otherwise.
func (o *Results) GetWebsocketRespHeaderBytes() int32 {
	if o == nil || o.WebsocketRespHeaderBytes == nil {
		var ret int32
		return ret
	}
	return *o.WebsocketRespHeaderBytes
}

// GetWebsocketRespHeaderBytesOk returns a tuple with the WebsocketRespHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetWebsocketRespHeaderBytesOk() (*int32, bool) {
	if o == nil || o.WebsocketRespHeaderBytes == nil {
		return nil, false
	}
	return o.WebsocketRespHeaderBytes, true
}

// HasWebsocketRespHeaderBytes returns a boolean if a field has been set.
func (o *Results) HasWebsocketRespHeaderBytes() bool {
	if o != nil && o.WebsocketRespHeaderBytes != nil {
		return true
	}

	return false
}

// SetWebsocketRespHeaderBytes gets a reference to the given int32 and assigns it to the WebsocketRespHeaderBytes field.
func (o *Results) SetWebsocketRespHeaderBytes(v int32) {
	o.WebsocketRespHeaderBytes = &v
}

// GetWebsocketRespBodyBytes returns the WebsocketRespBodyBytes field value if set, zero value otherwise.
func (o *Results) GetWebsocketRespBodyBytes() int32 {
	if o == nil || o.WebsocketRespBodyBytes == nil {
		var ret int32
		return ret
	}
	return *o.WebsocketRespBodyBytes
}

// GetWebsocketRespBodyBytesOk returns a tuple with the WebsocketRespBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetWebsocketRespBodyBytesOk() (*int32, bool) {
	if o == nil || o.WebsocketRespBodyBytes == nil {
		return nil, false
	}
	return o.WebsocketRespBodyBytes, true
}

// HasWebsocketRespBodyBytes returns a boolean if a field has been set.
func (o *Results) HasWebsocketRespBodyBytes() bool {
	if o != nil && o.WebsocketRespBodyBytes != nil {
		return true
	}

	return false
}

// SetWebsocketRespBodyBytes gets a reference to the given int32 and assigns it to the WebsocketRespBodyBytes field.
func (o *Results) SetWebsocketRespBodyBytes(v int32) {
	o.WebsocketRespBodyBytes = &v
}

// GetWebsocketBereqHeaderBytes returns the WebsocketBereqHeaderBytes field value if set, zero value otherwise.
func (o *Results) GetWebsocketBereqHeaderBytes() int32 {
	if o == nil || o.WebsocketBereqHeaderBytes == nil {
		var ret int32
		return ret
	}
	return *o.WebsocketBereqHeaderBytes
}

// GetWebsocketBereqHeaderBytesOk returns a tuple with the WebsocketBereqHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetWebsocketBereqHeaderBytesOk() (*int32, bool) {
	if o == nil || o.WebsocketBereqHeaderBytes == nil {
		return nil, false
	}
	return o.WebsocketBereqHeaderBytes, true
}

// HasWebsocketBereqHeaderBytes returns a boolean if a field has been set.
func (o *Results) HasWebsocketBereqHeaderBytes() bool {
	if o != nil && o.WebsocketBereqHeaderBytes != nil {
		return true
	}

	return false
}

// SetWebsocketBereqHeaderBytes gets a reference to the given int32 and assigns it to the WebsocketBereqHeaderBytes field.
func (o *Results) SetWebsocketBereqHeaderBytes(v int32) {
	o.WebsocketBereqHeaderBytes = &v
}

// GetWebsocketBereqBodyBytes returns the WebsocketBereqBodyBytes field value if set, zero value otherwise.
func (o *Results) GetWebsocketBereqBodyBytes() int32 {
	if o == nil || o.WebsocketBereqBodyBytes == nil {
		var ret int32
		return ret
	}
	return *o.WebsocketBereqBodyBytes
}

// GetWebsocketBereqBodyBytesOk returns a tuple with the WebsocketBereqBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetWebsocketBereqBodyBytesOk() (*int32, bool) {
	if o == nil || o.WebsocketBereqBodyBytes == nil {
		return nil, false
	}
	return o.WebsocketBereqBodyBytes, true
}

// HasWebsocketBereqBodyBytes returns a boolean if a field has been set.
func (o *Results) HasWebsocketBereqBodyBytes() bool {
	if o != nil && o.WebsocketBereqBodyBytes != nil {
		return true
	}

	return false
}

// SetWebsocketBereqBodyBytes gets a reference to the given int32 and assigns it to the WebsocketBereqBodyBytes field.
func (o *Results) SetWebsocketBereqBodyBytes(v int32) {
	o.WebsocketBereqBodyBytes = &v
}

// GetWebsocketBerespHeaderBytes returns the WebsocketBerespHeaderBytes field value if set, zero value otherwise.
func (o *Results) GetWebsocketBerespHeaderBytes() int32 {
	if o == nil || o.WebsocketBerespHeaderBytes == nil {
		var ret int32
		return ret
	}
	return *o.WebsocketBerespHeaderBytes
}

// GetWebsocketBerespHeaderBytesOk returns a tuple with the WebsocketBerespHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetWebsocketBerespHeaderBytesOk() (*int32, bool) {
	if o == nil || o.WebsocketBerespHeaderBytes == nil {
		return nil, false
	}
	return o.WebsocketBerespHeaderBytes, true
}

// HasWebsocketBerespHeaderBytes returns a boolean if a field has been set.
func (o *Results) HasWebsocketBerespHeaderBytes() bool {
	if o != nil && o.WebsocketBerespHeaderBytes != nil {
		return true
	}

	return false
}

// SetWebsocketBerespHeaderBytes gets a reference to the given int32 and assigns it to the WebsocketBerespHeaderBytes field.
func (o *Results) SetWebsocketBerespHeaderBytes(v int32) {
	o.WebsocketBerespHeaderBytes = &v
}

// GetWebsocketBerespBodyBytes returns the WebsocketBerespBodyBytes field value if set, zero value otherwise.
func (o *Results) GetWebsocketBerespBodyBytes() int32 {
	if o == nil || o.WebsocketBerespBodyBytes == nil {
		var ret int32
		return ret
	}
	return *o.WebsocketBerespBodyBytes
}

// GetWebsocketBerespBodyBytesOk returns a tuple with the WebsocketBerespBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetWebsocketBerespBodyBytesOk() (*int32, bool) {
	if o == nil || o.WebsocketBerespBodyBytes == nil {
		return nil, false
	}
	return o.WebsocketBerespBodyBytes, true
}

// HasWebsocketBerespBodyBytes returns a boolean if a field has been set.
func (o *Results) HasWebsocketBerespBodyBytes() bool {
	if o != nil && o.WebsocketBerespBodyBytes != nil {
		return true
	}

	return false
}

// SetWebsocketBerespBodyBytes gets a reference to the given int32 and assigns it to the WebsocketBerespBodyBytes field.
func (o *Results) SetWebsocketBerespBodyBytes(v int32) {
	o.WebsocketBerespBodyBytes = &v
}

// GetWebsocketConnTimeMs returns the WebsocketConnTimeMs field value if set, zero value otherwise.
func (o *Results) GetWebsocketConnTimeMs() int32 {
	if o == nil || o.WebsocketConnTimeMs == nil {
		var ret int32
		return ret
	}
	return *o.WebsocketConnTimeMs
}

// GetWebsocketConnTimeMsOk returns a tuple with the WebsocketConnTimeMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetWebsocketConnTimeMsOk() (*int32, bool) {
	if o == nil || o.WebsocketConnTimeMs == nil {
		return nil, false
	}
	return o.WebsocketConnTimeMs, true
}

// HasWebsocketConnTimeMs returns a boolean if a field has been set.
func (o *Results) HasWebsocketConnTimeMs() bool {
	if o != nil && o.WebsocketConnTimeMs != nil {
		return true
	}

	return false
}

// SetWebsocketConnTimeMs gets a reference to the given int32 and assigns it to the WebsocketConnTimeMs field.
func (o *Results) SetWebsocketConnTimeMs(v int32) {
	o.WebsocketConnTimeMs = &v
}

// GetFanoutRecvPublishes returns the FanoutRecvPublishes field value if set, zero value otherwise.
func (o *Results) GetFanoutRecvPublishes() int32 {
	if o == nil || o.FanoutRecvPublishes == nil {
		var ret int32
		return ret
	}
	return *o.FanoutRecvPublishes
}

// GetFanoutRecvPublishesOk returns a tuple with the FanoutRecvPublishes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetFanoutRecvPublishesOk() (*int32, bool) {
	if o == nil || o.FanoutRecvPublishes == nil {
		return nil, false
	}
	return o.FanoutRecvPublishes, true
}

// HasFanoutRecvPublishes returns a boolean if a field has been set.
func (o *Results) HasFanoutRecvPublishes() bool {
	if o != nil && o.FanoutRecvPublishes != nil {
		return true
	}

	return false
}

// SetFanoutRecvPublishes gets a reference to the given int32 and assigns it to the FanoutRecvPublishes field.
func (o *Results) SetFanoutRecvPublishes(v int32) {
	o.FanoutRecvPublishes = &v
}

// GetFanoutSendPublishes returns the FanoutSendPublishes field value if set, zero value otherwise.
func (o *Results) GetFanoutSendPublishes() int32 {
	if o == nil || o.FanoutSendPublishes == nil {
		var ret int32
		return ret
	}
	return *o.FanoutSendPublishes
}

// GetFanoutSendPublishesOk returns a tuple with the FanoutSendPublishes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetFanoutSendPublishesOk() (*int32, bool) {
	if o == nil || o.FanoutSendPublishes == nil {
		return nil, false
	}
	return o.FanoutSendPublishes, true
}

// HasFanoutSendPublishes returns a boolean if a field has been set.
func (o *Results) HasFanoutSendPublishes() bool {
	if o != nil && o.FanoutSendPublishes != nil {
		return true
	}

	return false
}

// SetFanoutSendPublishes gets a reference to the given int32 and assigns it to the FanoutSendPublishes field.
func (o *Results) SetFanoutSendPublishes(v int32) {
	o.FanoutSendPublishes = &v
}

// GetKvStoreClassAOperations returns the KvStoreClassAOperations field value if set, zero value otherwise.
func (o *Results) GetKvStoreClassAOperations() int32 {
	if o == nil || o.KvStoreClassAOperations == nil {
		var ret int32
		return ret
	}
	return *o.KvStoreClassAOperations
}

// GetKvStoreClassAOperationsOk returns a tuple with the KvStoreClassAOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetKvStoreClassAOperationsOk() (*int32, bool) {
	if o == nil || o.KvStoreClassAOperations == nil {
		return nil, false
	}
	return o.KvStoreClassAOperations, true
}

// HasKvStoreClassAOperations returns a boolean if a field has been set.
func (o *Results) HasKvStoreClassAOperations() bool {
	if o != nil && o.KvStoreClassAOperations != nil {
		return true
	}

	return false
}

// SetKvStoreClassAOperations gets a reference to the given int32 and assigns it to the KvStoreClassAOperations field.
func (o *Results) SetKvStoreClassAOperations(v int32) {
	o.KvStoreClassAOperations = &v
}

// GetKvStoreClassBOperations returns the KvStoreClassBOperations field value if set, zero value otherwise.
func (o *Results) GetKvStoreClassBOperations() int32 {
	if o == nil || o.KvStoreClassBOperations == nil {
		var ret int32
		return ret
	}
	return *o.KvStoreClassBOperations
}

// GetKvStoreClassBOperationsOk returns a tuple with the KvStoreClassBOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetKvStoreClassBOperationsOk() (*int32, bool) {
	if o == nil || o.KvStoreClassBOperations == nil {
		return nil, false
	}
	return o.KvStoreClassBOperations, true
}

// HasKvStoreClassBOperations returns a boolean if a field has been set.
func (o *Results) HasKvStoreClassBOperations() bool {
	if o != nil && o.KvStoreClassBOperations != nil {
		return true
	}

	return false
}

// SetKvStoreClassBOperations gets a reference to the given int32 and assigns it to the KvStoreClassBOperations field.
func (o *Results) SetKvStoreClassBOperations(v int32) {
	o.KvStoreClassBOperations = &v
}

// GetObjectStoreClassAOperations returns the ObjectStoreClassAOperations field value if set, zero value otherwise.
// Deprecated
func (o *Results) GetObjectStoreClassAOperations() int32 {
	if o == nil || o.ObjectStoreClassAOperations == nil {
		var ret int32
		return ret
	}
	return *o.ObjectStoreClassAOperations
}

// GetObjectStoreClassAOperationsOk returns a tuple with the ObjectStoreClassAOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *Results) GetObjectStoreClassAOperationsOk() (*int32, bool) {
	if o == nil || o.ObjectStoreClassAOperations == nil {
		return nil, false
	}
	return o.ObjectStoreClassAOperations, true
}

// HasObjectStoreClassAOperations returns a boolean if a field has been set.
func (o *Results) HasObjectStoreClassAOperations() bool {
	if o != nil && o.ObjectStoreClassAOperations != nil {
		return true
	}

	return false
}

// SetObjectStoreClassAOperations gets a reference to the given int32 and assigns it to the ObjectStoreClassAOperations field.
// Deprecated
func (o *Results) SetObjectStoreClassAOperations(v int32) {
	o.ObjectStoreClassAOperations = &v
}

// GetObjectStoreClassBOperations returns the ObjectStoreClassBOperations field value if set, zero value otherwise.
// Deprecated
func (o *Results) GetObjectStoreClassBOperations() int32 {
	if o == nil || o.ObjectStoreClassBOperations == nil {
		var ret int32
		return ret
	}
	return *o.ObjectStoreClassBOperations
}

// GetObjectStoreClassBOperationsOk returns a tuple with the ObjectStoreClassBOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *Results) GetObjectStoreClassBOperationsOk() (*int32, bool) {
	if o == nil || o.ObjectStoreClassBOperations == nil {
		return nil, false
	}
	return o.ObjectStoreClassBOperations, true
}

// HasObjectStoreClassBOperations returns a boolean if a field has been set.
func (o *Results) HasObjectStoreClassBOperations() bool {
	if o != nil && o.ObjectStoreClassBOperations != nil {
		return true
	}

	return false
}

// SetObjectStoreClassBOperations gets a reference to the given int32 and assigns it to the ObjectStoreClassBOperations field.
// Deprecated
func (o *Results) SetObjectStoreClassBOperations(v int32) {
	o.ObjectStoreClassBOperations = &v
}

// GetFanoutReqHeaderBytes returns the FanoutReqHeaderBytes field value if set, zero value otherwise.
func (o *Results) GetFanoutReqHeaderBytes() int32 {
	if o == nil || o.FanoutReqHeaderBytes == nil {
		var ret int32
		return ret
	}
	return *o.FanoutReqHeaderBytes
}

// GetFanoutReqHeaderBytesOk returns a tuple with the FanoutReqHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetFanoutReqHeaderBytesOk() (*int32, bool) {
	if o == nil || o.FanoutReqHeaderBytes == nil {
		return nil, false
	}
	return o.FanoutReqHeaderBytes, true
}

// HasFanoutReqHeaderBytes returns a boolean if a field has been set.
func (o *Results) HasFanoutReqHeaderBytes() bool {
	if o != nil && o.FanoutReqHeaderBytes != nil {
		return true
	}

	return false
}

// SetFanoutReqHeaderBytes gets a reference to the given int32 and assigns it to the FanoutReqHeaderBytes field.
func (o *Results) SetFanoutReqHeaderBytes(v int32) {
	o.FanoutReqHeaderBytes = &v
}

// GetFanoutReqBodyBytes returns the FanoutReqBodyBytes field value if set, zero value otherwise.
func (o *Results) GetFanoutReqBodyBytes() int32 {
	if o == nil || o.FanoutReqBodyBytes == nil {
		var ret int32
		return ret
	}
	return *o.FanoutReqBodyBytes
}

// GetFanoutReqBodyBytesOk returns a tuple with the FanoutReqBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetFanoutReqBodyBytesOk() (*int32, bool) {
	if o == nil || o.FanoutReqBodyBytes == nil {
		return nil, false
	}
	return o.FanoutReqBodyBytes, true
}

// HasFanoutReqBodyBytes returns a boolean if a field has been set.
func (o *Results) HasFanoutReqBodyBytes() bool {
	if o != nil && o.FanoutReqBodyBytes != nil {
		return true
	}

	return false
}

// SetFanoutReqBodyBytes gets a reference to the given int32 and assigns it to the FanoutReqBodyBytes field.
func (o *Results) SetFanoutReqBodyBytes(v int32) {
	o.FanoutReqBodyBytes = &v
}

// GetFanoutRespHeaderBytes returns the FanoutRespHeaderBytes field value if set, zero value otherwise.
func (o *Results) GetFanoutRespHeaderBytes() int32 {
	if o == nil || o.FanoutRespHeaderBytes == nil {
		var ret int32
		return ret
	}
	return *o.FanoutRespHeaderBytes
}

// GetFanoutRespHeaderBytesOk returns a tuple with the FanoutRespHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetFanoutRespHeaderBytesOk() (*int32, bool) {
	if o == nil || o.FanoutRespHeaderBytes == nil {
		return nil, false
	}
	return o.FanoutRespHeaderBytes, true
}

// HasFanoutRespHeaderBytes returns a boolean if a field has been set.
func (o *Results) HasFanoutRespHeaderBytes() bool {
	if o != nil && o.FanoutRespHeaderBytes != nil {
		return true
	}

	return false
}

// SetFanoutRespHeaderBytes gets a reference to the given int32 and assigns it to the FanoutRespHeaderBytes field.
func (o *Results) SetFanoutRespHeaderBytes(v int32) {
	o.FanoutRespHeaderBytes = &v
}

// GetFanoutRespBodyBytes returns the FanoutRespBodyBytes field value if set, zero value otherwise.
func (o *Results) GetFanoutRespBodyBytes() int32 {
	if o == nil || o.FanoutRespBodyBytes == nil {
		var ret int32
		return ret
	}
	return *o.FanoutRespBodyBytes
}

// GetFanoutRespBodyBytesOk returns a tuple with the FanoutRespBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetFanoutRespBodyBytesOk() (*int32, bool) {
	if o == nil || o.FanoutRespBodyBytes == nil {
		return nil, false
	}
	return o.FanoutRespBodyBytes, true
}

// HasFanoutRespBodyBytes returns a boolean if a field has been set.
func (o *Results) HasFanoutRespBodyBytes() bool {
	if o != nil && o.FanoutRespBodyBytes != nil {
		return true
	}

	return false
}

// SetFanoutRespBodyBytes gets a reference to the given int32 and assigns it to the FanoutRespBodyBytes field.
func (o *Results) SetFanoutRespBodyBytes(v int32) {
	o.FanoutRespBodyBytes = &v
}

// GetFanoutBereqHeaderBytes returns the FanoutBereqHeaderBytes field value if set, zero value otherwise.
func (o *Results) GetFanoutBereqHeaderBytes() int32 {
	if o == nil || o.FanoutBereqHeaderBytes == nil {
		var ret int32
		return ret
	}
	return *o.FanoutBereqHeaderBytes
}

// GetFanoutBereqHeaderBytesOk returns a tuple with the FanoutBereqHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetFanoutBereqHeaderBytesOk() (*int32, bool) {
	if o == nil || o.FanoutBereqHeaderBytes == nil {
		return nil, false
	}
	return o.FanoutBereqHeaderBytes, true
}

// HasFanoutBereqHeaderBytes returns a boolean if a field has been set.
func (o *Results) HasFanoutBereqHeaderBytes() bool {
	if o != nil && o.FanoutBereqHeaderBytes != nil {
		return true
	}

	return false
}

// SetFanoutBereqHeaderBytes gets a reference to the given int32 and assigns it to the FanoutBereqHeaderBytes field.
func (o *Results) SetFanoutBereqHeaderBytes(v int32) {
	o.FanoutBereqHeaderBytes = &v
}

// GetFanoutBereqBodyBytes returns the FanoutBereqBodyBytes field value if set, zero value otherwise.
func (o *Results) GetFanoutBereqBodyBytes() int32 {
	if o == nil || o.FanoutBereqBodyBytes == nil {
		var ret int32
		return ret
	}
	return *o.FanoutBereqBodyBytes
}

// GetFanoutBereqBodyBytesOk returns a tuple with the FanoutBereqBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetFanoutBereqBodyBytesOk() (*int32, bool) {
	if o == nil || o.FanoutBereqBodyBytes == nil {
		return nil, false
	}
	return o.FanoutBereqBodyBytes, true
}

// HasFanoutBereqBodyBytes returns a boolean if a field has been set.
func (o *Results) HasFanoutBereqBodyBytes() bool {
	if o != nil && o.FanoutBereqBodyBytes != nil {
		return true
	}

	return false
}

// SetFanoutBereqBodyBytes gets a reference to the given int32 and assigns it to the FanoutBereqBodyBytes field.
func (o *Results) SetFanoutBereqBodyBytes(v int32) {
	o.FanoutBereqBodyBytes = &v
}

// GetFanoutBerespHeaderBytes returns the FanoutBerespHeaderBytes field value if set, zero value otherwise.
func (o *Results) GetFanoutBerespHeaderBytes() int32 {
	if o == nil || o.FanoutBerespHeaderBytes == nil {
		var ret int32
		return ret
	}
	return *o.FanoutBerespHeaderBytes
}

// GetFanoutBerespHeaderBytesOk returns a tuple with the FanoutBerespHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetFanoutBerespHeaderBytesOk() (*int32, bool) {
	if o == nil || o.FanoutBerespHeaderBytes == nil {
		return nil, false
	}
	return o.FanoutBerespHeaderBytes, true
}

// HasFanoutBerespHeaderBytes returns a boolean if a field has been set.
func (o *Results) HasFanoutBerespHeaderBytes() bool {
	if o != nil && o.FanoutBerespHeaderBytes != nil {
		return true
	}

	return false
}

// SetFanoutBerespHeaderBytes gets a reference to the given int32 and assigns it to the FanoutBerespHeaderBytes field.
func (o *Results) SetFanoutBerespHeaderBytes(v int32) {
	o.FanoutBerespHeaderBytes = &v
}

// GetFanoutBerespBodyBytes returns the FanoutBerespBodyBytes field value if set, zero value otherwise.
func (o *Results) GetFanoutBerespBodyBytes() int32 {
	if o == nil || o.FanoutBerespBodyBytes == nil {
		var ret int32
		return ret
	}
	return *o.FanoutBerespBodyBytes
}

// GetFanoutBerespBodyBytesOk returns a tuple with the FanoutBerespBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetFanoutBerespBodyBytesOk() (*int32, bool) {
	if o == nil || o.FanoutBerespBodyBytes == nil {
		return nil, false
	}
	return o.FanoutBerespBodyBytes, true
}

// HasFanoutBerespBodyBytes returns a boolean if a field has been set.
func (o *Results) HasFanoutBerespBodyBytes() bool {
	if o != nil && o.FanoutBerespBodyBytes != nil {
		return true
	}

	return false
}

// SetFanoutBerespBodyBytes gets a reference to the given int32 and assigns it to the FanoutBerespBodyBytes field.
func (o *Results) SetFanoutBerespBodyBytes(v int32) {
	o.FanoutBerespBodyBytes = &v
}

// GetFanoutConnTimeMs returns the FanoutConnTimeMs field value if set, zero value otherwise.
func (o *Results) GetFanoutConnTimeMs() int32 {
	if o == nil || o.FanoutConnTimeMs == nil {
		var ret int32
		return ret
	}
	return *o.FanoutConnTimeMs
}

// GetFanoutConnTimeMsOk returns a tuple with the FanoutConnTimeMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetFanoutConnTimeMsOk() (*int32, bool) {
	if o == nil || o.FanoutConnTimeMs == nil {
		return nil, false
	}
	return o.FanoutConnTimeMs, true
}

// HasFanoutConnTimeMs returns a boolean if a field has been set.
func (o *Results) HasFanoutConnTimeMs() bool {
	if o != nil && o.FanoutConnTimeMs != nil {
		return true
	}

	return false
}

// SetFanoutConnTimeMs gets a reference to the given int32 and assigns it to the FanoutConnTimeMs field.
func (o *Results) SetFanoutConnTimeMs(v int32) {
	o.FanoutConnTimeMs = &v
}

// GetDdosActionLimitStreamsConnections returns the DdosActionLimitStreamsConnections field value if set, zero value otherwise.
func (o *Results) GetDdosActionLimitStreamsConnections() int32 {
	if o == nil || o.DdosActionLimitStreamsConnections == nil {
		var ret int32
		return ret
	}
	return *o.DdosActionLimitStreamsConnections
}

// GetDdosActionLimitStreamsConnectionsOk returns a tuple with the DdosActionLimitStreamsConnections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetDdosActionLimitStreamsConnectionsOk() (*int32, bool) {
	if o == nil || o.DdosActionLimitStreamsConnections == nil {
		return nil, false
	}
	return o.DdosActionLimitStreamsConnections, true
}

// HasDdosActionLimitStreamsConnections returns a boolean if a field has been set.
func (o *Results) HasDdosActionLimitStreamsConnections() bool {
	if o != nil && o.DdosActionLimitStreamsConnections != nil {
		return true
	}

	return false
}

// SetDdosActionLimitStreamsConnections gets a reference to the given int32 and assigns it to the DdosActionLimitStreamsConnections field.
func (o *Results) SetDdosActionLimitStreamsConnections(v int32) {
	o.DdosActionLimitStreamsConnections = &v
}

// GetDdosActionLimitStreamsRequests returns the DdosActionLimitStreamsRequests field value if set, zero value otherwise.
func (o *Results) GetDdosActionLimitStreamsRequests() int32 {
	if o == nil || o.DdosActionLimitStreamsRequests == nil {
		var ret int32
		return ret
	}
	return *o.DdosActionLimitStreamsRequests
}

// GetDdosActionLimitStreamsRequestsOk returns a tuple with the DdosActionLimitStreamsRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetDdosActionLimitStreamsRequestsOk() (*int32, bool) {
	if o == nil || o.DdosActionLimitStreamsRequests == nil {
		return nil, false
	}
	return o.DdosActionLimitStreamsRequests, true
}

// HasDdosActionLimitStreamsRequests returns a boolean if a field has been set.
func (o *Results) HasDdosActionLimitStreamsRequests() bool {
	if o != nil && o.DdosActionLimitStreamsRequests != nil {
		return true
	}

	return false
}

// SetDdosActionLimitStreamsRequests gets a reference to the given int32 and assigns it to the DdosActionLimitStreamsRequests field.
func (o *Results) SetDdosActionLimitStreamsRequests(v int32) {
	o.DdosActionLimitStreamsRequests = &v
}

// GetDdosActionTarpitAccept returns the DdosActionTarpitAccept field value if set, zero value otherwise.
func (o *Results) GetDdosActionTarpitAccept() int32 {
	if o == nil || o.DdosActionTarpitAccept == nil {
		var ret int32
		return ret
	}
	return *o.DdosActionTarpitAccept
}

// GetDdosActionTarpitAcceptOk returns a tuple with the DdosActionTarpitAccept field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetDdosActionTarpitAcceptOk() (*int32, bool) {
	if o == nil || o.DdosActionTarpitAccept == nil {
		return nil, false
	}
	return o.DdosActionTarpitAccept, true
}

// HasDdosActionTarpitAccept returns a boolean if a field has been set.
func (o *Results) HasDdosActionTarpitAccept() bool {
	if o != nil && o.DdosActionTarpitAccept != nil {
		return true
	}

	return false
}

// SetDdosActionTarpitAccept gets a reference to the given int32 and assigns it to the DdosActionTarpitAccept field.
func (o *Results) SetDdosActionTarpitAccept(v int32) {
	o.DdosActionTarpitAccept = &v
}

// GetDdosActionTarpit returns the DdosActionTarpit field value if set, zero value otherwise.
func (o *Results) GetDdosActionTarpit() int32 {
	if o == nil || o.DdosActionTarpit == nil {
		var ret int32
		return ret
	}
	return *o.DdosActionTarpit
}

// GetDdosActionTarpitOk returns a tuple with the DdosActionTarpit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetDdosActionTarpitOk() (*int32, bool) {
	if o == nil || o.DdosActionTarpit == nil {
		return nil, false
	}
	return o.DdosActionTarpit, true
}

// HasDdosActionTarpit returns a boolean if a field has been set.
func (o *Results) HasDdosActionTarpit() bool {
	if o != nil && o.DdosActionTarpit != nil {
		return true
	}

	return false
}

// SetDdosActionTarpit gets a reference to the given int32 and assigns it to the DdosActionTarpit field.
func (o *Results) SetDdosActionTarpit(v int32) {
	o.DdosActionTarpit = &v
}

// GetDdosActionClose returns the DdosActionClose field value if set, zero value otherwise.
func (o *Results) GetDdosActionClose() int32 {
	if o == nil || o.DdosActionClose == nil {
		var ret int32
		return ret
	}
	return *o.DdosActionClose
}

// GetDdosActionCloseOk returns a tuple with the DdosActionClose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetDdosActionCloseOk() (*int32, bool) {
	if o == nil || o.DdosActionClose == nil {
		return nil, false
	}
	return o.DdosActionClose, true
}

// HasDdosActionClose returns a boolean if a field has been set.
func (o *Results) HasDdosActionClose() bool {
	if o != nil && o.DdosActionClose != nil {
		return true
	}

	return false
}

// SetDdosActionClose gets a reference to the given int32 and assigns it to the DdosActionClose field.
func (o *Results) SetDdosActionClose(v int32) {
	o.DdosActionClose = &v
}

// GetDdosActionBlackhole returns the DdosActionBlackhole field value if set, zero value otherwise.
func (o *Results) GetDdosActionBlackhole() int32 {
	if o == nil || o.DdosActionBlackhole == nil {
		var ret int32
		return ret
	}
	return *o.DdosActionBlackhole
}

// GetDdosActionBlackholeOk returns a tuple with the DdosActionBlackhole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetDdosActionBlackholeOk() (*int32, bool) {
	if o == nil || o.DdosActionBlackhole == nil {
		return nil, false
	}
	return o.DdosActionBlackhole, true
}

// HasDdosActionBlackhole returns a boolean if a field has been set.
func (o *Results) HasDdosActionBlackhole() bool {
	if o != nil && o.DdosActionBlackhole != nil {
		return true
	}

	return false
}

// SetDdosActionBlackhole gets a reference to the given int32 and assigns it to the DdosActionBlackhole field.
func (o *Results) SetDdosActionBlackhole(v int32) {
	o.DdosActionBlackhole = &v
}

// GetBotChallengeStarts returns the BotChallengeStarts field value if set, zero value otherwise.
func (o *Results) GetBotChallengeStarts() int32 {
	if o == nil || o.BotChallengeStarts == nil {
		var ret int32
		return ret
	}
	return *o.BotChallengeStarts
}

// GetBotChallengeStartsOk returns a tuple with the BotChallengeStarts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetBotChallengeStartsOk() (*int32, bool) {
	if o == nil || o.BotChallengeStarts == nil {
		return nil, false
	}
	return o.BotChallengeStarts, true
}

// HasBotChallengeStarts returns a boolean if a field has been set.
func (o *Results) HasBotChallengeStarts() bool {
	if o != nil && o.BotChallengeStarts != nil {
		return true
	}

	return false
}

// SetBotChallengeStarts gets a reference to the given int32 and assigns it to the BotChallengeStarts field.
func (o *Results) SetBotChallengeStarts(v int32) {
	o.BotChallengeStarts = &v
}

// GetBotChallengeCompleteTokensPassed returns the BotChallengeCompleteTokensPassed field value if set, zero value otherwise.
func (o *Results) GetBotChallengeCompleteTokensPassed() int32 {
	if o == nil || o.BotChallengeCompleteTokensPassed == nil {
		var ret int32
		return ret
	}
	return *o.BotChallengeCompleteTokensPassed
}

// GetBotChallengeCompleteTokensPassedOk returns a tuple with the BotChallengeCompleteTokensPassed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetBotChallengeCompleteTokensPassedOk() (*int32, bool) {
	if o == nil || o.BotChallengeCompleteTokensPassed == nil {
		return nil, false
	}
	return o.BotChallengeCompleteTokensPassed, true
}

// HasBotChallengeCompleteTokensPassed returns a boolean if a field has been set.
func (o *Results) HasBotChallengeCompleteTokensPassed() bool {
	if o != nil && o.BotChallengeCompleteTokensPassed != nil {
		return true
	}

	return false
}

// SetBotChallengeCompleteTokensPassed gets a reference to the given int32 and assigns it to the BotChallengeCompleteTokensPassed field.
func (o *Results) SetBotChallengeCompleteTokensPassed(v int32) {
	o.BotChallengeCompleteTokensPassed = &v
}

// GetBotChallengeCompleteTokensFailed returns the BotChallengeCompleteTokensFailed field value if set, zero value otherwise.
func (o *Results) GetBotChallengeCompleteTokensFailed() int32 {
	if o == nil || o.BotChallengeCompleteTokensFailed == nil {
		var ret int32
		return ret
	}
	return *o.BotChallengeCompleteTokensFailed
}

// GetBotChallengeCompleteTokensFailedOk returns a tuple with the BotChallengeCompleteTokensFailed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetBotChallengeCompleteTokensFailedOk() (*int32, bool) {
	if o == nil || o.BotChallengeCompleteTokensFailed == nil {
		return nil, false
	}
	return o.BotChallengeCompleteTokensFailed, true
}

// HasBotChallengeCompleteTokensFailed returns a boolean if a field has been set.
func (o *Results) HasBotChallengeCompleteTokensFailed() bool {
	if o != nil && o.BotChallengeCompleteTokensFailed != nil {
		return true
	}

	return false
}

// SetBotChallengeCompleteTokensFailed gets a reference to the given int32 and assigns it to the BotChallengeCompleteTokensFailed field.
func (o *Results) SetBotChallengeCompleteTokensFailed(v int32) {
	o.BotChallengeCompleteTokensFailed = &v
}

// GetBotChallengeCompleteTokensChecked returns the BotChallengeCompleteTokensChecked field value if set, zero value otherwise.
func (o *Results) GetBotChallengeCompleteTokensChecked() int32 {
	if o == nil || o.BotChallengeCompleteTokensChecked == nil {
		var ret int32
		return ret
	}
	return *o.BotChallengeCompleteTokensChecked
}

// GetBotChallengeCompleteTokensCheckedOk returns a tuple with the BotChallengeCompleteTokensChecked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetBotChallengeCompleteTokensCheckedOk() (*int32, bool) {
	if o == nil || o.BotChallengeCompleteTokensChecked == nil {
		return nil, false
	}
	return o.BotChallengeCompleteTokensChecked, true
}

// HasBotChallengeCompleteTokensChecked returns a boolean if a field has been set.
func (o *Results) HasBotChallengeCompleteTokensChecked() bool {
	if o != nil && o.BotChallengeCompleteTokensChecked != nil {
		return true
	}

	return false
}

// SetBotChallengeCompleteTokensChecked gets a reference to the given int32 and assigns it to the BotChallengeCompleteTokensChecked field.
func (o *Results) SetBotChallengeCompleteTokensChecked(v int32) {
	o.BotChallengeCompleteTokensChecked = &v
}

// GetBotChallengeCompleteTokensDisabled returns the BotChallengeCompleteTokensDisabled field value if set, zero value otherwise.
func (o *Results) GetBotChallengeCompleteTokensDisabled() int32 {
	if o == nil || o.BotChallengeCompleteTokensDisabled == nil {
		var ret int32
		return ret
	}
	return *o.BotChallengeCompleteTokensDisabled
}

// GetBotChallengeCompleteTokensDisabledOk returns a tuple with the BotChallengeCompleteTokensDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetBotChallengeCompleteTokensDisabledOk() (*int32, bool) {
	if o == nil || o.BotChallengeCompleteTokensDisabled == nil {
		return nil, false
	}
	return o.BotChallengeCompleteTokensDisabled, true
}

// HasBotChallengeCompleteTokensDisabled returns a boolean if a field has been set.
func (o *Results) HasBotChallengeCompleteTokensDisabled() bool {
	if o != nil && o.BotChallengeCompleteTokensDisabled != nil {
		return true
	}

	return false
}

// SetBotChallengeCompleteTokensDisabled gets a reference to the given int32 and assigns it to the BotChallengeCompleteTokensDisabled field.
func (o *Results) SetBotChallengeCompleteTokensDisabled(v int32) {
	o.BotChallengeCompleteTokensDisabled = &v
}

// GetBotChallengeCompleteTokensIssued returns the BotChallengeCompleteTokensIssued field value if set, zero value otherwise.
func (o *Results) GetBotChallengeCompleteTokensIssued() int32 {
	if o == nil || o.BotChallengeCompleteTokensIssued == nil {
		var ret int32
		return ret
	}
	return *o.BotChallengeCompleteTokensIssued
}

// GetBotChallengeCompleteTokensIssuedOk returns a tuple with the BotChallengeCompleteTokensIssued field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetBotChallengeCompleteTokensIssuedOk() (*int32, bool) {
	if o == nil || o.BotChallengeCompleteTokensIssued == nil {
		return nil, false
	}
	return o.BotChallengeCompleteTokensIssued, true
}

// HasBotChallengeCompleteTokensIssued returns a boolean if a field has been set.
func (o *Results) HasBotChallengeCompleteTokensIssued() bool {
	if o != nil && o.BotChallengeCompleteTokensIssued != nil {
		return true
	}

	return false
}

// SetBotChallengeCompleteTokensIssued gets a reference to the given int32 and assigns it to the BotChallengeCompleteTokensIssued field.
func (o *Results) SetBotChallengeCompleteTokensIssued(v int32) {
	o.BotChallengeCompleteTokensIssued = &v
}

// GetBotChallengesIssued returns the BotChallengesIssued field value if set, zero value otherwise.
func (o *Results) GetBotChallengesIssued() int32 {
	if o == nil || o.BotChallengesIssued == nil {
		var ret int32
		return ret
	}
	return *o.BotChallengesIssued
}

// GetBotChallengesIssuedOk returns a tuple with the BotChallengesIssued field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetBotChallengesIssuedOk() (*int32, bool) {
	if o == nil || o.BotChallengesIssued == nil {
		return nil, false
	}
	return o.BotChallengesIssued, true
}

// HasBotChallengesIssued returns a boolean if a field has been set.
func (o *Results) HasBotChallengesIssued() bool {
	if o != nil && o.BotChallengesIssued != nil {
		return true
	}

	return false
}

// SetBotChallengesIssued gets a reference to the given int32 and assigns it to the BotChallengesIssued field.
func (o *Results) SetBotChallengesIssued(v int32) {
	o.BotChallengesIssued = &v
}

// GetBotChallengesSucceeded returns the BotChallengesSucceeded field value if set, zero value otherwise.
func (o *Results) GetBotChallengesSucceeded() int32 {
	if o == nil || o.BotChallengesSucceeded == nil {
		var ret int32
		return ret
	}
	return *o.BotChallengesSucceeded
}

// GetBotChallengesSucceededOk returns a tuple with the BotChallengesSucceeded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetBotChallengesSucceededOk() (*int32, bool) {
	if o == nil || o.BotChallengesSucceeded == nil {
		return nil, false
	}
	return o.BotChallengesSucceeded, true
}

// HasBotChallengesSucceeded returns a boolean if a field has been set.
func (o *Results) HasBotChallengesSucceeded() bool {
	if o != nil && o.BotChallengesSucceeded != nil {
		return true
	}

	return false
}

// SetBotChallengesSucceeded gets a reference to the given int32 and assigns it to the BotChallengesSucceeded field.
func (o *Results) SetBotChallengesSucceeded(v int32) {
	o.BotChallengesSucceeded = &v
}

// GetBotChallengesFailed returns the BotChallengesFailed field value if set, zero value otherwise.
func (o *Results) GetBotChallengesFailed() int32 {
	if o == nil || o.BotChallengesFailed == nil {
		var ret int32
		return ret
	}
	return *o.BotChallengesFailed
}

// GetBotChallengesFailedOk returns a tuple with the BotChallengesFailed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetBotChallengesFailedOk() (*int32, bool) {
	if o == nil || o.BotChallengesFailed == nil {
		return nil, false
	}
	return o.BotChallengesFailed, true
}

// HasBotChallengesFailed returns a boolean if a field has been set.
func (o *Results) HasBotChallengesFailed() bool {
	if o != nil && o.BotChallengesFailed != nil {
		return true
	}

	return false
}

// SetBotChallengesFailed gets a reference to the given int32 and assigns it to the BotChallengesFailed field.
func (o *Results) SetBotChallengesFailed(v int32) {
	o.BotChallengesFailed = &v
}

// GetDdosActionDowngrade returns the DdosActionDowngrade field value if set, zero value otherwise.
func (o *Results) GetDdosActionDowngrade() int32 {
	if o == nil || o.DdosActionDowngrade == nil {
		var ret int32
		return ret
	}
	return *o.DdosActionDowngrade
}

// GetDdosActionDowngradeOk returns a tuple with the DdosActionDowngrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetDdosActionDowngradeOk() (*int32, bool) {
	if o == nil || o.DdosActionDowngrade == nil {
		return nil, false
	}
	return o.DdosActionDowngrade, true
}

// HasDdosActionDowngrade returns a boolean if a field has been set.
func (o *Results) HasDdosActionDowngrade() bool {
	if o != nil && o.DdosActionDowngrade != nil {
		return true
	}

	return false
}

// SetDdosActionDowngrade gets a reference to the given int32 and assigns it to the DdosActionDowngrade field.
func (o *Results) SetDdosActionDowngrade(v int32) {
	o.DdosActionDowngrade = &v
}

// GetDdosActionDowngradedConnections returns the DdosActionDowngradedConnections field value if set, zero value otherwise.
func (o *Results) GetDdosActionDowngradedConnections() int32 {
	if o == nil || o.DdosActionDowngradedConnections == nil {
		var ret int32
		return ret
	}
	return *o.DdosActionDowngradedConnections
}

// GetDdosActionDowngradedConnectionsOk returns a tuple with the DdosActionDowngradedConnections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetDdosActionDowngradedConnectionsOk() (*int32, bool) {
	if o == nil || o.DdosActionDowngradedConnections == nil {
		return nil, false
	}
	return o.DdosActionDowngradedConnections, true
}

// HasDdosActionDowngradedConnections returns a boolean if a field has been set.
func (o *Results) HasDdosActionDowngradedConnections() bool {
	if o != nil && o.DdosActionDowngradedConnections != nil {
		return true
	}

	return false
}

// SetDdosActionDowngradedConnections gets a reference to the given int32 and assigns it to the DdosActionDowngradedConnections field.
func (o *Results) SetDdosActionDowngradedConnections(v int32) {
	o.DdosActionDowngradedConnections = &v
}

// GetVclOnComputeHitRequests returns the VclOnComputeHitRequests field value if set, zero value otherwise.
func (o *Results) GetVclOnComputeHitRequests() int32 {
	if o == nil || o.VclOnComputeHitRequests == nil {
		var ret int32
		return ret
	}
	return *o.VclOnComputeHitRequests
}

// GetVclOnComputeHitRequestsOk returns a tuple with the VclOnComputeHitRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetVclOnComputeHitRequestsOk() (*int32, bool) {
	if o == nil || o.VclOnComputeHitRequests == nil {
		return nil, false
	}
	return o.VclOnComputeHitRequests, true
}

// HasVclOnComputeHitRequests returns a boolean if a field has been set.
func (o *Results) HasVclOnComputeHitRequests() bool {
	if o != nil && o.VclOnComputeHitRequests != nil {
		return true
	}

	return false
}

// SetVclOnComputeHitRequests gets a reference to the given int32 and assigns it to the VclOnComputeHitRequests field.
func (o *Results) SetVclOnComputeHitRequests(v int32) {
	o.VclOnComputeHitRequests = &v
}

// GetVclOnComputeMissRequests returns the VclOnComputeMissRequests field value if set, zero value otherwise.
func (o *Results) GetVclOnComputeMissRequests() int32 {
	if o == nil || o.VclOnComputeMissRequests == nil {
		var ret int32
		return ret
	}
	return *o.VclOnComputeMissRequests
}

// GetVclOnComputeMissRequestsOk returns a tuple with the VclOnComputeMissRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetVclOnComputeMissRequestsOk() (*int32, bool) {
	if o == nil || o.VclOnComputeMissRequests == nil {
		return nil, false
	}
	return o.VclOnComputeMissRequests, true
}

// HasVclOnComputeMissRequests returns a boolean if a field has been set.
func (o *Results) HasVclOnComputeMissRequests() bool {
	if o != nil && o.VclOnComputeMissRequests != nil {
		return true
	}

	return false
}

// SetVclOnComputeMissRequests gets a reference to the given int32 and assigns it to the VclOnComputeMissRequests field.
func (o *Results) SetVclOnComputeMissRequests(v int32) {
	o.VclOnComputeMissRequests = &v
}

// GetVclOnComputePassRequests returns the VclOnComputePassRequests field value if set, zero value otherwise.
func (o *Results) GetVclOnComputePassRequests() int32 {
	if o == nil || o.VclOnComputePassRequests == nil {
		var ret int32
		return ret
	}
	return *o.VclOnComputePassRequests
}

// GetVclOnComputePassRequestsOk returns a tuple with the VclOnComputePassRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetVclOnComputePassRequestsOk() (*int32, bool) {
	if o == nil || o.VclOnComputePassRequests == nil {
		return nil, false
	}
	return o.VclOnComputePassRequests, true
}

// HasVclOnComputePassRequests returns a boolean if a field has been set.
func (o *Results) HasVclOnComputePassRequests() bool {
	if o != nil && o.VclOnComputePassRequests != nil {
		return true
	}

	return false
}

// SetVclOnComputePassRequests gets a reference to the given int32 and assigns it to the VclOnComputePassRequests field.
func (o *Results) SetVclOnComputePassRequests(v int32) {
	o.VclOnComputePassRequests = &v
}

// GetVclOnComputeErrorRequests returns the VclOnComputeErrorRequests field value if set, zero value otherwise.
func (o *Results) GetVclOnComputeErrorRequests() int32 {
	if o == nil || o.VclOnComputeErrorRequests == nil {
		var ret int32
		return ret
	}
	return *o.VclOnComputeErrorRequests
}

// GetVclOnComputeErrorRequestsOk returns a tuple with the VclOnComputeErrorRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetVclOnComputeErrorRequestsOk() (*int32, bool) {
	if o == nil || o.VclOnComputeErrorRequests == nil {
		return nil, false
	}
	return o.VclOnComputeErrorRequests, true
}

// HasVclOnComputeErrorRequests returns a boolean if a field has been set.
func (o *Results) HasVclOnComputeErrorRequests() bool {
	if o != nil && o.VclOnComputeErrorRequests != nil {
		return true
	}

	return false
}

// SetVclOnComputeErrorRequests gets a reference to the given int32 and assigns it to the VclOnComputeErrorRequests field.
func (o *Results) SetVclOnComputeErrorRequests(v int32) {
	o.VclOnComputeErrorRequests = &v
}

// GetVclOnComputeSynthRequests returns the VclOnComputeSynthRequests field value if set, zero value otherwise.
func (o *Results) GetVclOnComputeSynthRequests() int32 {
	if o == nil || o.VclOnComputeSynthRequests == nil {
		var ret int32
		return ret
	}
	return *o.VclOnComputeSynthRequests
}

// GetVclOnComputeSynthRequestsOk returns a tuple with the VclOnComputeSynthRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetVclOnComputeSynthRequestsOk() (*int32, bool) {
	if o == nil || o.VclOnComputeSynthRequests == nil {
		return nil, false
	}
	return o.VclOnComputeSynthRequests, true
}

// HasVclOnComputeSynthRequests returns a boolean if a field has been set.
func (o *Results) HasVclOnComputeSynthRequests() bool {
	if o != nil && o.VclOnComputeSynthRequests != nil {
		return true
	}

	return false
}

// SetVclOnComputeSynthRequests gets a reference to the given int32 and assigns it to the VclOnComputeSynthRequests field.
func (o *Results) SetVclOnComputeSynthRequests(v int32) {
	o.VclOnComputeSynthRequests = &v
}

// GetVclOnComputeEdgeHitRequests returns the VclOnComputeEdgeHitRequests field value if set, zero value otherwise.
func (o *Results) GetVclOnComputeEdgeHitRequests() int32 {
	if o == nil || o.VclOnComputeEdgeHitRequests == nil {
		var ret int32
		return ret
	}
	return *o.VclOnComputeEdgeHitRequests
}

// GetVclOnComputeEdgeHitRequestsOk returns a tuple with the VclOnComputeEdgeHitRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetVclOnComputeEdgeHitRequestsOk() (*int32, bool) {
	if o == nil || o.VclOnComputeEdgeHitRequests == nil {
		return nil, false
	}
	return o.VclOnComputeEdgeHitRequests, true
}

// HasVclOnComputeEdgeHitRequests returns a boolean if a field has been set.
func (o *Results) HasVclOnComputeEdgeHitRequests() bool {
	if o != nil && o.VclOnComputeEdgeHitRequests != nil {
		return true
	}

	return false
}

// SetVclOnComputeEdgeHitRequests gets a reference to the given int32 and assigns it to the VclOnComputeEdgeHitRequests field.
func (o *Results) SetVclOnComputeEdgeHitRequests(v int32) {
	o.VclOnComputeEdgeHitRequests = &v
}

// GetVclOnComputeEdgeMissRequests returns the VclOnComputeEdgeMissRequests field value if set, zero value otherwise.
func (o *Results) GetVclOnComputeEdgeMissRequests() int32 {
	if o == nil || o.VclOnComputeEdgeMissRequests == nil {
		var ret int32
		return ret
	}
	return *o.VclOnComputeEdgeMissRequests
}

// GetVclOnComputeEdgeMissRequestsOk returns a tuple with the VclOnComputeEdgeMissRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetVclOnComputeEdgeMissRequestsOk() (*int32, bool) {
	if o == nil || o.VclOnComputeEdgeMissRequests == nil {
		return nil, false
	}
	return o.VclOnComputeEdgeMissRequests, true
}

// HasVclOnComputeEdgeMissRequests returns a boolean if a field has been set.
func (o *Results) HasVclOnComputeEdgeMissRequests() bool {
	if o != nil && o.VclOnComputeEdgeMissRequests != nil {
		return true
	}

	return false
}

// SetVclOnComputeEdgeMissRequests gets a reference to the given int32 and assigns it to the VclOnComputeEdgeMissRequests field.
func (o *Results) SetVclOnComputeEdgeMissRequests(v int32) {
	o.VclOnComputeEdgeMissRequests = &v
}

// GetAllHitRequests returns the AllHitRequests field value if set, zero value otherwise.
func (o *Results) GetAllHitRequests() int32 {
	if o == nil || o.AllHitRequests == nil {
		var ret int32
		return ret
	}
	return *o.AllHitRequests
}

// GetAllHitRequestsOk returns a tuple with the AllHitRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetAllHitRequestsOk() (*int32, bool) {
	if o == nil || o.AllHitRequests == nil {
		return nil, false
	}
	return o.AllHitRequests, true
}

// HasAllHitRequests returns a boolean if a field has been set.
func (o *Results) HasAllHitRequests() bool {
	if o != nil && o.AllHitRequests != nil {
		return true
	}

	return false
}

// SetAllHitRequests gets a reference to the given int32 and assigns it to the AllHitRequests field.
func (o *Results) SetAllHitRequests(v int32) {
	o.AllHitRequests = &v
}

// GetAllMissRequests returns the AllMissRequests field value if set, zero value otherwise.
func (o *Results) GetAllMissRequests() int32 {
	if o == nil || o.AllMissRequests == nil {
		var ret int32
		return ret
	}
	return *o.AllMissRequests
}

// GetAllMissRequestsOk returns a tuple with the AllMissRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetAllMissRequestsOk() (*int32, bool) {
	if o == nil || o.AllMissRequests == nil {
		return nil, false
	}
	return o.AllMissRequests, true
}

// HasAllMissRequests returns a boolean if a field has been set.
func (o *Results) HasAllMissRequests() bool {
	if o != nil && o.AllMissRequests != nil {
		return true
	}

	return false
}

// SetAllMissRequests gets a reference to the given int32 and assigns it to the AllMissRequests field.
func (o *Results) SetAllMissRequests(v int32) {
	o.AllMissRequests = &v
}

// GetAllPassRequests returns the AllPassRequests field value if set, zero value otherwise.
func (o *Results) GetAllPassRequests() int32 {
	if o == nil || o.AllPassRequests == nil {
		var ret int32
		return ret
	}
	return *o.AllPassRequests
}

// GetAllPassRequestsOk returns a tuple with the AllPassRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetAllPassRequestsOk() (*int32, bool) {
	if o == nil || o.AllPassRequests == nil {
		return nil, false
	}
	return o.AllPassRequests, true
}

// HasAllPassRequests returns a boolean if a field has been set.
func (o *Results) HasAllPassRequests() bool {
	if o != nil && o.AllPassRequests != nil {
		return true
	}

	return false
}

// SetAllPassRequests gets a reference to the given int32 and assigns it to the AllPassRequests field.
func (o *Results) SetAllPassRequests(v int32) {
	o.AllPassRequests = &v
}

// GetAllErrorRequests returns the AllErrorRequests field value if set, zero value otherwise.
func (o *Results) GetAllErrorRequests() int32 {
	if o == nil || o.AllErrorRequests == nil {
		var ret int32
		return ret
	}
	return *o.AllErrorRequests
}

// GetAllErrorRequestsOk returns a tuple with the AllErrorRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetAllErrorRequestsOk() (*int32, bool) {
	if o == nil || o.AllErrorRequests == nil {
		return nil, false
	}
	return o.AllErrorRequests, true
}

// HasAllErrorRequests returns a boolean if a field has been set.
func (o *Results) HasAllErrorRequests() bool {
	if o != nil && o.AllErrorRequests != nil {
		return true
	}

	return false
}

// SetAllErrorRequests gets a reference to the given int32 and assigns it to the AllErrorRequests field.
func (o *Results) SetAllErrorRequests(v int32) {
	o.AllErrorRequests = &v
}

// GetAllSynthRequests returns the AllSynthRequests field value if set, zero value otherwise.
func (o *Results) GetAllSynthRequests() int32 {
	if o == nil || o.AllSynthRequests == nil {
		var ret int32
		return ret
	}
	return *o.AllSynthRequests
}

// GetAllSynthRequestsOk returns a tuple with the AllSynthRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetAllSynthRequestsOk() (*int32, bool) {
	if o == nil || o.AllSynthRequests == nil {
		return nil, false
	}
	return o.AllSynthRequests, true
}

// HasAllSynthRequests returns a boolean if a field has been set.
func (o *Results) HasAllSynthRequests() bool {
	if o != nil && o.AllSynthRequests != nil {
		return true
	}

	return false
}

// SetAllSynthRequests gets a reference to the given int32 and assigns it to the AllSynthRequests field.
func (o *Results) SetAllSynthRequests(v int32) {
	o.AllSynthRequests = &v
}

// GetAllEdgeHitRequests returns the AllEdgeHitRequests field value if set, zero value otherwise.
func (o *Results) GetAllEdgeHitRequests() int32 {
	if o == nil || o.AllEdgeHitRequests == nil {
		var ret int32
		return ret
	}
	return *o.AllEdgeHitRequests
}

// GetAllEdgeHitRequestsOk returns a tuple with the AllEdgeHitRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetAllEdgeHitRequestsOk() (*int32, bool) {
	if o == nil || o.AllEdgeHitRequests == nil {
		return nil, false
	}
	return o.AllEdgeHitRequests, true
}

// HasAllEdgeHitRequests returns a boolean if a field has been set.
func (o *Results) HasAllEdgeHitRequests() bool {
	if o != nil && o.AllEdgeHitRequests != nil {
		return true
	}

	return false
}

// SetAllEdgeHitRequests gets a reference to the given int32 and assigns it to the AllEdgeHitRequests field.
func (o *Results) SetAllEdgeHitRequests(v int32) {
	o.AllEdgeHitRequests = &v
}

// GetAllEdgeMissRequests returns the AllEdgeMissRequests field value if set, zero value otherwise.
func (o *Results) GetAllEdgeMissRequests() int32 {
	if o == nil || o.AllEdgeMissRequests == nil {
		var ret int32
		return ret
	}
	return *o.AllEdgeMissRequests
}

// GetAllEdgeMissRequestsOk returns a tuple with the AllEdgeMissRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetAllEdgeMissRequestsOk() (*int32, bool) {
	if o == nil || o.AllEdgeMissRequests == nil {
		return nil, false
	}
	return o.AllEdgeMissRequests, true
}

// HasAllEdgeMissRequests returns a boolean if a field has been set.
func (o *Results) HasAllEdgeMissRequests() bool {
	if o != nil && o.AllEdgeMissRequests != nil {
		return true
	}

	return false
}

// SetAllEdgeMissRequests gets a reference to the given int32 and assigns it to the AllEdgeMissRequests field.
func (o *Results) SetAllEdgeMissRequests(v int32) {
	o.AllEdgeMissRequests = &v
}

// GetAllStatus1xx returns the AllStatus1xx field value if set, zero value otherwise.
func (o *Results) GetAllStatus1xx() int32 {
	if o == nil || o.AllStatus1xx == nil {
		var ret int32
		return ret
	}
	return *o.AllStatus1xx
}

// GetAllStatus1xxOk returns a tuple with the AllStatus1xx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetAllStatus1xxOk() (*int32, bool) {
	if o == nil || o.AllStatus1xx == nil {
		return nil, false
	}
	return o.AllStatus1xx, true
}

// HasAllStatus1xx returns a boolean if a field has been set.
func (o *Results) HasAllStatus1xx() bool {
	if o != nil && o.AllStatus1xx != nil {
		return true
	}

	return false
}

// SetAllStatus1xx gets a reference to the given int32 and assigns it to the AllStatus1xx field.
func (o *Results) SetAllStatus1xx(v int32) {
	o.AllStatus1xx = &v
}

// GetAllStatus2xx returns the AllStatus2xx field value if set, zero value otherwise.
func (o *Results) GetAllStatus2xx() int32 {
	if o == nil || o.AllStatus2xx == nil {
		var ret int32
		return ret
	}
	return *o.AllStatus2xx
}

// GetAllStatus2xxOk returns a tuple with the AllStatus2xx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetAllStatus2xxOk() (*int32, bool) {
	if o == nil || o.AllStatus2xx == nil {
		return nil, false
	}
	return o.AllStatus2xx, true
}

// HasAllStatus2xx returns a boolean if a field has been set.
func (o *Results) HasAllStatus2xx() bool {
	if o != nil && o.AllStatus2xx != nil {
		return true
	}

	return false
}

// SetAllStatus2xx gets a reference to the given int32 and assigns it to the AllStatus2xx field.
func (o *Results) SetAllStatus2xx(v int32) {
	o.AllStatus2xx = &v
}

// GetAllStatus3xx returns the AllStatus3xx field value if set, zero value otherwise.
func (o *Results) GetAllStatus3xx() int32 {
	if o == nil || o.AllStatus3xx == nil {
		var ret int32
		return ret
	}
	return *o.AllStatus3xx
}

// GetAllStatus3xxOk returns a tuple with the AllStatus3xx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetAllStatus3xxOk() (*int32, bool) {
	if o == nil || o.AllStatus3xx == nil {
		return nil, false
	}
	return o.AllStatus3xx, true
}

// HasAllStatus3xx returns a boolean if a field has been set.
func (o *Results) HasAllStatus3xx() bool {
	if o != nil && o.AllStatus3xx != nil {
		return true
	}

	return false
}

// SetAllStatus3xx gets a reference to the given int32 and assigns it to the AllStatus3xx field.
func (o *Results) SetAllStatus3xx(v int32) {
	o.AllStatus3xx = &v
}

// GetAllStatus4xx returns the AllStatus4xx field value if set, zero value otherwise.
func (o *Results) GetAllStatus4xx() int32 {
	if o == nil || o.AllStatus4xx == nil {
		var ret int32
		return ret
	}
	return *o.AllStatus4xx
}

// GetAllStatus4xxOk returns a tuple with the AllStatus4xx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetAllStatus4xxOk() (*int32, bool) {
	if o == nil || o.AllStatus4xx == nil {
		return nil, false
	}
	return o.AllStatus4xx, true
}

// HasAllStatus4xx returns a boolean if a field has been set.
func (o *Results) HasAllStatus4xx() bool {
	if o != nil && o.AllStatus4xx != nil {
		return true
	}

	return false
}

// SetAllStatus4xx gets a reference to the given int32 and assigns it to the AllStatus4xx field.
func (o *Results) SetAllStatus4xx(v int32) {
	o.AllStatus4xx = &v
}

// GetAllStatus5xx returns the AllStatus5xx field value if set, zero value otherwise.
func (o *Results) GetAllStatus5xx() int32 {
	if o == nil || o.AllStatus5xx == nil {
		var ret int32
		return ret
	}
	return *o.AllStatus5xx
}

// GetAllStatus5xxOk returns a tuple with the AllStatus5xx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetAllStatus5xxOk() (*int32, bool) {
	if o == nil || o.AllStatus5xx == nil {
		return nil, false
	}
	return o.AllStatus5xx, true
}

// HasAllStatus5xx returns a boolean if a field has been set.
func (o *Results) HasAllStatus5xx() bool {
	if o != nil && o.AllStatus5xx != nil {
		return true
	}

	return false
}

// SetAllStatus5xx gets a reference to the given int32 and assigns it to the AllStatus5xx field.
func (o *Results) SetAllStatus5xx(v int32) {
	o.AllStatus5xx = &v
}

// GetOriginOffload returns the OriginOffload field value if set, zero value otherwise.
func (o *Results) GetOriginOffload() float32 {
	if o == nil || o.OriginOffload == nil {
		var ret float32
		return ret
	}
	return *o.OriginOffload
}

// GetOriginOffloadOk returns a tuple with the OriginOffload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetOriginOffloadOk() (*float32, bool) {
	if o == nil || o.OriginOffload == nil {
		return nil, false
	}
	return o.OriginOffload, true
}

// HasOriginOffload returns a boolean if a field has been set.
func (o *Results) HasOriginOffload() bool {
	if o != nil && o.OriginOffload != nil {
		return true
	}

	return false
}

// SetOriginOffload gets a reference to the given float32 and assigns it to the OriginOffload field.
func (o *Results) SetOriginOffload(v float32) {
	o.OriginOffload = &v
}

// GetRequestDeniedGetHeadBody returns the RequestDeniedGetHeadBody field value if set, zero value otherwise.
func (o *Results) GetRequestDeniedGetHeadBody() int32 {
	if o == nil || o.RequestDeniedGetHeadBody == nil {
		var ret int32
		return ret
	}
	return *o.RequestDeniedGetHeadBody
}

// GetRequestDeniedGetHeadBodyOk returns a tuple with the RequestDeniedGetHeadBody field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetRequestDeniedGetHeadBodyOk() (*int32, bool) {
	if o == nil || o.RequestDeniedGetHeadBody == nil {
		return nil, false
	}
	return o.RequestDeniedGetHeadBody, true
}

// HasRequestDeniedGetHeadBody returns a boolean if a field has been set.
func (o *Results) HasRequestDeniedGetHeadBody() bool {
	if o != nil && o.RequestDeniedGetHeadBody != nil {
		return true
	}

	return false
}

// SetRequestDeniedGetHeadBody gets a reference to the given int32 and assigns it to the RequestDeniedGetHeadBody field.
func (o *Results) SetRequestDeniedGetHeadBody(v int32) {
	o.RequestDeniedGetHeadBody = &v
}

// GetServiceID returns the ServiceID field value if set, zero value otherwise.
func (o *Results) GetServiceID() string {
	if o == nil || o.ServiceID == nil {
		var ret string
		return ret
	}
	return *o.ServiceID
}

// GetServiceIDOk returns a tuple with the ServiceID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetServiceIDOk() (*string, bool) {
	if o == nil || o.ServiceID == nil {
		return nil, false
	}
	return o.ServiceID, true
}

// HasServiceID returns a boolean if a field has been set.
func (o *Results) HasServiceID() bool {
	if o != nil && o.ServiceID != nil {
		return true
	}

	return false
}

// SetServiceID gets a reference to the given string and assigns it to the ServiceID field.
func (o *Results) SetServiceID(v string) {
	o.ServiceID = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *Results) GetStartTime() int32 {
	if o == nil || o.StartTime == nil {
		var ret int32
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Results) GetStartTimeOk() (*int32, bool) {
	if o == nil || o.StartTime == nil {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *Results) HasStartTime() bool {
	if o != nil && o.StartTime != nil {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given int32 and assigns it to the StartTime field.
func (o *Results) SetStartTime(v int32) {
	o.StartTime = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o Results) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Requests != nil {
		toSerialize["requests"] = o.Requests
	}
	if o.Hits != nil {
		toSerialize["hits"] = o.Hits
	}
	if o.HitsTime != nil {
		toSerialize["hits_time"] = o.HitsTime
	}
	if o.Miss != nil {
		toSerialize["miss"] = o.Miss
	}
	if o.MissTime != nil {
		toSerialize["miss_time"] = o.MissTime
	}
	if o.Pass != nil {
		toSerialize["pass"] = o.Pass
	}
	if o.PassTime != nil {
		toSerialize["pass_time"] = o.PassTime
	}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}
	if o.Restarts != nil {
		toSerialize["restarts"] = o.Restarts
	}
	if o.HitRatio.IsSet() {
		toSerialize["hit_ratio"] = o.HitRatio.Get()
	}
	if o.Bandwidth != nil {
		toSerialize["bandwidth"] = o.Bandwidth
	}
	if o.BodySize != nil {
		toSerialize["body_size"] = o.BodySize
	}
	if o.HeaderSize != nil {
		toSerialize["header_size"] = o.HeaderSize
	}
	if o.ReqBodyBytes != nil {
		toSerialize["req_body_bytes"] = o.ReqBodyBytes
	}
	if o.ReqHeaderBytes != nil {
		toSerialize["req_header_bytes"] = o.ReqHeaderBytes
	}
	if o.RespBodyBytes != nil {
		toSerialize["resp_body_bytes"] = o.RespBodyBytes
	}
	if o.RespHeaderBytes != nil {
		toSerialize["resp_header_bytes"] = o.RespHeaderBytes
	}
	if o.BereqBodyBytes != nil {
		toSerialize["bereq_body_bytes"] = o.BereqBodyBytes
	}
	if o.BereqHeaderBytes != nil {
		toSerialize["bereq_header_bytes"] = o.BereqHeaderBytes
	}
	if o.Uncacheable != nil {
		toSerialize["uncacheable"] = o.Uncacheable
	}
	if o.Pipe != nil {
		toSerialize["pipe"] = o.Pipe
	}
	if o.Synth != nil {
		toSerialize["synth"] = o.Synth
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
	if o.EdgeRequests != nil {
		toSerialize["edge_requests"] = o.EdgeRequests
	}
	if o.EdgeRespHeaderBytes != nil {
		toSerialize["edge_resp_header_bytes"] = o.EdgeRespHeaderBytes
	}
	if o.EdgeRespBodyBytes != nil {
		toSerialize["edge_resp_body_bytes"] = o.EdgeRespBodyBytes
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
	if o.OriginRevalidations != nil {
		toSerialize["origin_revalidations"] = o.OriginRevalidations
	}
	if o.OriginCacheFetches != nil {
		toSerialize["origin_cache_fetches"] = o.OriginCacheFetches
	}
	if o.Shield != nil {
		toSerialize["shield"] = o.Shield
	}
	if o.ShieldRespBodyBytes != nil {
		toSerialize["shield_resp_body_bytes"] = o.ShieldRespBodyBytes
	}
	if o.ShieldRespHeaderBytes != nil {
		toSerialize["shield_resp_header_bytes"] = o.ShieldRespHeaderBytes
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
	if o.ShieldRevalidations != nil {
		toSerialize["shield_revalidations"] = o.ShieldRevalidations
	}
	if o.ShieldCacheFetches != nil {
		toSerialize["shield_cache_fetches"] = o.ShieldCacheFetches
	}
	if o.Ipv6 != nil {
		toSerialize["ipv6"] = o.Ipv6
	}
	if o.Otfp != nil {
		toSerialize["otfp"] = o.Otfp
	}
	if o.OtfpRespBodyBytes != nil {
		toSerialize["otfp_resp_body_bytes"] = o.OtfpRespBodyBytes
	}
	if o.OtfpRespHeaderBytes != nil {
		toSerialize["otfp_resp_header_bytes"] = o.OtfpRespHeaderBytes
	}
	if o.OtfpShieldRespBodyBytes != nil {
		toSerialize["otfp_shield_resp_body_bytes"] = o.OtfpShieldRespBodyBytes
	}
	if o.OtfpShieldRespHeaderBytes != nil {
		toSerialize["otfp_shield_resp_header_bytes"] = o.OtfpShieldRespHeaderBytes
	}
	if o.OtfpManifests != nil {
		toSerialize["otfp_manifests"] = o.OtfpManifests
	}
	if o.OtfpDeliverTime != nil {
		toSerialize["otfp_deliver_time"] = o.OtfpDeliverTime
	}
	if o.OtfpShieldTime != nil {
		toSerialize["otfp_shield_time"] = o.OtfpShieldTime
	}
	if o.Video != nil {
		toSerialize["video"] = o.Video
	}
	if o.Pci != nil {
		toSerialize["pci"] = o.Pci
	}
	if o.Log != nil {
		toSerialize["log"] = o.Log
	}
	if o.LogBytes != nil {
		toSerialize["log_bytes"] = o.LogBytes
	}
	if o.HTTP2 != nil {
		toSerialize["http2"] = o.HTTP2
	}
	if o.HTTP3 != nil {
		toSerialize["http3"] = o.HTTP3
	}
	if o.WafLogged != nil {
		toSerialize["waf_logged"] = o.WafLogged
	}
	if o.WafBlocked != nil {
		toSerialize["waf_blocked"] = o.WafBlocked
	}
	if o.WafPassed != nil {
		toSerialize["waf_passed"] = o.WafPassed
	}
	if o.AttackReqBodyBytes != nil {
		toSerialize["attack_req_body_bytes"] = o.AttackReqBodyBytes
	}
	if o.AttackReqHeaderBytes != nil {
		toSerialize["attack_req_header_bytes"] = o.AttackReqHeaderBytes
	}
	if o.AttackLoggedReqBodyBytes != nil {
		toSerialize["attack_logged_req_body_bytes"] = o.AttackLoggedReqBodyBytes
	}
	if o.AttackLoggedReqHeaderBytes != nil {
		toSerialize["attack_logged_req_header_bytes"] = o.AttackLoggedReqHeaderBytes
	}
	if o.AttackBlockedReqBodyBytes != nil {
		toSerialize["attack_blocked_req_body_bytes"] = o.AttackBlockedReqBodyBytes
	}
	if o.AttackBlockedReqHeaderBytes != nil {
		toSerialize["attack_blocked_req_header_bytes"] = o.AttackBlockedReqHeaderBytes
	}
	if o.AttackPassedReqBodyBytes != nil {
		toSerialize["attack_passed_req_body_bytes"] = o.AttackPassedReqBodyBytes
	}
	if o.AttackPassedReqHeaderBytes != nil {
		toSerialize["attack_passed_req_header_bytes"] = o.AttackPassedReqHeaderBytes
	}
	if o.AttackRespSynthBytes != nil {
		toSerialize["attack_resp_synth_bytes"] = o.AttackRespSynthBytes
	}
	if o.Imgopto != nil {
		toSerialize["imgopto"] = o.Imgopto
	}
	if o.ImgoptoRespBodyBytes != nil {
		toSerialize["imgopto_resp_body_bytes"] = o.ImgoptoRespBodyBytes
	}
	if o.ImgoptoRespHeaderBytes != nil {
		toSerialize["imgopto_resp_header_bytes"] = o.ImgoptoRespHeaderBytes
	}
	if o.ImgoptoShield != nil {
		toSerialize["imgopto_shield"] = o.ImgoptoShield
	}
	if o.ImgoptoShieldRespBodyBytes != nil {
		toSerialize["imgopto_shield_resp_body_bytes"] = o.ImgoptoShieldRespBodyBytes
	}
	if o.ImgoptoShieldRespHeaderBytes != nil {
		toSerialize["imgopto_shield_resp_header_bytes"] = o.ImgoptoShieldRespHeaderBytes
	}
	if o.ImgoptoTransforms != nil {
		toSerialize["imgopto_transforms"] = o.ImgoptoTransforms
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
	if o.ImgvideoShieldRespHeaderBytes != nil {
		toSerialize["imgvideo_shield_resp_header_bytes"] = o.ImgvideoShieldRespHeaderBytes
	}
	if o.ImgvideoShieldRespBodyBytes != nil {
		toSerialize["imgvideo_shield_resp_body_bytes"] = o.ImgvideoShieldRespBodyBytes
	}
	if o.ImgvideoShield != nil {
		toSerialize["imgvideo_shield"] = o.ImgvideoShield
	}
	if o.ImgvideoShieldFrames != nil {
		toSerialize["imgvideo_shield_frames"] = o.ImgvideoShieldFrames
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
	if o.TLSHandshakeSentBytes != nil {
		toSerialize["tls_handshake_sent_bytes"] = o.TLSHandshakeSentBytes
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
	if o.SegblockOriginFetches != nil {
		toSerialize["segblock_origin_fetches"] = o.SegblockOriginFetches
	}
	if o.SegblockShieldFetches != nil {
		toSerialize["segblock_shield_fetches"] = o.SegblockShieldFetches
	}
	if o.ComputeRequests != nil {
		toSerialize["compute_requests"] = o.ComputeRequests
	}
	if o.ComputeRequestTimeMs != nil {
		toSerialize["compute_request_time_ms"] = o.ComputeRequestTimeMs
	}
	if o.ComputeRequestTimeBilledMs != nil {
		toSerialize["compute_request_time_billed_ms"] = o.ComputeRequestTimeBilledMs
	}
	if o.ComputeRAMUsed != nil {
		toSerialize["compute_ram_used"] = o.ComputeRAMUsed
	}
	if o.ComputeExecutionTimeMs != nil {
		toSerialize["compute_execution_time_ms"] = o.ComputeExecutionTimeMs
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
	if o.WebsocketRespBodyBytes != nil {
		toSerialize["websocket_resp_body_bytes"] = o.WebsocketRespBodyBytes
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
	if o.BotChallengeCompleteTokensIssued != nil {
		toSerialize["bot_challenge_complete_tokens_issued"] = o.BotChallengeCompleteTokensIssued
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
	if o.DdosActionDowngrade != nil {
		toSerialize["ddos_action_downgrade"] = o.DdosActionDowngrade
	}
	if o.DdosActionDowngradedConnections != nil {
		toSerialize["ddos_action_downgraded_connections"] = o.DdosActionDowngradedConnections
	}
	if o.VclOnComputeHitRequests != nil {
		toSerialize["vcl_on_compute_hit_requests"] = o.VclOnComputeHitRequests
	}
	if o.VclOnComputeMissRequests != nil {
		toSerialize["vcl_on_compute_miss_requests"] = o.VclOnComputeMissRequests
	}
	if o.VclOnComputePassRequests != nil {
		toSerialize["vcl_on_compute_pass_requests"] = o.VclOnComputePassRequests
	}
	if o.VclOnComputeErrorRequests != nil {
		toSerialize["vcl_on_compute_error_requests"] = o.VclOnComputeErrorRequests
	}
	if o.VclOnComputeSynthRequests != nil {
		toSerialize["vcl_on_compute_synth_requests"] = o.VclOnComputeSynthRequests
	}
	if o.VclOnComputeEdgeHitRequests != nil {
		toSerialize["vcl_on_compute_edge_hit_requests"] = o.VclOnComputeEdgeHitRequests
	}
	if o.VclOnComputeEdgeMissRequests != nil {
		toSerialize["vcl_on_compute_edge_miss_requests"] = o.VclOnComputeEdgeMissRequests
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
	if o.ServiceID != nil {
		toSerialize["service_id"] = o.ServiceID
	}
	if o.StartTime != nil {
		toSerialize["start_time"] = o.StartTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *Results) UnmarshalJSON(bytes []byte) (err error) {
	varResults := _Results{}

	if err = json.Unmarshal(bytes, &varResults); err == nil {
		*o = Results(varResults)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "requests")
		delete(additionalProperties, "hits")
		delete(additionalProperties, "hits_time")
		delete(additionalProperties, "miss")
		delete(additionalProperties, "miss_time")
		delete(additionalProperties, "pass")
		delete(additionalProperties, "pass_time")
		delete(additionalProperties, "errors")
		delete(additionalProperties, "restarts")
		delete(additionalProperties, "hit_ratio")
		delete(additionalProperties, "bandwidth")
		delete(additionalProperties, "body_size")
		delete(additionalProperties, "header_size")
		delete(additionalProperties, "req_body_bytes")
		delete(additionalProperties, "req_header_bytes")
		delete(additionalProperties, "resp_body_bytes")
		delete(additionalProperties, "resp_header_bytes")
		delete(additionalProperties, "bereq_body_bytes")
		delete(additionalProperties, "bereq_header_bytes")
		delete(additionalProperties, "uncacheable")
		delete(additionalProperties, "pipe")
		delete(additionalProperties, "synth")
		delete(additionalProperties, "tls")
		delete(additionalProperties, "tls_v10")
		delete(additionalProperties, "tls_v11")
		delete(additionalProperties, "tls_v12")
		delete(additionalProperties, "tls_v13")
		delete(additionalProperties, "edge_requests")
		delete(additionalProperties, "edge_resp_header_bytes")
		delete(additionalProperties, "edge_resp_body_bytes")
		delete(additionalProperties, "edge_hit_requests")
		delete(additionalProperties, "edge_miss_requests")
		delete(additionalProperties, "origin_fetches")
		delete(additionalProperties, "origin_fetch_header_bytes")
		delete(additionalProperties, "origin_fetch_body_bytes")
		delete(additionalProperties, "origin_fetch_resp_header_bytes")
		delete(additionalProperties, "origin_fetch_resp_body_bytes")
		delete(additionalProperties, "origin_revalidations")
		delete(additionalProperties, "origin_cache_fetches")
		delete(additionalProperties, "shield")
		delete(additionalProperties, "shield_resp_body_bytes")
		delete(additionalProperties, "shield_resp_header_bytes")
		delete(additionalProperties, "shield_fetches")
		delete(additionalProperties, "shield_fetch_header_bytes")
		delete(additionalProperties, "shield_fetch_body_bytes")
		delete(additionalProperties, "shield_fetch_resp_header_bytes")
		delete(additionalProperties, "shield_fetch_resp_body_bytes")
		delete(additionalProperties, "shield_revalidations")
		delete(additionalProperties, "shield_cache_fetches")
		delete(additionalProperties, "ipv6")
		delete(additionalProperties, "otfp")
		delete(additionalProperties, "otfp_resp_body_bytes")
		delete(additionalProperties, "otfp_resp_header_bytes")
		delete(additionalProperties, "otfp_shield_resp_body_bytes")
		delete(additionalProperties, "otfp_shield_resp_header_bytes")
		delete(additionalProperties, "otfp_manifests")
		delete(additionalProperties, "otfp_deliver_time")
		delete(additionalProperties, "otfp_shield_time")
		delete(additionalProperties, "video")
		delete(additionalProperties, "pci")
		delete(additionalProperties, "log")
		delete(additionalProperties, "log_bytes")
		delete(additionalProperties, "http2")
		delete(additionalProperties, "http3")
		delete(additionalProperties, "waf_logged")
		delete(additionalProperties, "waf_blocked")
		delete(additionalProperties, "waf_passed")
		delete(additionalProperties, "attack_req_body_bytes")
		delete(additionalProperties, "attack_req_header_bytes")
		delete(additionalProperties, "attack_logged_req_body_bytes")
		delete(additionalProperties, "attack_logged_req_header_bytes")
		delete(additionalProperties, "attack_blocked_req_body_bytes")
		delete(additionalProperties, "attack_blocked_req_header_bytes")
		delete(additionalProperties, "attack_passed_req_body_bytes")
		delete(additionalProperties, "attack_passed_req_header_bytes")
		delete(additionalProperties, "attack_resp_synth_bytes")
		delete(additionalProperties, "imgopto")
		delete(additionalProperties, "imgopto_resp_body_bytes")
		delete(additionalProperties, "imgopto_resp_header_bytes")
		delete(additionalProperties, "imgopto_shield")
		delete(additionalProperties, "imgopto_shield_resp_body_bytes")
		delete(additionalProperties, "imgopto_shield_resp_header_bytes")
		delete(additionalProperties, "imgopto_transforms")
		delete(additionalProperties, "imgvideo")
		delete(additionalProperties, "imgvideo_frames")
		delete(additionalProperties, "imgvideo_resp_header_bytes")
		delete(additionalProperties, "imgvideo_resp_body_bytes")
		delete(additionalProperties, "imgvideo_shield_resp_header_bytes")
		delete(additionalProperties, "imgvideo_shield_resp_body_bytes")
		delete(additionalProperties, "imgvideo_shield")
		delete(additionalProperties, "imgvideo_shield_frames")
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
		delete(additionalProperties, "status_1xx")
		delete(additionalProperties, "status_2xx")
		delete(additionalProperties, "status_3xx")
		delete(additionalProperties, "status_4xx")
		delete(additionalProperties, "status_5xx")
		delete(additionalProperties, "object_size_1k")
		delete(additionalProperties, "object_size_10k")
		delete(additionalProperties, "object_size_100k")
		delete(additionalProperties, "object_size_1m")
		delete(additionalProperties, "object_size_10m")
		delete(additionalProperties, "object_size_100m")
		delete(additionalProperties, "object_size_1g")
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
		delete(additionalProperties, "tls_handshake_sent_bytes")
		delete(additionalProperties, "hit_resp_body_bytes")
		delete(additionalProperties, "miss_resp_body_bytes")
		delete(additionalProperties, "pass_resp_body_bytes")
		delete(additionalProperties, "segblock_origin_fetches")
		delete(additionalProperties, "segblock_shield_fetches")
		delete(additionalProperties, "compute_requests")
		delete(additionalProperties, "compute_request_time_ms")
		delete(additionalProperties, "compute_request_time_billed_ms")
		delete(additionalProperties, "compute_ram_used")
		delete(additionalProperties, "compute_execution_time_ms")
		delete(additionalProperties, "compute_req_header_bytes")
		delete(additionalProperties, "compute_req_body_bytes")
		delete(additionalProperties, "compute_resp_header_bytes")
		delete(additionalProperties, "compute_resp_body_bytes")
		delete(additionalProperties, "compute_resp_status_1xx")
		delete(additionalProperties, "compute_resp_status_2xx")
		delete(additionalProperties, "compute_resp_status_3xx")
		delete(additionalProperties, "compute_resp_status_4xx")
		delete(additionalProperties, "compute_resp_status_5xx")
		delete(additionalProperties, "compute_bereq_header_bytes")
		delete(additionalProperties, "compute_bereq_body_bytes")
		delete(additionalProperties, "compute_beresp_header_bytes")
		delete(additionalProperties, "compute_beresp_body_bytes")
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
		delete(additionalProperties, "websocket_resp_body_bytes")
		delete(additionalProperties, "websocket_bereq_header_bytes")
		delete(additionalProperties, "websocket_bereq_body_bytes")
		delete(additionalProperties, "websocket_beresp_header_bytes")
		delete(additionalProperties, "websocket_beresp_body_bytes")
		delete(additionalProperties, "websocket_conn_time_ms")
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
		delete(additionalProperties, "bot_challenge_complete_tokens_issued")
		delete(additionalProperties, "bot_challenges_issued")
		delete(additionalProperties, "bot_challenges_succeeded")
		delete(additionalProperties, "bot_challenges_failed")
		delete(additionalProperties, "ddos_action_downgrade")
		delete(additionalProperties, "ddos_action_downgraded_connections")
		delete(additionalProperties, "vcl_on_compute_hit_requests")
		delete(additionalProperties, "vcl_on_compute_miss_requests")
		delete(additionalProperties, "vcl_on_compute_pass_requests")
		delete(additionalProperties, "vcl_on_compute_error_requests")
		delete(additionalProperties, "vcl_on_compute_synth_requests")
		delete(additionalProperties, "vcl_on_compute_edge_hit_requests")
		delete(additionalProperties, "vcl_on_compute_edge_miss_requests")
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
		delete(additionalProperties, "service_id")
		delete(additionalProperties, "start_time")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableResults is a helper abstraction for handling nullable results types.
type NullableResults struct {
	value *Results
	isSet bool
}

// Get returns the value.
func (v NullableResults) Get() *Results {
	return v.value
}

// Set modifies the value.
func (v *NullableResults) Set(val *Results) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableResults) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableResults) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableResults returns a pointer to a new instance of NullableResults.
func NewNullableResults(val *Results) *NullableResults {
	return &NullableResults{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableResults) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableResults) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
