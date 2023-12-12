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
	Use:   "plogger -t=<extractor-key>",
	Short: "Plot your logger output",
	Long:  `Plot your logger output`,
	Run: func(cmd *cobra.Command, args []string) {
		extractorType, err := cmd.Flags().GetString("type")
		if err != nil {
			slog.Error("Error parsing type flag", slog.Any("error", err))
			os.Exit(1)
		}

		p := tea.NewProgram(presentation.InitialModel)
		go readLog(p, extractorType)
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
	rootCmd.Flags().StringP("type", "t", "", "Log extractor type")
}

func readLog(p *tea.Program, extractorType string) {
	extractor, err := extract.GetExtractor(extractorType)
	if err != nil {
		slog.Info("Error getting extractor for key "+extractorType, slog.Any("errro", err))
		os.Exit(1)
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		data, _ := extractor.Extract(line)
		p.Send(data)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
