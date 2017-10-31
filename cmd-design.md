## 命令设计
### 命令使用:agenda [command]

##### 可使用的command:
1. help
2. register
3. login
4. logout

#### agenda help:列出命令的说明
例子：agenda help
***
#### agenda register【flags】:注册一个命令的账号
【flags】：
+ -u 用户名
+ -p 密码
+ -m 邮箱地址
+ -t 手机号码

例子：agenda register -u young -p password -m 1425434617@qq.com -t 134556276572

 ***
#### agenda login 【flags】：通过已经注册的账号进行登录
【flags】：
+ -u 用户名
+ -p 密码

例子：agenda login -u young -p password

***

#### agenda logout:登出agenda

例子：agenda logout
