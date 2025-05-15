NAME:
   soapstone - Leave messages tied to commands in the terminal and see them when executing

USAGE:
   soapstone [global options] command [command options]
   warning: WILL execute file from path in the command specified

   e.g.
     soapstone touch example
     soapstone gcc file.c -o file

COMMANDS:
   comment, add, c, a  Add comment to command, [COMMAND] [COMMENT...]
   delete, remove      Remove comment [COMMENT] [INDEX]
   help                Show help information

GLOBAL OPTIONS:
   --noex  Do not execute commands from system path (default: false)
Error: flag: help requested
