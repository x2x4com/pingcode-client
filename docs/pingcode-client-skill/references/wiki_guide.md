# PingCode 知识管理 (Wiki) 指南 (CLI)

本文档总结了使用 Go CLI 客户端管理 PingCode Wiki 模块的常用命令及核心流程。

`wiki` 是大模块名称，下设 `space`（空间管理）和 `page`（页面管理）两个子模块。

编译完成的二进制文件为 `pingcode-client`

## 全局输出格式

所有命令支持 `--output` 和 `--raw` 全局 flags：

```bash
# 默认 table 格式
pingcode-client wiki space list

# JSON 格式
pingcode-client wiki space list --output json

# 原始 API 响应
pingcode-client wiki space list --raw --output json
```

---

## 1. ID 获取速查表 (Quick Reference)

| 资源类型 | 获取命令 | 备注 |
| :--- | :--- | :--- |
| **空间列表** | `pingcode-client wiki space list` | 获取所有 Wiki 空间及其 ID |
| **页面列表** | `pingcode-client wiki page list --space-id {sid}` | 获取指定空间下的页面 |
| **空间成员** | `pingcode-client wiki space member list --space-id {sid}` | 查看空间成员 |

---

## 2. 空间管理 (Space)

### 2.1 基础操作
- **列出所有空间**:
  ```bash
  pingcode-client wiki space list
  ```
- **获取空间详情**:
  ```bash
  pingcode-client wiki space get -i {space_id}
  ```
- **创建空间**:
  ```bash
  pingcode-client wiki space create \
    --name "团队文档" \
    --identifier "TEAM" \
    --scope-type organization \
    --visibility public \
    --desc "团队共享知识库"
  ```
- **更新空间**:
  ```bash
  pingcode-client wiki space update \
    -i {space_id} \
    --name "新名称" \
    --desc "新描述"
  ```
- **删除空间**:
  ```bash
  pingcode-client wiki space delete -i {space_id}
  ```

### 2.2 成员管理
- **查看成员列表**:
  ```bash
  pingcode-client wiki space member list -i {space_id}
  ```
- **添加成员**:
  ```bash
  # role_id 需要使用系统角色ID（管理员: 100000000000000000000001）
  pingcode-client wiki space member add \
    -i {space_id} \
    --member-id {uid} \
    --member-type user \
    --role-id 100000000000000000000001
  ```
- **移除成员**:
  ```bash
  pingcode-client wiki space member remove \
    -i {space_id} \
    --member-id {uid}
  ```

---

## 3. 页面管理 (Page)

### 3.1 基础操作
- **列出空间下的页面**:
  ```bash
  pingcode-client wiki page list --space-id {space_id}
  ```
- **获取页面详情**:
  ```bash
  pingcode-client wiki page get -i {page_id}
  ```
- **创建页面**:
  ```bash
  pingcode-client wiki page create \
    --space-id {space_id} \
    --name "新页面标题" \
    --type doc
  # 创建文件夹
  pingcode-client wiki page create \
    --space-id {space_id} \
    --name "文档分类" \
    --type folder
  # 在指定父页面下创建
  pingcode-client wiki page create \
    --space-id {space_id} \
    --name "子页面" \
    --parent-id {parent_page_id}
  ```
- **更新页面标题**:
  ```bash
  pingcode-client wiki page update -i {page_id} --name "新标题"
  ```
- **删除页面**:
  ```bash
  pingcode-client wiki page delete -i {page_id}
  ```

### 3.2 页面内容
- **获取页面内容**:
  ```bash
  pingcode-client wiki page content get -i {page_id}
  ```
- **更新页面内容**:
  ```bash
  pingcode-client wiki page content update \
    -i {page_id} \
    --content "# 标题\n正文内容" \
    --format markdown
  ```

### 3.3 版本历史
- **查看页面历史版本**:
  ```bash
  pingcode-client wiki page versions -i {page_id}
  ```

---

## 4. 常见工作流

### 4.1 创建团队知识库
```bash
# 1. 创建 Wiki 空间
pingcode-client wiki space create \
  --name "技术文档" \
  --identifier "TECH" \
  --scope-type organization

# 2. 获取新建空间 ID
pingcode-client wiki space list

# 3. 添加团队成员（role_id: 100000000000000000000001 = 管理员）
pingcode-client wiki space member add \
  -i {space_id} \
  --member-id {uid} \
  --member-type user \
  --role-id 100000000000000000000001

# 4. 创建文档结构
pingcode-client wiki page create --space-id {space_id} --name "架构设计" --type folder
pingcode-client wiki page create --space-id {space_id} --name "API 文档" --parent-id {folder_id}

# 5. 写入内容
pingcode-client wiki page content update \
  -i {page_id} \
  --content "# API 文档\n..." \
  --format markdown
```

### 4.2 查阅文档
```bash
# 1. 找到目标空间
pingcode-client wiki space list

# 2. 列出页面
pingcode-client wiki page list --space-id {space_id}

# 3. 读取内容
pingcode-client wiki page content get -i {page_id}
```

---

## 5. 常见问题 (Troubleshooting)

- **全局 flags**: 所有命令支持 `--output`（json/yaml/markdown/table）和 `--raw`（原始 API 响应）全局 flags。
