FROM gcr.io/distroless/static

ENV PROJECT_NAME=hap-switch-command
ENTRYPOINT ["/hap-switch-command"]

ARG OS
ARG ARCH
COPY dist/${PROJECT_NAME}_${OS}_${ARCH}/${PROJECT_NAME} /
CMD [ "serve" ]
