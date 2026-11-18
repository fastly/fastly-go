# TLSCertificate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**TLSCertificateData**](TlsCertificateData.md) |  | [optional] 

## Methods

### NewTLSCertificate

`func NewTLSCertificate() *TLSCertificate`

NewTLSCertificate instantiates a new TLSCertificate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTLSCertificateWithDefaults

`func NewTLSCertificateWithDefaults() *TLSCertificate`

NewTLSCertificateWithDefaults instantiates a new TLSCertificate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *TLSCertificate) GetData() TLSCertificateData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TLSCertificate) GetDataOk() (*TLSCertificateData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TLSCertificate) SetData(v TLSCertificateData)`

SetData sets Data field to given value.

### HasData

`func (o *TLSCertificate) HasData() bool`

HasData returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
