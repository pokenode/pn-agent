# PokeNode Agent

[PokeNode](https://pokenode.com) is a online vps monitoring service.
The pn-agent is a component to be deployed on your vps.
It sends information of host machine to https://api.pokenode.com.

## Requirement

We use systemd to manage pn-agent service, so that your machines are required to support systemd.
There is no additional requirements.

## Collected Data

Here is what pn-agent collect from your machine:

CPU & Memory & Disk

Top Processes

Docker containers
