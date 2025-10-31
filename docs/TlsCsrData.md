# TlsCsrData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**TypeTlsCsr**](TypeTlsCsr.md) |  | [optional] [default to TYPETLSCSR_CSR]
**Attributes** | Pointer to [**TlsCsrDataAttributes**](TlsCsrDataAttributes.md) |  | [optional] 

## Methods

### NewTlsCsrData

`func NewTlsCsrData() *TlsCsrData`

NewTlsCsrData instantiates a new TlsCsrData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTlsCsrDataWithDefaults

`func NewTlsCsrDataWithDefaults() *TlsCsrData`

NewTlsCsrDataWithDefaults instantiates a new TlsCsrData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TlsCsrData) GetType() TypeTlsCsr`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TlsCsrData) GetTypeOk() (*TypeTlsCsr, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TlsCsrData) SetType(v TypeTlsCsr)`

SetType sets Type field to given value.

### HasType

`func (o *TlsCsrData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAttributes

`func (o *TlsCsrData) GetAttributes() TlsCsrDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TlsCsrData) GetAttributesOk() (*TlsCsrDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TlsCsrData) SetAttributes(v TlsCsrDataAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *TlsCsrData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


