#!/usr/bin/env python3
########################################################################################################################
# Some Settings
########################################################################################################################
GOLANG_GENERICS = [
    "bool",
    "int",
    "float64",
    "interface",
    "interface{}",
    "string"
]

IMPORT_MAPPING = {
    "time.Time": "time",
    "*time.Time": "time"
}

TYPE_MAPPING = {
    # Generic Types
    "boolean": "*bool",
    "integer": "*int",
    "number": "*float64",
    "object": "*interface{}",# This is a shortcut and might give problems in the future.
    "string": "*string",
    # Specific Empty Definitions, used as helpers within the API
    "Io_k8s_apimachinery_pkg_apis_meta_v1_Time": "time.Time",
    "Io_k8s_apimachinery_pkg_util_intstr_IntOrString": "int"
}

URL = 'https://raw.githubusercontent.com/kubernetes/kubernetes/master/api/openapi-spec/swagger.json'
