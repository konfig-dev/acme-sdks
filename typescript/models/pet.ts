/*
OpenAPI Petstore

This is a sample server Petstore server. For this sample, you can use the api key `special-key` to test the authorization filters.

The version of the OpenAPI document: 1.0.0


NOTE: This file is auto generated by Konfig (https://konfigthis.com).
*/
import type * as buffer from "buffer"

import { Category } from './category';
import { Tag } from './tag';

/**
 * A pet for sale in the pet store
 * @export
 * @interface Pet
 */
export interface Pet {
    /**
     * 
     * @type {Array<Tag>}
     * @memberof Pet
     */
    'tags'?: Array<Tag>;
    /**
     * 
     * @type {number}
     * @memberof Pet
     */
    'id'?: number;
    /**
     * 
     * @type {Category}
     * @memberof Pet
     */
    'category'?: Category;
    /**
     * 
     * @type {string}
     * @memberof Pet
     */
    'name': string;
    /**
     * 
     * @type {Array<string>}
     * @memberof Pet
     */
    'photoUrls': Array<string>;
    /**
     * pet status in the store
     * @type {string}
     * @memberof Pet
     * @deprecated
     */
    'status'?: PetStatusEnum;
}

type PetStatusEnum = 'available' | 'pending' | 'sold'


