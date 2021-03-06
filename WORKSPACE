git_repository(
    name = "io_bazel_rules_go",
    remote = "https://github.com/bazelbuild/rules_go.git",
    tag = "0.4.1",
)

load("@io_bazel_rules_go//go:def.bzl", "go_repositories", "new_go_repository")

go_repositories()

# for building docker base images
debs = (
    (
        "busybox_deb",
        "83d809a22d765e52390c0bc352fe30e9d1ac7c82fd509e0d779d8289bfc8a53d",
        "http://ftp.us.debian.org/debian/pool/main/b/busybox/busybox-static_1.22.0-9+deb8u1_amd64.deb",
    ),
    (
        "libstdcpp6_deb",
        "f1509bbabd78e89c861de16931aec5988e1215649688fd4f8dfe1af875a7fbef",
	    "http://ftp.us.debian.org/debian/pool/main/g/gcc-4.9/libstdc++6_4.9.2-10_amd64.deb",
    ),
    (
		"libgcc1_deb",
		"a1402290165e8d91b396a33d79580a4501041e92bdb62ef23929a0c207cd9af9",
		"http://ftp.us.debian.org/debian/pool/main/g/gcc-4.9/libgcc1_4.9.2-10_amd64.deb",
    ),
    (
        "libc_deb",
        "2d8de90c084a26c266fa8efa91564f99b2373a7949caa9a1db83460918e6e832",
        "http://ftp.us.debian.org/debian/pool/main/g/glibc/libc6_2.19-18+deb8u7_amd64.deb",
    ),
    (
        "ca_certificates_deb",
        "f58d646045855277c87f532ea5c18df319e91d9892437880c9a0169b834f1bd8",
        "http://ftp.us.debian.org/debian/pool/main/c/ca-certificates/ca-certificates_20141019+deb8u1_all.deb",
    ),
)

[http_file(
    name = name,
    sha256 = sha256,
    url = url,
) for name, sha256, url in debs]

## NETLINK

new_go_repository(
    name = "com_github_vishvananda_netlink",
    commit = "fe3b5664d23a11b52ba59bece4ff29c52772a56b",
    importpath = "github.com/vishvananda/netlink",
)

new_go_repository(
    name = "com_github_vishvananda_netns",
    commit = "54f0e4339ce73702a0607f49922aaa1e749b418d",
    importpath = "github.com/vishvananda/netns",
)

# Deps, per https://github.com/kubernetes/client-go/blob/v2.0.0/Godeps/Godeps.json

## CLIENT-GO DEPS
new_go_repository(
    name = "io_k8s_client_go",
    commit = "e121606b0d09b2e1c467183ee46217fa85a6b672",
    importpath = "k8s.io/client-go",
)

## GLOG

new_go_repository(
    name = "com_github_golang_glog",
    commit = "44145f04b68cf362d9c4df2182967c2275eaefed",
    importpath = "github.com/golang/glog",
)

new_go_repository(
    name = "com_github_spf13_pflag",
    commit = "5ccb023bc27df288a957c5e994cd44fd19619465",
    importpath = "github.com/spf13/pflag",
)

new_go_repository(
    name = "com_github_ghodss_yaml",
    commit = "73d445a93680fa1a78ae23a5839bad48f32ba1ee",
    importpath = "github.com/ghodss/yaml",
)

new_go_repository(
    name = "com_github_ugorji_go",
    commit = "f1f1a805ed361a0e078bb537e4ea78cd37dcf065",
    importpath = "github.com/ugorji/go",
)

new_go_repository(
    name = "com_github_google_gofuzz",
    commit = "bbcb9da2d746f8bdbd6a936686a0a6067ada0ec5",
    importpath = "github.com/google/gofuzz",
)

new_go_repository(
    name = "com_github_gogo_protobuf",
    commit = "e18d7aa8f8c624c915db340349aad4c49b10d173",
    importpath = "github.com/gogo/protobuf",
)

new_go_repository(
    name = "com_github_go_openapi_spec",
    commit = "6aced65f8501fe1217321abf0749d354824ba2ff",
    importpath = "github.com/go-openapi/spec",
)

new_go_repository(
    name = "com_github_go_openapi_swag",
    commit = "1d0bd113de87027671077d3c71eb3ac5d7dbba72",
    importpath = "github.com/go-openapi/swag",
)

new_go_repository(
    name = "in_gopkg_inf_v0",
    commit = "3887ee99ecf07df5b447e9b00d9c0b2adaa9f3e4",
    importpath = "gopkg.in/inf.v0",
)

new_go_repository(
    name = "com_github_emicklei_go_restful",
    commit = "89ef8af493ab468a45a42bb0d89a06fccdd2fb22",
    importpath = "github.com/emicklei/go-restful",
)

new_go_repository(
    name = "org_golang_x_net",
    commit = "e90d6d0afc4c315a0d87a568ae68577cc15149a0",
    importpath = "golang.org/x/net",
)

new_go_repository(
    name = "com_github_go_openapi_jsonreference",
    commit = "13c6e3589ad90f49bd3e3bbe2c2cb3d7a4142272",
    importpath = "github.com/go-openapi/jsonreference",
)

new_go_repository(
    name = "com_github_go_openapi_jsonpointer",
    commit = "46af16f9f7b149af66e5d1bd010e3574dc06de98",
    importpath = "github.com/go-openapi/jsonpointer",
)

new_go_repository(
    name = "com_github_PuerkitoBio_purell",
    commit = "8a290539e2e8629dbc4e6bad948158f790ec31f4",
    importpath = "github.com/PuerkitoBio/purell",
)

new_go_repository(
    name = "com_github_mailru_easyjson",
    commit = "d5b7844b561a7bc640052f1b935f7b800330d7e0",
    importpath = "github.com/mailru/easyjson",
)

new_go_repository(
    name = "org_golang_x_text",
    commit = "2910a502d2bf9e43193af9d68ca516529614eed3",
    importpath = "golang.org/x/text",
)

new_go_repository(
    name = "com_github_PuerkitoBio_urlesc",
    commit = "5bd2802263f21d8788851d5305584c82a5c75d7e",
    importpath = "github.com/PuerkitoBio/urlesc",
)

new_go_repository(
    name = "com_github_docker_distribution",
    commit = "cd27f179f2c10c5d300e6d09025b538c475b0d51",
    importpath = "github.com/docker/distribution",
)

new_go_repository(
    name = "com_github_davecgh_go_spew",
    commit = "5215b55f46b2b919f50a1df0eaa5886afe4e3b3d",
    importpath = "github.com/davecgh/go-spew",
)

new_go_repository(
    name = "com_github_pborman_uuid",
    commit = "ca53cad383cad2479bbba7f7a1a05797ec1386e4",
    importpath = "github.com/pborman/uuid",
)

new_go_repository(
    name = "in_gopkg_yaml_v2",
    commit = "53feefa2559fb8dfa8d81baad31be332c97d6c77",
    importpath = "gopkg.in/yaml.v2",
)

new_go_repository(
    name = "com_github_blang_semver",
    commit = "31b736133b98f26d5e078ec9eb591666edfd091f",
    importpath = "github.com/blang/semver",
)

new_go_repository(
    name = "com_github_juju_ratelimit",
    commit = "77ed1c8a01217656d2080ad51981f6e99adaa177",
    importpath = "github.com/juju/ratelimit",
)

new_go_repository(
    name = "org_golang_x_oauth2",
    commit = "3c3a985cb79f52a3190fbc056984415ca6763d01",
    importpath = "golang.org/x/oauth2",
)

new_go_repository(
    name = "com_github_coreos_go_oidc",
    commit = "5644a2f50e2d2d5ba0b474bc5bc55fea1925936d",
    importpath = "github.com/coreos/go-oidc",
)

new_go_repository(
    name = "com_github_jonboulle_clockwork",
    commit = "72f9bd7c4e0c2a40055ab3d0f09654f730cce982",
    importpath = "github.com/jonboulle/clockwork",
)

new_go_repository(
    name = "com_github_coreos_pkg",
    commit = "fa29b1d70f0beaddd4c7021607cc3c3be8ce94b8",
    importpath = "github.com/coreos/pkg",
)

new_go_repository(
    name = "com_google_cloud_go",
    commit = "3b1ae45394a234c385be014e9a488f2bb6eef821",
    importpath = "cloud.google.com/go",
)
