#!/usr/bin/env python3
"""
This script is used to generate Go scripts from the Kubernetes API endpoints. There is a 10.000 lines long swagger
document in JSON format that is used to populate the data at the website:
- https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.19/

This json file is available at:
- https://raw.githubusercontent.com/kubernetes/kubernetes/master/api/openapi-spec/swagger.json

E.g. too much to do by hand and we need to introduce some sort of automation in order for a strongly typed
language to still benefit from the rich environment of Kubernetes.
"""
import argparse
import os
import json
import requests
import textwrap

from settings import (
    GOLANG_GENERICS,
    IMPORT_MAPPING,
    TYPE_MAPPING,
    URL
)


EMPTY_DEFINITIONS = []
MANDATORY_POINTERS = []

########################################################################################################################
# Helpers
########################################################################################################################
###################### Readers ######################
def get_url(url):
    """ Get the URL page and dump it back in json format

    Args:
        url (str): the URL, ideally a FQDN

    Returns:
        json: A json-string containin the results
    """
    response = requests.get(url)
    return json.loads(response.text)

def read_json():
    """ Read from a json file

    Returns:
        json: A json-string containin the results
    """
    return json.loads(open("./raw.json").read())

###################### Parsers ######################
def clean_ref(ref):
    """ Clean $ref fields such that they can be used in Go.
    
    These files typically refer to other object definitions.
    Example:
        $/definitions/this-is.a.field$ref -> this__is_a_field_ref

    Args:
        ref (str): The raw $ref value as in the json

    Returns:
        string: The cleaned field.
    """
    return _capitalize(
        ref.replace("#/definitions/", "")
           .replace(".", "_")
           .replace("-", "__")
           .replace("$", "_")
    )

def _capitalize(string):
    """ Custom capitalize, only upper-casing the first character.

    Example:
        teST -> TeST

    Args:
        string (str): Input string

    Returns:
        string: "_capital-ized string.
    """
    if len(string) == 1:
        return string.capitalize()
    elif len(string) > 1:
        return string[0].capitalize() + string[1:]
    else:
        return string

def parse_required(definition):
    """ Get the required from the definition if they exist.

    Args:
        definition (Definition): An object definition

    Returns:
        list: A list containing the required object properties
    """
    if "required" in definition.keys():
        return definition["required"]
    else:
        return []

def parse_type(struct_name, spec):
    """ Parse type such that it can be used by Go

    Args:
        struct_name (str): The name of the resulting struct
        spec (Spec): The raw input spec, a field form the raw json body

    Returns:
        string: A cleaned type
    """
    if "type" in spec.keys():
        if "items" in spec.keys() and spec["type"] == "array":
            # When an items key exists and the type is array, Go would call it a slice
            if "type" in spec["items"].keys():
                result = "[]" + translate_spec(spec["items"]["type"])
            else:
                result = "[]" + translate_spec(clean_ref(spec["items"]["$ref"]))
        else:
            result = translate_spec(spec["type"])
    else:
        result = translate_spec(clean_ref(spec["$ref"]))
    
    # Handle recursive structures by formulating a pointer
    # You get errors if you have a struct like:
    # type struct_a struct {
    #   field_a struct_a `json:"field_a"`
    # }
    if struct_name == result:
        print("RECURSIVE: " + result)
        result = "*" + result
    return result

###################### Translators ######################
def translate_spec(spec):
    """ Perform a translation between the Kubernetes API types and Go types.

    Args:
        spec (str): Kubernetes API type

    Returns:
        string: Go type
    """
    if spec in TYPE_MAPPING.keys():
        return TYPE_MAPPING[spec]
    else:
        return spec

def translate_import(import_):
    """ A function to translate imports. 
    
    Some variable types might require imports, most notably time.Time.
    These can then be added here.

    Args:
        import_ (str): Raw Import, maybe a Go Object of which we need to deduce a package

    Returns:
        string: Indicating a Go Package
    """
    if import_ in IMPORT_MAPPING.keys():
        return IMPORT_MAPPING[import_]
    else:
        return import_

###################### Validators ######################
def check_and_add_import(imports, property_type):
    if property_type in IMPORT_MAPPING.keys():
        imports.append(translate_import(property_type))
    return imports

def CheckLeaf(child_type):
    if child_type not in GOLANG_GENERICS:
        return False
    else:
        return True

def CheckForDefinitionWithoutProperties(definition):
    if "properties" not in definition.keys():
        return True
    else:
        return False

###################### Formatter ######################
def formatToPointerIfNecessary(type_, isLeaf, isRequired):
    """ Format the field type into a pointer if necessary.

    Sometimes we want to want to make use of a pointer to be able to check for nil.

    Args:
        type_ (str): Go variable

    Returns:
        string: A pointer to the Go variable
    """
    # It should be a point if it is not a basic go object, and not required/omittable
    if ((type_ in MANDATORY_POINTERS or isLeaf or not isRequired or type_ in EMPTY_DEFINITIONS) and 
       (type_[0] != "*" and type_[0] != "[")):
        if type_ in MANDATORY_POINTERS:
           print("MANDATORY_POINTER: " + type_)
        elif isLeaf:
            print("ISLEAF: " + type_)
        elif not isRequired:
            print("NOT ISREQUIRED: " + type_)
        else:
            print("EMPTY_DEF: " + type_)
        return "*" + type_
    else:
        return type_

###################### Prettifiers ######################
def set_space_after_column(properties, object_name):
    """ Used to determine the spacing between fieldname-fieldtype and fieldtype-jsonfield/yamlfield.

    Args:
        properties ([type]): [description]
        object_name ([type]): [description]

    Returns:
        [type]: [description]
    """
    max_len = 0
    for property_ in properties:
        if len(property_[object_name]) > max_len:
            max_len = len(property_[object_name])

    return max_len

def set_newlines_description_struct(description):
    """Used to format the description of the struct such that it does not become one long line

    Args:
        description (str): The raw one line-description

    Returns:
        string: Line-out struct description
    """
    if not description:
        return ""
    return "\n\n\n// " + "\n// ".join(textwrap.wrap(description, 120, break_long_words=False, break_on_hyphens=True))

def set_newlines_description_properties(description, struct_name):
    """Used to format the description of the property such that it does not become one long line

    Args:
        description (str): The raw description
        struct_name (str): The name of the struct

    Returns:
        string: Line out properties description
    """
    if not description and struct_name not in GOLANG_GENERICS:
        return "\n\t// See: " + linkify_struct_name(struct_name)
    elif not description and struct_name in GOLANG_GENERICS:
        return "\n\t// NO DESCRIPTION"
    else:
        res = "\n\t// " + "\n\t// ".join(textwrap.wrap(description, 120, break_long_words=False, break_on_hyphens=True))
        res += "\n\t// See: " + linkify_struct_name(struct_name)
        return res

def linkify_struct_name(struct_name):
    """Used to format a struct name such that it can be used to 
    jump directly to its struct.

    Args:
        struct_name (str): The raw struct name.

    Returns:
        string: A reference to the file containing the struct corresponding to the definition.
    """
    parsed_name = struct_name.replace("*", "").replace("[]", "")
    return "file://" + os.path.dirname(os.path.dirname(os.path.abspath(__file__))) + "/model/" + parsed_name + ".go"

########################################################################################################################
# Toplevel
########################################################################################################################
# Create function to build Go code
def create_go_struct(
    struct_description, struct_name, properties, imports, package_name, 
    parents, depth, output_format="json", verbose=True):
    """ Create a Go-struct from the parsed struct name and properties

    Args:
        struct_description (str): The description to be used for the struct
        struct_name (str): The name to be used for the struct
        properties (dict): A dictionary containing the prepared properties
        imports (list): A list of strings containing the raw imports
        package_name (str): The name of the resulting Go package
        output_format (str, optional): The output-format, either "json" or "yaml". Defaults to "json".
        verbose (bool, optional): Whether or not description should be added to the files. Defaults to True.

    Returns:
        string: The contents of a Go file
    """
    # Some input validation:
    assert output_format in ["json", "yaml"]
    
    # Start the file with the package name
    full_result = """package """ + package_name

    # Then add imports if necessary
    if len(imports) > 0:
        full_result += "\n\nimport ("
        # Filter out duplicates by converting the list to a set
        for import_ in set(imports):
            full_result += '\n\t"' + import_ + '"'
        full_result += "\n)"
    
    if depth > 0 and verbose:
        full_result += "\n\n"
        full_result += "\n// Tree Depth: " + str(depth)
        full_result += "\n// REFERENCES:"
        for parent in parents:
            full_result += "\n// - " + linkify_struct_name(parent)

    # If verbose is true, add the struct description if verbose
    if verbose:
        full_result += set_newlines_description_struct(struct_description)
    else:
        full_result += "\n"

    # Start of the struct definition
    full_result += "\ntype " + struct_name + " struct {"

    # Determine the spacing to be used to space out the columns
    n_space_1 = set_space_after_column(properties, "property_name")
    n_space_2 = set_space_after_column(properties, "property_type")

    for i, property_ in enumerate(properties):
        # Only add an extra newline if:
        #   - This is not the first property
        #   - Nor the last property 
        #   - We expect to write comments with descriptions as well (e.g. verbose) 
        if i != 0 and i != len(properties) and verbose:
            full_result += "\n"

        # If verbose is true, add descriptions, else keep it simple
        if verbose:
            full_result += set_newlines_description_properties(property_["description"], property_["property_type"])
        
        # If property is required, there is not much we need to do, otherwise, we need to stipulate
        # that it can be omitted from the struct.
        if property_["required"]:
            full_result += (
                "\n\t" + property_["property_name"] + " "*(n_space_1-len(property_["property_name"])) + " " + 
                property_["property_type"] + " "*(n_space_2-len(property_["property_type"])) + " " +
                '`' + output_format + ':"' + property_["raw_property_name"] + '"`'
            )
        else:
            full_result += (
                "\n\t" + property_["property_name"] + " "*(n_space_1-len(property_["property_name"])) + " " + 
                property_["property_type"] + " "*(n_space_2-len(property_["property_type"])) + " " +
                '`' + output_format + ':"' + property_["raw_property_name"] + ',omitempty"`'
            )

    # Finally, close of the struct
    full_result += "\n}\n"

    return full_result


def parse_definition(struct_name, definition, tree):
    """Parse the raw definition, clean up the names, definitions and properties and pass the results on to have a Go script created

    Args:
        struct_name (str): The name we want to have for the struct
        definition (str): The definition that accompanies the struct
        tree 

    Returns:
        String: The contents of a Go file
    """  
    # Container variables
    properties = []
    imports = []
    clean_struct_name = clean_ref(struct_name)
    tree["children"][clean_struct_name] = []
    
    # Assess what properties are required
    required = parse_required(definition)

    # If no field is required, we could end up with an
    # empty struct, which we should avoid by making a
    # pointer out of it
    # Leaving this out now, since we only want leaves to 
    # be structs
    # if len(required) == 0:
    #    MANDATORY_POINTERS.append(clean_struct_name)

    # Assess and parse the Struct Definition
    if "description" in definition.keys():
        full_description = definition["description"]
    else:
        full_description = ""

    # Assess and parse the Struct Properties
    # StorageVersionSpec is an Empty spec: Hence, we need to check for this.
    if not CheckForDefinitionWithoutProperties(definition):
        for property_, specs in definition["properties"].items():   
            # Assemble property Name
            property_name = clean_ref(property_)

            # Check if is required
            isRequired = property_ in required

            # Assemble property Type
            property_type = parse_type(clean_struct_name, specs)
            isLeaf = CheckLeaf(property_type)
            property_type = formatToPointerIfNecessary(property_type, isLeaf, isRequired)

            # Assemble Property Description
            if "description" in definition["properties"][property_]:
                property_description = definition["properties"][property_]["description"]
            else:
                property_description = ""
            
            # TODO: Fix
            if property_name.lower() in ["httpget"]:
                # This gives errors later since this port receives both ints and strings.
                # Hence, we either need to apply some custom maatwerk or skip for now.
                continue

            # Assemble Property
            properties.append({
                "raw_property_name": property_, 
                "property_name": property_name,
                "property_type": property_type,
                "required": isRequired,
                "description": property_description,
                "isLeaf": isLeaf
            })

            # Add child
            tree["children"][clean_struct_name].append({
                "name": property_name,
                "type": property_type,
                "isLeaf": isLeaf
            })

            # Assess whether the property type requires an import
            imports = check_and_add_import(imports, property_type)
    else:
        MANDATORY_POINTERS.append(clean_struct_name)

    golang_struct = {
        "struct_description": full_description,
        "struct_name": clean_struct_name,
        "properties": properties,
        "imports": imports,
    }

    # Pass the parsed variables to the Go Struct Generation
    return golang_struct, tree


########################################################################################################################
# Children and Parent Assessments
########################################################################################################################
def get_parents(tree):
    for definition in tree["children"]:
        tree["parents"][definition] = find_parents(definition, tree)
    return tree

def generalize_type(type_):
    return type_.replace("*", "").replace("[]", "")

def find_parents(child_name, tree):
    res = []
    for definition in tree["children"]:
        if definition != child_name:
            for child in tree["children"][definition]:
                if generalize_type(child["type"]) == child_name:
                    res.append(definition)
    return list(set(res))

def assess_tree_depth(tree):
    remaining_depths = check_remaining_depths(tree)
    print("Total definitions: " + str(remaining_depths) + "\n")
    print("----------------------------------------------------------------------")

    for definition in tree["parents"]:
        if len(tree["parents"][definition]) == 0:
            tree["depths"][definition] = 0

    iteration = 0
    while remaining_depths > 0 and iteration < 50:
        print("Filling up depth. Iteration: " + str(iteration + 1))
        tree = update_tree_depth(tree, iteration)
        iteration += 1
        remaining_depths = check_remaining_depths(tree)
        print("Remaining definitions: " + str(remaining_depths) + "\n")

    return tree

def update_tree_depth(tree, last_depth):
    for definition in tree["parents"]:
        if not tree["depths"][definition] and tree["depths"][definition] != 0:
            for parent in tree["parents"][definition]:
                if tree["depths"][parent] == last_depth:
                    tree["depths"][definition] = last_depth + 1
    return tree

def check_remaining_depths(tree):
    non_filled = 0
    for definition in tree["parents"]:
        # Seems to be a python bug here: None == 0?
        if not tree["depths"][definition] and tree["depths"][definition] != 0:
            non_filled += 1
    
    return non_filled

########################################################################################################################
# Main: Since this is a python-script, we need to have this main-condition to kick-off execution
########################################################################################################################
def create_parser():
    """ Handle CLI arguments

    Returns:
        parser: Return a parser with specific arguments
    """    
    parser = argparse.ArgumentParser()
    # The output folder
    parser.add_argument(
        "-f", "--folder",
        help='The output folder, a relative path is expected. (default: "../model/")',
        default="../model/"
    )
    # Whether or not descriptions should be added
    parser.add_argument(
        "-n", "--no-descriptions",
        help="Whether or not the Go structs should be commented with descriptions (default: true)",
        action="store_true"
    )
    # The output format
    parser.add_argument(
        "-o", "--output",
        help='The output type, json and yaml supported (default: "json")',
        default="json"
    )
    # The name of the package
    parser.add_argument(
        "-p", "--package",
        help='The name of the Go-package that will be created (default: "kugo_model")',
        default="kugo_model"
    )
    # Whether or not the input json should be downloaded from a URL (Useful for updating!)
    parser.add_argument(
        "--url",
        help='Whether or not the input should be downloaded from source or from file (defaul: to be "raw.json")',
        action="store_true"
    )

    return parser


if __name__ == "__main__":
    # Get Arguments
    parser = create_parser()
    args = parser.parse_args()

    # Starting reporting:
    print("----------------------------------------------------------------------")
    if args.url:
        print("Read input from: " + URL)
    else:
        print("Read input from local file: raw.json")
    print("Adding descriptions:", str(not args.no_descriptions))
    print("Outputting the Golang struct definitions into folder: " + args.folder)
    print("The output type:", args.output)
    print("Start assemble of package: " + args.package)

    # Getting input
    if args.url:
        # In case you want to read directly from source
        jsonbody = get_url(URL)
    else:
        # Since the source is rate-limited, you might prefer to download it to a file and read it from there
        jsonbody = read_json()

    # Create the folder if it does not exist
    if not os.path.exists(args.folder):
        os.makedirs(args.folder)

    tree = {
        "children": {},
        "parents": {},
        "depths": {}
    }

    # Find Empty Definitions
    for definition in jsonbody["definitions"].keys():
        if CheckForDefinitionWithoutProperties(jsonbody["definitions"][definition]):
            EMPTY_DEFINITIONS.append(clean_ref(definition))

    # Assemble golang structs and tree
    golang_structs = []
    for definition in jsonbody["definitions"].keys():
        golang_struct, tree = parse_definition(definition, jsonbody["definitions"][definition], tree)
        golang_structs.append(golang_struct)
        tree["parents"][golang_struct["struct_name"]] = []
        tree["depths"][golang_struct["struct_name"]] = None

    tree = get_parents(tree)
    # Assemble tree depth
    print("----------------------------------------------------------------------")
    tree = assess_tree_depth(tree)

    if not args.no_descriptions:
        # Create the folder if it does not exist
        if not os.path.exists("./meta"):
            os.makedirs("./meta")

        open("./meta/children.txt", 'w+').write(str(tree["children"]))
        open("./meta/parents.txt", 'w+').write(str(tree["parents"]))
        open("./meta/depths.txt", 'w+').write(str(tree["depths"]))

    # Loop over the definitions and, clean the names and output them to the required folder
    total_files = 0
    for g in golang_structs:
        total_files += 1
        with open(args.folder + "/" + g["struct_name"] + ".go", 'w') as out_file: 
            out_file.write(
                create_go_struct(
                    g["struct_description"],
                    g["struct_name"],
                    g["properties"],
                    g["imports"],
                    parents=tree["parents"][g["struct_name"]],
                    depth=tree["depths"][g["struct_name"]],
                    package_name=args.package,
                    output_format=args.output, 
                    verbose=not args.no_descriptions
                )
            )

    # Conclude Reporting
    print("----------------------------------------------------------------------")
    print("Exported number of files: " + str(total_files))
    print("----------------------------------------------------------------------")

########################################################################################################################
# EOF
########################################################################################################################
