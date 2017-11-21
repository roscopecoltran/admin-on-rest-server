# swagger_client.DatasourcecontrollerApi

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**add_data_source_using_post**](DatasourcecontrollerApi.md#add_data_source_using_post) | **POST** /datasource/_datasource | addDataSource
[**edit_data_source_using_put**](DatasourcecontrollerApi.md#edit_data_source_using_put) | **PUT** /datasource/_datasource/put/{id} | editDataSource
[**find_role_using_get**](DatasourcecontrollerApi.md#find_role_using_get) | **GET** /datasource/_datasource/{datasourceId} | findRole
[**list_using_get**](DatasourcecontrollerApi.md#list_using_get) | **GET** /datasource/_datasource | list


# **add_data_source_using_post**
> DataSource add_data_source_using_post(data_source)

addDataSource

### Example 
```python
from __future__ import print_function
import time
import swagger_client
from swagger_client.rest import ApiException
from pprint import pprint

# create an instance of the API class
api_instance = swagger_client.DatasourcecontrollerApi()
data_source = swagger_client.DataSource() # DataSource | dataSource

try: 
    # addDataSource
    api_response = api_instance.add_data_source_using_post(data_source)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling DatasourcecontrollerApi->add_data_source_using_post: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data_source** | [**DataSource**](DataSource.md)| dataSource | 

### Return type

[**DataSource**](DataSource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **edit_data_source_using_put**
> DataSource edit_data_source_using_put(id, data_source)

editDataSource

### Example 
```python
from __future__ import print_function
import time
import swagger_client
from swagger_client.rest import ApiException
from pprint import pprint

# create an instance of the API class
api_instance = swagger_client.DatasourcecontrollerApi()
id = 'id_example' # str | id
data_source = swagger_client.DataSource() # DataSource | dataSource

try: 
    # editDataSource
    api_response = api_instance.edit_data_source_using_put(id, data_source)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling DatasourcecontrollerApi->edit_data_source_using_put: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **str**| id | 
 **data_source** | [**DataSource**](DataSource.md)| dataSource | 

### Return type

[**DataSource**](DataSource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **find_role_using_get**
> DataSource find_role_using_get(datasource_id)

findRole

### Example 
```python
from __future__ import print_function
import time
import swagger_client
from swagger_client.rest import ApiException
from pprint import pprint

# create an instance of the API class
api_instance = swagger_client.DatasourcecontrollerApi()
datasource_id = 'datasource_id_example' # str | datasourceId

try: 
    # findRole
    api_response = api_instance.find_role_using_get(datasource_id)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling DatasourcecontrollerApi->find_role_using_get: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **datasource_id** | **str**| datasourceId | 

### Return type

[**DataSource**](DataSource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_using_get**
> list[DataSource] list_using_get()

list

### Example 
```python
from __future__ import print_function
import time
import swagger_client
from swagger_client.rest import ApiException
from pprint import pprint

# create an instance of the API class
api_instance = swagger_client.DatasourcecontrollerApi()

try: 
    # list
    api_response = api_instance.list_using_get()
    pprint(api_response)
except ApiException as e:
    print("Exception when calling DatasourcecontrollerApi->list_using_get: %s\n" % e)
```

### Parameters
This endpoint does not need any parameter.

### Return type

[**list[DataSource]**](DataSource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

