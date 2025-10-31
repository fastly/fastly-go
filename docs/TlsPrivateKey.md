# TlsPrivateKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**TlsPrivateKeyData**](TlsPrivateKeyData.md) |  | [optional] 

## Methods

### NewTlsPrivateKey

`func NewTlsPrivateKey() *TlsPrivateKey`

NewTlsPrivateKey instantiates a new TlsPrivateKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTlsPrivateKeyWithDefaults

`func NewTlsPrivateKeyWithDefaults() *TlsPrivateKey`

NewTlsPrivateKeyWithDefaults instantiates a new TlsPrivateKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *TlsPrivateKey) GetData() TlsPrivateKeyData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TlsPrivateKey) GetDataOk() (*TlsPrivateKeyData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TlsPrivateKey) SetData(v TlsPrivateKeyData)`

SetData sets Data field to given value.

### HasData

`func (o *TlsPrivateKey) HasData() bool`

HasData returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


