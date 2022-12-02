/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { BaseHttpRequest } from './core/BaseHttpRequest';
import type { ConfigType } from './core/Config';
import { NodeHttpRequest } from './core/NodeHttpRequest';

import { PetService } from './services/PetService';
import { StoreService } from './services/StoreService';
import { UserService } from './services/UserService';

type HttpRequestConstructor = new (config: ConfigType) => BaseHttpRequest;

export class AcmeClient {

    public readonly pet: PetService;
    public readonly store: StoreService;
    public readonly user: UserService;

    public readonly request: BaseHttpRequest;

    constructor(config?: Partial<ConfigType>, HttpRequest: HttpRequestConstructor = NodeHttpRequest) {
        this.request = new HttpRequest({
            BASE: config?.BASE ?? 'http://petstore.swagger.io/v2',
            VERSION: config?.VERSION ?? '1.0.0',
            WITH_CREDENTIALS: config?.WITH_CREDENTIALS ?? false,
            CREDENTIALS: config?.CREDENTIALS ?? 'include',
            TOKEN: config?.TOKEN,
            USERNAME: config?.USERNAME,
            PASSWORD: config?.PASSWORD,
            HEADERS: config?.HEADERS,
            ENCODE_PATH: config?.ENCODE_PATH,
        });

        this.pet = new PetService(this.request);
        this.store = new StoreService(this.request);
        this.user = new UserService(this.request);
    }
}

