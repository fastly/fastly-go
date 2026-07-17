# WafSimulateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WafResponse** | **int32** | The HTTP status code the WAF would return for the simulated request (e.g., `200` for allowed, `406` for blocked). | 
**Signals** | [**[]WafSimulateSignal**](WafSimulateSignal.md) | List of signals detected by the WAF during simulation. Empty array when no signals are detected. | 

## Methods

### NewWafSimulateResponse

`func NewWafSimulateResponse(wafResponse int32, signals []WafSimulateSignal, ) *WafSimulateResponse`

NewWafSimulateResponse instantiates a new WafSimulateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafSimulateResponseWithDefaults

`func NewWafSimulateResponseWithDefaults() *WafSimulateResponse`

NewWafSimulateResponseWithDefaults instantiates a new WafSimulateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWafResponse

`func (o *WafSimulateResponse) GetWafResponse() int32`

GetWafResponse returns the WafResponse field if non-nil, zero value otherwise.

### GetWafResponseOk

`func (o *WafSimulateResponse) GetWafResponseOk() (*int32, bool)`

GetWafResponseOk returns a tuple with the WafResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafResponse

`func (o *WafSimulateResponse) SetWafResponse(v int32)`

SetWafResponse sets WafResponse field to given value.


### GetSignals

`func (o *WafSimulateResponse) GetSignals() []WafSimulateSignal`

GetSignals returns the Signals field if non-nil, zero value otherwise.

### GetSignalsOk

`func (o *WafSimulateResponse) GetSignalsOk() (*[]WafSimulateSignal, bool)`

GetSignalsOk returns a tuple with the Signals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignals

`func (o *WafSimulateResponse) SetSignals(v []WafSimulateSignal)`

SetSignals sets Signals field to given value.



[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


