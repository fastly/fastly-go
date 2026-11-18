# EventAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Admin** | Pointer to **bool** | Indicates if event was performed by Fastly. | [optional] 
**CreatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**CustomerID** | Pointer to **string** |  | [optional] [readonly] 
**Description** | Pointer to **string** | Description of the event. | [optional] 
**EventType** | Pointer to **string** | Type of event. Can be used with `filter[event_type]` | [optional] 
**IP** | Pointer to **string** | IP addresses that the event was requested from. | [optional] 
**Metadata** | Pointer to **map[string]map[string]any** | Hash of key value pairs of additional information. | [optional] 
**ServiceID** | Pointer to **string** |  | [optional] [readonly] 
**UserID** | Pointer to **string** |  | [optional] [readonly] 
**TokenID** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewEventAttributes

`func NewEventAttributes() *EventAttributes`

NewEventAttributes instantiates a new EventAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventAttributesWithDefaults

`func NewEventAttributesWithDefaults() *EventAttributes`

NewEventAttributesWithDefaults instantiates a new EventAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdmin

`func (o *EventAttributes) GetAdmin() bool`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *EventAttributes) GetAdminOk() (*bool, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *EventAttributes) SetAdmin(v bool)`

SetAdmin sets Admin field to given value.

### HasAdmin

`func (o *EventAttributes) HasAdmin() bool`

HasAdmin returns a boolean if a field has been set.

### GetCreatedAt

`func (o *EventAttributes) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EventAttributes) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EventAttributes) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EventAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *EventAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *EventAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetCustomerID

`func (o *EventAttributes) GetCustomerID() string`

GetCustomerID returns the CustomerID field if non-nil, zero value otherwise.

### GetCustomerIDOk

`func (o *EventAttributes) GetCustomerIDOk() (*string, bool)`

GetCustomerIDOk returns a tuple with the CustomerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerID

`func (o *EventAttributes) SetCustomerID(v string)`

SetCustomerID sets CustomerID field to given value.

### HasCustomerID

`func (o *EventAttributes) HasCustomerID() bool`

HasCustomerID returns a boolean if a field has been set.

### GetDescription

`func (o *EventAttributes) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EventAttributes) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EventAttributes) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EventAttributes) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEventType

`func (o *EventAttributes) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *EventAttributes) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *EventAttributes) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *EventAttributes) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetIP

`func (o *EventAttributes) GetIP() string`

GetIP returns the IP field if non-nil, zero value otherwise.

### GetIPOk

`func (o *EventAttributes) GetIPOk() (*string, bool)`

GetIPOk returns a tuple with the IP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIP

`func (o *EventAttributes) SetIP(v string)`

SetIP sets IP field to given value.

### HasIP

`func (o *EventAttributes) HasIP() bool`

HasIP returns a boolean if a field has been set.

### GetMetadata

`func (o *EventAttributes) GetMetadata() map[string]map[string]any`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *EventAttributes) GetMetadataOk() (*map[string]map[string]any, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *EventAttributes) SetMetadata(v map[string]map[string]any)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *EventAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetServiceID

`func (o *EventAttributes) GetServiceID() string`

GetServiceID returns the ServiceID field if non-nil, zero value otherwise.

### GetServiceIDOk

`func (o *EventAttributes) GetServiceIDOk() (*string, bool)`

GetServiceIDOk returns a tuple with the ServiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceID

`func (o *EventAttributes) SetServiceID(v string)`

SetServiceID sets ServiceID field to given value.

### HasServiceID

`func (o *EventAttributes) HasServiceID() bool`

HasServiceID returns a boolean if a field has been set.

### GetUserID

`func (o *EventAttributes) GetUserID() string`

GetUserID returns the UserID field if non-nil, zero value otherwise.

### GetUserIDOk

`func (o *EventAttributes) GetUserIDOk() (*string, bool)`

GetUserIDOk returns a tuple with the UserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserID

`func (o *EventAttributes) SetUserID(v string)`

SetUserID sets UserID field to given value.

### HasUserID

`func (o *EventAttributes) HasUserID() bool`

HasUserID returns a boolean if a field has been set.

### GetTokenID

`func (o *EventAttributes) GetTokenID() string`

GetTokenID returns the TokenID field if non-nil, zero value otherwise.

### GetTokenIDOk

`func (o *EventAttributes) GetTokenIDOk() (*string, bool)`

GetTokenIDOk returns a tuple with the TokenID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenID

`func (o *EventAttributes) SetTokenID(v string)`

SetTokenID sets TokenID field to given value.

### HasTokenID

`func (o *EventAttributes) HasTokenID() bool`

HasTokenID returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
