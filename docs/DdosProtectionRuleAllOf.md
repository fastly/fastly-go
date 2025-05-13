# DdosProtectionRuleAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ID** | Pointer to **string** | Unique ID of the rule. | [optional] 
**Name** | Pointer to **string** | A human-readable name for the rule. | [optional] 
**Action** | Pointer to [**DdosProtectionAction**](DdosProtectionAction.md) |  | [optional] [default to DDOSPROTECTIONACTION_DEFAULT]
**CustomerID** | Pointer to **string** | Alphanumeric string identifying the customer. | [optional] 
**ServiceID** | Pointer to **string** | Alphanumeric string identifying the service. | [optional] 
**SourceIP** | Pointer to **NullableString** | Source IP address attribute. | [optional] 
**CountryCode** | Pointer to **NullableString** | Country code attribute. | [optional] 
**Host** | Pointer to **NullableString** | Host attribute. | [optional] 
**Asn** | Pointer to **NullableString** | ASN attribute. | [optional] 
**SourceIPPrefix** | Pointer to **NullableString** | Source IP prefix attribute. | [optional] 
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

### GetID

`func (o *DdosProtectionRuleAllOf) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *DdosProtectionRuleAllOf) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *DdosProtectionRuleAllOf) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *DdosProtectionRuleAllOf) HasID() bool`

HasID returns a boolean if a field has been set.

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

`func (o *DdosProtectionRuleAllOf) GetAction() DdosProtectionAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *DdosProtectionRuleAllOf) GetActionOk() (*DdosProtectionAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *DdosProtectionRuleAllOf) SetAction(v DdosProtectionAction)`

SetAction sets Action field to given value.

### HasAction

`func (o *DdosProtectionRuleAllOf) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetCustomerID

`func (o *DdosProtectionRuleAllOf) GetCustomerID() string`

GetCustomerID returns the CustomerID field if non-nil, zero value otherwise.

### GetCustomerIDOk

`func (o *DdosProtectionRuleAllOf) GetCustomerIDOk() (*string, bool)`

GetCustomerIDOk returns a tuple with the CustomerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerID

`func (o *DdosProtectionRuleAllOf) SetCustomerID(v string)`

SetCustomerID sets CustomerID field to given value.

### HasCustomerID

`func (o *DdosProtectionRuleAllOf) HasCustomerID() bool`

HasCustomerID returns a boolean if a field has been set.

### GetServiceID

`func (o *DdosProtectionRuleAllOf) GetServiceID() string`

GetServiceID returns the ServiceID field if non-nil, zero value otherwise.

### GetServiceIDOk

`func (o *DdosProtectionRuleAllOf) GetServiceIDOk() (*string, bool)`

GetServiceIDOk returns a tuple with the ServiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceID

`func (o *DdosProtectionRuleAllOf) SetServiceID(v string)`

SetServiceID sets ServiceID field to given value.

### HasServiceID

`func (o *DdosProtectionRuleAllOf) HasServiceID() bool`

HasServiceID returns a boolean if a field has been set.

### GetSourceIP

`func (o *DdosProtectionRuleAllOf) GetSourceIP() string`

GetSourceIP returns the SourceIP field if non-nil, zero value otherwise.

### GetSourceIPOk

`func (o *DdosProtectionRuleAllOf) GetSourceIPOk() (*string, bool)`

GetSourceIPOk returns a tuple with the SourceIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIP

`func (o *DdosProtectionRuleAllOf) SetSourceIP(v string)`

SetSourceIP sets SourceIP field to given value.

### HasSourceIP

`func (o *DdosProtectionRuleAllOf) HasSourceIP() bool`

HasSourceIP returns a boolean if a field has been set.

### SetSourceIPNil

`func (o *DdosProtectionRuleAllOf) SetSourceIPNil(b bool)`

 SetSourceIPNil sets the value for SourceIP to be an explicit nil

### UnsetSourceIP
`func (o *DdosProtectionRuleAllOf) UnsetSourceIP()`

UnsetSourceIP ensures that no value is present for SourceIP, not even an explicit nil
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
### GetSourceIPPrefix

`func (o *DdosProtectionRuleAllOf) GetSourceIPPrefix() string`

GetSourceIPPrefix returns the SourceIPPrefix field if non-nil, zero value otherwise.

### GetSourceIPPrefixOk

`func (o *DdosProtectionRuleAllOf) GetSourceIPPrefixOk() (*string, bool)`

GetSourceIPPrefixOk returns a tuple with the SourceIPPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIPPrefix

`func (o *DdosProtectionRuleAllOf) SetSourceIPPrefix(v string)`

SetSourceIPPrefix sets SourceIPPrefix field to given value.

### HasSourceIPPrefix

`func (o *DdosProtectionRuleAllOf) HasSourceIPPrefix() bool`

HasSourceIPPrefix returns a boolean if a field has been set.

### SetSourceIPPrefixNil

`func (o *DdosProtectionRuleAllOf) SetSourceIPPrefixNil(b bool)`

 SetSourceIPPrefixNil sets the value for SourceIPPrefix to be an explicit nil

### UnsetSourceIPPrefix
`func (o *DdosProtectionRuleAllOf) UnsetSourceIPPrefix()`

UnsetSourceIPPrefix ensures that no value is present for SourceIPPrefix, not even an explicit nil
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
