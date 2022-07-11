# README

- Gin+JWT的Demo
- 尝试使用Air实现热部署

参考教程：

- **[Gin框架中使用JWT进行接口认证](https://juejin.cn/post/6844904090690912264)**
- [使用Air实现Go程序实时热重载](https://www.liwenzhou.com/posts/Go/live_reload_with_air/)
  - [Air Github](https://github.com/cosmtrek/air)

## 记录使用Air遇到的坑

这真是一个大坑。在输入`air init`命令，生成`.air.toml`文件之后，将生成的`.air.toml`文件内容替换如下：

```toml
# [Air](https://github.com/cosmtrek/air) TOML 格式的配置文件

# 工作目录
# 使用 . 或绝对路径，请注意 `tmp_dir` 目录必须在 `root` 目录下
root = "."
tmp_dir = "tmp"

[build]
# 只需要写你平常编译使用的shell命令。你也可以使用 `make`
# Windows平台示例: cmd = "go build -o tmp\main.exe ."
cmd = "go build -o ./tmp/main ."
# 由`cmd`命令得到的二进制文件名
# Windows平台示例：bin = "tmp\main.exe"
bin = "tmp/main"
# 自定义执行程序的命令，可以添加额外的编译标识例如添加 GIN_MODE=release
# Windows平台示例：full_bin = "tmp\main.exe"
full_bin = "APP_ENV=dev APP_USER=air ./tmp/main"
# 监听以下文件扩展名的文件.
include_ext = ["go", "tpl", "tmpl", "html"]
# 忽略这些文件扩展名或目录
exclude_dir = ["assets", "tmp", "vendor", "frontend/node_modules"]
# 监听以下指定目录的文件
include_dir = []
# 排除以下文件
exclude_file = []
# 如果文件更改过于频繁，则没有必要在每次更改时都触发构建。可以设置触发构建的延迟时间
delay = 1000 # ms
# 发生构建错误时，停止运行旧的二进制文件。
stop_on_error = true
# air的日志文件名，该日志文件放置在你的`tmp_dir`中
log = "air_errors.log"

[log]
# 显示日志时间
time = true

[color]
# 自定义每个部分显示的颜色。如果找不到颜色，使用原始的应用程序日志。
main = "magenta"
watcher = "cyan"
build = "yellow"
runner = "green"

[misc]
# 退出时删除tmp目录
clean_on_exit = true
```

> 之所以没有按照上面教程所说，手动创建`.air.conf`文件，是因为手动创建的`.air.conf`文件会变成`Read Only`，后续无法修改，非常麻烦。所以我们根据github官方的推荐，使用`air init`命令，生成`.air.toml`文件。

上述文件是李文周老师博客提供了一些汉化注释，大体还是github官方的`air_example.toml`文件（https://github.com/cosmtrek/air/blob/master/air_example.toml）。就是这个官方文件，让我踩了一个大坑。

注意看，上述`.air.toml`文件和官方的`air_example.toml`文件中，通过go build生成可执行文件时，是将其命名为`main`而不是`main.exe`，这就导致后续无论怎么启动`air`，都会报错：找不到`/tmp/main.exe`，因为我们生成的是`tmp.main`文件。

![](https://img-qingbo.oss-cn-beijing.aliyuncs.com/img/20220510183337.png)

所以我们要将`.air.toml`文件修改如下：

```toml
# [Air](https://github.com/cosmtrek/air) TOML 格式的配置文件

# 工作目录
# 使用 . 或绝对路径，请注意 `tmp_dir` 目录必须在 `root` 目录下
root = "."
tmp_dir = "tmp"

[build]
# 只需要写你平常编译使用的shell命令。你也可以使用 `make`
# Windows平台示例: cmd = "go build -o tmp\main.exe ."
cmd = "go build -o ./tmp/main.exe ."
# 由`cmd`命令得到的二进制文件名
# Windows平台示例：bin = "tmp\main.exe"
bin = "tmp/main"
# 自定义执行程序的命令，可以添加额外的编译标识例如添加 GIN_MODE=release
# Windows平台示例：full_bin = "tmp\main.exe"
#full_bin = "APP_ENV=dev APP_USER=air ./tmp/main"
full_bin = "./tmp/main"

# 监听以下文件扩展名的文件.
include_ext = ["go", "tpl", "tmpl", "html"]
# 忽略这些文件扩展名或目录
exclude_dir = ["assets", "tmp", "vendor", "frontend/node_modules"]
# 监听以下指定目录的文件
include_dir = []
# 排除以下文件
exclude_file = []
# 如果文件更改过于频繁，则没有必要在每次更改时都触发构建。可以设置触发构建的延迟时间
delay = 1000 # ms
# 发生构建错误时，停止运行旧的二进制文件。
stop_on_error = true
# air的日志文件名，该日志文件放置在你的`tmp_dir`中
log = "air_errors.log"

[log]
# 显示日志时间
time = true

[color]
# 自定义每个部分显示的颜色。如果找不到颜色，使用原始的应用程序日志。
main = "magenta"
watcher = "cyan"
build = "yellow"
runner = "green"

[misc]
# 退出时删除tmp目录
clean_on_exit = true
```

这样一来，我们直接在控制台输入`air`，就可以启动Air**实现热部署**了。

![](https://img-qingbo.oss-cn-beijing.aliyuncs.com/img/20220510183838.png)

> Tips：修改完代码后`Ctrl+S`，会立马重新部署。但是自动等待一会等AIr检测到更新了，也会重新部署的。

要想停止Air，直接在控制台`Ctrl+C`即可。

![](https://img-qingbo.oss-cn-beijing.aliyuncs.com/img/20220510184130.png)

> 最后好像还是有个bug，就是`Ctrl+C`之后tmp文件夹并不能成功删掉，后续需要的话要手动删除。



## JWT的实践

在代码都编写完成后，可以进行测试了。

首先登录，`localhost:8080/api/user/login`，这样就得到token了：

![](https://img-qingbo.oss-cn-beijing.aliyuncs.com/img/20220511150221.png)

先不在header中携带token，看看能否跑通`localhost:8080/api/test/ga1`：

![](https://img-qingbo.oss-cn-beijing.aliyuncs.com/img/20220511150330.png)

可以发现不携带token是无法访问的，说明我们的JWTAuth()中间件是生效的。

最后在header中携带token：

![](https://img-qingbo.oss-cn-beijing.aliyuncs.com/img/20220511150636.png)

