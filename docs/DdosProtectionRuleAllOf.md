# DdosProtectionRuleAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique ID of the rule. | [optional] 
**Name** | Pointer to **string** | A human-readable name for the rule. | [optional] 
**Action** | Pointer to **string** | Action types for a rule. Supported action values are default, block, log, off. The default action value follows the current protection mode of the associated service. | [optional] [default to "default"]
**CustomerId** | Pointer to **string** | Alphanumeric string identifying the customer. | [optional] 
**ServiceId** | Pointer to **string** | Alphanumeric string identifying the service. | [optional] 
**SourceIp** | Pointer to **NullableString** | Source IP address attribute. | [optional] 
**CountryCode** | Pointer to **NullableString** | Country code attribute. | [optional] 
**Host** | Pointer to **NullableString** | Host attribute. | [optional] 
**Asn** | Pointer to **NullableString** | ASN attribute. | [optional] 
**SourceIpPrefix** | Pointer to **NullableString** | Source IP prefix attribute. | [optional] 
**AdditionalAttributes** | Pointer to **[]string** | Attribute category for additional, unlisted attributes used in a rule. | [optional] 

## Methods

### NewDdosProtectionRuleAllOf

`func NewDdosProtectionRuleAllOf() *DdosProtectionRuleAllOf`

NewDdosProtectionRuleAllOf instantiates a new DdosProtectionRuleAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDdosProtectionRuleAllOfWithDefaults

`func NewDdosProtectionRuleAllOfWithDefaults() *DdosProtectionRuleAllOf`

NewDdosProtectionRuleAllOfWithDefaults instantiates a new DdosProtectionRuleAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DdosProtectionRuleAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DdosProtectionRuleAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DdosProtectionRuleAllOf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DdosProtectionRuleAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DdosProtectionRuleAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DdosProtectionRuleAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DdosProtectionRuleAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DdosProtectionRuleAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAction

`func (o *DdosProtectionRuleAllOf) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *DdosProtectionRuleAllOf) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *DdosProtectionRuleAllOf) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *DdosProtectionRuleAllOf) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetCustomerId

`func (o *DdosProtectionRuleAllOf) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *DdosProtectionRuleAllOf) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *DdosProtectionRuleAllOf) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *DdosProtectionRuleAllOf) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetServiceId

`func (o *DdosProtectionRuleAllOf) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *DdosProtectionRuleAllOf) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *DdosProtectionRuleAllOf) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *DdosProtectionRuleAllOf) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetSourceIp

`func (o *DdosProtectionRuleAllOf) GetSourceIp() string`

GetSourceIp returns the SourceIp field if non-nil, zero value otherwise.

### GetSourceIpOk

`func (o *DdosProtectionRuleAllOf) GetSourceIpOk() (*string, bool)`

GetSourceIpOk returns a tuple with the SourceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIp

`func (o *DdosProtectionRuleAllOf) SetSourceIp(v string)`

SetSourceIp sets SourceIp field to given value.

### HasSourceIp

`func (o *DdosProtectionRuleAllOf) HasSourceIp() bool`

HasSourceIp returns a boolean if a field has been set.

### SetSourceIpNil

`func (o *DdosProtectionRuleAllOf) SetSourceIpNil(b bool)`

 SetSourceIpNil sets the value for SourceIp to be an explicit nil

### UnsetSourceIp
`func (o *DdosProtectionRuleAllOf) UnsetSourceIp()`

UnsetSourceIp ensures that no value is present for SourceIp, not even an explicit nil
### GetCountryCode

`func (o *DdosProtectionRuleAllOf) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *DdosProtectionRuleAllOf) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *DdosProtectionRuleAllOf) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *DdosProtectionRuleAllOf) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### SetCountryCodeNil

`func (o *DdosProtectionRuleAllOf) SetCountryCodeNil(b bool)`

 SetCountryCodeNil sets the value for CountryCode to be an explicit nil

### UnsetCountryCode
`func (o *DdosProtectionRuleAllOf) UnsetCountryCode()`

UnsetCountryCode ensures that no value is present for CountryCode, not even an explicit nil
### GetHost

`func (o *DdosProtectionRuleAllOf) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *DdosProtectionRuleAllOf) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *DdosProtectionRuleAllOf) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *DdosProtectionRuleAllOf) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHostNil

`func (o *DdosProtectionRuleAllOf) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *DdosProtectionRuleAllOf) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetAsn

`func (o *DdosProtectionRuleAllOf) GetAsn() string`

GetAsn returns the Asn field if non-nil, zero value otherwise.

### GetAsnOk

`func (o *DdosProtectionRuleAllOf) GetAsnOk() (*string, bool)`

GetAsnOk returns a tuple with the Asn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsn

`func (o *DdosProtectionRuleAllOf) SetAsn(v string)`

SetAsn sets Asn field to given value.

### HasAsn

`func (o *DdosProtectionRuleAllOf) HasAsn() bool`

HasAsn returns a boolean if a field has been set.

### SetAsnNil

`func (o *DdosProtectionRuleAllOf) SetAsnNil(b bool)`

 SetAsnNil sets the value for Asn to be an explicit nil

### UnsetAsn
`func (o *DdosProtectionRuleAllOf) UnsetAsn()`

UnsetAsn ensures that no value is present for Asn, not even an explicit nil
### GetSourceIpPrefix

`func (o *DdosProtectionRuleAllOf) GetSourceIpPrefix() string`

GetSourceIpPrefix returns the SourceIpPrefix field if non-nil, zero value otherwise.

### GetSourceIpPrefixOk

`func (o *DdosProtectionRuleAllOf) GetSourceIpPrefixOk() (*string, bool)`

GetSourceIpPrefixOk returns a tuple with the SourceIpPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIpPrefix

`func (o *DdosProtectionRuleAllOf) SetSourceIpPrefix(v string)`

SetSourceIpPrefix sets SourceIpPrefix field to given value.

### HasSourceIpPrefix

`func (o *DdosProtectionRuleAllOf) HasSourceIpPrefix() bool`

HasSourceIpPrefix returns a boolean if a field has been set.

### SetSourceIpPrefixNil

`func (o *DdosProtectionRuleAllOf) SetSourceIpPrefixNil(b bool)`

 SetSourceIpPrefixNil sets the value for SourceIpPrefix to be an explicit nil

### UnsetSourceIpPrefix
`func (o *DdosProtectionRuleAllOf) UnsetSourceIpPrefix()`

UnsetSourceIpPrefix ensures that no value is present for SourceIpPrefix, not even an explicit nil
### GetAdditionalAttributes

`func (o *DdosProtectionRuleAllOf) GetAdditionalAttributes() []string`

GetAdditionalAttributes returns the AdditionalAttributes field if non-nil, zero value otherwise.

### GetAdditionalAttributesOk

`func (o *DdosProtectionRuleAllOf) GetAdditionalAttributesOk() (*[]string, bool)`

GetAdditionalAttributesOk returns a tuple with the AdditionalAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalAttributes

`func (o *DdosProtectionRuleAllOf) SetAdditionalAttributes(v []string)`

SetAdditionalAttributes sets AdditionalAttributes field to given value.

### HasAdditionalAttributes

`func (o *DdosProtectionRuleAllOf) HasAdditionalAttributes() bool`

HasAdditionalAttributes returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


