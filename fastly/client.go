// Package fastly is an API client library for interacting with most facets of the Fastly API.
package fastly

/*
Fastly API

Via the Fastly API you can perform any of the operations that are possible within the management console,  including creating services, domains, and backends, configuring rules or uploading your own application code, as well as account operations such as user administration and billing reports. The API is organized into collections of endpoints that allow manipulation of objects related to Fastly services and accounts. For the most accurate and up-to-date API reference content, visit our [Developer Hub](https://developer.fastly.com/reference/api/) 

API version: 1.0.0
Contact: oss@fastly.com
*/

// This code is auto-generated; DO NOT EDIT.


import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httputil"
	gourl "net/url"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"golang.org/x/oauth2"
)

// https://developer.fastly.com/reference/api/#rate-limiting
const defaultRateLimit = 1000

var (
	jsonCheck = regexp.MustCompile(`(?i:(?:application|text)/(?:vnd\.[^;]+\+)?json)`)
	xmlCheck  = regexp.MustCompile(`(?i:(?:application|text)/xml)`)
)

// APIClient manages communication with the Fastly API API v1.0.0
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	cfg    *Configuration
	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// RateLimitRemaining is the last observed value of http header 
	// Fastly-RateLimit-Remaining
	// https://developer.fastly.com/reference/api/#rate-limiting
	RateLimitRemaining int
	
	// RateLimitReset is the last observed value of http header 
	// Fastly-RateLimit-Reset
	RateLimitReset int

	// API Services

	ACLAPI ACLAPI

	ACLEntryAPI ACLEntryAPI

	ApexRedirectAPI ApexRedirectAPI

	BackendAPI BackendAPI

	BillingAPI BillingAPI

	BillingAddressAPI BillingAddressAPI

	CacheSettingsAPI CacheSettingsAPI

	ConditionAPI ConditionAPI

	ContactAPI ContactAPI

	ContentAPI ContentAPI

	CustomerAPI CustomerAPI

	DictionaryAPI DictionaryAPI

	DictionaryInfoAPI DictionaryInfoAPI

	DictionaryItemAPI DictionaryItemAPI

	DiffAPI DiffAPI

	DirectorAPI DirectorAPI

	DirectorBackendAPI DirectorBackendAPI

	DomainAPI DomainAPI

	EnabledProductsAPI EnabledProductsAPI

	EventsAPI EventsAPI

	GzipAPI GzipAPI

	HeaderAPI HeaderAPI

	HealthcheckAPI HealthcheckAPI

	HistoricalAPI HistoricalAPI

	HTTP3API HTTP3API

	IamPermissionsAPI IamPermissionsAPI

	IamRolesAPI IamRolesAPI

	IamServiceGroupsAPI IamServiceGroupsAPI

	IamUserGroupsAPI IamUserGroupsAPI

	InvitationsAPI InvitationsAPI

	LoggingAzureblobAPI LoggingAzureblobAPI

	LoggingBigqueryAPI LoggingBigqueryAPI

	LoggingCloudfilesAPI LoggingCloudfilesAPI

	LoggingDatadogAPI LoggingDatadogAPI

	LoggingDigitaloceanAPI LoggingDigitaloceanAPI

	LoggingElasticsearchAPI LoggingElasticsearchAPI

	LoggingFtpAPI LoggingFtpAPI

	LoggingGcsAPI LoggingGcsAPI

	LoggingHerokuAPI LoggingHerokuAPI

	LoggingHoneycombAPI LoggingHoneycombAPI

	LoggingHTTPSAPI LoggingHTTPSAPI

	LoggingKafkaAPI LoggingKafkaAPI

	LoggingKinesisAPI LoggingKinesisAPI

	LoggingLogentriesAPI LoggingLogentriesAPI

	LoggingLogglyAPI LoggingLogglyAPI

	LoggingLogshuttleAPI LoggingLogshuttleAPI

	LoggingNewrelicAPI LoggingNewrelicAPI

	LoggingOpenstackAPI LoggingOpenstackAPI

	LoggingPapertrailAPI LoggingPapertrailAPI

	LoggingPubsubAPI LoggingPubsubAPI

	LoggingS3API LoggingS3API

	LoggingScalyrAPI LoggingScalyrAPI

	LoggingSftpAPI LoggingSftpAPI

	LoggingSplunkAPI LoggingSplunkAPI

	LoggingSumologicAPI LoggingSumologicAPI

	LoggingSyslogAPI LoggingSyslogAPI

	MutualAuthenticationAPI MutualAuthenticationAPI

	ObjectStoreAPI ObjectStoreAPI

	PackageAPI PackageAPI

	PoolAPI PoolAPI

	PopAPI PopAPI

	PublicIPListAPI PublicIPListAPI

	PublishAPI PublishAPI

	PurgeAPI PurgeAPI

	RateLimiterAPI RateLimiterAPI

	RealtimeAPI RealtimeAPI

	RequestSettingsAPI RequestSettingsAPI

	ResourceAPI ResourceAPI

	ResponseObjectAPI ResponseObjectAPI

	ServerAPI ServerAPI

	ServiceAPI ServiceAPI

	ServiceAuthorizationsAPI ServiceAuthorizationsAPI

	SettingsAPI SettingsAPI

	SnippetAPI SnippetAPI

	StarAPI StarAPI

	StatsAPI StatsAPI

	TLSActivationsAPI TLSActivationsAPI

	TLSBulkCertificatesAPI TLSBulkCertificatesAPI

	TLSCertificatesAPI TLSCertificatesAPI

	TLSConfigurationsAPI TLSConfigurationsAPI

	TLSCsrsAPI TLSCsrsAPI

	TLSDomainsAPI TLSDomainsAPI

	TLSPrivateKeysAPI TLSPrivateKeysAPI

	TLSSubscriptionsAPI TLSSubscriptionsAPI

	TokensAPI TokensAPI

	UserAPI UserAPI

	VclAPI VclAPI

	VclDiffAPI VclDiffAPI

	VersionAPI VersionAPI

	WafActiveRulesAPI WafActiveRulesAPI

	WafExclusionsAPI WafExclusionsAPI

	WafFirewallVersionsAPI WafFirewallVersionsAPI

	WafFirewallsAPI WafFirewallsAPI

	WafRuleRevisionsAPI WafRuleRevisionsAPI

	WafRulesAPI WafRulesAPI

	WafTagsAPI WafTagsAPI
}

type service struct {
	client *APIClient
}

// NewAPIClient creates a new API client. Requires a userAgent string describing your application.
// optionally a custom http.Client to allow for advanced features such as caching.
func NewAPIClient(cfg *Configuration) *APIClient {
	if cfg.HTTPClient == nil {
		cfg.HTTPClient = http.DefaultClient
	}

	c := &APIClient{}
	c.cfg = cfg
	c.common.client = c
	c.RateLimitRemaining = defaultRateLimit

	// API Services
	c.ACLAPI = (*ACLAPIService)(&c.common)
	c.ACLEntryAPI = (*ACLEntryAPIService)(&c.common)
	c.ApexRedirectAPI = (*ApexRedirectAPIService)(&c.common)
	c.BackendAPI = (*BackendAPIService)(&c.common)
	c.BillingAPI = (*BillingAPIService)(&c.common)
	c.BillingAddressAPI = (*BillingAddressAPIService)(&c.common)
	c.CacheSettingsAPI = (*CacheSettingsAPIService)(&c.common)
	c.ConditionAPI = (*ConditionAPIService)(&c.common)
	c.ContactAPI = (*ContactAPIService)(&c.common)
	c.ContentAPI = (*ContentAPIService)(&c.common)
	c.CustomerAPI = (*CustomerAPIService)(&c.common)
	c.DictionaryAPI = (*DictionaryAPIService)(&c.common)
	c.DictionaryInfoAPI = (*DictionaryInfoAPIService)(&c.common)
	c.DictionaryItemAPI = (*DictionaryItemAPIService)(&c.common)
	c.DiffAPI = (*DiffAPIService)(&c.common)
	c.DirectorAPI = (*DirectorAPIService)(&c.common)
	c.DirectorBackendAPI = (*DirectorBackendAPIService)(&c.common)
	c.DomainAPI = (*DomainAPIService)(&c.common)
	c.EnabledProductsAPI = (*EnabledProductsAPIService)(&c.common)
	c.EventsAPI = (*EventsAPIService)(&c.common)
	c.GzipAPI = (*GzipAPIService)(&c.common)
	c.HeaderAPI = (*HeaderAPIService)(&c.common)
	c.HealthcheckAPI = (*HealthcheckAPIService)(&c.common)
	c.HistoricalAPI = (*HistoricalAPIService)(&c.common)
	c.HTTP3API = (*HTTP3APIService)(&c.common)
	c.IamPermissionsAPI = (*IamPermissionsAPIService)(&c.common)
	c.IamRolesAPI = (*IamRolesAPIService)(&c.common)
	c.IamServiceGroupsAPI = (*IamServiceGroupsAPIService)(&c.common)
	c.IamUserGroupsAPI = (*IamUserGroupsAPIService)(&c.common)
	c.InvitationsAPI = (*InvitationsAPIService)(&c.common)
	c.LoggingAzureblobAPI = (*LoggingAzureblobAPIService)(&c.common)
	c.LoggingBigqueryAPI = (*LoggingBigqueryAPIService)(&c.common)
	c.LoggingCloudfilesAPI = (*LoggingCloudfilesAPIService)(&c.common)
	c.LoggingDatadogAPI = (*LoggingDatadogAPIService)(&c.common)
	c.LoggingDigitaloceanAPI = (*LoggingDigitaloceanAPIService)(&c.common)
	c.LoggingElasticsearchAPI = (*LoggingElasticsearchAPIService)(&c.common)
	c.LoggingFtpAPI = (*LoggingFtpAPIService)(&c.common)
	c.LoggingGcsAPI = (*LoggingGcsAPIService)(&c.common)
	c.LoggingHerokuAPI = (*LoggingHerokuAPIService)(&c.common)
	c.LoggingHoneycombAPI = (*LoggingHoneycombAPIService)(&c.common)
	c.LoggingHTTPSAPI = (*LoggingHTTPSAPIService)(&c.common)
	c.LoggingKafkaAPI = (*LoggingKafkaAPIService)(&c.common)
	c.LoggingKinesisAPI = (*LoggingKinesisAPIService)(&c.common)
	c.LoggingLogentriesAPI = (*LoggingLogentriesAPIService)(&c.common)
	c.LoggingLogglyAPI = (*LoggingLogglyAPIService)(&c.common)
	c.LoggingLogshuttleAPI = (*LoggingLogshuttleAPIService)(&c.common)
	c.LoggingNewrelicAPI = (*LoggingNewrelicAPIService)(&c.common)
	c.LoggingOpenstackAPI = (*LoggingOpenstackAPIService)(&c.common)
	c.LoggingPapertrailAPI = (*LoggingPapertrailAPIService)(&c.common)
	c.LoggingPubsubAPI = (*LoggingPubsubAPIService)(&c.common)
	c.LoggingS3API = (*LoggingS3APIService)(&c.common)
	c.LoggingScalyrAPI = (*LoggingScalyrAPIService)(&c.common)
	c.LoggingSftpAPI = (*LoggingSftpAPIService)(&c.common)
	c.LoggingSplunkAPI = (*LoggingSplunkAPIService)(&c.common)
	c.LoggingSumologicAPI = (*LoggingSumologicAPIService)(&c.common)
	c.LoggingSyslogAPI = (*LoggingSyslogAPIService)(&c.common)
	c.MutualAuthenticationAPI = (*MutualAuthenticationAPIService)(&c.common)
	c.ObjectStoreAPI = (*ObjectStoreAPIService)(&c.common)
	c.PackageAPI = (*PackageAPIService)(&c.common)
	c.PoolAPI = (*PoolAPIService)(&c.common)
	c.PopAPI = (*PopAPIService)(&c.common)
	c.PublicIPListAPI = (*PublicIPListAPIService)(&c.common)
	c.PublishAPI = (*PublishAPIService)(&c.common)
	c.PurgeAPI = (*PurgeAPIService)(&c.common)
	c.RateLimiterAPI = (*RateLimiterAPIService)(&c.common)
	c.RealtimeAPI = (*RealtimeAPIService)(&c.common)
	c.RequestSettingsAPI = (*RequestSettingsAPIService)(&c.common)
	c.ResourceAPI = (*ResourceAPIService)(&c.common)
	c.ResponseObjectAPI = (*ResponseObjectAPIService)(&c.common)
	c.ServerAPI = (*ServerAPIService)(&c.common)
	c.ServiceAPI = (*ServiceAPIService)(&c.common)
	c.ServiceAuthorizationsAPI = (*ServiceAuthorizationsAPIService)(&c.common)
	c.SettingsAPI = (*SettingsAPIService)(&c.common)
	c.SnippetAPI = (*SnippetAPIService)(&c.common)
	c.StarAPI = (*StarAPIService)(&c.common)
	c.StatsAPI = (*StatsAPIService)(&c.common)
	c.TLSActivationsAPI = (*TLSActivationsAPIService)(&c.common)
	c.TLSBulkCertificatesAPI = (*TLSBulkCertificatesAPIService)(&c.common)
	c.TLSCertificatesAPI = (*TLSCertificatesAPIService)(&c.common)
	c.TLSConfigurationsAPI = (*TLSConfigurationsAPIService)(&c.common)
	c.TLSCsrsAPI = (*TLSCsrsAPIService)(&c.common)
	c.TLSDomainsAPI = (*TLSDomainsAPIService)(&c.common)
	c.TLSPrivateKeysAPI = (*TLSPrivateKeysAPIService)(&c.common)
	c.TLSSubscriptionsAPI = (*TLSSubscriptionsAPIService)(&c.common)
	c.TokensAPI = (*TokensAPIService)(&c.common)
	c.UserAPI = (*UserAPIService)(&c.common)
	c.VclAPI = (*VclAPIService)(&c.common)
	c.VclDiffAPI = (*VclDiffAPIService)(&c.common)
	c.VersionAPI = (*VersionAPIService)(&c.common)
	c.WafActiveRulesAPI = (*WafActiveRulesAPIService)(&c.common)
	c.WafExclusionsAPI = (*WafExclusionsAPIService)(&c.common)
	c.WafFirewallVersionsAPI = (*WafFirewallVersionsAPIService)(&c.common)
	c.WafFirewallsAPI = (*WafFirewallsAPIService)(&c.common)
	c.WafRuleRevisionsAPI = (*WafRuleRevisionsAPIService)(&c.common)
	c.WafRulesAPI = (*WafRulesAPIService)(&c.common)
	c.WafTagsAPI = (*WafTagsAPIService)(&c.common)

	return c
}

func atoi(in string) (int, error) {
	return strconv.Atoi(in)
}

// selectHeaderContentType select a content type from the available list.
func selectHeaderContentType(contentTypes []string) string {
	if len(contentTypes) == 0 {
		return ""
	}
	if contains(contentTypes, "application/json") {
		return "application/json"
	}
	return contentTypes[0] // use the first content type specified in 'consumes'
}

// selectHeaderAccept join all accept types and return
func selectHeaderAccept(accepts []string) string {
	if len(accepts) == 0 {
		return ""
	}

	if contains(accepts, "application/json") {
		return "application/json"
	}

	return strings.Join(accepts, ",")
}

// contains is a case insensitive match, finding needle in a haystack
func contains(haystack []string, needle string) bool {
	for _, a := range haystack {
		if strings.ToLower(a) == strings.ToLower(needle) {
			return true
		}
	}
	return false
}

// Verify optional parameters are of the correct type.
func typeCheckParameter(obj any, expected string, name string) error {
	// Make sure there is an object.
	if obj == nil {
		return nil
	}

	// Check the type is as expected.
	if reflect.TypeOf(obj).String() != expected {
		return fmt.Errorf("expected %s to be of type %s but received %s", name, expected, reflect.TypeOf(obj).String())
	}
	return nil
}

// parameterToString convert any parameters to string, using a delimiter if format is provided.
func parameterToString(obj any, collectionFormat string) string {
	var delimiter string

	switch collectionFormat {
	case "pipes":
		delimiter = "|"
	case "ssv":
		delimiter = " "
	case "tsv":
		delimiter = "\t"
	case "csv":
		delimiter = ","
	}

	if reflect.TypeOf(obj).Kind() == reflect.Slice {
		return strings.Trim(strings.ReplaceAll(fmt.Sprint(obj), " ", delimiter), "[]")
	} else if t, ok := obj.(time.Time); ok {
		return t.Format(time.RFC3339)
	}

	return fmt.Sprintf("%v", obj)
}

// helper for converting any parameters to json strings
func parameterToJSON(obj any) (string, error) {
	jsonBuf, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}
	return string(jsonBuf), err
}

// callAPI do the request.
func (c *APIClient) callAPI(request *http.Request) (*http.Response, error) {
	if c.cfg.Debug {
		dump, err := httputil.DumpRequestOut(request, true)
		if err != nil {
			return nil, err
		}
		log.Printf("\n%s\n", string(dump))
	}

	resp, err := c.cfg.HTTPClient.Do(request)
	if err != nil {
		return resp, err
	}

	if c.cfg.Debug {
		dump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			return resp, err
		}
		log.Printf("\n%s\n", string(dump))
	}
	return resp, err
}

// GetConfig allows modification of underlying config for alternate 
// implementations and testing. 
//
// Caution: modifying the configuration while live can cause data races and 
// potentially unwanted behavior
func (c *APIClient) GetConfig() *Configuration {
	return c.cfg
}

type formFile struct {
		fileBytes []byte
		fileName string
		formFileName string
}

// prepareRequest build the request
func (c *APIClient) prepareRequest(
	ctx context.Context,
	path string, method string,
	postBody any,
	headerParams map[string]string,
	queryParams gourl.Values,
	formParams gourl.Values,
	formFiles []formFile) (localVarRequest *http.Request, err error) {
	var body *bytes.Buffer

	// Detect postBody type and post.
	if postBody != nil {
		contentType := headerParams["Content-Type"]
		if contentType == "" {
			contentType = detectContentType(postBody)
			headerParams["Content-Type"] = contentType
		}

		body, err = setBody(postBody, contentType)
		if err != nil {
			return nil, err
		}
	}

	// add form parameters and file if available.
	if strings.HasPrefix(headerParams["Content-Type"], "multipart/form-data") && len(formParams) > 0 || (len(formFiles) > 0) {
		if body != nil {
			return nil, errors.New("cannot specify postBody and multipart form at the same time")
		}
		body = &bytes.Buffer{}
		w := multipart.NewWriter(body)

		for k, v := range formParams {
			for _, iv := range v {
				if strings.HasPrefix(k, "@") { // file
					err = addFile(w, k[1:], iv)
					if err != nil {
						return nil, err
					}
				} else { // form value
					_ = w.WriteField(k, iv)
				}
			}
		}
		for _, formFile := range formFiles {
			if len(formFile.fileBytes) > 0 && formFile.fileName != "" {
				w.Boundary()
				part, err := w.CreateFormFile(formFile.formFileName, filepath.Base(formFile.fileName))
				if err != nil {
						return nil, err
				}
				_, err = part.Write(formFile.fileBytes)
				if err != nil {
						return nil, err
				}
			}
		}

		// Set the Boundary in the Content-Type
		headerParams["Content-Type"] = w.FormDataContentType()

		// Set Content-Length
		headerParams["Content-Length"] = fmt.Sprintf("%d", body.Len())
		_ = w.Close()
	}

	if strings.HasPrefix(headerParams["Content-Type"], "application/x-www-form-urlencoded") && len(formParams) > 0 {
		if body != nil {
			return nil, errors.New("cannot specify postBody and x-www-form-urlencoded form at the same time")
		}
		body = &bytes.Buffer{}
		_, _ = body.WriteString(formParams.Encode())
		// Set Content-Length
		headerParams["Content-Length"] = fmt.Sprintf("%d", body.Len())
	}

	// Setup path and query parameters
	u, err := gourl.Parse(path)
	if err != nil {
		return nil, err
	}

	// Override request host, if applicable
	if c.cfg.Host != "" {
		u.Host = c.cfg.Host
	}

	// Override request scheme, if applicable
	if c.cfg.Scheme != "" {
		u.Scheme = c.cfg.Scheme
	}

	// Adding Query Param
	query := u.Query()
	for k, v := range queryParams {
		for _, iv := range v {
			query.Add(k, iv)
		}
	}

	// Encode the parameters.
	u.RawQuery = query.Encode()

	// Generate a new request
	if body != nil {
		localVarRequest, err = http.NewRequest(method, u.String(), body)
	} else {
		localVarRequest, err = http.NewRequest(method, u.String(), nil)
	}
	if err != nil {
		return nil, err
	}

	// add header parameters, if any
	if len(headerParams) > 0 {
		headers := http.Header{}
		for h, v := range headerParams {
			headers[h] = []string{v}
		}
		localVarRequest.Header = headers
	}

	// Add the user agent to the request.
	localVarRequest.Header.Add("User-Agent", c.cfg.UserAgent)

	if ctx != nil {
		// add context to the request
		localVarRequest = localVarRequest.WithContext(ctx)

		// Walk through any authentication.

		// OAuth2 authentication
		if tok, ok := ctx.Value(ContextOAuth2).(oauth2.TokenSource); ok {
			// We were able to grab an oauth2 token from the context
			var latestToken *oauth2.Token
			if latestToken, err = tok.Token(); err != nil {
				return nil, err
			}

			latestToken.SetAuthHeader(localVarRequest)
		}

		// Basic HTTP Authentication
		if auth, ok := ctx.Value(ContextBasicAuth).(BasicAuth); ok {
			localVarRequest.SetBasicAuth(auth.UserName, auth.Password)
		}

		// AccessToken Authentication
		if auth, ok := ctx.Value(ContextAccessToken).(string); ok {
			localVarRequest.Header.Add("Authorization", "Bearer "+auth)
		}
	}

	for header, value := range c.cfg.DefaultHeader {
		localVarRequest.Header.Add(header, value)
	}
	return localVarRequest, nil
}

func (c *APIClient) decode(v any, b []byte, contentType string) (err error) {
	if len(b) == 0 {
		return nil
	}
	if s, ok := v.(*string); ok {
		*s = string(b)
		return nil
	}
	if f, ok := v.(**os.File); ok {
		*f, err = ioutil.TempFile("", "HTTPClientFile")
		if err != nil {
			return err
		}
		_, err = (*f).Write(b)
		if err != nil {
			return err
		}
		_, err = (*f).Seek(0, io.SeekStart)
		return err
	}
	if xmlCheck.MatchString(contentType) {
		return xml.Unmarshal(b, v)
	}
	if jsonCheck.MatchString(contentType) {
		if actualObj, ok := v.(interface{ GetActualInstance() any }); ok { // oneOf, anyOf schemas
      // make sure UnmarshalJSON is defined on the type
      unmarshalObj, ok := actualObj.(interface{ UnmarshalJSON([]byte) error })
			if !ok {
				return errors.New("unknown type with GetActualInstance but no unmarshalObj.UnmarshalJSON defined")
      }
      if err = unmarshalObj.UnmarshalJSON(b); err != nil {
        return err
      }
		} else if err = json.Unmarshal(b, v); err != nil { // simple model
			return err
		}
		return nil
	}
	return errors.New("undefined response type")
}

// Add a file to the multipart request
func addFile(w *multipart.Writer, fieldName, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	part, err := w.CreateFormFile(fieldName, filepath.Base(path))
	if err != nil {
		return err
	}
	_, err = io.Copy(part, file)

	return err
}

// Prevent trying to import "fmt"
func reportError(format string, a ...any) error {
	return fmt.Errorf(format, a...)
}

// A wrapper for strict JSON decoding
func newStrictDecoder(data []byte) *json.Decoder {
	dec := json.NewDecoder(bytes.NewBuffer(data))
	dec.DisallowUnknownFields()
	return dec
}

// Set request body from an any
func setBody(body any, contentType string) (bodyBuf *bytes.Buffer, err error) {
	if bodyBuf == nil {
		bodyBuf = &bytes.Buffer{}
	}

	if reader, ok := body.(io.Reader); ok {
		_, err = bodyBuf.ReadFrom(reader)
	} else if fp, ok := body.(**os.File); ok {
		_, err = bodyBuf.ReadFrom(*fp)
	} else if b, ok := body.([]byte); ok {
		_, err = bodyBuf.Write(b)
	} else if s, ok := body.(string); ok {
		_, err = bodyBuf.WriteString(s)
	} else if s, ok := body.(*string); ok {
		_, err = bodyBuf.WriteString(*s)
	} else if jsonCheck.MatchString(contentType) {
		err = json.NewEncoder(bodyBuf).Encode(body)
	} else if xmlCheck.MatchString(contentType) {
		err = xml.NewEncoder(bodyBuf).Encode(body)
	}

	if err != nil {
		return nil, err
	}

	if bodyBuf.Len() == 0 {
		err = fmt.Errorf("invalid body type %s", contentType)
		return nil, err
	}
	return bodyBuf, nil
}

// detectContentType method is used to figure out `Request.Body` content type for request header
func detectContentType(body any) string {
	contentType := "text/plain; charset=utf-8"
	kind := reflect.TypeOf(body).Kind()

	switch kind {
	case reflect.Struct, reflect.Map, reflect.Ptr:
		contentType = "application/json; charset=utf-8"
	case reflect.String:
		contentType = "text/plain; charset=utf-8"
	default:
		if b, ok := body.([]byte); ok {
			contentType = http.DetectContentType(b)
		} else if kind == reflect.Slice {
			contentType = "application/json; charset=utf-8"
		}
	}

	return contentType
}

// Ripped from https://github.com/gregjones/httpcache/blob/master/httpcache.go
type cacheControl map[string]string

func parseCacheControl(headers http.Header) cacheControl {
	cc := cacheControl{}
	ccHeader := headers.Get("Cache-Control")
	for _, part := range strings.Split(ccHeader, ",") {
		part = strings.Trim(part, " ")
		if part == "" {
			continue
		}
		if strings.ContainsRune(part, '=') {
			keyval := strings.Split(part, "=")
			cc[strings.Trim(keyval[0], " ")] = strings.Trim(keyval[1], ",")
		} else {
			cc[part] = ""
		}
	}
	return cc
}

// CacheExpires is a helper function to determine remaining time before repeating a request.
func CacheExpires(r *http.Response) time.Time {
	// Figure out when the cache expires.
	var expires time.Time
	now, err := time.Parse(time.RFC1123, r.Header.Get("date"))
	if err != nil {
		return time.Now()
	}
	respCacheControl := parseCacheControl(r.Header)

	if maxAge, ok := respCacheControl["max-age"]; ok {
		lifetime, err := time.ParseDuration(maxAge + "s")
		if err != nil {
			expires = now
		} else {
			expires = now.Add(lifetime)
		}
	} else {
		expiresHeader := r.Header.Get("Expires")
		if expiresHeader != "" {
			expires, err = time.Parse(time.RFC1123, expiresHeader)
			if err != nil {
				expires = now
			}
		}
	}
	return expires
}

func strlen(s string) int {
	return utf8.RuneCountInString(s)
}

// GenericAPIError provides access to the body, error and model on returned errors.
type GenericAPIError struct {
	body  []byte
	error string
	model any
}

// Error returns non-empty string if there was an error.
func (e GenericAPIError) Error() string {
	return e.error
}

// Body returns the raw bytes of the response
func (e GenericAPIError) Body() []byte {
	return e.body
}

// Model returns the unpacked model of the error
func (e GenericAPIError) Model() any {
	return e.model
}
