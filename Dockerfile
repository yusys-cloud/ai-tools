FROM scratch
COPY main config.json /
COPY ui /ui/
CMD ["/main"]
