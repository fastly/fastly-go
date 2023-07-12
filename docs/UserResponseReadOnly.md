# UserResponseReadOnly

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ID** | Pointer to **string** |  | [optional] [readonly] 
**EmailHash** | Pointer to **string** | The alphanumeric string identifying a email login. | [optional] [readonly] 
**CustomerID** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewUserResponseReadOnly

`func NewUserResponseReadOnly() *UserResponseReadOnly`

NewUserResponseReadOnly instantiates a new UserResponseReadOnly object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserResponseReadOnlyWithDefaults

`func NewUserResponseReadOnlyWithDefaults() *UserResponseReadOnly`

NewUserResponseReadOnlyWithDefaults instantiates a new UserResponseReadOnly object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetID

`func (o *UserResponseReadOnly) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *UserResponseReadOnly) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *UserResponseReadOnly) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *UserResponseReadOnly) HasID() bool`

HasID returns a boolean if a field has been set.

### GetEmailHash

`func (o *UserResponseReadOnly) GetEmailHash() string`

GetEmailHash returns the EmailHash field if non-nil, zero value otherwise.

### GetEmailHashOk

`func (o *UserResponseReadOnly) GetEmailHashOk() (*string, bool)`

GetEmailHashOk returns a tuple with the EmailHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailHash

`func (o *UserResponseReadOnly) SetEmailHash(v string)`

SetEmailHash sets EmailHash field to given value.

### HasEmailHash

`func (o *UserResponseReadOnly) HasEmailHash() bool`

HasEmailHash returns a boolean if a field has been set.

### GetCustomerID

`func (o *UserResponseReadOnly) GetCustomerID() string`

GetCustomerID returns the CustomerID field if non-nil, zero value otherwise.

### GetCustomerIDOk

`func (o *UserResponseReadOnly) GetCustomerIDOk() (*string, bool)`

GetCustomerIDOk returns a tuple with the CustomerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerID

`func (o *UserResponseReadOnly) SetCustomerID(v string)`

SetCustomerID sets CustomerID field to given value.

### HasCustomerID

`func (o *UserResponseReadOnly) HasCustomerID() bool`

HasCustomerID returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
