# Website

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique website identifier | [optional] 
**Domain** | Pointer to **string** | Website domain | [optional] 
**PageIds** | Pointer to **[]string** | IDs of pages associated with this website | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewWebsite

`func NewWebsite() *Website`

NewWebsite instantiates a new Website object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebsiteWithDefaults

`func NewWebsiteWithDefaults() *Website`

NewWebsiteWithDefaults instantiates a new Website object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Website) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Website) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Website) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Website) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDomain

`func (o *Website) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *Website) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *Website) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *Website) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetPageIds

`func (o *Website) GetPageIds() []string`

GetPageIds returns the PageIds field if non-nil, zero value otherwise.

### GetPageIdsOk

`func (o *Website) GetPageIdsOk() (*[]string, bool)`

GetPageIdsOk returns a tuple with the PageIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageIds

`func (o *Website) SetPageIds(v []string)`

SetPageIds sets PageIds field to given value.

### HasPageIds

`func (o *Website) HasPageIds() bool`

HasPageIds returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Website) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Website) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Website) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Website) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Website) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Website) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Website) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Website) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


