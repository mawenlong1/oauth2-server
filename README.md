oauth2服务器

## go get翻墙(前提安装ss)
- windows
新建proxy.bat并添加到PATH中
```bat
@echo off
set http_proxy=socks5://127.0.0.1:1080
set https_proxy=socks5://127.0.0.1:1080
go get -u -v %*
echo ...
pause
```
然后就可以使用
proxy.bat golang.org/x/crypto
- mac或linux
需要修改为软件对应的端口号。
```
alias proxy='ALL_PROXY=socks5://127.0.0.1:1086/ \
        http_proxy=http://127.0.0.1:1087/ \
        https_proxy=http://127.0.0.1:1087/ \
        HTTP_PROXY=http://127.0.0.1:1087/ \
        HTTPS_PROXY=http://127.0.0.1:1087/'
```
直接运行上面命令或者添加到配置文件后，可以使用
proxy go get golang.org/x/crypto