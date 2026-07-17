# InlineResponse20019

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**AgentKeyMeta**](AgentKeyMeta.md) |  | [optional] 
**Data** | Pointer to [**[]AgentKey**](AgentKey.md) | The agent keys returned by the request. | [optional] 

## Methods

### NewInlineResponse20019

`func NewInlineResponse20019() *InlineResponse20019`

NewInlineResponse20019 instantiates a new InlineResponse20019 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20019WithDefaults

`func NewInlineResponse20019WithDefaults() *InlineResponse20019`

NewInlineResponse20019WithDefaults instantiates a new InlineResponse20019 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *InlineResponse20019) GetMeta() AgentKeyMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *InlineResponse20019) GetMetaOk() (*AgentKeyMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *InlineResponse20019) SetMeta(v AgentKeyMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *InlineResponse20019) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetData

`func (o *InlineResponse20019) GetData() []AgentKey`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InlineResponse20019) GetDataOk() (*[]AgentKey, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *InlineResponse20019) SetData(v []AgentKey)`

SetData sets Data field to given value.

### HasData

`func (o *InlineResponse20019) HasData() bool`

HasData returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


