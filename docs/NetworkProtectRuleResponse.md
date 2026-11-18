# NetworkProtectRuleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IPProtocol** | [**NetworkProtectRuleIPProtocol**](NetworkProtectRuleIPProtocol.md) |  | 
**SourcePrefixListID** | **string** | Match traffic whose source appears in this prefix list. | 
**DestinationIPAddress** | **string** | Match traffic sent to this address; omit for any destination. An address outside a Network using this Rule Set never matches and still consumes one of the 64 rules. The API does not enforce containment in a network prefix; attached Rule Sets require a compatible IP family. | 
**SourcePorts** | [**NetworkProtectRuleSourcePorts**](NetworkProtectRuleSourcePorts.md) |  | 
**DestinationPorts** | [**NetworkProtectRuleDestinationPorts**](NetworkProtectRuleDestinationPorts.md) |  | 
**CreatedAt** | Pointer to **time.Time** | An ISO 8601 timestamp in UTC. | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** | An ISO 8601 timestamp in UTC. | [optional] [readonly] 

## Methods

### NewNetworkProtectRuleResponse

`func NewNetworkProtectRuleResponse(ipProtocol NetworkProtectRuleIPProtocol, sourcePrefixListID string, destinationIPAddress string, sourcePorts NetworkProtectRuleSourcePorts, destinationPorts NetworkProtectRuleDestinationPorts, ) *NetworkProtectRuleResponse`

NewNetworkProtectRuleResponse instantiates a new NetworkProtectRuleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkProtectRuleResponseWithDefaults

`func NewNetworkProtectRuleResponseWithDefaults() *NetworkProtectRuleResponse`

NewNetworkProtectRuleResponseWithDefaults instantiates a new NetworkProtectRuleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIPProtocol

`func (o *NetworkProtectRuleResponse) GetIPProtocol() NetworkProtectRuleIPProtocol`

GetIPProtocol returns the IPProtocol field if non-nil, zero value otherwise.

### GetIPProtocolOk

`func (o *NetworkProtectRuleResponse) GetIPProtocolOk() (*NetworkProtectRuleIPProtocol, bool)`

GetIPProtocolOk returns a tuple with the IPProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIPProtocol

`func (o *NetworkProtectRuleResponse) SetIPProtocol(v NetworkProtectRuleIPProtocol)`

SetIPProtocol sets IPProtocol field to given value.


### GetSourcePrefixListID

`func (o *NetworkProtectRuleResponse) GetSourcePrefixListID() string`

GetSourcePrefixListID returns the SourcePrefixListID field if non-nil, zero value otherwise.

### GetSourcePrefixListIDOk

`func (o *NetworkProtectRuleResponse) GetSourcePrefixListIDOk() (*string, bool)`

GetSourcePrefixListIDOk returns a tuple with the SourcePrefixListID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePrefixListID

`func (o *NetworkProtectRuleResponse) SetSourcePrefixListID(v string)`

SetSourcePrefixListID sets SourcePrefixListID field to given value.


### GetDestinationIPAddress

`func (o *NetworkProtectRuleResponse) GetDestinationIPAddress() string`

GetDestinationIPAddress returns the DestinationIPAddress field if non-nil, zero value otherwise.

### GetDestinationIPAddressOk

`func (o *NetworkProtectRuleResponse) GetDestinationIPAddressOk() (*string, bool)`

GetDestinationIPAddressOk returns a tuple with the DestinationIPAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationIPAddress

`func (o *NetworkProtectRuleResponse) SetDestinationIPAddress(v string)`

SetDestinationIPAddress sets DestinationIPAddress field to given value.


### GetSourcePorts

`func (o *NetworkProtectRuleResponse) GetSourcePorts() NetworkProtectRuleSourcePorts`

GetSourcePorts returns the SourcePorts field if non-nil, zero value otherwise.

### GetSourcePortsOk

`func (o *NetworkProtectRuleResponse) GetSourcePortsOk() (*NetworkProtectRuleSourcePorts, bool)`

GetSourcePortsOk returns a tuple with the SourcePorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePorts

`func (o *NetworkProtectRuleResponse) SetSourcePorts(v NetworkProtectRuleSourcePorts)`

SetSourcePorts sets SourcePorts field to given value.


### GetDestinationPorts

`func (o *NetworkProtectRuleResponse) GetDestinationPorts() NetworkProtectRuleDestinationPorts`

GetDestinationPorts returns the DestinationPorts field if non-nil, zero value otherwise.

### GetDestinationPortsOk

`func (o *NetworkProtectRuleResponse) GetDestinationPortsOk() (*NetworkProtectRuleDestinationPorts, bool)`

GetDestinationPortsOk returns a tuple with the DestinationPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationPorts

`func (o *NetworkProtectRuleResponse) SetDestinationPorts(v NetworkProtectRuleDestinationPorts)`

SetDestinationPorts sets DestinationPorts field to given value.


### GetCreatedAt

`func (o *NetworkProtectRuleResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NetworkProtectRuleResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NetworkProtectRuleResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *NetworkProtectRuleResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *NetworkProtectRuleResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *NetworkProtectRuleResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *NetworkProtectRuleResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *NetworkProtectRuleResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
