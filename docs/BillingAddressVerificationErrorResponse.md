# BillingAddressVerificationErrorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to [**[]BillingAddressVerificationErrorResponseErrors**](BillingAddressVerificationErrorResponseErrors.md) |  | [optional] 

## Methods

### NewBillingAddressVerificationErrorResponse

`func NewBillingAddressVerificationErrorResponse() *BillingAddressVerificationErrorResponse`

NewBillingAddressVerificationErrorResponse instantiates a new BillingAddressVerificationErrorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingAddressVerificationErrorResponseWithDefaults

`func NewBillingAddressVerificationErrorResponseWithDefaults() *BillingAddressVerificationErrorResponse`

NewBillingAddressVerificationErrorResponseWithDefaults instantiates a new BillingAddressVerificationErrorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *BillingAddressVerificationErrorResponse) GetErrors() []BillingAddressVerificationErrorResponseErrors`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *BillingAddressVerificationErrorResponse) GetErrorsOk() (*[]BillingAddressVerificationErrorResponseErrors, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *BillingAddressVerificationErrorResponse) SetErrors(v []BillingAddressVerificationErrorResponseErrors)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *BillingAddressVerificationErrorResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
