package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ollama/ollama/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Config struct {
	Test     string    `mapstructure:"test"`
	Profiles []Profile `mapstructure:"profiles"`
}

type Profile struct {
	Name   string `mapstructure:"name"`
	Model  string `mapstructure:"model"`
	Prompt string `mapstructure:"prompt"`
}

var C Config

var v *viper.Viper

var rootCmd = &cobra.Command{
	Use:   "shellama",
	Short: "shellama is a simple utility for interacting with ollama in your shell",
	Args:  cobra.MinimumNArgs(1),
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		// initialize viper
		v = viper.New()

		// set the base name of the config file, without the file extension
		v.SetConfigName(".shellama")

		// add the config file path
		v.AddConfigPath("/etc/shellama/")
		v.AddConfigPath("$HOME")
		v.AddConfigPath(".")

		// read the config file, ignore errors if file doesn't exist
		if err := v.ReadInConfig(); err != nil {
			return err
		}

		// load the config
		if err := v.Unmarshal(&C); err != nil {
			return err
		}

		// select the profile
		// TODO: what if profile doesn't exist?
		profile := cmd.Flag("profile").Value.String()
		for _, p := range C.Profiles {
			if p.Name == profile {
				v.Set("model", p.Model)
				v.Set("prompt", p.Prompt)
				break
			}
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		client, err := api.ClientFromEnvironment()
		if err != nil {
			log.Fatal(err)
		}
		messages := []api.Message{
			api.Message{
				Role:    "system",
				Content: v.GetString("prompt"),
			},
			api.Message{
				Role:    "user",
				Content: strings.Join(args, " "),
			},
		}

		ctx := context.Background()
		req := &api.ChatRequest{
			Model:    v.GetString("model"),
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
	rootCmd.Flags().StringP("profile", "p", "default", "profile to use")
}

func Execute(version string) {
	rootCmd.Version = version
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
