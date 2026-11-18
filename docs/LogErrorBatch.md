# LogErrorBatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BatchID** | Pointer to **string** | Unique identifier for this batch of logs. | [optional] 
**Logs** | Pointer to [**[]LogError**](LogError.md) |  | [optional] 

## Methods

### NewLogErrorBatch

`func NewLogErrorBatch() *LogErrorBatch`

NewLogErrorBatch instantiates a new LogErrorBatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogErrorBatchWithDefaults

`func NewLogErrorBatchWithDefaults() *LogErrorBatch`

NewLogErrorBatchWithDefaults instantiates a new LogErrorBatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBatchID

`func (o *LogErrorBatch) GetBatchID() string`

GetBatchID returns the BatchID field if non-nil, zero value otherwise.

### GetBatchIDOk

`func (o *LogErrorBatch) GetBatchIDOk() (*string, bool)`

GetBatchIDOk returns a tuple with the BatchID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchID

`func (o *LogErrorBatch) SetBatchID(v string)`

SetBatchID sets BatchID field to given value.

### HasBatchID

`func (o *LogErrorBatch) HasBatchID() bool`

HasBatchID returns a boolean if a field has been set.

### GetLogs

`func (o *LogErrorBatch) GetLogs() []LogError`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *LogErrorBatch) GetLogsOk() (*[]LogError, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *LogErrorBatch) SetLogs(v []LogError)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *LogErrorBatch) HasLogs() bool`

HasLogs returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
