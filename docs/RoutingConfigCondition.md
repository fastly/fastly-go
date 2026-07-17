# RoutingConfigCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ConditionType**](ConditionType.md) |  | 
**Operator** | [**ConditionOperator**](ConditionOperator.md) |  | 
**Key** | Pointer to **string** | The key to evaluate. For `header` conditions this is the header name. Required for `header` conditions. | [optional] 
**Value** | **string** | The value to compare against using the operator. | 

## Methods

### NewRoutingConfigCondition

`func NewRoutingConfigCondition(type_ ConditionType, operator ConditionOperator, value string, ) *RoutingConfigCondition`

NewRoutingConfigCondition instantiates a new RoutingConfigCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutingConfigConditionWithDefaults

`func NewRoutingConfigConditionWithDefaults() *RoutingConfigCondition`

NewRoutingConfigConditionWithDefaults instantiates a new RoutingConfigCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RoutingConfigCondition) GetType() ConditionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RoutingConfigCondition) GetTypeOk() (*ConditionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RoutingConfigCondition) SetType(v ConditionType)`

SetType sets Type field to given value.


### GetOperator

`func (o *RoutingConfigCondition) GetOperator() ConditionOperator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *RoutingConfigCondition) GetOperatorOk() (*ConditionOperator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *RoutingConfigCondition) SetOperator(v ConditionOperator)`

SetOperator sets Operator field to given value.


### GetKey

`func (o *RoutingConfigCondition) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *RoutingConfigCondition) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *RoutingConfigCondition) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *RoutingConfigCondition) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValue

`func (o *RoutingConfigCondition) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RoutingConfigCondition) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RoutingConfigCondition) SetValue(v string)`

SetValue sets Value field to given value.



[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


