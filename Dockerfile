FROM    vikings/alpine:latest
LABEL   maintainer=ztao8607@gmail.com
COPY    bin/gotrigger /gotrigger
EXPOSE  80
ENTRYPOINT ["/gotrigger"]