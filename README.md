# Infrastructure management software.

Cod is a tool for managing infrastructure. It will support kubernetes,
traditonal virtual machines and microvms. Written in go of course.

## Salmon

Traditional virtual machine management. Cod so far is based on the Linux
KVM hypervisor. To make life easier for us, we use libvirt. Which is a
set of abstractions over KVM.

## Trout

The infrastructure often needs to communicate with each other. And since
kubernetes is based on scaleablity and redundancy we cannot rely on the
IP address of something. We need a highly capable DNS where we can
modify/add/delete records with API calls.
