# Blog API

这是一个基于 Go 语言的typecho博客后端 API 服务,使用 Gin 框架开发,支持文章的增删改查等基本功能。

## 功能特性

- RESTful API 设计
- 基于 Gin 框架
- GORM 数据库操作
- 支持配置文件
- 环境变量配置
- 文章 CRUD 操作
- 分页查询支持

## 技术栈

- Go 1.24
- Gin 1.10.0
- GORM v1.25.12 
- MySQL 5.7
- YAML

## 快速开始

1. 克隆项目
```bash
git clone https://github.com/yourusername/blog_api.git
cd blog_api
```

2. 修改配置文件
```yaml
# config/config.yaml
server:
  port: 8080
  mode: release  # debug/release

database:
  host: 127.0.0.1
  port: 3306
  username: your_username
  password: your_password
  dbname: your_database
  charset: utf8mb4
  parseTime: true
  loc: Local
```

3. 安装依赖
```bash
go mod tidy
```

4. 运行项目
```bash
go run main.go
```

## API 接口

### 文章接口

| 方法   | 路径            | 描述     |
|--------|----------------|----------|
| GET    | /api/v1/posts  | 获取文章列表 |
| POST   | /api/v1/posts  | 创建新文章  |
| GET    | /api/v1/posts/:id | 获取单篇文章 |
| PUT    | /api/v1/posts/:id | 更新文章  |
| DELETE | /api/v1/posts/:id | 删除文章  |

### 健康检查

| 方法   | 路径     | 描述      |
|--------|---------|-----------|
| GET    | /health | 健康检查接口 |

## 开发部署

### 环境变量

支持通过环境变量覆盖配置:

- DB_PASSWORD: 数据库密码
- DB_HOST: 数据库地址 
- DB_PORT: 数据库端口

### 部署

1. 构建
```bash
go build -o blog_api
```

2. 运行
```bash
./blog_api
```

或者使用 Docker:

```bash
# 构建镜像
docker build -t blog-api .

# 运行容器
docker run -d \
  --name blog-api \
  -p 8080:8080 \
  -e DB_PASSWORD=prod_password \
  blog-api
```

## 项目结构

```
blog_api/
├── config/         // 配置文件
├── controllers/    // 控制器
├── database/      // 数据库连接
├── model/         // 数据模型
├── routes/        // 路由定义
├── main.go        // 入口文件
└── README.md
```
