import * as api from './api/api';
import * as angular from 'angular';

const apiModule = angular.module('api', [])
.service('ApplycontrollerApi', api.ApplycontrollerApi)
.service('AuthenticationrestcontrollerApi', api.AuthenticationrestcontrollerApi)
.service('DatacontrollerApi', api.DatacontrollerApi)
.service('DatasourcecontrollerApi', api.DatasourcecontrollerApi)
.service('PermissioncontrollerApi', api.PermissioncontrollerApi)
.service('RolecontrollerApi', api.RolecontrollerApi)
.service('SchemacontrollerApi', api.SchemacontrollerApi)
.service('UsercontrollerApi', api.UsercontrollerApi)

export default apiModule;
