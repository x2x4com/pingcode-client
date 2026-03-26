# PingCode 测试管理 (TestHub) 指南 (CLI)

本文档总结了使用 Go CLI 客户端管理 PingCode TestHub 测试模块的常用命令及核心流程。

`testhub` 模块下设 `library`（测试库）、`case`（测试用例）、`plan`（测试计划）和 `run`（测试执行）四个子模块。

编译完成的二进制文件为 `pingcode-client`

## 1. ID 获取速查表 (Quick Reference)

| 资源类型 | 获取命令 | 备注 |
| :--- | :--- | :--- |
| **测试库列表** | `pingcode-client testhub library list` | 获取所有测试库及其 ID |
| **测试用例列表** | `pingcode-client testhub case list --library-id {lid}` | 获取指定库下的用例 |
| **测试计划列表** | `pingcode-client testhub plan list --library-id {lid}` | 获取指定库下的计划 |
| **用例模块列表** | `pingcode-client testhub library suite list --library-id {lid}` | 获取用例分组 ID |
| **库成员列表** | `pingcode-client testhub library member list --library-id {lid}` | 查看测试库成员 |

---

## 2. 测试库管理 (Library)

### 2.1 基础操作
- **列出所有测试库**:
  ```bash
  pingcode-client testhub library list
  ```
- **创建测试库**:
  ```bash
  pingcode-client testhub library create \
    --name "核心功能测试库" \
    --identifier "CORE" \
    --scope-type organization \
    --visibility public \
    --desc "覆盖核心业务场景的测试用例"
  ```
- **更新测试库**:
  ```bash
  pingcode-client testhub library update --id {lib_id} --name "新名称" --desc "新描述"
  ```

### 2.2 成员管理
- **查看成员**:
  ```bash
  pingcode-client testhub library member list --library-id {lib_id}
  ```
- **添加成员**:
  ```bash
  # role_id 需使用系统角色ID（管理员: 100000000000000000000001）
  pingcode-client testhub library member add \
    --library-id {lib_id} \
    --member-id {user_id} \
    --member-type user \
    --role-id 100000000000000000000001
  ```
- **移除成员**:
  ```bash
  pingcode-client testhub library member remove --library-id {lib_id} --member-id {user_id}
  ```

### 2.3 用例模块 (Suite)
- **列出模块**:
  ```bash
  pingcode-client testhub library suite list --library-id {lib_id}
  ```
- **创建模块**:
  ```bash
  pingcode-client testhub library suite create --library-id {lib_id} --name "登录功能"
  ```
- **删除模块**:
  ```bash
  pingcode-client testhub library suite delete --library-id {lib_id} --suite-id {sid}
  ```

---

## 3. 测试用例管理 (Case)

### 3.1 基础操作
- **列出测试用例**:
  ```bash
  pingcode-client testhub case list --library-id {lib_id}
  # 按模块过滤（可选）
  pingcode-client testhub case list --library-id {lib_id} --suite-id {suite_id}
  ```
- **获取用例详情**:
  ```bash
  pingcode-client testhub case get --id {case_id}
  ```
- **创建用例**:
  ```bash
  pingcode-client testhub case create \
    --library-id {lib_id} \
    --title "用户登录成功场景" \
    --suite-id {suite_id} \
    --maintenance-id {user_id} \
    --desc "验证用户使用正确账号密码登录" \
    --precondition "用户已注册且账号状态正常"
  ```
- **更新用例**:
  ```bash
  pingcode-client testhub case update --id {case_id} --title "新标题" --desc "更新的描述"
  ```
- **删除用例**:
  ```bash
  pingcode-client testhub case delete --id {case_id}
  ```

---

## 4. 测试计划管理 (Plan)

### 4.1 查看计划
```bash
pingcode-client testhub plan list --library-id {lib_id}
```

### 4.2 创建测试计划
```bash
pingcode-client testhub plan create \
  --library-id {lib_id} \
  --name "Sprint 10 回归测试" \
  --type-id {plan_type_id} \
  --assignee-id {user_id} \
  --start-at 1700000000 \
  --end-at 1700604800
```

---

## 5. 测试执行 (Run)

### 5.1 创建执行记录
```bash
pingcode-client testhub run create \
  --library-id {lib_id} \
  --plan-id {plan_id} \
  --case-id {case_id} \
  --executor-id {user_id}
```

---

## 6. 常见工作流

### 6.1 新版本回归测试
```bash
# 1. 列出测试库
pingcode-client testhub library list

# 2. 查看用例模块
pingcode-client testhub library suite list --library-id {lib_id}

# 3. 查看用例列表
pingcode-client testhub case list --library-id {lib_id}

# 4. 创建测试计划
pingcode-client testhub plan create \
  --library-id {lib_id} \
  --name "v2.1.0 回归测试" \
  --type-id {type_id} \
  --assignee-id {qa_user_id} \
  --start-at {start_ts} \
  --end-at {end_ts}

# 5. 将用例加入执行
pingcode-client testhub run create \
  --library-id {lib_id} \
  --plan-id {plan_id} \
  --case-id {case_id} \
  --executor-id {qa_user_id}
```

### 6.2 新功能用例维护
```bash
# 1. 创建用例模块
pingcode-client testhub library suite create --library-id {lib_id} --name "支付功能"

# 2. 创建测试用例
pingcode-client testhub case create \
  --library-id {lib_id} \
  --title "微信支付成功场景" \
  --suite-id {suite_id} \
  --maintenance-id {user_id} \
  --precondition "用户已绑定微信支付"

# 3. 查看用例列表确认
pingcode-client testhub case list --library-id {lib_id}
```
