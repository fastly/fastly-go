# WafFirewallResponseDataAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ID** | Pointer to **string** |  | [optional] [readonly] 
**Attributes** | Pointer to [**WafFirewallResponseDataAttributes**](WafFirewallResponseDataAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**RelationshipWafFirewallVersions**](RelationshipWafFirewallVersions.md) |  | [optional] 

## Methods

### NewWafFirewallResponseDataAllOf

`func NewWafFirewallResponseDataAllOf() *WafFirewallResponseDataAllOf`

NewWafFirewallResponseDataAllOf instantiates a new WafFirewallResponseDataAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafFirewallResponseDataAllOfWithDefaults

`func NewWafFirewallResponseDataAllOfWithDefaults() *WafFirewallResponseDataAllOf`

NewWafFirewallResponseDataAllOfWithDefaults instantiates a new WafFirewallResponseDataAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetID

`func (o *WafFirewallResponseDataAllOf) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *WafFirewallResponseDataAllOf) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *WafFirewallResponseDataAllOf) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *WafFirewallResponseDataAllOf) HasID() bool`

HasID returns a boolean if a field has been set.

### GetAttributes

`func (o *WafFirewallResponseDataAllOf) GetAttributes() WafFirewallResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *WafFirewallResponseDataAllOf) GetAttributesOk() (*WafFirewallResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *WafFirewallResponseDataAllOf) SetAttributes(v WafFirewallResponseDataAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *WafFirewallResponseDataAllOf) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *WafFirewallResponseDataAllOf) GetRelationships() RelationshipWafFirewallVersions`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *WafFirewallResponseDataAllOf) GetRelationshipsOk() (*RelationshipWafFirewallVersions, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *WafFirewallResponseDataAllOf) SetRelationships(v RelationshipWafFirewallVersions)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *WafFirewallResponseDataAllOf) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
