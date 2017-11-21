# IO.Swagger.Api.UsercontrollerApi

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddUserUsingPOST**](UsercontrollerApi.md#adduserusingpost) | **POST** /user/_users | addUser
[**EditFieldUsingPUT1**](UsercontrollerApi.md#editfieldusingput1) | **PUT** /user/_users/put/{id} | editField
[**FindUserUsingGET1**](UsercontrollerApi.md#finduserusingget1) | **GET** /user/_users/{userId} | findUser
[**GetAuthenticatedUserUsingGET**](UsercontrollerApi.md#getauthenticateduserusingget) | **GET** /user/me | getAuthenticatedUser
[**ListUsingGET3**](UsercontrollerApi.md#listusingget3) | **GET** /user/_users | list


<a name="adduserusingpost"></a>
# **AddUserUsingPOST**
> User AddUserUsingPOST (User user)

addUser

### Example
```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class AddUserUsingPOSTExample
    {
        public void main()
        {
            var apiInstance = new UsercontrollerApi();
            var user = new User(); // User | user

            try
            {
                // addUser
                User result = apiInstance.AddUserUsingPOST(user);
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling UsercontrollerApi.AddUserUsingPOST: " + e.Message );
            }
        }
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user** | [**User**](User.md)| user | 

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="editfieldusingput1"></a>
# **EditFieldUsingPUT1**
> User EditFieldUsingPUT1 (string id, User user)

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
    public class EditFieldUsingPUT1Example
    {
        public void main()
        {
            var apiInstance = new UsercontrollerApi();
            var id = id_example;  // string | id
            var user = new User(); // User | user

            try
            {
                // editField
                User result = apiInstance.EditFieldUsingPUT1(id, user);
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling UsercontrollerApi.EditFieldUsingPUT1: " + e.Message );
            }
        }
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| id | 
 **user** | [**User**](User.md)| user | 

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="finduserusingget1"></a>
# **FindUserUsingGET1**
> User FindUserUsingGET1 (string userId)

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
    public class FindUserUsingGET1Example
    {
        public void main()
        {
            var apiInstance = new UsercontrollerApi();
            var userId = userId_example;  // string | userId

            try
            {
                // findUser
                User result = apiInstance.FindUserUsingGET1(userId);
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling UsercontrollerApi.FindUserUsingGET1: " + e.Message );
            }
        }
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string**| userId | 

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="getauthenticateduserusingget"></a>
# **GetAuthenticatedUserUsingGET**
> JwtUser GetAuthenticatedUserUsingGET ()

getAuthenticatedUser

### Example
```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class GetAuthenticatedUserUsingGETExample
    {
        public void main()
        {
            var apiInstance = new UsercontrollerApi();

            try
            {
                // getAuthenticatedUser
                JwtUser result = apiInstance.GetAuthenticatedUserUsingGET();
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling UsercontrollerApi.GetAuthenticatedUserUsingGET: " + e.Message );
            }
        }
    }
}
```

### Parameters
This endpoint does not need any parameter.

### Return type

[**JwtUser**](JwtUser.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="listusingget3"></a>
# **ListUsingGET3**
> List<User> ListUsingGET3 ()

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
    public class ListUsingGET3Example
    {
        public void main()
        {
            var apiInstance = new UsercontrollerApi();

            try
            {
                // list
                List&lt;User&gt; result = apiInstance.ListUsingGET3();
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling UsercontrollerApi.ListUsingGET3: " + e.Message );
            }
        }
    }
}
```

### Parameters
This endpoint does not need any parameter.

### Return type

[**List<User>**](User.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

