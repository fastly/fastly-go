# ValuesBandwidth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AverageBandwidthBytes** | Pointer to **float32** | The average bandwidth in bytes for responses to requests to the URL in the current dimension. | [optional] 
**BandwidthPercentage** | Pointer to **float32** | The total bandwidth percentage for all responses to requests to the URL in the current dimension. | [optional] 

## Methods

### NewValuesBandwidth

`func NewValuesBandwidth() *ValuesBandwidth`

NewValuesBandwidth instantiates a new ValuesBandwidth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValuesBandwidthWithDefaults

`func NewValuesBandwidthWithDefaults() *ValuesBandwidth`

NewValuesBandwidthWithDefaults instantiates a new ValuesBandwidth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAverageBandwidthBytes

`func (o *ValuesBandwidth) GetAverageBandwidthBytes() float32`

GetAverageBandwidthBytes returns the AverageBandwidthBytes field if non-nil, zero value otherwise.

### GetAverageBandwidthBytesOk

`func (o *ValuesBandwidth) GetAverageBandwidthBytesOk() (*float32, bool)`

GetAverageBandwidthBytesOk returns a tuple with the AverageBandwidthBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageBandwidthBytes

`func (o *ValuesBandwidth) SetAverageBandwidthBytes(v float32)`

SetAverageBandwidthBytes sets AverageBandwidthBytes field to given value.

### HasAverageBandwidthBytes

`func (o *ValuesBandwidth) HasAverageBandwidthBytes() bool`

HasAverageBandwidthBytes returns a boolean if a field has been set.

### GetBandwidthPercentage

`func (o *ValuesBandwidth) GetBandwidthPercentage() float32`

GetBandwidthPercentage returns the BandwidthPercentage field if non-nil, zero value otherwise.

### GetBandwidthPercentageOk

`func (o *ValuesBandwidth) GetBandwidthPercentageOk() (*float32, bool)`

GetBandwidthPercentageOk returns a tuple with the BandwidthPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthPercentage

`func (o *ValuesBandwidth) SetBandwidthPercentage(v float32)`

SetBandwidthPercentage sets BandwidthPercentage field to given value.

### HasBandwidthPercentage

`func (o *ValuesBandwidth) HasBandwidthPercentage() bool`

HasBandwidthPercentage returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


