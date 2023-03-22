# BillingTotalExtras

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of this extra cost. | [optional] 
**Recurring** | Pointer to **float32** | Recurring monthly cost in USD. Not present if $0.0. | [optional] 
**Setup** | Pointer to **float32** | Initial set up cost in USD. Not present if $0.0 or this is not the month the extra was added. | [optional] 

## Methods

### NewBillingTotalExtras

`func NewBillingTotalExtras() *BillingTotalExtras`

NewBillingTotalExtras instantiates a new BillingTotalExtras object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingTotalExtrasWithDefaults

`func NewBillingTotalExtrasWithDefaults() *BillingTotalExtras`

NewBillingTotalExtrasWithDefaults instantiates a new BillingTotalExtras object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BillingTotalExtras) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BillingTotalExtras) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BillingTotalExtras) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BillingTotalExtras) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRecurring

`func (o *BillingTotalExtras) GetRecurring() float32`

GetRecurring returns the Recurring field if non-nil, zero value otherwise.

### GetRecurringOk

`func (o *BillingTotalExtras) GetRecurringOk() (*float32, bool)`

GetRecurringOk returns a tuple with the Recurring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurring

`func (o *BillingTotalExtras) SetRecurring(v float32)`

SetRecurring sets Recurring field to given value.

### HasRecurring

`func (o *BillingTotalExtras) HasRecurring() bool`

HasRecurring returns a boolean if a field has been set.

### GetSetup

`func (o *BillingTotalExtras) GetSetup() float32`

GetSetup returns the Setup field if non-nil, zero value otherwise.

### GetSetupOk

`func (o *BillingTotalExtras) GetSetupOk() (*float32, bool)`

GetSetupOk returns a tuple with the Setup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetup

`func (o *BillingTotalExtras) SetSetup(v float32)`

SetSetup sets Setup field to given value.

### HasSetup

`func (o *BillingTotalExtras) HasSetup() bool`

HasSetup returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
