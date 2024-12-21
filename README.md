该代码仓用于提供Kmesh for HuaweiCloud项目部分所需文件，为kmesh仓部分编译产物存放仓

采用[Kmesh](https://github.com/kmesh-net/kmesh)仓最新release，由于kmesh数据面部分编译产物和操作系统版本强耦合，在此基础上提供如下能力：

1. 归档Kmesh在不同操作系统下编译出的ebpf程序、ko文件、so文件并打包成kmesh-devel.rpm包；
2. 归档Kmesh在不同操作系统下编译出的ebpf程序对应bpf2go文件于代码仓中；
3. 添加编译相关脚本文件，支持从kmesh仓拉取代码并生成对应的编译产物并打包至kmesh-devel



### 目录结构说明：

hce：提供给hce 2.0操作系统使用的文件

​		bpf2go:提供给hce 2.0操作系统使用的bpf2go文件

​		ospatch:[增强该操作系统](https://github.com/kmesh-net/kmesh/blob/main/docs/kmesh_kernel_compile.md)所需要的内核patch

openeuler：提供给openeuler2303操作系统使用的文件

ubuntu：提供给ubuntu 22.04 LTS操作系统使用的文件

build_tools：编译工具，用于fork kmesh代码并制作kmesh-devel包



### 用户使用指南：

例：在openeuler 2303环境中

1. 根据自己操作系统环境选择并安装kmesh-devel.rpm
2. 将ebpf程序拷贝到bpf2go文件相同目录下，确保其他ko、so文件位置正确
3. 根据自己操作系统环境修改import路径，"github.com/kmesh-for-huaweicloud/kmesh-huaweicloud/openeuler/bpf2go"
4. 执行编译