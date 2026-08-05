taskcli % ./taskcli --help
TaskCLI is a command-line task manager that helps you 
manage your daily tasks efficiently.

Usage:
  taskcli [command]

Available Commands:
  Delete      Delete task
  add         Add a new task
  complete    Mark a task as completed
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  list        List all tasks

Flags:
  -h, --help   help for taskcli

Use "taskcli [command] --help" for more information about a command.



###################
###################

taskcli % ./taskcli add -t "Learn Cobra" -d "Study Cobra Framework"
✓ Task added: Learn Cobra



taskcli % ./taskcli add -t "Task 1" -d "On Diet"
✓ Task added: Task 1




taskcli % ./taskcli list

Your Tasks:
===========
[ ] ID: 1 - Learn Cobra
    Study Cobra Framework
    Created: 2026-08-05 22:05

[ ] ID: 2 - Task 1
    On Diet
    Created: 2026-08-05 22:06





taskcli % ./taskcli complete 1
✓ Task 1 marked as completed!




taskcli % ./taskcli list      

Your Tasks:
===========
[✓] ID: 1 - Learn Cobra
    Study Cobra Framework
    Created: 2026-08-05 22:05

[ ] ID: 2 - Task 1
    On Diet
    Created: 2026-08-05 22:06


taskcli % ./taskcli delete 1 
Task 1 deleted successfully%      



taskcli % ./taskcli list      

Your Tasks:
===========
[ ] ID: 2 - Task 1
    On Diet
    Created: 2026-08-05 22:06