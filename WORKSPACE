git_repository(
    name = "io_bazel_rules_go",
    remote = "https://github.com/bazelbuild/rules_go.git",
    tag = "0.4.3",
)

load("@io_bazel_rules_go//go:def.bzl", "go_repositories", "new_go_repository")

go_repositories()

#=============================================================================

git_repository(
    name = "org_pubref_rules_protobuf",
    remote = "https://github.com/pubref/rules_protobuf.git",
    tag = "v0.7.1",
)

load("@org_pubref_rules_protobuf//go:rules.bzl", "go_proto_repositories")

go_proto_repositories()

#=============================================================================

# for building docker base images
debs = (
    (
        "deb_busybox",
        "83d809a22d765e52390c0bc352fe30e9d1ac7c82fd509e0d779d8289bfc8a53d",
        "http://ftp.us.debian.org/debian/pool/main/b/busybox/busybox-static_1.22.0-9+deb8u1_amd64.deb",
    ),
    (
        "deb_libc",
        "2d8de90c084a26c266fa8efa91564f99b2373a7949caa9a1db83460918e6e832",
        "http://ftp.us.debian.org/debian/pool/main/g/glibc/libc6_2.19-18+deb8u7_amd64.deb",
    ),
    (
        "deb_ca_certificates",
        "f58d646045855277c87f532ea5c18df319e91d9892437880c9a0169b834f1bd8",
        "http://ftp.us.debian.org/debian/pool/main/c/ca-certificates/ca-certificates_20141019+deb8u1_all.deb",
    ),
)

[http_file(
    name = name,
    sha256 = sha256,
    url = url,
) for name, sha256, url in debs]

#=============================================================================
# client-go

new_go_repository(
    name = "io_k8s_client_go",
    commit = "d04a6004ffb19ecd067120ca99aee0176f1de810",
    importpath = "k8s.io/client-go",
)

new_go_repository(
    name = "com_github_PuerkitoBio_purell",
    commit = "8a290539e2e8629dbc4e6bad948158f790ec31f4",
    importpath = "github.com/PuerkitoBio/purell",
)

new_go_repository(
    name = "com_github_PuerkitoBio_urlesc",
    commit = "5bd2802263f21d8788851d5305584c82a5c75d7e",
    importpath = "github.com/PuerkitoBio/urlesc",
)

new_go_repository(
    name = "com_github_coreos_go_oidc",
    commit = "5644a2f50e2d2d5ba0b474bc5bc55fea1925936d",
    importpath = "github.com/coreos/go-oidc",
)

new_go_repository(
    name = "com_github_coreos_pkg",
    commit = "fa29b1d70f0beaddd4c7021607cc3c3be8ce94b8",
    importpath = "github.com/coreos/pkg",
)

new_go_repository(
    name = "com_github_davecgh_go_spew",
    commit = "5215b55f46b2b919f50a1df0eaa5886afe4e3b3d",
    importpath = "github.com/davecgh/go-spew",
)

new_go_repository(
    name = "com_github_docker_distribution",
    commit = "cd27f179f2c10c5d300e6d09025b538c475b0d51",
    importpath = "github.com/docker/distribution",
)

new_go_repository(
    name = "com_github_emicklei_go_restful",
    commit = "09691a3b6378b740595c1002f40c34dd5f218a22",
    importpath = "github.com/emicklei/go-restful",
)

new_go_repository(
    name = "com_github_ghodss_yaml",
    commit = "73d445a93680fa1a78ae23a5839bad48f32ba1ee",
    importpath = "github.com/ghodss/yaml",
)

new_go_repository(
    name = "com_github_go_openapi_jsonpointer",
    commit = "46af16f9f7b149af66e5d1bd010e3574dc06de98",
    importpath = "github.com/go-openapi/jsonpointer",
)

new_go_repository(
    name = "com_github_go_openapi_jsonreference",
    commit = "13c6e3589ad90f49bd3e3bbe2c2cb3d7a4142272",
    importpath = "github.com/go-openapi/jsonreference",
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
    name = "com_github_gogo_protobuf",
    commit = "e18d7aa8f8c624c915db340349aad4c49b10d173",
    importpath = "github.com/gogo/protobuf",
)

new_go_repository(
    name = "com_github_golang_glog",
    commit = "44145f04b68cf362d9c4df2182967c2275eaefed",
    importpath = "github.com/golang/glog",
)

new_go_repository(
    name = "com_github_golang_groupcache",
    commit = "02826c3e79038b59d737d3b1c0a1d937f71a4433",
    importpath = "github.com/golang/groupcache",
)

new_go_repository(
    name = "com_github_golang_protobuf",
    commit = "8616e8ee5e20a1704615e6c8d7afcdac06087a67",
    importpath = "github.com/golang/protobuf",
)

new_go_repository(
    name = "com_github_google_gofuzz",
    commit = "44d81051d367757e1c7c6a5a86423ece9afcf63c",
    importpath = "github.com/google/gofuzz",
)

new_go_repository(
    name = "com_github_howeyc_gopass",
    commit = "3ca23474a7c7203e0a0a070fd33508f6efdb9b3d",
    importpath = "github.com/howeyc/gopass",
)

new_go_repository(
    name = "com_github_imdario_mergo",
    commit = "6633656539c1639d9d78127b7d47c622b5d7b6dc",
    importpath = "github.com/imdario/mergo",
)

new_go_repository(
    name = "com_github_jonboulle_clockwork",
    commit = "72f9bd7c4e0c2a40055ab3d0f09654f730cce982",
    importpath = "github.com/jonboulle/clockwork",
)

new_go_repository(
    name = "com_github_juju_ratelimit",
    commit = "77ed1c8a01217656d2080ad51981f6e99adaa177",
    importpath = "github.com/juju/ratelimit",
)

new_go_repository(
    name = "com_github_mailru_easyjson",
    commit = "d5b7844b561a7bc640052f1b935f7b800330d7e0",
    importpath = "github.com/mailru/easyjson",
)

new_go_repository(
    name = "com_github_pmezard_go_difflib",
    commit = "d8ed2627bdf02c080bf22230dbb337003b7aba2d",
    importpath = "github.com/pmezard/go-difflib",
)

new_go_repository(
    name = "com_github_spf13_pflag",
    commit = "5ccb023bc27df288a957c5e994cd44fd19619465",
    importpath = "github.com/spf13/pflag",
)

new_go_repository(
    name = "com_github_stretchr_testify",
    commit = "e3a8ff8ce36581f87a15341206f205b1da467059",
    importpath = "github.com/stretchr/testify",
)

new_go_repository(
    name = "com_github_ugorji_go",
    commit = "ded73eae5db7e7a0ef6f55aace87a2873c5d2b74",
    importpath = "github.com/ugorji/go",
)

new_go_repository(
    name = "com_google_cloud_go_compute_metadata",
    commit = "3b1ae45394a234c385be014e9a488f2bb6eef821",
    importpath = "cloud.google.com/go/compute/metadata",
)

new_go_repository(
    name = "com_google_cloud_go_internal",
    commit = "3b1ae45394a234c385be014e9a488f2bb6eef821",
    importpath = "cloud.google.com/go/internal",
)

new_go_repository(
    name = "in_gopkg_inf_v0",
    commit = "3887ee99ecf07df5b447e9b00d9c0b2adaa9f3e4",
    importpath = "gopkg.in/inf.v0",
)

new_go_repository(
    name = "in_gopkg_yaml_v2",
    commit = "53feefa2559fb8dfa8d81baad31be332c97d6c77",
    importpath = "gopkg.in/yaml.v2",
)

new_go_repository(
    name = "io_k8s_apimachinery",
    commit = "d90aa2c8531f13b0ca734845934c10dcb6a56ca7",
    importpath = "k8s.io/apimachinery",
)

new_go_repository(
    name = "org_golang_google_appengine",
    commit = "4f7eeb5305a4ba1966344836ba4af9996b7b4e05",
    importpath = "google.golang.org/appengine",
)

new_go_repository(
    name = "org_golang_google_appengine_internal",
    commit = "4f7eeb5305a4ba1966344836ba4af9996b7b4e05",
    importpath = "google.golang.org/appengine/internal",
)

new_go_repository(
    name = "org_golang_google_appengine_internal_app_identity",
    commit = "4f7eeb5305a4ba1966344836ba4af9996b7b4e05",
    importpath = "google.golang.org/appengine/internal/app_identity",
)

new_go_repository(
    name = "org_golang_google_appengine_internal_base",
    commit = "4f7eeb5305a4ba1966344836ba4af9996b7b4e05",
    importpath = "google.golang.org/appengine/internal/base",
)

new_go_repository(
    name = "org_golang_google_appengine_internal_datastore",
    commit = "4f7eeb5305a4ba1966344836ba4af9996b7b4e05",
    importpath = "google.golang.org/appengine/internal/datastore",
)

new_go_repository(
    name = "org_golang_google_appengine_internal_log",
    commit = "4f7eeb5305a4ba1966344836ba4af9996b7b4e05",
    importpath = "google.golang.org/appengine/internal/log",
)

new_go_repository(
    name = "org_golang_google_appengine_internal_modules",
    commit = "4f7eeb5305a4ba1966344836ba4af9996b7b4e05",
    importpath = "google.golang.org/appengine/internal/modules",
)

new_go_repository(
    name = "org_golang_google_appengine_internal_remote_api",
    commit = "4f7eeb5305a4ba1966344836ba4af9996b7b4e05",
    importpath = "google.golang.org/appengine/internal/remote_api",
)

new_go_repository(
    name = "org_golang_x_crypto",
    commit = "d172538b2cfce0c13cee31e647d0367aa8cd2486",
    importpath = "golang.org/x/crypto",
)

new_go_repository(
    name = "org_golang_x_net",
    commit = "e90d6d0afc4c315a0d87a568ae68577cc15149a0",
    importpath = "golang.org/x/net",
)

new_go_repository(
    name = "org_golang_x_oauth2",
    commit = "3c3a985cb79f52a3190fbc056984415ca6763d01",
    importpath = "golang.org/x/oauth2",
)

new_go_repository(
    name = "org_golang_x_sys",
    commit = "8f0908ab3b2457e2e15403d3697c9ef5cb4b57a9",
    importpath = "golang.org/x/sys",
)

new_go_repository(
    name = "org_golang_x_text",
    commit = "2910a502d2bf9e43193af9d68ca516529614eed3",
    importpath = "golang.org/x/text",
)

#=============================================================================
# other deps

new_go_repository(
    name = "com_github_18F_hmacauth",
    commit = "9232a6386b737d7d1e5c1c6e817aa48d5d8ee7cd",
    importpath = "github.com/18F/hmacauth",
)

new_go_repository(
    name = "org_golang_google_api",
    commit = "650535c7d6201e8304c92f38c922a9a3a36c6877",
    importpath = "google.golang.org/api",
)

new_go_repository(
    name = "com_google_cloud_go",
    commit = "dbe4740b523eecbc19b2050f0691772c312aa07b",
    importpath = "cloud.google.com/go",
)

new_go_repository(
    name = "com_github_googleapis_gax_go",
    commit = "8c5154c0fe5bf18cf649634d4c6df50897a32751",
    importpath = "github.com/googleapis/gax-go",
)
