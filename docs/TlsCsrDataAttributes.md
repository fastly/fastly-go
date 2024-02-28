# TLSCsrDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sans** | **[]string** | Subject Alternate Names - An array of one or more fully qualified domain names or public IP addresses to be secured by this certificate. Required. | 
**CommonName** | Pointer to **string** | Common Name (CN) - The fully qualified domain name (FQDN) to be secured by this certificate. The common name should be one of the entries in the SANs parameter. | [optional] 
**Country** | Pointer to **string** | Country (C) - The two-letter ISO country code where the organization is located. | [optional] 
**State** | Pointer to **string** | State (S) - The state, province, region, or county where the organization is located. This should not be abbreviated. | [optional] 
**City** | Pointer to **string** | Locality (L) - The locality, city, town, or village where the organization is located. | [optional] 
**PostalCode** | Pointer to **string** | Postal Code - The postal code where the organization is located. | [optional] 
**StreetAddress** | Pointer to **string** | Street Address - The street address where the organization is located. | [optional] 
**Organization** | Pointer to **string** | Organization (O) - The legal name of the organization, including any suffixes. This should not be abbreviated. | [optional] 
**OrganizationalUnit** | Pointer to **string** | Organizational Unit (OU) - The internal division of the organization managing the certificate. | [optional] 
**Email** | Pointer to **string** | Email Address (EMAIL) - The organizational contact for this. | [optional] 
**KeyType** | Pointer to **string** | CSR Key Type. | [optional] 

## Methods

### NewTLSCsrDataAttributes

`func NewTLSCsrDataAttributes(sans []string, ) *TLSCsrDataAttributes`

NewTLSCsrDataAttributes instantiates a new TLSCsrDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTLSCsrDataAttributesWithDefaults

`func NewTLSCsrDataAttributesWithDefaults() *TLSCsrDataAttributes`

NewTLSCsrDataAttributesWithDefaults instantiates a new TLSCsrDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSans

`func (o *TLSCsrDataAttributes) GetSans() []string`

GetSans returns the Sans field if non-nil, zero value otherwise.

### GetSansOk

`func (o *TLSCsrDataAttributes) GetSansOk() (*[]string, bool)`

GetSansOk returns a tuple with the Sans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSans

`func (o *TLSCsrDataAttributes) SetSans(v []string)`

SetSans sets Sans field to given value.


### GetCommonName

`func (o *TLSCsrDataAttributes) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *TLSCsrDataAttributes) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *TLSCsrDataAttributes) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.

### HasCommonName

`func (o *TLSCsrDataAttributes) HasCommonName() bool`

HasCommonName returns a boolean if a field has been set.

### GetCountry

`func (o *TLSCsrDataAttributes) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *TLSCsrDataAttributes) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *TLSCsrDataAttributes) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *TLSCsrDataAttributes) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetState

`func (o *TLSCsrDataAttributes) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *TLSCsrDataAttributes) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *TLSCsrDataAttributes) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *TLSCsrDataAttributes) HasState() bool`

HasState returns a boolean if a field has been set.

### GetCity

`func (o *TLSCsrDataAttributes) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *TLSCsrDataAttributes) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *TLSCsrDataAttributes) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *TLSCsrDataAttributes) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetPostalCode

`func (o *TLSCsrDataAttributes) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *TLSCsrDataAttributes) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *TLSCsrDataAttributes) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *TLSCsrDataAttributes) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetStreetAddress

`func (o *TLSCsrDataAttributes) GetStreetAddress() string`

GetStreetAddress returns the StreetAddress field if non-nil, zero value otherwise.

### GetStreetAddressOk

`func (o *TLSCsrDataAttributes) GetStreetAddressOk() (*string, bool)`

GetStreetAddressOk returns a tuple with the StreetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetAddress

`func (o *TLSCsrDataAttributes) SetStreetAddress(v string)`

SetStreetAddress sets StreetAddress field to given value.

### HasStreetAddress

`func (o *TLSCsrDataAttributes) HasStreetAddress() bool`

HasStreetAddress returns a boolean if a field has been set.

### GetOrganization

`func (o *TLSCsrDataAttributes) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *TLSCsrDataAttributes) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *TLSCsrDataAttributes) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *TLSCsrDataAttributes) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetOrganizationalUnit

`func (o *TLSCsrDataAttributes) GetOrganizationalUnit() string`

GetOrganizationalUnit returns the OrganizationalUnit field if non-nil, zero value otherwise.

### GetOrganizationalUnitOk

`func (o *TLSCsrDataAttributes) GetOrganizationalUnitOk() (*string, bool)`

GetOrganizationalUnitOk returns a tuple with the OrganizationalUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationalUnit

`func (o *TLSCsrDataAttributes) SetOrganizationalUnit(v string)`

SetOrganizationalUnit sets OrganizationalUnit field to given value.

### HasOrganizationalUnit

`func (o *TLSCsrDataAttributes) HasOrganizationalUnit() bool`

HasOrganizationalUnit returns a boolean if a field has been set.

### GetEmail

`func (o *TLSCsrDataAttributes) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *TLSCsrDataAttributes) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *TLSCsrDataAttributes) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *TLSCsrDataAttributes) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetKeyType

`func (o *TLSCsrDataAttributes) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *TLSCsrDataAttributes) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *TLSCsrDataAttributes) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.

### HasKeyType

`func (o *TLSCsrDataAttributes) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
