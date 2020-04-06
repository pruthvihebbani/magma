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
query LocationDepsQuery($id: ID!) {
  location: node(id: $id) {
    ... on Location {
      files {
        id
      }
      images {
        id
      }
      children {
        id
      }
      surveys {
        id
      }
      equipments {
        id
      }
    }
  }
}

"""]

@dataclass
class LocationDepsQuery(DataClassJsonMixin):
    @dataclass
    class LocationDepsQueryData(DataClassJsonMixin):
        @dataclass
        class Node(DataClassJsonMixin):
            @dataclass
            class File(DataClassJsonMixin):
                id: str

            @dataclass
            class Location(DataClassJsonMixin):
                id: str

            @dataclass
            class Survey(DataClassJsonMixin):
                id: str

            @dataclass
            class Equipment(DataClassJsonMixin):
                id: str

            files: List[File]
            images: List[File]
            children: List[Location]
            surveys: List[Survey]
            equipments: List[Equipment]

        location: Optional[Node]

    data: LocationDepsQueryData

    @classmethod
    # fmt: off
    def execute(cls, client: GraphqlClient, id: str) -> LocationDepsQueryData:
        # fmt: off
        variables = {"id": id}
        response_text = client.call(''.join(set(QUERY)), variables=variables)
        return cls.from_json(response_text).data
