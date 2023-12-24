# kreate - WIP
[![Go Report Card](https://goreportcard.com/badge/github.com/azn-abel/kreate)](https://goreportcard.com/report/github.com/azn-abel/kreate)

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

While kreate doesn't allow the variety of options that `kubectl create` does (--image, --namespace, etc.), manual edits and review are likely needed before applying the manifests regardless, which (in my opinion) defeats the purpose of having so many options.

Additionally, `kubectl create` has some features that can get pretty annoying when it comes to generating manifests:
- Generating pod manifests is done with `kubectl run`, *not* `kubectl create`, which can lead to confusion
- Generating any manifest files requires the `--dry-run=client` and `-o yaml` options, which become annoying to type out every time
- Manifests for some components (such as NetworkPolicies) cannot be created with `kubectl create`.

**To be clear:** `kubectl create` is not a bad option for quickly creating K8s components, it's just not necessarily optimized for generating manifest files, which is sometimes more helpful.
