# InlineResponse20014

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]SuccessfulResponseAsObject**](SuccessfulResponseAsObject.md) |  | [optional] 
**Meta** | Pointer to **interface{}** | Meta for the pagination. | [optional] 

## Methods

### NewInlineResponse20014

`func NewInlineResponse20014() *InlineResponse20014`

NewInlineResponse20014 instantiates a new InlineResponse20014 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20014WithDefaults

`func NewInlineResponse20014WithDefaults() *InlineResponse20014`

NewInlineResponse20014WithDefaults instantiates a new InlineResponse20014 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *InlineResponse20014) GetData() []SuccessfulResponseAsObject`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InlineResponse20014) GetDataOk() (*[]SuccessfulResponseAsObject, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *InlineResponse20014) SetData(v []SuccessfulResponseAsObject)`

SetData sets Data field to given value.

### HasData

`func (o *InlineResponse20014) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *InlineResponse20014) GetMeta() interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *InlineResponse20014) GetMetaOk() (*interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *InlineResponse20014) SetMeta(v interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *InlineResponse20014) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### SetMetaNil

`func (o *InlineResponse20014) SetMetaNil(b bool)`

 SetMetaNil sets the value for Meta to be an explicit nil

### UnsetMeta
`func (o *InlineResponse20014) UnsetMeta()`

UnsetMeta ensures that no value is present for Meta, not even an explicit nil

[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


