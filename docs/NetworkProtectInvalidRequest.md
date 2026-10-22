# NetworkProtectInvalidRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** | Short summary of the validation problem. | 
**Status** | **int32** | HTTP status code for this error. | 
**Detail** | **string** | Corrective explanation for the invalid request. | 
**Errors** | [**[]NetworkProtectErrorDetail**](NetworkProtectErrorDetail.md) | Field-level validation errors, when available. | 

## Methods

### NewNetworkProtectInvalidRequest

`func NewNetworkProtectInvalidRequest(title string, status int32, detail string, errors []NetworkProtectErrorDetail, ) *NetworkProtectInvalidRequest`

NewNetworkProtectInvalidRequest instantiates a new NetworkProtectInvalidRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkProtectInvalidRequestWithDefaults

`func NewNetworkProtectInvalidRequestWithDefaults() *NetworkProtectInvalidRequest`

NewNetworkProtectInvalidRequestWithDefaults instantiates a new NetworkProtectInvalidRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *NetworkProtectInvalidRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *NetworkProtectInvalidRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *NetworkProtectInvalidRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetStatus

`func (o *NetworkProtectInvalidRequest) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NetworkProtectInvalidRequest) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NetworkProtectInvalidRequest) SetStatus(v int32)`

SetStatus sets Status field to given value.


### GetDetail

`func (o *NetworkProtectInvalidRequest) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *NetworkProtectInvalidRequest) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *NetworkProtectInvalidRequest) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetErrors

`func (o *NetworkProtectInvalidRequest) GetErrors() []NetworkProtectErrorDetail`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *NetworkProtectInvalidRequest) GetErrorsOk() (*[]NetworkProtectErrorDetail, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *NetworkProtectInvalidRequest) SetErrors(v []NetworkProtectErrorDetail)`

SetErrors sets Errors field to given value.



[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


