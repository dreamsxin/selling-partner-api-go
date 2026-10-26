# PrimaryContact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The full name of the seller&#39;s primary contact. | 
**Address** | [**Address**](Address.md) |  | 
**NonLatinName** | Pointer to **string** | The non-Latin script version of the primary contact&#39;s name, if applicable. | [optional] 

## Methods

### NewPrimaryContact

`func NewPrimaryContact(name string, address Address, ) *PrimaryContact`

NewPrimaryContact instantiates a new PrimaryContact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrimaryContactWithDefaults

`func NewPrimaryContactWithDefaults() *PrimaryContact`

NewPrimaryContactWithDefaults instantiates a new PrimaryContact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PrimaryContact) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PrimaryContact) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PrimaryContact) SetName(v string)`

SetName sets Name field to given value.


### GetAddress

`func (o *PrimaryContact) GetAddress() Address`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *PrimaryContact) GetAddressOk() (*Address, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *PrimaryContact) SetAddress(v Address)`

SetAddress sets Address field to given value.


### GetNonLatinName

`func (o *PrimaryContact) GetNonLatinName() string`

GetNonLatinName returns the NonLatinName field if non-nil, zero value otherwise.

### GetNonLatinNameOk

`func (o *PrimaryContact) GetNonLatinNameOk() (*string, bool)`

GetNonLatinNameOk returns a tuple with the NonLatinName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonLatinName

`func (o *PrimaryContact) SetNonLatinName(v string)`

SetNonLatinName sets NonLatinName field to given value.

### HasNonLatinName

`func (o *PrimaryContact) HasNonLatinName() bool`

HasNonLatinName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


