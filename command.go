//go:generate godocdown -o README.md

package cobraslack

import (
	"fmt"
	"github.com/slack-go/slack"
	"github.com/spf13/cobra"
	"net/http"
	"strings"
)

// SlashCommand creates a slack slash command that executes a cobra command
type SlashCommand struct {
	cmd *cobra.Command
}

// NewSlashCommand creates a SlashCommand from the give root cobra command
func NewSlashCommand(root *cobra.Command) *SlashCommand {
	return &SlashCommand{cmd: root}
}

// Handler returns an http handler that sets the arguments to the root cobra command from the text of the slash command
// the output of the cobra command is written directly to the response body.
func (s *SlashCommand) Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//w.Header().Set("Content-Type", "application/json")
		slash, err := slack.SlashCommandParse(r)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to parse command: %s", err), http.StatusBadRequest)
			return
		}
		s.cmd.SetArgs(strings.Split(slash.Text, " "))
		s.cmd.SetOut(w)
		if err := s.cmd.ExecuteContext(r.Context()); err != nil {
			http.Error(w, fmt.Sprintf("failed to execute command: %s", err), http.StatusInternalServerError)
			return
		}
	}
}
