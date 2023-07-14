/*
 * OpenAPI Petstore
 * This is a sample server Petstore server. For this sample, you can use the api key `special-key` to test the authorization filters.
 *
 * The version of the OpenAPI document: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by Konfig (https://konfigthis.com).
 * Do not edit the class manually.
 */


package com.acme.client.api;

import com.acme.client.ApiException;
import com.acme.client.ApiClient;
import com.acme.client.ApiException;
import com.acme.client.Configuration;
import java.time.OffsetDateTime;
import com.acme.client.model.User;
import com.acme.client.model.UserCreateRequest;
import org.junit.jupiter.api.Disabled;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.BeforeAll;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * API tests for UserApi
 */
@Disabled
public class UserApiTest {

    private static UserApi api;

    
    @BeforeAll
    public static void beforeClass() {
        ApiClient apiClient = Configuration.getDefaultApiClient();
        api = new UserApi(apiClient);
    }

    /**
     * Create user
     *
     * This can only be done by the logged in user.
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void createTest() throws ApiException {
        Long id = null;
        String username = null;
        String firstName = null;
        String lastName = null;
        String email = null;
        String password = null;
        String phone = null;
        Integer userStatus = null;
        api.create()
                .id(id)
                .username(username)
                .firstName(firstName)
                .lastName(lastName)
                .email(email)
                .password(password)
                .phone(phone)
                .userStatus(userStatus)
                .execute();
        // TODO: test validations
    }

    /**
     * Creates list of users with given input array
     *
     * 
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void createWithArrayTest() throws ApiException {
        api.createWithArray()
                .execute();
        // TODO: test validations
    }

    /**
     * Creates list of users with given input array
     *
     * 
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void createWithListTest() throws ApiException {
        api.createWithList()
                .execute();
        // TODO: test validations
    }

    /**
     * Delete user
     *
     * This can only be done by the logged in user.
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void deleteTest() throws ApiException {
        String username = null;
        api.delete(username)
                .execute();
        // TODO: test validations
    }

    /**
     * Get user by user name
     *
     * 
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void getByNameTest() throws ApiException {
        String username = null;
        User response = api.getByName(username)
                .execute();
        // TODO: test validations
    }

    /**
     * Logs user into the system
     *
     * 
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void loginTest() throws ApiException {
        String username = null;
        String password = null;
        String response = api.login(username, password)
                .execute();
        // TODO: test validations
    }

    /**
     * Logs out current logged in user session
     *
     * 
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void logoutTest() throws ApiException {
        api.logout()
                .execute();
        // TODO: test validations
    }

    /**
     * Updated user
     *
     * This can only be done by the logged in user.
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void updateTest() throws ApiException {
        String username = null;
        Long id = null;
        String firstName = null;
        String lastName = null;
        String email = null;
        String password = null;
        String phone = null;
        Integer userStatus = null;
        api.update(username)
                .id(id)
                .firstName(firstName)
                .lastName(lastName)
                .email(email)
                .password(password)
                .phone(phone)
                .userStatus(userStatus)
                .execute();
        // TODO: test validations
    }

}
