# cargo

Cargo is a simple utility for pulling in commonly used assets from nearby
directories.

For example, I often just want a copy of bootstrap.min.css and/or jquery.  I
also have a ~/src directory on every machine which has numerous web projects.
So, if I were to start a new one I might do something like this:

    /Users/john/src $ mkdir new-web-project
    /Users/john/src $ cd new-web-project
    /Users/john/src/new-web-project $ cargo within src bootstrap.min.css jquery.min.js
    Searching under: /Users/john/src
                for: [bootstrap.min.css jquery.min.js]
    Copying: /Users/john/src/yapa-search.old.and.busted/resources/public/css/bootstrap.min.css
         to: /Users/john/src/new-web-project/bootstrap.min.css
    Copying: /Users/john/src/whisper/static/js/jquery.min.js
         to: /Users/john/src/new-web-project/jquery.min.js

Cargo will automatically search the entire directory tree specified (in this
case /Users/john/src) and find all copies of any files you want.  It then
selects the most common version of each file using file size.

Or maybe I know I have one and only one `mime.types` file somewhere under etc
but I don't feel like hunting for it:

    /root/src/quickhack # cargo from /etc mime.types
    Searching under: /etc
                for: [mime.types]
    Copying: /etc/mime.types
         to: /root/src/quickhack/mime.types

You can use relative paths with `cargo from`, like `cargo from ..` if you like.
But for longer paths, if you already know you want to search within the subtree
5 levels up, the relatively cumbersome `cargo from ../../../../.. file` can be
avoided via the "up" shortcut:

    /Users/john/src/new-web-project/a/b/c/d/e $ cargo up 5 mime.types
    Searching under: /Users/john/src/new-web-project
                for: [mime.types]
    Copying: /Users/john/src/new-web-project/mime.types
         to: /Users/john/src/new-web-project/a/b/c/d/e/mime.types

(Of course in this situation, I would probably still just "cargo within src"
instead, but you get the idea).

## Install

For now:

    make dependency-install install INSTALL_DIR=/usr/local/bin

## Bugs

- Currently cargo does not follow symbolic links.  This is because it uses go's
  filepath.walk which itself does not follow symbolic links.


## Ideas

- Add "-m --min=<n>   Short-circuit at <n> identical copies."
- Add "--ignore-errors" or similar to keep going when a copy fails
- Add "--confirm" or "--dry-run" or something.
- Add "--move" option
- Add "--md5" option to be more accurate than size (but slower).
- Globs instead of exact filenames


## License

Cargo is licensed under the MIT License.
