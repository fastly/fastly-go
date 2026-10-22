# NetworkProtectConflict

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** | Short summary of the conflict. | 
**Status** | **int32** | HTTP status code for this error. | 
**Detail** | **string** | Explanation of the conflicting resource state. | 
**Errors** | Pointer to [**[]NetworkProtectConflictError**](NetworkProtectConflictError.md) | Optional conflict error details. | [optional] 

## Methods

### NewNetworkProtectConflict

`func NewNetworkProtectConflict(title string, status int32, detail string, ) *NetworkProtectConflict`

NewNetworkProtectConflict instantiates a new NetworkProtectConflict object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkProtectConflictWithDefaults

`func NewNetworkProtectConflictWithDefaults() *NetworkProtectConflict`

NewNetworkProtectConflictWithDefaults instantiates a new NetworkProtectConflict object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *NetworkProtectConflict) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *NetworkProtectConflict) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *NetworkProtectConflict) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetStatus

`func (o *NetworkProtectConflict) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NetworkProtectConflict) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NetworkProtectConflict) SetStatus(v int32)`

SetStatus sets Status field to given value.


### GetDetail

`func (o *NetworkProtectConflict) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *NetworkProtectConflict) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *NetworkProtectConflict) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetErrors

`func (o *NetworkProtectConflict) GetErrors() []NetworkProtectConflictError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *NetworkProtectConflict) GetErrorsOk() (*[]NetworkProtectConflictError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *NetworkProtectConflict) SetErrors(v []NetworkProtectConflictError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *NetworkProtectConflict) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


