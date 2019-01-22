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

## Proposition

We have to declare the architecture of the repository

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
    }
  ]
}
```

The project will create the checksum of each project, check which project changed and then resolve what has to be built.
For example, changing `tool-b` would output the following

```json
[
  {
    "name": "tool-b",
    "path": "common/cmd/tool-b"
  },
  {
    "name": "tool-a",
    "path": "common/cmd/tool-a"
  },
  {
    "name": "mac-app",
    "path": "mac/app"
  },
  {
    "name": "win-app",
    "path": "win/app"
  }
]
```

Now, to know what 