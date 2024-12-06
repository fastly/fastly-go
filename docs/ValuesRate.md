# ValuesRate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rate** | Pointer to **float32** | The percentage of requests matching the value in the current dimension. | [optional] 

## Methods

### NewValuesRate

`func NewValuesRate() *ValuesRate`

NewValuesRate instantiates a new ValuesRate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValuesRateWithDefaults

`func NewValuesRateWithDefaults() *ValuesRate`

NewValuesRateWithDefaults instantiates a new ValuesRate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRate

`func (o *ValuesRate) GetRate() float32`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *ValuesRate) GetRateOk() (*float32, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *ValuesRate) SetRate(v float32)`

SetRate sets Rate field to given value.

### HasRate

`func (o *ValuesRate) HasRate() bool`

HasRate returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
