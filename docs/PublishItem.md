# PublishItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | **string** | The channel to publish the message on. | 
**Id** | Pointer to **string** | The ID of the message. | [optional] 
**PrevId** | Pointer to **string** | The ID of the previous message published on the same channel. | [optional] 
**Formats** | [**PublishItemFormats**](PublishItemFormats.md) |  | 

## Methods

### NewPublishItem

`func NewPublishItem(channel string, formats PublishItemFormats, ) *PublishItem`

NewPublishItem instantiates a new PublishItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublishItemWithDefaults

`func NewPublishItemWithDefaults() *PublishItem`

NewPublishItemWithDefaults instantiates a new PublishItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *PublishItem) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *PublishItem) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *PublishItem) SetChannel(v string)`

SetChannel sets Channel field to given value.


### GetId

`func (o *PublishItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PublishItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PublishItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PublishItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPrevId

`func (o *PublishItem) GetPrevId() string`

GetPrevId returns the PrevId field if non-nil, zero value otherwise.

### GetPrevIdOk

`func (o *PublishItem) GetPrevIdOk() (*string, bool)`

GetPrevIdOk returns a tuple with the PrevId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevId

`func (o *PublishItem) SetPrevId(v string)`

SetPrevId sets PrevId field to given value.

### HasPrevId

`func (o *PublishItem) HasPrevId() bool`

HasPrevId returns a boolean if a field has been set.

### GetFormats

`func (o *PublishItem) GetFormats() PublishItemFormats`

GetFormats returns the Formats field if non-nil, zero value otherwise.

### GetFormatsOk

`func (o *PublishItem) GetFormatsOk() (*PublishItemFormats, bool)`

GetFormatsOk returns a tuple with the Formats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormats

`func (o *PublishItem) SetFormats(v PublishItemFormats)`

SetFormats sets Formats field to given value.



[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


