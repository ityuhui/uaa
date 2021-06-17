FROM scratch
COPY uaa /
COPY public /public
ENTRYPOINT ["/uaa"]