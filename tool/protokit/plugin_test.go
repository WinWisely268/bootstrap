package protokit_test

import (
	"github.com/stretchr/testify/suite"
	"google.golang.org/protobuf/proto"
	plugin_go "google.golang.org/protobuf/types/pluginpb"

	"bytes"
	"errors"
	"testing"

	"github.com/getcouragenow/bootstrap/tool/protokit"
	"github.com/getcouragenow/bootstrap/tool/protokit/utils"
)

type PluginTest struct {
	suite.Suite
}

func TestPlugin(t *testing.T) {
	suite.Run(t, new(PluginTest))
}

func (assert *PluginTest) TestRunPlugin() {
	fds, err := utils.LoadDescriptorSet("fixtures", "fileset.pb")
	assert.NoError(err)

	req := utils.CreateGenRequest(fds, "booking.proto", "todo.proto")
	data, err := proto.Marshal(req)
	assert.NoError(err)

	in := bytes.NewBuffer(data)
	out := new(bytes.Buffer)

	assert.NoError(protokit.RunPluginWithIO(new(OkPlugin), in, out))
	assert.NotEmpty(out)
}

func (assert *PluginTest) TestRunPluginInputError() {
	in := bytes.NewBufferString("Not a codegen request")
	out := new(bytes.Buffer)

	err := protokit.RunPluginWithIO(nil, in, out)
	assert.EqualError(err, "proto: cannot parse reserved wire type")
	assert.Empty(out)
}

func (assert *PluginTest) TestRunPluginNoFilesToGenerate() {
	fds, err := utils.LoadDescriptorSet("fixtures", "fileset.pb")
	assert.NoError(err)

	req := utils.CreateGenRequest(fds)
	data, err := proto.Marshal(req)
	assert.NoError(err)

	in := bytes.NewBuffer(data)
	out := new(bytes.Buffer)

	err = protokit.RunPluginWithIO(new(ErrorPlugin), in, out)
	assert.EqualError(err, "no files were supplied to the generator")
	assert.Empty(out)
}

func (assert *PluginTest) TestRunPluginGeneratorError() {
	fds, err := utils.LoadDescriptorSet("fixtures", "fileset.pb")
	assert.NoError(err)

	req := utils.CreateGenRequest(fds, "booking.proto", "todo.proto")
	data, err := proto.Marshal(req)
	assert.NoError(err)

	in := bytes.NewBuffer(data)
	out := new(bytes.Buffer)

	err = protokit.RunPluginWithIO(new(ErrorPlugin), in, out)
	assert.EqualError(err, "generator error")
	assert.Empty(out)
}

type ErrorPlugin struct{}

func (ep *ErrorPlugin) Generate(r *plugin_go.CodeGeneratorRequest) (*plugin_go.CodeGeneratorResponse, error) {
	return nil, errors.New("generator error")
}

type OkPlugin struct{}

func (op *OkPlugin) Generate(r *plugin_go.CodeGeneratorRequest) (*plugin_go.CodeGeneratorResponse, error) {
	resp := new(plugin_go.CodeGeneratorResponse)
	resp.File = append(resp.File, &plugin_go.CodeGeneratorResponse_File{
		Name:    proto.String("myfile.out"),
		Content: proto.String("someoutput"),
	})

	return resp, nil
}
