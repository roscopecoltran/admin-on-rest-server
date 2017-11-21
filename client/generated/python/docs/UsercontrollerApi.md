# swagger_client.UsercontrollerApi

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**add_user_using_post**](UsercontrollerApi.md#add_user_using_post) | **POST** /user/_users | addUser
[**edit_field_using_put1**](UsercontrollerApi.md#edit_field_using_put1) | **PUT** /user/_users/put/{id} | editField
[**find_user_using_get1**](UsercontrollerApi.md#find_user_using_get1) | **GET** /user/_users/{userId} | findUser
[**get_authenticated_user_using_get**](UsercontrollerApi.md#get_authenticated_user_using_get) | **GET** /user/me | getAuthenticatedUser
[**list_using_get3**](UsercontrollerApi.md#list_using_get3) | **GET** /user/_users | list


# **add_user_using_post**
> User add_user_using_post(user)

addUser

### Example 
```python
from __future__ import print_function
import time
import swagger_client
from swagger_client.rest import ApiException
from pprint import pprint

# create an instance of the API class
api_instance = swagger_client.UsercontrollerApi()
user = swagger_client.User() # User | user

try: 
    # addUser
    api_response = api_instance.add_user_using_post(user)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling UsercontrollerApi->add_user_using_post: %s\n" % e)
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

# **edit_field_using_put1**
> User edit_field_using_put1(id, user)

editField

### Example 
```python
from __future__ import print_function
import time
import swagger_client
from swagger_client.rest import ApiException
from pprint import pprint

# create an instance of the API class
api_instance = swagger_client.UsercontrollerApi()
id = 'id_example' # str | id
user = swagger_client.User() # User | user

try: 
    # editField
    api_response = api_instance.edit_field_using_put1(id, user)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling UsercontrollerApi->edit_field_using_put1: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **str**| id | 
 **user** | [**User**](User.md)| user | 

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **find_user_using_get1**
> User find_user_using_get1(user_id)

findUser

### Example 
```python
from __future__ import print_function
import time
import swagger_client
from swagger_client.rest import ApiException
from pprint import pprint

# create an instance of the API class
api_instance = swagger_client.UsercontrollerApi()
user_id = 'user_id_example' # str | userId

try: 
    # findUser
    api_response = api_instance.find_user_using_get1(user_id)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling UsercontrollerApi->find_user_using_get1: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user_id** | **str**| userId | 

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_authenticated_user_using_get**
> JwtUser get_authenticated_user_using_get()

getAuthenticatedUser

### Example 
```python
from __future__ import print_function
import time
import swagger_client
from swagger_client.rest import ApiException
from pprint import pprint

# create an instance of the API class
api_instance = swagger_client.UsercontrollerApi()

try: 
    # getAuthenticatedUser
    api_response = api_instance.get_authenticated_user_using_get()
    pprint(api_response)
except ApiException as e:
    print("Exception when calling UsercontrollerApi->get_authenticated_user_using_get: %s\n" % e)
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

# **list_using_get3**
> list[User] list_using_get3()

list

### Example 
```python
from __future__ import print_function
import time
import swagger_client
from swagger_client.rest import ApiException
from pprint import pprint

# create an instance of the API class
api_instance = swagger_client.UsercontrollerApi()

try: 
    # list
    api_response = api_instance.list_using_get3()
    pprint(api_response)
except ApiException as e:
    print("Exception when calling UsercontrollerApi->list_using_get3: %s\n" % e)
```

### Parameters
This endpoint does not need any parameter.

### Return type

[**list[User]**](User.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

