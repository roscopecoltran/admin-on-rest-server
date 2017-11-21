# swagger_client.SchemacontrollerApi

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**add_entity_using_post**](SchemacontrollerApi.md#add_entity_using_post) | **POST** /schemas/_entitys | addEntity
[**add_field_using_post**](SchemacontrollerApi.md#add_field_using_post) | **POST** /schemas/_fields | addField
[**edit_entity_using_put**](SchemacontrollerApi.md#edit_entity_using_put) | **PUT** /schemas/_entitys/put/{id} | editEntity
[**edit_field_using_put**](SchemacontrollerApi.md#edit_field_using_put) | **PUT** /schemas/_fields/put/{id} | editField
[**find_all_fields_using_get**](SchemacontrollerApi.md#find_all_fields_using_get) | **GET** /schemas/_fields | findAllFields
[**find_entity_fields_using_get**](SchemacontrollerApi.md#find_entity_fields_using_get) | **GET** /schemas/fields | findEntityFields
[**find_one_field_using_get**](SchemacontrollerApi.md#find_one_field_using_get) | **GET** /schemas/_fields/{fid} | findOneField
[**find_one_using_get1**](SchemacontrollerApi.md#find_one_using_get1) | **GET** /schemas/_entitys/{eid} | findOne
[**get_schemas_using_get**](SchemacontrollerApi.md#get_schemas_using_get) | **GET** /schemas/_entitys | getSchemas
[**reset_current_ds_using_put**](SchemacontrollerApi.md#reset_current_ds_using_put) | **PUT** /schemas/resetCurrentDs/{dataSourceId} | resetCurrentDs
[**sync_schemas_using_put**](SchemacontrollerApi.md#sync_schemas_using_put) | **PUT** /schemas/sync/{dataSourceId} | syncSchemas


# **add_entity_using_post**
> Entity add_entity_using_post(entity)

addEntity

### Example 
```python
from __future__ import print_function
import time
import swagger_client
from swagger_client.rest import ApiException
from pprint import pprint

# create an instance of the API class
api_instance = swagger_client.SchemacontrollerApi()
entity = swagger_client.Entity() # Entity | entity

try: 
    # addEntity
    api_response = api_instance.add_entity_using_post(entity)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling SchemacontrollerApi->add_entity_using_post: %s\n" % e)
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

# **add_field_using_post**
> Field add_field_using_post(field)

addField

### Example 
```python
from __future__ import print_function
import time
import swagger_client
from swagger_client.rest import ApiException
from pprint import pprint

# create an instance of the API class
api_instance = swagger_client.SchemacontrollerApi()
field = swagger_client.Field() # Field | field

try: 
    # addField
    api_response = api_instance.add_field_using_post(field)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling SchemacontrollerApi->add_field_using_post: %s\n" % e)
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

# **edit_entity_using_put**
> Entity edit_entity_using_put(id, entity)

editEntity

### Example 
```python
from __future__ import print_function
import time
import swagger_client
from swagger_client.rest import ApiException
from pprint import pprint

# create an instance of the API class
api_instance = swagger_client.SchemacontrollerApi()
id = 'id_example' # str | id
entity = swagger_client.Entity() # Entity | entity

try: 
    # editEntity
    api_response = api_instance.edit_entity_using_put(id, entity)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling SchemacontrollerApi->edit_entity_using_put: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **str**| id | 
 **entity** | [**Entity**](Entity.md)| entity | 

### Return type

[**Entity**](Entity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **edit_field_using_put**
> Field edit_field_using_put(id, field)

editField

### Example 
```python
from __future__ import print_function
import time
import swagger_client
from swagger_client.rest import ApiException
from pprint import pprint

# create an instance of the API class
api_instance = swagger_client.SchemacontrollerApi()
id = 'id_example' # str | id
field = swagger_client.Field() # Field | field

try: 
    # editField
    api_response = api_instance.edit_field_using_put(id, field)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling SchemacontrollerApi->edit_field_using_put: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **str**| id | 
 **field** | [**Field**](Field.md)| field | 

### Return type

[**Field**](Field.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **find_all_fields_using_get**
> list[Field] find_all_fields_using_get(eid)

findAllFields

### Example 
```python
from __future__ import print_function
import time
import swagger_client
from swagger_client.rest import ApiException
from pprint import pprint

# create an instance of the API class
api_instance = swagger_client.SchemacontrollerApi()
eid = 'eid_example' # str | eid

try: 
    # findAllFields
    api_response = api_instance.find_all_fields_using_get(eid)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling SchemacontrollerApi->find_all_fields_using_get: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eid** | **str**| eid | 

### Return type

[**list[Field]**](Field.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **find_entity_fields_using_get**
> list[Field] find_entity_fields_using_get(entity_name)

findEntityFields

### Example 
```python
from __future__ import print_function
import time
import swagger_client
from swagger_client.rest import ApiException
from pprint import pprint

# create an instance of the API class
api_instance = swagger_client.SchemacontrollerApi()
entity_name = 'entity_name_example' # str | entityName

try: 
    # findEntityFields
    api_response = api_instance.find_entity_fields_using_get(entity_name)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling SchemacontrollerApi->find_entity_fields_using_get: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entity_name** | **str**| entityName | 

### Return type

[**list[Field]**](Field.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **find_one_field_using_get**
> Field find_one_field_using_get(fid)

findOneField

### Example 
```python
from __future__ import print_function
import time
import swagger_client
from swagger_client.rest import ApiException
from pprint import pprint

# create an instance of the API class
api_instance = swagger_client.SchemacontrollerApi()
fid = 'fid_example' # str | fid

try: 
    # findOneField
    api_response = api_instance.find_one_field_using_get(fid)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling SchemacontrollerApi->find_one_field_using_get: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fid** | **str**| fid | 

### Return type

[**Field**](Field.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **find_one_using_get1**
> Entity find_one_using_get1(eid)

findOne

### Example 
```python
from __future__ import print_function
import time
import swagger_client
from swagger_client.rest import ApiException
from pprint import pprint

# create an instance of the API class
api_instance = swagger_client.SchemacontrollerApi()
eid = 'eid_example' # str | eid

try: 
    # findOne
    api_response = api_instance.find_one_using_get1(eid)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling SchemacontrollerApi->find_one_using_get1: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eid** | **str**| eid | 

### Return type

[**Entity**](Entity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_schemas_using_get**
> list[Entity] get_schemas_using_get()

getSchemas

### Example 
```python
from __future__ import print_function
import time
import swagger_client
from swagger_client.rest import ApiException
from pprint import pprint

# create an instance of the API class
api_instance = swagger_client.SchemacontrollerApi()

try: 
    # getSchemas
    api_response = api_instance.get_schemas_using_get()
    pprint(api_response)
except ApiException as e:
    print("Exception when calling SchemacontrollerApi->get_schemas_using_get: %s\n" % e)
```

### Parameters
This endpoint does not need any parameter.

### Return type

[**list[Entity]**](Entity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **reset_current_ds_using_put**
> ResponseEntity reset_current_ds_using_put(data_source_id)

resetCurrentDs

### Example 
```python
from __future__ import print_function
import time
import swagger_client
from swagger_client.rest import ApiException
from pprint import pprint

# create an instance of the API class
api_instance = swagger_client.SchemacontrollerApi()
data_source_id = 'data_source_id_example' # str | dataSourceId

try: 
    # resetCurrentDs
    api_response = api_instance.reset_current_ds_using_put(data_source_id)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling SchemacontrollerApi->reset_current_ds_using_put: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data_source_id** | **str**| dataSourceId | 

### Return type

[**ResponseEntity**](ResponseEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **sync_schemas_using_put**
> ResponseEntity sync_schemas_using_put(data_source_id)

syncSchemas

### Example 
```python
from __future__ import print_function
import time
import swagger_client
from swagger_client.rest import ApiException
from pprint import pprint

# create an instance of the API class
api_instance = swagger_client.SchemacontrollerApi()
data_source_id = 'data_source_id_example' # str | dataSourceId

try: 
    # syncSchemas
    api_response = api_instance.sync_schemas_using_put(data_source_id)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling SchemacontrollerApi->sync_schemas_using_put: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data_source_id** | **str**| dataSourceId | 

### Return type

[**ResponseEntity**](ResponseEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

