%global debug_package %{nil}
%global __strip /bin/true

Name:          kmesh
Version:       1.0.0
Release:       1
Summary:       %{name} is a eBPF-based service mesh kernel solution
License:       ASL 2.0 and GPL-2.0
URL:           https://github.com/kmesh-net
Source0:       %{name}-%{version}.tar.gz

%description
%{name} is a eBPF-based service mesh kernel solution.

ExclusiveArch: x86_64 aarch64

%package devel
Summary: Development files

%description devel
Development files

%prep
%autosetup -n %{name}-%{version}

%build
cd %{_builddir}/%{name}-%{version}
make build

%install
mkdir -p %{buildroot}%{_bindir}
install %{_builddir}/%{name}-%{version}/kmesh-daemon %{buildroot}%{_bindir}
install %{_builddir}/%{name}-%{version}/kmesh-cni %{buildroot}%{_bindir}

mkdir -p %{buildroot}/usr/lib64
install %{_builddir}/%{name}-%{version}/bpf/deserialization_to_bpf_map/libkmesh_deserial.so %{buildroot}/usr/lib64
install %{_builddir}/%{name}-%{version}/api/v2-c/libkmesh_api_v2_c.so %{buildroot}/usr/lib64

mkdir -p %{buildroot}/lib/modules/kmesh
install %{_builddir}/%{name}-%{version}/kernel/ko/kmesh.ko %{buildroot}/lib/modules/kmesh

mkdir -p %{buildroot}/usr/share/kmesh/bpf2go
cp -a %{_builddir}/%{name}-%{version}/bpf/kmesh/bpf2go/. %{buildroot}/usr/share/kmesh/bpf2go/
find %{buildroot}/usr/share/kmesh/bpf2go -type f -name '*.go' -delete


%check
cd %{_builddir}/%{name}-%{version}

%post
echo "installing ..."
ln -sf /lib/modules/kmesh/kmesh.ko /lib/modules/`uname -r`
depmod -a

%preun
if [ "$1" == "1" ]; then
    systemctl status kmesh | grep "active (running)"
    if [ "$?" == "0" ]; then
        systemctl restart kmesh.service
    fi
else
    systemctl stop kmesh.service
fi

%postun
if [ "$1" -ne "1" ]; then
    rm -rf /lib/modules/`uname -r`/kmesh.ko
fi
depmod -a

%clean
rm -rf %{buildroot}

%files
%defattr(-,root,root)
%attr(0500,root,root) %{_bindir}/kmesh-daemon
%attr(0500,root,root) %{_bindir}/kmesh-cni

%files devel
%attr(0500,root,root) %dir /lib/modules/kmesh
%attr(0400,root,root) /lib/modules/kmesh/kmesh.ko

%attr(0500,root,root) /usr/lib64/libkmesh_deserial.so
%attr(0500,root,root) /usr/lib64/libkmesh_api_v2_c.so

%attr(0500,root,root) %dir /usr/share/kmesh
%attr(0500,root,root) /usr/share/kmesh

%changelog
* Wed Sep 27 2023 kwb0523<kwb0523@163.com> - 0.0.1-1
- init package
