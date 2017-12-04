# Example

This directory houses a [Dockerfile](/example/Dockerfile) that has had the following run over it:

```
$ image-helpgen dockerfile -dockerfile Dockerfile
```

which produced [help.md](/example/help.md) and

```
$ image-helpgen man
```

which produced [help.1](/example/help.1).

The rendered man page loooks like:

![man.png](/example/man.png)

**Note**: Normally it's a good idea to fill in anything that can not be automatically filled in :smile:
