# Changelog

All notable changes to this project will be documented in this file.

## [0.2.0] - 2026-03-28

### Added

- **`--output` 全局输出格式参数**
  - 支持 `json`、`yaml`、`markdown`、`table` 四种输出格式
  - 默认为 `table` 表格格式

- **`--raw` 全局原始数据参数**
  - 返回 API 完整原始响应
  - 不经过简化处理

- **输出格式化模块** (`internal/pkg/output/output.go`)
  - 统一的输出格式化处理
  - 支持多种数据结构的格式化

- **测试脚本** (`tests/test_output.sh`)
  - 覆盖所有命令的 `--output` 和 `--raw` 参数测试
  - 76 个测试用例全部通过

### Changed

- **所有模块统一输出模式**
  - `Ship` 模块：产品、需求、工单相关命令
  - `Project` 模块：项目、工作项、迭代、版本、看板、成员相关命令
  - `Wiki` 模块：空间、页面、内容、成员相关命令
  - `TestHub` 模块：测试库、用例、计划、运行相关命令
  - `Directory` 模块：用户、部门相关命令

- **SDK 层与展示层分离**
  - `default.go` 函数返回 SDK 类型
  - 命令层统一使用 `output.FormatAndPrint()` 格式化输出

### Deprecated

- 直接打印输出的方式已废弃
  - 不再使用 `fmt.Printf` / `fmt.Println` 直接输出
  - 所有输出统一经过格式化模块

### Fixed

- 修复了部分命令缺少 `--output` 和 `--raw` 支持的问题

## [0.1.4] - 2026-03-27

### Fixed

- 修复 `ship ideas update` 命令异常问题（添加 `omitempty` 到 struct JSON tags）
- 修复 `/v1/prefix` 重复问题
