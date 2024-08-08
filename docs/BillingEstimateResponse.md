# BillingEstimateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndTime** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**StartTime** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**CustomerID** | Pointer to **string** |  | [optional] [readonly] 
**VendorState** | Pointer to **string** | The current state of our third-party billing vendor. One of `up` or `down`. | [optional] [readonly] 
**Status** | Pointer to [**BillingStatus**](BillingStatus.md) |  | [optional] 
**Total** | Pointer to [**BillingTotal**](BillingTotal.md) |  | [optional] 
**Regions** | Pointer to [**map[string]BillingRegions**](BillingRegions.md) | Breakdown of regional data for products that are region based. | [optional] 
**InvoiceID** | Pointer to **string** |  | [optional] [readonly] 
**LineItems** | Pointer to [**[]BillingEstimateLinesLineItems**](BillingEstimateLinesLineItems.md) |  | [optional] 

## Methods

### NewBillingEstimateResponse

`func NewBillingEstimateResponse() *BillingEstimateResponse`

NewBillingEstimateResponse instantiates a new BillingEstimateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingEstimateResponseWithDefaults

`func NewBillingEstimateResponseWithDefaults() *BillingEstimateResponse`

NewBillingEstimateResponseWithDefaults instantiates a new BillingEstimateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndTime

`func (o *BillingEstimateResponse) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *BillingEstimateResponse) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *BillingEstimateResponse) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *BillingEstimateResponse) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### SetEndTimeNil

`func (o *BillingEstimateResponse) SetEndTimeNil(b bool)`

 SetEndTimeNil sets the value for EndTime to be an explicit nil

### UnsetEndTime
`func (o *BillingEstimateResponse) UnsetEndTime()`

UnsetEndTime ensures that no value is present for EndTime, not even an explicit nil
### GetStartTime

`func (o *BillingEstimateResponse) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *BillingEstimateResponse) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *BillingEstimateResponse) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *BillingEstimateResponse) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTimeNil

`func (o *BillingEstimateResponse) SetStartTimeNil(b bool)`

 SetStartTimeNil sets the value for StartTime to be an explicit nil

### UnsetStartTime
`func (o *BillingEstimateResponse) UnsetStartTime()`

UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
### GetCustomerID

`func (o *BillingEstimateResponse) GetCustomerID() string`

GetCustomerID returns the CustomerID field if non-nil, zero value otherwise.

### GetCustomerIDOk

`func (o *BillingEstimateResponse) GetCustomerIDOk() (*string, bool)`

GetCustomerIDOk returns a tuple with the CustomerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerID

`func (o *BillingEstimateResponse) SetCustomerID(v string)`

SetCustomerID sets CustomerID field to given value.

### HasCustomerID

`func (o *BillingEstimateResponse) HasCustomerID() bool`

HasCustomerID returns a boolean if a field has been set.

### GetVendorState

`func (o *BillingEstimateResponse) GetVendorState() string`

GetVendorState returns the VendorState field if non-nil, zero value otherwise.

### GetVendorStateOk

`func (o *BillingEstimateResponse) GetVendorStateOk() (*string, bool)`

GetVendorStateOk returns a tuple with the VendorState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorState

`func (o *BillingEstimateResponse) SetVendorState(v string)`

SetVendorState sets VendorState field to given value.

### HasVendorState

`func (o *BillingEstimateResponse) HasVendorState() bool`

HasVendorState returns a boolean if a field has been set.

### GetStatus

`func (o *BillingEstimateResponse) GetStatus() BillingStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BillingEstimateResponse) GetStatusOk() (*BillingStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BillingEstimateResponse) SetStatus(v BillingStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BillingEstimateResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTotal

`func (o *BillingEstimateResponse) GetTotal() BillingTotal`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *BillingEstimateResponse) GetTotalOk() (*BillingTotal, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *BillingEstimateResponse) SetTotal(v BillingTotal)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *BillingEstimateResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetRegions

`func (o *BillingEstimateResponse) GetRegions() map[string]BillingRegions`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *BillingEstimateResponse) GetRegionsOk() (*map[string]BillingRegions, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *BillingEstimateResponse) SetRegions(v map[string]BillingRegions)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *BillingEstimateResponse) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### GetInvoiceID

`func (o *BillingEstimateResponse) GetInvoiceID() string`

GetInvoiceID returns the InvoiceID field if non-nil, zero value otherwise.

### GetInvoiceIDOk

`func (o *BillingEstimateResponse) GetInvoiceIDOk() (*string, bool)`

GetInvoiceIDOk returns a tuple with the InvoiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceID

`func (o *BillingEstimateResponse) SetInvoiceID(v string)`

SetInvoiceID sets InvoiceID field to given value.

### HasInvoiceID

`func (o *BillingEstimateResponse) HasInvoiceID() bool`

HasInvoiceID returns a boolean if a field has been set.

### GetLineItems

`func (o *BillingEstimateResponse) GetLineItems() []BillingEstimateLinesLineItems`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *BillingEstimateResponse) GetLineItemsOk() (*[]BillingEstimateLinesLineItems, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *BillingEstimateResponse) SetLineItems(v []BillingEstimateLinesLineItems)`

SetLineItems sets LineItems field to given value.

### HasLineItems

`func (o *BillingEstimateResponse) HasLineItems() bool`

HasLineItems returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
