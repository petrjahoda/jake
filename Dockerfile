FROM scratch
COPY /html html
COPY /js js
COPY /linux /
CMD ["/jake"]