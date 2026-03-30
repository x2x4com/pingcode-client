# PingCode 项目管理指南 (CLI)

本文档总结了使用 Go CLI 客户端管理 PingCode 项目模块中迭代（Sprint）、版本（Release）和看板（Kanban）的常用命令及核心流程。

编译完成的二进制文件为 `pingcode-client`

## 全局输出格式

所有命令支持 `--output` 和 `--raw` 全局 flags：

```bash
# 默认 table 格式
pingcode-client project iteration list --project-id {pid}

# JSON 格式
pingcode-client project iteration list --project-id {pid} --output json

# 原始 API 响应
pingcode-client project iteration list --project-id {pid} --raw --output json
```

---

## 1. ID 获取速查表 (Quick Reference)

| 资源类型 | 获取命令 | 备注 |
| :--- | :--- | :--- |
| **项目列表** | `pingcode-client project list` | 所有操作的基础 |
| **迭代列表** | `pingcode-client project iteration list --project-id {pid}` | 获取项目下所有 Sprint |
| **版本列表** | `pingcode-client project version list --project-id {pid}` | 获取项目下所有 Release |
| **看板列表** | `pingcode-client project kanban list --project-id {pid}` | 获取项目下所有看板 |
| **看板栏列表** | `pingcode-client project kanban entries --project-id {pid} --board-id {bid}` | 获取看板列 (Entry/Column) |

---

## 2. 迭代管理 (Sprint)

迭代用于规划一个开发周期内要完成的工作项集合。

### 2.1 查看迭代

```bash
pingcode-client project iteration list --project-id {project_id}
```

### 2.2 创建迭代

```bash
pingcode-client project iteration create \
  --project-id {project_id} \
  --name "Sprint 10" \
  --start-date "2026-04-01" \
  --end-date "2026-04-14"
```

### 2.3 更新迭代

```bash
pingcode-client project iteration update \
  -i {iteration_id} \
  --name "Sprint 10 (已调整)" \
  --end-date "2026-04-21"
```

### 2.4 删除迭代

```bash
pingcode-client project iteration delete -i {iteration_id}
```

### 2.5 将工作项加入迭代

通过更新工作项的 `--sprint-id` 字段实现：

```bash
pingcode-client project workitem update \
  -i {work_item_id} \
  --sprint-id {iteration_id}
```

---

## 3. 版本管理 (Release)

版本用于标记一个正式发布节点，将多个已完成的工作项归入同一个发布版本。

### 3.1 查看版本

```bash
pingcode-client project version list --project-id {project_id}
```

### 3.2 创建版本

```bash
pingcode-client project version create \
  --project-id {project_id} \
  --name "v1.2.0" \
  --desc "本版本包含 OAuth 登录、性能优化等功能" \
  --release-date "2026-04-30"
```

### 3.3 更新版本（确认发布时间）

通过设置 `--release-date` 来记录实际发布日期：

```bash
pingcode-client project version update \
  -i {version_id} \
  --name "v1.2.0" \
  --release-date "2026-04-30"
```

### 3.4 删除版本

```bash
pingcode-client project version delete -i {version_id}
```

### 3.5 将工作项关联到版本

> ⚠️ **注意**：当前 CLI 的 `workitem update` 未暴露 `--version-id` 参数，请在 PingCode Web 界面操作，或等待后续 CLI 版本支持。

---

## 4. 看板管理 (Kanban)

看板用于可视化工作项的流转状态，适合 Kanban 流程项目。

### 4.1 查看看板

```bash
pingcode-client project kanban list --project-id {project_id}
```

### 4.2 查看看板栏（列）

```bash
pingcode-client project kanban entries \
  --project-id {project_id} \
  --board-id {board_id}
```

### 4.3 创建看板

```bash
pingcode-client project kanban create \
  --project-id {project_id} \
  --name "主看板"
```

### 4.4 将工作项放入看板指定列

```bash
pingcode-client project workitem update \
  -i {work_item_id} \
  --board-id {board_id} \
  --entry-id {entry_id}
```

---

## 5. 典型 DevOps 发布流程示例

```bash
# 1. 查看当前迭代情况
pingcode-client project iteration list --project-id {project_id}

# 2. 查看待发布版本
pingcode-client project version list --project-id {project_id}

# 3. 创建新发布版本
pingcode-client project version create \
  --project-id {project_id} \
  --name "v1.2.0" \
  --desc "功能：OAuth登录、性能优化" \
  --release-date "2026-04-30"

# 4. 查询已完成工作项（可在 Web 端将其关联到版本）
#    先获取"已完成"状态 ID
pingcode-client project metadata statuses
#    查询已完成工作项
pingcode-client project workitem list \
  --project-id {project_id} \
  --state-ids {done_state_id}

# 5. 部署完成后更新版本的 release-date 确认发布时间
pingcode-client project version update -i {version_id} --release-date "2026-04-30"

# 6. 创建下一个迭代
pingcode-client project iteration create \
  --project-id {project_id} \
  --name "Sprint 11" \
  --start-date "2026-05-01" \
  --end-date "2026-05-14"
```

---

## 6. 常见问题 (Troubleshooting)

- **版本 vs 迭代的区别**: 迭代（Sprint）是时间盒，描述"什么时候做"；版本（Release）是发布节点，描述"发布了什么"。两者可以组合使用。
- **工作项必须在项目内**: 所有操作均需要 `--project-id`，请先通过 `pingcode-client project list` 确认。
- **看板可见性**: 在 Kanban 类型项目中，工作项若未分配 `--board-id` 和 `--entry-id`，则不会在看板视图中显示。
- **全局 flags**: 所有命令支持 `--output`（json/yaml/markdown/table）和 `--raw`（原始 API 响应）全局 flags。
