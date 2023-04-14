FROM scratch
ARG module
COPY $module /$module
COPY config.dist.yml /config.yml
ENTRYPOINT ["CONFIG_LOCATION=/config.yml", "/${module}"]

