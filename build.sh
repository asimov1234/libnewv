#!/bin/zsh
set -x
set -e

#python3 build/main.py apple go
cd *.xcframework/ios-arm64 && mkdir temp && mv *.a temp/ 
cd temp
# 保存.a文件的文件名
A_FILE_NAME=$(ls *.a)
echo "保存的文件名: $A_FILE_NAME"
# 解包.a文件
ar -x *.a
# 替换go.o中的Xray和xray为xxxx
if [ -f go.o ]; then
    echo "替换go.o中的Xray和xray为xxxx"
    echo "查找并替换Xray为xxxx:"
    # 显示替换前的匹配内容
    hexdump -C go.o | grep -i "xray"
    # 使用perl处理二进制文件
    perl -pi -e 's/Xray/xxxx/g' go.o
    echo "查找并替换xray为xxxx:"
    # 显示替换后的匹配内容
    hexdump -C go.o | grep -i "xray"
    perl -pi -e 's/xray/xxxx/g' go.o
    echo "替换完成，最终检查:"
    hexdump -C go.o | grep -i "xxxx"
fi
# 重新打包成原来的文件名
ar crs "$A_FILE_NAME" *.o
# 将重新打包的.a文件移动到temp目录下
mv "$A_FILE_NAME" ../
# 返回上级目录并删除temp目录
cd ..
rm -rf temp
