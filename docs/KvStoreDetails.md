# KvStoreDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ID** | Pointer to **string** | ID of the store. | [optional] 
**Name** | Pointer to **string** | Name of the store. | [optional] 

## Methods

### NewKvStoreDetails

`func NewKvStoreDetails() *KvStoreDetails`

NewKvStoreDetails instantiates a new KvStoreDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKvStoreDetailsWithDefaults

`func NewKvStoreDetailsWithDefaults() *KvStoreDetails`

NewKvStoreDetailsWithDefaults instantiates a new KvStoreDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetID

`func (o *KvStoreDetails) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *KvStoreDetails) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *KvStoreDetails) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *KvStoreDetails) HasID() bool`

HasID returns a boolean if a field has been set.

### GetName

`func (o *KvStoreDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KvStoreDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KvStoreDetails) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KvStoreDetails) HasName() bool`

HasName returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
