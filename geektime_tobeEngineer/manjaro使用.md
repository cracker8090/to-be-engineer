# 1 使用技巧

## 1.1 win10和manjaro双系统时差问题
系统时间又因为系统的不同使用了两种时间管理办法：

    localtime：本地时间，目前只有 Windows 在使用。
    UTC：是一种世界标准时间，Linux 这类类 UNIX 多数会使用，UTC 加减时区之后才是本地时间。

系统时间又因为系统的不同使用了两种时间管理办法：

    localtime：本地时间，目前只有 Windows 在使用。
    UTC：是一种世界标准时间，Linux 这类类 UNIX 多数会使用，UTC 加减时区之后才是本地时间。

```
# 以管理员身份使用运行
reg add "HKEY_LOCAL_MACHINE\System\CurrentControlSet\Control\TimeZoneInformation" /v RealTimeIsUniversal /d 1 /t REG_DWORD /f

# 以上方法无效或64位系统：
reg add "HKEY_LOCAL_MACHINE\System\CurrentControlSet\Control\TimeZoneInformation" /v RealTimeIsUniversal /d 1 /t REG_QWORD /f

sudo timedatectl set-local-rtc true
```



## 1.2 如何管理启动应用程序 

查看系统负载

top，ps，uptime，atop，w





## 如何查看系统启动时间几怎么加快系统的启动

### 查看开机启动时间systemd-analyze 
Startup finished in 4.543s (kernel) + 15.609s (userspace) = 20.152s
graphical.target reached after 15.029s in userspace

查看开机启动项及启动时间systemd-analyze blame
          6.726s systemd-journal-flush.service
          5.382s lvm2-monitor.service
          4.388s dev-sdb2.device
          4.161s ModemManager.service
          2.582s ldconfig.service
          1.918s avahi-daemon.service
          1.874s NetworkManager.service
          1.874s systemd-logind.service
          1.869s bluetooth.service
          1.635s systemd-journal-catalog-update.service
           863ms systemd-tmpfiles-setup-dev.service
           767ms polkit.service
           707ms systemd-udev-trigger.service
           667ms upower.service
           646ms dev-hugepages.mount
           645ms dev-mqueue.mount
           616ms systemd-binfmt.service
           575ms tlp.service
           572ms udisks2.service
           502ms tmp.mount
           501ms sys-kernel-debug.mount
           413ms systemd-remount-fs.service
           400ms proc-sys-fs-binfmt_misc.mount
           371ms systemd-rfkill.service
           329ms systemd-timesyncd.service
           322ms systemd-sysusers.service
           299ms org.cups.cupsd.service
           242ms systemd-udevd.service
           216ms sys-kernel-config.mount
           214ms systemd-backlight@backlight:intel_backlight.service
           211ms systemd-tmpfiles-clean.service
           178ms systemd-tmpfiles-setup.service
           163ms systemd-random-seed.service
           158ms systemd-sysctl.service
           137ms systemd-journald.service
           109ms haveged.service
           107ms systemd-update-utmp.service
           103ms wpa_supplicant.service
            89ms kmod-static-nodes.service
            77ms user@1000.service
            23ms systemd-user-sessions.service
            23ms systemd-update-done.service
            12ms rtkit-daemon.service
             8ms systemd-modules-load.service
             3ms sys-fs-fuse-connections.mount

查看出错启动项
ystemctl --all | grep not-found
● auditd.service                                                                              not-found inactive   dead      auditd.service
● lvm2-activation.service                                                                     not-found inactive   dead      lvm2-activation.service                                                      
● openvswitch.service                                                                         not-found inactive   dead      openvswitch.service                                                          
● plymouth-quit-wait.service                                                                  not-found inactive   dead      plymouth-quit-wait.service                                                   
● plymouth-quit.service                                                                       not-found inactive   dead      plymouth-quit.service                                                        
● plymouth-start.service                                                                      not-found inactive   dead      plymouth-start.service                                                       
● syslog.service                                                                              not-found inactive   dead      syslog.service 

关闭出错启动项（以 plymouth-start.service 为例）
ystemctl mask plymouth-start.service


如何管理双系统双硬盘默认启动哪个？

搜狗拼音系统启动后不能使用？

如何安装系统的驱动程序几默认配套电脑的驱动程序？



```
sudo vim /usr/share/applications/netease-cloud-music.desktop
Exec=

后更改为

netease-cloud-music %U --no-sandbox
```

# openstack, kvm, qemu-kvm以及libvirt之间的关系

虚拟化技术，分为全虚拟化、半虚拟化、部分虚拟化。

全虚拟化有 vmware、virtualbox、kvm、xen

半虚拟化有 openstack  等

部分虚拟化 LXC、cgroup、docker等



KVM是最底层的hypervisor，它是用来模拟CPU的运行，它缺少了对network和周边I/O的支持，所以我们是没法直接用它的。

QEMU-KVM就是一个完整的模拟器，它是构建基于KVM上面的，它提供了完整的网络和I/O支持。

Openstack不会直接控制qemu-kvm，它会用一个叫libvirt的库去间接控制qemu-kvm。libvirt提供了跨VM平台的功能，它可以控制除了QEMU之外的模拟器，包括vmware, virtualbox， xen等等。

所以为了openstack的跨VM性，所以openstack只会用libvirt而不直接用qemu-kvm。libvirt还提供了一些高级的功能，例如pool/vol管理。



qemu+kvm spice 很舒服，建议花时间配置一下qemu-kvm/libvirt，毕竟qemu更future-proof