# NetworkProtectConflictError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | Pointer to **string** | Request field associated with the conflict. | [optional] 
**Error** | Pointer to **string** | Explanation of the conflict for the field. | [optional] 

## Methods

### NewNetworkProtectConflictError

`func NewNetworkProtectConflictError() *NetworkProtectConflictError`

NewNetworkProtectConflictError instantiates a new NetworkProtectConflictError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkProtectConflictErrorWithDefaults

`func NewNetworkProtectConflictErrorWithDefaults() *NetworkProtectConflictError`

NewNetworkProtectConflictErrorWithDefaults instantiates a new NetworkProtectConflictError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *NetworkProtectConflictError) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *NetworkProtectConflictError) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *NetworkProtectConflictError) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *NetworkProtectConflictError) HasField() bool`

HasField returns a boolean if a field has been set.

### GetError

`func (o *NetworkProtectConflictError) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *NetworkProtectConflictError) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *NetworkProtectConflictError) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *NetworkProtectConflictError) HasError() bool`

HasError returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
