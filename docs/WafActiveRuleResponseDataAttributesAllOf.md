# WafActiveRuleResponseDataAttributesAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LatestRevision** | Pointer to **int32** | The latest rule revision number that is available for the associated rule revision. | [optional] [readonly] 
**Outdated** | Pointer to **bool** | Indicates if the associated rule revision is up to date or not. | [optional] [readonly] 

## Methods

### NewWafActiveRuleResponseDataAttributesAllOf

`func NewWafActiveRuleResponseDataAttributesAllOf() *WafActiveRuleResponseDataAttributesAllOf`

NewWafActiveRuleResponseDataAttributesAllOf instantiates a new WafActiveRuleResponseDataAttributesAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafActiveRuleResponseDataAttributesAllOfWithDefaults

`func NewWafActiveRuleResponseDataAttributesAllOfWithDefaults() *WafActiveRuleResponseDataAttributesAllOf`

NewWafActiveRuleResponseDataAttributesAllOfWithDefaults instantiates a new WafActiveRuleResponseDataAttributesAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLatestRevision

`func (o *WafActiveRuleResponseDataAttributesAllOf) GetLatestRevision() int32`

GetLatestRevision returns the LatestRevision field if non-nil, zero value otherwise.

### GetLatestRevisionOk

`func (o *WafActiveRuleResponseDataAttributesAllOf) GetLatestRevisionOk() (*int32, bool)`

GetLatestRevisionOk returns a tuple with the LatestRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestRevision

`func (o *WafActiveRuleResponseDataAttributesAllOf) SetLatestRevision(v int32)`

SetLatestRevision sets LatestRevision field to given value.

### HasLatestRevision

`func (o *WafActiveRuleResponseDataAttributesAllOf) HasLatestRevision() bool`

HasLatestRevision returns a boolean if a field has been set.

### GetOutdated

`func (o *WafActiveRuleResponseDataAttributesAllOf) GetOutdated() bool`

GetOutdated returns the Outdated field if non-nil, zero value otherwise.

### GetOutdatedOk

`func (o *WafActiveRuleResponseDataAttributesAllOf) GetOutdatedOk() (*bool, bool)`

GetOutdatedOk returns a tuple with the Outdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutdated

`func (o *WafActiveRuleResponseDataAttributesAllOf) SetOutdated(v bool)`

SetOutdated sets Outdated field to given value.

### HasOutdated

`func (o *WafActiveRuleResponseDataAttributesAllOf) HasOutdated() bool`

HasOutdated returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
