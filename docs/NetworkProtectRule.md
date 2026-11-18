# NetworkProtectRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | Optional rule metadata of 1 to 100 Unicode code points. It does not participate in traffic matching. Omit it to leave a rule without a comment; replacing rules with an item that omits it clears the stored comment. Omitting rules on PATCH preserves existing comments. Empty strings and null are rejected. | [optional] 
**IPProtocol** | [**NetworkProtectRuleIPProtocol**](NetworkProtectRuleIPProtocol.md) |  | 
**SourcePrefixListID** | **string** | Match traffic whose source appears in this prefix list. | 
**DestinationIPAddress** | **string** | Match traffic sent to this address; omit for any destination. An address outside a Network using this Rule Set never matches and still consumes one of the 64 rules. The API does not enforce containment in a network prefix; attached Rule Sets require a compatible IP family. | 
**SourcePorts** | [**NetworkProtectRuleSourcePorts**](NetworkProtectRuleSourcePorts.md) |  | 
**DestinationPorts** | [**NetworkProtectRuleDestinationPorts**](NetworkProtectRuleDestinationPorts.md) |  | 
**Action** | [**NetworkProtectRuleAction**](NetworkProtectRuleAction.md) |  | 

## Methods

### NewNetworkProtectRule

`func NewNetworkProtectRule(ipProtocol NetworkProtectRuleIPProtocol, sourcePrefixListID string, destinationIPAddress string, sourcePorts NetworkProtectRuleSourcePorts, destinationPorts NetworkProtectRuleDestinationPorts, action NetworkProtectRuleAction, ) *NetworkProtectRule`

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

### GetIPProtocol

`func (o *NetworkProtectRule) GetIPProtocol() NetworkProtectRuleIPProtocol`

GetIPProtocol returns the IPProtocol field if non-nil, zero value otherwise.

### GetIPProtocolOk

`func (o *NetworkProtectRule) GetIPProtocolOk() (*NetworkProtectRuleIPProtocol, bool)`

GetIPProtocolOk returns a tuple with the IPProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIPProtocol

`func (o *NetworkProtectRule) SetIPProtocol(v NetworkProtectRuleIPProtocol)`

SetIPProtocol sets IPProtocol field to given value.


### GetSourcePrefixListID

`func (o *NetworkProtectRule) GetSourcePrefixListID() string`

GetSourcePrefixListID returns the SourcePrefixListID field if non-nil, zero value otherwise.

### GetSourcePrefixListIDOk

`func (o *NetworkProtectRule) GetSourcePrefixListIDOk() (*string, bool)`

GetSourcePrefixListIDOk returns a tuple with the SourcePrefixListID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePrefixListID

`func (o *NetworkProtectRule) SetSourcePrefixListID(v string)`

SetSourcePrefixListID sets SourcePrefixListID field to given value.


### GetDestinationIPAddress

`func (o *NetworkProtectRule) GetDestinationIPAddress() string`

GetDestinationIPAddress returns the DestinationIPAddress field if non-nil, zero value otherwise.

### GetDestinationIPAddressOk

`func (o *NetworkProtectRule) GetDestinationIPAddressOk() (*string, bool)`

GetDestinationIPAddressOk returns a tuple with the DestinationIPAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationIPAddress

`func (o *NetworkProtectRule) SetDestinationIPAddress(v string)`

SetDestinationIPAddress sets DestinationIPAddress field to given value.


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
