package web

import (
	"log"

	"github.com/Azure/azure-sdk-for-go/services/web/mgmt/2019-08-01/web"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/utils"
)

func appServiceSiteSourceControlSchema() *schema.Schema {
	return &schema.Schema{
		Type:          schema.TypeList,
		Optional:      true,
		MaxItems:      1,
		Computed:      true,
		ConflictsWith: []string{"site_config.0.scm_type"},
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"repo_url": {
					Type:     schema.TypeString,
					Optional: true,
					Computed: true,
				},

				"branch": {
					Type:     schema.TypeString,
					Optional: true,
					Computed: true,
				},

				"manual_integration": {
					Type:     schema.TypeBool,
					Optional: true,
					Computed: true,
				},

				"use_mercurial": {
					Type:     schema.TypeBool,
					Optional: true,
					Computed: true,
				},

				"rollback_enabled": {
					Type:     schema.TypeBool,
					Optional: true,
					Computed: true,
				},
			},
		},
	}
}

func expandAppServiceSiteSourceControl(d *schema.ResourceData) *web.SiteSourceControlProperties {
	sourceControlRaw := d.Get("source_control").([]interface{})
	sourceControl := sourceControlRaw[0].(map[string]interface{})

	result := &web.SiteSourceControlProperties{
		RepoURL:                   utils.String(sourceControl["repo_url"].(string)),
		Branch:                    utils.String(sourceControl["branch"].(string)),
		IsManualIntegration:       utils.Bool(sourceControl["manual_integration"].(bool)),
		IsMercurial:               utils.Bool(sourceControl["use_mercurial"].(bool)),
		DeploymentRollbackEnabled: utils.Bool(sourceControl["rollback_enabled"].(bool)),
	}

	return result
}

func flattenAppServiceSourceControl(input *web.SiteSourceControlProperties) []interface{} {
	results := make([]interface{}, 0)
	result := make(map[string]interface{})

	if input == nil {
		log.Printf("[DEBUG] SiteSourceControlProperties is nil")
		return results
	}

	if input.RepoURL != nil && *input.RepoURL != "" {
		result["repo_url"] = *input.RepoURL
	}

	if input.Branch != nil && *input.Branch != "" {
		result["branch"] = *input.Branch
	} else {
		result["branch"] = "master"
	}

	result["use_mercurial"] = *input.IsMercurial

	result["manual_integration"] = *input.IsManualIntegration

	result["rollback_enabled"] = *input.DeploymentRollbackEnabled

	return append(results, result)
}