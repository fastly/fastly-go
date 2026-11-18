# CustomerResponse

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
**CreatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**DeletedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**UpdatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**ID** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewCustomerResponse

`func NewCustomerResponse() *CustomerResponse`

NewCustomerResponse instantiates a new CustomerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerResponseWithDefaults

`func NewCustomerResponseWithDefaults() *CustomerResponse`

NewCustomerResponseWithDefaults instantiates a new CustomerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingContactID

`func (o *CustomerResponse) GetBillingContactID() string`

GetBillingContactID returns the BillingContactID field if non-nil, zero value otherwise.

### GetBillingContactIDOk

`func (o *CustomerResponse) GetBillingContactIDOk() (*string, bool)`

GetBillingContactIDOk returns a tuple with the BillingContactID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingContactID

`func (o *CustomerResponse) SetBillingContactID(v string)`

SetBillingContactID sets BillingContactID field to given value.

### HasBillingContactID

`func (o *CustomerResponse) HasBillingContactID() bool`

HasBillingContactID returns a boolean if a field has been set.

### SetBillingContactIDNil

`func (o *CustomerResponse) SetBillingContactIDNil(b bool)`

 SetBillingContactIDNil sets the value for BillingContactID to be an explicit nil

### UnsetBillingContactID
`func (o *CustomerResponse) UnsetBillingContactID()`

UnsetBillingContactID ensures that no value is present for BillingContactID, not even an explicit nil
### GetBillingNetworkType

`func (o *CustomerResponse) GetBillingNetworkType() string`

GetBillingNetworkType returns the BillingNetworkType field if non-nil, zero value otherwise.

### GetBillingNetworkTypeOk

`func (o *CustomerResponse) GetBillingNetworkTypeOk() (*string, bool)`

GetBillingNetworkTypeOk returns a tuple with the BillingNetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingNetworkType

`func (o *CustomerResponse) SetBillingNetworkType(v string)`

SetBillingNetworkType sets BillingNetworkType field to given value.

### HasBillingNetworkType

`func (o *CustomerResponse) HasBillingNetworkType() bool`

HasBillingNetworkType returns a boolean if a field has been set.

### GetBillingRef

`func (o *CustomerResponse) GetBillingRef() string`

GetBillingRef returns the BillingRef field if non-nil, zero value otherwise.

### GetBillingRefOk

`func (o *CustomerResponse) GetBillingRefOk() (*string, bool)`

GetBillingRefOk returns a tuple with the BillingRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingRef

`func (o *CustomerResponse) SetBillingRef(v string)`

SetBillingRef sets BillingRef field to given value.

### HasBillingRef

`func (o *CustomerResponse) HasBillingRef() bool`

HasBillingRef returns a boolean if a field has been set.

### SetBillingRefNil

`func (o *CustomerResponse) SetBillingRefNil(b bool)`

 SetBillingRefNil sets the value for BillingRef to be an explicit nil

### UnsetBillingRef
`func (o *CustomerResponse) UnsetBillingRef()`

UnsetBillingRef ensures that no value is present for BillingRef, not even an explicit nil
### GetCanConfigureWordpress

`func (o *CustomerResponse) GetCanConfigureWordpress() bool`

GetCanConfigureWordpress returns the CanConfigureWordpress field if non-nil, zero value otherwise.

### GetCanConfigureWordpressOk

`func (o *CustomerResponse) GetCanConfigureWordpressOk() (*bool, bool)`

GetCanConfigureWordpressOk returns a tuple with the CanConfigureWordpress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanConfigureWordpress

`func (o *CustomerResponse) SetCanConfigureWordpress(v bool)`

SetCanConfigureWordpress sets CanConfigureWordpress field to given value.

### HasCanConfigureWordpress

`func (o *CustomerResponse) HasCanConfigureWordpress() bool`

HasCanConfigureWordpress returns a boolean if a field has been set.

### SetCanConfigureWordpressNil

`func (o *CustomerResponse) SetCanConfigureWordpressNil(b bool)`

 SetCanConfigureWordpressNil sets the value for CanConfigureWordpress to be an explicit nil

### UnsetCanConfigureWordpress
`func (o *CustomerResponse) UnsetCanConfigureWordpress()`

UnsetCanConfigureWordpress ensures that no value is present for CanConfigureWordpress, not even an explicit nil
### GetCanResetPasswords

`func (o *CustomerResponse) GetCanResetPasswords() bool`

GetCanResetPasswords returns the CanResetPasswords field if non-nil, zero value otherwise.

### GetCanResetPasswordsOk

`func (o *CustomerResponse) GetCanResetPasswordsOk() (*bool, bool)`

GetCanResetPasswordsOk returns a tuple with the CanResetPasswords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanResetPasswords

`func (o *CustomerResponse) SetCanResetPasswords(v bool)`

SetCanResetPasswords sets CanResetPasswords field to given value.

### HasCanResetPasswords

`func (o *CustomerResponse) HasCanResetPasswords() bool`

HasCanResetPasswords returns a boolean if a field has been set.

### GetCanUploadVcl

`func (o *CustomerResponse) GetCanUploadVcl() bool`

GetCanUploadVcl returns the CanUploadVcl field if non-nil, zero value otherwise.

### GetCanUploadVclOk

`func (o *CustomerResponse) GetCanUploadVclOk() (*bool, bool)`

GetCanUploadVclOk returns a tuple with the CanUploadVcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanUploadVcl

`func (o *CustomerResponse) SetCanUploadVcl(v bool)`

SetCanUploadVcl sets CanUploadVcl field to given value.

### HasCanUploadVcl

`func (o *CustomerResponse) HasCanUploadVcl() bool`

HasCanUploadVcl returns a boolean if a field has been set.

### GetForce2fa

`func (o *CustomerResponse) GetForce2fa() bool`

GetForce2fa returns the Force2fa field if non-nil, zero value otherwise.

### GetForce2faOk

`func (o *CustomerResponse) GetForce2faOk() (*bool, bool)`

GetForce2faOk returns a tuple with the Force2fa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForce2fa

`func (o *CustomerResponse) SetForce2fa(v bool)`

SetForce2fa sets Force2fa field to given value.

### HasForce2fa

`func (o *CustomerResponse) HasForce2fa() bool`

HasForce2fa returns a boolean if a field has been set.

### GetForceSso

`func (o *CustomerResponse) GetForceSso() bool`

GetForceSso returns the ForceSso field if non-nil, zero value otherwise.

### GetForceSsoOk

`func (o *CustomerResponse) GetForceSsoOk() (*bool, bool)`

GetForceSsoOk returns a tuple with the ForceSso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceSso

`func (o *CustomerResponse) SetForceSso(v bool)`

SetForceSso sets ForceSso field to given value.

### HasForceSso

`func (o *CustomerResponse) HasForceSso() bool`

HasForceSso returns a boolean if a field has been set.

### GetHasAccountPanel

`func (o *CustomerResponse) GetHasAccountPanel() bool`

GetHasAccountPanel returns the HasAccountPanel field if non-nil, zero value otherwise.

### GetHasAccountPanelOk

`func (o *CustomerResponse) GetHasAccountPanelOk() (*bool, bool)`

GetHasAccountPanelOk returns a tuple with the HasAccountPanel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAccountPanel

`func (o *CustomerResponse) SetHasAccountPanel(v bool)`

SetHasAccountPanel sets HasAccountPanel field to given value.

### HasHasAccountPanel

`func (o *CustomerResponse) HasHasAccountPanel() bool`

HasHasAccountPanel returns a boolean if a field has been set.

### GetHasImprovedEvents

`func (o *CustomerResponse) GetHasImprovedEvents() bool`

GetHasImprovedEvents returns the HasImprovedEvents field if non-nil, zero value otherwise.

### GetHasImprovedEventsOk

`func (o *CustomerResponse) GetHasImprovedEventsOk() (*bool, bool)`

GetHasImprovedEventsOk returns a tuple with the HasImprovedEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasImprovedEvents

`func (o *CustomerResponse) SetHasImprovedEvents(v bool)`

SetHasImprovedEvents sets HasImprovedEvents field to given value.

### HasHasImprovedEvents

`func (o *CustomerResponse) HasHasImprovedEvents() bool`

HasHasImprovedEvents returns a boolean if a field has been set.

### GetHasImprovedSslConfig

`func (o *CustomerResponse) GetHasImprovedSslConfig() bool`

GetHasImprovedSslConfig returns the HasImprovedSslConfig field if non-nil, zero value otherwise.

### GetHasImprovedSslConfigOk

`func (o *CustomerResponse) GetHasImprovedSslConfigOk() (*bool, bool)`

GetHasImprovedSslConfigOk returns a tuple with the HasImprovedSslConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasImprovedSslConfig

`func (o *CustomerResponse) SetHasImprovedSslConfig(v bool)`

SetHasImprovedSslConfig sets HasImprovedSslConfig field to given value.

### HasHasImprovedSslConfig

`func (o *CustomerResponse) HasHasImprovedSslConfig() bool`

HasHasImprovedSslConfig returns a boolean if a field has been set.

### GetHasOpenstackLogging

`func (o *CustomerResponse) GetHasOpenstackLogging() bool`

GetHasOpenstackLogging returns the HasOpenstackLogging field if non-nil, zero value otherwise.

### GetHasOpenstackLoggingOk

`func (o *CustomerResponse) GetHasOpenstackLoggingOk() (*bool, bool)`

GetHasOpenstackLoggingOk returns a tuple with the HasOpenstackLogging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasOpenstackLogging

`func (o *CustomerResponse) SetHasOpenstackLogging(v bool)`

SetHasOpenstackLogging sets HasOpenstackLogging field to given value.

### HasHasOpenstackLogging

`func (o *CustomerResponse) HasHasOpenstackLogging() bool`

HasHasOpenstackLogging returns a boolean if a field has been set.

### GetHasPci

`func (o *CustomerResponse) GetHasPci() bool`

GetHasPci returns the HasPci field if non-nil, zero value otherwise.

### GetHasPciOk

`func (o *CustomerResponse) GetHasPciOk() (*bool, bool)`

GetHasPciOk returns a tuple with the HasPci field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPci

`func (o *CustomerResponse) SetHasPci(v bool)`

SetHasPci sets HasPci field to given value.

### HasHasPci

`func (o *CustomerResponse) HasHasPci() bool`

HasHasPci returns a boolean if a field has been set.

### GetHasPciPasswords

`func (o *CustomerResponse) GetHasPciPasswords() bool`

GetHasPciPasswords returns the HasPciPasswords field if non-nil, zero value otherwise.

### GetHasPciPasswordsOk

`func (o *CustomerResponse) GetHasPciPasswordsOk() (*bool, bool)`

GetHasPciPasswordsOk returns a tuple with the HasPciPasswords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPciPasswords

`func (o *CustomerResponse) SetHasPciPasswords(v bool)`

SetHasPciPasswords sets HasPciPasswords field to given value.

### HasHasPciPasswords

`func (o *CustomerResponse) HasHasPciPasswords() bool`

HasHasPciPasswords returns a boolean if a field has been set.

### GetIPWhitelist

`func (o *CustomerResponse) GetIPWhitelist() string`

GetIPWhitelist returns the IPWhitelist field if non-nil, zero value otherwise.

### GetIPWhitelistOk

`func (o *CustomerResponse) GetIPWhitelistOk() (*string, bool)`

GetIPWhitelistOk returns a tuple with the IPWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIPWhitelist

`func (o *CustomerResponse) SetIPWhitelist(v string)`

SetIPWhitelist sets IPWhitelist field to given value.

### HasIPWhitelist

`func (o *CustomerResponse) HasIPWhitelist() bool`

HasIPWhitelist returns a boolean if a field has been set.

### GetLegalContactID

`func (o *CustomerResponse) GetLegalContactID() string`

GetLegalContactID returns the LegalContactID field if non-nil, zero value otherwise.

### GetLegalContactIDOk

`func (o *CustomerResponse) GetLegalContactIDOk() (*string, bool)`

GetLegalContactIDOk returns a tuple with the LegalContactID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalContactID

`func (o *CustomerResponse) SetLegalContactID(v string)`

SetLegalContactID sets LegalContactID field to given value.

### HasLegalContactID

`func (o *CustomerResponse) HasLegalContactID() bool`

HasLegalContactID returns a boolean if a field has been set.

### SetLegalContactIDNil

`func (o *CustomerResponse) SetLegalContactIDNil(b bool)`

 SetLegalContactIDNil sets the value for LegalContactID to be an explicit nil

### UnsetLegalContactID
`func (o *CustomerResponse) UnsetLegalContactID()`

UnsetLegalContactID ensures that no value is present for LegalContactID, not even an explicit nil
### GetName

`func (o *CustomerResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomerResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomerResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CustomerResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwnerID

`func (o *CustomerResponse) GetOwnerID() string`

GetOwnerID returns the OwnerID field if non-nil, zero value otherwise.

### GetOwnerIDOk

`func (o *CustomerResponse) GetOwnerIDOk() (*string, bool)`

GetOwnerIDOk returns a tuple with the OwnerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerID

`func (o *CustomerResponse) SetOwnerID(v string)`

SetOwnerID sets OwnerID field to given value.

### HasOwnerID

`func (o *CustomerResponse) HasOwnerID() bool`

HasOwnerID returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *CustomerResponse) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *CustomerResponse) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *CustomerResponse) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *CustomerResponse) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetPostalAddress

`func (o *CustomerResponse) GetPostalAddress() string`

GetPostalAddress returns the PostalAddress field if non-nil, zero value otherwise.

### GetPostalAddressOk

`func (o *CustomerResponse) GetPostalAddressOk() (*string, bool)`

GetPostalAddressOk returns a tuple with the PostalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalAddress

`func (o *CustomerResponse) SetPostalAddress(v string)`

SetPostalAddress sets PostalAddress field to given value.

### HasPostalAddress

`func (o *CustomerResponse) HasPostalAddress() bool`

HasPostalAddress returns a boolean if a field has been set.

### SetPostalAddressNil

`func (o *CustomerResponse) SetPostalAddressNil(b bool)`

 SetPostalAddressNil sets the value for PostalAddress to be an explicit nil

### UnsetPostalAddress
`func (o *CustomerResponse) UnsetPostalAddress()`

UnsetPostalAddress ensures that no value is present for PostalAddress, not even an explicit nil
### GetPricingPlan

`func (o *CustomerResponse) GetPricingPlan() string`

GetPricingPlan returns the PricingPlan field if non-nil, zero value otherwise.

### GetPricingPlanOk

`func (o *CustomerResponse) GetPricingPlanOk() (*string, bool)`

GetPricingPlanOk returns a tuple with the PricingPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingPlan

`func (o *CustomerResponse) SetPricingPlan(v string)`

SetPricingPlan sets PricingPlan field to given value.

### HasPricingPlan

`func (o *CustomerResponse) HasPricingPlan() bool`

HasPricingPlan returns a boolean if a field has been set.

### GetPricingPlanID

`func (o *CustomerResponse) GetPricingPlanID() string`

GetPricingPlanID returns the PricingPlanID field if non-nil, zero value otherwise.

### GetPricingPlanIDOk

`func (o *CustomerResponse) GetPricingPlanIDOk() (*string, bool)`

GetPricingPlanIDOk returns a tuple with the PricingPlanID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingPlanID

`func (o *CustomerResponse) SetPricingPlanID(v string)`

SetPricingPlanID sets PricingPlanID field to given value.

### HasPricingPlanID

`func (o *CustomerResponse) HasPricingPlanID() bool`

HasPricingPlanID returns a boolean if a field has been set.

### GetSecurityContactID

`func (o *CustomerResponse) GetSecurityContactID() string`

GetSecurityContactID returns the SecurityContactID field if non-nil, zero value otherwise.

### GetSecurityContactIDOk

`func (o *CustomerResponse) GetSecurityContactIDOk() (*string, bool)`

GetSecurityContactIDOk returns a tuple with the SecurityContactID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityContactID

`func (o *CustomerResponse) SetSecurityContactID(v string)`

SetSecurityContactID sets SecurityContactID field to given value.

### HasSecurityContactID

`func (o *CustomerResponse) HasSecurityContactID() bool`

HasSecurityContactID returns a boolean if a field has been set.

### SetSecurityContactIDNil

`func (o *CustomerResponse) SetSecurityContactIDNil(b bool)`

 SetSecurityContactIDNil sets the value for SecurityContactID to be an explicit nil

### UnsetSecurityContactID
`func (o *CustomerResponse) UnsetSecurityContactID()`

UnsetSecurityContactID ensures that no value is present for SecurityContactID, not even an explicit nil
### GetTechnicalContactID

`func (o *CustomerResponse) GetTechnicalContactID() string`

GetTechnicalContactID returns the TechnicalContactID field if non-nil, zero value otherwise.

### GetTechnicalContactIDOk

`func (o *CustomerResponse) GetTechnicalContactIDOk() (*string, bool)`

GetTechnicalContactIDOk returns a tuple with the TechnicalContactID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechnicalContactID

`func (o *CustomerResponse) SetTechnicalContactID(v string)`

SetTechnicalContactID sets TechnicalContactID field to given value.

### HasTechnicalContactID

`func (o *CustomerResponse) HasTechnicalContactID() bool`

HasTechnicalContactID returns a boolean if a field has been set.

### SetTechnicalContactIDNil

`func (o *CustomerResponse) SetTechnicalContactIDNil(b bool)`

 SetTechnicalContactIDNil sets the value for TechnicalContactID to be an explicit nil

### UnsetTechnicalContactID
`func (o *CustomerResponse) UnsetTechnicalContactID()`

UnsetTechnicalContactID ensures that no value is present for TechnicalContactID, not even an explicit nil
### GetCreatedAt

`func (o *CustomerResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CustomerResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CustomerResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CustomerResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *CustomerResponse) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *CustomerResponse) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetDeletedAt

`func (o *CustomerResponse) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *CustomerResponse) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *CustomerResponse) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *CustomerResponse) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *CustomerResponse) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *CustomerResponse) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetUpdatedAt

`func (o *CustomerResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CustomerResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CustomerResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CustomerResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *CustomerResponse) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *CustomerResponse) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetID

`func (o *CustomerResponse) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *CustomerResponse) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *CustomerResponse) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *CustomerResponse) HasID() bool`

HasID returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
