/**
 * OpenAPI Petstore
 * This is a sample server Petstore server. For this sample, you can use the api key `special-key` to test the authorization filters.
 *
 * The version of the OpenAPI document: 1.0.0
 * 
 *
 * NOTE: This file is auto generated by Konfig (https://konfigthis.com).
 * Do not edit the class manually.
 */

import { AxiosRequestConfig } from "axios";
import {
  MiscellaneousApi,
  PetApi,
  StoreApi,
  UserApi,
} from "./api";
import { Configuration, ConfigurationParameters } from "./configuration";
import { AcmeCustom } from "./client-custom";

export class Acme extends AcmeCustom {
  readonly miscellaneous: MiscellaneousApi;
  readonly pet: PetApi;
  readonly store: StoreApi;
  readonly user: UserApi;

  constructor(configurationParameters: ConfigurationParameters) {
    super(configurationParameters);
    const configuration = new Configuration(configurationParameters);
    this.miscellaneous = new MiscellaneousApi(configuration);
    this.pet = new PetApi(configuration);
    this.store = new StoreApi(configuration);
    this.user = new UserApi(configuration);
  }

}
