# monolize agent

管理多种 AI Agent 的配置文件。

## 概述

`agent` 命令用于查看、编辑和管理多种 AI Agent（如 Claude Code、OpenAI Codex、Kimi CLI、GLM）的配置文件。这简化了在不同 AI 工具之间切换和配置的过程。

## 支持的 Agent

| Agent | 名称 | 配置文件 |
|-------|------|----------|
| `claude-code` | Claude Code | `~/.claude.json`, `~/.claude/settings.json`, `~/.claude/settings.local.json` |
| `codex` | OpenAI Codex | `~/.codex/config.toml` |
| `kimi` | Kimi CLI | `~/.kimi/config.toml` |
| `glm` | GLM (智谱 AI) | `~/.claude.json`, `~/.claude/settings.json` |

## 使用方法

```bash
monolize agent [command]
```

## 子命令

### `agent list`

列出所有支持的 AI Agent 及其配置状态。

```bash
monolize agent list
```

**输出示例**:

```
╔══════════════════════════════════════════════════════════════════╗
║                    Supported AI Agents                           ║
╚══════════════════════════════════════════════════════════════════╝

Agent         Display Name     Config Files                    Status
claude-code   Claude Code      .claude.json                    ✅ Configured
                               .claude/settings.json
                               .claude/settings.local.json
codex         OpenAI Codex     .codex/config.toml              ❌ Not configured
kimi          Kimi CLI         .kimi/config.toml               ✅ Configured
glm           GLM (Zhipu AI)   .claude.json                    ❌ Not configured
                               .claude/settings.json
```

### `agent view`

查看指定 Agent 的配置文件内容。

```bash
monolize agent view <agent>
```

**示例**:

```bash
monolize agent view claude-code
monolize agent view kimi
```

**输出示例**:

```
╔══════════════════════════════════════════════════════════════════╗
║                  Claude Code Configuration                       ║
╚══════════════════════════════════════════════════════════════════╝

File: /Users/you/.claude.json
────────────────────────────────────────────────────────────────────
{
  "api_key": "your-api-key",
  "model": "claude-3-opus-20240229"
}
```

### `agent edit`

编辑指定 Agent 的配置文件。

```bash
monolize agent edit <agent> [config-index] [flags]
```

**参数**:

| 参数 | 说明 |
|------|------|
| `agent` | Agent 名称 |
| `config-index` | 配置文件索引（可选，从 0 开始） |

**标志**:

| 标志 | 说明 |
|------|------|
| `--tui` | 使用交互式终端 UI 选择配置文件 |

**示例**:

```bash
# 编辑第一个配置文件
monolize agent edit claude-code

# 编辑指定索引的配置文件
monolize agent edit claude-code 1

# 使用 TUI 选择要编辑的配置文件
monolize agent edit claude-code --tui
```

**编辑器选择**:

编辑器由 `$EDITOR` 环境变量决定：
- 如果设置了 `$EDITOR`，使用该编辑器
- macOS/Linux 默认使用 `vim`
- Windows 默认使用 `notepad`

## 标志

| 标志 | 适用命令 | 说明 |
|------|----------|------|
| `--tui` | `edit` | 启用交互式终端 UI |

## 完整示例

### 查看所有 Agent 状态

```bash
monolize agent list
```

### 查看并编辑 Claude Code 配置

```bash
# 查看当前配置
monolize agent view claude-code

# 编辑配置
monolize agent edit claude-code --tui
```

### 配置新的 AI Agent

```bash
# 查看需要配置的文件
monolize agent view kimi

# 创建/编辑配置文件
monolize agent edit kimi
```

## 配置文件格式

### Claude Code (`~/.claude.json`)

```json
{
  "api_key": "your-api-key",
  "model": "claude-3-opus-20240229"
}
```

### Codex / Kimi (TOML 格式)

```toml
api_key = "your-api-key"
model = "gpt-4"
```

## 故障排除

### 配置文件不存在

如果配置文件不存在，`view` 命令会报错。使用 `edit` 命令会自动创建目录和文件。

### 权限问题

确保有权限访问 `~` 目录下的配置文件。

### 编辑器问题

如果编辑器无法启动，检查 `$EDITOR` 环境变量：

```bash
echo $EDITOR
export EDITOR=nano  # 或其他编辑器
```
