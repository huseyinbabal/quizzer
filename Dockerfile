FROM scratch
LABEL org.opencontainers.image.source="https://github.com/huseyinbabal/$module"
ARG module
COPY $module /$module
COPY config.dist.yml /config.yml
ENTRYPOINT ["CONFIG_LOCATION=/config.yml", "/${module}"]

