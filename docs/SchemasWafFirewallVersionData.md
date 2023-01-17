# SchemasWafFirewallVersionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**TypeWafFirewallVersion**](TypeWafFirewallVersion.md) |  | [optional] [default to TYPEWAFFIREWALLVERSION_WAF_FIREWALL_VERSION]
**Attributes** | Pointer to [**WafFirewallVersionDataAttributes**](WafFirewallVersionDataAttributes.md) |  | [optional] 

## Methods

### NewSchemasWafFirewallVersionData

`func NewSchemasWafFirewallVersionData() *SchemasWafFirewallVersionData`

NewSchemasWafFirewallVersionData instantiates a new SchemasWafFirewallVersionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemasWafFirewallVersionDataWithDefaults

`func NewSchemasWafFirewallVersionDataWithDefaults() *SchemasWafFirewallVersionData`

NewSchemasWafFirewallVersionDataWithDefaults instantiates a new SchemasWafFirewallVersionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SchemasWafFirewallVersionData) GetType() TypeWafFirewallVersion`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SchemasWafFirewallVersionData) GetTypeOk() (*TypeWafFirewallVersion, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SchemasWafFirewallVersionData) SetType(v TypeWafFirewallVersion)`

SetType sets Type field to given value.

### HasType

`func (o *SchemasWafFirewallVersionData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAttributes

`func (o *SchemasWafFirewallVersionData) GetAttributes() WafFirewallVersionDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SchemasWafFirewallVersionData) GetAttributesOk() (*WafFirewallVersionDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SchemasWafFirewallVersionData) SetAttributes(v WafFirewallVersionDataAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *SchemasWafFirewallVersionData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
