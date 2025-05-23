# DdosProtectionRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**UpdatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
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
### GetID

`func (o *DdosProtectionRule) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *DdosProtectionRule) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *DdosProtectionRule) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *DdosProtectionRule) HasID() bool`

HasID returns a boolean if a field has been set.

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

`func (o *DdosProtectionRule) GetAction() DdosProtectionAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *DdosProtectionRule) GetActionOk() (*DdosProtectionAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *DdosProtectionRule) SetAction(v DdosProtectionAction)`

SetAction sets Action field to given value.

### HasAction

`func (o *DdosProtectionRule) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetCustomerID

`func (o *DdosProtectionRule) GetCustomerID() string`

GetCustomerID returns the CustomerID field if non-nil, zero value otherwise.

### GetCustomerIDOk

`func (o *DdosProtectionRule) GetCustomerIDOk() (*string, bool)`

GetCustomerIDOk returns a tuple with the CustomerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerID

`func (o *DdosProtectionRule) SetCustomerID(v string)`

SetCustomerID sets CustomerID field to given value.

### HasCustomerID

`func (o *DdosProtectionRule) HasCustomerID() bool`

HasCustomerID returns a boolean if a field has been set.

### GetServiceID

`func (o *DdosProtectionRule) GetServiceID() string`

GetServiceID returns the ServiceID field if non-nil, zero value otherwise.

### GetServiceIDOk

`func (o *DdosProtectionRule) GetServiceIDOk() (*string, bool)`

GetServiceIDOk returns a tuple with the ServiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceID

`func (o *DdosProtectionRule) SetServiceID(v string)`

SetServiceID sets ServiceID field to given value.

### HasServiceID

`func (o *DdosProtectionRule) HasServiceID() bool`

HasServiceID returns a boolean if a field has been set.

### GetSourceIP

`func (o *DdosProtectionRule) GetSourceIP() string`

GetSourceIP returns the SourceIP field if non-nil, zero value otherwise.

### GetSourceIPOk

`func (o *DdosProtectionRule) GetSourceIPOk() (*string, bool)`

GetSourceIPOk returns a tuple with the SourceIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIP

`func (o *DdosProtectionRule) SetSourceIP(v string)`

SetSourceIP sets SourceIP field to given value.

### HasSourceIP

`func (o *DdosProtectionRule) HasSourceIP() bool`

HasSourceIP returns a boolean if a field has been set.

### SetSourceIPNil

`func (o *DdosProtectionRule) SetSourceIPNil(b bool)`

 SetSourceIPNil sets the value for SourceIP to be an explicit nil

### UnsetSourceIP
`func (o *DdosProtectionRule) UnsetSourceIP()`

UnsetSourceIP ensures that no value is present for SourceIP, not even an explicit nil
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
### GetSourceIPPrefix

`func (o *DdosProtectionRule) GetSourceIPPrefix() string`

GetSourceIPPrefix returns the SourceIPPrefix field if non-nil, zero value otherwise.

### GetSourceIPPrefixOk

`func (o *DdosProtectionRule) GetSourceIPPrefixOk() (*string, bool)`

GetSourceIPPrefixOk returns a tuple with the SourceIPPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIPPrefix

`func (o *DdosProtectionRule) SetSourceIPPrefix(v string)`

SetSourceIPPrefix sets SourceIPPrefix field to given value.

### HasSourceIPPrefix

`func (o *DdosProtectionRule) HasSourceIPPrefix() bool`

HasSourceIPPrefix returns a boolean if a field has been set.

### SetSourceIPPrefixNil

`func (o *DdosProtectionRule) SetSourceIPPrefixNil(b bool)`

 SetSourceIPPrefixNil sets the value for SourceIPPrefix to be an explicit nil

### UnsetSourceIPPrefix
`func (o *DdosProtectionRule) UnsetSourceIPPrefix()`

UnsetSourceIPPrefix ensures that no value is present for SourceIPPrefix, not even an explicit nil
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
