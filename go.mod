module github.com/GoogleCloudPlatform/terraformer

go 1.16

require (
	cloud.google.com/go v0.77.0
	cloud.google.com/go/logging v1.1.2
	cloud.google.com/go/storage v1.14.0
	github.com/Azure/azure-sdk-for-go v43.0.0+incompatible
	github.com/Azure/azure-storage-blob-go v0.10.0
	github.com/Azure/go-autorest/autorest v0.11.12
	github.com/DataDog/datadog-api-client-go v1.0.0-beta.20
	github.com/IBM-Cloud/bluemix-go v0.0.0-20210203095940-db28d5e07b55
	github.com/IBM/go-sdk-core/v3 v3.3.1
	github.com/IBM/go-sdk-core/v4 v4.9.0
	github.com/IBM/ibm-cos-sdk-go v1.5.0
	github.com/IBM/keyprotect-go-client v0.6.0
	github.com/IBM/networking-go-sdk v0.13.0
	github.com/IBM/vpc-go-sdk v0.4.1
	github.com/Nvveen/Gotty v0.0.0-20120604004816-cd527374f1e5 // indirect
	github.com/OctopusDeploy/go-octopusdeploy v1.6.0
	github.com/PaloAltoNetworks/pango v0.5.2-0.20210514062125-b809b85eb51c
	github.com/SAP/go-hdb v0.105.2 // indirect
	github.com/SermoDigital/jose v0.9.1 // indirect
	github.com/aliyun/alibaba-cloud-sdk-go v1.60.295
	github.com/aliyun/aliyun-tablestore-go-sdk v4.1.2+incompatible
	github.com/apache/openwhisk-client-go v0.0.0-20210106144548-17d556327cd3
	github.com/aws/aws-sdk-go-v2 v1.6.0
	github.com/aws/aws-sdk-go-v2/config v1.3.0
	github.com/aws/aws-sdk-go-v2/credentials v1.2.1
	github.com/aws/aws-sdk-go-v2/service/accessanalyzer v1.2.0
	github.com/aws/aws-sdk-go-v2/service/acm v1.2.1
	github.com/aws/aws-sdk-go-v2/service/apigateway v1.2.1
	github.com/aws/aws-sdk-go-v2/service/appsync v1.2.1
	github.com/aws/aws-sdk-go-v2/service/autoscaling v1.3.1
	github.com/aws/aws-sdk-go-v2/service/batch v1.3.1
	github.com/aws/aws-sdk-go-v2/service/budgets v1.1.3
	github.com/aws/aws-sdk-go-v2/service/cloud9 v1.1.3
	github.com/aws/aws-sdk-go-v2/service/cloudformation v1.3.0
	github.com/aws/aws-sdk-go-v2/service/cloudfront v1.3.0
	github.com/aws/aws-sdk-go-v2/service/cloudhsmv2 v1.1.3
	github.com/aws/aws-sdk-go-v2/service/cloudtrail v1.2.1
	github.com/aws/aws-sdk-go-v2/service/cloudwatch v1.3.0
	github.com/aws/aws-sdk-go-v2/service/cloudwatchevents v1.2.0
	github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs v1.2.3
	github.com/aws/aws-sdk-go-v2/service/codebuild v1.2.1
	github.com/aws/aws-sdk-go-v2/service/codecommit v1.1.3
	github.com/aws/aws-sdk-go-v2/service/codedeploy v1.2.1
	github.com/aws/aws-sdk-go-v2/service/codepipeline v1.2.1
	github.com/aws/aws-sdk-go-v2/service/cognitoidentity v1.2.1
	github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider v1.2.1
	github.com/aws/aws-sdk-go-v2/service/configservice v1.3.0
	github.com/aws/aws-sdk-go-v2/service/datapipeline v1.1.3
	github.com/aws/aws-sdk-go-v2/service/devicefarm v1.1.3
	github.com/aws/aws-sdk-go-v2/service/dynamodb v1.2.1
	github.com/aws/aws-sdk-go-v2/service/ec2 v1.3.0
	github.com/aws/aws-sdk-go-v2/service/ecr v1.2.1
	github.com/aws/aws-sdk-go-v2/service/ecrpublic v1.2.0
	github.com/aws/aws-sdk-go-v2/service/ecs v1.2.1
	github.com/aws/aws-sdk-go-v2/service/efs v1.2.1
	github.com/aws/aws-sdk-go-v2/service/eks v1.2.1
	github.com/aws/aws-sdk-go-v2/service/elasticache v1.2.1
	github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk v1.2.1
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing v1.2.1
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 v1.2.1
	github.com/aws/aws-sdk-go-v2/service/elasticsearchservice v1.2.1
	github.com/aws/aws-sdk-go-v2/service/emr v1.2.1
	github.com/aws/aws-sdk-go-v2/service/firehose v1.2.1
	github.com/aws/aws-sdk-go-v2/service/glue v1.3.0
	github.com/aws/aws-sdk-go-v2/service/iam v1.3.0
	github.com/aws/aws-sdk-go-v2/service/iot v1.2.0
	github.com/aws/aws-sdk-go-v2/service/kafka v1.2.1
	github.com/aws/aws-sdk-go-v2/service/kinesis v1.2.1
	github.com/aws/aws-sdk-go-v2/service/kms v1.2.2
	github.com/aws/aws-sdk-go-v2/service/lambda v1.2.1
	github.com/aws/aws-sdk-go-v2/service/mediapackage v1.2.1
	github.com/aws/aws-sdk-go-v2/service/mediastore v1.1.4
	github.com/aws/aws-sdk-go-v2/service/opsworks v1.2.2
	github.com/aws/aws-sdk-go-v2/service/organizations v1.2.1
	github.com/aws/aws-sdk-go-v2/service/qldb v1.1.3
	github.com/aws/aws-sdk-go-v2/service/rds v1.2.1
	github.com/aws/aws-sdk-go-v2/service/resourcegroups v1.2.1
	github.com/aws/aws-sdk-go-v2/service/route53 v1.3.0
	github.com/aws/aws-sdk-go-v2/service/s3 v1.9.0
	github.com/aws/aws-sdk-go-v2/service/secretsmanager v1.2.1
	github.com/aws/aws-sdk-go-v2/service/securityhub v1.2.1
	github.com/aws/aws-sdk-go-v2/service/servicecatalog v1.2.1
	github.com/aws/aws-sdk-go-v2/service/ses v1.2.1
	github.com/aws/aws-sdk-go-v2/service/sfn v1.2.1
	github.com/aws/aws-sdk-go-v2/service/sns v1.2.1
	github.com/aws/aws-sdk-go-v2/service/sqs v1.3.0
	github.com/aws/aws-sdk-go-v2/service/ssm v1.3.0
	github.com/aws/aws-sdk-go-v2/service/sts v1.4.1
	github.com/aws/aws-sdk-go-v2/service/swf v1.2.2
	github.com/aws/aws-sdk-go-v2/service/waf v1.1.4
	github.com/aws/aws-sdk-go-v2/service/wafregional v1.2.1
	github.com/aws/aws-sdk-go-v2/service/workspaces v1.2.1
	github.com/aws/aws-sdk-go-v2/service/xray v1.2.1
	github.com/cenkalti/backoff v2.2.1+incompatible // indirect
	github.com/cloudflare/cloudflare-go v0.13.6
	github.com/cloudfoundry/jibber_jabber v0.0.0-20151120183258-bcc4c8345a21 // indirect
	github.com/ddelnano/terraform-provider-mikrotik/client v0.0.0-20210401060029-7f652169b2c4
	github.com/ddelnano/terraform-provider-xenorchestra/client v0.0.0-20210401070256-0d721c6762ef
	github.com/denverdino/aliyungo v0.0.0-20200327235253-d59c209c7e93
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/digitalocean/godo v1.57.0
	github.com/dollarshaveclub/new-relic-synthetics-go v0.0.0-20170605224734-4dc3dd6ae884
	github.com/duosecurity/duo_api_golang v0.0.0-20201112143038-0e07e9f869e3 // indirect
	github.com/elazarl/go-bindata-assetfs v1.0.1 // indirect
	github.com/fastly/go-fastly/v3 v3.6.0
	github.com/fatih/structs v1.1.0 // indirect
	github.com/gocql/gocql v0.0.0-20210707082121-9a3953d1826d // indirect
	github.com/google/go-github/v35 v35.1.0
	github.com/gophercloud/gophercloud v0.17.0
	github.com/grafana/grafana-api-golang-client v0.0.0-20210218192924-9ccd2365d2a6

	github.com/harvester/terraform-provider-harvester v0.1.0
	github.com/hashicorp/go-azure-helpers v0.10.0
	github.com/hashicorp/go-cleanhttp v0.5.2
	github.com/hashicorp/go-hclog v0.15.0
	github.com/hashicorp/go-memdb v1.3.2 // indirect
	github.com/hashicorp/go-plugin v1.4.0
	github.com/hashicorp/hcl v1.0.0
	github.com/hashicorp/terraform v0.12.31
	github.com/hashicorp/vault v0.10.4
	github.com/heimweh/go-pagerduty v0.0.0-20210412205347-cc0e5d3c14d4
	github.com/heroku/heroku-go/v5 v5.1.0
	github.com/hokaccha/go-prettyjson v0.0.0-20210113012101-fb4e108d2519 // indirect
	github.com/iancoleman/strcase v0.1.2
	github.com/jefferai/jsonx v1.0.1 // indirect
	github.com/jmespath/go-jmespath v0.4.0
	github.com/jonboydell/logzio_client v1.2.0
	github.com/labd/commercetools-go-sdk v0.3.1
	github.com/linode/linodego v0.24.1
	github.com/mrparkers/terraform-provider-keycloak v0.0.0-20200506151941-509881368409
	github.com/nicksnyder/go-i18n v1.10.1 // indirect
	github.com/ns1/ns1-go v2.4.0+incompatible
	github.com/ory/dockertest v3.3.5+incompatible // indirect
	github.com/packethost/packngo v0.9.0
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	github.com/paultyng/go-newrelic/v4 v4.10.0
	github.com/pkg/errors v0.9.1
	github.com/ryanuber/go-glob v1.0.0 // indirect
	github.com/spf13/cobra v1.1.1
	github.com/spf13/pflag v1.0.5
	github.com/tencentcloud/tencentcloud-sdk-go v3.0.233+incompatible
	github.com/tencentyun/cos-go-sdk-v5 v0.7.19
	github.com/vultr/govultr v0.5.0
	github.com/yandex-cloud/go-genproto v0.0.0-20200722140432-762fe965ce77
	github.com/yandex-cloud/go-sdk v0.0.0-20200722140627-2194e5077f13
	github.com/zclconf/go-cty v1.8.2
	github.com/zorkian/go-datadog-api v2.30.0+incompatible
	golang.org/x/oauth2 v0.0.0-20210218202405-ba52d332ba99
	golang.org/x/text v0.3.6
	gonum.org/v1/gonum v0.7.0
	google.golang.org/api v0.40.0
	google.golang.org/genproto v0.0.0-20210226172003-ab064af71705
	gopkg.in/jarcoal/httpmock.v1 v1.0.0-00010101000000-000000000000 // indirect
	k8s.io/apimachinery v0.21.0
	k8s.io/client-go v12.0.0+incompatible
)

replace (
	github.com/Azure/azure-sdk-for-go => github.com/Azure/azure-sdk-for-go v42.3.0+incompatible
	// start harvester
	github.com/dgrijalva/jwt-go => github.com/dgrijalva/jwt-go v3.2.1-0.20200107013213-dc14462fd587+incompatible
	github.com/docker/distribution => github.com/docker/distribution v0.0.0-20191216044856-a8371794149d
	github.com/docker/docker => github.com/docker/docker v1.4.2-0.20200203170920-46ec8731fbce
	github.com/go-kit/kit => github.com/go-kit/kit v0.3.0
	github.com/googleapis/gnostic => github.com/googleapis/gnostic v0.4.1

	github.com/harvester/harvester => github.com/harvester/harvester v0.0.2-0.20210621053253-abc91f0304f5

	github.com/knative/pkg => github.com/rancher/pkg v0.0.0-20190514055449-b30ab9de040e
	github.com/openshift/api => github.com/openshift/api v0.0.0-20191219222812-2987a591a72c
	github.com/openshift/client-go => github.com/openshift/client-go v0.0.0-20200521150516-05eb9880269c
	github.com/operator-framework/operator-lifecycle-manager => github.com/operator-framework/operator-lifecycle-manager v0.0.0-20190128024246-5eb7ae5bdb7a
	github.com/rancher/rancher/pkg/apis => github.com/rancher/rancher/pkg/apis v0.0.0-20210304063736-65f7c844267b
	github.com/rancher/rancher/pkg/client => github.com/rancher/rancher/pkg/client v0.0.0-20210304063736-65f7c844267b
	gopkg.in/jarcoal/httpmock.v1 => github.com/jarcoal/httpmock v1.0.5

	k8s.io/api => k8s.io/api v0.20.4
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.20.4
	k8s.io/apimachinery => k8s.io/apimachinery v0.20.4
	k8s.io/apiserver => k8s.io/apiserver v0.20.4
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.20.4
	k8s.io/client-go => k8s.io/client-go v0.20.4
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.20.4
	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.20.4
	k8s.io/code-generator => k8s.io/code-generator v0.20.4
	k8s.io/component-base => k8s.io/component-base v0.20.4
	k8s.io/component-helpers => k8s.io/component-helpers v0.20.4
	k8s.io/controller-manager => k8s.io/controller-manager v0.20.4
	k8s.io/cri-api => k8s.io/cri-api v0.20.4
	k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.20.4
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.20.4
	k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.20.4
	k8s.io/kube-proxy => k8s.io/kube-proxy v0.20.4
	k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.20.4
	k8s.io/kubectl => k8s.io/kubectl v0.20.4
	k8s.io/kubelet => k8s.io/kubelet v0.20.4
	k8s.io/kubernetes => k8s.io/kubernetes v1.20.2
	k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.20.4
	k8s.io/metrics => k8s.io/metrics v0.20.4
	k8s.io/mount-utils => k8s.io/mount-utils v0.20.4
	k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.20.4

	kubevirt.io/client-go => github.com/kubevirt/client-go v0.40.0
	kubevirt.io/containerized-data-importer => github.com/rancher/kubevirt-containerized-data-importer v1.26.1-0.20210303063201-9e7a78643487
	sigs.k8s.io/structured-merge-diff => sigs.k8s.io/structured-merge-diff v0.0.0-20190302045857-e85c7b244fd2
// end harvester
)
