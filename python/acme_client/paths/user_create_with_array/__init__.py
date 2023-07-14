# do not import all endpoints into this module because that uses a lot of memory and stack frames
# if you need the ability to import all endpoints from this module, import them with
# from acme_client.paths.user_create_with_array import Api

from acme_client.paths import PathValues

path = PathValues.USER_CREATE_WITH_ARRAY