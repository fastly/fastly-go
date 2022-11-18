# Billing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndTime** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**StartTime** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**InvoiceID** | Pointer to **string** |  | [optional] [readonly] 
**CustomerID** | Pointer to **string** |  | [optional] [readonly] 
**VendorState** | Pointer to **string** | The current state of our third-party billing vendor. One of `up` or `down`. | [optional] [readonly] 
**Status** | Pointer to [**BillingStatus**](BillingStatus.md) |  | [optional] 
**Total** | Pointer to [**BillingTotal**](BillingTotal.md) |  | [optional] 
**Regions** | Pointer to **map[string]map[string]map[string]any** | Breakdown of regional data for products that are region based. | [optional] 

## Methods

### NewBilling

`func NewBilling() *Billing`

NewBilling instantiates a new Billing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingWithDefaults

`func NewBillingWithDefaults() *Billing`

NewBillingWithDefaults instantiates a new Billing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndTime

`func (o *Billing) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *Billing) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *Billing) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *Billing) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### SetEndTimeNil

`func (o *Billing) SetEndTimeNil(b bool)`

 SetEndTimeNil sets the value for EndTime to be an explicit nil

### UnsetEndTime
`func (o *Billing) UnsetEndTime()`

UnsetEndTime ensures that no value is present for EndTime, not even an explicit nil
### GetStartTime

`func (o *Billing) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *Billing) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *Billing) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *Billing) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTimeNil

`func (o *Billing) SetStartTimeNil(b bool)`

 SetStartTimeNil sets the value for StartTime to be an explicit nil

### UnsetStartTime
`func (o *Billing) UnsetStartTime()`

UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
### GetInvoiceID

`func (o *Billing) GetInvoiceID() string`

GetInvoiceID returns the InvoiceID field if non-nil, zero value otherwise.

### GetInvoiceIDOk

`func (o *Billing) GetInvoiceIDOk() (*string, bool)`

GetInvoiceIDOk returns a tuple with the InvoiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceID

`func (o *Billing) SetInvoiceID(v string)`

SetInvoiceID sets InvoiceID field to given value.

### HasInvoiceID

`func (o *Billing) HasInvoiceID() bool`

HasInvoiceID returns a boolean if a field has been set.

### GetCustomerID

`func (o *Billing) GetCustomerID() string`

GetCustomerID returns the CustomerID field if non-nil, zero value otherwise.

### GetCustomerIDOk

`func (o *Billing) GetCustomerIDOk() (*string, bool)`

GetCustomerIDOk returns a tuple with the CustomerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerID

`func (o *Billing) SetCustomerID(v string)`

SetCustomerID sets CustomerID field to given value.

### HasCustomerID

`func (o *Billing) HasCustomerID() bool`

HasCustomerID returns a boolean if a field has been set.

### GetVendorState

`func (o *Billing) GetVendorState() string`

GetVendorState returns the VendorState field if non-nil, zero value otherwise.

### GetVendorStateOk

`func (o *Billing) GetVendorStateOk() (*string, bool)`

GetVendorStateOk returns a tuple with the VendorState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorState

`func (o *Billing) SetVendorState(v string)`

SetVendorState sets VendorState field to given value.

### HasVendorState

`func (o *Billing) HasVendorState() bool`

HasVendorState returns a boolean if a field has been set.

### GetStatus

`func (o *Billing) GetStatus() BillingStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Billing) GetStatusOk() (*BillingStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Billing) SetStatus(v BillingStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Billing) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTotal

`func (o *Billing) GetTotal() BillingTotal`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *Billing) GetTotalOk() (*BillingTotal, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *Billing) SetTotal(v BillingTotal)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *Billing) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetRegions

`func (o *Billing) GetRegions() map[string]map[string]map[string]any`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *Billing) GetRegionsOk() (*map[string]map[string]map[string]any, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *Billing) SetRegions(v map[string]map[string]map[string]any)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *Billing) HasRegions() bool`

HasRegions returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
