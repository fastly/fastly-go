# RateLimiter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A human readable name for the rate limiting rule. | [optional] 
**URIDictionaryName** | Pointer to **NullableString** | The name of an Edge Dictionary containing URIs as keys. If not defined or `null`, all origin URIs will be rate limited. | [optional] 
**HTTPMethods** | Pointer to **[]string** | Array of HTTP methods to apply rate limiting to. | [optional] 
**RpsLimit** | Pointer to **int32** | Upper limit of requests per second allowed by the rate limiter. | [optional] 
**WindowSize** | Pointer to **int32** | Number of seconds during which the RPS limit must be exceeded in order to trigger a violation. | [optional] 
**ClientKey** | Pointer to **[]string** | Array of VCL variables used to generate a counter key to identify a client. Example variables include `req.http.Fastly-Client-IP`, `req.http.User-Agent`, or a custom header like `req.http.API-Key`. | [optional] 
**PenaltyBoxDuration** | Pointer to **int32** | Length of time in minutes that the rate limiter is in effect after the initial violation is detected. | [optional] 
**Action** | Pointer to **string** | The action to take when a rate limiter violation is detected. | [optional] 
**Response** | Pointer to [**NullableRateLimiterResponse1**](RateLimiterResponse1.md) |  | [optional] 
**ResponseObjectName** | Pointer to **NullableString** | Name of existing response object. Required if `action` is `response_object`. Note that the rate limiter response is only updated to reflect the response object content when saving the rate limiter configuration. | [optional] 
**LoggerType** | Pointer to **string** | Name of the type of logging endpoint to be used when action is `log_only`. The logging endpoint type is used to determine the appropriate log format to use when emitting log entries. | [optional] 
**FeatureRevision** | Pointer to **int32** | Revision number of the rate limiting feature implementation. Defaults to the most recent revision. | [optional] 

## Methods

### NewRateLimiter

`func NewRateLimiter() *RateLimiter`

NewRateLimiter instantiates a new RateLimiter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRateLimiterWithDefaults

`func NewRateLimiterWithDefaults() *RateLimiter`

NewRateLimiterWithDefaults instantiates a new RateLimiter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RateLimiter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RateLimiter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RateLimiter) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RateLimiter) HasName() bool`

HasName returns a boolean if a field has been set.

### GetURIDictionaryName

`func (o *RateLimiter) GetURIDictionaryName() string`

GetURIDictionaryName returns the URIDictionaryName field if non-nil, zero value otherwise.

### GetURIDictionaryNameOk

`func (o *RateLimiter) GetURIDictionaryNameOk() (*string, bool)`

GetURIDictionaryNameOk returns a tuple with the URIDictionaryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetURIDictionaryName

`func (o *RateLimiter) SetURIDictionaryName(v string)`

SetURIDictionaryName sets URIDictionaryName field to given value.

### HasURIDictionaryName

`func (o *RateLimiter) HasURIDictionaryName() bool`

HasURIDictionaryName returns a boolean if a field has been set.

### SetURIDictionaryNameNil

`func (o *RateLimiter) SetURIDictionaryNameNil(b bool)`

 SetURIDictionaryNameNil sets the value for URIDictionaryName to be an explicit nil

### UnsetURIDictionaryName
`func (o *RateLimiter) UnsetURIDictionaryName()`

UnsetURIDictionaryName ensures that no value is present for URIDictionaryName, not even an explicit nil
### GetHTTPMethods

`func (o *RateLimiter) GetHTTPMethods() []string`

GetHTTPMethods returns the HTTPMethods field if non-nil, zero value otherwise.

### GetHTTPMethodsOk

`func (o *RateLimiter) GetHTTPMethodsOk() (*[]string, bool)`

GetHTTPMethodsOk returns a tuple with the HTTPMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHTTPMethods

`func (o *RateLimiter) SetHTTPMethods(v []string)`

SetHTTPMethods sets HTTPMethods field to given value.

### HasHTTPMethods

`func (o *RateLimiter) HasHTTPMethods() bool`

HasHTTPMethods returns a boolean if a field has been set.

### GetRpsLimit

`func (o *RateLimiter) GetRpsLimit() int32`

GetRpsLimit returns the RpsLimit field if non-nil, zero value otherwise.

### GetRpsLimitOk

`func (o *RateLimiter) GetRpsLimitOk() (*int32, bool)`

GetRpsLimitOk returns a tuple with the RpsLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpsLimit

`func (o *RateLimiter) SetRpsLimit(v int32)`

SetRpsLimit sets RpsLimit field to given value.

### HasRpsLimit

`func (o *RateLimiter) HasRpsLimit() bool`

HasRpsLimit returns a boolean if a field has been set.

### GetWindowSize

`func (o *RateLimiter) GetWindowSize() int32`

GetWindowSize returns the WindowSize field if non-nil, zero value otherwise.

### GetWindowSizeOk

`func (o *RateLimiter) GetWindowSizeOk() (*int32, bool)`

GetWindowSizeOk returns a tuple with the WindowSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowSize

`func (o *RateLimiter) SetWindowSize(v int32)`

SetWindowSize sets WindowSize field to given value.

### HasWindowSize

`func (o *RateLimiter) HasWindowSize() bool`

HasWindowSize returns a boolean if a field has been set.

### GetClientKey

`func (o *RateLimiter) GetClientKey() []string`

GetClientKey returns the ClientKey field if non-nil, zero value otherwise.

### GetClientKeyOk

`func (o *RateLimiter) GetClientKeyOk() (*[]string, bool)`

GetClientKeyOk returns a tuple with the ClientKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientKey

`func (o *RateLimiter) SetClientKey(v []string)`

SetClientKey sets ClientKey field to given value.

### HasClientKey

`func (o *RateLimiter) HasClientKey() bool`

HasClientKey returns a boolean if a field has been set.

### GetPenaltyBoxDuration

`func (o *RateLimiter) GetPenaltyBoxDuration() int32`

GetPenaltyBoxDuration returns the PenaltyBoxDuration field if non-nil, zero value otherwise.

### GetPenaltyBoxDurationOk

`func (o *RateLimiter) GetPenaltyBoxDurationOk() (*int32, bool)`

GetPenaltyBoxDurationOk returns a tuple with the PenaltyBoxDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPenaltyBoxDuration

`func (o *RateLimiter) SetPenaltyBoxDuration(v int32)`

SetPenaltyBoxDuration sets PenaltyBoxDuration field to given value.

### HasPenaltyBoxDuration

`func (o *RateLimiter) HasPenaltyBoxDuration() bool`

HasPenaltyBoxDuration returns a boolean if a field has been set.

### GetAction

`func (o *RateLimiter) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *RateLimiter) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *RateLimiter) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *RateLimiter) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetResponse

`func (o *RateLimiter) GetResponse() RateLimiterResponse1`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *RateLimiter) GetResponseOk() (*RateLimiterResponse1, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *RateLimiter) SetResponse(v RateLimiterResponse1)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *RateLimiter) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### SetResponseNil

`func (o *RateLimiter) SetResponseNil(b bool)`

 SetResponseNil sets the value for Response to be an explicit nil

### UnsetResponse
`func (o *RateLimiter) UnsetResponse()`

UnsetResponse ensures that no value is present for Response, not even an explicit nil
### GetResponseObjectName

`func (o *RateLimiter) GetResponseObjectName() string`

GetResponseObjectName returns the ResponseObjectName field if non-nil, zero value otherwise.

### GetResponseObjectNameOk

`func (o *RateLimiter) GetResponseObjectNameOk() (*string, bool)`

GetResponseObjectNameOk returns a tuple with the ResponseObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseObjectName

`func (o *RateLimiter) SetResponseObjectName(v string)`

SetResponseObjectName sets ResponseObjectName field to given value.

### HasResponseObjectName

`func (o *RateLimiter) HasResponseObjectName() bool`

HasResponseObjectName returns a boolean if a field has been set.

### SetResponseObjectNameNil

`func (o *RateLimiter) SetResponseObjectNameNil(b bool)`

 SetResponseObjectNameNil sets the value for ResponseObjectName to be an explicit nil

### UnsetResponseObjectName
`func (o *RateLimiter) UnsetResponseObjectName()`

UnsetResponseObjectName ensures that no value is present for ResponseObjectName, not even an explicit nil
### GetLoggerType

`func (o *RateLimiter) GetLoggerType() string`

GetLoggerType returns the LoggerType field if non-nil, zero value otherwise.

### GetLoggerTypeOk

`func (o *RateLimiter) GetLoggerTypeOk() (*string, bool)`

GetLoggerTypeOk returns a tuple with the LoggerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggerType

`func (o *RateLimiter) SetLoggerType(v string)`

SetLoggerType sets LoggerType field to given value.

### HasLoggerType

`func (o *RateLimiter) HasLoggerType() bool`

HasLoggerType returns a boolean if a field has been set.

### GetFeatureRevision

`func (o *RateLimiter) GetFeatureRevision() int32`

GetFeatureRevision returns the FeatureRevision field if non-nil, zero value otherwise.

### GetFeatureRevisionOk

`func (o *RateLimiter) GetFeatureRevisionOk() (*int32, bool)`

GetFeatureRevisionOk returns a tuple with the FeatureRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureRevision

`func (o *RateLimiter) SetFeatureRevision(v int32)`

SetFeatureRevision sets FeatureRevision field to given value.

### HasFeatureRevision

`func (o *RateLimiter) HasFeatureRevision() bool`

HasFeatureRevision returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
