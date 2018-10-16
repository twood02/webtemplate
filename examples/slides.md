---
layout: page
title: "Slides Example"
permalink: /slides/
---
<link type="text/css" rel="stylesheet" href="/assets/css/lightslider.min.css" />
<script src="https://ajax.googleapis.com/ajax/libs/jquery/1.11.0/jquery.min.js"></script>
<script src="/assets/js/lightslider.min.js"></script>

<script type="text/javascript">
$(document).ready(function() {
    $("#slider").lightSlider({
        item: 3,
        mode: "slide",
        auto: false,
        loop: false,
        controls: true,
        pager: true,
    });
});
</script>

Here is a slide show:

<ul id="slider">
	<li><img src="http://faculty.cs.gwu.edu/timwood/images/tim1.jpg"></li>
	<li><img src="http://faculty.cs.gwu.edu/timwood/images/tim2.jpg"></li>
	<li><img src="http://faculty.cs.gwu.edu/timwood/images/tim3.jpg"></li>
	<li><img src="http://faculty.cs.gwu.edu/timwood/images/tim4.jpg"></li>
	<li><img src="http://faculty.cs.gwu.edu/timwood/images/tim5.jpg"></li>
</ul>

LightSlider library: [http://sachinchoolur.github.io/lightslider/](http://sachinchoolur.github.io/lightslider/)

Please use this for something other than pictures of me...

To make it work:
	- Copy the script tags from the top of this file to the page with the slideshow, or put them in the template
	- For each slideshow, create a `<ul>` with a `<li>` tag for each "slide". Give the list a unique ID and then use that in the  `$("#slider").lightSlider({` tag in the top script.
	- To make multiple slideshows copy/paste that chunk of code and be careful about braces/parentheses.
	- Check the instructions for configuration options
	- You can use [for loops](https://shopify.github.io/liquid/tags/iteration/) from Liquid Tags if you want.