package help_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	screenmock "meegle-cli/internal/screen/mock"
	"meegle-cli/internal/screens/help"
)

func TestModelFocusTogglesVisibility(t *testing.T) {
	ctrl := gomock.NewController(t)
	app := screenmock.NewMockAppModel(ctrl)

	model := help.New()
	model.OnFocus(app)
	view := model.View(app)
	assert.True(t, strings.Contains(view, "Help"))

	model.OnBlur(app)
	assert.Equal(t, "", model.View(app))
}
