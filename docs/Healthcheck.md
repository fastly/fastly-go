# Healthcheck

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

## Methods

### NewHealthcheck

`func NewHealthcheck() *Healthcheck`

NewHealthcheck instantiates a new Healthcheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthcheckWithDefaults

`func NewHealthcheckWithDefaults() *Healthcheck`

NewHealthcheckWithDefaults instantiates a new Healthcheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckInterval

`func (o *Healthcheck) GetCheckInterval() int32`

GetCheckInterval returns the CheckInterval field if non-nil, zero value otherwise.

### GetCheckIntervalOk

`func (o *Healthcheck) GetCheckIntervalOk() (*int32, bool)`

GetCheckIntervalOk returns a tuple with the CheckInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckInterval

`func (o *Healthcheck) SetCheckInterval(v int32)`

SetCheckInterval sets CheckInterval field to given value.

### HasCheckInterval

`func (o *Healthcheck) HasCheckInterval() bool`

HasCheckInterval returns a boolean if a field has been set.

### GetComment

`func (o *Healthcheck) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Healthcheck) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Healthcheck) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Healthcheck) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *Healthcheck) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *Healthcheck) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetExpectedResponse

`func (o *Healthcheck) GetExpectedResponse() int32`

GetExpectedResponse returns the ExpectedResponse field if non-nil, zero value otherwise.

### GetExpectedResponseOk

`func (o *Healthcheck) GetExpectedResponseOk() (*int32, bool)`

GetExpectedResponseOk returns a tuple with the ExpectedResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedResponse

`func (o *Healthcheck) SetExpectedResponse(v int32)`

SetExpectedResponse sets ExpectedResponse field to given value.

### HasExpectedResponse

`func (o *Healthcheck) HasExpectedResponse() bool`

HasExpectedResponse returns a boolean if a field has been set.

### GetHeaders

`func (o *Healthcheck) GetHeaders() []string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *Healthcheck) GetHeadersOk() (*[]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *Healthcheck) SetHeaders(v []string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *Healthcheck) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetHost

`func (o *Healthcheck) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *Healthcheck) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *Healthcheck) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *Healthcheck) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetHTTPVersion

`func (o *Healthcheck) GetHTTPVersion() string`

GetHTTPVersion returns the HTTPVersion field if non-nil, zero value otherwise.

### GetHTTPVersionOk

`func (o *Healthcheck) GetHTTPVersionOk() (*string, bool)`

GetHTTPVersionOk returns a tuple with the HTTPVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHTTPVersion

`func (o *Healthcheck) SetHTTPVersion(v string)`

SetHTTPVersion sets HTTPVersion field to given value.

### HasHTTPVersion

`func (o *Healthcheck) HasHTTPVersion() bool`

HasHTTPVersion returns a boolean if a field has been set.

### GetInitial

`func (o *Healthcheck) GetInitial() int32`

GetInitial returns the Initial field if non-nil, zero value otherwise.

### GetInitialOk

`func (o *Healthcheck) GetInitialOk() (*int32, bool)`

GetInitialOk returns a tuple with the Initial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitial

`func (o *Healthcheck) SetInitial(v int32)`

SetInitial sets Initial field to given value.

### HasInitial

`func (o *Healthcheck) HasInitial() bool`

HasInitial returns a boolean if a field has been set.

### GetMethod

`func (o *Healthcheck) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *Healthcheck) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *Healthcheck) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *Healthcheck) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetName

`func (o *Healthcheck) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Healthcheck) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Healthcheck) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Healthcheck) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPath

`func (o *Healthcheck) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *Healthcheck) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *Healthcheck) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *Healthcheck) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetThreshold

`func (o *Healthcheck) GetThreshold() int32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *Healthcheck) GetThresholdOk() (*int32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *Healthcheck) SetThreshold(v int32)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *Healthcheck) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetTimeout

`func (o *Healthcheck) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *Healthcheck) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *Healthcheck) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *Healthcheck) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetWindow

`func (o *Healthcheck) GetWindow() int32`

GetWindow returns the Window field if non-nil, zero value otherwise.

### GetWindowOk

`func (o *Healthcheck) GetWindowOk() (*int32, bool)`

GetWindowOk returns a tuple with the Window field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindow

`func (o *Healthcheck) SetWindow(v int32)`

SetWindow sets Window field to given value.

### HasWindow

`func (o *Healthcheck) HasWindow() bool`

HasWindow returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
