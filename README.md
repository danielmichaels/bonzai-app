# [Bonzai] Cookiecutter

> A [cookiecutter] for quickly scaffolding [Bonzai] apps

## Usage

To create a new bonzai app using this repository you only need to run the following:

```shell
cookiecutter https://github.com/danielmichaels/bonzai-app
# or gh:danielmichaels/bonzai-app
```
And then answer the prompts. Here's an example run:

```shell
z ❯ cookiecutter https://github.com/danielmichaels/bonzai-app
github_username [danielmichaels]: 
project_name [bonzai app]: demo
cmd_name [app]: 
project_slug [demo]: 
project_description [A [bonzai](https://github.com/rwxrob/bonzai) cookiecutter project for quickly bootstraping new bonzai branches]: 
go_module_path [github.com/danielmichaels/demo]: 
Select go_version:
1 - 1.18
2 - 1.19
Choose from 1, 2 [1]: 2
Select license:
1 - Apache Software License 2.0
2 - MIT
3 - BSD-3
4 - GNU GPL v3.0
Choose from 1, 2, 3, 4 [1]: 1
license_owner [Boaty McBoatFace]: Daniel Michaels
```

This will create a directory called `demo` in the current working directory. All upper case
letters are converted to lowercase and hypens are used instead of spaces.

After `cookiecutter` has run the following output will be printed to the screen detailing
what to do next.

```shell
====================================================================================
Your project `demo` is ready!
The following is a *brief* overview of steps to push code to remote and
how to get your go module working.
- Move to project directory, and initialize a git repository:
    $ cd demo && git init
- Run go mod tidy to pull in dependencies:
    $ go mod tidy
- If you want to update upstream dependencies (optional; recommended)
    $ go get -u
- Check the code works
    $ go run cmd/app/main.go
- Upload initial code to git:
    $ git add -a
    $ git commit -m "Initial commit!"
    $ git remote add origin https://github.com/danielmichaels/demo.git
    $ git push -u origin --all
```

## Requirements

This project needs the python tool [cookiecutter]. 

```shell
pipx install cookiecutter
```

I'd recommend installing it using [pipx] (`pip install --user pipx`). Normal `pip` is
fine too. 

**How can I trust this guy? he's using python!?**

If there was a viable `go` tool with equivalent capabilities I'd use
that, but there isn't and this tool as served me well over the years.

[cookiecutter]: https://cookiecutter.readthedocs.io/en/stable/
[bonzai]: https://github.com/rwxrob/bonzai
[pipx]: https://pypa.github.io/pipx/
