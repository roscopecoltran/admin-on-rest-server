# IO.Swagger.Api.PermissioncontrollerApi

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPermissionUsingPOST**](PermissioncontrollerApi.md#addpermissionusingpost) | **POST** /permission/_permission | addPermission
[**EditPermissionUsingPUT**](PermissioncontrollerApi.md#editpermissionusingput) | **PUT** /permission/_permission/{id} | editPermission
[**FindUserUsingGET**](PermissioncontrollerApi.md#finduserusingget) | **GET** /permission/_permission/{id} | findUser
[**ListUsingGET1**](PermissioncontrollerApi.md#listusingget1) | **GET** /permission/_permission | list


<a name="addpermissionusingpost"></a>
# **AddPermissionUsingPOST**
> Permission AddPermissionUsingPOST (Permission permission)

addPermission

### Example
```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class AddPermissionUsingPOSTExample
    {
        public void main()
        {
            var apiInstance = new PermissioncontrollerApi();
            var permission = new Permission(); // Permission | permission

            try
            {
                // addPermission
                Permission result = apiInstance.AddPermissionUsingPOST(permission);
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling PermissioncontrollerApi.AddPermissionUsingPOST: " + e.Message );
            }
        }
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **permission** | [**Permission**](Permission.md)| permission | 

### Return type

[**Permission**](Permission.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="editpermissionusingput"></a>
# **EditPermissionUsingPUT**
> Permission EditPermissionUsingPUT (string id, Permission permission)

editPermission

### Example
```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class EditPermissionUsingPUTExample
    {
        public void main()
        {
            var apiInstance = new PermissioncontrollerApi();
            var id = id_example;  // string | id
            var permission = new Permission(); // Permission | permission

            try
            {
                // editPermission
                Permission result = apiInstance.EditPermissionUsingPUT(id, permission);
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling PermissioncontrollerApi.EditPermissionUsingPUT: " + e.Message );
            }
        }
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| id | 
 **permission** | [**Permission**](Permission.md)| permission | 

### Return type

[**Permission**](Permission.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="finduserusingget"></a>
# **FindUserUsingGET**
> Permission FindUserUsingGET (string id)

findUser

### Example
```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class FindUserUsingGETExample
    {
        public void main()
        {
            var apiInstance = new PermissioncontrollerApi();
            var id = id_example;  // string | id

            try
            {
                // findUser
                Permission result = apiInstance.FindUserUsingGET(id);
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling PermissioncontrollerApi.FindUserUsingGET: " + e.Message );
            }
        }
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| id | 

### Return type

[**Permission**](Permission.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="listusingget1"></a>
# **ListUsingGET1**
> List<Permission> ListUsingGET1 (string roleId)

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
    public class ListUsingGET1Example
    {
        public void main()
        {
            var apiInstance = new PermissioncontrollerApi();
            var roleId = roleId_example;  // string | roleId

            try
            {
                // list
                List&lt;Permission&gt; result = apiInstance.ListUsingGET1(roleId);
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling PermissioncontrollerApi.ListUsingGET1: " + e.Message );
            }
        }
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **roleId** | **string**| roleId | 

### Return type

[**List<Permission>**](Permission.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

