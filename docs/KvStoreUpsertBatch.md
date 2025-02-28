# KvStoreUpsertBatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | The key of the item to be added, updated, retrieved, or deleted. | 
**TimeToLiveSec** | Pointer to **int32** | Indicates that the item should be deleted after the specified number of seconds have elapsed. Deletion is not immediate but will occur within 24 hours of the requested expiration. | [optional] 
**Metadata** | Pointer to **string** | An arbitrary data field which can contain up to 2000 bytes of data. | [optional] 
**BackgroundFetch** | Pointer to **bool** | If set to true, the new value for the item will not be immediately visible to other users of the KV store; they will receive the existing (stale) value while the platform updates cached copies. Setting this to true ensures that other users of the KV store will receive responses to &#39;get&#39; operations for this item quickly, although they will be slightly out of date. | [optional] [default to false]
**Value** | **string** | Value for the item (in Base64 encoding). | 

## Methods

### NewKvStoreUpsertBatch

`func NewKvStoreUpsertBatch(key string, value string, ) *KvStoreUpsertBatch`

NewKvStoreUpsertBatch instantiates a new KvStoreUpsertBatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKvStoreUpsertBatchWithDefaults

`func NewKvStoreUpsertBatchWithDefaults() *KvStoreUpsertBatch`

NewKvStoreUpsertBatchWithDefaults instantiates a new KvStoreUpsertBatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *KvStoreUpsertBatch) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *KvStoreUpsertBatch) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *KvStoreUpsertBatch) SetKey(v string)`

SetKey sets Key field to given value.


### GetTimeToLiveSec

`func (o *KvStoreUpsertBatch) GetTimeToLiveSec() int32`

GetTimeToLiveSec returns the TimeToLiveSec field if non-nil, zero value otherwise.

### GetTimeToLiveSecOk

`func (o *KvStoreUpsertBatch) GetTimeToLiveSecOk() (*int32, bool)`

GetTimeToLiveSecOk returns a tuple with the TimeToLiveSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeToLiveSec

`func (o *KvStoreUpsertBatch) SetTimeToLiveSec(v int32)`

SetTimeToLiveSec sets TimeToLiveSec field to given value.

### HasTimeToLiveSec

`func (o *KvStoreUpsertBatch) HasTimeToLiveSec() bool`

HasTimeToLiveSec returns a boolean if a field has been set.

### GetMetadata

`func (o *KvStoreUpsertBatch) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *KvStoreUpsertBatch) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *KvStoreUpsertBatch) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *KvStoreUpsertBatch) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetBackgroundFetch

`func (o *KvStoreUpsertBatch) GetBackgroundFetch() bool`

GetBackgroundFetch returns the BackgroundFetch field if non-nil, zero value otherwise.

### GetBackgroundFetchOk

`func (o *KvStoreUpsertBatch) GetBackgroundFetchOk() (*bool, bool)`

GetBackgroundFetchOk returns a tuple with the BackgroundFetch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundFetch

`func (o *KvStoreUpsertBatch) SetBackgroundFetch(v bool)`

SetBackgroundFetch sets BackgroundFetch field to given value.

### HasBackgroundFetch

`func (o *KvStoreUpsertBatch) HasBackgroundFetch() bool`

HasBackgroundFetch returns a boolean if a field has been set.

### GetValue

`func (o *KvStoreUpsertBatch) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *KvStoreUpsertBatch) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *KvStoreUpsertBatch) SetValue(v string)`

SetValue sets Value field to given value.



[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
