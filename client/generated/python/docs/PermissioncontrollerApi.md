# swagger_client.PermissioncontrollerApi

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**add_permission_using_post**](PermissioncontrollerApi.md#add_permission_using_post) | **POST** /permission/_permission | addPermission
[**edit_permission_using_put**](PermissioncontrollerApi.md#edit_permission_using_put) | **PUT** /permission/_permission/{id} | editPermission
[**find_user_using_get**](PermissioncontrollerApi.md#find_user_using_get) | **GET** /permission/_permission/{id} | findUser
[**list_using_get1**](PermissioncontrollerApi.md#list_using_get1) | **GET** /permission/_permission | list


# **add_permission_using_post**
> Permission add_permission_using_post(permission)

addPermission

### Example 
```python
from __future__ import print_function
import time
import swagger_client
from swagger_client.rest import ApiException
from pprint import pprint

# create an instance of the API class
api_instance = swagger_client.PermissioncontrollerApi()
permission = swagger_client.Permission() # Permission | permission

try: 
    # addPermission
    api_response = api_instance.add_permission_using_post(permission)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling PermissioncontrollerApi->add_permission_using_post: %s\n" % e)
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

# **edit_permission_using_put**
> Permission edit_permission_using_put(id, permission)

editPermission

### Example 
```python
from __future__ import print_function
import time
import swagger_client
from swagger_client.rest import ApiException
from pprint import pprint

# create an instance of the API class
api_instance = swagger_client.PermissioncontrollerApi()
id = 'id_example' # str | id
permission = swagger_client.Permission() # Permission | permission

try: 
    # editPermission
    api_response = api_instance.edit_permission_using_put(id, permission)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling PermissioncontrollerApi->edit_permission_using_put: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **str**| id | 
 **permission** | [**Permission**](Permission.md)| permission | 

### Return type

[**Permission**](Permission.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **find_user_using_get**
> Permission find_user_using_get(id)

findUser

### Example 
```python
from __future__ import print_function
import time
import swagger_client
from swagger_client.rest import ApiException
from pprint import pprint

# create an instance of the API class
api_instance = swagger_client.PermissioncontrollerApi()
id = 'id_example' # str | id

try: 
    # findUser
    api_response = api_instance.find_user_using_get(id)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling PermissioncontrollerApi->find_user_using_get: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **str**| id | 

### Return type

[**Permission**](Permission.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_using_get1**
> list[Permission] list_using_get1(role_id)

list

### Example 
```python
from __future__ import print_function
import time
import swagger_client
from swagger_client.rest import ApiException
from pprint import pprint

# create an instance of the API class
api_instance = swagger_client.PermissioncontrollerApi()
role_id = 'role_id_example' # str | roleId

try: 
    # list
    api_response = api_instance.list_using_get1(role_id)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling PermissioncontrollerApi->list_using_get1: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **role_id** | **str**| roleId | 

### Return type

[**list[Permission]**](Permission.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

