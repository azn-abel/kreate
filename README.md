# kreate - WIP

kreate makes it easier for beginners to get started with Kubernetes by generating manifest files for common components, such as Pods, Deployments, Services, and NetworkPolicies.

## Getting started
1. Make sure you have the Go programming language installed locally
2. Clone this repository
3. Run `go build -o kreate` from the terminal in your local repo
4. Add the resulting "kreate" binary to your PATH.
5. Run `kreate help`



## Motivation for this project

TL;DR - I wanted to make it *a little less annoying* to generate Kubernetes manifest files.

Writing Kubernetes manifests is hardly ever done from scratch. Two common options for getting started (neither of which are very convenient) are:
- Sifting through the [docs](https://k8s.io/docs) to find a template for the desired component
- Running something like :
`kubectl create deployment sample-deployment --image=nginx --dry-run=client -o yaml > sample-deployment.yaml`

kreate simplifies the syntax of the above command to:
`kreate deployment > sample-deployment.yaml`

While kreate doesn't allow the variety of options that `kubectl create` does (--image, --namespace, etc.), the resulting manifest file can still be easily edited in VSCode, Vim, or any other editor of your choice. Even with all the options that kubectl offers, manual edits are likely needed before applying the manifests in any case.

Additionally, with `kubectl create`, some beginners may struggle with figuring out and remembering which command to run and which options are required for each component. To name a few of these disrepancies:
- Generating pod manifests is done with `kubectl run`, *not* `kubectl create`, which can lead to confusion
- Generating any manifest files requires the `--dry-run=client` and `-o yaml` options, which become annoying to type out every time
- Manifests for some components (such as NetworkPolicies) cannot be created with `kubectl create`.
