# HealthcheckResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CheckInterval** | Pointer to **int32** | How often to run the health check in milliseconds. | [optional] 
**Comment** | Pointer to **NullableString** | A freeform descriptive note. | [optional] 
**ExpectedResponse** | Pointer to **int32** | The status code expected from the host. | [optional] 
**Headers** | Pointer to **[]string** | Array of custom headers that will be added to the health check probes. | [optional] 
**Host** | Pointer to **string** | Which host to check. | [optional] 
**HTTPVersion** | Pointer to **string** | Whether to use version 1.0 or 1.1 HTTP. | [optional] 
**Initial** | Pointer to **int32** | When loading a config, the initial number of probes to be seen as OK. | [optional] 
**Method** | Pointer to **string** | Which HTTP method to use. | [optional] 
**Name** | Pointer to **string** | The name of the health check. | [optional] 
**Path** | Pointer to **string** | The path to check. | [optional] 
**Threshold** | Pointer to **int32** | How many health checks must succeed to be considered healthy. | [optional] 
**Timeout** | Pointer to **int32** | Timeout in milliseconds. | [optional] 
**Window** | Pointer to **int32** | The number of most recent health check queries to keep for this health check. | [optional] 
**ServiceID** | Pointer to **string** |  | [optional] [readonly] 
**Version** | Pointer to **int32** |  | [optional] [readonly] 
**CreatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**DeletedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**UpdatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 

## Methods

### NewHealthcheckResponse

`func NewHealthcheckResponse() *HealthcheckResponse`

NewHealthcheckResponse instantiates a new HealthcheckResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthcheckResponseWithDefaults

`func NewHealthcheckResponseWithDefaults() *HealthcheckResponse`

NewHealthcheckResponseWithDefaults instantiates a new HealthcheckResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckInterval

`func (o *HealthcheckResponse) GetCheckInterval() int32`

GetCheckInterval returns the CheckInterval field if non-nil, zero value otherwise.

### GetCheckIntervalOk

`func (o *HealthcheckResponse) GetCheckIntervalOk() (*int32, bool)`

GetCheckIntervalOk returns a tuple with the CheckInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckInterval

`func (o *HealthcheckResponse) SetCheckInterval(v int32)`

SetCheckInterval sets CheckInterval field to given value.

### HasCheckInterval

`func (o *HealthcheckResponse) HasCheckInterval() bool`

HasCheckInterval returns a boolean if a field has been set.

### GetComment

`func (o *HealthcheckResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *HealthcheckResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *HealthcheckResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *HealthcheckResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *HealthcheckResponse) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *HealthcheckResponse) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetExpectedResponse

`func (o *HealthcheckResponse) GetExpectedResponse() int32`

GetExpectedResponse returns the ExpectedResponse field if non-nil, zero value otherwise.

### GetExpectedResponseOk

`func (o *HealthcheckResponse) GetExpectedResponseOk() (*int32, bool)`

GetExpectedResponseOk returns a tuple with the ExpectedResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedResponse

`func (o *HealthcheckResponse) SetExpectedResponse(v int32)`

SetExpectedResponse sets ExpectedResponse field to given value.

### HasExpectedResponse

`func (o *HealthcheckResponse) HasExpectedResponse() bool`

HasExpectedResponse returns a boolean if a field has been set.

### GetHeaders

`func (o *HealthcheckResponse) GetHeaders() []string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *HealthcheckResponse) GetHeadersOk() (*[]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *HealthcheckResponse) SetHeaders(v []string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *HealthcheckResponse) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetHost

`func (o *HealthcheckResponse) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *HealthcheckResponse) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *HealthcheckResponse) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *HealthcheckResponse) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetHTTPVersion

`func (o *HealthcheckResponse) GetHTTPVersion() string`

GetHTTPVersion returns the HTTPVersion field if non-nil, zero value otherwise.

### GetHTTPVersionOk

`func (o *HealthcheckResponse) GetHTTPVersionOk() (*string, bool)`

GetHTTPVersionOk returns a tuple with the HTTPVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHTTPVersion

`func (o *HealthcheckResponse) SetHTTPVersion(v string)`

SetHTTPVersion sets HTTPVersion field to given value.

### HasHTTPVersion

`func (o *HealthcheckResponse) HasHTTPVersion() bool`

HasHTTPVersion returns a boolean if a field has been set.

### GetInitial

`func (o *HealthcheckResponse) GetInitial() int32`

GetInitial returns the Initial field if non-nil, zero value otherwise.

### GetInitialOk

`func (o *HealthcheckResponse) GetInitialOk() (*int32, bool)`

GetInitialOk returns a tuple with the Initial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitial

`func (o *HealthcheckResponse) SetInitial(v int32)`

SetInitial sets Initial field to given value.

### HasInitial

`func (o *HealthcheckResponse) HasInitial() bool`

HasInitial returns a boolean if a field has been set.

### GetMethod

`func (o *HealthcheckResponse) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *HealthcheckResponse) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *HealthcheckResponse) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *HealthcheckResponse) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetName

`func (o *HealthcheckResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HealthcheckResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HealthcheckResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HealthcheckResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPath

`func (o *HealthcheckResponse) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *HealthcheckResponse) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *HealthcheckResponse) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *HealthcheckResponse) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetThreshold

`func (o *HealthcheckResponse) GetThreshold() int32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *HealthcheckResponse) GetThresholdOk() (*int32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *HealthcheckResponse) SetThreshold(v int32)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *HealthcheckResponse) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetTimeout

`func (o *HealthcheckResponse) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *HealthcheckResponse) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *HealthcheckResponse) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *HealthcheckResponse) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetWindow

`func (o *HealthcheckResponse) GetWindow() int32`

GetWindow returns the Window field if non-nil, zero value otherwise.

### GetWindowOk

`func (o *HealthcheckResponse) GetWindowOk() (*int32, bool)`

GetWindowOk returns a tuple with the Window field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindow

`func (o *HealthcheckResponse) SetWindow(v int32)`

SetWindow sets Window field to given value.

### HasWindow

`func (o *HealthcheckResponse) HasWindow() bool`

HasWindow returns a boolean if a field has been set.

### GetServiceID

`func (o *HealthcheckResponse) GetServiceID() string`

GetServiceID returns the ServiceID field if non-nil, zero value otherwise.

### GetServiceIDOk

`func (o *HealthcheckResponse) GetServiceIDOk() (*string, bool)`

GetServiceIDOk returns a tuple with the ServiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceID

`func (o *HealthcheckResponse) SetServiceID(v string)`

SetServiceID sets ServiceID field to given value.

### HasServiceID

`func (o *HealthcheckResponse) HasServiceID() bool`

HasServiceID returns a boolean if a field has been set.

### GetVersion

`func (o *HealthcheckResponse) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *HealthcheckResponse) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *HealthcheckResponse) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *HealthcheckResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetCreatedAt

`func (o *HealthcheckResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *HealthcheckResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *HealthcheckResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *HealthcheckResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *HealthcheckResponse) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *HealthcheckResponse) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetDeletedAt

`func (o *HealthcheckResponse) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *HealthcheckResponse) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *HealthcheckResponse) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *HealthcheckResponse) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *HealthcheckResponse) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *HealthcheckResponse) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetUpdatedAt

`func (o *HealthcheckResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *HealthcheckResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *HealthcheckResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *HealthcheckResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *HealthcheckResponse) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *HealthcheckResponse) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
