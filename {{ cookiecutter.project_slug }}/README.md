# {{ cookiecutter.project_name.strip() }}

> {{ cookiecutter.project_description.strip() }}

## Install

This command can be installed as a standalone program or composed into a
Bonzai command tree.

Standalone

```
go install github.com/{{ cookiecutter.github_username.strip() }}/{{ cookiecutter.project_name.strip() }}/cmd/{{ cookiecutter.cmd_name.strip() }}@latest
```

Composed

```go
package z

import (
	Z "github.com/rwxrob/bonzai/z"
	example "{{ cookiecutter.go_module_path }}"
)

var Cmd = &Z.Cmd{
	Name:     `z`,
	Commands: []*Z.Cmd{help.Cmd, example.Cmd, example.BazCmd},
}
```

## Tab Completion

To activate bash completion just use the `complete -C` option from your
`.bashrc` or command line. There is no messy sourcing required. All the
completion is done by the program itself.

```
complete -C {{ cookiecutter.project_name.strip() }} {{ cookiecutter.project_name.strip() }}
```

If you don't have bash or tab completion check use the shortcut
commands instead.

## Embedded Documentation

All documentation (like manual pages) has been embedded into the source
code of the application. See the source or run the program with help to
access it.

## Other Examples

* <https://github.com/rwxrob/z> - the one that started it all by Bonzai's creator.
* <https://github.com/danielmichaels/zet-cmd> - my own zettelkasten bonzai app
