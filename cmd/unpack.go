package cmd

import (
	"archiver/lib/compression"
	"archiver/lib/compression/vlc"
	"github.com/spf13/cobra"
	"io"
	"os"
	"path/filepath"
	"strings"
)

const unpackedExtension = "txt"

func unpack(cmd *cobra.Command, args []string) {
	var decoder compression.Decoder
	if len(args) == 0 || args[0] == "" {
		handleErr(ErrEmptyPath)
	}
	method := cmd.Flag("method").Value.String()
	switch method {
	case "vlc":
		decoder = vlc.New()
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

	// packed := Encode(data)
	packed := decoder.Decode(data)
	err = os.WriteFile(unpackedFileName(filePath), []byte(packed), 0644)
	if err != nil {
		handleErr(err)
	}
}

func unpackedFileName(path string) string { // TODO: Убрать лишние переменные
	// path/to/file/myFile.txt -> myFile.vlc
	fileName := filepath.Base(path)
	return strings.TrimSuffix(fileName, filepath.Ext(fileName)) + "." + unpackedExtension
}

var unpackCmd = &cobra.Command{
	Use:   "unpack",
	Short: "Unpack file",
	Run:   unpack,
}

func init() {
	// Добавляем к корневой команде команду "pack"
	rootCmd.AddCommand(unpackCmd)
	unpackCmd.Flags().StringP("method", "m", "", "decompression method: vlc")
	if err := unpackCmd.MarkFlagRequired("method"); err != nil {
		panic(err)
	}
}
