# Condition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **NullableString** | A freeform descriptive note. | [optional] 
**Name** | Pointer to **string** | Name of the condition. Required. | [optional] 
**Priority** | Pointer to **string** | A numeric string. Priority determines execution order. Lower numbers execute first. | [optional] [default to "100"]
**Statement** | Pointer to **string** | A conditional expression in VCL used to determine if the condition is met. | [optional] 
**ServiceId** | Pointer to **string** |  | [optional] [readonly] 
**Version** | Pointer to **string** | A numeric string that represents the service version. | [optional] 
**Type** | Pointer to **string** | Type of the condition. Required. | [optional] 

## Methods

### NewCondition

`func NewCondition() *Condition`

NewCondition instantiates a new Condition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConditionWithDefaults

`func NewConditionWithDefaults() *Condition`

NewConditionWithDefaults instantiates a new Condition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *Condition) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Condition) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Condition) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Condition) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *Condition) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *Condition) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetName

`func (o *Condition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Condition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Condition) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Condition) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPriority

`func (o *Condition) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *Condition) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *Condition) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *Condition) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetStatement

`func (o *Condition) GetStatement() string`

GetStatement returns the Statement field if non-nil, zero value otherwise.

### GetStatementOk

`func (o *Condition) GetStatementOk() (*string, bool)`

GetStatementOk returns a tuple with the Statement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatement

`func (o *Condition) SetStatement(v string)`

SetStatement sets Statement field to given value.

### HasStatement

`func (o *Condition) HasStatement() bool`

HasStatement returns a boolean if a field has been set.

### GetServiceId

`func (o *Condition) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *Condition) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *Condition) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *Condition) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetVersion

`func (o *Condition) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Condition) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Condition) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Condition) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetType

`func (o *Condition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Condition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Condition) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Condition) HasType() bool`

HasType returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


