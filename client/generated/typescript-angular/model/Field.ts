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

import * as models from './models';

export interface Field {
    "choices"?: Array<models.IChoiceItem>;
    "component"?: Field.ComponentEnum;
    "dataSourceId"?: string;
    "defaultValue"?: string;
    "eid"?: string;
    "id"?: string;
    "inputType"?: Field.InputTypeEnum;
    "isAutoIncremented"?: boolean;
    "isPartOfPrimaryKey"?: boolean;
    "label"?: string;
    "maxLength"?: number;
    "maxValue"?: string;
    "minValue"?: string;
    "name"?: string;
    "reference"?: string;
    "referenceOptionText"?: string;
    "required"?: boolean;
    "showInCreate"?: boolean;
    "showInEdit"?: boolean;
    "showInFilter"?: boolean;
    "showInList"?: boolean;
    "showInShow"?: boolean;
}

export namespace Field {
    export enum ComponentEnum {
        Autocomplete = <any> 'Autocomplete',
        Boolean = <any> 'Boolean',
        NullableBoolean = <any> 'NullableBoolean',
        CheckboxGroup = <any> 'CheckboxGroup',
        Date = <any> 'Date',
        Disabled = <any> 'Disabled',
        File = <any> 'File',
        Image = <any> 'Image',
        LongText = <any> 'LongText',
        Number = <any> 'Number',
        RadioButtonGroup = <any> 'RadioButtonGroup',
        ReferenceArray = <any> 'ReferenceArray',
        Reference = <any> 'Reference',
        RichText = <any> 'RichText',
        SelectArray = <any> 'SelectArray',
        Select = <any> 'Select',
        Text = <any> 'Text'
    }
    export enum InputTypeEnum {
        Text = <any> 'text',
        Email = <any> 'email',
        Password = <any> 'password',
        Url = <any> 'url'
    }
}