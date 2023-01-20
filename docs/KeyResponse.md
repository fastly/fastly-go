# KeyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **[]string** |  | [optional] 
**Meta** | Pointer to [**GetStoresResponseMeta**](GetStoresResponseMeta.md) |  | [optional] 

## Methods

### NewKeyResponse

`func NewKeyResponse() *KeyResponse`

NewKeyResponse instantiates a new KeyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyResponseWithDefaults

`func NewKeyResponseWithDefaults() *KeyResponse`

NewKeyResponseWithDefaults instantiates a new KeyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *KeyResponse) GetData() []string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *KeyResponse) GetDataOk() (*[]string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *KeyResponse) SetData(v []string)`

SetData sets Data field to given value.

### HasData

`func (o *KeyResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *KeyResponse) GetMeta() GetStoresResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *KeyResponse) GetMetaOk() (*GetStoresResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *KeyResponse) SetMeta(v GetStoresResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *KeyResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
