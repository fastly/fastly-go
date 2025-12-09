# NgwafRequestEnable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkspaceId** | **string** | The workspace to link. | 
**TrafficRamp** | Pointer to **string** | The percentage of traffic to inspect. | [optional] 

## Methods

### NewNgwafRequestEnable

`func NewNgwafRequestEnable(workspaceId string, ) *NgwafRequestEnable`

NewNgwafRequestEnable instantiates a new NgwafRequestEnable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNgwafRequestEnableWithDefaults

`func NewNgwafRequestEnableWithDefaults() *NgwafRequestEnable`

NewNgwafRequestEnableWithDefaults instantiates a new NgwafRequestEnable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkspaceId

`func (o *NgwafRequestEnable) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *NgwafRequestEnable) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *NgwafRequestEnable) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.


### GetTrafficRamp

`func (o *NgwafRequestEnable) GetTrafficRamp() string`

GetTrafficRamp returns the TrafficRamp field if non-nil, zero value otherwise.

### GetTrafficRampOk

`func (o *NgwafRequestEnable) GetTrafficRampOk() (*string, bool)`

GetTrafficRampOk returns a tuple with the TrafficRamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficRamp

`func (o *NgwafRequestEnable) SetTrafficRamp(v string)`

SetTrafficRamp sets TrafficRamp field to given value.

### HasTrafficRamp

`func (o *NgwafRequestEnable) HasTrafficRamp() bool`

HasTrafficRamp returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


