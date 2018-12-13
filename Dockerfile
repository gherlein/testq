FROM iron/go
WORKDIR /bin
# Now just add the binary
ADD testq /bin/
ENTRYPOINT ["./testq"]
