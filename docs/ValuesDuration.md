# ValuesDuration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AverageResponseTime** | Pointer to **float32** | The average time in seconds to respond to requests to the URL in the current dimension. | [optional] 
**P95ResponseTime** | Pointer to **float32** | The P95 time in seconds to respond to requests to the URL in the current dimension. | [optional] 
**ResponseTimePercentage** | Pointer to **float32** | The total percentage of time to respond to all requests to the URL in the current dimension. | [optional] 

## Methods

### NewValuesDuration

`func NewValuesDuration() *ValuesDuration`

NewValuesDuration instantiates a new ValuesDuration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValuesDurationWithDefaults

`func NewValuesDurationWithDefaults() *ValuesDuration`

NewValuesDurationWithDefaults instantiates a new ValuesDuration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAverageResponseTime

`func (o *ValuesDuration) GetAverageResponseTime() float32`

GetAverageResponseTime returns the AverageResponseTime field if non-nil, zero value otherwise.

### GetAverageResponseTimeOk

`func (o *ValuesDuration) GetAverageResponseTimeOk() (*float32, bool)`

GetAverageResponseTimeOk returns a tuple with the AverageResponseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageResponseTime

`func (o *ValuesDuration) SetAverageResponseTime(v float32)`

SetAverageResponseTime sets AverageResponseTime field to given value.

### HasAverageResponseTime

`func (o *ValuesDuration) HasAverageResponseTime() bool`

HasAverageResponseTime returns a boolean if a field has been set.

### GetP95ResponseTime

`func (o *ValuesDuration) GetP95ResponseTime() float32`

GetP95ResponseTime returns the P95ResponseTime field if non-nil, zero value otherwise.

### GetP95ResponseTimeOk

`func (o *ValuesDuration) GetP95ResponseTimeOk() (*float32, bool)`

GetP95ResponseTimeOk returns a tuple with the P95ResponseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP95ResponseTime

`func (o *ValuesDuration) SetP95ResponseTime(v float32)`

SetP95ResponseTime sets P95ResponseTime field to given value.

### HasP95ResponseTime

`func (o *ValuesDuration) HasP95ResponseTime() bool`

HasP95ResponseTime returns a boolean if a field has been set.

### GetResponseTimePercentage

`func (o *ValuesDuration) GetResponseTimePercentage() float32`

GetResponseTimePercentage returns the ResponseTimePercentage field if non-nil, zero value otherwise.

### GetResponseTimePercentageOk

`func (o *ValuesDuration) GetResponseTimePercentageOk() (*float32, bool)`

GetResponseTimePercentageOk returns a tuple with the ResponseTimePercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTimePercentage

`func (o *ValuesDuration) SetResponseTimePercentage(v float32)`

SetResponseTimePercentage sets ResponseTimePercentage field to given value.

### HasResponseTimePercentage

`func (o *ValuesDuration) HasResponseTimePercentage() bool`

HasResponseTimePercentage returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


