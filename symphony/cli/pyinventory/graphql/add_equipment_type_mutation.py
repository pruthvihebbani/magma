#!/usr/bin/env python3
# @generated AUTOGENERATED file. Do not Change!

from dataclasses import dataclass, field
from datetime import datetime
from enum import Enum
from functools import partial
from numbers import Number
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
    service = "service"
    datetime_local = "datetime_local"
    MISSING_ENUM = ""

    @classmethod
    def _missing_(cls, value):
        return cls.MISSING_ENUM


@dataclass_json
@dataclass
class AddEquipmentTypeInput:
    @dataclass_json
    @dataclass
    class EquipmentPositionInput:
        name: str
        id: Optional[str] = None
        index: Optional[int] = None
        visibleLabel: Optional[str] = None

    @dataclass_json
    @dataclass
    class EquipmentPortInput:
        name: str
        id: Optional[str] = None
        index: Optional[int] = None
        visibleLabel: Optional[str] = None
        portTypeID: Optional[str] = None
        bandwidth: Optional[str] = None

    @dataclass_json
    @dataclass
    class PropertyTypeInput:
        name: str
        type: PropertyKind = enum_field(PropertyKind)
        id: Optional[str] = None
        index: Optional[int] = None
        category: Optional[str] = None
        stringValue: Optional[str] = None
        intValue: Optional[int] = None
        booleanValue: Optional[bool] = None
        floatValue: Optional[Number] = None
        latitudeValue: Optional[Number] = None
        longitudeValue: Optional[Number] = None
        rangeFromValue: Optional[Number] = None
        rangeToValue: Optional[Number] = None
        isEditable: Optional[bool] = None
        isInstanceProperty: Optional[bool] = None
        isMandatory: Optional[bool] = None
        isDeleted: Optional[bool] = None

    name: str
    positions: List[EquipmentPositionInput]
    ports: List[EquipmentPortInput]
    properties: List[PropertyTypeInput]
    category: Optional[str] = None


@dataclass_json
@dataclass
class AddEquipmentTypeMutation:
    __QUERY__ = """
    mutation AddEquipmentTypeMutation($input: AddEquipmentTypeInput!) {
  addEquipmentType(input: $input) {
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
    class AddEquipmentTypeMutationData:
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

        addEquipmentType: Optional[EquipmentType] = None

    data: Optional[AddEquipmentTypeMutationData] = None
    errors: Optional[Any] = None

    @classmethod
    # fmt: off
    def execute(cls, client, input: AddEquipmentTypeInput):
        # fmt: off
        variables = {"input": input}
        response_text = client.call(cls.__QUERY__, variables=variables)
        return cls.from_json(response_text).data
