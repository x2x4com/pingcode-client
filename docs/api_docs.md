# REST API
> 

**Base URL:** `https://open.pingcode.com`

---

## facade

### GetV1ShipIdea_priorities

**`GET /v1/ship/idea_priorities`**

---

### GetV1ShipIdea_prioritiesPriority_id

**`GET /v1/ship/idea_priorities/{priority_id}`**

---

## facade

### GetV1ShipProductsProduct_idChannelsChannel_id

**`GET /v1/ship/products/{product_id}/channels/{channel_id}`**

---

## facade

### GetV1ShipTicket_priorities

**`GET /v1/ship/ticket_priorities`**

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "5cb9466afda1ce4ca0090005",
            "url": "https://rest_api_root/v1/ship/ticket_priorities/5cb9466afda1ce4ca0090005",
            "name": "P0"
        }
    ]
}
```

---

## facade

### GetV1ShipTicket_solutions

**`GET /v1/ship/ticket_solutions`**

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "6422711c3f12e6c1e46d40e9",
            "url": "https://rest_api_root/v1/ship/ticket_solutions/6422711c3f12e6c1e46d40e9",
            "name": "进入需求池"
        }
    ]
}
```

---

### GetV1ShipTicket_solutionsTicket_solution_id

**`GET /v1/ship/ticket_solutions/{ticket_solution_id}`**

**响应示例：**

```json
{
    "id": "6422711c3f12e6c1e46d40e9",
    "url": "https://rest_api_root/v1/ship/ticket_solutions/6422711c3f12e6c1e46d40e9",
    "name": "进入需求池"
}
```

---

## facade

### GetV1ShipProductsProduct_idPlansPlan_id

**`GET /v1/ship/products/{product_id}/plans/{plan_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `product_id` | String | 产品的id。 |
| `plan_id` | String | 需求排期id。 |

**响应示例：**

```json
{
    "id": "63bb744414bd13c9def24ce4",
    "url": "http://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e6/plans/63bb744414bd13c9def24ce4",
    "product": {
        "id": "6422711c3f12e6c1e46d40e6",
        "url": "http://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e6",
        "identifier": "SLC",
        "name": "示例产品",
        "is_archived": 0,
        "is_deleted": 0
    },
    "name": "账号管理",
    "assignee": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "start_at": 1672704000,
    "end_at": 1672963199
}
```

---

## DevOps_数据集成

### 交付

---

### 代码

---

### 构建

---

## Kanban

### 创建一个泳道

**`POST /v1/project/projects/{project_id}/boards/{board_id}/swimlanes`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `project_id` | String | 项目的id。 |
| `board_id` | String | 看板的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `name` | String | 泳道的名称。在同一看板下该名称是唯一的。 |

**响应示例：**

```json
{
    "id": "5bb623f6a70571487ea44357",
    "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919/boards/5eb623f6a70571487ea47223/swimlanes/5bb623f6a70571487ea44357",
    "project": {
        "id": "5eb623f6a70571487ea41919",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919",
        "identifier": "KANBAN",
        "name": "kanban",
        "type": "kanban",
        "is_archived": 0,
        "is_deleted": 0
    },
    "board": {
        "id": "5eb623f6a70571487ea47223",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919/boards/5eb623f6a70571487ea47223",
        "name": "默认看板",
        "work_item_types": ["epic", "feature", "story"]
    },
    "name": "一个泳道",
    "is_system": false
}
```

---

### 创建一个看板

**`POST /v1/project/projects/{project_id}/boards`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `project_id` | String | 项目的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `name` | String | 看板的名字。在同一个项目中该名字是唯一的。 |
| `work_item_types` *(可选)* | String[] | 看板支持的工作项类型列表。自定义工作项类型只支持hybrid类型项目里的看板。 |

**响应示例：**

```json
{
    "id": "5eb623f6a70571487ea47222",
    "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919/boards/5eb623f6a70571487ea47222",
    "project": {
        "id": "5eb623f6a70571487ea41919",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919",
        "identifier": "KANBAN",
        "name": "kanban",
        "type": "kanban",
        "is_archived": 0,
        "is_deleted": 0
    },
    "name": "一个看板",
    "work_item_types": ["epic", "feature", "story", "6385c650fef18f2d7222d15d"]
}
```

---

### 创建一个看板栏

**`POST /v1/project/projects/{project_id}/boards/{board_id}/entries`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `project_id` | String | 项目的id。 |
| `board_id` | String | 看板的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `name` | String | 看板栏的名称。在同一看板下该名称是唯一的。 |
| `wip_limit` *(可选)* | Number | 在制品数量。 |
| `is_split` *(可选)* | Boolean | 是否将看板栏拆分为进行中和已完成。 |
| `definition_of_done` *(可选)* | String | 完成的定义。 |

**响应示例：**

```json
{
   "id": "5ab623f6a70571487ea45634",
   "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919/boards/5eb623f6a70571487ea47222/entries/5ab623f6a70571487ea45634",
   "project": {
      "id": "5eb623f6a70571487ea41919",
      "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919",
      "identifier": "KANBAN",
      "name": "kanban",
      "type": "kanban",
      "is_archived": 0,
      "is_deleted": 0
   },
   "board": {
      "id": "5eb623f6a70571487ea47222",
      "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919/boards/5eb623f6a70571487ea47222",
      "name": "kanban",
      "work_item_types": ["epic", "feature", "story"]
   },
   "name": "一个看板栏",
   "is_system": false,
   "is_split": true,
   "wip_limit": 1,
   "definition_of_done": "Unit test passed"
}
```

---

### 删除一个泳道

**`DELETE /v1/project/projects/{project_id}/boards/{board_id}/swimlanes/{swimlane_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `project_id` | String | 项目的id。 |
| `board_id` | String | 看板的id。 |
| `swimlane_id` | String | 泳道的id. |

**响应示例：**

```json
{
    "id": "5bb623f6a70571487ea44357",
    "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919/boards/5eb623f6a70571487ea47223/swimlanes/5bb623f6a70571487ea44357",
    "project": {
        "id": "5eb623f6a70571487ea41919",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919",
        "identifier": "KANBAN",
        "name": "kanban",
        "type": "kanban",
        "is_archived": 0,
        "is_deleted": 0
    },
    "board": {
        "id": "5eb623f6a70571487ea47223",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919/boards/5eb623f6a70571487ea47223",
        "name": "默认看板",
        "work_item_types": ["epic", "feature", "story"]
    },
    "name": "一个泳道",
    "is_system": false
}
```

---

### 删除一个看板

**`DELETE /v1/project/projects/{project_id}/boards/{board_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `project_id` | String | 项目的id。 |
| `board_id` | String | 看板的id。 |

**响应示例：**

```json
{
    "id": "5eb623f6a70571487ea47222",
    "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919/boards/5eb623f6a70571487ea47222",
    "project": {
        "id": "5eb623f6a70571487ea41919",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919",
        "identifier": "KANBAN",
        "name": "kanban",
        "type": "kanban",
        "is_archived": 0,
        "is_deleted": 0
    },
    "name": "一个看板",
    "work_item_types": ["epic", "feature"]
}
```

---

### 删除一个看板栏

**`DELETE /v1/project/projects/{project_id}/boards/{board_id}/entries/{entry_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `project_id` | String | 项目的id。 |
| `board_id` | String | 看板的id。 |
| `entry_id` | String | 看板栏的id。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "5ab623f6a70571487ea45634",
            "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919/boards/5eb623f6a70571487ea47223/entries/5ab623f6a70571487ea45634",
            "project": {
                "id": "5eb623f6a70571487ea41919",
                "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919",
                "identifier": "KANBAN",
                "name": "kanban",
                "type": "kanban",
                "is_archived": 0,
                "is_deleted": 0
            },
            "board": {
                "id": "5eb623f6a70571487ea47223",
                "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919/boards/5eb623f6a70571487ea47223",
                "name": "默认看板",
                "work_item_types": ["epic", "feature", "story"]
            },
            "name": "需求池",
            "is_system": false,
            "is_split": true,
            "wip_limit": 1,
            "definition_of_done": "Unit test passed"
        }
    ]
}
```

---

### 获取泳道列表

**`GET /v1/project/projects/{project_id}/boards/{board_id}/swimlanes`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `project_id` | String | 项目的id。 |
| `board_id` | String | 看板的id。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "5bb623f6a70571487ea44357",
            "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919/boards/5eb623f6a70571487ea47223/swimlanes/5bb623f6a70571487ea44357",
            "project": {
                "id": "5eb623f6a70571487ea41919",
                "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919",
                "identifier": "KANBAN",
                "name": "kanban",
                "type": "kanban",
                "is_archived": 0,
                "is_deleted": 0
            },
            "board": {
                "id": "5eb623f6a70571487ea47223",
                "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919/boards/5eb623f6a70571487ea47223",
                "name": "默认看板",
                "work_item_types": ["epic", "feature", "story"]
            },
            "name": "一个泳道",
            "is_system": false
        }
    ]
}
```

---

### 获取看板列表

**`GET /v1/project/projects/{project_id}/boards`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `project_id` | String | 项目的id。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "5eb623f6a70571487ea47222",
            "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919/boards/5eb623f6a70571487ea47222",
            "project": {
                "id": "5eb623f6a70571487ea41919",
                "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919",
                "identifier": "KANBAN",
                "name": "kanban",
                "type": "kanban",
                "is_archived": 0,
                "is_deleted": 0
            },
            "name": "kanban",
            "work_item_types": ["epic", "feature", "story"]
        }
    ]
}
```

---

### 获取看板栏列表

**`GET /v1/project/projects/{project_id}/boards/{board_id}/entries`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `project_id` | String | 项目的id。 |
| `board_id` | String | 看板的id。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "5ab623f6a70571487ea45634",
            "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919/boards/5eb623f6a70571487ea47223/entries/5ab623f6a70571487ea45634",
            "project": {
                "id": "5eb623f6a70571487ea41919",
                "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919",
                "identifier": "KANBAN",
                "name": "kanban",
                "type": "kanban",
                "is_archived": 0,
                "is_deleted": 0
            },
            "board": {
                "id": "5eb623f6a70571487ea47223",
                "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919/boards/5eb623f6a70571487ea47223",
                "name": "默认看板",
                "work_item_types": ["epic", "feature", "story"]
            },
            "name": "需求池",
            "is_system": false,
            "is_split": true,
            "wip_limit": 1,
            "definition_of_done": "Unit test passed"
        }
    ]
}
```

---

### 部分更新一个泳道

**`PATCH /v1/project/projects/{project_id}/boards/{board_id}/swimlanes/{swimlane_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `project_id` | String | 项目的id。 |
| `board_id` | String | 看板的id。 |
| `swimlane_id` | String | 泳道的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `name` *(可选)* | String | 泳道的名称。在同一看板下该名称是唯一的。 |

**响应示例：**

```json
{
    "id": "5bb623f6a70571487ea44357",
    "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919/boards/5eb623f6a70571487ea47223/swimlanes/5bb623f6a70571487ea44357",
    "project": {
        "id": "5eb623f6a70571487ea41919",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919",
        "identifier": "KANBAN",
        "name": "kanban",
        "type": "kanban",
        "is_archived": 0,
        "is_deleted": 0
    },
    "board": {
        "id": "5eb623f6a70571487ea47223",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919/boards/5eb623f6a70571487ea47223",
        "name": "默认看板",
        "work_item_types": ["epic", "feature", "story"]
    },
    "name": "一个泳道",
    "is_system": false
}
```

---

### 部分更新一个看板

**`PATCH /v1/project/projects/{project_id}/boards/{board_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `project_id` | String | 项目的id。 |
| `board_id` | String | 看板的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `name` *(可选)* | String | 看板的名字。在同一个项目中该名字是唯一的。 |
| `work_item_types` *(可选)* | String[] | 看板支持的工作项类型列表。 |

**响应示例：**

```json
{
    "id": "5eb623f6a70571487ea47222",
    "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919/boards/5eb623f6a70571487ea47222",
    "project": {
        "id": "5eb623f6a70571487ea41919",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919",
        "identifier": "KANBAN",
        "name": "kanban",
        "type": "kanban",
        "is_archived": 0,
        "is_deleted": 0
    },
    "name": "一个看板",
    "work_item_types": ["epic", "feature", "6385c650fef18f2d7222d15d"]
}
```

---

### 部分更新一个看板栏

**`PATCH /v1/project/projects/{project_id}/boards/{board_id}/entries/{entry_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `project_id` | String | 项目的id。 |
| `board_id` | String | 看板的id。 |
| `entry_id` | String | 看板栏的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `name` *(可选)* | String | 看板栏的名称。在同一看板下该名称是唯一的。 |
| `wip_limit` *(可选)* | Number | 在制品数量。 |
| `is_split` *(可选)* | Boolean | 是否将看板栏拆分为进行中和已完成 |
| `definition_of_done` *(可选)* | String | 完成的定义。 |

**响应示例：**

```json
{
    "id": "5ab623f6a70571487ea45634",
    "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919/boards/5eb623f6a70571487ea47223/entries/5ab623f6a70571487ea45634",
    "project": {
        "id": "5eb623f6a70571487ea41919",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919",
        "identifier": "KANBAN",
        "name": "kanban",
        "type": "kanban",
        "is_archived": 0,
        "is_deleted": 0
    },
    "board": {
        "id": "5eb623f6a70571487ea47223",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919/boards/5eb623f6a70571487ea47223",
        "name": "默认看板",
        "work_item_types": ["epic", "feature", "story"]
    },
    "name": "需求池",
    "is_system": false,
    "is_split": true,
    "wip_limit": 1,
    "definition_of_done": "Unit test passed"
}
```

---

## Scrum

### 创建一个迭代

**`POST /v1/project/projects/{project_id}/sprints`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `project_id` | String | 项目的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `name` | String | 迭代的名称。 |
| `start_at` | Number | 迭代开始的时间。 |
| `end_at` | Number | 迭代结束的时间。 |
| `assignee_id` | String | 迭代负责人的id。 |
| `description` *(可选)* | String | 迭代的描述。 |
| `status` *(可选)* | String | 迭代的状态。 |
| `category_ids` *(可选)* | String[] | 迭代类别的id数组。 |

**响应示例：**

```json
{
    "id": "5ecf7b74eaab845a2aa53132",
    "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/sprints/5ecf7b74eaab845a2aa53132",
    "project": {
        "id": "5eb623f6a70571487ea47000",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
        "identifier": "SCR",
        "name": "Scrum项目",
        "type": "scrum",
        "is_archived": 0,
        "is_deleted": 0
    },
    "name": "Sprint 2",
    "status": "pending",
    "assignee": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "start_at": 1589791860,
    "end_at": 1589791860,
    "description": "This is sprint 2",
    "started_at": 1589791860,
    "completed_at": 1589791960,
    "total_story_points": 0,
    "started_story_points": 0,
    "completed_story_points": 0,
    "categories": [
        {
            "id": "676a460a0fd987b7ea320887",
            "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/sprint_categories/676a460a0fd987b7ea320887",
            "name": "Category 1"
        }
    ],
    "created_at": 1676454024,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "updated_at": 1676454024,
    "updated_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    }
}
```

---

### 创建一个迭代分组

**`POST /v1/project/projects/{project_id}/sprint_sections`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `project_id` | String | 项目的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `name` | String | 迭代分组的名称。 |

**响应示例：**

```json
{
    "id": "634f869a0fd987b7ea320833",
    "url": "http://rest_api_root/v1/project/projects/63560f3ad02cbc4f9df91236/sprint_sections/634f869a0fd987b7ea320833",
    "project": {
        "id": "5eb623f6a70571487ea47000",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
        "identifier": "SCR",
        "name": "Scrum项目",
        "type": "scrum",
        "is_archived": 0,
        "is_deleted": 0
    },
    "name": "Section 1"
}
```

---

### 创建一个迭代类别

**`POST /v1/project/projects/{project_id}/sprint_categories`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `project_id` | String | 项目的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `name` | String | 迭代类别的名称。 |
| `section_id` *(可选)* | String | 迭代类别所属迭代分组id。 |

**响应示例：**

```json
{
    "id": "676a460a0fd987b7ea320887",
    "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/sprint_categories/676a460a0fd987b7ea320887",
    "project": {
        "id": "5eb623f6a70571487ea47000",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
        "identifier": "SCR",
        "name": "Scrum项目",
        "type": "scrum",
        "is_archived": 0,
        "is_deleted": 0
    },
    "name": "Category 1",
    "section": {
        "id": "634f869a0fd987b7ea320833",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/sprint_sections/634f869a0fd987b7ea320833",
        "name": "Section 1"
    }
}
```

---

### 删除一个迭代分组

**`DELETE /v1/project/projects/{project_id}/sprint_sections/{section_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `project_id` | String | 项目的id。 |
| `section_id` | String | 迭代分组的id。 |

**响应示例：**

```json
{
    "id": "634f869a0fd987b7ea320834",
    "url": "http://rest_api_root/v1/project/projects/63560f3ad02cbc4f9df91236/sprint_sections/634f869a0fd987b7ea320834",
    "project": {
        "id": "5eb623f6a70571487ea47000",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
        "identifier": "SCR",
        "name": "Scrum项目",
        "type": "scrum",
        "is_archived": 0,
        "is_deleted": 0
    },
    "name": "Section 2"
}
```

---

### 删除一个迭代类别

**`DELETE /v1/project/projects/{project_id}/sprint_categories/{sprint_category_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `project_id` | String | 项目的id。 |
| `sprint_category_id` | String | 迭代类别的id。 |

**响应示例：**

```json
{
    "id": "676a460a0fd987b7ea320888",
    "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/sprint_categories/676a460a0fd987b7ea320888",
    "project": {
        "id": "5eb623f6a70571487ea47000",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
        "identifier": "SCR",
        "name": "Scrum项目",
        "type": "scrum",
        "is_archived": 0,
        "is_deleted": 0
    },
    "name": "Category 2",
    "section": {
        "id": "634f869a0fd987b7ea320833",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/sprint_sections/634f869a0fd987b7ea320833",
        "name": "Section 1"
    }
}
```

---

### 批量创建迭代

**`POST /v1/project/sprints/bulk`**

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `sprints` | Object[] | 需要批量创建的迭代。该参数是一个对象数组（数组内对象不得超过100个）。 |
| `sprints.project_id` | String | 迭代所属项目的id。 |
| `sprints.name` | String | 迭代的名称。 |
| `sprints.start_at` | Number | 迭代开始的时间。 |
| `sprints.end_at` | Number | 迭代结束的时间。 |
| `sprints.assignee_id` | String | 迭代负责人的id。 |
| `sprints.description` *(可选)* | String | 迭代的描述。 |
| `sprints.status` *(可选)* | String | 迭代的状态。 |
| `sprints.category_ids` *(可选)* | String[] | 迭代类别的id列表。 |

**响应示例：**

```json
[
    {
        "state": "success",
        "sprint": {
            "id": "5ecf7b74eaab845a2aa53134",
            "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/sprints/5ecf7b74eaab845a2aa53134",
            "project": {
                "id": "5eb623f6a70571487ea47000",
                "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
                "identifier": "SCR",
                "name": "Scrum项目",
                "type": "scrum",
                "is_archived": 0,
                "is_deleted": 0
            },
            "name": "Sprint 3",
            "status": "pending",
            "assignee": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "start_at": 1589791860,
            "end_at": 1589791860,
            "description": "This is sprint 3",
            "started_at": 1589791860,
            "completed_at": 1589791960,
            "total_story_points": 0,
            "started_story_points": 0,
            "completed_story_points": 0,
            "categories": [
                {
                    "id": "676a460a0fd987b7ea320887",
                    "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/sprint_categories/676a460a0fd987b7ea320887",
                    "name": "Category 1"
                }
            ],
            "created_at": 1676454024,
            "created_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "updated_at": 1676454024,
            "updated_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            }
        }
    }
]
```

---

### 获取迭代分组列表

**`GET /v1/project/projects/{project_id}/sprint_sections`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `project_id` | String | 项目的id。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 2,
    "values": [
        {
            "id": "634f869a0fd987b7ea320833",
            "url": "http://rest_api_root/v1/project/projects/63560f3ad02cbc4f9df91236/sprint_sections/634f869a0fd987b7ea320833",
            "project": {
                "id": "5eb623f6a70571487ea47000",
                "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
                "identifier": "SCR",
                "name": "Scrum项目",
                "type": "scrum",
                "is_archived": 0,
                "is_deleted": 0
            },
            "name": "Section 1"
        },
        {
            "id": "634f869a0fd987b7ea320834",
            "url": "http://rest_api_root/v1/project/projects/63560f3ad02cbc4f9df91236/sprint_sections/634f869a0fd987b7ea320834",
            "project": {
                "id": "5eb623f6a70571487ea47000",
                "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
                "identifier": "SCR",
                "name": "Scrum项目",
                "type": "scrum",
                "is_archived": 0,
                "is_deleted": 0
            },
            "name": "Section 2"
        }
    ]
}
```

---

### 获取迭代列表

**`GET /v1/project/projects/{project_id}/sprints`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `project_id` | String | 项目的id。 |

**查询参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `name` *(可选)* | String | 迭代的名称。 |
| `status` *(可选)* | String | 迭代的状态。 |
| `created_between` *(可选)* | String | 创建时间介于的时间范围，通过','分割起始时间。 |
| `updated_between` *(可选)* | String | 更新时间介于的时间范围，通过','分割起始时间。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "5ecf7b74eaab845a2aa53138",
            "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/sprints/5ecf7b74eaab845a2aa53138",
            "project": {
                "id": "5eb623f6a70571487ea47000",
                "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
                "identifier": "SCR",
                "name": "Scrum项目",
                "type": "scrum",
                "is_archived": 0,
                "is_deleted": 0
            },
            "name": "Sprint 1",
            "status": "completed",
            "assignee": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "start_at": 1589791860,
            "end_at": 1589791860,
            "description": "This is sprint 1",
            "started_at": 1589791860,
            "completed_at": 1589791960,
            "total_story_points": 0,
            "started_story_points": 0,
            "completed_story_points": 0,
            "categories": [
                {
                    "id": "676a460a0fd987b7ea320887",
                    "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/sprint_categories/676a460a0fd987b7ea320887",
                    "name": "Category 1"
                }
            ],
            "created_at": 1676454024,
            "created_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "updated_at": 1676454024,
            "updated_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            }
        }
    ]
}
```

---

### 获取迭代类别列表

**`GET /v1/project/projects/{project_id}/sprint_categories`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `project_id` | String | 项目的id。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 2,
    "values": [
        {
            "id": "676a460a0fd987b7ea320887",
            "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/sprint_categories/676a460a0fd987b7ea320887",
            "project": {
                "id": "5eb623f6a70571487ea47000",
                "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
                "identifier": "SCR",
                "name": "Scrum项目",
                "type": "scrum",
                "is_archived": 0,
                "is_deleted": 0
            },
            "name": "Category 1",
            "section": {
                "id": "634f869a0fd987b7ea320833",
                "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/sprint_sections/634f869a0fd987b7ea320833",
                "name": "Section 1"
            }
        },
        {
            "id": "676a460a0fd987b7ea320888",
            "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/sprint_categories/676a460a0fd987b7ea320888",
            "project": {
                "id": "5eb623f6a70571487ea47000",
                "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
                "identifier": "SCR",
                "name": "Scrum项目",
                "type": "scrum",
                "is_archived": 0,
                "is_deleted": 0
            },
            "name": "Category 2",
            "section": {
                "id": "634f869a0fd987b7ea320833",
                "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/sprint_sections/634f869a0fd987b7ea320833",
                "name": "Section 1"
            }
        }
    ]
}
```

---

### 部分更新一个迭代

**`PATCH /v1/project/projects/{project_id}/sprints/{sprint_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `project_id` | String | 项目的id。 |
| `sprint_id` | String | 迭代的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `name` *(可选)* | String | 迭代的名称。 |
| `start_at` *(可选)* | Number | 迭代开始的时间。 |
| `end_at` *(可选)* | Number | 迭代结束的时间。 |
| `assignee_id` *(可选)* | String | 迭代负责人的id。 |
| `description` *(可选)* | String | 迭代的描述。 |
| `status` *(可选)* | String | 迭代的状态。 |
| `category_ids` *(可选)* | String[] | 迭代类别的id列表。 |

**响应示例：**

```json
{
    "id": "5ecf7b74eaab845a2aa53132",
    "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/sprints/5ecf7b74eaab845a2aa53132",
    "project": {
        "id": "5eb623f6a70571487ea47000",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
        "identifier": "SCR",
        "name": "Scrum项目",
        "type": "scrum",
        "is_archived": 0,
        "is_deleted": 0
    },
    "name": "Sprint 2",
    "status": "in_progress",
    "assignee": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "start_at": 1589791860,
    "end_at": 1589791860,
    "description": "This is sprint 2",
    "started_at": 1589791860,
    "completed_at": 1589791960,
    "total_story_points": 0,
    "started_story_points": 0,
    "completed_story_points": 0,
    "categories": [
        {
            "id": "676a460a0fd987b7ea320887",
            "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/sprint_categories/676a460a0fd987b7ea320887",
            "name": "Category 1"
        }
    ],
    "created_at": 1676454024,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "updated_at": 1676454024,
    "updated_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    }
}
```

---

### 部分更新一个迭代分组

**`PATCH /v1/project/projects/{project_id}/sprint_sections/{section_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `project_id` | String | 项目的id。 |
| `section_id` | String | 迭代分组的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `name` | String | 迭代分组的名称。 |

**响应示例：**

```json
{
    "id": "634f869a0fd987b7ea320833",
    "url": "http://rest_api_root/v1/project/projects/63560f3ad02cbc4f9df91236/sprint_sections/634f869a0fd987b7ea320833",
    "project": {
        "id": "5eb623f6a70571487ea47000",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
        "identifier": "SCR",
        "name": "Scrum项目",
        "type": "scrum",
        "is_archived": 0,
        "is_deleted": 0
    },
    "name": "Section 1"
}
```

---

### 部分更新一个迭代类别

**`PATCH /v1/project/projects/{project_id}/sprint_categories/{sprint_category_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `project_id` | String | 项目的id。 |
| `sprint_category_id` | String | 迭代类别的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `name` *(可选)* | String | 迭代类别的名称。 |
| `section_id` *(可选)* | String | 迭代类别所属迭代分组id。 |

**响应示例：**

```json
{
    "id": "676a460a0fd987b7ea320888",
    "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/sprint_categories/676a460a0fd987b7ea320888",
    "project": {
        "id": "5eb623f6a70571487ea47000",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
        "identifier": "SCR",
        "name": "Scrum项目",
        "type": "scrum",
        "is_archived": 0,
        "is_deleted": 0
    },
    "name": "Category 2",
    "section": {
        "id": "634f869a0fd987b7ea320833",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/sprint_sections/634f869a0fd987b7ea320833",
        "name": "Section 1"
    }
}
```

---

## 个人

### 获取个人信息

**`GET /v1/myself`**

**响应示例：**

```json
{
    "id": "a0417f68e846aae315c85d24643678a9",
    "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
    "name": "john",
    "display_name": "John",
    "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png",
    "email": "terry@pingcode.com",
    "mobile": "15000000000",
    "status": "enabled",
    "preferences": {
        "locale": "zh-cn",
        "timezone": "Asia/Shanghai"
    },
    "permissions": {
        "agile_create_project": true,
        "agile_configuration": true,
        "create_user": true
    }
}
```

---

## 交付

### 环境

企业内实际的部署环境，用于在PingCode中显示各个环境的部署信息。 资源地址： GET https://rest_api_root/v1/release/environments/{env_id}

**全量数据示例：**

```json
{
    "id": "564587fe700d43b81b080123",
    "url": "https://rest_api_root/v1/release/environments/564587fe700d43b81b080123",
    "name": "Production",
    "html_url": "https://your-environment-url"
}
```

**引用数据示例：**

```json
{
    "id": "564587fe700d43b81b080123",
    "url": "https://rest_api_root/v1/release/environments/564587fe700d43b81b080123",
    "name": "Production"
}
```

---

### 部署

企业内实际的部署记录，用于在PingCode中显示部署的详细信息。 资源地址： GET https://rest_api_root/v1/release/deploys/{deploy_id}

**全量数据示例：**

```json
{
    "id": "564587fe700d43b81b080339",
    "url": "https://rest_api_root/v1/release/deploys/564587fe700d43b81b080339",
    "status": "deployed",
    "release_name": "1.1.0",
    "environment": {
        "id": "564587fe700d43b81b080123",
        "url": "https://rest_api_root/v1/release/environments/564587fe700d43b81b080123",
        "name": "Production"
    },
    "release_url": "https://github.com/worktile/ngx-planet/releases/tag/1.1.0",
    "start_at": 1583143467,
    "end_at": 1583143667,
    "duration": 200,
    "work_items": [
        {
            "id": "564587fe700d43b81b080ab8",
            "url": "https://rest_api_root/v1/project/work_items/564587fe700d43b81b080ab8",
            "identifier": "PLM-001",
            "title": "这是一个用户故事",
            "type": "story",
            "start_at": 1583290309,
            "end_at": 1583290347,
            "parent_id": "5edca524cad2fa112b06105c",
            "short_id": "c9WqLmTO",
            "html_url": "https://yctech.pingcode.com/pjm/workitems/c9WqLmTO",
            "properties": {
                "prop_a": "prop_a_value",
                "prop_b": "prop_b_value"
            }
        }
    ]
}
```

**引用数据示例：**

```json
{
    "id": "564587fe700d43b81b080339",
    "url": "https://rest_api_root/v1/release/deploys/564587fe700d43b81b080339",
    "status": "deployed",
    "release_name": "1.1.0"
}
```

---

## 产品

### 创建一个产品

**`POST /v1/ship/products`**

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `name` | String | 产品的名称（不超过32个字符）。 |
| `identifier` | String | 产品的标识。在一个企业中这个标识是唯一的。产品的标识由大写英文字母/数字/下划线/连接线组成（不超过15个字符）。 |
| `description` *(可选)* | String | 产品的描述。 |
| `members` *(可选)* | Object[] | 产品成员的列表。 |
| `members.id` | String | 企业成员或团队的id。 |
| `members.type` | String | 产品成员的类型。 |

**响应示例：**

```json
{
    "id": "6422711c3f12e6c1e46d40e9",
    "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9",
    "identifier": "SLC",
    "name": "示例产品",
    "visibility": "private",
    "scope_type": "user_group",
    "scope_id": "6422711c3f12e6c1e46d40e9",
    "color": "#FA8888",
    "description": "示例产品描述",
    "members": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9/members/a0417f68e846aae315c85d24643678a9",
            "type": "user",
            "user": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            }
        },
        {
            "id": "63c8fb32729dee3334d96af7",
            "url": "https://rest_api_root/v1/wiki/spaces/642fd641209b56920a6c6e5e/members/63c8fb32729dee3334d96af7",
            "type": "user_group",
            "user_group": {
                 "id": "63c8fb32729dee3334d96af7",
                 "url": "https://rest_api_root/v1/directory/groups/63c8fb32729dee3334d96af7",
                 "name": "Open Team"
            }
        }
    ],
    "created_at": 1583290300,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "updated_at": 1583290300,
    "updated_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "is_archived": 0,
    "is_deleted": 0
}
```

---

### 向产品中添加一个成员

**`POST /v1/ship/products/{product_id}/members`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `product_id` | String | 产品的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `member` | Object | 产品的成员。 |
| `member.id` | String | 企业成员或团队的id。 |
| `member.type` | String | 项目成员的类型。 |
| `role_id` *(可选)* | String | 角色的id。 |

**响应示例：**

```json
{
    "id": "a0417f68e846aae315c85d24643678a9",
    "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9/members/a0417f68e846aae315c85d24643678a9",
    "product": {
        "id": "6422711c3f12e6c1e46d40e9",
        "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9",
        "name": "示例产品",
        "identifier": "SLC",
        "is_archived": 0,
        "is_deleted": 0
    },
    "type": "user",
    "user": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "role": {
        "id": "6422711c3f12e6c1e46d40e6",
        "url": "https://rest_api_root/v1/directory/roles/6422711c3f12e6c1e46d40e6",
        "name": "管理员"
    }
}
```

---

### 向产品中添加一个标签

**`POST /v1/ship/products/{product_id}/tags`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `product_id` | String | 产品的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `name` | String | 标签的名称。在一个产品中这个名称是唯一的。 |

**响应示例：**

```json
{
    "id": "63eca881a0a13a3efc8d49f0",
    "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e6/tags/63eca881a0a13a3efc8d49f0",
    "product": {
        "id": "6422711c3f12e6c1e46d40e6",
        "url": "http://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e6",
        "identifier": "SLC",
        "name": "示例产品",
        "is_archived": 0,
        "is_deleted": 0
    },
    "name": "标签-1",
    "color": "#56ABFB"
}
```

---

### 向产品中添加一个需求模块

**`POST /v1/ship/products/{product_id}/suites`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `product_id` | String | 产品的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `name` | String | 模块名称。需求模块为树形结构，同一层次下的名称不能重复。 |
| `type` | String | 模块类型。允许值分别表示子产品和模块。 |
| `parent_id` *(可选)* | String | 父模块的id。 |

**响应示例：**

```json
{
    "id": "63eca881a0a13a3efc8d49f0",
    "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e6/suites/63eca881a0a13a3efc8d49f0",
    "product": {
        "id": "6422711c3f12e6c1e46d40e6",
        "url": "http://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e6",
        "identifier": "SLC",
        "name": "示例产品",
        "is_archived": 0,
        "is_deleted": 0
    },
    "name": "技术支持确认",
    "type": "module",
    "parent": {
        "id": "63bb744414bd13c9def24ce3",
        "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e6/suites/63bb744414bd13c9def24ce3",
        "name": "父级产品模块",
        "type": "module"
    }
}
```

---

### 在产品中移除一个成员

**`DELETE /v1/ship/products/{product_id}/members/{member_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `product_id` | String | 产品的id。 |
| `member_id` | String | 企业成员或团队的id。 |

**响应示例：**

```json
{
    "id": "a0417f68e846aae315c85d24643678a9",
    "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9/members/a0417f68e846aae315c85d24643678a9",
    "product": {
        "id": "6422711c3f12e6c1e46d40e9",
        "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9",
        "name": "示例产品",
        "identifier": "SLC",
        "is_archived": 0,
        "is_deleted": 0
    },
    "type": "user",
    "user": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "role": {
        "id": "6422711c3f12e6c1e46d40e6",
        "url": "https://rest_api_root/v1/directory/roles/6422711c3f12e6c1e46d40e6",
        "name": "管理员"
    }
}
```

---

### 在产品中移除一个标签

**`DELETE /v1/ship/products/{product_id}/tags/{tag_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `product_id` | String | 产品的id。 |
| `tag_id` | String | 标签的id。 |

**响应示例：**

```json
{
    "id": "63eca881a0a13a3efc8d49f0",
    "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e6/tags/63eca881a0a13a3efc8d49f0",
    "product": {
        "id": "6422711c3f12e6c1e46d40e6",
        "url": "http://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e6",
        "identifier": "SLC",
        "name": "示例产品",
        "is_archived": 0,
        "is_deleted": 0
    },
    "name": "技术支持确认",
    "color": "#56ABFB"
}
```

---

### 在产品中移除一个需求模块

**`DELETE /v1/ship/products/{product_id}/suites/{suite_id}`**

请注意，删除一个模块会自动删除其所有的子模块。

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `product_id` | String | 产品的id。 |
| `suite_id` | String | 模块id。 |

**响应示例：**

```json
{
    "id": "63eca881a0a13a3efc8d49f0",
    "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e6/suites/63eca881a0a13a3efc8d49f0",
    "product": {
        "id": "6422711c3f12e6c1e46d40e6",
        "url": "http://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e6",
        "identifier": "SLC",
        "name": "示例产品",
        "is_archived": 0,
        "is_deleted": 0
    },
    "name": "技术支持确认",
    "type": "module",
    "parent": {
        "id": "63bb744414bd13c9def24ce3",
        "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e6/suites/63bb744414bd13c9def24ce3",
        "name": "父级产品模块",
        "type": "module"
    }
}
```

---

### 获取产品中的工单渠道列表

**`GET /v1/ship/products/{product_id}/channels`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `product_id` | String | 产品的id。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "63eca881a0a13a3efc8d49ed",
            "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e6/channels/63eca881a0a13a3efc8d49ed",
            "name": "客户反馈Web渠道",
            "product": {
                "id": "6422711c3f12e6c1e46d40e6",
                "url": "http://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e6",
                "identifier": "SLC",
                "name": "示例产品",
                "is_archived": 0,
                "is_deleted": 0
            },
            "description": "收集客户反馈意见的Web渠道"
        }
    ]
}
```

---

### 获取产品中的工单类型列表

**`GET /v1/ship/products/{product_id}/ticket_types`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `product_id` | String | 产品的id。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "63bb744214bd13c9def24ca9",
            "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9/ticket_types/63bb744214bd13c9def24ca9",
            "product": {
                "id": "6422711c3f12e6c1e46d40e9",
                "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9",
                "name": "自动化",
                "identifier": "FLOW",
                "is_archived": 0,
                "is_deleted": 0
            },
            "ticket_type": {
                "id": "63bb744214bd13c9def24ca9",
                "url": "https://rest_api_root/v1/ship/ticket_types/63bb744214bd13c9def24ca9",
                "name": "需求"
            }
        }
    ]
}
```

---

### 获取产品中的成员列表

**`GET /v1/ship/products/{product_id}/members`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `product_id` | String | 产品的id。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 2,
    "values": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9/members/a0417f68e846aae315c85d24643678a9",
            "type": "user",
            "product": {
                "id": "6422711c3f12e6c1e46d40e9",
                "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9",
                "name": "示例产品",
                "identifier": "SLC",
                "is_archived": 0,
                "is_deleted": 0
             },
            "user": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "role": {
                "id": "6422711c3f12e6c1e46d40e6",
                "url": "https://rest_api_root/v1/directory/roles/6422711c3f12e6c1e46d40e6",
                "name": "管理员"
            }
        },
        {
            "id": "63c8fb32729dee3334d96af7",
            "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9/members/63c8fb32729dee3334d96af7",
            "product": {
                "id": "6422711c3f12e6c1e46d40e9",
                "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9",
                "name": "示例产品",
                "identifier": "SLC",
                "is_archived": 0,
                "is_deleted": 0
             },
            "type": "user_group",
            "user_group": {
                "id": "63c8fb32729dee3334d96af7",
                "url": "https://rest_api_root/v1/directory/groups/63c8fb32729dee3334d96af7",
                "name": "Open Team"
            },
            "role": {
                "id": "6422711c3f12e6c1e46d40e6",
                "url": "https://rest_api_root/v1/directory/roles/6422711c3f12e6c1e46d40e6",
                "name": "管理员"
            }
        }
    ]
}
```

---

### 获取产品中的标签列表

**`GET /v1/ship/products/{product_id}/tags`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `product_id` | String | 产品的id。 |

**响应示例：**

```json
{
     "page_size": 30,
     "page_index": 0,
     "total": 1,
     "values": [
         {
             "id": "63eca881a0a13a3efc8d49f0",
             "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e6/tags/63eca881a0a13a3efc8d49f0",
             "product": {
                 "id": "6422711c3f12e6c1e46d40e6",
                 "url": "http://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e6",
                 "identifier": "SLC",
                 "name": "示例产品",
                 "is_archived": 0,
                 "is_deleted": 0
             },
             "name": "技术支持确认",
             "color": "#56ABFB"
         }
     ]
 }
```

---

### 获取产品中的需求排期列表

**`GET /v1/ship/products/{product_id}/plans`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `product_id` | String | 产品的id。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "63bb744414bd13c9def24ce4",
            "url": "http://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e6/plans/63bb744414bd13c9def24ce4",
            "product": {
                "id": "6422711c3f12e6c1e46d40e6",
                "url": "http://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e6",
                "identifier": "SLC",
                "name": "示例产品",
                "is_archived": 0,
                "is_deleted": 0
            },
            "name": "账号管理",
            "assignee": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "start_at": 1672704000,
            "end_at": 1672963199
        }
    ]
}
```

---

### 获取产品中的需求模块列表

**`GET /v1/ship/products/{product_id}/suites`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `product_id` | String | 产品的id。 |

**响应示例：**

```json
{
     "page_size": 30,
     "page_index": 0,
     "total": 1,
     "values": [
         {
             "id": "63eca881a0a13a3efc8d49f0",
             "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e6/suites/63eca881a0a13a3efc8d49f0",
             "product": {
                 "id": "6422711c3f12e6c1e46d40e6",
                 "url": "http://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e6",
                 "identifier": "SLC",
                 "name": "示例产品",
                 "is_archived": 0,
                 "is_deleted": 0
             },
            "name": "技术支持确认",
            "type": "module",
            "parent": {
                "id": "63bb744414bd13c9def24ce3",
                "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e6/suites/63bb744414bd13c9def24ce3",
                "name": "父级产品模块",
                "type": "module"
            }
         }
     ]
 }
```

---

### 获取产品列表

**`GET /v1/ship/products`**

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "6422711c3f12e6c1e46d40e9",
            "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9",
            "identifier": "SLC",
            "name": "示例产品",
            "visibility": "private",
            "scope_type": "user_group",
            "scope_id": "6422711c3f12e6c1e46d40e9",
            "color": "#FA8888",
            "description": "示例产品描述",
            "members": [
                {
                    "id": "a0417f68e846aae315c85d24643678a9",
                    "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9/members/a0417f68e846aae315c85d24643678a9",
                    "type": "user",
                    "user": {
                        "id": "a0417f68e846aae315c85d24643678a9",
                        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                        "name": "john",
                        "display_name": "John",
                        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
                    }
                },
                {
                     "id": "63c8fb32729dee3334d96af7",
                     "url": "https://rest_api_root/v1/wiki/spaces/642fd641209b56920a6c6e5e/members/63c8fb32729dee3334d96af7",
                     "type": "user_group",
                     "user_group": {
                         "id": "63c8fb32729dee3334d96af7",
                         "url": "https://rest_api_root/v1/directory/groups/63c8fb32729dee3334d96af7",
                         "name": "Open Team"
                     }
                }
            ],
            "created_at": 1583290300,
            "created_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "updated_at": 1583290300,
            "updated_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "is_archived": 0,
            "is_deleted": 0
        }
    ]
}
```

---

### 部分更新一个产品

**`PATCH /v1/ship/products/{product_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `product_id` | String | 产品的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `name` *(可选)* | String | 产品的名称（不超过32个字符）。 |
| `identifier` *(可选)* | String | 产品的标识。在一个企业中这个标识是唯一的。产品的标识由大写英文字母/数字/下划线/连接线组成（不超过15个字符）。 |
| `description` *(可选)* | String | 产品的描述。 |

**响应示例：**

```json
{
    "id": "6422711c3f12e6c1e46d40e9",
    "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9",
    "identifier": "SLC",
    "name": "示例产品",
    "scope_type": "user_group",
    "scope_id": "6422711c3f12e6c1e46d40e9",
    "visibility": "private",
    "color": "#FA8888",
    "description": "示例产品描述",
    "members": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9/members/a0417f68e846aae315c85d24643678a9",
            "type": "user",
            "user": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            }
        },
        {
            "id": "63c8fb32729dee3334d96af7",
            "url": "https://rest_api_root/v1/wiki/spaces/642fd641209b56920a6c6e5e/members/63c8fb32729dee3334d96af7",
            "type": "user_group",
            "user_group": {
                 "id": "63c8fb32729dee3334d96af7",
                 "url": "https://rest_api_root/v1/directory/groups/63c8fb32729dee3334d96af7",
                 "name": "Open Team"
            }
        }
    ],
    "created_at": 1583290300,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "updated_at": 1583293300,
    "updated_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "is_archived": 0,
    "is_deleted": 0
}
```

---

## 产品管理

### 产品

用于查看和管理产品相关的基本信息。 资源地址：{GET} /v1/ship/products/{product_id}

**全量数据示例：**

```json
{
    "id": "6422711c3f12e6c1e46d40e9",
    "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9",
    "identifier": "SLC",
    "name": "示例产品",
    "scope_type": "user_group",
    "scope_id": "6422711c3f12e6c1e46d40e9",
    "visibility": "private",
    "color": "#FA8888",
    "description": "示例产品描述",
    "members": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9/members/a0417f68e846aae315c85d24643678a9",
            "type": "user",
            "user": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            }
        },
        {
            "id": "63c8fb32729dee3334d96af7",
            "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9/members/63c8fb32729dee3334d96af7",
            "type": "user_group",
            "user_group": {
                "id": "63c8fb32729dee3334d96af7",
                "url": "https://rest_api_root/v1/directory/groups/63c8fb32729dee3334d96af7",
                "name": "Open Team"
            }
        }
    ],
    "created_at": 1583290300,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "updated_at": 1583290300,
    "updated_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "is_archived": 0,
    "is_deleted": 0
}
```

**引用数据示例：**

```json
{
    "id": "6422711c3f12e6c1e46d40e9",
    "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9",
    "identifier": "SLC",
    "name": "示例产品",
    "is_archived": 0,
    "is_deleted": 0
}
```

---

### 产品配置中心

用于查看和管理产品相关的配置信息。

---

### 外部用户

用于查看和管理外部用户相关的基本信息。 资源地址：{GET} /v1/ship/products/{product_id}/users/{user_id}

**全量数据示例：**

```json
{
    "id": "64a2b61c3a12e6c2e46d41e9",
    "url": "https://rest_api_root/v1/ship/products/64a2b61c3a12e6c2e46d41e9/users/64a2b61c3a12e6c2e46d41e9",
    "name": "jack",
    "display_name": "Jack",
    "avatar": "https://s3.amazonaws.com/bucket/a46ef40c-e21e-48cf-a579-cace9fba839a_160x160.png",
    "email": "jack@email.com",
    "mobile": null,
    "product": {
        "id": "6422711c3f12e6c1e46d40e9",
        "url": "http://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9",
        "identifier": "SLC",
        "name": "示例产品",
        "is_archived": 0,
        "is_deleted": 0
    },
    "customer": {
        "id": "64dd899e3f6383ba72ec2a01",
        "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9/customers/64dd899e3f6383ba72ec2a01",
        "name": "深圳XX新零售有限公司"
    }
}
```

**引用数据示例：**

```json
{
    "id": "64a2b61c3a12e6c2e46d41e9",
    "url": "https://rest_api_root/v1/ship/products/64a2b61c3a12e6c2e46d41e9/users/64a2b61c3a12e6c2e46d41e9",
    "name": "jack",
    "display_name": "Jack",
    "avatar": "https://s3.amazonaws.com/bucket/a46ef40c-e21e-48cf-a579-cace9fba839a_160x160.png"
}
```

---

### 客户

用于查看和管理客户相关的基本信息。 资源地址：{GET} https://rest_api_root/v1/ship/products/{product_id}/customers/{customer_id}

**全量数据示例：**

```json
{
    "id": "64dd899e3f6383ba72ec2a0d",
    "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e6/customers/64dd899e3f6383ba72ec2a0d",
    "product": {
        "id": "6422711c3f12e6c1e46d40e6",
        "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e6",
        "identifier": "SLC",
        "name": "示例产品",
        "is_archived": 0,
        "is_deleted": 0
    },
    "name": "上海XX新零售有限公司",
    "assignee": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "scale": 200,
    "description": "上海XX新零售有限公司的描述",
    "created_at": 1692240286,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "updated_at": 1692240286,
    "updated_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "is_archived": 0,
    "is_deleted": 0
}
```

**引用数据示例：**

```json
{
    "id": "64dd899e3f6383ba72ec2a0d",
    "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e6/customers/64dd899e3f6383ba72ec2a0d",
    "name": "上海XX新零售有限公司"
}
```

---

### 工单

用于查看和管理工单相关的基本信息。 资源地址：{GET} /v1/ship/tickets/{ticket_id}

**全量数据示例：**

```json
{
    "id": "63eca888a0a13a3efc8d4a43",
    "url": "https://rest_api_root/v1/ship/tickets/63eca888a0a13a3efc8d4a43",
    "product": {
        "id": "6422711c3f12e6c1e46d40e9",
        "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9",
        "identifier": "SLC",
        "name": "示例产品",
        "is_archived": 0,
        "is_deleted": 0
    },
    "identifier": "SLC-T1",
    "title": "希望新增支持第三方账号注册",
    "short_id": "pdMUgQzZ",
    "html_url": "https://yctech.pingcode.com/ship/tickets/pdMUgQzZ",
    "assignee": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "state": {
        "id": "63eca880a0a13a3efc8d49d9",
        "url": "https://rest_api_root/v1/ship/ticket_states/63eca880a0a13a3efc8d49d9",
        "name": "待处理",
        "type": "pending"
    },
    "estimated_at": {
        "from": 1701619200,
        "to": 1702742399,
        "granularity": "day"
    },
    "type": {
        "id": "63eca880a0a13a3efc8d49e0",
        "url": "https://rest_api_root/v1/ship/ticket_types/63eca880a0a13a3efc8d49e0",
        "name": "需求"
    },
    "customer": {
        "id": "63eca881a0a13a3efc8d49fc",
        "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9/customers/63eca881a0a13a3efc8d49fc",
        "name": "北京XX科技有限公司"
    },
    "solution": {
        "id": "62f217ae16e3661a20124330",
        "url": "https://rest_api_root/v1/ship/ticket_solutions/62f217ae16e3661a20124330",
        "name": "进入需求池"
    },
    "priority": {
        "url": "https://rest_api_root/v1/ship/ticket_priorities/5cb9466afda1ce4ca0090004",
        "id": "5cb9466afda1ce4ca0090004",
        "name": "P1"
    },
    "channel": {
        "id": "64550d9ec696b249b5fc607d",
        "url": "https://rest_api_root/v1/ship/channels/64550d9ec696b249b5fc607d",
        "name": "channel-1"
    },
    "description": "<p>希望支持其他更多第三方平台的账号注册，以便用第三方账号登录找回更换了手机号的账号，保障账号安全</p>",
    "properties": {
        "prop_a": "prop_a_value",
        "prop_b": "prop_b_value"
    },
    "tags": [
        {
            "id": "63eca881a0a13a3efc8d49f1",
            "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9/tags/63eca881a0a13a3efc8d49f1",
            "name": "已确认"
        }
    ],
    "participants": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "url": "https://rest_api_root/v1/participants/a0417f68e846aae315c85d24643678a9?principal_type=ticket&principal_id=63eca888a0a13a3efc8d4a43",
            "type": "user",
            "user": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            }
        },
        {
            "id": "63c8fb32729dee3334d96af7",
            "url": "https://rest_api_root/v1/participants/63c8fb32729dee3334d96af7?principal_type=ticket&principal_id=63eca888a0a13a3efc8d4a43",
            "type": "user_group",
            "user_group": {
                "id": "63c8fb32729dee3334d96af7",
                "url": "https://rest_api_root/v1/directory/groups/63c8fb32729dee3334d96af7",
                "name": "Open Team"
            }
        }
    ],
    "public_image_token": "N96GlJ4AMQoBCw9pZQ2EMl-AprLN_V_DYlghupBNkJA",
    "submitted_at": 1676454024,
    "submitted_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "completed_at": 1689579131,
    "completed_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "created_at": 1676454024,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "updated_at": 1676454024,
    "updated_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "is_archived": 0,
    "is_deleted": 0
}
```

**引用数据示例：**

```json
{
    "id": "63eca888a0a13a3efc8d4a43",
    "url": "https://rest_api_root/v1/ship/tickets/63eca888a0a13a3efc8d4a43",
    "identifier": "SLC-T1",
    "title": "希望新增支持第三方账号注册",
    "short_id": "pdMUgQzZ",
    "html_url": "https://yctech.pingcode.com/ship/tickets/pdMUgQzZ"
}
```

---

### 需求

用于查看和管理需求相关的基本信息。 资源地址：{GET} /v1/ship/ideas/{idea_id}

**全量数据示例：**

```json
{
    "id": "64b4d70ba368e6594360ea24",
    "url": "https://rest_api_root/v1/ship/ideas/64b4d70ba368e6594360ea24",
    "product": {
        "id": "6422711c3f12e6c1e46d40e9",
        "url": "http://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9",
        "identifier": "SLC",
        "name": "示例产品",
        "is_archived": 0,
        "is_deleted": 0
    },
    "identifier": "SLC-1",
    "title": "示例需求",
    "short_id": "Ogf1EYey",
    "html_url": "https://yctech.pingcode.com/ship/ideas/Ogf1EYey",
    "assignee": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "state": {
        "id": "63e1bf51898a0be5a2d21b2a",
        "url": "https://rest_api_root/v1/ship/idea_states/63e1bf51898a0be5a2d21b2a",
        "name": "待评审",
        "type": "pending"
    },
    "priority": {
        "id": "5cb9466afda1ce4ca0090005",
        "url": "https://rest_api_root/v1/ship/idea_priorities/5cb9466afda1ce4ca0090005",
        "name": "P0"
    },
    "plan": {
        "id": "63bb744414bd13c9def24ce4",
        "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9/plans/63bb744414bd13c9def24ce4",
        "name": "账号管理"
    },
    "suite": {
        "id": "63bb744414bd13c9def24ce4",
        "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9/suites/63bb744414bd13c9def24ce4",
        "name": "需求模块",
        "type": "module"
    },
    "plan_at": {
        "from": 1690732800,
        "to": 1691337599,
        "granularity": "day"
    },
    "real_at": {
        "from": 1690732800,
        "to": 1691337599,
        "granularity": "day"
    },
    "score": 0,
    "progress": 0.60,
    "description": "这是一段描述",
    "properties": {
        "backlog_from": "5cb7e6e2fda1ce4ca0000001",
        "backlog_type": "5cb7e763fda1ce4ca0010002"
    },
    "participants": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "url": "https://rest_api_root/v1/participants/a0417f68e846aae315c85d24643678a9?principal_type=idea&principal_id=64b4d70ba368e6594360ea24",
            "type": "user",
            "user": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            }
        },
        {
            "id": "63c8fb32729dee3334d96af7",
            "url": "https://rest_api_root/v1/participants/63c8fb32729dee3334d96af7?principal_type=idea&principal_id=64b4d70ba368e6594360ea24",
            "type": "user_group",
            "user_group": {
                "id": "63c8fb32729dee3334d96af7",
                "url": "https://rest_api_root/v1/directory/groups/63c8fb32729dee3334d96af7",
                "name": "Open Team"
            }
        }
    ],
    "public_image_token": "-fkvANQ2dcVECK6Xg45L3kG8VCbOTK8NrNckGkxljQD",
    "completed_at": 1689579131,
    "completed_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "created_at": 1689573131,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "updated_at": 1689579131,
    "updated_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "is_archived": 0,
    "is_deleted": 0
}
```

**引用数据示例：**

```json
{
    "id": "64b4d70ba368e6594360ea24",
    "url": "https://rest_api_root/v1/ship/ideas/64b4d70ba368e6594360ea24",
    "identifier": "SLC-1",
    "title": "示例需求",
    "short_id": "Ogf1EYey",
    "html_url": "https://yctech.pingcode.com/ship/ideas/Ogf1EYey"
}
```

---

## 产品配置中心

### 工单配置

用于查看和管理工单相关的配置信息。

---

### 需求配置

用于查看和管理需求相关的配置信息。

---

## 代码仓库

### 全量更新一个代码仓库

**`PUT /v1/scm/products/{product_id}/repositories/{repository_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `product_id` | String | 代码托管平台的id。 |
| `repository_id` | String | 代码仓库的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `name` | String | 代码仓库的名称。 |
| `full_name` | String | 代码仓库的全称。同一代码托管平台下，代码仓库的全称是唯一的。 |
| `description` *(可选)* | String | 代码仓库的简介。 |
| `is_fork` *(可选)* | Boolean | 是否是fork仓库。 |
| `is_private` *(可选)* | Boolean | 是否是私有仓库。 |
| `owner_name` *(可选)* | String | 代码仓库拥有者的用户名。 |
| `html_url` *(可选)* | String | 代码仓库在代码托管平台上的地址。如果为空，在PingCode中不显示相关跳转链接。 |
| `branches_url` *(可选)* | String | 代码仓库的分支地址，使用{branch}表示分支名。如果为空，在PingCode中不显示相关跳转链接。 |
| `commits_url` *(可选)* | String | 代码仓库的提交地址，使用{sha}表示提交的SHA值。如果为空，在PingCode中不显示相关跳转链接。 |
| `compare_url` *(可选)* | String | 代码仓库的分支对比地址，使用{base}和{head}表示基准分支名和需要进行对比的分支名。如果为空，在PingCode中不显示相关跳转链接。 |
| `pulls_url` *(可选)* | String | 代码仓库的拉取请求地址，使用{number}表示拉取请求的编号。如果为空，在PingCode中不显示相关跳转链接。 |

**响应示例：**

```json
{
    "id": "564587fe700d43b81b080766",
    "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766",
    "product": {
        "id": "564587fe700d43b81b080765",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765",
        "name": "Github",
        "type": "github"
    },
    "name": "ngx-planet",
    "full_name": "worktile/ngx-planet",
    "created_at": 1403018919,
    "owner": {
        "id": "5666aea91f99e33cb7c44964",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/users/5666aea91f99e33cb7c44964",
        "name": "terry"
    },
    "is_fork": false,
    "is_private": false,
    "description": "A powerful, reliable, fully-featured and production ready Micro Frontend library for Angular",
    "html_url": "https://github.com/worktile/ngx-planet",
    "branches_url": "https://github.com/worktile/ngx-planet/tree/{branch}",
    "commits_url": "https://github.com/worktile/ngx-planet/commit/{sha}",
    "compare_url": "https://github.com/worktile/ngx-planet/compare/{base}...{head}",
    "pulls_url": "https://github.com/worktile/ngx-planet/pull/{number}"
}
```

---

### 创建一个代码仓库

**`POST /v1/scm/products/{product_id}/repositories`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `product_id` | String | 代码托管平台的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `name` | String | 代码仓库的名称。 |
| `full_name` | String | 代码仓库的全称。同一代码托管平台下，代码仓库的全称是唯一的。 |
| `description` *(可选)* | String | 代码仓库的简介。 |
| `is_fork` *(可选)* | Boolean | 是否是fork仓库。 |
| `is_private` *(可选)* | Boolean | 是否是私有仓库。 |
| `owner_name` *(可选)* | String | 代码仓库拥有者的用户名。 |
| `html_url` *(可选)* | String | 代码仓库的地址。如果为空，在PingCode中不显示相关跳转链接。 |
| `branches_url` *(可选)* | String | 代码仓库的分支地址，使用{branch}表示分支名。如果为空，在PingCode中不显示相关跳转链接。 |
| `commits_url` *(可选)* | String | 代码仓库的提交地址，使用{sha}表示提交的SHA值。如果为空，在PingCode中不显示相关跳转链接。 |
| `compare_url` *(可选)* | String | 代码仓库的分支对比地址，使用{base}和{head}表示基准分支名和需要进行对比的分支名。如果为空，在PingCode中不显示相关跳转链接。 |
| `pulls_url` *(可选)* | String | 代码仓库的拉取请求地址，使用{number}表示拉取请求的编号。如果为空，在PingCode中不显示相关跳转链接。 |

**响应示例：**

```json
{
    "id": "564587fe700d43b81b080766",
    "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766",
    "product": {
        "id": "564587fe700d43b81b080765",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765",
        "name": "Github",
        "type": "github"
    },
    "name": "ngx-planet",
    "full_name": "worktile/ngx-planet",
    "created_at": 1403018919,
    "owner": {
        "id": "5666aea91f99e33cb7c44964",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/users/5666aea91f99e33cb7c44964",
        "name": "terry"
    },
    "is_fork": false,
    "is_private": false,
    "description": "A powerful, reliable, fully-featured and production ready Micro Frontend library for Angular",
    "html_url": "https://github.com/worktile/ngx-planet",
    "branches_url": "https://github.com/worktile/ngx-planet/tree/{branch}",
    "commits_url": "https://github.com/worktile/ngx-planet/commit/{sha}",
    "compare_url": "https://github.com/worktile/ngx-planet/compare/{base}...{head}",
    "pulls_url": "https://github.com/worktile/ngx-planet/pull/{number}"
}
```

---

### 获取代码仓库列表

**`GET /v1/scm/products/{product_id}/repositories`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `product_id` | String | 代码托管平台的id。 |

**查询参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `full_name` *(可选)* | String | 代码仓库的全称。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "564587fe700d43b81b080766",
            "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766",
            "product": {
                "id": "564587fe700d43b81b080765",
                "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765",
                "name": "Github",
                "type": "github"
            },
            "name": "ngx-planet",
            "full_name": "worktile/ngx-planet",
            "created_at": 1403018919,
            "owner": {
                "id": "5666aea91f99e33cb7c44964",
                "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/users/5666aea91f99e33cb7c44964",
                "name": "terry"
            },
            "is_fork": false,
            "is_private": false,
            "description": "A powerful, reliable, fully-featured and production ready Micro Frontend library for Angular",
            "html_url": "https://github.com/worktile/ngx-planet",
            "branches_url": "https://github.com/worktile/ngx-planet/tree/{branch}",
            "commits_url": "https://github.com/worktile/ngx-planet/commit/{sha}",
            "compare_url": "https://github.com/worktile/ngx-planet/compare/{base}...{head}",
            "pulls_url": "https://github.com/worktile/ngx-planet/pull/{number}"
        }
    ]
}
```

---

### 部分更新一个代码仓库

**`PATCH /v1/scm/products/{product_id}/repositories/{repository_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `product_id` | String | 代码托管平台的id。 |
| `repository_id` | String | 代码仓库的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `name` *(可选)* | String | 代码仓库的名称。 |
| `full_name` *(可选)* | String | 代码仓库的全称。同一代码托管平台下，代码仓库的全称是唯一的。 |
| `description` *(可选)* | String | 代码仓库的简介。 |
| `is_fork` *(可选)* | Boolean | 是否是fork仓库。 |
| `is_private` *(可选)* | Boolean | 是否是私有仓库。 |
| `owner_name` *(可选)* | String | 代码仓库拥有者的用户名。 |
| `html_url` *(可选)* | String | 代码仓库在代码托管平台上的地址。如果为空，在PingCode中不显示相关跳转链接。 |
| `branches_url` *(可选)* | String | 代码仓库的分支地址，使用{branch}表示分支名。如果为空，在PingCode中不显示相关跳转链接。 |
| `commits_url` *(可选)* | String | 代码仓库的提交地址，使用{sha}表示提交的SHA值。如果为空，在PingCode中不显示相关跳转链接。 |
| `compare_url` *(可选)* | String | 代码仓库的分支对比地址，使用{base}和{head}表示基准分支名和需要进行对比的分支名。如果为空，在PingCode中不显示相关跳转链接。 |
| `pulls_url` *(可选)* | String | 代码仓库的拉取请求地址，使用{number}表示拉取请求的编号。如果为空，在PingCode中不显示相关跳转链接。 |

**响应示例：**

```json
{
    "id": "564587fe700d43b81b080766",
    "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766",
    "product": {
        "id": "564587fe700d43b81b080765",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765",
        "name": "Github",
        "type": "github"
    },
    "name": "ngx-planet",
    "full_name": "worktile/ngx-planet",
    "created_at": 1403018919,
    "owner": {
        "id": "5666aea91f99e33cb7c44964",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/users/5666aea91f99e33cb7c44964",
        "name": "terry"
    },
    "is_fork": false,
    "is_private": false,
    "description": "A powerful, reliable, fully-featured and production ready Micro Frontend library for Angular",
    "html_url": "https://github.com/worktile/ngx-planet",
    "branches_url": "https://github.com/worktile/ngx-planet/tree/{branch}",
    "commits_url": "https://github.com/worktile/ngx-planet/commit/{sha}",
    "compare_url": "https://github.com/worktile/ngx-planet/compare/{base}...{head}",
    "pulls_url": "https://github.com/worktile/ngx-planet/pull/{number}"
}
```

---

## 代码

### 代码仓库

代码托管平台内实际的代码仓库，用于在PingCode中显示代码仓库的详细信息。 资源地址： GET https://rest_api_root/v1/scm/products/{product_id}/repositories/{repository_id}

**全量数据示例：**

```json
{
    "id": "564587fe700d43b81b080766",
    "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766",
    "product": {
        "id": "564587fe700d43b81b080765",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765",
        "name": "Github",
        "type": "github"
    },
    "name": "ngx-planet",
    "full_name": "worktile/ngx-planet",
    "created_at": 1403018919,
    "owner": {
        "id": "5666aea91f99e33cb7c44964",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/users/5666aea91f99e33cb7c44964",
        "name": "terry"
    },
    "is_fork": false,
    "is_private": false,
    "description": "A powerful, reliable, fully-featured and production ready Micro Frontend library for Angular",
    "html_url": "https://github.com/worktile/ngx-planet",
    "branches_url": "https://github.com/worktile/ngx-planet/tree/{branch}",
    "commits_url": "https://github.com/worktile/ngx-planet/commit/{sha}",
    "compare_url": "https://github.com/worktile/ngx-planet/compare/{base}...{head}",
    "pulls_url": "https://github.com/worktile/ngx-planet/pull/{number}"
}
```

**引用数据示例：**

```json
{
    "id": "564587fe700d43b81b080766",
    "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766",
    "name": "ngx-planet",
    "full_name": "worktile/ngx-planet",
    "created_at": 1403018919
}
```

---

### 代码分支

代码仓库内实际的分支，用于在PingCode中显示分支的详细信息。 资源地址： GET https://rest_api_root/v1/scm/products/{product_id}/repositories/{repository_id}/branches/{branch_id}

**全量数据示例：**

```json
{
    "id": "564587fe700d43b81b080767",
    "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766/branches/564587fe700d43b81b080767",
    "product": {
        "id": "564587fe700d43b81b080765",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765",
        "name": "Github",
        "type": "github"
    },
    "repository": {
        "id": "564587fe700d43b81b080766",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766",
        "name": "ngx-planet",
        "full_name": "worktile/ngx-planet",
        "created_at": 1403018919
    },
    "name": "terry/#PLM-001",
    "created_at": 1403018919,
    "sender": {
        "id": "5666aea91f99e33cb7c44964",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/users/5666aea91f99e33cb7c44964",
        "name": "terry"
    },
    "is_default": false,
    "work_items": [
        {
            "id": "564587fe700d43b81b080ab8",
            "url": "https://rest_api_root/v1/project/work_items/564587fe700d43b81b080ab8",
            "identifier": "PLM-001",
            "title": "这是一个用户故事",
            "type": "story",
            "start_at": 1583290309,
            "end_at": 1583290347,
            "parent_id": "5edca524cad2fa112b06105c",
            "short_id": "c9WqLmTO",
            "html_url": "https://yctech.pingcode.com/pjm/workitems/c9WqLmTO",
            "properties": {
                "prop_a": "prop_a_value",
                "prop_b": "prop_b_value"
            }
        }
    ]
}
```

**引用数据示例：**

```json
{
    "id": "564587fe700d43b81b080767",
    "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766/branches/564587fe700d43b81b080767",
    "name": "terry/#PLM-001",
    "created_at": 1403018919
}
```

---

### 代码评审

拉取请求实际的代码评审记录，用于在PingCode中显示代码评审的详细信息。 资源地址： GET https://rest_api_root/v1/scm/products/{product_id}/repositories/{repository_id}/pull_requests/{pull_request_id}/reviews/{review_id}

**全量数据示例：**

```json
{
  "id": "524587fe700d43b81b080988",
  "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766/pull_requests/594587fe700d43b81b080789/reviews/524587fe700d43b81b080988",
  "product": {
    "id": "564587fe700d43b81b080765",
    "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765",
    "name": "Github",
    "type": "github"
  },
  "repository": {
    "id": "564587fe700d43b81b080766",
    "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766",
    "name": "ngx-planet",
    "full_name": "worktile/ngx-planet",
    "created_at": 1403018919
  },
  "pull_request": {
    "id": "594587fe700d43b81b080789",
    "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766/pull_requests/594587fe700d43b81b080789",
    "number": 1
  },
  "reviewer": {
    "id": "5999aea91f99e33cb7c44964",
    "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/users/5999aea91f99e33cb7c44964",
    "name": "anytao"
  },
  "status": "approved",
  "description": "Review has approved",
  "submitted_at": 1403014111,
  "html_url": "https://github.com/worktile/ngx-planet/pull/127#pullrequestreview-384383294"
}
```

**引用数据示例：**

```json
{
  "id": "524587fe700d43b81b080988",
  "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766/pull_requests/594587fe700d43b81b080789/reviews/524587fe700d43b81b080988",
}
```

---

### 托管平台

企业内实际的代码托管平台，例如Github或私有部署的Gitlab。 资源地址： GET https://rest_api_root/v1/scm/products/{product_id}

**全量数据示例：**

```json
{
    "id": "564587fe700d43b81b080765",
    "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765",
    "name": "Github",
    "type": "github",
    "description": "Github公有云"
}
```

**引用数据示例：**

```json
{
    "id": "564587fe700d43b81b080765",
    "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765",
    "name": "Github",
    "type": "github"
}
```

---

### 托管平台用户

代码托管平台内实际的用户，用于在PingCode中显示该用户在代码托管平台上的名称、头像以及主页的信息。如果没有手动创建用户，在操作代码仓库、分支、拉取请求时，将自动创建仅包含该用户名的用户。 资源地址： GET https://rest_api_root/v1/scm/products/{product_id}/users/{user_id}

**全量数据示例：**

```json
{
    "id": "5666aea91f99e33cb7c44964",
    "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/users/5666aea91f99e33cb7c44964",
    "product": {
        "id": "564587fe700d43b81b080765",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765",
        "name": "Github",
        "type": "github"
    },
    "name": "terry",
    "display_name": "Terry",
    "html_url": "https://github.com/terrylee",
    "avatar_url": "https://avatars2.githubusercontent.com/u/694592?v=4"
}
```

**引用数据示例：**

```json
{
    "id": "5666aea91f99e33cb7c44964",
    "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/users/5666aea91f99e33cb7c44964",
    "product": {
        "id": "564587fe700d43b81b080765",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765",
        "name": "Github",
        "type": "github"
    },
    "name": "terry"
}
```

---

### 拉取请求

代码仓库内实际的拉取请求，用于在PingCode中显示拉取请求的详细信息。 资源地址： GET https://rest_api_root/v1/scm/products/{product_id}/repositories/{repository_id}/pull_requests/{pull_request_id}

**全量数据示例：**

```json
{
  "id": "594587fe700d43b81b080789",
  "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766/pull_requests/594587fe700d43b81b080789",
  "product": {
    "id": "564587fe700d43b81b080765",
    "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765",
    "name": "Github",
    "type": "github"
  },
  "repository": {
    "id": "564587fe700d43b81b080766",
    "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766",
    "name": "ngx-planet",
    "full_name": "worktile/ngx-planet",
    "created_at": 1403018919
  },
  "title": "fix(doc): #PLM-001 fix document title",
  "number": 1,
  "status": "merged",
  "description": "Please give some great suggestions",
  "author": {
    "id": "5666aea91f99e33cb7c44964",
    "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/users/5666aea91f99e33cb7c44964",
    "name": "terry"
  },
  "source_branch": {
    "id": "564587fe700d43b81b080767",
    "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766/branches/564587fe700d43b81b080767",
    "name": "terry/#PLM-001",
    "create_at": 1403018919
  },
  "target_branch": {
    "id": "564587fe700d43b81b080776",
    "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766/branches/564587fe700d43b81b080776",
    "name": "develop",
    "create_at": 1402018919
  },
  "created_at": 1403014000,
  "merged_at": 1473018919,
  "merged_commit_sha": "96a024347146ebdc5f481f45e6e6871e0c43af5f",
  "merged_by": {
    "id": "5666aea91f99e33cb7c44964",
    "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/users/5666aea91f99e33cb7c44964",
    "name": "terry"
  },
  "comments_count": 2,
  "review_comments_count": 2,
  "commits_count": 2,
  "additions_count": 0,
  "deletions_count": 0,
  "changed_files_count": 3,
  "work_items": [
    {
      "id": "564587fe700d43b81b080ab8",
      "url": "https://rest_api_root/v1/project/work_items/564587fe700d43b81b080ab8",
      "identifier": "PLM-001",
      "title": "这是一个用户故事",
      "type": "story",
      "start_at": 1583290309,
      "end_at": 1583290347,
      "parent_id": "5edca524cad2fa112b06105c",
      "short_id": "c9WqLmTO",
      "html_url": "https://yctech.pingcode.com/pjm/workitems/c9WqLmTO",
      "properties": {
        "prop_a": "prop_a_value",
        "prop_b": "prop_b_value"
      }
    }
  ]
}
```

**引用数据示例：**

```json
{
  "id": "594587fe700d43b81b080789",
  "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766/pull_requests/594587fe700d43b81b080789",
  "number": 1
}
```

---

### 提交

实际的代码提交记录，用于在PingCode中显示提交的详细信息。提交并不会自动和代码仓库关联，需要通过提交引用与之关联。 资源地址： GET https://rest_api_root/v1/scm/commits/{commit_id_or_sha}

**全量数据示例：**

```json
{
    "id": "5e3bb2128cfda459bbafa3fb",
    "url": "https://rest_api_root/v1/scm/commits/5e3bb2128cfda459bbafa3fb",
    "sha": "96a024347146ebdc5f481f45e6e6871e0c43af5f",
    "message": "feat(scope): #PLM-001 initialization code structure",
    "committer_name": "terry",
    "committed_at": 1403018919,
    "tree_id": "1bf8989985e70389c07daa5052464a9c6f4896bb",
    "files_added": [
        "index.ts"
    ],
    "files_removed": [
        "utilities.ts"
    ],
    "files_modified": [
        "README.md"
    ],
    "file_changed_count": 3,
    "work_items": [
        {
            "id": "564587fe700d43b81b080ab8",
            "url": "https://rest_api_root/v1/project/work_items/564587fe700d43b81b080ab8",
            "identifier": "PLM-001",
            "title": "这是一个用户故事",
            "type": "story",
            "start_at": 1583290309,
            "end_at": 1583290347,
            "parent_id": "5edca524cad2fa112b06105c",
            "short_id": "c9WqLmTO",
            "html_url": "https://yctech.pingcode.com/pjm/workitems/c9WqLmTO",
            "properties": {
                "prop_a": "prop_a_value",
                "prop_b": "prop_b_value"
            }
        }
    ]
}
```

**引用数据示例：**

```json
{
    "id": "5e3bb2128cfda459bbafa3fb",
    "url": "https://rest_api_root/v1/scm/commits/5e3bb2128cfda459bbafa3fb",
    "sha": "96a024347146ebdc5f481f45e6e6871e0c43af5f",
    "message": "feat(scope): #PLM-001 initialization code structure",
    "committer_name": "terry",
    "committed_at": 1403018919
}
```

---

### 提交引用

提交引用是提交与分支的一种关联关系，一个提交可以与多个分支关联，一个分支也可以与多个提交关联。 当提交与分支关联后，提交会自动与由此分支发起的拉取请求关联，当拉取请求合并后，这些关联的提交将自动被标记为“已合并”状态。 资源地址： GET https://rest_api_root/v1/scm/products/{product_id}/repositories/{repository_id}/refs/{ref_id}

**全量数据示例：**

```json
{
    "id": "5e451b7dd704c212f7de8b4f",
    "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766/refs/5e451b7dd704c212f7de8b4f",
    "product": {
        "id": "564587fe700d43b81b080765",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765",
        "name": "Github",
        "type": "github"
    },
    "repository": {
        "id": "564587fe700d43b81b080766",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766",
        "name": "ngx-planet",
        "full_name": "worktile/ngx-planet",
        "created_at": 1403018919
    },
    "commit": {
        "id": "5e3bb2128cfda459bbafa3fb",
        "url": "https://rest_api_root/v1/scm/commits/5e3bb2128cfda459bbafa3fb",
        "sha": "96a024347146ebdc5f481f45e6e6871e0c43af5f",
        "message": "feat(scope): #PLM-001 initialization code structure",
        "committer_name": "terry",
        "committed_at": 1403018919
    },
    "meta": {
        "id": "564587fe700d43b81b080767",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766/branches/564587fe700d43b81b080767",
        "name": "terry/#PLM-001",
        "type": "branch"
    }
}
```

**引用数据示例：**

```json
{
    "id": "5e451b7dd704c212f7de8b4f",
    "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766/refs/5e451b7dd704c212f7de8b4f"
}
```

---

## 代码分支

### 创建一个代码分支

**`POST /v1/scm/products/{product_id}/repositories/{repository_id}/branches`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `product_id` | String | 代码托管平台的id。 |
| `repository_id` | String | 代码仓库的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `name` | String | 分支的名称。同一代码仓库下，分支的名称是唯一的。 |
| `sender_name` | String | 分支创建者的用户名。 |
| `is_default` *(可选)* | Boolean | 是否设置为默认分支。当创建分支时，如果当前仓库的分支数为0，则新创建的分支会自动设置为该仓库的默认分支。如果创建分支时设置了该分支为默认分支，并且仓库已经有默认分支存在，那么其他分支将被取消默认属性，而该分支将被设置为新的默认分支。 |
| `work_item_identifiers` *(可选)* | String[] | PingCode工作项编号的列表。通过该参数将分支与一个或多个工作项关联，分支和工作项关联后，分支下所有的提交都会和该工作项关联。 |

**响应示例：**

```json
{
    "id": "564587fe700d43b81b080767",
    "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766/branches/564587fe700d43b81b080767",
    "product": {
        "id": "564587fe700d43b81b080765",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765",
        "name": "Github",
        "type": "github"
    },
    "repository": {
        "id": "564587fe700d43b81b080766",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766",
        "name": "ngx-planet",
        "full_name": "worktile/ngx-planet",
        "created_at": 1403018919
    },
    "name": "terry/#PLM-001",
    "created_at": 1403018919,
    "sender": {
        "id": "5666aea91f99e33cb7c44964",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/users/5666aea91f99e33cb7c44964",
        "name": "terry"
    },
    "is_default": true,
    "work_items": [
        {
            "id": "564587fe700d43b81b080ab8",
            "url": "https://rest_api_root/v1/project/work_items/564587fe700d43b81b080ab8",
            "identifier": "PLM-001",
            "title": "这是一个用户故事",
            "type": "story",
            "start_at": 1583290309,
            "end_at": 1583290347,
            "parent_id": "5edca524cad2fa112b06105c",
            "short_id": "c9WqLmTO",
            "html_url": "https://yctech.pingcode.com/pjm/workitems/c9WqLmTO",
            "properties": {
                "prop_a": "prop_a_value",
                "prop_b": "prop_b_value"
            }
        }
    ]
}
```

---

### 删除一个代码分支

**`DELETE /v1/scm/products/{product_id}/repositories/{repository_id}/branches/{branch_id}`**

删除分支后，不会移除该分支和工作项的关联历史，在关联历史中，依然可以查询到删除的分支。默认分支不能被删除。

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `product_id` | String | 代码托管平台的id。 |
| `repository_id` | String | 代码仓库的id。 |
| `branch_id` | String | 分支的id。 |

**响应示例：**

```json
{
    "id": "564587fe700d43b81b080768",
    "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766/branches/564587fe700d43b81b080768",
    "product": {
        "id": "564587fe700d43b81b080765",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765",
        "name": "Github",
        "type": "github"
    },
    "repository": {
        "id": "564587fe700d43b81b080766",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766",
        "name": "ngx-planet",
        "full_name": "worktile/ngx-planet",
        "created_at": 1403018919
    },
    "name": "terry/#PLM-001",
    "created_at": 1403018919,
    "sender": {
        "id": "5666aea91f99e33cb7c44964",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/users/5666aea91f99e33cb7c44964",
        "name": "terry"
    },
    "is_default": false,
    "work_items": [
        {
            "id": "564587fe700d43b81b080ab8",
            "url": "https://rest_api_root/v1/project/work_items/564587fe700d43b81b080ab8",
            "identifier": "PLM-001",
            "title": "这是一个用户故事",
            "type": "story",
            "start_at": 1583290309,
            "end_at": 1583290347,
            "parent_id": "5edca524cad2fa112b06105c",
            "short_id": "c9WqLmTO",
            "html_url": "https://yctech.pingcode.com/pjm/workitems/c9WqLmTO",
            "properties": {
                "prop_a": "prop_a_value",
                "prop_b": "prop_b_value"
            }
        }
    ]
}
```

---

### 获取代码分支列表

**`GET /v1/scm/products/{product_id}/repositories/{repository_id}/branches`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `product_id` | String | 代码托管平台的id。 |
| `repository_id` | String | 代码仓库的id。 |

**查询参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `name` *(可选)* | String | 分支的名称。 |
| `work_item_id` *(可选)* | String | 工作项的id。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "564587fe700d43b81b080767",
            "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766/branches/564587fe700d43b81b080767",
            "product": {
                "id": "564587fe700d43b81b080765",
                "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765",
                "name": "Github",
                "type": "github"
            },
            "repository": {
                "id": "564587fe700d43b81b080766",
                "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766",
                "name": "ngx-planet",
                "full_name": "worktile/ngx-planet",
                "created_at": 1403018919
            },
            "name": "terry/#PLM-001",
            "created_at": 1403018919,
            "sender": {
                "id": "5666aea91f99e33cb7c44964",
                "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/users/5666aea91f99e33cb7c44964",
                "name": "terry"
            },
            "is_default": false,
            "work_items": [
                {
                    "id": "564587fe700d43b81b080ab8",
                    "url": "https://rest_api_root/v1/project/work_items/564587fe700d43b81b080ab8",
                    "identifier": "PLM-001",
                    "title": "这是一个用户故事",
                    "type": "story",
                    "start_at": 1583290309,
                    "end_at": 1583290347,
                    "parent_id": "5edca524cad2fa112b06105c",
                    "short_id": "c9WqLmTO",
                    "html_url": "https://yctech.pingcode.com/pjm/workitems/c9WqLmTO",
                    "properties": {
                        "prop_a": "prop_a_value",
                        "prop_b": "prop_b_value"
                    }
                }
            ]
        }
    ]
}
```

---

### 部分更新一个代码分支

**`PATCH /v1/scm/products/{product_id}/repositories/{repository_id}/branches/{branch_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `product_id` | String | 代码托管平台的id。 |
| `repository_id` | String | 代码仓库的id。 |
| `branch_id` | String | 分支的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `is_default` *(可选)* | Boolean | 是否设置为默认分支。该值只能是true，设置该分支为默认分支后将取消其他分支的默认属性。 |
| `work_item_identifiers` *(可选)* | String[] | PingCode工作项编号的列表。通过该参数将分支与一个或多个工作项关联，分支和工作项关联后，分支下所有的提交都会和该工作项关联。 |

**响应示例：**

```json
{
    "id": "564587fe700d43b81b080767",
    "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766/branches/564587fe700d43b81b080767",
    "product": {
        "id": "564587fe700d43b81b080765",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765",
        "name": "Github",
        "type": "github"
    },
    "repository": {
        "id": "564587fe700d43b81b080766",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766",
        "name": "ngx-planet",
        "full_name": "worktile/ngx-planet",
        "created_at": 1403018919
    },
    "name": "terry/#PLM-001",
    "created_at": 1403018919,
    "sender": {
        "id": "5666aea91f99e33cb7c44964",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/users/5666aea91f99e33cb7c44964",
        "name": "terry"
    },
    "is_default": true,
    "work_items": [
        {
            "id": "564587fe700d43b81b080ab8",
            "url": "https://rest_api_root/v1/project/work_items/564587fe700d43b81b080ab8",
            "identifier": "PLM-001",
            "title": "这是一个用户故事",
            "type": "story",
            "start_at": 1583290309,
            "end_at": 1583290347,
            "parent_id": "5edca524cad2fa112b06105c",
            "short_id": "c9WqLmTO",
            "html_url": "https://yctech.pingcode.com/pjm/workitems/c9WqLmTO",
            "properties": {
                "prop_a": "prop_a_value",
                "prop_b": "prop_b_value"
            }
        }
    ]
}
```

---

## 代码评审

### 全量更新一个代码评审

**`PUT /v1/scm/products/{product_id}/repositories/{repository_id}/pull_requests/{pull_request_id}/reviews/{review_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `product_id` | String | 代码托管平台的id。 |
| `repository_id` | String | 代码仓库的id。 |
| `pull_request_id` | String | 拉取请求的id。 |
| `review_id` | String | 代码评审的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `reviewer_name` | String | 评审人的用户名。 |
| `status` | String | 代码评审的状态。 |
| `submitted_at` | Number | 提交的时间。 |
| `description` *(可选)* | String | 代码评审的描述。 |
| `html_url` *(可选)* | String | 代码评审的地址。如果为空，在PingCode中不显示相关跳转链接。 |

**响应示例：**

```json
{
   "id": "524587fe700d43b81b080988",
   "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766/pull_requests/594587fe700d43b81b080789/reviews/524587fe700d43b81b080988",
   "product": {
      "id": "564587fe700d43b81b080765",
      "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765",
      "name": "Github",
      "type": "github"
   },
   "repository": {
      "id": "564587fe700d43b81b080766",
      "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766",
      "name": "ngx-planet",
      "full_name": "worktile/ngx-planet",
      "created_at": 1403018919
   },
   "pull_request": {
      "id": "594587fe700d43b81b080789",
      "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766/pull_requests/594587fe700d43b81b080789",
      "number": 1
   },
   "reviewer": {
      "id": "5999aea91f99e33cb7c44964",
      "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/users/5999aea91f99e33cb7c44964",
      "name": "anytao"
   },
   "status": "approved",
   "description": "Review has approved",
   "submitted_at": 1403014111,
   "html_url": "https://github.com/worktile/ngx-planet/pull/127#pullrequestreview-384383294"
}
```

---

### 创建一个代码评审

**`POST /v1/scm/products/{product_id}/repositories/{repository_id}/pull_requests/{pull_request_id}/reviews`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `product_id` | String | 代码托管平台的id。 |
| `repository_id` | String | 代码仓库的id。 |
| `pull_request_id` | String | 拉取请求的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `status` | String | 代码评审的状态。 |
| `reviewer_name` | String | 评审人的用户名。 |
| `description` *(可选)* | String | 代码评审的描述。 |
| `submitted_at` | Number | 提交的时间。 |
| `html_url` *(可选)* | String | 代码评审的地址。如果为空，在PingCode中不显示相关跳转链接。 |

**响应示例：**

```json
{
   "id": "524587fe700d43b81b080988",
   "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766/pull_requests/594587fe700d43b81b080789/reviews/524587fe700d43b81b080988",
   "product": {
      "id": "564587fe700d43b81b080765",
      "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765",
      "name": "Github",
      "type": "github"
   },
   "repository": {
      "id": "564587fe700d43b81b080766",
      "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766",
      "name": "ngx-planet",
      "full_name": "worktile/ngx-planet",
      "created_at": 1403018919
   },
   "pull_request": {
      "id": "594587fe700d43b81b080789",
      "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766/pull_requests/594587fe700d43b81b080789",
      "number": 1
   },
   "reviewer": {
      "id": "5999aea91f99e33cb7c44964",
      "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/users/5999aea91f99e33cb7c44964",
      "name": "anytao"
   },
   "status": "approved",
   "description": "Review has approved",
   "submitted_at": 1403014111,
   "html_url": "https://github.com/worktile/ngx-planet/pull/127#pullrequestreview-384383294"
}
```

---

### 获取代码评审列表

**`GET /v1/scm/products/{product_id}/repositories/{repository_id}/pull_requests/{pull_request_id}/reviews`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `product_id` | String | 代码托管平台的id。 |
| `repository_id` | String | 代码仓库的id。 |
| `pull_request_id` | String | 拉取请求的id。 |

**响应示例：**

```json
{
   "page_size": 30,
   "page_index": 0,
   "total": 1,
   "values": [
      {
         "id": "524587fe700d43b81b080988",
         "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766/pull_requests/594587fe700d43b81b080789/reviews/524587fe700d43b81b080988",
         "product": {
            "id": "564587fe700d43b81b080765",
            "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765",
            "name": "Github",
            "type": "github"
         },
         "repository": {
            "id": "564587fe700d43b81b080766",
            "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766",
            "name": "ngx-planet",
            "full_name": "worktile/ngx-planet",
            "created_at": 1403018919
         },
         "pull_request": {
            "id": "594587fe700d43b81b080789",
            "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766/pull_requests/594587fe700d43b81b080789",
            "number": 1
         },
         "reviewer": {
            "id": "5999aea91f99e33cb7c44964",
            "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/users/5999aea91f99e33cb7c44964",
            "name": "anytao"
         },
         "status": "approved",
         "description": "Review has approved",
         "submitted_at": 1403014111,
         "html_url": "https://github.com/worktile/ngx-planet/pull/127#pullrequestreview-384383294"
      }
   ]
}
```

---

### 部分更新一个代码评审

**`PATCH /v1/scm/products/{product_id}/repositories/{repository_id}/pull_requests/{pull_request_id}/reviews/{review_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `product_id` | String | 代码托管平台的id。 |
| `repository_id` | String | 代码仓库的id。 |
| `pull_request_id` | String | 拉取请求的id。 |
| `review_id` | String | 代码评审的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `reviewer_name` *(可选)* | String | 评审人的用户名。 |
| `status` *(可选)* | String | 代码评审的状态。 |
| `description` *(可选)* | String | 代码评审的描述。 |
| `submitted_at` *(可选)* | Number | 提交的时间。 |
| `html_url` *(可选)* | String | 代码评审的地址。如果为空，在PingCode中不显示相关跳转链接。 |

**响应示例：**

```json
{
   "id": "524587fe700d43b81b080988",
   "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766/pull_requests/594587fe700d43b81b080789/reviews/524587fe700d43b81b080988",
   "product": {
      "id": "564587fe700d43b81b080765",
      "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765",
      "name": "Github",
      "type": "github"
   },
   "repository": {
      "id": "564587fe700d43b81b080766",
      "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766",
      "name": "ngx-planet",
      "full_name": "worktile/ngx-planet",
      "created_at": 1403018919
   },
   "pull_request": {
      "id": "594587fe700d43b81b080789",
      "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766/pull_requests/594587fe700d43b81b080789",
      "number": 1
   },
   "reviewer": {
      "id": "5999aea91f99e33cb7c44964",
      "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/users/5999aea91f99e33cb7c44964",
      "name": "anytao"
   },
   "status": "approved",
   "description": "Review has approved",
   "submitted_at": 1403014111,
   "html_url": "https://github.com/worktile/ngx-planet/pull/127#pullrequestreview-384383294"
}
```

---

## 企业成员

### 创建一个企业成员

**`POST /v1/directory/users`**

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `name` | String | 企业成员的名称，在一个企业中这个名称是唯一的。 |
| `display_name` | String | 企业成员的显示名。 |
| `email` *(可选)* | String | 企业成员的邮箱地址，在一个企业中这个邮箱地址是唯一的，邮箱地址和手机号其中一个必填。 |
| `mobile` *(可选)* | String | 企业成员的手机号，在一个企业中这个手机号是唯一，邮箱地址和手机号其中一个必填。 |
| `password` *(可选)* | String | 企业成员的密码，长度为6～200的字符串(包含6和200)。 |
| `department_id` *(可选)* | String | 企业成员的部门id。 |
| `job_id` *(可选)* | String | 企业成员的职位id。 |
| `employee_number` *(可选)* | String | 企业成员的工号。 |

**响应示例：**

```json
{
    "id": "a0417f68e846aae315c85d24643678a9",
    "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
    "name": "john",
    "display_name": "John",
    "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png",
    "department": {
        "id": "6422711c3f12e6c1e46d40e6",
        "url": "https://rest_api_root/v1/directory/departments/6422711c3f12e6c1e46d40e6",
        "name": "技术支持"
    },
    "job": {
        "id": "6440c881c56f557eb1aff6e5",
        "url": "https://rest_api_root/v1/directory/jobs/6440c881c56f557eb1aff6e5",
        "name": "后端工程师"
    },
    "email": "john@email.com",
    "mobile": "15000000000",
    "status": "init",
    "employee_number": "zxv"
}
```

---

### 批量更新企业成员属性

**`PATCH /v1/directory/users/bulk`**

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `user_ids` | String[] | 企业成员的id数组，不能包含自己和团队拥有者。 |
| `property_name` | String | 需要更新的企业成员属性名，目前仅支持：status（可选值为：enabled、disabled） |
| `property_value` | String | 需要更新的企业成员属性值。 |

**响应示例：**

```json
 [
    {
        "state": "success",
        "user_id": "a0417f68e846aae315c85d24643678a9"
    },
    {
        "state": "failure",
        "user_id": "a0417f68e846aae315c85d24643678a8",
        "message": "failure reason.."
    }
]
```

---

### 获取企业成员列表

**`GET /v1/directory/users`**

**查询参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `name` *(可选)* | String | 企业成员的名称，在一个企业中这个名称是唯一的。 |
| `keywords` *(可选)* | String | 关键词模糊搜索，支持姓名、用户名。 |
| `emails` *(可选)* | String | 企业成员的邮箱地址，使用','分割，最多只能20个。 |
| `mobiles` *(可选)* | String | 企业成员的手机号，使用','分割，最多只能20个。 |
| `department_ids` *(可选)* | String | 企业成员的部门id，使用','分割，最多只能20个。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
            "name": "john",
            "display_name": "John",
            "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png",
            "department": {
                "id": "6422711c3f12e6c1e46d40e6",
                "url": "https://rest_api_root/v1/directory/departments/6422711c3f12e6c1e46d40e6",
                "name": "技术支持"
            },
            "job": {
                "id": "6440c881c56f557eb1aff6e5",
                "url": "https://rest_api_root/v1/directory/jobs/6440c881c56f557eb1aff6e5",
                "name": "后端工程师"
            },
            "email": "john@email.com",
            "mobile": "15000000000",
            "status": "enabled",
            "employee_number": "zxv"
        }
    ]
}
```

---

### 部分更新一个企业成员

**`PATCH /v1/directory/users/{user_id}`**

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `name` *(可选)* | String | 企业成员的名称，在一个企业中这个名称是唯一的。 |
| `display_name` *(可选)* | String | 企业成员的显示名。 |
| `email` *(可选)* | String | 企业成员的邮箱地址，在一个企业中这个邮箱地址是唯一的。 |
| `mobile` *(可选)* | String | 企业成员的手机号，在一个企业中这个手机号是唯一的。 |
| `status` *(可选)* | String | 企业成员的状态。 |
| `employee_number` *(可选)* | String | 企业成员的工号。 |
| `department_id` *(可选)* | String | 企业成员的部门id。 |
| `job_id` *(可选)* | String | 企业成员的职位id。 |

**响应示例：**

```json
{
    "id": "a0417f68e846aae315c85d24643678a9",
    "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
    "name": "john",
    "display_name": "John",
    "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png",
    "department": {
        "id": "6422711c3f12e6c1e46d40e6",
        "url": "https://rest_api_root/v1/directory/departments/6422711c3f12e6c1e46d40e6",
        "name": "技术支持"
    },
    "job": {
        "id": "6440c881c56f557eb1aff6e5",
        "url": "https://rest_api_root/v1/directory/jobs/6440c881c56f557eb1aff6e5",
        "name": "后端工程师"
    },
    "email": "john@email.com",
    "mobile": "15000000000",
    "status": "enabled",
    "employee_number": "zxv"
}
```

---

## 企业

### 获取企业信息

**`GET /v1/directory/team`**

**响应示例：**

```json
{
    "id": "56ba35de87ad7153c2062f65",
    "url": "https://rest_api_root/v1/directory/team",
    "name": "YCtech",
    "secondary_domain": "yctech"
}
```

---

## 全局

### 个人

用于查看和管理个人的基本信息。

---

### 安全

---

### 组织

---

### 组织

---

### 组织

---

### 组织

---

### 组织

---

### 组织

---

### 通用

---

## 关注人

### 添加一个关注人

**`POST /v1/participants`**

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `principal_type` | String | 关注人主体的类型。 |
| `principal_id` *(可选)* | String | 关注人主体的id。 |
| `review_id` *(可选)* | String | 关注人评审主体的id。principal_id和review_id至少存在一个，若同时存在，则忽略review_id。 |
| `type` | String | 关注人的类型。 |
| `participant_id` | String | 关注人的id。用户的id或者团队的id。 |

**响应示例（工作项）：**

```json
{
    "id": "a0417f68e846aae315c85d24643678a9",
    "url": "https://rest_api_root/v1/participants/a0417f68e846aae315c85d24643678a9?principal_type=work_item&principal_id=63e1bf51760505c8795ebccc",
    "type": "user",
    "user": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    }
}
```

**响应示例（产品需求评审）：**

```json
{
    "id": "a0417f68e846aae315c85d24643678a9",
    "url": "https://rest_api_root/v1/participants/a0417f68e846aae315c85d24643678a9?principal_type=idea&review_id=6f168f764eba01a5278b87cd",
    "type": "user",
    "user": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    }
}
```

---

### 移除一个关注人

**`DELETE /v1/participants/{participant_id}?principal_type={principal_type}&principal_id={principal_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `participant_id` | String | 关注人的id。用户的id或者团队的id。 |

**查询参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `principal_type` | String | 关注人主体的类型。 |
| `principal_id` *(可选)* | String | 关注人主体的id。 |
| `review_id` *(可选)* | String | 注人评审主体的id。principal_id和review_id至少存在一个，若同时存在，则忽略review_id。 |

**响应示例：**

```json
{
    "id": "a0417f68e846aae315c85d24643678a9",
    "url": "https://rest_api_root/v1/participants/a0417f68e846aae315c85d24643678a9?principal_type=work_item&principal_id=63e1bf51760505c8795ebccc",
    "type": "user",
    "user": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    }
}
```

---

### 获取关注人列表

**`GET /v1/participants?principal_type={principal_type}&principal_id={principal_id}`**

**查询参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `principal_type` | String | 关注人主体的类型。 |
| `principal_id` *(可选)* | String | 关注人主体的id。 |
| `review_id` *(可选)* | String | 关注人评审主体的id。principal_id和review_id至少存在一个，若同时存在，则忽略review_id。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 2,
    "values": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "url": "https://rest_api_root/v1/participants/a0417f68e846aae315c85d24643678a9?principal_type=work_item&principal_id=63e1bf51760505c8795ebccc",
            "type": "user",
            "user": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            }
        },
        {
            "id": "63c8fb32729dee3334d96af7",
            "url": "https://rest_api_root/v1/participants/63c8fb32729dee3334d96af7?principal_type=work_item&principal_id=63e1bf51760505c8795ebccc",
            "type": "user_group",
            "user_group": {
                "id": "63c8fb32729dee3334d96af7",
                "url": "https://rest_api_root/v1/directory/groups/63c8fb32729dee3334d96af7",
                "name": "Open Team"
            }
        }
    ]
}
```

---

## 关联

### 创建一个关联

**`POST /v1/relations`**

**请求参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `principal_type` | String | 关联主体的类型。关联主体的类型和关联的目标类型需要搭配使用，比如： 需求关联工单，principal_type为idea，target_type为ticket； 需求关联工作项，principal_type为idea，target_type为work_item； 需求关联测试用例，principal_type为idea，target_type为test_case； 需求关联需求，principal_type为idea，target_type为idea； 需求关联页面，principal_type为idea，target_type为page； 工单关联需求，principal_type为ticket，target_type为idea； 工单关联工作项，principal_type为ticket，target_type为work_item； 工单关联工单，principal_type为ticket，target_type为ticket； 工单关联页面，principal_type为ticket，target_type为page； 工作项关联测试用例，principal_type为work_item，target_type为test_case； 工作项关联需求，principal_type为work_item，target_type为idea； 工作项关联工单，principal_type为work_item，target_type为ticket； 工作项关联页面，principal_type为work_item，target_type为page； 测试计划关联缺陷，principal_type为test_plan，target_type为work_item； 执行用例关联缺陷，principal_type为test_run，target_type为work_item； 测试用例关联需求，principal_type为test_case，target_type为idea； 测试用例关联工作项，principal_type为test_case，target_type为work_item； 测试用例关联页面，principal_type为test_case，target_type为page； 页面关联需求，暂不开放； 页面关联工单，暂不开放； 页面关联工作项，暂不开放； 页面关联测试用例，暂不开放； |
| `principal_id` | String | 关联主体的id。 |
| `target_type` | String | 关联的目标类型。 |
| `target_id` | String | 关联的目标id。 |

**响应示例：**

```json
{
    "id": "fa1125cb06305edca524cad2",
    "url": "https://rest_api_root/v1/relations/fa1125cb06305edca524cad2",
    "principal_type": "test_run",
    "principal": {
        "id": "547000eb6a70571487623fea",
        "url": "https://rest_api_root/v1/testhub/runs/547000eb6a70571487623fea",
        "status": "failure",
        "short_id": "Aq1u38yX",
        "html_url": "https://yctech.pingcode.com/testhub/runs/Aq1u38yX"
    },
    "target_type": "work_item",
    "target": {
        "id": "5edca524cad2fa1125cb0630",
        "url": "https://rest_api_root/v1/project/work_items/5edca524cad2fa1125cb0630",
        "identifier": "SCR-5",
        "title": "这是一个缺陷",
        "type": "bug",
        "start_at": 1583290309,
        "end_at": 1583290347,
        "parent_id": "5edca524cad2fa112b06105c",
        "short_id": "c9WqLmTO",
        "html_url": "https://yctech.pingcode.com/pjm/workitems/c9WqLmTO",
        "properties": {
            "prop_a": "prop_a_value",
            "prop_b": "prop_b_value"
        }
    }
}
```

---

### 删除一个关联

**`DELETE /v1/relations/{relation_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `relation_id` | String | 关联的id。 |

**响应示例：**

```json
{
    "id": "fa1125cb06305edca524cad2",
    "url": "https://rest_api_root/v1/relations/fa1125cb06305edca524cad2",
    "principal_type": "test_plan",
    "principal": {
        "id": "5eb6a70571487623fea47000",
        "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/plans/5eb6a70571487623fea47000",
        "name": "测试计划",
        "status": "in_progress",
        "start_at": 1589791860,
        "end_at": 1589791870
    },
    "target_type": "work_item",
    "target": {
        "id": "5edca524cad2fa1125cb0630",
        "url": "https://rest_api_root/v1/project/work_items/5edca524cad2fa1125cb0630",
        "identifier": "SCR-5",
        "title": "这是一个缺陷",
        "type": "bug",
        "start_at": 1583290309,
        "end_at": 1583290347,
        "parent_id": "5edca524cad2fa112b06105c",
        "short_id": "c9WqLmTO",
        "html_url": "https://yctech.pingcode.com/pjm/workitems/c9WqLmTO",
        "properties": {
            "prop_a": "prop_a_value",
            "prop_b": "prop_b_value"
        }
    }
}
```

---

### 获取关联列表

**`GET /v1/relations?principal_type={principal_type}&principal_id={principal_id}&target_type={target_type}`**

**查询参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `principal_type` | String | 关联主体的类型。关联主体的类型和关联的目标类型需要搭配使用，比如： 需求关联工单，principal_type为idea，target_type为ticket； 需求关联工作项，principal_type为idea，target_type为work_item； 需求关联测试用例，principal_type为idea，target_type为test_case； 需求关联需求，principal_type为idea，target_type为idea； 需求关联页面，principal_type为idea，target_type为page； 工单关联需求，principal_type为ticket，target_type为idea； 工单关联工作项，principal_type为ticket，target_type为work_item； 工单关联工单，principal_type为ticket，target_type为ticket； 工单关联页面，principal_type为ticket，target_type为page； 工作项关联测试用例，principal_type为work_item，target_type为test_case； 工作项关联需求，principal_type为work_item，target_type为idea； 工作项关联工单，principal_type为work_item，target_type为ticket； 工作项关联页面，principal_type为work_item，target_type为page； 测试计划关联缺陷，principal_type为test_plan，target_type为work_item； 执行用例关联缺陷，principal_type为test_run，target_type为work_item 测试用例关联需求，principal_type为test_case，target_type为idea； 测试用例关联工作项，principal_type为test_case，target_type为work_item  测试用例关联页面，principal_type为test_case，target_type为page； 页面关联需求，principal_type为page，target_type为idea； 页面关联工单，principal_type为page，target_type为ticket； 页面关联工作项，principal_type为page，target_type为work_item； 页面关联测试用例，principal_type为page，target_type为test_case； |
| `principal_id` | String | 关联主体的id。 |
| `target_type` | String | 关联的目标类型。 |

**Success-Response**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "64b4d70ba368e6594360ea79",
            "url": "https://rest_api_root/v1/relations/64b4d70ba368e6594360ea79",
            "principal_type": "idea",
            "principal": {
                "id": "64b4d70ba368e6594360ea24",
                "url": "https://rest_api_root/v1/ship/ideas/64b4d70ba368e6594360ea24",
                "identifier": "SLC-1",
                "title": "示例需求",
                "short_id": "Ogf1EYey",
                "html_url": "https://yctech.pingcode.com/ship/ideas/Ogf1EYey"
            },
            "target_type": "ticket",
            "target": {
                "id": "63eca888a0a13a3efc8d4a43",
                "url": "https://rest_api_root/v1/ship/tickets/63eca888a0a13a3efc8d4a43",
                "identifier": "SLC-T1",
                "title": "希望新增支持第三方账号注册",
                "short_id": "pdMUgQzZ",
                "html_url": "https://yctech.pingcode.com/ship/tickets/pdMUgQzZ"
            }
        }
    ]
}
```

---

## 发布

### 创建一个发布

**`POST /v1/project/projects/{project_id}/versions`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `project_id` | String | 项目的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `name` | String | 发布的名称。同一项目下该名称是唯一的。 |
| `start_at` | Number | 开始的时间。 |
| `end_at` | Number | 发布的时间。 |
| `assignee_id` | String | 负责人的id。 |
| `stage_id` *(可选)* | String | 发布阶段的id。 |
| `category_ids` *(可选)* | String[] | 发布类别id数组。 |

**响应示例：**

```json
{
    "id": "5eb623f6a70571487ea47001",
    "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/versions/5eb623f6a70571487ea47001",
    "project": {
        "id": "5eb623f6a70571487ea47000",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
        "identifier": "SCR",
        "type": "scrum",
        "name": "Scrum项目",
        "is_archived": 0,
        "is_deleted": 0
    },
    "name": "1.0.0",
    "start_at": 1565193600,
    "end_at": 1566403200,
    "stage": {
        "id": "5f44a8f8bb347b14b49624a1",
        "url": "https://rest_api_root/v1/project/stages/5f44a8f8bb347b14b49624a1",
        "name": "未开始",
        "type": "pending",
        "color": "#FA8888"
    },
    "assignee": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "stages": [
        {
            "id": "5f44a8f8bb347b14b49624a1",
            "url": "https://rest_api_root/v1/project/stages/5f44a8f8bb347b14b49624a1",
            "name": "未开始",
            "operate_at": null
        },
        {
            "id": "5f44a8f8bb347b14b49624a2",
            "url": "https://rest_api_root/v1/project/stages/5f44a8f8bb347b14b49624a2",
            "name": "进行中",
            "operate_at": null
        },
        {
            "id": "5f44a8f8bb347b14b49624a3",
            "url": "https://rest_api_root/v1/project/stages/5f44a8f8bb347b14b49624a3",
            "name": "已发布",
            "operate_at": null
        }
    ],
    "progress": 0,
    "changelog": null,
    "operate_at": 1565255712,
    "categories": [
        {
            "id": "676a460a0fd987b7ea320889",
            "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/version_categories/676a460a0fd987b7ea320889",
            "name": "私有部署发布"
        }
    ],
    "created_at": 1565366200,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "updated_at": 1565366200,
    "updated_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    }
}
```

---

### 创建一个发布分组

**`POST /v1/project/projects/{project_id}/version_sections`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `project_id` | String | 项目的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `name` | String | 发布分组的名称。 |
| `description` *(可选)* | String | 发布分组的描述。 |

**响应示例：**

```json
{
    "id": "63560f3ad02cbc4f9df91251",
    "url": "http://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/version_sections/63560f3ad02cbc4f9df91251",
    "project": {
        "id": "5eb623f6a70571487ea47000",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
        "identifier": "SCR",
        "name": "Scrum项目",
        "type": "scrum",
        "is_archived": 0,
        "is_deleted": 0
    },
    "name": "私有部署",
    "description": "私有部署发布分组"
}
```

---

### 创建一个发布类别

**`POST /v1/project/projects/{project_id}/version_categories`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `project_id` | String | 项目的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `name` | String | 发布类别的名称。 |
| `section_id` *(可选)* | String | 发布类别所属发布分组。 |

**响应示例：**

```json
{
    "id": "676a460a0fd987b7ea320889",
    "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/version_categories/676a460a0fd987b7ea320889",
    "project": {
        "id": "5eb623f6a70571487ea47000",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
        "identifier": "SCR",
        "name": "Scrum项目",
        "type": "scrum",
        "is_archived": 0,
        "is_deleted": 0
    },
    "name": "私有部署发布",
    "section": {
        "id": "63560f3ad02cbc4f9df91251",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/version_sections/63560f3ad02cbc4f9df91251",
        "name": "私有部署发布分组"
    }
}
```

---

### 创建一个发布阶段

**`POST /v1/project/stages`**

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `name` | String | 发布阶段的名称。在一个企业中这个名称是唯一的。 |
| `type` | String | 发布阶段的类型。 |

**响应示例：**

```json
{
   "id": "5c9b35de90ad7153c2062f18",
   "url": "https://rest_api_root/v1/project/stages/5c9b35de90ad7153c2062f18",
   "name": "新建",
   "type": "pending",
   "color": "#ff7575"
}
```

---

### 删除一个发布

**`DELETE /v1/project/projects/{project_id}/versions/{version_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `project_id` | String | 项目的id。 |
| `version_id` | String | 发布的id。 |

**响应示例：**

```json
{
    "id": "5eb623f6a70571487ea47001",
    "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/versions/5eb623f6a70571487ea47001",
    "project": {
        "id": "5eb623f6a70571487ea47000",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
        "identifier": "SCR",
        "type": "scrum",
        "name": "Scrum项目",
        "is_archived": 0,
        "is_deleted": 0
    },
    "name": "1.0.0",
    "start_at": 1565193600,
    "end_at": 1566403200,
    "stage": {
        "id": "5f44a8f8bb347b14b49624a1",
        "url": "https://rest_api_root/v1/project/stages/5f44a8f8bb347b14b49624a1",
        "name": "未开始",
        "type": "pending",
        "color": "#FA8888"
    },
    "assignee": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "stages": [
        {
            "id": "5f44a8f8bb347b14b49624a1",
            "url": "https://rest_api_root/v1/project/stages/5f44a8f8bb347b14b49624a1",
            "name": "未开始",
            "operate_at": null
        },
        {
            "id": "5f44a8f8bb347b14b49624a2",
            "url": "https://rest_api_root/v1/project/stages/5f44a8f8bb347b14b49624a2",
            "name": "进行中",
            "operate_at": null
        },
        {
            "id": "5f44a8f8bb347b14b49624a3",
            "url": "https://rest_api_root/v1/project/stages/5f44a8f8bb347b14b49624a3",
            "name": "已发布",
            "operate_at": null
        }
    ],
    "progress": 0,
    "changelog": null,
    "operate_at": 1565255712,
    "categories": [
        {
            "id": "676a460a0fd987b7ea320889",
            "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/version_categories/676a460a0fd987b7ea320889",
            "name": "私有部署发布"
        }
    ],
    "created_at": 1565366200,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "updated_at": 1565366200,
    "updated_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    }
}
```

---

### 删除一个发布分组

**`DELETE /v1/project/projects/{project_id}/version_sections/{section_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `project_id` | String | 项目的id。 |
| `section_id` | String | 发布分组的id。 |

**响应示例：**

```json
{
    "id": "63560f3ad02cbc4f9df91252",
    "url": "http://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/version_sections/63560f3ad02cbc4f9df91252",
    "project": {
        "id": "5eb623f6a70571487ea47000",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
        "identifier": "SCR",
        "name": "Scrum项目",
        "type": "scrum",
        "is_archived": 0,
        "is_deleted": 0
    },
    "name": "临时部署",
    "description": "临时发布分组"
}
```

---

### 删除一个发布类别

**`DELETE /v1/project/projects/{project_id}/version_categories/{version_category_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `project_id` | String | 项目的id。 |
| `version_category_id` | String | 发布类别的id。 |

**响应示例：**

```json
{
    "id": "676a460a0fd987b7ea320890",
    "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/version_categories/676a460a0fd987b7ea320890",
    "project": {
        "id": "5eb623f6a70571487ea47000",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
        "identifier": "SCR",
        "name": "Scrum项目",
        "type": "scrum",
        "is_archived": 0,
        "is_deleted": 0
    },
    "name": "这是一个发布类别",
    "section": {
        "id": "63560f3ad02cbc4f9df91251",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/version_sections/63560f3ad02cbc4f9df91251",
        "name": "私有部署发布分组"
    }
}
```

---

### 删除一个发布阶段

**`DELETE /v1/project/stages/{stage_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `stage_id` | String | 发布阶段的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `replace_id` *(可选)* | String | 替换的发布阶段id，如果一个发布阶段已经被发布使用，删除该发布阶段时需要提供一个替换的发布阶段。 |

**响应示例：**

```json
{
   "id": "5c9b35de90ad7153c2062f18",
   "url": "https://rest_api_root/v1/project/stages/5c9b35de90ad7153c2062f18",
   "name": "新建",
   "type": "in_progress",
   "color": "#ff7575"
}
```

---

### 批量创建发布

**`POST /v1/project/versions/bulk`**

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `versions` | Object[] | 需要批量创建的发布。该参数是一个对象数组（数组内对象不得超过100个）。 |
| `versions.project_id` | String | 发布所属项目的id。 |
| `versions.name` | String | 发布的名称。 |
| `versions.start_at` | Number | 发布的开始时间。 |
| `versions.end_at` | Number | 发布的时间。 |
| `versions.assignee_id` | String | 发布负责人的id。 |
| `versions.stage_id` | String | 发布的阶段id。 |
| `versions.category_ids` *(可选)* | String[] | 发布类别的id列表。 |

**响应示例：**

```json
[
    {
        "state": "success",
        "version": {
            "id": "5eb623f6a70571487ea47001",
            "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/versions/5eb623f6a70571487ea47001",
            "project": {
                "id": "5eb623f6a70571487ea47000",
                "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
                "identifier": "SCR",
                "type": "scrum",
                "name": "Scrum项目",
                "is_archived": 0,
                "is_deleted": 0
            },
            "name": "1.0.0",
            "start_at": 1565193600,
            "end_at": 1566403200,
            "stage": {
                "id": "5f44a8f8bb347b14b49624a1",
                "url": "https://rest_api_root/v1/project/stages/5f44a8f8bb347b14b49624a1",
                "name": "未开始",
                "type": "pending",
                "color": "#FA8888"
            },
            "assignee": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "stages": [
                {
                    "id": "5f44a8f8bb347b14b49624a1",
                    "url": "https://rest_api_root/v1/project/stages/5f44a8f8bb347b14b49624a1",
                    "name": "未开始",
                    "operate_at": 1565366400
                },
                {
                    "id": "5f44a8f8bb347b14b49624a2",
                    "url": "https://rest_api_root/v1/project/stages/5f44a8f8bb347b14b49624a2",
                    "name": "进行中",
                    "operate_at": null
                },
                {
                    "id": "5f44a8f8bb347b14b49624a3",
                    "url": "https://rest_api_root/v1/project/stages/5f44a8f8bb347b14b49624a3",
                    "name": "已发布",
                    "operate_at": null
                }
            ],
            "progress": 0,
            "changelog": null,
            "operate_at": 1565366400,
            "categories": [
                {
                    "id": "676a460a0fd987b7ea320889",
                    "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/version_categories/676a460a0fd987b7ea320889",
                    "name": "私有部署发布"
                }
            ],
            "created_at": 1565366200,
            "created_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "updated_at": 1565366200,
            "updated_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            }
        }
    }
]
```

---

### 获取发布分组列表

**`GET /v1/project/projects/{project_id}/version_sections`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `project_id` | String | 项目的id。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "63560f3ad02cbc4f9df91251",
            "url": "http://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/version_sections/63560f3ad02cbc4f9df91251",
            "project": {
                "id": "5eb623f6a70571487ea47000",
                "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
                "identifier": "SCR",
                "name": "Scrum项目",
                "type": "scrum",
                "is_archived": 0,
                "is_deleted": 0
            },
            "name": "这是一个发布分组",
            "description": "这是一个发布分组的描述"
        }
    ]
}
```

---

### 获取发布列表

**`GET /v1/project/projects/{project_id}/versions`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `project_id` | String | 项目的id。 |

**查询参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `name` *(可选)* | String | 发布的名字。 |
| `status` *(可选)* | String | 发布的状态。 |
| `created_between` *(可选)* | String | 创建时间介于的时间范围，通过','分割起始时间。 |
| `updated_between` *(可选)* | String | 更新时间介于的时间范围，通过','分割起始时间。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "5eb623f6a70571487ea47001",
            "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/versions/5eb623f6a70571487ea47001",
            "project": {
                "id": "5eb623f6a70571487ea47000",
                "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
                "identifier": "SCR",
                "type": "scrum",
                "name": "Scrum项目",
                "is_archived": 0,
                "is_deleted": 0
            },
            "name": "1.0.0",
            "start_at": 1565193600,
            "end_at": 1566403200,
            "stage": {
                "id": "5f44a8f8bb347b14b49624a1",
                "url": "https://rest_api_root/v1/project/stages/5f44a8f8bb347b14b49624a1",
                "name": "未开始",
                "type": "pending",
                "color": "#FA8888"
            },
            "assignee": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "stages": [
                {
                    "id": "5f44a8f8bb347b14b49624a1",
                    "url": "https://rest_api_root/v1/project/stages/5f44a8f8bb347b14b49624a1",
                    "name": "未开始",
                    "operate_at": 1565366400
                },
                {
                    "id": "5f44a8f8bb347b14b49624a2",
                    "url": "https://rest_api_root/v1/project/stages/5f44a8f8bb347b14b49624a2",
                    "name": "进行中",
                    "operate_at": null
                },
                {
                    "id": "5f44a8f8bb347b14b49624a3",
                    "url": "https://rest_api_root/v1/project/stages/5f44a8f8bb347b14b49624a3",
                    "name": "已发布",
                    "operate_at": null
                }
            ],
            "progress": 0,
            "changelog": "发布日志",
            "operate_at": 1565366400,
            "categories": [
                {
                    "id": "676a460a0fd987b7ea320889",
                    "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/version_categories/676a460a0fd987b7ea320889",
                    "name": "私有部署发布"
                }
            ],
            "created_at": 1565366200,
            "created_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "updated_at": 1565366200,
            "updated_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            }
        }
    ]
}
```

---

### 获取发布类别列表

**`GET /v1/project/projects/{project_id}/version_categories`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `project_id` | String | 项目的id。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "676a460a0fd987b7ea320889",
            "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/version_categories/676a460a0fd987b7ea320889",
            "project": {
                "id": "5eb623f6a70571487ea47000",
                "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
                "identifier": "SCR",
                "name": "Scrum项目",
                "type": "scrum",
                "is_archived": 0,
                "is_deleted": 0
            },
            "name": "私有部署发布",
            "section": {
                "id": "63560f3ad02cbc4f9df91251",
                "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/version_sections/63560f3ad02cbc4f9df91251",
                "name": "私有部署发布分组"
            }
        }
    ]
}
```

---

### 获取发布阶段列表

**`GET /v1/project/stages`**

**响应示例：**

```json
{
     "page_size": 30,
     "page_index": 0,
     "total": 1,
     "values": [
         {
             "id": "5c9b35de90ad7153c2062f18",
             "url": "https://rest_api_root/v1/project/stages/5c9b35de90ad7153c2062f18",
             "name": "新建",
             "type": "in_progress",
             "color": "#ff7575"
         }
     ]
}
```

---

### 部分更新一个发布

**`PATCH /v1/project/projects/{project_id}/versions/{version_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `project_id` | String | 项目的id。 |
| `version_id` | String | 发布的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `name` *(可选)* | String | 发布的名称。同一项目下该名称是唯一的。 |
| `start_at` *(可选)* | Number | 开始的时间。 |
| `end_at` *(可选)* | Number | 发布的时间。 |
| `assignee_id` *(可选)* | String | 负责人的id。 |
| `stage_id` *(可选)* | String | 发布阶段的id。 |
| `operate_at` *(可选)* | Number | 发布阶段的日期。需要和stage_id一起传递。 |
| `category_ids` *(可选)* | String[] | 发布类别id数组。 |

**响应示例：**

```json
{
    "id": "5eb623f6a70571487ea47001",
    "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/versions/5eb623f6a70571487ea47001",
    "project": {
        "id": "5eb623f6a70571487ea47000",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
        "identifier": "SCR",
        "type": "scrum",
        "name": "Scrum项目",
        "is_archived": 0,
        "is_deleted": 0
    },
    "name": "1.0.0",
    "start_at": 1565193600,
    "end_at": 1566403200,
    "stage": {
        "id": "5f44a8f8bb347b14b49624a1",
        "url": "https://rest_api_root/v1/project/stages/5f44a8f8bb347b14b49624a1",
        "name": "未开始",
        "type": "pending",
        "color": "#FA8888"
    },
    "assignee": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "stages": [
        {
            "id": "5f44a8f8bb347b14b49624a1",
            "url": "https://rest_api_root/v1/project/stages/5f44a8f8bb347b14b49624a1",
            "name": "未开始",
            "operate_at": 1565366400
        },
        {
            "id": "5f44a8f8bb347b14b49624a2",
            "url": "https://rest_api_root/v1/project/stages/5f44a8f8bb347b14b49624a2",
            "name": "进行中",
            "operate_at": null
        },
        {
            "id": "5f44a8f8bb347b14b49624a3",
            "url": "https://rest_api_root/v1/project/stages/5f44a8f8bb347b14b49624a3",
            "name": "已发布",
            "operate_at": null
        }
    ],
    "progress": 0,
    "changelog": "发布日志",
    "operate_at": 1565366400,
    "categories": [
        {
            "id": "676a460a0fd987b7ea320889",
            "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/version_categories/676a460a0fd987b7ea320889",
            "name": "私有部署发布"
        }
    ],
    "created_at": 1565366200,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "updated_at": 1565366200,
    "updated_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    }
}
```

---

### 部分更新一个发布分组

**`PATCH /v1/project/projects/{project_id}/version_sections/{section_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `project_id` | String | 项目的id。 |
| `section_id` | String | 发布分组的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `name` | String | 发布分组的名称。 |
| `description` *(可选)* | String | 发布分组的描述。 |

**响应示例：**

```json
{
    "id": "63560f3ad02cbc4f9df91251",
    "url": "http://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/version_sections/63560f3ad02cbc4f9df91251",
    "project": {
        "id": "5eb623f6a70571487ea47000",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
        "identifier": "SCR",
        "name": "Scrum项目",
        "type": "scrum",
        "is_archived": 0,
        "is_deleted": 0
    },
    "name": "私有部署",
    "description": "私有部署发布分组"
}
```

---

### 部分更新一个发布类别

**`PATCH /v1/project/projects/{project_id}/version_categories/{version_category_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `project_id` | String | 项目的id。 |
| `version_category_id` | String | 发布类别的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `name` *(可选)* | String | 发布类别的名称。 |
| `section_id` *(可选)* | String | 发布类别所属发布分组。 |

**响应示例：**

```json
{
    "id": "676a460a0fd987b7ea320889",
    "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/version_categories/676a460a0fd987b7ea320889",
    "project": {
        "id": "5eb623f6a70571487ea47000",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
        "identifier": "SCR",
        "name": "Scrum项目",
        "type": "scrum",
        "is_archived": 0,
        "is_deleted": 0
    },
    "name": "私有部署发布",
    "section": {
        "id": "63560f3ad02cbc4f9df91251",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/version_sections/63560f3ad02cbc4f9df91251",
        "name": "私有部署发布分组"
    }
}
```

---

### 部分更新一个发布阶段

**`PATCH /v1/project/stages/{stage_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `stage_id` | String | 发布阶段的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `name` *(可选)* | String | 发布阶段的名称。在一个企业中这个名称是唯一的。 |
| `type` *(可选)* | String | 发布阶段的类型。 |

**响应示例：**

```json
{
   "id": "5c9b35de90ad7153c2062f18",
   "url": "https://rest_api_root/v1/project/stages/5c9b35de90ad7153c2062f18",
   "name": "新建",
   "type": "in_progress",
   "color": "#ff7575"
}
```

---

## 团队

### 创建一个团队

**`POST /v1/directory/groups`**

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `name` | String | 团队的名称，在一个企业中这个名称是唯一的。 |
| `visibility` *(可选)* | String | 团队的可见范围。 |
| `description` *(可选)* | String | 团队的描述。 |

**响应示例：**

```json
{
    "id": "63c8fb32729dee3334d96af7",
    "url": "https://rest_api_root/v1/directory/groups/63c8fb32729dee3334d96af7",
    "name": "Open Team",
    "visibility": "private",
    "description": "This is Open Team."
}
```

---

### 向团队中添加一个成员

**`POST /v1/directory/groups/{group_id}/members`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `group_id` | String | 团队id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `user_id` | String | 用户id。 |
| `role` | String | 团队角色。 |

**响应示例：**

```json
{
    "id": "a0417f68e846aae315c85d24643678a9",
    "url": "https://rest_api_root/v1/directory/groups/64ca0f67cb78a0a80e1a999e/members/a0417f68e846aae315c85d24643678a9",
    "group": {
        "id": "64ca0f67cb78a0a80e1a999e",
        "url": "https://rest_api_root/v1/directory/groups/64ca0f67cb78a0a80e1a999e",
        "name": "PingCode"
    },
    "user": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "role": "manager"
}
```

---

### 在团队中移除一个成员

**`DELETE /v1/directory/groups/{group_id}/members/{member_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `group_id` | String | 团队id。 |
| `member_id` | String | 团队成员id。 |

**响应示例：**

```json
{
    "id": "a0417f68e846aae315c85d24643678a9",
    "url": "https://rest_api_root/v1/directory/groups/64ca0f67cb78a0a80e1a999e/members/a0417f68e846aae315c85d24643678a9",
    "group": {
        "id": "64ca0f67cb78a0a80e1a999e",
        "url": "https://rest_api_root/v1/directory/groups/64ca0f67cb78a0a80e1a999e",
        "name": "PingCode"
    },
    "user": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "role": "manager"
}
```

---

### 获取团队中的成员列表

**`GET /v1/directory/groups/{group_id}/members`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `group_id` | String | 团队id。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "url": "https://rest_api_root/v1/directory/groups/64ca0f67cb78a0a80e1a999e/members/a0417f68e846aae315c85d24643678a9",
            "group": {
                "id": "64ca0f67cb78a0a80e1a999e",
                "url": "https://rest_api_root/v1/directory/groups/64ca0f67cb78a0a80e1a999e",
                "name": "PingCode"
            },
            "user": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "role": "manager"
        }
    ]
}
```

---

### 获取团队列表

**`GET /v1/directory/groups`**

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 2,
    "values": [
        {
            "id": "63c8fb32729dee3334d96af7",
            "url": "https://rest_api_root/v1/directory/groups/63c8fb32729dee3334d96af7",
            "name": "Open Team",
            "visibility": "private",
            "description": "This is Open Team."
        },
        {
            "id": "64ca0f67cb78a0a80e1a999e",
            "url": "https://rest_api_root/v1/directory/groups/64ca0f67cb78a0a80e1a999e",
            "name": "PingCode",
            "visibility": "public",
            "description": "This is PingCode."
        }
    ]
}
```

---

### 部分更新一个团队

**`PATCH /v1/directory/groups/{group_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `group_id` | String | 团队id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `name` *(可选)* | String | 团队的名称，在一个企业中这个名称是唯一的。 |
| `visibility` *(可选)* | String | 团队的可见范围。 |
| `description` *(可选)* | String | 团队的描述。 |

**响应示例：**

```json
{
    "id": "63c8fb32729dee3334d96af7",
    "url": "https://rest_api_root/v1/directory/groups/63c8fb32729dee3334d96af7",
    "name": "Open Team Update",
    "visibility": "public",
    "description": "This is Update Open Team."
}
```

---

## 外部用户

### 创建一个外部用户

**`POST /v1/ship/products/{product_id}/users`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `product_id` | String | 产品的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `name` | String | 外部用户的名称。 |
| `email` *(可选)* | String | 外部用户的邮箱地址，邮箱地址和手机号其中一个必填。 |
| `mobile` *(可选)* | String | 外部用户的手机号，邮箱地址和手机号其中一个必填，如果同时存在，以手机号为准。 |
| `customer_id` *(可选)* | String | 外部用户所属客户的id。 |

**响应示例：**

```json
{
    "id": "64a2b61c3a12e6c2e46d41e9",
    "url": "https://rest_api_root/v1/ship/products/64a2b61c3a12e6c2e46d41e9/users/64a2b61c3a12e6c2e46d41e9",
    "name": "jack",
    "display_name": "Jack",
    "avatar": "https://s3.amazonaws.com/bucket/a46ef40c-e21e-48cf-a579-cace9fba839a_160x160.png",
    "email": "jack@email.com",
    "mobile": null,
    "product": {
        "id": "6422711c3f12e6c1e46d40e9",
        "url": "http://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9",
        "identifier": "SLC",
        "name": "示例产品",
        "is_archived": 0,
        "is_deleted": 0
    },
    "customer": {
        "id": "64dd899e3f6383ba72ec2a01",
        "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9/customers/64dd899e3f6383ba72ec2a01",
        "name": "深圳XX新零售有限公司"
    }
}
```

---

### 删除一个外部用户

**`DELETE /v1/ship/products/{product_id}/users/{user_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `product_id` | String | 产品的id。 |
| `user_id` | String | 外部用户的id。 |

**响应示例：**

```json
{
    "id": "64a2b61c3a12e6c2e46d41e9",
    "url": "https://rest_api_root/v1/ship/products/64a2b61c3a12e6c2e46d41e9/users/64a2b61c3a12e6c2e46d41e9",
    "name": "jack",
    "display_name": "Jack",
    "avatar": "https://s3.amazonaws.com/bucket/a46ef40c-e21e-48cf-a579-cace9fba839a_160x160.png",
    "email": "jack@email.com",
    "mobile": null,
    "product": {
        "id": "6422711c3f12e6c1e46d40e9",
        "url": "http://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9",
        "identifier": "SLC",
        "name": "示例产品",
        "is_archived": 0,
        "is_deleted": 0
    },
    "customer": {
        "id": "64dd899e3f6383ba72ec2a01",
        "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9/customers/64dd899e3f6383ba72ec2a01",
        "name": "深圳XX新零售有限公司"
    }
}
```

---

### 获取外部用户列表

**`GET /v1/ship/products/{product_id}/users`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `product_id` | String | 产品的id。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "64a2b61c3a12e6c2e46d41e9",
            "url": "https://rest_api_root/v1/ship/products/64a2b61c3a12e6c2e46d41e9/users/64a2b61c3a12e6c2e46d41e9",
            "name": "jack",
            "display_name": "Jack",
            "avatar": "https://s3.amazonaws.com/bucket/a46ef40c-e21e-48cf-a579-cace9fba839a_160x160.png",
            "email": "jack@email.com",
            "mobile": null,
            "product": {
                "id": "6422711c3f12e6c1e46d40e9",
                "url": "http://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9",
                "identifier": "SLC",
                "name": "示例产品",
                "is_archived": 0,
                "is_deleted": 0
            },
            "customer": {
                "id": "64dd899e3f6383ba72ec2a01",
                "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9/customers/64dd899e3f6383ba72ec2a01",
                "name": "深圳XX新零售有限公司"
            }
        }
    ]
}
```

---

### 部分更新一个外部用户

**`PATCH /v1/ship/products/{product_id}/users/{user_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `product_id` | String | 产品的id。 |
| `user_id` | String | 外部用户的id。 |

**响应示例：**

```json
{
    "id": "64a2b61c3a12e6c2e46d41e9",
    "url": "https://rest_api_root/v1/ship/products/64a2b61c3a12e6c2e46d41e9/users/64a2b61c3a12e6c2e46d41e9",
    "name": "jack",
    "display_name": "Jack",
    "avatar": "https://s3.amazonaws.com/bucket/a46ef40c-e21e-48cf-a579-cace9fba839a_160x160.png",
    "email": "jack@email.com",
    "mobile": null,
    "product": {
        "id": "6422711c3f12e6c1e46d40e9",
        "url": "http://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9",
        "identifier": "SLC",
        "name": "示例产品",
        "is_archived": 0,
        "is_deleted": 0
    },
    "customer": {
        "id": "64dd899e3f6383ba72ec2a01",
        "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9/customers/64dd899e3f6383ba72ec2a01",
        "name": "深圳XX新零售有限公司"
    }
}
```

---

## 安全

### 日志

用于查看企业的日志信息。

---

### 日志

用于查看企业的日志信息。

---

## 客户

### 创建一个客户

**`POST /v1/ship/products/{product_id}/customers`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `product_id` | String | 产品的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `name` | String | 客户的名称。 |
| `assignee_id` *(可选)* | String | 客户的负责人id。 |
| `scale` *(可选)* | Number | 客户的规模。 |
| `description` *(可选)* | String | 客户的描述。 |

**响应示例：**

```json
{
    "id": "64dd899e3f6383ba72ec2a0d",
    "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e6/customers/64dd899e3f6383ba72ec2a0d",
    "product": {
        "id": "6422711c3f12e6c1e46d40e6",
        "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e6",
        "identifier": "SLC",
        "name": "示例产品",
        "is_archived": 0,
        "is_deleted": 0
    },
    "name": "上海XX新零售有限公司",
    "assignee": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "scale": 200,
    "description": "上海XX新零售有限公司的描述",
    "created_at": 1692240286,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "updated_at": 1692240286,
    "updated_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "is_archived": 0,
    "is_deleted": 0
}
```

---

### 获取客户列表

**`GET /v1/ship/products/{product_id}/customers`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `product_id` | String | 产品的id。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "64dd899e3f6383ba72ec2a0d",
            "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e6/customers/64dd899e3f6383ba72ec2a0d",
            "product": {
                "id": "6422711c3f12e6c1e46d40e6",
                "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e6",
                "identifier": "SLC",
                "name": "示例产品",
                "is_archived": 0,
                "is_deleted": 0
            },
            "name": "上海XX新零售有限公司",
            "assignee": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "scale": 200,
            "description": "上海XX新零售有限公司的描述",
            "created_at": 1692240286,
            "created_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "updated_at": 1692240286,
            "updated_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "is_archived": 0,
            "is_deleted": 0
        }
    ]
}
```

---

### 部分更新一个客户

**`PATCH /v1/ship/products/{product_id}/customers/{customer_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `product_id` | String | 产品的id。 |
| `customer_id` | String | 客户的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `name` *(可选)* | String | 客户的名称。 |
| `assignee_id` *(可选)* | String | 客户的负责人id。 |
| `scale` *(可选)* | Number | 客户的规模。 |
| `description` *(可选)* | String | 客户的描述。 |

**响应示例：**

```json
{
    "id": "64dd899e3f6383ba72ec2a0d",
    "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e6/customers/64dd899e3f6383ba72ec2a0d",
    "product": {
        "id": "6422711c3f12e6c1e46d40e6",
        "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e6",
        "identifier": "SLC",
        "name": "示例产品",
        "is_archived": 0,
        "is_deleted": 0
    },
    "name": "上海XX新零售有限公司",
    "assignee": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "scale": 200,
    "description": "上海XX新零售有限公司的描述",
    "created_at": 1692240286,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "updated_at": 1692241286,
    "updated_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "is_archived": 0,
    "is_deleted": 0
}
```

---

## 客户端凭据

### 获取企业令牌

**`GET /v1/auth/token?grant_type=client_credentials`**

access_token的有效期为30天，删除PingCode的应用或重置应用的Secret都会导致access_token失效。

**查询参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `grant_type` | String | OAuth2的授予类型，这里固定为：client_credentials。 |
| `client_id` | String | PingCode应用的Client ID |
| `client_secret` | String | PingCode应用的Secret |

**响应示例：**

```json
{
   "access_token": "e7321ca8-f724-4abd-9169-d76d095c6acf",
   "token_type": "Bearer",
   "expires_in": 1577808000
}
```

---

## 工作项关注人

### 添加一个工作项关注人

**`POST /v1/project/work_items/{work_item_id}/participants`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `work_item_id` | String | 工作项的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `user_id` | String | 企业成员的id。 |

**响应示例：**

```json
{
    "id": "a0417f68e846aae315c85d24643678a9",
    "url": "https://rest_api_root/v1/project/work_items/5edca112b06305c524cad2fa/participants/a0417f68e846aae315c85d24643678a9",
    "work_item": {
        "id": "5edca112b06305c524cad2fa",
        "url": "https://rest_api_root/v1/project/work_items/5edca112b06305c524cad2fa",
        "identifier": "SCR-3",
        "title": "这是一个用户故事",
        "type": "story",
        "start_at": 1583290309,
        "end_at": 1583290347,
        "parent_id": "5edca524cad2fa112b06105c",
        "short_id": "c9WqLmTO",
        "html_url": "https://yctech.pingcode.com/pjm/workitems/c9WqLmTO",
        "properties": {
            "prop_a": "prop_a_value",
            "prop_b": "prop_b_value"
         }
    },
    "type": "user",
    "user": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    }
}
```

---

### 移除一个工作项关注人

**`DELETE /v1/project/work_items/{work_item_id}/participants/{participant_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `work_item_id` | String | 工作项的id。 |
| `participant_id` | String | 工作项关注人的id。 |

**响应示例：**

```json
{
    "id": "a0417f68e846aae315c85d24643678a9",
    "url": "https://rest_api_root/v1/project/work_items/5edca112b06305c524cad2fa/participants/a0417f68e846aae315c85d24643678a9",
    "work_item": {
        "id": "5edca112b06305c524cad2fa",
        "url": "https://rest_api_root/v1/project/work_items/5edca112b06305c524cad2fa",
        "identifier": "SCR-3",
        "title": "这是一个用户故事",
        "type": "story",
        "start_at": 1583290309,
        "end_at": 1583290347,
        "parent_id": "5edca524cad2fa112b06105c",
        "short_id": "c9WqLmTO",
        "html_url": "https://yctech.pingcode.com/pjm/workitems/c9WqLmTO",
        "properties": {
            "prop_a": "prop_a_value",
            "prop_b": "prop_b_value"
        }
    },
    "type": "user",
    "user": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    }
}
```

---

### 获取一个工作项关注人列表

**`GET /v1/project/work_items/{work_item_id}/participants`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `work_item_id` | String | 工作项的id。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 2,
    "values": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "url": "https://rest_api_root/v1/project/work_items/5edca112b06305c524cad2fa/participants/a0417f68e846aae315c85d24643678a9",
            "work_item": {
                "id": "5edca112b06305c524cad2fa",
                "url": "https://rest_api_root/v1/project/work_items/5edca112b06305c524cad2fa",
                "identifier": "SCR-3",
                "title": "这是一个用户故事",
                "type": "story",
                "start_at": 1583290309,
                "end_at": 1583290347,
                "parent_id": "5edca524cad2fa112b06105c",
                "short_id": "c9WqLmTO",
                "html_url": "https://yctech.pingcode.com/pjm/workitems/c9WqLmTO",
                "properties": {
                    "prop_a": "prop_a_value",
                    "prop_b": "prop_b_value"
                }
            },
            "type": "user",
            "user": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            }
        },
        {
            "id": "63c8fb32729dee3334d96af7",
            "url": "https://rest_api_root/v1/project/work_items/5edca112b06305c524cad2fa/participants/63c8fb32729dee3334d96af7",
            "work_item": {
                "id": "5edca112b06305c524cad2fa",
                "url": "https://rest_api_root/v1/project/work_items/5edca112b06305c524cad2fa",
                "identifier": "SCR-3",
                "title": "这是一个用户故事",
                "type": "story",
                "start_at": 1583290309,
                "end_at": 1583290347,
                "parent_id": "5edca524cad2fa112b06105c",
                "short_id": "c9WqLmTO",
                "html_url": "https://yctech.pingcode.com/pjm/workitems/c9WqLmTO",
                "properties": {
                   "prop_a": "prop_a_value",
                   "prop_b": "prop_b_value"
                 }
            },
            "type": "user_group",
            "user_group": {
                "id": "63c8fb32729dee3334d96af7",
                "url": "https://rest_api_root/v1/directory/users/63c8fb32729dee3334d96af7",
                "name": "Open Team"
            }
        }
    ]
 }
```

---

## 工作项

### 关联一个工作项

**`POST /v1/project/work_items/{work_item_id}/relations`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `work_item_id` | String | 工作项的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `target_work_item_id` | String | 目标工作项的id。 |
| `relation_type` | String | 关联的类型。 |

**响应示例：**

```json
{
    "id": "58fb35de50ef8153c2062e36",
    "url": "https://rest_api_root/v1/project/work_items/5edca524cad2b06305cfa112/relations/58fb35de50ef8153c2062e36",
    "relation_type": "relate",
    "origin_work_item": {
        "id": "5edca524cad2b06305cfa112",
        "url": "https://rest_api_root/v1/project/work_items/5edca524cad2b06305cfa112",
        "identifier": "SCR-4",
        "title": "这是一个任务",
        "type": "task",
        "start_at": 1583290309,
        "end_at": 1583290347,
        "parent_id": "5edca524cad2fa1125cb0629",
        "short_id": "c9WqLmTO",
        "html_url": "https://yctech.pingcode.com/pjm/workitems/c9WqLmTO",
        "properties": {
            "prop_a": "prop_a_value",
            "prop_b": "prop_b_value"
        }
    },
    "target_work_item": {
        "id": "5f9a65ef20ef8153c1462e64",
        "url": "https://rest_api_root/v1/project/work_items/5f9a65ef20ef8153c1462e64",
        "identifier": "SCR-5",
        "title": "这是一个缺陷",
        "type": "bug",
        "start_at": 1583290309,
        "end_at": 1583290347,
        "parent_id": null,
        "short_id": "a9WqLmTO",
        "html_url": "https://yctech.pingcode.com/pjm/workitems/a9WqLmTO",
        "properties": {
            "prop_a": "prop_a_value",
            "prop_b": "prop_b_value"
        }
    }
}
```

---

### 创建一个工作项

**`POST /v1/project/work_items`**

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `project_id` | String | 项目的id。 |
| `type_id` | String | 工作项类型的id。工作项类型分为9种系统类型和一些自定义类型。系统类型包括：史诗、特性、用户故事、阶段、里程碑、需求、任务、缺陷和事务。可以通过获取全部工作项类型列表获得。 |
| `title` | String | 工作项的标题。 |
| `description` *(可选)* | String | 工作项的描述。 |
| `start_at` *(可选)* | Number | 工作项的开始时间。当工作项类型为里程碑时，该参数无效。 |
| `end_at` *(可选)* | Number | 工作项的截止时间。 |
| `priority_id` *(可选)* | String | 工作项优先级的id。 |
| `state_id` *(可选)* | String | 工作项状态的id。工作项状态的id在设置时必须同时满足工作项类型的工作项状态方案和工作项状态流转的状态值才能成功完成设置。工作项状态方案可以通过获取工作项状态方案列表获取。工作项状态流转则可以通过获取状态方案中的工作项状态流转列表获取。 |
| `assignee_id` *(可选)* | String | 工作项负责人的id。 |
| `parent_id` *(可选)* | String | 工作项的父工作项的id。当前工作项类型支持的父工作类型包含parent_id对应的工作项类型时，该参数有效。具体配置见属性与视图子工作项配置。 |
| `sprint_id` *(可选)* | String | 所属迭代id。该字段只有项目类型为scrum和hybrid时有效。 |
| `version_ids` *(可选)* | String[] | 所属发布的id列表。 |
| `board_id` *(可选)* | String | 看板的id。该字段只有项目类型为kanban和hybrid时有效。 |
| `entry_id` *(可选)* | String | 看板栏的id。该字段只有项目类型为kanban和hybrid时有效。 |
| `swimlane_id` *(可选)* | String | 泳道的id。该字段只有项目类型为kanban和hybrid时有效。 |
| `story_points` *(可选)* | Number | 工作项的故事点。当工作项的属性在视图中配置了故事点属性时，该参数生效。故事点数值必须是大于等于0的整数或最多包含一位小数的正数。 |
| `estimated_workload` *(可选)* | Number | 工作项的预估工时。 |
| `remaining_workload` *(可选)* | Number | 工作项的剩余工时。 |
| `properties` *(可选)* | Object | 工作项属性的键值对集合，需要注意的是，当前工作项类型对应的工作项属性方案需要包含这些工作项属性，例如工作项属性方案中包含prop_a和prop_b。 |
| `properties.prop_a` *(可选)* | Object | 工作项属性prop_a。 |
| `properties.prop_b` *(可选)* | Object | 工作项属性prop_b。 |
| `participant_ids` *(可选)* | String[] | 工作项关注人的id列表。 |

**响应示例：**

```json
{
    "id": "5edca524cad2fa1125cb0630",
    "url": "https://rest_api_root/v1/project/work_items/5edca524cad2fa1125cb0630",
    "project": {
        "id": "5eb623f6a70571487ea47000",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
        "identifier": "SCR",
        "name": "Scrum项目",
        "type": "scrum",
        "is_archived": 0,
        "is_deleted": 0
    },
    "identifier": "SCR-5",
    "title": "这是一个缺陷",
    "type": "bug",
    "start_at": 1583290309,
    "end_at": 1583290347,
    "parent_id": "5edca112b06305c524cad2fa",
    "short_id": "1bAqLmTG",
    "html_url": "https://yctech.pingcode.com/pjm/workitems/1bAqLmTG",
    "parent": {
        "id": "5edca112b06305c524cad2fa",
        "url": "https://rest_api_root/v1/project/work_items/5edca112b06305c524cad2fa",
        "identifier": "SCR-3",
        "title": "这是一个用户故事",
        "type": "story",
        "start_at": 1583290309,
        "end_at": 1583290347,
        "parent_id": "5edca524cad2fa112b06309c",
        "properties": {
            "prop_a": "prop_a_value",
            "prop_b": "prop_b_value",
            "risk": null,
            "business_value": null,
            "effort": null,
            "backlog_type": null,
            "backlog_from": null
        }
    },
    "assignee": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "versions": [
       {
          "id": "5eb623487ea47001f6a70571",
          "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/versions/5eb623487ea47001f6a70571",
          "name": "1.0.1",
          "start_at": 1565255712,
          "end_at": 1565255879,
          "stage": {
              "id": "5f44a8f8bb347b14b49624a1",
              "url": "https://rest_api_root/v1/project/stages/5f44a8f8bb347b14b49624a1",
              "name": "未开始",
              "type": "pending",
              "color": "#FA8888"
          }
       }
    ],
    "sprint": {
        "id": "5ecf7b74eaab845a2aa53138",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/sprints/5ecf7b74eaab845a2aa53138",
        "name": "Sprint 1",
        "start_at": 1589791860,
        "end_at": 1589791860,
        "status": "completed"
    },
    "state": {
        "id": "5c9b35de90ad7153c2062f18",
        "url": "https://rest_api_root/v1/project/work_item_states/5c9b35de90ad7153c2062f18",
        "name": "新建",
        "type": "pending",
        "color": "#ff7575"
    },
    "priority": {
        "id": "5eb623f6a70571487ea47111",
        "url": "https://rest_api_root/v1/project/work_item_priorities/5eb623f6a70571487ea47111",
        "name": "最高"
    },
    "board": null,
    "entry": null,
    "swimlane": null,
    "phase": null,
    "description": "这是一个缺陷的描述",
    "completed_at": 1583290347,
    "story_points": 1,
    "estimated_workload": 1,
    "remaining_workload": 1,
    "properties": {
        "prop_a": "prop_a_value",
        "prop_b": "prop_b_value",
        "severity": null,
        "replay_version": null,
        "reappear_probability": null,
        "bug_type": null,
        "reason": null,
        "solution": null,
        "replay_step": null
    },
    "tags": [],
    "participants": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "url": "https://rest_api_root/v1/participants/a0417f68e846aae315c85d24643678a9?principal_type=work_item&principal_id=5edca524cad2fa1125cb0630",
            "type": "user",
            "user": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            }
        }
    ],
    "created_at": 1583290300,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "updated_at": 1583290300,
    "updated_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "is_archived": 0,
    "is_deleted": 0
}
```

---

### 创建一个工作项交付目标

**`POST /v1/project/deliverables`**

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `work_item_id` | String | 工作项的id。工作项所属项目类型必须为waterfall或hybrid。 |
| `name` | String | 工作项交付目标的名称。 |
| `content_type` *(可选)* | String | 工作项交付物的类型。工作项交付物的类型。只支持link。attachment类型的交付物通过向交付目标中上传一个文件接口上传附件。 |
| `content` *(可选)* | String | 工作项交付物。当工作项交付物的类型是link时，工作项交付物为包含name和href的对象，例如：{ &quot;name&quot;: &quot;链接工作项交付目标&quot;,  &quot;href&quot;: &quot;https://rest_api_root/wiki/spaces/public/pages/6472e6f2f1968d3fdb0aba15&quot; }。当工作项交付物不为空时，参数必须包含交付物类型。 |

**响应示例：**

```json
{
    "id": "63761fee31caaf7718981876",
    "url": "https://rest_api_root/v1/project/deliverables/63761fee31caaf7718981876",
    "name": "阶段交付目标",
    "content_type": "link",
    "content": {
        "name": "PingCode",
        "href": "https://www.pingcode.com"
    },
    "project": {
        "id": "6375cc81e3004de4ea14aa52",
        "url": "https://rest_api_root/v1/project/projects/6375cc81e3004de4ea14aa52",
        "identifier": "WTF",
        "name": "瀑布项目",
        "type": "waterfall",
        "is_archived": 0,
        "is_deleted": 0
    },
    "work_item": {
        "id": "63761fee31caaf77189816b4",
        "url": "https://rest_api_root/v1/project/work_items/63761fee31caaf77189816b4",
        "identifier": "WTF-5",
        "title": "这是一个阶段",
        "type": "630da48bc9443b1aa94ce3ee",
        "start_at": 1583290309,
        "end_at": 1583290347,
        "parent_id": "5edca524cad2fa1125cb0629",
        "short_id": "c9WqLmTO",
        "html_url": "https://yctech.pingcode.com/pjm/workitems/c9WqLmTO",
        "properties": {
            "prop_a": "prop_a_value",
            "prop_b": "prop_b_value"
        }
    },
    "created_at": 1668685806,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "updated_at": 1668685806,
    "updated_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "is_archived": 0,
    "is_deleted": 0
}
```

---

### 创建一个工作项工时

**`POST /v1/project/workloads`**

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `work_item_id` | String | 工作项的id。 |
| `workload_type_id` | String | 工时类型的id。 |
| `duration` | Number | 时长。时长的单位是小时。数值可以是为0-24之间，最多包含一位小数的正数。 |
| `report_at` | Number | 登记日期。该值为十位数字组成的时间戳，会被转换为该时间当天的零点零分零秒。 |
| `report_by` *(可选)* | String | 登记人，企业鉴权时必填。个人鉴权时不需要传递，即使传递了也会被忽略。 |
| `description` *(可选)* | String | 工时的描述信息。 |

**响应示例：**

```json
{
    "id": "5f168f764eba01a5278b87cd",
    "url": "https://rest_api_root/v1/project/workloads/5f168f764eba01a5278b87cd",
    "project": {
        "id": "5eb623f6a70571487ea47000",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
        "identifier": "SCR",
        "name": "Scrum项目",
        "type": "scrum",
        "is_archived": 0,
        "is_deleted": 0
    },
    "work_item": {
        "id": "5edca524cad2fa1125cb0630",
        "url": "https://rest_api_root/v1/project/work_items/5edca524cad2fa1125cb0630",
        "identifier": "SCR-5",
        "title": "这是一个缺陷",
        "type": "bug",
        "start_at": 1583290309,
        "end_at": 1583290347,
        "parent_id": "5edca524cad2fa1125cb0629",
        "short_id": "c9WqLmTO",
        "html_url": "https://yctech.pingcode.com/pjm/workitems/c9WqLmTO",
        "properties": {
            "prop_a": "prop_a_value",
            "prop_b": "prop_b_value"
        }
    },
    "type": {
        "id": "5a86eaf6a72585327ea46fge0",
        "url": "https://rest_api_root/v1/project/workload_types/5a86eaf6a72585327ea46fge0",
        "name": "研发",
        "color": "#56ABF"
    },
    "duration": 8,
    "description": "这是一个工时",
    "report_at": 1593290347,
    "report_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "created_at": 1593290347
}
```

---

### 删除一个工作项

**`DELETE /v1/project/work_items/{work_item_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `work_item_id` | String | 工作项的id。 |

**响应示例：**

```json
{
    "id": "5edca5d2fa112b06305c24ca",
    "url": "https://rest_api_root/v1/project/work_items/5edca5d2fa112b06305c24ca",
    "project": {
        "id": "5eb623f6a70571487ea41919",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919",
        "identifier": "KB",
        "name": "看板项目",
        "type": "kanban",
        "is_archived": 0,
        "is_deleted": 0
    },
    "identifier": "KB-1",
    "title": "这是一个事务",
    "type": "issue",
    "start_at": 1583290309,
    "end_at": 1583290347,
    "short_id": "c9WqLmTO",
    "html_url": "https://yctech.pingcode.com/pjm/workitems/c9WqLmTO",
    "assignee": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "versions": null,
    "sprint": null,
    "state": {
        "id": "5c9b35de90ad7153c2062f18",
        "url": "https://rest_api_root/v1/project/work_item_states/5c9b35de90ad7153c2062f18",
        "name": "新建",
        "type": "pending",
        "color": "#ff7575"
    },
    "priority": {
        "id": "5eb623f6a70571487ea47111",
        "url": "https://rest_api_root/v1/project/work_item_priorities/5eb623f6a70571487ea47111",
        "name": "最高"
    },
    "board": {
        "id": "5eb623f6a70571487ea47222",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919/boards/5eb623f6a70571487ea47222",
        "name": "kanban",
        "work_item_types": [
            "epic",
            "feature",
            "issue",
            "story"
        ]
    },
    "entry": {
        "id": "5ee1c4fac5b3c51206f0a861",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919/boards/5eb623f6a70571487ea47222/entries/5ee1c4fac5b3c51206f0a861",
        "name": "需求池"
    },
    "swimlane": {
        "id": "5ee1c4fac5b3c51206f0a866",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919/boards/5eb623f6a70571487ea47222/swimlanes/5ee1c4fac5b3c51206f0a866",
        "name": "默认泳道"
    },
    "phase": null,
    "description": "这是一个事务的描述",
    "completed_at": 1583290347,
    "story_points": null,
    "estimated_workload": 1,
    "remaining_workload": 1,
    "properties": {
        "prop_a": "prop_a_value",
        "prop_b": "prop_b_value",
        "entry_status": null,
        "entry_position": null,
        "operation_time": 1691571221
    },
    "tags": [
        {
            "id": "5e6b35de50ef8153c2062f70",
            "url": "https://rest_api_root/v1/project/tags/5e6b35de50ef8153c2062f70",
            "name": "标签-1"
        }
    ],
    "participants": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "url": "https://rest_api_root/v1/participants/a0417f68e846aae315c85d24643678a9?principal_type=work_item&principal_id=5edca5d2fa112b06305c24ca",
            "type": "user",
            "user": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            }
        }
    ],
    "created_at": 1583290300,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "updated_at": 1583290300,
    "updated_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "is_archived": 0,
    "is_deleted": 1
}
```

---

### 删除一个工作项交付目标

**`DELETE /v1/project/deliverables/{deliverable_target_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `deliverable_target_id` | String | 工作项交付目标的id。 |

**响应示例：**

```json
{
    "id": "63761fee31caaf7718981876",
    "url": "https://rest_api_root/v1/project/deliverables/63761fee31caaf7718981876",
    "name": "阶段交付目标",
    "content_type": "link",
    "content": {
        "name": "PingCode",
        "href": "https://www.pingcode.com"
    },
    "project": {
        "id": "6375cc81e3004de4ea14aa52",
        "url": "https://rest_api_root/v1/project/projects/6375cc81e3004de4ea14aa52",
        "identifier": "WTF",
        "name": "瀑布项目",
        "type": "waterfall",
        "is_archived": 0,
        "is_deleted": 0
    },
    "work_item": {
        "id": "63761fee31caaf77189816b4",
        "url": "https://rest_api_root/v1/project/work_items/63761fee31caaf77189816b4",
        "identifier": "WTF-5",
        "title": "这是一个阶段",
        "type": "630da48bc9443b1aa94ce3ee",
        "start_at": 1583290309,
        "end_at": 1583290347,
        "parent_id": "5edca524cad2fa1125cb0629",
        "short_id": "c9WqLmTO",
        "html_url": "https://yctech.pingcode.com/pjm/workitems/c9WqLmTO",
        "properties": {
            "prop_a": "prop_a_value",
            "prop_b": "prop_b_value"
        }
    },
    "created_at": 1668685806,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "updated_at": 1668685806,
    "updated_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "is_archived": 0,
    "is_deleted": 1
}
```

---

### 删除一个工作项工时

**`DELETE /v1/project/workloads/{workload_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `workload_id` | String | 工时的id。 |

**响应示例：**

```json
{
    "id": "5f168f764eba01a5278b87cd",
    "url": "https://rest_api_root/v1/project/workloads/5f168f764eba01a5278b87cd",
    "project": {
        "id": "5eb623f6a70571487ea47000",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
        "identifier": "SCR",
        "name": "Scrum项目",
        "type": "scrum",
        "is_archived": 0,
        "is_deleted": 0
    },
    "work_item": {
        "id": "5edca524cad2fa1125cb0630",
        "url": "https://rest_api_root/v1/project/work_items/5edca524cad2fa1125cb0630",
        "identifier": "SCR-5",
        "title": "这是一个缺陷",
        "type": "bug",
        "start_at": 1583290309,
        "end_at": 1583290347,
        "parent_id": "5edca524cad2fa1125cb0629",
        "short_id": "c9WqLmTO",
        "html_url": "https://yctech.pingcode.com/pjm/workitems/c9WqLmTO",
        "properties": {
            "prop_a": "prop_a_value",
            "prop_b": "prop_b_value"
        }
    },
    "type": {
        "id": "5a86eaf6a72585327ea46fge0",
        "url": "https://rest_api_root/v1/project/workload_types/5a86eaf6a72585327ea46fge0",
        "name": "研发",
        "color": "#56ABF"
    },
    "duration": 8,
    "description": "这是一个工时",
    "report_at": 1593290347,
    "report_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "created_at": 1593290347
}
```

---

### 取消关联一个工作项

**`DELETE /v1/project/work_items/{work_item_id}/relations/{relation_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `work_item_id` | String | 工作项的id。 |
| `relation_id` | String | 工作项关联的id。 |

**响应示例：**

```json
{
    "id": "58fb35de50ef8153c2062e36",
    "url": "https://rest_api_root/v1/project/work_items/5edca524cad2b06305cfa112/relations/58fb35de50ef8153c2062e36",
    "relation_type": "relate",
    "origin_work_item": {
        "id": "5edca524cad2b06305cfa112",
        "url": "https://rest_api_root/v1/project/work_items/5edca524cad2b06305cfa112",
        "identifier": "SCR-4",
        "title": "这是一个任务",
        "type": "task",
        "start_at": 1583290309,
        "end_at": 1583290347,
        "parent_id": "5edca524cad2fa1125cb0629",
        "short_id": "c9WqLmTO",
        "html_url": "https://yctech.pingcode.com/pjm/workitems/c9WqLmTO",
        "properties": {
            "prop_a": "prop_a_value",
            "prop_b": "prop_b_value"
        }
    },
    "target_work_item": {
        "id": "5f9a65ef20ef8153c1462e64",
        "url": "https://rest_api_root/v1/project/work_items/5f9a65ef20ef8153c1462e64",
        "identifier": "SCR-5",
        "title": "这是一个缺陷",
        "type": "bug",
        "start_at": 1583290309,
        "end_at": 1583290347,
        "parent_id": null,
        "short_id": "a9WqLmTO",
        "html_url": "https://yctech.pingcode.com/pjm/workitems/a9WqLmTO",
        "properties": {
            "prop_a": "prop_a_value",
            "prop_b": "prop_b_value"
        }
    }
}
```

---

### 向工作项中添加一个标签

**`POST /v1/project/work_items/{work_item_id}/tags`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `work_item_id` | String | 工作项的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `tag_id` | String | 标签的id。 |

**响应示例：**

```json
{
    "id": "5e6b35de50ef8153c2062f70",
    "url": "https://rest_api_root/v1/project/work_items/5edca524cad2fa1125cb0630/tags/5e6b35de50ef8153c2062f70",
    "work_item": {
        "id": "5edca524cad2fa1125cb0630",
        "url": "https://rest_api_root/v1/project/work_items/5edca524cad2fa1125cb0630",
        "identifier": "SCR-5",
        "title": "这是一个缺陷",
        "type": "bug",
        "start_at": 1583290309,
        "end_at":1583290347,
        "parent_id": "5edca524cad2fa112b05105c",
        "short_id": "c9WqLmTO",
        "html_url": "https://yctech.pingcode.com/pjm/workitems/c9WqLmTO",
        "properties": {
            "prop_a": "prop_a_value",
            "prop_b": "prop_b_value"
        }
    },
    "tag": {
        "id": "5e6b35de50ef8153c2062f70",
        "url": "https://rest_api_root/v1/project/tags/5e6b35de50ef8153c2062f70",
        "name": "标签-1"
    }
}
```

---

### 在工作项中移除一个标签

**`DELETE /v1/project/work_items/{work_item_id}/tags/{tag_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `work_item_id` | String | 工作项的id。 |
| `tag_id` | String | 标签的id。 |

**响应示例：**

```json
{
    "id": "5e6b35de50ef8153c2062f70",
    "url": "https://rest_api_root/v1/project/work_items/5edca524cad2fa1125cb0630/tags/5e6b35de50ef8153c2062f70",
    "work_item": {
        "id": "5edca524cad2fa1125cb0630",
        "url": "https://rest_api_root/v1/project/work_items/5edca524cad2fa1125cb0630",
        "identifier": "SCR-5",
        "title": "这是一个缺陷",
        "type": "bug",
        "start_at": 1583290309,
        "end_at":1583290347,
        "parent_id": "5edca524cad2fa112b05105c",
        "short_id": "c9WqLmTO",
        "html_url": "https://yctech.pingcode.com/pjm/workitems/c9WqLmTO",
        "properties": {
            "prop_a": "prop_a_value",
            "prop_b": "prop_b_value"
        }
    },
    "tag": {
        "id": "5e6b35de50ef8153c2062f70",
        "url": "https://rest_api_root/v1/project/tags/5e6b35de50ef8153c2062f70",
        "name": "标签-1"
    }
}
```

---

### 批量部分更新工作项属性

**`PATCH /v1/project/work_items`**

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `ids` | String[] | 需要更新的工作项的id列表。最多可以批量更新100个工作项。 |
| `property_name` | String | 需要更新的工作项属性名。 |
| `property_value` *(可选)* | String | 需要更新的工作项属性值。 |

**响应示例：**

```json
{
    "inserts": 0,
    "updates": 1,
    "deletes": 0
}
```

---

### 获取关联的工作项列表

**`GET /v1/project/work_items/{work_item_id}/relations`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `work_item_id` | String | 工作项的id。 |

**查询参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `relation_type` *(可选)* | String | 关联的类型。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "58fb35de50ef8153c2062e36",
            "url": "https://rest_api_root/v1/project/work_items/5edca524cad2b06305cfa112/relations/58fb35de50ef8153c2062e36",
            "relation_type": "relate",
            "origin_work_item": {
                "id": "5edca524cad2b06305cfa112",
                "url": "https://rest_api_root/v1/project/work_items/5edca524cad2b06305cfa112",
                "identifier": "SCR-4",
                "title": "这是一个任务",
                "type": "task",
                "start_at": 1583290309,
                "end_at": 1583290347,
                "parent_id": "5edca524cad2fa1125cb0629",
                "short_id": "c9WqLmTO",
                "html_url": "https://yctech.pingcode.com/pjm/workitems/c9WqLmTO",
                "properties": {
                    "prop_a": "prop_a_value",
                    "prop_b": "prop_b_value"
                }
            },
            "target_work_item": {
                "id": "5f9a65ef20ef8153c1462e64",
                "url": "https://rest_api_root/v1/project/work_items/5f9a65ef20ef8153c1462e64",
                "identifier": "SCR-5",
                "title": "这是一个缺陷",
                "type": "bug",
                "start_at": 1583290309,
                "end_at": 1583290347,
                "parent_id": null,
                "short_id": "a9WqLmTO",
                "html_url": "https://yctech.pingcode.com/pjm/workitems/a9WqLmTO",
                "properties": {
                    "prop_a": "prop_a_value",
                    "prop_b": "prop_b_value"
                }
            }
        }
    ]
}
```

---

### 获取工作项交付目标列表

**`GET /v1/project/deliverables`**

**查询参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `project_id` *(可选)* | String | 项目的id。项目类型必须为waterfall或hybrid。 |
| `work_item_id` *(可选)* | String | 工作项的id。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 2,
    "values": [
        {
            "id": "63761fee31caaf7718981876",
            "url": "https://rest_api_root/v1/project/deliverables/63761fee31caaf7718981876",
            "name": "阶段交付目标",
            "content_type": "link",
            "content": {
                "name": "PingCode",
                "href": "https://www.pingcode.com"
            },
            "project": {
                "id": "6375cc81e3004de4ea14aa52",
                "url": "https://rest_api_root/v1/project/projects/6375cc81e3004de4ea14aa52",
                "identifier": "WTF",
                "name": "瀑布项目",
                "type": "waterfall",
                "is_archived": 0,
                "is_deleted": 0
            },
            "work_item": {
                "id": "63761fee31caaf77189816b4",
                "url": "https://rest_api_root/v1/project/work_items/63761fee31caaf77189816b4",
                "identifier": "WTF-5",
                "title": "这是一个阶段",
                "type": "630da48bc9443b1aa94ce3ee",
                "start_at": 1583290309,
                "end_at": 1583290347,
                "parent_id": "5edca524cad2fa1125cb0629",
                "short_id": "c9WqLmTO",
                "html_url": "https://yctech.pingcode.com/pjm/workitems/c9WqLmTO",
                "properties": {
                   "prop_a": "prop_a_value",
                   "prop_b": "prop_b_value"
                }
            },
            "created_at": 1668685806,
            "created_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "updated_at": 1668685806,
            "updated_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "is_archived": 0,
            "is_deleted": 0
        },
        {
            "id": "63761fee31caaf7718981877",
            "url": "https://rest_api_root/v1/project/deliverables/63761fee31caaf7718981877",
            "name": "缺陷交付目标",
            "content_type": "attachment",
            "content": {
                "id": "64abd9050461799795b6ea3e",
                "url": "https://rest_api_root/v1/attachments/64abd9050461799795b6ea3e?deliverable_target_id=63761fee31caaf7718981877",
                "title": "fixed.png",
                "size": 11396,
                "type": "file"
            },
            "project": {
                "id": "6375cc81e3004de4ea14aa52",
                "url": "https://rest_api_root/v1/project/projects/6375cc81e3004de4ea14aa52",
                "identifier": "WTF",
                "name": "瀑布项目",
                "type": "waterfall",
                "is_archived": 0,
                "is_deleted": 0
            },
            "work_item": {
                "id": "63761fee31caaf77189816b5",
                "url": "https://rest_api_root/v1/project/work_items/63761fee31caaf77189816b5",
                "identifier": "WTF-6",
                "title": "这是一个缺陷",
                "type": "bug",
                "start_at": 1583290319,
                "end_at": 1583290357,
                "parent_id": "5edca524cad2fa1125cb0623",
                "short_id": "c9WqLmTO",
                "html_url": "https://yctech.pingcode.com/pjm/workitems/c9WqLmTO",
                "properties": {
                    "prop_a": "prop_a_value",
                    "prop_b": "prop_b_value"
                }
            },
            "created_at": 1668685816,
            "created_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "updated_at": 1668685816,
            "updated_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "is_archived": 0,
            "is_deleted": 0
        }
    ]
}
```

---

### 获取工作项优先级列表

**`GET /v1/project/work_item/priorities?project_id={project_id}`**

**查询参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `project_id` | String | 项目的id。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "5eb623f6a70571487ea47111",
            "url": "https://rest_api_root/v1/project/work_item_priorities/5eb623f6a70571487ea47111",
            "name": "最高"
        }
    ]
}
```

---

### 获取工作项关联类型列表

**`GET /v1/project/work_item/relation_types`**

**响应示例：**

```json
{
    "page_index": 0,
    "page_size": 30,
    "total": 1,
    "values": [
        {
            "id": "676510af06fd48a4a4e12616",
            "url": "https://rest_api_root/v1/project/work_item/relation_types/676510af06fd48a4a4e12616",
            "name": "重复",
            "category": "duplicate",
            "is_system": 1
        }
    ]
}
```

---

### 获取工作项列表

**`GET /v1/project/work_items`**

**查询参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `identifier` *(可选)* | String | 工作项编号。 |
| `project_ids` *(可选)* | String | 项目的id，使用','分割，最多只能20个。 |
| `type_ids` *(可选)* | String | 工作项类型的id。工作项类型分为9种系统类型和一些自定义类型。系统类型包括：史诗、特性、用户故事、阶段、里程碑、需求、任务、缺陷和事务。可以通过获取全部工作项类型列表获得。最多只能20个。 |
| `parent_ids` *(可选)* | String | 父工作项的id，使用','分割，最多只能20个。 |
| `assignee_ids` *(可选)* | String | 工作项负责人的id，使用','分割，最多只能20个。 |
| `state_ids` *(可选)* | String | 工作项状态的id，使用','分割，最多只能20个。 |
| `start_between` *(可选)* | String | 开始时间介于的时间范围，通过','分割起始时间。比如1580000000,1590000000表示开始时间介于两个时间之间；,1590000000表示开始时间小于该时间；1580000000表示开始时间大于该时间。 |
| `end_between` *(可选)* | String | 结束时间介于的时间范围，通过','分割起始时间。使用方式参考开始时间。 |
| `priority_ids` *(可选)* | String | 工作项优先级的id，使用','分割，最多只能20个。 |
| `bug_type_ids` *(可选)* | String | 缺陷类别的id，使用','分割，最多只能20个。 |
| `tag_ids` *(可选)* | String | 工作项标签的id，使用','分割，最多只能20个。 |
| `sprint_ids` *(可选)* | String | 迭代的id，使用','分割，最多只能20个。 |
| `board_ids` *(可选)* | String | 看板的id，使用','分割，最多只能20个。 |
| `entry_ids` *(可选)* | String | 看板栏的id，使用','分割，最多只能20个。 |
| `swimlane_ids` *(可选)* | String | 泳道的id，使用','分割，最多只能20个。 |
| `phase_ids` *(可选)* | String | 所属计划的id，使用','分割，最多只能20个。 |
| `version_ids` *(可选)* | String | 发布的id，使用','分割，最多只能20个。 |
| `created_by_ids` *(可选)* | String | 创建人的id，使用','分割，最多只能20个。 |
| `created_between` *(可选)* | String | 创建时间介于的时间范围，通过','分割起始时间。使用方式参考开始时间。 |
| `updated_between` *(可选)* | String | 更新时间介于的时间范围，通过','分割起始时间。使用方式参考开始时间。 |
| `participant_id` *(可选)* | String | 工作项关注人的id。 |
| `keywords` *(可选)* | String | 关键字。支持工作项编号和工作项标题。 |
| `include_public_image_token` *(可选)* | String | 包含获取图片资源token的属性。使用','分割，最多只能20个，支持description和自定义多行文本类型的属性。参数示例description,properties.prop_b。 |
| `include_deleted` *(可选)* | Boolean | 是否查询已删除的工作项。该值默认为false。 |
| `include_archived` *(可选)* | Boolean | 是否查询已归档的工作项。该值默认为false。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 4,
    "values": [
        {
            "id": "5edca524cad2fa112b06305c",
            "url": "https://rest_api_root/v1/project/work_items/5edca524cad2fa112b06305c",
            "project": {
                "id": "5eb623f6a70571487ea47000",
                "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
                "identifier": "SCR",
                "name": "Scrum项目",
                "type": "scrum",
                "is_archived": 0,
                "is_deleted": 0
            },
            "identifier": "SCR-1",
            "title": "这是一个史诗",
            "type": "epic",
            "start_at": 1583290309,
            "end_at": 1583290347,
            "parent_id": null,
            "short_id": "d9WqLmTO",
            "html_url": "https://yctech.pingcode.com/pjm/workitems/d9WqLmTO",
            "parent": null,
            "assignee": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "versions": null,
            "sprint": null,
            "state": {
                "id": "5c9b35de90ad7153c2062f18",
                "url": "https://rest_api_root/v1/project/work_item_states/5c9b35de90ad7153c2062f18",
                "name": "新建",
                "type": "pending",
                "color": "#ff7575"
            },
            "priority": {
                "id": "5eb623f6a70571487ea47111",
                "url": "https://rest_api_root/v1/project/work_item_priorities/5eb623f6a70571487ea47111",
                "name": "最高"
            },
            "board": null,
            "entry": null,
            "swimlane": null,
            "phase": null,
            "description": "这是一个史诗的描述",
            "completed_at": 1583290347,
            "story_points": null,
            "estimated_workload": 1,
            "remaining_workload": 1,
            "properties": {
                "prop_a": "prop_a_value",
                "prop_b": "prop_b_value",
                "risk": null,
                "business_value": null,
                "effort": 123,
                "backlog_from": null
            },
            "tags": [
                {
                    "id": "5e6b35de50ef8153c2062f70",
                    "url": "https://rest_api_root/v1/project/tags/5e6b35de50ef8153c2062f70",
                    "name": "标签-1"
                }
            ],
            "participants": [
                {
                    "id": "a0417f68e846aae315c85d24643678a9",
                    "url": "https://rest_api_root/v1/participants/a0417f68e846aae315c85d24643678a9?principal_type=work_item&principal_id=5edca524cad2fa112b06305c",
                    "type": "user",
                    "user": {
                        "id": "a0417f68e846aae315c85d24643678a9",
                        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                        "name": "john",
                        "display_name": "John",
                        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
                    }
                }
            ],
            "public_image_token": "73UNZyxnxUVvzKXe6KMs7ZUsEaTx3AGaBP3-Y9GE-Df",
            "created_at": 1583290300,
            "created_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "updated_at": 1583290300,
            "updated_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "is_archived": 0,
            "is_deleted": 0
        },
        {
            "id": "5edfa3b67463c1ee626c0979",
            "url": "https://rest_api_root/v1/project/work_items/5edfa3b67463c1ee626c0979",
            "project": {
                "id": "5eb623f6a70571487ea47000",
                "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
                "identifier": "SCR",
                "name": "Scrum项目",
                "type": "scrum",
                "is_archived": 0,
                "is_deleted": 0
            },
            "identifier": "SCR-2",
            "title": "这是一个特性",
            "type": "feature",
            "start_at": 1583290309,
            "end_at": 1583290347,
            "parent_id": "5edca524cad2fa112b06305c",
            "short_id": "a9WqLmTO",
            "html_url": "https://yctech.pingcode.com/pjm/workitems/a9WqLmTO",
            "parent": {
                "id": "5edca524cad2fa112b06305c",
                "url": "https://rest_api_root/v1/project/work_items/5edca524cad2fa112b06305c",
                "identifier": "SCR-1",
                "title": "这是一个史诗",
                "type": "epic",
                "start_at": 1583290309,
                "end_at": 1583290347,
                "parent_id": null,
                "properties": {
                    "prop_a": "prop_a_value",
                    "prop_b": "prop_b_value",
                    "risk": null,
                    "business_value": null,
                    "effort": null,
                    "backlog_type": null,
                    "backlog_from": null
                }
            },
            "assignee": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "versions": [
                {
                   "id": "5eb623487ea47001f6a70571",
                   "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/versions/5eb623487ea47001f6a70571",
                   "name": "1.0.1",
                   "start_at": 1565255712,
                   "end_at": 1565255879,
                   "stage": {
                       "id": "5f44a8f8bb347b14b49624a1",
                       "url": "https://rest_api_root/v1/project/stages/5f44a8f8bb347b14b49624a1",
                       "name": "未开始",
                       "type": "pending",
                       "color": "#FA8888"
                   }
                }
            ],
            "sprint": null,
            "state": {
                "id": "5c9b35de90ad7153c2062f18",
                "url": "https://rest_api_root/v1/project/work_item_states/5c9b35de90ad7153c2062f18",
                "name": "新建",
                "type": "pending",
                "color": "#ff7575"
            },
            "priority": {
                "id": "5eb623f6a70571487ea47111",
                "url": "https://rest_api_root/v1/project/work_item_priorities/5eb623f6a70571487ea47111",
                "name": "最高"
            },
            "board": null,
            "entry": null,
            "swimlane": null,
            "phase": null,
            "description": "这是一个特性的描述",
            "completed_at": 1583290347,
            "story_points": null,
            "estimated_workload": 1,
            "remaining_workload": 1,
            "properties": {
                "prop_a": "prop_a_value",
                "prop_b": "prop_b_value",
                "risk": null,
                "business_value": null,
                "effort": null,
                "backlog_type": null,
                "backlog_from": null
            },
            "tags": [
                {
                    "id": "5e6b35de50ef8153c2062f70",
                    "url": "https://rest_api_root/v1/project/tags/5e6b35de50ef8153c2062f70",
                    "name": "标签-1"
                }
            ],
            "participants": [
                {
                    "id": "a0417f68e846aae315c85d24643678a9",
                    "url": "https://rest_api_root/v1/participants/a0417f68e846aae315c85d24643678a9?principal_type=work_item&principal_id=5edfa3b67463c1ee626c0979",
                    "type": "user",
                    "user": {
                        "id": "a0417f68e846aae315c85d24643678a9",
                        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                        "name": "john",
                        "display_name": "John",
                        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
                    }
                }
            ],
            "public_image_token": "73UNZyxnxUVvzKXe6KMs7ZUsEaTx3AGaBP3-Y9GE-fC",
            "created_at": 1583290300,
            "created_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "updated_at": 1583290300,
            "updated_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "is_archived": 1,
            "is_deleted": 1
        },
        {
            "id": "5edca112b06305c524cad2fa",
            "url": "https://rest_api_root/v1/project/work_items/5edca112b06305c524cad2fa",
            "project": {
                "id": "5eb623f6a70571487ea47000",
                "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
                "identifier": "SCR",
                "name": "Scrum项目",
                "type": "scrum",
                "is_archived": 0,
                "is_deleted": 0
            },
            "identifier": "SCR-3",
            "title": "这是一个用户故事",
            "type": "story",
            "start_at": 1583290309,
            "end_at": 1583290347,
            "parent_id": "5edfa3b67463c1ee626c0979",
            "short_id": "b9WqLmTO",
            "html_url": "https://yctech.pingcode.com/pjm/workitems/b9WqLmTO",
            "parent": {
                "id": "5edfa3b67463c1ee626c0979",
                "url": "https://rest_api_root/v1/project/work_items/5edfa3b67463c1ee626c0979",
                "identifier": "SCR-2",
                "title": "这是一个特性",
                "type": "feature",
                "start_at": 1583290309,
                "end_at": 1583290347,
                "parent_id": "5edca524cad2fa112b06305g",
                "properties": {
                    "prop_a": "prop_a_value",
                    "prop_b": "prop_b_value",
                    "risk": null,
                    "business_value": null,
                    "effort": null,
                    "backlog_type": null,
                    "backlog_from": null
                }
            },
            "assignee": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "versions": [
                {
                   "id": "5eb623487ea47001f6a70571",
                   "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/versions/5eb623487ea47001f6a70571",
                   "name": "1.0.1",
                   "start_at": 1565255712,
                   "end_at": 1565255879,
                   "stage": {
                       "id": "5f44a8f8bb347b14b49624a1",
                       "url": "https://rest_api_root/v1/project/stages/5f44a8f8bb347b14b49624a1",
                       "name": "未开始",
                       "type": "pending",
                       "color": "#FA8888"
                   }
                }
            ],
            "sprint": {
                "id": "5ecf7b74eaab845a2aa53138",
                "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/sprints/5ecf7b74eaab845a2aa53138",
                "name": "Sprint 1",
                "start_at": 1589791860,
                "end_at": 1589791860,
                "status": "completed"
            },
            "state": {
                "id": "5c9b35de90ad7153c2062f18",
                "url": "https://rest_api_root/v1/project/work_item_states/5c9b35de90ad7153c2062f18",
                "name": "新建",
                "type": "pending",
                "color": "#ff7575"
            },
            "priority": {
                "id": "5eb623f6a70571487ea47111",
                "url": "https://rest_api_root/v1/project/work_item_priorities/5eb623f6a70571487ea47111",
                "name": "最高"
            },
            "board": null,
            "entry": null,
            "swimlane": null,
            "phase": null,
            "description": "这是一个用户故事的描述",
            "completed_at": 1583290347,
            "story_points": 1,
            "estimated_workload": 1,
            "remaining_workload": 1,
            "properties": {
                "prop_a": "prop_a_value",
                "prop_b": "prop_b_value",
                "risk": null,
                "backlog_type": null,
                "backlog_from": null
            },
            "tags": [
                {
                    "id": "5e6b35de50ef8153c2062f70",
                    "url": "https://rest_api_root/v1/project/tags/5e6b35de50ef8153c2062f70",
                    "name": "标签-1"
                }
            ],
            "participants": [
                {
                    "id": "a0417f68e846aae315c85d24643678a9",
                    "url": "https://rest_api_root/v1/participants/a0417f68e846aae315c85d24643678a9?principal_type=work_item&principal_id=5edca112b06305c524cad2fa",
                    "type": "user",
                    "user": {
                        "id": "a0417f68e846aae315c85d24643678a9",
                        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                        "name": "john",
                        "display_name": "John",
                        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
                    }
                }
            ],
            "public_image_token": "73UNZyxnxUVvzKXe6KMs7ZUsEaTx3AGaBP3-Y9GE-Hm",
            "created_at": 1583290300,
            "created_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "updated_at": 1583290300,
            "updated_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "is_archived": 0,
            "is_deleted": 0
        },
        {
            "id": "5edca5d2fa112b06305c24ca",
            "url": "https://rest_api_root/v1/project/work_items/5edca5d2fa112b06305c24ca",
            "project": {
                "id": "6375cc81e3004de4ea14aa52",
                "url": "http://rest_api_root/v1/project/projects/6375cc81e3004de4ea14aa52",
                "identifier": "WTF",
                "name": "瀑布项目",
                "type": "waterfall",
                "is_archived": 0,
                "is_deleted": 0
            },
            "identifier": "WTF-2",
            "title": "这是一个瀑布项目下需求类型的工作项",
            "type": "630da48bc9443b1aa94ce3ea",
            "start_at": 1583290309,
            "end_at": 1583290347,
            "short_id": "c9WqLmTO",
            "html_url": "https://yctech.pingcode.com/pjm/workitems/c9WqLmTO",
            "assignee": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "versions": null,
            "sprint": null,
            "state": {
                "id": "5c9b35de90ad7153c2062f18",
                "url": "https://rest_api_root/v1/project/work_item_states/5c9b35de90ad7153c2062f18",
                "name": "新建",
                "type": "pending",
                "color": "#ff7575"
            },
            "priority": {
                "id": "5eb623f6a70571487ea47111",
                "url": "https://rest_api_root/v1/project/work_item_priorities/5eb623f6a70571487ea47111",
                "name": "最高"
            },
            "board": null,
            "entry": null,
            "swimlane": null,
            "phase": {
                "id": "63761fee31caaf77189816b4",
                "url": "http://rest_api_root/v1/project/projects/63761fee31caaf7718981698/phases/63761fee31caaf77189816b4",
                "title": "这是一个阶段",
                "identifier": "WTF-1"
            },
            "description": "这是一个瀑布项目下需求类型的工作项描述",
            "completed_at": 1583290347,
            "story_points": null,
            "estimated_workload": 1,
            "remaining_workload": 1,
            "properties": {
                "prop_a": "prop_a_value",
                "prop_b": "prop_b_value",
                "backlog_type": null,
                "backlog_from": null
            },
            "tags": [
                {
                    "id": "5e6b35de50ef8153c2062f70",
                    "url": "https://rest_api_root/v1/project/tags/5e6b35de50ef8153c2062f70",
                    "name": "标签-1"
                }
            ],
            "participants": [
                {
                    "id": "a0417f68e846aae315c85d24643678a9",
                    "url": "https://rest_api_root/v1/participants/a0417f68e846aae315c85d24643678a9?principal_type=work_item&principal_id=5edca112b06305c524cad2fa",
                    "type": "user",
                    "user": {
                        "id": "a0417f68e846aae315c85d24643678a9",
                        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                        "name": "john",
                        "display_name": "John",
                        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
                    }
                }
            ],
            "public_image_token": "73UNZyxnxUVvzKXe6KMs7ZUsEaTx3AGaBP3-Y9GE-Ki",
            "created_at": 1583290300,
            "created_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "updated_at": 1583290300,
            "updated_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "is_archived": 0,
            "is_deleted": 0
        },
        {
            "id": "6efca131b06329c524cdd2fb",
            "url": "https://rest_api_root/v1/project/work_items/6efca131b06329c524cdd2fb",
            "project": {
                "id": "5eb623f6a70571487ea47004",
                "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47004",
                "identifier": "HBR",
                "name": "Hybrid项目",
                "type": "hybrid",
                "is_archived": 0,
                "is_deleted": 0
            },
            "identifier": "HBR-1",
            "title": "这是一个用户故事",
            "type": "story",
            "start_at": 1583290309,
            "end_at": 1583290347,
            "parent_id": "5edfa3b67463c1ee626c0980",
            "short_id": "e9WqLmTO",
            "html_url": "https://yctech.pingcode.com/pjm/workitems/e9WqLmTO",
            "parent": {
                "id": "5edfa3b67463c1ee626c0980",
                "url": "https://rest_api_root/v1/project/work_items/5edfa3b67463c1ee626c0980",
                "identifier": "HBR-2",
                "title": "这是一个特性",
                "type": "feature",
                "start_at": 1583290309,
                "end_at": 1583290347,
                "parent_id": null,
                "properties": {
                   "prop_a": "prop_a_value",
                   "prop_b": "prop_b_value",
                   "risk": null,
                   "business_value": null,
                   "effort": null,
                   "backlog_type": null,
                   "backlog_from": null
                }
            },
            "assignee": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "versions": [
                {
                   "id": "5eb623487ea47001f6a70571",
                   "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/versions/5eb623487ea47001f6a70571",
                   "name": "1.0.1",
                   "start_at": 1565255712,
                   "end_at": 1565255879,
                   "stage": {
                       "id": "5f44a8f8bb347b14b49624a1",
                       "url": "https://rest_api_root/v1/project/stages/5f44a8f8bb347b14b49624a1",
                       "name": "未开始",
                       "type": "pending",
                       "color": "#FA8888"
                   }
                }
            ],
            "sprint": {
                "id": "5ecf7b74eaab845a2aa53139",
                "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47004/sprints/5ecf7b74eaab845a2aa53139",
                "name": "Sprint 1",
                "start_at": 1589791860,
                "end_at": 1589791860,
                "status": "completed"
            },
            "state": {
                "id": "5c9b35de90ad7153c2062f18",
                "url": "https://rest_api_root/v1/project/work_item_states/5c9b35de90ad7153c2062f18",
                "name": "新建",
                "type": "pending",
                "color": "#ff7575"
            },
            "priority": {
                "id": "5eb623f6a70571487ea47111",
                "url": "https://rest_api_root/v1/project/work_item_priorities/5eb623f6a70571487ea47111",
                "name": "最高"
            },
            "board": {
                "id": "5eb623f6a70571487ea47333",
                "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47004/boards/5eb623f6a70571487ea47333",
                "name": "kanban",
                "work_item_types": [
                   "epic",
                   "feature",
                   "issue",
                   "story"
                ]
            },
            "entry": {
                "id": "5ee1c4fac5b3c51206f0a862",
                "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47004/boards/5eb623f6a70571487ea47333/entries/5ee1c4fac5b3c51206f0a862",
                "name": "需求池"
            },
            "swimlane": {
                "id": "5ee1c4fac5b3c51206f0a867",
                "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47004/boards/5eb623f6a70571487ea47333/swimlanes/5ee1c4fac5b3c51206f0a867",
                "name": "默认泳道"
            },
            "phase": {
                "id": "63761fee31caaf77189816b5",
                "url": "http://rest_api_root/v1/project/projects/5eb623f6a70571487ea47004/phases/63761fee31caaf77189816b5",
                "title": "这是一个阶段",
                "identifier": "WTF-1"
            },
            "description": "<p>这是一个用户故事的描述<p><img src=\"https://atlas.pingcode.com/files/public/689a9d124df436ef91923a3a\" originUrl=\"https://atlas.pingcode.com/files/public/689a9d124df436ef91923a3a\" alt=\"图片.png\" size=\"35460\" style=\"text-align: center;\" />",
            "completed_at": 1583290347,
            "story_points": 1,
            "estimated_workload": 1,
            "remaining_workload": 1,
            "properties": {
                "prop_a": "prop_a_value",
                "prop_b": "prop_b_value",
                "risk": null,
                "backlog_type": null,
                "backlog_from": null
            },
            "tags": [
                {
                    "id": "5e6b35de50ef8153c2062f70",
                    "url": "https://rest_api_root/v1/project/tags/5e6b35de50ef8153c2062f70",
                    "name": "标签-1"
                }
            ],
            "participants": [
                {
                    "id": "a0417f68e846aae315c85d24643678a9",
                    "url": "https://rest_api_root/v1/participants/a0417f68e846aae315c85d24643678a9?principal_type=work_item&principal_id=5edca112b06305c524cad2fa",
                    "type": "user",
                    "user": {
                        "id": "a0417f68e846aae315c85d24643678a9",
                        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                        "name": "john",
                        "display_name": "John",
                        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
                    }
                }
            ],
            "public_image_token": "73UNZyxnxUVvzKXe6KMs7ZUsEaTx3AGaBP3-Y9GE-Ac",
            "created_at": 1583290300,
            "created_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "updated_at": 1583290300,
            "updated_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "is_archived": 0,
            "is_deleted": 0
        }
    ]
}
```

---

### 获取工作项属性列表

**`GET /v1/project/work_item/properties?project_id={project_id}&work_item_type_id={work_item_type_id}`**

**查询参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `project_id` | String | 项目的id。 |
| `work_item_type_id` | String | 工作项类型的id。工作项类型分为9种系统类型和一些自定义类型。系统类型包括：史诗、特性、用户故事、阶段、里程碑、需求、任务、缺陷和事务。可以通过获取全部工作项类型列表获得。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "severity",
            "url": "https://rest_api_root/v1/project/work_item/work_item_properties/severity",
            "name": "严重程度",
            "type": "select",
            "options": [
                {
                    "_id": "5efb1859110533727a82c603",
                    "text": "严重"
                },
                {
                    "_id": "5efb1859110533727a82c604",
                    "text": "一般"
                }
            ]
        },
        {
            "id": "identifier",
            "url": "https://rest_api_root/v1/project/work_item_properties/identifier",
            "name": "编号",
            "type": "text",
            "options": null
        }
    ]
}
```

---

### 获取工作项工时列表

**`GET /v1/project/workloads`**

**查询参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `project_id` *(可选)* | String | 项目的id。使用该参数过滤数据时，登记日期查询的起始时间和登记日期查询的结束时间的跨度最大为3个月。 |
| `work_item_id` *(可选)* | String | 工作项的id。 |
| `start_at` *(可选)* | String | 登记日期查询的起始时间。 |
| `end_at` *(可选)* | String | 登记日期查询的结束时间。 |
| `user_id` *(可选)* | String | 登记人的id。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "5f168f764eba01a5278b87cd",
            "url": "https://rest_api_root/v1/project/workloads/5f168f764eba01a5278b87cd",
            "project": {
                "id": "5eb623f6a70571487ea47000",
                "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
                "identifier": "SCR",
                "name": "Scrum项目",
                "type": "scrum",
                "is_archived": 0,
                "is_deleted": 0
            },
            "work_item": {
                "id": "5edca524cad2fa1125cb0630",
                "url": "https://rest_api_root/v1/project/work_items/5edca524cad2fa1125cb0630",
                "identifier": "SCR-5",
                "title": "这是一个缺陷",
                "type": "bug",
                "start_at": 1583290309,
                "end_at": 1583290347,
                "parent_id": "5edca524cad2fa1125cb0629",
                "short_id": "c9WqLmTO",
                "html_url": "https://yctech.pingcode.com/pjm/workitems/c9WqLmTO",
                "properties": {
                    "prop_a": "prop_a_value",
                    "prop_b": "prop_b_value"
                }
            },
            "type": {
                "id": "5a86eaf6a72585327ea46fge0",
                "url": "https://rest_api_root/v1/project/workload_types/5a86eaf6a72585327ea46fge0",
                "name": "研发",
                "color": "#56ABF"
            },
            "duration": 8,
            "description": "这是一个工时",
            "report_at": 1593290347,
            "report_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "created_at": 1593290347
        }
    ]
}
```

---

### 获取工作项工时类型列表

**`GET /v1/project/workload_types`**

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "5a86eaf6a72585327ea46fge0",
            "url": "https://rest_api_root/v1/project/workload_types/5a86eaf6a72585327ea46fge0",
            "name": "研发",
            "color": "#56ABFB"
        }
    ]
}
```

---

### 获取工作项标签列表

**`GET /v1/project/work_item/tags`**

**查询参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `project_id` | String | 项目的id。 |
| `name` *(可选)* | String | 标签的名称。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "5e6b35de50ef8153c2062f70",
            "url": "https://rest_api_root/v1/project/work_item_tags/5e6b35de50ef8153c2062f70",
            "name": "标签-1"
        }
    ]
}
```

---

### 获取工作项流转记录列表

**`GET /v1/project/work_items/{work_item_id}/transition_histories`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `work_item_id` | String | 工作项的id。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "5e6b35de50ef8153c2062f70",
            "url": "https://rest_api_root/v1/project/work_items/5edca524cad2fa1125cb0630/transition_histories/5e6b35de50ef8153c2062f70",
            "work_item": {
                "id": "5edca524cad2fa1125cb0630",
                "url": "https://rest_api_root/v1/project/work_items/5edca524cad2fa1125cb0630",
                "identifier": "SCR-5",
                "title": "这是一个缺陷",
                "type": "bug",
                "start_at": 1674493200,
                "end_at": 1674493200,
                "parent_id": "5edca524cad2fa112b05105c",
                "short_id": "c9WqLmTO",
                "html_url": "https://yctech.pingcode.com/pjm/workitems/c9WqLmTO",
                "properties": {
                   "prop_a": "prop_a_value",
                   "prop_b": "prop_b_value"
                }
            },
            "from_state": {
                "id": "5c9b35de90ad7153c2062f18",
                "url": "https://rest_api_root/v1/project/work_item_states/5c9b35de90ad7153c2062f18",
                "name": "新建",
                "type": "pending",
                "color": "#56ABFB"
            },
            "to_state": {
                "id": "5ef85b1e9481936604da7f4c",
                "url": "https://rest_api_root/v1/project/work_item_states/5ef85b1e9481936604da7f4c",
                "name": "进行中",
                "type": "in_progress",
                "color": "#F6C659"
            },
            "created_at": 1674528614,
            "created_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            }
        }
    ]
}
```

---

### 获取工作项状态列表

**`GET /v1/project/work_item/states?project_id={project_id}&work_item_type_id={work_item_type_id}`**

**查询参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `project_id` | String | 项目的id。 |
| `work_item_type_id` | String | 工作项类型的id。工作项类型分为9种系统类型和一些自定义类型。系统类型包括：史诗、特性、用户故事、阶段、里程碑、需求、任务、缺陷和事务。可以通过获取全部工作项类型列表获得。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "5c9b35de90ad7153c2062f18",
            "url": "https://rest_api_root/v1/project/work_item_states/5c9b35de90ad7153c2062f18",
            "name": "新建",
            "type": "pending",
            "color": "#ff7575"
        }
    ]
}
```

---

### 获取工作项类型列表

**`GET /v1/project/work_item/types?project_id={project_id}`**

**查询参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `project_id` | String | 项目的id。 |

**响应示例（scrum项目）：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 6,
    "values": [
        {
            "id": "epic",
            "url": "http://rest_api_root/v1/project/work_item_types/epic",
            "name": "史诗",
            "group": "requirement"
        },
        {
            "id": "feature",
            "url": "http://rest_api_root/v1/project/work_item_types/feature",
            "name": "特性",
            "group": "requirement"
        },
        {
            "id": "story",
            "url": "http://rest_api_root/v1/project/work_item_types/story",
            "name": "用户故事",
            "group": "requirement"
        },
        {
            "id": "task",
            "url": "http://rest_api_root/v1/project/work_item_types/task",
            "name": "任务",
            "group": "task"
        },
        {
            "id": "bug",
            "url": "http://rest_api_root/v1/project/work_item_types/bug",
            "name": "缺陷",
            "group": "bug"
        }
    ]
}
```

**响应示例（kanban项目）：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 6,
    "values": [
        {
            "id": "epic",
            "url": "http://rest_api_root/v1/project/work_item_types/epic",
            "name": "史诗",
            "group": "requirement"
        },
        {
            "id": "feature",
            "url": "http://rest_api_root/v1/project/work_item_types/feature",
            "name": "特性",
            "group": "requirement"
        },
        {
            "id": "story",
            "url": "http://rest_api_root/v1/project/work_item_types/story",
            "name": "用户故事",
            "group": "requirement"
        },
        {
            "id": "task",
            "url": "http://rest_api_root/v1/project/work_item_types/task",
            "name": "任务",
            "group": "task"
        },
        {
            "id": "bug",
            "url": "http://rest_api_root/v1/project/work_item_types/bug",
            "name": "缺陷",
            "group": "bug"
        },
        {
            "id": "issue",
            "url": "http://rest_api_root/v1/project/work_item_types/issue",
            "name": "事务",
            "group": "issue"
        }
    ]
}
```

**响应示例（waterfall项目）：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 6,
    "values": [
        {
            "id": "630da48bc9443b1aa94ce3ea",
            "url": "https://rest_api_root/v1/project/work_item_types/630da48bc9443b1aa94ce3ea",
            "name": "需求",
            "group": "requirement"
        },
        {
            "id": "task",
            "url": "https://rest_api_root/v1/project/work_item_types/task",
            "name": "任务",
            "group": "task"
        },
        {
            "id": "bug",
            "url": "https://rest_api_root/v1/project/work_item_types/bug",
            "name": "缺陷",
            "group": "bug"
        },
        {
            "id": "6385c650fef18f2d7222d15d",
            "url": "https://rest_api_root/v1/project/work_item_types/6385c650fef18f2d7222d15d",
            "name": "自定义",
            "group": "issue"
        },
        {
            "id": "630da48bc9443b1aa94ce3ee",
            "url": "https://rest_api_root/v1/project/work_item_types/630da48bc9443b1aa94ce3ee",
            "name": "阶段",
            "group": "plan"
        },
        {
            "id": "630da48bc9443b1aa94ce3ef",
            "url": "https://rest_api_root/v1/project/work_item_types/630da48bc9443b1aa94ce3ef",
            "name": "里程碑",
            "group": "plan"
        }
    ]
}
```

**响应示例（hybrid项目）：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total":9,
    "values": [
        {
            "id": "epic",
            "url": "http://rest_api_root/v1/project/work_item_types/epic",
            "name": "史诗",
            "group": "requirement"
        },
        {
            "id": "feature",
            "url": "http://rest_api_root/v1/project/work_item_types/feature",
            "name": "特性",
            "group": "requirement"
        },
        {
            "id": "story",
            "url": "http://rest_api_root/v1/project/work_item_types/story",
            "name": "用户故事",
            "group": "requirement"
        },
        {
            "id": "task",
            "url": "http://rest_api_root/v1/project/work_item_types/task",
            "name": "任务",
            "group": "task"
        },
        {
            "id": "bug",
            "url": "http://rest_api_root/v1/project/work_item_types/bug",
            "name": "缺陷",
            "group": "bug"
        },
        {
            "id": "issue",
            "url": "http://rest_api_root/v1/project/work_item_types/issue",
            "name": "事务",
            "group": "issue"
        },
        {
            "id": "630da48bc9443b1aa94ce3ea",
            "url": "https://rest_api_root/v1/project/work_item_types/630da48bc9443b1aa94ce3ea",
            "name": "需求",
            "group": "requirement"
        },
        {
            "id": "630da48bc9443b1aa94ce3ee",
            "url": "https://rest_api_root/v1/project/work_item_types/630da48bc9443b1aa94ce3ee",
            "name": "阶段",
            "group": "plan"
        },
        {
            "id": "630da48bc9443b1aa94ce3ef",
            "url": "https://rest_api_root/v1/project/work_item_types/630da48bc9443b1aa94ce3ef",
            "name": "里程碑",
            "group": "plan"
        },
        {
            "id": "6385c650fef18f2d7222d15d",
            "url": "https://rest_api_root/v1/project/work_item_types/6385c650fef18f2d7222d15d",
            "name": "自定义",
            "group": "issue"
        }
    ]
}
```

---

### 部分更新一个工作项

**`PATCH /v1/project/work_items/{work_item_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `work_item_id` | String | 工作项的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `title` *(可选)* | String | 工作项的标题。 |
| `description` *(可选)* | String | 工作项的描述。 |
| `start_at` *(可选)* | Number | 工作项的开始时间。当工作项类型为里程碑时，该参数无效。 |
| `end_at` *(可选)* | Number | 工作项的截止时间。 |
| `priority_id` *(可选)* | String | 工作项优先级的id。 |
| `state_id` *(可选)* | String | 工作项状态的id。工作项状态的id在设置时必须同时满足工作项类型的工作项状态方案和工作项状态流转的状态值才能成功完成设置。工作项状态方案可以通过获取工作项状态方案列表获取。工作项状态流转则可以通过获取状态方案中的工作项状态流转列表获取。 |
| `assignee_id` *(可选)* | String | 工作项负责人的id。 |
| `parent_id` *(可选)* | String | 工作项的父工作项的id。当前工作项类型支持的父工作类型包含parent_id对应的工作项类型时，该参数有效。具体配置见属性与视图子工作项配置。 |
| `version_ids` *(可选)* | String[] | 所属发布的id列表。 |
| `board_id` *(可选)* | String | 看板的id。该字段只有项目类型为kanban和hybrid时有效。 |
| `entry_id` *(可选)* | String | 看板栏的id。该字段只有项目类型为kanban和hybrid时有效。 |
| `swimlane_id` *(可选)* | String | 泳道的id。该字段只有项目类型为kanban和hybrid时有效。 |
| `phase_id` *(可选)* | String | 所属计划的id。该字段只有项目类型为waterfall和hybrid时有效。 |
| `story_points` *(可选)* | Number | 工作项的故事点。当工作项的属性在视图中配置了故事点属性时，该参数生效。故事点数值必须是大于等于0的整数或最多包含一位小数的正数。 |
| `estimated_workload` *(可选)* | Number | 工作项的预估工时。 |
| `remaining_workload` *(可选)* | Number | 工作项的剩余工时。 |
| `properties` *(可选)* | Object | 工作项属性的键值对集合，需要注意的是，当前工作项类型对应的工作项属性方案需要包含这些工作项属性，例如工作项属性方案中包含prop_a和prop_b。 |
| `properties.prop_a` *(可选)* | Object | 工作项属性prop_a。 |
| `properties.prop_b` *(可选)* | Object | 工作项属性prop_b。 |

**响应示例：**

```json
{
    "id": "5edca112b06305c524cad2fa",
    "url": "https://rest_api_root/v1/project/work_items/5edca112b06305c524cad2fa",
    "project": {
        "id": "5eb623f6a70571487ea47000",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
        "identifier": "SCR",
        "name": "Scrum项目",
        "type": "scrum",
        "is_archived": 0,
        "is_deleted": 0
    },
    "identifier": "SCR-3",
    "title": "这是一个用户故事",
    "type": "story",
    "start_at": 1583290309,
    "end_at": 1583290347,
    "parent_id": "5edfa3b67463c1ee626c0979",
    "short_id": "b9WqLmTO",
    "html_url": "https://yctech.pingcode.com/pjm/workitems/b9WqLmTO",
    "parent": {
        "id": "5edfa3b67463c1ee626c0979",
        "url": "https://rest_api_root/v1/project/work_items/5edfa3b67463c1ee626c0979",
        "identifier": "SCR-2",
        "title": "这是一个特性",
        "type": "feature",
        "start_at": 1583290309,
        "end_at": 1583290347,
        "parent_id": "5edca524cad2fa112b06306c",
        "properties": {
            "prop_a": "prop_a_value",
            "prop_b": "prop_b_value",
            "risk": null,
            "business_value": null,
            "effort": null,
            "backlog_type": null,
            "backlog_from": null
        }
    },
    "assignee": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "versions": [
       {
          "id": "5eb623487ea47001f6a70571",
          "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/versions/5eb623487ea47001f6a70571",
          "name": "1.0.1",
          "start_at": 1565255712,
          "end_at": 1565255879,
          "stage": {
              "id": "5f44a8f8bb347b14b49624a1",
              "url": "https://rest_api_root/v1/project/stages/5f44a8f8bb347b14b49624a1",
              "name": "未开始",
              "type": "pending",
              "color": "#FA8888"
          }
       }
    ],
    "sprint": {
        "id": "5ecf7b74eaab845a2aa53138",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/sprints/5ecf7b74eaab845a2aa53138",
        "name": "Sprint 1",
        "start_at": 1589791860,
        "end_at": 1589791860,
        "status": "completed"
    },
    "state": {
        "id": "5c9b35de90ad7153c2062f18",
        "url": "https://rest_api_root/v1/project/work_item_states/5c9b35de90ad7153c2062f18",
        "name": "新建",
        "type": "pending",
        "color": "#ff7575"
    },
    "priority": {
        "id": "5eb623f6a70571487ea47111",
        "url": "https://rest_api_root/v1/project/work_item_priorities/5eb623f6a70571487ea47111",
        "name": "最高"
    },
    "board":null,
    "entry": null,
    "swimlane": null,
    "phase": null,
    "story_points": 1,
    "estimated_workload": 1,
    "remaining_workload": 1,
    "description": "这是一个用户故事的描述",
    "completed_at": 1583290347,
    "properties": {
        "prop_a": "prop_a_value",
        "prop_b": "prop_b_value",
        "risk": null,
        "backlog_type": null,
        "backlog_from": null
    },
    "tags": [],
    "participants": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "url": "https://rest_api_root/v1/participants/a0417f68e846aae315c85d24643678a9?principal_type=work_item&principal_id=5edca112b06305c524cad2fa",
            "type": "user",
            "user": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            }
        }
    ],
    "created_at": 1583290300,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "updated_at": 1583293300,
    "updated_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "is_archived": 0,
    "is_deleted": 0
}
```

---

### 部分更新一个工作项交付目标

**`PATCH /v1/project/deliverables/{deliverable_target_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `deliverable_target_id` | String | 工作项交付目标id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `work_item_id` *(可选)* | String | 工作项的id。 |
| `name` *(可选)* | String | 工作项交付目标的名称。 |
| `content_type` *(可选)* | String | 工作项交付物的类型。只支持link。attachment类型的交付物通过向交付目标中上传一个文件接口上传附件。 |
| `content` *(可选)* | String | 工作项交付物。当工作项交付物的类型是link时，工作项交付物为包含name和href的对象，例如：{ &quot;name&quot;: &quot;链接工作项交付目标&quot;,  &quot;href&quot;: &quot;https://rest_api_root/wiki/spaces/public/pages/6472e6f2f1968d3fdb0aba15&quot; }。当工作项交付物不为空时，参数必须包含交付物类型。 |

**响应示例：**

```json
{
    "id": "63761fee31caaf7718981876",
    "url": "https://rest_api_root/v1/project/deliverables/63761fee31caaf7718981876",
    "name": "阶段交付目标",
    "content_type": "link",
    "content": {
        "name": "PingCode",
        "href": "https://www.pingcode.com"
    },
    "project": {
        "id": "6375cc81e3004de4ea14aa52",
        "url": "https://rest_api_root/v1/project/projects/6375cc81e3004de4ea14aa52",
        "identifier": "WTF",
        "name": "瀑布项目",
        "type": "waterfall",
        "is_archived": 0,
        "is_deleted": 0
    },
    "work_item": {
        "id": "63761fee31caaf77189816b4",
        "url": "https://rest_api_root/v1/project/work_items/63761fee31caaf77189816b4",
        "identifier": "WTF-5",
        "title": "这是一个阶段",
        "type": "630da48bc9443b1aa94ce3ee",
        "start_at": 1583290309,
        "end_at": 1583290347,
        "parent_id": "5edca524cad2fa1125cb0629",
        "short_id": "c9WqLmTO",
        "html_url": "https://yctech.pingcode.com/pjm/workitems/c9WqLmTO",
        "properties": {
            "prop_a": "prop_a_value",
            "prop_b": "prop_b_value"
        }
    },
    "created_at": 1668685806,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "updated_at": 1668687806,
    "updated_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "is_archived": 0,
    "is_deleted": 0
}
```

---

### 部分更新一个工作项工时

**`PATCH /v1/project/workloads/{workload_id}`**

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `workload_type_id` *(可选)* | String | 工时类型的id。 |
| `duration` *(可选)* | Number | 时长。时长的单位是小时。数值可以是为0-24之间，最多包含一位小数的正数。 |
| `report_at` *(可选)* | Number | 登记日期。该值为十位数字组成的时间戳，会被转换为该时间当天的零点零分零秒。 |
| `report_by` *(可选)* | String | 登记人，企业鉴权时必填。个人鉴权时不需要传递，即使传递了也会被忽略。 |
| `description` *(可选)* | String | 工时的描述信息。 |

**响应示例：**

```json
{
    "id": "5f168f764eba01a5278b87cd",
    "url": "https://rest_api_root/v1/project/workloads/5f168f764eba01a5278b87cd",
    "project": {
        "id": "5eb623f6a70571487ea47000",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
        "identifier": "SCR",
        "name": "Scrum项目",
        "type": "scrum",
        "is_archived": 0,
        "is_deleted": 0
    },
    "work_item": {
        "id": "5edca524cad2fa1125cb0630",
        "url": "https://rest_api_root/v1/project/work_items/5edca524cad2fa1125cb0630",
        "identifier": "SCR-5",
        "title": "这是一个缺陷",
        "type": "bug",
        "start_at": 1583290309,
        "end_at": 1583290347,
        "parent_id": "5edca524cad2fa1125cb0629",
        "short_id": "c9WqLmTO",
        "html_url": "https://yctech.pingcode.com/pjm/workitems/c9WqLmTO",
        "properties": {
            "prop_a": "prop_a_value",
            "prop_b": "prop_b_value"
        }
    },
    "type": {
        "id": "5a86eaf6a72585327ea46fge0",
        "url": "https://rest_api_root/v1/project/workload_types/5a86eaf6a72585327ea46fge0",
        "name": "研发",
        "color": "#56ABF"
    },
    "duration": 8,
    "description": "这是一个工时",
    "report_at": 1593290347,
    "report_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "created_at": 1593290347
}
```

---

## 工作项配置

### 创建一个工作项属性

**`POST /v1/project/work_item_properties`**

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `name` | String | 工作项属性的名称。在一个企业中这个名称是唯一的。 |
| `type` | String | 工作项属性的类型。 |
| `options` *(可选)* | Object[] | 选择项列表。当工作项属性类型为select,multi_select,cascade_select,cascade_multi_select时可选填此项。 |
| `options._id` *(可选)* | String | 选择项id。该值在选择项中是唯一的。 |
| `options.text` | String | 选择项内容。该值在选择项中是唯一的。 |
| `options.parent_id` *(可选)* | String | 选择项父级id。当属性类型为cascade_select,cascade_multi_select时，parent_id参数有效，用于构建级联类型选择项之间的父子关系，最多存在4级。 |

**响应示例（select）：**

```json
{
    "id": "severity",
    "url": "https://rest_api_root/v1/project/work_item_properties/severity",
    "name": "严重程度",
    "type": "select",
    "options": [
        {
            "_id": "5efb1859110533727a82c603",
            "text": "严重"
        },
        {
            "_id": "5efb1859110533727a82c604",
            "text": "一般"
        }
    ],
    "is_removable": true,
    "is_name_editable": true,
    "is_options_editable": true,
    "select_all_level": false,
    "display_all_level": false,
    "display_separator": null
}
```

**响应示例（cascade_select）：**

```json
{
    "id": "jiliandanxuan",
    "url": "https://rest_api_root/v1/project/work_item_properties/jiliandanxuan",
    "name": "级联单选",
    "type": "cascade_select",
    "options": [
        {
            "_id": "64b9f741eec7fd94e63b411e",
            "text": "父"
        },
        {
            "_id": "64b9f741eec7fd94e63b411f",
            "text": "子",
            "parent_id": "64b9f741eec7fd94e63b411e"
        }
    ],
    "is_removable": true,
    "is_name_editable": true,
    "is_options_editable": true,
    "select_all_level": false,
    "display_all_level": false,
    "display_separator": "/"
}
```

---

### 创建一个工作项标签

**`POST /v1/project/work_item_tags`**

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `name` | String | 标签的名称。在一个企业中这个名称是唯一的。 |

**响应示例：**

```json
{
    "id": "5e6b35de50ef8153c2062f70",
    "url": "https://rest_api_root/v1/project/work_item_tags/5e6b35de50ef8153c2062f70",
    "name": "标签-1",
    "color": "#56ABFB"
}
```

---

### 创建一个工作项状态

**`POST /v1/project/work_item_states`**

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `name` | String | 工作项状态的名称。在一个企业中这个名称是唯一的。 |
| `type` | String | 工作项状态的类型。 |

**响应示例：**

```json
{
    "id": "5c9b35de90ad7153c2062f18",
    "url": "https://rest_api_root/v1/project/work_item_states/5c9b35de90ad7153c2062f18",
    "name": "新建",
    "type": "pending",
    "color": "#ff7575",
    "is_system": false
}
```

---

### 创建一个工作项类型

**`POST /v1/project/work_item_types`**

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `name` | String | 工作项类型的名称。在一个企业中这个名称是唯一的。 |
| `group` | String | 工作项类型的分组。 |

**响应示例：**

```json
{
    "id": "630da48bc9443b1aa94ce3fc",
    "url": "https://rest_api_root/v1/project/work_item_types/630da48bc9443b1aa94ce3fc",
    "name": "功能缺陷",
    "group": "bug",
    "is_system": false
}
```

---

### 删除一个工作项标签

**`DELETE /v1/project/work_item_tags/{tag_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `tag_id` | String | 标签的id。 |

**响应示例：**

```json
{
    "id": "5e6b35de50ef8153c2062f70",
    "url": "https://rest_api_root/v1/project/work_item_tags/5e6b35de50ef8153c2062f70",
    "name": "标签-2",
    "color": "#56ABFB"
}
```

---

### 删除一个工作项类型

**`DELETE /v1/project/work_item_types/{work_item_type_id}`**

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `work_item_type_id` | String | 工作项类型的id。 |

**响应示例：**

```json
{
    "id": "630da48bc9443b1aa94ce3df",
    "url": "https://rest_api_root/v1/project/work_item_types/630da48bc9443b1aa94ce3df",
    "name": "非功能性需求",
    "group": "requirement",
    "is_system": false
}
```

---

### 向属性方案中添加一个工作项属性

**`POST /v1/project/work_item_property_plans/{property_plan_id}/work_item_properties`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `property_plan_id` | String | 工作项属性方案的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `property_id` | String | 工作项属性的id。 |

**响应示例：**

```json
{
    "id": "severity",
    "url": "https://rest_api_root/v1/project/work_item_property_plans/5f8a21f18ef715265de90c21/properties/severity",
    "property_plan": {
        "id": "5f8a21f18ef715265de90c21",
        "url": "https://rest_api_root/v1/project/work_item_property_plans/5f8a21f18ef715265de90c21",
        "project_type": "scrum",
        "work_item_type": "story"
    },
    "property": {
        "id": "severity",
        "url": "https://rest_api_root/v1/project/work_item_properties/severity",
        "name": "严重程度",
        "type": "select"
    }
}
```

---

### 向工作项类型方案中添加一个工作项类型

**`POST /v1/project/work_item_type_plans/{work_item_type_plan_id}/work_item_types`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `work_item_type_plan_id` | String | 工作项类型方案的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `work_item_type_id` | String | 工作项类型的id。 |

**响应示例：**

```json
{
    "id": "630da48bc9443b1aa94ce3ea",
    "url": "https://rest_api_root/v1/project/work_item_type_plans/65b0d9fd9ee456e672657e39/work_item_types/630da48bc9443b1aa94ce3ea",
    "type_plan": {
        "id": "6abb92f18ad725395df86c45",
        "url": "https://rest_api_root/v1/project/work_item_type_plans/65b0d9fd9ee456e672657e39",
        "project_type": "waterfall"
    },
    "type": {
        "id": "630da48bc9443b1aa94ce3ea",
        "url": "https://rest_api_root/v1/project/work_item_types/630da48bc9443b1aa94ce3ea",
        "name": "需求",
        "group": "requirement"
    },
    "sub_types": [
        {
           "id": "bug",
           "url": "https://rest_api_root/v1/project/work_item_types/bug",
           "name": "缺陷",
           "group": "bug"
        },
        {
           "id": "6385c650fef18f2d7222d15d",
           "url": "https://rest_api_root/v1/project/work_item_types/6385c650fef18f2d7222d15d",
           "name": "自定义",
           "group": "issue"
        }
    ]
}
```

---

### 向状态方案中添加一个工作项状态

**`POST /v1/project/work_item_state_plans/{state_plan_id}/work_item_states`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `state_plan_id` | String | 工作项状态方案的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `state_id` | String | 工作项状态的id。 |

**响应示例：**

```json
{
    "id": "5cc2062f189b35de90ad7153",
    "url": "https://rest_api_root/v1/project/work_item_state_plans/5c9b62f18ad715335de90c20/work_item_states/5c9b35de90ad7153c2062f18",
    "state_plan": {
        "id": "5c9b62f18ad715335de90c20",
        "url": "https://rest_api_root/v1/project/work_item_state_plans/5c9b62f18ad715335de90c20",
        "project_type": "scrum",
        "work_item_type": "story"
    },
    "state": {
        "id": "5c9b35de90ad7153c2062f18",
        "url": "https://rest_api_root/v1/project/work_item_states/5c9b35de90ad7153c2062f18",
        "name": "新建",
        "type": "pending",
        "color": "#ff7575"
    }
}
```

---

### 向状态方案中添加一个工作项状态流转

**`POST /v1/project/work_item_state_plans/{state_plan_id}/work_item_state_flows`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `state_plan_id` | String | 工作项状态方案的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `from_state_id` | String | 起始工作项状态的id。 |
| `to_state_id` | String | 可更改为的工作项状态的id。 |

**响应示例：**

```json
{
    "id": "5ef85b1e9481936604da7fcd",
    "url": "https://rest_api_root/v1/project/work_item_state_plans/5c9b62f18ad715335de90c20/work_item_state_flows/5ef85b1e9481936604da7fcd",
    "state_plan": {
        "id": "5c9b62f18ad715335de90c20",
        "url": "https://rest_api_root/v1/project/work_item_state_plans/5c9b62f18ad715335de90c20",
        "project_type": "scrum",
        "work_item_type": "story"
    },
    "from_state": {
        "id": "5c9b35de90ad7153c2062f18",
        "url": "https://rest_api_root/v1/project/work_item_states/5c9b35de90ad7153c2062f18",
        "name": "新建",
        "type": "pending",
        "color": "#56ABFB"
    },
    "to_state": {
        "id": "5ef85b1e9481936604da7f4c",
        "url": "https://rest_api_root/v1/project/work_item_states/5ef85b1e9481936604da7f4c",
        "name": "进行中",
        "type": "in_progress",
        "color": "#F6C659"
    }
}
```

---

### 在属性方案中移除一个工作项属性

**`DELETE /v1/project/work_item_property_plans/{property_plan_id}/work_item_properties/{property_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `property_plan_id` | String | 工作项属性方案的id。 |
| `property_id` | String | 工作项属性的id。 |

**响应示例：**

```json
{
    "id": "severity",
    "url": "https://rest_api_root/v1/project/work_item_property_plans/5f8a21f18ef715265de90c21/properties/severity",
    "property_plan": {
        "id": "5f8a21f18ef715265de90c21",
        "url": "https://rest_api_root/v1/project/work_item_property_plans/5f8a21f18ef715265de90c21",
        "project_type": "scrum",
        "work_item_type": "story"
    },
    "property": {
        "id": "severity",
        "url": "https://rest_api_root/v1/project/work_item_properties/severity",
        "name": "严重程度",
        "type": "select"
    }
}
```

---

### 在工作项类型方案中移除一个工作项类型

**`DELETE /v1/project/work_item_type_plans/{work_item_type_plan_id}/work_item_types/{work_item_type_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `work_item_type_plan_id` | String | 工作项类型方案的id。 |
| `work_item_type_id` | String | 工作项类型的id。 |

**响应示例：**

```json
{
    "id": "630da48bc9443b1aa94ce3ea",
    "url": "https://rest_api_root/v1/project/work_item_type_plans/65b0d9fd9ee456e672657e39/work_item_types/630da48bc9443b1aa94ce3ea",
    "type_plan": {
        "id": "6abb92f18ad725395df86c45",
        "url": "https://rest_api_root/v1/project/work_item_type_plans/65b0d9fd9ee456e672657e39",
        "project_type": "waterfall"
    },
    "type": {
        "id": "630da48bc9443b1aa94ce3ea",
        "url": "https://rest_api_root/v1/project/work_item_types/630da48bc9443b1aa94ce3ea",
        "name": "需求",
        "group": "requirement"
    },
    "sub_types": [
        {
           "id": "bug",
           "url": "https://rest_api_root/v1/project/work_item_types/bug",
           "name": "缺陷",
           "group": "bug"
        },
        {
           "id": "6385c650fef18f2d7222d15d",
           "url": "https://rest_api_root/v1/project/work_item_types/6385c650fef18f2d7222d15d",
           "name": "自定义",
           "group": "issue"
        }
    ]
}
```

---

### 在状态方案中移除一个工作项状态

**`DELETE /v1/project/work_item_state_plans/{state_plan_id}/work_item_states/{state_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `state_plan_id` | String | 工作项状态方案的id。 |
| `state_id` | String | 工作项状态的id。 |

**响应示例：**

```json
{
    "id": "5cc2062f189b35de90ad7153",
    "url": "https://rest_api_root/v1/project/work_item_state_plans/5c9b62f18ad715335de90c20/work_item_states/5c9b35de90ad7153c2062f18",
    "state_plan": {
        "id": "5c9b62f18ad715335de90c20",
        "url": "https://rest_api_root/v1/project/work_item_state_plans/5c9b62f18ad715335de90c20",
        "project_type": "scrum",
        "work_item_type": "story"
    },
    "state": {
        "id": "5c9b35de90ad7153c2062f18",
        "url": "https://rest_api_root/v1/project/work_item_states/5c9b35de90ad7153c2062f18",
        "name": "新建",
        "type": "pending",
        "color": "#ff7575"
    }
}
```

---

### 在状态方案中移除一个工作项状态流转

**`DELETE /v1/project/work_item_state_plans/{state_plan_id}/work_item_state_flows/{flow_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `state_plan_id` | String | 工作项状态方案的id。 |
| `flow_id` | String | 工作项状态流转的id。 |

**响应示例：**

```json
{
    "id": "5ef85b1e9481936604da7fcd",
    "url": "https://rest_api_root/v1/project/work_item_state_plans/5c9b62f18ad715335de90c20/work_item_state_flows/5ef85b1e9481936604da7fcd",
    "state_plan": {
        "id": "5c9b62f18ad715335de90c20",
        "url": "https://rest_api_root/v1/project/work_item_state_plans/5c9b62f18ad715335de90c20",
        "project_type": "scrum",
        "work_item_type": "story"
    },
    "from_state": {
        "id": "5c9b35de90ad7153c2062f18",
        "url": "https://rest_api_root/v1/project/work_item_states/5c9b35de90ad7153c2062f18",
        "name": "新建",
        "type": "pending",
        "color": "#56ABFB"
    },
    "to_state": {
        "id": "5ef85b1e9481936604da7f4c",
        "url": "https://rest_api_root/v1/project/work_item_states/5ef85b1e9481936604da7f4c",
        "name": "进行中",
        "type": "in_progress",
        "color": "#F6C659"
    }
}
```

---

### 开启工作项本地配置

**`POST /v1/project/work_item_plans`**

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `project_id` | String | 项目的id。 |

**响应示例：**

```json
{
    "id": "5eb623f6a70571487ea47000",
    "url": "https://rest_api_root/v1/project/work_item_plans/5eb623f6a70571487ea47000",
    "project": {
        "id": "5eb623f6a70571487ea47000",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
        "identifier": "SCR",
        "name": "Scrum项目",
        "type": "scrum",
        "is_archived": 0,
        "is_deleted": 0
    }
}
```

---

### 获取全部工作项属性列表

**`GET /v1/project/work_item_properties`**

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 2,
    "values": [
        {
            "id": "severity",
            "url": "https://rest_api_root/v1/project/work_item_properties/severity",
            "name": "严重程度",
            "type": "select",
            "options": [
                {
                    "_id": "5efb1859110533727a82c603",
                    "text": "严重"
                },
                {
                    "_id": "5efb1859110533727a82c604",
                    "text": "一般"
                }
            ],
            "is_removable": true,
            "is_name_editable": true,
            "is_options_editable": true,
            "select_all_level": false,
            "display_all_level": false,
            "display_separator": null
        },
        {
            "id": "identifier",
            "url": "https://rest_api_root/v1/project/work_item_properties/identifier",
            "name": "编号",
            "type": "text",
            "options": null,
            "is_removable": false,
            "is_name_editable": false,
            "is_options_editable": false,
            "select_all_level": false,
            "display_all_level": false,
            "display_separator": null
        }
    ]
}
```

---

### 获取全部工作项标签列表

**`GET /v1/project/work_item_tags`**

**查询参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `name` *(可选)* | String | 标签的名称。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "5e6b35de50ef8153c2062f70",
            "url": "https://rest_api_root/v1/project/work_item_tags/5e6b35de50ef8153c2062f70",
            "name": "标签-2",
            "color": "#56ABFB"
        }
    ]
}
```

---

### 获取全部工作项状态列表

**`GET /v1/project/work_item_states`**

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "5c9b35de90ad7153c2062f18",
            "url": "https://rest_api_root/v1/project/work_item_states/5c9b35de90ad7153c2062f18",
            "name": "新建",
            "type": "pending",
            "color": "#ff7575",
            "is_system": true
        }
    ]
}
```

---

### 获取全部工作项类型列表

**`GET /v1/project/work_item_types`**

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 10,
    "values": [
        {
            "id": "epic",
            "url": "http://rest_api_root/v1/project/work_item_types/epic",
            "name": "史诗",
            "group": "requirement",
            "is_system": true
        },
        {
            "id": "feature",
            "url": "http://rest_api_root/v1/project/work_item_types/feature",
            "name": "特性",
            "group": "requirement",
            "is_system": true
        },
        {
            "id": "story",
            "url": "http://rest_api_root/v1/project/work_item_types/story",
            "name": "用户故事",
            "group": "requirement",
            "is_system": true
        },
        {
            "id": "task",
            "url": "http://rest_api_root/v1/project/work_item_types/task",
            "name": "任务",
            "group": "task",
            "is_system": true
        },
        {
            "id": "bug",
            "url": "http://rest_api_root/v1/project/work_item_types/bug",
            "name": "缺陷",
            "group": "bug",
            "is_system": true
        },
        {
            "id": "issue",
            "url": "http://rest_api_root/v1/project/work_item_types/issue",
            "name": "事务",
            "group": "issue",
            "is_system": true
        },
        {
            "id": "630da48bc9443b1aa94ce3ea",
            "url": "https://rest_api_root/v1/project/work_item_types/630da48bc9443b1aa94ce3ea",
            "name": "需求",
            "group": "requirement",
            "is_system": true
        },
        {
            "id": "6385c650fef18f2d7222d15d",
            "url": "https://rest_api_root/v1/project/work_item_types/6385c650fef18f2d7222d15d",
            "name": "自定义",
            "group": "issue",
            "is_system": false
        },
        {
            "id": "630da48bc9443b1aa94ce3ee",
            "url": "https://rest_api_root/v1/project/work_item_types/630da48bc9443b1aa94ce3ee",
            "name": "阶段",
            "group": "plan",
            "is_system": true
        },
        {
            "id": "630da48bc9443b1aa94ce3ef",
            "url": "https://rest_api_root/v1/project/work_item_types/630da48bc9443b1aa94ce3ef",
            "name": "里程碑",
            "group": "plan",
            "is_system": true
        }
    ]
}
```

---

### 获取属性方案中的工作项属性列表

**`GET /v1/project/work_item_property_plans/{property_plan_id}/work_item_properties`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `property_plan_id` | String | 工作项属性方案的id。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 2,
    "values": [
        {
            "id": "severity",
            "url": "https://rest_api_root/v1/project/work_item_property_plans/5f8a21f18ef715265de90c21/properties/severity",
            "property_plan": {
                "id": "5f8a21f18ef715265de90c21",
                "url": "https://rest_api_root/v1/project/work_item_property_plans/5f8a21f18ef715265de90c21",
                "project_type": "scrum",
                "work_item_type": "story"
            },
            "property": {
                "id": "severity",
                "url": "https://rest_api_root/v1/project/work_item_properties/severity",
                "name": "严重程度",
                "type": "select"
            }
        },
        {
            "id": "identifier",
            "url": "https://rest_api_root/v1/project/work_item_property_plans/5f8a21f18ef715265de90c21/properties/identifier",
            "property_plan": {
                "id": "5f8a21f18ef715265de90c21",
                "url": "https://rest_api_root/v1/project/work_item_property_plans/5f8a21f18ef715265de90c21",
                "project_type": "scrum",
                "work_item_type": "story"
            },
            "property": {
                "id": "identifier",
                "url": "https://rest_api_root/v1/project/work_item_properties/identifier",
                "name": "编号",
                "type": "text"
            }
        }
    ]
}
```

---

### 获取工作项属性方案列表

**`GET /v1/project/work_item_property_plans`**

**查询参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `project_id` *(可选)* | String | 项目的id。查询开启本地配置的工作项属性方案时传入。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "5f8a21f18ef715265de90c21",
            "url": "https://rest_api_root/v1/project/work_item_property_plans/5f8a21f18ef715265de90c21",
            "project_type": "scrum",
            "process_id": "5fa690f1ae0571487ea49030",
            "work_item_type": "story",
            "project": {
                "id": "5eb623f6a70571487ea47000",
                "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
                "identifier": "SCR",
                "name": "Scrum项目",
                "type": "scrum",
                "is_archived": 0,
                "is_deleted": 0
            }
        }
    ]
}
```

---

### 获取工作项状态方案列表

**`GET /v1/project/work_item_state_plans`**

**查询参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `project_id` *(可选)* | String | 项目的id。查询开启本地配置的工作项状态方案时传入。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "5c9b62f18ad715335de90c20",
            "url": "https://rest_api_root/v1/project/work_item_state_plans/5c9b62f18ad715335de90c20",
            "project_type": "scrum",
            "process_id": "5fa690f1ae0571487ea49030",
            "work_item_type": "story",
            "project": {
                "id": "5eb623f6a70571487ea47000",
                "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
                "identifier": "SCR",
                "name": "Scrum项目",
                "type": "scrum",
                "is_archived": 0,
                "is_deleted": 0
            }
        }
    ]
}
```

---

### 获取工作项类型方案中的工作项类型列表

**`GET /v1/project/work_item_type_plans/{work_item_type_plan_id}/work_item_types`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `work_item_type_plan_id` | String | 工作项类型方案的id。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "630da48bc9443b1aa94ce3ea",
            "url": "https://rest_api_root/v1/project/work_item_type_plans/65b0d9fd9ee456e672657e39/work_item_types/630da48bc9443b1aa94ce3ea",
            "type_plan": {
                "id": "6abb92f18ad725395df86c45",
                "url": "https://rest_api_root/v1/project/work_item_type_plans/65b0d9fd9ee456e672657e39",
                "project_type": "waterfall"
            },
            "type": {
                "id": "630da48bc9443b1aa94ce3ea",
                "url": "https://rest_api_root/v1/project/work_item_types/630da48bc9443b1aa94ce3ea",
                "name": "需求",
                "group": "requirement"
            },
            "sub_types": [
                {
                   "id": "bug",
                   "url": "https://rest_api_root/v1/project/work_item_types/bug",
                   "name": "缺陷",
                   "group": "bug"
                },
                {
                   "id": "6385c650fef18f2d7222d15d",
                   "url": "https://rest_api_root/v1/project/work_item_types/6385c650fef18f2d7222d15d",
                   "name": "自定义",
                   "group": "issue"
                }
            ]
        }
    ]
}
```

---

### 获取工作项类型方案列表

**`GET /v1/project/work_item_type_plans`**

**查询参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `project_id` *(可选)* | String | 项目的id。查询开启本地配置的工作项类型方案时传入。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "5eb623f6a70571487ea47000",
            "url": "https://rest_api_root/v1/project/work_item_type_plans/5eb623f6a70571487ea47000",
            "project_type": "scrum",
            "process_id": "5fa690f1ae0571487ea49030",
            "project": {
                "id": "5eb623f6a70571487ea47000",
                "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
                "identifier": "SCR",
                "name": "Scrum项目",
                "type": "scrum",
                "is_archived": 0,
                "is_deleted": 0
            }
        }
    ]
}
```

---

### 获取状态方案中的工作项状态列表

**`GET /v1/project/work_item_state_plans/{state_plan_id}/work_item_states`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `state_plan_id` | String | 工作项状态方案的id。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "5cc2062f189b35de90ad7153",
            "url": "https://rest_api_root/v1/project/work_item_state_plans/5c9b62f18ad715335de90c20/work_item_states/5c9b35de90ad7153c2062f18",
            "state_plan": {
                "id": "5c9b62f18ad715335de90c20",
                "url": "https://rest_api_root/v1/project/work_item_state_plans/5c9b62f18ad715335de90c20",
                "project_type": "scrum",
                "work_item_type": "story"
            },
            "state": {
                "id": "5c9b35de90ad7153c2062f18",
                "url": "https://rest_api_root/v1/project/work_item_states/5c9b35de90ad7153c2062f18",
                "name": "新建",
                "type": "pending",
                "color": "#ff7575"
            }
        }
    ]
}
```

---

### 获取状态方案中的工作项状态流转列表

**`GET /v1/project/work_item_state_plans/{state_plan_id}/work_item_state_flows`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `state_plan_id` | String | 工作项状态方案的id。 |

**查询参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `from_state_id` *(可选)* | String | 上一个工作项状态的id。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "5ef85b1e9481936604da7fcd",
            "url": "https://rest_api_root/v1/project/work_item_state_plans/5c9b62f18ad715335de90c20/work_item_state_flows/5ef85b1e9481936604da7fcd",
            "state_plan": {
                "id": "5c9b62f18ad715335de90c20",
                "url": "https://rest_api_root/v1/project/work_item_state_plans/5c9b62f18ad715335de90c20",
                "project_type": "scrum",
                "work_item_type": "story"
            },
            "from_state": {
                "id": "5c9b35de90ad7153c2062f18",
                "url": "https://rest_api_root/v1/project/work_item_states/5c9b35de90ad7153c2062f18",
                "name": "新建",
                "type": "pending",
                "color": "#56ABFB"
            },
            "to_state": {
                "id": "5ef85b1e9481936604da7f4c",
                "url": "https://rest_api_root/v1/project/work_item_states/5ef85b1e9481936604da7f4c",
                "name": "进行中",
                "type": "in_progress",
                "color": "#F6C659"
            }
        }
    ]
}
```

---

### 部分更新一个工作项属性

**`PATCH /v1/project/work_item_properties/{property_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `property_id` | String | 工作项属性的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `name` *(可选)* | String | 工作项属性的名称。在一个企业中这个名称是唯一的。 |
| `options` *(可选)* | Object[] | 选择项列表。options是整体更新的。 |
| `options._id` *(可选)* | String | 选择项id。该值在选择项中是唯一的。 |
| `options.text` | String | 选择项内容。该值在选择项中是唯一的。 |
| `options.parent_id` *(可选)* | String | 选择项父级id。当属性类型为cascade_select,cascade_multi_select时，parent_id参数有效，用于构建级联类型选择项之间的父子关系，最多存在4级。 |

**响应示例（select）：**

```json
{
    "id": "severity-update",
    "url": "https://rest_api_root/v1/project/work_item_properties/severity",
    "name": "严重程度-update",
    "type": "select",
    "options": [
        {
            "_id": "5efb1859110533727a82c603",
            "text": "严重-update"
        },
        {
            "_id": "5efb1859110533727a82c624",
            "text": "一般"
        }
    ],
    "is_removable": true,
    "is_name_editable": true,
    "is_options_editable": true,
    "select_all_level": false,
    "display_all_level": false,
    "display_separator": null
}
```

**响应示例（cascade_select）：**

```json
{
    "id": "jiliandanxuan",
    "url": "https://rest_api_root/v1/project/work_item_properties/jiliandanxuan",
    "name": "级联单选-update",
    "type": "cascade_select",
    "options": [
        {
            "_id": "64b9f741eec7fd94e63b411e",
            "text": "父-update"
        },
        {
            "_id": "64b9f741eec7fd94e63b411f",
            "text": "子-update",
            "parent_id": "64b9f741eec7fd94e63b411e"
        },
        {
            "_id": "64b9f741eec7fd94e63b411g",
            "text": "子-create",
            "parent_id": "64b9f741eec7fd94e63b411f"
        }
    ],
    "is_removable": true,
    "is_name_editable": true,
    "is_options_editable": true,
    "select_all_level": false,
    "display_all_level": false,
    "display_separator": "/"
}
```

---

### 部分更新一个工作项标签

**`PATCH /v1/project/work_item_tags/{tag_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `tag_id` | String | 标签的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `name` *(可选)* | String | 标签的名称。在一个企业中这个名称是唯一的。 |

**响应示例：**

```json
{
    "id": "5e6b35de50ef8153c2062f70",
    "url": "https://rest_api_root/v1/project/work_item_tags/5e6b35de50ef8153c2062f70",
    "name": "标签-2",
    "color": "#56ABFB"
}
```

---

### 部分更新一个工作项状态

**`PATCH /v1/project/work_item_states/{state_id}`**

只有非系统类型的工作项状态才能更新。

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `name` *(可选)* | String | 工作项状态的名称。在一个企业中这个名称是唯一的。 |
| `type` *(可选)* | String | 工作项状态的类型。 |

**响应示例：**

```json
{
    "id": "5c9b35de90ad7153c2062f18",
    "url": "https://rest_api_root/v1/project/work_item_states/5c9b35de90ad7153c2062f18",
    "name": "新建",
    "type": "pending",
    "color": "#ff7575",
    "is_system": false
}
```

---

### 部分更新一个工作项类型

**`PATCH /v1/project/work_item_types/{work_item_type_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `work_item_type_id` | String | 工作项类型的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `name` | String | 工作项类型的名称。在一个企业中这个名称是唯一的。 |

**响应示例：**

```json
{
    "id": "630da48bc9443b1aa94ce3df",
    "url": "https://rest_api_root/v1/project/work_item_types/630da48bc9443b1aa94ce3df",
    "name": "非功能性需求",
    "group": "requirement",
    "is_system": false
}
```

---

### 部分更新工作项类型方案中的工作项类型

**`PATCH /v1/project/work_item_type_plans/{work_item_type_plan_id}/work_item_types/{work_item_type_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `work_item_type_plan_id` | String | 工作项类型方案的id。 |
| `work_item_type_id` | String | 工作项类型的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `sub_type_ids` | String[] | 子工作项类型id的列表，使用','分割，最多只能20个。 |

**响应示例：**

```json
{
    "id": "630da48bc9443b1aa94ce3ea",
    "url": "https://rest_api_root/v1/project/work_item_type_plans/65b0d9fd9ee456e672657e39/work_item_types/630da48bc9443b1aa94ce3ea",
    "type_plan": {
        "id": "6abb92f18ad725395df86c45",
        "url": "https://rest_api_root/v1/project/work_item_type_plans/65b0d9fd9ee456e672657e39",
        "project_type": "waterfall"
    },
    "type": {
        "id": "630da48bc9443b1aa94ce3ea",
        "url": "https://rest_api_root/v1/project/work_item_types/630da48bc9443b1aa94ce3ea",
        "name": "需求",
        "group": "requirement"
    },
    "sub_types": [
        {
           "id": "bug",
           "url": "https://rest_api_root/v1/project/work_item_types/bug",
           "name": "缺陷",
           "group": "bug"
        },
        {
           "id": "6385c650fef18f2d7222d15d",
           "url": "https://rest_api_root/v1/project/work_item_types/6385c650fef18f2d7222d15d",
           "name": "自定义",
           "group": "issue"
        }
    ]
}
```

---

## 工单

### 创建一个工单

**`POST /v1/ship/tickets`**

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `product_id` | String | 工单的产品id。 |
| `title` | String | 工单的标题，最大长度为255。 |
| `type_id` | String | 工单的类型id，您可以在 “获取工单类型列表” API获取。 |
| `description` *(可选)* | String | 工单的描述。 |
| `submitter_id` *(可选)* | String | 工单的提交人id，企业授权时，该值有效；个人鉴权时，指定无效。 |
| `customer_id` *(可选)* | String | 工单的客户id，您可以在 “获取产品客户列表” API获取。 |
| `channel_id` *(可选)* | String | 工单的渠道id，您可以在 “获取渠道列表” API获取。 |
| `assignee_id` *(可选)* | String | 工单的负责人id。 |
| `priority_id` *(可选)* | String | 工单的优先级id，您可以在 “获取工单优先级列表” API获取。 |
| `properties` *(可选)* | Object | 工单的自定义属性。 |
| `properties.prop_a` *(可选)* | Object | 工单的自定义属性prop_a。 |
| `properties.prop_b` *(可选)* | Object | 工单的自定义属性prop_b。 |

**响应示例：**

```json
{
    "id": "63eca888a0a13a3efc8d4a43",
    "url": "https://rest_api_root/v1/ship/tickets/63eca888a0a13a3efc8d4a43",
    "product": {
        "id": "6422711c3f12e6c1e46d40e9",
        "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9",
        "identifier": "SLC",
        "name": "示例产品",
        "is_archived": 0,
        "is_deleted": 0
    },
    "identifier": "SLC-T1",
    "title": "希望新增支持第三方账号注册",
    "short_id": "pdMUgQzZ",
    "html_url": "https://yctech.pingcode.com/ship/tickets/pdMUgQzZ",
    "assignee": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "state": {
        "id": "63eca880a0a13a3efc8d49d9",
        "url": "https://rest_api_root/v1/ship/ticket_states/63eca880a0a13a3efc8d49d9",
        "name": "待处理",
        "type": "pending"
    },
    "type": {
        "id": "63eca880a0a13a3efc8d49e0",
        "url": "https://rest_api_root/v1/ship/ticket_types/63eca880a0a13a3efc8d49e0",
        "name": "需求"
    },
    "customer": {
        "id": "63eca881a0a13a3efc8d49fc",
        "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9/customers/63eca881a0a13a3efc8d49fc",
        "name": "北京XX科技有限公司"
    },
    "solution": {
        "id": "62f217ae16e3661a20124330",
        "url": "https://rest_api_root/v1/ship/ticket_solutions/62f217ae16e3661a20124330",
        "name": "进入需求池"
    },
    "priority": {
        "url": "https://rest_api_root/v1/ship/ticket_priorities/5cb9466afda1ce4ca0090004",
        "id": "5cb9466afda1ce4ca0090004",
        "name": "P1"
    },
    "channel": {
        "id": "64550d9ec696b249b5fc607d",
        "url": "https://rest_api_root/v1/ship/channels/64550d9ec696b249b5fc607d",
        "name": "channel-1"
    },
    "description": "<p>希望支持其他更多第三方平台的账号注册，以便用第三方账号登录找回更换了手机号的账号，保障账号安全</p>",
    "properties": {
        "prop_a": "prop_a_value",
        "prop_b": "prop_b_value"
    },
    "estimated_at": null,
    "tags": [],
    "participants": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "url": "https://rest_api_root/v1/participants/a0417f68e846aae315c85d24643678a9?principal_type=ticket&principal_id=63eca888a0a13a3efc8d4a43",
            "type": "user",
            "user": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            }
        },
        {
            "id": "63c8fb32729dee3334d96af7",
            "url": "https://rest_api_root/v1/participants/63c8fb32729dee3334d96af7?principal_type=ticket&principal_id=63eca888a0a13a3efc8d4a43",
            "type": "user_group",
            "user_group": {
                "id": "63c8fb32729dee3334d96af7",
                "url": "https://rest_api_root/v1/directory/groups/63c8fb32729dee3334d96af7",
                "name": "Open Team"
            }
        }
    ],
    "submitted_at": 1676454024,
    "submitted_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "completed_at": 1689579131,
    "completed_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "created_at": 1676454024,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "updated_at": 1676454024,
    "updated_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "is_archived": 0,
    "is_deleted": 0
}
```

---

### 获取工单优先级列表

**`GET /v1/ship/ticket/priorities?product_id={product_id}`**

**查询参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `product_id` | String | 产品的id。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "5cb9466afda1ce4ca0090005",
            "url": "https://rest_api_root/v1/ship/ticket_priorities/5cb9466afda1ce4ca0090005",
            "name": "P0"
        }
    ]
}
```

---

### 获取工单列表

**`GET /v1/ship/tickets`**

**查询参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `product_id` *(可选)* | String | 产品的id。 |
| `type_id` *(可选)* | String | 工单类型id。 |
| `state_id` *(可选)* | String | 工单状态id。 |
| `priority_id` *(可选)* | String | 工单优先级id。 |
| `created_between` *(可选)* | String | 创建时间介于的时间范围，通过','分割起始时间。 |
| `updated_between` *(可选)* | String | 更新时间介于的时间范围，通过','分割起始时间。 |
| `keywords` *(可选)* | String | 关键字。支持工单编号和工单标题。 |
| `include_public_image_token` *(可选)* | String | 包含获取图片资源token的属性。使用','分割，最多只能20个，支持description和自定义多行文本类型的属性。参数示例description,properties.prop_b。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "63eca888a0a13a3efc8d4a43",
            "url": "https://rest_api_root/v1/ship/tickets/63eca888a0a13a3efc8d4a43",
            "product": {
                "id": "6422711c3f12e6c1e46d40e9",
                "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9",
                "identifier": "SLC",
                "name": "示例产品",
                "is_archived": 0,
                "is_deleted": 0
            },
            "identifier": "SLC-T1",
            "title": "希望新增支持第三方账号注册",
            "short_id": "pdMUgQzZ",
            "html_url": "https://yctech.pingcode.com/ship/tickets/pdMUgQzZ",
            "assignee": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "state": {
                "id": "63eca880a0a13a3efc8d49d9",
                "url": "https://rest_api_root/v1/ship/ticket_states/63eca880a0a13a3efc8d49d9",
                "name": "待处理",
                "type": "pending"
            },
            "estimated_at": {
                "from": 1701619200,
                "to": 1702742399,
                "granularity": "day"
            },
            "type": {
                "id": "63eca880a0a13a3efc8d49e0",
                "url": "https://rest_api_root/v1/ship/ticket_types/63eca880a0a13a3efc8d49e0",
                "name": "需求"
            },
            "customer": {
                "id": "63eca881a0a13a3efc8d49fc",
                "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9/customers/63eca881a0a13a3efc8d49fc",
                "name": "北京XX科技有限公司"
            },
            "solution": {
                "id": "62f217ae16e3661a20124330",
                "url": "https://rest_api_root/v1/ship/ticket_solutions/62f217ae16e3661a20124330",
                "name": "进入需求池"
            },
            "priority": {
                "url": "https://rest_api_root/v1/ship/ticket_priorities/5cb9466afda1ce4ca0090004",
                "id": "5cb9466afda1ce4ca0090004",
                "name": "P1"
            },
            "channel": {
                "id": "64550d9ec696b249b5fc607d",
                "url": "https://rest_api_root/v1/ship/channels/64550d9ec696b249b5fc607d",
                "name": "channel-1"
            },
            "description": "<p>希望支持其他更多第三方平台的账号注册，以便用第三方账号登录找回更换了手机号的账号，保障账号安全</p>",
            "properties": {
                "prop_a": "prop_a_value",
                "prop_b": "prop_b_value"
            },
            "tags": [
                {
                    "id": "63eca881a0a13a3efc8d49f1",
                    "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9/tags/63eca881a0a13a3efc8d49f1",
                    "name": "已确认"
                }
            ],
            "participants": [
                {
                    "id": "a0417f68e846aae315c85d24643678a9",
                    "url": "https://rest_api_root/v1/participants/a0417f68e846aae315c85d24643678a9?principal_type=ticket&principal_id=63eca888a0a13a3efc8d4a43",
                    "type": "user",
                    "user": {
                        "id": "a0417f68e846aae315c85d24643678a9",
                        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                        "name": "john",
                        "display_name": "John",
                        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
                    }
                },
                {
                    "id": "63c8fb32729dee3334d96af7",
                    "url": "https://rest_api_root/v1/participants/63c8fb32729dee3334d96af7?principal_type=ticket&principal_id=63eca888a0a13a3efc8d4a43",
                    "type": "user_group",
                    "user_group": {
                        "id": "63c8fb32729dee3334d96af7",
                        "url": "https://rest_api_root/v1/directory/groups/63c8fb32729dee3334d96af7",
                        "name": "Open Team"
                    }
                }
            ],
            "public_image_token": "N96GlJ4AMQoBCw9pZQ2EMl-AprLN_V_DYlghupBNkJA",
            "submitted_at": 1676454024,
            "submitted_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "completed_at": 1689579131,
            "completed_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "created_at": 1676454024,
            "created_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "updated_at": 1676454024,
            "updated_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "is_archived": 0,
            "is_deleted": 0
        }
    ]
}
```

---

### 获取工单属性列表

**`GET /v1/ship/ticket/properties?product_id={product_id}`**

**查询参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `product_id` | String | 产品的id。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 2,
    "values": [
        {
            "id": "solution",
            "url": "https://rest_api_root/v1/ship/ticket_properties/solution",
            "name": "解决方案",
            "type": "select",
            "options": [
                {
                    "_id": "6422711c3f12e6c1e46d40e9",
                    "text": "进入需求池"
                }
            ]
        },
        {
            "id": "identifier",
            "url": "https://rest_api_root/v1/ship/ticket_properties/identifier",
            "name": "编号",
            "type": "text",
            "options": null
        }
    ]
}
```

---

### 获取工单标签列表

**`GET /v1/ship/ticket/tags?product_id={product_id}`**

**查询参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `product_id` | String | 产品的id。 |

**响应示例：**

```json
{
     "page_size": 30,
     "page_index": 0,
     "total": 1,
     "values": [
         {
             "id": "63eca881a0a13a3efc8d49f0",
             "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e6/tags/63eca881a0a13a3efc8d49f0",
             "name": "技术支持确认"
         }
     ]
 }
```

---

### 获取工单流转记录列表

**`GET /v1/ship/tickets/{ticket_id}/transition_histories`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `ticket_id` | String | 工单的id。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 2,
    "values": [
        {
            "id": "64c3676c983bb9481ee1eea5",
            "url": "https://rest_api_root/v1/ship/tickets/5edca524cad2fa1125cb0630/transition_histories/64c3676c983bb9481ee1eea5",
            "ticket": {
                "id": "63eca888a0a13a3efc8d4a43",
                "url": "https://rest_api_root/v1/ship/tickets/63eca888a0a13a3efc8d4a43",
                "identifier": "SLC-T1",
                "title": "希望新增支持第三方账号注册",
                "short_id": "pdMUgQzZ",
                "html_url": "https://yctech.pingcode.com/ship/tickets/pdMUgQzZ"
            },
            "from_state": null,
            "to_state": {
                "id": "63e1bf51898a0be5a2d21b29",
                "url": "https://rest_api_root/v1/ship/ticket_states/63e1bf51898a0be5a2d21b29",
                "name": "待处理",
                "type": "pending"
            },
            "created_at": 1674528614,
            "created_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            }
        },
        {
            "id": "658bdb79e5839f556562accf",
            "url": "https://rest_api_root/v1/ship/tickets/5edca524cad2fa1125cb0630/transition_histories/658bdb79e5839f556562accf",
            "ticket": {
                "id": "63eca888a0a13a3efc8d4a43",
                "url": "https://rest_api_root/v1/ship/tickets/63eca888a0a13a3efc8d4a43",
                "identifier": "SLC-T1",
                "title": "希望新增支持第三方账号注册",
                "short_id": "pdMUgQzZ",
                "html_url": "https://yctech.pingcode.com/ship/tickets/pdMUgQzZ"
            },
            "from_state": {
                "id": "63e1bf51898a0be5a2d21b29",
                "url": "https://rest_api_root/v1/ship/ticket_states/63e1bf51898a0be5a2d21b29",
                "name": "待处理",
                "type": "pending"
            },
            "to_state": {
                "id": "63e1bf51898a0be5a2d21b2b",
                "url": "https://rest_api_root/v1/ship/ticket_states/63e1bf51898a0be5a2d21b2b",
                "name": "处理中",
                "type": "in_progress"
            },
            "created_at": 1674528614,
            "created_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            }
        }
    ]
}
```

---

### 获取工单渠道列表

**`GET /v1/ship/ticket/channels?product_id={product_id}`**

**查询参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `product_id` | String | 产品的id。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "63eca881a0a13a3efc8d49ed",
            "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e6/channels/63eca881a0a13a3efc8d49ed",
            "name": "客户反馈Web渠道"
        }
    ]
}
```

---

### 获取工单状态列表

**`GET /v1/ship/ticket/states?product_id={product_id}`**

**查询参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `product_id` | String | 产品的id。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "6422711c3f12e6c1e46d40f2",
            "url": "https://rest_api_root/v1/ship/ticket_states/6422711c3f12e6c1e46d40f2",
            "name": "处理中",
            "type": "pending"
        }
    ]
}
```

---

### 获取工单类型列表

**`GET /v1/ship/ticket/types?product_id={product_id}`**

**查询参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `product_id` | String | 产品的id。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "63bb744214bd13c9def24ca9",
            "url": "https://rest_api_root/v1/ship/ticket_types/63bb744214bd13c9def24ca9",
            "name": "需求"
        }
    ]
}
```

---

### 获取工单解决方案列表

**`GET /v1/ship/ticket/solutions?product_id={product_id}`**

**查询参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `product_id` | String | 产品的id。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "6422711c3f12e6c1e46d40e9",
            "url": "https://rest_api_root/v1/ship/ticket_solutions/6422711c3f12e6c1e46d40e9",
            "name": "进入需求池"
        }
    ]
}
```

---

### 部分更新一个工单

**`PATCH /v1/ship/tickets/{ticket_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `ticket_id` | String | 工单id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `title` *(可选)* | String | 工单的标题，最大长度为255。 |
| `description` *(可选)* | String | 工单的描述。 |
| `type_id` *(可选)* | String | 工单的类型id，您可以在 “获取工单类型列表” API获取。 |
| `state_id` *(可选)* | String | 工单的状态id，您可以在 “获取工单状态列表” API获取。 |
| `assignee_id` *(可选)* | String | 工单的负责人id。 |
| `submitter_id` *(可选)* | String | 工单的提交人id，企业授权时，该值有效；个人鉴权时，指定无效。 |
| `solution_id` *(可选)* | String | 工单的解决方案id，您可以在 “获取工单解决方案列表” API获取。 |
| `priority_id` *(可选)* | String | 工单的优先级id，您可以在 “获取工单优先级列表” API获取。 |
| `customer_id` *(可选)* | String | 工单的客户id，您可以在 “获取产品客户列表” API获取。 |
| `properties` *(可选)* | Object | 工单的自定义属性。 |
| `properties.prop_a` *(可选)* | Object | 工单的自定义属性prop_a。 |
| `properties.prop_b` *(可选)* | Object | 工单的自定义属性prop_b。 |

**响应示例：**

```json
{
    "id": "63eca888a0a13a3efc8d4a43",
    "url": "https://rest_api_root/v1/ship/tickets/63eca888a0a13a3efc8d4a43",
    "product": {
        "id": "6422711c3f12e6c1e46d40e9",
        "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9",
        "identifier": "SLC",
        "name": "示例产品",
        "is_archived": 0,
        "is_deleted": 0
    },
    "identifier": "SLC-T1",
    "title": "希望新增支持第三方账号注册",
    "short_id": "pdMUgQzZ",
    "html_url": "https://yctech.pingcode.com/ship/tickets/pdMUgQzZ",
    "assignee": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "state": {
        "id": "63eca880a0a13a3efc8d49d9",
        "url": "https://rest_api_root/v1/ship/ticket_states/63eca880a0a13a3efc8d49d9",
        "name": "待处理",
        "type": "pending"
    },
    "estimated_at": {
        "from": 1701619200,
        "to": 1702742399,
        "granularity": "day"
    },
    "type": {
        "id": "63eca880a0a13a3efc8d49e0",
        "url": "https://rest_api_root/v1/ship/ticket_types/63eca880a0a13a3efc8d49e0",
        "name": "需求"
    },
    "customer": {
        "id": "63eca881a0a13a3efc8d49fc",
        "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9/customers/63eca881a0a13a3efc8d49fc",
        "name": "北京XX科技有限公司"
    },
    "solution": {
        "id": "62f217ae16e3661a20124330",
        "url": "https://rest_api_root/v1/ship/ticket_solutions/62f217ae16e3661a20124330",
        "name": "进入需求池"
    },
    "priority": {
        "id": "5cb9466afda1ce4ca0090004",
        "url": "https://rest_api_root/v1/ship/ticket_priorities/5cb9466afda1ce4ca0090004",
        "name": "P1"
    },
    "channel": {
        "id": "64550d9ec696b249b5fc607d",
        "url": "https://rest_api_root/v1/ship/channels/64550d9ec696b249b5fc607d",
        "name": "channel-1"
    },
    "description": "<p>希望支持其他更多第三方平台的账号注册，以便用第三方账号登录找回更换了手机号的账号，保障账号安全</p>",
    "properties": {
        "prop_a": "prop_a_value",
        "prop_b": "prop_b_value"
    },
    "tags": [],
    "participants": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "url": "https://rest_api_root/v1/participants/a0417f68e846aae315c85d24643678a9?principal_type=ticket&principal_id=63eca888a0a13a3efc8d4a43",
            "type": "user",
            "user": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            }
        },
        {
            "id": "63c8fb32729dee3334d96af7",
            "url": "https://rest_api_root/v1/participants/63c8fb32729dee3334d96af7?principal_type=ticket&principal_id=63eca888a0a13a3efc8d4a43",
            "type": "user_group",
            "user_group": {
                "id": "63c8fb32729dee3334d96af7",
                "url": "https://rest_api_root/v1/directory/groups/63c8fb32729dee3334d96af7",
                "name": "Open Team"
            }
        }
    ],
    "submitted_at": 1676454024,
    "submitted_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "completed_at": 1689579131,
    "completed_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "created_at": 1676454024,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "updated_at": 1676455024,
    "updated_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "is_archived": 0,
    "is_deleted": 0
}
```

---

## 工单配置

### 创建一个工单属性

**`POST /v1/ship/ticket_properties`**

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `name` | String | 工单属性的名称。在一个企业中这个名称是唯一的。 |
| `type` | String | 工单属性的类型。 |
| `options` *(可选)* | Object[] | 选择项列表。当工单属性类型为select,multi_select,cascade_select,cascade_multi_select时可选填此项。 |
| `options._id` *(可选)* | String | 选择项id。该值在选择项中是唯一的。 |
| `options.text` | String | 选择项内容。该值在选择项中是唯一的。 |
| `options.parent_id` *(可选)* | String | 选择项父级id。当属性类型为cascade_select,cascade_multi_select时，parent_id参数有效，用于构建级联类型选择项之间的父子关系，最多存在4级。 |

**响应示例（select）：**

```json
{
    "id": "severity",
    "url": "https://rest_api_root/v1/ship/ticket_properties/severity",
    "name": "严重程度",
    "type": "select",
    "options": [
        {
            "_id": "5efb1859110533727a82c603",
            "text": "严重"
        },
        {
            "_id": "5efb1859110533727a82c604",
            "text": "一般"
        }
    ],
    "is_removable": true,
    "is_name_editable": true,
    "is_options_editable": true,
    "select_all_level": false,
    "display_all_level": false,
    "display_separator": null
}
```

**响应示例（cascade_select）：**

```json
{
    "id": "jiliandanxuan",
    "url": "https://rest_api_root/v1/ship/ticket_properties/jiliandanxuan",
    "name": "级联单选",
    "type": "cascade_select",
    "options": [
        {
            "_id": "64b9f741eec7fd94e63b411e",
            "text": "父"
        },
        {
            "_id": "64b9f741eec7fd94e63b411f",
            "text": "子",
            "parent_id": "64b9f741eec7fd94e63b411e"
        }
    ],
    "is_removable": true,
    "is_name_editable": true,
    "is_options_editable": true,
    "select_all_level": false,
    "display_all_level": false,
    "display_separator": "/"
}
```

---

### 创建一个工单状态

**`POST /v1/ship/ticket_states`**

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `name` | String | 工单状态的名称，在一个企业中这个名称是唯一的。 |
| `type` | String | 工单状态的类型。 |

**响应示例：**

```json
{
    "id": "6422711c3f12e6c1e46d40f2",
    "url": "https://rest_api_root/v1/ship/ticket_states/6422711c3f12e6c1e46d40f2",
    "name": "处理中",
    "type": "pending",
    "color": "#56ABFB"
}
```

---

### 向工单属性方案中添加一个工单属性

**`POST /v1/ship/ticket_property_plans/{property_plan_id}/ticket_properties`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `property_plan_id` | String | 工单属性方案的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `property_id` | String | 工单属性的id。 |

**响应示例：**

```json
{
    "id": "solution",
    "url": "https://rest_api_root/v1/ship/ticket_property_plans/5f8a21f18ef715265de90c21/ticket_properties/solution",
    "property_plan": {
        "id": "5f8a21f18ef715265de90c21",
        "url": "https://rest_api_root/v1/ship/ticket_property_plans/5f8a21f18ef715265de90c21"
    },
    "property": {
        "id": "solution",
        "url": "https://rest_api_root/v1/ship/ticket_properties/solution",
        "name": "解决方案",
        "type": "select"
    }
}
```

---

### 向状态方案中添加一个工单状态

**`POST /v1/ship/ticket_state_plans/{state_plan_id}/ticket_states`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `state_plan_id` | String | 工单状态方案的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `state_id` | String | 工单状态的id。 |

**响应示例：**

```json
{
    "id": "63bb744214bd13c9def24ca2",
    "url": "https://rest_api_root/v1/ship/ticket_state_plans/63feb3da9cc1ead1d2be93f4/ticket_states/63bb744214bd13c9def24ca2",
    "state_plan": {
        "id": "63feb3da9cc1ead1d2be93f4",
        "url": "https://rest_api_root/v1/ship/ticket_state_plans/63feb3da9cc1ead1d2be93f4"
    },
    "state": {
        "id": "63bb744214bd13c9def24ca2",
        "url": "https://rest_api_root/v1/ship/ticket_states/63bb744214bd13c9def24ca2",
        "name": "待处理",
        "type": "pending"
    }
}
```

---

### 向状态方案中添加一个工单状态流转

**`POST /v1/ship/ticket_state_plans/{state_plan_id}/ticket_state_flows`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `state_plan_id` | String | 工单状态方案的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `from_state_id` | String | 起始工单状态的id。 |
| `to_state_id` | String | 可更改为的工单状态的id。 |

**响应示例：**

```json
{
    "id": "63feb3da9cc1ead1d2be93fd",
    "url": "https://rest_api_root/v1/ship/ticket_state_plans/63feb3da9cc1ead1d2be93f4/ticket_state_flows/63feb3da9cc1ead1d2be93fd",
    "state_plan": {
        "id": "63feb3da9cc1ead1d2be93f4",
        "url": "https://rest_api_root/v1/ship/ticket_state_plans/63feb3da9cc1ead1d2be93f4"
    },
    "form_state": {
        "id": "63bb744214bd13c9def24ca5",
        "url": "https://rest_api_root/v1/ship/ticket_states/63bb744214bd13c9def24ca5",
        "name": "已计划",
        "type": "in_progress"
    },
    "to_state": {
        "id": "63bb744214bd13c9def24ca2",
        "url": "https://rest_api_root/v1/ship/ticket_states/63bb744214bd13c9def24ca2",
        "name": "待处理",
        "type": "pending"
    }
}
```

---

### 在工单属性方案中移除一个工单属性

**`DELETE /v1/ship/ticket_property_plans/{property_plan_id}/ticket_properties/{property_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `property_plan_id` | String | 工单属性方案的id。 |
| `property_id` | String | 工单属性的id。 |

**响应示例：**

```json
{
    "id": "solution",
    "url": "https://rest_api_root/v1/ship/ticket_property_plans/5f8a21f18ef715265de90c21/ticket_properties/solution",
    "property_plan": {
        "id": "5f8a21f18ef715265de90c21",
        "url": "https://rest_api_root/v1/ship/ticket_property_plans/5f8a21f18ef715265de90c21"
    },
    "property": {
        "id": "solution",
        "url": "https://rest_api_root/v1/ship/ticket_properties/solution",
        "name": "解决方案",
        "type": "select"
    }
}
```

---

### 在状态方案中移除一个工单状态

**`DELETE /v1/ship/ticket_state_plans/{state_plan_id}/ticket_states/{state_id}`**

移除状态后，每种类型的状态至少存在一种，否则将无法移除。

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `state_plan_id` | String | 工单状态方案的id。 |
| `state_id` | String | 工单状态的id。 |

**响应示例：**

```json
{
    "id": "63bb744214bd13c9def24ca2",
    "url": "https://rest_api_root/v1/ship/ticket_state_plans/63feb3da9cc1ead1d2be93f4/ticket_states/63bb744214bd13c9def24ca2",
    "state_plan": {
        "id": "63feb3da9cc1ead1d2be93f4",
        "url": "https://rest_api_root/v1/ship/ticket_state_plans/63feb3da9cc1ead1d2be93f4"
    },
    "state": {
        "id": "63bb744214bd13c9def24ca2",
        "url": "https://rest_api_root/v1/ship/ticket_states/63bb744214bd13c9def24ca2",
        "name": "待处理",
        "type": "pending"
    }
}
```

---

### 在状态方案中移除一个工单状态流转

**`DELETE /v1/ship/ticket_state_plans/{state_plan_id}/ticket_state_flows/{state_flow_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `state_plan_id` | String | 工单状态方案的id。 |
| `state_flow_id` | String | 工单状态流转的id。 |

**响应示例：**

```json
{
    "id": "63feb3da9cc1ead1d2be93fd",
    "url": "https://rest_api_root/v1/ship/ticket_state_plans/63feb3da9cc1ead1d2be93f4/ticket_state_flows/63feb3da9cc1ead1d2be93fd",
    "state_plan": {
        "id": "63feb3da9cc1ead1d2be93f4",
        "url": "https://rest_api_root/v1/ship/ticket_state_plans/63feb3da9cc1ead1d2be93f4"
    },
    "form_state": {
        "id": "63bb744214bd13c9def24ca5",
        "url": "https://rest_api_root/v1/ship/ticket_states/63bb744214bd13c9def24ca5",
        "name": "已计划",
        "type": "in_progress"
    },
    "to_state": {
        "id": "63bb744214bd13c9def24ca2",
        "url": "https://rest_api_root/v1/ship/ticket_states/63bb744214bd13c9def24ca2",
        "name": "待处理",
        "type": "pending"
    }
}
```

---

### 获取全部工单属性列表

**`GET /v1/ship/ticket_properties`**

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 2,
    "values": [
        {
            "id": "solution",
            "url": "https://rest_api_root/v1/ship/ticket_properties/solution",
            "name": "解决方案",
            "type": "select",
            "options": [
                {
                    "_id": "6422711c3f12e6c1e46d40e9",
                    "text": "进入需求池"
                }
            ],
            "is_removable": false,
            "is_name_editable": false,
            "is_options_editable": false
        },
        {
            "id": "identifier",
            "url": "https://rest_api_root/v1/ship/ticket_properties/identifier",
            "name": "编号",
            "type": "text",
            "options": null,
            "is_removable": false,
            "is_name_editable": false,
            "is_options_editable": false
        }
    ]
}
```

---

### 获取全部工单状态列表

**`GET /v1/ship/ticket_states`**

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "6422711c3f12e6c1e46d40f2",
            "url": "https://rest_api_root/v1/ship/ticket_states/6422711c3f12e6c1e46d40f2",
            "name": "处理中",
            "type": "pending",
            "color": "#F6C659"
        }
    ]
}
```

---

### 获取全部工单类型列表

**`GET /v1/ship/ticket_types`**

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "63bb744214bd13c9def24ca9",
            "url": "https://rest_api_root/v1/ship/ticket_types/63bb744214bd13c9def24ca9",
            "name": "需求",
            "is_system": true
        }
    ]
}
```

---

### 获取工单属性方案中的工单属性列表

**`GET /v1/ship/ticket_property_plans/{property_plan_id}/ticket_properties`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `property_plan_id` | String | 工单属性方案的id。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "solution",
            "url": "https://rest_api_root/v1/ship/ticket_property_plans/5f8a21f18ef715265de90c21/ticket_properties/solution",
            "property_plan": {
                "id": "5f8a21f18ef715265de90c21",
                "url": "https://rest_api_root/v1/ship/ticket_property_plans/5f8a21f18ef715265de90c21"
            },
            "property": {
                "id": "solution",
                "url": "https://rest_api_root/v1/ship/ticket_properties/solution",
                "name": "解决方案",
                "type": "select"
            }
        }
    ]
}
```

---

### 获取工单属性方案列表

**`GET /v1/ship/ticket_property_plans`**

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 2,
    "values": [
        {
            "id": "5f8a21f18ef715265de90c21",
            "url": "https://rest_api_root/v1/ship/ticket_property_plans/5f8a21f18ef715265de90c21",
            "product": null
        },
        {
            "id": "5f8a21f18ef715265de90c22",
            "url": "https://rest_api_root/v1/ship/ticket_property_plans/5f8a21f18ef715265de90c22",
            "product": {
                 "id": "6422711c3f12e6c1e46d40e9",
                 "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9",
                 "identifier": "SLC",
                 "name": "示例产品",
                 "is_archived": 0,
                 "is_deleted": 0
            }
        }
    ]
}
```

---

### 获取工单状态方案列表

**`GET /v1/ship/ticket_state_plans`**

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "63feb3da9cc1ead1d2be93f4",
            "url": "https://rest_api_root/v1/ship/ticket_state_plans/63feb3da9cc1ead1d2be93f4",
            "product": null
        },
        {
            "id": "63feb3da9cc1ead1d2be93f5",
            "url": "https://rest_api_root/v1/ship/ticket_state_plans/63feb3da9cc1ead1d2be93f5",
            "product": {
                 "id": "6422711c3f12e6c1e46d40e9",
                 "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9",
                 "identifier": "SLC",
                 "name": "示例产品",
                 "is_archived": 0,
                 "is_deleted": 0
            }
        }
    ]
}
```

---

### 获取状态方案中的工单状态列表

**`GET /v1/ship/ticket_state_plans/{state_plan_id}/ticket_states`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `state_plan_id` | String | 工单状态方案的id。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "63bb744214bd13c9def24ca2",
            "url": "https://rest_api_root/v1/ship/ticket_state_plans/63feb3da9cc1ead1d2be93f4/ticket_states/63bb744214bd13c9def24ca2",
            "state_plan": {
                "id": "63feb3da9cc1ead1d2be93f4",
                "url": "https://rest_api_root/v1/ship/ticket_state_plans/63feb3da9cc1ead1d2be93f4"
            },
            "state": {
                "id": "63bb744214bd13c9def24ca2",
                "url": "https://rest_api_root/v1/ship/ticket_states/63bb744214bd13c9def24ca2",
                "name": "待处理",
                "type": "pending"
            }
        }
    ]
}
```

---

### 获取状态方案中的工单状态流转列表

**`GET /v1/ship/ticket_state_plans/{state_plan_id}/ticket_state_flows`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `state_plan_id` | String | 工单状态方案的id。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "63feb3da9cc1ead1d2be93f5",
            "url": "https://rest_api_root/v1/ship/ticket_state_plans/63feb3da9cc1ead1d2be93f4/ticket_states/63feb3da9cc1ead1d2be93f5",
            "state_plan": {
                "id": "63feb3da9cc1ead1d2be93f4",
                "url": "https://rest_api_root/v1/ship/ticket_state_plans/63feb3da9cc1ead1d2be93f4"
            },
            "form_state": {
                "id": "63bb744214bd13c9def24ca2",
                "url": "https://rest_api_root/v1/ship/ticket_states/63bb744214bd13c9def24ca2",
                "name": "待处理",
                "type": "pending"
            },
            "to_state": {
                "id": "63bb744214bd13c9def24ca4",
                "url": "https://rest_api_root/v1/ship/ticket_states/63bb744214bd13c9def24ca4",
                "name": "处理中",
                "type": "in_progress"
            }
        }
    ]
}
```

---

### 部分更新一个工单属性

**`PATCH /v1/ship/ticket_properties/{property_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `property_id` | String | 工单属性的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `name` *(可选)* | String | 工单属性的名称。在一个企业中这个名称是唯一的。 |
| `options` *(可选)* | Object[] | 选择项列表。options是整体更新的。 |
| `options._id` *(可选)* | String | 选择项id。该值在选择项中是唯一的。 |
| `options.text` | String | 选择项内容。该值在选择项中是唯一的。 |
| `options.parent_id` *(可选)* | String | 选择项父级id。当属性类型为cascade_select,cascade_multi_select时，parent_id参数有效，用于构建级联类型选择项之间的父子关系，最多存在4级。 |

**响应示例（select）：**

```json
{
    "id": "severity-update",
    "url": "https://rest_api_root/v1/ship/ticket_properties/severity",
    "name": "严重程度-update",
    "type": "select",
    "options": [
        {
            "_id": "5efb1859110533727a82c603",
            "text": "严重-update"
        },
        {
            "_id": "5efb1859110533727a82c624",
            "text": "一般"
        }
    ],
    "is_removable": true,
    "is_name_editable": true,
    "is_options_editable": true,
    "select_all_level": false,
    "display_all_level": false,
    "display_separator": null
}
```

**响应示例（cascade_select）：**

```json
{
    "id": "jiliandanxuan",
    "url": "https://rest_api_root/v1/ship/ticket_properties/jiliandanxuan",
    "name": "级联单选-update",
    "type": "cascade_select",
    "options": [
        {
            "_id": "64b9f741eec7fd94e63b411e",
            "text": "父-update"
        },
        {
            "_id": "64b9f741eec7fd94e63b411f",
            "text": "子",
            "parent_id": "64b9f741eec7fd94e63b411e"
        }
    ],
    "is_removable": true,
    "is_name_editable": true,
    "is_options_editable": true,
    "select_all_level": false,
    "display_all_level": false,
    "display_separator": "/"
}
```

---

### 部分更新一个工单状态

**`PATCH /v1/ship/ticket_states/{ticket_state_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `ticket_state_id` | String | 工单状态id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `name` *(可选)* | String | 工单状态的名称，在一个企业中这个名称是唯一的。 |
| `type` *(可选)* | String | 工单状态的类型。 |

**响应示例：**

```json
{
    "id": "6422711c3f12e6c1e46d40f2",
    "url": "https://rest_api_root/v1/ship/ticket_states/6422711c3f12e6c1e46d40f2",
    "name": "已完成",
    "type": "completed",
    "color": "#56AB25"
}
```

---

## 工时

### 创建一个工时

**`POST /v1/workloads`**

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `principal_type` | String | 工时主体的类型。 |
| `principal_id` | String | 工时主体的id。 |
| `type_id` *(可选)* | String | 工时类型的id。 |
| `duration` | Number | 工时的时长。单位是小时，数值可以是为0-24之间，最多包含一位小数的正数。 |
| `report_at` | Number | 工时的登记日期。该值为十位数字组成的时间戳，会被转换为该时间当天的零点零分零秒。 |
| `report_by_id` *(可选)* | String | 工时的登记人，企业鉴权时必填。个人鉴权时不需要传递，即使传递了也会被忽略。 |
| `recorded_at` *(可选)* | String | 工时的登记时间，默认是当前时间。 |
| `description` *(可选)* | String | 工时的说明。 |

**响应示例：**

```json
{
    "id": "5f168f764eba01a5278b87cd",
    "url": "https://rest_api_root/v1/workloads/5f168f764eba01a5278b87cd",
    "principal_type": "work_item",
    "principal": {
        "id": "5edca524cad2fa1125cb0630",
        "url": "https://rest_api_root/v1/project/work_items/5edca524cad2fa1125cb0630",
        "identifier": "SCR-5",
        "title": "这是一个缺陷",
        "type": "bug",
        "start_at": 1583290309,
        "end_at": 1583290347,
        "parent_id": "5edca524cad2fa1125cb0629",
        "short_id": "c9WqLmTO",
        "html_url": "https://yctech.pingcode.com/pjm/workitems/c9WqLmTO",
        "properties": {
            "prop_a": "prop_a_value",
            "prop_b": "prop_b_value"
        }
    },
    "type": {
        "id": "5a86eaf6a72585327ea46fge0",
        "url": "https://rest_api_root/v1/workload_types/5a86eaf6a72585327ea46fge0",
        "name": "研发"
    },
    "duration": 8,
    "review_state": "approved",
    "description": "这是一个工时",
    "report_at": 1593290347,
    "report_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "created_at": 1593290347,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    }
}
```

---

### 删除一个工时

**`DELETE /v1/workloads/{workload_id}`**

用户令牌只能删除用户自己登记的工时记录。

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `workload_id` | String | 工时的id。 |

**响应示例：**

```json
{
    "id": "5f168f764eba01a5278b87cd",
    "url": "https://rest_api_root/v1/workloads/5f168f764eba01a5278b87cd",
    "principal_type": "work_item",
    "principal": {
        "id": "5edca524cad2fa1125cb0630",
        "url": "https://rest_api_root/v1/project/work_items/5edca524cad2fa1125cb0630",
        "identifier": "SCR-5",
        "title": "这是一个缺陷",
        "type": "bug",
        "start_at": 1583290309,
        "end_at": 1583290347,
        "parent_id": "5edca524cad2fa1125cb0629",
        "short_id": "c9WqLmTO",
        "html_url": "https://yctech.pingcode.com/pjm/workitems/c9WqLmTO",
        "properties": {
            "prop_a": "prop_a_value",
            "prop_b": "prop_b_value"
        }
    },
    "type": {
        "id": "5a86eaf6a72585327ea46fge0",
        "url": "https://rest_api_root/v1/workload_types/5a86eaf6a72585327ea46fge0",
        "name": "研发"
    },
    "duration": 8,
    "review_state": "approved",
    "description": "这是一个工时",
    "report_at": 1593290347,
    "report_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "created_at": 1593290347,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    }
}
```

---

### 获取工时列表

**`GET /v1/workloads`**

**查询参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `principal_type` *(可选)* | String | 工时主体的类型。当查询参数含有pilot_id或principal_id时，principal_type参数必填。 |
| `pilot_id` *(可选)* | String | 工时主体所在产品、项目或测试库的id。使用该参数过滤数据时，登记日期查询的起始时间和登记日期查询的结束时间的跨度最大为3个月。 |
| `principal_id` *(可选)* | String | 工时主体的id。 |
| `start_at` *(可选)* | Number | 登记日期查询的起始时间。开始时间会转换为对应日期的开始时间点。开始时间和结束时间须同时存在。 |
| `end_at` *(可选)* | Number | 登记日期查询的结束时间。结束时间会转换为对应日期的结束时间点。 |
| `report_by_id` *(可选)* | String | 登记人的id。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
         {
             "id": "5f168f764eba01a5278b87cd",
             "url": "https://rest_api_root/v1/workloads/5f168f764eba01a5278b87cd",
             "principal_type": "work_item",
             "principal": {
                 "id": "5edca524cad2fa1125cb0630",
                 "url": "https://rest_api_root/v1/project/work_items/5edca524cad2fa1125cb0630",
                 "identifier": "SCR-5",
                 "title": "这是一个缺陷",
                 "type": "bug",
                 "start_at": 1583290309,
                 "end_at": 1583290347,
                 "parent_id": "5edca524cad2fa1125cb0629",
                 "short_id": "c9WqLmTO",
                 "html_url": "https://yctech.pingcode.com/pjm/workitems/c9WqLmTO",
                 "properties": {
                     "prop_a": "prop_a_value",
                     "prop_b": "prop_b_value"
                 }
             },
             "type": {
                 "id": "5a86eaf6a72585327ea46fge0",
                 "url": "https://rest_api_root/v1/workload_types/5a86eaf6a72585327ea46fge0",
                 "name": "研发"
             },
             "duration": 8,
             "review_state": "approved",
             "description": "这是一个工时",
             "report_at": 1593290347,
             "report_by": {
                 "id": "a0417f68e846aae315c85d24643678a9",
                 "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                 "name": "john",
                 "display_name": "John",
                 "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
             },
             "created_at": 1593290347,
             "created_by": {
                 "id": "a0417f68e846aae315c85d24643678a9",
                 "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                 "name": "john",
                 "display_name": "John",
                 "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
             }
         }
    ]
}
```

---

### 获取工时类型列表

**`GET /v1/workload_types`**

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "5a86eaf6a72585327ea46fge0",
            "url": "https://rest_api_root/v1/workload_types/5a86eaf6a72585327ea46fge0",
            "name": "研发"
        }
    ]
}
```

---

### 部分更新一个工时

**`PATCH /v1/workloads/{workload_id}`**

用户令牌只能更新属于用户自己登记的工时记录。

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `workload_id` | String | 工时的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `type_id` *(可选)* | String | 工时类型的id。 |
| `duration` *(可选)* | Number | 工时的时长。单位是小时，数值可以是为0-24之间，最多包含一位小数的正数。 |
| `report_at` *(可选)* | Number | 工时的登记日期。该值为十位数字组成的时间戳，会被转换为该时间当天的零点零分零秒。 |
| `description` *(可选)* | String | 工时的说明。 |

**响应示例：**

```json
{
    "id": "5f168f764eba01a5278b87cd",
    "url": "https://rest_api_root/v1/workloads/5f168f764eba01a5278b87cd",
    "principal_type": "work_item",
    "principal": {
        "id": "5edca524cad2fa1125cb0630",
        "url": "https://rest_api_root/v1/project/work_items/5edca524cad2fa1125cb0630",
        "identifier": "SCR-5",
        "title": "这是一个缺陷",
        "type": "bug",
        "start_at": 1583290309,
        "end_at": 1583290347,
        "parent_id": "5edca524cad2fa1125cb0629",
        "short_id": "c9WqLmTO",
        "html_url": "https://yctech.pingcode.com/pjm/workitems/c9WqLmTO",
        "properties": {
            "prop_a": "prop_a_value",
            "prop_b": "prop_b_value"
        }
    },
    "type": {
        "id": "5a86eaf6a72585327ea46fge0",
        "url": "https://rest_api_root/v1/workload_types/5a86eaf6a72585327ea46fge0",
        "name": "研发"
    },
    "duration": 8,
    "review_state": "approved",
    "description": "这是一个工时",
    "report_at": 1593290347,
    "report_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "created_at": 1593290347,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    }
}
```

---

## 托管平台

### 全量更新一个托管平台

**`PUT /v1/scm/products/{product_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `product_id` | String | 代码托管平台的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `name` | String | 代码托管平台的名称，在一个企业中这个名称是唯一的。 |
| `type` | String | 代码托管平台的产品类型，主要用于显示图标。 |
| `description` *(可选)* | String | 代码托管平台简介。 |

**响应示例：**

```json
{
    "id": "564587fe700d43b81b080765",
    "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765",
    "name": "Github",
    "type": "github",
    "description": "Github公有云"
}
```

---

### 创建一个托管平台

**`POST /v1/scm/products`**

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `name` | String | 代码托管平台的名称，在一个企业中这个名称是唯一的。 |
| `type` | String | 代码托管平台的产品类型，主要用于显示图标。 |
| `description` *(可选)* | String | 代码托管平台的简介 |

**响应示例：**

```json
{
    "id": "564587fe700d43b81b080765",
    "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765",
    "name": "Github",
    "type": "github",
    "description": "Github公有云"
}
```

---

### 获取托管平台列表

**`GET /v1/scm/products`**

**查询参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `name` *(可选)* | String | 代码托管平台的名称。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "564587fe700d43b81b080765",
            "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765",
            "name": "Github",
            "type": "github",
            "description": "Github公有云"
        }
    ]
}
```

---

### 部分更新一个托管平台

**`PATCH /v1/scm/products/{product_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `product_id` | String | 代码托管平台的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `name` *(可选)* | String | 代码托管平台的名称，在一个企业中这个名称是唯一的。 |
| `type` *(可选)* | String | 代码托管平台的产品类型，主要用于显示图标。 |
| `description` *(可选)* | String | 代码托管平台简介。 |

**响应示例：**

```json
{
    "id": "564587fe700d43b81b080765",
    "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765",
    "name": "Github",
    "type": "github",
    "description": "Github公有云"
}
```

---

## 托管平台用户

### 全量更新一个托管平台用户

**`PUT /v1/scm/products/{product_id}/users/{user_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `product_id` | String | 代码托管平台的id。 |
| `user_id` | String | 代码托管平台上的用户id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `name` | String | 代码托管平台上的用户名，同一代码托管平台下，该用户名是唯一的。 |
| `display_name` *(可选)* | String | 用户显示名。 |
| `html_url` *(可选)* | String | 代码托管平台上的用户主页地址。 |
| `avatar_url` *(可选)* | String | 代码托管平台上的用户头像地址。 |

**响应示例：**

```json
{
    "id": "5666aea91f99e33cb7c44964",
    "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/users/5666aea91f99e33cb7c44964",
    "product": {
        "id": "564587fe700d43b81b080765",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765",
        "name": "Github",
        "type": "github"
    },
    "name": "terry",
    "display_name": "Terry",
    "html_url": "https://github.com/terrylee",
    "avatar_url": "https://avatars2.githubusercontent.com/u/694592?v=4"
}
```

---

### 创建一个托管平台用户

**`POST /v1/scm/products/{product_id}/users`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `product_id` | String | 代码托管平台的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `name` | String | 代码托管平台上的用户名，同一代码托管平台下，该用户名是唯一的。 |
| `display_name` *(可选)* | String | 用户显示名。 |
| `html_url` *(可选)* | String | 代码托管平台上的用户主页地址。 |
| `avatar_url` *(可选)* | String | 代码托管平台上的用户头像地址。 |

**响应示例：**

```json
{
    "id": "5666aea91f99e33cb7c44964",
    "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/users/5666aea91f99e33cb7c44964",
    "product": {
        "id": "564587fe700d43b81b080765",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765",
        "name": "Github",
        "type": "github"
    },
    "name": "terry",
    "display_name": "Terry",
    "html_url": "https://github.com/terrylee",
    "avatar_url": "https://avatars2.githubusercontent.com/u/694592?v=4"
}
```

---

### 获取托管平台用户列表

**`GET /v1/scm/products/{product_id}/users`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `product_id` | String | 代码托管平台的id。 |

**查询参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `name` *(可选)* | String | 代码托管平台上的用户名。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "5666aea91f99e33cb7c44964",
            "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/users/5666aea91f99e33cb7c44964",
            "product": {
                "id": "564587fe700d43b81b080765",
                "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765",
                "name": "Github",
                "type": "github"
            },
            "name": "terry",
            "display_name": "Terry",
            "html_url": "https://github.com/terrylee",
            "avatar_url": "https://avatars2.githubusercontent.com/u/694592?v=4"
        }
    ]
}
```

---

### 部分更新一个托管平台用户

**`PATCH /v1/scm/products/{product_id}/users/{user_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `product_id` | String | 代码托管平台的id。 |
| `user_id` | String | 代码托管平台上的用户id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `name` *(可选)* | String | 代码托管平台上的用户名，同一代码托管平台下，该用户名是唯一的。 |
| `display_name` *(可选)* | String | 用户显示名。 |
| `html_url` *(可选)* | String | 代码托管平台上的用户主页地址。 |
| `avatar_url` *(可选)* | String | 代码托管平台上的用户头像地址。 |

**响应示例：**

```json
{
    "id": "5666aea91f99e33cb7c44964",
    "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/users/5666aea91f99e33cb7c44964",
    "product": {
        "id": "564587fe700d43b81b080765",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765",
        "name": "Github",
        "type": "github"
    },
    "name": "terry",
    "display_name": "Terry",
    "html_url": "https://github.com/terrylee",
    "avatar_url": "https://avatars2.githubusercontent.com/u/694592?v=4"
}
```

---

## 执行用例配置

### 获取全部执行用例执行结果列表

**`GET /v1/testhub/run_statuses`**

**响应示例：**

```json
{
    "page_index": 0,
    "page_size": 30,
    "total": 5,
    "values": [
        {
            "id": "68d117800d5dd2484a198261",
            "url": "https://rest_api_root/v1/testhub/run_statuses/68d117800d5dd2484a198261",
            "name": "通过",
            "is_system": "true"
        },
        {
            "id": "68d117800d5dd2484a198262",
            "url": "https://rest_api_root/v1/testhub/run_statuses/68d117800d5dd2484a198262",
            "name": "受阻",
            "is_system": "true"
        },
        {
            "id": "68d117800d5dd2484a198263",
            "url": "https://rest_api_root/v1/testhub/run_statuses/68d117800d5dd2484a198263",
            "name": "失败",
            "is_system": "true"
        },
        {
            "id": "68d117800d5dd2484a198264",
            "url": "https://rest_api_root/v1/testhub/run_statuses/68d117800d5dd2484a198264",
            "name": "跳过",
            "is_system": "true"
        },
        {
            "id": "68d117800d5dd2484a198265",
            "url": "https://rest_api_root/v1/testhub/run_statuses/68d117800d5dd2484a198265",
            "name": "未测",
            "is_system": "true"
        }
    ]
}
```

---

## 拉取请求

### 全量更新一个拉取请求

**`PUT /v1/scm/products/{product_id}/repositories/{repository_id}/pull_requests/{pull_request_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `product_id` | String | 代码托管平台的id。 |
| `repository_id` | String | 代码仓库的id。 |
| `pull_request_id` | String | 拉取请求的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `title` | String | 拉取请求的标题。 |
| `creator_name` | String | 拉取请求创建者的用户名。 |
| `source_branch_id` | String | 源分支的id。 |
| `target_branch_id` | String | 目标分支的id。 |
| `status` | String | 拉取请求的状态。 |
| `description` *(可选)* | String | 拉取请求的描述。 |
| `merged_at` *(可选)* | Number | 拉取请求合并的时间。当拉取请求状态为merged时，该值为必填。 |
| `merged_commit_sha` *(可选)* | String | 源分支最后一次提交的SHA值。当拉取请求状态为merged时，该值为必填。 |
| `merged_by_name` *(可选)* | String | 拉取请求合并者的用户名。当拉取请求状态为merged时，该值为必填。 |
| `comments_count` *(可选)* | Number | 拉取请求的评论数量。 |
| `review_comments_count` *(可选)* | Number | 代码评审的评论数量。 |
| `commits_count` *(可选)* | Number | 代码的提交数量。 |
| `additions_count` *(可选)* | Number | 新增文件的数量。 |
| `deletions_count` *(可选)* | Number | 删除文件的数量。 |
| `changed_files_count` *(可选)* | Number | 更改文件的数量。 |
| `work_item_identifiers` *(可选)* | String[] | PingCode工作项编号的列表。通过该参数将拉取请求与一个或多个工作项关联。 |

**响应示例：**

```json
{
   "id": "594587fe700d43b81b080789",
   "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766/pull_requests/594587fe700d43b81b080789",
   "product": {
      "id": "564587fe700d43b81b080765",
      "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765",
      "name": "Github",
      "type": "github"
   },
   "repository": {
      "id": "564587fe700d43b81b080766",
      "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766",
      "name": "ngx-planet",
      "full_name": "worktile/ngx-planet",
      "created_at": 1403018919
   },
   "title": "fix(doc): #PLM-001 fix document title",
   "number": 1,
   "status": "merged",
   "description": "Please give some great suggestions",
   "author": {
      "id": "5666aea91f99e33cb7c44964",
      "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/users/5666aea91f99e33cb7c44964",
      "name": "terry"
   },
   "source_branch": {
      "id": "564587fe700d43b81b080767",
      "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766/branches/564587fe700d43b81b080767",
      "name": "terry/#PLM-001",
      "create_at": 1403018919
   },
   "target_branch": {
      "id": "564587fe700d43b81b080776",
      "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766/branches/564587fe700d43b81b080776",
      "name": "develop",
      "create_at": 1402018919
   },
   "created_at": 1403014000,
   "merged_at": 1473018919,
   "merged_commit_sha": "96a024347146ebdc5f481f45e6e6871e0c43af5f",
   "merged_by": {
      "id": "5666aea91f99e33cb7c44968",
      "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/users/5666aea91f99e33cb7c44964",
      "name": "terry"
   },
   "comments_count": 2,
   "review_comments_count": 2,
   "commits_count": 2,
   "additions_count": 0,
   "deletions_count": 0,
   "changed_files_count": 3,
   "work_items": [
      {
          "id": "564587fe700d43b81b080ab8",
          "url": "https://rest_api_root/v1/project/work_items/564587fe700d43b81b080ab8",
          "identifier": "PLM-001",
          "title": "这是一个用户故事",
          "type": "story",
          "start_at": 1583290309,
          "end_at": 1583290347,
          "parent_id": "5edca524cad2fa112b06105c",
          "short_id": "c9WqLmTO",
          "html_url": "https://yctech.pingcode.com/pjm/workitems/c9WqLmTO",
          "properties": {
              "prop_a": "prop_a_value",
              "prop_b": "prop_b_value"
          }
      }
   ]
}
```

---

### 创建一个拉取请求

**`POST /v1/scm/products/{product_id}/repositories/{repository_id}/pull_requests`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `product_id` | String | 代码托管平台的id。 |
| `repository_id` | String | 代码仓库的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `title` | String | 拉取请求的标题。 |
| `number` | Number | 拉取请求的编号。同一代码仓库下，该值是唯一的。 |
| `creator_name` | String | 拉取请求创建者的用户名。 |
| `source_branch_id` *(可选)* | String | 源分支的id。 |
| `target_branch_id` | String | 目标分支的id。 |
| `status` | String | 拉取请求的状态。 |
| `description` *(可选)* | String | 拉取请求的描述。 |
| `merged_at` *(可选)* | Number | 拉取请求合并的时间。当拉取请求状态为merged时，该值为必填。 |
| `merged_commit_sha` *(可选)* | String | 源分支最后一次提交的SHA值。当拉取请求状态为merged时，该值为必填。 |
| `merged_by_name` *(可选)* | String | 拉取请求合并者的用户名。当拉取请求状态为merged时，该值为必填。 |
| `comments_count` *(可选)* | Number | 拉取请求的评论数量。 |
| `review_comments_count` *(可选)* | Number | 代码评审的评论数量。 |
| `commits_count` *(可选)* | Number | 代码的提交数量。 |
| `additions_count` *(可选)* | Number | 新增文件的数量。 |
| `deletions_count` *(可选)* | Number | 删除文件的数量。 |
| `changed_files_count` *(可选)* | Number | 更改文件的数量。 |
| `work_item_identifiers` *(可选)* | String[] | PingCode工作项编号的列表。通过该参数将拉取请求与一个或多个工作项关联。 |

**响应示例：**

```json
{
  "id": "594587fe700d43b81b080789",
  "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766/pull_requests/594587fe700d43b81b080789",
  "product": {
    "id": "564587fe700d43b81b080765",
    "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765",
    "name": "Github",
    "type": "github"
  },
  "repository": {
    "id": "564587fe700d43b81b080766",
    "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766",
    "name": "ngx-planet",
    "full_name": "worktile/ngx-planet",
    "created_at": 1403018919
  },
  "title": "fix(doc): #PLM-001 fix document title",
  "number": 1,
  "status": "merged",
  "description": "Please give some great suggestions",
  "author": {
    "id": "5666aea91f99e33cb7c44964",
    "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/users/5666aea91f99e33cb7c44964",
    "name": "terry"
  },
  "source_branch": {
    "id": "564587fe700d43b81b080767",
    "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766/branches/564587fe700d43b81b080767",
    "name": "terry/#PLM-001",
    "create_at": 1403018919
  },
  "target_branch": {
    "id": "564587fe700d43b81b080776",
    "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766/branches/564587fe700d43b81b080776",
    "name": "develop",
    "create_at": 1402018919
  },
      "created_at": 1463014000,
  "merged_at": 1473018919,
  "merged_commit_sha": "96a024347146ebdc5f481f45e6e6871e0c43af5f",
  "merged_by": {
    "id": "5666aea91f99e33cb7c44964",
    "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/users/5666aea91f99e33cb7c44964",
    "name": "terry"
  },
  "comments_count": 2,
  "review_comments_count": 2,
  "commits_count": 2,
  "additions_count": 0,
  "deletions_count": 0,
  "changed_files_count": 3,
  "work_items": [
    {
      "id": "564587fe700d43b81b080ab8",
      "url": "https://rest_api_root/v1/project/work_items/564587fe700d43b81b080ab8",
      "identifier": "PLM-001",
      "title": "这是一个用户故事",
      "type": "story",
      "start_at": 1583290309,
      "end_at": 1583290347,
      "parent_id": "5edca524cad2fa112b06105c",
      "short_id": "c9WqLmTO",
      "html_url": "https://yctech.pingcode.com/pjm/workitems/c9WqLmTO",
      "properties": {
        "prop_a": "prop_a_value",
        "prop_b": "prop_b_value"
      }
    }
  ]
}
```

---

### 获取拉取请求列表

**`GET /v1/scm/products/{product_id}/repositories/{repository_id}/pull_requests`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `product_id` | String | 代码托管平台的id。 |
| `repository_id` | String | 代码仓库的id。 |

**查询参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `number` *(可选)* | number | 拉取请求的编号。 |
| `work_item_id` *(可选)* | String | 工作项的id。 |

**响应示例：**

```json
{
     "page_size": 30,
     "page_index": 0,
     "total": 1,
     "values": [
         {
             "id": "594587fe700d43b81b080789",
             "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766/pull_requests/594587fe700d43b81b080789",
             "product": {
                "id": "564587fe700d43b81b080765",
                "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765",
                "name": "Github",
                "type": "github"
             },
             "repository": {
                "id": "564587fe700d43b81b080766",
                "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766",
                "name": "ngx-planet",
                "full_name": "worktile/ngx-planet",
                "created_at": 1403018919
             },
             "title": "fix(doc): #PLM-001 fix document title",
             "number": 1,
             "status": "merged",
             "description": "Please give some great suggestions",
             "author": {
                "id": "5666aea91f99e33cb7c44964",
                "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/users/5666aea91f99e33cb7c44964",
                "name": "terry"
             },
             "source_branch": {
                "id": "564587fe700d43b81b080767",
                "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766/branches/564587fe700d43b81b080767",
                "name": "terry/#PLM-001",
                "create_at": 1403018919
             },
             "target_branch": {
                "id": "564587fe700d43b81b080776",
                "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766/branches/564587fe700d43b81b080776",
                "name": "develop",
                "create_at": 1402018919
             },
             "created_at": 1403014000,
             "merged_at": 1473018919,
             "merged_commit_sha": "96a024347146ebdc5f481f45e6e6871e0c43af5f",
             "merged_by": {
                "id": "5666aea91f99e33cb7c44964",
                "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/users/5666aea91f99e33cb7c44964",
                "name": "terry"
             },
             "comments_count": 2,
             "review_comments_count": 2,
             "commits_count": 2,
             "additions_count": 0,
             "deletions_count": 0,
             "changed_files_count": 3,
             "work_items": [
                {
                    "id": "564587fe700d43b81b080ab8",
                    "url": "https://rest_api_root/v1/project/work_items/564587fe700d43b81b080ab8",
                    "identifier": "PLM-001",
                    "title": "这是一个用户故事",
                    "type": "story",
                    "start_at": 1583290309,
                    "end_at": 1583290347,
                    "parent_id": "5edca524cad2fa112b06105c",
                    "short_id": "c9WqLmTO",
                    "html_url": "https://yctech.pingcode.com/pjm/workitems/c9WqLmTO",
                    "properties": {
                        "prop_a": "prop_a_value",
                        "prop_b": "prop_b_value"
                    }
                }
             ]
         }
     ]
}
```

---

### 部分更新一个拉取请求

**`PATCH /v1/scm/products/{product_id}/repositories/{repository_id}/pull_requests/{pull_request_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `product_id` | String | 代码托管平台的id。 |
| `repository_id` | String | 代码仓库的id。 |
| `pull_request_id` | String | 拉取请求的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `status` | String | 拉取请求的状态。 |
| `title` *(可选)* | String | 拉取请求的标题。 |
| `creator_name` *(可选)* | String | 拉取请求创建者的用户名。 |
| `description` *(可选)* | String | 拉取请求的描述。 |
| `target_branch_id` *(可选)* | String | 目标分支的id。 |
| `source_branch_id` *(可选)* | String | 源分支的id。 |
| `merged_at` *(可选)* | Number | 拉取请求合并的时间。当拉取请求状态为merged时，该值为必填。 |
| `merged_commit_sha` *(可选)* | String | 源分支最后一次提交的SHA值。当拉取请求状态为merged时，该值为必填。 |
| `merged_by_name` *(可选)* | String | 拉取请求合并者的用户名。当拉取请求状态为merged时，该值为必填。 |
| `comments_count` *(可选)* | Number | 拉取请求的评论数量。 |
| `review_comments_count` *(可选)* | Number | 代码评审的评论数量。 |
| `commits_count` *(可选)* | Number | 代码的提交数量。 |
| `additions_count` *(可选)* | Number | 新增文件的数量。 |
| `deletions_count` *(可选)* | Number | 删除文件的数量。 |
| `changed_files_count` *(可选)* | Number | 更改文件的数量。 |
| `work_item_identifiers` *(可选)* | String[] | PingCode工作项编号的列表。通过该参数将拉取请求与一个或多个工作项关联。 |

**响应示例：**

```json
{
   "id": "594587fe700d43b81b080789",
   "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766/pull_requests/594587fe700d43b81b080789",
   "product": {
      "id": "564587fe700d43b81b080765",
      "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765",
      "name": "Github",
      "type": "github"
   },
   "repository": {
      "id": "564587fe700d43b81b080766",
      "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766",
      "name": "ngx-planet",
      "full_name": "worktile/ngx-planet",
      "created_at": 1403018919
   },
   "title": "fix(doc): #PLM-001 fix document title",
   "number": 1,
   "status": "merged",
   "description": "Please give some great suggestions",
   "author": {
      "id": "5666aea91f99e33cb7c44964",
      "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/users/5666aea91f99e33cb7c44964",
      "name": "terry"
   },
   "source_branch": {
      "id": "564587fe700d43b81b080767",
      "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766/branches/564587fe700d43b81b080767",
      "name": "terry/#PLM-001",
      "create_at": 1403018919
   },
   "target_branch": {
      "id": "564587fe700d43b81b080776",
      "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766/branches/564587fe700d43b81b080776",
      "name": "develop",
      "create_at": 1402018919
   },
   "created_at": 1403014000,
   "merged_at": 1473018919,
   "merged_commit_sha": "96a024347146ebdc5f481f45e6e6871e0c43af5f",
   "merged_by": {
      "id": "5666aea91f99e33cb7c44964",
      "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/users/5666aea91f99e33cb7c44964",
      "name": "terry"
   },
   "comments_count": 2,
   "review_comments_count": 2,
   "commits_count": 2,
   "additions_count": 0,
   "deletions_count": 0,
   "changed_files_count": 3,
   "work_items": [
      {
          "id": "564587fe700d43b81b080ab8",
          "url": "https://rest_api_root/v1/project/work_items/564587fe700d43b81b080ab8",
          "identifier": "PLM-001",
          "title": "这是一个用户故事",
          "type": "story",
          "start_at": 1583290309,
          "end_at": 1583290347,
          "parent_id": "5edca524cad2fa112b06105c",
          "short_id": "c9WqLmTO",
          "html_url": "https://yctech.pingcode.com/pjm/workitems/c9WqLmTO",
          "properties": {
              "prop_a": "prop_a_value",
              "prop_b": "prop_b_value"
          }
      }
   ]
}
```

---

## 授权码

### 刷新用户令牌

**`GET /v1/auth/token?grant_type=refresh_token`**

通过refresh_token重新获取access_token，可以避免用户频繁重新授权。refresh_token的有效期为90天，删除PingCode的应用或重置应用的Secret都会导致refresh_token失效。

**查询参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `grant_type` | String | OAuth2的授予类型，这里固定为：refresh_token。 |
| `refresh_token` | String | 通过authorization_code获取access_token时，一起返回的refresh_token。 |

**响应示例：**

```json
{
   "access_token": "e7321ca8-f724-4abd-9169-d76d095c6acf",
   "token_type": "Bearer",
   "expires_in": 1577808000
}
```

---

### 获取用户令牌

**`GET /v1/auth/token?grant_type=authorization_code`**

access_token的有效期为30天，删除PingCode的应用或重置应用的Secret都会导致access_token失效。

**查询参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `grant_type` | String | OAuth2的授予类型，这里固定为：authorization_code。 |
| `client_id` | String | PingCode应用的Client ID |
| `client_secret` | String | PingCode应用的Secret |
| `code` | String | 用户授权后，在回调地址中拿到的code值。 |

**响应示例：**

```json
{
   "access_token": "e7321ca8-f724-4abd-9169-d76d095c6acf",
   "refresh_token": "f724-4abd-9169-e7321ca8-d76d095c6acf",
   "token_type": "Bearer",
   "expires_in": 1577808000
}
```

---

### 请求授权

**`GET https://oauth2_root/authorize?response_type=code`**

用户访问请求授权页面前，需要登录PingCode，否则页面自动跳转到PingCode的登录页面。需要注意，私有环境的授权页面地址为：https://xxxxxx/oauth2/authorize 。

**查询参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `response_type` | String | 响应类型，这里固定为code类型。 |
| `client_id` | String | PingCode应用的Client ID |

---

## 提交

### 创建一个提交

**`POST /v1/scm/commits`**

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `sha` | String | 提交的SHA值。 |
| `message` | String | 提交的描述信息。 |
| `committer_name` | String | 提交者的用户名。 |
| `committed_at` | Number | 提交的时间。 |
| `tree_id` *(可选)* | String | 提交树的SHA值。 |
| `files_added` | String[] | 新增文件的文件名列表。 |
| `files_removed` | String[] | 移除文件的文件名列表。 |
| `files_modified` | String[] | 更新文件的文件名列表。 |
| `work_item_identifiers` *(可选)* | String[] | PingCode工作项编号的列表。通过该参数将提交与一个或多个工作项关联。 |

**响应示例：**

```json
{
    "id": "5e3bb2128cfda459bbafa3fb",
    "url": "https://rest_api_root/v1/scm/commits/5e3bb2128cfda459bbafa3fb",
    "sha": "96a024347146ebdc5f481f45e6e6871e0c43af5f",
    "message": "feat(scope): #PLM-001 initialization code structure",
    "committer_name": "terry",
    "committed_at": 1403018919,
    "tree_id": "1bf8989985e70389c07daa5052464a9c6f4896bb",
    "files_added": [
        "index.ts"
    ],
    "files_removed": [
        "utilities.ts"
    ],
    "files_modified": [
        "README.md"
    ],
    "file_changed_count": 3,
    "work_items": [
        {
            "id": "564587fe700d43b81b080ab8",
            "url": "https://rest_api_root/v1/project/work_items/564587fe700d43b81b080ab8",
            "identifier": "PLM-001",
            "title": "这是一个用户故事",
            "type": "story",
            "start_at": 1583290309,
            "end_at": 1583290347,
            "parent_id": "5edca524cad2fa112b06105c",
            "short_id": "c9WqLmTO",
            "html_url": "https://yctech.pingcode.com/pjm/workitems/c9WqLmTO",
            "properties": {
                "prop_a": "prop_a_value",
                "prop_b": "prop_b_value"
            }
        }
    ]
}
```

---

### 获取提交列表

**`GET /v1/scm/commits`**

**查询参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `sha` *(可选)* | String | 提交的SHA值。 |
| `work_item_id` *(可选)* | String | 工作项的id。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "5e3bb2128cfda459bbafa3fb",
            "url": "https://rest_api_root/v1/scm/commits/5e3bb2128cfda459bbafa3fb",
            "sha": "96a024347146ebdc5f481f45e6e6871e0c43af5f",
            "message": "feat(scope): #PLM-001 initialization code structure",
            "committer_name": "terry",
            "committed_at": 1403018919,
            "tree_id": "1bf8989985e70389c07daa5052464a9c6f4896bb",
            "files_added": [
                "index.ts"
            ],
            "files_removed": [
                "utilities.ts"
            ],
            "files_modified": [
                "README.md"
            ],
            "file_changed_count": 3,
            "work_items": [
                {
                    "id": "564587fe700d43b81b080ab8",
                    "url": "https://rest_api_root/v1/project/work_items/564587fe700d43b81b080ab8",
                    "identifier": "PLM-001",
                    "title": "这是一个用户故事",
                    "type": "story",
                    "start_at": 1583290309,
                    "end_at": 1583290347,
                    "parent_id": "5edca524cad2fa112b06105c",
                    "short_id": "c9WqLmTO",
                    "html_url": "https://yctech.pingcode.com/pjm/workitems/c9WqLmTO",
                    "properties": {
                        "prop_a": "prop_a_value",
                        "prop_b": "prop_b_value"
                    }
                }
            ]
        }
    ]
}
```

---

## 提交引用

### 创建一个提交引用

**`POST /v1/scm/products/{product_id}/repositories/{repository_id}/refs`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `product_id` | String | 代码托管平台的id。 |
| `repository_id` | String | 代码仓库的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `sha` | String | 提交的SHA值。 |
| `meta_type` | String | 引用实体的类型。 |
| `meta_id` | String | 引用实体的id，例如：branch_id。 |

**响应示例：**

```json
{
    "id": "5e451b7dd704c212f7de8b4f",
    "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766/refs/5e451b7dd704c212f7de8b4f",
    "product": {
        "id": "564587fe700d43b81b080765",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765",
        "name": "Github",
        "type": "github"
    },
    "repository": {
        "id": "564587fe700d43b81b080766",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766",
        "name": "ngx-planet",
        "full_name": "worktile/ngx-planet",
        "created_at": 1403018919
    },
    "commit": {
        "id": "5e3bb2128cfda459bbafa3fb",
        "url": "https://rest_api_root/v1/scm/commits/5e3bb2128cfda459bbafa3fb",
        "sha": "96a024347146ebdc5f481f45e6e6871e0c43af5f",
        "message": "feat(scope): #PLM-001 initialization code structure",
        "committer_name": "terry",
        "committed_at": 1403018919
    },
    "meta": {
        "id": "564587fe700d43b81b080767",
        "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766/branches/564587fe700d43b81b080767",
        "name": "terry/#PLM-001",
        "type": "branch"
    }
}
```

---

### 获取提交引用列表

**`GET /v1/scm/products/{product_id}/repositories/{repository_id}/refs`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `product_id` | String | 代码托管平台的id。 |
| `repository_id` | String | 代码仓库的id。 |

**查询参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `meta_type` | String | 引用实体的类型。 |
| `meta_id` | String | 引用实体的id。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "5e451b7dd704c212f7de8b4f",
            "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766/refs/5e451b7dd704c212f7de8b4f",
            "product": {
                "id": "564587fe700d43b81b080765",
                "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765",
                "name": "Github",
                "type": "github"
            },
            "repository": {
                "id": "564587fe700d43b81b080766",
                "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766",
                "name": "ngx-planet",
                "full_name": "worktile/ngx-planet",
                "created_at": 1403018919
            },
            "commit": {
                "id": "5e3bb2128cfda459bbafa3fb",
                "url": "https://rest_api_root/v1/scm/commits/5e3bb2128cfda459bbafa3fb",
                "sha": "96a024347146ebdc5f481f45e6e6871e0c43af5f",
                "message": "feat(scope): #PLM-001 initialization code structure",
                "committer_name": "terry",
                "committed_at": 1403018919
            },
            "meta": {
                "id": "564587fe700d43b81b080767",
                "url": "https://rest_api_root/v1/scm/products/564587fe700d43b81b080765/repositories/564587fe700d43b81b080766/branches/564587fe700d43b81b080767",
                "name": "terry/#PLM-001",
                "type": "branch"
            }
        }
    ]
}
```

---

## 日志

### 获取审计日志列表

**`GET /v1/security/audit_logs`**

**查询参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `operated_between` | String | 操作时间介于的时间范围，通过','分割起始时间。 |
| `operated_bys` *(可选)* | String | 操作人的列表，使用','分割，最多只能20个。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "5fca7b74efab845a2fa53181",
            "url": "https://rest_api_root/v1/security/audit_logs/5fca7b74efab845a2fa53181",
            "operated_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "operated_at": 1676454024,
            "operate_object": "规则",
            "application": "自动化",
            "ip": "45.251.20.42",
            "summary": "修改规则",
            "detail": "规则：规则1\n类型：自动化规则\n描述：-"
        }
    ]
}
```

---

### 获取登录日志列表

**`GET /v1/security/login_logs`**

**查询参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `logged_between` | String | 登录时间介于的时间范围，通过','分割起始时间。 |
| `user_ids` *(可选)* | String | 成员id的列表，使用','分割，最多只能20个。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "5fca7b74efab845a2fa53181",
            "url": "https://rest_api_root/v1/security/login_logs/5fca7b74efab845a2fa53181",
            "user": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "login_method": "账号密码",
            "region": "北京海淀区",
            "ip": "45.251.20.42",
            "user_agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36",
            "created_at": 1676454024
        }
    ]
}
```

---

## 构建

### 构建记录

企业内实际的构建记录，用于在PingCode中显示构建的详细信息。 资源地址： GET https://rest_api_root/v1/build/builds/{build_id}

**全量数据示例：**

```json
{
  "id": "564587fe700d43b81b080765",
  "url": "https://rest_api_root/v1/build/builds/564587fe700d43b81b080765",
  "identifier": "131",
  "name": "unit-test",
  "job_url": "https://your-job-url",
  "provider": "jenkins",
  "result_overview": "1000 test cases pass",
  "result_url": "https://your-result-url",
  "start_at": 1583290309,
  "status": "success",
  "end_at": 1583290347,
  "duration": 38,
  "work_items": [
    {
      "id": "564587fe700d43b81b080ab8",
      "url": "https://rest_api_root/v1/project/work_items/564587fe700d43b81b080ab8",
      "identifier": "PLM-001",
      "title": "这是一个用户故事",
      "type": "story",
      "start_at": 1583290309,
      "end_at": 1583290347,
      "parent_id": "5edca524cad2fa112b06105c",
      "short_id": "c9WqLmTO",
      "html_url": "https://yctech.pingcode.com/pjm/workitems/c9WqLmTO",
      "properties": {
        "prop_a": "prop_a_value",
        "prop_b": "prop_b_value"
      }
    }
  ]
}
```

**引用数据示例：**

```json
{
  "id": "564587fe700d43b81b080765",
  "url": "https://rest_api_root/v1/build/builds/564587fe700d43b81b080765",
  "identifier": "131",
  "name": "unit-test",
  "job_url": "https://your-job-url",
  "provider": "jenkins",
  "result_overview": "1000 test cases pass",
  "result_url": "https://your-result-url",
  "start_at": 1583290309,
  "status": "success",
  "end_at": 1583290347,
  "duration": 38
}
```

---

## 构建记录

### 全量更新一条构建记录

**`PUT /v1/build/builds/{build_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `build_id` | String | 构建的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `name` | String | 构建的名称。 |
| `identifier` | String | 构建的编号。 |
| `job_url` *(可选)* | String | 构建任务的地址。如果为空，在PingCode中不显示相关跳转链接。 |
| `provider` | String | 构建工具的名称。 |
| `result_overview` *(可选)* | String | 构建结果的概述。 |
| `result_url` *(可选)* | String | 构建结果的地址。如果为空，在PingCode中不显示相关的跳转链接。 |
| `start_at` | Number | 构建开始的时间。 |
| `end_at` | Number | 构建结束的时间。 |
| `duration` | Number | 构建持续的时间。单位为秒。 |
| `status` | String | 构建的状态。 |
| `work_item_identifiers` *(可选)* | String[] | PingCode工作项编号的列表。通过该参数将构建与一个或多个工作项关联。 |

**响应示例：**

```json
{
   "id": "564587fe700d43b81b080765",
   "url": "https://rest_api_root/v1/build/builds/564587fe700d43b81b080765",
   "identifier": "131",
   "name": "unit-test",
   "job_url": "https://your-job-url",
   "provider": "jenkins",
   "result_overview": "1000 test cases pass",
   "result_url": "https://your-result-url",
   "start_at": 1583290309,
   "status": "success",
   "end_at": 1583290347,
   "duration": 38,
   "work_items": [
      {
          "id": "564587fe700d43b81b080ab8",
          "url": "https://rest_api_root/v1/project/work_items/564587fe700d43b81b080ab8",
          "identifier": "PLM-001",
          "title": "这是一个用户故事",
          "type": "story",
          "start_at": 1583290309,
          "end_at": 1583290347,
          "parent_id": "5edca524cad2fa112b06105c",
          "short_id": "c9WqLmTO",
          "html_url": "https://yctech.pingcode.com/pjm/workitems/c9WqLmTO",
          "properties": {
              "prop_a": "prop_a_value",
              "prop_b": "prop_b_value"
          }
      }
   ]
}
```

---

### 创建一条构建记录

**`POST /v1/build/builds`**

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `name` | String | 构建的名称。 |
| `identifier` | String | 构建的编号。 |
| `job_url` *(可选)* | String | 构建任务的地址。如果为空，在PingCode中不显示相关跳转链接。 |
| `provider` | String | 构建工具的名称。 |
| `result_overview` *(可选)* | String | 构建结果的概述。 |
| `result_url` *(可选)* | String | 构建结果的地址。如果为空，在PingCode中不显示相关的跳转链接。 |
| `start_at` | Number | 构建开始的时间。 |
| `end_at` | Number | 构建结束的时间。 |
| `duration` | Number | 构建持续的时间。单位为秒。 |
| `status` | String | 构建的状态。 |
| `work_item_identifiers` *(可选)* | String[] | PingCode工作项编号的列表。通过该参数将构建与一个或多个工作项关联。 |

**响应示例：**

```json
{
   "id": "564587fe700d43b81b080765",
   "url": "https://rest_api_root/v1/build/builds/564587fe700d43b81b080765",
   "name": "unit-test",
   "identifier": "131",
   "job_url": "https://your-job-url",
   "provider": "jenkins",
   "result_overview": "1000 test cases pass",
   "result_url": "https://your-result-url",
   "start_at": 1583290309,
   "status": "success",
   "end_at": 1583290347,
   "duration": 38,
   "work_items": [
      {
          "id": "564587fe700d43b81b080ab8",
          "url": "https://rest_api_root/v1/project/work_items/564587fe700d43b81b080ab8",
          "identifier": "PLM-001",
          "title": "这是一个用户故事",
          "type": "story",
          "start_at": 1583290309,
          "end_at": 1583290347,
          "parent_id": "5edca524cad2fa112b06105c",
          "short_id": "c9WqLmTO",
          "html_url": "https://yctech.pingcode.com/pjm/workitems/c9WqLmTO",
          "properties": {
              "prop_a": "prop_a_value",
              "prop_b": "prop_b_value"
          }
      }
   ]
}
```

---

### 删除一条构建记录

**`DEL /v1/build/builds/{build_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `build_id` | String | 构建的id。 |

**响应示例：**

```json
{
   "id": "564587fe700d43b81b080765",
   "url": "https://rest_api_root/v1/build/builds/564587fe700d43b81b080765",
   "identifier": "131",
   "name": "unit-test",
   "job_url": "https://your-job-url",
   "provider": "jenkins",
   "result_overview": "1000 test cases pass",
   "result_url": "https://your-result-url",
   "start_at": 1583290309,
   "status": "success",
   "end_at": 1583290347,
   "duration": 38,
   "work_items": [
     {
        "id": "564587fe700d43b81b080ab8",
        "url": "https://rest_api_root/v1/project/work_items/564587fe700d43b81b080ab8",
        "identifier": "PLM-001",
        "title": "这是一个用户故事",
        "type": "story",
        "start_at": 1583290309,
        "end_at": 1583290347,
        "parent_id": "5edca524cad2fa112b06105c",
        "short_id": "c9WqLmTO",
        "html_url": "https://yctech.pingcode.com/pjm/workitems/c9WqLmTO",
        "properties": {
            "prop_a": "prop_a_value",
            "prop_b": "prop_b_value"
        }
     }
   ]
}
```

---

### 获取构建记录列表

**`GET /v1/build/builds`**

**响应示例：**

```json
{
   "page_size": 30,
   "page_index": 0,
   "total": 1,
   "values": [
       {
            "id": "564587fe700d43b81b080765",
            "url": "https://rest_api_root/v1/build/builds/564587fe700d43b81b080765",
            "identifier": "131",
            "name": "unit-test",
            "job_url": "https://your-job-url",
            "provider": "jenkins",
            "result_overview": "1000 test cases pass",
            "result_url": "https://your-result-url",
            "start_at": 1583290309,
            "status": "success",
            "end_at": 1583290347,
            "duration": 38,
            "work_items": [
                  {
                       "id": "564587fe700d43b81b080ab8",
                       "url": "https://rest_api_root/v1/project/work_items/564587fe700d43b81b080ab8",
                       "identifier": "PLM-001",
                       "title": "这是一个用户故事",
                       "type": "story",
                       "start_at": 1583290309,
                       "end_at": 1583290347,
                       "parent_id": "5edca524cad2fa112b06105c",
                       "short_id": "c9WqLmTO",
                       "html_url": "https://yctech.pingcode.com/pjm/workitems/c9WqLmTO",
                       "properties": {
                           "prop_a": "prop_a_value",
                           "prop_b": "prop_b_value"
                       }
                  }
              ]
       }
   ]
}
```

---

### 部分更新一条构建记录

**`PATCH /v1/build/builds/{build_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `build_id` | String | 构建的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `name` *(可选)* | String | 构建的名称。 |
| `identifier` *(可选)* | String | 构建的编号。 |
| `job_url` *(可选)* | String | 构建任务的地址。如果为空，在PingCode中不显示相关跳转链接。 |
| `provider` *(可选)* | String | 构建工具的名称。 |
| `result_overview` *(可选)* | String | 构建结果的概述。 |
| `result_url` *(可选)* | String | 构建结果的地址。如果为空，在PingCode中不显示相关的跳转链接。 |
| `start_at` *(可选)* | Number | 构建开始的时间。 |
| `end_at` *(可选)* | Number | 构建结束的时间。 |
| `status` *(可选)* | String | 构建的状态。 |
| `duration` *(可选)* | Number | 构建持续的时间。单位为秒。 |
| `work_item_identifiers` *(可选)* | String[] | PingCode工作项编号的列表。通过该参数将构建与一个或多个工作项关联。 |

**响应示例：**

```json
{
   "id": "564587fe700d43b81b080765",
   "url": "https://rest_api_root/v1/build/builds/564587fe700d43b81b080765",
   "identifier": "131",
   "name": "unit-test",
   "job_url": "https://your-job-url",
   "provider": "jenkins",
   "result_overview": "1000 test cases pass",
   "result_url": "https://your-result-url",
   "start_at": 1583290309,
   "status": "failure",
   "end_at": 1583290347,
   "duration": 38,
   "work_items": [
      {
          "id": "564587fe700d43b81b080ab8",
          "url": "https://rest_api_root/v1/project/work_items/564587fe700d43b81b080ab8",
          "identifier": "PLM-001",
          "title": "这是一个用户故事",
          "type": "story",
          "start_at": 1583290309,
          "end_at": 1583290347,
          "parent_id": "5edca524cad2fa112b06105c",
          "short_id": "c9WqLmTO",
          "html_url": "https://yctech.pingcode.com/pjm/workitems/c9WqLmTO",
          "properties": {
              "prop_a": "prop_a_value",
              "prop_b": "prop_b_value"
          }
      }
   ]
}
```

---

## 概述

### URI结构

PingCode REST API通过URI路径提供对资源的访问，使用{}将URI路径的一部分标记为可使用参数替换的部分，URI路径遵循以下规则： https://rest_api_root/v1[/{area}]/{resource} 例如： https://open.pingcode.com/v1/scm/productshttps://open.pingcode.com/v1/scm/products/{product_id}/repositorieshttps://open.pingcode.com/v1/release/environments rest_api_root表示REST API的根路径，在不同的环境中rest_api_root值有所不同： 公有云环境的rest_api_root值为：https://open.pingcode.com私有部署环境的rest_api_root值为：https://xxxxxx/open oauth2_root表示OAuth2页面的根路径，在不同的环境中oauth2_root值也有所不同： 公有云环境的oauth2_root值为：https://open.pingcode.com/oauth2私有部署环境的oauth2_root值为：https://xxxxxx/oauth2

---

### 使用方式

PingCode REST API支持 OPTIONS/GET/PUT/PATCH/POST/DELETE等标准的HTTP请求。 对于GET/DELETE请求，通过querystring传递参数；对于POST/PUT/PATCH请求，需要在headers中添加&quot;content-type&quot;: &quot;application/json&quot;，然后通过body传递参数。 PingCode REST API使用HTTP状态码指示已执行操作的状态； 使用response body传递数据。 单个资源 当创建、更新、获取、删除单个资源成功时，会返回当前操作的资源。 HTTP状态码：201Body：{     "id": "5e05d8448f8461dada9ba29c",     "url": "https://rest_api_root/v1/{resource}",     "name": "资源名称",     "desc": "资源简介",     "created_at": 1578897962} 分页数据 当请求多条数据时，默认每一页返回30条，最大返回100条。 通过在querystring中设置page_size和page_index，指定每一页的数据量和第几页的数据（page_index为0时，表示第一页）。 在返回的数据结构中，page_size表示当前每页的数据量，page_index表示当前在第几页，total表示资源总数量，values表示资源的数组。 HTTP状态码：200Body：{     "page_size": 30,     "page_index": 0,     "total": 100,     "values": [         {             "id": "5e05d8448f8461dada9ba29c",             "url": "https://rest_api_root/v1/{resource}",             "name": "资源名称",             "desc": "资源简介",             "created_at": 1578897962         },         ...     ]} 错误 当请求失败时，会返回错误码和错误信息。 HTTP状态码：500Body：{     "code": "100000",     "message": "Internal Server Error"}

---

### 数据结构

PingCode REST API使用json作为通讯格式，所有时间均使用10位数字组成的时间戳。 PingCode REST API为每一种资源定义两种数据结构，全量结构和引用结构。 全量结构包含资源的所有属性，引用结构只包含必要属性。当获取单个资源或分页获取资源列表时，PingCode REST API将返回全量结构； 当获取其他资源引用当前资源时，PingCode REST API将返回引用结构。 全量结构 {     "id": "5e05d8448f8461dada9ba29c",     "url": "https://rest_api_root/v1/{resource}",     "name": "资源名称",     "desc": "资源简介",     "created_at": 1578897962} 引用结构 {     "id": "5e05d8448f8461dada9ba29c",     "url": "https://rest_api_root/v1/{resource}",     "name": "资源名称"}

---

### 欢迎使用

欢迎使用PingCode Representational State Transfer APIs （简称PingCode REST API）。 PingCode REST API用于通过HTTP与PingCode服务端进行远程交互，例如创建、修改、查询、删除PingCode的资源。

---

### 频率限制

PingCode REST API限制使用者的请求频率，目的是保障核心服务的可靠且响应迅速。频率限制不用于区分客户和服务级别。 具体策略 根据使用者的身份标识，PingCode REST API最多允许每位使用者每分钟请求200次，单位分钟内超出限制数量的HTTP请求将统一返回错误信息。 HTTP状态码：429Headers：{     "x-pc-retry-after": 50}Body：{     "code": "100038",     "message": "请求频率过高"} x-pc-retry-after指示使用者在重新请求之前必须等待的秒数。如果使用者在到期之前重新发出请求，则请求会再次失败并返回相同的HTTP状态码和response body。 合理请求 由于频率限制的存在，最小化请求将十分必要，一个显而易见的策略是缓存不会轻易变更的数据。 另外使用PingCode Flow中的发送Webhook和发送HTTP请求来将PingCode中发生变更的数据发送给订阅者，也可以有效降低 PingCode REST API的请求数量，从而降低遇到频率限制的风险。

---

## 活动记录

### 获取活动记录列表

**`GET /v1/activities?principal_type={principal_type}&principal_id={principal_id}`**

**查询参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `principal_type` | String | 主体的类型。 |
| `principal_id` | String | 主体的id。 |

**响应示例：**

```json
{
    "page_index": 0,
    "page_size": 30,
    "total": 1,
    "values": [
        {
            "id": "694ae20fdb8e0baef70f7ddb",
            "url": "https://rest_api_root/v1/activities/694ae20fdb8e0baef70f7ddb?principal_type=idea&principal_id=683562430d684517b06b814b",
            "template": "update-property",
            "type": "update",
            "summary": "修改了引用多选",
            "content": {
                "property_key": "yinyongduoxuan",
                "origin": null,
                "target": [
                    {
                        "id": "65fa797d8f0358d376233220",
                        "name": "REST API 产品"
                    }
                ]
            },
            "client": "web",
            "created_at": 1766515215,
            "created_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            }
        }
    ]
}
```

---

## 测试库

### 创建一个测试库

**`POST /v1/testhub/libraries`**

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `scope_type` *(可选)* | String | 测试库的所属类型。默认值organization。允许值分别表示企业可见和团队可见。 |
| `scope_id` *(可选)* | String | 测试库的所属id。当scope_type为user_group时，该字段必填，且表示团队id；当scope_type为其他值时，该字段无效。 |
| `name` | String | 测试库的名称。 |
| `visibility` *(可选)* | String | 测试库的可见范围。允许值分别表示组织全部成员可见和仅项目成员可见。 |
| `identifier` | String | 测试库的标识。在一个企业中这个标识是唯一的。项目的标识由大写英文字母/数字/下划线/连接线组成（不超过15个字符）。 |
| `description` *(可选)* | String | 测试库的描述。 |
| `members` *(可选)* | Object[] | 测试库成员的列表。当测试库的scope_type变为user_group时，测试库成员必须是scope_id指定的团队内的成员；当scope_type为其他值时，测试库成员可以是任意成员或团队。 |
| `members.id` | String | 企业成员或团队的id。 |
| `members.type` | String | 测试库成员的类型。 |

**响应示例：**

```json
{
    "id": "5eb623f6a70571487ea47000",
    "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000",
    "identifier": "CSK",
    "name": "测试库",
    "scope_type": "user_group",
    "scope_id": "63c8fb32729dee3334d96af7",
    "visibility": "private",
    "color": "#F693E7",
    "description": "这是一个测试库",
    "members": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/members/a0417f68e846aae315c85d24643678a9",
            "type": "user",
            "user": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            }
        }
    ],
    "created_at": 1583290300,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "updated_at": 1583290300,
    "updated_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "is_archived": 0,
    "is_deleted": 0
}
```

---

### 向测试库中添加一个成员

**`POST /v1/testhub/libraries/{library_id}/members`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `library_id` | String | 测试库的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `member` *(可选)* | Object | 测试库的成员。 |
| `member.id` | String | 企业成员或团队的id。 |
| `member.type` | String | 成员的类型。 |
| `role_id` *(可选)* | String | 角色的id。 |

**响应示例：**

```json
{
    "id": "a0417f68e846aae315c85d24643678a9",
    "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/members/a0417f68e846aae315c85d24643678a9",
    "library": {
        "id": "5eb623f6a70571487ea47000",
        "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000",
        "identifier": "CSK",
        "name": "测试库",
        "is_archived": 0,
        "is_deleted": 0
    },
    "type": "user",
    "user": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "role": {
        "id": "6422711c3f12e6c1e46d40e6",
        "url": "https://rest_api_root/v1/directory/roles/6422711c3f12e6c1e46d40e6",
        "name": "管理员"
    }
}
```

---

### 向测试库中添加一个用例模块

**`POST /v1/testhub/libraries/{library_id}/suites`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `library_id` | String | 测试库的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `name` | String | 测试模块名称。测试模块为树形结构，同一层次下的名称不能重复。 |
| `parent_id` *(可选)* | String | 父测试模块的id。 |

**响应示例：**

```json
{
    "id": "55714870a70ea4eb623f6700",
    "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/suites/55714870a70ea4eb623f6700",
    "library": {
        "id": "5eb623f6a70571487ea47000",
        "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000",
        "identifier": "CSK",
        "name": "测试库",
        "is_archived": 0,
        "is_deleted": 0
    },
    "name": "登录",
    "parent": {
        "id": "5eb623f6a70571487ea46999",
        "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/suites/5eb623f6a70571487ea46999",
        "name": "用户"
    },
    "paths": "首页/用户"
}
```

---

### 在测试库中移除一个成员

**`DELETE /v1/testhub/libraries/{library_id}/members/{member_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `library_id` | String | 测试库的id。 |
| `member_id` | String | 企业成员或团队的id。 |

**响应示例：**

```json
{
    "id": "a0417f68e846aae315c85d24643678a9",
    "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/members/a0417f68e846aae315c85d24643678a9",
    "library": {
        "id": "5eb623f6a70571487ea47000",
        "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000",
        "identifier": "CSK",
        "name": "测试库",
        "is_archived": 0,
        "is_deleted": 0
    },
    "type": "user",
    "user": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "role": {
        "id": "6422711c3f12e6c1e46d40e6",
        "url": "https://rest_api_root/v1/directory/roles/6422711c3f12e6c1e46d40e6",
        "name": "管理员"
    }
}
```

---

### 在测试库中移除一个用例模块

**`DELETE /v1/testhub/libraries/{library_id}/suites/{suite_id}`**

请注意，删除一个模块会自动删除其所有的子模块。

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `library_id` | String | 测试库的id。 |
| `suite_id` | String | 测试模块的id。 |

**响应示例：**

```json
{
    "id": "55714870a70ea4eb623f6701",
    "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/suites/55714870a70ea4eb623f6701",
    "library": {
        "id": "5eb623f6a70571487ea47000",
        "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000",
        "identifier": "CSK",
        "name": "测试库",
        "is_archived": 0,
        "is_deleted": 0
    },
    "name": "注册",
    "parent": {
        "id": "5eb623f6a70571487ea46999",
        "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/suites/5eb623f6a70571487ea46999",
        "name": "用户"
    },
    "paths": "首页/用户"
}
```

---

### 获取测试库中的成员列表

**`GET /v1/testhub/libraries/{library_id}/members`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `library_id` | String | 测试库的id。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 2,
    "values": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/members/a0417f68e846aae315c85d24643678a9",
            "library": {
                "id": "5eb623f6a70571487ea47000",
                "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000",
                "identifier": "CSK",
                "name": "测试库",
                "is_archived": 0,
                "is_deleted": 0
            },
            "type": "user",
            "user": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "role": {
                "id": "6422711c3f12e6c1e46d40e6",
                "url": "https://rest_api_root/v1/directory/roles/6422711c3f12e6c1e46d40e6",
                "name": "管理员"
            }
        },
        {
            "id": "63c8fb32729dee3334d96af7",
            "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/members/63c8fb32729dee3334d96af7",
            "library": {
                "id": "5eb623f6a70571487ea47000",
                "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000",
                "identifier": "CSK",
                "name": "测试库",
                "is_archived": 0,
                "is_deleted": 0
            },
            "type": "user_group",
            "user_group": {
                "id": "63c8fb32729dee3334d96af7",
                "url": "https://rest_api_root/v1/directory/groups/63c8fb32729dee3334d96af7",
                "name": "Open Team"
            },
            "role": {
                "id": "6422711c3f12e6c1e46d40e6",
                "url": "https://rest_api_root/v1/directory/roles/6422711c3f12e6c1e46d40e6",
                "name": "管理员"
            }
        }
    ]
}
```

---

### 获取测试库中的用例模块列表

**`GET /v1/testhub/libraries/{library_id}/suites`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `library_id` | String | 测试库的id。 |

**查询参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `parent_id` *(可选)* | String | 父测试模块的id。值可以是root或者某个模块id，当为空时，获取到所有的模块，当为root时，获取到所有的一级模块，当为某个模块id时，获取到该模块的直属子模块。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "55714870a70ea4eb623f6700",
            "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/suites/55714870a70ea4eb623f6700",
            "library": {
                "id": "5eb623f6a70571487ea47000",
                "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000",
                "identifier": "CSK",
                "name": "测试库",
                "is_archived": 0,
                "is_deleted": 0
            },
            "name": "登录",
            "parent": {
                "id": "5eb623f6a70571487ea46999",
                "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/suites/5eb623f6a70571487ea46999",
                "name": "用户"
            },
            "paths": "首页/用户"
        }
    ]
}
```

---

### 获取测试库列表

**`GET /v1/testhub/libraries`**

**查询参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `name` *(可选)* | String | 测试库的名称。 |
| `scope_type` *(可选)* | String | 测试库的所属类型。 |
| `scope_id` *(可选)* | String | 测试库的所属id。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "5eb623f6a70571487ea47000",
            "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000",
            "identifier": "CSK",
            "name": "测试库",
            "scope_type": "user_group",
            "scope_id": "63c8fb32729dee3334d96af7",
            "visibility": "private",
            "color": "#F693E7",
            "description": "这是一个测试库",
            "members": [
                {
                    "id": "a0417f68e846aae315c85d24643678a9",
                    "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/members/a0417f68e846aae315c85d24643678a9",
                    "type": "user",
                    "user": {
                        "id": "a0417f68e846aae315c85d24643678a9",
                        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                        "name": "john",
                        "display_name": "John",
                        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
                    }
                }
            ],
            "created_at": 1583290300,
            "created_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "updated_at": 1583290300,
            "updated_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "is_archived": 0,
            "is_deleted": 0
        }
    ]
}
```

---

### 部分更新一个测试库

**`PATCH /v1/testhub/libraries/{library_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `library_id` | String | 测试库的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `name` *(可选)* | String | 测试库的名称。 |
| `identifier` *(可选)* | String | 测试库的标识。在一个企业中这个标识是唯一的。项目的标识由大写英文字母/数字/下划线/连接线组成（不超过15个字符）。 |
| `description` *(可选)* | String | 测试库的描述。 |

**响应示例：**

```json
{
    "id": "5eb623f6a70571487ea47000",
    "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000",
    "identifier": "CSK",
    "name": "测试库",
    "scope_type": "user_group",
    "scope_id": "63c8fb32729dee3334d96af7",
    "visibility": "private",
    "color": "#F693E7",
    "description": "这是一个测试库",
    "members": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/members/a0417f68e846aae315c85d24643678a9",
            "type": "user",
            "user": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            }
        }
    ],
    "created_at": 1583290300,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "updated_at": 1583290300,
    "updated_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "is_archived": 0,
    "is_deleted": 0
}
```

---

### 部分更新一个测试库中用例模块

**`PATCH /v1/testhub/libraries/{library_id}/suites/{suite_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `library_id` | String | 测试库的id。 |
| `suite_id` | String | 测试模块的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `name` *(可选)* | String | 测试模块名称。测试模块为树形结构，同一层次下的名称不能重复。 |
| `parent_id` *(可选)* | String | 父测试模块的id。 |

**响应示例：**

```json
{
    "id": "55714870a70ea4eb623f6700",
    "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/suites/55714870a70ea4eb623f6700",
    "library": {
        "id": "5eb623f6a70571487ea47000",
        "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000",
        "identifier": "CSK",
        "name": "测试库",
        "is_archived": 0,
        "is_deleted": 0
    },
    "name": "登录",
    "parent": {
        "id": "5eb623f6a70571487ea46999",
        "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/suites/5eb623f6a70571487ea46999",
        "name": "用户"
    },
    "paths": "首页/用户"
}
```

---

### 部分更新一个测试库内的成员

**`PATCH /v1/testhub/libraries/{library_id}/members/{member_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `library_id` | String | 测试库的id。 |
| `member_id` | String | 企业成员或团队的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `role_id` *(可选)* | String | 角色的id。管理员可以更改其他用户的角色，但非管理员用户无权更改其他用户的角色。 |

**响应示例：**

```json
{
    "id": "a0417f68e846aae315c85d24643678a9",
    "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/members/a0417f68e846aae315c85d24643678a9",
    "library": {
        "id": "5eb623f6a70571487ea47000",
        "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000",
        "identifier": "CSK",
        "name": "测试库",
        "is_archived": 0,
        "is_deleted": 0
    },
    "type": "user",
    "user": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "role": {
        "id": "6422711c3f12e6c1e46d40e6",
        "url": "https://rest_api_root/v1/directory/roles/6422711c3f12e6c1e46d40e6",
        "name": "管理员"
    }
}
```

---

## 测试管理

### 测试库

用于查看和管理测试库的基本信息。 资源地址：{GET} https://rest_api_root/v1/testhub/libraries/{library_id}

**全量数据示例：**

```json
{
    "id": "5eb623f6a70571487ea47000",
    "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000",
    "identifier": "CSK",
    "name": "测试库",
    "scope_type": "user_group",
    "scope_id": "63c8fb32729dee3334d96af7",
    "visibility": "private",
    "color": "#F693E7",
    "description": "这是一个测试库",
    "members": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/members/a0417f68e846aae315c85d24643678a9",
            "type": "user",
            "user": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            }
        }
    ],
    "created_at": 1583290300,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "updated_at": 1583290300,
    "updated_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "is_archived": 0,
    "is_deleted": 0
}
```

**引用数据示例：**

```json
{
    "id": "5eb623f6a70571487ea47000",
    "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000",
    "identifier": "CSK",
    "name": "测试库",
    "is_archived": 0,
    "is_deleted": 0
}
```

---

### 测试配置中心

用于查看和管理测试用例属性。

---

### 用例

用于查看和管理测试用例的基本信息。 资源地址：{GET} https://rest_api_root/v1/testhub/cases/{case_id}

**全量数据示例：**

```json
{
    "id": "5edca524cad2fa112b06305c",
    "url": "https://rest_api_root/v1/testhub/cases/5edca524cad2fa112b06305c",
    "library": {
        "id": "5eb623f6a70571487ea47000",
        "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000",
        "identifier": "CSK",
        "name": "测试库",
        "is_archived": 0,
        "is_deleted": 0
    },
    "identifier": "CSK-10",
    "title": "这是一个测试用例",
    "level": "P1",
    "short_id": "fdUw3C",
    "html_url": "https://yctech.pingcode.com/testhub/cases/fdUw3C",
    "important_level": {
        "id": "57a109b35ae8c20630fd7256",
        "url": "https://rest_api_root/v1/testhub/case_important_levels/57a109b35ae8c20630fd7256",
        "name": "P1"
    },
    "suite": {
        "id": "55714870a70ea4eb623f6700",
        "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/suites/55714870a70ea4eb623f6700",
        "name": "登录",
        "paths": "首页/账户"
    },
    "state": {
        "id": "686f62038668bbae4f4dd0c1",
        "url": "https://rest_api_root/v1/testhub/case_states/686f62038668bbae4f4dd0c1",
        "name": "设计",
        "type": "pending"
    },
    "type": {
        "id": "5cf189b35de9c20620ad7153",
        "url": "https://rest_api_root/v1/testhub/case_types/5cf189b35de9c20620ad7153",
        "name": "功能测试"
    },
    "maintenance": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "test_type": "automation",
    "description": "测试用例的备注",
    "precondition": "前置条件的描述信息",
    "properties": {
        "prop_a": "prop_a_value",
        "prop_b": "prop_b_value"
    },
    "estimated_workload": 8,
    "remaining_workload": 5,
    "steps": [
        {
             "step_id": "524cad5edb06305cca2fa111",
             "description": "网页访问",
             "expected_value": null,
             "is_group": 1,
             "group_id": null
        },
        {
             "step_id": "524cad5edb06305cca2fa112",
             "description": "在浏览器地址栏中输入https://pingcode.com",
             "expected_value": "成功访问PingCode官网",
             "is_group": 0,
             "group_id": "524cad5edb06305cca2fa111"
        }
    ],
    "participants": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "url": "https://rest_api_root/v1/participants/a0417f68e846aae315c85d24643678a9?principal_type=test_case&principal_id=5edca524cad2fa112b06305c",
            "type": "user",
            "user": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            }
        }
    ],
    "public_image_token": "IcF1VmJFF-UIi9lMU18m1NFFAruWXYx0ZAm90pJ2LGk",
    "created_at": 1583290300,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "updated_at": 1583290300,
    "updated_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "is_archived": 0,
    "is_deleted": 0
}
```

**引用数据示例：**

```json
{
    "id": "5edca524cad2fa112b06305c",
    "url": "https://rest_api_root/v1/testhub/cases/5edca524cad2fa112b06305c",
    "identifier": "CSK-10",
    "title": "这是一个测试用例",
    "level": "P1",
    "short_id": "fdUw3C",
    "html_url": "https://yctech.pingcode.com/testhub/cases/fdUw3C",
    "properties": {
        "prop_a": "prop_a_value",
        "prop_b": "prop_b_value"
    }
}
```

---

### 计划

用于查看和管理测试计划的基本信息。 资源地址：{GET} https://rest_api_root/v1/testhub/libraries/{library_id}/plans/{plan_id}

**全量数据示例：**

```json
{
    "id": "5eb6a70571487623fea47000",
    "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/plans/5eb6a70571487623fea47000",
    "library": {
        "id": "5eb623f6a70571487ea47000",
        "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000",
        "identifier": "CSK",
        "name": "测试库",
        "is_archived": 0,
        "is_deleted": 0
    },
    "name": "测试计划",
    "state": {
        "id": "652d0cb2b798f983d9c67c2b",
        "url": "http://rest_api_root/v1/testhub/plan_states/652d0cb2b798f983d9c67c2b",
        "name": "进行中",
        "type": "in_progress"
    },
    "start_at": 1589791860,
    "end_at": 1589791870,
    "short_id": "7nNkViOD",
    "html_url": "https://yctech.pingcode.com/testhub/libraries/CSK/plans/7nNkViOD/runs",
    "type": {
        "id": "641d0ab2b998f883f9c67b2c",
        "url": "http://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/plan_types/641d0ab2b998f883f9c67b2c",
        "name": "发布测试"
    },
    "project": {
        "id": "5eb623f6a70571487ea41919",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919",
        "identifier": "SCR",
        "name": "Scrum项目",
        "type": "scrum",
        "is_archived": 0,
        "is_deleted": 0
    },
    "sprint": null,
    "version": {
        "id": "5eb623487ea47001f6a70571",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919/versions/5eb623487ea47001f6a70571",
        "name": "1.0.0",
        "start_at": 1565255712,
        "end_at": 1565255879,
        "stage": {
            "id": "5f44a8f8bb347b14b49624a1",
            "url": "https://rest_api_root/v1/project/stages/5f44a8f8bb347b14b49624a1",
            "name": "未开始",
            "type": "pending",
            "color": "#FA8888"
        }
    },
    "assignee": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "summary": "测试报告的概要",
    "created_at": 1565366200,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "updated_at": 1565366200,
    "updated_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    }

}
```

**引用数据示例：**

```json
{
    "id": "5eb6a70571487623fea47000",
    "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/plans/5eb6a70571487623fea47000",
    "name": "测试计划",
    "status": "in_progress",
    "start_at": 1589791860,
    "end_at": 1589791870,
    "short_id": "7nNkViOD",
    "html_url": "https://yctech.pingcode.com/testhub/libraries/CSK/plans/7nNkViOD/runs"
}
```

---

## 测试配置中心

### 执行用例配置

用于查看和管理执行用例属性。

---

### 用例配置

用于查看和管理测试用例属性。

---

### 计划配置

用于查看和管理计划属性。

---

## 瀑布

### 向瀑布项目中添加一个工作项类型

**`POST /v1/project/projects/{project_id}/work_item_types`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `project_id` | String | 项目id。项目类型必须为waterfall，并且项目开启了本地配置。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `id` | String | 工作项类型的id，不支持增加计划类型（里程碑、阶段）。 |

**响应示例：**

```json
{
    "id": "630da48bc9443b1aa94ce3ea",
    "url": "https://rest_api_root/v1/project/projects/6375cc81e3004de4ea14aa52/work_item_types/630da48bc9443b1aa94ce3ea",
    "project": {
        "id": "6375cc81e3004de4ea14aa52",
        "url": "http://rest_api_root/v1/project/projects/6375cc81e3004de4ea14aa52",
        "identifier": "WTF",
        "name": "瀑布项目",
        "type": "waterfall",
        "is_archived": 0,
        "is_deleted": 0
    },
    "work_item_type": {
        "id": "630da48bc9443b1aa94ce3ea",
        "url": "https://rest_api_root/v1/project/waterfall/work_item_types/630da48bc9443b1aa94ce3ea",
        "name": "需求",
        "group": "requirement"
    }
}
```

---

### 在瀑布项目中移除一个工作项类型

**`DELETE /v1/project/projects/{project_id}/work_item_types/{work_item_type_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `project_id` | String | 项目的id。项目类型必须为waterfall，并且项目开启了本地配置。 |
| `work_item_type_id` | String | 工作项类型的id，不支持移除计划类型（里程碑、阶段）。 |

**响应示例：**

```json
{
    "id": "630da48bc9443b1aa94ce3ea",
    "url": "https://rest_api_root/v1/project/projects/6375cc81e3004de4ea14aa52/work_item_types/630da48bc9443b1aa94ce3ea",
    "project": {
        "id": "6375cc81e3004de4ea14aa52",
        "url": "http://rest_api_root/v1/project/projects/6375cc81e3004de4ea14aa52",
        "identifier": "WTF",
        "name": "瀑布项目",
        "type": "waterfall",
        "is_archived": 0,
        "is_deleted": 0
    },
    "work_item_type": {
        "id": "630da48bc9443b1aa94ce3ea",
        "url": "https://rest_api_root/v1/project/waterfall/work_item_types/630da48bc9443b1aa94ce3ea",
        "name": "需求",
        "group": "requirement"
    }
}
```

---

### 获取瀑布项目中的工作项类型列表

**`GET /v1/project/projects/{project_id}/work_item_types`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `project_id` | String | 项目的id。项目类型必须为waterfall，并且项目开启了本地配置。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `group` | String | 工作项类型的分组。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 6,
    "values": [
        {
            "id": "630da48bc9443b1aa94ce3ea",
            "url": "https://rest_api_root/v1/project/projects/6375cc81e3004de4ea14aa52/work_item_types/630da48bc9443b1aa94ce3ea",
            "project": {
                "id": "6375cc81e3004de4ea14aa52",
                "url": "http://rest_api_root/v1/project/projects/6375cc81e3004de4ea14aa52",
                "identifier": "WTF",
                "name": "瀑布项目",
                "type": "waterfall",
                "is_archived": 0,
                "is_deleted": 0
            },
            "work_item_type": {
                "id": "630da48bc9443b1aa94ce3ea",
                "url": "https://rest_api_root/v1/project/projects/6375cc81e3004de4ea14aa52/work_item_types/630da48bc9443b1aa94ce3ea",
                "name": "需求",
                "group": "requirement"
            }
        },
        {
            "id": "630da48bc9443b1aa94ce3eb",
            "url": "https://rest_api_root/v1/project/projects/6375cc81e3004de4ea14aa52/work_item_types/630da48bc9443b1aa94ce3eb",
            "project": {
                "id": "6375cc81e3004de4ea14aa52",
                "url": "http://rest_api_root/v1/project/projects/6375cc81e3004de4ea14aa52",
                "identifier": "WTF",
                "name": "瀑布项目",
                "type": "waterfall",
                "is_archived": 0,
                "is_deleted": 0
            },
            "work_item_type": {
                "id": "630da48bc9443b1aa94ce3eb",
                "url": "https://rest_api_root/v1/project/projects/6375cc81e3004de4ea14aa52/work_item_types/630da48bc9443b1aa94ce3eb",
                "name": "任务",
                "group": "task"
            }
        },
        {
            "id": "630da48bc9443b1aa94ce3ec",
            "url": "https://rest_api_root/v1/project/projects/6375cc81e3004de4ea14aa52/work_item_types/630da48bc9443b1aa94ce3ec",
            "project": {
                "id": "6375cc81e3004de4ea14aa52",
                "url": "http://rest_api_root/v1/project/projects/6375cc81e3004de4ea14aa52",
                "identifier": "WTF",
                "name": "瀑布项目",
                "type": "waterfall",
                "is_archived": 0,
                "is_deleted": 0
            },
            "work_item_type": {
                "id": "630da48bc9443b1aa94ce3ec",
                "url": "https://rest_api_root/v1/project/projects/6375cc81e3004de4ea14aa52/work_item_types/630da48bc9443b1aa94ce3ec",
                "name": "缺陷",
                "group": "bug"
            }
        },
        {
            "id": "630da48bc9443b1aa94ce3ed",
            "url": "https://rest_api_root/v1/project/projects/6375cc81e3004de4ea14aa52/work_item_types/630da48bc9443b1aa94ce3ed",
            "project": {
                "id": "6375cc81e3004de4ea14aa52",
                "url": "http://rest_api_root/v1/project/projects/6375cc81e3004de4ea14aa52",
                "identifier": "WTF",
                "name": "瀑布项目",
                "type": "waterfall",
                "is_archived": 0,
                "is_deleted": 0
            },
            "work_item_type": {
                "id": "630da48bc9443b1aa94ce3ed",
                "url": "https://rest_api_root/v1/project/projects/6375cc81e3004de4ea14aa52/work_item_types/630da48bc9443b1aa94ce3ed",
                "name": "自定义",
                "group": "issue"
            }
        },
        {
            "id": "630da48bc9443b1aa94ce3ee",
            "url": "https://rest_api_root/v1/project/projects/6375cc81e3004de4ea14aa52/work_item_types/630da48bc9443b1aa94ce3ee",
            "project": {
                "id": "6375cc81e3004de4ea14aa52",
                "url": "http://rest_api_root/v1/project/projects/6375cc81e3004de4ea14aa52",
                "identifier": "WTF",
                "name": "瀑布项目",
                "type": "waterfall",
                "is_archived": 0,
                "is_deleted": 0
            },
            "work_item_type": {
                "id": "630da48bc9443b1aa94ce3ee",
                "url": "https://rest_api_root/v1/project/projects/6375cc81e3004de4ea14aa52/work_item_types/630da48bc9443b1aa94ce3ee",
                "name": "阶段",
                "group": "plan"
            }
        },
        {
            "id": "630da48bc9443b1aa94ce3ef",
            "url": "https://rest_api_root/v1/project/projects/6375cc81e3004de4ea14aa52/work_item_types/630da48bc9443b1aa94ce3ef",
            "project": {
                "id": "6375cc81e3004de4ea14aa52",
                "url": "http://rest_api_root/v1/project/projects/6375cc81e3004de4ea14aa52",
                "identifier": "WTF",
                "name": "瀑布项目",
                "type": "waterfall",
                "is_archived": 0,
                "is_deleted": 0
            },
            "work_item_type": {
                "id": "630da48bc9443b1aa94ce3ef",
                "url": "https://rest_api_root/v1/project/projects/6375cc81e3004de4ea14aa52/work_item_types/630da48bc9443b1aa94ce3ef",
                "name": "里程碑",
                "group": "plan"
            }
        }
    ]
}
```

---

## 环境

### 全量更新一个环境

**`PUT /v1/release/environments/{env_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `env_id` | String | 环境的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `name` | String | 环境的名称。同一团队内，环境的名称是唯一的。 |
| `html_url` *(可选)* | String | 环境的地址。如果为空，在PingCode中不显示相关跳转链接。 |

**响应示例：**

```json
{
    "id": "564587fe700d43b81b080123",
    "url": "https://rest_api_root/v1/release/environments/564587fe700d43b81b080123",
    "name": "Production",
    "html_url": "https://your-environment-url"
}
```

---

### 创建一个环境

**`POST /v1/release/environments`**

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `name` | String | 环境的名称。同一团队内，环境的名称是唯一的。 |
| `html_url` *(可选)* | String | 环境的地址。如果为空，在PingCode中不显示相关跳转链接。 |

**响应示例：**

```json
{
    "id": "564587fe700d43b81b080123",
    "url": "https://rest_api_root/v1/release/environments/564587fe700d43b81b080123",
    "name": "Production",
    "html_url": "https://your-environment-url"
}
```

---

### 删除一个环境

**`DELETE /v1/release/environments/{env_id}`**

删除环境之前，需要先删除与该环境关联的部署。

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `env_id` | String | 环境的id。 |

**响应示例：**

```json
{
    "id": "564587fe700d43b81b080123",
    "url": "https://rest_api_root/v1/release/environments/564587fe700d43b81b080123",
    "name": "Production",
    "html_url": "https://your-environment-url"
}
```

---

### 获取环境列表

**`GET /v1/release/environments`**

**查询参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `name` | String | 环境的名称。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "564587fe700d43b81b080123",
            "url": "https://rest_api_root/v1/release/environments/564587fe700d43b81b080123",
            "name": "Production",
            "html_url": "https://your-environment-url"
        }
    ]
}
```

---

### 部分更新一个环境

**`PATCH /v1/release/environments/{env_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `env_id` | String | 环境的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `name` *(可选)* | String | 环境的名称。同一团队内，环境的名称是唯一的。 |
| `html_url` *(可选)* | String | 环境的地址。如果为空，在PingCode中不显示相关跳转链接。 |

**响应示例：**

```json
{
    "id": "564587fe700d43b81b080123",
    "url": "https://rest_api_root/v1/release/environments/564587fe700d43b81b080123",
    "name": "Production",
    "html_url": "https://your-environment-url"
}
```

---

## 用例

### 创建一个用例

**`POST /v1/testhub/cases`**

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `test_library_id` | String | 测试库的id。 |
| `title` | String | 测试用例的标题。 |
| `suite_id` *(可选)* | String | 测试用例所属模块的id。 |
| `type_id` *(可选)* | String | 测试用例类型的id。 |
| `important_level_id` *(可选)* | String | 测试用例重要程度的id。 |
| `maintenance_id` *(可选)* | String | 测试用例维护人的id。 |
| `participant_ids` *(可选)* | String[] | 测试用例关注人的id列表。 |
| `properties` *(可选)* | Object | 测试用例属性的键值对集合，需要注意的是，当前测试用例对应的属性方案需要包含这些测试用例属性，例如属性方案中包含prop_a和prop_b。 |
| `properties.prop_a` *(可选)* | Object | 测试用例属性prop_a。 |
| `properties.prop_b` *(可选)* | Object | 测试用例属性prop_b。 |
| `description` *(可选)* | String | 测试用例的描述。 |
| `precondition` *(可选)* | String | 测试用例的前置条件。 |
| `steps` *(可选)* | Object[] | 测试用例的步骤列表。steps是整体更新的。 |
| `steps.step_id` *(可选)* | String | 测试用例步骤的id。 |
| `steps.description` *(可选)* | String | 测试用例步骤的描述。 |
| `steps.expected_value` *(可选)* | String | 测试用例步骤的期望值。 |
| `steps.is_group` *(可选)* | Number | 测试用例步骤类型是否为分组。 |
| `steps.group_id` *(可选)* | String | 测试用例步骤所属分组的id。group_id是分组类型步骤的step_id，分组类型的步骤不需要该参数。 |

**响应示例：**

```json
{
    "id": "5edca524cad2fa112b06305c",
    "url": "https://rest_api_root/v1/testhub/cases/5edca524cad2fa112b06305c",
    "library": {
        "id": "5eb623f6a70571487ea47000",
        "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000",
        "identifier": "CSK",
        "name": "测试库",
        "is_archived": 0,
        "is_deleted": 0
    },
    "identifier": "CSK-10",
    "title": "这是一个测试用例",
    "level": "P1",
    "short_id": "fdUw3C",
    "html_url": "https://yctech.pingcode.com/testhub/cases/fdUw3C",
    "important_level": {
        "id": "57a109b35ae8c20630fd7256",
        "url": "https://rest_api_root/v1/testhub/case_important_levels/57a109b35ae8c20630fd7256",
        "name": "P1"
    },
    "suite": {
        "id": "55714870a70ea4eb623f6700",
        "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/suites/55714870a70ea4eb623f6700",
        "name": "登录",
        "paths": "首页/账户"
    },
   "state": {
        "id": "686f62038668bbae4f4dd0c1",
        "url": "https://rest_api_root/v1/testhub/case_states/686f62038668bbae4f4dd0c1",
        "name": "设计",
        "type": "pending"
    },
    "type": {
        "id": "5cf189b35de9c20620ad7153",
        "url": "https://rest_api_root/v1/testhub/case_types/5cf189b35de9c20620ad7153",
        "name": "功能测试"
    },
    "maintenance": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "test_type": "automation",
    "description": "测试用例的备注",
    "precondition": "前置条件的描述信息",
    "properties": {
        "prop_a": "prop_a_value",
        "prop_b": "prop_b_value"
    },
    "estimated_workload": null,
    "remaining_workload": null,
    "steps": [
        {
            "step_id": "524cad5edb06305cca2fa112",
            "description": "UI测试",
            "expected_value": null,
            "is_group": 1,
            "group_id": null
        },
        {
            "step_id": "524cad5edb06305cca2fa113",
            "description": "在浏览器地址栏中输入https://pingcode.com",
            "expected_value": "成功访问PingCode官网",
            "is_group": 0,
            "group_id": "524cad5edb06305cca2fa112"
        }
    ],
    "participants": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "url": "https://rest_api_root/v1/participants/a0417f68e846aae315c85d24643678a9?principal_type=test_case&principal_id=5edca524cad2fa112b06305c",
            "type": "user",
            "user": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            }
        }
    ],
    "created_at": 1583290300,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "updated_at": 1583290300,
    "updated_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "is_archived": 0,
    "is_deleted": 0
}
```

---

### 删除一个用例

**`DELETE /v1/testhub/cases/{case_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `case_id` | String | 测试用例的id。 |

**响应示例：**

```json
{
    "id": "5edca524cad2fa112b06305d",
    "url": "https://rest_api_root/v1/testhub/cases/5edca524cad2fa112b06305d",
    "library": {
        "id": "5eb623f6a70571487ea47000",
        "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000",
        "identifier": "CSK",
        "name": "测试库",
        "is_archived": 0,
        "is_deleted": 0
    },
    "identifier": "CSK-1",
    "title": "这是一个测试用例",
    "level": "P1",
    "short_id": "fdUw3C",
    "html_url": "https://yctech.pingcode.com/testhub/cases/fdUw3C",
    "important_level": {
        "id": "57a109b35ae8c20630fd7256",
        "url": "https://rest_api_root/v1/testhub/case_important_levels/57a109b35ae8c20630fd7256",
        "name": "P1"
    },
    "suite": {
        "id": "55714870a70ea4eb623f6700",
        "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/suites/55714870a70ea4eb623f6700",
        "name": "登录",
        "paths": "首页/账户"
    },
    "state": {
        "id": "686f62038668bbae4f4dd0c1",
        "url": "https://rest_api_root/v1/testhub/case_states/686f62038668bbae4f4dd0c1",
        "name": "设计",
        "type": "pending"
    },
    "type": {
        "id": "5cf189b35de9c20620ad7153",
        "url": "https://rest_api_root/v1/testhub/case_types/5cf189b35de9c20620ad7153",
        "name": "功能测试"
    },
    "maintenance": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "test_type": "automation",
    "description": "测试用例的备注",
    "precondition": "前置条件的描述信息",
    "properties": {
        "prop_a": "prop_a_value",
        "prop_b": "prop_b_value"
    },
    "estimated_workload": 8,
    "remaining_workload": 5,
    "steps": [
        {
            "step_id": "524cad5edb06305cca2fa112",
            "description": "在浏览器地址栏中输入https://pingcode.com",
            "expected_value": "成功访问PingCode官网",
            "is_group": 0,
            "group_id": null
        }
    ],
    "participants": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "url": "https://rest_api_root/v1/participants/a0417f68e846aae315c85d24643678a9?principal_type=test_case&principal_id=5edca524cad2fa112b06305d",
            "type": "user",
            "user": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            }
        }
    ],
    "created_at": 1583290300,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "updated_at": 1583293300,
    "updated_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "is_archived": 0,
    "is_deleted": 1
}
```

---

### 批量创建测试用例关联工作项

**`POST /v1/testhub/cases/{case_id}/work_item_relations/bulk`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `case_id` | String | 测试用例id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `work_item_ids` | String[] | 关联的工作项id数组。支持的工作项类型包括bug和story。 |

**响应示例：**

```json
[
    {
        "state": "success",
        "relation": {
            "id": "fa1125cb06305edca524cad2",
            "url": "https://rest_api_root/v1/relations/fa1125cb06305edca524cad2",
            "principal_type": "test_case",
            "principal": {
                "id": "5edca524cad2fa112b06305c",
                "url": "https://rest_api_root/v1/testhub/cases/5edca524cad2fa112b06305c",
                "identifier": "CSK-10",
                "title": "这是一个测试用例",
                "level": "P1",
                "short_id": "fdUw3C",
                "html_url": "https://yctech.pingcode.com/testhub/cases/fdUw3C"
            },
            "target_type": "work_item",
            "target": {
                "id": "5edca524cad2fa1125cb0630",
                "url": "https://rest_api_root/v1/project/work_items/5edca524cad2fa1125cb0630",
                "identifier": "SCR-5",
                "title": "这是一个缺陷",
                "type": "bug",
                "start_at": 1583290309,
                "end_at": 1583290347,
                "parent_id": "5edca524cad2fa112b06105c",
                "short_id": "c9WqLmTO",
                "html_url": "https://yctech.pingcode.com/pjm/workitems/c9WqLmTO"
            }
        }
    }
]
```

---

### 批量创建用例

**`POST /v1/testhub/cases/bulk`**

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `cases` | Object[] | 创建单个测试用例必要参数的数组，数组内对象不能超过100个。 |
| `cases.test_library_id` | String | 测试用例所属测试库id。 |
| `cases.title` | String | 测试用例的标题，长度1～200有效字符。 |
| `cases.important_level_id` *(可选)* | String | 测试用例重要程度的id。 |
| `cases.maintenance_id` *(可选)* | String | 测试用例维护人的id。 |
| `cases.participant_ids` *(可选)* | String[] | 测试用例关注人的id列表。 |
| `cases.properties` *(可选)* | String | 测试用例属性的键值对集合。 |
| `cases.description` *(可选)* | String | 测试用例的描述。 |
| `cases.precondition` *(可选)* | String | 测试用例的前置条件。 |
| `cases.steps` *(可选)* | Object[] | 测试用例的步骤列表。 |
| `cases.steps.step_id` *(可选)* | String | 测试用例步骤的id。 |
| `cases.steps.description` *(可选)* | String | 测试用例步骤的描述。 |
| `cases.steps.expected_value` *(可选)* | String | 测试用例步骤的期望值。 |
| `cases.steps.is_group` *(可选)* | Number | 测试用例步骤类型是否为分组。 |
| `cases.steps.group_id` *(可选)* | String | 测试用例步骤所属分组的id。group_id是分组类型步骤的step_id，分组类型的步骤不需要该参数。 |

**响应示例：**

```json
[
    {
        "state": "success",
        "case": {
            "id": "5edca524cad2fa112b06305d",
            "url": "https://rest_api_root/v1/testhub/cases/5edca524cad2fa112b06305d",
            "library": {
                "id": "5eb623f6a70571487ea47000",
                "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000",
                "identifier": "CSK",
                "name": "测试库",
                "is_archived": 0,
                "is_deleted": 0
            },
            "identifier": "CSK-1",
            "title": "这是一个测试用例",
            "level": "P1",
            "short_id": "fdUw3C",
            "html_url": "https://yctech.pingcode.com/testhub/cases/fdUw3C",
            "important_level": {
                "id": "57a109b35ae8c20630fd7256",
                "url": "https://rest_api_root/v1/testhub/case_important_levels/57a109b35ae8c20630fd7256",
                "name": "P1"
            },
            "suite": {
                "id": "55714870a70ea4eb623f6700",
                "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/suites/55714870a70ea4eb623f6700",
                "name": "登录",
                "paths": "首页/账户"
            },
            "state": {
                "id": "686f62038668bbae4f4dd0c1",
                "url": "https://rest_api_root/v1/testhub/case_states/686f62038668bbae4f4dd0c1",
                "name": "设计",
                "type": "pending"
            },
            "type": {
                "id": "5cf189b35de9c20620ad7153",
                "url": "https://rest_api_root/v1/testhub/case_types/5cf189b35de9c20620ad7153",
                "name": "功能测试"
            },
            "maintenance": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "test_type": "automation",
            "description": "测试用例的备注",
            "precondition": "前置条件的描述信息",
            "properties": {
                "prop_a": "prop_a_value",
                "prop_b": "prop_b_value"
            },
            "estimated_workload": null,
            "remaining_workload": null,
            "steps": [
                {
                    "step_id": "524cad5edb06305cca2fa112",
                    "description": "UI测试",
                    "expected_value": null,
                    "is_group": 1,
                    "group_id": null
                },
                {
                    "step_id": "524cad5edb06305cca2fa113",
                    "description": "在浏览器地址栏中输入https://pingcode.com",
                    "expected_value": "成功访问PingCode官网",
                    "is_group": 0,
                    "group_id": "524cad5edb06305cca2fa112"
                }
            ],
            "participants": [
                {
                    "id": "a0417f68e846aae315c85d24643678a9",
                    "url": "https://rest_api_root/v1/participants/a0417f68e846aae315c85d24643678a9?principal_type=test_case&principal_id=5edca524cad2fa112b06305d",
                    "type": "user",
                    "user": {
                        "id": "a0417f68e846aae315c85d24643678a9",
                        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                        "name": "john",
                        "display_name": "John",
                        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
                    }
                }
            ],
            "created_at": 1583290300,
            "created_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "updated_at": 1583293300,
            "updated_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "is_archived": 0,
            "is_deleted": 0
        }
    }
]
```

---

### 批量删除测试用例关联工作项

**`DELETE /v1/testhub/cases/{case_id}/work_item_relations/bulk`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `case_id` | String | 测试用例id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `work_item_ids` | String[] | 关联的工作项id数组。支持的工作项类型包括bug和story。 |

**响应示例：**

```json
[
    {
        "state": "success",
        "relation": {
            "id": "fa1125cb06305edca524cad2",
            "url": "https://rest_api_root/v1/relations/fa1125cb06305edca524cad2",
            "principal_type": "test_case",
            "principal": {
                "id": "5edca524cad2fa112b06305c",
                "url": "https://rest_api_root/v1/testhub/cases/5edca524cad2fa112b06305c",
                "identifier": "CSK-10",
                "title": "这是一个测试用例",
                "level": "P1",
                "short_id": "fdUw3C",
                "html_url": "https://yctech.pingcode.com/testhub/cases/fdUw3C"
            },
            "target_type": "work_item",
            "target": {
                "id": "5edca524cad2fa1125cb0630",
                "url": "https://rest_api_root/v1/project/work_items/5edca524cad2fa1125cb0630",
                "identifier": "SCR-5",
                "title": "这是一个缺陷",
                "type": "bug",
                "start_at": 1583290309,
                "end_at": 1583290347,
                "parent_id": "5edca524cad2fa112b06105c",
                "short_id": "c9WqLmTO",
                "html_url": "https://yctech.pingcode.com/pjm/workitems/c9WqLmTO",
                "properties": {
                   "prop_a": "prop_a_value",
                   "prop_b": "prop_b_value"
                }
            }
        }
    }
]
```

---

### 批量部分更新用例

**`PATCH /v1/testhub/cases/bulk`**

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `cases` | Object[] | 部分更新测试用例的数组。 |
| `cases.case_id` | String | 测试用例的id。 |
| `cases.state_id` *(可选)* | String | 测试用例状态的id。 |
| `cases.type_id` *(可选)* | String | 测试用例类型的id。 |
| `cases.title` *(可选)* | String | 测试用例的标题。 |
| `cases.important_level_id` *(可选)* | String | 测试用例重要程度的id。 |
| `cases.maintenance_id` *(可选)* | String | 测试用例维护人的id。 |
| `cases.properties` *(可选)* | Object[] | 测试用例属性的键值对集合，property中包含propertyKey、propertyValue和propertyType三个字段。需要注意的是，当前测试用例对应的属性方案需要包含这些测试用例属性。 |
| `cases.properties.prop_a` *(可选)* | Object | 测试用例属性的自定义属性prop_a。 |
| `cases.properties.prop_b` *(可选)* | Object | 测试用例属性的自定义属性prop_b。 |
| `cases.description` *(可选)* | String | 测试用例的备注。 |
| `cases.precondition` *(可选)* | String | 测试用例的前置条件。 |
| `cases.steps` *(可选)* | Object[] | 测试用例的步骤列表。steps是整体更新的。 |
| `cases.steps.step_id` *(可选)* | String | 测试用例的步骤的id。如果step中不包含用于确定唯一性的step_id，那么这个step将被视为新建，并为之创建新的step_id。 |
| `cases.steps.description` *(可选)* | String | 测试用例的步骤的描述。 |
| `cases.steps.expected_value` *(可选)* | String | 测试用例的步骤的期望值。 |
| `cases.steps.is_group` *(可选)* | Number | 测试用例步骤类型是否为分组。 |
| `cases.steps.group_id` *(可选)* | String | 测试用例步骤所属分组的id。group_id是分组类型步骤的step_id，分组类型的步骤不需要该参数。 |

**响应示例：**

```json
[
    {
        "state": "success",
        "case": {
            "id": "5edca524cad2fa112b06305c",
            "url": "https://rest_api_root/v1/testhub/cases/5edca524cad2fa112b06305c",
            "library": {
                "id": "5eb623f6a70571487ea47000",
                "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000",
                "identifier": "CSK",
                "name": "测试库",
                "is_archived": 0,
                "is_deleted": 0
            },
            "identifier": "CSK-10",
            "title": "这是一个测试用例",
            "level": "P1",
            "short_id": "fdUw3C",
            "html_url": "https://yctech.pingcode.com/testhub/cases/fdUw3C",
            "important_level": {
                "id": "57a109b35ae8c20630fd7256",
                "url": "https://rest_api_root/v1/testhub/case_important_levels/57a109b35ae8c20630fd7256",
                "name": "P1"
            },
            "suite": {
                "id": "55714870a70ea4eb623f6700",
                "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/suites/55714870a70ea4eb623f6700",
                "name": "登录",
                "paths": "首页/账户"
            },
            "state": {
                 "id": "686f62038668bbae4f4dd0c1",
                 "url": "https://rest_api_root/v1/testhub/case_states/686f62038668bbae4f4dd0c1",
                 "name": "设计",
                 "type": "pending"
            },
            "type": {
                "id": "5cf189b35de9c20620ad7153",
                "url": "https://rest_api_root/v1/testhub/case_types/5cf189b35de9c20620ad7153",
                "name": "功能测试"
            },
            "maintenance": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "test_type": "automation",
            "description": "测试用例的备注",
            "precondition": "前置条件的描述信息",
            "properties": {
                "prop_a": "prop_a_value",
                "prop_b": "prop_b_value"
            },
            "estimated_workload": 8,
            "remaining_workload": 5,
            "steps": [
                {
                    "step_id": "524cad5edb06305cca2fa112",
                    "description": "UI测试",
                    "expected_value": null,
                    "is_group": 1,
                    "group_id": null
                },
                {
                    "step_id": "524cad5edb06305cca2fa113",
                    "description": "点击下一页按钮",
                    "expected_value": "成功跳转至下一页",
                    "is_group": 0,
                    "group_id": "524cad5edb06305cca2fa112"
                },
                {
                    "step_id": "524cad5edb06305cca2fa114",
                    "description": "点击最后一页按钮",
                    "expected_value": "成功跳转至最后一页",
                    "is_group": 0,
                    "group_id": "524cad5edb06305cca2fa112"
                }
            ],
            "participants": [
                {
                    "id": "a0417f68e846aae315c85d24643678a9",
                    "url": "https://rest_api_root/v1/participants/a0417f68e846aae315c85d24643678a9?principal_type=test_case&principal_id=5edca524cad2fa112b06305c",
                    "type": "user",
                    "user": {
                        "id": "a0417f68e846aae315c85d24643678a9",
                        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                        "name": "john",
                        "display_name": "John",
                        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
                    }
                }
            ],
            "created_at": 1583290300,
            "created_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "updated_at": 15832903300,
            "updated_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "is_archived": 0,
            "is_deleted": 0
        }
    }
]
```

---

### 获取用例列表

**`GET /v1/testhub/cases`**

**查询参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `library_id` *(可选)* | String | 测试库的id。 |
| `maintenance_id` *(可选)* | String | 维护人的id。 |
| `state_ids` *(可选)* | String | 状态的id，使用','分割，最多只能20个。 |
| `important_level_ids` *(可选)* | String | 重要程度的id，使用','分割，最多只能20个。 |
| `tag_ids` *(可选)* | String | 标签的id，使用','分割，最多只能20个。 |
| `created_between` *(可选)* | String | 创建时间介于的时间范围，通过','分割起始时间。 |
| `updated_between` *(可选)* | String | 更新时间介于的时间范围，通过','分割起始时间。 |
| `include_public_image_token` *(可选)* | String | 包含获取图片资源token的属性。使用','分割，最多只能20个，支持description和自定义多行文本类型的属性。参数示例description,properties.prop_b。 |
| `include_deleted` *(可选)* | Boolean | 是否查询已删除的用例。该值默认为false。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "5edca524cad2fa112b06305c",
            "url": "https://rest_api_root/v1/testhub/cases/5edca524cad2fa112b06305c",
            "library": {
                "id": "5eb623f6a70571487ea47000",
                "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000",
                "identifier": "CSK",
                "name": "测试库",
                "is_archived": 0,
                "is_deleted": 0
            },
            "identifier": "CSK-10",
            "title": "这是一个测试用例",
            "level": "P1",
            "short_id": "fdUw3C",
            "html_url": "https://yctech.pingcode.com/testhub/cases/fdUw3C",
            "important_level": {
                "id": "57a109b35ae8c20630fd7256",
                "url": "https://rest_api_root/v1/testhub/case_important_levels/57a109b35ae8c20630fd7256",
                "name": "P1"
            },
            "suite": {
                "id": "55714870a70ea4eb623f6700",
                "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/suites/55714870a70ea4eb623f6700",
                "name": "登录",
                "paths": "首页/账户"
            },
            "state": {
                "id": "686f62038668bbae4f4dd0c1",
                "url": "https://rest_api_root/v1/testhub/case_states/686f62038668bbae4f4dd0c1",
                "name": "设计",
                "type": "pending"
             },
            "type": {
                "id": "5cf189b35de9c20620ad7153",
                "url": "https://rest_api_root/v1/testhub/case_types/5cf189b35de9c20620ad7153",
                "name": "功能测试"
            },
            "maintenance": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "test_type": "automation",
            "description": "测试用例的备注",
            "precondition": "前置条件的描述信息",
            "properties": {
                "prop_a": "prop_a_value",
                "prop_b": "prop_b_value"
            },
            "estimated_workload": 8,
            "remaining_workload": 5,
            "steps": [
                {
                    "step_id": "524cad5edb06305cca2fa100",
                    "description": "UI测试",
                    "expected_value": null,
                    "is_group": 1,
                    "group_id": null
                },
                {
                    "step_id": "524cad5edb06305cca2fa112",
                    "description": "在浏览器地址栏中输入https://pingcode.com",
                    "expected_value": "成功访问PingCode官网",
                    "is_group": 0,
                    "group_id": "524cad5edb06305cca2fa100"
                },
                {
                    "step_id": "524cad5edb06305cca2fa113",
                    "description": "点击下一页按钮",
                    "expected_value": "成功跳转至下一页",
                    "is_group": 0,
                    "group_id": "524cad5edb06305cca2fa100"
                }
            ],
            "participants": [
                {
                    "id": "a0417f68e846aae315c85d24643678a9",
                    "url": "https://rest_api_root/v1/participants/a0417f68e846aae315c85d24643678a9?principal_type=test_case&principal_id=5edca524cad2fa112b06305c",
                    "type": "user",
                    "user": {
                        "id": "a0417f68e846aae315c85d24643678a9",
                        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                        "name": "john",
                        "display_name": "John",
                        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
                    }
                }
            ],
            "public_image_token": "IcF1VmJFF-UIi9lMU18m1NFFAruWXYx0ZAm90pJ2LGk",
            "created_at": 1583290300,
            "created_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "updated_at": 1583293300,
            "updated_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "is_archived": 0,
            "is_deleted": 0
        }
    ]
}
```

---

### 获取用例属性列表

**`GET /v1/testhub/case/properties?library_id={library_id}`**

**查询参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `library_id` | String | 测试库的id。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 2,
    "values": [
        {
            "id": "environment",
            "url": "https://rest_api_root/v1/testhub/case_properties/environment",
            "name": "重现环境",
            "type": "select",
            "options": [
                {
                    "_id": "5efb1859110533727a82c603",
                    "text": "测试"
                },
                {
                    "_id": "5efb1859110533727a82c604",
                    "text": "生产"
                }
            ]
        },
        {
            "id": "estimated_workload",
            "url": "https://rest_api_root/v1/testhub/case_properties/estimated_workload",
            "name": "预估工时",
            "type": "number",
            "options": null
        }
    ]
}
```

---

### 获取用例模块列表

**`GET /v1/testhub/case/suites?library_id={library_id}`**

**查询参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `library_id` | String | 测试库的id。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "55714870a70ea4eb623f6700",
            "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/suites/55714870a70ea4eb623f6700",
            "name": "登录",
            "paths": "首页/用户"
        }
    ]
}
```

---

### 获取用例状态列表

**`GET /v1/testhub/case/states?library_id={library_id}`**

**查询参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `library_id` | String | 测试库的id。 |

**响应示例：**

```json
{
    "page_index": 0,
    "page_size": 30,
    "total": 3,
    "values": [
        {
            "id": "686f62038668bbae4f4dd0c1",
            "url": "http://rest_api_root/v1/testhub/case_states/686f62038668bbae4f4dd0c1",
            "name": "设计",
            "type": "pending"
        },
        {
            "id": "686f62038668bbae4f4dd0c2",
            "url": "http://rest_api_root/v1/testhub/case_states/686f62038668bbae4f4dd0c2",
            "name": "就绪",
            "type": "completed"
        },
        {
            "id": "686f62038668bbae4f4dd0c3",
            "url": "http://rest_api_root/v1/testhub/case_states/686f62038668bbae4f4dd0c3",
            "name": "废弃",
            "type": "closed"
        }
    ]
}
```

---

### 获取用例的执行历史列表

**`GET /v1/testhub/cases/{case_id}/histories`**

获取测试用例所有执行用例的最后一次执行结果。

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `case_id` | String | 测试用例的id。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "65115f0939286e26e05a66db",
            "url": "https://rest_api_root/v1/testhub/runs/547000eb6a70571487623fea/histories/65115f0939286e26e05a66db",
            "run": {
                "id": "547000eb6a70571487623fea",
                "url": "https://rest_api_root/v1/testhub/runs/547000eb6a70571487623fea",
                "status": "pass",
                "short_id": "Aq1u38yX",
                "html_url": "https://yctech.pingcode.com/testhub/runs/Aq1u38yX"
            },
            "library": {
                "id": "5eb623f6a70571487ea47000",
                "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000",
                "identifier": "CSK",
                "name": "测试库",
                "is_archived": 0,
                "is_deleted": 0
            },
            "plan": {
                "id": "5eb6a70571487623fea47000",
                "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/plans/5eb6a70571487623fea47000",
                "name": "测试计划",
                "status": "in_progress",
                "start_at": 1589791860,
                "end_at": 1589791870,
                "short_id": "7nNkViOD",
                "html_url": "https://yctech.pingcode.com/testhub/libraries/CSK/plans/7nNkViOD/runs"
            },
            "case": {
                "id": "5edca524cad2fa112b06305c",
                "url": "https://rest_api_root/v1/testhub/cases/5edca524cad2fa112b06305c",
                "identifier": "CSK-10",
                "title": "这是一个测试用例",
                "level": "P1",
                "short_id": "fdUw3C",
                "html_url": "https://yctech.pingcode.com/testhub/cases/fdUw3C",
                "test_type": "automation",
                "properties": {
                    "prop_a": "prop_a_value",
                    "prop_b": "prop_b_value"
                }
            },
            "status": "pass",
            "executed_at": 1583290300,
            "executed_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "steps": [
                {
                    "step_id": "524cad5edb06305cca2fa112",
                    "status": "pass",
                    "actual_value": null
                }
            ]
        }
    ]
}
```

---

### 获取用例类型列表

**`GET /v1/testhub/case/types?library_id={library_id}`**

**查询参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `library_id` | String | 测试库的id。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "5cf189b35de9c20620ad7153",
            "url": "https://rest_api_root/v1/testhub/case_types/5cf189b35de9c20620ad7153",
            "name": "功能测试"
        }
    ]
}
```

---

### 部分更新一个用例

**`PATCH /v1/testhub/cases/{case_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `case_id` | String | 测试用例的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `suite_id` *(可选)* | String | 测试用例所属模块的id。 |
| `state_id` *(可选)* | String | 测试用例状态的id。 |
| `type_id` *(可选)* | String | 测试用例类型的id。 |
| `title` *(可选)* | String | 测试用例的标题。 |
| `important_level_id` *(可选)* | String | 测试用例重要程度的id。 |
| `maintenance_id` *(可选)* | String | 测试用例维护人的id。 |
| `properties` *(可选)* | Object | 测试用例属性的键值对集合。需要注意的是，当前测试用例对应的属性方案需要包含这些测试用例属性。 |
| `properties.prop_a` *(可选)* | Object | 测试用例属性prop_a。 |
| `properties.prop_b` *(可选)* | Object | 测试用例属性prop_b。 |
| `description` *(可选)* | String | 测试用例的备注。 |
| `precondition` *(可选)* | String | 测试用例的前置条件。 |
| `steps` *(可选)* | Object[] | 测试用例的步骤列表。steps是整体更新的。 |
| `steps.step_id` *(可选)* | String | 测试用例的步骤的id。如果step中不包含用于确定唯一性的step_id，那么这个step将被视为新建，并为之创建新的step_id。 |
| `steps.description` *(可选)* | String | 测试用例的步骤的描述。 |
| `steps.expected_value` *(可选)* | String | 测试用例的步骤的期望值。 |
| `steps.is_group` *(可选)* | Number | 测试用例步骤类型是否为分组。 |
| `steps.group_id` *(可选)* | String | 测试用例步骤所属分组的id。group_id是分组类型步骤的step_id，分组类型的步骤不需要该参数。 |

**响应示例：**

```json
{
    "id": "5edca524cad2fa112b06305c",
    "url": "https://rest_api_root/v1/testhub/cases/5edca524cad2fa112b06305c",
    "library": {
        "id": "5eb623f6a70571487ea47000",
        "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000",
        "identifier": "CSK",
        "name": "测试库",
        "is_archived": 0,
        "is_deleted": 0
    },
    "identifier": "CSK-10",
    "title": "这是一个测试用例",
    "level": "P1",
    "short_id": "fdUw3C",
    "html_url": "https://yctech.pingcode.com/testhub/cases/fdUw3C",
    "important_level": {
        "id": "57a109b35ae8c20630fd7256",
        "url": "https://rest_api_root/v1/testhub/case_important_levels/57a109b35ae8c20630fd7256",
        "name": "P1"
    },
    "suite": {
        "id": "55714870a70ea4eb623f6700",
        "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/suites/55714870a70ea4eb623f6700",
        "name": "登录",
        "paths": "首页/账户"
    },
   "state": {
        "id": "686f62038668bbae4f4dd0c1",
        "url": "https://rest_api_root/v1/testhub/case_states/686f62038668bbae4f4dd0c1",
        "name": "设计",
        "type": "pending"
     },
    "type": {
        "id": "5cf189b35de9c20620ad7153",
        "url": "https://rest_api_root/v1/testhub/case_types/5cf189b35de9c20620ad7153",
        "name": "功能测试"
    },
    "maintenance": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "test_type": "automation",
    "description": "测试用例的备注",
    "precondition": "前置条件的描述信息",
    "properties": {
        "prop_a": "prop_a_value",
        "prop_b": "prop_b_value"
    },
    "estimated_workload": 8,
    "remaining_workload": 5,
    "steps": [
        {
            "step_id": "524cad5edb06305cca2fa112",
            "description": "UI测试",
            "expected_value": null,
            "is_group": 1,
            "group_id": null
        },
        {
            "step_id": "524cad5edb06305cca2fa113",
            "description": "点击下一页按钮",
            "expected_value": "成功跳转至下一页",
            "is_group": 0,
            "group_id": "524cad5edb06305cca2fa112"
        },
        {
            "step_id": "524cad5edb06305cca2fa114",
            "description": "点击最后一页按钮",
            "expected_value": "成功跳转至最后一页",
            "is_group": 0,
            "group_id": "524cad5edb06305cca2fa112"
        }
    ],
    "participants": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "url": "https://rest_api_root/v1/participants/a0417f68e846aae315c85d24643678a9?principal_type=test_case&principal_id=5edca524cad2fa112b06305c",
            "type": "user",
            "user": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            }
        }
    ],
    "created_at": 1583290300,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "updated_at": 1583293300,
    "updated_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "is_archived": 0,
    "is_deleted": 0
}
```

---

## 用例配置

### 创建一个用例属性

**`POST /v1/testhub/case_properties`**

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `name` | String | 用例属性的名称。在一个企业中这个名称是唯一的。 |
| `type` | String | 用例属性的类型。 |
| `options` *(可选)* | Object[] | 选择项列表。当用例属性类型为select,multi_select,cascade_select,cascade_multi_select时可选填此项。 |
| `options._id` *(可选)* | String | 选择项id。该值在选择项中是唯一的。 |
| `options.text` | String | 选择项内容。该值在选择项中是唯一的。 |
| `options.parent_id` *(可选)* | String | 选择项父级id。当属性类型为cascade_select,cascade_multi_select时，parent_id参数有效，用于构建级联类型选择项之间的父子关系，最多存在4级。 |

**响应示例：**

```json
{
    "id": "severity",
    "url": "https://rest_api_root/v1/testhub/case_properties/severity",
    "name": "严重程度",
    "type": "select",
    "options": [
        {
            "_id": "5efb1859110533727a82c603",
            "text": "严重"
        },
        {
            "_id": "5efb1859110533727a82c604",
            "text": "一般"
        }
    ],
    "is_removable": true,
    "is_name_editable": true,
    "is_options_editable": true
}
```

---

### 向属性方案中添加一个用例属性

**`POST /v1/testhub/case_property_plans/{property_plan_id}/case_properties`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `property_plan_id` | String | 测试用例属性方案的id。 |

**请求参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `property_id` | String | 测试用例属性的id。 |

**响应示例：**

```json
{
    "id": "environment",
    "url": "https://rest_api_root/v1/testhub/case_property_plans/5f8a21f18ef715265de90c21/case_properties/environment",
    "property_plan": {
        "id": "5f8a21f18ef715265de90c21",
        "url": "https://rest_api_root/v1/testhub/case_property_plans/5f8a21f18ef715265de90c21",
        "category": "library",
        "host": "case"
    },
    "property": {
        "id": "environment",
        "url": "https://rest_api_root/v1/testhub/case_properties/environment",
        "name": "重现环境",
        "type": "select",
        "options": [
            {
                "_id": "5efb1859110533727a82c603",
                "text": "test"
            },
            {
                "_id": "5efb1859110533727a82c604",
                "text": "beta"
            },
            {
                "_id": "5efb1859110533727a82c605",
                "text": "rc"
            }
        ]
    }
}
```

---

### 在属性方案中移除一个用例属性

**`DELETE /v1/testhub/case_property_plans/{property_plan_id}/case_properties/{property_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `property_plan_id` | String | 测试用例属性方案的id。 |
| `property_id` | String | 测试用例属性的id。 |

**响应示例：**

```json
{
    "id": "environment",
    "url": "https://rest_api_root/v1/testhub/case_property_plans/5f8a21f18ef715265de90c21/case_properties/environment",
    "property_plan": {
        "id": "5f8a21f18ef715265de90c21",
        "url": "https://rest_api_root/v1/testhub/case_property_plans/5f8a21f18ef715265de90c21",
        "category": "library",
        "host": "case"
    },
    "property": {
        "id": "environment",
        "url": "https://rest_api_root/v1/testhub/case_properties/environment",
        "name": "重现环境",
        "type": "select",
        "options": [
            {
                "_id": "5efb1859110533727a82c603",
                "text": "test"
            },
            {
                "_id": "5efb1859110533727a82c604",
                "text": "beta"
            },
            {
                "_id": "5efb1859110533727a82c605",
                "text": "rc"
            }
        ]
    }
}
```

---

### 获取全部用例属性列表

**`GET /v1/testhub/case_properties`**

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 2,
    "values": [
        {
            "id": "environment",
            "url": "https://rest_api_root/v1/testhub/case_properties/environment",
            "name": "重现环境",
            "type": "select",
            "options": [
                {
                    "_id": "5efb1859110533727a82c603",
                    "text": "测试"
                },
                {
                    "_id": "5efb1859110533727a82c604",
                    "text": "生产"
                }
            ],
            "is_removable": false,
            "is_name_editable": false,
            "is_options_editable": false
        },
        {
            "id": "estimated_workload",
            "url": "https://rest_api_root/v1/testhub/case_properties/estimated_workload",
            "name": "预估工时",
            "type": "number",
            "options": null,
            "is_removable": false,
            "is_name_editable": false,
            "is_options_editable": false
        }
    ]
}
```

---

### 获取全部用例状态列表

**`GET /v1/testhub/case_states`**

**响应示例：**

```json
{
    "page_index": 0,
    "page_size": 30,
    "total": 3,
    "values": [
        {
            "id": "686f62038668bbae4f4dd0c1",
            "url": "http://rest_api_root/v1/testhub/case_states/686f62038668bbae4f4dd0c1",
            "name": "设计",
            "type": "pending"
        },
        {
            "id": "686f62038668bbae4f4dd0c2",
            "url": "http://rest_api_root/v1/testhub/case_states/686f62038668bbae4f4dd0c2",
            "name": "就绪",
            "type": "completed"
        },
        {
            "id": "686f62038668bbae4f4dd0c3",
            "url": "http://rest_api_root/v1/testhub/case_states/686f62038668bbae4f4dd0c3",
            "name": "废弃",
            "type": "closed"
        }
    ]
}
```

---

### 获取全部用例类型列表

**`GET /v1/testhub/case_types`**

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "5cf189b35de9c20620ad7153",
            "url": "https://rest_api_root/v1/testhub/case_types/5cf189b35de9c20620ad7153",
            "name": "功能测试"
        }
    ]
}
```

---

### 获取全部重要程度列表

**`GET /v1/testhub/case_important_levels`**

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "57a109b35ae8c20630fd7256",
            "url": "https://rest_api_root/v1/testhub/case_important_levels/57a109b35ae8c20630fd7256",
            "name": "P11",
            "color": "#56ABFB"
        }
    ]
}
```

---

### 获取属性方案中的用例属性列表

**`GET /v1/testhub/case_property_plans/{property_plan_id}/case_properties`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `property_plan_id` | String | 测试用例属性方案的id。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 2,
    "values": [
        {
            "id": "environment",
            "url": "https://rest_api_root/v1/testhub/case_property_plans/5f8a21f18ef715265de90c21/case_properties/environment",
            "property_plan": {
                "id": "5f8a21f18ef715265de90c21",
                "url": "https://rest_api_root/v1/testhub/case_property_plans/5f8a21f18ef715265de90c21",
                "category": "library",
                "host": "case"
            },
            "property": {
                "id": "environment",
                "url": "https://rest_api_root/v1/testhub/case_properties/environment",
                "name": "重现环境",
                "type": "select",
                "options": [
                    {
                        "_id": "5efb1859110533727a82c603",
                        "text": "test"
                    },
                    {
                        "_id": "5efb1859110533727a82c604",
                        "text": "beta"
                    },
                    {
                        "_id": "5efb1859110533727a82c605",
                        "text": "rc"
                    }
                ]
            }
        },
        {
            "id": "estimated_workload",
            "url": "https://rest_api_root/v1/testhub/property_plans/5f8a21f18ef715265de90c21/properties/estimated_workload",
            "property_plan": {
                "id": "5f8a21f18ef715265de90c21",
                "url": "https://rest_api_root/v1/testhub/property_plans/5f8a21f18ef715265de90c21",
                "category": "library",
                "host": "case"
            },
            "property": {
                "id": "estimated_workload",
                "url": "https://rest_api_root/v1/testhub/case_properties/estimated_workload",
                "name": "预估工时",
                "type": "number",
                "options": null
            }
        }
    ]
}
```

---

### 获取用例属性方案列表

**`GET /v1/testhub/case_property_plans`**

**查询参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `library_id` *(可选)* | String | 测试库的id。查询开启本地配置的方案时传入。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 2,
    "values": [
        {
            "id": "5f8a21f18ef715265de90c21",
            "url": "https://rest_api_root/v1/testhub/case_property_plans/5f8a21f18ef715265de90c21",
            "category": "library",
            "host": "case",
            "library": null
        },
        {
            "id": "5f8a21f18ef715265de90c22",
            "url": "https://rest_api_root/v1/testhub/case_property_plans/5f8a21f18ef715265de90c22",
            "category": "library",
            "host": "case",
            "library": {
                "id": "5eb623f6a70571487ea47000",
                "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000",
                "identifier": "CSK",
                "name": "测试库",
                "is_archived": 0,
                "is_deleted": 0
            }
        }
    ]
}
```

---

### 部分更新一个用例属性

**`PATCH /v1/testhub/case_properties/{property_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `property_id` | String | 用例属性的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `name` *(可选)* | String | 用例属性的名称。在一个企业中这个名称是唯一的。 |
| `options` *(可选)* | Object[] | 选择项列表。options是整体更新的。当用例属性类型为select,multi_select,cascade_select,cascade_multi_select时可选填此项。 |
| `options._id` *(可选)* | String | 选择项id。该值在选择项中是唯一的。 |
| `options.text` | String | 选择项内容。该值在选择项中是唯一的。 |
| `options.parent_id` *(可选)* | String | 选择项父级id。当属性类型为cascade_select,cascade_multi_select时，parent_id参数有效，用于构建级联类型选择项之间的父子关系，最多存在4级。 |

**响应示例：**

```json
{
    "id": "severity-update",
    "url": "https://rest_api_root/v1/testhub/case_properties/severity",
    "name": "严重程度-update",
    "type": "select",
    "options": [
        {
            "_id": "5efb1859110533727a82c603",
            "text": "严重-update"
        },
        {
            "_id": "5efb1859110533727a82c624",
            "text": "一般"
        }
    ],
    "is_removable": true,
    "is_name_editable": true,
    "is_options_editable": true
}
```

---

## 知识管理

### 空间

用于查看和管理空间相关的基本信息。 资源地址：{GET} https://rest_api_root/v1/wiki/spaces/{space_id}

**全量数据示例：**

```json
{
    "id": "642fd641209b56920a6c6e5e",
    "url": "https://rest_api_root/v1/wiki/spaces/642fd641209b56920a6c6e5e",
    "identifier": "DEMO",
    "name": "示例空间",
    "scope_type": "user_group",
    "scope_id": "63c8fb32729dee3334d96af7",
    "visibility": "private",
    "color": "#FB7894",
    "description": "示例空间描述",
    "members": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "url": "https://rest_api_root/v1/wiki/spaces/642fd641209b56920a6c6e5e/members/a0417f68e846aae315c85d24643678a9",
            "type": "user",
            "user": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            }
        },
        {
            "id": "63c8fb32729dee3334d96af7",
            "url": "https://rest_api_root/v1/wiki/spaces/642fd641209b56920a6c6e5e/members/63c8fb32729dee3334d96af7",
            "type": "user_group",
            "user_group": {
                "id": "63c8fb32729dee3334d96af7",
                "url": "https://rest_api_root/v1/directory/groups/63c8fb32729dee3334d96af7",
                "name": "Open Team"
            }
        }
    ],
    "created_at": 1583290300,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "updated_at": 1583290400,
    "updated_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "is_archived": 0,
    "is_deleted": 0
}
```

**引用数据示例：**

```json
{
    "id": "642fd641209b56920a6c6e5e",
    "url": "https://rest_api_root/v1/wiki/spaces/642fd641209b56920a6c6e5e",
    "identifier": "DEMO",
    "name": "示例空间",
    "is_archived": 0,
    "is_deleted": 0
}
```

---

### 页面

用于查看和管理页面相关的基本信息。企业令牌只能管理非个人空间下的页面。 资源地址：{GET} https://rest_api_root/v1/wiki/pages/{page_id}

**全量数据示例：**

```json
{
    "id": "63e1bf51760505c8795ebccc",
    "url": "https://rest_api_root/v1/wiki/pages/63e1bf51760505c8795ebccc",
    "space": {
        "id": "63e1bf51760505c8795ebcc8",
        "url": "https://rest_api_root/v1/wiki/spaces/63e1bf51760505c8795ebcc8",
        "name": "示例空间",
        "identifier": "DEMO",
        "is_archived": 0,
        "is_deleted": 0
    },
    "name": "示例页面",
    "type": "document",
    "short_id": "5-x6NN",
    "html_url": "https://yctech.pingcode.com/wiki/pages/5-x6NN",
    "parent": {
        "id": "63e1bf51760505c8795ebcce",
        "url": "https://rest_api_root/v1/wiki/pages/63e1bf51760505c8795ebcce",
        "name": "模板",
        "type": "document",
        "short_id": "7nNkViOD",
        "html_url": "https://yctech.pingcode.com/wiki/pages/5-x6NN"
    },
    "icon": "file-fill",
    "readings": 10,
    "published_at": 1694065129,
    "published_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "participants": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "url": "https://rest_api_root/v1/participants/a0417f68e846aae315c85d24643678a9?principal_type=page&principal_id=63e1bf51760505c8795ebccc",
            "type": "user",
            "user": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            }
        }
    ],
    "created_at": 1675738962,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "updated_at": 1694065129,
    "updated_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "is_locked": 0,
    "is_archived": 0,
    "is_deleted": 0
}
```

**引用数据示例：**

```json
{
    "id": "63e1bf51760505c8795ebccc",
    "url": "https://rest_api_root/v1/wiki/pages/63e1bf51760505c8795ebccc",
    "name": "示例页面",
    "type": "document",
    "short_id": "5-x6NN",
    "html_url": "https://yctech.pingcode.com/wiki/pages/5-x6NN"
}
```

---

## 空间

### 创建一个空间

**`POST /v1/wiki/spaces`**

企业令牌不能创建scope_type为user类型的空间。

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `name` | String | 空间的名称（不超过32个字符）。 |
| `identifier` | String | 空间的标识。在一个企业中这个标识是唯一的。产品的标识由大写英文字母/数字/下划线/连接线组成（不超过15个字符）。 |
| `scope_type` | String | 空间的所属类型。允许值分别表示企业可见、团队可见和用户可见。 |
| `scope_id` *(可选)* | String | 空间的所属id。当scope_type为user_group时，该字段必填，且表示团队id；当scope_type为其他值时，该字段无效。 |
| `visibility` *(可选)* | String | 空间可见性。允许值分别表示组织全部成员可见和仅空间成员可见。 |
| `description` *(可选)* | String | 空间的描述。 |
| `members` *(可选)* | Object[] | 空间成员的列表。 |
| `members.id` | String | 企业成员或团队的id。 |
| `members.type` | String | 空间成员类型。 |

**响应示例：**

```json
{
    "id": "642fd641209b56920a6c6e5f",
    "url": "https://rest_api_root/v1/wiki/spaces/642fd641209b56920a6c6e5f",
    "identifier": "GROUP",
    "name": "团队空间",
    "scope_type": "user_group",
    "scope_id": "63c8fb32729dee3334d96af7",
    "visibility": "private",
    "color": "#FB7894",
    "description": "团队空间所属一个团队，只能添加所属团队内的成员。",
    "members": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "url": "https://rest_api_root/v1/wiki/spaces/642fd641209b56920a6c6e5f/members/a0417f68e846aae315c85d24643678a9",
            "type": "user",
            "user": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            }
        },
        {
            "id": "63c8fb32729dee3334d96af7",
            "url": "https://rest_api_root/v1/wiki/spaces/642fd641209b56920a6c6e5f/members/63c8fb32729dee3334d96af7",
            "type": "user_group",
            "user_group": {
                "id": "63c8fb32729dee3334d96af7",
                "url": "https://rest_api_root/v1/directory/groups/63c8fb32729dee3334d96af7",
                "name": "Open Team"
            }
        }
    ],
    "created_at": 1583290300,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "updated_at": 1583290300,
    "updated_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "is_archived": 0,
    "is_deleted": 0
}
```

---

### 删除一个空间

**`DELETE /v1/wiki/spaces/{space_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `space_id` | String | 空间的id。 |

**响应示例：**

```json
{
    "id": "642fd641209b56920a6c6e5e",
    "url": "https://rest_api_root/v1/wiki/spaces/642fd641209b56920a6c6e5e",
    "identifier": "DEMO",
    "name": "示例空间",
    "scope_type": "user_group",
    "scope_id": "63c8fb32729dee3334d96af7",
    "visibility": "private",
    "color": "#FB7894",
    "description": "示例空间描述",
    "members": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "url": "https://rest_api_root/v1/wiki/spaces/642fd641209b56920a6c6e5e/members/a0417f68e846aae315c85d24643678a9",
            "type": "user",
            "user": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            }
        },
        {
            "id": "63c8fb32729dee3334d96af7",
            "url": "https://rest_api_root/v1/wiki/spaces/642fd641209b56920a6c6e5e/members/63c8fb32729dee3334d96af7",
            "type": "user_group",
            "user_group": {
                "id": "63c8fb32729dee3334d96af7",
                "url": "https://rest_api_root/v1/directory/groups/63c8fb32729dee3334d96af7",
                "name": "Open Team"
            }
        }
    ],
    "created_at": 1583290300,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "updated_at": 1583290400,
    "updated_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "is_archived": 0,
    "is_deleted": 0
}
```

---

### 向空间中添加一个成员

**`POST /v1/wiki/spaces/{space_id}/members`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `space_id` | String | 空间的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `member` | Object | 空间的成员。 |
| `member.id` | String | 企业成员或团队的id。 |
| `member.type` | String | 空间成员的类型。 |
| `role_id` *(可选)* | String | 角色的id。 |

**响应示例：**

```json
{
    "id": "a0417f68e846aae315c85d24643678a9",
    "url": "https: //rest_api_root/v1/wiki/spaces/642fd641209b56920a6c6e5e/members/a0417f68e846aae315c85d24643678a9",
    "space": {
        "id": "642fd641209b56920a6c6e5e",
        "url": "https://rest_api_root/v1/wiki/spaces/642fd641209b56920a6c6e5e",
        "identifier": "DEMO",
        "name": "示例空间",
        "is_archived": 0,
        "is_deleted": 0
    },
    "type": "user",
    "user": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "role": {
        "id": "6422711c3f12e6c1e46d40e6",
        "url": "https://rest_api_root/v1/directory/roles/6422711c3f12e6c1e46d40e6",
        "name": "管理员"
    }
}
```

---

### 在空间中移除一个成员

**`DELETE /v1/wiki/spaces/{space_id}/members/{member_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `space_id` | String | 空间的id。 |
| `member_id` | String | 企业成员或团队的id。 |

**响应示例：**

```json
{
    "id": "a0417f68e846aae315c85d24643678a9",
    "url": "https: //rest_api_root/v1/wiki/spaces/642fd641209b56920a6c6e5e/members/a0417f68e846aae315c85d24643678a9",
    "space": {
        "id": "642fd641209b56920a6c6e5e",
        "url": "https://rest_api_root/v1/wiki/spaces/642fd641209b56920a6c6e5e",
        "identifier": "DEMO",
        "name": "示例空间",
        "is_archived": 0,
        "is_deleted": 0
    },
    "type": "user",
    "user": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "role": {
        "id": "6422711c3f12e6c1e46d40e6",
        "url": "https://rest_api_root/v1/directory/roles/6422711c3f12e6c1e46d40e6",
        "name": "管理员"
    }
}
```

---

### 获取空间中的成员列表

**`GET /v1/wiki/spaces/{space_id}/members`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `space_id` | String | 空间的id。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 2,
    "values": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "url": "https: //rest_api_root/v1/wiki/spaces/642fd641209b56920a6c6e5e/members/a0417f68e846aae315c85d24643678a9",
            "space": {
                "id": "642fd641209b56920a6c6e5e",
                "url": "https://rest_api_root/v1/wiki/spaces/642fd641209b56920a6c6e5e",
                "identifier": "DEMO",
                "name": "示例空间",
                "is_archived": 0,
                "is_deleted": 0
            },
            "type": "user",
            "user": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "role": {
                "id": "6422711c3f12e6c1e46d40e6",
                "url": "https://rest_api_root/v1/directory/roles/6422711c3f12e6c1e46d40e6",
                "name": "管理员"
            }
        },
        {
            "id": "63c8fb32729dee3334d96af7",
            "url": "https: //rest_api_root/v1/wiki/spaces/642fd641209b56920a6c6e5e/members/63c8fb32729dee3334d96af7",
            "space": {
                "id": "642fd641209b56920a6c6e5e",
                "url": "https://rest_api_root/v1/wiki/spaces/642fd641209b56920a6c6e5e",
                "identifier": "DEMO",
                "name": "示例空间",
                "is_archived": 0,
                "is_deleted": 0
            },
            "type": "userGroup",
            "user_group": {
                "id": "63c8fb32729dee3334d96af7",
                "url": "https://rest_api_root/v1/directory/groups/63c8fb32729dee3334d96af7",
                "name": "Open Team"
            },
            "role": {
                "id": "6422711c3f12e6c1e46d40e6",
                "url": "https://rest_api_root/v1/directory/roles/6422711c3f12e6c1e46d40e6",
                "name": "管理员"
            }
       }
    ]
}
```

---

### 获取空间列表

**`GET /v1/wiki/spaces`**

**查询参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `keywords` *(可选)* | String | 关键字。只支持name关键字搜索。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "642fd641209b56920a6c6e5e",
            "url": "https://rest_api_root/v1/wiki/spaces/642fd641209b56920a6c6e5e",
            "identifier": "DEMO",
            "name": "示例空间",
            "scope_type": "user_group",
            "scope_id": "63c8fb32729dee3334d96af7",
            "visibility": "private",
            "color": "#FB7894",
            "description": "示例空间描述",
            "members": [
                {
                    "id": "a0417f68e846aae315c85d24643678a9",
                    "url": "https://rest_api_root/v1/wiki/spaces/642fd641209b56920a6c6e5e/members/a0417f68e846aae315c85d24643678a9",
                    "type": "user",
                    "user": {
                        "id": "a0417f68e846aae315c85d24643678a9",
                        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                        "name": "john",
                        "display_name": "John",
                        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
                    }
                },
                {
                    "id": "63c8fb32729dee3334d96af7",
                    "url": "https://rest_api_root/v1/wiki/spaces/642fd641209b56920a6c6e5e/members/63c8fb32729dee3334d96af7",
                    "type": "user_group",
                    "user_group": {
                        "id": "63c8fb32729dee3334d96af7",
                        "url": "https://rest_api_root/v1/directory/groups/63c8fb32729dee3334d96af7",
                        "name": "Open Team"
                    }
                }
            ],
            "created_at": 1583290300,
            "created_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "updated_at": 1583290400,
            "updated_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "is_archived": 0,
            "is_deleted": 0
        }
    ]
}
```

---

### 部分更新一个空间

**`PATCH /v1/wiki/spaces/{space_id}`**

企业令牌不能更新scope_type为user类型的空间。

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `space_id` | String | 空间的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `name` *(可选)* | String | 空间的名称（不超过32个字符）。 |
| `identifier` *(可选)* | String | 空间的标识。在一个企业中这个标识是唯一的。产品的标识由大写英文字母/数字/下划线/连接线组成（不超过15个字符）。 |
| `description` *(可选)* | String | 空间的描述。 |

**响应示例：**

```json
{
    "id": "642fd641209b56920a6c6e5e",
    "url": "https://rest_api_root/v1/wiki/spaces/642fd641209b56920a6c6e5e",
    "identifier": "DEMO",
    "name": "示例空间",
    "scope_type": "user_group",
    "scope_id": "63c8fb32729dee3334d96af7",
    "visibility": "private",
    "color": "#FB7894",
    "description": "示例空间描述",
    "members": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "url": "https://rest_api_root/v1/wiki/spaces/642fd641209b56920a6c6e5e/members/a0417f68e846aae315c85d24643678a9",
            "type": "user",
            "user": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            }
        },
        {
            "id": "63c8fb32729dee3334d96af7",
            "url": "https://rest_api_root/v1/wiki/spaces/642fd641209b56920a6c6e5e/members/63c8fb32729dee3334d96af7",
            "type": "user_group",
            "user_group": {
                "id": "63c8fb32729dee3334d96af7",
                "url": "https://rest_api_root/v1/directory/groups/63c8fb32729dee3334d96af7",
                "name": "Open Team"
            }
        }
    ],
    "created_at": 1583290300,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "updated_at": 1583290400,
    "updated_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "is_archived": 0,
    "is_deleted": 0
}
```

---

## 组织

### 企业

用于查看企业的基本信息。

**全量数据示例：**

```json
{
     "id": "56ba35de87ad7153c2062f65",
     "url": "https://rest_api_root/v1/directory/team",
     "name": "YCtech",
     "secondary_domain": "yctech"
 }
```

**引用数据示例：**

```json
{
     "id": "56ba35de87ad7153c2062f65",
     "url": "https://rest_api_root/v1/directory/team",
     "name": "YCtech",
     "secondary_domain": "yctech"
 }
```

---

### 企业成员

用于查看和管理企业成员的基本信息。 资源地址：{GET} https://rest_api_root/v1/directory/users/{user_id}

**全量数据示例：**

```json
{
    "id": "a0417f68e846aae315c85d24643678a9",
    "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
    "name": "john",
    "display_name": "John",
    "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png",
    "department": {
        "id": "6422711c3f12e6c1e46d40e6",
        "url": "https://rest_api_root/v1/directory/departments/6422711c3f12e6c1e46d40e6",
        "name": "技术支持"
    },
    "job": {
        "id": "6440c881c56f557eb1aff6e5",
        "url": "https://rest_api_root/v1/directory/jobs/6440c881c56f557eb1aff6e5",
        "name": "后端工程师"
    },
    "email": "john@email.com",
    "mobile": "15000000000",
    "status": "enabled",
    "employee_number": "zxv"
}
```

**引用数据示例：**

```json
{
    "id": "a0417f68e846aae315c85d24643678a9",
    "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
    "name": "john",
    "display_name": "John",
    "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
}
```

---

### 团队

用于查看企业的团队信息。 资源地址：{GET} https://rest_api_root/v1/directory/groups/{group_id}

**全量数据示例：**

```json
{
    "id": "63c8fb32729dee3334d96af7",
    "url": "https://rest_api_root/v1/directory/groups/63c8fb32729dee3334d96af7",
    "name": "Open Team",
    "visibility": "private",
    "description": "This is Open Team."
}
```

**引用数据示例：**

```json
{
    "id": "63c8fb32729dee3334d96af7",
    "url": "https://rest_api_root/v1/directory/groups/63c8fb32729dee3334d96af7",
    "name": "Open Team"
}
```

---

### 团队

用于查看和管理团队成员的信息。

---

### 职位

用于查看企业的职位信息。 资源地址：{GET} https://rest_api_root/v1/directory/jobs/{job_id}

**全量数据示例：**

```json
{
    "id": "6422711c3f12e6c1e46d40e6",
    "url": "https://rest_api_root/v1/directory/jobs/6422711c3f12e6c1e46d40e6",
    "name": "技术总监",
    "is_system": true
}
```

**引用数据示例：**

```json
{
    "id": "6422711c3f12e6c1e46d40e6",
    "url": "https://rest_api_root/v1/directory/jobs/6422711c3f12e6c1e46d40e6",
    "name": "技术总监"
}
```

---

### 角色

用于查看企业的角色信息。 资源地址：{GET} https://rest_api_root/v1/directory/roles/{role_id}

**全量数据示例：**

```json
{
    "id": "6422711c3f12e6c1e46d40e6",
    "url": "https://rest_api_root/v1/directory/roles/6422711c3f12e6c1e46d40e6",
    "name": "管理员",
    "is_system": true
}
```

**引用数据示例：**

```json
{
    "id": "6422711c3f12e6c1e46d40e6",
    "url": "https://rest_api_root/v1/directory/roles/6422711c3f12e6c1e46d40e6",
    "name": "管理员"
}
```

---

### 部门

用于查看和管理企业的部门信息。 资源地址：{GET} https://rest_api_root/v1/directory/departments/{department_id}

**全量数据示例：**

```json
{
    "id": "6422711c3f12e6c1e46d40e6",
    "url": "https://rest_api_root/v1/directory/departments/6422711c3f12e6c1e46d40e6",
    "name": "技术支持",
    "head": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "parent": {
        "id": "6422711c3f12e6c1e46d40e2",
        "url": "https://rest_api_root/v1/directory/departments/6422711c3f12e6c1e46d40e2",
        "name": "技术中心"
    }
}
```

**引用数据示例：**

```json
{
    "id": "6422711c3f12e6c1e46d40e6",
    "url": "https://rest_api_root/v1/directory/departments/6422711c3f12e6c1e46d40e6",
    "name": "技术支持"
}
```

---

## 职位

### 获取职位列表

**`GET /v1/directory/jobs`**

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "6422711c3f12e6c1e46d40e6",
            "url": "https://rest_api_root/v1/directory/jobs/6422711c3f12e6c1e46d40e6",
            "name": "技术总监",
            "is_system": true
        }
    ]
}
```

---

## 角色

### 获取角色列表

**`GET /v1/directory/roles`**

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "6422711c3f12e6c1e46d40e6",
            "url": "https://rest_api_root/v1/directory/roles/6422711c3f12e6c1e46d40e6",
            "name": "管理员",
            "is_system": true
        }
    ]
}
```

---

## 计划

### 全量更新一个执行用例

**`PUT /v1/testhub/runs/{run_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `run_id` | String | 执行用例的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `status_id` | String | 执行用例执行结果的id。 |
| `remark` *(可选)* | String | 执行用例执行结果的备注。 |
| `steps` | Object[] | 执行用例步骤的列表。 |
| `steps.step_id` | String | 执行用例步骤的id。 |
| `steps.status_id` | String | 执行用例步骤执行结果的id。 |
| `steps.actual_value` *(可选)* | String | 执行用例步骤的实际值。 |
| `executor_id` *(可选)* | String | 执行人的id。不传默认执行人为空。 |

**响应示例：**

```json
{
    "id": "547000eb6a70571487623fea",
    "url": "https://rest_api_root/v1/testhub/runs/547000eb6a70571487623fea",
    "status": "pass",
    "short_id": "Aq1u38yX",
    "html_url": "https://yctech.pingcode.com/testhub/runs/Aq1u38yX",
    "library": {
        "id": "5eb623f6a70571487ea47000",
        "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000",
        "identifier": "CSK",
        "name": "测试库",
        "is_archived": 0,
        "is_deleted": 0
    },
    "plan": {
        "id": "5eb6a70571487623fea47000",
        "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/plans/5eb6a70571487623fea47000",
        "name": "测试计划",
        "status": "in_progress",
        "start_at": 1589791860,
        "end_at": 1589791870,
        "short_id": "7nNkViOD",
        "html_url": "https://yctech.pingcode.com/testhub/libraries/CSK/plans/7nNkViOD/runs"
    },
    "case": {
        "id": "5edca524cad2fa112b06305c",
        "url": "https://rest_api_root/v1/testhub/cases/5edca524cad2fa112b06305c",
        "identifier": "CSK-10",
        "title": "这是一个测试用例",
        "level": "P1",
        "short_id": "fdUw3C",
        "html_url": "https://yctech.pingcode.com/testhub/cases/fdUw3C",
        "test_type": "automation",
        "properties": {
            "prop_a": "prop_a_value",
            "prop_b": "prop_b_value"
        }
    },
    "latest_executed_status": {
        "id": "68d117800d5dd2484a198261",
        "url": "http://rest_api_root/v1/testhub/run_statuses/68d117800d5dd2484a198261",
        "name": "通过"
    },
    "suite": {
        "id": "55714870a70ea4eb623f6700",
        "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/suites/55714870a70ea4eb623f6700",
        "name": "登录",
        "paths": "首页/账户"
    },
    "remark": "执行用例执行结果的备注",
    "executor": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "steps": [
        {
            "step_id": "524cad5edb06305cca2fa112",
            "status": "pass",
            "actual_value": null
        }
    ],
    "created_at": 1583290300,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "updated_at": 1583293300,
    "updated_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "is_archived": 0,
    "is_deleted": 0
}
```

---

### 创建一个执行用例

**`POST /v1/testhub/runs`**

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `library_id` | String | 测试库的id。 |
| `plan_id` | String | 测试计划的id。 |
| `case_id` | String | 测试用例的id。 |
| `executor_id` *(可选)* | String | 执行人的id。 |

**响应示例：**

```json
{
    "id": "547000eb6a70571487623fea",
    "url": "https://rest_api_root/v1/testhub/runs/547000eb6a70571487623fea",
    "status": "not_start",
    "short_id": "Aq1u38yX",
    "html_url": "https://yctech.pingcode.com/testhub/runs/Aq1u38yX",
    "library": {
        "id": "5eb623f6a70571487ea47000",
        "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000",
        "identifier": "CSK",
        "name": "测试库",
        "is_archived": 0,
        "is_deleted": 0
    },
    "plan": {
        "id": "5eb6a70571487623fea47000",
        "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/plans/5eb6a70571487623fea47000",
        "name": "测试计划",
        "status": "in_progress",
        "start_at": 1589791860,
        "end_at": 1589791870,
        "short_id": "7nNkViOD",
        "html_url": "https://yctech.pingcode.com/testhub/libraries/CSK/plans/7nNkViOD/runs"
    },
    "case": {
        "id": "5edca524cad2fa112b06305c",
        "url": "https://rest_api_root/v1/testhub/cases/5edca524cad2fa112b06305c",
        "identifier": "CSK-10",
        "title": "这是一个测试用例",
        "level": "P1",
        "short_id": "fdUw3C",
        "html_url": "https://yctech.pingcode.com/testhub/cases/fdUw3C",
        "test_type": "automation",
        "properties": {
            "prop_a": "prop_a_value",
            "prop_b": "prop_b_value"
        }
    },
    "latest_executed_status": {
        "id": "68d117800d5dd2484a198265",
        "url": "http://rest_api_root/v1/testhub/run_statuses/68d117800d5dd2484a198265",
        "name": "未测"
    },
    "suite": {
        "id": "55714870a70ea4eb623f6700",
        "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/suites/55714870a70ea4eb623f6700",
        "name": "登录",
        "paths": "首页/账户"
    },
    "remark": null,
    "executor": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "steps": [
        {
            "step_id": "524cad5edb06305cca2fa112",
            "status": "not_start",
            "actual_value": null
        }
    ],
    "created_at": 1583290300,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "updated_at": 1583290300,
    "updated_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "is_archived": 0,
    "is_deleted": 0
}
```

---

### 创建一个计划

**`POST /v1/testhub/libraries/{library_id}/plans`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `library_id` | String | 测试库的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `name` | String | 测试计划的名称。名称在一个测试库里唯一。 |
| `type_id` | String | 测试计划类型的id。 |
| `start_at` | Number | 测试计划的开始时间。 |
| `end_at` | Number | 测试计划的结束时间。 |
| `assignee_id` | String | 测试计划负责人的id。 |
| `project_id` *(可选)* | String | 项目的id。该字段在 sprint_id 或 version_id 有值时必填。 |
| `sprint_id` *(可选)* | String | 迭代的id。该字段仅在 type_id 代表迭代测试时有效。 |
| `version_id` *(可选)* | String | 发布的id。该字段仅在 type_id 代表发布测试时有效。 |

**响应示例：**

```json
{
    "id": "5eb6a70571487623fea47000",
    "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/plans/5eb6a70571487623fea47000",
    "library": {
        "id": "5eb623f6a70571487ea47000",
        "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000",
        "identifier": "CSK",
        "name": "测试库",
        "is_archived": 0,
        "is_deleted": 0
    },
    "name": "测试计划",
    "state": {
        "id": "652d0cb2b798f983d9c67c2b",
        "url": "http://rest_api_root/v1/testhub/plan_states/652d0cb2b798f983d9c67c2c",
        "name": "进行中",
        "type": "in_progress"
    },
    "start_at": 1589791860,
    "end_at": 1589791870,
    "short_id": "7nNkViOD",
    "html_url": "https://yctech.pingcode.com/testhub/libraries/CSK/plans/7nNkViOD/runs",
    "type": {
        "id": "641d0ab2b998f883f9c67b2c",
        "url": "http://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/plan_types/641d0ab2b998f883f9c67b2c",
        "name": "发布测试"
    },
    "project": {
        "id": "5eb623f6a70571487ea41919",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919",
        "identifier": "SCR",
        "name": "Scrum项目",
        "type": "scrum",
        "is_archived": 0,
        "is_deleted": 0
    },
    "sprint": null,
    "version": {
        "id": "5eb623f6a70571487ea47001",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919/versions/5eb623f6a70571487ea47001",
        "name": "1.0.0",
        "start_at": 1565255712,
        "end_at": 1565255879,
        "stage": {
            "id": "5f44a8f8bb347b14b49624a1",
            "url": "https://rest_api_root/v1/project/stages/5f44a8f8bb347b14b49624a1",
            "name": "未开始",
            "type": "pending",
            "color": "#FA8888"
        }
    },
    "assignee": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "summary": "",
    "created_at": 1565366200,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "updated_at": 1565366200,
    "updated_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    }
}
```

---

### 批量创建执行用例

**`POST /v1/testhub/runs/bulk`**

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `runs` | Object[] | 创建单个执行用例必要参数的数组。数组长度不超过100。 |
| `runs.library_id` | String | 执行用例所属测试库的id。 |
| `runs.plan_id` | String | 执行用例所属测试计划的id。 |
| `runs.case_id` | String | 测试用例的id。 |
| `runs.executor_id` *(可选)* | String | 执行人的id。 |

**响应示例：**

```json
[
    {
        "state": "success",
        "run": {
            "id": "547000eb6a70571487623fea",
            "url": "https://rest_api_root/v1/testhub/runs/547000eb6a70571487623fea",
            "status": "not_start",
            "short_id": "Aq1u38yX",
            "html_url": "https://yctech.pingcode.com/testhub/runs/Aq1u38yX",
            "library": {
                "id": "5eb623f6a70571487ea47000",
                "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000",
                "identifier": "CSK",
                "name": "测试库",
                "is_archived": 0,
                "is_deleted": 0
            },
            "plan": {
                "id": "5eb6a70571487623fea47000",
                "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/plans/5eb6a70571487623fea47000",
                "name": "测试计划",
                "status": "in_progress",
                "start_at": 1589791860,
                "end_at": 1589791870
            },
            "case": {
                "id": "5edca524cad2fa112b06305c",
                "url": "https://rest_api_root/v1/testhub/cases/5edca524cad2fa112b06305c",
                "identifier": "CSK-10",
                "title": "这是一个测试用例",
                "level": "P1",
                "short_id": "fdUw3C",
                "html_url": "https://yctech.pingcode.com/testhub/cases/fdUw3C",
                "test_type": "automation",
                "properties": {
                    "prop_a": "prop_a_value",
                    "prop_b": "prop_b_value"
                }
            },
            "latest_executed_status": {
                "id": "68d117800d5dd2484a198265",
                "url": "http://rest_api_root/v1/testhub/run_statuses/68d117800d5dd2484a198265",
                "name": "未测"
            },
            "suite": {
                "id": "55714870a70ea4eb623f6700",
                "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/suites/55714870a70ea4eb623f6700",
                "name": "登录",
                "paths": "首页/账户"
            },
            "executor": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "remark": "执行用例执行结果的备注",
            "steps": [
                {
                    "step_id": "524cad5edb06305cca2fa112",
                    "status": "not_start",
                    "actual_value": null
                }
            ],
            "created_at": 1583290300,
            "created_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "updated_at": 1583290300,
            "updated_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "is_archived": 0,
            "is_deleted": 0
        }
    },
    {
        "state": "failure",
        "run": {
            "library_id": "5edca524cad2fa112b06305a",
            "plan_id": "5edca524cad2fa112b06305b",
            "case_id": "5edca524cad2fa112b06305d",
            "executor_id": "a0417f68e846aae315c85d24643678a9"
        },
        "message": "创建失败或已创建"
    }
]
```

---

### 批量操作执行用例

**`POST /v1/testhub/libraries/{library_id}/plans/{plan_id}/runs/bulk`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `library_id` | String | 测试库的id。 |
| `plan_id` | String | 测试计划的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `inserts` *(可选)* | Object[] | 需要批量创建的执行用例。该参数是一个对象数组（数组内对象不得超过50个）。 |
| `inserts.case_id` | String | 测试用例id。 |
| `inserts.executor_id` *(可选)* | String | 执行人id。 |
| `updates` *(可选)* | Object[] | 需要批量更新的执行用例。该参数是一个对象数组（数组内对象不得超过50个）。 |
| `updates.run_id` | String | 执行用例的id。 |
| `updates.status_id` | String | 执行用例执行结果的id。 |
| `updates.steps` *(可选)* | Object[] | 执行用例步骤的数组。 |
| `updates.steps.step_id` | String | 执行用例步骤的id。 |
| `updates.steps.status_id` | String | 执行用例步骤执行结果的id。 |
| `updates.steps.actual_value` *(可选)* | String | 执行用例步骤的实际值。 |
| `updates.executor_id` *(可选)* | String | 执行人的id。 |
| `deletes` *(可选)* | String[] | 需要批量删除的执行用例。该参数是一个执行用例id的数组（数组内id不得超过50个）。 |

**响应示例：**

```json
{
    "inserts": 1,
    "updates": 1,
    "deletes": 1
}
```

---

### 批量部分更新执行用例

**`PATCH /v1/testhub/runs/bulk`**

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `runs` | Object[] | 部分更新单个执行用例必要参数的数组。 |
| `runs.run_id` | String | 执行用例的id。 |
| `runs.status_id` | String | 执行用例执行结果的id。 |
| `runs.remark` *(可选)* | String | 执行用例执行结果的备注。 |
| `runs.steps` *(可选)* | Object[] | 执行用例的步骤列表。 |
| `runs.steps.step_id` | String | 执行用例步骤的id。 |
| `runs.steps.status_id` | String | 执行用例步骤执行结果的id。 |
| `runs.steps.actual_value` *(可选)* | String | 执行用例步骤的实际值。 |
| `runs.executor_id` *(可选)* | String | 执行人的id。 |

**响应示例：**

```json
[
    {
        "state": "success",
        "run": {
            "id": "5edca524cad2fa112b06305c",
            "url": "https://rest_api_root/v1/testhub/runs/547000eb6a70571487623fea",
            "status": "not_start",
            "short_id": "Aq1u38yX",
            "html_url": "https://yctech.pingcode.com/testhub/runs/Aq1u38yX",
            "library": {
                "id": "5eb623f6a70571487ea47000",
                "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000",
                "identifier": "CSK",
                "name": "测试库",
                "is_archived": 0,
                "is_deleted": 0
            },
            "plan": {
                "id": "5eb6a70571487623fea47000",
                "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/plans/5eb6a70571487623fea47000",
                "name": "测试计划",
                "status": "in_progress",
                "start_at": 1589791860,
                "end_at": 1589791870
            },
            "case": {
                "id": "5edca524cad2fa112b06305c",
                "url": "https://rest_api_root/v1/testhub/cases/5edca524cad2fa112b06305c",
                "identifier": "CSK-10",
                "title": "这是一个测试用例",
                "level": "P1",
                "short_id": "fdUw3C",
                "html_url": "https://yctech.pingcode.com/testhub/cases/fdUw3C",
                "test_type": "automation",
                "properties": {
                    "prop_a": "prop_a_value",
                    "prop_b": "prop_b_value"
                }
            },
            "latest_executed_status": {
                "id": "68d117800d5dd2484a198265",
                "url": "http://rest_api_root/v1/testhub/run_statuses/68d117800d5dd2484a198265",
                "name": "未测"
            },
            "suite": {
                "id": "55714870a70ea4eb623f6700",
                "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/suites/55714870a70ea4eb623f6700",
                "name": "登录",
                "paths": "首页/账户"
            },
            "executor": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "remark": null,
            "steps": [
                {
                    "step_id": "524cad5edb06305cca2fa112",
                    "status": "not_start",
                    "actual_value": "正常访问PingCode官网"
                }
            ],
            "created_at": 1583290300,
            "created_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "updated_at": 1583299300,
            "updated_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "is_archived": 0,
            "is_deleted": 0
        }
    },
    {
        "state": "success",
        "run": {
            "id": "5edca524cad2fa112b06305d",
            "url": "https://rest_api_root/v1/testhub/runs/547000eb6a70571487623fea",
            "status": "not_start",
            "short_id": "Aq1u38yX",
            "html_url": "https://yctech.pingcode.com/testhub/runs/Aq1u38yX",
            "library": {
                "id": "5eb623f6a70571487ea47000",
                "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000",
                "identifier": "CSK",
                "name": "测试库",
                "is_archived": 0,
                "is_deleted": 0
            },
            "plan": {
                "id": "5eb6a70571487623fea47000",
                "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/plans/5eb6a70571487623fea47000",
                "name": "测试计划",
                "status": "in_progress",
                "start_at": 1589791860,
                "end_at": 1589791870
            },
            "case": {
                "id": "5edca524cad2fa112b06305c",
                "url": "https://rest_api_root/v1/testhub/cases/5edca524cad2fa112b06305c",
                "identifier": "CSK-10",
                "title": "这是一个测试用例",
                "level": "P1",
                "short_id": "fdUw3C",
                "html_url": "https://yctech.pingcode.com/testhub/cases/fdUw3C",
                "test_type": "automation",
                "properties": {
                    "prop_a": "prop_a_value",
                    "prop_b": "prop_b_value"
                }
            },
            "latest_executed_status": {
                "id": "68d117800d5dd2484a198265",
                "url": "http://rest_api_root/v1/testhub/run_statuses/68d117800d5dd2484a198265",
                "name": "未测"
            },
            "suite": {
                "id": "55714870a70ea4eb623f6700",
                "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/suites/55714870a70ea4eb623f6700",
                "name": "登录",
                "paths": "首页/账户"
            },
            "executor": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "remark": "执行用例执行结果的备注",
            "steps": [
                {
                    "step_id": "524cad5edb06305cca2fa112",
                    "status": "not_start",
                    "actual_value": "正常访问PingCode官网"
                }
            ],
            "created_at": 1583290300,
            "created_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "updated_at": 1583299300,
            "updated_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "is_archived": 0,
            "is_deleted": 0
        }
    }
]
```

---

### 获取执行用例列表

**`GET /v1/testhub/runs`**

**查询参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `plan_id` *(可选)* | String | 测试计划的id。 |
| `case_id` *(可选)* | String | 测试用例的id。 |
| `suite_id` *(可选)* | String | 测试模块的id。 |
| `status_id` *(可选)* | String | 执行用例执行结果的id。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "547000eb6a70571487623fea",
            "url": "https://rest_api_root/v1/testhub/runs/547000eb6a70571487623fea",
            "status": "not_start",
            "short_id": "Aq1u38yX",
            "html_url": "https://yctech.pingcode.com/testhub/runs/Aq1u38yX",
            "library": {
                "id": "5eb623f6a70571487ea47000",
                "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000",
                "identifier": "CSK",
                "name": "测试库",
                "is_archived": 0,
                "is_deleted": 0
            },
            "plan": {
                "id": "5eb6a70571487623fea47000",
                "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/plans/5eb6a70571487623fea47000",
                "name": "测试计划",
                "status": "in_progress",
                "start_at": 1589791860,
                "end_at": 1589791870,
                "short_id": "7nNkViOD",
                "html_url": "https://yctech.pingcode.com/testhub/libraries/CSK/plans/7nNkViOD/runs"
            },
            "case": {
                "id": "5edca524cad2fa112b06305c",
                "url": "https://rest_api_root/v1/testhub/cases/5edca524cad2fa112b06305c",
                "identifier": "CSK-10",
                "title": "这是一个测试用例",
                "level": "P1",
                "html_url": "https://yctech.pingcode.com/testhub/cases/fdUw3C",
                "test_type": "automation",
                "properties": {
                    "prop_a": "prop_a_value",
                    "prop_b": "prop_b_value"
                }
            },
            "latest_executed_status": {
                "id": "68d117800d5dd2484a198265",
                "url": "http://rest_api_root/v1/testhub/run_statuses/68d117800d5dd2484a198265",
                "name": "未测"
            },
            "suite": {
                "id": "55714870a70ea4eb623f6700",
                "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/suites/55714870a70ea4eb623f6700",
                "name": "登录",
                "paths": "首页/账户"
            },
            "remark": "执行用例执行结果的备注",
            "executor": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "steps": [
                {
                    "step_id": "524cad5edb06305cca2fa112",
                    "status": "not_start",
                    "actual_value": null
                }
            ],
            "created_at": 1583290300,
            "created_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "updated_at": 1583290300,
            "updated_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "is_archived": 0,
            "is_deleted": 0
        }
    ]
}
```

---

### 获取执行用例执行结果列表

**`GET /v1/testhub/run/statuses?library_id={library_id}`**

**查询参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `library_id` | String | 测试库的id。 |

**响应示例：**

```json
{
    "page_index": 0,
    "page_size": 30,
    "total": 5,
    "values": [
        {
            "id": "68d117800d5dd2484a198261",
            "url": "https://rest_api_root/v1/testhub/run_statuses/68d117800d5dd2484a198261",
            "name": "通过"
        },
        {
            "id": "68d117800d5dd2484a198262",
            "url": "https://rest_api_root/v1/testhub/run_statuses/68d117800d5dd2484a198262",
            "name": "受阻"
        },
        {
            "id": "68d117800d5dd2484a198263",
            "url": "https://rest_api_root/v1/testhub/run_statuses/68d117800d5dd2484a198263",
            "name": "失败"
        },
        {
            "id": "68d117800d5dd2484a198264",
            "url": "https://rest_api_root/v1/testhub/run_statuses/68d117800d5dd2484a198264",
            "name": "跳过"
        },
        {
            "id": "68d117800d5dd2484a198265",
            "url": "https://rest_api_root/v1/testhub/run_statuses/68d117800d5dd2484a198265",
            "name": "未测"
        }
    ]
}
```

---

### 获取执行用例的结果记录

**`GET /v1/testhub/runs/{run_id}/histories`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `run_id` | String | 执行用例的id。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "65115f0939286e26e05a66db",
            "url": "https://rest_api_root/v1/testhub/runs/547000eb6a70571487623fea/histories/65115f0939286e26e05a66db",
            "run": {
                "id": "547000eb6a70571487623fea",
                "url": "https://rest_api_root/v1/testhub/runs/547000eb6a70571487623fea",
                "status": "pass",
                "short_id": "Aq1u38yX",
                "html_url": "https://yctech.pingcode.com/testhub/runs/Aq1u38yX"
            },
            "library": {
                "id": "5eb623f6a70571487ea47000",
                "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000",
                "identifier": "CSK",
                "name": "测试库",
                "is_archived": 0,
                "is_deleted": 0
            },
            "plan": {
                "id": "5eb6a70571487623fea47000",
                "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/plans/5eb6a70571487623fea47000",
                "name": "测试计划",
                "status": "in_progress",
                "start_at": 1589791860,
                "end_at": 1589791870,
                "short_id": "7nNkViOD",
                "html_url": "https://yctech.pingcode.com/testhub/libraries/CSK/plans/7nNkViOD/runs"
            },
            "case": {
                "id": "5edca524cad2fa112b06305c",
                "url": "https://rest_api_root/v1/testhub/cases/5edca524cad2fa112b06305c",
                "identifier": "CSK-10",
                "title": "这是一个测试用例",
                "level": "P1",
                "short_id": "fdUw3C",
                "html_url": "https://yctech.pingcode.com/testhub/cases/fdUw3C",
                "test_type": "automation",
                "properties": {
                    "prop_a": "prop_a_value",
                    "prop_b": "prop_b_value"
                }
            },
            "executed_status": {
                "id": "68d117800d5dd2484a198261",
                "url": "http://rest_api_root/v1/testhub/run_statuses/68d117800d5dd2484a198261",
                "name": "通过"
            },
            "remark": "执行用例执行结果的备注",
            "executed_at": 1583290300,
            "executed_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "steps": [
                {
                    "step_id": "524cad5edb06305cca2fa112",
                    "status": "pass",
                    "actual_value": null
                }
            ]
        }
    ]
}
```

---

### 获取计划列表

**`GET /v1/testhub/libraries/{library_id}/plans`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `library_id` | String | 测试库的id。 |

**查询参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `name` *(可选)* | String | 测试计划名称。 |
| `created_between` *(可选)* | String | 创建时间介于的时间范围，通过','分割起始时间。 |
| `updated_between` *(可选)* | String | 更新时间介于的时间范围，通过','分割起始时间。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "5eb6a70571487623fea47000",
            "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/plans/5eb6a70571487623fea47000",
            "library": {
                "id": "5eb623f6a70571487ea47000",
                "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000",
                "identifier": "CSK",
                "name": "测试库",
                "is_archived": 0,
                "is_deleted": 0
            },
            "name": "测试计划",
            "state": {
                "id": "652d0cb2b798f983d9c67c2b",
                "url": "http://rest_api_root/v1/testhub/plan_states/652d0cb2b798f983d9c67c2b",
                "name": "进行中",
                "type": "in_progress"
            },
            "start_at": 1589791860,
            "end_at": 1589791870,
            "short_id": "7nNkViOD",
            "html_url": "https://yctech.pingcode.com/testhub/libraries/CSK/plans/7nNkViOD/runs",
            "type": {
                "id": "641d0ab2b998f883f9c67b2c",
                "url": "http://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/plan_types/641d0ab2b998f883f9c67b2c",
                "name": "发布测试"
            },
            "project": {
                "id": "5eb623f6a70571487ea41919",
                "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919",
                "identifier": "SCR",
                "name": "Scrum项目",
                "type": "scrum",
                "is_archived": 0,
                "is_deleted": 0
            },
            "sprint": null,
            "version": {
                "id": "5eb623487ea47001f6a70571",
                "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919/versions/5eb623487ea47001f6a70571",
                "name": "1.0.0",
                "start_at": 1565255712,
                "end_at": 1565255879,
                "stage": {
                    "id": "5f44a8f8bb347b14b49624a1",
                    "url": "https://rest_api_root/v1/project/stages/5f44a8f8bb347b14b49624a1",
                    "name": "未开始",
                    "type": "pending",
                    "color": "#FA8888"
                }
            },
            "assignee": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "summary": "测试报告的概要",
            "created_at": 1565366200,
            "created_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "updated_at": 1565366200,
            "updated_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            }
        }
    ]
}
```

---

### 获取计划类型列表

**`GET /v1/testhub/libraries/{library_id}/plan_types`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `library_id` | String | 测试库的id。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "642f765b6950bc66cfa82f05",
            "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/plan_types/642f765b6950bc66cfa82f05",
            "library": {
                "id": "5eb623f6a70571487ea47000",
                "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000",
                "identifier": "CSK",
                "name": "测试库",
                "is_archived": 0,
                "is_deleted": 0
            },
            "name": "普通测试"
        }
    ]
}
```

---

### 部分更新一个执行用例

**`PATCH /v1/testhub/runs/{run_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `run_id` | String | 执行用例的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `status_id` | String | 执行用例执行结果的id。 |
| `remark` *(可选)* | String | 执行用例执行结果的备注。 |
| `steps` *(可选)* | Object[] | 执行用例步骤的列表。steps是整体更新的。 |
| `steps.step_id` | String | 执行用例步骤的id。 |
| `steps.status_id` | String | 执行用例步骤执行结果的id。 |
| `steps.actual_value` *(可选)* | String | 执行用例步骤的实际值。 |
| `executor_id` *(可选)* | String | 执行人的id。不传默认执行人为执行用例的创建人。 |

**响应示例：**

```json
{
    "id": "547000eb6a70571487623fea",
    "url": "https://rest_api_root/v1/testhub/runs/547000eb6a70571487623fea",
    "status": "not_start",
    "short_id": "Aq1u38yX",
    "html_url": "https://yctech.pingcode.com/testhub/runs/Aq1u38yX",
    "library": {
        "id": "5eb623f6a70571487ea47000",
        "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000",
        "identifier": "CSK",
        "name": "测试库",
        "is_archived": 0,
        "is_deleted": 0
    },
    "plan": {
        "id": "5eb6a70571487623fea47000",
        "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/plans/5eb6a70571487623fea47000",
        "name": "测试计划",
        "status": "in_progress",
        "start_at": 1589791860,
        "end_at": 1589791870,
        "short_id": "7nNkViOD",
        "html_url": "https://yctech.pingcode.com/testhub/libraries/CSK/plans/7nNkViOD/runs"
    },
    "case": {
        "id": "5edca524cad2fa112b06305c",
        "url": "https://rest_api_root/v1/testhub/cases/5edca524cad2fa112b06305c",
        "identifier": "CSK-10",
        "title": "这是一个测试用例",
        "level": "P1",
        "html_url": "https://yctech.pingcode.com/testhub/cases/fdUw3C",
        "test_type": "automation",
        "properties": {
            "prop_a": "prop_a_value",
            "prop_b": "prop_b_value"
        }
    },
    "latest_executed_status": {
        "id": "68d117800d5dd2484a198265",
        "url": "http://rest_api_root/v1/testhub/run_statuses/68d117800d5dd2484a198265",
        "name": "未测"
    },
    "suite": {
        "id": "55714870a70ea4eb623f6700",
        "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/suites/55714870a70ea4eb623f6700",
        "name": "登录",
        "paths": "首页/账户"
    },
    "remark": "执行用例执行结果的备注",
    "executor": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "steps": [
        {
            "step_id": "524cad5edb06305cca2fa112",
            "status": "not_start",
            "actual_value": null
        }
    ],
    "created_at": 1583290300,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "updated_at": 1583290300,
    "updated_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "is_archived": 0,
    "is_deleted": 0
}
```

---

### 部分更新一个计划

**`PATCH /v1/testhub/libraries/{library_id}/plans/{plan_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `library_id` | String | 测试库的id。 |
| `plan_id` | String | 测试计划的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `name` *(可选)* | String | 测试计划的名称。名称在一个测试库里唯一。 |
| `type_id` *(可选)* | String | 测试计划类型的id。指定测试计划类型时，建议同时指定对应的 sprint_id 或 version_id。 |
| `project_id` *(可选)* | String | 项目的id。 |
| `sprint_id` *(可选)* | String | 迭代的id。该字段仅在测试计划类型为迭代测试时有效。 |
| `version_id` *(可选)* | String | 发布的id。该字段仅在测试计划类型为发布测试时有效。 |
| `start_at` *(可选)* | Number | 测试计划的开始时间。 |
| `end_at` *(可选)* | Number | 测试计划的结束时间。 |
| `assignee_id` *(可选)* | String | 测试计划负责人的id。 |
| `state_id` *(可选)* | String | 测试计划状态的id。 |
| `summary` *(可选)* | String | 测试报告的概要信息。 |

**响应示例：**

```json
{
    "id": "5eb6a70571487623fea47000",
    "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/plans/5eb6a70571487623fea47000",
    "library": {
        "id": "5eb623f6a70571487ea47000",
        "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000",
        "identifier": "CSK",
        "name": "测试库",
        "is_archived": 0,
        "is_deleted": 0
    },
    "name": "测试计划",
    "state": {
        "id": "652d0cb2b798f983d9c67c2b",
        "url": "http://rest_api_root/v1/testhub/plan_states/652d0cb2b798f983d9c67c2c",
        "name": "进行中",
        "type": "in_progress"
    },
    "start_at": 1589791860,
    "end_at": 1589791870,
    "short_id": "7nNkViOD",
    "html_url": "https://yctech.pingcode.com/testhub/libraries/CSK/plans/7nNkViOD/runs",
    "type": {
        "id": "641d0ab2b998f883f9c67b2c",
        "url": "http://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/plan_types/641d0ab2b998f883f9c67b2c",
        "name": "发布测试"
    },
    "project": {
        "id": "5eb623f6a70571487ea41919",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919",
        "identifier": "SCR",
        "name": "Scrum项目",
        "type": "scrum",
        "is_archived": 0,
        "is_deleted": 0
    },
    "sprint": null,
    "version": {
        "id": "5eb623487ea47001f6a70571",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea41919/versions/5eb623487ea47001f6a70571",
        "name": "1.0.0",
        "start_at": 1565255712,
        "end_at": 1565255879,
        "stage": {
            "id": "5f44a8f8bb347b14b49624a1",
            "url": "https://rest_api_root/v1/project/stages/5f44a8f8bb347b14b49624a1",
            "name": "未开始",
            "type": "pending",
            "color": "#FA8888"
        }
    },
    "assignee": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "summary": "测试报告的概要",
    "created_at": 1565366200,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "updated_at": 1565366200,
    "updated_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    }
}
```

---

## 计划配置

### 获取全部计划状态列表

**`GET /v1/testhub/plan_states`**

**响应示例：**

```json
{
    "page_index": 0,
    "page_size": 30,
    "total": 4,
    "values": [
        {
            "id": "686f62038668bbae4f4dd0ca",
            "url": "http://rest_api_root/v1/testhub/plan_states/686f62038668bbae4f4dd0ca",
            "name": "未开始",
            "type": "pending",
            "is_system": true
        },
        {
            "id": "652d0cb2b798f983d9c67c2b",
            "url": "http://rest_api_root/v1/testhub/plan_states/652d0cb2b798f983d9c67c2b",
            "name": "进行中",
            "type": "in_progress",
            "is_system": true
        },
        {
            "id": "652d0cb2b798f983d9c67c2d",
            "url": "http://rest_api_root/v1/testhub/plan_states/652d0cb2b798f983d9c67c2d",
            "name": "已完成",
            "type": "completed",
            "is_system": true
        },
        {
            "id": "652d0cb2b798f983d9c67c2e",
            "url": "http://rest_api_root/v1/testhub/plan_states/652d0cb2b798f983d9c67c2e",
            "name": "已终止",
            "type": "completed",
            "is_system": false
        }
    ]
}
```

---

## 评审

### 创建一个评审

**`POST /v1/reviews`**

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `title` | String | 评审的标题。 |
| `pilot_id` | String | 评审主体所在产品、项目或测试库的id。 |
| `principal_type` | String | 评审主体的类型。 |
| `description` *(可选)* | String | 评审的描述。 |

**响应示例：**

```json
{
    "id": "5f168f764eba01a5278b87cd",
    "url": "https://rest_api_root/v1/reviews/5f168f764eba01a5278b87cd?principal_type=idea",
    "identifier": "SCR-R5",
    "title": "这是一个评审",
    "status": "pending",
    "principal_type": "idea",
    "short_id": "LsEy8mZF",
    "html_url": "https://yctech.pingcode.com/reviews/LsEy8mZF",
    "pilot": {
        "id": "63bb744314bd13c9def24cb4",
        "url": "https://rest_api_root/v1/ship/products/63bb744314bd13c9def24cb4",
        "name": "示例产品",
        "identifier": "SLC",
        "is_archived": 0,
        "is_deleted": 0
    },
    "description": "这是一个评审的描述",
    "participants": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "url": "https://rest_api_root/v1/participants/a0417f68e846aae315c85d24643678a9?principal_type=idea&review_id=6f168f764eba01a5278b87cd",
            "type": "user",
            "user": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            }
        }
    ],
    "submitted_at": null,
    "submitted_by": null,
    "completed_at": null,
    "completed_by": null,
    "created_at": 1593290347,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/avatar.png"
    },
    "updated_at": 1593291347,
    "updated_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/avatar.png"
    }
}
```

---

### 删除一个评审

**`DELETE /v1/reviews/{review_id}?principal_type={principal_type}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `review_id` | String | 评审的id。 |

**查询参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `principal_type` | String | 评审主体的类型。 |

**响应示例：**

```json
{
    "id": "5f168f764eba01a5278b87cd",
    "url": "https://rest_api_root/v1/reviews/5f168f764eba01a5278b87cd?principal_type=idea",
    "identifier": "SCR-R5",
    "title": "这是一个评审",
    "status": "pending",
    "principal_type": "idea",
    "short_id": "LsEy8mZF",
    "html_url": "https://yctech.pingcode.com/reviews/LsEy8mZF",
    "pilot": {
        "id": "63bb744314bd13c9def24cb4",
        "url": "https://rest_api_root/v1/ship/products/63bb744314bd13c9def24cb4",
        "name": "示例产品",
        "identifier": "SLC",
        "is_archived": 0,
        "is_deleted": 0
    },
    "description": "这是一个评审的描述",
    "participants": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "url": "https://rest_api_root/v1/participants/a0417f68e846aae315c85d24643678a9?principal_type=idea&review_id=6f168f764eba01a5278b87cd",
            "type": "user",
            "user": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            }
        }
    ],
    "submitted_at": null,
    "submitted_by": null,
    "completed_at": null,
    "completed_by": null,
    "created_at": 1593290347,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/avatar.png"
    },
    "updated_at": 1593291347,
    "updated_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/avatar.png"
    }
}
```

---

### 向评审中添加一个评审内容

**`POST /v1/reviews/{review_id}/principals`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `review_id` | String | 评审的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `principal_id` | String | 评审内容的id。 |
| `principal_type` | String | 评审主体的类型。 |

**响应示例：**

```json
{
    "id": "63bb744514bd13c9def24ceb",
    "url": "https://rest_api_root/v1/reviews/68ccfe6b3eef8131da564e4a/principals/63bb744514bd13c9def24ceb?principal_type=idea",
    "review": {
        "id": "68ccfe6b3eef8131da564e4a",
        "url": "https://rest_api_root/v1/reviews/68ccfe6b3eef8131da564e4a?principal_type=idea",
        "identifier": "SLC-R11",
        "title": "这是一个需求评审",
        "status": "pending",
        "principal_type": "idea",
        "short_id": "LK7Qy-NA",
        "html_url": "https://yctech.pingcode.com/ship/reviews/LK7Qy-NA"
    },
    "principal_type": "idea",
    "principal": {
        "id": "63bb744514bd13c9def24ceb",
        "url": "https://rest_api_root/v1/ship/ideas/63bb744514bd13c9def24ceb",
        "identifier": "SLC-1",
        "title": "这是一个产品需求",
        "short_id": "Omi8PFL0",
        "html_url": "https://yctech.pingcode.com/ship/ideas/Omi8PFL0"
    }
}
```

---

### 在评审中移除一个评审内容

**`DELETE /v1/reviews/{review_id}/principals/{principal_id}?principal_type={principal_type}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `review_id` | String | 评审的id。 |
| `principal_id` | String | 评审内容的id。 |

**查询参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `principal_type` | String | 评审主体的类型。 |

**响应示例：**

```json
{
    "id": "63bb744514bd13c9def24ceb",
    "url": "https://rest_api_root/v1/reviews/68ccfe6b3eef8131da564e4a/principals/63bb744514bd13c9def24ceb?principal_type=idea",
    "review": {
        "id": "68ccfe6b3eef8131da564e4a",
        "url": "https://rest_api_root/v1/reviews/68ccfe6b3eef8131da564e4a?principal_type=idea",
        "identifier": "SLC-R11",
        "title": "这是一个需求评审",
        "status": "pending",
        "principal_type": "idea",
        "short_id": "LK7Qy-NA",
        "html_url": "https://yctech.pingcode.com/ship/reviews/LK7Qy-NA"
    },
    "principal_type": "idea",
    "principal": {
        "id": "63bb744514bd13c9def24ceb",
        "url": "https://rest_api_root/v1/ship/ideas/63bb744514bd13c9def24ceb",
        "identifier": "SLC-1",
        "title": "这是一个产品需求",
        "short_id": "Omi8PFL0",
        "html_url": "https://yctech.pingcode.com/ship/ideas/Omi8PFL0"
    }
}
```

---

### 获取一个评审内容

**`GET /v1/reviews/{review_id}/principals/{principal_id}?principal_type={principal_type}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `review_id` | String | 评审的id。 |
| `principal_id` | String | 评审内容的id。 |

**查询参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `principal_type` | String | 评审主体的类型。 |

**响应示例：**

```json
{
    "id": "63bb744514bd13c9def24ceb",
    "url": "https://rest_api_root/v1/reviews/68ccfe6b3eef8131da564e4a/principals/63bb744514bd13c9def24ceb?principal_type=idea",
    "review": {
        "id": "68ccfe6b3eef8131da564e4a",
        "url": "https://rest_api_root/v1/reviews/68ccfe6b3eef8131da564e4a?principal_type=idea",
        "identifier": "SLC-R11",
        "title": "这是一个需求评审",
        "status": "pending",
        "principal_type": "idea",
        "short_id": "LK7Qy-NA",
        "html_url": "https://yctech.pingcode.com/ship/reviews/LK7Qy-NA"
    },
    "principal_type": "idea",
    "principal": {
        "id": "63bb744514bd13c9def24ceb",
        "url": "https://rest_api_root/v1/ship/ideas/63bb744514bd13c9def24ceb",
        "identifier": "SLC-1",
        "title": "这是一个产品需求",
        "short_id": "Omi8PFL0",
        "html_url": "https://yctech.pingcode.com/ship/ideas/Omi8PFL0"
    }
}
```

---

### 获取评审内容列表

**`GET /v1/reviews/{review_id}/principals?principal_type={principal_type}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `review_id` | String | 评审的id。 |

**查询参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `principal_type` | String | 评审主体的类型。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "63bb744514bd13c9def24ceb",
            "url": "https://rest_api_root/v1/reviews/68ccfe6b3eef8131da564e4a/principals/63bb744514bd13c9def24ceb?principal_type=idea",
            "review": {
                "id": "68ccfe6b3eef8131da564e4a",
                "url": "https://rest_api_root/v1/reviews/68ccfe6b3eef8131da564e4a?principal_type=idea",
                "identifier": "SLC-R11",
                "title": "这是一个需求评审",
                "status": "pending",
                "principal_type": "idea",
                "short_id": "LK7Qy-NA",
                "html_url": "https://yctech.pingcode.com/ship/reviews/LK7Qy-NA"
            },
            "principal_type": "idea",
            "principal": {
                "id": "63bb744514bd13c9def24ceb",
                "url": "https://rest_api_root/v1/ship/ideas/63bb744514bd13c9def24ceb",
                "identifier": "SLC-1",
                "title": "这是一个产品需求",
                "short_id": "Omi8PFL0",
                "html_url": "https://yctech.pingcode.com/ship/ideas/Omi8PFL0"
            }
        }
    ]
}
```

---

### 获取评审列表

**`GET /v1/reviews?principal_type={principal_type}&pilot_id={pilot_id}`**

**查询参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `principal_type` | String | 评审主体的类型。 |
| `pilot_id` | String | 评审主体所在产品、项目或测试库的id。 |
| `status` *(可选)* | String | 评审的状态。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "5f168f764eba01a5278b87cd",
            "url": "https://rest_api_root/v1/reviews/5f168f764eba01a5278b87cd?principal_type=idea",
            "identifier": "SCR-R5",
            "title": "这是一个评审",
            "status": "completed",
            "principal_type": "idea",
            "short_id": "LsEy8mZF",
            "html_url": "https://yctech.pingcode.com/reviews/LsEy8mZF",
            "pilot": {
                "id": "63bb744314bd13c9def24cb4",
                "url": "https://rest_api_root/v1/ship/products/63bb744314bd13c9def24cb4",
                "name": "示例产品",
                "identifier": "SLC",
                "is_archived": 0,
                "is_deleted": 0
            },
            "description": "这是一个评审的描述",
            "participants": [
                {
                    "id": "a0417f68e846aae315c85d24643678a9",
                    "url": "https://rest_api_root/v1/participants/a0417f68e846aae315c85d24643678a9?principal_type=idea&review_id=6f168f764eba01a5278b87cd",
                    "type": "user",
                    "user": {
                        "id": "a0417f68e846aae315c85d24643678a9",
                        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                        "name": "john",
                        "display_name": "John",
                        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
                    }
                }
            ],
            "submitted_at": 1593290347,
            "submitted_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/avatar.png"
            },
            "completed_at": 1593291347,
            "completed_by": {
                "id": "b2417f68e846aae315c85d24643678b0",
                "url": "https://rest_api_root/v1/directory/users/b2417f68e846aae315c85d24643678b0",
                "name": "mary",
                "display_name": "Mary",
                "avatar": "https://s3.amazonaws.com/bucket/avatar2.png"
            },
            "created_at": 1593290347,
            "created_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/avatar.png"
            },
            "updated_at": 1593291347,
            "updated_by": {
                "id": "b2417f68e846aae315c85d24643678b0",
                "url": "https://rest_api_root/v1/directory/users/b2417f68e846aae315c85d24643678b0",
                "name": "mary",
                "display_name": "Mary",
                "avatar": "https://s3.amazonaws.com/bucket/avatar2.png"
            }
        }
    ]
}
```

---

## 评论

### 创建一个评论

**`POST /v1/comments`**

**请求参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `principal_type` | String | 评论主体的类型。 |
| `principal_id` *(可选)* | String | 评论主体的id。 |
| `review_id` *(可选)* | String | 评论评审的id。principal_id和review_id至少存在一个，若同时存在，则忽略review_id。 |
| `content` | String | 评论的内容。 |
| `created_at` *(可选)* | Number | 评论的创建时间。 |
| `created_by` *(可选)* | String | 评论的创建人。 |

**响应示例：**

```json
{
    "id": "59f72dfaeadb5b5197b7da6d",
    "url": "https://rest_api_root/v1/comments/59f72dfaeadb5b5197b7da6d?principal_type=work_item&principal_id=5edca524cad2fa1125cb0630",
    "content": "这是一个工作项评论",
    "attachment_count": 0,
    "attachments": [],
    "is_reply_comment": false,
    "replied_comment": null,
    "created_at": 1565255712,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "is_deleted": 0
}
```

---

### 删除一个评论

**`DELETE /v1/comments/{comment_id}?principal_type={principal_type}&principal_id={principal_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `comment_id` | String | 评论的id。 |

**查询参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `principal_type` | String | 评论主体的类型。 |
| `principal_id` *(可选)* | String | 评论主体的id。 |
| `review_id` *(可选)* | String | 评论评审的id。principal_id和review_id至少存在一个，若同时存在，则忽略review_id。 |

**响应示例：**

```json
{
    "id": "59f72dfaeadb5b5197b7da6d",
    "url": "https://rest_api_root/v1/comments/59f72dfaeadb5b5197b7da6d?principal_type=work_item&principal_id=5edca524cad2fa1125cb0630",
    "content": "",
    "attachment_count": 0,
    "attachments": [],
    "is_reply_comment": false,
    "replied_comment": null,
    "created_at": 1565255712,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "is_deleted": 1
}
```

---

### 获取评论列表

**`GET /v1/comments?principal_type={principal_type}&principal_id={principal_id}`**

**查询参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `principal_type` | String | 评论主体的类型。 |
| `principal_id` *(可选)* | String | 评论主体的id。 |
| `review_id` *(可选)* | String | 评论评审的id。principal_id和review_id至少存在一个，若同时存在，则忽略review_id。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "59f72dfaeadb5b5197b7da6d",
            "url": "https://rest_api_root/v1/comments/59f72dfaeadb5b5197b7da6d?principal_type=work_item&principal_id=5edca524cad2fa1125cb0630",
            "content": "这是一个工作项评论",
            "attachment_count": 0,
            "attachments": [],
            "is_reply_comment": false,
            "replied_comment": null,
            "created_at": 1565255712,
            "created_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "is_deleted": 0
        },
        {
            "id": "50f72dfaeadb5b5197b7da6e",
            "url": "https://rest_api_root/v1/comments/59f72dfaeadb5b5197b7da6d?principal_type=work_item&principal_id=50f72dfaeadb5b5197b7da6e",
            "content": "这是一个工作项评论回复",
            "attachment_count": 0,
            "attachments": [],
            "is_reply_comment": true,
            "replied_comment": {
                "id": "59f72dfaeadb5b5197b7da6d",
                "url": "https://rest_api_root/v1/comments/59f72dfaeadb5b5197b7da6d?principal_type=work_item&principal_id=5edca524cad2fa1125cb0630",
                "content": "这是一个工作项评论",
                "is_deleted": 0
            },
            "created_at": 1565255712,
            "created_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "is_deleted": 0
        }
    ]
}
```

---

## 通用

### 关注人

用于查看和管理关注人的基本信息。 资源地址：{GET} https://rest_api_root/v1/participants/{participant_id}

**全量数据示例（用户类型）:**

```json
{
    "id": "a0417f68e846aae315c85d24643678a9",
    "url": "https://rest_api_root/v1/participants/a0417f68e846aae315c85d24643678a9?principal_type=work_item&principal_id=63e1bf51760505c8795ebccc",
    "type": "user",
    "user": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    }
}
```

**全量数据示例（评审用户类型）:**

```json
{
    "id": "a0417f68e846aae315c85d24643678a9",
    "url": "https://rest_api_root/v1/participants/a0417f68e846aae315c85d24643678a9?principal_type=idea&review_id=6f168f764eba01a5278b87cd",
    "type": "user",
    "user": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    }
}
```

**全量数据示例（团队类型）:**

```json
{
    "id": "63c8fb32729dee3334d96af7",
    "url": "https://rest_api_root/v1/participants/63c8fb32729dee3334d96af7?principal_type=work_item&principal_id=63e1bf51760505c8795ebccc",
    "type": "user_group",
    "user_group": {
        "id": "63c8fb32729dee3334d96af7",
        "url": "https://rest_api_root/v1/directory/groups/63c8fb32729dee3334d96af7",
        "name": "Open Team"
    }
}
```

**引用数据示例（用户类型）:**

```json
{
    "id": "a0417f68e846aae315c85d24643678a9",
    "url": "https://rest_api_root/v1/participants/a0417f68e846aae315c85d24643678a9?principal_type=work_item&principal_id=63e1bf51760505c8795ebccc",
    "type": "user",
    "user": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    }
}
```

**引用数据示例（团队类型）:**

```json
{
    "id": "63c8fb32729dee3334d96af7",
    "url": "https://rest_api_root/v1/participants/63c8fb32729dee3334d96af7?principal_type=work_item&principal_id=63e1bf51760505c8795ebccc",
    "type": "user_group",
    "user_group": {
        "id": "63c8fb32729dee3334d96af7",
        "url": "https://rest_api_root/v1/directory/groups/63c8fb32729dee3334d96af7",
        "name": "Open Team"
    }
}
```

---

### 关联

用于查看和管理关联的基本信息。 资源地址：{GET} https://rest_api_root/v1/relations/{relation_id}

**全量数据示例（工作项关联需求）：**

```json
{
    "id": "653b1b4a3113be5bb040e4c5",
    "url": "https://rest_api_root/v1/relations/653b1b4a3113be5bb040e4c5",
    "principal_type": "work_item",
    "principal": {
        "id": "5edca524cad2fa1125cb0630",
        "url": "https://rest_api_root/v1/project/work_items/5edca524cad2fa1125cb0630",
        "identifier": "SCR-5",
        "title": "这是一个缺陷",
        "type": "bug",
        "start_at": 1583290309,
        "end_at": 1583290347,
        "parent_id": "5edca524cad2fa112b06105c",
        "short_id": "c9WqLmTO",
        "html_url": "https://yctech.pingcode.com/pjm/workitems/c9WqLmTO",
        "properties": {
            "prop_a": "prop_a_value",
            "prop_b": "prop_b_value"
        }
    },
    "target_type": "idea",
    "target": {
        "id": "64b4d70ba368e6594360ea24",
        "url": "https://rest_api_root/v1/ship/ideas/64b4d70ba368e6594360ea24",
        "identifier": "SLC-1",
        "title": "示例需求",
        "short_id": "Ogf1EYey",
        "html_url": "https://yctech.pingcode.com/ship/ideas/Ogf1EYey"
    }
}
```

**全量数据示例（测试计划关联工作项）：**

```json
{
    "id": "fa1125cb06305edca524cad2",
    "url": "https://rest_api_root/v1/relations/fa1125cb06305edca524cad2",
    "principal_type": "test_plan",
    "principal": {
        "id": "5eb6a70571487623fea47000",
        "url": "https://rest_api_root/v1/testhub/libraries/5eb623f6a70571487ea47000/plans/5eb6a70571487623fea47000",
        "name": "测试计划",
        "status": "in_progress",
        "start_at": 1589791860,
        "end_at": 1589791870
    },
    "target_type": "work_item",
    "target": {
        "id": "5edca524cad2fa1125cb0630",
        "url": "https://rest_api_root/v1/project/work_items/5edca524cad2fa1125cb0630",
        "identifier": "SCR-5",
        "title": "这是一个缺陷",
        "type": "bug",
        "start_at": 1583290309,
        "end_at": 1583290347,
        "parent_id": "5edca524cad2fa112b06105c",
        "short_id": "c9WqLmTO",
        "html_url": "https://yctech.pingcode.com/pjm/workitems/c9WqLmTO",
        "properties": {
            "prop_a": "prop_a_value",
            "prop_b": "prop_b_value"
        }
    }
}
```

**全量数据示例（工单关联测试用例）：**

```json
{
    "id": "6552d7ab1bffb252b82645ba",
    "url": "https://rest_api_root/v1/relations/6552d7ab1bffb252b82645ba",
    "principal_type": "ticket",
    "principal": {
        "id": "63eca888a0a13a3efc8d4a43",
        "url": "https://rest_api_root/v1/ship/tickets/63eca888a0a13a3efc8d4a43",
        "identifier": "SLC-T1",
        "title": "希望新增支持第三方账号注册",
        "short_id": "pdMUgQzZ",
        "html_url": "https://yctech.pingcode.com/ship/tickets/pdMUgQzZ"
    },
    "target_type": "test_case",
    "target": {
        "id": "5edca524cad2fa112b06305c",
        "url": "https://rest_api_root/v1/testhub/cases/5edca524cad2fa112b06305c",
        "identifier": "CSK-10",
        "title": "这是一个测试用例",
        "level": "P1",
        "short_id": "fdUw3C",
        "html_url": "https://yctech.pingcode.com/testhub/cases/fdUw3C",
        "properties": {
            "prop_a": "prop_a_value",
            "prop_b": "prop_b_value"
        }
    }
}
```

**全量数据示例（工作项关联页面）：**

```json
{
    "id": "6552d9da1bffb252b8276cf7",
    "url": "https://rest_api_root/v1/relations/6552d9da1bffb252b8276cf7",
    "principal_type": "work_item",
    "principal": {
        "id": "5edca524cad2fa1125cb0630",
        "url": "https://rest_api_root/v1/project/work_items/5edca524cad2fa1125cb0630",
        "identifier": "SCR-5",
        "title": "这是一个缺陷",
        "type": "bug",
        "start_at": 1583290309,
        "end_at": 1583290347,
        "parent_id": "5edca524cad2fa112b06105c",
        "short_id": "c9WqLmTO",
        "html_url": "https://yctech.pingcode.com/pjm/workitems/c9WqLmTO",
        "properties": {
            "prop_a": "prop_a_value",
            "prop_b": "prop_b_value"
        }
    },
    "target_type": "page",
    "target": {
        "id": "63e1bf51760505c8795ebccc",
        "url": "https://rest_api_root/v1/wiki/pages/63e1bf51760505c8795ebccc",
        "name": "示例页面",
        "type": "document",
        "short_id": "5-x6NN",
        "html_url": "https://yctech.pingcode.com/wiki/pages/5-x6NN"
    }
}
```

**引用数据示例:**

```json
{
    "id": "fa1125cb06305edca524cad2",
    "url": "https://rest_api_root/v1/relations/fa1125cb06305edca524cad2"
}
```

---

### 工时

用于查看和管理工时的基本信息。 资源地址：{GET} https://rest_api_root/v1/workloads/{workload_id}

**全量数据示例：**

```json
{
    "id": "5f168f764eba01a5278b87cd",
    "url": "https://rest_api_root/v1/workloads/5f168f764eba01a5278b87cd",
    "principal_type": "work_item",
    "principal": {
        "id": "5edca524cad2fa1125cb0630",
        "url": "https://rest_api_root/v1/project/work_items/5edca524cad2fa1125cb0630",
        "identifier": "SCR-5",
        "title": "这是一个缺陷",
        "type": "bug",
        "start_at": 1583290309,
        "end_at": 1583290347,
        "parent_id": "5edca524cad2fa1125cb0629",
        "short_id": "c9WqLmTO",
        "html_url": "https://yctech.pingcode.com/pjm/workitems/c9WqLmTO",
        "properties": {
            "prop_a": "prop_a_value",
            "prop_b": "prop_b_value"
        }
    },
    "type": {
        "id": "5a86eaf6a72585327ea46fge0",
        "url": "https://rest_api_root/v1/workload_types/5a86eaf6a72585327ea46fge0",
        "name": "研发"
    },
    "duration": 8,
    "review_state": "approved",
    "description": "这是一个工时",
    "report_at": 1593290347,
    "report_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "created_at": 1593290347,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    }
}
```

**引用数据示例：**

```json
{
    "id": "5f168f764eba01a5278b87cd",
    "url": "https://rest_api_root/v1/workloads/5f168f764eba01a5278b87cd",
    "type": {
        "id": "5a86eaf6a72585327ea46fge0",
        "url": "https://rest_api_root/v1/workload_types/5a86eaf6a72585327ea46fge0",
        "name": "研发"
    },
    "duration": 8,
    "description": "这是一个工时",
    "report_at": 1593290347,
    "report_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    }
}
```

---

### 活动记录

用于查看和管理活动记录的基本信息。 资源地址：{GET} https://rest_api_root/v1/activitys/{activity_id}?principal_type={principal_type}&amp;principal_id={principal_id}

**全量数据示例：**

```json
{
    "id": "694ae20fdb8e0baef70f7ddb",
    "url": "https://rest_api_root/v1/activities/694ae20fdb8e0baef70f7ddb?principal_type=idea&principal_id=683562430d684517b06b814b",
    "template": "update-property",
    "type": "update",
    "summary": "修改了引用多选",
    "content": {
        "property_key": "yinyongduoxuan",
        "origin": null,
        "target": [
            {
                "id": "65fa797d8f0358d376233220",
                "name": "REST API 产品"
            }
        ]
    },
    "client": "web",
    "created_at": 1766515215,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    }
}
```

**引用数据示例：**

```json
{
    "id": "694ae20fdb8e0baef70f7ddb",
    "url": "https://rest_api_root/v1/activities/694ae20fdb8e0baef70f7ddb?principal_type=idea&principal_id=683562430d684517b06b814b",
    "summary": "修改了引用多选"
}
```

---

### 评审

用于查看和管理评审的基本信息。 资源地址：{GET} https://rest_api_root/v1/reviews/{review_id}?principal_type={principal_type}

**全量数据示例：**

```json
{
    "id": "6f168f764eba01a5278b87cd",
    "url": "https://rest_api_root/v1/reviews/6f168f764eba01a5278b87cd?principal_type=idea",
    "identifier": "SCR-R5",
    "title": "这是一个评审",
    "status": "completed",
    "principal_type": "idea",
    "short_id": "LsEy8mZF",
    "html_url": "https://yctech.pingcode.com/reviews/LsEy8mZF",
    "pilot": {
        "id": "63bb744314bd13c9def24cb4",
        "url": "https://rest_api_root/v1/ship/products/63bb744314bd13c9def24cb4",
        "name": "示例产品",
        "identifier": "SLC",
        "is_archived": 0,
        "is_deleted": 0
    },
    "description": "这是一个评审的描述",
    "participants": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "url": "https://rest_api_root/v1/participants/a0417f68e846aae315c85d24643678a9?principal_type=idea&review_id=6f168f764eba01a5278b87cd",
            "type": "user",
            "user": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            }
        }
    ],
    "submitted_at": 1593290347,
    "submitted_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/avatar.png"
    },
    "completed_at": 1593291347,
    "completed_by": {
        "id": "b2417f68e846aae315c85d24643678b0",
        "url": "https://rest_api_root/v1/directory/users/b2417f68e846aae315c85d24643678b0",
        "name": "mary",
        "display_name": "Mary",
        "avatar": "https://s3.amazonaws.com/bucket/avatar2.png"
    },
    "created_at": 1593290347,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/avatar.png"
    },
    "updated_at": 1593291347,
    "updated_by": {
        "id": "b2417f68e846aae315c85d24643678b0",
        "url": "https://rest_api_root/v1/directory/users/b2417f68e846aae315c85d24643678b0",
        "name": "mary",
        "display_name": "Mary",
        "avatar": "https://s3.amazonaws.com/bucket/avatar2.png"
    }
}
```

**引用数据示例：**

```json
{
    "id": "6f168f764eba01a5278b87cd",
    "url": "https://rest_api_root/v1/reviews/6f168f764eba01a5278b87cd?principal_type=idea",
    "identifier": "SCR-R5",
    "title": "这是一个评审",
    "status": "completed",
    "principal_type": "idea",
    "short_id": "LsEy8mZF",
    "html_url": "https://yctech.pingcode.com/reviews/LsEy8mZF"
}
```

---

### 评论

用于查看和管理评论的基本信息。 资源地址：{GET} https://rest_api_root/v1/comments/{comment_id}

**全量数据示例：**

```json
{
    "id": "59f72dfaeadb5b5197b7da6d",
    "url": "https://rest_api_root/v1/comments/59f72dfaeadb5b5197b7da6d?principal_type=work_item&principal_id=5edca524cad2fa1125cb0630",
    "content": "这是一个工作项评论",
    "attachment_count": 1,
    "attachments": [
        {
            "id": "5da588ca84c7377a5d327e6d",
            "url": "https://rest_api_root/v1/attachments/5da588ca84c7377a5d327e6d?principal_type=work_item&principal_id=5edca524cad2fa1125cb0630&comment_id=59f72dfaeadb5b5197b7da6d",
            "title": "这是一个代码片段",
            "size": 16,
            "type": "snippet"
        }
    ],
    "is_reply_comment": false,
    "replied_comment": null,
    "created_at": 1565255712,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "is_deleted": 0
}
```

**引用数据示例：**

```json
{
    "id": "59f72dfaeadb5b5197b7da6d",
    "url": "https://rest_api_root/v1/comments/59f72dfaeadb5b5197b7da6d?principal_type=work_item&principal_id=5edca524cad2fa1125cb0630",
    "content": "这是一个工作项评论",
    "is_deleted": 0
}
```

---

### 附件

用于查看和管理附件的基本信息。 资源地址：{GET} https://rest_api_root/v1/attachments/{attachment_id}

**全量数据示例（文件）:**

```json
{
    "id": "5da588ca84c7377a5d327e6c",
    "url": "https://rest_api_root/v1/attachments/5da588ca84c7377a5d327e6c?principal_type=work_item&principal_id=5edca524cad2fa1125cb0630",
    "title": "这是一个图片类型的附件",
    "size": 1024,
    "type": "file",
    "file_type": "image",
    "ext": "png",
    "download_url": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a.png",
    "created_at": 1583290347,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    }
}
```

**全量数据示例（代码段）:**

```json
{
    "id": "5da588ca84c7377a5d327e6d",
    "url": "https://rest_api_root/v1/attachments/5da588ca84c7377a5d327e6d?principal_type=work_item&principal_id=5edca524cad2fa1125cb0630&comment_id=59f72dfaeadb5b5197b7da6d",
    "title": "这是一个代码片段",
    "size": 16,
    "type": "snippet",
    "format": "javascript",
    "content": "const a = 'abc';",
    "line": 1,
    "created_at": 1583290347,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    }
}
```

**引用数据示例：**

```json
{
    "id": "5da588ca84c7377a5d327e6d",
    "url": "https://rest_api_root/v1/attachments/5da588ca84c7377a5d327e6d?principal_type=work_item&principal_id=5edca524cad2fa1125cb0630&comment_id=59f72dfaeadb5b5197b7da6d",
    "title": "这是一个代码片段",
    "size": 16,
    "type": "snippet"
}
```

---

## 部署

### 全量更新一个部署

**`PUT /v1/release/deploys/{deploy_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `deploy_id` | String | 部署的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `status` | String | 部署的状态。 |
| `env_id` | String | 环境的id。 |
| `release_name` | String | 版本发布的名称。 |
| `release_url` *(可选)* | String | 版本发布的地址。如果为空，在PingCode中不显示相关跳转链接。 |
| `start_at` | Number | 部署开始的时间。 |
| `end_at` | Number | 部署结束的时间。 |
| `duration` | Number | 部署持续的时间。单位为秒。 |
| `work_item_identifiers` *(可选)* | String[] | PingCode工作项编号的列表。通过该参数将部署与一个或多个工作项关联。 |

**响应示例：**

```json
{
    "id": "564587fe700d43b81b080339",
    "url": "https://rest_api_root/v1/release/deploys/564587fe700d43b81b080339",
    "status": "deployed",
    "release_name": "1.1.0",
    "environment": {
        "id": "564587fe700d43b81b080123",
        "url": "https://rest_api_root/v1/release/environments/564587fe700d43b81b080123",
        "name": "Production"
    },
    "release_url": "https://your-release-host/release",
    "start_at": 1583143467,
    "end_at": 1583143667,
    "duration": 200,
    "work_items": [
        {
            "id": "564587fe700d43b81b080ab8",
            "url": "https://rest_api_root/v1/project/work_items/564587fe700d43b81b080ab8",
            "identifier": "PLM-001",
            "title": "这是一个用户故事",
            "type": "story",
            "start_at": 1583290309,
            "end_at": 1583290347,
            "parent_id": "5edca524cad2fa112b06105c",
            "short_id": "c9WqLmTO",
            "html_url": "https://yctech.pingcode.com/pjm/workitems/c9WqLmTO",
            "properties": {
                "prop_a": "prop_a_value",
                "prop_b": "prop_b_value"
            }
        }
    ]
}
```

---

### 创建一个部署

**`POST /v1/release/deploys`**

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `status` | String | 部署的状态。 |
| `env_id` | String | 环境的id。 |
| `release_name` | String | 发布的名称。 |
| `release_url` *(可选)* | String | 版本发布的地址。如果为空，在PingCode中不显示相关跳转链接。 |
| `start_at` | Number | 部署开始的时间。 |
| `end_at` | Number | 部署结束的时间。 |
| `duration` | Number | 部署持续的时间。单位为秒。 |
| `work_item_identifiers` *(可选)* | String[] | PingCode工作项编号的列表。通过该参数将部署与一个或多个工作项关联。 |

**响应示例：**

```json
{
    "id": "564587fe700d43b81b080339",
    "url": "https://rest_api_root/v1/release/deploys/564587fe700d43b81b080339",
    "status": "deployed",
    "release_name": "1.1.0",
    "environment": {
        "id": "564587fe700d43b81b080123",
        "url": "https://rest_api_root/v1/release/environments/564587fe700d43b81b080123",
        "name": "Production"
    },
    "release_url": "https://github.com/worktile/ngx-planet/releases/tag/1.1.0",
    "start_at": 1583143467,
    "end_at": 1583143667,
    "duration": 200,
    "work_items": [
        {
            "id": "564587fe700d43b81b080ab8",
            "url": "https://rest_api_root/v1/project/work_items/564587fe700d43b81b080ab8",
            "identifier": "PLM-001",
            "title": "这是一个用户故事",
            "type": "story",
            "start_at": 1583290309,
            "end_at": 1583290347,
            "parent_id": "5edca524cad2fa112b06105c",
            "short_id": "c9WqLmTO",
            "html_url": "https://yctech.pingcode.com/pjm/workitems/c9WqLmTO",
            "properties": {
                "prop_a": "prop_a_value",
                "prop_b": "prop_b_value"
            }
        }
    ]
}
```

---

### 删除一个部署

**`DELETE /v1/release/deploys/{deploy_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `deploy_id` | String | 部署的id。 |

**响应示例：**

```json
{
    "id": "564587fe700d43b81b080339",
    "url": "https://rest_api_root/v1/release/deploys/564587fe700d43b81b080339",
    "status": "deployed",
    "release_name": "1.1.0",
    "environment": {
        "id": "564587fe700d43b81b080123",
        "url": "https://rest_api_root/v1/release/environments/564587fe700d43b81b080123",
        "name": "Production"
    },
    "release_url": "https://github.com/worktile/ngx-planet/releases/tag/1.1.0",
    "start_at": 1583143467,
    "end_at": 1583143667,
    "duration": 200,
    "work_items": [
        {
            "id": "564587fe700d43b81b080ab8",
            "url": "https://rest_api_root/v1/project/work_items/564587fe700d43b81b080ab8",
            "identifier": "PLM-001",
            "title": "这是一个用户故事",
            "type": "story",
            "start_at": 1583290309,
            "end_at": 1583290347,
            "parent_id": "5edca524cad2fa112b06105c",
            "short_id": "c9WqLmTO",
            "html_url": "https://yctech.pingcode.com/pjm/workitems/c9WqLmTO",
            "properties": {
                "prop_a": "prop_a_value",
                "prop_b": "prop_b_value"
            }
        }
    ]
}
```

---

### 获取部署列表

**`GET /v1/release/deploys`**

**查询参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `env_id` *(可选)* | String | 环境的id。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "564587fe700d43b81b080339",
            "url": "https://rest_api_root/v1/release/deploys/564587fe700d43b81b080339",
            "status": "deployed",
            "release_name": "1.1.0",
            "environment": {
                "id": "564587fe700d43b81b080123",
                "url": "https://rest_api_root/v1/release/environments/564587fe700d43b81b080123",
                "name": "Production"
            },
            "release_url": "https://github.com/worktile/ngx-planet/releases/tag/1.1.0",
            "start_at": 1583143467,
            "end_at": 1583143667,
            "duration": 200,
            "work_items": [
                {
                    "id": "564587fe700d43b81b080ab8",
                    "url": "https://rest_api_root/v1/project/work_items/564587fe700d43b81b080ab8",
                    "identifier": "PLM-001",
                    "title": "这是一个用户故事",
                    "type": "story",
                    "start_at": 1583290309,
                    "end_at": 1583290347,
                    "parent_id": "5edca524cad2fa112b06105c",
                    "short_id": "c9WqLmTO",
                    "html_url": "https://yctech.pingcode.com/pjm/workitems/c9WqLmTO",
                    "properties": {
                        "prop_a": "prop_a_value",
                        "prop_b": "prop_b_value"
                    }
                }
            ]
        }
    ]
}
```

---

### 部分更新一个部署

**`PATCH /v1/release/deploys/{deploy_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `deploy_id` | String | 部署的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `status` *(可选)* | String | 部署的状态。 |
| `env_id` *(可选)* | String | 环境的id。 |
| `release_name` *(可选)* | String | 版本发布的名称。 |
| `release_url` *(可选)* | String | 版本发布的地址。如果为空，在PingCode中不显示相关跳转链接。 |
| `start_at` *(可选)* | Number | 部署开始的时间。 |
| `end_at` *(可选)* | Number | 部署结束的时间。 |
| `duration` *(可选)* | Number | 部署持续的时间。单位为秒。 |
| `work_item_identifiers` *(可选)* | String[] | PingCode工作项编号的列表。通过该参数将部署与一个或多个工作项关联。 |

**响应示例：**

```json
{
    "id": "564587fe700d43b81b080339",
    "url": "https://rest_api_root/v1/release/deploys/564587fe700d43b81b080339",
    "status": "deployed",
    "release_name": "1.1.0",
    "environment": {
        "id": "564587fe700d43b81b080123",
        "url": "https://rest_api_root/v1/release/environments/564587fe700d43b81b080123",
        "name": "Production"
    },
    "release_url": "https://github.com/worktile/ngx-planet/releases/tag/1.1.0",
    "start_at": 1583143467,
    "end_at": 1583143667,
    "duration": 200,
    "work_items": [
        {
            "id": "564587fe700d43b81b080ab8",
            "url": "https://rest_api_root/v1/project/work_items/564587fe700d43b81b080ab8",
            "identifier": "PLM-001",
            "title": "这是一个用户故事",
            "type": "story",
            "start_at": 1583290309,
            "end_at": 1583290347,
            "parent_id": "5edca524cad2fa112b06105c",
            "short_id": "c9WqLmTO",
            "html_url": "https://yctech.pingcode.com/pjm/workitems/c9WqLmTO",
            "properties": {
                "prop_a": "prop_a_value",
                "prop_b": "prop_b_value"
            }
        }
    ]
}
```

---

## 部门

### 创建一个部门

**`POST /v1/directory/departments`**

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `name` | String | 部门的名称，在一个企业中这个名称是唯一的。 |
| `parent_id` *(可选)* | String | 父部门的id。 |
| `head_id` *(可选)* | String | 部门负责人的id。 |

**响应示例：**

```json
{
    "id": "6422711c3f12e6c1e46d40e6",
    "url": "https://rest_api_root/v1/directory/departments/6422711c3f12e6c1e46d40e6",
    "name": "技术支持",
    "head": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "parent": {
        "id": "6422711c3f12e6c1e46d40e2",
        "url": "https://rest_api_root/v1/directory/departments/6422711c3f12e6c1e46d40e2",
        "name": "技术中心"
    }
}
```

---

### 删除一个部门

**`DELETE /v1/directory/departments/{department_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `department_id` | String | 部门id。 |

**响应示例：**

```json
{
    "id": "6422711c3f12e6c1e46d40e6",
    "url": "https://rest_api_root/v1/directory/departments/6422711c3f12e6c1e46d40e6",
    "name": "技术支持",
    "head": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "parent": {
        "id": "6422711c3f12e6c1e46d40e2",
        "url": "https://rest_api_root/v1/directory/departments/6422711c3f12e6c1e46d40e2",
        "name": "技术中心"
    }
}
```

---

### 获取部门列表

**`GET /v1/directory/departments`**

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "6422711c3f12e6c1e46d40e6",
            "url": "https://rest_api_root/v1/directory/departments/6422711c3f12e6c1e46d40e6",
            "name": "技术支持",
            "head": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "parent": {
                "id": "6422711c3f12e6c1e46d40e2",
                "url": "https://rest_api_root/v1/directory/departments/6422711c3f12e6c1e46d40e2",
                "name": "技术中心"
            }
        }
    ]
}
```

---

### 部分更新一个部门

**`PATCH /v1/directory/departments/{department_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `department_id` | String | 部门id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `name` *(可选)* | String | 部门的名称，在一个企业中这个名称是唯一的。 |
| `parent_id` *(可选)* | String | 父部门的id。 |
| `head_id` *(可选)* | String | 部门负责人的id。 |

**响应示例：**

```json
{
    "id": "6422711c3f12e6c1e46d40e6",
    "url": "https://rest_api_root/v1/directory/departments/6422711c3f12e6c1e46d40e6",
    "name": "技术支持",
    "head": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "parent": {
        "id": "6422711c3f12e6c1e46d40e2",
        "url": "https://rest_api_root/v1/directory/departments/6422711c3f12e6c1e46d40e2",
        "name": "技术中心"
    }
}
```

---

## 鉴权

### 客户端凭据

客户端凭据（OAuth2 Client Credentials）是最简单、最直接的授权方式，通过客户端凭据获取的访问令牌（access_token）不区分用户身份， 在PingCode中将这类访问令牌称为“企业令牌”，企业令牌具有系统管理员的权限，主要用于访问和操作全局的数据，请谨慎管理。 获取企业令牌时，需要提供client_id和client_secret， 您可以在PingCode企业后台的凭据管理中创建一个应用，配置数据范围，然后拿到client_id和client_secret。 使用企业令牌时，只需要在HTTP请求在headers中添加&quot;authorization&quot;: &quot;Bearer {access_token}&quot;，即可访问受保护的数据。

---

### 授权码

授权码（OAuth2 Authorization Code）是最常用的授权方式，主要用于企业员工管理自己的数据。 通过授权码获取的访问令牌（access_token）在PingCode中称为“用户令牌”，用户令牌属于某个用户所有，仅能访问这个用户有权限操作的数据。 第三方应用可以通过用户的手动授权获得该用户的用户令牌，然后通过用户令牌访问该用户有权限操作的数据。 在通过授权码的方式获取用户令牌时，需要提供client_id和code， 您可以在PingCode企业后台的凭据管理中创建一个应用，配置数据范围，然后拿到client_id和redirect_uri。 在您的第三方应用中引导用户访问请求授权页面，PingCode会询问该用户是否授权给您的应用。 得到用户许可后，浏览器会跳转redirect_uri页面，并在URL的参数中包含一个临时的code，然后您的应用可以根据client_id和code获取该用户的用户令牌。 使用用户令牌时，只需要在HTTP请求在headers中添加&quot;authorization&quot;: &quot;Bearer {access_token}&quot;，即可访问受保护的数据。

---

## 附件

### 上传一个代码段

**`POST /v1/attachments`**

**请求参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `principal_type` | String | 附件主体的类型。 |
| `principal_id` | String | 附件主体的id。 |
| `comment_id` *(可选)* | String | 附件主体的评论id。当需要向附件主体的评论上传文件或者代码段时，需要传入该参数。 |
| `title` | String | 代码段的标题。 |
| `format` | String | 代码段的语言。 |
| `content` | String | 代码段的内容。 |

---

### 上传一个文件

**`POST /v1/attachments?principal_type={principal_type}&principal_id={principal_id}`**

**查询参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `principal_type` | String | 附件主体的类型。 |
| `principal_id` | String | 附件主体的id。 |
| `comment_id` *(可选)* | String | 附件主体的评论id。当需要向附件主体的评论上传文件或者代码段时，需要传入该参数。 |

**请求参数 form-data：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `title` | String | 这是一个图片类型的附件.png |
| `file` | File | /Users/ping-code/这是一个图片类型的附件.png |

---

### 删除一个附件

**`DELETE /v1/attachments/{attachment_id}?principal_type={principal_type}&principal_id={principal_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `attachment_id` | String | 附件的id。 |

**查询参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `principal_type` | String | 附件主体的类型。 |
| `principal_id` | String | 附件主体的id。 |
| `comment_id` *(可选)* | String | 附件主体的评论id。当需要获取附件主体的评论附件时，需要传入该参数。 |

**响应示例：**

```json
{
    "id": "5da588ca84c7377a5d327e6c",
    "url": "https://rest_api_root/v1/attachments/5da588ca84c7377a5d327e6c?principal_type=work_item&principal_id=5edca524cad2fa1125cb0630",
    "title": "这是一个图片类型的附件",
    "size": 1024,
    "type": "file",
    "file_type": "image",
    "ext": "png",
    "download_url": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a.png",
    "created_at": 1583290347,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    }
}
```

---

### 获取附件列表

**`GET /v1/attachments?principal_type={principal_type}&principal_id={principal_id}`**

**查询参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `principal_type` | String | 附件主体的类型。 |
| `principal_id` | String | 附件主体的id。 |
| `comment_id` *(可选)* | String | 附件主体的评论id。当需要获取附件主体的评论附件时，需要传入该参数。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 2,
    "values": [
        {
            "id": "5da588ca84c7377a5d327e6d",
            "url": "https://rest_api_root/v1/attachments/5da588ca84c7377a5d327e6d?principal_type=work_item&principal_id=5edca524cad2fa1125cb0630&comment_id=59f72dfaeadb5b5197b7da6d",
            "title": "这是一个代码片段",
            "size": 16,
            "type": "snippet",
            "format": "javascript",
            "content": "const a = 'abc';",
            "line": 1,
            "created_at": 1583290347,
            "created_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            }
        },
        {
            "id": "5da588ca84c7377a5d327e6f",
            "url": "https://rest_api_root/v1/attachments/5da588ca84c7377a5d327e6f?principal_type=work_item&principal_id=5edca524cad2fa1125cb0630&comment_id=59f72dfaeadb5b5197b7da6d",
            "title": "这是一个图片类型的附件",
            "size": 1024,
            "type": "file",
            "file_type": "image",
            "ext": "png",
            "download_url": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839b.png",
            "created_at": 1583290347,
            "created_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            }
        }
    ]
}
```

---

## 需求

### 创建一个需求

**`POST /v1/ship/ideas`**

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `product_id` | String | 需求的产品id。 |
| `title` | String | 需求的标题，最大长度为255。 |
| `assignee_id` *(可选)* | String | 需求负责人的id。 |
| `description` *(可选)* | String | 需求的描述。 |
| `suite_id` *(可选)* | String | 需求的产品模块id。 |
| `properties` *(可选)* | Object | 需求属性的键值对集合。要注意的是，当前产品的需求属性视图需要包含这些需求属性，例如需求属性视图中包含prop_a和prop_b。 |
| `properties.prop_a` *(可选)* | Object | 需求项属性prop_a。 |
| `properties.prop_b` *(可选)* | Object | 需求项属性prop_b。 |

**响应示例：**

```json
{
    "id": "64b4d70ba368e6594360ea24",
    "url": "https://rest_api_root/v1/ship/ideas/64b4d70ba368e6594360ea24",
    "product": {
        "id": "6422711c3f12e6c1e46d40e9",
        "url": "http://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9",
        "identifier": "SLC",
        "name": "示例产品",
        "is_archived": 0,
        "is_deleted": 0
    },
    "identifier": "SLC-1",
    "title": "示例需求",
    "short_id": "Ogf1EYey",
    "html_url": "https://yctech.pingcode.com/ship/ideas/Ogf1EYey"
    "assignee": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "state": {
        "id": "63e1bf51898a0be5a2d21b2a",
        "url": "https://rest_api_root/v1/ship/idea_states/63e1bf51898a0be5a2d21b2a",
        "name": "待评审",
        "type": "pending"
    },
    "priority": {
        "id": "5cb9466afda1ce4ca0090001",
        "url": "https://rest_api_root/v1/ship/idea_priorities/5cb9466afda1ce4ca0090001",
        "name": "P4"
    },
    "plan": null,
    "suite": {
        "id": "63bb744414bd13c9def24ce4",
        "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9/suites/63bb744414bd13c9def24ce4",
        "name": "需求模块",
        "type": "module"
     },
    "plan_at": null,
    "real_at": null,
    "score": 0,
    "progress": 0,
    "description": "这是一段描述",
    "properties": {
        "backlog_from": "5cb7e6e2fda1ce4ca0000001",
        "backlog_type": "5cb7e763fda1ce4ca0010002",
        "prop_a": "prop_a_value",
        "prop_b": "prop_b_value"
    },
    "participants": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "url": "https://rest_api_root/v1/participants/a0417f68e846aae315c85d24643678a9?principal_type=idea&principal_id=64b4d70ba368e6594360ea24",
            "type": "user",
            "user": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            }
        }
    ],
    "completed_at": null,
    "completed_by": null,
    "created_at": 1689573131,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "updated_at": 1689573131,
    "updated_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "is_archived": 0,
    "is_deleted": 0
}
```

---

### 获取需求优先级列表

**`GET /v1/ship/idea/priorities?product_id={product_id}`**

**查询参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `product_id` | String | 产品的id。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "5cb9466afda1ce4ca0090005",
            "url": "https://rest_api_root/v1/ship/idea_priorities/5cb9466afda1ce4ca0090005",
            "name": "P0"
        }
    ]
}
```

---

### 获取需求列表

**`GET /v1/ship/ideas`**

**查询参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `product_id` *(可选)* | String | 产品的id。 |
| `state_id` *(可选)* | String | 需求状态id。 |
| `priority_id` *(可选)* | String | 需求优先级id。 |
| `created_between` *(可选)* | String | 创建时间介于的时间范围，通过','分割起始时间。 |
| `updated_between` *(可选)* | String | 更新时间介于的时间范围，通过','分割起始时间。 |
| `keywords` *(可选)* | String | 搜索关键字。支持需求编号和需求标题。 |
| `include_public_image_token` *(可选)* | String | 包含获取图片资源token的属性。使用','分割，最多只能20个，支持description和自定义多行文本类型的属性。参数示例description,properties.xxx。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "64b4d70ba368e6594360ea24",
            "url": "https://rest_api_root/v1/ship/ideas/64b4d70ba368e6594360ea24",
            "product": {
                "id": "6422711c3f12e6c1e46d40e9",
                "url": "http://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9",
                "identifier": "SLC",
                "name": "示例产品",
                "is_archived": 0,
                "is_deleted": 0
            },
            "identifier": "SLC-1",
            "title": "示例需求",
            "short_id": "Ogf1EYey",
            "html_url": "https://yctech.pingcode.com/ship/ideas/Ogf1EYey",
            "assignee": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "state": {
                "id": "63e1bf51898a0be5a2d21b2a",
                "url": "https://rest_api_root/v1/ship/idea_states/63e1bf51898a0be5a2d21b2a",
                "name": "待评审",
                "type": "pending"
            },
            "priority": {
                "id": "5cb9466afda1ce4ca0090005",
                "url": "https://rest_api_root/v1/ship/idea_priorities/5cb9466afda1ce4ca0090005",
                "name": "P0"
            },
            "plan": {
                "id": "63bb744414bd13c9def24ce4",
                "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9/plans/63bb744414bd13c9def24ce4",
                "name": "账号管理"
            },
            "suite": {
                "id": "63bb744414bd13c9def24ce4",
                "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9/suites/63bb744414bd13c9def24ce4",
                "name": "需求模块",
                "type": "module"
            },
            "plan_at": {
                "from": 1690732800,
                "to": 1691337599,
                "granularity": "day"
            },
            "real_at": {
                "from": 1690732800,
                "to": 1691337599,
                "granularity": "day"
            },
            "score": 0,
            "progress": 0.60,
            "description": "这是一段描述",
            "properties": {
                "backlog_from": "5cb7e6e2fda1ce4ca0000001",
                "backlog_type": "5cb7e763fda1ce4ca0010002"
            },
            "public_image_token": "-fkvANQ2dcVECK6Xg45L3kG8VCbOTK8NrNckGkxljQd",
            "participants": [
                {
                    "id": "a0417f68e846aae315c85d24643678a9",
                    "url": "https://rest_api_root/v1/participants/a0417f68e846aae315c85d24643678a9?principal_type=idea&principal_id=64b4d70ba368e6594360ea24",
                    "type": "user",
                    "user": {
                        "id": "a0417f68e846aae315c85d24643678a9",
                        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                        "name": "john",
                        "display_name": "John",
                        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
                    }
                },
                {
                    "id": "63c8fb32729dee3334d96af7",
                    "url": "https://rest_api_root/v1/participants/63c8fb32729dee3334d96af7?principal_type=idea&principal_id=64b4d70ba368e6594360ea24",
                    "type": "user_group",
                    "user_group": {
                        "id": "63c8fb32729dee3334d96af7",
                        "url": "https://rest_api_root/v1/directory/groups/63c8fb32729dee3334d96af7",
                        "name": "Open Team"
                    }
                }
            ],
           "completed_at": 1689573131,
           "completed_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "created_at": 1689573131,
            "created_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "updated_at": 1689579131,
            "updated_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "is_archived": 0,
            "is_deleted": 0
        }
    ]
}
```

---

### 获取需求属性列表

**`GET /v1/ship/idea/properties?product_id={product_id}`**

**查询参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `product_id` | String | 产品的id。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 2,
    "values": [
        {
            "id": "backlog_type",
            "url": "https://rest_api_root/v1/ship/idea_properties/backlog_type",
            "name": "需求类型",
            "type": "select",
            "options": [
                {
                    "_id": "5cb7e763fda1ce4ca0010002",
                    "text": "功能需求"
                },
                {
                    "_id": "5cb7e763fda1ce4ca0010004",
                    "text": "体验优化"
                }
            ]
        },
        {
            "id": "identifier",
            "url": "https://rest_api_root/v1/ship/idea_properties/identifier",
            "name": "编号",
            "type": "text",
            "options": null
        }
    ]
}
```

---

### 获取需求排期列表

**`GET /v1/ship/idea/plans?product_id={product_id}`**

**查询参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `product_id` | String | 产品的id。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "63bb744414bd13c9def24ce4",
            "url": "http://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e6/plans/63bb744414bd13c9def24ce4",
            "name": "账号管理"
        }
    ]
}
```

---

### 获取需求模块列表

**`GET /v1/ship/idea/suites?product_id={product_id}`**

**查询参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `product_id` | String | 产品的id。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "63bb744414bd13c9def24ce4",
            "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9/suites/63bb744414bd13c9def24ce4",
            "name": "需求模块",
            "type": "module"
        }
    ]
}
```

---

### 获取需求流转记录列表

**`GET /v1/ship/ideas/{idea_id}/transition_histories`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `idea_id` | String | 需求的id。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 2,
    "values": [
        {
            "id": "64c3676c983bb9481ee1eea5",
            "url": "https://rest_api_root/v1/ship/ideas/64b4d70ba368e6594360ea24/transition_histories/64c3676c983bb9481ee1eea5",
            "idea": {
                "id": "64b4d70ba368e6594360ea24",
                "url": "https://rest_api_root/v1/ship/ideas/64b4d70ba368e6594360ea24",
                "identifier": "SLC-1",
                "title": "示例需求",
                "short_id": "Ogf1EYey",
                "html_url": "https://yctech.pingcode.com/ship/ideas/Ogf1EYey"
            },
            "from_state": null,
            "to_state": {
                "id": "63e1bf51898a0be5a2d21b29",
                "url": "https://rest_api_root/v1/ship/idea_states/63e1bf51898a0be5a2d21b29",
                "name": "待处理",
                "type": "pending"
            },
            "created_at": 1674528614,
            "created_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            }
        },
        {
            "id": "658bdb79e5839f556562accf",
            "url": "https://rest_api_root/v1/ship/ideas/64b4d70ba368e6594360ea24/transition_histories/658bdb79e5839f556562accf",
            "idea": {
                "id": "64b4d70ba368e6594360ea24",
                "url": "https://rest_api_root/v1/ship/ideas/64b4d70ba368e6594360ea24",
                "identifier": "SLC-1",
                "title": "示例需求",
                "short_id": "Ogf1EYey",
                "html_url": "https://yctech.pingcode.com/ship/ideas/Ogf1EYey"
            },
            "from_state": {
                "id": "63e1bf51898a0be5a2d21b29",
                "url": "https://rest_api_root/v1/ship/idea_states/63e1bf51898a0be5a2d21b29",
                "name": "待处理",
                "type": "pending"
            },
            "to_state": {
                "id": "63e1bf51898a0be5a2d21b2b",
                "url": "https://rest_api_root/v1/ship/idea_states/63e1bf51898a0be5a2d21b2b",
                "name": "处理中",
                "type": "in_progress"
            },
            "created_at": 1674528614,
            "created_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            }
        }
    ]
}
```

---

### 获取需求状态列表

**`GET /v1/ship/idea/states?product_id={product_id}`**

**查询参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `product_id` | String | 产品的id。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "63e1bf51898a0be5a2d21b2a",
            "url": "https://rest_api_root/v1/ship/idea_states/63e1bf51898a0be5a2d21b2a",
            "name": "待评审",
            "type": "pending"
        }
    ]
}
```

---

### 部分更新一个需求

**`PATCH /v1/ship/ideas/{idea_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `idea_id` | String | 需求id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `title` *(可选)* | String | 需求的标题，最大长度为255。 |
| `description` *(可选)* | String | 需求的描述。 |
| `state_id` *(可选)* | String | 需求状态的id，您可以在 获取需求状态列表 API获取。 |
| `priority_id` *(可选)* | String | 需求优先级的id，您可以在 获取需求优先级列表 API获取。 |
| `assignee_id` *(可选)* | String | 需求负责人的id。 |
| `progress` *(可选)* | Number | 需求的进度，取值范围为：0到1的两位小数。 |
| `plan_at` *(可选)* | Object | 需求的计划时间。plan_at是整体更新的，其中包含from、to、granularity三个属性，均为必填。 |
| `plan_at.from` | Number | 需求的计划开始时间。为秒级时间戳。 |
| `plan_at.to` | Number | 需求的计划结束时间。为秒级时间戳。 |
| `plan_at.granularity` | String | 需求的计划时间周期单位。 |
| `real_at` *(可选)* | Object | 需求的实际时间，参数格式同plan_at。 |
| `plan_id` *(可选)* | String | 产品排期的id，您可以在 获取产品排期列表 API获取。 |
| `suite_id` *(可选)* | String | 产品模块的id，您可以在 获取产品模块列表 API获取。 |
| `properties` *(可选)* | Object | 需求的自定义属性。 |
| `properties.prop_a` *(可选)* | Object | 需求的自定义属性prop_a。 |
| `properties.prop_b` *(可选)* | Object | 需求的自定义属性prop_b。 |

**响应示例：**

```json
{
    "id": "64b4d70ba368e6594360ea24",
    "url": "https://rest_api_root/v1/ship/ideas/64b4d70ba368e6594360ea24",
    "product": {
        "id": "6422711c3f12e6c1e46d40e9",
        "url": "http://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9",
        "identifier": "SLC",
        "name": "示例产品",
        "is_archived": 0,
        "is_deleted": 0
    },
    "identifier": "SLC-1",
    "title": "示例需求",
    "short_id": "Ogf1EYey",
    "html_url": "https://yctech.pingcode.com/ship/ideas/Ogf1EYey",
    "assignee": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "state": {
        "id": "63e1bf51898a0be5a2d21b2a",
        "url": "https://rest_api_root/v1/ship/idea_states/63e1bf51898a0be5a2d21b2a",
        "name": "待评审",
        "type": "pending"
    },
    "priority": {
        "id": "5cb9466afda1ce4ca0090005",
        "url": "https://rest_api_root/v1/ship/idea_priorities/5cb9466afda1ce4ca0090005",
        "name": "P0"
    },
    "plan": {
        "id": "63bb744414bd13c9def24ce4",
        "url": "http://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9/plans/63bb744414bd13c9def24ce4",
        "name": "账号管理"
    },
    "suite": {
        "id": "63bb744414bd13c9def24ce4",
        "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9/suites/63bb744414bd13c9def24ce4",
        "name": "需求模块",
        "type": "module"
     },
    "plan_at": {
        "from": 1690732800,
        "to": 1691337599,
        "granularity": "day"
    },
    "real_at": {
        "from": 1690732800,
        "to": 1691337599,
        "granularity": "day"
    },
    "score": 0,
    "progress": 0.88,
    "description": "这是一段描述",
    "properties": {
        "backlog_from": "5cb7e6e2fda1ce4ca0000001",
        "backlog_type": "5cb7e763fda1ce4ca0010002",
        "prop_a": "prop_a_value",
        "prop_b": "prop_b_value"
    },
    "participants": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "url": "https://rest_api_root/v1/participants/a0417f68e846aae315c85d24643678a9?principal_type=idea&principal_id=64b4d70ba368e6594360ea24",
            "type": "user",
            "user": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            }
        }
    ],
    "completed_at": 1689578888,
    "completed_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "created_at": 1689573131,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "updated_at": 1689578888,
    "updated_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "is_archived": 0,
    "is_deleted": 0
}
```

---

## 需求配置

### 创建一个需求属性

**`POST /v1/ship/idea_properties`**

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `name` | String | 需求属性的名称。在一个企业中这个名称是唯一的。 |
| `type` | String | 需求属性的类型。 |
| `options` *(可选)* | Object[] | 选择项列表。当需求属性类型为select,multi_select,cascade_select,cascade_multi_select时可选填此项。 |
| `options._id` *(可选)* | String | 选择项id。该值在选择项中是唯一的。 |
| `options.text` | String | 选择项内容。该值在选择项中是唯一的。 |
| `options.parent_id` *(可选)* | String | 选择项父级id。当属性类型为cascade_select,cascade_multi_select时，parent_id参数有效，用于构建级联类型选择项之间的父子关系，最多存在4级。 |

**响应示例（select）：**

```json
{
    "id": "severity",
    "url": "https://rest_api_root/v1/ship/idea_properties/severity",
    "name": "严重程度",
    "type": "select",
    "options": [
        {
            "_id": "5efb1859110533727a82c603",
            "text": "严重"
        },
        {
            "_id": "5efb1859110533727a82c604",
            "text": "一般"
        }
    ],
    "is_removable": true,
    "is_name_editable": true,
    "is_options_editable": true,
    "select_all_level": false,
    "display_all_level": false,
    "display_separator": null
}
```

**响应示例（cascade_select）：**

```json
{
    "id": "jiliandanxuan",
    "url": "https://rest_api_root/v1/ship/idea_properties/jiliandanxuan",
    "name": "级联单选",
    "type": "cascade_select",
    "options": [
        {
            "_id": "64b9f741eec7fd94e63b411e",
            "text": "父"
        },
        {
            "_id": "64b9f741eec7fd94e63b411f",
            "text": "子",
            "parent_id": "64b9f741eec7fd94e63b411e"
        }
    ],
    "is_removable": true,
    "is_name_editable": true,
    "is_options_editable": true,
    "select_all_level": false,
    "display_all_level": false,
    "display_separator": "/"
}
```

---

### 向需求属性方案中添加一个需求属性

**`POST /v1/ship/idea_property_plans/{property_plan_id}/idea_properties`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `property_plan_id` | String | 需求属性方案的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `property_id` | String | 需求属性的id。 |

**响应示例：**

```json
{
    "id": "solution",
    "url": "https://rest_api_root/v1/ship/idea_property_plans/5d7a21f19ef715269ae90c50/idea_properties/solution",
    "property_plan": {
        "id": "5d7a21f19ef715269ae90c50",
        "url": "https://rest_api_root/v1/ship/idea_property_plans/5d7a21f19ef715269ae90c50"
    },
    "property": {
        "id": "solution",
        "url": "https://rest_api_root/v1/ship/idea_properties/solution",
        "name": "解决方案",
        "type": "select"
    }
}
```

---

### 在需求属性方案中移除一个需求属性

**`DELETE /v1/ship/idea_property_plans/{property_plan_id}/idea_properties/{property_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `property_plan_id` | String | 需求属性方案的id。 |
| `property_id` | String | 需求属性的id。 |

**响应示例：**

```json
{
    "id": "solution",
    "url": "https://rest_api_root/v1/ship/idea_property_plans/5d7a21f19ef715269ae90c50/idea_properties/solution",
    "property_plan": {
        "id": "5d7a21f19ef715269ae90c50",
        "url": "https://rest_api_root/v1/ship/idea_property_plans/5d7a21f19ef715269ae90c50"
    },
    "property": {
        "id": "solution",
        "url": "https://rest_api_root/v1/ship/idea_properties/solution",
        "name": "解决方案",
        "type": "select"
    }
}
```

---

### 获取全部需求属性列表

**`GET /v1/ship/idea_properties`**

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 2,
    "values": [
        {
            "id": "backlog_type",
            "url": "https://rest_api_root/v1/ship/idea_properties/backlog_type",
            "name": "需求类型",
            "type": "select",
            "options": [
                {
                    "_id": "5cb7e763fda1ce4ca0010002",
                    "text": "功能需求"
                },
                {
                    "_id": "5cb7e763fda1ce4ca0010004",
                    "text": "体验优化"
                }
            ],
            "is_removable": false,
            "is_name_editable": false,
            "is_options_editable": false
        },
        {
            "id": "identifier",
            "url": "https://rest_api_root/v1/ship/idea_properties/identifier",
            "name": "编号",
            "type": "text",
            "options": null,
            "is_removable": false,
            "is_name_editable": false,
            "is_options_editable": false
        }
    ]
}
```

---

### 获取全部需求状态列表

**`GET /v1/ship/idea_states`**

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "63e1bf51898a0be5a2d21b2a",
            "url": "https://rest_api_root/v1/ship/idea_states/63e1bf51898a0be5a2d21b2a",
            "name": "待评审",
            "type": "pending",
            "color": "#56ABFB"
        }
    ]
}
```

---

### 获取需求属性方案中的需求属性列表

**`GET /v1/ship/idea_property_plans/{property_plan_id}/idea_properties`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `property_plan_id` | String | 需求属性方案的id。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "solution",
            "url": "https://rest_api_root/v1/ship/idea_property_plans/5d7a21f19ef715269ae90c50/idea_properties/solution",
            "property_plan": {
                "id": "5d7a21f19ef715269ae90c50",
                "url": "https://rest_api_root/v1/ship/idea_property_plans/5d7a21f19ef715269ae90c50"
            },
            "property": {
                "id": "solution",
                "url": "https://rest_api_root/v1/ship/idea_properties/solution",
                "name": "解决方案",
                "type": "select"
            }
        }
    ]
}
```

---

### 获取需求属性方案列表

**`GET /v1/ship/idea_property_plans`**

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 2,
    "values": [
        {
            "id": "5d7a21f19ef715269ae90c50",
            "url": "https://rest_api_root/v1/ship/idea_property_plans/5d7a21f19ef715269ae90c50",
            "product": null
        },
        {
            "id": "5f8a21f18ef715265de90c22",
            "url": "https://rest_api_root/v1/ship/idea_property_plans/5f8a21f18ef715265de90c22",
            "product": {
                 "id": "6422711c3f12e6c1e46d40e9",
                 "url": "https://rest_api_root/v1/ship/products/6422711c3f12e6c1e46d40e9",
                 "identifier": "SLC",
                 "name": "示例产品",
                 "is_archived": 0,
                 "is_deleted": 0
            }
        }
    ]
}
```

---

### 部分更新一个需求属性

**`PATCH /v1/ship/idea_properties/{property_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `property_id` | String | 需求属性的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `name` *(可选)* | String | 需求属性的名称。在一个企业中这个名称是唯一的。 |
| `options` *(可选)* | Object[] | 选择项列表。options是整体更新的。 |
| `options._id` *(可选)* | String | 选择项id。该值在选择项中是唯一的。 |
| `options.text` | String | 选择项内容。该值在选择项中是唯一的。 |
| `options.parent_id` *(可选)* | String | 选择项父级id。当属性类型为cascade_select,cascade_multi_select时，parent_id参数有效，用于构建级联类型选择项之间的父子关系，最多存在4级。 |

**响应示例（select）：**

```json
{
    "id": "severity-update",
    "url": "https://rest_api_root/v1/ship/idea_properties/severity",
    "name": "严重程度-update",
    "type": "select",
    "options": [
        {
            "_id": "64b9f741eec7fd94e63b411e",
            "text": "父"
        },
        {
            "_id": "64b9f741eec7fd94e63b411f",
            "text": "子"
        }
    ],
    "is_removable": true,
    "is_name_editable": true,
    "is_options_editable": true,
    "select_all_level": false,
    "display_all_level": false,
    "display_separator": null
}
```

**响应示例（cascade_select）：**

```json
{
    "id": "jiliandanxuan",
    "url": "https://rest_api_root/v1/ship/idea_properties/jiliandanxuan",
    "name": "级联单选-update",
    "type": "cascade_select",
    "options": [
        {
            "_id": "64b9f741eec7fd94e63b411e",
            "text": "父-update"
        },
        {
            "_id": "64b9f741eec7fd94e63b411f",
            "text": "子-update",
            "parent_id": "64b9f741eec7fd94e63b411e"
        }
    ],
    "is_removable": true,
    "is_name_editable": true,
    "is_options_editable": true,
    "select_all_level": false,
    "display_all_level": false,
    "display_separator": "/"
}
```

---

## 页面

### 创建一个页面

**`POST /v1/wiki/pages`**

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `space_id` | String | 空间的id。 |
| `name` | String | 页面的名称。 |
| `parent_id` *(可选)* | String | 父页面的id。 |
| `content` *(可选)* | String | 页面的内容。 |
| `format_type` *(可选)* | String | 页面内容的格式。content和format_type字段必须同时传递。 |

**响应示例：**

```json
{
    "id": "63e1bf51760505c8795ebccc",
    "url": "https://rest_api_root/v1/wiki/pages/63e1bf51760505c8795ebccc",
    "space": {
        "id": "63e1bf51760505c8795ebcc8",
        "url": "https://rest_api_root/v1/wiki/spaces/63e1bf51760505c8795ebcc8",
        "name": "示例空间",
        "identifier": "DEMO",
        "is_archived": 0,
        "is_deleted": 0
    },
    "name": "示例页面",
    "type": "document",
    "short_id": "5-x6NN",
    "html_url": "https://yctech.pingcode.com/wiki/pages/5-x6NN",
    "parent": {
        "id": "63e1bf51760505c8795ebcce",
        "url": "https://rest_api_root/v1/wiki/pages/63e1bf51760505c8795ebcce",
        "name": "模板",
        "type": "document",
        "short_id": "c-x6NN",
        "html_url": "https://yctech.pingcode.com/wiki/pages/c-x6NN"
    },
    "type": "document",
    "icon": "file-fill",
    "readings": 0,
    "published_at": 1675738962,
    "published_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "participants": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "url": "https://rest_api_root/v1/participants/a0417f68e846aae315c85d24643678a9?principal_type=page&principal_id=63e1bf51760505c8795ebccc",
            "type": "user",
            "user": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            }
        }
    ],
    "created_at": 1675738962,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "updated_at": 1675738962,
    "updated_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "is_locked": 0,
    "is_archived": 0,
    "is_deleted": 0
}
```

---

### 删除一个页面

**`DELETE /v1/wiki/pages/{page_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `page_id` | String | 页面的id。 |

**响应示例：**

```json
{
    "id": "63e1bf51760505c8795ebccc",
    "url": "https://rest_api_root/v1/wiki/pages/63e1bf51760505c8795ebccc",
    "name": "示例页面updated",
    "type": "document",
    "short_id": "5-x6NN",
    "html_url": "https://yctech.pingcode.com/wiki/pages/5-x6NN",
    "space": {
        "id": "63e1bf51760505c8795ebcc8",
        "url": "https://rest_api_root/v1/wiki/spaces/63e1bf51760505c8795ebcc8",
        "name": "示例空间",
        "identifier": "DEMO",
        "is_archived": 0,
        "is_deleted": 0
    },
    "parent": {
        "id": "63e1bf51760505c8795ebcce",
        "url": "https://rest_api_root/v1/wiki/pages/63e1bf51760505c8795ebcce",
        "name": "模板",
        "type": "document",
        "short_id": "c-x6NN",
        "html_url": "https://yctech.pingcode.com/wiki/pages/c-x6NN"
    },
    "icon": "file-fill",
    "readings": 0,
    "published_at": 1675739999,
    "published_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "participants": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "url": "https://rest_api_root/v1/participants/a0417f68e846aae315c85d24643678a9?principal_type=page&principal_id=63e1bf51760505c8795ebccc",
            "type": "user",
            "user": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            }
        }
    ],
    "created_at": 1675738962,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "updated_at": 1675739999,
    "updated_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "is_locked": 0,
    "is_archived": 0,
    "is_deleted": 1
}
```

---

### 恢复一个页面到指定版本

**`POST /v1/wiki/pages/{page_id}/versions/{version_id}/restore`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `page_id` | String | 页面的id。 |
| `version_id` | String | 页面版本的id。 |

**响应示例：**

```json
{
    "id": "65093abf4d4c8ca623da8fff",
    "url": "https://rest_api_root/v1/wiki/pages/65093a8e4d4c8ca623da8fcd/versions/65093abf4d4c8ca623da8fff",
    "page": {
        "id": "65093a8e4d4c8ca623da8fcd",
        "url": "https://rest_api_root/v1/wiki/pages/65093a8e4d4c8ca623da8fcd",
        "name": "主页",
        "type": "document",
        "short_id": "5-x6NN",
        "html_url": "https://yctech.pingcode.com/wiki/pages/5-x6NN"
    },
    "name": "v2恢复自v1",
    "order": 2,
    "status": "published",
    "created_at": 1695103832,
    "created_by": {
         "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    }
}
```

---

### 更新一个文档正文

**`PUT /v1/wiki/pages/{page_id}/content`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `page_id` | String | 页面的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `content` | String | 页面的内容。 |
| `format_type` | String | 页面内容的格式。 |

**响应示例：**

```json
{
    "id": "65093a8e4d4c8ca623da8fcd",
    "url": "https://rest_api_root/v1/wiki/pages/65093a8e4d4c8ca623da8fcd/content",
    "version": {
        "id": "65093abf4d4c8ca623da8ffe",
        "url": "https://rest_api_root/v1/wiki/63e1bf51760505c8795ebccc/versions/65093abf4d4c8ca623da8ffe",
        "name": "v3",
        "order": 3
    },
    "format_type": "markdown",
    "content": "**空间是记录信息和知识的页面集合，通过组织页面层级将知识系统化、结构化，便于团队沉淀经验、共享资源，实现知识增值，加快知识流动，在知识管理层面助力企业更快更好的发布产品。** *PingCode* 空间支持以下特性： 页面支持插入多种元素以及关联页面，满足编写需要 编辑过程自动保存草稿，无需担心内容丢失 提供丰富的模板，使用模板保持页面的一致性，让空间更加规范 使用锁定功能锁定页面最终版本 删除的页面放进回收站，随时恢复 树状页面结构，直接拖动页面编排目录，让知识管理更方便高效 通过设置成员角色来进行权限控制 通过页面评论实现成员沟通交流，形成反馈闭环 **【PingCode 空间】当前处于不断迭代过程中，更多功能即将呈现，敬请期待~**"
}
```

---

### 获取一个文档正文

**`GET /v1/wiki/pages/{page_id}/content`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `page_id` | String | 页面的id。 |

**查询参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `format_type` *(可选)* | String | 正文格式。 |
| `version_id` *(可选)* | String | 页面版本的id。默认值为页面当前版本的id。 |

**响应示例：**

```json
{
    "id": "65093a8e4d4c8ca623da8fcd",
    "url": "https://rest_api_root/v1/wiki/pages/65093a8e4d4c8ca623da8fcd/content",
    "version": {
        "id": "65093abf4d4c8ca623da8ffe",
        "url": "https://rest_api_root/v1/wiki/63e1bf51760505c8795ebccc/versions/65093abf4d4c8ca623da8ffe",
        "name": "v3",
        "order": 3
    },
    "format_type": "text",
    "content": "空间是记录信息和知识的页面集合，通过组织页面层级将知识系统化、结构化，便于团队沉淀经验、共享资源，实现知识增值，加快知识流动，在知识管理层面助力企业更快更好的发布产品。 PingCode 空间支持以下特性： 页面支持插入多种元素以及关联页面，满足编写需要 编辑过程自动保存草稿，无需担心内容丢失 提供丰富的模板，使用模板保持页面的一致性，让空间更加规范 使用锁定功能锁定页面最终版本 删除的页面放进回收站，随时恢复 树状页面结构，直接拖动页面编排目录，让知识管理更方便高效 通过设置成员角色来进行权限控制 通过页面评论实现成员沟通交流，形成反馈闭环  【PingCode 空间】当前处于不断迭代过程中，更多功能即将呈现，敬请期待~"
}
```

---

### 获取一个页面的版本列表

**`GET /v1/wiki/pages/{page_id}/versions`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `page_id` | String | 页面的id。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "65093abf4d4c8ca623da8ffe",
            "url": "https://rest_api_root/v1/wiki/pages/65093a8e4d4c8ca623da8fcd/versions/65093abf4d4c8ca623da8ffe",
            "page": {
                "id": "65093a8e4d4c8ca623da8fcd",
                "url": "https://rest_api_root/v1/wiki/pages/65093a8e4d4c8ca623da8fcd",
                "name": "主页",
                "type": "document",
                "short_id": "AAx6NN",
                "html_url": "https://yctech.pingcode.com/wiki/pages/AAx6NN"
            },
            "name": "v1",
            "order": 1,
            "status": "published",
            "created_at": 1695103679,
            "created_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            }
        }
    ]
}
```

---

### 获取统计页面

**`GET /v1/dashboard/pages`**

**查询参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `category` | String | 统计页面的类型。 |
| `count` *(可选)* | String | 统计页面的数量。 |

**响应示例：**

```json
[
    {
       "id": "63e1bf51760505c8795ebccc",
       "url": "https://rest_api_root/v1/wiki/pages/63e1bf51760505c8795ebccc",
       "space": {
           "id": "63e1bf51760505c8795ebcc8",
           "url": "https://rest_api_root/v1/wiki/spaces/63e1bf51760505c8795ebcc8",
           "name": "Demo space",
           "identifier": "DEMO",
           "is_archived": 0,
           "is_deleted": 0
       },
       "name": "Demo page",
       "type": "document",
       "parent": {
           "id": "63e1bf51760505c8795ebcce",
           "url": "https://rest_api_root/v1/wiki/pages/63e1bf51760505c8795ebcce",
           "name": "Demo parent page",
           "type": "document",
           "short_id": "x-x6NN",
           "html_url": "https://yctech.pingcode.com/wiki/pages/x-x6NN"
       },
       "icon": "file-fill",
       "readings": 10,
       "published_at": 1694065129,
       "published_by": {
           "id": "a0417f68e846aae315c85d24643678a9",
           "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
           "name": "john",
           "display_name": "John",
           "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
       },
       "participants": [
           {
               "id": "a0417f68e846aae315c85d24643678a9",
               "url": "https://rest_api_root/v1/participants/a0417f68e846aae315c85d24643678a9?principal_type=page&principal_id=63e1bf51760505c8795ebccc",
               "type": "user",
               "user": {
                   "id": "a0417f68e846aae315c85d24643678a9",
                   "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                   "name": "john",
                   "display_name": "John",
                   "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
               }
           }
        ],
        "created_at": 1675738962,
        "created_by": {
            "id": "a0417f68e846aae315c85d24643678a9",
            "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
            "name": "john",
            "display_name": "John",
            "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
        },
        "updated_at": 1694065129,
        "updated_by": {
            "id": "a0417f68e846aae315c85d24643678a9",
            "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
            "name": "john",
            "display_name": "John",
            "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
         },
         "is_archived": 0,
         "is_deleted": 0
    }
]
```

---

### 获取页面列表

**`GET /v1/wiki/pages?space_id={space_id}`**

**查询参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `space_id` | String | 空间的id。 |
| `parent_id` *(可选)* | String | 父页面的id。 |
| `created_between` *(可选)* | String | 创建时间介于的时间范围，通过','分割起始时间。 |
| `updated_between` *(可选)* | String | 更新时间介于的时间范围，通过','分割起始时间。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 2,
    "values": [
        {
            "id": "63e1bf51760505c8795ebccc",
            "url": "https://rest_api_root/v1/wiki/pages/63e1bf51760505c8795ebccc",
            "space": {
                "id": "63e1bf51760505c8795ebcc8",
                "url": "https://rest_api_root/v1/wiki/spaces/63e1bf51760505c8795ebcc8",
                "name": "示例空间",
                "identifier": "DEMO",
                "is_archived": 0,
                "is_deleted": 0
            },
            "name": "示例页面",
            "type": "document",
            "short_id": "5-x6NN",
            "html_url": "https://yctech.pingcode.com/wiki/pages/5-x6NN",
            "parent": {
                "id": "63e1bf51760505c8795ebcce",
                "url": "https://rest_api_root/v1/wiki/pages/63e1bf51760505c8795ebcce",
                "name": "模板",
                "type": "document",
                "short_id": "5-x6NN",
                "html_url": "https://yctech.pingcode.com/wiki/pages/5-x6NN"
            },
            "icon": "file-fill",
            "readings": 10,
            "published_at": 1694065129,
            "published_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "participants": [
                {
                    "id": "a0417f68e846aae315c85d24643678a9",
                    "url": "https://rest_api_root/v1/participants/a0417f68e846aae315c85d24643678a9?principal_type=page&principal_id=63e1bf51760505c8795ebccc",
                    "type": "user",
                    "user": {
                        "id": "a0417f68e846aae315c85d24643678a9",
                        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                        "name": "john",
                        "display_name": "John",
                        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
                    }
                }
            ],
            "created_at": 1675738962,
            "created_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "updated_at": 1694065129,
            "updated_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "is_locked": 0,
            "is_archived": 0,
            "is_deleted": 0
        },
        {
            "id": "63e1bf51760505c8795ebcce",
            "url": "https://rest_api_root/v1/wiki/pages/63e1bf51760505c8795ebcce",
            "space": {
                "id": "63e1bf51760505c8795ebcc8",
                "url": "https://rest_api_root/v1/wiki/spaces/63e1bf51760505c8795ebcc8",
                "name": "示例空间",
                "identifier": "DEMO",
                "is_archived": 0,
                "is_deleted": 0
            },
            "name": "模板",
            "type": "document",
            "short_id": "5-x6NN",
            "html_url": "https://yctech.pingcode.com/wiki/pages/5-x6NN",
            "parent": null,
            "icon": "file-fill",
            "readings": 0,
            "published_at": 1694065129,
            "published_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "participants": [
                {
                    "id": "a0417f68e846aae315c85d24643678a9",
                    "url": "https://rest_api_root/v1/participants/a0417f68e846aae315c85d24643678a9?principal_type=page&principal_id=63e1bf51760505c8795ebcce",
                    "type": "user",
                    "user": {
                        "id": "a0417f68e846aae315c85d24643678a9",
                        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                        "name": "john",
                        "display_name": "John",
                        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
                    }
                }
            ],
            "created_at": 1675738962,
            "created_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "updated_at": 1694065129,
            "updated_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "is_locked": 0,
            "is_archived": 0,
            "is_deleted": 0
        }
    ]
}
```

---

### 部分更新一个页面

**`PATCH /v1/wiki/pages/{page_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `page_id` | String | 页面的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `name` *(可选)* | String | 页面的名称。 |
| `parent_id` *(可选)* | String | 父页面的id。 |
| `lock` *(可选)* | Number | 是否锁定页面。 |

**响应示例：**

```json
{
    "id": "63e1bf51760505c8795ebccc",
    "url": "https://rest_api_root/v1/wiki/pages/63e1bf51760505c8795ebccc",
    "space": {
        "id": "63e1bf51760505c8795ebcc8",
        "url": "https://rest_api_root/v1/wiki/spaces/63e1bf51760505c8795ebcc8",
        "name": "示例空间",
        "identifier": "DEMO",
        "is_archived": 0,
        "is_deleted": 0
    },
    "name": "示例页面updated",
    "type": "document",
    "short_id": "5-x6NN",
    "html_url": "https://yctech.pingcode.com/wiki/pages/5-x6NN",
    "parent": {
        "id": "63e1bf51760505c8795ebcce",
        "url": "https://rest_api_root/v1/wiki/pages/63e1bf51760505c8795ebcce",
        "name": "模板",
        "type": "document",
        "short_id": "c-x6NN",
        "html_url": "https://yctech.pingcode.com/wiki/pages/c-x6NN"
    },
    "icon": "file-fill",
    "readings": 0,
    "published_at": 1675739999,
    "published_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "participants": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "url": "https://rest_api_root/v1/participants/a0417f68e846aae315c85d24643678a9?principal_type=page&principal_id=63e1bf51760505c8795ebccc",
            "type": "user",
            "user": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            }
        }
    ],
    "created_at": 1675738962,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "updated_at": 1675739999,
    "updated_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "is_locked": 1,
    "is_archived": 0,
    "is_deleted": 0
}
```

---

## 项目

### 创建一个项目

**`POST /v1/project/projects`**

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `scope_type` *(可选)* | String | 项目的所属类型。默认值organization。允许值分别表示企业可见和团队可见。 |
| `scope_id` *(可选)* | String | 项目的所属id。当scope_type为user_group时，该字段必填，且表示团队id；当scope_type为其他值时，该字段无效。 |
| `name` | String | 项目的名称。最大长度为255。 |
| `visibility` *(可选)* | String | 项目的可见范围。允许值分别表示组织全部成员可见和仅项目成员可见。 |
| `type` | String | 项目的类型。 |
| `identifier` | String | 项目的标识。在一个企业中这个标识是唯一的。项目的标识由大写英文字母/数字/下划线/连接线组成（不超过15个字符）。 |
| `description` *(可选)* | String | 项目的描述。 |
| `members` *(可选)* | Object[] | 项目成员的列表。当项目的scope_type变为user_group时，项目成员必须是scope_id指定的团队内的成员；当scope_type为其他值时，项目成员可以是任意成员或团队。 |
| `members.id` | String | 企业成员或团队的id。 |
| `members.type` | String | 项目成员的类型。 |
| `start_at` *(可选)* | Number | 项目开始的时间。 |
| `end_at` *(可选)* | Number | 项目结束的时间。 |
| `assignee_id` *(可选)* | String | 项目负责人的id。 |

**响应示例：**

```json
{
    "id": "5eb623f6a70571487ea47000",
    "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
    "type": "scrum",
    "process_id": "5fa690f1ae0571487ea49030",
    "scope_type": "user_group",
    "scope_id": "63c8fb32729dee3334d96af7",
    "name": "Scrum项目",
    "color": "#F693E7",
    "identifier": "SCR",
    "visibility": "private",
    "description": "这是一个scrum类型的项目",
    "state": {
        "id": "66cbf3b4b78a55fcd1a76296",
        "url": "http://rest_api_root/v1/project/project_states/66cbf3b4b78a55fcd1a76296",
        "name": "正常",
        "type": "in_progress"
    },
    "assignee": {
        "id": "8168dd057b104c81bceb2e48bdf283d0",
        "url": "http://rest_api_root/v1/directory/users/8168dd057b104c81bceb2e48bdf283d0",
        "name": "john",
        "display_name": "john",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "start_at": 1680278400,
    "end_at": 1682870399,
    "properties": {},
    "members": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/members/a0417f68e846aae315c85d24643678a9",
            "type": "user",
            "user": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            }
        }
    ],
    "created_at": 1583290300,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "updated_at": 1583290300,
    "updated_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "is_archived": 0,
    "is_deleted": 0
}
```

---

### 向项目中添加一个成员

**`POST /v1/project/projects/{project_id}/members`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `project_id` | String | 项目的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `member` | Object | 项目的成员。 |
| `member.id` | String | 企业成员或团队的id。 |
| `member.type` | String | 项目成员的类型。 |
| `role_id` *(可选)* | String | 角色的id。 |

**响应示例：**

```json
{
    "id": "a0417f68e846aae315c85d24643678a9",
    "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/members/a0417f68e846aae315c85d24643678a9",
    "project": {
        "id": "5eb623f6a70571487ea47000",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
        "identifier": "SCR",
        "name": "Scrum项目",
        "type": "scrum",
        "is_archived": 0,
        "is_deleted": 0
    },
    "type": "user",
    "user": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "role": {
        "id": "6422711c3f12e6c1e46d40e6",
        "url": "https://rest_api_root/v1/directory/roles/6422711c3f12e6c1e46d40e6",
        "name": "管理员"
    }
}
```

---

### 向项目中添加一个项目属性

**`POST /v1/project/projects/{project_id}/project_properties`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `project_id` | String | 项目的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `property_id` | String | 项目属性的id。 |

**响应示例：**

```json
{
  "id": "risk",
  "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/project_properties/risk",
  "project": {
    "id": "5eb623f6a70571487ea47000",
    "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
    "identifier": "SCR",
    "name": "Scrum项目",
    "type": "scrum",
    "is_archived": 0,
    "is_deleted": 0
  },
  "property": {
    "id": "risk",
    "url": "http://rest_api_root/v1/project/project_properties/risk",
    "name": "项目风险",
    "type": "select",
    "options": [
      {
        "_id": "5efb1859110533727a82c603",
        "text": "高"
      },
      {
        "_id": "5efb1859110533727a82c604",
        "text": "中"
      },
      {
        "_id": "5efb1859110533727a82c605",
        "text": "低"
      }
    ]
  }
}
```

---

### 在项目中移除一个成员

**`DELETE /v1/project/projects/{project_id}/members/{member_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `project_id` | String | 项目的id。 |
| `member_id` | String | 企业成员或团队的id。 |

**响应示例：**

```json
{
    "id": "a0417f68e846aae315c85d24643678a9",
    "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/members/a0417f68e846aae315c85d24643678a9",
    "project": {
        "id": "5eb623f6a70571487ea47000",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
        "identifier": "SCR",
        "name": "Scrum项目",
        "type": "scrum",
        "is_archived": 0,
        "is_deleted": 0
    },
    "type": "user",
    "user": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "role": {
        "id": "6422711c3f12e6c1e46d40e6",
        "url": "https://rest_api_root/v1/directory/roles/6422711c3f12e6c1e46d40e6",
        "name": "管理员"
    }
}
```

---

### 在项目中移除一个项目属性

**`DELETE /v1/project/projects/{project_id}/project_properties/{property_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `project_id` | String | 项目的id。 |
| `property_id` | String | 项目属性的id。 |

**响应示例：**

```json
{
  "id": "risk",
  "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/project_properties/risk",
  "project": {
    "id": "5eb623f6a70571487ea47000",
    "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
    "identifier": "SCR",
    "name": "Scrum项目",
    "type": "scrum",
    "is_archived": 0,
    "is_deleted": 0
  },
  "property": {
    "id": "risk",
    "url": "http://rest_api_root/v1/project/project_properties/risk",
    "name": "项目风险",
    "type": "select",
    "options": [
      {
        "_id": "5efb1859110533727a82c603",
        "text": "高"
      },
      {
        "_id": "5efb1859110533727a82c604",
        "text": "中"
      },
      {
        "_id": "5efb1859110533727a82c605",
        "text": "低"
      }
    ]
  }
}
```

---

### 复制一个项目

**`POST /v1/project/projects/{project_id}/clone`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `project_id` | String | 项目的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `scope_type` *(可选)* | String | 项目的所属类型。默认使用原项目的所属。允许值分别表示企业可见和团队可见。 |
| `scope_id` *(可选)* | String | 项目的所属id。当scope_type为user_group时，该字段必填，且表示团队id；当scope_type为其他值时，该字段无效。 |
| `name` *(可选)* | String | 项目的名称。最大长度为255。默认使用原项目的名称。 |
| `visibility` *(可选)* | String | 项目的可见范围。默认使用原项目的可见范围。允许值分别表示组织全部成员可见和仅项目成员可见。 |
| `identifier` | String | 项目的标识。在一个企业中这个标识是唯一的。项目的标识由大写英文字母/数字/下划线/连接线组成（不超过15个字符）。 |
| `description` *(可选)* | String | 项目的描述。默认使用原项目的描述。 |
| `members` *(可选)* | Object[] | 项目成员的列表。当项目的scope_type变为user_group时，项目成员必须是scope_id指定的团队内的成员；当scope_type为其他值时，项目成员可以是任意成员或团队，默认使用原项目的成员列表。 |
| `members.id` | String | 企业成员或团队的id。 |
| `members.type` | String | 项目成员的类型。 |

**响应示例：**

```json
{
    "id": "5ab623f6a70571487ea47001",
    "url": "https://rest_api_root/v1/project/projects/5ab623f6a70571487ea47001",
    "type": "scrum",
    "process_id": "5fa690f1ae0571487ea49030",
    "scope_type": "user_group",
    "scope_id": "63c8fb32729dee3334d96af7",
    "name": "复制的Scrum项目",
    "color": "#F693E7",
    "identifier": "SCRC",
    "visibility": "public",
    "description": "这是一个复制的Scrum类型的项目",
    "state": {
        "id": "66cbf3b4b78a55fcd1a76296",
        "url": "http://rest_api_root/v1/project/project_states/66cbf3b4b78a55fcd1a76296",
        "name": "正常",
        "type": "in_progress"
    },
    "assignee": {
        "id": "8168dd057b104c81bceb2e48bdf283d0",
        "url": "http://rest_api_root/v1/directory/users/8168dd057b104c81bceb2e48bdf283d0",
        "name": "john",
        "display_name": "john",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "start_at": 1680278400,
    "end_at": 1682870399,
    "properties": {},
    "members": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/members/a0417f68e846aae315c85d24643678a9",
            "type": "user",
            "user": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            }
        }
    ],
    "created_at": 1583290300,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "updated_at": 1583290300,
    "updated_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "is_archived": 0,
    "is_deleted": 0
}
```

---

### 获取一个项目进度

**`GET /v1/project/projects/{project_id}/progress`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `project_id` | String | 项目的id。 |

**响应示例：**

```json
{
    "work_item": {
        "total": 4,
        "pending_count": 1,
        "in_progress_count": 2,
        "completed_count": 1
    }
}
```

---

### 获取项目中的成员列表

**`GET /v1/project/projects/{project_id}/members`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `project_id` | String | 项目的id。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 2,
    "values": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/members/a0417f68e846aae315c85d24643678a9",
            "project": {
                "id": "5eb623f6a70571487ea47000",
                "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
                "identifier": "SCR",
                "name": "Scrum项目",
                "type": "scrum",
                "is_archived": 0,
                "is_deleted": 0
            },
            "type": "user",
            "user": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "role": {
                "id": "6422711c3f12e6c1e46d40e6",
                "url": "https://rest_api_root/v1/directory/roles/6422711c3f12e6c1e46d40e6",
                "name": "管理员"
            }
        },
        {
            "id": "63c8fb32729dee3334d96af7",
            "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/members/63c8fb32729dee3334d96af7",
            "project": {
                "id": "5eb623f6a70571487ea47000",
                "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
                "identifier": "SCR",
                "name": "Scrum项目",
                "type": "scrum",
                "is_archived": 0,
                "is_deleted": 0
            },
            "type": "user_group",
            "user_group": {
                "id": "63c8fb32729dee3334d96af7",
                "url": "https://rest_api_root/v1/directory/groups/63c8fb32729dee3334d96af7",
                "name": "Open Team"
            },
            "role": {
                "id": "6422711c3f12e6c1e46d40e6",
                "url": "https://rest_api_root/v1/directory/roles/6422711c3f12e6c1e46d40e6",
                "name": "管理员"
            }
        }
    ]
}
```

---

### 获取项目中的项目属性列表

**`GET /v1/project/projects/{project_id}/project_properties`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `project_id` | String | 项目的id。 |

**响应示例：**

```json
{
  "page_size": 30,
  "page_index": 0,
  "total": 1,
  "values": [
    {
      "id": "risk",
      "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/project_properties/risk",
      "project": {
        "id": "5eb623f6a70571487ea47000",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
        "identifier": "SCR",
        "name": "Scrum项目",
        "type": "scrum",
        "is_archived": 0,
        "is_deleted": 0
      },
      "property": {
        "id": "risk",
        "url": "http://rest_api_root/v1/project/project_properties/risk",
        "name": "项目风险",
        "type": "select",
        "options": [
          {
            "_id": "5efb1859110533727a82c603",
            "text": "高"
          },
          {
            "_id": "5efb1859110533727a82c604",
            "text": "中"
          },
          {
            "_id": "5efb1859110533727a82c605",
            "text": "低"
          }
        ]
      }
    },
    {
      "id": "name",
      "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/project_properties/name",
      "project": {
        "id": "5eb623f6a70571487ea47000",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
        "identifier": "SCR",
        "name": "Scrum项目",
        "type": "scrum",
        "is_archived": 0,
        "is_deleted": 0
      },
      "property": {
        "id": "name",
        "url": "https://rest_api_root/v1/project/project_properties/name",
        "name": "名称",
        "type": "text",
        "options": null,
        "is_removable": false,
        "is_name_editable": false,
        "is_options_editable": false
      }
    }
  ]
}
```

---

### 获取项目列表

**`GET /v1/project/projects`**

**查询参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `identifier` *(可选)* | String | 项目的标识。 |
| `type` *(可选)* | String | 项目的类型。 |
| `include_deleted` *(可选)* | Boolean | 是否查询已删除的项目。该值默认为false。 |
| `include_archived` *(可选)* | Boolean | 是否查询已归档的项目。该值默认为false。 |

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 1,
    "values": [
        {
            "id": "5eb623f6a70571487ea47000",
            "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
            "type": "scrum",
            "process_id": "5fa690f1ae0571487ea49030",
            "scope_type": "user_group",
            "scope_id": "63c8fb32729dee3334d96af7",
            "name": "Scrum项目",
            "color": "#F693E7",
            "identifier": "SCR",
            "visibility": "private",
            "description": "这是一个scrum类型的项目",
            "state": {
                "id": "66cbf3b4b78a55fcd1a76296",
                "url": "http://rest_api_root/v1/project/project_states/66cbf3b4b78a55fcd1a76296",
                "name": "正常",
                "type": "in_progress"
            },
            "assignee": {
                "id": "8168dd057b104c81bceb2e48bdf283d0",
                "url": "http://rest_api_root/v1/directory/users/8168dd057b104c81bceb2e48bdf283d0",
                "name": "john",
                "display_name": "john",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "start_at": 1680278400,
            "end_at": 1682870399,
            "properties": {
                "prop_a": "prop_a_value",
                "prop_b": "prop_b_value"
            },
            "members": [
                {
                    "id": "a0417f68e846aae315c85d24643678a9",
                    "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/members/a0417f68e846aae315c85d24643678a9",
                    "type": "user",
                    "user": {
                        "id": "a0417f68e846aae315c85d24643678a9",
                        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                        "name": "john",
                        "display_name": "John",
                        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
                    }
                }
            ],
            "created_at": 1583290300,
            "created_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "updated_at": 1583290300,
            "updated_by": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            },
            "is_archived": 0,
            "is_deleted": 0
        }
    ]
}
```

---

### 获取项目状态列表

**`GET /v1/project/project/states?project_id={project_id}`**

**查询参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `project_id` | String | 项目的id。 |

**响应示例：**

```json
{
    "page_index": 0,
    "page_size": 30,
    "total": 5,
    "values": [
        {
            "id": "66cbf5401e7cc374c85acb1b",
            "url": "https://rest_api_root/v1/project/project_states/66cbf5401e7cc374c85acb1b",
            "name": "未开始",
            "type": "pending"
        },
        {
            "id": "66cbf5401e7cc374c85acb1c",
            "url": "https://rest_api_root/v1/project/project_states/66cbf5401e7cc374c85acb1c",
            "name": "正常",
            "type": "in_progress"
        },
        {
            "id": "66cbf5401e7cc374c85acb1d",
            "url": "https://rest_api_root/v1/project/project_states/66cbf5401e7cc374c85acb1d",
            "name": "预警",
            "type": "in_progress"
        },
        {
            "id": "66cbf5401e7cc374c85acb1e",
            "url": "https://rest_api_root/v1/project/project_states/66cbf5401e7cc374c85acb1e",
            "name": "延期",
            "type": "in_progress"
        },
        {
            "id": "66cbf5401e7cc374c85acb1f",
            "url": "https://rest_api_root/v1/project/project_states/66cbf5401e7cc374c85acb1f",
            "name": "结束",
            "type": "completed"
        }
    ]
}
```

---

### 部分更新一个项目

**`PATCH /v1/project/projects/{project_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `project_id` | String | 项目的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `name` *(可选)* | String | 项目的名称。最大长度为255。 |
| `identifier` *(可选)* | String | 项目的标识。在一个企业中这个标识是唯一的。项目的标识由大写英文字母/数字/下划线/连接线组成（不超过15个字符）。 |
| `description` *(可选)* | String | 项目的描述。 |
| `start_at` *(可选)* | Number | 项目开始的时间。 |
| `end_at` *(可选)* | Number | 项目结束的时间。 |
| `assignee_id` *(可选)* | String | 项目负责人的id。 |
| `state_id` *(可选)* | String | 项目状态的id。项目状态可以通过获取项目状态列表获取。 |
| `properties` *(可选)* | Object | 项目自定义属性的键值对集合。需要注意自定义属性需要包含在项目属性方案中。例如属性方案中包含prop_a和prop_b。 |
| `properties.prop_a` *(可选)* | Object | 项目自定义属性prop_a。 |
| `properties.prop_b` *(可选)* | Object | 项目自定义属性prop_b。 |

**响应示例：**

```json
{
    "id": "5eb623f6a70571487ea47000",
    "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
    "type": "scrum",
    "process_id": "5fa690f1ae0571487ea49030",
    "scope_type": "user_group",
    "scope_id": "63c8fb32729dee3334d96af7",
    "name": "Scrum项目",
    "color": "#F693E7",
    "identifier": "SCR",
    "visibility": "private",
    "description": "这是一个scrum类型的项目",
    "state": {
        "id": "66cbf3b4b78a55fcd1a76296",
        "url": "http://rest_api_root/v1/project/project_states/66cbf3b4b78a55fcd1a76296",
        "name": "正常",
        "type": "in_progress"
    },
    "assignee": {
        "id": "8168dd057b104c81bceb2e48bdf283d0",
        "url": "http://rest_api_root/v1/directory/users/8168dd057b104c81bceb2e48bdf283d0",
        "name": "john",
        "display_name": "john",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "start_at": 1680278400,
    "end_at": 1682870399,
    "properties": {
        "prop_a": "prop_a_value",
        "prop_b": "prop_b_value"
    },
    "members": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/members/a0417f68e846aae315c85d24643678a9",
            "type": "user",
            "user": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            }
        }
    ],
    "created_at": 1583290300,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "updated_at": 1583293300,
    "updated_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "is_archived": 0,
    "is_deleted": 0
}
```

---

### 部分更新一个项目内的成员

**`PATCH /v1/project/projects/{project_id}/members/{member_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `project_id` | String | 项目的id。 |
| `member_id` | String | 企业成员或团队的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `role_id` *(可选)* | String | 角色的id。管理员可以更改其他用户的角色，但非管理员用户无权更改其他用户的角色。 |

**响应示例：**

```json
{
    "id": "a0417f68e846aae315c85d24643678a9",
    "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/members/a0417f68e846aae315c85d24643678a9",
    "project": {
        "id": "5eb623f6a70571487ea47000",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
        "identifier": "SCR",
        "name": "Scrum项目",
        "type": "scrum",
        "is_archived": 0,
        "is_deleted": 0
    },
    "type": "user",
    "user": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "role": {
        "id": "6422711c3f12e6c1e46d40e6",
        "url": "https://rest_api_root/v1/directory/roles/6422711c3f12e6c1e46d40e6",
        "name": "管理员"
    }
}
```

---

## 项目管理

### Kanban

用于查看和管理看板相关信息。

---

### Scrum

用于查看和管理迭代相关信息。

---

### 发布

用于查看和管理发布相关信息。 资源地址：{GET} https://rest_api_root/v1/project/projects/{project_id}/versions/{version_id}

**全量数据示例：**

```json
{
    "id": "5eb623f6a70571487ea47001",
    "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/versions/5eb623f6a70571487ea47001",
    "project": {
        "id": "5eb623f6a70571487ea47000",
        "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
        "identifier": "SCR",
        "type": "scrum",
        "name": "Scrum项目",
        "is_archived": 0,
        "is_deleted": 0
    },
    "name": "1.0.0",
    "start_at": 1565193600,
    "end_at": 1566403200,
    "stage": {
        "id": "5f44a8f8bb347b14b49624a1",
        "url": "https://rest_api_root/v1/project/stages/5f44a8f8bb347b14b49624a1",
        "name": "未开始",
        "type": "pending",
        "color": "#FA8888"
    },
    "assignee": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "stages": [
        {
            "id": "5f44a8f8bb347b14b49624a1",
            "url": "https://rest_api_root/v1/project/stages/5f44a8f8bb347b14b49624a1",
            "name": "未开始",
            "operate_at": 1565366400
        },
        {
            "id": "5f44a8f8bb347b14b49624a2",
            "url": "https://rest_api_root/v1/project/stages/5f44a8f8bb347b14b49624a2",
            "name": "进行中",
            "operate_at": null
        },
        {
            "id": "5f44a8f8bb347b14b49624a3",
            "url": "https://rest_api_root/v1/project/stages/5f44a8f8bb347b14b49624a3",
            "name": "已发布",
            "operate_at": null
        }
    ],
    "progress": 0,
    "changelog": "发布日志",
    "operate_at": 1565366400,
    "categories": [
        {
            "id": "676a460a0fd987b7ea320889",
            "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/version_categories/676a460a0fd987b7ea320889",
            "name": "私有部署发布"
        }
    ],
    "created_at": 1565366200,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "updated_at": 1565366200,
    "updated_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    }
}
```

**引用数据示例：**

```json
{
    "id": "5eb623f6a70571487ea47001",
    "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/versions/5eb623f6a70571487ea47001",
    "name": "1.0.0",
    "start_at": 1565193600,
    "end_at": 1566403200,
    "stage": {
        "id": "5f44a8f8bb347b14b49624a1",
        "url": "https://rest_api_root/v1/project/stages/5f44a8f8bb347b14b49624a1",
        "name": "未开始",
        "type": "pending",
        "color": "#FA8888"
    }
}
```

---

### 工作项

用于查看和管理工作项相关信息。 资源地址：{GET} https://rest_api_root/v1/project/work_items/{work_item_id}

**全量数据示例：**

```json
{
     "id": "6efca131b06329c524cdd2fb",
     "url": "https://rest_api_root/v1/project/work_items/6efca131b06329c524cdd2fb",
     "project": {
         "id": "5eb623f6a70571487ea47004",
         "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47004",
         "identifier": "HBR",
         "name": "Hybrid项目",
         "type": "hybrid",
         "is_archived": 0,
         "is_deleted": 0
     },
     "identifier": "HBR-1",
     "title": "这是一个用户故事",
     "type": "story",
     "start_at": 1583290309,
     "end_at": 1583290347,
     "parent_id": "5edfa3b67463c1ee626c0980",
     "short_id": "eqWqLmTO",
     "html_url": "https://yctech.pingcode.com/pjm/workitems/eqWqLmTO",
     "parent": {
         "id": "5edfa3b67463c1ee626c0980",
         "url": "https://rest_api_root/v1/project/work_items/5edfa3b67463c1ee626c0980",
         "identifier": "HBR-2",
         "title": "这是一个特性",
         "type": "feature",
         "start_at": 1583290309,
         "end_at": 1583290347,
         "parent_id": null,
         "properties": {
             "prop_a": "prop_a_value",
             "prop_b": "prop_b_value",
             "risk": null,
             "business_value": null,
             "effort": null,
             "backlog_type": null,
             "backlog_from": null
         }
      },
      "assignee": {
          "id": "a0417f68e846aae315c85d24643678a9",
          "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
          "name": "john",
          "display_name": "John",
          "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
      },
      "version": {
          "id": "5eb623487ea47001f6a70571",
          "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47004/versions/5eb623487ea47001f6a70582",
          "name": "1.0.1",
          "start_at": 1565255712,
          "end_at": 1565255879,
          "stage": {
              "id": "5f44a8f8bb347b14b49624a1",
              "url": "https://rest_api_root/v1/project/stages/5f44a8f8bb347b14b49624a1",
              "name": "未开始",
              "type": "pending",
              "color": "#FA8888"
          }
      },
      "sprint": {
          "id": "5ecf7b74eaab845a2aa53139",
          "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47004/sprints/5ecf7b74eaab845a2aa53139",
          "name": "Sprint 1",
          "start_at": 1589791860,
          "end_at": 1589791860,
          "status": "completed"
      },
      "state": {
          "id": "5c9b35de90ad7153c2062f18",
          "url": "https://rest_api_root/v1/project/work_item_states/5c9b35de90ad7153c2062f18",
          "name": "新建",
          "type": "pending",
          "color": "#ff7575"
      },
      "priority": {
          "id": "5eb623f6a70571487ea47111",
          "url": "https://rest_api_root/v1/project/work_item_priorities/5eb623f6a70571487ea47111",
          "name": "最高"
      },
      "board": {
          "id": "5eb623f6a70571487ea47333",
          "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47004/boards/5eb623f6a70571487ea47333",
          "name": "kanban",
          "work_item_types": [
             "epic",
             "feature",
             "issue",
             "story"
          ]
      },
      "entry": {
          "id": "5ee1c4fac5b3c51206f0a862",
          "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47004/boards/5eb623f6a70571487ea47333/entries/5ee1c4fac5b3c51206f0a862",
          "name": "需求池"
      },
      "swimlane": {
          "id": "5ee1c4fac5b3c51206f0a867",
          "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47004/boards/5eb623f6a70571487ea47333/swimlanes/5ee1c4fac5b3c51206f0a867",
          "name": "默认泳道"
      },
      "phase": {
          "id": "63761fee31caaf77189816b5",
          "url": "http://rest_api_root/v1/project/projects/5eb623f6a70571487ea47004/phases/63761fee31caaf77189816b5",
          "title": "这是一个阶段",
          "identifier": "WTF-1"
      },
      "description": "<p>这是一个用户故事的描述<p><img src=\"https://atlas.pingcode.com/files/public/689a9d124df436ef91923a3a\" originUrl=\"https://atlas.pingcode.com/files/public/689a9d124df436ef91923a3a\" alt=\"图片.png\" size=\"35460\" style=\"text-align: center;\" />",
      "completed_at": 1583290347,
      "story_points": 1,
      "estimated_workload": 1,
      "remaining_workload": 1,
      "properties": {
          "prop_a": "prop_a_value",
          "prop_b": "prop_b_value",
          "risk": null,
          "backlog_type": null,
          "backlog_from": null
      },
      "tags": [
          {
              "id": "5e6b35de50ef8153c2062f70",
              "url": "https://rest_api_root/v1/project/tags/5e6b35de50ef8153c2062f70",
              "name": "标签-1"
          }
      ],
      "participants": [
          {
              "id": "a0417f68e846aae315c85d24643678a9",
              "url": "https://rest_api_root/v1/participants/a0417f68e846aae315c85d24643678a9?principal_type=work_item&principal_id=5edca112b06305c524cad2fa",
              "type": "user",
              "user": {
                  "id": "a0417f68e846aae315c85d24643678a9",
                  "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                  "name": "john",
                  "display_name": "John",
                  "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
              }
          }
      ],
      "public_image_token": "-fkvANQ2dcVECK6Xg45L3kG8VCbOTK8NrNckGkxljRY",
      "created_at": 1583290300,
      "created_by": {
          "id": "a0417f68e846aae315c85d24643678a9",
          "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
          "name": "john",
          "display_name": "John",
          "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
      },
      "updated_at": 1583290300,
      "updated_by": {
          "id": "a0417f68e846aae315c85d24643678a9",
          "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
          "name": "john",
          "display_name": "John",
          "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
      },
      "is_archived": 0,
      "is_deleted": 0
}
```

**引用数据示例：**

```json
{
    "id": "5edca5d2fa112b06305c24ca",
    "url": "https://rest_api_root/v1/project/work_items/5edca5d2fa112b06305c24ca",
    "identifier": "KB-1",
    "title": "这是一个事务",
    "type": "issue",
    "start_at": 1583290309,
    "end_at": 1583290347,
    "parent_id": null,
    "short_id": "c9WqLmTO",
    "html_url": "https://yctech.pingcode.com/pjm/workitems/c9WqLmTO",
    "properties": {
        "prop_a": "prop_a_value",
        "prop_b": "prop_b_value",
        "entry_status": null,
        "entry_position": null,
        "operation_time": 1691571221
    }
}
```

---

### 项目

用于查看和管理项目相关信息。 资源地址：{GET} https://rest_api_root/v1/project/projects/{project_id}

**全量数据示例：**

```json
{
    "id": "5eb623f6a70571487ea47000",
    "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
    "type": "scrum",
    "process_id": "5fa690f1ae0571487ea49030",
    "scope_type": "user_group",
    "scope_id": "63c8fb32729dee3334d96af7",
    "name": "Scrum项目",
    "color": "#F693E7",
    "identifier": "SCR",
    "visibility": "private",
    "description": "这是一个scrum类型的项目",
    "state": {
        "id": "66cbf3b4b78a55fcd1a76296",
        "url": "http://rest_api_root/v1/project/project_states/66cbf3b4b78a55fcd1a76296",
        "name": "正常",
        "type": "in_progress"
    },
    "assignee": {
        "id": "8168dd057b104c81bceb2e48bdf283d0",
        "url": "http://rest_api_root/v1/directory/users/8168dd057b104c81bceb2e48bdf283d0",
        "name": "john",
        "display_name": "john",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "start_at": 1680278400,
    "end_at": 1682870399,
    "properties": {
        "prop_a": "prop_a_value",
        "prop_b": "prop_b_value"
    },
    "members": [
        {
            "id": "a0417f68e846aae315c85d24643678a9",
            "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/members/a0417f68e846aae315c85d24643678a9",
            "type": "user",
            "user": {
                "id": "a0417f68e846aae315c85d24643678a9",
                "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
                "name": "john",
                "display_name": "John",
                "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
            }
        },
        {
            "id": "63c8fb32729dee3334d96af7",
            "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000/members/63c8fb32729dee3334d96af7",
            "type": "user_group",
            "user_group": {
                "id": "63c8fb32729dee3334d96af7",
                "url": "https://rest_api_root/v1/directory/groups/63c8fb32729dee3334d96af7",
                "name": "Open Team"
            }
        }
    ],
    "created_at": 1583290300,
    "created_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "updated_at": 1583290300,
    "updated_by": {
        "id": "a0417f68e846aae315c85d24643678a9",
        "url": "https://rest_api_root/v1/directory/users/a0417f68e846aae315c85d24643678a9",
        "name": "john",
        "display_name": "John",
        "avatar": "https://s3.amazonaws.com/bucket/b46ef40c-e22e-4ecf-a599-cace9fba839a_160x160.png"
    },
    "is_archived": 0,
    "is_deleted": 0
}
```

**引用数据示例：**

```json
{
    "id": "5eb623f6a70571487ea47000",
    "url": "https://rest_api_root/v1/project/projects/5eb623f6a70571487ea47000",
    "identifier": "SCR",
    "name": "Scrum项目",
    "type": "scrum",
    "is_archived": 0,
    "is_deleted": 0
}
```

---

### 项目配置中心

---

### 项目配置中心

---

## 项目配置

### PatchV1ProjectProject_propertiesProperty_id

**`PATCH /v1/project/project_properties/{property_id}`**

**路径参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `property_id` | String | 项目属性的id。 |

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `name` *(可选)* | String | 项目属性的名称。在一个企业中这个名称是唯一的。 |
| `options` *(可选)* | Object[] | 选择项列表。options是整体更新的。 |
| `options._id` *(可选)* | Sting | 选择项的_id。如果option中不包含用于确定唯一性的_id，那么这个option将被视为新建，并为之创建新的_id。 |
| `options.text` | String | 选择项内容。 |
| `options.parent_id` *(可选)* | String | 选择项父级id。当属性类型为cascade_select,cascade_multi_select时，parent_id参数有效，用于构建级联类型选择项之间的父子关系，最多存在4级。 |

**响应示例：**

```json
{
    "id": "xiangmuguimo",
    "url": "https://rest_api_root/v1/project/project_properties/xiangmuguimo",
    "name": "项目规模",
    "type": "select",
    "options": [
        {
            "_id": "5efb1859110533727a82c605",
            "text": "大"
        },
        {
            "_id": "5efb1859110533727a82c606",
            "text": "小"
        }
    ],
    "is_removable": true,
    "is_name_editable": true,
    "is_options_editable": true
}
```

---

### 创建一个项目属性

**`POST /v1/project/project_properties`**

**Parameter：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `name` | String | 项目属性的名称。在一个企业中这个名称是唯一的。 |
| `type` | String | 项目属性的类型。 |
| `options` *(可选)* | Object[] | 选择项列表。当属性类型为select,multi_select,cascade_select,cascade_multi_select时可选填此项。 |
| `options._id` *(可选)* | String | 选择项id。该值在选择项中是唯一的。 |
| `options.text` | String | 选择项内容。该值在选择项中是唯一的。 |
| `options.parent_id` *(可选)* | String | 选择项父级id。当属性类型为cascade_select,cascade_multi_select时，parent_id参数有效，用于构建级联类型选择项之间的父子关系，最多存在4级。 |

**响应示例：**

```json
{
    "id": "risk",
    "url": "https://rest_api_root/v1/project/project_properties/risk",
    "name": "项目风险",
    "type": "select",
    "options": [
        {
            "_id": "5efb1859110533727a82c603",
            "text": "高"
        },
        {
            "_id": "5efb1859110533727a82c604",
            "text": "中"
        },
        {
            "_id": "5efb1859110533727a82c605",
            "text": "低"
        }
    ],
    "is_removable": false,
    "is_name_editable": false,
    "is_options_editable": false
}
```

---

### 获取项目属性列表

**`GET /v1/project/project_properties`**

**响应示例：**

```json
{
    "page_size": 30,
    "page_index": 0,
    "total": 2,
    "values": [
        {
            "id": "risk",
            "url": "https://rest_api_root/v1/project/project_properties/risk",
            "name": "项目风险",
            "type": "select",
            "options": [
                {
                    "_id": "5efb1859110533727a82c603",
                    "text": "高"
                },
                {
                    "_id": "5efb1859110533727a82c604",
                    "text": "中"
                },
                {
                    "_id": "5efb1859110533727a82c605",
                    "text": "低"
                }
            ],
            "is_removable": false,
            "is_name_editable": false,
            "is_options_editable": false
        },
        {
            "id": "name",
            "url": "https://rest_api_root/v1/project/project_properties/name",
            "name": "名称",
            "type": "text",
            "options": null,
            "is_removable": false,
            "is_name_editable": false,
            "is_options_editable": false
        }
    ]
}
```

---

## 项目配置中心

### 工作项配置

用于查看和管理工作项相关配置。

---

### 项目配置

用于查看和管理项目相关配置。

---

### 项目配置

用于查看和管理项目相关配置。

---

