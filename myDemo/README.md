#

## 账号注册登录服务 (Go + Gin + GORM)

### 1. 简介

本需求基于 Go 语言，使用 Gin Web 框架和 GORM ORM 框架实现了基础的用户注册与登录功能。
- 数据库: MySQL 8.0
- 密码安全: 使用 bcrypt 进行加盐哈希存储

### 2. Start the procedure

1. 创建 MySQL 数据库 `go_login_demo`。
2. 修改 `config/db.go` 中的数据库账号密码。
3. 运行:
   ```bash
   go run main.go
   ```
### 业务流程图

```mermaid
sequenceDiagram
    participant User as 用户
    participant API as Go服务
    participant DB as MySQL数据库

    Note over User, DB: 注册流程
    User->>API: POST /api/register (user, pass)
    API->>DB: 查询用户名是否存在
    alt 存在
        DB-->>API: return true
        API-->>User: 400 错误: 用户名已存在
    else 不存在
        API->>API: bcrypt 加密密码
        API->>DB: 插入新用户
        DB-->>API: 成功
        API-->>User: 200 注册成功
    end

    Note over User, DB: 登录流程
    User->>API: POST /api/login (user, pass)
    API->>DB: 根据用户名查找用户
    alt 用户不存在
        DB-->>API: return null
        API-->>User: 401 错误: 用户不存在
    else 用户存在
        API->>API: bcrypt 比对密码(Hash vs Input)
        alt 密码匹配
            API-->>User: 200 登录成功
        else 密码错误
            API-->>User: 401 错误: 密码错误
        end
    end
```