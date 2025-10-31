# DdosProtectionRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**UpdatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
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

### NewDdosProtectionRule

`func NewDdosProtectionRule() *DdosProtectionRule`

NewDdosProtectionRule instantiates a new DdosProtectionRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDdosProtectionRuleWithDefaults

`func NewDdosProtectionRuleWithDefaults() *DdosProtectionRule`

NewDdosProtectionRuleWithDefaults instantiates a new DdosProtectionRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *DdosProtectionRule) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DdosProtectionRule) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DdosProtectionRule) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DdosProtectionRule) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *DdosProtectionRule) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *DdosProtectionRule) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *DdosProtectionRule) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DdosProtectionRule) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DdosProtectionRule) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DdosProtectionRule) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *DdosProtectionRule) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *DdosProtectionRule) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetId

`func (o *DdosProtectionRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DdosProtectionRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DdosProtectionRule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DdosProtectionRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DdosProtectionRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DdosProtectionRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DdosProtectionRule) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DdosProtectionRule) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAction

`func (o *DdosProtectionRule) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *DdosProtectionRule) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *DdosProtectionRule) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *DdosProtectionRule) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetCustomerId

`func (o *DdosProtectionRule) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *DdosProtectionRule) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *DdosProtectionRule) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *DdosProtectionRule) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetServiceId

`func (o *DdosProtectionRule) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *DdosProtectionRule) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *DdosProtectionRule) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *DdosProtectionRule) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetSourceIp

`func (o *DdosProtectionRule) GetSourceIp() string`

GetSourceIp returns the SourceIp field if non-nil, zero value otherwise.

### GetSourceIpOk

`func (o *DdosProtectionRule) GetSourceIpOk() (*string, bool)`

GetSourceIpOk returns a tuple with the SourceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIp

`func (o *DdosProtectionRule) SetSourceIp(v string)`

SetSourceIp sets SourceIp field to given value.

### HasSourceIp

`func (o *DdosProtectionRule) HasSourceIp() bool`

HasSourceIp returns a boolean if a field has been set.

### SetSourceIpNil

`func (o *DdosProtectionRule) SetSourceIpNil(b bool)`

 SetSourceIpNil sets the value for SourceIp to be an explicit nil

### UnsetSourceIp
`func (o *DdosProtectionRule) UnsetSourceIp()`

UnsetSourceIp ensures that no value is present for SourceIp, not even an explicit nil
### GetCountryCode

`func (o *DdosProtectionRule) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *DdosProtectionRule) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *DdosProtectionRule) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *DdosProtectionRule) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### SetCountryCodeNil

`func (o *DdosProtectionRule) SetCountryCodeNil(b bool)`

 SetCountryCodeNil sets the value for CountryCode to be an explicit nil

### UnsetCountryCode
`func (o *DdosProtectionRule) UnsetCountryCode()`

UnsetCountryCode ensures that no value is present for CountryCode, not even an explicit nil
### GetHost

`func (o *DdosProtectionRule) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *DdosProtectionRule) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *DdosProtectionRule) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *DdosProtectionRule) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHostNil

`func (o *DdosProtectionRule) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *DdosProtectionRule) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetAsn

`func (o *DdosProtectionRule) GetAsn() string`

GetAsn returns the Asn field if non-nil, zero value otherwise.

### GetAsnOk

`func (o *DdosProtectionRule) GetAsnOk() (*string, bool)`

GetAsnOk returns a tuple with the Asn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsn

`func (o *DdosProtectionRule) SetAsn(v string)`

SetAsn sets Asn field to given value.

### HasAsn

`func (o *DdosProtectionRule) HasAsn() bool`

HasAsn returns a boolean if a field has been set.

### SetAsnNil

`func (o *DdosProtectionRule) SetAsnNil(b bool)`

 SetAsnNil sets the value for Asn to be an explicit nil

### UnsetAsn
`func (o *DdosProtectionRule) UnsetAsn()`

UnsetAsn ensures that no value is present for Asn, not even an explicit nil
### GetSourceIpPrefix

`func (o *DdosProtectionRule) GetSourceIpPrefix() string`

GetSourceIpPrefix returns the SourceIpPrefix field if non-nil, zero value otherwise.

### GetSourceIpPrefixOk

`func (o *DdosProtectionRule) GetSourceIpPrefixOk() (*string, bool)`

GetSourceIpPrefixOk returns a tuple with the SourceIpPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIpPrefix

`func (o *DdosProtectionRule) SetSourceIpPrefix(v string)`

SetSourceIpPrefix sets SourceIpPrefix field to given value.

### HasSourceIpPrefix

`func (o *DdosProtectionRule) HasSourceIpPrefix() bool`

HasSourceIpPrefix returns a boolean if a field has been set.

### SetSourceIpPrefixNil

`func (o *DdosProtectionRule) SetSourceIpPrefixNil(b bool)`

 SetSourceIpPrefixNil sets the value for SourceIpPrefix to be an explicit nil

### UnsetSourceIpPrefix
`func (o *DdosProtectionRule) UnsetSourceIpPrefix()`

UnsetSourceIpPrefix ensures that no value is present for SourceIpPrefix, not even an explicit nil
### GetAdditionalAttributes

`func (o *DdosProtectionRule) GetAdditionalAttributes() []string`

GetAdditionalAttributes returns the AdditionalAttributes field if non-nil, zero value otherwise.

### GetAdditionalAttributesOk

`func (o *DdosProtectionRule) GetAdditionalAttributesOk() (*[]string, bool)`

GetAdditionalAttributesOk returns a tuple with the AdditionalAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalAttributes

`func (o *DdosProtectionRule) SetAdditionalAttributes(v []string)`

SetAdditionalAttributes sets AdditionalAttributes field to given value.

### HasAdditionalAttributes

`func (o *DdosProtectionRule) HasAdditionalAttributes() bool`

HasAdditionalAttributes returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


