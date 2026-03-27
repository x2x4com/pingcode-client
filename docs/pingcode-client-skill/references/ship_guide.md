# PingCode 产品管理 (Ship) 指南 (CLI)

本文档总结了使用 Go CLI 客户端管理 PingCode 产品管理 (Ship) 模块的常用命令及核心流程。

`ship` 是大模块名称，下设 `product`（产品管理）、`ideas`（需求管理）和 `tickets`（工单管理）三个核心子模块。

编译完成的二进制文件为 `pingcode-client`

## 1. ID 获取速查表 (Quick Reference)

### 1.1 产品相关 (Product)
| 资源类型 | 获取命令 | 备注 |
| :--- | :--- | :--- |
| **产品列表** | `pingcode-client ship product list` | 获取所有产品及其 ID |
| **产品成员** | `pingcode-client ship product members list --product-id {pid}` | 查看产品成员 |
| **产品模块** | `pingcode-client ship product modules list --product-id {pid}` | 查看产品下的需求模块 |
| **产品排期** | `pingcode-client ship product plans --product-id {pid}` | 查看产品的排期计划 |
| **产品标签** | `pingcode-client ship product tags list --product-id {pid}` | 查看产品的标签 |
| **工单渠道** | `pingcode-client ship product channels --product-id {pid}` | 查看产品的工单渠道 |
| **工单类型** | `pingcode-client ship product ticket-types --product-id {pid}` | 查看产品的工单类型 |

### 1.2 需求相关 (Idea)
| 资源类型 | 获取命令 | 备注 |
| :--- | :--- | :--- |
| **需求列表** | `pingcode-client ship ideas list --product-id {pid}` | 获取现有需求 ID |
| **需求状态** | `pingcode-client ship ideas states --product-id {pid}` | 需求流转的目标状态 |
| **优先级** | `pingcode-client ship ideas priorities --product-id {pid}` | 需求的紧急程度 |
| **属性列表** | `pingcode-client ship ideas properties --product-id {pid}` | 获取自定义属性 ID |

---

## 2. 产品管理 (Product)

### 2.1 产品基础操作
- **创建产品**:
  ```bash
  pingcode-client ship product create --name "新产品" --identifier "PROD" --desc "描述信息"
  ```
- **更新产品**:
  ```bash
  pingcode-client ship product update --product-id {pid} --name "新名称"
  ```

### 2.2 成员管理
- **添加成员**:
  ```bash
  pingcode-client ship product members add --product-id {pid} --member-id {uid} --member-type user --role-id {rid}
  ```
- **移除成员**:
  ```bash
  pingcode-client ship product members remove --product-id {pid} --member-id {uid}
  ```

### 2.3 需求模块管理
- **添加模块**:
  ```bash
  pingcode-client ship product modules add --product-id {pid} --name "模块A" --type module
  ```
- **移除模块**:
  ```bash
  pingcode-client ship product modules remove --product-id {pid} --suite-id {sid}
  ```

### 2.4 标签管理
- **添加标签**:
  ```bash
  pingcode-client ship product tags add --product-id {pid} --name "核心"
  ```
- **移除标签**:
  ```bash
  pingcode-client ship product tags remove --product-id {pid} --tag-id {tid}
  ```

---

## 3. 需求管理 (Ideas)

### 3.1 需求基础操作
- **创建需求**:
  ```bash
  pingcode-client ship ideas create --product-id {pid} --title "需求标题" --desc "描述" --priority-id {prid} --assignee-id {uid} --suite-id {sid}
  ```
- **更新需求**:
  ```bash
  pingcode-client ship ideas update --product-id {pid} --idea-id {iid} --state-id {sid} --priority-id {prid}
  ```

### 3.2 查看流转历史
```bash
pingcode-client ship ideas histories --product-id {pid} --idea-id {idea_id}
```

---

## 4. 常见问题 (Troubleshooting)

*   **模块区分**: 始终记住 `ship` 是顶层模块，`product` 用于配置和管理产品本身，而 `ideas` 用于管理产品下的具体需求。
*   **ID 必填**: 绝大多数子命令都需要 `--product-id`，请先通过 `pingcode-client ship product list` 获取。

---

## 5. 工单管理 (Tickets)

工单模块用于管理来自客户的问题反馈和功能请求。

### 5.1 ID 速查
| 资源类型 | 获取命令 | 备注 |
| :--- | :--- | :--- |
| **工单优先级** | `pingcode-client ship tickets priorities --product-id {pid}` | 获取优先级 ID |
| **工单状态** | `pingcode-client ship tickets states --product-id {pid}` | 获取状态 ID |
| **工单类型** | `pingcode-client ship product ticket-types --product-id {pid}` | 获取工单类型 ID |
| **渠道列表** | `pingcode-client ship product channels --product-id {pid}` | 获取渠道 ID |

### 5.2 工单基础操作
- **列出工单**:
  ```bash
  pingcode-client ship tickets list --product-id {pid}
  # 按状态筛选
  pingcode-client ship tickets list --product-id {pid} --state-id {sid}
  # 关键字搜索
  pingcode-client ship tickets list --product-id {pid} --keywords "登录"
  ```
- **获取工单详情**:
  ```bash
  pingcode-client ship tickets get --product-id {pid} --id {ticket_id}
  ```
- **创建工单**:
  ```bash
  pingcode-client ship tickets create \
    --product-id {pid} \
    --title "用户无法登录" \
    --type-id {type_id} \
    --assignee-id {user_id} \
    --priority-id {priority_id} \
    --channel-id {channel_id}
  ```
- **更新工单**:
  ```bash
  # 更新状态
  pingcode-client ship tickets update --product-id {pid} --id {ticket_id} --state-id {new_state_id}
  # 更改负责人
  pingcode-client ship tickets update --product-id {pid} --id {ticket_id} --assignee-id {user_id}
  # 设置解决方案
  pingcode-client ship tickets update --id {ticket_id} --solution-id {solution_id}
  ```
