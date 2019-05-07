# oauth2服务器
## 构建步骤
### 注意
> 由于项目使用的是Go Modules作为包管理因此go的版本至少为go 1.11。
### 构建
```
go build 
```
下载不了的包需要翻墙或者在.mod里面添加替代的库，参考[在go modules中使用replace替换无法直接获取的package]('https://www.cnblogs.com/apocelipes/p/9609895.html')
### 初始化
```
./oauth2-server init
```
运行会创建相关的表
### 加载数据
```
./oauth2-server loadData oauth/fixtures/scopes.yml oauth/fixtures/roles.yml oauth/fixtures/test_clients.yml oauth/fixtures/test_users.yml
```
加载测试数据数据(数据在oauth/fixtures目录下)
### 运行
```
./oauth2-server run
```

## 程序测试
### 授权码模式
- 参考(http://tools.ietf.org/html/rfc6749#section-4.1)
- 浏览器发起授权码请求
```
http://localhost:3001/web/authorize?client_id=test_client_1&redirect_uri=https://www.example.com&response_type=code&state=somestate&scope=read_write
```
- 提示登录，可以注册然后登录
![login](images/login.png "login")
- 提示用户是否授权应用(test_client_1)访问用户的用户信息
![login](images/auth.png "auth")
- 同意授权，则重定向到应用制定的uri（redirect_uri的值），同时在uri中有授权码。
![login](images/success.png "success")
响应：
```
https://www.example.com/?code=2d6779fd-41a6-41a5-95b4-d881a24c9a39&state=somestate
```
- 拒绝授权
![login](images/failed.png "failed")
响应：
```
https://www.example.com/?error=access_denied&state=somestate
```
- 使用授权码可以获取访问的token
```
curl -X POST --user test_client_1:test_secret http://localhost:3001/v1/oauth/tokens -d "code=2d6779fd-41a6-41a5-95b4-d881a24c9a39&grant_type=authorization_code&redirect_uri=https://www.example.com&scope=read_write"
```
响应
```json
{
    "user_id":"e8ee95a8-a919-487c-a335-5fd9baff66ba",
    "access_token":"984dc8fd-7ee6-42c5-86d4-6b5aa4c10734",
    "expires_in":3600,
    "token_type":"Bearer",
    "scope":"read_write",
    "refresh_token":"be3f9f56-385c-4dc8-8cf7-aa68be7f5214"
}
```

### 简化模式
- 浏览器请求
```
http://localhost:3001/web/authorize?client_id=test_client_1&redirect_uri=https://www.example.com&response_type=token&state=somestate&scope=read_write
```
- 获取响应
```
https://www.example.com/#access_token=9769f172-3e97-4f0d-9207-180612b047bd&expires_in=604800&scope=read_write&state=somestate&token_type=Bearer
```
### 用户名密码模式
- post请求
```
curl -X POST --user test_client_1:test_secret http://localhost:3001/v1/oauth/tokens -d "grant_type=password&username=mwl@123&password=123&scope=read_write"
```
- 响应
```json
{
    "user_id":"e8ee95a8-a919-487c-a335-5fd9baff66ba",
    "access_token":"6ffbfcf1-e9f3-45e1-87e4-a0107f4a94c5",
    "expires_in":3600,
    "token_type":"Bearer",
    "scope":"read_write",
    "refresh_token":"92e5ab18-1cff-43ff-8b6f-7b552bbd4a94"
}
```
### 客户端模式
- 请求
```
curl -X POST --user test_client_1:test_secret http://localhost:3001/v1/oauth/tokens -d "grant_type=client_credentials&scope=read_write"
```
- 响应
```json
{
    "access_token":"fc37b71c-936c-48b1-a218-30095b93133d",
    "expires_in":3600,
    "token_type":"Bearer",
    "scope":"read_write"
}
```

### 令牌校验
- 使用前面的请求获取accesstoken
- 请求
```
curl -X POST --user test_client_2:test_secret http://localhost:3001/v1/oauth/introspect -d "token=6ffbfcf1-e9f3-45e1-87e4-a0107f4a94c5&token_type_hint=access_token"
```
- 响应
```json
{
    "active":true,
    "scope":"read_write",
    "client_id":"test_client_1",
    "username":"mwl@123",
    "token_type":"Bearer",
    "exp":1557240145
}
```
### 令牌刷新
- 使用的是前面获取的refreshtoken
- 请求
```
curl -X POST --user test_client_1:test_secret http://localhost:3001/v1/oauth/tokens -d "grant_type=refresh_token&refresh_token=92e5ab18-1cff-43ff-8b6f-7b552bbd4a94"
```
- 响应
```json
{
    "user_id":"e8ee95a8-a919-487c-a335-5fd9baff66ba",
    "access_token":"548d0510-049a-4b41-b9c7-f43b4abb9c87",
    "expires_in":3600,
    "token_type":"Bearer",
    "scope":"read_write",
    "refresh_token":"ccbf323b-5fc4-4321-83bd-cb85b32e593f"
}
```

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