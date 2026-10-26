# Alias

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AliasId** | **string** | The alternative identifier value that can be used to reference this order. | 
**AliasType** | **string** | The kind of alternative identifier this represents.  **Possible values**: &#x60;SELLER_ORDER_ID&#x60; | 

## Methods

### NewAlias

`func NewAlias(aliasId string, aliasType string, ) *Alias`

NewAlias instantiates a new Alias object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAliasWithDefaults

`func NewAliasWithDefaults() *Alias`

NewAliasWithDefaults instantiates a new Alias object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAliasId

`func (o *Alias) GetAliasId() string`

GetAliasId returns the AliasId field if non-nil, zero value otherwise.

### GetAliasIdOk

`func (o *Alias) GetAliasIdOk() (*string, bool)`

GetAliasIdOk returns a tuple with the AliasId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasId

`func (o *Alias) SetAliasId(v string)`

SetAliasId sets AliasId field to given value.


### GetAliasType

`func (o *Alias) GetAliasType() string`

GetAliasType returns the AliasType field if non-nil, zero value otherwise.

### GetAliasTypeOk

`func (o *Alias) GetAliasTypeOk() (*string, bool)`

GetAliasTypeOk returns a tuple with the AliasType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasType

`func (o *Alias) SetAliasType(v string)`

SetAliasType sets AliasType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


