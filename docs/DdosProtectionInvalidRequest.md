# DdosProtectionInvalidRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** |  | 
**Status** | **int32** |  | 
**Detail** | **string** |  | 
**Errors** | [**[]DdosProtectionErrorErrors**](DdosProtectionErrorErrors.md) |  | 

## Methods

### NewDdosProtectionInvalidRequest

`func NewDdosProtectionInvalidRequest(title string, status int32, detail string, errors []DdosProtectionErrorErrors, ) *DdosProtectionInvalidRequest`

NewDdosProtectionInvalidRequest instantiates a new DdosProtectionInvalidRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDdosProtectionInvalidRequestWithDefaults

`func NewDdosProtectionInvalidRequestWithDefaults() *DdosProtectionInvalidRequest`

NewDdosProtectionInvalidRequestWithDefaults instantiates a new DdosProtectionInvalidRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *DdosProtectionInvalidRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DdosProtectionInvalidRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DdosProtectionInvalidRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetStatus

`func (o *DdosProtectionInvalidRequest) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DdosProtectionInvalidRequest) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DdosProtectionInvalidRequest) SetStatus(v int32)`

SetStatus sets Status field to given value.


### GetDetail

`func (o *DdosProtectionInvalidRequest) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *DdosProtectionInvalidRequest) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *DdosProtectionInvalidRequest) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetErrors

`func (o *DdosProtectionInvalidRequest) GetErrors() []DdosProtectionErrorErrors`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *DdosProtectionInvalidRequest) GetErrorsOk() (*[]DdosProtectionErrorErrors, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *DdosProtectionInvalidRequest) SetErrors(v []DdosProtectionErrorErrors)`

SetErrors sets Errors field to given value.



[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


