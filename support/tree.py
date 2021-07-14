#!/usr/bin/env python3
"""

"""

from .settings import (
    TYPE_MAPPING
)

def assess_parents_children(definitions):
    children = [{k: []} for k in definitions.keys()]
    parents = [{k: []} for k in definitions.keys()]

    for k in definitions.keys():
        if "properties" in definitions[k].keys():
            for property_ in definitions[k]["properties"]:

def isEdge(child):
    if child["type"] not in TYPE_MAPPING.keys():
        return False
    else:
        return True
