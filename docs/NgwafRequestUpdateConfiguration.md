# NgwafRequestUpdateConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkspaceID** | Pointer to **string** | The workspace to link. | [optional] 
**TrafficRamp** | Pointer to **string** | The percentage of traffic to inspect. | [optional] 

## Methods

### NewNgwafRequestUpdateConfiguration

`func NewNgwafRequestUpdateConfiguration() *NgwafRequestUpdateConfiguration`

NewNgwafRequestUpdateConfiguration instantiates a new NgwafRequestUpdateConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNgwafRequestUpdateConfigurationWithDefaults

`func NewNgwafRequestUpdateConfigurationWithDefaults() *NgwafRequestUpdateConfiguration`

NewNgwafRequestUpdateConfigurationWithDefaults instantiates a new NgwafRequestUpdateConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkspaceID

`func (o *NgwafRequestUpdateConfiguration) GetWorkspaceID() string`

GetWorkspaceID returns the WorkspaceID field if non-nil, zero value otherwise.

### GetWorkspaceIDOk

`func (o *NgwafRequestUpdateConfiguration) GetWorkspaceIDOk() (*string, bool)`

GetWorkspaceIDOk returns a tuple with the WorkspaceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceID

`func (o *NgwafRequestUpdateConfiguration) SetWorkspaceID(v string)`

SetWorkspaceID sets WorkspaceID field to given value.

### HasWorkspaceID

`func (o *NgwafRequestUpdateConfiguration) HasWorkspaceID() bool`

HasWorkspaceID returns a boolean if a field has been set.

### GetTrafficRamp

`func (o *NgwafRequestUpdateConfiguration) GetTrafficRamp() string`

GetTrafficRamp returns the TrafficRamp field if non-nil, zero value otherwise.

### GetTrafficRampOk

`func (o *NgwafRequestUpdateConfiguration) GetTrafficRampOk() (*string, bool)`

GetTrafficRampOk returns a tuple with the TrafficRamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficRamp

`func (o *NgwafRequestUpdateConfiguration) SetTrafficRamp(v string)`

SetTrafficRamp sets TrafficRamp field to given value.

### HasTrafficRamp

`func (o *NgwafRequestUpdateConfiguration) HasTrafficRamp() bool`

HasTrafficRamp returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
