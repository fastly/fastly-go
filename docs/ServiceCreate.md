# ServiceCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **NullableString** | A freeform descriptive note. | [optional] 
**Name** | Pointer to **string** | The name of the service. | [optional] 
**CustomerId** | Pointer to **string** | Alphanumeric string identifying the customer. | [optional] 
**Type** | Pointer to **string** | The type of this service. | [optional] 

## Methods

### NewServiceCreate

`func NewServiceCreate() *ServiceCreate`

NewServiceCreate instantiates a new ServiceCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceCreateWithDefaults

`func NewServiceCreateWithDefaults() *ServiceCreate`

NewServiceCreateWithDefaults instantiates a new ServiceCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *ServiceCreate) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ServiceCreate) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ServiceCreate) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ServiceCreate) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *ServiceCreate) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *ServiceCreate) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetName

`func (o *ServiceCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceCreate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServiceCreate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCustomerId

`func (o *ServiceCreate) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *ServiceCreate) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *ServiceCreate) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *ServiceCreate) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetType

`func (o *ServiceCreate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServiceCreate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServiceCreate) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ServiceCreate) HasType() bool`

HasType returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


