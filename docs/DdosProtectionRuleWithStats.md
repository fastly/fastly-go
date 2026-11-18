# DdosProtectionRuleWithStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**UpdatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**ID** | Pointer to **string** | Unique ID of the rule. | [optional] 
**Name** | Pointer to **string** | A human-readable name for the rule. | [optional] 
**Action** | Pointer to **string** | Action types for a rule. Supported action values are default, block, log, off, client_challenge. The default action value follows the current protection mode of the associated service. | [optional] [default to "default"]
**CustomerID** | Pointer to **string** | Alphanumeric string identifying the customer. | [optional] 
**ServiceID** | Pointer to **string** | Alphanumeric string identifying the service. | [optional] 
**SourceIP** | Pointer to **NullableString** | Source IP address attribute. | [optional] 
**CountryCode** | Pointer to **NullableString** | Country code attribute. | [optional] 
**Host** | Pointer to **NullableString** | Host attribute. | [optional] 
**Asn** | Pointer to **NullableString** | ASN attribute. | [optional] 
**SourceIPPrefix** | Pointer to **NullableString** | Source IP prefix attribute. | [optional] 
**Path** | Pointer to **NullableString** | URL path attribute. | [optional] 
**AdditionalAttributes** | Pointer to **[]string** | Attribute category for additional, unlisted attributes used in a rule. | [optional] 
**TrafficStats** | Pointer to [**DdosProtectionTrafficStats**](DdosProtectionTrafficStats.md) |  | [optional] 

## Methods

### NewDdosProtectionRuleWithStats

`func NewDdosProtectionRuleWithStats() *DdosProtectionRuleWithStats`

NewDdosProtectionRuleWithStats instantiates a new DdosProtectionRuleWithStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDdosProtectionRuleWithStatsWithDefaults

`func NewDdosProtectionRuleWithStatsWithDefaults() *DdosProtectionRuleWithStats`

NewDdosProtectionRuleWithStatsWithDefaults instantiates a new DdosProtectionRuleWithStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *DdosProtectionRuleWithStats) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DdosProtectionRuleWithStats) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DdosProtectionRuleWithStats) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DdosProtectionRuleWithStats) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *DdosProtectionRuleWithStats) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *DdosProtectionRuleWithStats) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *DdosProtectionRuleWithStats) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DdosProtectionRuleWithStats) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DdosProtectionRuleWithStats) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DdosProtectionRuleWithStats) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *DdosProtectionRuleWithStats) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *DdosProtectionRuleWithStats) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetID

`func (o *DdosProtectionRuleWithStats) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *DdosProtectionRuleWithStats) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *DdosProtectionRuleWithStats) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *DdosProtectionRuleWithStats) HasID() bool`

HasID returns a boolean if a field has been set.

### GetName

`func (o *DdosProtectionRuleWithStats) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DdosProtectionRuleWithStats) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DdosProtectionRuleWithStats) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DdosProtectionRuleWithStats) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAction

`func (o *DdosProtectionRuleWithStats) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *DdosProtectionRuleWithStats) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *DdosProtectionRuleWithStats) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *DdosProtectionRuleWithStats) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetCustomerID

`func (o *DdosProtectionRuleWithStats) GetCustomerID() string`

GetCustomerID returns the CustomerID field if non-nil, zero value otherwise.

### GetCustomerIDOk

`func (o *DdosProtectionRuleWithStats) GetCustomerIDOk() (*string, bool)`

GetCustomerIDOk returns a tuple with the CustomerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerID

`func (o *DdosProtectionRuleWithStats) SetCustomerID(v string)`

SetCustomerID sets CustomerID field to given value.

### HasCustomerID

`func (o *DdosProtectionRuleWithStats) HasCustomerID() bool`

HasCustomerID returns a boolean if a field has been set.

### GetServiceID

`func (o *DdosProtectionRuleWithStats) GetServiceID() string`

GetServiceID returns the ServiceID field if non-nil, zero value otherwise.

### GetServiceIDOk

`func (o *DdosProtectionRuleWithStats) GetServiceIDOk() (*string, bool)`

GetServiceIDOk returns a tuple with the ServiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceID

`func (o *DdosProtectionRuleWithStats) SetServiceID(v string)`

SetServiceID sets ServiceID field to given value.

### HasServiceID

`func (o *DdosProtectionRuleWithStats) HasServiceID() bool`

HasServiceID returns a boolean if a field has been set.

### GetSourceIP

`func (o *DdosProtectionRuleWithStats) GetSourceIP() string`

GetSourceIP returns the SourceIP field if non-nil, zero value otherwise.

### GetSourceIPOk

`func (o *DdosProtectionRuleWithStats) GetSourceIPOk() (*string, bool)`

GetSourceIPOk returns a tuple with the SourceIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIP

`func (o *DdosProtectionRuleWithStats) SetSourceIP(v string)`

SetSourceIP sets SourceIP field to given value.

### HasSourceIP

`func (o *DdosProtectionRuleWithStats) HasSourceIP() bool`

HasSourceIP returns a boolean if a field has been set.

### SetSourceIPNil

`func (o *DdosProtectionRuleWithStats) SetSourceIPNil(b bool)`

 SetSourceIPNil sets the value for SourceIP to be an explicit nil

### UnsetSourceIP
`func (o *DdosProtectionRuleWithStats) UnsetSourceIP()`

UnsetSourceIP ensures that no value is present for SourceIP, not even an explicit nil
### GetCountryCode

`func (o *DdosProtectionRuleWithStats) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *DdosProtectionRuleWithStats) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *DdosProtectionRuleWithStats) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *DdosProtectionRuleWithStats) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### SetCountryCodeNil

`func (o *DdosProtectionRuleWithStats) SetCountryCodeNil(b bool)`

 SetCountryCodeNil sets the value for CountryCode to be an explicit nil

### UnsetCountryCode
`func (o *DdosProtectionRuleWithStats) UnsetCountryCode()`

UnsetCountryCode ensures that no value is present for CountryCode, not even an explicit nil
### GetHost

`func (o *DdosProtectionRuleWithStats) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *DdosProtectionRuleWithStats) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *DdosProtectionRuleWithStats) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *DdosProtectionRuleWithStats) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHostNil

`func (o *DdosProtectionRuleWithStats) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *DdosProtectionRuleWithStats) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetAsn

`func (o *DdosProtectionRuleWithStats) GetAsn() string`

GetAsn returns the Asn field if non-nil, zero value otherwise.

### GetAsnOk

`func (o *DdosProtectionRuleWithStats) GetAsnOk() (*string, bool)`

GetAsnOk returns a tuple with the Asn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsn

`func (o *DdosProtectionRuleWithStats) SetAsn(v string)`

SetAsn sets Asn field to given value.

### HasAsn

`func (o *DdosProtectionRuleWithStats) HasAsn() bool`

HasAsn returns a boolean if a field has been set.

### SetAsnNil

`func (o *DdosProtectionRuleWithStats) SetAsnNil(b bool)`

 SetAsnNil sets the value for Asn to be an explicit nil

### UnsetAsn
`func (o *DdosProtectionRuleWithStats) UnsetAsn()`

UnsetAsn ensures that no value is present for Asn, not even an explicit nil
### GetSourceIPPrefix

`func (o *DdosProtectionRuleWithStats) GetSourceIPPrefix() string`

GetSourceIPPrefix returns the SourceIPPrefix field if non-nil, zero value otherwise.

### GetSourceIPPrefixOk

`func (o *DdosProtectionRuleWithStats) GetSourceIPPrefixOk() (*string, bool)`

GetSourceIPPrefixOk returns a tuple with the SourceIPPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIPPrefix

`func (o *DdosProtectionRuleWithStats) SetSourceIPPrefix(v string)`

SetSourceIPPrefix sets SourceIPPrefix field to given value.

### HasSourceIPPrefix

`func (o *DdosProtectionRuleWithStats) HasSourceIPPrefix() bool`

HasSourceIPPrefix returns a boolean if a field has been set.

### SetSourceIPPrefixNil

`func (o *DdosProtectionRuleWithStats) SetSourceIPPrefixNil(b bool)`

 SetSourceIPPrefixNil sets the value for SourceIPPrefix to be an explicit nil

### UnsetSourceIPPrefix
`func (o *DdosProtectionRuleWithStats) UnsetSourceIPPrefix()`

UnsetSourceIPPrefix ensures that no value is present for SourceIPPrefix, not even an explicit nil
### GetPath

`func (o *DdosProtectionRuleWithStats) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *DdosProtectionRuleWithStats) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *DdosProtectionRuleWithStats) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *DdosProtectionRuleWithStats) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *DdosProtectionRuleWithStats) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *DdosProtectionRuleWithStats) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetAdditionalAttributes

`func (o *DdosProtectionRuleWithStats) GetAdditionalAttributes() []string`

GetAdditionalAttributes returns the AdditionalAttributes field if non-nil, zero value otherwise.

### GetAdditionalAttributesOk

`func (o *DdosProtectionRuleWithStats) GetAdditionalAttributesOk() (*[]string, bool)`

GetAdditionalAttributesOk returns a tuple with the AdditionalAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalAttributes

`func (o *DdosProtectionRuleWithStats) SetAdditionalAttributes(v []string)`

SetAdditionalAttributes sets AdditionalAttributes field to given value.

### HasAdditionalAttributes

`func (o *DdosProtectionRuleWithStats) HasAdditionalAttributes() bool`

HasAdditionalAttributes returns a boolean if a field has been set.

### GetTrafficStats

`func (o *DdosProtectionRuleWithStats) GetTrafficStats() DdosProtectionTrafficStats`

GetTrafficStats returns the TrafficStats field if non-nil, zero value otherwise.

### GetTrafficStatsOk

`func (o *DdosProtectionRuleWithStats) GetTrafficStatsOk() (*DdosProtectionTrafficStats, bool)`

GetTrafficStatsOk returns a tuple with the TrafficStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficStats

`func (o *DdosProtectionRuleWithStats) SetTrafficStats(v DdosProtectionTrafficStats)`

SetTrafficStats sets TrafficStats field to given value.

### HasTrafficStats

`func (o *DdosProtectionRuleWithStats) HasTrafficStats() bool`

HasTrafficStats returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
