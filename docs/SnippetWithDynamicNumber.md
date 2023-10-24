# SnippetWithDynamicNumber

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name for the snippet. | [optional] 
**Type** | Pointer to **string** | The location in generated VCL where the snippet should be placed. | [optional] 
**Content** | Pointer to **string** | The VCL code that specifies exactly what the snippet does. | [optional] 
**Priority** | Pointer to **string** | Priority determines execution order. Lower numbers execute first. | [optional] [default to "100"]
**Dynamic** | Pointer to **float32** | Sets the snippet version. | [optional] 

## Methods

### NewSnippetWithDynamicNumber

`func NewSnippetWithDynamicNumber() *SnippetWithDynamicNumber`

NewSnippetWithDynamicNumber instantiates a new SnippetWithDynamicNumber object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnippetWithDynamicNumberWithDefaults

`func NewSnippetWithDynamicNumberWithDefaults() *SnippetWithDynamicNumber`

NewSnippetWithDynamicNumberWithDefaults instantiates a new SnippetWithDynamicNumber object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SnippetWithDynamicNumber) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SnippetWithDynamicNumber) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SnippetWithDynamicNumber) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SnippetWithDynamicNumber) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *SnippetWithDynamicNumber) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SnippetWithDynamicNumber) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SnippetWithDynamicNumber) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SnippetWithDynamicNumber) HasType() bool`

HasType returns a boolean if a field has been set.

### GetContent

`func (o *SnippetWithDynamicNumber) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *SnippetWithDynamicNumber) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *SnippetWithDynamicNumber) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *SnippetWithDynamicNumber) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetPriority

`func (o *SnippetWithDynamicNumber) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *SnippetWithDynamicNumber) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *SnippetWithDynamicNumber) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *SnippetWithDynamicNumber) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetDynamic

`func (o *SnippetWithDynamicNumber) GetDynamic() float32`

GetDynamic returns the Dynamic field if non-nil, zero value otherwise.

### GetDynamicOk

`func (o *SnippetWithDynamicNumber) GetDynamicOk() (*float32, bool)`

GetDynamicOk returns a tuple with the Dynamic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamic

`func (o *SnippetWithDynamicNumber) SetDynamic(v float32)`

SetDynamic sets Dynamic field to given value.

### HasDynamic

`func (o *SnippetWithDynamicNumber) HasDynamic() bool`

HasDynamic returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
