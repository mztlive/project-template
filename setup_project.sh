#!/bin/bash

# Get user input for package name and project directory
read -p "请输入包名: " package_name
read -p "请输入项目目录名: " project_dir

# Clone the repository
echo "正在克隆项目模板..."
git clone --depth 1 git@github.com:mztlive/project-template.git "${project_dir}"
echo "克隆完成."

# Change to the project directory
cd "${project_dir}"

# Remove the .git directory
echo "删除 .git 目录..."
rm -rf .git
echo ".git 目录已删除."

# Find all files and replace the template package name with the user's input
echo "正在替换包名..."
find . -type f -exec sed -i "s|github.com/mztlive/project-template|${package_name}|g" {} \;
echo "包名替换完成."

# Done
echo "脚本执行完毕，项目已设置为用户指定的包名."