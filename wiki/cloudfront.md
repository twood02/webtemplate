---
layout: page
title: AWS CloudFront
permalink: /wiki/cloudfront
---
<link type="text/css" rel="stylesheet" href="/assets/css/lightslider.min.css" />
<script src="https://ajax.googleapis.com/ajax/libs/jquery/1.11.0/jquery.min.js"></script>
<script src="/assets/js/ligrun.sh
htslider.min.js"></script>


### A Brief History about CDNs

Before I dive directly into Amazon Cloudfront I thought it would be useful to the reader to familiarize themselves with the origins of Amazon Cloudfront. They are based on Content Delivery Networks or CDNs. This technology has been around for quite some time now. They were first developed in the late 90's to address higher demand for video and audio streaming and the ever growing content on the web. Other technologies such as server farms and hierarchical caching played instrumental roles in paving the ground of building the infrastructure that supported the advent of the internet boom.

### Benefits of CDNs. How does it work?

Now you may be wondering why are CDNs important and how do they benefit me? Before CDNs, we would have to retrieve static data from the origin each time a user requested. Naturally, this created many inefficiency and latency issues which only get exacerbated as more users make request to the server. CDNs, allow us to cache data that are frequently requested. What this means is that when you search for cute pictures of cats you will be able to view those adorable animals at a much greater speeds.

### Amazon Cloudfront History. How does it work?

Amazon launched their own CDN CloudFront in 2008. Amazon's global network have now amassed a staggering 205 data centers which they call 'Edge Locations' in 84 cities across 42 countries. When a user makes a request with Amazon CloudFront, the request will initially be sent to the nearest Edge Location from the user. If the information that they have requested has already been cached at that Edge Location, the user will then receive a response. If the information has yet been cached, Cloudfront will retrieve the network packets from the origin.

For example, if a user located in Singapore makes a request to a web server in California our network packets will have to circumvent the globe before they receive an acknowledgement. If the web server was hosted with Cloudfront the request only has to travel as far as one the nearest edge location located in Kuala Lumpur, Malaysia.


### Amazon Cloudfront vs Regular S3

For the purpose of this blog I have created two sites. One hosted on CloudFront and the other on a regular S3. S3 is a service offered by AWS 'that provides object storage through a web service interface.' In other terms it is where our two sites are going to be hosted.

[Cute Corgi site hosted on S3](http://bendogpicture.s3-website-ap-southeast-1.amazonaws.com/)
[Cute Tea Cup cat site hosted on CloudFront](d2wutah124er3v.cloudfront.net)

Note: I created the buckets in Singapore and I am running our network speed tests in Washington DC. Why did I choose Singapore? Well it is pretty far away so hopefully we should be able to see some considerable performance benefits with CloudFront.

### Conclusion

I hope that after reading this blog you have become just as excited about CloudFront or CDNs as I am. If not maybe learnt something about how infrastructure and CDNs support the ever growing volume of data on the internet and increase the quality of the User Experience. At the very least you got to see some cute animals.


https://aws.amazon.com/cloudfront/features/
https://aws.amazon.com/s3/
