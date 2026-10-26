# CreateReportScheduleSpecification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReportType** | **string** | The report type. Refer to [Report Type Values](https://developer-docs.amazon.com/sp-api/docs/report-type-values) for more information. | 
**MarketplaceIds** | **[]string** | A list of marketplace identifiers for the report schedule. | 
**ReportOptions** | Pointer to **map[string]string** | Additional information passed to reports. This varies by report type. | [optional] 
**Period** | **string** | One of a set of predefined &lt;a href&#x3D;&#39;https://developer-docs.amazon.com/sp-api/docs/iso-8601&#39;&gt;ISO 8601&lt;/a&gt; periods that specifies how often a report should be created. | 
**NextReportCreationTime** | Pointer to **time.Time** | The date and time when the schedule will create its next report, in &lt;a href&#x3D;&#39;https://developer-docs.amazon.com/sp-api/docs/iso-8601&#39;&gt;ISO 8601&lt;/a&gt; date time format. | [optional] 

## Methods

### NewCreateReportScheduleSpecification

`func NewCreateReportScheduleSpecification(reportType string, marketplaceIds []string, period string, ) *CreateReportScheduleSpecification`

NewCreateReportScheduleSpecification instantiates a new CreateReportScheduleSpecification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateReportScheduleSpecificationWithDefaults

`func NewCreateReportScheduleSpecificationWithDefaults() *CreateReportScheduleSpecification`

NewCreateReportScheduleSpecificationWithDefaults instantiates a new CreateReportScheduleSpecification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReportType

`func (o *CreateReportScheduleSpecification) GetReportType() string`

GetReportType returns the ReportType field if non-nil, zero value otherwise.

### GetReportTypeOk

`func (o *CreateReportScheduleSpecification) GetReportTypeOk() (*string, bool)`

GetReportTypeOk returns a tuple with the ReportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportType

`func (o *CreateReportScheduleSpecification) SetReportType(v string)`

SetReportType sets ReportType field to given value.


### GetMarketplaceIds

`func (o *CreateReportScheduleSpecification) GetMarketplaceIds() []string`

GetMarketplaceIds returns the MarketplaceIds field if non-nil, zero value otherwise.

### GetMarketplaceIdsOk

`func (o *CreateReportScheduleSpecification) GetMarketplaceIdsOk() (*[]string, bool)`

GetMarketplaceIdsOk returns a tuple with the MarketplaceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplaceIds

`func (o *CreateReportScheduleSpecification) SetMarketplaceIds(v []string)`

SetMarketplaceIds sets MarketplaceIds field to given value.


### GetReportOptions

`func (o *CreateReportScheduleSpecification) GetReportOptions() map[string]string`

GetReportOptions returns the ReportOptions field if non-nil, zero value otherwise.

### GetReportOptionsOk

`func (o *CreateReportScheduleSpecification) GetReportOptionsOk() (*map[string]string, bool)`

GetReportOptionsOk returns a tuple with the ReportOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportOptions

`func (o *CreateReportScheduleSpecification) SetReportOptions(v map[string]string)`

SetReportOptions sets ReportOptions field to given value.

### HasReportOptions

`func (o *CreateReportScheduleSpecification) HasReportOptions() bool`

HasReportOptions returns a boolean if a field has been set.

### GetPeriod

`func (o *CreateReportScheduleSpecification) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *CreateReportScheduleSpecification) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *CreateReportScheduleSpecification) SetPeriod(v string)`

SetPeriod sets Period field to given value.


### GetNextReportCreationTime

`func (o *CreateReportScheduleSpecification) GetNextReportCreationTime() time.Time`

GetNextReportCreationTime returns the NextReportCreationTime field if non-nil, zero value otherwise.

### GetNextReportCreationTimeOk

`func (o *CreateReportScheduleSpecification) GetNextReportCreationTimeOk() (*time.Time, bool)`

GetNextReportCreationTimeOk returns a tuple with the NextReportCreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextReportCreationTime

`func (o *CreateReportScheduleSpecification) SetNextReportCreationTime(v time.Time)`

SetNextReportCreationTime sets NextReportCreationTime field to given value.

### HasNextReportCreationTime

`func (o *CreateReportScheduleSpecification) HasNextReportCreationTime() bool`

HasNextReportCreationTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


