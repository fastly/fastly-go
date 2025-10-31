# CustomerAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of the address. | [optional] 
**Address1** | Pointer to **string** | The street number and name of the address. | [optional] 
**Address2** | Pointer to **string** | Additional address line for unit number, P.O. Box, etc. | [optional] 
**Locality** | Pointer to **string** | City, town, or locality name the address is located. | [optional] 
**Region** | Pointer to **string** | State, province, or region of the address. | [optional] 
**Country** | Pointer to **string** | ISO 3166-1 alpha-2 country code (e.g., \&quot;US\&quot;, \&quot;CA\&quot;, \&quot;NZ\&quot;) | [optional] 
**PostalCode** | Pointer to **string** | Postal or Zip code of the address. | [optional] 

## Methods

### NewCustomerAddress

`func NewCustomerAddress() *CustomerAddress`

NewCustomerAddress instantiates a new CustomerAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerAddressWithDefaults

`func NewCustomerAddressWithDefaults() *CustomerAddress`

NewCustomerAddressWithDefaults instantiates a new CustomerAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CustomerAddress) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomerAddress) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomerAddress) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CustomerAddress) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAddress1

`func (o *CustomerAddress) GetAddress1() string`

GetAddress1 returns the Address1 field if non-nil, zero value otherwise.

### GetAddress1Ok

`func (o *CustomerAddress) GetAddress1Ok() (*string, bool)`

GetAddress1Ok returns a tuple with the Address1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress1

`func (o *CustomerAddress) SetAddress1(v string)`

SetAddress1 sets Address1 field to given value.

### HasAddress1

`func (o *CustomerAddress) HasAddress1() bool`

HasAddress1 returns a boolean if a field has been set.

### GetAddress2

`func (o *CustomerAddress) GetAddress2() string`

GetAddress2 returns the Address2 field if non-nil, zero value otherwise.

### GetAddress2Ok

`func (o *CustomerAddress) GetAddress2Ok() (*string, bool)`

GetAddress2Ok returns a tuple with the Address2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress2

`func (o *CustomerAddress) SetAddress2(v string)`

SetAddress2 sets Address2 field to given value.

### HasAddress2

`func (o *CustomerAddress) HasAddress2() bool`

HasAddress2 returns a boolean if a field has been set.

### GetLocality

`func (o *CustomerAddress) GetLocality() string`

GetLocality returns the Locality field if non-nil, zero value otherwise.

### GetLocalityOk

`func (o *CustomerAddress) GetLocalityOk() (*string, bool)`

GetLocalityOk returns a tuple with the Locality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocality

`func (o *CustomerAddress) SetLocality(v string)`

SetLocality sets Locality field to given value.

### HasLocality

`func (o *CustomerAddress) HasLocality() bool`

HasLocality returns a boolean if a field has been set.

### GetRegion

`func (o *CustomerAddress) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CustomerAddress) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CustomerAddress) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *CustomerAddress) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetCountry

`func (o *CustomerAddress) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *CustomerAddress) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *CustomerAddress) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *CustomerAddress) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetPostalCode

`func (o *CustomerAddress) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *CustomerAddress) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *CustomerAddress) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *CustomerAddress) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


