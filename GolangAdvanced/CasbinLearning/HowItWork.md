# How it work ?

> 我觉得这部分内容得捋一遍，不然马上忘干净

在 Casbin 中，访问控制模型被抽象为基于 **PERM** 元模型（Policy 策略，Effect 效果，Request 请求，Matchers 匹配器）的 CONF 文件

更改项目的授权机制的方法是：更新配置文件

## Request（请求）

这个就是定义请求所需的参数，最基本的请求至少包含元组形式的三个元素：subject, object, action;

也就是说，一个典型的请求定义：r = {sub, obj, act}

## Policy （策略）

确定访问规则的结构，这个组件定义策略规则文档中的字段名称及其顺序：

定义示例：p = {sub, obj, act} 或者 p = {sub, obj, act, eft}

注意这边这个定义，如果不包含 eft，那么默认 eft 为 “allow”

## Matcher（匹配器）

这边是指定如何将请求与策略进行匹配

## Effect （效果）

使用逻辑运算组合已匹配策略的结果

比如说这条匹配式子：e = some(where (p.eft == allow)) && !some(where (p.eft == deny))

## PERM 元模型

上面这四个东西：Policy、Matcher、Effect、Request，组成了 PERM 元模型

## role_definition 角色域

g = _, _ 表示以角色为基础
g = _, _, _ 表示以域为基础（多商户模式）

这个 g 可以理解为授权的意思，比如说，g(Alice, data_admin) 表示的是，给 Alice data_admin 的权限；

“g”写什么都OK，反正就是个占位符。然后这个实际上就是 g = 用户，角色；

## 多租户模型

这个不知道什么玩意
