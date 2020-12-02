import requests
import json


########################################################################################################################
# Helpers
########################################################################################################################
# Readers
def get_url(url):
    response = requests.get(url)
    return json.loads(response.text)

def read_json():
    return json.loads(open("./raw.json").read())

# Parsers
def parse_ref(ref):
    return _capitalize(
        ref.replace("#/definitions/", "")
           .replace(".", "_")
           .replace("-", "__")
           .replace("$", "_")
    )

def _capitalize(string):
    if len(string) == 1:
        return string.capitalize()
    elif len(string) > 1:
        return string[0].capitalize() + string[1:]
    else:
        return string

def parse_required(definition):
    if "required" in definition.keys():
        return definition["required"]
    else:
        return []

def parse_spec(struct_name, spec):
    if "type" in spec.keys():
        if "items" in spec.keys() and spec["type"] == "array":
            if "type" in spec["items"].keys():
                result = "[]" + translate_spec(spec["items"]["type"])
            else:
                result = "[]" + translate_spec(parse_ref(spec["items"]["$ref"]))
        else:
            result = translate_spec(spec["type"])
    else:
        result = translate_spec(parse_ref(spec["$ref"]))
    
    # Handle recursive structures by formulating a pointer
    if struct_name == result:
        result = "*" + result
    return result

# Translators
def translate_import(import_):
    if import_ == "time.Time":
        return "time"

def translate_spec(spec):
    if spec == "integer":
        return "int"
    elif spec == "number":
        return "float64"
    elif spec == "object":
        return "interface{}"
    elif spec == "string":
        return "string"
    elif spec == "boolean":
        return "bool"
    elif spec == "Io_k8s_apimachinery_pkg_apis_meta_v1_Time":
        return "time.Time"
    elif spec == "Io_k8s_apimachinery_pkg_util_intstr_IntOrString":
        return "int"
    else:
        return spec

def formatToPointerIfNecessary(type_, required):
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


# Formatters
def set_space_after_column(properties, object_name):
    max_len = 0
    for property_ in properties:
        if len(property_[object_name]) > max_len:
            max_len = len(property_[object_name])

    return max_len

########################################################################################################################
# Toplevel
########################################################################################################################
# Create function to build Go code
def to_go_struct(struct_name, properties, imports, package_name):
    full_result = """package """ + package_name
    if len(imports) > 0:
        full_result = full_result + "\n\nimport ("
        # Filter out duplicates by converting the list to a set
        for import_ in set(imports):
            full_result = full_result + '\n\t"' + translate_import(import_) + '"'
        full_result = full_result + "\n)"
    
    full_result = full_result + "\n\ntype " + struct_name + " struct {"

    n_space_1 = set_space_after_column(properties, "property_name")
    n_space_2 = set_space_after_column(properties, "property_type")

    for property_ in properties:
        if property_["required"]:
            full_result += (
                "\n\t" + property_["property_name"] + " "*(n_space_1-len(property_["property_name"])) + " " + 
                property_["property_type"] + " "*(n_space_2-len(property_["property_type"])) + " " +
                '`json:"' + property_["raw_property_name"] + '"`'
            )
        else:
            full_result += (
                "\n\t" + property_["property_name"] + " "*(n_space_1-len(property_["property_name"])) + " " + 
                property_["property_type"] + " "*(n_space_2-len(property_["property_type"])) + " " +
                '`json:"' + property_["raw_property_name"] + ',omitempty"`'
            )

    full_result += "\n}\n\n"
    return full_result


def parse_definition(struct_name, definition, package_name):
    properties = []
    required = parse_required(definition)

    imports = []
    # StorageVersionSpec is an Empty spec
    if "properties" in definition:
        for property_, specs in definition["properties"].items():
            isRequired = property_ in required
            property_type = parse_spec(parse_ref(struct_name), specs)
            properties.append({
                "raw_property_name": property_, 
                "property_name": parse_ref(property_),
                "property_type": formatToPointerIfNecessary(property_type, isRequired),
                "required": isRequired
            })
            if property_type in ["time.Time"]:
                imports.append(property_type)
    return to_go_struct(parse_ref(struct_name), properties, imports, package_name)


########################################################################################################################
# Main
########################################################################################################################
if __name__ == "__main__":
    url = 'https://raw.githubusercontent.com/kubernetes/kubernetes/master/api/openapi-spec/swagger.json'
    jsonbody = get_url(url)
    #jsonbody = read_json()
    output_folder = "../model/"
    package_name = "kugo_model"
    total_files = 0

    for k in jsonbody["definitions"].keys():
        total_files += 1
        name = parse_ref(k)
        with open(output_folder + name + ".go", 'w') as out_file: 
            out_file.write(parse_definition(k, jsonbody["definitions"][k], package_name))

    print("Exported number of files: " + str(total_files))