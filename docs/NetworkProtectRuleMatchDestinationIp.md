# NetworkProtectRuleMatchDestinationIp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationIpAddress** | **string** | Match traffic sent to this address; omit for any destination. An address outside a Network using this Rule Set never matches and still consumes one of the 64 rules. The API does not enforce containment in a network prefix; attached Rule Sets require a compatible IP family. | 

## Methods

### NewNetworkProtectRuleMatchDestinationIp

`func NewNetworkProtectRuleMatchDestinationIp(destinationIpAddress string, ) *NetworkProtectRuleMatchDestinationIp`

NewNetworkProtectRuleMatchDestinationIp instantiates a new NetworkProtectRuleMatchDestinationIp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkProtectRuleMatchDestinationIpWithDefaults

`func NewNetworkProtectRuleMatchDestinationIpWithDefaults() *NetworkProtectRuleMatchDestinationIp`

NewNetworkProtectRuleMatchDestinationIpWithDefaults instantiates a new NetworkProtectRuleMatchDestinationIp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationIpAddress

`func (o *NetworkProtectRuleMatchDestinationIp) GetDestinationIpAddress() string`

GetDestinationIpAddress returns the DestinationIpAddress field if non-nil, zero value otherwise.

### GetDestinationIpAddressOk

`func (o *NetworkProtectRuleMatchDestinationIp) GetDestinationIpAddressOk() (*string, bool)`

GetDestinationIpAddressOk returns a tuple with the DestinationIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationIpAddress

`func (o *NetworkProtectRuleMatchDestinationIp) SetDestinationIpAddress(v string)`

SetDestinationIpAddress sets DestinationIpAddress field to given value.



[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


