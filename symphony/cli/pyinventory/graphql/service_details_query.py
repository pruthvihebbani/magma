#!/usr/bin/env python3
# @generated AUTOGENERATED file. Do not Change!

from dataclasses import dataclass
from datetime import datetime
from gql.gql.datetime_utils import DATETIME_FIELD
from gql.gql.graphql_client import GraphqlClient
from gql.gql.client import OperationException
from gql.gql.reporter import FailedOperationException
from functools import partial
from numbers import Number
from typing import Any, Callable, List, Mapping, Optional
from time import perf_counter
from dataclasses_json import DataClassJsonMixin

from .customer_fragment import CustomerFragment, QUERY as CustomerFragmentQuery
from .property_fragment import PropertyFragment, QUERY as PropertyFragmentQuery

QUERY: List[str] = CustomerFragmentQuery + PropertyFragmentQuery + ["""
query ServiceDetailsQuery($id: ID!) {
  service: node(id: $id) {
    ... on Service {
      id
      name
      externalId
      customer {
        ...CustomerFragment
      }
      endpoints {
        id
        port {
          id
          properties {
            ...PropertyFragment
          }
          definition {
            id
            name
          }
          link {
            id
            properties {
              ...PropertyFragment
            }
            services {
              id
            }
          }
        }
        definition {
          role
        }
      }
      links {
        id
        properties {
          ...PropertyFragment
        }
        services {
          id
        }
      }
    }
  }
}

"""]

@dataclass
class ServiceDetailsQuery(DataClassJsonMixin):
    @dataclass
    class ServiceDetailsQueryData(DataClassJsonMixin):
        @dataclass
        class Node(DataClassJsonMixin):
            @dataclass
            class Customer(CustomerFragment):
                pass

            @dataclass
            class ServiceEndpoint(DataClassJsonMixin):
                @dataclass
                class EquipmentPort(DataClassJsonMixin):
                    @dataclass
                    class Property(PropertyFragment):
                        pass

                    @dataclass
                    class EquipmentPortDefinition(DataClassJsonMixin):
                        id: str
                        name: str

                    @dataclass
                    class Link(DataClassJsonMixin):
                        @dataclass
                        class Property(PropertyFragment):
                            pass

                        @dataclass
                        class Service(DataClassJsonMixin):
                            id: str

                        id: str
                        properties: List[Property]
                        services: List[Service]

                    id: str
                    properties: List[Property]
                    definition: EquipmentPortDefinition
                    link: Optional[Link]

                @dataclass
                class ServiceEndpointDefinition(DataClassJsonMixin):
                    role: Optional[str]

                id: str
                definition: ServiceEndpointDefinition
                port: Optional[EquipmentPort]

            @dataclass
            class Link(DataClassJsonMixin):
                @dataclass
                class Property(PropertyFragment):
                    pass

                @dataclass
                class Service(DataClassJsonMixin):
                    id: str

                id: str
                properties: List[Property]
                services: List[Service]

            id: str
            name: str
            endpoints: List[ServiceEndpoint]
            links: List[Link]
            externalId: Optional[str]
            customer: Optional[Customer]

        service: Optional[Node]

    data: ServiceDetailsQueryData

    @classmethod
    # fmt: off
    def execute(cls, client: GraphqlClient, id: str) -> Optional[ServiceDetailsQueryData.Node]:
        # fmt: off
        variables = {"id": id}
        try:
            start_time = perf_counter()
            response_text = client.call(''.join(set(QUERY)), variables=variables)
            res = cls.from_json(response_text).data
            elapsed_time = perf_counter() - start_time
            client.reporter.log_successful_operation("ServiceDetailsQuery", variables, elapsed_time)
            return res.service
        except OperationException as e:
            raise FailedOperationException(
                client.reporter,
                e.err_msg,
                e.err_id,
                "ServiceDetailsQuery",
                variables,
            )
