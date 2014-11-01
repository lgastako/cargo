# cargo

Cargo is a simple utility for pulling in commonly used assets from near by
source directories.

You type "cargo cult <filename>" and it finds all files with the same name under
a target directory and then picks whichever one is the most common and copies it
to the current working directory.

The idea is you can do something like this:

    ~/src/project_foo $ cargo cult bootstrap.min.js


