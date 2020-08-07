---
layout: home

---
<div class="wrapper" markdown="0"><div class="footer-col-wrapper">
<div class="footer-col two-col-1">
    <ul class="contact-list">
        <li><a href="https://faculty.cs.gwu.edu/timwood/"><b>Prof. Tim Wood</b></a></li>
        <li><a href="mailto:timwood@gwu.edu">timwood@gwu.edu</a></li>
        <li>Office Hours: By appointment (email/slack me)</li>
        <li>Class: Tuesdays 10AM-12:30PM on Zoom</li>
    </ul>
</div>
<div class="footer-col two-col-2">
    <ul class="contact-list">
        <li><a href="https://www.linkedin.com/in/lucas-ch"><b>Prof. Lucas Chaufournier</b></a></li>
        <li><a href="mailto:lucaschaufournier@gmail.com">lucaschaufournier@gmail.com</a></li>
    </ul>
    </div>
</div></div>

>  The course will be a hands-on introduction to networking, distributed systems, and cloud computing.  The course will be fairly programming intensive (group projects) and you might need to pick up some new languages along the way (C, java, python, and go). <br>
> The course will be co-taught by Prof. Tim Wood and Lucas Chaufournier (GW CS class of 2015, now at Capital One) to give a mix of perspectives from academia and industry.



## Announcements ##
 - The rest of the semester after Spring Break will all be done online. Please contact the instructors if you have any difficulties or concerns! Please continue to reserve 10-12:30 on Tuesdays for this class, although most weeks we will try to have our online meeting shorter than that.

## Class Resources ##
 - [AWS Educate](https://www.awseducate.com) 
 - [Slack Message Board](https://gwadvnet20.slack.com) - [Join here](https://join.slack.com/t/gwadvnet20/shared_invite/enQtODkxMjAwOTE3NTA4LWQ0ZWI0YzAyZjRmYTBmOThjYWEzNWQ2YjcxOGNlZWQzZmEyZDBkNzRlNTVlMTM4MjZlZTViZmM1MDIwNjU4MTc)
 - [gwAdvNet20 GitHub Org](https://github.com/gwAdvNet20)

## Schedule  ##

<div style="font-size:90%">

| Part 1: Networking | Materials |
|:---  |:--- |
| [Network Programming](/slides/1-network-programming.pdf) <br> Weeks 1-3 | **1/14/20:** [Day 1 Survey](https://forms.gle/jPAQtEpsTajsiC4y7) - [Cloud 9 setup](/c9/) <br> **1/21/20:** [HelloInternet Exercise](https://github.com/gwAdvNet20/HelloInternet) (submit PR by 1/27) <br> **1/28/20:** [tshark Wiki](/wiki/tshark)- [Forensics Exercise](/forensics/) <br> **Videos:** [LAN Routing](https://youtu.be/XP61HtbGPbA) - [How DNS Works](https://youtu.be/S8G63sJPPj0) - [HTTP Basics](https://youtu.be/LZJNj-HHfII) - [OSI and TCP Models](https://youtu.be/i9RL5jD9cTI) <br>**Our Videos:** [TCP Reliability](https://expl.ai/YHVVJHG)<br> **Assignments:** [HelloInternet Exercise](https://github.com/gwAdvNet20/HelloInternet) due 1/27/20, [Reliable UDP](/assignments/reliable-udp) due 2/2/20 |
| [Scalability & Performance](/slides/2-scalability-performance.pdf) <br> Weeks 4-5 | **2/4/20:** [Python Select Server](/code/selectserver.py) <br> **2/11/20:** [Jupyter Guide](wiki/jupyter/) <br> **Readings:** [Latency at LinkedIn](https://engineering.linkedin.com/performance/who-moved-my-99th-percentile-latency) <br> **Assignments:** [Code Reviews](/assignments/helloInternet/) due Tuesday 2/18, [Tech Blog](/assignments/technical-blog/) due Thursday 2/20 |
| [High Performance Middleboxes](/slides/3-middleboxes.pdf) <br> Week 6 | **Assignments:** Fixed HelloInternet due Sunday 3/1|

| Part 2: Distributed Systems | Materials
|:---  |:--- |
| [Distributed Systems Basics](/slides/4-dist-sys-intro.pdf) <br>Week 7| **Midterm:** Tuesday Feb 25th - you can bring one sheet of hand written notes. |
| [Scalable Apps & Comm Frameworks](5-microservices.pdf) <br>Weeks 8-9 | **Exercise:** [RestServ](/assignments/httprest) <br> [Quiz 3/10](https://forms.gle/Pn6s8wP8hQG2oitp9)|
| [Spring break!]() <br> Week 10 | |
| [Moving Online & Microservices](/slides/6-going-online.pdf) <br> Week 11 | [Go Resources](/wiki/go/) <br> **Assignment:** [ETL Pipeline](/assignments/etl-pipeline/) due April 5th/April 26th<br> **3/24/20:**  [Class Video](/na/) --- [Office hour ETL walkthrough video](/na/) |
| [Eventual vs Strong Consistency <br> and State Machine Replication](/slides/7-replicated-services-notes.pdf) <br> Week 12 | **3/31/20:** [Class Video](/na/) / [Slide PDF](/slides/7-replicated-services-notes.pdf) --- [Lucas's advice on ETL](/na/) <br> **Reading:**  [Eventually Consistent](https://www.allthingsdistributed.com/2008/12/eventually_consistent.html) by Werner Vogels, Amazon CTO <br> **Reading Quiz:** [Consistency Models Quiz](https://docs.google.com/forms/d/e/1FAIpQLSeS0AzQFawefZDwKYUmT_0202lP_W7XjZDJIUrwcg3KbeYGQw/viewform?usp=sf_link), due May 3rd (see [Tasks](/tasks/) for quiz info)| 
| [Consensus & Fault Tolerance](/slides/8-consensus-annotated.pdf) <br> Week 13 | **4/7/20:** [Class Video](/na/) / [Slide PDF](/slides/8-consensus-annotated.pdf) <br> **Reading:** [Raft Resources](/wiki/raft/) <br> **Reading Quiz:** [Raft Consensus Quiz](https://docs.google.com/forms/d/e/1FAIpQLSda8Ew9m-J3-Dw7V8JSZWoYDQ6wgB-NoVoo-4Gq3piOEmIPzA/viewform?usp=sf_link), due May 3rd (see [Tasks](/tasks/) for quiz info)|
| [Clouds, VMs, and Containers](/slides/9-cloud-vms-containers.pdf) <br> Week 14 | **4/14/20:** [Class Video](/na/) / [Slide PDF](/slides/9-cloud-vms-containers.pdf) <br> **Assignment:** [Raft Leader Election](/assignments/raft-election/) or [Dist. Sys. Tech Blog](/assignments/technical-blog-2/) due April 28th/May 10th | 
| [Distributed Clouds](/slides/10-distributed-clouds.pdf) <br> Week 15 | **4/21/20:** [Slide PDF](/slides/10-distributed-clouds.pdf) / [Class Video](/na/) <br> **Quiz:** [VMs, Containers, Clouds, CDNs, and Serverless](https://docs.google.com/forms/d/e/1FAIpQLSdUsegrnPKAXfCdG6DrYbgaWZkgJstD0s24OowdvHjRgxBCGA/viewform?usp=sf_link) due May <del>3rd</del> 6th (see [Tasks](/tasks/)) |

</div>
