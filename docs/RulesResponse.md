# RulesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]RuleResponse**](RuleResponse.md) |  | [optional] 
**Meta** | Pointer to [**PaginationMeta**](PaginationMeta.md) |  | [optional] 

## Methods

### NewRulesResponse

`func NewRulesResponse() *RulesResponse`

NewRulesResponse instantiates a new RulesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRulesResponseWithDefaults

`func NewRulesResponseWithDefaults() *RulesResponse`

NewRulesResponseWithDefaults instantiates a new RulesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *RulesResponse) GetData() []RuleResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RulesResponse) GetDataOk() (*[]RuleResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RulesResponse) SetData(v []RuleResponse)`

SetData sets Data field to given value.

### HasData

`func (o *RulesResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *RulesResponse) GetMeta() PaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *RulesResponse) GetMetaOk() (*PaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *RulesResponse) SetMeta(v PaginationMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *RulesResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


