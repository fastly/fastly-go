# BillingEstimateResponseAllOfLine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlanNo** | Pointer to **int32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Units** | Pointer to **float32** |  | [optional] 
**PerUnitCost** | Pointer to **float32** |  | [optional] 
**ServiceNo** | Pointer to **float32** |  | [optional] 
**ServiceType** | Pointer to **string** |  | [optional] 
**Amount** | Pointer to **float32** |  | [optional] 
**ClientServiceID** | Pointer to **string** |  | [optional] 
**ClientPlanID** | Pointer to **string** |  | [optional] 

## Methods

### NewBillingEstimateResponseAllOfLine

`func NewBillingEstimateResponseAllOfLine() *BillingEstimateResponseAllOfLine`

NewBillingEstimateResponseAllOfLine instantiates a new BillingEstimateResponseAllOfLine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingEstimateResponseAllOfLineWithDefaults

`func NewBillingEstimateResponseAllOfLineWithDefaults() *BillingEstimateResponseAllOfLine`

NewBillingEstimateResponseAllOfLineWithDefaults instantiates a new BillingEstimateResponseAllOfLine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlanNo

`func (o *BillingEstimateResponseAllOfLine) GetPlanNo() int32`

GetPlanNo returns the PlanNo field if non-nil, zero value otherwise.

### GetPlanNoOk

`func (o *BillingEstimateResponseAllOfLine) GetPlanNoOk() (*int32, bool)`

GetPlanNoOk returns a tuple with the PlanNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanNo

`func (o *BillingEstimateResponseAllOfLine) SetPlanNo(v int32)`

SetPlanNo sets PlanNo field to given value.

### HasPlanNo

`func (o *BillingEstimateResponseAllOfLine) HasPlanNo() bool`

HasPlanNo returns a boolean if a field has been set.

### GetDescription

`func (o *BillingEstimateResponseAllOfLine) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BillingEstimateResponseAllOfLine) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BillingEstimateResponseAllOfLine) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BillingEstimateResponseAllOfLine) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUnits

`func (o *BillingEstimateResponseAllOfLine) GetUnits() float32`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *BillingEstimateResponseAllOfLine) GetUnitsOk() (*float32, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *BillingEstimateResponseAllOfLine) SetUnits(v float32)`

SetUnits sets Units field to given value.

### HasUnits

`func (o *BillingEstimateResponseAllOfLine) HasUnits() bool`

HasUnits returns a boolean if a field has been set.

### GetPerUnitCost

`func (o *BillingEstimateResponseAllOfLine) GetPerUnitCost() float32`

GetPerUnitCost returns the PerUnitCost field if non-nil, zero value otherwise.

### GetPerUnitCostOk

`func (o *BillingEstimateResponseAllOfLine) GetPerUnitCostOk() (*float32, bool)`

GetPerUnitCostOk returns a tuple with the PerUnitCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerUnitCost

`func (o *BillingEstimateResponseAllOfLine) SetPerUnitCost(v float32)`

SetPerUnitCost sets PerUnitCost field to given value.

### HasPerUnitCost

`func (o *BillingEstimateResponseAllOfLine) HasPerUnitCost() bool`

HasPerUnitCost returns a boolean if a field has been set.

### GetServiceNo

`func (o *BillingEstimateResponseAllOfLine) GetServiceNo() float32`

GetServiceNo returns the ServiceNo field if non-nil, zero value otherwise.

### GetServiceNoOk

`func (o *BillingEstimateResponseAllOfLine) GetServiceNoOk() (*float32, bool)`

GetServiceNoOk returns a tuple with the ServiceNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceNo

`func (o *BillingEstimateResponseAllOfLine) SetServiceNo(v float32)`

SetServiceNo sets ServiceNo field to given value.

### HasServiceNo

`func (o *BillingEstimateResponseAllOfLine) HasServiceNo() bool`

HasServiceNo returns a boolean if a field has been set.

### GetServiceType

`func (o *BillingEstimateResponseAllOfLine) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *BillingEstimateResponseAllOfLine) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *BillingEstimateResponseAllOfLine) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *BillingEstimateResponseAllOfLine) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### GetAmount

`func (o *BillingEstimateResponseAllOfLine) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *BillingEstimateResponseAllOfLine) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *BillingEstimateResponseAllOfLine) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *BillingEstimateResponseAllOfLine) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetClientServiceID

`func (o *BillingEstimateResponseAllOfLine) GetClientServiceID() string`

GetClientServiceID returns the ClientServiceID field if non-nil, zero value otherwise.

### GetClientServiceIDOk

`func (o *BillingEstimateResponseAllOfLine) GetClientServiceIDOk() (*string, bool)`

GetClientServiceIDOk returns a tuple with the ClientServiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientServiceID

`func (o *BillingEstimateResponseAllOfLine) SetClientServiceID(v string)`

SetClientServiceID sets ClientServiceID field to given value.

### HasClientServiceID

`func (o *BillingEstimateResponseAllOfLine) HasClientServiceID() bool`

HasClientServiceID returns a boolean if a field has been set.

### GetClientPlanID

`func (o *BillingEstimateResponseAllOfLine) GetClientPlanID() string`

GetClientPlanID returns the ClientPlanID field if non-nil, zero value otherwise.

### GetClientPlanIDOk

`func (o *BillingEstimateResponseAllOfLine) GetClientPlanIDOk() (*string, bool)`

GetClientPlanIDOk returns a tuple with the ClientPlanID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientPlanID

`func (o *BillingEstimateResponseAllOfLine) SetClientPlanID(v string)`

SetClientPlanID sets ClientPlanID field to given value.

### HasClientPlanID

`func (o *BillingEstimateResponseAllOfLine) HasClientPlanID() bool`

HasClientPlanID returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
