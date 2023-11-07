# Changelog

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

* Object Store API [commit](https://github.com/fastly/fastly-go/commit/e69498474f02c2208072160821a0d49c6999087d) 

## [v1.0.0-beta.2](https://github.com/fastly/fastly-go/releases/tag/v1.0.0-beta.2) (2023-01-18)

**Bug fixes:**

* Fixed OpenAPI schemas to produce missing methods for updating service settings [commit](https://github.com/fastly/fastly-go/commit/4c0423bfccbb4f62cb90f894f630b26306ffdc1a) 

## [v1.0.0-beta.1](https://github.com/fastly/fastly-go/releases/tag/v1.0.0-beta.1) (2023-01-17)

**Enhancements:**

* Update service settings API endpoint [commit](https://github.com/fastly/fastly-go/commit/0c29f6af943304085de0c999e45407e151600e3a) 
* Config Store API endpoints [commit](https://github.com/fastly/fastly-go/commit/0c29f6af943304085de0c999e45407e151600e3a) 

## [v1.0.0-beta.0](https://github.com/fastly/fastly-go/releases/tag/v1.0.0-beta.0) (2022-12-15)

**Enhancements:**

* New interface from code-generated API client [commit](https://github.com/fastly/fastly-go/commit/6b36bdea0aacc79321a1a970c57f0a31ca09ca45) 
  * [Blog post: Better Fastly API clients with OpenAPI Generator](https://dev.to/fastly/better-fastly-api-clients-with-openapi-generator-3lno)
  * [Documentation](https://github.com/fastly/fastly-go#documentation-for-api-endpoints)
  * [Unsupported API endpoints](https://github.com/fastly/fastly-go#issues)
