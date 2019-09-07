



torbrowser







校验

[https://www.gnupg.org](https://www.gnupg.org/) 

Gpg4win

gpg --verify python-3.5.1.exe.asc

gpg: assuming signed data in 'python-3.5.1.exe'
gpg: Signature made 12/08/15 05:59:22 中国标准时间 using RSA key ID 487034E5
gpg: Can't check signature: No public key

这一步可以看到RSA key ID为487034E5，由于没有公钥，所以我们无法检查签名
下一步我们要通过一些公钥服务器下载公钥。命令指定公钥服务器为hkp://pool.sks-keyservers.net，要下载的ID为487034E5

gpg --keyserver hkp://pool.sks-keyservers.net --recv-keys 487034E5



常见的密钥服务器，来自维基百科https://en.wikipedia.org/wiki/Key_server_%28cryptographic%29
1.keys.gnupg.net
2.hkp://subkeys.pgp.net (服务器池)
3.http://pgp.mit.edu
4.hkp://pool.sks-keyservers.net (服务器池, 也支持TLS: hkps://hkps.pool.sks-keyservers.net)
5.hkp://zimmermann.mayfirst.org (也支持TLS)
6.http://keyserver.ubuntu.com



















































































































