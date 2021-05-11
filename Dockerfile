FROM gcr.io/virtuan-env/virtuan/base:v2

COPY default /etc/nginx/sites-available/default

WORKDIR /root

RUN mkdir -p /root/.ssh

COPY ./services /etc/services.d/

RUN echo 'export PATH=$PATH:/root/go/bin' >> /root/.bashrc
RUN echo 'source  /root/.bashrc'
ENV PATH="/root/go/bin:${PATH}"

ENTRYPOINT ["/init"]

CMD ["sh", "/root/nginxserver.sh"]
