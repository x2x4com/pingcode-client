# PingCode 工作项管理指南 (CLI)

本文档总结了使用 Go CLI 客户端管理 PingCode 工作项的常用命令及核心流程。

编译完成的二进制文件为 `pingcode-client`

## 1. ID 获取速查表 (Quick Reference)

执行任何操作前，通常需要先获取相关资源的 ID。

| 资源类型 | 获取命令 | 备注 |
| :--- | :--- | :--- |
| **项目 (Project)** | `pingcode-client project list` | 所有操作的基础 |
| **工作项类型 (Type)** | `pingcode-client project metadata types` | 如 `task`, `bug`, `story` |
| **状态 (State)** | `pingcode-client project metadata statuses` | 状态流转需符合工作流 |
| **优先级 (Priority)** | `pingcode-client project metadata priorities -p {pid}` | 需指定项目 ID |
| **负责人 (Assignee)** | `pingcode-client project metadata members -p {pid}` | 获取项目成员 ID |
| **迭代 (Sprint)** | `pingcode-client project iteration list -p {pid}` | 获取迭代 ID |
| **看板 (Board)** | `pingcode-client project kanban list -p {pid}` | 获取看板 ID |
| **看板栏 (Entry)** | `pingcode-client project kanban entries -p {pid} -b {bid}` | 获取具体的列 ID |
| **标签 (Tag)** | `pingcode-client project metadata tags list` | 获取标签 ID |
| **父工作项 (Parent)** | `pingcode-client project workitem list -p {pid}` | 获取已有工作项 ID |

---

## 2. 工作项操作

### 2.1 创建一个工作项
```bash
pingcode-client project workitem create \
  -p {project_id} \
  -t "工作项标题" \
  -d "工作项描述" \
  --type-id task \
  --priority-id {priority_id} \
  --assignee-id {member_id} \
  --parent-id {parent_workitem_id}
```
*   **注意**: 成功创建后会返回 `ID` 和 `Identifier` (例如: `CSX-2`)。

### 2.2 更新工作项状态
```bash
pingcode-client project workitem update \
  -i {work_item_id} \
  --state-id {state_id}
```
*   **排错**: 如果报错 `100379 工作项状态不能流转到当前状态`，表示 PingCode 系统中定义的工作流不支持从“当前状态”直接跳到“目标状态”。
    *   **步骤 1：查看当前状态**:
        ```bash
        pingcode-client project workitem get -i {work_item_id}
        ```
        确认输出中的 `Status` 字段。
    *   **步骤 2：列出所有可用状态 ID**:
        ```bash
        pingcode-client project metadata statuses
        ```
        确认你使用的 `{state_id}` 是否正确对应你想要的目标状态。
    *   **步骤 3：中间状态过渡**: 如果必须到达该状态，可能需要先流转到一个中间状态（例如：从“已拒绝”流转回“打开”，再流转到“进行中”）。

### 2.3 将工作项放入看板
在 Kanban 类型的项目中，工作项必须分配到具体的 **看板 (Board)** 和 **看板栏 (Entry/Column)** 才能在视图中可见。

1.  **分配看板和栏位**:
    ```bash
    pingcode-client project workitem update \
      -i {work_item_id} \
      --board-id {board_id} \
      --entry-id {entry_id}
    ```

### 2.4 查看工作项详情
可以使用 `ID` 或 `Identifier` (Key) 查询：
```bash
# 使用 ID
pingcode-client project workitem get -i {work_item_id}

# 使用 Identifier (Key)
pingcode-client project workitem get -k CSX-2
```

### 2.5 标签管理 (Tags)
工作项可以添加多个标签用于分类。

1.  **创建新标签**:
    ```bash
    pingcode-client project metadata tags create -n "重要"
    ```
    *注意: 如果返回 403 权限不足，说明当前 API 凭据没有管理标签的权限。*

2.  **向工作项添加标签**:
    ```bash
    pingcode-client project workitem tag add -i {work_item_id} -t {tag_id}
    ```

3.  **从工作项移除标签**:
    ```bash
    pingcode-client project workitem tag remove -i {work_item_id} -t {tag_id}
    ```

### 2.6 组合过滤查询 (Advanced Filtering)
如果需要执行如“查找标签为 `backend` 且状态为 `打开` 的所有工作项”这样的组合查询，请遵循以下详细步骤：

#### 第一步：获取项目 ID
过滤查询必须指定项目 ID。
```bash
pingcode-client project list
# 假设我们要查询的项目 ID 为: 69beeb90a36e928af7cd9dab
```

#### 第二步：获取标签 ID (Tag IDs)
找到你想要过滤的标签对应的唯一 ID。
```bash
pingcode-client project metadata tags list
# 假设找到 backend 的 ID 为: 69bf0cb2f4befdc8e29f63a1
```

#### 第三步：获取状态 ID (State IDs)
找到你想要过滤的状态对应的唯一 ID。
```bash
pingcode-client project metadata statuses
# 假设找到 “打开” 状态的 ID 为: 6742ced12e4ea7d120649d29
```

#### 第四步：执行组合查询命令
使用 `--tag-ids` 和 `--state-ids` 参数传递这些 ID。

**命令格式：**
```bash
pingcode-client project workitem list \
  -p {项目ID} \
  --tag-ids {标签ID} \
  --state-ids {状态ID}
```

**实操示例：**
```bash
pingcode-client project workitem list \
  -p 69beeb90a36e928af7cd9dab \
  --tag-ids 69bf0cb2f4befdc8e29f63a1 \
  --state-ids 6742ced12e4ea7d120649d29
```

*   **多选支持**: 两个参数都支持以 **逗号 (`,`)** 分隔传递多个 ID。
    *   例如：`--tag-ids id1,id2` 表示包含 id1 **或** id2 标签的工作项。
    *   组合使用时，表示同时满足标签条件 **且** 满足状态条件。

---

## 3. 开发经验总结 (Agent Notes)

*   **omitempty**: 在 SDK 定义中，更新操作 (`PATCH`) 必须在字段上使用 `omitempty`。否则，未赋值的字符串字段会以空字符串形式发送，导致 API 报错（如 `title 不能为空`）。
*   **Identifier 支持**: 系统并不直接支持通过 Identifier 获取详情，CLI 内部通过 `List` 全量获取后在本地进行过滤。
*   **错误处理**: PingCode API 返回的 400 错误通常包含有价值的 JSON 错误信息。在 SDK 中读取并打印响应体对于调试（如检查状态流转规则、必填项缺失）至关重要。
