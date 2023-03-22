# ServiceDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**DeletedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**UpdatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**Comment** | Pointer to **NullableString** | A freeform descriptive note. | [optional] 
**Name** | Pointer to **string** | The name of the service. | [optional] 
**CustomerID** | Pointer to **string** | Alphanumeric string identifying the customer. | [optional] 
**Type** | Pointer to **string** | The type of this service. | [optional] 
**ID** | Pointer to **string** |  | [optional] [readonly] 
**PublishKey** | Pointer to **string** | Unused at this time. | [optional] 
**Paused** | Pointer to **bool** | Whether the service is paused. Services are paused due to a lack of traffic for an extended period of time. Services are resumed either when a draft version is activated or a locked version is cloned and reactivated. | [optional] 
**Versions** | Pointer to [**[]SchemasVersionResponse**](SchemasVersionResponse.md) | A list of [versions](/reference/api/services/version/) associated with the service. | [optional] 
**ActiveVersion** | Pointer to [**NullableServiceVersionDetailOrNull**](ServiceVersionDetailOrNull.md) |  | [optional] 
**Version** | Pointer to [**ServiceVersionDetail**](ServiceVersionDetail.md) |  | [optional] 

## Methods

### NewServiceDetail

`func NewServiceDetail() *ServiceDetail`

NewServiceDetail instantiates a new ServiceDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceDetailWithDefaults

`func NewServiceDetailWithDefaults() *ServiceDetail`

NewServiceDetailWithDefaults instantiates a new ServiceDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ServiceDetail) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ServiceDetail) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ServiceDetail) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ServiceDetail) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *ServiceDetail) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *ServiceDetail) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetDeletedAt

`func (o *ServiceDetail) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *ServiceDetail) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *ServiceDetail) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *ServiceDetail) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *ServiceDetail) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *ServiceDetail) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetUpdatedAt

`func (o *ServiceDetail) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ServiceDetail) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ServiceDetail) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ServiceDetail) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *ServiceDetail) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *ServiceDetail) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetComment

`func (o *ServiceDetail) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ServiceDetail) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ServiceDetail) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ServiceDetail) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *ServiceDetail) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *ServiceDetail) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetName

`func (o *ServiceDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceDetail) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServiceDetail) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCustomerID

`func (o *ServiceDetail) GetCustomerID() string`

GetCustomerID returns the CustomerID field if non-nil, zero value otherwise.

### GetCustomerIDOk

`func (o *ServiceDetail) GetCustomerIDOk() (*string, bool)`

GetCustomerIDOk returns a tuple with the CustomerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerID

`func (o *ServiceDetail) SetCustomerID(v string)`

SetCustomerID sets CustomerID field to given value.

### HasCustomerID

`func (o *ServiceDetail) HasCustomerID() bool`

HasCustomerID returns a boolean if a field has been set.

### GetType

`func (o *ServiceDetail) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServiceDetail) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServiceDetail) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ServiceDetail) HasType() bool`

HasType returns a boolean if a field has been set.

### GetID

`func (o *ServiceDetail) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *ServiceDetail) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *ServiceDetail) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *ServiceDetail) HasID() bool`

HasID returns a boolean if a field has been set.

### GetPublishKey

`func (o *ServiceDetail) GetPublishKey() string`

GetPublishKey returns the PublishKey field if non-nil, zero value otherwise.

### GetPublishKeyOk

`func (o *ServiceDetail) GetPublishKeyOk() (*string, bool)`

GetPublishKeyOk returns a tuple with the PublishKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishKey

`func (o *ServiceDetail) SetPublishKey(v string)`

SetPublishKey sets PublishKey field to given value.

### HasPublishKey

`func (o *ServiceDetail) HasPublishKey() bool`

HasPublishKey returns a boolean if a field has been set.

### GetPaused

`func (o *ServiceDetail) GetPaused() bool`

GetPaused returns the Paused field if non-nil, zero value otherwise.

### GetPausedOk

`func (o *ServiceDetail) GetPausedOk() (*bool, bool)`

GetPausedOk returns a tuple with the Paused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaused

`func (o *ServiceDetail) SetPaused(v bool)`

SetPaused sets Paused field to given value.

### HasPaused

`func (o *ServiceDetail) HasPaused() bool`

HasPaused returns a boolean if a field has been set.

### GetVersions

`func (o *ServiceDetail) GetVersions() []SchemasVersionResponse`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *ServiceDetail) GetVersionsOk() (*[]SchemasVersionResponse, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *ServiceDetail) SetVersions(v []SchemasVersionResponse)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *ServiceDetail) HasVersions() bool`

HasVersions returns a boolean if a field has been set.

### GetActiveVersion

`func (o *ServiceDetail) GetActiveVersion() ServiceVersionDetailOrNull`

GetActiveVersion returns the ActiveVersion field if non-nil, zero value otherwise.

### GetActiveVersionOk

`func (o *ServiceDetail) GetActiveVersionOk() (*ServiceVersionDetailOrNull, bool)`

GetActiveVersionOk returns a tuple with the ActiveVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveVersion

`func (o *ServiceDetail) SetActiveVersion(v ServiceVersionDetailOrNull)`

SetActiveVersion sets ActiveVersion field to given value.

### HasActiveVersion

`func (o *ServiceDetail) HasActiveVersion() bool`

HasActiveVersion returns a boolean if a field has been set.

### SetActiveVersionNil

`func (o *ServiceDetail) SetActiveVersionNil(b bool)`

 SetActiveVersionNil sets the value for ActiveVersion to be an explicit nil

### UnsetActiveVersion
`func (o *ServiceDetail) UnsetActiveVersion()`

UnsetActiveVersion ensures that no value is present for ActiveVersion, not even an explicit nil
### GetVersion

`func (o *ServiceDetail) GetVersion() ServiceVersionDetail`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ServiceDetail) GetVersionOk() (*ServiceVersionDetail, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ServiceDetail) SetVersion(v ServiceVersionDetail)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ServiceDetail) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
