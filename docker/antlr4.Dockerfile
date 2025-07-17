FROM openjdk:17-slim
ARG ANTLR_VERSION=4.13.1

CMD ["bash"]

# Install CURL
RUN apt-get update && \
    apt-get install -y curl && \
    rm -rf /var/lib/apt/lists/*

# Download and install ANTLR

RUN curl -o antlr.jar \
    https://www.antlr.org/download/antlr-${ANTLR_VERSION}-complete.jar

RUN mkdir /usr/local/lib/antlr

RUN mv antlr.jar /usr/local/lib/antlr

RUN echo '#!/bin/bash\njava -jar /usr/local/lib/antlr/antlr.jar "$@"' > /usr/local/bin/antlr

RUN chmod +x /usr/local/bin/antlr

# Install TAR (Probaly should do this in a separate layer)

RUN apt-get update && \
    apt-get install -y tar && \
    rm -rf /var/lib/apt/lists/*

# Download and install GO


RUN echo "export PATH=\$PATH:/usr/local/go/bin" >> ~/.bashrc

RUN curl -o go.tar.gz \
    https://dl.google.com/go/go1.24.5.linux-amd64.tar.gz 

RUN mkdir -p /usr/local/go/bin

RUN tar -xzf go.tar.gz -C /usr/local

RUN bash -c "source ~/.bashrc"

# ANTLR Go runtime

RUN bash -c "go install github.com/antlr/antlr4/runtime/Go/antlr"
