#!/usr/bin/env python3
# @generated AUTOGENERATED file. Do not Change!

from dataclasses import dataclass, field
from datetime import datetime
from functools import partial
from gql.gql.datetime_utils import fromisoformat
from numbers import Number
from typing import Any, Callable, List, Mapping, Optional

from dataclasses_json import dataclass_json
from marshmallow import fields as marshmallow_fields

from .equipment_port_input import EquipmentPortInput
from .equipment_position_input import EquipmentPositionInput
from .property_type_input import PropertyTypeInput


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
class EditEquipmentTypeInput:
    id: str
    name: str
    positions: List[EquipmentPositionInput]
    ports: List[EquipmentPortInput]
    properties: List[PropertyTypeInput]
    category: Optional[str] = None

