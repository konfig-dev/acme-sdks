import time
import sys
from pathlib import Path
python_sdk_dir = str(Path(__file__).parents[1]) + '/sdks/python'
sys.path.append(python_sdk_dir)
import acme_client
from pprint import pprint
from acme_client.api import pet_api
from acme_client.model.api_response import ApiResponse
from acme_client.model.tag import Tag
from acme_client.model.category import Category
from acme_client.model.pet import Pet
# Defining the host is optional and defaults to http://petstore.swagger.io/v2
# See configuration.py for a list of all supported configuration parameters.
configuration = acme_client.Configuration(
    host = "http://petstore.swagger.io/v2"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure OAuth2 access token for authorization: petstore_auth
configuration = acme_client.Configuration(
    host = "http://petstore.swagger.io/v2"
)
configuration.access_token = 'YOUR_ACCESS_TOKEN'


# Enter a context with an instance of the API client
with acme_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = pet_api.PetApi(api_client)
    pet = Pet(
        id=1,
        category=Category(
            id=1,
            name="CbUUGjjNSwg0_bs9ZayIMrKdgNvb6gvxmPb9GcsM61ate1RA89q3w1l4eH4XxEz.5awLMdeXylwK0lMGUSM4jsrh4dstlnQUN5vVdMLPA",
        ),
        name="doggie",
        photo_urls=[
            "photo_urls_example",
        ],
        tags=[
            Tag(
                id=1,
                name="name_example",
            ),
        ],
        status="available",
    ) # Pet | Pet object that needs to be added to the store

    try:
        # Add a new pet to the store
        api_response = api_instance.add_pet(pet)
        pprint(api_response)
    except acme_client.ApiException as e:
        print("Exception when calling PetApi->add_pet: %s\n" % e)