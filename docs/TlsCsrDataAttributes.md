# TlsCsrDataAttributes

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
**RelationshipsTlsPrivateKeyId** | Pointer to **string** | Optional. An alphanumeric string identifying the private key you&#39;ve uploaded for use with your TLS certificate. If left blank, Fastly will create and manage a key for you. | [optional] 

## Methods

### NewTlsCsrDataAttributes

`func NewTlsCsrDataAttributes(sans []string, ) *TlsCsrDataAttributes`

NewTlsCsrDataAttributes instantiates a new TlsCsrDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTlsCsrDataAttributesWithDefaults

`func NewTlsCsrDataAttributesWithDefaults() *TlsCsrDataAttributes`

NewTlsCsrDataAttributesWithDefaults instantiates a new TlsCsrDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSans

`func (o *TlsCsrDataAttributes) GetSans() []string`

GetSans returns the Sans field if non-nil, zero value otherwise.

### GetSansOk

`func (o *TlsCsrDataAttributes) GetSansOk() (*[]string, bool)`

GetSansOk returns a tuple with the Sans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSans

`func (o *TlsCsrDataAttributes) SetSans(v []string)`

SetSans sets Sans field to given value.


### GetCommonName

`func (o *TlsCsrDataAttributes) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *TlsCsrDataAttributes) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *TlsCsrDataAttributes) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.

### HasCommonName

`func (o *TlsCsrDataAttributes) HasCommonName() bool`

HasCommonName returns a boolean if a field has been set.

### GetCountry

`func (o *TlsCsrDataAttributes) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *TlsCsrDataAttributes) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *TlsCsrDataAttributes) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *TlsCsrDataAttributes) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetState

`func (o *TlsCsrDataAttributes) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *TlsCsrDataAttributes) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *TlsCsrDataAttributes) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *TlsCsrDataAttributes) HasState() bool`

HasState returns a boolean if a field has been set.

### GetCity

`func (o *TlsCsrDataAttributes) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *TlsCsrDataAttributes) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *TlsCsrDataAttributes) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *TlsCsrDataAttributes) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetPostalCode

`func (o *TlsCsrDataAttributes) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *TlsCsrDataAttributes) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *TlsCsrDataAttributes) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *TlsCsrDataAttributes) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetStreetAddress

`func (o *TlsCsrDataAttributes) GetStreetAddress() string`

GetStreetAddress returns the StreetAddress field if non-nil, zero value otherwise.

### GetStreetAddressOk

`func (o *TlsCsrDataAttributes) GetStreetAddressOk() (*string, bool)`

GetStreetAddressOk returns a tuple with the StreetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetAddress

`func (o *TlsCsrDataAttributes) SetStreetAddress(v string)`

SetStreetAddress sets StreetAddress field to given value.

### HasStreetAddress

`func (o *TlsCsrDataAttributes) HasStreetAddress() bool`

HasStreetAddress returns a boolean if a field has been set.

### GetOrganization

`func (o *TlsCsrDataAttributes) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *TlsCsrDataAttributes) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *TlsCsrDataAttributes) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *TlsCsrDataAttributes) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetOrganizationalUnit

`func (o *TlsCsrDataAttributes) GetOrganizationalUnit() string`

GetOrganizationalUnit returns the OrganizationalUnit field if non-nil, zero value otherwise.

### GetOrganizationalUnitOk

`func (o *TlsCsrDataAttributes) GetOrganizationalUnitOk() (*string, bool)`

GetOrganizationalUnitOk returns a tuple with the OrganizationalUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationalUnit

`func (o *TlsCsrDataAttributes) SetOrganizationalUnit(v string)`

SetOrganizationalUnit sets OrganizationalUnit field to given value.

### HasOrganizationalUnit

`func (o *TlsCsrDataAttributes) HasOrganizationalUnit() bool`

HasOrganizationalUnit returns a boolean if a field has been set.

### GetEmail

`func (o *TlsCsrDataAttributes) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *TlsCsrDataAttributes) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *TlsCsrDataAttributes) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *TlsCsrDataAttributes) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetKeyType

`func (o *TlsCsrDataAttributes) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *TlsCsrDataAttributes) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *TlsCsrDataAttributes) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.

### HasKeyType

`func (o *TlsCsrDataAttributes) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.

### GetRelationshipsTlsPrivateKeyId

`func (o *TlsCsrDataAttributes) GetRelationshipsTlsPrivateKeyId() string`

GetRelationshipsTlsPrivateKeyId returns the RelationshipsTlsPrivateKeyId field if non-nil, zero value otherwise.

### GetRelationshipsTlsPrivateKeyIdOk

`func (o *TlsCsrDataAttributes) GetRelationshipsTlsPrivateKeyIdOk() (*string, bool)`

GetRelationshipsTlsPrivateKeyIdOk returns a tuple with the RelationshipsTlsPrivateKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationshipsTlsPrivateKeyId

`func (o *TlsCsrDataAttributes) SetRelationshipsTlsPrivateKeyId(v string)`

SetRelationshipsTlsPrivateKeyId sets RelationshipsTlsPrivateKeyId field to given value.

### HasRelationshipsTlsPrivateKeyId

`func (o *TlsCsrDataAttributes) HasRelationshipsTlsPrivateKeyId() bool`

HasRelationshipsTlsPrivateKeyId returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


