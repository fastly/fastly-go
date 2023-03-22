# BillingAddressVerificationErrorResponseErrors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The error type. | 
**Title** | **string** |  | 
**Detail** | **string** |  | 
**Status** | **float32** |  | 
**Candidates** | Pointer to [**[]BillingAddressAttributes**](BillingAddressAttributes.md) |  | [optional] 

## Methods

### NewBillingAddressVerificationErrorResponseErrors

`func NewBillingAddressVerificationErrorResponseErrors(resourceType string, title string, detail string, status float32, ) *BillingAddressVerificationErrorResponseErrors`

NewBillingAddressVerificationErrorResponseErrors instantiates a new BillingAddressVerificationErrorResponseErrors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingAddressVerificationErrorResponseErrorsWithDefaults

`func NewBillingAddressVerificationErrorResponseErrorsWithDefaults() *BillingAddressVerificationErrorResponseErrors`

NewBillingAddressVerificationErrorResponseErrorsWithDefaults instantiates a new BillingAddressVerificationErrorResponseErrors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BillingAddressVerificationErrorResponseErrors) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BillingAddressVerificationErrorResponseErrors) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BillingAddressVerificationErrorResponseErrors) SetType(v string)`

SetType sets Type field to given value.


### GetTitle

`func (o *BillingAddressVerificationErrorResponseErrors) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *BillingAddressVerificationErrorResponseErrors) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *BillingAddressVerificationErrorResponseErrors) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDetail

`func (o *BillingAddressVerificationErrorResponseErrors) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *BillingAddressVerificationErrorResponseErrors) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *BillingAddressVerificationErrorResponseErrors) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetStatus

`func (o *BillingAddressVerificationErrorResponseErrors) GetStatus() float32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BillingAddressVerificationErrorResponseErrors) GetStatusOk() (*float32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BillingAddressVerificationErrorResponseErrors) SetStatus(v float32)`

SetStatus sets Status field to given value.


### GetCandidates

`func (o *BillingAddressVerificationErrorResponseErrors) GetCandidates() []BillingAddressAttributes`

GetCandidates returns the Candidates field if non-nil, zero value otherwise.

### GetCandidatesOk

`func (o *BillingAddressVerificationErrorResponseErrors) GetCandidatesOk() (*[]BillingAddressAttributes, bool)`

GetCandidatesOk returns a tuple with the Candidates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCandidates

`func (o *BillingAddressVerificationErrorResponseErrors) SetCandidates(v []BillingAddressAttributes)`

SetCandidates sets Candidates field to given value.

### HasCandidates

`func (o *BillingAddressVerificationErrorResponseErrors) HasCandidates() bool`

HasCandidates returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
