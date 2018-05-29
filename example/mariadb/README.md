# Example container help page for mariadb

Here are the basic steps to create the help.1 and help.md files
from the Dockerfile in this directory:

Generated the help.md file from the Dockerfile with this command:

```
  image-helpgen dockerfile -dockerfile Dockerfile 

```
Manually edited the help.md file to change TODO entries to descriptions
and add host mount points and host port numbers, then ran:

```
  image-helpgen man 
```

To check the edits in the resulting help.1 file, run:

```
   man ./help.1
```
