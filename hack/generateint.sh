#!/bin/bash


TYPES=$(find pkg/operator/aws/v1alpha1 -type d | awk -F/ '{print $NF}')
for i in $TYPES; do
#	echo "$i.DefaultInformer(resClient): &$i.Handler{},"
	echo "\"github.com/chronojam/terraform-operator/pkg/operator/aws/v1alpha1/$i\""
done
