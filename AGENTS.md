# AGENTS.md

本文档记录了 AI 助手在本项目中执行的关键任务和系统集成工作。

## 自动化任务记录

### 1. BDD 测试集成 (2026-02-26)
- **任务**: 为 `internal` 包添加 BDD 风格的单元测试。
- **工具**: 引入了 `Ginkgo` 和 `Gomega` 框架。
- **覆盖范围**: `internal/config` 和 `internal/git`。
- **验证**: 所有测试已通过 `make test-bdd` 验证。

### 2. 跨平台 Makefile 构建 (2026-02-26)
- **任务**: 创建支持 Windows, Linux, Mac 的构建系统。
- **功能**:
    - 自动 OS 检测。
    - 交叉编译支持 (`build-linux`, `build-darwin`)。
    - 统一的清理和测试接口。

### 3. VS Code 环境标准化 (2026-02-26)
- **任务**: 优化 `.vscode` 目录配置。
- **成果**:
    - `tasks.json`: 与 Makefile 深度绑定。
    - `launch.json`: 提供标准化的调试模板。
    - `settings.json`: 统一 Go 语言开发规范。

## 使用说明

以下为本项目在日常开发与验证中的常用用法摘要：

- 构建
    - `make build`：为当前系统编译可执行文件（Windows 生成 `monolize.exe`）。
    - `make build-linux`：交叉编译 Linux 版。
    - `make build-darwin`：交叉编译 macOS 版。
- 测试与校验
    - `make test`：运行所有单元测试。
    - `make test-bdd`：以 BDD 风格运行测试（基于 Ginkgo/Gomega）。
    - `make lint`：运行 `go vet` 做静态检查。
- 常用 CLI
    - 更新仓库：`monolize update --path <目录或仓库路径> [--path <更多路径>]`
    - 创建 Mono-repo：`monolize create --path <目录或仓库路径> [-p ...] --name <名称> [-n] [--output <输出路径>]`
    - 同步子模块（在 mono-repo 根目录内）：`monolize sync`
- 获取帮助
    - `monolize --help` 或 `monolize <command> --help` 查看命令与参数说明。

## 助手指令参考
本项目旨在保持高内聚、低耦合的 Go 代码风格。在进行后续开发时，请务必：
- 优先更新 `Makefile` 以保持构建一致性。
- 遵循现有的 BDD 测试模式添加新功能测试。
- 确保 `.vscode` 配置的通用性。
