load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "doc.go",
        "fake_auth_client.go",
        "fake_user.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//pkg/apis/auth/v1alpha1:go_default_library",
        "//pkg/client/clientset_generated/clientset/typed/auth/v1alpha1:go_default_library",
        "@io_k8s_apimachinery//pkg/apis/meta/v1:go_default_library",
        "@io_k8s_apimachinery//pkg/labels:go_default_library",
        "@io_k8s_apimachinery//pkg/runtime/schema:go_default_library",
        "@io_k8s_apimachinery//pkg/types:go_default_library",
        "@io_k8s_apimachinery//pkg/watch:go_default_library",
        "@io_k8s_client_go//rest:go_default_library",
        "@io_k8s_client_go//testing:go_default_library",
    ],
)
