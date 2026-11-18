# ArcVirtualKeysResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]ArcVirtualKeyListItemResponse**](ArcVirtualKeyListItemResponse.md) | The virtual keys returned by the request. | 
**Meta** | [**ArcVirtualKeysMeta**](ArcVirtualKeysMeta.md) |  | 

## Methods

### NewArcVirtualKeysResponse

`func NewArcVirtualKeysResponse(data []ArcVirtualKeyListItemResponse, meta ArcVirtualKeysMeta, ) *ArcVirtualKeysResponse`

NewArcVirtualKeysResponse instantiates a new ArcVirtualKeysResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArcVirtualKeysResponseWithDefaults

`func NewArcVirtualKeysResponseWithDefaults() *ArcVirtualKeysResponse`

NewArcVirtualKeysResponseWithDefaults instantiates a new ArcVirtualKeysResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ArcVirtualKeysResponse) GetData() []ArcVirtualKeyListItemResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ArcVirtualKeysResponse) GetDataOk() (*[]ArcVirtualKeyListItemResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ArcVirtualKeysResponse) SetData(v []ArcVirtualKeyListItemResponse)`

SetData sets Data field to given value.


### GetMeta

`func (o *ArcVirtualKeysResponse) GetMeta() ArcVirtualKeysMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ArcVirtualKeysResponse) GetMetaOk() (*ArcVirtualKeysMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ArcVirtualKeysResponse) SetMeta(v ArcVirtualKeysMeta)`

SetMeta sets Meta field to given value.



[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
