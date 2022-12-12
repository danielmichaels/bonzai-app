package main

import cmd "{{ cookiecutter.go_module_path.strip() }}"

func main() { cmd.Cmd.Run() }
