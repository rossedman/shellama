package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ollama/ollama/api"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "shellama",
	Short: "shellama is a simple utility for interacting with ollama in your shell",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := api.ClientFromEnvironment()
		if err != nil {
			log.Fatal(err)
		}
		model, err := cmd.Flags().GetString("model")
		if err != nil {
			log.Fatal(err)
		}

		messages := []api.Message{
			api.Message{
				Role:    "system",
				Content: "Provide very brief responses that only contain code, do not wrap in markdown blocks",
			},
			api.Message{
				Role:    "user",
				Content: strings.Join(args, ""),
			},
		}

		ctx := context.Background()
		req := &api.ChatRequest{
			Model:    model,
			Messages: messages,
		}

		respFunc := func(resp api.ChatResponse) error {
			fmt.Print(resp.Message.Content)
			return nil
		}

		err = client.Chat(ctx, req, respFunc)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.Flags().StringP("model", "m", "llama3.1", "model to use")
}

func Execute(version string) {
	rootCmd.Version = version
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
