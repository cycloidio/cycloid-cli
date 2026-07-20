package common_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/gen/models"
	"github.com/cycloidio/cycloid-cli/utils/ptr"
)

func TestFormUseCaseToFormVars(t *testing.T) {
	t.Run("GroupNameEqualsSectionName", func(t *testing.T) {
		// A group named the same as its own section used to bypass the nested
		// map init guard (CLI-145), causing a nil map panic on write.
		stackConfig := models.ServiceCatalogConfigs{
			"default": models.ServiceCatalogConfig{
				Forms: &models.FormUseCase{
					Sections: []*models.FormSection{
						{
							Name: ptr.Ptr("shared"),
							Groups: []*models.FormGroup{
								{
									Name: ptr.Ptr("shared"),
									Vars: []*models.FormEntity{
										{Key: ptr.Ptr("key1"), Default: "value1"},
									},
								},
							},
						},
					},
				},
			},
		}

		var res models.FormVariables
		var err error
		assert.NotPanics(t, func() {
			res, err = common.FormUseCaseToFormVars(stackConfig, "default")
		})
		assert.NoError(t, err)

		assert.Equal(t, models.FormVariables{
			"shared": {
				"shared": {
					"key1": "value1",
				},
			},
		}, res)
	})

	t.Run("GroupNameCollidesWithAnotherSectionName", func(t *testing.T) {
		// The outer output map only ever holds section names as top-level keys,
		// so the wrong guard (output[*group.Name]) falsely reports "already
		// initialized" whenever a group's name matches ANY already-processed
		// section — not just its own. Reproduce with sectionB's group named
		// after sectionA, which was already allocated by the time sectionB
		// is processed.
		stackConfig := models.ServiceCatalogConfigs{
			"default": models.ServiceCatalogConfig{
				Forms: &models.FormUseCase{
					Sections: []*models.FormSection{
						{
							Name: ptr.Ptr("sectionA"),
							Groups: []*models.FormGroup{
								{
									Name: ptr.Ptr("g1"),
									Vars: []*models.FormEntity{
										{Key: ptr.Ptr("key1"), Default: "valueA"},
									},
								},
							},
						},
						{
							Name: ptr.Ptr("sectionB"),
							Groups: []*models.FormGroup{
								{
									Name: ptr.Ptr("sectionA"),
									Vars: []*models.FormEntity{
										{Key: ptr.Ptr("key1"), Default: "valueB"},
									},
								},
							},
						},
					},
				},
			},
		}

		var res models.FormVariables
		var err error
		assert.NotPanics(t, func() {
			res, err = common.FormUseCaseToFormVars(stackConfig, "default")
		})
		assert.NoError(t, err)

		assert.Equal(t, models.FormVariables{
			"sectionA": {
				"g1": {
					"key1": "valueA",
				},
			},
			"sectionB": {
				"sectionA": {
					"key1": "valueB",
				},
			},
		}, res)
	})
}
