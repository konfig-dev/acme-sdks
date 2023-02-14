/*
OpenAPI Petstore

Testing PetApiService

*/

// Code generated by Konfig (https://konfigthis.com);

package acme

import (
    "os"
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    acme "github.com/konfig-dev/acme-sdks/go"
)

func Test_acme_PetApiService(t *testing.T) {

    configuration := acme.NewConfiguration()
    apiKey := os.Getenv("API_KEY")
    configuration.Context = context.WithValue(configuration.Context, acme.ContextAPIKeys, map[string]acme.APIKey{
        "api_key": {Key: apiKey},
    })
    apiClient := acme.NewAPIClient(configuration)

    t.Run("Test PetApiService AddPet", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.PetApi.AddPet().Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test PetApiService Delete", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var petId int64

        resp, httpRes, err := apiClient.PetApi.Delete(petId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test PetApiService FindPetsByStatus", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.PetApi.FindPetsByStatus().Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test PetApiService FindPetsByTags", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.PetApi.FindPetsByTags().Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test PetApiService GetPetById", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var petId int64

        resp, httpRes, err := apiClient.PetApi.GetPetById(petId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test PetApiService UpdatePet", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.PetApi.UpdatePet().Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test PetApiService UpdatePetWithForm", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var petId int64

        resp, httpRes, err := apiClient.PetApi.UpdatePetWithForm(petId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test PetApiService UploadFile", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var petId int64

        resp, httpRes, err := apiClient.PetApi.UploadFile(petId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
