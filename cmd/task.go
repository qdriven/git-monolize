package cmd

import (
	"fmt"
	"monolize/internal/task"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	taskDir     string
	githubOwner string
	workDir     string
)

var taskCmd = &cobra.Command{
	Use:   "task",
	Short: "Task management commands",
	Long:  `Manage tasks: dispatch tasks to new directories and sync back changes.`,
}

var taskDispatchCmd = &cobra.Command{
	Use:   "dispatch <task-name>",
	Short: "Dispatch a task to a new working directory",
	Long: `Copy a task from the task directory to a new working directory,
initialize git repository and create a GitHub repository.

Example:
  monolize task dispatch my-task --dest ./workspace/my-task`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		taskName := args[0]
		destPath, _ := cmd.Flags().GetString("dest")

		if taskDir == "" {
			return fmt.Errorf("task directory is required, use --task-dir or set in config")
		}
		if githubOwner == "" {
			return fmt.Errorf("github owner is required, use --owner or set in config")
		}

		mgr := task.NewManager(taskDir, githubOwner, workDir)
		return mgr.Dispatch(taskName, destPath)
	},
}

var taskSyncCmd = &cobra.Command{
	Use:   "sync <task-name>",
	Short: "Sync task changes back to the task directory",
	Long: `Copy the task implementation from working directory back to the original task directory.

Example:
  monolize task sync my-task --work-path ./workspace/my-task`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		taskName := args[0]
		workPath, _ := cmd.Flags().GetString("work-path")

		if taskDir == "" {
			return fmt.Errorf("task directory is required, use --task-dir or set in config")
		}

		mgr := task.NewManager(taskDir, githubOwner, workDir)
		return mgr.SyncBack(taskName, workPath)
	},
}

var taskListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks in the task directory",
	Long:  `List all tasks (directories) in the configured task directory.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if taskDir == "" {
			return fmt.Errorf("task directory is required, use --task-dir or set in config")
		}

		mgr := task.NewManager(taskDir, githubOwner, workDir)
		tasks, err := mgr.ListTasks()
		if err != nil {
			return err
		}

		if len(tasks) == 0 {
			fmt.Println("No tasks found.")
			return nil
		}

		fmt.Printf("Tasks in %s:\n", taskDir)
		for _, t := range tasks {
			fmt.Printf("  - %s\n", t)
		}
		return nil
	},
}

func init() {
	taskCmd.PersistentFlags().StringVar(&taskDir, "task-dir", "", "Task directory containing all tasks")
	taskCmd.PersistentFlags().StringVar(&githubOwner, "owner", "", "GitHub owner for creating repositories")
	taskCmd.PersistentFlags().StringVar(&workDir, "work-dir", ".", "Working directory for dispatched tasks")

	viper.BindPFlag("task_dir", taskCmd.PersistentFlags().Lookup("task-dir"))
	viper.BindPFlag("github_owner", taskCmd.PersistentFlags().Lookup("owner"))
	viper.BindPFlag("work_dir", taskCmd.PersistentFlags().Lookup("work-dir"))

	taskDispatchCmd.Flags().String("dest", "", "Destination path for the dispatched task (default: <work-dir>/<task-name>)")
	taskSyncCmd.Flags().String("work-path", "", "Working path of the task to sync (default: <work-dir>/<task-name>)")

	taskCmd.AddCommand(taskDispatchCmd)
	taskCmd.AddCommand(taskSyncCmd)
	taskCmd.AddCommand(taskListCmd)

	rootCmd.AddCommand(taskCmd)
}
