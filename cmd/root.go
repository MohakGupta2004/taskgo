package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/MohakGupta2004/taskgo/internal/storage"
	"github.com/MohakGupta2004/taskgo/internal/task"
	"github.com/spf13/cobra"
)

var (
	taskManager *task.Manager
	rootCmd     = &cobra.Command{
		Use:   "taskgo",
		Short: "A beautiful CLI Todo List application",
		Long:  `TaskGo is a production-grade CLI Todo List application written in Go.`,
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	storagePath := filepath.Join(home, ".taskgo", "tasks.json")
	store := storage.NewJSONStorage(storagePath)
	taskManager = task.NewManager(store)
}
