# RuleResponseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Alphanumeric string identifying the rule. Stable across versions of the routing config. | [optional] [readonly] 
**IsDefault** | Pointer to **bool** | Whether this is the default (catch-all) rule for the path. | [optional] [readonly] 
**Action** | Pointer to [**Action**](Action.md) |  | [optional] 
**Conditions** | Pointer to [**[]RoutingConfigCondition**](RoutingConfigCondition.md) | The conditions a request must satisfy for this rule to match. Empty for the default rule. | [optional] 

## Methods

### NewRuleResponseAllOf

`func NewRuleResponseAllOf() *RuleResponseAllOf`

NewRuleResponseAllOf instantiates a new RuleResponseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleResponseAllOfWithDefaults

`func NewRuleResponseAllOfWithDefaults() *RuleResponseAllOf`

NewRuleResponseAllOfWithDefaults instantiates a new RuleResponseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RuleResponseAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RuleResponseAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RuleResponseAllOf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RuleResponseAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsDefault

`func (o *RuleResponseAllOf) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *RuleResponseAllOf) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *RuleResponseAllOf) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *RuleResponseAllOf) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetAction

`func (o *RuleResponseAllOf) GetAction() Action`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *RuleResponseAllOf) GetActionOk() (*Action, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *RuleResponseAllOf) SetAction(v Action)`

SetAction sets Action field to given value.

### HasAction

`func (o *RuleResponseAllOf) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetConditions

`func (o *RuleResponseAllOf) GetConditions() []RoutingConfigCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *RuleResponseAllOf) GetConditionsOk() (*[]RoutingConfigCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *RuleResponseAllOf) SetConditions(v []RoutingConfigCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *RuleResponseAllOf) HasConditions() bool`

HasConditions returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


