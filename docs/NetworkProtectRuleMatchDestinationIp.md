# NetworkProtectRuleMatchDestinationIP

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationIPAddress** | **string** | Match traffic sent to this address; omit for any destination. An address outside a Network using this Rule Set never matches and still consumes one of the 64 rules. The API does not enforce containment in a network prefix; attached Rule Sets require a compatible IP family. | 

## Methods

### NewNetworkProtectRuleMatchDestinationIP

`func NewNetworkProtectRuleMatchDestinationIP(destinationIPAddress string, ) *NetworkProtectRuleMatchDestinationIP`

NewNetworkProtectRuleMatchDestinationIP instantiates a new NetworkProtectRuleMatchDestinationIP object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkProtectRuleMatchDestinationIPWithDefaults

`func NewNetworkProtectRuleMatchDestinationIPWithDefaults() *NetworkProtectRuleMatchDestinationIP`

NewNetworkProtectRuleMatchDestinationIPWithDefaults instantiates a new NetworkProtectRuleMatchDestinationIP object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationIPAddress

`func (o *NetworkProtectRuleMatchDestinationIP) GetDestinationIPAddress() string`

GetDestinationIPAddress returns the DestinationIPAddress field if non-nil, zero value otherwise.

### GetDestinationIPAddressOk

`func (o *NetworkProtectRuleMatchDestinationIP) GetDestinationIPAddressOk() (*string, bool)`

GetDestinationIPAddressOk returns a tuple with the DestinationIPAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationIPAddress

`func (o *NetworkProtectRuleMatchDestinationIP) SetDestinationIPAddress(v string)`

SetDestinationIPAddress sets DestinationIPAddress field to given value.



[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
