package com.acme.client;

import com.acme.client.api.MiscellaneousApi;
import com.acme.client.api.PetApi;
import com.acme.client.api.StoreApi;
import com.acme.client.api.UserApi;

public class Java {
    private ApiClient apiClient;
    public final MiscellaneousApi miscellaneous;
    public final PetApi pet;
    public final StoreApi store;
    public final UserApi user;

    public Java() {
        this(null);
    }

    public Java(Configuration configuration) {
        this.apiClient = new ApiClient(null, configuration);
        this.miscellaneous = new MiscellaneousApi(this.apiClient);
        this.pet = new PetApi(this.apiClient);
        this.store = new StoreApi(this.apiClient);
        this.user = new UserApi(this.apiClient);
    }

}
