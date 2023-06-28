# HistoricalFieldResultsAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Requests** | Pointer to **int32** | Number of requests processed. | [optional] 
**Hits** | Pointer to **int32** | Number of cache hits. | [optional] 
**HitsTime** | Pointer to **float32** | Total amount of time spent processing cache hits (in seconds). | [optional] 
**Miss** | Pointer to **int32** | Number of cache misses. | [optional] 
**MissTime** | Pointer to **float32** | Total amount of time spent processing cache misses (in seconds). | [optional] 
**Pass** | Pointer to **int32** | Number of requests that passed through the CDN without being cached. | [optional] 
**PassTime** | Pointer to **float32** | Total amount of time spent processing cache passes (in seconds). | [optional] 
**Errors** | Pointer to **int32** | Number of cache errors. | [optional] 
**Restarts** | Pointer to **int32** | Number of restarts performed. | [optional] 
**HitRatio** | Pointer to **NullableFloat32** | Ratio of cache hits to cache misses (between 0 and 1). | [optional] 
**Bandwidth** | Pointer to **int32** | Total bytes delivered (`resp_header_bytes` + `resp_body_bytes` + `bereq_header_bytes` + `bereq_body_bytes` + `compute_resp_header_bytes` + `compute_resp_body_bytes` + `compute_bereq_header_bytes` + `compute_bereq_body_bytes` + `websocket_resp_header_bytes` + `websocket_resp_body_bytes` + `websocket_bereq_header_bytes` + `websocket_bereq_body_bytes` + `fanout_resp_header_bytes` + `fanout_resp_body_bytes` + `fanout_bereq_header_bytes` + `fanout_bereq_body_bytes`). | [optional] 
**BodySize** | Pointer to **int32** | Total body bytes delivered (alias for resp_body_bytes). | [optional] 
**HeaderSize** | Pointer to **int32** | Total header bytes delivered (alias for resp_header_bytes). | [optional] 
**ReqBodyBytes** | Pointer to **int32** | Total body bytes received. | [optional] 
**ReqHeaderBytes** | Pointer to **int32** | Total header bytes received. | [optional] 
**RespBodyBytes** | Pointer to **int32** | Total body bytes delivered (edge_resp_body_bytes + shield_resp_body_bytes). | [optional] 
**RespHeaderBytes** | Pointer to **int32** | Total header bytes delivered (edge_resp_header_bytes + shield_resp_header_bytes). | [optional] 
**BereqBodyBytes** | Pointer to **int32** | Total body bytes sent to origin. | [optional] 
**BereqHeaderBytes** | Pointer to **int32** | Total header bytes sent to origin. | [optional] 
**Uncacheable** | Pointer to **int32** | Number of requests that were designated uncachable. | [optional] 
**Pipe** | Pointer to **int32** | Optional. Pipe operations performed (legacy feature). | [optional] 
**Synth** | Pointer to **int32** | Number of requests that returned a synthetic response (i.e., response objects created with the `synthetic` VCL statement). | [optional] 
**TLS** | Pointer to **int32** | Number of requests that were received over TLS. | [optional] 
**TLSV10** | Pointer to **int32** | Number of requests received over TLS 1.0. | [optional] 
**TLSV11** | Pointer to **int32** | Number of requests received over TLS 1.1. | [optional] 
**TLSV12** | Pointer to **int32** | Number of requests received over TLS 1.2. | [optional] 
**TLSV13** | Pointer to **int32** | Number of requests received over TLS 1.3. | [optional] 
**EdgeRequests** | Pointer to **int32** | Number of requests sent by end users to Fastly. | [optional] 
**EdgeRespHeaderBytes** | Pointer to **int32** | Total header bytes delivered from Fastly to the end user. | [optional] 
**EdgeRespBodyBytes** | Pointer to **int32** | Total body bytes delivered from Fastly to the end user. | [optional] 
**EdgeHitRequests** | Pointer to **int32** | Number of requests sent by end users to Fastly that resulted in a hit at the edge. | [optional] 
**EdgeMissRequests** | Pointer to **int32** | Number of requests sent by end users to Fastly that resulted in a miss at the edge. | [optional] 
**OriginFetches** | Pointer to **int32** | Number of requests sent to origin. | [optional] 
**OriginFetchHeaderBytes** | Pointer to **int32** | Total request header bytes sent to origin. | [optional] 
**OriginFetchBodyBytes** | Pointer to **int32** | Total request body bytes sent to origin. | [optional] 
**OriginFetchRespHeaderBytes** | Pointer to **int32** | Total header bytes received from origin. | [optional] 
**OriginFetchRespBodyBytes** | Pointer to **int32** | Total body bytes received from origin. | [optional] 
**OriginRevalidations** | Pointer to **int32** | Number of responses received from origin with a `304` status code in response to an `If-Modified-Since` or `If-None-Match` request. Under regular scenarios, a revalidation will imply a cache hit. However, if using Fastly Image Optimizer or segmented caching this may result in a cache miss. | [optional] 
**OriginCacheFetches** | Pointer to **int32** | The total number of completed requests made to backends (origins) that returned cacheable content. | [optional] 
**Shield** | Pointer to **int32** | Number of requests from edge to the shield POP. | [optional] 
**ShieldRespBodyBytes** | Pointer to **int32** | Total body bytes delivered via a shield. | [optional] 
**ShieldRespHeaderBytes** | Pointer to **int32** | Total header bytes delivered via a shield. | [optional] 
**ShieldFetches** | Pointer to **int32** | Number of requests made from one Fastly POP to another, as part of shielding. | [optional] 
**ShieldFetchHeaderBytes** | Pointer to **int32** | Total request header bytes sent to a shield. | [optional] 
**ShieldFetchBodyBytes** | Pointer to **int32** | Total request body bytes sent to a shield. | [optional] 
**ShieldFetchRespHeaderBytes** | Pointer to **int32** | Total response header bytes sent from a shield to the edge. | [optional] 
**ShieldFetchRespBodyBytes** | Pointer to **int32** | Total response body bytes sent from a shield to the edge. | [optional] 
**ShieldRevalidations** | Pointer to **int32** | Number of responses received from origin with a `304` status code, in response to an `If-Modified-Since` or `If-None-Match` request to a shield. Under regular scenarios, a revalidation will imply a cache hit. However, if using segmented caching this may result in a cache miss. | [optional] 
**ShieldCacheFetches** | Pointer to **int32** | The total number of completed requests made to shields that returned cacheable content. | [optional] 
**Ipv6** | Pointer to **int32** | Number of requests that were received over IPv6. | [optional] 
**Otfp** | Pointer to **int32** | Number of responses that came from the Fastly On-the-Fly Packaging service for video-on-demand. | [optional] 
**OtfpRespBodyBytes** | Pointer to **int32** | Total body bytes delivered from the Fastly On-the-Fly Packaging service for video-on-demand. | [optional] 
**OtfpRespHeaderBytes** | Pointer to **int32** | Total header bytes delivered from the Fastly On-the-Fly Packaging service for video-on-demand. | [optional] 
**OtfpShieldRespBodyBytes** | Pointer to **int32** | Total body bytes delivered via a shield for the Fastly On-the-Fly Packaging service for video-on-demand. | [optional] 
**OtfpShieldRespHeaderBytes** | Pointer to **int32** | Total header bytes delivered via a shield for the Fastly On-the-Fly Packaging service for video-on-demand. | [optional] 
**OtfpManifests** | Pointer to **int32** | Number of responses that were manifest files from the Fastly On-the-Fly Packaging service for video-on-demand. | [optional] 
**OtfpDeliverTime** | Pointer to **float32** | Total amount of time spent delivering a response from the Fastly On-the-Fly Packaging service for video-on-demand (in seconds). | [optional] 
**OtfpShieldTime** | Pointer to **float32** | Total amount of time spent delivering a response via a shield from the Fastly On-the-Fly Packaging service for video-on-demand (in seconds). | [optional] 
**Video** | Pointer to **int32** | Number of responses with the video segment or video manifest MIME type (i.e., application/x-mpegurl, application/vnd.apple.mpegurl, application/f4m, application/dash+xml, application/vnd.ms-sstr+xml, ideo/mp2t, audio/aac, video/f4f, video/x-flv, video/mp4, audio/mp4). | [optional] 
**Pci** | Pointer to **int32** | Number of responses with the PCI flag turned on. | [optional] 
**Log** | Pointer to **int32** | Number of log lines sent. | [optional] 
**LogBytes** | Pointer to **int32** | Total log bytes sent. | [optional] 
**HTTP2** | Pointer to **int32** | Number of requests received over HTTP/2. | [optional] 
**HTTP3** | Pointer to **int32** | Number of requests received over HTTP/3. | [optional] 
**WafLogged** | Pointer to **int32** | Number of requests that triggered a WAF rule and were logged. | [optional] 
**WafBlocked** | Pointer to **int32** | Number of requests that triggered a WAF rule and were blocked. | [optional] 
**WafPassed** | Pointer to **int32** | Number of requests that triggered a WAF rule and were passed. | [optional] 
**AttackReqBodyBytes** | Pointer to **int32** | Total body bytes received from requests that triggered a WAF rule. | [optional] 
**AttackReqHeaderBytes** | Pointer to **int32** | Total header bytes received from requests that triggered a WAF rule. | [optional] 
**AttackLoggedReqBodyBytes** | Pointer to **int32** | Total body bytes received from requests that triggered a WAF rule that was logged. | [optional] 
**AttackLoggedReqHeaderBytes** | Pointer to **int32** | Total header bytes received from requests that triggered a WAF rule that was logged. | [optional] 
**AttackBlockedReqBodyBytes** | Pointer to **int32** | Total body bytes received from requests that triggered a WAF rule that was blocked. | [optional] 
**AttackBlockedReqHeaderBytes** | Pointer to **int32** | Total header bytes received from requests that triggered a WAF rule that was blocked. | [optional] 
**AttackPassedReqBodyBytes** | Pointer to **int32** | Total body bytes received from requests that triggered a WAF rule that was passed. | [optional] 
**AttackPassedReqHeaderBytes** | Pointer to **int32** | Total header bytes received from requests that triggered a WAF rule that was passed. | [optional] 
**AttackRespSynthBytes** | Pointer to **int32** | Total bytes delivered for requests that triggered a WAF rule and returned a synthetic response. | [optional] 
**Imgopto** | Pointer to **int32** | Number of responses that came from the Fastly Image Optimizer service. If the service receives 10 requests for an image, this stat will be 10 regardless of how many times the image was transformed. | [optional] 
**ImgoptoRespBodyBytes** | Pointer to **int32** | Total body bytes delivered from the Fastly Image Optimizer service, including shield traffic. | [optional] 
**ImgoptoRespHeaderBytes** | Pointer to **int32** | Total header bytes delivered from the Fastly Image Optimizer service, including shield traffic. | [optional] 
**ImgoptoShieldRespBodyBytes** | Pointer to **int32** | Total body bytes delivered via a shield from the Fastly Image Optimizer service. | [optional] 
**ImgoptoShieldRespHeaderBytes** | Pointer to **int32** | Total header bytes delivered via a shield from the Fastly Image Optimizer service. | [optional] 
**Imgvideo** | Pointer to **int32** | Number of video responses that came from the Fastly Image Optimizer service. | [optional] 
**ImgvideoFrames** | Pointer to **int32** | Number of video frames that came from the Fastly Image Optimizer service. A video frame is an individual image within a sequence of video. | [optional] 
**ImgvideoRespHeaderBytes** | Pointer to **int32** | Total header bytes of video delivered from the Fastly Image Optimizer service. | [optional] 
**ImgvideoRespBodyBytes** | Pointer to **int32** | Total body bytes of video delivered from the Fastly Image Optimizer service. | [optional] 
**ImgvideoShieldRespHeaderBytes** | Pointer to **int32** | Total header bytes of video delivered via a shield from the Fastly Image Optimizer service. | [optional] 
**ImgvideoShieldRespBodyBytes** | Pointer to **int32** | Total body bytes of video delivered via a shield from the Fastly Image Optimizer service. | [optional] 
**ImgvideoShield** | Pointer to **int32** | Number of video responses delivered via a shield that came from the Fastly Image Optimizer service. | [optional] 
**ImgvideoShieldFrames** | Pointer to **int32** | Number of video frames delivered via a shield that came from the Fastly Image Optimizer service. A video frame is an individual image within a sequence of video. | [optional] 
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
**Status1xx** | Pointer to **int32** | Number of \&quot;Informational\&quot; category status codes delivered. | [optional] 
**Status2xx** | Pointer to **int32** | Number of \&quot;Success\&quot; status codes delivered. | [optional] 
**Status3xx** | Pointer to **int32** | Number of \&quot;Redirection\&quot; codes delivered. | [optional] 
**Status4xx** | Pointer to **int32** | Number of \&quot;Client Error\&quot; codes delivered. | [optional] 
**Status5xx** | Pointer to **int32** | Number of \&quot;Server Error\&quot; codes delivered. | [optional] 
**ObjectSize1k** | Pointer to **int32** | Number of objects served that were under 1KB in size. | [optional] 
**ObjectSize10k** | Pointer to **int32** | Number of objects served that were between 1KB and 10KB in size. | [optional] 
**ObjectSize100k** | Pointer to **int32** | Number of objects served that were between 10KB and 100KB in size. | [optional] 
**ObjectSize1m** | Pointer to **int32** | Number of objects served that were between 100KB and 1MB in size. | [optional] 
**ObjectSize10m** | Pointer to **int32** | Number of objects served that were between 1MB and 10MB in size. | [optional] 
**ObjectSize100m** | Pointer to **int32** | Number of objects served that were between 10MB and 100MB in size. | [optional] 
**ObjectSize1g** | Pointer to **int32** | Number of objects served that were between 100MB and 1GB in size. | [optional] 
**RecvSubTime** | Pointer to **float32** | Time spent inside the `vcl_recv` Varnish subroutine (in seconds). | [optional] 
**RecvSubCount** | Pointer to **int32** | Number of executions of the `vcl_recv` Varnish subroutine. | [optional] 
**HashSubTime** | Pointer to **float32** | Time spent inside the `vcl_hash` Varnish subroutine (in seconds). | [optional] 
**HashSubCount** | Pointer to **int32** | Number of executions of the `vcl_hash` Varnish subroutine. | [optional] 
**MissSubTime** | Pointer to **float32** | Time spent inside the `vcl_miss` Varnish subroutine (in seconds). | [optional] 
**MissSubCount** | Pointer to **int32** | Number of executions of the `vcl_miss` Varnish subroutine. | [optional] 
**FetchSubTime** | Pointer to **float32** | Time spent inside the `vcl_fetch` Varnish subroutine (in seconds). | [optional] 
**FetchSubCount** | Pointer to **int32** | Number of executions of the `vcl_fetch` Varnish subroutine. | [optional] 
**PassSubTime** | Pointer to **float32** | Time spent inside the `vcl_pass` Varnish subroutine (in seconds). | [optional] 
**PassSubCount** | Pointer to **int32** | Number of executions of the `vcl_pass` Varnish subroutine. | [optional] 
**PipeSubTime** | Pointer to **float32** | Time spent inside the `vcl_pipe` Varnish subroutine (in seconds). | [optional] 
**PipeSubCount** | Pointer to **int32** | Number of executions of the `vcl_pipe` Varnish subroutine. | [optional] 
**DeliverSubTime** | Pointer to **float32** | Time spent inside the `vcl_deliver` Varnish subroutine (in seconds). | [optional] 
**DeliverSubCount** | Pointer to **int32** | Number of executions of the `vcl_deliver` Varnish subroutine. | [optional] 
**ErrorSubTime** | Pointer to **float32** | Time spent inside the `vcl_error` Varnish subroutine (in seconds). | [optional] 
**ErrorSubCount** | Pointer to **int32** | Number of executions of the `vcl_error` Varnish subroutine. | [optional] 
**HitSubTime** | Pointer to **float32** | Time spent inside the `vcl_hit` Varnish subroutine (in seconds). | [optional] 
**HitSubCount** | Pointer to **int32** | Number of executions of the `vcl_hit` Varnish subroutine. | [optional] 
**PrehashSubTime** | Pointer to **float32** | Time spent inside the `vcl_prehash` Varnish subroutine (in seconds). | [optional] 
**PrehashSubCount** | Pointer to **int32** | Number of executions of the `vcl_prehash` Varnish subroutine. | [optional] 
**PredeliverSubTime** | Pointer to **float32** | Time spent inside the `vcl_predeliver` Varnish subroutine (in seconds). | [optional] 
**PredeliverSubCount** | Pointer to **int32** | Number of executions of the `vcl_predeliver` Varnish subroutine. | [optional] 
**TLSHandshakeSentBytes** | Pointer to **int32** | Number of bytes transferred during TLS handshake. | [optional] 
**HitRespBodyBytes** | Pointer to **int32** | Total body bytes delivered for cache hits. | [optional] 
**MissRespBodyBytes** | Pointer to **int32** | Total body bytes delivered for cache misses. | [optional] 
**PassRespBodyBytes** | Pointer to **int32** | Total body bytes delivered for cache passes. | [optional] 
**SegblockOriginFetches** | Pointer to **int32** | Number of `Range` requests to origin for segments of resources when using segmented caching. | [optional] 
**SegblockShieldFetches** | Pointer to **int32** | Number of `Range` requests to a shield for segments of resources when using segmented caching. | [optional] 
**ComputeRequests** | Pointer to **int32** | The total number of requests that were received for your service by Fastly. | [optional] 
**ComputeRequestTimeMs** | Pointer to **float32** | The total, actual amount of time used to process your requests, including active CPU time (in milliseconds). | [optional] 
**ComputeRequestTimeBilledMs** | Pointer to **float32** | The total amount of request processing time you will be billed for, measured in 50 millisecond increments. | [optional] 
**ComputeRAMUsed** | Pointer to **int32** | The amount of RAM used for your service by Fastly (in bytes). | [optional] 
**ComputeExecutionTimeMs** | Pointer to **float32** | The amount of active CPU time used to process your requests (in milliseconds). | [optional] 
**ComputeReqHeaderBytes** | Pointer to **int32** | Total header bytes received by Compute@Edge. | [optional] 
**ComputeReqBodyBytes** | Pointer to **int32** | Total body bytes received by Compute@Edge. | [optional] 
**ComputeRespHeaderBytes** | Pointer to **int32** | Total header bytes sent from Compute@Edge to end user. | [optional] 
**ComputeRespBodyBytes** | Pointer to **int32** | Total body bytes sent from Compute@Edge to end user. | [optional] 
**ComputeRespStatus1xx** | Pointer to **int32** | Number of \&quot;Informational\&quot; category status codes delivered by Compute@Edge. | [optional] 
**ComputeRespStatus2xx** | Pointer to **int32** | Number of \&quot;Success\&quot; category status codes delivered by Compute@Edge. | [optional] 
**ComputeRespStatus3xx** | Pointer to **int32** | Number of \&quot;Redirection\&quot; category status codes delivered by Compute@Edge. | [optional] 
**ComputeRespStatus4xx** | Pointer to **int32** | Number of \&quot;Client Error\&quot; category status codes delivered by Compute@Edge. | [optional] 
**ComputeRespStatus5xx** | Pointer to **int32** | Number of \&quot;Server Error\&quot; category status codes delivered by Compute@Edge. | [optional] 
**ComputeBereqHeaderBytes** | Pointer to **int32** | Total header bytes sent to backends (origins) by Compute@Edge. | [optional] 
**ComputeBereqBodyBytes** | Pointer to **int32** | Total body bytes sent to backends (origins) by Compute@Edge. | [optional] 
**ComputeBerespHeaderBytes** | Pointer to **int32** | Total header bytes received from backends (origins) by Compute@Edge. | [optional] 
**ComputeBerespBodyBytes** | Pointer to **int32** | Total body bytes received from backends (origins) by Compute@Edge. | [optional] 
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
**WebsocketRespBodyBytes** | Pointer to **int32** | Total message content bytes sent to end users over passthrough WebSocket connections. | [optional] 
**WebsocketBereqHeaderBytes** | Pointer to **int32** | Total header bytes sent to backends over passthrough WebSocket connections. | [optional] 
**WebsocketBereqBodyBytes** | Pointer to **int32** | Total message content bytes sent to backends over passthrough WebSocket connections. | [optional] 
**WebsocketBerespHeaderBytes** | Pointer to **int32** | Total header bytes received from backends over passthrough WebSocket connections. | [optional] 
**WebsocketBerespBodyBytes** | Pointer to **int32** | Total message content bytes received from backends over passthrough WebSocket connections. | [optional] 
**WebsocketConnTimeMs** | Pointer to **int32** | Total duration of passthrough WebSocket connections with end users. | [optional] 
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
**ServiceID** | Pointer to [**ReadOnlyIDService**](ReadOnlyIDService.md) |  | [optional] 
**StartTime** | Pointer to **int32** |  | [optional] 

## Methods

### NewHistoricalFieldResultsAttributes

`func NewHistoricalFieldResultsAttributes() *HistoricalFieldResultsAttributes`

NewHistoricalFieldResultsAttributes instantiates a new HistoricalFieldResultsAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoricalFieldResultsAttributesWithDefaults

`func NewHistoricalFieldResultsAttributesWithDefaults() *HistoricalFieldResultsAttributes`

NewHistoricalFieldResultsAttributesWithDefaults instantiates a new HistoricalFieldResultsAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequests

`func (o *HistoricalFieldResultsAttributes) GetRequests() int32`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *HistoricalFieldResultsAttributes) GetRequestsOk() (*int32, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *HistoricalFieldResultsAttributes) SetRequests(v int32)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *HistoricalFieldResultsAttributes) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### GetHits

`func (o *HistoricalFieldResultsAttributes) GetHits() int32`

GetHits returns the Hits field if non-nil, zero value otherwise.

### GetHitsOk

`func (o *HistoricalFieldResultsAttributes) GetHitsOk() (*int32, bool)`

GetHitsOk returns a tuple with the Hits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHits

`func (o *HistoricalFieldResultsAttributes) SetHits(v int32)`

SetHits sets Hits field to given value.

### HasHits

`func (o *HistoricalFieldResultsAttributes) HasHits() bool`

HasHits returns a boolean if a field has been set.

### GetHitsTime

`func (o *HistoricalFieldResultsAttributes) GetHitsTime() float32`

GetHitsTime returns the HitsTime field if non-nil, zero value otherwise.

### GetHitsTimeOk

`func (o *HistoricalFieldResultsAttributes) GetHitsTimeOk() (*float32, bool)`

GetHitsTimeOk returns a tuple with the HitsTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHitsTime

`func (o *HistoricalFieldResultsAttributes) SetHitsTime(v float32)`

SetHitsTime sets HitsTime field to given value.

### HasHitsTime

`func (o *HistoricalFieldResultsAttributes) HasHitsTime() bool`

HasHitsTime returns a boolean if a field has been set.

### GetMiss

`func (o *HistoricalFieldResultsAttributes) GetMiss() int32`

GetMiss returns the Miss field if non-nil, zero value otherwise.

### GetMissOk

`func (o *HistoricalFieldResultsAttributes) GetMissOk() (*int32, bool)`

GetMissOk returns a tuple with the Miss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiss

`func (o *HistoricalFieldResultsAttributes) SetMiss(v int32)`

SetMiss sets Miss field to given value.

### HasMiss

`func (o *HistoricalFieldResultsAttributes) HasMiss() bool`

HasMiss returns a boolean if a field has been set.

### GetMissTime

`func (o *HistoricalFieldResultsAttributes) GetMissTime() float32`

GetMissTime returns the MissTime field if non-nil, zero value otherwise.

### GetMissTimeOk

`func (o *HistoricalFieldResultsAttributes) GetMissTimeOk() (*float32, bool)`

GetMissTimeOk returns a tuple with the MissTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissTime

`func (o *HistoricalFieldResultsAttributes) SetMissTime(v float32)`

SetMissTime sets MissTime field to given value.

### HasMissTime

`func (o *HistoricalFieldResultsAttributes) HasMissTime() bool`

HasMissTime returns a boolean if a field has been set.

### GetPass

`func (o *HistoricalFieldResultsAttributes) GetPass() int32`

GetPass returns the Pass field if non-nil, zero value otherwise.

### GetPassOk

`func (o *HistoricalFieldResultsAttributes) GetPassOk() (*int32, bool)`

GetPassOk returns a tuple with the Pass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPass

`func (o *HistoricalFieldResultsAttributes) SetPass(v int32)`

SetPass sets Pass field to given value.

### HasPass

`func (o *HistoricalFieldResultsAttributes) HasPass() bool`

HasPass returns a boolean if a field has been set.

### GetPassTime

`func (o *HistoricalFieldResultsAttributes) GetPassTime() float32`

GetPassTime returns the PassTime field if non-nil, zero value otherwise.

### GetPassTimeOk

`func (o *HistoricalFieldResultsAttributes) GetPassTimeOk() (*float32, bool)`

GetPassTimeOk returns a tuple with the PassTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassTime

`func (o *HistoricalFieldResultsAttributes) SetPassTime(v float32)`

SetPassTime sets PassTime field to given value.

### HasPassTime

`func (o *HistoricalFieldResultsAttributes) HasPassTime() bool`

HasPassTime returns a boolean if a field has been set.

### GetErrors

`func (o *HistoricalFieldResultsAttributes) GetErrors() int32`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *HistoricalFieldResultsAttributes) GetErrorsOk() (*int32, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *HistoricalFieldResultsAttributes) SetErrors(v int32)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *HistoricalFieldResultsAttributes) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetRestarts

`func (o *HistoricalFieldResultsAttributes) GetRestarts() int32`

GetRestarts returns the Restarts field if non-nil, zero value otherwise.

### GetRestartsOk

`func (o *HistoricalFieldResultsAttributes) GetRestartsOk() (*int32, bool)`

GetRestartsOk returns a tuple with the Restarts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestarts

`func (o *HistoricalFieldResultsAttributes) SetRestarts(v int32)`

SetRestarts sets Restarts field to given value.

### HasRestarts

`func (o *HistoricalFieldResultsAttributes) HasRestarts() bool`

HasRestarts returns a boolean if a field has been set.

### GetHitRatio

`func (o *HistoricalFieldResultsAttributes) GetHitRatio() float32`

GetHitRatio returns the HitRatio field if non-nil, zero value otherwise.

### GetHitRatioOk

`func (o *HistoricalFieldResultsAttributes) GetHitRatioOk() (*float32, bool)`

GetHitRatioOk returns a tuple with the HitRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHitRatio

`func (o *HistoricalFieldResultsAttributes) SetHitRatio(v float32)`

SetHitRatio sets HitRatio field to given value.

### HasHitRatio

`func (o *HistoricalFieldResultsAttributes) HasHitRatio() bool`

HasHitRatio returns a boolean if a field has been set.

### SetHitRatioNil

`func (o *HistoricalFieldResultsAttributes) SetHitRatioNil(b bool)`

 SetHitRatioNil sets the value for HitRatio to be an explicit nil

### UnsetHitRatio
`func (o *HistoricalFieldResultsAttributes) UnsetHitRatio()`

UnsetHitRatio ensures that no value is present for HitRatio, not even an explicit nil
### GetBandwidth

`func (o *HistoricalFieldResultsAttributes) GetBandwidth() int32`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *HistoricalFieldResultsAttributes) GetBandwidthOk() (*int32, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *HistoricalFieldResultsAttributes) SetBandwidth(v int32)`

SetBandwidth sets Bandwidth field to given value.

### HasBandwidth

`func (o *HistoricalFieldResultsAttributes) HasBandwidth() bool`

HasBandwidth returns a boolean if a field has been set.

### GetBodySize

`func (o *HistoricalFieldResultsAttributes) GetBodySize() int32`

GetBodySize returns the BodySize field if non-nil, zero value otherwise.

### GetBodySizeOk

`func (o *HistoricalFieldResultsAttributes) GetBodySizeOk() (*int32, bool)`

GetBodySizeOk returns a tuple with the BodySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodySize

`func (o *HistoricalFieldResultsAttributes) SetBodySize(v int32)`

SetBodySize sets BodySize field to given value.

### HasBodySize

`func (o *HistoricalFieldResultsAttributes) HasBodySize() bool`

HasBodySize returns a boolean if a field has been set.

### GetHeaderSize

`func (o *HistoricalFieldResultsAttributes) GetHeaderSize() int32`

GetHeaderSize returns the HeaderSize field if non-nil, zero value otherwise.

### GetHeaderSizeOk

`func (o *HistoricalFieldResultsAttributes) GetHeaderSizeOk() (*int32, bool)`

GetHeaderSizeOk returns a tuple with the HeaderSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderSize

`func (o *HistoricalFieldResultsAttributes) SetHeaderSize(v int32)`

SetHeaderSize sets HeaderSize field to given value.

### HasHeaderSize

`func (o *HistoricalFieldResultsAttributes) HasHeaderSize() bool`

HasHeaderSize returns a boolean if a field has been set.

### GetReqBodyBytes

`func (o *HistoricalFieldResultsAttributes) GetReqBodyBytes() int32`

GetReqBodyBytes returns the ReqBodyBytes field if non-nil, zero value otherwise.

### GetReqBodyBytesOk

`func (o *HistoricalFieldResultsAttributes) GetReqBodyBytesOk() (*int32, bool)`

GetReqBodyBytesOk returns a tuple with the ReqBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqBodyBytes

`func (o *HistoricalFieldResultsAttributes) SetReqBodyBytes(v int32)`

SetReqBodyBytes sets ReqBodyBytes field to given value.

### HasReqBodyBytes

`func (o *HistoricalFieldResultsAttributes) HasReqBodyBytes() bool`

HasReqBodyBytes returns a boolean if a field has been set.

### GetReqHeaderBytes

`func (o *HistoricalFieldResultsAttributes) GetReqHeaderBytes() int32`

GetReqHeaderBytes returns the ReqHeaderBytes field if non-nil, zero value otherwise.

### GetReqHeaderBytesOk

`func (o *HistoricalFieldResultsAttributes) GetReqHeaderBytesOk() (*int32, bool)`

GetReqHeaderBytesOk returns a tuple with the ReqHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqHeaderBytes

`func (o *HistoricalFieldResultsAttributes) SetReqHeaderBytes(v int32)`

SetReqHeaderBytes sets ReqHeaderBytes field to given value.

### HasReqHeaderBytes

`func (o *HistoricalFieldResultsAttributes) HasReqHeaderBytes() bool`

HasReqHeaderBytes returns a boolean if a field has been set.

### GetRespBodyBytes

`func (o *HistoricalFieldResultsAttributes) GetRespBodyBytes() int32`

GetRespBodyBytes returns the RespBodyBytes field if non-nil, zero value otherwise.

### GetRespBodyBytesOk

`func (o *HistoricalFieldResultsAttributes) GetRespBodyBytesOk() (*int32, bool)`

GetRespBodyBytesOk returns a tuple with the RespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRespBodyBytes

`func (o *HistoricalFieldResultsAttributes) SetRespBodyBytes(v int32)`

SetRespBodyBytes sets RespBodyBytes field to given value.

### HasRespBodyBytes

`func (o *HistoricalFieldResultsAttributes) HasRespBodyBytes() bool`

HasRespBodyBytes returns a boolean if a field has been set.

### GetRespHeaderBytes

`func (o *HistoricalFieldResultsAttributes) GetRespHeaderBytes() int32`

GetRespHeaderBytes returns the RespHeaderBytes field if non-nil, zero value otherwise.

### GetRespHeaderBytesOk

`func (o *HistoricalFieldResultsAttributes) GetRespHeaderBytesOk() (*int32, bool)`

GetRespHeaderBytesOk returns a tuple with the RespHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRespHeaderBytes

`func (o *HistoricalFieldResultsAttributes) SetRespHeaderBytes(v int32)`

SetRespHeaderBytes sets RespHeaderBytes field to given value.

### HasRespHeaderBytes

`func (o *HistoricalFieldResultsAttributes) HasRespHeaderBytes() bool`

HasRespHeaderBytes returns a boolean if a field has been set.

### GetBereqBodyBytes

`func (o *HistoricalFieldResultsAttributes) GetBereqBodyBytes() int32`

GetBereqBodyBytes returns the BereqBodyBytes field if non-nil, zero value otherwise.

### GetBereqBodyBytesOk

`func (o *HistoricalFieldResultsAttributes) GetBereqBodyBytesOk() (*int32, bool)`

GetBereqBodyBytesOk returns a tuple with the BereqBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBereqBodyBytes

`func (o *HistoricalFieldResultsAttributes) SetBereqBodyBytes(v int32)`

SetBereqBodyBytes sets BereqBodyBytes field to given value.

### HasBereqBodyBytes

`func (o *HistoricalFieldResultsAttributes) HasBereqBodyBytes() bool`

HasBereqBodyBytes returns a boolean if a field has been set.

### GetBereqHeaderBytes

`func (o *HistoricalFieldResultsAttributes) GetBereqHeaderBytes() int32`

GetBereqHeaderBytes returns the BereqHeaderBytes field if non-nil, zero value otherwise.

### GetBereqHeaderBytesOk

`func (o *HistoricalFieldResultsAttributes) GetBereqHeaderBytesOk() (*int32, bool)`

GetBereqHeaderBytesOk returns a tuple with the BereqHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBereqHeaderBytes

`func (o *HistoricalFieldResultsAttributes) SetBereqHeaderBytes(v int32)`

SetBereqHeaderBytes sets BereqHeaderBytes field to given value.

### HasBereqHeaderBytes

`func (o *HistoricalFieldResultsAttributes) HasBereqHeaderBytes() bool`

HasBereqHeaderBytes returns a boolean if a field has been set.

### GetUncacheable

`func (o *HistoricalFieldResultsAttributes) GetUncacheable() int32`

GetUncacheable returns the Uncacheable field if non-nil, zero value otherwise.

### GetUncacheableOk

`func (o *HistoricalFieldResultsAttributes) GetUncacheableOk() (*int32, bool)`

GetUncacheableOk returns a tuple with the Uncacheable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUncacheable

`func (o *HistoricalFieldResultsAttributes) SetUncacheable(v int32)`

SetUncacheable sets Uncacheable field to given value.

### HasUncacheable

`func (o *HistoricalFieldResultsAttributes) HasUncacheable() bool`

HasUncacheable returns a boolean if a field has been set.

### GetPipe

`func (o *HistoricalFieldResultsAttributes) GetPipe() int32`

GetPipe returns the Pipe field if non-nil, zero value otherwise.

### GetPipeOk

`func (o *HistoricalFieldResultsAttributes) GetPipeOk() (*int32, bool)`

GetPipeOk returns a tuple with the Pipe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipe

`func (o *HistoricalFieldResultsAttributes) SetPipe(v int32)`

SetPipe sets Pipe field to given value.

### HasPipe

`func (o *HistoricalFieldResultsAttributes) HasPipe() bool`

HasPipe returns a boolean if a field has been set.

### GetSynth

`func (o *HistoricalFieldResultsAttributes) GetSynth() int32`

GetSynth returns the Synth field if non-nil, zero value otherwise.

### GetSynthOk

`func (o *HistoricalFieldResultsAttributes) GetSynthOk() (*int32, bool)`

GetSynthOk returns a tuple with the Synth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynth

`func (o *HistoricalFieldResultsAttributes) SetSynth(v int32)`

SetSynth sets Synth field to given value.

### HasSynth

`func (o *HistoricalFieldResultsAttributes) HasSynth() bool`

HasSynth returns a boolean if a field has been set.

### GetTLS

`func (o *HistoricalFieldResultsAttributes) GetTLS() int32`

GetTLS returns the TLS field if non-nil, zero value otherwise.

### GetTLSOk

`func (o *HistoricalFieldResultsAttributes) GetTLSOk() (*int32, bool)`

GetTLSOk returns a tuple with the TLS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTLS

`func (o *HistoricalFieldResultsAttributes) SetTLS(v int32)`

SetTLS sets TLS field to given value.

### HasTLS

`func (o *HistoricalFieldResultsAttributes) HasTLS() bool`

HasTLS returns a boolean if a field has been set.

### GetTLSV10

`func (o *HistoricalFieldResultsAttributes) GetTLSV10() int32`

GetTLSV10 returns the TLSV10 field if non-nil, zero value otherwise.

### GetTLSV10Ok

`func (o *HistoricalFieldResultsAttributes) GetTLSV10Ok() (*int32, bool)`

GetTLSV10Ok returns a tuple with the TLSV10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTLSV10

`func (o *HistoricalFieldResultsAttributes) SetTLSV10(v int32)`

SetTLSV10 sets TLSV10 field to given value.

### HasTLSV10

`func (o *HistoricalFieldResultsAttributes) HasTLSV10() bool`

HasTLSV10 returns a boolean if a field has been set.

### GetTLSV11

`func (o *HistoricalFieldResultsAttributes) GetTLSV11() int32`

GetTLSV11 returns the TLSV11 field if non-nil, zero value otherwise.

### GetTLSV11Ok

`func (o *HistoricalFieldResultsAttributes) GetTLSV11Ok() (*int32, bool)`

GetTLSV11Ok returns a tuple with the TLSV11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTLSV11

`func (o *HistoricalFieldResultsAttributes) SetTLSV11(v int32)`

SetTLSV11 sets TLSV11 field to given value.

### HasTLSV11

`func (o *HistoricalFieldResultsAttributes) HasTLSV11() bool`

HasTLSV11 returns a boolean if a field has been set.

### GetTLSV12

`func (o *HistoricalFieldResultsAttributes) GetTLSV12() int32`

GetTLSV12 returns the TLSV12 field if non-nil, zero value otherwise.

### GetTLSV12Ok

`func (o *HistoricalFieldResultsAttributes) GetTLSV12Ok() (*int32, bool)`

GetTLSV12Ok returns a tuple with the TLSV12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTLSV12

`func (o *HistoricalFieldResultsAttributes) SetTLSV12(v int32)`

SetTLSV12 sets TLSV12 field to given value.

### HasTLSV12

`func (o *HistoricalFieldResultsAttributes) HasTLSV12() bool`

HasTLSV12 returns a boolean if a field has been set.

### GetTLSV13

`func (o *HistoricalFieldResultsAttributes) GetTLSV13() int32`

GetTLSV13 returns the TLSV13 field if non-nil, zero value otherwise.

### GetTLSV13Ok

`func (o *HistoricalFieldResultsAttributes) GetTLSV13Ok() (*int32, bool)`

GetTLSV13Ok returns a tuple with the TLSV13 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTLSV13

`func (o *HistoricalFieldResultsAttributes) SetTLSV13(v int32)`

SetTLSV13 sets TLSV13 field to given value.

### HasTLSV13

`func (o *HistoricalFieldResultsAttributes) HasTLSV13() bool`

HasTLSV13 returns a boolean if a field has been set.

### GetEdgeRequests

`func (o *HistoricalFieldResultsAttributes) GetEdgeRequests() int32`

GetEdgeRequests returns the EdgeRequests field if non-nil, zero value otherwise.

### GetEdgeRequestsOk

`func (o *HistoricalFieldResultsAttributes) GetEdgeRequestsOk() (*int32, bool)`

GetEdgeRequestsOk returns a tuple with the EdgeRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeRequests

`func (o *HistoricalFieldResultsAttributes) SetEdgeRequests(v int32)`

SetEdgeRequests sets EdgeRequests field to given value.

### HasEdgeRequests

`func (o *HistoricalFieldResultsAttributes) HasEdgeRequests() bool`

HasEdgeRequests returns a boolean if a field has been set.

### GetEdgeRespHeaderBytes

`func (o *HistoricalFieldResultsAttributes) GetEdgeRespHeaderBytes() int32`

GetEdgeRespHeaderBytes returns the EdgeRespHeaderBytes field if non-nil, zero value otherwise.

### GetEdgeRespHeaderBytesOk

`func (o *HistoricalFieldResultsAttributes) GetEdgeRespHeaderBytesOk() (*int32, bool)`

GetEdgeRespHeaderBytesOk returns a tuple with the EdgeRespHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeRespHeaderBytes

`func (o *HistoricalFieldResultsAttributes) SetEdgeRespHeaderBytes(v int32)`

SetEdgeRespHeaderBytes sets EdgeRespHeaderBytes field to given value.

### HasEdgeRespHeaderBytes

`func (o *HistoricalFieldResultsAttributes) HasEdgeRespHeaderBytes() bool`

HasEdgeRespHeaderBytes returns a boolean if a field has been set.

### GetEdgeRespBodyBytes

`func (o *HistoricalFieldResultsAttributes) GetEdgeRespBodyBytes() int32`

GetEdgeRespBodyBytes returns the EdgeRespBodyBytes field if non-nil, zero value otherwise.

### GetEdgeRespBodyBytesOk

`func (o *HistoricalFieldResultsAttributes) GetEdgeRespBodyBytesOk() (*int32, bool)`

GetEdgeRespBodyBytesOk returns a tuple with the EdgeRespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeRespBodyBytes

`func (o *HistoricalFieldResultsAttributes) SetEdgeRespBodyBytes(v int32)`

SetEdgeRespBodyBytes sets EdgeRespBodyBytes field to given value.

### HasEdgeRespBodyBytes

`func (o *HistoricalFieldResultsAttributes) HasEdgeRespBodyBytes() bool`

HasEdgeRespBodyBytes returns a boolean if a field has been set.

### GetEdgeHitRequests

`func (o *HistoricalFieldResultsAttributes) GetEdgeHitRequests() int32`

GetEdgeHitRequests returns the EdgeHitRequests field if non-nil, zero value otherwise.

### GetEdgeHitRequestsOk

`func (o *HistoricalFieldResultsAttributes) GetEdgeHitRequestsOk() (*int32, bool)`

GetEdgeHitRequestsOk returns a tuple with the EdgeHitRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeHitRequests

`func (o *HistoricalFieldResultsAttributes) SetEdgeHitRequests(v int32)`

SetEdgeHitRequests sets EdgeHitRequests field to given value.

### HasEdgeHitRequests

`func (o *HistoricalFieldResultsAttributes) HasEdgeHitRequests() bool`

HasEdgeHitRequests returns a boolean if a field has been set.

### GetEdgeMissRequests

`func (o *HistoricalFieldResultsAttributes) GetEdgeMissRequests() int32`

GetEdgeMissRequests returns the EdgeMissRequests field if non-nil, zero value otherwise.

### GetEdgeMissRequestsOk

`func (o *HistoricalFieldResultsAttributes) GetEdgeMissRequestsOk() (*int32, bool)`

GetEdgeMissRequestsOk returns a tuple with the EdgeMissRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeMissRequests

`func (o *HistoricalFieldResultsAttributes) SetEdgeMissRequests(v int32)`

SetEdgeMissRequests sets EdgeMissRequests field to given value.

### HasEdgeMissRequests

`func (o *HistoricalFieldResultsAttributes) HasEdgeMissRequests() bool`

HasEdgeMissRequests returns a boolean if a field has been set.

### GetOriginFetches

`func (o *HistoricalFieldResultsAttributes) GetOriginFetches() int32`

GetOriginFetches returns the OriginFetches field if non-nil, zero value otherwise.

### GetOriginFetchesOk

`func (o *HistoricalFieldResultsAttributes) GetOriginFetchesOk() (*int32, bool)`

GetOriginFetchesOk returns a tuple with the OriginFetches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginFetches

`func (o *HistoricalFieldResultsAttributes) SetOriginFetches(v int32)`

SetOriginFetches sets OriginFetches field to given value.

### HasOriginFetches

`func (o *HistoricalFieldResultsAttributes) HasOriginFetches() bool`

HasOriginFetches returns a boolean if a field has been set.

### GetOriginFetchHeaderBytes

`func (o *HistoricalFieldResultsAttributes) GetOriginFetchHeaderBytes() int32`

GetOriginFetchHeaderBytes returns the OriginFetchHeaderBytes field if non-nil, zero value otherwise.

### GetOriginFetchHeaderBytesOk

`func (o *HistoricalFieldResultsAttributes) GetOriginFetchHeaderBytesOk() (*int32, bool)`

GetOriginFetchHeaderBytesOk returns a tuple with the OriginFetchHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginFetchHeaderBytes

`func (o *HistoricalFieldResultsAttributes) SetOriginFetchHeaderBytes(v int32)`

SetOriginFetchHeaderBytes sets OriginFetchHeaderBytes field to given value.

### HasOriginFetchHeaderBytes

`func (o *HistoricalFieldResultsAttributes) HasOriginFetchHeaderBytes() bool`

HasOriginFetchHeaderBytes returns a boolean if a field has been set.

### GetOriginFetchBodyBytes

`func (o *HistoricalFieldResultsAttributes) GetOriginFetchBodyBytes() int32`

GetOriginFetchBodyBytes returns the OriginFetchBodyBytes field if non-nil, zero value otherwise.

### GetOriginFetchBodyBytesOk

`func (o *HistoricalFieldResultsAttributes) GetOriginFetchBodyBytesOk() (*int32, bool)`

GetOriginFetchBodyBytesOk returns a tuple with the OriginFetchBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginFetchBodyBytes

`func (o *HistoricalFieldResultsAttributes) SetOriginFetchBodyBytes(v int32)`

SetOriginFetchBodyBytes sets OriginFetchBodyBytes field to given value.

### HasOriginFetchBodyBytes

`func (o *HistoricalFieldResultsAttributes) HasOriginFetchBodyBytes() bool`

HasOriginFetchBodyBytes returns a boolean if a field has been set.

### GetOriginFetchRespHeaderBytes

`func (o *HistoricalFieldResultsAttributes) GetOriginFetchRespHeaderBytes() int32`

GetOriginFetchRespHeaderBytes returns the OriginFetchRespHeaderBytes field if non-nil, zero value otherwise.

### GetOriginFetchRespHeaderBytesOk

`func (o *HistoricalFieldResultsAttributes) GetOriginFetchRespHeaderBytesOk() (*int32, bool)`

GetOriginFetchRespHeaderBytesOk returns a tuple with the OriginFetchRespHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginFetchRespHeaderBytes

`func (o *HistoricalFieldResultsAttributes) SetOriginFetchRespHeaderBytes(v int32)`

SetOriginFetchRespHeaderBytes sets OriginFetchRespHeaderBytes field to given value.

### HasOriginFetchRespHeaderBytes

`func (o *HistoricalFieldResultsAttributes) HasOriginFetchRespHeaderBytes() bool`

HasOriginFetchRespHeaderBytes returns a boolean if a field has been set.

### GetOriginFetchRespBodyBytes

`func (o *HistoricalFieldResultsAttributes) GetOriginFetchRespBodyBytes() int32`

GetOriginFetchRespBodyBytes returns the OriginFetchRespBodyBytes field if non-nil, zero value otherwise.

### GetOriginFetchRespBodyBytesOk

`func (o *HistoricalFieldResultsAttributes) GetOriginFetchRespBodyBytesOk() (*int32, bool)`

GetOriginFetchRespBodyBytesOk returns a tuple with the OriginFetchRespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginFetchRespBodyBytes

`func (o *HistoricalFieldResultsAttributes) SetOriginFetchRespBodyBytes(v int32)`

SetOriginFetchRespBodyBytes sets OriginFetchRespBodyBytes field to given value.

### HasOriginFetchRespBodyBytes

`func (o *HistoricalFieldResultsAttributes) HasOriginFetchRespBodyBytes() bool`

HasOriginFetchRespBodyBytes returns a boolean if a field has been set.

### GetOriginRevalidations

`func (o *HistoricalFieldResultsAttributes) GetOriginRevalidations() int32`

GetOriginRevalidations returns the OriginRevalidations field if non-nil, zero value otherwise.

### GetOriginRevalidationsOk

`func (o *HistoricalFieldResultsAttributes) GetOriginRevalidationsOk() (*int32, bool)`

GetOriginRevalidationsOk returns a tuple with the OriginRevalidations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginRevalidations

`func (o *HistoricalFieldResultsAttributes) SetOriginRevalidations(v int32)`

SetOriginRevalidations sets OriginRevalidations field to given value.

### HasOriginRevalidations

`func (o *HistoricalFieldResultsAttributes) HasOriginRevalidations() bool`

HasOriginRevalidations returns a boolean if a field has been set.

### GetOriginCacheFetches

`func (o *HistoricalFieldResultsAttributes) GetOriginCacheFetches() int32`

GetOriginCacheFetches returns the OriginCacheFetches field if non-nil, zero value otherwise.

### GetOriginCacheFetchesOk

`func (o *HistoricalFieldResultsAttributes) GetOriginCacheFetchesOk() (*int32, bool)`

GetOriginCacheFetchesOk returns a tuple with the OriginCacheFetches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginCacheFetches

`func (o *HistoricalFieldResultsAttributes) SetOriginCacheFetches(v int32)`

SetOriginCacheFetches sets OriginCacheFetches field to given value.

### HasOriginCacheFetches

`func (o *HistoricalFieldResultsAttributes) HasOriginCacheFetches() bool`

HasOriginCacheFetches returns a boolean if a field has been set.

### GetShield

`func (o *HistoricalFieldResultsAttributes) GetShield() int32`

GetShield returns the Shield field if non-nil, zero value otherwise.

### GetShieldOk

`func (o *HistoricalFieldResultsAttributes) GetShieldOk() (*int32, bool)`

GetShieldOk returns a tuple with the Shield field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShield

`func (o *HistoricalFieldResultsAttributes) SetShield(v int32)`

SetShield sets Shield field to given value.

### HasShield

`func (o *HistoricalFieldResultsAttributes) HasShield() bool`

HasShield returns a boolean if a field has been set.

### GetShieldRespBodyBytes

`func (o *HistoricalFieldResultsAttributes) GetShieldRespBodyBytes() int32`

GetShieldRespBodyBytes returns the ShieldRespBodyBytes field if non-nil, zero value otherwise.

### GetShieldRespBodyBytesOk

`func (o *HistoricalFieldResultsAttributes) GetShieldRespBodyBytesOk() (*int32, bool)`

GetShieldRespBodyBytesOk returns a tuple with the ShieldRespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShieldRespBodyBytes

`func (o *HistoricalFieldResultsAttributes) SetShieldRespBodyBytes(v int32)`

SetShieldRespBodyBytes sets ShieldRespBodyBytes field to given value.

### HasShieldRespBodyBytes

`func (o *HistoricalFieldResultsAttributes) HasShieldRespBodyBytes() bool`

HasShieldRespBodyBytes returns a boolean if a field has been set.

### GetShieldRespHeaderBytes

`func (o *HistoricalFieldResultsAttributes) GetShieldRespHeaderBytes() int32`

GetShieldRespHeaderBytes returns the ShieldRespHeaderBytes field if non-nil, zero value otherwise.

### GetShieldRespHeaderBytesOk

`func (o *HistoricalFieldResultsAttributes) GetShieldRespHeaderBytesOk() (*int32, bool)`

GetShieldRespHeaderBytesOk returns a tuple with the ShieldRespHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShieldRespHeaderBytes

`func (o *HistoricalFieldResultsAttributes) SetShieldRespHeaderBytes(v int32)`

SetShieldRespHeaderBytes sets ShieldRespHeaderBytes field to given value.

### HasShieldRespHeaderBytes

`func (o *HistoricalFieldResultsAttributes) HasShieldRespHeaderBytes() bool`

HasShieldRespHeaderBytes returns a boolean if a field has been set.

### GetShieldFetches

`func (o *HistoricalFieldResultsAttributes) GetShieldFetches() int32`

GetShieldFetches returns the ShieldFetches field if non-nil, zero value otherwise.

### GetShieldFetchesOk

`func (o *HistoricalFieldResultsAttributes) GetShieldFetchesOk() (*int32, bool)`

GetShieldFetchesOk returns a tuple with the ShieldFetches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShieldFetches

`func (o *HistoricalFieldResultsAttributes) SetShieldFetches(v int32)`

SetShieldFetches sets ShieldFetches field to given value.

### HasShieldFetches

`func (o *HistoricalFieldResultsAttributes) HasShieldFetches() bool`

HasShieldFetches returns a boolean if a field has been set.

### GetShieldFetchHeaderBytes

`func (o *HistoricalFieldResultsAttributes) GetShieldFetchHeaderBytes() int32`

GetShieldFetchHeaderBytes returns the ShieldFetchHeaderBytes field if non-nil, zero value otherwise.

### GetShieldFetchHeaderBytesOk

`func (o *HistoricalFieldResultsAttributes) GetShieldFetchHeaderBytesOk() (*int32, bool)`

GetShieldFetchHeaderBytesOk returns a tuple with the ShieldFetchHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShieldFetchHeaderBytes

`func (o *HistoricalFieldResultsAttributes) SetShieldFetchHeaderBytes(v int32)`

SetShieldFetchHeaderBytes sets ShieldFetchHeaderBytes field to given value.

### HasShieldFetchHeaderBytes

`func (o *HistoricalFieldResultsAttributes) HasShieldFetchHeaderBytes() bool`

HasShieldFetchHeaderBytes returns a boolean if a field has been set.

### GetShieldFetchBodyBytes

`func (o *HistoricalFieldResultsAttributes) GetShieldFetchBodyBytes() int32`

GetShieldFetchBodyBytes returns the ShieldFetchBodyBytes field if non-nil, zero value otherwise.

### GetShieldFetchBodyBytesOk

`func (o *HistoricalFieldResultsAttributes) GetShieldFetchBodyBytesOk() (*int32, bool)`

GetShieldFetchBodyBytesOk returns a tuple with the ShieldFetchBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShieldFetchBodyBytes

`func (o *HistoricalFieldResultsAttributes) SetShieldFetchBodyBytes(v int32)`

SetShieldFetchBodyBytes sets ShieldFetchBodyBytes field to given value.

### HasShieldFetchBodyBytes

`func (o *HistoricalFieldResultsAttributes) HasShieldFetchBodyBytes() bool`

HasShieldFetchBodyBytes returns a boolean if a field has been set.

### GetShieldFetchRespHeaderBytes

`func (o *HistoricalFieldResultsAttributes) GetShieldFetchRespHeaderBytes() int32`

GetShieldFetchRespHeaderBytes returns the ShieldFetchRespHeaderBytes field if non-nil, zero value otherwise.

### GetShieldFetchRespHeaderBytesOk

`func (o *HistoricalFieldResultsAttributes) GetShieldFetchRespHeaderBytesOk() (*int32, bool)`

GetShieldFetchRespHeaderBytesOk returns a tuple with the ShieldFetchRespHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShieldFetchRespHeaderBytes

`func (o *HistoricalFieldResultsAttributes) SetShieldFetchRespHeaderBytes(v int32)`

SetShieldFetchRespHeaderBytes sets ShieldFetchRespHeaderBytes field to given value.

### HasShieldFetchRespHeaderBytes

`func (o *HistoricalFieldResultsAttributes) HasShieldFetchRespHeaderBytes() bool`

HasShieldFetchRespHeaderBytes returns a boolean if a field has been set.

### GetShieldFetchRespBodyBytes

`func (o *HistoricalFieldResultsAttributes) GetShieldFetchRespBodyBytes() int32`

GetShieldFetchRespBodyBytes returns the ShieldFetchRespBodyBytes field if non-nil, zero value otherwise.

### GetShieldFetchRespBodyBytesOk

`func (o *HistoricalFieldResultsAttributes) GetShieldFetchRespBodyBytesOk() (*int32, bool)`

GetShieldFetchRespBodyBytesOk returns a tuple with the ShieldFetchRespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShieldFetchRespBodyBytes

`func (o *HistoricalFieldResultsAttributes) SetShieldFetchRespBodyBytes(v int32)`

SetShieldFetchRespBodyBytes sets ShieldFetchRespBodyBytes field to given value.

### HasShieldFetchRespBodyBytes

`func (o *HistoricalFieldResultsAttributes) HasShieldFetchRespBodyBytes() bool`

HasShieldFetchRespBodyBytes returns a boolean if a field has been set.

### GetShieldRevalidations

`func (o *HistoricalFieldResultsAttributes) GetShieldRevalidations() int32`

GetShieldRevalidations returns the ShieldRevalidations field if non-nil, zero value otherwise.

### GetShieldRevalidationsOk

`func (o *HistoricalFieldResultsAttributes) GetShieldRevalidationsOk() (*int32, bool)`

GetShieldRevalidationsOk returns a tuple with the ShieldRevalidations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShieldRevalidations

`func (o *HistoricalFieldResultsAttributes) SetShieldRevalidations(v int32)`

SetShieldRevalidations sets ShieldRevalidations field to given value.

### HasShieldRevalidations

`func (o *HistoricalFieldResultsAttributes) HasShieldRevalidations() bool`

HasShieldRevalidations returns a boolean if a field has been set.

### GetShieldCacheFetches

`func (o *HistoricalFieldResultsAttributes) GetShieldCacheFetches() int32`

GetShieldCacheFetches returns the ShieldCacheFetches field if non-nil, zero value otherwise.

### GetShieldCacheFetchesOk

`func (o *HistoricalFieldResultsAttributes) GetShieldCacheFetchesOk() (*int32, bool)`

GetShieldCacheFetchesOk returns a tuple with the ShieldCacheFetches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShieldCacheFetches

`func (o *HistoricalFieldResultsAttributes) SetShieldCacheFetches(v int32)`

SetShieldCacheFetches sets ShieldCacheFetches field to given value.

### HasShieldCacheFetches

`func (o *HistoricalFieldResultsAttributes) HasShieldCacheFetches() bool`

HasShieldCacheFetches returns a boolean if a field has been set.

### GetIpv6

`func (o *HistoricalFieldResultsAttributes) GetIpv6() int32`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *HistoricalFieldResultsAttributes) GetIpv6Ok() (*int32, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *HistoricalFieldResultsAttributes) SetIpv6(v int32)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *HistoricalFieldResultsAttributes) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.

### GetOtfp

`func (o *HistoricalFieldResultsAttributes) GetOtfp() int32`

GetOtfp returns the Otfp field if non-nil, zero value otherwise.

### GetOtfpOk

`func (o *HistoricalFieldResultsAttributes) GetOtfpOk() (*int32, bool)`

GetOtfpOk returns a tuple with the Otfp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtfp

`func (o *HistoricalFieldResultsAttributes) SetOtfp(v int32)`

SetOtfp sets Otfp field to given value.

### HasOtfp

`func (o *HistoricalFieldResultsAttributes) HasOtfp() bool`

HasOtfp returns a boolean if a field has been set.

### GetOtfpRespBodyBytes

`func (o *HistoricalFieldResultsAttributes) GetOtfpRespBodyBytes() int32`

GetOtfpRespBodyBytes returns the OtfpRespBodyBytes field if non-nil, zero value otherwise.

### GetOtfpRespBodyBytesOk

`func (o *HistoricalFieldResultsAttributes) GetOtfpRespBodyBytesOk() (*int32, bool)`

GetOtfpRespBodyBytesOk returns a tuple with the OtfpRespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtfpRespBodyBytes

`func (o *HistoricalFieldResultsAttributes) SetOtfpRespBodyBytes(v int32)`

SetOtfpRespBodyBytes sets OtfpRespBodyBytes field to given value.

### HasOtfpRespBodyBytes

`func (o *HistoricalFieldResultsAttributes) HasOtfpRespBodyBytes() bool`

HasOtfpRespBodyBytes returns a boolean if a field has been set.

### GetOtfpRespHeaderBytes

`func (o *HistoricalFieldResultsAttributes) GetOtfpRespHeaderBytes() int32`

GetOtfpRespHeaderBytes returns the OtfpRespHeaderBytes field if non-nil, zero value otherwise.

### GetOtfpRespHeaderBytesOk

`func (o *HistoricalFieldResultsAttributes) GetOtfpRespHeaderBytesOk() (*int32, bool)`

GetOtfpRespHeaderBytesOk returns a tuple with the OtfpRespHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtfpRespHeaderBytes

`func (o *HistoricalFieldResultsAttributes) SetOtfpRespHeaderBytes(v int32)`

SetOtfpRespHeaderBytes sets OtfpRespHeaderBytes field to given value.

### HasOtfpRespHeaderBytes

`func (o *HistoricalFieldResultsAttributes) HasOtfpRespHeaderBytes() bool`

HasOtfpRespHeaderBytes returns a boolean if a field has been set.

### GetOtfpShieldRespBodyBytes

`func (o *HistoricalFieldResultsAttributes) GetOtfpShieldRespBodyBytes() int32`

GetOtfpShieldRespBodyBytes returns the OtfpShieldRespBodyBytes field if non-nil, zero value otherwise.

### GetOtfpShieldRespBodyBytesOk

`func (o *HistoricalFieldResultsAttributes) GetOtfpShieldRespBodyBytesOk() (*int32, bool)`

GetOtfpShieldRespBodyBytesOk returns a tuple with the OtfpShieldRespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtfpShieldRespBodyBytes

`func (o *HistoricalFieldResultsAttributes) SetOtfpShieldRespBodyBytes(v int32)`

SetOtfpShieldRespBodyBytes sets OtfpShieldRespBodyBytes field to given value.

### HasOtfpShieldRespBodyBytes

`func (o *HistoricalFieldResultsAttributes) HasOtfpShieldRespBodyBytes() bool`

HasOtfpShieldRespBodyBytes returns a boolean if a field has been set.

### GetOtfpShieldRespHeaderBytes

`func (o *HistoricalFieldResultsAttributes) GetOtfpShieldRespHeaderBytes() int32`

GetOtfpShieldRespHeaderBytes returns the OtfpShieldRespHeaderBytes field if non-nil, zero value otherwise.

### GetOtfpShieldRespHeaderBytesOk

`func (o *HistoricalFieldResultsAttributes) GetOtfpShieldRespHeaderBytesOk() (*int32, bool)`

GetOtfpShieldRespHeaderBytesOk returns a tuple with the OtfpShieldRespHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtfpShieldRespHeaderBytes

`func (o *HistoricalFieldResultsAttributes) SetOtfpShieldRespHeaderBytes(v int32)`

SetOtfpShieldRespHeaderBytes sets OtfpShieldRespHeaderBytes field to given value.

### HasOtfpShieldRespHeaderBytes

`func (o *HistoricalFieldResultsAttributes) HasOtfpShieldRespHeaderBytes() bool`

HasOtfpShieldRespHeaderBytes returns a boolean if a field has been set.

### GetOtfpManifests

`func (o *HistoricalFieldResultsAttributes) GetOtfpManifests() int32`

GetOtfpManifests returns the OtfpManifests field if non-nil, zero value otherwise.

### GetOtfpManifestsOk

`func (o *HistoricalFieldResultsAttributes) GetOtfpManifestsOk() (*int32, bool)`

GetOtfpManifestsOk returns a tuple with the OtfpManifests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtfpManifests

`func (o *HistoricalFieldResultsAttributes) SetOtfpManifests(v int32)`

SetOtfpManifests sets OtfpManifests field to given value.

### HasOtfpManifests

`func (o *HistoricalFieldResultsAttributes) HasOtfpManifests() bool`

HasOtfpManifests returns a boolean if a field has been set.

### GetOtfpDeliverTime

`func (o *HistoricalFieldResultsAttributes) GetOtfpDeliverTime() float32`

GetOtfpDeliverTime returns the OtfpDeliverTime field if non-nil, zero value otherwise.

### GetOtfpDeliverTimeOk

`func (o *HistoricalFieldResultsAttributes) GetOtfpDeliverTimeOk() (*float32, bool)`

GetOtfpDeliverTimeOk returns a tuple with the OtfpDeliverTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtfpDeliverTime

`func (o *HistoricalFieldResultsAttributes) SetOtfpDeliverTime(v float32)`

SetOtfpDeliverTime sets OtfpDeliverTime field to given value.

### HasOtfpDeliverTime

`func (o *HistoricalFieldResultsAttributes) HasOtfpDeliverTime() bool`

HasOtfpDeliverTime returns a boolean if a field has been set.

### GetOtfpShieldTime

`func (o *HistoricalFieldResultsAttributes) GetOtfpShieldTime() float32`

GetOtfpShieldTime returns the OtfpShieldTime field if non-nil, zero value otherwise.

### GetOtfpShieldTimeOk

`func (o *HistoricalFieldResultsAttributes) GetOtfpShieldTimeOk() (*float32, bool)`

GetOtfpShieldTimeOk returns a tuple with the OtfpShieldTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtfpShieldTime

`func (o *HistoricalFieldResultsAttributes) SetOtfpShieldTime(v float32)`

SetOtfpShieldTime sets OtfpShieldTime field to given value.

### HasOtfpShieldTime

`func (o *HistoricalFieldResultsAttributes) HasOtfpShieldTime() bool`

HasOtfpShieldTime returns a boolean if a field has been set.

### GetVideo

`func (o *HistoricalFieldResultsAttributes) GetVideo() int32`

GetVideo returns the Video field if non-nil, zero value otherwise.

### GetVideoOk

`func (o *HistoricalFieldResultsAttributes) GetVideoOk() (*int32, bool)`

GetVideoOk returns a tuple with the Video field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideo

`func (o *HistoricalFieldResultsAttributes) SetVideo(v int32)`

SetVideo sets Video field to given value.

### HasVideo

`func (o *HistoricalFieldResultsAttributes) HasVideo() bool`

HasVideo returns a boolean if a field has been set.

### GetPci

`func (o *HistoricalFieldResultsAttributes) GetPci() int32`

GetPci returns the Pci field if non-nil, zero value otherwise.

### GetPciOk

`func (o *HistoricalFieldResultsAttributes) GetPciOk() (*int32, bool)`

GetPciOk returns a tuple with the Pci field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPci

`func (o *HistoricalFieldResultsAttributes) SetPci(v int32)`

SetPci sets Pci field to given value.

### HasPci

`func (o *HistoricalFieldResultsAttributes) HasPci() bool`

HasPci returns a boolean if a field has been set.

### GetLog

`func (o *HistoricalFieldResultsAttributes) GetLog() int32`

GetLog returns the Log field if non-nil, zero value otherwise.

### GetLogOk

`func (o *HistoricalFieldResultsAttributes) GetLogOk() (*int32, bool)`

GetLogOk returns a tuple with the Log field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLog

`func (o *HistoricalFieldResultsAttributes) SetLog(v int32)`

SetLog sets Log field to given value.

### HasLog

`func (o *HistoricalFieldResultsAttributes) HasLog() bool`

HasLog returns a boolean if a field has been set.

### GetLogBytes

`func (o *HistoricalFieldResultsAttributes) GetLogBytes() int32`

GetLogBytes returns the LogBytes field if non-nil, zero value otherwise.

### GetLogBytesOk

`func (o *HistoricalFieldResultsAttributes) GetLogBytesOk() (*int32, bool)`

GetLogBytesOk returns a tuple with the LogBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogBytes

`func (o *HistoricalFieldResultsAttributes) SetLogBytes(v int32)`

SetLogBytes sets LogBytes field to given value.

### HasLogBytes

`func (o *HistoricalFieldResultsAttributes) HasLogBytes() bool`

HasLogBytes returns a boolean if a field has been set.

### GetHTTP2

`func (o *HistoricalFieldResultsAttributes) GetHTTP2() int32`

GetHTTP2 returns the HTTP2 field if non-nil, zero value otherwise.

### GetHTTP2Ok

`func (o *HistoricalFieldResultsAttributes) GetHTTP2Ok() (*int32, bool)`

GetHTTP2Ok returns a tuple with the HTTP2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHTTP2

`func (o *HistoricalFieldResultsAttributes) SetHTTP2(v int32)`

SetHTTP2 sets HTTP2 field to given value.

### HasHTTP2

`func (o *HistoricalFieldResultsAttributes) HasHTTP2() bool`

HasHTTP2 returns a boolean if a field has been set.

### GetHTTP3

`func (o *HistoricalFieldResultsAttributes) GetHTTP3() int32`

GetHTTP3 returns the HTTP3 field if non-nil, zero value otherwise.

### GetHTTP3Ok

`func (o *HistoricalFieldResultsAttributes) GetHTTP3Ok() (*int32, bool)`

GetHTTP3Ok returns a tuple with the HTTP3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHTTP3

`func (o *HistoricalFieldResultsAttributes) SetHTTP3(v int32)`

SetHTTP3 sets HTTP3 field to given value.

### HasHTTP3

`func (o *HistoricalFieldResultsAttributes) HasHTTP3() bool`

HasHTTP3 returns a boolean if a field has been set.

### GetWafLogged

`func (o *HistoricalFieldResultsAttributes) GetWafLogged() int32`

GetWafLogged returns the WafLogged field if non-nil, zero value otherwise.

### GetWafLoggedOk

`func (o *HistoricalFieldResultsAttributes) GetWafLoggedOk() (*int32, bool)`

GetWafLoggedOk returns a tuple with the WafLogged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafLogged

`func (o *HistoricalFieldResultsAttributes) SetWafLogged(v int32)`

SetWafLogged sets WafLogged field to given value.

### HasWafLogged

`func (o *HistoricalFieldResultsAttributes) HasWafLogged() bool`

HasWafLogged returns a boolean if a field has been set.

### GetWafBlocked

`func (o *HistoricalFieldResultsAttributes) GetWafBlocked() int32`

GetWafBlocked returns the WafBlocked field if non-nil, zero value otherwise.

### GetWafBlockedOk

`func (o *HistoricalFieldResultsAttributes) GetWafBlockedOk() (*int32, bool)`

GetWafBlockedOk returns a tuple with the WafBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafBlocked

`func (o *HistoricalFieldResultsAttributes) SetWafBlocked(v int32)`

SetWafBlocked sets WafBlocked field to given value.

### HasWafBlocked

`func (o *HistoricalFieldResultsAttributes) HasWafBlocked() bool`

HasWafBlocked returns a boolean if a field has been set.

### GetWafPassed

`func (o *HistoricalFieldResultsAttributes) GetWafPassed() int32`

GetWafPassed returns the WafPassed field if non-nil, zero value otherwise.

### GetWafPassedOk

`func (o *HistoricalFieldResultsAttributes) GetWafPassedOk() (*int32, bool)`

GetWafPassedOk returns a tuple with the WafPassed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafPassed

`func (o *HistoricalFieldResultsAttributes) SetWafPassed(v int32)`

SetWafPassed sets WafPassed field to given value.

### HasWafPassed

`func (o *HistoricalFieldResultsAttributes) HasWafPassed() bool`

HasWafPassed returns a boolean if a field has been set.

### GetAttackReqBodyBytes

`func (o *HistoricalFieldResultsAttributes) GetAttackReqBodyBytes() int32`

GetAttackReqBodyBytes returns the AttackReqBodyBytes field if non-nil, zero value otherwise.

### GetAttackReqBodyBytesOk

`func (o *HistoricalFieldResultsAttributes) GetAttackReqBodyBytesOk() (*int32, bool)`

GetAttackReqBodyBytesOk returns a tuple with the AttackReqBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackReqBodyBytes

`func (o *HistoricalFieldResultsAttributes) SetAttackReqBodyBytes(v int32)`

SetAttackReqBodyBytes sets AttackReqBodyBytes field to given value.

### HasAttackReqBodyBytes

`func (o *HistoricalFieldResultsAttributes) HasAttackReqBodyBytes() bool`

HasAttackReqBodyBytes returns a boolean if a field has been set.

### GetAttackReqHeaderBytes

`func (o *HistoricalFieldResultsAttributes) GetAttackReqHeaderBytes() int32`

GetAttackReqHeaderBytes returns the AttackReqHeaderBytes field if non-nil, zero value otherwise.

### GetAttackReqHeaderBytesOk

`func (o *HistoricalFieldResultsAttributes) GetAttackReqHeaderBytesOk() (*int32, bool)`

GetAttackReqHeaderBytesOk returns a tuple with the AttackReqHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackReqHeaderBytes

`func (o *HistoricalFieldResultsAttributes) SetAttackReqHeaderBytes(v int32)`

SetAttackReqHeaderBytes sets AttackReqHeaderBytes field to given value.

### HasAttackReqHeaderBytes

`func (o *HistoricalFieldResultsAttributes) HasAttackReqHeaderBytes() bool`

HasAttackReqHeaderBytes returns a boolean if a field has been set.

### GetAttackLoggedReqBodyBytes

`func (o *HistoricalFieldResultsAttributes) GetAttackLoggedReqBodyBytes() int32`

GetAttackLoggedReqBodyBytes returns the AttackLoggedReqBodyBytes field if non-nil, zero value otherwise.

### GetAttackLoggedReqBodyBytesOk

`func (o *HistoricalFieldResultsAttributes) GetAttackLoggedReqBodyBytesOk() (*int32, bool)`

GetAttackLoggedReqBodyBytesOk returns a tuple with the AttackLoggedReqBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackLoggedReqBodyBytes

`func (o *HistoricalFieldResultsAttributes) SetAttackLoggedReqBodyBytes(v int32)`

SetAttackLoggedReqBodyBytes sets AttackLoggedReqBodyBytes field to given value.

### HasAttackLoggedReqBodyBytes

`func (o *HistoricalFieldResultsAttributes) HasAttackLoggedReqBodyBytes() bool`

HasAttackLoggedReqBodyBytes returns a boolean if a field has been set.

### GetAttackLoggedReqHeaderBytes

`func (o *HistoricalFieldResultsAttributes) GetAttackLoggedReqHeaderBytes() int32`

GetAttackLoggedReqHeaderBytes returns the AttackLoggedReqHeaderBytes field if non-nil, zero value otherwise.

### GetAttackLoggedReqHeaderBytesOk

`func (o *HistoricalFieldResultsAttributes) GetAttackLoggedReqHeaderBytesOk() (*int32, bool)`

GetAttackLoggedReqHeaderBytesOk returns a tuple with the AttackLoggedReqHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackLoggedReqHeaderBytes

`func (o *HistoricalFieldResultsAttributes) SetAttackLoggedReqHeaderBytes(v int32)`

SetAttackLoggedReqHeaderBytes sets AttackLoggedReqHeaderBytes field to given value.

### HasAttackLoggedReqHeaderBytes

`func (o *HistoricalFieldResultsAttributes) HasAttackLoggedReqHeaderBytes() bool`

HasAttackLoggedReqHeaderBytes returns a boolean if a field has been set.

### GetAttackBlockedReqBodyBytes

`func (o *HistoricalFieldResultsAttributes) GetAttackBlockedReqBodyBytes() int32`

GetAttackBlockedReqBodyBytes returns the AttackBlockedReqBodyBytes field if non-nil, zero value otherwise.

### GetAttackBlockedReqBodyBytesOk

`func (o *HistoricalFieldResultsAttributes) GetAttackBlockedReqBodyBytesOk() (*int32, bool)`

GetAttackBlockedReqBodyBytesOk returns a tuple with the AttackBlockedReqBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackBlockedReqBodyBytes

`func (o *HistoricalFieldResultsAttributes) SetAttackBlockedReqBodyBytes(v int32)`

SetAttackBlockedReqBodyBytes sets AttackBlockedReqBodyBytes field to given value.

### HasAttackBlockedReqBodyBytes

`func (o *HistoricalFieldResultsAttributes) HasAttackBlockedReqBodyBytes() bool`

HasAttackBlockedReqBodyBytes returns a boolean if a field has been set.

### GetAttackBlockedReqHeaderBytes

`func (o *HistoricalFieldResultsAttributes) GetAttackBlockedReqHeaderBytes() int32`

GetAttackBlockedReqHeaderBytes returns the AttackBlockedReqHeaderBytes field if non-nil, zero value otherwise.

### GetAttackBlockedReqHeaderBytesOk

`func (o *HistoricalFieldResultsAttributes) GetAttackBlockedReqHeaderBytesOk() (*int32, bool)`

GetAttackBlockedReqHeaderBytesOk returns a tuple with the AttackBlockedReqHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackBlockedReqHeaderBytes

`func (o *HistoricalFieldResultsAttributes) SetAttackBlockedReqHeaderBytes(v int32)`

SetAttackBlockedReqHeaderBytes sets AttackBlockedReqHeaderBytes field to given value.

### HasAttackBlockedReqHeaderBytes

`func (o *HistoricalFieldResultsAttributes) HasAttackBlockedReqHeaderBytes() bool`

HasAttackBlockedReqHeaderBytes returns a boolean if a field has been set.

### GetAttackPassedReqBodyBytes

`func (o *HistoricalFieldResultsAttributes) GetAttackPassedReqBodyBytes() int32`

GetAttackPassedReqBodyBytes returns the AttackPassedReqBodyBytes field if non-nil, zero value otherwise.

### GetAttackPassedReqBodyBytesOk

`func (o *HistoricalFieldResultsAttributes) GetAttackPassedReqBodyBytesOk() (*int32, bool)`

GetAttackPassedReqBodyBytesOk returns a tuple with the AttackPassedReqBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackPassedReqBodyBytes

`func (o *HistoricalFieldResultsAttributes) SetAttackPassedReqBodyBytes(v int32)`

SetAttackPassedReqBodyBytes sets AttackPassedReqBodyBytes field to given value.

### HasAttackPassedReqBodyBytes

`func (o *HistoricalFieldResultsAttributes) HasAttackPassedReqBodyBytes() bool`

HasAttackPassedReqBodyBytes returns a boolean if a field has been set.

### GetAttackPassedReqHeaderBytes

`func (o *HistoricalFieldResultsAttributes) GetAttackPassedReqHeaderBytes() int32`

GetAttackPassedReqHeaderBytes returns the AttackPassedReqHeaderBytes field if non-nil, zero value otherwise.

### GetAttackPassedReqHeaderBytesOk

`func (o *HistoricalFieldResultsAttributes) GetAttackPassedReqHeaderBytesOk() (*int32, bool)`

GetAttackPassedReqHeaderBytesOk returns a tuple with the AttackPassedReqHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackPassedReqHeaderBytes

`func (o *HistoricalFieldResultsAttributes) SetAttackPassedReqHeaderBytes(v int32)`

SetAttackPassedReqHeaderBytes sets AttackPassedReqHeaderBytes field to given value.

### HasAttackPassedReqHeaderBytes

`func (o *HistoricalFieldResultsAttributes) HasAttackPassedReqHeaderBytes() bool`

HasAttackPassedReqHeaderBytes returns a boolean if a field has been set.

### GetAttackRespSynthBytes

`func (o *HistoricalFieldResultsAttributes) GetAttackRespSynthBytes() int32`

GetAttackRespSynthBytes returns the AttackRespSynthBytes field if non-nil, zero value otherwise.

### GetAttackRespSynthBytesOk

`func (o *HistoricalFieldResultsAttributes) GetAttackRespSynthBytesOk() (*int32, bool)`

GetAttackRespSynthBytesOk returns a tuple with the AttackRespSynthBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackRespSynthBytes

`func (o *HistoricalFieldResultsAttributes) SetAttackRespSynthBytes(v int32)`

SetAttackRespSynthBytes sets AttackRespSynthBytes field to given value.

### HasAttackRespSynthBytes

`func (o *HistoricalFieldResultsAttributes) HasAttackRespSynthBytes() bool`

HasAttackRespSynthBytes returns a boolean if a field has been set.

### GetImgopto

`func (o *HistoricalFieldResultsAttributes) GetImgopto() int32`

GetImgopto returns the Imgopto field if non-nil, zero value otherwise.

### GetImgoptoOk

`func (o *HistoricalFieldResultsAttributes) GetImgoptoOk() (*int32, bool)`

GetImgoptoOk returns a tuple with the Imgopto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImgopto

`func (o *HistoricalFieldResultsAttributes) SetImgopto(v int32)`

SetImgopto sets Imgopto field to given value.

### HasImgopto

`func (o *HistoricalFieldResultsAttributes) HasImgopto() bool`

HasImgopto returns a boolean if a field has been set.

### GetImgoptoRespBodyBytes

`func (o *HistoricalFieldResultsAttributes) GetImgoptoRespBodyBytes() int32`

GetImgoptoRespBodyBytes returns the ImgoptoRespBodyBytes field if non-nil, zero value otherwise.

### GetImgoptoRespBodyBytesOk

`func (o *HistoricalFieldResultsAttributes) GetImgoptoRespBodyBytesOk() (*int32, bool)`

GetImgoptoRespBodyBytesOk returns a tuple with the ImgoptoRespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImgoptoRespBodyBytes

`func (o *HistoricalFieldResultsAttributes) SetImgoptoRespBodyBytes(v int32)`

SetImgoptoRespBodyBytes sets ImgoptoRespBodyBytes field to given value.

### HasImgoptoRespBodyBytes

`func (o *HistoricalFieldResultsAttributes) HasImgoptoRespBodyBytes() bool`

HasImgoptoRespBodyBytes returns a boolean if a field has been set.

### GetImgoptoRespHeaderBytes

`func (o *HistoricalFieldResultsAttributes) GetImgoptoRespHeaderBytes() int32`

GetImgoptoRespHeaderBytes returns the ImgoptoRespHeaderBytes field if non-nil, zero value otherwise.

### GetImgoptoRespHeaderBytesOk

`func (o *HistoricalFieldResultsAttributes) GetImgoptoRespHeaderBytesOk() (*int32, bool)`

GetImgoptoRespHeaderBytesOk returns a tuple with the ImgoptoRespHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImgoptoRespHeaderBytes

`func (o *HistoricalFieldResultsAttributes) SetImgoptoRespHeaderBytes(v int32)`

SetImgoptoRespHeaderBytes sets ImgoptoRespHeaderBytes field to given value.

### HasImgoptoRespHeaderBytes

`func (o *HistoricalFieldResultsAttributes) HasImgoptoRespHeaderBytes() bool`

HasImgoptoRespHeaderBytes returns a boolean if a field has been set.

### GetImgoptoShieldRespBodyBytes

`func (o *HistoricalFieldResultsAttributes) GetImgoptoShieldRespBodyBytes() int32`

GetImgoptoShieldRespBodyBytes returns the ImgoptoShieldRespBodyBytes field if non-nil, zero value otherwise.

### GetImgoptoShieldRespBodyBytesOk

`func (o *HistoricalFieldResultsAttributes) GetImgoptoShieldRespBodyBytesOk() (*int32, bool)`

GetImgoptoShieldRespBodyBytesOk returns a tuple with the ImgoptoShieldRespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImgoptoShieldRespBodyBytes

`func (o *HistoricalFieldResultsAttributes) SetImgoptoShieldRespBodyBytes(v int32)`

SetImgoptoShieldRespBodyBytes sets ImgoptoShieldRespBodyBytes field to given value.

### HasImgoptoShieldRespBodyBytes

`func (o *HistoricalFieldResultsAttributes) HasImgoptoShieldRespBodyBytes() bool`

HasImgoptoShieldRespBodyBytes returns a boolean if a field has been set.

### GetImgoptoShieldRespHeaderBytes

`func (o *HistoricalFieldResultsAttributes) GetImgoptoShieldRespHeaderBytes() int32`

GetImgoptoShieldRespHeaderBytes returns the ImgoptoShieldRespHeaderBytes field if non-nil, zero value otherwise.

### GetImgoptoShieldRespHeaderBytesOk

`func (o *HistoricalFieldResultsAttributes) GetImgoptoShieldRespHeaderBytesOk() (*int32, bool)`

GetImgoptoShieldRespHeaderBytesOk returns a tuple with the ImgoptoShieldRespHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImgoptoShieldRespHeaderBytes

`func (o *HistoricalFieldResultsAttributes) SetImgoptoShieldRespHeaderBytes(v int32)`

SetImgoptoShieldRespHeaderBytes sets ImgoptoShieldRespHeaderBytes field to given value.

### HasImgoptoShieldRespHeaderBytes

`func (o *HistoricalFieldResultsAttributes) HasImgoptoShieldRespHeaderBytes() bool`

HasImgoptoShieldRespHeaderBytes returns a boolean if a field has been set.

### GetImgvideo

`func (o *HistoricalFieldResultsAttributes) GetImgvideo() int32`

GetImgvideo returns the Imgvideo field if non-nil, zero value otherwise.

### GetImgvideoOk

`func (o *HistoricalFieldResultsAttributes) GetImgvideoOk() (*int32, bool)`

GetImgvideoOk returns a tuple with the Imgvideo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImgvideo

`func (o *HistoricalFieldResultsAttributes) SetImgvideo(v int32)`

SetImgvideo sets Imgvideo field to given value.

### HasImgvideo

`func (o *HistoricalFieldResultsAttributes) HasImgvideo() bool`

HasImgvideo returns a boolean if a field has been set.

### GetImgvideoFrames

`func (o *HistoricalFieldResultsAttributes) GetImgvideoFrames() int32`

GetImgvideoFrames returns the ImgvideoFrames field if non-nil, zero value otherwise.

### GetImgvideoFramesOk

`func (o *HistoricalFieldResultsAttributes) GetImgvideoFramesOk() (*int32, bool)`

GetImgvideoFramesOk returns a tuple with the ImgvideoFrames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImgvideoFrames

`func (o *HistoricalFieldResultsAttributes) SetImgvideoFrames(v int32)`

SetImgvideoFrames sets ImgvideoFrames field to given value.

### HasImgvideoFrames

`func (o *HistoricalFieldResultsAttributes) HasImgvideoFrames() bool`

HasImgvideoFrames returns a boolean if a field has been set.

### GetImgvideoRespHeaderBytes

`func (o *HistoricalFieldResultsAttributes) GetImgvideoRespHeaderBytes() int32`

GetImgvideoRespHeaderBytes returns the ImgvideoRespHeaderBytes field if non-nil, zero value otherwise.

### GetImgvideoRespHeaderBytesOk

`func (o *HistoricalFieldResultsAttributes) GetImgvideoRespHeaderBytesOk() (*int32, bool)`

GetImgvideoRespHeaderBytesOk returns a tuple with the ImgvideoRespHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImgvideoRespHeaderBytes

`func (o *HistoricalFieldResultsAttributes) SetImgvideoRespHeaderBytes(v int32)`

SetImgvideoRespHeaderBytes sets ImgvideoRespHeaderBytes field to given value.

### HasImgvideoRespHeaderBytes

`func (o *HistoricalFieldResultsAttributes) HasImgvideoRespHeaderBytes() bool`

HasImgvideoRespHeaderBytes returns a boolean if a field has been set.

### GetImgvideoRespBodyBytes

`func (o *HistoricalFieldResultsAttributes) GetImgvideoRespBodyBytes() int32`

GetImgvideoRespBodyBytes returns the ImgvideoRespBodyBytes field if non-nil, zero value otherwise.

### GetImgvideoRespBodyBytesOk

`func (o *HistoricalFieldResultsAttributes) GetImgvideoRespBodyBytesOk() (*int32, bool)`

GetImgvideoRespBodyBytesOk returns a tuple with the ImgvideoRespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImgvideoRespBodyBytes

`func (o *HistoricalFieldResultsAttributes) SetImgvideoRespBodyBytes(v int32)`

SetImgvideoRespBodyBytes sets ImgvideoRespBodyBytes field to given value.

### HasImgvideoRespBodyBytes

`func (o *HistoricalFieldResultsAttributes) HasImgvideoRespBodyBytes() bool`

HasImgvideoRespBodyBytes returns a boolean if a field has been set.

### GetImgvideoShieldRespHeaderBytes

`func (o *HistoricalFieldResultsAttributes) GetImgvideoShieldRespHeaderBytes() int32`

GetImgvideoShieldRespHeaderBytes returns the ImgvideoShieldRespHeaderBytes field if non-nil, zero value otherwise.

### GetImgvideoShieldRespHeaderBytesOk

`func (o *HistoricalFieldResultsAttributes) GetImgvideoShieldRespHeaderBytesOk() (*int32, bool)`

GetImgvideoShieldRespHeaderBytesOk returns a tuple with the ImgvideoShieldRespHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImgvideoShieldRespHeaderBytes

`func (o *HistoricalFieldResultsAttributes) SetImgvideoShieldRespHeaderBytes(v int32)`

SetImgvideoShieldRespHeaderBytes sets ImgvideoShieldRespHeaderBytes field to given value.

### HasImgvideoShieldRespHeaderBytes

`func (o *HistoricalFieldResultsAttributes) HasImgvideoShieldRespHeaderBytes() bool`

HasImgvideoShieldRespHeaderBytes returns a boolean if a field has been set.

### GetImgvideoShieldRespBodyBytes

`func (o *HistoricalFieldResultsAttributes) GetImgvideoShieldRespBodyBytes() int32`

GetImgvideoShieldRespBodyBytes returns the ImgvideoShieldRespBodyBytes field if non-nil, zero value otherwise.

### GetImgvideoShieldRespBodyBytesOk

`func (o *HistoricalFieldResultsAttributes) GetImgvideoShieldRespBodyBytesOk() (*int32, bool)`

GetImgvideoShieldRespBodyBytesOk returns a tuple with the ImgvideoShieldRespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImgvideoShieldRespBodyBytes

`func (o *HistoricalFieldResultsAttributes) SetImgvideoShieldRespBodyBytes(v int32)`

SetImgvideoShieldRespBodyBytes sets ImgvideoShieldRespBodyBytes field to given value.

### HasImgvideoShieldRespBodyBytes

`func (o *HistoricalFieldResultsAttributes) HasImgvideoShieldRespBodyBytes() bool`

HasImgvideoShieldRespBodyBytes returns a boolean if a field has been set.

### GetImgvideoShield

`func (o *HistoricalFieldResultsAttributes) GetImgvideoShield() int32`

GetImgvideoShield returns the ImgvideoShield field if non-nil, zero value otherwise.

### GetImgvideoShieldOk

`func (o *HistoricalFieldResultsAttributes) GetImgvideoShieldOk() (*int32, bool)`

GetImgvideoShieldOk returns a tuple with the ImgvideoShield field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImgvideoShield

`func (o *HistoricalFieldResultsAttributes) SetImgvideoShield(v int32)`

SetImgvideoShield sets ImgvideoShield field to given value.

### HasImgvideoShield

`func (o *HistoricalFieldResultsAttributes) HasImgvideoShield() bool`

HasImgvideoShield returns a boolean if a field has been set.

### GetImgvideoShieldFrames

`func (o *HistoricalFieldResultsAttributes) GetImgvideoShieldFrames() int32`

GetImgvideoShieldFrames returns the ImgvideoShieldFrames field if non-nil, zero value otherwise.

### GetImgvideoShieldFramesOk

`func (o *HistoricalFieldResultsAttributes) GetImgvideoShieldFramesOk() (*int32, bool)`

GetImgvideoShieldFramesOk returns a tuple with the ImgvideoShieldFrames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImgvideoShieldFrames

`func (o *HistoricalFieldResultsAttributes) SetImgvideoShieldFrames(v int32)`

SetImgvideoShieldFrames sets ImgvideoShieldFrames field to given value.

### HasImgvideoShieldFrames

`func (o *HistoricalFieldResultsAttributes) HasImgvideoShieldFrames() bool`

HasImgvideoShieldFrames returns a boolean if a field has been set.

### GetStatus200

`func (o *HistoricalFieldResultsAttributes) GetStatus200() int32`

GetStatus200 returns the Status200 field if non-nil, zero value otherwise.

### GetStatus200Ok

`func (o *HistoricalFieldResultsAttributes) GetStatus200Ok() (*int32, bool)`

GetStatus200Ok returns a tuple with the Status200 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus200

`func (o *HistoricalFieldResultsAttributes) SetStatus200(v int32)`

SetStatus200 sets Status200 field to given value.

### HasStatus200

`func (o *HistoricalFieldResultsAttributes) HasStatus200() bool`

HasStatus200 returns a boolean if a field has been set.

### GetStatus204

`func (o *HistoricalFieldResultsAttributes) GetStatus204() int32`

GetStatus204 returns the Status204 field if non-nil, zero value otherwise.

### GetStatus204Ok

`func (o *HistoricalFieldResultsAttributes) GetStatus204Ok() (*int32, bool)`

GetStatus204Ok returns a tuple with the Status204 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus204

`func (o *HistoricalFieldResultsAttributes) SetStatus204(v int32)`

SetStatus204 sets Status204 field to given value.

### HasStatus204

`func (o *HistoricalFieldResultsAttributes) HasStatus204() bool`

HasStatus204 returns a boolean if a field has been set.

### GetStatus206

`func (o *HistoricalFieldResultsAttributes) GetStatus206() int32`

GetStatus206 returns the Status206 field if non-nil, zero value otherwise.

### GetStatus206Ok

`func (o *HistoricalFieldResultsAttributes) GetStatus206Ok() (*int32, bool)`

GetStatus206Ok returns a tuple with the Status206 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus206

`func (o *HistoricalFieldResultsAttributes) SetStatus206(v int32)`

SetStatus206 sets Status206 field to given value.

### HasStatus206

`func (o *HistoricalFieldResultsAttributes) HasStatus206() bool`

HasStatus206 returns a boolean if a field has been set.

### GetStatus301

`func (o *HistoricalFieldResultsAttributes) GetStatus301() int32`

GetStatus301 returns the Status301 field if non-nil, zero value otherwise.

### GetStatus301Ok

`func (o *HistoricalFieldResultsAttributes) GetStatus301Ok() (*int32, bool)`

GetStatus301Ok returns a tuple with the Status301 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus301

`func (o *HistoricalFieldResultsAttributes) SetStatus301(v int32)`

SetStatus301 sets Status301 field to given value.

### HasStatus301

`func (o *HistoricalFieldResultsAttributes) HasStatus301() bool`

HasStatus301 returns a boolean if a field has been set.

### GetStatus302

`func (o *HistoricalFieldResultsAttributes) GetStatus302() int32`

GetStatus302 returns the Status302 field if non-nil, zero value otherwise.

### GetStatus302Ok

`func (o *HistoricalFieldResultsAttributes) GetStatus302Ok() (*int32, bool)`

GetStatus302Ok returns a tuple with the Status302 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus302

`func (o *HistoricalFieldResultsAttributes) SetStatus302(v int32)`

SetStatus302 sets Status302 field to given value.

### HasStatus302

`func (o *HistoricalFieldResultsAttributes) HasStatus302() bool`

HasStatus302 returns a boolean if a field has been set.

### GetStatus304

`func (o *HistoricalFieldResultsAttributes) GetStatus304() int32`

GetStatus304 returns the Status304 field if non-nil, zero value otherwise.

### GetStatus304Ok

`func (o *HistoricalFieldResultsAttributes) GetStatus304Ok() (*int32, bool)`

GetStatus304Ok returns a tuple with the Status304 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus304

`func (o *HistoricalFieldResultsAttributes) SetStatus304(v int32)`

SetStatus304 sets Status304 field to given value.

### HasStatus304

`func (o *HistoricalFieldResultsAttributes) HasStatus304() bool`

HasStatus304 returns a boolean if a field has been set.

### GetStatus400

`func (o *HistoricalFieldResultsAttributes) GetStatus400() int32`

GetStatus400 returns the Status400 field if non-nil, zero value otherwise.

### GetStatus400Ok

`func (o *HistoricalFieldResultsAttributes) GetStatus400Ok() (*int32, bool)`

GetStatus400Ok returns a tuple with the Status400 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus400

`func (o *HistoricalFieldResultsAttributes) SetStatus400(v int32)`

SetStatus400 sets Status400 field to given value.

### HasStatus400

`func (o *HistoricalFieldResultsAttributes) HasStatus400() bool`

HasStatus400 returns a boolean if a field has been set.

### GetStatus401

`func (o *HistoricalFieldResultsAttributes) GetStatus401() int32`

GetStatus401 returns the Status401 field if non-nil, zero value otherwise.

### GetStatus401Ok

`func (o *HistoricalFieldResultsAttributes) GetStatus401Ok() (*int32, bool)`

GetStatus401Ok returns a tuple with the Status401 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus401

`func (o *HistoricalFieldResultsAttributes) SetStatus401(v int32)`

SetStatus401 sets Status401 field to given value.

### HasStatus401

`func (o *HistoricalFieldResultsAttributes) HasStatus401() bool`

HasStatus401 returns a boolean if a field has been set.

### GetStatus403

`func (o *HistoricalFieldResultsAttributes) GetStatus403() int32`

GetStatus403 returns the Status403 field if non-nil, zero value otherwise.

### GetStatus403Ok

`func (o *HistoricalFieldResultsAttributes) GetStatus403Ok() (*int32, bool)`

GetStatus403Ok returns a tuple with the Status403 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus403

`func (o *HistoricalFieldResultsAttributes) SetStatus403(v int32)`

SetStatus403 sets Status403 field to given value.

### HasStatus403

`func (o *HistoricalFieldResultsAttributes) HasStatus403() bool`

HasStatus403 returns a boolean if a field has been set.

### GetStatus404

`func (o *HistoricalFieldResultsAttributes) GetStatus404() int32`

GetStatus404 returns the Status404 field if non-nil, zero value otherwise.

### GetStatus404Ok

`func (o *HistoricalFieldResultsAttributes) GetStatus404Ok() (*int32, bool)`

GetStatus404Ok returns a tuple with the Status404 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus404

`func (o *HistoricalFieldResultsAttributes) SetStatus404(v int32)`

SetStatus404 sets Status404 field to given value.

### HasStatus404

`func (o *HistoricalFieldResultsAttributes) HasStatus404() bool`

HasStatus404 returns a boolean if a field has been set.

### GetStatus406

`func (o *HistoricalFieldResultsAttributes) GetStatus406() int32`

GetStatus406 returns the Status406 field if non-nil, zero value otherwise.

### GetStatus406Ok

`func (o *HistoricalFieldResultsAttributes) GetStatus406Ok() (*int32, bool)`

GetStatus406Ok returns a tuple with the Status406 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus406

`func (o *HistoricalFieldResultsAttributes) SetStatus406(v int32)`

SetStatus406 sets Status406 field to given value.

### HasStatus406

`func (o *HistoricalFieldResultsAttributes) HasStatus406() bool`

HasStatus406 returns a boolean if a field has been set.

### GetStatus416

`func (o *HistoricalFieldResultsAttributes) GetStatus416() int32`

GetStatus416 returns the Status416 field if non-nil, zero value otherwise.

### GetStatus416Ok

`func (o *HistoricalFieldResultsAttributes) GetStatus416Ok() (*int32, bool)`

GetStatus416Ok returns a tuple with the Status416 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus416

`func (o *HistoricalFieldResultsAttributes) SetStatus416(v int32)`

SetStatus416 sets Status416 field to given value.

### HasStatus416

`func (o *HistoricalFieldResultsAttributes) HasStatus416() bool`

HasStatus416 returns a boolean if a field has been set.

### GetStatus429

`func (o *HistoricalFieldResultsAttributes) GetStatus429() int32`

GetStatus429 returns the Status429 field if non-nil, zero value otherwise.

### GetStatus429Ok

`func (o *HistoricalFieldResultsAttributes) GetStatus429Ok() (*int32, bool)`

GetStatus429Ok returns a tuple with the Status429 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus429

`func (o *HistoricalFieldResultsAttributes) SetStatus429(v int32)`

SetStatus429 sets Status429 field to given value.

### HasStatus429

`func (o *HistoricalFieldResultsAttributes) HasStatus429() bool`

HasStatus429 returns a boolean if a field has been set.

### GetStatus500

`func (o *HistoricalFieldResultsAttributes) GetStatus500() int32`

GetStatus500 returns the Status500 field if non-nil, zero value otherwise.

### GetStatus500Ok

`func (o *HistoricalFieldResultsAttributes) GetStatus500Ok() (*int32, bool)`

GetStatus500Ok returns a tuple with the Status500 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus500

`func (o *HistoricalFieldResultsAttributes) SetStatus500(v int32)`

SetStatus500 sets Status500 field to given value.

### HasStatus500

`func (o *HistoricalFieldResultsAttributes) HasStatus500() bool`

HasStatus500 returns a boolean if a field has been set.

### GetStatus501

`func (o *HistoricalFieldResultsAttributes) GetStatus501() int32`

GetStatus501 returns the Status501 field if non-nil, zero value otherwise.

### GetStatus501Ok

`func (o *HistoricalFieldResultsAttributes) GetStatus501Ok() (*int32, bool)`

GetStatus501Ok returns a tuple with the Status501 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus501

`func (o *HistoricalFieldResultsAttributes) SetStatus501(v int32)`

SetStatus501 sets Status501 field to given value.

### HasStatus501

`func (o *HistoricalFieldResultsAttributes) HasStatus501() bool`

HasStatus501 returns a boolean if a field has been set.

### GetStatus502

`func (o *HistoricalFieldResultsAttributes) GetStatus502() int32`

GetStatus502 returns the Status502 field if non-nil, zero value otherwise.

### GetStatus502Ok

`func (o *HistoricalFieldResultsAttributes) GetStatus502Ok() (*int32, bool)`

GetStatus502Ok returns a tuple with the Status502 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus502

`func (o *HistoricalFieldResultsAttributes) SetStatus502(v int32)`

SetStatus502 sets Status502 field to given value.

### HasStatus502

`func (o *HistoricalFieldResultsAttributes) HasStatus502() bool`

HasStatus502 returns a boolean if a field has been set.

### GetStatus503

`func (o *HistoricalFieldResultsAttributes) GetStatus503() int32`

GetStatus503 returns the Status503 field if non-nil, zero value otherwise.

### GetStatus503Ok

`func (o *HistoricalFieldResultsAttributes) GetStatus503Ok() (*int32, bool)`

GetStatus503Ok returns a tuple with the Status503 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus503

`func (o *HistoricalFieldResultsAttributes) SetStatus503(v int32)`

SetStatus503 sets Status503 field to given value.

### HasStatus503

`func (o *HistoricalFieldResultsAttributes) HasStatus503() bool`

HasStatus503 returns a boolean if a field has been set.

### GetStatus504

`func (o *HistoricalFieldResultsAttributes) GetStatus504() int32`

GetStatus504 returns the Status504 field if non-nil, zero value otherwise.

### GetStatus504Ok

`func (o *HistoricalFieldResultsAttributes) GetStatus504Ok() (*int32, bool)`

GetStatus504Ok returns a tuple with the Status504 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus504

`func (o *HistoricalFieldResultsAttributes) SetStatus504(v int32)`

SetStatus504 sets Status504 field to given value.

### HasStatus504

`func (o *HistoricalFieldResultsAttributes) HasStatus504() bool`

HasStatus504 returns a boolean if a field has been set.

### GetStatus505

`func (o *HistoricalFieldResultsAttributes) GetStatus505() int32`

GetStatus505 returns the Status505 field if non-nil, zero value otherwise.

### GetStatus505Ok

`func (o *HistoricalFieldResultsAttributes) GetStatus505Ok() (*int32, bool)`

GetStatus505Ok returns a tuple with the Status505 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus505

`func (o *HistoricalFieldResultsAttributes) SetStatus505(v int32)`

SetStatus505 sets Status505 field to given value.

### HasStatus505

`func (o *HistoricalFieldResultsAttributes) HasStatus505() bool`

HasStatus505 returns a boolean if a field has been set.

### GetStatus1xx

`func (o *HistoricalFieldResultsAttributes) GetStatus1xx() int32`

GetStatus1xx returns the Status1xx field if non-nil, zero value otherwise.

### GetStatus1xxOk

`func (o *HistoricalFieldResultsAttributes) GetStatus1xxOk() (*int32, bool)`

GetStatus1xxOk returns a tuple with the Status1xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus1xx

`func (o *HistoricalFieldResultsAttributes) SetStatus1xx(v int32)`

SetStatus1xx sets Status1xx field to given value.

### HasStatus1xx

`func (o *HistoricalFieldResultsAttributes) HasStatus1xx() bool`

HasStatus1xx returns a boolean if a field has been set.

### GetStatus2xx

`func (o *HistoricalFieldResultsAttributes) GetStatus2xx() int32`

GetStatus2xx returns the Status2xx field if non-nil, zero value otherwise.

### GetStatus2xxOk

`func (o *HistoricalFieldResultsAttributes) GetStatus2xxOk() (*int32, bool)`

GetStatus2xxOk returns a tuple with the Status2xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus2xx

`func (o *HistoricalFieldResultsAttributes) SetStatus2xx(v int32)`

SetStatus2xx sets Status2xx field to given value.

### HasStatus2xx

`func (o *HistoricalFieldResultsAttributes) HasStatus2xx() bool`

HasStatus2xx returns a boolean if a field has been set.

### GetStatus3xx

`func (o *HistoricalFieldResultsAttributes) GetStatus3xx() int32`

GetStatus3xx returns the Status3xx field if non-nil, zero value otherwise.

### GetStatus3xxOk

`func (o *HistoricalFieldResultsAttributes) GetStatus3xxOk() (*int32, bool)`

GetStatus3xxOk returns a tuple with the Status3xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus3xx

`func (o *HistoricalFieldResultsAttributes) SetStatus3xx(v int32)`

SetStatus3xx sets Status3xx field to given value.

### HasStatus3xx

`func (o *HistoricalFieldResultsAttributes) HasStatus3xx() bool`

HasStatus3xx returns a boolean if a field has been set.

### GetStatus4xx

`func (o *HistoricalFieldResultsAttributes) GetStatus4xx() int32`

GetStatus4xx returns the Status4xx field if non-nil, zero value otherwise.

### GetStatus4xxOk

`func (o *HistoricalFieldResultsAttributes) GetStatus4xxOk() (*int32, bool)`

GetStatus4xxOk returns a tuple with the Status4xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus4xx

`func (o *HistoricalFieldResultsAttributes) SetStatus4xx(v int32)`

SetStatus4xx sets Status4xx field to given value.

### HasStatus4xx

`func (o *HistoricalFieldResultsAttributes) HasStatus4xx() bool`

HasStatus4xx returns a boolean if a field has been set.

### GetStatus5xx

`func (o *HistoricalFieldResultsAttributes) GetStatus5xx() int32`

GetStatus5xx returns the Status5xx field if non-nil, zero value otherwise.

### GetStatus5xxOk

`func (o *HistoricalFieldResultsAttributes) GetStatus5xxOk() (*int32, bool)`

GetStatus5xxOk returns a tuple with the Status5xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus5xx

`func (o *HistoricalFieldResultsAttributes) SetStatus5xx(v int32)`

SetStatus5xx sets Status5xx field to given value.

### HasStatus5xx

`func (o *HistoricalFieldResultsAttributes) HasStatus5xx() bool`

HasStatus5xx returns a boolean if a field has been set.

### GetObjectSize1k

`func (o *HistoricalFieldResultsAttributes) GetObjectSize1k() int32`

GetObjectSize1k returns the ObjectSize1k field if non-nil, zero value otherwise.

### GetObjectSize1kOk

`func (o *HistoricalFieldResultsAttributes) GetObjectSize1kOk() (*int32, bool)`

GetObjectSize1kOk returns a tuple with the ObjectSize1k field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectSize1k

`func (o *HistoricalFieldResultsAttributes) SetObjectSize1k(v int32)`

SetObjectSize1k sets ObjectSize1k field to given value.

### HasObjectSize1k

`func (o *HistoricalFieldResultsAttributes) HasObjectSize1k() bool`

HasObjectSize1k returns a boolean if a field has been set.

### GetObjectSize10k

`func (o *HistoricalFieldResultsAttributes) GetObjectSize10k() int32`

GetObjectSize10k returns the ObjectSize10k field if non-nil, zero value otherwise.

### GetObjectSize10kOk

`func (o *HistoricalFieldResultsAttributes) GetObjectSize10kOk() (*int32, bool)`

GetObjectSize10kOk returns a tuple with the ObjectSize10k field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectSize10k

`func (o *HistoricalFieldResultsAttributes) SetObjectSize10k(v int32)`

SetObjectSize10k sets ObjectSize10k field to given value.

### HasObjectSize10k

`func (o *HistoricalFieldResultsAttributes) HasObjectSize10k() bool`

HasObjectSize10k returns a boolean if a field has been set.

### GetObjectSize100k

`func (o *HistoricalFieldResultsAttributes) GetObjectSize100k() int32`

GetObjectSize100k returns the ObjectSize100k field if non-nil, zero value otherwise.

### GetObjectSize100kOk

`func (o *HistoricalFieldResultsAttributes) GetObjectSize100kOk() (*int32, bool)`

GetObjectSize100kOk returns a tuple with the ObjectSize100k field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectSize100k

`func (o *HistoricalFieldResultsAttributes) SetObjectSize100k(v int32)`

SetObjectSize100k sets ObjectSize100k field to given value.

### HasObjectSize100k

`func (o *HistoricalFieldResultsAttributes) HasObjectSize100k() bool`

HasObjectSize100k returns a boolean if a field has been set.

### GetObjectSize1m

`func (o *HistoricalFieldResultsAttributes) GetObjectSize1m() int32`

GetObjectSize1m returns the ObjectSize1m field if non-nil, zero value otherwise.

### GetObjectSize1mOk

`func (o *HistoricalFieldResultsAttributes) GetObjectSize1mOk() (*int32, bool)`

GetObjectSize1mOk returns a tuple with the ObjectSize1m field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectSize1m

`func (o *HistoricalFieldResultsAttributes) SetObjectSize1m(v int32)`

SetObjectSize1m sets ObjectSize1m field to given value.

### HasObjectSize1m

`func (o *HistoricalFieldResultsAttributes) HasObjectSize1m() bool`

HasObjectSize1m returns a boolean if a field has been set.

### GetObjectSize10m

`func (o *HistoricalFieldResultsAttributes) GetObjectSize10m() int32`

GetObjectSize10m returns the ObjectSize10m field if non-nil, zero value otherwise.

### GetObjectSize10mOk

`func (o *HistoricalFieldResultsAttributes) GetObjectSize10mOk() (*int32, bool)`

GetObjectSize10mOk returns a tuple with the ObjectSize10m field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectSize10m

`func (o *HistoricalFieldResultsAttributes) SetObjectSize10m(v int32)`

SetObjectSize10m sets ObjectSize10m field to given value.

### HasObjectSize10m

`func (o *HistoricalFieldResultsAttributes) HasObjectSize10m() bool`

HasObjectSize10m returns a boolean if a field has been set.

### GetObjectSize100m

`func (o *HistoricalFieldResultsAttributes) GetObjectSize100m() int32`

GetObjectSize100m returns the ObjectSize100m field if non-nil, zero value otherwise.

### GetObjectSize100mOk

`func (o *HistoricalFieldResultsAttributes) GetObjectSize100mOk() (*int32, bool)`

GetObjectSize100mOk returns a tuple with the ObjectSize100m field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectSize100m

`func (o *HistoricalFieldResultsAttributes) SetObjectSize100m(v int32)`

SetObjectSize100m sets ObjectSize100m field to given value.

### HasObjectSize100m

`func (o *HistoricalFieldResultsAttributes) HasObjectSize100m() bool`

HasObjectSize100m returns a boolean if a field has been set.

### GetObjectSize1g

`func (o *HistoricalFieldResultsAttributes) GetObjectSize1g() int32`

GetObjectSize1g returns the ObjectSize1g field if non-nil, zero value otherwise.

### GetObjectSize1gOk

`func (o *HistoricalFieldResultsAttributes) GetObjectSize1gOk() (*int32, bool)`

GetObjectSize1gOk returns a tuple with the ObjectSize1g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectSize1g

`func (o *HistoricalFieldResultsAttributes) SetObjectSize1g(v int32)`

SetObjectSize1g sets ObjectSize1g field to given value.

### HasObjectSize1g

`func (o *HistoricalFieldResultsAttributes) HasObjectSize1g() bool`

HasObjectSize1g returns a boolean if a field has been set.

### GetRecvSubTime

`func (o *HistoricalFieldResultsAttributes) GetRecvSubTime() float32`

GetRecvSubTime returns the RecvSubTime field if non-nil, zero value otherwise.

### GetRecvSubTimeOk

`func (o *HistoricalFieldResultsAttributes) GetRecvSubTimeOk() (*float32, bool)`

GetRecvSubTimeOk returns a tuple with the RecvSubTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecvSubTime

`func (o *HistoricalFieldResultsAttributes) SetRecvSubTime(v float32)`

SetRecvSubTime sets RecvSubTime field to given value.

### HasRecvSubTime

`func (o *HistoricalFieldResultsAttributes) HasRecvSubTime() bool`

HasRecvSubTime returns a boolean if a field has been set.

### GetRecvSubCount

`func (o *HistoricalFieldResultsAttributes) GetRecvSubCount() int32`

GetRecvSubCount returns the RecvSubCount field if non-nil, zero value otherwise.

### GetRecvSubCountOk

`func (o *HistoricalFieldResultsAttributes) GetRecvSubCountOk() (*int32, bool)`

GetRecvSubCountOk returns a tuple with the RecvSubCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecvSubCount

`func (o *HistoricalFieldResultsAttributes) SetRecvSubCount(v int32)`

SetRecvSubCount sets RecvSubCount field to given value.

### HasRecvSubCount

`func (o *HistoricalFieldResultsAttributes) HasRecvSubCount() bool`

HasRecvSubCount returns a boolean if a field has been set.

### GetHashSubTime

`func (o *HistoricalFieldResultsAttributes) GetHashSubTime() float32`

GetHashSubTime returns the HashSubTime field if non-nil, zero value otherwise.

### GetHashSubTimeOk

`func (o *HistoricalFieldResultsAttributes) GetHashSubTimeOk() (*float32, bool)`

GetHashSubTimeOk returns a tuple with the HashSubTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashSubTime

`func (o *HistoricalFieldResultsAttributes) SetHashSubTime(v float32)`

SetHashSubTime sets HashSubTime field to given value.

### HasHashSubTime

`func (o *HistoricalFieldResultsAttributes) HasHashSubTime() bool`

HasHashSubTime returns a boolean if a field has been set.

### GetHashSubCount

`func (o *HistoricalFieldResultsAttributes) GetHashSubCount() int32`

GetHashSubCount returns the HashSubCount field if non-nil, zero value otherwise.

### GetHashSubCountOk

`func (o *HistoricalFieldResultsAttributes) GetHashSubCountOk() (*int32, bool)`

GetHashSubCountOk returns a tuple with the HashSubCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashSubCount

`func (o *HistoricalFieldResultsAttributes) SetHashSubCount(v int32)`

SetHashSubCount sets HashSubCount field to given value.

### HasHashSubCount

`func (o *HistoricalFieldResultsAttributes) HasHashSubCount() bool`

HasHashSubCount returns a boolean if a field has been set.

### GetMissSubTime

`func (o *HistoricalFieldResultsAttributes) GetMissSubTime() float32`

GetMissSubTime returns the MissSubTime field if non-nil, zero value otherwise.

### GetMissSubTimeOk

`func (o *HistoricalFieldResultsAttributes) GetMissSubTimeOk() (*float32, bool)`

GetMissSubTimeOk returns a tuple with the MissSubTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissSubTime

`func (o *HistoricalFieldResultsAttributes) SetMissSubTime(v float32)`

SetMissSubTime sets MissSubTime field to given value.

### HasMissSubTime

`func (o *HistoricalFieldResultsAttributes) HasMissSubTime() bool`

HasMissSubTime returns a boolean if a field has been set.

### GetMissSubCount

`func (o *HistoricalFieldResultsAttributes) GetMissSubCount() int32`

GetMissSubCount returns the MissSubCount field if non-nil, zero value otherwise.

### GetMissSubCountOk

`func (o *HistoricalFieldResultsAttributes) GetMissSubCountOk() (*int32, bool)`

GetMissSubCountOk returns a tuple with the MissSubCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissSubCount

`func (o *HistoricalFieldResultsAttributes) SetMissSubCount(v int32)`

SetMissSubCount sets MissSubCount field to given value.

### HasMissSubCount

`func (o *HistoricalFieldResultsAttributes) HasMissSubCount() bool`

HasMissSubCount returns a boolean if a field has been set.

### GetFetchSubTime

`func (o *HistoricalFieldResultsAttributes) GetFetchSubTime() float32`

GetFetchSubTime returns the FetchSubTime field if non-nil, zero value otherwise.

### GetFetchSubTimeOk

`func (o *HistoricalFieldResultsAttributes) GetFetchSubTimeOk() (*float32, bool)`

GetFetchSubTimeOk returns a tuple with the FetchSubTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFetchSubTime

`func (o *HistoricalFieldResultsAttributes) SetFetchSubTime(v float32)`

SetFetchSubTime sets FetchSubTime field to given value.

### HasFetchSubTime

`func (o *HistoricalFieldResultsAttributes) HasFetchSubTime() bool`

HasFetchSubTime returns a boolean if a field has been set.

### GetFetchSubCount

`func (o *HistoricalFieldResultsAttributes) GetFetchSubCount() int32`

GetFetchSubCount returns the FetchSubCount field if non-nil, zero value otherwise.

### GetFetchSubCountOk

`func (o *HistoricalFieldResultsAttributes) GetFetchSubCountOk() (*int32, bool)`

GetFetchSubCountOk returns a tuple with the FetchSubCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFetchSubCount

`func (o *HistoricalFieldResultsAttributes) SetFetchSubCount(v int32)`

SetFetchSubCount sets FetchSubCount field to given value.

### HasFetchSubCount

`func (o *HistoricalFieldResultsAttributes) HasFetchSubCount() bool`

HasFetchSubCount returns a boolean if a field has been set.

### GetPassSubTime

`func (o *HistoricalFieldResultsAttributes) GetPassSubTime() float32`

GetPassSubTime returns the PassSubTime field if non-nil, zero value otherwise.

### GetPassSubTimeOk

`func (o *HistoricalFieldResultsAttributes) GetPassSubTimeOk() (*float32, bool)`

GetPassSubTimeOk returns a tuple with the PassSubTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassSubTime

`func (o *HistoricalFieldResultsAttributes) SetPassSubTime(v float32)`

SetPassSubTime sets PassSubTime field to given value.

### HasPassSubTime

`func (o *HistoricalFieldResultsAttributes) HasPassSubTime() bool`

HasPassSubTime returns a boolean if a field has been set.

### GetPassSubCount

`func (o *HistoricalFieldResultsAttributes) GetPassSubCount() int32`

GetPassSubCount returns the PassSubCount field if non-nil, zero value otherwise.

### GetPassSubCountOk

`func (o *HistoricalFieldResultsAttributes) GetPassSubCountOk() (*int32, bool)`

GetPassSubCountOk returns a tuple with the PassSubCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassSubCount

`func (o *HistoricalFieldResultsAttributes) SetPassSubCount(v int32)`

SetPassSubCount sets PassSubCount field to given value.

### HasPassSubCount

`func (o *HistoricalFieldResultsAttributes) HasPassSubCount() bool`

HasPassSubCount returns a boolean if a field has been set.

### GetPipeSubTime

`func (o *HistoricalFieldResultsAttributes) GetPipeSubTime() float32`

GetPipeSubTime returns the PipeSubTime field if non-nil, zero value otherwise.

### GetPipeSubTimeOk

`func (o *HistoricalFieldResultsAttributes) GetPipeSubTimeOk() (*float32, bool)`

GetPipeSubTimeOk returns a tuple with the PipeSubTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipeSubTime

`func (o *HistoricalFieldResultsAttributes) SetPipeSubTime(v float32)`

SetPipeSubTime sets PipeSubTime field to given value.

### HasPipeSubTime

`func (o *HistoricalFieldResultsAttributes) HasPipeSubTime() bool`

HasPipeSubTime returns a boolean if a field has been set.

### GetPipeSubCount

`func (o *HistoricalFieldResultsAttributes) GetPipeSubCount() int32`

GetPipeSubCount returns the PipeSubCount field if non-nil, zero value otherwise.

### GetPipeSubCountOk

`func (o *HistoricalFieldResultsAttributes) GetPipeSubCountOk() (*int32, bool)`

GetPipeSubCountOk returns a tuple with the PipeSubCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipeSubCount

`func (o *HistoricalFieldResultsAttributes) SetPipeSubCount(v int32)`

SetPipeSubCount sets PipeSubCount field to given value.

### HasPipeSubCount

`func (o *HistoricalFieldResultsAttributes) HasPipeSubCount() bool`

HasPipeSubCount returns a boolean if a field has been set.

### GetDeliverSubTime

`func (o *HistoricalFieldResultsAttributes) GetDeliverSubTime() float32`

GetDeliverSubTime returns the DeliverSubTime field if non-nil, zero value otherwise.

### GetDeliverSubTimeOk

`func (o *HistoricalFieldResultsAttributes) GetDeliverSubTimeOk() (*float32, bool)`

GetDeliverSubTimeOk returns a tuple with the DeliverSubTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliverSubTime

`func (o *HistoricalFieldResultsAttributes) SetDeliverSubTime(v float32)`

SetDeliverSubTime sets DeliverSubTime field to given value.

### HasDeliverSubTime

`func (o *HistoricalFieldResultsAttributes) HasDeliverSubTime() bool`

HasDeliverSubTime returns a boolean if a field has been set.

### GetDeliverSubCount

`func (o *HistoricalFieldResultsAttributes) GetDeliverSubCount() int32`

GetDeliverSubCount returns the DeliverSubCount field if non-nil, zero value otherwise.

### GetDeliverSubCountOk

`func (o *HistoricalFieldResultsAttributes) GetDeliverSubCountOk() (*int32, bool)`

GetDeliverSubCountOk returns a tuple with the DeliverSubCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliverSubCount

`func (o *HistoricalFieldResultsAttributes) SetDeliverSubCount(v int32)`

SetDeliverSubCount sets DeliverSubCount field to given value.

### HasDeliverSubCount

`func (o *HistoricalFieldResultsAttributes) HasDeliverSubCount() bool`

HasDeliverSubCount returns a boolean if a field has been set.

### GetErrorSubTime

`func (o *HistoricalFieldResultsAttributes) GetErrorSubTime() float32`

GetErrorSubTime returns the ErrorSubTime field if non-nil, zero value otherwise.

### GetErrorSubTimeOk

`func (o *HistoricalFieldResultsAttributes) GetErrorSubTimeOk() (*float32, bool)`

GetErrorSubTimeOk returns a tuple with the ErrorSubTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorSubTime

`func (o *HistoricalFieldResultsAttributes) SetErrorSubTime(v float32)`

SetErrorSubTime sets ErrorSubTime field to given value.

### HasErrorSubTime

`func (o *HistoricalFieldResultsAttributes) HasErrorSubTime() bool`

HasErrorSubTime returns a boolean if a field has been set.

### GetErrorSubCount

`func (o *HistoricalFieldResultsAttributes) GetErrorSubCount() int32`

GetErrorSubCount returns the ErrorSubCount field if non-nil, zero value otherwise.

### GetErrorSubCountOk

`func (o *HistoricalFieldResultsAttributes) GetErrorSubCountOk() (*int32, bool)`

GetErrorSubCountOk returns a tuple with the ErrorSubCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorSubCount

`func (o *HistoricalFieldResultsAttributes) SetErrorSubCount(v int32)`

SetErrorSubCount sets ErrorSubCount field to given value.

### HasErrorSubCount

`func (o *HistoricalFieldResultsAttributes) HasErrorSubCount() bool`

HasErrorSubCount returns a boolean if a field has been set.

### GetHitSubTime

`func (o *HistoricalFieldResultsAttributes) GetHitSubTime() float32`

GetHitSubTime returns the HitSubTime field if non-nil, zero value otherwise.

### GetHitSubTimeOk

`func (o *HistoricalFieldResultsAttributes) GetHitSubTimeOk() (*float32, bool)`

GetHitSubTimeOk returns a tuple with the HitSubTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHitSubTime

`func (o *HistoricalFieldResultsAttributes) SetHitSubTime(v float32)`

SetHitSubTime sets HitSubTime field to given value.

### HasHitSubTime

`func (o *HistoricalFieldResultsAttributes) HasHitSubTime() bool`

HasHitSubTime returns a boolean if a field has been set.

### GetHitSubCount

`func (o *HistoricalFieldResultsAttributes) GetHitSubCount() int32`

GetHitSubCount returns the HitSubCount field if non-nil, zero value otherwise.

### GetHitSubCountOk

`func (o *HistoricalFieldResultsAttributes) GetHitSubCountOk() (*int32, bool)`

GetHitSubCountOk returns a tuple with the HitSubCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHitSubCount

`func (o *HistoricalFieldResultsAttributes) SetHitSubCount(v int32)`

SetHitSubCount sets HitSubCount field to given value.

### HasHitSubCount

`func (o *HistoricalFieldResultsAttributes) HasHitSubCount() bool`

HasHitSubCount returns a boolean if a field has been set.

### GetPrehashSubTime

`func (o *HistoricalFieldResultsAttributes) GetPrehashSubTime() float32`

GetPrehashSubTime returns the PrehashSubTime field if non-nil, zero value otherwise.

### GetPrehashSubTimeOk

`func (o *HistoricalFieldResultsAttributes) GetPrehashSubTimeOk() (*float32, bool)`

GetPrehashSubTimeOk returns a tuple with the PrehashSubTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrehashSubTime

`func (o *HistoricalFieldResultsAttributes) SetPrehashSubTime(v float32)`

SetPrehashSubTime sets PrehashSubTime field to given value.

### HasPrehashSubTime

`func (o *HistoricalFieldResultsAttributes) HasPrehashSubTime() bool`

HasPrehashSubTime returns a boolean if a field has been set.

### GetPrehashSubCount

`func (o *HistoricalFieldResultsAttributes) GetPrehashSubCount() int32`

GetPrehashSubCount returns the PrehashSubCount field if non-nil, zero value otherwise.

### GetPrehashSubCountOk

`func (o *HistoricalFieldResultsAttributes) GetPrehashSubCountOk() (*int32, bool)`

GetPrehashSubCountOk returns a tuple with the PrehashSubCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrehashSubCount

`func (o *HistoricalFieldResultsAttributes) SetPrehashSubCount(v int32)`

SetPrehashSubCount sets PrehashSubCount field to given value.

### HasPrehashSubCount

`func (o *HistoricalFieldResultsAttributes) HasPrehashSubCount() bool`

HasPrehashSubCount returns a boolean if a field has been set.

### GetPredeliverSubTime

`func (o *HistoricalFieldResultsAttributes) GetPredeliverSubTime() float32`

GetPredeliverSubTime returns the PredeliverSubTime field if non-nil, zero value otherwise.

### GetPredeliverSubTimeOk

`func (o *HistoricalFieldResultsAttributes) GetPredeliverSubTimeOk() (*float32, bool)`

GetPredeliverSubTimeOk returns a tuple with the PredeliverSubTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredeliverSubTime

`func (o *HistoricalFieldResultsAttributes) SetPredeliverSubTime(v float32)`

SetPredeliverSubTime sets PredeliverSubTime field to given value.

### HasPredeliverSubTime

`func (o *HistoricalFieldResultsAttributes) HasPredeliverSubTime() bool`

HasPredeliverSubTime returns a boolean if a field has been set.

### GetPredeliverSubCount

`func (o *HistoricalFieldResultsAttributes) GetPredeliverSubCount() int32`

GetPredeliverSubCount returns the PredeliverSubCount field if non-nil, zero value otherwise.

### GetPredeliverSubCountOk

`func (o *HistoricalFieldResultsAttributes) GetPredeliverSubCountOk() (*int32, bool)`

GetPredeliverSubCountOk returns a tuple with the PredeliverSubCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredeliverSubCount

`func (o *HistoricalFieldResultsAttributes) SetPredeliverSubCount(v int32)`

SetPredeliverSubCount sets PredeliverSubCount field to given value.

### HasPredeliverSubCount

`func (o *HistoricalFieldResultsAttributes) HasPredeliverSubCount() bool`

HasPredeliverSubCount returns a boolean if a field has been set.

### GetTLSHandshakeSentBytes

`func (o *HistoricalFieldResultsAttributes) GetTLSHandshakeSentBytes() int32`

GetTLSHandshakeSentBytes returns the TLSHandshakeSentBytes field if non-nil, zero value otherwise.

### GetTLSHandshakeSentBytesOk

`func (o *HistoricalFieldResultsAttributes) GetTLSHandshakeSentBytesOk() (*int32, bool)`

GetTLSHandshakeSentBytesOk returns a tuple with the TLSHandshakeSentBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTLSHandshakeSentBytes

`func (o *HistoricalFieldResultsAttributes) SetTLSHandshakeSentBytes(v int32)`

SetTLSHandshakeSentBytes sets TLSHandshakeSentBytes field to given value.

### HasTLSHandshakeSentBytes

`func (o *HistoricalFieldResultsAttributes) HasTLSHandshakeSentBytes() bool`

HasTLSHandshakeSentBytes returns a boolean if a field has been set.

### GetHitRespBodyBytes

`func (o *HistoricalFieldResultsAttributes) GetHitRespBodyBytes() int32`

GetHitRespBodyBytes returns the HitRespBodyBytes field if non-nil, zero value otherwise.

### GetHitRespBodyBytesOk

`func (o *HistoricalFieldResultsAttributes) GetHitRespBodyBytesOk() (*int32, bool)`

GetHitRespBodyBytesOk returns a tuple with the HitRespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHitRespBodyBytes

`func (o *HistoricalFieldResultsAttributes) SetHitRespBodyBytes(v int32)`

SetHitRespBodyBytes sets HitRespBodyBytes field to given value.

### HasHitRespBodyBytes

`func (o *HistoricalFieldResultsAttributes) HasHitRespBodyBytes() bool`

HasHitRespBodyBytes returns a boolean if a field has been set.

### GetMissRespBodyBytes

`func (o *HistoricalFieldResultsAttributes) GetMissRespBodyBytes() int32`

GetMissRespBodyBytes returns the MissRespBodyBytes field if non-nil, zero value otherwise.

### GetMissRespBodyBytesOk

`func (o *HistoricalFieldResultsAttributes) GetMissRespBodyBytesOk() (*int32, bool)`

GetMissRespBodyBytesOk returns a tuple with the MissRespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissRespBodyBytes

`func (o *HistoricalFieldResultsAttributes) SetMissRespBodyBytes(v int32)`

SetMissRespBodyBytes sets MissRespBodyBytes field to given value.

### HasMissRespBodyBytes

`func (o *HistoricalFieldResultsAttributes) HasMissRespBodyBytes() bool`

HasMissRespBodyBytes returns a boolean if a field has been set.

### GetPassRespBodyBytes

`func (o *HistoricalFieldResultsAttributes) GetPassRespBodyBytes() int32`

GetPassRespBodyBytes returns the PassRespBodyBytes field if non-nil, zero value otherwise.

### GetPassRespBodyBytesOk

`func (o *HistoricalFieldResultsAttributes) GetPassRespBodyBytesOk() (*int32, bool)`

GetPassRespBodyBytesOk returns a tuple with the PassRespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassRespBodyBytes

`func (o *HistoricalFieldResultsAttributes) SetPassRespBodyBytes(v int32)`

SetPassRespBodyBytes sets PassRespBodyBytes field to given value.

### HasPassRespBodyBytes

`func (o *HistoricalFieldResultsAttributes) HasPassRespBodyBytes() bool`

HasPassRespBodyBytes returns a boolean if a field has been set.

### GetSegblockOriginFetches

`func (o *HistoricalFieldResultsAttributes) GetSegblockOriginFetches() int32`

GetSegblockOriginFetches returns the SegblockOriginFetches field if non-nil, zero value otherwise.

### GetSegblockOriginFetchesOk

`func (o *HistoricalFieldResultsAttributes) GetSegblockOriginFetchesOk() (*int32, bool)`

GetSegblockOriginFetchesOk returns a tuple with the SegblockOriginFetches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegblockOriginFetches

`func (o *HistoricalFieldResultsAttributes) SetSegblockOriginFetches(v int32)`

SetSegblockOriginFetches sets SegblockOriginFetches field to given value.

### HasSegblockOriginFetches

`func (o *HistoricalFieldResultsAttributes) HasSegblockOriginFetches() bool`

HasSegblockOriginFetches returns a boolean if a field has been set.

### GetSegblockShieldFetches

`func (o *HistoricalFieldResultsAttributes) GetSegblockShieldFetches() int32`

GetSegblockShieldFetches returns the SegblockShieldFetches field if non-nil, zero value otherwise.

### GetSegblockShieldFetchesOk

`func (o *HistoricalFieldResultsAttributes) GetSegblockShieldFetchesOk() (*int32, bool)`

GetSegblockShieldFetchesOk returns a tuple with the SegblockShieldFetches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegblockShieldFetches

`func (o *HistoricalFieldResultsAttributes) SetSegblockShieldFetches(v int32)`

SetSegblockShieldFetches sets SegblockShieldFetches field to given value.

### HasSegblockShieldFetches

`func (o *HistoricalFieldResultsAttributes) HasSegblockShieldFetches() bool`

HasSegblockShieldFetches returns a boolean if a field has been set.

### GetComputeRequests

`func (o *HistoricalFieldResultsAttributes) GetComputeRequests() int32`

GetComputeRequests returns the ComputeRequests field if non-nil, zero value otherwise.

### GetComputeRequestsOk

`func (o *HistoricalFieldResultsAttributes) GetComputeRequestsOk() (*int32, bool)`

GetComputeRequestsOk returns a tuple with the ComputeRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRequests

`func (o *HistoricalFieldResultsAttributes) SetComputeRequests(v int32)`

SetComputeRequests sets ComputeRequests field to given value.

### HasComputeRequests

`func (o *HistoricalFieldResultsAttributes) HasComputeRequests() bool`

HasComputeRequests returns a boolean if a field has been set.

### GetComputeRequestTimeMs

`func (o *HistoricalFieldResultsAttributes) GetComputeRequestTimeMs() float32`

GetComputeRequestTimeMs returns the ComputeRequestTimeMs field if non-nil, zero value otherwise.

### GetComputeRequestTimeMsOk

`func (o *HistoricalFieldResultsAttributes) GetComputeRequestTimeMsOk() (*float32, bool)`

GetComputeRequestTimeMsOk returns a tuple with the ComputeRequestTimeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRequestTimeMs

`func (o *HistoricalFieldResultsAttributes) SetComputeRequestTimeMs(v float32)`

SetComputeRequestTimeMs sets ComputeRequestTimeMs field to given value.

### HasComputeRequestTimeMs

`func (o *HistoricalFieldResultsAttributes) HasComputeRequestTimeMs() bool`

HasComputeRequestTimeMs returns a boolean if a field has been set.

### GetComputeRequestTimeBilledMs

`func (o *HistoricalFieldResultsAttributes) GetComputeRequestTimeBilledMs() float32`

GetComputeRequestTimeBilledMs returns the ComputeRequestTimeBilledMs field if non-nil, zero value otherwise.

### GetComputeRequestTimeBilledMsOk

`func (o *HistoricalFieldResultsAttributes) GetComputeRequestTimeBilledMsOk() (*float32, bool)`

GetComputeRequestTimeBilledMsOk returns a tuple with the ComputeRequestTimeBilledMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRequestTimeBilledMs

`func (o *HistoricalFieldResultsAttributes) SetComputeRequestTimeBilledMs(v float32)`

SetComputeRequestTimeBilledMs sets ComputeRequestTimeBilledMs field to given value.

### HasComputeRequestTimeBilledMs

`func (o *HistoricalFieldResultsAttributes) HasComputeRequestTimeBilledMs() bool`

HasComputeRequestTimeBilledMs returns a boolean if a field has been set.

### GetComputeRAMUsed

`func (o *HistoricalFieldResultsAttributes) GetComputeRAMUsed() int32`

GetComputeRAMUsed returns the ComputeRAMUsed field if non-nil, zero value otherwise.

### GetComputeRAMUsedOk

`func (o *HistoricalFieldResultsAttributes) GetComputeRAMUsedOk() (*int32, bool)`

GetComputeRAMUsedOk returns a tuple with the ComputeRAMUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRAMUsed

`func (o *HistoricalFieldResultsAttributes) SetComputeRAMUsed(v int32)`

SetComputeRAMUsed sets ComputeRAMUsed field to given value.

### HasComputeRAMUsed

`func (o *HistoricalFieldResultsAttributes) HasComputeRAMUsed() bool`

HasComputeRAMUsed returns a boolean if a field has been set.

### GetComputeExecutionTimeMs

`func (o *HistoricalFieldResultsAttributes) GetComputeExecutionTimeMs() float32`

GetComputeExecutionTimeMs returns the ComputeExecutionTimeMs field if non-nil, zero value otherwise.

### GetComputeExecutionTimeMsOk

`func (o *HistoricalFieldResultsAttributes) GetComputeExecutionTimeMsOk() (*float32, bool)`

GetComputeExecutionTimeMsOk returns a tuple with the ComputeExecutionTimeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeExecutionTimeMs

`func (o *HistoricalFieldResultsAttributes) SetComputeExecutionTimeMs(v float32)`

SetComputeExecutionTimeMs sets ComputeExecutionTimeMs field to given value.

### HasComputeExecutionTimeMs

`func (o *HistoricalFieldResultsAttributes) HasComputeExecutionTimeMs() bool`

HasComputeExecutionTimeMs returns a boolean if a field has been set.

### GetComputeReqHeaderBytes

`func (o *HistoricalFieldResultsAttributes) GetComputeReqHeaderBytes() int32`

GetComputeReqHeaderBytes returns the ComputeReqHeaderBytes field if non-nil, zero value otherwise.

### GetComputeReqHeaderBytesOk

`func (o *HistoricalFieldResultsAttributes) GetComputeReqHeaderBytesOk() (*int32, bool)`

GetComputeReqHeaderBytesOk returns a tuple with the ComputeReqHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeReqHeaderBytes

`func (o *HistoricalFieldResultsAttributes) SetComputeReqHeaderBytes(v int32)`

SetComputeReqHeaderBytes sets ComputeReqHeaderBytes field to given value.

### HasComputeReqHeaderBytes

`func (o *HistoricalFieldResultsAttributes) HasComputeReqHeaderBytes() bool`

HasComputeReqHeaderBytes returns a boolean if a field has been set.

### GetComputeReqBodyBytes

`func (o *HistoricalFieldResultsAttributes) GetComputeReqBodyBytes() int32`

GetComputeReqBodyBytes returns the ComputeReqBodyBytes field if non-nil, zero value otherwise.

### GetComputeReqBodyBytesOk

`func (o *HistoricalFieldResultsAttributes) GetComputeReqBodyBytesOk() (*int32, bool)`

GetComputeReqBodyBytesOk returns a tuple with the ComputeReqBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeReqBodyBytes

`func (o *HistoricalFieldResultsAttributes) SetComputeReqBodyBytes(v int32)`

SetComputeReqBodyBytes sets ComputeReqBodyBytes field to given value.

### HasComputeReqBodyBytes

`func (o *HistoricalFieldResultsAttributes) HasComputeReqBodyBytes() bool`

HasComputeReqBodyBytes returns a boolean if a field has been set.

### GetComputeRespHeaderBytes

`func (o *HistoricalFieldResultsAttributes) GetComputeRespHeaderBytes() int32`

GetComputeRespHeaderBytes returns the ComputeRespHeaderBytes field if non-nil, zero value otherwise.

### GetComputeRespHeaderBytesOk

`func (o *HistoricalFieldResultsAttributes) GetComputeRespHeaderBytesOk() (*int32, bool)`

GetComputeRespHeaderBytesOk returns a tuple with the ComputeRespHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespHeaderBytes

`func (o *HistoricalFieldResultsAttributes) SetComputeRespHeaderBytes(v int32)`

SetComputeRespHeaderBytes sets ComputeRespHeaderBytes field to given value.

### HasComputeRespHeaderBytes

`func (o *HistoricalFieldResultsAttributes) HasComputeRespHeaderBytes() bool`

HasComputeRespHeaderBytes returns a boolean if a field has been set.

### GetComputeRespBodyBytes

`func (o *HistoricalFieldResultsAttributes) GetComputeRespBodyBytes() int32`

GetComputeRespBodyBytes returns the ComputeRespBodyBytes field if non-nil, zero value otherwise.

### GetComputeRespBodyBytesOk

`func (o *HistoricalFieldResultsAttributes) GetComputeRespBodyBytesOk() (*int32, bool)`

GetComputeRespBodyBytesOk returns a tuple with the ComputeRespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespBodyBytes

`func (o *HistoricalFieldResultsAttributes) SetComputeRespBodyBytes(v int32)`

SetComputeRespBodyBytes sets ComputeRespBodyBytes field to given value.

### HasComputeRespBodyBytes

`func (o *HistoricalFieldResultsAttributes) HasComputeRespBodyBytes() bool`

HasComputeRespBodyBytes returns a boolean if a field has been set.

### GetComputeRespStatus1xx

`func (o *HistoricalFieldResultsAttributes) GetComputeRespStatus1xx() int32`

GetComputeRespStatus1xx returns the ComputeRespStatus1xx field if non-nil, zero value otherwise.

### GetComputeRespStatus1xxOk

`func (o *HistoricalFieldResultsAttributes) GetComputeRespStatus1xxOk() (*int32, bool)`

GetComputeRespStatus1xxOk returns a tuple with the ComputeRespStatus1xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespStatus1xx

`func (o *HistoricalFieldResultsAttributes) SetComputeRespStatus1xx(v int32)`

SetComputeRespStatus1xx sets ComputeRespStatus1xx field to given value.

### HasComputeRespStatus1xx

`func (o *HistoricalFieldResultsAttributes) HasComputeRespStatus1xx() bool`

HasComputeRespStatus1xx returns a boolean if a field has been set.

### GetComputeRespStatus2xx

`func (o *HistoricalFieldResultsAttributes) GetComputeRespStatus2xx() int32`

GetComputeRespStatus2xx returns the ComputeRespStatus2xx field if non-nil, zero value otherwise.

### GetComputeRespStatus2xxOk

`func (o *HistoricalFieldResultsAttributes) GetComputeRespStatus2xxOk() (*int32, bool)`

GetComputeRespStatus2xxOk returns a tuple with the ComputeRespStatus2xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespStatus2xx

`func (o *HistoricalFieldResultsAttributes) SetComputeRespStatus2xx(v int32)`

SetComputeRespStatus2xx sets ComputeRespStatus2xx field to given value.

### HasComputeRespStatus2xx

`func (o *HistoricalFieldResultsAttributes) HasComputeRespStatus2xx() bool`

HasComputeRespStatus2xx returns a boolean if a field has been set.

### GetComputeRespStatus3xx

`func (o *HistoricalFieldResultsAttributes) GetComputeRespStatus3xx() int32`

GetComputeRespStatus3xx returns the ComputeRespStatus3xx field if non-nil, zero value otherwise.

### GetComputeRespStatus3xxOk

`func (o *HistoricalFieldResultsAttributes) GetComputeRespStatus3xxOk() (*int32, bool)`

GetComputeRespStatus3xxOk returns a tuple with the ComputeRespStatus3xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespStatus3xx

`func (o *HistoricalFieldResultsAttributes) SetComputeRespStatus3xx(v int32)`

SetComputeRespStatus3xx sets ComputeRespStatus3xx field to given value.

### HasComputeRespStatus3xx

`func (o *HistoricalFieldResultsAttributes) HasComputeRespStatus3xx() bool`

HasComputeRespStatus3xx returns a boolean if a field has been set.

### GetComputeRespStatus4xx

`func (o *HistoricalFieldResultsAttributes) GetComputeRespStatus4xx() int32`

GetComputeRespStatus4xx returns the ComputeRespStatus4xx field if non-nil, zero value otherwise.

### GetComputeRespStatus4xxOk

`func (o *HistoricalFieldResultsAttributes) GetComputeRespStatus4xxOk() (*int32, bool)`

GetComputeRespStatus4xxOk returns a tuple with the ComputeRespStatus4xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespStatus4xx

`func (o *HistoricalFieldResultsAttributes) SetComputeRespStatus4xx(v int32)`

SetComputeRespStatus4xx sets ComputeRespStatus4xx field to given value.

### HasComputeRespStatus4xx

`func (o *HistoricalFieldResultsAttributes) HasComputeRespStatus4xx() bool`

HasComputeRespStatus4xx returns a boolean if a field has been set.

### GetComputeRespStatus5xx

`func (o *HistoricalFieldResultsAttributes) GetComputeRespStatus5xx() int32`

GetComputeRespStatus5xx returns the ComputeRespStatus5xx field if non-nil, zero value otherwise.

### GetComputeRespStatus5xxOk

`func (o *HistoricalFieldResultsAttributes) GetComputeRespStatus5xxOk() (*int32, bool)`

GetComputeRespStatus5xxOk returns a tuple with the ComputeRespStatus5xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespStatus5xx

`func (o *HistoricalFieldResultsAttributes) SetComputeRespStatus5xx(v int32)`

SetComputeRespStatus5xx sets ComputeRespStatus5xx field to given value.

### HasComputeRespStatus5xx

`func (o *HistoricalFieldResultsAttributes) HasComputeRespStatus5xx() bool`

HasComputeRespStatus5xx returns a boolean if a field has been set.

### GetComputeBereqHeaderBytes

`func (o *HistoricalFieldResultsAttributes) GetComputeBereqHeaderBytes() int32`

GetComputeBereqHeaderBytes returns the ComputeBereqHeaderBytes field if non-nil, zero value otherwise.

### GetComputeBereqHeaderBytesOk

`func (o *HistoricalFieldResultsAttributes) GetComputeBereqHeaderBytesOk() (*int32, bool)`

GetComputeBereqHeaderBytesOk returns a tuple with the ComputeBereqHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBereqHeaderBytes

`func (o *HistoricalFieldResultsAttributes) SetComputeBereqHeaderBytes(v int32)`

SetComputeBereqHeaderBytes sets ComputeBereqHeaderBytes field to given value.

### HasComputeBereqHeaderBytes

`func (o *HistoricalFieldResultsAttributes) HasComputeBereqHeaderBytes() bool`

HasComputeBereqHeaderBytes returns a boolean if a field has been set.

### GetComputeBereqBodyBytes

`func (o *HistoricalFieldResultsAttributes) GetComputeBereqBodyBytes() int32`

GetComputeBereqBodyBytes returns the ComputeBereqBodyBytes field if non-nil, zero value otherwise.

### GetComputeBereqBodyBytesOk

`func (o *HistoricalFieldResultsAttributes) GetComputeBereqBodyBytesOk() (*int32, bool)`

GetComputeBereqBodyBytesOk returns a tuple with the ComputeBereqBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBereqBodyBytes

`func (o *HistoricalFieldResultsAttributes) SetComputeBereqBodyBytes(v int32)`

SetComputeBereqBodyBytes sets ComputeBereqBodyBytes field to given value.

### HasComputeBereqBodyBytes

`func (o *HistoricalFieldResultsAttributes) HasComputeBereqBodyBytes() bool`

HasComputeBereqBodyBytes returns a boolean if a field has been set.

### GetComputeBerespHeaderBytes

`func (o *HistoricalFieldResultsAttributes) GetComputeBerespHeaderBytes() int32`

GetComputeBerespHeaderBytes returns the ComputeBerespHeaderBytes field if non-nil, zero value otherwise.

### GetComputeBerespHeaderBytesOk

`func (o *HistoricalFieldResultsAttributes) GetComputeBerespHeaderBytesOk() (*int32, bool)`

GetComputeBerespHeaderBytesOk returns a tuple with the ComputeBerespHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBerespHeaderBytes

`func (o *HistoricalFieldResultsAttributes) SetComputeBerespHeaderBytes(v int32)`

SetComputeBerespHeaderBytes sets ComputeBerespHeaderBytes field to given value.

### HasComputeBerespHeaderBytes

`func (o *HistoricalFieldResultsAttributes) HasComputeBerespHeaderBytes() bool`

HasComputeBerespHeaderBytes returns a boolean if a field has been set.

### GetComputeBerespBodyBytes

`func (o *HistoricalFieldResultsAttributes) GetComputeBerespBodyBytes() int32`

GetComputeBerespBodyBytes returns the ComputeBerespBodyBytes field if non-nil, zero value otherwise.

### GetComputeBerespBodyBytesOk

`func (o *HistoricalFieldResultsAttributes) GetComputeBerespBodyBytesOk() (*int32, bool)`

GetComputeBerespBodyBytesOk returns a tuple with the ComputeBerespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBerespBodyBytes

`func (o *HistoricalFieldResultsAttributes) SetComputeBerespBodyBytes(v int32)`

SetComputeBerespBodyBytes sets ComputeBerespBodyBytes field to given value.

### HasComputeBerespBodyBytes

`func (o *HistoricalFieldResultsAttributes) HasComputeBerespBodyBytes() bool`

HasComputeBerespBodyBytes returns a boolean if a field has been set.

### GetComputeBereqs

`func (o *HistoricalFieldResultsAttributes) GetComputeBereqs() int32`

GetComputeBereqs returns the ComputeBereqs field if non-nil, zero value otherwise.

### GetComputeBereqsOk

`func (o *HistoricalFieldResultsAttributes) GetComputeBereqsOk() (*int32, bool)`

GetComputeBereqsOk returns a tuple with the ComputeBereqs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBereqs

`func (o *HistoricalFieldResultsAttributes) SetComputeBereqs(v int32)`

SetComputeBereqs sets ComputeBereqs field to given value.

### HasComputeBereqs

`func (o *HistoricalFieldResultsAttributes) HasComputeBereqs() bool`

HasComputeBereqs returns a boolean if a field has been set.

### GetComputeBereqErrors

`func (o *HistoricalFieldResultsAttributes) GetComputeBereqErrors() int32`

GetComputeBereqErrors returns the ComputeBereqErrors field if non-nil, zero value otherwise.

### GetComputeBereqErrorsOk

`func (o *HistoricalFieldResultsAttributes) GetComputeBereqErrorsOk() (*int32, bool)`

GetComputeBereqErrorsOk returns a tuple with the ComputeBereqErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBereqErrors

`func (o *HistoricalFieldResultsAttributes) SetComputeBereqErrors(v int32)`

SetComputeBereqErrors sets ComputeBereqErrors field to given value.

### HasComputeBereqErrors

`func (o *HistoricalFieldResultsAttributes) HasComputeBereqErrors() bool`

HasComputeBereqErrors returns a boolean if a field has been set.

### GetComputeResourceLimitExceeded

`func (o *HistoricalFieldResultsAttributes) GetComputeResourceLimitExceeded() int32`

GetComputeResourceLimitExceeded returns the ComputeResourceLimitExceeded field if non-nil, zero value otherwise.

### GetComputeResourceLimitExceededOk

`func (o *HistoricalFieldResultsAttributes) GetComputeResourceLimitExceededOk() (*int32, bool)`

GetComputeResourceLimitExceededOk returns a tuple with the ComputeResourceLimitExceeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeResourceLimitExceeded

`func (o *HistoricalFieldResultsAttributes) SetComputeResourceLimitExceeded(v int32)`

SetComputeResourceLimitExceeded sets ComputeResourceLimitExceeded field to given value.

### HasComputeResourceLimitExceeded

`func (o *HistoricalFieldResultsAttributes) HasComputeResourceLimitExceeded() bool`

HasComputeResourceLimitExceeded returns a boolean if a field has been set.

### GetComputeHeapLimitExceeded

`func (o *HistoricalFieldResultsAttributes) GetComputeHeapLimitExceeded() int32`

GetComputeHeapLimitExceeded returns the ComputeHeapLimitExceeded field if non-nil, zero value otherwise.

### GetComputeHeapLimitExceededOk

`func (o *HistoricalFieldResultsAttributes) GetComputeHeapLimitExceededOk() (*int32, bool)`

GetComputeHeapLimitExceededOk returns a tuple with the ComputeHeapLimitExceeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeHeapLimitExceeded

`func (o *HistoricalFieldResultsAttributes) SetComputeHeapLimitExceeded(v int32)`

SetComputeHeapLimitExceeded sets ComputeHeapLimitExceeded field to given value.

### HasComputeHeapLimitExceeded

`func (o *HistoricalFieldResultsAttributes) HasComputeHeapLimitExceeded() bool`

HasComputeHeapLimitExceeded returns a boolean if a field has been set.

### GetComputeStackLimitExceeded

`func (o *HistoricalFieldResultsAttributes) GetComputeStackLimitExceeded() int32`

GetComputeStackLimitExceeded returns the ComputeStackLimitExceeded field if non-nil, zero value otherwise.

### GetComputeStackLimitExceededOk

`func (o *HistoricalFieldResultsAttributes) GetComputeStackLimitExceededOk() (*int32, bool)`

GetComputeStackLimitExceededOk returns a tuple with the ComputeStackLimitExceeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeStackLimitExceeded

`func (o *HistoricalFieldResultsAttributes) SetComputeStackLimitExceeded(v int32)`

SetComputeStackLimitExceeded sets ComputeStackLimitExceeded field to given value.

### HasComputeStackLimitExceeded

`func (o *HistoricalFieldResultsAttributes) HasComputeStackLimitExceeded() bool`

HasComputeStackLimitExceeded returns a boolean if a field has been set.

### GetComputeGlobalsLimitExceeded

`func (o *HistoricalFieldResultsAttributes) GetComputeGlobalsLimitExceeded() int32`

GetComputeGlobalsLimitExceeded returns the ComputeGlobalsLimitExceeded field if non-nil, zero value otherwise.

### GetComputeGlobalsLimitExceededOk

`func (o *HistoricalFieldResultsAttributes) GetComputeGlobalsLimitExceededOk() (*int32, bool)`

GetComputeGlobalsLimitExceededOk returns a tuple with the ComputeGlobalsLimitExceeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeGlobalsLimitExceeded

`func (o *HistoricalFieldResultsAttributes) SetComputeGlobalsLimitExceeded(v int32)`

SetComputeGlobalsLimitExceeded sets ComputeGlobalsLimitExceeded field to given value.

### HasComputeGlobalsLimitExceeded

`func (o *HistoricalFieldResultsAttributes) HasComputeGlobalsLimitExceeded() bool`

HasComputeGlobalsLimitExceeded returns a boolean if a field has been set.

### GetComputeGuestErrors

`func (o *HistoricalFieldResultsAttributes) GetComputeGuestErrors() int32`

GetComputeGuestErrors returns the ComputeGuestErrors field if non-nil, zero value otherwise.

### GetComputeGuestErrorsOk

`func (o *HistoricalFieldResultsAttributes) GetComputeGuestErrorsOk() (*int32, bool)`

GetComputeGuestErrorsOk returns a tuple with the ComputeGuestErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeGuestErrors

`func (o *HistoricalFieldResultsAttributes) SetComputeGuestErrors(v int32)`

SetComputeGuestErrors sets ComputeGuestErrors field to given value.

### HasComputeGuestErrors

`func (o *HistoricalFieldResultsAttributes) HasComputeGuestErrors() bool`

HasComputeGuestErrors returns a boolean if a field has been set.

### GetComputeRuntimeErrors

`func (o *HistoricalFieldResultsAttributes) GetComputeRuntimeErrors() int32`

GetComputeRuntimeErrors returns the ComputeRuntimeErrors field if non-nil, zero value otherwise.

### GetComputeRuntimeErrorsOk

`func (o *HistoricalFieldResultsAttributes) GetComputeRuntimeErrorsOk() (*int32, bool)`

GetComputeRuntimeErrorsOk returns a tuple with the ComputeRuntimeErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRuntimeErrors

`func (o *HistoricalFieldResultsAttributes) SetComputeRuntimeErrors(v int32)`

SetComputeRuntimeErrors sets ComputeRuntimeErrors field to given value.

### HasComputeRuntimeErrors

`func (o *HistoricalFieldResultsAttributes) HasComputeRuntimeErrors() bool`

HasComputeRuntimeErrors returns a boolean if a field has been set.

### GetEdgeHitRespBodyBytes

`func (o *HistoricalFieldResultsAttributes) GetEdgeHitRespBodyBytes() int32`

GetEdgeHitRespBodyBytes returns the EdgeHitRespBodyBytes field if non-nil, zero value otherwise.

### GetEdgeHitRespBodyBytesOk

`func (o *HistoricalFieldResultsAttributes) GetEdgeHitRespBodyBytesOk() (*int32, bool)`

GetEdgeHitRespBodyBytesOk returns a tuple with the EdgeHitRespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeHitRespBodyBytes

`func (o *HistoricalFieldResultsAttributes) SetEdgeHitRespBodyBytes(v int32)`

SetEdgeHitRespBodyBytes sets EdgeHitRespBodyBytes field to given value.

### HasEdgeHitRespBodyBytes

`func (o *HistoricalFieldResultsAttributes) HasEdgeHitRespBodyBytes() bool`

HasEdgeHitRespBodyBytes returns a boolean if a field has been set.

### GetEdgeHitRespHeaderBytes

`func (o *HistoricalFieldResultsAttributes) GetEdgeHitRespHeaderBytes() int32`

GetEdgeHitRespHeaderBytes returns the EdgeHitRespHeaderBytes field if non-nil, zero value otherwise.

### GetEdgeHitRespHeaderBytesOk

`func (o *HistoricalFieldResultsAttributes) GetEdgeHitRespHeaderBytesOk() (*int32, bool)`

GetEdgeHitRespHeaderBytesOk returns a tuple with the EdgeHitRespHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeHitRespHeaderBytes

`func (o *HistoricalFieldResultsAttributes) SetEdgeHitRespHeaderBytes(v int32)`

SetEdgeHitRespHeaderBytes sets EdgeHitRespHeaderBytes field to given value.

### HasEdgeHitRespHeaderBytes

`func (o *HistoricalFieldResultsAttributes) HasEdgeHitRespHeaderBytes() bool`

HasEdgeHitRespHeaderBytes returns a boolean if a field has been set.

### GetEdgeMissRespBodyBytes

`func (o *HistoricalFieldResultsAttributes) GetEdgeMissRespBodyBytes() int32`

GetEdgeMissRespBodyBytes returns the EdgeMissRespBodyBytes field if non-nil, zero value otherwise.

### GetEdgeMissRespBodyBytesOk

`func (o *HistoricalFieldResultsAttributes) GetEdgeMissRespBodyBytesOk() (*int32, bool)`

GetEdgeMissRespBodyBytesOk returns a tuple with the EdgeMissRespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeMissRespBodyBytes

`func (o *HistoricalFieldResultsAttributes) SetEdgeMissRespBodyBytes(v int32)`

SetEdgeMissRespBodyBytes sets EdgeMissRespBodyBytes field to given value.

### HasEdgeMissRespBodyBytes

`func (o *HistoricalFieldResultsAttributes) HasEdgeMissRespBodyBytes() bool`

HasEdgeMissRespBodyBytes returns a boolean if a field has been set.

### GetEdgeMissRespHeaderBytes

`func (o *HistoricalFieldResultsAttributes) GetEdgeMissRespHeaderBytes() int32`

GetEdgeMissRespHeaderBytes returns the EdgeMissRespHeaderBytes field if non-nil, zero value otherwise.

### GetEdgeMissRespHeaderBytesOk

`func (o *HistoricalFieldResultsAttributes) GetEdgeMissRespHeaderBytesOk() (*int32, bool)`

GetEdgeMissRespHeaderBytesOk returns a tuple with the EdgeMissRespHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeMissRespHeaderBytes

`func (o *HistoricalFieldResultsAttributes) SetEdgeMissRespHeaderBytes(v int32)`

SetEdgeMissRespHeaderBytes sets EdgeMissRespHeaderBytes field to given value.

### HasEdgeMissRespHeaderBytes

`func (o *HistoricalFieldResultsAttributes) HasEdgeMissRespHeaderBytes() bool`

HasEdgeMissRespHeaderBytes returns a boolean if a field has been set.

### GetOriginCacheFetchRespBodyBytes

`func (o *HistoricalFieldResultsAttributes) GetOriginCacheFetchRespBodyBytes() int32`

GetOriginCacheFetchRespBodyBytes returns the OriginCacheFetchRespBodyBytes field if non-nil, zero value otherwise.

### GetOriginCacheFetchRespBodyBytesOk

`func (o *HistoricalFieldResultsAttributes) GetOriginCacheFetchRespBodyBytesOk() (*int32, bool)`

GetOriginCacheFetchRespBodyBytesOk returns a tuple with the OriginCacheFetchRespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginCacheFetchRespBodyBytes

`func (o *HistoricalFieldResultsAttributes) SetOriginCacheFetchRespBodyBytes(v int32)`

SetOriginCacheFetchRespBodyBytes sets OriginCacheFetchRespBodyBytes field to given value.

### HasOriginCacheFetchRespBodyBytes

`func (o *HistoricalFieldResultsAttributes) HasOriginCacheFetchRespBodyBytes() bool`

HasOriginCacheFetchRespBodyBytes returns a boolean if a field has been set.

### GetOriginCacheFetchRespHeaderBytes

`func (o *HistoricalFieldResultsAttributes) GetOriginCacheFetchRespHeaderBytes() int32`

GetOriginCacheFetchRespHeaderBytes returns the OriginCacheFetchRespHeaderBytes field if non-nil, zero value otherwise.

### GetOriginCacheFetchRespHeaderBytesOk

`func (o *HistoricalFieldResultsAttributes) GetOriginCacheFetchRespHeaderBytesOk() (*int32, bool)`

GetOriginCacheFetchRespHeaderBytesOk returns a tuple with the OriginCacheFetchRespHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginCacheFetchRespHeaderBytes

`func (o *HistoricalFieldResultsAttributes) SetOriginCacheFetchRespHeaderBytes(v int32)`

SetOriginCacheFetchRespHeaderBytes sets OriginCacheFetchRespHeaderBytes field to given value.

### HasOriginCacheFetchRespHeaderBytes

`func (o *HistoricalFieldResultsAttributes) HasOriginCacheFetchRespHeaderBytes() bool`

HasOriginCacheFetchRespHeaderBytes returns a boolean if a field has been set.

### GetShieldHitRequests

`func (o *HistoricalFieldResultsAttributes) GetShieldHitRequests() int32`

GetShieldHitRequests returns the ShieldHitRequests field if non-nil, zero value otherwise.

### GetShieldHitRequestsOk

`func (o *HistoricalFieldResultsAttributes) GetShieldHitRequestsOk() (*int32, bool)`

GetShieldHitRequestsOk returns a tuple with the ShieldHitRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShieldHitRequests

`func (o *HistoricalFieldResultsAttributes) SetShieldHitRequests(v int32)`

SetShieldHitRequests sets ShieldHitRequests field to given value.

### HasShieldHitRequests

`func (o *HistoricalFieldResultsAttributes) HasShieldHitRequests() bool`

HasShieldHitRequests returns a boolean if a field has been set.

### GetShieldMissRequests

`func (o *HistoricalFieldResultsAttributes) GetShieldMissRequests() int32`

GetShieldMissRequests returns the ShieldMissRequests field if non-nil, zero value otherwise.

### GetShieldMissRequestsOk

`func (o *HistoricalFieldResultsAttributes) GetShieldMissRequestsOk() (*int32, bool)`

GetShieldMissRequestsOk returns a tuple with the ShieldMissRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShieldMissRequests

`func (o *HistoricalFieldResultsAttributes) SetShieldMissRequests(v int32)`

SetShieldMissRequests sets ShieldMissRequests field to given value.

### HasShieldMissRequests

`func (o *HistoricalFieldResultsAttributes) HasShieldMissRequests() bool`

HasShieldMissRequests returns a boolean if a field has been set.

### GetShieldHitRespHeaderBytes

`func (o *HistoricalFieldResultsAttributes) GetShieldHitRespHeaderBytes() int32`

GetShieldHitRespHeaderBytes returns the ShieldHitRespHeaderBytes field if non-nil, zero value otherwise.

### GetShieldHitRespHeaderBytesOk

`func (o *HistoricalFieldResultsAttributes) GetShieldHitRespHeaderBytesOk() (*int32, bool)`

GetShieldHitRespHeaderBytesOk returns a tuple with the ShieldHitRespHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShieldHitRespHeaderBytes

`func (o *HistoricalFieldResultsAttributes) SetShieldHitRespHeaderBytes(v int32)`

SetShieldHitRespHeaderBytes sets ShieldHitRespHeaderBytes field to given value.

### HasShieldHitRespHeaderBytes

`func (o *HistoricalFieldResultsAttributes) HasShieldHitRespHeaderBytes() bool`

HasShieldHitRespHeaderBytes returns a boolean if a field has been set.

### GetShieldHitRespBodyBytes

`func (o *HistoricalFieldResultsAttributes) GetShieldHitRespBodyBytes() int32`

GetShieldHitRespBodyBytes returns the ShieldHitRespBodyBytes field if non-nil, zero value otherwise.

### GetShieldHitRespBodyBytesOk

`func (o *HistoricalFieldResultsAttributes) GetShieldHitRespBodyBytesOk() (*int32, bool)`

GetShieldHitRespBodyBytesOk returns a tuple with the ShieldHitRespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShieldHitRespBodyBytes

`func (o *HistoricalFieldResultsAttributes) SetShieldHitRespBodyBytes(v int32)`

SetShieldHitRespBodyBytes sets ShieldHitRespBodyBytes field to given value.

### HasShieldHitRespBodyBytes

`func (o *HistoricalFieldResultsAttributes) HasShieldHitRespBodyBytes() bool`

HasShieldHitRespBodyBytes returns a boolean if a field has been set.

### GetShieldMissRespHeaderBytes

`func (o *HistoricalFieldResultsAttributes) GetShieldMissRespHeaderBytes() int32`

GetShieldMissRespHeaderBytes returns the ShieldMissRespHeaderBytes field if non-nil, zero value otherwise.

### GetShieldMissRespHeaderBytesOk

`func (o *HistoricalFieldResultsAttributes) GetShieldMissRespHeaderBytesOk() (*int32, bool)`

GetShieldMissRespHeaderBytesOk returns a tuple with the ShieldMissRespHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShieldMissRespHeaderBytes

`func (o *HistoricalFieldResultsAttributes) SetShieldMissRespHeaderBytes(v int32)`

SetShieldMissRespHeaderBytes sets ShieldMissRespHeaderBytes field to given value.

### HasShieldMissRespHeaderBytes

`func (o *HistoricalFieldResultsAttributes) HasShieldMissRespHeaderBytes() bool`

HasShieldMissRespHeaderBytes returns a boolean if a field has been set.

### GetShieldMissRespBodyBytes

`func (o *HistoricalFieldResultsAttributes) GetShieldMissRespBodyBytes() int32`

GetShieldMissRespBodyBytes returns the ShieldMissRespBodyBytes field if non-nil, zero value otherwise.

### GetShieldMissRespBodyBytesOk

`func (o *HistoricalFieldResultsAttributes) GetShieldMissRespBodyBytesOk() (*int32, bool)`

GetShieldMissRespBodyBytesOk returns a tuple with the ShieldMissRespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShieldMissRespBodyBytes

`func (o *HistoricalFieldResultsAttributes) SetShieldMissRespBodyBytes(v int32)`

SetShieldMissRespBodyBytes sets ShieldMissRespBodyBytes field to given value.

### HasShieldMissRespBodyBytes

`func (o *HistoricalFieldResultsAttributes) HasShieldMissRespBodyBytes() bool`

HasShieldMissRespBodyBytes returns a boolean if a field has been set.

### GetWebsocketReqHeaderBytes

`func (o *HistoricalFieldResultsAttributes) GetWebsocketReqHeaderBytes() int32`

GetWebsocketReqHeaderBytes returns the WebsocketReqHeaderBytes field if non-nil, zero value otherwise.

### GetWebsocketReqHeaderBytesOk

`func (o *HistoricalFieldResultsAttributes) GetWebsocketReqHeaderBytesOk() (*int32, bool)`

GetWebsocketReqHeaderBytesOk returns a tuple with the WebsocketReqHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsocketReqHeaderBytes

`func (o *HistoricalFieldResultsAttributes) SetWebsocketReqHeaderBytes(v int32)`

SetWebsocketReqHeaderBytes sets WebsocketReqHeaderBytes field to given value.

### HasWebsocketReqHeaderBytes

`func (o *HistoricalFieldResultsAttributes) HasWebsocketReqHeaderBytes() bool`

HasWebsocketReqHeaderBytes returns a boolean if a field has been set.

### GetWebsocketReqBodyBytes

`func (o *HistoricalFieldResultsAttributes) GetWebsocketReqBodyBytes() int32`

GetWebsocketReqBodyBytes returns the WebsocketReqBodyBytes field if non-nil, zero value otherwise.

### GetWebsocketReqBodyBytesOk

`func (o *HistoricalFieldResultsAttributes) GetWebsocketReqBodyBytesOk() (*int32, bool)`

GetWebsocketReqBodyBytesOk returns a tuple with the WebsocketReqBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsocketReqBodyBytes

`func (o *HistoricalFieldResultsAttributes) SetWebsocketReqBodyBytes(v int32)`

SetWebsocketReqBodyBytes sets WebsocketReqBodyBytes field to given value.

### HasWebsocketReqBodyBytes

`func (o *HistoricalFieldResultsAttributes) HasWebsocketReqBodyBytes() bool`

HasWebsocketReqBodyBytes returns a boolean if a field has been set.

### GetWebsocketRespHeaderBytes

`func (o *HistoricalFieldResultsAttributes) GetWebsocketRespHeaderBytes() int32`

GetWebsocketRespHeaderBytes returns the WebsocketRespHeaderBytes field if non-nil, zero value otherwise.

### GetWebsocketRespHeaderBytesOk

`func (o *HistoricalFieldResultsAttributes) GetWebsocketRespHeaderBytesOk() (*int32, bool)`

GetWebsocketRespHeaderBytesOk returns a tuple with the WebsocketRespHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsocketRespHeaderBytes

`func (o *HistoricalFieldResultsAttributes) SetWebsocketRespHeaderBytes(v int32)`

SetWebsocketRespHeaderBytes sets WebsocketRespHeaderBytes field to given value.

### HasWebsocketRespHeaderBytes

`func (o *HistoricalFieldResultsAttributes) HasWebsocketRespHeaderBytes() bool`

HasWebsocketRespHeaderBytes returns a boolean if a field has been set.

### GetWebsocketRespBodyBytes

`func (o *HistoricalFieldResultsAttributes) GetWebsocketRespBodyBytes() int32`

GetWebsocketRespBodyBytes returns the WebsocketRespBodyBytes field if non-nil, zero value otherwise.

### GetWebsocketRespBodyBytesOk

`func (o *HistoricalFieldResultsAttributes) GetWebsocketRespBodyBytesOk() (*int32, bool)`

GetWebsocketRespBodyBytesOk returns a tuple with the WebsocketRespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsocketRespBodyBytes

`func (o *HistoricalFieldResultsAttributes) SetWebsocketRespBodyBytes(v int32)`

SetWebsocketRespBodyBytes sets WebsocketRespBodyBytes field to given value.

### HasWebsocketRespBodyBytes

`func (o *HistoricalFieldResultsAttributes) HasWebsocketRespBodyBytes() bool`

HasWebsocketRespBodyBytes returns a boolean if a field has been set.

### GetWebsocketBereqHeaderBytes

`func (o *HistoricalFieldResultsAttributes) GetWebsocketBereqHeaderBytes() int32`

GetWebsocketBereqHeaderBytes returns the WebsocketBereqHeaderBytes field if non-nil, zero value otherwise.

### GetWebsocketBereqHeaderBytesOk

`func (o *HistoricalFieldResultsAttributes) GetWebsocketBereqHeaderBytesOk() (*int32, bool)`

GetWebsocketBereqHeaderBytesOk returns a tuple with the WebsocketBereqHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsocketBereqHeaderBytes

`func (o *HistoricalFieldResultsAttributes) SetWebsocketBereqHeaderBytes(v int32)`

SetWebsocketBereqHeaderBytes sets WebsocketBereqHeaderBytes field to given value.

### HasWebsocketBereqHeaderBytes

`func (o *HistoricalFieldResultsAttributes) HasWebsocketBereqHeaderBytes() bool`

HasWebsocketBereqHeaderBytes returns a boolean if a field has been set.

### GetWebsocketBereqBodyBytes

`func (o *HistoricalFieldResultsAttributes) GetWebsocketBereqBodyBytes() int32`

GetWebsocketBereqBodyBytes returns the WebsocketBereqBodyBytes field if non-nil, zero value otherwise.

### GetWebsocketBereqBodyBytesOk

`func (o *HistoricalFieldResultsAttributes) GetWebsocketBereqBodyBytesOk() (*int32, bool)`

GetWebsocketBereqBodyBytesOk returns a tuple with the WebsocketBereqBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsocketBereqBodyBytes

`func (o *HistoricalFieldResultsAttributes) SetWebsocketBereqBodyBytes(v int32)`

SetWebsocketBereqBodyBytes sets WebsocketBereqBodyBytes field to given value.

### HasWebsocketBereqBodyBytes

`func (o *HistoricalFieldResultsAttributes) HasWebsocketBereqBodyBytes() bool`

HasWebsocketBereqBodyBytes returns a boolean if a field has been set.

### GetWebsocketBerespHeaderBytes

`func (o *HistoricalFieldResultsAttributes) GetWebsocketBerespHeaderBytes() int32`

GetWebsocketBerespHeaderBytes returns the WebsocketBerespHeaderBytes field if non-nil, zero value otherwise.

### GetWebsocketBerespHeaderBytesOk

`func (o *HistoricalFieldResultsAttributes) GetWebsocketBerespHeaderBytesOk() (*int32, bool)`

GetWebsocketBerespHeaderBytesOk returns a tuple with the WebsocketBerespHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsocketBerespHeaderBytes

`func (o *HistoricalFieldResultsAttributes) SetWebsocketBerespHeaderBytes(v int32)`

SetWebsocketBerespHeaderBytes sets WebsocketBerespHeaderBytes field to given value.

### HasWebsocketBerespHeaderBytes

`func (o *HistoricalFieldResultsAttributes) HasWebsocketBerespHeaderBytes() bool`

HasWebsocketBerespHeaderBytes returns a boolean if a field has been set.

### GetWebsocketBerespBodyBytes

`func (o *HistoricalFieldResultsAttributes) GetWebsocketBerespBodyBytes() int32`

GetWebsocketBerespBodyBytes returns the WebsocketBerespBodyBytes field if non-nil, zero value otherwise.

### GetWebsocketBerespBodyBytesOk

`func (o *HistoricalFieldResultsAttributes) GetWebsocketBerespBodyBytesOk() (*int32, bool)`

GetWebsocketBerespBodyBytesOk returns a tuple with the WebsocketBerespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsocketBerespBodyBytes

`func (o *HistoricalFieldResultsAttributes) SetWebsocketBerespBodyBytes(v int32)`

SetWebsocketBerespBodyBytes sets WebsocketBerespBodyBytes field to given value.

### HasWebsocketBerespBodyBytes

`func (o *HistoricalFieldResultsAttributes) HasWebsocketBerespBodyBytes() bool`

HasWebsocketBerespBodyBytes returns a boolean if a field has been set.

### GetWebsocketConnTimeMs

`func (o *HistoricalFieldResultsAttributes) GetWebsocketConnTimeMs() int32`

GetWebsocketConnTimeMs returns the WebsocketConnTimeMs field if non-nil, zero value otherwise.

### GetWebsocketConnTimeMsOk

`func (o *HistoricalFieldResultsAttributes) GetWebsocketConnTimeMsOk() (*int32, bool)`

GetWebsocketConnTimeMsOk returns a tuple with the WebsocketConnTimeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsocketConnTimeMs

`func (o *HistoricalFieldResultsAttributes) SetWebsocketConnTimeMs(v int32)`

SetWebsocketConnTimeMs sets WebsocketConnTimeMs field to given value.

### HasWebsocketConnTimeMs

`func (o *HistoricalFieldResultsAttributes) HasWebsocketConnTimeMs() bool`

HasWebsocketConnTimeMs returns a boolean if a field has been set.

### GetFanoutRecvPublishes

`func (o *HistoricalFieldResultsAttributes) GetFanoutRecvPublishes() int32`

GetFanoutRecvPublishes returns the FanoutRecvPublishes field if non-nil, zero value otherwise.

### GetFanoutRecvPublishesOk

`func (o *HistoricalFieldResultsAttributes) GetFanoutRecvPublishesOk() (*int32, bool)`

GetFanoutRecvPublishesOk returns a tuple with the FanoutRecvPublishes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFanoutRecvPublishes

`func (o *HistoricalFieldResultsAttributes) SetFanoutRecvPublishes(v int32)`

SetFanoutRecvPublishes sets FanoutRecvPublishes field to given value.

### HasFanoutRecvPublishes

`func (o *HistoricalFieldResultsAttributes) HasFanoutRecvPublishes() bool`

HasFanoutRecvPublishes returns a boolean if a field has been set.

### GetFanoutSendPublishes

`func (o *HistoricalFieldResultsAttributes) GetFanoutSendPublishes() int32`

GetFanoutSendPublishes returns the FanoutSendPublishes field if non-nil, zero value otherwise.

### GetFanoutSendPublishesOk

`func (o *HistoricalFieldResultsAttributes) GetFanoutSendPublishesOk() (*int32, bool)`

GetFanoutSendPublishesOk returns a tuple with the FanoutSendPublishes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFanoutSendPublishes

`func (o *HistoricalFieldResultsAttributes) SetFanoutSendPublishes(v int32)`

SetFanoutSendPublishes sets FanoutSendPublishes field to given value.

### HasFanoutSendPublishes

`func (o *HistoricalFieldResultsAttributes) HasFanoutSendPublishes() bool`

HasFanoutSendPublishes returns a boolean if a field has been set.

### GetKvStoreClassAOperations

`func (o *HistoricalFieldResultsAttributes) GetKvStoreClassAOperations() int32`

GetKvStoreClassAOperations returns the KvStoreClassAOperations field if non-nil, zero value otherwise.

### GetKvStoreClassAOperationsOk

`func (o *HistoricalFieldResultsAttributes) GetKvStoreClassAOperationsOk() (*int32, bool)`

GetKvStoreClassAOperationsOk returns a tuple with the KvStoreClassAOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKvStoreClassAOperations

`func (o *HistoricalFieldResultsAttributes) SetKvStoreClassAOperations(v int32)`

SetKvStoreClassAOperations sets KvStoreClassAOperations field to given value.

### HasKvStoreClassAOperations

`func (o *HistoricalFieldResultsAttributes) HasKvStoreClassAOperations() bool`

HasKvStoreClassAOperations returns a boolean if a field has been set.

### GetKvStoreClassBOperations

`func (o *HistoricalFieldResultsAttributes) GetKvStoreClassBOperations() int32`

GetKvStoreClassBOperations returns the KvStoreClassBOperations field if non-nil, zero value otherwise.

### GetKvStoreClassBOperationsOk

`func (o *HistoricalFieldResultsAttributes) GetKvStoreClassBOperationsOk() (*int32, bool)`

GetKvStoreClassBOperationsOk returns a tuple with the KvStoreClassBOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKvStoreClassBOperations

`func (o *HistoricalFieldResultsAttributes) SetKvStoreClassBOperations(v int32)`

SetKvStoreClassBOperations sets KvStoreClassBOperations field to given value.

### HasKvStoreClassBOperations

`func (o *HistoricalFieldResultsAttributes) HasKvStoreClassBOperations() bool`

HasKvStoreClassBOperations returns a boolean if a field has been set.

### GetObjectStoreClassAOperations

`func (o *HistoricalFieldResultsAttributes) GetObjectStoreClassAOperations() int32`

GetObjectStoreClassAOperations returns the ObjectStoreClassAOperations field if non-nil, zero value otherwise.

### GetObjectStoreClassAOperationsOk

`func (o *HistoricalFieldResultsAttributes) GetObjectStoreClassAOperationsOk() (*int32, bool)`

GetObjectStoreClassAOperationsOk returns a tuple with the ObjectStoreClassAOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectStoreClassAOperations

`func (o *HistoricalFieldResultsAttributes) SetObjectStoreClassAOperations(v int32)`

SetObjectStoreClassAOperations sets ObjectStoreClassAOperations field to given value.

### HasObjectStoreClassAOperations

`func (o *HistoricalFieldResultsAttributes) HasObjectStoreClassAOperations() bool`

HasObjectStoreClassAOperations returns a boolean if a field has been set.

### GetObjectStoreClassBOperations

`func (o *HistoricalFieldResultsAttributes) GetObjectStoreClassBOperations() int32`

GetObjectStoreClassBOperations returns the ObjectStoreClassBOperations field if non-nil, zero value otherwise.

### GetObjectStoreClassBOperationsOk

`func (o *HistoricalFieldResultsAttributes) GetObjectStoreClassBOperationsOk() (*int32, bool)`

GetObjectStoreClassBOperationsOk returns a tuple with the ObjectStoreClassBOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectStoreClassBOperations

`func (o *HistoricalFieldResultsAttributes) SetObjectStoreClassBOperations(v int32)`

SetObjectStoreClassBOperations sets ObjectStoreClassBOperations field to given value.

### HasObjectStoreClassBOperations

`func (o *HistoricalFieldResultsAttributes) HasObjectStoreClassBOperations() bool`

HasObjectStoreClassBOperations returns a boolean if a field has been set.

### GetFanoutReqHeaderBytes

`func (o *HistoricalFieldResultsAttributes) GetFanoutReqHeaderBytes() int32`

GetFanoutReqHeaderBytes returns the FanoutReqHeaderBytes field if non-nil, zero value otherwise.

### GetFanoutReqHeaderBytesOk

`func (o *HistoricalFieldResultsAttributes) GetFanoutReqHeaderBytesOk() (*int32, bool)`

GetFanoutReqHeaderBytesOk returns a tuple with the FanoutReqHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFanoutReqHeaderBytes

`func (o *HistoricalFieldResultsAttributes) SetFanoutReqHeaderBytes(v int32)`

SetFanoutReqHeaderBytes sets FanoutReqHeaderBytes field to given value.

### HasFanoutReqHeaderBytes

`func (o *HistoricalFieldResultsAttributes) HasFanoutReqHeaderBytes() bool`

HasFanoutReqHeaderBytes returns a boolean if a field has been set.

### GetFanoutReqBodyBytes

`func (o *HistoricalFieldResultsAttributes) GetFanoutReqBodyBytes() int32`

GetFanoutReqBodyBytes returns the FanoutReqBodyBytes field if non-nil, zero value otherwise.

### GetFanoutReqBodyBytesOk

`func (o *HistoricalFieldResultsAttributes) GetFanoutReqBodyBytesOk() (*int32, bool)`

GetFanoutReqBodyBytesOk returns a tuple with the FanoutReqBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFanoutReqBodyBytes

`func (o *HistoricalFieldResultsAttributes) SetFanoutReqBodyBytes(v int32)`

SetFanoutReqBodyBytes sets FanoutReqBodyBytes field to given value.

### HasFanoutReqBodyBytes

`func (o *HistoricalFieldResultsAttributes) HasFanoutReqBodyBytes() bool`

HasFanoutReqBodyBytes returns a boolean if a field has been set.

### GetFanoutRespHeaderBytes

`func (o *HistoricalFieldResultsAttributes) GetFanoutRespHeaderBytes() int32`

GetFanoutRespHeaderBytes returns the FanoutRespHeaderBytes field if non-nil, zero value otherwise.

### GetFanoutRespHeaderBytesOk

`func (o *HistoricalFieldResultsAttributes) GetFanoutRespHeaderBytesOk() (*int32, bool)`

GetFanoutRespHeaderBytesOk returns a tuple with the FanoutRespHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFanoutRespHeaderBytes

`func (o *HistoricalFieldResultsAttributes) SetFanoutRespHeaderBytes(v int32)`

SetFanoutRespHeaderBytes sets FanoutRespHeaderBytes field to given value.

### HasFanoutRespHeaderBytes

`func (o *HistoricalFieldResultsAttributes) HasFanoutRespHeaderBytes() bool`

HasFanoutRespHeaderBytes returns a boolean if a field has been set.

### GetFanoutRespBodyBytes

`func (o *HistoricalFieldResultsAttributes) GetFanoutRespBodyBytes() int32`

GetFanoutRespBodyBytes returns the FanoutRespBodyBytes field if non-nil, zero value otherwise.

### GetFanoutRespBodyBytesOk

`func (o *HistoricalFieldResultsAttributes) GetFanoutRespBodyBytesOk() (*int32, bool)`

GetFanoutRespBodyBytesOk returns a tuple with the FanoutRespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFanoutRespBodyBytes

`func (o *HistoricalFieldResultsAttributes) SetFanoutRespBodyBytes(v int32)`

SetFanoutRespBodyBytes sets FanoutRespBodyBytes field to given value.

### HasFanoutRespBodyBytes

`func (o *HistoricalFieldResultsAttributes) HasFanoutRespBodyBytes() bool`

HasFanoutRespBodyBytes returns a boolean if a field has been set.

### GetFanoutBereqHeaderBytes

`func (o *HistoricalFieldResultsAttributes) GetFanoutBereqHeaderBytes() int32`

GetFanoutBereqHeaderBytes returns the FanoutBereqHeaderBytes field if non-nil, zero value otherwise.

### GetFanoutBereqHeaderBytesOk

`func (o *HistoricalFieldResultsAttributes) GetFanoutBereqHeaderBytesOk() (*int32, bool)`

GetFanoutBereqHeaderBytesOk returns a tuple with the FanoutBereqHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFanoutBereqHeaderBytes

`func (o *HistoricalFieldResultsAttributes) SetFanoutBereqHeaderBytes(v int32)`

SetFanoutBereqHeaderBytes sets FanoutBereqHeaderBytes field to given value.

### HasFanoutBereqHeaderBytes

`func (o *HistoricalFieldResultsAttributes) HasFanoutBereqHeaderBytes() bool`

HasFanoutBereqHeaderBytes returns a boolean if a field has been set.

### GetFanoutBereqBodyBytes

`func (o *HistoricalFieldResultsAttributes) GetFanoutBereqBodyBytes() int32`

GetFanoutBereqBodyBytes returns the FanoutBereqBodyBytes field if non-nil, zero value otherwise.

### GetFanoutBereqBodyBytesOk

`func (o *HistoricalFieldResultsAttributes) GetFanoutBereqBodyBytesOk() (*int32, bool)`

GetFanoutBereqBodyBytesOk returns a tuple with the FanoutBereqBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFanoutBereqBodyBytes

`func (o *HistoricalFieldResultsAttributes) SetFanoutBereqBodyBytes(v int32)`

SetFanoutBereqBodyBytes sets FanoutBereqBodyBytes field to given value.

### HasFanoutBereqBodyBytes

`func (o *HistoricalFieldResultsAttributes) HasFanoutBereqBodyBytes() bool`

HasFanoutBereqBodyBytes returns a boolean if a field has been set.

### GetFanoutBerespHeaderBytes

`func (o *HistoricalFieldResultsAttributes) GetFanoutBerespHeaderBytes() int32`

GetFanoutBerespHeaderBytes returns the FanoutBerespHeaderBytes field if non-nil, zero value otherwise.

### GetFanoutBerespHeaderBytesOk

`func (o *HistoricalFieldResultsAttributes) GetFanoutBerespHeaderBytesOk() (*int32, bool)`

GetFanoutBerespHeaderBytesOk returns a tuple with the FanoutBerespHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFanoutBerespHeaderBytes

`func (o *HistoricalFieldResultsAttributes) SetFanoutBerespHeaderBytes(v int32)`

SetFanoutBerespHeaderBytes sets FanoutBerespHeaderBytes field to given value.

### HasFanoutBerespHeaderBytes

`func (o *HistoricalFieldResultsAttributes) HasFanoutBerespHeaderBytes() bool`

HasFanoutBerespHeaderBytes returns a boolean if a field has been set.

### GetFanoutBerespBodyBytes

`func (o *HistoricalFieldResultsAttributes) GetFanoutBerespBodyBytes() int32`

GetFanoutBerespBodyBytes returns the FanoutBerespBodyBytes field if non-nil, zero value otherwise.

### GetFanoutBerespBodyBytesOk

`func (o *HistoricalFieldResultsAttributes) GetFanoutBerespBodyBytesOk() (*int32, bool)`

GetFanoutBerespBodyBytesOk returns a tuple with the FanoutBerespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFanoutBerespBodyBytes

`func (o *HistoricalFieldResultsAttributes) SetFanoutBerespBodyBytes(v int32)`

SetFanoutBerespBodyBytes sets FanoutBerespBodyBytes field to given value.

### HasFanoutBerespBodyBytes

`func (o *HistoricalFieldResultsAttributes) HasFanoutBerespBodyBytes() bool`

HasFanoutBerespBodyBytes returns a boolean if a field has been set.

### GetFanoutConnTimeMs

`func (o *HistoricalFieldResultsAttributes) GetFanoutConnTimeMs() int32`

GetFanoutConnTimeMs returns the FanoutConnTimeMs field if non-nil, zero value otherwise.

### GetFanoutConnTimeMsOk

`func (o *HistoricalFieldResultsAttributes) GetFanoutConnTimeMsOk() (*int32, bool)`

GetFanoutConnTimeMsOk returns a tuple with the FanoutConnTimeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFanoutConnTimeMs

`func (o *HistoricalFieldResultsAttributes) SetFanoutConnTimeMs(v int32)`

SetFanoutConnTimeMs sets FanoutConnTimeMs field to given value.

### HasFanoutConnTimeMs

`func (o *HistoricalFieldResultsAttributes) HasFanoutConnTimeMs() bool`

HasFanoutConnTimeMs returns a boolean if a field has been set.

### GetDdosActionLimitStreamsConnections

`func (o *HistoricalFieldResultsAttributes) GetDdosActionLimitStreamsConnections() int32`

GetDdosActionLimitStreamsConnections returns the DdosActionLimitStreamsConnections field if non-nil, zero value otherwise.

### GetDdosActionLimitStreamsConnectionsOk

`func (o *HistoricalFieldResultsAttributes) GetDdosActionLimitStreamsConnectionsOk() (*int32, bool)`

GetDdosActionLimitStreamsConnectionsOk returns a tuple with the DdosActionLimitStreamsConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdosActionLimitStreamsConnections

`func (o *HistoricalFieldResultsAttributes) SetDdosActionLimitStreamsConnections(v int32)`

SetDdosActionLimitStreamsConnections sets DdosActionLimitStreamsConnections field to given value.

### HasDdosActionLimitStreamsConnections

`func (o *HistoricalFieldResultsAttributes) HasDdosActionLimitStreamsConnections() bool`

HasDdosActionLimitStreamsConnections returns a boolean if a field has been set.

### GetDdosActionLimitStreamsRequests

`func (o *HistoricalFieldResultsAttributes) GetDdosActionLimitStreamsRequests() int32`

GetDdosActionLimitStreamsRequests returns the DdosActionLimitStreamsRequests field if non-nil, zero value otherwise.

### GetDdosActionLimitStreamsRequestsOk

`func (o *HistoricalFieldResultsAttributes) GetDdosActionLimitStreamsRequestsOk() (*int32, bool)`

GetDdosActionLimitStreamsRequestsOk returns a tuple with the DdosActionLimitStreamsRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdosActionLimitStreamsRequests

`func (o *HistoricalFieldResultsAttributes) SetDdosActionLimitStreamsRequests(v int32)`

SetDdosActionLimitStreamsRequests sets DdosActionLimitStreamsRequests field to given value.

### HasDdosActionLimitStreamsRequests

`func (o *HistoricalFieldResultsAttributes) HasDdosActionLimitStreamsRequests() bool`

HasDdosActionLimitStreamsRequests returns a boolean if a field has been set.

### GetDdosActionTarpitAccept

`func (o *HistoricalFieldResultsAttributes) GetDdosActionTarpitAccept() int32`

GetDdosActionTarpitAccept returns the DdosActionTarpitAccept field if non-nil, zero value otherwise.

### GetDdosActionTarpitAcceptOk

`func (o *HistoricalFieldResultsAttributes) GetDdosActionTarpitAcceptOk() (*int32, bool)`

GetDdosActionTarpitAcceptOk returns a tuple with the DdosActionTarpitAccept field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdosActionTarpitAccept

`func (o *HistoricalFieldResultsAttributes) SetDdosActionTarpitAccept(v int32)`

SetDdosActionTarpitAccept sets DdosActionTarpitAccept field to given value.

### HasDdosActionTarpitAccept

`func (o *HistoricalFieldResultsAttributes) HasDdosActionTarpitAccept() bool`

HasDdosActionTarpitAccept returns a boolean if a field has been set.

### GetDdosActionTarpit

`func (o *HistoricalFieldResultsAttributes) GetDdosActionTarpit() int32`

GetDdosActionTarpit returns the DdosActionTarpit field if non-nil, zero value otherwise.

### GetDdosActionTarpitOk

`func (o *HistoricalFieldResultsAttributes) GetDdosActionTarpitOk() (*int32, bool)`

GetDdosActionTarpitOk returns a tuple with the DdosActionTarpit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdosActionTarpit

`func (o *HistoricalFieldResultsAttributes) SetDdosActionTarpit(v int32)`

SetDdosActionTarpit sets DdosActionTarpit field to given value.

### HasDdosActionTarpit

`func (o *HistoricalFieldResultsAttributes) HasDdosActionTarpit() bool`

HasDdosActionTarpit returns a boolean if a field has been set.

### GetDdosActionClose

`func (o *HistoricalFieldResultsAttributes) GetDdosActionClose() int32`

GetDdosActionClose returns the DdosActionClose field if non-nil, zero value otherwise.

### GetDdosActionCloseOk

`func (o *HistoricalFieldResultsAttributes) GetDdosActionCloseOk() (*int32, bool)`

GetDdosActionCloseOk returns a tuple with the DdosActionClose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdosActionClose

`func (o *HistoricalFieldResultsAttributes) SetDdosActionClose(v int32)`

SetDdosActionClose sets DdosActionClose field to given value.

### HasDdosActionClose

`func (o *HistoricalFieldResultsAttributes) HasDdosActionClose() bool`

HasDdosActionClose returns a boolean if a field has been set.

### GetDdosActionBlackhole

`func (o *HistoricalFieldResultsAttributes) GetDdosActionBlackhole() int32`

GetDdosActionBlackhole returns the DdosActionBlackhole field if non-nil, zero value otherwise.

### GetDdosActionBlackholeOk

`func (o *HistoricalFieldResultsAttributes) GetDdosActionBlackholeOk() (*int32, bool)`

GetDdosActionBlackholeOk returns a tuple with the DdosActionBlackhole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdosActionBlackhole

`func (o *HistoricalFieldResultsAttributes) SetDdosActionBlackhole(v int32)`

SetDdosActionBlackhole sets DdosActionBlackhole field to given value.

### HasDdosActionBlackhole

`func (o *HistoricalFieldResultsAttributes) HasDdosActionBlackhole() bool`

HasDdosActionBlackhole returns a boolean if a field has been set.

### GetServiceID

`func (o *HistoricalFieldResultsAttributes) GetServiceID() ReadOnlyIDService`

GetServiceID returns the ServiceID field if non-nil, zero value otherwise.

### GetServiceIDOk

`func (o *HistoricalFieldResultsAttributes) GetServiceIDOk() (*ReadOnlyIDService, bool)`

GetServiceIDOk returns a tuple with the ServiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceID

`func (o *HistoricalFieldResultsAttributes) SetServiceID(v ReadOnlyIDService)`

SetServiceID sets ServiceID field to given value.

### HasServiceID

`func (o *HistoricalFieldResultsAttributes) HasServiceID() bool`

HasServiceID returns a boolean if a field has been set.

### GetStartTime

`func (o *HistoricalFieldResultsAttributes) GetStartTime() int32`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *HistoricalFieldResultsAttributes) GetStartTimeOk() (*int32, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *HistoricalFieldResultsAttributes) SetStartTime(v int32)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *HistoricalFieldResultsAttributes) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
