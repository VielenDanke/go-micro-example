# be aware that from custom nexus repositories you have to write <nexus_repository>:<port>/<image_name>
FROM alpine

COPY example .
COPY config.json .

# We can't built our binaries inside private network in docker because of absence internet connection
# We built it before dockerfile build command, and then just copy binary inside (and all necessary files)
CMD ./example