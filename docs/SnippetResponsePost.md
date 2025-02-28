# SnippetResponsePost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name for the snippet. | [optional] 
**Type** | Pointer to **string** | The location in generated VCL where the snippet should be placed. | [optional] 
**Content** | Pointer to **NullableString** | The VCL code that specifies exactly what the snippet does. | [optional] 
**Priority** | Pointer to **string** | Priority determines execution order. Lower numbers execute first. | [optional] [default to "100"]
**Dynamic** | Pointer to **float32** | Sets the snippet version. | [optional] 
**CreatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**DeletedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**UpdatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**ServiceID** | Pointer to **string** |  | [optional] [readonly] 
**Version** | Pointer to **string** | String representing the number identifying a version of the service. | [optional] [readonly] 
**ID** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewSnippetResponsePost

`func NewSnippetResponsePost() *SnippetResponsePost`

NewSnippetResponsePost instantiates a new SnippetResponsePost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnippetResponsePostWithDefaults

`func NewSnippetResponsePostWithDefaults() *SnippetResponsePost`

NewSnippetResponsePostWithDefaults instantiates a new SnippetResponsePost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SnippetResponsePost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SnippetResponsePost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SnippetResponsePost) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SnippetResponsePost) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *SnippetResponsePost) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SnippetResponsePost) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SnippetResponsePost) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SnippetResponsePost) HasType() bool`

HasType returns a boolean if a field has been set.

### GetContent

`func (o *SnippetResponsePost) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *SnippetResponsePost) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *SnippetResponsePost) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *SnippetResponsePost) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *SnippetResponsePost) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *SnippetResponsePost) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetPriority

`func (o *SnippetResponsePost) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *SnippetResponsePost) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *SnippetResponsePost) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *SnippetResponsePost) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetDynamic

`func (o *SnippetResponsePost) GetDynamic() float32`

GetDynamic returns the Dynamic field if non-nil, zero value otherwise.

### GetDynamicOk

`func (o *SnippetResponsePost) GetDynamicOk() (*float32, bool)`

GetDynamicOk returns a tuple with the Dynamic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamic

`func (o *SnippetResponsePost) SetDynamic(v float32)`

SetDynamic sets Dynamic field to given value.

### HasDynamic

`func (o *SnippetResponsePost) HasDynamic() bool`

HasDynamic returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SnippetResponsePost) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SnippetResponsePost) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SnippetResponsePost) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SnippetResponsePost) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *SnippetResponsePost) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *SnippetResponsePost) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetDeletedAt

`func (o *SnippetResponsePost) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *SnippetResponsePost) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *SnippetResponsePost) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *SnippetResponsePost) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *SnippetResponsePost) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *SnippetResponsePost) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetUpdatedAt

`func (o *SnippetResponsePost) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SnippetResponsePost) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SnippetResponsePost) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SnippetResponsePost) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *SnippetResponsePost) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *SnippetResponsePost) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetServiceID

`func (o *SnippetResponsePost) GetServiceID() string`

GetServiceID returns the ServiceID field if non-nil, zero value otherwise.

### GetServiceIDOk

`func (o *SnippetResponsePost) GetServiceIDOk() (*string, bool)`

GetServiceIDOk returns a tuple with the ServiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceID

`func (o *SnippetResponsePost) SetServiceID(v string)`

SetServiceID sets ServiceID field to given value.

### HasServiceID

`func (o *SnippetResponsePost) HasServiceID() bool`

HasServiceID returns a boolean if a field has been set.

### GetVersion

`func (o *SnippetResponsePost) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SnippetResponsePost) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SnippetResponsePost) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SnippetResponsePost) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetID

`func (o *SnippetResponsePost) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *SnippetResponsePost) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *SnippetResponsePost) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *SnippetResponsePost) HasID() bool`

HasID returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
