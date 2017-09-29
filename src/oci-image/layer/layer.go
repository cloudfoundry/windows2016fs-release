package layer

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/Microsoft/go-winio/archive/tar"

	winio "github.com/Microsoft/go-winio"
	"github.com/Microsoft/go-winio/backuptar"
	"github.com/Microsoft/hcsshim"
)

type Manager struct {
	driverInfo hcsshim.DriverInfo
}

func NewManager(driverInfo hcsshim.DriverInfo) *Manager {
	return &Manager{
		driverInfo: driverInfo,
	}
}

func (m *Manager) Extract(layerGzipFile, layerId string, parentLayerPaths []string) error {
	if err := winio.EnableProcessPrivileges([]string{winio.SeBackupPrivilege, winio.SeRestorePrivilege}); err != nil {
		return err
	}
	defer winio.DisableProcessPrivileges([]string{winio.SeBackupPrivilege, winio.SeRestorePrivilege})

	layerWriter, err := hcsshim.NewLayerWriter(m.driverInfo, layerId, parentLayerPaths)
	if err != nil {
		return fmt.Errorf("Failed to create new layer writer: %s", err.Error())
	}

	gf, err := os.Open(layerGzipFile)
	if err != nil {
		return err
	}
	defer gf.Close()

	gr, err := gzip.NewReader(gf)
	if err != nil {
		return err
	}
	defer gr.Close()

	tr := tar.NewReader(gr)

	hdr, err := tr.Next()
	buf := bufio.NewWriter(nil)

	for err == nil {
		base := path.Base(hdr.Name)
		if strings.HasPrefix(base, ".wh.") {
			name := path.Join(path.Dir(hdr.Name), base[len(".wh."):])
			err = layerWriter.Remove(filepath.FromSlash(name))
			if err != nil {
				return fmt.Errorf("Failed to remove: %s", err.Error())
			}
			hdr, err = tr.Next()
		} else if hdr.Typeflag == tar.TypeLink {
			err = layerWriter.AddLink(filepath.FromSlash(hdr.Name), filepath.FromSlash(hdr.Linkname))
			if err != nil {
				return fmt.Errorf("Failed to add link: %s", err.Error())
			}
			hdr, err = tr.Next()
		} else {
			var (
				name     string
				fileInfo *winio.FileBasicInfo
			)
			name, _, fileInfo, err = backuptar.FileInfoFromHeader(hdr)
			if err != nil {
				return fmt.Errorf("Failed to get file info: %s", err.Error())
			}
			err = layerWriter.Add(filepath.FromSlash(name), fileInfo)
			if err != nil {
				return fmt.Errorf("Failed to add layer: %s", err.Error())
			}
			buf.Reset(layerWriter)

			hdr, err = backuptar.WriteBackupStreamFromTarFile(buf, tr, hdr)
			ferr := buf.Flush()
			if ferr != nil {
				err = ferr
			}
		}
	}

	if err != io.EOF {
		return err
	}

	return layerWriter.Close()
}
