FROM scratch

COPY auth /bin/

WORKDIR /

ENTRYPOINT ["/bin/auth"]