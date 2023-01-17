# BillingTotal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bandwidth** | Pointer to **float32** | The total amount of bandwidth used this month (See bandwidth_units for measurement). | [optional] 
**BandwidthCost** | Pointer to **float32** | The cost of the bandwidth used this month in USD. | [optional] 
**BandwidthUnits** | Pointer to **NullableString** | Bandwidth measurement units based on billing plan. | [optional] 
**Cost** | Pointer to **float32** | The final amount to be paid. | [optional] 
**CostBeforeDiscount** | Pointer to **float32** | Total incurred cost plus extras cost. | [optional] 
**Discount** | Pointer to **float32** | Calculated discount rate. | [optional] 
**Extras** | Pointer to [**[]BillingTotalExtras**](BillingTotalExtras.md) | A list of any extras for this invoice. | [optional] 
**ExtrasCost** | Pointer to **float32** | Total cost of all extras. | [optional] 
**IncurredCost** | Pointer to **float32** | The total cost of bandwidth and requests used this month. | [optional] 
**Overage** | Pointer to **float32** | How much over the plan minimum has been incurred. | [optional] 
**PlanCode** | Pointer to **string** | The short code the plan this invoice was generated under. | [optional] 
**PlanMinimum** | Pointer to **float32** | The minimum cost of this plan. | [optional] 
**PlanName** | Pointer to **string** | The name of the plan this invoice was generated under. | [optional] 
**Requests** | Pointer to **float32** | The total number of requests used this month. | [optional] 
**RequestsCost** | Pointer to **float32** | The cost of the requests used this month. | [optional] 
**Terms** | Pointer to **string** | Payment terms. Almost always Net15. | [optional] 

## Methods

### NewBillingTotal

`func NewBillingTotal() *BillingTotal`

NewBillingTotal instantiates a new BillingTotal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingTotalWithDefaults

`func NewBillingTotalWithDefaults() *BillingTotal`

NewBillingTotalWithDefaults instantiates a new BillingTotal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBandwidth

`func (o *BillingTotal) GetBandwidth() float32`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *BillingTotal) GetBandwidthOk() (*float32, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *BillingTotal) SetBandwidth(v float32)`

SetBandwidth sets Bandwidth field to given value.

### HasBandwidth

`func (o *BillingTotal) HasBandwidth() bool`

HasBandwidth returns a boolean if a field has been set.

### GetBandwidthCost

`func (o *BillingTotal) GetBandwidthCost() float32`

GetBandwidthCost returns the BandwidthCost field if non-nil, zero value otherwise.

### GetBandwidthCostOk

`func (o *BillingTotal) GetBandwidthCostOk() (*float32, bool)`

GetBandwidthCostOk returns a tuple with the BandwidthCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthCost

`func (o *BillingTotal) SetBandwidthCost(v float32)`

SetBandwidthCost sets BandwidthCost field to given value.

### HasBandwidthCost

`func (o *BillingTotal) HasBandwidthCost() bool`

HasBandwidthCost returns a boolean if a field has been set.

### GetBandwidthUnits

`func (o *BillingTotal) GetBandwidthUnits() string`

GetBandwidthUnits returns the BandwidthUnits field if non-nil, zero value otherwise.

### GetBandwidthUnitsOk

`func (o *BillingTotal) GetBandwidthUnitsOk() (*string, bool)`

GetBandwidthUnitsOk returns a tuple with the BandwidthUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthUnits

`func (o *BillingTotal) SetBandwidthUnits(v string)`

SetBandwidthUnits sets BandwidthUnits field to given value.

### HasBandwidthUnits

`func (o *BillingTotal) HasBandwidthUnits() bool`

HasBandwidthUnits returns a boolean if a field has been set.

### SetBandwidthUnitsNil

`func (o *BillingTotal) SetBandwidthUnitsNil(b bool)`

 SetBandwidthUnitsNil sets the value for BandwidthUnits to be an explicit nil

### UnsetBandwidthUnits
`func (o *BillingTotal) UnsetBandwidthUnits()`

UnsetBandwidthUnits ensures that no value is present for BandwidthUnits, not even an explicit nil
### GetCost

`func (o *BillingTotal) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *BillingTotal) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *BillingTotal) SetCost(v float32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *BillingTotal) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetCostBeforeDiscount

`func (o *BillingTotal) GetCostBeforeDiscount() float32`

GetCostBeforeDiscount returns the CostBeforeDiscount field if non-nil, zero value otherwise.

### GetCostBeforeDiscountOk

`func (o *BillingTotal) GetCostBeforeDiscountOk() (*float32, bool)`

GetCostBeforeDiscountOk returns a tuple with the CostBeforeDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostBeforeDiscount

`func (o *BillingTotal) SetCostBeforeDiscount(v float32)`

SetCostBeforeDiscount sets CostBeforeDiscount field to given value.

### HasCostBeforeDiscount

`func (o *BillingTotal) HasCostBeforeDiscount() bool`

HasCostBeforeDiscount returns a boolean if a field has been set.

### GetDiscount

`func (o *BillingTotal) GetDiscount() float32`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *BillingTotal) GetDiscountOk() (*float32, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *BillingTotal) SetDiscount(v float32)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *BillingTotal) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### GetExtras

`func (o *BillingTotal) GetExtras() []BillingTotalExtras`

GetExtras returns the Extras field if non-nil, zero value otherwise.

### GetExtrasOk

`func (o *BillingTotal) GetExtrasOk() (*[]BillingTotalExtras, bool)`

GetExtrasOk returns a tuple with the Extras field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtras

`func (o *BillingTotal) SetExtras(v []BillingTotalExtras)`

SetExtras sets Extras field to given value.

### HasExtras

`func (o *BillingTotal) HasExtras() bool`

HasExtras returns a boolean if a field has been set.

### GetExtrasCost

`func (o *BillingTotal) GetExtrasCost() float32`

GetExtrasCost returns the ExtrasCost field if non-nil, zero value otherwise.

### GetExtrasCostOk

`func (o *BillingTotal) GetExtrasCostOk() (*float32, bool)`

GetExtrasCostOk returns a tuple with the ExtrasCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtrasCost

`func (o *BillingTotal) SetExtrasCost(v float32)`

SetExtrasCost sets ExtrasCost field to given value.

### HasExtrasCost

`func (o *BillingTotal) HasExtrasCost() bool`

HasExtrasCost returns a boolean if a field has been set.

### GetIncurredCost

`func (o *BillingTotal) GetIncurredCost() float32`

GetIncurredCost returns the IncurredCost field if non-nil, zero value otherwise.

### GetIncurredCostOk

`func (o *BillingTotal) GetIncurredCostOk() (*float32, bool)`

GetIncurredCostOk returns a tuple with the IncurredCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncurredCost

`func (o *BillingTotal) SetIncurredCost(v float32)`

SetIncurredCost sets IncurredCost field to given value.

### HasIncurredCost

`func (o *BillingTotal) HasIncurredCost() bool`

HasIncurredCost returns a boolean if a field has been set.

### GetOverage

`func (o *BillingTotal) GetOverage() float32`

GetOverage returns the Overage field if non-nil, zero value otherwise.

### GetOverageOk

`func (o *BillingTotal) GetOverageOk() (*float32, bool)`

GetOverageOk returns a tuple with the Overage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverage

`func (o *BillingTotal) SetOverage(v float32)`

SetOverage sets Overage field to given value.

### HasOverage

`func (o *BillingTotal) HasOverage() bool`

HasOverage returns a boolean if a field has been set.

### GetPlanCode

`func (o *BillingTotal) GetPlanCode() string`

GetPlanCode returns the PlanCode field if non-nil, zero value otherwise.

### GetPlanCodeOk

`func (o *BillingTotal) GetPlanCodeOk() (*string, bool)`

GetPlanCodeOk returns a tuple with the PlanCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanCode

`func (o *BillingTotal) SetPlanCode(v string)`

SetPlanCode sets PlanCode field to given value.

### HasPlanCode

`func (o *BillingTotal) HasPlanCode() bool`

HasPlanCode returns a boolean if a field has been set.

### GetPlanMinimum

`func (o *BillingTotal) GetPlanMinimum() float32`

GetPlanMinimum returns the PlanMinimum field if non-nil, zero value otherwise.

### GetPlanMinimumOk

`func (o *BillingTotal) GetPlanMinimumOk() (*float32, bool)`

GetPlanMinimumOk returns a tuple with the PlanMinimum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanMinimum

`func (o *BillingTotal) SetPlanMinimum(v float32)`

SetPlanMinimum sets PlanMinimum field to given value.

### HasPlanMinimum

`func (o *BillingTotal) HasPlanMinimum() bool`

HasPlanMinimum returns a boolean if a field has been set.

### GetPlanName

`func (o *BillingTotal) GetPlanName() string`

GetPlanName returns the PlanName field if non-nil, zero value otherwise.

### GetPlanNameOk

`func (o *BillingTotal) GetPlanNameOk() (*string, bool)`

GetPlanNameOk returns a tuple with the PlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanName

`func (o *BillingTotal) SetPlanName(v string)`

SetPlanName sets PlanName field to given value.

### HasPlanName

`func (o *BillingTotal) HasPlanName() bool`

HasPlanName returns a boolean if a field has been set.

### GetRequests

`func (o *BillingTotal) GetRequests() float32`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *BillingTotal) GetRequestsOk() (*float32, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *BillingTotal) SetRequests(v float32)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *BillingTotal) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### GetRequestsCost

`func (o *BillingTotal) GetRequestsCost() float32`

GetRequestsCost returns the RequestsCost field if non-nil, zero value otherwise.

### GetRequestsCostOk

`func (o *BillingTotal) GetRequestsCostOk() (*float32, bool)`

GetRequestsCostOk returns a tuple with the RequestsCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestsCost

`func (o *BillingTotal) SetRequestsCost(v float32)`

SetRequestsCost sets RequestsCost field to given value.

### HasRequestsCost

`func (o *BillingTotal) HasRequestsCost() bool`

HasRequestsCost returns a boolean if a field has been set.

### GetTerms

`func (o *BillingTotal) GetTerms() string`

GetTerms returns the Terms field if non-nil, zero value otherwise.

### GetTermsOk

`func (o *BillingTotal) GetTermsOk() (*string, bool)`

GetTermsOk returns a tuple with the Terms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerms

`func (o *BillingTotal) SetTerms(v string)`

SetTerms sets Terms field to given value.

### HasTerms

`func (o *BillingTotal) HasTerms() bool`

HasTerms returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
