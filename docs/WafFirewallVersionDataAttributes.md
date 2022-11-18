# WafFirewallVersionDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedHTTPVersions** | Pointer to **string** | Allowed HTTP versions. | [optional] [default to "HTTP/1.0 HTTP/1.1 HTTP/2"]
**AllowedMethods** | Pointer to **string** | A space-separated list of HTTP method names. | [optional] [default to "GET HEAD POST OPTIONS PUT PATCH DELETE"]
**AllowedRequestContentType** | Pointer to **string** | Allowed request content types. | [optional] [default to "application/x-www-form-urlencoded|multipart/form-data|text/xml|application/xml|application/x-amf|application/json|text/plain"]
**AllowedRequestContentTypeCharset** | Pointer to **string** | Allowed request content type charset. | [optional] [default to "utf-8|iso-8859-1|iso-8859-15|windows-1252"]
**ArgNameLength** | Pointer to **int32** | The maximum allowed argument name length. | [optional] [default to 100]
**ArgLength** | Pointer to **int32** | The maximum allowed length of an argument. | [optional] [default to 400]
**CombinedFileSizes** | Pointer to **int32** | The maximum allowed size of all files (in bytes). | [optional] [default to 10000000]
**Comment** | Pointer to **NullableString** | A freeform descriptive note. | [optional] 
**CriticalAnomalyScore** | Pointer to **int32** | Score value to add for critical anomalies. | [optional] [default to 6]
**CrsValidateUtf8Encoding** | Pointer to **bool** | CRS validate UTF8 encoding. | [optional] 
**ErrorAnomalyScore** | Pointer to **int32** | Score value to add for error anomalies. | [optional] [default to 5]
**HighRiskCountryCodes** | Pointer to **string** | A space-separated list of country codes in ISO 3166-1 (two-letter) format. | [optional] 
**HTTPViolationScoreThreshold** | Pointer to **int32** | HTTP violation threshold. | [optional] 
**InboundAnomalyScoreThreshold** | Pointer to **int32** | Inbound anomaly threshold. | [optional] 
**LfiScoreThreshold** | Pointer to **int32** | Local file inclusion attack threshold. | [optional] 
**Locked** | Pointer to **bool** | Whether a specific firewall version is locked from being modified. | [optional] [default to false]
**MaxFileSize** | Pointer to **int32** | The maximum allowed file size, in bytes. | [optional] [default to 10000000]
**MaxNumArgs** | Pointer to **int32** | The maximum number of arguments allowed. | [optional] [default to 255]
**NoticeAnomalyScore** | Pointer to **int32** | Score value to add for notice anomalies. | [optional] [default to 4]
**Number** | Pointer to **int32** |  | [optional] [readonly] 
**ParanoiaLevel** | Pointer to **int32** | The configured paranoia level. | [optional] [default to 1]
**PhpInjectionScoreThreshold** | Pointer to **int32** | PHP injection threshold. | [optional] 
**RceScoreThreshold** | Pointer to **int32** | Remote code execution threshold. | [optional] 
**RestrictedExtensions** | Pointer to **string** | A space-separated list of allowed file extensions. | [optional] [default to ".asa/ .asax/ .ascx/ .axd/ .backup/ .bak/ .bat/ .cdx/ .cer/ .cfg/ .cmd/ .com/ .config/ .conf/ .cs/ .csproj/ .csr/ .dat/ .db/ .dbf/ .dll/ .dos/ .htr/ .htw/ .ida/ .idc/ .idq/ .inc/ .ini/ .key/ .licx/ .lnk/ .log/ .mdb/ .old/ .pass/ .pdb/ .pol/ .printer/ .pwd/ .resources/ .resx/ .sql/ .sys/ .vb/ .vbs/ .vbproj/ .vsdisco/ .webinfo/ .xsd/ .xsx"]
**RestrictedHeaders** | Pointer to **string** | A space-separated list of allowed header names. | [optional] [default to "/proxy/ /lock-token/ /content-range/ /translate/ /if/"]
**RfiScoreThreshold** | Pointer to **int32** | Remote file inclusion attack threshold. | [optional] 
**SessionFixationScoreThreshold** | Pointer to **int32** | Session fixation attack threshold. | [optional] 
**SQLInjectionScoreThreshold** | Pointer to **int32** | SQL injection attack threshold. | [optional] 
**TotalArgLength** | Pointer to **int32** | The maximum size of argument names and values. | [optional] [default to 6400]
**WarningAnomalyScore** | Pointer to **int32** | Score value to add for warning anomalies. | [optional] 
**XSSScoreThreshold** | Pointer to **int32** | XSS attack threshold. | [optional] 

## Methods

### NewWafFirewallVersionDataAttributes

`func NewWafFirewallVersionDataAttributes() *WafFirewallVersionDataAttributes`

NewWafFirewallVersionDataAttributes instantiates a new WafFirewallVersionDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafFirewallVersionDataAttributesWithDefaults

`func NewWafFirewallVersionDataAttributesWithDefaults() *WafFirewallVersionDataAttributes`

NewWafFirewallVersionDataAttributesWithDefaults instantiates a new WafFirewallVersionDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedHTTPVersions

`func (o *WafFirewallVersionDataAttributes) GetAllowedHTTPVersions() string`

GetAllowedHTTPVersions returns the AllowedHTTPVersions field if non-nil, zero value otherwise.

### GetAllowedHTTPVersionsOk

`func (o *WafFirewallVersionDataAttributes) GetAllowedHTTPVersionsOk() (*string, bool)`

GetAllowedHTTPVersionsOk returns a tuple with the AllowedHTTPVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedHTTPVersions

`func (o *WafFirewallVersionDataAttributes) SetAllowedHTTPVersions(v string)`

SetAllowedHTTPVersions sets AllowedHTTPVersions field to given value.

### HasAllowedHTTPVersions

`func (o *WafFirewallVersionDataAttributes) HasAllowedHTTPVersions() bool`

HasAllowedHTTPVersions returns a boolean if a field has been set.

### GetAllowedMethods

`func (o *WafFirewallVersionDataAttributes) GetAllowedMethods() string`

GetAllowedMethods returns the AllowedMethods field if non-nil, zero value otherwise.

### GetAllowedMethodsOk

`func (o *WafFirewallVersionDataAttributes) GetAllowedMethodsOk() (*string, bool)`

GetAllowedMethodsOk returns a tuple with the AllowedMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedMethods

`func (o *WafFirewallVersionDataAttributes) SetAllowedMethods(v string)`

SetAllowedMethods sets AllowedMethods field to given value.

### HasAllowedMethods

`func (o *WafFirewallVersionDataAttributes) HasAllowedMethods() bool`

HasAllowedMethods returns a boolean if a field has been set.

### GetAllowedRequestContentType

`func (o *WafFirewallVersionDataAttributes) GetAllowedRequestContentType() string`

GetAllowedRequestContentType returns the AllowedRequestContentType field if non-nil, zero value otherwise.

### GetAllowedRequestContentTypeOk

`func (o *WafFirewallVersionDataAttributes) GetAllowedRequestContentTypeOk() (*string, bool)`

GetAllowedRequestContentTypeOk returns a tuple with the AllowedRequestContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedRequestContentType

`func (o *WafFirewallVersionDataAttributes) SetAllowedRequestContentType(v string)`

SetAllowedRequestContentType sets AllowedRequestContentType field to given value.

### HasAllowedRequestContentType

`func (o *WafFirewallVersionDataAttributes) HasAllowedRequestContentType() bool`

HasAllowedRequestContentType returns a boolean if a field has been set.

### GetAllowedRequestContentTypeCharset

`func (o *WafFirewallVersionDataAttributes) GetAllowedRequestContentTypeCharset() string`

GetAllowedRequestContentTypeCharset returns the AllowedRequestContentTypeCharset field if non-nil, zero value otherwise.

### GetAllowedRequestContentTypeCharsetOk

`func (o *WafFirewallVersionDataAttributes) GetAllowedRequestContentTypeCharsetOk() (*string, bool)`

GetAllowedRequestContentTypeCharsetOk returns a tuple with the AllowedRequestContentTypeCharset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedRequestContentTypeCharset

`func (o *WafFirewallVersionDataAttributes) SetAllowedRequestContentTypeCharset(v string)`

SetAllowedRequestContentTypeCharset sets AllowedRequestContentTypeCharset field to given value.

### HasAllowedRequestContentTypeCharset

`func (o *WafFirewallVersionDataAttributes) HasAllowedRequestContentTypeCharset() bool`

HasAllowedRequestContentTypeCharset returns a boolean if a field has been set.

### GetArgNameLength

`func (o *WafFirewallVersionDataAttributes) GetArgNameLength() int32`

GetArgNameLength returns the ArgNameLength field if non-nil, zero value otherwise.

### GetArgNameLengthOk

`func (o *WafFirewallVersionDataAttributes) GetArgNameLengthOk() (*int32, bool)`

GetArgNameLengthOk returns a tuple with the ArgNameLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgNameLength

`func (o *WafFirewallVersionDataAttributes) SetArgNameLength(v int32)`

SetArgNameLength sets ArgNameLength field to given value.

### HasArgNameLength

`func (o *WafFirewallVersionDataAttributes) HasArgNameLength() bool`

HasArgNameLength returns a boolean if a field has been set.

### GetArgLength

`func (o *WafFirewallVersionDataAttributes) GetArgLength() int32`

GetArgLength returns the ArgLength field if non-nil, zero value otherwise.

### GetArgLengthOk

`func (o *WafFirewallVersionDataAttributes) GetArgLengthOk() (*int32, bool)`

GetArgLengthOk returns a tuple with the ArgLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgLength

`func (o *WafFirewallVersionDataAttributes) SetArgLength(v int32)`

SetArgLength sets ArgLength field to given value.

### HasArgLength

`func (o *WafFirewallVersionDataAttributes) HasArgLength() bool`

HasArgLength returns a boolean if a field has been set.

### GetCombinedFileSizes

`func (o *WafFirewallVersionDataAttributes) GetCombinedFileSizes() int32`

GetCombinedFileSizes returns the CombinedFileSizes field if non-nil, zero value otherwise.

### GetCombinedFileSizesOk

`func (o *WafFirewallVersionDataAttributes) GetCombinedFileSizesOk() (*int32, bool)`

GetCombinedFileSizesOk returns a tuple with the CombinedFileSizes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCombinedFileSizes

`func (o *WafFirewallVersionDataAttributes) SetCombinedFileSizes(v int32)`

SetCombinedFileSizes sets CombinedFileSizes field to given value.

### HasCombinedFileSizes

`func (o *WafFirewallVersionDataAttributes) HasCombinedFileSizes() bool`

HasCombinedFileSizes returns a boolean if a field has been set.

### GetComment

`func (o *WafFirewallVersionDataAttributes) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *WafFirewallVersionDataAttributes) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *WafFirewallVersionDataAttributes) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *WafFirewallVersionDataAttributes) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *WafFirewallVersionDataAttributes) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *WafFirewallVersionDataAttributes) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetCriticalAnomalyScore

`func (o *WafFirewallVersionDataAttributes) GetCriticalAnomalyScore() int32`

GetCriticalAnomalyScore returns the CriticalAnomalyScore field if non-nil, zero value otherwise.

### GetCriticalAnomalyScoreOk

`func (o *WafFirewallVersionDataAttributes) GetCriticalAnomalyScoreOk() (*int32, bool)`

GetCriticalAnomalyScoreOk returns a tuple with the CriticalAnomalyScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticalAnomalyScore

`func (o *WafFirewallVersionDataAttributes) SetCriticalAnomalyScore(v int32)`

SetCriticalAnomalyScore sets CriticalAnomalyScore field to given value.

### HasCriticalAnomalyScore

`func (o *WafFirewallVersionDataAttributes) HasCriticalAnomalyScore() bool`

HasCriticalAnomalyScore returns a boolean if a field has been set.

### GetCrsValidateUtf8Encoding

`func (o *WafFirewallVersionDataAttributes) GetCrsValidateUtf8Encoding() bool`

GetCrsValidateUtf8Encoding returns the CrsValidateUtf8Encoding field if non-nil, zero value otherwise.

### GetCrsValidateUtf8EncodingOk

`func (o *WafFirewallVersionDataAttributes) GetCrsValidateUtf8EncodingOk() (*bool, bool)`

GetCrsValidateUtf8EncodingOk returns a tuple with the CrsValidateUtf8Encoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrsValidateUtf8Encoding

`func (o *WafFirewallVersionDataAttributes) SetCrsValidateUtf8Encoding(v bool)`

SetCrsValidateUtf8Encoding sets CrsValidateUtf8Encoding field to given value.

### HasCrsValidateUtf8Encoding

`func (o *WafFirewallVersionDataAttributes) HasCrsValidateUtf8Encoding() bool`

HasCrsValidateUtf8Encoding returns a boolean if a field has been set.

### GetErrorAnomalyScore

`func (o *WafFirewallVersionDataAttributes) GetErrorAnomalyScore() int32`

GetErrorAnomalyScore returns the ErrorAnomalyScore field if non-nil, zero value otherwise.

### GetErrorAnomalyScoreOk

`func (o *WafFirewallVersionDataAttributes) GetErrorAnomalyScoreOk() (*int32, bool)`

GetErrorAnomalyScoreOk returns a tuple with the ErrorAnomalyScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorAnomalyScore

`func (o *WafFirewallVersionDataAttributes) SetErrorAnomalyScore(v int32)`

SetErrorAnomalyScore sets ErrorAnomalyScore field to given value.

### HasErrorAnomalyScore

`func (o *WafFirewallVersionDataAttributes) HasErrorAnomalyScore() bool`

HasErrorAnomalyScore returns a boolean if a field has been set.

### GetHighRiskCountryCodes

`func (o *WafFirewallVersionDataAttributes) GetHighRiskCountryCodes() string`

GetHighRiskCountryCodes returns the HighRiskCountryCodes field if non-nil, zero value otherwise.

### GetHighRiskCountryCodesOk

`func (o *WafFirewallVersionDataAttributes) GetHighRiskCountryCodesOk() (*string, bool)`

GetHighRiskCountryCodesOk returns a tuple with the HighRiskCountryCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighRiskCountryCodes

`func (o *WafFirewallVersionDataAttributes) SetHighRiskCountryCodes(v string)`

SetHighRiskCountryCodes sets HighRiskCountryCodes field to given value.

### HasHighRiskCountryCodes

`func (o *WafFirewallVersionDataAttributes) HasHighRiskCountryCodes() bool`

HasHighRiskCountryCodes returns a boolean if a field has been set.

### GetHTTPViolationScoreThreshold

`func (o *WafFirewallVersionDataAttributes) GetHTTPViolationScoreThreshold() int32`

GetHTTPViolationScoreThreshold returns the HTTPViolationScoreThreshold field if non-nil, zero value otherwise.

### GetHTTPViolationScoreThresholdOk

`func (o *WafFirewallVersionDataAttributes) GetHTTPViolationScoreThresholdOk() (*int32, bool)`

GetHTTPViolationScoreThresholdOk returns a tuple with the HTTPViolationScoreThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHTTPViolationScoreThreshold

`func (o *WafFirewallVersionDataAttributes) SetHTTPViolationScoreThreshold(v int32)`

SetHTTPViolationScoreThreshold sets HTTPViolationScoreThreshold field to given value.

### HasHTTPViolationScoreThreshold

`func (o *WafFirewallVersionDataAttributes) HasHTTPViolationScoreThreshold() bool`

HasHTTPViolationScoreThreshold returns a boolean if a field has been set.

### GetInboundAnomalyScoreThreshold

`func (o *WafFirewallVersionDataAttributes) GetInboundAnomalyScoreThreshold() int32`

GetInboundAnomalyScoreThreshold returns the InboundAnomalyScoreThreshold field if non-nil, zero value otherwise.

### GetInboundAnomalyScoreThresholdOk

`func (o *WafFirewallVersionDataAttributes) GetInboundAnomalyScoreThresholdOk() (*int32, bool)`

GetInboundAnomalyScoreThresholdOk returns a tuple with the InboundAnomalyScoreThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundAnomalyScoreThreshold

`func (o *WafFirewallVersionDataAttributes) SetInboundAnomalyScoreThreshold(v int32)`

SetInboundAnomalyScoreThreshold sets InboundAnomalyScoreThreshold field to given value.

### HasInboundAnomalyScoreThreshold

`func (o *WafFirewallVersionDataAttributes) HasInboundAnomalyScoreThreshold() bool`

HasInboundAnomalyScoreThreshold returns a boolean if a field has been set.

### GetLfiScoreThreshold

`func (o *WafFirewallVersionDataAttributes) GetLfiScoreThreshold() int32`

GetLfiScoreThreshold returns the LfiScoreThreshold field if non-nil, zero value otherwise.

### GetLfiScoreThresholdOk

`func (o *WafFirewallVersionDataAttributes) GetLfiScoreThresholdOk() (*int32, bool)`

GetLfiScoreThresholdOk returns a tuple with the LfiScoreThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLfiScoreThreshold

`func (o *WafFirewallVersionDataAttributes) SetLfiScoreThreshold(v int32)`

SetLfiScoreThreshold sets LfiScoreThreshold field to given value.

### HasLfiScoreThreshold

`func (o *WafFirewallVersionDataAttributes) HasLfiScoreThreshold() bool`

HasLfiScoreThreshold returns a boolean if a field has been set.

### GetLocked

`func (o *WafFirewallVersionDataAttributes) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *WafFirewallVersionDataAttributes) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *WafFirewallVersionDataAttributes) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *WafFirewallVersionDataAttributes) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetMaxFileSize

`func (o *WafFirewallVersionDataAttributes) GetMaxFileSize() int32`

GetMaxFileSize returns the MaxFileSize field if non-nil, zero value otherwise.

### GetMaxFileSizeOk

`func (o *WafFirewallVersionDataAttributes) GetMaxFileSizeOk() (*int32, bool)`

GetMaxFileSizeOk returns a tuple with the MaxFileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFileSize

`func (o *WafFirewallVersionDataAttributes) SetMaxFileSize(v int32)`

SetMaxFileSize sets MaxFileSize field to given value.

### HasMaxFileSize

`func (o *WafFirewallVersionDataAttributes) HasMaxFileSize() bool`

HasMaxFileSize returns a boolean if a field has been set.

### GetMaxNumArgs

`func (o *WafFirewallVersionDataAttributes) GetMaxNumArgs() int32`

GetMaxNumArgs returns the MaxNumArgs field if non-nil, zero value otherwise.

### GetMaxNumArgsOk

`func (o *WafFirewallVersionDataAttributes) GetMaxNumArgsOk() (*int32, bool)`

GetMaxNumArgsOk returns a tuple with the MaxNumArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumArgs

`func (o *WafFirewallVersionDataAttributes) SetMaxNumArgs(v int32)`

SetMaxNumArgs sets MaxNumArgs field to given value.

### HasMaxNumArgs

`func (o *WafFirewallVersionDataAttributes) HasMaxNumArgs() bool`

HasMaxNumArgs returns a boolean if a field has been set.

### GetNoticeAnomalyScore

`func (o *WafFirewallVersionDataAttributes) GetNoticeAnomalyScore() int32`

GetNoticeAnomalyScore returns the NoticeAnomalyScore field if non-nil, zero value otherwise.

### GetNoticeAnomalyScoreOk

`func (o *WafFirewallVersionDataAttributes) GetNoticeAnomalyScoreOk() (*int32, bool)`

GetNoticeAnomalyScoreOk returns a tuple with the NoticeAnomalyScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoticeAnomalyScore

`func (o *WafFirewallVersionDataAttributes) SetNoticeAnomalyScore(v int32)`

SetNoticeAnomalyScore sets NoticeAnomalyScore field to given value.

### HasNoticeAnomalyScore

`func (o *WafFirewallVersionDataAttributes) HasNoticeAnomalyScore() bool`

HasNoticeAnomalyScore returns a boolean if a field has been set.

### GetNumber

`func (o *WafFirewallVersionDataAttributes) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *WafFirewallVersionDataAttributes) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *WafFirewallVersionDataAttributes) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *WafFirewallVersionDataAttributes) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetParanoiaLevel

`func (o *WafFirewallVersionDataAttributes) GetParanoiaLevel() int32`

GetParanoiaLevel returns the ParanoiaLevel field if non-nil, zero value otherwise.

### GetParanoiaLevelOk

`func (o *WafFirewallVersionDataAttributes) GetParanoiaLevelOk() (*int32, bool)`

GetParanoiaLevelOk returns a tuple with the ParanoiaLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParanoiaLevel

`func (o *WafFirewallVersionDataAttributes) SetParanoiaLevel(v int32)`

SetParanoiaLevel sets ParanoiaLevel field to given value.

### HasParanoiaLevel

`func (o *WafFirewallVersionDataAttributes) HasParanoiaLevel() bool`

HasParanoiaLevel returns a boolean if a field has been set.

### GetPhpInjectionScoreThreshold

`func (o *WafFirewallVersionDataAttributes) GetPhpInjectionScoreThreshold() int32`

GetPhpInjectionScoreThreshold returns the PhpInjectionScoreThreshold field if non-nil, zero value otherwise.

### GetPhpInjectionScoreThresholdOk

`func (o *WafFirewallVersionDataAttributes) GetPhpInjectionScoreThresholdOk() (*int32, bool)`

GetPhpInjectionScoreThresholdOk returns a tuple with the PhpInjectionScoreThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhpInjectionScoreThreshold

`func (o *WafFirewallVersionDataAttributes) SetPhpInjectionScoreThreshold(v int32)`

SetPhpInjectionScoreThreshold sets PhpInjectionScoreThreshold field to given value.

### HasPhpInjectionScoreThreshold

`func (o *WafFirewallVersionDataAttributes) HasPhpInjectionScoreThreshold() bool`

HasPhpInjectionScoreThreshold returns a boolean if a field has been set.

### GetRceScoreThreshold

`func (o *WafFirewallVersionDataAttributes) GetRceScoreThreshold() int32`

GetRceScoreThreshold returns the RceScoreThreshold field if non-nil, zero value otherwise.

### GetRceScoreThresholdOk

`func (o *WafFirewallVersionDataAttributes) GetRceScoreThresholdOk() (*int32, bool)`

GetRceScoreThresholdOk returns a tuple with the RceScoreThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRceScoreThreshold

`func (o *WafFirewallVersionDataAttributes) SetRceScoreThreshold(v int32)`

SetRceScoreThreshold sets RceScoreThreshold field to given value.

### HasRceScoreThreshold

`func (o *WafFirewallVersionDataAttributes) HasRceScoreThreshold() bool`

HasRceScoreThreshold returns a boolean if a field has been set.

### GetRestrictedExtensions

`func (o *WafFirewallVersionDataAttributes) GetRestrictedExtensions() string`

GetRestrictedExtensions returns the RestrictedExtensions field if non-nil, zero value otherwise.

### GetRestrictedExtensionsOk

`func (o *WafFirewallVersionDataAttributes) GetRestrictedExtensionsOk() (*string, bool)`

GetRestrictedExtensionsOk returns a tuple with the RestrictedExtensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedExtensions

`func (o *WafFirewallVersionDataAttributes) SetRestrictedExtensions(v string)`

SetRestrictedExtensions sets RestrictedExtensions field to given value.

### HasRestrictedExtensions

`func (o *WafFirewallVersionDataAttributes) HasRestrictedExtensions() bool`

HasRestrictedExtensions returns a boolean if a field has been set.

### GetRestrictedHeaders

`func (o *WafFirewallVersionDataAttributes) GetRestrictedHeaders() string`

GetRestrictedHeaders returns the RestrictedHeaders field if non-nil, zero value otherwise.

### GetRestrictedHeadersOk

`func (o *WafFirewallVersionDataAttributes) GetRestrictedHeadersOk() (*string, bool)`

GetRestrictedHeadersOk returns a tuple with the RestrictedHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedHeaders

`func (o *WafFirewallVersionDataAttributes) SetRestrictedHeaders(v string)`

SetRestrictedHeaders sets RestrictedHeaders field to given value.

### HasRestrictedHeaders

`func (o *WafFirewallVersionDataAttributes) HasRestrictedHeaders() bool`

HasRestrictedHeaders returns a boolean if a field has been set.

### GetRfiScoreThreshold

`func (o *WafFirewallVersionDataAttributes) GetRfiScoreThreshold() int32`

GetRfiScoreThreshold returns the RfiScoreThreshold field if non-nil, zero value otherwise.

### GetRfiScoreThresholdOk

`func (o *WafFirewallVersionDataAttributes) GetRfiScoreThresholdOk() (*int32, bool)`

GetRfiScoreThresholdOk returns a tuple with the RfiScoreThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRfiScoreThreshold

`func (o *WafFirewallVersionDataAttributes) SetRfiScoreThreshold(v int32)`

SetRfiScoreThreshold sets RfiScoreThreshold field to given value.

### HasRfiScoreThreshold

`func (o *WafFirewallVersionDataAttributes) HasRfiScoreThreshold() bool`

HasRfiScoreThreshold returns a boolean if a field has been set.

### GetSessionFixationScoreThreshold

`func (o *WafFirewallVersionDataAttributes) GetSessionFixationScoreThreshold() int32`

GetSessionFixationScoreThreshold returns the SessionFixationScoreThreshold field if non-nil, zero value otherwise.

### GetSessionFixationScoreThresholdOk

`func (o *WafFirewallVersionDataAttributes) GetSessionFixationScoreThresholdOk() (*int32, bool)`

GetSessionFixationScoreThresholdOk returns a tuple with the SessionFixationScoreThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionFixationScoreThreshold

`func (o *WafFirewallVersionDataAttributes) SetSessionFixationScoreThreshold(v int32)`

SetSessionFixationScoreThreshold sets SessionFixationScoreThreshold field to given value.

### HasSessionFixationScoreThreshold

`func (o *WafFirewallVersionDataAttributes) HasSessionFixationScoreThreshold() bool`

HasSessionFixationScoreThreshold returns a boolean if a field has been set.

### GetSQLInjectionScoreThreshold

`func (o *WafFirewallVersionDataAttributes) GetSQLInjectionScoreThreshold() int32`

GetSQLInjectionScoreThreshold returns the SQLInjectionScoreThreshold field if non-nil, zero value otherwise.

### GetSQLInjectionScoreThresholdOk

`func (o *WafFirewallVersionDataAttributes) GetSQLInjectionScoreThresholdOk() (*int32, bool)`

GetSQLInjectionScoreThresholdOk returns a tuple with the SQLInjectionScoreThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSQLInjectionScoreThreshold

`func (o *WafFirewallVersionDataAttributes) SetSQLInjectionScoreThreshold(v int32)`

SetSQLInjectionScoreThreshold sets SQLInjectionScoreThreshold field to given value.

### HasSQLInjectionScoreThreshold

`func (o *WafFirewallVersionDataAttributes) HasSQLInjectionScoreThreshold() bool`

HasSQLInjectionScoreThreshold returns a boolean if a field has been set.

### GetTotalArgLength

`func (o *WafFirewallVersionDataAttributes) GetTotalArgLength() int32`

GetTotalArgLength returns the TotalArgLength field if non-nil, zero value otherwise.

### GetTotalArgLengthOk

`func (o *WafFirewallVersionDataAttributes) GetTotalArgLengthOk() (*int32, bool)`

GetTotalArgLengthOk returns a tuple with the TotalArgLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalArgLength

`func (o *WafFirewallVersionDataAttributes) SetTotalArgLength(v int32)`

SetTotalArgLength sets TotalArgLength field to given value.

### HasTotalArgLength

`func (o *WafFirewallVersionDataAttributes) HasTotalArgLength() bool`

HasTotalArgLength returns a boolean if a field has been set.

### GetWarningAnomalyScore

`func (o *WafFirewallVersionDataAttributes) GetWarningAnomalyScore() int32`

GetWarningAnomalyScore returns the WarningAnomalyScore field if non-nil, zero value otherwise.

### GetWarningAnomalyScoreOk

`func (o *WafFirewallVersionDataAttributes) GetWarningAnomalyScoreOk() (*int32, bool)`

GetWarningAnomalyScoreOk returns a tuple with the WarningAnomalyScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningAnomalyScore

`func (o *WafFirewallVersionDataAttributes) SetWarningAnomalyScore(v int32)`

SetWarningAnomalyScore sets WarningAnomalyScore field to given value.

### HasWarningAnomalyScore

`func (o *WafFirewallVersionDataAttributes) HasWarningAnomalyScore() bool`

HasWarningAnomalyScore returns a boolean if a field has been set.

### GetXSSScoreThreshold

`func (o *WafFirewallVersionDataAttributes) GetXSSScoreThreshold() int32`

GetXSSScoreThreshold returns the XSSScoreThreshold field if non-nil, zero value otherwise.

### GetXSSScoreThresholdOk

`func (o *WafFirewallVersionDataAttributes) GetXSSScoreThresholdOk() (*int32, bool)`

GetXSSScoreThresholdOk returns a tuple with the XSSScoreThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXSSScoreThreshold

`func (o *WafFirewallVersionDataAttributes) SetXSSScoreThreshold(v int32)`

SetXSSScoreThreshold sets XSSScoreThreshold field to given value.

### HasXSSScoreThreshold

`func (o *WafFirewallVersionDataAttributes) HasXSSScoreThreshold() bool`

HasXSSScoreThreshold returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
