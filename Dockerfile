FROM golang:1.17

RUN mkdir /tmp/app
ADD . /tmp/app
WORKDIR /tmp/app

RUN ssh-keygen -f /tmp/sss -N ''
COPY configdata/config /root/.ssh/config
COPY configdata/id_ed25519-scrollodex.pub /root/.ssh/id_ed25519-scrollodex.pub
COPY configdata/id_ed25519-scrollodex /root/.ssh/id_ed25519-scrollodex
RUN chmod 0640 .ssh/config
RUN chmod 0640 /root/.ssh/id_ed25519-scrollodex.pub
RUN chmod 0600 /root/.ssh/id_ed25519-scrollodex

RUN git config --global pull.rebase false
RUN git config --global pull.ff only

# Used by the filesystem session driver. (but not redis)
RUN mkdir /tmp/sessions

RUN go build -o main .
CMD ["/tmp/app/main"]
EXPOSE 3000
