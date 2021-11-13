FROM alpine

COPY pingbot.out /usr/bin/pingbot

ENTRYPOINT ["/usr/bin/pingbot", "--no-update"]
