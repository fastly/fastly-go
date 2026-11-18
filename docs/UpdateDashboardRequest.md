# UpdateDashboardRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A human-readable name | [optional] 
**Description** | Pointer to **string** | A short description of the dashboard | [optional] 
**Items** | Pointer to [**[]DashboardItem**](DashboardItem.md) | A list of [dashboard items](#dashboard-item). | [optional] 

## Methods

### NewUpdateDashboardRequest

`func NewUpdateDashboardRequest() *UpdateDashboardRequest`

NewUpdateDashboardRequest instantiates a new UpdateDashboardRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDashboardRequestWithDefaults

`func NewUpdateDashboardRequestWithDefaults() *UpdateDashboardRequest`

NewUpdateDashboardRequestWithDefaults instantiates a new UpdateDashboardRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateDashboardRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateDashboardRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateDashboardRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateDashboardRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateDashboardRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateDashboardRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateDashboardRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateDashboardRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetItems

`func (o *UpdateDashboardRequest) GetItems() []DashboardItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *UpdateDashboardRequest) GetItemsOk() (*[]DashboardItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *UpdateDashboardRequest) SetItems(v []DashboardItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *UpdateDashboardRequest) HasItems() bool`

HasItems returns a boolean if a field has been set.

### SetItemsNil

`func (o *UpdateDashboardRequest) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *UpdateDashboardRequest) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil

[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
