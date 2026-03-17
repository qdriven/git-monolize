# monolize task

任务管理命令：分发任务到新目录并同步回更改。

## 概述

`task` 命令用于管理开发任务的工作流程：
1. **Dispatch**: 将任务从任务目录复制到工作目录，初始化 Git 仓库并创建 GitHub 远程仓库
2. **Sync**: 将工作目录中的任务实现同步回原始任务目录
3. **List**: 列出任务目录中的所有任务

支持交互式终端 UI (TUI) 模式，提供更友好的操作体验。

## 使用方法

```bash
monolize task [command] [flags]
```

## 全局标志

| 标志 | 默认值 | 说明 |
|------|--------|------|
| `--task-dir` | | 任务目录路径（必需） |
| `--owner` | | GitHub 用户名（dispatch 必需） |
| `--work-path` | `.` | 工作目录路径 |
| `--tui` | `false` | 启用交互式终端 UI |

## 子命令

### `task list`

列出任务目录中的所有任务。

```bash
monolize task list --task-dir ./tasks
```

**示例**:

```bash
# 基本列表
monolize task list --task-dir ~/tasks

# TUI 模式
monolize task list --task-dir ~/tasks --tui
```

**输出示例**:

```
Tasks in /Users/you/tasks
#  Task Name
1  feature-auth
2  feature-payment
3  bugfix-login
4  refactor-api
```

### `task dispatch`

将任务分发到新的工作目录。

```bash
monolize task dispatch [task-name] [flags]
```

**参数**:

| 参数 | 必填 | 说明 |
|------|------|------|
| `task-name` | CLI 模式必填 | 任务名称 |

**标志**:

| 标志 | 说明 |
|------|------|
| `--dest` | 目标路径 |
| `--tui` | 启用 TUI 模式交互选择任务 |

**示例**:

```bash
# CLI 模式
monolize task dispatch my-task \
  --task-dir ./tasks \
  --owner myuser \
  --dest ./workspace/my-task

# TUI 模式（交互选择任务）
monolize task dispatch --tui \
  --task-dir ./tasks \
  --owner myuser
```

**工作流程**:

1. **复制任务文件**: 将任务从任务目录复制到工作目录
2. **初始化 Git**: 在工作目录中初始化 Git 仓库
3. **创建 GitHub 仓库**: 使用 `gh repo create` 创建远程仓库
4. **推送代码**: 将代码推送到 GitHub

**输出示例**:

```
Copying task files...
  From: /Users/you/tasks/my-task
  To:   /Users/you/workspace/my-task

Initializing git repository...
  ✓ Done

Creating GitHub repository myuser/my-task...
  ✓ Done

Task dispatched successfully!
  Location: /Users/you/workspace/my-task
  GitHub:   https://github.com/myuser/my-task
```

### `task sync`

将任务实现从工作目录同步回任务目录。

```bash
monolize task sync [task-name] [flags]
```

**参数**:

| 参数 | 必填 | 说明 |
|------|------|------|
| `task-name` | CLI 模式必填 | 任务名称 |

**标志**:

| 标志 | 说明 |
|------|------|
| `--work-path` | 工作目录路径 |
| `--tui` | 启用 TUI 模式交互选择任务 |

**示例**:

```bash
# CLI 模式
monolize task sync my-task \
  --task-dir ./tasks \
  --work-path ./workspace

# TUI 模式
monolize task sync --tui \
  --task-dir ./tasks \
  --work-path ./workspace
```

**工作流程**:

1. **查找工作目录**: 在工作目录中查找任务实现
2. **复制回任务目录**: 将实现复制回原始任务目录

**输出示例**:

```
Syncing task files...
  From: /Users/you/workspace/my-task/task
  To:   /Users/you/tasks/my-task

Task synced successfully!
  Location: /Users/you/tasks/my-task
```

## 目录结构

### 任务目录

```
tasks/
├── feature-auth/        # 任务文件夹
│   ├── README.md
│   └── specs.md
├── feature-payment/
│   └── ...
└── bugfix-login/
    └── ...
```

### 工作目录

```
workspace/
├── feature-auth/        # 分发后的任务
│   ├── .git/           # Git 仓库
│   ├── task/           # 任务文件（复制自任务目录）
│   │   ├── README.md
│   │   └── specs.md
│   └── src/            # 你的实现
└── ...
```

## TUI 模式

使用 `--tui` 标志启用交互式终端 UI：

- **任务选择**: 使用方向键导航，回车选择
- **确认对话框**: 确认或取消操作
- **进度指示器**: 显示操作进度

```bash
monolize task dispatch --tui --task-dir ./tasks --owner myuser
```

## 配置文件

在 `~/.monolize.yaml` 中配置默认值：

```yaml
task-dir: /Users/you/tasks
github-owner: your-username
work-dir: /Users/you/workspace
```

配置后可以简化命令：

```bash
monolize task list
monolize task dispatch my-task
monolize task sync my-task
```

## 前置条件

- **GitHub CLI (`gh`)**: 必须安装并登录
  ```bash
  brew install gh
  gh auth login
  ```

- **Git**: 必须配置用户信息
  ```bash
  git config --global user.name "Your Name"
  git config --global user.email "your@email.com"
  ```

## 故障排除

### GitHub CLI 未安装

```
Error: gh command not found
```

解决: 安装 GitHub CLI

### 未登录 GitHub

```
Error: not logged in to gh
```

解决: 运行 `gh auth login`

### 任务目录不存在

```
Error: task directory is required
```

解决: 使用 `--task-dir` 指定目录或在配置文件中设置

### 任务不存在

```
Error: task not found: my-task
```

解决: 检查任务名称或使用 `task list` 查看可用任务
