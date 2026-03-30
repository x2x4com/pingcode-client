# PingCode 企业组织 (Directory) 指南 (CLI)

本文档总结了使用 Go CLI 客户端管理 PingCode 企业组织目录的常用命令。

`directory` 模块下设 `user`（企业成员）和 `department`（部门）两个子模块。

编译完成的二进制文件为 `pingcode-client`

## 全局输出格式

所有命令支持 `--output` 和 `--raw` 全局 flags：

```bash
# 默认 table 格式
pingcode-client directory user list

# JSON 格式
pingcode-client directory user list --output json

# 原始 API 响应
pingcode-client directory user list --raw --output json
```

---

## 1. ID 获取速查表 (Quick Reference)

| 资源类型 | 获取命令 | 备注 |
| :--- | :--- | :--- |
| **企业成员列表** | `pingcode-client directory user list` | 获取所有成员及其 ID |
| **按关键字搜索成员** | `pingcode-client directory user list --keywords {name}` | 模糊匹配姓名/用户名 |
| **按部门筛选成员** | `pingcode-client directory user list --department-ids {did}` | 获取部门下成员 |
| **部门列表** | `pingcode-client directory department list` | 获取所有部门及其 ID |

---

## 2. 企业成员管理 (User)

### 2.1 查询成员
- **列出所有成员**:
  ```bash
  pingcode-client directory user list
  ```
- **关键字搜索**:
  ```bash
  pingcode-client directory user list --keywords "张三"
  ```
- **按部门筛选**:
  ```bash
  pingcode-client directory user list --department-ids {dept_id}
  ```

### 2.2 创建成员
```bash
pingcode-client directory user create \
  --name "zhangsan" \
  --display-name "张三" \
  --email "zhangsan@company.com" \
  --department-id {dept_id}
```

> **注意**: `--email` 和 `--mobile` 至少填写一个。

### 2.3 更新成员信息
```bash
# 更新显示名
pingcode-client directory user update \
  --id {user_id} \
  --display-name "张三(SRE)"

# 禁用账号
pingcode-client directory user update \
  --id {user_id} \
  --status disabled

# 启用账号
pingcode-client directory user update \
  --id {user_id} \
  --status enabled

# 更新部门
pingcode-client directory user update \
  --id {user_id} \
  --department-id {new_dept_id}
```

---

## 3. 部门管理 (Department)

### 3.1 查询部门
- **列出所有部门**:
  ```bash
  pingcode-client directory department list
  ```

### 3.2 创建部门
```bash
# 创建顶级部门
pingcode-client directory department create --name "技术中心"

# 创建子部门
pingcode-client directory department create \
  --name "后端研发组" \
  --parent-id {parent_dept_id}

# 指定部门负责人
pingcode-client directory department create \
  --name "QA组" \
  --parent-id {pid} \
  --head-id {user_id}
```

### 3.3 更新部门
```bash
# 重命名部门
pingcode-client directory department update \
  --id {dept_id} \
  --name "研发中心"

# 更换负责人
pingcode-client directory department update \
  --id {dept_id} \
  --head-id {new_user_id}

# 迁移到其他父部门
pingcode-client directory department update \
  --id {dept_id} \
  --parent-id {new_parent_id}
```

### 3.4 删除部门
```bash
pingcode-client directory department delete --id {dept_id}
```

---

## 4. 常见工作流

### 4.1 新员工入职
```bash
# 1. 查找目标部门 ID
pingcode-client directory department list

# 2. 创建账号
pingcode-client directory user create \
  --name "lisi" \
  --display-name "李四" \
  --email "lisi@company.com" \
  --mobile "138xxxxxxxx" \
  --department-id {dept_id}
```

### 4.2 组织架构调整
```bash
# 1. 查看现有部门
pingcode-client directory department list

# 2. 新建子部门
pingcode-client directory department create \
  --name "平台组" \
  --parent-id {tech_dept_id}

# 3. 将成员迁移到新部门
pingcode-client directory user update \
  --id {user_id} \
  --department-id {new_dept_id}
```

---

## 5. 常见问题 (Troubleshooting)

- **全局 flags**: 所有命令支持 `--output`（json/yaml/markdown/table）和 `--raw`（原始 API 响应）全局 flags。
