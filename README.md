# PingCode Go Client CLI

PingCode Go Client CLI 是一个基于 Go 语言开发的命令行工具，旨在帮助开发者更方便地与 PingCode API 进行交互。它支持 Client Credentials 认证方式，并涵盖了 PingCode 的多个核心模块。

## 功能特性

- **项目管理 (Project/Agile)**: 支持项目、工作项、迭代、版本、看板的完整 CRUD 操作及元数据管理。
- **产品管理 (Ship)**: 支持产品及其成员、特性、规划、测试用例的管理。
- **知识库管理 (Wiki)**: 支持查看和列出 Wiki 空间及其页面信息。
- **自动化支持**: 易于集成到 CI/CD 流水线中，通过 CLI 快速执行 PingCode 操作。

## 安装

### 前提条件

- 已安装 [Go 1.21+](https://golang.org/dl/)

### 从源码编译

```bash
# 克隆仓库
git clone https://github.com/x2x4com/pingcode-client.git
cd pingcode-client

# 编译所有平台的二进制文件
make build

# 运行本地版本
make run
```

编译后的二进制文件将位于 `dist/` 目录下。

## 认证配置

使用 CLI 前需要配置 PingCode 的 Client ID 和 Client Secret。你可以通过以下两种方式进行认证：

### 1. 环境变量 (推荐)

在你的 shell 配置文件（如 `.bashrc` 或 `.zshrc`）中设置：

```bash
export PINGCODE_CLIENT_ID="你的_CLIENT_ID"
export PINGCODE_CLIENT_SECRET="你的_CLIENT_SECRET"
```

### 2. 命令行参数

在执行命令时直接通过参数提供：

```bash
pingcode-client --client-id="xxx" --client-secret="xxx" [command]
```

## 使用指南

### 基础用法

```bash
pingcode-client [command] --help
```

### 核心模块命令

#### 1. 项目管理 (`project`)

用于管理 PingCode 中的敏捷项目信息。

- **列出项目**: `pingcode-client project list`
- **查看工作项**: `pingcode-client project workitem get --id <ITEM_ID>`
- **创建工作项**: `pingcode-client project workitem create --title "新任务" --type-id <TYPE_ID>`

#### 2. 产品管理 (`ship`)

用于管理和列出 PingCode 中的产品信息。

- **列出产品**: `pingcode-client ship product list`
- **创建产品**: `pingcode-client ship product create --name "新产品" --identifier "PROD-01"`

#### 3. 知识库管理 (`wiki`)

用于列出和查看 PingCode Wiki 空间信息。

- **列出空间**: `pingcode-client wiki`

## 项目结构

- [cmd/](cmd/): 命令行接口定义 (使用 Cobra 框架)
- [internal/app/](internal/app/): 各模块的具体业务逻辑实现
- [internal/pkg/sdk/](internal/pkg/sdk/): PingCode API 的 Go SDK 实现
- [docs/](docs/): 项目相关文档和指南
- [info/](info/): 版本和构建信息

## 扩展文档

更多详细的模块指南和 API 映射，请参考：
- [Ship 模块指南](docs/pingcode-skills/references/ship_guide.md)
- [工作项指南](docs/pingcode-skills/references/workitem_guide.md)
- [技能库文档](docs/pingcode-skills/skills.md)

## 贡献指南

欢迎提交 Issue 和 Pull Request 来改进此项目。
