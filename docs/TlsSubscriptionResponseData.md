# TlsSubscriptionResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Attributes** | Pointer to [**TlsSubscriptionResponseAttributes**](TlsSubscriptionResponseAttributes.md) |  | [optional] 

## Methods

### NewTlsSubscriptionResponseData

`func NewTlsSubscriptionResponseData() *TlsSubscriptionResponseData`

NewTlsSubscriptionResponseData instantiates a new TlsSubscriptionResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTlsSubscriptionResponseDataWithDefaults

`func NewTlsSubscriptionResponseDataWithDefaults() *TlsSubscriptionResponseData`

NewTlsSubscriptionResponseDataWithDefaults instantiates a new TlsSubscriptionResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TlsSubscriptionResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TlsSubscriptionResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TlsSubscriptionResponseData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TlsSubscriptionResponseData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAttributes

`func (o *TlsSubscriptionResponseData) GetAttributes() TlsSubscriptionResponseAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TlsSubscriptionResponseData) GetAttributesOk() (*TlsSubscriptionResponseAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TlsSubscriptionResponseData) SetAttributes(v TlsSubscriptionResponseAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *TlsSubscriptionResponseData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


