FROM ubuntu:22.04

RUN apt-get update && apt-get install -y python3.10 python3-pip openjdk-8-jdk curl unzip zip
RUN curl -sS https://get.sdkman.io | bash

ENV JAVA_HOME=/usr/lib/jvm/java-8-openjdk-amd64
ENV PATH="$JAVA_HOME/bin:$PATH"
ENV SDKMAN_DIR="/root/.sdkman"
ENV PATH="$SDKMAN_DIR/candidates/gradle/current/bin:$PATH"

RUN bash -c "source $SDKMAN_DIR/bin/sdkman-init.sh && sdk install kotlin && sdk install gradle"

WORKDIR /app

COPY .. .

RUN bash -c "source $SDKMAN_DIR/bin/sdkman-init.sh && gradle build"

CMD ["gradle", "run"]
