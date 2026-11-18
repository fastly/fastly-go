# TLSSubscriptionResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ID** | Pointer to **string** |  | [optional] [readonly] 
**Attributes** | Pointer to [**TLSSubscriptionResponseAttributes**](TlsSubscriptionResponseAttributes.md) |  | [optional] 

## Methods

### NewTLSSubscriptionResponseData

`func NewTLSSubscriptionResponseData() *TLSSubscriptionResponseData`

NewTLSSubscriptionResponseData instantiates a new TLSSubscriptionResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTLSSubscriptionResponseDataWithDefaults

`func NewTLSSubscriptionResponseDataWithDefaults() *TLSSubscriptionResponseData`

NewTLSSubscriptionResponseDataWithDefaults instantiates a new TLSSubscriptionResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetID

`func (o *TLSSubscriptionResponseData) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *TLSSubscriptionResponseData) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *TLSSubscriptionResponseData) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *TLSSubscriptionResponseData) HasID() bool`

HasID returns a boolean if a field has been set.

### GetAttributes

`func (o *TLSSubscriptionResponseData) GetAttributes() TLSSubscriptionResponseAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TLSSubscriptionResponseData) GetAttributesOk() (*TLSSubscriptionResponseAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TLSSubscriptionResponseData) SetAttributes(v TLSSubscriptionResponseAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *TLSSubscriptionResponseData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
