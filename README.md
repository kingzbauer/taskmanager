## Structure of the tasks.yaml
```ymal
project: what is the overall name of the project
working_dir: what directory to be used as the working directory
tasks: # now contains the tasks that should be run concurrently
	- name: what is the name of the of the task
	  command: what command should be executed, can be a path to the executable or just a name in the $PATH env variable
	  args: [arg1, arg2, arg3]
	  dir: can be relative to working_dir or absolute
	  tags: [tag1, tag2]
```

#### Example yml file
```yaml
---
project: React applications
NPM:
	command: &NPM npm
    args: &NPMARGS [run, build]  
tasks:
	- name: newJobs
	  command: *NPM
	  args: *NPMARGS
	  dir: "./react-apps/apps/newJobs"
	  tags: [users, new]
					  
	- name: ongoing
	  command: *NPM
	  args: *NPMARGS
	  dir: "./react-apps/apps/ongoing"
	  tags:
		  - users
		  - ongoing
```

`taskmanager --taskfile mytasks.yml --tags "new, users"`
