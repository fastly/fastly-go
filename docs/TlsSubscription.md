# TLSSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**TLSSubscriptionData**](TlsSubscriptionData.md) |  | [optional] 

## Methods

### NewTLSSubscription

`func NewTLSSubscription() *TLSSubscription`

NewTLSSubscription instantiates a new TLSSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTLSSubscriptionWithDefaults

`func NewTLSSubscriptionWithDefaults() *TLSSubscription`

NewTLSSubscriptionWithDefaults instantiates a new TLSSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *TLSSubscription) GetData() TLSSubscriptionData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TLSSubscription) GetDataOk() (*TLSSubscriptionData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TLSSubscription) SetData(v TLSSubscriptionData)`

SetData sets Data field to given value.

### HasData

`func (o *TLSSubscription) HasData() bool`

HasData returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
