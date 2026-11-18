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
	"time"
)

// PlatformValues Metric values for the timestamp field. Note that this dataset is sparse: only the keys with non-zero values will be included in the record.
type PlatformValues struct {
	// Timestamp for the start of the time period being reported
	Timestamp *time.Time `json:"timestamp,omitempty"`
	// Datacenter for this data point. Only present if the request was made with `group_by=datacenter`.
	Datacenter *string `json:"datacenter,omitempty"`
	// Region for this data point. Only present if the request was made with `group_by=region`.
	Region *string `json:"region,omitempty"`
	// Number of requests processed.
	Requests *int64 `json:"requests,omitempty"`
	// Number of cache hits.
	Hits *int64 `json:"hits,omitempty"`
	// Total amount of time spent processing cache hits (in seconds).
	HitsTime *float32 `json:"hits_time,omitempty"`
	// Number of cache misses.
	Miss *int64 `json:"miss,omitempty"`
	// Total amount of time spent processing cache misses (in seconds).
	MissTime *float32 `json:"miss_time,omitempty"`
	// Number of requests that passed through the CDN without being cached.
	Pass *int64 `json:"pass,omitempty"`
	// Total amount of time spent processing cache passes (in seconds).
	PassTime *float32 `json:"pass_time,omitempty"`
	// Number of cache errors.
	Errors *int64 `json:"errors,omitempty"`
	// Number of restarts performed.
	Restarts *int64 `json:"restarts,omitempty"`
	// Ratio of cache hits to cache misses (between 0 and 1).
	HitRatio NullableFloat32 `json:"hit_ratio,omitempty"`
	// Total bytes delivered (`resp_header_bytes` + `resp_body_bytes` + `bereq_header_bytes` + `bereq_body_bytes` + `compute_resp_header_bytes` + `compute_resp_body_bytes` + `compute_bereq_header_bytes` + `compute_bereq_body_bytes` + `websocket_resp_header_bytes` + `websocket_resp_body_bytes` + `websocket_bereq_header_bytes` + `websocket_bereq_body_bytes` + `fanout_resp_header_bytes` + `fanout_resp_body_bytes` + `fanout_bereq_header_bytes` + `fanout_bereq_body_bytes`).
	Bandwidth *int64 `json:"bandwidth,omitempty"`
	// Total body bytes received.
	ReqBodyBytes *int64 `json:"req_body_bytes,omitempty"`
	// Total header bytes received.
	ReqHeaderBytes *int64 `json:"req_header_bytes,omitempty"`
	// Total body bytes delivered (edge_resp_body_bytes + shield_resp_body_bytes).
	RespBodyBytes *int64 `json:"resp_body_bytes,omitempty"`
	// Total header bytes delivered (edge_resp_header_bytes + shield_resp_header_bytes).
	RespHeaderBytes *int64 `json:"resp_header_bytes,omitempty"`
	// Total body bytes sent to origin.
	BereqBodyBytes *int64 `json:"bereq_body_bytes,omitempty"`
	// Total header bytes sent to origin.
	BereqHeaderBytes *int64 `json:"bereq_header_bytes,omitempty"`
	// Number of requests that were designated uncachable.
	Uncacheable *int64 `json:"uncacheable,omitempty"`
	// Optional. Pipe operations performed (legacy feature).
	Pipe *int64 `json:"pipe,omitempty"`
	// Number of requests that returned a synthetic response (i.e., response objects created with the `synthetic` VCL statement).
	Synth *int64 `json:"synth,omitempty"`
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
	// Number of requests sent by end users to Fastly.
	EdgeRequests *int64 `json:"edge_requests,omitempty"`
	// Total header bytes delivered from Fastly to the end user.
	EdgeRespHeaderBytes *int64 `json:"edge_resp_header_bytes,omitempty"`
	// Total body bytes delivered from Fastly to the end user.
	EdgeRespBodyBytes *int64 `json:"edge_resp_body_bytes,omitempty"`
	// Number of requests sent by end users to Fastly that resulted in a hit at the edge.
	EdgeHitRequests *int64 `json:"edge_hit_requests,omitempty"`
	// Number of requests sent by end users to Fastly that resulted in a miss at the edge.
	EdgeMissRequests *int64 `json:"edge_miss_requests,omitempty"`
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
	// Number of responses received from origin with a `304` status code in response to an `If-Modified-Since` or `If-None-Match` request. Under regular scenarios, a revalidation will imply a cache hit. However, if using Fastly Image Optimizer or segmented caching this may result in a cache miss.
	OriginRevalidations *int64 `json:"origin_revalidations,omitempty"`
	// The total number of completed requests made to backends (origins) that returned cacheable content.
	OriginCacheFetches *int64 `json:"origin_cache_fetches,omitempty"`
	// Number of requests from edge to the shield POP.
	Shield *int64 `json:"shield,omitempty"`
	// Total body bytes delivered via a shield.
	ShieldRespBodyBytes *int64 `json:"shield_resp_body_bytes,omitempty"`
	// Total header bytes delivered via a shield.
	ShieldRespHeaderBytes *int64 `json:"shield_resp_header_bytes,omitempty"`
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
	// Number of responses received from origin with a `304` status code, in response to an `If-Modified-Since` or `If-None-Match` request to a shield. Under regular scenarios, a revalidation will imply a cache hit. However, if using segmented caching this may result in a cache miss.
	ShieldRevalidations *int64 `json:"shield_revalidations,omitempty"`
	// The total number of completed requests made to shields that returned cacheable content.
	ShieldCacheFetches *int64 `json:"shield_cache_fetches,omitempty"`
	// Number of requests that were received over IPv6.
	Ipv6 *int64 `json:"ipv6,omitempty"`
	// Number of responses that came from the Fastly On-the-Fly Packaging service for video-on-demand.
	Otfp *int64 `json:"otfp,omitempty"`
	// Total body bytes delivered from the Fastly On-the-Fly Packaging service for video-on-demand.
	OtfpRespBodyBytes *int64 `json:"otfp_resp_body_bytes,omitempty"`
	// Total header bytes delivered from the Fastly On-the-Fly Packaging service for video-on-demand.
	OtfpRespHeaderBytes *int64 `json:"otfp_resp_header_bytes,omitempty"`
	// Total body bytes delivered via a shield for the Fastly On-the-Fly Packaging service for video-on-demand.
	OtfpShieldRespBodyBytes *int64 `json:"otfp_shield_resp_body_bytes,omitempty"`
	// Total header bytes delivered via a shield for the Fastly On-the-Fly Packaging service for video-on-demand.
	OtfpShieldRespHeaderBytes *int64 `json:"otfp_shield_resp_header_bytes,omitempty"`
	// Number of responses that were manifest files from the Fastly On-the-Fly Packaging service for video-on-demand.
	OtfpManifests *int64 `json:"otfp_manifests,omitempty"`
	// Total amount of time spent delivering a response from the Fastly On-the-Fly Packaging service for video-on-demand (in seconds).
	OtfpDeliverTime *float32 `json:"otfp_deliver_time,omitempty"`
	// Total amount of time spent delivering a response via a shield from the Fastly On-the-Fly Packaging service for video-on-demand (in seconds).
	OtfpShieldTime *float32 `json:"otfp_shield_time,omitempty"`
	// Number of responses with the video segment or video manifest MIME type (i.e., application/x-mpegurl, application/vnd.apple.mpegurl, application/f4m, application/dash+xml, application/vnd.ms-sstr+xml, ideo/mp2t, audio/aac, video/f4f, video/x-flv, video/mp4, audio/mp4).
	Video *int64 `json:"video,omitempty"`
	// Number of responses with the PCI flag turned on.
	Pci *int64 `json:"pci,omitempty"`
	// Number of log lines sent.
	Log *int64 `json:"log,omitempty"`
	// Total log bytes sent.
	LogBytes *int64 `json:"log_bytes,omitempty"`
	// Number of requests received over HTTP/2.
	HTTP2 *int64 `json:"http2,omitempty"`
	// Number of requests received over HTTP/3.
	HTTP3 *int64 `json:"http3,omitempty"`
	// Total body bytes received from requests that triggered a WAF rule.
	AttackReqBodyBytes *int64 `json:"attack_req_body_bytes,omitempty"`
	// Total header bytes received from requests that triggered a WAF rule.
	AttackReqHeaderBytes *int64 `json:"attack_req_header_bytes,omitempty"`
	// Total body bytes received from requests that triggered a WAF rule that was logged.
	AttackLoggedReqBodyBytes *int64 `json:"attack_logged_req_body_bytes,omitempty"`
	// Total header bytes received from requests that triggered a WAF rule that was logged.
	AttackLoggedReqHeaderBytes *int64 `json:"attack_logged_req_header_bytes,omitempty"`
	// Total body bytes received from requests that triggered a WAF rule that was blocked.
	AttackBlockedReqBodyBytes *int64 `json:"attack_blocked_req_body_bytes,omitempty"`
	// Total header bytes received from requests that triggered a WAF rule that was blocked.
	AttackBlockedReqHeaderBytes *int64 `json:"attack_blocked_req_header_bytes,omitempty"`
	// Total body bytes received from requests that triggered a WAF rule that was passed.
	AttackPassedReqBodyBytes *int64 `json:"attack_passed_req_body_bytes,omitempty"`
	// Total header bytes received from requests that triggered a WAF rule that was passed.
	AttackPassedReqHeaderBytes *int64 `json:"attack_passed_req_header_bytes,omitempty"`
	// Total bytes delivered for requests that triggered a WAF rule and returned a synthetic response.
	AttackRespSynthBytes *int64 `json:"attack_resp_synth_bytes,omitempty"`
	// Number of responses that came from the Fastly Image Optimizer service. If the service receives 10 requests for an image, this stat will be 10 regardless of how many times the image was transformed.
	Imgopto *int64 `json:"imgopto,omitempty"`
	// The total image bytes that would have been delivered if the images had not been transformed by the Fastly Image Optimizer. Bytes are counted once per client request. Byte counts exclude shield traffic and requests where Fastly Image Optimizer proxied a request without attempting to transform it (e.g., file types that are not supported by Image Optimizer).
	ImgoptoInputImageBytes *int64 `json:"imgopto_input_image_bytes,omitempty"`
	// The total image bytes delivered that have been transformed by the Fastly Image Optimizer. Bytes are counted once per client request. Byte counts exclude shield traffic and requests where Fastly Image Optimizer proxied a request without attempting to transform it (e.g., file types that are not supported by Image Optimizer).
	ImgoptoOutputImageBytes *int64 `json:"imgopto_output_image_bytes,omitempty"`
	// Total body bytes delivered from the Fastly Image Optimizer service, including shield traffic.
	ImgoptoRespBodyBytes *int64 `json:"imgopto_resp_body_bytes,omitempty"`
	// Total header bytes delivered from the Fastly Image Optimizer service, including shield traffic.
	ImgoptoRespHeaderBytes *int64 `json:"imgopto_resp_header_bytes,omitempty"`
	// Number of responses that came from the Fastly Image Optimizer service via a shield.
	ImgoptoShield *int64 `json:"imgopto_shield,omitempty"`
	// Total body bytes delivered via a shield from the Fastly Image Optimizer service.
	ImgoptoShieldRespBodyBytes *int64 `json:"imgopto_shield_resp_body_bytes,omitempty"`
	// Total header bytes delivered via a shield from the Fastly Image Optimizer service.
	ImgoptoShieldRespHeaderBytes *int64 `json:"imgopto_shield_resp_header_bytes,omitempty"`
	// Number of transforms performed by the Fastly Image Optimizer service.
	ImgoptoTransforms *int64 `json:"imgopto_transforms,omitempty"`
	// Number of video responses that came from the Fastly Image Optimizer service.
	Imgvideo *int64 `json:"imgvideo,omitempty"`
	// Number of video frames that came from the Fastly Image Optimizer service. A video frame is an individual image within a sequence of video.
	ImgvideoFrames *int64 `json:"imgvideo_frames,omitempty"`
	// Total header bytes of video delivered from the Fastly Image Optimizer service.
	ImgvideoRespHeaderBytes *int64 `json:"imgvideo_resp_header_bytes,omitempty"`
	// Total body bytes of video delivered from the Fastly Image Optimizer service.
	ImgvideoRespBodyBytes *int64 `json:"imgvideo_resp_body_bytes,omitempty"`
	// Total header bytes of video delivered via a shield from the Fastly Image Optimizer service.
	ImgvideoShieldRespHeaderBytes *int64 `json:"imgvideo_shield_resp_header_bytes,omitempty"`
	// Total body bytes of video delivered via a shield from the Fastly Image Optimizer service.
	ImgvideoShieldRespBodyBytes *int64 `json:"imgvideo_shield_resp_body_bytes,omitempty"`
	// Number of video responses delivered via a shield that came from the Fastly Image Optimizer service.
	ImgvideoShield *int64 `json:"imgvideo_shield,omitempty"`
	// Number of video frames delivered via a shield that came from the Fastly Image Optimizer service. A video frame is an individual image within a sequence of video.
	ImgvideoShieldFrames *int64 `json:"imgvideo_shield_frames,omitempty"`
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
	// Number of responses sent with status code 530.
	Status530 *int64 `json:"status_530,omitempty"`
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
	// Time spent inside the `vcl_recv` Varnish subroutine (in seconds).
	RecvSubTime *float32 `json:"recv_sub_time,omitempty"`
	// Number of executions of the `vcl_recv` Varnish subroutine.
	RecvSubCount *int64 `json:"recv_sub_count,omitempty"`
	// Time spent inside the `vcl_hash` Varnish subroutine (in seconds).
	HashSubTime *float32 `json:"hash_sub_time,omitempty"`
	// Number of executions of the `vcl_hash` Varnish subroutine.
	HashSubCount *int64 `json:"hash_sub_count,omitempty"`
	// Time spent inside the `vcl_miss` Varnish subroutine (in seconds).
	MissSubTime *float32 `json:"miss_sub_time,omitempty"`
	// Number of executions of the `vcl_miss` Varnish subroutine.
	MissSubCount *int64 `json:"miss_sub_count,omitempty"`
	// Time spent inside the `vcl_fetch` Varnish subroutine (in seconds).
	FetchSubTime *float32 `json:"fetch_sub_time,omitempty"`
	// Number of executions of the `vcl_fetch` Varnish subroutine.
	FetchSubCount *int64 `json:"fetch_sub_count,omitempty"`
	// Time spent inside the `vcl_pass` Varnish subroutine (in seconds).
	PassSubTime *float32 `json:"pass_sub_time,omitempty"`
	// Number of executions of the `vcl_pass` Varnish subroutine.
	PassSubCount *int64 `json:"pass_sub_count,omitempty"`
	// Time spent inside the `vcl_pipe` Varnish subroutine (in seconds).
	PipeSubTime *float32 `json:"pipe_sub_time,omitempty"`
	// Number of executions of the `vcl_pipe` Varnish subroutine.
	PipeSubCount *int64 `json:"pipe_sub_count,omitempty"`
	// Time spent inside the `vcl_deliver` Varnish subroutine (in seconds).
	DeliverSubTime *float32 `json:"deliver_sub_time,omitempty"`
	// Number of executions of the `vcl_deliver` Varnish subroutine.
	DeliverSubCount *int64 `json:"deliver_sub_count,omitempty"`
	// Time spent inside the `vcl_error` Varnish subroutine (in seconds).
	ErrorSubTime *float32 `json:"error_sub_time,omitempty"`
	// Number of executions of the `vcl_error` Varnish subroutine.
	ErrorSubCount *int64 `json:"error_sub_count,omitempty"`
	// Time spent inside the `vcl_hit` Varnish subroutine (in seconds).
	HitSubTime *float32 `json:"hit_sub_time,omitempty"`
	// Number of executions of the `vcl_hit` Varnish subroutine.
	HitSubCount *int64 `json:"hit_sub_count,omitempty"`
	// Time spent inside the `vcl_prehash` Varnish subroutine (in seconds).
	PrehashSubTime *float32 `json:"prehash_sub_time,omitempty"`
	// Number of executions of the `vcl_prehash` Varnish subroutine.
	PrehashSubCount *int64 `json:"prehash_sub_count,omitempty"`
	// Time spent inside the `vcl_predeliver` Varnish subroutine (in seconds).
	PredeliverSubTime *float32 `json:"predeliver_sub_time,omitempty"`
	// Number of executions of the `vcl_predeliver` Varnish subroutine.
	PredeliverSubCount *int64 `json:"predeliver_sub_count,omitempty"`
	// Number of bytes transferred during TLS handshake.
	TLSHandshakeSentBytes *int64 `json:"tls_handshake_sent_bytes,omitempty"`
	// Total body bytes delivered for cache hits.
	HitRespBodyBytes *int64 `json:"hit_resp_body_bytes,omitempty"`
	// Total body bytes delivered for cache misses.
	MissRespBodyBytes *int64 `json:"miss_resp_body_bytes,omitempty"`
	// Total body bytes delivered for cache passes.
	PassRespBodyBytes *int64 `json:"pass_resp_body_bytes,omitempty"`
	// Number of `Range` requests to origin for segments of resources when using segmented caching.
	SegblockOriginFetches *int64 `json:"segblock_origin_fetches,omitempty"`
	// Number of `Range` requests to a shield for segments of resources when using segmented caching.
	SegblockShieldFetches *int64 `json:"segblock_shield_fetches,omitempty"`
	// The total number of requests that were received for your service by Fastly.
	ComputeRequests *int64 `json:"compute_requests,omitempty"`
	// The total, actual amount of time used to process your requests, including active CPU time (in milliseconds).
	ComputeRequestTimeMs *float32 `json:"compute_request_time_ms,omitempty"`
	// The total amount of request processing time you will be billed for, measured in 50 millisecond increments.
	ComputeRequestTimeBilledMs *float32 `json:"compute_request_time_billed_ms,omitempty"`
	// The amount of RAM used for your service by Fastly (in bytes).
	ComputeRAMUsed *int64 `json:"compute_ram_used,omitempty"`
	// The amount of active CPU time used to process your requests (in milliseconds).
	ComputeExecutionTimeMs *float32 `json:"compute_execution_time_ms,omitempty"`
	// Total header bytes received by the Compute platform.
	ComputeReqHeaderBytes *int64 `json:"compute_req_header_bytes,omitempty"`
	// Total body bytes received by the Compute platform.
	ComputeReqBodyBytes *int64 `json:"compute_req_body_bytes,omitempty"`
	// Total header bytes sent from Compute to end user.
	ComputeRespHeaderBytes *int64 `json:"compute_resp_header_bytes,omitempty"`
	// Total body bytes sent from Compute to end user.
	ComputeRespBodyBytes *int64 `json:"compute_resp_body_bytes,omitempty"`
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
	// Total header bytes sent to backends (origins) by the Compute platform.
	ComputeBereqHeaderBytes *int64 `json:"compute_bereq_header_bytes,omitempty"`
	// Total body bytes sent to backends (origins) by the Compute platform.
	ComputeBereqBodyBytes *int64 `json:"compute_bereq_body_bytes,omitempty"`
	// Total header bytes received from backends (origins) by the Compute platform.
	ComputeBerespHeaderBytes *int64 `json:"compute_beresp_header_bytes,omitempty"`
	// Total body bytes received from backends (origins) by the Compute platform.
	ComputeBerespBodyBytes *int64 `json:"compute_beresp_body_bytes,omitempty"`
	// Number of backend requests started.
	ComputeBereqs *int64 `json:"compute_bereqs,omitempty"`
	// Number of backend request errors, including timeouts.
	ComputeServiceBereqError *int64 `json:"compute_service_bereq_error,omitempty"`
	// Number of times a guest exceeded its heap limit.
	ComputeServiceMemoryExceededError *int64 `json:"compute_service_memory_exceeded_error,omitempty"`
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
	// Total message content bytes sent to end users over passthrough WebSocket connections.
	WebsocketRespBodyBytes *int64 `json:"websocket_resp_body_bytes,omitempty"`
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
	// Total published messages received from the publish API endpoint.
	FanoutRecvPublishes *int64 `json:"fanout_recv_publishes,omitempty"`
	// Total published messages sent to end users.
	FanoutSendPublishes *int64 `json:"fanout_send_publishes,omitempty"`
	// The total number of class a operations for the KV store.
	KvStoreClassAOperations *int64 `json:"kv_store_class_a_operations,omitempty"`
	// The total number of class b operations for the KV store.
	KvStoreClassBOperations *int64 `json:"kv_store_class_b_operations,omitempty"`
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
	// The number of challenge-complete tokens issued. For example, issuing a challenge-complete token after a series of CAPTCHA challenges ending in success.
	BotChallengeCompleteTokensIssued *int64 `json:"bot_challenge_complete_tokens_issued,omitempty"`
	// The number of challenges issued. For example, the issuance of a CAPTCHA challenge.
	BotChallengesIssued *int64 `json:"bot_challenges_issued,omitempty"`
	// The number of successful challenge solutions processed. For example, a correct CAPTCHA solution.
	BotChallengesSucceeded *int64 `json:"bot_challenges_succeeded,omitempty"`
	// The number of failed challenge solutions processed. For example, an incorrect CAPTCHA solution.
	BotChallengesFailed *int64 `json:"bot_challenges_failed,omitempty"`
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
	DdosProtectionRequestsDetectCount *int64 `json:"ddos_protection_requests_detect_count,omitempty"`
	// Number of requests classified as a DDoS attack against a customer origin or service that were mitigated by the Fastly platform.
	DdosProtectionRequestsMitigateCount *int64 `json:"ddos_protection_requests_mitigate_count,omitempty"`
	// Number of requests analyzed for DDoS attacks against a customer origin or service, but with no DDoS detected.
	DdosProtectionRequestsAllowCount *int64 `json:"ddos_protection_requests_allow_count,omitempty"`
	// A count of the number of Class A Object Storage operations.
	ObjectStorageClassAOperationsCount *int64 `json:"object_storage_class_a_operations_count,omitempty"`
	// A count of the number of Class B Object Storage operations.
	ObjectStorageClassBOperationsCount *int64 `json:"object_storage_class_b_operations_count,omitempty"`
	// Number of requests received by AI Accelerator.
	AiaRequests *int64 `json:"aia_requests,omitempty"`
	// Number of \"Informational\" category status codes received from AI provider.
	AiaStatus1xx *int64 `json:"aia_status_1xx,omitempty"`
	// Number of \"Success\" status codes received from AI provider.
	AiaStatus2xx *int64 `json:"aia_status_2xx,omitempty"`
	// Number of \"Redirection\" received from AI provider.
	AiaStatus3xx *int64 `json:"aia_status_3xx,omitempty"`
	// Number of \"Client Error\" received from AI provider.
	AiaStatus4xx *int64 `json:"aia_status_4xx,omitempty"`
	// Number of \"Server Error\" received from AI provider.
	AiaStatus5xx *int64 `json:"aia_status_5xx,omitempty"`
	// The usage tokens associated with the response returned from the AI Accelerator cache.
	AiaResponseUsageTokens *int64 `json:"aia_response_usage_tokens,omitempty"`
	// The number of usage tokens reported by the request to the origin from AI Accelerator.
	AiaOriginUsageTokens *int64 `json:"aia_origin_usage_tokens,omitempty"`
	// The estimated amount of time saved by responses served from the AI Accelerator semantic cache.
	AiaEstimatedTimeSavedMs *int64 `json:"aia_estimated_time_saved_ms,omitempty"`
	// Number of requests that were collapsed and satisfied by a usable cache object.
	RequestCollapseUsableCount *int64 `json:"request_collapse_usable_count,omitempty"`
	// Number of requests that were collapsed and unable to be satisfied by the resulting cache object.
	RequestCollapseUnusableCount *int64 `json:"request_collapse_unusable_count,omitempty"`
	// Number of cache operations executed by the Compute platform.
	ComputeCacheOperationsCount *int64 `json:"compute_cache_operations_count,omitempty"`
	// Total number of Next-Gen WAF (Edge WAF & Core WAF) requests.
	NgwafRequestsTotalCount *int32 `json:"ngwaf_requests_total_count,omitempty"`
	// Count of Edge WAF requests with an unknown outcome.
	NgwafRequestsUnknownCount *int32 `json:"ngwaf_requests_unknown_count,omitempty"`
	// Count of Next-Gen WAF (Edge WAF & Core WAF) requests allowed.
	NgwafRequestsAllowedCount *int32 `json:"ngwaf_requests_allowed_count,omitempty"`
	// Count of Edge WAF requests logged.
	NgwafRequestsLoggedCount *int32 `json:"ngwaf_requests_logged_count,omitempty"`
	// Count of Next-Gen WAF (Edge WAF & Core WAF) requests blocked.
	NgwafRequestsBlockedCount *int32 `json:"ngwaf_requests_blocked_count,omitempty"`
	// Count of Edge WAF requests timed outcome.
	NgwafRequestsTimeoutCount *int32 `json:"ngwaf_requests_timeout_count,omitempty"`
	// Count of Edge WAF requests challenged.
	NgwafRequestsChallengedCount *int32 `json:"ngwaf_requests_challenged_count,omitempty"`
	// Number of requests processed by the API Discovery engine.
	ApiDiscoveryRequestsCount *int32 `json:"api_discovery_requests_count,omitempty"`
	// Number of responses delivered with status code 103 (Early Hints) by the Compute platform.
	ComputeRespStatus103 *int32 `json:"compute_resp_status_103,omitempty"`
	// Number of responses delivered with status code 200 (Success) by the Compute platform.
	ComputeRespStatus200 *int32 `json:"compute_resp_status_200,omitempty"`
	// Number of responses delivered with status code 204 (No Content) by the Compute platform.
	ComputeRespStatus204 *int32 `json:"compute_resp_status_204,omitempty"`
	// Number of responses delivered with status code 206 (Partial Content) by the Compute platform.
	ComputeRespStatus206 *int32 `json:"compute_resp_status_206,omitempty"`
	// Number of responses delivered with status code 301 (Moved Permanently) by the Compute platform.
	ComputeRespStatus301 *int32 `json:"compute_resp_status_301,omitempty"`
	// Number of responses delivered with status code 302 (Found) by the Compute platform.
	ComputeRespStatus302 *int32 `json:"compute_resp_status_302,omitempty"`
	// Number of responses delivered with status code 304 (Not Modified) by the Compute platform.
	ComputeRespStatus304 *int32 `json:"compute_resp_status_304,omitempty"`
	// Number of responses delivered with status code 400 (Bad Request) by the Compute platform.
	ComputeRespStatus400 *int32 `json:"compute_resp_status_400,omitempty"`
	// Number of responses delivered with status code 401 (Unauthorized) by the Compute platform.
	ComputeRespStatus401 *int32 `json:"compute_resp_status_401,omitempty"`
	// Number of responses delivered with status code 403 (Forbidden) by the Compute platform.
	ComputeRespStatus403 *int32 `json:"compute_resp_status_403,omitempty"`
	// Number of responses delivered with status code 404 (Not Found) by the Compute platform.
	ComputeRespStatus404 *int32 `json:"compute_resp_status_404,omitempty"`
	// Number of responses delivered with status code 416 (Range Not Satisfiable) by the Compute platform.
	ComputeRespStatus416 *int32 `json:"compute_resp_status_416,omitempty"`
	// Number of responses delivered with status code 429 (Too Many Requests) by the Compute platform.
	ComputeRespStatus429 *int32 `json:"compute_resp_status_429,omitempty"`
	// Number of responses delivered with status code 500 (Internal Server Error) by the Compute platform.
	ComputeRespStatus500 *int32 `json:"compute_resp_status_500,omitempty"`
	// Number of responses delivered with status code 501 (Not Implemented) by the Compute platform.
	ComputeRespStatus501 *int32 `json:"compute_resp_status_501,omitempty"`
	// Number of responses delivered with status code 502 (Bad Gateway) by the Compute platform.
	ComputeRespStatus502 *int32 `json:"compute_resp_status_502,omitempty"`
	// Number of responses delivered with status code 503 (Service Unavailable) by the Compute platform.
	ComputeRespStatus503 *int32 `json:"compute_resp_status_503,omitempty"`
	// Number of responses delivered with status code 504 (Gateway Timeout) by the Compute platform.
	ComputeRespStatus504 *int32 `json:"compute_resp_status_504,omitempty"`
	// Number of responses delivered with status code 505 (HTTP Version Not Supported) by the Compute platform.
	ComputeRespStatus505 *int32 `json:"compute_resp_status_505,omitempty"`
	// Number of responses delivered with status code 530 by the Compute platform.
	ComputeRespStatus530 *int32 `json:"compute_resp_status_530,omitempty"`
	// The number of Image Optimizer requests made from Compute services.
	ImgoptoComputeRequests *int32 `json:"imgopto_compute_requests,omitempty"`
	// Number of billable DNS responses (e.g., A, CNAME).
	DNSBillableResponsesCount *int32 `json:"dns_billable_responses_count,omitempty"`
	// Number of non-billable DNS responses (e.g., NODATA, NXDOMAIN).
	DNSNonbillableResponsesCount *int32 `json:"dns_nonbillable_responses_count,omitempty"`
	// Number of requests that resulted in a WebSocket upgrade.
	Upgrade *int32 `json:"upgrade,omitempty"`
	// Count of Next-Gen WAF Bot Management requests.
	NgwafBotAnalysisRequestCount *int32 `json:"ngwaf_bot_analysis_request_count,omitempty"`
	// Count of AVIF images delivered to end user by Image Optimizer.
	ImgoptoAvifCount *int32 `json:"imgopto_avif_count,omitempty"`
	// Count of JPEG images delivered to end user by Image Optimizer.
	ImgoptoJpegCount *int32 `json:"imgopto_jpeg_count,omitempty"`
	// Count of PNG images delivered to end user by Image Optimizer.
	ImgoptoPngCount *int32 `json:"imgopto_png_count,omitempty"`
	// Count of GIF images delivered to end user by Image Optimizer.
	ImgoptoGifCount *int32 `json:"imgopto_gif_count,omitempty"`
	// Count of WebP images delivered to end user by Image Optimizer.
	ImgoptoWebpCount *int32 `json:"imgopto_webp_count,omitempty"`
	// Count of JPEGXL images delivered to end user by Image Optimizer.
	ImgoptoJpegxlCount *int32 `json:"imgopto_jpegxl_count,omitempty"`
	// Count of SVG images delivered to end user by Image Optimizer.
	ImgoptoSvgCount *int32 `json:"imgopto_svg_count,omitempty"`
	// Count of MP4s delivered to end user by Image Optimizer.
	ImgoptoMp4Count *int32 `json:"imgopto_mp4_count,omitempty"`
	// Aggregate of fatal errors caused by exceeding allocated resource limits, specifically runtime duration, vCPU usage, and heap memory limits.
	ComputeServiceResourceLimitsError *int32 `json:"compute_service_resource_limits_error,omitempty"`
	// Fatal errors caused by service logic faults, including stack overflows, unreachable code traps, illegal memory access, or attempts to send multiple responses.
	ComputeServiceRuntimeError *int32 `json:"compute_service_runtime_error,omitempty"`
	// Fatal errors caused by the service path exceeding hop or service limits, or where a forwarding loop is detected via CDN-Loop headers.
	ComputeServiceChainError *int32 `json:"compute_service_chain_error,omitempty"`
	// Fatal errors caused by internal errors in Fastly’s Compute platform.
	ComputePlatformInternalError *int32 `json:"compute_platform_internal_error,omitempty"`
	// Fatal errors caused by exceeding the per-request runtime limit.
	ComputeServiceTimeoutError *int32 `json:"compute_service_timeout_error,omitempty"`
	// Fatal errors caused by exceeding the per-request vCPU time limit.
	ComputeServiceVcpuExceededError *int32 `json:"compute_service_vcpu_exceeded_error,omitempty"`
	// Non-fatal errors caused by attempts to exceed defined operational limits, such as simultaneous backend requests or cache transactions.
	ComputeServiceLimitsError *int32 `json:"compute_service_limits_error,omitempty"`
	// Fatal errors caused by unprocessable requests to the service, such as requests with malformed CDN-Loop headers or invalid purge credentials.
	ComputePlatformInvalidRequestError *int32 `json:"compute_platform_invalid_request_error,omitempty"`
	// Number of WebAssembly (Wasm) sandboxes created.
	ComputeSandboxes *int32 `json:"compute_sandboxes,omitempty"`
	// Total number of Bot Management requests across all deployments.
	BotRequestsTotalCount *int32 `json:"bot_requests_total_count,omitempty"`
	// Count of edge requests where bot detection analysis was performed.
	BotEdgeRequestsAnalyzedCount *int32 `json:"bot_edge_requests_analyzed_count,omitempty"`
	// Count of edge requests where a bot was detected.
	BotEdgeRequestsDetectedCount *int32 `json:"bot_edge_requests_detected_count,omitempty"`
	// Count of edge requests where a verified bot was detected.
	BotEdgeRequestsVerifiedCount *int32 `json:"bot_edge_requests_verified_count,omitempty"`
	// Count of edge requests where an AI crawler was detected.
	BotEdgeRequestsAiCrawlerCount *int32 `json:"bot_edge_requests_ai_crawler_count,omitempty"`
	// Count of edge requests where an AI fetcher was detected.
	BotEdgeRequestsAiFetcherCount *int32 `json:"bot_edge_requests_ai_fetcher_count,omitempty"`
	// Count of edge requests where an accessibility bot was detected.
	BotEdgeRequestsAccessibilityCount *int32 `json:"bot_edge_requests_accessibility_count,omitempty"`
	// Count of edge requests where a content fetcher was detected.
	BotEdgeRequestsContentFetcherCount *int32 `json:"bot_edge_requests_content_fetcher_count,omitempty"`
	// Count of edge requests where a monitoring and site tool was detected.
	BotEdgeRequestsMonitoringCount *int32 `json:"bot_edge_requests_monitoring_count,omitempty"`
	// Count of edge requests where an online marketing bot was detected.
	BotEdgeRequestsOnlineMarketingCount *int32 `json:"bot_edge_requests_online_marketing_count,omitempty"`
	// Count of edge requests where a page preview bot was detected.
	BotEdgeRequestsPagePreviewCount *int32 `json:"bot_edge_requests_page_preview_count,omitempty"`
	// Count of edge requests where a platform integration was detected.
	BotEdgeRequestsPlatformIntegrationsCount *int32 `json:"bot_edge_requests_platform_integrations_count,omitempty"`
	// Count of edge requests where a research bot was detected.
	BotEdgeRequestsResearchCount *int32 `json:"bot_edge_requests_research_count,omitempty"`
	// Count of edge requests where a search engine crawler was detected.
	BotEdgeRequestsSearchEngineCrawlerCount *int32 `json:"bot_edge_requests_search_engine_crawler_count,omitempty"`
	// Count of edge requests where a search engine optimization bot was detected.
	BotEdgeRequestsSearchEngineOptimizationCount *int32 `json:"bot_edge_requests_search_engine_optimization_count,omitempty"`
	// Count of edge requests where a security tool was detected.
	BotEdgeRequestsSecurityToolsCount *int32 `json:"bot_edge_requests_security_tools_count,omitempty"`
	// The number of times Compute has handed off a request to the Fanout proxy or WebSocket proxy.
	ComputeHandoff *int32 `json:"compute_handoff,omitempty"`
	// Number of backend requests from a Compute service that failed during DNS resolution.
	ComputeServiceBereqDNSError *int32 `json:"compute_service_bereq_dns_error,omitempty"`
	// Number of backend requests from a Compute service where the connection to the origin timed out before being established.
	ComputeServiceBereqConnTimeoutError *int32 `json:"compute_service_bereq_conn_timeout_error,omitempty"`
	// Number of backend requests from a Compute service where the origin actively refused the connection.
	ComputeServiceBereqConnRefusedError *int32 `json:"compute_service_bereq_conn_refused_error,omitempty"`
	// Number of backend requests from a Compute service that failed due to a connection error not classified as a timeout or refusal.
	ComputeServiceBereqConnOtherError *int32 `json:"compute_service_bereq_conn_other_error,omitempty"`
	// Number of backend requests from a Compute service that failed due to a TLS certificate validation error (e.g., expired, untrusted CA, hostname mismatch).
	ComputeServiceBereqTLSServerCertError *int32 `json:"compute_service_bereq_tls_server_cert_error,omitempty"`
	// Number of backend requests from a Compute service that failed due to a TLS error not classified as a certificate error.
	ComputeServiceBereqTLSOtherError *int32 `json:"compute_service_bereq_tls_other_error,omitempty"`
	// Number of backend requests from a Compute service that failed due to an HTTP/1.x protocol violation after the request was transmitted.
	ComputeServiceBereqHTTPProtoV1Error *int32 `json:"compute_service_bereq_http_proto_v1_error,omitempty"`
	// Number of backend requests from a Compute service that failed due to an HTTP/2 protocol error, typically a `RST_STREAM` or `GO_AWAY` from the origin.
	ComputeServiceBereqHTTPProtoV2Error *int32 `json:"compute_service_bereq_http_proto_v2_error,omitempty"`
	// Number of backend requests from a Compute service where the origin sent an incomplete HTTP response.
	ComputeServiceBereqHTTPIncompleteError *int32 `json:"compute_service_bereq_http_incomplete_error,omitempty"`
	// Number of backend requests from a Compute service where the origin did not respond within the configured timeout period.
	ComputeServiceBereqHTTPTimeoutError *int32 `json:"compute_service_bereq_http_timeout_error,omitempty"`
	// Number of backend requests from a Compute service that failed due to an HTTP-level error not classified in any category.
	ComputeServiceBereqHTTPOtherError *int32 `json:"compute_service_bereq_http_other_error,omitempty"`
	// Number of backend requests from a Compute service that failed due to an error not classified into the DNS, connection, TLS, or HTTP categories.
	ComputeServiceBereqOtherError *int32 `json:"compute_service_bereq_other_error,omitempty"`
	// Number of backend requests from a Compute service where the origin returned a 5xx status code.
	ComputeServiceBereq5xxError *int32 `json:"compute_service_bereq_5xx_error,omitempty"`
	// Number of backend requests from a Compute service that failed at the TCP connection level. Sum of `compute_service_bereq_conn_timeout_error`, `compute_service_bereq_conn_refused_error`, and `compute_service_bereq_conn_other_error`.
	ComputeServiceBereqConnError *int32 `json:"compute_service_bereq_conn_error,omitempty"`
	// Number of backend requests from a Compute service that failed during the TLS handshake or session with the origin. Sum of `compute_service_bereq_tls_server_cert_error` and `compute_service_bereq_tls_other_error`.
	ComputeServiceBereqTLSError *int32 `json:"compute_service_bereq_tls_error,omitempty"`
	// Number of backend requests from a Compute service that failed at the HTTP protocol level. Sum of `compute_service_bereq_http_proto_v1_error`, `compute_service_bereq_http_proto_v2_error`, `compute_service_bereq_http_incomplete_error`, `compute_service_bereq_http_timeout_error`, and `compute_service_bereq_http_other_error`.
	ComputeServiceBereqHTTPError *int32 `json:"compute_service_bereq_http_error,omitempty"`
	// Number of Private Access Token challenges issued.
	BotChallengesPatsIssued *int32 `json:"bot_challenges_pats_issued,omitempty"`
	// Number of successful Private Access Token challenge solutions processed.
	BotChallengesPatsSucceeded *int32 `json:"bot_challenges_pats_succeeded,omitempty"`
	// Number of edge requests where a headless bot was detected.
	BotEdgeRequestsHeadlessCount *int32 `json:"bot_edge_requests_headless_count,omitempty"`
	// Number of responses sent with status code 499 (Client Disconnected).
	Status499 *int32 `json:"status_499,omitempty"`
	// Number of requests received by AI Runtime Control.
	ArcRequests *int32 `json:"arc_requests,omitempty"`
	// The number of challenge-complete tokens checked while enforcing DDoS protection rules.
	DdosChallengeCompleteTokensChecked *int32 `json:"ddos_challenge_complete_tokens_checked,omitempty"`
	// The number of challenge-complete tokens that passed validation while enforcing DDoS protection rules.
	DdosChallengeCompleteTokensPassed *int32 `json:"ddos_challenge_complete_tokens_passed,omitempty"`
	// The number of challenge-complete tokens that failed validation while enforcing DDoS protection rules.
	DdosChallengeCompleteTokensFailed *int32 `json:"ddos_challenge_complete_tokens_failed,omitempty"`
	// The number of challenge-start tokens created while enforcing DDoS protection rules.
	DdosChallengeStarts *int32 `json:"ddos_challenge_starts,omitempty"`
	// Number of challenge-complete tokens issued as a result of a DDoS action. For example, issuing a challenge-complete token after a series of CAPTCHA challenges ending in success.
	DdosChallengeCompleteTokensIssued *int32 `json:"ddos_challenge_complete_tokens_issued,omitempty"`
	// Number of challenges issued from a DDoS action.
	DdosChallengesIssued *int32 `json:"ddos_challenges_issued,omitempty"`
	// Number of successful challenge solutions processed as a result of a DDoS action. For example, a correct CAPTCHA solution.
	DdosChallengesSucceeded *int32 `json:"ddos_challenges_succeeded,omitempty"`
	// Number of failed challenge solutions processed as a result of a DDoS action. For example, an incorrect CAPTCHA solution.
	DdosChallengesFailed *int32 `json:"ddos_challenges_failed,omitempty"`
	// Number of Private Access Token challenges issued as a result of a DDoS action.
	DdosChallengesPatsIssued *int32 `json:"ddos_challenges_pats_issued,omitempty"`
	// Number of successful Private Access Token challenge solutions processed as a result of a DDoS action.
	DdosChallengesPatsSucceeded *int32 `json:"ddos_challenges_pats_succeeded,omitempty"`
	// **Gauge.** The number of pages monitored by Content Security Policy.
	CspPagesGauge *int64 `json:"csp_pages_gauge,omitempty"`
	// **Gauge.** The number of Next-Gen WAF rules defined at the account level.
	NgwafAccountRulesGauge *int64 `json:"ngwaf_account_rules_gauge,omitempty"`
	// **Gauge.** The number of Next-Gen WAF workspaces in the account.
	NgwafAccountWorkspaces *int64 `json:"ngwaf_account_workspaces,omitempty"`
	// **Gauge.** The number of Next-Gen WAF rules defined in a workspace.
	NgwafWorkspaceRulesGauge *int64 `json:"ngwaf_workspace_rules_gauge,omitempty"`
	// **Gauge.** The number of Next-Gen WAF thresholds defined in a workspace.
	NgwafWorkspaceThresholdsGauge *int64 `json:"ngwaf_workspace_thresholds_gauge,omitempty"`
	// **Gauge.** The number of Object Storage buckets.
	ObjectStorageBucketsCount *int64 `json:"object_storage_buckets_count,omitempty"`
	// **Gauge.** The number of bytes stored in Object Storage.
	ObjectStorageBytesCount *int64 `json:"object_storage_bytes_count,omitempty"`
	AdditionalProperties    map[string]any
}

type _PlatformValues PlatformValues

// NewPlatformValues instantiates a new PlatformValues object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlatformValues() *PlatformValues {
	this := PlatformValues{}
	return &this
}

// NewPlatformValuesWithDefaults instantiates a new PlatformValues object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlatformValuesWithDefaults() *PlatformValues {
	this := PlatformValues{}
	return &this
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *PlatformValues) GetTimestamp() time.Time {
	if o == nil || o.Timestamp == nil {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetTimestampOk() (*time.Time, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *PlatformValues) HasTimestamp() bool {
	if o != nil && o.Timestamp != nil {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *PlatformValues) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetDatacenter returns the Datacenter field value if set, zero value otherwise.
func (o *PlatformValues) GetDatacenter() string {
	if o == nil || o.Datacenter == nil {
		var ret string
		return ret
	}
	return *o.Datacenter
}

// GetDatacenterOk returns a tuple with the Datacenter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetDatacenterOk() (*string, bool) {
	if o == nil || o.Datacenter == nil {
		return nil, false
	}
	return o.Datacenter, true
}

// HasDatacenter returns a boolean if a field has been set.
func (o *PlatformValues) HasDatacenter() bool {
	if o != nil && o.Datacenter != nil {
		return true
	}

	return false
}

// SetDatacenter gets a reference to the given string and assigns it to the Datacenter field.
func (o *PlatformValues) SetDatacenter(v string) {
	o.Datacenter = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *PlatformValues) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *PlatformValues) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *PlatformValues) SetRegion(v string) {
	o.Region = &v
}

// GetRequests returns the Requests field value if set, zero value otherwise.
func (o *PlatformValues) GetRequests() int64 {
	if o == nil || o.Requests == nil {
		var ret int64
		return ret
	}
	return *o.Requests
}

// GetRequestsOk returns a tuple with the Requests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetRequestsOk() (*int64, bool) {
	if o == nil || o.Requests == nil {
		return nil, false
	}
	return o.Requests, true
}

// HasRequests returns a boolean if a field has been set.
func (o *PlatformValues) HasRequests() bool {
	if o != nil && o.Requests != nil {
		return true
	}

	return false
}

// SetRequests gets a reference to the given int64 and assigns it to the Requests field.
func (o *PlatformValues) SetRequests(v int64) {
	o.Requests = &v
}

// GetHits returns the Hits field value if set, zero value otherwise.
func (o *PlatformValues) GetHits() int64 {
	if o == nil || o.Hits == nil {
		var ret int64
		return ret
	}
	return *o.Hits
}

// GetHitsOk returns a tuple with the Hits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetHitsOk() (*int64, bool) {
	if o == nil || o.Hits == nil {
		return nil, false
	}
	return o.Hits, true
}

// HasHits returns a boolean if a field has been set.
func (o *PlatformValues) HasHits() bool {
	if o != nil && o.Hits != nil {
		return true
	}

	return false
}

// SetHits gets a reference to the given int64 and assigns it to the Hits field.
func (o *PlatformValues) SetHits(v int64) {
	o.Hits = &v
}

// GetHitsTime returns the HitsTime field value if set, zero value otherwise.
func (o *PlatformValues) GetHitsTime() float32 {
	if o == nil || o.HitsTime == nil {
		var ret float32
		return ret
	}
	return *o.HitsTime
}

// GetHitsTimeOk returns a tuple with the HitsTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetHitsTimeOk() (*float32, bool) {
	if o == nil || o.HitsTime == nil {
		return nil, false
	}
	return o.HitsTime, true
}

// HasHitsTime returns a boolean if a field has been set.
func (o *PlatformValues) HasHitsTime() bool {
	if o != nil && o.HitsTime != nil {
		return true
	}

	return false
}

// SetHitsTime gets a reference to the given float32 and assigns it to the HitsTime field.
func (o *PlatformValues) SetHitsTime(v float32) {
	o.HitsTime = &v
}

// GetMiss returns the Miss field value if set, zero value otherwise.
func (o *PlatformValues) GetMiss() int64 {
	if o == nil || o.Miss == nil {
		var ret int64
		return ret
	}
	return *o.Miss
}

// GetMissOk returns a tuple with the Miss field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetMissOk() (*int64, bool) {
	if o == nil || o.Miss == nil {
		return nil, false
	}
	return o.Miss, true
}

// HasMiss returns a boolean if a field has been set.
func (o *PlatformValues) HasMiss() bool {
	if o != nil && o.Miss != nil {
		return true
	}

	return false
}

// SetMiss gets a reference to the given int64 and assigns it to the Miss field.
func (o *PlatformValues) SetMiss(v int64) {
	o.Miss = &v
}

// GetMissTime returns the MissTime field value if set, zero value otherwise.
func (o *PlatformValues) GetMissTime() float32 {
	if o == nil || o.MissTime == nil {
		var ret float32
		return ret
	}
	return *o.MissTime
}

// GetMissTimeOk returns a tuple with the MissTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetMissTimeOk() (*float32, bool) {
	if o == nil || o.MissTime == nil {
		return nil, false
	}
	return o.MissTime, true
}

// HasMissTime returns a boolean if a field has been set.
func (o *PlatformValues) HasMissTime() bool {
	if o != nil && o.MissTime != nil {
		return true
	}

	return false
}

// SetMissTime gets a reference to the given float32 and assigns it to the MissTime field.
func (o *PlatformValues) SetMissTime(v float32) {
	o.MissTime = &v
}

// GetPass returns the Pass field value if set, zero value otherwise.
func (o *PlatformValues) GetPass() int64 {
	if o == nil || o.Pass == nil {
		var ret int64
		return ret
	}
	return *o.Pass
}

// GetPassOk returns a tuple with the Pass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetPassOk() (*int64, bool) {
	if o == nil || o.Pass == nil {
		return nil, false
	}
	return o.Pass, true
}

// HasPass returns a boolean if a field has been set.
func (o *PlatformValues) HasPass() bool {
	if o != nil && o.Pass != nil {
		return true
	}

	return false
}

// SetPass gets a reference to the given int64 and assigns it to the Pass field.
func (o *PlatformValues) SetPass(v int64) {
	o.Pass = &v
}

// GetPassTime returns the PassTime field value if set, zero value otherwise.
func (o *PlatformValues) GetPassTime() float32 {
	if o == nil || o.PassTime == nil {
		var ret float32
		return ret
	}
	return *o.PassTime
}

// GetPassTimeOk returns a tuple with the PassTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetPassTimeOk() (*float32, bool) {
	if o == nil || o.PassTime == nil {
		return nil, false
	}
	return o.PassTime, true
}

// HasPassTime returns a boolean if a field has been set.
func (o *PlatformValues) HasPassTime() bool {
	if o != nil && o.PassTime != nil {
		return true
	}

	return false
}

// SetPassTime gets a reference to the given float32 and assigns it to the PassTime field.
func (o *PlatformValues) SetPassTime(v float32) {
	o.PassTime = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *PlatformValues) GetErrors() int64 {
	if o == nil || o.Errors == nil {
		var ret int64
		return ret
	}
	return *o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetErrorsOk() (*int64, bool) {
	if o == nil || o.Errors == nil {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *PlatformValues) HasErrors() bool {
	if o != nil && o.Errors != nil {
		return true
	}

	return false
}

// SetErrors gets a reference to the given int64 and assigns it to the Errors field.
func (o *PlatformValues) SetErrors(v int64) {
	o.Errors = &v
}

// GetRestarts returns the Restarts field value if set, zero value otherwise.
func (o *PlatformValues) GetRestarts() int64 {
	if o == nil || o.Restarts == nil {
		var ret int64
		return ret
	}
	return *o.Restarts
}

// GetRestartsOk returns a tuple with the Restarts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetRestartsOk() (*int64, bool) {
	if o == nil || o.Restarts == nil {
		return nil, false
	}
	return o.Restarts, true
}

// HasRestarts returns a boolean if a field has been set.
func (o *PlatformValues) HasRestarts() bool {
	if o != nil && o.Restarts != nil {
		return true
	}

	return false
}

// SetRestarts gets a reference to the given int64 and assigns it to the Restarts field.
func (o *PlatformValues) SetRestarts(v int64) {
	o.Restarts = &v
}

// GetHitRatio returns the HitRatio field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlatformValues) GetHitRatio() float32 {
	if o == nil || o.HitRatio.Get() == nil {
		var ret float32
		return ret
	}
	return *o.HitRatio.Get()
}

// GetHitRatioOk returns a tuple with the HitRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlatformValues) GetHitRatioOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.HitRatio.Get(), o.HitRatio.IsSet()
}

// HasHitRatio returns a boolean if a field has been set.
func (o *PlatformValues) HasHitRatio() bool {
	if o != nil && o.HitRatio.IsSet() {
		return true
	}

	return false
}

// SetHitRatio gets a reference to the given NullableFloat32 and assigns it to the HitRatio field.
func (o *PlatformValues) SetHitRatio(v float32) {
	o.HitRatio.Set(&v)
}

// SetHitRatioNil sets the value for HitRatio to be an explicit nil
func (o *PlatformValues) SetHitRatioNil() {
	o.HitRatio.Set(nil)
}

// UnsetHitRatio ensures that no value is present for HitRatio, not even an explicit nil
func (o *PlatformValues) UnsetHitRatio() {
	o.HitRatio.Unset()
}

// GetBandwidth returns the Bandwidth field value if set, zero value otherwise.
func (o *PlatformValues) GetBandwidth() int64 {
	if o == nil || o.Bandwidth == nil {
		var ret int64
		return ret
	}
	return *o.Bandwidth
}

// GetBandwidthOk returns a tuple with the Bandwidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetBandwidthOk() (*int64, bool) {
	if o == nil || o.Bandwidth == nil {
		return nil, false
	}
	return o.Bandwidth, true
}

// HasBandwidth returns a boolean if a field has been set.
func (o *PlatformValues) HasBandwidth() bool {
	if o != nil && o.Bandwidth != nil {
		return true
	}

	return false
}

// SetBandwidth gets a reference to the given int64 and assigns it to the Bandwidth field.
func (o *PlatformValues) SetBandwidth(v int64) {
	o.Bandwidth = &v
}

// GetReqBodyBytes returns the ReqBodyBytes field value if set, zero value otherwise.
func (o *PlatformValues) GetReqBodyBytes() int64 {
	if o == nil || o.ReqBodyBytes == nil {
		var ret int64
		return ret
	}
	return *o.ReqBodyBytes
}

// GetReqBodyBytesOk returns a tuple with the ReqBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetReqBodyBytesOk() (*int64, bool) {
	if o == nil || o.ReqBodyBytes == nil {
		return nil, false
	}
	return o.ReqBodyBytes, true
}

// HasReqBodyBytes returns a boolean if a field has been set.
func (o *PlatformValues) HasReqBodyBytes() bool {
	if o != nil && o.ReqBodyBytes != nil {
		return true
	}

	return false
}

// SetReqBodyBytes gets a reference to the given int64 and assigns it to the ReqBodyBytes field.
func (o *PlatformValues) SetReqBodyBytes(v int64) {
	o.ReqBodyBytes = &v
}

// GetReqHeaderBytes returns the ReqHeaderBytes field value if set, zero value otherwise.
func (o *PlatformValues) GetReqHeaderBytes() int64 {
	if o == nil || o.ReqHeaderBytes == nil {
		var ret int64
		return ret
	}
	return *o.ReqHeaderBytes
}

// GetReqHeaderBytesOk returns a tuple with the ReqHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetReqHeaderBytesOk() (*int64, bool) {
	if o == nil || o.ReqHeaderBytes == nil {
		return nil, false
	}
	return o.ReqHeaderBytes, true
}

// HasReqHeaderBytes returns a boolean if a field has been set.
func (o *PlatformValues) HasReqHeaderBytes() bool {
	if o != nil && o.ReqHeaderBytes != nil {
		return true
	}

	return false
}

// SetReqHeaderBytes gets a reference to the given int64 and assigns it to the ReqHeaderBytes field.
func (o *PlatformValues) SetReqHeaderBytes(v int64) {
	o.ReqHeaderBytes = &v
}

// GetRespBodyBytes returns the RespBodyBytes field value if set, zero value otherwise.
func (o *PlatformValues) GetRespBodyBytes() int64 {
	if o == nil || o.RespBodyBytes == nil {
		var ret int64
		return ret
	}
	return *o.RespBodyBytes
}

// GetRespBodyBytesOk returns a tuple with the RespBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetRespBodyBytesOk() (*int64, bool) {
	if o == nil || o.RespBodyBytes == nil {
		return nil, false
	}
	return o.RespBodyBytes, true
}

// HasRespBodyBytes returns a boolean if a field has been set.
func (o *PlatformValues) HasRespBodyBytes() bool {
	if o != nil && o.RespBodyBytes != nil {
		return true
	}

	return false
}

// SetRespBodyBytes gets a reference to the given int64 and assigns it to the RespBodyBytes field.
func (o *PlatformValues) SetRespBodyBytes(v int64) {
	o.RespBodyBytes = &v
}

// GetRespHeaderBytes returns the RespHeaderBytes field value if set, zero value otherwise.
func (o *PlatformValues) GetRespHeaderBytes() int64 {
	if o == nil || o.RespHeaderBytes == nil {
		var ret int64
		return ret
	}
	return *o.RespHeaderBytes
}

// GetRespHeaderBytesOk returns a tuple with the RespHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetRespHeaderBytesOk() (*int64, bool) {
	if o == nil || o.RespHeaderBytes == nil {
		return nil, false
	}
	return o.RespHeaderBytes, true
}

// HasRespHeaderBytes returns a boolean if a field has been set.
func (o *PlatformValues) HasRespHeaderBytes() bool {
	if o != nil && o.RespHeaderBytes != nil {
		return true
	}

	return false
}

// SetRespHeaderBytes gets a reference to the given int64 and assigns it to the RespHeaderBytes field.
func (o *PlatformValues) SetRespHeaderBytes(v int64) {
	o.RespHeaderBytes = &v
}

// GetBereqBodyBytes returns the BereqBodyBytes field value if set, zero value otherwise.
func (o *PlatformValues) GetBereqBodyBytes() int64 {
	if o == nil || o.BereqBodyBytes == nil {
		var ret int64
		return ret
	}
	return *o.BereqBodyBytes
}

// GetBereqBodyBytesOk returns a tuple with the BereqBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetBereqBodyBytesOk() (*int64, bool) {
	if o == nil || o.BereqBodyBytes == nil {
		return nil, false
	}
	return o.BereqBodyBytes, true
}

// HasBereqBodyBytes returns a boolean if a field has been set.
func (o *PlatformValues) HasBereqBodyBytes() bool {
	if o != nil && o.BereqBodyBytes != nil {
		return true
	}

	return false
}

// SetBereqBodyBytes gets a reference to the given int64 and assigns it to the BereqBodyBytes field.
func (o *PlatformValues) SetBereqBodyBytes(v int64) {
	o.BereqBodyBytes = &v
}

// GetBereqHeaderBytes returns the BereqHeaderBytes field value if set, zero value otherwise.
func (o *PlatformValues) GetBereqHeaderBytes() int64 {
	if o == nil || o.BereqHeaderBytes == nil {
		var ret int64
		return ret
	}
	return *o.BereqHeaderBytes
}

// GetBereqHeaderBytesOk returns a tuple with the BereqHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetBereqHeaderBytesOk() (*int64, bool) {
	if o == nil || o.BereqHeaderBytes == nil {
		return nil, false
	}
	return o.BereqHeaderBytes, true
}

// HasBereqHeaderBytes returns a boolean if a field has been set.
func (o *PlatformValues) HasBereqHeaderBytes() bool {
	if o != nil && o.BereqHeaderBytes != nil {
		return true
	}

	return false
}

// SetBereqHeaderBytes gets a reference to the given int64 and assigns it to the BereqHeaderBytes field.
func (o *PlatformValues) SetBereqHeaderBytes(v int64) {
	o.BereqHeaderBytes = &v
}

// GetUncacheable returns the Uncacheable field value if set, zero value otherwise.
func (o *PlatformValues) GetUncacheable() int64 {
	if o == nil || o.Uncacheable == nil {
		var ret int64
		return ret
	}
	return *o.Uncacheable
}

// GetUncacheableOk returns a tuple with the Uncacheable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetUncacheableOk() (*int64, bool) {
	if o == nil || o.Uncacheable == nil {
		return nil, false
	}
	return o.Uncacheable, true
}

// HasUncacheable returns a boolean if a field has been set.
func (o *PlatformValues) HasUncacheable() bool {
	if o != nil && o.Uncacheable != nil {
		return true
	}

	return false
}

// SetUncacheable gets a reference to the given int64 and assigns it to the Uncacheable field.
func (o *PlatformValues) SetUncacheable(v int64) {
	o.Uncacheable = &v
}

// GetPipe returns the Pipe field value if set, zero value otherwise.
func (o *PlatformValues) GetPipe() int64 {
	if o == nil || o.Pipe == nil {
		var ret int64
		return ret
	}
	return *o.Pipe
}

// GetPipeOk returns a tuple with the Pipe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetPipeOk() (*int64, bool) {
	if o == nil || o.Pipe == nil {
		return nil, false
	}
	return o.Pipe, true
}

// HasPipe returns a boolean if a field has been set.
func (o *PlatformValues) HasPipe() bool {
	if o != nil && o.Pipe != nil {
		return true
	}

	return false
}

// SetPipe gets a reference to the given int64 and assigns it to the Pipe field.
func (o *PlatformValues) SetPipe(v int64) {
	o.Pipe = &v
}

// GetSynth returns the Synth field value if set, zero value otherwise.
func (o *PlatformValues) GetSynth() int64 {
	if o == nil || o.Synth == nil {
		var ret int64
		return ret
	}
	return *o.Synth
}

// GetSynthOk returns a tuple with the Synth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetSynthOk() (*int64, bool) {
	if o == nil || o.Synth == nil {
		return nil, false
	}
	return o.Synth, true
}

// HasSynth returns a boolean if a field has been set.
func (o *PlatformValues) HasSynth() bool {
	if o != nil && o.Synth != nil {
		return true
	}

	return false
}

// SetSynth gets a reference to the given int64 and assigns it to the Synth field.
func (o *PlatformValues) SetSynth(v int64) {
	o.Synth = &v
}

// GetTLS returns the TLS field value if set, zero value otherwise.
func (o *PlatformValues) GetTLS() int64 {
	if o == nil || o.TLS == nil {
		var ret int64
		return ret
	}
	return *o.TLS
}

// GetTLSOk returns a tuple with the TLS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetTLSOk() (*int64, bool) {
	if o == nil || o.TLS == nil {
		return nil, false
	}
	return o.TLS, true
}

// HasTLS returns a boolean if a field has been set.
func (o *PlatformValues) HasTLS() bool {
	if o != nil && o.TLS != nil {
		return true
	}

	return false
}

// SetTLS gets a reference to the given int64 and assigns it to the TLS field.
func (o *PlatformValues) SetTLS(v int64) {
	o.TLS = &v
}

// GetTLSV10 returns the TLSV10 field value if set, zero value otherwise.
func (o *PlatformValues) GetTLSV10() int64 {
	if o == nil || o.TLSV10 == nil {
		var ret int64
		return ret
	}
	return *o.TLSV10
}

// GetTLSV10Ok returns a tuple with the TLSV10 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetTLSV10Ok() (*int64, bool) {
	if o == nil || o.TLSV10 == nil {
		return nil, false
	}
	return o.TLSV10, true
}

// HasTLSV10 returns a boolean if a field has been set.
func (o *PlatformValues) HasTLSV10() bool {
	if o != nil && o.TLSV10 != nil {
		return true
	}

	return false
}

// SetTLSV10 gets a reference to the given int64 and assigns it to the TLSV10 field.
func (o *PlatformValues) SetTLSV10(v int64) {
	o.TLSV10 = &v
}

// GetTLSV11 returns the TLSV11 field value if set, zero value otherwise.
func (o *PlatformValues) GetTLSV11() int64 {
	if o == nil || o.TLSV11 == nil {
		var ret int64
		return ret
	}
	return *o.TLSV11
}

// GetTLSV11Ok returns a tuple with the TLSV11 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetTLSV11Ok() (*int64, bool) {
	if o == nil || o.TLSV11 == nil {
		return nil, false
	}
	return o.TLSV11, true
}

// HasTLSV11 returns a boolean if a field has been set.
func (o *PlatformValues) HasTLSV11() bool {
	if o != nil && o.TLSV11 != nil {
		return true
	}

	return false
}

// SetTLSV11 gets a reference to the given int64 and assigns it to the TLSV11 field.
func (o *PlatformValues) SetTLSV11(v int64) {
	o.TLSV11 = &v
}

// GetTLSV12 returns the TLSV12 field value if set, zero value otherwise.
func (o *PlatformValues) GetTLSV12() int64 {
	if o == nil || o.TLSV12 == nil {
		var ret int64
		return ret
	}
	return *o.TLSV12
}

// GetTLSV12Ok returns a tuple with the TLSV12 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetTLSV12Ok() (*int64, bool) {
	if o == nil || o.TLSV12 == nil {
		return nil, false
	}
	return o.TLSV12, true
}

// HasTLSV12 returns a boolean if a field has been set.
func (o *PlatformValues) HasTLSV12() bool {
	if o != nil && o.TLSV12 != nil {
		return true
	}

	return false
}

// SetTLSV12 gets a reference to the given int64 and assigns it to the TLSV12 field.
func (o *PlatformValues) SetTLSV12(v int64) {
	o.TLSV12 = &v
}

// GetTLSV13 returns the TLSV13 field value if set, zero value otherwise.
func (o *PlatformValues) GetTLSV13() int64 {
	if o == nil || o.TLSV13 == nil {
		var ret int64
		return ret
	}
	return *o.TLSV13
}

// GetTLSV13Ok returns a tuple with the TLSV13 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetTLSV13Ok() (*int64, bool) {
	if o == nil || o.TLSV13 == nil {
		return nil, false
	}
	return o.TLSV13, true
}

// HasTLSV13 returns a boolean if a field has been set.
func (o *PlatformValues) HasTLSV13() bool {
	if o != nil && o.TLSV13 != nil {
		return true
	}

	return false
}

// SetTLSV13 gets a reference to the given int64 and assigns it to the TLSV13 field.
func (o *PlatformValues) SetTLSV13(v int64) {
	o.TLSV13 = &v
}

// GetEdgeRequests returns the EdgeRequests field value if set, zero value otherwise.
func (o *PlatformValues) GetEdgeRequests() int64 {
	if o == nil || o.EdgeRequests == nil {
		var ret int64
		return ret
	}
	return *o.EdgeRequests
}

// GetEdgeRequestsOk returns a tuple with the EdgeRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetEdgeRequestsOk() (*int64, bool) {
	if o == nil || o.EdgeRequests == nil {
		return nil, false
	}
	return o.EdgeRequests, true
}

// HasEdgeRequests returns a boolean if a field has been set.
func (o *PlatformValues) HasEdgeRequests() bool {
	if o != nil && o.EdgeRequests != nil {
		return true
	}

	return false
}

// SetEdgeRequests gets a reference to the given int64 and assigns it to the EdgeRequests field.
func (o *PlatformValues) SetEdgeRequests(v int64) {
	o.EdgeRequests = &v
}

// GetEdgeRespHeaderBytes returns the EdgeRespHeaderBytes field value if set, zero value otherwise.
func (o *PlatformValues) GetEdgeRespHeaderBytes() int64 {
	if o == nil || o.EdgeRespHeaderBytes == nil {
		var ret int64
		return ret
	}
	return *o.EdgeRespHeaderBytes
}

// GetEdgeRespHeaderBytesOk returns a tuple with the EdgeRespHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetEdgeRespHeaderBytesOk() (*int64, bool) {
	if o == nil || o.EdgeRespHeaderBytes == nil {
		return nil, false
	}
	return o.EdgeRespHeaderBytes, true
}

// HasEdgeRespHeaderBytes returns a boolean if a field has been set.
func (o *PlatformValues) HasEdgeRespHeaderBytes() bool {
	if o != nil && o.EdgeRespHeaderBytes != nil {
		return true
	}

	return false
}

// SetEdgeRespHeaderBytes gets a reference to the given int64 and assigns it to the EdgeRespHeaderBytes field.
func (o *PlatformValues) SetEdgeRespHeaderBytes(v int64) {
	o.EdgeRespHeaderBytes = &v
}

// GetEdgeRespBodyBytes returns the EdgeRespBodyBytes field value if set, zero value otherwise.
func (o *PlatformValues) GetEdgeRespBodyBytes() int64 {
	if o == nil || o.EdgeRespBodyBytes == nil {
		var ret int64
		return ret
	}
	return *o.EdgeRespBodyBytes
}

// GetEdgeRespBodyBytesOk returns a tuple with the EdgeRespBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetEdgeRespBodyBytesOk() (*int64, bool) {
	if o == nil || o.EdgeRespBodyBytes == nil {
		return nil, false
	}
	return o.EdgeRespBodyBytes, true
}

// HasEdgeRespBodyBytes returns a boolean if a field has been set.
func (o *PlatformValues) HasEdgeRespBodyBytes() bool {
	if o != nil && o.EdgeRespBodyBytes != nil {
		return true
	}

	return false
}

// SetEdgeRespBodyBytes gets a reference to the given int64 and assigns it to the EdgeRespBodyBytes field.
func (o *PlatformValues) SetEdgeRespBodyBytes(v int64) {
	o.EdgeRespBodyBytes = &v
}

// GetEdgeHitRequests returns the EdgeHitRequests field value if set, zero value otherwise.
func (o *PlatformValues) GetEdgeHitRequests() int64 {
	if o == nil || o.EdgeHitRequests == nil {
		var ret int64
		return ret
	}
	return *o.EdgeHitRequests
}

// GetEdgeHitRequestsOk returns a tuple with the EdgeHitRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetEdgeHitRequestsOk() (*int64, bool) {
	if o == nil || o.EdgeHitRequests == nil {
		return nil, false
	}
	return o.EdgeHitRequests, true
}

// HasEdgeHitRequests returns a boolean if a field has been set.
func (o *PlatformValues) HasEdgeHitRequests() bool {
	if o != nil && o.EdgeHitRequests != nil {
		return true
	}

	return false
}

// SetEdgeHitRequests gets a reference to the given int64 and assigns it to the EdgeHitRequests field.
func (o *PlatformValues) SetEdgeHitRequests(v int64) {
	o.EdgeHitRequests = &v
}

// GetEdgeMissRequests returns the EdgeMissRequests field value if set, zero value otherwise.
func (o *PlatformValues) GetEdgeMissRequests() int64 {
	if o == nil || o.EdgeMissRequests == nil {
		var ret int64
		return ret
	}
	return *o.EdgeMissRequests
}

// GetEdgeMissRequestsOk returns a tuple with the EdgeMissRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetEdgeMissRequestsOk() (*int64, bool) {
	if o == nil || o.EdgeMissRequests == nil {
		return nil, false
	}
	return o.EdgeMissRequests, true
}

// HasEdgeMissRequests returns a boolean if a field has been set.
func (o *PlatformValues) HasEdgeMissRequests() bool {
	if o != nil && o.EdgeMissRequests != nil {
		return true
	}

	return false
}

// SetEdgeMissRequests gets a reference to the given int64 and assigns it to the EdgeMissRequests field.
func (o *PlatformValues) SetEdgeMissRequests(v int64) {
	o.EdgeMissRequests = &v
}

// GetOriginFetches returns the OriginFetches field value if set, zero value otherwise.
func (o *PlatformValues) GetOriginFetches() int64 {
	if o == nil || o.OriginFetches == nil {
		var ret int64
		return ret
	}
	return *o.OriginFetches
}

// GetOriginFetchesOk returns a tuple with the OriginFetches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetOriginFetchesOk() (*int64, bool) {
	if o == nil || o.OriginFetches == nil {
		return nil, false
	}
	return o.OriginFetches, true
}

// HasOriginFetches returns a boolean if a field has been set.
func (o *PlatformValues) HasOriginFetches() bool {
	if o != nil && o.OriginFetches != nil {
		return true
	}

	return false
}

// SetOriginFetches gets a reference to the given int64 and assigns it to the OriginFetches field.
func (o *PlatformValues) SetOriginFetches(v int64) {
	o.OriginFetches = &v
}

// GetOriginFetchHeaderBytes returns the OriginFetchHeaderBytes field value if set, zero value otherwise.
func (o *PlatformValues) GetOriginFetchHeaderBytes() int64 {
	if o == nil || o.OriginFetchHeaderBytes == nil {
		var ret int64
		return ret
	}
	return *o.OriginFetchHeaderBytes
}

// GetOriginFetchHeaderBytesOk returns a tuple with the OriginFetchHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetOriginFetchHeaderBytesOk() (*int64, bool) {
	if o == nil || o.OriginFetchHeaderBytes == nil {
		return nil, false
	}
	return o.OriginFetchHeaderBytes, true
}

// HasOriginFetchHeaderBytes returns a boolean if a field has been set.
func (o *PlatformValues) HasOriginFetchHeaderBytes() bool {
	if o != nil && o.OriginFetchHeaderBytes != nil {
		return true
	}

	return false
}

// SetOriginFetchHeaderBytes gets a reference to the given int64 and assigns it to the OriginFetchHeaderBytes field.
func (o *PlatformValues) SetOriginFetchHeaderBytes(v int64) {
	o.OriginFetchHeaderBytes = &v
}

// GetOriginFetchBodyBytes returns the OriginFetchBodyBytes field value if set, zero value otherwise.
func (o *PlatformValues) GetOriginFetchBodyBytes() int64 {
	if o == nil || o.OriginFetchBodyBytes == nil {
		var ret int64
		return ret
	}
	return *o.OriginFetchBodyBytes
}

// GetOriginFetchBodyBytesOk returns a tuple with the OriginFetchBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetOriginFetchBodyBytesOk() (*int64, bool) {
	if o == nil || o.OriginFetchBodyBytes == nil {
		return nil, false
	}
	return o.OriginFetchBodyBytes, true
}

// HasOriginFetchBodyBytes returns a boolean if a field has been set.
func (o *PlatformValues) HasOriginFetchBodyBytes() bool {
	if o != nil && o.OriginFetchBodyBytes != nil {
		return true
	}

	return false
}

// SetOriginFetchBodyBytes gets a reference to the given int64 and assigns it to the OriginFetchBodyBytes field.
func (o *PlatformValues) SetOriginFetchBodyBytes(v int64) {
	o.OriginFetchBodyBytes = &v
}

// GetOriginFetchRespHeaderBytes returns the OriginFetchRespHeaderBytes field value if set, zero value otherwise.
func (o *PlatformValues) GetOriginFetchRespHeaderBytes() int64 {
	if o == nil || o.OriginFetchRespHeaderBytes == nil {
		var ret int64
		return ret
	}
	return *o.OriginFetchRespHeaderBytes
}

// GetOriginFetchRespHeaderBytesOk returns a tuple with the OriginFetchRespHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetOriginFetchRespHeaderBytesOk() (*int64, bool) {
	if o == nil || o.OriginFetchRespHeaderBytes == nil {
		return nil, false
	}
	return o.OriginFetchRespHeaderBytes, true
}

// HasOriginFetchRespHeaderBytes returns a boolean if a field has been set.
func (o *PlatformValues) HasOriginFetchRespHeaderBytes() bool {
	if o != nil && o.OriginFetchRespHeaderBytes != nil {
		return true
	}

	return false
}

// SetOriginFetchRespHeaderBytes gets a reference to the given int64 and assigns it to the OriginFetchRespHeaderBytes field.
func (o *PlatformValues) SetOriginFetchRespHeaderBytes(v int64) {
	o.OriginFetchRespHeaderBytes = &v
}

// GetOriginFetchRespBodyBytes returns the OriginFetchRespBodyBytes field value if set, zero value otherwise.
func (o *PlatformValues) GetOriginFetchRespBodyBytes() int64 {
	if o == nil || o.OriginFetchRespBodyBytes == nil {
		var ret int64
		return ret
	}
	return *o.OriginFetchRespBodyBytes
}

// GetOriginFetchRespBodyBytesOk returns a tuple with the OriginFetchRespBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetOriginFetchRespBodyBytesOk() (*int64, bool) {
	if o == nil || o.OriginFetchRespBodyBytes == nil {
		return nil, false
	}
	return o.OriginFetchRespBodyBytes, true
}

// HasOriginFetchRespBodyBytes returns a boolean if a field has been set.
func (o *PlatformValues) HasOriginFetchRespBodyBytes() bool {
	if o != nil && o.OriginFetchRespBodyBytes != nil {
		return true
	}

	return false
}

// SetOriginFetchRespBodyBytes gets a reference to the given int64 and assigns it to the OriginFetchRespBodyBytes field.
func (o *PlatformValues) SetOriginFetchRespBodyBytes(v int64) {
	o.OriginFetchRespBodyBytes = &v
}

// GetOriginRevalidations returns the OriginRevalidations field value if set, zero value otherwise.
func (o *PlatformValues) GetOriginRevalidations() int64 {
	if o == nil || o.OriginRevalidations == nil {
		var ret int64
		return ret
	}
	return *o.OriginRevalidations
}

// GetOriginRevalidationsOk returns a tuple with the OriginRevalidations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetOriginRevalidationsOk() (*int64, bool) {
	if o == nil || o.OriginRevalidations == nil {
		return nil, false
	}
	return o.OriginRevalidations, true
}

// HasOriginRevalidations returns a boolean if a field has been set.
func (o *PlatformValues) HasOriginRevalidations() bool {
	if o != nil && o.OriginRevalidations != nil {
		return true
	}

	return false
}

// SetOriginRevalidations gets a reference to the given int64 and assigns it to the OriginRevalidations field.
func (o *PlatformValues) SetOriginRevalidations(v int64) {
	o.OriginRevalidations = &v
}

// GetOriginCacheFetches returns the OriginCacheFetches field value if set, zero value otherwise.
func (o *PlatformValues) GetOriginCacheFetches() int64 {
	if o == nil || o.OriginCacheFetches == nil {
		var ret int64
		return ret
	}
	return *o.OriginCacheFetches
}

// GetOriginCacheFetchesOk returns a tuple with the OriginCacheFetches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetOriginCacheFetchesOk() (*int64, bool) {
	if o == nil || o.OriginCacheFetches == nil {
		return nil, false
	}
	return o.OriginCacheFetches, true
}

// HasOriginCacheFetches returns a boolean if a field has been set.
func (o *PlatformValues) HasOriginCacheFetches() bool {
	if o != nil && o.OriginCacheFetches != nil {
		return true
	}

	return false
}

// SetOriginCacheFetches gets a reference to the given int64 and assigns it to the OriginCacheFetches field.
func (o *PlatformValues) SetOriginCacheFetches(v int64) {
	o.OriginCacheFetches = &v
}

// GetShield returns the Shield field value if set, zero value otherwise.
func (o *PlatformValues) GetShield() int64 {
	if o == nil || o.Shield == nil {
		var ret int64
		return ret
	}
	return *o.Shield
}

// GetShieldOk returns a tuple with the Shield field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetShieldOk() (*int64, bool) {
	if o == nil || o.Shield == nil {
		return nil, false
	}
	return o.Shield, true
}

// HasShield returns a boolean if a field has been set.
func (o *PlatformValues) HasShield() bool {
	if o != nil && o.Shield != nil {
		return true
	}

	return false
}

// SetShield gets a reference to the given int64 and assigns it to the Shield field.
func (o *PlatformValues) SetShield(v int64) {
	o.Shield = &v
}

// GetShieldRespBodyBytes returns the ShieldRespBodyBytes field value if set, zero value otherwise.
func (o *PlatformValues) GetShieldRespBodyBytes() int64 {
	if o == nil || o.ShieldRespBodyBytes == nil {
		var ret int64
		return ret
	}
	return *o.ShieldRespBodyBytes
}

// GetShieldRespBodyBytesOk returns a tuple with the ShieldRespBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetShieldRespBodyBytesOk() (*int64, bool) {
	if o == nil || o.ShieldRespBodyBytes == nil {
		return nil, false
	}
	return o.ShieldRespBodyBytes, true
}

// HasShieldRespBodyBytes returns a boolean if a field has been set.
func (o *PlatformValues) HasShieldRespBodyBytes() bool {
	if o != nil && o.ShieldRespBodyBytes != nil {
		return true
	}

	return false
}

// SetShieldRespBodyBytes gets a reference to the given int64 and assigns it to the ShieldRespBodyBytes field.
func (o *PlatformValues) SetShieldRespBodyBytes(v int64) {
	o.ShieldRespBodyBytes = &v
}

// GetShieldRespHeaderBytes returns the ShieldRespHeaderBytes field value if set, zero value otherwise.
func (o *PlatformValues) GetShieldRespHeaderBytes() int64 {
	if o == nil || o.ShieldRespHeaderBytes == nil {
		var ret int64
		return ret
	}
	return *o.ShieldRespHeaderBytes
}

// GetShieldRespHeaderBytesOk returns a tuple with the ShieldRespHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetShieldRespHeaderBytesOk() (*int64, bool) {
	if o == nil || o.ShieldRespHeaderBytes == nil {
		return nil, false
	}
	return o.ShieldRespHeaderBytes, true
}

// HasShieldRespHeaderBytes returns a boolean if a field has been set.
func (o *PlatformValues) HasShieldRespHeaderBytes() bool {
	if o != nil && o.ShieldRespHeaderBytes != nil {
		return true
	}

	return false
}

// SetShieldRespHeaderBytes gets a reference to the given int64 and assigns it to the ShieldRespHeaderBytes field.
func (o *PlatformValues) SetShieldRespHeaderBytes(v int64) {
	o.ShieldRespHeaderBytes = &v
}

// GetShieldFetches returns the ShieldFetches field value if set, zero value otherwise.
func (o *PlatformValues) GetShieldFetches() int64 {
	if o == nil || o.ShieldFetches == nil {
		var ret int64
		return ret
	}
	return *o.ShieldFetches
}

// GetShieldFetchesOk returns a tuple with the ShieldFetches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetShieldFetchesOk() (*int64, bool) {
	if o == nil || o.ShieldFetches == nil {
		return nil, false
	}
	return o.ShieldFetches, true
}

// HasShieldFetches returns a boolean if a field has been set.
func (o *PlatformValues) HasShieldFetches() bool {
	if o != nil && o.ShieldFetches != nil {
		return true
	}

	return false
}

// SetShieldFetches gets a reference to the given int64 and assigns it to the ShieldFetches field.
func (o *PlatformValues) SetShieldFetches(v int64) {
	o.ShieldFetches = &v
}

// GetShieldFetchHeaderBytes returns the ShieldFetchHeaderBytes field value if set, zero value otherwise.
func (o *PlatformValues) GetShieldFetchHeaderBytes() int64 {
	if o == nil || o.ShieldFetchHeaderBytes == nil {
		var ret int64
		return ret
	}
	return *o.ShieldFetchHeaderBytes
}

// GetShieldFetchHeaderBytesOk returns a tuple with the ShieldFetchHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetShieldFetchHeaderBytesOk() (*int64, bool) {
	if o == nil || o.ShieldFetchHeaderBytes == nil {
		return nil, false
	}
	return o.ShieldFetchHeaderBytes, true
}

// HasShieldFetchHeaderBytes returns a boolean if a field has been set.
func (o *PlatformValues) HasShieldFetchHeaderBytes() bool {
	if o != nil && o.ShieldFetchHeaderBytes != nil {
		return true
	}

	return false
}

// SetShieldFetchHeaderBytes gets a reference to the given int64 and assigns it to the ShieldFetchHeaderBytes field.
func (o *PlatformValues) SetShieldFetchHeaderBytes(v int64) {
	o.ShieldFetchHeaderBytes = &v
}

// GetShieldFetchBodyBytes returns the ShieldFetchBodyBytes field value if set, zero value otherwise.
func (o *PlatformValues) GetShieldFetchBodyBytes() int64 {
	if o == nil || o.ShieldFetchBodyBytes == nil {
		var ret int64
		return ret
	}
	return *o.ShieldFetchBodyBytes
}

// GetShieldFetchBodyBytesOk returns a tuple with the ShieldFetchBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetShieldFetchBodyBytesOk() (*int64, bool) {
	if o == nil || o.ShieldFetchBodyBytes == nil {
		return nil, false
	}
	return o.ShieldFetchBodyBytes, true
}

// HasShieldFetchBodyBytes returns a boolean if a field has been set.
func (o *PlatformValues) HasShieldFetchBodyBytes() bool {
	if o != nil && o.ShieldFetchBodyBytes != nil {
		return true
	}

	return false
}

// SetShieldFetchBodyBytes gets a reference to the given int64 and assigns it to the ShieldFetchBodyBytes field.
func (o *PlatformValues) SetShieldFetchBodyBytes(v int64) {
	o.ShieldFetchBodyBytes = &v
}

// GetShieldFetchRespHeaderBytes returns the ShieldFetchRespHeaderBytes field value if set, zero value otherwise.
func (o *PlatformValues) GetShieldFetchRespHeaderBytes() int64 {
	if o == nil || o.ShieldFetchRespHeaderBytes == nil {
		var ret int64
		return ret
	}
	return *o.ShieldFetchRespHeaderBytes
}

// GetShieldFetchRespHeaderBytesOk returns a tuple with the ShieldFetchRespHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetShieldFetchRespHeaderBytesOk() (*int64, bool) {
	if o == nil || o.ShieldFetchRespHeaderBytes == nil {
		return nil, false
	}
	return o.ShieldFetchRespHeaderBytes, true
}

// HasShieldFetchRespHeaderBytes returns a boolean if a field has been set.
func (o *PlatformValues) HasShieldFetchRespHeaderBytes() bool {
	if o != nil && o.ShieldFetchRespHeaderBytes != nil {
		return true
	}

	return false
}

// SetShieldFetchRespHeaderBytes gets a reference to the given int64 and assigns it to the ShieldFetchRespHeaderBytes field.
func (o *PlatformValues) SetShieldFetchRespHeaderBytes(v int64) {
	o.ShieldFetchRespHeaderBytes = &v
}

// GetShieldFetchRespBodyBytes returns the ShieldFetchRespBodyBytes field value if set, zero value otherwise.
func (o *PlatformValues) GetShieldFetchRespBodyBytes() int64 {
	if o == nil || o.ShieldFetchRespBodyBytes == nil {
		var ret int64
		return ret
	}
	return *o.ShieldFetchRespBodyBytes
}

// GetShieldFetchRespBodyBytesOk returns a tuple with the ShieldFetchRespBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetShieldFetchRespBodyBytesOk() (*int64, bool) {
	if o == nil || o.ShieldFetchRespBodyBytes == nil {
		return nil, false
	}
	return o.ShieldFetchRespBodyBytes, true
}

// HasShieldFetchRespBodyBytes returns a boolean if a field has been set.
func (o *PlatformValues) HasShieldFetchRespBodyBytes() bool {
	if o != nil && o.ShieldFetchRespBodyBytes != nil {
		return true
	}

	return false
}

// SetShieldFetchRespBodyBytes gets a reference to the given int64 and assigns it to the ShieldFetchRespBodyBytes field.
func (o *PlatformValues) SetShieldFetchRespBodyBytes(v int64) {
	o.ShieldFetchRespBodyBytes = &v
}

// GetShieldRevalidations returns the ShieldRevalidations field value if set, zero value otherwise.
func (o *PlatformValues) GetShieldRevalidations() int64 {
	if o == nil || o.ShieldRevalidations == nil {
		var ret int64
		return ret
	}
	return *o.ShieldRevalidations
}

// GetShieldRevalidationsOk returns a tuple with the ShieldRevalidations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetShieldRevalidationsOk() (*int64, bool) {
	if o == nil || o.ShieldRevalidations == nil {
		return nil, false
	}
	return o.ShieldRevalidations, true
}

// HasShieldRevalidations returns a boolean if a field has been set.
func (o *PlatformValues) HasShieldRevalidations() bool {
	if o != nil && o.ShieldRevalidations != nil {
		return true
	}

	return false
}

// SetShieldRevalidations gets a reference to the given int64 and assigns it to the ShieldRevalidations field.
func (o *PlatformValues) SetShieldRevalidations(v int64) {
	o.ShieldRevalidations = &v
}

// GetShieldCacheFetches returns the ShieldCacheFetches field value if set, zero value otherwise.
func (o *PlatformValues) GetShieldCacheFetches() int64 {
	if o == nil || o.ShieldCacheFetches == nil {
		var ret int64
		return ret
	}
	return *o.ShieldCacheFetches
}

// GetShieldCacheFetchesOk returns a tuple with the ShieldCacheFetches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetShieldCacheFetchesOk() (*int64, bool) {
	if o == nil || o.ShieldCacheFetches == nil {
		return nil, false
	}
	return o.ShieldCacheFetches, true
}

// HasShieldCacheFetches returns a boolean if a field has been set.
func (o *PlatformValues) HasShieldCacheFetches() bool {
	if o != nil && o.ShieldCacheFetches != nil {
		return true
	}

	return false
}

// SetShieldCacheFetches gets a reference to the given int64 and assigns it to the ShieldCacheFetches field.
func (o *PlatformValues) SetShieldCacheFetches(v int64) {
	o.ShieldCacheFetches = &v
}

// GetIpv6 returns the Ipv6 field value if set, zero value otherwise.
func (o *PlatformValues) GetIpv6() int64 {
	if o == nil || o.Ipv6 == nil {
		var ret int64
		return ret
	}
	return *o.Ipv6
}

// GetIpv6Ok returns a tuple with the Ipv6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetIpv6Ok() (*int64, bool) {
	if o == nil || o.Ipv6 == nil {
		return nil, false
	}
	return o.Ipv6, true
}

// HasIpv6 returns a boolean if a field has been set.
func (o *PlatformValues) HasIpv6() bool {
	if o != nil && o.Ipv6 != nil {
		return true
	}

	return false
}

// SetIpv6 gets a reference to the given int64 and assigns it to the Ipv6 field.
func (o *PlatformValues) SetIpv6(v int64) {
	o.Ipv6 = &v
}

// GetOtfp returns the Otfp field value if set, zero value otherwise.
func (o *PlatformValues) GetOtfp() int64 {
	if o == nil || o.Otfp == nil {
		var ret int64
		return ret
	}
	return *o.Otfp
}

// GetOtfpOk returns a tuple with the Otfp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetOtfpOk() (*int64, bool) {
	if o == nil || o.Otfp == nil {
		return nil, false
	}
	return o.Otfp, true
}

// HasOtfp returns a boolean if a field has been set.
func (o *PlatformValues) HasOtfp() bool {
	if o != nil && o.Otfp != nil {
		return true
	}

	return false
}

// SetOtfp gets a reference to the given int64 and assigns it to the Otfp field.
func (o *PlatformValues) SetOtfp(v int64) {
	o.Otfp = &v
}

// GetOtfpRespBodyBytes returns the OtfpRespBodyBytes field value if set, zero value otherwise.
func (o *PlatformValues) GetOtfpRespBodyBytes() int64 {
	if o == nil || o.OtfpRespBodyBytes == nil {
		var ret int64
		return ret
	}
	return *o.OtfpRespBodyBytes
}

// GetOtfpRespBodyBytesOk returns a tuple with the OtfpRespBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetOtfpRespBodyBytesOk() (*int64, bool) {
	if o == nil || o.OtfpRespBodyBytes == nil {
		return nil, false
	}
	return o.OtfpRespBodyBytes, true
}

// HasOtfpRespBodyBytes returns a boolean if a field has been set.
func (o *PlatformValues) HasOtfpRespBodyBytes() bool {
	if o != nil && o.OtfpRespBodyBytes != nil {
		return true
	}

	return false
}

// SetOtfpRespBodyBytes gets a reference to the given int64 and assigns it to the OtfpRespBodyBytes field.
func (o *PlatformValues) SetOtfpRespBodyBytes(v int64) {
	o.OtfpRespBodyBytes = &v
}

// GetOtfpRespHeaderBytes returns the OtfpRespHeaderBytes field value if set, zero value otherwise.
func (o *PlatformValues) GetOtfpRespHeaderBytes() int64 {
	if o == nil || o.OtfpRespHeaderBytes == nil {
		var ret int64
		return ret
	}
	return *o.OtfpRespHeaderBytes
}

// GetOtfpRespHeaderBytesOk returns a tuple with the OtfpRespHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetOtfpRespHeaderBytesOk() (*int64, bool) {
	if o == nil || o.OtfpRespHeaderBytes == nil {
		return nil, false
	}
	return o.OtfpRespHeaderBytes, true
}

// HasOtfpRespHeaderBytes returns a boolean if a field has been set.
func (o *PlatformValues) HasOtfpRespHeaderBytes() bool {
	if o != nil && o.OtfpRespHeaderBytes != nil {
		return true
	}

	return false
}

// SetOtfpRespHeaderBytes gets a reference to the given int64 and assigns it to the OtfpRespHeaderBytes field.
func (o *PlatformValues) SetOtfpRespHeaderBytes(v int64) {
	o.OtfpRespHeaderBytes = &v
}

// GetOtfpShieldRespBodyBytes returns the OtfpShieldRespBodyBytes field value if set, zero value otherwise.
func (o *PlatformValues) GetOtfpShieldRespBodyBytes() int64 {
	if o == nil || o.OtfpShieldRespBodyBytes == nil {
		var ret int64
		return ret
	}
	return *o.OtfpShieldRespBodyBytes
}

// GetOtfpShieldRespBodyBytesOk returns a tuple with the OtfpShieldRespBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetOtfpShieldRespBodyBytesOk() (*int64, bool) {
	if o == nil || o.OtfpShieldRespBodyBytes == nil {
		return nil, false
	}
	return o.OtfpShieldRespBodyBytes, true
}

// HasOtfpShieldRespBodyBytes returns a boolean if a field has been set.
func (o *PlatformValues) HasOtfpShieldRespBodyBytes() bool {
	if o != nil && o.OtfpShieldRespBodyBytes != nil {
		return true
	}

	return false
}

// SetOtfpShieldRespBodyBytes gets a reference to the given int64 and assigns it to the OtfpShieldRespBodyBytes field.
func (o *PlatformValues) SetOtfpShieldRespBodyBytes(v int64) {
	o.OtfpShieldRespBodyBytes = &v
}

// GetOtfpShieldRespHeaderBytes returns the OtfpShieldRespHeaderBytes field value if set, zero value otherwise.
func (o *PlatformValues) GetOtfpShieldRespHeaderBytes() int64 {
	if o == nil || o.OtfpShieldRespHeaderBytes == nil {
		var ret int64
		return ret
	}
	return *o.OtfpShieldRespHeaderBytes
}

// GetOtfpShieldRespHeaderBytesOk returns a tuple with the OtfpShieldRespHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetOtfpShieldRespHeaderBytesOk() (*int64, bool) {
	if o == nil || o.OtfpShieldRespHeaderBytes == nil {
		return nil, false
	}
	return o.OtfpShieldRespHeaderBytes, true
}

// HasOtfpShieldRespHeaderBytes returns a boolean if a field has been set.
func (o *PlatformValues) HasOtfpShieldRespHeaderBytes() bool {
	if o != nil && o.OtfpShieldRespHeaderBytes != nil {
		return true
	}

	return false
}

// SetOtfpShieldRespHeaderBytes gets a reference to the given int64 and assigns it to the OtfpShieldRespHeaderBytes field.
func (o *PlatformValues) SetOtfpShieldRespHeaderBytes(v int64) {
	o.OtfpShieldRespHeaderBytes = &v
}

// GetOtfpManifests returns the OtfpManifests field value if set, zero value otherwise.
func (o *PlatformValues) GetOtfpManifests() int64 {
	if o == nil || o.OtfpManifests == nil {
		var ret int64
		return ret
	}
	return *o.OtfpManifests
}

// GetOtfpManifestsOk returns a tuple with the OtfpManifests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetOtfpManifestsOk() (*int64, bool) {
	if o == nil || o.OtfpManifests == nil {
		return nil, false
	}
	return o.OtfpManifests, true
}

// HasOtfpManifests returns a boolean if a field has been set.
func (o *PlatformValues) HasOtfpManifests() bool {
	if o != nil && o.OtfpManifests != nil {
		return true
	}

	return false
}

// SetOtfpManifests gets a reference to the given int64 and assigns it to the OtfpManifests field.
func (o *PlatformValues) SetOtfpManifests(v int64) {
	o.OtfpManifests = &v
}

// GetOtfpDeliverTime returns the OtfpDeliverTime field value if set, zero value otherwise.
func (o *PlatformValues) GetOtfpDeliverTime() float32 {
	if o == nil || o.OtfpDeliverTime == nil {
		var ret float32
		return ret
	}
	return *o.OtfpDeliverTime
}

// GetOtfpDeliverTimeOk returns a tuple with the OtfpDeliverTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetOtfpDeliverTimeOk() (*float32, bool) {
	if o == nil || o.OtfpDeliverTime == nil {
		return nil, false
	}
	return o.OtfpDeliverTime, true
}

// HasOtfpDeliverTime returns a boolean if a field has been set.
func (o *PlatformValues) HasOtfpDeliverTime() bool {
	if o != nil && o.OtfpDeliverTime != nil {
		return true
	}

	return false
}

// SetOtfpDeliverTime gets a reference to the given float32 and assigns it to the OtfpDeliverTime field.
func (o *PlatformValues) SetOtfpDeliverTime(v float32) {
	o.OtfpDeliverTime = &v
}

// GetOtfpShieldTime returns the OtfpShieldTime field value if set, zero value otherwise.
func (o *PlatformValues) GetOtfpShieldTime() float32 {
	if o == nil || o.OtfpShieldTime == nil {
		var ret float32
		return ret
	}
	return *o.OtfpShieldTime
}

// GetOtfpShieldTimeOk returns a tuple with the OtfpShieldTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetOtfpShieldTimeOk() (*float32, bool) {
	if o == nil || o.OtfpShieldTime == nil {
		return nil, false
	}
	return o.OtfpShieldTime, true
}

// HasOtfpShieldTime returns a boolean if a field has been set.
func (o *PlatformValues) HasOtfpShieldTime() bool {
	if o != nil && o.OtfpShieldTime != nil {
		return true
	}

	return false
}

// SetOtfpShieldTime gets a reference to the given float32 and assigns it to the OtfpShieldTime field.
func (o *PlatformValues) SetOtfpShieldTime(v float32) {
	o.OtfpShieldTime = &v
}

// GetVideo returns the Video field value if set, zero value otherwise.
func (o *PlatformValues) GetVideo() int64 {
	if o == nil || o.Video == nil {
		var ret int64
		return ret
	}
	return *o.Video
}

// GetVideoOk returns a tuple with the Video field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetVideoOk() (*int64, bool) {
	if o == nil || o.Video == nil {
		return nil, false
	}
	return o.Video, true
}

// HasVideo returns a boolean if a field has been set.
func (o *PlatformValues) HasVideo() bool {
	if o != nil && o.Video != nil {
		return true
	}

	return false
}

// SetVideo gets a reference to the given int64 and assigns it to the Video field.
func (o *PlatformValues) SetVideo(v int64) {
	o.Video = &v
}

// GetPci returns the Pci field value if set, zero value otherwise.
func (o *PlatformValues) GetPci() int64 {
	if o == nil || o.Pci == nil {
		var ret int64
		return ret
	}
	return *o.Pci
}

// GetPciOk returns a tuple with the Pci field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetPciOk() (*int64, bool) {
	if o == nil || o.Pci == nil {
		return nil, false
	}
	return o.Pci, true
}

// HasPci returns a boolean if a field has been set.
func (o *PlatformValues) HasPci() bool {
	if o != nil && o.Pci != nil {
		return true
	}

	return false
}

// SetPci gets a reference to the given int64 and assigns it to the Pci field.
func (o *PlatformValues) SetPci(v int64) {
	o.Pci = &v
}

// GetLog returns the Log field value if set, zero value otherwise.
func (o *PlatformValues) GetLog() int64 {
	if o == nil || o.Log == nil {
		var ret int64
		return ret
	}
	return *o.Log
}

// GetLogOk returns a tuple with the Log field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetLogOk() (*int64, bool) {
	if o == nil || o.Log == nil {
		return nil, false
	}
	return o.Log, true
}

// HasLog returns a boolean if a field has been set.
func (o *PlatformValues) HasLog() bool {
	if o != nil && o.Log != nil {
		return true
	}

	return false
}

// SetLog gets a reference to the given int64 and assigns it to the Log field.
func (o *PlatformValues) SetLog(v int64) {
	o.Log = &v
}

// GetLogBytes returns the LogBytes field value if set, zero value otherwise.
func (o *PlatformValues) GetLogBytes() int64 {
	if o == nil || o.LogBytes == nil {
		var ret int64
		return ret
	}
	return *o.LogBytes
}

// GetLogBytesOk returns a tuple with the LogBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetLogBytesOk() (*int64, bool) {
	if o == nil || o.LogBytes == nil {
		return nil, false
	}
	return o.LogBytes, true
}

// HasLogBytes returns a boolean if a field has been set.
func (o *PlatformValues) HasLogBytes() bool {
	if o != nil && o.LogBytes != nil {
		return true
	}

	return false
}

// SetLogBytes gets a reference to the given int64 and assigns it to the LogBytes field.
func (o *PlatformValues) SetLogBytes(v int64) {
	o.LogBytes = &v
}

// GetHTTP2 returns the HTTP2 field value if set, zero value otherwise.
func (o *PlatformValues) GetHTTP2() int64 {
	if o == nil || o.HTTP2 == nil {
		var ret int64
		return ret
	}
	return *o.HTTP2
}

// GetHTTP2Ok returns a tuple with the HTTP2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetHTTP2Ok() (*int64, bool) {
	if o == nil || o.HTTP2 == nil {
		return nil, false
	}
	return o.HTTP2, true
}

// HasHTTP2 returns a boolean if a field has been set.
func (o *PlatformValues) HasHTTP2() bool {
	if o != nil && o.HTTP2 != nil {
		return true
	}

	return false
}

// SetHTTP2 gets a reference to the given int64 and assigns it to the HTTP2 field.
func (o *PlatformValues) SetHTTP2(v int64) {
	o.HTTP2 = &v
}

// GetHTTP3 returns the HTTP3 field value if set, zero value otherwise.
func (o *PlatformValues) GetHTTP3() int64 {
	if o == nil || o.HTTP3 == nil {
		var ret int64
		return ret
	}
	return *o.HTTP3
}

// GetHTTP3Ok returns a tuple with the HTTP3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetHTTP3Ok() (*int64, bool) {
	if o == nil || o.HTTP3 == nil {
		return nil, false
	}
	return o.HTTP3, true
}

// HasHTTP3 returns a boolean if a field has been set.
func (o *PlatformValues) HasHTTP3() bool {
	if o != nil && o.HTTP3 != nil {
		return true
	}

	return false
}

// SetHTTP3 gets a reference to the given int64 and assigns it to the HTTP3 field.
func (o *PlatformValues) SetHTTP3(v int64) {
	o.HTTP3 = &v
}

// GetAttackReqBodyBytes returns the AttackReqBodyBytes field value if set, zero value otherwise.
func (o *PlatformValues) GetAttackReqBodyBytes() int64 {
	if o == nil || o.AttackReqBodyBytes == nil {
		var ret int64
		return ret
	}
	return *o.AttackReqBodyBytes
}

// GetAttackReqBodyBytesOk returns a tuple with the AttackReqBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetAttackReqBodyBytesOk() (*int64, bool) {
	if o == nil || o.AttackReqBodyBytes == nil {
		return nil, false
	}
	return o.AttackReqBodyBytes, true
}

// HasAttackReqBodyBytes returns a boolean if a field has been set.
func (o *PlatformValues) HasAttackReqBodyBytes() bool {
	if o != nil && o.AttackReqBodyBytes != nil {
		return true
	}

	return false
}

// SetAttackReqBodyBytes gets a reference to the given int64 and assigns it to the AttackReqBodyBytes field.
func (o *PlatformValues) SetAttackReqBodyBytes(v int64) {
	o.AttackReqBodyBytes = &v
}

// GetAttackReqHeaderBytes returns the AttackReqHeaderBytes field value if set, zero value otherwise.
func (o *PlatformValues) GetAttackReqHeaderBytes() int64 {
	if o == nil || o.AttackReqHeaderBytes == nil {
		var ret int64
		return ret
	}
	return *o.AttackReqHeaderBytes
}

// GetAttackReqHeaderBytesOk returns a tuple with the AttackReqHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetAttackReqHeaderBytesOk() (*int64, bool) {
	if o == nil || o.AttackReqHeaderBytes == nil {
		return nil, false
	}
	return o.AttackReqHeaderBytes, true
}

// HasAttackReqHeaderBytes returns a boolean if a field has been set.
func (o *PlatformValues) HasAttackReqHeaderBytes() bool {
	if o != nil && o.AttackReqHeaderBytes != nil {
		return true
	}

	return false
}

// SetAttackReqHeaderBytes gets a reference to the given int64 and assigns it to the AttackReqHeaderBytes field.
func (o *PlatformValues) SetAttackReqHeaderBytes(v int64) {
	o.AttackReqHeaderBytes = &v
}

// GetAttackLoggedReqBodyBytes returns the AttackLoggedReqBodyBytes field value if set, zero value otherwise.
func (o *PlatformValues) GetAttackLoggedReqBodyBytes() int64 {
	if o == nil || o.AttackLoggedReqBodyBytes == nil {
		var ret int64
		return ret
	}
	return *o.AttackLoggedReqBodyBytes
}

// GetAttackLoggedReqBodyBytesOk returns a tuple with the AttackLoggedReqBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetAttackLoggedReqBodyBytesOk() (*int64, bool) {
	if o == nil || o.AttackLoggedReqBodyBytes == nil {
		return nil, false
	}
	return o.AttackLoggedReqBodyBytes, true
}

// HasAttackLoggedReqBodyBytes returns a boolean if a field has been set.
func (o *PlatformValues) HasAttackLoggedReqBodyBytes() bool {
	if o != nil && o.AttackLoggedReqBodyBytes != nil {
		return true
	}

	return false
}

// SetAttackLoggedReqBodyBytes gets a reference to the given int64 and assigns it to the AttackLoggedReqBodyBytes field.
func (o *PlatformValues) SetAttackLoggedReqBodyBytes(v int64) {
	o.AttackLoggedReqBodyBytes = &v
}

// GetAttackLoggedReqHeaderBytes returns the AttackLoggedReqHeaderBytes field value if set, zero value otherwise.
func (o *PlatformValues) GetAttackLoggedReqHeaderBytes() int64 {
	if o == nil || o.AttackLoggedReqHeaderBytes == nil {
		var ret int64
		return ret
	}
	return *o.AttackLoggedReqHeaderBytes
}

// GetAttackLoggedReqHeaderBytesOk returns a tuple with the AttackLoggedReqHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetAttackLoggedReqHeaderBytesOk() (*int64, bool) {
	if o == nil || o.AttackLoggedReqHeaderBytes == nil {
		return nil, false
	}
	return o.AttackLoggedReqHeaderBytes, true
}

// HasAttackLoggedReqHeaderBytes returns a boolean if a field has been set.
func (o *PlatformValues) HasAttackLoggedReqHeaderBytes() bool {
	if o != nil && o.AttackLoggedReqHeaderBytes != nil {
		return true
	}

	return false
}

// SetAttackLoggedReqHeaderBytes gets a reference to the given int64 and assigns it to the AttackLoggedReqHeaderBytes field.
func (o *PlatformValues) SetAttackLoggedReqHeaderBytes(v int64) {
	o.AttackLoggedReqHeaderBytes = &v
}

// GetAttackBlockedReqBodyBytes returns the AttackBlockedReqBodyBytes field value if set, zero value otherwise.
func (o *PlatformValues) GetAttackBlockedReqBodyBytes() int64 {
	if o == nil || o.AttackBlockedReqBodyBytes == nil {
		var ret int64
		return ret
	}
	return *o.AttackBlockedReqBodyBytes
}

// GetAttackBlockedReqBodyBytesOk returns a tuple with the AttackBlockedReqBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetAttackBlockedReqBodyBytesOk() (*int64, bool) {
	if o == nil || o.AttackBlockedReqBodyBytes == nil {
		return nil, false
	}
	return o.AttackBlockedReqBodyBytes, true
}

// HasAttackBlockedReqBodyBytes returns a boolean if a field has been set.
func (o *PlatformValues) HasAttackBlockedReqBodyBytes() bool {
	if o != nil && o.AttackBlockedReqBodyBytes != nil {
		return true
	}

	return false
}

// SetAttackBlockedReqBodyBytes gets a reference to the given int64 and assigns it to the AttackBlockedReqBodyBytes field.
func (o *PlatformValues) SetAttackBlockedReqBodyBytes(v int64) {
	o.AttackBlockedReqBodyBytes = &v
}

// GetAttackBlockedReqHeaderBytes returns the AttackBlockedReqHeaderBytes field value if set, zero value otherwise.
func (o *PlatformValues) GetAttackBlockedReqHeaderBytes() int64 {
	if o == nil || o.AttackBlockedReqHeaderBytes == nil {
		var ret int64
		return ret
	}
	return *o.AttackBlockedReqHeaderBytes
}

// GetAttackBlockedReqHeaderBytesOk returns a tuple with the AttackBlockedReqHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetAttackBlockedReqHeaderBytesOk() (*int64, bool) {
	if o == nil || o.AttackBlockedReqHeaderBytes == nil {
		return nil, false
	}
	return o.AttackBlockedReqHeaderBytes, true
}

// HasAttackBlockedReqHeaderBytes returns a boolean if a field has been set.
func (o *PlatformValues) HasAttackBlockedReqHeaderBytes() bool {
	if o != nil && o.AttackBlockedReqHeaderBytes != nil {
		return true
	}

	return false
}

// SetAttackBlockedReqHeaderBytes gets a reference to the given int64 and assigns it to the AttackBlockedReqHeaderBytes field.
func (o *PlatformValues) SetAttackBlockedReqHeaderBytes(v int64) {
	o.AttackBlockedReqHeaderBytes = &v
}

// GetAttackPassedReqBodyBytes returns the AttackPassedReqBodyBytes field value if set, zero value otherwise.
func (o *PlatformValues) GetAttackPassedReqBodyBytes() int64 {
	if o == nil || o.AttackPassedReqBodyBytes == nil {
		var ret int64
		return ret
	}
	return *o.AttackPassedReqBodyBytes
}

// GetAttackPassedReqBodyBytesOk returns a tuple with the AttackPassedReqBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetAttackPassedReqBodyBytesOk() (*int64, bool) {
	if o == nil || o.AttackPassedReqBodyBytes == nil {
		return nil, false
	}
	return o.AttackPassedReqBodyBytes, true
}

// HasAttackPassedReqBodyBytes returns a boolean if a field has been set.
func (o *PlatformValues) HasAttackPassedReqBodyBytes() bool {
	if o != nil && o.AttackPassedReqBodyBytes != nil {
		return true
	}

	return false
}

// SetAttackPassedReqBodyBytes gets a reference to the given int64 and assigns it to the AttackPassedReqBodyBytes field.
func (o *PlatformValues) SetAttackPassedReqBodyBytes(v int64) {
	o.AttackPassedReqBodyBytes = &v
}

// GetAttackPassedReqHeaderBytes returns the AttackPassedReqHeaderBytes field value if set, zero value otherwise.
func (o *PlatformValues) GetAttackPassedReqHeaderBytes() int64 {
	if o == nil || o.AttackPassedReqHeaderBytes == nil {
		var ret int64
		return ret
	}
	return *o.AttackPassedReqHeaderBytes
}

// GetAttackPassedReqHeaderBytesOk returns a tuple with the AttackPassedReqHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetAttackPassedReqHeaderBytesOk() (*int64, bool) {
	if o == nil || o.AttackPassedReqHeaderBytes == nil {
		return nil, false
	}
	return o.AttackPassedReqHeaderBytes, true
}

// HasAttackPassedReqHeaderBytes returns a boolean if a field has been set.
func (o *PlatformValues) HasAttackPassedReqHeaderBytes() bool {
	if o != nil && o.AttackPassedReqHeaderBytes != nil {
		return true
	}

	return false
}

// SetAttackPassedReqHeaderBytes gets a reference to the given int64 and assigns it to the AttackPassedReqHeaderBytes field.
func (o *PlatformValues) SetAttackPassedReqHeaderBytes(v int64) {
	o.AttackPassedReqHeaderBytes = &v
}

// GetAttackRespSynthBytes returns the AttackRespSynthBytes field value if set, zero value otherwise.
func (o *PlatformValues) GetAttackRespSynthBytes() int64 {
	if o == nil || o.AttackRespSynthBytes == nil {
		var ret int64
		return ret
	}
	return *o.AttackRespSynthBytes
}

// GetAttackRespSynthBytesOk returns a tuple with the AttackRespSynthBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetAttackRespSynthBytesOk() (*int64, bool) {
	if o == nil || o.AttackRespSynthBytes == nil {
		return nil, false
	}
	return o.AttackRespSynthBytes, true
}

// HasAttackRespSynthBytes returns a boolean if a field has been set.
func (o *PlatformValues) HasAttackRespSynthBytes() bool {
	if o != nil && o.AttackRespSynthBytes != nil {
		return true
	}

	return false
}

// SetAttackRespSynthBytes gets a reference to the given int64 and assigns it to the AttackRespSynthBytes field.
func (o *PlatformValues) SetAttackRespSynthBytes(v int64) {
	o.AttackRespSynthBytes = &v
}

// GetImgopto returns the Imgopto field value if set, zero value otherwise.
func (o *PlatformValues) GetImgopto() int64 {
	if o == nil || o.Imgopto == nil {
		var ret int64
		return ret
	}
	return *o.Imgopto
}

// GetImgoptoOk returns a tuple with the Imgopto field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetImgoptoOk() (*int64, bool) {
	if o == nil || o.Imgopto == nil {
		return nil, false
	}
	return o.Imgopto, true
}

// HasImgopto returns a boolean if a field has been set.
func (o *PlatformValues) HasImgopto() bool {
	if o != nil && o.Imgopto != nil {
		return true
	}

	return false
}

// SetImgopto gets a reference to the given int64 and assigns it to the Imgopto field.
func (o *PlatformValues) SetImgopto(v int64) {
	o.Imgopto = &v
}

// GetImgoptoInputImageBytes returns the ImgoptoInputImageBytes field value if set, zero value otherwise.
func (o *PlatformValues) GetImgoptoInputImageBytes() int64 {
	if o == nil || o.ImgoptoInputImageBytes == nil {
		var ret int64
		return ret
	}
	return *o.ImgoptoInputImageBytes
}

// GetImgoptoInputImageBytesOk returns a tuple with the ImgoptoInputImageBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetImgoptoInputImageBytesOk() (*int64, bool) {
	if o == nil || o.ImgoptoInputImageBytes == nil {
		return nil, false
	}
	return o.ImgoptoInputImageBytes, true
}

// HasImgoptoInputImageBytes returns a boolean if a field has been set.
func (o *PlatformValues) HasImgoptoInputImageBytes() bool {
	if o != nil && o.ImgoptoInputImageBytes != nil {
		return true
	}

	return false
}

// SetImgoptoInputImageBytes gets a reference to the given int64 and assigns it to the ImgoptoInputImageBytes field.
func (o *PlatformValues) SetImgoptoInputImageBytes(v int64) {
	o.ImgoptoInputImageBytes = &v
}

// GetImgoptoOutputImageBytes returns the ImgoptoOutputImageBytes field value if set, zero value otherwise.
func (o *PlatformValues) GetImgoptoOutputImageBytes() int64 {
	if o == nil || o.ImgoptoOutputImageBytes == nil {
		var ret int64
		return ret
	}
	return *o.ImgoptoOutputImageBytes
}

// GetImgoptoOutputImageBytesOk returns a tuple with the ImgoptoOutputImageBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetImgoptoOutputImageBytesOk() (*int64, bool) {
	if o == nil || o.ImgoptoOutputImageBytes == nil {
		return nil, false
	}
	return o.ImgoptoOutputImageBytes, true
}

// HasImgoptoOutputImageBytes returns a boolean if a field has been set.
func (o *PlatformValues) HasImgoptoOutputImageBytes() bool {
	if o != nil && o.ImgoptoOutputImageBytes != nil {
		return true
	}

	return false
}

// SetImgoptoOutputImageBytes gets a reference to the given int64 and assigns it to the ImgoptoOutputImageBytes field.
func (o *PlatformValues) SetImgoptoOutputImageBytes(v int64) {
	o.ImgoptoOutputImageBytes = &v
}

// GetImgoptoRespBodyBytes returns the ImgoptoRespBodyBytes field value if set, zero value otherwise.
func (o *PlatformValues) GetImgoptoRespBodyBytes() int64 {
	if o == nil || o.ImgoptoRespBodyBytes == nil {
		var ret int64
		return ret
	}
	return *o.ImgoptoRespBodyBytes
}

// GetImgoptoRespBodyBytesOk returns a tuple with the ImgoptoRespBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetImgoptoRespBodyBytesOk() (*int64, bool) {
	if o == nil || o.ImgoptoRespBodyBytes == nil {
		return nil, false
	}
	return o.ImgoptoRespBodyBytes, true
}

// HasImgoptoRespBodyBytes returns a boolean if a field has been set.
func (o *PlatformValues) HasImgoptoRespBodyBytes() bool {
	if o != nil && o.ImgoptoRespBodyBytes != nil {
		return true
	}

	return false
}

// SetImgoptoRespBodyBytes gets a reference to the given int64 and assigns it to the ImgoptoRespBodyBytes field.
func (o *PlatformValues) SetImgoptoRespBodyBytes(v int64) {
	o.ImgoptoRespBodyBytes = &v
}

// GetImgoptoRespHeaderBytes returns the ImgoptoRespHeaderBytes field value if set, zero value otherwise.
func (o *PlatformValues) GetImgoptoRespHeaderBytes() int64 {
	if o == nil || o.ImgoptoRespHeaderBytes == nil {
		var ret int64
		return ret
	}
	return *o.ImgoptoRespHeaderBytes
}

// GetImgoptoRespHeaderBytesOk returns a tuple with the ImgoptoRespHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetImgoptoRespHeaderBytesOk() (*int64, bool) {
	if o == nil || o.ImgoptoRespHeaderBytes == nil {
		return nil, false
	}
	return o.ImgoptoRespHeaderBytes, true
}

// HasImgoptoRespHeaderBytes returns a boolean if a field has been set.
func (o *PlatformValues) HasImgoptoRespHeaderBytes() bool {
	if o != nil && o.ImgoptoRespHeaderBytes != nil {
		return true
	}

	return false
}

// SetImgoptoRespHeaderBytes gets a reference to the given int64 and assigns it to the ImgoptoRespHeaderBytes field.
func (o *PlatformValues) SetImgoptoRespHeaderBytes(v int64) {
	o.ImgoptoRespHeaderBytes = &v
}

// GetImgoptoShield returns the ImgoptoShield field value if set, zero value otherwise.
func (o *PlatformValues) GetImgoptoShield() int64 {
	if o == nil || o.ImgoptoShield == nil {
		var ret int64
		return ret
	}
	return *o.ImgoptoShield
}

// GetImgoptoShieldOk returns a tuple with the ImgoptoShield field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetImgoptoShieldOk() (*int64, bool) {
	if o == nil || o.ImgoptoShield == nil {
		return nil, false
	}
	return o.ImgoptoShield, true
}

// HasImgoptoShield returns a boolean if a field has been set.
func (o *PlatformValues) HasImgoptoShield() bool {
	if o != nil && o.ImgoptoShield != nil {
		return true
	}

	return false
}

// SetImgoptoShield gets a reference to the given int64 and assigns it to the ImgoptoShield field.
func (o *PlatformValues) SetImgoptoShield(v int64) {
	o.ImgoptoShield = &v
}

// GetImgoptoShieldRespBodyBytes returns the ImgoptoShieldRespBodyBytes field value if set, zero value otherwise.
func (o *PlatformValues) GetImgoptoShieldRespBodyBytes() int64 {
	if o == nil || o.ImgoptoShieldRespBodyBytes == nil {
		var ret int64
		return ret
	}
	return *o.ImgoptoShieldRespBodyBytes
}

// GetImgoptoShieldRespBodyBytesOk returns a tuple with the ImgoptoShieldRespBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetImgoptoShieldRespBodyBytesOk() (*int64, bool) {
	if o == nil || o.ImgoptoShieldRespBodyBytes == nil {
		return nil, false
	}
	return o.ImgoptoShieldRespBodyBytes, true
}

// HasImgoptoShieldRespBodyBytes returns a boolean if a field has been set.
func (o *PlatformValues) HasImgoptoShieldRespBodyBytes() bool {
	if o != nil && o.ImgoptoShieldRespBodyBytes != nil {
		return true
	}

	return false
}

// SetImgoptoShieldRespBodyBytes gets a reference to the given int64 and assigns it to the ImgoptoShieldRespBodyBytes field.
func (o *PlatformValues) SetImgoptoShieldRespBodyBytes(v int64) {
	o.ImgoptoShieldRespBodyBytes = &v
}

// GetImgoptoShieldRespHeaderBytes returns the ImgoptoShieldRespHeaderBytes field value if set, zero value otherwise.
func (o *PlatformValues) GetImgoptoShieldRespHeaderBytes() int64 {
	if o == nil || o.ImgoptoShieldRespHeaderBytes == nil {
		var ret int64
		return ret
	}
	return *o.ImgoptoShieldRespHeaderBytes
}

// GetImgoptoShieldRespHeaderBytesOk returns a tuple with the ImgoptoShieldRespHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetImgoptoShieldRespHeaderBytesOk() (*int64, bool) {
	if o == nil || o.ImgoptoShieldRespHeaderBytes == nil {
		return nil, false
	}
	return o.ImgoptoShieldRespHeaderBytes, true
}

// HasImgoptoShieldRespHeaderBytes returns a boolean if a field has been set.
func (o *PlatformValues) HasImgoptoShieldRespHeaderBytes() bool {
	if o != nil && o.ImgoptoShieldRespHeaderBytes != nil {
		return true
	}

	return false
}

// SetImgoptoShieldRespHeaderBytes gets a reference to the given int64 and assigns it to the ImgoptoShieldRespHeaderBytes field.
func (o *PlatformValues) SetImgoptoShieldRespHeaderBytes(v int64) {
	o.ImgoptoShieldRespHeaderBytes = &v
}

// GetImgoptoTransforms returns the ImgoptoTransforms field value if set, zero value otherwise.
func (o *PlatformValues) GetImgoptoTransforms() int64 {
	if o == nil || o.ImgoptoTransforms == nil {
		var ret int64
		return ret
	}
	return *o.ImgoptoTransforms
}

// GetImgoptoTransformsOk returns a tuple with the ImgoptoTransforms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetImgoptoTransformsOk() (*int64, bool) {
	if o == nil || o.ImgoptoTransforms == nil {
		return nil, false
	}
	return o.ImgoptoTransforms, true
}

// HasImgoptoTransforms returns a boolean if a field has been set.
func (o *PlatformValues) HasImgoptoTransforms() bool {
	if o != nil && o.ImgoptoTransforms != nil {
		return true
	}

	return false
}

// SetImgoptoTransforms gets a reference to the given int64 and assigns it to the ImgoptoTransforms field.
func (o *PlatformValues) SetImgoptoTransforms(v int64) {
	o.ImgoptoTransforms = &v
}

// GetImgvideo returns the Imgvideo field value if set, zero value otherwise.
func (o *PlatformValues) GetImgvideo() int64 {
	if o == nil || o.Imgvideo == nil {
		var ret int64
		return ret
	}
	return *o.Imgvideo
}

// GetImgvideoOk returns a tuple with the Imgvideo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetImgvideoOk() (*int64, bool) {
	if o == nil || o.Imgvideo == nil {
		return nil, false
	}
	return o.Imgvideo, true
}

// HasImgvideo returns a boolean if a field has been set.
func (o *PlatformValues) HasImgvideo() bool {
	if o != nil && o.Imgvideo != nil {
		return true
	}

	return false
}

// SetImgvideo gets a reference to the given int64 and assigns it to the Imgvideo field.
func (o *PlatformValues) SetImgvideo(v int64) {
	o.Imgvideo = &v
}

// GetImgvideoFrames returns the ImgvideoFrames field value if set, zero value otherwise.
func (o *PlatformValues) GetImgvideoFrames() int64 {
	if o == nil || o.ImgvideoFrames == nil {
		var ret int64
		return ret
	}
	return *o.ImgvideoFrames
}

// GetImgvideoFramesOk returns a tuple with the ImgvideoFrames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetImgvideoFramesOk() (*int64, bool) {
	if o == nil || o.ImgvideoFrames == nil {
		return nil, false
	}
	return o.ImgvideoFrames, true
}

// HasImgvideoFrames returns a boolean if a field has been set.
func (o *PlatformValues) HasImgvideoFrames() bool {
	if o != nil && o.ImgvideoFrames != nil {
		return true
	}

	return false
}

// SetImgvideoFrames gets a reference to the given int64 and assigns it to the ImgvideoFrames field.
func (o *PlatformValues) SetImgvideoFrames(v int64) {
	o.ImgvideoFrames = &v
}

// GetImgvideoRespHeaderBytes returns the ImgvideoRespHeaderBytes field value if set, zero value otherwise.
func (o *PlatformValues) GetImgvideoRespHeaderBytes() int64 {
	if o == nil || o.ImgvideoRespHeaderBytes == nil {
		var ret int64
		return ret
	}
	return *o.ImgvideoRespHeaderBytes
}

// GetImgvideoRespHeaderBytesOk returns a tuple with the ImgvideoRespHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetImgvideoRespHeaderBytesOk() (*int64, bool) {
	if o == nil || o.ImgvideoRespHeaderBytes == nil {
		return nil, false
	}
	return o.ImgvideoRespHeaderBytes, true
}

// HasImgvideoRespHeaderBytes returns a boolean if a field has been set.
func (o *PlatformValues) HasImgvideoRespHeaderBytes() bool {
	if o != nil && o.ImgvideoRespHeaderBytes != nil {
		return true
	}

	return false
}

// SetImgvideoRespHeaderBytes gets a reference to the given int64 and assigns it to the ImgvideoRespHeaderBytes field.
func (o *PlatformValues) SetImgvideoRespHeaderBytes(v int64) {
	o.ImgvideoRespHeaderBytes = &v
}

// GetImgvideoRespBodyBytes returns the ImgvideoRespBodyBytes field value if set, zero value otherwise.
func (o *PlatformValues) GetImgvideoRespBodyBytes() int64 {
	if o == nil || o.ImgvideoRespBodyBytes == nil {
		var ret int64
		return ret
	}
	return *o.ImgvideoRespBodyBytes
}

// GetImgvideoRespBodyBytesOk returns a tuple with the ImgvideoRespBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetImgvideoRespBodyBytesOk() (*int64, bool) {
	if o == nil || o.ImgvideoRespBodyBytes == nil {
		return nil, false
	}
	return o.ImgvideoRespBodyBytes, true
}

// HasImgvideoRespBodyBytes returns a boolean if a field has been set.
func (o *PlatformValues) HasImgvideoRespBodyBytes() bool {
	if o != nil && o.ImgvideoRespBodyBytes != nil {
		return true
	}

	return false
}

// SetImgvideoRespBodyBytes gets a reference to the given int64 and assigns it to the ImgvideoRespBodyBytes field.
func (o *PlatformValues) SetImgvideoRespBodyBytes(v int64) {
	o.ImgvideoRespBodyBytes = &v
}

// GetImgvideoShieldRespHeaderBytes returns the ImgvideoShieldRespHeaderBytes field value if set, zero value otherwise.
func (o *PlatformValues) GetImgvideoShieldRespHeaderBytes() int64 {
	if o == nil || o.ImgvideoShieldRespHeaderBytes == nil {
		var ret int64
		return ret
	}
	return *o.ImgvideoShieldRespHeaderBytes
}

// GetImgvideoShieldRespHeaderBytesOk returns a tuple with the ImgvideoShieldRespHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetImgvideoShieldRespHeaderBytesOk() (*int64, bool) {
	if o == nil || o.ImgvideoShieldRespHeaderBytes == nil {
		return nil, false
	}
	return o.ImgvideoShieldRespHeaderBytes, true
}

// HasImgvideoShieldRespHeaderBytes returns a boolean if a field has been set.
func (o *PlatformValues) HasImgvideoShieldRespHeaderBytes() bool {
	if o != nil && o.ImgvideoShieldRespHeaderBytes != nil {
		return true
	}

	return false
}

// SetImgvideoShieldRespHeaderBytes gets a reference to the given int64 and assigns it to the ImgvideoShieldRespHeaderBytes field.
func (o *PlatformValues) SetImgvideoShieldRespHeaderBytes(v int64) {
	o.ImgvideoShieldRespHeaderBytes = &v
}

// GetImgvideoShieldRespBodyBytes returns the ImgvideoShieldRespBodyBytes field value if set, zero value otherwise.
func (o *PlatformValues) GetImgvideoShieldRespBodyBytes() int64 {
	if o == nil || o.ImgvideoShieldRespBodyBytes == nil {
		var ret int64
		return ret
	}
	return *o.ImgvideoShieldRespBodyBytes
}

// GetImgvideoShieldRespBodyBytesOk returns a tuple with the ImgvideoShieldRespBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetImgvideoShieldRespBodyBytesOk() (*int64, bool) {
	if o == nil || o.ImgvideoShieldRespBodyBytes == nil {
		return nil, false
	}
	return o.ImgvideoShieldRespBodyBytes, true
}

// HasImgvideoShieldRespBodyBytes returns a boolean if a field has been set.
func (o *PlatformValues) HasImgvideoShieldRespBodyBytes() bool {
	if o != nil && o.ImgvideoShieldRespBodyBytes != nil {
		return true
	}

	return false
}

// SetImgvideoShieldRespBodyBytes gets a reference to the given int64 and assigns it to the ImgvideoShieldRespBodyBytes field.
func (o *PlatformValues) SetImgvideoShieldRespBodyBytes(v int64) {
	o.ImgvideoShieldRespBodyBytes = &v
}

// GetImgvideoShield returns the ImgvideoShield field value if set, zero value otherwise.
func (o *PlatformValues) GetImgvideoShield() int64 {
	if o == nil || o.ImgvideoShield == nil {
		var ret int64
		return ret
	}
	return *o.ImgvideoShield
}

// GetImgvideoShieldOk returns a tuple with the ImgvideoShield field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetImgvideoShieldOk() (*int64, bool) {
	if o == nil || o.ImgvideoShield == nil {
		return nil, false
	}
	return o.ImgvideoShield, true
}

// HasImgvideoShield returns a boolean if a field has been set.
func (o *PlatformValues) HasImgvideoShield() bool {
	if o != nil && o.ImgvideoShield != nil {
		return true
	}

	return false
}

// SetImgvideoShield gets a reference to the given int64 and assigns it to the ImgvideoShield field.
func (o *PlatformValues) SetImgvideoShield(v int64) {
	o.ImgvideoShield = &v
}

// GetImgvideoShieldFrames returns the ImgvideoShieldFrames field value if set, zero value otherwise.
func (o *PlatformValues) GetImgvideoShieldFrames() int64 {
	if o == nil || o.ImgvideoShieldFrames == nil {
		var ret int64
		return ret
	}
	return *o.ImgvideoShieldFrames
}

// GetImgvideoShieldFramesOk returns a tuple with the ImgvideoShieldFrames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetImgvideoShieldFramesOk() (*int64, bool) {
	if o == nil || o.ImgvideoShieldFrames == nil {
		return nil, false
	}
	return o.ImgvideoShieldFrames, true
}

// HasImgvideoShieldFrames returns a boolean if a field has been set.
func (o *PlatformValues) HasImgvideoShieldFrames() bool {
	if o != nil && o.ImgvideoShieldFrames != nil {
		return true
	}

	return false
}

// SetImgvideoShieldFrames gets a reference to the given int64 and assigns it to the ImgvideoShieldFrames field.
func (o *PlatformValues) SetImgvideoShieldFrames(v int64) {
	o.ImgvideoShieldFrames = &v
}

// GetStatus200 returns the Status200 field value if set, zero value otherwise.
func (o *PlatformValues) GetStatus200() int64 {
	if o == nil || o.Status200 == nil {
		var ret int64
		return ret
	}
	return *o.Status200
}

// GetStatus200Ok returns a tuple with the Status200 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetStatus200Ok() (*int64, bool) {
	if o == nil || o.Status200 == nil {
		return nil, false
	}
	return o.Status200, true
}

// HasStatus200 returns a boolean if a field has been set.
func (o *PlatformValues) HasStatus200() bool {
	if o != nil && o.Status200 != nil {
		return true
	}

	return false
}

// SetStatus200 gets a reference to the given int64 and assigns it to the Status200 field.
func (o *PlatformValues) SetStatus200(v int64) {
	o.Status200 = &v
}

// GetStatus204 returns the Status204 field value if set, zero value otherwise.
func (o *PlatformValues) GetStatus204() int64 {
	if o == nil || o.Status204 == nil {
		var ret int64
		return ret
	}
	return *o.Status204
}

// GetStatus204Ok returns a tuple with the Status204 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetStatus204Ok() (*int64, bool) {
	if o == nil || o.Status204 == nil {
		return nil, false
	}
	return o.Status204, true
}

// HasStatus204 returns a boolean if a field has been set.
func (o *PlatformValues) HasStatus204() bool {
	if o != nil && o.Status204 != nil {
		return true
	}

	return false
}

// SetStatus204 gets a reference to the given int64 and assigns it to the Status204 field.
func (o *PlatformValues) SetStatus204(v int64) {
	o.Status204 = &v
}

// GetStatus206 returns the Status206 field value if set, zero value otherwise.
func (o *PlatformValues) GetStatus206() int64 {
	if o == nil || o.Status206 == nil {
		var ret int64
		return ret
	}
	return *o.Status206
}

// GetStatus206Ok returns a tuple with the Status206 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetStatus206Ok() (*int64, bool) {
	if o == nil || o.Status206 == nil {
		return nil, false
	}
	return o.Status206, true
}

// HasStatus206 returns a boolean if a field has been set.
func (o *PlatformValues) HasStatus206() bool {
	if o != nil && o.Status206 != nil {
		return true
	}

	return false
}

// SetStatus206 gets a reference to the given int64 and assigns it to the Status206 field.
func (o *PlatformValues) SetStatus206(v int64) {
	o.Status206 = &v
}

// GetStatus301 returns the Status301 field value if set, zero value otherwise.
func (o *PlatformValues) GetStatus301() int64 {
	if o == nil || o.Status301 == nil {
		var ret int64
		return ret
	}
	return *o.Status301
}

// GetStatus301Ok returns a tuple with the Status301 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetStatus301Ok() (*int64, bool) {
	if o == nil || o.Status301 == nil {
		return nil, false
	}
	return o.Status301, true
}

// HasStatus301 returns a boolean if a field has been set.
func (o *PlatformValues) HasStatus301() bool {
	if o != nil && o.Status301 != nil {
		return true
	}

	return false
}

// SetStatus301 gets a reference to the given int64 and assigns it to the Status301 field.
func (o *PlatformValues) SetStatus301(v int64) {
	o.Status301 = &v
}

// GetStatus302 returns the Status302 field value if set, zero value otherwise.
func (o *PlatformValues) GetStatus302() int64 {
	if o == nil || o.Status302 == nil {
		var ret int64
		return ret
	}
	return *o.Status302
}

// GetStatus302Ok returns a tuple with the Status302 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetStatus302Ok() (*int64, bool) {
	if o == nil || o.Status302 == nil {
		return nil, false
	}
	return o.Status302, true
}

// HasStatus302 returns a boolean if a field has been set.
func (o *PlatformValues) HasStatus302() bool {
	if o != nil && o.Status302 != nil {
		return true
	}

	return false
}

// SetStatus302 gets a reference to the given int64 and assigns it to the Status302 field.
func (o *PlatformValues) SetStatus302(v int64) {
	o.Status302 = &v
}

// GetStatus304 returns the Status304 field value if set, zero value otherwise.
func (o *PlatformValues) GetStatus304() int64 {
	if o == nil || o.Status304 == nil {
		var ret int64
		return ret
	}
	return *o.Status304
}

// GetStatus304Ok returns a tuple with the Status304 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetStatus304Ok() (*int64, bool) {
	if o == nil || o.Status304 == nil {
		return nil, false
	}
	return o.Status304, true
}

// HasStatus304 returns a boolean if a field has been set.
func (o *PlatformValues) HasStatus304() bool {
	if o != nil && o.Status304 != nil {
		return true
	}

	return false
}

// SetStatus304 gets a reference to the given int64 and assigns it to the Status304 field.
func (o *PlatformValues) SetStatus304(v int64) {
	o.Status304 = &v
}

// GetStatus400 returns the Status400 field value if set, zero value otherwise.
func (o *PlatformValues) GetStatus400() int64 {
	if o == nil || o.Status400 == nil {
		var ret int64
		return ret
	}
	return *o.Status400
}

// GetStatus400Ok returns a tuple with the Status400 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetStatus400Ok() (*int64, bool) {
	if o == nil || o.Status400 == nil {
		return nil, false
	}
	return o.Status400, true
}

// HasStatus400 returns a boolean if a field has been set.
func (o *PlatformValues) HasStatus400() bool {
	if o != nil && o.Status400 != nil {
		return true
	}

	return false
}

// SetStatus400 gets a reference to the given int64 and assigns it to the Status400 field.
func (o *PlatformValues) SetStatus400(v int64) {
	o.Status400 = &v
}

// GetStatus401 returns the Status401 field value if set, zero value otherwise.
func (o *PlatformValues) GetStatus401() int64 {
	if o == nil || o.Status401 == nil {
		var ret int64
		return ret
	}
	return *o.Status401
}

// GetStatus401Ok returns a tuple with the Status401 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetStatus401Ok() (*int64, bool) {
	if o == nil || o.Status401 == nil {
		return nil, false
	}
	return o.Status401, true
}

// HasStatus401 returns a boolean if a field has been set.
func (o *PlatformValues) HasStatus401() bool {
	if o != nil && o.Status401 != nil {
		return true
	}

	return false
}

// SetStatus401 gets a reference to the given int64 and assigns it to the Status401 field.
func (o *PlatformValues) SetStatus401(v int64) {
	o.Status401 = &v
}

// GetStatus403 returns the Status403 field value if set, zero value otherwise.
func (o *PlatformValues) GetStatus403() int64 {
	if o == nil || o.Status403 == nil {
		var ret int64
		return ret
	}
	return *o.Status403
}

// GetStatus403Ok returns a tuple with the Status403 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetStatus403Ok() (*int64, bool) {
	if o == nil || o.Status403 == nil {
		return nil, false
	}
	return o.Status403, true
}

// HasStatus403 returns a boolean if a field has been set.
func (o *PlatformValues) HasStatus403() bool {
	if o != nil && o.Status403 != nil {
		return true
	}

	return false
}

// SetStatus403 gets a reference to the given int64 and assigns it to the Status403 field.
func (o *PlatformValues) SetStatus403(v int64) {
	o.Status403 = &v
}

// GetStatus404 returns the Status404 field value if set, zero value otherwise.
func (o *PlatformValues) GetStatus404() int64 {
	if o == nil || o.Status404 == nil {
		var ret int64
		return ret
	}
	return *o.Status404
}

// GetStatus404Ok returns a tuple with the Status404 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetStatus404Ok() (*int64, bool) {
	if o == nil || o.Status404 == nil {
		return nil, false
	}
	return o.Status404, true
}

// HasStatus404 returns a boolean if a field has been set.
func (o *PlatformValues) HasStatus404() bool {
	if o != nil && o.Status404 != nil {
		return true
	}

	return false
}

// SetStatus404 gets a reference to the given int64 and assigns it to the Status404 field.
func (o *PlatformValues) SetStatus404(v int64) {
	o.Status404 = &v
}

// GetStatus406 returns the Status406 field value if set, zero value otherwise.
func (o *PlatformValues) GetStatus406() int64 {
	if o == nil || o.Status406 == nil {
		var ret int64
		return ret
	}
	return *o.Status406
}

// GetStatus406Ok returns a tuple with the Status406 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetStatus406Ok() (*int64, bool) {
	if o == nil || o.Status406 == nil {
		return nil, false
	}
	return o.Status406, true
}

// HasStatus406 returns a boolean if a field has been set.
func (o *PlatformValues) HasStatus406() bool {
	if o != nil && o.Status406 != nil {
		return true
	}

	return false
}

// SetStatus406 gets a reference to the given int64 and assigns it to the Status406 field.
func (o *PlatformValues) SetStatus406(v int64) {
	o.Status406 = &v
}

// GetStatus416 returns the Status416 field value if set, zero value otherwise.
func (o *PlatformValues) GetStatus416() int64 {
	if o == nil || o.Status416 == nil {
		var ret int64
		return ret
	}
	return *o.Status416
}

// GetStatus416Ok returns a tuple with the Status416 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetStatus416Ok() (*int64, bool) {
	if o == nil || o.Status416 == nil {
		return nil, false
	}
	return o.Status416, true
}

// HasStatus416 returns a boolean if a field has been set.
func (o *PlatformValues) HasStatus416() bool {
	if o != nil && o.Status416 != nil {
		return true
	}

	return false
}

// SetStatus416 gets a reference to the given int64 and assigns it to the Status416 field.
func (o *PlatformValues) SetStatus416(v int64) {
	o.Status416 = &v
}

// GetStatus429 returns the Status429 field value if set, zero value otherwise.
func (o *PlatformValues) GetStatus429() int64 {
	if o == nil || o.Status429 == nil {
		var ret int64
		return ret
	}
	return *o.Status429
}

// GetStatus429Ok returns a tuple with the Status429 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetStatus429Ok() (*int64, bool) {
	if o == nil || o.Status429 == nil {
		return nil, false
	}
	return o.Status429, true
}

// HasStatus429 returns a boolean if a field has been set.
func (o *PlatformValues) HasStatus429() bool {
	if o != nil && o.Status429 != nil {
		return true
	}

	return false
}

// SetStatus429 gets a reference to the given int64 and assigns it to the Status429 field.
func (o *PlatformValues) SetStatus429(v int64) {
	o.Status429 = &v
}

// GetStatus500 returns the Status500 field value if set, zero value otherwise.
func (o *PlatformValues) GetStatus500() int64 {
	if o == nil || o.Status500 == nil {
		var ret int64
		return ret
	}
	return *o.Status500
}

// GetStatus500Ok returns a tuple with the Status500 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetStatus500Ok() (*int64, bool) {
	if o == nil || o.Status500 == nil {
		return nil, false
	}
	return o.Status500, true
}

// HasStatus500 returns a boolean if a field has been set.
func (o *PlatformValues) HasStatus500() bool {
	if o != nil && o.Status500 != nil {
		return true
	}

	return false
}

// SetStatus500 gets a reference to the given int64 and assigns it to the Status500 field.
func (o *PlatformValues) SetStatus500(v int64) {
	o.Status500 = &v
}

// GetStatus501 returns the Status501 field value if set, zero value otherwise.
func (o *PlatformValues) GetStatus501() int64 {
	if o == nil || o.Status501 == nil {
		var ret int64
		return ret
	}
	return *o.Status501
}

// GetStatus501Ok returns a tuple with the Status501 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetStatus501Ok() (*int64, bool) {
	if o == nil || o.Status501 == nil {
		return nil, false
	}
	return o.Status501, true
}

// HasStatus501 returns a boolean if a field has been set.
func (o *PlatformValues) HasStatus501() bool {
	if o != nil && o.Status501 != nil {
		return true
	}

	return false
}

// SetStatus501 gets a reference to the given int64 and assigns it to the Status501 field.
func (o *PlatformValues) SetStatus501(v int64) {
	o.Status501 = &v
}

// GetStatus502 returns the Status502 field value if set, zero value otherwise.
func (o *PlatformValues) GetStatus502() int64 {
	if o == nil || o.Status502 == nil {
		var ret int64
		return ret
	}
	return *o.Status502
}

// GetStatus502Ok returns a tuple with the Status502 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetStatus502Ok() (*int64, bool) {
	if o == nil || o.Status502 == nil {
		return nil, false
	}
	return o.Status502, true
}

// HasStatus502 returns a boolean if a field has been set.
func (o *PlatformValues) HasStatus502() bool {
	if o != nil && o.Status502 != nil {
		return true
	}

	return false
}

// SetStatus502 gets a reference to the given int64 and assigns it to the Status502 field.
func (o *PlatformValues) SetStatus502(v int64) {
	o.Status502 = &v
}

// GetStatus503 returns the Status503 field value if set, zero value otherwise.
func (o *PlatformValues) GetStatus503() int64 {
	if o == nil || o.Status503 == nil {
		var ret int64
		return ret
	}
	return *o.Status503
}

// GetStatus503Ok returns a tuple with the Status503 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetStatus503Ok() (*int64, bool) {
	if o == nil || o.Status503 == nil {
		return nil, false
	}
	return o.Status503, true
}

// HasStatus503 returns a boolean if a field has been set.
func (o *PlatformValues) HasStatus503() bool {
	if o != nil && o.Status503 != nil {
		return true
	}

	return false
}

// SetStatus503 gets a reference to the given int64 and assigns it to the Status503 field.
func (o *PlatformValues) SetStatus503(v int64) {
	o.Status503 = &v
}

// GetStatus504 returns the Status504 field value if set, zero value otherwise.
func (o *PlatformValues) GetStatus504() int64 {
	if o == nil || o.Status504 == nil {
		var ret int64
		return ret
	}
	return *o.Status504
}

// GetStatus504Ok returns a tuple with the Status504 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetStatus504Ok() (*int64, bool) {
	if o == nil || o.Status504 == nil {
		return nil, false
	}
	return o.Status504, true
}

// HasStatus504 returns a boolean if a field has been set.
func (o *PlatformValues) HasStatus504() bool {
	if o != nil && o.Status504 != nil {
		return true
	}

	return false
}

// SetStatus504 gets a reference to the given int64 and assigns it to the Status504 field.
func (o *PlatformValues) SetStatus504(v int64) {
	o.Status504 = &v
}

// GetStatus505 returns the Status505 field value if set, zero value otherwise.
func (o *PlatformValues) GetStatus505() int64 {
	if o == nil || o.Status505 == nil {
		var ret int64
		return ret
	}
	return *o.Status505
}

// GetStatus505Ok returns a tuple with the Status505 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetStatus505Ok() (*int64, bool) {
	if o == nil || o.Status505 == nil {
		return nil, false
	}
	return o.Status505, true
}

// HasStatus505 returns a boolean if a field has been set.
func (o *PlatformValues) HasStatus505() bool {
	if o != nil && o.Status505 != nil {
		return true
	}

	return false
}

// SetStatus505 gets a reference to the given int64 and assigns it to the Status505 field.
func (o *PlatformValues) SetStatus505(v int64) {
	o.Status505 = &v
}

// GetStatus530 returns the Status530 field value if set, zero value otherwise.
func (o *PlatformValues) GetStatus530() int64 {
	if o == nil || o.Status530 == nil {
		var ret int64
		return ret
	}
	return *o.Status530
}

// GetStatus530Ok returns a tuple with the Status530 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetStatus530Ok() (*int64, bool) {
	if o == nil || o.Status530 == nil {
		return nil, false
	}
	return o.Status530, true
}

// HasStatus530 returns a boolean if a field has been set.
func (o *PlatformValues) HasStatus530() bool {
	if o != nil && o.Status530 != nil {
		return true
	}

	return false
}

// SetStatus530 gets a reference to the given int64 and assigns it to the Status530 field.
func (o *PlatformValues) SetStatus530(v int64) {
	o.Status530 = &v
}

// GetStatus1xx returns the Status1xx field value if set, zero value otherwise.
func (o *PlatformValues) GetStatus1xx() int64 {
	if o == nil || o.Status1xx == nil {
		var ret int64
		return ret
	}
	return *o.Status1xx
}

// GetStatus1xxOk returns a tuple with the Status1xx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetStatus1xxOk() (*int64, bool) {
	if o == nil || o.Status1xx == nil {
		return nil, false
	}
	return o.Status1xx, true
}

// HasStatus1xx returns a boolean if a field has been set.
func (o *PlatformValues) HasStatus1xx() bool {
	if o != nil && o.Status1xx != nil {
		return true
	}

	return false
}

// SetStatus1xx gets a reference to the given int64 and assigns it to the Status1xx field.
func (o *PlatformValues) SetStatus1xx(v int64) {
	o.Status1xx = &v
}

// GetStatus2xx returns the Status2xx field value if set, zero value otherwise.
func (o *PlatformValues) GetStatus2xx() int64 {
	if o == nil || o.Status2xx == nil {
		var ret int64
		return ret
	}
	return *o.Status2xx
}

// GetStatus2xxOk returns a tuple with the Status2xx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetStatus2xxOk() (*int64, bool) {
	if o == nil || o.Status2xx == nil {
		return nil, false
	}
	return o.Status2xx, true
}

// HasStatus2xx returns a boolean if a field has been set.
func (o *PlatformValues) HasStatus2xx() bool {
	if o != nil && o.Status2xx != nil {
		return true
	}

	return false
}

// SetStatus2xx gets a reference to the given int64 and assigns it to the Status2xx field.
func (o *PlatformValues) SetStatus2xx(v int64) {
	o.Status2xx = &v
}

// GetStatus3xx returns the Status3xx field value if set, zero value otherwise.
func (o *PlatformValues) GetStatus3xx() int64 {
	if o == nil || o.Status3xx == nil {
		var ret int64
		return ret
	}
	return *o.Status3xx
}

// GetStatus3xxOk returns a tuple with the Status3xx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetStatus3xxOk() (*int64, bool) {
	if o == nil || o.Status3xx == nil {
		return nil, false
	}
	return o.Status3xx, true
}

// HasStatus3xx returns a boolean if a field has been set.
func (o *PlatformValues) HasStatus3xx() bool {
	if o != nil && o.Status3xx != nil {
		return true
	}

	return false
}

// SetStatus3xx gets a reference to the given int64 and assigns it to the Status3xx field.
func (o *PlatformValues) SetStatus3xx(v int64) {
	o.Status3xx = &v
}

// GetStatus4xx returns the Status4xx field value if set, zero value otherwise.
func (o *PlatformValues) GetStatus4xx() int64 {
	if o == nil || o.Status4xx == nil {
		var ret int64
		return ret
	}
	return *o.Status4xx
}

// GetStatus4xxOk returns a tuple with the Status4xx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetStatus4xxOk() (*int64, bool) {
	if o == nil || o.Status4xx == nil {
		return nil, false
	}
	return o.Status4xx, true
}

// HasStatus4xx returns a boolean if a field has been set.
func (o *PlatformValues) HasStatus4xx() bool {
	if o != nil && o.Status4xx != nil {
		return true
	}

	return false
}

// SetStatus4xx gets a reference to the given int64 and assigns it to the Status4xx field.
func (o *PlatformValues) SetStatus4xx(v int64) {
	o.Status4xx = &v
}

// GetStatus5xx returns the Status5xx field value if set, zero value otherwise.
func (o *PlatformValues) GetStatus5xx() int64 {
	if o == nil || o.Status5xx == nil {
		var ret int64
		return ret
	}
	return *o.Status5xx
}

// GetStatus5xxOk returns a tuple with the Status5xx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetStatus5xxOk() (*int64, bool) {
	if o == nil || o.Status5xx == nil {
		return nil, false
	}
	return o.Status5xx, true
}

// HasStatus5xx returns a boolean if a field has been set.
func (o *PlatformValues) HasStatus5xx() bool {
	if o != nil && o.Status5xx != nil {
		return true
	}

	return false
}

// SetStatus5xx gets a reference to the given int64 and assigns it to the Status5xx field.
func (o *PlatformValues) SetStatus5xx(v int64) {
	o.Status5xx = &v
}

// GetObjectSize1k returns the ObjectSize1k field value if set, zero value otherwise.
func (o *PlatformValues) GetObjectSize1k() int64 {
	if o == nil || o.ObjectSize1k == nil {
		var ret int64
		return ret
	}
	return *o.ObjectSize1k
}

// GetObjectSize1kOk returns a tuple with the ObjectSize1k field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetObjectSize1kOk() (*int64, bool) {
	if o == nil || o.ObjectSize1k == nil {
		return nil, false
	}
	return o.ObjectSize1k, true
}

// HasObjectSize1k returns a boolean if a field has been set.
func (o *PlatformValues) HasObjectSize1k() bool {
	if o != nil && o.ObjectSize1k != nil {
		return true
	}

	return false
}

// SetObjectSize1k gets a reference to the given int64 and assigns it to the ObjectSize1k field.
func (o *PlatformValues) SetObjectSize1k(v int64) {
	o.ObjectSize1k = &v
}

// GetObjectSize10k returns the ObjectSize10k field value if set, zero value otherwise.
func (o *PlatformValues) GetObjectSize10k() int64 {
	if o == nil || o.ObjectSize10k == nil {
		var ret int64
		return ret
	}
	return *o.ObjectSize10k
}

// GetObjectSize10kOk returns a tuple with the ObjectSize10k field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetObjectSize10kOk() (*int64, bool) {
	if o == nil || o.ObjectSize10k == nil {
		return nil, false
	}
	return o.ObjectSize10k, true
}

// HasObjectSize10k returns a boolean if a field has been set.
func (o *PlatformValues) HasObjectSize10k() bool {
	if o != nil && o.ObjectSize10k != nil {
		return true
	}

	return false
}

// SetObjectSize10k gets a reference to the given int64 and assigns it to the ObjectSize10k field.
func (o *PlatformValues) SetObjectSize10k(v int64) {
	o.ObjectSize10k = &v
}

// GetObjectSize100k returns the ObjectSize100k field value if set, zero value otherwise.
func (o *PlatformValues) GetObjectSize100k() int64 {
	if o == nil || o.ObjectSize100k == nil {
		var ret int64
		return ret
	}
	return *o.ObjectSize100k
}

// GetObjectSize100kOk returns a tuple with the ObjectSize100k field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetObjectSize100kOk() (*int64, bool) {
	if o == nil || o.ObjectSize100k == nil {
		return nil, false
	}
	return o.ObjectSize100k, true
}

// HasObjectSize100k returns a boolean if a field has been set.
func (o *PlatformValues) HasObjectSize100k() bool {
	if o != nil && o.ObjectSize100k != nil {
		return true
	}

	return false
}

// SetObjectSize100k gets a reference to the given int64 and assigns it to the ObjectSize100k field.
func (o *PlatformValues) SetObjectSize100k(v int64) {
	o.ObjectSize100k = &v
}

// GetObjectSize1m returns the ObjectSize1m field value if set, zero value otherwise.
func (o *PlatformValues) GetObjectSize1m() int64 {
	if o == nil || o.ObjectSize1m == nil {
		var ret int64
		return ret
	}
	return *o.ObjectSize1m
}

// GetObjectSize1mOk returns a tuple with the ObjectSize1m field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetObjectSize1mOk() (*int64, bool) {
	if o == nil || o.ObjectSize1m == nil {
		return nil, false
	}
	return o.ObjectSize1m, true
}

// HasObjectSize1m returns a boolean if a field has been set.
func (o *PlatformValues) HasObjectSize1m() bool {
	if o != nil && o.ObjectSize1m != nil {
		return true
	}

	return false
}

// SetObjectSize1m gets a reference to the given int64 and assigns it to the ObjectSize1m field.
func (o *PlatformValues) SetObjectSize1m(v int64) {
	o.ObjectSize1m = &v
}

// GetObjectSize10m returns the ObjectSize10m field value if set, zero value otherwise.
func (o *PlatformValues) GetObjectSize10m() int64 {
	if o == nil || o.ObjectSize10m == nil {
		var ret int64
		return ret
	}
	return *o.ObjectSize10m
}

// GetObjectSize10mOk returns a tuple with the ObjectSize10m field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetObjectSize10mOk() (*int64, bool) {
	if o == nil || o.ObjectSize10m == nil {
		return nil, false
	}
	return o.ObjectSize10m, true
}

// HasObjectSize10m returns a boolean if a field has been set.
func (o *PlatformValues) HasObjectSize10m() bool {
	if o != nil && o.ObjectSize10m != nil {
		return true
	}

	return false
}

// SetObjectSize10m gets a reference to the given int64 and assigns it to the ObjectSize10m field.
func (o *PlatformValues) SetObjectSize10m(v int64) {
	o.ObjectSize10m = &v
}

// GetObjectSize100m returns the ObjectSize100m field value if set, zero value otherwise.
func (o *PlatformValues) GetObjectSize100m() int64 {
	if o == nil || o.ObjectSize100m == nil {
		var ret int64
		return ret
	}
	return *o.ObjectSize100m
}

// GetObjectSize100mOk returns a tuple with the ObjectSize100m field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetObjectSize100mOk() (*int64, bool) {
	if o == nil || o.ObjectSize100m == nil {
		return nil, false
	}
	return o.ObjectSize100m, true
}

// HasObjectSize100m returns a boolean if a field has been set.
func (o *PlatformValues) HasObjectSize100m() bool {
	if o != nil && o.ObjectSize100m != nil {
		return true
	}

	return false
}

// SetObjectSize100m gets a reference to the given int64 and assigns it to the ObjectSize100m field.
func (o *PlatformValues) SetObjectSize100m(v int64) {
	o.ObjectSize100m = &v
}

// GetObjectSize1g returns the ObjectSize1g field value if set, zero value otherwise.
func (o *PlatformValues) GetObjectSize1g() int64 {
	if o == nil || o.ObjectSize1g == nil {
		var ret int64
		return ret
	}
	return *o.ObjectSize1g
}

// GetObjectSize1gOk returns a tuple with the ObjectSize1g field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetObjectSize1gOk() (*int64, bool) {
	if o == nil || o.ObjectSize1g == nil {
		return nil, false
	}
	return o.ObjectSize1g, true
}

// HasObjectSize1g returns a boolean if a field has been set.
func (o *PlatformValues) HasObjectSize1g() bool {
	if o != nil && o.ObjectSize1g != nil {
		return true
	}

	return false
}

// SetObjectSize1g gets a reference to the given int64 and assigns it to the ObjectSize1g field.
func (o *PlatformValues) SetObjectSize1g(v int64) {
	o.ObjectSize1g = &v
}

// GetRecvSubTime returns the RecvSubTime field value if set, zero value otherwise.
func (o *PlatformValues) GetRecvSubTime() float32 {
	if o == nil || o.RecvSubTime == nil {
		var ret float32
		return ret
	}
	return *o.RecvSubTime
}

// GetRecvSubTimeOk returns a tuple with the RecvSubTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetRecvSubTimeOk() (*float32, bool) {
	if o == nil || o.RecvSubTime == nil {
		return nil, false
	}
	return o.RecvSubTime, true
}

// HasRecvSubTime returns a boolean if a field has been set.
func (o *PlatformValues) HasRecvSubTime() bool {
	if o != nil && o.RecvSubTime != nil {
		return true
	}

	return false
}

// SetRecvSubTime gets a reference to the given float32 and assigns it to the RecvSubTime field.
func (o *PlatformValues) SetRecvSubTime(v float32) {
	o.RecvSubTime = &v
}

// GetRecvSubCount returns the RecvSubCount field value if set, zero value otherwise.
func (o *PlatformValues) GetRecvSubCount() int64 {
	if o == nil || o.RecvSubCount == nil {
		var ret int64
		return ret
	}
	return *o.RecvSubCount
}

// GetRecvSubCountOk returns a tuple with the RecvSubCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetRecvSubCountOk() (*int64, bool) {
	if o == nil || o.RecvSubCount == nil {
		return nil, false
	}
	return o.RecvSubCount, true
}

// HasRecvSubCount returns a boolean if a field has been set.
func (o *PlatformValues) HasRecvSubCount() bool {
	if o != nil && o.RecvSubCount != nil {
		return true
	}

	return false
}

// SetRecvSubCount gets a reference to the given int64 and assigns it to the RecvSubCount field.
func (o *PlatformValues) SetRecvSubCount(v int64) {
	o.RecvSubCount = &v
}

// GetHashSubTime returns the HashSubTime field value if set, zero value otherwise.
func (o *PlatformValues) GetHashSubTime() float32 {
	if o == nil || o.HashSubTime == nil {
		var ret float32
		return ret
	}
	return *o.HashSubTime
}

// GetHashSubTimeOk returns a tuple with the HashSubTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetHashSubTimeOk() (*float32, bool) {
	if o == nil || o.HashSubTime == nil {
		return nil, false
	}
	return o.HashSubTime, true
}

// HasHashSubTime returns a boolean if a field has been set.
func (o *PlatformValues) HasHashSubTime() bool {
	if o != nil && o.HashSubTime != nil {
		return true
	}

	return false
}

// SetHashSubTime gets a reference to the given float32 and assigns it to the HashSubTime field.
func (o *PlatformValues) SetHashSubTime(v float32) {
	o.HashSubTime = &v
}

// GetHashSubCount returns the HashSubCount field value if set, zero value otherwise.
func (o *PlatformValues) GetHashSubCount() int64 {
	if o == nil || o.HashSubCount == nil {
		var ret int64
		return ret
	}
	return *o.HashSubCount
}

// GetHashSubCountOk returns a tuple with the HashSubCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetHashSubCountOk() (*int64, bool) {
	if o == nil || o.HashSubCount == nil {
		return nil, false
	}
	return o.HashSubCount, true
}

// HasHashSubCount returns a boolean if a field has been set.
func (o *PlatformValues) HasHashSubCount() bool {
	if o != nil && o.HashSubCount != nil {
		return true
	}

	return false
}

// SetHashSubCount gets a reference to the given int64 and assigns it to the HashSubCount field.
func (o *PlatformValues) SetHashSubCount(v int64) {
	o.HashSubCount = &v
}

// GetMissSubTime returns the MissSubTime field value if set, zero value otherwise.
func (o *PlatformValues) GetMissSubTime() float32 {
	if o == nil || o.MissSubTime == nil {
		var ret float32
		return ret
	}
	return *o.MissSubTime
}

// GetMissSubTimeOk returns a tuple with the MissSubTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetMissSubTimeOk() (*float32, bool) {
	if o == nil || o.MissSubTime == nil {
		return nil, false
	}
	return o.MissSubTime, true
}

// HasMissSubTime returns a boolean if a field has been set.
func (o *PlatformValues) HasMissSubTime() bool {
	if o != nil && o.MissSubTime != nil {
		return true
	}

	return false
}

// SetMissSubTime gets a reference to the given float32 and assigns it to the MissSubTime field.
func (o *PlatformValues) SetMissSubTime(v float32) {
	o.MissSubTime = &v
}

// GetMissSubCount returns the MissSubCount field value if set, zero value otherwise.
func (o *PlatformValues) GetMissSubCount() int64 {
	if o == nil || o.MissSubCount == nil {
		var ret int64
		return ret
	}
	return *o.MissSubCount
}

// GetMissSubCountOk returns a tuple with the MissSubCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetMissSubCountOk() (*int64, bool) {
	if o == nil || o.MissSubCount == nil {
		return nil, false
	}
	return o.MissSubCount, true
}

// HasMissSubCount returns a boolean if a field has been set.
func (o *PlatformValues) HasMissSubCount() bool {
	if o != nil && o.MissSubCount != nil {
		return true
	}

	return false
}

// SetMissSubCount gets a reference to the given int64 and assigns it to the MissSubCount field.
func (o *PlatformValues) SetMissSubCount(v int64) {
	o.MissSubCount = &v
}

// GetFetchSubTime returns the FetchSubTime field value if set, zero value otherwise.
func (o *PlatformValues) GetFetchSubTime() float32 {
	if o == nil || o.FetchSubTime == nil {
		var ret float32
		return ret
	}
	return *o.FetchSubTime
}

// GetFetchSubTimeOk returns a tuple with the FetchSubTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetFetchSubTimeOk() (*float32, bool) {
	if o == nil || o.FetchSubTime == nil {
		return nil, false
	}
	return o.FetchSubTime, true
}

// HasFetchSubTime returns a boolean if a field has been set.
func (o *PlatformValues) HasFetchSubTime() bool {
	if o != nil && o.FetchSubTime != nil {
		return true
	}

	return false
}

// SetFetchSubTime gets a reference to the given float32 and assigns it to the FetchSubTime field.
func (o *PlatformValues) SetFetchSubTime(v float32) {
	o.FetchSubTime = &v
}

// GetFetchSubCount returns the FetchSubCount field value if set, zero value otherwise.
func (o *PlatformValues) GetFetchSubCount() int64 {
	if o == nil || o.FetchSubCount == nil {
		var ret int64
		return ret
	}
	return *o.FetchSubCount
}

// GetFetchSubCountOk returns a tuple with the FetchSubCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetFetchSubCountOk() (*int64, bool) {
	if o == nil || o.FetchSubCount == nil {
		return nil, false
	}
	return o.FetchSubCount, true
}

// HasFetchSubCount returns a boolean if a field has been set.
func (o *PlatformValues) HasFetchSubCount() bool {
	if o != nil && o.FetchSubCount != nil {
		return true
	}

	return false
}

// SetFetchSubCount gets a reference to the given int64 and assigns it to the FetchSubCount field.
func (o *PlatformValues) SetFetchSubCount(v int64) {
	o.FetchSubCount = &v
}

// GetPassSubTime returns the PassSubTime field value if set, zero value otherwise.
func (o *PlatformValues) GetPassSubTime() float32 {
	if o == nil || o.PassSubTime == nil {
		var ret float32
		return ret
	}
	return *o.PassSubTime
}

// GetPassSubTimeOk returns a tuple with the PassSubTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetPassSubTimeOk() (*float32, bool) {
	if o == nil || o.PassSubTime == nil {
		return nil, false
	}
	return o.PassSubTime, true
}

// HasPassSubTime returns a boolean if a field has been set.
func (o *PlatformValues) HasPassSubTime() bool {
	if o != nil && o.PassSubTime != nil {
		return true
	}

	return false
}

// SetPassSubTime gets a reference to the given float32 and assigns it to the PassSubTime field.
func (o *PlatformValues) SetPassSubTime(v float32) {
	o.PassSubTime = &v
}

// GetPassSubCount returns the PassSubCount field value if set, zero value otherwise.
func (o *PlatformValues) GetPassSubCount() int64 {
	if o == nil || o.PassSubCount == nil {
		var ret int64
		return ret
	}
	return *o.PassSubCount
}

// GetPassSubCountOk returns a tuple with the PassSubCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetPassSubCountOk() (*int64, bool) {
	if o == nil || o.PassSubCount == nil {
		return nil, false
	}
	return o.PassSubCount, true
}

// HasPassSubCount returns a boolean if a field has been set.
func (o *PlatformValues) HasPassSubCount() bool {
	if o != nil && o.PassSubCount != nil {
		return true
	}

	return false
}

// SetPassSubCount gets a reference to the given int64 and assigns it to the PassSubCount field.
func (o *PlatformValues) SetPassSubCount(v int64) {
	o.PassSubCount = &v
}

// GetPipeSubTime returns the PipeSubTime field value if set, zero value otherwise.
func (o *PlatformValues) GetPipeSubTime() float32 {
	if o == nil || o.PipeSubTime == nil {
		var ret float32
		return ret
	}
	return *o.PipeSubTime
}

// GetPipeSubTimeOk returns a tuple with the PipeSubTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetPipeSubTimeOk() (*float32, bool) {
	if o == nil || o.PipeSubTime == nil {
		return nil, false
	}
	return o.PipeSubTime, true
}

// HasPipeSubTime returns a boolean if a field has been set.
func (o *PlatformValues) HasPipeSubTime() bool {
	if o != nil && o.PipeSubTime != nil {
		return true
	}

	return false
}

// SetPipeSubTime gets a reference to the given float32 and assigns it to the PipeSubTime field.
func (o *PlatformValues) SetPipeSubTime(v float32) {
	o.PipeSubTime = &v
}

// GetPipeSubCount returns the PipeSubCount field value if set, zero value otherwise.
func (o *PlatformValues) GetPipeSubCount() int64 {
	if o == nil || o.PipeSubCount == nil {
		var ret int64
		return ret
	}
	return *o.PipeSubCount
}

// GetPipeSubCountOk returns a tuple with the PipeSubCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetPipeSubCountOk() (*int64, bool) {
	if o == nil || o.PipeSubCount == nil {
		return nil, false
	}
	return o.PipeSubCount, true
}

// HasPipeSubCount returns a boolean if a field has been set.
func (o *PlatformValues) HasPipeSubCount() bool {
	if o != nil && o.PipeSubCount != nil {
		return true
	}

	return false
}

// SetPipeSubCount gets a reference to the given int64 and assigns it to the PipeSubCount field.
func (o *PlatformValues) SetPipeSubCount(v int64) {
	o.PipeSubCount = &v
}

// GetDeliverSubTime returns the DeliverSubTime field value if set, zero value otherwise.
func (o *PlatformValues) GetDeliverSubTime() float32 {
	if o == nil || o.DeliverSubTime == nil {
		var ret float32
		return ret
	}
	return *o.DeliverSubTime
}

// GetDeliverSubTimeOk returns a tuple with the DeliverSubTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetDeliverSubTimeOk() (*float32, bool) {
	if o == nil || o.DeliverSubTime == nil {
		return nil, false
	}
	return o.DeliverSubTime, true
}

// HasDeliverSubTime returns a boolean if a field has been set.
func (o *PlatformValues) HasDeliverSubTime() bool {
	if o != nil && o.DeliverSubTime != nil {
		return true
	}

	return false
}

// SetDeliverSubTime gets a reference to the given float32 and assigns it to the DeliverSubTime field.
func (o *PlatformValues) SetDeliverSubTime(v float32) {
	o.DeliverSubTime = &v
}

// GetDeliverSubCount returns the DeliverSubCount field value if set, zero value otherwise.
func (o *PlatformValues) GetDeliverSubCount() int64 {
	if o == nil || o.DeliverSubCount == nil {
		var ret int64
		return ret
	}
	return *o.DeliverSubCount
}

// GetDeliverSubCountOk returns a tuple with the DeliverSubCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetDeliverSubCountOk() (*int64, bool) {
	if o == nil || o.DeliverSubCount == nil {
		return nil, false
	}
	return o.DeliverSubCount, true
}

// HasDeliverSubCount returns a boolean if a field has been set.
func (o *PlatformValues) HasDeliverSubCount() bool {
	if o != nil && o.DeliverSubCount != nil {
		return true
	}

	return false
}

// SetDeliverSubCount gets a reference to the given int64 and assigns it to the DeliverSubCount field.
func (o *PlatformValues) SetDeliverSubCount(v int64) {
	o.DeliverSubCount = &v
}

// GetErrorSubTime returns the ErrorSubTime field value if set, zero value otherwise.
func (o *PlatformValues) GetErrorSubTime() float32 {
	if o == nil || o.ErrorSubTime == nil {
		var ret float32
		return ret
	}
	return *o.ErrorSubTime
}

// GetErrorSubTimeOk returns a tuple with the ErrorSubTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetErrorSubTimeOk() (*float32, bool) {
	if o == nil || o.ErrorSubTime == nil {
		return nil, false
	}
	return o.ErrorSubTime, true
}

// HasErrorSubTime returns a boolean if a field has been set.
func (o *PlatformValues) HasErrorSubTime() bool {
	if o != nil && o.ErrorSubTime != nil {
		return true
	}

	return false
}

// SetErrorSubTime gets a reference to the given float32 and assigns it to the ErrorSubTime field.
func (o *PlatformValues) SetErrorSubTime(v float32) {
	o.ErrorSubTime = &v
}

// GetErrorSubCount returns the ErrorSubCount field value if set, zero value otherwise.
func (o *PlatformValues) GetErrorSubCount() int64 {
	if o == nil || o.ErrorSubCount == nil {
		var ret int64
		return ret
	}
	return *o.ErrorSubCount
}

// GetErrorSubCountOk returns a tuple with the ErrorSubCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetErrorSubCountOk() (*int64, bool) {
	if o == nil || o.ErrorSubCount == nil {
		return nil, false
	}
	return o.ErrorSubCount, true
}

// HasErrorSubCount returns a boolean if a field has been set.
func (o *PlatformValues) HasErrorSubCount() bool {
	if o != nil && o.ErrorSubCount != nil {
		return true
	}

	return false
}

// SetErrorSubCount gets a reference to the given int64 and assigns it to the ErrorSubCount field.
func (o *PlatformValues) SetErrorSubCount(v int64) {
	o.ErrorSubCount = &v
}

// GetHitSubTime returns the HitSubTime field value if set, zero value otherwise.
func (o *PlatformValues) GetHitSubTime() float32 {
	if o == nil || o.HitSubTime == nil {
		var ret float32
		return ret
	}
	return *o.HitSubTime
}

// GetHitSubTimeOk returns a tuple with the HitSubTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetHitSubTimeOk() (*float32, bool) {
	if o == nil || o.HitSubTime == nil {
		return nil, false
	}
	return o.HitSubTime, true
}

// HasHitSubTime returns a boolean if a field has been set.
func (o *PlatformValues) HasHitSubTime() bool {
	if o != nil && o.HitSubTime != nil {
		return true
	}

	return false
}

// SetHitSubTime gets a reference to the given float32 and assigns it to the HitSubTime field.
func (o *PlatformValues) SetHitSubTime(v float32) {
	o.HitSubTime = &v
}

// GetHitSubCount returns the HitSubCount field value if set, zero value otherwise.
func (o *PlatformValues) GetHitSubCount() int64 {
	if o == nil || o.HitSubCount == nil {
		var ret int64
		return ret
	}
	return *o.HitSubCount
}

// GetHitSubCountOk returns a tuple with the HitSubCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetHitSubCountOk() (*int64, bool) {
	if o == nil || o.HitSubCount == nil {
		return nil, false
	}
	return o.HitSubCount, true
}

// HasHitSubCount returns a boolean if a field has been set.
func (o *PlatformValues) HasHitSubCount() bool {
	if o != nil && o.HitSubCount != nil {
		return true
	}

	return false
}

// SetHitSubCount gets a reference to the given int64 and assigns it to the HitSubCount field.
func (o *PlatformValues) SetHitSubCount(v int64) {
	o.HitSubCount = &v
}

// GetPrehashSubTime returns the PrehashSubTime field value if set, zero value otherwise.
func (o *PlatformValues) GetPrehashSubTime() float32 {
	if o == nil || o.PrehashSubTime == nil {
		var ret float32
		return ret
	}
	return *o.PrehashSubTime
}

// GetPrehashSubTimeOk returns a tuple with the PrehashSubTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetPrehashSubTimeOk() (*float32, bool) {
	if o == nil || o.PrehashSubTime == nil {
		return nil, false
	}
	return o.PrehashSubTime, true
}

// HasPrehashSubTime returns a boolean if a field has been set.
func (o *PlatformValues) HasPrehashSubTime() bool {
	if o != nil && o.PrehashSubTime != nil {
		return true
	}

	return false
}

// SetPrehashSubTime gets a reference to the given float32 and assigns it to the PrehashSubTime field.
func (o *PlatformValues) SetPrehashSubTime(v float32) {
	o.PrehashSubTime = &v
}

// GetPrehashSubCount returns the PrehashSubCount field value if set, zero value otherwise.
func (o *PlatformValues) GetPrehashSubCount() int64 {
	if o == nil || o.PrehashSubCount == nil {
		var ret int64
		return ret
	}
	return *o.PrehashSubCount
}

// GetPrehashSubCountOk returns a tuple with the PrehashSubCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetPrehashSubCountOk() (*int64, bool) {
	if o == nil || o.PrehashSubCount == nil {
		return nil, false
	}
	return o.PrehashSubCount, true
}

// HasPrehashSubCount returns a boolean if a field has been set.
func (o *PlatformValues) HasPrehashSubCount() bool {
	if o != nil && o.PrehashSubCount != nil {
		return true
	}

	return false
}

// SetPrehashSubCount gets a reference to the given int64 and assigns it to the PrehashSubCount field.
func (o *PlatformValues) SetPrehashSubCount(v int64) {
	o.PrehashSubCount = &v
}

// GetPredeliverSubTime returns the PredeliverSubTime field value if set, zero value otherwise.
func (o *PlatformValues) GetPredeliverSubTime() float32 {
	if o == nil || o.PredeliverSubTime == nil {
		var ret float32
		return ret
	}
	return *o.PredeliverSubTime
}

// GetPredeliverSubTimeOk returns a tuple with the PredeliverSubTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetPredeliverSubTimeOk() (*float32, bool) {
	if o == nil || o.PredeliverSubTime == nil {
		return nil, false
	}
	return o.PredeliverSubTime, true
}

// HasPredeliverSubTime returns a boolean if a field has been set.
func (o *PlatformValues) HasPredeliverSubTime() bool {
	if o != nil && o.PredeliverSubTime != nil {
		return true
	}

	return false
}

// SetPredeliverSubTime gets a reference to the given float32 and assigns it to the PredeliverSubTime field.
func (o *PlatformValues) SetPredeliverSubTime(v float32) {
	o.PredeliverSubTime = &v
}

// GetPredeliverSubCount returns the PredeliverSubCount field value if set, zero value otherwise.
func (o *PlatformValues) GetPredeliverSubCount() int64 {
	if o == nil || o.PredeliverSubCount == nil {
		var ret int64
		return ret
	}
	return *o.PredeliverSubCount
}

// GetPredeliverSubCountOk returns a tuple with the PredeliverSubCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetPredeliverSubCountOk() (*int64, bool) {
	if o == nil || o.PredeliverSubCount == nil {
		return nil, false
	}
	return o.PredeliverSubCount, true
}

// HasPredeliverSubCount returns a boolean if a field has been set.
func (o *PlatformValues) HasPredeliverSubCount() bool {
	if o != nil && o.PredeliverSubCount != nil {
		return true
	}

	return false
}

// SetPredeliverSubCount gets a reference to the given int64 and assigns it to the PredeliverSubCount field.
func (o *PlatformValues) SetPredeliverSubCount(v int64) {
	o.PredeliverSubCount = &v
}

// GetTLSHandshakeSentBytes returns the TLSHandshakeSentBytes field value if set, zero value otherwise.
func (o *PlatformValues) GetTLSHandshakeSentBytes() int64 {
	if o == nil || o.TLSHandshakeSentBytes == nil {
		var ret int64
		return ret
	}
	return *o.TLSHandshakeSentBytes
}

// GetTLSHandshakeSentBytesOk returns a tuple with the TLSHandshakeSentBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetTLSHandshakeSentBytesOk() (*int64, bool) {
	if o == nil || o.TLSHandshakeSentBytes == nil {
		return nil, false
	}
	return o.TLSHandshakeSentBytes, true
}

// HasTLSHandshakeSentBytes returns a boolean if a field has been set.
func (o *PlatformValues) HasTLSHandshakeSentBytes() bool {
	if o != nil && o.TLSHandshakeSentBytes != nil {
		return true
	}

	return false
}

// SetTLSHandshakeSentBytes gets a reference to the given int64 and assigns it to the TLSHandshakeSentBytes field.
func (o *PlatformValues) SetTLSHandshakeSentBytes(v int64) {
	o.TLSHandshakeSentBytes = &v
}

// GetHitRespBodyBytes returns the HitRespBodyBytes field value if set, zero value otherwise.
func (o *PlatformValues) GetHitRespBodyBytes() int64 {
	if o == nil || o.HitRespBodyBytes == nil {
		var ret int64
		return ret
	}
	return *o.HitRespBodyBytes
}

// GetHitRespBodyBytesOk returns a tuple with the HitRespBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetHitRespBodyBytesOk() (*int64, bool) {
	if o == nil || o.HitRespBodyBytes == nil {
		return nil, false
	}
	return o.HitRespBodyBytes, true
}

// HasHitRespBodyBytes returns a boolean if a field has been set.
func (o *PlatformValues) HasHitRespBodyBytes() bool {
	if o != nil && o.HitRespBodyBytes != nil {
		return true
	}

	return false
}

// SetHitRespBodyBytes gets a reference to the given int64 and assigns it to the HitRespBodyBytes field.
func (o *PlatformValues) SetHitRespBodyBytes(v int64) {
	o.HitRespBodyBytes = &v
}

// GetMissRespBodyBytes returns the MissRespBodyBytes field value if set, zero value otherwise.
func (o *PlatformValues) GetMissRespBodyBytes() int64 {
	if o == nil || o.MissRespBodyBytes == nil {
		var ret int64
		return ret
	}
	return *o.MissRespBodyBytes
}

// GetMissRespBodyBytesOk returns a tuple with the MissRespBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetMissRespBodyBytesOk() (*int64, bool) {
	if o == nil || o.MissRespBodyBytes == nil {
		return nil, false
	}
	return o.MissRespBodyBytes, true
}

// HasMissRespBodyBytes returns a boolean if a field has been set.
func (o *PlatformValues) HasMissRespBodyBytes() bool {
	if o != nil && o.MissRespBodyBytes != nil {
		return true
	}

	return false
}

// SetMissRespBodyBytes gets a reference to the given int64 and assigns it to the MissRespBodyBytes field.
func (o *PlatformValues) SetMissRespBodyBytes(v int64) {
	o.MissRespBodyBytes = &v
}

// GetPassRespBodyBytes returns the PassRespBodyBytes field value if set, zero value otherwise.
func (o *PlatformValues) GetPassRespBodyBytes() int64 {
	if o == nil || o.PassRespBodyBytes == nil {
		var ret int64
		return ret
	}
	return *o.PassRespBodyBytes
}

// GetPassRespBodyBytesOk returns a tuple with the PassRespBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetPassRespBodyBytesOk() (*int64, bool) {
	if o == nil || o.PassRespBodyBytes == nil {
		return nil, false
	}
	return o.PassRespBodyBytes, true
}

// HasPassRespBodyBytes returns a boolean if a field has been set.
func (o *PlatformValues) HasPassRespBodyBytes() bool {
	if o != nil && o.PassRespBodyBytes != nil {
		return true
	}

	return false
}

// SetPassRespBodyBytes gets a reference to the given int64 and assigns it to the PassRespBodyBytes field.
func (o *PlatformValues) SetPassRespBodyBytes(v int64) {
	o.PassRespBodyBytes = &v
}

// GetSegblockOriginFetches returns the SegblockOriginFetches field value if set, zero value otherwise.
func (o *PlatformValues) GetSegblockOriginFetches() int64 {
	if o == nil || o.SegblockOriginFetches == nil {
		var ret int64
		return ret
	}
	return *o.SegblockOriginFetches
}

// GetSegblockOriginFetchesOk returns a tuple with the SegblockOriginFetches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetSegblockOriginFetchesOk() (*int64, bool) {
	if o == nil || o.SegblockOriginFetches == nil {
		return nil, false
	}
	return o.SegblockOriginFetches, true
}

// HasSegblockOriginFetches returns a boolean if a field has been set.
func (o *PlatformValues) HasSegblockOriginFetches() bool {
	if o != nil && o.SegblockOriginFetches != nil {
		return true
	}

	return false
}

// SetSegblockOriginFetches gets a reference to the given int64 and assigns it to the SegblockOriginFetches field.
func (o *PlatformValues) SetSegblockOriginFetches(v int64) {
	o.SegblockOriginFetches = &v
}

// GetSegblockShieldFetches returns the SegblockShieldFetches field value if set, zero value otherwise.
func (o *PlatformValues) GetSegblockShieldFetches() int64 {
	if o == nil || o.SegblockShieldFetches == nil {
		var ret int64
		return ret
	}
	return *o.SegblockShieldFetches
}

// GetSegblockShieldFetchesOk returns a tuple with the SegblockShieldFetches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetSegblockShieldFetchesOk() (*int64, bool) {
	if o == nil || o.SegblockShieldFetches == nil {
		return nil, false
	}
	return o.SegblockShieldFetches, true
}

// HasSegblockShieldFetches returns a boolean if a field has been set.
func (o *PlatformValues) HasSegblockShieldFetches() bool {
	if o != nil && o.SegblockShieldFetches != nil {
		return true
	}

	return false
}

// SetSegblockShieldFetches gets a reference to the given int64 and assigns it to the SegblockShieldFetches field.
func (o *PlatformValues) SetSegblockShieldFetches(v int64) {
	o.SegblockShieldFetches = &v
}

// GetComputeRequests returns the ComputeRequests field value if set, zero value otherwise.
func (o *PlatformValues) GetComputeRequests() int64 {
	if o == nil || o.ComputeRequests == nil {
		var ret int64
		return ret
	}
	return *o.ComputeRequests
}

// GetComputeRequestsOk returns a tuple with the ComputeRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetComputeRequestsOk() (*int64, bool) {
	if o == nil || o.ComputeRequests == nil {
		return nil, false
	}
	return o.ComputeRequests, true
}

// HasComputeRequests returns a boolean if a field has been set.
func (o *PlatformValues) HasComputeRequests() bool {
	if o != nil && o.ComputeRequests != nil {
		return true
	}

	return false
}

// SetComputeRequests gets a reference to the given int64 and assigns it to the ComputeRequests field.
func (o *PlatformValues) SetComputeRequests(v int64) {
	o.ComputeRequests = &v
}

// GetComputeRequestTimeMs returns the ComputeRequestTimeMs field value if set, zero value otherwise.
func (o *PlatformValues) GetComputeRequestTimeMs() float32 {
	if o == nil || o.ComputeRequestTimeMs == nil {
		var ret float32
		return ret
	}
	return *o.ComputeRequestTimeMs
}

// GetComputeRequestTimeMsOk returns a tuple with the ComputeRequestTimeMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetComputeRequestTimeMsOk() (*float32, bool) {
	if o == nil || o.ComputeRequestTimeMs == nil {
		return nil, false
	}
	return o.ComputeRequestTimeMs, true
}

// HasComputeRequestTimeMs returns a boolean if a field has been set.
func (o *PlatformValues) HasComputeRequestTimeMs() bool {
	if o != nil && o.ComputeRequestTimeMs != nil {
		return true
	}

	return false
}

// SetComputeRequestTimeMs gets a reference to the given float32 and assigns it to the ComputeRequestTimeMs field.
func (o *PlatformValues) SetComputeRequestTimeMs(v float32) {
	o.ComputeRequestTimeMs = &v
}

// GetComputeRequestTimeBilledMs returns the ComputeRequestTimeBilledMs field value if set, zero value otherwise.
func (o *PlatformValues) GetComputeRequestTimeBilledMs() float32 {
	if o == nil || o.ComputeRequestTimeBilledMs == nil {
		var ret float32
		return ret
	}
	return *o.ComputeRequestTimeBilledMs
}

// GetComputeRequestTimeBilledMsOk returns a tuple with the ComputeRequestTimeBilledMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetComputeRequestTimeBilledMsOk() (*float32, bool) {
	if o == nil || o.ComputeRequestTimeBilledMs == nil {
		return nil, false
	}
	return o.ComputeRequestTimeBilledMs, true
}

// HasComputeRequestTimeBilledMs returns a boolean if a field has been set.
func (o *PlatformValues) HasComputeRequestTimeBilledMs() bool {
	if o != nil && o.ComputeRequestTimeBilledMs != nil {
		return true
	}

	return false
}

// SetComputeRequestTimeBilledMs gets a reference to the given float32 and assigns it to the ComputeRequestTimeBilledMs field.
func (o *PlatformValues) SetComputeRequestTimeBilledMs(v float32) {
	o.ComputeRequestTimeBilledMs = &v
}

// GetComputeRAMUsed returns the ComputeRAMUsed field value if set, zero value otherwise.
func (o *PlatformValues) GetComputeRAMUsed() int64 {
	if o == nil || o.ComputeRAMUsed == nil {
		var ret int64
		return ret
	}
	return *o.ComputeRAMUsed
}

// GetComputeRAMUsedOk returns a tuple with the ComputeRAMUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetComputeRAMUsedOk() (*int64, bool) {
	if o == nil || o.ComputeRAMUsed == nil {
		return nil, false
	}
	return o.ComputeRAMUsed, true
}

// HasComputeRAMUsed returns a boolean if a field has been set.
func (o *PlatformValues) HasComputeRAMUsed() bool {
	if o != nil && o.ComputeRAMUsed != nil {
		return true
	}

	return false
}

// SetComputeRAMUsed gets a reference to the given int64 and assigns it to the ComputeRAMUsed field.
func (o *PlatformValues) SetComputeRAMUsed(v int64) {
	o.ComputeRAMUsed = &v
}

// GetComputeExecutionTimeMs returns the ComputeExecutionTimeMs field value if set, zero value otherwise.
func (o *PlatformValues) GetComputeExecutionTimeMs() float32 {
	if o == nil || o.ComputeExecutionTimeMs == nil {
		var ret float32
		return ret
	}
	return *o.ComputeExecutionTimeMs
}

// GetComputeExecutionTimeMsOk returns a tuple with the ComputeExecutionTimeMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetComputeExecutionTimeMsOk() (*float32, bool) {
	if o == nil || o.ComputeExecutionTimeMs == nil {
		return nil, false
	}
	return o.ComputeExecutionTimeMs, true
}

// HasComputeExecutionTimeMs returns a boolean if a field has been set.
func (o *PlatformValues) HasComputeExecutionTimeMs() bool {
	if o != nil && o.ComputeExecutionTimeMs != nil {
		return true
	}

	return false
}

// SetComputeExecutionTimeMs gets a reference to the given float32 and assigns it to the ComputeExecutionTimeMs field.
func (o *PlatformValues) SetComputeExecutionTimeMs(v float32) {
	o.ComputeExecutionTimeMs = &v
}

// GetComputeReqHeaderBytes returns the ComputeReqHeaderBytes field value if set, zero value otherwise.
func (o *PlatformValues) GetComputeReqHeaderBytes() int64 {
	if o == nil || o.ComputeReqHeaderBytes == nil {
		var ret int64
		return ret
	}
	return *o.ComputeReqHeaderBytes
}

// GetComputeReqHeaderBytesOk returns a tuple with the ComputeReqHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetComputeReqHeaderBytesOk() (*int64, bool) {
	if o == nil || o.ComputeReqHeaderBytes == nil {
		return nil, false
	}
	return o.ComputeReqHeaderBytes, true
}

// HasComputeReqHeaderBytes returns a boolean if a field has been set.
func (o *PlatformValues) HasComputeReqHeaderBytes() bool {
	if o != nil && o.ComputeReqHeaderBytes != nil {
		return true
	}

	return false
}

// SetComputeReqHeaderBytes gets a reference to the given int64 and assigns it to the ComputeReqHeaderBytes field.
func (o *PlatformValues) SetComputeReqHeaderBytes(v int64) {
	o.ComputeReqHeaderBytes = &v
}

// GetComputeReqBodyBytes returns the ComputeReqBodyBytes field value if set, zero value otherwise.
func (o *PlatformValues) GetComputeReqBodyBytes() int64 {
	if o == nil || o.ComputeReqBodyBytes == nil {
		var ret int64
		return ret
	}
	return *o.ComputeReqBodyBytes
}

// GetComputeReqBodyBytesOk returns a tuple with the ComputeReqBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetComputeReqBodyBytesOk() (*int64, bool) {
	if o == nil || o.ComputeReqBodyBytes == nil {
		return nil, false
	}
	return o.ComputeReqBodyBytes, true
}

// HasComputeReqBodyBytes returns a boolean if a field has been set.
func (o *PlatformValues) HasComputeReqBodyBytes() bool {
	if o != nil && o.ComputeReqBodyBytes != nil {
		return true
	}

	return false
}

// SetComputeReqBodyBytes gets a reference to the given int64 and assigns it to the ComputeReqBodyBytes field.
func (o *PlatformValues) SetComputeReqBodyBytes(v int64) {
	o.ComputeReqBodyBytes = &v
}

// GetComputeRespHeaderBytes returns the ComputeRespHeaderBytes field value if set, zero value otherwise.
func (o *PlatformValues) GetComputeRespHeaderBytes() int64 {
	if o == nil || o.ComputeRespHeaderBytes == nil {
		var ret int64
		return ret
	}
	return *o.ComputeRespHeaderBytes
}

// GetComputeRespHeaderBytesOk returns a tuple with the ComputeRespHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetComputeRespHeaderBytesOk() (*int64, bool) {
	if o == nil || o.ComputeRespHeaderBytes == nil {
		return nil, false
	}
	return o.ComputeRespHeaderBytes, true
}

// HasComputeRespHeaderBytes returns a boolean if a field has been set.
func (o *PlatformValues) HasComputeRespHeaderBytes() bool {
	if o != nil && o.ComputeRespHeaderBytes != nil {
		return true
	}

	return false
}

// SetComputeRespHeaderBytes gets a reference to the given int64 and assigns it to the ComputeRespHeaderBytes field.
func (o *PlatformValues) SetComputeRespHeaderBytes(v int64) {
	o.ComputeRespHeaderBytes = &v
}

// GetComputeRespBodyBytes returns the ComputeRespBodyBytes field value if set, zero value otherwise.
func (o *PlatformValues) GetComputeRespBodyBytes() int64 {
	if o == nil || o.ComputeRespBodyBytes == nil {
		var ret int64
		return ret
	}
	return *o.ComputeRespBodyBytes
}

// GetComputeRespBodyBytesOk returns a tuple with the ComputeRespBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetComputeRespBodyBytesOk() (*int64, bool) {
	if o == nil || o.ComputeRespBodyBytes == nil {
		return nil, false
	}
	return o.ComputeRespBodyBytes, true
}

// HasComputeRespBodyBytes returns a boolean if a field has been set.
func (o *PlatformValues) HasComputeRespBodyBytes() bool {
	if o != nil && o.ComputeRespBodyBytes != nil {
		return true
	}

	return false
}

// SetComputeRespBodyBytes gets a reference to the given int64 and assigns it to the ComputeRespBodyBytes field.
func (o *PlatformValues) SetComputeRespBodyBytes(v int64) {
	o.ComputeRespBodyBytes = &v
}

// GetComputeRespStatus1xx returns the ComputeRespStatus1xx field value if set, zero value otherwise.
func (o *PlatformValues) GetComputeRespStatus1xx() int64 {
	if o == nil || o.ComputeRespStatus1xx == nil {
		var ret int64
		return ret
	}
	return *o.ComputeRespStatus1xx
}

// GetComputeRespStatus1xxOk returns a tuple with the ComputeRespStatus1xx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetComputeRespStatus1xxOk() (*int64, bool) {
	if o == nil || o.ComputeRespStatus1xx == nil {
		return nil, false
	}
	return o.ComputeRespStatus1xx, true
}

// HasComputeRespStatus1xx returns a boolean if a field has been set.
func (o *PlatformValues) HasComputeRespStatus1xx() bool {
	if o != nil && o.ComputeRespStatus1xx != nil {
		return true
	}

	return false
}

// SetComputeRespStatus1xx gets a reference to the given int64 and assigns it to the ComputeRespStatus1xx field.
func (o *PlatformValues) SetComputeRespStatus1xx(v int64) {
	o.ComputeRespStatus1xx = &v
}

// GetComputeRespStatus2xx returns the ComputeRespStatus2xx field value if set, zero value otherwise.
func (o *PlatformValues) GetComputeRespStatus2xx() int64 {
	if o == nil || o.ComputeRespStatus2xx == nil {
		var ret int64
		return ret
	}
	return *o.ComputeRespStatus2xx
}

// GetComputeRespStatus2xxOk returns a tuple with the ComputeRespStatus2xx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetComputeRespStatus2xxOk() (*int64, bool) {
	if o == nil || o.ComputeRespStatus2xx == nil {
		return nil, false
	}
	return o.ComputeRespStatus2xx, true
}

// HasComputeRespStatus2xx returns a boolean if a field has been set.
func (o *PlatformValues) HasComputeRespStatus2xx() bool {
	if o != nil && o.ComputeRespStatus2xx != nil {
		return true
	}

	return false
}

// SetComputeRespStatus2xx gets a reference to the given int64 and assigns it to the ComputeRespStatus2xx field.
func (o *PlatformValues) SetComputeRespStatus2xx(v int64) {
	o.ComputeRespStatus2xx = &v
}

// GetComputeRespStatus3xx returns the ComputeRespStatus3xx field value if set, zero value otherwise.
func (o *PlatformValues) GetComputeRespStatus3xx() int64 {
	if o == nil || o.ComputeRespStatus3xx == nil {
		var ret int64
		return ret
	}
	return *o.ComputeRespStatus3xx
}

// GetComputeRespStatus3xxOk returns a tuple with the ComputeRespStatus3xx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetComputeRespStatus3xxOk() (*int64, bool) {
	if o == nil || o.ComputeRespStatus3xx == nil {
		return nil, false
	}
	return o.ComputeRespStatus3xx, true
}

// HasComputeRespStatus3xx returns a boolean if a field has been set.
func (o *PlatformValues) HasComputeRespStatus3xx() bool {
	if o != nil && o.ComputeRespStatus3xx != nil {
		return true
	}

	return false
}

// SetComputeRespStatus3xx gets a reference to the given int64 and assigns it to the ComputeRespStatus3xx field.
func (o *PlatformValues) SetComputeRespStatus3xx(v int64) {
	o.ComputeRespStatus3xx = &v
}

// GetComputeRespStatus4xx returns the ComputeRespStatus4xx field value if set, zero value otherwise.
func (o *PlatformValues) GetComputeRespStatus4xx() int64 {
	if o == nil || o.ComputeRespStatus4xx == nil {
		var ret int64
		return ret
	}
	return *o.ComputeRespStatus4xx
}

// GetComputeRespStatus4xxOk returns a tuple with the ComputeRespStatus4xx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetComputeRespStatus4xxOk() (*int64, bool) {
	if o == nil || o.ComputeRespStatus4xx == nil {
		return nil, false
	}
	return o.ComputeRespStatus4xx, true
}

// HasComputeRespStatus4xx returns a boolean if a field has been set.
func (o *PlatformValues) HasComputeRespStatus4xx() bool {
	if o != nil && o.ComputeRespStatus4xx != nil {
		return true
	}

	return false
}

// SetComputeRespStatus4xx gets a reference to the given int64 and assigns it to the ComputeRespStatus4xx field.
func (o *PlatformValues) SetComputeRespStatus4xx(v int64) {
	o.ComputeRespStatus4xx = &v
}

// GetComputeRespStatus5xx returns the ComputeRespStatus5xx field value if set, zero value otherwise.
func (o *PlatformValues) GetComputeRespStatus5xx() int64 {
	if o == nil || o.ComputeRespStatus5xx == nil {
		var ret int64
		return ret
	}
	return *o.ComputeRespStatus5xx
}

// GetComputeRespStatus5xxOk returns a tuple with the ComputeRespStatus5xx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetComputeRespStatus5xxOk() (*int64, bool) {
	if o == nil || o.ComputeRespStatus5xx == nil {
		return nil, false
	}
	return o.ComputeRespStatus5xx, true
}

// HasComputeRespStatus5xx returns a boolean if a field has been set.
func (o *PlatformValues) HasComputeRespStatus5xx() bool {
	if o != nil && o.ComputeRespStatus5xx != nil {
		return true
	}

	return false
}

// SetComputeRespStatus5xx gets a reference to the given int64 and assigns it to the ComputeRespStatus5xx field.
func (o *PlatformValues) SetComputeRespStatus5xx(v int64) {
	o.ComputeRespStatus5xx = &v
}

// GetComputeBereqHeaderBytes returns the ComputeBereqHeaderBytes field value if set, zero value otherwise.
func (o *PlatformValues) GetComputeBereqHeaderBytes() int64 {
	if o == nil || o.ComputeBereqHeaderBytes == nil {
		var ret int64
		return ret
	}
	return *o.ComputeBereqHeaderBytes
}

// GetComputeBereqHeaderBytesOk returns a tuple with the ComputeBereqHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetComputeBereqHeaderBytesOk() (*int64, bool) {
	if o == nil || o.ComputeBereqHeaderBytes == nil {
		return nil, false
	}
	return o.ComputeBereqHeaderBytes, true
}

// HasComputeBereqHeaderBytes returns a boolean if a field has been set.
func (o *PlatformValues) HasComputeBereqHeaderBytes() bool {
	if o != nil && o.ComputeBereqHeaderBytes != nil {
		return true
	}

	return false
}

// SetComputeBereqHeaderBytes gets a reference to the given int64 and assigns it to the ComputeBereqHeaderBytes field.
func (o *PlatformValues) SetComputeBereqHeaderBytes(v int64) {
	o.ComputeBereqHeaderBytes = &v
}

// GetComputeBereqBodyBytes returns the ComputeBereqBodyBytes field value if set, zero value otherwise.
func (o *PlatformValues) GetComputeBereqBodyBytes() int64 {
	if o == nil || o.ComputeBereqBodyBytes == nil {
		var ret int64
		return ret
	}
	return *o.ComputeBereqBodyBytes
}

// GetComputeBereqBodyBytesOk returns a tuple with the ComputeBereqBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetComputeBereqBodyBytesOk() (*int64, bool) {
	if o == nil || o.ComputeBereqBodyBytes == nil {
		return nil, false
	}
	return o.ComputeBereqBodyBytes, true
}

// HasComputeBereqBodyBytes returns a boolean if a field has been set.
func (o *PlatformValues) HasComputeBereqBodyBytes() bool {
	if o != nil && o.ComputeBereqBodyBytes != nil {
		return true
	}

	return false
}

// SetComputeBereqBodyBytes gets a reference to the given int64 and assigns it to the ComputeBereqBodyBytes field.
func (o *PlatformValues) SetComputeBereqBodyBytes(v int64) {
	o.ComputeBereqBodyBytes = &v
}

// GetComputeBerespHeaderBytes returns the ComputeBerespHeaderBytes field value if set, zero value otherwise.
func (o *PlatformValues) GetComputeBerespHeaderBytes() int64 {
	if o == nil || o.ComputeBerespHeaderBytes == nil {
		var ret int64
		return ret
	}
	return *o.ComputeBerespHeaderBytes
}

// GetComputeBerespHeaderBytesOk returns a tuple with the ComputeBerespHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetComputeBerespHeaderBytesOk() (*int64, bool) {
	if o == nil || o.ComputeBerespHeaderBytes == nil {
		return nil, false
	}
	return o.ComputeBerespHeaderBytes, true
}

// HasComputeBerespHeaderBytes returns a boolean if a field has been set.
func (o *PlatformValues) HasComputeBerespHeaderBytes() bool {
	if o != nil && o.ComputeBerespHeaderBytes != nil {
		return true
	}

	return false
}

// SetComputeBerespHeaderBytes gets a reference to the given int64 and assigns it to the ComputeBerespHeaderBytes field.
func (o *PlatformValues) SetComputeBerespHeaderBytes(v int64) {
	o.ComputeBerespHeaderBytes = &v
}

// GetComputeBerespBodyBytes returns the ComputeBerespBodyBytes field value if set, zero value otherwise.
func (o *PlatformValues) GetComputeBerespBodyBytes() int64 {
	if o == nil || o.ComputeBerespBodyBytes == nil {
		var ret int64
		return ret
	}
	return *o.ComputeBerespBodyBytes
}

// GetComputeBerespBodyBytesOk returns a tuple with the ComputeBerespBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetComputeBerespBodyBytesOk() (*int64, bool) {
	if o == nil || o.ComputeBerespBodyBytes == nil {
		return nil, false
	}
	return o.ComputeBerespBodyBytes, true
}

// HasComputeBerespBodyBytes returns a boolean if a field has been set.
func (o *PlatformValues) HasComputeBerespBodyBytes() bool {
	if o != nil && o.ComputeBerespBodyBytes != nil {
		return true
	}

	return false
}

// SetComputeBerespBodyBytes gets a reference to the given int64 and assigns it to the ComputeBerespBodyBytes field.
func (o *PlatformValues) SetComputeBerespBodyBytes(v int64) {
	o.ComputeBerespBodyBytes = &v
}

// GetComputeBereqs returns the ComputeBereqs field value if set, zero value otherwise.
func (o *PlatformValues) GetComputeBereqs() int64 {
	if o == nil || o.ComputeBereqs == nil {
		var ret int64
		return ret
	}
	return *o.ComputeBereqs
}

// GetComputeBereqsOk returns a tuple with the ComputeBereqs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetComputeBereqsOk() (*int64, bool) {
	if o == nil || o.ComputeBereqs == nil {
		return nil, false
	}
	return o.ComputeBereqs, true
}

// HasComputeBereqs returns a boolean if a field has been set.
func (o *PlatformValues) HasComputeBereqs() bool {
	if o != nil && o.ComputeBereqs != nil {
		return true
	}

	return false
}

// SetComputeBereqs gets a reference to the given int64 and assigns it to the ComputeBereqs field.
func (o *PlatformValues) SetComputeBereqs(v int64) {
	o.ComputeBereqs = &v
}

// GetComputeServiceBereqError returns the ComputeServiceBereqError field value if set, zero value otherwise.
func (o *PlatformValues) GetComputeServiceBereqError() int64 {
	if o == nil || o.ComputeServiceBereqError == nil {
		var ret int64
		return ret
	}
	return *o.ComputeServiceBereqError
}

// GetComputeServiceBereqErrorOk returns a tuple with the ComputeServiceBereqError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetComputeServiceBereqErrorOk() (*int64, bool) {
	if o == nil || o.ComputeServiceBereqError == nil {
		return nil, false
	}
	return o.ComputeServiceBereqError, true
}

// HasComputeServiceBereqError returns a boolean if a field has been set.
func (o *PlatformValues) HasComputeServiceBereqError() bool {
	if o != nil && o.ComputeServiceBereqError != nil {
		return true
	}

	return false
}

// SetComputeServiceBereqError gets a reference to the given int64 and assigns it to the ComputeServiceBereqError field.
func (o *PlatformValues) SetComputeServiceBereqError(v int64) {
	o.ComputeServiceBereqError = &v
}

// GetComputeServiceMemoryExceededError returns the ComputeServiceMemoryExceededError field value if set, zero value otherwise.
func (o *PlatformValues) GetComputeServiceMemoryExceededError() int64 {
	if o == nil || o.ComputeServiceMemoryExceededError == nil {
		var ret int64
		return ret
	}
	return *o.ComputeServiceMemoryExceededError
}

// GetComputeServiceMemoryExceededErrorOk returns a tuple with the ComputeServiceMemoryExceededError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetComputeServiceMemoryExceededErrorOk() (*int64, bool) {
	if o == nil || o.ComputeServiceMemoryExceededError == nil {
		return nil, false
	}
	return o.ComputeServiceMemoryExceededError, true
}

// HasComputeServiceMemoryExceededError returns a boolean if a field has been set.
func (o *PlatformValues) HasComputeServiceMemoryExceededError() bool {
	if o != nil && o.ComputeServiceMemoryExceededError != nil {
		return true
	}

	return false
}

// SetComputeServiceMemoryExceededError gets a reference to the given int64 and assigns it to the ComputeServiceMemoryExceededError field.
func (o *PlatformValues) SetComputeServiceMemoryExceededError(v int64) {
	o.ComputeServiceMemoryExceededError = &v
}

// GetEdgeHitRespBodyBytes returns the EdgeHitRespBodyBytes field value if set, zero value otherwise.
func (o *PlatformValues) GetEdgeHitRespBodyBytes() int64 {
	if o == nil || o.EdgeHitRespBodyBytes == nil {
		var ret int64
		return ret
	}
	return *o.EdgeHitRespBodyBytes
}

// GetEdgeHitRespBodyBytesOk returns a tuple with the EdgeHitRespBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetEdgeHitRespBodyBytesOk() (*int64, bool) {
	if o == nil || o.EdgeHitRespBodyBytes == nil {
		return nil, false
	}
	return o.EdgeHitRespBodyBytes, true
}

// HasEdgeHitRespBodyBytes returns a boolean if a field has been set.
func (o *PlatformValues) HasEdgeHitRespBodyBytes() bool {
	if o != nil && o.EdgeHitRespBodyBytes != nil {
		return true
	}

	return false
}

// SetEdgeHitRespBodyBytes gets a reference to the given int64 and assigns it to the EdgeHitRespBodyBytes field.
func (o *PlatformValues) SetEdgeHitRespBodyBytes(v int64) {
	o.EdgeHitRespBodyBytes = &v
}

// GetEdgeHitRespHeaderBytes returns the EdgeHitRespHeaderBytes field value if set, zero value otherwise.
func (o *PlatformValues) GetEdgeHitRespHeaderBytes() int64 {
	if o == nil || o.EdgeHitRespHeaderBytes == nil {
		var ret int64
		return ret
	}
	return *o.EdgeHitRespHeaderBytes
}

// GetEdgeHitRespHeaderBytesOk returns a tuple with the EdgeHitRespHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetEdgeHitRespHeaderBytesOk() (*int64, bool) {
	if o == nil || o.EdgeHitRespHeaderBytes == nil {
		return nil, false
	}
	return o.EdgeHitRespHeaderBytes, true
}

// HasEdgeHitRespHeaderBytes returns a boolean if a field has been set.
func (o *PlatformValues) HasEdgeHitRespHeaderBytes() bool {
	if o != nil && o.EdgeHitRespHeaderBytes != nil {
		return true
	}

	return false
}

// SetEdgeHitRespHeaderBytes gets a reference to the given int64 and assigns it to the EdgeHitRespHeaderBytes field.
func (o *PlatformValues) SetEdgeHitRespHeaderBytes(v int64) {
	o.EdgeHitRespHeaderBytes = &v
}

// GetEdgeMissRespBodyBytes returns the EdgeMissRespBodyBytes field value if set, zero value otherwise.
func (o *PlatformValues) GetEdgeMissRespBodyBytes() int64 {
	if o == nil || o.EdgeMissRespBodyBytes == nil {
		var ret int64
		return ret
	}
	return *o.EdgeMissRespBodyBytes
}

// GetEdgeMissRespBodyBytesOk returns a tuple with the EdgeMissRespBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetEdgeMissRespBodyBytesOk() (*int64, bool) {
	if o == nil || o.EdgeMissRespBodyBytes == nil {
		return nil, false
	}
	return o.EdgeMissRespBodyBytes, true
}

// HasEdgeMissRespBodyBytes returns a boolean if a field has been set.
func (o *PlatformValues) HasEdgeMissRespBodyBytes() bool {
	if o != nil && o.EdgeMissRespBodyBytes != nil {
		return true
	}

	return false
}

// SetEdgeMissRespBodyBytes gets a reference to the given int64 and assigns it to the EdgeMissRespBodyBytes field.
func (o *PlatformValues) SetEdgeMissRespBodyBytes(v int64) {
	o.EdgeMissRespBodyBytes = &v
}

// GetEdgeMissRespHeaderBytes returns the EdgeMissRespHeaderBytes field value if set, zero value otherwise.
func (o *PlatformValues) GetEdgeMissRespHeaderBytes() int64 {
	if o == nil || o.EdgeMissRespHeaderBytes == nil {
		var ret int64
		return ret
	}
	return *o.EdgeMissRespHeaderBytes
}

// GetEdgeMissRespHeaderBytesOk returns a tuple with the EdgeMissRespHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetEdgeMissRespHeaderBytesOk() (*int64, bool) {
	if o == nil || o.EdgeMissRespHeaderBytes == nil {
		return nil, false
	}
	return o.EdgeMissRespHeaderBytes, true
}

// HasEdgeMissRespHeaderBytes returns a boolean if a field has been set.
func (o *PlatformValues) HasEdgeMissRespHeaderBytes() bool {
	if o != nil && o.EdgeMissRespHeaderBytes != nil {
		return true
	}

	return false
}

// SetEdgeMissRespHeaderBytes gets a reference to the given int64 and assigns it to the EdgeMissRespHeaderBytes field.
func (o *PlatformValues) SetEdgeMissRespHeaderBytes(v int64) {
	o.EdgeMissRespHeaderBytes = &v
}

// GetOriginCacheFetchRespBodyBytes returns the OriginCacheFetchRespBodyBytes field value if set, zero value otherwise.
func (o *PlatformValues) GetOriginCacheFetchRespBodyBytes() int64 {
	if o == nil || o.OriginCacheFetchRespBodyBytes == nil {
		var ret int64
		return ret
	}
	return *o.OriginCacheFetchRespBodyBytes
}

// GetOriginCacheFetchRespBodyBytesOk returns a tuple with the OriginCacheFetchRespBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetOriginCacheFetchRespBodyBytesOk() (*int64, bool) {
	if o == nil || o.OriginCacheFetchRespBodyBytes == nil {
		return nil, false
	}
	return o.OriginCacheFetchRespBodyBytes, true
}

// HasOriginCacheFetchRespBodyBytes returns a boolean if a field has been set.
func (o *PlatformValues) HasOriginCacheFetchRespBodyBytes() bool {
	if o != nil && o.OriginCacheFetchRespBodyBytes != nil {
		return true
	}

	return false
}

// SetOriginCacheFetchRespBodyBytes gets a reference to the given int64 and assigns it to the OriginCacheFetchRespBodyBytes field.
func (o *PlatformValues) SetOriginCacheFetchRespBodyBytes(v int64) {
	o.OriginCacheFetchRespBodyBytes = &v
}

// GetOriginCacheFetchRespHeaderBytes returns the OriginCacheFetchRespHeaderBytes field value if set, zero value otherwise.
func (o *PlatformValues) GetOriginCacheFetchRespHeaderBytes() int64 {
	if o == nil || o.OriginCacheFetchRespHeaderBytes == nil {
		var ret int64
		return ret
	}
	return *o.OriginCacheFetchRespHeaderBytes
}

// GetOriginCacheFetchRespHeaderBytesOk returns a tuple with the OriginCacheFetchRespHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetOriginCacheFetchRespHeaderBytesOk() (*int64, bool) {
	if o == nil || o.OriginCacheFetchRespHeaderBytes == nil {
		return nil, false
	}
	return o.OriginCacheFetchRespHeaderBytes, true
}

// HasOriginCacheFetchRespHeaderBytes returns a boolean if a field has been set.
func (o *PlatformValues) HasOriginCacheFetchRespHeaderBytes() bool {
	if o != nil && o.OriginCacheFetchRespHeaderBytes != nil {
		return true
	}

	return false
}

// SetOriginCacheFetchRespHeaderBytes gets a reference to the given int64 and assigns it to the OriginCacheFetchRespHeaderBytes field.
func (o *PlatformValues) SetOriginCacheFetchRespHeaderBytes(v int64) {
	o.OriginCacheFetchRespHeaderBytes = &v
}

// GetShieldHitRequests returns the ShieldHitRequests field value if set, zero value otherwise.
func (o *PlatformValues) GetShieldHitRequests() int64 {
	if o == nil || o.ShieldHitRequests == nil {
		var ret int64
		return ret
	}
	return *o.ShieldHitRequests
}

// GetShieldHitRequestsOk returns a tuple with the ShieldHitRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetShieldHitRequestsOk() (*int64, bool) {
	if o == nil || o.ShieldHitRequests == nil {
		return nil, false
	}
	return o.ShieldHitRequests, true
}

// HasShieldHitRequests returns a boolean if a field has been set.
func (o *PlatformValues) HasShieldHitRequests() bool {
	if o != nil && o.ShieldHitRequests != nil {
		return true
	}

	return false
}

// SetShieldHitRequests gets a reference to the given int64 and assigns it to the ShieldHitRequests field.
func (o *PlatformValues) SetShieldHitRequests(v int64) {
	o.ShieldHitRequests = &v
}

// GetShieldMissRequests returns the ShieldMissRequests field value if set, zero value otherwise.
func (o *PlatformValues) GetShieldMissRequests() int64 {
	if o == nil || o.ShieldMissRequests == nil {
		var ret int64
		return ret
	}
	return *o.ShieldMissRequests
}

// GetShieldMissRequestsOk returns a tuple with the ShieldMissRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetShieldMissRequestsOk() (*int64, bool) {
	if o == nil || o.ShieldMissRequests == nil {
		return nil, false
	}
	return o.ShieldMissRequests, true
}

// HasShieldMissRequests returns a boolean if a field has been set.
func (o *PlatformValues) HasShieldMissRequests() bool {
	if o != nil && o.ShieldMissRequests != nil {
		return true
	}

	return false
}

// SetShieldMissRequests gets a reference to the given int64 and assigns it to the ShieldMissRequests field.
func (o *PlatformValues) SetShieldMissRequests(v int64) {
	o.ShieldMissRequests = &v
}

// GetShieldHitRespHeaderBytes returns the ShieldHitRespHeaderBytes field value if set, zero value otherwise.
func (o *PlatformValues) GetShieldHitRespHeaderBytes() int64 {
	if o == nil || o.ShieldHitRespHeaderBytes == nil {
		var ret int64
		return ret
	}
	return *o.ShieldHitRespHeaderBytes
}

// GetShieldHitRespHeaderBytesOk returns a tuple with the ShieldHitRespHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetShieldHitRespHeaderBytesOk() (*int64, bool) {
	if o == nil || o.ShieldHitRespHeaderBytes == nil {
		return nil, false
	}
	return o.ShieldHitRespHeaderBytes, true
}

// HasShieldHitRespHeaderBytes returns a boolean if a field has been set.
func (o *PlatformValues) HasShieldHitRespHeaderBytes() bool {
	if o != nil && o.ShieldHitRespHeaderBytes != nil {
		return true
	}

	return false
}

// SetShieldHitRespHeaderBytes gets a reference to the given int64 and assigns it to the ShieldHitRespHeaderBytes field.
func (o *PlatformValues) SetShieldHitRespHeaderBytes(v int64) {
	o.ShieldHitRespHeaderBytes = &v
}

// GetShieldHitRespBodyBytes returns the ShieldHitRespBodyBytes field value if set, zero value otherwise.
func (o *PlatformValues) GetShieldHitRespBodyBytes() int64 {
	if o == nil || o.ShieldHitRespBodyBytes == nil {
		var ret int64
		return ret
	}
	return *o.ShieldHitRespBodyBytes
}

// GetShieldHitRespBodyBytesOk returns a tuple with the ShieldHitRespBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetShieldHitRespBodyBytesOk() (*int64, bool) {
	if o == nil || o.ShieldHitRespBodyBytes == nil {
		return nil, false
	}
	return o.ShieldHitRespBodyBytes, true
}

// HasShieldHitRespBodyBytes returns a boolean if a field has been set.
func (o *PlatformValues) HasShieldHitRespBodyBytes() bool {
	if o != nil && o.ShieldHitRespBodyBytes != nil {
		return true
	}

	return false
}

// SetShieldHitRespBodyBytes gets a reference to the given int64 and assigns it to the ShieldHitRespBodyBytes field.
func (o *PlatformValues) SetShieldHitRespBodyBytes(v int64) {
	o.ShieldHitRespBodyBytes = &v
}

// GetShieldMissRespHeaderBytes returns the ShieldMissRespHeaderBytes field value if set, zero value otherwise.
func (o *PlatformValues) GetShieldMissRespHeaderBytes() int64 {
	if o == nil || o.ShieldMissRespHeaderBytes == nil {
		var ret int64
		return ret
	}
	return *o.ShieldMissRespHeaderBytes
}

// GetShieldMissRespHeaderBytesOk returns a tuple with the ShieldMissRespHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetShieldMissRespHeaderBytesOk() (*int64, bool) {
	if o == nil || o.ShieldMissRespHeaderBytes == nil {
		return nil, false
	}
	return o.ShieldMissRespHeaderBytes, true
}

// HasShieldMissRespHeaderBytes returns a boolean if a field has been set.
func (o *PlatformValues) HasShieldMissRespHeaderBytes() bool {
	if o != nil && o.ShieldMissRespHeaderBytes != nil {
		return true
	}

	return false
}

// SetShieldMissRespHeaderBytes gets a reference to the given int64 and assigns it to the ShieldMissRespHeaderBytes field.
func (o *PlatformValues) SetShieldMissRespHeaderBytes(v int64) {
	o.ShieldMissRespHeaderBytes = &v
}

// GetShieldMissRespBodyBytes returns the ShieldMissRespBodyBytes field value if set, zero value otherwise.
func (o *PlatformValues) GetShieldMissRespBodyBytes() int64 {
	if o == nil || o.ShieldMissRespBodyBytes == nil {
		var ret int64
		return ret
	}
	return *o.ShieldMissRespBodyBytes
}

// GetShieldMissRespBodyBytesOk returns a tuple with the ShieldMissRespBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetShieldMissRespBodyBytesOk() (*int64, bool) {
	if o == nil || o.ShieldMissRespBodyBytes == nil {
		return nil, false
	}
	return o.ShieldMissRespBodyBytes, true
}

// HasShieldMissRespBodyBytes returns a boolean if a field has been set.
func (o *PlatformValues) HasShieldMissRespBodyBytes() bool {
	if o != nil && o.ShieldMissRespBodyBytes != nil {
		return true
	}

	return false
}

// SetShieldMissRespBodyBytes gets a reference to the given int64 and assigns it to the ShieldMissRespBodyBytes field.
func (o *PlatformValues) SetShieldMissRespBodyBytes(v int64) {
	o.ShieldMissRespBodyBytes = &v
}

// GetWebsocketReqHeaderBytes returns the WebsocketReqHeaderBytes field value if set, zero value otherwise.
func (o *PlatformValues) GetWebsocketReqHeaderBytes() int64 {
	if o == nil || o.WebsocketReqHeaderBytes == nil {
		var ret int64
		return ret
	}
	return *o.WebsocketReqHeaderBytes
}

// GetWebsocketReqHeaderBytesOk returns a tuple with the WebsocketReqHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetWebsocketReqHeaderBytesOk() (*int64, bool) {
	if o == nil || o.WebsocketReqHeaderBytes == nil {
		return nil, false
	}
	return o.WebsocketReqHeaderBytes, true
}

// HasWebsocketReqHeaderBytes returns a boolean if a field has been set.
func (o *PlatformValues) HasWebsocketReqHeaderBytes() bool {
	if o != nil && o.WebsocketReqHeaderBytes != nil {
		return true
	}

	return false
}

// SetWebsocketReqHeaderBytes gets a reference to the given int64 and assigns it to the WebsocketReqHeaderBytes field.
func (o *PlatformValues) SetWebsocketReqHeaderBytes(v int64) {
	o.WebsocketReqHeaderBytes = &v
}

// GetWebsocketReqBodyBytes returns the WebsocketReqBodyBytes field value if set, zero value otherwise.
func (o *PlatformValues) GetWebsocketReqBodyBytes() int64 {
	if o == nil || o.WebsocketReqBodyBytes == nil {
		var ret int64
		return ret
	}
	return *o.WebsocketReqBodyBytes
}

// GetWebsocketReqBodyBytesOk returns a tuple with the WebsocketReqBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetWebsocketReqBodyBytesOk() (*int64, bool) {
	if o == nil || o.WebsocketReqBodyBytes == nil {
		return nil, false
	}
	return o.WebsocketReqBodyBytes, true
}

// HasWebsocketReqBodyBytes returns a boolean if a field has been set.
func (o *PlatformValues) HasWebsocketReqBodyBytes() bool {
	if o != nil && o.WebsocketReqBodyBytes != nil {
		return true
	}

	return false
}

// SetWebsocketReqBodyBytes gets a reference to the given int64 and assigns it to the WebsocketReqBodyBytes field.
func (o *PlatformValues) SetWebsocketReqBodyBytes(v int64) {
	o.WebsocketReqBodyBytes = &v
}

// GetWebsocketRespHeaderBytes returns the WebsocketRespHeaderBytes field value if set, zero value otherwise.
func (o *PlatformValues) GetWebsocketRespHeaderBytes() int64 {
	if o == nil || o.WebsocketRespHeaderBytes == nil {
		var ret int64
		return ret
	}
	return *o.WebsocketRespHeaderBytes
}

// GetWebsocketRespHeaderBytesOk returns a tuple with the WebsocketRespHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetWebsocketRespHeaderBytesOk() (*int64, bool) {
	if o == nil || o.WebsocketRespHeaderBytes == nil {
		return nil, false
	}
	return o.WebsocketRespHeaderBytes, true
}

// HasWebsocketRespHeaderBytes returns a boolean if a field has been set.
func (o *PlatformValues) HasWebsocketRespHeaderBytes() bool {
	if o != nil && o.WebsocketRespHeaderBytes != nil {
		return true
	}

	return false
}

// SetWebsocketRespHeaderBytes gets a reference to the given int64 and assigns it to the WebsocketRespHeaderBytes field.
func (o *PlatformValues) SetWebsocketRespHeaderBytes(v int64) {
	o.WebsocketRespHeaderBytes = &v
}

// GetWebsocketRespBodyBytes returns the WebsocketRespBodyBytes field value if set, zero value otherwise.
func (o *PlatformValues) GetWebsocketRespBodyBytes() int64 {
	if o == nil || o.WebsocketRespBodyBytes == nil {
		var ret int64
		return ret
	}
	return *o.WebsocketRespBodyBytes
}

// GetWebsocketRespBodyBytesOk returns a tuple with the WebsocketRespBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetWebsocketRespBodyBytesOk() (*int64, bool) {
	if o == nil || o.WebsocketRespBodyBytes == nil {
		return nil, false
	}
	return o.WebsocketRespBodyBytes, true
}

// HasWebsocketRespBodyBytes returns a boolean if a field has been set.
func (o *PlatformValues) HasWebsocketRespBodyBytes() bool {
	if o != nil && o.WebsocketRespBodyBytes != nil {
		return true
	}

	return false
}

// SetWebsocketRespBodyBytes gets a reference to the given int64 and assigns it to the WebsocketRespBodyBytes field.
func (o *PlatformValues) SetWebsocketRespBodyBytes(v int64) {
	o.WebsocketRespBodyBytes = &v
}

// GetWebsocketBereqHeaderBytes returns the WebsocketBereqHeaderBytes field value if set, zero value otherwise.
func (o *PlatformValues) GetWebsocketBereqHeaderBytes() int64 {
	if o == nil || o.WebsocketBereqHeaderBytes == nil {
		var ret int64
		return ret
	}
	return *o.WebsocketBereqHeaderBytes
}

// GetWebsocketBereqHeaderBytesOk returns a tuple with the WebsocketBereqHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetWebsocketBereqHeaderBytesOk() (*int64, bool) {
	if o == nil || o.WebsocketBereqHeaderBytes == nil {
		return nil, false
	}
	return o.WebsocketBereqHeaderBytes, true
}

// HasWebsocketBereqHeaderBytes returns a boolean if a field has been set.
func (o *PlatformValues) HasWebsocketBereqHeaderBytes() bool {
	if o != nil && o.WebsocketBereqHeaderBytes != nil {
		return true
	}

	return false
}

// SetWebsocketBereqHeaderBytes gets a reference to the given int64 and assigns it to the WebsocketBereqHeaderBytes field.
func (o *PlatformValues) SetWebsocketBereqHeaderBytes(v int64) {
	o.WebsocketBereqHeaderBytes = &v
}

// GetWebsocketBereqBodyBytes returns the WebsocketBereqBodyBytes field value if set, zero value otherwise.
func (o *PlatformValues) GetWebsocketBereqBodyBytes() int64 {
	if o == nil || o.WebsocketBereqBodyBytes == nil {
		var ret int64
		return ret
	}
	return *o.WebsocketBereqBodyBytes
}

// GetWebsocketBereqBodyBytesOk returns a tuple with the WebsocketBereqBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetWebsocketBereqBodyBytesOk() (*int64, bool) {
	if o == nil || o.WebsocketBereqBodyBytes == nil {
		return nil, false
	}
	return o.WebsocketBereqBodyBytes, true
}

// HasWebsocketBereqBodyBytes returns a boolean if a field has been set.
func (o *PlatformValues) HasWebsocketBereqBodyBytes() bool {
	if o != nil && o.WebsocketBereqBodyBytes != nil {
		return true
	}

	return false
}

// SetWebsocketBereqBodyBytes gets a reference to the given int64 and assigns it to the WebsocketBereqBodyBytes field.
func (o *PlatformValues) SetWebsocketBereqBodyBytes(v int64) {
	o.WebsocketBereqBodyBytes = &v
}

// GetWebsocketBerespHeaderBytes returns the WebsocketBerespHeaderBytes field value if set, zero value otherwise.
func (o *PlatformValues) GetWebsocketBerespHeaderBytes() int64 {
	if o == nil || o.WebsocketBerespHeaderBytes == nil {
		var ret int64
		return ret
	}
	return *o.WebsocketBerespHeaderBytes
}

// GetWebsocketBerespHeaderBytesOk returns a tuple with the WebsocketBerespHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetWebsocketBerespHeaderBytesOk() (*int64, bool) {
	if o == nil || o.WebsocketBerespHeaderBytes == nil {
		return nil, false
	}
	return o.WebsocketBerespHeaderBytes, true
}

// HasWebsocketBerespHeaderBytes returns a boolean if a field has been set.
func (o *PlatformValues) HasWebsocketBerespHeaderBytes() bool {
	if o != nil && o.WebsocketBerespHeaderBytes != nil {
		return true
	}

	return false
}

// SetWebsocketBerespHeaderBytes gets a reference to the given int64 and assigns it to the WebsocketBerespHeaderBytes field.
func (o *PlatformValues) SetWebsocketBerespHeaderBytes(v int64) {
	o.WebsocketBerespHeaderBytes = &v
}

// GetWebsocketBerespBodyBytes returns the WebsocketBerespBodyBytes field value if set, zero value otherwise.
func (o *PlatformValues) GetWebsocketBerespBodyBytes() int64 {
	if o == nil || o.WebsocketBerespBodyBytes == nil {
		var ret int64
		return ret
	}
	return *o.WebsocketBerespBodyBytes
}

// GetWebsocketBerespBodyBytesOk returns a tuple with the WebsocketBerespBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetWebsocketBerespBodyBytesOk() (*int64, bool) {
	if o == nil || o.WebsocketBerespBodyBytes == nil {
		return nil, false
	}
	return o.WebsocketBerespBodyBytes, true
}

// HasWebsocketBerespBodyBytes returns a boolean if a field has been set.
func (o *PlatformValues) HasWebsocketBerespBodyBytes() bool {
	if o != nil && o.WebsocketBerespBodyBytes != nil {
		return true
	}

	return false
}

// SetWebsocketBerespBodyBytes gets a reference to the given int64 and assigns it to the WebsocketBerespBodyBytes field.
func (o *PlatformValues) SetWebsocketBerespBodyBytes(v int64) {
	o.WebsocketBerespBodyBytes = &v
}

// GetWebsocketConnTimeMs returns the WebsocketConnTimeMs field value if set, zero value otherwise.
func (o *PlatformValues) GetWebsocketConnTimeMs() int64 {
	if o == nil || o.WebsocketConnTimeMs == nil {
		var ret int64
		return ret
	}
	return *o.WebsocketConnTimeMs
}

// GetWebsocketConnTimeMsOk returns a tuple with the WebsocketConnTimeMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetWebsocketConnTimeMsOk() (*int64, bool) {
	if o == nil || o.WebsocketConnTimeMs == nil {
		return nil, false
	}
	return o.WebsocketConnTimeMs, true
}

// HasWebsocketConnTimeMs returns a boolean if a field has been set.
func (o *PlatformValues) HasWebsocketConnTimeMs() bool {
	if o != nil && o.WebsocketConnTimeMs != nil {
		return true
	}

	return false
}

// SetWebsocketConnTimeMs gets a reference to the given int64 and assigns it to the WebsocketConnTimeMs field.
func (o *PlatformValues) SetWebsocketConnTimeMs(v int64) {
	o.WebsocketConnTimeMs = &v
}

// GetFanoutRecvPublishes returns the FanoutRecvPublishes field value if set, zero value otherwise.
func (o *PlatformValues) GetFanoutRecvPublishes() int64 {
	if o == nil || o.FanoutRecvPublishes == nil {
		var ret int64
		return ret
	}
	return *o.FanoutRecvPublishes
}

// GetFanoutRecvPublishesOk returns a tuple with the FanoutRecvPublishes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetFanoutRecvPublishesOk() (*int64, bool) {
	if o == nil || o.FanoutRecvPublishes == nil {
		return nil, false
	}
	return o.FanoutRecvPublishes, true
}

// HasFanoutRecvPublishes returns a boolean if a field has been set.
func (o *PlatformValues) HasFanoutRecvPublishes() bool {
	if o != nil && o.FanoutRecvPublishes != nil {
		return true
	}

	return false
}

// SetFanoutRecvPublishes gets a reference to the given int64 and assigns it to the FanoutRecvPublishes field.
func (o *PlatformValues) SetFanoutRecvPublishes(v int64) {
	o.FanoutRecvPublishes = &v
}

// GetFanoutSendPublishes returns the FanoutSendPublishes field value if set, zero value otherwise.
func (o *PlatformValues) GetFanoutSendPublishes() int64 {
	if o == nil || o.FanoutSendPublishes == nil {
		var ret int64
		return ret
	}
	return *o.FanoutSendPublishes
}

// GetFanoutSendPublishesOk returns a tuple with the FanoutSendPublishes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetFanoutSendPublishesOk() (*int64, bool) {
	if o == nil || o.FanoutSendPublishes == nil {
		return nil, false
	}
	return o.FanoutSendPublishes, true
}

// HasFanoutSendPublishes returns a boolean if a field has been set.
func (o *PlatformValues) HasFanoutSendPublishes() bool {
	if o != nil && o.FanoutSendPublishes != nil {
		return true
	}

	return false
}

// SetFanoutSendPublishes gets a reference to the given int64 and assigns it to the FanoutSendPublishes field.
func (o *PlatformValues) SetFanoutSendPublishes(v int64) {
	o.FanoutSendPublishes = &v
}

// GetKvStoreClassAOperations returns the KvStoreClassAOperations field value if set, zero value otherwise.
func (o *PlatformValues) GetKvStoreClassAOperations() int64 {
	if o == nil || o.KvStoreClassAOperations == nil {
		var ret int64
		return ret
	}
	return *o.KvStoreClassAOperations
}

// GetKvStoreClassAOperationsOk returns a tuple with the KvStoreClassAOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetKvStoreClassAOperationsOk() (*int64, bool) {
	if o == nil || o.KvStoreClassAOperations == nil {
		return nil, false
	}
	return o.KvStoreClassAOperations, true
}

// HasKvStoreClassAOperations returns a boolean if a field has been set.
func (o *PlatformValues) HasKvStoreClassAOperations() bool {
	if o != nil && o.KvStoreClassAOperations != nil {
		return true
	}

	return false
}

// SetKvStoreClassAOperations gets a reference to the given int64 and assigns it to the KvStoreClassAOperations field.
func (o *PlatformValues) SetKvStoreClassAOperations(v int64) {
	o.KvStoreClassAOperations = &v
}

// GetKvStoreClassBOperations returns the KvStoreClassBOperations field value if set, zero value otherwise.
func (o *PlatformValues) GetKvStoreClassBOperations() int64 {
	if o == nil || o.KvStoreClassBOperations == nil {
		var ret int64
		return ret
	}
	return *o.KvStoreClassBOperations
}

// GetKvStoreClassBOperationsOk returns a tuple with the KvStoreClassBOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetKvStoreClassBOperationsOk() (*int64, bool) {
	if o == nil || o.KvStoreClassBOperations == nil {
		return nil, false
	}
	return o.KvStoreClassBOperations, true
}

// HasKvStoreClassBOperations returns a boolean if a field has been set.
func (o *PlatformValues) HasKvStoreClassBOperations() bool {
	if o != nil && o.KvStoreClassBOperations != nil {
		return true
	}

	return false
}

// SetKvStoreClassBOperations gets a reference to the given int64 and assigns it to the KvStoreClassBOperations field.
func (o *PlatformValues) SetKvStoreClassBOperations(v int64) {
	o.KvStoreClassBOperations = &v
}

// GetFanoutReqHeaderBytes returns the FanoutReqHeaderBytes field value if set, zero value otherwise.
func (o *PlatformValues) GetFanoutReqHeaderBytes() int64 {
	if o == nil || o.FanoutReqHeaderBytes == nil {
		var ret int64
		return ret
	}
	return *o.FanoutReqHeaderBytes
}

// GetFanoutReqHeaderBytesOk returns a tuple with the FanoutReqHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetFanoutReqHeaderBytesOk() (*int64, bool) {
	if o == nil || o.FanoutReqHeaderBytes == nil {
		return nil, false
	}
	return o.FanoutReqHeaderBytes, true
}

// HasFanoutReqHeaderBytes returns a boolean if a field has been set.
func (o *PlatformValues) HasFanoutReqHeaderBytes() bool {
	if o != nil && o.FanoutReqHeaderBytes != nil {
		return true
	}

	return false
}

// SetFanoutReqHeaderBytes gets a reference to the given int64 and assigns it to the FanoutReqHeaderBytes field.
func (o *PlatformValues) SetFanoutReqHeaderBytes(v int64) {
	o.FanoutReqHeaderBytes = &v
}

// GetFanoutReqBodyBytes returns the FanoutReqBodyBytes field value if set, zero value otherwise.
func (o *PlatformValues) GetFanoutReqBodyBytes() int64 {
	if o == nil || o.FanoutReqBodyBytes == nil {
		var ret int64
		return ret
	}
	return *o.FanoutReqBodyBytes
}

// GetFanoutReqBodyBytesOk returns a tuple with the FanoutReqBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetFanoutReqBodyBytesOk() (*int64, bool) {
	if o == nil || o.FanoutReqBodyBytes == nil {
		return nil, false
	}
	return o.FanoutReqBodyBytes, true
}

// HasFanoutReqBodyBytes returns a boolean if a field has been set.
func (o *PlatformValues) HasFanoutReqBodyBytes() bool {
	if o != nil && o.FanoutReqBodyBytes != nil {
		return true
	}

	return false
}

// SetFanoutReqBodyBytes gets a reference to the given int64 and assigns it to the FanoutReqBodyBytes field.
func (o *PlatformValues) SetFanoutReqBodyBytes(v int64) {
	o.FanoutReqBodyBytes = &v
}

// GetFanoutRespHeaderBytes returns the FanoutRespHeaderBytes field value if set, zero value otherwise.
func (o *PlatformValues) GetFanoutRespHeaderBytes() int64 {
	if o == nil || o.FanoutRespHeaderBytes == nil {
		var ret int64
		return ret
	}
	return *o.FanoutRespHeaderBytes
}

// GetFanoutRespHeaderBytesOk returns a tuple with the FanoutRespHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetFanoutRespHeaderBytesOk() (*int64, bool) {
	if o == nil || o.FanoutRespHeaderBytes == nil {
		return nil, false
	}
	return o.FanoutRespHeaderBytes, true
}

// HasFanoutRespHeaderBytes returns a boolean if a field has been set.
func (o *PlatformValues) HasFanoutRespHeaderBytes() bool {
	if o != nil && o.FanoutRespHeaderBytes != nil {
		return true
	}

	return false
}

// SetFanoutRespHeaderBytes gets a reference to the given int64 and assigns it to the FanoutRespHeaderBytes field.
func (o *PlatformValues) SetFanoutRespHeaderBytes(v int64) {
	o.FanoutRespHeaderBytes = &v
}

// GetFanoutRespBodyBytes returns the FanoutRespBodyBytes field value if set, zero value otherwise.
func (o *PlatformValues) GetFanoutRespBodyBytes() int64 {
	if o == nil || o.FanoutRespBodyBytes == nil {
		var ret int64
		return ret
	}
	return *o.FanoutRespBodyBytes
}

// GetFanoutRespBodyBytesOk returns a tuple with the FanoutRespBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetFanoutRespBodyBytesOk() (*int64, bool) {
	if o == nil || o.FanoutRespBodyBytes == nil {
		return nil, false
	}
	return o.FanoutRespBodyBytes, true
}

// HasFanoutRespBodyBytes returns a boolean if a field has been set.
func (o *PlatformValues) HasFanoutRespBodyBytes() bool {
	if o != nil && o.FanoutRespBodyBytes != nil {
		return true
	}

	return false
}

// SetFanoutRespBodyBytes gets a reference to the given int64 and assigns it to the FanoutRespBodyBytes field.
func (o *PlatformValues) SetFanoutRespBodyBytes(v int64) {
	o.FanoutRespBodyBytes = &v
}

// GetFanoutBereqHeaderBytes returns the FanoutBereqHeaderBytes field value if set, zero value otherwise.
func (o *PlatformValues) GetFanoutBereqHeaderBytes() int64 {
	if o == nil || o.FanoutBereqHeaderBytes == nil {
		var ret int64
		return ret
	}
	return *o.FanoutBereqHeaderBytes
}

// GetFanoutBereqHeaderBytesOk returns a tuple with the FanoutBereqHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetFanoutBereqHeaderBytesOk() (*int64, bool) {
	if o == nil || o.FanoutBereqHeaderBytes == nil {
		return nil, false
	}
	return o.FanoutBereqHeaderBytes, true
}

// HasFanoutBereqHeaderBytes returns a boolean if a field has been set.
func (o *PlatformValues) HasFanoutBereqHeaderBytes() bool {
	if o != nil && o.FanoutBereqHeaderBytes != nil {
		return true
	}

	return false
}

// SetFanoutBereqHeaderBytes gets a reference to the given int64 and assigns it to the FanoutBereqHeaderBytes field.
func (o *PlatformValues) SetFanoutBereqHeaderBytes(v int64) {
	o.FanoutBereqHeaderBytes = &v
}

// GetFanoutBereqBodyBytes returns the FanoutBereqBodyBytes field value if set, zero value otherwise.
func (o *PlatformValues) GetFanoutBereqBodyBytes() int64 {
	if o == nil || o.FanoutBereqBodyBytes == nil {
		var ret int64
		return ret
	}
	return *o.FanoutBereqBodyBytes
}

// GetFanoutBereqBodyBytesOk returns a tuple with the FanoutBereqBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetFanoutBereqBodyBytesOk() (*int64, bool) {
	if o == nil || o.FanoutBereqBodyBytes == nil {
		return nil, false
	}
	return o.FanoutBereqBodyBytes, true
}

// HasFanoutBereqBodyBytes returns a boolean if a field has been set.
func (o *PlatformValues) HasFanoutBereqBodyBytes() bool {
	if o != nil && o.FanoutBereqBodyBytes != nil {
		return true
	}

	return false
}

// SetFanoutBereqBodyBytes gets a reference to the given int64 and assigns it to the FanoutBereqBodyBytes field.
func (o *PlatformValues) SetFanoutBereqBodyBytes(v int64) {
	o.FanoutBereqBodyBytes = &v
}

// GetFanoutBerespHeaderBytes returns the FanoutBerespHeaderBytes field value if set, zero value otherwise.
func (o *PlatformValues) GetFanoutBerespHeaderBytes() int64 {
	if o == nil || o.FanoutBerespHeaderBytes == nil {
		var ret int64
		return ret
	}
	return *o.FanoutBerespHeaderBytes
}

// GetFanoutBerespHeaderBytesOk returns a tuple with the FanoutBerespHeaderBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetFanoutBerespHeaderBytesOk() (*int64, bool) {
	if o == nil || o.FanoutBerespHeaderBytes == nil {
		return nil, false
	}
	return o.FanoutBerespHeaderBytes, true
}

// HasFanoutBerespHeaderBytes returns a boolean if a field has been set.
func (o *PlatformValues) HasFanoutBerespHeaderBytes() bool {
	if o != nil && o.FanoutBerespHeaderBytes != nil {
		return true
	}

	return false
}

// SetFanoutBerespHeaderBytes gets a reference to the given int64 and assigns it to the FanoutBerespHeaderBytes field.
func (o *PlatformValues) SetFanoutBerespHeaderBytes(v int64) {
	o.FanoutBerespHeaderBytes = &v
}

// GetFanoutBerespBodyBytes returns the FanoutBerespBodyBytes field value if set, zero value otherwise.
func (o *PlatformValues) GetFanoutBerespBodyBytes() int64 {
	if o == nil || o.FanoutBerespBodyBytes == nil {
		var ret int64
		return ret
	}
	return *o.FanoutBerespBodyBytes
}

// GetFanoutBerespBodyBytesOk returns a tuple with the FanoutBerespBodyBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetFanoutBerespBodyBytesOk() (*int64, bool) {
	if o == nil || o.FanoutBerespBodyBytes == nil {
		return nil, false
	}
	return o.FanoutBerespBodyBytes, true
}

// HasFanoutBerespBodyBytes returns a boolean if a field has been set.
func (o *PlatformValues) HasFanoutBerespBodyBytes() bool {
	if o != nil && o.FanoutBerespBodyBytes != nil {
		return true
	}

	return false
}

// SetFanoutBerespBodyBytes gets a reference to the given int64 and assigns it to the FanoutBerespBodyBytes field.
func (o *PlatformValues) SetFanoutBerespBodyBytes(v int64) {
	o.FanoutBerespBodyBytes = &v
}

// GetFanoutConnTimeMs returns the FanoutConnTimeMs field value if set, zero value otherwise.
func (o *PlatformValues) GetFanoutConnTimeMs() int64 {
	if o == nil || o.FanoutConnTimeMs == nil {
		var ret int64
		return ret
	}
	return *o.FanoutConnTimeMs
}

// GetFanoutConnTimeMsOk returns a tuple with the FanoutConnTimeMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetFanoutConnTimeMsOk() (*int64, bool) {
	if o == nil || o.FanoutConnTimeMs == nil {
		return nil, false
	}
	return o.FanoutConnTimeMs, true
}

// HasFanoutConnTimeMs returns a boolean if a field has been set.
func (o *PlatformValues) HasFanoutConnTimeMs() bool {
	if o != nil && o.FanoutConnTimeMs != nil {
		return true
	}

	return false
}

// SetFanoutConnTimeMs gets a reference to the given int64 and assigns it to the FanoutConnTimeMs field.
func (o *PlatformValues) SetFanoutConnTimeMs(v int64) {
	o.FanoutConnTimeMs = &v
}

// GetDdosActionLimitStreamsConnections returns the DdosActionLimitStreamsConnections field value if set, zero value otherwise.
func (o *PlatformValues) GetDdosActionLimitStreamsConnections() int64 {
	if o == nil || o.DdosActionLimitStreamsConnections == nil {
		var ret int64
		return ret
	}
	return *o.DdosActionLimitStreamsConnections
}

// GetDdosActionLimitStreamsConnectionsOk returns a tuple with the DdosActionLimitStreamsConnections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetDdosActionLimitStreamsConnectionsOk() (*int64, bool) {
	if o == nil || o.DdosActionLimitStreamsConnections == nil {
		return nil, false
	}
	return o.DdosActionLimitStreamsConnections, true
}

// HasDdosActionLimitStreamsConnections returns a boolean if a field has been set.
func (o *PlatformValues) HasDdosActionLimitStreamsConnections() bool {
	if o != nil && o.DdosActionLimitStreamsConnections != nil {
		return true
	}

	return false
}

// SetDdosActionLimitStreamsConnections gets a reference to the given int64 and assigns it to the DdosActionLimitStreamsConnections field.
func (o *PlatformValues) SetDdosActionLimitStreamsConnections(v int64) {
	o.DdosActionLimitStreamsConnections = &v
}

// GetDdosActionLimitStreamsRequests returns the DdosActionLimitStreamsRequests field value if set, zero value otherwise.
func (o *PlatformValues) GetDdosActionLimitStreamsRequests() int64 {
	if o == nil || o.DdosActionLimitStreamsRequests == nil {
		var ret int64
		return ret
	}
	return *o.DdosActionLimitStreamsRequests
}

// GetDdosActionLimitStreamsRequestsOk returns a tuple with the DdosActionLimitStreamsRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetDdosActionLimitStreamsRequestsOk() (*int64, bool) {
	if o == nil || o.DdosActionLimitStreamsRequests == nil {
		return nil, false
	}
	return o.DdosActionLimitStreamsRequests, true
}

// HasDdosActionLimitStreamsRequests returns a boolean if a field has been set.
func (o *PlatformValues) HasDdosActionLimitStreamsRequests() bool {
	if o != nil && o.DdosActionLimitStreamsRequests != nil {
		return true
	}

	return false
}

// SetDdosActionLimitStreamsRequests gets a reference to the given int64 and assigns it to the DdosActionLimitStreamsRequests field.
func (o *PlatformValues) SetDdosActionLimitStreamsRequests(v int64) {
	o.DdosActionLimitStreamsRequests = &v
}

// GetDdosActionTarpitAccept returns the DdosActionTarpitAccept field value if set, zero value otherwise.
func (o *PlatformValues) GetDdosActionTarpitAccept() int64 {
	if o == nil || o.DdosActionTarpitAccept == nil {
		var ret int64
		return ret
	}
	return *o.DdosActionTarpitAccept
}

// GetDdosActionTarpitAcceptOk returns a tuple with the DdosActionTarpitAccept field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetDdosActionTarpitAcceptOk() (*int64, bool) {
	if o == nil || o.DdosActionTarpitAccept == nil {
		return nil, false
	}
	return o.DdosActionTarpitAccept, true
}

// HasDdosActionTarpitAccept returns a boolean if a field has been set.
func (o *PlatformValues) HasDdosActionTarpitAccept() bool {
	if o != nil && o.DdosActionTarpitAccept != nil {
		return true
	}

	return false
}

// SetDdosActionTarpitAccept gets a reference to the given int64 and assigns it to the DdosActionTarpitAccept field.
func (o *PlatformValues) SetDdosActionTarpitAccept(v int64) {
	o.DdosActionTarpitAccept = &v
}

// GetDdosActionTarpit returns the DdosActionTarpit field value if set, zero value otherwise.
func (o *PlatformValues) GetDdosActionTarpit() int64 {
	if o == nil || o.DdosActionTarpit == nil {
		var ret int64
		return ret
	}
	return *o.DdosActionTarpit
}

// GetDdosActionTarpitOk returns a tuple with the DdosActionTarpit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetDdosActionTarpitOk() (*int64, bool) {
	if o == nil || o.DdosActionTarpit == nil {
		return nil, false
	}
	return o.DdosActionTarpit, true
}

// HasDdosActionTarpit returns a boolean if a field has been set.
func (o *PlatformValues) HasDdosActionTarpit() bool {
	if o != nil && o.DdosActionTarpit != nil {
		return true
	}

	return false
}

// SetDdosActionTarpit gets a reference to the given int64 and assigns it to the DdosActionTarpit field.
func (o *PlatformValues) SetDdosActionTarpit(v int64) {
	o.DdosActionTarpit = &v
}

// GetDdosActionClose returns the DdosActionClose field value if set, zero value otherwise.
func (o *PlatformValues) GetDdosActionClose() int64 {
	if o == nil || o.DdosActionClose == nil {
		var ret int64
		return ret
	}
	return *o.DdosActionClose
}

// GetDdosActionCloseOk returns a tuple with the DdosActionClose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetDdosActionCloseOk() (*int64, bool) {
	if o == nil || o.DdosActionClose == nil {
		return nil, false
	}
	return o.DdosActionClose, true
}

// HasDdosActionClose returns a boolean if a field has been set.
func (o *PlatformValues) HasDdosActionClose() bool {
	if o != nil && o.DdosActionClose != nil {
		return true
	}

	return false
}

// SetDdosActionClose gets a reference to the given int64 and assigns it to the DdosActionClose field.
func (o *PlatformValues) SetDdosActionClose(v int64) {
	o.DdosActionClose = &v
}

// GetDdosActionBlackhole returns the DdosActionBlackhole field value if set, zero value otherwise.
func (o *PlatformValues) GetDdosActionBlackhole() int64 {
	if o == nil || o.DdosActionBlackhole == nil {
		var ret int64
		return ret
	}
	return *o.DdosActionBlackhole
}

// GetDdosActionBlackholeOk returns a tuple with the DdosActionBlackhole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetDdosActionBlackholeOk() (*int64, bool) {
	if o == nil || o.DdosActionBlackhole == nil {
		return nil, false
	}
	return o.DdosActionBlackhole, true
}

// HasDdosActionBlackhole returns a boolean if a field has been set.
func (o *PlatformValues) HasDdosActionBlackhole() bool {
	if o != nil && o.DdosActionBlackhole != nil {
		return true
	}

	return false
}

// SetDdosActionBlackhole gets a reference to the given int64 and assigns it to the DdosActionBlackhole field.
func (o *PlatformValues) SetDdosActionBlackhole(v int64) {
	o.DdosActionBlackhole = &v
}

// GetBotChallengeStarts returns the BotChallengeStarts field value if set, zero value otherwise.
func (o *PlatformValues) GetBotChallengeStarts() int64 {
	if o == nil || o.BotChallengeStarts == nil {
		var ret int64
		return ret
	}
	return *o.BotChallengeStarts
}

// GetBotChallengeStartsOk returns a tuple with the BotChallengeStarts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetBotChallengeStartsOk() (*int64, bool) {
	if o == nil || o.BotChallengeStarts == nil {
		return nil, false
	}
	return o.BotChallengeStarts, true
}

// HasBotChallengeStarts returns a boolean if a field has been set.
func (o *PlatformValues) HasBotChallengeStarts() bool {
	if o != nil && o.BotChallengeStarts != nil {
		return true
	}

	return false
}

// SetBotChallengeStarts gets a reference to the given int64 and assigns it to the BotChallengeStarts field.
func (o *PlatformValues) SetBotChallengeStarts(v int64) {
	o.BotChallengeStarts = &v
}

// GetBotChallengeCompleteTokensPassed returns the BotChallengeCompleteTokensPassed field value if set, zero value otherwise.
func (o *PlatformValues) GetBotChallengeCompleteTokensPassed() int64 {
	if o == nil || o.BotChallengeCompleteTokensPassed == nil {
		var ret int64
		return ret
	}
	return *o.BotChallengeCompleteTokensPassed
}

// GetBotChallengeCompleteTokensPassedOk returns a tuple with the BotChallengeCompleteTokensPassed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetBotChallengeCompleteTokensPassedOk() (*int64, bool) {
	if o == nil || o.BotChallengeCompleteTokensPassed == nil {
		return nil, false
	}
	return o.BotChallengeCompleteTokensPassed, true
}

// HasBotChallengeCompleteTokensPassed returns a boolean if a field has been set.
func (o *PlatformValues) HasBotChallengeCompleteTokensPassed() bool {
	if o != nil && o.BotChallengeCompleteTokensPassed != nil {
		return true
	}

	return false
}

// SetBotChallengeCompleteTokensPassed gets a reference to the given int64 and assigns it to the BotChallengeCompleteTokensPassed field.
func (o *PlatformValues) SetBotChallengeCompleteTokensPassed(v int64) {
	o.BotChallengeCompleteTokensPassed = &v
}

// GetBotChallengeCompleteTokensFailed returns the BotChallengeCompleteTokensFailed field value if set, zero value otherwise.
func (o *PlatformValues) GetBotChallengeCompleteTokensFailed() int64 {
	if o == nil || o.BotChallengeCompleteTokensFailed == nil {
		var ret int64
		return ret
	}
	return *o.BotChallengeCompleteTokensFailed
}

// GetBotChallengeCompleteTokensFailedOk returns a tuple with the BotChallengeCompleteTokensFailed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetBotChallengeCompleteTokensFailedOk() (*int64, bool) {
	if o == nil || o.BotChallengeCompleteTokensFailed == nil {
		return nil, false
	}
	return o.BotChallengeCompleteTokensFailed, true
}

// HasBotChallengeCompleteTokensFailed returns a boolean if a field has been set.
func (o *PlatformValues) HasBotChallengeCompleteTokensFailed() bool {
	if o != nil && o.BotChallengeCompleteTokensFailed != nil {
		return true
	}

	return false
}

// SetBotChallengeCompleteTokensFailed gets a reference to the given int64 and assigns it to the BotChallengeCompleteTokensFailed field.
func (o *PlatformValues) SetBotChallengeCompleteTokensFailed(v int64) {
	o.BotChallengeCompleteTokensFailed = &v
}

// GetBotChallengeCompleteTokensChecked returns the BotChallengeCompleteTokensChecked field value if set, zero value otherwise.
func (o *PlatformValues) GetBotChallengeCompleteTokensChecked() int64 {
	if o == nil || o.BotChallengeCompleteTokensChecked == nil {
		var ret int64
		return ret
	}
	return *o.BotChallengeCompleteTokensChecked
}

// GetBotChallengeCompleteTokensCheckedOk returns a tuple with the BotChallengeCompleteTokensChecked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetBotChallengeCompleteTokensCheckedOk() (*int64, bool) {
	if o == nil || o.BotChallengeCompleteTokensChecked == nil {
		return nil, false
	}
	return o.BotChallengeCompleteTokensChecked, true
}

// HasBotChallengeCompleteTokensChecked returns a boolean if a field has been set.
func (o *PlatformValues) HasBotChallengeCompleteTokensChecked() bool {
	if o != nil && o.BotChallengeCompleteTokensChecked != nil {
		return true
	}

	return false
}

// SetBotChallengeCompleteTokensChecked gets a reference to the given int64 and assigns it to the BotChallengeCompleteTokensChecked field.
func (o *PlatformValues) SetBotChallengeCompleteTokensChecked(v int64) {
	o.BotChallengeCompleteTokensChecked = &v
}

// GetBotChallengeCompleteTokensDisabled returns the BotChallengeCompleteTokensDisabled field value if set, zero value otherwise.
func (o *PlatformValues) GetBotChallengeCompleteTokensDisabled() int64 {
	if o == nil || o.BotChallengeCompleteTokensDisabled == nil {
		var ret int64
		return ret
	}
	return *o.BotChallengeCompleteTokensDisabled
}

// GetBotChallengeCompleteTokensDisabledOk returns a tuple with the BotChallengeCompleteTokensDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetBotChallengeCompleteTokensDisabledOk() (*int64, bool) {
	if o == nil || o.BotChallengeCompleteTokensDisabled == nil {
		return nil, false
	}
	return o.BotChallengeCompleteTokensDisabled, true
}

// HasBotChallengeCompleteTokensDisabled returns a boolean if a field has been set.
func (o *PlatformValues) HasBotChallengeCompleteTokensDisabled() bool {
	if o != nil && o.BotChallengeCompleteTokensDisabled != nil {
		return true
	}

	return false
}

// SetBotChallengeCompleteTokensDisabled gets a reference to the given int64 and assigns it to the BotChallengeCompleteTokensDisabled field.
func (o *PlatformValues) SetBotChallengeCompleteTokensDisabled(v int64) {
	o.BotChallengeCompleteTokensDisabled = &v
}

// GetBotChallengeCompleteTokensIssued returns the BotChallengeCompleteTokensIssued field value if set, zero value otherwise.
func (o *PlatformValues) GetBotChallengeCompleteTokensIssued() int64 {
	if o == nil || o.BotChallengeCompleteTokensIssued == nil {
		var ret int64
		return ret
	}
	return *o.BotChallengeCompleteTokensIssued
}

// GetBotChallengeCompleteTokensIssuedOk returns a tuple with the BotChallengeCompleteTokensIssued field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetBotChallengeCompleteTokensIssuedOk() (*int64, bool) {
	if o == nil || o.BotChallengeCompleteTokensIssued == nil {
		return nil, false
	}
	return o.BotChallengeCompleteTokensIssued, true
}

// HasBotChallengeCompleteTokensIssued returns a boolean if a field has been set.
func (o *PlatformValues) HasBotChallengeCompleteTokensIssued() bool {
	if o != nil && o.BotChallengeCompleteTokensIssued != nil {
		return true
	}

	return false
}

// SetBotChallengeCompleteTokensIssued gets a reference to the given int64 and assigns it to the BotChallengeCompleteTokensIssued field.
func (o *PlatformValues) SetBotChallengeCompleteTokensIssued(v int64) {
	o.BotChallengeCompleteTokensIssued = &v
}

// GetBotChallengesIssued returns the BotChallengesIssued field value if set, zero value otherwise.
func (o *PlatformValues) GetBotChallengesIssued() int64 {
	if o == nil || o.BotChallengesIssued == nil {
		var ret int64
		return ret
	}
	return *o.BotChallengesIssued
}

// GetBotChallengesIssuedOk returns a tuple with the BotChallengesIssued field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetBotChallengesIssuedOk() (*int64, bool) {
	if o == nil || o.BotChallengesIssued == nil {
		return nil, false
	}
	return o.BotChallengesIssued, true
}

// HasBotChallengesIssued returns a boolean if a field has been set.
func (o *PlatformValues) HasBotChallengesIssued() bool {
	if o != nil && o.BotChallengesIssued != nil {
		return true
	}

	return false
}

// SetBotChallengesIssued gets a reference to the given int64 and assigns it to the BotChallengesIssued field.
func (o *PlatformValues) SetBotChallengesIssued(v int64) {
	o.BotChallengesIssued = &v
}

// GetBotChallengesSucceeded returns the BotChallengesSucceeded field value if set, zero value otherwise.
func (o *PlatformValues) GetBotChallengesSucceeded() int64 {
	if o == nil || o.BotChallengesSucceeded == nil {
		var ret int64
		return ret
	}
	return *o.BotChallengesSucceeded
}

// GetBotChallengesSucceededOk returns a tuple with the BotChallengesSucceeded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetBotChallengesSucceededOk() (*int64, bool) {
	if o == nil || o.BotChallengesSucceeded == nil {
		return nil, false
	}
	return o.BotChallengesSucceeded, true
}

// HasBotChallengesSucceeded returns a boolean if a field has been set.
func (o *PlatformValues) HasBotChallengesSucceeded() bool {
	if o != nil && o.BotChallengesSucceeded != nil {
		return true
	}

	return false
}

// SetBotChallengesSucceeded gets a reference to the given int64 and assigns it to the BotChallengesSucceeded field.
func (o *PlatformValues) SetBotChallengesSucceeded(v int64) {
	o.BotChallengesSucceeded = &v
}

// GetBotChallengesFailed returns the BotChallengesFailed field value if set, zero value otherwise.
func (o *PlatformValues) GetBotChallengesFailed() int64 {
	if o == nil || o.BotChallengesFailed == nil {
		var ret int64
		return ret
	}
	return *o.BotChallengesFailed
}

// GetBotChallengesFailedOk returns a tuple with the BotChallengesFailed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetBotChallengesFailedOk() (*int64, bool) {
	if o == nil || o.BotChallengesFailed == nil {
		return nil, false
	}
	return o.BotChallengesFailed, true
}

// HasBotChallengesFailed returns a boolean if a field has been set.
func (o *PlatformValues) HasBotChallengesFailed() bool {
	if o != nil && o.BotChallengesFailed != nil {
		return true
	}

	return false
}

// SetBotChallengesFailed gets a reference to the given int64 and assigns it to the BotChallengesFailed field.
func (o *PlatformValues) SetBotChallengesFailed(v int64) {
	o.BotChallengesFailed = &v
}

// GetDdosActionDowngrade returns the DdosActionDowngrade field value if set, zero value otherwise.
func (o *PlatformValues) GetDdosActionDowngrade() int64 {
	if o == nil || o.DdosActionDowngrade == nil {
		var ret int64
		return ret
	}
	return *o.DdosActionDowngrade
}

// GetDdosActionDowngradeOk returns a tuple with the DdosActionDowngrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetDdosActionDowngradeOk() (*int64, bool) {
	if o == nil || o.DdosActionDowngrade == nil {
		return nil, false
	}
	return o.DdosActionDowngrade, true
}

// HasDdosActionDowngrade returns a boolean if a field has been set.
func (o *PlatformValues) HasDdosActionDowngrade() bool {
	if o != nil && o.DdosActionDowngrade != nil {
		return true
	}

	return false
}

// SetDdosActionDowngrade gets a reference to the given int64 and assigns it to the DdosActionDowngrade field.
func (o *PlatformValues) SetDdosActionDowngrade(v int64) {
	o.DdosActionDowngrade = &v
}

// GetDdosActionDowngradedConnections returns the DdosActionDowngradedConnections field value if set, zero value otherwise.
func (o *PlatformValues) GetDdosActionDowngradedConnections() int64 {
	if o == nil || o.DdosActionDowngradedConnections == nil {
		var ret int64
		return ret
	}
	return *o.DdosActionDowngradedConnections
}

// GetDdosActionDowngradedConnectionsOk returns a tuple with the DdosActionDowngradedConnections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetDdosActionDowngradedConnectionsOk() (*int64, bool) {
	if o == nil || o.DdosActionDowngradedConnections == nil {
		return nil, false
	}
	return o.DdosActionDowngradedConnections, true
}

// HasDdosActionDowngradedConnections returns a boolean if a field has been set.
func (o *PlatformValues) HasDdosActionDowngradedConnections() bool {
	if o != nil && o.DdosActionDowngradedConnections != nil {
		return true
	}

	return false
}

// SetDdosActionDowngradedConnections gets a reference to the given int64 and assigns it to the DdosActionDowngradedConnections field.
func (o *PlatformValues) SetDdosActionDowngradedConnections(v int64) {
	o.DdosActionDowngradedConnections = &v
}

// GetAllHitRequests returns the AllHitRequests field value if set, zero value otherwise.
func (o *PlatformValues) GetAllHitRequests() int64 {
	if o == nil || o.AllHitRequests == nil {
		var ret int64
		return ret
	}
	return *o.AllHitRequests
}

// GetAllHitRequestsOk returns a tuple with the AllHitRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetAllHitRequestsOk() (*int64, bool) {
	if o == nil || o.AllHitRequests == nil {
		return nil, false
	}
	return o.AllHitRequests, true
}

// HasAllHitRequests returns a boolean if a field has been set.
func (o *PlatformValues) HasAllHitRequests() bool {
	if o != nil && o.AllHitRequests != nil {
		return true
	}

	return false
}

// SetAllHitRequests gets a reference to the given int64 and assigns it to the AllHitRequests field.
func (o *PlatformValues) SetAllHitRequests(v int64) {
	o.AllHitRequests = &v
}

// GetAllMissRequests returns the AllMissRequests field value if set, zero value otherwise.
func (o *PlatformValues) GetAllMissRequests() int64 {
	if o == nil || o.AllMissRequests == nil {
		var ret int64
		return ret
	}
	return *o.AllMissRequests
}

// GetAllMissRequestsOk returns a tuple with the AllMissRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetAllMissRequestsOk() (*int64, bool) {
	if o == nil || o.AllMissRequests == nil {
		return nil, false
	}
	return o.AllMissRequests, true
}

// HasAllMissRequests returns a boolean if a field has been set.
func (o *PlatformValues) HasAllMissRequests() bool {
	if o != nil && o.AllMissRequests != nil {
		return true
	}

	return false
}

// SetAllMissRequests gets a reference to the given int64 and assigns it to the AllMissRequests field.
func (o *PlatformValues) SetAllMissRequests(v int64) {
	o.AllMissRequests = &v
}

// GetAllPassRequests returns the AllPassRequests field value if set, zero value otherwise.
func (o *PlatformValues) GetAllPassRequests() int64 {
	if o == nil || o.AllPassRequests == nil {
		var ret int64
		return ret
	}
	return *o.AllPassRequests
}

// GetAllPassRequestsOk returns a tuple with the AllPassRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetAllPassRequestsOk() (*int64, bool) {
	if o == nil || o.AllPassRequests == nil {
		return nil, false
	}
	return o.AllPassRequests, true
}

// HasAllPassRequests returns a boolean if a field has been set.
func (o *PlatformValues) HasAllPassRequests() bool {
	if o != nil && o.AllPassRequests != nil {
		return true
	}

	return false
}

// SetAllPassRequests gets a reference to the given int64 and assigns it to the AllPassRequests field.
func (o *PlatformValues) SetAllPassRequests(v int64) {
	o.AllPassRequests = &v
}

// GetAllErrorRequests returns the AllErrorRequests field value if set, zero value otherwise.
func (o *PlatformValues) GetAllErrorRequests() int64 {
	if o == nil || o.AllErrorRequests == nil {
		var ret int64
		return ret
	}
	return *o.AllErrorRequests
}

// GetAllErrorRequestsOk returns a tuple with the AllErrorRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetAllErrorRequestsOk() (*int64, bool) {
	if o == nil || o.AllErrorRequests == nil {
		return nil, false
	}
	return o.AllErrorRequests, true
}

// HasAllErrorRequests returns a boolean if a field has been set.
func (o *PlatformValues) HasAllErrorRequests() bool {
	if o != nil && o.AllErrorRequests != nil {
		return true
	}

	return false
}

// SetAllErrorRequests gets a reference to the given int64 and assigns it to the AllErrorRequests field.
func (o *PlatformValues) SetAllErrorRequests(v int64) {
	o.AllErrorRequests = &v
}

// GetAllSynthRequests returns the AllSynthRequests field value if set, zero value otherwise.
func (o *PlatformValues) GetAllSynthRequests() int64 {
	if o == nil || o.AllSynthRequests == nil {
		var ret int64
		return ret
	}
	return *o.AllSynthRequests
}

// GetAllSynthRequestsOk returns a tuple with the AllSynthRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetAllSynthRequestsOk() (*int64, bool) {
	if o == nil || o.AllSynthRequests == nil {
		return nil, false
	}
	return o.AllSynthRequests, true
}

// HasAllSynthRequests returns a boolean if a field has been set.
func (o *PlatformValues) HasAllSynthRequests() bool {
	if o != nil && o.AllSynthRequests != nil {
		return true
	}

	return false
}

// SetAllSynthRequests gets a reference to the given int64 and assigns it to the AllSynthRequests field.
func (o *PlatformValues) SetAllSynthRequests(v int64) {
	o.AllSynthRequests = &v
}

// GetAllEdgeHitRequests returns the AllEdgeHitRequests field value if set, zero value otherwise.
func (o *PlatformValues) GetAllEdgeHitRequests() int64 {
	if o == nil || o.AllEdgeHitRequests == nil {
		var ret int64
		return ret
	}
	return *o.AllEdgeHitRequests
}

// GetAllEdgeHitRequestsOk returns a tuple with the AllEdgeHitRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetAllEdgeHitRequestsOk() (*int64, bool) {
	if o == nil || o.AllEdgeHitRequests == nil {
		return nil, false
	}
	return o.AllEdgeHitRequests, true
}

// HasAllEdgeHitRequests returns a boolean if a field has been set.
func (o *PlatformValues) HasAllEdgeHitRequests() bool {
	if o != nil && o.AllEdgeHitRequests != nil {
		return true
	}

	return false
}

// SetAllEdgeHitRequests gets a reference to the given int64 and assigns it to the AllEdgeHitRequests field.
func (o *PlatformValues) SetAllEdgeHitRequests(v int64) {
	o.AllEdgeHitRequests = &v
}

// GetAllEdgeMissRequests returns the AllEdgeMissRequests field value if set, zero value otherwise.
func (o *PlatformValues) GetAllEdgeMissRequests() int64 {
	if o == nil || o.AllEdgeMissRequests == nil {
		var ret int64
		return ret
	}
	return *o.AllEdgeMissRequests
}

// GetAllEdgeMissRequestsOk returns a tuple with the AllEdgeMissRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetAllEdgeMissRequestsOk() (*int64, bool) {
	if o == nil || o.AllEdgeMissRequests == nil {
		return nil, false
	}
	return o.AllEdgeMissRequests, true
}

// HasAllEdgeMissRequests returns a boolean if a field has been set.
func (o *PlatformValues) HasAllEdgeMissRequests() bool {
	if o != nil && o.AllEdgeMissRequests != nil {
		return true
	}

	return false
}

// SetAllEdgeMissRequests gets a reference to the given int64 and assigns it to the AllEdgeMissRequests field.
func (o *PlatformValues) SetAllEdgeMissRequests(v int64) {
	o.AllEdgeMissRequests = &v
}

// GetAllStatus1xx returns the AllStatus1xx field value if set, zero value otherwise.
func (o *PlatformValues) GetAllStatus1xx() int64 {
	if o == nil || o.AllStatus1xx == nil {
		var ret int64
		return ret
	}
	return *o.AllStatus1xx
}

// GetAllStatus1xxOk returns a tuple with the AllStatus1xx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetAllStatus1xxOk() (*int64, bool) {
	if o == nil || o.AllStatus1xx == nil {
		return nil, false
	}
	return o.AllStatus1xx, true
}

// HasAllStatus1xx returns a boolean if a field has been set.
func (o *PlatformValues) HasAllStatus1xx() bool {
	if o != nil && o.AllStatus1xx != nil {
		return true
	}

	return false
}

// SetAllStatus1xx gets a reference to the given int64 and assigns it to the AllStatus1xx field.
func (o *PlatformValues) SetAllStatus1xx(v int64) {
	o.AllStatus1xx = &v
}

// GetAllStatus2xx returns the AllStatus2xx field value if set, zero value otherwise.
func (o *PlatformValues) GetAllStatus2xx() int64 {
	if o == nil || o.AllStatus2xx == nil {
		var ret int64
		return ret
	}
	return *o.AllStatus2xx
}

// GetAllStatus2xxOk returns a tuple with the AllStatus2xx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetAllStatus2xxOk() (*int64, bool) {
	if o == nil || o.AllStatus2xx == nil {
		return nil, false
	}
	return o.AllStatus2xx, true
}

// HasAllStatus2xx returns a boolean if a field has been set.
func (o *PlatformValues) HasAllStatus2xx() bool {
	if o != nil && o.AllStatus2xx != nil {
		return true
	}

	return false
}

// SetAllStatus2xx gets a reference to the given int64 and assigns it to the AllStatus2xx field.
func (o *PlatformValues) SetAllStatus2xx(v int64) {
	o.AllStatus2xx = &v
}

// GetAllStatus3xx returns the AllStatus3xx field value if set, zero value otherwise.
func (o *PlatformValues) GetAllStatus3xx() int64 {
	if o == nil || o.AllStatus3xx == nil {
		var ret int64
		return ret
	}
	return *o.AllStatus3xx
}

// GetAllStatus3xxOk returns a tuple with the AllStatus3xx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetAllStatus3xxOk() (*int64, bool) {
	if o == nil || o.AllStatus3xx == nil {
		return nil, false
	}
	return o.AllStatus3xx, true
}

// HasAllStatus3xx returns a boolean if a field has been set.
func (o *PlatformValues) HasAllStatus3xx() bool {
	if o != nil && o.AllStatus3xx != nil {
		return true
	}

	return false
}

// SetAllStatus3xx gets a reference to the given int64 and assigns it to the AllStatus3xx field.
func (o *PlatformValues) SetAllStatus3xx(v int64) {
	o.AllStatus3xx = &v
}

// GetAllStatus4xx returns the AllStatus4xx field value if set, zero value otherwise.
func (o *PlatformValues) GetAllStatus4xx() int64 {
	if o == nil || o.AllStatus4xx == nil {
		var ret int64
		return ret
	}
	return *o.AllStatus4xx
}

// GetAllStatus4xxOk returns a tuple with the AllStatus4xx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetAllStatus4xxOk() (*int64, bool) {
	if o == nil || o.AllStatus4xx == nil {
		return nil, false
	}
	return o.AllStatus4xx, true
}

// HasAllStatus4xx returns a boolean if a field has been set.
func (o *PlatformValues) HasAllStatus4xx() bool {
	if o != nil && o.AllStatus4xx != nil {
		return true
	}

	return false
}

// SetAllStatus4xx gets a reference to the given int64 and assigns it to the AllStatus4xx field.
func (o *PlatformValues) SetAllStatus4xx(v int64) {
	o.AllStatus4xx = &v
}

// GetAllStatus5xx returns the AllStatus5xx field value if set, zero value otherwise.
func (o *PlatformValues) GetAllStatus5xx() int64 {
	if o == nil || o.AllStatus5xx == nil {
		var ret int64
		return ret
	}
	return *o.AllStatus5xx
}

// GetAllStatus5xxOk returns a tuple with the AllStatus5xx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetAllStatus5xxOk() (*int64, bool) {
	if o == nil || o.AllStatus5xx == nil {
		return nil, false
	}
	return o.AllStatus5xx, true
}

// HasAllStatus5xx returns a boolean if a field has been set.
func (o *PlatformValues) HasAllStatus5xx() bool {
	if o != nil && o.AllStatus5xx != nil {
		return true
	}

	return false
}

// SetAllStatus5xx gets a reference to the given int64 and assigns it to the AllStatus5xx field.
func (o *PlatformValues) SetAllStatus5xx(v int64) {
	o.AllStatus5xx = &v
}

// GetOriginOffload returns the OriginOffload field value if set, zero value otherwise.
func (o *PlatformValues) GetOriginOffload() float32 {
	if o == nil || o.OriginOffload == nil {
		var ret float32
		return ret
	}
	return *o.OriginOffload
}

// GetOriginOffloadOk returns a tuple with the OriginOffload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetOriginOffloadOk() (*float32, bool) {
	if o == nil || o.OriginOffload == nil {
		return nil, false
	}
	return o.OriginOffload, true
}

// HasOriginOffload returns a boolean if a field has been set.
func (o *PlatformValues) HasOriginOffload() bool {
	if o != nil && o.OriginOffload != nil {
		return true
	}

	return false
}

// SetOriginOffload gets a reference to the given float32 and assigns it to the OriginOffload field.
func (o *PlatformValues) SetOriginOffload(v float32) {
	o.OriginOffload = &v
}

// GetRequestDeniedGetHeadBody returns the RequestDeniedGetHeadBody field value if set, zero value otherwise.
func (o *PlatformValues) GetRequestDeniedGetHeadBody() int64 {
	if o == nil || o.RequestDeniedGetHeadBody == nil {
		var ret int64
		return ret
	}
	return *o.RequestDeniedGetHeadBody
}

// GetRequestDeniedGetHeadBodyOk returns a tuple with the RequestDeniedGetHeadBody field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetRequestDeniedGetHeadBodyOk() (*int64, bool) {
	if o == nil || o.RequestDeniedGetHeadBody == nil {
		return nil, false
	}
	return o.RequestDeniedGetHeadBody, true
}

// HasRequestDeniedGetHeadBody returns a boolean if a field has been set.
func (o *PlatformValues) HasRequestDeniedGetHeadBody() bool {
	if o != nil && o.RequestDeniedGetHeadBody != nil {
		return true
	}

	return false
}

// SetRequestDeniedGetHeadBody gets a reference to the given int64 and assigns it to the RequestDeniedGetHeadBody field.
func (o *PlatformValues) SetRequestDeniedGetHeadBody(v int64) {
	o.RequestDeniedGetHeadBody = &v
}

// GetDdosProtectionRequestsDetectCount returns the DdosProtectionRequestsDetectCount field value if set, zero value otherwise.
func (o *PlatformValues) GetDdosProtectionRequestsDetectCount() int64 {
	if o == nil || o.DdosProtectionRequestsDetectCount == nil {
		var ret int64
		return ret
	}
	return *o.DdosProtectionRequestsDetectCount
}

// GetDdosProtectionRequestsDetectCountOk returns a tuple with the DdosProtectionRequestsDetectCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetDdosProtectionRequestsDetectCountOk() (*int64, bool) {
	if o == nil || o.DdosProtectionRequestsDetectCount == nil {
		return nil, false
	}
	return o.DdosProtectionRequestsDetectCount, true
}

// HasDdosProtectionRequestsDetectCount returns a boolean if a field has been set.
func (o *PlatformValues) HasDdosProtectionRequestsDetectCount() bool {
	if o != nil && o.DdosProtectionRequestsDetectCount != nil {
		return true
	}

	return false
}

// SetDdosProtectionRequestsDetectCount gets a reference to the given int64 and assigns it to the DdosProtectionRequestsDetectCount field.
func (o *PlatformValues) SetDdosProtectionRequestsDetectCount(v int64) {
	o.DdosProtectionRequestsDetectCount = &v
}

// GetDdosProtectionRequestsMitigateCount returns the DdosProtectionRequestsMitigateCount field value if set, zero value otherwise.
func (o *PlatformValues) GetDdosProtectionRequestsMitigateCount() int64 {
	if o == nil || o.DdosProtectionRequestsMitigateCount == nil {
		var ret int64
		return ret
	}
	return *o.DdosProtectionRequestsMitigateCount
}

// GetDdosProtectionRequestsMitigateCountOk returns a tuple with the DdosProtectionRequestsMitigateCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetDdosProtectionRequestsMitigateCountOk() (*int64, bool) {
	if o == nil || o.DdosProtectionRequestsMitigateCount == nil {
		return nil, false
	}
	return o.DdosProtectionRequestsMitigateCount, true
}

// HasDdosProtectionRequestsMitigateCount returns a boolean if a field has been set.
func (o *PlatformValues) HasDdosProtectionRequestsMitigateCount() bool {
	if o != nil && o.DdosProtectionRequestsMitigateCount != nil {
		return true
	}

	return false
}

// SetDdosProtectionRequestsMitigateCount gets a reference to the given int64 and assigns it to the DdosProtectionRequestsMitigateCount field.
func (o *PlatformValues) SetDdosProtectionRequestsMitigateCount(v int64) {
	o.DdosProtectionRequestsMitigateCount = &v
}

// GetDdosProtectionRequestsAllowCount returns the DdosProtectionRequestsAllowCount field value if set, zero value otherwise.
func (o *PlatformValues) GetDdosProtectionRequestsAllowCount() int64 {
	if o == nil || o.DdosProtectionRequestsAllowCount == nil {
		var ret int64
		return ret
	}
	return *o.DdosProtectionRequestsAllowCount
}

// GetDdosProtectionRequestsAllowCountOk returns a tuple with the DdosProtectionRequestsAllowCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetDdosProtectionRequestsAllowCountOk() (*int64, bool) {
	if o == nil || o.DdosProtectionRequestsAllowCount == nil {
		return nil, false
	}
	return o.DdosProtectionRequestsAllowCount, true
}

// HasDdosProtectionRequestsAllowCount returns a boolean if a field has been set.
func (o *PlatformValues) HasDdosProtectionRequestsAllowCount() bool {
	if o != nil && o.DdosProtectionRequestsAllowCount != nil {
		return true
	}

	return false
}

// SetDdosProtectionRequestsAllowCount gets a reference to the given int64 and assigns it to the DdosProtectionRequestsAllowCount field.
func (o *PlatformValues) SetDdosProtectionRequestsAllowCount(v int64) {
	o.DdosProtectionRequestsAllowCount = &v
}

// GetObjectStorageClassAOperationsCount returns the ObjectStorageClassAOperationsCount field value if set, zero value otherwise.
func (o *PlatformValues) GetObjectStorageClassAOperationsCount() int64 {
	if o == nil || o.ObjectStorageClassAOperationsCount == nil {
		var ret int64
		return ret
	}
	return *o.ObjectStorageClassAOperationsCount
}

// GetObjectStorageClassAOperationsCountOk returns a tuple with the ObjectStorageClassAOperationsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetObjectStorageClassAOperationsCountOk() (*int64, bool) {
	if o == nil || o.ObjectStorageClassAOperationsCount == nil {
		return nil, false
	}
	return o.ObjectStorageClassAOperationsCount, true
}

// HasObjectStorageClassAOperationsCount returns a boolean if a field has been set.
func (o *PlatformValues) HasObjectStorageClassAOperationsCount() bool {
	if o != nil && o.ObjectStorageClassAOperationsCount != nil {
		return true
	}

	return false
}

// SetObjectStorageClassAOperationsCount gets a reference to the given int64 and assigns it to the ObjectStorageClassAOperationsCount field.
func (o *PlatformValues) SetObjectStorageClassAOperationsCount(v int64) {
	o.ObjectStorageClassAOperationsCount = &v
}

// GetObjectStorageClassBOperationsCount returns the ObjectStorageClassBOperationsCount field value if set, zero value otherwise.
func (o *PlatformValues) GetObjectStorageClassBOperationsCount() int64 {
	if o == nil || o.ObjectStorageClassBOperationsCount == nil {
		var ret int64
		return ret
	}
	return *o.ObjectStorageClassBOperationsCount
}

// GetObjectStorageClassBOperationsCountOk returns a tuple with the ObjectStorageClassBOperationsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetObjectStorageClassBOperationsCountOk() (*int64, bool) {
	if o == nil || o.ObjectStorageClassBOperationsCount == nil {
		return nil, false
	}
	return o.ObjectStorageClassBOperationsCount, true
}

// HasObjectStorageClassBOperationsCount returns a boolean if a field has been set.
func (o *PlatformValues) HasObjectStorageClassBOperationsCount() bool {
	if o != nil && o.ObjectStorageClassBOperationsCount != nil {
		return true
	}

	return false
}

// SetObjectStorageClassBOperationsCount gets a reference to the given int64 and assigns it to the ObjectStorageClassBOperationsCount field.
func (o *PlatformValues) SetObjectStorageClassBOperationsCount(v int64) {
	o.ObjectStorageClassBOperationsCount = &v
}

// GetAiaRequests returns the AiaRequests field value if set, zero value otherwise.
func (o *PlatformValues) GetAiaRequests() int64 {
	if o == nil || o.AiaRequests == nil {
		var ret int64
		return ret
	}
	return *o.AiaRequests
}

// GetAiaRequestsOk returns a tuple with the AiaRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetAiaRequestsOk() (*int64, bool) {
	if o == nil || o.AiaRequests == nil {
		return nil, false
	}
	return o.AiaRequests, true
}

// HasAiaRequests returns a boolean if a field has been set.
func (o *PlatformValues) HasAiaRequests() bool {
	if o != nil && o.AiaRequests != nil {
		return true
	}

	return false
}

// SetAiaRequests gets a reference to the given int64 and assigns it to the AiaRequests field.
func (o *PlatformValues) SetAiaRequests(v int64) {
	o.AiaRequests = &v
}

// GetAiaStatus1xx returns the AiaStatus1xx field value if set, zero value otherwise.
func (o *PlatformValues) GetAiaStatus1xx() int64 {
	if o == nil || o.AiaStatus1xx == nil {
		var ret int64
		return ret
	}
	return *o.AiaStatus1xx
}

// GetAiaStatus1xxOk returns a tuple with the AiaStatus1xx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetAiaStatus1xxOk() (*int64, bool) {
	if o == nil || o.AiaStatus1xx == nil {
		return nil, false
	}
	return o.AiaStatus1xx, true
}

// HasAiaStatus1xx returns a boolean if a field has been set.
func (o *PlatformValues) HasAiaStatus1xx() bool {
	if o != nil && o.AiaStatus1xx != nil {
		return true
	}

	return false
}

// SetAiaStatus1xx gets a reference to the given int64 and assigns it to the AiaStatus1xx field.
func (o *PlatformValues) SetAiaStatus1xx(v int64) {
	o.AiaStatus1xx = &v
}

// GetAiaStatus2xx returns the AiaStatus2xx field value if set, zero value otherwise.
func (o *PlatformValues) GetAiaStatus2xx() int64 {
	if o == nil || o.AiaStatus2xx == nil {
		var ret int64
		return ret
	}
	return *o.AiaStatus2xx
}

// GetAiaStatus2xxOk returns a tuple with the AiaStatus2xx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetAiaStatus2xxOk() (*int64, bool) {
	if o == nil || o.AiaStatus2xx == nil {
		return nil, false
	}
	return o.AiaStatus2xx, true
}

// HasAiaStatus2xx returns a boolean if a field has been set.
func (o *PlatformValues) HasAiaStatus2xx() bool {
	if o != nil && o.AiaStatus2xx != nil {
		return true
	}

	return false
}

// SetAiaStatus2xx gets a reference to the given int64 and assigns it to the AiaStatus2xx field.
func (o *PlatformValues) SetAiaStatus2xx(v int64) {
	o.AiaStatus2xx = &v
}

// GetAiaStatus3xx returns the AiaStatus3xx field value if set, zero value otherwise.
func (o *PlatformValues) GetAiaStatus3xx() int64 {
	if o == nil || o.AiaStatus3xx == nil {
		var ret int64
		return ret
	}
	return *o.AiaStatus3xx
}

// GetAiaStatus3xxOk returns a tuple with the AiaStatus3xx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetAiaStatus3xxOk() (*int64, bool) {
	if o == nil || o.AiaStatus3xx == nil {
		return nil, false
	}
	return o.AiaStatus3xx, true
}

// HasAiaStatus3xx returns a boolean if a field has been set.
func (o *PlatformValues) HasAiaStatus3xx() bool {
	if o != nil && o.AiaStatus3xx != nil {
		return true
	}

	return false
}

// SetAiaStatus3xx gets a reference to the given int64 and assigns it to the AiaStatus3xx field.
func (o *PlatformValues) SetAiaStatus3xx(v int64) {
	o.AiaStatus3xx = &v
}

// GetAiaStatus4xx returns the AiaStatus4xx field value if set, zero value otherwise.
func (o *PlatformValues) GetAiaStatus4xx() int64 {
	if o == nil || o.AiaStatus4xx == nil {
		var ret int64
		return ret
	}
	return *o.AiaStatus4xx
}

// GetAiaStatus4xxOk returns a tuple with the AiaStatus4xx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetAiaStatus4xxOk() (*int64, bool) {
	if o == nil || o.AiaStatus4xx == nil {
		return nil, false
	}
	return o.AiaStatus4xx, true
}

// HasAiaStatus4xx returns a boolean if a field has been set.
func (o *PlatformValues) HasAiaStatus4xx() bool {
	if o != nil && o.AiaStatus4xx != nil {
		return true
	}

	return false
}

// SetAiaStatus4xx gets a reference to the given int64 and assigns it to the AiaStatus4xx field.
func (o *PlatformValues) SetAiaStatus4xx(v int64) {
	o.AiaStatus4xx = &v
}

// GetAiaStatus5xx returns the AiaStatus5xx field value if set, zero value otherwise.
func (o *PlatformValues) GetAiaStatus5xx() int64 {
	if o == nil || o.AiaStatus5xx == nil {
		var ret int64
		return ret
	}
	return *o.AiaStatus5xx
}

// GetAiaStatus5xxOk returns a tuple with the AiaStatus5xx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetAiaStatus5xxOk() (*int64, bool) {
	if o == nil || o.AiaStatus5xx == nil {
		return nil, false
	}
	return o.AiaStatus5xx, true
}

// HasAiaStatus5xx returns a boolean if a field has been set.
func (o *PlatformValues) HasAiaStatus5xx() bool {
	if o != nil && o.AiaStatus5xx != nil {
		return true
	}

	return false
}

// SetAiaStatus5xx gets a reference to the given int64 and assigns it to the AiaStatus5xx field.
func (o *PlatformValues) SetAiaStatus5xx(v int64) {
	o.AiaStatus5xx = &v
}

// GetAiaResponseUsageTokens returns the AiaResponseUsageTokens field value if set, zero value otherwise.
func (o *PlatformValues) GetAiaResponseUsageTokens() int64 {
	if o == nil || o.AiaResponseUsageTokens == nil {
		var ret int64
		return ret
	}
	return *o.AiaResponseUsageTokens
}

// GetAiaResponseUsageTokensOk returns a tuple with the AiaResponseUsageTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetAiaResponseUsageTokensOk() (*int64, bool) {
	if o == nil || o.AiaResponseUsageTokens == nil {
		return nil, false
	}
	return o.AiaResponseUsageTokens, true
}

// HasAiaResponseUsageTokens returns a boolean if a field has been set.
func (o *PlatformValues) HasAiaResponseUsageTokens() bool {
	if o != nil && o.AiaResponseUsageTokens != nil {
		return true
	}

	return false
}

// SetAiaResponseUsageTokens gets a reference to the given int64 and assigns it to the AiaResponseUsageTokens field.
func (o *PlatformValues) SetAiaResponseUsageTokens(v int64) {
	o.AiaResponseUsageTokens = &v
}

// GetAiaOriginUsageTokens returns the AiaOriginUsageTokens field value if set, zero value otherwise.
func (o *PlatformValues) GetAiaOriginUsageTokens() int64 {
	if o == nil || o.AiaOriginUsageTokens == nil {
		var ret int64
		return ret
	}
	return *o.AiaOriginUsageTokens
}

// GetAiaOriginUsageTokensOk returns a tuple with the AiaOriginUsageTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetAiaOriginUsageTokensOk() (*int64, bool) {
	if o == nil || o.AiaOriginUsageTokens == nil {
		return nil, false
	}
	return o.AiaOriginUsageTokens, true
}

// HasAiaOriginUsageTokens returns a boolean if a field has been set.
func (o *PlatformValues) HasAiaOriginUsageTokens() bool {
	if o != nil && o.AiaOriginUsageTokens != nil {
		return true
	}

	return false
}

// SetAiaOriginUsageTokens gets a reference to the given int64 and assigns it to the AiaOriginUsageTokens field.
func (o *PlatformValues) SetAiaOriginUsageTokens(v int64) {
	o.AiaOriginUsageTokens = &v
}

// GetAiaEstimatedTimeSavedMs returns the AiaEstimatedTimeSavedMs field value if set, zero value otherwise.
func (o *PlatformValues) GetAiaEstimatedTimeSavedMs() int64 {
	if o == nil || o.AiaEstimatedTimeSavedMs == nil {
		var ret int64
		return ret
	}
	return *o.AiaEstimatedTimeSavedMs
}

// GetAiaEstimatedTimeSavedMsOk returns a tuple with the AiaEstimatedTimeSavedMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetAiaEstimatedTimeSavedMsOk() (*int64, bool) {
	if o == nil || o.AiaEstimatedTimeSavedMs == nil {
		return nil, false
	}
	return o.AiaEstimatedTimeSavedMs, true
}

// HasAiaEstimatedTimeSavedMs returns a boolean if a field has been set.
func (o *PlatformValues) HasAiaEstimatedTimeSavedMs() bool {
	if o != nil && o.AiaEstimatedTimeSavedMs != nil {
		return true
	}

	return false
}

// SetAiaEstimatedTimeSavedMs gets a reference to the given int64 and assigns it to the AiaEstimatedTimeSavedMs field.
func (o *PlatformValues) SetAiaEstimatedTimeSavedMs(v int64) {
	o.AiaEstimatedTimeSavedMs = &v
}

// GetRequestCollapseUsableCount returns the RequestCollapseUsableCount field value if set, zero value otherwise.
func (o *PlatformValues) GetRequestCollapseUsableCount() int64 {
	if o == nil || o.RequestCollapseUsableCount == nil {
		var ret int64
		return ret
	}
	return *o.RequestCollapseUsableCount
}

// GetRequestCollapseUsableCountOk returns a tuple with the RequestCollapseUsableCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetRequestCollapseUsableCountOk() (*int64, bool) {
	if o == nil || o.RequestCollapseUsableCount == nil {
		return nil, false
	}
	return o.RequestCollapseUsableCount, true
}

// HasRequestCollapseUsableCount returns a boolean if a field has been set.
func (o *PlatformValues) HasRequestCollapseUsableCount() bool {
	if o != nil && o.RequestCollapseUsableCount != nil {
		return true
	}

	return false
}

// SetRequestCollapseUsableCount gets a reference to the given int64 and assigns it to the RequestCollapseUsableCount field.
func (o *PlatformValues) SetRequestCollapseUsableCount(v int64) {
	o.RequestCollapseUsableCount = &v
}

// GetRequestCollapseUnusableCount returns the RequestCollapseUnusableCount field value if set, zero value otherwise.
func (o *PlatformValues) GetRequestCollapseUnusableCount() int64 {
	if o == nil || o.RequestCollapseUnusableCount == nil {
		var ret int64
		return ret
	}
	return *o.RequestCollapseUnusableCount
}

// GetRequestCollapseUnusableCountOk returns a tuple with the RequestCollapseUnusableCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetRequestCollapseUnusableCountOk() (*int64, bool) {
	if o == nil || o.RequestCollapseUnusableCount == nil {
		return nil, false
	}
	return o.RequestCollapseUnusableCount, true
}

// HasRequestCollapseUnusableCount returns a boolean if a field has been set.
func (o *PlatformValues) HasRequestCollapseUnusableCount() bool {
	if o != nil && o.RequestCollapseUnusableCount != nil {
		return true
	}

	return false
}

// SetRequestCollapseUnusableCount gets a reference to the given int64 and assigns it to the RequestCollapseUnusableCount field.
func (o *PlatformValues) SetRequestCollapseUnusableCount(v int64) {
	o.RequestCollapseUnusableCount = &v
}

// GetComputeCacheOperationsCount returns the ComputeCacheOperationsCount field value if set, zero value otherwise.
func (o *PlatformValues) GetComputeCacheOperationsCount() int64 {
	if o == nil || o.ComputeCacheOperationsCount == nil {
		var ret int64
		return ret
	}
	return *o.ComputeCacheOperationsCount
}

// GetComputeCacheOperationsCountOk returns a tuple with the ComputeCacheOperationsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetComputeCacheOperationsCountOk() (*int64, bool) {
	if o == nil || o.ComputeCacheOperationsCount == nil {
		return nil, false
	}
	return o.ComputeCacheOperationsCount, true
}

// HasComputeCacheOperationsCount returns a boolean if a field has been set.
func (o *PlatformValues) HasComputeCacheOperationsCount() bool {
	if o != nil && o.ComputeCacheOperationsCount != nil {
		return true
	}

	return false
}

// SetComputeCacheOperationsCount gets a reference to the given int64 and assigns it to the ComputeCacheOperationsCount field.
func (o *PlatformValues) SetComputeCacheOperationsCount(v int64) {
	o.ComputeCacheOperationsCount = &v
}

// GetNgwafRequestsTotalCount returns the NgwafRequestsTotalCount field value if set, zero value otherwise.
func (o *PlatformValues) GetNgwafRequestsTotalCount() int32 {
	if o == nil || o.NgwafRequestsTotalCount == nil {
		var ret int32
		return ret
	}
	return *o.NgwafRequestsTotalCount
}

// GetNgwafRequestsTotalCountOk returns a tuple with the NgwafRequestsTotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetNgwafRequestsTotalCountOk() (*int32, bool) {
	if o == nil || o.NgwafRequestsTotalCount == nil {
		return nil, false
	}
	return o.NgwafRequestsTotalCount, true
}

// HasNgwafRequestsTotalCount returns a boolean if a field has been set.
func (o *PlatformValues) HasNgwafRequestsTotalCount() bool {
	if o != nil && o.NgwafRequestsTotalCount != nil {
		return true
	}

	return false
}

// SetNgwafRequestsTotalCount gets a reference to the given int32 and assigns it to the NgwafRequestsTotalCount field.
func (o *PlatformValues) SetNgwafRequestsTotalCount(v int32) {
	o.NgwafRequestsTotalCount = &v
}

// GetNgwafRequestsUnknownCount returns the NgwafRequestsUnknownCount field value if set, zero value otherwise.
func (o *PlatformValues) GetNgwafRequestsUnknownCount() int32 {
	if o == nil || o.NgwafRequestsUnknownCount == nil {
		var ret int32
		return ret
	}
	return *o.NgwafRequestsUnknownCount
}

// GetNgwafRequestsUnknownCountOk returns a tuple with the NgwafRequestsUnknownCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetNgwafRequestsUnknownCountOk() (*int32, bool) {
	if o == nil || o.NgwafRequestsUnknownCount == nil {
		return nil, false
	}
	return o.NgwafRequestsUnknownCount, true
}

// HasNgwafRequestsUnknownCount returns a boolean if a field has been set.
func (o *PlatformValues) HasNgwafRequestsUnknownCount() bool {
	if o != nil && o.NgwafRequestsUnknownCount != nil {
		return true
	}

	return false
}

// SetNgwafRequestsUnknownCount gets a reference to the given int32 and assigns it to the NgwafRequestsUnknownCount field.
func (o *PlatformValues) SetNgwafRequestsUnknownCount(v int32) {
	o.NgwafRequestsUnknownCount = &v
}

// GetNgwafRequestsAllowedCount returns the NgwafRequestsAllowedCount field value if set, zero value otherwise.
func (o *PlatformValues) GetNgwafRequestsAllowedCount() int32 {
	if o == nil || o.NgwafRequestsAllowedCount == nil {
		var ret int32
		return ret
	}
	return *o.NgwafRequestsAllowedCount
}

// GetNgwafRequestsAllowedCountOk returns a tuple with the NgwafRequestsAllowedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetNgwafRequestsAllowedCountOk() (*int32, bool) {
	if o == nil || o.NgwafRequestsAllowedCount == nil {
		return nil, false
	}
	return o.NgwafRequestsAllowedCount, true
}

// HasNgwafRequestsAllowedCount returns a boolean if a field has been set.
func (o *PlatformValues) HasNgwafRequestsAllowedCount() bool {
	if o != nil && o.NgwafRequestsAllowedCount != nil {
		return true
	}

	return false
}

// SetNgwafRequestsAllowedCount gets a reference to the given int32 and assigns it to the NgwafRequestsAllowedCount field.
func (o *PlatformValues) SetNgwafRequestsAllowedCount(v int32) {
	o.NgwafRequestsAllowedCount = &v
}

// GetNgwafRequestsLoggedCount returns the NgwafRequestsLoggedCount field value if set, zero value otherwise.
func (o *PlatformValues) GetNgwafRequestsLoggedCount() int32 {
	if o == nil || o.NgwafRequestsLoggedCount == nil {
		var ret int32
		return ret
	}
	return *o.NgwafRequestsLoggedCount
}

// GetNgwafRequestsLoggedCountOk returns a tuple with the NgwafRequestsLoggedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetNgwafRequestsLoggedCountOk() (*int32, bool) {
	if o == nil || o.NgwafRequestsLoggedCount == nil {
		return nil, false
	}
	return o.NgwafRequestsLoggedCount, true
}

// HasNgwafRequestsLoggedCount returns a boolean if a field has been set.
func (o *PlatformValues) HasNgwafRequestsLoggedCount() bool {
	if o != nil && o.NgwafRequestsLoggedCount != nil {
		return true
	}

	return false
}

// SetNgwafRequestsLoggedCount gets a reference to the given int32 and assigns it to the NgwafRequestsLoggedCount field.
func (o *PlatformValues) SetNgwafRequestsLoggedCount(v int32) {
	o.NgwafRequestsLoggedCount = &v
}

// GetNgwafRequestsBlockedCount returns the NgwafRequestsBlockedCount field value if set, zero value otherwise.
func (o *PlatformValues) GetNgwafRequestsBlockedCount() int32 {
	if o == nil || o.NgwafRequestsBlockedCount == nil {
		var ret int32
		return ret
	}
	return *o.NgwafRequestsBlockedCount
}

// GetNgwafRequestsBlockedCountOk returns a tuple with the NgwafRequestsBlockedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetNgwafRequestsBlockedCountOk() (*int32, bool) {
	if o == nil || o.NgwafRequestsBlockedCount == nil {
		return nil, false
	}
	return o.NgwafRequestsBlockedCount, true
}

// HasNgwafRequestsBlockedCount returns a boolean if a field has been set.
func (o *PlatformValues) HasNgwafRequestsBlockedCount() bool {
	if o != nil && o.NgwafRequestsBlockedCount != nil {
		return true
	}

	return false
}

// SetNgwafRequestsBlockedCount gets a reference to the given int32 and assigns it to the NgwafRequestsBlockedCount field.
func (o *PlatformValues) SetNgwafRequestsBlockedCount(v int32) {
	o.NgwafRequestsBlockedCount = &v
}

// GetNgwafRequestsTimeoutCount returns the NgwafRequestsTimeoutCount field value if set, zero value otherwise.
func (o *PlatformValues) GetNgwafRequestsTimeoutCount() int32 {
	if o == nil || o.NgwafRequestsTimeoutCount == nil {
		var ret int32
		return ret
	}
	return *o.NgwafRequestsTimeoutCount
}

// GetNgwafRequestsTimeoutCountOk returns a tuple with the NgwafRequestsTimeoutCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetNgwafRequestsTimeoutCountOk() (*int32, bool) {
	if o == nil || o.NgwafRequestsTimeoutCount == nil {
		return nil, false
	}
	return o.NgwafRequestsTimeoutCount, true
}

// HasNgwafRequestsTimeoutCount returns a boolean if a field has been set.
func (o *PlatformValues) HasNgwafRequestsTimeoutCount() bool {
	if o != nil && o.NgwafRequestsTimeoutCount != nil {
		return true
	}

	return false
}

// SetNgwafRequestsTimeoutCount gets a reference to the given int32 and assigns it to the NgwafRequestsTimeoutCount field.
func (o *PlatformValues) SetNgwafRequestsTimeoutCount(v int32) {
	o.NgwafRequestsTimeoutCount = &v
}

// GetNgwafRequestsChallengedCount returns the NgwafRequestsChallengedCount field value if set, zero value otherwise.
func (o *PlatformValues) GetNgwafRequestsChallengedCount() int32 {
	if o == nil || o.NgwafRequestsChallengedCount == nil {
		var ret int32
		return ret
	}
	return *o.NgwafRequestsChallengedCount
}

// GetNgwafRequestsChallengedCountOk returns a tuple with the NgwafRequestsChallengedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetNgwafRequestsChallengedCountOk() (*int32, bool) {
	if o == nil || o.NgwafRequestsChallengedCount == nil {
		return nil, false
	}
	return o.NgwafRequestsChallengedCount, true
}

// HasNgwafRequestsChallengedCount returns a boolean if a field has been set.
func (o *PlatformValues) HasNgwafRequestsChallengedCount() bool {
	if o != nil && o.NgwafRequestsChallengedCount != nil {
		return true
	}

	return false
}

// SetNgwafRequestsChallengedCount gets a reference to the given int32 and assigns it to the NgwafRequestsChallengedCount field.
func (o *PlatformValues) SetNgwafRequestsChallengedCount(v int32) {
	o.NgwafRequestsChallengedCount = &v
}

// GetApiDiscoveryRequestsCount returns the ApiDiscoveryRequestsCount field value if set, zero value otherwise.
func (o *PlatformValues) GetApiDiscoveryRequestsCount() int32 {
	if o == nil || o.ApiDiscoveryRequestsCount == nil {
		var ret int32
		return ret
	}
	return *o.ApiDiscoveryRequestsCount
}

// GetApiDiscoveryRequestsCountOk returns a tuple with the ApiDiscoveryRequestsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetApiDiscoveryRequestsCountOk() (*int32, bool) {
	if o == nil || o.ApiDiscoveryRequestsCount == nil {
		return nil, false
	}
	return o.ApiDiscoveryRequestsCount, true
}

// HasApiDiscoveryRequestsCount returns a boolean if a field has been set.
func (o *PlatformValues) HasApiDiscoveryRequestsCount() bool {
	if o != nil && o.ApiDiscoveryRequestsCount != nil {
		return true
	}

	return false
}

// SetApiDiscoveryRequestsCount gets a reference to the given int32 and assigns it to the ApiDiscoveryRequestsCount field.
func (o *PlatformValues) SetApiDiscoveryRequestsCount(v int32) {
	o.ApiDiscoveryRequestsCount = &v
}

// GetComputeRespStatus103 returns the ComputeRespStatus103 field value if set, zero value otherwise.
func (o *PlatformValues) GetComputeRespStatus103() int32 {
	if o == nil || o.ComputeRespStatus103 == nil {
		var ret int32
		return ret
	}
	return *o.ComputeRespStatus103
}

// GetComputeRespStatus103Ok returns a tuple with the ComputeRespStatus103 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetComputeRespStatus103Ok() (*int32, bool) {
	if o == nil || o.ComputeRespStatus103 == nil {
		return nil, false
	}
	return o.ComputeRespStatus103, true
}

// HasComputeRespStatus103 returns a boolean if a field has been set.
func (o *PlatformValues) HasComputeRespStatus103() bool {
	if o != nil && o.ComputeRespStatus103 != nil {
		return true
	}

	return false
}

// SetComputeRespStatus103 gets a reference to the given int32 and assigns it to the ComputeRespStatus103 field.
func (o *PlatformValues) SetComputeRespStatus103(v int32) {
	o.ComputeRespStatus103 = &v
}

// GetComputeRespStatus200 returns the ComputeRespStatus200 field value if set, zero value otherwise.
func (o *PlatformValues) GetComputeRespStatus200() int32 {
	if o == nil || o.ComputeRespStatus200 == nil {
		var ret int32
		return ret
	}
	return *o.ComputeRespStatus200
}

// GetComputeRespStatus200Ok returns a tuple with the ComputeRespStatus200 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetComputeRespStatus200Ok() (*int32, bool) {
	if o == nil || o.ComputeRespStatus200 == nil {
		return nil, false
	}
	return o.ComputeRespStatus200, true
}

// HasComputeRespStatus200 returns a boolean if a field has been set.
func (o *PlatformValues) HasComputeRespStatus200() bool {
	if o != nil && o.ComputeRespStatus200 != nil {
		return true
	}

	return false
}

// SetComputeRespStatus200 gets a reference to the given int32 and assigns it to the ComputeRespStatus200 field.
func (o *PlatformValues) SetComputeRespStatus200(v int32) {
	o.ComputeRespStatus200 = &v
}

// GetComputeRespStatus204 returns the ComputeRespStatus204 field value if set, zero value otherwise.
func (o *PlatformValues) GetComputeRespStatus204() int32 {
	if o == nil || o.ComputeRespStatus204 == nil {
		var ret int32
		return ret
	}
	return *o.ComputeRespStatus204
}

// GetComputeRespStatus204Ok returns a tuple with the ComputeRespStatus204 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetComputeRespStatus204Ok() (*int32, bool) {
	if o == nil || o.ComputeRespStatus204 == nil {
		return nil, false
	}
	return o.ComputeRespStatus204, true
}

// HasComputeRespStatus204 returns a boolean if a field has been set.
func (o *PlatformValues) HasComputeRespStatus204() bool {
	if o != nil && o.ComputeRespStatus204 != nil {
		return true
	}

	return false
}

// SetComputeRespStatus204 gets a reference to the given int32 and assigns it to the ComputeRespStatus204 field.
func (o *PlatformValues) SetComputeRespStatus204(v int32) {
	o.ComputeRespStatus204 = &v
}

// GetComputeRespStatus206 returns the ComputeRespStatus206 field value if set, zero value otherwise.
func (o *PlatformValues) GetComputeRespStatus206() int32 {
	if o == nil || o.ComputeRespStatus206 == nil {
		var ret int32
		return ret
	}
	return *o.ComputeRespStatus206
}

// GetComputeRespStatus206Ok returns a tuple with the ComputeRespStatus206 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetComputeRespStatus206Ok() (*int32, bool) {
	if o == nil || o.ComputeRespStatus206 == nil {
		return nil, false
	}
	return o.ComputeRespStatus206, true
}

// HasComputeRespStatus206 returns a boolean if a field has been set.
func (o *PlatformValues) HasComputeRespStatus206() bool {
	if o != nil && o.ComputeRespStatus206 != nil {
		return true
	}

	return false
}

// SetComputeRespStatus206 gets a reference to the given int32 and assigns it to the ComputeRespStatus206 field.
func (o *PlatformValues) SetComputeRespStatus206(v int32) {
	o.ComputeRespStatus206 = &v
}

// GetComputeRespStatus301 returns the ComputeRespStatus301 field value if set, zero value otherwise.
func (o *PlatformValues) GetComputeRespStatus301() int32 {
	if o == nil || o.ComputeRespStatus301 == nil {
		var ret int32
		return ret
	}
	return *o.ComputeRespStatus301
}

// GetComputeRespStatus301Ok returns a tuple with the ComputeRespStatus301 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetComputeRespStatus301Ok() (*int32, bool) {
	if o == nil || o.ComputeRespStatus301 == nil {
		return nil, false
	}
	return o.ComputeRespStatus301, true
}

// HasComputeRespStatus301 returns a boolean if a field has been set.
func (o *PlatformValues) HasComputeRespStatus301() bool {
	if o != nil && o.ComputeRespStatus301 != nil {
		return true
	}

	return false
}

// SetComputeRespStatus301 gets a reference to the given int32 and assigns it to the ComputeRespStatus301 field.
func (o *PlatformValues) SetComputeRespStatus301(v int32) {
	o.ComputeRespStatus301 = &v
}

// GetComputeRespStatus302 returns the ComputeRespStatus302 field value if set, zero value otherwise.
func (o *PlatformValues) GetComputeRespStatus302() int32 {
	if o == nil || o.ComputeRespStatus302 == nil {
		var ret int32
		return ret
	}
	return *o.ComputeRespStatus302
}

// GetComputeRespStatus302Ok returns a tuple with the ComputeRespStatus302 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetComputeRespStatus302Ok() (*int32, bool) {
	if o == nil || o.ComputeRespStatus302 == nil {
		return nil, false
	}
	return o.ComputeRespStatus302, true
}

// HasComputeRespStatus302 returns a boolean if a field has been set.
func (o *PlatformValues) HasComputeRespStatus302() bool {
	if o != nil && o.ComputeRespStatus302 != nil {
		return true
	}

	return false
}

// SetComputeRespStatus302 gets a reference to the given int32 and assigns it to the ComputeRespStatus302 field.
func (o *PlatformValues) SetComputeRespStatus302(v int32) {
	o.ComputeRespStatus302 = &v
}

// GetComputeRespStatus304 returns the ComputeRespStatus304 field value if set, zero value otherwise.
func (o *PlatformValues) GetComputeRespStatus304() int32 {
	if o == nil || o.ComputeRespStatus304 == nil {
		var ret int32
		return ret
	}
	return *o.ComputeRespStatus304
}

// GetComputeRespStatus304Ok returns a tuple with the ComputeRespStatus304 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetComputeRespStatus304Ok() (*int32, bool) {
	if o == nil || o.ComputeRespStatus304 == nil {
		return nil, false
	}
	return o.ComputeRespStatus304, true
}

// HasComputeRespStatus304 returns a boolean if a field has been set.
func (o *PlatformValues) HasComputeRespStatus304() bool {
	if o != nil && o.ComputeRespStatus304 != nil {
		return true
	}

	return false
}

// SetComputeRespStatus304 gets a reference to the given int32 and assigns it to the ComputeRespStatus304 field.
func (o *PlatformValues) SetComputeRespStatus304(v int32) {
	o.ComputeRespStatus304 = &v
}

// GetComputeRespStatus400 returns the ComputeRespStatus400 field value if set, zero value otherwise.
func (o *PlatformValues) GetComputeRespStatus400() int32 {
	if o == nil || o.ComputeRespStatus400 == nil {
		var ret int32
		return ret
	}
	return *o.ComputeRespStatus400
}

// GetComputeRespStatus400Ok returns a tuple with the ComputeRespStatus400 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetComputeRespStatus400Ok() (*int32, bool) {
	if o == nil || o.ComputeRespStatus400 == nil {
		return nil, false
	}
	return o.ComputeRespStatus400, true
}

// HasComputeRespStatus400 returns a boolean if a field has been set.
func (o *PlatformValues) HasComputeRespStatus400() bool {
	if o != nil && o.ComputeRespStatus400 != nil {
		return true
	}

	return false
}

// SetComputeRespStatus400 gets a reference to the given int32 and assigns it to the ComputeRespStatus400 field.
func (o *PlatformValues) SetComputeRespStatus400(v int32) {
	o.ComputeRespStatus400 = &v
}

// GetComputeRespStatus401 returns the ComputeRespStatus401 field value if set, zero value otherwise.
func (o *PlatformValues) GetComputeRespStatus401() int32 {
	if o == nil || o.ComputeRespStatus401 == nil {
		var ret int32
		return ret
	}
	return *o.ComputeRespStatus401
}

// GetComputeRespStatus401Ok returns a tuple with the ComputeRespStatus401 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetComputeRespStatus401Ok() (*int32, bool) {
	if o == nil || o.ComputeRespStatus401 == nil {
		return nil, false
	}
	return o.ComputeRespStatus401, true
}

// HasComputeRespStatus401 returns a boolean if a field has been set.
func (o *PlatformValues) HasComputeRespStatus401() bool {
	if o != nil && o.ComputeRespStatus401 != nil {
		return true
	}

	return false
}

// SetComputeRespStatus401 gets a reference to the given int32 and assigns it to the ComputeRespStatus401 field.
func (o *PlatformValues) SetComputeRespStatus401(v int32) {
	o.ComputeRespStatus401 = &v
}

// GetComputeRespStatus403 returns the ComputeRespStatus403 field value if set, zero value otherwise.
func (o *PlatformValues) GetComputeRespStatus403() int32 {
	if o == nil || o.ComputeRespStatus403 == nil {
		var ret int32
		return ret
	}
	return *o.ComputeRespStatus403
}

// GetComputeRespStatus403Ok returns a tuple with the ComputeRespStatus403 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetComputeRespStatus403Ok() (*int32, bool) {
	if o == nil || o.ComputeRespStatus403 == nil {
		return nil, false
	}
	return o.ComputeRespStatus403, true
}

// HasComputeRespStatus403 returns a boolean if a field has been set.
func (o *PlatformValues) HasComputeRespStatus403() bool {
	if o != nil && o.ComputeRespStatus403 != nil {
		return true
	}

	return false
}

// SetComputeRespStatus403 gets a reference to the given int32 and assigns it to the ComputeRespStatus403 field.
func (o *PlatformValues) SetComputeRespStatus403(v int32) {
	o.ComputeRespStatus403 = &v
}

// GetComputeRespStatus404 returns the ComputeRespStatus404 field value if set, zero value otherwise.
func (o *PlatformValues) GetComputeRespStatus404() int32 {
	if o == nil || o.ComputeRespStatus404 == nil {
		var ret int32
		return ret
	}
	return *o.ComputeRespStatus404
}

// GetComputeRespStatus404Ok returns a tuple with the ComputeRespStatus404 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetComputeRespStatus404Ok() (*int32, bool) {
	if o == nil || o.ComputeRespStatus404 == nil {
		return nil, false
	}
	return o.ComputeRespStatus404, true
}

// HasComputeRespStatus404 returns a boolean if a field has been set.
func (o *PlatformValues) HasComputeRespStatus404() bool {
	if o != nil && o.ComputeRespStatus404 != nil {
		return true
	}

	return false
}

// SetComputeRespStatus404 gets a reference to the given int32 and assigns it to the ComputeRespStatus404 field.
func (o *PlatformValues) SetComputeRespStatus404(v int32) {
	o.ComputeRespStatus404 = &v
}

// GetComputeRespStatus416 returns the ComputeRespStatus416 field value if set, zero value otherwise.
func (o *PlatformValues) GetComputeRespStatus416() int32 {
	if o == nil || o.ComputeRespStatus416 == nil {
		var ret int32
		return ret
	}
	return *o.ComputeRespStatus416
}

// GetComputeRespStatus416Ok returns a tuple with the ComputeRespStatus416 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetComputeRespStatus416Ok() (*int32, bool) {
	if o == nil || o.ComputeRespStatus416 == nil {
		return nil, false
	}
	return o.ComputeRespStatus416, true
}

// HasComputeRespStatus416 returns a boolean if a field has been set.
func (o *PlatformValues) HasComputeRespStatus416() bool {
	if o != nil && o.ComputeRespStatus416 != nil {
		return true
	}

	return false
}

// SetComputeRespStatus416 gets a reference to the given int32 and assigns it to the ComputeRespStatus416 field.
func (o *PlatformValues) SetComputeRespStatus416(v int32) {
	o.ComputeRespStatus416 = &v
}

// GetComputeRespStatus429 returns the ComputeRespStatus429 field value if set, zero value otherwise.
func (o *PlatformValues) GetComputeRespStatus429() int32 {
	if o == nil || o.ComputeRespStatus429 == nil {
		var ret int32
		return ret
	}
	return *o.ComputeRespStatus429
}

// GetComputeRespStatus429Ok returns a tuple with the ComputeRespStatus429 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetComputeRespStatus429Ok() (*int32, bool) {
	if o == nil || o.ComputeRespStatus429 == nil {
		return nil, false
	}
	return o.ComputeRespStatus429, true
}

// HasComputeRespStatus429 returns a boolean if a field has been set.
func (o *PlatformValues) HasComputeRespStatus429() bool {
	if o != nil && o.ComputeRespStatus429 != nil {
		return true
	}

	return false
}

// SetComputeRespStatus429 gets a reference to the given int32 and assigns it to the ComputeRespStatus429 field.
func (o *PlatformValues) SetComputeRespStatus429(v int32) {
	o.ComputeRespStatus429 = &v
}

// GetComputeRespStatus500 returns the ComputeRespStatus500 field value if set, zero value otherwise.
func (o *PlatformValues) GetComputeRespStatus500() int32 {
	if o == nil || o.ComputeRespStatus500 == nil {
		var ret int32
		return ret
	}
	return *o.ComputeRespStatus500
}

// GetComputeRespStatus500Ok returns a tuple with the ComputeRespStatus500 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetComputeRespStatus500Ok() (*int32, bool) {
	if o == nil || o.ComputeRespStatus500 == nil {
		return nil, false
	}
	return o.ComputeRespStatus500, true
}

// HasComputeRespStatus500 returns a boolean if a field has been set.
func (o *PlatformValues) HasComputeRespStatus500() bool {
	if o != nil && o.ComputeRespStatus500 != nil {
		return true
	}

	return false
}

// SetComputeRespStatus500 gets a reference to the given int32 and assigns it to the ComputeRespStatus500 field.
func (o *PlatformValues) SetComputeRespStatus500(v int32) {
	o.ComputeRespStatus500 = &v
}

// GetComputeRespStatus501 returns the ComputeRespStatus501 field value if set, zero value otherwise.
func (o *PlatformValues) GetComputeRespStatus501() int32 {
	if o == nil || o.ComputeRespStatus501 == nil {
		var ret int32
		return ret
	}
	return *o.ComputeRespStatus501
}

// GetComputeRespStatus501Ok returns a tuple with the ComputeRespStatus501 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetComputeRespStatus501Ok() (*int32, bool) {
	if o == nil || o.ComputeRespStatus501 == nil {
		return nil, false
	}
	return o.ComputeRespStatus501, true
}

// HasComputeRespStatus501 returns a boolean if a field has been set.
func (o *PlatformValues) HasComputeRespStatus501() bool {
	if o != nil && o.ComputeRespStatus501 != nil {
		return true
	}

	return false
}

// SetComputeRespStatus501 gets a reference to the given int32 and assigns it to the ComputeRespStatus501 field.
func (o *PlatformValues) SetComputeRespStatus501(v int32) {
	o.ComputeRespStatus501 = &v
}

// GetComputeRespStatus502 returns the ComputeRespStatus502 field value if set, zero value otherwise.
func (o *PlatformValues) GetComputeRespStatus502() int32 {
	if o == nil || o.ComputeRespStatus502 == nil {
		var ret int32
		return ret
	}
	return *o.ComputeRespStatus502
}

// GetComputeRespStatus502Ok returns a tuple with the ComputeRespStatus502 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetComputeRespStatus502Ok() (*int32, bool) {
	if o == nil || o.ComputeRespStatus502 == nil {
		return nil, false
	}
	return o.ComputeRespStatus502, true
}

// HasComputeRespStatus502 returns a boolean if a field has been set.
func (o *PlatformValues) HasComputeRespStatus502() bool {
	if o != nil && o.ComputeRespStatus502 != nil {
		return true
	}

	return false
}

// SetComputeRespStatus502 gets a reference to the given int32 and assigns it to the ComputeRespStatus502 field.
func (o *PlatformValues) SetComputeRespStatus502(v int32) {
	o.ComputeRespStatus502 = &v
}

// GetComputeRespStatus503 returns the ComputeRespStatus503 field value if set, zero value otherwise.
func (o *PlatformValues) GetComputeRespStatus503() int32 {
	if o == nil || o.ComputeRespStatus503 == nil {
		var ret int32
		return ret
	}
	return *o.ComputeRespStatus503
}

// GetComputeRespStatus503Ok returns a tuple with the ComputeRespStatus503 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetComputeRespStatus503Ok() (*int32, bool) {
	if o == nil || o.ComputeRespStatus503 == nil {
		return nil, false
	}
	return o.ComputeRespStatus503, true
}

// HasComputeRespStatus503 returns a boolean if a field has been set.
func (o *PlatformValues) HasComputeRespStatus503() bool {
	if o != nil && o.ComputeRespStatus503 != nil {
		return true
	}

	return false
}

// SetComputeRespStatus503 gets a reference to the given int32 and assigns it to the ComputeRespStatus503 field.
func (o *PlatformValues) SetComputeRespStatus503(v int32) {
	o.ComputeRespStatus503 = &v
}

// GetComputeRespStatus504 returns the ComputeRespStatus504 field value if set, zero value otherwise.
func (o *PlatformValues) GetComputeRespStatus504() int32 {
	if o == nil || o.ComputeRespStatus504 == nil {
		var ret int32
		return ret
	}
	return *o.ComputeRespStatus504
}

// GetComputeRespStatus504Ok returns a tuple with the ComputeRespStatus504 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetComputeRespStatus504Ok() (*int32, bool) {
	if o == nil || o.ComputeRespStatus504 == nil {
		return nil, false
	}
	return o.ComputeRespStatus504, true
}

// HasComputeRespStatus504 returns a boolean if a field has been set.
func (o *PlatformValues) HasComputeRespStatus504() bool {
	if o != nil && o.ComputeRespStatus504 != nil {
		return true
	}

	return false
}

// SetComputeRespStatus504 gets a reference to the given int32 and assigns it to the ComputeRespStatus504 field.
func (o *PlatformValues) SetComputeRespStatus504(v int32) {
	o.ComputeRespStatus504 = &v
}

// GetComputeRespStatus505 returns the ComputeRespStatus505 field value if set, zero value otherwise.
func (o *PlatformValues) GetComputeRespStatus505() int32 {
	if o == nil || o.ComputeRespStatus505 == nil {
		var ret int32
		return ret
	}
	return *o.ComputeRespStatus505
}

// GetComputeRespStatus505Ok returns a tuple with the ComputeRespStatus505 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetComputeRespStatus505Ok() (*int32, bool) {
	if o == nil || o.ComputeRespStatus505 == nil {
		return nil, false
	}
	return o.ComputeRespStatus505, true
}

// HasComputeRespStatus505 returns a boolean if a field has been set.
func (o *PlatformValues) HasComputeRespStatus505() bool {
	if o != nil && o.ComputeRespStatus505 != nil {
		return true
	}

	return false
}

// SetComputeRespStatus505 gets a reference to the given int32 and assigns it to the ComputeRespStatus505 field.
func (o *PlatformValues) SetComputeRespStatus505(v int32) {
	o.ComputeRespStatus505 = &v
}

// GetComputeRespStatus530 returns the ComputeRespStatus530 field value if set, zero value otherwise.
func (o *PlatformValues) GetComputeRespStatus530() int32 {
	if o == nil || o.ComputeRespStatus530 == nil {
		var ret int32
		return ret
	}
	return *o.ComputeRespStatus530
}

// GetComputeRespStatus530Ok returns a tuple with the ComputeRespStatus530 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetComputeRespStatus530Ok() (*int32, bool) {
	if o == nil || o.ComputeRespStatus530 == nil {
		return nil, false
	}
	return o.ComputeRespStatus530, true
}

// HasComputeRespStatus530 returns a boolean if a field has been set.
func (o *PlatformValues) HasComputeRespStatus530() bool {
	if o != nil && o.ComputeRespStatus530 != nil {
		return true
	}

	return false
}

// SetComputeRespStatus530 gets a reference to the given int32 and assigns it to the ComputeRespStatus530 field.
func (o *PlatformValues) SetComputeRespStatus530(v int32) {
	o.ComputeRespStatus530 = &v
}

// GetImgoptoComputeRequests returns the ImgoptoComputeRequests field value if set, zero value otherwise.
func (o *PlatformValues) GetImgoptoComputeRequests() int32 {
	if o == nil || o.ImgoptoComputeRequests == nil {
		var ret int32
		return ret
	}
	return *o.ImgoptoComputeRequests
}

// GetImgoptoComputeRequestsOk returns a tuple with the ImgoptoComputeRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetImgoptoComputeRequestsOk() (*int32, bool) {
	if o == nil || o.ImgoptoComputeRequests == nil {
		return nil, false
	}
	return o.ImgoptoComputeRequests, true
}

// HasImgoptoComputeRequests returns a boolean if a field has been set.
func (o *PlatformValues) HasImgoptoComputeRequests() bool {
	if o != nil && o.ImgoptoComputeRequests != nil {
		return true
	}

	return false
}

// SetImgoptoComputeRequests gets a reference to the given int32 and assigns it to the ImgoptoComputeRequests field.
func (o *PlatformValues) SetImgoptoComputeRequests(v int32) {
	o.ImgoptoComputeRequests = &v
}

// GetDNSBillableResponsesCount returns the DNSBillableResponsesCount field value if set, zero value otherwise.
func (o *PlatformValues) GetDNSBillableResponsesCount() int32 {
	if o == nil || o.DNSBillableResponsesCount == nil {
		var ret int32
		return ret
	}
	return *o.DNSBillableResponsesCount
}

// GetDNSBillableResponsesCountOk returns a tuple with the DNSBillableResponsesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetDNSBillableResponsesCountOk() (*int32, bool) {
	if o == nil || o.DNSBillableResponsesCount == nil {
		return nil, false
	}
	return o.DNSBillableResponsesCount, true
}

// HasDNSBillableResponsesCount returns a boolean if a field has been set.
func (o *PlatformValues) HasDNSBillableResponsesCount() bool {
	if o != nil && o.DNSBillableResponsesCount != nil {
		return true
	}

	return false
}

// SetDNSBillableResponsesCount gets a reference to the given int32 and assigns it to the DNSBillableResponsesCount field.
func (o *PlatformValues) SetDNSBillableResponsesCount(v int32) {
	o.DNSBillableResponsesCount = &v
}

// GetDNSNonbillableResponsesCount returns the DNSNonbillableResponsesCount field value if set, zero value otherwise.
func (o *PlatformValues) GetDNSNonbillableResponsesCount() int32 {
	if o == nil || o.DNSNonbillableResponsesCount == nil {
		var ret int32
		return ret
	}
	return *o.DNSNonbillableResponsesCount
}

// GetDNSNonbillableResponsesCountOk returns a tuple with the DNSNonbillableResponsesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetDNSNonbillableResponsesCountOk() (*int32, bool) {
	if o == nil || o.DNSNonbillableResponsesCount == nil {
		return nil, false
	}
	return o.DNSNonbillableResponsesCount, true
}

// HasDNSNonbillableResponsesCount returns a boolean if a field has been set.
func (o *PlatformValues) HasDNSNonbillableResponsesCount() bool {
	if o != nil && o.DNSNonbillableResponsesCount != nil {
		return true
	}

	return false
}

// SetDNSNonbillableResponsesCount gets a reference to the given int32 and assigns it to the DNSNonbillableResponsesCount field.
func (o *PlatformValues) SetDNSNonbillableResponsesCount(v int32) {
	o.DNSNonbillableResponsesCount = &v
}

// GetUpgrade returns the Upgrade field value if set, zero value otherwise.
func (o *PlatformValues) GetUpgrade() int32 {
	if o == nil || o.Upgrade == nil {
		var ret int32
		return ret
	}
	return *o.Upgrade
}

// GetUpgradeOk returns a tuple with the Upgrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetUpgradeOk() (*int32, bool) {
	if o == nil || o.Upgrade == nil {
		return nil, false
	}
	return o.Upgrade, true
}

// HasUpgrade returns a boolean if a field has been set.
func (o *PlatformValues) HasUpgrade() bool {
	if o != nil && o.Upgrade != nil {
		return true
	}

	return false
}

// SetUpgrade gets a reference to the given int32 and assigns it to the Upgrade field.
func (o *PlatformValues) SetUpgrade(v int32) {
	o.Upgrade = &v
}

// GetNgwafBotAnalysisRequestCount returns the NgwafBotAnalysisRequestCount field value if set, zero value otherwise.
func (o *PlatformValues) GetNgwafBotAnalysisRequestCount() int32 {
	if o == nil || o.NgwafBotAnalysisRequestCount == nil {
		var ret int32
		return ret
	}
	return *o.NgwafBotAnalysisRequestCount
}

// GetNgwafBotAnalysisRequestCountOk returns a tuple with the NgwafBotAnalysisRequestCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetNgwafBotAnalysisRequestCountOk() (*int32, bool) {
	if o == nil || o.NgwafBotAnalysisRequestCount == nil {
		return nil, false
	}
	return o.NgwafBotAnalysisRequestCount, true
}

// HasNgwafBotAnalysisRequestCount returns a boolean if a field has been set.
func (o *PlatformValues) HasNgwafBotAnalysisRequestCount() bool {
	if o != nil && o.NgwafBotAnalysisRequestCount != nil {
		return true
	}

	return false
}

// SetNgwafBotAnalysisRequestCount gets a reference to the given int32 and assigns it to the NgwafBotAnalysisRequestCount field.
func (o *PlatformValues) SetNgwafBotAnalysisRequestCount(v int32) {
	o.NgwafBotAnalysisRequestCount = &v
}

// GetImgoptoAvifCount returns the ImgoptoAvifCount field value if set, zero value otherwise.
func (o *PlatformValues) GetImgoptoAvifCount() int32 {
	if o == nil || o.ImgoptoAvifCount == nil {
		var ret int32
		return ret
	}
	return *o.ImgoptoAvifCount
}

// GetImgoptoAvifCountOk returns a tuple with the ImgoptoAvifCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetImgoptoAvifCountOk() (*int32, bool) {
	if o == nil || o.ImgoptoAvifCount == nil {
		return nil, false
	}
	return o.ImgoptoAvifCount, true
}

// HasImgoptoAvifCount returns a boolean if a field has been set.
func (o *PlatformValues) HasImgoptoAvifCount() bool {
	if o != nil && o.ImgoptoAvifCount != nil {
		return true
	}

	return false
}

// SetImgoptoAvifCount gets a reference to the given int32 and assigns it to the ImgoptoAvifCount field.
func (o *PlatformValues) SetImgoptoAvifCount(v int32) {
	o.ImgoptoAvifCount = &v
}

// GetImgoptoJpegCount returns the ImgoptoJpegCount field value if set, zero value otherwise.
func (o *PlatformValues) GetImgoptoJpegCount() int32 {
	if o == nil || o.ImgoptoJpegCount == nil {
		var ret int32
		return ret
	}
	return *o.ImgoptoJpegCount
}

// GetImgoptoJpegCountOk returns a tuple with the ImgoptoJpegCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetImgoptoJpegCountOk() (*int32, bool) {
	if o == nil || o.ImgoptoJpegCount == nil {
		return nil, false
	}
	return o.ImgoptoJpegCount, true
}

// HasImgoptoJpegCount returns a boolean if a field has been set.
func (o *PlatformValues) HasImgoptoJpegCount() bool {
	if o != nil && o.ImgoptoJpegCount != nil {
		return true
	}

	return false
}

// SetImgoptoJpegCount gets a reference to the given int32 and assigns it to the ImgoptoJpegCount field.
func (o *PlatformValues) SetImgoptoJpegCount(v int32) {
	o.ImgoptoJpegCount = &v
}

// GetImgoptoPngCount returns the ImgoptoPngCount field value if set, zero value otherwise.
func (o *PlatformValues) GetImgoptoPngCount() int32 {
	if o == nil || o.ImgoptoPngCount == nil {
		var ret int32
		return ret
	}
	return *o.ImgoptoPngCount
}

// GetImgoptoPngCountOk returns a tuple with the ImgoptoPngCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetImgoptoPngCountOk() (*int32, bool) {
	if o == nil || o.ImgoptoPngCount == nil {
		return nil, false
	}
	return o.ImgoptoPngCount, true
}

// HasImgoptoPngCount returns a boolean if a field has been set.
func (o *PlatformValues) HasImgoptoPngCount() bool {
	if o != nil && o.ImgoptoPngCount != nil {
		return true
	}

	return false
}

// SetImgoptoPngCount gets a reference to the given int32 and assigns it to the ImgoptoPngCount field.
func (o *PlatformValues) SetImgoptoPngCount(v int32) {
	o.ImgoptoPngCount = &v
}

// GetImgoptoGifCount returns the ImgoptoGifCount field value if set, zero value otherwise.
func (o *PlatformValues) GetImgoptoGifCount() int32 {
	if o == nil || o.ImgoptoGifCount == nil {
		var ret int32
		return ret
	}
	return *o.ImgoptoGifCount
}

// GetImgoptoGifCountOk returns a tuple with the ImgoptoGifCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetImgoptoGifCountOk() (*int32, bool) {
	if o == nil || o.ImgoptoGifCount == nil {
		return nil, false
	}
	return o.ImgoptoGifCount, true
}

// HasImgoptoGifCount returns a boolean if a field has been set.
func (o *PlatformValues) HasImgoptoGifCount() bool {
	if o != nil && o.ImgoptoGifCount != nil {
		return true
	}

	return false
}

// SetImgoptoGifCount gets a reference to the given int32 and assigns it to the ImgoptoGifCount field.
func (o *PlatformValues) SetImgoptoGifCount(v int32) {
	o.ImgoptoGifCount = &v
}

// GetImgoptoWebpCount returns the ImgoptoWebpCount field value if set, zero value otherwise.
func (o *PlatformValues) GetImgoptoWebpCount() int32 {
	if o == nil || o.ImgoptoWebpCount == nil {
		var ret int32
		return ret
	}
	return *o.ImgoptoWebpCount
}

// GetImgoptoWebpCountOk returns a tuple with the ImgoptoWebpCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetImgoptoWebpCountOk() (*int32, bool) {
	if o == nil || o.ImgoptoWebpCount == nil {
		return nil, false
	}
	return o.ImgoptoWebpCount, true
}

// HasImgoptoWebpCount returns a boolean if a field has been set.
func (o *PlatformValues) HasImgoptoWebpCount() bool {
	if o != nil && o.ImgoptoWebpCount != nil {
		return true
	}

	return false
}

// SetImgoptoWebpCount gets a reference to the given int32 and assigns it to the ImgoptoWebpCount field.
func (o *PlatformValues) SetImgoptoWebpCount(v int32) {
	o.ImgoptoWebpCount = &v
}

// GetImgoptoJpegxlCount returns the ImgoptoJpegxlCount field value if set, zero value otherwise.
func (o *PlatformValues) GetImgoptoJpegxlCount() int32 {
	if o == nil || o.ImgoptoJpegxlCount == nil {
		var ret int32
		return ret
	}
	return *o.ImgoptoJpegxlCount
}

// GetImgoptoJpegxlCountOk returns a tuple with the ImgoptoJpegxlCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetImgoptoJpegxlCountOk() (*int32, bool) {
	if o == nil || o.ImgoptoJpegxlCount == nil {
		return nil, false
	}
	return o.ImgoptoJpegxlCount, true
}

// HasImgoptoJpegxlCount returns a boolean if a field has been set.
func (o *PlatformValues) HasImgoptoJpegxlCount() bool {
	if o != nil && o.ImgoptoJpegxlCount != nil {
		return true
	}

	return false
}

// SetImgoptoJpegxlCount gets a reference to the given int32 and assigns it to the ImgoptoJpegxlCount field.
func (o *PlatformValues) SetImgoptoJpegxlCount(v int32) {
	o.ImgoptoJpegxlCount = &v
}

// GetImgoptoSvgCount returns the ImgoptoSvgCount field value if set, zero value otherwise.
func (o *PlatformValues) GetImgoptoSvgCount() int32 {
	if o == nil || o.ImgoptoSvgCount == nil {
		var ret int32
		return ret
	}
	return *o.ImgoptoSvgCount
}

// GetImgoptoSvgCountOk returns a tuple with the ImgoptoSvgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetImgoptoSvgCountOk() (*int32, bool) {
	if o == nil || o.ImgoptoSvgCount == nil {
		return nil, false
	}
	return o.ImgoptoSvgCount, true
}

// HasImgoptoSvgCount returns a boolean if a field has been set.
func (o *PlatformValues) HasImgoptoSvgCount() bool {
	if o != nil && o.ImgoptoSvgCount != nil {
		return true
	}

	return false
}

// SetImgoptoSvgCount gets a reference to the given int32 and assigns it to the ImgoptoSvgCount field.
func (o *PlatformValues) SetImgoptoSvgCount(v int32) {
	o.ImgoptoSvgCount = &v
}

// GetImgoptoMp4Count returns the ImgoptoMp4Count field value if set, zero value otherwise.
func (o *PlatformValues) GetImgoptoMp4Count() int32 {
	if o == nil || o.ImgoptoMp4Count == nil {
		var ret int32
		return ret
	}
	return *o.ImgoptoMp4Count
}

// GetImgoptoMp4CountOk returns a tuple with the ImgoptoMp4Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetImgoptoMp4CountOk() (*int32, bool) {
	if o == nil || o.ImgoptoMp4Count == nil {
		return nil, false
	}
	return o.ImgoptoMp4Count, true
}

// HasImgoptoMp4Count returns a boolean if a field has been set.
func (o *PlatformValues) HasImgoptoMp4Count() bool {
	if o != nil && o.ImgoptoMp4Count != nil {
		return true
	}

	return false
}

// SetImgoptoMp4Count gets a reference to the given int32 and assigns it to the ImgoptoMp4Count field.
func (o *PlatformValues) SetImgoptoMp4Count(v int32) {
	o.ImgoptoMp4Count = &v
}

// GetComputeServiceResourceLimitsError returns the ComputeServiceResourceLimitsError field value if set, zero value otherwise.
func (o *PlatformValues) GetComputeServiceResourceLimitsError() int32 {
	if o == nil || o.ComputeServiceResourceLimitsError == nil {
		var ret int32
		return ret
	}
	return *o.ComputeServiceResourceLimitsError
}

// GetComputeServiceResourceLimitsErrorOk returns a tuple with the ComputeServiceResourceLimitsError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetComputeServiceResourceLimitsErrorOk() (*int32, bool) {
	if o == nil || o.ComputeServiceResourceLimitsError == nil {
		return nil, false
	}
	return o.ComputeServiceResourceLimitsError, true
}

// HasComputeServiceResourceLimitsError returns a boolean if a field has been set.
func (o *PlatformValues) HasComputeServiceResourceLimitsError() bool {
	if o != nil && o.ComputeServiceResourceLimitsError != nil {
		return true
	}

	return false
}

// SetComputeServiceResourceLimitsError gets a reference to the given int32 and assigns it to the ComputeServiceResourceLimitsError field.
func (o *PlatformValues) SetComputeServiceResourceLimitsError(v int32) {
	o.ComputeServiceResourceLimitsError = &v
}

// GetComputeServiceRuntimeError returns the ComputeServiceRuntimeError field value if set, zero value otherwise.
func (o *PlatformValues) GetComputeServiceRuntimeError() int32 {
	if o == nil || o.ComputeServiceRuntimeError == nil {
		var ret int32
		return ret
	}
	return *o.ComputeServiceRuntimeError
}

// GetComputeServiceRuntimeErrorOk returns a tuple with the ComputeServiceRuntimeError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetComputeServiceRuntimeErrorOk() (*int32, bool) {
	if o == nil || o.ComputeServiceRuntimeError == nil {
		return nil, false
	}
	return o.ComputeServiceRuntimeError, true
}

// HasComputeServiceRuntimeError returns a boolean if a field has been set.
func (o *PlatformValues) HasComputeServiceRuntimeError() bool {
	if o != nil && o.ComputeServiceRuntimeError != nil {
		return true
	}

	return false
}

// SetComputeServiceRuntimeError gets a reference to the given int32 and assigns it to the ComputeServiceRuntimeError field.
func (o *PlatformValues) SetComputeServiceRuntimeError(v int32) {
	o.ComputeServiceRuntimeError = &v
}

// GetComputeServiceChainError returns the ComputeServiceChainError field value if set, zero value otherwise.
func (o *PlatformValues) GetComputeServiceChainError() int32 {
	if o == nil || o.ComputeServiceChainError == nil {
		var ret int32
		return ret
	}
	return *o.ComputeServiceChainError
}

// GetComputeServiceChainErrorOk returns a tuple with the ComputeServiceChainError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetComputeServiceChainErrorOk() (*int32, bool) {
	if o == nil || o.ComputeServiceChainError == nil {
		return nil, false
	}
	return o.ComputeServiceChainError, true
}

// HasComputeServiceChainError returns a boolean if a field has been set.
func (o *PlatformValues) HasComputeServiceChainError() bool {
	if o != nil && o.ComputeServiceChainError != nil {
		return true
	}

	return false
}

// SetComputeServiceChainError gets a reference to the given int32 and assigns it to the ComputeServiceChainError field.
func (o *PlatformValues) SetComputeServiceChainError(v int32) {
	o.ComputeServiceChainError = &v
}

// GetComputePlatformInternalError returns the ComputePlatformInternalError field value if set, zero value otherwise.
func (o *PlatformValues) GetComputePlatformInternalError() int32 {
	if o == nil || o.ComputePlatformInternalError == nil {
		var ret int32
		return ret
	}
	return *o.ComputePlatformInternalError
}

// GetComputePlatformInternalErrorOk returns a tuple with the ComputePlatformInternalError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetComputePlatformInternalErrorOk() (*int32, bool) {
	if o == nil || o.ComputePlatformInternalError == nil {
		return nil, false
	}
	return o.ComputePlatformInternalError, true
}

// HasComputePlatformInternalError returns a boolean if a field has been set.
func (o *PlatformValues) HasComputePlatformInternalError() bool {
	if o != nil && o.ComputePlatformInternalError != nil {
		return true
	}

	return false
}

// SetComputePlatformInternalError gets a reference to the given int32 and assigns it to the ComputePlatformInternalError field.
func (o *PlatformValues) SetComputePlatformInternalError(v int32) {
	o.ComputePlatformInternalError = &v
}

// GetComputeServiceTimeoutError returns the ComputeServiceTimeoutError field value if set, zero value otherwise.
func (o *PlatformValues) GetComputeServiceTimeoutError() int32 {
	if o == nil || o.ComputeServiceTimeoutError == nil {
		var ret int32
		return ret
	}
	return *o.ComputeServiceTimeoutError
}

// GetComputeServiceTimeoutErrorOk returns a tuple with the ComputeServiceTimeoutError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetComputeServiceTimeoutErrorOk() (*int32, bool) {
	if o == nil || o.ComputeServiceTimeoutError == nil {
		return nil, false
	}
	return o.ComputeServiceTimeoutError, true
}

// HasComputeServiceTimeoutError returns a boolean if a field has been set.
func (o *PlatformValues) HasComputeServiceTimeoutError() bool {
	if o != nil && o.ComputeServiceTimeoutError != nil {
		return true
	}

	return false
}

// SetComputeServiceTimeoutError gets a reference to the given int32 and assigns it to the ComputeServiceTimeoutError field.
func (o *PlatformValues) SetComputeServiceTimeoutError(v int32) {
	o.ComputeServiceTimeoutError = &v
}

// GetComputeServiceVcpuExceededError returns the ComputeServiceVcpuExceededError field value if set, zero value otherwise.
func (o *PlatformValues) GetComputeServiceVcpuExceededError() int32 {
	if o == nil || o.ComputeServiceVcpuExceededError == nil {
		var ret int32
		return ret
	}
	return *o.ComputeServiceVcpuExceededError
}

// GetComputeServiceVcpuExceededErrorOk returns a tuple with the ComputeServiceVcpuExceededError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetComputeServiceVcpuExceededErrorOk() (*int32, bool) {
	if o == nil || o.ComputeServiceVcpuExceededError == nil {
		return nil, false
	}
	return o.ComputeServiceVcpuExceededError, true
}

// HasComputeServiceVcpuExceededError returns a boolean if a field has been set.
func (o *PlatformValues) HasComputeServiceVcpuExceededError() bool {
	if o != nil && o.ComputeServiceVcpuExceededError != nil {
		return true
	}

	return false
}

// SetComputeServiceVcpuExceededError gets a reference to the given int32 and assigns it to the ComputeServiceVcpuExceededError field.
func (o *PlatformValues) SetComputeServiceVcpuExceededError(v int32) {
	o.ComputeServiceVcpuExceededError = &v
}

// GetComputeServiceLimitsError returns the ComputeServiceLimitsError field value if set, zero value otherwise.
func (o *PlatformValues) GetComputeServiceLimitsError() int32 {
	if o == nil || o.ComputeServiceLimitsError == nil {
		var ret int32
		return ret
	}
	return *o.ComputeServiceLimitsError
}

// GetComputeServiceLimitsErrorOk returns a tuple with the ComputeServiceLimitsError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetComputeServiceLimitsErrorOk() (*int32, bool) {
	if o == nil || o.ComputeServiceLimitsError == nil {
		return nil, false
	}
	return o.ComputeServiceLimitsError, true
}

// HasComputeServiceLimitsError returns a boolean if a field has been set.
func (o *PlatformValues) HasComputeServiceLimitsError() bool {
	if o != nil && o.ComputeServiceLimitsError != nil {
		return true
	}

	return false
}

// SetComputeServiceLimitsError gets a reference to the given int32 and assigns it to the ComputeServiceLimitsError field.
func (o *PlatformValues) SetComputeServiceLimitsError(v int32) {
	o.ComputeServiceLimitsError = &v
}

// GetComputePlatformInvalidRequestError returns the ComputePlatformInvalidRequestError field value if set, zero value otherwise.
func (o *PlatformValues) GetComputePlatformInvalidRequestError() int32 {
	if o == nil || o.ComputePlatformInvalidRequestError == nil {
		var ret int32
		return ret
	}
	return *o.ComputePlatformInvalidRequestError
}

// GetComputePlatformInvalidRequestErrorOk returns a tuple with the ComputePlatformInvalidRequestError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetComputePlatformInvalidRequestErrorOk() (*int32, bool) {
	if o == nil || o.ComputePlatformInvalidRequestError == nil {
		return nil, false
	}
	return o.ComputePlatformInvalidRequestError, true
}

// HasComputePlatformInvalidRequestError returns a boolean if a field has been set.
func (o *PlatformValues) HasComputePlatformInvalidRequestError() bool {
	if o != nil && o.ComputePlatformInvalidRequestError != nil {
		return true
	}

	return false
}

// SetComputePlatformInvalidRequestError gets a reference to the given int32 and assigns it to the ComputePlatformInvalidRequestError field.
func (o *PlatformValues) SetComputePlatformInvalidRequestError(v int32) {
	o.ComputePlatformInvalidRequestError = &v
}

// GetComputeSandboxes returns the ComputeSandboxes field value if set, zero value otherwise.
func (o *PlatformValues) GetComputeSandboxes() int32 {
	if o == nil || o.ComputeSandboxes == nil {
		var ret int32
		return ret
	}
	return *o.ComputeSandboxes
}

// GetComputeSandboxesOk returns a tuple with the ComputeSandboxes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetComputeSandboxesOk() (*int32, bool) {
	if o == nil || o.ComputeSandboxes == nil {
		return nil, false
	}
	return o.ComputeSandboxes, true
}

// HasComputeSandboxes returns a boolean if a field has been set.
func (o *PlatformValues) HasComputeSandboxes() bool {
	if o != nil && o.ComputeSandboxes != nil {
		return true
	}

	return false
}

// SetComputeSandboxes gets a reference to the given int32 and assigns it to the ComputeSandboxes field.
func (o *PlatformValues) SetComputeSandboxes(v int32) {
	o.ComputeSandboxes = &v
}

// GetBotRequestsTotalCount returns the BotRequestsTotalCount field value if set, zero value otherwise.
func (o *PlatformValues) GetBotRequestsTotalCount() int32 {
	if o == nil || o.BotRequestsTotalCount == nil {
		var ret int32
		return ret
	}
	return *o.BotRequestsTotalCount
}

// GetBotRequestsTotalCountOk returns a tuple with the BotRequestsTotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetBotRequestsTotalCountOk() (*int32, bool) {
	if o == nil || o.BotRequestsTotalCount == nil {
		return nil, false
	}
	return o.BotRequestsTotalCount, true
}

// HasBotRequestsTotalCount returns a boolean if a field has been set.
func (o *PlatformValues) HasBotRequestsTotalCount() bool {
	if o != nil && o.BotRequestsTotalCount != nil {
		return true
	}

	return false
}

// SetBotRequestsTotalCount gets a reference to the given int32 and assigns it to the BotRequestsTotalCount field.
func (o *PlatformValues) SetBotRequestsTotalCount(v int32) {
	o.BotRequestsTotalCount = &v
}

// GetBotEdgeRequestsAnalyzedCount returns the BotEdgeRequestsAnalyzedCount field value if set, zero value otherwise.
func (o *PlatformValues) GetBotEdgeRequestsAnalyzedCount() int32 {
	if o == nil || o.BotEdgeRequestsAnalyzedCount == nil {
		var ret int32
		return ret
	}
	return *o.BotEdgeRequestsAnalyzedCount
}

// GetBotEdgeRequestsAnalyzedCountOk returns a tuple with the BotEdgeRequestsAnalyzedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetBotEdgeRequestsAnalyzedCountOk() (*int32, bool) {
	if o == nil || o.BotEdgeRequestsAnalyzedCount == nil {
		return nil, false
	}
	return o.BotEdgeRequestsAnalyzedCount, true
}

// HasBotEdgeRequestsAnalyzedCount returns a boolean if a field has been set.
func (o *PlatformValues) HasBotEdgeRequestsAnalyzedCount() bool {
	if o != nil && o.BotEdgeRequestsAnalyzedCount != nil {
		return true
	}

	return false
}

// SetBotEdgeRequestsAnalyzedCount gets a reference to the given int32 and assigns it to the BotEdgeRequestsAnalyzedCount field.
func (o *PlatformValues) SetBotEdgeRequestsAnalyzedCount(v int32) {
	o.BotEdgeRequestsAnalyzedCount = &v
}

// GetBotEdgeRequestsDetectedCount returns the BotEdgeRequestsDetectedCount field value if set, zero value otherwise.
func (o *PlatformValues) GetBotEdgeRequestsDetectedCount() int32 {
	if o == nil || o.BotEdgeRequestsDetectedCount == nil {
		var ret int32
		return ret
	}
	return *o.BotEdgeRequestsDetectedCount
}

// GetBotEdgeRequestsDetectedCountOk returns a tuple with the BotEdgeRequestsDetectedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetBotEdgeRequestsDetectedCountOk() (*int32, bool) {
	if o == nil || o.BotEdgeRequestsDetectedCount == nil {
		return nil, false
	}
	return o.BotEdgeRequestsDetectedCount, true
}

// HasBotEdgeRequestsDetectedCount returns a boolean if a field has been set.
func (o *PlatformValues) HasBotEdgeRequestsDetectedCount() bool {
	if o != nil && o.BotEdgeRequestsDetectedCount != nil {
		return true
	}

	return false
}

// SetBotEdgeRequestsDetectedCount gets a reference to the given int32 and assigns it to the BotEdgeRequestsDetectedCount field.
func (o *PlatformValues) SetBotEdgeRequestsDetectedCount(v int32) {
	o.BotEdgeRequestsDetectedCount = &v
}

// GetBotEdgeRequestsVerifiedCount returns the BotEdgeRequestsVerifiedCount field value if set, zero value otherwise.
func (o *PlatformValues) GetBotEdgeRequestsVerifiedCount() int32 {
	if o == nil || o.BotEdgeRequestsVerifiedCount == nil {
		var ret int32
		return ret
	}
	return *o.BotEdgeRequestsVerifiedCount
}

// GetBotEdgeRequestsVerifiedCountOk returns a tuple with the BotEdgeRequestsVerifiedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetBotEdgeRequestsVerifiedCountOk() (*int32, bool) {
	if o == nil || o.BotEdgeRequestsVerifiedCount == nil {
		return nil, false
	}
	return o.BotEdgeRequestsVerifiedCount, true
}

// HasBotEdgeRequestsVerifiedCount returns a boolean if a field has been set.
func (o *PlatformValues) HasBotEdgeRequestsVerifiedCount() bool {
	if o != nil && o.BotEdgeRequestsVerifiedCount != nil {
		return true
	}

	return false
}

// SetBotEdgeRequestsVerifiedCount gets a reference to the given int32 and assigns it to the BotEdgeRequestsVerifiedCount field.
func (o *PlatformValues) SetBotEdgeRequestsVerifiedCount(v int32) {
	o.BotEdgeRequestsVerifiedCount = &v
}

// GetBotEdgeRequestsAiCrawlerCount returns the BotEdgeRequestsAiCrawlerCount field value if set, zero value otherwise.
func (o *PlatformValues) GetBotEdgeRequestsAiCrawlerCount() int32 {
	if o == nil || o.BotEdgeRequestsAiCrawlerCount == nil {
		var ret int32
		return ret
	}
	return *o.BotEdgeRequestsAiCrawlerCount
}

// GetBotEdgeRequestsAiCrawlerCountOk returns a tuple with the BotEdgeRequestsAiCrawlerCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetBotEdgeRequestsAiCrawlerCountOk() (*int32, bool) {
	if o == nil || o.BotEdgeRequestsAiCrawlerCount == nil {
		return nil, false
	}
	return o.BotEdgeRequestsAiCrawlerCount, true
}

// HasBotEdgeRequestsAiCrawlerCount returns a boolean if a field has been set.
func (o *PlatformValues) HasBotEdgeRequestsAiCrawlerCount() bool {
	if o != nil && o.BotEdgeRequestsAiCrawlerCount != nil {
		return true
	}

	return false
}

// SetBotEdgeRequestsAiCrawlerCount gets a reference to the given int32 and assigns it to the BotEdgeRequestsAiCrawlerCount field.
func (o *PlatformValues) SetBotEdgeRequestsAiCrawlerCount(v int32) {
	o.BotEdgeRequestsAiCrawlerCount = &v
}

// GetBotEdgeRequestsAiFetcherCount returns the BotEdgeRequestsAiFetcherCount field value if set, zero value otherwise.
func (o *PlatformValues) GetBotEdgeRequestsAiFetcherCount() int32 {
	if o == nil || o.BotEdgeRequestsAiFetcherCount == nil {
		var ret int32
		return ret
	}
	return *o.BotEdgeRequestsAiFetcherCount
}

// GetBotEdgeRequestsAiFetcherCountOk returns a tuple with the BotEdgeRequestsAiFetcherCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetBotEdgeRequestsAiFetcherCountOk() (*int32, bool) {
	if o == nil || o.BotEdgeRequestsAiFetcherCount == nil {
		return nil, false
	}
	return o.BotEdgeRequestsAiFetcherCount, true
}

// HasBotEdgeRequestsAiFetcherCount returns a boolean if a field has been set.
func (o *PlatformValues) HasBotEdgeRequestsAiFetcherCount() bool {
	if o != nil && o.BotEdgeRequestsAiFetcherCount != nil {
		return true
	}

	return false
}

// SetBotEdgeRequestsAiFetcherCount gets a reference to the given int32 and assigns it to the BotEdgeRequestsAiFetcherCount field.
func (o *PlatformValues) SetBotEdgeRequestsAiFetcherCount(v int32) {
	o.BotEdgeRequestsAiFetcherCount = &v
}

// GetBotEdgeRequestsAccessibilityCount returns the BotEdgeRequestsAccessibilityCount field value if set, zero value otherwise.
func (o *PlatformValues) GetBotEdgeRequestsAccessibilityCount() int32 {
	if o == nil || o.BotEdgeRequestsAccessibilityCount == nil {
		var ret int32
		return ret
	}
	return *o.BotEdgeRequestsAccessibilityCount
}

// GetBotEdgeRequestsAccessibilityCountOk returns a tuple with the BotEdgeRequestsAccessibilityCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetBotEdgeRequestsAccessibilityCountOk() (*int32, bool) {
	if o == nil || o.BotEdgeRequestsAccessibilityCount == nil {
		return nil, false
	}
	return o.BotEdgeRequestsAccessibilityCount, true
}

// HasBotEdgeRequestsAccessibilityCount returns a boolean if a field has been set.
func (o *PlatformValues) HasBotEdgeRequestsAccessibilityCount() bool {
	if o != nil && o.BotEdgeRequestsAccessibilityCount != nil {
		return true
	}

	return false
}

// SetBotEdgeRequestsAccessibilityCount gets a reference to the given int32 and assigns it to the BotEdgeRequestsAccessibilityCount field.
func (o *PlatformValues) SetBotEdgeRequestsAccessibilityCount(v int32) {
	o.BotEdgeRequestsAccessibilityCount = &v
}

// GetBotEdgeRequestsContentFetcherCount returns the BotEdgeRequestsContentFetcherCount field value if set, zero value otherwise.
func (o *PlatformValues) GetBotEdgeRequestsContentFetcherCount() int32 {
	if o == nil || o.BotEdgeRequestsContentFetcherCount == nil {
		var ret int32
		return ret
	}
	return *o.BotEdgeRequestsContentFetcherCount
}

// GetBotEdgeRequestsContentFetcherCountOk returns a tuple with the BotEdgeRequestsContentFetcherCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetBotEdgeRequestsContentFetcherCountOk() (*int32, bool) {
	if o == nil || o.BotEdgeRequestsContentFetcherCount == nil {
		return nil, false
	}
	return o.BotEdgeRequestsContentFetcherCount, true
}

// HasBotEdgeRequestsContentFetcherCount returns a boolean if a field has been set.
func (o *PlatformValues) HasBotEdgeRequestsContentFetcherCount() bool {
	if o != nil && o.BotEdgeRequestsContentFetcherCount != nil {
		return true
	}

	return false
}

// SetBotEdgeRequestsContentFetcherCount gets a reference to the given int32 and assigns it to the BotEdgeRequestsContentFetcherCount field.
func (o *PlatformValues) SetBotEdgeRequestsContentFetcherCount(v int32) {
	o.BotEdgeRequestsContentFetcherCount = &v
}

// GetBotEdgeRequestsMonitoringCount returns the BotEdgeRequestsMonitoringCount field value if set, zero value otherwise.
func (o *PlatformValues) GetBotEdgeRequestsMonitoringCount() int32 {
	if o == nil || o.BotEdgeRequestsMonitoringCount == nil {
		var ret int32
		return ret
	}
	return *o.BotEdgeRequestsMonitoringCount
}

// GetBotEdgeRequestsMonitoringCountOk returns a tuple with the BotEdgeRequestsMonitoringCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetBotEdgeRequestsMonitoringCountOk() (*int32, bool) {
	if o == nil || o.BotEdgeRequestsMonitoringCount == nil {
		return nil, false
	}
	return o.BotEdgeRequestsMonitoringCount, true
}

// HasBotEdgeRequestsMonitoringCount returns a boolean if a field has been set.
func (o *PlatformValues) HasBotEdgeRequestsMonitoringCount() bool {
	if o != nil && o.BotEdgeRequestsMonitoringCount != nil {
		return true
	}

	return false
}

// SetBotEdgeRequestsMonitoringCount gets a reference to the given int32 and assigns it to the BotEdgeRequestsMonitoringCount field.
func (o *PlatformValues) SetBotEdgeRequestsMonitoringCount(v int32) {
	o.BotEdgeRequestsMonitoringCount = &v
}

// GetBotEdgeRequestsOnlineMarketingCount returns the BotEdgeRequestsOnlineMarketingCount field value if set, zero value otherwise.
func (o *PlatformValues) GetBotEdgeRequestsOnlineMarketingCount() int32 {
	if o == nil || o.BotEdgeRequestsOnlineMarketingCount == nil {
		var ret int32
		return ret
	}
	return *o.BotEdgeRequestsOnlineMarketingCount
}

// GetBotEdgeRequestsOnlineMarketingCountOk returns a tuple with the BotEdgeRequestsOnlineMarketingCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetBotEdgeRequestsOnlineMarketingCountOk() (*int32, bool) {
	if o == nil || o.BotEdgeRequestsOnlineMarketingCount == nil {
		return nil, false
	}
	return o.BotEdgeRequestsOnlineMarketingCount, true
}

// HasBotEdgeRequestsOnlineMarketingCount returns a boolean if a field has been set.
func (o *PlatformValues) HasBotEdgeRequestsOnlineMarketingCount() bool {
	if o != nil && o.BotEdgeRequestsOnlineMarketingCount != nil {
		return true
	}

	return false
}

// SetBotEdgeRequestsOnlineMarketingCount gets a reference to the given int32 and assigns it to the BotEdgeRequestsOnlineMarketingCount field.
func (o *PlatformValues) SetBotEdgeRequestsOnlineMarketingCount(v int32) {
	o.BotEdgeRequestsOnlineMarketingCount = &v
}

// GetBotEdgeRequestsPagePreviewCount returns the BotEdgeRequestsPagePreviewCount field value if set, zero value otherwise.
func (o *PlatformValues) GetBotEdgeRequestsPagePreviewCount() int32 {
	if o == nil || o.BotEdgeRequestsPagePreviewCount == nil {
		var ret int32
		return ret
	}
	return *o.BotEdgeRequestsPagePreviewCount
}

// GetBotEdgeRequestsPagePreviewCountOk returns a tuple with the BotEdgeRequestsPagePreviewCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetBotEdgeRequestsPagePreviewCountOk() (*int32, bool) {
	if o == nil || o.BotEdgeRequestsPagePreviewCount == nil {
		return nil, false
	}
	return o.BotEdgeRequestsPagePreviewCount, true
}

// HasBotEdgeRequestsPagePreviewCount returns a boolean if a field has been set.
func (o *PlatformValues) HasBotEdgeRequestsPagePreviewCount() bool {
	if o != nil && o.BotEdgeRequestsPagePreviewCount != nil {
		return true
	}

	return false
}

// SetBotEdgeRequestsPagePreviewCount gets a reference to the given int32 and assigns it to the BotEdgeRequestsPagePreviewCount field.
func (o *PlatformValues) SetBotEdgeRequestsPagePreviewCount(v int32) {
	o.BotEdgeRequestsPagePreviewCount = &v
}

// GetBotEdgeRequestsPlatformIntegrationsCount returns the BotEdgeRequestsPlatformIntegrationsCount field value if set, zero value otherwise.
func (o *PlatformValues) GetBotEdgeRequestsPlatformIntegrationsCount() int32 {
	if o == nil || o.BotEdgeRequestsPlatformIntegrationsCount == nil {
		var ret int32
		return ret
	}
	return *o.BotEdgeRequestsPlatformIntegrationsCount
}

// GetBotEdgeRequestsPlatformIntegrationsCountOk returns a tuple with the BotEdgeRequestsPlatformIntegrationsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetBotEdgeRequestsPlatformIntegrationsCountOk() (*int32, bool) {
	if o == nil || o.BotEdgeRequestsPlatformIntegrationsCount == nil {
		return nil, false
	}
	return o.BotEdgeRequestsPlatformIntegrationsCount, true
}

// HasBotEdgeRequestsPlatformIntegrationsCount returns a boolean if a field has been set.
func (o *PlatformValues) HasBotEdgeRequestsPlatformIntegrationsCount() bool {
	if o != nil && o.BotEdgeRequestsPlatformIntegrationsCount != nil {
		return true
	}

	return false
}

// SetBotEdgeRequestsPlatformIntegrationsCount gets a reference to the given int32 and assigns it to the BotEdgeRequestsPlatformIntegrationsCount field.
func (o *PlatformValues) SetBotEdgeRequestsPlatformIntegrationsCount(v int32) {
	o.BotEdgeRequestsPlatformIntegrationsCount = &v
}

// GetBotEdgeRequestsResearchCount returns the BotEdgeRequestsResearchCount field value if set, zero value otherwise.
func (o *PlatformValues) GetBotEdgeRequestsResearchCount() int32 {
	if o == nil || o.BotEdgeRequestsResearchCount == nil {
		var ret int32
		return ret
	}
	return *o.BotEdgeRequestsResearchCount
}

// GetBotEdgeRequestsResearchCountOk returns a tuple with the BotEdgeRequestsResearchCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetBotEdgeRequestsResearchCountOk() (*int32, bool) {
	if o == nil || o.BotEdgeRequestsResearchCount == nil {
		return nil, false
	}
	return o.BotEdgeRequestsResearchCount, true
}

// HasBotEdgeRequestsResearchCount returns a boolean if a field has been set.
func (o *PlatformValues) HasBotEdgeRequestsResearchCount() bool {
	if o != nil && o.BotEdgeRequestsResearchCount != nil {
		return true
	}

	return false
}

// SetBotEdgeRequestsResearchCount gets a reference to the given int32 and assigns it to the BotEdgeRequestsResearchCount field.
func (o *PlatformValues) SetBotEdgeRequestsResearchCount(v int32) {
	o.BotEdgeRequestsResearchCount = &v
}

// GetBotEdgeRequestsSearchEngineCrawlerCount returns the BotEdgeRequestsSearchEngineCrawlerCount field value if set, zero value otherwise.
func (o *PlatformValues) GetBotEdgeRequestsSearchEngineCrawlerCount() int32 {
	if o == nil || o.BotEdgeRequestsSearchEngineCrawlerCount == nil {
		var ret int32
		return ret
	}
	return *o.BotEdgeRequestsSearchEngineCrawlerCount
}

// GetBotEdgeRequestsSearchEngineCrawlerCountOk returns a tuple with the BotEdgeRequestsSearchEngineCrawlerCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetBotEdgeRequestsSearchEngineCrawlerCountOk() (*int32, bool) {
	if o == nil || o.BotEdgeRequestsSearchEngineCrawlerCount == nil {
		return nil, false
	}
	return o.BotEdgeRequestsSearchEngineCrawlerCount, true
}

// HasBotEdgeRequestsSearchEngineCrawlerCount returns a boolean if a field has been set.
func (o *PlatformValues) HasBotEdgeRequestsSearchEngineCrawlerCount() bool {
	if o != nil && o.BotEdgeRequestsSearchEngineCrawlerCount != nil {
		return true
	}

	return false
}

// SetBotEdgeRequestsSearchEngineCrawlerCount gets a reference to the given int32 and assigns it to the BotEdgeRequestsSearchEngineCrawlerCount field.
func (o *PlatformValues) SetBotEdgeRequestsSearchEngineCrawlerCount(v int32) {
	o.BotEdgeRequestsSearchEngineCrawlerCount = &v
}

// GetBotEdgeRequestsSearchEngineOptimizationCount returns the BotEdgeRequestsSearchEngineOptimizationCount field value if set, zero value otherwise.
func (o *PlatformValues) GetBotEdgeRequestsSearchEngineOptimizationCount() int32 {
	if o == nil || o.BotEdgeRequestsSearchEngineOptimizationCount == nil {
		var ret int32
		return ret
	}
	return *o.BotEdgeRequestsSearchEngineOptimizationCount
}

// GetBotEdgeRequestsSearchEngineOptimizationCountOk returns a tuple with the BotEdgeRequestsSearchEngineOptimizationCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetBotEdgeRequestsSearchEngineOptimizationCountOk() (*int32, bool) {
	if o == nil || o.BotEdgeRequestsSearchEngineOptimizationCount == nil {
		return nil, false
	}
	return o.BotEdgeRequestsSearchEngineOptimizationCount, true
}

// HasBotEdgeRequestsSearchEngineOptimizationCount returns a boolean if a field has been set.
func (o *PlatformValues) HasBotEdgeRequestsSearchEngineOptimizationCount() bool {
	if o != nil && o.BotEdgeRequestsSearchEngineOptimizationCount != nil {
		return true
	}

	return false
}

// SetBotEdgeRequestsSearchEngineOptimizationCount gets a reference to the given int32 and assigns it to the BotEdgeRequestsSearchEngineOptimizationCount field.
func (o *PlatformValues) SetBotEdgeRequestsSearchEngineOptimizationCount(v int32) {
	o.BotEdgeRequestsSearchEngineOptimizationCount = &v
}

// GetBotEdgeRequestsSecurityToolsCount returns the BotEdgeRequestsSecurityToolsCount field value if set, zero value otherwise.
func (o *PlatformValues) GetBotEdgeRequestsSecurityToolsCount() int32 {
	if o == nil || o.BotEdgeRequestsSecurityToolsCount == nil {
		var ret int32
		return ret
	}
	return *o.BotEdgeRequestsSecurityToolsCount
}

// GetBotEdgeRequestsSecurityToolsCountOk returns a tuple with the BotEdgeRequestsSecurityToolsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetBotEdgeRequestsSecurityToolsCountOk() (*int32, bool) {
	if o == nil || o.BotEdgeRequestsSecurityToolsCount == nil {
		return nil, false
	}
	return o.BotEdgeRequestsSecurityToolsCount, true
}

// HasBotEdgeRequestsSecurityToolsCount returns a boolean if a field has been set.
func (o *PlatformValues) HasBotEdgeRequestsSecurityToolsCount() bool {
	if o != nil && o.BotEdgeRequestsSecurityToolsCount != nil {
		return true
	}

	return false
}

// SetBotEdgeRequestsSecurityToolsCount gets a reference to the given int32 and assigns it to the BotEdgeRequestsSecurityToolsCount field.
func (o *PlatformValues) SetBotEdgeRequestsSecurityToolsCount(v int32) {
	o.BotEdgeRequestsSecurityToolsCount = &v
}

// GetComputeHandoff returns the ComputeHandoff field value if set, zero value otherwise.
func (o *PlatformValues) GetComputeHandoff() int32 {
	if o == nil || o.ComputeHandoff == nil {
		var ret int32
		return ret
	}
	return *o.ComputeHandoff
}

// GetComputeHandoffOk returns a tuple with the ComputeHandoff field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetComputeHandoffOk() (*int32, bool) {
	if o == nil || o.ComputeHandoff == nil {
		return nil, false
	}
	return o.ComputeHandoff, true
}

// HasComputeHandoff returns a boolean if a field has been set.
func (o *PlatformValues) HasComputeHandoff() bool {
	if o != nil && o.ComputeHandoff != nil {
		return true
	}

	return false
}

// SetComputeHandoff gets a reference to the given int32 and assigns it to the ComputeHandoff field.
func (o *PlatformValues) SetComputeHandoff(v int32) {
	o.ComputeHandoff = &v
}

// GetComputeServiceBereqDNSError returns the ComputeServiceBereqDNSError field value if set, zero value otherwise.
func (o *PlatformValues) GetComputeServiceBereqDNSError() int32 {
	if o == nil || o.ComputeServiceBereqDNSError == nil {
		var ret int32
		return ret
	}
	return *o.ComputeServiceBereqDNSError
}

// GetComputeServiceBereqDNSErrorOk returns a tuple with the ComputeServiceBereqDNSError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetComputeServiceBereqDNSErrorOk() (*int32, bool) {
	if o == nil || o.ComputeServiceBereqDNSError == nil {
		return nil, false
	}
	return o.ComputeServiceBereqDNSError, true
}

// HasComputeServiceBereqDNSError returns a boolean if a field has been set.
func (o *PlatformValues) HasComputeServiceBereqDNSError() bool {
	if o != nil && o.ComputeServiceBereqDNSError != nil {
		return true
	}

	return false
}

// SetComputeServiceBereqDNSError gets a reference to the given int32 and assigns it to the ComputeServiceBereqDNSError field.
func (o *PlatformValues) SetComputeServiceBereqDNSError(v int32) {
	o.ComputeServiceBereqDNSError = &v
}

// GetComputeServiceBereqConnTimeoutError returns the ComputeServiceBereqConnTimeoutError field value if set, zero value otherwise.
func (o *PlatformValues) GetComputeServiceBereqConnTimeoutError() int32 {
	if o == nil || o.ComputeServiceBereqConnTimeoutError == nil {
		var ret int32
		return ret
	}
	return *o.ComputeServiceBereqConnTimeoutError
}

// GetComputeServiceBereqConnTimeoutErrorOk returns a tuple with the ComputeServiceBereqConnTimeoutError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetComputeServiceBereqConnTimeoutErrorOk() (*int32, bool) {
	if o == nil || o.ComputeServiceBereqConnTimeoutError == nil {
		return nil, false
	}
	return o.ComputeServiceBereqConnTimeoutError, true
}

// HasComputeServiceBereqConnTimeoutError returns a boolean if a field has been set.
func (o *PlatformValues) HasComputeServiceBereqConnTimeoutError() bool {
	if o != nil && o.ComputeServiceBereqConnTimeoutError != nil {
		return true
	}

	return false
}

// SetComputeServiceBereqConnTimeoutError gets a reference to the given int32 and assigns it to the ComputeServiceBereqConnTimeoutError field.
func (o *PlatformValues) SetComputeServiceBereqConnTimeoutError(v int32) {
	o.ComputeServiceBereqConnTimeoutError = &v
}

// GetComputeServiceBereqConnRefusedError returns the ComputeServiceBereqConnRefusedError field value if set, zero value otherwise.
func (o *PlatformValues) GetComputeServiceBereqConnRefusedError() int32 {
	if o == nil || o.ComputeServiceBereqConnRefusedError == nil {
		var ret int32
		return ret
	}
	return *o.ComputeServiceBereqConnRefusedError
}

// GetComputeServiceBereqConnRefusedErrorOk returns a tuple with the ComputeServiceBereqConnRefusedError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetComputeServiceBereqConnRefusedErrorOk() (*int32, bool) {
	if o == nil || o.ComputeServiceBereqConnRefusedError == nil {
		return nil, false
	}
	return o.ComputeServiceBereqConnRefusedError, true
}

// HasComputeServiceBereqConnRefusedError returns a boolean if a field has been set.
func (o *PlatformValues) HasComputeServiceBereqConnRefusedError() bool {
	if o != nil && o.ComputeServiceBereqConnRefusedError != nil {
		return true
	}

	return false
}

// SetComputeServiceBereqConnRefusedError gets a reference to the given int32 and assigns it to the ComputeServiceBereqConnRefusedError field.
func (o *PlatformValues) SetComputeServiceBereqConnRefusedError(v int32) {
	o.ComputeServiceBereqConnRefusedError = &v
}

// GetComputeServiceBereqConnOtherError returns the ComputeServiceBereqConnOtherError field value if set, zero value otherwise.
func (o *PlatformValues) GetComputeServiceBereqConnOtherError() int32 {
	if o == nil || o.ComputeServiceBereqConnOtherError == nil {
		var ret int32
		return ret
	}
	return *o.ComputeServiceBereqConnOtherError
}

// GetComputeServiceBereqConnOtherErrorOk returns a tuple with the ComputeServiceBereqConnOtherError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetComputeServiceBereqConnOtherErrorOk() (*int32, bool) {
	if o == nil || o.ComputeServiceBereqConnOtherError == nil {
		return nil, false
	}
	return o.ComputeServiceBereqConnOtherError, true
}

// HasComputeServiceBereqConnOtherError returns a boolean if a field has been set.
func (o *PlatformValues) HasComputeServiceBereqConnOtherError() bool {
	if o != nil && o.ComputeServiceBereqConnOtherError != nil {
		return true
	}

	return false
}

// SetComputeServiceBereqConnOtherError gets a reference to the given int32 and assigns it to the ComputeServiceBereqConnOtherError field.
func (o *PlatformValues) SetComputeServiceBereqConnOtherError(v int32) {
	o.ComputeServiceBereqConnOtherError = &v
}

// GetComputeServiceBereqTLSServerCertError returns the ComputeServiceBereqTLSServerCertError field value if set, zero value otherwise.
func (o *PlatformValues) GetComputeServiceBereqTLSServerCertError() int32 {
	if o == nil || o.ComputeServiceBereqTLSServerCertError == nil {
		var ret int32
		return ret
	}
	return *o.ComputeServiceBereqTLSServerCertError
}

// GetComputeServiceBereqTLSServerCertErrorOk returns a tuple with the ComputeServiceBereqTLSServerCertError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetComputeServiceBereqTLSServerCertErrorOk() (*int32, bool) {
	if o == nil || o.ComputeServiceBereqTLSServerCertError == nil {
		return nil, false
	}
	return o.ComputeServiceBereqTLSServerCertError, true
}

// HasComputeServiceBereqTLSServerCertError returns a boolean if a field has been set.
func (o *PlatformValues) HasComputeServiceBereqTLSServerCertError() bool {
	if o != nil && o.ComputeServiceBereqTLSServerCertError != nil {
		return true
	}

	return false
}

// SetComputeServiceBereqTLSServerCertError gets a reference to the given int32 and assigns it to the ComputeServiceBereqTLSServerCertError field.
func (o *PlatformValues) SetComputeServiceBereqTLSServerCertError(v int32) {
	o.ComputeServiceBereqTLSServerCertError = &v
}

// GetComputeServiceBereqTLSOtherError returns the ComputeServiceBereqTLSOtherError field value if set, zero value otherwise.
func (o *PlatformValues) GetComputeServiceBereqTLSOtherError() int32 {
	if o == nil || o.ComputeServiceBereqTLSOtherError == nil {
		var ret int32
		return ret
	}
	return *o.ComputeServiceBereqTLSOtherError
}

// GetComputeServiceBereqTLSOtherErrorOk returns a tuple with the ComputeServiceBereqTLSOtherError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetComputeServiceBereqTLSOtherErrorOk() (*int32, bool) {
	if o == nil || o.ComputeServiceBereqTLSOtherError == nil {
		return nil, false
	}
	return o.ComputeServiceBereqTLSOtherError, true
}

// HasComputeServiceBereqTLSOtherError returns a boolean if a field has been set.
func (o *PlatformValues) HasComputeServiceBereqTLSOtherError() bool {
	if o != nil && o.ComputeServiceBereqTLSOtherError != nil {
		return true
	}

	return false
}

// SetComputeServiceBereqTLSOtherError gets a reference to the given int32 and assigns it to the ComputeServiceBereqTLSOtherError field.
func (o *PlatformValues) SetComputeServiceBereqTLSOtherError(v int32) {
	o.ComputeServiceBereqTLSOtherError = &v
}

// GetComputeServiceBereqHTTPProtoV1Error returns the ComputeServiceBereqHTTPProtoV1Error field value if set, zero value otherwise.
func (o *PlatformValues) GetComputeServiceBereqHTTPProtoV1Error() int32 {
	if o == nil || o.ComputeServiceBereqHTTPProtoV1Error == nil {
		var ret int32
		return ret
	}
	return *o.ComputeServiceBereqHTTPProtoV1Error
}

// GetComputeServiceBereqHTTPProtoV1ErrorOk returns a tuple with the ComputeServiceBereqHTTPProtoV1Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetComputeServiceBereqHTTPProtoV1ErrorOk() (*int32, bool) {
	if o == nil || o.ComputeServiceBereqHTTPProtoV1Error == nil {
		return nil, false
	}
	return o.ComputeServiceBereqHTTPProtoV1Error, true
}

// HasComputeServiceBereqHTTPProtoV1Error returns a boolean if a field has been set.
func (o *PlatformValues) HasComputeServiceBereqHTTPProtoV1Error() bool {
	if o != nil && o.ComputeServiceBereqHTTPProtoV1Error != nil {
		return true
	}

	return false
}

// SetComputeServiceBereqHTTPProtoV1Error gets a reference to the given int32 and assigns it to the ComputeServiceBereqHTTPProtoV1Error field.
func (o *PlatformValues) SetComputeServiceBereqHTTPProtoV1Error(v int32) {
	o.ComputeServiceBereqHTTPProtoV1Error = &v
}

// GetComputeServiceBereqHTTPProtoV2Error returns the ComputeServiceBereqHTTPProtoV2Error field value if set, zero value otherwise.
func (o *PlatformValues) GetComputeServiceBereqHTTPProtoV2Error() int32 {
	if o == nil || o.ComputeServiceBereqHTTPProtoV2Error == nil {
		var ret int32
		return ret
	}
	return *o.ComputeServiceBereqHTTPProtoV2Error
}

// GetComputeServiceBereqHTTPProtoV2ErrorOk returns a tuple with the ComputeServiceBereqHTTPProtoV2Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetComputeServiceBereqHTTPProtoV2ErrorOk() (*int32, bool) {
	if o == nil || o.ComputeServiceBereqHTTPProtoV2Error == nil {
		return nil, false
	}
	return o.ComputeServiceBereqHTTPProtoV2Error, true
}

// HasComputeServiceBereqHTTPProtoV2Error returns a boolean if a field has been set.
func (o *PlatformValues) HasComputeServiceBereqHTTPProtoV2Error() bool {
	if o != nil && o.ComputeServiceBereqHTTPProtoV2Error != nil {
		return true
	}

	return false
}

// SetComputeServiceBereqHTTPProtoV2Error gets a reference to the given int32 and assigns it to the ComputeServiceBereqHTTPProtoV2Error field.
func (o *PlatformValues) SetComputeServiceBereqHTTPProtoV2Error(v int32) {
	o.ComputeServiceBereqHTTPProtoV2Error = &v
}

// GetComputeServiceBereqHTTPIncompleteError returns the ComputeServiceBereqHTTPIncompleteError field value if set, zero value otherwise.
func (o *PlatformValues) GetComputeServiceBereqHTTPIncompleteError() int32 {
	if o == nil || o.ComputeServiceBereqHTTPIncompleteError == nil {
		var ret int32
		return ret
	}
	return *o.ComputeServiceBereqHTTPIncompleteError
}

// GetComputeServiceBereqHTTPIncompleteErrorOk returns a tuple with the ComputeServiceBereqHTTPIncompleteError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetComputeServiceBereqHTTPIncompleteErrorOk() (*int32, bool) {
	if o == nil || o.ComputeServiceBereqHTTPIncompleteError == nil {
		return nil, false
	}
	return o.ComputeServiceBereqHTTPIncompleteError, true
}

// HasComputeServiceBereqHTTPIncompleteError returns a boolean if a field has been set.
func (o *PlatformValues) HasComputeServiceBereqHTTPIncompleteError() bool {
	if o != nil && o.ComputeServiceBereqHTTPIncompleteError != nil {
		return true
	}

	return false
}

// SetComputeServiceBereqHTTPIncompleteError gets a reference to the given int32 and assigns it to the ComputeServiceBereqHTTPIncompleteError field.
func (o *PlatformValues) SetComputeServiceBereqHTTPIncompleteError(v int32) {
	o.ComputeServiceBereqHTTPIncompleteError = &v
}

// GetComputeServiceBereqHTTPTimeoutError returns the ComputeServiceBereqHTTPTimeoutError field value if set, zero value otherwise.
func (o *PlatformValues) GetComputeServiceBereqHTTPTimeoutError() int32 {
	if o == nil || o.ComputeServiceBereqHTTPTimeoutError == nil {
		var ret int32
		return ret
	}
	return *o.ComputeServiceBereqHTTPTimeoutError
}

// GetComputeServiceBereqHTTPTimeoutErrorOk returns a tuple with the ComputeServiceBereqHTTPTimeoutError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetComputeServiceBereqHTTPTimeoutErrorOk() (*int32, bool) {
	if o == nil || o.ComputeServiceBereqHTTPTimeoutError == nil {
		return nil, false
	}
	return o.ComputeServiceBereqHTTPTimeoutError, true
}

// HasComputeServiceBereqHTTPTimeoutError returns a boolean if a field has been set.
func (o *PlatformValues) HasComputeServiceBereqHTTPTimeoutError() bool {
	if o != nil && o.ComputeServiceBereqHTTPTimeoutError != nil {
		return true
	}

	return false
}

// SetComputeServiceBereqHTTPTimeoutError gets a reference to the given int32 and assigns it to the ComputeServiceBereqHTTPTimeoutError field.
func (o *PlatformValues) SetComputeServiceBereqHTTPTimeoutError(v int32) {
	o.ComputeServiceBereqHTTPTimeoutError = &v
}

// GetComputeServiceBereqHTTPOtherError returns the ComputeServiceBereqHTTPOtherError field value if set, zero value otherwise.
func (o *PlatformValues) GetComputeServiceBereqHTTPOtherError() int32 {
	if o == nil || o.ComputeServiceBereqHTTPOtherError == nil {
		var ret int32
		return ret
	}
	return *o.ComputeServiceBereqHTTPOtherError
}

// GetComputeServiceBereqHTTPOtherErrorOk returns a tuple with the ComputeServiceBereqHTTPOtherError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetComputeServiceBereqHTTPOtherErrorOk() (*int32, bool) {
	if o == nil || o.ComputeServiceBereqHTTPOtherError == nil {
		return nil, false
	}
	return o.ComputeServiceBereqHTTPOtherError, true
}

// HasComputeServiceBereqHTTPOtherError returns a boolean if a field has been set.
func (o *PlatformValues) HasComputeServiceBereqHTTPOtherError() bool {
	if o != nil && o.ComputeServiceBereqHTTPOtherError != nil {
		return true
	}

	return false
}

// SetComputeServiceBereqHTTPOtherError gets a reference to the given int32 and assigns it to the ComputeServiceBereqHTTPOtherError field.
func (o *PlatformValues) SetComputeServiceBereqHTTPOtherError(v int32) {
	o.ComputeServiceBereqHTTPOtherError = &v
}

// GetComputeServiceBereqOtherError returns the ComputeServiceBereqOtherError field value if set, zero value otherwise.
func (o *PlatformValues) GetComputeServiceBereqOtherError() int32 {
	if o == nil || o.ComputeServiceBereqOtherError == nil {
		var ret int32
		return ret
	}
	return *o.ComputeServiceBereqOtherError
}

// GetComputeServiceBereqOtherErrorOk returns a tuple with the ComputeServiceBereqOtherError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetComputeServiceBereqOtherErrorOk() (*int32, bool) {
	if o == nil || o.ComputeServiceBereqOtherError == nil {
		return nil, false
	}
	return o.ComputeServiceBereqOtherError, true
}

// HasComputeServiceBereqOtherError returns a boolean if a field has been set.
func (o *PlatformValues) HasComputeServiceBereqOtherError() bool {
	if o != nil && o.ComputeServiceBereqOtherError != nil {
		return true
	}

	return false
}

// SetComputeServiceBereqOtherError gets a reference to the given int32 and assigns it to the ComputeServiceBereqOtherError field.
func (o *PlatformValues) SetComputeServiceBereqOtherError(v int32) {
	o.ComputeServiceBereqOtherError = &v
}

// GetComputeServiceBereq5xxError returns the ComputeServiceBereq5xxError field value if set, zero value otherwise.
func (o *PlatformValues) GetComputeServiceBereq5xxError() int32 {
	if o == nil || o.ComputeServiceBereq5xxError == nil {
		var ret int32
		return ret
	}
	return *o.ComputeServiceBereq5xxError
}

// GetComputeServiceBereq5xxErrorOk returns a tuple with the ComputeServiceBereq5xxError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetComputeServiceBereq5xxErrorOk() (*int32, bool) {
	if o == nil || o.ComputeServiceBereq5xxError == nil {
		return nil, false
	}
	return o.ComputeServiceBereq5xxError, true
}

// HasComputeServiceBereq5xxError returns a boolean if a field has been set.
func (o *PlatformValues) HasComputeServiceBereq5xxError() bool {
	if o != nil && o.ComputeServiceBereq5xxError != nil {
		return true
	}

	return false
}

// SetComputeServiceBereq5xxError gets a reference to the given int32 and assigns it to the ComputeServiceBereq5xxError field.
func (o *PlatformValues) SetComputeServiceBereq5xxError(v int32) {
	o.ComputeServiceBereq5xxError = &v
}

// GetComputeServiceBereqConnError returns the ComputeServiceBereqConnError field value if set, zero value otherwise.
func (o *PlatformValues) GetComputeServiceBereqConnError() int32 {
	if o == nil || o.ComputeServiceBereqConnError == nil {
		var ret int32
		return ret
	}
	return *o.ComputeServiceBereqConnError
}

// GetComputeServiceBereqConnErrorOk returns a tuple with the ComputeServiceBereqConnError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetComputeServiceBereqConnErrorOk() (*int32, bool) {
	if o == nil || o.ComputeServiceBereqConnError == nil {
		return nil, false
	}
	return o.ComputeServiceBereqConnError, true
}

// HasComputeServiceBereqConnError returns a boolean if a field has been set.
func (o *PlatformValues) HasComputeServiceBereqConnError() bool {
	if o != nil && o.ComputeServiceBereqConnError != nil {
		return true
	}

	return false
}

// SetComputeServiceBereqConnError gets a reference to the given int32 and assigns it to the ComputeServiceBereqConnError field.
func (o *PlatformValues) SetComputeServiceBereqConnError(v int32) {
	o.ComputeServiceBereqConnError = &v
}

// GetComputeServiceBereqTLSError returns the ComputeServiceBereqTLSError field value if set, zero value otherwise.
func (o *PlatformValues) GetComputeServiceBereqTLSError() int32 {
	if o == nil || o.ComputeServiceBereqTLSError == nil {
		var ret int32
		return ret
	}
	return *o.ComputeServiceBereqTLSError
}

// GetComputeServiceBereqTLSErrorOk returns a tuple with the ComputeServiceBereqTLSError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetComputeServiceBereqTLSErrorOk() (*int32, bool) {
	if o == nil || o.ComputeServiceBereqTLSError == nil {
		return nil, false
	}
	return o.ComputeServiceBereqTLSError, true
}

// HasComputeServiceBereqTLSError returns a boolean if a field has been set.
func (o *PlatformValues) HasComputeServiceBereqTLSError() bool {
	if o != nil && o.ComputeServiceBereqTLSError != nil {
		return true
	}

	return false
}

// SetComputeServiceBereqTLSError gets a reference to the given int32 and assigns it to the ComputeServiceBereqTLSError field.
func (o *PlatformValues) SetComputeServiceBereqTLSError(v int32) {
	o.ComputeServiceBereqTLSError = &v
}

// GetComputeServiceBereqHTTPError returns the ComputeServiceBereqHTTPError field value if set, zero value otherwise.
func (o *PlatformValues) GetComputeServiceBereqHTTPError() int32 {
	if o == nil || o.ComputeServiceBereqHTTPError == nil {
		var ret int32
		return ret
	}
	return *o.ComputeServiceBereqHTTPError
}

// GetComputeServiceBereqHTTPErrorOk returns a tuple with the ComputeServiceBereqHTTPError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetComputeServiceBereqHTTPErrorOk() (*int32, bool) {
	if o == nil || o.ComputeServiceBereqHTTPError == nil {
		return nil, false
	}
	return o.ComputeServiceBereqHTTPError, true
}

// HasComputeServiceBereqHTTPError returns a boolean if a field has been set.
func (o *PlatformValues) HasComputeServiceBereqHTTPError() bool {
	if o != nil && o.ComputeServiceBereqHTTPError != nil {
		return true
	}

	return false
}

// SetComputeServiceBereqHTTPError gets a reference to the given int32 and assigns it to the ComputeServiceBereqHTTPError field.
func (o *PlatformValues) SetComputeServiceBereqHTTPError(v int32) {
	o.ComputeServiceBereqHTTPError = &v
}

// GetBotChallengesPatsIssued returns the BotChallengesPatsIssued field value if set, zero value otherwise.
func (o *PlatformValues) GetBotChallengesPatsIssued() int32 {
	if o == nil || o.BotChallengesPatsIssued == nil {
		var ret int32
		return ret
	}
	return *o.BotChallengesPatsIssued
}

// GetBotChallengesPatsIssuedOk returns a tuple with the BotChallengesPatsIssued field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetBotChallengesPatsIssuedOk() (*int32, bool) {
	if o == nil || o.BotChallengesPatsIssued == nil {
		return nil, false
	}
	return o.BotChallengesPatsIssued, true
}

// HasBotChallengesPatsIssued returns a boolean if a field has been set.
func (o *PlatformValues) HasBotChallengesPatsIssued() bool {
	if o != nil && o.BotChallengesPatsIssued != nil {
		return true
	}

	return false
}

// SetBotChallengesPatsIssued gets a reference to the given int32 and assigns it to the BotChallengesPatsIssued field.
func (o *PlatformValues) SetBotChallengesPatsIssued(v int32) {
	o.BotChallengesPatsIssued = &v
}

// GetBotChallengesPatsSucceeded returns the BotChallengesPatsSucceeded field value if set, zero value otherwise.
func (o *PlatformValues) GetBotChallengesPatsSucceeded() int32 {
	if o == nil || o.BotChallengesPatsSucceeded == nil {
		var ret int32
		return ret
	}
	return *o.BotChallengesPatsSucceeded
}

// GetBotChallengesPatsSucceededOk returns a tuple with the BotChallengesPatsSucceeded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetBotChallengesPatsSucceededOk() (*int32, bool) {
	if o == nil || o.BotChallengesPatsSucceeded == nil {
		return nil, false
	}
	return o.BotChallengesPatsSucceeded, true
}

// HasBotChallengesPatsSucceeded returns a boolean if a field has been set.
func (o *PlatformValues) HasBotChallengesPatsSucceeded() bool {
	if o != nil && o.BotChallengesPatsSucceeded != nil {
		return true
	}

	return false
}

// SetBotChallengesPatsSucceeded gets a reference to the given int32 and assigns it to the BotChallengesPatsSucceeded field.
func (o *PlatformValues) SetBotChallengesPatsSucceeded(v int32) {
	o.BotChallengesPatsSucceeded = &v
}

// GetBotEdgeRequestsHeadlessCount returns the BotEdgeRequestsHeadlessCount field value if set, zero value otherwise.
func (o *PlatformValues) GetBotEdgeRequestsHeadlessCount() int32 {
	if o == nil || o.BotEdgeRequestsHeadlessCount == nil {
		var ret int32
		return ret
	}
	return *o.BotEdgeRequestsHeadlessCount
}

// GetBotEdgeRequestsHeadlessCountOk returns a tuple with the BotEdgeRequestsHeadlessCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetBotEdgeRequestsHeadlessCountOk() (*int32, bool) {
	if o == nil || o.BotEdgeRequestsHeadlessCount == nil {
		return nil, false
	}
	return o.BotEdgeRequestsHeadlessCount, true
}

// HasBotEdgeRequestsHeadlessCount returns a boolean if a field has been set.
func (o *PlatformValues) HasBotEdgeRequestsHeadlessCount() bool {
	if o != nil && o.BotEdgeRequestsHeadlessCount != nil {
		return true
	}

	return false
}

// SetBotEdgeRequestsHeadlessCount gets a reference to the given int32 and assigns it to the BotEdgeRequestsHeadlessCount field.
func (o *PlatformValues) SetBotEdgeRequestsHeadlessCount(v int32) {
	o.BotEdgeRequestsHeadlessCount = &v
}

// GetStatus499 returns the Status499 field value if set, zero value otherwise.
func (o *PlatformValues) GetStatus499() int32 {
	if o == nil || o.Status499 == nil {
		var ret int32
		return ret
	}
	return *o.Status499
}

// GetStatus499Ok returns a tuple with the Status499 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetStatus499Ok() (*int32, bool) {
	if o == nil || o.Status499 == nil {
		return nil, false
	}
	return o.Status499, true
}

// HasStatus499 returns a boolean if a field has been set.
func (o *PlatformValues) HasStatus499() bool {
	if o != nil && o.Status499 != nil {
		return true
	}

	return false
}

// SetStatus499 gets a reference to the given int32 and assigns it to the Status499 field.
func (o *PlatformValues) SetStatus499(v int32) {
	o.Status499 = &v
}

// GetArcRequests returns the ArcRequests field value if set, zero value otherwise.
func (o *PlatformValues) GetArcRequests() int32 {
	if o == nil || o.ArcRequests == nil {
		var ret int32
		return ret
	}
	return *o.ArcRequests
}

// GetArcRequestsOk returns a tuple with the ArcRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetArcRequestsOk() (*int32, bool) {
	if o == nil || o.ArcRequests == nil {
		return nil, false
	}
	return o.ArcRequests, true
}

// HasArcRequests returns a boolean if a field has been set.
func (o *PlatformValues) HasArcRequests() bool {
	if o != nil && o.ArcRequests != nil {
		return true
	}

	return false
}

// SetArcRequests gets a reference to the given int32 and assigns it to the ArcRequests field.
func (o *PlatformValues) SetArcRequests(v int32) {
	o.ArcRequests = &v
}

// GetDdosChallengeCompleteTokensChecked returns the DdosChallengeCompleteTokensChecked field value if set, zero value otherwise.
func (o *PlatformValues) GetDdosChallengeCompleteTokensChecked() int32 {
	if o == nil || o.DdosChallengeCompleteTokensChecked == nil {
		var ret int32
		return ret
	}
	return *o.DdosChallengeCompleteTokensChecked
}

// GetDdosChallengeCompleteTokensCheckedOk returns a tuple with the DdosChallengeCompleteTokensChecked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetDdosChallengeCompleteTokensCheckedOk() (*int32, bool) {
	if o == nil || o.DdosChallengeCompleteTokensChecked == nil {
		return nil, false
	}
	return o.DdosChallengeCompleteTokensChecked, true
}

// HasDdosChallengeCompleteTokensChecked returns a boolean if a field has been set.
func (o *PlatformValues) HasDdosChallengeCompleteTokensChecked() bool {
	if o != nil && o.DdosChallengeCompleteTokensChecked != nil {
		return true
	}

	return false
}

// SetDdosChallengeCompleteTokensChecked gets a reference to the given int32 and assigns it to the DdosChallengeCompleteTokensChecked field.
func (o *PlatformValues) SetDdosChallengeCompleteTokensChecked(v int32) {
	o.DdosChallengeCompleteTokensChecked = &v
}

// GetDdosChallengeCompleteTokensPassed returns the DdosChallengeCompleteTokensPassed field value if set, zero value otherwise.
func (o *PlatformValues) GetDdosChallengeCompleteTokensPassed() int32 {
	if o == nil || o.DdosChallengeCompleteTokensPassed == nil {
		var ret int32
		return ret
	}
	return *o.DdosChallengeCompleteTokensPassed
}

// GetDdosChallengeCompleteTokensPassedOk returns a tuple with the DdosChallengeCompleteTokensPassed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetDdosChallengeCompleteTokensPassedOk() (*int32, bool) {
	if o == nil || o.DdosChallengeCompleteTokensPassed == nil {
		return nil, false
	}
	return o.DdosChallengeCompleteTokensPassed, true
}

// HasDdosChallengeCompleteTokensPassed returns a boolean if a field has been set.
func (o *PlatformValues) HasDdosChallengeCompleteTokensPassed() bool {
	if o != nil && o.DdosChallengeCompleteTokensPassed != nil {
		return true
	}

	return false
}

// SetDdosChallengeCompleteTokensPassed gets a reference to the given int32 and assigns it to the DdosChallengeCompleteTokensPassed field.
func (o *PlatformValues) SetDdosChallengeCompleteTokensPassed(v int32) {
	o.DdosChallengeCompleteTokensPassed = &v
}

// GetDdosChallengeCompleteTokensFailed returns the DdosChallengeCompleteTokensFailed field value if set, zero value otherwise.
func (o *PlatformValues) GetDdosChallengeCompleteTokensFailed() int32 {
	if o == nil || o.DdosChallengeCompleteTokensFailed == nil {
		var ret int32
		return ret
	}
	return *o.DdosChallengeCompleteTokensFailed
}

// GetDdosChallengeCompleteTokensFailedOk returns a tuple with the DdosChallengeCompleteTokensFailed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetDdosChallengeCompleteTokensFailedOk() (*int32, bool) {
	if o == nil || o.DdosChallengeCompleteTokensFailed == nil {
		return nil, false
	}
	return o.DdosChallengeCompleteTokensFailed, true
}

// HasDdosChallengeCompleteTokensFailed returns a boolean if a field has been set.
func (o *PlatformValues) HasDdosChallengeCompleteTokensFailed() bool {
	if o != nil && o.DdosChallengeCompleteTokensFailed != nil {
		return true
	}

	return false
}

// SetDdosChallengeCompleteTokensFailed gets a reference to the given int32 and assigns it to the DdosChallengeCompleteTokensFailed field.
func (o *PlatformValues) SetDdosChallengeCompleteTokensFailed(v int32) {
	o.DdosChallengeCompleteTokensFailed = &v
}

// GetDdosChallengeStarts returns the DdosChallengeStarts field value if set, zero value otherwise.
func (o *PlatformValues) GetDdosChallengeStarts() int32 {
	if o == nil || o.DdosChallengeStarts == nil {
		var ret int32
		return ret
	}
	return *o.DdosChallengeStarts
}

// GetDdosChallengeStartsOk returns a tuple with the DdosChallengeStarts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetDdosChallengeStartsOk() (*int32, bool) {
	if o == nil || o.DdosChallengeStarts == nil {
		return nil, false
	}
	return o.DdosChallengeStarts, true
}

// HasDdosChallengeStarts returns a boolean if a field has been set.
func (o *PlatformValues) HasDdosChallengeStarts() bool {
	if o != nil && o.DdosChallengeStarts != nil {
		return true
	}

	return false
}

// SetDdosChallengeStarts gets a reference to the given int32 and assigns it to the DdosChallengeStarts field.
func (o *PlatformValues) SetDdosChallengeStarts(v int32) {
	o.DdosChallengeStarts = &v
}

// GetDdosChallengeCompleteTokensIssued returns the DdosChallengeCompleteTokensIssued field value if set, zero value otherwise.
func (o *PlatformValues) GetDdosChallengeCompleteTokensIssued() int32 {
	if o == nil || o.DdosChallengeCompleteTokensIssued == nil {
		var ret int32
		return ret
	}
	return *o.DdosChallengeCompleteTokensIssued
}

// GetDdosChallengeCompleteTokensIssuedOk returns a tuple with the DdosChallengeCompleteTokensIssued field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetDdosChallengeCompleteTokensIssuedOk() (*int32, bool) {
	if o == nil || o.DdosChallengeCompleteTokensIssued == nil {
		return nil, false
	}
	return o.DdosChallengeCompleteTokensIssued, true
}

// HasDdosChallengeCompleteTokensIssued returns a boolean if a field has been set.
func (o *PlatformValues) HasDdosChallengeCompleteTokensIssued() bool {
	if o != nil && o.DdosChallengeCompleteTokensIssued != nil {
		return true
	}

	return false
}

// SetDdosChallengeCompleteTokensIssued gets a reference to the given int32 and assigns it to the DdosChallengeCompleteTokensIssued field.
func (o *PlatformValues) SetDdosChallengeCompleteTokensIssued(v int32) {
	o.DdosChallengeCompleteTokensIssued = &v
}

// GetDdosChallengesIssued returns the DdosChallengesIssued field value if set, zero value otherwise.
func (o *PlatformValues) GetDdosChallengesIssued() int32 {
	if o == nil || o.DdosChallengesIssued == nil {
		var ret int32
		return ret
	}
	return *o.DdosChallengesIssued
}

// GetDdosChallengesIssuedOk returns a tuple with the DdosChallengesIssued field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetDdosChallengesIssuedOk() (*int32, bool) {
	if o == nil || o.DdosChallengesIssued == nil {
		return nil, false
	}
	return o.DdosChallengesIssued, true
}

// HasDdosChallengesIssued returns a boolean if a field has been set.
func (o *PlatformValues) HasDdosChallengesIssued() bool {
	if o != nil && o.DdosChallengesIssued != nil {
		return true
	}

	return false
}

// SetDdosChallengesIssued gets a reference to the given int32 and assigns it to the DdosChallengesIssued field.
func (o *PlatformValues) SetDdosChallengesIssued(v int32) {
	o.DdosChallengesIssued = &v
}

// GetDdosChallengesSucceeded returns the DdosChallengesSucceeded field value if set, zero value otherwise.
func (o *PlatformValues) GetDdosChallengesSucceeded() int32 {
	if o == nil || o.DdosChallengesSucceeded == nil {
		var ret int32
		return ret
	}
	return *o.DdosChallengesSucceeded
}

// GetDdosChallengesSucceededOk returns a tuple with the DdosChallengesSucceeded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetDdosChallengesSucceededOk() (*int32, bool) {
	if o == nil || o.DdosChallengesSucceeded == nil {
		return nil, false
	}
	return o.DdosChallengesSucceeded, true
}

// HasDdosChallengesSucceeded returns a boolean if a field has been set.
func (o *PlatformValues) HasDdosChallengesSucceeded() bool {
	if o != nil && o.DdosChallengesSucceeded != nil {
		return true
	}

	return false
}

// SetDdosChallengesSucceeded gets a reference to the given int32 and assigns it to the DdosChallengesSucceeded field.
func (o *PlatformValues) SetDdosChallengesSucceeded(v int32) {
	o.DdosChallengesSucceeded = &v
}

// GetDdosChallengesFailed returns the DdosChallengesFailed field value if set, zero value otherwise.
func (o *PlatformValues) GetDdosChallengesFailed() int32 {
	if o == nil || o.DdosChallengesFailed == nil {
		var ret int32
		return ret
	}
	return *o.DdosChallengesFailed
}

// GetDdosChallengesFailedOk returns a tuple with the DdosChallengesFailed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetDdosChallengesFailedOk() (*int32, bool) {
	if o == nil || o.DdosChallengesFailed == nil {
		return nil, false
	}
	return o.DdosChallengesFailed, true
}

// HasDdosChallengesFailed returns a boolean if a field has been set.
func (o *PlatformValues) HasDdosChallengesFailed() bool {
	if o != nil && o.DdosChallengesFailed != nil {
		return true
	}

	return false
}

// SetDdosChallengesFailed gets a reference to the given int32 and assigns it to the DdosChallengesFailed field.
func (o *PlatformValues) SetDdosChallengesFailed(v int32) {
	o.DdosChallengesFailed = &v
}

// GetDdosChallengesPatsIssued returns the DdosChallengesPatsIssued field value if set, zero value otherwise.
func (o *PlatformValues) GetDdosChallengesPatsIssued() int32 {
	if o == nil || o.DdosChallengesPatsIssued == nil {
		var ret int32
		return ret
	}
	return *o.DdosChallengesPatsIssued
}

// GetDdosChallengesPatsIssuedOk returns a tuple with the DdosChallengesPatsIssued field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetDdosChallengesPatsIssuedOk() (*int32, bool) {
	if o == nil || o.DdosChallengesPatsIssued == nil {
		return nil, false
	}
	return o.DdosChallengesPatsIssued, true
}

// HasDdosChallengesPatsIssued returns a boolean if a field has been set.
func (o *PlatformValues) HasDdosChallengesPatsIssued() bool {
	if o != nil && o.DdosChallengesPatsIssued != nil {
		return true
	}

	return false
}

// SetDdosChallengesPatsIssued gets a reference to the given int32 and assigns it to the DdosChallengesPatsIssued field.
func (o *PlatformValues) SetDdosChallengesPatsIssued(v int32) {
	o.DdosChallengesPatsIssued = &v
}

// GetDdosChallengesPatsSucceeded returns the DdosChallengesPatsSucceeded field value if set, zero value otherwise.
func (o *PlatformValues) GetDdosChallengesPatsSucceeded() int32 {
	if o == nil || o.DdosChallengesPatsSucceeded == nil {
		var ret int32
		return ret
	}
	return *o.DdosChallengesPatsSucceeded
}

// GetDdosChallengesPatsSucceededOk returns a tuple with the DdosChallengesPatsSucceeded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetDdosChallengesPatsSucceededOk() (*int32, bool) {
	if o == nil || o.DdosChallengesPatsSucceeded == nil {
		return nil, false
	}
	return o.DdosChallengesPatsSucceeded, true
}

// HasDdosChallengesPatsSucceeded returns a boolean if a field has been set.
func (o *PlatformValues) HasDdosChallengesPatsSucceeded() bool {
	if o != nil && o.DdosChallengesPatsSucceeded != nil {
		return true
	}

	return false
}

// SetDdosChallengesPatsSucceeded gets a reference to the given int32 and assigns it to the DdosChallengesPatsSucceeded field.
func (o *PlatformValues) SetDdosChallengesPatsSucceeded(v int32) {
	o.DdosChallengesPatsSucceeded = &v
}

// GetCspPagesGauge returns the CspPagesGauge field value if set, zero value otherwise.
func (o *PlatformValues) GetCspPagesGauge() int64 {
	if o == nil || o.CspPagesGauge == nil {
		var ret int64
		return ret
	}
	return *o.CspPagesGauge
}

// GetCspPagesGaugeOk returns a tuple with the CspPagesGauge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetCspPagesGaugeOk() (*int64, bool) {
	if o == nil || o.CspPagesGauge == nil {
		return nil, false
	}
	return o.CspPagesGauge, true
}

// HasCspPagesGauge returns a boolean if a field has been set.
func (o *PlatformValues) HasCspPagesGauge() bool {
	if o != nil && o.CspPagesGauge != nil {
		return true
	}

	return false
}

// SetCspPagesGauge gets a reference to the given int64 and assigns it to the CspPagesGauge field.
func (o *PlatformValues) SetCspPagesGauge(v int64) {
	o.CspPagesGauge = &v
}

// GetNgwafAccountRulesGauge returns the NgwafAccountRulesGauge field value if set, zero value otherwise.
func (o *PlatformValues) GetNgwafAccountRulesGauge() int64 {
	if o == nil || o.NgwafAccountRulesGauge == nil {
		var ret int64
		return ret
	}
	return *o.NgwafAccountRulesGauge
}

// GetNgwafAccountRulesGaugeOk returns a tuple with the NgwafAccountRulesGauge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetNgwafAccountRulesGaugeOk() (*int64, bool) {
	if o == nil || o.NgwafAccountRulesGauge == nil {
		return nil, false
	}
	return o.NgwafAccountRulesGauge, true
}

// HasNgwafAccountRulesGauge returns a boolean if a field has been set.
func (o *PlatformValues) HasNgwafAccountRulesGauge() bool {
	if o != nil && o.NgwafAccountRulesGauge != nil {
		return true
	}

	return false
}

// SetNgwafAccountRulesGauge gets a reference to the given int64 and assigns it to the NgwafAccountRulesGauge field.
func (o *PlatformValues) SetNgwafAccountRulesGauge(v int64) {
	o.NgwafAccountRulesGauge = &v
}

// GetNgwafAccountWorkspaces returns the NgwafAccountWorkspaces field value if set, zero value otherwise.
func (o *PlatformValues) GetNgwafAccountWorkspaces() int64 {
	if o == nil || o.NgwafAccountWorkspaces == nil {
		var ret int64
		return ret
	}
	return *o.NgwafAccountWorkspaces
}

// GetNgwafAccountWorkspacesOk returns a tuple with the NgwafAccountWorkspaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetNgwafAccountWorkspacesOk() (*int64, bool) {
	if o == nil || o.NgwafAccountWorkspaces == nil {
		return nil, false
	}
	return o.NgwafAccountWorkspaces, true
}

// HasNgwafAccountWorkspaces returns a boolean if a field has been set.
func (o *PlatformValues) HasNgwafAccountWorkspaces() bool {
	if o != nil && o.NgwafAccountWorkspaces != nil {
		return true
	}

	return false
}

// SetNgwafAccountWorkspaces gets a reference to the given int64 and assigns it to the NgwafAccountWorkspaces field.
func (o *PlatformValues) SetNgwafAccountWorkspaces(v int64) {
	o.NgwafAccountWorkspaces = &v
}

// GetNgwafWorkspaceRulesGauge returns the NgwafWorkspaceRulesGauge field value if set, zero value otherwise.
func (o *PlatformValues) GetNgwafWorkspaceRulesGauge() int64 {
	if o == nil || o.NgwafWorkspaceRulesGauge == nil {
		var ret int64
		return ret
	}
	return *o.NgwafWorkspaceRulesGauge
}

// GetNgwafWorkspaceRulesGaugeOk returns a tuple with the NgwafWorkspaceRulesGauge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetNgwafWorkspaceRulesGaugeOk() (*int64, bool) {
	if o == nil || o.NgwafWorkspaceRulesGauge == nil {
		return nil, false
	}
	return o.NgwafWorkspaceRulesGauge, true
}

// HasNgwafWorkspaceRulesGauge returns a boolean if a field has been set.
func (o *PlatformValues) HasNgwafWorkspaceRulesGauge() bool {
	if o != nil && o.NgwafWorkspaceRulesGauge != nil {
		return true
	}

	return false
}

// SetNgwafWorkspaceRulesGauge gets a reference to the given int64 and assigns it to the NgwafWorkspaceRulesGauge field.
func (o *PlatformValues) SetNgwafWorkspaceRulesGauge(v int64) {
	o.NgwafWorkspaceRulesGauge = &v
}

// GetNgwafWorkspaceThresholdsGauge returns the NgwafWorkspaceThresholdsGauge field value if set, zero value otherwise.
func (o *PlatformValues) GetNgwafWorkspaceThresholdsGauge() int64 {
	if o == nil || o.NgwafWorkspaceThresholdsGauge == nil {
		var ret int64
		return ret
	}
	return *o.NgwafWorkspaceThresholdsGauge
}

// GetNgwafWorkspaceThresholdsGaugeOk returns a tuple with the NgwafWorkspaceThresholdsGauge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetNgwafWorkspaceThresholdsGaugeOk() (*int64, bool) {
	if o == nil || o.NgwafWorkspaceThresholdsGauge == nil {
		return nil, false
	}
	return o.NgwafWorkspaceThresholdsGauge, true
}

// HasNgwafWorkspaceThresholdsGauge returns a boolean if a field has been set.
func (o *PlatformValues) HasNgwafWorkspaceThresholdsGauge() bool {
	if o != nil && o.NgwafWorkspaceThresholdsGauge != nil {
		return true
	}

	return false
}

// SetNgwafWorkspaceThresholdsGauge gets a reference to the given int64 and assigns it to the NgwafWorkspaceThresholdsGauge field.
func (o *PlatformValues) SetNgwafWorkspaceThresholdsGauge(v int64) {
	o.NgwafWorkspaceThresholdsGauge = &v
}

// GetObjectStorageBucketsCount returns the ObjectStorageBucketsCount field value if set, zero value otherwise.
func (o *PlatformValues) GetObjectStorageBucketsCount() int64 {
	if o == nil || o.ObjectStorageBucketsCount == nil {
		var ret int64
		return ret
	}
	return *o.ObjectStorageBucketsCount
}

// GetObjectStorageBucketsCountOk returns a tuple with the ObjectStorageBucketsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetObjectStorageBucketsCountOk() (*int64, bool) {
	if o == nil || o.ObjectStorageBucketsCount == nil {
		return nil, false
	}
	return o.ObjectStorageBucketsCount, true
}

// HasObjectStorageBucketsCount returns a boolean if a field has been set.
func (o *PlatformValues) HasObjectStorageBucketsCount() bool {
	if o != nil && o.ObjectStorageBucketsCount != nil {
		return true
	}

	return false
}

// SetObjectStorageBucketsCount gets a reference to the given int64 and assigns it to the ObjectStorageBucketsCount field.
func (o *PlatformValues) SetObjectStorageBucketsCount(v int64) {
	o.ObjectStorageBucketsCount = &v
}

// GetObjectStorageBytesCount returns the ObjectStorageBytesCount field value if set, zero value otherwise.
func (o *PlatformValues) GetObjectStorageBytesCount() int64 {
	if o == nil || o.ObjectStorageBytesCount == nil {
		var ret int64
		return ret
	}
	return *o.ObjectStorageBytesCount
}

// GetObjectStorageBytesCountOk returns a tuple with the ObjectStorageBytesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetObjectStorageBytesCountOk() (*int64, bool) {
	if o == nil || o.ObjectStorageBytesCount == nil {
		return nil, false
	}
	return o.ObjectStorageBytesCount, true
}

// HasObjectStorageBytesCount returns a boolean if a field has been set.
func (o *PlatformValues) HasObjectStorageBytesCount() bool {
	if o != nil && o.ObjectStorageBytesCount != nil {
		return true
	}

	return false
}

// SetObjectStorageBytesCount gets a reference to the given int64 and assigns it to the ObjectStorageBytesCount field.
func (o *PlatformValues) SetObjectStorageBytesCount(v int64) {
	o.ObjectStorageBytesCount = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o PlatformValues) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Timestamp != nil {
		toSerialize["timestamp"] = o.Timestamp
	}
	if o.Datacenter != nil {
		toSerialize["datacenter"] = o.Datacenter
	}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}
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
	if o.ImgoptoInputImageBytes != nil {
		toSerialize["imgopto_input_image_bytes"] = o.ImgoptoInputImageBytes
	}
	if o.ImgoptoOutputImageBytes != nil {
		toSerialize["imgopto_output_image_bytes"] = o.ImgoptoOutputImageBytes
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
	if o.Status530 != nil {
		toSerialize["status_530"] = o.Status530
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
	if o.ComputeServiceBereqError != nil {
		toSerialize["compute_service_bereq_error"] = o.ComputeServiceBereqError
	}
	if o.ComputeServiceMemoryExceededError != nil {
		toSerialize["compute_service_memory_exceeded_error"] = o.ComputeServiceMemoryExceededError
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
	if o.DdosProtectionRequestsDetectCount != nil {
		toSerialize["ddos_protection_requests_detect_count"] = o.DdosProtectionRequestsDetectCount
	}
	if o.DdosProtectionRequestsMitigateCount != nil {
		toSerialize["ddos_protection_requests_mitigate_count"] = o.DdosProtectionRequestsMitigateCount
	}
	if o.DdosProtectionRequestsAllowCount != nil {
		toSerialize["ddos_protection_requests_allow_count"] = o.DdosProtectionRequestsAllowCount
	}
	if o.ObjectStorageClassAOperationsCount != nil {
		toSerialize["object_storage_class_a_operations_count"] = o.ObjectStorageClassAOperationsCount
	}
	if o.ObjectStorageClassBOperationsCount != nil {
		toSerialize["object_storage_class_b_operations_count"] = o.ObjectStorageClassBOperationsCount
	}
	if o.AiaRequests != nil {
		toSerialize["aia_requests"] = o.AiaRequests
	}
	if o.AiaStatus1xx != nil {
		toSerialize["aia_status_1xx"] = o.AiaStatus1xx
	}
	if o.AiaStatus2xx != nil {
		toSerialize["aia_status_2xx"] = o.AiaStatus2xx
	}
	if o.AiaStatus3xx != nil {
		toSerialize["aia_status_3xx"] = o.AiaStatus3xx
	}
	if o.AiaStatus4xx != nil {
		toSerialize["aia_status_4xx"] = o.AiaStatus4xx
	}
	if o.AiaStatus5xx != nil {
		toSerialize["aia_status_5xx"] = o.AiaStatus5xx
	}
	if o.AiaResponseUsageTokens != nil {
		toSerialize["aia_response_usage_tokens"] = o.AiaResponseUsageTokens
	}
	if o.AiaOriginUsageTokens != nil {
		toSerialize["aia_origin_usage_tokens"] = o.AiaOriginUsageTokens
	}
	if o.AiaEstimatedTimeSavedMs != nil {
		toSerialize["aia_estimated_time_saved_ms"] = o.AiaEstimatedTimeSavedMs
	}
	if o.RequestCollapseUsableCount != nil {
		toSerialize["request_collapse_usable_count"] = o.RequestCollapseUsableCount
	}
	if o.RequestCollapseUnusableCount != nil {
		toSerialize["request_collapse_unusable_count"] = o.RequestCollapseUnusableCount
	}
	if o.ComputeCacheOperationsCount != nil {
		toSerialize["compute_cache_operations_count"] = o.ComputeCacheOperationsCount
	}
	if o.NgwafRequestsTotalCount != nil {
		toSerialize["ngwaf_requests_total_count"] = o.NgwafRequestsTotalCount
	}
	if o.NgwafRequestsUnknownCount != nil {
		toSerialize["ngwaf_requests_unknown_count"] = o.NgwafRequestsUnknownCount
	}
	if o.NgwafRequestsAllowedCount != nil {
		toSerialize["ngwaf_requests_allowed_count"] = o.NgwafRequestsAllowedCount
	}
	if o.NgwafRequestsLoggedCount != nil {
		toSerialize["ngwaf_requests_logged_count"] = o.NgwafRequestsLoggedCount
	}
	if o.NgwafRequestsBlockedCount != nil {
		toSerialize["ngwaf_requests_blocked_count"] = o.NgwafRequestsBlockedCount
	}
	if o.NgwafRequestsTimeoutCount != nil {
		toSerialize["ngwaf_requests_timeout_count"] = o.NgwafRequestsTimeoutCount
	}
	if o.NgwafRequestsChallengedCount != nil {
		toSerialize["ngwaf_requests_challenged_count"] = o.NgwafRequestsChallengedCount
	}
	if o.ApiDiscoveryRequestsCount != nil {
		toSerialize["api_discovery_requests_count"] = o.ApiDiscoveryRequestsCount
	}
	if o.ComputeRespStatus103 != nil {
		toSerialize["compute_resp_status_103"] = o.ComputeRespStatus103
	}
	if o.ComputeRespStatus200 != nil {
		toSerialize["compute_resp_status_200"] = o.ComputeRespStatus200
	}
	if o.ComputeRespStatus204 != nil {
		toSerialize["compute_resp_status_204"] = o.ComputeRespStatus204
	}
	if o.ComputeRespStatus206 != nil {
		toSerialize["compute_resp_status_206"] = o.ComputeRespStatus206
	}
	if o.ComputeRespStatus301 != nil {
		toSerialize["compute_resp_status_301"] = o.ComputeRespStatus301
	}
	if o.ComputeRespStatus302 != nil {
		toSerialize["compute_resp_status_302"] = o.ComputeRespStatus302
	}
	if o.ComputeRespStatus304 != nil {
		toSerialize["compute_resp_status_304"] = o.ComputeRespStatus304
	}
	if o.ComputeRespStatus400 != nil {
		toSerialize["compute_resp_status_400"] = o.ComputeRespStatus400
	}
	if o.ComputeRespStatus401 != nil {
		toSerialize["compute_resp_status_401"] = o.ComputeRespStatus401
	}
	if o.ComputeRespStatus403 != nil {
		toSerialize["compute_resp_status_403"] = o.ComputeRespStatus403
	}
	if o.ComputeRespStatus404 != nil {
		toSerialize["compute_resp_status_404"] = o.ComputeRespStatus404
	}
	if o.ComputeRespStatus416 != nil {
		toSerialize["compute_resp_status_416"] = o.ComputeRespStatus416
	}
	if o.ComputeRespStatus429 != nil {
		toSerialize["compute_resp_status_429"] = o.ComputeRespStatus429
	}
	if o.ComputeRespStatus500 != nil {
		toSerialize["compute_resp_status_500"] = o.ComputeRespStatus500
	}
	if o.ComputeRespStatus501 != nil {
		toSerialize["compute_resp_status_501"] = o.ComputeRespStatus501
	}
	if o.ComputeRespStatus502 != nil {
		toSerialize["compute_resp_status_502"] = o.ComputeRespStatus502
	}
	if o.ComputeRespStatus503 != nil {
		toSerialize["compute_resp_status_503"] = o.ComputeRespStatus503
	}
	if o.ComputeRespStatus504 != nil {
		toSerialize["compute_resp_status_504"] = o.ComputeRespStatus504
	}
	if o.ComputeRespStatus505 != nil {
		toSerialize["compute_resp_status_505"] = o.ComputeRespStatus505
	}
	if o.ComputeRespStatus530 != nil {
		toSerialize["compute_resp_status_530"] = o.ComputeRespStatus530
	}
	if o.ImgoptoComputeRequests != nil {
		toSerialize["imgopto_compute_requests"] = o.ImgoptoComputeRequests
	}
	if o.DNSBillableResponsesCount != nil {
		toSerialize["dns_billable_responses_count"] = o.DNSBillableResponsesCount
	}
	if o.DNSNonbillableResponsesCount != nil {
		toSerialize["dns_nonbillable_responses_count"] = o.DNSNonbillableResponsesCount
	}
	if o.Upgrade != nil {
		toSerialize["upgrade"] = o.Upgrade
	}
	if o.NgwafBotAnalysisRequestCount != nil {
		toSerialize["ngwaf_bot_analysis_request_count"] = o.NgwafBotAnalysisRequestCount
	}
	if o.ImgoptoAvifCount != nil {
		toSerialize["imgopto_avif_count"] = o.ImgoptoAvifCount
	}
	if o.ImgoptoJpegCount != nil {
		toSerialize["imgopto_jpeg_count"] = o.ImgoptoJpegCount
	}
	if o.ImgoptoPngCount != nil {
		toSerialize["imgopto_png_count"] = o.ImgoptoPngCount
	}
	if o.ImgoptoGifCount != nil {
		toSerialize["imgopto_gif_count"] = o.ImgoptoGifCount
	}
	if o.ImgoptoWebpCount != nil {
		toSerialize["imgopto_webp_count"] = o.ImgoptoWebpCount
	}
	if o.ImgoptoJpegxlCount != nil {
		toSerialize["imgopto_jpegxl_count"] = o.ImgoptoJpegxlCount
	}
	if o.ImgoptoSvgCount != nil {
		toSerialize["imgopto_svg_count"] = o.ImgoptoSvgCount
	}
	if o.ImgoptoMp4Count != nil {
		toSerialize["imgopto_mp4_count"] = o.ImgoptoMp4Count
	}
	if o.ComputeServiceResourceLimitsError != nil {
		toSerialize["compute_service_resource_limits_error"] = o.ComputeServiceResourceLimitsError
	}
	if o.ComputeServiceRuntimeError != nil {
		toSerialize["compute_service_runtime_error"] = o.ComputeServiceRuntimeError
	}
	if o.ComputeServiceChainError != nil {
		toSerialize["compute_service_chain_error"] = o.ComputeServiceChainError
	}
	if o.ComputePlatformInternalError != nil {
		toSerialize["compute_platform_internal_error"] = o.ComputePlatformInternalError
	}
	if o.ComputeServiceTimeoutError != nil {
		toSerialize["compute_service_timeout_error"] = o.ComputeServiceTimeoutError
	}
	if o.ComputeServiceVcpuExceededError != nil {
		toSerialize["compute_service_vcpu_exceeded_error"] = o.ComputeServiceVcpuExceededError
	}
	if o.ComputeServiceLimitsError != nil {
		toSerialize["compute_service_limits_error"] = o.ComputeServiceLimitsError
	}
	if o.ComputePlatformInvalidRequestError != nil {
		toSerialize["compute_platform_invalid_request_error"] = o.ComputePlatformInvalidRequestError
	}
	if o.ComputeSandboxes != nil {
		toSerialize["compute_sandboxes"] = o.ComputeSandboxes
	}
	if o.BotRequestsTotalCount != nil {
		toSerialize["bot_requests_total_count"] = o.BotRequestsTotalCount
	}
	if o.BotEdgeRequestsAnalyzedCount != nil {
		toSerialize["bot_edge_requests_analyzed_count"] = o.BotEdgeRequestsAnalyzedCount
	}
	if o.BotEdgeRequestsDetectedCount != nil {
		toSerialize["bot_edge_requests_detected_count"] = o.BotEdgeRequestsDetectedCount
	}
	if o.BotEdgeRequestsVerifiedCount != nil {
		toSerialize["bot_edge_requests_verified_count"] = o.BotEdgeRequestsVerifiedCount
	}
	if o.BotEdgeRequestsAiCrawlerCount != nil {
		toSerialize["bot_edge_requests_ai_crawler_count"] = o.BotEdgeRequestsAiCrawlerCount
	}
	if o.BotEdgeRequestsAiFetcherCount != nil {
		toSerialize["bot_edge_requests_ai_fetcher_count"] = o.BotEdgeRequestsAiFetcherCount
	}
	if o.BotEdgeRequestsAccessibilityCount != nil {
		toSerialize["bot_edge_requests_accessibility_count"] = o.BotEdgeRequestsAccessibilityCount
	}
	if o.BotEdgeRequestsContentFetcherCount != nil {
		toSerialize["bot_edge_requests_content_fetcher_count"] = o.BotEdgeRequestsContentFetcherCount
	}
	if o.BotEdgeRequestsMonitoringCount != nil {
		toSerialize["bot_edge_requests_monitoring_count"] = o.BotEdgeRequestsMonitoringCount
	}
	if o.BotEdgeRequestsOnlineMarketingCount != nil {
		toSerialize["bot_edge_requests_online_marketing_count"] = o.BotEdgeRequestsOnlineMarketingCount
	}
	if o.BotEdgeRequestsPagePreviewCount != nil {
		toSerialize["bot_edge_requests_page_preview_count"] = o.BotEdgeRequestsPagePreviewCount
	}
	if o.BotEdgeRequestsPlatformIntegrationsCount != nil {
		toSerialize["bot_edge_requests_platform_integrations_count"] = o.BotEdgeRequestsPlatformIntegrationsCount
	}
	if o.BotEdgeRequestsResearchCount != nil {
		toSerialize["bot_edge_requests_research_count"] = o.BotEdgeRequestsResearchCount
	}
	if o.BotEdgeRequestsSearchEngineCrawlerCount != nil {
		toSerialize["bot_edge_requests_search_engine_crawler_count"] = o.BotEdgeRequestsSearchEngineCrawlerCount
	}
	if o.BotEdgeRequestsSearchEngineOptimizationCount != nil {
		toSerialize["bot_edge_requests_search_engine_optimization_count"] = o.BotEdgeRequestsSearchEngineOptimizationCount
	}
	if o.BotEdgeRequestsSecurityToolsCount != nil {
		toSerialize["bot_edge_requests_security_tools_count"] = o.BotEdgeRequestsSecurityToolsCount
	}
	if o.ComputeHandoff != nil {
		toSerialize["compute_handoff"] = o.ComputeHandoff
	}
	if o.ComputeServiceBereqDNSError != nil {
		toSerialize["compute_service_bereq_dns_error"] = o.ComputeServiceBereqDNSError
	}
	if o.ComputeServiceBereqConnTimeoutError != nil {
		toSerialize["compute_service_bereq_conn_timeout_error"] = o.ComputeServiceBereqConnTimeoutError
	}
	if o.ComputeServiceBereqConnRefusedError != nil {
		toSerialize["compute_service_bereq_conn_refused_error"] = o.ComputeServiceBereqConnRefusedError
	}
	if o.ComputeServiceBereqConnOtherError != nil {
		toSerialize["compute_service_bereq_conn_other_error"] = o.ComputeServiceBereqConnOtherError
	}
	if o.ComputeServiceBereqTLSServerCertError != nil {
		toSerialize["compute_service_bereq_tls_server_cert_error"] = o.ComputeServiceBereqTLSServerCertError
	}
	if o.ComputeServiceBereqTLSOtherError != nil {
		toSerialize["compute_service_bereq_tls_other_error"] = o.ComputeServiceBereqTLSOtherError
	}
	if o.ComputeServiceBereqHTTPProtoV1Error != nil {
		toSerialize["compute_service_bereq_http_proto_v1_error"] = o.ComputeServiceBereqHTTPProtoV1Error
	}
	if o.ComputeServiceBereqHTTPProtoV2Error != nil {
		toSerialize["compute_service_bereq_http_proto_v2_error"] = o.ComputeServiceBereqHTTPProtoV2Error
	}
	if o.ComputeServiceBereqHTTPIncompleteError != nil {
		toSerialize["compute_service_bereq_http_incomplete_error"] = o.ComputeServiceBereqHTTPIncompleteError
	}
	if o.ComputeServiceBereqHTTPTimeoutError != nil {
		toSerialize["compute_service_bereq_http_timeout_error"] = o.ComputeServiceBereqHTTPTimeoutError
	}
	if o.ComputeServiceBereqHTTPOtherError != nil {
		toSerialize["compute_service_bereq_http_other_error"] = o.ComputeServiceBereqHTTPOtherError
	}
	if o.ComputeServiceBereqOtherError != nil {
		toSerialize["compute_service_bereq_other_error"] = o.ComputeServiceBereqOtherError
	}
	if o.ComputeServiceBereq5xxError != nil {
		toSerialize["compute_service_bereq_5xx_error"] = o.ComputeServiceBereq5xxError
	}
	if o.ComputeServiceBereqConnError != nil {
		toSerialize["compute_service_bereq_conn_error"] = o.ComputeServiceBereqConnError
	}
	if o.ComputeServiceBereqTLSError != nil {
		toSerialize["compute_service_bereq_tls_error"] = o.ComputeServiceBereqTLSError
	}
	if o.ComputeServiceBereqHTTPError != nil {
		toSerialize["compute_service_bereq_http_error"] = o.ComputeServiceBereqHTTPError
	}
	if o.BotChallengesPatsIssued != nil {
		toSerialize["bot_challenges_pats_issued"] = o.BotChallengesPatsIssued
	}
	if o.BotChallengesPatsSucceeded != nil {
		toSerialize["bot_challenges_pats_succeeded"] = o.BotChallengesPatsSucceeded
	}
	if o.BotEdgeRequestsHeadlessCount != nil {
		toSerialize["bot_edge_requests_headless_count"] = o.BotEdgeRequestsHeadlessCount
	}
	if o.Status499 != nil {
		toSerialize["status_499"] = o.Status499
	}
	if o.ArcRequests != nil {
		toSerialize["arc_requests"] = o.ArcRequests
	}
	if o.DdosChallengeCompleteTokensChecked != nil {
		toSerialize["ddos_challenge_complete_tokens_checked"] = o.DdosChallengeCompleteTokensChecked
	}
	if o.DdosChallengeCompleteTokensPassed != nil {
		toSerialize["ddos_challenge_complete_tokens_passed"] = o.DdosChallengeCompleteTokensPassed
	}
	if o.DdosChallengeCompleteTokensFailed != nil {
		toSerialize["ddos_challenge_complete_tokens_failed"] = o.DdosChallengeCompleteTokensFailed
	}
	if o.DdosChallengeStarts != nil {
		toSerialize["ddos_challenge_starts"] = o.DdosChallengeStarts
	}
	if o.DdosChallengeCompleteTokensIssued != nil {
		toSerialize["ddos_challenge_complete_tokens_issued"] = o.DdosChallengeCompleteTokensIssued
	}
	if o.DdosChallengesIssued != nil {
		toSerialize["ddos_challenges_issued"] = o.DdosChallengesIssued
	}
	if o.DdosChallengesSucceeded != nil {
		toSerialize["ddos_challenges_succeeded"] = o.DdosChallengesSucceeded
	}
	if o.DdosChallengesFailed != nil {
		toSerialize["ddos_challenges_failed"] = o.DdosChallengesFailed
	}
	if o.DdosChallengesPatsIssued != nil {
		toSerialize["ddos_challenges_pats_issued"] = o.DdosChallengesPatsIssued
	}
	if o.DdosChallengesPatsSucceeded != nil {
		toSerialize["ddos_challenges_pats_succeeded"] = o.DdosChallengesPatsSucceeded
	}
	if o.CspPagesGauge != nil {
		toSerialize["csp_pages_gauge"] = o.CspPagesGauge
	}
	if o.NgwafAccountRulesGauge != nil {
		toSerialize["ngwaf_account_rules_gauge"] = o.NgwafAccountRulesGauge
	}
	if o.NgwafAccountWorkspaces != nil {
		toSerialize["ngwaf_account_workspaces"] = o.NgwafAccountWorkspaces
	}
	if o.NgwafWorkspaceRulesGauge != nil {
		toSerialize["ngwaf_workspace_rules_gauge"] = o.NgwafWorkspaceRulesGauge
	}
	if o.NgwafWorkspaceThresholdsGauge != nil {
		toSerialize["ngwaf_workspace_thresholds_gauge"] = o.NgwafWorkspaceThresholdsGauge
	}
	if o.ObjectStorageBucketsCount != nil {
		toSerialize["object_storage_buckets_count"] = o.ObjectStorageBucketsCount
	}
	if o.ObjectStorageBytesCount != nil {
		toSerialize["object_storage_bytes_count"] = o.ObjectStorageBytesCount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *PlatformValues) UnmarshalJSON(bytes []byte) (err error) {
	varPlatformValues := _PlatformValues{}

	if err = json.Unmarshal(bytes, &varPlatformValues); err == nil {
		*o = PlatformValues(varPlatformValues)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "timestamp")
		delete(additionalProperties, "datacenter")
		delete(additionalProperties, "region")
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
		delete(additionalProperties, "imgopto_input_image_bytes")
		delete(additionalProperties, "imgopto_output_image_bytes")
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
		delete(additionalProperties, "status_530")
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
		delete(additionalProperties, "compute_service_bereq_error")
		delete(additionalProperties, "compute_service_memory_exceeded_error")
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
		delete(additionalProperties, "ddos_protection_requests_detect_count")
		delete(additionalProperties, "ddos_protection_requests_mitigate_count")
		delete(additionalProperties, "ddos_protection_requests_allow_count")
		delete(additionalProperties, "object_storage_class_a_operations_count")
		delete(additionalProperties, "object_storage_class_b_operations_count")
		delete(additionalProperties, "aia_requests")
		delete(additionalProperties, "aia_status_1xx")
		delete(additionalProperties, "aia_status_2xx")
		delete(additionalProperties, "aia_status_3xx")
		delete(additionalProperties, "aia_status_4xx")
		delete(additionalProperties, "aia_status_5xx")
		delete(additionalProperties, "aia_response_usage_tokens")
		delete(additionalProperties, "aia_origin_usage_tokens")
		delete(additionalProperties, "aia_estimated_time_saved_ms")
		delete(additionalProperties, "request_collapse_usable_count")
		delete(additionalProperties, "request_collapse_unusable_count")
		delete(additionalProperties, "compute_cache_operations_count")
		delete(additionalProperties, "ngwaf_requests_total_count")
		delete(additionalProperties, "ngwaf_requests_unknown_count")
		delete(additionalProperties, "ngwaf_requests_allowed_count")
		delete(additionalProperties, "ngwaf_requests_logged_count")
		delete(additionalProperties, "ngwaf_requests_blocked_count")
		delete(additionalProperties, "ngwaf_requests_timeout_count")
		delete(additionalProperties, "ngwaf_requests_challenged_count")
		delete(additionalProperties, "api_discovery_requests_count")
		delete(additionalProperties, "compute_resp_status_103")
		delete(additionalProperties, "compute_resp_status_200")
		delete(additionalProperties, "compute_resp_status_204")
		delete(additionalProperties, "compute_resp_status_206")
		delete(additionalProperties, "compute_resp_status_301")
		delete(additionalProperties, "compute_resp_status_302")
		delete(additionalProperties, "compute_resp_status_304")
		delete(additionalProperties, "compute_resp_status_400")
		delete(additionalProperties, "compute_resp_status_401")
		delete(additionalProperties, "compute_resp_status_403")
		delete(additionalProperties, "compute_resp_status_404")
		delete(additionalProperties, "compute_resp_status_416")
		delete(additionalProperties, "compute_resp_status_429")
		delete(additionalProperties, "compute_resp_status_500")
		delete(additionalProperties, "compute_resp_status_501")
		delete(additionalProperties, "compute_resp_status_502")
		delete(additionalProperties, "compute_resp_status_503")
		delete(additionalProperties, "compute_resp_status_504")
		delete(additionalProperties, "compute_resp_status_505")
		delete(additionalProperties, "compute_resp_status_530")
		delete(additionalProperties, "imgopto_compute_requests")
		delete(additionalProperties, "dns_billable_responses_count")
		delete(additionalProperties, "dns_nonbillable_responses_count")
		delete(additionalProperties, "upgrade")
		delete(additionalProperties, "ngwaf_bot_analysis_request_count")
		delete(additionalProperties, "imgopto_avif_count")
		delete(additionalProperties, "imgopto_jpeg_count")
		delete(additionalProperties, "imgopto_png_count")
		delete(additionalProperties, "imgopto_gif_count")
		delete(additionalProperties, "imgopto_webp_count")
		delete(additionalProperties, "imgopto_jpegxl_count")
		delete(additionalProperties, "imgopto_svg_count")
		delete(additionalProperties, "imgopto_mp4_count")
		delete(additionalProperties, "compute_service_resource_limits_error")
		delete(additionalProperties, "compute_service_runtime_error")
		delete(additionalProperties, "compute_service_chain_error")
		delete(additionalProperties, "compute_platform_internal_error")
		delete(additionalProperties, "compute_service_timeout_error")
		delete(additionalProperties, "compute_service_vcpu_exceeded_error")
		delete(additionalProperties, "compute_service_limits_error")
		delete(additionalProperties, "compute_platform_invalid_request_error")
		delete(additionalProperties, "compute_sandboxes")
		delete(additionalProperties, "bot_requests_total_count")
		delete(additionalProperties, "bot_edge_requests_analyzed_count")
		delete(additionalProperties, "bot_edge_requests_detected_count")
		delete(additionalProperties, "bot_edge_requests_verified_count")
		delete(additionalProperties, "bot_edge_requests_ai_crawler_count")
		delete(additionalProperties, "bot_edge_requests_ai_fetcher_count")
		delete(additionalProperties, "bot_edge_requests_accessibility_count")
		delete(additionalProperties, "bot_edge_requests_content_fetcher_count")
		delete(additionalProperties, "bot_edge_requests_monitoring_count")
		delete(additionalProperties, "bot_edge_requests_online_marketing_count")
		delete(additionalProperties, "bot_edge_requests_page_preview_count")
		delete(additionalProperties, "bot_edge_requests_platform_integrations_count")
		delete(additionalProperties, "bot_edge_requests_research_count")
		delete(additionalProperties, "bot_edge_requests_search_engine_crawler_count")
		delete(additionalProperties, "bot_edge_requests_search_engine_optimization_count")
		delete(additionalProperties, "bot_edge_requests_security_tools_count")
		delete(additionalProperties, "compute_handoff")
		delete(additionalProperties, "compute_service_bereq_dns_error")
		delete(additionalProperties, "compute_service_bereq_conn_timeout_error")
		delete(additionalProperties, "compute_service_bereq_conn_refused_error")
		delete(additionalProperties, "compute_service_bereq_conn_other_error")
		delete(additionalProperties, "compute_service_bereq_tls_server_cert_error")
		delete(additionalProperties, "compute_service_bereq_tls_other_error")
		delete(additionalProperties, "compute_service_bereq_http_proto_v1_error")
		delete(additionalProperties, "compute_service_bereq_http_proto_v2_error")
		delete(additionalProperties, "compute_service_bereq_http_incomplete_error")
		delete(additionalProperties, "compute_service_bereq_http_timeout_error")
		delete(additionalProperties, "compute_service_bereq_http_other_error")
		delete(additionalProperties, "compute_service_bereq_other_error")
		delete(additionalProperties, "compute_service_bereq_5xx_error")
		delete(additionalProperties, "compute_service_bereq_conn_error")
		delete(additionalProperties, "compute_service_bereq_tls_error")
		delete(additionalProperties, "compute_service_bereq_http_error")
		delete(additionalProperties, "bot_challenges_pats_issued")
		delete(additionalProperties, "bot_challenges_pats_succeeded")
		delete(additionalProperties, "bot_edge_requests_headless_count")
		delete(additionalProperties, "status_499")
		delete(additionalProperties, "arc_requests")
		delete(additionalProperties, "ddos_challenge_complete_tokens_checked")
		delete(additionalProperties, "ddos_challenge_complete_tokens_passed")
		delete(additionalProperties, "ddos_challenge_complete_tokens_failed")
		delete(additionalProperties, "ddos_challenge_starts")
		delete(additionalProperties, "ddos_challenge_complete_tokens_issued")
		delete(additionalProperties, "ddos_challenges_issued")
		delete(additionalProperties, "ddos_challenges_succeeded")
		delete(additionalProperties, "ddos_challenges_failed")
		delete(additionalProperties, "ddos_challenges_pats_issued")
		delete(additionalProperties, "ddos_challenges_pats_succeeded")
		delete(additionalProperties, "csp_pages_gauge")
		delete(additionalProperties, "ngwaf_account_rules_gauge")
		delete(additionalProperties, "ngwaf_account_workspaces")
		delete(additionalProperties, "ngwaf_workspace_rules_gauge")
		delete(additionalProperties, "ngwaf_workspace_thresholds_gauge")
		delete(additionalProperties, "object_storage_buckets_count")
		delete(additionalProperties, "object_storage_bytes_count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullablePlatformValues is a helper abstraction for handling nullable platformvalues types.
type NullablePlatformValues struct {
	value *PlatformValues
	isSet bool
}

// Get returns the value.
func (v NullablePlatformValues) Get() *PlatformValues {
	return v.value
}

// Set modifies the value.
func (v *NullablePlatformValues) Set(val *PlatformValues) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullablePlatformValues) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullablePlatformValues) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullablePlatformValues returns a pointer to a new instance of NullablePlatformValues.
func NewNullablePlatformValues(val *PlatformValues) *NullablePlatformValues {
	return &NullablePlatformValues{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullablePlatformValues) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullablePlatformValues) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
