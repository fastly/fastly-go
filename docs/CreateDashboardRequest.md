# CreateDashboardRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A human-readable name | 
**Description** | Pointer to **string** | A short description of the dashboard | [optional] 
**Items** | Pointer to [**[]DashboardItem**](DashboardItem.md) | A list of [dashboard items](#dashboard-item). | [optional] 

## Methods

### NewCreateDashboardRequest

`func NewCreateDashboardRequest(name string, ) *CreateDashboardRequest`

NewCreateDashboardRequest instantiates a new CreateDashboardRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDashboardRequestWithDefaults

`func NewCreateDashboardRequestWithDefaults() *CreateDashboardRequest`

NewCreateDashboardRequestWithDefaults instantiates a new CreateDashboardRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateDashboardRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateDashboardRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateDashboardRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateDashboardRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateDashboardRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateDashboardRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateDashboardRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetItems

`func (o *CreateDashboardRequest) GetItems() []DashboardItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CreateDashboardRequest) GetItemsOk() (*[]DashboardItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CreateDashboardRequest) SetItems(v []DashboardItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *CreateDashboardRequest) HasItems() bool`

HasItems returns a boolean if a field has been set.

### SetItemsNil

`func (o *CreateDashboardRequest) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *CreateDashboardRequest) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil

[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
