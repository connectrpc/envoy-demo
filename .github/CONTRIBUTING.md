# Contributing

We'd love your help making `envoy-demo` better.

If you'd like to make a substantial change, please [open an issue][open-issue]
describing your proposal &mdash; discussing large changes ahead of time makes
pull request review much smoother. In your issue, pull request, and any other
communications, please remember to treat your fellow contributors with
respect!

Note that for a contribution to be accepted, you must sign off on all commits
in order to affirm that they comply with the [Developer Certificate of Origin][dco].

## Setup

[Fork][fork], then clone the repository:

```bash
git clone git@github.com:your_github_username/envoy-demo.git connect-envoy-demo
cd connect-envoy-demo
git remote add upstream https://github.com/connectrpc/envoy-demo.git
git fetch upstream
```

You will need an up-to-date installation of the Go programming language to work
on the gRPC server, and an installation of Docker or Podman to with Docker
Compose to run the demo setup.

To start the demo setup, run the default Make target:

```bash
make
```

## Making Changes

Start by creating a new branch for your changes:

```bash
git checkout main
git fetch upstream
git rebase upstream/main
git checkout -b cool_new_feature
```

Make your changes, then ensure that `make` still works.
When you're satisfied with your changes, push them to your fork.

```bash
git commit -a
git push origin cool_new_feature
```

Then use the GitHub UI to open a pull request.

At this point, you're waiting on us to review your changes. We *try* to respond
to issues and pull requests within a few business days, and we may suggest some
improvements or alternatives. Once your changes are approved, one of the
project maintainers will merge them.

We're much more likely to approve your changes if you:

- Document added code.
- Write a [good commit message][commit-message].
- Maintain backward compatibility.

[fork]: https://github.com/connectrpc/envoy-demo/fork
[open-issue]: https://github.com/connectrpc/envoy-demo/issues/new
[dco]: https://developercertificate.org
[commit-message]: http://tbaggery.com/2008/04/19/a-note-about-git-commit-messages.html

