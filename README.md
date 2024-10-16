# templ-watch-bug
Reproduction of templ issue

The issue happends with you change something while using `templ generate --watch`.
This repository has a simple example of the issue.

## Steps to reproduce

1. Clone this repository
2. Run `task dev` to start the `templ generate --watch` (and go server)
3. Open the browser and go to `http://localhost:8080` (or `http://localhost:7331` for the proxyed version)
4. Uncomment `@aComponent` in `templates/landing.templ` and save the file
5. Reload the browser window and the component will be rendered wrong

## Notes on the setup

I wired up a simple Taskfile to show the version of Taskfile, and how it all runs
in my environment.
