FROM python:3-slim
ADD python3.sh fetch.py /root/
ENV CRED null
ENV TOKEN null
WORKDIR /root/
RUN pip install --upgrade google-api-python-client google-auth-httplib2 google-auth-oauthlib && \
    chmod 770 /root/python3.sh && \
    echo $CRED > /root/credentials.json && \
    echo $TOKEN > /root/token.pickle 
CMD /bin/sh -c /root/python3.sh
