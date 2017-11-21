# \DatasourcecontrollerApi

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddDataSourceUsingPOST**](DatasourcecontrollerApi.md#AddDataSourceUsingPOST) | **Post** /datasource/_datasource | addDataSource
[**EditDataSourceUsingPUT**](DatasourcecontrollerApi.md#EditDataSourceUsingPUT) | **Put** /datasource/_datasource/put/{id} | editDataSource
[**FindRoleUsingGET**](DatasourcecontrollerApi.md#FindRoleUsingGET) | **Get** /datasource/_datasource/{datasourceId} | findRole
[**ListUsingGET**](DatasourcecontrollerApi.md#ListUsingGET) | **Get** /datasource/_datasource | list


# **AddDataSourceUsingPOST**
> DataSource AddDataSourceUsingPOST($dataSource)

addDataSource


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dataSource** | [**DataSource**](DataSource.md)| dataSource | 

### Return type

[**DataSource**](DataSource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditDataSourceUsingPUT**
> DataSource EditDataSourceUsingPUT($id, $dataSource)

editDataSource


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| id | 
 **dataSource** | [**DataSource**](DataSource.md)| dataSource | 

### Return type

[**DataSource**](DataSource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindRoleUsingGET**
> DataSource FindRoleUsingGET($datasourceId)

findRole


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **datasourceId** | **string**| datasourceId | 

### Return type

[**DataSource**](DataSource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListUsingGET**
> []DataSource ListUsingGET()

list


### Parameters
This endpoint does not need any parameter.

### Return type

[**[]DataSource**](DataSource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

