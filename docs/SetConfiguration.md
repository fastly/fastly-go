# SetConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkspaceID** | Pointer to **string** | The new workspace_id. Required in the `PUT` request body when `product_id` is `ngwaf`. Optional in the `PATCH` request body for `ngwaf`. | [optional] 
**TrafficRamp** | Pointer to **string** | The new traffic ramp. Optional in the `PATCH` request body for `ngwaf`. | [optional] 

## Methods

### NewSetConfiguration

`func NewSetConfiguration() *SetConfiguration`

NewSetConfiguration instantiates a new SetConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetConfigurationWithDefaults

`func NewSetConfigurationWithDefaults() *SetConfiguration`

NewSetConfigurationWithDefaults instantiates a new SetConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkspaceID

`func (o *SetConfiguration) GetWorkspaceID() string`

GetWorkspaceID returns the WorkspaceID field if non-nil, zero value otherwise.

### GetWorkspaceIDOk

`func (o *SetConfiguration) GetWorkspaceIDOk() (*string, bool)`

GetWorkspaceIDOk returns a tuple with the WorkspaceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceID

`func (o *SetConfiguration) SetWorkspaceID(v string)`

SetWorkspaceID sets WorkspaceID field to given value.

### HasWorkspaceID

`func (o *SetConfiguration) HasWorkspaceID() bool`

HasWorkspaceID returns a boolean if a field has been set.

### GetTrafficRamp

`func (o *SetConfiguration) GetTrafficRamp() string`

GetTrafficRamp returns the TrafficRamp field if non-nil, zero value otherwise.

### GetTrafficRampOk

`func (o *SetConfiguration) GetTrafficRampOk() (*string, bool)`

GetTrafficRampOk returns a tuple with the TrafficRamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficRamp

`func (o *SetConfiguration) SetTrafficRamp(v string)`

SetTrafficRamp sets TrafficRamp field to given value.

### HasTrafficRamp

`func (o *SetConfiguration) HasTrafficRamp() bool`

HasTrafficRamp returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
