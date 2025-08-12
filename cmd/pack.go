package cmd

import (
	"archiver/lib/compression"
	"archiver/lib/compression/vlc"
	"errors"
	"github.com/spf13/cobra"
	"io"
	"os"
	"path/filepath"
	"strings"
)

const packedExtension = "vlc"

var ErrEmptyPath = errors.New("please specify a file path")

func pack(cmd *cobra.Command, args []string) {
	var encoder compression.Encoder
	if len(args) == 0 || args[0] == "" {
		handleErr(ErrEmptyPath)
	}

	method := cmd.Flag("method").Value.String()

	switch method {
	case "vlc":
		encoder = vlc.New()
	default:
		cmd.PrintErr("unknown method")
	}

	filePath := args[0]
	r, err := os.Open(filePath)
	if err != nil {
		handleErr(err)
	}
	defer r.Close()
	data, err := io.ReadAll(r)
	if err != nil {
		handleErr(err)
	}

	packed := encoder.Encode(string(data))
	err = os.WriteFile(packedFileName(filePath), packed, 0644)
	if err != nil {
		handleErr(err)
	}
}

func packedFileName(path string) string { // TODO: Убрать лишние переменные
	// path/to/file/myFile.txt -> myFile.vlc
	fileName := filepath.Base(path)
	return strings.TrimSuffix(fileName, filepath.Ext(fileName)) + "." + packedExtension
}

var packCmd = &cobra.Command{
	Use:   "pack",
	Short: "Pack file",
	Run:   pack,
}

func init() {
	// Добавляем к корневой команде команду "pack"
	rootCmd.AddCommand(packCmd)
	packCmd.Flags().StringP("method", "m", "", "compression method: vlc")
	if err := packCmd.MarkFlagRequired("method"); err != nil {
		panic(err)
	}
}
