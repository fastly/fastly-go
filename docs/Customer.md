# Customer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingContactId** | Pointer to **NullableString** | The alphanumeric string representing the primary billing contact. | [optional] 
**BillingNetworkType** | Pointer to **string** | Customer&#39;s current network revenue type. | [optional] 
**BillingRef** | Pointer to **NullableString** | Used for adding purchased orders to customer&#39;s account. | [optional] 
**CanConfigureWordpress** | Pointer to **NullableBool** | Whether this customer can view or edit wordpress. | [optional] [readonly] 
**CanResetPasswords** | Pointer to **bool** | Whether this customer can reset passwords. | [optional] [readonly] 
**CanUploadVcl** | Pointer to **bool** | Whether this customer can upload VCL. | [optional] [readonly] 
**Force2fa** | Pointer to **bool** | Specifies whether 2FA is forced or not forced on the customer account. Logs out non-2FA users once 2FA is force enabled. | [optional] 
**ForceSso** | Pointer to **bool** | Specifies whether SSO is forced or not forced on the customer account. | [optional] 
**HasAccountPanel** | Pointer to **bool** | Specifies whether the account has access or does not have access to the account panel. | [optional] 
**HasImprovedEvents** | Pointer to **bool** | Specifies whether the account has access or does not have access to the improved events. | [optional] 
**HasImprovedSslConfig** | Pointer to **bool** | Whether this customer can view or edit the SSL config. | [optional] [readonly] 
**HasOpenstackLogging** | Pointer to **bool** | Specifies whether the account has enabled or not enabled openstack logging. | [optional] 
**HasPci** | Pointer to **bool** | Specifies whether the account can edit PCI for a service. | [optional] 
**HasPciPasswords** | Pointer to **bool** | Specifies whether PCI passwords are required for the account. | [optional] [readonly] 
**IpWhitelist** | Pointer to **string** | The range of IP addresses authorized to access the customer account. | [optional] 
**LegalContactId** | Pointer to **NullableString** | The alphanumeric string identifying the account&#39;s legal contact. | [optional] 
**Name** | Pointer to **string** | The name of the customer, generally the company name. | [optional] 
**OwnerId** | Pointer to **string** | The alphanumeric string identifying the account owner. | [optional] 
**PhoneNumber** | Pointer to **string** | The phone number associated with the account. | [optional] 
**PostalAddress** | Pointer to **NullableString** | The postal address associated with the account. | [optional] 
**PricingPlan** | Pointer to **string** | The pricing plan this customer is under. | [optional] 
**PricingPlanId** | Pointer to **string** | The alphanumeric string identifying the pricing plan. | [optional] 
**SecurityContactId** | Pointer to **NullableString** | The alphanumeric string identifying the account&#39;s security contact. | [optional] 
**TechnicalContactId** | Pointer to **NullableString** | The alphanumeric string identifying the account&#39;s technical contact. | [optional] 

## Methods

### NewCustomer

`func NewCustomer() *Customer`

NewCustomer instantiates a new Customer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerWithDefaults

`func NewCustomerWithDefaults() *Customer`

NewCustomerWithDefaults instantiates a new Customer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingContactId

`func (o *Customer) GetBillingContactId() string`

GetBillingContactId returns the BillingContactId field if non-nil, zero value otherwise.

### GetBillingContactIdOk

`func (o *Customer) GetBillingContactIdOk() (*string, bool)`

GetBillingContactIdOk returns a tuple with the BillingContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingContactId

`func (o *Customer) SetBillingContactId(v string)`

SetBillingContactId sets BillingContactId field to given value.

### HasBillingContactId

`func (o *Customer) HasBillingContactId() bool`

HasBillingContactId returns a boolean if a field has been set.

### SetBillingContactIdNil

`func (o *Customer) SetBillingContactIdNil(b bool)`

 SetBillingContactIdNil sets the value for BillingContactId to be an explicit nil

### UnsetBillingContactId
`func (o *Customer) UnsetBillingContactId()`

UnsetBillingContactId ensures that no value is present for BillingContactId, not even an explicit nil
### GetBillingNetworkType

`func (o *Customer) GetBillingNetworkType() string`

GetBillingNetworkType returns the BillingNetworkType field if non-nil, zero value otherwise.

### GetBillingNetworkTypeOk

`func (o *Customer) GetBillingNetworkTypeOk() (*string, bool)`

GetBillingNetworkTypeOk returns a tuple with the BillingNetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingNetworkType

`func (o *Customer) SetBillingNetworkType(v string)`

SetBillingNetworkType sets BillingNetworkType field to given value.

### HasBillingNetworkType

`func (o *Customer) HasBillingNetworkType() bool`

HasBillingNetworkType returns a boolean if a field has been set.

### GetBillingRef

`func (o *Customer) GetBillingRef() string`

GetBillingRef returns the BillingRef field if non-nil, zero value otherwise.

### GetBillingRefOk

`func (o *Customer) GetBillingRefOk() (*string, bool)`

GetBillingRefOk returns a tuple with the BillingRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingRef

`func (o *Customer) SetBillingRef(v string)`

SetBillingRef sets BillingRef field to given value.

### HasBillingRef

`func (o *Customer) HasBillingRef() bool`

HasBillingRef returns a boolean if a field has been set.

### SetBillingRefNil

`func (o *Customer) SetBillingRefNil(b bool)`

 SetBillingRefNil sets the value for BillingRef to be an explicit nil

### UnsetBillingRef
`func (o *Customer) UnsetBillingRef()`

UnsetBillingRef ensures that no value is present for BillingRef, not even an explicit nil
### GetCanConfigureWordpress

`func (o *Customer) GetCanConfigureWordpress() bool`

GetCanConfigureWordpress returns the CanConfigureWordpress field if non-nil, zero value otherwise.

### GetCanConfigureWordpressOk

`func (o *Customer) GetCanConfigureWordpressOk() (*bool, bool)`

GetCanConfigureWordpressOk returns a tuple with the CanConfigureWordpress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanConfigureWordpress

`func (o *Customer) SetCanConfigureWordpress(v bool)`

SetCanConfigureWordpress sets CanConfigureWordpress field to given value.

### HasCanConfigureWordpress

`func (o *Customer) HasCanConfigureWordpress() bool`

HasCanConfigureWordpress returns a boolean if a field has been set.

### SetCanConfigureWordpressNil

`func (o *Customer) SetCanConfigureWordpressNil(b bool)`

 SetCanConfigureWordpressNil sets the value for CanConfigureWordpress to be an explicit nil

### UnsetCanConfigureWordpress
`func (o *Customer) UnsetCanConfigureWordpress()`

UnsetCanConfigureWordpress ensures that no value is present for CanConfigureWordpress, not even an explicit nil
### GetCanResetPasswords

`func (o *Customer) GetCanResetPasswords() bool`

GetCanResetPasswords returns the CanResetPasswords field if non-nil, zero value otherwise.

### GetCanResetPasswordsOk

`func (o *Customer) GetCanResetPasswordsOk() (*bool, bool)`

GetCanResetPasswordsOk returns a tuple with the CanResetPasswords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanResetPasswords

`func (o *Customer) SetCanResetPasswords(v bool)`

SetCanResetPasswords sets CanResetPasswords field to given value.

### HasCanResetPasswords

`func (o *Customer) HasCanResetPasswords() bool`

HasCanResetPasswords returns a boolean if a field has been set.

### GetCanUploadVcl

`func (o *Customer) GetCanUploadVcl() bool`

GetCanUploadVcl returns the CanUploadVcl field if non-nil, zero value otherwise.

### GetCanUploadVclOk

`func (o *Customer) GetCanUploadVclOk() (*bool, bool)`

GetCanUploadVclOk returns a tuple with the CanUploadVcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanUploadVcl

`func (o *Customer) SetCanUploadVcl(v bool)`

SetCanUploadVcl sets CanUploadVcl field to given value.

### HasCanUploadVcl

`func (o *Customer) HasCanUploadVcl() bool`

HasCanUploadVcl returns a boolean if a field has been set.

### GetForce2fa

`func (o *Customer) GetForce2fa() bool`

GetForce2fa returns the Force2fa field if non-nil, zero value otherwise.

### GetForce2faOk

`func (o *Customer) GetForce2faOk() (*bool, bool)`

GetForce2faOk returns a tuple with the Force2fa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForce2fa

`func (o *Customer) SetForce2fa(v bool)`

SetForce2fa sets Force2fa field to given value.

### HasForce2fa

`func (o *Customer) HasForce2fa() bool`

HasForce2fa returns a boolean if a field has been set.

### GetForceSso

`func (o *Customer) GetForceSso() bool`

GetForceSso returns the ForceSso field if non-nil, zero value otherwise.

### GetForceSsoOk

`func (o *Customer) GetForceSsoOk() (*bool, bool)`

GetForceSsoOk returns a tuple with the ForceSso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceSso

`func (o *Customer) SetForceSso(v bool)`

SetForceSso sets ForceSso field to given value.

### HasForceSso

`func (o *Customer) HasForceSso() bool`

HasForceSso returns a boolean if a field has been set.

### GetHasAccountPanel

`func (o *Customer) GetHasAccountPanel() bool`

GetHasAccountPanel returns the HasAccountPanel field if non-nil, zero value otherwise.

### GetHasAccountPanelOk

`func (o *Customer) GetHasAccountPanelOk() (*bool, bool)`

GetHasAccountPanelOk returns a tuple with the HasAccountPanel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAccountPanel

`func (o *Customer) SetHasAccountPanel(v bool)`

SetHasAccountPanel sets HasAccountPanel field to given value.

### HasHasAccountPanel

`func (o *Customer) HasHasAccountPanel() bool`

HasHasAccountPanel returns a boolean if a field has been set.

### GetHasImprovedEvents

`func (o *Customer) GetHasImprovedEvents() bool`

GetHasImprovedEvents returns the HasImprovedEvents field if non-nil, zero value otherwise.

### GetHasImprovedEventsOk

`func (o *Customer) GetHasImprovedEventsOk() (*bool, bool)`

GetHasImprovedEventsOk returns a tuple with the HasImprovedEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasImprovedEvents

`func (o *Customer) SetHasImprovedEvents(v bool)`

SetHasImprovedEvents sets HasImprovedEvents field to given value.

### HasHasImprovedEvents

`func (o *Customer) HasHasImprovedEvents() bool`

HasHasImprovedEvents returns a boolean if a field has been set.

### GetHasImprovedSslConfig

`func (o *Customer) GetHasImprovedSslConfig() bool`

GetHasImprovedSslConfig returns the HasImprovedSslConfig field if non-nil, zero value otherwise.

### GetHasImprovedSslConfigOk

`func (o *Customer) GetHasImprovedSslConfigOk() (*bool, bool)`

GetHasImprovedSslConfigOk returns a tuple with the HasImprovedSslConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasImprovedSslConfig

`func (o *Customer) SetHasImprovedSslConfig(v bool)`

SetHasImprovedSslConfig sets HasImprovedSslConfig field to given value.

### HasHasImprovedSslConfig

`func (o *Customer) HasHasImprovedSslConfig() bool`

HasHasImprovedSslConfig returns a boolean if a field has been set.

### GetHasOpenstackLogging

`func (o *Customer) GetHasOpenstackLogging() bool`

GetHasOpenstackLogging returns the HasOpenstackLogging field if non-nil, zero value otherwise.

### GetHasOpenstackLoggingOk

`func (o *Customer) GetHasOpenstackLoggingOk() (*bool, bool)`

GetHasOpenstackLoggingOk returns a tuple with the HasOpenstackLogging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasOpenstackLogging

`func (o *Customer) SetHasOpenstackLogging(v bool)`

SetHasOpenstackLogging sets HasOpenstackLogging field to given value.

### HasHasOpenstackLogging

`func (o *Customer) HasHasOpenstackLogging() bool`

HasHasOpenstackLogging returns a boolean if a field has been set.

### GetHasPci

`func (o *Customer) GetHasPci() bool`

GetHasPci returns the HasPci field if non-nil, zero value otherwise.

### GetHasPciOk

`func (o *Customer) GetHasPciOk() (*bool, bool)`

GetHasPciOk returns a tuple with the HasPci field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPci

`func (o *Customer) SetHasPci(v bool)`

SetHasPci sets HasPci field to given value.

### HasHasPci

`func (o *Customer) HasHasPci() bool`

HasHasPci returns a boolean if a field has been set.

### GetHasPciPasswords

`func (o *Customer) GetHasPciPasswords() bool`

GetHasPciPasswords returns the HasPciPasswords field if non-nil, zero value otherwise.

### GetHasPciPasswordsOk

`func (o *Customer) GetHasPciPasswordsOk() (*bool, bool)`

GetHasPciPasswordsOk returns a tuple with the HasPciPasswords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPciPasswords

`func (o *Customer) SetHasPciPasswords(v bool)`

SetHasPciPasswords sets HasPciPasswords field to given value.

### HasHasPciPasswords

`func (o *Customer) HasHasPciPasswords() bool`

HasHasPciPasswords returns a boolean if a field has been set.

### GetIpWhitelist

`func (o *Customer) GetIpWhitelist() string`

GetIpWhitelist returns the IpWhitelist field if non-nil, zero value otherwise.

### GetIpWhitelistOk

`func (o *Customer) GetIpWhitelistOk() (*string, bool)`

GetIpWhitelistOk returns a tuple with the IpWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpWhitelist

`func (o *Customer) SetIpWhitelist(v string)`

SetIpWhitelist sets IpWhitelist field to given value.

### HasIpWhitelist

`func (o *Customer) HasIpWhitelist() bool`

HasIpWhitelist returns a boolean if a field has been set.

### GetLegalContactId

`func (o *Customer) GetLegalContactId() string`

GetLegalContactId returns the LegalContactId field if non-nil, zero value otherwise.

### GetLegalContactIdOk

`func (o *Customer) GetLegalContactIdOk() (*string, bool)`

GetLegalContactIdOk returns a tuple with the LegalContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalContactId

`func (o *Customer) SetLegalContactId(v string)`

SetLegalContactId sets LegalContactId field to given value.

### HasLegalContactId

`func (o *Customer) HasLegalContactId() bool`

HasLegalContactId returns a boolean if a field has been set.

### SetLegalContactIdNil

`func (o *Customer) SetLegalContactIdNil(b bool)`

 SetLegalContactIdNil sets the value for LegalContactId to be an explicit nil

### UnsetLegalContactId
`func (o *Customer) UnsetLegalContactId()`

UnsetLegalContactId ensures that no value is present for LegalContactId, not even an explicit nil
### GetName

`func (o *Customer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Customer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Customer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Customer) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwnerId

`func (o *Customer) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *Customer) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *Customer) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *Customer) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *Customer) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *Customer) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *Customer) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *Customer) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetPostalAddress

`func (o *Customer) GetPostalAddress() string`

GetPostalAddress returns the PostalAddress field if non-nil, zero value otherwise.

### GetPostalAddressOk

`func (o *Customer) GetPostalAddressOk() (*string, bool)`

GetPostalAddressOk returns a tuple with the PostalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalAddress

`func (o *Customer) SetPostalAddress(v string)`

SetPostalAddress sets PostalAddress field to given value.

### HasPostalAddress

`func (o *Customer) HasPostalAddress() bool`

HasPostalAddress returns a boolean if a field has been set.

### SetPostalAddressNil

`func (o *Customer) SetPostalAddressNil(b bool)`

 SetPostalAddressNil sets the value for PostalAddress to be an explicit nil

### UnsetPostalAddress
`func (o *Customer) UnsetPostalAddress()`

UnsetPostalAddress ensures that no value is present for PostalAddress, not even an explicit nil
### GetPricingPlan

`func (o *Customer) GetPricingPlan() string`

GetPricingPlan returns the PricingPlan field if non-nil, zero value otherwise.

### GetPricingPlanOk

`func (o *Customer) GetPricingPlanOk() (*string, bool)`

GetPricingPlanOk returns a tuple with the PricingPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingPlan

`func (o *Customer) SetPricingPlan(v string)`

SetPricingPlan sets PricingPlan field to given value.

### HasPricingPlan

`func (o *Customer) HasPricingPlan() bool`

HasPricingPlan returns a boolean if a field has been set.

### GetPricingPlanId

`func (o *Customer) GetPricingPlanId() string`

GetPricingPlanId returns the PricingPlanId field if non-nil, zero value otherwise.

### GetPricingPlanIdOk

`func (o *Customer) GetPricingPlanIdOk() (*string, bool)`

GetPricingPlanIdOk returns a tuple with the PricingPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingPlanId

`func (o *Customer) SetPricingPlanId(v string)`

SetPricingPlanId sets PricingPlanId field to given value.

### HasPricingPlanId

`func (o *Customer) HasPricingPlanId() bool`

HasPricingPlanId returns a boolean if a field has been set.

### GetSecurityContactId

`func (o *Customer) GetSecurityContactId() string`

GetSecurityContactId returns the SecurityContactId field if non-nil, zero value otherwise.

### GetSecurityContactIdOk

`func (o *Customer) GetSecurityContactIdOk() (*string, bool)`

GetSecurityContactIdOk returns a tuple with the SecurityContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityContactId

`func (o *Customer) SetSecurityContactId(v string)`

SetSecurityContactId sets SecurityContactId field to given value.

### HasSecurityContactId

`func (o *Customer) HasSecurityContactId() bool`

HasSecurityContactId returns a boolean if a field has been set.

### SetSecurityContactIdNil

`func (o *Customer) SetSecurityContactIdNil(b bool)`

 SetSecurityContactIdNil sets the value for SecurityContactId to be an explicit nil

### UnsetSecurityContactId
`func (o *Customer) UnsetSecurityContactId()`

UnsetSecurityContactId ensures that no value is present for SecurityContactId, not even an explicit nil
### GetTechnicalContactId

`func (o *Customer) GetTechnicalContactId() string`

GetTechnicalContactId returns the TechnicalContactId field if non-nil, zero value otherwise.

### GetTechnicalContactIdOk

`func (o *Customer) GetTechnicalContactIdOk() (*string, bool)`

GetTechnicalContactIdOk returns a tuple with the TechnicalContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechnicalContactId

`func (o *Customer) SetTechnicalContactId(v string)`

SetTechnicalContactId sets TechnicalContactId field to given value.

### HasTechnicalContactId

`func (o *Customer) HasTechnicalContactId() bool`

HasTechnicalContactId returns a boolean if a field has been set.

### SetTechnicalContactIdNil

`func (o *Customer) SetTechnicalContactIdNil(b bool)`

 SetTechnicalContactIdNil sets the value for TechnicalContactId to be an explicit nil

### UnsetTechnicalContactId
`func (o *Customer) UnsetTechnicalContactId()`

UnsetTechnicalContactId ensures that no value is present for TechnicalContactId, not even an explicit nil

[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


