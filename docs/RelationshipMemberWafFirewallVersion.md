# RelationshipMemberWafFirewallVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**TypeWafFirewallVersion**](TypeWafFirewallVersion.md) |  | [optional] [default to TYPEWAFFIREWALLVERSION_WAF_FIREWALL_VERSION]
**ID** | Pointer to **string** | Alphanumeric string identifying a Firewall version. | [optional] [readonly] 

## Methods

### NewRelationshipMemberWafFirewallVersion

`func NewRelationshipMemberWafFirewallVersion() *RelationshipMemberWafFirewallVersion`

NewRelationshipMemberWafFirewallVersion instantiates a new RelationshipMemberWafFirewallVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelationshipMemberWafFirewallVersionWithDefaults

`func NewRelationshipMemberWafFirewallVersionWithDefaults() *RelationshipMemberWafFirewallVersion`

NewRelationshipMemberWafFirewallVersionWithDefaults instantiates a new RelationshipMemberWafFirewallVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RelationshipMemberWafFirewallVersion) GetType() TypeWafFirewallVersion`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RelationshipMemberWafFirewallVersion) GetTypeOk() (*TypeWafFirewallVersion, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RelationshipMemberWafFirewallVersion) SetType(v TypeWafFirewallVersion)`

SetType sets Type field to given value.

### HasType

`func (o *RelationshipMemberWafFirewallVersion) HasType() bool`

HasType returns a boolean if a field has been set.

### GetID

`func (o *RelationshipMemberWafFirewallVersion) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *RelationshipMemberWafFirewallVersion) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *RelationshipMemberWafFirewallVersion) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *RelationshipMemberWafFirewallVersion) HasID() bool`

HasID returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
