FROM scratch

LABEL Maintainer="Calum Gardner <calumgardner23@gmail.com>"

COPY .build/operator /terraform-operator
COPY terraform /usr/local/bin/terraform

ENTRYPOINT ["/terraform-operator"] 
