#!/usr/bin/bash
BRANCH=${1:-main}
git clone https://github.com/kmesh-net/kmesh.git
git switch "$BRANCH"
if [ $? -eq 0 ]; then
    echo "switch to $BRANCH success"
else
    echo "switch to $BRANCH failed"
fi

CURRENT_PATH=$(pwd)
KMESH_PATH=$(pwd)/kmesh
echo "当前路径: $KMESH_PATH"

# modify Environment Variables HOME 
# mkdir rpmbuild
cd $CURRENT_PATH/
KMESH_VERSION=$(grep Version: kmesh.spec | egrep -o [0-9]+.[0-9]+.[0-9]+)

cd $CURRENT_PATH/
rm -rf rpmbuild
mkdir rpmbuild
cd rpmbuild
mkdir BUILD SOURCES BUILDROOT RPMS SPECS SRPMS   

# rpmbuild files prepare
cd $CURRENT_PATH/
cp -r $KMESH_PATH kmesh_tmp
mv kmesh_tmp $CURRENT_PATH/kmesh-$KMESH_VERSION
cp kmesh.spec $CURRENT_PATH/rpmbuild/SPECS/

cd $CURRENT_PATH
#(cd kmesh-$KMESH_VERSION/ && ./build.sh -c)
tar zcf kmesh-$KMESH_VERSION.tar.gz kmesh-$KMESH_VERSION/
rm -rf kmesh-$KMESH_VERSION/
mv kmesh-$KMESH_VERSION.tar.gz rpmbuild/SOURCES/

# rpm build 
cd $CURRENT_PATH/rpmbuild/SPECS/
rpmbuild --define="_topdir $CURRENT_PATH/rpmbuild/" -bb kmesh.spec

rpm -qlp $CURRENT_PATH/rpmbuild/RPMS/x86_64/kmesh-devel-0.0.1-1.x86_64.rpm


