# SnippetCommon

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name for the snippet. | [optional] 
**Type** | Pointer to **string** | The location in generated VCL where the snippet should be placed. | [optional] 
**Content** | Pointer to **string** | The VCL code that specifies exactly what the snippet does. | [optional] 
**Priority** | Pointer to **string** | Priority determines execution order. Lower numbers execute first. | [optional] [default to "100"]

## Methods

### NewSnippetCommon

`func NewSnippetCommon() *SnippetCommon`

NewSnippetCommon instantiates a new SnippetCommon object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnippetCommonWithDefaults

`func NewSnippetCommonWithDefaults() *SnippetCommon`

NewSnippetCommonWithDefaults instantiates a new SnippetCommon object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SnippetCommon) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SnippetCommon) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SnippetCommon) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SnippetCommon) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *SnippetCommon) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SnippetCommon) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SnippetCommon) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SnippetCommon) HasType() bool`

HasType returns a boolean if a field has been set.

### GetContent

`func (o *SnippetCommon) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *SnippetCommon) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *SnippetCommon) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *SnippetCommon) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetPriority

`func (o *SnippetCommon) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *SnippetCommon) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *SnippetCommon) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *SnippetCommon) HasPriority() bool`

HasPriority returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
