load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

# Pull in bazel_diff for testing.
http_archive(
    name = "bazel_diff",
    patch_args = ["-p1"],
    patches = ["@//:third_party/patches/bazel-diff-only-just-non-external-rules.patch"],
    sha256 = "bdc3ef2192e9aeb288506e2f348aef57bce5e4facf7374b6185ac5ccdd4a9001",
    strip_prefix = "bazel-diff-3.5.0",
    url = "https://github.com/Tinder/bazel-diff/archive/refs/tags/3.5.0.tar.gz",
)

load("@bazel_diff//:artifacts.bzl", "BAZEL_DIFF_MAVEN_ARTIFACTS")
load("@rules_jvm_external//:defs.bzl", "maven_install")

maven_install(
    name = "bazel_diff_maven",
    artifacts = BAZEL_DIFF_MAVEN_ARTIFACTS,
    fail_if_repin_required = True,
    maven_install_json = "@//:bazel_diff_maven_install.json",
    repositories = [
        "https://repo1.maven.org/maven2",
    ],
)

load("@bazel_diff_maven//:defs.bzl", bazel_diff_pinned_maven_install = "pinned_maven_install")

bazel_diff_pinned_maven_install()
