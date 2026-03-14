package task

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
)

type Manager struct {
	TaskDir     string
	GitHubOwner string
	WorkDir     string
}

func NewManager(taskDir, owner, workDir string) *Manager {
	if workDir == "" {
		workDir = "."
	}
	return &Manager{
		TaskDir:     taskDir,
		GitHubOwner: owner,
		WorkDir:     workDir,
	}
}

func (m *Manager) Dispatch(taskName string, destPath string) error {
	srcPath := filepath.Join(m.TaskDir, taskName)
	if _, err := os.Stat(srcPath); os.IsNotExist(err) {
		return fmt.Errorf("task not found: %s", srcPath)
	}

	if destPath == "" {
		destPath = filepath.Join(m.WorkDir, taskName)
	}

	if _, err := os.Stat(destPath); os.IsNotExist(err) {
		if err := os.MkdirAll(destPath, 0755); err != nil {
			return fmt.Errorf("failed to create destination directory: %w", err)
		}
	} else {
		taskSubDir := filepath.Join(destPath, "task")
		if _, err := os.Stat(taskSubDir); os.IsNotExist(err) {
			if err := os.MkdirAll(taskSubDir, 0755); err != nil {
				return fmt.Errorf("failed to create task subdirectory: %w", err)
			}
		}
		destPath = taskSubDir
	}

	fmt.Printf("Copying task from %s to %s\n", srcPath, destPath)
	if err := copyDirectory(srcPath, destPath); err != nil {
		return fmt.Errorf("failed to copy task: %w", err)
	}

	fmt.Println("Initializing git repository...")
	if err := initGitRepo(destPath); err != nil {
		return fmt.Errorf("failed to initialize git: %w", err)
	}

	fmt.Printf("Creating GitHub repository %s/%s...\n", m.GitHubOwner, taskName)
	if err := createGitHubRepo(destPath, m.GitHubOwner, taskName); err != nil {
		return fmt.Errorf("failed to create GitHub repository: %w", err)
	}

	fmt.Printf("\nTask dispatched successfully!\n")
	fmt.Printf("Location: %s\n", destPath)
	fmt.Printf("GitHub: https://github.com/%s/%s\n", m.GitHubOwner, taskName)

	return nil
}

func (m *Manager) SyncBack(taskName string, workPath string) error {
	if workPath == "" {
		workPath = filepath.Join(m.WorkDir, taskName)
	}

	taskInWork := filepath.Join(workPath, "task")
	if _, err := os.Stat(taskInWork); os.IsNotExist(err) {
		taskInWork = workPath
	}

	destPath := filepath.Join(m.TaskDir, taskName)

	if _, err := os.Stat(taskInWork); os.IsNotExist(err) {
		return fmt.Errorf("task working directory not found: %s", taskInWork)
	}

	fmt.Printf("Syncing task from %s to %s\n", taskInWork, destPath)
	if err := copyDirectory(taskInWork, destPath); err != nil {
		return fmt.Errorf("failed to sync task: %w", err)
	}

	fmt.Printf("\nTask synced successfully!\n")
	fmt.Printf("Location: %s\n", destPath)

	return nil
}

func (m *Manager) ListTasks() ([]string, error) {
	entries, err := os.ReadDir(m.TaskDir)
	if err != nil {
		return nil, fmt.Errorf("failed to read task directory: %w", err)
	}

	var tasks []string
	for _, entry := range entries {
		if entry.IsDir() {
			tasks = append(tasks, entry.Name())
		}
	}
	return tasks, nil
}

func copyDirectory(src, dst string) error {
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		dstPath := filepath.Join(dst, relPath)

		if info.IsDir() {
			return os.MkdirAll(dstPath, info.Mode())
		}

		return copyFile(path, dstPath)
	})
}

func copyFile(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	return err
}

func initGitRepo(path string) error {
	cmd := exec.Command("git", "init")
	cmd.Dir = path
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}

	cmd = exec.Command("git", "add", ".")
	cmd.Dir = path
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}

	cmd = exec.Command("git", "commit", "-m", "Initial commit")
	cmd.Dir = path
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func createGitHubRepo(path, owner, name string) error {
	cmd := exec.Command("gh", "repo", "create", fmt.Sprintf("%s/%s", owner, name),
		"--public",
		"--source=.",
		"--push")
	cmd.Dir = path
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
