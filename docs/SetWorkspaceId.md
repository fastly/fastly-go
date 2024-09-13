# SetWorkspaceID

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkspaceID** | Pointer to **string** | The workspace to link with the Next-Gen WAF product. Note this body is only supported and required when `product_id` is `ngwaf` | [optional] 

## Methods

### NewSetWorkspaceID

`func NewSetWorkspaceID() *SetWorkspaceID`

NewSetWorkspaceID instantiates a new SetWorkspaceID object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetWorkspaceIDWithDefaults

`func NewSetWorkspaceIDWithDefaults() *SetWorkspaceID`

NewSetWorkspaceIDWithDefaults instantiates a new SetWorkspaceID object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkspaceID

`func (o *SetWorkspaceID) GetWorkspaceID() string`

GetWorkspaceID returns the WorkspaceID field if non-nil, zero value otherwise.

### GetWorkspaceIDOk

`func (o *SetWorkspaceID) GetWorkspaceIDOk() (*string, bool)`

GetWorkspaceIDOk returns a tuple with the WorkspaceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceID

`func (o *SetWorkspaceID) SetWorkspaceID(v string)`

SetWorkspaceID sets WorkspaceID field to given value.

### HasWorkspaceID

`func (o *SetWorkspaceID) HasWorkspaceID() bool`

HasWorkspaceID returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
