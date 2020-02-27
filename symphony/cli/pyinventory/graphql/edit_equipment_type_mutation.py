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

from .property_kind_enum import PropertyKind

from .edit_equipment_type_input import EditEquipmentTypeInput


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



@dataclass_json
@dataclass
class EditEquipmentTypeMutation:
    __QUERY__ = """
    mutation EditEquipmentTypeMutation($input: EditEquipmentTypeInput!) {
  editEquipmentType(input: $input) {
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

    """

    @dataclass_json
    @dataclass
    class EditEquipmentTypeMutationData:
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
                floatValue: Optional[Number] = None
                latitudeValue: Optional[Number] = None
                longitudeValue: Optional[Number] = None
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

        editEquipmentType: Optional[EquipmentType] = None

    data: Optional[EditEquipmentTypeMutationData] = None
    errors: Optional[Any] = None

    @classmethod
    # fmt: off
    def execute(cls, client: GraphqlClient, input: EditEquipmentTypeInput):
        # fmt: off
        variables = {"input": input}
        response_text = client.call(cls.__QUERY__, variables=variables)
        return cls.from_json(response_text).data
