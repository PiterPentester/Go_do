#!/bin/bash
### Get app name & image url ###

echo "go-do" > properties
echo "${REGISTRY_URL}/${REGISTRY_NAMESPACE}/${IMAGE_NAME}:${BUILD_NUMBER}" >> properties

input="properties"
N=0
while IFS= read -r line
do
  if [ $N == 0 ];then
    echo "IDS_PROJECT_NAME=$line"
    export "IDS_PROJECT_NAME=$line"
    N=1
  elif [ $N == 1 ]
  then
    echo "PIPELINE_IMAGE_URL=$line"
    export "PIPELINE_IMAGE_URL=$line"
  fi
done < "$input"
