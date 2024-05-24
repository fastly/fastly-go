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

// WafFirewallVersionDataAttributes struct for WafFirewallVersionDataAttributes
type WafFirewallVersionDataAttributes struct {
	// Allowed HTTP versions.
	AllowedHTTPVersions *string `json:"allowed_http_versions,omitempty"`
	// A space-separated list of HTTP method names.
	AllowedMethods *string `json:"allowed_methods,omitempty"`
	// Allowed request content types.
	AllowedRequestContentType *string `json:"allowed_request_content_type,omitempty"`
	// Allowed request content type charset.
	AllowedRequestContentTypeCharset *string `json:"allowed_request_content_type_charset,omitempty"`
	// The maximum allowed argument name length.
	ArgNameLength *int32 `json:"arg_name_length,omitempty"`
	// The maximum allowed length of an argument.
	ArgLength *int32 `json:"arg_length,omitempty"`
	// The maximum allowed size of all files (in bytes).
	CombinedFileSizes *int32 `json:"combined_file_sizes,omitempty"`
	// A freeform descriptive note.
	Comment NullableString `json:"comment,omitempty"`
	// Score value to add for critical anomalies.
	CriticalAnomalyScore *int32 `json:"critical_anomaly_score,omitempty"`
	// CRS validate UTF8 encoding.
	CrsValidateUtf8Encoding *bool `json:"crs_validate_utf8_encoding,omitempty"`
	// Score value to add for error anomalies.
	ErrorAnomalyScore *int32 `json:"error_anomaly_score,omitempty"`
	// A space-separated list of country codes in ISO 3166-1 (two-letter) format.
	HighRiskCountryCodes *string `json:"high_risk_country_codes,omitempty"`
	// HTTP violation threshold.
	HTTPViolationScoreThreshold *int32 `json:"http_violation_score_threshold,omitempty"`
	// Inbound anomaly threshold.
	InboundAnomalyScoreThreshold *int32 `json:"inbound_anomaly_score_threshold,omitempty"`
	// Local file inclusion attack threshold.
	LfiScoreThreshold *int32 `json:"lfi_score_threshold,omitempty"`
	// Whether a specific firewall version is locked from being modified.
	Locked *bool `json:"locked,omitempty"`
	// The maximum allowed file size, in bytes.
	MaxFileSize *int32 `json:"max_file_size,omitempty"`
	// The maximum number of arguments allowed.
	MaxNumArgs *int32 `json:"max_num_args,omitempty"`
	// Score value to add for notice anomalies.
	NoticeAnomalyScore *int32 `json:"notice_anomaly_score,omitempty"`
	Number *int32 `json:"number,omitempty"`
	// The configured paranoia level.
	ParanoiaLevel *int32 `json:"paranoia_level,omitempty"`
	// PHP injection threshold.
	PhpInjectionScoreThreshold *int32 `json:"php_injection_score_threshold,omitempty"`
	// Remote code execution threshold.
	RceScoreThreshold *int32 `json:"rce_score_threshold,omitempty"`
	// A space-separated list of allowed file extensions.
	RestrictedExtensions *string `json:"restricted_extensions,omitempty"`
	// A space-separated list of allowed header names.
	RestrictedHeaders *string `json:"restricted_headers,omitempty"`
	// Remote file inclusion attack threshold.
	RfiScoreThreshold *int32 `json:"rfi_score_threshold,omitempty"`
	// Session fixation attack threshold.
	SessionFixationScoreThreshold *int32 `json:"session_fixation_score_threshold,omitempty"`
	// SQL injection attack threshold.
	SQLInjectionScoreThreshold *int32 `json:"sql_injection_score_threshold,omitempty"`
	// The maximum size of argument names and values.
	TotalArgLength *int32 `json:"total_arg_length,omitempty"`
	// Score value to add for warning anomalies.
	WarningAnomalyScore *int32 `json:"warning_anomaly_score,omitempty"`
	// XSS attack threshold.
	XSSScoreThreshold *int32 `json:"xss_score_threshold,omitempty"`
	AdditionalProperties map[string]any
}

type _WafFirewallVersionDataAttributes WafFirewallVersionDataAttributes

// NewWafFirewallVersionDataAttributes instantiates a new WafFirewallVersionDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWafFirewallVersionDataAttributes() *WafFirewallVersionDataAttributes {
	this := WafFirewallVersionDataAttributes{}
	var allowedHTTPVersions string = "HTTP/1.0 HTTP/1.1 HTTP/2"
	this.AllowedHTTPVersions = &allowedHTTPVersions
	var allowedMethods string = "GET HEAD POST OPTIONS PUT PATCH DELETE"
	this.AllowedMethods = &allowedMethods
	var allowedRequestContentType string = "application/x-www-form-urlencoded|multipart/form-data|text/xml|application/xml|application/x-amf|application/json|text/plain"
	this.AllowedRequestContentType = &allowedRequestContentType
	var allowedRequestContentTypeCharset string = "utf-8|iso-8859-1|iso-8859-15|windows-1252"
	this.AllowedRequestContentTypeCharset = &allowedRequestContentTypeCharset
	var argNameLength int32 = 100
	this.ArgNameLength = &argNameLength
	var argLength int32 = 400
	this.ArgLength = &argLength
	var combinedFileSizes int32 = 10000000
	this.CombinedFileSizes = &combinedFileSizes
	var criticalAnomalyScore int32 = 6
	this.CriticalAnomalyScore = &criticalAnomalyScore
	var errorAnomalyScore int32 = 5
	this.ErrorAnomalyScore = &errorAnomalyScore
	var locked bool = false
	this.Locked = &locked
	var maxFileSize int32 = 10000000
	this.MaxFileSize = &maxFileSize
	var maxNumArgs int32 = 255
	this.MaxNumArgs = &maxNumArgs
	var noticeAnomalyScore int32 = 4
	this.NoticeAnomalyScore = &noticeAnomalyScore
	var paranoiaLevel int32 = 1
	this.ParanoiaLevel = &paranoiaLevel
	var restrictedExtensions string = ".asa/ .asax/ .ascx/ .axd/ .backup/ .bak/ .bat/ .cdx/ .cer/ .cfg/ .cmd/ .com/ .config/ .conf/ .cs/ .csproj/ .csr/ .dat/ .db/ .dbf/ .dll/ .dos/ .htr/ .htw/ .ida/ .idc/ .idq/ .inc/ .ini/ .key/ .licx/ .lnk/ .log/ .mdb/ .old/ .pass/ .pdb/ .pol/ .printer/ .pwd/ .resources/ .resx/ .sql/ .sys/ .vb/ .vbs/ .vbproj/ .vsdisco/ .webinfo/ .xsd/ .xsx"
	this.RestrictedExtensions = &restrictedExtensions
	var restrictedHeaders string = "/proxy/ /lock-token/ /content-range/ /translate/ /if/"
	this.RestrictedHeaders = &restrictedHeaders
	var totalArgLength int32 = 6400
	this.TotalArgLength = &totalArgLength
	return &this
}

// NewWafFirewallVersionDataAttributesWithDefaults instantiates a new WafFirewallVersionDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWafFirewallVersionDataAttributesWithDefaults() *WafFirewallVersionDataAttributes {
	this := WafFirewallVersionDataAttributes{}
	var allowedHTTPVersions string = "HTTP/1.0 HTTP/1.1 HTTP/2"
	this.AllowedHTTPVersions = &allowedHTTPVersions
	var allowedMethods string = "GET HEAD POST OPTIONS PUT PATCH DELETE"
	this.AllowedMethods = &allowedMethods
	var allowedRequestContentType string = "application/x-www-form-urlencoded|multipart/form-data|text/xml|application/xml|application/x-amf|application/json|text/plain"
	this.AllowedRequestContentType = &allowedRequestContentType
	var allowedRequestContentTypeCharset string = "utf-8|iso-8859-1|iso-8859-15|windows-1252"
	this.AllowedRequestContentTypeCharset = &allowedRequestContentTypeCharset
	var argNameLength int32 = 100
	this.ArgNameLength = &argNameLength
	var argLength int32 = 400
	this.ArgLength = &argLength
	var combinedFileSizes int32 = 10000000
	this.CombinedFileSizes = &combinedFileSizes
	var criticalAnomalyScore int32 = 6
	this.CriticalAnomalyScore = &criticalAnomalyScore
	var errorAnomalyScore int32 = 5
	this.ErrorAnomalyScore = &errorAnomalyScore
	var locked bool = false
	this.Locked = &locked
	var maxFileSize int32 = 10000000
	this.MaxFileSize = &maxFileSize
	var maxNumArgs int32 = 255
	this.MaxNumArgs = &maxNumArgs
	var noticeAnomalyScore int32 = 4
	this.NoticeAnomalyScore = &noticeAnomalyScore
	var paranoiaLevel int32 = 1
	this.ParanoiaLevel = &paranoiaLevel
	var restrictedExtensions string = ".asa/ .asax/ .ascx/ .axd/ .backup/ .bak/ .bat/ .cdx/ .cer/ .cfg/ .cmd/ .com/ .config/ .conf/ .cs/ .csproj/ .csr/ .dat/ .db/ .dbf/ .dll/ .dos/ .htr/ .htw/ .ida/ .idc/ .idq/ .inc/ .ini/ .key/ .licx/ .lnk/ .log/ .mdb/ .old/ .pass/ .pdb/ .pol/ .printer/ .pwd/ .resources/ .resx/ .sql/ .sys/ .vb/ .vbs/ .vbproj/ .vsdisco/ .webinfo/ .xsd/ .xsx"
	this.RestrictedExtensions = &restrictedExtensions
	var restrictedHeaders string = "/proxy/ /lock-token/ /content-range/ /translate/ /if/"
	this.RestrictedHeaders = &restrictedHeaders
	var totalArgLength int32 = 6400
	this.TotalArgLength = &totalArgLength
	return &this
}

// GetAllowedHTTPVersions returns the AllowedHTTPVersions field value if set, zero value otherwise.
func (o *WafFirewallVersionDataAttributes) GetAllowedHTTPVersions() string {
	if o == nil || o.AllowedHTTPVersions == nil {
		var ret string
		return ret
	}
	return *o.AllowedHTTPVersions
}

// GetAllowedHTTPVersionsOk returns a tuple with the AllowedHTTPVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallVersionDataAttributes) GetAllowedHTTPVersionsOk() (*string, bool) {
	if o == nil || o.AllowedHTTPVersions == nil {
		return nil, false
	}
	return o.AllowedHTTPVersions, true
}

// HasAllowedHTTPVersions returns a boolean if a field has been set.
func (o *WafFirewallVersionDataAttributes) HasAllowedHTTPVersions() bool {
	if o != nil && o.AllowedHTTPVersions != nil {
		return true
	}

	return false
}

// SetAllowedHTTPVersions gets a reference to the given string and assigns it to the AllowedHTTPVersions field.
func (o *WafFirewallVersionDataAttributes) SetAllowedHTTPVersions(v string) {
	o.AllowedHTTPVersions = &v
}

// GetAllowedMethods returns the AllowedMethods field value if set, zero value otherwise.
func (o *WafFirewallVersionDataAttributes) GetAllowedMethods() string {
	if o == nil || o.AllowedMethods == nil {
		var ret string
		return ret
	}
	return *o.AllowedMethods
}

// GetAllowedMethodsOk returns a tuple with the AllowedMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallVersionDataAttributes) GetAllowedMethodsOk() (*string, bool) {
	if o == nil || o.AllowedMethods == nil {
		return nil, false
	}
	return o.AllowedMethods, true
}

// HasAllowedMethods returns a boolean if a field has been set.
func (o *WafFirewallVersionDataAttributes) HasAllowedMethods() bool {
	if o != nil && o.AllowedMethods != nil {
		return true
	}

	return false
}

// SetAllowedMethods gets a reference to the given string and assigns it to the AllowedMethods field.
func (o *WafFirewallVersionDataAttributes) SetAllowedMethods(v string) {
	o.AllowedMethods = &v
}

// GetAllowedRequestContentType returns the AllowedRequestContentType field value if set, zero value otherwise.
func (o *WafFirewallVersionDataAttributes) GetAllowedRequestContentType() string {
	if o == nil || o.AllowedRequestContentType == nil {
		var ret string
		return ret
	}
	return *o.AllowedRequestContentType
}

// GetAllowedRequestContentTypeOk returns a tuple with the AllowedRequestContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallVersionDataAttributes) GetAllowedRequestContentTypeOk() (*string, bool) {
	if o == nil || o.AllowedRequestContentType == nil {
		return nil, false
	}
	return o.AllowedRequestContentType, true
}

// HasAllowedRequestContentType returns a boolean if a field has been set.
func (o *WafFirewallVersionDataAttributes) HasAllowedRequestContentType() bool {
	if o != nil && o.AllowedRequestContentType != nil {
		return true
	}

	return false
}

// SetAllowedRequestContentType gets a reference to the given string and assigns it to the AllowedRequestContentType field.
func (o *WafFirewallVersionDataAttributes) SetAllowedRequestContentType(v string) {
	o.AllowedRequestContentType = &v
}

// GetAllowedRequestContentTypeCharset returns the AllowedRequestContentTypeCharset field value if set, zero value otherwise.
func (o *WafFirewallVersionDataAttributes) GetAllowedRequestContentTypeCharset() string {
	if o == nil || o.AllowedRequestContentTypeCharset == nil {
		var ret string
		return ret
	}
	return *o.AllowedRequestContentTypeCharset
}

// GetAllowedRequestContentTypeCharsetOk returns a tuple with the AllowedRequestContentTypeCharset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallVersionDataAttributes) GetAllowedRequestContentTypeCharsetOk() (*string, bool) {
	if o == nil || o.AllowedRequestContentTypeCharset == nil {
		return nil, false
	}
	return o.AllowedRequestContentTypeCharset, true
}

// HasAllowedRequestContentTypeCharset returns a boolean if a field has been set.
func (o *WafFirewallVersionDataAttributes) HasAllowedRequestContentTypeCharset() bool {
	if o != nil && o.AllowedRequestContentTypeCharset != nil {
		return true
	}

	return false
}

// SetAllowedRequestContentTypeCharset gets a reference to the given string and assigns it to the AllowedRequestContentTypeCharset field.
func (o *WafFirewallVersionDataAttributes) SetAllowedRequestContentTypeCharset(v string) {
	o.AllowedRequestContentTypeCharset = &v
}

// GetArgNameLength returns the ArgNameLength field value if set, zero value otherwise.
func (o *WafFirewallVersionDataAttributes) GetArgNameLength() int32 {
	if o == nil || o.ArgNameLength == nil {
		var ret int32
		return ret
	}
	return *o.ArgNameLength
}

// GetArgNameLengthOk returns a tuple with the ArgNameLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallVersionDataAttributes) GetArgNameLengthOk() (*int32, bool) {
	if o == nil || o.ArgNameLength == nil {
		return nil, false
	}
	return o.ArgNameLength, true
}

// HasArgNameLength returns a boolean if a field has been set.
func (o *WafFirewallVersionDataAttributes) HasArgNameLength() bool {
	if o != nil && o.ArgNameLength != nil {
		return true
	}

	return false
}

// SetArgNameLength gets a reference to the given int32 and assigns it to the ArgNameLength field.
func (o *WafFirewallVersionDataAttributes) SetArgNameLength(v int32) {
	o.ArgNameLength = &v
}

// GetArgLength returns the ArgLength field value if set, zero value otherwise.
func (o *WafFirewallVersionDataAttributes) GetArgLength() int32 {
	if o == nil || o.ArgLength == nil {
		var ret int32
		return ret
	}
	return *o.ArgLength
}

// GetArgLengthOk returns a tuple with the ArgLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallVersionDataAttributes) GetArgLengthOk() (*int32, bool) {
	if o == nil || o.ArgLength == nil {
		return nil, false
	}
	return o.ArgLength, true
}

// HasArgLength returns a boolean if a field has been set.
func (o *WafFirewallVersionDataAttributes) HasArgLength() bool {
	if o != nil && o.ArgLength != nil {
		return true
	}

	return false
}

// SetArgLength gets a reference to the given int32 and assigns it to the ArgLength field.
func (o *WafFirewallVersionDataAttributes) SetArgLength(v int32) {
	o.ArgLength = &v
}

// GetCombinedFileSizes returns the CombinedFileSizes field value if set, zero value otherwise.
func (o *WafFirewallVersionDataAttributes) GetCombinedFileSizes() int32 {
	if o == nil || o.CombinedFileSizes == nil {
		var ret int32
		return ret
	}
	return *o.CombinedFileSizes
}

// GetCombinedFileSizesOk returns a tuple with the CombinedFileSizes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallVersionDataAttributes) GetCombinedFileSizesOk() (*int32, bool) {
	if o == nil || o.CombinedFileSizes == nil {
		return nil, false
	}
	return o.CombinedFileSizes, true
}

// HasCombinedFileSizes returns a boolean if a field has been set.
func (o *WafFirewallVersionDataAttributes) HasCombinedFileSizes() bool {
	if o != nil && o.CombinedFileSizes != nil {
		return true
	}

	return false
}

// SetCombinedFileSizes gets a reference to the given int32 and assigns it to the CombinedFileSizes field.
func (o *WafFirewallVersionDataAttributes) SetCombinedFileSizes(v int32) {
	o.CombinedFileSizes = &v
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WafFirewallVersionDataAttributes) GetComment() string {
	if o == nil || o.Comment.Get() == nil {
		var ret string
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WafFirewallVersionDataAttributes) GetCommentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *WafFirewallVersionDataAttributes) HasComment() bool {
	if o != nil && o.Comment.IsSet() {
		return true
	}

	return false
}

// SetComment gets a reference to the given NullableString and assigns it to the Comment field.
func (o *WafFirewallVersionDataAttributes) SetComment(v string) {
	o.Comment.Set(&v)
}
// SetCommentNil sets the value for Comment to be an explicit nil
func (o *WafFirewallVersionDataAttributes) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil
func (o *WafFirewallVersionDataAttributes) UnsetComment() {
	o.Comment.Unset()
}

// GetCriticalAnomalyScore returns the CriticalAnomalyScore field value if set, zero value otherwise.
func (o *WafFirewallVersionDataAttributes) GetCriticalAnomalyScore() int32 {
	if o == nil || o.CriticalAnomalyScore == nil {
		var ret int32
		return ret
	}
	return *o.CriticalAnomalyScore
}

// GetCriticalAnomalyScoreOk returns a tuple with the CriticalAnomalyScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallVersionDataAttributes) GetCriticalAnomalyScoreOk() (*int32, bool) {
	if o == nil || o.CriticalAnomalyScore == nil {
		return nil, false
	}
	return o.CriticalAnomalyScore, true
}

// HasCriticalAnomalyScore returns a boolean if a field has been set.
func (o *WafFirewallVersionDataAttributes) HasCriticalAnomalyScore() bool {
	if o != nil && o.CriticalAnomalyScore != nil {
		return true
	}

	return false
}

// SetCriticalAnomalyScore gets a reference to the given int32 and assigns it to the CriticalAnomalyScore field.
func (o *WafFirewallVersionDataAttributes) SetCriticalAnomalyScore(v int32) {
	o.CriticalAnomalyScore = &v
}

// GetCrsValidateUtf8Encoding returns the CrsValidateUtf8Encoding field value if set, zero value otherwise.
func (o *WafFirewallVersionDataAttributes) GetCrsValidateUtf8Encoding() bool {
	if o == nil || o.CrsValidateUtf8Encoding == nil {
		var ret bool
		return ret
	}
	return *o.CrsValidateUtf8Encoding
}

// GetCrsValidateUtf8EncodingOk returns a tuple with the CrsValidateUtf8Encoding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallVersionDataAttributes) GetCrsValidateUtf8EncodingOk() (*bool, bool) {
	if o == nil || o.CrsValidateUtf8Encoding == nil {
		return nil, false
	}
	return o.CrsValidateUtf8Encoding, true
}

// HasCrsValidateUtf8Encoding returns a boolean if a field has been set.
func (o *WafFirewallVersionDataAttributes) HasCrsValidateUtf8Encoding() bool {
	if o != nil && o.CrsValidateUtf8Encoding != nil {
		return true
	}

	return false
}

// SetCrsValidateUtf8Encoding gets a reference to the given bool and assigns it to the CrsValidateUtf8Encoding field.
func (o *WafFirewallVersionDataAttributes) SetCrsValidateUtf8Encoding(v bool) {
	o.CrsValidateUtf8Encoding = &v
}

// GetErrorAnomalyScore returns the ErrorAnomalyScore field value if set, zero value otherwise.
func (o *WafFirewallVersionDataAttributes) GetErrorAnomalyScore() int32 {
	if o == nil || o.ErrorAnomalyScore == nil {
		var ret int32
		return ret
	}
	return *o.ErrorAnomalyScore
}

// GetErrorAnomalyScoreOk returns a tuple with the ErrorAnomalyScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallVersionDataAttributes) GetErrorAnomalyScoreOk() (*int32, bool) {
	if o == nil || o.ErrorAnomalyScore == nil {
		return nil, false
	}
	return o.ErrorAnomalyScore, true
}

// HasErrorAnomalyScore returns a boolean if a field has been set.
func (o *WafFirewallVersionDataAttributes) HasErrorAnomalyScore() bool {
	if o != nil && o.ErrorAnomalyScore != nil {
		return true
	}

	return false
}

// SetErrorAnomalyScore gets a reference to the given int32 and assigns it to the ErrorAnomalyScore field.
func (o *WafFirewallVersionDataAttributes) SetErrorAnomalyScore(v int32) {
	o.ErrorAnomalyScore = &v
}

// GetHighRiskCountryCodes returns the HighRiskCountryCodes field value if set, zero value otherwise.
func (o *WafFirewallVersionDataAttributes) GetHighRiskCountryCodes() string {
	if o == nil || o.HighRiskCountryCodes == nil {
		var ret string
		return ret
	}
	return *o.HighRiskCountryCodes
}

// GetHighRiskCountryCodesOk returns a tuple with the HighRiskCountryCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallVersionDataAttributes) GetHighRiskCountryCodesOk() (*string, bool) {
	if o == nil || o.HighRiskCountryCodes == nil {
		return nil, false
	}
	return o.HighRiskCountryCodes, true
}

// HasHighRiskCountryCodes returns a boolean if a field has been set.
func (o *WafFirewallVersionDataAttributes) HasHighRiskCountryCodes() bool {
	if o != nil && o.HighRiskCountryCodes != nil {
		return true
	}

	return false
}

// SetHighRiskCountryCodes gets a reference to the given string and assigns it to the HighRiskCountryCodes field.
func (o *WafFirewallVersionDataAttributes) SetHighRiskCountryCodes(v string) {
	o.HighRiskCountryCodes = &v
}

// GetHTTPViolationScoreThreshold returns the HTTPViolationScoreThreshold field value if set, zero value otherwise.
func (o *WafFirewallVersionDataAttributes) GetHTTPViolationScoreThreshold() int32 {
	if o == nil || o.HTTPViolationScoreThreshold == nil {
		var ret int32
		return ret
	}
	return *o.HTTPViolationScoreThreshold
}

// GetHTTPViolationScoreThresholdOk returns a tuple with the HTTPViolationScoreThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallVersionDataAttributes) GetHTTPViolationScoreThresholdOk() (*int32, bool) {
	if o == nil || o.HTTPViolationScoreThreshold == nil {
		return nil, false
	}
	return o.HTTPViolationScoreThreshold, true
}

// HasHTTPViolationScoreThreshold returns a boolean if a field has been set.
func (o *WafFirewallVersionDataAttributes) HasHTTPViolationScoreThreshold() bool {
	if o != nil && o.HTTPViolationScoreThreshold != nil {
		return true
	}

	return false
}

// SetHTTPViolationScoreThreshold gets a reference to the given int32 and assigns it to the HTTPViolationScoreThreshold field.
func (o *WafFirewallVersionDataAttributes) SetHTTPViolationScoreThreshold(v int32) {
	o.HTTPViolationScoreThreshold = &v
}

// GetInboundAnomalyScoreThreshold returns the InboundAnomalyScoreThreshold field value if set, zero value otherwise.
func (o *WafFirewallVersionDataAttributes) GetInboundAnomalyScoreThreshold() int32 {
	if o == nil || o.InboundAnomalyScoreThreshold == nil {
		var ret int32
		return ret
	}
	return *o.InboundAnomalyScoreThreshold
}

// GetInboundAnomalyScoreThresholdOk returns a tuple with the InboundAnomalyScoreThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallVersionDataAttributes) GetInboundAnomalyScoreThresholdOk() (*int32, bool) {
	if o == nil || o.InboundAnomalyScoreThreshold == nil {
		return nil, false
	}
	return o.InboundAnomalyScoreThreshold, true
}

// HasInboundAnomalyScoreThreshold returns a boolean if a field has been set.
func (o *WafFirewallVersionDataAttributes) HasInboundAnomalyScoreThreshold() bool {
	if o != nil && o.InboundAnomalyScoreThreshold != nil {
		return true
	}

	return false
}

// SetInboundAnomalyScoreThreshold gets a reference to the given int32 and assigns it to the InboundAnomalyScoreThreshold field.
func (o *WafFirewallVersionDataAttributes) SetInboundAnomalyScoreThreshold(v int32) {
	o.InboundAnomalyScoreThreshold = &v
}

// GetLfiScoreThreshold returns the LfiScoreThreshold field value if set, zero value otherwise.
func (o *WafFirewallVersionDataAttributes) GetLfiScoreThreshold() int32 {
	if o == nil || o.LfiScoreThreshold == nil {
		var ret int32
		return ret
	}
	return *o.LfiScoreThreshold
}

// GetLfiScoreThresholdOk returns a tuple with the LfiScoreThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallVersionDataAttributes) GetLfiScoreThresholdOk() (*int32, bool) {
	if o == nil || o.LfiScoreThreshold == nil {
		return nil, false
	}
	return o.LfiScoreThreshold, true
}

// HasLfiScoreThreshold returns a boolean if a field has been set.
func (o *WafFirewallVersionDataAttributes) HasLfiScoreThreshold() bool {
	if o != nil && o.LfiScoreThreshold != nil {
		return true
	}

	return false
}

// SetLfiScoreThreshold gets a reference to the given int32 and assigns it to the LfiScoreThreshold field.
func (o *WafFirewallVersionDataAttributes) SetLfiScoreThreshold(v int32) {
	o.LfiScoreThreshold = &v
}

// GetLocked returns the Locked field value if set, zero value otherwise.
func (o *WafFirewallVersionDataAttributes) GetLocked() bool {
	if o == nil || o.Locked == nil {
		var ret bool
		return ret
	}
	return *o.Locked
}

// GetLockedOk returns a tuple with the Locked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallVersionDataAttributes) GetLockedOk() (*bool, bool) {
	if o == nil || o.Locked == nil {
		return nil, false
	}
	return o.Locked, true
}

// HasLocked returns a boolean if a field has been set.
func (o *WafFirewallVersionDataAttributes) HasLocked() bool {
	if o != nil && o.Locked != nil {
		return true
	}

	return false
}

// SetLocked gets a reference to the given bool and assigns it to the Locked field.
func (o *WafFirewallVersionDataAttributes) SetLocked(v bool) {
	o.Locked = &v
}

// GetMaxFileSize returns the MaxFileSize field value if set, zero value otherwise.
func (o *WafFirewallVersionDataAttributes) GetMaxFileSize() int32 {
	if o == nil || o.MaxFileSize == nil {
		var ret int32
		return ret
	}
	return *o.MaxFileSize
}

// GetMaxFileSizeOk returns a tuple with the MaxFileSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallVersionDataAttributes) GetMaxFileSizeOk() (*int32, bool) {
	if o == nil || o.MaxFileSize == nil {
		return nil, false
	}
	return o.MaxFileSize, true
}

// HasMaxFileSize returns a boolean if a field has been set.
func (o *WafFirewallVersionDataAttributes) HasMaxFileSize() bool {
	if o != nil && o.MaxFileSize != nil {
		return true
	}

	return false
}

// SetMaxFileSize gets a reference to the given int32 and assigns it to the MaxFileSize field.
func (o *WafFirewallVersionDataAttributes) SetMaxFileSize(v int32) {
	o.MaxFileSize = &v
}

// GetMaxNumArgs returns the MaxNumArgs field value if set, zero value otherwise.
func (o *WafFirewallVersionDataAttributes) GetMaxNumArgs() int32 {
	if o == nil || o.MaxNumArgs == nil {
		var ret int32
		return ret
	}
	return *o.MaxNumArgs
}

// GetMaxNumArgsOk returns a tuple with the MaxNumArgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallVersionDataAttributes) GetMaxNumArgsOk() (*int32, bool) {
	if o == nil || o.MaxNumArgs == nil {
		return nil, false
	}
	return o.MaxNumArgs, true
}

// HasMaxNumArgs returns a boolean if a field has been set.
func (o *WafFirewallVersionDataAttributes) HasMaxNumArgs() bool {
	if o != nil && o.MaxNumArgs != nil {
		return true
	}

	return false
}

// SetMaxNumArgs gets a reference to the given int32 and assigns it to the MaxNumArgs field.
func (o *WafFirewallVersionDataAttributes) SetMaxNumArgs(v int32) {
	o.MaxNumArgs = &v
}

// GetNoticeAnomalyScore returns the NoticeAnomalyScore field value if set, zero value otherwise.
func (o *WafFirewallVersionDataAttributes) GetNoticeAnomalyScore() int32 {
	if o == nil || o.NoticeAnomalyScore == nil {
		var ret int32
		return ret
	}
	return *o.NoticeAnomalyScore
}

// GetNoticeAnomalyScoreOk returns a tuple with the NoticeAnomalyScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallVersionDataAttributes) GetNoticeAnomalyScoreOk() (*int32, bool) {
	if o == nil || o.NoticeAnomalyScore == nil {
		return nil, false
	}
	return o.NoticeAnomalyScore, true
}

// HasNoticeAnomalyScore returns a boolean if a field has been set.
func (o *WafFirewallVersionDataAttributes) HasNoticeAnomalyScore() bool {
	if o != nil && o.NoticeAnomalyScore != nil {
		return true
	}

	return false
}

// SetNoticeAnomalyScore gets a reference to the given int32 and assigns it to the NoticeAnomalyScore field.
func (o *WafFirewallVersionDataAttributes) SetNoticeAnomalyScore(v int32) {
	o.NoticeAnomalyScore = &v
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *WafFirewallVersionDataAttributes) GetNumber() int32 {
	if o == nil || o.Number == nil {
		var ret int32
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallVersionDataAttributes) GetNumberOk() (*int32, bool) {
	if o == nil || o.Number == nil {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *WafFirewallVersionDataAttributes) HasNumber() bool {
	if o != nil && o.Number != nil {
		return true
	}

	return false
}

// SetNumber gets a reference to the given int32 and assigns it to the Number field.
func (o *WafFirewallVersionDataAttributes) SetNumber(v int32) {
	o.Number = &v
}

// GetParanoiaLevel returns the ParanoiaLevel field value if set, zero value otherwise.
func (o *WafFirewallVersionDataAttributes) GetParanoiaLevel() int32 {
	if o == nil || o.ParanoiaLevel == nil {
		var ret int32
		return ret
	}
	return *o.ParanoiaLevel
}

// GetParanoiaLevelOk returns a tuple with the ParanoiaLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallVersionDataAttributes) GetParanoiaLevelOk() (*int32, bool) {
	if o == nil || o.ParanoiaLevel == nil {
		return nil, false
	}
	return o.ParanoiaLevel, true
}

// HasParanoiaLevel returns a boolean if a field has been set.
func (o *WafFirewallVersionDataAttributes) HasParanoiaLevel() bool {
	if o != nil && o.ParanoiaLevel != nil {
		return true
	}

	return false
}

// SetParanoiaLevel gets a reference to the given int32 and assigns it to the ParanoiaLevel field.
func (o *WafFirewallVersionDataAttributes) SetParanoiaLevel(v int32) {
	o.ParanoiaLevel = &v
}

// GetPhpInjectionScoreThreshold returns the PhpInjectionScoreThreshold field value if set, zero value otherwise.
func (o *WafFirewallVersionDataAttributes) GetPhpInjectionScoreThreshold() int32 {
	if o == nil || o.PhpInjectionScoreThreshold == nil {
		var ret int32
		return ret
	}
	return *o.PhpInjectionScoreThreshold
}

// GetPhpInjectionScoreThresholdOk returns a tuple with the PhpInjectionScoreThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallVersionDataAttributes) GetPhpInjectionScoreThresholdOk() (*int32, bool) {
	if o == nil || o.PhpInjectionScoreThreshold == nil {
		return nil, false
	}
	return o.PhpInjectionScoreThreshold, true
}

// HasPhpInjectionScoreThreshold returns a boolean if a field has been set.
func (o *WafFirewallVersionDataAttributes) HasPhpInjectionScoreThreshold() bool {
	if o != nil && o.PhpInjectionScoreThreshold != nil {
		return true
	}

	return false
}

// SetPhpInjectionScoreThreshold gets a reference to the given int32 and assigns it to the PhpInjectionScoreThreshold field.
func (o *WafFirewallVersionDataAttributes) SetPhpInjectionScoreThreshold(v int32) {
	o.PhpInjectionScoreThreshold = &v
}

// GetRceScoreThreshold returns the RceScoreThreshold field value if set, zero value otherwise.
func (o *WafFirewallVersionDataAttributes) GetRceScoreThreshold() int32 {
	if o == nil || o.RceScoreThreshold == nil {
		var ret int32
		return ret
	}
	return *o.RceScoreThreshold
}

// GetRceScoreThresholdOk returns a tuple with the RceScoreThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallVersionDataAttributes) GetRceScoreThresholdOk() (*int32, bool) {
	if o == nil || o.RceScoreThreshold == nil {
		return nil, false
	}
	return o.RceScoreThreshold, true
}

// HasRceScoreThreshold returns a boolean if a field has been set.
func (o *WafFirewallVersionDataAttributes) HasRceScoreThreshold() bool {
	if o != nil && o.RceScoreThreshold != nil {
		return true
	}

	return false
}

// SetRceScoreThreshold gets a reference to the given int32 and assigns it to the RceScoreThreshold field.
func (o *WafFirewallVersionDataAttributes) SetRceScoreThreshold(v int32) {
	o.RceScoreThreshold = &v
}

// GetRestrictedExtensions returns the RestrictedExtensions field value if set, zero value otherwise.
func (o *WafFirewallVersionDataAttributes) GetRestrictedExtensions() string {
	if o == nil || o.RestrictedExtensions == nil {
		var ret string
		return ret
	}
	return *o.RestrictedExtensions
}

// GetRestrictedExtensionsOk returns a tuple with the RestrictedExtensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallVersionDataAttributes) GetRestrictedExtensionsOk() (*string, bool) {
	if o == nil || o.RestrictedExtensions == nil {
		return nil, false
	}
	return o.RestrictedExtensions, true
}

// HasRestrictedExtensions returns a boolean if a field has been set.
func (o *WafFirewallVersionDataAttributes) HasRestrictedExtensions() bool {
	if o != nil && o.RestrictedExtensions != nil {
		return true
	}

	return false
}

// SetRestrictedExtensions gets a reference to the given string and assigns it to the RestrictedExtensions field.
func (o *WafFirewallVersionDataAttributes) SetRestrictedExtensions(v string) {
	o.RestrictedExtensions = &v
}

// GetRestrictedHeaders returns the RestrictedHeaders field value if set, zero value otherwise.
func (o *WafFirewallVersionDataAttributes) GetRestrictedHeaders() string {
	if o == nil || o.RestrictedHeaders == nil {
		var ret string
		return ret
	}
	return *o.RestrictedHeaders
}

// GetRestrictedHeadersOk returns a tuple with the RestrictedHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallVersionDataAttributes) GetRestrictedHeadersOk() (*string, bool) {
	if o == nil || o.RestrictedHeaders == nil {
		return nil, false
	}
	return o.RestrictedHeaders, true
}

// HasRestrictedHeaders returns a boolean if a field has been set.
func (o *WafFirewallVersionDataAttributes) HasRestrictedHeaders() bool {
	if o != nil && o.RestrictedHeaders != nil {
		return true
	}

	return false
}

// SetRestrictedHeaders gets a reference to the given string and assigns it to the RestrictedHeaders field.
func (o *WafFirewallVersionDataAttributes) SetRestrictedHeaders(v string) {
	o.RestrictedHeaders = &v
}

// GetRfiScoreThreshold returns the RfiScoreThreshold field value if set, zero value otherwise.
func (o *WafFirewallVersionDataAttributes) GetRfiScoreThreshold() int32 {
	if o == nil || o.RfiScoreThreshold == nil {
		var ret int32
		return ret
	}
	return *o.RfiScoreThreshold
}

// GetRfiScoreThresholdOk returns a tuple with the RfiScoreThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallVersionDataAttributes) GetRfiScoreThresholdOk() (*int32, bool) {
	if o == nil || o.RfiScoreThreshold == nil {
		return nil, false
	}
	return o.RfiScoreThreshold, true
}

// HasRfiScoreThreshold returns a boolean if a field has been set.
func (o *WafFirewallVersionDataAttributes) HasRfiScoreThreshold() bool {
	if o != nil && o.RfiScoreThreshold != nil {
		return true
	}

	return false
}

// SetRfiScoreThreshold gets a reference to the given int32 and assigns it to the RfiScoreThreshold field.
func (o *WafFirewallVersionDataAttributes) SetRfiScoreThreshold(v int32) {
	o.RfiScoreThreshold = &v
}

// GetSessionFixationScoreThreshold returns the SessionFixationScoreThreshold field value if set, zero value otherwise.
func (o *WafFirewallVersionDataAttributes) GetSessionFixationScoreThreshold() int32 {
	if o == nil || o.SessionFixationScoreThreshold == nil {
		var ret int32
		return ret
	}
	return *o.SessionFixationScoreThreshold
}

// GetSessionFixationScoreThresholdOk returns a tuple with the SessionFixationScoreThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallVersionDataAttributes) GetSessionFixationScoreThresholdOk() (*int32, bool) {
	if o == nil || o.SessionFixationScoreThreshold == nil {
		return nil, false
	}
	return o.SessionFixationScoreThreshold, true
}

// HasSessionFixationScoreThreshold returns a boolean if a field has been set.
func (o *WafFirewallVersionDataAttributes) HasSessionFixationScoreThreshold() bool {
	if o != nil && o.SessionFixationScoreThreshold != nil {
		return true
	}

	return false
}

// SetSessionFixationScoreThreshold gets a reference to the given int32 and assigns it to the SessionFixationScoreThreshold field.
func (o *WafFirewallVersionDataAttributes) SetSessionFixationScoreThreshold(v int32) {
	o.SessionFixationScoreThreshold = &v
}

// GetSQLInjectionScoreThreshold returns the SQLInjectionScoreThreshold field value if set, zero value otherwise.
func (o *WafFirewallVersionDataAttributes) GetSQLInjectionScoreThreshold() int32 {
	if o == nil || o.SQLInjectionScoreThreshold == nil {
		var ret int32
		return ret
	}
	return *o.SQLInjectionScoreThreshold
}

// GetSQLInjectionScoreThresholdOk returns a tuple with the SQLInjectionScoreThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallVersionDataAttributes) GetSQLInjectionScoreThresholdOk() (*int32, bool) {
	if o == nil || o.SQLInjectionScoreThreshold == nil {
		return nil, false
	}
	return o.SQLInjectionScoreThreshold, true
}

// HasSQLInjectionScoreThreshold returns a boolean if a field has been set.
func (o *WafFirewallVersionDataAttributes) HasSQLInjectionScoreThreshold() bool {
	if o != nil && o.SQLInjectionScoreThreshold != nil {
		return true
	}

	return false
}

// SetSQLInjectionScoreThreshold gets a reference to the given int32 and assigns it to the SQLInjectionScoreThreshold field.
func (o *WafFirewallVersionDataAttributes) SetSQLInjectionScoreThreshold(v int32) {
	o.SQLInjectionScoreThreshold = &v
}

// GetTotalArgLength returns the TotalArgLength field value if set, zero value otherwise.
func (o *WafFirewallVersionDataAttributes) GetTotalArgLength() int32 {
	if o == nil || o.TotalArgLength == nil {
		var ret int32
		return ret
	}
	return *o.TotalArgLength
}

// GetTotalArgLengthOk returns a tuple with the TotalArgLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallVersionDataAttributes) GetTotalArgLengthOk() (*int32, bool) {
	if o == nil || o.TotalArgLength == nil {
		return nil, false
	}
	return o.TotalArgLength, true
}

// HasTotalArgLength returns a boolean if a field has been set.
func (o *WafFirewallVersionDataAttributes) HasTotalArgLength() bool {
	if o != nil && o.TotalArgLength != nil {
		return true
	}

	return false
}

// SetTotalArgLength gets a reference to the given int32 and assigns it to the TotalArgLength field.
func (o *WafFirewallVersionDataAttributes) SetTotalArgLength(v int32) {
	o.TotalArgLength = &v
}

// GetWarningAnomalyScore returns the WarningAnomalyScore field value if set, zero value otherwise.
func (o *WafFirewallVersionDataAttributes) GetWarningAnomalyScore() int32 {
	if o == nil || o.WarningAnomalyScore == nil {
		var ret int32
		return ret
	}
	return *o.WarningAnomalyScore
}

// GetWarningAnomalyScoreOk returns a tuple with the WarningAnomalyScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallVersionDataAttributes) GetWarningAnomalyScoreOk() (*int32, bool) {
	if o == nil || o.WarningAnomalyScore == nil {
		return nil, false
	}
	return o.WarningAnomalyScore, true
}

// HasWarningAnomalyScore returns a boolean if a field has been set.
func (o *WafFirewallVersionDataAttributes) HasWarningAnomalyScore() bool {
	if o != nil && o.WarningAnomalyScore != nil {
		return true
	}

	return false
}

// SetWarningAnomalyScore gets a reference to the given int32 and assigns it to the WarningAnomalyScore field.
func (o *WafFirewallVersionDataAttributes) SetWarningAnomalyScore(v int32) {
	o.WarningAnomalyScore = &v
}

// GetXSSScoreThreshold returns the XSSScoreThreshold field value if set, zero value otherwise.
func (o *WafFirewallVersionDataAttributes) GetXSSScoreThreshold() int32 {
	if o == nil || o.XSSScoreThreshold == nil {
		var ret int32
		return ret
	}
	return *o.XSSScoreThreshold
}

// GetXSSScoreThresholdOk returns a tuple with the XSSScoreThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallVersionDataAttributes) GetXSSScoreThresholdOk() (*int32, bool) {
	if o == nil || o.XSSScoreThreshold == nil {
		return nil, false
	}
	return o.XSSScoreThreshold, true
}

// HasXSSScoreThreshold returns a boolean if a field has been set.
func (o *WafFirewallVersionDataAttributes) HasXSSScoreThreshold() bool {
	if o != nil && o.XSSScoreThreshold != nil {
		return true
	}

	return false
}

// SetXSSScoreThreshold gets a reference to the given int32 and assigns it to the XSSScoreThreshold field.
func (o *WafFirewallVersionDataAttributes) SetXSSScoreThreshold(v int32) {
	o.XSSScoreThreshold = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o WafFirewallVersionDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.AllowedHTTPVersions != nil {
		toSerialize["allowed_http_versions"] = o.AllowedHTTPVersions
	}
	if o.AllowedMethods != nil {
		toSerialize["allowed_methods"] = o.AllowedMethods
	}
	if o.AllowedRequestContentType != nil {
		toSerialize["allowed_request_content_type"] = o.AllowedRequestContentType
	}
	if o.AllowedRequestContentTypeCharset != nil {
		toSerialize["allowed_request_content_type_charset"] = o.AllowedRequestContentTypeCharset
	}
	if o.ArgNameLength != nil {
		toSerialize["arg_name_length"] = o.ArgNameLength
	}
	if o.ArgLength != nil {
		toSerialize["arg_length"] = o.ArgLength
	}
	if o.CombinedFileSizes != nil {
		toSerialize["combined_file_sizes"] = o.CombinedFileSizes
	}
	if o.Comment.IsSet() {
		toSerialize["comment"] = o.Comment.Get()
	}
	if o.CriticalAnomalyScore != nil {
		toSerialize["critical_anomaly_score"] = o.CriticalAnomalyScore
	}
	if o.CrsValidateUtf8Encoding != nil {
		toSerialize["crs_validate_utf8_encoding"] = o.CrsValidateUtf8Encoding
	}
	if o.ErrorAnomalyScore != nil {
		toSerialize["error_anomaly_score"] = o.ErrorAnomalyScore
	}
	if o.HighRiskCountryCodes != nil {
		toSerialize["high_risk_country_codes"] = o.HighRiskCountryCodes
	}
	if o.HTTPViolationScoreThreshold != nil {
		toSerialize["http_violation_score_threshold"] = o.HTTPViolationScoreThreshold
	}
	if o.InboundAnomalyScoreThreshold != nil {
		toSerialize["inbound_anomaly_score_threshold"] = o.InboundAnomalyScoreThreshold
	}
	if o.LfiScoreThreshold != nil {
		toSerialize["lfi_score_threshold"] = o.LfiScoreThreshold
	}
	if o.Locked != nil {
		toSerialize["locked"] = o.Locked
	}
	if o.MaxFileSize != nil {
		toSerialize["max_file_size"] = o.MaxFileSize
	}
	if o.MaxNumArgs != nil {
		toSerialize["max_num_args"] = o.MaxNumArgs
	}
	if o.NoticeAnomalyScore != nil {
		toSerialize["notice_anomaly_score"] = o.NoticeAnomalyScore
	}
	if o.Number != nil {
		toSerialize["number"] = o.Number
	}
	if o.ParanoiaLevel != nil {
		toSerialize["paranoia_level"] = o.ParanoiaLevel
	}
	if o.PhpInjectionScoreThreshold != nil {
		toSerialize["php_injection_score_threshold"] = o.PhpInjectionScoreThreshold
	}
	if o.RceScoreThreshold != nil {
		toSerialize["rce_score_threshold"] = o.RceScoreThreshold
	}
	if o.RestrictedExtensions != nil {
		toSerialize["restricted_extensions"] = o.RestrictedExtensions
	}
	if o.RestrictedHeaders != nil {
		toSerialize["restricted_headers"] = o.RestrictedHeaders
	}
	if o.RfiScoreThreshold != nil {
		toSerialize["rfi_score_threshold"] = o.RfiScoreThreshold
	}
	if o.SessionFixationScoreThreshold != nil {
		toSerialize["session_fixation_score_threshold"] = o.SessionFixationScoreThreshold
	}
	if o.SQLInjectionScoreThreshold != nil {
		toSerialize["sql_injection_score_threshold"] = o.SQLInjectionScoreThreshold
	}
	if o.TotalArgLength != nil {
		toSerialize["total_arg_length"] = o.TotalArgLength
	}
	if o.WarningAnomalyScore != nil {
		toSerialize["warning_anomaly_score"] = o.WarningAnomalyScore
	}
	if o.XSSScoreThreshold != nil {
		toSerialize["xss_score_threshold"] = o.XSSScoreThreshold
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *WafFirewallVersionDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	varWafFirewallVersionDataAttributes := _WafFirewallVersionDataAttributes{}

	if err = json.Unmarshal(bytes, &varWafFirewallVersionDataAttributes); err == nil {
		*o = WafFirewallVersionDataAttributes(varWafFirewallVersionDataAttributes)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "allowed_http_versions")
		delete(additionalProperties, "allowed_methods")
		delete(additionalProperties, "allowed_request_content_type")
		delete(additionalProperties, "allowed_request_content_type_charset")
		delete(additionalProperties, "arg_name_length")
		delete(additionalProperties, "arg_length")
		delete(additionalProperties, "combined_file_sizes")
		delete(additionalProperties, "comment")
		delete(additionalProperties, "critical_anomaly_score")
		delete(additionalProperties, "crs_validate_utf8_encoding")
		delete(additionalProperties, "error_anomaly_score")
		delete(additionalProperties, "high_risk_country_codes")
		delete(additionalProperties, "http_violation_score_threshold")
		delete(additionalProperties, "inbound_anomaly_score_threshold")
		delete(additionalProperties, "lfi_score_threshold")
		delete(additionalProperties, "locked")
		delete(additionalProperties, "max_file_size")
		delete(additionalProperties, "max_num_args")
		delete(additionalProperties, "notice_anomaly_score")
		delete(additionalProperties, "number")
		delete(additionalProperties, "paranoia_level")
		delete(additionalProperties, "php_injection_score_threshold")
		delete(additionalProperties, "rce_score_threshold")
		delete(additionalProperties, "restricted_extensions")
		delete(additionalProperties, "restricted_headers")
		delete(additionalProperties, "rfi_score_threshold")
		delete(additionalProperties, "session_fixation_score_threshold")
		delete(additionalProperties, "sql_injection_score_threshold")
		delete(additionalProperties, "total_arg_length")
		delete(additionalProperties, "warning_anomaly_score")
		delete(additionalProperties, "xss_score_threshold")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableWafFirewallVersionDataAttributes is a helper abstraction for handling nullable waffirewallversiondataattributes types. 
type NullableWafFirewallVersionDataAttributes struct {
	value *WafFirewallVersionDataAttributes
	isSet bool
}

// Get returns the value.
func (v NullableWafFirewallVersionDataAttributes) Get() *WafFirewallVersionDataAttributes {
	return v.value
}

// Set modifies the value.
func (v *NullableWafFirewallVersionDataAttributes) Set(val *WafFirewallVersionDataAttributes) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableWafFirewallVersionDataAttributes) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableWafFirewallVersionDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableWafFirewallVersionDataAttributes returns a pointer to a new instance of NullableWafFirewallVersionDataAttributes.
func NewNullableWafFirewallVersionDataAttributes(val *WafFirewallVersionDataAttributes) *NullableWafFirewallVersionDataAttributes {
	return &NullableWafFirewallVersionDataAttributes{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableWafFirewallVersionDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableWafFirewallVersionDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
