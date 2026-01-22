# Notes of Casbin learning

## what is Casbin ?

Casbin 是一个**授权库**，适用于需要对资源进行受控访问的场景。

在一个典型的使用场景中，**主体/subject**（用户或者服务）请求访问**对象/object**（资源或者实体）以执行**操作/action**（比如读取、写入或者删除）。

也就是说，这代表了Casbin最常见的处理的原则：{subject, object, action}操作流程

除了上面说的这种标准模型之外，Casbin还支持RBAC（基于角色的访问控制），ABAC（基于属性的访问控制)，还有其他高级模式；

## What Casbin Does ?

- 通过经典的{subject, object, action}格式，或者别的自定义格式。执行策略，allow or deny authorizations（这个是什么意思呢，就是**允许**或则**拒绝**授权
- Manages storage for the access control model and associated policies (管理访问控制模型及其关联策略的存储)
- 处理 用户-角色 和 角色-角色 关系（这个是RBAC中的角色层级概念）
- 识别内置的超级用户，比如说root或者是administrator。这些用户是拥有不受限制的访问权限的，无需显示的权限规则；
- 提供各种内置运算符，用于规则中的模式匹配；

## What Casbin Does Not Do ｜ 就是不能做的！

- 用户身份验证（登陆期间，验证用户名和密码凭据）
- 用户或者角色列表管理

注意到，Casbin 的设计初衷并不是为了密码存储系统。其专注于的是**授权**，不过 Casbin 在 RBAC 模式下，还是会进行维护 用户-角色 关联；

## Casbin Enforcer

这个所谓的Enforcer，就是 强制执行器（虽然不知道为什么翻译的这么奇怪）

Casbin relies on configuration files to specify the access control model.

Casbin 需要两个配置文件，一个是 model.conf, 一个是 policy.csv；其中，model.conf 定义其访问控制模型，而 policy.csv 包含具体的权限规则。

使用 Casbin 是怎么一回事呢？简单的来讲，是和一个结构体进行交互，这个结构体就是小标题里那个 **Enforcer**。初始化的时候，Enforcer会自动加载这两个配置文件。

一种情况：直接用文件加载两个配置

```go
import (
	"github.com/casbin/casbin/v3"
)

e, err := casbin.NewEnforcer("path/to/model.conf", "path/to/policy.csv")
```

还一种情况（只是举个例子，实际上那两个文件爱怎么配置怎么配置，渠道随意）：

```go
import (
    "log"

    "github.com/casbin/casbin/v3"
    "github.com/casbin/casbin/v3/model"
    xormadapter "github.com/casbin/xorm-adapter/v2"
    _ "github.com/go-sql-driver/mysql"
)

// Initialize a Xorm adapter with MySQL database.
a, err := xormadapter.NewAdapter("mysql", "mysql_username:mysql_password@tcp(127.0.0.1:3306)/")
if err != nil {
    log.Fatalf("error: adapter: %s", err)
}

m, err := model.NewModelFromString(`
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = r.sub == p.sub && r.obj == p.obj && r.act == p.act
`)
if err != nil {
    log.Fatalf("error: model: %s", err)
}

e, err := casbin.NewEnforcer(m, a)
if err != nil {
    log.Fatalf("error: enforcer: %s", err)
}
```

上面那段代码里，那个 a 就是存的权限规则，从数据库里去拿；

## 关于检查权限

在发生资源访问之前，先去检查权限！
