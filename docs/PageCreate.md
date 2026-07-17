# PageCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebsiteId** | **string** |  | 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Notifications** | Pointer to [**[]Notification**](Notification.md) |  | [optional] 
**Paths** | Pointer to **[]string** |  | [optional] 

## Methods

### NewPageCreate

`func NewPageCreate(websiteId string, name string, ) *PageCreate`

NewPageCreate instantiates a new PageCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageCreateWithDefaults

`func NewPageCreateWithDefaults() *PageCreate`

NewPageCreateWithDefaults instantiates a new PageCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebsiteId

`func (o *PageCreate) GetWebsiteId() string`

GetWebsiteId returns the WebsiteId field if non-nil, zero value otherwise.

### GetWebsiteIdOk

`func (o *PageCreate) GetWebsiteIdOk() (*string, bool)`

GetWebsiteIdOk returns a tuple with the WebsiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteId

`func (o *PageCreate) SetWebsiteId(v string)`

SetWebsiteId sets WebsiteId field to given value.


### GetName

`func (o *PageCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PageCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PageCreate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *PageCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PageCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PageCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PageCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetNotifications

`func (o *PageCreate) GetNotifications() []Notification`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *PageCreate) GetNotificationsOk() (*[]Notification, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *PageCreate) SetNotifications(v []Notification)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *PageCreate) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### GetPaths

`func (o *PageCreate) GetPaths() []string`

GetPaths returns the Paths field if non-nil, zero value otherwise.

### GetPathsOk

`func (o *PageCreate) GetPathsOk() (*[]string, bool)`

GetPathsOk returns a tuple with the Paths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaths

`func (o *PageCreate) SetPaths(v []string)`

SetPaths sets Paths field to given value.

### HasPaths

`func (o *PageCreate) HasPaths() bool`

HasPaths returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


