# PublishItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | **string** | The channel to publish the message on. | 
**ID** | Pointer to **string** | The ID of the message. | [optional] 
**PrevID** | Pointer to **string** | The ID of the previous message published on the same channel. | [optional] 
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


### GetID

`func (o *PublishItem) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *PublishItem) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *PublishItem) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *PublishItem) HasID() bool`

HasID returns a boolean if a field has been set.

### GetPrevID

`func (o *PublishItem) GetPrevID() string`

GetPrevID returns the PrevID field if non-nil, zero value otherwise.

### GetPrevIDOk

`func (o *PublishItem) GetPrevIDOk() (*string, bool)`

GetPrevIDOk returns a tuple with the PrevID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevID

`func (o *PublishItem) SetPrevID(v string)`

SetPrevID sets PrevID field to given value.

### HasPrevID

`func (o *PublishItem) HasPrevID() bool`

HasPrevID returns a boolean if a field has been set.

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
