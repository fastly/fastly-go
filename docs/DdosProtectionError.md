# DdosProtectionError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int32** |  | [optional] 
**Detail** | Pointer to **string** |  | [optional] 
**Errors** | Pointer to [**[]DdosProtectionErrorErrors**](DdosProtectionErrorErrors.md) |  | [optional] 

## Methods

### NewDdosProtectionError

`func NewDdosProtectionError() *DdosProtectionError`

NewDdosProtectionError instantiates a new DdosProtectionError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDdosProtectionErrorWithDefaults

`func NewDdosProtectionErrorWithDefaults() *DdosProtectionError`

NewDdosProtectionErrorWithDefaults instantiates a new DdosProtectionError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *DdosProtectionError) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DdosProtectionError) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DdosProtectionError) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *DdosProtectionError) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetStatus

`func (o *DdosProtectionError) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DdosProtectionError) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DdosProtectionError) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DdosProtectionError) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDetail

`func (o *DdosProtectionError) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *DdosProtectionError) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *DdosProtectionError) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *DdosProtectionError) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetErrors

`func (o *DdosProtectionError) GetErrors() []DdosProtectionErrorErrors`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *DdosProtectionError) GetErrorsOk() (*[]DdosProtectionErrorErrors, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *DdosProtectionError) SetErrors(v []DdosProtectionErrorErrors)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *DdosProtectionError) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


