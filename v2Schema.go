/*
Copyright © 2022 Rich Rose <richardrose@google.com>

*/
package v2

type InstructionType struct {
	Type string `yaml:"type"`
	URI  string `yaml:"uri"`
}

type CustomPropertiesType struct {
	Key       string `yaml:"key"`
	Value     string `yaml:"value,omitempty"`
	Reference string `yaml:"reference,omitempty"`
}

type StartupScriptType struct {
	Type             string `yaml:"type,omitempty"`
	Path             string `yaml:"path,omitempty"`
	CustomProperties []CustomPropertiesType
	//	CustomProperties []struct {
	//		Key       string `yaml:"key"`
	//		Value     string `yaml:"value,omitempty"`
	//		Reference string `yaml:"reference,omitempty"`
	//	} `yaml:"custom_properties"`
}

type StudentVisibleOutputsType struct {
	Label     string `yaml:"label"`
	Reference string `yaml:"reference"`
}

type EnvironmentResourceType struct {
	Type          string `yaml:"type"`
	ID            string `yaml:"id"`
	Variant       string `yaml:"variant,omitempty"`
	SshKeyUser    string `yaml:"ssh_key_user,omitempty"`
	StartupScript StartupScriptType
	//			StartupScript struct {
	//				Type             string `yaml:"type"`
	//				Path             string `yaml:"path"`
	//				CustomProperties []struct {
	//					Key       string `yaml:"key"`
	//					Value     string `yaml:"value,omitempty"`
	//					Reference string `yaml:"reference,omitempty"`
	//				} `yaml:"custom_properties"`
	//			} `yaml:"startup_script,omitempty"`
	CleanupScript struct {
		Type             string `yaml:"type,omitempty"`
		Path             string `yaml:"path,omitempty"`
		CustomProperties []struct {
			Key       string `yaml:"key,omitempty"`
			Reference string `yaml:"reference,omitempty"`
		} `yaml:"custom_properties"`
	} `yaml:"cleanup_script,omitempty"`
	AllowedLocations []string `yaml:"allowed_locations,omitempty"`
	Permissions      []struct {
		Project string   `yaml:"project"`
		Roles   []string `yaml:"roles"`
	} `yaml:"permissions,omitempty"`
	UserPolicy          string `yaml:"user_policy,omitempty"`
	AccountRestrictions struct {
		AllowSpotInstances  bool     `yaml:"allow_spot_instances"`
		AllowSubnetDeletion bool     `yaml:"allow_subnet_deletion"`
		AllowVpcDeletion    bool     `yaml:"allow_vpc_deletion"`
		AllowedRdsInstances []string `yaml:"allowed_rds_instances"`
	} `yaml:"account_restrictions,omitempty"`
}

type EnvironmentType struct {
	Resources             []EnvironmentResourceType
	StudentVisibleOutputs []StudentVisibleOutputsType
}

// NOTE: Schema is autogenerated
// https://mholt.github.io/json-to-go/
// https://zhwt.github.io/yaml-to-go/
type SchemaV2 struct {
	EntityType    string   `yaml:"entity_type"`
	SchemaVersion int      `yaml:"schema_version"`
	Title         string   `yaml:"title"`
	Description   string   `yaml:"description"`
	DefaultLocale string   `yaml:"default_locale"`
	Duration      int      `yaml:"duration"`
	MaxDuration   int      `yaml:"max_duration"`
	Credits       int      `yaml:"credits"`
	Level         string   `yaml:"level"`
	Tags          []string `yaml:"tags"`
	Instruction   InstructionType
	// 	Instruction   struct {
	// 		Type string `yaml:"type"`
	// 		URI  string `yaml:"uri"`
	// 	} `yaml:"instruction"`
	Resources []struct {
		Type        string `yaml:"type"`
		ID          string `yaml:"id"`
		Title       string `yaml:"title"`
		Description string `yaml:"description"`
		URI         string `yaml:"uri"`
	} `yaml:"resources"`
	Environment EnvironmentType
	//	Environment struct {
	//    Resources   []ResourceType
	//		Resources []struct {
	//			Type          string `yaml:"type"`
	//			ID            string `yaml:"id"`
	//			Variant       string `yaml:"variant,omitempty"`
	//SshKeyUser    string `yaml:"ssh_key_user,omitempty"`
	//			StartupScript struct {
	//				Type             string `yaml:"type"`
	//				Path             string `yaml:"path"`
	//				CustomProperties []struct {
	//					Key       string `yaml:"key"`
	//					Value     string `yaml:"value,omitempty"`
	//					Reference string `yaml:"reference,omitempty"`
	//				} `yaml:"custom_properties"`
	//			} `yaml:"startup_script,omitempty"`
	//			CleanupScript struct {
	//				Type             string `yaml:"type"`
	//				Path             string `yaml:"path"`
	//				CustomProperties []struct {
	//					Key       string `yaml:"key"`
	//					Reference string `yaml:"reference"`
	//				} `yaml:"custom_properties"`
	//			} `yaml:"cleanup_script,omitempty"`
	//			AllowedLocations []string `yaml:"allowed_locations,omitempty"`
	//			Permissions      []struct {
	//				Project string   `yaml:"project"`
	//				Roles   []string `yaml:"roles"`
	//			} `yaml:"permissions,omitempty"`
	//			UserPolicy          string `yaml:"user_policy,omitempty"`
	//			AccountRestrictions struct {
	//				AllowSpotInstances  bool     `yaml:"allow_spot_instances"`
	//				AllowSubnetDeletion bool     `yaml:"allow_subnet_deletion"`
	//				AllowVpcDeletion    bool     `yaml:"allow_vpc_deletion"`
	//				AllowedRdsInstances []string `yaml:"allowed_rds_instances"`
	//			} `yaml:"account_restrictions,omitempty"`
	//		} `yaml:"resources"`
	//		StudentVisibleOutputs []struct {
	//			Label     string `yaml:"label"`
	//			Reference string `yaml:"reference"`
	//		} `yaml:"student_visible_outputs"`
	//	} `yaml:"environment"`
	Assessment string `yaml:"assessment"`

	// Amend example to use the simple form
	//	Assessment struct {
	//		PassingPercentage int `yaml:"passing_percentage"`
	//		Steps             []struct {
	//			Title           string `yaml:"title"`
	//			LocaleID        string `yaml:"locale_id"`
	//			MaximumScore    int    `yaml:"maximum_score"`
	//			StudentMessages struct {
	//				Success             string `yaml:"success"`
	//				BucketMissing       string `yaml:"bucket_missing"`
	//				BucketMisconfigured string `yaml:"bucket_misconfigured"`
	//			} `yaml:"student_messages"`
	//			Services   []string `yaml:"services"`
	//			MethodName string   `yaml:"method_name"`
	//		} `yaml:"steps"`
	//	} `yaml:"assessment"`
}
