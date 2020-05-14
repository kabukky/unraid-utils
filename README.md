# unraid-utils
## Purpose
A Go program that will regularly run the utils defined in the program. Tailored to my personal use on my Unraid box.
## Utils
### Cleaner
Will regularly search through directories defined via the `DIRS_TO_CLEAN` env var and delete all empty directories that are older than 24 hours. Multiple directories are defined comma separated, e. g. `DIRS_TO_CLEAN=/tv,/movies`.