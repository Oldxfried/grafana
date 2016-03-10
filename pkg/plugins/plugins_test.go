package plugins

import (
	"path/filepath"
	"testing"

	"github.com/grafana/grafana/pkg/setting"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/ini.v1"
)

func TestPluginScans(t *testing.T) {

	Convey("When scaning for plugins", t, func() {
		setting.StaticRootPath, _ = filepath.Abs("../../public/")
		setting.Cfg = ini.Empty()
		err := Init()

		So(err, ShouldBeNil)
		So(len(DataSources), ShouldBeGreaterThan, 1)
		So(len(Panels), ShouldBeGreaterThan, 1)

		Convey("Should set module automatically", func() {
			So(DataSources["graphite"].Module, ShouldEqual, "app/plugins/datasource/graphite/module")
		})
	})

	Convey("When reading app plugin definition", t, func() {
		setting.Cfg = ini.Empty()
		sec, _ := setting.Cfg.NewSection("plugin.nginx-app")
		sec.NewKey("path", "../../examples/nginx-app")
		err := Init()

		So(err, ShouldBeNil)
		So(len(Apps), ShouldBeGreaterThan, 0)

		So(Apps["nginx-app"].Info.Logos.Large, ShouldEqual, "public/plugins/nginx-app/img/logo_large.png")
		So(Apps["nginx-app"].Info.Screenshots[1].Path, ShouldEqual, "public/plugins/nginx-app/img/screenshot2.png")
	})

}
