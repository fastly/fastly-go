# RelationshipMemberCustomer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**TypeCustomer**](TypeCustomer.md) |  | [optional] [default to TYPECUSTOMER_CUSTOMER]
**ID** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewRelationshipMemberCustomer

`func NewRelationshipMemberCustomer() *RelationshipMemberCustomer`

NewRelationshipMemberCustomer instantiates a new RelationshipMemberCustomer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelationshipMemberCustomerWithDefaults

`func NewRelationshipMemberCustomerWithDefaults() *RelationshipMemberCustomer`

NewRelationshipMemberCustomerWithDefaults instantiates a new RelationshipMemberCustomer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RelationshipMemberCustomer) GetType() TypeCustomer`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RelationshipMemberCustomer) GetTypeOk() (*TypeCustomer, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RelationshipMemberCustomer) SetType(v TypeCustomer)`

SetType sets Type field to given value.

### HasType

`func (o *RelationshipMemberCustomer) HasType() bool`

HasType returns a boolean if a field has been set.

### GetID

`func (o *RelationshipMemberCustomer) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *RelationshipMemberCustomer) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *RelationshipMemberCustomer) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *RelationshipMemberCustomer) HasID() bool`

HasID returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
