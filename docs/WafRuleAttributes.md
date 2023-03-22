# WafRuleAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ModsecRuleID** | Pointer to **int32** | Corresponding ModSecurity rule ID. | [optional] [readonly] 
**Publisher** | Pointer to **string** | Rule publisher. | [optional] [readonly] 
**Type** | Pointer to **string** | The rule&#39;s [type](https://docs.fastly.com/en/guides/managing-rules-on-the-fastly-waf#understanding-the-types-of-rules). | [optional] [readonly] 

## Methods

### NewWafRuleAttributes

`func NewWafRuleAttributes() *WafRuleAttributes`

NewWafRuleAttributes instantiates a new WafRuleAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafRuleAttributesWithDefaults

`func NewWafRuleAttributesWithDefaults() *WafRuleAttributes`

NewWafRuleAttributesWithDefaults instantiates a new WafRuleAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModsecRuleID

`func (o *WafRuleAttributes) GetModsecRuleID() int32`

GetModsecRuleID returns the ModsecRuleID field if non-nil, zero value otherwise.

### GetModsecRuleIDOk

`func (o *WafRuleAttributes) GetModsecRuleIDOk() (*int32, bool)`

GetModsecRuleIDOk returns a tuple with the ModsecRuleID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModsecRuleID

`func (o *WafRuleAttributes) SetModsecRuleID(v int32)`

SetModsecRuleID sets ModsecRuleID field to given value.

### HasModsecRuleID

`func (o *WafRuleAttributes) HasModsecRuleID() bool`

HasModsecRuleID returns a boolean if a field has been set.

### GetPublisher

`func (o *WafRuleAttributes) GetPublisher() string`

GetPublisher returns the Publisher field if non-nil, zero value otherwise.

### GetPublisherOk

`func (o *WafRuleAttributes) GetPublisherOk() (*string, bool)`

GetPublisherOk returns a tuple with the Publisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublisher

`func (o *WafRuleAttributes) SetPublisher(v string)`

SetPublisher sets Publisher field to given value.

### HasPublisher

`func (o *WafRuleAttributes) HasPublisher() bool`

HasPublisher returns a boolean if a field has been set.

### GetType

`func (o *WafRuleAttributes) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WafRuleAttributes) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WafRuleAttributes) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WafRuleAttributes) HasType() bool`

HasType returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
