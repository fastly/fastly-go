# LogError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SequenceNumber** | Pointer to **int32** | Sequence number for ordering messages. | [optional] 
**ErrorTimeUs** | Pointer to **int64** | Timestamp of the error in microseconds. | [optional] 
**Stream** | Pointer to **string** | The stream type, always &#39;logging_error&#39; for logging endpoint errors. | [optional] 
**Message** | Pointer to **string** | User-friendly error message. | [optional] 
**Endpoint** | Pointer to **string** | Name of the logging endpoint that generated the error. | [optional] 
**Details** | Pointer to **string** | Additional error details as a JSON string. | [optional] 

## Methods

### NewLogError

`func NewLogError() *LogError`

NewLogError instantiates a new LogError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogErrorWithDefaults

`func NewLogErrorWithDefaults() *LogError`

NewLogErrorWithDefaults instantiates a new LogError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSequenceNumber

`func (o *LogError) GetSequenceNumber() int32`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *LogError) GetSequenceNumberOk() (*int32, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *LogError) SetSequenceNumber(v int32)`

SetSequenceNumber sets SequenceNumber field to given value.

### HasSequenceNumber

`func (o *LogError) HasSequenceNumber() bool`

HasSequenceNumber returns a boolean if a field has been set.

### GetErrorTimeUs

`func (o *LogError) GetErrorTimeUs() int64`

GetErrorTimeUs returns the ErrorTimeUs field if non-nil, zero value otherwise.

### GetErrorTimeUsOk

`func (o *LogError) GetErrorTimeUsOk() (*int64, bool)`

GetErrorTimeUsOk returns a tuple with the ErrorTimeUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorTimeUs

`func (o *LogError) SetErrorTimeUs(v int64)`

SetErrorTimeUs sets ErrorTimeUs field to given value.

### HasErrorTimeUs

`func (o *LogError) HasErrorTimeUs() bool`

HasErrorTimeUs returns a boolean if a field has been set.

### GetStream

`func (o *LogError) GetStream() string`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *LogError) GetStreamOk() (*string, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *LogError) SetStream(v string)`

SetStream sets Stream field to given value.

### HasStream

`func (o *LogError) HasStream() bool`

HasStream returns a boolean if a field has been set.

### GetMessage

`func (o *LogError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *LogError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *LogError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *LogError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetEndpoint

`func (o *LogError) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *LogError) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *LogError) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *LogError) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetDetails

`func (o *LogError) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *LogError) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *LogError) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *LogError) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


