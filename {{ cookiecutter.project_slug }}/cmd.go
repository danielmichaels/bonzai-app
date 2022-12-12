package cmd

import (
	"fmt"
	"text/template"

	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/conf"
	"github.com/rwxrob/help"
	"github.com/rwxrob/vars"
)

func init() {
	Z.Conf.SoftInit()
	Z.Vars.SoftInit()
}

var Cmd = &Z.Cmd{
	Name:        `{{ cookiecutter.project_name.strip() }}`,
	Summary:     `Summary of what this command does`,
	Usage:       `Usage of this command`,
	Version:     `v0.0.1`,
	License:     `{{ cookiecutter.license.strip() }}`,
	Source:      `{{ cookiecutter.go_module_path.strip() }}`,
	Description: `
		You can start the description here and wrap it to look nice and it
		will just work.  Descriptions are written in BonzaiMark,
		a simplified combination of CommonMark, "go doc", and text/template
		that uses the Cmd itself as a data source and has a rich set of
		builtin template functions ({% raw %}{{pre "pre"}}{% endraw %}, {% raw %}{{pre "exename"}}{% endraw %},
		{% raw %}{pre "indent"}}{% endraw %}, etc.). There are four block types and four span types in
		BonzaiMark:
		Spans
		    Plain
		    *Italic*
		    **Bold**
		    ***BoldItalic***
		    <Under> (brackets remain)
		Note that on most terminals italic is rendered as underlining and
		depending on how old the terminal, other formatting might not appear
		as expected. If you know how to set LESS_TERMCAP_* variables they
		will be observed when output is to the terminal.
		Blocks
		1. Paragraph
		2. Verbatim (block begins with '    ', never first)
		3. Numbered (block begins with '* ')
		4. Bulleted (block begins with '1. ')
		Currently, a verbatim block must never be first because of the
		stripping of initial white space.
		Templates
		Anything from Cmd that fulfills the requirement to be included in
		a Go text/template may be used. This includes {{ "{{ cmd .Name }}" }}
		and the rest. A number of builtin template functions have also been
		added (such as {{ "indent" }}) which can receive piped input. You
		can add your own functions (or overwrite existing ones) by adding
		your own Dynamic template.FuncMap (see text/template for more about
		Go templates). Note that verbatim blocks will need to indented to work:
		    {% raw %}{{ "{{ dir | indent 4 }}" }}{% endraw %}
		Produces a nice verbatim block:
		    {% raw %}{{ dir | indent 4 }}{% endraw %}
		Note this is different for every user and their specific system. The
		ability to incorporate dynamic data into any help documentation is
		a game-changer not only for creating very consumable tools, but
		creating intelligent, interactive training and education materials
		as well.
		Templates Within Templates
		Sometimes you will need more text than can easily fit within
		a single action. (Actions may not span new lines.) For such things
		defining a template with that text is required and they you can
		include it with the {% raw %}{{pre "template" }}{% endraw %} tag.
		    {% raw %}{{define "long" -}}{% endraw %}
		    Here is something
		    that spans multiple
		    lines that would otherwise be too long for a single action.
		    {% raw %}{{- end}}{% endraw %}
		Something`,
	Commands:    []*Z.Cmd{
	    // external imports below
	    help.Cmd, conf.Cmd, vars.Cmd,

	    // internal commands below
	    Welcome,
	    },
    Dynamic: template.FuncMap{
		"uname": func(_ *Z.Cmd) string { return Z.Out("uname", "-a") },
		"dir":   func() string { return Z.Out("dir") },
	},
}

var Welcome = &Z.Cmd{
    Name:       `welcome`,
    Commands:   []*Z.Cmd{help.Cmd},
	Call: func(_ *Z.Cmd, _ ...string) error {
	    fmt.Println("Welcome to {{ cookiecutter.project_name.strip() }}.")
	    return nil
	},
}
