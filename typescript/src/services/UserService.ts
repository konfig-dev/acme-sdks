/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { User } from '../models/User';

import type { CancelablePromise } from '../core/CancelablePromise';
import type { BaseHttpRequest } from '../core/BaseHttpRequest';

export class UserService {

    constructor(public readonly httpRequest: BaseHttpRequest) {}

    /**
     * Create user
     * This can only be done by the logged in user.
     * @returns any successful operation
     * @throws ApiError
     */
    public createUser({
        requestBody,
    }: {
        /**
         * Created user object
         */
        requestBody: User,
    }): CancelablePromise<any> {
        return this.httpRequest.request({
            method: 'POST',
            url: '/user',
            body: requestBody,
            mediaType: 'application/json',
        });
    }

    /**
     * Creates list of users with given input array
     * @returns any successful operation
     * @throws ApiError
     */
    public createUsersWithArrayInput({
        requestBody,
    }: {
        /**
         * List of user object
         */
        requestBody: Array<User>,
    }): CancelablePromise<any> {
        return this.httpRequest.request({
            method: 'POST',
            url: '/user/createWithArray',
            body: requestBody,
            mediaType: 'application/json',
        });
    }

    /**
     * Creates list of users with given input array
     * @returns any successful operation
     * @throws ApiError
     */
    public createUsersWithListInput({
        requestBody,
    }: {
        /**
         * List of user object
         */
        requestBody: Array<User>,
    }): CancelablePromise<any> {
        return this.httpRequest.request({
            method: 'POST',
            url: '/user/createWithList',
            body: requestBody,
            mediaType: 'application/json',
        });
    }

    /**
     * Logs user into the system
     * @returns string successful operation
     * @throws ApiError
     */
    public loginUser({
        username,
        password,
    }: {
        /**
         * The user name for login
         */
        username: string,
        /**
         * The password for login in clear text
         */
        password: string,
    }): CancelablePromise<string> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/user/login',
            query: {
                'username': username,
                'password': password,
            },
            errors: {
                400: `Invalid username/password supplied`,
            },
        });
    }

    /**
     * Logs out current logged in user session
     * @returns any successful operation
     * @throws ApiError
     */
    public logoutUser(): CancelablePromise<any> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/user/logout',
        });
    }

    /**
     * Get user by user name
     * @returns User successful operation
     * @throws ApiError
     */
    public getUserByName({
        username,
    }: {
        /**
         * The name that needs to be fetched. Use user1 for testing.
         */
        username: string,
    }): CancelablePromise<User> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/user/{username}',
            path: {
                'username': username,
            },
            errors: {
                400: `Invalid username supplied`,
                404: `User not found`,
            },
        });
    }

    /**
     * Updated user
     * This can only be done by the logged in user.
     * @returns void
     * @throws ApiError
     */
    public updateUser({
        username,
        requestBody,
    }: {
        /**
         * name that need to be deleted
         */
        username: string,
        /**
         * Updated user object
         */
        requestBody: User,
    }): CancelablePromise<void> {
        return this.httpRequest.request({
            method: 'PUT',
            url: '/user/{username}',
            path: {
                'username': username,
            },
            body: requestBody,
            mediaType: 'application/json',
            errors: {
                400: `Invalid user supplied`,
                404: `User not found`,
            },
        });
    }

    /**
     * Delete user
     * This can only be done by the logged in user.
     * @returns void
     * @throws ApiError
     */
    public deleteUser({
        username,
    }: {
        /**
         * The name that needs to be deleted
         */
        username: string,
    }): CancelablePromise<void> {
        return this.httpRequest.request({
            method: 'DELETE',
            url: '/user/{username}',
            path: {
                'username': username,
            },
            errors: {
                400: `Invalid username supplied`,
                404: `User not found`,
            },
        });
    }

}
