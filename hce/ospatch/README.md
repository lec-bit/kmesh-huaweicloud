After compiling the kernel following [Enhance the kernel](https://github.com/kmesh-net/kmesh/blob/main/docs/kmesh_kernel_compile.md), perform the following steps to replace the kernel:

**Copy the three RPM packages to the corresponding device and install them:**

```bash

rpm -ivh kernel-5.10.0_60.18.0.50.r1083_58.hce2.x86_64-1.x86_64.rpm --force
rpm -ivh kernel-devel-5.10.0_60.18.0.50.r1083_58.hce2.x86_64-1.x86_64.rpm
rpm -ivh kernel-headers-5.10.0_60.18.0.50.r1083_58.hce2.x86_64-1.x86_64.rpm

```

Execute the following commands to check how many kernels are present in the system. If you cannot see any, you need to reinstall the kernel package.

```
(x86)cat /boot/grub2/grub.cfg | grep menuentry
```

```
(arm)cat /boot/efi/EFI/euleros/grub.cfg |grep menuentry
```

```
Reinstall the kernel package if necessary:
rpm -e kernel
yum localinstall -y kernel-5.10.0_60.18.0.50.r1083_58.hce2.x86_64-1.x86_64.rpm
```

**Configure to boot from the default kernel. Replace the kernel name with the actual name found in the system:**

```
grub2-set-default 'Huawei Cloud EulerOS (5.10.0-60.18.0.50.r1083_58.hce2.x86_64) 2.0 (x86_64)'
```

**Run the following command to confirm the configuration was successful:**

```
grub2-editenv list


[root@localhost ~]# grub2-editenv list
saved_entry=Huawei Cloud EulerOS (5.10.0-60.18.0.50.r1083_58.hce2.x86_64) 2.0 (x86_64)
boot_success=0

```

**Reboot the machine.**