# TLSActivationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**TLSActivationResponseData**](TlsActivationResponseData.md) |  | [optional] 

## Methods

### NewTLSActivationResponse

`func NewTLSActivationResponse() *TLSActivationResponse`

NewTLSActivationResponse instantiates a new TLSActivationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTLSActivationResponseWithDefaults

`func NewTLSActivationResponseWithDefaults() *TLSActivationResponse`

NewTLSActivationResponseWithDefaults instantiates a new TLSActivationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *TLSActivationResponse) GetData() TLSActivationResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TLSActivationResponse) GetDataOk() (*TLSActivationResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TLSActivationResponse) SetData(v TLSActivationResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *TLSActivationResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
