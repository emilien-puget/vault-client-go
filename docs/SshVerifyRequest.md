# SshVerifyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Otp** | Pointer to **string** | [Required] One-Time-Key that needs to be validated | [optional] 

## Methods

### NewSshVerifyRequest

`func NewSshVerifyRequest() *SshVerifyRequest`

NewSshVerifyRequest instantiates a new SshVerifyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSshVerifyRequestWithDefaults

`func NewSshVerifyRequestWithDefaults() *SshVerifyRequest`

NewSshVerifyRequestWithDefaults instantiates a new SshVerifyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOtp

`func (o *SshVerifyRequest) GetOtp() string`

GetOtp returns the Otp field if non-nil, zero value otherwise.

### GetOtpOk

`func (o *SshVerifyRequest) GetOtpOk() (*string, bool)`

GetOtpOk returns a tuple with the Otp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtp

`func (o *SshVerifyRequest) SetOtp(v string)`

SetOtp sets Otp field to given value.

### HasOtp

`func (o *SshVerifyRequest) HasOtp() bool`

HasOtp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

