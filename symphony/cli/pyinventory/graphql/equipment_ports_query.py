#!/usr/bin/env python3
# @generated AUTOGENERATED file. Do not Change!

from dataclasses import dataclass, field
from datetime import datetime
from gql.gql.datetime_utils import fromisoformat
from gql.gql.graphql_client import GraphqlClient
from functools import partial
from numbers import Number
from typing import Any, Callable, List, Mapping, Optional

from dataclasses_json import dataclass_json
from marshmallow import fields as marshmallow_fields


DATETIME_FIELD = field(
    metadata={
        "dataclasses_json": {
            "encoder": datetime.isoformat,
            "decoder": fromisoformat,
            "mm_field": marshmallow_fields.DateTime(format="iso"),
        }
    }
)


@dataclass_json
@dataclass
class EquipmentPortsQuery:
    __QUERY__ = """
    query EquipmentPortsQuery($id: ID!) {
  equipment: node(id: $id) {
    ... on Equipment {
      ports {
        id
        definition {
          id
          name
        }
        link {
          id
        }
      }
    }
  }
}

    """

    @dataclass_json
    @dataclass
    class EquipmentPortsQueryData:
        @dataclass_json
        @dataclass
        class Node:
            @dataclass_json
            @dataclass
            class EquipmentPort:
                @dataclass_json
                @dataclass
                class EquipmentPortDefinition:
                    id: str
                    name: str

                @dataclass_json
                @dataclass
                class Link:
                    id: str

                id: str
                definition: EquipmentPortDefinition
                link: Optional[Link] = None

            ports: List[EquipmentPort]

        equipment: Optional[Node] = None

    data: Optional[EquipmentPortsQueryData] = None
    errors: Optional[Any] = None

    @classmethod
    # fmt: off
    def execute(cls, client: GraphqlClient, id: str):
        # fmt: off
        variables = {"id": id}
        response_text = client.call(cls.__QUERY__, variables=variables)
        return cls.from_json(response_text).data
