// Author: yangzq80@gmail.com
// Date: 2023/7/18
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	bio "github.com/yusys-cloud/ai-tools/base/io"
	"net/http"
)

var uploadCmd = &cobra.Command{
	Use:   "upload [filename]",
	Short: "简单文件上传和查看的HTTP服务",
	Long:  "通过发送一个 multipart/form-data 格式的 POST 请求上传文件;也可用作简单静态资源HTTP服务",
	Example: `上传文件:	curl -F "file=@/path/to/your/another_file.txt" http://localhost:9999/upload
查看已上传:	curl http://localhost:9999/
`,

	//Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		startUploadServer()
		//filename := args[0]
	},
}
var dataPath string
var port string

func init() {
	rootCmd.AddCommand(uploadCmd)

	uploadCmd.Flags().StringVarP(&dataPath, "data", "d", "./data", "自定义上传路径")
	uploadCmd.Flags().StringVarP(&port, "port", "p", "9999", "HTTP端口")
}

func startUploadServer() {
	http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			file, header, err := r.FormFile("file")
			if err != nil {
				http.Error(w, "Error retrieving file", http.StatusInternalServerError)
				return
			}
			defer file.Close()

			err = bio.UploadFile(dataPath, header.Filename)

			if err != nil {
				http.Error(w, "Error copying file", http.StatusInternalServerError)
				return
			}
			fmt.Fprintf(w, "File uploaded successfully to: %s\n", header.Filename)
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	})
	http.Handle("/", http.FileServer(http.Dir(dataPath)))
	fmt.Printf("Upload server started at http://localhost:%v/upload\nFile server started at http://localhost:%v\n", port, port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("Error starting upload server:", err)
		return
	}
}
