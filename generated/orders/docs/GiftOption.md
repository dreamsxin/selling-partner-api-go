# GiftOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GiftMessage** | Pointer to **string** | Personal message from the buyer to be included with the gift-wrapped item. | [optional] 
**GiftWrapLevel** | Pointer to **string** | Type or quality level of gift wrapping service selected by the customer. | [optional] 

## Methods

### NewGiftOption

`func NewGiftOption() *GiftOption`

NewGiftOption instantiates a new GiftOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGiftOptionWithDefaults

`func NewGiftOptionWithDefaults() *GiftOption`

NewGiftOptionWithDefaults instantiates a new GiftOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGiftMessage

`func (o *GiftOption) GetGiftMessage() string`

GetGiftMessage returns the GiftMessage field if non-nil, zero value otherwise.

### GetGiftMessageOk

`func (o *GiftOption) GetGiftMessageOk() (*string, bool)`

GetGiftMessageOk returns a tuple with the GiftMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiftMessage

`func (o *GiftOption) SetGiftMessage(v string)`

SetGiftMessage sets GiftMessage field to given value.

### HasGiftMessage

`func (o *GiftOption) HasGiftMessage() bool`

HasGiftMessage returns a boolean if a field has been set.

### GetGiftWrapLevel

`func (o *GiftOption) GetGiftWrapLevel() string`

GetGiftWrapLevel returns the GiftWrapLevel field if non-nil, zero value otherwise.

### GetGiftWrapLevelOk

`func (o *GiftOption) GetGiftWrapLevelOk() (*string, bool)`

GetGiftWrapLevelOk returns a tuple with the GiftWrapLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiftWrapLevel

`func (o *GiftOption) SetGiftWrapLevel(v string)`

SetGiftWrapLevel sets GiftWrapLevel field to given value.

### HasGiftWrapLevel

`func (o *GiftOption) HasGiftWrapLevel() bool`

HasGiftWrapLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


