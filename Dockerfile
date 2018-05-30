FROM scratch
COPY ./server /
EXPOSE ${PORT}
CMD ["/server"]
