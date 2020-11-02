# cobraslack
--
    import "."


## Usage

#### func  QueryHandler

```go
func QueryHandler(cmd *cobra.Command) http.HandlerFunc
```
QueryHandler returns an http handler that sets the arguments to the root cobra
command from the text of the 'text' query paramater the output of the cobra
command is written directly to the response body. example executing "echo -h"
subcommand: curl -X GET "$(host)/command?text=echo%20-h"

#### func  SlackHandler

```go
func SlackHandler(cmd *cobra.Command, verificationToken string) http.HandlerFunc
```
SlackHandler returns an http handler that sets the arguments to the root cobra
command from the text of the slash command the output of the cobra command is
written directly to the response body as a slack message. example executing
"echo -h" subcommand: curl -X POST "$(host)/command" --data "text=echo -h"
