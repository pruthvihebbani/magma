#!/usr/bin/env python3
# @generated AUTOGENERATED file. Do not Change!

from dataclasses import dataclass, field
from datetime import datetime
from enum import Enum
from functools import partial
from typing import Any, Callable, List, Mapping, Optional

from dataclasses_json import dataclass_json
from marshmallow import fields as marshmallow_fields

from .datetime_utils import fromisoformat


DATETIME_FIELD = field(
    metadata={
        "dataclasses_json": {
            "encoder": datetime.isoformat,
            "decoder": fromisoformat,
            "mm_field": marshmallow_fields.DateTime(format="iso"),
        }
    }
)


def enum_field(enum_type):
    def encode_enum(value):
        return value.value

    def decode_enum(t, value):
        return t(value)

    return field(
        metadata={
            "dataclasses_json": {
                "encoder": encode_enum,
                "decoder": partial(decode_enum, enum_type),
            }
        }
    )


class PropertyKind(Enum):
    string = "string"
    int = "int"
    bool = "bool"
    float = "float"
    date = "date"
    enum = "enum"
    range = "range"
    email = "email"
    gps_location = "gps_location"
    equipment = "equipment"
    location = "location"
    datetime_local = "datetime_local"


@dataclass_json
@dataclass
class EquipmentTypesQuery:
    __QUERY__ = """
    query EquipmentTypesQuery {
  equipmentTypes {
    edges {
      node {
        id
        name
        category
        propertyTypes {
          id
          name
          type
          index
          stringValue
          intValue
          booleanValue
          floatValue
          latitudeValue
          longitudeValue
          isEditable
          isInstanceProperty
        }
        positionDefinitions {
          id
          name
          index
          visibleLabel
        }
        portDefinitions {
          id
          name
          index
          visibleLabel
        }
      }
    }
  }
}

    """

    @dataclass_json
    @dataclass
    class EquipmentTypesQueryData:
        @dataclass_json
        @dataclass
        class EquipmentTypeConnection:
            @dataclass_json
            @dataclass
            class EquipmentTypeEdge:
                @dataclass_json
                @dataclass
                class EquipmentType:
                    @dataclass_json
                    @dataclass
                    class PropertyType:
                        id: str
                        name: str
                        type: PropertyKind = enum_field(PropertyKind)
                        index: Optional[int] = None
                        stringValue: Optional[str] = None
                        intValue: Optional[int] = None
                        booleanValue: Optional[bool] = None
                        floatValue: Optional[float] = None
                        latitudeValue: Optional[float] = None
                        longitudeValue: Optional[float] = None
                        isEditable: Optional[bool] = None
                        isInstanceProperty: Optional[bool] = None

                    @dataclass_json
                    @dataclass
                    class EquipmentPositionDefinition:
                        id: str
                        name: str
                        index: Optional[int] = None
                        visibleLabel: Optional[str] = None

                    @dataclass_json
                    @dataclass
                    class EquipmentPortDefinition:
                        id: str
                        name: str
                        index: Optional[int] = None
                        visibleLabel: Optional[str] = None

                    id: str
                    name: str
                    propertyTypes: List[PropertyType]
                    positionDefinitions: List[EquipmentPositionDefinition]
                    portDefinitions: List[EquipmentPortDefinition]
                    category: Optional[str] = None

                node: Optional[EquipmentType] = None

            edges: Optional[List[EquipmentTypeEdge]] = None

        equipmentTypes: EquipmentTypeConnection

    data: Optional[EquipmentTypesQueryData] = None
    errors: Any = None

    @classmethod
    # fmt: off
    def execute(cls, client):
        # fmt: off
        variables = None
        response_text = client.call(cls.__QUERY__, variables=variables)
        return cls.from_json(response_text).data
