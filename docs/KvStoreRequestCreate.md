# KvStoreRequestCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A human-readable name for the store. Refer to https://docs.fastly.com/products/compute-resource-limits#kv-store for limitations on the KV store name. | 

## Methods

### NewKvStoreRequestCreate

`func NewKvStoreRequestCreate(name string, ) *KvStoreRequestCreate`

NewKvStoreRequestCreate instantiates a new KvStoreRequestCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKvStoreRequestCreateWithDefaults

`func NewKvStoreRequestCreateWithDefaults() *KvStoreRequestCreate`

NewKvStoreRequestCreateWithDefaults instantiates a new KvStoreRequestCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *KvStoreRequestCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KvStoreRequestCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KvStoreRequestCreate) SetName(v string)`

SetName sets Name field to given value.



[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
