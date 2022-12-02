/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { ApiRequestOptions } from './ApiRequestOptions';
import type { CancelablePromise } from './CancelablePromise';
import type { ConfigType } from './Config';

export abstract class BaseHttpRequest {

    constructor(public readonly config: ConfigType) {}

    public abstract request<T>(options: ApiRequestOptions): CancelablePromise<T>;
}
