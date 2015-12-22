// Copyright 2015 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package cloud

// Generated code - do not edit.

const fallbackPublicCloudInfo = `
# DO NOT EDIT, will be overwritten, use “juju update-clouds” to refresh.
clouds:
  aws:
    type: aws
    auth-types: [ access-key ]
    regions:
      us-east-1:
        endpoint: https://us-east-1.aws.amazon.com/v1.2/
      us-west-1:
        endpoint: https://us-west-1.aws.amazon.com/v1.2/
      us-west-2:
        endpoint: https://us-west-2.aws.amazon.com/v1.2/
      eu-west-1:
        endpoint: https://eu-west-1.aws.amazon.com/v1.2/
      eu-central-1:
        endpoint: https://eu-central-1.aws.amazon.com/v1.2/
      ap-southeast-1:
        endpoint: https://ap-southeast-1.aws.amazon.com/v1.2/
      ap-southeast-2:
        endpoint: https://ap-southeast-2.aws.amazon.com/v1.2/
      ap-northeast-1:
        endpoint: https://ap-northeast-1.aws.amazon.com/v1.2/
      sa-east-1:
        endpoint: https://sa-east-1.aws.amazon.com/v1.2/
  aws-china:
    type: aws
    auth-types: [ access-key ]
    regions:
      cn-north-1:
        endpoint: https://ec2.cn-north-1.amazonaws.com.cn/
  aws-gov:
    type: aws
    auth-types: [ access-key ]
    regions:
      us-gov-west-1:
        endpoint: https://ec2.us-gov-west-1.amazonaws-govcloud.com
  google:
    type: gce
    auth-types: [ oauth2 ]
    endpoint: https://www.googleapis.com
    regions:
      us-east1:
      us-central1:
      europe-west1:
      asia-east1:
  azure:
    type: azure
    auth-types: [ userpass ]
    endpoint: https://management.core.windows.net/
    regions:
      Central US:
      East US:
      East US 2:
      North Central US:
      South Central US:
      West US:
      North Europe:
      West Europe:
      East Asia:
      Southeast Asia:
      Japan East:
      Japan West:
      Brazil South:
      Australia East:
      Australia Southeast:
      Central India:
      South India:
      West India:
  azure-china:
    type: azure
    auth-types: [ userpass ]
    endpoint: https://management.core.chinacloudapi.cn/
    regions:
      China East:
      China North:
  rackspace:
    type: openstack
    auth-types: [ access-key, userpass ]
    endpoint: https://identity.api.rackspacecloud.com/v2.0
    regions:
      Dallas-Fort Worth:
      Chicago:
      Northern Virginia:
      London:
        endpoint: https://lon.identity.api.rackspacecloud.com/v2.0
      Sydney:
      Hong Kong:
`
