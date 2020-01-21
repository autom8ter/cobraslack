//go:generate godocdown -o README.md

package cobraslack

import (
	"github.com/nlopes/slack"
	"github.com/spf13/cobra"
	"net/http"
	"strings"
)

type SlashCommand struct {
	rootCmd *cobra.Command
}

func New(rootCmd *cobra.Command) *SlashCommand {
	return &SlashCommand{
		rootCmd: rootCmd,
	}
}

func (s *SlashCommand) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	slash, err := slack.SlashCommandParse(r)
	if err != nil {
		http.Error(w, "failed to parse slash command", http.StatusBadRequest)
		return
	}
	in := strings.NewReader(slash.Text)
	s.rootCmd.SetIn(in)
	s.rootCmd.SetOut(w)
	if err := s.rootCmd.Execute(); err != nil {
		http.Error(w, "failed to execute command", http.StatusBadRequest)
		return
	}
	return
}
