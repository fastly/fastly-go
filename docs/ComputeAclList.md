# ComputeACLList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]ComputeACLCreateAclsResponse**](ComputeACLCreateAclsResponse.md) |  | [optional] 
**Meta** | Pointer to [**ComputeACLListMeta**](ComputeACLListMeta.md) |  | [optional] 

## Methods

### NewComputeACLList

`func NewComputeACLList() *ComputeACLList`

NewComputeACLList instantiates a new ComputeACLList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputeACLListWithDefaults

`func NewComputeACLListWithDefaults() *ComputeACLList`

NewComputeACLListWithDefaults instantiates a new ComputeACLList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ComputeACLList) GetData() []ComputeACLCreateAclsResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ComputeACLList) GetDataOk() (*[]ComputeACLCreateAclsResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ComputeACLList) SetData(v []ComputeACLCreateAclsResponse)`

SetData sets Data field to given value.

### HasData

`func (o *ComputeACLList) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *ComputeACLList) GetMeta() ComputeACLListMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ComputeACLList) GetMetaOk() (*ComputeACLListMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ComputeACLList) SetMeta(v ComputeACLListMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ComputeACLList) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
