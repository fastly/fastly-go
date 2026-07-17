# Page

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique page identifier | [optional] 
**WebsiteId** | Pointer to **string** | Parent website ID | [optional] 
**Name** | Pointer to **string** | Page name | [optional] 
**Description** | Pointer to **string** | Page description | [optional] 
**Notifications** | Pointer to [**[]Notification**](Notification.md) | Notification configurations for this page | [optional] 
**Paths** | Pointer to **[]string** | URL paths to monitor | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewPage

`func NewPage() *Page`

NewPage instantiates a new Page object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageWithDefaults

`func NewPageWithDefaults() *Page`

NewPageWithDefaults instantiates a new Page object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Page) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Page) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Page) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Page) HasId() bool`

HasId returns a boolean if a field has been set.

### GetWebsiteId

`func (o *Page) GetWebsiteId() string`

GetWebsiteId returns the WebsiteId field if non-nil, zero value otherwise.

### GetWebsiteIdOk

`func (o *Page) GetWebsiteIdOk() (*string, bool)`

GetWebsiteIdOk returns a tuple with the WebsiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteId

`func (o *Page) SetWebsiteId(v string)`

SetWebsiteId sets WebsiteId field to given value.

### HasWebsiteId

`func (o *Page) HasWebsiteId() bool`

HasWebsiteId returns a boolean if a field has been set.

### GetName

`func (o *Page) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Page) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Page) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Page) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *Page) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Page) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Page) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Page) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetNotifications

`func (o *Page) GetNotifications() []Notification`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *Page) GetNotificationsOk() (*[]Notification, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *Page) SetNotifications(v []Notification)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *Page) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### GetPaths

`func (o *Page) GetPaths() []string`

GetPaths returns the Paths field if non-nil, zero value otherwise.

### GetPathsOk

`func (o *Page) GetPathsOk() (*[]string, bool)`

GetPathsOk returns a tuple with the Paths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaths

`func (o *Page) SetPaths(v []string)`

SetPaths sets Paths field to given value.

### HasPaths

`func (o *Page) HasPaths() bool`

HasPaths returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Page) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Page) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Page) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Page) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Page) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Page) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Page) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Page) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


