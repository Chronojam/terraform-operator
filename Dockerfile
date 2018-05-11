FROM scratch

LABEL Maintainer="Calum Gardner <calumgardner23@gmail.com>"

COPY .build/operator /terraform-operator
COPY terraform /terraform

ENTRYPOINT ["/terraform-operator"] 
