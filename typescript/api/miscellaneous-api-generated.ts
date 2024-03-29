/* tslint:disable */
/* eslint-disable */
/*
OpenAPI Petstore

This is a sample server Petstore server. For this sample, you can use the api key `special-key` to test the authorization filters.

The version of the OpenAPI document: 1.0.0


NOTE: This file is auto generated by Konfig (https://konfigthis.com).
*/

import globalAxios, { AxiosPromise, AxiosInstance, AxiosRequestConfig } from 'axios';
import { Configuration } from '../configuration';
// Some imports not used depending on template conditions
// @ts-ignore
import { DUMMY_BASE_URL, assertParamExists, setApiKeyToObject, setBasicAuthToObject, setBearerAuthToObject, setSearchParams, serializeDataIfNeeded, toPathString, createRequestFunction, isBrowser } from '../common';
import { fromBuffer } from "file-type/browser"
const FormData = require("form-data")
import { setOAuthToObject } from '../common';
// @ts-ignore
import { BASE_PATH, COLLECTION_FORMATS, RequestArgs, BaseAPI, RequiredError } from '../base';
// @ts-ignore
import { PaginateRequest } from '../models';
// @ts-ignore
import { PaginateResponse } from '../models';
import { paginate } from "../pagination/paginate";
import type * as buffer from "buffer"
import { requestBeforeHook } from '../requestBeforeHook';
/**
 * MiscellaneousApi - axios parameter creator
 * @export
 */
export const MiscellaneousApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        /**
         * Iterate through a bunch of items
         * @summary Pagination sandbox
         * @param {PaginateRequest} [paginateRequest] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        paginate: async (paginateRequest?: PaginateRequest, options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            const localVarPath = `/pagination`;
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions: AxiosRequestConfig = { method: 'GET', ...baseOptions, ...options};
            const localVarHeaderParameter = configuration && !isBrowser() ? { "User-Agent": configuration.userAgent } : {} as any;
            const localVarQueryParameter = {} as any;


    
            localVarHeaderParameter['Content-Type'] = 'application/json';


            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};
            requestBeforeHook({
                requestBody: paginateRequest,
                queryParameters: localVarQueryParameter,
                requestConfig: localVarRequestOptions,
                path: localVarPath,
                configuration
            });
            localVarRequestOptions.data = serializeDataIfNeeded(paginateRequest, localVarRequestOptions, configuration)

            setSearchParams(localVarUrlObj, localVarQueryParameter);
            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
    }
};

/**
 * MiscellaneousApi - functional programming interface
 * @export
 */
export const MiscellaneousApiFp = function(configuration?: Configuration) {
    const localVarAxiosParamCreator = MiscellaneousApiAxiosParamCreator(configuration)
    return {
        /**
         * Iterate through a bunch of items
         * @summary Pagination sandbox
         * @param {MiscellaneousApiPaginateRequest} requestParameters Request parameters.
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async paginate(requestParameters: MiscellaneousApiPaginateRequest = {}, options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<PaginateResponse>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.paginate(requestParameters, options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
    }
};

/**
 * MiscellaneousApi - factory interface
 * @export
 */
export const MiscellaneousApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    const localVarFp = MiscellaneousApiFp(configuration)
    return {
        /**
         * Iterate through a bunch of items
         * @summary Pagination sandbox
         * @param {MiscellaneousApiPaginateRequest} requestParameters Request parameters.
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        paginate(requestParameters: MiscellaneousApiPaginateRequest = {}, options?: AxiosRequestConfig): AxiosPromise<PaginateResponse> {
            return localVarFp.paginate(requestParameters, options).then((request) => request(axios, basePath));
        },
    };
};

/**
 * Request parameters for paginate operation in MiscellaneousApi.
 * @export
 * @interface MiscellaneousApiPaginateRequest
 */
export type MiscellaneousApiPaginateRequest = {
    
} & PaginateRequest

/**
 * MiscellaneousApiGenerated - object-oriented interface
 * @export
 * @class MiscellaneousApiGenerated
 * @extends {BaseAPI}
 */
export class MiscellaneousApiGenerated extends BaseAPI {
    /**
     * Iterate through a bunch of items
     * @summary Pagination sandbox
     * @param {MiscellaneousApiPaginateRequest} requestParameters Request parameters.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof MiscellaneousApiGenerated
     */
    public paginate(requestParameters: MiscellaneousApiPaginateRequest = {}, options?: AxiosRequestConfig) {
        return MiscellaneousApiFp(this.configuration).paginate(requestParameters, options).then((request) => request(this.axios, this.basePath));
    }
}
