# IO.Swagger - the C# library for the DataHive RESTful APIs

No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)

This C# SDK is automatically generated by the [Swagger Codegen](https://github.com/swagger-api/swagger-codegen) project:

- API version: 1.0
- SDK version: 1.0.0
- Build package: io.swagger.codegen.languages.CSharpClientCodegen

<a name="frameworks-supported"></a>
## Frameworks supported
- .NET 4.0 or later
- Windows Phone 7.1 (Mango)

<a name="dependencies"></a>
## Dependencies
- [RestSharp](https://www.nuget.org/packages/RestSharp) - 105.1.0 or later
- [Json.NET](https://www.nuget.org/packages/Newtonsoft.Json/) - 7.0.0 or later

The DLLs included in the package may not be the latest version. We recommend using [NuGet] (https://docs.nuget.org/consume/installing-nuget) to obtain the latest version of the packages:
```
Install-Package RestSharp
Install-Package Newtonsoft.Json
```

NOTE: RestSharp versions greater than 105.1.0 have a bug which causes file uploads to fail. See [RestSharp#742](https://github.com/restsharp/RestSharp/issues/742)

<a name="installation"></a>
## Installation
Run the following command to generate the DLL
- [Mac/Linux] `/bin/sh build.sh`
- [Windows] `build.bat`

Then include the DLL (under the `bin` folder) in the C# project, and use the namespaces:
```csharp
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;
```
<a name="packaging"></a>
## Packaging

A `.nuspec` is included with the project. You can follow the Nuget quickstart to [create](https://docs.microsoft.com/en-us/nuget/quickstart/create-and-publish-a-package#create-the-package) and [publish](https://docs.microsoft.com/en-us/nuget/quickstart/create-and-publish-a-package#publish-the-package) packages.

This `.nuspec` uses placeholders from the `.csproj`, so build the `.csproj` directly:

```
nuget pack -Build -OutputDirectory out IO.Swagger.csproj
```

Then, publish to a [local feed](https://docs.microsoft.com/en-us/nuget/hosting-packages/local-feeds) or [other host](https://docs.microsoft.com/en-us/nuget/hosting-packages/overview) and consume the new package via Nuget as usual.

<a name="getting-started"></a>
## Getting Started

```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class Example
    {
        public void main()
        {

            var apiInstance = new ApplycontrollerApi();
            var apply = new Apply(); // Apply | apply

            try
            {
                // apply
                ResponseEntity result = apiInstance.ApplyUsingPOST(apply);
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling ApplycontrollerApi.ApplyUsingPOST: " + e.Message );
            }

        }
    }
}
```

<a name="documentation-for-api-endpoints"></a>
## Documentation for API Endpoints

All URIs are relative to *https://localhost:8080*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ApplycontrollerApi* | [**ApplyUsingPOST**](docs/ApplycontrollerApi.md#applyusingpost) | **POST** /apply | apply
*AuthenticationrestcontrollerApi* | [**CreateAuthenticationTokenUsingPOST**](docs/AuthenticationrestcontrollerApi.md#createauthenticationtokenusingpost) | **POST** /auth | createAuthenticationToken
*AuthenticationrestcontrollerApi* | [**RefreshAndGetAuthenticationTokenUsingGET**](docs/AuthenticationrestcontrollerApi.md#refreshandgetauthenticationtokenusingget) | **GET** /refresh | refreshAndGetAuthenticationToken
*DatacontrollerApi* | [**DataMutationUsingDELETE**](docs/DatacontrollerApi.md#datamutationusingdelete) | **DELETE** /api/{entity}/{id} | dataMutation
*DatacontrollerApi* | [**DataMutationUsingPOST**](docs/DatacontrollerApi.md#datamutationusingpost) | **POST** /api/{entity} | dataMutation
*DatacontrollerApi* | [**DataMutationUsingPUT**](docs/DatacontrollerApi.md#datamutationusingput) | **PUT** /api/{entity}/{id} | dataMutation
*DatacontrollerApi* | [**DataQueryUsingGET**](docs/DatacontrollerApi.md#dataqueryusingget) | **GET** /api/{entity} | dataQuery
*DatacontrollerApi* | [**FindOneUsingGET**](docs/DatacontrollerApi.md#findoneusingget) | **GET** /api/{entity}/{id} | findOne
*DatasourcecontrollerApi* | [**AddDataSourceUsingPOST**](docs/DatasourcecontrollerApi.md#adddatasourceusingpost) | **POST** /datasource/_datasource | addDataSource
*DatasourcecontrollerApi* | [**EditDataSourceUsingPUT**](docs/DatasourcecontrollerApi.md#editdatasourceusingput) | **PUT** /datasource/_datasource/put/{id} | editDataSource
*DatasourcecontrollerApi* | [**FindRoleUsingGET**](docs/DatasourcecontrollerApi.md#findroleusingget) | **GET** /datasource/_datasource/{datasourceId} | findRole
*DatasourcecontrollerApi* | [**ListUsingGET**](docs/DatasourcecontrollerApi.md#listusingget) | **GET** /datasource/_datasource | list
*PermissioncontrollerApi* | [**AddPermissionUsingPOST**](docs/PermissioncontrollerApi.md#addpermissionusingpost) | **POST** /permission/_permission | addPermission
*PermissioncontrollerApi* | [**EditPermissionUsingPUT**](docs/PermissioncontrollerApi.md#editpermissionusingput) | **PUT** /permission/_permission/{id} | editPermission
*PermissioncontrollerApi* | [**FindUserUsingGET**](docs/PermissioncontrollerApi.md#finduserusingget) | **GET** /permission/_permission/{id} | findUser
*PermissioncontrollerApi* | [**ListUsingGET1**](docs/PermissioncontrollerApi.md#listusingget1) | **GET** /permission/_permission | list
*RolecontrollerApi* | [**AddRoleUsingPOST**](docs/RolecontrollerApi.md#addroleusingpost) | **POST** /role/_roles | addRole
*RolecontrollerApi* | [**EditRoleUsingPUT**](docs/RolecontrollerApi.md#editroleusingput) | **PUT** /role/_roles/put/{id} | editRole
*RolecontrollerApi* | [**FindRoleUsingGET1**](docs/RolecontrollerApi.md#findroleusingget1) | **GET** /role/_roles/{roleId} | findRole
*RolecontrollerApi* | [**ListUsingGET2**](docs/RolecontrollerApi.md#listusingget2) | **GET** /role/_roles | list
*SchemacontrollerApi* | [**AddEntityUsingPOST**](docs/SchemacontrollerApi.md#addentityusingpost) | **POST** /schemas/_entitys | addEntity
*SchemacontrollerApi* | [**AddFieldUsingPOST**](docs/SchemacontrollerApi.md#addfieldusingpost) | **POST** /schemas/_fields | addField
*SchemacontrollerApi* | [**EditEntityUsingPUT**](docs/SchemacontrollerApi.md#editentityusingput) | **PUT** /schemas/_entitys/put/{id} | editEntity
*SchemacontrollerApi* | [**EditFieldUsingPUT**](docs/SchemacontrollerApi.md#editfieldusingput) | **PUT** /schemas/_fields/put/{id} | editField
*SchemacontrollerApi* | [**FindAllFieldsUsingGET**](docs/SchemacontrollerApi.md#findallfieldsusingget) | **GET** /schemas/_fields | findAllFields
*SchemacontrollerApi* | [**FindEntityFieldsUsingGET**](docs/SchemacontrollerApi.md#findentityfieldsusingget) | **GET** /schemas/fields | findEntityFields
*SchemacontrollerApi* | [**FindOneFieldUsingGET**](docs/SchemacontrollerApi.md#findonefieldusingget) | **GET** /schemas/_fields/{fid} | findOneField
*SchemacontrollerApi* | [**FindOneUsingGET1**](docs/SchemacontrollerApi.md#findoneusingget1) | **GET** /schemas/_entitys/{eid} | findOne
*SchemacontrollerApi* | [**GetSchemasUsingGET**](docs/SchemacontrollerApi.md#getschemasusingget) | **GET** /schemas/_entitys | getSchemas
*SchemacontrollerApi* | [**ResetCurrentDsUsingPUT**](docs/SchemacontrollerApi.md#resetcurrentdsusingput) | **PUT** /schemas/resetCurrentDs/{dataSourceId} | resetCurrentDs
*SchemacontrollerApi* | [**SyncSchemasUsingPUT**](docs/SchemacontrollerApi.md#syncschemasusingput) | **PUT** /schemas/sync/{dataSourceId} | syncSchemas
*UsercontrollerApi* | [**AddUserUsingPOST**](docs/UsercontrollerApi.md#adduserusingpost) | **POST** /user/_users | addUser
*UsercontrollerApi* | [**EditFieldUsingPUT1**](docs/UsercontrollerApi.md#editfieldusingput1) | **PUT** /user/_users/put/{id} | editField
*UsercontrollerApi* | [**FindUserUsingGET1**](docs/UsercontrollerApi.md#finduserusingget1) | **GET** /user/_users/{userId} | findUser
*UsercontrollerApi* | [**GetAuthenticatedUserUsingGET**](docs/UsercontrollerApi.md#getauthenticateduserusingget) | **GET** /user/me | getAuthenticatedUser
*UsercontrollerApi* | [**ListUsingGET3**](docs/UsercontrollerApi.md#listusingget3) | **GET** /user/_users | list


<a name="documentation-for-models"></a>
## Documentation for Models

 - [Model.Apply](docs/Apply.md)
 - [Model.ChoiceItem](docs/ChoiceItem.md)
 - [Model.DataSource](docs/DataSource.md)
 - [Model.Entity](docs/Entity.md)
 - [Model.Field](docs/Field.md)
 - [Model.GrantedAuthority](docs/GrantedAuthority.md)
 - [Model.IChoiceItem](docs/IChoiceItem.md)
 - [Model.JwtAuthenticationRequest](docs/JwtAuthenticationRequest.md)
 - [Model.JwtUser](docs/JwtUser.md)
 - [Model.Permission](docs/Permission.md)
 - [Model.ResponseEntity](docs/ResponseEntity.md)
 - [Model.Role](docs/Role.md)
 - [Model.User](docs/User.md)


<a name="documentation-for-authorization"></a>
## Documentation for Authorization

All endpoints do not require authorization.