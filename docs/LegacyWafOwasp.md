# LegacyWafOwasp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedHTTPVersions** | Pointer to **string** | Allowed HTTP versions. | [optional] [default to "HTTP/1.0 HTTP/1.1 HTTP/2"]
**AllowedMethods** | Pointer to **string** | A space-separated list of HTTP method names. | [optional] [default to "GET HEAD POST OPTIONS PUT PATCH DELETE"]
**AllowedRequestContentType** | Pointer to **string** | Allowed request content types. | [optional] [default to "application/x-www-form-urlencoded|multipart/form-data|text/xml|application/xml|application/x-amf|application/json|text/plain"]
**ArgLength** | Pointer to **int32** | The maximum allowed length of an argument. | [optional] [default to 400]
**ArgNameLength** | Pointer to **int32** | The maximum allowed argument name length. | [optional] [default to 100]
**CombinedFileSizes** | Pointer to **int32** | The maximum allowed size of all files (in bytes). | [optional] [default to 10000000]
**CreatedAt** | Pointer to **string** | Date and time that the settings object was created. | [optional] 
**CriticalAnomalyScore** | Pointer to **int32** | Score value to add for critical anomalies. | [optional] [default to 6]
**CrsValidateUtf8Encoding** | Pointer to **bool** | CRS validate UTF8 encoding. | [optional] 
**ErrorAnomalyScore** | Pointer to **int32** | Score value to add for error anomalies. | [optional] [default to 5]
**HighRiskCountryCodes** | Pointer to **string** | A space-separated list of country codes in ISO 3166-1 (two-letter) format. | [optional] 
**HTTPViolationScoreThreshold** | Pointer to **int32** | HTTP violation threshold. | [optional] 
**InboundAnomalyScoreThreshold** | Pointer to **int32** | Inbound anomaly threshold. | [optional] 
**LfiScoreThreshold** | Pointer to **int32** | Local file inclusion attack threshold. | [optional] 
**MaxFileSize** | Pointer to **int32** | The maximum allowed file size (in bytes). | [optional] [default to 10000000]
**MaxNumArgs** | Pointer to **int32** | The maximum number of arguments allowed. | [optional] [default to 255]
**NoticeAnomalyScore** | Pointer to **int32** | Score value to add for notice anomalies. | [optional] [default to 4]
**ParanoiaLevel** | Pointer to **int32** | The configured paranoia level. | [optional] [default to 1]
**PhpInjectionScoreThreshold** | Pointer to **int32** | PHP injection threshold. | [optional] 
**RceScoreThreshold** | Pointer to **int32** | Remote code execution threshold. | [optional] 
**RestrictedExtensions** | Pointer to **string** | A space-separated list of disallowed file extensions. | [optional] [default to ".asa/ .asax/ .ascx/ .axd/ .backup/ .bak/ .bat/ .cdx/ .cer/ .cfg/ .cmd/ .com/ .config/ .conf/ .cs/ .csproj/ .csr/ .dat/ .db/ .dbf/ .dll/ .dos/ .htr/ .htw/ .ida/ .idc/ .idq/ .inc/ .ini/ .key/ .licx/ .lnk/ .log/ .mdb/ .old/ .pass/ .pdb/ .pol/ .printer/ .pwd/ .resources/ .resx/ .sql/ .sys/ .vb/ .vbs/ .vbproj/ .vsdisco/ .webinfo/ .xsd/ .xsx"]
**RestrictedHeaders** | Pointer to **string** | A space-separated list of disallowed header names. | [optional] [default to "/proxy/ /lock-token/ /content-range/ /translate/ /if/"]
**RfiScoreThreshold** | Pointer to **int32** | Remote file inclusion attack threshold. | [optional] 
**SessionFixationScoreThreshold** | Pointer to **int32** | Session fixation attack threshold. | [optional] 
**SQLInjectionScoreThreshold** | Pointer to **int32** | SQL injection attack threshold. | [optional] 
**TotalArgLength** | Pointer to **int32** | The maximum size of argument names and values. | [optional] [default to 6400]
**UpdatedAt** | Pointer to **string** | Date and time that the settings object was last updated. | [optional] 
**WarningAnomalyScore** | Pointer to **int32** | Score value to add for warning anomalies. | [optional] 
**XSSScoreThreshold** | Pointer to **int32** | XSS attack threshold. | [optional] 

## Methods

### NewLegacyWafOwasp

`func NewLegacyWafOwasp() *LegacyWafOwasp`

NewLegacyWafOwasp instantiates a new LegacyWafOwasp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLegacyWafOwaspWithDefaults

`func NewLegacyWafOwaspWithDefaults() *LegacyWafOwasp`

NewLegacyWafOwaspWithDefaults instantiates a new LegacyWafOwasp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedHTTPVersions

`func (o *LegacyWafOwasp) GetAllowedHTTPVersions() string`

GetAllowedHTTPVersions returns the AllowedHTTPVersions field if non-nil, zero value otherwise.

### GetAllowedHTTPVersionsOk

`func (o *LegacyWafOwasp) GetAllowedHTTPVersionsOk() (*string, bool)`

GetAllowedHTTPVersionsOk returns a tuple with the AllowedHTTPVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedHTTPVersions

`func (o *LegacyWafOwasp) SetAllowedHTTPVersions(v string)`

SetAllowedHTTPVersions sets AllowedHTTPVersions field to given value.

### HasAllowedHTTPVersions

`func (o *LegacyWafOwasp) HasAllowedHTTPVersions() bool`

HasAllowedHTTPVersions returns a boolean if a field has been set.

### GetAllowedMethods

`func (o *LegacyWafOwasp) GetAllowedMethods() string`

GetAllowedMethods returns the AllowedMethods field if non-nil, zero value otherwise.

### GetAllowedMethodsOk

`func (o *LegacyWafOwasp) GetAllowedMethodsOk() (*string, bool)`

GetAllowedMethodsOk returns a tuple with the AllowedMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedMethods

`func (o *LegacyWafOwasp) SetAllowedMethods(v string)`

SetAllowedMethods sets AllowedMethods field to given value.

### HasAllowedMethods

`func (o *LegacyWafOwasp) HasAllowedMethods() bool`

HasAllowedMethods returns a boolean if a field has been set.

### GetAllowedRequestContentType

`func (o *LegacyWafOwasp) GetAllowedRequestContentType() string`

GetAllowedRequestContentType returns the AllowedRequestContentType field if non-nil, zero value otherwise.

### GetAllowedRequestContentTypeOk

`func (o *LegacyWafOwasp) GetAllowedRequestContentTypeOk() (*string, bool)`

GetAllowedRequestContentTypeOk returns a tuple with the AllowedRequestContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedRequestContentType

`func (o *LegacyWafOwasp) SetAllowedRequestContentType(v string)`

SetAllowedRequestContentType sets AllowedRequestContentType field to given value.

### HasAllowedRequestContentType

`func (o *LegacyWafOwasp) HasAllowedRequestContentType() bool`

HasAllowedRequestContentType returns a boolean if a field has been set.

### GetArgLength

`func (o *LegacyWafOwasp) GetArgLength() int32`

GetArgLength returns the ArgLength field if non-nil, zero value otherwise.

### GetArgLengthOk

`func (o *LegacyWafOwasp) GetArgLengthOk() (*int32, bool)`

GetArgLengthOk returns a tuple with the ArgLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgLength

`func (o *LegacyWafOwasp) SetArgLength(v int32)`

SetArgLength sets ArgLength field to given value.

### HasArgLength

`func (o *LegacyWafOwasp) HasArgLength() bool`

HasArgLength returns a boolean if a field has been set.

### GetArgNameLength

`func (o *LegacyWafOwasp) GetArgNameLength() int32`

GetArgNameLength returns the ArgNameLength field if non-nil, zero value otherwise.

### GetArgNameLengthOk

`func (o *LegacyWafOwasp) GetArgNameLengthOk() (*int32, bool)`

GetArgNameLengthOk returns a tuple with the ArgNameLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgNameLength

`func (o *LegacyWafOwasp) SetArgNameLength(v int32)`

SetArgNameLength sets ArgNameLength field to given value.

### HasArgNameLength

`func (o *LegacyWafOwasp) HasArgNameLength() bool`

HasArgNameLength returns a boolean if a field has been set.

### GetCombinedFileSizes

`func (o *LegacyWafOwasp) GetCombinedFileSizes() int32`

GetCombinedFileSizes returns the CombinedFileSizes field if non-nil, zero value otherwise.

### GetCombinedFileSizesOk

`func (o *LegacyWafOwasp) GetCombinedFileSizesOk() (*int32, bool)`

GetCombinedFileSizesOk returns a tuple with the CombinedFileSizes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCombinedFileSizes

`func (o *LegacyWafOwasp) SetCombinedFileSizes(v int32)`

SetCombinedFileSizes sets CombinedFileSizes field to given value.

### HasCombinedFileSizes

`func (o *LegacyWafOwasp) HasCombinedFileSizes() bool`

HasCombinedFileSizes returns a boolean if a field has been set.

### GetCreatedAt

`func (o *LegacyWafOwasp) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LegacyWafOwasp) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LegacyWafOwasp) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *LegacyWafOwasp) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCriticalAnomalyScore

`func (o *LegacyWafOwasp) GetCriticalAnomalyScore() int32`

GetCriticalAnomalyScore returns the CriticalAnomalyScore field if non-nil, zero value otherwise.

### GetCriticalAnomalyScoreOk

`func (o *LegacyWafOwasp) GetCriticalAnomalyScoreOk() (*int32, bool)`

GetCriticalAnomalyScoreOk returns a tuple with the CriticalAnomalyScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticalAnomalyScore

`func (o *LegacyWafOwasp) SetCriticalAnomalyScore(v int32)`

SetCriticalAnomalyScore sets CriticalAnomalyScore field to given value.

### HasCriticalAnomalyScore

`func (o *LegacyWafOwasp) HasCriticalAnomalyScore() bool`

HasCriticalAnomalyScore returns a boolean if a field has been set.

### GetCrsValidateUtf8Encoding

`func (o *LegacyWafOwasp) GetCrsValidateUtf8Encoding() bool`

GetCrsValidateUtf8Encoding returns the CrsValidateUtf8Encoding field if non-nil, zero value otherwise.

### GetCrsValidateUtf8EncodingOk

`func (o *LegacyWafOwasp) GetCrsValidateUtf8EncodingOk() (*bool, bool)`

GetCrsValidateUtf8EncodingOk returns a tuple with the CrsValidateUtf8Encoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrsValidateUtf8Encoding

`func (o *LegacyWafOwasp) SetCrsValidateUtf8Encoding(v bool)`

SetCrsValidateUtf8Encoding sets CrsValidateUtf8Encoding field to given value.

### HasCrsValidateUtf8Encoding

`func (o *LegacyWafOwasp) HasCrsValidateUtf8Encoding() bool`

HasCrsValidateUtf8Encoding returns a boolean if a field has been set.

### GetErrorAnomalyScore

`func (o *LegacyWafOwasp) GetErrorAnomalyScore() int32`

GetErrorAnomalyScore returns the ErrorAnomalyScore field if non-nil, zero value otherwise.

### GetErrorAnomalyScoreOk

`func (o *LegacyWafOwasp) GetErrorAnomalyScoreOk() (*int32, bool)`

GetErrorAnomalyScoreOk returns a tuple with the ErrorAnomalyScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorAnomalyScore

`func (o *LegacyWafOwasp) SetErrorAnomalyScore(v int32)`

SetErrorAnomalyScore sets ErrorAnomalyScore field to given value.

### HasErrorAnomalyScore

`func (o *LegacyWafOwasp) HasErrorAnomalyScore() bool`

HasErrorAnomalyScore returns a boolean if a field has been set.

### GetHighRiskCountryCodes

`func (o *LegacyWafOwasp) GetHighRiskCountryCodes() string`

GetHighRiskCountryCodes returns the HighRiskCountryCodes field if non-nil, zero value otherwise.

### GetHighRiskCountryCodesOk

`func (o *LegacyWafOwasp) GetHighRiskCountryCodesOk() (*string, bool)`

GetHighRiskCountryCodesOk returns a tuple with the HighRiskCountryCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighRiskCountryCodes

`func (o *LegacyWafOwasp) SetHighRiskCountryCodes(v string)`

SetHighRiskCountryCodes sets HighRiskCountryCodes field to given value.

### HasHighRiskCountryCodes

`func (o *LegacyWafOwasp) HasHighRiskCountryCodes() bool`

HasHighRiskCountryCodes returns a boolean if a field has been set.

### GetHTTPViolationScoreThreshold

`func (o *LegacyWafOwasp) GetHTTPViolationScoreThreshold() int32`

GetHTTPViolationScoreThreshold returns the HTTPViolationScoreThreshold field if non-nil, zero value otherwise.

### GetHTTPViolationScoreThresholdOk

`func (o *LegacyWafOwasp) GetHTTPViolationScoreThresholdOk() (*int32, bool)`

GetHTTPViolationScoreThresholdOk returns a tuple with the HTTPViolationScoreThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHTTPViolationScoreThreshold

`func (o *LegacyWafOwasp) SetHTTPViolationScoreThreshold(v int32)`

SetHTTPViolationScoreThreshold sets HTTPViolationScoreThreshold field to given value.

### HasHTTPViolationScoreThreshold

`func (o *LegacyWafOwasp) HasHTTPViolationScoreThreshold() bool`

HasHTTPViolationScoreThreshold returns a boolean if a field has been set.

### GetInboundAnomalyScoreThreshold

`func (o *LegacyWafOwasp) GetInboundAnomalyScoreThreshold() int32`

GetInboundAnomalyScoreThreshold returns the InboundAnomalyScoreThreshold field if non-nil, zero value otherwise.

### GetInboundAnomalyScoreThresholdOk

`func (o *LegacyWafOwasp) GetInboundAnomalyScoreThresholdOk() (*int32, bool)`

GetInboundAnomalyScoreThresholdOk returns a tuple with the InboundAnomalyScoreThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundAnomalyScoreThreshold

`func (o *LegacyWafOwasp) SetInboundAnomalyScoreThreshold(v int32)`

SetInboundAnomalyScoreThreshold sets InboundAnomalyScoreThreshold field to given value.

### HasInboundAnomalyScoreThreshold

`func (o *LegacyWafOwasp) HasInboundAnomalyScoreThreshold() bool`

HasInboundAnomalyScoreThreshold returns a boolean if a field has been set.

### GetLfiScoreThreshold

`func (o *LegacyWafOwasp) GetLfiScoreThreshold() int32`

GetLfiScoreThreshold returns the LfiScoreThreshold field if non-nil, zero value otherwise.

### GetLfiScoreThresholdOk

`func (o *LegacyWafOwasp) GetLfiScoreThresholdOk() (*int32, bool)`

GetLfiScoreThresholdOk returns a tuple with the LfiScoreThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLfiScoreThreshold

`func (o *LegacyWafOwasp) SetLfiScoreThreshold(v int32)`

SetLfiScoreThreshold sets LfiScoreThreshold field to given value.

### HasLfiScoreThreshold

`func (o *LegacyWafOwasp) HasLfiScoreThreshold() bool`

HasLfiScoreThreshold returns a boolean if a field has been set.

### GetMaxFileSize

`func (o *LegacyWafOwasp) GetMaxFileSize() int32`

GetMaxFileSize returns the MaxFileSize field if non-nil, zero value otherwise.

### GetMaxFileSizeOk

`func (o *LegacyWafOwasp) GetMaxFileSizeOk() (*int32, bool)`

GetMaxFileSizeOk returns a tuple with the MaxFileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFileSize

`func (o *LegacyWafOwasp) SetMaxFileSize(v int32)`

SetMaxFileSize sets MaxFileSize field to given value.

### HasMaxFileSize

`func (o *LegacyWafOwasp) HasMaxFileSize() bool`

HasMaxFileSize returns a boolean if a field has been set.

### GetMaxNumArgs

`func (o *LegacyWafOwasp) GetMaxNumArgs() int32`

GetMaxNumArgs returns the MaxNumArgs field if non-nil, zero value otherwise.

### GetMaxNumArgsOk

`func (o *LegacyWafOwasp) GetMaxNumArgsOk() (*int32, bool)`

GetMaxNumArgsOk returns a tuple with the MaxNumArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumArgs

`func (o *LegacyWafOwasp) SetMaxNumArgs(v int32)`

SetMaxNumArgs sets MaxNumArgs field to given value.

### HasMaxNumArgs

`func (o *LegacyWafOwasp) HasMaxNumArgs() bool`

HasMaxNumArgs returns a boolean if a field has been set.

### GetNoticeAnomalyScore

`func (o *LegacyWafOwasp) GetNoticeAnomalyScore() int32`

GetNoticeAnomalyScore returns the NoticeAnomalyScore field if non-nil, zero value otherwise.

### GetNoticeAnomalyScoreOk

`func (o *LegacyWafOwasp) GetNoticeAnomalyScoreOk() (*int32, bool)`

GetNoticeAnomalyScoreOk returns a tuple with the NoticeAnomalyScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoticeAnomalyScore

`func (o *LegacyWafOwasp) SetNoticeAnomalyScore(v int32)`

SetNoticeAnomalyScore sets NoticeAnomalyScore field to given value.

### HasNoticeAnomalyScore

`func (o *LegacyWafOwasp) HasNoticeAnomalyScore() bool`

HasNoticeAnomalyScore returns a boolean if a field has been set.

### GetParanoiaLevel

`func (o *LegacyWafOwasp) GetParanoiaLevel() int32`

GetParanoiaLevel returns the ParanoiaLevel field if non-nil, zero value otherwise.

### GetParanoiaLevelOk

`func (o *LegacyWafOwasp) GetParanoiaLevelOk() (*int32, bool)`

GetParanoiaLevelOk returns a tuple with the ParanoiaLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParanoiaLevel

`func (o *LegacyWafOwasp) SetParanoiaLevel(v int32)`

SetParanoiaLevel sets ParanoiaLevel field to given value.

### HasParanoiaLevel

`func (o *LegacyWafOwasp) HasParanoiaLevel() bool`

HasParanoiaLevel returns a boolean if a field has been set.

### GetPhpInjectionScoreThreshold

`func (o *LegacyWafOwasp) GetPhpInjectionScoreThreshold() int32`

GetPhpInjectionScoreThreshold returns the PhpInjectionScoreThreshold field if non-nil, zero value otherwise.

### GetPhpInjectionScoreThresholdOk

`func (o *LegacyWafOwasp) GetPhpInjectionScoreThresholdOk() (*int32, bool)`

GetPhpInjectionScoreThresholdOk returns a tuple with the PhpInjectionScoreThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhpInjectionScoreThreshold

`func (o *LegacyWafOwasp) SetPhpInjectionScoreThreshold(v int32)`

SetPhpInjectionScoreThreshold sets PhpInjectionScoreThreshold field to given value.

### HasPhpInjectionScoreThreshold

`func (o *LegacyWafOwasp) HasPhpInjectionScoreThreshold() bool`

HasPhpInjectionScoreThreshold returns a boolean if a field has been set.

### GetRceScoreThreshold

`func (o *LegacyWafOwasp) GetRceScoreThreshold() int32`

GetRceScoreThreshold returns the RceScoreThreshold field if non-nil, zero value otherwise.

### GetRceScoreThresholdOk

`func (o *LegacyWafOwasp) GetRceScoreThresholdOk() (*int32, bool)`

GetRceScoreThresholdOk returns a tuple with the RceScoreThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRceScoreThreshold

`func (o *LegacyWafOwasp) SetRceScoreThreshold(v int32)`

SetRceScoreThreshold sets RceScoreThreshold field to given value.

### HasRceScoreThreshold

`func (o *LegacyWafOwasp) HasRceScoreThreshold() bool`

HasRceScoreThreshold returns a boolean if a field has been set.

### GetRestrictedExtensions

`func (o *LegacyWafOwasp) GetRestrictedExtensions() string`

GetRestrictedExtensions returns the RestrictedExtensions field if non-nil, zero value otherwise.

### GetRestrictedExtensionsOk

`func (o *LegacyWafOwasp) GetRestrictedExtensionsOk() (*string, bool)`

GetRestrictedExtensionsOk returns a tuple with the RestrictedExtensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedExtensions

`func (o *LegacyWafOwasp) SetRestrictedExtensions(v string)`

SetRestrictedExtensions sets RestrictedExtensions field to given value.

### HasRestrictedExtensions

`func (o *LegacyWafOwasp) HasRestrictedExtensions() bool`

HasRestrictedExtensions returns a boolean if a field has been set.

### GetRestrictedHeaders

`func (o *LegacyWafOwasp) GetRestrictedHeaders() string`

GetRestrictedHeaders returns the RestrictedHeaders field if non-nil, zero value otherwise.

### GetRestrictedHeadersOk

`func (o *LegacyWafOwasp) GetRestrictedHeadersOk() (*string, bool)`

GetRestrictedHeadersOk returns a tuple with the RestrictedHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedHeaders

`func (o *LegacyWafOwasp) SetRestrictedHeaders(v string)`

SetRestrictedHeaders sets RestrictedHeaders field to given value.

### HasRestrictedHeaders

`func (o *LegacyWafOwasp) HasRestrictedHeaders() bool`

HasRestrictedHeaders returns a boolean if a field has been set.

### GetRfiScoreThreshold

`func (o *LegacyWafOwasp) GetRfiScoreThreshold() int32`

GetRfiScoreThreshold returns the RfiScoreThreshold field if non-nil, zero value otherwise.

### GetRfiScoreThresholdOk

`func (o *LegacyWafOwasp) GetRfiScoreThresholdOk() (*int32, bool)`

GetRfiScoreThresholdOk returns a tuple with the RfiScoreThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRfiScoreThreshold

`func (o *LegacyWafOwasp) SetRfiScoreThreshold(v int32)`

SetRfiScoreThreshold sets RfiScoreThreshold field to given value.

### HasRfiScoreThreshold

`func (o *LegacyWafOwasp) HasRfiScoreThreshold() bool`

HasRfiScoreThreshold returns a boolean if a field has been set.

### GetSessionFixationScoreThreshold

`func (o *LegacyWafOwasp) GetSessionFixationScoreThreshold() int32`

GetSessionFixationScoreThreshold returns the SessionFixationScoreThreshold field if non-nil, zero value otherwise.

### GetSessionFixationScoreThresholdOk

`func (o *LegacyWafOwasp) GetSessionFixationScoreThresholdOk() (*int32, bool)`

GetSessionFixationScoreThresholdOk returns a tuple with the SessionFixationScoreThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionFixationScoreThreshold

`func (o *LegacyWafOwasp) SetSessionFixationScoreThreshold(v int32)`

SetSessionFixationScoreThreshold sets SessionFixationScoreThreshold field to given value.

### HasSessionFixationScoreThreshold

`func (o *LegacyWafOwasp) HasSessionFixationScoreThreshold() bool`

HasSessionFixationScoreThreshold returns a boolean if a field has been set.

### GetSQLInjectionScoreThreshold

`func (o *LegacyWafOwasp) GetSQLInjectionScoreThreshold() int32`

GetSQLInjectionScoreThreshold returns the SQLInjectionScoreThreshold field if non-nil, zero value otherwise.

### GetSQLInjectionScoreThresholdOk

`func (o *LegacyWafOwasp) GetSQLInjectionScoreThresholdOk() (*int32, bool)`

GetSQLInjectionScoreThresholdOk returns a tuple with the SQLInjectionScoreThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSQLInjectionScoreThreshold

`func (o *LegacyWafOwasp) SetSQLInjectionScoreThreshold(v int32)`

SetSQLInjectionScoreThreshold sets SQLInjectionScoreThreshold field to given value.

### HasSQLInjectionScoreThreshold

`func (o *LegacyWafOwasp) HasSQLInjectionScoreThreshold() bool`

HasSQLInjectionScoreThreshold returns a boolean if a field has been set.

### GetTotalArgLength

`func (o *LegacyWafOwasp) GetTotalArgLength() int32`

GetTotalArgLength returns the TotalArgLength field if non-nil, zero value otherwise.

### GetTotalArgLengthOk

`func (o *LegacyWafOwasp) GetTotalArgLengthOk() (*int32, bool)`

GetTotalArgLengthOk returns a tuple with the TotalArgLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalArgLength

`func (o *LegacyWafOwasp) SetTotalArgLength(v int32)`

SetTotalArgLength sets TotalArgLength field to given value.

### HasTotalArgLength

`func (o *LegacyWafOwasp) HasTotalArgLength() bool`

HasTotalArgLength returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *LegacyWafOwasp) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *LegacyWafOwasp) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *LegacyWafOwasp) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *LegacyWafOwasp) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetWarningAnomalyScore

`func (o *LegacyWafOwasp) GetWarningAnomalyScore() int32`

GetWarningAnomalyScore returns the WarningAnomalyScore field if non-nil, zero value otherwise.

### GetWarningAnomalyScoreOk

`func (o *LegacyWafOwasp) GetWarningAnomalyScoreOk() (*int32, bool)`

GetWarningAnomalyScoreOk returns a tuple with the WarningAnomalyScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningAnomalyScore

`func (o *LegacyWafOwasp) SetWarningAnomalyScore(v int32)`

SetWarningAnomalyScore sets WarningAnomalyScore field to given value.

### HasWarningAnomalyScore

`func (o *LegacyWafOwasp) HasWarningAnomalyScore() bool`

HasWarningAnomalyScore returns a boolean if a field has been set.

### GetXSSScoreThreshold

`func (o *LegacyWafOwasp) GetXSSScoreThreshold() int32`

GetXSSScoreThreshold returns the XSSScoreThreshold field if non-nil, zero value otherwise.

### GetXSSScoreThresholdOk

`func (o *LegacyWafOwasp) GetXSSScoreThresholdOk() (*int32, bool)`

GetXSSScoreThresholdOk returns a tuple with the XSSScoreThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXSSScoreThreshold

`func (o *LegacyWafOwasp) SetXSSScoreThreshold(v int32)`

SetXSSScoreThreshold sets XSSScoreThreshold field to given value.

### HasXSSScoreThreshold

`func (o *LegacyWafOwasp) HasXSSScoreThreshold() bool`

HasXSSScoreThreshold returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
