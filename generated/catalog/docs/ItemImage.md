# ItemImage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Variant** | **string** | Variant of the image, such as &#x60;MAIN&#x60; or &#x60;PT01&#x60;. | 
**Link** | **string** | URL for the image. | 
**Height** | **int32** | Height of the image in pixels. | 
**Width** | **int32** | Width of the image in pixels. | 

## Methods

### NewItemImage

`func NewItemImage(variant string, link string, height int32, width int32, ) *ItemImage`

NewItemImage instantiates a new ItemImage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemImageWithDefaults

`func NewItemImageWithDefaults() *ItemImage`

NewItemImageWithDefaults instantiates a new ItemImage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVariant

`func (o *ItemImage) GetVariant() string`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *ItemImage) GetVariantOk() (*string, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *ItemImage) SetVariant(v string)`

SetVariant sets Variant field to given value.


### GetLink

`func (o *ItemImage) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *ItemImage) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *ItemImage) SetLink(v string)`

SetLink sets Link field to given value.


### GetHeight

`func (o *ItemImage) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *ItemImage) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *ItemImage) SetHeight(v int32)`

SetHeight sets Height field to given value.


### GetWidth

`func (o *ItemImage) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *ItemImage) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *ItemImage) SetWidth(v int32)`

SetWidth sets Width field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


