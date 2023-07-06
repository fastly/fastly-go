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
	"encoding/json"
)

// LegacyWafOwasp struct for LegacyWafOwasp
type LegacyWafOwasp struct {
	// Allowed HTTP versions.
	AllowedHTTPVersions *string `json:"allowed_http_versions,omitempty"`
	// A space-separated list of HTTP method names.
	AllowedMethods *string `json:"allowed_methods,omitempty"`
	// Allowed request content types.
	AllowedRequestContentType *string `json:"allowed_request_content_type,omitempty"`
	// The maximum allowed length of an argument.
	ArgLength *int32 `json:"arg_length,omitempty"`
	// The maximum allowed argument name length.
	ArgNameLength *int32 `json:"arg_name_length,omitempty"`
	// The maximum allowed size of all files (in bytes).
	CombinedFileSizes *int32 `json:"combined_file_sizes,omitempty"`
	// Date and time that the settings object was created.
	CreatedAt *string `json:"created_at,omitempty"`
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
	// The maximum allowed file size (in bytes).
	MaxFileSize *int32 `json:"max_file_size,omitempty"`
	// The maximum number of arguments allowed.
	MaxNumArgs *int32 `json:"max_num_args,omitempty"`
	// Score value to add for notice anomalies.
	NoticeAnomalyScore *int32 `json:"notice_anomaly_score,omitempty"`
	// The configured paranoia level.
	ParanoiaLevel *int32 `json:"paranoia_level,omitempty"`
	// PHP injection threshold.
	PhpInjectionScoreThreshold *int32 `json:"php_injection_score_threshold,omitempty"`
	// Remote code execution threshold.
	RceScoreThreshold *int32 `json:"rce_score_threshold,omitempty"`
	// A space-separated list of disallowed file extensions.
	RestrictedExtensions *string `json:"restricted_extensions,omitempty"`
	// A space-separated list of disallowed header names.
	RestrictedHeaders *string `json:"restricted_headers,omitempty"`
	// Remote file inclusion attack threshold.
	RfiScoreThreshold *int32 `json:"rfi_score_threshold,omitempty"`
	// Session fixation attack threshold.
	SessionFixationScoreThreshold *int32 `json:"session_fixation_score_threshold,omitempty"`
	// SQL injection attack threshold.
	SQLInjectionScoreThreshold *int32 `json:"sql_injection_score_threshold,omitempty"`
	// The maximum size of argument names and values.
	TotalArgLength *int32 `json:"total_arg_length,omitempty"`
	// Date and time that the settings object was last updated.
	UpdatedAt *string `json:"updated_at,omitempty"`
	// Score value to add for warning anomalies.
	WarningAnomalyScore *int32 `json:"warning_anomaly_score,omitempty"`
	// XSS attack threshold.
	XSSScoreThreshold *int32 `json:"xss_score_threshold,omitempty"`
	AdditionalProperties map[string]any
}

type _LegacyWafOwasp LegacyWafOwasp

// NewLegacyWafOwasp instantiates a new LegacyWafOwasp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLegacyWafOwasp() *LegacyWafOwasp {
	this := LegacyWafOwasp{}
	var allowedHTTPVersions string = "HTTP/1.0 HTTP/1.1 HTTP/2"
	this.AllowedHTTPVersions = &allowedHTTPVersions
	var allowedMethods string = "GET HEAD POST OPTIONS PUT PATCH DELETE"
	this.AllowedMethods = &allowedMethods
	var allowedRequestContentType string = "application/x-www-form-urlencoded|multipart/form-data|text/xml|application/xml|application/x-amf|application/json|text/plain"
	this.AllowedRequestContentType = &allowedRequestContentType
	var argLength int32 = 400
	this.ArgLength = &argLength
	var argNameLength int32 = 100
	this.ArgNameLength = &argNameLength
	var combinedFileSizes int32 = 10000000
	this.CombinedFileSizes = &combinedFileSizes
	var criticalAnomalyScore int32 = 6
	this.CriticalAnomalyScore = &criticalAnomalyScore
	var errorAnomalyScore int32 = 5
	this.ErrorAnomalyScore = &errorAnomalyScore
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

// NewLegacyWafOwaspWithDefaults instantiates a new LegacyWafOwasp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLegacyWafOwaspWithDefaults() *LegacyWafOwasp {
	this := LegacyWafOwasp{}
	var allowedHTTPVersions string = "HTTP/1.0 HTTP/1.1 HTTP/2"
	this.AllowedHTTPVersions = &allowedHTTPVersions
	var allowedMethods string = "GET HEAD POST OPTIONS PUT PATCH DELETE"
	this.AllowedMethods = &allowedMethods
	var allowedRequestContentType string = "application/x-www-form-urlencoded|multipart/form-data|text/xml|application/xml|application/x-amf|application/json|text/plain"
	this.AllowedRequestContentType = &allowedRequestContentType
	var argLength int32 = 400
	this.ArgLength = &argLength
	var argNameLength int32 = 100
	this.ArgNameLength = &argNameLength
	var combinedFileSizes int32 = 10000000
	this.CombinedFileSizes = &combinedFileSizes
	var criticalAnomalyScore int32 = 6
	this.CriticalAnomalyScore = &criticalAnomalyScore
	var errorAnomalyScore int32 = 5
	this.ErrorAnomalyScore = &errorAnomalyScore
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
func (o *LegacyWafOwasp) GetAllowedHTTPVersions() string {
	if o == nil || o.AllowedHTTPVersions == nil {
		var ret string
		return ret
	}
	return *o.AllowedHTTPVersions
}

// GetAllowedHTTPVersionsOk returns a tuple with the AllowedHTTPVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyWafOwasp) GetAllowedHTTPVersionsOk() (*string, bool) {
	if o == nil || o.AllowedHTTPVersions == nil {
		return nil, false
	}
	return o.AllowedHTTPVersions, true
}

// HasAllowedHTTPVersions returns a boolean if a field has been set.
func (o *LegacyWafOwasp) HasAllowedHTTPVersions() bool {
	if o != nil && o.AllowedHTTPVersions != nil {
		return true
	}

	return false
}

// SetAllowedHTTPVersions gets a reference to the given string and assigns it to the AllowedHTTPVersions field.
func (o *LegacyWafOwasp) SetAllowedHTTPVersions(v string) {
	o.AllowedHTTPVersions = &v
}

// GetAllowedMethods returns the AllowedMethods field value if set, zero value otherwise.
func (o *LegacyWafOwasp) GetAllowedMethods() string {
	if o == nil || o.AllowedMethods == nil {
		var ret string
		return ret
	}
	return *o.AllowedMethods
}

// GetAllowedMethodsOk returns a tuple with the AllowedMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyWafOwasp) GetAllowedMethodsOk() (*string, bool) {
	if o == nil || o.AllowedMethods == nil {
		return nil, false
	}
	return o.AllowedMethods, true
}

// HasAllowedMethods returns a boolean if a field has been set.
func (o *LegacyWafOwasp) HasAllowedMethods() bool {
	if o != nil && o.AllowedMethods != nil {
		return true
	}

	return false
}

// SetAllowedMethods gets a reference to the given string and assigns it to the AllowedMethods field.
func (o *LegacyWafOwasp) SetAllowedMethods(v string) {
	o.AllowedMethods = &v
}

// GetAllowedRequestContentType returns the AllowedRequestContentType field value if set, zero value otherwise.
func (o *LegacyWafOwasp) GetAllowedRequestContentType() string {
	if o == nil || o.AllowedRequestContentType == nil {
		var ret string
		return ret
	}
	return *o.AllowedRequestContentType
}

// GetAllowedRequestContentTypeOk returns a tuple with the AllowedRequestContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyWafOwasp) GetAllowedRequestContentTypeOk() (*string, bool) {
	if o == nil || o.AllowedRequestContentType == nil {
		return nil, false
	}
	return o.AllowedRequestContentType, true
}

// HasAllowedRequestContentType returns a boolean if a field has been set.
func (o *LegacyWafOwasp) HasAllowedRequestContentType() bool {
	if o != nil && o.AllowedRequestContentType != nil {
		return true
	}

	return false
}

// SetAllowedRequestContentType gets a reference to the given string and assigns it to the AllowedRequestContentType field.
func (o *LegacyWafOwasp) SetAllowedRequestContentType(v string) {
	o.AllowedRequestContentType = &v
}

// GetArgLength returns the ArgLength field value if set, zero value otherwise.
func (o *LegacyWafOwasp) GetArgLength() int32 {
	if o == nil || o.ArgLength == nil {
		var ret int32
		return ret
	}
	return *o.ArgLength
}

// GetArgLengthOk returns a tuple with the ArgLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyWafOwasp) GetArgLengthOk() (*int32, bool) {
	if o == nil || o.ArgLength == nil {
		return nil, false
	}
	return o.ArgLength, true
}

// HasArgLength returns a boolean if a field has been set.
func (o *LegacyWafOwasp) HasArgLength() bool {
	if o != nil && o.ArgLength != nil {
		return true
	}

	return false
}

// SetArgLength gets a reference to the given int32 and assigns it to the ArgLength field.
func (o *LegacyWafOwasp) SetArgLength(v int32) {
	o.ArgLength = &v
}

// GetArgNameLength returns the ArgNameLength field value if set, zero value otherwise.
func (o *LegacyWafOwasp) GetArgNameLength() int32 {
	if o == nil || o.ArgNameLength == nil {
		var ret int32
		return ret
	}
	return *o.ArgNameLength
}

// GetArgNameLengthOk returns a tuple with the ArgNameLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyWafOwasp) GetArgNameLengthOk() (*int32, bool) {
	if o == nil || o.ArgNameLength == nil {
		return nil, false
	}
	return o.ArgNameLength, true
}

// HasArgNameLength returns a boolean if a field has been set.
func (o *LegacyWafOwasp) HasArgNameLength() bool {
	if o != nil && o.ArgNameLength != nil {
		return true
	}

	return false
}

// SetArgNameLength gets a reference to the given int32 and assigns it to the ArgNameLength field.
func (o *LegacyWafOwasp) SetArgNameLength(v int32) {
	o.ArgNameLength = &v
}

// GetCombinedFileSizes returns the CombinedFileSizes field value if set, zero value otherwise.
func (o *LegacyWafOwasp) GetCombinedFileSizes() int32 {
	if o == nil || o.CombinedFileSizes == nil {
		var ret int32
		return ret
	}
	return *o.CombinedFileSizes
}

// GetCombinedFileSizesOk returns a tuple with the CombinedFileSizes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyWafOwasp) GetCombinedFileSizesOk() (*int32, bool) {
	if o == nil || o.CombinedFileSizes == nil {
		return nil, false
	}
	return o.CombinedFileSizes, true
}

// HasCombinedFileSizes returns a boolean if a field has been set.
func (o *LegacyWafOwasp) HasCombinedFileSizes() bool {
	if o != nil && o.CombinedFileSizes != nil {
		return true
	}

	return false
}

// SetCombinedFileSizes gets a reference to the given int32 and assigns it to the CombinedFileSizes field.
func (o *LegacyWafOwasp) SetCombinedFileSizes(v int32) {
	o.CombinedFileSizes = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *LegacyWafOwasp) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyWafOwasp) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *LegacyWafOwasp) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *LegacyWafOwasp) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetCriticalAnomalyScore returns the CriticalAnomalyScore field value if set, zero value otherwise.
func (o *LegacyWafOwasp) GetCriticalAnomalyScore() int32 {
	if o == nil || o.CriticalAnomalyScore == nil {
		var ret int32
		return ret
	}
	return *o.CriticalAnomalyScore
}

// GetCriticalAnomalyScoreOk returns a tuple with the CriticalAnomalyScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyWafOwasp) GetCriticalAnomalyScoreOk() (*int32, bool) {
	if o == nil || o.CriticalAnomalyScore == nil {
		return nil, false
	}
	return o.CriticalAnomalyScore, true
}

// HasCriticalAnomalyScore returns a boolean if a field has been set.
func (o *LegacyWafOwasp) HasCriticalAnomalyScore() bool {
	if o != nil && o.CriticalAnomalyScore != nil {
		return true
	}

	return false
}

// SetCriticalAnomalyScore gets a reference to the given int32 and assigns it to the CriticalAnomalyScore field.
func (o *LegacyWafOwasp) SetCriticalAnomalyScore(v int32) {
	o.CriticalAnomalyScore = &v
}

// GetCrsValidateUtf8Encoding returns the CrsValidateUtf8Encoding field value if set, zero value otherwise.
func (o *LegacyWafOwasp) GetCrsValidateUtf8Encoding() bool {
	if o == nil || o.CrsValidateUtf8Encoding == nil {
		var ret bool
		return ret
	}
	return *o.CrsValidateUtf8Encoding
}

// GetCrsValidateUtf8EncodingOk returns a tuple with the CrsValidateUtf8Encoding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyWafOwasp) GetCrsValidateUtf8EncodingOk() (*bool, bool) {
	if o == nil || o.CrsValidateUtf8Encoding == nil {
		return nil, false
	}
	return o.CrsValidateUtf8Encoding, true
}

// HasCrsValidateUtf8Encoding returns a boolean if a field has been set.
func (o *LegacyWafOwasp) HasCrsValidateUtf8Encoding() bool {
	if o != nil && o.CrsValidateUtf8Encoding != nil {
		return true
	}

	return false
}

// SetCrsValidateUtf8Encoding gets a reference to the given bool and assigns it to the CrsValidateUtf8Encoding field.
func (o *LegacyWafOwasp) SetCrsValidateUtf8Encoding(v bool) {
	o.CrsValidateUtf8Encoding = &v
}

// GetErrorAnomalyScore returns the ErrorAnomalyScore field value if set, zero value otherwise.
func (o *LegacyWafOwasp) GetErrorAnomalyScore() int32 {
	if o == nil || o.ErrorAnomalyScore == nil {
		var ret int32
		return ret
	}
	return *o.ErrorAnomalyScore
}

// GetErrorAnomalyScoreOk returns a tuple with the ErrorAnomalyScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyWafOwasp) GetErrorAnomalyScoreOk() (*int32, bool) {
	if o == nil || o.ErrorAnomalyScore == nil {
		return nil, false
	}
	return o.ErrorAnomalyScore, true
}

// HasErrorAnomalyScore returns a boolean if a field has been set.
func (o *LegacyWafOwasp) HasErrorAnomalyScore() bool {
	if o != nil && o.ErrorAnomalyScore != nil {
		return true
	}

	return false
}

// SetErrorAnomalyScore gets a reference to the given int32 and assigns it to the ErrorAnomalyScore field.
func (o *LegacyWafOwasp) SetErrorAnomalyScore(v int32) {
	o.ErrorAnomalyScore = &v
}

// GetHighRiskCountryCodes returns the HighRiskCountryCodes field value if set, zero value otherwise.
func (o *LegacyWafOwasp) GetHighRiskCountryCodes() string {
	if o == nil || o.HighRiskCountryCodes == nil {
		var ret string
		return ret
	}
	return *o.HighRiskCountryCodes
}

// GetHighRiskCountryCodesOk returns a tuple with the HighRiskCountryCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyWafOwasp) GetHighRiskCountryCodesOk() (*string, bool) {
	if o == nil || o.HighRiskCountryCodes == nil {
		return nil, false
	}
	return o.HighRiskCountryCodes, true
}

// HasHighRiskCountryCodes returns a boolean if a field has been set.
func (o *LegacyWafOwasp) HasHighRiskCountryCodes() bool {
	if o != nil && o.HighRiskCountryCodes != nil {
		return true
	}

	return false
}

// SetHighRiskCountryCodes gets a reference to the given string and assigns it to the HighRiskCountryCodes field.
func (o *LegacyWafOwasp) SetHighRiskCountryCodes(v string) {
	o.HighRiskCountryCodes = &v
}

// GetHTTPViolationScoreThreshold returns the HTTPViolationScoreThreshold field value if set, zero value otherwise.
func (o *LegacyWafOwasp) GetHTTPViolationScoreThreshold() int32 {
	if o == nil || o.HTTPViolationScoreThreshold == nil {
		var ret int32
		return ret
	}
	return *o.HTTPViolationScoreThreshold
}

// GetHTTPViolationScoreThresholdOk returns a tuple with the HTTPViolationScoreThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyWafOwasp) GetHTTPViolationScoreThresholdOk() (*int32, bool) {
	if o == nil || o.HTTPViolationScoreThreshold == nil {
		return nil, false
	}
	return o.HTTPViolationScoreThreshold, true
}

// HasHTTPViolationScoreThreshold returns a boolean if a field has been set.
func (o *LegacyWafOwasp) HasHTTPViolationScoreThreshold() bool {
	if o != nil && o.HTTPViolationScoreThreshold != nil {
		return true
	}

	return false
}

// SetHTTPViolationScoreThreshold gets a reference to the given int32 and assigns it to the HTTPViolationScoreThreshold field.
func (o *LegacyWafOwasp) SetHTTPViolationScoreThreshold(v int32) {
	o.HTTPViolationScoreThreshold = &v
}

// GetInboundAnomalyScoreThreshold returns the InboundAnomalyScoreThreshold field value if set, zero value otherwise.
func (o *LegacyWafOwasp) GetInboundAnomalyScoreThreshold() int32 {
	if o == nil || o.InboundAnomalyScoreThreshold == nil {
		var ret int32
		return ret
	}
	return *o.InboundAnomalyScoreThreshold
}

// GetInboundAnomalyScoreThresholdOk returns a tuple with the InboundAnomalyScoreThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyWafOwasp) GetInboundAnomalyScoreThresholdOk() (*int32, bool) {
	if o == nil || o.InboundAnomalyScoreThreshold == nil {
		return nil, false
	}
	return o.InboundAnomalyScoreThreshold, true
}

// HasInboundAnomalyScoreThreshold returns a boolean if a field has been set.
func (o *LegacyWafOwasp) HasInboundAnomalyScoreThreshold() bool {
	if o != nil && o.InboundAnomalyScoreThreshold != nil {
		return true
	}

	return false
}

// SetInboundAnomalyScoreThreshold gets a reference to the given int32 and assigns it to the InboundAnomalyScoreThreshold field.
func (o *LegacyWafOwasp) SetInboundAnomalyScoreThreshold(v int32) {
	o.InboundAnomalyScoreThreshold = &v
}

// GetLfiScoreThreshold returns the LfiScoreThreshold field value if set, zero value otherwise.
func (o *LegacyWafOwasp) GetLfiScoreThreshold() int32 {
	if o == nil || o.LfiScoreThreshold == nil {
		var ret int32
		return ret
	}
	return *o.LfiScoreThreshold
}

// GetLfiScoreThresholdOk returns a tuple with the LfiScoreThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyWafOwasp) GetLfiScoreThresholdOk() (*int32, bool) {
	if o == nil || o.LfiScoreThreshold == nil {
		return nil, false
	}
	return o.LfiScoreThreshold, true
}

// HasLfiScoreThreshold returns a boolean if a field has been set.
func (o *LegacyWafOwasp) HasLfiScoreThreshold() bool {
	if o != nil && o.LfiScoreThreshold != nil {
		return true
	}

	return false
}

// SetLfiScoreThreshold gets a reference to the given int32 and assigns it to the LfiScoreThreshold field.
func (o *LegacyWafOwasp) SetLfiScoreThreshold(v int32) {
	o.LfiScoreThreshold = &v
}

// GetMaxFileSize returns the MaxFileSize field value if set, zero value otherwise.
func (o *LegacyWafOwasp) GetMaxFileSize() int32 {
	if o == nil || o.MaxFileSize == nil {
		var ret int32
		return ret
	}
	return *o.MaxFileSize
}

// GetMaxFileSizeOk returns a tuple with the MaxFileSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyWafOwasp) GetMaxFileSizeOk() (*int32, bool) {
	if o == nil || o.MaxFileSize == nil {
		return nil, false
	}
	return o.MaxFileSize, true
}

// HasMaxFileSize returns a boolean if a field has been set.
func (o *LegacyWafOwasp) HasMaxFileSize() bool {
	if o != nil && o.MaxFileSize != nil {
		return true
	}

	return false
}

// SetMaxFileSize gets a reference to the given int32 and assigns it to the MaxFileSize field.
func (o *LegacyWafOwasp) SetMaxFileSize(v int32) {
	o.MaxFileSize = &v
}

// GetMaxNumArgs returns the MaxNumArgs field value if set, zero value otherwise.
func (o *LegacyWafOwasp) GetMaxNumArgs() int32 {
	if o == nil || o.MaxNumArgs == nil {
		var ret int32
		return ret
	}
	return *o.MaxNumArgs
}

// GetMaxNumArgsOk returns a tuple with the MaxNumArgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyWafOwasp) GetMaxNumArgsOk() (*int32, bool) {
	if o == nil || o.MaxNumArgs == nil {
		return nil, false
	}
	return o.MaxNumArgs, true
}

// HasMaxNumArgs returns a boolean if a field has been set.
func (o *LegacyWafOwasp) HasMaxNumArgs() bool {
	if o != nil && o.MaxNumArgs != nil {
		return true
	}

	return false
}

// SetMaxNumArgs gets a reference to the given int32 and assigns it to the MaxNumArgs field.
func (o *LegacyWafOwasp) SetMaxNumArgs(v int32) {
	o.MaxNumArgs = &v
}

// GetNoticeAnomalyScore returns the NoticeAnomalyScore field value if set, zero value otherwise.
func (o *LegacyWafOwasp) GetNoticeAnomalyScore() int32 {
	if o == nil || o.NoticeAnomalyScore == nil {
		var ret int32
		return ret
	}
	return *o.NoticeAnomalyScore
}

// GetNoticeAnomalyScoreOk returns a tuple with the NoticeAnomalyScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyWafOwasp) GetNoticeAnomalyScoreOk() (*int32, bool) {
	if o == nil || o.NoticeAnomalyScore == nil {
		return nil, false
	}
	return o.NoticeAnomalyScore, true
}

// HasNoticeAnomalyScore returns a boolean if a field has been set.
func (o *LegacyWafOwasp) HasNoticeAnomalyScore() bool {
	if o != nil && o.NoticeAnomalyScore != nil {
		return true
	}

	return false
}

// SetNoticeAnomalyScore gets a reference to the given int32 and assigns it to the NoticeAnomalyScore field.
func (o *LegacyWafOwasp) SetNoticeAnomalyScore(v int32) {
	o.NoticeAnomalyScore = &v
}

// GetParanoiaLevel returns the ParanoiaLevel field value if set, zero value otherwise.
func (o *LegacyWafOwasp) GetParanoiaLevel() int32 {
	if o == nil || o.ParanoiaLevel == nil {
		var ret int32
		return ret
	}
	return *o.ParanoiaLevel
}

// GetParanoiaLevelOk returns a tuple with the ParanoiaLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyWafOwasp) GetParanoiaLevelOk() (*int32, bool) {
	if o == nil || o.ParanoiaLevel == nil {
		return nil, false
	}
	return o.ParanoiaLevel, true
}

// HasParanoiaLevel returns a boolean if a field has been set.
func (o *LegacyWafOwasp) HasParanoiaLevel() bool {
	if o != nil && o.ParanoiaLevel != nil {
		return true
	}

	return false
}

// SetParanoiaLevel gets a reference to the given int32 and assigns it to the ParanoiaLevel field.
func (o *LegacyWafOwasp) SetParanoiaLevel(v int32) {
	o.ParanoiaLevel = &v
}

// GetPhpInjectionScoreThreshold returns the PhpInjectionScoreThreshold field value if set, zero value otherwise.
func (o *LegacyWafOwasp) GetPhpInjectionScoreThreshold() int32 {
	if o == nil || o.PhpInjectionScoreThreshold == nil {
		var ret int32
		return ret
	}
	return *o.PhpInjectionScoreThreshold
}

// GetPhpInjectionScoreThresholdOk returns a tuple with the PhpInjectionScoreThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyWafOwasp) GetPhpInjectionScoreThresholdOk() (*int32, bool) {
	if o == nil || o.PhpInjectionScoreThreshold == nil {
		return nil, false
	}
	return o.PhpInjectionScoreThreshold, true
}

// HasPhpInjectionScoreThreshold returns a boolean if a field has been set.
func (o *LegacyWafOwasp) HasPhpInjectionScoreThreshold() bool {
	if o != nil && o.PhpInjectionScoreThreshold != nil {
		return true
	}

	return false
}

// SetPhpInjectionScoreThreshold gets a reference to the given int32 and assigns it to the PhpInjectionScoreThreshold field.
func (o *LegacyWafOwasp) SetPhpInjectionScoreThreshold(v int32) {
	o.PhpInjectionScoreThreshold = &v
}

// GetRceScoreThreshold returns the RceScoreThreshold field value if set, zero value otherwise.
func (o *LegacyWafOwasp) GetRceScoreThreshold() int32 {
	if o == nil || o.RceScoreThreshold == nil {
		var ret int32
		return ret
	}
	return *o.RceScoreThreshold
}

// GetRceScoreThresholdOk returns a tuple with the RceScoreThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyWafOwasp) GetRceScoreThresholdOk() (*int32, bool) {
	if o == nil || o.RceScoreThreshold == nil {
		return nil, false
	}
	return o.RceScoreThreshold, true
}

// HasRceScoreThreshold returns a boolean if a field has been set.
func (o *LegacyWafOwasp) HasRceScoreThreshold() bool {
	if o != nil && o.RceScoreThreshold != nil {
		return true
	}

	return false
}

// SetRceScoreThreshold gets a reference to the given int32 and assigns it to the RceScoreThreshold field.
func (o *LegacyWafOwasp) SetRceScoreThreshold(v int32) {
	o.RceScoreThreshold = &v
}

// GetRestrictedExtensions returns the RestrictedExtensions field value if set, zero value otherwise.
func (o *LegacyWafOwasp) GetRestrictedExtensions() string {
	if o == nil || o.RestrictedExtensions == nil {
		var ret string
		return ret
	}
	return *o.RestrictedExtensions
}

// GetRestrictedExtensionsOk returns a tuple with the RestrictedExtensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyWafOwasp) GetRestrictedExtensionsOk() (*string, bool) {
	if o == nil || o.RestrictedExtensions == nil {
		return nil, false
	}
	return o.RestrictedExtensions, true
}

// HasRestrictedExtensions returns a boolean if a field has been set.
func (o *LegacyWafOwasp) HasRestrictedExtensions() bool {
	if o != nil && o.RestrictedExtensions != nil {
		return true
	}

	return false
}

// SetRestrictedExtensions gets a reference to the given string and assigns it to the RestrictedExtensions field.
func (o *LegacyWafOwasp) SetRestrictedExtensions(v string) {
	o.RestrictedExtensions = &v
}

// GetRestrictedHeaders returns the RestrictedHeaders field value if set, zero value otherwise.
func (o *LegacyWafOwasp) GetRestrictedHeaders() string {
	if o == nil || o.RestrictedHeaders == nil {
		var ret string
		return ret
	}
	return *o.RestrictedHeaders
}

// GetRestrictedHeadersOk returns a tuple with the RestrictedHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyWafOwasp) GetRestrictedHeadersOk() (*string, bool) {
	if o == nil || o.RestrictedHeaders == nil {
		return nil, false
	}
	return o.RestrictedHeaders, true
}

// HasRestrictedHeaders returns a boolean if a field has been set.
func (o *LegacyWafOwasp) HasRestrictedHeaders() bool {
	if o != nil && o.RestrictedHeaders != nil {
		return true
	}

	return false
}

// SetRestrictedHeaders gets a reference to the given string and assigns it to the RestrictedHeaders field.
func (o *LegacyWafOwasp) SetRestrictedHeaders(v string) {
	o.RestrictedHeaders = &v
}

// GetRfiScoreThreshold returns the RfiScoreThreshold field value if set, zero value otherwise.
func (o *LegacyWafOwasp) GetRfiScoreThreshold() int32 {
	if o == nil || o.RfiScoreThreshold == nil {
		var ret int32
		return ret
	}
	return *o.RfiScoreThreshold
}

// GetRfiScoreThresholdOk returns a tuple with the RfiScoreThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyWafOwasp) GetRfiScoreThresholdOk() (*int32, bool) {
	if o == nil || o.RfiScoreThreshold == nil {
		return nil, false
	}
	return o.RfiScoreThreshold, true
}

// HasRfiScoreThreshold returns a boolean if a field has been set.
func (o *LegacyWafOwasp) HasRfiScoreThreshold() bool {
	if o != nil && o.RfiScoreThreshold != nil {
		return true
	}

	return false
}

// SetRfiScoreThreshold gets a reference to the given int32 and assigns it to the RfiScoreThreshold field.
func (o *LegacyWafOwasp) SetRfiScoreThreshold(v int32) {
	o.RfiScoreThreshold = &v
}

// GetSessionFixationScoreThreshold returns the SessionFixationScoreThreshold field value if set, zero value otherwise.
func (o *LegacyWafOwasp) GetSessionFixationScoreThreshold() int32 {
	if o == nil || o.SessionFixationScoreThreshold == nil {
		var ret int32
		return ret
	}
	return *o.SessionFixationScoreThreshold
}

// GetSessionFixationScoreThresholdOk returns a tuple with the SessionFixationScoreThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyWafOwasp) GetSessionFixationScoreThresholdOk() (*int32, bool) {
	if o == nil || o.SessionFixationScoreThreshold == nil {
		return nil, false
	}
	return o.SessionFixationScoreThreshold, true
}

// HasSessionFixationScoreThreshold returns a boolean if a field has been set.
func (o *LegacyWafOwasp) HasSessionFixationScoreThreshold() bool {
	if o != nil && o.SessionFixationScoreThreshold != nil {
		return true
	}

	return false
}

// SetSessionFixationScoreThreshold gets a reference to the given int32 and assigns it to the SessionFixationScoreThreshold field.
func (o *LegacyWafOwasp) SetSessionFixationScoreThreshold(v int32) {
	o.SessionFixationScoreThreshold = &v
}

// GetSQLInjectionScoreThreshold returns the SQLInjectionScoreThreshold field value if set, zero value otherwise.
func (o *LegacyWafOwasp) GetSQLInjectionScoreThreshold() int32 {
	if o == nil || o.SQLInjectionScoreThreshold == nil {
		var ret int32
		return ret
	}
	return *o.SQLInjectionScoreThreshold
}

// GetSQLInjectionScoreThresholdOk returns a tuple with the SQLInjectionScoreThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyWafOwasp) GetSQLInjectionScoreThresholdOk() (*int32, bool) {
	if o == nil || o.SQLInjectionScoreThreshold == nil {
		return nil, false
	}
	return o.SQLInjectionScoreThreshold, true
}

// HasSQLInjectionScoreThreshold returns a boolean if a field has been set.
func (o *LegacyWafOwasp) HasSQLInjectionScoreThreshold() bool {
	if o != nil && o.SQLInjectionScoreThreshold != nil {
		return true
	}

	return false
}

// SetSQLInjectionScoreThreshold gets a reference to the given int32 and assigns it to the SQLInjectionScoreThreshold field.
func (o *LegacyWafOwasp) SetSQLInjectionScoreThreshold(v int32) {
	o.SQLInjectionScoreThreshold = &v
}

// GetTotalArgLength returns the TotalArgLength field value if set, zero value otherwise.
func (o *LegacyWafOwasp) GetTotalArgLength() int32 {
	if o == nil || o.TotalArgLength == nil {
		var ret int32
		return ret
	}
	return *o.TotalArgLength
}

// GetTotalArgLengthOk returns a tuple with the TotalArgLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyWafOwasp) GetTotalArgLengthOk() (*int32, bool) {
	if o == nil || o.TotalArgLength == nil {
		return nil, false
	}
	return o.TotalArgLength, true
}

// HasTotalArgLength returns a boolean if a field has been set.
func (o *LegacyWafOwasp) HasTotalArgLength() bool {
	if o != nil && o.TotalArgLength != nil {
		return true
	}

	return false
}

// SetTotalArgLength gets a reference to the given int32 and assigns it to the TotalArgLength field.
func (o *LegacyWafOwasp) SetTotalArgLength(v int32) {
	o.TotalArgLength = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *LegacyWafOwasp) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyWafOwasp) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *LegacyWafOwasp) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *LegacyWafOwasp) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetWarningAnomalyScore returns the WarningAnomalyScore field value if set, zero value otherwise.
func (o *LegacyWafOwasp) GetWarningAnomalyScore() int32 {
	if o == nil || o.WarningAnomalyScore == nil {
		var ret int32
		return ret
	}
	return *o.WarningAnomalyScore
}

// GetWarningAnomalyScoreOk returns a tuple with the WarningAnomalyScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyWafOwasp) GetWarningAnomalyScoreOk() (*int32, bool) {
	if o == nil || o.WarningAnomalyScore == nil {
		return nil, false
	}
	return o.WarningAnomalyScore, true
}

// HasWarningAnomalyScore returns a boolean if a field has been set.
func (o *LegacyWafOwasp) HasWarningAnomalyScore() bool {
	if o != nil && o.WarningAnomalyScore != nil {
		return true
	}

	return false
}

// SetWarningAnomalyScore gets a reference to the given int32 and assigns it to the WarningAnomalyScore field.
func (o *LegacyWafOwasp) SetWarningAnomalyScore(v int32) {
	o.WarningAnomalyScore = &v
}

// GetXSSScoreThreshold returns the XSSScoreThreshold field value if set, zero value otherwise.
func (o *LegacyWafOwasp) GetXSSScoreThreshold() int32 {
	if o == nil || o.XSSScoreThreshold == nil {
		var ret int32
		return ret
	}
	return *o.XSSScoreThreshold
}

// GetXSSScoreThresholdOk returns a tuple with the XSSScoreThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyWafOwasp) GetXSSScoreThresholdOk() (*int32, bool) {
	if o == nil || o.XSSScoreThreshold == nil {
		return nil, false
	}
	return o.XSSScoreThreshold, true
}

// HasXSSScoreThreshold returns a boolean if a field has been set.
func (o *LegacyWafOwasp) HasXSSScoreThreshold() bool {
	if o != nil && o.XSSScoreThreshold != nil {
		return true
	}

	return false
}

// SetXSSScoreThreshold gets a reference to the given int32 and assigns it to the XSSScoreThreshold field.
func (o *LegacyWafOwasp) SetXSSScoreThreshold(v int32) {
	o.XSSScoreThreshold = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o LegacyWafOwasp) MarshalJSON() ([]byte, error) {
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
	if o.ArgLength != nil {
		toSerialize["arg_length"] = o.ArgLength
	}
	if o.ArgNameLength != nil {
		toSerialize["arg_name_length"] = o.ArgNameLength
	}
	if o.CombinedFileSizes != nil {
		toSerialize["combined_file_sizes"] = o.CombinedFileSizes
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
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
	if o.MaxFileSize != nil {
		toSerialize["max_file_size"] = o.MaxFileSize
	}
	if o.MaxNumArgs != nil {
		toSerialize["max_num_args"] = o.MaxNumArgs
	}
	if o.NoticeAnomalyScore != nil {
		toSerialize["notice_anomaly_score"] = o.NoticeAnomalyScore
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
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
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
func (o *LegacyWafOwasp) UnmarshalJSON(bytes []byte) (err error) {
	varLegacyWafOwasp := _LegacyWafOwasp{}

	if err = json.Unmarshal(bytes, &varLegacyWafOwasp); err == nil {
		*o = LegacyWafOwasp(varLegacyWafOwasp)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "allowed_http_versions")
		delete(additionalProperties, "allowed_methods")
		delete(additionalProperties, "allowed_request_content_type")
		delete(additionalProperties, "arg_length")
		delete(additionalProperties, "arg_name_length")
		delete(additionalProperties, "combined_file_sizes")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "critical_anomaly_score")
		delete(additionalProperties, "crs_validate_utf8_encoding")
		delete(additionalProperties, "error_anomaly_score")
		delete(additionalProperties, "high_risk_country_codes")
		delete(additionalProperties, "http_violation_score_threshold")
		delete(additionalProperties, "inbound_anomaly_score_threshold")
		delete(additionalProperties, "lfi_score_threshold")
		delete(additionalProperties, "max_file_size")
		delete(additionalProperties, "max_num_args")
		delete(additionalProperties, "notice_anomaly_score")
		delete(additionalProperties, "paranoia_level")
		delete(additionalProperties, "php_injection_score_threshold")
		delete(additionalProperties, "rce_score_threshold")
		delete(additionalProperties, "restricted_extensions")
		delete(additionalProperties, "restricted_headers")
		delete(additionalProperties, "rfi_score_threshold")
		delete(additionalProperties, "session_fixation_score_threshold")
		delete(additionalProperties, "sql_injection_score_threshold")
		delete(additionalProperties, "total_arg_length")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "warning_anomaly_score")
		delete(additionalProperties, "xss_score_threshold")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableLegacyWafOwasp is a helper abstraction for handling nullable legacywafowasp types. 
type NullableLegacyWafOwasp struct {
	value *LegacyWafOwasp
	isSet bool
}

// Get returns the value.
func (v NullableLegacyWafOwasp) Get() *LegacyWafOwasp {
	return v.value
}

// Set modifies the value.
func (v *NullableLegacyWafOwasp) Set(val *LegacyWafOwasp) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableLegacyWafOwasp) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableLegacyWafOwasp) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableLegacyWafOwasp returns a pointer to a new instance of NullableLegacyWafOwasp.
func NewNullableLegacyWafOwasp(val *LegacyWafOwasp) *NullableLegacyWafOwasp {
	return &NullableLegacyWafOwasp{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableLegacyWafOwasp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableLegacyWafOwasp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
