# coding: utf-8

"""
    OpenAPI Petstore

    This is a sample server Petstore server. For this sample, you can use the api key `special-key` to test the authorization filters.

    The version of the OpenAPI document: 1.0.0
    Generated by: https://konfigthis.com
"""

from datetime import date, datetime  # noqa: F401
import decimal  # noqa: F401
import functools  # noqa: F401
import io  # noqa: F401
import re  # noqa: F401
import typing  # noqa: F401
import typing_extensions  # noqa: F401
import uuid  # noqa: F401

import frozendict  # noqa: F401

from acme_client import schemas  # noqa: F401


class PaginateResponsePageInfo(
    schemas.DictSchema
):
    """
    This class is auto generated by Konfig (https://konfigthis.com)
    """


    class MetaOapg:
        
        class properties:
            hasNextPage = schemas.BoolSchema
            hasPreviousPage = schemas.BoolSchema
            startCursor = schemas.StrSchema
            endCursor = schemas.StrSchema
            __annotations__ = {
                "hasNextPage": hasNextPage,
                "hasPreviousPage": hasPreviousPage,
                "startCursor": startCursor,
                "endCursor": endCursor,
            }
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["hasNextPage"]) -> MetaOapg.properties.hasNextPage: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["hasPreviousPage"]) -> MetaOapg.properties.hasPreviousPage: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["startCursor"]) -> MetaOapg.properties.startCursor: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["endCursor"]) -> MetaOapg.properties.endCursor: ...
    
    @typing.overload
    def __getitem__(self, name: str) -> schemas.UnsetAnyTypeSchema: ...
    
    def __getitem__(self, name: typing.Union[typing_extensions.Literal["hasNextPage", "hasPreviousPage", "startCursor", "endCursor", ], str]):
        # dict_instance[name] accessor
        return super().__getitem__(name)
    
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["hasNextPage"]) -> typing.Union[MetaOapg.properties.hasNextPage, schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["hasPreviousPage"]) -> typing.Union[MetaOapg.properties.hasPreviousPage, schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["startCursor"]) -> typing.Union[MetaOapg.properties.startCursor, schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["endCursor"]) -> typing.Union[MetaOapg.properties.endCursor, schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: str) -> typing.Union[schemas.UnsetAnyTypeSchema, schemas.Unset]: ...
    
    def get_item_oapg(self, name: typing.Union[typing_extensions.Literal["hasNextPage", "hasPreviousPage", "startCursor", "endCursor", ], str]):
        return super().get_item_oapg(name)
    

    def __new__(
        cls,
        *args: typing.Union[dict, frozendict.frozendict, ],
        hasNextPage: typing.Union[MetaOapg.properties.hasNextPage, bool, schemas.Unset] = schemas.unset,
        hasPreviousPage: typing.Union[MetaOapg.properties.hasPreviousPage, bool, schemas.Unset] = schemas.unset,
        startCursor: typing.Union[MetaOapg.properties.startCursor, str, schemas.Unset] = schemas.unset,
        endCursor: typing.Union[MetaOapg.properties.endCursor, str, schemas.Unset] = schemas.unset,
        _configuration: typing.Optional[schemas.Configuration] = None,
        **kwargs: typing.Union[schemas.AnyTypeSchema, dict, frozendict.frozendict, str, date, datetime, uuid.UUID, int, float, decimal.Decimal, None, list, tuple, bytes],
    ) -> 'PaginateResponsePageInfo':
        return super().__new__(
            cls,
            *args,
            hasNextPage=hasNextPage,
            hasPreviousPage=hasPreviousPage,
            startCursor=startCursor,
            endCursor=endCursor,
            _configuration=_configuration,
            **kwargs,
        )
