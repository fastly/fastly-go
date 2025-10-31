# SchemasUserResponseReadOnly

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**EmailHash** | Pointer to **string** | The alphanumeric string identifying a email login. | [optional] [readonly] 
**CustomerId** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewSchemasUserResponseReadOnly

`func NewSchemasUserResponseReadOnly() *SchemasUserResponseReadOnly`

NewSchemasUserResponseReadOnly instantiates a new SchemasUserResponseReadOnly object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemasUserResponseReadOnlyWithDefaults

`func NewSchemasUserResponseReadOnlyWithDefaults() *SchemasUserResponseReadOnly`

NewSchemasUserResponseReadOnlyWithDefaults instantiates a new SchemasUserResponseReadOnly object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SchemasUserResponseReadOnly) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SchemasUserResponseReadOnly) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SchemasUserResponseReadOnly) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SchemasUserResponseReadOnly) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEmailHash

`func (o *SchemasUserResponseReadOnly) GetEmailHash() string`

GetEmailHash returns the EmailHash field if non-nil, zero value otherwise.

### GetEmailHashOk

`func (o *SchemasUserResponseReadOnly) GetEmailHashOk() (*string, bool)`

GetEmailHashOk returns a tuple with the EmailHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailHash

`func (o *SchemasUserResponseReadOnly) SetEmailHash(v string)`

SetEmailHash sets EmailHash field to given value.

### HasEmailHash

`func (o *SchemasUserResponseReadOnly) HasEmailHash() bool`

HasEmailHash returns a boolean if a field has been set.

### GetCustomerId

`func (o *SchemasUserResponseReadOnly) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *SchemasUserResponseReadOnly) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *SchemasUserResponseReadOnly) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *SchemasUserResponseReadOnly) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


