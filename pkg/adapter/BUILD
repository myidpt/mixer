package(default_visibility = ["//visibility:public"])

load("@io_bazel_rules_go//go:def.bzl", "go_prefix", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "accessLogs.go",
        "adapter.go",
        "applicationLogs.go",
        "builder.go",
        "configError.go",
        "denials.go",
        "lists.go",
        "metrics.go",
        "quotas.go",
        "registrar.go",
    ],
    deps = [
        "@com_github_golang_protobuf//proto:go_default_library",
        "@com_github_google_go_genproto//googleapis/rpc/code:go_default_library",
        "@com_github_google_go_genproto//googleapis/rpc/status:go_default_library",
        "@com_github_hashicorp_go_multierror//:go_default_library",
        "@com_github_istio_api//:mixer/v1/config/descriptor",
    ],
)

go_test(
    name = "small_tests",
    size = "small",
    srcs = ["configError_test.go"],
    library = ":go_default_library",
)
