# Changelog

## [v1.0.0-beta.39](https://github.com/fastly/fastly-go/releases/tag/release/v1.0.0-beta.39) (2025-02-28)

**Breaking Changes:**


- fix(snippet): Renamed `snippet_id` and `snippet_name` properties to `id` and `name` respectively.
- feat(kv_store, kv_store_item): KV store endpoints have new operation names, and the associated data models also have
                                 new names.
- feat(product): Product enablement and configuration endpoints have been renamed into per-product operations.
- fix(historical): Widened integer stats fields to 64 bits.
- fix(realtime, domain_inspector): Widened integer stats fields to 64 bits.
- fix(historical): Corrected field names of `ddos_protection_requests_detect_count`,
                   `ddos_protection_requests_mitigate_count`, and `ddos_protection_requests_allow_count`.
- fix(realtime): Corrected field names of `ddos_protection_requests_detect_count`,
                 `ddos_protection_requests_mitigate_count`, and `ddos_protection_requests_allow_count`.
- fix(billing_usage_metrics): added parameters `req_start_month` and `req_end_month`.
- deprecated(billing): Billing v2 API has been deprecated.
- removed(legacy_waf): Legacy WAF API has been removed.

**Bug fixes:**

- fix(snippet): Marked `content` as nullable to support dynamic snippets.
- fix(billing_usage_metrics): `product_id` and `usage_type_name` parameters are no longer required.

**Enhancements:**

- feat(products): Added `object_storage` and `ai_accelerator` products.
- feat(historical): Added new metrics: `object_storage_class_a_operations_count`,
                    `object_storage_class_b_operations_count`, `aia_requests`, `aia_status_1xx`,
                    `aia_status_2xx`, `aia_status_3xx`, `aia_status_5xx`, `aia_response_usage_tokens`,
                    `aia_origin_usage_tokens`, `aia_estimated_time_saved_ms`, `request_collapse_usable_count`,
                    `request_collapse_unusable_count`, `status_530`, and `compute_cache_operations_count`.
- feat(realtime): Added new metrics: `object_storage_class_a_operations_count`,
                  `object_storage_class_b_operations_count`, `aia_requests`, `aia_status_1xx`,
                  `aia_status_2xx`, `aia_status_3xx`, `aia_status_5xx`, `aia_response_usage_tokens`,
                  `aia_origin_usage_tokens`, `aia_estimated_time_saved_ms`, `request_collapse_usable_count`,
                  `request_collapse_unusable_count`, `status_530`, and `compute_cache_operations_count`.
- feat(object_storage_access_keys): Added Object Storage Access Keys API.
- feat(domain_inspector_historical, domain_inspector_realtime): Added new metrics: `status_530` and `origin_status_530`.
- feat(origin_inspector_historical, origin_inspector_realtime): Added new metrics: `status_530`, `waf_status_530`,
                                                                `compute_status_530`, and `all_status_530`.
- feat(kv_store_item): Added support for `if-generation-match`.
- feat(acls_in_compute): Added ACLs in Compute API.

**Documentation:**

- doc(billing_usage_metrics): Documented that `start_month` and `end_month` are required.
- doc(historical): Correct default start time in description of Historical Stats API.
- doc(domain_inspector_historical): Removed duplicate description.
- doc(domain_inspector_realtime): Removed duplicate description.
- doc(origin_inspector_historical): Removed duplicate description.
- doc(origin_inspector_realtime): Removed duplicate description.
- doc(backend): Updated description of Backend API.
- doc(backend): Clarified default behavior of keepalive_time.
- doc(waf): Updated EOL notice of legacy-waf.

## [v1.0.0-beta.38](https://github.com/fastly/fastly-go/releases/tag/release/v1.0.0-beta.38) (2024-12-17)

**Breaking Changes:**

- fix(domain-inspector-realtime, origin-inspector-realtime, historical, realtime): Stats fields widened to 64 bits.

## [v1.0.0-beta.37](https://github.com/fastly/fastly-go/releases/tag/release/v1.0.0-beta.37) (2024-12-05)

**Breaking Changes:**

- fix(udm-domains): Use v1 versioned HTTP endpoints for UDM domains
- fix(historical): Remove references to VCL on Compute metrics
- fix(realtime): Remove references to VCL on Compute metrics
- feat(billing-usage-metrics): Updated to v3 of the API

**Bug fixes:**

- fix(acls-in-compute): Corrected shape of `compute-acl-update` model.

**Enhancements:**

- feat(object-storage-access-keys): Added Object Storage Access Keys API
- feat(logging-grafanacloudlogs): Added Grafana Cloud Logs Loggin API
- feat(insights): Added Insights API
- feat(kv-store-item): Added `consistency` parameter to `get-keys` operation.
- feat(enabled-products): Added `mode` property to `set-configuration`.

**Documentation:**

- doc(observability-custom-dashboards): Fix bad link in custom dashboard documentation
- doc(observability-custom-dashboards): Reformat some enums to fix docs rendering
- doc(acls-in-compute): Corrections to descriptions of `compute-acl-lookup`, `compute-acl-list-entries`,
                        `compute-acl-update-entry`, and `compute-acl-update` structures
- doc(enabled-products): Added support for product ID `log_explorer_insights`.
- doc(enabled-products): Added support for product ID `ddos_protection`.

## [v1.0.0-beta.36](https://github.com/fastly/fastly-go/releases/tag/release/v1.0.0-beta.36) (2024-10-23)

**Enhancements:**

- feat(go): Run 'gofmt' on generated Go code.

**Documentation:**

- doc(backend): Correct spelling in `connect_timeout` and `first_byte_timeout` field descriptions.

## [v1.0.0-beta.35](https://github.com/fastly/fastly-go/releases/tag/release/v1.0.0-beta.35) (2024-09-13)

**Bug fixes:**

- bugfix(rust): Implement std::fmt::Display instead of std::string::ToString for enum models.
- fix(contact): Corrected endpoint called by `delete-contact` operation
- fix(origin-inspector, domain-inspector): Use integer type for Timestamp.
- fix(billing): Make all fields on billing list item data nullable

**Enhancements:**

- feat(acls-in-compute): Add ACLs in Compute API
- feat(enabled-products): Updated to use `v1` versioned endpoints.
- feat(enabled-products): Added `get-product-configuration`, `set-product-configuration` operations.
- feat(realtime, historical): Added `request_denied_get_head_body` metric.
- feat(tls): Add definitions for values of sort parameter
- feat(staging): Add activate/deactivate endpoints for staging
- feat(customer-addresses): Add Customer Addresses API

**Documentation:**

- doc(enabled-products): Added support for product IDs `bot_management` and `ngwaf`.

## [v1.0.0-beta.34](https://github.com/fastly/fastly-go/releases/tag/release/v1.0.0-beta.34) (2024-08-21)

**Bug fixes:**

- fix(billing): make rate-per-unit nullable

## [v1.0.0-beta.33](https://github.com/fastly/fastly-go/releases/tag/release/v1.0.0-beta.33) (2024-08-08)

**Bug fixes:**

- fix(billing): Adjust type of regional data to help the generator
- fix(billing): Correct type of invoice_id field

## [v1.0.0-beta.32](https://github.com/fastly/fastly-go/releases/tag/release/v1.0.0-beta.32) (2024-08-04)

**Enhancements:**

- feat(observability): Adds new Observability Custom Dashboards API
- feat(billing-invoices): Adds month-to-date invoice information
- feat(billing-usage-metrics): Adds information on service-level usage.
- feat(tls): Adds an endpoint to get a TLS certificate blob (Limited Availability)
- feat(stats): Add `origin_offload` metric

## [v1.0.0-beta.31](https://github.com/fastly/fastly-go/releases/tag/v1.0.0-beta.31) (2024-07-01)

**Bug fixes:**

- fix(logging): For several endpoints, correct use_tls to be string

**Enhancements:**

- feat(content): `/content/edge_check` endpoint now returns informational values in `hash` when a timeout occurs or when
  an object is too large.
- feat(logging-datalog): Added additional regions

## [v1.0.0-beta.30](https://github.com/fastly/fastly-go/releases/tag/v1.0.0-beta.30) (2024-06-27)

**Bug fixes:**

- bugfix(alerts-definitions): For Origin derived metrics, correct `all_bandwidth` type to `integer`

**Enhancements:**

- feat(logging-s3): Add `file_max_bytes` configuration field
- feat(alerts-definitions): Add `integration_id` parameter to the List Alert Definitions endpoint
- feat(alerts-definitions): For Origin derived metrics, add `all_status_4xx_excl_404_rate` and `all_status_404_rate` properties
- feat(alerts-definitions): For Domain derived metrics, add `status_4xx_excl_404_rate` and `status_404_rate` properties
- feat(alerts-definitions): For Stats derived metrics, add `status_4xx_excl_404_rate`, `status_404_rate`, `all_status_5xx_rate`,
                            `all_status_4xx_rate`, `all_status_gte_400_rate`, and `all_status_lt_500_rate` properties
- feat(billing-invoices): For invoice line items, added `ProductLine` property

**Documentation:**

- doc(billing-invoices): "Billing Invoices API" relabeled to "Invoices API"
- doc(billing-invoices): Updated documentation of `billing_start_date` and `billing_end_date` query parameters of
                         List of invoices endpoint
- doc(alerts-definitions): Updated documentation of several fields
- doc(notification-service): Updated documentation examples for several fields

## [v1.0.0-beta.29](https://github.com/fastly/fastly-go/releases/tag/v1.0.0-beta.29) (2024-05-24)

**Bug fixes:**

- fix(backend): Correct tcp_keepalive_enable to be nullable
- bugfix(resource): Correct `resourceTyperesource` accepted values
- bugfix(sudo): Mark API to require authentication token

**Enhancements:**

- feat(image-optimizer-default-settings): Add Image Optimizer Default Settings APIs
- feat(tls-subscriptions): Add `certificate_authority` filter parameter

**Documentation:**

- docs(generator): Update links in docs and comments with unified docs site URL structure
- docs(generator): Clean up README by using GitHub alert icons and collapsible sections
- doc(logging-kinesis): Update description of `format` field.
- doc(tls-subscriptions): Update descriptions of `certificate_authority` and `tls_subscription_include` fields.

## [v1.0.0-beta.28](https://github.com/fastly/fastly-go/releases/tag/v1.0.0-beta.28) (2024-04-25)

**Bug fixes:**

- bugfix(billing_address, invitations): Correct customer relationship schema to be single entry rather than array
- bugfix(request_settings): Mark request_settings fields as nullable: bypass_busy_wait, force_miss, geo_headers,
  max_stale_age, timer_support, and xff

**Enhancements:**

- feat(generator): The API Client Generator now uses a new automatic changelog generation process.
- feat(historical, realtime): Add the following new metrics: `ddos_action_downgrade`,
  `ddos_action_downgraded_connections`, `vcl_on_compute_hit_requests`, `vcl_on_compute_miss_requests`,
  `vcl_on_compute_pass_requests`, `vcl_on_compute_error_requests`, `vcl_on_compute_synth_requests`,
  `vcl_on_compute_edge_hit_requests`, `vcl_on_compute_edge_miss_requests`, `all_hit_requests`,
  `all_miss_requests`, `all_pass_requests`, `all_error_requests`, `all_synth_requests`, `all_edge_hit_requests`,
  `all_edge_miss_requests`, `all_status_1xx`, `all_status_2xx`, `all_status_3xx`, `all_status_4xx`, and
  `all_status_5xx`.
- feat(backend): Add `tcp_keepalive_*` properties to the Backend API, which allow configuring TCP keepalives for
  backend connections.

**Documentation:**

- doc(realtime): Correct description of miss_histogram structure in real-time stats.

## [v1.0.0-beta.27](https://github.com/fastly/fastly-go/releases/tag/v1.0.0-beta.27) (2024-02-28)

**Enhancements:**

- feat(billing): add 'get invoice by invoice ID' endpoint.
- feat(billing): add 'list invoices' endpoint.
- feat(image_opto): add 'shield' and 'transform' methods and fields.
- feat(secret_store): add 'name' filter.
- feat(tls): add 'create CSR' endpoint.
- feat(tls_subscription): add 'active order' methods and fields.

**Documentation:**

- doc(backend) fix openssl link.
- doc(director,domain_inspector,pool,realtime) fix POPs link.
- doc(service) fix service version link.
- doc(tls) fix activation link.

## [v1.0.0-beta.26](https://github.com/fastly/fastly-go/releases/tag/v1.0.0-beta.26) (2024-02-19)

**Bug fixes:**

- fix(response_object): strongly type response_object create_update requests
- fix(tls_configurations): fix `tls_protocols` field to be a string array type

## [v1.0.0-beta.25](https://github.com/fastly/fastly-go/releases/tag/v1.0.0-beta.25) (2023-11-29)

**Breaking:**

- breaking(historical): restructure OpenAPI schema to avoid duplicated data rendering.

**Bug fixes:**

- fix: no longer display Markdown syntax in code comments.

**Enhancements:**

- feat(stats): expose new `service_id` and `start_time` fields.

## [v1.0.0-beta.24](https://github.com/fastly/fastly-go/releases/tag/v1.0.0-beta.24) (2023-11-07)

**Enhancements:**

- feat(config_store): add `name` query param to list endpoint.

## [v1.0.0-beta.23](https://github.com/fastly/fastly-go/releases/tag/v1.0.0-beta.23) (2023-10-27)

**Documentation:**

- docs: rename Compute@Edge to Compute.

## [v1.0.0-beta.22](https://github.com/fastly/fastly-go/releases/tag/v1.0.0-beta.22) (2023-10-24)

**Enhancements:**

- feat(stats): add historical DDoS metrics.
- feat(stats): add bot challenges.

**Bug fixes:**

- fix(snippets): ensure POST response's dynamic field is numerical.

## [v1.0.0-beta.21](https://github.com/fastly/fastly-go/releases/tag/v1.0.0-beta.21) (2023-09-01)

**Enhancements:**

- feat(events): support extra created_at filters.

## [v1.0.0-beta.20](https://github.com/fastly/fastly-go/releases/tag/v1.0.0-beta.20) (2023-09-01)

**Enhancements:**

- feat(backend): support share_key field.
- feat(events): support extra created_at filters.
- feat(logging/newrelic): add OTLP endpoints.
- feat(tls/subscriptions): support self_managed_http_challenge field.

**Documentation:**

- doc(secretstore): correct description for GET endpoint.

## [v1.0.0-beta.19](https://github.com/fastly/fastly-go/releases/tag/v1.0.0-beta.19) (2023-07-31)

**Breaking:**

The following restructures have helped resolve some issues with our OpenAPI schemas but as a side-effect this has resulted in a break to our API client interface as different types are now being generated.

- refactor: general restructure OpenAPI schemas.
- refactor(domain): remove explicit schema type for 'any'.

**Bug fixes:**

- fix: change response `version` type to string.
- fix(cache_settings): change response `stale_ttl` and `ttl` types to strings.
- fix(header): change response `ignore_if_set` and `priority` types to strings.
- fix(logging): change response `period` and `gzip_level` types to strings.
- fix(pool): change response `use_tls`, `max_conn_default`, `first_byte_timeout`, `quorum` and `tls_check_cert` types to strings.
- fix(request_settings): change response `bypass_busy_wait`, `force_miss`, `force_ssl`, `geo_headers`, `max_stale_age` and `timer_support` types to strings.
- fix(response_object): change response `status` type to string.

## [v1.0.0-beta.18](https://github.com/fastly/fastly-go/releases/tag/v1.0.0-beta.18) (2023-07-13)

**Bug fixes:**

- fix(logging_gcs): set expected default value for 'path'.
- fix(origin_inspector_historical): use correct type for 'values'.
- fix(tls_subscriptions): fix argument order for deleting globalsign email challenge.

## [v1.0.0-beta.17](https://github.com/fastly/fastly-go/releases/tag/v1.0.0-beta.17) (2023-07-12)

**Bug fixes:**

- fix(billing): rename response field 'lines' to 'line_items'.
- fix(billing): restructure response models like 'aria_invoice_id'.
- fix(billing): make 'sent_at', 'locked', 'require_new_password', 'two_factor_auth_enabled' nullable.

## [v1.0.0-beta.16](https://github.com/fastly/fastly-go/releases/tag/v1.0.0-beta.16) (2023-07-06)

Substantial changes were made to the underlying OpenAPI specification that produces this API client. These changes have resulted in multiple new endpoints being supported as well as multiple breaking type changes and so we're publishing these changes as a new major release.

**Enhancements:**

- feat(apex_redirect): support all endpoints.
- feat(contact): support 'create' endpoint.
- feat(director): support 'update' endpoint.
- feat(domain_inspector): support all endpoints.
- feat(iam_roles): support 'add permissions' endpoint.
- feat(iam_roles): support 'create role' endpoint.
- feat(iam_roles): support 'delete permissions' endpoint.
- feat(iam_roles): support 'update role' endpoint.
- feat(iam_services): support 'add services' endpoint.
- feat(iam_services): support 'create service group' endpoint.
- feat(iam_services): support 'remove services' endpoint.
- feat(iam_services): support 'update service group' endpoint.
- feat(iam_users): support 'add members' endpoint.
- feat(iam_users): support 'add roles' endpoint.
- feat(iam_users): support 'add service groups' endpoint.
- feat(iam_users): support 'create user group' endpoint.
- feat(iam_users): support 'remove members' endpoint.
- feat(iam_users): support 'remove roles' endpoint.
- feat(iam_users): support 'remove service groups' endpoint.
- feat(iam_users): support 'update user group' endpoint.
- feat(legacy_waf): support all endpoints.
- feat(logging_kafka): support 'update' endpoint.
- feat(logging_kinesis): support 'update' endpoint.
- feat(origin_inspector): support all endpoints.
- feat(request_settings): support 'create' endpoint.
- feat(response_object): support 'create' endpoint.
- feat(response_object): support 'update' endpoint.
- feat(secret_store): support all endpoints.
- feat(service_authorizations): support 'delete' endpoint.
- feat(service_authorizations): support 'update' endpoint.
- feat(snippet): support 'update versioned snippet' endpoint.
- feat(sudo): support 'request sudo access' endpoint.
- feat(tokens): support 'revoke multiple tokens' endpoint.
- feat(tokens): support 'create token' endpoint.
- feat(waf_active_rules): support 'delete' endpoint.

**Bug fixes:**

- fix(content): update request/response types.
- fix(events): update metadata type.
- fix(realtime_entry): update recorded/aggregated type.
- fix(realtime_measurements): update miss_histogram type.

## [v1.0.0-beta.15](https://github.com/fastly/fastly-go/releases/tag/v1.0.0-beta.15) (2023-07-05)

**Enhancements:**

- feat(purge): support purge of multiple surrogate keys.
- feat(vcl): support vcl content endpoints.
- feat(vcl): support vcl linting endpoints.

**Bug fixes:**

- fix(snippet): dynamic field switched from int to string.

**Documentation:**

- docs: remove deprecated docs endpoints from README 'issues' list.

## [v1.0.0-beta.14](https://github.com/fastly/fastly-go/releases/tag/v1.0.0-beta.14) (2023-06-27)

**Enhancements:**

- feat(automation_tokens): implement endpoints.
- feat(rate_limiter): implement POST/PUT endpoints.

**Bug fixes:**

- fix(historical_stats): extract primitive into custom type.

## [v1.0.0-beta.13](https://github.com/fastly/fastly-go/releases/tag/v1.0.0-beta.13) (2023-06-23)

**Bug fixes:**

- fix(historical_stats): generate missing models.

## [v1.0.0-beta.12](https://github.com/fastly/fastly-go/releases/tag/v1.0.0-beta.12) (2023-06-21)

**Bug fixes:**

- fix(tls_activation): add tls_configuration and tls_domains.
- fix(tls_subscription): add tls_configuration and common name.

## [v1.0.0-beta.11](https://github.com/fastly/fastly-go/releases/tag/v1.0.0-beta.11) (2023-06-20)

**Enhancements:**

- feat(realtime_measurements): add billable request processing time.
- feat(tokens): add support for the 'get token' endpoint.

**Bug fixes:**

- fix(config): add realtime hostname.
- fix(historical_stats): generate field results model.
- fix(kv_store): remove the 'force' property from the 'delete store' endpoint.
- feat(realtime_measurements): rename object store to kv store.

## [v1.0.0-beta.10](https://github.com/fastly/fastly-go/releases/tag/v1.0.0-beta.10) (2023-05-22)

**Bug fixes:**

- fix(acl): change `version` from int to string.
- fix(acl): add missing methods for `service_id` and `service_version` properties.
- fix(backend): make `ssl_check_cert` nullable.
- fix(purge): skip URL escape for `surrogate_key` param.
- fix(snippets): change `priority` and `version` from int to string.
- fix(snippets): add missing methods for `service_id` and `service_version` properties.

## [v1.0.0-beta.9](https://github.com/fastly/fastly-go/releases/tag/v1.0.0-beta.9) (2023-05-17)

**Enhancements:**

- feat(config_store): add Config Store endpoints.

## [v1.0.0-beta.8](https://github.com/fastly/fastly-go/releases/tag/v1.0.0-beta.8) (2023-05-16)

**Breaking changes:**

- breaking(object_store): rename to kv_store

**Enhancements:**

- feat(dictionary_item): add 'bulk' PATCH endpoint.
- feat(package): add `files_hash` metadata property.
- feat(tls_certificates): add `filter[in_use]` parameter.

## [v1.0.0-beta.7](https://github.com/fastly/fastly-go/releases/tag/v1.0.0-beta.7) (2023-04-26)

**Bug fixes:**

- fix(object-store-item): use correct type for key value.
- fix(tls-csrs): remove internal endpoint.

## [v1.0.0-beta.6](https://github.com/fastly/fastly-go/releases/tag/v1.0.0-beta.6) (2023-04-04)

**Bug fixes:**

- fix(purge): avoid PathEscape with x-allow-reserved

## [v1.0.0-beta.5](https://github.com/fastly/fastly-go/releases/tag/v1.0.0-beta.5) (2023-04-03)

**Enhancements:**

- feat(object-store): add 'location' property to 'create_store'
- feat(object-store): add 'force' property to 'delete_store'
- feat(realtime): additional DDoS properties

**Documentation:**

- docs(acl-entries): document batch updating
- docs(logentries): deprecation notice
- docs(object-store): new properties for 'SetValueForKey'

## [v1.0.0-beta.4](https://github.com/fastly/fastly-go/releases/tag/v1.0.0-beta.4) (2023-03-22)

**Bug fixes:**

- fix(purge): switch authentication type to 'token'

**Enhancements:**

- feat(domain-ownerships): List domain-ownerships
- feat(events): implement 'filter_created_at' property

**Documentation:**

- docs(backend): keepalive_time
- docs(object-store): restructure of the API documentation
- docs(pop): region, shield, latitude, longitude
- docs(product-enablement): brotli_compression
- docs(resource): terminology
- docs(results): fanout properties
- docs(tls/subscriptions): new 'failed' state
- docs(user): 'login' modification note removed

## [v1.0.0-beta.3](https://github.com/fastly/fastly-go/releases/tag/v1.0.0-beta.3) (2023-01-20)

**Enhancements:**

- Object Store API [commit](https://github.com/fastly/fastly-go/commit/e69498474f02c2208072160821a0d49c6999087d)

## [v1.0.0-beta.2](https://github.com/fastly/fastly-go/releases/tag/v1.0.0-beta.2) (2023-01-18)

**Bug fixes:**

- Fixed OpenAPI schemas to produce missing methods for updating service settings [commit](https://github.com/fastly/fastly-go/commit/4c0423bfccbb4f62cb90f894f630b26306ffdc1a)

## [v1.0.0-beta.1](https://github.com/fastly/fastly-go/releases/tag/v1.0.0-beta.1) (2023-01-17)

**Enhancements:**

- Update service settings API endpoint [commit](https://github.com/fastly/fastly-go/commit/0c29f6af943304085de0c999e45407e151600e3a)
- Config Store API endpoints [commit](https://github.com/fastly/fastly-go/commit/0c29f6af943304085de0c999e45407e151600e3a)

## [v1.0.0-beta.0](https://github.com/fastly/fastly-go/releases/tag/v1.0.0-beta.0) (2022-12-15)

**Enhancements:**

- New interface from code-generated API client [commit](https://github.com/fastly/fastly-go/commit/6b36bdea0aacc79321a1a970c57f0a31ca09ca45)
  - [Blog post: Better Fastly API clients with OpenAPI Generator](https://dev.to/fastly/better-fastly-api-clients-with-openapi-generator-3lno)
  - [Documentation](https://github.com/fastly/fastly-go#documentation-for-api-endpoints)
  - [Unsupported API endpoints](https://github.com/fastly/fastly-go#issues)
