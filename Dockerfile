FROM golang:1.14

RUN apt update && apt install -y \
  vim

RUN curl https://raw.githubusercontent.com/Shougo/neobundle.vim/master/bin/install.sh | sh
RUN curl https://gist.githubusercontent.com/joshuarubin/03badc5775848289ddc4/raw/1d0255160fb317eeb06bf5d6e2be760c340894ae/.vimrc > /root/.vimrc

RUN vim +NeoBundleInstall +qall
RUN vim +GoInstallBinaries +qall

WORKDIR /go/src/app

COPY . .

RUN go mod init kmode
RUN go build kmode

CMD ["bash"]

