FROM puppeth/blockscout@sha256:54f3ab56488a73e172fcfe236b3a44fce750cf7e19ecf83e89bc8a11f1ad5087

WORKDIR /opt/app

ADD ./start.sh .
RUN chmod +x start.sh
CMD ./start.sh
# CMD bash