# Go Embed

Embed files in current or sub directories so they can be included in go binary.

Can embed config files and shell scripts and then copy them to correct locations.

## Things TODO

- Embed files
- in `main`, define `init()`. In this function:
  - check if config file or script already exist
  - if they do not, copy the files
