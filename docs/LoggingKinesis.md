# LoggingKinesis

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name for the real-time logging configuration. | [optional] 
**Placement** | Pointer to [**NullableLoggingPlacement**](LoggingPlacement.md) |  | [optional] 
**FormatVersion** | Pointer to [**LoggingFormatVersion**](LoggingFormatVersion.md) |  | [optional] [default to LOGGINGFORMATVERSION_v2]
**Format** | Pointer to **string** | A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats). Must produce valid JSON that Kinesis can ingest. | [optional] [default to "{\"timestamp\":\"%{begin:%Y-%m-%dT%H:%M:%S}t\",\"time_elapsed\":\"%{time.elapsed.usec}V\",\"is_tls\":\"%{if(req.is_ssl, \\\"true\\\", \\\"false\\\")}V\",\"client_ip\":\"%{req.http.Fastly-Client-IP}V\",\"geo_city\":\"%{client.geo.city}V\",\"geo_country_code\":\"%{client.geo.country_code}V\",\"request\":\"%{req.request}V\",\"host\":\"%{req.http.Fastly-Orig-Host}V\",\"url\":\"%{json.escape(req.url)}V\",\"request_referer\":\"%{json.escape(req.http.Referer)}V\",\"request_user_agent\":\"%{json.escape(req.http.User-Agent)}V\",\"request_accept_language\":\"%{json.escape(req.http.Accept-Language)}V\",\"request_accept_charset\":\"%{json.escape(req.http.Accept-Charset)}V\",\"cache_status\":\"%{regsub(fastly_info.state, \\\"^(HIT-(SYNTH)|(HITPASS|HIT|MISS|PASS|ERROR|PIPE)).*\\\", \\\"\\\\2\\\\3\\\") }V\"}"]
**Topic** | Pointer to **string** | The Amazon Kinesis stream to send logs to. Required. | [optional] 
**Region** | Pointer to [**AwsRegion**](AwsRegion.md) |  | [optional] 
**SecretKey** | Pointer to **NullableString** | The secret key associated with the target Amazon Kinesis stream. Not required if `iam_role` is specified. | [optional] 
**AccessKey** | Pointer to **NullableString** | The access key associated with the target Amazon Kinesis stream. Not required if `iam_role` is specified. | [optional] 
**IamRole** | Pointer to **NullableString** | The ARN for an IAM role granting Fastly access to the target Amazon Kinesis stream. Not required if `access_key` and `secret_key` are provided. | [optional] 

## Methods

### NewLoggingKinesis

`func NewLoggingKinesis() *LoggingKinesis`

NewLoggingKinesis instantiates a new LoggingKinesis object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoggingKinesisWithDefaults

`func NewLoggingKinesisWithDefaults() *LoggingKinesis`

NewLoggingKinesisWithDefaults instantiates a new LoggingKinesis object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *LoggingKinesis) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LoggingKinesis) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LoggingKinesis) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LoggingKinesis) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlacement

`func (o *LoggingKinesis) GetPlacement() LoggingPlacement`

GetPlacement returns the Placement field if non-nil, zero value otherwise.

### GetPlacementOk

`func (o *LoggingKinesis) GetPlacementOk() (*LoggingPlacement, bool)`

GetPlacementOk returns a tuple with the Placement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacement

`func (o *LoggingKinesis) SetPlacement(v LoggingPlacement)`

SetPlacement sets Placement field to given value.

### HasPlacement

`func (o *LoggingKinesis) HasPlacement() bool`

HasPlacement returns a boolean if a field has been set.

### SetPlacementNil

`func (o *LoggingKinesis) SetPlacementNil(b bool)`

 SetPlacementNil sets the value for Placement to be an explicit nil

### UnsetPlacement
`func (o *LoggingKinesis) UnsetPlacement()`

UnsetPlacement ensures that no value is present for Placement, not even an explicit nil
### GetFormatVersion

`func (o *LoggingKinesis) GetFormatVersion() LoggingFormatVersion`

GetFormatVersion returns the FormatVersion field if non-nil, zero value otherwise.

### GetFormatVersionOk

`func (o *LoggingKinesis) GetFormatVersionOk() (*LoggingFormatVersion, bool)`

GetFormatVersionOk returns a tuple with the FormatVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatVersion

`func (o *LoggingKinesis) SetFormatVersion(v LoggingFormatVersion)`

SetFormatVersion sets FormatVersion field to given value.

### HasFormatVersion

`func (o *LoggingKinesis) HasFormatVersion() bool`

HasFormatVersion returns a boolean if a field has been set.

### GetFormat

`func (o *LoggingKinesis) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *LoggingKinesis) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *LoggingKinesis) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *LoggingKinesis) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetTopic

`func (o *LoggingKinesis) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *LoggingKinesis) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *LoggingKinesis) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *LoggingKinesis) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### GetRegion

`func (o *LoggingKinesis) GetRegion() AwsRegion`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *LoggingKinesis) GetRegionOk() (*AwsRegion, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *LoggingKinesis) SetRegion(v AwsRegion)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *LoggingKinesis) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetSecretKey

`func (o *LoggingKinesis) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *LoggingKinesis) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *LoggingKinesis) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *LoggingKinesis) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### SetSecretKeyNil

`func (o *LoggingKinesis) SetSecretKeyNil(b bool)`

 SetSecretKeyNil sets the value for SecretKey to be an explicit nil

### UnsetSecretKey
`func (o *LoggingKinesis) UnsetSecretKey()`

UnsetSecretKey ensures that no value is present for SecretKey, not even an explicit nil
### GetAccessKey

`func (o *LoggingKinesis) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *LoggingKinesis) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *LoggingKinesis) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *LoggingKinesis) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### SetAccessKeyNil

`func (o *LoggingKinesis) SetAccessKeyNil(b bool)`

 SetAccessKeyNil sets the value for AccessKey to be an explicit nil

### UnsetAccessKey
`func (o *LoggingKinesis) UnsetAccessKey()`

UnsetAccessKey ensures that no value is present for AccessKey, not even an explicit nil
### GetIamRole

`func (o *LoggingKinesis) GetIamRole() string`

GetIamRole returns the IamRole field if non-nil, zero value otherwise.

### GetIamRoleOk

`func (o *LoggingKinesis) GetIamRoleOk() (*string, bool)`

GetIamRoleOk returns a tuple with the IamRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIamRole

`func (o *LoggingKinesis) SetIamRole(v string)`

SetIamRole sets IamRole field to given value.

### HasIamRole

`func (o *LoggingKinesis) HasIamRole() bool`

HasIamRole returns a boolean if a field has been set.

### SetIamRoleNil

`func (o *LoggingKinesis) SetIamRoleNil(b bool)`

 SetIamRoleNil sets the value for IamRole to be an explicit nil

### UnsetIamRole
`func (o *LoggingKinesis) UnsetIamRole()`

UnsetIamRole ensures that no value is present for IamRole, not even an explicit nil

[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
