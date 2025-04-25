FROM sbtscala/scala-sbt:eclipse-temurin-17.0.14_7_1.10.11_3.3.5

COPY . /app
WORKDIR /app

RUN curl -sSL https://ngrok-agent.s3.amazonaws.com/ngrok.asc \
      | tee /etc/apt/trusted.gpg.d/ngrok.asc >/dev/null \
      && echo "deb https://ngrok-agent.s3.amazonaws.com buster main" \
      | tee /etc/apt/sources.list.d/ngrok.list \
      && apt update \
      && apt install ngrok

ARG NGROK_AUTH_TOKEN
COPY ngrok.yml /root/.ngrok2/ngrok.yml
RUN bash -c ngrok config add-authtoken $NGROK_AUTH_TOKEN --config ./ngrok.yml
EXPOSE 9000