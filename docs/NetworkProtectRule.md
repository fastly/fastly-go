# NetworkProtectRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | Optional rule metadata of 1 to 100 Unicode code points. It does not participate in traffic matching. Omit it to leave a rule without a comment; replacing rules with an item that omits it clears the stored comment. Omitting rules on PATCH preserves existing comments. Empty strings and null are rejected. | [optional] 
**IpProtocol** | [**NetworkProtectRuleIpProtocol**](NetworkProtectRuleIpProtocol.md) |  | 
**SourcePrefixListId** | **string** | Match traffic whose source appears in this prefix list. | 
**DestinationIpAddress** | **string** | Match traffic sent to this address; omit for any destination. An address outside a Network using this Rule Set never matches and still consumes one of the 64 rules. The API does not enforce containment in a network prefix; attached Rule Sets require a compatible IP family. | 
**SourcePorts** | [**NetworkProtectRuleSourcePorts**](NetworkProtectRuleSourcePorts.md) |  | 
**DestinationPorts** | [**NetworkProtectRuleDestinationPorts**](NetworkProtectRuleDestinationPorts.md) |  | 
**Action** | [**NetworkProtectRuleAction**](NetworkProtectRuleAction.md) |  | 

## Methods

### NewNetworkProtectRule

`func NewNetworkProtectRule(ipProtocol NetworkProtectRuleIpProtocol, sourcePrefixListId string, destinationIpAddress string, sourcePorts NetworkProtectRuleSourcePorts, destinationPorts NetworkProtectRuleDestinationPorts, action NetworkProtectRuleAction, ) *NetworkProtectRule`

NewNetworkProtectRule instantiates a new NetworkProtectRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkProtectRuleWithDefaults

`func NewNetworkProtectRuleWithDefaults() *NetworkProtectRule`

NewNetworkProtectRuleWithDefaults instantiates a new NetworkProtectRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *NetworkProtectRule) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *NetworkProtectRule) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *NetworkProtectRule) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *NetworkProtectRule) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetIpProtocol

`func (o *NetworkProtectRule) GetIpProtocol() NetworkProtectRuleIpProtocol`

GetIpProtocol returns the IpProtocol field if non-nil, zero value otherwise.

### GetIpProtocolOk

`func (o *NetworkProtectRule) GetIpProtocolOk() (*NetworkProtectRuleIpProtocol, bool)`

GetIpProtocolOk returns a tuple with the IpProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpProtocol

`func (o *NetworkProtectRule) SetIpProtocol(v NetworkProtectRuleIpProtocol)`

SetIpProtocol sets IpProtocol field to given value.


### GetSourcePrefixListId

`func (o *NetworkProtectRule) GetSourcePrefixListId() string`

GetSourcePrefixListId returns the SourcePrefixListId field if non-nil, zero value otherwise.

### GetSourcePrefixListIdOk

`func (o *NetworkProtectRule) GetSourcePrefixListIdOk() (*string, bool)`

GetSourcePrefixListIdOk returns a tuple with the SourcePrefixListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePrefixListId

`func (o *NetworkProtectRule) SetSourcePrefixListId(v string)`

SetSourcePrefixListId sets SourcePrefixListId field to given value.


### GetDestinationIpAddress

`func (o *NetworkProtectRule) GetDestinationIpAddress() string`

GetDestinationIpAddress returns the DestinationIpAddress field if non-nil, zero value otherwise.

### GetDestinationIpAddressOk

`func (o *NetworkProtectRule) GetDestinationIpAddressOk() (*string, bool)`

GetDestinationIpAddressOk returns a tuple with the DestinationIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationIpAddress

`func (o *NetworkProtectRule) SetDestinationIpAddress(v string)`

SetDestinationIpAddress sets DestinationIpAddress field to given value.


### GetSourcePorts

`func (o *NetworkProtectRule) GetSourcePorts() NetworkProtectRuleSourcePorts`

GetSourcePorts returns the SourcePorts field if non-nil, zero value otherwise.

### GetSourcePortsOk

`func (o *NetworkProtectRule) GetSourcePortsOk() (*NetworkProtectRuleSourcePorts, bool)`

GetSourcePortsOk returns a tuple with the SourcePorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePorts

`func (o *NetworkProtectRule) SetSourcePorts(v NetworkProtectRuleSourcePorts)`

SetSourcePorts sets SourcePorts field to given value.


### GetDestinationPorts

`func (o *NetworkProtectRule) GetDestinationPorts() NetworkProtectRuleDestinationPorts`

GetDestinationPorts returns the DestinationPorts field if non-nil, zero value otherwise.

### GetDestinationPortsOk

`func (o *NetworkProtectRule) GetDestinationPortsOk() (*NetworkProtectRuleDestinationPorts, bool)`

GetDestinationPortsOk returns a tuple with the DestinationPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationPorts

`func (o *NetworkProtectRule) SetDestinationPorts(v NetworkProtectRuleDestinationPorts)`

SetDestinationPorts sets DestinationPorts field to given value.


### GetAction

`func (o *NetworkProtectRule) GetAction() NetworkProtectRuleAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *NetworkProtectRule) GetActionOk() (*NetworkProtectRuleAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *NetworkProtectRule) SetAction(v NetworkProtectRuleAction)`

SetAction sets Action field to given value.



[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


