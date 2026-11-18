# PlatformValues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | Pointer to **time.Time** | Timestamp for the start of the time period being reported | [optional] 
**Datacenter** | Pointer to **string** | Datacenter for this data point. Only present if the request was made with `group_by&#x3D;datacenter`. | [optional] 
**Region** | Pointer to **string** | Region for this data point. Only present if the request was made with `group_by&#x3D;region`. | [optional] 
**Requests** | Pointer to **int64** | Number of requests processed. | [optional] 
**Hits** | Pointer to **int64** | Number of cache hits. | [optional] 
**HitsTime** | Pointer to **float32** | Total amount of time spent processing cache hits (in seconds). | [optional] 
**Miss** | Pointer to **int64** | Number of cache misses. | [optional] 
**MissTime** | Pointer to **float32** | Total amount of time spent processing cache misses (in seconds). | [optional] 
**Pass** | Pointer to **int64** | Number of requests that passed through the CDN without being cached. | [optional] 
**PassTime** | Pointer to **float32** | Total amount of time spent processing cache passes (in seconds). | [optional] 
**Errors** | Pointer to **int64** | Number of cache errors. | [optional] 
**Restarts** | Pointer to **int64** | Number of restarts performed. | [optional] 
**HitRatio** | Pointer to **NullableFloat32** | Ratio of cache hits to cache misses (between 0 and 1). | [optional] 
**Bandwidth** | Pointer to **int64** | Total bytes delivered (`resp_header_bytes` + `resp_body_bytes` + `bereq_header_bytes` + `bereq_body_bytes` + `compute_resp_header_bytes` + `compute_resp_body_bytes` + `compute_bereq_header_bytes` + `compute_bereq_body_bytes` + `websocket_resp_header_bytes` + `websocket_resp_body_bytes` + `websocket_bereq_header_bytes` + `websocket_bereq_body_bytes` + `fanout_resp_header_bytes` + `fanout_resp_body_bytes` + `fanout_bereq_header_bytes` + `fanout_bereq_body_bytes`). | [optional] 
**ReqBodyBytes** | Pointer to **int64** | Total body bytes received. | [optional] 
**ReqHeaderBytes** | Pointer to **int64** | Total header bytes received. | [optional] 
**RespBodyBytes** | Pointer to **int64** | Total body bytes delivered (edge_resp_body_bytes + shield_resp_body_bytes). | [optional] 
**RespHeaderBytes** | Pointer to **int64** | Total header bytes delivered (edge_resp_header_bytes + shield_resp_header_bytes). | [optional] 
**BereqBodyBytes** | Pointer to **int64** | Total body bytes sent to origin. | [optional] 
**BereqHeaderBytes** | Pointer to **int64** | Total header bytes sent to origin. | [optional] 
**Uncacheable** | Pointer to **int64** | Number of requests that were designated uncachable. | [optional] 
**Pipe** | Pointer to **int64** | Optional. Pipe operations performed (legacy feature). | [optional] 
**Synth** | Pointer to **int64** | Number of requests that returned a synthetic response (i.e., response objects created with the `synthetic` VCL statement). | [optional] 
**TLS** | Pointer to **int64** | Number of requests that were received over TLS. | [optional] 
**TLSV10** | Pointer to **int64** | Number of requests received over TLS 1.0. | [optional] 
**TLSV11** | Pointer to **int64** | Number of requests received over TLS 1.1. | [optional] 
**TLSV12** | Pointer to **int64** | Number of requests received over TLS 1.2. | [optional] 
**TLSV13** | Pointer to **int64** | Number of requests received over TLS 1.3. | [optional] 
**EdgeRequests** | Pointer to **int64** | Number of requests sent by end users to Fastly. | [optional] 
**EdgeRespHeaderBytes** | Pointer to **int64** | Total header bytes delivered from Fastly to the end user. | [optional] 
**EdgeRespBodyBytes** | Pointer to **int64** | Total body bytes delivered from Fastly to the end user. | [optional] 
**EdgeHitRequests** | Pointer to **int64** | Number of requests sent by end users to Fastly that resulted in a hit at the edge. | [optional] 
**EdgeMissRequests** | Pointer to **int64** | Number of requests sent by end users to Fastly that resulted in a miss at the edge. | [optional] 
**OriginFetches** | Pointer to **int64** | Number of requests sent to origin. | [optional] 
**OriginFetchHeaderBytes** | Pointer to **int64** | Total request header bytes sent to origin. | [optional] 
**OriginFetchBodyBytes** | Pointer to **int64** | Total request body bytes sent to origin. | [optional] 
**OriginFetchRespHeaderBytes** | Pointer to **int64** | Total header bytes received from origin. | [optional] 
**OriginFetchRespBodyBytes** | Pointer to **int64** | Total body bytes received from origin. | [optional] 
**OriginRevalidations** | Pointer to **int64** | Number of responses received from origin with a `304` status code in response to an `If-Modified-Since` or `If-None-Match` request. Under regular scenarios, a revalidation will imply a cache hit. However, if using Fastly Image Optimizer or segmented caching this may result in a cache miss. | [optional] 
**OriginCacheFetches** | Pointer to **int64** | The total number of completed requests made to backends (origins) that returned cacheable content. | [optional] 
**Shield** | Pointer to **int64** | Number of requests from edge to the shield POP. | [optional] 
**ShieldRespBodyBytes** | Pointer to **int64** | Total body bytes delivered via a shield. | [optional] 
**ShieldRespHeaderBytes** | Pointer to **int64** | Total header bytes delivered via a shield. | [optional] 
**ShieldFetches** | Pointer to **int64** | Number of requests made from one Fastly POP to another, as part of shielding. | [optional] 
**ShieldFetchHeaderBytes** | Pointer to **int64** | Total request header bytes sent to a shield. | [optional] 
**ShieldFetchBodyBytes** | Pointer to **int64** | Total request body bytes sent to a shield. | [optional] 
**ShieldFetchRespHeaderBytes** | Pointer to **int64** | Total response header bytes sent from a shield to the edge. | [optional] 
**ShieldFetchRespBodyBytes** | Pointer to **int64** | Total response body bytes sent from a shield to the edge. | [optional] 
**ShieldRevalidations** | Pointer to **int64** | Number of responses received from origin with a `304` status code, in response to an `If-Modified-Since` or `If-None-Match` request to a shield. Under regular scenarios, a revalidation will imply a cache hit. However, if using segmented caching this may result in a cache miss. | [optional] 
**ShieldCacheFetches** | Pointer to **int64** | The total number of completed requests made to shields that returned cacheable content. | [optional] 
**Ipv6** | Pointer to **int64** | Number of requests that were received over IPv6. | [optional] 
**Otfp** | Pointer to **int64** | Number of responses that came from the Fastly On-the-Fly Packaging service for video-on-demand. | [optional] 
**OtfpRespBodyBytes** | Pointer to **int64** | Total body bytes delivered from the Fastly On-the-Fly Packaging service for video-on-demand. | [optional] 
**OtfpRespHeaderBytes** | Pointer to **int64** | Total header bytes delivered from the Fastly On-the-Fly Packaging service for video-on-demand. | [optional] 
**OtfpShieldRespBodyBytes** | Pointer to **int64** | Total body bytes delivered via a shield for the Fastly On-the-Fly Packaging service for video-on-demand. | [optional] 
**OtfpShieldRespHeaderBytes** | Pointer to **int64** | Total header bytes delivered via a shield for the Fastly On-the-Fly Packaging service for video-on-demand. | [optional] 
**OtfpManifests** | Pointer to **int64** | Number of responses that were manifest files from the Fastly On-the-Fly Packaging service for video-on-demand. | [optional] 
**OtfpDeliverTime** | Pointer to **float32** | Total amount of time spent delivering a response from the Fastly On-the-Fly Packaging service for video-on-demand (in seconds). | [optional] 
**OtfpShieldTime** | Pointer to **float32** | Total amount of time spent delivering a response via a shield from the Fastly On-the-Fly Packaging service for video-on-demand (in seconds). | [optional] 
**Video** | Pointer to **int64** | Number of responses with the video segment or video manifest MIME type (i.e., application/x-mpegurl, application/vnd.apple.mpegurl, application/f4m, application/dash+xml, application/vnd.ms-sstr+xml, ideo/mp2t, audio/aac, video/f4f, video/x-flv, video/mp4, audio/mp4). | [optional] 
**Pci** | Pointer to **int64** | Number of responses with the PCI flag turned on. | [optional] 
**Log** | Pointer to **int64** | Number of log lines sent. | [optional] 
**LogBytes** | Pointer to **int64** | Total log bytes sent. | [optional] 
**HTTP2** | Pointer to **int64** | Number of requests received over HTTP/2. | [optional] 
**HTTP3** | Pointer to **int64** | Number of requests received over HTTP/3. | [optional] 
**AttackReqBodyBytes** | Pointer to **int64** | Total body bytes received from requests that triggered a WAF rule. | [optional] 
**AttackReqHeaderBytes** | Pointer to **int64** | Total header bytes received from requests that triggered a WAF rule. | [optional] 
**AttackLoggedReqBodyBytes** | Pointer to **int64** | Total body bytes received from requests that triggered a WAF rule that was logged. | [optional] 
**AttackLoggedReqHeaderBytes** | Pointer to **int64** | Total header bytes received from requests that triggered a WAF rule that was logged. | [optional] 
**AttackBlockedReqBodyBytes** | Pointer to **int64** | Total body bytes received from requests that triggered a WAF rule that was blocked. | [optional] 
**AttackBlockedReqHeaderBytes** | Pointer to **int64** | Total header bytes received from requests that triggered a WAF rule that was blocked. | [optional] 
**AttackPassedReqBodyBytes** | Pointer to **int64** | Total body bytes received from requests that triggered a WAF rule that was passed. | [optional] 
**AttackPassedReqHeaderBytes** | Pointer to **int64** | Total header bytes received from requests that triggered a WAF rule that was passed. | [optional] 
**AttackRespSynthBytes** | Pointer to **int64** | Total bytes delivered for requests that triggered a WAF rule and returned a synthetic response. | [optional] 
**Imgopto** | Pointer to **int64** | Number of responses that came from the Fastly Image Optimizer service. If the service receives 10 requests for an image, this stat will be 10 regardless of how many times the image was transformed. | [optional] 
**ImgoptoInputImageBytes** | Pointer to **int64** | The total image bytes that would have been delivered if the images had not been transformed by the Fastly Image Optimizer. Bytes are counted once per client request. Byte counts exclude shield traffic and requests where Fastly Image Optimizer proxied a request without attempting to transform it (e.g., file types that are not supported by Image Optimizer). | [optional] 
**ImgoptoOutputImageBytes** | Pointer to **int64** | The total image bytes delivered that have been transformed by the Fastly Image Optimizer. Bytes are counted once per client request. Byte counts exclude shield traffic and requests where Fastly Image Optimizer proxied a request without attempting to transform it (e.g., file types that are not supported by Image Optimizer). | [optional] 
**ImgoptoRespBodyBytes** | Pointer to **int64** | Total body bytes delivered from the Fastly Image Optimizer service, including shield traffic. | [optional] 
**ImgoptoRespHeaderBytes** | Pointer to **int64** | Total header bytes delivered from the Fastly Image Optimizer service, including shield traffic. | [optional] 
**ImgoptoShield** | Pointer to **int64** | Number of responses that came from the Fastly Image Optimizer service via a shield. | [optional] 
**ImgoptoShieldRespBodyBytes** | Pointer to **int64** | Total body bytes delivered via a shield from the Fastly Image Optimizer service. | [optional] 
**ImgoptoShieldRespHeaderBytes** | Pointer to **int64** | Total header bytes delivered via a shield from the Fastly Image Optimizer service. | [optional] 
**ImgoptoTransforms** | Pointer to **int64** | Number of transforms performed by the Fastly Image Optimizer service. | [optional] 
**Imgvideo** | Pointer to **int64** | Number of video responses that came from the Fastly Image Optimizer service. | [optional] 
**ImgvideoFrames** | Pointer to **int64** | Number of video frames that came from the Fastly Image Optimizer service. A video frame is an individual image within a sequence of video. | [optional] 
**ImgvideoRespHeaderBytes** | Pointer to **int64** | Total header bytes of video delivered from the Fastly Image Optimizer service. | [optional] 
**ImgvideoRespBodyBytes** | Pointer to **int64** | Total body bytes of video delivered from the Fastly Image Optimizer service. | [optional] 
**ImgvideoShieldRespHeaderBytes** | Pointer to **int64** | Total header bytes of video delivered via a shield from the Fastly Image Optimizer service. | [optional] 
**ImgvideoShieldRespBodyBytes** | Pointer to **int64** | Total body bytes of video delivered via a shield from the Fastly Image Optimizer service. | [optional] 
**ImgvideoShield** | Pointer to **int64** | Number of video responses delivered via a shield that came from the Fastly Image Optimizer service. | [optional] 
**ImgvideoShieldFrames** | Pointer to **int64** | Number of video frames delivered via a shield that came from the Fastly Image Optimizer service. A video frame is an individual image within a sequence of video. | [optional] 
**Status200** | Pointer to **int64** | Number of responses sent with status code 200 (Success). | [optional] 
**Status204** | Pointer to **int64** | Number of responses sent with status code 204 (No Content). | [optional] 
**Status206** | Pointer to **int64** | Number of responses sent with status code 206 (Partial Content). | [optional] 
**Status301** | Pointer to **int64** | Number of responses sent with status code 301 (Moved Permanently). | [optional] 
**Status302** | Pointer to **int64** | Number of responses sent with status code 302 (Found). | [optional] 
**Status304** | Pointer to **int64** | Number of responses sent with status code 304 (Not Modified). | [optional] 
**Status400** | Pointer to **int64** | Number of responses sent with status code 400 (Bad Request). | [optional] 
**Status401** | Pointer to **int64** | Number of responses sent with status code 401 (Unauthorized). | [optional] 
**Status403** | Pointer to **int64** | Number of responses sent with status code 403 (Forbidden). | [optional] 
**Status404** | Pointer to **int64** | Number of responses sent with status code 404 (Not Found). | [optional] 
**Status406** | Pointer to **int64** | Number of responses sent with status code 406 (Not Acceptable). | [optional] 
**Status416** | Pointer to **int64** | Number of responses sent with status code 416 (Range Not Satisfiable). | [optional] 
**Status429** | Pointer to **int64** | Number of responses sent with status code 429 (Too Many Requests). | [optional] 
**Status500** | Pointer to **int64** | Number of responses sent with status code 500 (Internal Server Error). | [optional] 
**Status501** | Pointer to **int64** | Number of responses sent with status code 501 (Not Implemented). | [optional] 
**Status502** | Pointer to **int64** | Number of responses sent with status code 502 (Bad Gateway). | [optional] 
**Status503** | Pointer to **int64** | Number of responses sent with status code 503 (Service Unavailable). | [optional] 
**Status504** | Pointer to **int64** | Number of responses sent with status code 504 (Gateway Timeout). | [optional] 
**Status505** | Pointer to **int64** | Number of responses sent with status code 505 (HTTP Version Not Supported). | [optional] 
**Status530** | Pointer to **int64** | Number of responses sent with status code 530. | [optional] 
**Status1xx** | Pointer to **int64** | Number of \&quot;Informational\&quot; category status codes delivered. | [optional] 
**Status2xx** | Pointer to **int64** | Number of \&quot;Success\&quot; status codes delivered. | [optional] 
**Status3xx** | Pointer to **int64** | Number of \&quot;Redirection\&quot; codes delivered. | [optional] 
**Status4xx** | Pointer to **int64** | Number of \&quot;Client Error\&quot; codes delivered. | [optional] 
**Status5xx** | Pointer to **int64** | Number of \&quot;Server Error\&quot; codes delivered. | [optional] 
**ObjectSize1k** | Pointer to **int64** | Number of objects served that were under 1KB in size. | [optional] 
**ObjectSize10k** | Pointer to **int64** | Number of objects served that were between 1KB and 10KB in size. | [optional] 
**ObjectSize100k** | Pointer to **int64** | Number of objects served that were between 10KB and 100KB in size. | [optional] 
**ObjectSize1m** | Pointer to **int64** | Number of objects served that were between 100KB and 1MB in size. | [optional] 
**ObjectSize10m** | Pointer to **int64** | Number of objects served that were between 1MB and 10MB in size. | [optional] 
**ObjectSize100m** | Pointer to **int64** | Number of objects served that were between 10MB and 100MB in size. | [optional] 
**ObjectSize1g** | Pointer to **int64** | Number of objects served that were between 100MB and 1GB in size. | [optional] 
**RecvSubTime** | Pointer to **float32** | Time spent inside the `vcl_recv` Varnish subroutine (in seconds). | [optional] 
**RecvSubCount** | Pointer to **int64** | Number of executions of the `vcl_recv` Varnish subroutine. | [optional] 
**HashSubTime** | Pointer to **float32** | Time spent inside the `vcl_hash` Varnish subroutine (in seconds). | [optional] 
**HashSubCount** | Pointer to **int64** | Number of executions of the `vcl_hash` Varnish subroutine. | [optional] 
**MissSubTime** | Pointer to **float32** | Time spent inside the `vcl_miss` Varnish subroutine (in seconds). | [optional] 
**MissSubCount** | Pointer to **int64** | Number of executions of the `vcl_miss` Varnish subroutine. | [optional] 
**FetchSubTime** | Pointer to **float32** | Time spent inside the `vcl_fetch` Varnish subroutine (in seconds). | [optional] 
**FetchSubCount** | Pointer to **int64** | Number of executions of the `vcl_fetch` Varnish subroutine. | [optional] 
**PassSubTime** | Pointer to **float32** | Time spent inside the `vcl_pass` Varnish subroutine (in seconds). | [optional] 
**PassSubCount** | Pointer to **int64** | Number of executions of the `vcl_pass` Varnish subroutine. | [optional] 
**PipeSubTime** | Pointer to **float32** | Time spent inside the `vcl_pipe` Varnish subroutine (in seconds). | [optional] 
**PipeSubCount** | Pointer to **int64** | Number of executions of the `vcl_pipe` Varnish subroutine. | [optional] 
**DeliverSubTime** | Pointer to **float32** | Time spent inside the `vcl_deliver` Varnish subroutine (in seconds). | [optional] 
**DeliverSubCount** | Pointer to **int64** | Number of executions of the `vcl_deliver` Varnish subroutine. | [optional] 
**ErrorSubTime** | Pointer to **float32** | Time spent inside the `vcl_error` Varnish subroutine (in seconds). | [optional] 
**ErrorSubCount** | Pointer to **int64** | Number of executions of the `vcl_error` Varnish subroutine. | [optional] 
**HitSubTime** | Pointer to **float32** | Time spent inside the `vcl_hit` Varnish subroutine (in seconds). | [optional] 
**HitSubCount** | Pointer to **int64** | Number of executions of the `vcl_hit` Varnish subroutine. | [optional] 
**PrehashSubTime** | Pointer to **float32** | Time spent inside the `vcl_prehash` Varnish subroutine (in seconds). | [optional] 
**PrehashSubCount** | Pointer to **int64** | Number of executions of the `vcl_prehash` Varnish subroutine. | [optional] 
**PredeliverSubTime** | Pointer to **float32** | Time spent inside the `vcl_predeliver` Varnish subroutine (in seconds). | [optional] 
**PredeliverSubCount** | Pointer to **int64** | Number of executions of the `vcl_predeliver` Varnish subroutine. | [optional] 
**TLSHandshakeSentBytes** | Pointer to **int64** | Number of bytes transferred during TLS handshake. | [optional] 
**HitRespBodyBytes** | Pointer to **int64** | Total body bytes delivered for cache hits. | [optional] 
**MissRespBodyBytes** | Pointer to **int64** | Total body bytes delivered for cache misses. | [optional] 
**PassRespBodyBytes** | Pointer to **int64** | Total body bytes delivered for cache passes. | [optional] 
**SegblockOriginFetches** | Pointer to **int64** | Number of `Range` requests to origin for segments of resources when using segmented caching. | [optional] 
**SegblockShieldFetches** | Pointer to **int64** | Number of `Range` requests to a shield for segments of resources when using segmented caching. | [optional] 
**ComputeRequests** | Pointer to **int64** | The total number of requests that were received for your service by Fastly. | [optional] 
**ComputeRequestTimeMs** | Pointer to **float32** | The total, actual amount of time used to process your requests, including active CPU time (in milliseconds). | [optional] 
**ComputeRequestTimeBilledMs** | Pointer to **float32** | The total amount of request processing time you will be billed for, measured in 50 millisecond increments. | [optional] 
**ComputeRAMUsed** | Pointer to **int64** | The amount of RAM used for your service by Fastly (in bytes). | [optional] 
**ComputeExecutionTimeMs** | Pointer to **float32** | The amount of active CPU time used to process your requests (in milliseconds). | [optional] 
**ComputeReqHeaderBytes** | Pointer to **int64** | Total header bytes received by the Compute platform. | [optional] 
**ComputeReqBodyBytes** | Pointer to **int64** | Total body bytes received by the Compute platform. | [optional] 
**ComputeRespHeaderBytes** | Pointer to **int64** | Total header bytes sent from Compute to end user. | [optional] 
**ComputeRespBodyBytes** | Pointer to **int64** | Total body bytes sent from Compute to end user. | [optional] 
**ComputeRespStatus1xx** | Pointer to **int64** | Number of \&quot;Informational\&quot; category status codes delivered by the Compute platform. | [optional] 
**ComputeRespStatus2xx** | Pointer to **int64** | Number of \&quot;Success\&quot; category status codes delivered by the Compute platform. | [optional] 
**ComputeRespStatus3xx** | Pointer to **int64** | Number of \&quot;Redirection\&quot; category status codes delivered by the Compute platform. | [optional] 
**ComputeRespStatus4xx** | Pointer to **int64** | Number of \&quot;Client Error\&quot; category status codes delivered by the Compute platform. | [optional] 
**ComputeRespStatus5xx** | Pointer to **int64** | Number of \&quot;Server Error\&quot; category status codes delivered by the Compute platform. | [optional] 
**ComputeBereqHeaderBytes** | Pointer to **int64** | Total header bytes sent to backends (origins) by the Compute platform. | [optional] 
**ComputeBereqBodyBytes** | Pointer to **int64** | Total body bytes sent to backends (origins) by the Compute platform. | [optional] 
**ComputeBerespHeaderBytes** | Pointer to **int64** | Total header bytes received from backends (origins) by the Compute platform. | [optional] 
**ComputeBerespBodyBytes** | Pointer to **int64** | Total body bytes received from backends (origins) by the Compute platform. | [optional] 
**ComputeBereqs** | Pointer to **int64** | Number of backend requests started. | [optional] 
**ComputeServiceBereqError** | Pointer to **int64** | Number of backend request errors, including timeouts. | [optional] 
**ComputeServiceMemoryExceededError** | Pointer to **int64** | Number of times a guest exceeded its heap limit. | [optional] 
**EdgeHitRespBodyBytes** | Pointer to **int64** | Body bytes delivered for edge hits. | [optional] 
**EdgeHitRespHeaderBytes** | Pointer to **int64** | Header bytes delivered for edge hits. | [optional] 
**EdgeMissRespBodyBytes** | Pointer to **int64** | Body bytes delivered for edge misses. | [optional] 
**EdgeMissRespHeaderBytes** | Pointer to **int64** | Header bytes delivered for edge misses. | [optional] 
**OriginCacheFetchRespBodyBytes** | Pointer to **int64** | Body bytes received from origin for cacheable content. | [optional] 
**OriginCacheFetchRespHeaderBytes** | Pointer to **int64** | Header bytes received from an origin for cacheable content. | [optional] 
**ShieldHitRequests** | Pointer to **int64** | Number of requests that resulted in a hit at a shield. | [optional] 
**ShieldMissRequests** | Pointer to **int64** | Number of requests that resulted in a miss at a shield. | [optional] 
**ShieldHitRespHeaderBytes** | Pointer to **int64** | Header bytes delivered for shield hits. | [optional] 
**ShieldHitRespBodyBytes** | Pointer to **int64** | Body bytes delivered for shield hits. | [optional] 
**ShieldMissRespHeaderBytes** | Pointer to **int64** | Header bytes delivered for shield misses. | [optional] 
**ShieldMissRespBodyBytes** | Pointer to **int64** | Body bytes delivered for shield misses. | [optional] 
**WebsocketReqHeaderBytes** | Pointer to **int64** | Total header bytes received from end users over passthrough WebSocket connections. | [optional] 
**WebsocketReqBodyBytes** | Pointer to **int64** | Total message content bytes received from end users over passthrough WebSocket connections. | [optional] 
**WebsocketRespHeaderBytes** | Pointer to **int64** | Total header bytes sent to end users over passthrough WebSocket connections. | [optional] 
**WebsocketRespBodyBytes** | Pointer to **int64** | Total message content bytes sent to end users over passthrough WebSocket connections. | [optional] 
**WebsocketBereqHeaderBytes** | Pointer to **int64** | Total header bytes sent to backends over passthrough WebSocket connections. | [optional] 
**WebsocketBereqBodyBytes** | Pointer to **int64** | Total message content bytes sent to backends over passthrough WebSocket connections. | [optional] 
**WebsocketBerespHeaderBytes** | Pointer to **int64** | Total header bytes received from backends over passthrough WebSocket connections. | [optional] 
**WebsocketBerespBodyBytes** | Pointer to **int64** | Total message content bytes received from backends over passthrough WebSocket connections. | [optional] 
**WebsocketConnTimeMs** | Pointer to **int64** | Total duration of passthrough WebSocket connections with end users. | [optional] 
**FanoutRecvPublishes** | Pointer to **int64** | Total published messages received from the publish API endpoint. | [optional] 
**FanoutSendPublishes** | Pointer to **int64** | Total published messages sent to end users. | [optional] 
**KvStoreClassAOperations** | Pointer to **int64** | The total number of class a operations for the KV store. | [optional] 
**KvStoreClassBOperations** | Pointer to **int64** | The total number of class b operations for the KV store. | [optional] 
**FanoutReqHeaderBytes** | Pointer to **int64** | Total header bytes received from end users over Fanout connections. | [optional] 
**FanoutReqBodyBytes** | Pointer to **int64** | Total body or message content bytes received from end users over Fanout connections. | [optional] 
**FanoutRespHeaderBytes** | Pointer to **int64** | Total header bytes sent to end users over Fanout connections. | [optional] 
**FanoutRespBodyBytes** | Pointer to **int64** | Total body or message content bytes sent to end users over Fanout connections, excluding published message content. | [optional] 
**FanoutBereqHeaderBytes** | Pointer to **int64** | Total header bytes sent to backends over Fanout connections. | [optional] 
**FanoutBereqBodyBytes** | Pointer to **int64** | Total body or message content bytes sent to backends over Fanout connections. | [optional] 
**FanoutBerespHeaderBytes** | Pointer to **int64** | Total header bytes received from backends over Fanout connections. | [optional] 
**FanoutBerespBodyBytes** | Pointer to **int64** | Total body or message content bytes received from backends over Fanout connections. | [optional] 
**FanoutConnTimeMs** | Pointer to **int64** | Total duration of Fanout connections with end users. | [optional] 
**DdosActionLimitStreamsConnections** | Pointer to **int64** | For HTTP/2, the number of connections the limit-streams action was applied to. The limit-streams action caps the allowed number of concurrent streams in a connection. | [optional] 
**DdosActionLimitStreamsRequests** | Pointer to **int64** | For HTTP/2, the number of requests made on a connection for which the limit-streams action was taken. The limit-streams action caps the allowed number of concurrent streams in a connection. | [optional] 
**DdosActionTarpitAccept** | Pointer to **int64** | The number of times the tarpit-accept action was taken. The tarpit-accept action adds a delay when accepting future connections. | [optional] 
**DdosActionTarpit** | Pointer to **int64** | The number of times the tarpit action was taken. The tarpit action delays writing the response to the client. | [optional] 
**DdosActionClose** | Pointer to **int64** | The number of times the close action was taken. The close action aborts the connection as soon as possible. The close action takes effect either right after accept, right after the client hello, or right after the response was sent. | [optional] 
**DdosActionBlackhole** | Pointer to **int64** | The number of times the blackhole action was taken. The blackhole action quietly closes a TCP connection without sending a reset. The blackhole action quietly closes a TCP connection without notifying its peer (all TCP state is dropped). | [optional] 
**BotChallengeStarts** | Pointer to **int64** | The number of challenge-start tokens created. | [optional] 
**BotChallengeCompleteTokensPassed** | Pointer to **int64** | The number of challenge-complete tokens that passed validation. | [optional] 
**BotChallengeCompleteTokensFailed** | Pointer to **int64** | The number of challenge-complete tokens that failed validation. | [optional] 
**BotChallengeCompleteTokensChecked** | Pointer to **int64** | The number of challenge-complete tokens checked. | [optional] 
**BotChallengeCompleteTokensDisabled** | Pointer to **int64** | The number of challenge-complete tokens not checked because the feature was disabled. | [optional] 
**BotChallengeCompleteTokensIssued** | Pointer to **int64** | The number of challenge-complete tokens issued. For example, issuing a challenge-complete token after a series of CAPTCHA challenges ending in success. | [optional] 
**BotChallengesIssued** | Pointer to **int64** | The number of challenges issued. For example, the issuance of a CAPTCHA challenge. | [optional] 
**BotChallengesSucceeded** | Pointer to **int64** | The number of successful challenge solutions processed. For example, a correct CAPTCHA solution. | [optional] 
**BotChallengesFailed** | Pointer to **int64** | The number of failed challenge solutions processed. For example, an incorrect CAPTCHA solution. | [optional] 
**DdosActionDowngrade** | Pointer to **int64** | The number of times the downgrade action was taken. The downgrade action restricts the client to http1. | [optional] 
**DdosActionDowngradedConnections** | Pointer to **int64** | The number of connections the downgrade action was applied to. The downgrade action restricts the connection to http1. | [optional] 
**AllHitRequests** | Pointer to **int64** | Number of cache hits for a VCL service. | [optional] 
**AllMissRequests** | Pointer to **int64** | Number of cache misses for a VCL service. | [optional] 
**AllPassRequests** | Pointer to **int64** | Number of requests that passed through the CDN without being cached for a VCL service. | [optional] 
**AllErrorRequests** | Pointer to **int64** | Number of cache errors for a VCL service. | [optional] 
**AllSynthRequests** | Pointer to **int64** | Number of requests that returned a synthetic response (i.e., response objects created with the `synthetic` VCL statement) for a VCL service. | [optional] 
**AllEdgeHitRequests** | Pointer to **int64** | Number of requests sent by end users to Fastly that resulted in a hit at the edge for a VCL service. | [optional] 
**AllEdgeMissRequests** | Pointer to **int64** | Number of requests sent by end users to Fastly that resulted in a miss at the edge for a VCL service. | [optional] 
**AllStatus1xx** | Pointer to **int64** | Number of \&quot;Informational\&quot; category status codes delivered for all sources. | [optional] 
**AllStatus2xx** | Pointer to **int64** | Number of \&quot;Success\&quot; status codes delivered for all sources. | [optional] 
**AllStatus3xx** | Pointer to **int64** | Number of \&quot;Redirection\&quot; codes delivered for all sources. | [optional] 
**AllStatus4xx** | Pointer to **int64** | Number of \&quot;Client Error\&quot; codes delivered for all sources. | [optional] 
**AllStatus5xx** | Pointer to **int64** | Number of \&quot;Server Error\&quot; codes delivered for all sources. | [optional] 
**OriginOffload** | Pointer to **float32** | Origin Offload measures the ratio of bytes served to end users that were cached by Fastly, over the bytes served to end users, between 0 and 1. ((`edge_resp_body_bytes` + `edge_resp_header_bytes`) - (`origin_fetch_resp_body_bytes` + `origin_fetch_resp_header_bytes`)) / (`edge_resp_body_bytes` + `edge_resp_header_bytes`). | [optional] 
**RequestDeniedGetHeadBody** | Pointer to **int64** | Number of requests where Fastly responded with 400 due to the request being a GET or HEAD request containing a body. | [optional] 
**DdosProtectionRequestsDetectCount** | Pointer to **int64** | Number of requests classified as a DDoS attack against a customer origin or service. | [optional] 
**DdosProtectionRequestsMitigateCount** | Pointer to **int64** | Number of requests classified as a DDoS attack against a customer origin or service that were mitigated by the Fastly platform. | [optional] 
**DdosProtectionRequestsAllowCount** | Pointer to **int64** | Number of requests analyzed for DDoS attacks against a customer origin or service, but with no DDoS detected. | [optional] 
**ObjectStorageClassAOperationsCount** | Pointer to **int64** | A count of the number of Class A Object Storage operations. | [optional] 
**ObjectStorageClassBOperationsCount** | Pointer to **int64** | A count of the number of Class B Object Storage operations. | [optional] 
**AiaRequests** | Pointer to **int64** | Number of requests received by AI Accelerator. | [optional] 
**AiaStatus1xx** | Pointer to **int64** | Number of \&quot;Informational\&quot; category status codes received from AI provider. | [optional] 
**AiaStatus2xx** | Pointer to **int64** | Number of \&quot;Success\&quot; status codes received from AI provider. | [optional] 
**AiaStatus3xx** | Pointer to **int64** | Number of \&quot;Redirection\&quot; received from AI provider. | [optional] 
**AiaStatus4xx** | Pointer to **int64** | Number of \&quot;Client Error\&quot; received from AI provider. | [optional] 
**AiaStatus5xx** | Pointer to **int64** | Number of \&quot;Server Error\&quot; received from AI provider. | [optional] 
**AiaResponseUsageTokens** | Pointer to **int64** | The usage tokens associated with the response returned from the AI Accelerator cache. | [optional] 
**AiaOriginUsageTokens** | Pointer to **int64** | The number of usage tokens reported by the request to the origin from AI Accelerator. | [optional] 
**AiaEstimatedTimeSavedMs** | Pointer to **int64** | The estimated amount of time saved by responses served from the AI Accelerator semantic cache. | [optional] 
**RequestCollapseUsableCount** | Pointer to **int64** | Number of requests that were collapsed and satisfied by a usable cache object. | [optional] 
**RequestCollapseUnusableCount** | Pointer to **int64** | Number of requests that were collapsed and unable to be satisfied by the resulting cache object. | [optional] 
**ComputeCacheOperationsCount** | Pointer to **int64** | Number of cache operations executed by the Compute platform. | [optional] 
**NgwafRequestsTotalCount** | Pointer to **int32** | Total number of Next-Gen WAF (Edge WAF &amp; Core WAF) requests. | [optional] 
**NgwafRequestsUnknownCount** | Pointer to **int32** | Count of Edge WAF requests with an unknown outcome. | [optional] 
**NgwafRequestsAllowedCount** | Pointer to **int32** | Count of Next-Gen WAF (Edge WAF &amp; Core WAF) requests allowed. | [optional] 
**NgwafRequestsLoggedCount** | Pointer to **int32** | Count of Edge WAF requests logged. | [optional] 
**NgwafRequestsBlockedCount** | Pointer to **int32** | Count of Next-Gen WAF (Edge WAF &amp; Core WAF) requests blocked. | [optional] 
**NgwafRequestsTimeoutCount** | Pointer to **int32** | Count of Edge WAF requests timed outcome. | [optional] 
**NgwafRequestsChallengedCount** | Pointer to **int32** | Count of Edge WAF requests challenged. | [optional] 
**ApiDiscoveryRequestsCount** | Pointer to **int32** | Number of requests processed by the API Discovery engine. | [optional] 
**ComputeRespStatus103** | Pointer to **int32** | Number of responses delivered with status code 103 (Early Hints) by the Compute platform. | [optional] 
**ComputeRespStatus200** | Pointer to **int32** | Number of responses delivered with status code 200 (Success) by the Compute platform. | [optional] 
**ComputeRespStatus204** | Pointer to **int32** | Number of responses delivered with status code 204 (No Content) by the Compute platform. | [optional] 
**ComputeRespStatus206** | Pointer to **int32** | Number of responses delivered with status code 206 (Partial Content) by the Compute platform. | [optional] 
**ComputeRespStatus301** | Pointer to **int32** | Number of responses delivered with status code 301 (Moved Permanently) by the Compute platform. | [optional] 
**ComputeRespStatus302** | Pointer to **int32** | Number of responses delivered with status code 302 (Found) by the Compute platform. | [optional] 
**ComputeRespStatus304** | Pointer to **int32** | Number of responses delivered with status code 304 (Not Modified) by the Compute platform. | [optional] 
**ComputeRespStatus400** | Pointer to **int32** | Number of responses delivered with status code 400 (Bad Request) by the Compute platform. | [optional] 
**ComputeRespStatus401** | Pointer to **int32** | Number of responses delivered with status code 401 (Unauthorized) by the Compute platform. | [optional] 
**ComputeRespStatus403** | Pointer to **int32** | Number of responses delivered with status code 403 (Forbidden) by the Compute platform. | [optional] 
**ComputeRespStatus404** | Pointer to **int32** | Number of responses delivered with status code 404 (Not Found) by the Compute platform. | [optional] 
**ComputeRespStatus416** | Pointer to **int32** | Number of responses delivered with status code 416 (Range Not Satisfiable) by the Compute platform. | [optional] 
**ComputeRespStatus429** | Pointer to **int32** | Number of responses delivered with status code 429 (Too Many Requests) by the Compute platform. | [optional] 
**ComputeRespStatus500** | Pointer to **int32** | Number of responses delivered with status code 500 (Internal Server Error) by the Compute platform. | [optional] 
**ComputeRespStatus501** | Pointer to **int32** | Number of responses delivered with status code 501 (Not Implemented) by the Compute platform. | [optional] 
**ComputeRespStatus502** | Pointer to **int32** | Number of responses delivered with status code 502 (Bad Gateway) by the Compute platform. | [optional] 
**ComputeRespStatus503** | Pointer to **int32** | Number of responses delivered with status code 503 (Service Unavailable) by the Compute platform. | [optional] 
**ComputeRespStatus504** | Pointer to **int32** | Number of responses delivered with status code 504 (Gateway Timeout) by the Compute platform. | [optional] 
**ComputeRespStatus505** | Pointer to **int32** | Number of responses delivered with status code 505 (HTTP Version Not Supported) by the Compute platform. | [optional] 
**ComputeRespStatus530** | Pointer to **int32** | Number of responses delivered with status code 530 by the Compute platform. | [optional] 
**ImgoptoComputeRequests** | Pointer to **int32** | The number of Image Optimizer requests made from Compute services. | [optional] 
**DNSBillableResponsesCount** | Pointer to **int32** | Number of billable DNS responses (e.g., A, CNAME). | [optional] 
**DNSNonbillableResponsesCount** | Pointer to **int32** | Number of non-billable DNS responses (e.g., NODATA, NXDOMAIN). | [optional] 
**Upgrade** | Pointer to **int32** | Number of requests that resulted in a WebSocket upgrade. | [optional] 
**NgwafBotAnalysisRequestCount** | Pointer to **int32** | Count of Next-Gen WAF Bot Management requests. | [optional] 
**ImgoptoAvifCount** | Pointer to **int32** | Count of AVIF images delivered to end user by Image Optimizer. | [optional] 
**ImgoptoJpegCount** | Pointer to **int32** | Count of JPEG images delivered to end user by Image Optimizer. | [optional] 
**ImgoptoPngCount** | Pointer to **int32** | Count of PNG images delivered to end user by Image Optimizer. | [optional] 
**ImgoptoGifCount** | Pointer to **int32** | Count of GIF images delivered to end user by Image Optimizer. | [optional] 
**ImgoptoWebpCount** | Pointer to **int32** | Count of WebP images delivered to end user by Image Optimizer. | [optional] 
**ImgoptoJpegxlCount** | Pointer to **int32** | Count of JPEGXL images delivered to end user by Image Optimizer. | [optional] 
**ImgoptoSvgCount** | Pointer to **int32** | Count of SVG images delivered to end user by Image Optimizer. | [optional] 
**ImgoptoMp4Count** | Pointer to **int32** | Count of MP4s delivered to end user by Image Optimizer. | [optional] 
**ComputeServiceResourceLimitsError** | Pointer to **int32** | Aggregate of fatal errors caused by exceeding allocated resource limits, specifically runtime duration, vCPU usage, and heap memory limits. | [optional] 
**ComputeServiceRuntimeError** | Pointer to **int32** | Fatal errors caused by service logic faults, including stack overflows, unreachable code traps, illegal memory access, or attempts to send multiple responses. | [optional] 
**ComputeServiceChainError** | Pointer to **int32** | Fatal errors caused by the service path exceeding hop or service limits, or where a forwarding loop is detected via CDN-Loop headers. | [optional] 
**ComputePlatformInternalError** | Pointer to **int32** | Fatal errors caused by internal errors in Fastly’s Compute platform. | [optional] 
**ComputeServiceTimeoutError** | Pointer to **int32** | Fatal errors caused by exceeding the per-request runtime limit. | [optional] 
**ComputeServiceVcpuExceededError** | Pointer to **int32** | Fatal errors caused by exceeding the per-request vCPU time limit. | [optional] 
**ComputeServiceLimitsError** | Pointer to **int32** | Non-fatal errors caused by attempts to exceed defined operational limits, such as simultaneous backend requests or cache transactions. | [optional] 
**ComputePlatformInvalidRequestError** | Pointer to **int32** | Fatal errors caused by unprocessable requests to the service, such as requests with malformed CDN-Loop headers or invalid purge credentials. | [optional] 
**ComputeSandboxes** | Pointer to **int32** | Number of WebAssembly (Wasm) sandboxes created. | [optional] 
**BotRequestsTotalCount** | Pointer to **int32** | Total number of Bot Management requests across all deployments. | [optional] 
**BotEdgeRequestsAnalyzedCount** | Pointer to **int32** | Count of edge requests where bot detection analysis was performed. | [optional] 
**BotEdgeRequestsDetectedCount** | Pointer to **int32** | Count of edge requests where a bot was detected. | [optional] 
**BotEdgeRequestsVerifiedCount** | Pointer to **int32** | Count of edge requests where a verified bot was detected. | [optional] 
**BotEdgeRequestsAiCrawlerCount** | Pointer to **int32** | Count of edge requests where an AI crawler was detected. | [optional] 
**BotEdgeRequestsAiFetcherCount** | Pointer to **int32** | Count of edge requests where an AI fetcher was detected. | [optional] 
**BotEdgeRequestsAccessibilityCount** | Pointer to **int32** | Count of edge requests where an accessibility bot was detected. | [optional] 
**BotEdgeRequestsContentFetcherCount** | Pointer to **int32** | Count of edge requests where a content fetcher was detected. | [optional] 
**BotEdgeRequestsMonitoringCount** | Pointer to **int32** | Count of edge requests where a monitoring and site tool was detected. | [optional] 
**BotEdgeRequestsOnlineMarketingCount** | Pointer to **int32** | Count of edge requests where an online marketing bot was detected. | [optional] 
**BotEdgeRequestsPagePreviewCount** | Pointer to **int32** | Count of edge requests where a page preview bot was detected. | [optional] 
**BotEdgeRequestsPlatformIntegrationsCount** | Pointer to **int32** | Count of edge requests where a platform integration was detected. | [optional] 
**BotEdgeRequestsResearchCount** | Pointer to **int32** | Count of edge requests where a research bot was detected. | [optional] 
**BotEdgeRequestsSearchEngineCrawlerCount** | Pointer to **int32** | Count of edge requests where a search engine crawler was detected. | [optional] 
**BotEdgeRequestsSearchEngineOptimizationCount** | Pointer to **int32** | Count of edge requests where a search engine optimization bot was detected. | [optional] 
**BotEdgeRequestsSecurityToolsCount** | Pointer to **int32** | Count of edge requests where a security tool was detected. | [optional] 
**ComputeHandoff** | Pointer to **int32** | The number of times Compute has handed off a request to the Fanout proxy or WebSocket proxy. | [optional] 
**ComputeServiceBereqDNSError** | Pointer to **int32** | Number of backend requests from a Compute service that failed during DNS resolution. | [optional] 
**ComputeServiceBereqConnTimeoutError** | Pointer to **int32** | Number of backend requests from a Compute service where the connection to the origin timed out before being established. | [optional] 
**ComputeServiceBereqConnRefusedError** | Pointer to **int32** | Number of backend requests from a Compute service where the origin actively refused the connection. | [optional] 
**ComputeServiceBereqConnOtherError** | Pointer to **int32** | Number of backend requests from a Compute service that failed due to a connection error not classified as a timeout or refusal. | [optional] 
**ComputeServiceBereqTLSServerCertError** | Pointer to **int32** | Number of backend requests from a Compute service that failed due to a TLS certificate validation error (e.g., expired, untrusted CA, hostname mismatch). | [optional] 
**ComputeServiceBereqTLSOtherError** | Pointer to **int32** | Number of backend requests from a Compute service that failed due to a TLS error not classified as a certificate error. | [optional] 
**ComputeServiceBereqHTTPProtoV1Error** | Pointer to **int32** | Number of backend requests from a Compute service that failed due to an HTTP/1.x protocol violation after the request was transmitted. | [optional] 
**ComputeServiceBereqHTTPProtoV2Error** | Pointer to **int32** | Number of backend requests from a Compute service that failed due to an HTTP/2 protocol error, typically a `RST_STREAM` or `GO_AWAY` from the origin. | [optional] 
**ComputeServiceBereqHTTPIncompleteError** | Pointer to **int32** | Number of backend requests from a Compute service where the origin sent an incomplete HTTP response. | [optional] 
**ComputeServiceBereqHTTPTimeoutError** | Pointer to **int32** | Number of backend requests from a Compute service where the origin did not respond within the configured timeout period. | [optional] 
**ComputeServiceBereqHTTPOtherError** | Pointer to **int32** | Number of backend requests from a Compute service that failed due to an HTTP-level error not classified in any category. | [optional] 
**ComputeServiceBereqOtherError** | Pointer to **int32** | Number of backend requests from a Compute service that failed due to an error not classified into the DNS, connection, TLS, or HTTP categories. | [optional] 
**ComputeServiceBereq5xxError** | Pointer to **int32** | Number of backend requests from a Compute service where the origin returned a 5xx status code. | [optional] 
**ComputeServiceBereqConnError** | Pointer to **int32** | Number of backend requests from a Compute service that failed at the TCP connection level. Sum of `compute_service_bereq_conn_timeout_error`, `compute_service_bereq_conn_refused_error`, and `compute_service_bereq_conn_other_error`. | [optional] 
**ComputeServiceBereqTLSError** | Pointer to **int32** | Number of backend requests from a Compute service that failed during the TLS handshake or session with the origin. Sum of `compute_service_bereq_tls_server_cert_error` and `compute_service_bereq_tls_other_error`. | [optional] 
**ComputeServiceBereqHTTPError** | Pointer to **int32** | Number of backend requests from a Compute service that failed at the HTTP protocol level. Sum of `compute_service_bereq_http_proto_v1_error`, `compute_service_bereq_http_proto_v2_error`, `compute_service_bereq_http_incomplete_error`, `compute_service_bereq_http_timeout_error`, and `compute_service_bereq_http_other_error`. | [optional] 
**BotChallengesPatsIssued** | Pointer to **int32** | Number of Private Access Token challenges issued. | [optional] 
**BotChallengesPatsSucceeded** | Pointer to **int32** | Number of successful Private Access Token challenge solutions processed. | [optional] 
**BotEdgeRequestsHeadlessCount** | Pointer to **int32** | Number of edge requests where a headless bot was detected. | [optional] 
**Status499** | Pointer to **int32** | Number of responses sent with status code 499 (Client Disconnected). | [optional] 
**ArcRequests** | Pointer to **int32** | Number of requests received by AI Runtime Control. | [optional] 
**DdosChallengeCompleteTokensChecked** | Pointer to **int32** | The number of challenge-complete tokens checked while enforcing DDoS protection rules. | [optional] 
**DdosChallengeCompleteTokensPassed** | Pointer to **int32** | The number of challenge-complete tokens that passed validation while enforcing DDoS protection rules. | [optional] 
**DdosChallengeCompleteTokensFailed** | Pointer to **int32** | The number of challenge-complete tokens that failed validation while enforcing DDoS protection rules. | [optional] 
**DdosChallengeStarts** | Pointer to **int32** | The number of challenge-start tokens created while enforcing DDoS protection rules. | [optional] 
**DdosChallengeCompleteTokensIssued** | Pointer to **int32** | Number of challenge-complete tokens issued as a result of a DDoS action. For example, issuing a challenge-complete token after a series of CAPTCHA challenges ending in success. | [optional] 
**DdosChallengesIssued** | Pointer to **int32** | Number of challenges issued from a DDoS action. | [optional] 
**DdosChallengesSucceeded** | Pointer to **int32** | Number of successful challenge solutions processed as a result of a DDoS action. For example, a correct CAPTCHA solution. | [optional] 
**DdosChallengesFailed** | Pointer to **int32** | Number of failed challenge solutions processed as a result of a DDoS action. For example, an incorrect CAPTCHA solution. | [optional] 
**DdosChallengesPatsIssued** | Pointer to **int32** | Number of Private Access Token challenges issued as a result of a DDoS action. | [optional] 
**DdosChallengesPatsSucceeded** | Pointer to **int32** | Number of successful Private Access Token challenge solutions processed as a result of a DDoS action. | [optional] 
**CspPagesGauge** | Pointer to **int64** | **Gauge.** The number of pages monitored by Content Security Policy. | [optional] 
**NgwafAccountRulesGauge** | Pointer to **int64** | **Gauge.** The number of Next-Gen WAF rules defined at the account level. | [optional] 
**NgwafAccountWorkspaces** | Pointer to **int64** | **Gauge.** The number of Next-Gen WAF workspaces in the account. | [optional] 
**NgwafWorkspaceRulesGauge** | Pointer to **int64** | **Gauge.** The number of Next-Gen WAF rules defined in a workspace. | [optional] 
**NgwafWorkspaceThresholdsGauge** | Pointer to **int64** | **Gauge.** The number of Next-Gen WAF thresholds defined in a workspace. | [optional] 
**ObjectStorageBucketsCount** | Pointer to **int64** | **Gauge.** The number of Object Storage buckets. | [optional] 
**ObjectStorageBytesCount** | Pointer to **int64** | **Gauge.** The number of bytes stored in Object Storage. | [optional] 

## Methods

### NewPlatformValues

`func NewPlatformValues() *PlatformValues`

NewPlatformValues instantiates a new PlatformValues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformValuesWithDefaults

`func NewPlatformValuesWithDefaults() *PlatformValues`

NewPlatformValuesWithDefaults instantiates a new PlatformValues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *PlatformValues) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *PlatformValues) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *PlatformValues) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *PlatformValues) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetDatacenter

`func (o *PlatformValues) GetDatacenter() string`

GetDatacenter returns the Datacenter field if non-nil, zero value otherwise.

### GetDatacenterOk

`func (o *PlatformValues) GetDatacenterOk() (*string, bool)`

GetDatacenterOk returns a tuple with the Datacenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenter

`func (o *PlatformValues) SetDatacenter(v string)`

SetDatacenter sets Datacenter field to given value.

### HasDatacenter

`func (o *PlatformValues) HasDatacenter() bool`

HasDatacenter returns a boolean if a field has been set.

### GetRegion

`func (o *PlatformValues) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *PlatformValues) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *PlatformValues) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *PlatformValues) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetRequests

`func (o *PlatformValues) GetRequests() int64`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *PlatformValues) GetRequestsOk() (*int64, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *PlatformValues) SetRequests(v int64)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *PlatformValues) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### GetHits

`func (o *PlatformValues) GetHits() int64`

GetHits returns the Hits field if non-nil, zero value otherwise.

### GetHitsOk

`func (o *PlatformValues) GetHitsOk() (*int64, bool)`

GetHitsOk returns a tuple with the Hits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHits

`func (o *PlatformValues) SetHits(v int64)`

SetHits sets Hits field to given value.

### HasHits

`func (o *PlatformValues) HasHits() bool`

HasHits returns a boolean if a field has been set.

### GetHitsTime

`func (o *PlatformValues) GetHitsTime() float32`

GetHitsTime returns the HitsTime field if non-nil, zero value otherwise.

### GetHitsTimeOk

`func (o *PlatformValues) GetHitsTimeOk() (*float32, bool)`

GetHitsTimeOk returns a tuple with the HitsTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHitsTime

`func (o *PlatformValues) SetHitsTime(v float32)`

SetHitsTime sets HitsTime field to given value.

### HasHitsTime

`func (o *PlatformValues) HasHitsTime() bool`

HasHitsTime returns a boolean if a field has been set.

### GetMiss

`func (o *PlatformValues) GetMiss() int64`

GetMiss returns the Miss field if non-nil, zero value otherwise.

### GetMissOk

`func (o *PlatformValues) GetMissOk() (*int64, bool)`

GetMissOk returns a tuple with the Miss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiss

`func (o *PlatformValues) SetMiss(v int64)`

SetMiss sets Miss field to given value.

### HasMiss

`func (o *PlatformValues) HasMiss() bool`

HasMiss returns a boolean if a field has been set.

### GetMissTime

`func (o *PlatformValues) GetMissTime() float32`

GetMissTime returns the MissTime field if non-nil, zero value otherwise.

### GetMissTimeOk

`func (o *PlatformValues) GetMissTimeOk() (*float32, bool)`

GetMissTimeOk returns a tuple with the MissTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissTime

`func (o *PlatformValues) SetMissTime(v float32)`

SetMissTime sets MissTime field to given value.

### HasMissTime

`func (o *PlatformValues) HasMissTime() bool`

HasMissTime returns a boolean if a field has been set.

### GetPass

`func (o *PlatformValues) GetPass() int64`

GetPass returns the Pass field if non-nil, zero value otherwise.

### GetPassOk

`func (o *PlatformValues) GetPassOk() (*int64, bool)`

GetPassOk returns a tuple with the Pass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPass

`func (o *PlatformValues) SetPass(v int64)`

SetPass sets Pass field to given value.

### HasPass

`func (o *PlatformValues) HasPass() bool`

HasPass returns a boolean if a field has been set.

### GetPassTime

`func (o *PlatformValues) GetPassTime() float32`

GetPassTime returns the PassTime field if non-nil, zero value otherwise.

### GetPassTimeOk

`func (o *PlatformValues) GetPassTimeOk() (*float32, bool)`

GetPassTimeOk returns a tuple with the PassTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassTime

`func (o *PlatformValues) SetPassTime(v float32)`

SetPassTime sets PassTime field to given value.

### HasPassTime

`func (o *PlatformValues) HasPassTime() bool`

HasPassTime returns a boolean if a field has been set.

### GetErrors

`func (o *PlatformValues) GetErrors() int64`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *PlatformValues) GetErrorsOk() (*int64, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *PlatformValues) SetErrors(v int64)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *PlatformValues) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetRestarts

`func (o *PlatformValues) GetRestarts() int64`

GetRestarts returns the Restarts field if non-nil, zero value otherwise.

### GetRestartsOk

`func (o *PlatformValues) GetRestartsOk() (*int64, bool)`

GetRestartsOk returns a tuple with the Restarts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestarts

`func (o *PlatformValues) SetRestarts(v int64)`

SetRestarts sets Restarts field to given value.

### HasRestarts

`func (o *PlatformValues) HasRestarts() bool`

HasRestarts returns a boolean if a field has been set.

### GetHitRatio

`func (o *PlatformValues) GetHitRatio() float32`

GetHitRatio returns the HitRatio field if non-nil, zero value otherwise.

### GetHitRatioOk

`func (o *PlatformValues) GetHitRatioOk() (*float32, bool)`

GetHitRatioOk returns a tuple with the HitRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHitRatio

`func (o *PlatformValues) SetHitRatio(v float32)`

SetHitRatio sets HitRatio field to given value.

### HasHitRatio

`func (o *PlatformValues) HasHitRatio() bool`

HasHitRatio returns a boolean if a field has been set.

### SetHitRatioNil

`func (o *PlatformValues) SetHitRatioNil(b bool)`

 SetHitRatioNil sets the value for HitRatio to be an explicit nil

### UnsetHitRatio
`func (o *PlatformValues) UnsetHitRatio()`

UnsetHitRatio ensures that no value is present for HitRatio, not even an explicit nil
### GetBandwidth

`func (o *PlatformValues) GetBandwidth() int64`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *PlatformValues) GetBandwidthOk() (*int64, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *PlatformValues) SetBandwidth(v int64)`

SetBandwidth sets Bandwidth field to given value.

### HasBandwidth

`func (o *PlatformValues) HasBandwidth() bool`

HasBandwidth returns a boolean if a field has been set.

### GetReqBodyBytes

`func (o *PlatformValues) GetReqBodyBytes() int64`

GetReqBodyBytes returns the ReqBodyBytes field if non-nil, zero value otherwise.

### GetReqBodyBytesOk

`func (o *PlatformValues) GetReqBodyBytesOk() (*int64, bool)`

GetReqBodyBytesOk returns a tuple with the ReqBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqBodyBytes

`func (o *PlatformValues) SetReqBodyBytes(v int64)`

SetReqBodyBytes sets ReqBodyBytes field to given value.

### HasReqBodyBytes

`func (o *PlatformValues) HasReqBodyBytes() bool`

HasReqBodyBytes returns a boolean if a field has been set.

### GetReqHeaderBytes

`func (o *PlatformValues) GetReqHeaderBytes() int64`

GetReqHeaderBytes returns the ReqHeaderBytes field if non-nil, zero value otherwise.

### GetReqHeaderBytesOk

`func (o *PlatformValues) GetReqHeaderBytesOk() (*int64, bool)`

GetReqHeaderBytesOk returns a tuple with the ReqHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqHeaderBytes

`func (o *PlatformValues) SetReqHeaderBytes(v int64)`

SetReqHeaderBytes sets ReqHeaderBytes field to given value.

### HasReqHeaderBytes

`func (o *PlatformValues) HasReqHeaderBytes() bool`

HasReqHeaderBytes returns a boolean if a field has been set.

### GetRespBodyBytes

`func (o *PlatformValues) GetRespBodyBytes() int64`

GetRespBodyBytes returns the RespBodyBytes field if non-nil, zero value otherwise.

### GetRespBodyBytesOk

`func (o *PlatformValues) GetRespBodyBytesOk() (*int64, bool)`

GetRespBodyBytesOk returns a tuple with the RespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRespBodyBytes

`func (o *PlatformValues) SetRespBodyBytes(v int64)`

SetRespBodyBytes sets RespBodyBytes field to given value.

### HasRespBodyBytes

`func (o *PlatformValues) HasRespBodyBytes() bool`

HasRespBodyBytes returns a boolean if a field has been set.

### GetRespHeaderBytes

`func (o *PlatformValues) GetRespHeaderBytes() int64`

GetRespHeaderBytes returns the RespHeaderBytes field if non-nil, zero value otherwise.

### GetRespHeaderBytesOk

`func (o *PlatformValues) GetRespHeaderBytesOk() (*int64, bool)`

GetRespHeaderBytesOk returns a tuple with the RespHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRespHeaderBytes

`func (o *PlatformValues) SetRespHeaderBytes(v int64)`

SetRespHeaderBytes sets RespHeaderBytes field to given value.

### HasRespHeaderBytes

`func (o *PlatformValues) HasRespHeaderBytes() bool`

HasRespHeaderBytes returns a boolean if a field has been set.

### GetBereqBodyBytes

`func (o *PlatformValues) GetBereqBodyBytes() int64`

GetBereqBodyBytes returns the BereqBodyBytes field if non-nil, zero value otherwise.

### GetBereqBodyBytesOk

`func (o *PlatformValues) GetBereqBodyBytesOk() (*int64, bool)`

GetBereqBodyBytesOk returns a tuple with the BereqBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBereqBodyBytes

`func (o *PlatformValues) SetBereqBodyBytes(v int64)`

SetBereqBodyBytes sets BereqBodyBytes field to given value.

### HasBereqBodyBytes

`func (o *PlatformValues) HasBereqBodyBytes() bool`

HasBereqBodyBytes returns a boolean if a field has been set.

### GetBereqHeaderBytes

`func (o *PlatformValues) GetBereqHeaderBytes() int64`

GetBereqHeaderBytes returns the BereqHeaderBytes field if non-nil, zero value otherwise.

### GetBereqHeaderBytesOk

`func (o *PlatformValues) GetBereqHeaderBytesOk() (*int64, bool)`

GetBereqHeaderBytesOk returns a tuple with the BereqHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBereqHeaderBytes

`func (o *PlatformValues) SetBereqHeaderBytes(v int64)`

SetBereqHeaderBytes sets BereqHeaderBytes field to given value.

### HasBereqHeaderBytes

`func (o *PlatformValues) HasBereqHeaderBytes() bool`

HasBereqHeaderBytes returns a boolean if a field has been set.

### GetUncacheable

`func (o *PlatformValues) GetUncacheable() int64`

GetUncacheable returns the Uncacheable field if non-nil, zero value otherwise.

### GetUncacheableOk

`func (o *PlatformValues) GetUncacheableOk() (*int64, bool)`

GetUncacheableOk returns a tuple with the Uncacheable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUncacheable

`func (o *PlatformValues) SetUncacheable(v int64)`

SetUncacheable sets Uncacheable field to given value.

### HasUncacheable

`func (o *PlatformValues) HasUncacheable() bool`

HasUncacheable returns a boolean if a field has been set.

### GetPipe

`func (o *PlatformValues) GetPipe() int64`

GetPipe returns the Pipe field if non-nil, zero value otherwise.

### GetPipeOk

`func (o *PlatformValues) GetPipeOk() (*int64, bool)`

GetPipeOk returns a tuple with the Pipe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipe

`func (o *PlatformValues) SetPipe(v int64)`

SetPipe sets Pipe field to given value.

### HasPipe

`func (o *PlatformValues) HasPipe() bool`

HasPipe returns a boolean if a field has been set.

### GetSynth

`func (o *PlatformValues) GetSynth() int64`

GetSynth returns the Synth field if non-nil, zero value otherwise.

### GetSynthOk

`func (o *PlatformValues) GetSynthOk() (*int64, bool)`

GetSynthOk returns a tuple with the Synth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynth

`func (o *PlatformValues) SetSynth(v int64)`

SetSynth sets Synth field to given value.

### HasSynth

`func (o *PlatformValues) HasSynth() bool`

HasSynth returns a boolean if a field has been set.

### GetTLS

`func (o *PlatformValues) GetTLS() int64`

GetTLS returns the TLS field if non-nil, zero value otherwise.

### GetTLSOk

`func (o *PlatformValues) GetTLSOk() (*int64, bool)`

GetTLSOk returns a tuple with the TLS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTLS

`func (o *PlatformValues) SetTLS(v int64)`

SetTLS sets TLS field to given value.

### HasTLS

`func (o *PlatformValues) HasTLS() bool`

HasTLS returns a boolean if a field has been set.

### GetTLSV10

`func (o *PlatformValues) GetTLSV10() int64`

GetTLSV10 returns the TLSV10 field if non-nil, zero value otherwise.

### GetTLSV10Ok

`func (o *PlatformValues) GetTLSV10Ok() (*int64, bool)`

GetTLSV10Ok returns a tuple with the TLSV10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTLSV10

`func (o *PlatformValues) SetTLSV10(v int64)`

SetTLSV10 sets TLSV10 field to given value.

### HasTLSV10

`func (o *PlatformValues) HasTLSV10() bool`

HasTLSV10 returns a boolean if a field has been set.

### GetTLSV11

`func (o *PlatformValues) GetTLSV11() int64`

GetTLSV11 returns the TLSV11 field if non-nil, zero value otherwise.

### GetTLSV11Ok

`func (o *PlatformValues) GetTLSV11Ok() (*int64, bool)`

GetTLSV11Ok returns a tuple with the TLSV11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTLSV11

`func (o *PlatformValues) SetTLSV11(v int64)`

SetTLSV11 sets TLSV11 field to given value.

### HasTLSV11

`func (o *PlatformValues) HasTLSV11() bool`

HasTLSV11 returns a boolean if a field has been set.

### GetTLSV12

`func (o *PlatformValues) GetTLSV12() int64`

GetTLSV12 returns the TLSV12 field if non-nil, zero value otherwise.

### GetTLSV12Ok

`func (o *PlatformValues) GetTLSV12Ok() (*int64, bool)`

GetTLSV12Ok returns a tuple with the TLSV12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTLSV12

`func (o *PlatformValues) SetTLSV12(v int64)`

SetTLSV12 sets TLSV12 field to given value.

### HasTLSV12

`func (o *PlatformValues) HasTLSV12() bool`

HasTLSV12 returns a boolean if a field has been set.

### GetTLSV13

`func (o *PlatformValues) GetTLSV13() int64`

GetTLSV13 returns the TLSV13 field if non-nil, zero value otherwise.

### GetTLSV13Ok

`func (o *PlatformValues) GetTLSV13Ok() (*int64, bool)`

GetTLSV13Ok returns a tuple with the TLSV13 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTLSV13

`func (o *PlatformValues) SetTLSV13(v int64)`

SetTLSV13 sets TLSV13 field to given value.

### HasTLSV13

`func (o *PlatformValues) HasTLSV13() bool`

HasTLSV13 returns a boolean if a field has been set.

### GetEdgeRequests

`func (o *PlatformValues) GetEdgeRequests() int64`

GetEdgeRequests returns the EdgeRequests field if non-nil, zero value otherwise.

### GetEdgeRequestsOk

`func (o *PlatformValues) GetEdgeRequestsOk() (*int64, bool)`

GetEdgeRequestsOk returns a tuple with the EdgeRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeRequests

`func (o *PlatformValues) SetEdgeRequests(v int64)`

SetEdgeRequests sets EdgeRequests field to given value.

### HasEdgeRequests

`func (o *PlatformValues) HasEdgeRequests() bool`

HasEdgeRequests returns a boolean if a field has been set.

### GetEdgeRespHeaderBytes

`func (o *PlatformValues) GetEdgeRespHeaderBytes() int64`

GetEdgeRespHeaderBytes returns the EdgeRespHeaderBytes field if non-nil, zero value otherwise.

### GetEdgeRespHeaderBytesOk

`func (o *PlatformValues) GetEdgeRespHeaderBytesOk() (*int64, bool)`

GetEdgeRespHeaderBytesOk returns a tuple with the EdgeRespHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeRespHeaderBytes

`func (o *PlatformValues) SetEdgeRespHeaderBytes(v int64)`

SetEdgeRespHeaderBytes sets EdgeRespHeaderBytes field to given value.

### HasEdgeRespHeaderBytes

`func (o *PlatformValues) HasEdgeRespHeaderBytes() bool`

HasEdgeRespHeaderBytes returns a boolean if a field has been set.

### GetEdgeRespBodyBytes

`func (o *PlatformValues) GetEdgeRespBodyBytes() int64`

GetEdgeRespBodyBytes returns the EdgeRespBodyBytes field if non-nil, zero value otherwise.

### GetEdgeRespBodyBytesOk

`func (o *PlatformValues) GetEdgeRespBodyBytesOk() (*int64, bool)`

GetEdgeRespBodyBytesOk returns a tuple with the EdgeRespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeRespBodyBytes

`func (o *PlatformValues) SetEdgeRespBodyBytes(v int64)`

SetEdgeRespBodyBytes sets EdgeRespBodyBytes field to given value.

### HasEdgeRespBodyBytes

`func (o *PlatformValues) HasEdgeRespBodyBytes() bool`

HasEdgeRespBodyBytes returns a boolean if a field has been set.

### GetEdgeHitRequests

`func (o *PlatformValues) GetEdgeHitRequests() int64`

GetEdgeHitRequests returns the EdgeHitRequests field if non-nil, zero value otherwise.

### GetEdgeHitRequestsOk

`func (o *PlatformValues) GetEdgeHitRequestsOk() (*int64, bool)`

GetEdgeHitRequestsOk returns a tuple with the EdgeHitRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeHitRequests

`func (o *PlatformValues) SetEdgeHitRequests(v int64)`

SetEdgeHitRequests sets EdgeHitRequests field to given value.

### HasEdgeHitRequests

`func (o *PlatformValues) HasEdgeHitRequests() bool`

HasEdgeHitRequests returns a boolean if a field has been set.

### GetEdgeMissRequests

`func (o *PlatformValues) GetEdgeMissRequests() int64`

GetEdgeMissRequests returns the EdgeMissRequests field if non-nil, zero value otherwise.

### GetEdgeMissRequestsOk

`func (o *PlatformValues) GetEdgeMissRequestsOk() (*int64, bool)`

GetEdgeMissRequestsOk returns a tuple with the EdgeMissRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeMissRequests

`func (o *PlatformValues) SetEdgeMissRequests(v int64)`

SetEdgeMissRequests sets EdgeMissRequests field to given value.

### HasEdgeMissRequests

`func (o *PlatformValues) HasEdgeMissRequests() bool`

HasEdgeMissRequests returns a boolean if a field has been set.

### GetOriginFetches

`func (o *PlatformValues) GetOriginFetches() int64`

GetOriginFetches returns the OriginFetches field if non-nil, zero value otherwise.

### GetOriginFetchesOk

`func (o *PlatformValues) GetOriginFetchesOk() (*int64, bool)`

GetOriginFetchesOk returns a tuple with the OriginFetches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginFetches

`func (o *PlatformValues) SetOriginFetches(v int64)`

SetOriginFetches sets OriginFetches field to given value.

### HasOriginFetches

`func (o *PlatformValues) HasOriginFetches() bool`

HasOriginFetches returns a boolean if a field has been set.

### GetOriginFetchHeaderBytes

`func (o *PlatformValues) GetOriginFetchHeaderBytes() int64`

GetOriginFetchHeaderBytes returns the OriginFetchHeaderBytes field if non-nil, zero value otherwise.

### GetOriginFetchHeaderBytesOk

`func (o *PlatformValues) GetOriginFetchHeaderBytesOk() (*int64, bool)`

GetOriginFetchHeaderBytesOk returns a tuple with the OriginFetchHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginFetchHeaderBytes

`func (o *PlatformValues) SetOriginFetchHeaderBytes(v int64)`

SetOriginFetchHeaderBytes sets OriginFetchHeaderBytes field to given value.

### HasOriginFetchHeaderBytes

`func (o *PlatformValues) HasOriginFetchHeaderBytes() bool`

HasOriginFetchHeaderBytes returns a boolean if a field has been set.

### GetOriginFetchBodyBytes

`func (o *PlatformValues) GetOriginFetchBodyBytes() int64`

GetOriginFetchBodyBytes returns the OriginFetchBodyBytes field if non-nil, zero value otherwise.

### GetOriginFetchBodyBytesOk

`func (o *PlatformValues) GetOriginFetchBodyBytesOk() (*int64, bool)`

GetOriginFetchBodyBytesOk returns a tuple with the OriginFetchBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginFetchBodyBytes

`func (o *PlatformValues) SetOriginFetchBodyBytes(v int64)`

SetOriginFetchBodyBytes sets OriginFetchBodyBytes field to given value.

### HasOriginFetchBodyBytes

`func (o *PlatformValues) HasOriginFetchBodyBytes() bool`

HasOriginFetchBodyBytes returns a boolean if a field has been set.

### GetOriginFetchRespHeaderBytes

`func (o *PlatformValues) GetOriginFetchRespHeaderBytes() int64`

GetOriginFetchRespHeaderBytes returns the OriginFetchRespHeaderBytes field if non-nil, zero value otherwise.

### GetOriginFetchRespHeaderBytesOk

`func (o *PlatformValues) GetOriginFetchRespHeaderBytesOk() (*int64, bool)`

GetOriginFetchRespHeaderBytesOk returns a tuple with the OriginFetchRespHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginFetchRespHeaderBytes

`func (o *PlatformValues) SetOriginFetchRespHeaderBytes(v int64)`

SetOriginFetchRespHeaderBytes sets OriginFetchRespHeaderBytes field to given value.

### HasOriginFetchRespHeaderBytes

`func (o *PlatformValues) HasOriginFetchRespHeaderBytes() bool`

HasOriginFetchRespHeaderBytes returns a boolean if a field has been set.

### GetOriginFetchRespBodyBytes

`func (o *PlatformValues) GetOriginFetchRespBodyBytes() int64`

GetOriginFetchRespBodyBytes returns the OriginFetchRespBodyBytes field if non-nil, zero value otherwise.

### GetOriginFetchRespBodyBytesOk

`func (o *PlatformValues) GetOriginFetchRespBodyBytesOk() (*int64, bool)`

GetOriginFetchRespBodyBytesOk returns a tuple with the OriginFetchRespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginFetchRespBodyBytes

`func (o *PlatformValues) SetOriginFetchRespBodyBytes(v int64)`

SetOriginFetchRespBodyBytes sets OriginFetchRespBodyBytes field to given value.

### HasOriginFetchRespBodyBytes

`func (o *PlatformValues) HasOriginFetchRespBodyBytes() bool`

HasOriginFetchRespBodyBytes returns a boolean if a field has been set.

### GetOriginRevalidations

`func (o *PlatformValues) GetOriginRevalidations() int64`

GetOriginRevalidations returns the OriginRevalidations field if non-nil, zero value otherwise.

### GetOriginRevalidationsOk

`func (o *PlatformValues) GetOriginRevalidationsOk() (*int64, bool)`

GetOriginRevalidationsOk returns a tuple with the OriginRevalidations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginRevalidations

`func (o *PlatformValues) SetOriginRevalidations(v int64)`

SetOriginRevalidations sets OriginRevalidations field to given value.

### HasOriginRevalidations

`func (o *PlatformValues) HasOriginRevalidations() bool`

HasOriginRevalidations returns a boolean if a field has been set.

### GetOriginCacheFetches

`func (o *PlatformValues) GetOriginCacheFetches() int64`

GetOriginCacheFetches returns the OriginCacheFetches field if non-nil, zero value otherwise.

### GetOriginCacheFetchesOk

`func (o *PlatformValues) GetOriginCacheFetchesOk() (*int64, bool)`

GetOriginCacheFetchesOk returns a tuple with the OriginCacheFetches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginCacheFetches

`func (o *PlatformValues) SetOriginCacheFetches(v int64)`

SetOriginCacheFetches sets OriginCacheFetches field to given value.

### HasOriginCacheFetches

`func (o *PlatformValues) HasOriginCacheFetches() bool`

HasOriginCacheFetches returns a boolean if a field has been set.

### GetShield

`func (o *PlatformValues) GetShield() int64`

GetShield returns the Shield field if non-nil, zero value otherwise.

### GetShieldOk

`func (o *PlatformValues) GetShieldOk() (*int64, bool)`

GetShieldOk returns a tuple with the Shield field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShield

`func (o *PlatformValues) SetShield(v int64)`

SetShield sets Shield field to given value.

### HasShield

`func (o *PlatformValues) HasShield() bool`

HasShield returns a boolean if a field has been set.

### GetShieldRespBodyBytes

`func (o *PlatformValues) GetShieldRespBodyBytes() int64`

GetShieldRespBodyBytes returns the ShieldRespBodyBytes field if non-nil, zero value otherwise.

### GetShieldRespBodyBytesOk

`func (o *PlatformValues) GetShieldRespBodyBytesOk() (*int64, bool)`

GetShieldRespBodyBytesOk returns a tuple with the ShieldRespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShieldRespBodyBytes

`func (o *PlatformValues) SetShieldRespBodyBytes(v int64)`

SetShieldRespBodyBytes sets ShieldRespBodyBytes field to given value.

### HasShieldRespBodyBytes

`func (o *PlatformValues) HasShieldRespBodyBytes() bool`

HasShieldRespBodyBytes returns a boolean if a field has been set.

### GetShieldRespHeaderBytes

`func (o *PlatformValues) GetShieldRespHeaderBytes() int64`

GetShieldRespHeaderBytes returns the ShieldRespHeaderBytes field if non-nil, zero value otherwise.

### GetShieldRespHeaderBytesOk

`func (o *PlatformValues) GetShieldRespHeaderBytesOk() (*int64, bool)`

GetShieldRespHeaderBytesOk returns a tuple with the ShieldRespHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShieldRespHeaderBytes

`func (o *PlatformValues) SetShieldRespHeaderBytes(v int64)`

SetShieldRespHeaderBytes sets ShieldRespHeaderBytes field to given value.

### HasShieldRespHeaderBytes

`func (o *PlatformValues) HasShieldRespHeaderBytes() bool`

HasShieldRespHeaderBytes returns a boolean if a field has been set.

### GetShieldFetches

`func (o *PlatformValues) GetShieldFetches() int64`

GetShieldFetches returns the ShieldFetches field if non-nil, zero value otherwise.

### GetShieldFetchesOk

`func (o *PlatformValues) GetShieldFetchesOk() (*int64, bool)`

GetShieldFetchesOk returns a tuple with the ShieldFetches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShieldFetches

`func (o *PlatformValues) SetShieldFetches(v int64)`

SetShieldFetches sets ShieldFetches field to given value.

### HasShieldFetches

`func (o *PlatformValues) HasShieldFetches() bool`

HasShieldFetches returns a boolean if a field has been set.

### GetShieldFetchHeaderBytes

`func (o *PlatformValues) GetShieldFetchHeaderBytes() int64`

GetShieldFetchHeaderBytes returns the ShieldFetchHeaderBytes field if non-nil, zero value otherwise.

### GetShieldFetchHeaderBytesOk

`func (o *PlatformValues) GetShieldFetchHeaderBytesOk() (*int64, bool)`

GetShieldFetchHeaderBytesOk returns a tuple with the ShieldFetchHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShieldFetchHeaderBytes

`func (o *PlatformValues) SetShieldFetchHeaderBytes(v int64)`

SetShieldFetchHeaderBytes sets ShieldFetchHeaderBytes field to given value.

### HasShieldFetchHeaderBytes

`func (o *PlatformValues) HasShieldFetchHeaderBytes() bool`

HasShieldFetchHeaderBytes returns a boolean if a field has been set.

### GetShieldFetchBodyBytes

`func (o *PlatformValues) GetShieldFetchBodyBytes() int64`

GetShieldFetchBodyBytes returns the ShieldFetchBodyBytes field if non-nil, zero value otherwise.

### GetShieldFetchBodyBytesOk

`func (o *PlatformValues) GetShieldFetchBodyBytesOk() (*int64, bool)`

GetShieldFetchBodyBytesOk returns a tuple with the ShieldFetchBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShieldFetchBodyBytes

`func (o *PlatformValues) SetShieldFetchBodyBytes(v int64)`

SetShieldFetchBodyBytes sets ShieldFetchBodyBytes field to given value.

### HasShieldFetchBodyBytes

`func (o *PlatformValues) HasShieldFetchBodyBytes() bool`

HasShieldFetchBodyBytes returns a boolean if a field has been set.

### GetShieldFetchRespHeaderBytes

`func (o *PlatformValues) GetShieldFetchRespHeaderBytes() int64`

GetShieldFetchRespHeaderBytes returns the ShieldFetchRespHeaderBytes field if non-nil, zero value otherwise.

### GetShieldFetchRespHeaderBytesOk

`func (o *PlatformValues) GetShieldFetchRespHeaderBytesOk() (*int64, bool)`

GetShieldFetchRespHeaderBytesOk returns a tuple with the ShieldFetchRespHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShieldFetchRespHeaderBytes

`func (o *PlatformValues) SetShieldFetchRespHeaderBytes(v int64)`

SetShieldFetchRespHeaderBytes sets ShieldFetchRespHeaderBytes field to given value.

### HasShieldFetchRespHeaderBytes

`func (o *PlatformValues) HasShieldFetchRespHeaderBytes() bool`

HasShieldFetchRespHeaderBytes returns a boolean if a field has been set.

### GetShieldFetchRespBodyBytes

`func (o *PlatformValues) GetShieldFetchRespBodyBytes() int64`

GetShieldFetchRespBodyBytes returns the ShieldFetchRespBodyBytes field if non-nil, zero value otherwise.

### GetShieldFetchRespBodyBytesOk

`func (o *PlatformValues) GetShieldFetchRespBodyBytesOk() (*int64, bool)`

GetShieldFetchRespBodyBytesOk returns a tuple with the ShieldFetchRespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShieldFetchRespBodyBytes

`func (o *PlatformValues) SetShieldFetchRespBodyBytes(v int64)`

SetShieldFetchRespBodyBytes sets ShieldFetchRespBodyBytes field to given value.

### HasShieldFetchRespBodyBytes

`func (o *PlatformValues) HasShieldFetchRespBodyBytes() bool`

HasShieldFetchRespBodyBytes returns a boolean if a field has been set.

### GetShieldRevalidations

`func (o *PlatformValues) GetShieldRevalidations() int64`

GetShieldRevalidations returns the ShieldRevalidations field if non-nil, zero value otherwise.

### GetShieldRevalidationsOk

`func (o *PlatformValues) GetShieldRevalidationsOk() (*int64, bool)`

GetShieldRevalidationsOk returns a tuple with the ShieldRevalidations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShieldRevalidations

`func (o *PlatformValues) SetShieldRevalidations(v int64)`

SetShieldRevalidations sets ShieldRevalidations field to given value.

### HasShieldRevalidations

`func (o *PlatformValues) HasShieldRevalidations() bool`

HasShieldRevalidations returns a boolean if a field has been set.

### GetShieldCacheFetches

`func (o *PlatformValues) GetShieldCacheFetches() int64`

GetShieldCacheFetches returns the ShieldCacheFetches field if non-nil, zero value otherwise.

### GetShieldCacheFetchesOk

`func (o *PlatformValues) GetShieldCacheFetchesOk() (*int64, bool)`

GetShieldCacheFetchesOk returns a tuple with the ShieldCacheFetches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShieldCacheFetches

`func (o *PlatformValues) SetShieldCacheFetches(v int64)`

SetShieldCacheFetches sets ShieldCacheFetches field to given value.

### HasShieldCacheFetches

`func (o *PlatformValues) HasShieldCacheFetches() bool`

HasShieldCacheFetches returns a boolean if a field has been set.

### GetIpv6

`func (o *PlatformValues) GetIpv6() int64`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *PlatformValues) GetIpv6Ok() (*int64, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *PlatformValues) SetIpv6(v int64)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *PlatformValues) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.

### GetOtfp

`func (o *PlatformValues) GetOtfp() int64`

GetOtfp returns the Otfp field if non-nil, zero value otherwise.

### GetOtfpOk

`func (o *PlatformValues) GetOtfpOk() (*int64, bool)`

GetOtfpOk returns a tuple with the Otfp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtfp

`func (o *PlatformValues) SetOtfp(v int64)`

SetOtfp sets Otfp field to given value.

### HasOtfp

`func (o *PlatformValues) HasOtfp() bool`

HasOtfp returns a boolean if a field has been set.

### GetOtfpRespBodyBytes

`func (o *PlatformValues) GetOtfpRespBodyBytes() int64`

GetOtfpRespBodyBytes returns the OtfpRespBodyBytes field if non-nil, zero value otherwise.

### GetOtfpRespBodyBytesOk

`func (o *PlatformValues) GetOtfpRespBodyBytesOk() (*int64, bool)`

GetOtfpRespBodyBytesOk returns a tuple with the OtfpRespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtfpRespBodyBytes

`func (o *PlatformValues) SetOtfpRespBodyBytes(v int64)`

SetOtfpRespBodyBytes sets OtfpRespBodyBytes field to given value.

### HasOtfpRespBodyBytes

`func (o *PlatformValues) HasOtfpRespBodyBytes() bool`

HasOtfpRespBodyBytes returns a boolean if a field has been set.

### GetOtfpRespHeaderBytes

`func (o *PlatformValues) GetOtfpRespHeaderBytes() int64`

GetOtfpRespHeaderBytes returns the OtfpRespHeaderBytes field if non-nil, zero value otherwise.

### GetOtfpRespHeaderBytesOk

`func (o *PlatformValues) GetOtfpRespHeaderBytesOk() (*int64, bool)`

GetOtfpRespHeaderBytesOk returns a tuple with the OtfpRespHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtfpRespHeaderBytes

`func (o *PlatformValues) SetOtfpRespHeaderBytes(v int64)`

SetOtfpRespHeaderBytes sets OtfpRespHeaderBytes field to given value.

### HasOtfpRespHeaderBytes

`func (o *PlatformValues) HasOtfpRespHeaderBytes() bool`

HasOtfpRespHeaderBytes returns a boolean if a field has been set.

### GetOtfpShieldRespBodyBytes

`func (o *PlatformValues) GetOtfpShieldRespBodyBytes() int64`

GetOtfpShieldRespBodyBytes returns the OtfpShieldRespBodyBytes field if non-nil, zero value otherwise.

### GetOtfpShieldRespBodyBytesOk

`func (o *PlatformValues) GetOtfpShieldRespBodyBytesOk() (*int64, bool)`

GetOtfpShieldRespBodyBytesOk returns a tuple with the OtfpShieldRespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtfpShieldRespBodyBytes

`func (o *PlatformValues) SetOtfpShieldRespBodyBytes(v int64)`

SetOtfpShieldRespBodyBytes sets OtfpShieldRespBodyBytes field to given value.

### HasOtfpShieldRespBodyBytes

`func (o *PlatformValues) HasOtfpShieldRespBodyBytes() bool`

HasOtfpShieldRespBodyBytes returns a boolean if a field has been set.

### GetOtfpShieldRespHeaderBytes

`func (o *PlatformValues) GetOtfpShieldRespHeaderBytes() int64`

GetOtfpShieldRespHeaderBytes returns the OtfpShieldRespHeaderBytes field if non-nil, zero value otherwise.

### GetOtfpShieldRespHeaderBytesOk

`func (o *PlatformValues) GetOtfpShieldRespHeaderBytesOk() (*int64, bool)`

GetOtfpShieldRespHeaderBytesOk returns a tuple with the OtfpShieldRespHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtfpShieldRespHeaderBytes

`func (o *PlatformValues) SetOtfpShieldRespHeaderBytes(v int64)`

SetOtfpShieldRespHeaderBytes sets OtfpShieldRespHeaderBytes field to given value.

### HasOtfpShieldRespHeaderBytes

`func (o *PlatformValues) HasOtfpShieldRespHeaderBytes() bool`

HasOtfpShieldRespHeaderBytes returns a boolean if a field has been set.

### GetOtfpManifests

`func (o *PlatformValues) GetOtfpManifests() int64`

GetOtfpManifests returns the OtfpManifests field if non-nil, zero value otherwise.

### GetOtfpManifestsOk

`func (o *PlatformValues) GetOtfpManifestsOk() (*int64, bool)`

GetOtfpManifestsOk returns a tuple with the OtfpManifests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtfpManifests

`func (o *PlatformValues) SetOtfpManifests(v int64)`

SetOtfpManifests sets OtfpManifests field to given value.

### HasOtfpManifests

`func (o *PlatformValues) HasOtfpManifests() bool`

HasOtfpManifests returns a boolean if a field has been set.

### GetOtfpDeliverTime

`func (o *PlatformValues) GetOtfpDeliverTime() float32`

GetOtfpDeliverTime returns the OtfpDeliverTime field if non-nil, zero value otherwise.

### GetOtfpDeliverTimeOk

`func (o *PlatformValues) GetOtfpDeliverTimeOk() (*float32, bool)`

GetOtfpDeliverTimeOk returns a tuple with the OtfpDeliverTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtfpDeliverTime

`func (o *PlatformValues) SetOtfpDeliverTime(v float32)`

SetOtfpDeliverTime sets OtfpDeliverTime field to given value.

### HasOtfpDeliverTime

`func (o *PlatformValues) HasOtfpDeliverTime() bool`

HasOtfpDeliverTime returns a boolean if a field has been set.

### GetOtfpShieldTime

`func (o *PlatformValues) GetOtfpShieldTime() float32`

GetOtfpShieldTime returns the OtfpShieldTime field if non-nil, zero value otherwise.

### GetOtfpShieldTimeOk

`func (o *PlatformValues) GetOtfpShieldTimeOk() (*float32, bool)`

GetOtfpShieldTimeOk returns a tuple with the OtfpShieldTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtfpShieldTime

`func (o *PlatformValues) SetOtfpShieldTime(v float32)`

SetOtfpShieldTime sets OtfpShieldTime field to given value.

### HasOtfpShieldTime

`func (o *PlatformValues) HasOtfpShieldTime() bool`

HasOtfpShieldTime returns a boolean if a field has been set.

### GetVideo

`func (o *PlatformValues) GetVideo() int64`

GetVideo returns the Video field if non-nil, zero value otherwise.

### GetVideoOk

`func (o *PlatformValues) GetVideoOk() (*int64, bool)`

GetVideoOk returns a tuple with the Video field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideo

`func (o *PlatformValues) SetVideo(v int64)`

SetVideo sets Video field to given value.

### HasVideo

`func (o *PlatformValues) HasVideo() bool`

HasVideo returns a boolean if a field has been set.

### GetPci

`func (o *PlatformValues) GetPci() int64`

GetPci returns the Pci field if non-nil, zero value otherwise.

### GetPciOk

`func (o *PlatformValues) GetPciOk() (*int64, bool)`

GetPciOk returns a tuple with the Pci field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPci

`func (o *PlatformValues) SetPci(v int64)`

SetPci sets Pci field to given value.

### HasPci

`func (o *PlatformValues) HasPci() bool`

HasPci returns a boolean if a field has been set.

### GetLog

`func (o *PlatformValues) GetLog() int64`

GetLog returns the Log field if non-nil, zero value otherwise.

### GetLogOk

`func (o *PlatformValues) GetLogOk() (*int64, bool)`

GetLogOk returns a tuple with the Log field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLog

`func (o *PlatformValues) SetLog(v int64)`

SetLog sets Log field to given value.

### HasLog

`func (o *PlatformValues) HasLog() bool`

HasLog returns a boolean if a field has been set.

### GetLogBytes

`func (o *PlatformValues) GetLogBytes() int64`

GetLogBytes returns the LogBytes field if non-nil, zero value otherwise.

### GetLogBytesOk

`func (o *PlatformValues) GetLogBytesOk() (*int64, bool)`

GetLogBytesOk returns a tuple with the LogBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogBytes

`func (o *PlatformValues) SetLogBytes(v int64)`

SetLogBytes sets LogBytes field to given value.

### HasLogBytes

`func (o *PlatformValues) HasLogBytes() bool`

HasLogBytes returns a boolean if a field has been set.

### GetHTTP2

`func (o *PlatformValues) GetHTTP2() int64`

GetHTTP2 returns the HTTP2 field if non-nil, zero value otherwise.

### GetHTTP2Ok

`func (o *PlatformValues) GetHTTP2Ok() (*int64, bool)`

GetHTTP2Ok returns a tuple with the HTTP2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHTTP2

`func (o *PlatformValues) SetHTTP2(v int64)`

SetHTTP2 sets HTTP2 field to given value.

### HasHTTP2

`func (o *PlatformValues) HasHTTP2() bool`

HasHTTP2 returns a boolean if a field has been set.

### GetHTTP3

`func (o *PlatformValues) GetHTTP3() int64`

GetHTTP3 returns the HTTP3 field if non-nil, zero value otherwise.

### GetHTTP3Ok

`func (o *PlatformValues) GetHTTP3Ok() (*int64, bool)`

GetHTTP3Ok returns a tuple with the HTTP3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHTTP3

`func (o *PlatformValues) SetHTTP3(v int64)`

SetHTTP3 sets HTTP3 field to given value.

### HasHTTP3

`func (o *PlatformValues) HasHTTP3() bool`

HasHTTP3 returns a boolean if a field has been set.

### GetAttackReqBodyBytes

`func (o *PlatformValues) GetAttackReqBodyBytes() int64`

GetAttackReqBodyBytes returns the AttackReqBodyBytes field if non-nil, zero value otherwise.

### GetAttackReqBodyBytesOk

`func (o *PlatformValues) GetAttackReqBodyBytesOk() (*int64, bool)`

GetAttackReqBodyBytesOk returns a tuple with the AttackReqBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackReqBodyBytes

`func (o *PlatformValues) SetAttackReqBodyBytes(v int64)`

SetAttackReqBodyBytes sets AttackReqBodyBytes field to given value.

### HasAttackReqBodyBytes

`func (o *PlatformValues) HasAttackReqBodyBytes() bool`

HasAttackReqBodyBytes returns a boolean if a field has been set.

### GetAttackReqHeaderBytes

`func (o *PlatformValues) GetAttackReqHeaderBytes() int64`

GetAttackReqHeaderBytes returns the AttackReqHeaderBytes field if non-nil, zero value otherwise.

### GetAttackReqHeaderBytesOk

`func (o *PlatformValues) GetAttackReqHeaderBytesOk() (*int64, bool)`

GetAttackReqHeaderBytesOk returns a tuple with the AttackReqHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackReqHeaderBytes

`func (o *PlatformValues) SetAttackReqHeaderBytes(v int64)`

SetAttackReqHeaderBytes sets AttackReqHeaderBytes field to given value.

### HasAttackReqHeaderBytes

`func (o *PlatformValues) HasAttackReqHeaderBytes() bool`

HasAttackReqHeaderBytes returns a boolean if a field has been set.

### GetAttackLoggedReqBodyBytes

`func (o *PlatformValues) GetAttackLoggedReqBodyBytes() int64`

GetAttackLoggedReqBodyBytes returns the AttackLoggedReqBodyBytes field if non-nil, zero value otherwise.

### GetAttackLoggedReqBodyBytesOk

`func (o *PlatformValues) GetAttackLoggedReqBodyBytesOk() (*int64, bool)`

GetAttackLoggedReqBodyBytesOk returns a tuple with the AttackLoggedReqBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackLoggedReqBodyBytes

`func (o *PlatformValues) SetAttackLoggedReqBodyBytes(v int64)`

SetAttackLoggedReqBodyBytes sets AttackLoggedReqBodyBytes field to given value.

### HasAttackLoggedReqBodyBytes

`func (o *PlatformValues) HasAttackLoggedReqBodyBytes() bool`

HasAttackLoggedReqBodyBytes returns a boolean if a field has been set.

### GetAttackLoggedReqHeaderBytes

`func (o *PlatformValues) GetAttackLoggedReqHeaderBytes() int64`

GetAttackLoggedReqHeaderBytes returns the AttackLoggedReqHeaderBytes field if non-nil, zero value otherwise.

### GetAttackLoggedReqHeaderBytesOk

`func (o *PlatformValues) GetAttackLoggedReqHeaderBytesOk() (*int64, bool)`

GetAttackLoggedReqHeaderBytesOk returns a tuple with the AttackLoggedReqHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackLoggedReqHeaderBytes

`func (o *PlatformValues) SetAttackLoggedReqHeaderBytes(v int64)`

SetAttackLoggedReqHeaderBytes sets AttackLoggedReqHeaderBytes field to given value.

### HasAttackLoggedReqHeaderBytes

`func (o *PlatformValues) HasAttackLoggedReqHeaderBytes() bool`

HasAttackLoggedReqHeaderBytes returns a boolean if a field has been set.

### GetAttackBlockedReqBodyBytes

`func (o *PlatformValues) GetAttackBlockedReqBodyBytes() int64`

GetAttackBlockedReqBodyBytes returns the AttackBlockedReqBodyBytes field if non-nil, zero value otherwise.

### GetAttackBlockedReqBodyBytesOk

`func (o *PlatformValues) GetAttackBlockedReqBodyBytesOk() (*int64, bool)`

GetAttackBlockedReqBodyBytesOk returns a tuple with the AttackBlockedReqBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackBlockedReqBodyBytes

`func (o *PlatformValues) SetAttackBlockedReqBodyBytes(v int64)`

SetAttackBlockedReqBodyBytes sets AttackBlockedReqBodyBytes field to given value.

### HasAttackBlockedReqBodyBytes

`func (o *PlatformValues) HasAttackBlockedReqBodyBytes() bool`

HasAttackBlockedReqBodyBytes returns a boolean if a field has been set.

### GetAttackBlockedReqHeaderBytes

`func (o *PlatformValues) GetAttackBlockedReqHeaderBytes() int64`

GetAttackBlockedReqHeaderBytes returns the AttackBlockedReqHeaderBytes field if non-nil, zero value otherwise.

### GetAttackBlockedReqHeaderBytesOk

`func (o *PlatformValues) GetAttackBlockedReqHeaderBytesOk() (*int64, bool)`

GetAttackBlockedReqHeaderBytesOk returns a tuple with the AttackBlockedReqHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackBlockedReqHeaderBytes

`func (o *PlatformValues) SetAttackBlockedReqHeaderBytes(v int64)`

SetAttackBlockedReqHeaderBytes sets AttackBlockedReqHeaderBytes field to given value.

### HasAttackBlockedReqHeaderBytes

`func (o *PlatformValues) HasAttackBlockedReqHeaderBytes() bool`

HasAttackBlockedReqHeaderBytes returns a boolean if a field has been set.

### GetAttackPassedReqBodyBytes

`func (o *PlatformValues) GetAttackPassedReqBodyBytes() int64`

GetAttackPassedReqBodyBytes returns the AttackPassedReqBodyBytes field if non-nil, zero value otherwise.

### GetAttackPassedReqBodyBytesOk

`func (o *PlatformValues) GetAttackPassedReqBodyBytesOk() (*int64, bool)`

GetAttackPassedReqBodyBytesOk returns a tuple with the AttackPassedReqBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackPassedReqBodyBytes

`func (o *PlatformValues) SetAttackPassedReqBodyBytes(v int64)`

SetAttackPassedReqBodyBytes sets AttackPassedReqBodyBytes field to given value.

### HasAttackPassedReqBodyBytes

`func (o *PlatformValues) HasAttackPassedReqBodyBytes() bool`

HasAttackPassedReqBodyBytes returns a boolean if a field has been set.

### GetAttackPassedReqHeaderBytes

`func (o *PlatformValues) GetAttackPassedReqHeaderBytes() int64`

GetAttackPassedReqHeaderBytes returns the AttackPassedReqHeaderBytes field if non-nil, zero value otherwise.

### GetAttackPassedReqHeaderBytesOk

`func (o *PlatformValues) GetAttackPassedReqHeaderBytesOk() (*int64, bool)`

GetAttackPassedReqHeaderBytesOk returns a tuple with the AttackPassedReqHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackPassedReqHeaderBytes

`func (o *PlatformValues) SetAttackPassedReqHeaderBytes(v int64)`

SetAttackPassedReqHeaderBytes sets AttackPassedReqHeaderBytes field to given value.

### HasAttackPassedReqHeaderBytes

`func (o *PlatformValues) HasAttackPassedReqHeaderBytes() bool`

HasAttackPassedReqHeaderBytes returns a boolean if a field has been set.

### GetAttackRespSynthBytes

`func (o *PlatformValues) GetAttackRespSynthBytes() int64`

GetAttackRespSynthBytes returns the AttackRespSynthBytes field if non-nil, zero value otherwise.

### GetAttackRespSynthBytesOk

`func (o *PlatformValues) GetAttackRespSynthBytesOk() (*int64, bool)`

GetAttackRespSynthBytesOk returns a tuple with the AttackRespSynthBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackRespSynthBytes

`func (o *PlatformValues) SetAttackRespSynthBytes(v int64)`

SetAttackRespSynthBytes sets AttackRespSynthBytes field to given value.

### HasAttackRespSynthBytes

`func (o *PlatformValues) HasAttackRespSynthBytes() bool`

HasAttackRespSynthBytes returns a boolean if a field has been set.

### GetImgopto

`func (o *PlatformValues) GetImgopto() int64`

GetImgopto returns the Imgopto field if non-nil, zero value otherwise.

### GetImgoptoOk

`func (o *PlatformValues) GetImgoptoOk() (*int64, bool)`

GetImgoptoOk returns a tuple with the Imgopto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImgopto

`func (o *PlatformValues) SetImgopto(v int64)`

SetImgopto sets Imgopto field to given value.

### HasImgopto

`func (o *PlatformValues) HasImgopto() bool`

HasImgopto returns a boolean if a field has been set.

### GetImgoptoInputImageBytes

`func (o *PlatformValues) GetImgoptoInputImageBytes() int64`

GetImgoptoInputImageBytes returns the ImgoptoInputImageBytes field if non-nil, zero value otherwise.

### GetImgoptoInputImageBytesOk

`func (o *PlatformValues) GetImgoptoInputImageBytesOk() (*int64, bool)`

GetImgoptoInputImageBytesOk returns a tuple with the ImgoptoInputImageBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImgoptoInputImageBytes

`func (o *PlatformValues) SetImgoptoInputImageBytes(v int64)`

SetImgoptoInputImageBytes sets ImgoptoInputImageBytes field to given value.

### HasImgoptoInputImageBytes

`func (o *PlatformValues) HasImgoptoInputImageBytes() bool`

HasImgoptoInputImageBytes returns a boolean if a field has been set.

### GetImgoptoOutputImageBytes

`func (o *PlatformValues) GetImgoptoOutputImageBytes() int64`

GetImgoptoOutputImageBytes returns the ImgoptoOutputImageBytes field if non-nil, zero value otherwise.

### GetImgoptoOutputImageBytesOk

`func (o *PlatformValues) GetImgoptoOutputImageBytesOk() (*int64, bool)`

GetImgoptoOutputImageBytesOk returns a tuple with the ImgoptoOutputImageBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImgoptoOutputImageBytes

`func (o *PlatformValues) SetImgoptoOutputImageBytes(v int64)`

SetImgoptoOutputImageBytes sets ImgoptoOutputImageBytes field to given value.

### HasImgoptoOutputImageBytes

`func (o *PlatformValues) HasImgoptoOutputImageBytes() bool`

HasImgoptoOutputImageBytes returns a boolean if a field has been set.

### GetImgoptoRespBodyBytes

`func (o *PlatformValues) GetImgoptoRespBodyBytes() int64`

GetImgoptoRespBodyBytes returns the ImgoptoRespBodyBytes field if non-nil, zero value otherwise.

### GetImgoptoRespBodyBytesOk

`func (o *PlatformValues) GetImgoptoRespBodyBytesOk() (*int64, bool)`

GetImgoptoRespBodyBytesOk returns a tuple with the ImgoptoRespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImgoptoRespBodyBytes

`func (o *PlatformValues) SetImgoptoRespBodyBytes(v int64)`

SetImgoptoRespBodyBytes sets ImgoptoRespBodyBytes field to given value.

### HasImgoptoRespBodyBytes

`func (o *PlatformValues) HasImgoptoRespBodyBytes() bool`

HasImgoptoRespBodyBytes returns a boolean if a field has been set.

### GetImgoptoRespHeaderBytes

`func (o *PlatformValues) GetImgoptoRespHeaderBytes() int64`

GetImgoptoRespHeaderBytes returns the ImgoptoRespHeaderBytes field if non-nil, zero value otherwise.

### GetImgoptoRespHeaderBytesOk

`func (o *PlatformValues) GetImgoptoRespHeaderBytesOk() (*int64, bool)`

GetImgoptoRespHeaderBytesOk returns a tuple with the ImgoptoRespHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImgoptoRespHeaderBytes

`func (o *PlatformValues) SetImgoptoRespHeaderBytes(v int64)`

SetImgoptoRespHeaderBytes sets ImgoptoRespHeaderBytes field to given value.

### HasImgoptoRespHeaderBytes

`func (o *PlatformValues) HasImgoptoRespHeaderBytes() bool`

HasImgoptoRespHeaderBytes returns a boolean if a field has been set.

### GetImgoptoShield

`func (o *PlatformValues) GetImgoptoShield() int64`

GetImgoptoShield returns the ImgoptoShield field if non-nil, zero value otherwise.

### GetImgoptoShieldOk

`func (o *PlatformValues) GetImgoptoShieldOk() (*int64, bool)`

GetImgoptoShieldOk returns a tuple with the ImgoptoShield field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImgoptoShield

`func (o *PlatformValues) SetImgoptoShield(v int64)`

SetImgoptoShield sets ImgoptoShield field to given value.

### HasImgoptoShield

`func (o *PlatformValues) HasImgoptoShield() bool`

HasImgoptoShield returns a boolean if a field has been set.

### GetImgoptoShieldRespBodyBytes

`func (o *PlatformValues) GetImgoptoShieldRespBodyBytes() int64`

GetImgoptoShieldRespBodyBytes returns the ImgoptoShieldRespBodyBytes field if non-nil, zero value otherwise.

### GetImgoptoShieldRespBodyBytesOk

`func (o *PlatformValues) GetImgoptoShieldRespBodyBytesOk() (*int64, bool)`

GetImgoptoShieldRespBodyBytesOk returns a tuple with the ImgoptoShieldRespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImgoptoShieldRespBodyBytes

`func (o *PlatformValues) SetImgoptoShieldRespBodyBytes(v int64)`

SetImgoptoShieldRespBodyBytes sets ImgoptoShieldRespBodyBytes field to given value.

### HasImgoptoShieldRespBodyBytes

`func (o *PlatformValues) HasImgoptoShieldRespBodyBytes() bool`

HasImgoptoShieldRespBodyBytes returns a boolean if a field has been set.

### GetImgoptoShieldRespHeaderBytes

`func (o *PlatformValues) GetImgoptoShieldRespHeaderBytes() int64`

GetImgoptoShieldRespHeaderBytes returns the ImgoptoShieldRespHeaderBytes field if non-nil, zero value otherwise.

### GetImgoptoShieldRespHeaderBytesOk

`func (o *PlatformValues) GetImgoptoShieldRespHeaderBytesOk() (*int64, bool)`

GetImgoptoShieldRespHeaderBytesOk returns a tuple with the ImgoptoShieldRespHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImgoptoShieldRespHeaderBytes

`func (o *PlatformValues) SetImgoptoShieldRespHeaderBytes(v int64)`

SetImgoptoShieldRespHeaderBytes sets ImgoptoShieldRespHeaderBytes field to given value.

### HasImgoptoShieldRespHeaderBytes

`func (o *PlatformValues) HasImgoptoShieldRespHeaderBytes() bool`

HasImgoptoShieldRespHeaderBytes returns a boolean if a field has been set.

### GetImgoptoTransforms

`func (o *PlatformValues) GetImgoptoTransforms() int64`

GetImgoptoTransforms returns the ImgoptoTransforms field if non-nil, zero value otherwise.

### GetImgoptoTransformsOk

`func (o *PlatformValues) GetImgoptoTransformsOk() (*int64, bool)`

GetImgoptoTransformsOk returns a tuple with the ImgoptoTransforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImgoptoTransforms

`func (o *PlatformValues) SetImgoptoTransforms(v int64)`

SetImgoptoTransforms sets ImgoptoTransforms field to given value.

### HasImgoptoTransforms

`func (o *PlatformValues) HasImgoptoTransforms() bool`

HasImgoptoTransforms returns a boolean if a field has been set.

### GetImgvideo

`func (o *PlatformValues) GetImgvideo() int64`

GetImgvideo returns the Imgvideo field if non-nil, zero value otherwise.

### GetImgvideoOk

`func (o *PlatformValues) GetImgvideoOk() (*int64, bool)`

GetImgvideoOk returns a tuple with the Imgvideo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImgvideo

`func (o *PlatformValues) SetImgvideo(v int64)`

SetImgvideo sets Imgvideo field to given value.

### HasImgvideo

`func (o *PlatformValues) HasImgvideo() bool`

HasImgvideo returns a boolean if a field has been set.

### GetImgvideoFrames

`func (o *PlatformValues) GetImgvideoFrames() int64`

GetImgvideoFrames returns the ImgvideoFrames field if non-nil, zero value otherwise.

### GetImgvideoFramesOk

`func (o *PlatformValues) GetImgvideoFramesOk() (*int64, bool)`

GetImgvideoFramesOk returns a tuple with the ImgvideoFrames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImgvideoFrames

`func (o *PlatformValues) SetImgvideoFrames(v int64)`

SetImgvideoFrames sets ImgvideoFrames field to given value.

### HasImgvideoFrames

`func (o *PlatformValues) HasImgvideoFrames() bool`

HasImgvideoFrames returns a boolean if a field has been set.

### GetImgvideoRespHeaderBytes

`func (o *PlatformValues) GetImgvideoRespHeaderBytes() int64`

GetImgvideoRespHeaderBytes returns the ImgvideoRespHeaderBytes field if non-nil, zero value otherwise.

### GetImgvideoRespHeaderBytesOk

`func (o *PlatformValues) GetImgvideoRespHeaderBytesOk() (*int64, bool)`

GetImgvideoRespHeaderBytesOk returns a tuple with the ImgvideoRespHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImgvideoRespHeaderBytes

`func (o *PlatformValues) SetImgvideoRespHeaderBytes(v int64)`

SetImgvideoRespHeaderBytes sets ImgvideoRespHeaderBytes field to given value.

### HasImgvideoRespHeaderBytes

`func (o *PlatformValues) HasImgvideoRespHeaderBytes() bool`

HasImgvideoRespHeaderBytes returns a boolean if a field has been set.

### GetImgvideoRespBodyBytes

`func (o *PlatformValues) GetImgvideoRespBodyBytes() int64`

GetImgvideoRespBodyBytes returns the ImgvideoRespBodyBytes field if non-nil, zero value otherwise.

### GetImgvideoRespBodyBytesOk

`func (o *PlatformValues) GetImgvideoRespBodyBytesOk() (*int64, bool)`

GetImgvideoRespBodyBytesOk returns a tuple with the ImgvideoRespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImgvideoRespBodyBytes

`func (o *PlatformValues) SetImgvideoRespBodyBytes(v int64)`

SetImgvideoRespBodyBytes sets ImgvideoRespBodyBytes field to given value.

### HasImgvideoRespBodyBytes

`func (o *PlatformValues) HasImgvideoRespBodyBytes() bool`

HasImgvideoRespBodyBytes returns a boolean if a field has been set.

### GetImgvideoShieldRespHeaderBytes

`func (o *PlatformValues) GetImgvideoShieldRespHeaderBytes() int64`

GetImgvideoShieldRespHeaderBytes returns the ImgvideoShieldRespHeaderBytes field if non-nil, zero value otherwise.

### GetImgvideoShieldRespHeaderBytesOk

`func (o *PlatformValues) GetImgvideoShieldRespHeaderBytesOk() (*int64, bool)`

GetImgvideoShieldRespHeaderBytesOk returns a tuple with the ImgvideoShieldRespHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImgvideoShieldRespHeaderBytes

`func (o *PlatformValues) SetImgvideoShieldRespHeaderBytes(v int64)`

SetImgvideoShieldRespHeaderBytes sets ImgvideoShieldRespHeaderBytes field to given value.

### HasImgvideoShieldRespHeaderBytes

`func (o *PlatformValues) HasImgvideoShieldRespHeaderBytes() bool`

HasImgvideoShieldRespHeaderBytes returns a boolean if a field has been set.

### GetImgvideoShieldRespBodyBytes

`func (o *PlatformValues) GetImgvideoShieldRespBodyBytes() int64`

GetImgvideoShieldRespBodyBytes returns the ImgvideoShieldRespBodyBytes field if non-nil, zero value otherwise.

### GetImgvideoShieldRespBodyBytesOk

`func (o *PlatformValues) GetImgvideoShieldRespBodyBytesOk() (*int64, bool)`

GetImgvideoShieldRespBodyBytesOk returns a tuple with the ImgvideoShieldRespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImgvideoShieldRespBodyBytes

`func (o *PlatformValues) SetImgvideoShieldRespBodyBytes(v int64)`

SetImgvideoShieldRespBodyBytes sets ImgvideoShieldRespBodyBytes field to given value.

### HasImgvideoShieldRespBodyBytes

`func (o *PlatformValues) HasImgvideoShieldRespBodyBytes() bool`

HasImgvideoShieldRespBodyBytes returns a boolean if a field has been set.

### GetImgvideoShield

`func (o *PlatformValues) GetImgvideoShield() int64`

GetImgvideoShield returns the ImgvideoShield field if non-nil, zero value otherwise.

### GetImgvideoShieldOk

`func (o *PlatformValues) GetImgvideoShieldOk() (*int64, bool)`

GetImgvideoShieldOk returns a tuple with the ImgvideoShield field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImgvideoShield

`func (o *PlatformValues) SetImgvideoShield(v int64)`

SetImgvideoShield sets ImgvideoShield field to given value.

### HasImgvideoShield

`func (o *PlatformValues) HasImgvideoShield() bool`

HasImgvideoShield returns a boolean if a field has been set.

### GetImgvideoShieldFrames

`func (o *PlatformValues) GetImgvideoShieldFrames() int64`

GetImgvideoShieldFrames returns the ImgvideoShieldFrames field if non-nil, zero value otherwise.

### GetImgvideoShieldFramesOk

`func (o *PlatformValues) GetImgvideoShieldFramesOk() (*int64, bool)`

GetImgvideoShieldFramesOk returns a tuple with the ImgvideoShieldFrames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImgvideoShieldFrames

`func (o *PlatformValues) SetImgvideoShieldFrames(v int64)`

SetImgvideoShieldFrames sets ImgvideoShieldFrames field to given value.

### HasImgvideoShieldFrames

`func (o *PlatformValues) HasImgvideoShieldFrames() bool`

HasImgvideoShieldFrames returns a boolean if a field has been set.

### GetStatus200

`func (o *PlatformValues) GetStatus200() int64`

GetStatus200 returns the Status200 field if non-nil, zero value otherwise.

### GetStatus200Ok

`func (o *PlatformValues) GetStatus200Ok() (*int64, bool)`

GetStatus200Ok returns a tuple with the Status200 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus200

`func (o *PlatformValues) SetStatus200(v int64)`

SetStatus200 sets Status200 field to given value.

### HasStatus200

`func (o *PlatformValues) HasStatus200() bool`

HasStatus200 returns a boolean if a field has been set.

### GetStatus204

`func (o *PlatformValues) GetStatus204() int64`

GetStatus204 returns the Status204 field if non-nil, zero value otherwise.

### GetStatus204Ok

`func (o *PlatformValues) GetStatus204Ok() (*int64, bool)`

GetStatus204Ok returns a tuple with the Status204 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus204

`func (o *PlatformValues) SetStatus204(v int64)`

SetStatus204 sets Status204 field to given value.

### HasStatus204

`func (o *PlatformValues) HasStatus204() bool`

HasStatus204 returns a boolean if a field has been set.

### GetStatus206

`func (o *PlatformValues) GetStatus206() int64`

GetStatus206 returns the Status206 field if non-nil, zero value otherwise.

### GetStatus206Ok

`func (o *PlatformValues) GetStatus206Ok() (*int64, bool)`

GetStatus206Ok returns a tuple with the Status206 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus206

`func (o *PlatformValues) SetStatus206(v int64)`

SetStatus206 sets Status206 field to given value.

### HasStatus206

`func (o *PlatformValues) HasStatus206() bool`

HasStatus206 returns a boolean if a field has been set.

### GetStatus301

`func (o *PlatformValues) GetStatus301() int64`

GetStatus301 returns the Status301 field if non-nil, zero value otherwise.

### GetStatus301Ok

`func (o *PlatformValues) GetStatus301Ok() (*int64, bool)`

GetStatus301Ok returns a tuple with the Status301 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus301

`func (o *PlatformValues) SetStatus301(v int64)`

SetStatus301 sets Status301 field to given value.

### HasStatus301

`func (o *PlatformValues) HasStatus301() bool`

HasStatus301 returns a boolean if a field has been set.

### GetStatus302

`func (o *PlatformValues) GetStatus302() int64`

GetStatus302 returns the Status302 field if non-nil, zero value otherwise.

### GetStatus302Ok

`func (o *PlatformValues) GetStatus302Ok() (*int64, bool)`

GetStatus302Ok returns a tuple with the Status302 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus302

`func (o *PlatformValues) SetStatus302(v int64)`

SetStatus302 sets Status302 field to given value.

### HasStatus302

`func (o *PlatformValues) HasStatus302() bool`

HasStatus302 returns a boolean if a field has been set.

### GetStatus304

`func (o *PlatformValues) GetStatus304() int64`

GetStatus304 returns the Status304 field if non-nil, zero value otherwise.

### GetStatus304Ok

`func (o *PlatformValues) GetStatus304Ok() (*int64, bool)`

GetStatus304Ok returns a tuple with the Status304 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus304

`func (o *PlatformValues) SetStatus304(v int64)`

SetStatus304 sets Status304 field to given value.

### HasStatus304

`func (o *PlatformValues) HasStatus304() bool`

HasStatus304 returns a boolean if a field has been set.

### GetStatus400

`func (o *PlatformValues) GetStatus400() int64`

GetStatus400 returns the Status400 field if non-nil, zero value otherwise.

### GetStatus400Ok

`func (o *PlatformValues) GetStatus400Ok() (*int64, bool)`

GetStatus400Ok returns a tuple with the Status400 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus400

`func (o *PlatformValues) SetStatus400(v int64)`

SetStatus400 sets Status400 field to given value.

### HasStatus400

`func (o *PlatformValues) HasStatus400() bool`

HasStatus400 returns a boolean if a field has been set.

### GetStatus401

`func (o *PlatformValues) GetStatus401() int64`

GetStatus401 returns the Status401 field if non-nil, zero value otherwise.

### GetStatus401Ok

`func (o *PlatformValues) GetStatus401Ok() (*int64, bool)`

GetStatus401Ok returns a tuple with the Status401 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus401

`func (o *PlatformValues) SetStatus401(v int64)`

SetStatus401 sets Status401 field to given value.

### HasStatus401

`func (o *PlatformValues) HasStatus401() bool`

HasStatus401 returns a boolean if a field has been set.

### GetStatus403

`func (o *PlatformValues) GetStatus403() int64`

GetStatus403 returns the Status403 field if non-nil, zero value otherwise.

### GetStatus403Ok

`func (o *PlatformValues) GetStatus403Ok() (*int64, bool)`

GetStatus403Ok returns a tuple with the Status403 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus403

`func (o *PlatformValues) SetStatus403(v int64)`

SetStatus403 sets Status403 field to given value.

### HasStatus403

`func (o *PlatformValues) HasStatus403() bool`

HasStatus403 returns a boolean if a field has been set.

### GetStatus404

`func (o *PlatformValues) GetStatus404() int64`

GetStatus404 returns the Status404 field if non-nil, zero value otherwise.

### GetStatus404Ok

`func (o *PlatformValues) GetStatus404Ok() (*int64, bool)`

GetStatus404Ok returns a tuple with the Status404 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus404

`func (o *PlatformValues) SetStatus404(v int64)`

SetStatus404 sets Status404 field to given value.

### HasStatus404

`func (o *PlatformValues) HasStatus404() bool`

HasStatus404 returns a boolean if a field has been set.

### GetStatus406

`func (o *PlatformValues) GetStatus406() int64`

GetStatus406 returns the Status406 field if non-nil, zero value otherwise.

### GetStatus406Ok

`func (o *PlatformValues) GetStatus406Ok() (*int64, bool)`

GetStatus406Ok returns a tuple with the Status406 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus406

`func (o *PlatformValues) SetStatus406(v int64)`

SetStatus406 sets Status406 field to given value.

### HasStatus406

`func (o *PlatformValues) HasStatus406() bool`

HasStatus406 returns a boolean if a field has been set.

### GetStatus416

`func (o *PlatformValues) GetStatus416() int64`

GetStatus416 returns the Status416 field if non-nil, zero value otherwise.

### GetStatus416Ok

`func (o *PlatformValues) GetStatus416Ok() (*int64, bool)`

GetStatus416Ok returns a tuple with the Status416 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus416

`func (o *PlatformValues) SetStatus416(v int64)`

SetStatus416 sets Status416 field to given value.

### HasStatus416

`func (o *PlatformValues) HasStatus416() bool`

HasStatus416 returns a boolean if a field has been set.

### GetStatus429

`func (o *PlatformValues) GetStatus429() int64`

GetStatus429 returns the Status429 field if non-nil, zero value otherwise.

### GetStatus429Ok

`func (o *PlatformValues) GetStatus429Ok() (*int64, bool)`

GetStatus429Ok returns a tuple with the Status429 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus429

`func (o *PlatformValues) SetStatus429(v int64)`

SetStatus429 sets Status429 field to given value.

### HasStatus429

`func (o *PlatformValues) HasStatus429() bool`

HasStatus429 returns a boolean if a field has been set.

### GetStatus500

`func (o *PlatformValues) GetStatus500() int64`

GetStatus500 returns the Status500 field if non-nil, zero value otherwise.

### GetStatus500Ok

`func (o *PlatformValues) GetStatus500Ok() (*int64, bool)`

GetStatus500Ok returns a tuple with the Status500 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus500

`func (o *PlatformValues) SetStatus500(v int64)`

SetStatus500 sets Status500 field to given value.

### HasStatus500

`func (o *PlatformValues) HasStatus500() bool`

HasStatus500 returns a boolean if a field has been set.

### GetStatus501

`func (o *PlatformValues) GetStatus501() int64`

GetStatus501 returns the Status501 field if non-nil, zero value otherwise.

### GetStatus501Ok

`func (o *PlatformValues) GetStatus501Ok() (*int64, bool)`

GetStatus501Ok returns a tuple with the Status501 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus501

`func (o *PlatformValues) SetStatus501(v int64)`

SetStatus501 sets Status501 field to given value.

### HasStatus501

`func (o *PlatformValues) HasStatus501() bool`

HasStatus501 returns a boolean if a field has been set.

### GetStatus502

`func (o *PlatformValues) GetStatus502() int64`

GetStatus502 returns the Status502 field if non-nil, zero value otherwise.

### GetStatus502Ok

`func (o *PlatformValues) GetStatus502Ok() (*int64, bool)`

GetStatus502Ok returns a tuple with the Status502 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus502

`func (o *PlatformValues) SetStatus502(v int64)`

SetStatus502 sets Status502 field to given value.

### HasStatus502

`func (o *PlatformValues) HasStatus502() bool`

HasStatus502 returns a boolean if a field has been set.

### GetStatus503

`func (o *PlatformValues) GetStatus503() int64`

GetStatus503 returns the Status503 field if non-nil, zero value otherwise.

### GetStatus503Ok

`func (o *PlatformValues) GetStatus503Ok() (*int64, bool)`

GetStatus503Ok returns a tuple with the Status503 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus503

`func (o *PlatformValues) SetStatus503(v int64)`

SetStatus503 sets Status503 field to given value.

### HasStatus503

`func (o *PlatformValues) HasStatus503() bool`

HasStatus503 returns a boolean if a field has been set.

### GetStatus504

`func (o *PlatformValues) GetStatus504() int64`

GetStatus504 returns the Status504 field if non-nil, zero value otherwise.

### GetStatus504Ok

`func (o *PlatformValues) GetStatus504Ok() (*int64, bool)`

GetStatus504Ok returns a tuple with the Status504 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus504

`func (o *PlatformValues) SetStatus504(v int64)`

SetStatus504 sets Status504 field to given value.

### HasStatus504

`func (o *PlatformValues) HasStatus504() bool`

HasStatus504 returns a boolean if a field has been set.

### GetStatus505

`func (o *PlatformValues) GetStatus505() int64`

GetStatus505 returns the Status505 field if non-nil, zero value otherwise.

### GetStatus505Ok

`func (o *PlatformValues) GetStatus505Ok() (*int64, bool)`

GetStatus505Ok returns a tuple with the Status505 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus505

`func (o *PlatformValues) SetStatus505(v int64)`

SetStatus505 sets Status505 field to given value.

### HasStatus505

`func (o *PlatformValues) HasStatus505() bool`

HasStatus505 returns a boolean if a field has been set.

### GetStatus530

`func (o *PlatformValues) GetStatus530() int64`

GetStatus530 returns the Status530 field if non-nil, zero value otherwise.

### GetStatus530Ok

`func (o *PlatformValues) GetStatus530Ok() (*int64, bool)`

GetStatus530Ok returns a tuple with the Status530 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus530

`func (o *PlatformValues) SetStatus530(v int64)`

SetStatus530 sets Status530 field to given value.

### HasStatus530

`func (o *PlatformValues) HasStatus530() bool`

HasStatus530 returns a boolean if a field has been set.

### GetStatus1xx

`func (o *PlatformValues) GetStatus1xx() int64`

GetStatus1xx returns the Status1xx field if non-nil, zero value otherwise.

### GetStatus1xxOk

`func (o *PlatformValues) GetStatus1xxOk() (*int64, bool)`

GetStatus1xxOk returns a tuple with the Status1xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus1xx

`func (o *PlatformValues) SetStatus1xx(v int64)`

SetStatus1xx sets Status1xx field to given value.

### HasStatus1xx

`func (o *PlatformValues) HasStatus1xx() bool`

HasStatus1xx returns a boolean if a field has been set.

### GetStatus2xx

`func (o *PlatformValues) GetStatus2xx() int64`

GetStatus2xx returns the Status2xx field if non-nil, zero value otherwise.

### GetStatus2xxOk

`func (o *PlatformValues) GetStatus2xxOk() (*int64, bool)`

GetStatus2xxOk returns a tuple with the Status2xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus2xx

`func (o *PlatformValues) SetStatus2xx(v int64)`

SetStatus2xx sets Status2xx field to given value.

### HasStatus2xx

`func (o *PlatformValues) HasStatus2xx() bool`

HasStatus2xx returns a boolean if a field has been set.

### GetStatus3xx

`func (o *PlatformValues) GetStatus3xx() int64`

GetStatus3xx returns the Status3xx field if non-nil, zero value otherwise.

### GetStatus3xxOk

`func (o *PlatformValues) GetStatus3xxOk() (*int64, bool)`

GetStatus3xxOk returns a tuple with the Status3xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus3xx

`func (o *PlatformValues) SetStatus3xx(v int64)`

SetStatus3xx sets Status3xx field to given value.

### HasStatus3xx

`func (o *PlatformValues) HasStatus3xx() bool`

HasStatus3xx returns a boolean if a field has been set.

### GetStatus4xx

`func (o *PlatformValues) GetStatus4xx() int64`

GetStatus4xx returns the Status4xx field if non-nil, zero value otherwise.

### GetStatus4xxOk

`func (o *PlatformValues) GetStatus4xxOk() (*int64, bool)`

GetStatus4xxOk returns a tuple with the Status4xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus4xx

`func (o *PlatformValues) SetStatus4xx(v int64)`

SetStatus4xx sets Status4xx field to given value.

### HasStatus4xx

`func (o *PlatformValues) HasStatus4xx() bool`

HasStatus4xx returns a boolean if a field has been set.

### GetStatus5xx

`func (o *PlatformValues) GetStatus5xx() int64`

GetStatus5xx returns the Status5xx field if non-nil, zero value otherwise.

### GetStatus5xxOk

`func (o *PlatformValues) GetStatus5xxOk() (*int64, bool)`

GetStatus5xxOk returns a tuple with the Status5xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus5xx

`func (o *PlatformValues) SetStatus5xx(v int64)`

SetStatus5xx sets Status5xx field to given value.

### HasStatus5xx

`func (o *PlatformValues) HasStatus5xx() bool`

HasStatus5xx returns a boolean if a field has been set.

### GetObjectSize1k

`func (o *PlatformValues) GetObjectSize1k() int64`

GetObjectSize1k returns the ObjectSize1k field if non-nil, zero value otherwise.

### GetObjectSize1kOk

`func (o *PlatformValues) GetObjectSize1kOk() (*int64, bool)`

GetObjectSize1kOk returns a tuple with the ObjectSize1k field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectSize1k

`func (o *PlatformValues) SetObjectSize1k(v int64)`

SetObjectSize1k sets ObjectSize1k field to given value.

### HasObjectSize1k

`func (o *PlatformValues) HasObjectSize1k() bool`

HasObjectSize1k returns a boolean if a field has been set.

### GetObjectSize10k

`func (o *PlatformValues) GetObjectSize10k() int64`

GetObjectSize10k returns the ObjectSize10k field if non-nil, zero value otherwise.

### GetObjectSize10kOk

`func (o *PlatformValues) GetObjectSize10kOk() (*int64, bool)`

GetObjectSize10kOk returns a tuple with the ObjectSize10k field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectSize10k

`func (o *PlatformValues) SetObjectSize10k(v int64)`

SetObjectSize10k sets ObjectSize10k field to given value.

### HasObjectSize10k

`func (o *PlatformValues) HasObjectSize10k() bool`

HasObjectSize10k returns a boolean if a field has been set.

### GetObjectSize100k

`func (o *PlatformValues) GetObjectSize100k() int64`

GetObjectSize100k returns the ObjectSize100k field if non-nil, zero value otherwise.

### GetObjectSize100kOk

`func (o *PlatformValues) GetObjectSize100kOk() (*int64, bool)`

GetObjectSize100kOk returns a tuple with the ObjectSize100k field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectSize100k

`func (o *PlatformValues) SetObjectSize100k(v int64)`

SetObjectSize100k sets ObjectSize100k field to given value.

### HasObjectSize100k

`func (o *PlatformValues) HasObjectSize100k() bool`

HasObjectSize100k returns a boolean if a field has been set.

### GetObjectSize1m

`func (o *PlatformValues) GetObjectSize1m() int64`

GetObjectSize1m returns the ObjectSize1m field if non-nil, zero value otherwise.

### GetObjectSize1mOk

`func (o *PlatformValues) GetObjectSize1mOk() (*int64, bool)`

GetObjectSize1mOk returns a tuple with the ObjectSize1m field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectSize1m

`func (o *PlatformValues) SetObjectSize1m(v int64)`

SetObjectSize1m sets ObjectSize1m field to given value.

### HasObjectSize1m

`func (o *PlatformValues) HasObjectSize1m() bool`

HasObjectSize1m returns a boolean if a field has been set.

### GetObjectSize10m

`func (o *PlatformValues) GetObjectSize10m() int64`

GetObjectSize10m returns the ObjectSize10m field if non-nil, zero value otherwise.

### GetObjectSize10mOk

`func (o *PlatformValues) GetObjectSize10mOk() (*int64, bool)`

GetObjectSize10mOk returns a tuple with the ObjectSize10m field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectSize10m

`func (o *PlatformValues) SetObjectSize10m(v int64)`

SetObjectSize10m sets ObjectSize10m field to given value.

### HasObjectSize10m

`func (o *PlatformValues) HasObjectSize10m() bool`

HasObjectSize10m returns a boolean if a field has been set.

### GetObjectSize100m

`func (o *PlatformValues) GetObjectSize100m() int64`

GetObjectSize100m returns the ObjectSize100m field if non-nil, zero value otherwise.

### GetObjectSize100mOk

`func (o *PlatformValues) GetObjectSize100mOk() (*int64, bool)`

GetObjectSize100mOk returns a tuple with the ObjectSize100m field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectSize100m

`func (o *PlatformValues) SetObjectSize100m(v int64)`

SetObjectSize100m sets ObjectSize100m field to given value.

### HasObjectSize100m

`func (o *PlatformValues) HasObjectSize100m() bool`

HasObjectSize100m returns a boolean if a field has been set.

### GetObjectSize1g

`func (o *PlatformValues) GetObjectSize1g() int64`

GetObjectSize1g returns the ObjectSize1g field if non-nil, zero value otherwise.

### GetObjectSize1gOk

`func (o *PlatformValues) GetObjectSize1gOk() (*int64, bool)`

GetObjectSize1gOk returns a tuple with the ObjectSize1g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectSize1g

`func (o *PlatformValues) SetObjectSize1g(v int64)`

SetObjectSize1g sets ObjectSize1g field to given value.

### HasObjectSize1g

`func (o *PlatformValues) HasObjectSize1g() bool`

HasObjectSize1g returns a boolean if a field has been set.

### GetRecvSubTime

`func (o *PlatformValues) GetRecvSubTime() float32`

GetRecvSubTime returns the RecvSubTime field if non-nil, zero value otherwise.

### GetRecvSubTimeOk

`func (o *PlatformValues) GetRecvSubTimeOk() (*float32, bool)`

GetRecvSubTimeOk returns a tuple with the RecvSubTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecvSubTime

`func (o *PlatformValues) SetRecvSubTime(v float32)`

SetRecvSubTime sets RecvSubTime field to given value.

### HasRecvSubTime

`func (o *PlatformValues) HasRecvSubTime() bool`

HasRecvSubTime returns a boolean if a field has been set.

### GetRecvSubCount

`func (o *PlatformValues) GetRecvSubCount() int64`

GetRecvSubCount returns the RecvSubCount field if non-nil, zero value otherwise.

### GetRecvSubCountOk

`func (o *PlatformValues) GetRecvSubCountOk() (*int64, bool)`

GetRecvSubCountOk returns a tuple with the RecvSubCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecvSubCount

`func (o *PlatformValues) SetRecvSubCount(v int64)`

SetRecvSubCount sets RecvSubCount field to given value.

### HasRecvSubCount

`func (o *PlatformValues) HasRecvSubCount() bool`

HasRecvSubCount returns a boolean if a field has been set.

### GetHashSubTime

`func (o *PlatformValues) GetHashSubTime() float32`

GetHashSubTime returns the HashSubTime field if non-nil, zero value otherwise.

### GetHashSubTimeOk

`func (o *PlatformValues) GetHashSubTimeOk() (*float32, bool)`

GetHashSubTimeOk returns a tuple with the HashSubTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashSubTime

`func (o *PlatformValues) SetHashSubTime(v float32)`

SetHashSubTime sets HashSubTime field to given value.

### HasHashSubTime

`func (o *PlatformValues) HasHashSubTime() bool`

HasHashSubTime returns a boolean if a field has been set.

### GetHashSubCount

`func (o *PlatformValues) GetHashSubCount() int64`

GetHashSubCount returns the HashSubCount field if non-nil, zero value otherwise.

### GetHashSubCountOk

`func (o *PlatformValues) GetHashSubCountOk() (*int64, bool)`

GetHashSubCountOk returns a tuple with the HashSubCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashSubCount

`func (o *PlatformValues) SetHashSubCount(v int64)`

SetHashSubCount sets HashSubCount field to given value.

### HasHashSubCount

`func (o *PlatformValues) HasHashSubCount() bool`

HasHashSubCount returns a boolean if a field has been set.

### GetMissSubTime

`func (o *PlatformValues) GetMissSubTime() float32`

GetMissSubTime returns the MissSubTime field if non-nil, zero value otherwise.

### GetMissSubTimeOk

`func (o *PlatformValues) GetMissSubTimeOk() (*float32, bool)`

GetMissSubTimeOk returns a tuple with the MissSubTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissSubTime

`func (o *PlatformValues) SetMissSubTime(v float32)`

SetMissSubTime sets MissSubTime field to given value.

### HasMissSubTime

`func (o *PlatformValues) HasMissSubTime() bool`

HasMissSubTime returns a boolean if a field has been set.

### GetMissSubCount

`func (o *PlatformValues) GetMissSubCount() int64`

GetMissSubCount returns the MissSubCount field if non-nil, zero value otherwise.

### GetMissSubCountOk

`func (o *PlatformValues) GetMissSubCountOk() (*int64, bool)`

GetMissSubCountOk returns a tuple with the MissSubCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissSubCount

`func (o *PlatformValues) SetMissSubCount(v int64)`

SetMissSubCount sets MissSubCount field to given value.

### HasMissSubCount

`func (o *PlatformValues) HasMissSubCount() bool`

HasMissSubCount returns a boolean if a field has been set.

### GetFetchSubTime

`func (o *PlatformValues) GetFetchSubTime() float32`

GetFetchSubTime returns the FetchSubTime field if non-nil, zero value otherwise.

### GetFetchSubTimeOk

`func (o *PlatformValues) GetFetchSubTimeOk() (*float32, bool)`

GetFetchSubTimeOk returns a tuple with the FetchSubTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFetchSubTime

`func (o *PlatformValues) SetFetchSubTime(v float32)`

SetFetchSubTime sets FetchSubTime field to given value.

### HasFetchSubTime

`func (o *PlatformValues) HasFetchSubTime() bool`

HasFetchSubTime returns a boolean if a field has been set.

### GetFetchSubCount

`func (o *PlatformValues) GetFetchSubCount() int64`

GetFetchSubCount returns the FetchSubCount field if non-nil, zero value otherwise.

### GetFetchSubCountOk

`func (o *PlatformValues) GetFetchSubCountOk() (*int64, bool)`

GetFetchSubCountOk returns a tuple with the FetchSubCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFetchSubCount

`func (o *PlatformValues) SetFetchSubCount(v int64)`

SetFetchSubCount sets FetchSubCount field to given value.

### HasFetchSubCount

`func (o *PlatformValues) HasFetchSubCount() bool`

HasFetchSubCount returns a boolean if a field has been set.

### GetPassSubTime

`func (o *PlatformValues) GetPassSubTime() float32`

GetPassSubTime returns the PassSubTime field if non-nil, zero value otherwise.

### GetPassSubTimeOk

`func (o *PlatformValues) GetPassSubTimeOk() (*float32, bool)`

GetPassSubTimeOk returns a tuple with the PassSubTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassSubTime

`func (o *PlatformValues) SetPassSubTime(v float32)`

SetPassSubTime sets PassSubTime field to given value.

### HasPassSubTime

`func (o *PlatformValues) HasPassSubTime() bool`

HasPassSubTime returns a boolean if a field has been set.

### GetPassSubCount

`func (o *PlatformValues) GetPassSubCount() int64`

GetPassSubCount returns the PassSubCount field if non-nil, zero value otherwise.

### GetPassSubCountOk

`func (o *PlatformValues) GetPassSubCountOk() (*int64, bool)`

GetPassSubCountOk returns a tuple with the PassSubCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassSubCount

`func (o *PlatformValues) SetPassSubCount(v int64)`

SetPassSubCount sets PassSubCount field to given value.

### HasPassSubCount

`func (o *PlatformValues) HasPassSubCount() bool`

HasPassSubCount returns a boolean if a field has been set.

### GetPipeSubTime

`func (o *PlatformValues) GetPipeSubTime() float32`

GetPipeSubTime returns the PipeSubTime field if non-nil, zero value otherwise.

### GetPipeSubTimeOk

`func (o *PlatformValues) GetPipeSubTimeOk() (*float32, bool)`

GetPipeSubTimeOk returns a tuple with the PipeSubTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipeSubTime

`func (o *PlatformValues) SetPipeSubTime(v float32)`

SetPipeSubTime sets PipeSubTime field to given value.

### HasPipeSubTime

`func (o *PlatformValues) HasPipeSubTime() bool`

HasPipeSubTime returns a boolean if a field has been set.

### GetPipeSubCount

`func (o *PlatformValues) GetPipeSubCount() int64`

GetPipeSubCount returns the PipeSubCount field if non-nil, zero value otherwise.

### GetPipeSubCountOk

`func (o *PlatformValues) GetPipeSubCountOk() (*int64, bool)`

GetPipeSubCountOk returns a tuple with the PipeSubCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipeSubCount

`func (o *PlatformValues) SetPipeSubCount(v int64)`

SetPipeSubCount sets PipeSubCount field to given value.

### HasPipeSubCount

`func (o *PlatformValues) HasPipeSubCount() bool`

HasPipeSubCount returns a boolean if a field has been set.

### GetDeliverSubTime

`func (o *PlatformValues) GetDeliverSubTime() float32`

GetDeliverSubTime returns the DeliverSubTime field if non-nil, zero value otherwise.

### GetDeliverSubTimeOk

`func (o *PlatformValues) GetDeliverSubTimeOk() (*float32, bool)`

GetDeliverSubTimeOk returns a tuple with the DeliverSubTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliverSubTime

`func (o *PlatformValues) SetDeliverSubTime(v float32)`

SetDeliverSubTime sets DeliverSubTime field to given value.

### HasDeliverSubTime

`func (o *PlatformValues) HasDeliverSubTime() bool`

HasDeliverSubTime returns a boolean if a field has been set.

### GetDeliverSubCount

`func (o *PlatformValues) GetDeliverSubCount() int64`

GetDeliverSubCount returns the DeliverSubCount field if non-nil, zero value otherwise.

### GetDeliverSubCountOk

`func (o *PlatformValues) GetDeliverSubCountOk() (*int64, bool)`

GetDeliverSubCountOk returns a tuple with the DeliverSubCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliverSubCount

`func (o *PlatformValues) SetDeliverSubCount(v int64)`

SetDeliverSubCount sets DeliverSubCount field to given value.

### HasDeliverSubCount

`func (o *PlatformValues) HasDeliverSubCount() bool`

HasDeliverSubCount returns a boolean if a field has been set.

### GetErrorSubTime

`func (o *PlatformValues) GetErrorSubTime() float32`

GetErrorSubTime returns the ErrorSubTime field if non-nil, zero value otherwise.

### GetErrorSubTimeOk

`func (o *PlatformValues) GetErrorSubTimeOk() (*float32, bool)`

GetErrorSubTimeOk returns a tuple with the ErrorSubTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorSubTime

`func (o *PlatformValues) SetErrorSubTime(v float32)`

SetErrorSubTime sets ErrorSubTime field to given value.

### HasErrorSubTime

`func (o *PlatformValues) HasErrorSubTime() bool`

HasErrorSubTime returns a boolean if a field has been set.

### GetErrorSubCount

`func (o *PlatformValues) GetErrorSubCount() int64`

GetErrorSubCount returns the ErrorSubCount field if non-nil, zero value otherwise.

### GetErrorSubCountOk

`func (o *PlatformValues) GetErrorSubCountOk() (*int64, bool)`

GetErrorSubCountOk returns a tuple with the ErrorSubCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorSubCount

`func (o *PlatformValues) SetErrorSubCount(v int64)`

SetErrorSubCount sets ErrorSubCount field to given value.

### HasErrorSubCount

`func (o *PlatformValues) HasErrorSubCount() bool`

HasErrorSubCount returns a boolean if a field has been set.

### GetHitSubTime

`func (o *PlatformValues) GetHitSubTime() float32`

GetHitSubTime returns the HitSubTime field if non-nil, zero value otherwise.

### GetHitSubTimeOk

`func (o *PlatformValues) GetHitSubTimeOk() (*float32, bool)`

GetHitSubTimeOk returns a tuple with the HitSubTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHitSubTime

`func (o *PlatformValues) SetHitSubTime(v float32)`

SetHitSubTime sets HitSubTime field to given value.

### HasHitSubTime

`func (o *PlatformValues) HasHitSubTime() bool`

HasHitSubTime returns a boolean if a field has been set.

### GetHitSubCount

`func (o *PlatformValues) GetHitSubCount() int64`

GetHitSubCount returns the HitSubCount field if non-nil, zero value otherwise.

### GetHitSubCountOk

`func (o *PlatformValues) GetHitSubCountOk() (*int64, bool)`

GetHitSubCountOk returns a tuple with the HitSubCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHitSubCount

`func (o *PlatformValues) SetHitSubCount(v int64)`

SetHitSubCount sets HitSubCount field to given value.

### HasHitSubCount

`func (o *PlatformValues) HasHitSubCount() bool`

HasHitSubCount returns a boolean if a field has been set.

### GetPrehashSubTime

`func (o *PlatformValues) GetPrehashSubTime() float32`

GetPrehashSubTime returns the PrehashSubTime field if non-nil, zero value otherwise.

### GetPrehashSubTimeOk

`func (o *PlatformValues) GetPrehashSubTimeOk() (*float32, bool)`

GetPrehashSubTimeOk returns a tuple with the PrehashSubTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrehashSubTime

`func (o *PlatformValues) SetPrehashSubTime(v float32)`

SetPrehashSubTime sets PrehashSubTime field to given value.

### HasPrehashSubTime

`func (o *PlatformValues) HasPrehashSubTime() bool`

HasPrehashSubTime returns a boolean if a field has been set.

### GetPrehashSubCount

`func (o *PlatformValues) GetPrehashSubCount() int64`

GetPrehashSubCount returns the PrehashSubCount field if non-nil, zero value otherwise.

### GetPrehashSubCountOk

`func (o *PlatformValues) GetPrehashSubCountOk() (*int64, bool)`

GetPrehashSubCountOk returns a tuple with the PrehashSubCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrehashSubCount

`func (o *PlatformValues) SetPrehashSubCount(v int64)`

SetPrehashSubCount sets PrehashSubCount field to given value.

### HasPrehashSubCount

`func (o *PlatformValues) HasPrehashSubCount() bool`

HasPrehashSubCount returns a boolean if a field has been set.

### GetPredeliverSubTime

`func (o *PlatformValues) GetPredeliverSubTime() float32`

GetPredeliverSubTime returns the PredeliverSubTime field if non-nil, zero value otherwise.

### GetPredeliverSubTimeOk

`func (o *PlatformValues) GetPredeliverSubTimeOk() (*float32, bool)`

GetPredeliverSubTimeOk returns a tuple with the PredeliverSubTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredeliverSubTime

`func (o *PlatformValues) SetPredeliverSubTime(v float32)`

SetPredeliverSubTime sets PredeliverSubTime field to given value.

### HasPredeliverSubTime

`func (o *PlatformValues) HasPredeliverSubTime() bool`

HasPredeliverSubTime returns a boolean if a field has been set.

### GetPredeliverSubCount

`func (o *PlatformValues) GetPredeliverSubCount() int64`

GetPredeliverSubCount returns the PredeliverSubCount field if non-nil, zero value otherwise.

### GetPredeliverSubCountOk

`func (o *PlatformValues) GetPredeliverSubCountOk() (*int64, bool)`

GetPredeliverSubCountOk returns a tuple with the PredeliverSubCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredeliverSubCount

`func (o *PlatformValues) SetPredeliverSubCount(v int64)`

SetPredeliverSubCount sets PredeliverSubCount field to given value.

### HasPredeliverSubCount

`func (o *PlatformValues) HasPredeliverSubCount() bool`

HasPredeliverSubCount returns a boolean if a field has been set.

### GetTLSHandshakeSentBytes

`func (o *PlatformValues) GetTLSHandshakeSentBytes() int64`

GetTLSHandshakeSentBytes returns the TLSHandshakeSentBytes field if non-nil, zero value otherwise.

### GetTLSHandshakeSentBytesOk

`func (o *PlatformValues) GetTLSHandshakeSentBytesOk() (*int64, bool)`

GetTLSHandshakeSentBytesOk returns a tuple with the TLSHandshakeSentBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTLSHandshakeSentBytes

`func (o *PlatformValues) SetTLSHandshakeSentBytes(v int64)`

SetTLSHandshakeSentBytes sets TLSHandshakeSentBytes field to given value.

### HasTLSHandshakeSentBytes

`func (o *PlatformValues) HasTLSHandshakeSentBytes() bool`

HasTLSHandshakeSentBytes returns a boolean if a field has been set.

### GetHitRespBodyBytes

`func (o *PlatformValues) GetHitRespBodyBytes() int64`

GetHitRespBodyBytes returns the HitRespBodyBytes field if non-nil, zero value otherwise.

### GetHitRespBodyBytesOk

`func (o *PlatformValues) GetHitRespBodyBytesOk() (*int64, bool)`

GetHitRespBodyBytesOk returns a tuple with the HitRespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHitRespBodyBytes

`func (o *PlatformValues) SetHitRespBodyBytes(v int64)`

SetHitRespBodyBytes sets HitRespBodyBytes field to given value.

### HasHitRespBodyBytes

`func (o *PlatformValues) HasHitRespBodyBytes() bool`

HasHitRespBodyBytes returns a boolean if a field has been set.

### GetMissRespBodyBytes

`func (o *PlatformValues) GetMissRespBodyBytes() int64`

GetMissRespBodyBytes returns the MissRespBodyBytes field if non-nil, zero value otherwise.

### GetMissRespBodyBytesOk

`func (o *PlatformValues) GetMissRespBodyBytesOk() (*int64, bool)`

GetMissRespBodyBytesOk returns a tuple with the MissRespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissRespBodyBytes

`func (o *PlatformValues) SetMissRespBodyBytes(v int64)`

SetMissRespBodyBytes sets MissRespBodyBytes field to given value.

### HasMissRespBodyBytes

`func (o *PlatformValues) HasMissRespBodyBytes() bool`

HasMissRespBodyBytes returns a boolean if a field has been set.

### GetPassRespBodyBytes

`func (o *PlatformValues) GetPassRespBodyBytes() int64`

GetPassRespBodyBytes returns the PassRespBodyBytes field if non-nil, zero value otherwise.

### GetPassRespBodyBytesOk

`func (o *PlatformValues) GetPassRespBodyBytesOk() (*int64, bool)`

GetPassRespBodyBytesOk returns a tuple with the PassRespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassRespBodyBytes

`func (o *PlatformValues) SetPassRespBodyBytes(v int64)`

SetPassRespBodyBytes sets PassRespBodyBytes field to given value.

### HasPassRespBodyBytes

`func (o *PlatformValues) HasPassRespBodyBytes() bool`

HasPassRespBodyBytes returns a boolean if a field has been set.

### GetSegblockOriginFetches

`func (o *PlatformValues) GetSegblockOriginFetches() int64`

GetSegblockOriginFetches returns the SegblockOriginFetches field if non-nil, zero value otherwise.

### GetSegblockOriginFetchesOk

`func (o *PlatformValues) GetSegblockOriginFetchesOk() (*int64, bool)`

GetSegblockOriginFetchesOk returns a tuple with the SegblockOriginFetches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegblockOriginFetches

`func (o *PlatformValues) SetSegblockOriginFetches(v int64)`

SetSegblockOriginFetches sets SegblockOriginFetches field to given value.

### HasSegblockOriginFetches

`func (o *PlatformValues) HasSegblockOriginFetches() bool`

HasSegblockOriginFetches returns a boolean if a field has been set.

### GetSegblockShieldFetches

`func (o *PlatformValues) GetSegblockShieldFetches() int64`

GetSegblockShieldFetches returns the SegblockShieldFetches field if non-nil, zero value otherwise.

### GetSegblockShieldFetchesOk

`func (o *PlatformValues) GetSegblockShieldFetchesOk() (*int64, bool)`

GetSegblockShieldFetchesOk returns a tuple with the SegblockShieldFetches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegblockShieldFetches

`func (o *PlatformValues) SetSegblockShieldFetches(v int64)`

SetSegblockShieldFetches sets SegblockShieldFetches field to given value.

### HasSegblockShieldFetches

`func (o *PlatformValues) HasSegblockShieldFetches() bool`

HasSegblockShieldFetches returns a boolean if a field has been set.

### GetComputeRequests

`func (o *PlatformValues) GetComputeRequests() int64`

GetComputeRequests returns the ComputeRequests field if non-nil, zero value otherwise.

### GetComputeRequestsOk

`func (o *PlatformValues) GetComputeRequestsOk() (*int64, bool)`

GetComputeRequestsOk returns a tuple with the ComputeRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRequests

`func (o *PlatformValues) SetComputeRequests(v int64)`

SetComputeRequests sets ComputeRequests field to given value.

### HasComputeRequests

`func (o *PlatformValues) HasComputeRequests() bool`

HasComputeRequests returns a boolean if a field has been set.

### GetComputeRequestTimeMs

`func (o *PlatformValues) GetComputeRequestTimeMs() float32`

GetComputeRequestTimeMs returns the ComputeRequestTimeMs field if non-nil, zero value otherwise.

### GetComputeRequestTimeMsOk

`func (o *PlatformValues) GetComputeRequestTimeMsOk() (*float32, bool)`

GetComputeRequestTimeMsOk returns a tuple with the ComputeRequestTimeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRequestTimeMs

`func (o *PlatformValues) SetComputeRequestTimeMs(v float32)`

SetComputeRequestTimeMs sets ComputeRequestTimeMs field to given value.

### HasComputeRequestTimeMs

`func (o *PlatformValues) HasComputeRequestTimeMs() bool`

HasComputeRequestTimeMs returns a boolean if a field has been set.

### GetComputeRequestTimeBilledMs

`func (o *PlatformValues) GetComputeRequestTimeBilledMs() float32`

GetComputeRequestTimeBilledMs returns the ComputeRequestTimeBilledMs field if non-nil, zero value otherwise.

### GetComputeRequestTimeBilledMsOk

`func (o *PlatformValues) GetComputeRequestTimeBilledMsOk() (*float32, bool)`

GetComputeRequestTimeBilledMsOk returns a tuple with the ComputeRequestTimeBilledMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRequestTimeBilledMs

`func (o *PlatformValues) SetComputeRequestTimeBilledMs(v float32)`

SetComputeRequestTimeBilledMs sets ComputeRequestTimeBilledMs field to given value.

### HasComputeRequestTimeBilledMs

`func (o *PlatformValues) HasComputeRequestTimeBilledMs() bool`

HasComputeRequestTimeBilledMs returns a boolean if a field has been set.

### GetComputeRAMUsed

`func (o *PlatformValues) GetComputeRAMUsed() int64`

GetComputeRAMUsed returns the ComputeRAMUsed field if non-nil, zero value otherwise.

### GetComputeRAMUsedOk

`func (o *PlatformValues) GetComputeRAMUsedOk() (*int64, bool)`

GetComputeRAMUsedOk returns a tuple with the ComputeRAMUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRAMUsed

`func (o *PlatformValues) SetComputeRAMUsed(v int64)`

SetComputeRAMUsed sets ComputeRAMUsed field to given value.

### HasComputeRAMUsed

`func (o *PlatformValues) HasComputeRAMUsed() bool`

HasComputeRAMUsed returns a boolean if a field has been set.

### GetComputeExecutionTimeMs

`func (o *PlatformValues) GetComputeExecutionTimeMs() float32`

GetComputeExecutionTimeMs returns the ComputeExecutionTimeMs field if non-nil, zero value otherwise.

### GetComputeExecutionTimeMsOk

`func (o *PlatformValues) GetComputeExecutionTimeMsOk() (*float32, bool)`

GetComputeExecutionTimeMsOk returns a tuple with the ComputeExecutionTimeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeExecutionTimeMs

`func (o *PlatformValues) SetComputeExecutionTimeMs(v float32)`

SetComputeExecutionTimeMs sets ComputeExecutionTimeMs field to given value.

### HasComputeExecutionTimeMs

`func (o *PlatformValues) HasComputeExecutionTimeMs() bool`

HasComputeExecutionTimeMs returns a boolean if a field has been set.

### GetComputeReqHeaderBytes

`func (o *PlatformValues) GetComputeReqHeaderBytes() int64`

GetComputeReqHeaderBytes returns the ComputeReqHeaderBytes field if non-nil, zero value otherwise.

### GetComputeReqHeaderBytesOk

`func (o *PlatformValues) GetComputeReqHeaderBytesOk() (*int64, bool)`

GetComputeReqHeaderBytesOk returns a tuple with the ComputeReqHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeReqHeaderBytes

`func (o *PlatformValues) SetComputeReqHeaderBytes(v int64)`

SetComputeReqHeaderBytes sets ComputeReqHeaderBytes field to given value.

### HasComputeReqHeaderBytes

`func (o *PlatformValues) HasComputeReqHeaderBytes() bool`

HasComputeReqHeaderBytes returns a boolean if a field has been set.

### GetComputeReqBodyBytes

`func (o *PlatformValues) GetComputeReqBodyBytes() int64`

GetComputeReqBodyBytes returns the ComputeReqBodyBytes field if non-nil, zero value otherwise.

### GetComputeReqBodyBytesOk

`func (o *PlatformValues) GetComputeReqBodyBytesOk() (*int64, bool)`

GetComputeReqBodyBytesOk returns a tuple with the ComputeReqBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeReqBodyBytes

`func (o *PlatformValues) SetComputeReqBodyBytes(v int64)`

SetComputeReqBodyBytes sets ComputeReqBodyBytes field to given value.

### HasComputeReqBodyBytes

`func (o *PlatformValues) HasComputeReqBodyBytes() bool`

HasComputeReqBodyBytes returns a boolean if a field has been set.

### GetComputeRespHeaderBytes

`func (o *PlatformValues) GetComputeRespHeaderBytes() int64`

GetComputeRespHeaderBytes returns the ComputeRespHeaderBytes field if non-nil, zero value otherwise.

### GetComputeRespHeaderBytesOk

`func (o *PlatformValues) GetComputeRespHeaderBytesOk() (*int64, bool)`

GetComputeRespHeaderBytesOk returns a tuple with the ComputeRespHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespHeaderBytes

`func (o *PlatformValues) SetComputeRespHeaderBytes(v int64)`

SetComputeRespHeaderBytes sets ComputeRespHeaderBytes field to given value.

### HasComputeRespHeaderBytes

`func (o *PlatformValues) HasComputeRespHeaderBytes() bool`

HasComputeRespHeaderBytes returns a boolean if a field has been set.

### GetComputeRespBodyBytes

`func (o *PlatformValues) GetComputeRespBodyBytes() int64`

GetComputeRespBodyBytes returns the ComputeRespBodyBytes field if non-nil, zero value otherwise.

### GetComputeRespBodyBytesOk

`func (o *PlatformValues) GetComputeRespBodyBytesOk() (*int64, bool)`

GetComputeRespBodyBytesOk returns a tuple with the ComputeRespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespBodyBytes

`func (o *PlatformValues) SetComputeRespBodyBytes(v int64)`

SetComputeRespBodyBytes sets ComputeRespBodyBytes field to given value.

### HasComputeRespBodyBytes

`func (o *PlatformValues) HasComputeRespBodyBytes() bool`

HasComputeRespBodyBytes returns a boolean if a field has been set.

### GetComputeRespStatus1xx

`func (o *PlatformValues) GetComputeRespStatus1xx() int64`

GetComputeRespStatus1xx returns the ComputeRespStatus1xx field if non-nil, zero value otherwise.

### GetComputeRespStatus1xxOk

`func (o *PlatformValues) GetComputeRespStatus1xxOk() (*int64, bool)`

GetComputeRespStatus1xxOk returns a tuple with the ComputeRespStatus1xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespStatus1xx

`func (o *PlatformValues) SetComputeRespStatus1xx(v int64)`

SetComputeRespStatus1xx sets ComputeRespStatus1xx field to given value.

### HasComputeRespStatus1xx

`func (o *PlatformValues) HasComputeRespStatus1xx() bool`

HasComputeRespStatus1xx returns a boolean if a field has been set.

### GetComputeRespStatus2xx

`func (o *PlatformValues) GetComputeRespStatus2xx() int64`

GetComputeRespStatus2xx returns the ComputeRespStatus2xx field if non-nil, zero value otherwise.

### GetComputeRespStatus2xxOk

`func (o *PlatformValues) GetComputeRespStatus2xxOk() (*int64, bool)`

GetComputeRespStatus2xxOk returns a tuple with the ComputeRespStatus2xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespStatus2xx

`func (o *PlatformValues) SetComputeRespStatus2xx(v int64)`

SetComputeRespStatus2xx sets ComputeRespStatus2xx field to given value.

### HasComputeRespStatus2xx

`func (o *PlatformValues) HasComputeRespStatus2xx() bool`

HasComputeRespStatus2xx returns a boolean if a field has been set.

### GetComputeRespStatus3xx

`func (o *PlatformValues) GetComputeRespStatus3xx() int64`

GetComputeRespStatus3xx returns the ComputeRespStatus3xx field if non-nil, zero value otherwise.

### GetComputeRespStatus3xxOk

`func (o *PlatformValues) GetComputeRespStatus3xxOk() (*int64, bool)`

GetComputeRespStatus3xxOk returns a tuple with the ComputeRespStatus3xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespStatus3xx

`func (o *PlatformValues) SetComputeRespStatus3xx(v int64)`

SetComputeRespStatus3xx sets ComputeRespStatus3xx field to given value.

### HasComputeRespStatus3xx

`func (o *PlatformValues) HasComputeRespStatus3xx() bool`

HasComputeRespStatus3xx returns a boolean if a field has been set.

### GetComputeRespStatus4xx

`func (o *PlatformValues) GetComputeRespStatus4xx() int64`

GetComputeRespStatus4xx returns the ComputeRespStatus4xx field if non-nil, zero value otherwise.

### GetComputeRespStatus4xxOk

`func (o *PlatformValues) GetComputeRespStatus4xxOk() (*int64, bool)`

GetComputeRespStatus4xxOk returns a tuple with the ComputeRespStatus4xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespStatus4xx

`func (o *PlatformValues) SetComputeRespStatus4xx(v int64)`

SetComputeRespStatus4xx sets ComputeRespStatus4xx field to given value.

### HasComputeRespStatus4xx

`func (o *PlatformValues) HasComputeRespStatus4xx() bool`

HasComputeRespStatus4xx returns a boolean if a field has been set.

### GetComputeRespStatus5xx

`func (o *PlatformValues) GetComputeRespStatus5xx() int64`

GetComputeRespStatus5xx returns the ComputeRespStatus5xx field if non-nil, zero value otherwise.

### GetComputeRespStatus5xxOk

`func (o *PlatformValues) GetComputeRespStatus5xxOk() (*int64, bool)`

GetComputeRespStatus5xxOk returns a tuple with the ComputeRespStatus5xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespStatus5xx

`func (o *PlatformValues) SetComputeRespStatus5xx(v int64)`

SetComputeRespStatus5xx sets ComputeRespStatus5xx field to given value.

### HasComputeRespStatus5xx

`func (o *PlatformValues) HasComputeRespStatus5xx() bool`

HasComputeRespStatus5xx returns a boolean if a field has been set.

### GetComputeBereqHeaderBytes

`func (o *PlatformValues) GetComputeBereqHeaderBytes() int64`

GetComputeBereqHeaderBytes returns the ComputeBereqHeaderBytes field if non-nil, zero value otherwise.

### GetComputeBereqHeaderBytesOk

`func (o *PlatformValues) GetComputeBereqHeaderBytesOk() (*int64, bool)`

GetComputeBereqHeaderBytesOk returns a tuple with the ComputeBereqHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBereqHeaderBytes

`func (o *PlatformValues) SetComputeBereqHeaderBytes(v int64)`

SetComputeBereqHeaderBytes sets ComputeBereqHeaderBytes field to given value.

### HasComputeBereqHeaderBytes

`func (o *PlatformValues) HasComputeBereqHeaderBytes() bool`

HasComputeBereqHeaderBytes returns a boolean if a field has been set.

### GetComputeBereqBodyBytes

`func (o *PlatformValues) GetComputeBereqBodyBytes() int64`

GetComputeBereqBodyBytes returns the ComputeBereqBodyBytes field if non-nil, zero value otherwise.

### GetComputeBereqBodyBytesOk

`func (o *PlatformValues) GetComputeBereqBodyBytesOk() (*int64, bool)`

GetComputeBereqBodyBytesOk returns a tuple with the ComputeBereqBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBereqBodyBytes

`func (o *PlatformValues) SetComputeBereqBodyBytes(v int64)`

SetComputeBereqBodyBytes sets ComputeBereqBodyBytes field to given value.

### HasComputeBereqBodyBytes

`func (o *PlatformValues) HasComputeBereqBodyBytes() bool`

HasComputeBereqBodyBytes returns a boolean if a field has been set.

### GetComputeBerespHeaderBytes

`func (o *PlatformValues) GetComputeBerespHeaderBytes() int64`

GetComputeBerespHeaderBytes returns the ComputeBerespHeaderBytes field if non-nil, zero value otherwise.

### GetComputeBerespHeaderBytesOk

`func (o *PlatformValues) GetComputeBerespHeaderBytesOk() (*int64, bool)`

GetComputeBerespHeaderBytesOk returns a tuple with the ComputeBerespHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBerespHeaderBytes

`func (o *PlatformValues) SetComputeBerespHeaderBytes(v int64)`

SetComputeBerespHeaderBytes sets ComputeBerespHeaderBytes field to given value.

### HasComputeBerespHeaderBytes

`func (o *PlatformValues) HasComputeBerespHeaderBytes() bool`

HasComputeBerespHeaderBytes returns a boolean if a field has been set.

### GetComputeBerespBodyBytes

`func (o *PlatformValues) GetComputeBerespBodyBytes() int64`

GetComputeBerespBodyBytes returns the ComputeBerespBodyBytes field if non-nil, zero value otherwise.

### GetComputeBerespBodyBytesOk

`func (o *PlatformValues) GetComputeBerespBodyBytesOk() (*int64, bool)`

GetComputeBerespBodyBytesOk returns a tuple with the ComputeBerespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBerespBodyBytes

`func (o *PlatformValues) SetComputeBerespBodyBytes(v int64)`

SetComputeBerespBodyBytes sets ComputeBerespBodyBytes field to given value.

### HasComputeBerespBodyBytes

`func (o *PlatformValues) HasComputeBerespBodyBytes() bool`

HasComputeBerespBodyBytes returns a boolean if a field has been set.

### GetComputeBereqs

`func (o *PlatformValues) GetComputeBereqs() int64`

GetComputeBereqs returns the ComputeBereqs field if non-nil, zero value otherwise.

### GetComputeBereqsOk

`func (o *PlatformValues) GetComputeBereqsOk() (*int64, bool)`

GetComputeBereqsOk returns a tuple with the ComputeBereqs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBereqs

`func (o *PlatformValues) SetComputeBereqs(v int64)`

SetComputeBereqs sets ComputeBereqs field to given value.

### HasComputeBereqs

`func (o *PlatformValues) HasComputeBereqs() bool`

HasComputeBereqs returns a boolean if a field has been set.

### GetComputeServiceBereqError

`func (o *PlatformValues) GetComputeServiceBereqError() int64`

GetComputeServiceBereqError returns the ComputeServiceBereqError field if non-nil, zero value otherwise.

### GetComputeServiceBereqErrorOk

`func (o *PlatformValues) GetComputeServiceBereqErrorOk() (*int64, bool)`

GetComputeServiceBereqErrorOk returns a tuple with the ComputeServiceBereqError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeServiceBereqError

`func (o *PlatformValues) SetComputeServiceBereqError(v int64)`

SetComputeServiceBereqError sets ComputeServiceBereqError field to given value.

### HasComputeServiceBereqError

`func (o *PlatformValues) HasComputeServiceBereqError() bool`

HasComputeServiceBereqError returns a boolean if a field has been set.

### GetComputeServiceMemoryExceededError

`func (o *PlatformValues) GetComputeServiceMemoryExceededError() int64`

GetComputeServiceMemoryExceededError returns the ComputeServiceMemoryExceededError field if non-nil, zero value otherwise.

### GetComputeServiceMemoryExceededErrorOk

`func (o *PlatformValues) GetComputeServiceMemoryExceededErrorOk() (*int64, bool)`

GetComputeServiceMemoryExceededErrorOk returns a tuple with the ComputeServiceMemoryExceededError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeServiceMemoryExceededError

`func (o *PlatformValues) SetComputeServiceMemoryExceededError(v int64)`

SetComputeServiceMemoryExceededError sets ComputeServiceMemoryExceededError field to given value.

### HasComputeServiceMemoryExceededError

`func (o *PlatformValues) HasComputeServiceMemoryExceededError() bool`

HasComputeServiceMemoryExceededError returns a boolean if a field has been set.

### GetEdgeHitRespBodyBytes

`func (o *PlatformValues) GetEdgeHitRespBodyBytes() int64`

GetEdgeHitRespBodyBytes returns the EdgeHitRespBodyBytes field if non-nil, zero value otherwise.

### GetEdgeHitRespBodyBytesOk

`func (o *PlatformValues) GetEdgeHitRespBodyBytesOk() (*int64, bool)`

GetEdgeHitRespBodyBytesOk returns a tuple with the EdgeHitRespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeHitRespBodyBytes

`func (o *PlatformValues) SetEdgeHitRespBodyBytes(v int64)`

SetEdgeHitRespBodyBytes sets EdgeHitRespBodyBytes field to given value.

### HasEdgeHitRespBodyBytes

`func (o *PlatformValues) HasEdgeHitRespBodyBytes() bool`

HasEdgeHitRespBodyBytes returns a boolean if a field has been set.

### GetEdgeHitRespHeaderBytes

`func (o *PlatformValues) GetEdgeHitRespHeaderBytes() int64`

GetEdgeHitRespHeaderBytes returns the EdgeHitRespHeaderBytes field if non-nil, zero value otherwise.

### GetEdgeHitRespHeaderBytesOk

`func (o *PlatformValues) GetEdgeHitRespHeaderBytesOk() (*int64, bool)`

GetEdgeHitRespHeaderBytesOk returns a tuple with the EdgeHitRespHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeHitRespHeaderBytes

`func (o *PlatformValues) SetEdgeHitRespHeaderBytes(v int64)`

SetEdgeHitRespHeaderBytes sets EdgeHitRespHeaderBytes field to given value.

### HasEdgeHitRespHeaderBytes

`func (o *PlatformValues) HasEdgeHitRespHeaderBytes() bool`

HasEdgeHitRespHeaderBytes returns a boolean if a field has been set.

### GetEdgeMissRespBodyBytes

`func (o *PlatformValues) GetEdgeMissRespBodyBytes() int64`

GetEdgeMissRespBodyBytes returns the EdgeMissRespBodyBytes field if non-nil, zero value otherwise.

### GetEdgeMissRespBodyBytesOk

`func (o *PlatformValues) GetEdgeMissRespBodyBytesOk() (*int64, bool)`

GetEdgeMissRespBodyBytesOk returns a tuple with the EdgeMissRespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeMissRespBodyBytes

`func (o *PlatformValues) SetEdgeMissRespBodyBytes(v int64)`

SetEdgeMissRespBodyBytes sets EdgeMissRespBodyBytes field to given value.

### HasEdgeMissRespBodyBytes

`func (o *PlatformValues) HasEdgeMissRespBodyBytes() bool`

HasEdgeMissRespBodyBytes returns a boolean if a field has been set.

### GetEdgeMissRespHeaderBytes

`func (o *PlatformValues) GetEdgeMissRespHeaderBytes() int64`

GetEdgeMissRespHeaderBytes returns the EdgeMissRespHeaderBytes field if non-nil, zero value otherwise.

### GetEdgeMissRespHeaderBytesOk

`func (o *PlatformValues) GetEdgeMissRespHeaderBytesOk() (*int64, bool)`

GetEdgeMissRespHeaderBytesOk returns a tuple with the EdgeMissRespHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeMissRespHeaderBytes

`func (o *PlatformValues) SetEdgeMissRespHeaderBytes(v int64)`

SetEdgeMissRespHeaderBytes sets EdgeMissRespHeaderBytes field to given value.

### HasEdgeMissRespHeaderBytes

`func (o *PlatformValues) HasEdgeMissRespHeaderBytes() bool`

HasEdgeMissRespHeaderBytes returns a boolean if a field has been set.

### GetOriginCacheFetchRespBodyBytes

`func (o *PlatformValues) GetOriginCacheFetchRespBodyBytes() int64`

GetOriginCacheFetchRespBodyBytes returns the OriginCacheFetchRespBodyBytes field if non-nil, zero value otherwise.

### GetOriginCacheFetchRespBodyBytesOk

`func (o *PlatformValues) GetOriginCacheFetchRespBodyBytesOk() (*int64, bool)`

GetOriginCacheFetchRespBodyBytesOk returns a tuple with the OriginCacheFetchRespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginCacheFetchRespBodyBytes

`func (o *PlatformValues) SetOriginCacheFetchRespBodyBytes(v int64)`

SetOriginCacheFetchRespBodyBytes sets OriginCacheFetchRespBodyBytes field to given value.

### HasOriginCacheFetchRespBodyBytes

`func (o *PlatformValues) HasOriginCacheFetchRespBodyBytes() bool`

HasOriginCacheFetchRespBodyBytes returns a boolean if a field has been set.

### GetOriginCacheFetchRespHeaderBytes

`func (o *PlatformValues) GetOriginCacheFetchRespHeaderBytes() int64`

GetOriginCacheFetchRespHeaderBytes returns the OriginCacheFetchRespHeaderBytes field if non-nil, zero value otherwise.

### GetOriginCacheFetchRespHeaderBytesOk

`func (o *PlatformValues) GetOriginCacheFetchRespHeaderBytesOk() (*int64, bool)`

GetOriginCacheFetchRespHeaderBytesOk returns a tuple with the OriginCacheFetchRespHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginCacheFetchRespHeaderBytes

`func (o *PlatformValues) SetOriginCacheFetchRespHeaderBytes(v int64)`

SetOriginCacheFetchRespHeaderBytes sets OriginCacheFetchRespHeaderBytes field to given value.

### HasOriginCacheFetchRespHeaderBytes

`func (o *PlatformValues) HasOriginCacheFetchRespHeaderBytes() bool`

HasOriginCacheFetchRespHeaderBytes returns a boolean if a field has been set.

### GetShieldHitRequests

`func (o *PlatformValues) GetShieldHitRequests() int64`

GetShieldHitRequests returns the ShieldHitRequests field if non-nil, zero value otherwise.

### GetShieldHitRequestsOk

`func (o *PlatformValues) GetShieldHitRequestsOk() (*int64, bool)`

GetShieldHitRequestsOk returns a tuple with the ShieldHitRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShieldHitRequests

`func (o *PlatformValues) SetShieldHitRequests(v int64)`

SetShieldHitRequests sets ShieldHitRequests field to given value.

### HasShieldHitRequests

`func (o *PlatformValues) HasShieldHitRequests() bool`

HasShieldHitRequests returns a boolean if a field has been set.

### GetShieldMissRequests

`func (o *PlatformValues) GetShieldMissRequests() int64`

GetShieldMissRequests returns the ShieldMissRequests field if non-nil, zero value otherwise.

### GetShieldMissRequestsOk

`func (o *PlatformValues) GetShieldMissRequestsOk() (*int64, bool)`

GetShieldMissRequestsOk returns a tuple with the ShieldMissRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShieldMissRequests

`func (o *PlatformValues) SetShieldMissRequests(v int64)`

SetShieldMissRequests sets ShieldMissRequests field to given value.

### HasShieldMissRequests

`func (o *PlatformValues) HasShieldMissRequests() bool`

HasShieldMissRequests returns a boolean if a field has been set.

### GetShieldHitRespHeaderBytes

`func (o *PlatformValues) GetShieldHitRespHeaderBytes() int64`

GetShieldHitRespHeaderBytes returns the ShieldHitRespHeaderBytes field if non-nil, zero value otherwise.

### GetShieldHitRespHeaderBytesOk

`func (o *PlatformValues) GetShieldHitRespHeaderBytesOk() (*int64, bool)`

GetShieldHitRespHeaderBytesOk returns a tuple with the ShieldHitRespHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShieldHitRespHeaderBytes

`func (o *PlatformValues) SetShieldHitRespHeaderBytes(v int64)`

SetShieldHitRespHeaderBytes sets ShieldHitRespHeaderBytes field to given value.

### HasShieldHitRespHeaderBytes

`func (o *PlatformValues) HasShieldHitRespHeaderBytes() bool`

HasShieldHitRespHeaderBytes returns a boolean if a field has been set.

### GetShieldHitRespBodyBytes

`func (o *PlatformValues) GetShieldHitRespBodyBytes() int64`

GetShieldHitRespBodyBytes returns the ShieldHitRespBodyBytes field if non-nil, zero value otherwise.

### GetShieldHitRespBodyBytesOk

`func (o *PlatformValues) GetShieldHitRespBodyBytesOk() (*int64, bool)`

GetShieldHitRespBodyBytesOk returns a tuple with the ShieldHitRespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShieldHitRespBodyBytes

`func (o *PlatformValues) SetShieldHitRespBodyBytes(v int64)`

SetShieldHitRespBodyBytes sets ShieldHitRespBodyBytes field to given value.

### HasShieldHitRespBodyBytes

`func (o *PlatformValues) HasShieldHitRespBodyBytes() bool`

HasShieldHitRespBodyBytes returns a boolean if a field has been set.

### GetShieldMissRespHeaderBytes

`func (o *PlatformValues) GetShieldMissRespHeaderBytes() int64`

GetShieldMissRespHeaderBytes returns the ShieldMissRespHeaderBytes field if non-nil, zero value otherwise.

### GetShieldMissRespHeaderBytesOk

`func (o *PlatformValues) GetShieldMissRespHeaderBytesOk() (*int64, bool)`

GetShieldMissRespHeaderBytesOk returns a tuple with the ShieldMissRespHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShieldMissRespHeaderBytes

`func (o *PlatformValues) SetShieldMissRespHeaderBytes(v int64)`

SetShieldMissRespHeaderBytes sets ShieldMissRespHeaderBytes field to given value.

### HasShieldMissRespHeaderBytes

`func (o *PlatformValues) HasShieldMissRespHeaderBytes() bool`

HasShieldMissRespHeaderBytes returns a boolean if a field has been set.

### GetShieldMissRespBodyBytes

`func (o *PlatformValues) GetShieldMissRespBodyBytes() int64`

GetShieldMissRespBodyBytes returns the ShieldMissRespBodyBytes field if non-nil, zero value otherwise.

### GetShieldMissRespBodyBytesOk

`func (o *PlatformValues) GetShieldMissRespBodyBytesOk() (*int64, bool)`

GetShieldMissRespBodyBytesOk returns a tuple with the ShieldMissRespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShieldMissRespBodyBytes

`func (o *PlatformValues) SetShieldMissRespBodyBytes(v int64)`

SetShieldMissRespBodyBytes sets ShieldMissRespBodyBytes field to given value.

### HasShieldMissRespBodyBytes

`func (o *PlatformValues) HasShieldMissRespBodyBytes() bool`

HasShieldMissRespBodyBytes returns a boolean if a field has been set.

### GetWebsocketReqHeaderBytes

`func (o *PlatformValues) GetWebsocketReqHeaderBytes() int64`

GetWebsocketReqHeaderBytes returns the WebsocketReqHeaderBytes field if non-nil, zero value otherwise.

### GetWebsocketReqHeaderBytesOk

`func (o *PlatformValues) GetWebsocketReqHeaderBytesOk() (*int64, bool)`

GetWebsocketReqHeaderBytesOk returns a tuple with the WebsocketReqHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsocketReqHeaderBytes

`func (o *PlatformValues) SetWebsocketReqHeaderBytes(v int64)`

SetWebsocketReqHeaderBytes sets WebsocketReqHeaderBytes field to given value.

### HasWebsocketReqHeaderBytes

`func (o *PlatformValues) HasWebsocketReqHeaderBytes() bool`

HasWebsocketReqHeaderBytes returns a boolean if a field has been set.

### GetWebsocketReqBodyBytes

`func (o *PlatformValues) GetWebsocketReqBodyBytes() int64`

GetWebsocketReqBodyBytes returns the WebsocketReqBodyBytes field if non-nil, zero value otherwise.

### GetWebsocketReqBodyBytesOk

`func (o *PlatformValues) GetWebsocketReqBodyBytesOk() (*int64, bool)`

GetWebsocketReqBodyBytesOk returns a tuple with the WebsocketReqBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsocketReqBodyBytes

`func (o *PlatformValues) SetWebsocketReqBodyBytes(v int64)`

SetWebsocketReqBodyBytes sets WebsocketReqBodyBytes field to given value.

### HasWebsocketReqBodyBytes

`func (o *PlatformValues) HasWebsocketReqBodyBytes() bool`

HasWebsocketReqBodyBytes returns a boolean if a field has been set.

### GetWebsocketRespHeaderBytes

`func (o *PlatformValues) GetWebsocketRespHeaderBytes() int64`

GetWebsocketRespHeaderBytes returns the WebsocketRespHeaderBytes field if non-nil, zero value otherwise.

### GetWebsocketRespHeaderBytesOk

`func (o *PlatformValues) GetWebsocketRespHeaderBytesOk() (*int64, bool)`

GetWebsocketRespHeaderBytesOk returns a tuple with the WebsocketRespHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsocketRespHeaderBytes

`func (o *PlatformValues) SetWebsocketRespHeaderBytes(v int64)`

SetWebsocketRespHeaderBytes sets WebsocketRespHeaderBytes field to given value.

### HasWebsocketRespHeaderBytes

`func (o *PlatformValues) HasWebsocketRespHeaderBytes() bool`

HasWebsocketRespHeaderBytes returns a boolean if a field has been set.

### GetWebsocketRespBodyBytes

`func (o *PlatformValues) GetWebsocketRespBodyBytes() int64`

GetWebsocketRespBodyBytes returns the WebsocketRespBodyBytes field if non-nil, zero value otherwise.

### GetWebsocketRespBodyBytesOk

`func (o *PlatformValues) GetWebsocketRespBodyBytesOk() (*int64, bool)`

GetWebsocketRespBodyBytesOk returns a tuple with the WebsocketRespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsocketRespBodyBytes

`func (o *PlatformValues) SetWebsocketRespBodyBytes(v int64)`

SetWebsocketRespBodyBytes sets WebsocketRespBodyBytes field to given value.

### HasWebsocketRespBodyBytes

`func (o *PlatformValues) HasWebsocketRespBodyBytes() bool`

HasWebsocketRespBodyBytes returns a boolean if a field has been set.

### GetWebsocketBereqHeaderBytes

`func (o *PlatformValues) GetWebsocketBereqHeaderBytes() int64`

GetWebsocketBereqHeaderBytes returns the WebsocketBereqHeaderBytes field if non-nil, zero value otherwise.

### GetWebsocketBereqHeaderBytesOk

`func (o *PlatformValues) GetWebsocketBereqHeaderBytesOk() (*int64, bool)`

GetWebsocketBereqHeaderBytesOk returns a tuple with the WebsocketBereqHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsocketBereqHeaderBytes

`func (o *PlatformValues) SetWebsocketBereqHeaderBytes(v int64)`

SetWebsocketBereqHeaderBytes sets WebsocketBereqHeaderBytes field to given value.

### HasWebsocketBereqHeaderBytes

`func (o *PlatformValues) HasWebsocketBereqHeaderBytes() bool`

HasWebsocketBereqHeaderBytes returns a boolean if a field has been set.

### GetWebsocketBereqBodyBytes

`func (o *PlatformValues) GetWebsocketBereqBodyBytes() int64`

GetWebsocketBereqBodyBytes returns the WebsocketBereqBodyBytes field if non-nil, zero value otherwise.

### GetWebsocketBereqBodyBytesOk

`func (o *PlatformValues) GetWebsocketBereqBodyBytesOk() (*int64, bool)`

GetWebsocketBereqBodyBytesOk returns a tuple with the WebsocketBereqBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsocketBereqBodyBytes

`func (o *PlatformValues) SetWebsocketBereqBodyBytes(v int64)`

SetWebsocketBereqBodyBytes sets WebsocketBereqBodyBytes field to given value.

### HasWebsocketBereqBodyBytes

`func (o *PlatformValues) HasWebsocketBereqBodyBytes() bool`

HasWebsocketBereqBodyBytes returns a boolean if a field has been set.

### GetWebsocketBerespHeaderBytes

`func (o *PlatformValues) GetWebsocketBerespHeaderBytes() int64`

GetWebsocketBerespHeaderBytes returns the WebsocketBerespHeaderBytes field if non-nil, zero value otherwise.

### GetWebsocketBerespHeaderBytesOk

`func (o *PlatformValues) GetWebsocketBerespHeaderBytesOk() (*int64, bool)`

GetWebsocketBerespHeaderBytesOk returns a tuple with the WebsocketBerespHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsocketBerespHeaderBytes

`func (o *PlatformValues) SetWebsocketBerespHeaderBytes(v int64)`

SetWebsocketBerespHeaderBytes sets WebsocketBerespHeaderBytes field to given value.

### HasWebsocketBerespHeaderBytes

`func (o *PlatformValues) HasWebsocketBerespHeaderBytes() bool`

HasWebsocketBerespHeaderBytes returns a boolean if a field has been set.

### GetWebsocketBerespBodyBytes

`func (o *PlatformValues) GetWebsocketBerespBodyBytes() int64`

GetWebsocketBerespBodyBytes returns the WebsocketBerespBodyBytes field if non-nil, zero value otherwise.

### GetWebsocketBerespBodyBytesOk

`func (o *PlatformValues) GetWebsocketBerespBodyBytesOk() (*int64, bool)`

GetWebsocketBerespBodyBytesOk returns a tuple with the WebsocketBerespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsocketBerespBodyBytes

`func (o *PlatformValues) SetWebsocketBerespBodyBytes(v int64)`

SetWebsocketBerespBodyBytes sets WebsocketBerespBodyBytes field to given value.

### HasWebsocketBerespBodyBytes

`func (o *PlatformValues) HasWebsocketBerespBodyBytes() bool`

HasWebsocketBerespBodyBytes returns a boolean if a field has been set.

### GetWebsocketConnTimeMs

`func (o *PlatformValues) GetWebsocketConnTimeMs() int64`

GetWebsocketConnTimeMs returns the WebsocketConnTimeMs field if non-nil, zero value otherwise.

### GetWebsocketConnTimeMsOk

`func (o *PlatformValues) GetWebsocketConnTimeMsOk() (*int64, bool)`

GetWebsocketConnTimeMsOk returns a tuple with the WebsocketConnTimeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsocketConnTimeMs

`func (o *PlatformValues) SetWebsocketConnTimeMs(v int64)`

SetWebsocketConnTimeMs sets WebsocketConnTimeMs field to given value.

### HasWebsocketConnTimeMs

`func (o *PlatformValues) HasWebsocketConnTimeMs() bool`

HasWebsocketConnTimeMs returns a boolean if a field has been set.

### GetFanoutRecvPublishes

`func (o *PlatformValues) GetFanoutRecvPublishes() int64`

GetFanoutRecvPublishes returns the FanoutRecvPublishes field if non-nil, zero value otherwise.

### GetFanoutRecvPublishesOk

`func (o *PlatformValues) GetFanoutRecvPublishesOk() (*int64, bool)`

GetFanoutRecvPublishesOk returns a tuple with the FanoutRecvPublishes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFanoutRecvPublishes

`func (o *PlatformValues) SetFanoutRecvPublishes(v int64)`

SetFanoutRecvPublishes sets FanoutRecvPublishes field to given value.

### HasFanoutRecvPublishes

`func (o *PlatformValues) HasFanoutRecvPublishes() bool`

HasFanoutRecvPublishes returns a boolean if a field has been set.

### GetFanoutSendPublishes

`func (o *PlatformValues) GetFanoutSendPublishes() int64`

GetFanoutSendPublishes returns the FanoutSendPublishes field if non-nil, zero value otherwise.

### GetFanoutSendPublishesOk

`func (o *PlatformValues) GetFanoutSendPublishesOk() (*int64, bool)`

GetFanoutSendPublishesOk returns a tuple with the FanoutSendPublishes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFanoutSendPublishes

`func (o *PlatformValues) SetFanoutSendPublishes(v int64)`

SetFanoutSendPublishes sets FanoutSendPublishes field to given value.

### HasFanoutSendPublishes

`func (o *PlatformValues) HasFanoutSendPublishes() bool`

HasFanoutSendPublishes returns a boolean if a field has been set.

### GetKvStoreClassAOperations

`func (o *PlatformValues) GetKvStoreClassAOperations() int64`

GetKvStoreClassAOperations returns the KvStoreClassAOperations field if non-nil, zero value otherwise.

### GetKvStoreClassAOperationsOk

`func (o *PlatformValues) GetKvStoreClassAOperationsOk() (*int64, bool)`

GetKvStoreClassAOperationsOk returns a tuple with the KvStoreClassAOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKvStoreClassAOperations

`func (o *PlatformValues) SetKvStoreClassAOperations(v int64)`

SetKvStoreClassAOperations sets KvStoreClassAOperations field to given value.

### HasKvStoreClassAOperations

`func (o *PlatformValues) HasKvStoreClassAOperations() bool`

HasKvStoreClassAOperations returns a boolean if a field has been set.

### GetKvStoreClassBOperations

`func (o *PlatformValues) GetKvStoreClassBOperations() int64`

GetKvStoreClassBOperations returns the KvStoreClassBOperations field if non-nil, zero value otherwise.

### GetKvStoreClassBOperationsOk

`func (o *PlatformValues) GetKvStoreClassBOperationsOk() (*int64, bool)`

GetKvStoreClassBOperationsOk returns a tuple with the KvStoreClassBOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKvStoreClassBOperations

`func (o *PlatformValues) SetKvStoreClassBOperations(v int64)`

SetKvStoreClassBOperations sets KvStoreClassBOperations field to given value.

### HasKvStoreClassBOperations

`func (o *PlatformValues) HasKvStoreClassBOperations() bool`

HasKvStoreClassBOperations returns a boolean if a field has been set.

### GetFanoutReqHeaderBytes

`func (o *PlatformValues) GetFanoutReqHeaderBytes() int64`

GetFanoutReqHeaderBytes returns the FanoutReqHeaderBytes field if non-nil, zero value otherwise.

### GetFanoutReqHeaderBytesOk

`func (o *PlatformValues) GetFanoutReqHeaderBytesOk() (*int64, bool)`

GetFanoutReqHeaderBytesOk returns a tuple with the FanoutReqHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFanoutReqHeaderBytes

`func (o *PlatformValues) SetFanoutReqHeaderBytes(v int64)`

SetFanoutReqHeaderBytes sets FanoutReqHeaderBytes field to given value.

### HasFanoutReqHeaderBytes

`func (o *PlatformValues) HasFanoutReqHeaderBytes() bool`

HasFanoutReqHeaderBytes returns a boolean if a field has been set.

### GetFanoutReqBodyBytes

`func (o *PlatformValues) GetFanoutReqBodyBytes() int64`

GetFanoutReqBodyBytes returns the FanoutReqBodyBytes field if non-nil, zero value otherwise.

### GetFanoutReqBodyBytesOk

`func (o *PlatformValues) GetFanoutReqBodyBytesOk() (*int64, bool)`

GetFanoutReqBodyBytesOk returns a tuple with the FanoutReqBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFanoutReqBodyBytes

`func (o *PlatformValues) SetFanoutReqBodyBytes(v int64)`

SetFanoutReqBodyBytes sets FanoutReqBodyBytes field to given value.

### HasFanoutReqBodyBytes

`func (o *PlatformValues) HasFanoutReqBodyBytes() bool`

HasFanoutReqBodyBytes returns a boolean if a field has been set.

### GetFanoutRespHeaderBytes

`func (o *PlatformValues) GetFanoutRespHeaderBytes() int64`

GetFanoutRespHeaderBytes returns the FanoutRespHeaderBytes field if non-nil, zero value otherwise.

### GetFanoutRespHeaderBytesOk

`func (o *PlatformValues) GetFanoutRespHeaderBytesOk() (*int64, bool)`

GetFanoutRespHeaderBytesOk returns a tuple with the FanoutRespHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFanoutRespHeaderBytes

`func (o *PlatformValues) SetFanoutRespHeaderBytes(v int64)`

SetFanoutRespHeaderBytes sets FanoutRespHeaderBytes field to given value.

### HasFanoutRespHeaderBytes

`func (o *PlatformValues) HasFanoutRespHeaderBytes() bool`

HasFanoutRespHeaderBytes returns a boolean if a field has been set.

### GetFanoutRespBodyBytes

`func (o *PlatformValues) GetFanoutRespBodyBytes() int64`

GetFanoutRespBodyBytes returns the FanoutRespBodyBytes field if non-nil, zero value otherwise.

### GetFanoutRespBodyBytesOk

`func (o *PlatformValues) GetFanoutRespBodyBytesOk() (*int64, bool)`

GetFanoutRespBodyBytesOk returns a tuple with the FanoutRespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFanoutRespBodyBytes

`func (o *PlatformValues) SetFanoutRespBodyBytes(v int64)`

SetFanoutRespBodyBytes sets FanoutRespBodyBytes field to given value.

### HasFanoutRespBodyBytes

`func (o *PlatformValues) HasFanoutRespBodyBytes() bool`

HasFanoutRespBodyBytes returns a boolean if a field has been set.

### GetFanoutBereqHeaderBytes

`func (o *PlatformValues) GetFanoutBereqHeaderBytes() int64`

GetFanoutBereqHeaderBytes returns the FanoutBereqHeaderBytes field if non-nil, zero value otherwise.

### GetFanoutBereqHeaderBytesOk

`func (o *PlatformValues) GetFanoutBereqHeaderBytesOk() (*int64, bool)`

GetFanoutBereqHeaderBytesOk returns a tuple with the FanoutBereqHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFanoutBereqHeaderBytes

`func (o *PlatformValues) SetFanoutBereqHeaderBytes(v int64)`

SetFanoutBereqHeaderBytes sets FanoutBereqHeaderBytes field to given value.

### HasFanoutBereqHeaderBytes

`func (o *PlatformValues) HasFanoutBereqHeaderBytes() bool`

HasFanoutBereqHeaderBytes returns a boolean if a field has been set.

### GetFanoutBereqBodyBytes

`func (o *PlatformValues) GetFanoutBereqBodyBytes() int64`

GetFanoutBereqBodyBytes returns the FanoutBereqBodyBytes field if non-nil, zero value otherwise.

### GetFanoutBereqBodyBytesOk

`func (o *PlatformValues) GetFanoutBereqBodyBytesOk() (*int64, bool)`

GetFanoutBereqBodyBytesOk returns a tuple with the FanoutBereqBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFanoutBereqBodyBytes

`func (o *PlatformValues) SetFanoutBereqBodyBytes(v int64)`

SetFanoutBereqBodyBytes sets FanoutBereqBodyBytes field to given value.

### HasFanoutBereqBodyBytes

`func (o *PlatformValues) HasFanoutBereqBodyBytes() bool`

HasFanoutBereqBodyBytes returns a boolean if a field has been set.

### GetFanoutBerespHeaderBytes

`func (o *PlatformValues) GetFanoutBerespHeaderBytes() int64`

GetFanoutBerespHeaderBytes returns the FanoutBerespHeaderBytes field if non-nil, zero value otherwise.

### GetFanoutBerespHeaderBytesOk

`func (o *PlatformValues) GetFanoutBerespHeaderBytesOk() (*int64, bool)`

GetFanoutBerespHeaderBytesOk returns a tuple with the FanoutBerespHeaderBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFanoutBerespHeaderBytes

`func (o *PlatformValues) SetFanoutBerespHeaderBytes(v int64)`

SetFanoutBerespHeaderBytes sets FanoutBerespHeaderBytes field to given value.

### HasFanoutBerespHeaderBytes

`func (o *PlatformValues) HasFanoutBerespHeaderBytes() bool`

HasFanoutBerespHeaderBytes returns a boolean if a field has been set.

### GetFanoutBerespBodyBytes

`func (o *PlatformValues) GetFanoutBerespBodyBytes() int64`

GetFanoutBerespBodyBytes returns the FanoutBerespBodyBytes field if non-nil, zero value otherwise.

### GetFanoutBerespBodyBytesOk

`func (o *PlatformValues) GetFanoutBerespBodyBytesOk() (*int64, bool)`

GetFanoutBerespBodyBytesOk returns a tuple with the FanoutBerespBodyBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFanoutBerespBodyBytes

`func (o *PlatformValues) SetFanoutBerespBodyBytes(v int64)`

SetFanoutBerespBodyBytes sets FanoutBerespBodyBytes field to given value.

### HasFanoutBerespBodyBytes

`func (o *PlatformValues) HasFanoutBerespBodyBytes() bool`

HasFanoutBerespBodyBytes returns a boolean if a field has been set.

### GetFanoutConnTimeMs

`func (o *PlatformValues) GetFanoutConnTimeMs() int64`

GetFanoutConnTimeMs returns the FanoutConnTimeMs field if non-nil, zero value otherwise.

### GetFanoutConnTimeMsOk

`func (o *PlatformValues) GetFanoutConnTimeMsOk() (*int64, bool)`

GetFanoutConnTimeMsOk returns a tuple with the FanoutConnTimeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFanoutConnTimeMs

`func (o *PlatformValues) SetFanoutConnTimeMs(v int64)`

SetFanoutConnTimeMs sets FanoutConnTimeMs field to given value.

### HasFanoutConnTimeMs

`func (o *PlatformValues) HasFanoutConnTimeMs() bool`

HasFanoutConnTimeMs returns a boolean if a field has been set.

### GetDdosActionLimitStreamsConnections

`func (o *PlatformValues) GetDdosActionLimitStreamsConnections() int64`

GetDdosActionLimitStreamsConnections returns the DdosActionLimitStreamsConnections field if non-nil, zero value otherwise.

### GetDdosActionLimitStreamsConnectionsOk

`func (o *PlatformValues) GetDdosActionLimitStreamsConnectionsOk() (*int64, bool)`

GetDdosActionLimitStreamsConnectionsOk returns a tuple with the DdosActionLimitStreamsConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdosActionLimitStreamsConnections

`func (o *PlatformValues) SetDdosActionLimitStreamsConnections(v int64)`

SetDdosActionLimitStreamsConnections sets DdosActionLimitStreamsConnections field to given value.

### HasDdosActionLimitStreamsConnections

`func (o *PlatformValues) HasDdosActionLimitStreamsConnections() bool`

HasDdosActionLimitStreamsConnections returns a boolean if a field has been set.

### GetDdosActionLimitStreamsRequests

`func (o *PlatformValues) GetDdosActionLimitStreamsRequests() int64`

GetDdosActionLimitStreamsRequests returns the DdosActionLimitStreamsRequests field if non-nil, zero value otherwise.

### GetDdosActionLimitStreamsRequestsOk

`func (o *PlatformValues) GetDdosActionLimitStreamsRequestsOk() (*int64, bool)`

GetDdosActionLimitStreamsRequestsOk returns a tuple with the DdosActionLimitStreamsRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdosActionLimitStreamsRequests

`func (o *PlatformValues) SetDdosActionLimitStreamsRequests(v int64)`

SetDdosActionLimitStreamsRequests sets DdosActionLimitStreamsRequests field to given value.

### HasDdosActionLimitStreamsRequests

`func (o *PlatformValues) HasDdosActionLimitStreamsRequests() bool`

HasDdosActionLimitStreamsRequests returns a boolean if a field has been set.

### GetDdosActionTarpitAccept

`func (o *PlatformValues) GetDdosActionTarpitAccept() int64`

GetDdosActionTarpitAccept returns the DdosActionTarpitAccept field if non-nil, zero value otherwise.

### GetDdosActionTarpitAcceptOk

`func (o *PlatformValues) GetDdosActionTarpitAcceptOk() (*int64, bool)`

GetDdosActionTarpitAcceptOk returns a tuple with the DdosActionTarpitAccept field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdosActionTarpitAccept

`func (o *PlatformValues) SetDdosActionTarpitAccept(v int64)`

SetDdosActionTarpitAccept sets DdosActionTarpitAccept field to given value.

### HasDdosActionTarpitAccept

`func (o *PlatformValues) HasDdosActionTarpitAccept() bool`

HasDdosActionTarpitAccept returns a boolean if a field has been set.

### GetDdosActionTarpit

`func (o *PlatformValues) GetDdosActionTarpit() int64`

GetDdosActionTarpit returns the DdosActionTarpit field if non-nil, zero value otherwise.

### GetDdosActionTarpitOk

`func (o *PlatformValues) GetDdosActionTarpitOk() (*int64, bool)`

GetDdosActionTarpitOk returns a tuple with the DdosActionTarpit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdosActionTarpit

`func (o *PlatformValues) SetDdosActionTarpit(v int64)`

SetDdosActionTarpit sets DdosActionTarpit field to given value.

### HasDdosActionTarpit

`func (o *PlatformValues) HasDdosActionTarpit() bool`

HasDdosActionTarpit returns a boolean if a field has been set.

### GetDdosActionClose

`func (o *PlatformValues) GetDdosActionClose() int64`

GetDdosActionClose returns the DdosActionClose field if non-nil, zero value otherwise.

### GetDdosActionCloseOk

`func (o *PlatformValues) GetDdosActionCloseOk() (*int64, bool)`

GetDdosActionCloseOk returns a tuple with the DdosActionClose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdosActionClose

`func (o *PlatformValues) SetDdosActionClose(v int64)`

SetDdosActionClose sets DdosActionClose field to given value.

### HasDdosActionClose

`func (o *PlatformValues) HasDdosActionClose() bool`

HasDdosActionClose returns a boolean if a field has been set.

### GetDdosActionBlackhole

`func (o *PlatformValues) GetDdosActionBlackhole() int64`

GetDdosActionBlackhole returns the DdosActionBlackhole field if non-nil, zero value otherwise.

### GetDdosActionBlackholeOk

`func (o *PlatformValues) GetDdosActionBlackholeOk() (*int64, bool)`

GetDdosActionBlackholeOk returns a tuple with the DdosActionBlackhole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdosActionBlackhole

`func (o *PlatformValues) SetDdosActionBlackhole(v int64)`

SetDdosActionBlackhole sets DdosActionBlackhole field to given value.

### HasDdosActionBlackhole

`func (o *PlatformValues) HasDdosActionBlackhole() bool`

HasDdosActionBlackhole returns a boolean if a field has been set.

### GetBotChallengeStarts

`func (o *PlatformValues) GetBotChallengeStarts() int64`

GetBotChallengeStarts returns the BotChallengeStarts field if non-nil, zero value otherwise.

### GetBotChallengeStartsOk

`func (o *PlatformValues) GetBotChallengeStartsOk() (*int64, bool)`

GetBotChallengeStartsOk returns a tuple with the BotChallengeStarts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotChallengeStarts

`func (o *PlatformValues) SetBotChallengeStarts(v int64)`

SetBotChallengeStarts sets BotChallengeStarts field to given value.

### HasBotChallengeStarts

`func (o *PlatformValues) HasBotChallengeStarts() bool`

HasBotChallengeStarts returns a boolean if a field has been set.

### GetBotChallengeCompleteTokensPassed

`func (o *PlatformValues) GetBotChallengeCompleteTokensPassed() int64`

GetBotChallengeCompleteTokensPassed returns the BotChallengeCompleteTokensPassed field if non-nil, zero value otherwise.

### GetBotChallengeCompleteTokensPassedOk

`func (o *PlatformValues) GetBotChallengeCompleteTokensPassedOk() (*int64, bool)`

GetBotChallengeCompleteTokensPassedOk returns a tuple with the BotChallengeCompleteTokensPassed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotChallengeCompleteTokensPassed

`func (o *PlatformValues) SetBotChallengeCompleteTokensPassed(v int64)`

SetBotChallengeCompleteTokensPassed sets BotChallengeCompleteTokensPassed field to given value.

### HasBotChallengeCompleteTokensPassed

`func (o *PlatformValues) HasBotChallengeCompleteTokensPassed() bool`

HasBotChallengeCompleteTokensPassed returns a boolean if a field has been set.

### GetBotChallengeCompleteTokensFailed

`func (o *PlatformValues) GetBotChallengeCompleteTokensFailed() int64`

GetBotChallengeCompleteTokensFailed returns the BotChallengeCompleteTokensFailed field if non-nil, zero value otherwise.

### GetBotChallengeCompleteTokensFailedOk

`func (o *PlatformValues) GetBotChallengeCompleteTokensFailedOk() (*int64, bool)`

GetBotChallengeCompleteTokensFailedOk returns a tuple with the BotChallengeCompleteTokensFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotChallengeCompleteTokensFailed

`func (o *PlatformValues) SetBotChallengeCompleteTokensFailed(v int64)`

SetBotChallengeCompleteTokensFailed sets BotChallengeCompleteTokensFailed field to given value.

### HasBotChallengeCompleteTokensFailed

`func (o *PlatformValues) HasBotChallengeCompleteTokensFailed() bool`

HasBotChallengeCompleteTokensFailed returns a boolean if a field has been set.

### GetBotChallengeCompleteTokensChecked

`func (o *PlatformValues) GetBotChallengeCompleteTokensChecked() int64`

GetBotChallengeCompleteTokensChecked returns the BotChallengeCompleteTokensChecked field if non-nil, zero value otherwise.

### GetBotChallengeCompleteTokensCheckedOk

`func (o *PlatformValues) GetBotChallengeCompleteTokensCheckedOk() (*int64, bool)`

GetBotChallengeCompleteTokensCheckedOk returns a tuple with the BotChallengeCompleteTokensChecked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotChallengeCompleteTokensChecked

`func (o *PlatformValues) SetBotChallengeCompleteTokensChecked(v int64)`

SetBotChallengeCompleteTokensChecked sets BotChallengeCompleteTokensChecked field to given value.

### HasBotChallengeCompleteTokensChecked

`func (o *PlatformValues) HasBotChallengeCompleteTokensChecked() bool`

HasBotChallengeCompleteTokensChecked returns a boolean if a field has been set.

### GetBotChallengeCompleteTokensDisabled

`func (o *PlatformValues) GetBotChallengeCompleteTokensDisabled() int64`

GetBotChallengeCompleteTokensDisabled returns the BotChallengeCompleteTokensDisabled field if non-nil, zero value otherwise.

### GetBotChallengeCompleteTokensDisabledOk

`func (o *PlatformValues) GetBotChallengeCompleteTokensDisabledOk() (*int64, bool)`

GetBotChallengeCompleteTokensDisabledOk returns a tuple with the BotChallengeCompleteTokensDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotChallengeCompleteTokensDisabled

`func (o *PlatformValues) SetBotChallengeCompleteTokensDisabled(v int64)`

SetBotChallengeCompleteTokensDisabled sets BotChallengeCompleteTokensDisabled field to given value.

### HasBotChallengeCompleteTokensDisabled

`func (o *PlatformValues) HasBotChallengeCompleteTokensDisabled() bool`

HasBotChallengeCompleteTokensDisabled returns a boolean if a field has been set.

### GetBotChallengeCompleteTokensIssued

`func (o *PlatformValues) GetBotChallengeCompleteTokensIssued() int64`

GetBotChallengeCompleteTokensIssued returns the BotChallengeCompleteTokensIssued field if non-nil, zero value otherwise.

### GetBotChallengeCompleteTokensIssuedOk

`func (o *PlatformValues) GetBotChallengeCompleteTokensIssuedOk() (*int64, bool)`

GetBotChallengeCompleteTokensIssuedOk returns a tuple with the BotChallengeCompleteTokensIssued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotChallengeCompleteTokensIssued

`func (o *PlatformValues) SetBotChallengeCompleteTokensIssued(v int64)`

SetBotChallengeCompleteTokensIssued sets BotChallengeCompleteTokensIssued field to given value.

### HasBotChallengeCompleteTokensIssued

`func (o *PlatformValues) HasBotChallengeCompleteTokensIssued() bool`

HasBotChallengeCompleteTokensIssued returns a boolean if a field has been set.

### GetBotChallengesIssued

`func (o *PlatformValues) GetBotChallengesIssued() int64`

GetBotChallengesIssued returns the BotChallengesIssued field if non-nil, zero value otherwise.

### GetBotChallengesIssuedOk

`func (o *PlatformValues) GetBotChallengesIssuedOk() (*int64, bool)`

GetBotChallengesIssuedOk returns a tuple with the BotChallengesIssued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotChallengesIssued

`func (o *PlatformValues) SetBotChallengesIssued(v int64)`

SetBotChallengesIssued sets BotChallengesIssued field to given value.

### HasBotChallengesIssued

`func (o *PlatformValues) HasBotChallengesIssued() bool`

HasBotChallengesIssued returns a boolean if a field has been set.

### GetBotChallengesSucceeded

`func (o *PlatformValues) GetBotChallengesSucceeded() int64`

GetBotChallengesSucceeded returns the BotChallengesSucceeded field if non-nil, zero value otherwise.

### GetBotChallengesSucceededOk

`func (o *PlatformValues) GetBotChallengesSucceededOk() (*int64, bool)`

GetBotChallengesSucceededOk returns a tuple with the BotChallengesSucceeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotChallengesSucceeded

`func (o *PlatformValues) SetBotChallengesSucceeded(v int64)`

SetBotChallengesSucceeded sets BotChallengesSucceeded field to given value.

### HasBotChallengesSucceeded

`func (o *PlatformValues) HasBotChallengesSucceeded() bool`

HasBotChallengesSucceeded returns a boolean if a field has been set.

### GetBotChallengesFailed

`func (o *PlatformValues) GetBotChallengesFailed() int64`

GetBotChallengesFailed returns the BotChallengesFailed field if non-nil, zero value otherwise.

### GetBotChallengesFailedOk

`func (o *PlatformValues) GetBotChallengesFailedOk() (*int64, bool)`

GetBotChallengesFailedOk returns a tuple with the BotChallengesFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotChallengesFailed

`func (o *PlatformValues) SetBotChallengesFailed(v int64)`

SetBotChallengesFailed sets BotChallengesFailed field to given value.

### HasBotChallengesFailed

`func (o *PlatformValues) HasBotChallengesFailed() bool`

HasBotChallengesFailed returns a boolean if a field has been set.

### GetDdosActionDowngrade

`func (o *PlatformValues) GetDdosActionDowngrade() int64`

GetDdosActionDowngrade returns the DdosActionDowngrade field if non-nil, zero value otherwise.

### GetDdosActionDowngradeOk

`func (o *PlatformValues) GetDdosActionDowngradeOk() (*int64, bool)`

GetDdosActionDowngradeOk returns a tuple with the DdosActionDowngrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdosActionDowngrade

`func (o *PlatformValues) SetDdosActionDowngrade(v int64)`

SetDdosActionDowngrade sets DdosActionDowngrade field to given value.

### HasDdosActionDowngrade

`func (o *PlatformValues) HasDdosActionDowngrade() bool`

HasDdosActionDowngrade returns a boolean if a field has been set.

### GetDdosActionDowngradedConnections

`func (o *PlatformValues) GetDdosActionDowngradedConnections() int64`

GetDdosActionDowngradedConnections returns the DdosActionDowngradedConnections field if non-nil, zero value otherwise.

### GetDdosActionDowngradedConnectionsOk

`func (o *PlatformValues) GetDdosActionDowngradedConnectionsOk() (*int64, bool)`

GetDdosActionDowngradedConnectionsOk returns a tuple with the DdosActionDowngradedConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdosActionDowngradedConnections

`func (o *PlatformValues) SetDdosActionDowngradedConnections(v int64)`

SetDdosActionDowngradedConnections sets DdosActionDowngradedConnections field to given value.

### HasDdosActionDowngradedConnections

`func (o *PlatformValues) HasDdosActionDowngradedConnections() bool`

HasDdosActionDowngradedConnections returns a boolean if a field has been set.

### GetAllHitRequests

`func (o *PlatformValues) GetAllHitRequests() int64`

GetAllHitRequests returns the AllHitRequests field if non-nil, zero value otherwise.

### GetAllHitRequestsOk

`func (o *PlatformValues) GetAllHitRequestsOk() (*int64, bool)`

GetAllHitRequestsOk returns a tuple with the AllHitRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllHitRequests

`func (o *PlatformValues) SetAllHitRequests(v int64)`

SetAllHitRequests sets AllHitRequests field to given value.

### HasAllHitRequests

`func (o *PlatformValues) HasAllHitRequests() bool`

HasAllHitRequests returns a boolean if a field has been set.

### GetAllMissRequests

`func (o *PlatformValues) GetAllMissRequests() int64`

GetAllMissRequests returns the AllMissRequests field if non-nil, zero value otherwise.

### GetAllMissRequestsOk

`func (o *PlatformValues) GetAllMissRequestsOk() (*int64, bool)`

GetAllMissRequestsOk returns a tuple with the AllMissRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllMissRequests

`func (o *PlatformValues) SetAllMissRequests(v int64)`

SetAllMissRequests sets AllMissRequests field to given value.

### HasAllMissRequests

`func (o *PlatformValues) HasAllMissRequests() bool`

HasAllMissRequests returns a boolean if a field has been set.

### GetAllPassRequests

`func (o *PlatformValues) GetAllPassRequests() int64`

GetAllPassRequests returns the AllPassRequests field if non-nil, zero value otherwise.

### GetAllPassRequestsOk

`func (o *PlatformValues) GetAllPassRequestsOk() (*int64, bool)`

GetAllPassRequestsOk returns a tuple with the AllPassRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllPassRequests

`func (o *PlatformValues) SetAllPassRequests(v int64)`

SetAllPassRequests sets AllPassRequests field to given value.

### HasAllPassRequests

`func (o *PlatformValues) HasAllPassRequests() bool`

HasAllPassRequests returns a boolean if a field has been set.

### GetAllErrorRequests

`func (o *PlatformValues) GetAllErrorRequests() int64`

GetAllErrorRequests returns the AllErrorRequests field if non-nil, zero value otherwise.

### GetAllErrorRequestsOk

`func (o *PlatformValues) GetAllErrorRequestsOk() (*int64, bool)`

GetAllErrorRequestsOk returns a tuple with the AllErrorRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllErrorRequests

`func (o *PlatformValues) SetAllErrorRequests(v int64)`

SetAllErrorRequests sets AllErrorRequests field to given value.

### HasAllErrorRequests

`func (o *PlatformValues) HasAllErrorRequests() bool`

HasAllErrorRequests returns a boolean if a field has been set.

### GetAllSynthRequests

`func (o *PlatformValues) GetAllSynthRequests() int64`

GetAllSynthRequests returns the AllSynthRequests field if non-nil, zero value otherwise.

### GetAllSynthRequestsOk

`func (o *PlatformValues) GetAllSynthRequestsOk() (*int64, bool)`

GetAllSynthRequestsOk returns a tuple with the AllSynthRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllSynthRequests

`func (o *PlatformValues) SetAllSynthRequests(v int64)`

SetAllSynthRequests sets AllSynthRequests field to given value.

### HasAllSynthRequests

`func (o *PlatformValues) HasAllSynthRequests() bool`

HasAllSynthRequests returns a boolean if a field has been set.

### GetAllEdgeHitRequests

`func (o *PlatformValues) GetAllEdgeHitRequests() int64`

GetAllEdgeHitRequests returns the AllEdgeHitRequests field if non-nil, zero value otherwise.

### GetAllEdgeHitRequestsOk

`func (o *PlatformValues) GetAllEdgeHitRequestsOk() (*int64, bool)`

GetAllEdgeHitRequestsOk returns a tuple with the AllEdgeHitRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllEdgeHitRequests

`func (o *PlatformValues) SetAllEdgeHitRequests(v int64)`

SetAllEdgeHitRequests sets AllEdgeHitRequests field to given value.

### HasAllEdgeHitRequests

`func (o *PlatformValues) HasAllEdgeHitRequests() bool`

HasAllEdgeHitRequests returns a boolean if a field has been set.

### GetAllEdgeMissRequests

`func (o *PlatformValues) GetAllEdgeMissRequests() int64`

GetAllEdgeMissRequests returns the AllEdgeMissRequests field if non-nil, zero value otherwise.

### GetAllEdgeMissRequestsOk

`func (o *PlatformValues) GetAllEdgeMissRequestsOk() (*int64, bool)`

GetAllEdgeMissRequestsOk returns a tuple with the AllEdgeMissRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllEdgeMissRequests

`func (o *PlatformValues) SetAllEdgeMissRequests(v int64)`

SetAllEdgeMissRequests sets AllEdgeMissRequests field to given value.

### HasAllEdgeMissRequests

`func (o *PlatformValues) HasAllEdgeMissRequests() bool`

HasAllEdgeMissRequests returns a boolean if a field has been set.

### GetAllStatus1xx

`func (o *PlatformValues) GetAllStatus1xx() int64`

GetAllStatus1xx returns the AllStatus1xx field if non-nil, zero value otherwise.

### GetAllStatus1xxOk

`func (o *PlatformValues) GetAllStatus1xxOk() (*int64, bool)`

GetAllStatus1xxOk returns a tuple with the AllStatus1xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllStatus1xx

`func (o *PlatformValues) SetAllStatus1xx(v int64)`

SetAllStatus1xx sets AllStatus1xx field to given value.

### HasAllStatus1xx

`func (o *PlatformValues) HasAllStatus1xx() bool`

HasAllStatus1xx returns a boolean if a field has been set.

### GetAllStatus2xx

`func (o *PlatformValues) GetAllStatus2xx() int64`

GetAllStatus2xx returns the AllStatus2xx field if non-nil, zero value otherwise.

### GetAllStatus2xxOk

`func (o *PlatformValues) GetAllStatus2xxOk() (*int64, bool)`

GetAllStatus2xxOk returns a tuple with the AllStatus2xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllStatus2xx

`func (o *PlatformValues) SetAllStatus2xx(v int64)`

SetAllStatus2xx sets AllStatus2xx field to given value.

### HasAllStatus2xx

`func (o *PlatformValues) HasAllStatus2xx() bool`

HasAllStatus2xx returns a boolean if a field has been set.

### GetAllStatus3xx

`func (o *PlatformValues) GetAllStatus3xx() int64`

GetAllStatus3xx returns the AllStatus3xx field if non-nil, zero value otherwise.

### GetAllStatus3xxOk

`func (o *PlatformValues) GetAllStatus3xxOk() (*int64, bool)`

GetAllStatus3xxOk returns a tuple with the AllStatus3xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllStatus3xx

`func (o *PlatformValues) SetAllStatus3xx(v int64)`

SetAllStatus3xx sets AllStatus3xx field to given value.

### HasAllStatus3xx

`func (o *PlatformValues) HasAllStatus3xx() bool`

HasAllStatus3xx returns a boolean if a field has been set.

### GetAllStatus4xx

`func (o *PlatformValues) GetAllStatus4xx() int64`

GetAllStatus4xx returns the AllStatus4xx field if non-nil, zero value otherwise.

### GetAllStatus4xxOk

`func (o *PlatformValues) GetAllStatus4xxOk() (*int64, bool)`

GetAllStatus4xxOk returns a tuple with the AllStatus4xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllStatus4xx

`func (o *PlatformValues) SetAllStatus4xx(v int64)`

SetAllStatus4xx sets AllStatus4xx field to given value.

### HasAllStatus4xx

`func (o *PlatformValues) HasAllStatus4xx() bool`

HasAllStatus4xx returns a boolean if a field has been set.

### GetAllStatus5xx

`func (o *PlatformValues) GetAllStatus5xx() int64`

GetAllStatus5xx returns the AllStatus5xx field if non-nil, zero value otherwise.

### GetAllStatus5xxOk

`func (o *PlatformValues) GetAllStatus5xxOk() (*int64, bool)`

GetAllStatus5xxOk returns a tuple with the AllStatus5xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllStatus5xx

`func (o *PlatformValues) SetAllStatus5xx(v int64)`

SetAllStatus5xx sets AllStatus5xx field to given value.

### HasAllStatus5xx

`func (o *PlatformValues) HasAllStatus5xx() bool`

HasAllStatus5xx returns a boolean if a field has been set.

### GetOriginOffload

`func (o *PlatformValues) GetOriginOffload() float32`

GetOriginOffload returns the OriginOffload field if non-nil, zero value otherwise.

### GetOriginOffloadOk

`func (o *PlatformValues) GetOriginOffloadOk() (*float32, bool)`

GetOriginOffloadOk returns a tuple with the OriginOffload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginOffload

`func (o *PlatformValues) SetOriginOffload(v float32)`

SetOriginOffload sets OriginOffload field to given value.

### HasOriginOffload

`func (o *PlatformValues) HasOriginOffload() bool`

HasOriginOffload returns a boolean if a field has been set.

### GetRequestDeniedGetHeadBody

`func (o *PlatformValues) GetRequestDeniedGetHeadBody() int64`

GetRequestDeniedGetHeadBody returns the RequestDeniedGetHeadBody field if non-nil, zero value otherwise.

### GetRequestDeniedGetHeadBodyOk

`func (o *PlatformValues) GetRequestDeniedGetHeadBodyOk() (*int64, bool)`

GetRequestDeniedGetHeadBodyOk returns a tuple with the RequestDeniedGetHeadBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestDeniedGetHeadBody

`func (o *PlatformValues) SetRequestDeniedGetHeadBody(v int64)`

SetRequestDeniedGetHeadBody sets RequestDeniedGetHeadBody field to given value.

### HasRequestDeniedGetHeadBody

`func (o *PlatformValues) HasRequestDeniedGetHeadBody() bool`

HasRequestDeniedGetHeadBody returns a boolean if a field has been set.

### GetDdosProtectionRequestsDetectCount

`func (o *PlatformValues) GetDdosProtectionRequestsDetectCount() int64`

GetDdosProtectionRequestsDetectCount returns the DdosProtectionRequestsDetectCount field if non-nil, zero value otherwise.

### GetDdosProtectionRequestsDetectCountOk

`func (o *PlatformValues) GetDdosProtectionRequestsDetectCountOk() (*int64, bool)`

GetDdosProtectionRequestsDetectCountOk returns a tuple with the DdosProtectionRequestsDetectCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdosProtectionRequestsDetectCount

`func (o *PlatformValues) SetDdosProtectionRequestsDetectCount(v int64)`

SetDdosProtectionRequestsDetectCount sets DdosProtectionRequestsDetectCount field to given value.

### HasDdosProtectionRequestsDetectCount

`func (o *PlatformValues) HasDdosProtectionRequestsDetectCount() bool`

HasDdosProtectionRequestsDetectCount returns a boolean if a field has been set.

### GetDdosProtectionRequestsMitigateCount

`func (o *PlatformValues) GetDdosProtectionRequestsMitigateCount() int64`

GetDdosProtectionRequestsMitigateCount returns the DdosProtectionRequestsMitigateCount field if non-nil, zero value otherwise.

### GetDdosProtectionRequestsMitigateCountOk

`func (o *PlatformValues) GetDdosProtectionRequestsMitigateCountOk() (*int64, bool)`

GetDdosProtectionRequestsMitigateCountOk returns a tuple with the DdosProtectionRequestsMitigateCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdosProtectionRequestsMitigateCount

`func (o *PlatformValues) SetDdosProtectionRequestsMitigateCount(v int64)`

SetDdosProtectionRequestsMitigateCount sets DdosProtectionRequestsMitigateCount field to given value.

### HasDdosProtectionRequestsMitigateCount

`func (o *PlatformValues) HasDdosProtectionRequestsMitigateCount() bool`

HasDdosProtectionRequestsMitigateCount returns a boolean if a field has been set.

### GetDdosProtectionRequestsAllowCount

`func (o *PlatformValues) GetDdosProtectionRequestsAllowCount() int64`

GetDdosProtectionRequestsAllowCount returns the DdosProtectionRequestsAllowCount field if non-nil, zero value otherwise.

### GetDdosProtectionRequestsAllowCountOk

`func (o *PlatformValues) GetDdosProtectionRequestsAllowCountOk() (*int64, bool)`

GetDdosProtectionRequestsAllowCountOk returns a tuple with the DdosProtectionRequestsAllowCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdosProtectionRequestsAllowCount

`func (o *PlatformValues) SetDdosProtectionRequestsAllowCount(v int64)`

SetDdosProtectionRequestsAllowCount sets DdosProtectionRequestsAllowCount field to given value.

### HasDdosProtectionRequestsAllowCount

`func (o *PlatformValues) HasDdosProtectionRequestsAllowCount() bool`

HasDdosProtectionRequestsAllowCount returns a boolean if a field has been set.

### GetObjectStorageClassAOperationsCount

`func (o *PlatformValues) GetObjectStorageClassAOperationsCount() int64`

GetObjectStorageClassAOperationsCount returns the ObjectStorageClassAOperationsCount field if non-nil, zero value otherwise.

### GetObjectStorageClassAOperationsCountOk

`func (o *PlatformValues) GetObjectStorageClassAOperationsCountOk() (*int64, bool)`

GetObjectStorageClassAOperationsCountOk returns a tuple with the ObjectStorageClassAOperationsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectStorageClassAOperationsCount

`func (o *PlatformValues) SetObjectStorageClassAOperationsCount(v int64)`

SetObjectStorageClassAOperationsCount sets ObjectStorageClassAOperationsCount field to given value.

### HasObjectStorageClassAOperationsCount

`func (o *PlatformValues) HasObjectStorageClassAOperationsCount() bool`

HasObjectStorageClassAOperationsCount returns a boolean if a field has been set.

### GetObjectStorageClassBOperationsCount

`func (o *PlatformValues) GetObjectStorageClassBOperationsCount() int64`

GetObjectStorageClassBOperationsCount returns the ObjectStorageClassBOperationsCount field if non-nil, zero value otherwise.

### GetObjectStorageClassBOperationsCountOk

`func (o *PlatformValues) GetObjectStorageClassBOperationsCountOk() (*int64, bool)`

GetObjectStorageClassBOperationsCountOk returns a tuple with the ObjectStorageClassBOperationsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectStorageClassBOperationsCount

`func (o *PlatformValues) SetObjectStorageClassBOperationsCount(v int64)`

SetObjectStorageClassBOperationsCount sets ObjectStorageClassBOperationsCount field to given value.

### HasObjectStorageClassBOperationsCount

`func (o *PlatformValues) HasObjectStorageClassBOperationsCount() bool`

HasObjectStorageClassBOperationsCount returns a boolean if a field has been set.

### GetAiaRequests

`func (o *PlatformValues) GetAiaRequests() int64`

GetAiaRequests returns the AiaRequests field if non-nil, zero value otherwise.

### GetAiaRequestsOk

`func (o *PlatformValues) GetAiaRequestsOk() (*int64, bool)`

GetAiaRequestsOk returns a tuple with the AiaRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiaRequests

`func (o *PlatformValues) SetAiaRequests(v int64)`

SetAiaRequests sets AiaRequests field to given value.

### HasAiaRequests

`func (o *PlatformValues) HasAiaRequests() bool`

HasAiaRequests returns a boolean if a field has been set.

### GetAiaStatus1xx

`func (o *PlatformValues) GetAiaStatus1xx() int64`

GetAiaStatus1xx returns the AiaStatus1xx field if non-nil, zero value otherwise.

### GetAiaStatus1xxOk

`func (o *PlatformValues) GetAiaStatus1xxOk() (*int64, bool)`

GetAiaStatus1xxOk returns a tuple with the AiaStatus1xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiaStatus1xx

`func (o *PlatformValues) SetAiaStatus1xx(v int64)`

SetAiaStatus1xx sets AiaStatus1xx field to given value.

### HasAiaStatus1xx

`func (o *PlatformValues) HasAiaStatus1xx() bool`

HasAiaStatus1xx returns a boolean if a field has been set.

### GetAiaStatus2xx

`func (o *PlatformValues) GetAiaStatus2xx() int64`

GetAiaStatus2xx returns the AiaStatus2xx field if non-nil, zero value otherwise.

### GetAiaStatus2xxOk

`func (o *PlatformValues) GetAiaStatus2xxOk() (*int64, bool)`

GetAiaStatus2xxOk returns a tuple with the AiaStatus2xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiaStatus2xx

`func (o *PlatformValues) SetAiaStatus2xx(v int64)`

SetAiaStatus2xx sets AiaStatus2xx field to given value.

### HasAiaStatus2xx

`func (o *PlatformValues) HasAiaStatus2xx() bool`

HasAiaStatus2xx returns a boolean if a field has been set.

### GetAiaStatus3xx

`func (o *PlatformValues) GetAiaStatus3xx() int64`

GetAiaStatus3xx returns the AiaStatus3xx field if non-nil, zero value otherwise.

### GetAiaStatus3xxOk

`func (o *PlatformValues) GetAiaStatus3xxOk() (*int64, bool)`

GetAiaStatus3xxOk returns a tuple with the AiaStatus3xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiaStatus3xx

`func (o *PlatformValues) SetAiaStatus3xx(v int64)`

SetAiaStatus3xx sets AiaStatus3xx field to given value.

### HasAiaStatus3xx

`func (o *PlatformValues) HasAiaStatus3xx() bool`

HasAiaStatus3xx returns a boolean if a field has been set.

### GetAiaStatus4xx

`func (o *PlatformValues) GetAiaStatus4xx() int64`

GetAiaStatus4xx returns the AiaStatus4xx field if non-nil, zero value otherwise.

### GetAiaStatus4xxOk

`func (o *PlatformValues) GetAiaStatus4xxOk() (*int64, bool)`

GetAiaStatus4xxOk returns a tuple with the AiaStatus4xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiaStatus4xx

`func (o *PlatformValues) SetAiaStatus4xx(v int64)`

SetAiaStatus4xx sets AiaStatus4xx field to given value.

### HasAiaStatus4xx

`func (o *PlatformValues) HasAiaStatus4xx() bool`

HasAiaStatus4xx returns a boolean if a field has been set.

### GetAiaStatus5xx

`func (o *PlatformValues) GetAiaStatus5xx() int64`

GetAiaStatus5xx returns the AiaStatus5xx field if non-nil, zero value otherwise.

### GetAiaStatus5xxOk

`func (o *PlatformValues) GetAiaStatus5xxOk() (*int64, bool)`

GetAiaStatus5xxOk returns a tuple with the AiaStatus5xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiaStatus5xx

`func (o *PlatformValues) SetAiaStatus5xx(v int64)`

SetAiaStatus5xx sets AiaStatus5xx field to given value.

### HasAiaStatus5xx

`func (o *PlatformValues) HasAiaStatus5xx() bool`

HasAiaStatus5xx returns a boolean if a field has been set.

### GetAiaResponseUsageTokens

`func (o *PlatformValues) GetAiaResponseUsageTokens() int64`

GetAiaResponseUsageTokens returns the AiaResponseUsageTokens field if non-nil, zero value otherwise.

### GetAiaResponseUsageTokensOk

`func (o *PlatformValues) GetAiaResponseUsageTokensOk() (*int64, bool)`

GetAiaResponseUsageTokensOk returns a tuple with the AiaResponseUsageTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiaResponseUsageTokens

`func (o *PlatformValues) SetAiaResponseUsageTokens(v int64)`

SetAiaResponseUsageTokens sets AiaResponseUsageTokens field to given value.

### HasAiaResponseUsageTokens

`func (o *PlatformValues) HasAiaResponseUsageTokens() bool`

HasAiaResponseUsageTokens returns a boolean if a field has been set.

### GetAiaOriginUsageTokens

`func (o *PlatformValues) GetAiaOriginUsageTokens() int64`

GetAiaOriginUsageTokens returns the AiaOriginUsageTokens field if non-nil, zero value otherwise.

### GetAiaOriginUsageTokensOk

`func (o *PlatformValues) GetAiaOriginUsageTokensOk() (*int64, bool)`

GetAiaOriginUsageTokensOk returns a tuple with the AiaOriginUsageTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiaOriginUsageTokens

`func (o *PlatformValues) SetAiaOriginUsageTokens(v int64)`

SetAiaOriginUsageTokens sets AiaOriginUsageTokens field to given value.

### HasAiaOriginUsageTokens

`func (o *PlatformValues) HasAiaOriginUsageTokens() bool`

HasAiaOriginUsageTokens returns a boolean if a field has been set.

### GetAiaEstimatedTimeSavedMs

`func (o *PlatformValues) GetAiaEstimatedTimeSavedMs() int64`

GetAiaEstimatedTimeSavedMs returns the AiaEstimatedTimeSavedMs field if non-nil, zero value otherwise.

### GetAiaEstimatedTimeSavedMsOk

`func (o *PlatformValues) GetAiaEstimatedTimeSavedMsOk() (*int64, bool)`

GetAiaEstimatedTimeSavedMsOk returns a tuple with the AiaEstimatedTimeSavedMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiaEstimatedTimeSavedMs

`func (o *PlatformValues) SetAiaEstimatedTimeSavedMs(v int64)`

SetAiaEstimatedTimeSavedMs sets AiaEstimatedTimeSavedMs field to given value.

### HasAiaEstimatedTimeSavedMs

`func (o *PlatformValues) HasAiaEstimatedTimeSavedMs() bool`

HasAiaEstimatedTimeSavedMs returns a boolean if a field has been set.

### GetRequestCollapseUsableCount

`func (o *PlatformValues) GetRequestCollapseUsableCount() int64`

GetRequestCollapseUsableCount returns the RequestCollapseUsableCount field if non-nil, zero value otherwise.

### GetRequestCollapseUsableCountOk

`func (o *PlatformValues) GetRequestCollapseUsableCountOk() (*int64, bool)`

GetRequestCollapseUsableCountOk returns a tuple with the RequestCollapseUsableCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCollapseUsableCount

`func (o *PlatformValues) SetRequestCollapseUsableCount(v int64)`

SetRequestCollapseUsableCount sets RequestCollapseUsableCount field to given value.

### HasRequestCollapseUsableCount

`func (o *PlatformValues) HasRequestCollapseUsableCount() bool`

HasRequestCollapseUsableCount returns a boolean if a field has been set.

### GetRequestCollapseUnusableCount

`func (o *PlatformValues) GetRequestCollapseUnusableCount() int64`

GetRequestCollapseUnusableCount returns the RequestCollapseUnusableCount field if non-nil, zero value otherwise.

### GetRequestCollapseUnusableCountOk

`func (o *PlatformValues) GetRequestCollapseUnusableCountOk() (*int64, bool)`

GetRequestCollapseUnusableCountOk returns a tuple with the RequestCollapseUnusableCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCollapseUnusableCount

`func (o *PlatformValues) SetRequestCollapseUnusableCount(v int64)`

SetRequestCollapseUnusableCount sets RequestCollapseUnusableCount field to given value.

### HasRequestCollapseUnusableCount

`func (o *PlatformValues) HasRequestCollapseUnusableCount() bool`

HasRequestCollapseUnusableCount returns a boolean if a field has been set.

### GetComputeCacheOperationsCount

`func (o *PlatformValues) GetComputeCacheOperationsCount() int64`

GetComputeCacheOperationsCount returns the ComputeCacheOperationsCount field if non-nil, zero value otherwise.

### GetComputeCacheOperationsCountOk

`func (o *PlatformValues) GetComputeCacheOperationsCountOk() (*int64, bool)`

GetComputeCacheOperationsCountOk returns a tuple with the ComputeCacheOperationsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeCacheOperationsCount

`func (o *PlatformValues) SetComputeCacheOperationsCount(v int64)`

SetComputeCacheOperationsCount sets ComputeCacheOperationsCount field to given value.

### HasComputeCacheOperationsCount

`func (o *PlatformValues) HasComputeCacheOperationsCount() bool`

HasComputeCacheOperationsCount returns a boolean if a field has been set.

### GetNgwafRequestsTotalCount

`func (o *PlatformValues) GetNgwafRequestsTotalCount() int32`

GetNgwafRequestsTotalCount returns the NgwafRequestsTotalCount field if non-nil, zero value otherwise.

### GetNgwafRequestsTotalCountOk

`func (o *PlatformValues) GetNgwafRequestsTotalCountOk() (*int32, bool)`

GetNgwafRequestsTotalCountOk returns a tuple with the NgwafRequestsTotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNgwafRequestsTotalCount

`func (o *PlatformValues) SetNgwafRequestsTotalCount(v int32)`

SetNgwafRequestsTotalCount sets NgwafRequestsTotalCount field to given value.

### HasNgwafRequestsTotalCount

`func (o *PlatformValues) HasNgwafRequestsTotalCount() bool`

HasNgwafRequestsTotalCount returns a boolean if a field has been set.

### GetNgwafRequestsUnknownCount

`func (o *PlatformValues) GetNgwafRequestsUnknownCount() int32`

GetNgwafRequestsUnknownCount returns the NgwafRequestsUnknownCount field if non-nil, zero value otherwise.

### GetNgwafRequestsUnknownCountOk

`func (o *PlatformValues) GetNgwafRequestsUnknownCountOk() (*int32, bool)`

GetNgwafRequestsUnknownCountOk returns a tuple with the NgwafRequestsUnknownCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNgwafRequestsUnknownCount

`func (o *PlatformValues) SetNgwafRequestsUnknownCount(v int32)`

SetNgwafRequestsUnknownCount sets NgwafRequestsUnknownCount field to given value.

### HasNgwafRequestsUnknownCount

`func (o *PlatformValues) HasNgwafRequestsUnknownCount() bool`

HasNgwafRequestsUnknownCount returns a boolean if a field has been set.

### GetNgwafRequestsAllowedCount

`func (o *PlatformValues) GetNgwafRequestsAllowedCount() int32`

GetNgwafRequestsAllowedCount returns the NgwafRequestsAllowedCount field if non-nil, zero value otherwise.

### GetNgwafRequestsAllowedCountOk

`func (o *PlatformValues) GetNgwafRequestsAllowedCountOk() (*int32, bool)`

GetNgwafRequestsAllowedCountOk returns a tuple with the NgwafRequestsAllowedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNgwafRequestsAllowedCount

`func (o *PlatformValues) SetNgwafRequestsAllowedCount(v int32)`

SetNgwafRequestsAllowedCount sets NgwafRequestsAllowedCount field to given value.

### HasNgwafRequestsAllowedCount

`func (o *PlatformValues) HasNgwafRequestsAllowedCount() bool`

HasNgwafRequestsAllowedCount returns a boolean if a field has been set.

### GetNgwafRequestsLoggedCount

`func (o *PlatformValues) GetNgwafRequestsLoggedCount() int32`

GetNgwafRequestsLoggedCount returns the NgwafRequestsLoggedCount field if non-nil, zero value otherwise.

### GetNgwafRequestsLoggedCountOk

`func (o *PlatformValues) GetNgwafRequestsLoggedCountOk() (*int32, bool)`

GetNgwafRequestsLoggedCountOk returns a tuple with the NgwafRequestsLoggedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNgwafRequestsLoggedCount

`func (o *PlatformValues) SetNgwafRequestsLoggedCount(v int32)`

SetNgwafRequestsLoggedCount sets NgwafRequestsLoggedCount field to given value.

### HasNgwafRequestsLoggedCount

`func (o *PlatformValues) HasNgwafRequestsLoggedCount() bool`

HasNgwafRequestsLoggedCount returns a boolean if a field has been set.

### GetNgwafRequestsBlockedCount

`func (o *PlatformValues) GetNgwafRequestsBlockedCount() int32`

GetNgwafRequestsBlockedCount returns the NgwafRequestsBlockedCount field if non-nil, zero value otherwise.

### GetNgwafRequestsBlockedCountOk

`func (o *PlatformValues) GetNgwafRequestsBlockedCountOk() (*int32, bool)`

GetNgwafRequestsBlockedCountOk returns a tuple with the NgwafRequestsBlockedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNgwafRequestsBlockedCount

`func (o *PlatformValues) SetNgwafRequestsBlockedCount(v int32)`

SetNgwafRequestsBlockedCount sets NgwafRequestsBlockedCount field to given value.

### HasNgwafRequestsBlockedCount

`func (o *PlatformValues) HasNgwafRequestsBlockedCount() bool`

HasNgwafRequestsBlockedCount returns a boolean if a field has been set.

### GetNgwafRequestsTimeoutCount

`func (o *PlatformValues) GetNgwafRequestsTimeoutCount() int32`

GetNgwafRequestsTimeoutCount returns the NgwafRequestsTimeoutCount field if non-nil, zero value otherwise.

### GetNgwafRequestsTimeoutCountOk

`func (o *PlatformValues) GetNgwafRequestsTimeoutCountOk() (*int32, bool)`

GetNgwafRequestsTimeoutCountOk returns a tuple with the NgwafRequestsTimeoutCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNgwafRequestsTimeoutCount

`func (o *PlatformValues) SetNgwafRequestsTimeoutCount(v int32)`

SetNgwafRequestsTimeoutCount sets NgwafRequestsTimeoutCount field to given value.

### HasNgwafRequestsTimeoutCount

`func (o *PlatformValues) HasNgwafRequestsTimeoutCount() bool`

HasNgwafRequestsTimeoutCount returns a boolean if a field has been set.

### GetNgwafRequestsChallengedCount

`func (o *PlatformValues) GetNgwafRequestsChallengedCount() int32`

GetNgwafRequestsChallengedCount returns the NgwafRequestsChallengedCount field if non-nil, zero value otherwise.

### GetNgwafRequestsChallengedCountOk

`func (o *PlatformValues) GetNgwafRequestsChallengedCountOk() (*int32, bool)`

GetNgwafRequestsChallengedCountOk returns a tuple with the NgwafRequestsChallengedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNgwafRequestsChallengedCount

`func (o *PlatformValues) SetNgwafRequestsChallengedCount(v int32)`

SetNgwafRequestsChallengedCount sets NgwafRequestsChallengedCount field to given value.

### HasNgwafRequestsChallengedCount

`func (o *PlatformValues) HasNgwafRequestsChallengedCount() bool`

HasNgwafRequestsChallengedCount returns a boolean if a field has been set.

### GetApiDiscoveryRequestsCount

`func (o *PlatformValues) GetApiDiscoveryRequestsCount() int32`

GetApiDiscoveryRequestsCount returns the ApiDiscoveryRequestsCount field if non-nil, zero value otherwise.

### GetApiDiscoveryRequestsCountOk

`func (o *PlatformValues) GetApiDiscoveryRequestsCountOk() (*int32, bool)`

GetApiDiscoveryRequestsCountOk returns a tuple with the ApiDiscoveryRequestsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiDiscoveryRequestsCount

`func (o *PlatformValues) SetApiDiscoveryRequestsCount(v int32)`

SetApiDiscoveryRequestsCount sets ApiDiscoveryRequestsCount field to given value.

### HasApiDiscoveryRequestsCount

`func (o *PlatformValues) HasApiDiscoveryRequestsCount() bool`

HasApiDiscoveryRequestsCount returns a boolean if a field has been set.

### GetComputeRespStatus103

`func (o *PlatformValues) GetComputeRespStatus103() int32`

GetComputeRespStatus103 returns the ComputeRespStatus103 field if non-nil, zero value otherwise.

### GetComputeRespStatus103Ok

`func (o *PlatformValues) GetComputeRespStatus103Ok() (*int32, bool)`

GetComputeRespStatus103Ok returns a tuple with the ComputeRespStatus103 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespStatus103

`func (o *PlatformValues) SetComputeRespStatus103(v int32)`

SetComputeRespStatus103 sets ComputeRespStatus103 field to given value.

### HasComputeRespStatus103

`func (o *PlatformValues) HasComputeRespStatus103() bool`

HasComputeRespStatus103 returns a boolean if a field has been set.

### GetComputeRespStatus200

`func (o *PlatformValues) GetComputeRespStatus200() int32`

GetComputeRespStatus200 returns the ComputeRespStatus200 field if non-nil, zero value otherwise.

### GetComputeRespStatus200Ok

`func (o *PlatformValues) GetComputeRespStatus200Ok() (*int32, bool)`

GetComputeRespStatus200Ok returns a tuple with the ComputeRespStatus200 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespStatus200

`func (o *PlatformValues) SetComputeRespStatus200(v int32)`

SetComputeRespStatus200 sets ComputeRespStatus200 field to given value.

### HasComputeRespStatus200

`func (o *PlatformValues) HasComputeRespStatus200() bool`

HasComputeRespStatus200 returns a boolean if a field has been set.

### GetComputeRespStatus204

`func (o *PlatformValues) GetComputeRespStatus204() int32`

GetComputeRespStatus204 returns the ComputeRespStatus204 field if non-nil, zero value otherwise.

### GetComputeRespStatus204Ok

`func (o *PlatformValues) GetComputeRespStatus204Ok() (*int32, bool)`

GetComputeRespStatus204Ok returns a tuple with the ComputeRespStatus204 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespStatus204

`func (o *PlatformValues) SetComputeRespStatus204(v int32)`

SetComputeRespStatus204 sets ComputeRespStatus204 field to given value.

### HasComputeRespStatus204

`func (o *PlatformValues) HasComputeRespStatus204() bool`

HasComputeRespStatus204 returns a boolean if a field has been set.

### GetComputeRespStatus206

`func (o *PlatformValues) GetComputeRespStatus206() int32`

GetComputeRespStatus206 returns the ComputeRespStatus206 field if non-nil, zero value otherwise.

### GetComputeRespStatus206Ok

`func (o *PlatformValues) GetComputeRespStatus206Ok() (*int32, bool)`

GetComputeRespStatus206Ok returns a tuple with the ComputeRespStatus206 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespStatus206

`func (o *PlatformValues) SetComputeRespStatus206(v int32)`

SetComputeRespStatus206 sets ComputeRespStatus206 field to given value.

### HasComputeRespStatus206

`func (o *PlatformValues) HasComputeRespStatus206() bool`

HasComputeRespStatus206 returns a boolean if a field has been set.

### GetComputeRespStatus301

`func (o *PlatformValues) GetComputeRespStatus301() int32`

GetComputeRespStatus301 returns the ComputeRespStatus301 field if non-nil, zero value otherwise.

### GetComputeRespStatus301Ok

`func (o *PlatformValues) GetComputeRespStatus301Ok() (*int32, bool)`

GetComputeRespStatus301Ok returns a tuple with the ComputeRespStatus301 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespStatus301

`func (o *PlatformValues) SetComputeRespStatus301(v int32)`

SetComputeRespStatus301 sets ComputeRespStatus301 field to given value.

### HasComputeRespStatus301

`func (o *PlatformValues) HasComputeRespStatus301() bool`

HasComputeRespStatus301 returns a boolean if a field has been set.

### GetComputeRespStatus302

`func (o *PlatformValues) GetComputeRespStatus302() int32`

GetComputeRespStatus302 returns the ComputeRespStatus302 field if non-nil, zero value otherwise.

### GetComputeRespStatus302Ok

`func (o *PlatformValues) GetComputeRespStatus302Ok() (*int32, bool)`

GetComputeRespStatus302Ok returns a tuple with the ComputeRespStatus302 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespStatus302

`func (o *PlatformValues) SetComputeRespStatus302(v int32)`

SetComputeRespStatus302 sets ComputeRespStatus302 field to given value.

### HasComputeRespStatus302

`func (o *PlatformValues) HasComputeRespStatus302() bool`

HasComputeRespStatus302 returns a boolean if a field has been set.

### GetComputeRespStatus304

`func (o *PlatformValues) GetComputeRespStatus304() int32`

GetComputeRespStatus304 returns the ComputeRespStatus304 field if non-nil, zero value otherwise.

### GetComputeRespStatus304Ok

`func (o *PlatformValues) GetComputeRespStatus304Ok() (*int32, bool)`

GetComputeRespStatus304Ok returns a tuple with the ComputeRespStatus304 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespStatus304

`func (o *PlatformValues) SetComputeRespStatus304(v int32)`

SetComputeRespStatus304 sets ComputeRespStatus304 field to given value.

### HasComputeRespStatus304

`func (o *PlatformValues) HasComputeRespStatus304() bool`

HasComputeRespStatus304 returns a boolean if a field has been set.

### GetComputeRespStatus400

`func (o *PlatformValues) GetComputeRespStatus400() int32`

GetComputeRespStatus400 returns the ComputeRespStatus400 field if non-nil, zero value otherwise.

### GetComputeRespStatus400Ok

`func (o *PlatformValues) GetComputeRespStatus400Ok() (*int32, bool)`

GetComputeRespStatus400Ok returns a tuple with the ComputeRespStatus400 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespStatus400

`func (o *PlatformValues) SetComputeRespStatus400(v int32)`

SetComputeRespStatus400 sets ComputeRespStatus400 field to given value.

### HasComputeRespStatus400

`func (o *PlatformValues) HasComputeRespStatus400() bool`

HasComputeRespStatus400 returns a boolean if a field has been set.

### GetComputeRespStatus401

`func (o *PlatformValues) GetComputeRespStatus401() int32`

GetComputeRespStatus401 returns the ComputeRespStatus401 field if non-nil, zero value otherwise.

### GetComputeRespStatus401Ok

`func (o *PlatformValues) GetComputeRespStatus401Ok() (*int32, bool)`

GetComputeRespStatus401Ok returns a tuple with the ComputeRespStatus401 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespStatus401

`func (o *PlatformValues) SetComputeRespStatus401(v int32)`

SetComputeRespStatus401 sets ComputeRespStatus401 field to given value.

### HasComputeRespStatus401

`func (o *PlatformValues) HasComputeRespStatus401() bool`

HasComputeRespStatus401 returns a boolean if a field has been set.

### GetComputeRespStatus403

`func (o *PlatformValues) GetComputeRespStatus403() int32`

GetComputeRespStatus403 returns the ComputeRespStatus403 field if non-nil, zero value otherwise.

### GetComputeRespStatus403Ok

`func (o *PlatformValues) GetComputeRespStatus403Ok() (*int32, bool)`

GetComputeRespStatus403Ok returns a tuple with the ComputeRespStatus403 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespStatus403

`func (o *PlatformValues) SetComputeRespStatus403(v int32)`

SetComputeRespStatus403 sets ComputeRespStatus403 field to given value.

### HasComputeRespStatus403

`func (o *PlatformValues) HasComputeRespStatus403() bool`

HasComputeRespStatus403 returns a boolean if a field has been set.

### GetComputeRespStatus404

`func (o *PlatformValues) GetComputeRespStatus404() int32`

GetComputeRespStatus404 returns the ComputeRespStatus404 field if non-nil, zero value otherwise.

### GetComputeRespStatus404Ok

`func (o *PlatformValues) GetComputeRespStatus404Ok() (*int32, bool)`

GetComputeRespStatus404Ok returns a tuple with the ComputeRespStatus404 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespStatus404

`func (o *PlatformValues) SetComputeRespStatus404(v int32)`

SetComputeRespStatus404 sets ComputeRespStatus404 field to given value.

### HasComputeRespStatus404

`func (o *PlatformValues) HasComputeRespStatus404() bool`

HasComputeRespStatus404 returns a boolean if a field has been set.

### GetComputeRespStatus416

`func (o *PlatformValues) GetComputeRespStatus416() int32`

GetComputeRespStatus416 returns the ComputeRespStatus416 field if non-nil, zero value otherwise.

### GetComputeRespStatus416Ok

`func (o *PlatformValues) GetComputeRespStatus416Ok() (*int32, bool)`

GetComputeRespStatus416Ok returns a tuple with the ComputeRespStatus416 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespStatus416

`func (o *PlatformValues) SetComputeRespStatus416(v int32)`

SetComputeRespStatus416 sets ComputeRespStatus416 field to given value.

### HasComputeRespStatus416

`func (o *PlatformValues) HasComputeRespStatus416() bool`

HasComputeRespStatus416 returns a boolean if a field has been set.

### GetComputeRespStatus429

`func (o *PlatformValues) GetComputeRespStatus429() int32`

GetComputeRespStatus429 returns the ComputeRespStatus429 field if non-nil, zero value otherwise.

### GetComputeRespStatus429Ok

`func (o *PlatformValues) GetComputeRespStatus429Ok() (*int32, bool)`

GetComputeRespStatus429Ok returns a tuple with the ComputeRespStatus429 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespStatus429

`func (o *PlatformValues) SetComputeRespStatus429(v int32)`

SetComputeRespStatus429 sets ComputeRespStatus429 field to given value.

### HasComputeRespStatus429

`func (o *PlatformValues) HasComputeRespStatus429() bool`

HasComputeRespStatus429 returns a boolean if a field has been set.

### GetComputeRespStatus500

`func (o *PlatformValues) GetComputeRespStatus500() int32`

GetComputeRespStatus500 returns the ComputeRespStatus500 field if non-nil, zero value otherwise.

### GetComputeRespStatus500Ok

`func (o *PlatformValues) GetComputeRespStatus500Ok() (*int32, bool)`

GetComputeRespStatus500Ok returns a tuple with the ComputeRespStatus500 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespStatus500

`func (o *PlatformValues) SetComputeRespStatus500(v int32)`

SetComputeRespStatus500 sets ComputeRespStatus500 field to given value.

### HasComputeRespStatus500

`func (o *PlatformValues) HasComputeRespStatus500() bool`

HasComputeRespStatus500 returns a boolean if a field has been set.

### GetComputeRespStatus501

`func (o *PlatformValues) GetComputeRespStatus501() int32`

GetComputeRespStatus501 returns the ComputeRespStatus501 field if non-nil, zero value otherwise.

### GetComputeRespStatus501Ok

`func (o *PlatformValues) GetComputeRespStatus501Ok() (*int32, bool)`

GetComputeRespStatus501Ok returns a tuple with the ComputeRespStatus501 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespStatus501

`func (o *PlatformValues) SetComputeRespStatus501(v int32)`

SetComputeRespStatus501 sets ComputeRespStatus501 field to given value.

### HasComputeRespStatus501

`func (o *PlatformValues) HasComputeRespStatus501() bool`

HasComputeRespStatus501 returns a boolean if a field has been set.

### GetComputeRespStatus502

`func (o *PlatformValues) GetComputeRespStatus502() int32`

GetComputeRespStatus502 returns the ComputeRespStatus502 field if non-nil, zero value otherwise.

### GetComputeRespStatus502Ok

`func (o *PlatformValues) GetComputeRespStatus502Ok() (*int32, bool)`

GetComputeRespStatus502Ok returns a tuple with the ComputeRespStatus502 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespStatus502

`func (o *PlatformValues) SetComputeRespStatus502(v int32)`

SetComputeRespStatus502 sets ComputeRespStatus502 field to given value.

### HasComputeRespStatus502

`func (o *PlatformValues) HasComputeRespStatus502() bool`

HasComputeRespStatus502 returns a boolean if a field has been set.

### GetComputeRespStatus503

`func (o *PlatformValues) GetComputeRespStatus503() int32`

GetComputeRespStatus503 returns the ComputeRespStatus503 field if non-nil, zero value otherwise.

### GetComputeRespStatus503Ok

`func (o *PlatformValues) GetComputeRespStatus503Ok() (*int32, bool)`

GetComputeRespStatus503Ok returns a tuple with the ComputeRespStatus503 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespStatus503

`func (o *PlatformValues) SetComputeRespStatus503(v int32)`

SetComputeRespStatus503 sets ComputeRespStatus503 field to given value.

### HasComputeRespStatus503

`func (o *PlatformValues) HasComputeRespStatus503() bool`

HasComputeRespStatus503 returns a boolean if a field has been set.

### GetComputeRespStatus504

`func (o *PlatformValues) GetComputeRespStatus504() int32`

GetComputeRespStatus504 returns the ComputeRespStatus504 field if non-nil, zero value otherwise.

### GetComputeRespStatus504Ok

`func (o *PlatformValues) GetComputeRespStatus504Ok() (*int32, bool)`

GetComputeRespStatus504Ok returns a tuple with the ComputeRespStatus504 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespStatus504

`func (o *PlatformValues) SetComputeRespStatus504(v int32)`

SetComputeRespStatus504 sets ComputeRespStatus504 field to given value.

### HasComputeRespStatus504

`func (o *PlatformValues) HasComputeRespStatus504() bool`

HasComputeRespStatus504 returns a boolean if a field has been set.

### GetComputeRespStatus505

`func (o *PlatformValues) GetComputeRespStatus505() int32`

GetComputeRespStatus505 returns the ComputeRespStatus505 field if non-nil, zero value otherwise.

### GetComputeRespStatus505Ok

`func (o *PlatformValues) GetComputeRespStatus505Ok() (*int32, bool)`

GetComputeRespStatus505Ok returns a tuple with the ComputeRespStatus505 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespStatus505

`func (o *PlatformValues) SetComputeRespStatus505(v int32)`

SetComputeRespStatus505 sets ComputeRespStatus505 field to given value.

### HasComputeRespStatus505

`func (o *PlatformValues) HasComputeRespStatus505() bool`

HasComputeRespStatus505 returns a boolean if a field has been set.

### GetComputeRespStatus530

`func (o *PlatformValues) GetComputeRespStatus530() int32`

GetComputeRespStatus530 returns the ComputeRespStatus530 field if non-nil, zero value otherwise.

### GetComputeRespStatus530Ok

`func (o *PlatformValues) GetComputeRespStatus530Ok() (*int32, bool)`

GetComputeRespStatus530Ok returns a tuple with the ComputeRespStatus530 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRespStatus530

`func (o *PlatformValues) SetComputeRespStatus530(v int32)`

SetComputeRespStatus530 sets ComputeRespStatus530 field to given value.

### HasComputeRespStatus530

`func (o *PlatformValues) HasComputeRespStatus530() bool`

HasComputeRespStatus530 returns a boolean if a field has been set.

### GetImgoptoComputeRequests

`func (o *PlatformValues) GetImgoptoComputeRequests() int32`

GetImgoptoComputeRequests returns the ImgoptoComputeRequests field if non-nil, zero value otherwise.

### GetImgoptoComputeRequestsOk

`func (o *PlatformValues) GetImgoptoComputeRequestsOk() (*int32, bool)`

GetImgoptoComputeRequestsOk returns a tuple with the ImgoptoComputeRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImgoptoComputeRequests

`func (o *PlatformValues) SetImgoptoComputeRequests(v int32)`

SetImgoptoComputeRequests sets ImgoptoComputeRequests field to given value.

### HasImgoptoComputeRequests

`func (o *PlatformValues) HasImgoptoComputeRequests() bool`

HasImgoptoComputeRequests returns a boolean if a field has been set.

### GetDNSBillableResponsesCount

`func (o *PlatformValues) GetDNSBillableResponsesCount() int32`

GetDNSBillableResponsesCount returns the DNSBillableResponsesCount field if non-nil, zero value otherwise.

### GetDNSBillableResponsesCountOk

`func (o *PlatformValues) GetDNSBillableResponsesCountOk() (*int32, bool)`

GetDNSBillableResponsesCountOk returns a tuple with the DNSBillableResponsesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDNSBillableResponsesCount

`func (o *PlatformValues) SetDNSBillableResponsesCount(v int32)`

SetDNSBillableResponsesCount sets DNSBillableResponsesCount field to given value.

### HasDNSBillableResponsesCount

`func (o *PlatformValues) HasDNSBillableResponsesCount() bool`

HasDNSBillableResponsesCount returns a boolean if a field has been set.

### GetDNSNonbillableResponsesCount

`func (o *PlatformValues) GetDNSNonbillableResponsesCount() int32`

GetDNSNonbillableResponsesCount returns the DNSNonbillableResponsesCount field if non-nil, zero value otherwise.

### GetDNSNonbillableResponsesCountOk

`func (o *PlatformValues) GetDNSNonbillableResponsesCountOk() (*int32, bool)`

GetDNSNonbillableResponsesCountOk returns a tuple with the DNSNonbillableResponsesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDNSNonbillableResponsesCount

`func (o *PlatformValues) SetDNSNonbillableResponsesCount(v int32)`

SetDNSNonbillableResponsesCount sets DNSNonbillableResponsesCount field to given value.

### HasDNSNonbillableResponsesCount

`func (o *PlatformValues) HasDNSNonbillableResponsesCount() bool`

HasDNSNonbillableResponsesCount returns a boolean if a field has been set.

### GetUpgrade

`func (o *PlatformValues) GetUpgrade() int32`

GetUpgrade returns the Upgrade field if non-nil, zero value otherwise.

### GetUpgradeOk

`func (o *PlatformValues) GetUpgradeOk() (*int32, bool)`

GetUpgradeOk returns a tuple with the Upgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgrade

`func (o *PlatformValues) SetUpgrade(v int32)`

SetUpgrade sets Upgrade field to given value.

### HasUpgrade

`func (o *PlatformValues) HasUpgrade() bool`

HasUpgrade returns a boolean if a field has been set.

### GetNgwafBotAnalysisRequestCount

`func (o *PlatformValues) GetNgwafBotAnalysisRequestCount() int32`

GetNgwafBotAnalysisRequestCount returns the NgwafBotAnalysisRequestCount field if non-nil, zero value otherwise.

### GetNgwafBotAnalysisRequestCountOk

`func (o *PlatformValues) GetNgwafBotAnalysisRequestCountOk() (*int32, bool)`

GetNgwafBotAnalysisRequestCountOk returns a tuple with the NgwafBotAnalysisRequestCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNgwafBotAnalysisRequestCount

`func (o *PlatformValues) SetNgwafBotAnalysisRequestCount(v int32)`

SetNgwafBotAnalysisRequestCount sets NgwafBotAnalysisRequestCount field to given value.

### HasNgwafBotAnalysisRequestCount

`func (o *PlatformValues) HasNgwafBotAnalysisRequestCount() bool`

HasNgwafBotAnalysisRequestCount returns a boolean if a field has been set.

### GetImgoptoAvifCount

`func (o *PlatformValues) GetImgoptoAvifCount() int32`

GetImgoptoAvifCount returns the ImgoptoAvifCount field if non-nil, zero value otherwise.

### GetImgoptoAvifCountOk

`func (o *PlatformValues) GetImgoptoAvifCountOk() (*int32, bool)`

GetImgoptoAvifCountOk returns a tuple with the ImgoptoAvifCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImgoptoAvifCount

`func (o *PlatformValues) SetImgoptoAvifCount(v int32)`

SetImgoptoAvifCount sets ImgoptoAvifCount field to given value.

### HasImgoptoAvifCount

`func (o *PlatformValues) HasImgoptoAvifCount() bool`

HasImgoptoAvifCount returns a boolean if a field has been set.

### GetImgoptoJpegCount

`func (o *PlatformValues) GetImgoptoJpegCount() int32`

GetImgoptoJpegCount returns the ImgoptoJpegCount field if non-nil, zero value otherwise.

### GetImgoptoJpegCountOk

`func (o *PlatformValues) GetImgoptoJpegCountOk() (*int32, bool)`

GetImgoptoJpegCountOk returns a tuple with the ImgoptoJpegCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImgoptoJpegCount

`func (o *PlatformValues) SetImgoptoJpegCount(v int32)`

SetImgoptoJpegCount sets ImgoptoJpegCount field to given value.

### HasImgoptoJpegCount

`func (o *PlatformValues) HasImgoptoJpegCount() bool`

HasImgoptoJpegCount returns a boolean if a field has been set.

### GetImgoptoPngCount

`func (o *PlatformValues) GetImgoptoPngCount() int32`

GetImgoptoPngCount returns the ImgoptoPngCount field if non-nil, zero value otherwise.

### GetImgoptoPngCountOk

`func (o *PlatformValues) GetImgoptoPngCountOk() (*int32, bool)`

GetImgoptoPngCountOk returns a tuple with the ImgoptoPngCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImgoptoPngCount

`func (o *PlatformValues) SetImgoptoPngCount(v int32)`

SetImgoptoPngCount sets ImgoptoPngCount field to given value.

### HasImgoptoPngCount

`func (o *PlatformValues) HasImgoptoPngCount() bool`

HasImgoptoPngCount returns a boolean if a field has been set.

### GetImgoptoGifCount

`func (o *PlatformValues) GetImgoptoGifCount() int32`

GetImgoptoGifCount returns the ImgoptoGifCount field if non-nil, zero value otherwise.

### GetImgoptoGifCountOk

`func (o *PlatformValues) GetImgoptoGifCountOk() (*int32, bool)`

GetImgoptoGifCountOk returns a tuple with the ImgoptoGifCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImgoptoGifCount

`func (o *PlatformValues) SetImgoptoGifCount(v int32)`

SetImgoptoGifCount sets ImgoptoGifCount field to given value.

### HasImgoptoGifCount

`func (o *PlatformValues) HasImgoptoGifCount() bool`

HasImgoptoGifCount returns a boolean if a field has been set.

### GetImgoptoWebpCount

`func (o *PlatformValues) GetImgoptoWebpCount() int32`

GetImgoptoWebpCount returns the ImgoptoWebpCount field if non-nil, zero value otherwise.

### GetImgoptoWebpCountOk

`func (o *PlatformValues) GetImgoptoWebpCountOk() (*int32, bool)`

GetImgoptoWebpCountOk returns a tuple with the ImgoptoWebpCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImgoptoWebpCount

`func (o *PlatformValues) SetImgoptoWebpCount(v int32)`

SetImgoptoWebpCount sets ImgoptoWebpCount field to given value.

### HasImgoptoWebpCount

`func (o *PlatformValues) HasImgoptoWebpCount() bool`

HasImgoptoWebpCount returns a boolean if a field has been set.

### GetImgoptoJpegxlCount

`func (o *PlatformValues) GetImgoptoJpegxlCount() int32`

GetImgoptoJpegxlCount returns the ImgoptoJpegxlCount field if non-nil, zero value otherwise.

### GetImgoptoJpegxlCountOk

`func (o *PlatformValues) GetImgoptoJpegxlCountOk() (*int32, bool)`

GetImgoptoJpegxlCountOk returns a tuple with the ImgoptoJpegxlCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImgoptoJpegxlCount

`func (o *PlatformValues) SetImgoptoJpegxlCount(v int32)`

SetImgoptoJpegxlCount sets ImgoptoJpegxlCount field to given value.

### HasImgoptoJpegxlCount

`func (o *PlatformValues) HasImgoptoJpegxlCount() bool`

HasImgoptoJpegxlCount returns a boolean if a field has been set.

### GetImgoptoSvgCount

`func (o *PlatformValues) GetImgoptoSvgCount() int32`

GetImgoptoSvgCount returns the ImgoptoSvgCount field if non-nil, zero value otherwise.

### GetImgoptoSvgCountOk

`func (o *PlatformValues) GetImgoptoSvgCountOk() (*int32, bool)`

GetImgoptoSvgCountOk returns a tuple with the ImgoptoSvgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImgoptoSvgCount

`func (o *PlatformValues) SetImgoptoSvgCount(v int32)`

SetImgoptoSvgCount sets ImgoptoSvgCount field to given value.

### HasImgoptoSvgCount

`func (o *PlatformValues) HasImgoptoSvgCount() bool`

HasImgoptoSvgCount returns a boolean if a field has been set.

### GetImgoptoMp4Count

`func (o *PlatformValues) GetImgoptoMp4Count() int32`

GetImgoptoMp4Count returns the ImgoptoMp4Count field if non-nil, zero value otherwise.

### GetImgoptoMp4CountOk

`func (o *PlatformValues) GetImgoptoMp4CountOk() (*int32, bool)`

GetImgoptoMp4CountOk returns a tuple with the ImgoptoMp4Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImgoptoMp4Count

`func (o *PlatformValues) SetImgoptoMp4Count(v int32)`

SetImgoptoMp4Count sets ImgoptoMp4Count field to given value.

### HasImgoptoMp4Count

`func (o *PlatformValues) HasImgoptoMp4Count() bool`

HasImgoptoMp4Count returns a boolean if a field has been set.

### GetComputeServiceResourceLimitsError

`func (o *PlatformValues) GetComputeServiceResourceLimitsError() int32`

GetComputeServiceResourceLimitsError returns the ComputeServiceResourceLimitsError field if non-nil, zero value otherwise.

### GetComputeServiceResourceLimitsErrorOk

`func (o *PlatformValues) GetComputeServiceResourceLimitsErrorOk() (*int32, bool)`

GetComputeServiceResourceLimitsErrorOk returns a tuple with the ComputeServiceResourceLimitsError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeServiceResourceLimitsError

`func (o *PlatformValues) SetComputeServiceResourceLimitsError(v int32)`

SetComputeServiceResourceLimitsError sets ComputeServiceResourceLimitsError field to given value.

### HasComputeServiceResourceLimitsError

`func (o *PlatformValues) HasComputeServiceResourceLimitsError() bool`

HasComputeServiceResourceLimitsError returns a boolean if a field has been set.

### GetComputeServiceRuntimeError

`func (o *PlatformValues) GetComputeServiceRuntimeError() int32`

GetComputeServiceRuntimeError returns the ComputeServiceRuntimeError field if non-nil, zero value otherwise.

### GetComputeServiceRuntimeErrorOk

`func (o *PlatformValues) GetComputeServiceRuntimeErrorOk() (*int32, bool)`

GetComputeServiceRuntimeErrorOk returns a tuple with the ComputeServiceRuntimeError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeServiceRuntimeError

`func (o *PlatformValues) SetComputeServiceRuntimeError(v int32)`

SetComputeServiceRuntimeError sets ComputeServiceRuntimeError field to given value.

### HasComputeServiceRuntimeError

`func (o *PlatformValues) HasComputeServiceRuntimeError() bool`

HasComputeServiceRuntimeError returns a boolean if a field has been set.

### GetComputeServiceChainError

`func (o *PlatformValues) GetComputeServiceChainError() int32`

GetComputeServiceChainError returns the ComputeServiceChainError field if non-nil, zero value otherwise.

### GetComputeServiceChainErrorOk

`func (o *PlatformValues) GetComputeServiceChainErrorOk() (*int32, bool)`

GetComputeServiceChainErrorOk returns a tuple with the ComputeServiceChainError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeServiceChainError

`func (o *PlatformValues) SetComputeServiceChainError(v int32)`

SetComputeServiceChainError sets ComputeServiceChainError field to given value.

### HasComputeServiceChainError

`func (o *PlatformValues) HasComputeServiceChainError() bool`

HasComputeServiceChainError returns a boolean if a field has been set.

### GetComputePlatformInternalError

`func (o *PlatformValues) GetComputePlatformInternalError() int32`

GetComputePlatformInternalError returns the ComputePlatformInternalError field if non-nil, zero value otherwise.

### GetComputePlatformInternalErrorOk

`func (o *PlatformValues) GetComputePlatformInternalErrorOk() (*int32, bool)`

GetComputePlatformInternalErrorOk returns a tuple with the ComputePlatformInternalError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputePlatformInternalError

`func (o *PlatformValues) SetComputePlatformInternalError(v int32)`

SetComputePlatformInternalError sets ComputePlatformInternalError field to given value.

### HasComputePlatformInternalError

`func (o *PlatformValues) HasComputePlatformInternalError() bool`

HasComputePlatformInternalError returns a boolean if a field has been set.

### GetComputeServiceTimeoutError

`func (o *PlatformValues) GetComputeServiceTimeoutError() int32`

GetComputeServiceTimeoutError returns the ComputeServiceTimeoutError field if non-nil, zero value otherwise.

### GetComputeServiceTimeoutErrorOk

`func (o *PlatformValues) GetComputeServiceTimeoutErrorOk() (*int32, bool)`

GetComputeServiceTimeoutErrorOk returns a tuple with the ComputeServiceTimeoutError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeServiceTimeoutError

`func (o *PlatformValues) SetComputeServiceTimeoutError(v int32)`

SetComputeServiceTimeoutError sets ComputeServiceTimeoutError field to given value.

### HasComputeServiceTimeoutError

`func (o *PlatformValues) HasComputeServiceTimeoutError() bool`

HasComputeServiceTimeoutError returns a boolean if a field has been set.

### GetComputeServiceVcpuExceededError

`func (o *PlatformValues) GetComputeServiceVcpuExceededError() int32`

GetComputeServiceVcpuExceededError returns the ComputeServiceVcpuExceededError field if non-nil, zero value otherwise.

### GetComputeServiceVcpuExceededErrorOk

`func (o *PlatformValues) GetComputeServiceVcpuExceededErrorOk() (*int32, bool)`

GetComputeServiceVcpuExceededErrorOk returns a tuple with the ComputeServiceVcpuExceededError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeServiceVcpuExceededError

`func (o *PlatformValues) SetComputeServiceVcpuExceededError(v int32)`

SetComputeServiceVcpuExceededError sets ComputeServiceVcpuExceededError field to given value.

### HasComputeServiceVcpuExceededError

`func (o *PlatformValues) HasComputeServiceVcpuExceededError() bool`

HasComputeServiceVcpuExceededError returns a boolean if a field has been set.

### GetComputeServiceLimitsError

`func (o *PlatformValues) GetComputeServiceLimitsError() int32`

GetComputeServiceLimitsError returns the ComputeServiceLimitsError field if non-nil, zero value otherwise.

### GetComputeServiceLimitsErrorOk

`func (o *PlatformValues) GetComputeServiceLimitsErrorOk() (*int32, bool)`

GetComputeServiceLimitsErrorOk returns a tuple with the ComputeServiceLimitsError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeServiceLimitsError

`func (o *PlatformValues) SetComputeServiceLimitsError(v int32)`

SetComputeServiceLimitsError sets ComputeServiceLimitsError field to given value.

### HasComputeServiceLimitsError

`func (o *PlatformValues) HasComputeServiceLimitsError() bool`

HasComputeServiceLimitsError returns a boolean if a field has been set.

### GetComputePlatformInvalidRequestError

`func (o *PlatformValues) GetComputePlatformInvalidRequestError() int32`

GetComputePlatformInvalidRequestError returns the ComputePlatformInvalidRequestError field if non-nil, zero value otherwise.

### GetComputePlatformInvalidRequestErrorOk

`func (o *PlatformValues) GetComputePlatformInvalidRequestErrorOk() (*int32, bool)`

GetComputePlatformInvalidRequestErrorOk returns a tuple with the ComputePlatformInvalidRequestError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputePlatformInvalidRequestError

`func (o *PlatformValues) SetComputePlatformInvalidRequestError(v int32)`

SetComputePlatformInvalidRequestError sets ComputePlatformInvalidRequestError field to given value.

### HasComputePlatformInvalidRequestError

`func (o *PlatformValues) HasComputePlatformInvalidRequestError() bool`

HasComputePlatformInvalidRequestError returns a boolean if a field has been set.

### GetComputeSandboxes

`func (o *PlatformValues) GetComputeSandboxes() int32`

GetComputeSandboxes returns the ComputeSandboxes field if non-nil, zero value otherwise.

### GetComputeSandboxesOk

`func (o *PlatformValues) GetComputeSandboxesOk() (*int32, bool)`

GetComputeSandboxesOk returns a tuple with the ComputeSandboxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeSandboxes

`func (o *PlatformValues) SetComputeSandboxes(v int32)`

SetComputeSandboxes sets ComputeSandboxes field to given value.

### HasComputeSandboxes

`func (o *PlatformValues) HasComputeSandboxes() bool`

HasComputeSandboxes returns a boolean if a field has been set.

### GetBotRequestsTotalCount

`func (o *PlatformValues) GetBotRequestsTotalCount() int32`

GetBotRequestsTotalCount returns the BotRequestsTotalCount field if non-nil, zero value otherwise.

### GetBotRequestsTotalCountOk

`func (o *PlatformValues) GetBotRequestsTotalCountOk() (*int32, bool)`

GetBotRequestsTotalCountOk returns a tuple with the BotRequestsTotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotRequestsTotalCount

`func (o *PlatformValues) SetBotRequestsTotalCount(v int32)`

SetBotRequestsTotalCount sets BotRequestsTotalCount field to given value.

### HasBotRequestsTotalCount

`func (o *PlatformValues) HasBotRequestsTotalCount() bool`

HasBotRequestsTotalCount returns a boolean if a field has been set.

### GetBotEdgeRequestsAnalyzedCount

`func (o *PlatformValues) GetBotEdgeRequestsAnalyzedCount() int32`

GetBotEdgeRequestsAnalyzedCount returns the BotEdgeRequestsAnalyzedCount field if non-nil, zero value otherwise.

### GetBotEdgeRequestsAnalyzedCountOk

`func (o *PlatformValues) GetBotEdgeRequestsAnalyzedCountOk() (*int32, bool)`

GetBotEdgeRequestsAnalyzedCountOk returns a tuple with the BotEdgeRequestsAnalyzedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotEdgeRequestsAnalyzedCount

`func (o *PlatformValues) SetBotEdgeRequestsAnalyzedCount(v int32)`

SetBotEdgeRequestsAnalyzedCount sets BotEdgeRequestsAnalyzedCount field to given value.

### HasBotEdgeRequestsAnalyzedCount

`func (o *PlatformValues) HasBotEdgeRequestsAnalyzedCount() bool`

HasBotEdgeRequestsAnalyzedCount returns a boolean if a field has been set.

### GetBotEdgeRequestsDetectedCount

`func (o *PlatformValues) GetBotEdgeRequestsDetectedCount() int32`

GetBotEdgeRequestsDetectedCount returns the BotEdgeRequestsDetectedCount field if non-nil, zero value otherwise.

### GetBotEdgeRequestsDetectedCountOk

`func (o *PlatformValues) GetBotEdgeRequestsDetectedCountOk() (*int32, bool)`

GetBotEdgeRequestsDetectedCountOk returns a tuple with the BotEdgeRequestsDetectedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotEdgeRequestsDetectedCount

`func (o *PlatformValues) SetBotEdgeRequestsDetectedCount(v int32)`

SetBotEdgeRequestsDetectedCount sets BotEdgeRequestsDetectedCount field to given value.

### HasBotEdgeRequestsDetectedCount

`func (o *PlatformValues) HasBotEdgeRequestsDetectedCount() bool`

HasBotEdgeRequestsDetectedCount returns a boolean if a field has been set.

### GetBotEdgeRequestsVerifiedCount

`func (o *PlatformValues) GetBotEdgeRequestsVerifiedCount() int32`

GetBotEdgeRequestsVerifiedCount returns the BotEdgeRequestsVerifiedCount field if non-nil, zero value otherwise.

### GetBotEdgeRequestsVerifiedCountOk

`func (o *PlatformValues) GetBotEdgeRequestsVerifiedCountOk() (*int32, bool)`

GetBotEdgeRequestsVerifiedCountOk returns a tuple with the BotEdgeRequestsVerifiedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotEdgeRequestsVerifiedCount

`func (o *PlatformValues) SetBotEdgeRequestsVerifiedCount(v int32)`

SetBotEdgeRequestsVerifiedCount sets BotEdgeRequestsVerifiedCount field to given value.

### HasBotEdgeRequestsVerifiedCount

`func (o *PlatformValues) HasBotEdgeRequestsVerifiedCount() bool`

HasBotEdgeRequestsVerifiedCount returns a boolean if a field has been set.

### GetBotEdgeRequestsAiCrawlerCount

`func (o *PlatformValues) GetBotEdgeRequestsAiCrawlerCount() int32`

GetBotEdgeRequestsAiCrawlerCount returns the BotEdgeRequestsAiCrawlerCount field if non-nil, zero value otherwise.

### GetBotEdgeRequestsAiCrawlerCountOk

`func (o *PlatformValues) GetBotEdgeRequestsAiCrawlerCountOk() (*int32, bool)`

GetBotEdgeRequestsAiCrawlerCountOk returns a tuple with the BotEdgeRequestsAiCrawlerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotEdgeRequestsAiCrawlerCount

`func (o *PlatformValues) SetBotEdgeRequestsAiCrawlerCount(v int32)`

SetBotEdgeRequestsAiCrawlerCount sets BotEdgeRequestsAiCrawlerCount field to given value.

### HasBotEdgeRequestsAiCrawlerCount

`func (o *PlatformValues) HasBotEdgeRequestsAiCrawlerCount() bool`

HasBotEdgeRequestsAiCrawlerCount returns a boolean if a field has been set.

### GetBotEdgeRequestsAiFetcherCount

`func (o *PlatformValues) GetBotEdgeRequestsAiFetcherCount() int32`

GetBotEdgeRequestsAiFetcherCount returns the BotEdgeRequestsAiFetcherCount field if non-nil, zero value otherwise.

### GetBotEdgeRequestsAiFetcherCountOk

`func (o *PlatformValues) GetBotEdgeRequestsAiFetcherCountOk() (*int32, bool)`

GetBotEdgeRequestsAiFetcherCountOk returns a tuple with the BotEdgeRequestsAiFetcherCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotEdgeRequestsAiFetcherCount

`func (o *PlatformValues) SetBotEdgeRequestsAiFetcherCount(v int32)`

SetBotEdgeRequestsAiFetcherCount sets BotEdgeRequestsAiFetcherCount field to given value.

### HasBotEdgeRequestsAiFetcherCount

`func (o *PlatformValues) HasBotEdgeRequestsAiFetcherCount() bool`

HasBotEdgeRequestsAiFetcherCount returns a boolean if a field has been set.

### GetBotEdgeRequestsAccessibilityCount

`func (o *PlatformValues) GetBotEdgeRequestsAccessibilityCount() int32`

GetBotEdgeRequestsAccessibilityCount returns the BotEdgeRequestsAccessibilityCount field if non-nil, zero value otherwise.

### GetBotEdgeRequestsAccessibilityCountOk

`func (o *PlatformValues) GetBotEdgeRequestsAccessibilityCountOk() (*int32, bool)`

GetBotEdgeRequestsAccessibilityCountOk returns a tuple with the BotEdgeRequestsAccessibilityCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotEdgeRequestsAccessibilityCount

`func (o *PlatformValues) SetBotEdgeRequestsAccessibilityCount(v int32)`

SetBotEdgeRequestsAccessibilityCount sets BotEdgeRequestsAccessibilityCount field to given value.

### HasBotEdgeRequestsAccessibilityCount

`func (o *PlatformValues) HasBotEdgeRequestsAccessibilityCount() bool`

HasBotEdgeRequestsAccessibilityCount returns a boolean if a field has been set.

### GetBotEdgeRequestsContentFetcherCount

`func (o *PlatformValues) GetBotEdgeRequestsContentFetcherCount() int32`

GetBotEdgeRequestsContentFetcherCount returns the BotEdgeRequestsContentFetcherCount field if non-nil, zero value otherwise.

### GetBotEdgeRequestsContentFetcherCountOk

`func (o *PlatformValues) GetBotEdgeRequestsContentFetcherCountOk() (*int32, bool)`

GetBotEdgeRequestsContentFetcherCountOk returns a tuple with the BotEdgeRequestsContentFetcherCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotEdgeRequestsContentFetcherCount

`func (o *PlatformValues) SetBotEdgeRequestsContentFetcherCount(v int32)`

SetBotEdgeRequestsContentFetcherCount sets BotEdgeRequestsContentFetcherCount field to given value.

### HasBotEdgeRequestsContentFetcherCount

`func (o *PlatformValues) HasBotEdgeRequestsContentFetcherCount() bool`

HasBotEdgeRequestsContentFetcherCount returns a boolean if a field has been set.

### GetBotEdgeRequestsMonitoringCount

`func (o *PlatformValues) GetBotEdgeRequestsMonitoringCount() int32`

GetBotEdgeRequestsMonitoringCount returns the BotEdgeRequestsMonitoringCount field if non-nil, zero value otherwise.

### GetBotEdgeRequestsMonitoringCountOk

`func (o *PlatformValues) GetBotEdgeRequestsMonitoringCountOk() (*int32, bool)`

GetBotEdgeRequestsMonitoringCountOk returns a tuple with the BotEdgeRequestsMonitoringCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotEdgeRequestsMonitoringCount

`func (o *PlatformValues) SetBotEdgeRequestsMonitoringCount(v int32)`

SetBotEdgeRequestsMonitoringCount sets BotEdgeRequestsMonitoringCount field to given value.

### HasBotEdgeRequestsMonitoringCount

`func (o *PlatformValues) HasBotEdgeRequestsMonitoringCount() bool`

HasBotEdgeRequestsMonitoringCount returns a boolean if a field has been set.

### GetBotEdgeRequestsOnlineMarketingCount

`func (o *PlatformValues) GetBotEdgeRequestsOnlineMarketingCount() int32`

GetBotEdgeRequestsOnlineMarketingCount returns the BotEdgeRequestsOnlineMarketingCount field if non-nil, zero value otherwise.

### GetBotEdgeRequestsOnlineMarketingCountOk

`func (o *PlatformValues) GetBotEdgeRequestsOnlineMarketingCountOk() (*int32, bool)`

GetBotEdgeRequestsOnlineMarketingCountOk returns a tuple with the BotEdgeRequestsOnlineMarketingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotEdgeRequestsOnlineMarketingCount

`func (o *PlatformValues) SetBotEdgeRequestsOnlineMarketingCount(v int32)`

SetBotEdgeRequestsOnlineMarketingCount sets BotEdgeRequestsOnlineMarketingCount field to given value.

### HasBotEdgeRequestsOnlineMarketingCount

`func (o *PlatformValues) HasBotEdgeRequestsOnlineMarketingCount() bool`

HasBotEdgeRequestsOnlineMarketingCount returns a boolean if a field has been set.

### GetBotEdgeRequestsPagePreviewCount

`func (o *PlatformValues) GetBotEdgeRequestsPagePreviewCount() int32`

GetBotEdgeRequestsPagePreviewCount returns the BotEdgeRequestsPagePreviewCount field if non-nil, zero value otherwise.

### GetBotEdgeRequestsPagePreviewCountOk

`func (o *PlatformValues) GetBotEdgeRequestsPagePreviewCountOk() (*int32, bool)`

GetBotEdgeRequestsPagePreviewCountOk returns a tuple with the BotEdgeRequestsPagePreviewCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotEdgeRequestsPagePreviewCount

`func (o *PlatformValues) SetBotEdgeRequestsPagePreviewCount(v int32)`

SetBotEdgeRequestsPagePreviewCount sets BotEdgeRequestsPagePreviewCount field to given value.

### HasBotEdgeRequestsPagePreviewCount

`func (o *PlatformValues) HasBotEdgeRequestsPagePreviewCount() bool`

HasBotEdgeRequestsPagePreviewCount returns a boolean if a field has been set.

### GetBotEdgeRequestsPlatformIntegrationsCount

`func (o *PlatformValues) GetBotEdgeRequestsPlatformIntegrationsCount() int32`

GetBotEdgeRequestsPlatformIntegrationsCount returns the BotEdgeRequestsPlatformIntegrationsCount field if non-nil, zero value otherwise.

### GetBotEdgeRequestsPlatformIntegrationsCountOk

`func (o *PlatformValues) GetBotEdgeRequestsPlatformIntegrationsCountOk() (*int32, bool)`

GetBotEdgeRequestsPlatformIntegrationsCountOk returns a tuple with the BotEdgeRequestsPlatformIntegrationsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotEdgeRequestsPlatformIntegrationsCount

`func (o *PlatformValues) SetBotEdgeRequestsPlatformIntegrationsCount(v int32)`

SetBotEdgeRequestsPlatformIntegrationsCount sets BotEdgeRequestsPlatformIntegrationsCount field to given value.

### HasBotEdgeRequestsPlatformIntegrationsCount

`func (o *PlatformValues) HasBotEdgeRequestsPlatformIntegrationsCount() bool`

HasBotEdgeRequestsPlatformIntegrationsCount returns a boolean if a field has been set.

### GetBotEdgeRequestsResearchCount

`func (o *PlatformValues) GetBotEdgeRequestsResearchCount() int32`

GetBotEdgeRequestsResearchCount returns the BotEdgeRequestsResearchCount field if non-nil, zero value otherwise.

### GetBotEdgeRequestsResearchCountOk

`func (o *PlatformValues) GetBotEdgeRequestsResearchCountOk() (*int32, bool)`

GetBotEdgeRequestsResearchCountOk returns a tuple with the BotEdgeRequestsResearchCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotEdgeRequestsResearchCount

`func (o *PlatformValues) SetBotEdgeRequestsResearchCount(v int32)`

SetBotEdgeRequestsResearchCount sets BotEdgeRequestsResearchCount field to given value.

### HasBotEdgeRequestsResearchCount

`func (o *PlatformValues) HasBotEdgeRequestsResearchCount() bool`

HasBotEdgeRequestsResearchCount returns a boolean if a field has been set.

### GetBotEdgeRequestsSearchEngineCrawlerCount

`func (o *PlatformValues) GetBotEdgeRequestsSearchEngineCrawlerCount() int32`

GetBotEdgeRequestsSearchEngineCrawlerCount returns the BotEdgeRequestsSearchEngineCrawlerCount field if non-nil, zero value otherwise.

### GetBotEdgeRequestsSearchEngineCrawlerCountOk

`func (o *PlatformValues) GetBotEdgeRequestsSearchEngineCrawlerCountOk() (*int32, bool)`

GetBotEdgeRequestsSearchEngineCrawlerCountOk returns a tuple with the BotEdgeRequestsSearchEngineCrawlerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotEdgeRequestsSearchEngineCrawlerCount

`func (o *PlatformValues) SetBotEdgeRequestsSearchEngineCrawlerCount(v int32)`

SetBotEdgeRequestsSearchEngineCrawlerCount sets BotEdgeRequestsSearchEngineCrawlerCount field to given value.

### HasBotEdgeRequestsSearchEngineCrawlerCount

`func (o *PlatformValues) HasBotEdgeRequestsSearchEngineCrawlerCount() bool`

HasBotEdgeRequestsSearchEngineCrawlerCount returns a boolean if a field has been set.

### GetBotEdgeRequestsSearchEngineOptimizationCount

`func (o *PlatformValues) GetBotEdgeRequestsSearchEngineOptimizationCount() int32`

GetBotEdgeRequestsSearchEngineOptimizationCount returns the BotEdgeRequestsSearchEngineOptimizationCount field if non-nil, zero value otherwise.

### GetBotEdgeRequestsSearchEngineOptimizationCountOk

`func (o *PlatformValues) GetBotEdgeRequestsSearchEngineOptimizationCountOk() (*int32, bool)`

GetBotEdgeRequestsSearchEngineOptimizationCountOk returns a tuple with the BotEdgeRequestsSearchEngineOptimizationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotEdgeRequestsSearchEngineOptimizationCount

`func (o *PlatformValues) SetBotEdgeRequestsSearchEngineOptimizationCount(v int32)`

SetBotEdgeRequestsSearchEngineOptimizationCount sets BotEdgeRequestsSearchEngineOptimizationCount field to given value.

### HasBotEdgeRequestsSearchEngineOptimizationCount

`func (o *PlatformValues) HasBotEdgeRequestsSearchEngineOptimizationCount() bool`

HasBotEdgeRequestsSearchEngineOptimizationCount returns a boolean if a field has been set.

### GetBotEdgeRequestsSecurityToolsCount

`func (o *PlatformValues) GetBotEdgeRequestsSecurityToolsCount() int32`

GetBotEdgeRequestsSecurityToolsCount returns the BotEdgeRequestsSecurityToolsCount field if non-nil, zero value otherwise.

### GetBotEdgeRequestsSecurityToolsCountOk

`func (o *PlatformValues) GetBotEdgeRequestsSecurityToolsCountOk() (*int32, bool)`

GetBotEdgeRequestsSecurityToolsCountOk returns a tuple with the BotEdgeRequestsSecurityToolsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotEdgeRequestsSecurityToolsCount

`func (o *PlatformValues) SetBotEdgeRequestsSecurityToolsCount(v int32)`

SetBotEdgeRequestsSecurityToolsCount sets BotEdgeRequestsSecurityToolsCount field to given value.

### HasBotEdgeRequestsSecurityToolsCount

`func (o *PlatformValues) HasBotEdgeRequestsSecurityToolsCount() bool`

HasBotEdgeRequestsSecurityToolsCount returns a boolean if a field has been set.

### GetComputeHandoff

`func (o *PlatformValues) GetComputeHandoff() int32`

GetComputeHandoff returns the ComputeHandoff field if non-nil, zero value otherwise.

### GetComputeHandoffOk

`func (o *PlatformValues) GetComputeHandoffOk() (*int32, bool)`

GetComputeHandoffOk returns a tuple with the ComputeHandoff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeHandoff

`func (o *PlatformValues) SetComputeHandoff(v int32)`

SetComputeHandoff sets ComputeHandoff field to given value.

### HasComputeHandoff

`func (o *PlatformValues) HasComputeHandoff() bool`

HasComputeHandoff returns a boolean if a field has been set.

### GetComputeServiceBereqDNSError

`func (o *PlatformValues) GetComputeServiceBereqDNSError() int32`

GetComputeServiceBereqDNSError returns the ComputeServiceBereqDNSError field if non-nil, zero value otherwise.

### GetComputeServiceBereqDNSErrorOk

`func (o *PlatformValues) GetComputeServiceBereqDNSErrorOk() (*int32, bool)`

GetComputeServiceBereqDNSErrorOk returns a tuple with the ComputeServiceBereqDNSError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeServiceBereqDNSError

`func (o *PlatformValues) SetComputeServiceBereqDNSError(v int32)`

SetComputeServiceBereqDNSError sets ComputeServiceBereqDNSError field to given value.

### HasComputeServiceBereqDNSError

`func (o *PlatformValues) HasComputeServiceBereqDNSError() bool`

HasComputeServiceBereqDNSError returns a boolean if a field has been set.

### GetComputeServiceBereqConnTimeoutError

`func (o *PlatformValues) GetComputeServiceBereqConnTimeoutError() int32`

GetComputeServiceBereqConnTimeoutError returns the ComputeServiceBereqConnTimeoutError field if non-nil, zero value otherwise.

### GetComputeServiceBereqConnTimeoutErrorOk

`func (o *PlatformValues) GetComputeServiceBereqConnTimeoutErrorOk() (*int32, bool)`

GetComputeServiceBereqConnTimeoutErrorOk returns a tuple with the ComputeServiceBereqConnTimeoutError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeServiceBereqConnTimeoutError

`func (o *PlatformValues) SetComputeServiceBereqConnTimeoutError(v int32)`

SetComputeServiceBereqConnTimeoutError sets ComputeServiceBereqConnTimeoutError field to given value.

### HasComputeServiceBereqConnTimeoutError

`func (o *PlatformValues) HasComputeServiceBereqConnTimeoutError() bool`

HasComputeServiceBereqConnTimeoutError returns a boolean if a field has been set.

### GetComputeServiceBereqConnRefusedError

`func (o *PlatformValues) GetComputeServiceBereqConnRefusedError() int32`

GetComputeServiceBereqConnRefusedError returns the ComputeServiceBereqConnRefusedError field if non-nil, zero value otherwise.

### GetComputeServiceBereqConnRefusedErrorOk

`func (o *PlatformValues) GetComputeServiceBereqConnRefusedErrorOk() (*int32, bool)`

GetComputeServiceBereqConnRefusedErrorOk returns a tuple with the ComputeServiceBereqConnRefusedError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeServiceBereqConnRefusedError

`func (o *PlatformValues) SetComputeServiceBereqConnRefusedError(v int32)`

SetComputeServiceBereqConnRefusedError sets ComputeServiceBereqConnRefusedError field to given value.

### HasComputeServiceBereqConnRefusedError

`func (o *PlatformValues) HasComputeServiceBereqConnRefusedError() bool`

HasComputeServiceBereqConnRefusedError returns a boolean if a field has been set.

### GetComputeServiceBereqConnOtherError

`func (o *PlatformValues) GetComputeServiceBereqConnOtherError() int32`

GetComputeServiceBereqConnOtherError returns the ComputeServiceBereqConnOtherError field if non-nil, zero value otherwise.

### GetComputeServiceBereqConnOtherErrorOk

`func (o *PlatformValues) GetComputeServiceBereqConnOtherErrorOk() (*int32, bool)`

GetComputeServiceBereqConnOtherErrorOk returns a tuple with the ComputeServiceBereqConnOtherError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeServiceBereqConnOtherError

`func (o *PlatformValues) SetComputeServiceBereqConnOtherError(v int32)`

SetComputeServiceBereqConnOtherError sets ComputeServiceBereqConnOtherError field to given value.

### HasComputeServiceBereqConnOtherError

`func (o *PlatformValues) HasComputeServiceBereqConnOtherError() bool`

HasComputeServiceBereqConnOtherError returns a boolean if a field has been set.

### GetComputeServiceBereqTLSServerCertError

`func (o *PlatformValues) GetComputeServiceBereqTLSServerCertError() int32`

GetComputeServiceBereqTLSServerCertError returns the ComputeServiceBereqTLSServerCertError field if non-nil, zero value otherwise.

### GetComputeServiceBereqTLSServerCertErrorOk

`func (o *PlatformValues) GetComputeServiceBereqTLSServerCertErrorOk() (*int32, bool)`

GetComputeServiceBereqTLSServerCertErrorOk returns a tuple with the ComputeServiceBereqTLSServerCertError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeServiceBereqTLSServerCertError

`func (o *PlatformValues) SetComputeServiceBereqTLSServerCertError(v int32)`

SetComputeServiceBereqTLSServerCertError sets ComputeServiceBereqTLSServerCertError field to given value.

### HasComputeServiceBereqTLSServerCertError

`func (o *PlatformValues) HasComputeServiceBereqTLSServerCertError() bool`

HasComputeServiceBereqTLSServerCertError returns a boolean if a field has been set.

### GetComputeServiceBereqTLSOtherError

`func (o *PlatformValues) GetComputeServiceBereqTLSOtherError() int32`

GetComputeServiceBereqTLSOtherError returns the ComputeServiceBereqTLSOtherError field if non-nil, zero value otherwise.

### GetComputeServiceBereqTLSOtherErrorOk

`func (o *PlatformValues) GetComputeServiceBereqTLSOtherErrorOk() (*int32, bool)`

GetComputeServiceBereqTLSOtherErrorOk returns a tuple with the ComputeServiceBereqTLSOtherError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeServiceBereqTLSOtherError

`func (o *PlatformValues) SetComputeServiceBereqTLSOtherError(v int32)`

SetComputeServiceBereqTLSOtherError sets ComputeServiceBereqTLSOtherError field to given value.

### HasComputeServiceBereqTLSOtherError

`func (o *PlatformValues) HasComputeServiceBereqTLSOtherError() bool`

HasComputeServiceBereqTLSOtherError returns a boolean if a field has been set.

### GetComputeServiceBereqHTTPProtoV1Error

`func (o *PlatformValues) GetComputeServiceBereqHTTPProtoV1Error() int32`

GetComputeServiceBereqHTTPProtoV1Error returns the ComputeServiceBereqHTTPProtoV1Error field if non-nil, zero value otherwise.

### GetComputeServiceBereqHTTPProtoV1ErrorOk

`func (o *PlatformValues) GetComputeServiceBereqHTTPProtoV1ErrorOk() (*int32, bool)`

GetComputeServiceBereqHTTPProtoV1ErrorOk returns a tuple with the ComputeServiceBereqHTTPProtoV1Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeServiceBereqHTTPProtoV1Error

`func (o *PlatformValues) SetComputeServiceBereqHTTPProtoV1Error(v int32)`

SetComputeServiceBereqHTTPProtoV1Error sets ComputeServiceBereqHTTPProtoV1Error field to given value.

### HasComputeServiceBereqHTTPProtoV1Error

`func (o *PlatformValues) HasComputeServiceBereqHTTPProtoV1Error() bool`

HasComputeServiceBereqHTTPProtoV1Error returns a boolean if a field has been set.

### GetComputeServiceBereqHTTPProtoV2Error

`func (o *PlatformValues) GetComputeServiceBereqHTTPProtoV2Error() int32`

GetComputeServiceBereqHTTPProtoV2Error returns the ComputeServiceBereqHTTPProtoV2Error field if non-nil, zero value otherwise.

### GetComputeServiceBereqHTTPProtoV2ErrorOk

`func (o *PlatformValues) GetComputeServiceBereqHTTPProtoV2ErrorOk() (*int32, bool)`

GetComputeServiceBereqHTTPProtoV2ErrorOk returns a tuple with the ComputeServiceBereqHTTPProtoV2Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeServiceBereqHTTPProtoV2Error

`func (o *PlatformValues) SetComputeServiceBereqHTTPProtoV2Error(v int32)`

SetComputeServiceBereqHTTPProtoV2Error sets ComputeServiceBereqHTTPProtoV2Error field to given value.

### HasComputeServiceBereqHTTPProtoV2Error

`func (o *PlatformValues) HasComputeServiceBereqHTTPProtoV2Error() bool`

HasComputeServiceBereqHTTPProtoV2Error returns a boolean if a field has been set.

### GetComputeServiceBereqHTTPIncompleteError

`func (o *PlatformValues) GetComputeServiceBereqHTTPIncompleteError() int32`

GetComputeServiceBereqHTTPIncompleteError returns the ComputeServiceBereqHTTPIncompleteError field if non-nil, zero value otherwise.

### GetComputeServiceBereqHTTPIncompleteErrorOk

`func (o *PlatformValues) GetComputeServiceBereqHTTPIncompleteErrorOk() (*int32, bool)`

GetComputeServiceBereqHTTPIncompleteErrorOk returns a tuple with the ComputeServiceBereqHTTPIncompleteError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeServiceBereqHTTPIncompleteError

`func (o *PlatformValues) SetComputeServiceBereqHTTPIncompleteError(v int32)`

SetComputeServiceBereqHTTPIncompleteError sets ComputeServiceBereqHTTPIncompleteError field to given value.

### HasComputeServiceBereqHTTPIncompleteError

`func (o *PlatformValues) HasComputeServiceBereqHTTPIncompleteError() bool`

HasComputeServiceBereqHTTPIncompleteError returns a boolean if a field has been set.

### GetComputeServiceBereqHTTPTimeoutError

`func (o *PlatformValues) GetComputeServiceBereqHTTPTimeoutError() int32`

GetComputeServiceBereqHTTPTimeoutError returns the ComputeServiceBereqHTTPTimeoutError field if non-nil, zero value otherwise.

### GetComputeServiceBereqHTTPTimeoutErrorOk

`func (o *PlatformValues) GetComputeServiceBereqHTTPTimeoutErrorOk() (*int32, bool)`

GetComputeServiceBereqHTTPTimeoutErrorOk returns a tuple with the ComputeServiceBereqHTTPTimeoutError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeServiceBereqHTTPTimeoutError

`func (o *PlatformValues) SetComputeServiceBereqHTTPTimeoutError(v int32)`

SetComputeServiceBereqHTTPTimeoutError sets ComputeServiceBereqHTTPTimeoutError field to given value.

### HasComputeServiceBereqHTTPTimeoutError

`func (o *PlatformValues) HasComputeServiceBereqHTTPTimeoutError() bool`

HasComputeServiceBereqHTTPTimeoutError returns a boolean if a field has been set.

### GetComputeServiceBereqHTTPOtherError

`func (o *PlatformValues) GetComputeServiceBereqHTTPOtherError() int32`

GetComputeServiceBereqHTTPOtherError returns the ComputeServiceBereqHTTPOtherError field if non-nil, zero value otherwise.

### GetComputeServiceBereqHTTPOtherErrorOk

`func (o *PlatformValues) GetComputeServiceBereqHTTPOtherErrorOk() (*int32, bool)`

GetComputeServiceBereqHTTPOtherErrorOk returns a tuple with the ComputeServiceBereqHTTPOtherError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeServiceBereqHTTPOtherError

`func (o *PlatformValues) SetComputeServiceBereqHTTPOtherError(v int32)`

SetComputeServiceBereqHTTPOtherError sets ComputeServiceBereqHTTPOtherError field to given value.

### HasComputeServiceBereqHTTPOtherError

`func (o *PlatformValues) HasComputeServiceBereqHTTPOtherError() bool`

HasComputeServiceBereqHTTPOtherError returns a boolean if a field has been set.

### GetComputeServiceBereqOtherError

`func (o *PlatformValues) GetComputeServiceBereqOtherError() int32`

GetComputeServiceBereqOtherError returns the ComputeServiceBereqOtherError field if non-nil, zero value otherwise.

### GetComputeServiceBereqOtherErrorOk

`func (o *PlatformValues) GetComputeServiceBereqOtherErrorOk() (*int32, bool)`

GetComputeServiceBereqOtherErrorOk returns a tuple with the ComputeServiceBereqOtherError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeServiceBereqOtherError

`func (o *PlatformValues) SetComputeServiceBereqOtherError(v int32)`

SetComputeServiceBereqOtherError sets ComputeServiceBereqOtherError field to given value.

### HasComputeServiceBereqOtherError

`func (o *PlatformValues) HasComputeServiceBereqOtherError() bool`

HasComputeServiceBereqOtherError returns a boolean if a field has been set.

### GetComputeServiceBereq5xxError

`func (o *PlatformValues) GetComputeServiceBereq5xxError() int32`

GetComputeServiceBereq5xxError returns the ComputeServiceBereq5xxError field if non-nil, zero value otherwise.

### GetComputeServiceBereq5xxErrorOk

`func (o *PlatformValues) GetComputeServiceBereq5xxErrorOk() (*int32, bool)`

GetComputeServiceBereq5xxErrorOk returns a tuple with the ComputeServiceBereq5xxError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeServiceBereq5xxError

`func (o *PlatformValues) SetComputeServiceBereq5xxError(v int32)`

SetComputeServiceBereq5xxError sets ComputeServiceBereq5xxError field to given value.

### HasComputeServiceBereq5xxError

`func (o *PlatformValues) HasComputeServiceBereq5xxError() bool`

HasComputeServiceBereq5xxError returns a boolean if a field has been set.

### GetComputeServiceBereqConnError

`func (o *PlatformValues) GetComputeServiceBereqConnError() int32`

GetComputeServiceBereqConnError returns the ComputeServiceBereqConnError field if non-nil, zero value otherwise.

### GetComputeServiceBereqConnErrorOk

`func (o *PlatformValues) GetComputeServiceBereqConnErrorOk() (*int32, bool)`

GetComputeServiceBereqConnErrorOk returns a tuple with the ComputeServiceBereqConnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeServiceBereqConnError

`func (o *PlatformValues) SetComputeServiceBereqConnError(v int32)`

SetComputeServiceBereqConnError sets ComputeServiceBereqConnError field to given value.

### HasComputeServiceBereqConnError

`func (o *PlatformValues) HasComputeServiceBereqConnError() bool`

HasComputeServiceBereqConnError returns a boolean if a field has been set.

### GetComputeServiceBereqTLSError

`func (o *PlatformValues) GetComputeServiceBereqTLSError() int32`

GetComputeServiceBereqTLSError returns the ComputeServiceBereqTLSError field if non-nil, zero value otherwise.

### GetComputeServiceBereqTLSErrorOk

`func (o *PlatformValues) GetComputeServiceBereqTLSErrorOk() (*int32, bool)`

GetComputeServiceBereqTLSErrorOk returns a tuple with the ComputeServiceBereqTLSError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeServiceBereqTLSError

`func (o *PlatformValues) SetComputeServiceBereqTLSError(v int32)`

SetComputeServiceBereqTLSError sets ComputeServiceBereqTLSError field to given value.

### HasComputeServiceBereqTLSError

`func (o *PlatformValues) HasComputeServiceBereqTLSError() bool`

HasComputeServiceBereqTLSError returns a boolean if a field has been set.

### GetComputeServiceBereqHTTPError

`func (o *PlatformValues) GetComputeServiceBereqHTTPError() int32`

GetComputeServiceBereqHTTPError returns the ComputeServiceBereqHTTPError field if non-nil, zero value otherwise.

### GetComputeServiceBereqHTTPErrorOk

`func (o *PlatformValues) GetComputeServiceBereqHTTPErrorOk() (*int32, bool)`

GetComputeServiceBereqHTTPErrorOk returns a tuple with the ComputeServiceBereqHTTPError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeServiceBereqHTTPError

`func (o *PlatformValues) SetComputeServiceBereqHTTPError(v int32)`

SetComputeServiceBereqHTTPError sets ComputeServiceBereqHTTPError field to given value.

### HasComputeServiceBereqHTTPError

`func (o *PlatformValues) HasComputeServiceBereqHTTPError() bool`

HasComputeServiceBereqHTTPError returns a boolean if a field has been set.

### GetBotChallengesPatsIssued

`func (o *PlatformValues) GetBotChallengesPatsIssued() int32`

GetBotChallengesPatsIssued returns the BotChallengesPatsIssued field if non-nil, zero value otherwise.

### GetBotChallengesPatsIssuedOk

`func (o *PlatformValues) GetBotChallengesPatsIssuedOk() (*int32, bool)`

GetBotChallengesPatsIssuedOk returns a tuple with the BotChallengesPatsIssued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotChallengesPatsIssued

`func (o *PlatformValues) SetBotChallengesPatsIssued(v int32)`

SetBotChallengesPatsIssued sets BotChallengesPatsIssued field to given value.

### HasBotChallengesPatsIssued

`func (o *PlatformValues) HasBotChallengesPatsIssued() bool`

HasBotChallengesPatsIssued returns a boolean if a field has been set.

### GetBotChallengesPatsSucceeded

`func (o *PlatformValues) GetBotChallengesPatsSucceeded() int32`

GetBotChallengesPatsSucceeded returns the BotChallengesPatsSucceeded field if non-nil, zero value otherwise.

### GetBotChallengesPatsSucceededOk

`func (o *PlatformValues) GetBotChallengesPatsSucceededOk() (*int32, bool)`

GetBotChallengesPatsSucceededOk returns a tuple with the BotChallengesPatsSucceeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotChallengesPatsSucceeded

`func (o *PlatformValues) SetBotChallengesPatsSucceeded(v int32)`

SetBotChallengesPatsSucceeded sets BotChallengesPatsSucceeded field to given value.

### HasBotChallengesPatsSucceeded

`func (o *PlatformValues) HasBotChallengesPatsSucceeded() bool`

HasBotChallengesPatsSucceeded returns a boolean if a field has been set.

### GetBotEdgeRequestsHeadlessCount

`func (o *PlatformValues) GetBotEdgeRequestsHeadlessCount() int32`

GetBotEdgeRequestsHeadlessCount returns the BotEdgeRequestsHeadlessCount field if non-nil, zero value otherwise.

### GetBotEdgeRequestsHeadlessCountOk

`func (o *PlatformValues) GetBotEdgeRequestsHeadlessCountOk() (*int32, bool)`

GetBotEdgeRequestsHeadlessCountOk returns a tuple with the BotEdgeRequestsHeadlessCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotEdgeRequestsHeadlessCount

`func (o *PlatformValues) SetBotEdgeRequestsHeadlessCount(v int32)`

SetBotEdgeRequestsHeadlessCount sets BotEdgeRequestsHeadlessCount field to given value.

### HasBotEdgeRequestsHeadlessCount

`func (o *PlatformValues) HasBotEdgeRequestsHeadlessCount() bool`

HasBotEdgeRequestsHeadlessCount returns a boolean if a field has been set.

### GetStatus499

`func (o *PlatformValues) GetStatus499() int32`

GetStatus499 returns the Status499 field if non-nil, zero value otherwise.

### GetStatus499Ok

`func (o *PlatformValues) GetStatus499Ok() (*int32, bool)`

GetStatus499Ok returns a tuple with the Status499 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus499

`func (o *PlatformValues) SetStatus499(v int32)`

SetStatus499 sets Status499 field to given value.

### HasStatus499

`func (o *PlatformValues) HasStatus499() bool`

HasStatus499 returns a boolean if a field has been set.

### GetArcRequests

`func (o *PlatformValues) GetArcRequests() int32`

GetArcRequests returns the ArcRequests field if non-nil, zero value otherwise.

### GetArcRequestsOk

`func (o *PlatformValues) GetArcRequestsOk() (*int32, bool)`

GetArcRequestsOk returns a tuple with the ArcRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArcRequests

`func (o *PlatformValues) SetArcRequests(v int32)`

SetArcRequests sets ArcRequests field to given value.

### HasArcRequests

`func (o *PlatformValues) HasArcRequests() bool`

HasArcRequests returns a boolean if a field has been set.

### GetDdosChallengeCompleteTokensChecked

`func (o *PlatformValues) GetDdosChallengeCompleteTokensChecked() int32`

GetDdosChallengeCompleteTokensChecked returns the DdosChallengeCompleteTokensChecked field if non-nil, zero value otherwise.

### GetDdosChallengeCompleteTokensCheckedOk

`func (o *PlatformValues) GetDdosChallengeCompleteTokensCheckedOk() (*int32, bool)`

GetDdosChallengeCompleteTokensCheckedOk returns a tuple with the DdosChallengeCompleteTokensChecked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdosChallengeCompleteTokensChecked

`func (o *PlatformValues) SetDdosChallengeCompleteTokensChecked(v int32)`

SetDdosChallengeCompleteTokensChecked sets DdosChallengeCompleteTokensChecked field to given value.

### HasDdosChallengeCompleteTokensChecked

`func (o *PlatformValues) HasDdosChallengeCompleteTokensChecked() bool`

HasDdosChallengeCompleteTokensChecked returns a boolean if a field has been set.

### GetDdosChallengeCompleteTokensPassed

`func (o *PlatformValues) GetDdosChallengeCompleteTokensPassed() int32`

GetDdosChallengeCompleteTokensPassed returns the DdosChallengeCompleteTokensPassed field if non-nil, zero value otherwise.

### GetDdosChallengeCompleteTokensPassedOk

`func (o *PlatformValues) GetDdosChallengeCompleteTokensPassedOk() (*int32, bool)`

GetDdosChallengeCompleteTokensPassedOk returns a tuple with the DdosChallengeCompleteTokensPassed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdosChallengeCompleteTokensPassed

`func (o *PlatformValues) SetDdosChallengeCompleteTokensPassed(v int32)`

SetDdosChallengeCompleteTokensPassed sets DdosChallengeCompleteTokensPassed field to given value.

### HasDdosChallengeCompleteTokensPassed

`func (o *PlatformValues) HasDdosChallengeCompleteTokensPassed() bool`

HasDdosChallengeCompleteTokensPassed returns a boolean if a field has been set.

### GetDdosChallengeCompleteTokensFailed

`func (o *PlatformValues) GetDdosChallengeCompleteTokensFailed() int32`

GetDdosChallengeCompleteTokensFailed returns the DdosChallengeCompleteTokensFailed field if non-nil, zero value otherwise.

### GetDdosChallengeCompleteTokensFailedOk

`func (o *PlatformValues) GetDdosChallengeCompleteTokensFailedOk() (*int32, bool)`

GetDdosChallengeCompleteTokensFailedOk returns a tuple with the DdosChallengeCompleteTokensFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdosChallengeCompleteTokensFailed

`func (o *PlatformValues) SetDdosChallengeCompleteTokensFailed(v int32)`

SetDdosChallengeCompleteTokensFailed sets DdosChallengeCompleteTokensFailed field to given value.

### HasDdosChallengeCompleteTokensFailed

`func (o *PlatformValues) HasDdosChallengeCompleteTokensFailed() bool`

HasDdosChallengeCompleteTokensFailed returns a boolean if a field has been set.

### GetDdosChallengeStarts

`func (o *PlatformValues) GetDdosChallengeStarts() int32`

GetDdosChallengeStarts returns the DdosChallengeStarts field if non-nil, zero value otherwise.

### GetDdosChallengeStartsOk

`func (o *PlatformValues) GetDdosChallengeStartsOk() (*int32, bool)`

GetDdosChallengeStartsOk returns a tuple with the DdosChallengeStarts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdosChallengeStarts

`func (o *PlatformValues) SetDdosChallengeStarts(v int32)`

SetDdosChallengeStarts sets DdosChallengeStarts field to given value.

### HasDdosChallengeStarts

`func (o *PlatformValues) HasDdosChallengeStarts() bool`

HasDdosChallengeStarts returns a boolean if a field has been set.

### GetDdosChallengeCompleteTokensIssued

`func (o *PlatformValues) GetDdosChallengeCompleteTokensIssued() int32`

GetDdosChallengeCompleteTokensIssued returns the DdosChallengeCompleteTokensIssued field if non-nil, zero value otherwise.

### GetDdosChallengeCompleteTokensIssuedOk

`func (o *PlatformValues) GetDdosChallengeCompleteTokensIssuedOk() (*int32, bool)`

GetDdosChallengeCompleteTokensIssuedOk returns a tuple with the DdosChallengeCompleteTokensIssued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdosChallengeCompleteTokensIssued

`func (o *PlatformValues) SetDdosChallengeCompleteTokensIssued(v int32)`

SetDdosChallengeCompleteTokensIssued sets DdosChallengeCompleteTokensIssued field to given value.

### HasDdosChallengeCompleteTokensIssued

`func (o *PlatformValues) HasDdosChallengeCompleteTokensIssued() bool`

HasDdosChallengeCompleteTokensIssued returns a boolean if a field has been set.

### GetDdosChallengesIssued

`func (o *PlatformValues) GetDdosChallengesIssued() int32`

GetDdosChallengesIssued returns the DdosChallengesIssued field if non-nil, zero value otherwise.

### GetDdosChallengesIssuedOk

`func (o *PlatformValues) GetDdosChallengesIssuedOk() (*int32, bool)`

GetDdosChallengesIssuedOk returns a tuple with the DdosChallengesIssued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdosChallengesIssued

`func (o *PlatformValues) SetDdosChallengesIssued(v int32)`

SetDdosChallengesIssued sets DdosChallengesIssued field to given value.

### HasDdosChallengesIssued

`func (o *PlatformValues) HasDdosChallengesIssued() bool`

HasDdosChallengesIssued returns a boolean if a field has been set.

### GetDdosChallengesSucceeded

`func (o *PlatformValues) GetDdosChallengesSucceeded() int32`

GetDdosChallengesSucceeded returns the DdosChallengesSucceeded field if non-nil, zero value otherwise.

### GetDdosChallengesSucceededOk

`func (o *PlatformValues) GetDdosChallengesSucceededOk() (*int32, bool)`

GetDdosChallengesSucceededOk returns a tuple with the DdosChallengesSucceeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdosChallengesSucceeded

`func (o *PlatformValues) SetDdosChallengesSucceeded(v int32)`

SetDdosChallengesSucceeded sets DdosChallengesSucceeded field to given value.

### HasDdosChallengesSucceeded

`func (o *PlatformValues) HasDdosChallengesSucceeded() bool`

HasDdosChallengesSucceeded returns a boolean if a field has been set.

### GetDdosChallengesFailed

`func (o *PlatformValues) GetDdosChallengesFailed() int32`

GetDdosChallengesFailed returns the DdosChallengesFailed field if non-nil, zero value otherwise.

### GetDdosChallengesFailedOk

`func (o *PlatformValues) GetDdosChallengesFailedOk() (*int32, bool)`

GetDdosChallengesFailedOk returns a tuple with the DdosChallengesFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdosChallengesFailed

`func (o *PlatformValues) SetDdosChallengesFailed(v int32)`

SetDdosChallengesFailed sets DdosChallengesFailed field to given value.

### HasDdosChallengesFailed

`func (o *PlatformValues) HasDdosChallengesFailed() bool`

HasDdosChallengesFailed returns a boolean if a field has been set.

### GetDdosChallengesPatsIssued

`func (o *PlatformValues) GetDdosChallengesPatsIssued() int32`

GetDdosChallengesPatsIssued returns the DdosChallengesPatsIssued field if non-nil, zero value otherwise.

### GetDdosChallengesPatsIssuedOk

`func (o *PlatformValues) GetDdosChallengesPatsIssuedOk() (*int32, bool)`

GetDdosChallengesPatsIssuedOk returns a tuple with the DdosChallengesPatsIssued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdosChallengesPatsIssued

`func (o *PlatformValues) SetDdosChallengesPatsIssued(v int32)`

SetDdosChallengesPatsIssued sets DdosChallengesPatsIssued field to given value.

### HasDdosChallengesPatsIssued

`func (o *PlatformValues) HasDdosChallengesPatsIssued() bool`

HasDdosChallengesPatsIssued returns a boolean if a field has been set.

### GetDdosChallengesPatsSucceeded

`func (o *PlatformValues) GetDdosChallengesPatsSucceeded() int32`

GetDdosChallengesPatsSucceeded returns the DdosChallengesPatsSucceeded field if non-nil, zero value otherwise.

### GetDdosChallengesPatsSucceededOk

`func (o *PlatformValues) GetDdosChallengesPatsSucceededOk() (*int32, bool)`

GetDdosChallengesPatsSucceededOk returns a tuple with the DdosChallengesPatsSucceeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdosChallengesPatsSucceeded

`func (o *PlatformValues) SetDdosChallengesPatsSucceeded(v int32)`

SetDdosChallengesPatsSucceeded sets DdosChallengesPatsSucceeded field to given value.

### HasDdosChallengesPatsSucceeded

`func (o *PlatformValues) HasDdosChallengesPatsSucceeded() bool`

HasDdosChallengesPatsSucceeded returns a boolean if a field has been set.

### GetCspPagesGauge

`func (o *PlatformValues) GetCspPagesGauge() int64`

GetCspPagesGauge returns the CspPagesGauge field if non-nil, zero value otherwise.

### GetCspPagesGaugeOk

`func (o *PlatformValues) GetCspPagesGaugeOk() (*int64, bool)`

GetCspPagesGaugeOk returns a tuple with the CspPagesGauge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspPagesGauge

`func (o *PlatformValues) SetCspPagesGauge(v int64)`

SetCspPagesGauge sets CspPagesGauge field to given value.

### HasCspPagesGauge

`func (o *PlatformValues) HasCspPagesGauge() bool`

HasCspPagesGauge returns a boolean if a field has been set.

### GetNgwafAccountRulesGauge

`func (o *PlatformValues) GetNgwafAccountRulesGauge() int64`

GetNgwafAccountRulesGauge returns the NgwafAccountRulesGauge field if non-nil, zero value otherwise.

### GetNgwafAccountRulesGaugeOk

`func (o *PlatformValues) GetNgwafAccountRulesGaugeOk() (*int64, bool)`

GetNgwafAccountRulesGaugeOk returns a tuple with the NgwafAccountRulesGauge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNgwafAccountRulesGauge

`func (o *PlatformValues) SetNgwafAccountRulesGauge(v int64)`

SetNgwafAccountRulesGauge sets NgwafAccountRulesGauge field to given value.

### HasNgwafAccountRulesGauge

`func (o *PlatformValues) HasNgwafAccountRulesGauge() bool`

HasNgwafAccountRulesGauge returns a boolean if a field has been set.

### GetNgwafAccountWorkspaces

`func (o *PlatformValues) GetNgwafAccountWorkspaces() int64`

GetNgwafAccountWorkspaces returns the NgwafAccountWorkspaces field if non-nil, zero value otherwise.

### GetNgwafAccountWorkspacesOk

`func (o *PlatformValues) GetNgwafAccountWorkspacesOk() (*int64, bool)`

GetNgwafAccountWorkspacesOk returns a tuple with the NgwafAccountWorkspaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNgwafAccountWorkspaces

`func (o *PlatformValues) SetNgwafAccountWorkspaces(v int64)`

SetNgwafAccountWorkspaces sets NgwafAccountWorkspaces field to given value.

### HasNgwafAccountWorkspaces

`func (o *PlatformValues) HasNgwafAccountWorkspaces() bool`

HasNgwafAccountWorkspaces returns a boolean if a field has been set.

### GetNgwafWorkspaceRulesGauge

`func (o *PlatformValues) GetNgwafWorkspaceRulesGauge() int64`

GetNgwafWorkspaceRulesGauge returns the NgwafWorkspaceRulesGauge field if non-nil, zero value otherwise.

### GetNgwafWorkspaceRulesGaugeOk

`func (o *PlatformValues) GetNgwafWorkspaceRulesGaugeOk() (*int64, bool)`

GetNgwafWorkspaceRulesGaugeOk returns a tuple with the NgwafWorkspaceRulesGauge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNgwafWorkspaceRulesGauge

`func (o *PlatformValues) SetNgwafWorkspaceRulesGauge(v int64)`

SetNgwafWorkspaceRulesGauge sets NgwafWorkspaceRulesGauge field to given value.

### HasNgwafWorkspaceRulesGauge

`func (o *PlatformValues) HasNgwafWorkspaceRulesGauge() bool`

HasNgwafWorkspaceRulesGauge returns a boolean if a field has been set.

### GetNgwafWorkspaceThresholdsGauge

`func (o *PlatformValues) GetNgwafWorkspaceThresholdsGauge() int64`

GetNgwafWorkspaceThresholdsGauge returns the NgwafWorkspaceThresholdsGauge field if non-nil, zero value otherwise.

### GetNgwafWorkspaceThresholdsGaugeOk

`func (o *PlatformValues) GetNgwafWorkspaceThresholdsGaugeOk() (*int64, bool)`

GetNgwafWorkspaceThresholdsGaugeOk returns a tuple with the NgwafWorkspaceThresholdsGauge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNgwafWorkspaceThresholdsGauge

`func (o *PlatformValues) SetNgwafWorkspaceThresholdsGauge(v int64)`

SetNgwafWorkspaceThresholdsGauge sets NgwafWorkspaceThresholdsGauge field to given value.

### HasNgwafWorkspaceThresholdsGauge

`func (o *PlatformValues) HasNgwafWorkspaceThresholdsGauge() bool`

HasNgwafWorkspaceThresholdsGauge returns a boolean if a field has been set.

### GetObjectStorageBucketsCount

`func (o *PlatformValues) GetObjectStorageBucketsCount() int64`

GetObjectStorageBucketsCount returns the ObjectStorageBucketsCount field if non-nil, zero value otherwise.

### GetObjectStorageBucketsCountOk

`func (o *PlatformValues) GetObjectStorageBucketsCountOk() (*int64, bool)`

GetObjectStorageBucketsCountOk returns a tuple with the ObjectStorageBucketsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectStorageBucketsCount

`func (o *PlatformValues) SetObjectStorageBucketsCount(v int64)`

SetObjectStorageBucketsCount sets ObjectStorageBucketsCount field to given value.

### HasObjectStorageBucketsCount

`func (o *PlatformValues) HasObjectStorageBucketsCount() bool`

HasObjectStorageBucketsCount returns a boolean if a field has been set.

### GetObjectStorageBytesCount

`func (o *PlatformValues) GetObjectStorageBytesCount() int64`

GetObjectStorageBytesCount returns the ObjectStorageBytesCount field if non-nil, zero value otherwise.

### GetObjectStorageBytesCountOk

`func (o *PlatformValues) GetObjectStorageBytesCountOk() (*int64, bool)`

GetObjectStorageBytesCountOk returns a tuple with the ObjectStorageBytesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectStorageBytesCount

`func (o *PlatformValues) SetObjectStorageBytesCount(v int64)`

SetObjectStorageBytesCount sets ObjectStorageBytesCount field to given value.

### HasObjectStorageBytesCount

`func (o *PlatformValues) HasObjectStorageBytesCount() bool`

HasObjectStorageBytesCount returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
