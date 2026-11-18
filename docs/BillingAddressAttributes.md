# BillingAddressAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address1** | Pointer to **string** | The first address line. | [optional] 
**Address2** | Pointer to **NullableString** | The second address line. | [optional] 
**City** | Pointer to **NullableString** | The city name. | [optional] 
**Country** | Pointer to **string** | ISO 3166-1 two-letter country code. | [optional] 
**Locality** | Pointer to **NullableString** | Other locality. | [optional] 
**PostalCode** | Pointer to **string** | Postal code (ZIP code for US addresses). | [optional] 
**State** | Pointer to **NullableString** | The state or province name. | [optional] 
**CustomerID** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewBillingAddressAttributes

`func NewBillingAddressAttributes() *BillingAddressAttributes`

NewBillingAddressAttributes instantiates a new BillingAddressAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingAddressAttributesWithDefaults

`func NewBillingAddressAttributesWithDefaults() *BillingAddressAttributes`

NewBillingAddressAttributesWithDefaults instantiates a new BillingAddressAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress1

`func (o *BillingAddressAttributes) GetAddress1() string`

GetAddress1 returns the Address1 field if non-nil, zero value otherwise.

### GetAddress1Ok

`func (o *BillingAddressAttributes) GetAddress1Ok() (*string, bool)`

GetAddress1Ok returns a tuple with the Address1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress1

`func (o *BillingAddressAttributes) SetAddress1(v string)`

SetAddress1 sets Address1 field to given value.

### HasAddress1

`func (o *BillingAddressAttributes) HasAddress1() bool`

HasAddress1 returns a boolean if a field has been set.

### GetAddress2

`func (o *BillingAddressAttributes) GetAddress2() string`

GetAddress2 returns the Address2 field if non-nil, zero value otherwise.

### GetAddress2Ok

`func (o *BillingAddressAttributes) GetAddress2Ok() (*string, bool)`

GetAddress2Ok returns a tuple with the Address2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress2

`func (o *BillingAddressAttributes) SetAddress2(v string)`

SetAddress2 sets Address2 field to given value.

### HasAddress2

`func (o *BillingAddressAttributes) HasAddress2() bool`

HasAddress2 returns a boolean if a field has been set.

### SetAddress2Nil

`func (o *BillingAddressAttributes) SetAddress2Nil(b bool)`

 SetAddress2Nil sets the value for Address2 to be an explicit nil

### UnsetAddress2
`func (o *BillingAddressAttributes) UnsetAddress2()`

UnsetAddress2 ensures that no value is present for Address2, not even an explicit nil
### GetCity

`func (o *BillingAddressAttributes) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *BillingAddressAttributes) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *BillingAddressAttributes) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *BillingAddressAttributes) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *BillingAddressAttributes) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *BillingAddressAttributes) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetCountry

`func (o *BillingAddressAttributes) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *BillingAddressAttributes) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *BillingAddressAttributes) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *BillingAddressAttributes) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetLocality

`func (o *BillingAddressAttributes) GetLocality() string`

GetLocality returns the Locality field if non-nil, zero value otherwise.

### GetLocalityOk

`func (o *BillingAddressAttributes) GetLocalityOk() (*string, bool)`

GetLocalityOk returns a tuple with the Locality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocality

`func (o *BillingAddressAttributes) SetLocality(v string)`

SetLocality sets Locality field to given value.

### HasLocality

`func (o *BillingAddressAttributes) HasLocality() bool`

HasLocality returns a boolean if a field has been set.

### SetLocalityNil

`func (o *BillingAddressAttributes) SetLocalityNil(b bool)`

 SetLocalityNil sets the value for Locality to be an explicit nil

### UnsetLocality
`func (o *BillingAddressAttributes) UnsetLocality()`

UnsetLocality ensures that no value is present for Locality, not even an explicit nil
### GetPostalCode

`func (o *BillingAddressAttributes) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *BillingAddressAttributes) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *BillingAddressAttributes) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *BillingAddressAttributes) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetState

`func (o *BillingAddressAttributes) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BillingAddressAttributes) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BillingAddressAttributes) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *BillingAddressAttributes) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *BillingAddressAttributes) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *BillingAddressAttributes) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetCustomerID

`func (o *BillingAddressAttributes) GetCustomerID() string`

GetCustomerID returns the CustomerID field if non-nil, zero value otherwise.

### GetCustomerIDOk

`func (o *BillingAddressAttributes) GetCustomerIDOk() (*string, bool)`

GetCustomerIDOk returns a tuple with the CustomerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerID

`func (o *BillingAddressAttributes) SetCustomerID(v string)`

SetCustomerID sets CustomerID field to given value.

### HasCustomerID

`func (o *BillingAddressAttributes) HasCustomerID() bool`

HasCustomerID returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
