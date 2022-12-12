import os
import re
import shutil
import textwrap
from pathlib import Path

CWD = Path.cwd().absolute()

def generate_license():
    """
    Generates the appropriate license for the template from the given options
    """

    # Get the selected license - the license file to be used will be the capitalized version
    # of the first word in the license name
    selected_license: str = "{{ cookiecutter.license }}".upper()
    license_file: str = re.findall(r"^[a-zA-Z]+", selected_license)[0]

    source = os.path.join(CWD, "_licenses", license_file)
    destination = os.path.join(CWD, "LICENSE")

    # Move the selected license file, remove the `_licenses` directory when done
    shutil.move(source, destination)
    shutil.rmtree(os.path.join(CWD, "_licenses"), ignore_errors=True)
def print_final_instructions():
    """
    Simply prints final instructions for users to follow once they generate a project
    using this template!
    """

    message = """
    ====================================================================================
    Your project `{{ cookiecutter.project_name.strip() }}` is ready!
    The following is a *brief* overview of steps to push code to remote and
    how to get your go module working.
    - Move to project directory, and initialize a git repository:
        $ cd {{ cookiecutter.project_name.strip() }} && git init
    - Run go mod tidy to pull in dependencies:
        $ go mod tidy
    - If you want to update upstream dependencies (optional; recommended)
        $ go get -u
    - Check the code works
        $ go run cmd/{{ cookiecutter.cmd_name.strip() }}/main.go
    - Upload initial code to git:
        $ git add -a
        $ git commit -m "Initial commit!"
        $ git remote add origin https://{{ cookiecutter.go_module_path.strip('/') }}.git
        $ git push -u origin --all
    """

    print(textwrap.dedent(message))

runners = [
    generate_license,
    print_final_instructions,
]

for runner in runners:
    try:
        runner()
    except ValueError as e:
        print(e)
        exit(-10)
