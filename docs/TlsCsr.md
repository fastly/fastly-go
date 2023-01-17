# TLSCsr

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**TLSCsrData**](TlsCsrData.md) |  | [optional] 

## Methods

### NewTLSCsr

`func NewTLSCsr() *TLSCsr`

NewTLSCsr instantiates a new TLSCsr object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTLSCsrWithDefaults

`func NewTLSCsrWithDefaults() *TLSCsr`

NewTLSCsrWithDefaults instantiates a new TLSCsr object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *TLSCsr) GetData() TLSCsrData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TLSCsr) GetDataOk() (*TLSCsrData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TLSCsr) SetData(v TLSCsrData)`

SetData sets Data field to given value.

### HasData

`func (o *TLSCsr) HasData() bool`

HasData returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
