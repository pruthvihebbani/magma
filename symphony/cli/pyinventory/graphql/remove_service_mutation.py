#!/usr/bin/env python3
# @generated AUTOGENERATED file. Do not Change!

from dataclasses import dataclass
from datetime import datetime
from gql.gql.datetime_utils import DATETIME_FIELD
from gql.gql.graphql_client import GraphqlClient
from functools import partial
from numbers import Number
from typing import Any, Callable, List, Mapping, Optional

from dataclasses_json import DataClassJsonMixin


QUERY: List[str] = ["""
mutation RemoveServiceMutation($id: ID!) {
  removeService(id: $id)
}

"""]

@dataclass
class RemoveServiceMutation(DataClassJsonMixin):
    @dataclass
    class RemoveServiceMutationData(DataClassJsonMixin):
        removeService: str

    data: RemoveServiceMutationData

    @classmethod
    # fmt: off
    def execute(cls, client: GraphqlClient, id: str) -> RemoveServiceMutationData:
        # fmt: off
        variables = {"id": id}
        response_text = client.call(''.join(set(QUERY)), variables=variables)
        return cls.from_json(response_text).data
