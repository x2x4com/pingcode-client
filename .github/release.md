# Release 流程说明

## 概述

本项目使用 GitHub Actions 自动构建并发布多平台二进制文件到 GitHub Releases。
发布由 **Git Tag** 驱动，推送符合 `v*` 格式的 tag 即自动触发。

---

## 触发方式

```bash
# 1. 更新 VERSION 文件
echo "0.2.0" > VERSION
git add VERSION
git commit -m "chore: bump version to v0.2.0"

# 2. 打 tag（需与 VERSION 文件保持一致）
git tag v0.2.0

# 3. 推送 tag 触发 workflow
git push origin v0.2.0
```

> **注意**：只有 `v` 开头的 tag（如 `v1.0.0`、`v0.2.1-beta`）才会触发 Release workflow。
> 普通 `git push` 不会触发。

---

## 构建产物

每次发布会并行构建以下 4 个平台的二进制文件：

| 文件名 | 平台 | 架构 | 说明 |
| :----- | :--- | :--- | :--- |
| `pingcode-client.amd64` | Linux | x86_64 | 服务器 / CI 环境常用 |
| `pingcode-client.arm64` | Linux | ARM64 | ARM 服务器 / 树莓派等 |
| `pingcode-client.darwin-amd64` | macOS | Intel | Intel Mac |
| `pingcode-client.darwin-arm64` | macOS | Apple Silicon | M 系列 Mac |

构建时自动注入以下版本信息（可通过 `pingcode-client --version` 查看）：

- `Version` — 读取自 `VERSION` 文件
- `GitHash` — 当前 commit SHA
- `BuildDate` — 构建时间

---

## Workflow 配置

Workflow 文件位于 `.github/workflows/release.yml`，结构分为两个阶段：

### 阶段一：build（矩阵并行）

4 个平台同时构建，使用原生 runner 避免跨编译问题：

- **Linux** 在 `ubuntu-latest` 上构建，`CGO_ENABLED=0`
- **macOS Intel** 在 `macos-13`（Intel runner）上构建，`CGO_ENABLED=1`
- **macOS ARM** 在 `macos-latest`（Apple Silicon runner）上构建，`CGO_ENABLED=1`

### 阶段二：release（依赖 build 全部完成）

收集 4 个构建产物，调用 `softprops/action-gh-release` 创建 Release，并：

- 自动以 tag 名称作为 Release 标题
- 根据 commit 历史自动生成 Release Notes
- 上传所有二进制文件作为 Release Assets

---

## 所需权限

Workflow 使用 `GITHUB_TOKEN`（Actions 自动提供），无需手动配置 Secret。

顶层已声明：

```yaml
permissions:
  contents: write  # 允许创建 Release 和上传 Assets
```

如果仓库开启了**分支保护**或**限制 Actions 写权限**，需在以下位置确认设置：

> `Settings` → `Actions` → `General` → `Workflow permissions` → 选择 **Read and write permissions**

---

## 常见问题

### Q：如何发布预发布版本（Pre-release）？

在 `.github/workflows/release.yml` 的 `Create Release` 步骤中增加 `prerelease: true`：

```yaml
- name: Create Release
  uses: softprops/action-gh-release@v2
  with:
    files: dist/*
    generate_release_notes: true
    prerelease: true   # 标记为预发布
```

也可以通过 tag 名称约定区分（如 `v1.0.0-beta`），配合条件判断自动设置：

```yaml
prerelease: ${{ contains(github.ref_name, '-') }}
```

### Q：如何重新触发同一版本的发布？

先删除已有的 tag 和 Release，再重新推送：

```bash
# 删除本地 tag
git tag -d v0.2.0
# 删除远端 tag
git push origin :refs/tags/v0.2.0
# 重新打 tag 并推送
git tag v0.2.0
git push origin v0.2.0
```

### Q：构建失败如何排查？

1. 进入仓库 `Actions` 页面，找到对应的 workflow run
2. 展开失败的 build job，查看具体报错
3. 常见原因：
   - `go.mod` 中 Go 版本与 runner 不匹配
   - 代码编译错误
   - `VERSION` 文件内容格式问题（应为纯版本号，如 `0.2.0`，不含 `v` 前缀）

### Q：VERSION 文件和 tag 不一致会怎样？

不会报错，但构建出的二进制 `--version` 显示的是 `VERSION` 文件内容，而 Release 页面标题是 tag 名。
**建议保持两者一致**，避免混淆。
