# Customer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingContactID** | Pointer to **NullableString** | The alphanumeric string representing the primary billing contact. | [optional] 
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
**IPWhitelist** | Pointer to **string** | The range of IP addresses authorized to access the customer account. | [optional] 
**LegalContactID** | Pointer to **NullableString** | The alphanumeric string identifying the account&#39;s legal contact. | [optional] 
**Name** | Pointer to **string** | The name of the customer, generally the company name. | [optional] 
**OwnerID** | Pointer to **string** | The alphanumeric string identifying the account owner. | [optional] 
**PhoneNumber** | Pointer to **string** | The phone number associated with the account. | [optional] 
**PostalAddress** | Pointer to **NullableString** | The postal address associated with the account. | [optional] 
**PricingPlan** | Pointer to **string** | The pricing plan this customer is under. | [optional] 
**PricingPlanID** | Pointer to **string** | The alphanumeric string identifying the pricing plan. | [optional] 
**SecurityContactID** | Pointer to **NullableString** | The alphanumeric string identifying the account&#39;s security contact. | [optional] 
**TechnicalContactID** | Pointer to **NullableString** | The alphanumeric string identifying the account&#39;s technical contact. | [optional] 

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

### GetBillingContactID

`func (o *Customer) GetBillingContactID() string`

GetBillingContactID returns the BillingContactID field if non-nil, zero value otherwise.

### GetBillingContactIDOk

`func (o *Customer) GetBillingContactIDOk() (*string, bool)`

GetBillingContactIDOk returns a tuple with the BillingContactID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingContactID

`func (o *Customer) SetBillingContactID(v string)`

SetBillingContactID sets BillingContactID field to given value.

### HasBillingContactID

`func (o *Customer) HasBillingContactID() bool`

HasBillingContactID returns a boolean if a field has been set.

### SetBillingContactIDNil

`func (o *Customer) SetBillingContactIDNil(b bool)`

 SetBillingContactIDNil sets the value for BillingContactID to be an explicit nil

### UnsetBillingContactID
`func (o *Customer) UnsetBillingContactID()`

UnsetBillingContactID ensures that no value is present for BillingContactID, not even an explicit nil
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

### GetIPWhitelist

`func (o *Customer) GetIPWhitelist() string`

GetIPWhitelist returns the IPWhitelist field if non-nil, zero value otherwise.

### GetIPWhitelistOk

`func (o *Customer) GetIPWhitelistOk() (*string, bool)`

GetIPWhitelistOk returns a tuple with the IPWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIPWhitelist

`func (o *Customer) SetIPWhitelist(v string)`

SetIPWhitelist sets IPWhitelist field to given value.

### HasIPWhitelist

`func (o *Customer) HasIPWhitelist() bool`

HasIPWhitelist returns a boolean if a field has been set.

### GetLegalContactID

`func (o *Customer) GetLegalContactID() string`

GetLegalContactID returns the LegalContactID field if non-nil, zero value otherwise.

### GetLegalContactIDOk

`func (o *Customer) GetLegalContactIDOk() (*string, bool)`

GetLegalContactIDOk returns a tuple with the LegalContactID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalContactID

`func (o *Customer) SetLegalContactID(v string)`

SetLegalContactID sets LegalContactID field to given value.

### HasLegalContactID

`func (o *Customer) HasLegalContactID() bool`

HasLegalContactID returns a boolean if a field has been set.

### SetLegalContactIDNil

`func (o *Customer) SetLegalContactIDNil(b bool)`

 SetLegalContactIDNil sets the value for LegalContactID to be an explicit nil

### UnsetLegalContactID
`func (o *Customer) UnsetLegalContactID()`

UnsetLegalContactID ensures that no value is present for LegalContactID, not even an explicit nil
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

### GetOwnerID

`func (o *Customer) GetOwnerID() string`

GetOwnerID returns the OwnerID field if non-nil, zero value otherwise.

### GetOwnerIDOk

`func (o *Customer) GetOwnerIDOk() (*string, bool)`

GetOwnerIDOk returns a tuple with the OwnerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerID

`func (o *Customer) SetOwnerID(v string)`

SetOwnerID sets OwnerID field to given value.

### HasOwnerID

`func (o *Customer) HasOwnerID() bool`

HasOwnerID returns a boolean if a field has been set.

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

### GetPricingPlanID

`func (o *Customer) GetPricingPlanID() string`

GetPricingPlanID returns the PricingPlanID field if non-nil, zero value otherwise.

### GetPricingPlanIDOk

`func (o *Customer) GetPricingPlanIDOk() (*string, bool)`

GetPricingPlanIDOk returns a tuple with the PricingPlanID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingPlanID

`func (o *Customer) SetPricingPlanID(v string)`

SetPricingPlanID sets PricingPlanID field to given value.

### HasPricingPlanID

`func (o *Customer) HasPricingPlanID() bool`

HasPricingPlanID returns a boolean if a field has been set.

### GetSecurityContactID

`func (o *Customer) GetSecurityContactID() string`

GetSecurityContactID returns the SecurityContactID field if non-nil, zero value otherwise.

### GetSecurityContactIDOk

`func (o *Customer) GetSecurityContactIDOk() (*string, bool)`

GetSecurityContactIDOk returns a tuple with the SecurityContactID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityContactID

`func (o *Customer) SetSecurityContactID(v string)`

SetSecurityContactID sets SecurityContactID field to given value.

### HasSecurityContactID

`func (o *Customer) HasSecurityContactID() bool`

HasSecurityContactID returns a boolean if a field has been set.

### SetSecurityContactIDNil

`func (o *Customer) SetSecurityContactIDNil(b bool)`

 SetSecurityContactIDNil sets the value for SecurityContactID to be an explicit nil

### UnsetSecurityContactID
`func (o *Customer) UnsetSecurityContactID()`

UnsetSecurityContactID ensures that no value is present for SecurityContactID, not even an explicit nil
### GetTechnicalContactID

`func (o *Customer) GetTechnicalContactID() string`

GetTechnicalContactID returns the TechnicalContactID field if non-nil, zero value otherwise.

### GetTechnicalContactIDOk

`func (o *Customer) GetTechnicalContactIDOk() (*string, bool)`

GetTechnicalContactIDOk returns a tuple with the TechnicalContactID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechnicalContactID

`func (o *Customer) SetTechnicalContactID(v string)`

SetTechnicalContactID sets TechnicalContactID field to given value.

### HasTechnicalContactID

`func (o *Customer) HasTechnicalContactID() bool`

HasTechnicalContactID returns a boolean if a field has been set.

### SetTechnicalContactIDNil

`func (o *Customer) SetTechnicalContactIDNil(b bool)`

 SetTechnicalContactIDNil sets the value for TechnicalContactID to be an explicit nil

### UnsetTechnicalContactID
`func (o *Customer) UnsetTechnicalContactID()`

UnsetTechnicalContactID ensures that no value is present for TechnicalContactID, not even an explicit nil

[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
