# NetworkProtectNetworkCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** | Optional descriptive text for the network. This metadata does not change traffic handling. Null clears it on update. | [optional] 
**Name** | **string** | A human-readable name, unique within the customer. | 
**Prefix** | **string** | The CIDR block to protect, masked to its network address. It cannot be changed after creation. | 
**RuleSetId** | Pointer to **string** | Custom rule-set attachment is unavailable. Omit this value or send an empty string to use the backend-assigned default rule set. | [optional] 
**EndpointPolicies** | Pointer to [**[]NetworkProtectEndpointPolicy**](NetworkProtectEndpointPolicy.md) | Complete endpoint policy configuration. A network has at most one policy per region and may have no policies. | [optional] 

## Methods

### NewNetworkProtectNetworkCreate

`func NewNetworkProtectNetworkCreate(name string, prefix string, ) *NetworkProtectNetworkCreate`

NewNetworkProtectNetworkCreate instantiates a new NetworkProtectNetworkCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkProtectNetworkCreateWithDefaults

`func NewNetworkProtectNetworkCreateWithDefaults() *NetworkProtectNetworkCreate`

NewNetworkProtectNetworkCreateWithDefaults instantiates a new NetworkProtectNetworkCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *NetworkProtectNetworkCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NetworkProtectNetworkCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NetworkProtectNetworkCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NetworkProtectNetworkCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *NetworkProtectNetworkCreate) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *NetworkProtectNetworkCreate) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetName

`func (o *NetworkProtectNetworkCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkProtectNetworkCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkProtectNetworkCreate) SetName(v string)`

SetName sets Name field to given value.


### GetPrefix

`func (o *NetworkProtectNetworkCreate) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *NetworkProtectNetworkCreate) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *NetworkProtectNetworkCreate) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.


### GetRuleSetId

`func (o *NetworkProtectNetworkCreate) GetRuleSetId() string`

GetRuleSetId returns the RuleSetId field if non-nil, zero value otherwise.

### GetRuleSetIdOk

`func (o *NetworkProtectNetworkCreate) GetRuleSetIdOk() (*string, bool)`

GetRuleSetIdOk returns a tuple with the RuleSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleSetId

`func (o *NetworkProtectNetworkCreate) SetRuleSetId(v string)`

SetRuleSetId sets RuleSetId field to given value.

### HasRuleSetId

`func (o *NetworkProtectNetworkCreate) HasRuleSetId() bool`

HasRuleSetId returns a boolean if a field has been set.

### GetEndpointPolicies

`func (o *NetworkProtectNetworkCreate) GetEndpointPolicies() []NetworkProtectEndpointPolicy`

GetEndpointPolicies returns the EndpointPolicies field if non-nil, zero value otherwise.

### GetEndpointPoliciesOk

`func (o *NetworkProtectNetworkCreate) GetEndpointPoliciesOk() (*[]NetworkProtectEndpointPolicy, bool)`

GetEndpointPoliciesOk returns a tuple with the EndpointPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointPolicies

`func (o *NetworkProtectNetworkCreate) SetEndpointPolicies(v []NetworkProtectEndpointPolicy)`

SetEndpointPolicies sets EndpointPolicies field to given value.

### HasEndpointPolicies

`func (o *NetworkProtectNetworkCreate) HasEndpointPolicies() bool`

HasEndpointPolicies returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


