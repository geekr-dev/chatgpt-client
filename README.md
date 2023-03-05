# chatgpt-client

命令行版 ChatGPT 应用

## 设置环境变量

```bash
export OPENAI_API_KEY='你的KEY'
```
可添加到系统的 `.bashrc` 或 `.zshrc` 文件中


## 启动
```bash
go mod tidy
go run chat.go
```


# 安装到系统
go build .
mv chat-client /usr/local/bin/chatgpt



