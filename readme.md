# Marionette

## Motivation

The goal is to build a monorepo like a multirepo.

We want to be able to have a repo looking like following

```
- mac
  - app
  - app_helper
- common
  - cmd
    - tool-a
    - tool-b
    - tool-c
- win
  - app
```

with the following dependencies

```
- mac/app
  - common/cmd/tool-a
- win/app
  - common/cmd/tool-a
- common/cmd/tool-a
  - common/cmd/tool-b
  - common/cmd/tool-c
```

And that when we make a change into `mac/app`, we can only rebuild `mac/app`.
Or when we make a change into `common/cmd/tool-b`, we rebuild the projects depending on it.

## Howto use it

First you have to define a project file `marionette.json` in your repo.

```json
{
  "projects": [
    {
      "name": "mac-app",
      "path": "mac/app",
      "dependencies": [
        "tool-a"
      ]
    },
    {
      "name": "win-app",
      "path": "win/app",
      "dependencies": [
        "tool-a"
      ]
    },
    {
      "name": "tool-a",
      "path": "common/cmd/tool-a",
      "dependencies": [
        "tool-b",
        "tool-c"
      ]
    },
    {
      "name": "tool-b",
      "path": "common/cmd/tool-b"
    },
    {
      "name": "tool-c",
      "path": "common/cmd/tool-c"
    },
  ]
}
```

Then when you are on a branch and you wanna know which project you have to rebuild, you just have to do

```bash
# with all the arguments
marionette -context <the root of you repository> -reference <the branch of reference> -config <the path to the configuration file>
# or without
marionette # -context $(pwd) -reference master -config marionette.json
```

And it will give you an ordered list of project you have to rebuild

In the current repo, if you do a change on main.go and that your run the previous command. It should return
```bash
$ marionette -config sample/simple.json
[
  {
    "name": "main",
    "path": "cmd/main.go",
    "dependencies": [
      "model",
      "parser"
    ]
  }
]
```

Now, if you just want to check if one of the project needs to be rebuilt.
```bash
$ marionette -config sample/simple.json -check main
true
```

## TODO

- Add filters on projects not to rebuild a project when you modify the readme

