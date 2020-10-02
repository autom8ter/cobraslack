package cobraslack_test

import (
	"github.com/autom8ter/cobraslack"
	"github.com/spf13/cobra"
	"io/ioutil"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"
)

func Test(t *testing.T) {
	cmd := &cobra.Command{
		Use:     "test",
		Short:   "this is a test command",
		Version: "v0.0.0",
	}
	f := ""
	cmd2 := &cobra.Command{
		Use:   "create-file",
		Short: "this is a test command to create a file",
		Run: func(cmd *cobra.Command, args []string) {
			os.Create(f)
		},
	}
	cmd2.Flags().StringVarP(&f, "file", "f", "", "file name to create")
	cmd.AddCommand(cmd2)
	srv := httptest.NewServer(cobraslack.NewSlashCommand(cmd).Handler())
	defer srv.Close()
	endpoint := srv.URL
	client := srv.Client()
	vals := url.Values{}
	//vals["command"] = []string{"/test"}
	vals["text"] = []string{"-v"}
	resp, err := client.PostForm(endpoint, vals)
	if err != nil {
		t.Fatal(err.Error())
	}
	defer resp.Body.Close()
	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err.Error())
	}
	if !strings.Contains(string(response), string(response)) {
		t.Fatalf("invalid response from version request: %s", string(response))
	}
	t.Logf("response = %s", string(response))
}
