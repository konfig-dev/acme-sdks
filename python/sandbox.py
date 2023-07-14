from pprint import pprint
from acme_client import Acme, ApiException

acme = Acme(
    # Defining the host is optional and defaults to http://petstore.swagger.io/v2
    # See configuration.py for a list of all supported configuration parameters.
    host="http://petstore.swagger.io/v2",
)

try:
    # Pagination sandbox
    paginate_response = acme.miscellaneous.paginate(
        first=1,  # optional
        after="string_example",  # optional
    )
    pprint(paginate_response.body)
    pprint(paginate_response.body["edges"])
    pprint(paginate_response.body["pageInfo"])
    pprint(paginate_response.headers)
    pprint(paginate_response.status)
    pprint(paginate_response.round_trip_time)
except ApiException as e:
    print("Exception when calling MiscellaneousApi.paginate: %s\n" % e)
    pprint(e.body)
    pprint(e.headers)
    pprint(e.status)
    pprint(e.reason)
    pprint(e.round_trip_time)
