# swagger_client.RolecontrollerApi

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**add_role_using_post**](RolecontrollerApi.md#add_role_using_post) | **POST** /role/_roles | addRole
[**edit_role_using_put**](RolecontrollerApi.md#edit_role_using_put) | **PUT** /role/_roles/put/{id} | editRole
[**find_role_using_get1**](RolecontrollerApi.md#find_role_using_get1) | **GET** /role/_roles/{roleId} | findRole
[**list_using_get2**](RolecontrollerApi.md#list_using_get2) | **GET** /role/_roles | list


# **add_role_using_post**
> Role add_role_using_post(role)

addRole

### Example 
```python
from __future__ import print_function
import time
import swagger_client
from swagger_client.rest import ApiException
from pprint import pprint

# create an instance of the API class
api_instance = swagger_client.RolecontrollerApi()
role = swagger_client.Role() # Role | role

try: 
    # addRole
    api_response = api_instance.add_role_using_post(role)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling RolecontrollerApi->add_role_using_post: %s\n" % e)
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

# **edit_role_using_put**
> Role edit_role_using_put(id, role)

editRole

### Example 
```python
from __future__ import print_function
import time
import swagger_client
from swagger_client.rest import ApiException
from pprint import pprint

# create an instance of the API class
api_instance = swagger_client.RolecontrollerApi()
id = 'id_example' # str | id
role = swagger_client.Role() # Role | role

try: 
    # editRole
    api_response = api_instance.edit_role_using_put(id, role)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling RolecontrollerApi->edit_role_using_put: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **str**| id | 
 **role** | [**Role**](Role.md)| role | 

### Return type

[**Role**](Role.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **find_role_using_get1**
> Role find_role_using_get1(role_id)

findRole

### Example 
```python
from __future__ import print_function
import time
import swagger_client
from swagger_client.rest import ApiException
from pprint import pprint

# create an instance of the API class
api_instance = swagger_client.RolecontrollerApi()
role_id = 'role_id_example' # str | roleId

try: 
    # findRole
    api_response = api_instance.find_role_using_get1(role_id)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling RolecontrollerApi->find_role_using_get1: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **role_id** | **str**| roleId | 

### Return type

[**Role**](Role.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_using_get2**
> list[Role] list_using_get2()

list

### Example 
```python
from __future__ import print_function
import time
import swagger_client
from swagger_client.rest import ApiException
from pprint import pprint

# create an instance of the API class
api_instance = swagger_client.RolecontrollerApi()

try: 
    # list
    api_response = api_instance.list_using_get2()
    pprint(api_response)
except ApiException as e:
    print("Exception when calling RolecontrollerApi->list_using_get2: %s\n" % e)
```

### Parameters
This endpoint does not need any parameter.

### Return type

[**list[Role]**](Role.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

