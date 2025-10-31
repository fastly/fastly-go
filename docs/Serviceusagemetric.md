# Serviceusagemetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | Pointer to **string** |  | [optional] [readonly] 
**ServiceId** | Pointer to **string** | Service ID associated with the usage. | [optional] 
**ServiceName** | Pointer to **string** | Name of the service associated with the usage. | [optional] 
**UsageUnits** | Pointer to **float32** | The quantity of the usage for the billing period. Amount will be in the units provided in the parent object (e.g., a quantity of `1.3` with a unit of `gb` would have a usage amount of 1.3 gigabytes). | [optional] 

## Methods

### NewServiceusagemetric

`func NewServiceusagemetric() *Serviceusagemetric`

NewServiceusagemetric instantiates a new Serviceusagemetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceusagemetricWithDefaults

`func NewServiceusagemetricWithDefaults() *Serviceusagemetric`

NewServiceusagemetricWithDefaults instantiates a new Serviceusagemetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerId

`func (o *Serviceusagemetric) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *Serviceusagemetric) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *Serviceusagemetric) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *Serviceusagemetric) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetServiceId

`func (o *Serviceusagemetric) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *Serviceusagemetric) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *Serviceusagemetric) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *Serviceusagemetric) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetServiceName

`func (o *Serviceusagemetric) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *Serviceusagemetric) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *Serviceusagemetric) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *Serviceusagemetric) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetUsageUnits

`func (o *Serviceusagemetric) GetUsageUnits() float32`

GetUsageUnits returns the UsageUnits field if non-nil, zero value otherwise.

### GetUsageUnitsOk

`func (o *Serviceusagemetric) GetUsageUnitsOk() (*float32, bool)`

GetUsageUnitsOk returns a tuple with the UsageUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageUnits

`func (o *Serviceusagemetric) SetUsageUnits(v float32)`

SetUsageUnits sets UsageUnits field to given value.

### HasUsageUnits

`func (o *Serviceusagemetric) HasUsageUnits() bool`

HasUsageUnits returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


