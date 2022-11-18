# DomainCheckItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **NullableString** | A freeform descriptive note. | [optional] 
**Name** | Pointer to **string** | The name of the domain or domains associated with this service. | [optional] 

## Methods

### NewDomainCheckItem

`func NewDomainCheckItem() *DomainCheckItem`

NewDomainCheckItem instantiates a new DomainCheckItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainCheckItemWithDefaults

`func NewDomainCheckItemWithDefaults() *DomainCheckItem`

NewDomainCheckItemWithDefaults instantiates a new DomainCheckItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *DomainCheckItem) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *DomainCheckItem) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *DomainCheckItem) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *DomainCheckItem) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *DomainCheckItem) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *DomainCheckItem) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetName

`func (o *DomainCheckItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DomainCheckItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DomainCheckItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DomainCheckItem) HasName() bool`

HasName returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
