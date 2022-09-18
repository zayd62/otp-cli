# """
# This is a skeleton file that can serve as a starting point for a Python
# console script. To run this script uncomment the following lines in the
# ``[options.entry_points]`` section in ``setup.cfg``::

#     console_scripts =
#          fibonacci = otp_cli.skeleton:run

# Then run ``pip install .`` (or ``pip install -e .`` for editable mode)
# which will install the command ``fibonacci`` inside your current environment.

# Besides console scripts, the header (i.e. until ``_logger``...) of this file can
# also be used as template for Python modules.

# Note:
#     This file can be renamed depending on your needs or safely removed if not needed.

# References:
#     - https://setuptools.pypa.io/en/latest/userguide/entry_point.html
#     - https://pip.pypa.io/en/stable/reference/pip_install
# """

import logging

import click
from rich import box
from rich.console import Console
from rich.markdown import Markdown
from rich.table import Table

_logger = logging.getLogger(__name__)
console = Console()

FORMAT_TYPE = ["otpauth"]
FORMAT_DESCRIPTION = ["a file containg a list of otpauth URIs"]


class HelpfulImportCmd(click.Command):
    """click command with custom help string"""

    def format_help(self, ctx, formatter):
        help_text = """
Usage: otp-cli import FORMAT PATH

imports auth codes from a file. has the following arguments:

- `FORMAT`: is the format of the auth codes.
- `PATH`: path to the file

The available formats are in the table below:
        """
        help_md = Markdown(help_text)
        format_table = Table(box=box.ASCII)
        format_table.add_column("Format")
        format_table.add_column("Description")

        for i in range(0, len(FORMAT_TYPE)):
            format_table.add_row(FORMAT_TYPE[i], FORMAT_DESCRIPTION[i])

        console.print(help_md)
        console.print(format_table)


@click.group(invoke_without_command=True)
@click.pass_context
def cli1(ctx):
    if ctx.invoked_subcommand is None:
        click.echo("I was invoked without subcommand")
    else:
        click.echo(f"I am about to invoke {ctx.invoked_subcommand}")


@cli1.command(name="import", cls=HelpfulImportCmd)
@click.argument("format")
@click.argument("path")
def import_from_file(format, path):
    """imports auth codes from a file."""
    click.echo(f"format {format} path {path}")


if __name__ == "__main__":
    cli1()
