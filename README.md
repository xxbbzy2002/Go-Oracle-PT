# Go Oracle Project Template

使用 go-ora 连接 Oracle 数据库的 Go 项目模板。

## 项目结构

```
├── cmd/app/          # 主程序入口
├── internal/
│   ├── config/       # 配置加载
│   ├── database/     # Oracle 数据库连接
│   └── logger/       # 日志模块
├── configs/          # 配置文件
└── logs/             # 日志输出目录
```

## 快速开始

1. 修改配置文件 `configs/config.yaml`
2. 运行程序：
   ```bash
   go run ./cmd/app
   ```

## 构建

```bash
go build -o bin/app ./cmd/app
```

## 依赖

- [go-ora](https://github.com/sijms/go-ora) - Oracle 数据库驱动
- [zap](https://github.com/uber-go/zap) - 日志库
