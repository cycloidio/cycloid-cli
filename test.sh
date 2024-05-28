#!/usr/bin/env bash

org=cycloid-sandbox


echo -e '
{
  "Cloud provider": {
  	"Terraform": {
  		"gcp_credentials_json": "((gcp.json_key))",
  		"gcp_project": "($ .organization_canonical $)",
  		"gcp_region": "europe-west1",
  		"gcp_zone": "europe-west1-b",
  		"terraform_storage_bucket_name": "($ .organization_canonical $)-terraform-remote-state"
  	}
  },
  "Instance": {
  	"Boot-disk": {
  		"module.vm.boot_disk_auto_delete": true,
  		"module.vm.boot_disk_device_name": "",
  		"module.vm.boot_disk_image": "debian-cloud/debian-12",
  		"module.vm.boot_disk_size": 5,
  		"module.vm.boot_disk_type": "pd-standard"
  	},
  	"Details": {
  		"module.vm.allow_stopping_for_update": true,
  		"module.vm.file_content": "",
  		"module.vm.instance_name": "${var.customer}-${var.project}-${var.env}-vm",
  		"module.vm.instance_tags": [
  			"${var.customer}-${var.project}-${var.env}-network-tag"
  		],
  		"module.vm.machine_type": "e2-small"
  	},
  	"Firewall-egress": {
  		"module.vm.egress_allow_protocol": "",
  		"module.vm.egress_disabled": true,
  		"module.vm.egress_firewall_name": ""
  	},
  	"Firewall-ingress": {
  		"module.vm.ingress_allow_ports": [
  			"22"
  		],
  		"module.vm.ingress_allow_protocol": "tcp",
  		"module.vm.ingress_disabled": false,
  		"module.vm.ingress_firewall_name": "${var.customer}-${var.project}-${var.env}-ingress"
  	},
  	"Network": {
  		"module.vm.network": "qscqsc",
  		"module.vm.network_ip": ""
  	}
  }
}
' | go run . projects create-stackforms-env \
  --verbosity debug \
  --org $org \
  --project "test-env" \
  --env "test2" \
  --use-case "gcp-gce" \
  --var-file '-'


# export CY_CREATE_ENV_VARS='{"viaEnv": "ok", "file4": "overriden", "subMap": {"two": 2}}'
# echo -e '{\n"file1": "titi",\n"toto": "toto"\n}{"file4": 3, "subMap": {"one": 1}}' | go run . projects create-stackforms-env \
#   --verbosity debug \
#   --org $org \
#   --project "test-env" \
#   --env "test" \
#   --vars '{"toto": "vars1"}' \
#   --var-file '-' \
#   --var-file <(echo '{"file2": "tata"}')
