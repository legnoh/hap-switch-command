FROM alpine

ARG package_name
ENV PKGNAME=${package_name}

COPY ${PKGNAME} /${PKGNAME}

ENTRYPOINT [ "/hap-switch-command" ]
CMD [ "serve" ]
