After compiling the kernel following [Enhance the kernel](https://github.com/kmesh-net/kmesh/blob/main/docs/kmesh_kernel_compile.md), perform the following steps to replace the kernel:

**Copy the three RPM packages to the corresponding device and install them:**
for example, we use openeuler 22.03 SP4

```bash

rpm -ivh kernel-5.10.0-216.0.0.115.oe2203sp4.x86_64.rpm --force
rpm -ivh kernel-devel-5.10.0-216.0.0.115.oe2203sp4.x86_64.rpm
rpm -ivh kernel-headers-5.10.0-216.0.0.115.oe2203sp4.x86_64.rpm

```

Execute the following commands to check how many kernels are present in the system. 
Find the corresponding version of the kernel.
If you cannot see any, you need to reinstall the kernel package.

```bash
(x86)cat /boot/grub2/grub.cfg | grep menuentry
...
menuentry 'openEuler (5.10.0-216.0.0.115.oe2203sp4.x86_64) 22.03 (LTS-SP4)'
...

```

```bash
(arm)cat /boot/efi/EFI/euleros/grub.cfg |grep menuentry
```

```bash
Reinstall the kernel package if necessary:
# backup old grub.cfg
cp /boot/grub2/grub.cfg /boot/grub2/grub.cfg.bak
rpm -e kernel
yum localinstall -y kernel-5.10.0-216.0.0.115.oe2203sp4.x86_64.rpm
```

**Configure to boot from the default kernel. Replace the kernel name with the actual name found in the system:**

```
grub2-set-default 'openEuler (5.10.0-216.0.0.115.oe2203sp4.x86_64) 22.03 (LTS-SP4)'
```

**Run the following command to confirm the configuration was successful:**

```bash
grub2-editenv list

    saved_entry=openEuler (5.10.0-216.0.0.115.oe2203sp4.x86_64) 22.03 (LTS-SP4)
    boot_success=0

```

**Reboot the machine.**
```bash
reboot
```