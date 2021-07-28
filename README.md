# aotw

![Docker Pulls](https://img.shields.io/docker/pulls/utrecht/aotw.svg)

Administrator Of The Week

## Configure sasm

Create a directory:

```bash
mkdir ~/.aotw
```

Create a config file:

```bash
vim ~/.aotw/config.yml
```

Add a Slack token:

```bash
---
admins:
  # aaa
  - code
  # bbb
  - code2
slack_token: someSlackToken
```

Save the file. Update the permissions to ensure that you are the
only one that is able to read and write to the config file:

```bash
chmod 0600 ~/.aotw/config.yml
```

## Send a slack message

[![dockeri.co](https://dockeri.co/image/utrecht/aotw)](https://hub.docker.com/r/utrecht/aotw)

```bash
docker run \
    -v /home/${USER}/.aotw/:/home/aotw/.aotw/ utrecht/aotw:0.1.0 \
    aotw --config /home/aotw/.aotw/config.yml
```
