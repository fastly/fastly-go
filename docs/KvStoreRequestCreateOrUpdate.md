# KvStoreRequestCreateOrUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A human-readable name for the store. Refer to https://docs.fastly.com/products/compute-resource-limits#kv-store for limitations on the KV store name. | 

## Methods

### NewKvStoreRequestCreateOrUpdate

`func NewKvStoreRequestCreateOrUpdate(name string, ) *KvStoreRequestCreateOrUpdate`

NewKvStoreRequestCreateOrUpdate instantiates a new KvStoreRequestCreateOrUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKvStoreRequestCreateOrUpdateWithDefaults

`func NewKvStoreRequestCreateOrUpdateWithDefaults() *KvStoreRequestCreateOrUpdate`

NewKvStoreRequestCreateOrUpdateWithDefaults instantiates a new KvStoreRequestCreateOrUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *KvStoreRequestCreateOrUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KvStoreRequestCreateOrUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KvStoreRequestCreateOrUpdate) SetName(v string)`

SetName sets Name field to given value.



[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


