# nohup 命令

用途：不挂断地运行命令。

语法：nohup Command [ Arg … ] [　& ]

描述：nohup 命令运行由 Command 参数和任何相关的 Arg 参数指定的命令，忽略所有挂断（SIGHUP）信号。在注销后使用 nohup 命令运行后台中的程序。要运行后台中的 nohup 命令，添加 & （ 表示”and”的符号）到命令的尾部。



运行mysql在后台： /usr/local/mysql/bin/mysqld_safe –user=mysql &

可是有很多程序并不想mysqld一样，这样我们就需要nohup命令

无论是否将 nohup 命令的输出重定向到终端，输出都将附加到当前目录的 nohup.out 文件中。如果当前目录的 nohup.out 文件不可写，输出重定向到 $HOME/nohup.out 文件中。如果没有文件能创建或打开以用于追加，那么 Command 参数指定的命令不可调用。如果标准错误是一个终端，那么把指定的命令写给标准错误的所有输出作为标准输出重定向到相同的文件描述符。

退出状态：该命令返回下列出口值：

126 可以查找但不能调用 Command 参数指定的命令。

127 nohup 命令发生错误或不能查找由 Command 参数指定的命令。

否则，nohup 命令的退出状态是 Command 参数指定命令的退出状态。

如果使用nohup命令提交作业，那么在缺省情况下该作业的所有输出都被重定向到一个名为nohup.out的文件中，除非另外指定了输出文件：

nohup command > myout.file 2>&1 &

输出被重定向到myout.file文件中。

使用 jobs 查看任务。

使用 fg %n　关闭。

另外有两个常用的ftp工具ncftpget和ncftpput，可以实现后台的ftp上传和下载，这样就可以利用这些命令在后台上传和下载文件了。



nohup /root/start.sh &

在shell中回车后提示：

[~]$ appending output to nohup.out

原程序的的标准输出被自动改向到当前目录下的nohup.out文件，起到了log的作用。

但是有时候在这一步会有问题，当把终端关闭后，进程会自动被关闭，察看nohup.out可以看到在关闭终端瞬间服务自动关闭。

咨询红旗Linux工程师后，他也不得其解，在我的终端上执行后，他启动的进程竟然在关闭终端后依然运行。

在第二遍给我演示时，我才发现我和他操作终端时的一个细节不同：他是在当shell中提示了nohup成功后还需要按终端上键盘任意键退回到shell输入命令窗口，然后通过在shell中输入exit来退出终端；而我是每次在nohup执行成功后直接点关闭程序按钮关闭终端.。所以这时候会断掉该命令所对应的session，导致nohup对应的进程被通知需要一起shutdown。















