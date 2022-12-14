/*
OpenAPI Petstore

Testing StoreApiService

*/

// Code generated by Konfig (https://konfigthis.com);

package acme

import (
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "github.com/konfig-dev/acme-sdks/go"
)

func Test_acme_StoreApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test StoreApiService DeleteOrder", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var orderId string

        resp, httpRes, err := apiClient.StoreApi.DeleteOrder(orderId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test StoreApiService GetInventory", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.StoreApi.GetInventory().Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test StoreApiService GetOrderById", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var orderId int64

        resp, httpRes, err := apiClient.StoreApi.GetOrderById(orderId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test StoreApiService PlaceOrder", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.StoreApi.PlaceOrder().Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
