# cobraslack
--
    import "github.com/autom8ter/cobraslack"


## Usage

#### type SlashCommand

```go
type SlashCommand struct {
}
```


#### func  New

```go
func New(rootCmd *cobra.Command) *SlashCommand
```

#### func (*SlashCommand) ServeHTTP

```go
func (s *SlashCommand) ServeHTTP(w http.ResponseWriter, r *http.Request)
```
