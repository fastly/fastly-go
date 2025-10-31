# ComputeAclList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]ComputeAclCreateAclsResponse**](ComputeAclCreateAclsResponse.md) |  | [optional] 
**Meta** | Pointer to [**ComputeAclListMeta**](ComputeAclListMeta.md) |  | [optional] 

## Methods

### NewComputeAclList

`func NewComputeAclList() *ComputeAclList`

NewComputeAclList instantiates a new ComputeAclList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputeAclListWithDefaults

`func NewComputeAclListWithDefaults() *ComputeAclList`

NewComputeAclListWithDefaults instantiates a new ComputeAclList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ComputeAclList) GetData() []ComputeAclCreateAclsResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ComputeAclList) GetDataOk() (*[]ComputeAclCreateAclsResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ComputeAclList) SetData(v []ComputeAclCreateAclsResponse)`

SetData sets Data field to given value.

### HasData

`func (o *ComputeAclList) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *ComputeAclList) GetMeta() ComputeAclListMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ComputeAclList) GetMetaOk() (*ComputeAclListMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ComputeAclList) SetMeta(v ComputeAclListMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ComputeAclList) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


