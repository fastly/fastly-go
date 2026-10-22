# NetworkProtectRuleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpProtocol** | [**NetworkProtectRuleIpProtocol**](NetworkProtectRuleIpProtocol.md) |  | 
**SourcePrefixListId** | **string** | Match traffic whose source appears in this prefix list. | 
**DestinationIpAddress** | **string** | Match traffic sent to this address; omit for any destination. An address outside a Network using this Rule Set never matches and still consumes one of the 64 rules. The API does not enforce containment in a network prefix; attached Rule Sets require a compatible IP family. | 
**SourcePorts** | [**NetworkProtectRuleSourcePorts**](NetworkProtectRuleSourcePorts.md) |  | 
**DestinationPorts** | [**NetworkProtectRuleDestinationPorts**](NetworkProtectRuleDestinationPorts.md) |  | 
**CreatedAt** | Pointer to **time.Time** | An ISO 8601 timestamp in UTC. | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** | An ISO 8601 timestamp in UTC. | [optional] [readonly] 

## Methods

### NewNetworkProtectRuleResponse

`func NewNetworkProtectRuleResponse(ipProtocol NetworkProtectRuleIpProtocol, sourcePrefixListId string, destinationIpAddress string, sourcePorts NetworkProtectRuleSourcePorts, destinationPorts NetworkProtectRuleDestinationPorts, ) *NetworkProtectRuleResponse`

NewNetworkProtectRuleResponse instantiates a new NetworkProtectRuleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkProtectRuleResponseWithDefaults

`func NewNetworkProtectRuleResponseWithDefaults() *NetworkProtectRuleResponse`

NewNetworkProtectRuleResponseWithDefaults instantiates a new NetworkProtectRuleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpProtocol

`func (o *NetworkProtectRuleResponse) GetIpProtocol() NetworkProtectRuleIpProtocol`

GetIpProtocol returns the IpProtocol field if non-nil, zero value otherwise.

### GetIpProtocolOk

`func (o *NetworkProtectRuleResponse) GetIpProtocolOk() (*NetworkProtectRuleIpProtocol, bool)`

GetIpProtocolOk returns a tuple with the IpProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpProtocol

`func (o *NetworkProtectRuleResponse) SetIpProtocol(v NetworkProtectRuleIpProtocol)`

SetIpProtocol sets IpProtocol field to given value.


### GetSourcePrefixListId

`func (o *NetworkProtectRuleResponse) GetSourcePrefixListId() string`

GetSourcePrefixListId returns the SourcePrefixListId field if non-nil, zero value otherwise.

### GetSourcePrefixListIdOk

`func (o *NetworkProtectRuleResponse) GetSourcePrefixListIdOk() (*string, bool)`

GetSourcePrefixListIdOk returns a tuple with the SourcePrefixListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePrefixListId

`func (o *NetworkProtectRuleResponse) SetSourcePrefixListId(v string)`

SetSourcePrefixListId sets SourcePrefixListId field to given value.


### GetDestinationIpAddress

`func (o *NetworkProtectRuleResponse) GetDestinationIpAddress() string`

GetDestinationIpAddress returns the DestinationIpAddress field if non-nil, zero value otherwise.

### GetDestinationIpAddressOk

`func (o *NetworkProtectRuleResponse) GetDestinationIpAddressOk() (*string, bool)`

GetDestinationIpAddressOk returns a tuple with the DestinationIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationIpAddress

`func (o *NetworkProtectRuleResponse) SetDestinationIpAddress(v string)`

SetDestinationIpAddress sets DestinationIpAddress field to given value.


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


