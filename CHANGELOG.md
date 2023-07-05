# Changelog

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
