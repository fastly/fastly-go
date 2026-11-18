# NetworkProtectRuleSetsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]NetworkProtectRuleSetResponse**](NetworkProtectRuleSetResponse.md) | Rule sets returned on the current page. | 
**Meta** | [**NetworkProtectListMeta**](NetworkProtectListMeta.md) |  | 

## Methods

### NewNetworkProtectRuleSetsResponse

`func NewNetworkProtectRuleSetsResponse(data []NetworkProtectRuleSetResponse, meta NetworkProtectListMeta, ) *NetworkProtectRuleSetsResponse`

NewNetworkProtectRuleSetsResponse instantiates a new NetworkProtectRuleSetsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkProtectRuleSetsResponseWithDefaults

`func NewNetworkProtectRuleSetsResponseWithDefaults() *NetworkProtectRuleSetsResponse`

NewNetworkProtectRuleSetsResponseWithDefaults instantiates a new NetworkProtectRuleSetsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *NetworkProtectRuleSetsResponse) GetData() []NetworkProtectRuleSetResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *NetworkProtectRuleSetsResponse) GetDataOk() (*[]NetworkProtectRuleSetResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *NetworkProtectRuleSetsResponse) SetData(v []NetworkProtectRuleSetResponse)`

SetData sets Data field to given value.


### GetMeta

`func (o *NetworkProtectRuleSetsResponse) GetMeta() NetworkProtectListMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *NetworkProtectRuleSetsResponse) GetMetaOk() (*NetworkProtectListMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *NetworkProtectRuleSetsResponse) SetMeta(v NetworkProtectListMeta)`

SetMeta sets Meta field to given value.



[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
