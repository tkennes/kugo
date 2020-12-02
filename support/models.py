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


########################################################################################################################
# Some Settings
########################################################################################################################
NO_DESCRIPTION_REQUIRED_IF_MISSING = [
    "bool",
    "int",
    "float64",
    "interface",
    "string"
]

IMPORT_MAPPING = {
    "time.Time": "time"
}

TYPE_MAPPING = {
    # Generic Types
    "boolean": "bool",
    "integer": "int",
    "number": "float64",
    "object": "interface",# This is a shortcut and might give problems in the future.
    "string": "string",
    # Specific Empty Definitions, used as helpers within the API
    "Io_k8s_apimachinery_pkg_apis_meta_v1_Time": "time.Time",
    "Io_k8s_apimachinery_pkg_util_intstr_IntOrString": "int"
}

URL = 'https://raw.githubusercontent.com/kubernetes/kubernetes/master/api/openapi-spec/swagger.json'
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

def check_and_add_import(imports, property_type):
    if property_type in IMPORT_MAPPING.keys():
        imports.append(translate_import(property_type))
    return imports

###################### Formatter ######################
def formatToPointerIfNecessary(type_, required):
    """ Format the field type into a pointer if necessary.

    Why do we want this? 
        - If a field is not required, and not available, using a pointer
        would result in a "nil" value. Otherwise, we need to dive deeper
        into the tree to investigate whether or not this branch exists.

    When is this necessary? If and only if:
        - The field is not required
        - The field does not start with []
        - The field is not of a generic Go Type
            - string
            - int
            - float64
            - bool

    Args:
        type_ (str): Go variable
        required (bool): Whether or not the variable is a required variable

    Returns:
        string: A pointer to the Go variable
    """
    # It should be a point if it is not a basic go object, and not required/omittable
    if not required:
        if len(type_) > 2:
            if type_[0:2] != "[]" and type_ not in ["string", "int", "float64", "bool"]:
                return "*" + type_
            else:
                return type_
        else:
            return type_
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
    if not description and struct_name not in NO_DESCRIPTION_REQUIRED_IF_MISSING:
        return "\n\t// See: " + linkify_struct_name(struct_name)
    elif not description and struct_name in NO_DESCRIPTION_REQUIRED_IF_MISSING:
        return "\n\t// NO DESCRIPTION"
    return "\n\t// " + "\n\t// ".join(textwrap.wrap(description, 120, break_long_words=False, break_on_hyphens=True))

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
def create_go_struct(struct_description, struct_name, properties, imports, package_name, output_format="json", verbose=True):
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

    # Finally, close of the struct and add a newline to comply with coding standards.
    full_result += "\n}\n"

    return full_result


def parse_definition(struct_name, definition, package_name, output_format="json", verbose=True):
    """Parse the raw definition, clean up the names, definitions and properties and pass the results on to have a Go script created

    Args:
        struct_name (str): The name we want to have for the struct
        definition (str): The definition that accompanies the struct
        package_name (str): The name we have chosen for the package
        output_format (str, optional): The format we want to request from the API. Defaults to "json".
        verbose (bool, optional): Whether or not we want to add descriptions in general. Defaults to True.

    Returns:
        String: The contents of a Go file
    """
    # Some input validation:
    assert output_format in ["json", "yaml"]
    
    # Container variables
    properties = []
    imports = []

    # Assess what properties are required
    required = parse_required(definition)

    # Assess and parse the Struct Definition
    if "description" in definition.keys():
        full_description = definition["description"]
    else:
        full_description = ""

    # Assess and parse the Struct Properties
    # StorageVersionSpec is an Empty spec: Hence, we need to check for this.
    if "properties" in definition:
        for property_, specs in definition["properties"].items():
            isRequired = property_ in required

            # Assemble property Type
            property_type = parse_type(clean_ref(struct_name), specs)
            property_type = formatToPointerIfNecessary(property_type, isRequired)

            # Assemble Property Description
            if "description" in definition["properties"][property_]:
                property_description = definition["properties"][property_]["description"]
            else:
                property_description = ""
            
            # Assemble Property
            properties.append({
                "raw_property_name": property_, 
                "property_name": clean_ref(property_),
                "property_type": formatToPointerIfNecessary(property_type, isRequired),
                "required": isRequired,
                "description": property_description
            })

            # Assess whether the property type requires an import
            imports = check_and_add_import(imports, property_type)
    
    # Pass the parsed variables to the Go Struct Generation
    return create_go_struct(full_description, clean_ref(struct_name), properties, imports, package_name, output_format, verbose)


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

    # Constants
    package_name = "kugo_model"

    # Create the folder if it does not exist
    if not os.path.exists(args.folder):
        os.makedirs(args.folder)

    # Loop over the definitions and, clean the names and output them to the required folder
    total_files = 0
    for k in jsonbody["definitions"].keys():
        total_files += 1
        name = clean_ref(k)
        with open(args.folder + "/" + name + ".go", 'w') as out_file: 
            out_file.write(parse_definition(k, jsonbody["definitions"][k], args.package, output_format=args.output, verbose=not args.no_descriptions))

    # Conclude Reporting
    print("----------------------------------------------------------------------")
    print("Exported number of files: " + str(total_files))
    print("----------------------------------------------------------------------")

########################################################################################################################
# EOF
########################################################################################################################