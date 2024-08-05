# ListDashboardsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]Dashboard**](Dashboard.md) |  | [optional] 
**Meta** | Pointer to **any** | Meta for the pagination. | [optional] 

## Methods

### NewListDashboardsResponse

`func NewListDashboardsResponse() *ListDashboardsResponse`

NewListDashboardsResponse instantiates a new ListDashboardsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDashboardsResponseWithDefaults

`func NewListDashboardsResponseWithDefaults() *ListDashboardsResponse`

NewListDashboardsResponseWithDefaults instantiates a new ListDashboardsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListDashboardsResponse) GetData() []Dashboard`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListDashboardsResponse) GetDataOk() (*[]Dashboard, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListDashboardsResponse) SetData(v []Dashboard)`

SetData sets Data field to given value.

### HasData

`func (o *ListDashboardsResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *ListDashboardsResponse) GetMeta() any`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListDashboardsResponse) GetMetaOk() (*any, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListDashboardsResponse) SetMeta(v any)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListDashboardsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### SetMetaNil

`func (o *ListDashboardsResponse) SetMetaNil(b bool)`

 SetMetaNil sets the value for Meta to be an explicit nil

### UnsetMeta
`func (o *ListDashboardsResponse) UnsetMeta()`

UnsetMeta ensures that no value is present for Meta, not even an explicit nil

[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
