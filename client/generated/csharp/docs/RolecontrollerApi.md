# IO.Swagger.Api.RolecontrollerApi

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddRoleUsingPOST**](RolecontrollerApi.md#addroleusingpost) | **POST** /role/_roles | addRole
[**EditRoleUsingPUT**](RolecontrollerApi.md#editroleusingput) | **PUT** /role/_roles/put/{id} | editRole
[**FindRoleUsingGET1**](RolecontrollerApi.md#findroleusingget1) | **GET** /role/_roles/{roleId} | findRole
[**ListUsingGET2**](RolecontrollerApi.md#listusingget2) | **GET** /role/_roles | list


<a name="addroleusingpost"></a>
# **AddRoleUsingPOST**
> Role AddRoleUsingPOST (Role role)

addRole

### Example
```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class AddRoleUsingPOSTExample
    {
        public void main()
        {
            var apiInstance = new RolecontrollerApi();
            var role = new Role(); // Role | role

            try
            {
                // addRole
                Role result = apiInstance.AddRoleUsingPOST(role);
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling RolecontrollerApi.AddRoleUsingPOST: " + e.Message );
            }
        }
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **role** | [**Role**](Role.md)| role | 

### Return type

[**Role**](Role.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="editroleusingput"></a>
# **EditRoleUsingPUT**
> Role EditRoleUsingPUT (string id, Role role)

editRole

### Example
```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class EditRoleUsingPUTExample
    {
        public void main()
        {
            var apiInstance = new RolecontrollerApi();
            var id = id_example;  // string | id
            var role = new Role(); // Role | role

            try
            {
                // editRole
                Role result = apiInstance.EditRoleUsingPUT(id, role);
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling RolecontrollerApi.EditRoleUsingPUT: " + e.Message );
            }
        }
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| id | 
 **role** | [**Role**](Role.md)| role | 

### Return type

[**Role**](Role.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="findroleusingget1"></a>
# **FindRoleUsingGET1**
> Role FindRoleUsingGET1 (string roleId)

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
    public class FindRoleUsingGET1Example
    {
        public void main()
        {
            var apiInstance = new RolecontrollerApi();
            var roleId = roleId_example;  // string | roleId

            try
            {
                // findRole
                Role result = apiInstance.FindRoleUsingGET1(roleId);
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling RolecontrollerApi.FindRoleUsingGET1: " + e.Message );
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

[**Role**](Role.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="listusingget2"></a>
# **ListUsingGET2**
> List<Role> ListUsingGET2 ()

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
    public class ListUsingGET2Example
    {
        public void main()
        {
            var apiInstance = new RolecontrollerApi();

            try
            {
                // list
                List&lt;Role&gt; result = apiInstance.ListUsingGET2();
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling RolecontrollerApi.ListUsingGET2: " + e.Message );
            }
        }
    }
}
```

### Parameters
This endpoint does not need any parameter.

### Return type

[**List<Role>**](Role.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

