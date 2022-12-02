/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { ApiResponse } from '../models/ApiResponse';
import type { Pet } from '../models/Pet';

import type { CancelablePromise } from '../core/CancelablePromise';
import type { BaseHttpRequest } from '../core/BaseHttpRequest';

export class PetService {

    constructor(public readonly httpRequest: BaseHttpRequest) {}

    /**
     * Add a new pet to the store
     * @returns Pet successful operation
     * @throws ApiError
     */
    public addPet({
        requestBody,
    }: {
        /**
         * Pet object that needs to be added to the store
         */
        requestBody: Pet,
    }): CancelablePromise<Pet> {
        return this.httpRequest.request({
            method: 'POST',
            url: '/pet',
            body: requestBody,
            mediaType: 'application/json',
            errors: {
                405: `Invalid input`,
            },
        });
    }

    /**
     * Update an existing pet
     * @returns Pet successful operation
     * @throws ApiError
     */
    public updatePet({
        requestBody,
    }: {
        /**
         * Pet object that needs to be added to the store
         */
        requestBody: Pet,
    }): CancelablePromise<Pet> {
        return this.httpRequest.request({
            method: 'PUT',
            url: '/pet',
            body: requestBody,
            mediaType: 'application/json',
            errors: {
                400: `Invalid ID supplied`,
                404: `Pet not found`,
                405: `Validation exception`,
            },
        });
    }

    /**
     * Finds Pets by status
     * Multiple status values can be provided with comma separated strings
     * @returns Pet successful operation
     * @throws ApiError
     */
    public findPetsByStatus({
        status,
    }: {
        /**
         * Status values that need to be considered for filter
         * @deprecated
         */
        status: Array<'available' | 'pending' | 'sold'>,
    }): CancelablePromise<Array<Pet>> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/pet/findByStatus',
            query: {
                'status': status,
            },
            errors: {
                400: `Invalid status value`,
            },
        });
    }

    /**
     * @deprecated
     * Finds Pets by tags
     * Multiple tags can be provided with comma separated strings. Use tag1, tag2, tag3 for testing.
     * @returns Pet successful operation
     * @throws ApiError
     */
    public findPetsByTags({
        tags,
    }: {
        /**
         * Tags to filter by
         */
        tags: Array<string>,
    }): CancelablePromise<Array<Pet>> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/pet/findByTags',
            query: {
                'tags': tags,
            },
            errors: {
                400: `Invalid tag value`,
            },
        });
    }

    /**
     * Find pet by ID
     * Returns a single pet
     * @returns Pet successful operation
     * @throws ApiError
     */
    public getPetById({
        petId,
    }: {
        /**
         * ID of pet to return
         */
        petId: number,
    }): CancelablePromise<Pet> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/pet/{petId}',
            path: {
                'petId': petId,
            },
            errors: {
                400: `Invalid ID supplied`,
                404: `Pet not found`,
            },
        });
    }

    /**
     * Updates a pet in the store with form data
     * @returns void
     * @throws ApiError
     */
    public updatePetWithForm({
        petId,
        formData,
    }: {
        /**
         * ID of pet that needs to be updated
         */
        petId: number,
        formData?: {
            /**
             * Updated name of the pet
             */
            name?: string;
            /**
             * Updated status of the pet
             */
            status?: string;
        },
    }): CancelablePromise<void> {
        return this.httpRequest.request({
            method: 'POST',
            url: '/pet/{petId}',
            path: {
                'petId': petId,
            },
            formData: formData,
            mediaType: 'application/x-www-form-urlencoded',
            errors: {
                405: `Invalid input`,
            },
        });
    }

    /**
     * Deletes a pet
     * @returns void
     * @throws ApiError
     */
    public deletePet({
        petId,
        apiKey,
    }: {
        /**
         * Pet id to delete
         */
        petId: number,
        apiKey?: string,
    }): CancelablePromise<void> {
        return this.httpRequest.request({
            method: 'DELETE',
            url: '/pet/{petId}',
            path: {
                'petId': petId,
            },
            headers: {
                'api_key': apiKey,
            },
            errors: {
                400: `Invalid pet value`,
            },
        });
    }

    /**
     * uploads an image
     * @returns ApiResponse successful operation
     * @throws ApiError
     */
    public uploadFile({
        petId,
        formData,
    }: {
        /**
         * ID of pet to update
         */
        petId: number,
        formData?: {
            /**
             * Additional data to pass to server
             */
            additionalMetadata?: string;
            /**
             * file to upload
             */
            file?: Blob;
        },
    }): CancelablePromise<ApiResponse> {
        return this.httpRequest.request({
            method: 'POST',
            url: '/pet/{petId}/uploadImage',
            path: {
                'petId': petId,
            },
            formData: formData,
            mediaType: 'multipart/form-data',
        });
    }

}
