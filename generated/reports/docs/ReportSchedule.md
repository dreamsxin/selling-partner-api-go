# ReportSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReportScheduleId** | **string** | The identifier for the report schedule. This identifier is unique only in combination with a seller ID. | 
**ReportType** | **string** | The report type. Refer to [Report Type Values](https://developer-docs.amazon.com/sp-api/docs/report-type-values) for more information. | 
**MarketplaceIds** | Pointer to **[]string** | A list of marketplace identifiers. The report document&#39;s contents will contain data for all of the specified marketplaces, unless the report type indicates otherwise. | [optional] 
**ReportOptions** | Pointer to **map[string]string** | Additional information passed to reports. This varies by report type. | [optional] 
**Period** | **string** | An &lt;a href&#x3D;&#39;https://developer-docs.amazon.com/sp-api/docs/iso-8601&#39;&gt;ISO 8601&lt;/a&gt; period value that indicates how often a report should be created. | 
**NextReportCreationTime** | Pointer to **time.Time** | The date and time when the schedule will create its next report, in &lt;a href&#x3D;&#39;https://developer-docs.amazon.com/sp-api/docs/iso-8601&#39;&gt;ISO 8601&lt;/a&gt; date time format. | [optional] 

## Methods

### NewReportSchedule

`func NewReportSchedule(reportScheduleId string, reportType string, period string, ) *ReportSchedule`

NewReportSchedule instantiates a new ReportSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportScheduleWithDefaults

`func NewReportScheduleWithDefaults() *ReportSchedule`

NewReportScheduleWithDefaults instantiates a new ReportSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReportScheduleId

`func (o *ReportSchedule) GetReportScheduleId() string`

GetReportScheduleId returns the ReportScheduleId field if non-nil, zero value otherwise.

### GetReportScheduleIdOk

`func (o *ReportSchedule) GetReportScheduleIdOk() (*string, bool)`

GetReportScheduleIdOk returns a tuple with the ReportScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportScheduleId

`func (o *ReportSchedule) SetReportScheduleId(v string)`

SetReportScheduleId sets ReportScheduleId field to given value.


### GetReportType

`func (o *ReportSchedule) GetReportType() string`

GetReportType returns the ReportType field if non-nil, zero value otherwise.

### GetReportTypeOk

`func (o *ReportSchedule) GetReportTypeOk() (*string, bool)`

GetReportTypeOk returns a tuple with the ReportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportType

`func (o *ReportSchedule) SetReportType(v string)`

SetReportType sets ReportType field to given value.


### GetMarketplaceIds

`func (o *ReportSchedule) GetMarketplaceIds() []string`

GetMarketplaceIds returns the MarketplaceIds field if non-nil, zero value otherwise.

### GetMarketplaceIdsOk

`func (o *ReportSchedule) GetMarketplaceIdsOk() (*[]string, bool)`

GetMarketplaceIdsOk returns a tuple with the MarketplaceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplaceIds

`func (o *ReportSchedule) SetMarketplaceIds(v []string)`

SetMarketplaceIds sets MarketplaceIds field to given value.

### HasMarketplaceIds

`func (o *ReportSchedule) HasMarketplaceIds() bool`

HasMarketplaceIds returns a boolean if a field has been set.

### GetReportOptions

`func (o *ReportSchedule) GetReportOptions() map[string]string`

GetReportOptions returns the ReportOptions field if non-nil, zero value otherwise.

### GetReportOptionsOk

`func (o *ReportSchedule) GetReportOptionsOk() (*map[string]string, bool)`

GetReportOptionsOk returns a tuple with the ReportOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportOptions

`func (o *ReportSchedule) SetReportOptions(v map[string]string)`

SetReportOptions sets ReportOptions field to given value.

### HasReportOptions

`func (o *ReportSchedule) HasReportOptions() bool`

HasReportOptions returns a boolean if a field has been set.

### GetPeriod

`func (o *ReportSchedule) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *ReportSchedule) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *ReportSchedule) SetPeriod(v string)`

SetPeriod sets Period field to given value.


### GetNextReportCreationTime

`func (o *ReportSchedule) GetNextReportCreationTime() time.Time`

GetNextReportCreationTime returns the NextReportCreationTime field if non-nil, zero value otherwise.

### GetNextReportCreationTimeOk

`func (o *ReportSchedule) GetNextReportCreationTimeOk() (*time.Time, bool)`

GetNextReportCreationTimeOk returns a tuple with the NextReportCreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextReportCreationTime

`func (o *ReportSchedule) SetNextReportCreationTime(v time.Time)`

SetNextReportCreationTime sets NextReportCreationTime field to given value.

### HasNextReportCreationTime

`func (o *ReportSchedule) HasNextReportCreationTime() bool`

HasNextReportCreationTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


