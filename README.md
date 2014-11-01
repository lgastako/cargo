# cargo

Cargo is a simple utility for pulling in commonly used assets from near by
source directories.

You type "cargo cult <filename>" and it finds all files with the same name under
a target directory and then picks whichever one is the most common and copies it
to the current working directory.

The idea is you can do something like this:

    ~/src/project_foo $ cargo cult bootstrap.min.js
    Searching for 'bootstrap.min.js' to clone ...
    Copying '/Users/john/src/foo/public/js/bootstrap.min.js' to '/Users/john/src/bar/bootstrap.min.js'.
## Ideas

- Add "-d --dir=<dir>" option to, instead of going up 1 dir, search <dir>.
- Add "-u --upto=<dir>" option to, instead of going up 1 dir, go up until you encounter <dir>.
- Add "-l --levels=<n>" option to, instead of going up 1 dir, go up N levels.
- Add "-r --race=<n>" option to copy the first file that you find <n> copies of.
- Add .cargo.ini or similar to store defaults, so you could eg default to "--upto=src"

