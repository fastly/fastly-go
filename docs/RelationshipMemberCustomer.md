# RelationshipMemberCustomer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**TypeCustomer**](TypeCustomer.md) |  | [optional] [default to TYPECUSTOMER_CUSTOMER]
**Id** | Pointer to **string** |  | [optional] [readonly] 

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

### GetId

`func (o *RelationshipMemberCustomer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RelationshipMemberCustomer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RelationshipMemberCustomer) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RelationshipMemberCustomer) HasId() bool`

HasId returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


