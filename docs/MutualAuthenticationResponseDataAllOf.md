# MutualAuthenticationResponseDataAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ID** | Pointer to **string** |  | [optional] [readonly] 
**Attributes** | Pointer to [**MutualAuthenticationResponseAttributes**](MutualAuthenticationResponseAttributes.md) |  | [optional] 

## Methods

### NewMutualAuthenticationResponseDataAllOf

`func NewMutualAuthenticationResponseDataAllOf() *MutualAuthenticationResponseDataAllOf`

NewMutualAuthenticationResponseDataAllOf instantiates a new MutualAuthenticationResponseDataAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMutualAuthenticationResponseDataAllOfWithDefaults

`func NewMutualAuthenticationResponseDataAllOfWithDefaults() *MutualAuthenticationResponseDataAllOf`

NewMutualAuthenticationResponseDataAllOfWithDefaults instantiates a new MutualAuthenticationResponseDataAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetID

`func (o *MutualAuthenticationResponseDataAllOf) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *MutualAuthenticationResponseDataAllOf) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *MutualAuthenticationResponseDataAllOf) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *MutualAuthenticationResponseDataAllOf) HasID() bool`

HasID returns a boolean if a field has been set.

### GetAttributes

`func (o *MutualAuthenticationResponseDataAllOf) GetAttributes() MutualAuthenticationResponseAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *MutualAuthenticationResponseDataAllOf) GetAttributesOk() (*MutualAuthenticationResponseAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *MutualAuthenticationResponseDataAllOf) SetAttributes(v MutualAuthenticationResponseAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *MutualAuthenticationResponseDataAllOf) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
