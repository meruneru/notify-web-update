#!/usr/bin/bash

export LINE_NOTIFY_TOKEN="Input your token"

sam build
sam local invoke NotifyWebUpdateFunction

