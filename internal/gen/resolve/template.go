package resolve

import "github.com/valyala/fasttemplate"

// Args Args
type Args map[string]interface{}

// Tpl Tpl
func Tpl(template string, args Args) string {
	t := fasttemplate.New(template, "[", "]")
	return t.ExecuteString(args)
}

// Res Resolver
var Res = `
func (R Res[structure]) [name]() [type] {
	return R.s.[name]
}`

// ResStr to string
var ResStr = `
func (R Res[structure]) [name]() string {
	return fmt.Sprintf("%d", R.s.[name])
}`

// ResStrP point to string
var ResStrP = `
func (R Res[structure]) [name]() *string {
	str:= fmt.Sprintf("%d", R.s.[name])
	return &str
}`
