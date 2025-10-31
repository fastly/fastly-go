# TopWorkspace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of the workspace. | [optional] 
**Name** | Pointer to **string** | Name of the workspace. | [optional] 
**Count** | Pointer to **int32** | Count of attacks on this workspace for the specific attack type. | [optional] 

## Methods

### NewTopWorkspace

`func NewTopWorkspace() *TopWorkspace`

NewTopWorkspace instantiates a new TopWorkspace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopWorkspaceWithDefaults

`func NewTopWorkspaceWithDefaults() *TopWorkspace`

NewTopWorkspaceWithDefaults instantiates a new TopWorkspace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TopWorkspace) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TopWorkspace) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TopWorkspace) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TopWorkspace) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *TopWorkspace) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TopWorkspace) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TopWorkspace) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TopWorkspace) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCount

`func (o *TopWorkspace) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *TopWorkspace) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *TopWorkspace) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *TopWorkspace) HasCount() bool`

HasCount returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


