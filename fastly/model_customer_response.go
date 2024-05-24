// Package fastly is an API client library for interacting with most facets of the Fastly API.
package fastly

/*
Fastly API

Via the Fastly API you can perform any of the operations that are possible within the management console,  including creating services, domains, and backends, configuring rules or uploading your own application code, as well as account operations such as user administration and billing reports. The API is organized into collections of endpoints that allow manipulation of objects related to Fastly services and accounts. For the most accurate and up-to-date API reference content, visit our [Developer Hub](https://www.fastly.com/documentation/reference/api/) 

API version: 1.0.0
Contact: oss@fastly.com
*/

// This code is auto-generated; DO NOT EDIT.


import (
	"encoding/json"
	"time"
)

// CustomerResponse struct for CustomerResponse
type CustomerResponse struct {
	// The alphanumeric string representing the primary billing contact.
	BillingContactID NullableString `json:"billing_contact_id,omitempty"`
	// Customer's current network revenue type.
	BillingNetworkType *string `json:"billing_network_type,omitempty"`
	// Used for adding purchased orders to customer's account.
	BillingRef NullableString `json:"billing_ref,omitempty"`
	// Whether this customer can view or edit wordpress.
	CanConfigureWordpress NullableBool `json:"can_configure_wordpress,omitempty"`
	// Whether this customer can reset passwords.
	CanResetPasswords *bool `json:"can_reset_passwords,omitempty"`
	// Whether this customer can upload VCL.
	CanUploadVcl *bool `json:"can_upload_vcl,omitempty"`
	// Specifies whether 2FA is forced or not forced on the customer account. Logs out non-2FA users once 2FA is force enabled.
	Force2fa *bool `json:"force_2fa,omitempty"`
	// Specifies whether SSO is forced or not forced on the customer account.
	ForceSso *bool `json:"force_sso,omitempty"`
	// Specifies whether the account has access or does not have access to the account panel.
	HasAccountPanel *bool `json:"has_account_panel,omitempty"`
	// Specifies whether the account has access or does not have access to the improved events.
	HasImprovedEvents *bool `json:"has_improved_events,omitempty"`
	// Whether this customer can view or edit the SSL config.
	HasImprovedSslConfig *bool `json:"has_improved_ssl_config,omitempty"`
	// Specifies whether the account has enabled or not enabled openstack logging.
	HasOpenstackLogging *bool `json:"has_openstack_logging,omitempty"`
	// Specifies whether the account can edit PCI for a service.
	HasPci *bool `json:"has_pci,omitempty"`
	// Specifies whether PCI passwords are required for the account.
	HasPciPasswords *bool `json:"has_pci_passwords,omitempty"`
	// The range of IP addresses authorized to access the customer account.
	IPWhitelist *string `json:"ip_whitelist,omitempty"`
	// The alphanumeric string identifying the account's legal contact.
	LegalContactID NullableString `json:"legal_contact_id,omitempty"`
	// The name of the customer, generally the company name.
	Name *string `json:"name,omitempty"`
	// The alphanumeric string identifying the account owner.
	OwnerID *string `json:"owner_id,omitempty"`
	// The phone number associated with the account.
	PhoneNumber *string `json:"phone_number,omitempty"`
	// The postal address associated with the account.
	PostalAddress NullableString `json:"postal_address,omitempty"`
	// The pricing plan this customer is under.
	PricingPlan *string `json:"pricing_plan,omitempty"`
	// The alphanumeric string identifying the pricing plan.
	PricingPlanID *string `json:"pricing_plan_id,omitempty"`
	// The alphanumeric string identifying the account's security contact.
	SecurityContactID NullableString `json:"security_contact_id,omitempty"`
	// The alphanumeric string identifying the account's technical contact.
	TechnicalContactID NullableString `json:"technical_contact_id,omitempty"`
	// Date and time in ISO 8601 format.
	CreatedAt NullableTime `json:"created_at,omitempty"`
	// Date and time in ISO 8601 format.
	DeletedAt NullableTime `json:"deleted_at,omitempty"`
	// Date and time in ISO 8601 format.
	UpdatedAt NullableTime `json:"updated_at,omitempty"`
	ID *string `json:"id,omitempty"`
	AdditionalProperties map[string]any
}

type _CustomerResponse CustomerResponse

// NewCustomerResponse instantiates a new CustomerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerResponse() *CustomerResponse {
	this := CustomerResponse{}
	return &this
}

// NewCustomerResponseWithDefaults instantiates a new CustomerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerResponseWithDefaults() *CustomerResponse {
	this := CustomerResponse{}
	return &this
}

// GetBillingContactID returns the BillingContactID field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerResponse) GetBillingContactID() string {
	if o == nil || o.BillingContactID.Get() == nil {
		var ret string
		return ret
	}
	return *o.BillingContactID.Get()
}

// GetBillingContactIDOk returns a tuple with the BillingContactID field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerResponse) GetBillingContactIDOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.BillingContactID.Get(), o.BillingContactID.IsSet()
}

// HasBillingContactID returns a boolean if a field has been set.
func (o *CustomerResponse) HasBillingContactID() bool {
	if o != nil && o.BillingContactID.IsSet() {
		return true
	}

	return false
}

// SetBillingContactID gets a reference to the given NullableString and assigns it to the BillingContactID field.
func (o *CustomerResponse) SetBillingContactID(v string) {
	o.BillingContactID.Set(&v)
}
// SetBillingContactIDNil sets the value for BillingContactID to be an explicit nil
func (o *CustomerResponse) SetBillingContactIDNil() {
	o.BillingContactID.Set(nil)
}

// UnsetBillingContactID ensures that no value is present for BillingContactID, not even an explicit nil
func (o *CustomerResponse) UnsetBillingContactID() {
	o.BillingContactID.Unset()
}

// GetBillingNetworkType returns the BillingNetworkType field value if set, zero value otherwise.
func (o *CustomerResponse) GetBillingNetworkType() string {
	if o == nil || o.BillingNetworkType == nil {
		var ret string
		return ret
	}
	return *o.BillingNetworkType
}

// GetBillingNetworkTypeOk returns a tuple with the BillingNetworkType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerResponse) GetBillingNetworkTypeOk() (*string, bool) {
	if o == nil || o.BillingNetworkType == nil {
		return nil, false
	}
	return o.BillingNetworkType, true
}

// HasBillingNetworkType returns a boolean if a field has been set.
func (o *CustomerResponse) HasBillingNetworkType() bool {
	if o != nil && o.BillingNetworkType != nil {
		return true
	}

	return false
}

// SetBillingNetworkType gets a reference to the given string and assigns it to the BillingNetworkType field.
func (o *CustomerResponse) SetBillingNetworkType(v string) {
	o.BillingNetworkType = &v
}

// GetBillingRef returns the BillingRef field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerResponse) GetBillingRef() string {
	if o == nil || o.BillingRef.Get() == nil {
		var ret string
		return ret
	}
	return *o.BillingRef.Get()
}

// GetBillingRefOk returns a tuple with the BillingRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerResponse) GetBillingRefOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.BillingRef.Get(), o.BillingRef.IsSet()
}

// HasBillingRef returns a boolean if a field has been set.
func (o *CustomerResponse) HasBillingRef() bool {
	if o != nil && o.BillingRef.IsSet() {
		return true
	}

	return false
}

// SetBillingRef gets a reference to the given NullableString and assigns it to the BillingRef field.
func (o *CustomerResponse) SetBillingRef(v string) {
	o.BillingRef.Set(&v)
}
// SetBillingRefNil sets the value for BillingRef to be an explicit nil
func (o *CustomerResponse) SetBillingRefNil() {
	o.BillingRef.Set(nil)
}

// UnsetBillingRef ensures that no value is present for BillingRef, not even an explicit nil
func (o *CustomerResponse) UnsetBillingRef() {
	o.BillingRef.Unset()
}

// GetCanConfigureWordpress returns the CanConfigureWordpress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerResponse) GetCanConfigureWordpress() bool {
	if o == nil || o.CanConfigureWordpress.Get() == nil {
		var ret bool
		return ret
	}
	return *o.CanConfigureWordpress.Get()
}

// GetCanConfigureWordpressOk returns a tuple with the CanConfigureWordpress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerResponse) GetCanConfigureWordpressOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CanConfigureWordpress.Get(), o.CanConfigureWordpress.IsSet()
}

// HasCanConfigureWordpress returns a boolean if a field has been set.
func (o *CustomerResponse) HasCanConfigureWordpress() bool {
	if o != nil && o.CanConfigureWordpress.IsSet() {
		return true
	}

	return false
}

// SetCanConfigureWordpress gets a reference to the given NullableBool and assigns it to the CanConfigureWordpress field.
func (o *CustomerResponse) SetCanConfigureWordpress(v bool) {
	o.CanConfigureWordpress.Set(&v)
}
// SetCanConfigureWordpressNil sets the value for CanConfigureWordpress to be an explicit nil
func (o *CustomerResponse) SetCanConfigureWordpressNil() {
	o.CanConfigureWordpress.Set(nil)
}

// UnsetCanConfigureWordpress ensures that no value is present for CanConfigureWordpress, not even an explicit nil
func (o *CustomerResponse) UnsetCanConfigureWordpress() {
	o.CanConfigureWordpress.Unset()
}

// GetCanResetPasswords returns the CanResetPasswords field value if set, zero value otherwise.
func (o *CustomerResponse) GetCanResetPasswords() bool {
	if o == nil || o.CanResetPasswords == nil {
		var ret bool
		return ret
	}
	return *o.CanResetPasswords
}

// GetCanResetPasswordsOk returns a tuple with the CanResetPasswords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerResponse) GetCanResetPasswordsOk() (*bool, bool) {
	if o == nil || o.CanResetPasswords == nil {
		return nil, false
	}
	return o.CanResetPasswords, true
}

// HasCanResetPasswords returns a boolean if a field has been set.
func (o *CustomerResponse) HasCanResetPasswords() bool {
	if o != nil && o.CanResetPasswords != nil {
		return true
	}

	return false
}

// SetCanResetPasswords gets a reference to the given bool and assigns it to the CanResetPasswords field.
func (o *CustomerResponse) SetCanResetPasswords(v bool) {
	o.CanResetPasswords = &v
}

// GetCanUploadVcl returns the CanUploadVcl field value if set, zero value otherwise.
func (o *CustomerResponse) GetCanUploadVcl() bool {
	if o == nil || o.CanUploadVcl == nil {
		var ret bool
		return ret
	}
	return *o.CanUploadVcl
}

// GetCanUploadVclOk returns a tuple with the CanUploadVcl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerResponse) GetCanUploadVclOk() (*bool, bool) {
	if o == nil || o.CanUploadVcl == nil {
		return nil, false
	}
	return o.CanUploadVcl, true
}

// HasCanUploadVcl returns a boolean if a field has been set.
func (o *CustomerResponse) HasCanUploadVcl() bool {
	if o != nil && o.CanUploadVcl != nil {
		return true
	}

	return false
}

// SetCanUploadVcl gets a reference to the given bool and assigns it to the CanUploadVcl field.
func (o *CustomerResponse) SetCanUploadVcl(v bool) {
	o.CanUploadVcl = &v
}

// GetForce2fa returns the Force2fa field value if set, zero value otherwise.
func (o *CustomerResponse) GetForce2fa() bool {
	if o == nil || o.Force2fa == nil {
		var ret bool
		return ret
	}
	return *o.Force2fa
}

// GetForce2faOk returns a tuple with the Force2fa field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerResponse) GetForce2faOk() (*bool, bool) {
	if o == nil || o.Force2fa == nil {
		return nil, false
	}
	return o.Force2fa, true
}

// HasForce2fa returns a boolean if a field has been set.
func (o *CustomerResponse) HasForce2fa() bool {
	if o != nil && o.Force2fa != nil {
		return true
	}

	return false
}

// SetForce2fa gets a reference to the given bool and assigns it to the Force2fa field.
func (o *CustomerResponse) SetForce2fa(v bool) {
	o.Force2fa = &v
}

// GetForceSso returns the ForceSso field value if set, zero value otherwise.
func (o *CustomerResponse) GetForceSso() bool {
	if o == nil || o.ForceSso == nil {
		var ret bool
		return ret
	}
	return *o.ForceSso
}

// GetForceSsoOk returns a tuple with the ForceSso field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerResponse) GetForceSsoOk() (*bool, bool) {
	if o == nil || o.ForceSso == nil {
		return nil, false
	}
	return o.ForceSso, true
}

// HasForceSso returns a boolean if a field has been set.
func (o *CustomerResponse) HasForceSso() bool {
	if o != nil && o.ForceSso != nil {
		return true
	}

	return false
}

// SetForceSso gets a reference to the given bool and assigns it to the ForceSso field.
func (o *CustomerResponse) SetForceSso(v bool) {
	o.ForceSso = &v
}

// GetHasAccountPanel returns the HasAccountPanel field value if set, zero value otherwise.
func (o *CustomerResponse) GetHasAccountPanel() bool {
	if o == nil || o.HasAccountPanel == nil {
		var ret bool
		return ret
	}
	return *o.HasAccountPanel
}

// GetHasAccountPanelOk returns a tuple with the HasAccountPanel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerResponse) GetHasAccountPanelOk() (*bool, bool) {
	if o == nil || o.HasAccountPanel == nil {
		return nil, false
	}
	return o.HasAccountPanel, true
}

// HasHasAccountPanel returns a boolean if a field has been set.
func (o *CustomerResponse) HasHasAccountPanel() bool {
	if o != nil && o.HasAccountPanel != nil {
		return true
	}

	return false
}

// SetHasAccountPanel gets a reference to the given bool and assigns it to the HasAccountPanel field.
func (o *CustomerResponse) SetHasAccountPanel(v bool) {
	o.HasAccountPanel = &v
}

// GetHasImprovedEvents returns the HasImprovedEvents field value if set, zero value otherwise.
func (o *CustomerResponse) GetHasImprovedEvents() bool {
	if o == nil || o.HasImprovedEvents == nil {
		var ret bool
		return ret
	}
	return *o.HasImprovedEvents
}

// GetHasImprovedEventsOk returns a tuple with the HasImprovedEvents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerResponse) GetHasImprovedEventsOk() (*bool, bool) {
	if o == nil || o.HasImprovedEvents == nil {
		return nil, false
	}
	return o.HasImprovedEvents, true
}

// HasHasImprovedEvents returns a boolean if a field has been set.
func (o *CustomerResponse) HasHasImprovedEvents() bool {
	if o != nil && o.HasImprovedEvents != nil {
		return true
	}

	return false
}

// SetHasImprovedEvents gets a reference to the given bool and assigns it to the HasImprovedEvents field.
func (o *CustomerResponse) SetHasImprovedEvents(v bool) {
	o.HasImprovedEvents = &v
}

// GetHasImprovedSslConfig returns the HasImprovedSslConfig field value if set, zero value otherwise.
func (o *CustomerResponse) GetHasImprovedSslConfig() bool {
	if o == nil || o.HasImprovedSslConfig == nil {
		var ret bool
		return ret
	}
	return *o.HasImprovedSslConfig
}

// GetHasImprovedSslConfigOk returns a tuple with the HasImprovedSslConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerResponse) GetHasImprovedSslConfigOk() (*bool, bool) {
	if o == nil || o.HasImprovedSslConfig == nil {
		return nil, false
	}
	return o.HasImprovedSslConfig, true
}

// HasHasImprovedSslConfig returns a boolean if a field has been set.
func (o *CustomerResponse) HasHasImprovedSslConfig() bool {
	if o != nil && o.HasImprovedSslConfig != nil {
		return true
	}

	return false
}

// SetHasImprovedSslConfig gets a reference to the given bool and assigns it to the HasImprovedSslConfig field.
func (o *CustomerResponse) SetHasImprovedSslConfig(v bool) {
	o.HasImprovedSslConfig = &v
}

// GetHasOpenstackLogging returns the HasOpenstackLogging field value if set, zero value otherwise.
func (o *CustomerResponse) GetHasOpenstackLogging() bool {
	if o == nil || o.HasOpenstackLogging == nil {
		var ret bool
		return ret
	}
	return *o.HasOpenstackLogging
}

// GetHasOpenstackLoggingOk returns a tuple with the HasOpenstackLogging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerResponse) GetHasOpenstackLoggingOk() (*bool, bool) {
	if o == nil || o.HasOpenstackLogging == nil {
		return nil, false
	}
	return o.HasOpenstackLogging, true
}

// HasHasOpenstackLogging returns a boolean if a field has been set.
func (o *CustomerResponse) HasHasOpenstackLogging() bool {
	if o != nil && o.HasOpenstackLogging != nil {
		return true
	}

	return false
}

// SetHasOpenstackLogging gets a reference to the given bool and assigns it to the HasOpenstackLogging field.
func (o *CustomerResponse) SetHasOpenstackLogging(v bool) {
	o.HasOpenstackLogging = &v
}

// GetHasPci returns the HasPci field value if set, zero value otherwise.
func (o *CustomerResponse) GetHasPci() bool {
	if o == nil || o.HasPci == nil {
		var ret bool
		return ret
	}
	return *o.HasPci
}

// GetHasPciOk returns a tuple with the HasPci field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerResponse) GetHasPciOk() (*bool, bool) {
	if o == nil || o.HasPci == nil {
		return nil, false
	}
	return o.HasPci, true
}

// HasHasPci returns a boolean if a field has been set.
func (o *CustomerResponse) HasHasPci() bool {
	if o != nil && o.HasPci != nil {
		return true
	}

	return false
}

// SetHasPci gets a reference to the given bool and assigns it to the HasPci field.
func (o *CustomerResponse) SetHasPci(v bool) {
	o.HasPci = &v
}

// GetHasPciPasswords returns the HasPciPasswords field value if set, zero value otherwise.
func (o *CustomerResponse) GetHasPciPasswords() bool {
	if o == nil || o.HasPciPasswords == nil {
		var ret bool
		return ret
	}
	return *o.HasPciPasswords
}

// GetHasPciPasswordsOk returns a tuple with the HasPciPasswords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerResponse) GetHasPciPasswordsOk() (*bool, bool) {
	if o == nil || o.HasPciPasswords == nil {
		return nil, false
	}
	return o.HasPciPasswords, true
}

// HasHasPciPasswords returns a boolean if a field has been set.
func (o *CustomerResponse) HasHasPciPasswords() bool {
	if o != nil && o.HasPciPasswords != nil {
		return true
	}

	return false
}

// SetHasPciPasswords gets a reference to the given bool and assigns it to the HasPciPasswords field.
func (o *CustomerResponse) SetHasPciPasswords(v bool) {
	o.HasPciPasswords = &v
}

// GetIPWhitelist returns the IPWhitelist field value if set, zero value otherwise.
func (o *CustomerResponse) GetIPWhitelist() string {
	if o == nil || o.IPWhitelist == nil {
		var ret string
		return ret
	}
	return *o.IPWhitelist
}

// GetIPWhitelistOk returns a tuple with the IPWhitelist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerResponse) GetIPWhitelistOk() (*string, bool) {
	if o == nil || o.IPWhitelist == nil {
		return nil, false
	}
	return o.IPWhitelist, true
}

// HasIPWhitelist returns a boolean if a field has been set.
func (o *CustomerResponse) HasIPWhitelist() bool {
	if o != nil && o.IPWhitelist != nil {
		return true
	}

	return false
}

// SetIPWhitelist gets a reference to the given string and assigns it to the IPWhitelist field.
func (o *CustomerResponse) SetIPWhitelist(v string) {
	o.IPWhitelist = &v
}

// GetLegalContactID returns the LegalContactID field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerResponse) GetLegalContactID() string {
	if o == nil || o.LegalContactID.Get() == nil {
		var ret string
		return ret
	}
	return *o.LegalContactID.Get()
}

// GetLegalContactIDOk returns a tuple with the LegalContactID field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerResponse) GetLegalContactIDOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LegalContactID.Get(), o.LegalContactID.IsSet()
}

// HasLegalContactID returns a boolean if a field has been set.
func (o *CustomerResponse) HasLegalContactID() bool {
	if o != nil && o.LegalContactID.IsSet() {
		return true
	}

	return false
}

// SetLegalContactID gets a reference to the given NullableString and assigns it to the LegalContactID field.
func (o *CustomerResponse) SetLegalContactID(v string) {
	o.LegalContactID.Set(&v)
}
// SetLegalContactIDNil sets the value for LegalContactID to be an explicit nil
func (o *CustomerResponse) SetLegalContactIDNil() {
	o.LegalContactID.Set(nil)
}

// UnsetLegalContactID ensures that no value is present for LegalContactID, not even an explicit nil
func (o *CustomerResponse) UnsetLegalContactID() {
	o.LegalContactID.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CustomerResponse) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerResponse) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CustomerResponse) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CustomerResponse) SetName(v string) {
	o.Name = &v
}

// GetOwnerID returns the OwnerID field value if set, zero value otherwise.
func (o *CustomerResponse) GetOwnerID() string {
	if o == nil || o.OwnerID == nil {
		var ret string
		return ret
	}
	return *o.OwnerID
}

// GetOwnerIDOk returns a tuple with the OwnerID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerResponse) GetOwnerIDOk() (*string, bool) {
	if o == nil || o.OwnerID == nil {
		return nil, false
	}
	return o.OwnerID, true
}

// HasOwnerID returns a boolean if a field has been set.
func (o *CustomerResponse) HasOwnerID() bool {
	if o != nil && o.OwnerID != nil {
		return true
	}

	return false
}

// SetOwnerID gets a reference to the given string and assigns it to the OwnerID field.
func (o *CustomerResponse) SetOwnerID(v string) {
	o.OwnerID = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *CustomerResponse) GetPhoneNumber() string {
	if o == nil || o.PhoneNumber == nil {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerResponse) GetPhoneNumberOk() (*string, bool) {
	if o == nil || o.PhoneNumber == nil {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *CustomerResponse) HasPhoneNumber() bool {
	if o != nil && o.PhoneNumber != nil {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *CustomerResponse) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetPostalAddress returns the PostalAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerResponse) GetPostalAddress() string {
	if o == nil || o.PostalAddress.Get() == nil {
		var ret string
		return ret
	}
	return *o.PostalAddress.Get()
}

// GetPostalAddressOk returns a tuple with the PostalAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerResponse) GetPostalAddressOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PostalAddress.Get(), o.PostalAddress.IsSet()
}

// HasPostalAddress returns a boolean if a field has been set.
func (o *CustomerResponse) HasPostalAddress() bool {
	if o != nil && o.PostalAddress.IsSet() {
		return true
	}

	return false
}

// SetPostalAddress gets a reference to the given NullableString and assigns it to the PostalAddress field.
func (o *CustomerResponse) SetPostalAddress(v string) {
	o.PostalAddress.Set(&v)
}
// SetPostalAddressNil sets the value for PostalAddress to be an explicit nil
func (o *CustomerResponse) SetPostalAddressNil() {
	o.PostalAddress.Set(nil)
}

// UnsetPostalAddress ensures that no value is present for PostalAddress, not even an explicit nil
func (o *CustomerResponse) UnsetPostalAddress() {
	o.PostalAddress.Unset()
}

// GetPricingPlan returns the PricingPlan field value if set, zero value otherwise.
func (o *CustomerResponse) GetPricingPlan() string {
	if o == nil || o.PricingPlan == nil {
		var ret string
		return ret
	}
	return *o.PricingPlan
}

// GetPricingPlanOk returns a tuple with the PricingPlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerResponse) GetPricingPlanOk() (*string, bool) {
	if o == nil || o.PricingPlan == nil {
		return nil, false
	}
	return o.PricingPlan, true
}

// HasPricingPlan returns a boolean if a field has been set.
func (o *CustomerResponse) HasPricingPlan() bool {
	if o != nil && o.PricingPlan != nil {
		return true
	}

	return false
}

// SetPricingPlan gets a reference to the given string and assigns it to the PricingPlan field.
func (o *CustomerResponse) SetPricingPlan(v string) {
	o.PricingPlan = &v
}

// GetPricingPlanID returns the PricingPlanID field value if set, zero value otherwise.
func (o *CustomerResponse) GetPricingPlanID() string {
	if o == nil || o.PricingPlanID == nil {
		var ret string
		return ret
	}
	return *o.PricingPlanID
}

// GetPricingPlanIDOk returns a tuple with the PricingPlanID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerResponse) GetPricingPlanIDOk() (*string, bool) {
	if o == nil || o.PricingPlanID == nil {
		return nil, false
	}
	return o.PricingPlanID, true
}

// HasPricingPlanID returns a boolean if a field has been set.
func (o *CustomerResponse) HasPricingPlanID() bool {
	if o != nil && o.PricingPlanID != nil {
		return true
	}

	return false
}

// SetPricingPlanID gets a reference to the given string and assigns it to the PricingPlanID field.
func (o *CustomerResponse) SetPricingPlanID(v string) {
	o.PricingPlanID = &v
}

// GetSecurityContactID returns the SecurityContactID field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerResponse) GetSecurityContactID() string {
	if o == nil || o.SecurityContactID.Get() == nil {
		var ret string
		return ret
	}
	return *o.SecurityContactID.Get()
}

// GetSecurityContactIDOk returns a tuple with the SecurityContactID field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerResponse) GetSecurityContactIDOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SecurityContactID.Get(), o.SecurityContactID.IsSet()
}

// HasSecurityContactID returns a boolean if a field has been set.
func (o *CustomerResponse) HasSecurityContactID() bool {
	if o != nil && o.SecurityContactID.IsSet() {
		return true
	}

	return false
}

// SetSecurityContactID gets a reference to the given NullableString and assigns it to the SecurityContactID field.
func (o *CustomerResponse) SetSecurityContactID(v string) {
	o.SecurityContactID.Set(&v)
}
// SetSecurityContactIDNil sets the value for SecurityContactID to be an explicit nil
func (o *CustomerResponse) SetSecurityContactIDNil() {
	o.SecurityContactID.Set(nil)
}

// UnsetSecurityContactID ensures that no value is present for SecurityContactID, not even an explicit nil
func (o *CustomerResponse) UnsetSecurityContactID() {
	o.SecurityContactID.Unset()
}

// GetTechnicalContactID returns the TechnicalContactID field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerResponse) GetTechnicalContactID() string {
	if o == nil || o.TechnicalContactID.Get() == nil {
		var ret string
		return ret
	}
	return *o.TechnicalContactID.Get()
}

// GetTechnicalContactIDOk returns a tuple with the TechnicalContactID field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerResponse) GetTechnicalContactIDOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TechnicalContactID.Get(), o.TechnicalContactID.IsSet()
}

// HasTechnicalContactID returns a boolean if a field has been set.
func (o *CustomerResponse) HasTechnicalContactID() bool {
	if o != nil && o.TechnicalContactID.IsSet() {
		return true
	}

	return false
}

// SetTechnicalContactID gets a reference to the given NullableString and assigns it to the TechnicalContactID field.
func (o *CustomerResponse) SetTechnicalContactID(v string) {
	o.TechnicalContactID.Set(&v)
}
// SetTechnicalContactIDNil sets the value for TechnicalContactID to be an explicit nil
func (o *CustomerResponse) SetTechnicalContactIDNil() {
	o.TechnicalContactID.Set(nil)
}

// UnsetTechnicalContactID ensures that no value is present for TechnicalContactID, not even an explicit nil
func (o *CustomerResponse) UnsetTechnicalContactID() {
	o.TechnicalContactID.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerResponse) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *CustomerResponse) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableTime and assigns it to the CreatedAt field.
func (o *CustomerResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt.Set(&v)
}
// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *CustomerResponse) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *CustomerResponse) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerResponse) GetDeletedAt() time.Time {
	if o == nil || o.DeletedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt.Get()
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerResponse) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DeletedAt.Get(), o.DeletedAt.IsSet()
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *CustomerResponse) HasDeletedAt() bool {
	if o != nil && o.DeletedAt.IsSet() {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given NullableTime and assigns it to the DeletedAt field.
func (o *CustomerResponse) SetDeletedAt(v time.Time) {
	o.DeletedAt.Set(&v)
}
// SetDeletedAtNil sets the value for DeletedAt to be an explicit nil
func (o *CustomerResponse) SetDeletedAtNil() {
	o.DeletedAt.Set(nil)
}

// UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
func (o *CustomerResponse) UnsetDeletedAt() {
	o.DeletedAt.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerResponse) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *CustomerResponse) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given NullableTime and assigns it to the UpdatedAt field.
func (o *CustomerResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt.Set(&v)
}
// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *CustomerResponse) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *CustomerResponse) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *CustomerResponse) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerResponse) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *CustomerResponse) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *CustomerResponse) SetID(v string) {
	o.ID = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o CustomerResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.BillingContactID.IsSet() {
		toSerialize["billing_contact_id"] = o.BillingContactID.Get()
	}
	if o.BillingNetworkType != nil {
		toSerialize["billing_network_type"] = o.BillingNetworkType
	}
	if o.BillingRef.IsSet() {
		toSerialize["billing_ref"] = o.BillingRef.Get()
	}
	if o.CanConfigureWordpress.IsSet() {
		toSerialize["can_configure_wordpress"] = o.CanConfigureWordpress.Get()
	}
	if o.CanResetPasswords != nil {
		toSerialize["can_reset_passwords"] = o.CanResetPasswords
	}
	if o.CanUploadVcl != nil {
		toSerialize["can_upload_vcl"] = o.CanUploadVcl
	}
	if o.Force2fa != nil {
		toSerialize["force_2fa"] = o.Force2fa
	}
	if o.ForceSso != nil {
		toSerialize["force_sso"] = o.ForceSso
	}
	if o.HasAccountPanel != nil {
		toSerialize["has_account_panel"] = o.HasAccountPanel
	}
	if o.HasImprovedEvents != nil {
		toSerialize["has_improved_events"] = o.HasImprovedEvents
	}
	if o.HasImprovedSslConfig != nil {
		toSerialize["has_improved_ssl_config"] = o.HasImprovedSslConfig
	}
	if o.HasOpenstackLogging != nil {
		toSerialize["has_openstack_logging"] = o.HasOpenstackLogging
	}
	if o.HasPci != nil {
		toSerialize["has_pci"] = o.HasPci
	}
	if o.HasPciPasswords != nil {
		toSerialize["has_pci_passwords"] = o.HasPciPasswords
	}
	if o.IPWhitelist != nil {
		toSerialize["ip_whitelist"] = o.IPWhitelist
	}
	if o.LegalContactID.IsSet() {
		toSerialize["legal_contact_id"] = o.LegalContactID.Get()
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.OwnerID != nil {
		toSerialize["owner_id"] = o.OwnerID
	}
	if o.PhoneNumber != nil {
		toSerialize["phone_number"] = o.PhoneNumber
	}
	if o.PostalAddress.IsSet() {
		toSerialize["postal_address"] = o.PostalAddress.Get()
	}
	if o.PricingPlan != nil {
		toSerialize["pricing_plan"] = o.PricingPlan
	}
	if o.PricingPlanID != nil {
		toSerialize["pricing_plan_id"] = o.PricingPlanID
	}
	if o.SecurityContactID.IsSet() {
		toSerialize["security_contact_id"] = o.SecurityContactID.Get()
	}
	if o.TechnicalContactID.IsSet() {
		toSerialize["technical_contact_id"] = o.TechnicalContactID.Get()
	}
	if o.CreatedAt.IsSet() {
		toSerialize["created_at"] = o.CreatedAt.Get()
	}
	if o.DeletedAt.IsSet() {
		toSerialize["deleted_at"] = o.DeletedAt.Get()
	}
	if o.UpdatedAt.IsSet() {
		toSerialize["updated_at"] = o.UpdatedAt.Get()
	}
	if o.ID != nil {
		toSerialize["id"] = o.ID
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *CustomerResponse) UnmarshalJSON(bytes []byte) (err error) {
	varCustomerResponse := _CustomerResponse{}

	if err = json.Unmarshal(bytes, &varCustomerResponse); err == nil {
		*o = CustomerResponse(varCustomerResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "billing_contact_id")
		delete(additionalProperties, "billing_network_type")
		delete(additionalProperties, "billing_ref")
		delete(additionalProperties, "can_configure_wordpress")
		delete(additionalProperties, "can_reset_passwords")
		delete(additionalProperties, "can_upload_vcl")
		delete(additionalProperties, "force_2fa")
		delete(additionalProperties, "force_sso")
		delete(additionalProperties, "has_account_panel")
		delete(additionalProperties, "has_improved_events")
		delete(additionalProperties, "has_improved_ssl_config")
		delete(additionalProperties, "has_openstack_logging")
		delete(additionalProperties, "has_pci")
		delete(additionalProperties, "has_pci_passwords")
		delete(additionalProperties, "ip_whitelist")
		delete(additionalProperties, "legal_contact_id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "owner_id")
		delete(additionalProperties, "phone_number")
		delete(additionalProperties, "postal_address")
		delete(additionalProperties, "pricing_plan")
		delete(additionalProperties, "pricing_plan_id")
		delete(additionalProperties, "security_contact_id")
		delete(additionalProperties, "technical_contact_id")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "deleted_at")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableCustomerResponse is a helper abstraction for handling nullable customerresponse types. 
type NullableCustomerResponse struct {
	value *CustomerResponse
	isSet bool
}

// Get returns the value.
func (v NullableCustomerResponse) Get() *CustomerResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableCustomerResponse) Set(val *CustomerResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableCustomerResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableCustomerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableCustomerResponse returns a pointer to a new instance of NullableCustomerResponse.
func NewNullableCustomerResponse(val *CustomerResponse) *NullableCustomerResponse {
	return &NullableCustomerResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableCustomerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableCustomerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
