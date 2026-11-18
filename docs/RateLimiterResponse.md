# RateLimiterResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A human readable name for the rate limiting rule. | [optional] 
**URIDictionaryName** | Pointer to **NullableString** | The name of a Dictionary containing URIs as keys. If not defined or `null`, all origin URIs will be rate limited. | [optional] 
**HTTPMethods** | Pointer to **[]string** | Array of HTTP methods to apply rate limiting to. | [optional] 
**RpsLimit** | Pointer to **int32** | Upper limit of requests per second allowed by the rate limiter. | [optional] 
**WindowSize** | Pointer to **int32** | Number of seconds during which the RPS limit must be exceeded in order to trigger a violation. | [optional] 
**ClientKey** | Pointer to **[]string** | Array of VCL variables used to generate a counter key to identify a client. Example variables include `req.http.Fastly-Client-IP`, `req.http.User-Agent`, or a custom header like `req.http.API-Key`. | [optional] 
**PenaltyBoxDuration** | Pointer to **int32** | Length of time in minutes that the rate limiter is in effect after the initial violation is detected. | [optional] 
**Action** | Pointer to **string** | The action to take when a rate limiter violation is detected. | [optional] 
**Response** | Pointer to **map[string]string** | Custom response to be sent when the rate limit is exceeded. Required if `action` is `response`. | [optional] 
**ResponseObjectName** | Pointer to **NullableString** | Name of existing response object. Required if `action` is `response_object`. Note that the rate limiter response is only updated to reflect the response object content when saving the rate limiter configuration. | [optional] 
**LoggerType** | Pointer to **string** | Name of the type of logging endpoint to be used when action is `log_only`. The logging endpoint type is used to determine the appropriate log format to use when emitting log entries. | [optional] 
**FeatureRevision** | Pointer to **int32** | Revision number of the rate limiting feature implementation. Defaults to the most recent revision. | [optional] 
**ServiceID** | Pointer to **string** |  | [optional] [readonly] 
**Version** | Pointer to **int32** |  | [optional] [readonly] 
**CreatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**DeletedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**UpdatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**ID** | Pointer to **string** | Alphanumeric string identifying the rate limiter. | [optional] 

## Methods

### NewRateLimiterResponse

`func NewRateLimiterResponse() *RateLimiterResponse`

NewRateLimiterResponse instantiates a new RateLimiterResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRateLimiterResponseWithDefaults

`func NewRateLimiterResponseWithDefaults() *RateLimiterResponse`

NewRateLimiterResponseWithDefaults instantiates a new RateLimiterResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RateLimiterResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RateLimiterResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RateLimiterResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RateLimiterResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetURIDictionaryName

`func (o *RateLimiterResponse) GetURIDictionaryName() string`

GetURIDictionaryName returns the URIDictionaryName field if non-nil, zero value otherwise.

### GetURIDictionaryNameOk

`func (o *RateLimiterResponse) GetURIDictionaryNameOk() (*string, bool)`

GetURIDictionaryNameOk returns a tuple with the URIDictionaryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetURIDictionaryName

`func (o *RateLimiterResponse) SetURIDictionaryName(v string)`

SetURIDictionaryName sets URIDictionaryName field to given value.

### HasURIDictionaryName

`func (o *RateLimiterResponse) HasURIDictionaryName() bool`

HasURIDictionaryName returns a boolean if a field has been set.

### SetURIDictionaryNameNil

`func (o *RateLimiterResponse) SetURIDictionaryNameNil(b bool)`

 SetURIDictionaryNameNil sets the value for URIDictionaryName to be an explicit nil

### UnsetURIDictionaryName
`func (o *RateLimiterResponse) UnsetURIDictionaryName()`

UnsetURIDictionaryName ensures that no value is present for URIDictionaryName, not even an explicit nil
### GetHTTPMethods

`func (o *RateLimiterResponse) GetHTTPMethods() []string`

GetHTTPMethods returns the HTTPMethods field if non-nil, zero value otherwise.

### GetHTTPMethodsOk

`func (o *RateLimiterResponse) GetHTTPMethodsOk() (*[]string, bool)`

GetHTTPMethodsOk returns a tuple with the HTTPMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHTTPMethods

`func (o *RateLimiterResponse) SetHTTPMethods(v []string)`

SetHTTPMethods sets HTTPMethods field to given value.

### HasHTTPMethods

`func (o *RateLimiterResponse) HasHTTPMethods() bool`

HasHTTPMethods returns a boolean if a field has been set.

### GetRpsLimit

`func (o *RateLimiterResponse) GetRpsLimit() int32`

GetRpsLimit returns the RpsLimit field if non-nil, zero value otherwise.

### GetRpsLimitOk

`func (o *RateLimiterResponse) GetRpsLimitOk() (*int32, bool)`

GetRpsLimitOk returns a tuple with the RpsLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpsLimit

`func (o *RateLimiterResponse) SetRpsLimit(v int32)`

SetRpsLimit sets RpsLimit field to given value.

### HasRpsLimit

`func (o *RateLimiterResponse) HasRpsLimit() bool`

HasRpsLimit returns a boolean if a field has been set.

### GetWindowSize

`func (o *RateLimiterResponse) GetWindowSize() int32`

GetWindowSize returns the WindowSize field if non-nil, zero value otherwise.

### GetWindowSizeOk

`func (o *RateLimiterResponse) GetWindowSizeOk() (*int32, bool)`

GetWindowSizeOk returns a tuple with the WindowSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowSize

`func (o *RateLimiterResponse) SetWindowSize(v int32)`

SetWindowSize sets WindowSize field to given value.

### HasWindowSize

`func (o *RateLimiterResponse) HasWindowSize() bool`

HasWindowSize returns a boolean if a field has been set.

### GetClientKey

`func (o *RateLimiterResponse) GetClientKey() []string`

GetClientKey returns the ClientKey field if non-nil, zero value otherwise.

### GetClientKeyOk

`func (o *RateLimiterResponse) GetClientKeyOk() (*[]string, bool)`

GetClientKeyOk returns a tuple with the ClientKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientKey

`func (o *RateLimiterResponse) SetClientKey(v []string)`

SetClientKey sets ClientKey field to given value.

### HasClientKey

`func (o *RateLimiterResponse) HasClientKey() bool`

HasClientKey returns a boolean if a field has been set.

### GetPenaltyBoxDuration

`func (o *RateLimiterResponse) GetPenaltyBoxDuration() int32`

GetPenaltyBoxDuration returns the PenaltyBoxDuration field if non-nil, zero value otherwise.

### GetPenaltyBoxDurationOk

`func (o *RateLimiterResponse) GetPenaltyBoxDurationOk() (*int32, bool)`

GetPenaltyBoxDurationOk returns a tuple with the PenaltyBoxDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPenaltyBoxDuration

`func (o *RateLimiterResponse) SetPenaltyBoxDuration(v int32)`

SetPenaltyBoxDuration sets PenaltyBoxDuration field to given value.

### HasPenaltyBoxDuration

`func (o *RateLimiterResponse) HasPenaltyBoxDuration() bool`

HasPenaltyBoxDuration returns a boolean if a field has been set.

### GetAction

`func (o *RateLimiterResponse) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *RateLimiterResponse) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *RateLimiterResponse) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *RateLimiterResponse) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetResponse

`func (o *RateLimiterResponse) GetResponse() map[string]string`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *RateLimiterResponse) GetResponseOk() (*map[string]string, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *RateLimiterResponse) SetResponse(v map[string]string)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *RateLimiterResponse) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### SetResponseNil

`func (o *RateLimiterResponse) SetResponseNil(b bool)`

 SetResponseNil sets the value for Response to be an explicit nil

### UnsetResponse
`func (o *RateLimiterResponse) UnsetResponse()`

UnsetResponse ensures that no value is present for Response, not even an explicit nil
### GetResponseObjectName

`func (o *RateLimiterResponse) GetResponseObjectName() string`

GetResponseObjectName returns the ResponseObjectName field if non-nil, zero value otherwise.

### GetResponseObjectNameOk

`func (o *RateLimiterResponse) GetResponseObjectNameOk() (*string, bool)`

GetResponseObjectNameOk returns a tuple with the ResponseObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseObjectName

`func (o *RateLimiterResponse) SetResponseObjectName(v string)`

SetResponseObjectName sets ResponseObjectName field to given value.

### HasResponseObjectName

`func (o *RateLimiterResponse) HasResponseObjectName() bool`

HasResponseObjectName returns a boolean if a field has been set.

### SetResponseObjectNameNil

`func (o *RateLimiterResponse) SetResponseObjectNameNil(b bool)`

 SetResponseObjectNameNil sets the value for ResponseObjectName to be an explicit nil

### UnsetResponseObjectName
`func (o *RateLimiterResponse) UnsetResponseObjectName()`

UnsetResponseObjectName ensures that no value is present for ResponseObjectName, not even an explicit nil
### GetLoggerType

`func (o *RateLimiterResponse) GetLoggerType() string`

GetLoggerType returns the LoggerType field if non-nil, zero value otherwise.

### GetLoggerTypeOk

`func (o *RateLimiterResponse) GetLoggerTypeOk() (*string, bool)`

GetLoggerTypeOk returns a tuple with the LoggerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggerType

`func (o *RateLimiterResponse) SetLoggerType(v string)`

SetLoggerType sets LoggerType field to given value.

### HasLoggerType

`func (o *RateLimiterResponse) HasLoggerType() bool`

HasLoggerType returns a boolean if a field has been set.

### GetFeatureRevision

`func (o *RateLimiterResponse) GetFeatureRevision() int32`

GetFeatureRevision returns the FeatureRevision field if non-nil, zero value otherwise.

### GetFeatureRevisionOk

`func (o *RateLimiterResponse) GetFeatureRevisionOk() (*int32, bool)`

GetFeatureRevisionOk returns a tuple with the FeatureRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureRevision

`func (o *RateLimiterResponse) SetFeatureRevision(v int32)`

SetFeatureRevision sets FeatureRevision field to given value.

### HasFeatureRevision

`func (o *RateLimiterResponse) HasFeatureRevision() bool`

HasFeatureRevision returns a boolean if a field has been set.

### GetServiceID

`func (o *RateLimiterResponse) GetServiceID() string`

GetServiceID returns the ServiceID field if non-nil, zero value otherwise.

### GetServiceIDOk

`func (o *RateLimiterResponse) GetServiceIDOk() (*string, bool)`

GetServiceIDOk returns a tuple with the ServiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceID

`func (o *RateLimiterResponse) SetServiceID(v string)`

SetServiceID sets ServiceID field to given value.

### HasServiceID

`func (o *RateLimiterResponse) HasServiceID() bool`

HasServiceID returns a boolean if a field has been set.

### GetVersion

`func (o *RateLimiterResponse) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *RateLimiterResponse) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *RateLimiterResponse) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *RateLimiterResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RateLimiterResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RateLimiterResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RateLimiterResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RateLimiterResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *RateLimiterResponse) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *RateLimiterResponse) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetDeletedAt

`func (o *RateLimiterResponse) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *RateLimiterResponse) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *RateLimiterResponse) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *RateLimiterResponse) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *RateLimiterResponse) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *RateLimiterResponse) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetUpdatedAt

`func (o *RateLimiterResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RateLimiterResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RateLimiterResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RateLimiterResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *RateLimiterResponse) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *RateLimiterResponse) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetID

`func (o *RateLimiterResponse) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *RateLimiterResponse) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *RateLimiterResponse) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *RateLimiterResponse) HasID() bool`

HasID returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
