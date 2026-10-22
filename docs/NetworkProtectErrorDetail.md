# NetworkProtectErrorDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | **string** | Request field associated with the error. | 
**Error** | **string** | Explanation of why the field is invalid. | 

## Methods

### NewNetworkProtectErrorDetail

`func NewNetworkProtectErrorDetail(field string, error_ string, ) *NetworkProtectErrorDetail`

NewNetworkProtectErrorDetail instantiates a new NetworkProtectErrorDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkProtectErrorDetailWithDefaults

`func NewNetworkProtectErrorDetailWithDefaults() *NetworkProtectErrorDetail`

NewNetworkProtectErrorDetailWithDefaults instantiates a new NetworkProtectErrorDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *NetworkProtectErrorDetail) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *NetworkProtectErrorDetail) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *NetworkProtectErrorDetail) SetField(v string)`

SetField sets Field field to given value.


### GetError

`func (o *NetworkProtectErrorDetail) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *NetworkProtectErrorDetail) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *NetworkProtectErrorDetail) SetError(v string)`

SetError sets Error field to given value.



[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


