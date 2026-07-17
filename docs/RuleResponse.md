# RuleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**UpdatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**Id** | Pointer to **string** | Alphanumeric string identifying the rule. Stable across versions of the routing config. | [optional] [readonly] 
**IsDefault** | Pointer to **bool** | Whether this is the default (catch-all) rule for the path. | [optional] [readonly] 
**Action** | Pointer to [**Action**](Action.md) |  | [optional] 
**Conditions** | Pointer to [**[]RoutingConfigCondition**](RoutingConfigCondition.md) | The conditions a request must satisfy for this rule to match. Empty for the default rule. | [optional] 

## Methods

### NewRuleResponse

`func NewRuleResponse() *RuleResponse`

NewRuleResponse instantiates a new RuleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleResponseWithDefaults

`func NewRuleResponseWithDefaults() *RuleResponse`

NewRuleResponseWithDefaults instantiates a new RuleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *RuleResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RuleResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RuleResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RuleResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *RuleResponse) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *RuleResponse) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *RuleResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RuleResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RuleResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RuleResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *RuleResponse) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *RuleResponse) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetId

`func (o *RuleResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RuleResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RuleResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RuleResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsDefault

`func (o *RuleResponse) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *RuleResponse) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *RuleResponse) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *RuleResponse) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetAction

`func (o *RuleResponse) GetAction() Action`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *RuleResponse) GetActionOk() (*Action, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *RuleResponse) SetAction(v Action)`

SetAction sets Action field to given value.

### HasAction

`func (o *RuleResponse) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetConditions

`func (o *RuleResponse) GetConditions() []RoutingConfigCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *RuleResponse) GetConditionsOk() (*[]RoutingConfigCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *RuleResponse) SetConditions(v []RoutingConfigCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *RuleResponse) HasConditions() bool`

HasConditions returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


