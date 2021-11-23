## Patches

## 001-identify-instances-by-name.patch

Find nodes in the cloud by the `Name` tag containing Node privateDNSName.

## 002-non-type-lb.patch

Adds handling for our special LB type `none`, defined by annotation `service.beta.kubernetes.io/aws-load-balancer-type`. It will create only Target Group without LoadBalanacer. Functionality used for creating custom ALBs.


## 003-dont-delete-ingress-sg-rules-elb.patch

We shouldn't delete Ingress SG rule, if it allows access from configured "ElbSecurityGroup", so that we won't disrupt access to Nodes from other ELBs.
