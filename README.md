# Z-Admin 我自己的后台管理项目模板

## 目录结构
- api api文件
- configs 配置文件夹
  - config.yaml 配置文件
- controller 控制器文件夹
  - other.go 其他类型的业务逻辑 例如验证码，文件上传等
  - user.go 用户逻辑
- global 全局变量
  - config.go 配置文件结构体
  - global.go 初始化全局变量
  - setting.go 读取配置文件
- model 数据模型
  - index.go 初始化数据库
  - user.go 用户数据模型
- pak 
  - response 统一发送响应
  - log 日志
  - err 错误
- router 路由文件夹
- utils 工具函数文件夹
- web 前端工作目录
- main.go 启动目录


## 实现的功能
- [x] 用户
  - [x] 用户注册
  - [x] 用户登陆
  - [x] 修改密码
- [ ] 角色管理
- [ ] 验证码
- [ ] 动态路由
- [ ] 权限管理
- [ ] 文件上传
- [ ] 断电续传