import logging

ch = logging.StreamHandler()
ch.setLevel(logging.INFO)
log = logging.getLogger()
log.setLevel(logging.DEBUG)
log.addHandler(ch)

SUPPORTED_DISTRO_LIST = ["ubuntu1804"]
SUPPORTED_PYVERSION_LIST = ["3.9.13", "3.11.7“]
SUPPORTED_GOLANG_LIST = ["1.20"]
#UPPORTED_CUDA_LIST = ["10.1", "10.2", "11.2", "11.6"]
SUPPORTED_CUDA_LIST = ["11.6","11.7","12.1"]

# ECR_REPO = "public.ecr.aws/iflytek-open"
ECR_REPO = "iflyopensource"
INNER_REPO = "artifacts.iflytek.com/docker-private/atp"
TEMP_GEN_DIR = "./dist/aiges"
Dockerfile = "Dockerfile"
