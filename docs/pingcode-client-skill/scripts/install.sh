#!/usr/bin/env bash
# install.sh — 自动下载 pingcode-client 到当前目录 (scripts/)
# 用法: bash scripts/install.sh [VERSION]
#   VERSION 可选，默认下载最新版本，如: bash scripts/install.sh v0.1.1

set -euo pipefail

REPO="x2x4com/pingcode-client"
BINARY_NAME="pingcode-client"
INSTALL_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

# ── 版本解析 ──────────────────────────────────────────────
if [[ "${1:-}" != "" ]]; then
  VERSION="${1}"
else
  echo ">>> 查询最新 Release 版本..."
  VERSION=$(curl -fsSL "https://api.github.com/repos/${REPO}/releases/latest" \
    | grep '"tag_name"' | head -1 | sed 's/.*"tag_name": *"\(.*\)".*/\1/')
fi
echo ">>> 版本: ${VERSION}"

# ── 平台检测 ──────────────────────────────────────────────
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

case "${OS}" in
  linux)  PLATFORM="linux" ;;
  darwin) PLATFORM="darwin" ;;
  *)
    echo "❌ 不支持的操作系统: ${OS}"
    exit 1
    ;;
esac

case "${ARCH}" in
  x86_64 | amd64) ARCH_SUFFIX="amd64" ;;
  arm64 | aarch64) ARCH_SUFFIX="arm64" ;;
  *)
    echo "❌ 不支持的架构: ${ARCH}"
    exit 1
    ;;
esac

# 组合文件名
if [[ "${PLATFORM}" == "linux" ]]; then
  ARTIFACT="${BINARY_NAME}.${ARCH_SUFFIX}"
else
  ARTIFACT="${BINARY_NAME}.${PLATFORM}-${ARCH_SUFFIX}"
fi

DOWNLOAD_URL="https://github.com/${REPO}/releases/download/${VERSION}/${ARTIFACT}"
DEST="${INSTALL_DIR}/${BINARY_NAME}"

# ── 下载 ──────────────────────────────────────────────────
echo ">>> 下载: ${DOWNLOAD_URL}"
curl -fSL --progress-bar "${DOWNLOAD_URL}" -o "${DEST}"
chmod +x "${DEST}"

echo ""
echo "✅ 安装完成: ${DEST}"
echo "   版本信息:"
"${DEST}" --version 2>/dev/null || true
