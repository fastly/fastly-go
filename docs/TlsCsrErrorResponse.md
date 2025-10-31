# TlsCsrErrorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to [**[]ErrorResponseData**](ErrorResponseData.md) |  | [optional] 

## Methods

### NewTlsCsrErrorResponse

`func NewTlsCsrErrorResponse() *TlsCsrErrorResponse`

NewTlsCsrErrorResponse instantiates a new TlsCsrErrorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTlsCsrErrorResponseWithDefaults

`func NewTlsCsrErrorResponseWithDefaults() *TlsCsrErrorResponse`

NewTlsCsrErrorResponseWithDefaults instantiates a new TlsCsrErrorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *TlsCsrErrorResponse) GetErrors() []ErrorResponseData`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *TlsCsrErrorResponse) GetErrorsOk() (*[]ErrorResponseData, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *TlsCsrErrorResponse) SetErrors(v []ErrorResponseData)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *TlsCsrErrorResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


