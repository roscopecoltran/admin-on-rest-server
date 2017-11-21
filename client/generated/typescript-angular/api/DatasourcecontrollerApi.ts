/**
 * DataHive RESTful APIs
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * OpenAPI spec version: 1.0
 * 
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 * Do not edit the class manually.
 */

import * as models from '../model/models';

/* tslint:disable:no-unused-variable member-ordering */

export class DatasourcecontrollerApi {
    protected basePath = 'https://localhost:8080';
    public defaultHeaders : any = {};

    static $inject: string[] = ['$http', '$httpParamSerializer', 'basePath'];

    constructor(protected $http: ng.IHttpService, protected $httpParamSerializer?: (d: any) => any, basePath?: string) {
        if (basePath !== undefined) {
            this.basePath = basePath;
        }
    }

    /**
     * 
     * @summary addDataSource
     * @param dataSource dataSource
     */
    public addDataSourceUsingPOST (dataSource: models.DataSource, extraHttpRequestParams?: any ) : ng.IHttpPromise<models.DataSource> {
        const localVarPath = this.basePath + '/datasource/_datasource';

        let queryParameters: any = {};
        let headerParams: any = (<any>Object).assign({}, this.defaultHeaders);
        // verify required parameter 'dataSource' is not null or undefined
        if (dataSource === null || dataSource === undefined) {
            throw new Error('Required parameter dataSource was null or undefined when calling addDataSourceUsingPOST.');
        }
        let httpRequestParams: ng.IRequestConfig = {
            method: 'POST',
            url: localVarPath,
            data: dataSource,
                        params: queryParameters,
            headers: headerParams
        };

        if (extraHttpRequestParams) {
            httpRequestParams = (<any>Object).assign(httpRequestParams, extraHttpRequestParams);
        }

        return this.$http(httpRequestParams);
    }
    /**
     * 
     * @summary editDataSource
     * @param id id
     * @param dataSource dataSource
     */
    public editDataSourceUsingPUT (id: string, dataSource: models.DataSource, extraHttpRequestParams?: any ) : ng.IHttpPromise<models.DataSource> {
        const localVarPath = this.basePath + '/datasource/_datasource/put/{id}'
            .replace('{' + 'id' + '}', String(id));

        let queryParameters: any = {};
        let headerParams: any = (<any>Object).assign({}, this.defaultHeaders);
        // verify required parameter 'id' is not null or undefined
        if (id === null || id === undefined) {
            throw new Error('Required parameter id was null or undefined when calling editDataSourceUsingPUT.');
        }
        // verify required parameter 'dataSource' is not null or undefined
        if (dataSource === null || dataSource === undefined) {
            throw new Error('Required parameter dataSource was null or undefined when calling editDataSourceUsingPUT.');
        }
        let httpRequestParams: ng.IRequestConfig = {
            method: 'PUT',
            url: localVarPath,
            data: dataSource,
                        params: queryParameters,
            headers: headerParams
        };

        if (extraHttpRequestParams) {
            httpRequestParams = (<any>Object).assign(httpRequestParams, extraHttpRequestParams);
        }

        return this.$http(httpRequestParams);
    }
    /**
     * 
     * @summary findRole
     * @param datasourceId datasourceId
     */
    public findRoleUsingGET (datasourceId: string, extraHttpRequestParams?: any ) : ng.IHttpPromise<models.DataSource> {
        const localVarPath = this.basePath + '/datasource/_datasource/{datasourceId}'
            .replace('{' + 'datasourceId' + '}', String(datasourceId));

        let queryParameters: any = {};
        let headerParams: any = (<any>Object).assign({}, this.defaultHeaders);
        // verify required parameter 'datasourceId' is not null or undefined
        if (datasourceId === null || datasourceId === undefined) {
            throw new Error('Required parameter datasourceId was null or undefined when calling findRoleUsingGET.');
        }
        let httpRequestParams: ng.IRequestConfig = {
            method: 'GET',
            url: localVarPath,
                                    params: queryParameters,
            headers: headerParams
        };

        if (extraHttpRequestParams) {
            httpRequestParams = (<any>Object).assign(httpRequestParams, extraHttpRequestParams);
        }

        return this.$http(httpRequestParams);
    }
    /**
     * 
     * @summary list
     */
    public listUsingGET (extraHttpRequestParams?: any ) : ng.IHttpPromise<Array<models.DataSource>> {
        const localVarPath = this.basePath + '/datasource/_datasource';

        let queryParameters: any = {};
        let headerParams: any = (<any>Object).assign({}, this.defaultHeaders);
        let httpRequestParams: ng.IRequestConfig = {
            method: 'GET',
            url: localVarPath,
                                    params: queryParameters,
            headers: headerParams
        };

        if (extraHttpRequestParams) {
            httpRequestParams = (<any>Object).assign(httpRequestParams, extraHttpRequestParams);
        }

        return this.$http(httpRequestParams);
    }
}
