# PackageMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the Compute@Edge package. | [optional] 
**Description** | Pointer to **string** | Description of the Compute@Edge package. | [optional] 
**Authors** | Pointer to **[]string** | A list of package authors&#39; email addresses. | [optional] 
**Language** | Pointer to **string** | The language of the Compute@Edge package. | [optional] 
**Size** | Pointer to **int32** | Size of the Compute@Edge package in bytes. | [optional] 
**Hashsum** | Pointer to **string** | Hash of the Compute@Edge package. | [optional] 

## Methods

### NewPackageMetadata

`func NewPackageMetadata() *PackageMetadata`

NewPackageMetadata instantiates a new PackageMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackageMetadataWithDefaults

`func NewPackageMetadataWithDefaults() *PackageMetadata`

NewPackageMetadataWithDefaults instantiates a new PackageMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PackageMetadata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PackageMetadata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PackageMetadata) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PackageMetadata) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PackageMetadata) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PackageMetadata) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PackageMetadata) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PackageMetadata) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAuthors

`func (o *PackageMetadata) GetAuthors() []string`

GetAuthors returns the Authors field if non-nil, zero value otherwise.

### GetAuthorsOk

`func (o *PackageMetadata) GetAuthorsOk() (*[]string, bool)`

GetAuthorsOk returns a tuple with the Authors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthors

`func (o *PackageMetadata) SetAuthors(v []string)`

SetAuthors sets Authors field to given value.

### HasAuthors

`func (o *PackageMetadata) HasAuthors() bool`

HasAuthors returns a boolean if a field has been set.

### GetLanguage

`func (o *PackageMetadata) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *PackageMetadata) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *PackageMetadata) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *PackageMetadata) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetSize

`func (o *PackageMetadata) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *PackageMetadata) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *PackageMetadata) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *PackageMetadata) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetHashsum

`func (o *PackageMetadata) GetHashsum() string`

GetHashsum returns the Hashsum field if non-nil, zero value otherwise.

### GetHashsumOk

`func (o *PackageMetadata) GetHashsumOk() (*string, bool)`

GetHashsumOk returns a tuple with the Hashsum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashsum

`func (o *PackageMetadata) SetHashsum(v string)`

SetHashsum sets Hashsum field to given value.

### HasHashsum

`func (o *PackageMetadata) HasHashsum() bool`

HasHashsum returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
