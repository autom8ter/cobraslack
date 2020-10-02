# cobraslack
--
    import "."


## Usage

#### type SlashCommand

```go
type SlashCommand struct {
}
```

SlashCommand creates a slack slash command that executes a cobra command

#### func  NewSlashCommand

```go
func NewSlashCommand(root *cobra.Command) *SlashCommand
```
NewSlashCommand creates a SlashCommand from the give root cobra command

#### func (*SlashCommand) Handler

```go
func (s *SlashCommand) Handler() http.HandlerFunc
```
Handler returns an http handler that sets the arguments to the root cobra
command from the text of the slash command the output of the cobra command is
written directly to the response body.
