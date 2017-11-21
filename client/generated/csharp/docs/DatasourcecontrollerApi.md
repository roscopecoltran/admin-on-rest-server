# IO.Swagger.Api.DatasourcecontrollerApi

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddDataSourceUsingPOST**](DatasourcecontrollerApi.md#adddatasourceusingpost) | **POST** /datasource/_datasource | addDataSource
[**EditDataSourceUsingPUT**](DatasourcecontrollerApi.md#editdatasourceusingput) | **PUT** /datasource/_datasource/put/{id} | editDataSource
[**FindRoleUsingGET**](DatasourcecontrollerApi.md#findroleusingget) | **GET** /datasource/_datasource/{datasourceId} | findRole
[**ListUsingGET**](DatasourcecontrollerApi.md#listusingget) | **GET** /datasource/_datasource | list


<a name="adddatasourceusingpost"></a>
# **AddDataSourceUsingPOST**
> DataSource AddDataSourceUsingPOST (DataSource dataSource)

addDataSource

### Example
```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class AddDataSourceUsingPOSTExample
    {
        public void main()
        {
            var apiInstance = new DatasourcecontrollerApi();
            var dataSource = new DataSource(); // DataSource | dataSource

            try
            {
                // addDataSource
                DataSource result = apiInstance.AddDataSourceUsingPOST(dataSource);
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling DatasourcecontrollerApi.AddDataSourceUsingPOST: " + e.Message );
            }
        }
    }
}
```

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

<a name="editdatasourceusingput"></a>
# **EditDataSourceUsingPUT**
> DataSource EditDataSourceUsingPUT (string id, DataSource dataSource)

editDataSource

### Example
```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class EditDataSourceUsingPUTExample
    {
        public void main()
        {
            var apiInstance = new DatasourcecontrollerApi();
            var id = id_example;  // string | id
            var dataSource = new DataSource(); // DataSource | dataSource

            try
            {
                // editDataSource
                DataSource result = apiInstance.EditDataSourceUsingPUT(id, dataSource);
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling DatasourcecontrollerApi.EditDataSourceUsingPUT: " + e.Message );
            }
        }
    }
}
```

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

<a name="findroleusingget"></a>
# **FindRoleUsingGET**
> DataSource FindRoleUsingGET (string datasourceId)

findRole

### Example
```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class FindRoleUsingGETExample
    {
        public void main()
        {
            var apiInstance = new DatasourcecontrollerApi();
            var datasourceId = datasourceId_example;  // string | datasourceId

            try
            {
                // findRole
                DataSource result = apiInstance.FindRoleUsingGET(datasourceId);
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling DatasourcecontrollerApi.FindRoleUsingGET: " + e.Message );
            }
        }
    }
}
```

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

<a name="listusingget"></a>
# **ListUsingGET**
> List<DataSource> ListUsingGET ()

list

### Example
```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class ListUsingGETExample
    {
        public void main()
        {
            var apiInstance = new DatasourcecontrollerApi();

            try
            {
                // list
                List&lt;DataSource&gt; result = apiInstance.ListUsingGET();
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling DatasourcecontrollerApi.ListUsingGET: " + e.Message );
            }
        }
    }
}
```

### Parameters
This endpoint does not need any parameter.

### Return type

[**List<DataSource>**](DataSource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

