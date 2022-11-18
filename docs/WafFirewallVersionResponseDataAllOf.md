# WafFirewallVersionResponseDataAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ID** | Pointer to **string** | Alphanumeric string identifying a Firewall version. | [optional] [readonly] 
**Attributes** | Pointer to [**WafFirewallVersionResponseDataAttributes**](WafFirewallVersionResponseDataAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**RelationshipsForWafFirewallVersion**](RelationshipsForWafFirewallVersion.md) |  | [optional] 

## Methods

### NewWafFirewallVersionResponseDataAllOf

`func NewWafFirewallVersionResponseDataAllOf() *WafFirewallVersionResponseDataAllOf`

NewWafFirewallVersionResponseDataAllOf instantiates a new WafFirewallVersionResponseDataAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafFirewallVersionResponseDataAllOfWithDefaults

`func NewWafFirewallVersionResponseDataAllOfWithDefaults() *WafFirewallVersionResponseDataAllOf`

NewWafFirewallVersionResponseDataAllOfWithDefaults instantiates a new WafFirewallVersionResponseDataAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetID

`func (o *WafFirewallVersionResponseDataAllOf) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *WafFirewallVersionResponseDataAllOf) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *WafFirewallVersionResponseDataAllOf) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *WafFirewallVersionResponseDataAllOf) HasID() bool`

HasID returns a boolean if a field has been set.

### GetAttributes

`func (o *WafFirewallVersionResponseDataAllOf) GetAttributes() WafFirewallVersionResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *WafFirewallVersionResponseDataAllOf) GetAttributesOk() (*WafFirewallVersionResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *WafFirewallVersionResponseDataAllOf) SetAttributes(v WafFirewallVersionResponseDataAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *WafFirewallVersionResponseDataAllOf) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *WafFirewallVersionResponseDataAllOf) GetRelationships() RelationshipsForWafFirewallVersion`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *WafFirewallVersionResponseDataAllOf) GetRelationshipsOk() (*RelationshipsForWafFirewallVersion, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *WafFirewallVersionResponseDataAllOf) SetRelationships(v RelationshipsForWafFirewallVersion)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *WafFirewallVersionResponseDataAllOf) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
