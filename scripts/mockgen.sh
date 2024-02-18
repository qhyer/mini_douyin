#!/bin/bash

# 设置app目录的路径，您可以根据需要修改这个路径
app_dir="../app"

# 遍历app目录下所有子文件夹
find "$app_dir" -type d -name 'biz' | while read -r biz_dir; do
    # 进入到biz目录下
    cd "$biz_dir" || exit
    # 查找当前biz目录下所有文件，排除biz.go和以_test.go结尾的文件
    find . -maxdepth 1 -type f ! -name 'biz.go' ! -name '*_test.go' | while read -r file; do
        # 移除文件名开头的"./"
        trimmed_file=${file#./}
        # 创建mock目录，如果不存在的话
        mkdir -p mock
        # 执行mockgen命令
        mockgen -source="$trimmed_file" -destination="mock/$trimmed_file"
        echo "mockgen -source=$trimmed_file -destination=mock/$trimmed_file 完成"
    done
    # 返回到脚本开始执行的目录
    cd - || exit
done