# CreateReportSpecification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReportOptions** | Pointer to **map[string]string** | Additional information passed to reports. This varies by report type. | [optional] 
**ReportType** | **string** | The report type. Refer to [Report Type Values](https://developer-docs.amazon.com/sp-api/docs/report-type-values) for more information. | 
**DataStartTime** | Pointer to **time.Time** | The start of a date and time range, in &lt;a href&#x3D;&#39;https://developer-docs.amazon.com/sp-api/docs/iso-8601&#39;&gt;ISO 8601&lt;/a&gt; date time format, used for selecting the data to report. The default is now. The value must be prior to or equal to the current date and time. Not all report types make use of this. | [optional] 
**DataEndTime** | Pointer to **time.Time** | The end of a date and time range, in &lt;a href&#x3D;&#39;https://developer-docs.amazon.com/sp-api/docs/iso-8601&#39;&gt;ISO 8601&lt;/a&gt; date time format, used for selecting the data to report. The default is now. The value must be prior to or equal to the current date and time. Not all report types make use of this. | [optional] 
**MarketplaceIds** | **[]string** | A list of marketplace identifiers. The report document&#39;s contents will contain data for all of the specified marketplaces, unless the report type indicates otherwise. | 

## Methods

### NewCreateReportSpecification

`func NewCreateReportSpecification(reportType string, marketplaceIds []string, ) *CreateReportSpecification`

NewCreateReportSpecification instantiates a new CreateReportSpecification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateReportSpecificationWithDefaults

`func NewCreateReportSpecificationWithDefaults() *CreateReportSpecification`

NewCreateReportSpecificationWithDefaults instantiates a new CreateReportSpecification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReportOptions

`func (o *CreateReportSpecification) GetReportOptions() map[string]string`

GetReportOptions returns the ReportOptions field if non-nil, zero value otherwise.

### GetReportOptionsOk

`func (o *CreateReportSpecification) GetReportOptionsOk() (*map[string]string, bool)`

GetReportOptionsOk returns a tuple with the ReportOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportOptions

`func (o *CreateReportSpecification) SetReportOptions(v map[string]string)`

SetReportOptions sets ReportOptions field to given value.

### HasReportOptions

`func (o *CreateReportSpecification) HasReportOptions() bool`

HasReportOptions returns a boolean if a field has been set.

### GetReportType

`func (o *CreateReportSpecification) GetReportType() string`

GetReportType returns the ReportType field if non-nil, zero value otherwise.

### GetReportTypeOk

`func (o *CreateReportSpecification) GetReportTypeOk() (*string, bool)`

GetReportTypeOk returns a tuple with the ReportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportType

`func (o *CreateReportSpecification) SetReportType(v string)`

SetReportType sets ReportType field to given value.


### GetDataStartTime

`func (o *CreateReportSpecification) GetDataStartTime() time.Time`

GetDataStartTime returns the DataStartTime field if non-nil, zero value otherwise.

### GetDataStartTimeOk

`func (o *CreateReportSpecification) GetDataStartTimeOk() (*time.Time, bool)`

GetDataStartTimeOk returns a tuple with the DataStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataStartTime

`func (o *CreateReportSpecification) SetDataStartTime(v time.Time)`

SetDataStartTime sets DataStartTime field to given value.

### HasDataStartTime

`func (o *CreateReportSpecification) HasDataStartTime() bool`

HasDataStartTime returns a boolean if a field has been set.

### GetDataEndTime

`func (o *CreateReportSpecification) GetDataEndTime() time.Time`

GetDataEndTime returns the DataEndTime field if non-nil, zero value otherwise.

### GetDataEndTimeOk

`func (o *CreateReportSpecification) GetDataEndTimeOk() (*time.Time, bool)`

GetDataEndTimeOk returns a tuple with the DataEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataEndTime

`func (o *CreateReportSpecification) SetDataEndTime(v time.Time)`

SetDataEndTime sets DataEndTime field to given value.

### HasDataEndTime

`func (o *CreateReportSpecification) HasDataEndTime() bool`

HasDataEndTime returns a boolean if a field has been set.

### GetMarketplaceIds

`func (o *CreateReportSpecification) GetMarketplaceIds() []string`

GetMarketplaceIds returns the MarketplaceIds field if non-nil, zero value otherwise.

### GetMarketplaceIdsOk

`func (o *CreateReportSpecification) GetMarketplaceIdsOk() (*[]string, bool)`

GetMarketplaceIdsOk returns a tuple with the MarketplaceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplaceIds

`func (o *CreateReportSpecification) SetMarketplaceIds(v []string)`

SetMarketplaceIds sets MarketplaceIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


