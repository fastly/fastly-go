# PageUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebsiteID** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Notifications** | Pointer to [**[]Notification**](Notification.md) |  | [optional] 
**Paths** | Pointer to **[]string** |  | [optional] 

## Methods

### NewPageUpdate

`func NewPageUpdate() *PageUpdate`

NewPageUpdate instantiates a new PageUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageUpdateWithDefaults

`func NewPageUpdateWithDefaults() *PageUpdate`

NewPageUpdateWithDefaults instantiates a new PageUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebsiteID

`func (o *PageUpdate) GetWebsiteID() string`

GetWebsiteID returns the WebsiteID field if non-nil, zero value otherwise.

### GetWebsiteIDOk

`func (o *PageUpdate) GetWebsiteIDOk() (*string, bool)`

GetWebsiteIDOk returns a tuple with the WebsiteID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteID

`func (o *PageUpdate) SetWebsiteID(v string)`

SetWebsiteID sets WebsiteID field to given value.

### HasWebsiteID

`func (o *PageUpdate) HasWebsiteID() bool`

HasWebsiteID returns a boolean if a field has been set.

### GetName

`func (o *PageUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PageUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PageUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PageUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PageUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PageUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PageUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PageUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetNotifications

`func (o *PageUpdate) GetNotifications() []Notification`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *PageUpdate) GetNotificationsOk() (*[]Notification, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *PageUpdate) SetNotifications(v []Notification)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *PageUpdate) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### GetPaths

`func (o *PageUpdate) GetPaths() []string`

GetPaths returns the Paths field if non-nil, zero value otherwise.

### GetPathsOk

`func (o *PageUpdate) GetPathsOk() (*[]string, bool)`

GetPathsOk returns a tuple with the Paths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaths

`func (o *PageUpdate) SetPaths(v []string)`

SetPaths sets Paths field to given value.

### HasPaths

`func (o *PageUpdate) HasPaths() bool`

HasPaths returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
