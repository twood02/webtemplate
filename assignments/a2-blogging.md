---
layout: page
title:  "Assignment 2: Technical Blogging"
permalink: /assignments/technical-blog
---



In this assignment, you will write a blog post on a networking topic that interests you. You should choose a topic that you feel you can cover in enough depth for a 10 minute presentation. While you will **not** be presenting, it is a good gauge for how much detail to cover. The presentation of go routines given in class is a good example.

<blockquote>
**Requirements:** 
 - Must be long enough to be interesting
 - You must write some code or run experiments
 - Present useful information in an understandable way 
 - Present useful information in a visually appealing way
 - Work in a group with 1-3 members
 - *All* members of the group must contribute (i.e., commits cannot all be made by a single user)

**Deadline:** February 20th, 11:59pm

**Problems/Questions?** Post to the #a2blog channel on Slack!
</blockquote>

The instructions below assume you are creating a standard web page. As an alternative, you are allowed to make a "zine" [similar to these](https://wizardzines.com/). In that case, check with the instructor for any special instructions.

## Setting up your site
Your post will eventually appear on the course website, but you can test it locally first. The course website is built using [Markdown format](https://github.com/adam-p/markdown-here/wiki/Markdown-Cheatsheet) files transformed into HTML using [Jekyll](https://jekyllrb.com) and hosted with [GitHub pages](https://pages.github.com/). Your search engine of choice can provide you with plenty of information about each of these technologies.

One member of the group must [fork the class website repository here](https://github.com/gwAdvNet20/gwAdvNet20.github.io) and share your fork with the rest of your group.

Your post should be created in a reasonably named folder within the `/wiki/` directory. Create a markdown file in your directory by pasting in this content:

```
---
layout: page
title:  YOUR TITLE
permalink: /wiki/SHORTNAME/
---

*by:* Name 1 and Name2


A short description of your post goes here.

---

The rest of your post goes here.

```

This provides the jekyll "front matter" needed at the top of each file; fill in the title of your article and a short name that will be used as the URL for your post. This `shortname` should be the same name as your folder.

Proceed to write your post in Markdown format in the remainder of the file.  You can look at other `.md` files in the website's repository to see how they are formatted.  The [Cloud9 tutorial](/c9/) (made from the `c9.md` file) is a good example that includes a mix of text, images, and code segments.  It even shows how to make slideshows if you want to display multiple images together (more info in `examples/slides.md`).

To view your rendered page, you will need to [install Jekyll](https://jekyllrb.com/docs/installation/). Once complete, you should be able to run the `./run.sh` script in the repository's root to generate the site which you can then visit in your browser at [http://localhost:4000/wiki/shortname](http://localhost:4000/wiki/shortname).

You will then submit by making a pull request to merge your blog post back into the class website.

## What to write
Your writeup should not simply be a list of commands. It should:
 - motivate why someone should read your post and care about the topic
 - explain the concepts behind the topic, not simply code segments or commands
 - include detailed examples or experimental results so the reader can understand the topic in depth
 - include visuals that help illustrate the concepts you are trying to explain

We strongly encourage you to get other students/friends to look at your post before you submit it! Get feedback on what is confusing or lacks detail.

You will be graded on both the quality of your material and the quality of your writing.

Your post is likely to be composed of multiple files (one Markdown file, plus images and possibly code files). You should include all of these files in your folder in a well organized manner.


## What To Turn In
 - A blog post for the class made via pull request into the class website.
 - Your pull request should be titled with the title of your post. The PR description should explain how the work was divided and the major contributions of each team member.

If you pass the assignment, your content will be included on the course website and released under the <a href="https://creativecommons.org/licenses/by-nc/4.0/">CC BY-NC 4.0</a> license under your name.  If you have a strong opinion for why you want a different license, alert the instructor.
