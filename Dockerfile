FROM gcr.io/distroless/static

ARG package_name
ENV PKGNAME=${package_name}

COPY ${PKGNAME} /${PKGNAME}

ENTRYPOINT [ "/$PKGNAME" ]
CMD [ "serve" ]
