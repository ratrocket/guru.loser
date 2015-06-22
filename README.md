# guru.loser

## Workflow...

- Get binary files `assemble` and `serve` (from:
  https://github.com/ratrocket/guru.loser/releases/tag/v0.1.0.  If there's a
  newer release, use its tag.).  Put them in the project `bin/` directory.  (Or
  you can build them with `make progs` if you have the [Go programming
  language](http://golang.org) installed.)

- Modify main.tmpl (html files are generated based on this).

- For local testing, run `make localhost` to generate html based on main.tmpl.
  Don't check these generated files (index.html, and files named by numbers)
  into the master or gh-pages branches -- the site will break!

- Locally, run `./bin/serve` to view site.  Navigate to http://localhost:8080.

- For a non-local (deployable) version of the site, run `make guru.loser`.
  This replaces local paths for assets & images with AWS Cloudfront CDN paths.
  (It will work just fine locally, unless you don't have a network.)

- If you need more control over the `assemble` process, it takes options.
  Generally, running `./bin/assemble [options] > index.html` will re-create all
  html files.  Options are:

  - `-tmpl <file>`: template file (default: main.tmpl)
  - `-json <file>`: json file to drive template (default: json.large)
  - `-img <path-prefix>`: path to prepend to image paths (default: ""; see
    "Notes on images, assets, and paths" below)
  - `-asst <path-prefix>`: similar to `-img` but for assets (default: ""; see
    same note below as for `-img`)
  - Note that `assemble` writes the index.html file to stdout (again, lazy), so
    capture it (`[command here, possibly with options] > index.html`).  The
    "interior" files are written directly to file.

## Notes on images, assets, and paths.

It's assumed that images are in a directory named "img".  You can use the `-img
<path-prefix>` option to tell `assemble` to prepend all image paths (which, at
minimum look like: `img/image-file.jpg`) with a string.  The string can be
local or remote.

The same goes for assets (stylesheets and javascripts): They are assumed to be
in a directories named `assets/stylesheets` and `assets/javascripts`.  This can
be controlled by the `-asst <path-prefix>` option to `assemble`.

### Example

To use a Cloudfront URL (and again, your S3 bucket must be set up with the
right directories for this to work: `BUCKET/img` and
`BUCKET/assets/{stylesheets,javascripts}`):

`./bin/assemble -img http://d1nubnl3w13qxh.cloudfront.net/ -asst http://d1nubnl3w13qxh.cloudfront.net/`

The trailing slash on both options is necessary!  (Because of laziness.)

The above command will create image links that look like:

`http://d1nubnl3w13qxh.cloudfront.net/img/image-file.jpg`

and asset paths like:

`http://d1nubnl3w13qxh.cloudfront.net/assets/javascripts/script.js`

If you want exactly the above `assemble` command, there's a make target for it.
Just execute `make guru.loser`.

## Misc info

- Cropped image sizes:

  - sm is 308x
  - md is 616x
  - lg is 1232x

- Been wanting to check this this instead of (or in addition to) Google
  Analytics: http://get.gaug.es/

## Javascript/CSS packages used:

- browserstate/history.js: https://github.com/browserstate/history.js

- Zepto.js: the aerogel-weight jQuery-compatible JavaScript library:
  http://zeptojs.com/

- Magnific Popup: Responsive jQuery (or Zepto) Lightbox Plugin:
  http://dimsemenov.com/plugins/magnific-popup/

- Grids – Pure: http://purecss.io/grids/
