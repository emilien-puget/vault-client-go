# IdentityLookupEntityRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AliasId** | Pointer to **string** | ID of the alias. | [optional] 
**AliasMountAccessor** | Pointer to **string** | Accessor of the mount to which the alias belongs to. This should be supplied in conjunction with &#39;alias_name&#39;. | [optional] 
**AliasName** | Pointer to **string** | Name of the alias. This should be supplied in conjunction with &#39;alias_mount_accessor&#39;. | [optional] 
**Id** | Pointer to **string** | ID of the entity. | [optional] 
**Name** | Pointer to **string** | Name of the entity. | [optional] 

## Methods

### NewIdentityLookupEntityRequest

`func NewIdentityLookupEntityRequest() *IdentityLookupEntityRequest`

NewIdentityLookupEntityRequest instantiates a new IdentityLookupEntityRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityLookupEntityRequestWithDefaults

`func NewIdentityLookupEntityRequestWithDefaults() *IdentityLookupEntityRequest`

NewIdentityLookupEntityRequestWithDefaults instantiates a new IdentityLookupEntityRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAliasId

`func (o *IdentityLookupEntityRequest) GetAliasId() string`

GetAliasId returns the AliasId field if non-nil, zero value otherwise.

### GetAliasIdOk

`func (o *IdentityLookupEntityRequest) GetAliasIdOk() (*string, bool)`

GetAliasIdOk returns a tuple with the AliasId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasId

`func (o *IdentityLookupEntityRequest) SetAliasId(v string)`

SetAliasId sets AliasId field to given value.

### HasAliasId

`func (o *IdentityLookupEntityRequest) HasAliasId() bool`

HasAliasId returns a boolean if a field has been set.

### GetAliasMountAccessor

`func (o *IdentityLookupEntityRequest) GetAliasMountAccessor() string`

GetAliasMountAccessor returns the AliasMountAccessor field if non-nil, zero value otherwise.

### GetAliasMountAccessorOk

`func (o *IdentityLookupEntityRequest) GetAliasMountAccessorOk() (*string, bool)`

GetAliasMountAccessorOk returns a tuple with the AliasMountAccessor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasMountAccessor

`func (o *IdentityLookupEntityRequest) SetAliasMountAccessor(v string)`

SetAliasMountAccessor sets AliasMountAccessor field to given value.

### HasAliasMountAccessor

`func (o *IdentityLookupEntityRequest) HasAliasMountAccessor() bool`

HasAliasMountAccessor returns a boolean if a field has been set.

### GetAliasName

`func (o *IdentityLookupEntityRequest) GetAliasName() string`

GetAliasName returns the AliasName field if non-nil, zero value otherwise.

### GetAliasNameOk

`func (o *IdentityLookupEntityRequest) GetAliasNameOk() (*string, bool)`

GetAliasNameOk returns a tuple with the AliasName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasName

`func (o *IdentityLookupEntityRequest) SetAliasName(v string)`

SetAliasName sets AliasName field to given value.

### HasAliasName

`func (o *IdentityLookupEntityRequest) HasAliasName() bool`

HasAliasName returns a boolean if a field has been set.

### GetId

`func (o *IdentityLookupEntityRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IdentityLookupEntityRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IdentityLookupEntityRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IdentityLookupEntityRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *IdentityLookupEntityRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IdentityLookupEntityRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IdentityLookupEntityRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IdentityLookupEntityRequest) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

