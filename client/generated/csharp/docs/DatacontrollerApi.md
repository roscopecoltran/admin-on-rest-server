# IO.Swagger.Api.DatacontrollerApi

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DataMutationUsingDELETE**](DatacontrollerApi.md#datamutationusingdelete) | **DELETE** /api/{entity}/{id} | dataMutation
[**DataMutationUsingPOST**](DatacontrollerApi.md#datamutationusingpost) | **POST** /api/{entity} | dataMutation
[**DataMutationUsingPUT**](DatacontrollerApi.md#datamutationusingput) | **PUT** /api/{entity}/{id} | dataMutation
[**DataQueryUsingGET**](DatacontrollerApi.md#dataqueryusingget) | **GET** /api/{entity} | dataQuery
[**FindOneUsingGET**](DatacontrollerApi.md#findoneusingget) | **GET** /api/{entity}/{id} | findOne


<a name="datamutationusingdelete"></a>
# **DataMutationUsingDELETE**
> Dictionary<string, Object> DataMutationUsingDELETE (string entity, string id)

dataMutation

### Example
```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class DataMutationUsingDELETEExample
    {
        public void main()
        {
            var apiInstance = new DatacontrollerApi();
            var entity = entity_example;  // string | entity
            var id = id_example;  // string | id

            try
            {
                // dataMutation
                Dictionary&lt;string, Object&gt; result = apiInstance.DataMutationUsingDELETE(entity, id);
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling DatacontrollerApi.DataMutationUsingDELETE: " + e.Message );
            }
        }
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entity** | **string**| entity | 
 **id** | **string**| id | 

### Return type

**Dictionary<string, Object>**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="datamutationusingpost"></a>
# **DataMutationUsingPOST**
> Dictionary<string, Object> DataMutationUsingPOST (string entity, Object allRequestParams)

dataMutation

### Example
```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class DataMutationUsingPOSTExample
    {
        public void main()
        {
            var apiInstance = new DatacontrollerApi();
            var entity = entity_example;  // string | entity
            var allRequestParams = ;  // Object | allRequestParams

            try
            {
                // dataMutation
                Dictionary&lt;string, Object&gt; result = apiInstance.DataMutationUsingPOST(entity, allRequestParams);
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling DatacontrollerApi.DataMutationUsingPOST: " + e.Message );
            }
        }
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entity** | **string**| entity | 
 **allRequestParams** | **Object**| allRequestParams | 

### Return type

**Dictionary<string, Object>**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="datamutationusingput"></a>
# **DataMutationUsingPUT**
> Dictionary<string, Object> DataMutationUsingPUT (string entity, string id, Object allRequestParams)

dataMutation

### Example
```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class DataMutationUsingPUTExample
    {
        public void main()
        {
            var apiInstance = new DatacontrollerApi();
            var entity = entity_example;  // string | entity
            var id = id_example;  // string | id
            var allRequestParams = ;  // Object | allRequestParams

            try
            {
                // dataMutation
                Dictionary&lt;string, Object&gt; result = apiInstance.DataMutationUsingPUT(entity, id, allRequestParams);
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling DatacontrollerApi.DataMutationUsingPUT: " + e.Message );
            }
        }
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entity** | **string**| entity | 
 **id** | **string**| id | 
 **allRequestParams** | **Object**| allRequestParams | 

### Return type

**Dictionary<string, Object>**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="dataqueryusingget"></a>
# **DataQueryUsingGET**
> Dictionary<string, Object> DataQueryUsingGET (string entity)

dataQuery

### Example
```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class DataQueryUsingGETExample
    {
        public void main()
        {
            var apiInstance = new DatacontrollerApi();
            var entity = entity_example;  // string | entity

            try
            {
                // dataQuery
                Dictionary&lt;string, Object&gt; result = apiInstance.DataQueryUsingGET(entity);
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling DatacontrollerApi.DataQueryUsingGET: " + e.Message );
            }
        }
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entity** | **string**| entity | 

### Return type

**Dictionary<string, Object>**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="findoneusingget"></a>
# **FindOneUsingGET**
> Dictionary<string, Object> FindOneUsingGET (string entity, string id)

findOne

### Example
```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class FindOneUsingGETExample
    {
        public void main()
        {
            var apiInstance = new DatacontrollerApi();
            var entity = entity_example;  // string | entity
            var id = id_example;  // string | id

            try
            {
                // findOne
                Dictionary&lt;string, Object&gt; result = apiInstance.FindOneUsingGET(entity, id);
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling DatacontrollerApi.FindOneUsingGET: " + e.Message );
            }
        }
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entity** | **string**| entity | 
 **id** | **string**| id | 

### Return type

**Dictionary<string, Object>**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

