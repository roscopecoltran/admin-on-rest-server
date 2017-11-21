# IO.Swagger.Api.SchemacontrollerApi

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddEntityUsingPOST**](SchemacontrollerApi.md#addentityusingpost) | **POST** /schemas/_entitys | addEntity
[**AddFieldUsingPOST**](SchemacontrollerApi.md#addfieldusingpost) | **POST** /schemas/_fields | addField
[**EditEntityUsingPUT**](SchemacontrollerApi.md#editentityusingput) | **PUT** /schemas/_entitys/put/{id} | editEntity
[**EditFieldUsingPUT**](SchemacontrollerApi.md#editfieldusingput) | **PUT** /schemas/_fields/put/{id} | editField
[**FindAllFieldsUsingGET**](SchemacontrollerApi.md#findallfieldsusingget) | **GET** /schemas/_fields | findAllFields
[**FindEntityFieldsUsingGET**](SchemacontrollerApi.md#findentityfieldsusingget) | **GET** /schemas/fields | findEntityFields
[**FindOneFieldUsingGET**](SchemacontrollerApi.md#findonefieldusingget) | **GET** /schemas/_fields/{fid} | findOneField
[**FindOneUsingGET1**](SchemacontrollerApi.md#findoneusingget1) | **GET** /schemas/_entitys/{eid} | findOne
[**GetSchemasUsingGET**](SchemacontrollerApi.md#getschemasusingget) | **GET** /schemas/_entitys | getSchemas
[**ResetCurrentDsUsingPUT**](SchemacontrollerApi.md#resetcurrentdsusingput) | **PUT** /schemas/resetCurrentDs/{dataSourceId} | resetCurrentDs
[**SyncSchemasUsingPUT**](SchemacontrollerApi.md#syncschemasusingput) | **PUT** /schemas/sync/{dataSourceId} | syncSchemas


<a name="addentityusingpost"></a>
# **AddEntityUsingPOST**
> Entity AddEntityUsingPOST (Entity entity)

addEntity

### Example
```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class AddEntityUsingPOSTExample
    {
        public void main()
        {
            var apiInstance = new SchemacontrollerApi();
            var entity = new Entity(); // Entity | entity

            try
            {
                // addEntity
                Entity result = apiInstance.AddEntityUsingPOST(entity);
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling SchemacontrollerApi.AddEntityUsingPOST: " + e.Message );
            }
        }
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entity** | [**Entity**](Entity.md)| entity | 

### Return type

[**Entity**](Entity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="addfieldusingpost"></a>
# **AddFieldUsingPOST**
> Field AddFieldUsingPOST (Field field)

addField

### Example
```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class AddFieldUsingPOSTExample
    {
        public void main()
        {
            var apiInstance = new SchemacontrollerApi();
            var field = new Field(); // Field | field

            try
            {
                // addField
                Field result = apiInstance.AddFieldUsingPOST(field);
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling SchemacontrollerApi.AddFieldUsingPOST: " + e.Message );
            }
        }
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **field** | [**Field**](Field.md)| field | 

### Return type

[**Field**](Field.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="editentityusingput"></a>
# **EditEntityUsingPUT**
> Entity EditEntityUsingPUT (string id, Entity entity)

editEntity

### Example
```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class EditEntityUsingPUTExample
    {
        public void main()
        {
            var apiInstance = new SchemacontrollerApi();
            var id = id_example;  // string | id
            var entity = new Entity(); // Entity | entity

            try
            {
                // editEntity
                Entity result = apiInstance.EditEntityUsingPUT(id, entity);
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling SchemacontrollerApi.EditEntityUsingPUT: " + e.Message );
            }
        }
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| id | 
 **entity** | [**Entity**](Entity.md)| entity | 

### Return type

[**Entity**](Entity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="editfieldusingput"></a>
# **EditFieldUsingPUT**
> Field EditFieldUsingPUT (string id, Field field)

editField

### Example
```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class EditFieldUsingPUTExample
    {
        public void main()
        {
            var apiInstance = new SchemacontrollerApi();
            var id = id_example;  // string | id
            var field = new Field(); // Field | field

            try
            {
                // editField
                Field result = apiInstance.EditFieldUsingPUT(id, field);
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling SchemacontrollerApi.EditFieldUsingPUT: " + e.Message );
            }
        }
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| id | 
 **field** | [**Field**](Field.md)| field | 

### Return type

[**Field**](Field.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="findallfieldsusingget"></a>
# **FindAllFieldsUsingGET**
> List<Field> FindAllFieldsUsingGET (string eid)

findAllFields

### Example
```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class FindAllFieldsUsingGETExample
    {
        public void main()
        {
            var apiInstance = new SchemacontrollerApi();
            var eid = eid_example;  // string | eid

            try
            {
                // findAllFields
                List&lt;Field&gt; result = apiInstance.FindAllFieldsUsingGET(eid);
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling SchemacontrollerApi.FindAllFieldsUsingGET: " + e.Message );
            }
        }
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eid** | **string**| eid | 

### Return type

[**List<Field>**](Field.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="findentityfieldsusingget"></a>
# **FindEntityFieldsUsingGET**
> List<Field> FindEntityFieldsUsingGET (string entityName)

findEntityFields

### Example
```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class FindEntityFieldsUsingGETExample
    {
        public void main()
        {
            var apiInstance = new SchemacontrollerApi();
            var entityName = entityName_example;  // string | entityName

            try
            {
                // findEntityFields
                List&lt;Field&gt; result = apiInstance.FindEntityFieldsUsingGET(entityName);
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling SchemacontrollerApi.FindEntityFieldsUsingGET: " + e.Message );
            }
        }
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entityName** | **string**| entityName | 

### Return type

[**List<Field>**](Field.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="findonefieldusingget"></a>
# **FindOneFieldUsingGET**
> Field FindOneFieldUsingGET (string fid)

findOneField

### Example
```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class FindOneFieldUsingGETExample
    {
        public void main()
        {
            var apiInstance = new SchemacontrollerApi();
            var fid = fid_example;  // string | fid

            try
            {
                // findOneField
                Field result = apiInstance.FindOneFieldUsingGET(fid);
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling SchemacontrollerApi.FindOneFieldUsingGET: " + e.Message );
            }
        }
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fid** | **string**| fid | 

### Return type

[**Field**](Field.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="findoneusingget1"></a>
# **FindOneUsingGET1**
> Entity FindOneUsingGET1 (string eid)

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
    public class FindOneUsingGET1Example
    {
        public void main()
        {
            var apiInstance = new SchemacontrollerApi();
            var eid = eid_example;  // string | eid

            try
            {
                // findOne
                Entity result = apiInstance.FindOneUsingGET1(eid);
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling SchemacontrollerApi.FindOneUsingGET1: " + e.Message );
            }
        }
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eid** | **string**| eid | 

### Return type

[**Entity**](Entity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="getschemasusingget"></a>
# **GetSchemasUsingGET**
> List<Entity> GetSchemasUsingGET ()

getSchemas

### Example
```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class GetSchemasUsingGETExample
    {
        public void main()
        {
            var apiInstance = new SchemacontrollerApi();

            try
            {
                // getSchemas
                List&lt;Entity&gt; result = apiInstance.GetSchemasUsingGET();
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling SchemacontrollerApi.GetSchemasUsingGET: " + e.Message );
            }
        }
    }
}
```

### Parameters
This endpoint does not need any parameter.

### Return type

[**List<Entity>**](Entity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="resetcurrentdsusingput"></a>
# **ResetCurrentDsUsingPUT**
> ResponseEntity ResetCurrentDsUsingPUT (string dataSourceId)

resetCurrentDs

### Example
```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class ResetCurrentDsUsingPUTExample
    {
        public void main()
        {
            var apiInstance = new SchemacontrollerApi();
            var dataSourceId = dataSourceId_example;  // string | dataSourceId

            try
            {
                // resetCurrentDs
                ResponseEntity result = apiInstance.ResetCurrentDsUsingPUT(dataSourceId);
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling SchemacontrollerApi.ResetCurrentDsUsingPUT: " + e.Message );
            }
        }
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dataSourceId** | **string**| dataSourceId | 

### Return type

[**ResponseEntity**](ResponseEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="syncschemasusingput"></a>
# **SyncSchemasUsingPUT**
> ResponseEntity SyncSchemasUsingPUT (string dataSourceId)

syncSchemas

### Example
```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class SyncSchemasUsingPUTExample
    {
        public void main()
        {
            var apiInstance = new SchemacontrollerApi();
            var dataSourceId = dataSourceId_example;  // string | dataSourceId

            try
            {
                // syncSchemas
                ResponseEntity result = apiInstance.SyncSchemasUsingPUT(dataSourceId);
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling SchemacontrollerApi.SyncSchemasUsingPUT: " + e.Message );
            }
        }
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dataSourceId** | **string**| dataSourceId | 

### Return type

[**ResponseEntity**](ResponseEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

