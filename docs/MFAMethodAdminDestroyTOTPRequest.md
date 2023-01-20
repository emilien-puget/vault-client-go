# MFAMethodAdminDestroyTOTPRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityId** | **string** | Identifier of the entity from which the MFA method secret needs to be removed. | 
**MethodId** | **string** | The unique identifier for this MFA method. | 

## Methods

### NewMFAMethodAdminDestroyTOTPRequest

`func NewMFAMethodAdminDestroyTOTPRequest(entityId string, methodId string, ) *MFAMethodAdminDestroyTOTPRequest`

NewMFAMethodAdminDestroyTOTPRequest instantiates a new MFAMethodAdminDestroyTOTPRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMFAMethodAdminDestroyTOTPRequestWithDefaults

`func NewMFAMethodAdminDestroyTOTPRequestWithDefaults() *MFAMethodAdminDestroyTOTPRequest`

NewMFAMethodAdminDestroyTOTPRequestWithDefaults instantiates a new MFAMethodAdminDestroyTOTPRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityId

`func (o *MFAMethodAdminDestroyTOTPRequest) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *MFAMethodAdminDestroyTOTPRequest) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *MFAMethodAdminDestroyTOTPRequest) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.


### GetMethodId

`func (o *MFAMethodAdminDestroyTOTPRequest) GetMethodId() string`

GetMethodId returns the MethodId field if non-nil, zero value otherwise.

### GetMethodIdOk

`func (o *MFAMethodAdminDestroyTOTPRequest) GetMethodIdOk() (*string, bool)`

GetMethodIdOk returns a tuple with the MethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodId

`func (o *MFAMethodAdminDestroyTOTPRequest) SetMethodId(v string)`

SetMethodId sets MethodId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

