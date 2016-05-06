## Structure of the tasks.yaml
```ymal
	project: what is the overall name of the project
	working_dir: what directory to be used as the working directory
	tasks: # now contains the tasks that should be run concurrently
		- name: what is the name of the of the task
		  command: what command should be executed, can be a path to the executable or just a name in the $PATH env variable
		  args: [arg1, arg2, arg3]
		  dir: can be relative to working_dir or absolute
```
