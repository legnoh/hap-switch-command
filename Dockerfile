FROM gcr.io/distroless/static

ENV PROJECT_NAME=hap-switch-command
ENTRYPOINT ["/hap-switch-command"]

ARG OS
ARG ARCH
COPY ${PROJECT_NAME} /${PROJECT_NAME}
CMD [ "serve" ]
