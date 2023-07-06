# LegacyWafUpdateStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompletedAt** | Pointer to **string** | Date and time that job was completed. | [optional] 
**CreatedAt** | Pointer to **string** | Date and time that job was created. | [optional] 
**Data** | Pointer to **string** | This field can contain data passed to the background worker as well as output from the background job. | [optional] 
**Message** | Pointer to **string** | Message with information about the status of the update. | [optional] 
**Status** | Pointer to **string** | Current status of the update. | [optional] 
**UpdatedAt** | Pointer to **string** | Date and time that job was last updated. | [optional] 

## Methods

### NewLegacyWafUpdateStatus

`func NewLegacyWafUpdateStatus() *LegacyWafUpdateStatus`

NewLegacyWafUpdateStatus instantiates a new LegacyWafUpdateStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLegacyWafUpdateStatusWithDefaults

`func NewLegacyWafUpdateStatusWithDefaults() *LegacyWafUpdateStatus`

NewLegacyWafUpdateStatusWithDefaults instantiates a new LegacyWafUpdateStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompletedAt

`func (o *LegacyWafUpdateStatus) GetCompletedAt() string`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *LegacyWafUpdateStatus) GetCompletedAtOk() (*string, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *LegacyWafUpdateStatus) SetCompletedAt(v string)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *LegacyWafUpdateStatus) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *LegacyWafUpdateStatus) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LegacyWafUpdateStatus) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LegacyWafUpdateStatus) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *LegacyWafUpdateStatus) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetData

`func (o *LegacyWafUpdateStatus) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *LegacyWafUpdateStatus) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *LegacyWafUpdateStatus) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *LegacyWafUpdateStatus) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMessage

`func (o *LegacyWafUpdateStatus) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *LegacyWafUpdateStatus) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *LegacyWafUpdateStatus) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *LegacyWafUpdateStatus) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *LegacyWafUpdateStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LegacyWafUpdateStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LegacyWafUpdateStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *LegacyWafUpdateStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *LegacyWafUpdateStatus) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *LegacyWafUpdateStatus) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *LegacyWafUpdateStatus) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *LegacyWafUpdateStatus) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
