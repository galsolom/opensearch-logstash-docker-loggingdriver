import os
import sys
import json

project_folder='testproject'
dotnet_delimiter="|tc="
_tests=[]
# dir = r'E:\dev\opensearch-logstash-plugin\testproject'

def is_in_list(params,list):
    return any(param in params for param in list)
    
labels_json = open("labels.json").read()
input_filters = sys.argv # passed identity
input_filters.pop(0) # remove file name from argument list
filters = input_filters
_json=json.loads(labels_json)
for tc_object in _json:
    tc_id=str(tc_object)
    tc_json = _json[tc_id]
    if is_in_list(filters,tc_json["domains"]) or is_in_list(filters,tc_json["labels"]):
        _tests.append(tc_id)

tests_list_syntaxed_dotnet = dotnet_delimiter.join(_tests)
command = 'dotnet test --filter="tc='+tests_list_syntaxed_dotnet+'"'


print(command)
#change directory
os.chdir(project_folder)
test = os.popen(command,'r')
#results
print (test.read())