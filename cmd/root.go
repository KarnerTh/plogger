package cmd

import (
	"bufio"
	"log/slog"
	"os"

	"github.com/KarnerTh/plogger/extract"
	"github.com/KarnerTh/plogger/presentation"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "plogger",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		p := tea.NewProgram(presentation.InitialModel)
		go readLog(p)
		if _, err := p.Run(); err != nil {
			slog.Error("Error in reading log", slog.Any("error", err))
			os.Exit(1)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.plogger.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func readLog(p *tea.Program) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		data, _ := extract.NewPingExtrator().Extract(line)
		p.Send(data)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
