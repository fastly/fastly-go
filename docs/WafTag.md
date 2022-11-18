# WafTag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**TypeWafTag**](TypeWafTag.md) |  | [optional] [default to TYPEWAFTAG_WAF_TAG]
**ID** | Pointer to **string** | Alphanumeric string identifying a WAF tag. | [optional] [readonly] 
**Attributes** | Pointer to [**WafTagAttributes**](WafTagAttributes.md) |  | [optional] 

## Methods

### NewWafTag

`func NewWafTag() *WafTag`

NewWafTag instantiates a new WafTag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafTagWithDefaults

`func NewWafTagWithDefaults() *WafTag`

NewWafTagWithDefaults instantiates a new WafTag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *WafTag) GetType() TypeWafTag`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WafTag) GetTypeOk() (*TypeWafTag, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WafTag) SetType(v TypeWafTag)`

SetType sets Type field to given value.

### HasType

`func (o *WafTag) HasType() bool`

HasType returns a boolean if a field has been set.

### GetID

`func (o *WafTag) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *WafTag) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *WafTag) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *WafTag) HasID() bool`

HasID returns a boolean if a field has been set.

### GetAttributes

`func (o *WafTag) GetAttributes() WafTagAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *WafTag) GetAttributesOk() (*WafTagAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *WafTag) SetAttributes(v WafTagAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *WafTag) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
