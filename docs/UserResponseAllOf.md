# UserResponseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ID** | Pointer to **string** |  | [optional] [readonly] 
**EmailHash** | Pointer to **string** | The alphanumeric string identifying a email login. | [optional] [readonly] 
**CustomerID** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewUserResponseAllOf

`func NewUserResponseAllOf() *UserResponseAllOf`

NewUserResponseAllOf instantiates a new UserResponseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserResponseAllOfWithDefaults

`func NewUserResponseAllOfWithDefaults() *UserResponseAllOf`

NewUserResponseAllOfWithDefaults instantiates a new UserResponseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetID

`func (o *UserResponseAllOf) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *UserResponseAllOf) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *UserResponseAllOf) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *UserResponseAllOf) HasID() bool`

HasID returns a boolean if a field has been set.

### GetEmailHash

`func (o *UserResponseAllOf) GetEmailHash() string`

GetEmailHash returns the EmailHash field if non-nil, zero value otherwise.

### GetEmailHashOk

`func (o *UserResponseAllOf) GetEmailHashOk() (*string, bool)`

GetEmailHashOk returns a tuple with the EmailHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailHash

`func (o *UserResponseAllOf) SetEmailHash(v string)`

SetEmailHash sets EmailHash field to given value.

### HasEmailHash

`func (o *UserResponseAllOf) HasEmailHash() bool`

HasEmailHash returns a boolean if a field has been set.

### GetCustomerID

`func (o *UserResponseAllOf) GetCustomerID() string`

GetCustomerID returns the CustomerID field if non-nil, zero value otherwise.

### GetCustomerIDOk

`func (o *UserResponseAllOf) GetCustomerIDOk() (*string, bool)`

GetCustomerIDOk returns a tuple with the CustomerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerID

`func (o *UserResponseAllOf) SetCustomerID(v string)`

SetCustomerID sets CustomerID field to given value.

### HasCustomerID

`func (o *UserResponseAllOf) HasCustomerID() bool`

HasCustomerID returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
