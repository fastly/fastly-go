# NetworkProtectNetworkUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** | Optional descriptive text for the network. This metadata does not change traffic handling. Null clears it on update. | [optional] 
**Name** | Pointer to **string** | A human-readable name, unique within the customer. | [optional] 
**RuleSetID** | Pointer to **NullableString** | Custom rule-set attachment is unavailable. Omission, null, and an empty string preserve the current backend-assigned rule set. | [optional] 
**EndpointPolicies** | Pointer to [**[]NetworkProtectEndpointPolicy**](NetworkProtectEndpointPolicy.md) | Complete endpoint policy configuration. A network has at most one policy per region and may have no policies. | [optional] 
**Advertise** | Pointer to **bool** | Desired advertisement state, which is not the same as observed route state. Enabling it requires approved LOA, valid ROA, and a ready default-region delivery policy. A later ROA failure does not turn an existing advertisement off; losing LOA approval or the last ready default endpoint does. | [optional] [default to false]

## Methods

### NewNetworkProtectNetworkUpdate

`func NewNetworkProtectNetworkUpdate() *NetworkProtectNetworkUpdate`

NewNetworkProtectNetworkUpdate instantiates a new NetworkProtectNetworkUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkProtectNetworkUpdateWithDefaults

`func NewNetworkProtectNetworkUpdateWithDefaults() *NetworkProtectNetworkUpdate`

NewNetworkProtectNetworkUpdateWithDefaults instantiates a new NetworkProtectNetworkUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *NetworkProtectNetworkUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NetworkProtectNetworkUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NetworkProtectNetworkUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NetworkProtectNetworkUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *NetworkProtectNetworkUpdate) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *NetworkProtectNetworkUpdate) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetName

`func (o *NetworkProtectNetworkUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkProtectNetworkUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkProtectNetworkUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NetworkProtectNetworkUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRuleSetID

`func (o *NetworkProtectNetworkUpdate) GetRuleSetID() string`

GetRuleSetID returns the RuleSetID field if non-nil, zero value otherwise.

### GetRuleSetIDOk

`func (o *NetworkProtectNetworkUpdate) GetRuleSetIDOk() (*string, bool)`

GetRuleSetIDOk returns a tuple with the RuleSetID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleSetID

`func (o *NetworkProtectNetworkUpdate) SetRuleSetID(v string)`

SetRuleSetID sets RuleSetID field to given value.

### HasRuleSetID

`func (o *NetworkProtectNetworkUpdate) HasRuleSetID() bool`

HasRuleSetID returns a boolean if a field has been set.

### SetRuleSetIDNil

`func (o *NetworkProtectNetworkUpdate) SetRuleSetIDNil(b bool)`

 SetRuleSetIDNil sets the value for RuleSetID to be an explicit nil

### UnsetRuleSetID
`func (o *NetworkProtectNetworkUpdate) UnsetRuleSetID()`

UnsetRuleSetID ensures that no value is present for RuleSetID, not even an explicit nil
### GetEndpointPolicies

`func (o *NetworkProtectNetworkUpdate) GetEndpointPolicies() []NetworkProtectEndpointPolicy`

GetEndpointPolicies returns the EndpointPolicies field if non-nil, zero value otherwise.

### GetEndpointPoliciesOk

`func (o *NetworkProtectNetworkUpdate) GetEndpointPoliciesOk() (*[]NetworkProtectEndpointPolicy, bool)`

GetEndpointPoliciesOk returns a tuple with the EndpointPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointPolicies

`func (o *NetworkProtectNetworkUpdate) SetEndpointPolicies(v []NetworkProtectEndpointPolicy)`

SetEndpointPolicies sets EndpointPolicies field to given value.

### HasEndpointPolicies

`func (o *NetworkProtectNetworkUpdate) HasEndpointPolicies() bool`

HasEndpointPolicies returns a boolean if a field has been set.

### GetAdvertise

`func (o *NetworkProtectNetworkUpdate) GetAdvertise() bool`

GetAdvertise returns the Advertise field if non-nil, zero value otherwise.

### GetAdvertiseOk

`func (o *NetworkProtectNetworkUpdate) GetAdvertiseOk() (*bool, bool)`

GetAdvertiseOk returns a tuple with the Advertise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvertise

`func (o *NetworkProtectNetworkUpdate) SetAdvertise(v bool)`

SetAdvertise sets Advertise field to given value.

### HasAdvertise

`func (o *NetworkProtectNetworkUpdate) HasAdvertise() bool`

HasAdvertise returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
