FROM alpine

COPY example .
COPY config.json .

CMD ./example