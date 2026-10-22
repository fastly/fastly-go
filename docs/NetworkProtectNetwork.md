# NetworkProtectNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** | Optional descriptive text for the network. This metadata does not change traffic handling. Null clears it on update. | [optional] 
**Name** | **string** | A human-readable name, unique within the customer. | 
**Prefix** | **string** | The CIDR block to protect, masked to its network address. It cannot be changed after creation. | 
**RuleSetId** | Pointer to **NullableString** | The backend-assigned rule set applied to this network. | [optional] 
**EndpointPolicies** | [**[]NetworkProtectEndpointPolicy**](NetworkProtectEndpointPolicy.md) | Complete endpoint policy configuration. A network has at most one policy per region and may have no policies. | 
**Advertise** | **bool** | Desired advertisement state, which is not the same as observed route state. Enabling it requires approved LOA, valid ROA, and a ready default-region delivery policy. A later ROA failure does not turn an existing advertisement off; losing LOA approval or the last ready default endpoint does. | [default to false]

## Methods

### NewNetworkProtectNetwork

`func NewNetworkProtectNetwork(name string, prefix string, endpointPolicies []NetworkProtectEndpointPolicy, advertise bool, ) *NetworkProtectNetwork`

NewNetworkProtectNetwork instantiates a new NetworkProtectNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkProtectNetworkWithDefaults

`func NewNetworkProtectNetworkWithDefaults() *NetworkProtectNetwork`

NewNetworkProtectNetworkWithDefaults instantiates a new NetworkProtectNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *NetworkProtectNetwork) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NetworkProtectNetwork) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NetworkProtectNetwork) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NetworkProtectNetwork) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *NetworkProtectNetwork) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *NetworkProtectNetwork) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetName

`func (o *NetworkProtectNetwork) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkProtectNetwork) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkProtectNetwork) SetName(v string)`

SetName sets Name field to given value.


### GetPrefix

`func (o *NetworkProtectNetwork) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *NetworkProtectNetwork) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *NetworkProtectNetwork) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.


### GetRuleSetId

`func (o *NetworkProtectNetwork) GetRuleSetId() string`

GetRuleSetId returns the RuleSetId field if non-nil, zero value otherwise.

### GetRuleSetIdOk

`func (o *NetworkProtectNetwork) GetRuleSetIdOk() (*string, bool)`

GetRuleSetIdOk returns a tuple with the RuleSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleSetId

`func (o *NetworkProtectNetwork) SetRuleSetId(v string)`

SetRuleSetId sets RuleSetId field to given value.

### HasRuleSetId

`func (o *NetworkProtectNetwork) HasRuleSetId() bool`

HasRuleSetId returns a boolean if a field has been set.

### SetRuleSetIdNil

`func (o *NetworkProtectNetwork) SetRuleSetIdNil(b bool)`

 SetRuleSetIdNil sets the value for RuleSetId to be an explicit nil

### UnsetRuleSetId
`func (o *NetworkProtectNetwork) UnsetRuleSetId()`

UnsetRuleSetId ensures that no value is present for RuleSetId, not even an explicit nil
### GetEndpointPolicies

`func (o *NetworkProtectNetwork) GetEndpointPolicies() []NetworkProtectEndpointPolicy`

GetEndpointPolicies returns the EndpointPolicies field if non-nil, zero value otherwise.

### GetEndpointPoliciesOk

`func (o *NetworkProtectNetwork) GetEndpointPoliciesOk() (*[]NetworkProtectEndpointPolicy, bool)`

GetEndpointPoliciesOk returns a tuple with the EndpointPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointPolicies

`func (o *NetworkProtectNetwork) SetEndpointPolicies(v []NetworkProtectEndpointPolicy)`

SetEndpointPolicies sets EndpointPolicies field to given value.


### GetAdvertise

`func (o *NetworkProtectNetwork) GetAdvertise() bool`

GetAdvertise returns the Advertise field if non-nil, zero value otherwise.

### GetAdvertiseOk

`func (o *NetworkProtectNetwork) GetAdvertiseOk() (*bool, bool)`

GetAdvertiseOk returns a tuple with the Advertise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvertise

`func (o *NetworkProtectNetwork) SetAdvertise(v bool)`

SetAdvertise sets Advertise field to given value.



[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


