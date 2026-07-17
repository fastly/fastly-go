# RoutingConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The user-defined name for the routing config. | 
**InitialVersion** | Pointer to [**InitialVersion**](InitialVersion.md) |  | [optional] 

## Methods

### NewRoutingConfig

`func NewRoutingConfig(name string, ) *RoutingConfig`

NewRoutingConfig instantiates a new RoutingConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutingConfigWithDefaults

`func NewRoutingConfigWithDefaults() *RoutingConfig`

NewRoutingConfigWithDefaults instantiates a new RoutingConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RoutingConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RoutingConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RoutingConfig) SetName(v string)`

SetName sets Name field to given value.


### GetInitialVersion

`func (o *RoutingConfig) GetInitialVersion() InitialVersion`

GetInitialVersion returns the InitialVersion field if non-nil, zero value otherwise.

### GetInitialVersionOk

`func (o *RoutingConfig) GetInitialVersionOk() (*InitialVersion, bool)`

GetInitialVersionOk returns a tuple with the InitialVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialVersion

`func (o *RoutingConfig) SetInitialVersion(v InitialVersion)`

SetInitialVersion sets InitialVersion field to given value.

### HasInitialVersion

`func (o *RoutingConfig) HasInitialVersion() bool`

HasInitialVersion returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


