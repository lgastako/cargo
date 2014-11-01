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

Or maybe I know I have one and only one `mime.types` file somewhere under etc
but I don't feel like hunting for it:

    /root/src/quickhack # cargo from /etc mime.types
    Searching under: /etc
                for: [mime.types]
    Copying: /etc/mime.types
         to: /root/src/quickhack/mime.types

You can use relative paths with "cargo from", like "cargo from .." but for
longer paths if you know you want to search within the subtree 5 levels up,
instead of having to type "cargo from ../../../../.. file" you can use the "up"
shortcut:

    /Users/john/src/new-web-project/a/b/c/d $ cargo up 4 mime.types
    Searching under: /Users/john/src/new-web-project
                for: [mime.types]
    Copying: /Users/john/src/new-web-project/mime.types
         to: /Users/john/src/new-web-project/a/b/c/d/mime.types


## Bugs

- Currently cargo does not follow symbolic links.  This is because it uses go's
  filepath.walk which itself does not follow symbolic links.


## Ideas

- Add "-m --min=<n>   Short-circuit at <n> identical copies."
- Add "--ignore-errors" or similar to keep going when a copy fails
- Add "--confirm" or "--dry-run" or something.
- Add "--move" option
- Globs instead of exact filenames
