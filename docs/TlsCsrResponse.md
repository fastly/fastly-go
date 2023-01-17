# TLSCsrResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**TLSCsrResponseData**](TlsCsrResponseData.md) |  | [optional] 

## Methods

### NewTLSCsrResponse

`func NewTLSCsrResponse() *TLSCsrResponse`

NewTLSCsrResponse instantiates a new TLSCsrResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTLSCsrResponseWithDefaults

`func NewTLSCsrResponseWithDefaults() *TLSCsrResponse`

NewTLSCsrResponseWithDefaults instantiates a new TLSCsrResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *TLSCsrResponse) GetData() TLSCsrResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TLSCsrResponse) GetDataOk() (*TLSCsrResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TLSCsrResponse) SetData(v TLSCsrResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *TLSCsrResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
