# TODO: Move entirely to bazel?
.PHONY: images

DOCKER_REGISTRY?=kopeio
DOCKER_TAG=1.0.20170619

all: images

gofmt:
	gofmt -w -s cmd/
	gofmt -w -s pkg/

portal:
	cd webapp/portal; npm run build
	gzip --force --keep --best webapp/portal/build/static/js/main.*.js

portal-image:
	bazel run //images:auth-portal
	docker tag bazel/images:auth-portal ${DOCKER_REGISTRY}/auth-portal:${DOCKER_TAG}

portal-push: portal-image
	docker push ${DOCKER_REGISTRY}/auth-portal:${DOCKER_TAG}

portal-bounce:
	kubectl delete pod -n kopeio-auth -l app=auth-portal

api-image:
	bazel run //images:auth-api
	docker tag bazel/images:auth-api ${DOCKER_REGISTRY}/auth-api:${DOCKER_TAG}

api-push: api-image
	docker push ${DOCKER_REGISTRY}/auth-api:${DOCKER_TAG}

api-bounce:
	kubectl delete pod -n kopeio-auth -l app=auth-api

push: portal-push api-push
	echo "pushed images"

images: portal-image api-image
	echo "built images"

# -----------------------------------------------------
# api machinery regenerate

apimachinery:
	#./hack/make-apimachinery.sh
	${GOPATH}/bin/conversion-gen --skip-unsafe=true --input-dirs kope.io/auth/pkg/apis/auth/v1alpha1 --v=0  --output-file-base=zz_generated.conversion
	${GOPATH}/bin/conversion-gen --skip-unsafe=true --input-dirs kope.io/auth/pkg/apis/componentconfig/v1alpha1 --v=0  --output-file-base=zz_generated.conversion
	${GOPATH}/bin/defaulter-gen --input-dirs kope.io/auth/pkg/apis/auth/v1alpha1 --v=0  --output-file-base=zz_generated.defaults
	${GOPATH}/bin/defaulter-gen --input-dirs kope.io/auth/pkg/apis/componentconfig/v1alpha1 --v=0  --output-file-base=zz_generated.defaults
	${GOPATH}/bin/deepcopy-gen --input-dirs kope.io/auth/pkg/apis/auth/v1alpha1 --v=0  --output-file-base=zz_generated.deepcopy
	${GOPATH}/bin/deepcopy-gen --input-dirs kope.io/auth/pkg/apis/componentconfig/v1alpha1 --v=0  --output-file-base=zz_generated.deepcopy
	#go install github.com/ugorji/go/codec/codecgen
	# codecgen works only if invoked from directory where the file is located.
	#cd pkg/apis/kops/v1alpha2/ && ~/k8s/bin/codecgen -d 1234 -o types.generated.go instancegroup.go cluster.go federation.go
	#cd pkg/apis/kops/v1alpha1/ && ~/k8s/bin/codecgen -d 1234 -o types.generated.go instancegroup.go cluster.go federation.go
	#cd pkg/apis/kops/ && ~/k8s/bin/codecgen -d 1234 -o types.generated.go instancegroup.go cluster.go federation.go
	${GOPATH}/bin/client-gen  --input-base kope.io/auth/pkg/apis --input="auth/,auth/v1alpha1,componentconfig/,componentconfig/v1alpha1" --clientset-path kope.io/auth/pkg/client/clientset_generated/
	${GOPATH}/bin/client-gen  --clientset-name="clientset" --input-base kope.io/auth/pkg/apis --input="auth/v1alpha1,componentconfig/v1alpha1" --clientset-path kope.io/auth/pkg/client/clientset_generated/
