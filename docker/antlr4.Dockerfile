FROM openjdk:17-slim
ARG ANTLR_VERSION=4.13.1

CMD ["bash"]

# Install CURL
RUN apt-get update && \
    apt-get install -y curl && \
    rm -rf /var/lib/apt/lists/*

# Download ANTLR
RUN curl -o antlr.jar \
    https://www.antlr.org/download/antlr-${ANTLR_VERSION}-complete.jar

RUN mkdir /usr/local/lib/antlr

RUN mv antlr.jar /usr/local/lib/antlr


RUN echo '#!/bin/bash\njava -jar /usr/local/lib/antlr/antlr.jar "$@"' > /usr/local/bin/antlr

RUN chmod +x /usr/local/bin/antlr

