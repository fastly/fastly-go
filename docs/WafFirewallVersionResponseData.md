# WafFirewallVersionResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**TypeWafFirewallVersion**](TypeWafFirewallVersion.md) |  | [optional] [default to TYPEWAFFIREWALLVERSION_WAF_FIREWALL_VERSION]
**Attributes** | Pointer to [**WafFirewallVersionResponseDataAttributes**](WafFirewallVersionResponseDataAttributes.md) |  | [optional] 
**ID** | Pointer to **string** | Alphanumeric string identifying a Firewall version. | [optional] [readonly] 
**Relationships** | Pointer to [**RelationshipsForWafFirewallVersion**](RelationshipsForWafFirewallVersion.md) |  | [optional] 

## Methods

### NewWafFirewallVersionResponseData

`func NewWafFirewallVersionResponseData() *WafFirewallVersionResponseData`

NewWafFirewallVersionResponseData instantiates a new WafFirewallVersionResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafFirewallVersionResponseDataWithDefaults

`func NewWafFirewallVersionResponseDataWithDefaults() *WafFirewallVersionResponseData`

NewWafFirewallVersionResponseDataWithDefaults instantiates a new WafFirewallVersionResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *WafFirewallVersionResponseData) GetType() TypeWafFirewallVersion`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WafFirewallVersionResponseData) GetTypeOk() (*TypeWafFirewallVersion, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WafFirewallVersionResponseData) SetType(v TypeWafFirewallVersion)`

SetType sets Type field to given value.

### HasType

`func (o *WafFirewallVersionResponseData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAttributes

`func (o *WafFirewallVersionResponseData) GetAttributes() WafFirewallVersionResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *WafFirewallVersionResponseData) GetAttributesOk() (*WafFirewallVersionResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *WafFirewallVersionResponseData) SetAttributes(v WafFirewallVersionResponseDataAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *WafFirewallVersionResponseData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetID

`func (o *WafFirewallVersionResponseData) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *WafFirewallVersionResponseData) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *WafFirewallVersionResponseData) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *WafFirewallVersionResponseData) HasID() bool`

HasID returns a boolean if a field has been set.

### GetRelationships

`func (o *WafFirewallVersionResponseData) GetRelationships() RelationshipsForWafFirewallVersion`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *WafFirewallVersionResponseData) GetRelationshipsOk() (*RelationshipsForWafFirewallVersion, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *WafFirewallVersionResponseData) SetRelationships(v RelationshipsForWafFirewallVersion)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *WafFirewallVersionResponseData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
