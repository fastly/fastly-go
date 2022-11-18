# RelationshipMemberWafFirewall

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**TypeWafFirewall**](TypeWafFirewall.md) |  | [optional] [default to TYPEWAFFIREWALL_WAF_FIREWALL]
**ID** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewRelationshipMemberWafFirewall

`func NewRelationshipMemberWafFirewall() *RelationshipMemberWafFirewall`

NewRelationshipMemberWafFirewall instantiates a new RelationshipMemberWafFirewall object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelationshipMemberWafFirewallWithDefaults

`func NewRelationshipMemberWafFirewallWithDefaults() *RelationshipMemberWafFirewall`

NewRelationshipMemberWafFirewallWithDefaults instantiates a new RelationshipMemberWafFirewall object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RelationshipMemberWafFirewall) GetType() TypeWafFirewall`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RelationshipMemberWafFirewall) GetTypeOk() (*TypeWafFirewall, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RelationshipMemberWafFirewall) SetType(v TypeWafFirewall)`

SetType sets Type field to given value.

### HasType

`func (o *RelationshipMemberWafFirewall) HasType() bool`

HasType returns a boolean if a field has been set.

### GetID

`func (o *RelationshipMemberWafFirewall) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *RelationshipMemberWafFirewall) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *RelationshipMemberWafFirewall) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *RelationshipMemberWafFirewall) HasID() bool`

HasID returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
