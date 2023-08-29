# Web Template
[Jekyll](https://jekyllrb.com) based template for GW CS.

## How to setup this template

Setup the repository
 - Create an organization for your class
 - Create a blank repository in your organization (do not have it create a readme for you)
   - If you plan to use github pages for hosting the site, you must name it `ORGNAME.github.io`. This guide assumes that, but 
- Tell Github to import the contents of this repository (`https://github.com/twood02/webtemplate`) into your new one 

You now can try to preview your site at `https://orgname.github.io`, however, you may see that the site's formatting is mostly broken. To fix this we need to adjust some settings specific to your site.

Edit the `_config.yml` file either by checking out the repo or through the github web editor
  - Fill in the `title`, `email`, `author`, and `description` tags.  The `description` will appear in the footer of each page, so typically I just list the course number
  - Edit the `baseurl` field:
    - If your website will appear at the root of a domain (e.g., orgname.github.io), set the `baseurl` to `/`
    - If your site will appear in a folder (e.g., https://www2.seas.gwu.edu/~USERNAME/), set the `baseurl` to the path to the site (in this case `/~USERNAME`)
 - Other settings you can probably leave as is
 
If you edited the config file directly on github, github pages will automatically rebuild the site and you should now see a proper layout when you visit `ORGNAME.github.io`.  It may take a few minutes for the website to rebuild.

## How to make edits to the site

The site is built using Jekyll (installation instructions below) which supports markdown and HTML formats. I generally use markdown files and add in HTML where necessary to do more complex formatting.

If you have the repository checked out to your computer, you can test the site locally by running: `./run.sh` in the repo root. This will start a webserver at `localhost:4000` where you can see the site (the link will be different if you adjusted the baseurl field).  As soon as you edit and save a file, the site will automatically rebuild.

If you aren't using GitHub pages to automatically render the site, then you can run `JEKYLL_ENV=production bundle exec jekyll build --incremental` to create a local version of the site in `build/` which you can upload somewhere directly.

## How to install Jekyll

I think all you need to is the following:
  - Make sure Ruby 3.X is installed. By default, MacOS has an outdated version. If you have Homebrew you can run: `brew install ruby`
  - Run `gem install bundler jekyll`
  - Then: `bundle install`
  
Check [https://jekyllrb.com/](https://jekyllrb.com/) for more details.
