# ItemDisplayGroupSalesRank

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebsiteDisplayGroup** | **string** | Name of the website display group that is associated with the sales rank | 
**Title** | **string** | Name of the sales rank. | 
**Link** | Pointer to **string** | Corresponding Amazon retail website URL for the sales rank. | [optional] 
**Rank** | **int32** | Sales rank. | 

## Methods

### NewItemDisplayGroupSalesRank

`func NewItemDisplayGroupSalesRank(websiteDisplayGroup string, title string, rank int32, ) *ItemDisplayGroupSalesRank`

NewItemDisplayGroupSalesRank instantiates a new ItemDisplayGroupSalesRank object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemDisplayGroupSalesRankWithDefaults

`func NewItemDisplayGroupSalesRankWithDefaults() *ItemDisplayGroupSalesRank`

NewItemDisplayGroupSalesRankWithDefaults instantiates a new ItemDisplayGroupSalesRank object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebsiteDisplayGroup

`func (o *ItemDisplayGroupSalesRank) GetWebsiteDisplayGroup() string`

GetWebsiteDisplayGroup returns the WebsiteDisplayGroup field if non-nil, zero value otherwise.

### GetWebsiteDisplayGroupOk

`func (o *ItemDisplayGroupSalesRank) GetWebsiteDisplayGroupOk() (*string, bool)`

GetWebsiteDisplayGroupOk returns a tuple with the WebsiteDisplayGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteDisplayGroup

`func (o *ItemDisplayGroupSalesRank) SetWebsiteDisplayGroup(v string)`

SetWebsiteDisplayGroup sets WebsiteDisplayGroup field to given value.


### GetTitle

`func (o *ItemDisplayGroupSalesRank) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ItemDisplayGroupSalesRank) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ItemDisplayGroupSalesRank) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetLink

`func (o *ItemDisplayGroupSalesRank) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *ItemDisplayGroupSalesRank) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *ItemDisplayGroupSalesRank) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *ItemDisplayGroupSalesRank) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetRank

`func (o *ItemDisplayGroupSalesRank) GetRank() int32`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *ItemDisplayGroupSalesRank) GetRankOk() (*int32, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *ItemDisplayGroupSalesRank) SetRank(v int32)`

SetRank sets Rank field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


