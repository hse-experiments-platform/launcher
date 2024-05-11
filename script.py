import re
import os

def snake_case(s):
    return re.sub(r'(?<!^)(?=[A-Z])', '_', s).lower()

def extract_methods(interface_content):
    # Исправленное регулярное выражение для поиска методов
    method_pattern = r'(\w+)\(context\.Context, \*(\w+)\) \(\*(\w+), error\)'
    return re.findall(method_pattern, interface_content)

def write_method_file(method, file_dir='output', service_name='launcher'):
    method_name, request_type, response_type = method
    file_name = f"{snake_case(method_name)}.go"
    if not os.path.exists(file_dir):
        os.makedirs(file_dir)
    file_path = os.path.join(file_dir, file_name)
    with open(file_path, 'w') as f:
        f.write(f'''package {service_name}

import (
    "context"

    pb "github.com/hse-experiments-platform/{service_name}/pkg/{service_name}"
)

func (s *{service_name}Service) {method_name}(ctx context.Context, req *pb.{request_type}) (*pb.{response_type}, error) {{
    panic("unimplemented")
}}
''')


interface_content = '''
type LauncherServiceServer interface {
	GetLaunches(context.Context, *GetLaunchesRequest) (*GetLaunchesResponse, error)
	LaunchDatasetUpload(context.Context, *LaunchDatasetUploadRequest) (*LaunchDatasetUploadResponse, error)
	GetDatasetUploadLaunch(context.Context, *GetDatasetUploadLaunchRequest) (*GetDatasetUploadLaunchResponse, error)
	LaunchDatasetSetTypes(context.Context, *LaunchDatasetSetTypesRequest) (*LaunchDatasetSetTypesResponse, error)
	GetDatasetSetTypesLaunch(context.Context, *GetDatasetSetTypesLaunchRequest) (*GetDatasetSetTypesLaunchResponse, error)
	LaunchDatasetPreprocess(context.Context, *LaunchDatasetPreprocessRequest) (*LaunchDatasetPreprocessResponse, error)
	GetDatasetPreprocessLaunch(context.Context, *GetDatasetPreprocessLaunchRequest) (*GetDatasetPreprocessLaunchResponse, error)
	LaunchTrain(context.Context, *LaunchTrainRequest) (*LaunchTrainResponse, error)
	GetTrainLaunch(context.Context, *GetTrainLaunchRequest) (*GetTrainLaunchResponse, error)
	LaunchPredict(context.Context, *LaunchPredictRequest) (*LaunchPredictResponse, error)
	GetPredictLaunch(context.Context, *GetPredictLaunchRequest) (*GetPredictLaunchResponse, error)
	LaunchGenericConvert(context.Context, *LaunchGenericConvertRequest) (*LaunchGenericConvertResponse, error)
	GetGenericConvertLaunch(context.Context, *GetGenericConvertLaunchRequest) (*GetGenericConvertLaunchResponse, error)
}
'''

def main():
    methods = extract_methods(interface_content)
    print(methods)
    for method in methods:
        print(method)
        write_method_file(method)

main()