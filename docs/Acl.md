# ACL

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name for the ACL. Must start with an alphanumeric character and contain only alphanumeric characters, underscores, and whitespace. | [optional] 

## Methods

### NewACL

`func NewACL() *ACL`

NewACL instantiates a new ACL object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewACLWithDefaults

`func NewACLWithDefaults() *ACL`

NewACLWithDefaults instantiates a new ACL object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ACL) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ACL) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ACL) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ACL) HasName() bool`

HasName returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
