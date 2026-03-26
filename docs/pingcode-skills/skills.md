# PingCode 工作流管理指南 (CLI)

## 边界
当前这个版本的 PingCode 工作流管理指南，是实现了对产品模块和开发流程的管理

## 工具
这个skills需要一个pingcode-client的工具来执行, 可以从[这里](https://github.com/x2x4com/pingcode-client)下载并编译，然后将编译后的二进制文件放到 scripts 中。
- scripts/pingcode-client

工具需要2个环境变量
- PINGCODE_CLIENT_ID
- PINGCODE_CLIENT_SECRET

## 角色

| 角色名称 | 角色描述 | tag标签 | Agent 角色 |
| -------- | -------- | ------- | ---------- |
| 产品经理 | 负责产品的需求管理、功能设计、技术方案等 | product | 需求分析 |
| 研发经理 | 负责将产品的需求转换为技术任务，代码审核 | arch | 技术架构 |
| 后端开发 | 负责产品的后端代码实现、测试、调试等 | backend | 后端开发 |
| 前端开发 | 负责产品的前端代码实现、测试、调试等 | frontend | 前端开发 |
| QA | 负责产品的功能测试、性能测试、安全测试等 | qa | QA |
| DevOps | 负责产品的部署、监控、维护等 | devops | DevOps |
| 安全专家 | 负责产品的安全设计、测试、维护等, 也是代码审核的一员 | security | 安全专家 |

通过 tag 标签，来标识不同的角色，下吗的命令可以列出和创建 tag 标签
```
$ pingcode-client project metadata tags list # 列出所有的 tag 标签
$ pingcode-client project metadata tags create --name "security" # 创建一个新的 tag 标签
```

注意标签名唯一，如果创建重复名称标签会报错
```
Error creating tag: create tag failed with status: 400, body: {"code":"100352","message":"'tag'资源已经存在"}
```


## 工作流程

### 产品经理/需求分析

#### 需求创建流程
| 步骤 | 描述 | 操作手册ID |
| ---- | ---- | ---------- |
| 1 | 确定产品范围，获取产品列表，与用户确认操作目标，记录产品ID | 参考`references/ship_guide.md`的产品列表 |
| 2 | 与用户沟通明确需求目标以及交付内容，时间点，然后创建需求 | 参考`references/ship_guide.md`创建需求 |

Example
```bash
# 步骤 1: 获取产品列表，与用户确认目标产品，记录 product_id
pingcode-client ship product list

# 步骤 2: 获取需求状态和优先级 ID（备用）
pingcode-client ship ideas states --product-id {product_id}
pingcode-client ship ideas priorities --product-id {product_id}

# 步骤 3: 创建需求
pingcode-client ship ideas create \
  --product-id {product_id} \
  --title "用户登录支持第三方 OAuth" \
  --desc "用户可通过 GitHub / Google 账号授权登录，减少注册摩擦，提升转化率"
```

### 研发经理/技术架构

#### 需求转化为技术任务
| 步骤 | 描述 | 操作手册ID |
| ---- | ---- | ---------- |
| 1 | 获取需求列表，与产品经理确认本迭代要实现的需求，记录 idea 标题和描述 | 参考 `references/ship_guide.md` 需求列表 |
| 2 | 获取开发项目列表，确认目标项目 ID | 参考 `references/workitem_guide.md` ID速查表 |
| 3 | 获取工作项类型、优先级、项目成员等元数据 | 参考 `references/workitem_guide.md` ID速查表 |
| 4 | 创建 Story 工作项，对应该需求，分配给合适的负责人，关联当前迭代 | 参考 `references/workitem_guide.md` 创建工作项 |
| 5 | 将 Story 拆解为具体子任务（task），按后端/前端/测试分配，打上对应角色标签 | 参考 `references/workitem_guide.md` 标签管理 |

Example
```bash
# 步骤 1: 查看待实现的需求
pingcode-client ship ideas list --product-id {product_id}

# 步骤 2: 获取项目和元数据
pingcode-client project list
pingcode-client project metadata types --project-id {project_id}
pingcode-client project metadata priorities --project-id {project_id}
pingcode-client project metadata members --project-id {project_id}
pingcode-client project iteration list --project-id {project_id}

# 步骤 3: 创建 Story（对应一条需求）
pingcode-client project workitem create \
  --project-id {project_id} \
  --title "实现第三方 OAuth 登录" \
  --desc "对应产品需求：用户登录支持第三方 OAuth" \
  --type-id story \
  --priority-id {priority_id} \
  --assignee-id {arch_member_id} \
  --sprint-id {iteration_id}

# 步骤 4: 拆解后端子任务
pingcode-client project workitem create \
  --project-id {project_id} \
  --title "后端: 实现 OAuth 回调接口" \
  --type-id task \
  --priority-id {priority_id} \
  --assignee-id {backend_member_id} \
  --parent-id {story_id}

# 步骤 5: 为后端任务打上 backend 角色标签
pingcode-client project metadata tags list
pingcode-client project workitem tag add -i {backend_task_id} -t {backend_tag_id}

# 步骤 6: 拆解前端子任务并打标签
pingcode-client project workitem create \
  --project-id {project_id} \
  --title "前端: 实现 OAuth 登录页面" \
  --type-id task \
  --priority-id {priority_id} \
  --assignee-id {frontend_member_id} \
  --parent-id {story_id}
pingcode-client project workitem tag add -i {frontend_task_id} -t {frontend_tag_id}
```

#### 代码审核流程
| 步骤 | 描述 | 操作手册ID |
| ---- | ---- | ---------- |
| 1 | 查询处于"代码审核"状态的工作项，或筛选带 `arch` / `security` 标签的工作项 | 参考 `references/workitem_guide.md` 组合过滤查询 |
| 2 | 查看工作项详情，结合代码仓库进行审核 | 参考 `references/workitem_guide.md` 查看工作项详情 |
| 3 | 审核通过：将工作项状态流转到"待测试" | 参考 `references/workitem_guide.md` 更新工作项状态 |
| 4 | 审核不通过：退回工作项并在描述中写明审核意见 | 参考 `references/workitem_guide.md` 更新工作项状态 |

Example
```bash
# 步骤 1: 获取"代码审核"和"待测试"等状态 ID
pingcode-client project metadata statuses

# 步骤 2: 查询待审核工作项（arch 标签 + 代码审核状态）
pingcode-client project metadata tags list
pingcode-client project workitem list \
  --project-id {project_id} \
  --tag-ids {arch_tag_id} \
  --state-ids {code_review_state_id}

# 步骤 3: 查看工作项详情
pingcode-client project workitem get -i {work_item_id}

# 步骤 4a: 审核通过 → 流转到"待测试"
pingcode-client project workitem update \
  -i {work_item_id} \
  --state-id {testing_state_id}

# 步骤 4b: 审核不通过 → 退回重新开发
pingcode-client project workitem update \
  -i {work_item_id} \
  --state-id {reopen_state_id} \
  --desc "审核意见：需要补充单元测试以覆盖边界场景，并修复 XSS 风险"
```


#### 开发流程
| 步骤 | 描述 | 操作手册ID |
| ---- | ---- | ---------- |
| 1 | 查询分配给本人且带 `backend` 标签的工作项，确认当前任务 | 参考 `references/workitem_guide.md` 组合过滤查询 |
| 2 | 领取工作项：将状态更新为"进行中" | 参考 `references/workitem_guide.md` 更新工作项状态 |
| 3 | 开发完成后，将工作项状态更新为"代码审核"，等待研发经理审核 | 参考 `references/workitem_guide.md` 更新工作项状态 |

Example
```bash
# 步骤 1: 获取 backend 标签 ID 和"打开"状态 ID
pingcode-client project metadata tags list
pingcode-client project metadata statuses

# 步骤 2: 查询 backend 标签 + 打开状态的工作项
pingcode-client project workitem list \
  --project-id {project_id} \
  --tag-ids {backend_tag_id} \
  --state-ids {open_state_id}

# 步骤 3: 领取任务 → 更新状态为"进行中"
pingcode-client project workitem update \
  -i {work_item_id} \
  --state-id {in_progress_state_id}

# 步骤 4: 开发完成 → 提交代码审核
pingcode-client project workitem update \
  -i {work_item_id} \
  --state-id {code_review_state_id}
```

#### 开发流程
| 步骤 | 描述 | 操作手册ID |
| ---- | ---- | ---------- |
| 1 | 查询分配给本人且带 `frontend` 标签的工作项，确认当前任务 | 参考 `references/workitem_guide.md` 组合过滤查询 |
| 2 | 领取工作项：将状态更新为"进行中" | 参考 `references/workitem_guide.md` 更新工作项状态 |
| 3 | 开发完成后，将工作项状态更新为"代码审核"，等待研发经理审核 | 参考 `references/workitem_guide.md` 更新工作项状态 |

Example
```bash
# 步骤 1: 获取 frontend 标签 ID 和"打开"状态 ID
pingcode-client project metadata tags list
pingcode-client project metadata statuses

# 步骤 2: 查询 frontend 标签 + 打开状态的工作项
pingcode-client project workitem list \
  --project-id {project_id} \
  --tag-ids {frontend_tag_id} \
  --state-ids {open_state_id}

# 步骤 3: 领取任务 → 更新状态为"进行中"
pingcode-client project workitem update \
  -i {work_item_id} \
  --state-id {in_progress_state_id}

# 步骤 4: 开发完成 → 提交代码审核
pingcode-client project workitem update \
  -i {work_item_id} \
  --state-id {code_review_state_id}
```

### QA

#### 测试流程
| 步骤 | 描述 | 操作手册ID |
| ---- | ---- | ---------- |
| 1 | 查询状态为"待测试"的工作项，确认本轮测试范围 | 参考 `references/workitem_guide.md` 组合过滤查询 |
| 2 | 领取测试任务：将状态更新为"测试中" | 参考 `references/workitem_guide.md` 更新工作项状态 |
| 3 | 测试通过：将工作项状态更新为"已完成" | 参考 `references/workitem_guide.md` 更新工作项状态 |
| 4 | 测试不通过：退回工作项，在描述中记录缺陷复现步骤，打 `qa` 标签标记 | 参考 `references/workitem_guide.md` 更新工作项状态 |

Example
```bash
# 步骤 1: 获取"待测试"、"测试中"、"已完成"、"重新打开"等状态 ID
pingcode-client project metadata statuses

# 步骤 2: 查询待测试的工作项
pingcode-client project workitem list \
  --project-id {project_id} \
  --state-ids {testing_state_id}

# 步骤 3: 领取任务 → 更新状态为"测试中"
pingcode-client project workitem update \
  -i {work_item_id} \
  --state-id {in_testing_state_id}

# 步骤 4a: 测试通过 → 标记为"已完成"
pingcode-client project workitem update \
  -i {work_item_id} \
  --state-id {done_state_id}

# 步骤 4b: 测试不通过 → 退回，记录缺陷信息
pingcode-client project workitem update \
  -i {work_item_id} \
  --state-id {reopen_state_id} \
  --desc "缺陷复现步骤：1. 点击 GitHub 登录按钮 2. 授权后跳转 500 错误。期望：成功跳回首页并完成登录"

# 为退回的工作项打 qa 标签方便追踪
pingcode-client project metadata tags list
pingcode-client project workitem tag add -i {work_item_id} -t {qa_tag_id}
```

### DevOps

#### 发布管理流程
| 步骤 | 描述 | 操作手册ID |
| ---- | ---- | ---------- |
| 1 | 查询当前迭代和版本状态，确认本次发布范围 | 参考 `references/project_guide.md` 迭代/版本管理 |
| 2 | 查询所有"已完成"的工作项，确认入版本范围 | 参考 `references/workitem_guide.md` 组合过滤查询 |
| 3 | 创建发布版本，将已完成工作项逐一关联到该版本 | 参考 `references/project_guide.md` 版本管理 |
| 4 | 部署完成后，将版本标记为"已发布"，并更新工作项状态 | 参考 `references/project_guide.md` 版本管理 |
| 5 | 创建下一个迭代，规划下一开发周期 | 参考 `references/project_guide.md` 迭代管理 |

Example
```bash
# 步骤 1: 查看当前迭代和版本
pingcode-client project iteration list --project-id {project_id}
pingcode-client project version list --project-id {project_id}

# 步骤 2: 获取"已完成"状态 ID，查询已完成的工作项
pingcode-client project metadata statuses
pingcode-client project workitem list \
  --project-id {project_id} \
  --state-ids {done_state_id}

# 步骤 3: 创建发布版本
pingcode-client project version create \
  --project-id {project_id} \
  --name "v1.2.0" \
  --desc "本版本包含：OAuth 登录、性能优化" \
  --release-date "2026-04-30"

# 步骤 4: 查询已完成工作项（手动记录后在 PingCode Web 端关联版本，CLI 暂不支持 workitem --version-id）
pingcode-client project workitem list \
  --project-id {project_id} \
  --state-ids {done_state_id}

# 步骤 5: 部署完成 → 更新版本 release-date 确认发布时间（CLI 暂无 --released 标记，以 release-date 为准）
pingcode-client project version update -i {version_id} --release-date "2026-04-30"

# 步骤 6: 创建下一个迭代
pingcode-client project iteration create \
  --project-id {project_id} \
  --name "Sprint 11" \
  --start-date "2026-05-01" \
  --end-date "2026-05-14"
```

#### 部署与监控流程
| 步骤 | 描述 | 操作手册ID |
| ---- | ---- | ---------- |
| 1 | 查询带 `devops` 标签的工作项，确认部署相关任务 | 参考 `references/workitem_guide.md` 组合过滤查询 |
| 2 | 创建部署工作项（type: task），关联到发布版本和当前迭代 | 参考 `references/workitem_guide.md` 创建工作项 |
| 3 | 按部署进展更新工作项状态（进行中 → 已完成） | 参考 `references/workitem_guide.md` 更新工作项状态 |

Example
```bash
# 步骤 1: 查询 devops 标签的工作项
pingcode-client project metadata tags list
pingcode-client project workitem list \
  --project-id {project_id} \
  --tag-ids {devops_tag_id}

# 步骤 2: 创建部署任务工作项
pingcode-client project workitem create \
  --project-id {project_id} \
  --title "部署 v1.2.0 到生产环境" \
  --type-id task \
  --assignee-id {devops_member_id} \
  --sprint-id {iteration_id}
pingcode-client project workitem tag add -i {deploy_task_id} -t {devops_tag_id}

# 步骤 3: 部署开始 → 进行中
pingcode-client project workitem update -i {deploy_task_id} --state-id {in_progress_state_id}

# 步骤 4: 部署完成 → 已完成
pingcode-client project workitem update -i {deploy_task_id} --state-id {done_state_id}
```

### 安全专家

#### 安全审核流程
| 步骤 | 描述 | 操作手册ID |
| ---- | ---- | ---------- |
| 1 | 查询带 `security` 标签的工作项，或处于"代码审核"状态的工作项 | 参考 `references/workitem_guide.md` 组合过滤查询 |
| 2 | 查看工作项详情，结合代码仓库进行安全评审（XSS、SQL注入、权限漏洞等） | 参考 `references/workitem_guide.md` 查看工作项详情 |
| 3 | 审核通过：将工作项状态流转到"待测试" | 参考 `references/workitem_guide.md` 更新工作项状态 |
| 4 | 发现安全问题：退回工作项，创建独立的安全缺陷工作项并打 `security` 标签 | 参考 `references/workitem_guide.md` 创建工作项 |

Example
```bash
# 步骤 1: 获取 security 标签 ID 和相关状态 ID
pingcode-client project metadata tags list
pingcode-client project metadata statuses

# 步骤 2: 查询待安全审核的工作项
pingcode-client project workitem list \
  --project-id {project_id} \
  --tag-ids {security_tag_id} \
  --state-ids {code_review_state_id}

# 步骤 3: 查看工作项详情
pingcode-client project workitem get -i {work_item_id}

# 步骤 4a: 安全审核通过 → 流转到"待测试"
pingcode-client project workitem update \
  -i {work_item_id} \
  --state-id {testing_state_id}

# 步骤 4b: 发现安全问题 → 退回工作项
pingcode-client project workitem update \
  -i {work_item_id} \
  --state-id {reopen_state_id} \
  --desc "安全审核意见：接口未做输入过滤，存在 SQL 注入风险，需修复后重新提交"

# 步骤 5: 创建独立安全缺陷工作项（严重级别）
pingcode-client project workitem create \
  --project-id {project_id} \
  --title "[安全] 用户输入未过滤导致 SQL 注入" \
  --type-id bug \
  --priority-id {urgent_priority_id} \
  --assignee-id {backend_member_id} \
  --parent-id {work_item_id}
pingcode-client project workitem tag add -i {security_bug_id} -t {security_tag_id}
```
