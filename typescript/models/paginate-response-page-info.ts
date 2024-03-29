/*
OpenAPI Petstore

This is a sample server Petstore server. For this sample, you can use the api key `special-key` to test the authorization filters.

The version of the OpenAPI document: 1.0.0


NOTE: This file is auto generated by Konfig (https://konfigthis.com).
*/
import type * as buffer from "buffer"


/**
 * 
 * @export
 * @interface PaginateResponsePageInfo
 */
export interface PaginateResponsePageInfo {
    /**
     * 
     * @type {boolean}
     * @memberof PaginateResponsePageInfo
     */
    'hasNextPage'?: boolean;
    /**
     * 
     * @type {boolean}
     * @memberof PaginateResponsePageInfo
     */
    'hasPreviousPage'?: boolean;
    /**
     * 
     * @type {string}
     * @memberof PaginateResponsePageInfo
     */
    'startCursor'?: string;
    /**
     * 
     * @type {string}
     * @memberof PaginateResponsePageInfo
     */
    'endCursor'?: string;
}

