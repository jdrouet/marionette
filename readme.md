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

