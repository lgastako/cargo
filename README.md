# cargo

Cargo is a simple utility for pulling in commonly used assets from near by
source directories.

You type "cargo cult <filename>" and it finds all files with the same name under
a target directory and then picks whichever one is the most common and copies it
to the current working directory.

The idea is that any one of these are equivalent:

1. This will copy files from "up 1 directory" (which is `~/src/foo` in this
   example, "up 2" would be `~src`):

    ~/src/foo/bar $ cargo up 1 bootstrap.min.js

1. This will copy files from the exact path "/var/www":

    ~/src/foo/bar $ cargo at /var/www bootstrap.min.js

1. This will walk up the directory tree from the current working directory until
   it finds a match.  In this case it will be `~/src`:

    ~/src/foo/bar $ cargo from src bootstrap.min.js

And all will three produce output something like:

    Searching under: /Users/john/src
                for: [bootstrap.min.js bootstrap.min.css]
    Copying: /Users/john/src/baz/bif/bootstrap.min.js
         to: /Users/john/src/foo/bar/bootstrap.min.js
    Copying: /Users/john/src/bim/bop/bootstrap.min.css
         to: /Users/john/src/foo/bar/bootstrap.min.css


## Ideas

- Add "-m --min=<n>   Short-circuit at <n> identical copies."
- Add "--ignore-errors" or similar to keep going when a copy fails
- Globs instead of exact filenames
