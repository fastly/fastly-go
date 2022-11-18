# WafFirewallResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**TypeWafFirewall**](TypeWafFirewall.md) |  | [optional] [default to TYPEWAFFIREWALL_WAF_FIREWALL]
**Attributes** | Pointer to [**WafFirewallResponseDataAttributes**](WafFirewallResponseDataAttributes.md) |  | [optional] 
**ID** | Pointer to **string** |  | [optional] [readonly] 
**Relationships** | Pointer to [**RelationshipWafFirewallVersions**](RelationshipWafFirewallVersions.md) |  | [optional] 

## Methods

### NewWafFirewallResponseData

`func NewWafFirewallResponseData() *WafFirewallResponseData`

NewWafFirewallResponseData instantiates a new WafFirewallResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafFirewallResponseDataWithDefaults

`func NewWafFirewallResponseDataWithDefaults() *WafFirewallResponseData`

NewWafFirewallResponseDataWithDefaults instantiates a new WafFirewallResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *WafFirewallResponseData) GetType() TypeWafFirewall`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WafFirewallResponseData) GetTypeOk() (*TypeWafFirewall, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WafFirewallResponseData) SetType(v TypeWafFirewall)`

SetType sets Type field to given value.

### HasType

`func (o *WafFirewallResponseData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAttributes

`func (o *WafFirewallResponseData) GetAttributes() WafFirewallResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *WafFirewallResponseData) GetAttributesOk() (*WafFirewallResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *WafFirewallResponseData) SetAttributes(v WafFirewallResponseDataAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *WafFirewallResponseData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetID

`func (o *WafFirewallResponseData) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *WafFirewallResponseData) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *WafFirewallResponseData) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *WafFirewallResponseData) HasID() bool`

HasID returns a boolean if a field has been set.

### GetRelationships

`func (o *WafFirewallResponseData) GetRelationships() RelationshipWafFirewallVersions`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *WafFirewallResponseData) GetRelationshipsOk() (*RelationshipWafFirewallVersions, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *WafFirewallResponseData) SetRelationships(v RelationshipWafFirewallVersions)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *WafFirewallResponseData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
