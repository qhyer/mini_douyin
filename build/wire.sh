#!/bin/bash

# 遍历../app目录下的所有子目录
for dir in $(find ../app -type d)
do
  # 如果目录中存在wire.go，则执行wire命令
  if [[ -f "${dir}/wire.go" ]]; then
    echo "找到 wire.go 在 ${dir}"
    # 在包含wire.go文件的目录中运行wire命令
    cd "${dir}"
    wire
    cd -
  fi
done