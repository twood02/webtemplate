---
layout: page
title: Network Forensics
permalink: /forensics/index.html
---

You have been hired to perform forensic analysis for Anarchy-R-Us Inc, a company that believes it has been hacked by a rogue employee, Ann Dercover. 

## Stage 0: Snooping Browsers

To warm up your investigative skills, we captured packet traffic on one of our webserver's to snoop on clients that are fetching files from our server. We captured a trace for a few minutes and recorded several clients downloading files. You are the forensic investigator. Your mission is to investigate the IP addresses below and figure out what they browsed. 

### IP Addresses

```
172.17.0.1
172.17.0.2
172.17.0.3
172.17.0.4
172.17.0.5
172.17.0.6
172.17.0.7
172.17.0.8
172.17.0.9
172.17.0.10
172.17.0.11
172.17.0.12
172.17.0.13
172.17.0.14
172.17.0.15
172.17.0.16
172.17.0.17
172.17.0.18
172.17.0.19
172.17.0.20
172.17.0.21
```

For a given IP address, you should figure out:

1. What is the name of the text files they downloaded?
2. What is in the content of each of the text files?
3. What is the name of the images they downloaded?
4. What is the content of each image?

[Here is your evidence file](evidence05.pcap).  You can download it to Cloud9 with: 

`wget https://gwadvnet20.github.io/forensics/evidence05.pcap`

> These problems are mirrored from the [LMG Network Forensics Puzzle Contest](http://forensicscontest.com). They have been copied here to ensure they remain accessible for students in this class.


## Stage 1: Illicit Messaging


Anarchy-R-Us, Inc. suspects that one of their employees, Ann Dercover, is really a secret agent working for their competitor. Ann has access to the company’s prize asset, the secret recipe. Security staff are worried that Ann may try to leak the company’s secret recipe.

Security staff have been monitoring Ann’s activity for some time, but haven’t found anything suspicious– until now. Today an unexpected laptop briefly appeared on the company wireless network. Staff hypothesize it may have been someone in the parking lot, because no strangers were seen in the building. Ann’s computer, (192.168.1.158) sent IMs over the wireless network to this computer. The rogue laptop disappeared shortly thereafter.

“We have a packet capture of the activity,” said security staff, “but we can’t figure out what’s going on. Can you help?”

You are the forensic investigator. Your mission is to figure out who Ann was IM-ing, what she sent, and recover evidence including:

1. What port number did the IM application communicate over? Is this a standard use of that port number?
1. What is the name of Ann’s IM buddy?
2. What was the first comment in the captured IM conversation?
3. What is the name of the file Ann transferred?
4. What is the magic number of the file you want to extract (first four bytes)?
5. What was the MD5sum of the file?
6. What is the secret recipe?

[Here is your evidence file](evidence01.pcap).  You can download it to Cloud9 with: 

`wget https://gwadvnet20.github.io/forensics/evidence01.pcap`

## Stage 2: Ann Skips Bail

After being released on bail, Ann Dercover disappears! Fortunately, investigators were carefully monitoring her network activity before she skipped town.

“We believe Ann may have communicated with her secret lover, Mr. X, before she left,” says the police chief. “The packet capture may contain clues to her whereabouts.”

You are the forensic investigator. Your mission is to figure out what Ann emailed, where she went, and recover evidence including:

1. What is Ann’s email address?
2. What is Ann’s email password?
3. What is Ann’s secret lover’s email address?
4. What two items did Ann tell her secret lover to bring?
5. What is the NAME of the attachment Ann sent to her secret lover?
6. What is the MD5sum of the attachment Ann sent to her secret lover?
7. In what CITY and COUNTRY is their rendez-vous point?
8. What is the MD5sum of the image embedded in the document?

[Here is your evidence file](evidence02.pcap).

## Stage 3: Ann’s AppleTV
Ann and Mr. X have set up their new base of operations. While waiting for the extradition paperwork to go through, you and your team of investigators covertly monitor her activity. Recently, Ann got a brand new AppleTV, and configured it with the static IP address 192.168.1.10. Here is the packet capture with her latest activity.

You are the forensic investigator. Your mission is to find out what Ann searched for, build a profile of her interests, and recover evidence including:

1. What is the MAC address of Ann’s AppleTV?
2. What User-Agent string did Ann’s AppleTV use in HTTP requests?
3. What were Ann’s first four search terms on the AppleTV (all incremental searches count)?
4. What was the title of the first movie Ann clicked on?
5. What was the full URL to the movie trailer (defined by “preview-url”)?
6. What was the title of the second movie Ann clicked on?
7. What was the price to buy it (defined by “price-display”)?
8. What was the last full term Ann searched for?

[Here is your evidence file](evidence03.pcap).

## Stage 4: The Curious Mr. X
While a fugitive in Mexico, Mr. X remotely infiltrates the Arctic Nuclear Fusion Research Facility’s (ANFRF) lab subnet over the Interwebs. Virtually inside the facility (pivoting through a compromised system), he conducts some noisy network reconnaissance. Sadly, Mr. X is not yet very stealthy.

Unfortunately for Mr. X, the lab’s network is instrumented to capture all traffic (with full content). His activities are discovered and analyzed… by you!

Here is the packet capture containing Mr. X’s activity. As the network forensic investigator, your mission is to answer the following questions:

1. What was the IP address of Mr. X’s scanner?
2. For the FIRST port scan that Mr. X conducted, what type of port scan was it? (Note: the scan consisted of many thousands of packets.) Pick one:
  - TCP SYN
  - TCP ACK
  - UDP
  - TCP Connect
  - TCP XMAS
  - TCP RST
3. What were the IP addresses of the targets Mr. X discovered?
4. What was the MAC address of the Apple system he found?
5. What was the IP address of the Windows system he found?
6. What TCP ports were open on the Windows system? (Please list the decimal numbers from lowest to highest.)

[Here is your evidence file](evidence04.pcap).