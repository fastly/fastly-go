# RealtimeMeasurements

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Requests** | Pointer to **int32** | Number of requests processed. | [optional] 
**Logging** | Pointer to **int32** | Number of log lines sent (alias for `log`). | [optional] 
**Log** | Pointer to **int32** | Number of log lines sent. | [optional] 
**RespHeaderBytes** | Pointer to **int32** | Total header bytes delivered (edge_resp_header_bytes + shield_resp_header_bytes). | [optional] 
**HeaderSize** | Pointer to **int32** | Total header bytes delivered (alias for resp_header_bytes). | [optional] 
**RespBodyBytes** | Pointer to **int32** | Total body bytes delivered (edge_resp_body_bytes + shield_resp_body_bytes). | [optional] 
**BodySize** | Pointer to **int32** | Total body bytes delivered (alias for resp_body_bytes). | [optional] 
**Hits** | Pointer to **int32** | Number of cache hits. | [optional] 
**Miss** | Pointer to **int32** | Number of cache misses. | [optional] 
**Pass** | Pointer to **int32** | Number of requests that passed through the CDN without being cached. | [optional] 
**Synth** | Pointer to **int32** | Number of requests that returned a synthetic response (i.e., response objects created with the `synthetic` VCL statement). | [optional] 
**Errors** | Pointer to **int32** | Number of cache errors. | [optional] 
**HitsTime** | Pointer to **float32** | Total amount of time spent processing cache hits (in seconds). | [optional] 
**MissTime** | Pointer to **float32** | Total amount of time spent processing cache misses (in seconds). | [optional] 
**MissHistogram** | Pointer to **map[string]map[string]any** | A histogram. The value in each bucket is the number of requests to the origin whose responses arrived during the time period represented by the bucket. The key of each bucket represents the upper bound (in response time) of that bucket. The buckets vary in width and cover the time periods 0-10ms (in 1ms increments), 10-250ms (in 10ms increments), 250-1,000ms (in 50ms increments), 1,000-3,000ms (in 100ms increments), 3,000-10,000ms (in 500 ms increments), 10,000-20,000ms (in 1,000ms increments), 20,000-60,000ms (in 5,000ms increments), and 60,000ms through infinity (in a single bucket). | [optional] 
**ComputeRequests** | Pointer to **int32** | The total number of requests that were received for your service by Fastly. | [optional] 
**ComputeExecutionTimeMs** | Pointer to **float32** | The amount of active CPU time used to process your requests (in milliseconds). | [optional] 
**ComputeRAMUsed** | Pointer to **int32** | The amount of RAM used for your service by Fastly (in bytes). | [optional] 
**ComputeRequestTimeMs** | Pointer to **float32** | The total, actual amount of time used to process your requests, including active CPU time (in milliseconds). | [optional] 
**ComputeRequestTimeBilledMs** | Pointer to **float32** | The total amount of request processing time you will be billed for, measured in 50 millisecond increments. | [optional] 
**Shield** | Pointer to **int32** | Number of requests from edge to the shield POP. | [optional] 
**Ipv6** | Pointer to **int32** | Number of requests that were received over IPv6. | [optional] 
**Imgopto** | Pointer to **int32** | Number of responses that came from the Fastly Image Optimizer service. If the service receives 10 requests for an image, this stat will be 10 regardless of how many times the image was transformed. | [optional] 
**ImgoptoShield** | Pointer to **int32** | Number of responses that came from the Fastly Image Optimizer service via a shield. | [optional] 
**ImgoptoTransforms** | Pointer to **int32** | Number of transforms performed by the Fastly Image Optimizer service. | [optional] 
**Otfp** | Pointer to **int32** | Number of responses that came from the Fastly On-the-Fly Packaging service for video-on-demand. | [optional] 
**OtfpShield** | Pointer to **int32** | Number of responses that came from the Fastly On-the-Fly Packaging service for video-on-demand via a shield. | [optional] 
**OtfpManifests** | Pointer to **int32** | Number of responses that were manifest files from the Fastly On-the-Fly Packaging service for video-on-demand. | [optional] 
**Video** | Pointer to **int32** | Number of responses with the video segment or video manifest MIME type (i.e., application/x-mpegurl, application/vnd.apple.mpegurl, application/f4m, application/dash+xml, application/vnd.ms-sstr+xml, ideo/mp2t, audio/aac, video/f4f, video/x-flv, video/mp4, audio/mp4). | [optional] 
**Pci** | Pointer to **int32** | Number of responses with the PCI flag turned on. | [optional] 
**HTTP2** | Pointer to **int32** | Number of requests received over HTTP/2. | [optional] 
**HTTP3** | Pointer to **int32** | Number of requests received over HTTP/3. | [optional] 
**Restarts** | Pointer to **int32** | Number of restarts performed. | [optional] 
**ReqHeaderBytes** | Pointer to **int32** | Total header bytes received. | [optional] 
**ReqBodyBytes** | Pointer to **int32** | Total body bytes received. | [optional] 
**BereqHeaderBytes** | Pointer to **int32** | Total header bytes sent to origin. | [optional] 
**BereqBodyBytes** | Pointer to **int32** | Total body bytes sent to origin. | [optional] 
**WafBlocked** | Pointer to **int32** | Number of requests that triggered a WAF rule and were blocked. | [optional] 
**WafLogged** | Pointer to **int32** | Number of requests that triggered a WAF rule and were logged. | [optional] 
**WafPassed** | Pointer to **int32** | Number of requests that triggered a WAF rule and were passed. | [optional] 
**AttackReqHeaderBytes** | Pointer to **int32** | Total header bytes received from requests that triggered a WAF rule. | [optional] 
**AttackReqBodyBytes** | Pointer to **int32** | Total body bytes received from requests that triggered a WAF rule. | [optional] 
**AttackRespSynthBytes** | Pointer to **int32** | Total bytes delivered for requests that triggered a WAF rule and returned a synthetic response. | [optional] 
**AttackLoggedReqHeaderBytes** | Pointer to **int32** | Total header bytes received from requests that triggered a WAF rule that was logged. | [optional] 
**AttackLoggedReqBodyBytes** | Pointer to **int32** | Total body bytes received from requests that triggered a WAF rule that was logged. | [optional] 
**AttackBlockedReqHeaderBytes** | Pointer to **int32** | Total header bytes received from requests that triggered a WAF rule that was blocked. | [optional] 
**AttackBlockedReqBodyBytes** | Pointer to **int32** | Total body bytes received from requests that triggered a WAF rule that was blocked. | [optional] 
**AttackPassedReqHeaderBytes** | Pointer to **int32** | Total header bytes received from requests that triggered a WAF rule that was passed. | [optional] 
**AttackPassedReqBodyBytes** | Pointer to **int32** | Total body bytes received from requests that triggered a WAF rule that was passed. | [optional] 
**ShieldRespHeaderBytes** | Pointer to **int32** | Total header bytes delivered via a shield. | [optional] 
**ShieldRespBodyBytes** | Pointer to **int32** | Total body bytes delivered via a shield. | [optional] 
**OtfpRespHeaderBytes** | Pointer to **int32** | Total header bytes delivered from the Fastly On-the-Fly Packaging service for video-on-demand. | [optional] 
**OtfpRespBodyBytes** | Pointer to **int32** | Total body bytes delivered from the Fastly On-the-Fly Packaging service for video-on-demand. | [optional] 
**OtfpShieldRespHeaderBytes** | Pointer to **int32** | Total header bytes delivered via a shield for the Fastly On-the-Fly Packaging service for video-on-demand. | [optional] 
**OtfpShieldRespBodyBytes** | Pointer to **int32** | Total body bytes delivered via a shield for the Fastly On-the-Fly Packaging service for video-on-demand. | [optional] 
**OtfpShieldTime** | Pointer to **float32** | Total amount of time spent delivering a response via a shield from the Fastly On-the-Fly Packaging service for video-on-demand (in seconds). | [optional] 
**OtfpDeliverTime** | Pointer to **float32** | Total amount of time spent delivering a response from the Fastly On-the-Fly Packaging service for video-on-demand (in seconds). | [optional] 
**ImgoptoRespHeaderBytes** | Pointer to **int32** | Total header bytes delivered from the Fastly Image Optimizer service, including shield traffic. | [optional] 
**ImgoptoRespBodyBytes** | Pointer to **int32** | Total body bytes delivered from the Fastly Image Optimizer service, including shield traffic. | [optional] 
**ImgoptoShieldRespHeaderBytes** | Pointer to **int32** | Total header bytes delivered via a shield from the Fastly Image Optimizer service. | [optional] 
**ImgoptoShieldRespBodyBytes** | Pointer to **int32** | Total body bytes delivered via a shield from the Fastly Image Optimizer service. | [optional] 
**Status1xx** | Pointer to **int32** | Number of \&quot;Informational\&quot; category status codes delivered. | [optional] 
**Status2xx** | Pointer to **int32** | Number of \&quot;Success\&quot; status codes delivered. | [optional] 
**Status3xx** | Pointer to **int32** | Number of \&quot;Redirection\&quot; codes delivered. | [optional] 
**Status4xx** | Pointer to **int32** | Number of \&quot;Client Error\&quot; codes delivered. | [optional] 
**Status5xx** | Pointer to **int32** | Number of \&quot;Server Error\&quot; codes delivered. | [optional] 
**Status200** | Pointer to **int32** | Number of responses sent with status code 200 (Success). | [optional] 
**Status204** | Pointer to **int32** | Number of responses sent with status code 204 (No Content). | [optional] 
**Status206** | Pointer to **int32** | Number of responses sent with status code 206 (Partial Content). | [optional] 
**Status301** | Pointer to **int32** | Number of responses sent with status code 301 (Moved Permanently). | [optional] 
**Status302** | Pointer to **int32** | Number of responses sent with status code 302 (Found). | [optional] 
**Status304** | Pointer to **int32** | Number of responses sent with status code 304 (Not Modified). | [optional] 
**Status400** | Pointer to **int32** | Number of responses sent with status code 400 (Bad Request). | [optional] 
**Status401** | Pointer to **int32** | Number of responses sent with status code 401 (Unauthorized). | [optional] 
**Status403** | Pointer to **int32** | Number of responses sent with status code 403 (Forbidden). | [optional] 
**Status404** | Pointer to **int32** | Number of responses sent with status code 404 (Not Found). | [optional] 
**Status406** | Pointer to **int32** | Number of responses sent with status code 406 (Not Acceptable). | [optional] 
**Status416** | Pointer to **int32** | Number of responses sent with status code 416 (Range Not Satisfiable). | [optional] 
**Status429** | Pointer to **int32** | Number of responses sent with status code 429 (Too Many Requests). | [optional] 
**Status500** | Pointer to **int32** | Number of responses sent with status code 500 (Internal Server Error). | [optional] 
**Status501** | Pointer to **int32** | Number of responses sent with status code 501 (Not Implemented). | [optional] 
**Status502** | Pointer to **int32** | Number of responses sent with status code 502 (Bad Gateway). | [optional] 
**Status503** | Pointer to **int32** | Number of responses sent with status code 503 (Service Unavailable). | [optional] 
**Status504** | Pointer to **int32** | Number of responses sent with status code 504 (Gateway Timeout). | [optional] 
**Status505** | Pointer to **int32** | Number of responses sent with status code 505 (HTTP Version Not Supported). | [optional] 
**Uncacheable** | Pointer to **int32** | Number of requests that were designated uncachable. | [optional] 
**PassTime** | Pointer to **float32** | Total amount of time spent processing cache passes (in seconds). | [optional] 
**TLS** | Pointer to **int32** | Number of requests that were received over TLS. | [optional] 
**TLSV10** | Pointer to **int32** | Number of requests received over TLS 1.0. | [optional] 
**TLSV11** | Pointer to **int32** | Number of requests received over TLS 1.1. | [optional] 
**TLSV12** | Pointer to **int32** | Number of requests received over TLS 1.2. | [optional] 
**TLSV13** | Pointer to **int32** | Number of requests received over TLS 1.3. | [optional] 
**ObjectSize1k** | Pointer to **int32** | Number of objects served that were under 1KB in size. | [optional] 
**ObjectSize10k** | Pointer to **int32** | Number of objects served that were between 1KB and 10KB in size. | [optional] 
**ObjectSize100k** | Pointer to **int32** | Number of objects served that were between 10KB and 100KB in size. | [optional] 
**ObjectSize1m** | Pointer to **int32** | Number of objects served that were between 100KB and 1MB in size. | [optional] 
**ObjectSize10m** | Pointer to **int32** | Number of objects served that were between 1MB and 10MB in size. | [optional] 
**ObjectSize100m** | Pointer to **int32** | Number of objects served that were between 10MB and 100MB in size. | [optional] 
**ObjectSize1g** | Pointer to **int32** | Number of objects served that were between 100MB and 1GB in size. | [optional] 
**ObjectSizeOther** | Pointer to **int32** | Number of objects served that were larger than 1GB in size. | [optional] 
**RecvSubTime** | Pointer to **float32** | Time spent inside the `vcl_recv` Varnish subroutine (in nanoseconds). | [optional] 
**RecvSubCount** | Pointer to **int32** | Number of executions of the `vcl_recv` Varnish subroutine. | [optional] 
**HashSubTime** | Pointer to **float32** | Time spent inside the `vcl_hash` Varnish subroutine (in nanoseconds). | [optional] 
**HashSubCount** | Pointer to **int32** | Number of executions of the `vcl_hash` Varnish subroutine. | [optional] 
**MissSubTime** | Pointer to **float32** | Time spent inside the `vcl_miss` Varnish subroutine (in nanoseconds). | [optional] 
**MissSubCount** | Pointer to **int32** | Number of executions of the `vcl_miss` Varnish subroutine. | [optional] 
**FetchSubTime** | Pointer to **float32** | Time spent inside the `vcl_fetch` Varnish subroutine (in nanoseconds). | [optional] 
**FetchSubCount** | Pointer to **int32** | Number of executions of the `vcl_fetch` Varnish subroutine. | [optional] 
**PassSubTime** | Pointer to **float32** | Time spent inside the `vcl_pass` Varnish subroutine (in nanoseconds). | [optional] 
**PassSubCount** | Pointer to **int32** | Number of executions of the `vcl_pass` Varnish subroutine. | [optional] 
**PipeSubTime** | Pointer to **float32** | Time spent inside the `vcl_pipe` Varnish subroutine (in nanoseconds). | [optional] 
**PipeSubCount** | Pointer to **int32** | Number of executions of the `vcl_pipe` Varnish subroutine. | [optional] 
**DeliverSubTime** | Pointer to **float32** | Time spent inside the `vcl_deliver` Varnish subroutine (in nanoseconds). | [optional] 
**DeliverSubCount** | Pointer to **int32** | Number of executions of the `vcl_deliver` Varnish subroutine. | [optional] 
**ErrorSubTime** | Pointer to **float32** | Time spent inside the `vcl_error` Varnish subroutine (in nanoseconds). | [optional] 
**ErrorSubCount** | Pointer to **int32** | Number of executions of the `vcl_error` Varnish subroutine. | [optional] 
**HitSubTime** | Pointer to **float32** | Time spent inside the `vcl_hit` Varnish subroutine (in nanoseconds). | [optional] 
**HitSubCount** | Pointer to **int32** | Number of executions of the `vcl_hit` Varnish subroutine. | [optional] 
**PrehashSubTime** | Pointer to **float32** | Time spent inside the `vcl_prehash` Varnish subroutine (in nanoseconds). | [optional] 
**PrehashSubCount** | Pointer to **int32** | Number of executions of the `vcl_prehash` Varnish subroutine. | [optional] 
**PredeliverSubTime** | Pointer to **float32** | Time spent inside the `vcl_predeliver` Varnish subroutine (in nanoseconds). | [optional] 
**PredeliverSubCount** | Pointer to **int32** | Number of executions of the `vcl_predeliver` Varnish subroutine. | [optional] 
**HitRespBodyBytes** | Pointer to **int32** | Total body bytes delivered for cache hits. | [optional] 
**MissRespBodyBytes** | Pointer to **int32** | Total body bytes delivered for cache misses. | [optional] 
**PassRespBodyBytes** | Pointer to **int32** | Total body bytes delivered for cache passes. | [optional] 
**ComputeReqHeaderBytes** | Pointer to **int32** | Total header bytes received by the Compute platform. | [optional] 
**ComputeReqBodyBytes** | Pointer to **int32** | Total body bytes received by the Compute platform. | [optional] 
**ComputeRespHeaderBytes** | Pointer to **int32** | Total header bytes sent from Compute to end user. | [optional] 
**ComputeRespBodyBytes** | Pointer to **int32** | Total body bytes sent from Compute to end user. | [optional] 
**Imgvideo** | Pointer to **int32** | Number of video responses that came from the Fastly Image Optimizer service. | [optional] 
**ImgvideoFrames** | Pointer to **int32** | Number of video frames that came from the Fastly Image Optimizer service. A video frame is an individual image within a sequence of video. | [optional] 
**ImgvideoRespHeaderBytes** | Pointer to **int32** | Total header bytes of video delivered from the Fastly Image Optimizer service. | [optional] 
**ImgvideoRespBodyBytes** | Pointer to **int32** | Total body bytes of video delivered from the Fastly Image Optimizer service. | [optional] 
**ImgvideoShield** | Pointer to **int32** | Number of video responses delivered via a shield that came from the Fastly Image Optimizer service. | [optional] 
**ImgvideoShieldFrames** | Pointer to **int32** | Number of video frames delivered via a shield that came from the Fastly Image Optimizer service. A video frame is an individual image within a sequence of video. | [optional] 
**ImgvideoShieldRespHeaderBytes** | Pointer to **int32** | Total header bytes of video delivered via a shield from the Fastly Image Optimizer service. | [optional] 
**ImgvideoShieldRespBodyBytes** | Pointer to **int32** | Total body bytes of video delivered via a shield from the Fastly Image Optimizer service. | [optional] 
**LogBytes** | Pointer to **int32** | Total log bytes sent. | [optional] 
**EdgeRequests** | Pointer to **int32** | Number of requests sent by end users to Fastly. | [optional] 
**EdgeRespHeaderBytes** | Pointer to **int32** | Total header bytes delivered from Fastly to the end user. | [optional] 
**EdgeRespBodyBytes** | Pointer to **int32** | Total body bytes delivered from Fastly to the end user. | [optional] 
**OriginRevalidations** | Pointer to **int32** | Number of responses received from origin with a `304` status code in response to an `If-Modified-Since` or `If-None-Match` request. Under regular scenarios, a revalidation will imply a cache hit. However, if using Fastly Image Optimizer or segmented caching this may result in a cache miss. | [optional] 
**OriginFetches** | Pointer to **int32** | Number of requests sent to origin. | [optional] 
**OriginFetchHeaderBytes** | Pointer to **int32** | Total request header bytes sent to origin. | [optional] 
**OriginFetchBodyBytes** | Pointer to **int32** | Total request body bytes sent to origin. | [optional] 
**OriginFetchRespHeaderBytes** | Pointer to **int32** | Total header bytes received from origin. | [optional] 
**OriginFetchRespBodyBytes** | Pointer to **int32** | Total body bytes received from origin. | [optional] 
**ShieldRevalidations** | Pointer to **int32** | Number of responses received from origin with a `304` status code, in response to an `If-Modified-Since` or `If-None-Match` request to a shield. Under regular scenarios, a revalidation will imply a cache hit. However, if using segmented caching this may result in a cache miss. | [optional] 
**ShieldFetches** | Pointer to **int32** | Number of requests made from one Fastly POP to another, as part of shielding. | [optional] 
**ShieldFetchHeaderBytes** | Pointer to **int32** | Total request header bytes sent to a shield. | [optional] 
**ShieldFetchBodyBytes** | Pointer to **int32** | Total request body bytes sent to a shield. | [optional] 
**ShieldFetchRespHeaderBytes** | Pointer to **int32** | Total response header bytes sent from a shield to the edge. | [optional] 
**ShieldFetchRespBodyBytes** | Pointer to **int32** | Total response body bytes sent from a shield to the edge. | [optional] 
**SegblockOriginFetches** | Pointer to **int32** | Number of `Range` requests to origin for segments of resources when using segmented caching. | [optional] 
**SegblockShieldFetches** | Pointer to **int32** | Number of `Range` requests to a shield for segments of resources when using segmented caching. | [optional] 
**ComputeRespStatus1xx** | Pointer to **int32** | Number of \&quot;Informational\&quot; category status codes delivered by the Compute platform. | [optional] 
**ComputeRespStatus2xx** | Pointer to **int32** | Number of \&quot;Success\&quot; category status codes delivered by the Compute platform. | [optional] 
**ComputeRespStatus3xx** | Pointer to **int32** | Number of \&quot;Redirection\&quot; category status codes delivered by the Compute platform. | [optional] 
**ComputeRespStatus4xx** | Pointer to **int32** | Number of \&quot;Client Error\&quot; category status codes delivered by the Compute platform. | [optional] 
**ComputeRespStatus5xx** | Pointer to **int32** | Number of \&quot;Server Error\&quot; category status codes delivered by the Compute platform. | [optional] 
**EdgeHitRequests** | Pointer to **int32** | Number of requests sent by end users to Fastly that resulted in a hit at the edge. | [optional] 
**EdgeMissRequests** | Pointer to **int32** | Number of requests sent by end users to Fastly that resulted in a miss at the edge. | [optional] 
**ComputeBereqHeaderBytes** | Pointer to **int32** | Total header bytes sent to backends (origins) by the Compute platform. | [optional] 
**ComputeBereqBodyBytes** | Pointer to **int32** | Total body bytes sent to backends (origins) by the Compute platform. | [optional] 
**ComputeBerespHeaderBytes** | Pointer to **int32** | Total header bytes received from backends (origins) by the Compute platform. | [optional] 
**ComputeBerespBodyBytes** | Pointer to **int32** | Total body bytes received from backends (origins) by the Compute platform. | [optional] 
**OriginCacheFetches** | Pointer to **int32** | The total number of completed requests made to backends (origins) that returned cacheable content. | [optional] 
**ShieldCacheFetches** | Pointer to **int32** | The total number of completed requests made to shields that returned cacheable content. | [optional] 
**ComputeBereqs** | Pointer to **int32** | Number of backend requests started. | [optional] 
**ComputeBereqErrors** | Pointer to **int32** | Number of backend request errors, including timeouts. | [optional] 
**ComputeResourceLimitExceeded** | Pointer to **int32** | Number of times a guest exceeded its resource limit, includes heap, stack, globals, and code execution timeout. | [optional] 
**ComputeHeapLimitExceeded** | Pointer to **int32** | Number of times a guest exceeded its heap limit. | [optional] 
**ComputeStackLimitExceeded** | Pointer to **int32** | Number of times a guest exceeded its stack limit. | [optional] 
**ComputeGlobalsLimitExceeded** | Pointer to **int32** | Number of times a guest exceeded its globals limit. | [optional] 
**ComputeGuestErrors** | Pointer to **int32** | Number of times a service experienced a guest code error. | [optional] 
**ComputeRuntimeErrors** | Pointer to **int32** | Number of times a service experienced a guest runtime error. | [optional] 
**EdgeHitRespBodyBytes** | Pointer to **int32** | Body bytes delivered for edge hits. | [optional] 
**EdgeHitRespHeaderBytes** | Pointer to **int32** | Header bytes delivered for edge hits. | [optional] 
**EdgeMissRespBodyBytes** | Pointer to **int32** | Body bytes delivered for edge misses. | [optional] 
**EdgeMissRespHeaderBytes** | Pointer to **int32** | Header bytes delivered for edge misses. | [optional] 
**OriginCacheFetchRespBodyBytes** | Pointer to **int32** | Body bytes received from origin for cacheable content. | [optional] 
**OriginCacheFetchRespHeaderBytes** | Pointer to **int32** | Header bytes received from an origin for cacheable content. | [optional] 
**ShieldHitRequests** | Pointer to **int32** | Number of requests that resulted in a hit at a shield. | [optional] 
**ShieldMissRequests** | Pointer to **int32** | Number of requests that resulted in a miss at a shield. | [optional] 
**ShieldHitRespHeaderBytes** | Pointer to **int32** | Header bytes delivered for shield hits. | [optional] 
**ShieldHitRespBodyBytes** | Pointer to **int32** | Body bytes delivered for shield hits. | [optional] 
**ShieldMissRespHeaderBytes** | Pointer to **int32** | Header bytes delivered for shield misses. | [optional] 
**ShieldMissRespBodyBytes** | Pointer to **int32** | Body bytes delivered for shield misses. | [optional] 
**WebsocketReqHeaderBytes** | Pointer to **int32** | Total header bytes received from end users over passthrough WebSocket connections. | [optional] 
**WebsocketReqBodyBytes** | Pointer to **int32** | Total message content bytes received from end users over passthrough WebSocket connections. | [optional] 
**WebsocketRespHeaderBytes** | Pointer to **int32** | Total header bytes sent to end users over passthrough WebSocket connections. | [optional] 
**WebsocketBereqHeaderBytes** | Pointer to **int32** | Total header bytes sent to backends over passthrough WebSocket connections. | [optional] 
**WebsocketBereqBodyBytes** | Pointer to **int32** | Total message content bytes sent to backends over passthrough WebSocket connections. | [optional] 
**WebsocketBerespHeaderBytes** | Pointer to **int32** | Total header bytes received from backends over passthrough WebSocket connections. | [optional] 
**WebsocketBerespBodyBytes** | Pointer to **int32** | Total message content bytes received from backends over passthrough WebSocket connections. | [optional] 
**WebsocketConnTimeMs** | Pointer to **int32** | Total duration of passthrough WebSocket connections with end users. | [optional] 
**WebsocketRespBodyBytes** | Pointer to **int32** | Total message content bytes sent to end users over passthrough WebSocket connections. | [optional] 
**FanoutRecvPublishes** | Pointer to **int32** | Total published messages received from the publish API endpoint. | [optional] 
**FanoutSendPublishes** | Pointer to **int32** | Total published messages sent to end users. | [optional] 
**KvStoreClassAOperations** | Pointer to **int32** | The total number of class a operations for the KV store. | [optional] 
**KvStoreClassBOperations** | Pointer to **int32** | The total number of class b operations for the KV store. | [optional] 
**ObjectStoreClassAOperations** | Pointer to **int32** | Use kv_store_class_a_operations. | [optional] 
**ObjectStoreClassBOperations** | Pointer to **int32** | Use kv_store_class_b_operations. | [optional] 
**FanoutReqHeaderBytes** | Pointer to **int32** | Total header bytes received from end users over Fanout connections. | [optional] 
**FanoutReqBodyBytes** | Pointer to **int32** | Total body or message content bytes received from end users over Fanout connections. | [optional] 
**FanoutRespHeaderBytes** | Pointer to **int32** | Total header bytes sent to end users over Fanout connections. | [optional] 
**FanoutRespBodyBytes** | Pointer to **int32** | Total body or message content bytes sent to end users over Fanout connections, excluding published message content. | [optional] 
**FanoutBereqHeaderBytes** | Pointer to **int32** | Total header bytes sent to backends over Fanout connections. | [optional] 
**FanoutBereqBodyBytes** | Pointer to **int32** | Total body or message content bytes sent to backends over Fanout connections. | [optional] 
**FanoutBerespHeaderBytes** | Pointer to **int32** | Total header bytes received from backends over Fanout connections. | [optional] 
**FanoutBerespBodyBytes** | Pointer to **int32** | Total body or message content bytes received from backends over Fanout connections. | [optional] 
**FanoutConnTimeMs** | Pointer to **int32** | Total duration of Fanout connections with end users. | [optional] 
**DdosActionLimitStreamsConnections** | Pointer to **int32** | For HTTP/2, the number of connections the limit-streams action was applied to. The limit-streams action caps the allowed number of concurrent streams in a connection. | [optional] 
**DdosActionLimitStreamsRequests** | Pointer to **int32** | For HTTP/2, the number of requests made on a connection for which the limit-streams action was taken. The limit-streams action caps the allowed number of concurrent streams in a connection. | [optional] 
**DdosActionTarpitAccept** | Pointer to **int32** | The number of times the tarpit-accept action was taken. The tarpit-accept action adds a delay when accepting future connections. | [optional] 
**DdosActionTarpit** | Pointer to **int32** | The number of times the tarpit action was taken. The tarpit action delays writing the response to the client. | [optional] 
**DdosActionClose** | Pointer to **int32** | The number of times the close action was taken. The close action aborts the connection as soon as possible. The close action takes effect either right after accept, right after the client hello, or right after the response was sent. | [optional] 
**DdosActionBlackhole** | Pointer to **int32** | The number of times the blackhole action was taken. The blackhole action quietly closes a TCP connection without sending a reset. The blackhole action quietly closes a TCP connection without notifying its peer (all TCP state is dropped). | [optional] 
**BotChallengeStarts** | Pointer to **int32** | The number of challenge-start tokens created. | [optional] 
**BotChallengeCompleteTokensPassed** | Pointer to **int32** | The number of challenge-complete tokens that passed validation. | [optional] 
**BotChallengeCompleteTokensFailed** | Pointer to **int32** | The number of challenge-complete tokens that failed validation. | [optional] 
**BotChallengeCompleteTokensChecked** | Pointer to **int32** | The number of challenge-complete tokens checked. | [optional] 
**BotChallengeCompleteTokensDisabled** | Pointer to **int32** | The number of challenge-complete tokens not checked because the feature was disabled. | [optional] 
**BotChallengesIssued** | Pointer to **int32** | The number of challenges issued. For example, the issuance of a CAPTCHA challenge. | [optional] 
**BotChallengesSucceeded** | Pointer to **int32** | The number of successful challenge solutions processed. For example, a correct CAPTCHA solution. | [optional] 
**BotChallengesFailed** | Pointer to **int32** | The number of failed challenge solutions processed. For example, an incorrect CAPTCHA solution. | [optional] 
**BotChallengeCompleteTokensIssued** | Pointer to **int32** | The number of challenge-complete tokens issued. For example, issuing a challenge-complete token after a series of CAPTCHA challenges ending in success. | [optional] 
**DdosActionDowngrade** | Pointer to **int32** | The number of times the downgrade action was taken. The downgrade action restricts the client to http1. | [optional] 
**DdosActionDowngradedConnections** | Pointer to **int32** | The number of connections the downgrade action was applied to. The downgrade action restricts the connection to http1. | [optional] 
**VclOnComputeHitRequests** | Pointer to **int32** | Number of cache hits for a VCL service running on Compute. | [optional] 
**VclOnComputeMissRequests** | Pointer to **int32** | Number of cache misses for a VCL service running on Compute. | [optional] 
**VclOnComputePassRequests** | Pointer to **int32** | Number of requests that passed through the CDN without being cached for a VCL service running on Compute. | [optional] 
**VclOnComputeErrorRequests** | Pointer to **int32** | Number of cache errors for a VCL service running on Compute. | [optional] 
**VclOnComputeSynthRequests** | Pointer to **int32** | Number of requests that returned a synthetic response (i.e., response objects created with the `synthetic` VCL statement) for a VCL service running on Compute. | [optional] 
**VclOnComputeEdgeHitRequests** | Pointer to **int32** | Number of requests sent by end users to Fastly that resulted in a hit at the edge for a VCL service running on Compute. | [optional] 
**VclOnComputeEdgeMissRequests** | Pointer to **int32** | Number of requests sent by end users to Fastly that resulted in a miss at the edge for a VCL service running on Compute. | [optional] 
**AllHitRequests** | Pointer to **int32** | Number of cache hits for a VCL service. | [optional] 
**AllMissRequests** | Pointer to **int32** | Number of cache misses for a VCL service. | [optional] 
**AllPassRequests** | Pointer to **int32** | Number of requests that passed through the CDN without being cached for a VCL service. | [optional] 
**AllErrorRequests** | Pointer to **int32** | Number of cache errors for a VCL service. | [optional] 
**AllSynthRequests** | Pointer to **int32** | Number of requests that returned a synthetic response (i.e., response objects created with the `synthetic` VCL statement) for a VCL service. | [optional] 
**AllEdgeHitRequests** | Pointer to **int32** | Number of requests sent by end users to Fastly that resulted in a hit at the edge for a VCL service. | [optional] 
**AllEdgeMissRequests** | Pointer to **int32** | Number of requests sent by end users to Fastly that resulted in a miss at the edge for a VCL service. | [optional] 
**AllStatus1xx** | Pointer to **int32** | Number of \&quot;Informational\&quot; category status codes delivered for all sources. | [optional] 
**AllStatus2xx** | Pointer to **int32** | Number of \&quot;Success\&quot; status codes delivered for all sources. | [optional] 
**AllStatus3xx** | Pointer to **int32** | Number of \&quot;Redirection\&quot; codes delivered for all sources. | [optional] 
**AllStatus4xx** | Pointer to **int32** | Number of \&quot;Client Error\&quot; codes delivered for all sources. | [optional] 
**AllStatus5xx** | Pointer to **int32** | Number of \&quot;Server Error\&quot; codes delivered for all sources. | [optional] 
**OriginOffload** | Pointer to **float32** | Origin Offload measures the ratio of bytes served to end users that were cached by Fastly, over the bytes served to end users, between 0 and 1. ((`edge_resp_body_bytes` + `edge_resp_header_bytes`) - (`origin_fetch_resp_body_bytes` + `origin_fetch_resp_header_bytes`)) / (`edge_resp_body_bytes` + `edge_resp_header_bytes`). | [optional] 
**RequestDeniedGetHeadBody** | Pointer to **int32** | Number of requests where Fastly responded with 400 due to the request being a GET or HEAD request containing a body. | [optional] 

## Methods

### NewRealtimeMeasurements

`func NewRealtimeMeasurements() *RealtimeMeasurements`

NewRealtimeMeasurements instantiates a new RealtimeMeasurements object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRealtimeMeasurementsWithDefaults

`func NewRealtimeMeasurementsWithDefaults() *RealtimeMeasurements`

NewRealtimeMeasurementsWithDefaults instantiates a new RealtimeMeasurements object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequests

`func (o *RealtimeMeasurements) GetRequests() int32`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *RealtimeMeasurements) GetRequestsOk() (*int32, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *RealtimeMeasurements) SetRequests(v int32)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *RealtimeMeasurements) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### GetLogging

`func (o *RealtimeMeasurements) GetLogging() int32`

GetLogging returns the Logging field if non-nil, zero value otherwise.

### GetLoggingOk

`func (o *RealtimeMeasurements) GetLoggingOk() (*int32, bool)`

GetLoggingOk returns a tuple with the Logging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogging

`func (o *RealtimeMeasurements) SetLogging(v int32)`

SetLogging sets Logging field to given value.

### HasLogging

`func (o *RealtimeMeasurements) HasLogging() bool`

HasLogging returns a boolean if a field has been set.

### GetLog

`func (o *RealtimeMeasurements) GetLog() int32`

GetLog returns the Log field if non-nil, zero value otherwise.

### GetLogOk

`func (o *RealtimeMeasurements) GetLogOk() (*int32, bool)`

GetLogOk returns a tuple with the Log field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLog

`func (o *RealtimeMeasurements) SetLog(v int32)`

SetLog sets Log field to given value.

### HasLog

`func (o *RealtimeMeasurements) HasLog() bool`

HasLog returns a boolean if a field has been set.

### GetRespHeaderBytes

`func (o *RealtimeMeasurements) GetRespHeaderBytes() int32`

GetRespHeaderBytes returns the RespHeaderBytes field if non-nil, zero value otherwise.

### GetRespHeaderBytesOk

`func (o *RealtimeMeasurements) GetRespHeaderBytesOk() (*int32, bool)`

GetRespHeaderBytesOk returns a tuple with the RespHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRespHeaderBytes

`func (o *RealtimeMeasurements) SetRespHeaderBytes(v int32)`

SetRespHeaderBytes sets RespHeaderBytes field to given value.

### HasRespHeaderBytes

`func (o *RealtimeMeasurements) HasRespHeaderBytes() bool`

HasRespHeaderBytes returns a boolean if a field has been set.

### GetHeaderSize

`func (o *RealtimeMeasurements) GetHeaderSize() int32`

GetHeaderSize returns the HeaderSize field if non-nil, zero value otherwise.

### GetHeaderSizeOk

`func (o *RealtimeMeasurements) GetHeaderSizeOk() (*int32, bool)`

GetHeaderSizeOk returns a tuple with the HeaderSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderSize

`func (o *RealtimeMeasurements) SetHeaderSize(v int32)`

SetHeaderSize sets HeaderSize field to given value.

### HasHeaderSize

`func (o *RealtimeMeasurements) HasHeaderSize() bool`

HasHeaderSize returns a boolean if a field has been set.

### GetRespBodyBytes

`func (o *RealtimeMeasurements) GetRespBodyBytes() int32`

GetRespBodyBytes returns the RespBodyBytes field if non-nil, zero value otherwise.

### GetRespBodyBytesOk

`func (o *RealtimeMeasurements) GetRespBodyBytesOk() (*int32, bool)`

GetRespBodyBytesOk returns a tuple with the RespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRespBodyBytes

`func (o *RealtimeMeasurements) SetRespBodyBytes(v int32)`

SetRespBodyBytes sets RespBodyBytes field to given value.

### HasRespBodyBytes

`func (o *RealtimeMeasurements) HasRespBodyBytes() bool`

HasRespBodyBytes returns a boolean if a field has been set.

### GetBodySize

`func (o *RealtimeMeasurements) GetBodySize() int32`

GetBodySize returns the BodySize field if non-nil, zero value otherwise.

### GetBodySizeOk

`func (o *RealtimeMeasurements) GetBodySizeOk() (*int32, bool)`

GetBodySizeOk returns a tuple with the BodySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodySize

`func (o *RealtimeMeasurements) SetBodySize(v int32)`

SetBodySize sets BodySize field to given value.

### HasBodySize

`func (o *RealtimeMeasurements) HasBodySize() bool`

HasBodySize returns a boolean if a field has been set.

### GetHits

`func (o *RealtimeMeasurements) GetHits() int32`

GetHits returns the Hits field if non-nil, zero value otherwise.

### GetHitsOk

`func (o *RealtimeMeasurements) GetHitsOk() (*int32, bool)`

GetHitsOk returns a tuple with the Hits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHits

`func (o *RealtimeMeasurements) SetHits(v int32)`

SetHits sets Hits field to given value.

### HasHits

`func (o *RealtimeMeasurements) HasHits() bool`

HasHits returns a boolean if a field has been set.

### GetMiss

`func (o *RealtimeMeasurements) GetMiss() int32`

GetMiss returns the Miss field if non-nil, zero value otherwise.

### GetMissOk

`func (o *RealtimeMeasurements) GetMissOk() (*int32, bool)`

GetMissOk returns a tuple with the Miss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiss

`func (o *RealtimeMeasurements) SetMiss(v int32)`

SetMiss sets Miss field to given value.

### HasMiss

`func (o *RealtimeMeasurements) HasMiss() bool`

HasMiss returns a boolean if a field has been set.

### GetPass

`func (o *RealtimeMeasurements) GetPass() int32`

GetPass returns the Pass field if non-nil, zero value otherwise.

### GetPassOk

`func (o *RealtimeMeasurements) GetPassOk() (*int32, bool)`

GetPassOk returns a tuple with the Pass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPass

`func (o *RealtimeMeasurements) SetPass(v int32)`

SetPass sets Pass field to given value.

### HasPass

`func (o *RealtimeMeasurements) HasPass() bool`

HasPass returns a boolean if a field has been set.

### GetSynth

`func (o *RealtimeMeasurements) GetSynth() int32`

GetSynth returns the Synth field if non-nil, zero value otherwise.

### GetSynthOk

`func (o *RealtimeMeasurements) GetSynthOk() (*int32, bool)`

GetSynthOk returns a tuple with the Synth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynth

`func (o *RealtimeMeasurements) SetSynth(v int32)`

SetSynth sets Synth field to given value.

### HasSynth

`func (o *RealtimeMeasurements) HasSynth() bool`

HasSynth returns a boolean if a field has been set.

### GetErrors

`func (o *RealtimeMeasurements) GetErrors() int32`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *RealtimeMeasurements) GetErrorsOk() (*int32, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *RealtimeMeasurements) SetErrors(v int32)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *RealtimeMeasurements) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetHitsTime

`func (o *RealtimeMeasurements) GetHitsTime() float32`

GetHitsTime returns the HitsTime field if non-nil, zero value otherwise.

### GetHitsTimeOk

`func (o *RealtimeMeasurements) GetHitsTimeOk() (*float32, bool)`

GetHitsTimeOk returns a tuple with the HitsTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHitsTime

`func (o *RealtimeMeasurements) SetHitsTime(v float32)`

SetHitsTime sets HitsTime field to given value.

### HasHitsTime

`func (o *RealtimeMeasurements) HasHitsTime() bool`

HasHitsTime returns a boolean if a field has been set.

### GetMissTime

`func (o *RealtimeMeasurements) GetMissTime() float32`

GetMissTime returns the MissTime field if non-nil, zero value otherwise.

### GetMissTimeOk

`func (o *RealtimeMeasurements) GetMissTimeOk() (*float32, bool)`

GetMissTimeOk returns a tuple with the MissTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissTime

`func (o *RealtimeMeasurements) SetMissTime(v float32)`

SetMissTime sets MissTime field to given value.

### HasMissTime

`func (o *RealtimeMeasurements) HasMissTime() bool`

HasMissTime returns a boolean if a field has been set.

### GetMissHistogram

`func (o *RealtimeMeasurements) GetMissHistogram() map[string]map[string]any`

GetMissHistogram returns the MissHistogram field if non-nil, zero value otherwise.

### GetMissHistogramOk

`func (o *RealtimeMeasurements) GetMissHistogramOk() (*map[string]map[string]any, bool)`

GetMissHistogramOk returns a tuple with the MissHistogram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissHistogram

`func (o *RealtimeMeasurements) SetMissHistogram(v map[string]map[string]any)`

SetMissHistogram sets MissHistogram field to given value.

### HasMissHistogram

`func (o *RealtimeMeasurements) HasMissHistogram() bool`

HasMissHistogram returns a boolean if a field has been set.

### GetComputeRequests

`func (o *RealtimeMeasurements) GetComputeRequests() int32`

GetComputeRequests returns the ComputeRequests field if non-nil, zero value otherwise.

### GetComputeRequestsOk

`func (o *RealtimeMeasurements) GetComputeRequestsOk() (*int32, bool)`

GetComputeRequestsOk returns a tuple with the ComputeRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRequests

`func (o *RealtimeMeasurements) SetComputeRequests(v int32)`

SetComputeRequests sets ComputeRequests field to given value.

### HasComputeRequests

`func (o *RealtimeMeasurements) HasComputeRequests() bool`

HasComputeRequests returns a boolean if a field has been set.

### GetComputeExecutionTimeMs

`func (o *RealtimeMeasurements) GetComputeExecutionTimeMs() float32`

GetComputeExecutionTimeMs returns the ComputeExecutionTimeMs field if non-nil, zero value otherwise.

### GetComputeExecutionTimeMsOk

`func (o *RealtimeMeasurements) GetComputeExecutionTimeMsOk() (*float32, bool)`

GetComputeExecutionTimeMsOk returns a tuple with the ComputeExecutionTimeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeExecutionTimeMs

`func (o *RealtimeMeasurements) SetComputeExecutionTimeMs(v float32)`

SetComputeExecutionTimeMs sets ComputeExecutionTimeMs field to given value.

### HasComputeExecutionTimeMs

`func (o *RealtimeMeasurements) HasComputeExecutionTimeMs() bool`

HasComputeExecutionTimeMs returns a boolean if a field has been set.

### GetComputeRAMUsed

`func (o *RealtimeMeasurements) GetComputeRAMUsed() int32`

GetComputeRAMUsed returns the ComputeRAMUsed field if non-nil, zero value otherwise.

### GetComputeRAMUsedOk

`func (o *RealtimeMeasurements) GetComputeRAMUsedOk() (*int32, bool)`

GetComputeRAMUsedOk returns a tuple with the ComputeRAMUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRAMUsed

`func (o *RealtimeMeasurements) SetComputeRAMUsed(v int32)`

SetComputeRAMUsed sets ComputeRAMUsed field to given value.

### HasComputeRAMUsed

`func (o *RealtimeMeasurements) HasComputeRAMUsed() bool`

HasComputeRAMUsed returns a boolean if a field has been set.

### GetComputeRequestTimeMs

`func (o *RealtimeMeasurements) GetComputeRequestTimeMs() float32`

GetComputeRequestTimeMs returns the ComputeRequestTimeMs field if non-nil, zero value otherwise.

### GetComputeRequestTimeMsOk

`func (o *RealtimeMeasurements) GetComputeRequestTimeMsOk() (*float32, bool)`

GetComputeRequestTimeMsOk returns a tuple with the ComputeRequestTimeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRequestTimeMs

`func (o *RealtimeMeasurements) SetComputeRequestTimeMs(v float32)`

SetComputeRequestTimeMs sets ComputeRequestTimeMs field to given value.

### HasComputeRequestTimeMs

`func (o *RealtimeMeasurements) HasComputeRequestTimeMs() bool`

HasComputeRequestTimeMs returns a boolean if a field has been set.

### GetComputeRequestTimeBilledMs

`func (o *RealtimeMeasurements) GetComputeRequestTimeBilledMs() float32`

GetComputeRequestTimeBilledMs returns the ComputeRequestTimeBilledMs field if non-nil, zero value otherwise.

### GetComputeRequestTimeBilledMsOk

`func (o *RealtimeMeasurements) GetComputeRequestTimeBilledMsOk() (*float32, bool)`

GetComputeRequestTimeBilledMsOk returns a tuple with the ComputeRequestTimeBilledMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRequestTimeBilledMs

`func (o *RealtimeMeasurements) SetComputeRequestTimeBilledMs(v float32)`

SetComputeRequestTimeBilledMs sets ComputeRequestTimeBilledMs field to given value.

### HasComputeRequestTimeBilledMs

`func (o *RealtimeMeasurements) HasComputeRequestTimeBilledMs() bool`

HasComputeRequestTimeBilledMs returns a boolean if a field has been set.

### GetShield

`func (o *RealtimeMeasurements) GetShield() int32`

GetShield returns the Shield field if non-nil, zero value otherwise.

### GetShieldOk

`func (o *RealtimeMeasurements) GetShieldOk() (*int32, bool)`

GetShieldOk returns a tuple with the Shield field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShield

`func (o *RealtimeMeasurements) SetShield(v int32)`

SetShield sets Shield field to given value.

### HasShield

`func (o *RealtimeMeasurements) HasShield() bool`

HasShield returns a boolean if a field has been set.

### GetIpv6

`func (o *RealtimeMeasurements) GetIpv6() int32`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *RealtimeMeasurements) GetIpv6Ok() (*int32, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *RealtimeMeasurements) SetIpv6(v int32)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *RealtimeMeasurements) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.

### GetImgopto

`func (o *RealtimeMeasurements) GetImgopto() int32`

GetImgopto returns the Imgopto field if non-nil, zero value otherwise.

### GetImgoptoOk

`func (o *RealtimeMeasurements) GetImgoptoOk() (*int32, bool)`

GetImgoptoOk returns a tuple with the Imgopto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImgopto

`func (o *RealtimeMeasurements) SetImgopto(v int32)`

SetImgopto sets Imgopto field to given value.

### HasImgopto

`func (o *RealtimeMeasurements) HasImgopto() bool`

HasImgopto returns a boolean if a field has been set.

### GetImgoptoShield

`func (o *RealtimeMeasurements) GetImgoptoShield() int32`

GetImgoptoShield returns the ImgoptoShield field if non-nil, zero value otherwise.

### GetImgoptoShieldOk

`func (o *RealtimeMeasurements) GetImgoptoShieldOk() (*int32, bool)`

GetImgoptoShieldOk returns a tuple with the ImgoptoShield field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImgoptoShield

`func (o *RealtimeMeasurements) SetImgoptoShield(v int32)`

SetImgoptoShield sets ImgoptoShield field to given value.

### HasImgoptoShield

`func (o *RealtimeMeasurements) HasImgoptoShield() bool`

HasImgoptoShield returns a boolean if a field has been set.

### GetImgoptoTransforms

`func (o *RealtimeMeasurements) GetImgoptoTransforms() int32`

GetImgoptoTransforms returns the ImgoptoTransforms field if non-nil, zero value otherwise.

### GetImgoptoTransformsOk

`func (o *RealtimeMeasurements) GetImgoptoTransformsOk() (*int32, bool)`

GetImgoptoTransformsOk returns a tuple with the ImgoptoTransforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImgoptoTransforms

`func (o *RealtimeMeasurements) SetImgoptoTransforms(v int32)`

SetImgoptoTransforms sets ImgoptoTransforms field to given value.

### HasImgoptoTransforms

`func (o *RealtimeMeasurements) HasImgoptoTransforms() bool`

HasImgoptoTransforms returns a boolean if a field has been set.

### GetOtfp

`func (o *RealtimeMeasurements) GetOtfp() int32`

GetOtfp returns the Otfp field if non-nil, zero value otherwise.

### GetOtfpOk

`func (o *RealtimeMeasurements) GetOtfpOk() (*int32, bool)`

GetOtfpOk returns a tuple with the Otfp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtfp

`func (o *RealtimeMeasurements) SetOtfp(v int32)`

SetOtfp sets Otfp field to given value.

### HasOtfp

`func (o *RealtimeMeasurements) HasOtfp() bool`

HasOtfp returns a boolean if a field has been set.

### GetOtfpShield

`func (o *RealtimeMeasurements) GetOtfpShield() int32`

GetOtfpShield returns the OtfpShield field if non-nil, zero value otherwise.

### GetOtfpShieldOk

`func (o *RealtimeMeasurements) GetOtfpShieldOk() (*int32, bool)`

GetOtfpShieldOk returns a tuple with the OtfpShield field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtfpShield

`func (o *RealtimeMeasurements) SetOtfpShield(v int32)`

SetOtfpShield sets OtfpShield field to given value.

### HasOtfpShield

`func (o *RealtimeMeasurements) HasOtfpShield() bool`

HasOtfpShield returns a boolean if a field has been set.

### GetOtfpManifests

`func (o *RealtimeMeasurements) GetOtfpManifests() int32`

GetOtfpManifests returns the OtfpManifests field if non-nil, zero value otherwise.

### GetOtfpManifestsOk

`func (o *RealtimeMeasurements) GetOtfpManifestsOk() (*int32, bool)`

GetOtfpManifestsOk returns a tuple with the OtfpManifests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtfpManifests

`func (o *RealtimeMeasurements) SetOtfpManifests(v int32)`

SetOtfpManifests sets OtfpManifests field to given value.

### HasOtfpManifests

`func (o *RealtimeMeasurements) HasOtfpManifests() bool`

HasOtfpManifests returns a boolean if a field has been set.

### GetVideo

`func (o *RealtimeMeasurements) GetVideo() int32`

GetVideo returns the Video field if non-nil, zero value otherwise.

### GetVideoOk

`func (o *RealtimeMeasurements) GetVideoOk() (*int32, bool)`

GetVideoOk returns a tuple with the Video field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideo

`func (o *RealtimeMeasurements) SetVideo(v int32)`

SetVideo sets Video field to given value.

### HasVideo

`func (o *RealtimeMeasurements) HasVideo() bool`

HasVideo returns a boolean if a field has been set.

### GetPci

`func (o *RealtimeMeasurements) GetPci() int32`

GetPci returns the Pci field if non-nil, zero value otherwise.

### GetPciOk

`func (o *RealtimeMeasurements) GetPciOk() (*int32, bool)`

GetPciOk returns a tuple with the Pci field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPci

`func (o *RealtimeMeasurements) SetPci(v int32)`

SetPci sets Pci field to given value.

### HasPci

`func (o *RealtimeMeasurements) HasPci() bool`

HasPci returns a boolean if a field has been set.

### GetHTTP2

`func (o *RealtimeMeasurements) GetHTTP2() int32`

GetHTTP2 returns the HTTP2 field if non-nil, zero value otherwise.

### GetHTTP2Ok

`func (o *RealtimeMeasurements) GetHTTP2Ok() (*int32, bool)`

GetHTTP2Ok returns a tuple with the HTTP2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHTTP2

`func (o *RealtimeMeasurements) SetHTTP2(v int32)`

SetHTTP2 sets HTTP2 field to given value.

### HasHTTP2

`func (o *RealtimeMeasurements) HasHTTP2() bool`

HasHTTP2 returns a boolean if a field has been set.

### GetHTTP3

`func (o *RealtimeMeasurements) GetHTTP3() int32`

GetHTTP3 returns the HTTP3 field if non-nil, zero value otherwise.

### GetHTTP3Ok

`func (o *RealtimeMeasurements) GetHTTP3Ok() (*int32, bool)`

GetHTTP3Ok returns a tuple with the HTTP3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHTTP3

`func (o *RealtimeMeasurements) SetHTTP3(v int32)`

SetHTTP3 sets HTTP3 field to given value.

### HasHTTP3

`func (o *RealtimeMeasurements) HasHTTP3() bool`

HasHTTP3 returns a boolean if a field has been set.

### GetRestarts

`func (o *RealtimeMeasurements) GetRestarts() int32`

GetRestarts returns the Restarts field if non-nil, zero value otherwise.

### GetRestartsOk

`func (o *RealtimeMeasurements) GetRestartsOk() (*int32, bool)`

GetRestartsOk returns a tuple with the Restarts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestarts

`func (o *RealtimeMeasurements) SetRestarts(v int32)`

SetRestarts sets Restarts field to given value.

### HasRestarts

`func (o *RealtimeMeasurements) HasRestarts() bool`

HasRestarts returns a boolean if a field has been set.

### GetReqHeaderBytes

`func (o *RealtimeMeasurements) GetReqHeaderBytes() int32`

GetReqHeaderBytes returns the ReqHeaderBytes field if non-nil, zero value otherwise.

### GetReqHeaderBytesOk

`func (o *RealtimeMeasurements) GetReqHeaderBytesOk() (*int32, bool)`

GetReqHeaderBytesOk returns a tuple with the ReqHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqHeaderBytes

`func (o *RealtimeMeasurements) SetReqHeaderBytes(v int32)`

SetReqHeaderBytes sets ReqHeaderBytes field to given value.

### HasReqHeaderBytes

`func (o *RealtimeMeasurements) HasReqHeaderBytes() bool`

HasReqHeaderBytes returns a boolean if a field has been set.

### GetReqBodyBytes

`func (o *RealtimeMeasurements) GetReqBodyBytes() int32`

GetReqBodyBytes returns the ReqBodyBytes field if non-nil, zero value otherwise.

### GetReqBodyBytesOk

`func (o *RealtimeMeasurements) GetReqBodyBytesOk() (*int32, bool)`

GetReqBodyBytesOk returns a tuple with the ReqBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqBodyBytes

`func (o *RealtimeMeasurements) SetReqBodyBytes(v int32)`

SetReqBodyBytes sets ReqBodyBytes field to given value.

### HasReqBodyBytes

`func (o *RealtimeMeasurements) HasReqBodyBytes() bool`

HasReqBodyBytes returns a boolean if a field has been set.

### GetBereqHeaderBytes

`func (o *RealtimeMeasurements) GetBereqHeaderBytes() int32`

GetBereqHeaderBytes returns the BereqHeaderBytes field if non-nil, zero value otherwise.

### GetBereqHeaderBytesOk

`func (o *RealtimeMeasurements) GetBereqHeaderBytesOk() (*int32, bool)`

GetBereqHeaderBytesOk returns a tuple with the BereqHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBereqHeaderBytes

`func (o *RealtimeMeasurements) SetBereqHeaderBytes(v int32)`

SetBereqHeaderBytes sets BereqHeaderBytes field to given value.

### HasBereqHeaderBytes

`func (o *RealtimeMeasurements) HasBereqHeaderBytes() bool`

HasBereqHeaderBytes returns a boolean if a field has been set.

### GetBereqBodyBytes

`func (o *RealtimeMeasurements) GetBereqBodyBytes() int32`

GetBereqBodyBytes returns the BereqBodyBytes field if non-nil, zero value otherwise.

### GetBereqBodyBytesOk

`func (o *RealtimeMeasurements) GetBereqBodyBytesOk() (*int32, bool)`

GetBereqBodyBytesOk returns a tuple with the BereqBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBereqBodyBytes

`func (o *RealtimeMeasurements) SetBereqBodyBytes(v int32)`

SetBereqBodyBytes sets BereqBodyBytes field to given value.

### HasBereqBodyBytes

`func (o *RealtimeMeasurements) HasBereqBodyBytes() bool`

HasBereqBodyBytes returns a boolean if a field has been set.

### GetWafBlocked

`func (o *RealtimeMeasurements) GetWafBlocked() int32`

GetWafBlocked returns the WafBlocked field if non-nil, zero value otherwise.

### GetWafBlockedOk

`func (o *RealtimeMeasurements) GetWafBlockedOk() (*int32, bool)`

GetWafBlockedOk returns a tuple with the WafBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafBlocked

`func (o *RealtimeMeasurements) SetWafBlocked(v int32)`

SetWafBlocked sets WafBlocked field to given value.

### HasWafBlocked

`func (o *RealtimeMeasurements) HasWafBlocked() bool`

HasWafBlocked returns a boolean if a field has been set.

### GetWafLogged

`func (o *RealtimeMeasurements) GetWafLogged() int32`

GetWafLogged returns the WafLogged field if non-nil, zero value otherwise.

### GetWafLoggedOk

`func (o *RealtimeMeasurements) GetWafLoggedOk() (*int32, bool)`

GetWafLoggedOk returns a tuple with the WafLogged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafLogged

`func (o *RealtimeMeasurements) SetWafLogged(v int32)`

SetWafLogged sets WafLogged field to given value.

### HasWafLogged

`func (o *RealtimeMeasurements) HasWafLogged() bool`

HasWafLogged returns a boolean if a field has been set.

### GetWafPassed

`func (o *RealtimeMeasurements) GetWafPassed() int32`

GetWafPassed returns the WafPassed field if non-nil, zero value otherwise.

### GetWafPassedOk

`func (o *RealtimeMeasurements) GetWafPassedOk() (*int32, bool)`

GetWafPassedOk returns a tuple with the WafPassed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafPassed

`func (o *RealtimeMeasurements) SetWafPassed(v int32)`

SetWafPassed sets WafPassed field to given value.

### HasWafPassed

`func (o *RealtimeMeasurements) HasWafPassed() bool`

HasWafPassed returns a boolean if a field has been set.

### GetAttackReqHeaderBytes

`func (o *RealtimeMeasurements) GetAttackReqHeaderBytes() int32`

GetAttackReqHeaderBytes returns the AttackReqHeaderBytes field if non-nil, zero value otherwise.

### GetAttackReqHeaderBytesOk

`func (o *RealtimeMeasurements) GetAttackReqHeaderBytesOk() (*int32, bool)`

GetAttackReqHeaderBytesOk returns a tuple with the AttackReqHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackReqHeaderBytes

`func (o *RealtimeMeasurements) SetAttackReqHeaderBytes(v int32)`

SetAttackReqHeaderBytes sets AttackReqHeaderBytes field to given value.

### HasAttackReqHeaderBytes

`func (o *RealtimeMeasurements) HasAttackReqHeaderBytes() bool`

HasAttackReqHeaderBytes returns a boolean if a field has been set.

### GetAttackReqBodyBytes

`func (o *RealtimeMeasurements) GetAttackReqBodyBytes() int32`

GetAttackReqBodyBytes returns the AttackReqBodyBytes field if non-nil, zero value otherwise.

### GetAttackReqBodyBytesOk

`func (o *RealtimeMeasurements) GetAttackReqBodyBytesOk() (*int32, bool)`

GetAttackReqBodyBytesOk returns a tuple with the AttackReqBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackReqBodyBytes

`func (o *RealtimeMeasurements) SetAttackReqBodyBytes(v int32)`

SetAttackReqBodyBytes sets AttackReqBodyBytes field to given value.

### HasAttackReqBodyBytes

`func (o *RealtimeMeasurements) HasAttackReqBodyBytes() bool`

HasAttackReqBodyBytes returns a boolean if a field has been set.

### GetAttackRespSynthBytes

`func (o *RealtimeMeasurements) GetAttackRespSynthBytes() int32`

GetAttackRespSynthBytes returns the AttackRespSynthBytes field if non-nil, zero value otherwise.

### GetAttackRespSynthBytesOk

`func (o *RealtimeMeasurements) GetAttackRespSynthBytesOk() (*int32, bool)`

GetAttackRespSynthBytesOk returns a tuple with the AttackRespSynthBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackRespSynthBytes

`func (o *RealtimeMeasurements) SetAttackRespSynthBytes(v int32)`

SetAttackRespSynthBytes sets AttackRespSynthBytes field to given value.

### HasAttackRespSynthBytes

`func (o *RealtimeMeasurements) HasAttackRespSynthBytes() bool`

HasAttackRespSynthBytes returns a boolean if a field has been set.

### GetAttackLoggedReqHeaderBytes

`func (o *RealtimeMeasurements) GetAttackLoggedReqHeaderBytes() int32`

GetAttackLoggedReqHeaderBytes returns the AttackLoggedReqHeaderBytes field if non-nil, zero value otherwise.

### GetAttackLoggedReqHeaderBytesOk

`func (o *RealtimeMeasurements) GetAttackLoggedReqHeaderBytesOk() (*int32, bool)`

GetAttackLoggedReqHeaderBytesOk returns a tuple with the AttackLoggedReqHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackLoggedReqHeaderBytes

`func (o *RealtimeMeasurements) SetAttackLoggedReqHeaderBytes(v int32)`

SetAttackLoggedReqHeaderBytes sets AttackLoggedReqHeaderBytes field to given value.

### HasAttackLoggedReqHeaderBytes

`func (o *RealtimeMeasurements) HasAttackLoggedReqHeaderBytes() bool`

HasAttackLoggedReqHeaderBytes returns a boolean if a field has been set.

### GetAttackLoggedReqBodyBytes

`func (o *RealtimeMeasurements) GetAttackLoggedReqBodyBytes() int32`

GetAttackLoggedReqBodyBytes returns the AttackLoggedReqBodyBytes field if non-nil, zero value otherwise.

### GetAttackLoggedReqBodyBytesOk

`func (o *RealtimeMeasurements) GetAttackLoggedReqBodyBytesOk() (*int32, bool)`

GetAttackLoggedReqBodyBytesOk returns a tuple with the AttackLoggedReqBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackLoggedReqBodyBytes

`func (o *RealtimeMeasurements) SetAttackLoggedReqBodyBytes(v int32)`

SetAttackLoggedReqBodyBytes sets AttackLoggedReqBodyBytes field to given value.

### HasAttackLoggedReqBodyBytes

`func (o *RealtimeMeasurements) HasAttackLoggedReqBodyBytes() bool`

HasAttackLoggedReqBodyBytes returns a boolean if a field has been set.

### GetAttackBlockedReqHeaderBytes

`func (o *RealtimeMeasurements) GetAttackBlockedReqHeaderBytes() int32`

GetAttackBlockedReqHeaderBytes returns the AttackBlockedReqHeaderBytes field if non-nil, zero value otherwise.

### GetAttackBlockedReqHeaderBytesOk

`func (o *RealtimeMeasurements) GetAttackBlockedReqHeaderBytesOk() (*int32, bool)`

GetAttackBlockedReqHeaderBytesOk returns a tuple with the AttackBlockedReqHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackBlockedReqHeaderBytes

`func (o *RealtimeMeasurements) SetAttackBlockedReqHeaderBytes(v int32)`

SetAttackBlockedReqHeaderBytes sets AttackBlockedReqHeaderBytes field to given value.

### HasAttackBlockedReqHeaderBytes

`func (o *RealtimeMeasurements) HasAttackBlockedReqHeaderBytes() bool`

HasAttackBlockedReqHeaderBytes returns a boolean if a field has been set.

### GetAttackBlockedReqBodyBytes

`func (o *RealtimeMeasurements) GetAttackBlockedReqBodyBytes() int32`

GetAttackBlockedReqBodyBytes returns the AttackBlockedReqBodyBytes field if non-nil, zero value otherwise.

### GetAttackBlockedReqBodyBytesOk

`func (o *RealtimeMeasurements) GetAttackBlockedReqBodyBytesOk() (*int32, bool)`

GetAttackBlockedReqBodyBytesOk returns a tuple with the AttackBlockedReqBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackBlockedReqBodyBytes

`func (o *RealtimeMeasurements) SetAttackBlockedReqBodyBytes(v int32)`

SetAttackBlockedReqBodyBytes sets AttackBlockedReqBodyBytes field to given value.

### HasAttackBlockedReqBodyBytes

`func (o *RealtimeMeasurements) HasAttackBlockedReqBodyBytes() bool`

HasAttackBlockedReqBodyBytes returns a boolean if a field has been set.

### GetAttackPassedReqHeaderBytes

`func (o *RealtimeMeasurements) GetAttackPassedReqHeaderBytes() int32`

GetAttackPassedReqHeaderBytes returns the AttackPassedReqHeaderBytes field if non-nil, zero value otherwise.

### GetAttackPassedReqHeaderBytesOk

`func (o *RealtimeMeasurements) GetAttackPassedReqHeaderBytesOk() (*int32, bool)`

GetAttackPassedReqHeaderBytesOk returns a tuple with the AttackPassedReqHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackPassedReqHeaderBytes

`func (o *RealtimeMeasurements) SetAttackPassedReqHeaderBytes(v int32)`

SetAttackPassedReqHeaderBytes sets AttackPassedReqHeaderBytes field to given value.

### HasAttackPassedReqHeaderBytes

`func (o *RealtimeMeasurements) HasAttackPassedReqHeaderBytes() bool`

HasAttackPassedReqHeaderBytes returns a boolean if a field has been set.

### GetAttackPassedReqBodyBytes

`func (o *RealtimeMeasurements) GetAttackPassedReqBodyBytes() int32`

GetAttackPassedReqBodyBytes returns the AttackPassedReqBodyBytes field if non-nil, zero value otherwise.

### GetAttackPassedReqBodyBytesOk

`func (o *RealtimeMeasurements) GetAttackPassedReqBodyBytesOk() (*int32, bool)`

GetAttackPassedReqBodyBytesOk returns a tuple with the AttackPassedReqBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackPassedReqBodyBytes

`func (o *RealtimeMeasurements) SetAttackPassedReqBodyBytes(v int32)`

SetAttackPassedReqBodyBytes sets AttackPassedReqBodyBytes field to given value.

### HasAttackPassedReqBodyBytes

`func (o *RealtimeMeasurements) HasAttackPassedReqBodyBytes() bool`

HasAttackPassedReqBodyBytes returns a boolean if a field has been set.

### GetShieldRespHeaderBytes

`func (o *RealtimeMeasurements) GetShieldRespHeaderBytes() int32`

GetShieldRespHeaderBytes returns the ShieldRespHeaderBytes field if non-nil, zero value otherwise.

### GetShieldRespHeaderBytesOk

`func (o *RealtimeMeasurements) GetShieldRespHeaderBytesOk() (*int32, bool)`

GetShieldRespHeaderBytesOk returns a tuple with the ShieldRespHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShieldRespHeaderBytes

`func (o *RealtimeMeasurements) SetShieldRespHeaderBytes(v int32)`

SetShieldRespHeaderBytes sets ShieldRespHeaderBytes field to given value.

### HasShieldRespHeaderBytes

`func (o *RealtimeMeasurements) HasShieldRespHeaderBytes() bool`

HasShieldRespHeaderBytes returns a boolean if a field has been set.

### GetShieldRespBodyBytes

`func (o *RealtimeMeasurements) GetShieldRespBodyBytes() int32`

GetShieldRespBodyBytes returns the ShieldRespBodyBytes field if non-nil, zero value otherwise.

### GetShieldRespBodyBytesOk

`func (o *RealtimeMeasurements) GetShieldRespBodyBytesOk() (*int32, bool)`

GetShieldRespBodyBytesOk returns a tuple with the ShieldRespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShieldRespBodyBytes

`func (o *RealtimeMeasurements) SetShieldRespBodyBytes(v int32)`

SetShieldRespBodyBytes sets ShieldRespBodyBytes field to given value.

### HasShieldRespBodyBytes

`func (o *RealtimeMeasurements) HasShieldRespBodyBytes() bool`

HasShieldRespBodyBytes returns a boolean if a field has been set.

### GetOtfpRespHeaderBytes

`func (o *RealtimeMeasurements) GetOtfpRespHeaderBytes() int32`

GetOtfpRespHeaderBytes returns the OtfpRespHeaderBytes field if non-nil, zero value otherwise.

### GetOtfpRespHeaderBytesOk

`func (o *RealtimeMeasurements) GetOtfpRespHeaderBytesOk() (*int32, bool)`

GetOtfpRespHeaderBytesOk returns a tuple with the OtfpRespHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtfpRespHeaderBytes

`func (o *RealtimeMeasurements) SetOtfpRespHeaderBytes(v int32)`

SetOtfpRespHeaderBytes sets OtfpRespHeaderBytes field to given value.

### HasOtfpRespHeaderBytes

`func (o *RealtimeMeasurements) HasOtfpRespHeaderBytes() bool`

HasOtfpRespHeaderBytes returns a boolean if a field has been set.

### GetOtfpRespBodyBytes

`func (o *RealtimeMeasurements) GetOtfpRespBodyBytes() int32`

GetOtfpRespBodyBytes returns the OtfpRespBodyBytes field if non-nil, zero value otherwise.

### GetOtfpRespBodyBytesOk

`func (o *RealtimeMeasurements) GetOtfpRespBodyBytesOk() (*int32, bool)`

GetOtfpRespBodyBytesOk returns a tuple with the OtfpRespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtfpRespBodyBytes

`func (o *RealtimeMeasurements) SetOtfpRespBodyBytes(v int32)`

SetOtfpRespBodyBytes sets OtfpRespBodyBytes field to given value.

### HasOtfpRespBodyBytes

`func (o *RealtimeMeasurements) HasOtfpRespBodyBytes() bool`

HasOtfpRespBodyBytes returns a boolean if a field has been set.

### GetOtfpShieldRespHeaderBytes

`func (o *RealtimeMeasurements) GetOtfpShieldRespHeaderBytes() int32`

GetOtfpShieldRespHeaderBytes returns the OtfpShieldRespHeaderBytes field if non-nil, zero value otherwise.

### GetOtfpShieldRespHeaderBytesOk

`func (o *RealtimeMeasurements) GetOtfpShieldRespHeaderBytesOk() (*int32, bool)`

GetOtfpShieldRespHeaderBytesOk returns a tuple with the OtfpShieldRespHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtfpShieldRespHeaderBytes

`func (o *RealtimeMeasurements) SetOtfpShieldRespHeaderBytes(v int32)`

SetOtfpShieldRespHeaderBytes sets OtfpShieldRespHeaderBytes field to given value.

### HasOtfpShieldRespHeaderBytes

`func (o *RealtimeMeasurements) HasOtfpShieldRespHeaderBytes() bool`

HasOtfpShieldRespHeaderBytes returns a boolean if a field has been set.

### GetOtfpShieldRespBodyBytes

`func (o *RealtimeMeasurements) GetOtfpShieldRespBodyBytes() int32`

GetOtfpShieldRespBodyBytes returns the OtfpShieldRespBodyBytes field if non-nil, zero value otherwise.

### GetOtfpShieldRespBodyBytesOk

`func (o *RealtimeMeasurements) GetOtfpShieldRespBodyBytesOk() (*int32, bool)`

GetOtfpShieldRespBodyBytesOk returns a tuple with the OtfpShieldRespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtfpShieldRespBodyBytes

`func (o *RealtimeMeasurements) SetOtfpShieldRespBodyBytes(v int32)`

SetOtfpShieldRespBodyBytes sets OtfpShieldRespBodyBytes field to given value.

### HasOtfpShieldRespBodyBytes

`func (o *RealtimeMeasurements) HasOtfpShieldRespBodyBytes() bool`

HasOtfpShieldRespBodyBytes returns a boolean if a field has been set.

### GetOtfpShieldTime

`func (o *RealtimeMeasurements) GetOtfpShieldTime() float32`

GetOtfpShieldTime returns the OtfpShieldTime field if non-nil, zero value otherwise.

### GetOtfpShieldTimeOk

`func (o *RealtimeMeasurements) GetOtfpShieldTimeOk() (*float32, bool)`

GetOtfpShieldTimeOk returns a tuple with the OtfpShieldTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtfpShieldTime

`func (o *RealtimeMeasurements) SetOtfpShieldTime(v float32)`

SetOtfpShieldTime sets OtfpShieldTime field to given value.

### HasOtfpShieldTime

`func (o *RealtimeMeasurements) HasOtfpShieldTime() bool`

HasOtfpShieldTime returns a boolean if a field has been set.

### GetOtfpDeliverTime

`func (o *RealtimeMeasurements) GetOtfpDeliverTime() float32`

GetOtfpDeliverTime returns the OtfpDeliverTime field if non-nil, zero value otherwise.

### GetOtfpDeliverTimeOk

`func (o *RealtimeMeasurements) GetOtfpDeliverTimeOk() (*float32, bool)`

GetOtfpDeliverTimeOk returns a tuple with the OtfpDeliverTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtfpDeliverTime

`func (o *RealtimeMeasurements) SetOtfpDeliverTime(v float32)`

SetOtfpDeliverTime sets OtfpDeliverTime field to given value.

### HasOtfpDeliverTime

`func (o *RealtimeMeasurements) HasOtfpDeliverTime() bool`

HasOtfpDeliverTime returns a boolean if a field has been set.

### GetImgoptoRespHeaderBytes

`func (o *RealtimeMeasurements) GetImgoptoRespHeaderBytes() int32`

GetImgoptoRespHeaderBytes returns the ImgoptoRespHeaderBytes field if non-nil, zero value otherwise.

### GetImgoptoRespHeaderBytesOk

`func (o *RealtimeMeasurements) GetImgoptoRespHeaderBytesOk() (*int32, bool)`

GetImgoptoRespHeaderBytesOk returns a tuple with the ImgoptoRespHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImgoptoRespHeaderBytes

`func (o *RealtimeMeasurements) SetImgoptoRespHeaderBytes(v int32)`

SetImgoptoRespHeaderBytes sets ImgoptoRespHeaderBytes field to given value.

### HasImgoptoRespHeaderBytes

`func (o *RealtimeMeasurements) HasImgoptoRespHeaderBytes() bool`

HasImgoptoRespHeaderBytes returns a boolean if a field has been set.

### GetImgoptoRespBodyBytes

`func (o *RealtimeMeasurements) GetImgoptoRespBodyBytes() int32`

GetImgoptoRespBodyBytes returns the ImgoptoRespBodyBytes field if non-nil, zero value otherwise.

### GetImgoptoRespBodyBytesOk

`func (o *RealtimeMeasurements) GetImgoptoRespBodyBytesOk() (*int32, bool)`

GetImgoptoRespBodyBytesOk returns a tuple with the ImgoptoRespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImgoptoRespBodyBytes

`func (o *RealtimeMeasurements) SetImgoptoRespBodyBytes(v int32)`

SetImgoptoRespBodyBytes sets ImgoptoRespBodyBytes field to given value.

### HasImgoptoRespBodyBytes

`func (o *RealtimeMeasurements) HasImgoptoRespBodyBytes() bool`

HasImgoptoRespBodyBytes returns a boolean if a field has been set.

### GetImgoptoShieldRespHeaderBytes

`func (o *RealtimeMeasurements) GetImgoptoShieldRespHeaderBytes() int32`

GetImgoptoShieldRespHeaderBytes returns the ImgoptoShieldRespHeaderBytes field if non-nil, zero value otherwise.

### GetImgoptoShieldRespHeaderBytesOk

`func (o *RealtimeMeasurements) GetImgoptoShieldRespHeaderBytesOk() (*int32, bool)`

GetImgoptoShieldRespHeaderBytesOk returns a tuple with the ImgoptoShieldRespHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImgoptoShieldRespHeaderBytes

`func (o *RealtimeMeasurements) SetImgoptoShieldRespHeaderBytes(v int32)`

SetImgoptoShieldRespHeaderBytes sets ImgoptoShieldRespHeaderBytes field to given value.

### HasImgoptoShieldRespHeaderBytes

`func (o *RealtimeMeasurements) HasImgoptoShieldRespHeaderBytes() bool`

HasImgoptoShieldRespHeaderBytes returns a boolean if a field has been set.

### GetImgoptoShieldRespBodyBytes

`func (o *RealtimeMeasurements) GetImgoptoShieldRespBodyBytes() int32`

GetImgoptoShieldRespBodyBytes returns the ImgoptoShieldRespBodyBytes field if non-nil, zero value otherwise.

### GetImgoptoShieldRespBodyBytesOk

`func (o *RealtimeMeasurements) GetImgoptoShieldRespBodyBytesOk() (*int32, bool)`

GetImgoptoShieldRespBodyBytesOk returns a tuple with the ImgoptoShieldRespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImgoptoShieldRespBodyBytes

`func (o *RealtimeMeasurements) SetImgoptoShieldRespBodyBytes(v int32)`

SetImgoptoShieldRespBodyBytes sets ImgoptoShieldRespBodyBytes field to given value.

### HasImgoptoShieldRespBodyBytes

`func (o *RealtimeMeasurements) HasImgoptoShieldRespBodyBytes() bool`

HasImgoptoShieldRespBodyBytes returns a boolean if a field has been set.

### GetStatus1xx

`func (o *RealtimeMeasurements) GetStatus1xx() int32`

GetStatus1xx returns the Status1xx field if non-nil, zero value otherwise.

### GetStatus1xxOk

`func (o *RealtimeMeasurements) GetStatus1xxOk() (*int32, bool)`

GetStatus1xxOk returns a tuple with the Status1xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus1xx

`func (o *RealtimeMeasurements) SetStatus1xx(v int32)`

SetStatus1xx sets Status1xx field to given value.

### HasStatus1xx

`func (o *RealtimeMeasurements) HasStatus1xx() bool`

HasStatus1xx returns a boolean if a field has been set.

### GetStatus2xx

`func (o *RealtimeMeasurements) GetStatus2xx() int32`

GetStatus2xx returns the Status2xx field if non-nil, zero value otherwise.

### GetStatus2xxOk

`func (o *RealtimeMeasurements) GetStatus2xxOk() (*int32, bool)`

GetStatus2xxOk returns a tuple with the Status2xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus2xx

`func (o *RealtimeMeasurements) SetStatus2xx(v int32)`

SetStatus2xx sets Status2xx field to given value.

### HasStatus2xx

`func (o *RealtimeMeasurements) HasStatus2xx() bool`

HasStatus2xx returns a boolean if a field has been set.

### GetStatus3xx

`func (o *RealtimeMeasurements) GetStatus3xx() int32`

GetStatus3xx returns the Status3xx field if non-nil, zero value otherwise.

### GetStatus3xxOk

`func (o *RealtimeMeasurements) GetStatus3xxOk() (*int32, bool)`

GetStatus3xxOk returns a tuple with the Status3xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus3xx

`func (o *RealtimeMeasurements) SetStatus3xx(v int32)`

SetStatus3xx sets Status3xx field to given value.

### HasStatus3xx

`func (o *RealtimeMeasurements) HasStatus3xx() bool`

HasStatus3xx returns a boolean if a field has been set.

### GetStatus4xx

`func (o *RealtimeMeasurements) GetStatus4xx() int32`

GetStatus4xx returns the Status4xx field if non-nil, zero value otherwise.

### GetStatus4xxOk

`func (o *RealtimeMeasurements) GetStatus4xxOk() (*int32, bool)`

GetStatus4xxOk returns a tuple with the Status4xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus4xx

`func (o *RealtimeMeasurements) SetStatus4xx(v int32)`

SetStatus4xx sets Status4xx field to given value.

### HasStatus4xx

`func (o *RealtimeMeasurements) HasStatus4xx() bool`

HasStatus4xx returns a boolean if a field has been set.

### GetStatus5xx

`func (o *RealtimeMeasurements) GetStatus5xx() int32`

GetStatus5xx returns the Status5xx field if non-nil, zero value otherwise.

### GetStatus5xxOk

`func (o *RealtimeMeasurements) GetStatus5xxOk() (*int32, bool)`

GetStatus5xxOk returns a tuple with the Status5xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus5xx

`func (o *RealtimeMeasurements) SetStatus5xx(v int32)`

SetStatus5xx sets Status5xx field to given value.

### HasStatus5xx

`func (o *RealtimeMeasurements) HasStatus5xx() bool`

HasStatus5xx returns a boolean if a field has been set.

### GetStatus200

`func (o *RealtimeMeasurements) GetStatus200() int32`

GetStatus200 returns the Status200 field if non-nil, zero value otherwise.

### GetStatus200Ok

`func (o *RealtimeMeasurements) GetStatus200Ok() (*int32, bool)`

GetStatus200Ok returns a tuple with the Status200 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus200

`func (o *RealtimeMeasurements) SetStatus200(v int32)`

SetStatus200 sets Status200 field to given value.

### HasStatus200

`func (o *RealtimeMeasurements) HasStatus200() bool`

HasStatus200 returns a boolean if a field has been set.

### GetStatus204

`func (o *RealtimeMeasurements) GetStatus204() int32`

GetStatus204 returns the Status204 field if non-nil, zero value otherwise.

### GetStatus204Ok

`func (o *RealtimeMeasurements) GetStatus204Ok() (*int32, bool)`

GetStatus204Ok returns a tuple with the Status204 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus204

`func (o *RealtimeMeasurements) SetStatus204(v int32)`

SetStatus204 sets Status204 field to given value.

### HasStatus204

`func (o *RealtimeMeasurements) HasStatus204() bool`

HasStatus204 returns a boolean if a field has been set.

### GetStatus206

`func (o *RealtimeMeasurements) GetStatus206() int32`

GetStatus206 returns the Status206 field if non-nil, zero value otherwise.

### GetStatus206Ok

`func (o *RealtimeMeasurements) GetStatus206Ok() (*int32, bool)`

GetStatus206Ok returns a tuple with the Status206 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus206

`func (o *RealtimeMeasurements) SetStatus206(v int32)`

SetStatus206 sets Status206 field to given value.

### HasStatus206

`func (o *RealtimeMeasurements) HasStatus206() bool`

HasStatus206 returns a boolean if a field has been set.

### GetStatus301

`func (o *RealtimeMeasurements) GetStatus301() int32`

GetStatus301 returns the Status301 field if non-nil, zero value otherwise.

### GetStatus301Ok

`func (o *RealtimeMeasurements) GetStatus301Ok() (*int32, bool)`

GetStatus301Ok returns a tuple with the Status301 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus301

`func (o *RealtimeMeasurements) SetStatus301(v int32)`

SetStatus301 sets Status301 field to given value.

### HasStatus301

`func (o *RealtimeMeasurements) HasStatus301() bool`

HasStatus301 returns a boolean if a field has been set.

### GetStatus302

`func (o *RealtimeMeasurements) GetStatus302() int32`

GetStatus302 returns the Status302 field if non-nil, zero value otherwise.

### GetStatus302Ok

`func (o *RealtimeMeasurements) GetStatus302Ok() (*int32, bool)`

GetStatus302Ok returns a tuple with the Status302 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus302

`func (o *RealtimeMeasurements) SetStatus302(v int32)`

SetStatus302 sets Status302 field to given value.

### HasStatus302

`func (o *RealtimeMeasurements) HasStatus302() bool`

HasStatus302 returns a boolean if a field has been set.

### GetStatus304

`func (o *RealtimeMeasurements) GetStatus304() int32`

GetStatus304 returns the Status304 field if non-nil, zero value otherwise.

### GetStatus304Ok

`func (o *RealtimeMeasurements) GetStatus304Ok() (*int32, bool)`

GetStatus304Ok returns a tuple with the Status304 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus304

`func (o *RealtimeMeasurements) SetStatus304(v int32)`

SetStatus304 sets Status304 field to given value.

### HasStatus304

`func (o *RealtimeMeasurements) HasStatus304() bool`

HasStatus304 returns a boolean if a field has been set.

### GetStatus400

`func (o *RealtimeMeasurements) GetStatus400() int32`

GetStatus400 returns the Status400 field if non-nil, zero value otherwise.

### GetStatus400Ok

`func (o *RealtimeMeasurements) GetStatus400Ok() (*int32, bool)`

GetStatus400Ok returns a tuple with the Status400 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus400

`func (o *RealtimeMeasurements) SetStatus400(v int32)`

SetStatus400 sets Status400 field to given value.

### HasStatus400

`func (o *RealtimeMeasurements) HasStatus400() bool`

HasStatus400 returns a boolean if a field has been set.

### GetStatus401

`func (o *RealtimeMeasurements) GetStatus401() int32`

GetStatus401 returns the Status401 field if non-nil, zero value otherwise.

### GetStatus401Ok

`func (o *RealtimeMeasurements) GetStatus401Ok() (*int32, bool)`

GetStatus401Ok returns a tuple with the Status401 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus401

`func (o *RealtimeMeasurements) SetStatus401(v int32)`

SetStatus401 sets Status401 field to given value.

### HasStatus401

`func (o *RealtimeMeasurements) HasStatus401() bool`

HasStatus401 returns a boolean if a field has been set.

### GetStatus403

`func (o *RealtimeMeasurements) GetStatus403() int32`

GetStatus403 returns the Status403 field if non-nil, zero value otherwise.

### GetStatus403Ok

`func (o *RealtimeMeasurements) GetStatus403Ok() (*int32, bool)`

GetStatus403Ok returns a tuple with the Status403 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus403

`func (o *RealtimeMeasurements) SetStatus403(v int32)`

SetStatus403 sets Status403 field to given value.

### HasStatus403

`func (o *RealtimeMeasurements) HasStatus403() bool`

HasStatus403 returns a boolean if a field has been set.

### GetStatus404

`func (o *RealtimeMeasurements) GetStatus404() int32`

GetStatus404 returns the Status404 field if non-nil, zero value otherwise.

### GetStatus404Ok

`func (o *RealtimeMeasurements) GetStatus404Ok() (*int32, bool)`

GetStatus404Ok returns a tuple with the Status404 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus404

`func (o *RealtimeMeasurements) SetStatus404(v int32)`

SetStatus404 sets Status404 field to given value.

### HasStatus404

`func (o *RealtimeMeasurements) HasStatus404() bool`

HasStatus404 returns a boolean if a field has been set.

### GetStatus406

`func (o *RealtimeMeasurements) GetStatus406() int32`

GetStatus406 returns the Status406 field if non-nil, zero value otherwise.

### GetStatus406Ok

`func (o *RealtimeMeasurements) GetStatus406Ok() (*int32, bool)`

GetStatus406Ok returns a tuple with the Status406 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus406

`func (o *RealtimeMeasurements) SetStatus406(v int32)`

SetStatus406 sets Status406 field to given value.

### HasStatus406

`func (o *RealtimeMeasurements) HasStatus406() bool`

HasStatus406 returns a boolean if a field has been set.

### GetStatus416

`func (o *RealtimeMeasurements) GetStatus416() int32`

GetStatus416 returns the Status416 field if non-nil, zero value otherwise.

### GetStatus416Ok

`func (o *RealtimeMeasurements) GetStatus416Ok() (*int32, bool)`

GetStatus416Ok returns a tuple with the Status416 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus416

`func (o *RealtimeMeasurements) SetStatus416(v int32)`

SetStatus416 sets Status416 field to given value.

### HasStatus416

`func (o *RealtimeMeasurements) HasStatus416() bool`

HasStatus416 returns a boolean if a field has been set.

### GetStatus429

`func (o *RealtimeMeasurements) GetStatus429() int32`

GetStatus429 returns the Status429 field if non-nil, zero value otherwise.

### GetStatus429Ok

`func (o *RealtimeMeasurements) GetStatus429Ok() (*int32, bool)`

GetStatus429Ok returns a tuple with the Status429 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus429

`func (o *RealtimeMeasurements) SetStatus429(v int32)`

SetStatus429 sets Status429 field to given value.

### HasStatus429

`func (o *RealtimeMeasurements) HasStatus429() bool`

HasStatus429 returns a boolean if a field has been set.

### GetStatus500

`func (o *RealtimeMeasurements) GetStatus500() int32`

GetStatus500 returns the Status500 field if non-nil, zero value otherwise.

### GetStatus500Ok

`func (o *RealtimeMeasurements) GetStatus500Ok() (*int32, bool)`

GetStatus500Ok returns a tuple with the Status500 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus500

`func (o *RealtimeMeasurements) SetStatus500(v int32)`

SetStatus500 sets Status500 field to given value.

### HasStatus500

`func (o *RealtimeMeasurements) HasStatus500() bool`

HasStatus500 returns a boolean if a field has been set.

### GetStatus501

`func (o *RealtimeMeasurements) GetStatus501() int32`

GetStatus501 returns the Status501 field if non-nil, zero value otherwise.

### GetStatus501Ok

`func (o *RealtimeMeasurements) GetStatus501Ok() (*int32, bool)`

GetStatus501Ok returns a tuple with the Status501 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus501

`func (o *RealtimeMeasurements) SetStatus501(v int32)`

SetStatus501 sets Status501 field to given value.

### HasStatus501

`func (o *RealtimeMeasurements) HasStatus501() bool`

HasStatus501 returns a boolean if a field has been set.

### GetStatus502

`func (o *RealtimeMeasurements) GetStatus502() int32`

GetStatus502 returns the Status502 field if non-nil, zero value otherwise.

### GetStatus502Ok

`func (o *RealtimeMeasurements) GetStatus502Ok() (*int32, bool)`

GetStatus502Ok returns a tuple with the Status502 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus502

`func (o *RealtimeMeasurements) SetStatus502(v int32)`

SetStatus502 sets Status502 field to given value.

### HasStatus502

`func (o *RealtimeMeasurements) HasStatus502() bool`

HasStatus502 returns a boolean if a field has been set.

### GetStatus503

`func (o *RealtimeMeasurements) GetStatus503() int32`

GetStatus503 returns the Status503 field if non-nil, zero value otherwise.

### GetStatus503Ok

`func (o *RealtimeMeasurements) GetStatus503Ok() (*int32, bool)`

GetStatus503Ok returns a tuple with the Status503 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus503

`func (o *RealtimeMeasurements) SetStatus503(v int32)`

SetStatus503 sets Status503 field to given value.

### HasStatus503

`func (o *RealtimeMeasurements) HasStatus503() bool`

HasStatus503 returns a boolean if a field has been set.

### GetStatus504

`func (o *RealtimeMeasurements) GetStatus504() int32`

GetStatus504 returns the Status504 field if non-nil, zero value otherwise.

### GetStatus504Ok

`func (o *RealtimeMeasurements) GetStatus504Ok() (*int32, bool)`

GetStatus504Ok returns a tuple with the Status504 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus504

`func (o *RealtimeMeasurements) SetStatus504(v int32)`

SetStatus504 sets Status504 field to given value.

### HasStatus504

`func (o *RealtimeMeasurements) HasStatus504() bool`

HasStatus504 returns a boolean if a field has been set.

### GetStatus505

`func (o *RealtimeMeasurements) GetStatus505() int32`

GetStatus505 returns the Status505 field if non-nil, zero value otherwise.

### GetStatus505Ok

`func (o *RealtimeMeasurements) GetStatus505Ok() (*int32, bool)`

GetStatus505Ok returns a tuple with the Status505 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus505

`func (o *RealtimeMeasurements) SetStatus505(v int32)`

SetStatus505 sets Status505 field to given value.

### HasStatus505

`func (o *RealtimeMeasurements) HasStatus505() bool`

HasStatus505 returns a boolean if a field has been set.

### GetUncacheable

`func (o *RealtimeMeasurements) GetUncacheable() int32`

GetUncacheable returns the Uncacheable field if non-nil, zero value otherwise.

### GetUncacheableOk

`func (o *RealtimeMeasurements) GetUncacheableOk() (*int32, bool)`

GetUncacheableOk returns a tuple with the Uncacheable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUncacheable

`func (o *RealtimeMeasurements) SetUncacheable(v int32)`

SetUncacheable sets Uncacheable field to given value.

### HasUncacheable

`func (o *RealtimeMeasurements) HasUncacheable() bool`

HasUncacheable returns a boolean if a field has been set.

### GetPassTime

`func (o *RealtimeMeasurements) GetPassTime() float32`

GetPassTime returns the PassTime field if non-nil, zero value otherwise.

### GetPassTimeOk

`func (o *RealtimeMeasurements) GetPassTimeOk() (*float32, bool)`

GetPassTimeOk returns a tuple with the PassTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassTime

`func (o *RealtimeMeasurements) SetPassTime(v float32)`

SetPassTime sets PassTime field to given value.

### HasPassTime

`func (o *RealtimeMeasurements) HasPassTime() bool`

HasPassTime returns a boolean if a field has been set.

### GetTLS

`func (o *RealtimeMeasurements) GetTLS() int32`

GetTLS returns the TLS field if non-nil, zero value otherwise.

### GetTLSOk

`func (o *RealtimeMeasurements) GetTLSOk() (*int32, bool)`

GetTLSOk returns a tuple with the TLS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTLS

`func (o *RealtimeMeasurements) SetTLS(v int32)`

SetTLS sets TLS field to given value.

### HasTLS

`func (o *RealtimeMeasurements) HasTLS() bool`

HasTLS returns a boolean if a field has been set.

### GetTLSV10

`func (o *RealtimeMeasurements) GetTLSV10() int32`

GetTLSV10 returns the TLSV10 field if non-nil, zero value otherwise.

### GetTLSV10Ok

`func (o *RealtimeMeasurements) GetTLSV10Ok() (*int32, bool)`

GetTLSV10Ok returns a tuple with the TLSV10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTLSV10

`func (o *RealtimeMeasurements) SetTLSV10(v int32)`

SetTLSV10 sets TLSV10 field to given value.

### HasTLSV10

`func (o *RealtimeMeasurements) HasTLSV10() bool`

HasTLSV10 returns a boolean if a field has been set.

### GetTLSV11

`func (o *RealtimeMeasurements) GetTLSV11() int32`

GetTLSV11 returns the TLSV11 field if non-nil, zero value otherwise.

### GetTLSV11Ok

`func (o *RealtimeMeasurements) GetTLSV11Ok() (*int32, bool)`

GetTLSV11Ok returns a tuple with the TLSV11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTLSV11

`func (o *RealtimeMeasurements) SetTLSV11(v int32)`

SetTLSV11 sets TLSV11 field to given value.

### HasTLSV11

`func (o *RealtimeMeasurements) HasTLSV11() bool`

HasTLSV11 returns a boolean if a field has been set.

### GetTLSV12

`func (o *RealtimeMeasurements) GetTLSV12() int32`

GetTLSV12 returns the TLSV12 field if non-nil, zero value otherwise.

### GetTLSV12Ok

`func (o *RealtimeMeasurements) GetTLSV12Ok() (*int32, bool)`

GetTLSV12Ok returns a tuple with the TLSV12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTLSV12

`func (o *RealtimeMeasurements) SetTLSV12(v int32)`

SetTLSV12 sets TLSV12 field to given value.

### HasTLSV12

`func (o *RealtimeMeasurements) HasTLSV12() bool`

HasTLSV12 returns a boolean if a field has been set.

### GetTLSV13

`func (o *RealtimeMeasurements) GetTLSV13() int32`

GetTLSV13 returns the TLSV13 field if non-nil, zero value otherwise.

### GetTLSV13Ok

`func (o *RealtimeMeasurements) GetTLSV13Ok() (*int32, bool)`

GetTLSV13Ok returns a tuple with the TLSV13 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTLSV13

`func (o *RealtimeMeasurements) SetTLSV13(v int32)`

SetTLSV13 sets TLSV13 field to given value.

### HasTLSV13

`func (o *RealtimeMeasurements) HasTLSV13() bool`

HasTLSV13 returns a boolean if a field has been set.

### GetObjectSize1k

`func (o *RealtimeMeasurements) GetObjectSize1k() int32`

GetObjectSize1k returns the ObjectSize1k field if non-nil, zero value otherwise.

### GetObjectSize1kOk

`func (o *RealtimeMeasurements) GetObjectSize1kOk() (*int32, bool)`

GetObjectSize1kOk returns a tuple with the ObjectSize1k field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectSize1k

`func (o *RealtimeMeasurements) SetObjectSize1k(v int32)`

SetObjectSize1k sets ObjectSize1k field to given value.

### HasObjectSize1k

`func (o *RealtimeMeasurements) HasObjectSize1k() bool`

HasObjectSize1k returns a boolean if a field has been set.

### GetObjectSize10k

`func (o *RealtimeMeasurements) GetObjectSize10k() int32`

GetObjectSize10k returns the ObjectSize10k field if non-nil, zero value otherwise.

### GetObjectSize10kOk

`func (o *RealtimeMeasurements) GetObjectSize10kOk() (*int32, bool)`

GetObjectSize10kOk returns a tuple with the ObjectSize10k field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectSize10k

`func (o *RealtimeMeasurements) SetObjectSize10k(v int32)`

SetObjectSize10k sets ObjectSize10k field to given value.

### HasObjectSize10k

`func (o *RealtimeMeasurements) HasObjectSize10k() bool`

HasObjectSize10k returns a boolean if a field has been set.

### GetObjectSize100k

`func (o *RealtimeMeasurements) GetObjectSize100k() int32`

GetObjectSize100k returns the ObjectSize100k field if non-nil, zero value otherwise.

### GetObjectSize100kOk

`func (o *RealtimeMeasurements) GetObjectSize100kOk() (*int32, bool)`

GetObjectSize100kOk returns a tuple with the ObjectSize100k field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectSize100k

`func (o *RealtimeMeasurements) SetObjectSize100k(v int32)`

SetObjectSize100k sets ObjectSize100k field to given value.

### HasObjectSize100k

`func (o *RealtimeMeasurements) HasObjectSize100k() bool`

HasObjectSize100k returns a boolean if a field has been set.

### GetObjectSize1m

`func (o *RealtimeMeasurements) GetObjectSize1m() int32`

GetObjectSize1m returns the ObjectSize1m field if non-nil, zero value otherwise.

### GetObjectSize1mOk

`func (o *RealtimeMeasurements) GetObjectSize1mOk() (*int32, bool)`

GetObjectSize1mOk returns a tuple with the ObjectSize1m field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectSize1m

`func (o *RealtimeMeasurements) SetObjectSize1m(v int32)`

SetObjectSize1m sets ObjectSize1m field to given value.

### HasObjectSize1m

`func (o *RealtimeMeasurements) HasObjectSize1m() bool`

HasObjectSize1m returns a boolean if a field has been set.

### GetObjectSize10m

`func (o *RealtimeMeasurements) GetObjectSize10m() int32`

GetObjectSize10m returns the ObjectSize10m field if non-nil, zero value otherwise.

### GetObjectSize10mOk

`func (o *RealtimeMeasurements) GetObjectSize10mOk() (*int32, bool)`

GetObjectSize10mOk returns a tuple with the ObjectSize10m field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectSize10m

`func (o *RealtimeMeasurements) SetObjectSize10m(v int32)`

SetObjectSize10m sets ObjectSize10m field to given value.

### HasObjectSize10m

`func (o *RealtimeMeasurements) HasObjectSize10m() bool`

HasObjectSize10m returns a boolean if a field has been set.

### GetObjectSize100m

`func (o *RealtimeMeasurements) GetObjectSize100m() int32`

GetObjectSize100m returns the ObjectSize100m field if non-nil, zero value otherwise.

### GetObjectSize100mOk

`func (o *RealtimeMeasurements) GetObjectSize100mOk() (*int32, bool)`

GetObjectSize100mOk returns a tuple with the ObjectSize100m field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectSize100m

`func (o *RealtimeMeasurements) SetObjectSize100m(v int32)`

SetObjectSize100m sets ObjectSize100m field to given value.

### HasObjectSize100m

`func (o *RealtimeMeasurements) HasObjectSize100m() bool`

HasObjectSize100m returns a boolean if a field has been set.

### GetObjectSize1g

`func (o *RealtimeMeasurements) GetObjectSize1g() int32`

GetObjectSize1g returns the ObjectSize1g field if non-nil, zero value otherwise.

### GetObjectSize1gOk

`func (o *RealtimeMeasurements) GetObjectSize1gOk() (*int32, bool)`

GetObjectSize1gOk returns a tuple with the ObjectSize1g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectSize1g

`func (o *RealtimeMeasurements) SetObjectSize1g(v int32)`

SetObjectSize1g sets ObjectSize1g field to given value.

### HasObjectSize1g

`func (o *RealtimeMeasurements) HasObjectSize1g() bool`

HasObjectSize1g returns a boolean if a field has been set.

### GetObjectSizeOther

`func (o *RealtimeMeasurements) GetObjectSizeOther() int32`

GetObjectSizeOther returns the ObjectSizeOther field if non-nil, zero value otherwise.

### GetObjectSizeOtherOk

`func (o *RealtimeMeasurements) GetObjectSizeOtherOk() (*int32, bool)`

GetObjectSizeOtherOk returns a tuple with the ObjectSizeOther field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectSizeOther

`func (o *RealtimeMeasurements) SetObjectSizeOther(v int32)`

SetObjectSizeOther sets ObjectSizeOther field to given value.

### HasObjectSizeOther

`func (o *RealtimeMeasurements) HasObjectSizeOther() bool`

HasObjectSizeOther returns a boolean if a field has been set.

### GetRecvSubTime

`func (o *RealtimeMeasurements) GetRecvSubTime() float32`

GetRecvSubTime returns the RecvSubTime field if non-nil, zero value otherwise.

### GetRecvSubTimeOk

`func (o *RealtimeMeasurements) GetRecvSubTimeOk() (*float32, bool)`

GetRecvSubTimeOk returns a tuple with the RecvSubTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecvSubTime

`func (o *RealtimeMeasurements) SetRecvSubTime(v float32)`

SetRecvSubTime sets RecvSubTime field to given value.

### HasRecvSubTime

`func (o *RealtimeMeasurements) HasRecvSubTime() bool`

HasRecvSubTime returns a boolean if a field has been set.

### GetRecvSubCount

`func (o *RealtimeMeasurements) GetRecvSubCount() int32`

GetRecvSubCount returns the RecvSubCount field if non-nil, zero value otherwise.

### GetRecvSubCountOk

`func (o *RealtimeMeasurements) GetRecvSubCountOk() (*int32, bool)`

GetRecvSubCountOk returns a tuple with the RecvSubCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecvSubCount

`func (o *RealtimeMeasurements) SetRecvSubCount(v int32)`

SetRecvSubCount sets RecvSubCount field to given value.

### HasRecvSubCount

`func (o *RealtimeMeasurements) HasRecvSubCount() bool`

HasRecvSubCount returns a boolean if a field has been set.

### GetHashSubTime

`func (o *RealtimeMeasurements) GetHashSubTime() float32`

GetHashSubTime returns the HashSubTime field if non-nil, zero value otherwise.

### GetHashSubTimeOk

`func (o *RealtimeMeasurements) GetHashSubTimeOk() (*float32, bool)`

GetHashSubTimeOk returns a tuple with the HashSubTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashSubTime

`func (o *RealtimeMeasurements) SetHashSubTime(v float32)`

SetHashSubTime sets HashSubTime field to given value.

### HasHashSubTime

`func (o *RealtimeMeasurements) HasHashSubTime() bool`

HasHashSubTime returns a boolean if a field has been set.

### GetHashSubCount

`func (o *RealtimeMeasurements) GetHashSubCount() int32`

GetHashSubCount returns the HashSubCount field if non-nil, zero value otherwise.

### GetHashSubCountOk

`func (o *RealtimeMeasurements) GetHashSubCountOk() (*int32, bool)`

GetHashSubCountOk returns a tuple with the HashSubCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashSubCount

`func (o *RealtimeMeasurements) SetHashSubCount(v int32)`

SetHashSubCount sets HashSubCount field to given value.

### HasHashSubCount

`func (o *RealtimeMeasurements) HasHashSubCount() bool`

HasHashSubCount returns a boolean if a field has been set.

### GetMissSubTime

`func (o *RealtimeMeasurements) GetMissSubTime() float32`

GetMissSubTime returns the MissSubTime field if non-nil, zero value otherwise.

### GetMissSubTimeOk

`func (o *RealtimeMeasurements) GetMissSubTimeOk() (*float32, bool)`

GetMissSubTimeOk returns a tuple with the MissSubTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissSubTime

`func (o *RealtimeMeasurements) SetMissSubTime(v float32)`

SetMissSubTime sets MissSubTime field to given value.

### HasMissSubTime

`func (o *RealtimeMeasurements) HasMissSubTime() bool`

HasMissSubTime returns a boolean if a field has been set.

### GetMissSubCount

`func (o *RealtimeMeasurements) GetMissSubCount() int32`

GetMissSubCount returns the MissSubCount field if non-nil, zero value otherwise.

### GetMissSubCountOk

`func (o *RealtimeMeasurements) GetMissSubCountOk() (*int32, bool)`

GetMissSubCountOk returns a tuple with the MissSubCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissSubCount

`func (o *RealtimeMeasurements) SetMissSubCount(v int32)`

SetMissSubCount sets MissSubCount field to given value.

### HasMissSubCount

`func (o *RealtimeMeasurements) HasMissSubCount() bool`

HasMissSubCount returns a boolean if a field has been set.

### GetFetchSubTime

`func (o *RealtimeMeasurements) GetFetchSubTime() float32`

GetFetchSubTime returns the FetchSubTime field if non-nil, zero value otherwise.

### GetFetchSubTimeOk

`func (o *RealtimeMeasurements) GetFetchSubTimeOk() (*float32, bool)`

GetFetchSubTimeOk returns a tuple with the FetchSubTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFetchSubTime

`func (o *RealtimeMeasurements) SetFetchSubTime(v float32)`

SetFetchSubTime sets FetchSubTime field to given value.

### HasFetchSubTime

`func (o *RealtimeMeasurements) HasFetchSubTime() bool`

HasFetchSubTime returns a boolean if a field has been set.

### GetFetchSubCount

`func (o *RealtimeMeasurements) GetFetchSubCount() int32`

GetFetchSubCount returns the FetchSubCount field if non-nil, zero value otherwise.

### GetFetchSubCountOk

`func (o *RealtimeMeasurements) GetFetchSubCountOk() (*int32, bool)`

GetFetchSubCountOk returns a tuple with the FetchSubCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFetchSubCount

`func (o *RealtimeMeasurements) SetFetchSubCount(v int32)`

SetFetchSubCount sets FetchSubCount field to given value.

### HasFetchSubCount

`func (o *RealtimeMeasurements) HasFetchSubCount() bool`

HasFetchSubCount returns a boolean if a field has been set.

### GetPassSubTime

`func (o *RealtimeMeasurements) GetPassSubTime() float32`

GetPassSubTime returns the PassSubTime field if non-nil, zero value otherwise.

### GetPassSubTimeOk

`func (o *RealtimeMeasurements) GetPassSubTimeOk() (*float32, bool)`

GetPassSubTimeOk returns a tuple with the PassSubTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassSubTime

`func (o *RealtimeMeasurements) SetPassSubTime(v float32)`

SetPassSubTime sets PassSubTime field to given value.

### HasPassSubTime

`func (o *RealtimeMeasurements) HasPassSubTime() bool`

HasPassSubTime returns a boolean if a field has been set.

### GetPassSubCount

`func (o *RealtimeMeasurements) GetPassSubCount() int32`

GetPassSubCount returns the PassSubCount field if non-nil, zero value otherwise.

### GetPassSubCountOk

`func (o *RealtimeMeasurements) GetPassSubCountOk() (*int32, bool)`

GetPassSubCountOk returns a tuple with the PassSubCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassSubCount

`func (o *RealtimeMeasurements) SetPassSubCount(v int32)`

SetPassSubCount sets PassSubCount field to given value.

### HasPassSubCount

`func (o *RealtimeMeasurements) HasPassSubCount() bool`

HasPassSubCount returns a boolean if a field has been set.

### GetPipeSubTime

`func (o *RealtimeMeasurements) GetPipeSubTime() float32`

GetPipeSubTime returns the PipeSubTime field if non-nil, zero value otherwise.

### GetPipeSubTimeOk

`func (o *RealtimeMeasurements) GetPipeSubTimeOk() (*float32, bool)`

GetPipeSubTimeOk returns a tuple with the PipeSubTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipeSubTime

`func (o *RealtimeMeasurements) SetPipeSubTime(v float32)`

SetPipeSubTime sets PipeSubTime field to given value.

### HasPipeSubTime

`func (o *RealtimeMeasurements) HasPipeSubTime() bool`

HasPipeSubTime returns a boolean if a field has been set.

### GetPipeSubCount

`func (o *RealtimeMeasurements) GetPipeSubCount() int32`

GetPipeSubCount returns the PipeSubCount field if non-nil, zero value otherwise.

### GetPipeSubCountOk

`func (o *RealtimeMeasurements) GetPipeSubCountOk() (*int32, bool)`

GetPipeSubCountOk returns a tuple with the PipeSubCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipeSubCount

`func (o *RealtimeMeasurements) SetPipeSubCount(v int32)`

SetPipeSubCount sets PipeSubCount field to given value.

### HasPipeSubCount

`func (o *RealtimeMeasurements) HasPipeSubCount() bool`

HasPipeSubCount returns a boolean if a field has been set.

### GetDeliverSubTime

`func (o *RealtimeMeasurements) GetDeliverSubTime() float32`

GetDeliverSubTime returns the DeliverSubTime field if non-nil, zero value otherwise.

### GetDeliverSubTimeOk

`func (o *RealtimeMeasurements) GetDeliverSubTimeOk() (*float32, bool)`

GetDeliverSubTimeOk returns a tuple with the DeliverSubTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliverSubTime

`func (o *RealtimeMeasurements) SetDeliverSubTime(v float32)`

SetDeliverSubTime sets DeliverSubTime field to given value.

### HasDeliverSubTime

`func (o *RealtimeMeasurements) HasDeliverSubTime() bool`

HasDeliverSubTime returns a boolean if a field has been set.

### GetDeliverSubCount

`func (o *RealtimeMeasurements) GetDeliverSubCount() int32`

GetDeliverSubCount returns the DeliverSubCount field if non-nil, zero value otherwise.

### GetDeliverSubCountOk

`func (o *RealtimeMeasurements) GetDeliverSubCountOk() (*int32, bool)`

GetDeliverSubCountOk returns a tuple with the DeliverSubCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliverSubCount

`func (o *RealtimeMeasurements) SetDeliverSubCount(v int32)`

SetDeliverSubCount sets DeliverSubCount field to given value.

### HasDeliverSubCount

`func (o *RealtimeMeasurements) HasDeliverSubCount() bool`

HasDeliverSubCount returns a boolean if a field has been set.

### GetErrorSubTime

`func (o *RealtimeMeasurements) GetErrorSubTime() float32`

GetErrorSubTime returns the ErrorSubTime field if non-nil, zero value otherwise.

### GetErrorSubTimeOk

`func (o *RealtimeMeasurements) GetErrorSubTimeOk() (*float32, bool)`

GetErrorSubTimeOk returns a tuple with the ErrorSubTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorSubTime

`func (o *RealtimeMeasurements) SetErrorSubTime(v float32)`

SetErrorSubTime sets ErrorSubTime field to given value.

### HasErrorSubTime

`func (o *RealtimeMeasurements) HasErrorSubTime() bool`

HasErrorSubTime returns a boolean if a field has been set.

### GetErrorSubCount

`func (o *RealtimeMeasurements) GetErrorSubCount() int32`

GetErrorSubCount returns the ErrorSubCount field if non-nil, zero value otherwise.

### GetErrorSubCountOk

`func (o *RealtimeMeasurements) GetErrorSubCountOk() (*int32, bool)`

GetErrorSubCountOk returns a tuple with the ErrorSubCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorSubCount

`func (o *RealtimeMeasurements) SetErrorSubCount(v int32)`

SetErrorSubCount sets ErrorSubCount field to given value.

### HasErrorSubCount

`func (o *RealtimeMeasurements) HasErrorSubCount() bool`

HasErrorSubCount returns a boolean if a field has been set.

### GetHitSubTime

`func (o *RealtimeMeasurements) GetHitSubTime() float32`

GetHitSubTime returns the HitSubTime field if non-nil, zero value otherwise.

### GetHitSubTimeOk

`func (o *RealtimeMeasurements) GetHitSubTimeOk() (*float32, bool)`

GetHitSubTimeOk returns a tuple with the HitSubTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHitSubTime

`func (o *RealtimeMeasurements) SetHitSubTime(v float32)`

SetHitSubTime sets HitSubTime field to given value.

### HasHitSubTime

`func (o *RealtimeMeasurements) HasHitSubTime() bool`

HasHitSubTime returns a boolean if a field has been set.

### GetHitSubCount

`func (o *RealtimeMeasurements) GetHitSubCount() int32`

GetHitSubCount returns the HitSubCount field if non-nil, zero value otherwise.

### GetHitSubCountOk

`func (o *RealtimeMeasurements) GetHitSubCountOk() (*int32, bool)`

GetHitSubCountOk returns a tuple with the HitSubCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHitSubCount

`func (o *RealtimeMeasurements) SetHitSubCount(v int32)`

SetHitSubCount sets HitSubCount field to given value.

### HasHitSubCount

`func (o *RealtimeMeasurements) HasHitSubCount() bool`

HasHitSubCount returns a boolean if a field has been set.

### GetPrehashSubTime

`func (o *RealtimeMeasurements) GetPrehashSubTime() float32`

GetPrehashSubTime returns the PrehashSubTime field if non-nil, zero value otherwise.

### GetPrehashSubTimeOk

`func (o *RealtimeMeasurements) GetPrehashSubTimeOk() (*float32, bool)`

GetPrehashSubTimeOk returns a tuple with the PrehashSubTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrehashSubTime

`func (o *RealtimeMeasurements) SetPrehashSubTime(v float32)`

SetPrehashSubTime sets PrehashSubTime field to given value.

### HasPrehashSubTime

`func (o *RealtimeMeasurements) HasPrehashSubTime() bool`

HasPrehashSubTime returns a boolean if a field has been set.

### GetPrehashSubCount

`func (o *RealtimeMeasurements) GetPrehashSubCount() int32`

GetPrehashSubCount returns the PrehashSubCount field if non-nil, zero value otherwise.

### GetPrehashSubCountOk

`func (o *RealtimeMeasurements) GetPrehashSubCountOk() (*int32, bool)`

GetPrehashSubCountOk returns a tuple with the PrehashSubCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrehashSubCount

`func (o *RealtimeMeasurements) SetPrehashSubCount(v int32)`

SetPrehashSubCount sets PrehashSubCount field to given value.

### HasPrehashSubCount

`func (o *RealtimeMeasurements) HasPrehashSubCount() bool`

HasPrehashSubCount returns a boolean if a field has been set.

### GetPredeliverSubTime

`func (o *RealtimeMeasurements) GetPredeliverSubTime() float32`

GetPredeliverSubTime returns the PredeliverSubTime field if non-nil, zero value otherwise.

### GetPredeliverSubTimeOk

`func (o *RealtimeMeasurements) GetPredeliverSubTimeOk() (*float32, bool)`

GetPredeliverSubTimeOk returns a tuple with the PredeliverSubTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredeliverSubTime

`func (o *RealtimeMeasurements) SetPredeliverSubTime(v float32)`

SetPredeliverSubTime sets PredeliverSubTime field to given value.

### HasPredeliverSubTime

`func (o *RealtimeMeasurements) HasPredeliverSubTime() bool`

HasPredeliverSubTime returns a boolean if a field has been set.

### GetPredeliverSubCount

`func (o *RealtimeMeasurements) GetPredeliverSubCount() int32`

GetPredeliverSubCount returns the PredeliverSubCount field if non-nil, zero value otherwise.

### GetPredeliverSubCountOk

`func (o *RealtimeMeasurements) GetPredeliverSubCountOk() (*int32, bool)`

GetPredeliverSubCountOk returns a tuple with the PredeliverSubCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredeliverSubCount

`func (o *RealtimeMeasurements) SetPredeliverSubCount(v int32)`

SetPredeliverSubCount sets PredeliverSubCount field to given value.

### HasPredeliverSubCount

`func (o *RealtimeMeasurements) HasPredeliverSubCount() bool`

HasPredeliverSubCount returns a boolean if a field has been set.

### GetHitRespBodyBytes

`func (o *RealtimeMeasurements) GetHitRespBodyBytes() int32`

GetHitRespBodyBytes returns the HitRespBodyBytes field if non-nil, zero value otherwise.

### GetHitRespBodyBytesOk

`func (o *RealtimeMeasurements) GetHitRespBodyBytesOk() (*int32, bool)`

GetHitRespBodyBytesOk returns a tuple with the HitRespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHitRespBodyBytes

`func (o *RealtimeMeasurements) SetHitRespBodyBytes(v int32)`

SetHitRespBodyBytes sets HitRespBodyBytes field to given value.

### HasHitRespBodyBytes

`func (o *RealtimeMeasurements) HasHitRespBodyBytes() bool`

HasHitRespBodyBytes returns a boolean if a field has been set.

### GetMissRespBodyBytes

`func (o *RealtimeMeasurements) GetMissRespBodyBytes() int32`

GetMissRespBodyBytes returns the MissRespBodyBytes field if non-nil, zero value otherwise.

### GetMissRespBodyBytesOk

`func (o *RealtimeMeasurements) GetMissRespBodyBytesOk() (*int32, bool)`

GetMissRespBodyBytesOk returns a tuple with the MissRespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissRespBodyBytes

`func (o *RealtimeMeasurements) SetMissRespBodyBytes(v int32)`

SetMissRespBodyBytes sets MissRespBodyBytes field to given value.

### HasMissRespBodyBytes

`func (o *RealtimeMeasurements) HasMissRespBodyBytes() bool`

HasMissRespBodyBytes returns a boolean if a field has been set.

### GetPassRespBodyBytes

`func (o *RealtimeMeasurements) GetPassRespBodyBytes() int32`

GetPassRespBodyBytes returns the PassRespBodyBytes field if non-nil, zero value otherwise.

### GetPassRespBodyBytesOk

`func (o *RealtimeMeasurements) GetPassRespBodyBytesOk() (*int32, bool)`

GetPassRespBodyBytesOk returns a tuple with the PassRespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassRespBodyBytes

`func (o *RealtimeMeasurements) SetPassRespBodyBytes(v int32)`

SetPassRespBodyBytes sets PassRespBodyBytes field to given value.

### HasPassRespBodyBytes

`func (o *RealtimeMeasurements) HasPassRespBodyBytes() bool`

HasPassRespBodyBytes returns a boolean if a field has been set.

### GetComputeReqHeaderBytes

`func (o *RealtimeMeasurements) GetComputeReqHeaderBytes() int32`

GetComputeReqHeaderBytes returns the ComputeReqHeaderBytes field if non-nil, zero value otherwise.

### GetComputeReqHeaderBytesOk

`func (o *RealtimeMeasurements) GetComputeReqHeaderBytesOk() (*int32, bool)`

GetComputeReqHeaderBytesOk returns a tuple with the ComputeReqHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeReqHeaderBytes

`func (o *RealtimeMeasurements) SetComputeReqHeaderBytes(v int32)`

SetComputeReqHeaderBytes sets ComputeReqHeaderBytes field to given value.

### HasComputeReqHeaderBytes

`func (o *RealtimeMeasurements) HasComputeReqHeaderBytes() bool`

HasComputeReqHeaderBytes returns a boolean if a field has been set.

### GetComputeReqBodyBytes

`func (o *RealtimeMeasurements) GetComputeReqBodyBytes() int32`

GetComputeReqBodyBytes returns the ComputeReqBodyBytes field if non-nil, zero value otherwise.

### GetComputeReqBodyBytesOk

`func (o *RealtimeMeasurements) GetComputeReqBodyBytesOk() (*int32, bool)`

GetComputeReqBodyBytesOk returns a tuple with the ComputeReqBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeReqBodyBytes

`func (o *RealtimeMeasurements) SetComputeReqBodyBytes(v int32)`

SetComputeReqBodyBytes sets ComputeReqBodyBytes field to given value.

### HasComputeReqBodyBytes

`func (o *RealtimeMeasurements) HasComputeReqBodyBytes() bool`

HasComputeReqBodyBytes returns a boolean if a field has been set.

### GetComputeRespHeaderBytes

`func (o *RealtimeMeasurements) GetComputeRespHeaderBytes() int32`

GetComputeRespHeaderBytes returns the ComputeRespHeaderBytes field if non-nil, zero value otherwise.

### GetComputeRespHeaderBytesOk

`func (o *RealtimeMeasurements) GetComputeRespHeaderBytesOk() (*int32, bool)`

GetComputeRespHeaderBytesOk returns a tuple with the ComputeRespHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespHeaderBytes

`func (o *RealtimeMeasurements) SetComputeRespHeaderBytes(v int32)`

SetComputeRespHeaderBytes sets ComputeRespHeaderBytes field to given value.

### HasComputeRespHeaderBytes

`func (o *RealtimeMeasurements) HasComputeRespHeaderBytes() bool`

HasComputeRespHeaderBytes returns a boolean if a field has been set.

### GetComputeRespBodyBytes

`func (o *RealtimeMeasurements) GetComputeRespBodyBytes() int32`

GetComputeRespBodyBytes returns the ComputeRespBodyBytes field if non-nil, zero value otherwise.

### GetComputeRespBodyBytesOk

`func (o *RealtimeMeasurements) GetComputeRespBodyBytesOk() (*int32, bool)`

GetComputeRespBodyBytesOk returns a tuple with the ComputeRespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespBodyBytes

`func (o *RealtimeMeasurements) SetComputeRespBodyBytes(v int32)`

SetComputeRespBodyBytes sets ComputeRespBodyBytes field to given value.

### HasComputeRespBodyBytes

`func (o *RealtimeMeasurements) HasComputeRespBodyBytes() bool`

HasComputeRespBodyBytes returns a boolean if a field has been set.

### GetImgvideo

`func (o *RealtimeMeasurements) GetImgvideo() int32`

GetImgvideo returns the Imgvideo field if non-nil, zero value otherwise.

### GetImgvideoOk

`func (o *RealtimeMeasurements) GetImgvideoOk() (*int32, bool)`

GetImgvideoOk returns a tuple with the Imgvideo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImgvideo

`func (o *RealtimeMeasurements) SetImgvideo(v int32)`

SetImgvideo sets Imgvideo field to given value.

### HasImgvideo

`func (o *RealtimeMeasurements) HasImgvideo() bool`

HasImgvideo returns a boolean if a field has been set.

### GetImgvideoFrames

`func (o *RealtimeMeasurements) GetImgvideoFrames() int32`

GetImgvideoFrames returns the ImgvideoFrames field if non-nil, zero value otherwise.

### GetImgvideoFramesOk

`func (o *RealtimeMeasurements) GetImgvideoFramesOk() (*int32, bool)`

GetImgvideoFramesOk returns a tuple with the ImgvideoFrames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImgvideoFrames

`func (o *RealtimeMeasurements) SetImgvideoFrames(v int32)`

SetImgvideoFrames sets ImgvideoFrames field to given value.

### HasImgvideoFrames

`func (o *RealtimeMeasurements) HasImgvideoFrames() bool`

HasImgvideoFrames returns a boolean if a field has been set.

### GetImgvideoRespHeaderBytes

`func (o *RealtimeMeasurements) GetImgvideoRespHeaderBytes() int32`

GetImgvideoRespHeaderBytes returns the ImgvideoRespHeaderBytes field if non-nil, zero value otherwise.

### GetImgvideoRespHeaderBytesOk

`func (o *RealtimeMeasurements) GetImgvideoRespHeaderBytesOk() (*int32, bool)`

GetImgvideoRespHeaderBytesOk returns a tuple with the ImgvideoRespHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImgvideoRespHeaderBytes

`func (o *RealtimeMeasurements) SetImgvideoRespHeaderBytes(v int32)`

SetImgvideoRespHeaderBytes sets ImgvideoRespHeaderBytes field to given value.

### HasImgvideoRespHeaderBytes

`func (o *RealtimeMeasurements) HasImgvideoRespHeaderBytes() bool`

HasImgvideoRespHeaderBytes returns a boolean if a field has been set.

### GetImgvideoRespBodyBytes

`func (o *RealtimeMeasurements) GetImgvideoRespBodyBytes() int32`

GetImgvideoRespBodyBytes returns the ImgvideoRespBodyBytes field if non-nil, zero value otherwise.

### GetImgvideoRespBodyBytesOk

`func (o *RealtimeMeasurements) GetImgvideoRespBodyBytesOk() (*int32, bool)`

GetImgvideoRespBodyBytesOk returns a tuple with the ImgvideoRespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImgvideoRespBodyBytes

`func (o *RealtimeMeasurements) SetImgvideoRespBodyBytes(v int32)`

SetImgvideoRespBodyBytes sets ImgvideoRespBodyBytes field to given value.

### HasImgvideoRespBodyBytes

`func (o *RealtimeMeasurements) HasImgvideoRespBodyBytes() bool`

HasImgvideoRespBodyBytes returns a boolean if a field has been set.

### GetImgvideoShield

`func (o *RealtimeMeasurements) GetImgvideoShield() int32`

GetImgvideoShield returns the ImgvideoShield field if non-nil, zero value otherwise.

### GetImgvideoShieldOk

`func (o *RealtimeMeasurements) GetImgvideoShieldOk() (*int32, bool)`

GetImgvideoShieldOk returns a tuple with the ImgvideoShield field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImgvideoShield

`func (o *RealtimeMeasurements) SetImgvideoShield(v int32)`

SetImgvideoShield sets ImgvideoShield field to given value.

### HasImgvideoShield

`func (o *RealtimeMeasurements) HasImgvideoShield() bool`

HasImgvideoShield returns a boolean if a field has been set.

### GetImgvideoShieldFrames

`func (o *RealtimeMeasurements) GetImgvideoShieldFrames() int32`

GetImgvideoShieldFrames returns the ImgvideoShieldFrames field if non-nil, zero value otherwise.

### GetImgvideoShieldFramesOk

`func (o *RealtimeMeasurements) GetImgvideoShieldFramesOk() (*int32, bool)`

GetImgvideoShieldFramesOk returns a tuple with the ImgvideoShieldFrames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImgvideoShieldFrames

`func (o *RealtimeMeasurements) SetImgvideoShieldFrames(v int32)`

SetImgvideoShieldFrames sets ImgvideoShieldFrames field to given value.

### HasImgvideoShieldFrames

`func (o *RealtimeMeasurements) HasImgvideoShieldFrames() bool`

HasImgvideoShieldFrames returns a boolean if a field has been set.

### GetImgvideoShieldRespHeaderBytes

`func (o *RealtimeMeasurements) GetImgvideoShieldRespHeaderBytes() int32`

GetImgvideoShieldRespHeaderBytes returns the ImgvideoShieldRespHeaderBytes field if non-nil, zero value otherwise.

### GetImgvideoShieldRespHeaderBytesOk

`func (o *RealtimeMeasurements) GetImgvideoShieldRespHeaderBytesOk() (*int32, bool)`

GetImgvideoShieldRespHeaderBytesOk returns a tuple with the ImgvideoShieldRespHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImgvideoShieldRespHeaderBytes

`func (o *RealtimeMeasurements) SetImgvideoShieldRespHeaderBytes(v int32)`

SetImgvideoShieldRespHeaderBytes sets ImgvideoShieldRespHeaderBytes field to given value.

### HasImgvideoShieldRespHeaderBytes

`func (o *RealtimeMeasurements) HasImgvideoShieldRespHeaderBytes() bool`

HasImgvideoShieldRespHeaderBytes returns a boolean if a field has been set.

### GetImgvideoShieldRespBodyBytes

`func (o *RealtimeMeasurements) GetImgvideoShieldRespBodyBytes() int32`

GetImgvideoShieldRespBodyBytes returns the ImgvideoShieldRespBodyBytes field if non-nil, zero value otherwise.

### GetImgvideoShieldRespBodyBytesOk

`func (o *RealtimeMeasurements) GetImgvideoShieldRespBodyBytesOk() (*int32, bool)`

GetImgvideoShieldRespBodyBytesOk returns a tuple with the ImgvideoShieldRespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImgvideoShieldRespBodyBytes

`func (o *RealtimeMeasurements) SetImgvideoShieldRespBodyBytes(v int32)`

SetImgvideoShieldRespBodyBytes sets ImgvideoShieldRespBodyBytes field to given value.

### HasImgvideoShieldRespBodyBytes

`func (o *RealtimeMeasurements) HasImgvideoShieldRespBodyBytes() bool`

HasImgvideoShieldRespBodyBytes returns a boolean if a field has been set.

### GetLogBytes

`func (o *RealtimeMeasurements) GetLogBytes() int32`

GetLogBytes returns the LogBytes field if non-nil, zero value otherwise.

### GetLogBytesOk

`func (o *RealtimeMeasurements) GetLogBytesOk() (*int32, bool)`

GetLogBytesOk returns a tuple with the LogBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogBytes

`func (o *RealtimeMeasurements) SetLogBytes(v int32)`

SetLogBytes sets LogBytes field to given value.

### HasLogBytes

`func (o *RealtimeMeasurements) HasLogBytes() bool`

HasLogBytes returns a boolean if a field has been set.

### GetEdgeRequests

`func (o *RealtimeMeasurements) GetEdgeRequests() int32`

GetEdgeRequests returns the EdgeRequests field if non-nil, zero value otherwise.

### GetEdgeRequestsOk

`func (o *RealtimeMeasurements) GetEdgeRequestsOk() (*int32, bool)`

GetEdgeRequestsOk returns a tuple with the EdgeRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeRequests

`func (o *RealtimeMeasurements) SetEdgeRequests(v int32)`

SetEdgeRequests sets EdgeRequests field to given value.

### HasEdgeRequests

`func (o *RealtimeMeasurements) HasEdgeRequests() bool`

HasEdgeRequests returns a boolean if a field has been set.

### GetEdgeRespHeaderBytes

`func (o *RealtimeMeasurements) GetEdgeRespHeaderBytes() int32`

GetEdgeRespHeaderBytes returns the EdgeRespHeaderBytes field if non-nil, zero value otherwise.

### GetEdgeRespHeaderBytesOk

`func (o *RealtimeMeasurements) GetEdgeRespHeaderBytesOk() (*int32, bool)`

GetEdgeRespHeaderBytesOk returns a tuple with the EdgeRespHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeRespHeaderBytes

`func (o *RealtimeMeasurements) SetEdgeRespHeaderBytes(v int32)`

SetEdgeRespHeaderBytes sets EdgeRespHeaderBytes field to given value.

### HasEdgeRespHeaderBytes

`func (o *RealtimeMeasurements) HasEdgeRespHeaderBytes() bool`

HasEdgeRespHeaderBytes returns a boolean if a field has been set.

### GetEdgeRespBodyBytes

`func (o *RealtimeMeasurements) GetEdgeRespBodyBytes() int32`

GetEdgeRespBodyBytes returns the EdgeRespBodyBytes field if non-nil, zero value otherwise.

### GetEdgeRespBodyBytesOk

`func (o *RealtimeMeasurements) GetEdgeRespBodyBytesOk() (*int32, bool)`

GetEdgeRespBodyBytesOk returns a tuple with the EdgeRespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeRespBodyBytes

`func (o *RealtimeMeasurements) SetEdgeRespBodyBytes(v int32)`

SetEdgeRespBodyBytes sets EdgeRespBodyBytes field to given value.

### HasEdgeRespBodyBytes

`func (o *RealtimeMeasurements) HasEdgeRespBodyBytes() bool`

HasEdgeRespBodyBytes returns a boolean if a field has been set.

### GetOriginRevalidations

`func (o *RealtimeMeasurements) GetOriginRevalidations() int32`

GetOriginRevalidations returns the OriginRevalidations field if non-nil, zero value otherwise.

### GetOriginRevalidationsOk

`func (o *RealtimeMeasurements) GetOriginRevalidationsOk() (*int32, bool)`

GetOriginRevalidationsOk returns a tuple with the OriginRevalidations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginRevalidations

`func (o *RealtimeMeasurements) SetOriginRevalidations(v int32)`

SetOriginRevalidations sets OriginRevalidations field to given value.

### HasOriginRevalidations

`func (o *RealtimeMeasurements) HasOriginRevalidations() bool`

HasOriginRevalidations returns a boolean if a field has been set.

### GetOriginFetches

`func (o *RealtimeMeasurements) GetOriginFetches() int32`

GetOriginFetches returns the OriginFetches field if non-nil, zero value otherwise.

### GetOriginFetchesOk

`func (o *RealtimeMeasurements) GetOriginFetchesOk() (*int32, bool)`

GetOriginFetchesOk returns a tuple with the OriginFetches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginFetches

`func (o *RealtimeMeasurements) SetOriginFetches(v int32)`

SetOriginFetches sets OriginFetches field to given value.

### HasOriginFetches

`func (o *RealtimeMeasurements) HasOriginFetches() bool`

HasOriginFetches returns a boolean if a field has been set.

### GetOriginFetchHeaderBytes

`func (o *RealtimeMeasurements) GetOriginFetchHeaderBytes() int32`

GetOriginFetchHeaderBytes returns the OriginFetchHeaderBytes field if non-nil, zero value otherwise.

### GetOriginFetchHeaderBytesOk

`func (o *RealtimeMeasurements) GetOriginFetchHeaderBytesOk() (*int32, bool)`

GetOriginFetchHeaderBytesOk returns a tuple with the OriginFetchHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginFetchHeaderBytes

`func (o *RealtimeMeasurements) SetOriginFetchHeaderBytes(v int32)`

SetOriginFetchHeaderBytes sets OriginFetchHeaderBytes field to given value.

### HasOriginFetchHeaderBytes

`func (o *RealtimeMeasurements) HasOriginFetchHeaderBytes() bool`

HasOriginFetchHeaderBytes returns a boolean if a field has been set.

### GetOriginFetchBodyBytes

`func (o *RealtimeMeasurements) GetOriginFetchBodyBytes() int32`

GetOriginFetchBodyBytes returns the OriginFetchBodyBytes field if non-nil, zero value otherwise.

### GetOriginFetchBodyBytesOk

`func (o *RealtimeMeasurements) GetOriginFetchBodyBytesOk() (*int32, bool)`

GetOriginFetchBodyBytesOk returns a tuple with the OriginFetchBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginFetchBodyBytes

`func (o *RealtimeMeasurements) SetOriginFetchBodyBytes(v int32)`

SetOriginFetchBodyBytes sets OriginFetchBodyBytes field to given value.

### HasOriginFetchBodyBytes

`func (o *RealtimeMeasurements) HasOriginFetchBodyBytes() bool`

HasOriginFetchBodyBytes returns a boolean if a field has been set.

### GetOriginFetchRespHeaderBytes

`func (o *RealtimeMeasurements) GetOriginFetchRespHeaderBytes() int32`

GetOriginFetchRespHeaderBytes returns the OriginFetchRespHeaderBytes field if non-nil, zero value otherwise.

### GetOriginFetchRespHeaderBytesOk

`func (o *RealtimeMeasurements) GetOriginFetchRespHeaderBytesOk() (*int32, bool)`

GetOriginFetchRespHeaderBytesOk returns a tuple with the OriginFetchRespHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginFetchRespHeaderBytes

`func (o *RealtimeMeasurements) SetOriginFetchRespHeaderBytes(v int32)`

SetOriginFetchRespHeaderBytes sets OriginFetchRespHeaderBytes field to given value.

### HasOriginFetchRespHeaderBytes

`func (o *RealtimeMeasurements) HasOriginFetchRespHeaderBytes() bool`

HasOriginFetchRespHeaderBytes returns a boolean if a field has been set.

### GetOriginFetchRespBodyBytes

`func (o *RealtimeMeasurements) GetOriginFetchRespBodyBytes() int32`

GetOriginFetchRespBodyBytes returns the OriginFetchRespBodyBytes field if non-nil, zero value otherwise.

### GetOriginFetchRespBodyBytesOk

`func (o *RealtimeMeasurements) GetOriginFetchRespBodyBytesOk() (*int32, bool)`

GetOriginFetchRespBodyBytesOk returns a tuple with the OriginFetchRespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginFetchRespBodyBytes

`func (o *RealtimeMeasurements) SetOriginFetchRespBodyBytes(v int32)`

SetOriginFetchRespBodyBytes sets OriginFetchRespBodyBytes field to given value.

### HasOriginFetchRespBodyBytes

`func (o *RealtimeMeasurements) HasOriginFetchRespBodyBytes() bool`

HasOriginFetchRespBodyBytes returns a boolean if a field has been set.

### GetShieldRevalidations

`func (o *RealtimeMeasurements) GetShieldRevalidations() int32`

GetShieldRevalidations returns the ShieldRevalidations field if non-nil, zero value otherwise.

### GetShieldRevalidationsOk

`func (o *RealtimeMeasurements) GetShieldRevalidationsOk() (*int32, bool)`

GetShieldRevalidationsOk returns a tuple with the ShieldRevalidations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShieldRevalidations

`func (o *RealtimeMeasurements) SetShieldRevalidations(v int32)`

SetShieldRevalidations sets ShieldRevalidations field to given value.

### HasShieldRevalidations

`func (o *RealtimeMeasurements) HasShieldRevalidations() bool`

HasShieldRevalidations returns a boolean if a field has been set.

### GetShieldFetches

`func (o *RealtimeMeasurements) GetShieldFetches() int32`

GetShieldFetches returns the ShieldFetches field if non-nil, zero value otherwise.

### GetShieldFetchesOk

`func (o *RealtimeMeasurements) GetShieldFetchesOk() (*int32, bool)`

GetShieldFetchesOk returns a tuple with the ShieldFetches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShieldFetches

`func (o *RealtimeMeasurements) SetShieldFetches(v int32)`

SetShieldFetches sets ShieldFetches field to given value.

### HasShieldFetches

`func (o *RealtimeMeasurements) HasShieldFetches() bool`

HasShieldFetches returns a boolean if a field has been set.

### GetShieldFetchHeaderBytes

`func (o *RealtimeMeasurements) GetShieldFetchHeaderBytes() int32`

GetShieldFetchHeaderBytes returns the ShieldFetchHeaderBytes field if non-nil, zero value otherwise.

### GetShieldFetchHeaderBytesOk

`func (o *RealtimeMeasurements) GetShieldFetchHeaderBytesOk() (*int32, bool)`

GetShieldFetchHeaderBytesOk returns a tuple with the ShieldFetchHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShieldFetchHeaderBytes

`func (o *RealtimeMeasurements) SetShieldFetchHeaderBytes(v int32)`

SetShieldFetchHeaderBytes sets ShieldFetchHeaderBytes field to given value.

### HasShieldFetchHeaderBytes

`func (o *RealtimeMeasurements) HasShieldFetchHeaderBytes() bool`

HasShieldFetchHeaderBytes returns a boolean if a field has been set.

### GetShieldFetchBodyBytes

`func (o *RealtimeMeasurements) GetShieldFetchBodyBytes() int32`

GetShieldFetchBodyBytes returns the ShieldFetchBodyBytes field if non-nil, zero value otherwise.

### GetShieldFetchBodyBytesOk

`func (o *RealtimeMeasurements) GetShieldFetchBodyBytesOk() (*int32, bool)`

GetShieldFetchBodyBytesOk returns a tuple with the ShieldFetchBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShieldFetchBodyBytes

`func (o *RealtimeMeasurements) SetShieldFetchBodyBytes(v int32)`

SetShieldFetchBodyBytes sets ShieldFetchBodyBytes field to given value.

### HasShieldFetchBodyBytes

`func (o *RealtimeMeasurements) HasShieldFetchBodyBytes() bool`

HasShieldFetchBodyBytes returns a boolean if a field has been set.

### GetShieldFetchRespHeaderBytes

`func (o *RealtimeMeasurements) GetShieldFetchRespHeaderBytes() int32`

GetShieldFetchRespHeaderBytes returns the ShieldFetchRespHeaderBytes field if non-nil, zero value otherwise.

### GetShieldFetchRespHeaderBytesOk

`func (o *RealtimeMeasurements) GetShieldFetchRespHeaderBytesOk() (*int32, bool)`

GetShieldFetchRespHeaderBytesOk returns a tuple with the ShieldFetchRespHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShieldFetchRespHeaderBytes

`func (o *RealtimeMeasurements) SetShieldFetchRespHeaderBytes(v int32)`

SetShieldFetchRespHeaderBytes sets ShieldFetchRespHeaderBytes field to given value.

### HasShieldFetchRespHeaderBytes

`func (o *RealtimeMeasurements) HasShieldFetchRespHeaderBytes() bool`

HasShieldFetchRespHeaderBytes returns a boolean if a field has been set.

### GetShieldFetchRespBodyBytes

`func (o *RealtimeMeasurements) GetShieldFetchRespBodyBytes() int32`

GetShieldFetchRespBodyBytes returns the ShieldFetchRespBodyBytes field if non-nil, zero value otherwise.

### GetShieldFetchRespBodyBytesOk

`func (o *RealtimeMeasurements) GetShieldFetchRespBodyBytesOk() (*int32, bool)`

GetShieldFetchRespBodyBytesOk returns a tuple with the ShieldFetchRespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShieldFetchRespBodyBytes

`func (o *RealtimeMeasurements) SetShieldFetchRespBodyBytes(v int32)`

SetShieldFetchRespBodyBytes sets ShieldFetchRespBodyBytes field to given value.

### HasShieldFetchRespBodyBytes

`func (o *RealtimeMeasurements) HasShieldFetchRespBodyBytes() bool`

HasShieldFetchRespBodyBytes returns a boolean if a field has been set.

### GetSegblockOriginFetches

`func (o *RealtimeMeasurements) GetSegblockOriginFetches() int32`

GetSegblockOriginFetches returns the SegblockOriginFetches field if non-nil, zero value otherwise.

### GetSegblockOriginFetchesOk

`func (o *RealtimeMeasurements) GetSegblockOriginFetchesOk() (*int32, bool)`

GetSegblockOriginFetchesOk returns a tuple with the SegblockOriginFetches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegblockOriginFetches

`func (o *RealtimeMeasurements) SetSegblockOriginFetches(v int32)`

SetSegblockOriginFetches sets SegblockOriginFetches field to given value.

### HasSegblockOriginFetches

`func (o *RealtimeMeasurements) HasSegblockOriginFetches() bool`

HasSegblockOriginFetches returns a boolean if a field has been set.

### GetSegblockShieldFetches

`func (o *RealtimeMeasurements) GetSegblockShieldFetches() int32`

GetSegblockShieldFetches returns the SegblockShieldFetches field if non-nil, zero value otherwise.

### GetSegblockShieldFetchesOk

`func (o *RealtimeMeasurements) GetSegblockShieldFetchesOk() (*int32, bool)`

GetSegblockShieldFetchesOk returns a tuple with the SegblockShieldFetches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegblockShieldFetches

`func (o *RealtimeMeasurements) SetSegblockShieldFetches(v int32)`

SetSegblockShieldFetches sets SegblockShieldFetches field to given value.

### HasSegblockShieldFetches

`func (o *RealtimeMeasurements) HasSegblockShieldFetches() bool`

HasSegblockShieldFetches returns a boolean if a field has been set.

### GetComputeRespStatus1xx

`func (o *RealtimeMeasurements) GetComputeRespStatus1xx() int32`

GetComputeRespStatus1xx returns the ComputeRespStatus1xx field if non-nil, zero value otherwise.

### GetComputeRespStatus1xxOk

`func (o *RealtimeMeasurements) GetComputeRespStatus1xxOk() (*int32, bool)`

GetComputeRespStatus1xxOk returns a tuple with the ComputeRespStatus1xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespStatus1xx

`func (o *RealtimeMeasurements) SetComputeRespStatus1xx(v int32)`

SetComputeRespStatus1xx sets ComputeRespStatus1xx field to given value.

### HasComputeRespStatus1xx

`func (o *RealtimeMeasurements) HasComputeRespStatus1xx() bool`

HasComputeRespStatus1xx returns a boolean if a field has been set.

### GetComputeRespStatus2xx

`func (o *RealtimeMeasurements) GetComputeRespStatus2xx() int32`

GetComputeRespStatus2xx returns the ComputeRespStatus2xx field if non-nil, zero value otherwise.

### GetComputeRespStatus2xxOk

`func (o *RealtimeMeasurements) GetComputeRespStatus2xxOk() (*int32, bool)`

GetComputeRespStatus2xxOk returns a tuple with the ComputeRespStatus2xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespStatus2xx

`func (o *RealtimeMeasurements) SetComputeRespStatus2xx(v int32)`

SetComputeRespStatus2xx sets ComputeRespStatus2xx field to given value.

### HasComputeRespStatus2xx

`func (o *RealtimeMeasurements) HasComputeRespStatus2xx() bool`

HasComputeRespStatus2xx returns a boolean if a field has been set.

### GetComputeRespStatus3xx

`func (o *RealtimeMeasurements) GetComputeRespStatus3xx() int32`

GetComputeRespStatus3xx returns the ComputeRespStatus3xx field if non-nil, zero value otherwise.

### GetComputeRespStatus3xxOk

`func (o *RealtimeMeasurements) GetComputeRespStatus3xxOk() (*int32, bool)`

GetComputeRespStatus3xxOk returns a tuple with the ComputeRespStatus3xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespStatus3xx

`func (o *RealtimeMeasurements) SetComputeRespStatus3xx(v int32)`

SetComputeRespStatus3xx sets ComputeRespStatus3xx field to given value.

### HasComputeRespStatus3xx

`func (o *RealtimeMeasurements) HasComputeRespStatus3xx() bool`

HasComputeRespStatus3xx returns a boolean if a field has been set.

### GetComputeRespStatus4xx

`func (o *RealtimeMeasurements) GetComputeRespStatus4xx() int32`

GetComputeRespStatus4xx returns the ComputeRespStatus4xx field if non-nil, zero value otherwise.

### GetComputeRespStatus4xxOk

`func (o *RealtimeMeasurements) GetComputeRespStatus4xxOk() (*int32, bool)`

GetComputeRespStatus4xxOk returns a tuple with the ComputeRespStatus4xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespStatus4xx

`func (o *RealtimeMeasurements) SetComputeRespStatus4xx(v int32)`

SetComputeRespStatus4xx sets ComputeRespStatus4xx field to given value.

### HasComputeRespStatus4xx

`func (o *RealtimeMeasurements) HasComputeRespStatus4xx() bool`

HasComputeRespStatus4xx returns a boolean if a field has been set.

### GetComputeRespStatus5xx

`func (o *RealtimeMeasurements) GetComputeRespStatus5xx() int32`

GetComputeRespStatus5xx returns the ComputeRespStatus5xx field if non-nil, zero value otherwise.

### GetComputeRespStatus5xxOk

`func (o *RealtimeMeasurements) GetComputeRespStatus5xxOk() (*int32, bool)`

GetComputeRespStatus5xxOk returns a tuple with the ComputeRespStatus5xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespStatus5xx

`func (o *RealtimeMeasurements) SetComputeRespStatus5xx(v int32)`

SetComputeRespStatus5xx sets ComputeRespStatus5xx field to given value.

### HasComputeRespStatus5xx

`func (o *RealtimeMeasurements) HasComputeRespStatus5xx() bool`

HasComputeRespStatus5xx returns a boolean if a field has been set.

### GetEdgeHitRequests

`func (o *RealtimeMeasurements) GetEdgeHitRequests() int32`

GetEdgeHitRequests returns the EdgeHitRequests field if non-nil, zero value otherwise.

### GetEdgeHitRequestsOk

`func (o *RealtimeMeasurements) GetEdgeHitRequestsOk() (*int32, bool)`

GetEdgeHitRequestsOk returns a tuple with the EdgeHitRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeHitRequests

`func (o *RealtimeMeasurements) SetEdgeHitRequests(v int32)`

SetEdgeHitRequests sets EdgeHitRequests field to given value.

### HasEdgeHitRequests

`func (o *RealtimeMeasurements) HasEdgeHitRequests() bool`

HasEdgeHitRequests returns a boolean if a field has been set.

### GetEdgeMissRequests

`func (o *RealtimeMeasurements) GetEdgeMissRequests() int32`

GetEdgeMissRequests returns the EdgeMissRequests field if non-nil, zero value otherwise.

### GetEdgeMissRequestsOk

`func (o *RealtimeMeasurements) GetEdgeMissRequestsOk() (*int32, bool)`

GetEdgeMissRequestsOk returns a tuple with the EdgeMissRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeMissRequests

`func (o *RealtimeMeasurements) SetEdgeMissRequests(v int32)`

SetEdgeMissRequests sets EdgeMissRequests field to given value.

### HasEdgeMissRequests

`func (o *RealtimeMeasurements) HasEdgeMissRequests() bool`

HasEdgeMissRequests returns a boolean if a field has been set.

### GetComputeBereqHeaderBytes

`func (o *RealtimeMeasurements) GetComputeBereqHeaderBytes() int32`

GetComputeBereqHeaderBytes returns the ComputeBereqHeaderBytes field if non-nil, zero value otherwise.

### GetComputeBereqHeaderBytesOk

`func (o *RealtimeMeasurements) GetComputeBereqHeaderBytesOk() (*int32, bool)`

GetComputeBereqHeaderBytesOk returns a tuple with the ComputeBereqHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBereqHeaderBytes

`func (o *RealtimeMeasurements) SetComputeBereqHeaderBytes(v int32)`

SetComputeBereqHeaderBytes sets ComputeBereqHeaderBytes field to given value.

### HasComputeBereqHeaderBytes

`func (o *RealtimeMeasurements) HasComputeBereqHeaderBytes() bool`

HasComputeBereqHeaderBytes returns a boolean if a field has been set.

### GetComputeBereqBodyBytes

`func (o *RealtimeMeasurements) GetComputeBereqBodyBytes() int32`

GetComputeBereqBodyBytes returns the ComputeBereqBodyBytes field if non-nil, zero value otherwise.

### GetComputeBereqBodyBytesOk

`func (o *RealtimeMeasurements) GetComputeBereqBodyBytesOk() (*int32, bool)`

GetComputeBereqBodyBytesOk returns a tuple with the ComputeBereqBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBereqBodyBytes

`func (o *RealtimeMeasurements) SetComputeBereqBodyBytes(v int32)`

SetComputeBereqBodyBytes sets ComputeBereqBodyBytes field to given value.

### HasComputeBereqBodyBytes

`func (o *RealtimeMeasurements) HasComputeBereqBodyBytes() bool`

HasComputeBereqBodyBytes returns a boolean if a field has been set.

### GetComputeBerespHeaderBytes

`func (o *RealtimeMeasurements) GetComputeBerespHeaderBytes() int32`

GetComputeBerespHeaderBytes returns the ComputeBerespHeaderBytes field if non-nil, zero value otherwise.

### GetComputeBerespHeaderBytesOk

`func (o *RealtimeMeasurements) GetComputeBerespHeaderBytesOk() (*int32, bool)`

GetComputeBerespHeaderBytesOk returns a tuple with the ComputeBerespHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBerespHeaderBytes

`func (o *RealtimeMeasurements) SetComputeBerespHeaderBytes(v int32)`

SetComputeBerespHeaderBytes sets ComputeBerespHeaderBytes field to given value.

### HasComputeBerespHeaderBytes

`func (o *RealtimeMeasurements) HasComputeBerespHeaderBytes() bool`

HasComputeBerespHeaderBytes returns a boolean if a field has been set.

### GetComputeBerespBodyBytes

`func (o *RealtimeMeasurements) GetComputeBerespBodyBytes() int32`

GetComputeBerespBodyBytes returns the ComputeBerespBodyBytes field if non-nil, zero value otherwise.

### GetComputeBerespBodyBytesOk

`func (o *RealtimeMeasurements) GetComputeBerespBodyBytesOk() (*int32, bool)`

GetComputeBerespBodyBytesOk returns a tuple with the ComputeBerespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBerespBodyBytes

`func (o *RealtimeMeasurements) SetComputeBerespBodyBytes(v int32)`

SetComputeBerespBodyBytes sets ComputeBerespBodyBytes field to given value.

### HasComputeBerespBodyBytes

`func (o *RealtimeMeasurements) HasComputeBerespBodyBytes() bool`

HasComputeBerespBodyBytes returns a boolean if a field has been set.

### GetOriginCacheFetches

`func (o *RealtimeMeasurements) GetOriginCacheFetches() int32`

GetOriginCacheFetches returns the OriginCacheFetches field if non-nil, zero value otherwise.

### GetOriginCacheFetchesOk

`func (o *RealtimeMeasurements) GetOriginCacheFetchesOk() (*int32, bool)`

GetOriginCacheFetchesOk returns a tuple with the OriginCacheFetches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginCacheFetches

`func (o *RealtimeMeasurements) SetOriginCacheFetches(v int32)`

SetOriginCacheFetches sets OriginCacheFetches field to given value.

### HasOriginCacheFetches

`func (o *RealtimeMeasurements) HasOriginCacheFetches() bool`

HasOriginCacheFetches returns a boolean if a field has been set.

### GetShieldCacheFetches

`func (o *RealtimeMeasurements) GetShieldCacheFetches() int32`

GetShieldCacheFetches returns the ShieldCacheFetches field if non-nil, zero value otherwise.

### GetShieldCacheFetchesOk

`func (o *RealtimeMeasurements) GetShieldCacheFetchesOk() (*int32, bool)`

GetShieldCacheFetchesOk returns a tuple with the ShieldCacheFetches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShieldCacheFetches

`func (o *RealtimeMeasurements) SetShieldCacheFetches(v int32)`

SetShieldCacheFetches sets ShieldCacheFetches field to given value.

### HasShieldCacheFetches

`func (o *RealtimeMeasurements) HasShieldCacheFetches() bool`

HasShieldCacheFetches returns a boolean if a field has been set.

### GetComputeBereqs

`func (o *RealtimeMeasurements) GetComputeBereqs() int32`

GetComputeBereqs returns the ComputeBereqs field if non-nil, zero value otherwise.

### GetComputeBereqsOk

`func (o *RealtimeMeasurements) GetComputeBereqsOk() (*int32, bool)`

GetComputeBereqsOk returns a tuple with the ComputeBereqs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBereqs

`func (o *RealtimeMeasurements) SetComputeBereqs(v int32)`

SetComputeBereqs sets ComputeBereqs field to given value.

### HasComputeBereqs

`func (o *RealtimeMeasurements) HasComputeBereqs() bool`

HasComputeBereqs returns a boolean if a field has been set.

### GetComputeBereqErrors

`func (o *RealtimeMeasurements) GetComputeBereqErrors() int32`

GetComputeBereqErrors returns the ComputeBereqErrors field if non-nil, zero value otherwise.

### GetComputeBereqErrorsOk

`func (o *RealtimeMeasurements) GetComputeBereqErrorsOk() (*int32, bool)`

GetComputeBereqErrorsOk returns a tuple with the ComputeBereqErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBereqErrors

`func (o *RealtimeMeasurements) SetComputeBereqErrors(v int32)`

SetComputeBereqErrors sets ComputeBereqErrors field to given value.

### HasComputeBereqErrors

`func (o *RealtimeMeasurements) HasComputeBereqErrors() bool`

HasComputeBereqErrors returns a boolean if a field has been set.

### GetComputeResourceLimitExceeded

`func (o *RealtimeMeasurements) GetComputeResourceLimitExceeded() int32`

GetComputeResourceLimitExceeded returns the ComputeResourceLimitExceeded field if non-nil, zero value otherwise.

### GetComputeResourceLimitExceededOk

`func (o *RealtimeMeasurements) GetComputeResourceLimitExceededOk() (*int32, bool)`

GetComputeResourceLimitExceededOk returns a tuple with the ComputeResourceLimitExceeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeResourceLimitExceeded

`func (o *RealtimeMeasurements) SetComputeResourceLimitExceeded(v int32)`

SetComputeResourceLimitExceeded sets ComputeResourceLimitExceeded field to given value.

### HasComputeResourceLimitExceeded

`func (o *RealtimeMeasurements) HasComputeResourceLimitExceeded() bool`

HasComputeResourceLimitExceeded returns a boolean if a field has been set.

### GetComputeHeapLimitExceeded

`func (o *RealtimeMeasurements) GetComputeHeapLimitExceeded() int32`

GetComputeHeapLimitExceeded returns the ComputeHeapLimitExceeded field if non-nil, zero value otherwise.

### GetComputeHeapLimitExceededOk

`func (o *RealtimeMeasurements) GetComputeHeapLimitExceededOk() (*int32, bool)`

GetComputeHeapLimitExceededOk returns a tuple with the ComputeHeapLimitExceeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeHeapLimitExceeded

`func (o *RealtimeMeasurements) SetComputeHeapLimitExceeded(v int32)`

SetComputeHeapLimitExceeded sets ComputeHeapLimitExceeded field to given value.

### HasComputeHeapLimitExceeded

`func (o *RealtimeMeasurements) HasComputeHeapLimitExceeded() bool`

HasComputeHeapLimitExceeded returns a boolean if a field has been set.

### GetComputeStackLimitExceeded

`func (o *RealtimeMeasurements) GetComputeStackLimitExceeded() int32`

GetComputeStackLimitExceeded returns the ComputeStackLimitExceeded field if non-nil, zero value otherwise.

### GetComputeStackLimitExceededOk

`func (o *RealtimeMeasurements) GetComputeStackLimitExceededOk() (*int32, bool)`

GetComputeStackLimitExceededOk returns a tuple with the ComputeStackLimitExceeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeStackLimitExceeded

`func (o *RealtimeMeasurements) SetComputeStackLimitExceeded(v int32)`

SetComputeStackLimitExceeded sets ComputeStackLimitExceeded field to given value.

### HasComputeStackLimitExceeded

`func (o *RealtimeMeasurements) HasComputeStackLimitExceeded() bool`

HasComputeStackLimitExceeded returns a boolean if a field has been set.

### GetComputeGlobalsLimitExceeded

`func (o *RealtimeMeasurements) GetComputeGlobalsLimitExceeded() int32`

GetComputeGlobalsLimitExceeded returns the ComputeGlobalsLimitExceeded field if non-nil, zero value otherwise.

### GetComputeGlobalsLimitExceededOk

`func (o *RealtimeMeasurements) GetComputeGlobalsLimitExceededOk() (*int32, bool)`

GetComputeGlobalsLimitExceededOk returns a tuple with the ComputeGlobalsLimitExceeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeGlobalsLimitExceeded

`func (o *RealtimeMeasurements) SetComputeGlobalsLimitExceeded(v int32)`

SetComputeGlobalsLimitExceeded sets ComputeGlobalsLimitExceeded field to given value.

### HasComputeGlobalsLimitExceeded

`func (o *RealtimeMeasurements) HasComputeGlobalsLimitExceeded() bool`

HasComputeGlobalsLimitExceeded returns a boolean if a field has been set.

### GetComputeGuestErrors

`func (o *RealtimeMeasurements) GetComputeGuestErrors() int32`

GetComputeGuestErrors returns the ComputeGuestErrors field if non-nil, zero value otherwise.

### GetComputeGuestErrorsOk

`func (o *RealtimeMeasurements) GetComputeGuestErrorsOk() (*int32, bool)`

GetComputeGuestErrorsOk returns a tuple with the ComputeGuestErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeGuestErrors

`func (o *RealtimeMeasurements) SetComputeGuestErrors(v int32)`

SetComputeGuestErrors sets ComputeGuestErrors field to given value.

### HasComputeGuestErrors

`func (o *RealtimeMeasurements) HasComputeGuestErrors() bool`

HasComputeGuestErrors returns a boolean if a field has been set.

### GetComputeRuntimeErrors

`func (o *RealtimeMeasurements) GetComputeRuntimeErrors() int32`

GetComputeRuntimeErrors returns the ComputeRuntimeErrors field if non-nil, zero value otherwise.

### GetComputeRuntimeErrorsOk

`func (o *RealtimeMeasurements) GetComputeRuntimeErrorsOk() (*int32, bool)`

GetComputeRuntimeErrorsOk returns a tuple with the ComputeRuntimeErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRuntimeErrors

`func (o *RealtimeMeasurements) SetComputeRuntimeErrors(v int32)`

SetComputeRuntimeErrors sets ComputeRuntimeErrors field to given value.

### HasComputeRuntimeErrors

`func (o *RealtimeMeasurements) HasComputeRuntimeErrors() bool`

HasComputeRuntimeErrors returns a boolean if a field has been set.

### GetEdgeHitRespBodyBytes

`func (o *RealtimeMeasurements) GetEdgeHitRespBodyBytes() int32`

GetEdgeHitRespBodyBytes returns the EdgeHitRespBodyBytes field if non-nil, zero value otherwise.

### GetEdgeHitRespBodyBytesOk

`func (o *RealtimeMeasurements) GetEdgeHitRespBodyBytesOk() (*int32, bool)`

GetEdgeHitRespBodyBytesOk returns a tuple with the EdgeHitRespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeHitRespBodyBytes

`func (o *RealtimeMeasurements) SetEdgeHitRespBodyBytes(v int32)`

SetEdgeHitRespBodyBytes sets EdgeHitRespBodyBytes field to given value.

### HasEdgeHitRespBodyBytes

`func (o *RealtimeMeasurements) HasEdgeHitRespBodyBytes() bool`

HasEdgeHitRespBodyBytes returns a boolean if a field has been set.

### GetEdgeHitRespHeaderBytes

`func (o *RealtimeMeasurements) GetEdgeHitRespHeaderBytes() int32`

GetEdgeHitRespHeaderBytes returns the EdgeHitRespHeaderBytes field if non-nil, zero value otherwise.

### GetEdgeHitRespHeaderBytesOk

`func (o *RealtimeMeasurements) GetEdgeHitRespHeaderBytesOk() (*int32, bool)`

GetEdgeHitRespHeaderBytesOk returns a tuple with the EdgeHitRespHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeHitRespHeaderBytes

`func (o *RealtimeMeasurements) SetEdgeHitRespHeaderBytes(v int32)`

SetEdgeHitRespHeaderBytes sets EdgeHitRespHeaderBytes field to given value.

### HasEdgeHitRespHeaderBytes

`func (o *RealtimeMeasurements) HasEdgeHitRespHeaderBytes() bool`

HasEdgeHitRespHeaderBytes returns a boolean if a field has been set.

### GetEdgeMissRespBodyBytes

`func (o *RealtimeMeasurements) GetEdgeMissRespBodyBytes() int32`

GetEdgeMissRespBodyBytes returns the EdgeMissRespBodyBytes field if non-nil, zero value otherwise.

### GetEdgeMissRespBodyBytesOk

`func (o *RealtimeMeasurements) GetEdgeMissRespBodyBytesOk() (*int32, bool)`

GetEdgeMissRespBodyBytesOk returns a tuple with the EdgeMissRespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeMissRespBodyBytes

`func (o *RealtimeMeasurements) SetEdgeMissRespBodyBytes(v int32)`

SetEdgeMissRespBodyBytes sets EdgeMissRespBodyBytes field to given value.

### HasEdgeMissRespBodyBytes

`func (o *RealtimeMeasurements) HasEdgeMissRespBodyBytes() bool`

HasEdgeMissRespBodyBytes returns a boolean if a field has been set.

### GetEdgeMissRespHeaderBytes

`func (o *RealtimeMeasurements) GetEdgeMissRespHeaderBytes() int32`

GetEdgeMissRespHeaderBytes returns the EdgeMissRespHeaderBytes field if non-nil, zero value otherwise.

### GetEdgeMissRespHeaderBytesOk

`func (o *RealtimeMeasurements) GetEdgeMissRespHeaderBytesOk() (*int32, bool)`

GetEdgeMissRespHeaderBytesOk returns a tuple with the EdgeMissRespHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeMissRespHeaderBytes

`func (o *RealtimeMeasurements) SetEdgeMissRespHeaderBytes(v int32)`

SetEdgeMissRespHeaderBytes sets EdgeMissRespHeaderBytes field to given value.

### HasEdgeMissRespHeaderBytes

`func (o *RealtimeMeasurements) HasEdgeMissRespHeaderBytes() bool`

HasEdgeMissRespHeaderBytes returns a boolean if a field has been set.

### GetOriginCacheFetchRespBodyBytes

`func (o *RealtimeMeasurements) GetOriginCacheFetchRespBodyBytes() int32`

GetOriginCacheFetchRespBodyBytes returns the OriginCacheFetchRespBodyBytes field if non-nil, zero value otherwise.

### GetOriginCacheFetchRespBodyBytesOk

`func (o *RealtimeMeasurements) GetOriginCacheFetchRespBodyBytesOk() (*int32, bool)`

GetOriginCacheFetchRespBodyBytesOk returns a tuple with the OriginCacheFetchRespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginCacheFetchRespBodyBytes

`func (o *RealtimeMeasurements) SetOriginCacheFetchRespBodyBytes(v int32)`

SetOriginCacheFetchRespBodyBytes sets OriginCacheFetchRespBodyBytes field to given value.

### HasOriginCacheFetchRespBodyBytes

`func (o *RealtimeMeasurements) HasOriginCacheFetchRespBodyBytes() bool`

HasOriginCacheFetchRespBodyBytes returns a boolean if a field has been set.

### GetOriginCacheFetchRespHeaderBytes

`func (o *RealtimeMeasurements) GetOriginCacheFetchRespHeaderBytes() int32`

GetOriginCacheFetchRespHeaderBytes returns the OriginCacheFetchRespHeaderBytes field if non-nil, zero value otherwise.

### GetOriginCacheFetchRespHeaderBytesOk

`func (o *RealtimeMeasurements) GetOriginCacheFetchRespHeaderBytesOk() (*int32, bool)`

GetOriginCacheFetchRespHeaderBytesOk returns a tuple with the OriginCacheFetchRespHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginCacheFetchRespHeaderBytes

`func (o *RealtimeMeasurements) SetOriginCacheFetchRespHeaderBytes(v int32)`

SetOriginCacheFetchRespHeaderBytes sets OriginCacheFetchRespHeaderBytes field to given value.

### HasOriginCacheFetchRespHeaderBytes

`func (o *RealtimeMeasurements) HasOriginCacheFetchRespHeaderBytes() bool`

HasOriginCacheFetchRespHeaderBytes returns a boolean if a field has been set.

### GetShieldHitRequests

`func (o *RealtimeMeasurements) GetShieldHitRequests() int32`

GetShieldHitRequests returns the ShieldHitRequests field if non-nil, zero value otherwise.

### GetShieldHitRequestsOk

`func (o *RealtimeMeasurements) GetShieldHitRequestsOk() (*int32, bool)`

GetShieldHitRequestsOk returns a tuple with the ShieldHitRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShieldHitRequests

`func (o *RealtimeMeasurements) SetShieldHitRequests(v int32)`

SetShieldHitRequests sets ShieldHitRequests field to given value.

### HasShieldHitRequests

`func (o *RealtimeMeasurements) HasShieldHitRequests() bool`

HasShieldHitRequests returns a boolean if a field has been set.

### GetShieldMissRequests

`func (o *RealtimeMeasurements) GetShieldMissRequests() int32`

GetShieldMissRequests returns the ShieldMissRequests field if non-nil, zero value otherwise.

### GetShieldMissRequestsOk

`func (o *RealtimeMeasurements) GetShieldMissRequestsOk() (*int32, bool)`

GetShieldMissRequestsOk returns a tuple with the ShieldMissRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShieldMissRequests

`func (o *RealtimeMeasurements) SetShieldMissRequests(v int32)`

SetShieldMissRequests sets ShieldMissRequests field to given value.

### HasShieldMissRequests

`func (o *RealtimeMeasurements) HasShieldMissRequests() bool`

HasShieldMissRequests returns a boolean if a field has been set.

### GetShieldHitRespHeaderBytes

`func (o *RealtimeMeasurements) GetShieldHitRespHeaderBytes() int32`

GetShieldHitRespHeaderBytes returns the ShieldHitRespHeaderBytes field if non-nil, zero value otherwise.

### GetShieldHitRespHeaderBytesOk

`func (o *RealtimeMeasurements) GetShieldHitRespHeaderBytesOk() (*int32, bool)`

GetShieldHitRespHeaderBytesOk returns a tuple with the ShieldHitRespHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShieldHitRespHeaderBytes

`func (o *RealtimeMeasurements) SetShieldHitRespHeaderBytes(v int32)`

SetShieldHitRespHeaderBytes sets ShieldHitRespHeaderBytes field to given value.

### HasShieldHitRespHeaderBytes

`func (o *RealtimeMeasurements) HasShieldHitRespHeaderBytes() bool`

HasShieldHitRespHeaderBytes returns a boolean if a field has been set.

### GetShieldHitRespBodyBytes

`func (o *RealtimeMeasurements) GetShieldHitRespBodyBytes() int32`

GetShieldHitRespBodyBytes returns the ShieldHitRespBodyBytes field if non-nil, zero value otherwise.

### GetShieldHitRespBodyBytesOk

`func (o *RealtimeMeasurements) GetShieldHitRespBodyBytesOk() (*int32, bool)`

GetShieldHitRespBodyBytesOk returns a tuple with the ShieldHitRespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShieldHitRespBodyBytes

`func (o *RealtimeMeasurements) SetShieldHitRespBodyBytes(v int32)`

SetShieldHitRespBodyBytes sets ShieldHitRespBodyBytes field to given value.

### HasShieldHitRespBodyBytes

`func (o *RealtimeMeasurements) HasShieldHitRespBodyBytes() bool`

HasShieldHitRespBodyBytes returns a boolean if a field has been set.

### GetShieldMissRespHeaderBytes

`func (o *RealtimeMeasurements) GetShieldMissRespHeaderBytes() int32`

GetShieldMissRespHeaderBytes returns the ShieldMissRespHeaderBytes field if non-nil, zero value otherwise.

### GetShieldMissRespHeaderBytesOk

`func (o *RealtimeMeasurements) GetShieldMissRespHeaderBytesOk() (*int32, bool)`

GetShieldMissRespHeaderBytesOk returns a tuple with the ShieldMissRespHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShieldMissRespHeaderBytes

`func (o *RealtimeMeasurements) SetShieldMissRespHeaderBytes(v int32)`

SetShieldMissRespHeaderBytes sets ShieldMissRespHeaderBytes field to given value.

### HasShieldMissRespHeaderBytes

`func (o *RealtimeMeasurements) HasShieldMissRespHeaderBytes() bool`

HasShieldMissRespHeaderBytes returns a boolean if a field has been set.

### GetShieldMissRespBodyBytes

`func (o *RealtimeMeasurements) GetShieldMissRespBodyBytes() int32`

GetShieldMissRespBodyBytes returns the ShieldMissRespBodyBytes field if non-nil, zero value otherwise.

### GetShieldMissRespBodyBytesOk

`func (o *RealtimeMeasurements) GetShieldMissRespBodyBytesOk() (*int32, bool)`

GetShieldMissRespBodyBytesOk returns a tuple with the ShieldMissRespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShieldMissRespBodyBytes

`func (o *RealtimeMeasurements) SetShieldMissRespBodyBytes(v int32)`

SetShieldMissRespBodyBytes sets ShieldMissRespBodyBytes field to given value.

### HasShieldMissRespBodyBytes

`func (o *RealtimeMeasurements) HasShieldMissRespBodyBytes() bool`

HasShieldMissRespBodyBytes returns a boolean if a field has been set.

### GetWebsocketReqHeaderBytes

`func (o *RealtimeMeasurements) GetWebsocketReqHeaderBytes() int32`

GetWebsocketReqHeaderBytes returns the WebsocketReqHeaderBytes field if non-nil, zero value otherwise.

### GetWebsocketReqHeaderBytesOk

`func (o *RealtimeMeasurements) GetWebsocketReqHeaderBytesOk() (*int32, bool)`

GetWebsocketReqHeaderBytesOk returns a tuple with the WebsocketReqHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsocketReqHeaderBytes

`func (o *RealtimeMeasurements) SetWebsocketReqHeaderBytes(v int32)`

SetWebsocketReqHeaderBytes sets WebsocketReqHeaderBytes field to given value.

### HasWebsocketReqHeaderBytes

`func (o *RealtimeMeasurements) HasWebsocketReqHeaderBytes() bool`

HasWebsocketReqHeaderBytes returns a boolean if a field has been set.

### GetWebsocketReqBodyBytes

`func (o *RealtimeMeasurements) GetWebsocketReqBodyBytes() int32`

GetWebsocketReqBodyBytes returns the WebsocketReqBodyBytes field if non-nil, zero value otherwise.

### GetWebsocketReqBodyBytesOk

`func (o *RealtimeMeasurements) GetWebsocketReqBodyBytesOk() (*int32, bool)`

GetWebsocketReqBodyBytesOk returns a tuple with the WebsocketReqBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsocketReqBodyBytes

`func (o *RealtimeMeasurements) SetWebsocketReqBodyBytes(v int32)`

SetWebsocketReqBodyBytes sets WebsocketReqBodyBytes field to given value.

### HasWebsocketReqBodyBytes

`func (o *RealtimeMeasurements) HasWebsocketReqBodyBytes() bool`

HasWebsocketReqBodyBytes returns a boolean if a field has been set.

### GetWebsocketRespHeaderBytes

`func (o *RealtimeMeasurements) GetWebsocketRespHeaderBytes() int32`

GetWebsocketRespHeaderBytes returns the WebsocketRespHeaderBytes field if non-nil, zero value otherwise.

### GetWebsocketRespHeaderBytesOk

`func (o *RealtimeMeasurements) GetWebsocketRespHeaderBytesOk() (*int32, bool)`

GetWebsocketRespHeaderBytesOk returns a tuple with the WebsocketRespHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsocketRespHeaderBytes

`func (o *RealtimeMeasurements) SetWebsocketRespHeaderBytes(v int32)`

SetWebsocketRespHeaderBytes sets WebsocketRespHeaderBytes field to given value.

### HasWebsocketRespHeaderBytes

`func (o *RealtimeMeasurements) HasWebsocketRespHeaderBytes() bool`

HasWebsocketRespHeaderBytes returns a boolean if a field has been set.

### GetWebsocketBereqHeaderBytes

`func (o *RealtimeMeasurements) GetWebsocketBereqHeaderBytes() int32`

GetWebsocketBereqHeaderBytes returns the WebsocketBereqHeaderBytes field if non-nil, zero value otherwise.

### GetWebsocketBereqHeaderBytesOk

`func (o *RealtimeMeasurements) GetWebsocketBereqHeaderBytesOk() (*int32, bool)`

GetWebsocketBereqHeaderBytesOk returns a tuple with the WebsocketBereqHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsocketBereqHeaderBytes

`func (o *RealtimeMeasurements) SetWebsocketBereqHeaderBytes(v int32)`

SetWebsocketBereqHeaderBytes sets WebsocketBereqHeaderBytes field to given value.

### HasWebsocketBereqHeaderBytes

`func (o *RealtimeMeasurements) HasWebsocketBereqHeaderBytes() bool`

HasWebsocketBereqHeaderBytes returns a boolean if a field has been set.

### GetWebsocketBereqBodyBytes

`func (o *RealtimeMeasurements) GetWebsocketBereqBodyBytes() int32`

GetWebsocketBereqBodyBytes returns the WebsocketBereqBodyBytes field if non-nil, zero value otherwise.

### GetWebsocketBereqBodyBytesOk

`func (o *RealtimeMeasurements) GetWebsocketBereqBodyBytesOk() (*int32, bool)`

GetWebsocketBereqBodyBytesOk returns a tuple with the WebsocketBereqBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsocketBereqBodyBytes

`func (o *RealtimeMeasurements) SetWebsocketBereqBodyBytes(v int32)`

SetWebsocketBereqBodyBytes sets WebsocketBereqBodyBytes field to given value.

### HasWebsocketBereqBodyBytes

`func (o *RealtimeMeasurements) HasWebsocketBereqBodyBytes() bool`

HasWebsocketBereqBodyBytes returns a boolean if a field has been set.

### GetWebsocketBerespHeaderBytes

`func (o *RealtimeMeasurements) GetWebsocketBerespHeaderBytes() int32`

GetWebsocketBerespHeaderBytes returns the WebsocketBerespHeaderBytes field if non-nil, zero value otherwise.

### GetWebsocketBerespHeaderBytesOk

`func (o *RealtimeMeasurements) GetWebsocketBerespHeaderBytesOk() (*int32, bool)`

GetWebsocketBerespHeaderBytesOk returns a tuple with the WebsocketBerespHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsocketBerespHeaderBytes

`func (o *RealtimeMeasurements) SetWebsocketBerespHeaderBytes(v int32)`

SetWebsocketBerespHeaderBytes sets WebsocketBerespHeaderBytes field to given value.

### HasWebsocketBerespHeaderBytes

`func (o *RealtimeMeasurements) HasWebsocketBerespHeaderBytes() bool`

HasWebsocketBerespHeaderBytes returns a boolean if a field has been set.

### GetWebsocketBerespBodyBytes

`func (o *RealtimeMeasurements) GetWebsocketBerespBodyBytes() int32`

GetWebsocketBerespBodyBytes returns the WebsocketBerespBodyBytes field if non-nil, zero value otherwise.

### GetWebsocketBerespBodyBytesOk

`func (o *RealtimeMeasurements) GetWebsocketBerespBodyBytesOk() (*int32, bool)`

GetWebsocketBerespBodyBytesOk returns a tuple with the WebsocketBerespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsocketBerespBodyBytes

`func (o *RealtimeMeasurements) SetWebsocketBerespBodyBytes(v int32)`

SetWebsocketBerespBodyBytes sets WebsocketBerespBodyBytes field to given value.

### HasWebsocketBerespBodyBytes

`func (o *RealtimeMeasurements) HasWebsocketBerespBodyBytes() bool`

HasWebsocketBerespBodyBytes returns a boolean if a field has been set.

### GetWebsocketConnTimeMs

`func (o *RealtimeMeasurements) GetWebsocketConnTimeMs() int32`

GetWebsocketConnTimeMs returns the WebsocketConnTimeMs field if non-nil, zero value otherwise.

### GetWebsocketConnTimeMsOk

`func (o *RealtimeMeasurements) GetWebsocketConnTimeMsOk() (*int32, bool)`

GetWebsocketConnTimeMsOk returns a tuple with the WebsocketConnTimeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsocketConnTimeMs

`func (o *RealtimeMeasurements) SetWebsocketConnTimeMs(v int32)`

SetWebsocketConnTimeMs sets WebsocketConnTimeMs field to given value.

### HasWebsocketConnTimeMs

`func (o *RealtimeMeasurements) HasWebsocketConnTimeMs() bool`

HasWebsocketConnTimeMs returns a boolean if a field has been set.

### GetWebsocketRespBodyBytes

`func (o *RealtimeMeasurements) GetWebsocketRespBodyBytes() int32`

GetWebsocketRespBodyBytes returns the WebsocketRespBodyBytes field if non-nil, zero value otherwise.

### GetWebsocketRespBodyBytesOk

`func (o *RealtimeMeasurements) GetWebsocketRespBodyBytesOk() (*int32, bool)`

GetWebsocketRespBodyBytesOk returns a tuple with the WebsocketRespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsocketRespBodyBytes

`func (o *RealtimeMeasurements) SetWebsocketRespBodyBytes(v int32)`

SetWebsocketRespBodyBytes sets WebsocketRespBodyBytes field to given value.

### HasWebsocketRespBodyBytes

`func (o *RealtimeMeasurements) HasWebsocketRespBodyBytes() bool`

HasWebsocketRespBodyBytes returns a boolean if a field has been set.

### GetFanoutRecvPublishes

`func (o *RealtimeMeasurements) GetFanoutRecvPublishes() int32`

GetFanoutRecvPublishes returns the FanoutRecvPublishes field if non-nil, zero value otherwise.

### GetFanoutRecvPublishesOk

`func (o *RealtimeMeasurements) GetFanoutRecvPublishesOk() (*int32, bool)`

GetFanoutRecvPublishesOk returns a tuple with the FanoutRecvPublishes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFanoutRecvPublishes

`func (o *RealtimeMeasurements) SetFanoutRecvPublishes(v int32)`

SetFanoutRecvPublishes sets FanoutRecvPublishes field to given value.

### HasFanoutRecvPublishes

`func (o *RealtimeMeasurements) HasFanoutRecvPublishes() bool`

HasFanoutRecvPublishes returns a boolean if a field has been set.

### GetFanoutSendPublishes

`func (o *RealtimeMeasurements) GetFanoutSendPublishes() int32`

GetFanoutSendPublishes returns the FanoutSendPublishes field if non-nil, zero value otherwise.

### GetFanoutSendPublishesOk

`func (o *RealtimeMeasurements) GetFanoutSendPublishesOk() (*int32, bool)`

GetFanoutSendPublishesOk returns a tuple with the FanoutSendPublishes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFanoutSendPublishes

`func (o *RealtimeMeasurements) SetFanoutSendPublishes(v int32)`

SetFanoutSendPublishes sets FanoutSendPublishes field to given value.

### HasFanoutSendPublishes

`func (o *RealtimeMeasurements) HasFanoutSendPublishes() bool`

HasFanoutSendPublishes returns a boolean if a field has been set.

### GetKvStoreClassAOperations

`func (o *RealtimeMeasurements) GetKvStoreClassAOperations() int32`

GetKvStoreClassAOperations returns the KvStoreClassAOperations field if non-nil, zero value otherwise.

### GetKvStoreClassAOperationsOk

`func (o *RealtimeMeasurements) GetKvStoreClassAOperationsOk() (*int32, bool)`

GetKvStoreClassAOperationsOk returns a tuple with the KvStoreClassAOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKvStoreClassAOperations

`func (o *RealtimeMeasurements) SetKvStoreClassAOperations(v int32)`

SetKvStoreClassAOperations sets KvStoreClassAOperations field to given value.

### HasKvStoreClassAOperations

`func (o *RealtimeMeasurements) HasKvStoreClassAOperations() bool`

HasKvStoreClassAOperations returns a boolean if a field has been set.

### GetKvStoreClassBOperations

`func (o *RealtimeMeasurements) GetKvStoreClassBOperations() int32`

GetKvStoreClassBOperations returns the KvStoreClassBOperations field if non-nil, zero value otherwise.

### GetKvStoreClassBOperationsOk

`func (o *RealtimeMeasurements) GetKvStoreClassBOperationsOk() (*int32, bool)`

GetKvStoreClassBOperationsOk returns a tuple with the KvStoreClassBOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKvStoreClassBOperations

`func (o *RealtimeMeasurements) SetKvStoreClassBOperations(v int32)`

SetKvStoreClassBOperations sets KvStoreClassBOperations field to given value.

### HasKvStoreClassBOperations

`func (o *RealtimeMeasurements) HasKvStoreClassBOperations() bool`

HasKvStoreClassBOperations returns a boolean if a field has been set.

### GetObjectStoreClassAOperations

`func (o *RealtimeMeasurements) GetObjectStoreClassAOperations() int32`

GetObjectStoreClassAOperations returns the ObjectStoreClassAOperations field if non-nil, zero value otherwise.

### GetObjectStoreClassAOperationsOk

`func (o *RealtimeMeasurements) GetObjectStoreClassAOperationsOk() (*int32, bool)`

GetObjectStoreClassAOperationsOk returns a tuple with the ObjectStoreClassAOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectStoreClassAOperations

`func (o *RealtimeMeasurements) SetObjectStoreClassAOperations(v int32)`

SetObjectStoreClassAOperations sets ObjectStoreClassAOperations field to given value.

### HasObjectStoreClassAOperations

`func (o *RealtimeMeasurements) HasObjectStoreClassAOperations() bool`

HasObjectStoreClassAOperations returns a boolean if a field has been set.

### GetObjectStoreClassBOperations

`func (o *RealtimeMeasurements) GetObjectStoreClassBOperations() int32`

GetObjectStoreClassBOperations returns the ObjectStoreClassBOperations field if non-nil, zero value otherwise.

### GetObjectStoreClassBOperationsOk

`func (o *RealtimeMeasurements) GetObjectStoreClassBOperationsOk() (*int32, bool)`

GetObjectStoreClassBOperationsOk returns a tuple with the ObjectStoreClassBOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectStoreClassBOperations

`func (o *RealtimeMeasurements) SetObjectStoreClassBOperations(v int32)`

SetObjectStoreClassBOperations sets ObjectStoreClassBOperations field to given value.

### HasObjectStoreClassBOperations

`func (o *RealtimeMeasurements) HasObjectStoreClassBOperations() bool`

HasObjectStoreClassBOperations returns a boolean if a field has been set.

### GetFanoutReqHeaderBytes

`func (o *RealtimeMeasurements) GetFanoutReqHeaderBytes() int32`

GetFanoutReqHeaderBytes returns the FanoutReqHeaderBytes field if non-nil, zero value otherwise.

### GetFanoutReqHeaderBytesOk

`func (o *RealtimeMeasurements) GetFanoutReqHeaderBytesOk() (*int32, bool)`

GetFanoutReqHeaderBytesOk returns a tuple with the FanoutReqHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFanoutReqHeaderBytes

`func (o *RealtimeMeasurements) SetFanoutReqHeaderBytes(v int32)`

SetFanoutReqHeaderBytes sets FanoutReqHeaderBytes field to given value.

### HasFanoutReqHeaderBytes

`func (o *RealtimeMeasurements) HasFanoutReqHeaderBytes() bool`

HasFanoutReqHeaderBytes returns a boolean if a field has been set.

### GetFanoutReqBodyBytes

`func (o *RealtimeMeasurements) GetFanoutReqBodyBytes() int32`

GetFanoutReqBodyBytes returns the FanoutReqBodyBytes field if non-nil, zero value otherwise.

### GetFanoutReqBodyBytesOk

`func (o *RealtimeMeasurements) GetFanoutReqBodyBytesOk() (*int32, bool)`

GetFanoutReqBodyBytesOk returns a tuple with the FanoutReqBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFanoutReqBodyBytes

`func (o *RealtimeMeasurements) SetFanoutReqBodyBytes(v int32)`

SetFanoutReqBodyBytes sets FanoutReqBodyBytes field to given value.

### HasFanoutReqBodyBytes

`func (o *RealtimeMeasurements) HasFanoutReqBodyBytes() bool`

HasFanoutReqBodyBytes returns a boolean if a field has been set.

### GetFanoutRespHeaderBytes

`func (o *RealtimeMeasurements) GetFanoutRespHeaderBytes() int32`

GetFanoutRespHeaderBytes returns the FanoutRespHeaderBytes field if non-nil, zero value otherwise.

### GetFanoutRespHeaderBytesOk

`func (o *RealtimeMeasurements) GetFanoutRespHeaderBytesOk() (*int32, bool)`

GetFanoutRespHeaderBytesOk returns a tuple with the FanoutRespHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFanoutRespHeaderBytes

`func (o *RealtimeMeasurements) SetFanoutRespHeaderBytes(v int32)`

SetFanoutRespHeaderBytes sets FanoutRespHeaderBytes field to given value.

### HasFanoutRespHeaderBytes

`func (o *RealtimeMeasurements) HasFanoutRespHeaderBytes() bool`

HasFanoutRespHeaderBytes returns a boolean if a field has been set.

### GetFanoutRespBodyBytes

`func (o *RealtimeMeasurements) GetFanoutRespBodyBytes() int32`

GetFanoutRespBodyBytes returns the FanoutRespBodyBytes field if non-nil, zero value otherwise.

### GetFanoutRespBodyBytesOk

`func (o *RealtimeMeasurements) GetFanoutRespBodyBytesOk() (*int32, bool)`

GetFanoutRespBodyBytesOk returns a tuple with the FanoutRespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFanoutRespBodyBytes

`func (o *RealtimeMeasurements) SetFanoutRespBodyBytes(v int32)`

SetFanoutRespBodyBytes sets FanoutRespBodyBytes field to given value.

### HasFanoutRespBodyBytes

`func (o *RealtimeMeasurements) HasFanoutRespBodyBytes() bool`

HasFanoutRespBodyBytes returns a boolean if a field has been set.

### GetFanoutBereqHeaderBytes

`func (o *RealtimeMeasurements) GetFanoutBereqHeaderBytes() int32`

GetFanoutBereqHeaderBytes returns the FanoutBereqHeaderBytes field if non-nil, zero value otherwise.

### GetFanoutBereqHeaderBytesOk

`func (o *RealtimeMeasurements) GetFanoutBereqHeaderBytesOk() (*int32, bool)`

GetFanoutBereqHeaderBytesOk returns a tuple with the FanoutBereqHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFanoutBereqHeaderBytes

`func (o *RealtimeMeasurements) SetFanoutBereqHeaderBytes(v int32)`

SetFanoutBereqHeaderBytes sets FanoutBereqHeaderBytes field to given value.

### HasFanoutBereqHeaderBytes

`func (o *RealtimeMeasurements) HasFanoutBereqHeaderBytes() bool`

HasFanoutBereqHeaderBytes returns a boolean if a field has been set.

### GetFanoutBereqBodyBytes

`func (o *RealtimeMeasurements) GetFanoutBereqBodyBytes() int32`

GetFanoutBereqBodyBytes returns the FanoutBereqBodyBytes field if non-nil, zero value otherwise.

### GetFanoutBereqBodyBytesOk

`func (o *RealtimeMeasurements) GetFanoutBereqBodyBytesOk() (*int32, bool)`

GetFanoutBereqBodyBytesOk returns a tuple with the FanoutBereqBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFanoutBereqBodyBytes

`func (o *RealtimeMeasurements) SetFanoutBereqBodyBytes(v int32)`

SetFanoutBereqBodyBytes sets FanoutBereqBodyBytes field to given value.

### HasFanoutBereqBodyBytes

`func (o *RealtimeMeasurements) HasFanoutBereqBodyBytes() bool`

HasFanoutBereqBodyBytes returns a boolean if a field has been set.

### GetFanoutBerespHeaderBytes

`func (o *RealtimeMeasurements) GetFanoutBerespHeaderBytes() int32`

GetFanoutBerespHeaderBytes returns the FanoutBerespHeaderBytes field if non-nil, zero value otherwise.

### GetFanoutBerespHeaderBytesOk

`func (o *RealtimeMeasurements) GetFanoutBerespHeaderBytesOk() (*int32, bool)`

GetFanoutBerespHeaderBytesOk returns a tuple with the FanoutBerespHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFanoutBerespHeaderBytes

`func (o *RealtimeMeasurements) SetFanoutBerespHeaderBytes(v int32)`

SetFanoutBerespHeaderBytes sets FanoutBerespHeaderBytes field to given value.

### HasFanoutBerespHeaderBytes

`func (o *RealtimeMeasurements) HasFanoutBerespHeaderBytes() bool`

HasFanoutBerespHeaderBytes returns a boolean if a field has been set.

### GetFanoutBerespBodyBytes

`func (o *RealtimeMeasurements) GetFanoutBerespBodyBytes() int32`

GetFanoutBerespBodyBytes returns the FanoutBerespBodyBytes field if non-nil, zero value otherwise.

### GetFanoutBerespBodyBytesOk

`func (o *RealtimeMeasurements) GetFanoutBerespBodyBytesOk() (*int32, bool)`

GetFanoutBerespBodyBytesOk returns a tuple with the FanoutBerespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFanoutBerespBodyBytes

`func (o *RealtimeMeasurements) SetFanoutBerespBodyBytes(v int32)`

SetFanoutBerespBodyBytes sets FanoutBerespBodyBytes field to given value.

### HasFanoutBerespBodyBytes

`func (o *RealtimeMeasurements) HasFanoutBerespBodyBytes() bool`

HasFanoutBerespBodyBytes returns a boolean if a field has been set.

### GetFanoutConnTimeMs

`func (o *RealtimeMeasurements) GetFanoutConnTimeMs() int32`

GetFanoutConnTimeMs returns the FanoutConnTimeMs field if non-nil, zero value otherwise.

### GetFanoutConnTimeMsOk

`func (o *RealtimeMeasurements) GetFanoutConnTimeMsOk() (*int32, bool)`

GetFanoutConnTimeMsOk returns a tuple with the FanoutConnTimeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFanoutConnTimeMs

`func (o *RealtimeMeasurements) SetFanoutConnTimeMs(v int32)`

SetFanoutConnTimeMs sets FanoutConnTimeMs field to given value.

### HasFanoutConnTimeMs

`func (o *RealtimeMeasurements) HasFanoutConnTimeMs() bool`

HasFanoutConnTimeMs returns a boolean if a field has been set.

### GetDdosActionLimitStreamsConnections

`func (o *RealtimeMeasurements) GetDdosActionLimitStreamsConnections() int32`

GetDdosActionLimitStreamsConnections returns the DdosActionLimitStreamsConnections field if non-nil, zero value otherwise.

### GetDdosActionLimitStreamsConnectionsOk

`func (o *RealtimeMeasurements) GetDdosActionLimitStreamsConnectionsOk() (*int32, bool)`

GetDdosActionLimitStreamsConnectionsOk returns a tuple with the DdosActionLimitStreamsConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdosActionLimitStreamsConnections

`func (o *RealtimeMeasurements) SetDdosActionLimitStreamsConnections(v int32)`

SetDdosActionLimitStreamsConnections sets DdosActionLimitStreamsConnections field to given value.

### HasDdosActionLimitStreamsConnections

`func (o *RealtimeMeasurements) HasDdosActionLimitStreamsConnections() bool`

HasDdosActionLimitStreamsConnections returns a boolean if a field has been set.

### GetDdosActionLimitStreamsRequests

`func (o *RealtimeMeasurements) GetDdosActionLimitStreamsRequests() int32`

GetDdosActionLimitStreamsRequests returns the DdosActionLimitStreamsRequests field if non-nil, zero value otherwise.

### GetDdosActionLimitStreamsRequestsOk

`func (o *RealtimeMeasurements) GetDdosActionLimitStreamsRequestsOk() (*int32, bool)`

GetDdosActionLimitStreamsRequestsOk returns a tuple with the DdosActionLimitStreamsRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdosActionLimitStreamsRequests

`func (o *RealtimeMeasurements) SetDdosActionLimitStreamsRequests(v int32)`

SetDdosActionLimitStreamsRequests sets DdosActionLimitStreamsRequests field to given value.

### HasDdosActionLimitStreamsRequests

`func (o *RealtimeMeasurements) HasDdosActionLimitStreamsRequests() bool`

HasDdosActionLimitStreamsRequests returns a boolean if a field has been set.

### GetDdosActionTarpitAccept

`func (o *RealtimeMeasurements) GetDdosActionTarpitAccept() int32`

GetDdosActionTarpitAccept returns the DdosActionTarpitAccept field if non-nil, zero value otherwise.

### GetDdosActionTarpitAcceptOk

`func (o *RealtimeMeasurements) GetDdosActionTarpitAcceptOk() (*int32, bool)`

GetDdosActionTarpitAcceptOk returns a tuple with the DdosActionTarpitAccept field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdosActionTarpitAccept

`func (o *RealtimeMeasurements) SetDdosActionTarpitAccept(v int32)`

SetDdosActionTarpitAccept sets DdosActionTarpitAccept field to given value.

### HasDdosActionTarpitAccept

`func (o *RealtimeMeasurements) HasDdosActionTarpitAccept() bool`

HasDdosActionTarpitAccept returns a boolean if a field has been set.

### GetDdosActionTarpit

`func (o *RealtimeMeasurements) GetDdosActionTarpit() int32`

GetDdosActionTarpit returns the DdosActionTarpit field if non-nil, zero value otherwise.

### GetDdosActionTarpitOk

`func (o *RealtimeMeasurements) GetDdosActionTarpitOk() (*int32, bool)`

GetDdosActionTarpitOk returns a tuple with the DdosActionTarpit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdosActionTarpit

`func (o *RealtimeMeasurements) SetDdosActionTarpit(v int32)`

SetDdosActionTarpit sets DdosActionTarpit field to given value.

### HasDdosActionTarpit

`func (o *RealtimeMeasurements) HasDdosActionTarpit() bool`

HasDdosActionTarpit returns a boolean if a field has been set.

### GetDdosActionClose

`func (o *RealtimeMeasurements) GetDdosActionClose() int32`

GetDdosActionClose returns the DdosActionClose field if non-nil, zero value otherwise.

### GetDdosActionCloseOk

`func (o *RealtimeMeasurements) GetDdosActionCloseOk() (*int32, bool)`

GetDdosActionCloseOk returns a tuple with the DdosActionClose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdosActionClose

`func (o *RealtimeMeasurements) SetDdosActionClose(v int32)`

SetDdosActionClose sets DdosActionClose field to given value.

### HasDdosActionClose

`func (o *RealtimeMeasurements) HasDdosActionClose() bool`

HasDdosActionClose returns a boolean if a field has been set.

### GetDdosActionBlackhole

`func (o *RealtimeMeasurements) GetDdosActionBlackhole() int32`

GetDdosActionBlackhole returns the DdosActionBlackhole field if non-nil, zero value otherwise.

### GetDdosActionBlackholeOk

`func (o *RealtimeMeasurements) GetDdosActionBlackholeOk() (*int32, bool)`

GetDdosActionBlackholeOk returns a tuple with the DdosActionBlackhole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdosActionBlackhole

`func (o *RealtimeMeasurements) SetDdosActionBlackhole(v int32)`

SetDdosActionBlackhole sets DdosActionBlackhole field to given value.

### HasDdosActionBlackhole

`func (o *RealtimeMeasurements) HasDdosActionBlackhole() bool`

HasDdosActionBlackhole returns a boolean if a field has been set.

### GetBotChallengeStarts

`func (o *RealtimeMeasurements) GetBotChallengeStarts() int32`

GetBotChallengeStarts returns the BotChallengeStarts field if non-nil, zero value otherwise.

### GetBotChallengeStartsOk

`func (o *RealtimeMeasurements) GetBotChallengeStartsOk() (*int32, bool)`

GetBotChallengeStartsOk returns a tuple with the BotChallengeStarts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotChallengeStarts

`func (o *RealtimeMeasurements) SetBotChallengeStarts(v int32)`

SetBotChallengeStarts sets BotChallengeStarts field to given value.

### HasBotChallengeStarts

`func (o *RealtimeMeasurements) HasBotChallengeStarts() bool`

HasBotChallengeStarts returns a boolean if a field has been set.

### GetBotChallengeCompleteTokensPassed

`func (o *RealtimeMeasurements) GetBotChallengeCompleteTokensPassed() int32`

GetBotChallengeCompleteTokensPassed returns the BotChallengeCompleteTokensPassed field if non-nil, zero value otherwise.

### GetBotChallengeCompleteTokensPassedOk

`func (o *RealtimeMeasurements) GetBotChallengeCompleteTokensPassedOk() (*int32, bool)`

GetBotChallengeCompleteTokensPassedOk returns a tuple with the BotChallengeCompleteTokensPassed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotChallengeCompleteTokensPassed

`func (o *RealtimeMeasurements) SetBotChallengeCompleteTokensPassed(v int32)`

SetBotChallengeCompleteTokensPassed sets BotChallengeCompleteTokensPassed field to given value.

### HasBotChallengeCompleteTokensPassed

`func (o *RealtimeMeasurements) HasBotChallengeCompleteTokensPassed() bool`

HasBotChallengeCompleteTokensPassed returns a boolean if a field has been set.

### GetBotChallengeCompleteTokensFailed

`func (o *RealtimeMeasurements) GetBotChallengeCompleteTokensFailed() int32`

GetBotChallengeCompleteTokensFailed returns the BotChallengeCompleteTokensFailed field if non-nil, zero value otherwise.

### GetBotChallengeCompleteTokensFailedOk

`func (o *RealtimeMeasurements) GetBotChallengeCompleteTokensFailedOk() (*int32, bool)`

GetBotChallengeCompleteTokensFailedOk returns a tuple with the BotChallengeCompleteTokensFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotChallengeCompleteTokensFailed

`func (o *RealtimeMeasurements) SetBotChallengeCompleteTokensFailed(v int32)`

SetBotChallengeCompleteTokensFailed sets BotChallengeCompleteTokensFailed field to given value.

### HasBotChallengeCompleteTokensFailed

`func (o *RealtimeMeasurements) HasBotChallengeCompleteTokensFailed() bool`

HasBotChallengeCompleteTokensFailed returns a boolean if a field has been set.

### GetBotChallengeCompleteTokensChecked

`func (o *RealtimeMeasurements) GetBotChallengeCompleteTokensChecked() int32`

GetBotChallengeCompleteTokensChecked returns the BotChallengeCompleteTokensChecked field if non-nil, zero value otherwise.

### GetBotChallengeCompleteTokensCheckedOk

`func (o *RealtimeMeasurements) GetBotChallengeCompleteTokensCheckedOk() (*int32, bool)`

GetBotChallengeCompleteTokensCheckedOk returns a tuple with the BotChallengeCompleteTokensChecked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotChallengeCompleteTokensChecked

`func (o *RealtimeMeasurements) SetBotChallengeCompleteTokensChecked(v int32)`

SetBotChallengeCompleteTokensChecked sets BotChallengeCompleteTokensChecked field to given value.

### HasBotChallengeCompleteTokensChecked

`func (o *RealtimeMeasurements) HasBotChallengeCompleteTokensChecked() bool`

HasBotChallengeCompleteTokensChecked returns a boolean if a field has been set.

### GetBotChallengeCompleteTokensDisabled

`func (o *RealtimeMeasurements) GetBotChallengeCompleteTokensDisabled() int32`

GetBotChallengeCompleteTokensDisabled returns the BotChallengeCompleteTokensDisabled field if non-nil, zero value otherwise.

### GetBotChallengeCompleteTokensDisabledOk

`func (o *RealtimeMeasurements) GetBotChallengeCompleteTokensDisabledOk() (*int32, bool)`

GetBotChallengeCompleteTokensDisabledOk returns a tuple with the BotChallengeCompleteTokensDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotChallengeCompleteTokensDisabled

`func (o *RealtimeMeasurements) SetBotChallengeCompleteTokensDisabled(v int32)`

SetBotChallengeCompleteTokensDisabled sets BotChallengeCompleteTokensDisabled field to given value.

### HasBotChallengeCompleteTokensDisabled

`func (o *RealtimeMeasurements) HasBotChallengeCompleteTokensDisabled() bool`

HasBotChallengeCompleteTokensDisabled returns a boolean if a field has been set.

### GetBotChallengesIssued

`func (o *RealtimeMeasurements) GetBotChallengesIssued() int32`

GetBotChallengesIssued returns the BotChallengesIssued field if non-nil, zero value otherwise.

### GetBotChallengesIssuedOk

`func (o *RealtimeMeasurements) GetBotChallengesIssuedOk() (*int32, bool)`

GetBotChallengesIssuedOk returns a tuple with the BotChallengesIssued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotChallengesIssued

`func (o *RealtimeMeasurements) SetBotChallengesIssued(v int32)`

SetBotChallengesIssued sets BotChallengesIssued field to given value.

### HasBotChallengesIssued

`func (o *RealtimeMeasurements) HasBotChallengesIssued() bool`

HasBotChallengesIssued returns a boolean if a field has been set.

### GetBotChallengesSucceeded

`func (o *RealtimeMeasurements) GetBotChallengesSucceeded() int32`

GetBotChallengesSucceeded returns the BotChallengesSucceeded field if non-nil, zero value otherwise.

### GetBotChallengesSucceededOk

`func (o *RealtimeMeasurements) GetBotChallengesSucceededOk() (*int32, bool)`

GetBotChallengesSucceededOk returns a tuple with the BotChallengesSucceeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotChallengesSucceeded

`func (o *RealtimeMeasurements) SetBotChallengesSucceeded(v int32)`

SetBotChallengesSucceeded sets BotChallengesSucceeded field to given value.

### HasBotChallengesSucceeded

`func (o *RealtimeMeasurements) HasBotChallengesSucceeded() bool`

HasBotChallengesSucceeded returns a boolean if a field has been set.

### GetBotChallengesFailed

`func (o *RealtimeMeasurements) GetBotChallengesFailed() int32`

GetBotChallengesFailed returns the BotChallengesFailed field if non-nil, zero value otherwise.

### GetBotChallengesFailedOk

`func (o *RealtimeMeasurements) GetBotChallengesFailedOk() (*int32, bool)`

GetBotChallengesFailedOk returns a tuple with the BotChallengesFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotChallengesFailed

`func (o *RealtimeMeasurements) SetBotChallengesFailed(v int32)`

SetBotChallengesFailed sets BotChallengesFailed field to given value.

### HasBotChallengesFailed

`func (o *RealtimeMeasurements) HasBotChallengesFailed() bool`

HasBotChallengesFailed returns a boolean if a field has been set.

### GetBotChallengeCompleteTokensIssued

`func (o *RealtimeMeasurements) GetBotChallengeCompleteTokensIssued() int32`

GetBotChallengeCompleteTokensIssued returns the BotChallengeCompleteTokensIssued field if non-nil, zero value otherwise.

### GetBotChallengeCompleteTokensIssuedOk

`func (o *RealtimeMeasurements) GetBotChallengeCompleteTokensIssuedOk() (*int32, bool)`

GetBotChallengeCompleteTokensIssuedOk returns a tuple with the BotChallengeCompleteTokensIssued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotChallengeCompleteTokensIssued

`func (o *RealtimeMeasurements) SetBotChallengeCompleteTokensIssued(v int32)`

SetBotChallengeCompleteTokensIssued sets BotChallengeCompleteTokensIssued field to given value.

### HasBotChallengeCompleteTokensIssued

`func (o *RealtimeMeasurements) HasBotChallengeCompleteTokensIssued() bool`

HasBotChallengeCompleteTokensIssued returns a boolean if a field has been set.

### GetDdosActionDowngrade

`func (o *RealtimeMeasurements) GetDdosActionDowngrade() int32`

GetDdosActionDowngrade returns the DdosActionDowngrade field if non-nil, zero value otherwise.

### GetDdosActionDowngradeOk

`func (o *RealtimeMeasurements) GetDdosActionDowngradeOk() (*int32, bool)`

GetDdosActionDowngradeOk returns a tuple with the DdosActionDowngrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdosActionDowngrade

`func (o *RealtimeMeasurements) SetDdosActionDowngrade(v int32)`

SetDdosActionDowngrade sets DdosActionDowngrade field to given value.

### HasDdosActionDowngrade

`func (o *RealtimeMeasurements) HasDdosActionDowngrade() bool`

HasDdosActionDowngrade returns a boolean if a field has been set.

### GetDdosActionDowngradedConnections

`func (o *RealtimeMeasurements) GetDdosActionDowngradedConnections() int32`

GetDdosActionDowngradedConnections returns the DdosActionDowngradedConnections field if non-nil, zero value otherwise.

### GetDdosActionDowngradedConnectionsOk

`func (o *RealtimeMeasurements) GetDdosActionDowngradedConnectionsOk() (*int32, bool)`

GetDdosActionDowngradedConnectionsOk returns a tuple with the DdosActionDowngradedConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdosActionDowngradedConnections

`func (o *RealtimeMeasurements) SetDdosActionDowngradedConnections(v int32)`

SetDdosActionDowngradedConnections sets DdosActionDowngradedConnections field to given value.

### HasDdosActionDowngradedConnections

`func (o *RealtimeMeasurements) HasDdosActionDowngradedConnections() bool`

HasDdosActionDowngradedConnections returns a boolean if a field has been set.

### GetVclOnComputeHitRequests

`func (o *RealtimeMeasurements) GetVclOnComputeHitRequests() int32`

GetVclOnComputeHitRequests returns the VclOnComputeHitRequests field if non-nil, zero value otherwise.

### GetVclOnComputeHitRequestsOk

`func (o *RealtimeMeasurements) GetVclOnComputeHitRequestsOk() (*int32, bool)`

GetVclOnComputeHitRequestsOk returns a tuple with the VclOnComputeHitRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVclOnComputeHitRequests

`func (o *RealtimeMeasurements) SetVclOnComputeHitRequests(v int32)`

SetVclOnComputeHitRequests sets VclOnComputeHitRequests field to given value.

### HasVclOnComputeHitRequests

`func (o *RealtimeMeasurements) HasVclOnComputeHitRequests() bool`

HasVclOnComputeHitRequests returns a boolean if a field has been set.

### GetVclOnComputeMissRequests

`func (o *RealtimeMeasurements) GetVclOnComputeMissRequests() int32`

GetVclOnComputeMissRequests returns the VclOnComputeMissRequests field if non-nil, zero value otherwise.

### GetVclOnComputeMissRequestsOk

`func (o *RealtimeMeasurements) GetVclOnComputeMissRequestsOk() (*int32, bool)`

GetVclOnComputeMissRequestsOk returns a tuple with the VclOnComputeMissRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVclOnComputeMissRequests

`func (o *RealtimeMeasurements) SetVclOnComputeMissRequests(v int32)`

SetVclOnComputeMissRequests sets VclOnComputeMissRequests field to given value.

### HasVclOnComputeMissRequests

`func (o *RealtimeMeasurements) HasVclOnComputeMissRequests() bool`

HasVclOnComputeMissRequests returns a boolean if a field has been set.

### GetVclOnComputePassRequests

`func (o *RealtimeMeasurements) GetVclOnComputePassRequests() int32`

GetVclOnComputePassRequests returns the VclOnComputePassRequests field if non-nil, zero value otherwise.

### GetVclOnComputePassRequestsOk

`func (o *RealtimeMeasurements) GetVclOnComputePassRequestsOk() (*int32, bool)`

GetVclOnComputePassRequestsOk returns a tuple with the VclOnComputePassRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVclOnComputePassRequests

`func (o *RealtimeMeasurements) SetVclOnComputePassRequests(v int32)`

SetVclOnComputePassRequests sets VclOnComputePassRequests field to given value.

### HasVclOnComputePassRequests

`func (o *RealtimeMeasurements) HasVclOnComputePassRequests() bool`

HasVclOnComputePassRequests returns a boolean if a field has been set.

### GetVclOnComputeErrorRequests

`func (o *RealtimeMeasurements) GetVclOnComputeErrorRequests() int32`

GetVclOnComputeErrorRequests returns the VclOnComputeErrorRequests field if non-nil, zero value otherwise.

### GetVclOnComputeErrorRequestsOk

`func (o *RealtimeMeasurements) GetVclOnComputeErrorRequestsOk() (*int32, bool)`

GetVclOnComputeErrorRequestsOk returns a tuple with the VclOnComputeErrorRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVclOnComputeErrorRequests

`func (o *RealtimeMeasurements) SetVclOnComputeErrorRequests(v int32)`

SetVclOnComputeErrorRequests sets VclOnComputeErrorRequests field to given value.

### HasVclOnComputeErrorRequests

`func (o *RealtimeMeasurements) HasVclOnComputeErrorRequests() bool`

HasVclOnComputeErrorRequests returns a boolean if a field has been set.

### GetVclOnComputeSynthRequests

`func (o *RealtimeMeasurements) GetVclOnComputeSynthRequests() int32`

GetVclOnComputeSynthRequests returns the VclOnComputeSynthRequests field if non-nil, zero value otherwise.

### GetVclOnComputeSynthRequestsOk

`func (o *RealtimeMeasurements) GetVclOnComputeSynthRequestsOk() (*int32, bool)`

GetVclOnComputeSynthRequestsOk returns a tuple with the VclOnComputeSynthRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVclOnComputeSynthRequests

`func (o *RealtimeMeasurements) SetVclOnComputeSynthRequests(v int32)`

SetVclOnComputeSynthRequests sets VclOnComputeSynthRequests field to given value.

### HasVclOnComputeSynthRequests

`func (o *RealtimeMeasurements) HasVclOnComputeSynthRequests() bool`

HasVclOnComputeSynthRequests returns a boolean if a field has been set.

### GetVclOnComputeEdgeHitRequests

`func (o *RealtimeMeasurements) GetVclOnComputeEdgeHitRequests() int32`

GetVclOnComputeEdgeHitRequests returns the VclOnComputeEdgeHitRequests field if non-nil, zero value otherwise.

### GetVclOnComputeEdgeHitRequestsOk

`func (o *RealtimeMeasurements) GetVclOnComputeEdgeHitRequestsOk() (*int32, bool)`

GetVclOnComputeEdgeHitRequestsOk returns a tuple with the VclOnComputeEdgeHitRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVclOnComputeEdgeHitRequests

`func (o *RealtimeMeasurements) SetVclOnComputeEdgeHitRequests(v int32)`

SetVclOnComputeEdgeHitRequests sets VclOnComputeEdgeHitRequests field to given value.

### HasVclOnComputeEdgeHitRequests

`func (o *RealtimeMeasurements) HasVclOnComputeEdgeHitRequests() bool`

HasVclOnComputeEdgeHitRequests returns a boolean if a field has been set.

### GetVclOnComputeEdgeMissRequests

`func (o *RealtimeMeasurements) GetVclOnComputeEdgeMissRequests() int32`

GetVclOnComputeEdgeMissRequests returns the VclOnComputeEdgeMissRequests field if non-nil, zero value otherwise.

### GetVclOnComputeEdgeMissRequestsOk

`func (o *RealtimeMeasurements) GetVclOnComputeEdgeMissRequestsOk() (*int32, bool)`

GetVclOnComputeEdgeMissRequestsOk returns a tuple with the VclOnComputeEdgeMissRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVclOnComputeEdgeMissRequests

`func (o *RealtimeMeasurements) SetVclOnComputeEdgeMissRequests(v int32)`

SetVclOnComputeEdgeMissRequests sets VclOnComputeEdgeMissRequests field to given value.

### HasVclOnComputeEdgeMissRequests

`func (o *RealtimeMeasurements) HasVclOnComputeEdgeMissRequests() bool`

HasVclOnComputeEdgeMissRequests returns a boolean if a field has been set.

### GetAllHitRequests

`func (o *RealtimeMeasurements) GetAllHitRequests() int32`

GetAllHitRequests returns the AllHitRequests field if non-nil, zero value otherwise.

### GetAllHitRequestsOk

`func (o *RealtimeMeasurements) GetAllHitRequestsOk() (*int32, bool)`

GetAllHitRequestsOk returns a tuple with the AllHitRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllHitRequests

`func (o *RealtimeMeasurements) SetAllHitRequests(v int32)`

SetAllHitRequests sets AllHitRequests field to given value.

### HasAllHitRequests

`func (o *RealtimeMeasurements) HasAllHitRequests() bool`

HasAllHitRequests returns a boolean if a field has been set.

### GetAllMissRequests

`func (o *RealtimeMeasurements) GetAllMissRequests() int32`

GetAllMissRequests returns the AllMissRequests field if non-nil, zero value otherwise.

### GetAllMissRequestsOk

`func (o *RealtimeMeasurements) GetAllMissRequestsOk() (*int32, bool)`

GetAllMissRequestsOk returns a tuple with the AllMissRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllMissRequests

`func (o *RealtimeMeasurements) SetAllMissRequests(v int32)`

SetAllMissRequests sets AllMissRequests field to given value.

### HasAllMissRequests

`func (o *RealtimeMeasurements) HasAllMissRequests() bool`

HasAllMissRequests returns a boolean if a field has been set.

### GetAllPassRequests

`func (o *RealtimeMeasurements) GetAllPassRequests() int32`

GetAllPassRequests returns the AllPassRequests field if non-nil, zero value otherwise.

### GetAllPassRequestsOk

`func (o *RealtimeMeasurements) GetAllPassRequestsOk() (*int32, bool)`

GetAllPassRequestsOk returns a tuple with the AllPassRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllPassRequests

`func (o *RealtimeMeasurements) SetAllPassRequests(v int32)`

SetAllPassRequests sets AllPassRequests field to given value.

### HasAllPassRequests

`func (o *RealtimeMeasurements) HasAllPassRequests() bool`

HasAllPassRequests returns a boolean if a field has been set.

### GetAllErrorRequests

`func (o *RealtimeMeasurements) GetAllErrorRequests() int32`

GetAllErrorRequests returns the AllErrorRequests field if non-nil, zero value otherwise.

### GetAllErrorRequestsOk

`func (o *RealtimeMeasurements) GetAllErrorRequestsOk() (*int32, bool)`

GetAllErrorRequestsOk returns a tuple with the AllErrorRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllErrorRequests

`func (o *RealtimeMeasurements) SetAllErrorRequests(v int32)`

SetAllErrorRequests sets AllErrorRequests field to given value.

### HasAllErrorRequests

`func (o *RealtimeMeasurements) HasAllErrorRequests() bool`

HasAllErrorRequests returns a boolean if a field has been set.

### GetAllSynthRequests

`func (o *RealtimeMeasurements) GetAllSynthRequests() int32`

GetAllSynthRequests returns the AllSynthRequests field if non-nil, zero value otherwise.

### GetAllSynthRequestsOk

`func (o *RealtimeMeasurements) GetAllSynthRequestsOk() (*int32, bool)`

GetAllSynthRequestsOk returns a tuple with the AllSynthRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllSynthRequests

`func (o *RealtimeMeasurements) SetAllSynthRequests(v int32)`

SetAllSynthRequests sets AllSynthRequests field to given value.

### HasAllSynthRequests

`func (o *RealtimeMeasurements) HasAllSynthRequests() bool`

HasAllSynthRequests returns a boolean if a field has been set.

### GetAllEdgeHitRequests

`func (o *RealtimeMeasurements) GetAllEdgeHitRequests() int32`

GetAllEdgeHitRequests returns the AllEdgeHitRequests field if non-nil, zero value otherwise.

### GetAllEdgeHitRequestsOk

`func (o *RealtimeMeasurements) GetAllEdgeHitRequestsOk() (*int32, bool)`

GetAllEdgeHitRequestsOk returns a tuple with the AllEdgeHitRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllEdgeHitRequests

`func (o *RealtimeMeasurements) SetAllEdgeHitRequests(v int32)`

SetAllEdgeHitRequests sets AllEdgeHitRequests field to given value.

### HasAllEdgeHitRequests

`func (o *RealtimeMeasurements) HasAllEdgeHitRequests() bool`

HasAllEdgeHitRequests returns a boolean if a field has been set.

### GetAllEdgeMissRequests

`func (o *RealtimeMeasurements) GetAllEdgeMissRequests() int32`

GetAllEdgeMissRequests returns the AllEdgeMissRequests field if non-nil, zero value otherwise.

### GetAllEdgeMissRequestsOk

`func (o *RealtimeMeasurements) GetAllEdgeMissRequestsOk() (*int32, bool)`

GetAllEdgeMissRequestsOk returns a tuple with the AllEdgeMissRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllEdgeMissRequests

`func (o *RealtimeMeasurements) SetAllEdgeMissRequests(v int32)`

SetAllEdgeMissRequests sets AllEdgeMissRequests field to given value.

### HasAllEdgeMissRequests

`func (o *RealtimeMeasurements) HasAllEdgeMissRequests() bool`

HasAllEdgeMissRequests returns a boolean if a field has been set.

### GetAllStatus1xx

`func (o *RealtimeMeasurements) GetAllStatus1xx() int32`

GetAllStatus1xx returns the AllStatus1xx field if non-nil, zero value otherwise.

### GetAllStatus1xxOk

`func (o *RealtimeMeasurements) GetAllStatus1xxOk() (*int32, bool)`

GetAllStatus1xxOk returns a tuple with the AllStatus1xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllStatus1xx

`func (o *RealtimeMeasurements) SetAllStatus1xx(v int32)`

SetAllStatus1xx sets AllStatus1xx field to given value.

### HasAllStatus1xx

`func (o *RealtimeMeasurements) HasAllStatus1xx() bool`

HasAllStatus1xx returns a boolean if a field has been set.

### GetAllStatus2xx

`func (o *RealtimeMeasurements) GetAllStatus2xx() int32`

GetAllStatus2xx returns the AllStatus2xx field if non-nil, zero value otherwise.

### GetAllStatus2xxOk

`func (o *RealtimeMeasurements) GetAllStatus2xxOk() (*int32, bool)`

GetAllStatus2xxOk returns a tuple with the AllStatus2xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllStatus2xx

`func (o *RealtimeMeasurements) SetAllStatus2xx(v int32)`

SetAllStatus2xx sets AllStatus2xx field to given value.

### HasAllStatus2xx

`func (o *RealtimeMeasurements) HasAllStatus2xx() bool`

HasAllStatus2xx returns a boolean if a field has been set.

### GetAllStatus3xx

`func (o *RealtimeMeasurements) GetAllStatus3xx() int32`

GetAllStatus3xx returns the AllStatus3xx field if non-nil, zero value otherwise.

### GetAllStatus3xxOk

`func (o *RealtimeMeasurements) GetAllStatus3xxOk() (*int32, bool)`

GetAllStatus3xxOk returns a tuple with the AllStatus3xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllStatus3xx

`func (o *RealtimeMeasurements) SetAllStatus3xx(v int32)`

SetAllStatus3xx sets AllStatus3xx field to given value.

### HasAllStatus3xx

`func (o *RealtimeMeasurements) HasAllStatus3xx() bool`

HasAllStatus3xx returns a boolean if a field has been set.

### GetAllStatus4xx

`func (o *RealtimeMeasurements) GetAllStatus4xx() int32`

GetAllStatus4xx returns the AllStatus4xx field if non-nil, zero value otherwise.

### GetAllStatus4xxOk

`func (o *RealtimeMeasurements) GetAllStatus4xxOk() (*int32, bool)`

GetAllStatus4xxOk returns a tuple with the AllStatus4xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllStatus4xx

`func (o *RealtimeMeasurements) SetAllStatus4xx(v int32)`

SetAllStatus4xx sets AllStatus4xx field to given value.

### HasAllStatus4xx

`func (o *RealtimeMeasurements) HasAllStatus4xx() bool`

HasAllStatus4xx returns a boolean if a field has been set.

### GetAllStatus5xx

`func (o *RealtimeMeasurements) GetAllStatus5xx() int32`

GetAllStatus5xx returns the AllStatus5xx field if non-nil, zero value otherwise.

### GetAllStatus5xxOk

`func (o *RealtimeMeasurements) GetAllStatus5xxOk() (*int32, bool)`

GetAllStatus5xxOk returns a tuple with the AllStatus5xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllStatus5xx

`func (o *RealtimeMeasurements) SetAllStatus5xx(v int32)`

SetAllStatus5xx sets AllStatus5xx field to given value.

### HasAllStatus5xx

`func (o *RealtimeMeasurements) HasAllStatus5xx() bool`

HasAllStatus5xx returns a boolean if a field has been set.

### GetOriginOffload

`func (o *RealtimeMeasurements) GetOriginOffload() float32`

GetOriginOffload returns the OriginOffload field if non-nil, zero value otherwise.

### GetOriginOffloadOk

`func (o *RealtimeMeasurements) GetOriginOffloadOk() (*float32, bool)`

GetOriginOffloadOk returns a tuple with the OriginOffload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginOffload

`func (o *RealtimeMeasurements) SetOriginOffload(v float32)`

SetOriginOffload sets OriginOffload field to given value.

### HasOriginOffload

`func (o *RealtimeMeasurements) HasOriginOffload() bool`

HasOriginOffload returns a boolean if a field has been set.

### GetRequestDeniedGetHeadBody

`func (o *RealtimeMeasurements) GetRequestDeniedGetHeadBody() int32`

GetRequestDeniedGetHeadBody returns the RequestDeniedGetHeadBody field if non-nil, zero value otherwise.

### GetRequestDeniedGetHeadBodyOk

`func (o *RealtimeMeasurements) GetRequestDeniedGetHeadBodyOk() (*int32, bool)`

GetRequestDeniedGetHeadBodyOk returns a tuple with the RequestDeniedGetHeadBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestDeniedGetHeadBody

`func (o *RealtimeMeasurements) SetRequestDeniedGetHeadBody(v int32)`

SetRequestDeniedGetHeadBody sets RequestDeniedGetHeadBody field to given value.

### HasRequestDeniedGetHeadBody

`func (o *RealtimeMeasurements) HasRequestDeniedGetHeadBody() bool`

HasRequestDeniedGetHeadBody returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
