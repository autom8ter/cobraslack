//go:generate godocdown -o README.md

package cobraslack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/slack-go/slack"
	"github.com/spf13/cobra"
	"net/http"
	"strings"
)

// SlackHandler returns an http handler that sets the arguments to the root cobra command from the text of the slash command
// the output of the cobra command is written directly to the response body as a slack message.
// example executing "echo -h" subcommand: curl -X POST "$(host)/command" --data "text=echo -h"
func SlackHandler(cmd *cobra.Command) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		slash, err := slack.SlashCommandParse(r)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to parse command: %s", err), http.StatusBadRequest)
			return
		}
		cmd.SetArgs(strings.Split(slash.Text, " "))
		buf := bytes.NewBuffer(nil)
		cmd.SetOut(buf)
		if err := cmd.ExecuteContext(r.Context()); err != nil {
			http.Error(w, fmt.Sprintf("failed to execute command: %s", err), http.StatusInternalServerError)
			return
		}
		if err := json.NewEncoder(w).Encode(&slack.Msg{
			Text: buf.String(),
		}); err != nil {
			http.Error(w, fmt.Sprintf("failed to encode response: %s", err), http.StatusInternalServerError)
			return
		}
	}
}

// QueryHandler returns an http handler that sets the arguments to the root cobra command from the text of the 'text' query paramater
// the output of the cobra command is written directly to the response body.
// example executing "echo -h" subcommand: curl -X GET "$(host)/command?text=echo%20-h"
func QueryHandler(cmd *cobra.Command) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		exec := r.URL.Query().Get("text")
		cmd.SetArgs(strings.Split(exec, " "))
		cmd.SetOut(w)
		if err := cmd.ExecuteContext(r.Context()); err != nil {
			http.Error(w, fmt.Sprintf("failed to execute command: %s", err), http.StatusInternalServerError)
			return
		}
	}
}
