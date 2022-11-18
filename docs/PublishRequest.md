# PublishRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]PublishItem**](PublishItem.md) | The messages to publish. | 

## Methods

### NewPublishRequest

`func NewPublishRequest(items []PublishItem, ) *PublishRequest`

NewPublishRequest instantiates a new PublishRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublishRequestWithDefaults

`func NewPublishRequestWithDefaults() *PublishRequest`

NewPublishRequestWithDefaults instantiates a new PublishRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *PublishRequest) GetItems() []PublishItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *PublishRequest) GetItemsOk() (*[]PublishItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *PublishRequest) SetItems(v []PublishItem)`

SetItems sets Items field to given value.



[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
