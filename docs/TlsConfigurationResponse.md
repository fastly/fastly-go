# TLSConfigurationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**TLSConfigurationResponseData**](TlsConfigurationResponseData.md) |  | [optional] 

## Methods

### NewTLSConfigurationResponse

`func NewTLSConfigurationResponse() *TLSConfigurationResponse`

NewTLSConfigurationResponse instantiates a new TLSConfigurationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTLSConfigurationResponseWithDefaults

`func NewTLSConfigurationResponseWithDefaults() *TLSConfigurationResponse`

NewTLSConfigurationResponseWithDefaults instantiates a new TLSConfigurationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *TLSConfigurationResponse) GetData() TLSConfigurationResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TLSConfigurationResponse) GetDataOk() (*TLSConfigurationResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TLSConfigurationResponse) SetData(v TLSConfigurationResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *TLSConfigurationResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
